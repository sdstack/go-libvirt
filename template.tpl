package {{PkgName}}

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type {{ExportName}} struct {
	conn   *Conn
	object dbus.BusObject
	{{range .Properties}}
	{{.Name}} {{GuessType .Name .Type DbusInterface}}{{end}}
}

// New{{ExportName}}() TODO
func New{{ExportName}}(c *Conn, path dbus.ObjectPath) (*{{ExportName}}) {
	m := &{{ExportName}}{conn: c}
	if path != "" {
	m.object = c.conn.Object("org.libvirt", path)
  } else {
  m.object = c.object
  }
	return m
}

{{range .Methods}}
{{$methodName := .Name}}
{{- range .Annotations}}// {{$methodName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) {{.Name}}({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}{{with GetParamterOuts .Args}}, {{end}}err error) {
	err = m.object.Call("{{DbusInterface}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	return
}
{{end}}

