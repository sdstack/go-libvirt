package {{PkgName}}

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type {{ExportName}} struct {
	conn   *Conn
	path dbus.ObjectPath
	{{range .Properties}}
	{{.Name}} {{GuessType .Name .Type DbusInterface}}{{end}}
}

// New{{ExportName}}() TODO
func New{{ExportName}}(c *Conn, path dbus.ObjectPath) (*{{ExportName}}) {
	return &{{ExportName}}{
   conn: c,
   path: path,
  }
}

{{range .Methods}}
{{$methodName := .Name}}
{{- range .Annotations}}// {{$methodName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) {{.Name}}({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}{{with GetParamterOuts .Args}}, {{end}}err error) {
	err = m.conn.object.Call("{{DbusInterface}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	return
}
{{end}}

