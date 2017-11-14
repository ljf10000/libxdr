package libxdr

import (
	. "asdf"
	"net"
)

const (
	DnsIp4 = 0
)

const SizeofXdrDns = 8*SizeofByte +
	2*SizeofInt32 +
	1*SizeofXdrArray +
	1*SizeofXdrString

type XdrDns struct {
	V            byte
	IpVersion    byte // 0: ipv4
	IpCount      byte
	ResponseCode byte

	CountRequest        byte
	CountResponseRecord byte
	CountResponseAuth   byte
	CountResponseExtra  byte

	Delay uint32
	/*
	 * if 1==ip_count, 0==ip_version
	 *   then
	 *       ip4 is the ip address
	 *       the ip array is not used
	 */
	Ip4 XdrIp4Addr

	Ip     XdrArray // ip4
	Domain XdrString
}

func (me *XdrDns) Size() int {
	return SizeofXdrDns
}

func (me *XdrReader) dumpDns(xdr *Xdr, obj *XdrDns, tab int) {
	dump(TabN(tab) + "dns:")

	tab += 1
	dump(TabN(tab)+"ip-version:%d", obj.IpVersion)
	dump(TabN(tab)+"ip-count:%d", obj.IpCount)
	dump(TabN(tab)+"response-code:%d", obj.ResponseCode)
	dump(TabN(tab)+"count-request:%d", obj.CountRequest)
	dump(TabN(tab)+"count-response-record:%d", obj.CountResponseRecord)
	dump(TabN(tab)+"count-response-auth:%d", obj.CountResponseAuth)
	dump(TabN(tab)+"count-response-extra:%d", obj.CountResponseExtra)
	dump(TabN(tab)+"delay:%d", obj.Delay)
	dump(TabN(tab)+"ip:%d", obj.Ip4)
	dump(TabN(tab)+"domain:%s", string(me.DnsDomain(xdr, obj)))

	for i := 0; i < int(obj.IpCount); i++ {
		if obj.IsIp4() {
			// host sort
			ip4 := me.DnsIp4(xdr, obj, i)

			ip := net.IPAddr{
				IP: net.IP{
					byte(ip4 >> 24 & 0xff),
					byte(ip4 >> 16 & 0xff),
					byte(ip4 >> 8 & 0xff),
					byte(ip4 >> 0 & 0xff),
				},
			}

			dump(TabN(tab)+"ip%d:%s", i, ip.String())
		} else {
			dump(TabN(tab)+"ip%d:ipv6address", i)
		}
	}
}

func (me *XdrDns) IsIp4() bool {
	return DnsIp4 == me.IpVersion
}

func (me *XdrDns) IsIp6() bool {
	return DnsIp4 != me.IpVersion
}

func (me *XdrDns) HaveIp4Addrs() bool {
	return me.IsIp4() && me.IpCount > 1
}

func (me *XdrDns) HaveIp6Addrs() bool {
	return me.IsIp6() && me.IpCount > 0
}

func (me *XdrReader) Dns(xdr *Xdr) *XdrDns {
	return (*XdrDns)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) DnsDomain(xdr *Xdr, obj *XdrDns) []byte {
	return me.xdrString(xdr, obj.Domain)
}

func (me *XdrReader) DnsIp4(xdr *Xdr, obj *XdrDns, idx int) XdrIp4Addr {
	entry := me.xdrArrayEntry(xdr, &obj.Ip, idx)
	if nil != entry {
		return *(*XdrIp4Addr)(entry)
	} else {
		return 0
	}
}

func (me *XdrReader) DnsIp6(xdr *Xdr, obj *XdrDns, idx int) XdrIp6Addr {
	entry := me.xdrArrayEntry(xdr, &obj.Ip, idx)
	if nil != entry {
		return XdrIp6Addr(ObjToSlice(entry, int(obj.Ip.Len)))
	} else {
		return nil
	}
}
