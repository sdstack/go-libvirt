package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type Secret struct {
	conn   *dbus.Conn
	object dbus.BusObject

	UUID      string
	UsageID   string
	UsageType int32
}

// NewSecret() establishes a connection to the system bus and authenticates.
func NewSecret() (*Secret, error) {
	m := new(Secret)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Secret) initConnection() error {
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

// GetValue See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretGetValue
func (m *Secret) GetValue(flags uint32) (value []byte, err error) {
	err = m.object.Call("org.libvirt.Secret.GetValue", 0, flags).Store(&value)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretGetXMLDesc
func (m *Secret) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.Secret.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// SetValue See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretSetValue
func (m *Secret) SetValue(value []byte, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Secret.SetValue", 0, value, flags).Store()
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretUndefine
func (m *Secret) Undefine() (err error) {
	err = m.object.Call("org.libvirt.Secret.Undefine", 0).Store()
	return
}
