package libvirt

import (
	"sync"

	"github.com/godbus/dbus"
)

type Domain struct {
	conn   *Conn
	object dbus.BusObject
	path   dbus.ObjectPath

	sigs  map[<-chan *dbus.Signal]struct{}
	sigmu sync.Mutex

	//Active bool
	//Autostart bool
	//Id uint32
	//Name string
	//OSType string
	//Persistent bool
	//SchedulerType interface{}
	//Updated bool
	//UUID string
}

// NewDomain() TODO
func NewDomain(c *Conn, path dbus.ObjectPath) *Domain {
	m := &Domain{conn: c}
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

// SubscribeAgentEvent See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventAgentLifecycleCallback
func (m *Domain) SubscribeAgentEvent(callback func(state int32, reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='AgentEvent'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.AgentEvent" || 2 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32), v.Body[1].(int32))
		}
	}()
	return ch
}

// UnSubscribeAgentEvent See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventAgentLifecycleCallback
func (m *Domain) UnSubscribeAgentEvent(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='AgentEvent'")
}

// SubscribeBalloonChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventBalloonChangeCallback
func (m *Domain) SubscribeBalloonChange(callback func(actual uint64)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='BalloonChange'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.BalloonChange" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(uint64))
		}
	}()
	return ch
}

// UnSubscribeBalloonChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventBalloonChangeCallback
func (m *Domain) UnSubscribeBalloonChange(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='BalloonChange'")
}

// SubscribeBlockJob See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventBlockJobCallback Callback was registered using VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2
func (m *Domain) SubscribeBlockJob(callback func(disk string, otype int32, status int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='BlockJob'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.BlockJob" || 3 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string), v.Body[1].(int32), v.Body[2].(int32))
		}
	}()
	return ch
}

// UnSubscribeBlockJob See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventBlockJobCallback Callback was registered using VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2
func (m *Domain) UnSubscribeBlockJob(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='BlockJob'")
}

// SubscribeControlError See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGenericCallback
func (m *Domain) SubscribeControlError(callback func()) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='ControlError'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.ControlError" || 0 != len(v.Body) {
				continue
			}
			callback()
		}
	}()
	return ch
}

// UnSubscribeControlError See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGenericCallback
func (m *Domain) UnSubscribeControlError(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='ControlError'")
}

// SubscribeDeviceAdded See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceAddedCallback
func (m *Domain) SubscribeDeviceAdded(callback func(device string)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceAdded'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.DeviceAdded" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string))
		}
	}()
	return ch
}

// UnSubscribeDeviceAdded See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceAddedCallback
func (m *Domain) UnSubscribeDeviceAdded(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceAdded'")
}

// SubscribeDeviceRemovalFailed See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceRemovalFailedCallback
func (m *Domain) SubscribeDeviceRemovalFailed(callback func(device string)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceRemovalFailed'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.DeviceRemovalFailed" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string))
		}
	}()
	return ch
}

// UnSubscribeDeviceRemovalFailed See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceRemovalFailedCallback
func (m *Domain) UnSubscribeDeviceRemovalFailed(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceRemovalFailed'")
}

// SubscribeDeviceRemoved See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceRemovedCallback
func (m *Domain) SubscribeDeviceRemoved(callback func(device string)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceRemoved'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.DeviceRemoved" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string))
		}
	}()
	return ch
}

// UnSubscribeDeviceRemoved See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDeviceRemovedCallback
func (m *Domain) UnSubscribeDeviceRemoved(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DeviceRemoved'")
}

// SubscribeDiskChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDiskChangeCallback
func (m *Domain) SubscribeDiskChange(callback func(oldSrcPath string, newSrcPath string, device string, reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DiskChange'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.DiskChange" || 4 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string), v.Body[1].(string), v.Body[2].(string), v.Body[3].(int32))
		}
	}()
	return ch
}

// UnSubscribeDiskChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventDiskChangeCallback
func (m *Domain) UnSubscribeDiskChange(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='DiskChange'")
}

