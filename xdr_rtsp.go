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

func (me *XdrReader) dumpRtsp(xdr *Xdr, obj *XdrRtsp, tab int) {
	dump(TabN(tab) + "rtsp:")

	tab += 1
	dump(TabN(tab)+"port-client-start:%d", obj.PortClientStart)
	dump(TabN(tab)+"port-client-end:%d", obj.PortClientEnd)
	dump(TabN(tab)+"port-server-start:%d", obj.PortServerStart)
	dump(TabN(tab)+"port-server-end:%d", obj.PortServerEnd)
	dump(TabN(tab)+"count-video:%d", obj.CountVideo)
	dump(TabN(tab)+"count-audio:%d", obj.CountAudio)
	dump(TabN(tab)+"describe-delay:%d", obj.DescribeDelay)

	dump(TabN(tab)+"url:%s", string(me.RtspUrl(xdr, obj)))
	dump(TabN(tab)+"user-agent:%s", string(me.RtspUserAgent(xdr, obj)))
	dump(TabN(tab)+"server-ip:%s", string(me.RtspServerIp(xdr, obj)))

}

func (me *XdrReader) Rtsp(xdr *Xdr) *XdrRtsp {
	return (*XdrRtsp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) RtspUrl(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrBinary(xdr, &obj.Url)
}

func (me *XdrReader) RtspUserAgent(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrBinary(xdr, &obj.UserAgent)
}

func (me *XdrReader) RtspServerIp(xdr *Xdr, obj *XdrRtsp) []byte {
	return me.xdrBinary(xdr, &obj.ServerIp)
}
