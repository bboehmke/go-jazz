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
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
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
		filterList = append(filterList, fmt.Sprintf("%s='%s'", key, url.QueryEscape(value)))
	}
	return fmt.Sprintf("?fields=feed/entry/content/%s[%s]", o.ResourceID, strings.Join(filterList, " and ")), nil
}

// GetURL returns the URL to get an object
//  https://jazz.net/wiki/bin/view/Main/RqmApi#integrationUrl
//  https://jazz.net/wiki/bin/view/Main/RqmApi#single_ProjectFeedUrl
//  https://jazz.net/wiki/bin/view/Main/RqmApi#Resource_Objects_and_their_Relat
func (o *QMObjectSpec) GetURL(proj *QMProject, id string) string {
	if strings.HasPrefix(id, "http://") ||
		strings.HasPrefix(id, "https://") ||
		strings.HasPrefix(id, "qm/service/") {
		return id
	}

	if _, err := strconv.Atoi(id); err == nil {
		id = fmt.Sprintf("urn:com.ibm.rqm:%s:%s", o.ResourceID, id)
	}

	return fmt.Sprintf(
		"qm/service/com.ibm.rqm.integration.service.IIntegrationService/resources/%s/%s/%s",
		proj.Alias, o.ResourceID, id)
}

// DumpXml for update or creation of objects
func (o *QMObjectSpec) DumpXml(obj QMObject) []byte {
	buffer := bytes.NewBuffer(nil)

	// build xml namespace attributes
	ns := map[string]string{
		"qm":       "http://jazz.net/xmlns/alm/qm/v0.1/",
		"alm":      "http://jazz.net/xmlns/alm/v0.1/",
		"qmresult": "http://jazz.net/xmlns/alm/qm/v0.1/executionresult/v0.1",
	}
	var xmlns []string
	for key, value := range ns {
		xmlns = append(xmlns, fmt.Sprintf("xmlns:%s=\"%s\"", key, value))
	}

	_, _ = fmt.Fprintf(buffer, "<qm:%s %s>\n",
		o.ResourceID, strings.Join(xmlns, " "))

	val := reflect.ValueOf(obj).Elem()
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)

		// only handle jazz fields
		fieldName := field.Tag.Get("jazz")
		if fieldName == "" {
			continue
		}

		// add value to dump
		switch v := fieldValue.Interface().(type) {
		case QMRef:
			if v.Href != "" {
				_, _ = fmt.Fprintf(buffer, " <%s href=\"%s\"/>\n", fieldName, v.Href)
			}
		case QMRefList:
			for _, ref := range v {
				if ref.Href != "" {
					_, _ = fmt.Fprintf(buffer, " <%s href=\"%s\"/>\n", fieldName, ref.Href)
				}
			}
		case QMVariableMap:
			_, _ = fmt.Fprintf(buffer, " <%s>\n", fieldName)
			for key, value := range v {
				_, _ = fmt.Fprintf(buffer,
					"  <qm:variable><qm:name>%s</qm:name><qm:value>%s</qm:value></qm:variable>\n",
					key, value)
			}
			_, _ = fmt.Fprintf(buffer, " </%s>\n", fieldName)
		case string:
			if len(v) > 0 {
				_, _ = fmt.Fprintf(buffer, " <%s>%s</%s>\n", fieldName, v, fieldName)
			}
		case time.Time:
			if !v.IsZero() {
				_, _ = fmt.Fprintf(buffer, " <%s>%s</%s>\n", fieldName, v.Format(time.RFC3339), fieldName)
			}

		case fmt.Stringer:
			str := v.String()
			if len(str) > 0 {
				_, _ = fmt.Fprintf(buffer, " <%s>%s</%s>\n", fieldName, str, fieldName)
			}

		case int:
			if v != 0 {
				_, _ = fmt.Fprintf(buffer, " <%s>%d</%s>\n", fieldName, v, fieldName)
			}

		default:
			panic("unknown type")
		}
	}
	_, _ = fmt.Fprintf(buffer, "</qm:%s>\n", o.ResourceID)

	return buffer.Bytes()
}
