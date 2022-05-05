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
		return "CCM" + strings.Title(m.ElementID)
	}

	return "CCM" + m.TypeID[strings.LastIndex(m.TypeID, ".")+1:]
}

func (m Model) String() string {
	return fmt.Sprintf("%s, %s, %s", m.ResourceID, m.ElementID, m.TypeID)
}
func (m Model) IsLoadable() bool {
	return m.ElementID != ""
}

func (m Model) CCMFields() []Field {
	fields := make([]Field, 0, len(m.Fields))
	for _, field := range m.Fields {
		if field.IsCCMType() {
			fields = append(fields, field)
		}
	}
	return fields
}
