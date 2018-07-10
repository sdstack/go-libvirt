package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type StorageVol struct {
	conn   *dbus.Conn
	object dbus.BusObject

	Name string
	Key  string
	Path string
}

// NewStorageVol() establishes a connection to the system bus and authenticates.
func NewStorageVol() (*StorageVol, error) {
	m := new(StorageVol)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *StorageVol) initConnection() error {
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

	m.object = m.conn.Object("org.libvirt.StorageVol", dbus.ObjectPath("/org/libvirt/storagevol"))

	return nil
}

// Delete See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolDelete
func (m *StorageVol) Delete(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StorageVol.Delete", 0, flags).Store()
	return
}

// GetInfo See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolGetInfoFlags
func (m *StorageVol) GetInfo(flags uint32) (info interface{}, err error) {
	err = m.object.Call("org.libvirt.StorageVol.GetInfo", 0, flags).Store(&info)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolGetXMLDesc
func (m *StorageVol) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.StorageVol.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// Resize See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolResize
func (m *StorageVol) Resize(capacity uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StorageVol.Resize", 0, capacity, flags).Store()
	return
}

// Wipe See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolWipePattern
func (m *StorageVol) Wipe(pattern uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StorageVol.Wipe", 0, pattern, flags).Store()
	return
}
