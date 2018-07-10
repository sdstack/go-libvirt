package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type NWFilter struct {
	conn   *dbus.Conn
	object dbus.BusObject

	Name string
	UUID string
}

// NewNWFilter() establishes a connection to the system bus and authenticates.
func NewNWFilter() (*NWFilter, error) {
	m := new(NWFilter)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *NWFilter) initConnection() error {
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

	m.object = m.conn.Object("org.libvirt.NWFilter", dbus.ObjectPath("/org/libvirt/nwfilter"))

	return nil
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterGetXMLDesc
func (m *NWFilter) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.NWFilter.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterUndefine
func (m *NWFilter) Undefine() (err error) {
	err = m.object.Call("org.libvirt.NWFilter.Undefine", 0).Store()
	return
}
