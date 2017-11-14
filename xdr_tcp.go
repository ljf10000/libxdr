package libxdr

import (
	. "asdf"
)

const SizeofXdrTcp = 2*SizeofInt16 +
	4*SizeofByte +
	3*SizeofInt32 +
	1*SizeofInt16 +
	6*SizeofByte

type XdrTcp struct {
	SynAckToSynTime uint16
	AckToSynTime    uint16

	Complete    byte
	CloseReason byte
	_           byte
	_           byte

	FirstRequestDelay  uint32
	FirstResponseDelay uint32
	Window             uint32

	Mss           uint16
	CountRetry    byte
	CountRetryAck byte

	CountAck      byte
	ConnectStatus byte
	HandShake12   byte
	HandShake23   byte
}

func (me *XdrTcp) Size() int {
	return SizeofXdrTcp
}

func (me *XdrReader) dumpTcp(xdr *Xdr, obj *XdrTcp, tab int) {

}

func (me *XdrReader) Tcp(xdr *Xdr) *XdrTcp {
	return (*XdrTcp)(me.xdrMember(xdr, xdr.OffsetofL4))
}
