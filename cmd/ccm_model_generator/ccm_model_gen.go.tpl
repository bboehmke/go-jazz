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
// {{ .Name }} (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#{{ .LinkRef }})
{{- range .Description}}
// {{.}}
{{- end}}
type {{ .Name }} struct {
	BaseObject
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
func (o *{{ .Name }}) Spec() *ObjectSpec {
	return &ObjectSpec{
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