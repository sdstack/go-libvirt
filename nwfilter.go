package libvirt

import "github.com/godbus/dbus"

type NWFilter struct {
	conn   *Conn
	object dbus.BusObject

	Name string
	UUID string
}

// NewNWFilter() TODO
func NewNWFilter(c *Conn, path dbus.ObjectPath) *NWFilter {
	m := &NWFilter{conn: c}
	if path != "" {
		m.object = c.conn.Object("org.libvirt", path)
	} else {
		m.object = c.object
	}
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
