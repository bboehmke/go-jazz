package main

import (
	"fmt"
	"strings"
)

// Model of an RTC CCM object
type Model struct {
	// Description of object
	Description []string
	// Name of link in documentation
	LinkRef string

	// Resource identifier of object.
	ResourceID string
	// Identifier of element inside resource.
	ElementID string
	// Identifier of Type
	TypeID string

	// Fields of this object
	Fields []Field
}

// Name of Go object
func (m Model) Name() string {
	if m.ElementID != "" {
		return strings.Title(m.ElementID)
	}

	return m.TypeID[strings.LastIndex(m.TypeID, ".")+1:]
}

func (m Model) String() string {
	return fmt.Sprintf("%s, %s, %s", m.ResourceID, m.ElementID, m.TypeID)
}
