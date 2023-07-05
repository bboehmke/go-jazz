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
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/spf13/cast"
)

// CCMFilter is used to filter results in CCM list queries
type CCMFilter map[string][]interface{}

// CCMRawFilter creates a filter from a raw query
func CCMRawFilter(query string) CCMFilter {
	return map[string][]interface{}{
		"_raw": {query},
	}
}

// ccmObjectSpecs contains specifications of all supported object types
var ccmObjectSpecs = make(map[string]*CCMObjectSpec)

// ccmRegisterType is used to add an object type to ccmObjectSpecs (called in init())
func ccmRegisterType(obj CCMObject) {
	spec := obj.Spec()
	ccmObjectSpecs[spec.Type.String()] = spec
}

type CCMObjectSpec struct {
	// Resource identifier of object.
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	//
	// * foundation: Common artifacts such as project areas, team areas, contributors, iterations and links
	// * scm: Source Control artifacts such as streams and components, as well as stream sizing deltas
	// * build: Build artifacts such as build results, build result contributions, build definitions, and build engines
	// * apt: Agile Planning artifacts such as team capacity and resource schedules and absences
	// * workitem: Work Item artifacts such as work items, categories, severities, and priorities
	ResourceID string

	// Identifier of element inside resource.
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	ElementID string

	// Identifier of Type
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	TypeID string

	Type reflect.Type
}

// buildFilterQuery for the given CCMFilter
func (o *CCMObjectSpec) buildFilterQuery(filter CCMFilter) (string, error) {
	if len(filter) == 0 {
		return "", nil
	}

	var filterList []string
	for key, values := range filter {
		orFilter := make([]string, len(values))

		// special handling to pass raw filter queries
		if key == "_raw" {
			for i, value := range values {
				orFilter[i] = cast.ToString(value)
			}

		} else {
			// error if field with name does not exist
			field, ok := o.Type.FieldByName(key)
			if !ok {
				return "", fmt.Errorf("no field with name \"%s\"", key)
			}

			// non jazz fields are not supported
			fieldName := field.Tag.Get("jazz")
			if fieldName == "" {
				return "", fmt.Errorf("no field with name \"%s\"", key)
			}

			// special handling for jazz fields
			_, err := CCMLoadObjectSpec(field.Type)
			if err == nil {
				fieldName = fieldName + "/itemId"
			}

			for i, value := range values {
				orFilter[i] = fmt.Sprintf("%s=\"%s\"", fieldName, cast.ToString(value))
			}
		}

		if len(orFilter) == 1 {
			filterList = append(filterList, orFilter[0])
		} else {
			filterList = append(filterList, fmt.Sprintf("(%s)", strings.Join(orFilter, " or ")))
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(filterList, " and ")), nil
}

// CCMLoadObjectSpec from given type
func CCMLoadObjectSpec(t reflect.Type) (*CCMObjectSpec, error) {
	// get inner type of pointer or lists
	if t.Kind() == reflect.Ptr ||
		t.Kind() == reflect.Array ||
		t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Chan {
		return CCMLoadObjectSpec(t.Elem())
	}

	// only structs can be CCM objects
	if t.Kind() != reflect.Struct {
		return nil, errors.New("invalid ccm object given")
	}

	// check if CCM object
	spec, ok := ccmObjectSpecs[t.String()]
	if !ok {
		return nil, errors.New("invalid ccm object given")
	}
	return spec, nil
}

// ListURL returns the URL to get a list of objects
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *CCMObjectSpec) ListURL(filter CCMFilter) (string, error) {
	filterQuery, err := o.buildFilterQuery(filter)
	if err != nil {
		return "", fmt.Errorf("failed to build filter: %w", err)
	}

	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s",
		o.ResourceID,
		url.QueryEscape(fmt.Sprintf("%s/%s%s/(itemId)", o.ElementID, o.ElementID, filterQuery))), nil
}

// GetURL returns the URL to get an object
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *CCMObjectSpec) GetURL(id string) string {
	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s/%s[itemId=%s]/(%s)",
		o.ResourceID, o.ElementID, o.ElementID,
		id,
		strings.Join(o.getLoadFields(o.Type), "|")) // field selector
}

