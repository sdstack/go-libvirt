package libvirt

import "github.com/godbus/dbus"

type NodeDevice struct {
	conn   *Conn
	object dbus.BusObject

	Name   string
	Parent string
}

// NewNodeDevice() TODO
func NewNodeDevice(c *Conn, path dbus.ObjectPath) *NodeDevice {
	m := &NodeDevice{conn: c}
	if path != "" {
		m.object = c.conn.Object("org.libvirt", path)
	} else {
		m.object = c.object
	}
	return m
}

// Destroy See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceDestroy
func (m *NodeDevice) Destroy() (err error) {
	err = m.object.Call("org.libvirt.NodeDevice.Destroy", 0).Store()
	return
}

// Detach See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceDetachFlags
func (m *NodeDevice) Detach(driverName string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.NodeDevice.Detach", 0, driverName, flags).Store()
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceGetXMLDesc
func (m *NodeDevice) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.NodeDevice.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// ListCaps See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceListCaps
func (m *NodeDevice) ListCaps() (names []string, err error) {
	err = m.object.Call("org.libvirt.NodeDevice.ListCaps", 0).Store(&names)
	return
}

// ReAttach See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceReAttach
func (m *NodeDevice) ReAttach() (err error) {
	err = m.object.Call("org.libvirt.NodeDevice.ReAttach", 0).Store()
	return
}

// Reset See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceReset
func (m *NodeDevice) Reset() (err error) {
	err = m.object.Call("org.libvirt.NodeDevice.Reset", 0).Store()
	return
}
