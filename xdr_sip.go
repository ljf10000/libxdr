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

func (me *XdrReader) Sip(xdr *Xdr) *XdrSip {
	return (*XdrSip)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) SipCallingNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.CallingNumber)
}

func (me *XdrReader) SipCalledNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.CalledNumber)
}

func (me *XdrReader) SipSessionId(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.SessionId)
}
