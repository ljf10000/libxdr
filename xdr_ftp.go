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

func (me *XdrFtp) Size() int {
	return SizeofXdrFtp
}

func (me *XdrReader) dumpFtp(xdr *Xdr, obj *XdrFtp, tab int) {

}

func (me *XdrReader) Ftp(xdr *Xdr) *XdrFtp {
	return (*XdrFtp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) FtpStatus(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Status)
}

func (me *XdrReader) FtpUser(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrReader) FtpPwd(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Pwd)
}

func (me *XdrReader) FtpFileName(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.FileName)
}
