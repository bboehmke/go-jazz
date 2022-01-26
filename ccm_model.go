package jazz

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/beevik/etree"
	"github.com/spf13/cast"
)

var BaseObjectType = reflect.TypeOf(BaseObject{})

type BaseObject struct {
	// Common fields of every object
	// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Common_properties

	init sync.Once
	ccm  *CCMApplication

	//  The UUID representing the item in storage. This is technically an internal detail, and resources should mostly be referred to by their unique URLs. In some cases the itemId may be the only unique identifier, however.
	ItemId string `jazz:"itemId"`

	// An MD5 hash of the URI for this element
	UniqueId string `jazz:"uniqueId"`

	// The UUID of the state for this item in storage. This is an internal detail.
	StateId string `jazz:"stateId"`

	// The UUID of a context object used for read access. This is an internal detail.
	ContextId string `jazz:"contextId"`

	// The timestamp of the last modification date of this resource.
	Modified time.Time `jazz:"modified"`

	// A boolean indicating whether or not the resource is "archived". Archived resources are typically hidden from the UI and filtered out of queries.
	Archived bool `jazz:"archived"`

	ReportableUrl string `jazz:"reportableUrl"`

	ModifiedBy *Contributor `jazz:"modifiedBy"`
}

type ObjectSpec struct {
	// Resource identifier of object. Should be overridden
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	//
	// * foundation: Common artifacts such as project areas, team areas, contributors, iterations and links
	// * scm: Source Control artifacts such as streams and components, as well as stream sizing deltas
	// * build: Build artifacts such as build results, build result contributions, build definitions, and build engines
	// * apt: Agile Planning artifacts such as team capacity and resource schedules and absences
	// * workitem: Work Item artifacts such as work items, categories, severities, and priorities
	ResourceID string

	// Identifier of element inside resource. Should be overridden
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	ElementID string

	// Identifier of Type
	// # https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Resources_provided_by_RTC
	TypeID string

	Type reflect.Type
}

func LoadObjectSpec(model interface{}) (*ObjectSpec, error) {
	t, ok := model.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(model)
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("invalid model type")
	}

	field, _ := t.FieldByName("BaseObject")
	ids := strings.Split(field.Tag.Get("jazz"), ",")

	if len(ids) < 2 {
		return nil, fmt.Errorf("invalid object spec: %s", field.Tag.Get("jazz"))
	}

	spec := ObjectSpec{
		ResourceID: ids[0],
		TypeID:     ids[1],
		Type:       t,
	}

	if len(ids) > 2 {
		spec.ElementID = ids[2]
	}

	return &spec, nil
}

// ListURL returns the URL to get a list of objects
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *ObjectSpec) ListURL() string {
	// TODO filter
	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s/%s/(itemId)",
		o.ResourceID, o.ElementID, o.ElementID)
}

// GetURL returns the URL to get a object
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *ObjectSpec) GetURL(id string) string {
	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s/%s[itemId=%s]/(%s)",
		o.ResourceID, o.ElementID, o.ElementID,
		id,
		strings.Join(getLoadFields(o.Type), "|")) // field selector
}

func getLoadFields(t reflect.Type) []string {
	fields := make([]string, 0)
	simpleFields := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// skip base object
		if field.Type == BaseObjectType {
			fields = append(fields, getLoadFields(field.Type)...)
			continue
		}

		// only handle jazz fields
		fieldName := field.Tag.Get("jazz")
		if fieldName == "" {
			continue
		}

		// skip non jazz elements
		spec, err := LoadObjectSpec(field.Type)
		if err != nil {
			simpleFields = true
			continue
		}

		// object with an element ID can be loaded later -> only itemId required
		if spec.ElementID != "" {
			fields = append(fields, fieldName+"/itemId")
			continue
		}

		for _, field := range getLoadFields(spec.Type) {
			fields = append(fields, fieldName+"/"+field)
		}
	}
	if simpleFields {
		fields = append(fields, "*")
	}

	return SliceUniqString(fields)
}

func SliceUniqString(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func (o *ObjectSpec) NewList() interface{} {
	return reflect.MakeSlice(reflect.SliceOf(o.Type), 0, 0).Interface()
}
func (o *ObjectSpec) Load(element *etree.Element) (interface{}, error) {
	objPtr := reflect.New(o.Type)
	err := o.loadFields(objPtr.Elem(), element.FindElement(o.ElementID))
	return objPtr.Interface(), err
}

func (o *ObjectSpec) loadFields(obj reflect.Value, element *etree.Element) error {
	t := obj.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := obj.FieldByName(field.Name)

		var err error
		if field.Type == BaseObjectType {
			err = o.loadFields(value, element)
		} else {
			tag := field.Tag.Get("jazz")

			// skip fields without tag
			if tag == "" {
				continue
			}

			// skip empty
			fieldElement := element.SelectElement(tag)
			if len(fieldElement.Child) == 0 {
				continue
			}

			var data interface{}
			data, err = parseValue(value.Type(), fieldElement)
			if err == nil {
				value.Set(reflect.ValueOf(data))
			}
		}
		if err != nil {
			return fmt.Errorf("failed to load %s: %w", field.Name, err)
		}
	}
	return nil
}

func parseValue(t reflect.Type, element *etree.Element) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cast.ToInt64E(element.Text())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cast.ToUint64E(element.Text())

	case reflect.String:
		return element.Text(), nil

	case reflect.Bool:
		return cast.ToBoolE(element.Text())

	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return time.Parse("2006-01-02T15:04:05.000-0700", element.Text())

		} else if _, ok := t.FieldByName("BaseObject"); ok {
			spec, err := LoadObjectSpec(t)
			if err != nil {
				return nil, err
			}
			return spec.Load(element)
		} else {
			panic("unknown type")
		}
	case reflect.Ptr:
		return parseValue(t.Elem(), element)

	default:
		panic("unknown type")
	}
}