// SubscribeGraphics See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGraphicsCallback
func (m *Domain) SubscribeGraphics(callback func(phase int32, local interface{}, remote interface{}, authScheme string, identities []interface{})) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Graphics'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.Graphics" || 5 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32), v.Body[1].(interface{}), v.Body[2].(interface{}), v.Body[3].(string), v.Body[4].([]interface{}))
		}
	}()
	return ch
}

// UnSubscribeGraphics See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGraphicsCallback
func (m *Domain) UnSubscribeGraphics(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Graphics'")
}

// SubscribeIOError See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventIOErrorReasonCallback
func (m *Domain) SubscribeIOError(callback func(srcPath string, device string, action int32, reason string)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='IOError'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.IOError" || 4 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string), v.Body[1].(string), v.Body[2].(int32), v.Body[3].(string))
		}
	}()
	return ch
}

// UnSubscribeIOError See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventIOErrorReasonCallback
func (m *Domain) UnSubscribeIOError(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='IOError'")
}

// SubscribeJobCompleted See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventJobCompletedCallback
func (m *Domain) SubscribeJobCompleted(callback func(params map[string]interface{})) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='JobCompleted'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.JobCompleted" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(map[string]interface{}))
		}
	}()
	return ch
}

// UnSubscribeJobCompleted See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventJobCompletedCallback
func (m *Domain) UnSubscribeJobCompleted(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='JobCompleted'")
}

// SubscribeMetadataChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventMetadataChangeCallback
func (m *Domain) SubscribeMetadataChange(callback func(otype int32, nsuri string)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='MetadataChange'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.MetadataChange" || 2 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32), v.Body[1].(string))
		}
	}()
	return ch
}

// UnSubscribeMetadataChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventMetadataChangeCallback
func (m *Domain) UnSubscribeMetadataChange(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='MetadataChange'")
}

// SubscribeMigrationIteration See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventMigrationIterationCallback
func (m *Domain) SubscribeMigrationIteration(callback func(iteration int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='MigrationIteration'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.MigrationIteration" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32))
		}
	}()
	return ch
}

// UnSubscribeMigrationIteration See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventMigrationIterationCallback
func (m *Domain) UnSubscribeMigrationIteration(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='MigrationIteration'")
}

// SubscribePMSuspend See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMSuspendCallback
func (m *Domain) SubscribePMSuspend(callback func(reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMSuspend'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.PMSuspend" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32))
		}
	}()
	return ch
}

// UnSubscribePMSuspend See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMSuspendCallback
func (m *Domain) UnSubscribePMSuspend(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMSuspend'")
}

// SubscribePMSuspendDisk See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMSuspendDiskCallback
func (m *Domain) SubscribePMSuspendDisk(callback func(reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMSuspendDisk'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.PMSuspendDisk" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32))
		}
	}()
	return ch
}

// UnSubscribePMSuspendDisk See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMSuspendDiskCallback
func (m *Domain) UnSubscribePMSuspendDisk(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMSuspendDisk'")
}

// SubscribePMWakeup See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMWakeupCallback
func (m *Domain) SubscribePMWakeup(callback func(reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMWakeup'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.PMWakeup" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32))
		}
	}()
	return ch
}

// UnSubscribePMWakeup See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventPMWakeupCallback
func (m *Domain) UnSubscribePMWakeup(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='PMWakeup'")
}

// SubscribeReboot See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGenericCallback
func (m *Domain) SubscribeReboot(callback func()) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Reboot'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.Reboot" || 0 != len(v.Body) {
				continue
			}
			callback()
		}
	}()
	return ch
}

// UnSubscribeReboot See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventGenericCallback
func (m *Domain) UnSubscribeReboot(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Reboot'")
}

// SubscribeRTCChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventRTCChangeCallback
func (m *Domain) SubscribeRTCChange(callback func(utcoffset int64)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='RTCChange'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.RTCChange" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int64))
		}
	}()
	return ch
}

// UnSubscribeRTCChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventRTCChangeCallback
func (m *Domain) UnSubscribeRTCChange(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='RTCChange'")
}

// SubscribeTrayChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventTrayChangeCallback
func (m *Domain) SubscribeTrayChange(callback func(device string, reason int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='TrayChange'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.TrayChange" || 2 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(string), v.Body[1].(int32))
		}
	}()
	return ch
}

