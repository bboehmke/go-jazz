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
)

// QMApplication interface
type QMApplication struct {
	client *Client
}

// Name of application
func (a *QMApplication) Name() string {
	return "Quality Management"
}

// ID of application
func (a *QMApplication) ID() string {
	return "qm"
}

// Client instance used for communication
func (a *QMApplication) Client() *Client {
	return a.client
}

// Projects of available (and accessible)
func (a *QMApplication) Projects(ctx context.Context) ([]*QMProject, error) {
	// https://jazz.net/wiki/bin/view/Main/RqmApi#Project_Feed_Service
	entries, err := Chan2List[FeedEntry](func(ch chan FeedEntry) error {
		return a.client.requestFeed(ctx,
			"qm/service/com.ibm.rqm.integration.service.IIntegrationService/projects",
			ch, true)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	projects := make([]*QMProject, len(entries))
	for i, entry := range entries {
		projects[i] = &QMProject{
			Title: entry.Title,
			Alias: entry.Alias,
			qm:    a,
		}
	}
	return projects, nil
}

// GetProject with the given title
func (a *QMApplication) GetProject(ctx context.Context, title string) (*QMProject, error) {
	projects, err := a.Projects(ctx)
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.Title == title {
			return project, nil
		}
	}
	return nil, fmt.Errorf("failed to find project \"%s\"", title)
}
