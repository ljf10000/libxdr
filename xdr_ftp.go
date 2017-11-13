package libxdr

type XdrFtp struct {
	TransMode byte
	TransType byte
	_         byte
	_         byte

	FileSize      uint32
	ResponseDelay uint64
	TransDuration uint64

	Status   XdrString
	User     XdrString
	Pwd      XdrString
	FileName XdrString
}

func (me *XdrMemFile) Ftp(xdr *Xdr) *XdrFtp {
	return (*XdrFtp)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) FtpStatus(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Status)
}

func (me *XdrMemFile) FtpUser(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.User)
}

func (me *XdrMemFile) FtpPwd(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.Pwd)
}

func (me *XdrMemFile) FtpFileName(xdr *Xdr, obj *XdrFtp) []byte {
	return me.xdrString(xdr, obj.FileName)
}
