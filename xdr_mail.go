package libxdr

import (
	. "asdf"
)

const SizeofXdrMail = 4*SizeofByte +
	2*SizeofInt16 +
	1*SizeofInt32 +
	5*SizeofXdrString

type XdrMail struct {
	V       byte
	AcsType byte
	_       byte
	_       byte

	MsgType    uint16
	StatusCode int16
	Length     uint32

	User   XdrString
	Domain XdrString
	Sender XdrString
	Recver XdrString
	Hdr    XdrString
}

func (me *XdrHandle) Mail(xdr *Xdr) *XdrMail {
	return (*XdrMail)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrHandle) MailUser(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrHandle) MailDomain(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Domain)
}

func (me *XdrHandle) MailSender(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Sender)
}

func (me *XdrHandle) MailRecver(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Recver)
}

func (me *XdrHandle) MailHdr(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Hdr)
}