// UnSubscribeTrayChange See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventTrayChangeCallback
func (m *Domain) UnSubscribeTrayChange(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='TrayChange'")
}

// SubscribeTunable See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventTunableCallback
func (m *Domain) SubscribeTunable(callback func(params map[string]interface{})) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Tunable'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.Tunable" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(map[string]interface{}))
		}
	}()
	return ch
}

// UnSubscribeTunable See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventTunableCallback
func (m *Domain) UnSubscribeTunable(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Tunable'")
}

// SubscribeWatchdog See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventWatchdogCallback
func (m *Domain) SubscribeWatchdog(callback func(action int32)) <-chan *dbus.Signal {
	if callback == nil {
		return nil
	}
	m.sigmu.Lock()
	ch := make(chan *dbus.Signal)
	m.sigs[ch] = struct{}{}
	m.conn.conn.Signal(ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Watchdog'")
	go func() {
		for v := range ch {
			if v.Path != m.path || v.Name != "org.libvirt.Domain.Watchdog" || 1 != len(v.Body) {
				continue
			}
			callback(v.Body[0].(int32))
		}
	}()
	return ch
}

// UnSubscribeWatchdog See https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectDomainEventWatchdogCallback
func (m *Domain) UnSubscribeWatchdog(ch <-chan *dbus.Signal) {
	m.sigmu.Lock()
	delete(m.sigs, ch)
	m.sigmu.Unlock()
	m.conn.conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, "type='signal',interface='org.libvirt.Domain',member='Watchdog'")
}

// AbortJob See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainAbortJob
func (m *Domain) AbortJob() (err error) {
	err = m.object.Call("org.libvirt.Domain.AbortJob", 0).Store()
	return
}

// AddIOThread See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainAddIOThread
func (m *Domain) AddIOThread(iothreadId uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.AddIOThread", 0, iothreadId, flags).Store()
	return
}

// AttachDevice See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainAttachDeviceFlags
func (m *Domain) AttachDevice(xml string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.AttachDevice", 0, xml, flags).Store()
	return
}

// BlockCommit See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockCommit
func (m *Domain) BlockCommit(disk string, base string, top string, bandwidth uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockCommit", 0, disk, base, top, bandwidth, flags).Store()
	return
}

// BlockCopy See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockCopy
func (m *Domain) BlockCopy(disk string, destxml string, params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockCopy", 0, disk, destxml, params, flags).Store()
	return
}

// BlockJobAbort See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockJobAbort
func (m *Domain) BlockJobAbort(disk string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockJobAbort", 0, disk, flags).Store()
	return
}

// BlockPeek See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockPeek
func (m *Domain) BlockPeek(disk string, offset uint64, size uint64, flags uint32) (buffer []byte, err error) {
	err = m.object.Call("org.libvirt.Domain.BlockPeek", 0, disk, offset, size, flags).Store(&buffer)
	return
}

// BlockPull See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockPull
func (m *Domain) BlockPull(disk string, bandwidth uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockPull", 0, disk, bandwidth, flags).Store()
	return
}

// BlockRebase See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockRebase Empty string can be used to pass a NULL as @base argument.
func (m *Domain) BlockRebase(disk string, base string, bandwidth uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockRebase", 0, disk, base, bandwidth, flags).Store()
	return
}

// BlockResize See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockResize
func (m *Domain) BlockResize(disk string, size uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockResize", 0, disk, size, flags).Store()
	return
}

// BlockJobSetSpeed See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainBlockJobSetSpeed
func (m *Domain) BlockJobSetSpeed(disk string, bandwidth uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.BlockJobSetSpeed", 0, disk, bandwidth, flags).Store()
	return
}

// CoreDump See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCoreDumpWithFormat
func (m *Domain) CoreDump(to string, dumpformat uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.CoreDump", 0, to, dumpformat, flags).Store()
	return
}

// Create See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCreateWithFlags
func (m *Domain) Create(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Create", 0, flags).Store()
	return
}

// CreateWithFiles See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCreateWithFiles
func (m *Domain) CreateWithFiles(files []uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.CreateWithFiles", 0, files, flags).Store()
	return
}

