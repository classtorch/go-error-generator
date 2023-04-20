package internal

import (
	"bytes"
	"text/template"
)

var importsTemplate = `
import (
{{ range .Imports }}
	{{.Alias}}    "{{.PackagePath}}"
{{- end }}
)
`

var errorsTemplate = `
var (
{{ range .Errors }}
	{{.EnumValue}}            = &errors.{{ .ErrorStructName }}{ {{ .CodeFieldName }}: {{ .Code }}, {{ .MsgFieldName }}: "{{ .Msg }}"} //{{ .Msg }}
{{- end }}
)
`

var mapsTemplate = `
{{ $errorStructName := .ErrorStructName }}

var (
{{ range $key, $value := .MapInfo }}
  {{ $key }} = map[int32]*errors.{{ $errorStructName }}{
    {{- range $value }}
         {{ .Code }}: &errors.{{ $errorStructName }}{ {{ .CodeFieldName }}: {{ .Code }}, {{ .MsgFieldName }}: "{{ .Msg }}"},
    {{- end }}
  }
{{ end }}
)
`

type ErrorInfo struct {
	ErrorStructName string
	EnumName        string
	EnumValue       string
	CodeFieldName   string
	Code            int32
	MsgFieldName    string
	Msg             string
}

type ErrorWrapper struct {
	Errors []ErrorInfo
}

type MapWrapper struct {
	ErrorStructName string
	MapInfo         map[string][]ErrorInfo
}

type ImportPackInfo struct {
	Alias       string
	PackagePath string
}

type ImportWrapper struct {
	Imports []ImportPackInfo
}

func (e *ErrorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}

func (m *MapWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("maps").Parse(mapsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, m); err != nil {
		panic(err)
	}
	return buf.String()
}

func (i *ImportWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("imports").Parse(importsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, i); err != nil {
		panic(err)
	}
	return buf.String()
}