// getLoadFields for the given CCM object type
func (o *CCMObjectSpec) getLoadFields(t reflect.Type) []string {
	fields := make([]string, 0)
	simpleFields := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// skip base object
		if field.Type == CCMBaseObjectType {
			if o.ElementID != "" {
				fields = append(fields, o.getLoadFields(field.Type)...)
			}
			continue
		}

		// only handle jazz fields
		fieldName := field.Tag.Get("jazz")
		if fieldName == "" {
			continue
		}

		// skip non jazz elements
		spec, err := CCMLoadObjectSpec(field.Type)
		if err != nil {
			simpleFields = true
			continue
		}

		// object with an element ID can be loaded later -> only itemId required
		if spec.ElementID != "" {
			fields = append(fields, fieldName+"/itemId")
			continue
		}

		subFields := spec.getLoadFields(spec.Type)
		if len(subFields) > 1 {
			fields = append(fields,
				fmt.Sprintf("%s/(%s)", fieldName,
					strings.Join(subFields, "|")))
		} else {
			fields = append(fields,
				fmt.Sprintf("%s/%s", fieldName, subFields[0]))
		}
	}
	if simpleFields {
		fields = append(fields, "*")
	}

	// make fields unique
	seen := make(map[string]struct{}, len(fields))
	j := 0
	for _, v := range fields {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		fields[j] = v
		j++
	}
	return fields[:j]
}

// Load object from XML element
func (o *CCMObjectSpec) Load(ccm *CCMApplication, value reflect.Value, element *etree.Element) error {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			bla := value.Interface()
			_ = bla
			value.Set(reflect.New(value.Type().Elem()))
		}

		err := o.Load(ccm, value.Elem(), element)
		if err != nil {
			return err
		}
		return nil

	case reflect.Slice:
		valueType := value.Type()
		sliceValue := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), 0, 0)

		elemType := valueType.Elem()
		for _, child := range element.ChildElements() {
			v := reflect.New(elemType)
			err := o.Load(ccm, v, child)
			if err != nil {
				return err
			}
			reflect.Append(sliceValue, v)
		}
		value.Set(sliceValue)
		return nil

	case reflect.Struct:
		return o.loadFields(ccm, value, element)

	default:
		panic("unsupported type")
	}
}

// loadFields of object from XML element
func (o *CCMObjectSpec) loadFields(ccm *CCMApplication, obj reflect.Value, element *etree.Element) error {
	t := obj.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := obj.FieldByName(field.Name)

		var err error
		if field.Type == CCMBaseObjectType {
			// add ccm instance to CCM objects
			if initObj, ok := value.Addr().Interface().(*CCMBaseObject); ok {
				initObj.setCCM(ccm)
			}
			err = o.loadFields(ccm, value, element)
		} else {
			tag := field.Tag.Get("jazz")

			// skip fields without tag
			if tag == "" {
				continue
			}

			// skip empty
			fieldElements := element.SelectElements(tag)
			if len(fieldElements) == 0 {
				continue
			}

			if value.Kind() == reflect.Slice {
				err = o.loadListValue(ccm, value, field.Type, fieldElements)
			} else if len(fieldElements[0].Child) > 0 {
				err = o.loadValue(ccm, value, field.Type, fieldElements[0])
			}

		}
		if err != nil {
			return fmt.Errorf("failed to load %s: %w", field.Name, err)
		}
	}
	return nil
}

// loadListValue from XML element
func (o *CCMObjectSpec) loadListValue(ccm *CCMApplication, value reflect.Value, valueType reflect.Type, element []*etree.Element) error {
	objList := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), 0, len(element))

	for _, e := range element {
		v := reflect.New(valueType.Elem()).Elem()

		err := o.loadValue(ccm, v, valueType.Elem(), e)
		if err != nil {
			return err
		}
		objList = reflect.Append(objList, v)
	}

	value.Set(objList)
	return nil
}

// loadValue from XML element
func (o *CCMObjectSpec) loadValue(ccm *CCMApplication, value reflect.Value, valueType reflect.Type, element *etree.Element) error {
	switch valueType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(cast.ToInt64(element.Text()))
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(cast.ToUint64(element.Text()))
		return nil

	case reflect.Float32, reflect.Float64:
		value.SetFloat(cast.ToFloat64(element.Text()))
		return nil

	case reflect.String:
		value.SetString(element.Text())
		return nil

	case reflect.Bool:
		value.SetBool(cast.ToBool(element.Text()))
		return nil

	case reflect.Struct:
		if valueType == reflect.TypeOf(time.Time{}) {
			parsedTime, err := time.Parse("2006-01-02T15:04:05.000-0700", element.Text())
			if err != nil {
				return err
			}

			value.Set(reflect.ValueOf(parsedTime))
			return nil

		} else if _, ok := valueType.FieldByName("CCMBaseObject"); ok {
			spec, err := CCMLoadObjectSpec(valueType)
			if err != nil {
				return err
			}
			return spec.Load(ccm, value, element)
		} else {
			panic("unknown type")
		}
	case reflect.Ptr:
		value.Set(reflect.New(valueType.Elem()))
		return o.loadValue(ccm, value.Elem(), valueType.Elem(), element)

	default:
		panic("unknown type")
	}
}
