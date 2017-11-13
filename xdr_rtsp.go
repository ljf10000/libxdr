package libxdr

type XdrRtsp struct {
	PortClientStart uint16
	PortClientEnd   uint16
	PortServerStart uint16
	PortServerEnd   uint16
	CountVideo      uint16
	CountAudio      uint16

	DescribeDelay uint32

	Url       XdrString
	UserAgent XdrString
	ServerIp  XdrString
}

func (me *XdrMemFile) Rtsp(xdr *Xdr) *XdrRtsp {
	return (*XdrRtsp)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) RtspUrl(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.Url)
}

func (me *XdrMemFile) RtspUserAgent(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.UserAgent)
}

func (me *XdrMemFile) RtspServerIp(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.ServerIp)
}
