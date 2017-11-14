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
	dump(TabN(tab) + "tcp:")

	tab += 1
	dump(TabN(tab)+"syn-ack-to-syn-time:%d", obj.SynAckToSynTime)
	dump(TabN(tab)+"ack-to-syn-time:%d", obj.AckToSynTime)
	dump(TabN(tab)+"complete:%d", obj.Complete)
	dump(TabN(tab)+"close-reason:%d", obj.CloseReason)
	dump(TabN(tab)+"first-request-delay:%d", obj.FirstRequestDelay)
	dump(TabN(tab)+"first-response-delay:%d", obj.FirstResponseDelay)
	dump(TabN(tab)+"window:%d", obj.Window)
	dump(TabN(tab)+"mss:%d", obj.Mss)
	dump(TabN(tab)+"count-retry:%d", obj.CountRetry)
	dump(TabN(tab)+"count-retry-ack:%d", obj.CountRetryAck)
	dump(TabN(tab)+"count-ack:%d", obj.CountAck)
	dump(TabN(tab)+"connect-status:%d", obj.ConnectStatus)
	dump(TabN(tab)+"handshake12:%d", obj.HandShake12)
	dump(TabN(tab)+"handshake23:%d", obj.HandShake23)
}

func (me *XdrReader) Tcp(xdr *Xdr) *XdrTcp {
	return (*XdrTcp)(me.xdrMember(xdr, xdr.OffsetofL4))
}
