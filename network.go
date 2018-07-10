package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type Network struct {
	conn   *dbus.Conn
	object dbus.BusObject

	Active     uint
	Autostart  uint
	Name       string
	Persistent uint
	UUID       string
}

// NewNetwork() establishes a connection to the system bus and authenticates.
func NewNetwork() (*Network, error) {
	m := new(Network)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Network) initConnection() error {
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

	m.object = m.conn.Object("org.libvirt", dbus.ObjectPath("/org/libvirt/QEMU"))

	return nil
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
