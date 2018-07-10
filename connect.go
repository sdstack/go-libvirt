package libvirt

import "github.com/godbus/dbus"

type Connect struct {
	conn *Conn
	path dbus.ObjectPath

	Encrypted  uint
	Hostname   string
	LibVersion uint64
	Secure     uint
	Version    uint64
}

// NewConnect() TODO
func NewConnect(c *Conn, path dbus.ObjectPath) *Connect {
	return &Connect{
		conn: c,
		path: path,
	}
}

// BaselineCPU See https://libvirt.org/html/libvirt-libvirt-host.html#virConnectBaselineCPU
func (m *Connect) BaselineCPU(xmlCPUs []string, flags uint32) (cpu string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.BaselineCPU", 0, xmlCPUs, flags).Store(&cpu)
	return
}

// CompareCPU See https://libvirt.org/html/libvirt-libvirt-host.html#virConnectCompareCPU
func (m *Connect) CompareCPU(xmlDesc string, flags uint32) (compareResult int32, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.CompareCPU", 0, xmlDesc, flags).Store(&compareResult)
	return
}

// DomainCreateXML See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCreateXML
func (m *Connect) DomainCreateXML(xml string, flags uint32) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainCreateXML", 0, xml, flags).Store(&domain)
	return
}

// DomainCreateXMLWithFiles See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCreateXMLWithFiles
func (m *Connect) DomainCreateXMLWithFiles(xml string, files []uint32, flags uint32) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainCreateXMLWithFiles", 0, xml, files, flags).Store(&domain)
	return
}

// DomainDefineXML See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDefineXML
func (m *Connect) DomainDefineXML(xml string) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainDefineXML", 0, xml).Store(&domain)
	return
}

// DomainLookupByID See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainLookupByID
func (m *Connect) DomainLookupByID(id int32) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainLookupByID", 0, id).Store(&domain)
	return
}

// DomainLookupByName See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainLookupByName
func (m *Connect) DomainLookupByName(name string) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainLookupByName", 0, name).Store(&domain)
	return
}

// DomainLookupByUUID See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainLookupByUUIDString
func (m *Connect) DomainLookupByUUID(uuid string) (domain dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainLookupByUUID", 0, uuid).Store(&domain)
	return
}

// DomainRestore See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainRestoreFlags Empty string can be used to pass a NULL as @xml argument.
func (m *Connect) DomainRestore(from string, xml string, flags uint32) (err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainRestore", 0, from, xml, flags).Store()
	return
}

// DomainSaveImageDefineXML See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSaveImageDefineXML
func (m *Connect) DomainSaveImageDefineXML(file string, xml string, flags uint32) (err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainSaveImageDefineXML", 0, file, xml, flags).Store()
	return
}

// DomainSaveImageGetXMLDesc See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSaveImageGetXMLDesc
func (m *Connect) DomainSaveImageGetXMLDesc(file string, flags uint32) (xml string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.DomainSaveImageGetXMLDesc", 0, file, flags).Store(&xml)
	return
}

// FindStoragePoolSources See https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectFindStoragePoolSources Empty string can be used to pass a NULL as @srcSpec argument.
func (m *Connect) FindStoragePoolSources(itype string, srcSpec string, flags uint32) (storagePoolSources string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.FindStoragePoolSources", 0, itype, srcSpec, flags).Store(&storagePoolSources)
	return
}

// GetAllDomainStats See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectGetAllDomainStats
func (m *Connect) GetAllDomainStats(stats uint32, flags uint32) (records []interface{}, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.GetAllDomainStats", 0, stats, flags).Store(&records)
	return
}

// GetCapabilities See https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetCapabilities
func (m *Connect) GetCapabilities() (capabilities string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.GetCapabilities", 0).Store(&capabilities)
	return
}

// GetCPUModelNames See https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetCPUModelNames
func (m *Connect) GetCPUModelNames(arch string, flags uint32) (models []string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.GetCPUModelNames", 0, arch, flags).Store(&models)
	return
}

