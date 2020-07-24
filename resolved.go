package resolved

import (
	"io"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

var (
	managerPath = "org.freedesktop.resolve1.Manager"
	likPath     = "org.freedesktop.resolve1.Link"
)

// New instantiates a new Resolved object allowing to manage systemd-resolved
func New() (*Resolved, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return &Resolved{
		conn:   conn,
		sysObj: conn.Object("org.freedesktop.resolve1", dbus.ObjectPath("/org/freedesktop/resolve1")),
	}, nil
}

// Resolved Central type for invoking methods or reading properties from systemd-resolved
type Resolved struct {
	conn   *dbus.Conn
	sysObj dbus.BusObject
	io.Closer
}

// Close closes the underlying dbus connection
func (r *Resolved) Close() error {
	return r.conn.Close()
}

// Introspect list available methods
func (r *Resolved) Introspect() (*introspect.Node, error) {
	node, err := introspect.Call(r.conn.Object("org.freedesktop.resolve1", "/org/freedesktop/resolve1"))
	if err != nil {
		return nil, err
	}
	return node, nil
}

// GetLink takes a network interface index and returns a Link object allowing to manipulate it
func (r *Resolved) GetLink(idx int32) (l *Link, err error) {
	var obj dbus.ObjectPath
	if err = r.sysObj.Call(managerPath+".GetLink", 0, idx).Store(&obj); err != nil {
		return nil, err
	}
	l = &Link{
		r.conn.Object("org.freedesktop.resolve1", obj),
	}
	return
}

// SetLinkDNS set dns entries for a given link
func (r *Resolved) SetLinkDNS(link int32, dnsEntries []*DnsEntry) error {
	c := r.sysObj.Call(managerPath+".SetLinkDNS", 0, link, dnsEntries)
	return c.Err
}

func (r *Resolved) ResetStatistics() {
	//TODO
}

func (r *Resolved) SetLinkDomains() {
	//TODO
}

func (r *Resolved) SetLinkLLMNR() {
	//TODO
}

func (r *Resolved) SetLinkMulticastDNS() {
	//TODO
}

func (r *Resolved) SetLinkDNSSEC() {
	//TODO
}

func (r *Resolved) SetLinkDNSSECNegativeTrustAnchors() {
	//TODO
}

func (r *Resolved) RevertLink() {
	//TODO
}
