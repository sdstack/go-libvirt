package libvirt

import (
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type StoragePool struct {
	conn   *dbus.Conn
	object dbus.BusObject

	Active     uint
	Autostart  uint
	Name       string
	Persistent uint
	UUID       string
}

// NewStoragePool() establishes a connection to the system bus and authenticates.
func NewStoragePool() (*StoragePool, error) {
	m := new(StoragePool)

	if err := m.initConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *StoragePool) initConnection() error {
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

// Build See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolBuild
func (m *StoragePool) Build(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Build", 0, flags).Store()
	return
}

// Create See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolCreate
func (m *StoragePool) Create(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Create", 0, flags).Store()
	return
}

// Delete See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDelete
func (m *StoragePool) Delete(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Delete", 0, flags).Store()
	return
}

// Destroy See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDestroy
func (m *StoragePool) Destroy() (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Destroy", 0).Store()
	return
}

// GetInfo See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolGetInfo
func (m *StoragePool) GetInfo() (info interface{}, err error) {
	err = m.object.Call("org.libvirt.StoragePool.GetInfo", 0).Store(&info)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolGetXMLDesc
func (m *StoragePool) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.StoragePool.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// ListStorageVolumes See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolListAllVolumes
func (m *StoragePool) ListStorageVolumes(flags uint32) (storageVols []dbus.ObjectPath, err error) {
	err = m.object.Call("org.libvirt.StoragePool.ListStorageVolumes", 0, flags).Store(&storageVols)
	return
}

// Refresh See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolRefresh
func (m *StoragePool) Refresh(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Refresh", 0, flags).Store()
	return
}

// StorageVolCreateXML See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolCreateXML
func (m *StoragePool) StorageVolCreateXML(xml string, flags uint32) (storageVol dbus.ObjectPath, err error) {
	err = m.object.Call("org.libvirt.StoragePool.StorageVolCreateXML", 0, xml, flags).Store(&storageVol)
	return
}

// StorageVolCreateXMLFrom See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolCreateXMLFrom Call with @key argument set to the key of the storage volume to be cloned.
func (m *StoragePool) StorageVolCreateXMLFrom(xml string, key string, flags uint32) (storageVol dbus.ObjectPath, err error) {
	err = m.object.Call("org.libvirt.StoragePool.StorageVolCreateXMLFrom", 0, xml, key, flags).Store(&storageVol)
	return
}

// StorageVolLookupByName See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolLookupByName
func (m *StoragePool) StorageVolLookupByName(name string) (storageVol dbus.ObjectPath, err error) {
	err = m.object.Call("org.libvirt.StoragePool.StorageVolLookupByName", 0, name).Store(&storageVol)
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolUndefine
func (m *StoragePool) Undefine() (err error) {
	err = m.object.Call("org.libvirt.StoragePool.Undefine", 0).Store()
	return
}
