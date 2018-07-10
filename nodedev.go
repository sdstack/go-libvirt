package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type NodeDevice struct {
	conn   *dbus.Conn
	object dbus.BusObject

	Name   string
	Parent string
}

// NewNodeDevice() establishes a connection to the system bus and authenticates.
func NewNodeDevice() (*NodeDevice, error) {
	m := new(NodeDevice)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *NodeDevice) initConnection() error {
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

	m.object = m.conn.Object("org.libvirt.NodeDevice", dbus.ObjectPath("/org/libvirt/nodedev"))

	return nil
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
