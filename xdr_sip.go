package libxdr

import (
	. "asdf"
)

const SizeofXdrSip = 4*SizeofByte + 2*SizeofInt16 + 3*SizeofXdrString

type XdrSip struct {
	CallDirection byte
	CallType      byte
	HangupReason  byte
	SignalType    byte

	DataFlowCount uint16
	Flags         uint16

	CallingNumber XdrString
	CalledNumber  XdrString
	SessionId     XdrString
}

func (me *XdrSip) Size() int {
	return SizeofXdrSip
}

func (me *XdrReader) dumpSip(xdr *Xdr, obj *XdrSip, tab int) {
	dump(TabN(tab) + "sip:")

	tab += 1
	dump(TabN(tab)+"call-direction:%d", obj.CallDirection)
	dump(TabN(tab)+"call-type:%d", obj.CallType)
	dump(TabN(tab)+"hangup-reason:%d", obj.HangupReason)
	dump(TabN(tab)+"signal-type:%d", obj.SignalType)
	dump(TabN(tab)+"data-flow-count:%d", obj.DataFlowCount)
	dump(TabN(tab)+"flags:%d", obj.Flags)

	dump(TabN(tab)+"calling-number:%s", string(me.SipCallingNumber(xdr, obj)))
	dump(TabN(tab)+"called-number:%s", string(me.SipCalledNumber(xdr, obj)))
	dump(TabN(tab)+"session-id:%s", string(me.SipSessionId(xdr, obj)))
}

func (me *XdrReader) Sip(xdr *Xdr) *XdrSip {
	return (*XdrSip)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) SipCallingNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrBinary(xdr, &obj.CallingNumber)
}

func (me *XdrReader) SipCalledNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrBinary(xdr, &obj.CalledNumber)
}

func (me *XdrReader) SipSessionId(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrBinary(xdr, &obj.SessionId)
}
