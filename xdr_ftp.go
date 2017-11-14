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
	dump(TabN(tab) + "ftp:")

	tab += 1
	dump(TabN(tab)+"trans-mode:%d", obj.TransMode)
	dump(TabN(tab)+"trans-type:%d", obj.TransType)
	dump(TabN(tab)+"file-size:%d", obj.FileSize)
	dump(TabN(tab)+"response-delay:%d", obj.ResponseDelay)
	dump(TabN(tab)+"duration:%d", obj.TransDuration)

	dump(TabN(tab)+"status:%s", string(me.FtpStatus(xdr, obj)))
	dump(TabN(tab)+"user:%s", string(me.FtpUser(xdr, obj)))
	dump(TabN(tab)+"pwd:%s", string(me.FtpPwd(xdr, obj)))
	dump(TabN(tab)+"filename:%s", string(me.FtpFileName(xdr, obj)))
}

func (me *XdrReader) Ftp(xdr *Xdr) *XdrFtp {
	return (*XdrFtp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) FtpStatus(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrBinary(xdr, &obj.Status)
}

func (me *XdrReader) FtpUser(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrBinary(xdr, &obj.User)
}

func (me *XdrReader) FtpPwd(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrBinary(xdr, &obj.Pwd)
}

func (me *XdrReader) FtpFileName(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrBinary(xdr, &obj.FileName)
}
