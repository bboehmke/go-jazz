package jazz

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/beevik/etree"
	"go.uber.org/zap"
	"golang.org/x/net/publicsuffix"
)

type Client struct {
	client *http.Client

	baseUrl   string
	user      string
	password  string
	basicAuth bool

	GC *GCApplication
}

func NewClient(baseUrl, user, password string) (*Client, error) {
	// cookie store
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookiejar: %w", err)
	}

	// ensure base url ends with /
	baseUrl = strings.TrimSpace(baseUrl)
	/*if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}*/

	client := &Client{
		client: &http.Client{
			Jar: jar,
		},
		baseUrl: baseUrl,
		user:    user,
		// hide password in debugger
		password: base64.StdEncoding.EncodeToString([]byte(password)),
	}

	client.GC = &GCApplication{client: client}

	return client, nil
}

func (c *Client) buildUrl(url string) string {
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		return c.baseUrl + url
	}
	return url
}

func (c *Client) SimpleGet(url, contentType, errorMessage string, statusCode int) (*http.Response, *etree.Element, error) {
	response, err := c.Get(url, contentType)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", errorMessage, err)
	}
	defer response.Body.Close()

	if statusCode != 0 && response.StatusCode != statusCode {
		return nil, nil, errors.New(errorMessage)
	}

	doc := etree.NewDocument()
	_, err = doc.ReadFrom(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	return response, doc.Root(), nil
}

func (c *Client) Get(url, contentType string) (*http.Response, error) {
	request, err := http.NewRequest("GET", c.buildUrl(url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get request: %w", err)
	}
	request.Header.Set("Content-type", contentType)

	return c.sendRequest(request)
}

func (c *Client) sendRequest(request *http.Request) (*http.Response, error) {
	// send request
	response, err := c.sendRawRequest(request, true)
	if err != nil {
		return nil, err
	}

	// check if auth is required
	// https://jazz.net/wiki/bin/view/Main/NativeClientAuthentication
	// > form challenge
	authMsg := response.Header.Get("x-com-ibm-team-repository-web-auth-msg")
	if authMsg != "" {
		// close response as it is not used
		_ = response.Body.Close()

		// if content of header if not "authrequired" there is an error
		if authMsg != "authrequired" {
			return nil, fmt.Errorf("server authentication error: %s", authMsg)
		}

		zap.S().Debugf("Login to %s as %s (Form challenge)", c.baseUrl, c.user)

		// send the login request
		values := make(url.Values)
		values.Set("j_username", c.user)
		pass, _ := base64.StdEncoding.DecodeString(c.password)
		values.Set("j_password", string(pass))

		jtsRequest, err := http.NewRequest("GET", c.buildUrl("jts/j_security_check?"+values.Encode()), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create JTS request: %w", err)
		}
		response, err = c.sendRawRequest(jtsRequest, false)
		// close response as it is not used
		_ = response.Body.Close()

		// if header is still set the login failed
		authMsg = response.Header.Get("x-com-ibm-team-repository-web-auth-msg")
		if authMsg != "" {
			return nil, fmt.Errorf("server authentication failed: %s", authMsg)
		}

		// resend original request
		return c.sendRawRequest(request, true)
	}

	// > basic auth
	if response.StatusCode == 401 && response.Header.Get("www-authenticate") != "" {
		// close response as it is not used
		_ = response.Body.Close()
		zap.S().Debugf("Login to %s as %s (basic auth)", c.baseUrl, c.user)

		c.basicAuth = true

		// resend original request
		response, err = c.sendRawRequest(request, true)
		if err != nil {
			return nil, err
		}
		if response.StatusCode == 401 {
			return nil, errors.New("server authentication failed")
		}

		return response, nil
	}

	//  unknown auth method
	if response.StatusCode == 401 {
		// close response as it is not used
		_ = response.Body.Close()
		return nil, fmt.Errorf("unknown auth method")
	}
	return response, nil
}

func (c *Client) sendRawRequest(request *http.Request, log bool) (*http.Response, error) {
	if log {
		zap.S().Debugf("Send get request to %s", request.URL)
	}

	if request.Header.Get("Accept") == "" {
		request.Header.Set("Accept", request.Header.Get("Content-Type"))
	}

	// normally only required for OSLC request but has no effect for others
	request.Header.Set("OSLC-Core-Version", "2.0")

	// TODO global config

	// only add basic auth data if it was enabled
	if c.basicAuth {
		pass, _ := base64.StdEncoding.DecodeString(c.password)
		request.SetBasicAuth(c.user, string(pass))
	}

	return c.client.Do(request)
}
