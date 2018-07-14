package libvirt

import (
	"sync"

	"github.com/godbus/dbus"
)

type Network struct {
	conn   *Conn
	object dbus.BusObject
	path   dbus.ObjectPath

	sigs  map[<-chan *dbus.Signal]struct{}
	sigmu sync.Mutex

	//Active bool
	//Autostart bool
	//Name string
	//Persistent bool
	//UUID string
}

// NewNetwork() TODO
func NewNetwork(c *Conn, path dbus.ObjectPath) *Network {
	m := &Network{conn: c}
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

// Create See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkCreate
func (m *Network) Create() (err error) {
	err = m.object.Call("org.libvirt.Network.Create", 0).Store()
	return
}

// Destroy See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkDestroy
func (m *Network) Destroy() (err error) {
	err = m.object.Call("org.libvirt.Network.Destroy", 0).Store()
	return
}

// GetDHCPLeases See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetDHCPLeases Empty string can be used to pass a NULL as @mac argument. Empty string will be returned in output for NULL variables.
func (m *Network) GetDHCPLeases(mac string, flags uint32) (leases []interface{}, err error) {
	err = m.object.Call("org.libvirt.Network.GetDHCPLeases", 0, mac, flags).Store(&leases)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetXMLDesc
func (m *Network) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.Network.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkUndefine
func (m *Network) Undefine() (err error) {
	err = m.object.Call("org.libvirt.Network.Undefine", 0).Store()
	return
}

// Update See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkUpdate
func (m *Network) Update(command uint32, section uint32, parentIndex int32, xml string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Network.Update", 0, command, section, parentIndex, xml, flags).Store()
	return
}

// GetActive See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkIsActive
func (m *Network) GetActive() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Network", "Active").Store(&v)
	return
}

// SetAutostart See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetAutostart and https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkSetAutostart
func (m *Network) SetAutostart(v bool) (err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Set", 0, "org.libvirt.Network", "Autostart", dbus.MakeVariant(v)).Store()
	return
}

// GetAutostart See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetAutostart and https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkSetAutostart
func (m *Network) GetAutostart() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Network", "Autostart").Store(&v)
	return
}

// GetName See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetName
func (m *Network) GetName() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Network", "Name").Store(&v)
	return
}

// GetPersistent See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkIsPersistent
func (m *Network) GetPersistent() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Network", "Persistent").Store(&v)
	return
}

// GetUUID See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkGetUUIDString
func (m *Network) GetUUID() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Network", "UUID").Store(&v)
	return
}
