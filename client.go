// Copyright 2022 Benjamin Böhmke <benjamin@boehmke.net>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jazz

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/beevik/etree"
	"go.uber.org/zap"
	"golang.org/x/net/publicsuffix"
)

// Client to communicate with various application of a jazz server
type Client struct {
	// HttpClient used for requests to the jazz server
	HttpClient *http.Client

	// maximum amount of actions that should run in parallel
	Worker int

	// GC provides all functionalities to access the "Global Configuration Management" application
	GC *GCApplication
	// CCM provides all functionalities to access the "Change and Configuration Management"  application
	CCM *CCMApplication
	// QM provides all functionalities to access the "Quality Management"  application
	QM *QMApplication

	// GlobalConfiguration used for API requests
	configContext *GlobalConfiguration

	// Logger instance used for debug logging
	Logger *zap.Logger

	baseUrl   string
	user      string
	password  string
	basicAuth bool
}

// NewClient creates a new client for the given server
func NewClient(baseUrl, user, password string) (*Client, error) {
	// cookie store
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookiejar: %w", err)
	}

	// ensure base url ends with /
	baseUrl = strings.TrimSpace(baseUrl)

	client := &Client{
		HttpClient: &http.Client{
			Jar: jar,
		},
		baseUrl: baseUrl,
		user:    user,
		// hide password in debugger
		password: base64.StdEncoding.EncodeToString([]byte(password)),
		Worker:   20,

		Logger: zap.NewNop(),
	}

	// register applications
	client.GC = &GCApplication{client: client}
	client.CCM = &CCMApplication{client: client}
	client.QM = &QMApplication{client: client}

	return client, nil
}

// WithConfig creates a new client (copy of existing) with the given global configuration
func (c *Client) WithConfig(config *GlobalConfiguration) *Client {
	client := &Client{
		HttpClient: c.HttpClient,
		baseUrl:    c.baseUrl,
		user:       c.user,
		password:   c.password,
		Worker:     c.Worker,
		basicAuth:  c.basicAuth,

		Logger: c.Logger,

		configContext: config,
	}

	// register applications
	client.GC = &GCApplication{client: client}
	client.CCM = &CCMApplication{client: client}
	client.QM = &QMApplication{client: client}

	return client
}

// buildUrl for the given path.
// If path is already a complete URL return the value.
func (c *Client) buildUrl(path string) string {
	if !strings.HasPrefix(path, "http:") && !strings.HasPrefix(path, "https:") {
		return c.baseUrl + path
	}
	return path
}

// getEtree send a GET HTTP request (same as get) and return content as XML etree
func (c *Client) getEtree(ctx context.Context, url, contentType, errorMessage string, statusCode int) (*http.Response, *etree.Element, error) {
	response, err := c.get(ctx, url, contentType, false)
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

// get sends GET request to server
func (c *Client) get(ctx context.Context, url, contentType string, noGc bool) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", c.buildUrl(url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get request: %w", err)
	}
	request.Header.Set("Content-type", contentType)

	return c.sendRequest(request, noGc)
}

// put sends GET request to server
func (c *Client) put(ctx context.Context, url, contentType string, reader io.Reader) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, "PUT", c.buildUrl(url), reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create put request: %w", err)
	}
	request.Header.Set("Content-type", contentType)

	return c.sendRequest(request, false)
}

// sendRequest to server and handle auth if required
func (c *Client) sendRequest(request *http.Request, noGc bool) (*http.Response, error) {
	// send request
	response, err := c.sendRawRequest(request, true, noGc)
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

		c.Logger.Sugar().Debugf("Login to %s as %s (Form challenge)", c.baseUrl, c.user)

		// send the login request
		values := make(url.Values)
		values.Set("j_username", c.user)
		pass, _ := base64.StdEncoding.DecodeString(c.password)
		values.Set("j_password", string(pass))

		jtsRequest, err := http.NewRequestWithContext(request.Context(), "GET", c.buildUrl("jts/j_security_check?"+values.Encode()), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create JTS request: %w", err)
		}
		response, err = c.sendRawRequest(jtsRequest, false, noGc)
		// close response as it is not used
		_ = response.Body.Close()

		// if header is still set the login failed
		authMsg = response.Header.Get("x-com-ibm-team-repository-web-auth-msg")
		if authMsg != "" {
			return nil, fmt.Errorf("server authentication failed: %s", authMsg)
		}

		// resend original request
		return c.sendRawRequest(request, true, noGc)
	}

	// > basic auth
	if response.StatusCode == 401 && response.Header.Get("www-authenticate") != "" {
		// close response as it is not used
		_ = response.Body.Close()
		c.Logger.Sugar().Debugf("Login to %s as %s (basic auth)", c.baseUrl, c.user)

		c.basicAuth = true

		// resend original request
		response, err = c.sendRawRequest(request, true, noGc)
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

// sendRawRequest to server
func (c *Client) sendRawRequest(request *http.Request, log, noGc bool) (*http.Response, error) {
	if log {
		c.Logger.Sugar().Debugf("Send %s request to %s", request.Method, request.URL)
	}

	if request.Header.Get("Accept") == "" {
		request.Header.Set("Accept", request.Header.Get("Content-Type"))
	}

	// normally only required for OSLC request but has no effect for others
	request.Header.Set("OSLC-Core-Version", "2.0")

	// set configuration context if enabled and set
	if !noGc && c.configContext != nil {
		request.Header.Set("Configuration-Context", c.configContext.URL)
	}

	// only add basic auth data if it was enabled
	if c.basicAuth {
		pass, _ := base64.StdEncoding.DecodeString(c.password)
		request.SetBasicAuth(c.user, string(pass))
	}

	return c.HttpClient.Do(request)
}
