package libxdr

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

func (me *XdrMemFile) Sip(xdr *Xdr) *XdrSip {
	return (*XdrSip)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) SipCallingNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.CallingNumber)
}

func (me *XdrMemFile) SipCalledNumber(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.CalledNumber)
}

func (me *XdrMemFile) SipSessionId(xdr *Xdr, obj *XdrSip) []byte {
	return me.xdrString(xdr, obj.SessionId)
}
