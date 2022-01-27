package main

import (
	"regexp"
	"strings"
)

// regex for handling list fields
var listFieldRegEx = regexp.MustCompile(`, *maxOccurs.*`)

// Field of an RTC CCM object
type Field struct {
	// Name of field in API
	Name string

	// Type of field in API
	Type string

	// Description of field
	Description []string
}

// GoName returns the name for the Go struct field
func (f Field) GoName() string {
	return strings.Title(f.Name)
}

// GoType returns the type for the Go struct field
func (f Field) GoType() string {
	// remove note about list entries from type
	t := listFieldRegEx.ReplaceAllString(f.Type, "")

	// set slice prefix
	var prefix string
	if t != f.Type {
		prefix = "[]"
	}

	// replace invalid types
	if fixedType, ok := invalidTypes[t]; ok {
		t = fixedType
	}

	// check if the type is an object
	if model, ok := modelTypeRef[t]; ok {
		return prefix + "*" + model.Name()
	}
	if name, ok := preDefinedTypes[t]; ok {
		return prefix + "*" + name
	}

	// handle basic type
	switch t {
	case "xs:string":
		return prefix + "string"
	case "xs:time", "xs:date":
		return prefix + "time.Time"
	case "xs:boolean":
		return prefix + "bool"
	case "xs:integer":
		return prefix + "int"
	case "xs:double", "xs:decimal":
		return prefix + "float64"
	case "xs:long":
		return prefix + "int64"
	default:
		panic("unknown type: " + t)
	}
}

func (f Field) String() string {
	return f.Name
}
