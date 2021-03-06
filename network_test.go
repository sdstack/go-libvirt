package libvirt

import "testing"

func TestListNetworks(t *testing.T) {
	c, err := NewConn(DriverQEMU)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	conn := NewConnect(c, "")
	npath, err := conn.NetworkLookupByName("default")
	if err != nil {
		t.Fatal(err)
	}
	network := NewNetwork(c, npath)
	_, err = network.GetXMLDesc(0)
	if err != nil {
		t.Fatal(err)
	}

}
