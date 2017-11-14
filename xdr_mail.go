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

func (me *XdrMail) Size() int {
	return SizeofXdrMail
}

func (me *XdrReader) dumpMail(xdr *Xdr, obj *XdrMail, tab int) {

}

func (me *XdrReader) Mail(xdr *Xdr) *XdrMail {
	return (*XdrMail)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) MailUser(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrReader) MailDomain(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Domain)
}

func (me *XdrReader) MailSender(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Sender)
}

func (me *XdrReader) MailRecver(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Recver)
}

func (me *XdrReader) MailHdr(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Hdr)
}
