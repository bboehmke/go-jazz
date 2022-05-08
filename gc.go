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
	"fmt"
	"net/url"
)

// https://jazz.net/sandbox02-gc/doc/scenarios

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
func (a *GCApplication) GlobalConfigs(ctx context.Context) ([]*GlobalConfiguration, error) {
	_, xml, err := a.client.getEtree(ctx,
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
func (a *GCApplication) GetGlobalConfig(ctx context.Context, title string) (*GlobalConfiguration, error) {
	// https://jazz.net/sandbox02-gc/doc/scenario?id=QueryConfigurations
	_, xml, err := a.client.getEtree(ctx,
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
