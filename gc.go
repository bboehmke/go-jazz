package jazz

import (
	"fmt"
	"net/url"
)

type GCApplication struct {
	client *Client
}

func (a *GCApplication) Name() string {
	return "Global Configuration Management"
}

func (a *GCApplication) ID() string {
	return "gc"
}

func (a *GCApplication) Client() *Client {
	return a.Client()
}

// GlobalConfiguration of Jazz application
type GlobalConfiguration struct {
	Title string
	URL   string
}

// GlobalConfigs available on server
func (a *GCApplication) GlobalConfigs() ([]*GlobalConfiguration, error) {
	_, xml, err := a.client.SimpleGet(
		"gc/configuration",
		"application/rdf+xml",
		"failed to get global configurations",
		200)
	if err != nil {
		return nil, err
	}

	configs := make([]*GlobalConfiguration, 0, len(xml.Child))
	for _, e := range xml.FindElements("//rdf:Description[dcterms:title]") {
		configs = append(configs, &GlobalConfiguration{
			Title: e.SelectElement("dcterms:title").Text(),
			URL:   e.SelectAttr("rdf:about").Value,
		})
	}
	return configs, nil
}

// GetGlobalConfig by title
func (a *GCApplication) GetGlobalConfig(title string) (*GlobalConfiguration, error) {
	// https://jazz.net/sandbox02-gc/doc/scenario?id=QueryConfigurations
	_, xml, err := a.client.SimpleGet(
		"gc/oslc-query/configurations?oslc.where="+url.QueryEscape(fmt.Sprintf("dcterms:title=\"%s\"", title)),
		"application/rdf+xml",
		"failed to get global configuration",
		200)
	if err != nil {
		return nil, err
	}

	element := xml.FindElement("//rdf:Description[dcterms:title]")
	if element == nil {
		return nil, fmt.Errorf("failed to find global configuration \"%s\"", title)
	}

	return &GlobalConfiguration{
		Title: element.SelectElement("dcterms:title").Text(),
		URL:   element.SelectAttr("rdf:about").Value,
	}, nil
}
