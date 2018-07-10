package libvirt

import (
	"sync"

	"github.com/godbus/dbus"
)

type StorageVol struct {
	conn   *Conn
	object dbus.BusObject
	path   dbus.ObjectPath

	sigs  map[<-chan *dbus.Signal]struct{}
	sigmu sync.Mutex

	Name string
	Key  string
	Path string
}

// NewStorageVol() TODO
func NewStorageVol(c *Conn, path dbus.ObjectPath) *StorageVol {
	m := &StorageVol{conn: c}
	if path != "" {
		m.object = c.conn.Object("org.libvirt", path)
	} else {
		m.object = c.object
	}
	m.path = c.object.Path()

	m.sigs = make(map[<-chan *dbus.Signal]struct{})

	return m
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
