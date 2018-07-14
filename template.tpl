package {{PkgName}}

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type {{ExportName}} struct {
	conn   *Conn
	object dbus.BusObject
	path dbus.ObjectPath

  {{if or .Properties .Signals}}
  sigs map[<-chan *dbus.Signal]struct{}
  sigmu sync.Mutex
  {{end}}
	{{range .Properties}}
	//{{.Name}} {{GuessType .Name .Type DbusInterface}}{{end}}
}

// New{{ExportName}}() TODO
func New{{ExportName}}(c *Conn, path dbus.ObjectPath) (*{{ExportName}}) {
	m := &{{ExportName}}{conn: c}
	if path != "" {
	  m.object = c.conn.Object("org.libvirt", path)
  } else {
    m.object = c.object
  }
  m.path = c.object.Path()

  {{if or .Properties .Signals}}
  m.sigmu.Lock()
  m.sigs = make(map[<-chan *dbus.Signal]struct{})
  m.sigmu.Unlock()
  {{end}}
	return m
}

{{range .Signals}}
{{$methodName := .Name}}
{{- range .Annotations}}// Subscribe{{$methodName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) Subscribe{{.Name}}(callback func({{GetParamterOutsProto .Args}})) <-chan *dbus.Signal {
  if callback == nil {
    return nil
  }
  m.sigmu.Lock()
  ch := make(chan *dbus.Signal)
  m.sigs[ch] = struct{}{}
  m.conn.conn.Signal(ch)
  m.sigmu.Unlock()
  m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='{{DbusInterface}}',member='{{.Name}}'")
  go func() {
    for v := range ch {
      if v.Path != m.path || v.Name != "{{DbusInterface}}.{{.Name}}" || {{len .Args}} != len(v.Body) {
        continue
      }
      callback({{range $index, $arg := .Args}}{{if $index}},{{end}}v.Body[{{$index}}].({{GuessType $arg.Name $arg.Type ""}}){{end}})
     }
  }()
  return ch
}

{{ range .Annotations}}// UnSubscribe{{$methodName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) UnSubscribe{{.Name}}(ch <-chan *dbus.Signal) {
  m.sigmu.Lock()
  delete(m.sigs, ch)
  m.sigmu.Unlock()
  m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='{{DbusInterface}}',member='{{.Name}}'")
}
{{end}}

{{range .Methods}}
{{$methodName := .Name}}
{{- range .Annotations}}// {{$methodName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) {{.Name}}({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}{{with GetParamterOuts .Args}}, {{end}}err error) {
	err = m.object.Call("{{DbusInterface}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	return
}
{{end}}

{{range .Properties}}
{{$propName := .Name}}
{{if PropWritable .}}{{range .Annotations}}// Set{{$propName}} {{AnnotationComment .Value}}{{end}}
func (m *{{ExportName}}) Set{{.Name}}(v {{GuessType .Name .Type ""}}) (err error) {
  err = m.object.Call("org.freedesktop.DBus.Properties.Set", 0, "{{DbusInterface}}", "{{.Name}}", dbus.MakeVariant(v)).Store()
  return
}
{{end}}
{{- range .Annotations}}// Get{{$propName}} {{AnnotationComment .Value}}
{{- end}}
func (m *{{ExportName}}) Get{{.Name}}() (v {{GuessType .Name .Type ""}}, err error) {
  err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "{{DbusInterface}}", "{{.Name}}").Store(&v)
  return
}
{{end}}
