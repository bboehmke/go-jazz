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

// Code generated! DO NOT EDIT

import (
	"reflect"
	"time"
)

func init() {
{{- range .}}
	ccmRegisterType(new({{ .Name }}))
{{- end }}
}
{{ range .}}

{{- if .LinkRef }}
// {{ .Name }} (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#{{ .LinkRef }})
{{- end }}
{{- range .Description}}
// {{.}}
{{- end}}
type {{ .Name }} struct {
	CCMBaseObject
{{- range .Fields}}
{{ range .Description}}
	// {{.}}
{{- end}}
	{{ .GoName }} {{ .GoType }} `jazz:"{{ .Name }}"`
{{- end}}
}

// {{ .Name }}Type contains the reflection type of {{ .Name }}
var go{{ .Name }}Type = reflect.TypeOf({{ .Name }}{})

// Spec returns the specification object for {{ .Name }}
func (o *{{ .Name }}) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "{{ .ResourceID }}",
		ElementID:  "{{ .ElementID }}",
		TypeID:     "{{ .TypeID }}",
		Type:       go{{ .Name }}Type,
	}
}
{{ if .IsLoadable }}
// Load {{ .Name }} object
func (o *{{ .Name }}) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}
{{ end }}
// LoadAllFields of {{ .Name }} object
func (o *{{ .Name }}) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	{{- range .CCMFields}}
		o.{{ .GoName }},
	{{- end}}
	)
}
{{ end }}