package jazz

import (
	"fmt"
	"strconv"
	"strings"
)

// QMFilter is used to filter results in QM list queries
type QMFilter map[string]string

type QMObjectSpec struct {
	// Resource identifier of object.
	// https://jazz.net/wiki/bin/view/Main/RqmApi#Resources_and_their_Supported_Op
	ResourceID string
}

func (o *QMObjectSpec) ListURL(proj *QMProject, filter QMFilter) (string, error) {
	filterQuery, err := o.buildFilterQuery(filter)
	if err != nil {
		return "", fmt.Errorf("failed to build filter: %w", err)
	}

	return fmt.Sprintf(
		"qm/service/com.ibm.rqm.integration.service.IIntegrationService/resources/%s/%s%s",
		proj.Alias, o.ResourceID, filterQuery), nil
}

// buildFilterQuery for the given QMFilter
func (o *QMObjectSpec) buildFilterQuery(filter QMFilter) (string, error) {
	if len(filter) == 0 {
		return "", nil
	}

	var filterList []string
	for key, value := range filter {
		filterList = append(filterList, fmt.Sprintf("%s='%s'", key, value))
	}
	return fmt.Sprintf("[%s]", strings.Join(filterList, " and ")), nil
}

// GetURL returns the URL to get an object
//  https://jazz.net/wiki/bin/view/Main/RqmApi#integrationUrl
//  https://jazz.net/wiki/bin/view/Main/RqmApi#single_ProjectFeedUrl
//  https://jazz.net/wiki/bin/view/Main/RqmApi#Resource_Objects_and_their_Relat
func (o *QMObjectSpec) GetURL(proj *QMProject, id string) string {
	if strings.HasPrefix(id, "http://") || strings.HasPrefix(id, "https://") {
		return id
	}

	if _, err := strconv.Atoi(id); err == nil {
		id = fmt.Sprintf("urn:com.ibm.rqm:%s:%s", o.ResourceID, id)
	}

	return fmt.Sprintf(
		"qm/service/com.ibm.rqm.integration.service.IIntegrationService/resources/%s/%s/%s",
		proj.Alias, o.ResourceID, id)
}
