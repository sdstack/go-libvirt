package libvirt

import (
	"sync"

	"github.com/godbus/dbus"
)

type NWFilter struct {
	conn   *Conn
	object dbus.BusObject
	path   dbus.ObjectPath

	sigs  map[<-chan *dbus.Signal]struct{}
	sigmu sync.Mutex

	//Name string
	//UUID string
}

// NewNWFilter() TODO
func NewNWFilter(c *Conn, path dbus.ObjectPath) *NWFilter {
	m := &NWFilter{conn: c}
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

// GetName See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterGetName// GetName const
func (m *NWFilter) GetName() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.NWFilter", "Name").Store(&v)
	return
}

// GetUUID See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterGetUUIDString// GetUUID const
func (m *NWFilter) GetUUID() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.NWFilter", "UUID").Store(&v)
	return
}
