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
	conn := NewConnect(c, "")
	domain, err := conn.DomainLookupByName("winxp")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#+v\n", domain)
}
