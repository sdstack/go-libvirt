package libvirt

import "github.com/godbus/dbus"

type Secret struct {
	conn *Conn
	path dbus.ObjectPath

	UUID      string
	UsageID   string
	UsageType int32
}

// NewSecret() TODO
func NewSecret(c *Conn, path dbus.ObjectPath) *Secret {
	return &Secret{
		conn: c,
		path: path,
	}
}

// GetValue See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretGetValue
func (m *Secret) GetValue(flags uint32) (value []byte, err error) {
	err = m.conn.object.Call("org.libvirt.Secret.GetValue", 0, flags).Store(&value)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretGetXMLDesc
func (m *Secret) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.conn.object.Call("org.libvirt.Secret.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// SetValue See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretSetValue
func (m *Secret) SetValue(value []byte, flags uint32) (err error) {
	err = m.conn.object.Call("org.libvirt.Secret.SetValue", 0, value, flags).Store()
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretUndefine
func (m *Secret) Undefine() (err error) {
	err = m.conn.object.Call("org.libvirt.Secret.Undefine", 0).Store()
	return
}
