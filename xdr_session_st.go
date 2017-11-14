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