// GetDomainCapabilities See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectGetDomainCapabilities Empty string can be used to pass a NULL as @emulatorbin, @arch, @machine or @virttype argument.
func (m *Connect) GetDomainCapabilities(emulatorbin string, arch string, machine string, virttype string, flags uint32) (domCapabilities string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.GetDomainCapabilities", 0, emulatorbin, arch, machine, virttype, flags).Store(&domCapabilities)
	return
}

// GetSysinfo See https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetSysinfo
func (m *Connect) GetSysinfo(flags uint32) (sysinfo string, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.GetSysinfo", 0, flags).Store(&sysinfo)
	return
}

// ListDomains See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
func (m *Connect) ListDomains(flags uint32) (domains []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListDomains", 0, flags).Store(&domains)
	return
}

// ListNetworks See https://libvirt.org/html/libvirt-libvirt-network.html#virConnectListAllNetworks
func (m *Connect) ListNetworks(flags uint32) (networks []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListNetworks", 0, flags).Store(&networks)
	return
}

// ListNodeDevices See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virConnectListAllNodeDevices
func (m *Connect) ListNodeDevices(flags uint32) (devs []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListNodeDevices", 0, flags).Store(&devs)
	return
}

// ListNWFilters See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virConnectListAllNWFilters
func (m *Connect) ListNWFilters(flags uint32) (nwfilters []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListNWFilters", 0, flags).Store(&nwfilters)
	return
}

// ListSecrets See https://libvirt.org/html/libvirt-libvirt-secret.html#virConnectListAllSecrets
func (m *Connect) ListSecrets(flags uint32) (secrets []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListSecrets", 0, flags).Store(&secrets)
	return
}

// ListStoragePools See https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectListAllStoragePools
func (m *Connect) ListStoragePools(flags uint32) (storagePools []dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.ListStoragePools", 0, flags).Store(&storagePools)
	return
}

// NetworkCreateXML See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkCreateXML
func (m *Connect) NetworkCreateXML(xml string) (network dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NetworkCreateXML", 0, xml).Store(&network)
	return
}

// NetworkDefineXML See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkDefineXML
func (m *Connect) NetworkDefineXML(xml string) (network dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NetworkDefineXML", 0, xml).Store(&network)
	return
}

// NetworkLookupByName See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkLookupByName
func (m *Connect) NetworkLookupByName(name string) (network dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NetworkLookupByName", 0, name).Store(&network)
	return
}

// NetworkLookupByUUID See https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkLookupByUUIDString
func (m *Connect) NetworkLookupByUUID(uuid string) (network dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NetworkLookupByUUID", 0, uuid).Store(&network)
	return
}

// NodeDeviceCreateXML See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceCreateXML
func (m *Connect) NodeDeviceCreateXML(xml string, flags uint32) (dev dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeDeviceCreateXML", 0, xml, flags).Store(&dev)
	return
}

// NodeDeviceLookupByName See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceLookupByName
func (m *Connect) NodeDeviceLookupByName(name string) (dev dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeDeviceLookupByName", 0, name).Store(&dev)
	return
}

// NodeDeviceLookupSCSIHostByWWN See https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceLookupSCSIHostByWWN
func (m *Connect) NodeDeviceLookupSCSIHostByWWN(wwnn string, wwpn string, flags uint32) (dev dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeDeviceLookupSCSIHostByWWN", 0, wwnn, wwpn, flags).Store(&dev)
	return
}

// NWFilterDefineXML See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterDefineXML
func (m *Connect) NWFilterDefineXML(xml string) (nwfilter dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NWFilterDefineXML", 0, xml).Store(&nwfilter)
	return
}

// NWFilterLookupByName See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterLookupByName
func (m *Connect) NWFilterLookupByName(name string) (nwfilter dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NWFilterLookupByName", 0, name).Store(&nwfilter)
	return
}

// NWFilterLookupByUUID See https://libvirt.org/html/libvirt-libvirt-nwfilter.html#virNWFilterLookupByUUIDString
func (m *Connect) NWFilterLookupByUUID(uuid string) (nwfilter dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NWFilterLookupByUUID", 0, uuid).Store(&nwfilter)
	return
}

