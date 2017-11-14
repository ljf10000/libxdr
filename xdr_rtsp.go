package libxdr

import (
	. "asdf"
)

const SizeofXdrRtsp = 6*SizeofInt16 +
	1*SizeofInt32 +
	3*SizeofXdrString

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

func (me *XdrRtsp) Size() int {
	return SizeofXdrRtsp
}

func (me *XdrReader) Rtsp(xdr *Xdr) *XdrRtsp {
	return (*XdrRtsp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) RtspUrl(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.Url)
}

func (me *XdrReader) RtspUserAgent(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.UserAgent)
}

func (me *XdrReader) RtspServerIp(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrString(xdr, obj.ServerIp)
}
