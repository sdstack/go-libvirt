package libvirt

import (
	"sync"

	"github.com/godbus/dbus"
)

type Interface struct {
	conn   *Conn
	object dbus.BusObject
	path   dbus.ObjectPath

	sigs  map[<-chan *dbus.Signal]struct{}
	sigmu sync.Mutex

	//Active bool
	//MAC string
	//Name string
}

// NewInterface() TODO
func NewInterface(c *Conn, path dbus.ObjectPath) *Interface {
	m := &Interface{conn: c}
	if path != "" {
		m.object = c.conn.Object("org.libvirt", path)
	} else {
		m.object = c.object
	}
	m.path = c.object.Path()

	m.sigmu.Lock()
	m.sigs = make(map[<-chan *dbus.Signal]struct{})
	m.sigmu.Unlock()

	return m
}

// Create See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceCreate
func (m *Interface) Create(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Interface.Create", 0, flags).Store()
	return
}

// Destroy See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceDestroy
func (m *Interface) Destroy(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Interface.Destroy", 0, flags).Store()
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceGetXMLDesc
func (m *Interface) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.Interface.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceUndefine
func (m *Interface) Undefine() (err error) {
	err = m.object.Call("org.libvirt.Interface.Undefine", 0).Store()
	return
}

// GetActive See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceIsActive
func (m *Interface) GetActive() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Interface", "Active").Store(&v)
	return
}

// GetMAC See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceGetMACString
func (m *Interface) GetMAC() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Interface", "MAC").Store(&v)
	return
}

// GetName See https://libvirt.org/html/libvirt-libvirt-interface.html#virInterfaceGetName
func (m *Interface) GetName() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Interface", "Name").Store(&v)
	return
}
