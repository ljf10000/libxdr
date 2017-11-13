package libxdr

const (
	DnsIp4 = 0
)

type XdrDns struct {
	ResponseCode        byte
	CountRequest        byte
	CountResponseRecord byte
	CountResponseAuth   byte

	CountResponseExtra byte
	IpVersion          byte // 0: ipv4
	IpCount            byte
	_                  byte

	Delay uint32
	/*
	 * if 1==ip_count, 0==ip_version
	 *   then
	 *       ip4 is the ip address
	 *       the ip array is not used
	 */
	Ip4    XdrIp4Addr
	Ip     XdrArray // ip4
	Domain XdrString
}

func (me *XdrDns) IsIp4() bool {
	return DnsIp4 == me.IpVersion
}

func (me *XdrMemFile) Dns(xdr *Xdr) *XdrDns {
	return (*XdrDns)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) DnsDomain(xdr *Xdr, obj *XdrDns) []byte {
	return me.xdrString(xdr, obj.Domain)
}

// safe, needn't copy
func (me *XdrMemFile) DnsIp4(xdr *Xdr, obj *XdrDns) []XdrIp4Addr {
	if !obj.IsIp4() || obj.IpCount < 2 {
		return nil
	}

	body := me.xdrArrayBody(xdr, obj.Ip)
	if nil == body {
		return nil
	}

	count := int(obj.IpCount)
	addrs := make([]XdrIp4Addr, 0, count)
	for i := 0; i < count; i++ {
		entry := me.xdrArrayEntry(xdr, body, int(obj.Ip.Size), i)
		addrs = append(addrs, *(*XdrIp4Addr)(entry))
	}

	return addrs
}

// safe, needn't copy
func (me *XdrMemFile) DnsIp6(xdr *Xdr, obj *XdrDns) []XdrIp6Addr {
	if obj.IsIp4() || obj.IpCount < 1 {
		return nil
	}

	body := me.xdrArrayBody(xdr, obj.Ip)
	if nil == body {
		return nil
	}

	count := int(obj.IpCount)
	addrs := make([]XdrIp6Addr, 0, count)
	for i := 0; i < count; i++ {
		addr := me.xdrArrayEntrySlice(xdr, body, int(obj.Ip.Size), i)

		addrs = append(addrs, XdrIp6Addr(addr))
	}

	return addrs
}
