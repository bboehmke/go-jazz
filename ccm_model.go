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

type CCMObject interface {
	Spec() *ObjectSpec
}

type CCMLoadableObject interface {
	CCMObject
	Load() error
}

type initObject interface {
	setCCM(ccm *CCMApplication)
}

var ccmObjects = make(map[string]*ObjectSpec)

func ccmRegisterType(obj CCMObject) {
	spec := obj.Spec()
	ccmObjects[spec.Type.String()] = spec
}

type BaseObject struct {
	// Common fields of every object
	// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Common_properties

	init sync.Once
	ccm  *CCMApplication

	// The UUID representing the item in storage. This is technically an internal
	// detail, and resources should mostly be referred to by their unique URLs.
	// In some cases the itemId may be the only unique identifier, however.
	ItemId string `jazz:"itemId"`

	// An MD5 hash of the URI for this element
	UniqueId string `jazz:"uniqueId"`

	// The UUID of the state for this item in storage. This is an internal detail.
	StateId string `jazz:"stateId"`

	// The UUID of a context object used for read access. This is an internal detail.
	ContextId string `jazz:"contextId"`

	// The timestamp of the last modification date of this resource.
	Modified *time.Time `jazz:"modified"`

	// A boolean indicating whether the resource is "archived". Archived
	// resources are typically hidden from the UI and filtered out of queries.
	Archived bool `jazz:"archived"`

	ReportableUrl string `jazz:"reportableUrl"`

	ModifiedBy *Contributor `jazz:"modifiedBy"`
}

// String returns the ItemId of this object (used for filter)
func (o *BaseObject) String() string {
	return o.ItemId
}

// setCCM application used for read and write actions
func (o *BaseObject) setCCM(ccm *CCMApplication) {
	o.ccm = ccm
}

// loadFields of the given object
func (o *BaseObject) loadFields(fields ...interface{}) error {
	for _, field := range fields {
		if fields, ok := field.([]CCMLoadableObject); ok {
			for _, f := range fields {
				if err := f.Load(); err != nil {
					return err
				}
			}
		} else if f, ok := field.(CCMLoadableObject); ok {
			if reflect.ValueOf(f).IsNil() {
				continue
			}
			if err := f.Load(); err != nil {
				return err
			}
		}
	}
	return nil
}

type ObjectSpec struct {
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

// LoadObjectSpec from given type
func LoadObjectSpec(t reflect.Type) (*ObjectSpec, error) {
	// get inner type of pointer or lists
	if t.Kind() == reflect.Ptr ||
		t.Kind() == reflect.Array ||
		t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Chan {
		return LoadObjectSpec(t.Elem())
	}

	// only structs can be CCM objects
	if t.Kind() != reflect.Struct {
		return nil, errors.New("invalid ccm object given")
	}

	// check if CCM object
	spec, ok := ccmObjects[t.String()]
	if !ok {
		return nil, errors.New("invalid ccm object given")
	}
	return spec, nil
}

// ListURL returns the URL to get a list of objects
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *ObjectSpec) ListURL() string {
	// TODO filter
	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s/%s/(itemId)",
		o.ResourceID, o.ElementID, o.ElementID)
}

// GetURL returns the URL to get an object
// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Examples
func (o *ObjectSpec) GetURL(id string) string {
	return fmt.Sprintf(
		"ccm/rpt/repository/%s?fields=%s/%s[itemId=%s]/(%s)",
		o.ResourceID, o.ElementID, o.ElementID,
		id,
		strings.Join(o.getLoadFields(o.Type), "|")) // field selector
}

// getLoadFields for the given CCM object type
func (o *ObjectSpec) getLoadFields(t reflect.Type) []string {
	fields := make([]string, 0)
	simpleFields := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// skip base object
		if field.Type == BaseObjectType {
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

		for _, subFieldName := range spec.getLoadFields(spec.Type) {
			fields = append(fields, fieldName+"/"+subFieldName)
		}
	}
	if simpleFields {
		fields = append(fields, "*")
	}

	return sliceUniqString(fields)
}

// sliceUniqString removes duplicates from string slice
func sliceUniqString(s []string) []string {
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

// Load object from XML element
func (o *ObjectSpec) Load(ccm *CCMApplication, value reflect.Value, element *etree.Element) error {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
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
func (o *ObjectSpec) loadFields(ccm *CCMApplication, obj reflect.Value, element *etree.Element) error {
	t := obj.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := obj.FieldByName(field.Name)

		var err error
		if field.Type == BaseObjectType {
			// add ccm instance to CCM objects
			if initObj, ok := value.Addr().Interface().(initObject); ok {
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
				err = loadListValue(ccm, value, field.Type, fieldElements)
			} else if len(fieldElements[0].Child) > 0 {
				err = loadValue(ccm, value, field.Type, fieldElements[0])
			}

		}
		if err != nil {
			return fmt.Errorf("failed to load %s: %w", field.Name, err)
		}
	}
	return nil
}

// loadListValue from XML element
func loadListValue(ccm *CCMApplication, value reflect.Value, valueType reflect.Type, element []*etree.Element) error {
	objList := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), 0, len(element))

	for _, e := range element {
		v := reflect.New(valueType.Elem()).Elem()

		err := loadValue(ccm, v, valueType.Elem(), e)
		if err != nil {
			return err
		}
		objList = reflect.Append(objList, v)
	}

	value.Set(objList)
	return nil
}

// loadValue from XML element
func loadValue(ccm *CCMApplication, value reflect.Value, valueType reflect.Type, element *etree.Element) error {
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

		} else if _, ok := valueType.FieldByName("BaseObject"); ok {
			spec, err := LoadObjectSpec(valueType)
			if err != nil {
				return err
			}
			return spec.Load(ccm, value, element)
		} else {
			panic("unknown type")
		}
	case reflect.Ptr:
		value.Set(reflect.New(valueType.Elem()))
		return loadValue(ccm, value.Elem(), valueType.Elem(), element)

	default:
		panic("unknown type")
	}
}
