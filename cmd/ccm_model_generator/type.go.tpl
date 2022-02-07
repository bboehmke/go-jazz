
// {{ .Name }} (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#{{ .LinkRef }})
{{- range .Description}}
// {{.}}
{{- end}}
type {{ .Name }} struct {
	BaseObject `jazz_resource:"{{ .ResourceID }}" jazz_type:"{{ .TypeID }}" jazz_element:"{{ .ElementID }}"`
{{range .Fields}}
{{- range .Description}}
	// {{.}}
{{- end}}
	{{ .GoName }} {{ .GoType }} `jazz:"{{ .Name }}"`
{{end}}
}
