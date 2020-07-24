package resolved

import (
	"fmt"
	"net"
	"testing"

	"github.com/godbus/dbus/v5"
)

func TestGetLink(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Error(err)
	}
	defer r.Close()
	obj, err := r.GetLink(2)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}

func TestSetLinkDNS(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Error(err)
	}
	defer r.Close()
	entries := []*DnsEntry{
		NewDnsEntry(net.IPv4(10, 0, 0, 138)),
		NewDnsEntry(net.IPv4(8, 8, 8, 8)),
		NewDnsEntry(net.IPv4(8, 8, 4, 4)),
	}
	fmt.Println(dbus.SignatureOf(entries))
	err = r.SetLinkDNS(2, entries)
	if err != nil {
		t.Error(err)
	}
	print("Success")
}

func TestInvokeOnLink(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Error(err)
	}
	defer r.Close()
	l, err := r.GetLink(2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l.connObj)
	dns, err := l.PropDNS()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dns)
	m, err := l.PropScopesMask()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
	d, err := l.PropDomains()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d)
}
