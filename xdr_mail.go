package libxdr

type XdrMail struct {
	MsgType    uint16
	StatusCode int16
	Length     uint32

	AcsType byte
	_       byte
	_       byte
	_       byte

	User   XdrString
	Domain XdrString
	Sender XdrString
	Recver XdrString
	Hdr    XdrString
}

func (me *XdrMemFile) Mail(xdr *Xdr) *XdrMail {
	return (*XdrMail)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) MailUser(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrMemFile) MailDomain(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Domain)
}

func (me *XdrMemFile) MailSender(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Sender)
}

func (me *XdrMemFile) MailRecver(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Recver)
}

func (me *XdrMemFile) MailHdr(xdr *Xdr, obj *XdrMail) []byte {
	return me.xdrString(xdr, obj.Hdr)
}
