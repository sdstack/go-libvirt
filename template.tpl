package {{PkgName}}

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type {{ExportName}} struct {
	conn   *dbus.Conn
	object dbus.BusObject
	{{range .Properties}}
	{{.Name}} {{GuessType .Name .Type DbusInterface}}{{end}}
}

// New{{ExportName}}() establishes a connection to the system bus and authenticates.
func New{{ExportName}}() (*{{ExportName}}, error) {
	m := new({{ExportName}})

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *{{ExportName}}) initConnection() error {
	var err error
	m.conn, err = dbus.SystemBusPrivate()
	if err != nil {
		return err
	}

	// Only use EXTERNAL method, and hardcode the uid (not username)
	// to avoid a username lookup (which requires a dynamically linked
	// libc)
	methods := []dbus.Auth{dbus.AuthExternal(strconv.Itoa(os.Getuid()))}

	err = m.conn.Auth(methods)
	if err != nil {
		m.conn.Close()
		return err
	}

	err = m.conn.Hello()
	if err != nil {
		m.conn.Close()
		return err
	}

	m.object = m.conn.Object("{{DbusInterface}}", dbus.ObjectPath("{{DbusPath}}"))

	return nil
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

