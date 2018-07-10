package libvirt

//go:generate go run gen.go

import (
	"errors"
	"os"
	"strconv"

	"github.com/godbus/dbus"
)

type Conn struct {
	conn   *dbus.Conn
	object dbus.BusObject
}

type Driver uint8

const (
	DriverVBox Driver = iota
	DriverVZ
	DriverQEMU
	DriverOpenVZ
	DriverBHyve
	DriverLXC
	DriverTest
	DriverXen
	DriverUML
)

// NewConn() establishes a connection to the system bus and authenticates.
func NewConn(d Driver) (*Conn, error) {
	c := new(Conn)

	if err := c.initConnection(d); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) initConnection(d Driver) error {
	var err error

	drivers := make(map[Driver]string)
	drivers[DriverVBox] = "VBox"
	drivers[DriverVZ] = "VZ"
	drivers[DriverQEMU] = "QEMU"
	drivers[DriverOpenVZ] = "OpenVZ"
	drivers[DriverBHyve] = "BHyve"
	drivers[DriverLXC] = "LXC"
	drivers[DriverTest] = "Test"
	drivers[DriverXen] = "Xen"
	drivers[DriverUML] = "UML"

	c.conn, err = dbus.SystemBusPrivate()
	if err != nil {
		return err
	}

	// Only use EXTERNAL method, and hardcode the uid (not username)
	// to avoid a username lookup (which requires a dynamically linked
	// libc)
	methods := []dbus.Auth{dbus.AuthExternal(strconv.Itoa(os.Getuid()))}

	err = c.conn.Auth(methods)
	if err != nil {
		c.conn.Close()
		return err
	}

	err = c.conn.Hello()
	if err != nil {
		c.conn.Close()
		return err
	}

	driver, ok := drivers[d]
	if !ok {
		c.conn.Close()
		return errors.New("unknown driver specified")
	}

	c.object = c.conn.Object("org.libvirt", dbus.ObjectPath("/org/libvirt/"+driver))

	return nil
}
