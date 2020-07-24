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

type Resolved struct {
	conn   *dbus.Conn
	sysObj dbus.BusObject
	io.Closer
}

func (r *Resolved) Close() error {
	return r.conn.Close()
}

func (r *Resolved) Introspect() (*introspect.Node, error) {
	node, err := introspect.Call(r.conn.Object("org.freedesktop.resolve1", "/org/freedesktop/resolve1"))
	if err != nil {
		return nil, err
	}
	return node, nil
}

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

func (r *Resolved) SetLinkDNS(link int32, dnsEntries []*DnsEntry) error {
	c := r.sysObj.Call(managerPath+".SetLinkDNS", 0, link, dnsEntries)
	return c.Err
}

func (r *Resolved) ResetStatistics() {

}

func (r *Resolved) SetLinkDomains() {

}

func (r *Resolved) SetLinkLLMNR() {

}

func (r *Resolved) SetLinkMulticastDNS() {

}

func (r *Resolved) SetLinkDNSSEC() {

}

func (r *Resolved) SetLinkDNSSECNegativeTrustAnchors() {

}

func (r *Resolved) RevertLink() {

}
