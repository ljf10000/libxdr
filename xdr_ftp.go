package libxdr

import (
	. "asdf"
)

const SizeofXdrFtp = 4*SizeofByte + 1*SizeofInt32 + 2*SizeofInt64 + 4*SizeofXdrString

type XdrFtp struct {
	V         byte
	TransMode byte
	TransType byte
	_         byte

	FileSize uint32

	ResponseDelay uint64
	TransDuration uint64

	Status   XdrString
	User     XdrString
	Pwd      XdrString
	FileName XdrString
}

func (me *XdrHandle) Ftp(xdr *Xdr) *XdrFtp {
	return (*XdrFtp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrHandle) FtpStatus(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Status)
}

func (me *XdrHandle) FtpUser(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrHandle) FtpPwd(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Pwd)
}

func (me *XdrHandle) FtpFileName(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.FileName)
}
