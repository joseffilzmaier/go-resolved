package resolved

import (
	"net"
	"syscall"
)

func NewDnsEntry(ip net.IP) *DnsEntry {
	if ip.To4() != nil {
		return &DnsEntry{
			Af:   syscall.AF_INET,
			Data: ip.To4(),
		}
	}
	return &DnsEntry{
		Af:   syscall.AF_INET6,
		Data: ip.To16(),
	}
}

type DnsEntry struct {
	Af   int32
	Data []byte
}

type Domain struct {
	Domain          string
	SearchOrRouting bool
}
