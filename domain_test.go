package libvirt

import (
	"fmt"
	"testing"
)

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
	xml, err := domain.GetXMLDesc(0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", xml)

}
