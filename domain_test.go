package libvirt

import (
	"fmt"
	"testing"

	"github.com/godbus/dbus"
)

func DomainEventCb(domain dbus.ObjectPath, event int32, detail int32) {
	fmt.Printf("%s %d %d\n", domain, event, detail)
}

func TestListDomains(t *testing.T) {
	c, err := NewConn(DriverQEMU)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	conn := NewConnect(c, "")
	dpath, err := conn.DomainLookupByName("winxp")
	if err != nil {
		t.Fatal(err)
	}
	domain := NewDomain(c, dpath)
	_, err = domain.GetXMLDesc(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.GetAllDomainStats(0, 1)
	if err != nil {
		panic(err)
	}

	ch := conn.SubscribeDomainEvent(DomainEventCb)
	defer conn.UnSubscribeDomainEvent(ch)

	st, err := domain.GetAutostart()
	if err != nil {
		panic(err)
	}
	fmt.Printf("autostart %#+v\n", st)
	select {}
}