// NodeGetCPUMap See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetCPUMap
func (m *Connect) NodeGetCPUMap(flags uint32) (res []uint, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetCPUMap", 0, flags).Store(&res)
	return
}

// NodeGetCPUStats See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetCPUStats
func (m *Connect) NodeGetCPUStats(cpuNum int32, flags uint32) (cpuStats map[string]uint64, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetCPUStats", 0, cpuNum, flags).Store(&cpuStats)
	return
}

// NodeGetFreeMemory See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetFreeMemory
func (m *Connect) NodeGetFreeMemory() (freemem uint64, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetFreeMemory", 0).Store(&freemem)
	return
}

// NodeGetMemoryParameters See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetMemoryParameters
func (m *Connect) NodeGetMemoryParameters(flags uint32) (memoryParameters map[string]interface{}, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetMemoryParameters", 0, flags).Store(&memoryParameters)
	return
}

// NodeGetMemoryStats See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetMemoryStats
func (m *Connect) NodeGetMemoryStats(cellNum int32, flags uint32) (stats map[string]uint64, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetMemoryStats", 0, cellNum, flags).Store(&stats)
	return
}

// NodeGetSecurityModel See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeGetSecurityModel
func (m *Connect) NodeGetSecurityModel() (secModel interface{}, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeGetSecurityModel", 0).Store(&secModel)
	return
}

// NodeSetMemoryParameters See https://libvirt.org/html/libvirt-libvirt-host.html#virNodeSetMemoryParameters
func (m *Connect) NodeSetMemoryParameters(params map[string]interface{}, flags uint32) (err error) {
	err = m.conn.object.Call("org.libvirt.Connect.NodeSetMemoryParameters", 0, params, flags).Store()
	return
}

// SecretDefineXML See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretDefineXML
func (m *Connect) SecretDefineXML(xml string, flags uint32) (secret dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.SecretDefineXML", 0, xml, flags).Store(&secret)
	return
}

// SecretLookupByUUID See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretLookupByUUIDString
func (m *Connect) SecretLookupByUUID(uuid string) (secret dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.SecretLookupByUUID", 0, uuid).Store(&secret)
	return
}

// SecretLookupByUsage See https://libvirt.org/html/libvirt-libvirt-secret.html#virSecretLookupByUsage
func (m *Connect) SecretLookupByUsage(usageType int32, usageID string) (secret dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.SecretLookupByUsage", 0, usageType, usageID).Store(&secret)
	return
}

// StoragePoolCreateXML See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolCreateXML
func (m *Connect) StoragePoolCreateXML(xml string, flags uint32) (storagePool dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StoragePoolCreateXML", 0, xml, flags).Store(&storagePool)
	return
}

// StoragePoolDefineXML See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDefineXML
func (m *Connect) StoragePoolDefineXML(xml string, flags uint32) (storagePool dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StoragePoolDefineXML", 0, xml, flags).Store(&storagePool)
	return
}

// StoragePoolLookupByName See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolLookupByName
func (m *Connect) StoragePoolLookupByName(name string) (storagePool dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StoragePoolLookupByName", 0, name).Store(&storagePool)
	return
}

// StoragePoolLookupByUUID See https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolLookupByUUIDString
func (m *Connect) StoragePoolLookupByUUID(uuid string) (storagePool dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StoragePoolLookupByUUID", 0, uuid).Store(&storagePool)
	return
}

// StorageVolLookupByKey See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolLookupByKey
func (m *Connect) StorageVolLookupByKey(key string) (storageVol dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StorageVolLookupByKey", 0, key).Store(&storageVol)
	return
}

// StorageVolLookupByPath See https://libvirt.org/html/libvirt-libvirt-storage.html#virStorageVolLookupByPath
func (m *Connect) StorageVolLookupByPath(path string) (storageVol dbus.ObjectPath, err error) {
	err = m.conn.object.Call("org.libvirt.Connect.StorageVolLookupByPath", 0, path).Store(&storageVol)
	return
}
