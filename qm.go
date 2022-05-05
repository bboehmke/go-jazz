package jazz

import (
	"fmt"
	"io"
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
func (a *QMApplication) Projects() ([]*QMProject, error) {
	// https://jazz.net/wiki/bin/view/Main/RqmApi#Project_Feed_Service
	entries, err := Chan2List[feedEntry](func(ch chan feedEntry) error {
		return a.client.requestFeed(
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

// NewUUID returns a new UUID generated on the server
func (a *QMApplication) NewUUID() (string, error) {
	response, err := a.client.Get(
		"qm/service/com.ibm.rqm.integration.service.IIntegrationService/UUID/new",
		"application/json",
		true)
	if err != nil {
		return "", fmt.Errorf("failed to get UUID: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to get UUID: %w", err)
	}
	return string(data), nil
}