// DelIOThread See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDelIOThread
func (m *Domain) DelIOThread(iothreadId uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.DelIOThread", 0, iothreadId, flags).Store()
	return
}

// Destroy See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDestroyFlags
func (m *Domain) Destroy(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Destroy", 0, flags).Store()
	return
}

// DetachDevice See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDetachDeviceFlags
func (m *Domain) DetachDevice(xml string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.DetachDevice", 0, xml, flags).Store()
	return
}

// FSFreeze See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainFSFreeze
func (m *Domain) FSFreeze(mountpoints []string, flags uint32) (frozenFilesystems uint32, err error) {
	err = m.object.Call("org.libvirt.Domain.FSFreeze", 0, mountpoints, flags).Store(&frozenFilesystems)
	return
}

// FSThaw See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainFSThaw
func (m *Domain) FSThaw(mountpoints []string, flags uint32) (thawedFilesystems uint32, err error) {
	err = m.object.Call("org.libvirt.Domain.FSThaw", 0, mountpoints, flags).Store(&thawedFilesystems)
	return
}

// FSTrim See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainFSTrim Empty string can be used to pass a NULL as @mountpoint argument.
func (m *Domain) FSTrim(mountpoint string, minimum uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.FSTrim", 0, mountpoint, minimum, flags).Store()
	return
}

// GetBlockIOParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetBlkioParameters
func (m *Domain) GetBlockIOParameters(flags uint32) (BlkioParameters map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetBlockIOParameters", 0, flags).Store(&BlkioParameters)
	return
}

// GetBlockIOTune See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetBlockIoTune
func (m *Domain) GetBlockIOTune(disk string, flags uint32) (blockIOTune map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetBlockIOTune", 0, disk, flags).Store(&blockIOTune)
	return
}

// GetBlockJobInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetBlockJobInfo
func (m *Domain) GetBlockJobInfo(disk string, flags uint32) (blockJobInfo interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetBlockJobInfo", 0, disk, flags).Store(&blockJobInfo)
	return
}

// GetControlInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetControlInfo
func (m *Domain) GetControlInfo(flags uint32) (controlInfo interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetControlInfo", 0, flags).Store(&controlInfo)
	return
}

// GetDiskErrors See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetDiskErrors
func (m *Domain) GetDiskErrors(flags uint32) (diskErrors []interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetDiskErrors", 0, flags).Store(&diskErrors)
	return
}

// GetEmulatorPinInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetEmulatorPinInfo
func (m *Domain) GetEmulatorPinInfo(flags uint32) (cpumap []bool, err error) {
	err = m.object.Call("org.libvirt.Domain.GetEmulatorPinInfo", 0, flags).Store(&cpumap)
	return
}

// GetFSInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetFSInfo
func (m *Domain) GetFSInfo(flags uint32) (fsInfo []interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetFSInfo", 0, flags).Store(&fsInfo)
	return
}

// GetGuestVcpus See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetGuestVcpus
func (m *Domain) GetGuestVcpus(flags uint32) (vcpus map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetGuestVcpus", 0, flags).Store(&vcpus)
	return
}

// GetHostname See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetHostname
func (m *Domain) GetHostname(flags uint32) (hostname string, err error) {
	err = m.object.Call("org.libvirt.Domain.GetHostname", 0, flags).Store(&hostname)
	return
}

// GetInterfaceParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetInterfaceParameters
func (m *Domain) GetInterfaceParameters(device string, flags uint32) (interfaceParameters map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetInterfaceParameters", 0, device, flags).Store(&interfaceParameters)
	return
}

// GetIOThreadInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetIOThreadInfo
func (m *Domain) GetIOThreadInfo(flags uint32) (ioThreadInfo []interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetIOThreadInfo", 0, flags).Store(&ioThreadInfo)
	return
}

// GetJobInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetJobInfo
func (m *Domain) GetJobInfo() (jobInfo interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetJobInfo", 0).Store(&jobInfo)
	return
}

// GetJobStats See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetJobStats
func (m *Domain) GetJobStats(flags uint32) (stats interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetJobStats", 0, flags).Store(&stats)
	return
}

// GetMemoryParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetMemoryParameters
func (m *Domain) GetMemoryParameters(flags uint32) (memoryParameters map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetMemoryParameters", 0, flags).Store(&memoryParameters)
	return
}

