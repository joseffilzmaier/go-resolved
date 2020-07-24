package resolved

import (
	"fmt"
	"reflect"

	"github.com/godbus/dbus/v5"
)

var (
	linkPath = "org.freedesktop.resolve1.Link"
)

type Link struct {
	connObj dbus.BusObject
}

func (l *Link) PropDNS() ([]DnsEntry, error) {
	v, err := l.connObj.GetProperty(linkPath + ".DNS")
	if err != nil {
		return nil, err
	}
	in := v.Value().([][]interface{})
	retval := make([]DnsEntry, len(in))
	for i, e := range in {
		retval[i] = DnsEntry{
			Af:   e[0].(int32),
			Data: e[1].([]byte),
		}
	}
	return retval, nil
}

func (l *Link) PropScopesMask() (uint64, error) {
	v, err := l.connObj.GetProperty(linkPath + ".ScopesMask")
	if err != nil {
		return 0, err
	}
	return v.Value().(uint64), nil
}

func (l *Link) PropDomains() ([]Domain, error) {
	v, err := l.connObj.GetProperty(linkPath + ".Domains")
	if err != nil {
		return nil, err
	}
	in := v.Value().([][]interface{})
	retval := make([]Domain, len(in))
	for i, e := range in {
		retval[i] = Domain{
			Domain:          e[0].(string),
			SearchOrRouting: e[1].(bool),
		}
	}
	fmt.Println(reflect.TypeOf(v.Value()))
	return retval, nil
}
