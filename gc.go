package jazz

import (
	"strings"

	"go.uber.org/zap"
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

type GlobalConfiguration struct {
	Title string
	URL   string
}

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
	zap.S().Debugf("Found %d global configurations", len(configs))
	return configs, nil
}

func (a *GCApplication) GetGlobalConfig(title string) (*GlobalConfiguration, error) {
	configs, err := a.GlobalConfigs()
	if err != nil {
		return nil, err
	}

	title = strings.ToLower(strings.TrimSpace(title))

	for _, config := range configs {
		if strings.ToLower(strings.TrimSpace(config.Title)) == title {
			return config, nil
		}
	}

	zap.S().Debugf("Global configuration %s not found", title)
	return nil, nil
}