// GetMetadata See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetMetadata Empty string can be used to pass a NULL as @uri argument.
func (m *Domain) GetMetadata(itype int32, uri string, flags uint32) (metadata string, err error) {
	err = m.object.Call("org.libvirt.Domain.GetMetadata", 0, itype, uri, flags).Store(&metadata)
	return
}

// GetNumaParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetNumaParameters
func (m *Domain) GetNumaParameters(flags uint32) (numaParameters map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetNumaParameters", 0, flags).Store(&numaParameters)
	return
}

// GetPerfEvents See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetPerfEvents
func (m *Domain) GetPerfEvents(flags uint32) (perfEvents map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetPerfEvents", 0, flags).Store(&perfEvents)
	return
}

// GetSchedulerParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetSchedulerParametersFlags
func (m *Domain) GetSchedulerParameters(flags uint32) (SchedulerParameters map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetSchedulerParameters", 0, flags).Store(&SchedulerParameters)
	return
}

// GetSecurityLabelList See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetSecurityLabelList
func (m *Domain) GetSecurityLabelList() (securityLabels []interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetSecurityLabelList", 0).Store(&securityLabels)
	return
}

// GetState See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetState
func (m *Domain) GetState(flags uint32) (state interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetState", 0, flags).Store(&state)
	return
}

// GetStats See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainListGetStats
func (m *Domain) GetStats(stats uint32, flags uint32) (records map[string]interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetStats", 0, stats, flags).Store(&records)
	return
}

// GetTime See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetTime
func (m *Domain) GetTime(flags uint32) (time interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.GetTime", 0, flags).Store(&time)
	return
}

// GetVcpuPinInfo See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetVcpuPinInfo
func (m *Domain) GetVcpuPinInfo(flags uint32) (vcpuPinInfo [][]bool, err error) {
	err = m.object.Call("org.libvirt.Domain.GetVcpuPinInfo", 0, flags).Store(&vcpuPinInfo)
	return
}

// GetVcpus See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetVcpusFlags
func (m *Domain) GetVcpus(flags uint32) (vcpus uint32, err error) {
	err = m.object.Call("org.libvirt.Domain.GetVcpus", 0, flags).Store(&vcpus)
	return
}

// GetXMLDesc See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetXMLDesc
func (m *Domain) GetXMLDesc(flags uint32) (xml string, err error) {
	err = m.object.Call("org.libvirt.Domain.GetXMLDesc", 0, flags).Store(&xml)
	return
}

// HasManagedSaveImage See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainHasManagedSaveImage
func (m *Domain) HasManagedSaveImage(flags uint32) (managedSaveImage bool, err error) {
	err = m.object.Call("org.libvirt.Domain.HasManagedSaveImage", 0, flags).Store(&managedSaveImage)
	return
}

// InjectNMI See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainInjectNMI
func (m *Domain) InjectNMI(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.InjectNMI", 0, flags).Store()
	return
}

// InterfaceAddresses See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainInterfaceAddresses
func (m *Domain) InterfaceAddresses(source uint32, flags uint32) (ifaces []interface{}, err error) {
	err = m.object.Call("org.libvirt.Domain.InterfaceAddresses", 0, source, flags).Store(&ifaces)
	return
}

// ManagedSave See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainManagedSave
func (m *Domain) ManagedSave(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.ManagedSave", 0, flags).Store()
	return
}

// ManagedSaveRemove See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainManagedSaveRemove
func (m *Domain) ManagedSaveRemove(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.ManagedSaveRemove", 0, flags).Store()
	return
}

// MemoryPeek See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMemoryPeek
func (m *Domain) MemoryPeek(offset uint64, size uint64, flags uint32) (buffer []byte, err error) {
	err = m.object.Call("org.libvirt.Domain.MemoryPeek", 0, offset, size, flags).Store(&buffer)
	return
}

// MemoryStats See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMemoryStats
func (m *Domain) MemoryStats(flags uint32) (stats map[int32]uint64, err error) {
	err = m.object.Call("org.libvirt.Domain.MemoryStats", 0, flags).Store(&stats)
	return
}

