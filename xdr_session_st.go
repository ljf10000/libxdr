package libxdr

import (
	. "asdf"
)

const SizeofXdrSessionSt = 5*SizeofInt32*2 + SizeofInt16*2

type XdrSessionSt struct {
	Flow          [2]uint32
	IpPacket      [2]uint32
	TcpDisorder   [2]uint32
	TcpRetransmit [2]uint32
	IpFrag        [2]uint32
	Duration      [2]uint16
}

func (me *XdrSessionSt) Size() int {
	return SizeofXdrSessionSt
}

func (me *XdrReader) dumpSessionSt(xdr *Xdr, obj *XdrSessionSt, tab int) {
	dump(TabN(tab) + "session-st:")

	tab += 1
	for i := 0; i < 2; i++ {
		prefix := "up"
		if 1 == i {
			prefix = "down"
		}

		dump(TabN(tab)+"%s-flow:%d", prefix, obj.Flow[i])
		dump(TabN(tab)+"%s-ip-packet:%d", prefix, obj.IpPacket[i])
		dump(TabN(tab)+"%s-tcp-disorder:%d", prefix, obj.TcpDisorder[i])
		dump(TabN(tab)+"%s-tcp-retransmit:%d", prefix, obj.TcpRetransmit[i])
		dump(TabN(tab)+"%s-ip-frag:%d", prefix, obj.IpFrag[i])
		dump(TabN(tab)+"%s-duration:%d", prefix, obj.Duration[i])
	}
}

func (me *XdrReader) SessionSt(xdr *Xdr) *XdrSessionSt {
	return (*XdrSessionSt)(me.xdrMember(xdr, xdr.OffsetofSessionSt))
}

type XdrServiceSt = XdrSessionSt

func (me *XdrReader) ServiceSt(xdr *Xdr) *XdrServiceSt {
	return (*XdrServiceSt)(me.xdrMember(xdr, xdr.OffsetofServiceSt))
}

func (me *XdrReader) dumpServiceSt(xdr *Xdr, obj *XdrServiceSt, tab int) {
	me.dumpSessionSt(xdr, obj, tab)
}
