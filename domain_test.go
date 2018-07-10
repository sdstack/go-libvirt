package libvirt

import (
	"fmt"
	"testing"
)

func TestListDomains(t *testing.T) {
	conn, err := NewConnect()
	if err != nil {
		t.Fatal(err)
	}
	domain, err := conn.DomainLookupByName("winxp")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#+v\n", domain)
}