// MigrateGetCompressionCache See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateGetCompressionCache
func (m *Domain) MigrateGetCompressionCache(flags uint32) (cacheSize uint64, err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateGetCompressionCache", 0, flags).Store(&cacheSize)
	return
}

// MigrateGetMaxSpeed See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateGetMaxSpeed
func (m *Domain) MigrateGetMaxSpeed(flags uint32) (bandwidth uint64, err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateGetMaxSpeed", 0, flags).Store(&bandwidth)
	return
}

// MigrateSetCompressionCache See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateSetCompressionCache
func (m *Domain) MigrateSetCompressionCache(cacheSize uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateSetCompressionCache", 0, cacheSize, flags).Store()
	return
}

// MigrateSetMaxDowntime See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateSetMaxDowntime
func (m *Domain) MigrateSetMaxDowntime(downtime uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateSetMaxDowntime", 0, downtime, flags).Store()
	return
}

// MigrateSetMaxSpeed See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateSetMaxSpeed
func (m *Domain) MigrateSetMaxSpeed(bandwidth uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateSetMaxSpeed", 0, bandwidth, flags).Store()
	return
}

// MigrateStartPostCopy See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateStartPostCopy
func (m *Domain) MigrateStartPostCopy(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateStartPostCopy", 0, flags).Store()
	return
}

// MigrateToURI3 See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainMigrateToURI3
func (m *Domain) MigrateToURI3(dconuri string, params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.MigrateToURI3", 0, dconuri, params, flags).Store()
	return
}

// OpenGraphicsFD See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainOpenGraphicsFD
func (m *Domain) OpenGraphicsFD(idx uint32, flags uint32) (fd uint32, err error) {
	err = m.object.Call("org.libvirt.Domain.OpenGraphicsFD", 0, idx, flags).Store(&fd)
	return
}

// PinEmulator See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainPinEmulator
func (m *Domain) PinEmulator(cpumap []bool, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.PinEmulator", 0, cpumap, flags).Store()
	return
}

// PinIOThread See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainPinIOThread
func (m *Domain) PinIOThread(iothreadId uint32, cpumap []bool, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.PinIOThread", 0, iothreadId, cpumap, flags).Store()
	return
}

// PinVcpu See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainPinVcpuFlags
func (m *Domain) PinVcpu(vcpu uint32, cpumap []bool, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.PinVcpu", 0, vcpu, cpumap, flags).Store()
	return
}

// PMWakeup See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainPMWakeup
func (m *Domain) PMWakeup(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.PMWakeup", 0, flags).Store()
	return
}

// Reboot See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainReboot
func (m *Domain) Reboot(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Reboot", 0, flags).Store()
	return
}

// Rename See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainRename
func (m *Domain) Rename(name string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Rename", 0, name, flags).Store()
	return
}

// Reset See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainReset
func (m *Domain) Reset(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Reset", 0, flags).Store()
	return
}

// Resume See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainResume
func (m *Domain) Resume() (err error) {
	err = m.object.Call("org.libvirt.Domain.Resume", 0).Store()
	return
}

// Save See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSaveFlags Empty string can be used to pass a NULL as @xml argument.
func (m *Domain) Save(to string, xml string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Save", 0, to, xml, flags).Store()
	return
}

// SendKey See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSendKey
func (m *Domain) SendKey(codeset uint32, holdtime uint32, keycodes []uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SendKey", 0, codeset, holdtime, keycodes, flags).Store()
	return
}

// SendProcessSignal See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSendProcessSignal
func (m *Domain) SendProcessSignal(pidValue int64, sigNum uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SendProcessSignal", 0, pidValue, sigNum, flags).Store()
	return
}

// SetBlockIOParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetBlkioParameters
func (m *Domain) SetBlockIOParameters(params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetBlockIOParameters", 0, params, flags).Store()
	return
}

// SetBlockIOTune See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetBlockIoTune
func (m *Domain) SetBlockIOTune(disk string, params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetBlockIOTune", 0, disk, params, flags).Store()
	return
}

// SetGuestVcpus See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetGuestVcpus
func (m *Domain) SetGuestVcpus(vcpumap []bool, state int32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetGuestVcpus", 0, vcpumap, state, flags).Store()
	return
}

// SetInterfaceParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetInterfaceParameters
func (m *Domain) SetInterfaceParameters(device string, params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetInterfaceParameters", 0, device, params, flags).Store()
	return
}

// SetMemory See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetMemoryFlags
func (m *Domain) SetMemory(memory uint64, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetMemory", 0, memory, flags).Store()
	return
}

// SetMemoryParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetMemoryParameters
func (m *Domain) SetMemoryParameters(params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetMemoryParameters", 0, params, flags).Store()
	return
}

// SetMemoryStatsPeriod See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetMemoryStatsPeriod
func (m *Domain) SetMemoryStatsPeriod(period int32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetMemoryStatsPeriod", 0, period, flags).Store()
	return
}

// SetMetadata See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetMetadata Empty string can be used to pass a NULL as @key or @uri argument.
func (m *Domain) SetMetadata(itype int32, metadata string, key string, uri string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetMetadata", 0, itype, metadata, key, uri, flags).Store()
	return
}

// SetNumaParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetNumaParameters
func (m *Domain) SetNumaParameters(params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetNumaParameters", 0, params, flags).Store()
	return
}

// SetPerfEvents See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetPerfEvents
func (m *Domain) SetPerfEvents(params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetPerfEvents", 0, params, flags).Store()
	return
}

// SetSchedulerParameters See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetSchedulerParametersFlags
func (m *Domain) SetSchedulerParameters(params map[string]interface{}, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetSchedulerParameters", 0, params, flags).Store()
	return
}

// SetUserPassword See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetUserPassword
func (m *Domain) SetUserPassword(user string, password string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetUserPassword", 0, user, password, flags).Store()
	return
}

// SetTime See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetTime
func (m *Domain) SetTime(seconds uint64, nseconds uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetTime", 0, seconds, nseconds, flags).Store()
	return
}

// SetVcpus See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetVcpusFlags
func (m *Domain) SetVcpus(vcpus uint32, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.SetVcpus", 0, vcpus, flags).Store()
	return
}

// Shutdown See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainShutdownFlags
func (m *Domain) Shutdown(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Shutdown", 0, flags).Store()
	return
}

// Suspend See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSuspend
func (m *Domain) Suspend() (err error) {
	err = m.object.Call("org.libvirt.Domain.Suspend", 0).Store()
	return
}

// Undefine See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainUndefineFlags
func (m *Domain) Undefine(flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.Undefine", 0, flags).Store()
	return
}

// UpdateDevice See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainUpdateDeviceFlags
func (m *Domain) UpdateDevice(xml string, flags uint32) (err error) {
	err = m.object.Call("org.libvirt.Domain.UpdateDevice", 0, xml, flags).Store()
	return
}

// GetActive See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainIsActive
func (m *Domain) GetActive() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Active").Store(&v)
	return
}

// SetAutostart See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetAutostart and https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetAutostart
func (m *Domain) SetAutostart(v bool) (err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Set", 0, "org.libvirt.Domain", "Autostart", dbus.MakeVariant(v)).Store()
	return
}

// GetAutostart See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetAutostart and https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainSetAutostart
func (m *Domain) GetAutostart() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Autostart").Store(&v)
	return
}

// GetId See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetID
func (m *Domain) GetId() (v uint32, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Id").Store(&v)
	return
}

// GetName See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetName
func (m *Domain) GetName() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Name").Store(&v)
	return
}

// GetOSType See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetOSType
func (m *Domain) GetOSType() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "OSType").Store(&v)
	return
}

// GetPersistent See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainIsPersistent
func (m *Domain) GetPersistent() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Persistent").Store(&v)
	return
}

// GetSchedulerType See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetSchedulerType
func (m *Domain) GetSchedulerType() (v interface{}, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "SchedulerType").Store(&v)
	return
}

// GetUpdated See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainIsUpdated
func (m *Domain) GetUpdated() (v bool, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "Updated").Store(&v)
	return
}

// GetUUID See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetUUIDString
func (m *Domain) GetUUID() (v string, err error) {
	err = m.object.Call("org.freedesktop.DBus.Properties.Get", 0, "org.libvirt.Domain", "UUID").Store(&v)
	return
}
