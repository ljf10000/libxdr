package libxdr

import (
	. "asdf"
)

const SizeofXdrHttp = 4*SizeofInt64 +
	1*SizeofInt32 +
	1*SizeofInt16 +
	10*SizeofByte +
	2*SizeofInt32 +
	8*SizeofXdrString

type XdrHttp struct {
	TimeRequest       uint64
	TimeFirstResponse uint64
	TimeLastContent   uint64
	ServiceDelay      uint64

	ContentLength uint32
	StatusCode    uint16
	Method        byte
	Version       byte

	Flags  byte
	Ie     byte
	Portal byte
	_      byte

	V byte
	_ byte
	_ byte
	_ byte

	OffsetofRequest  uint32
	OffsetofResponse uint32

	Host        XdrString
	Url         XdrString
	HostXonline XdrString
	UserAgent   XdrString
	Content     XdrString
	Refer       XdrString
	Cookie      XdrString
	Location    XdrString
}

func (me *XdrHttp) Size() int {
	return SizeofXdrHttp
}

func (me *XdrReader) dumpHttp(xdr *Xdr, obj *XdrHttp, tab int) {

}

func (me *XdrReader) Http(xdr *Xdr) *XdrHttp {
	return (*XdrHttp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) HttpHost(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Host)
}

func (me *XdrReader) HttpUrl(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Url)
}

func (me *XdrReader) HttpHostXonline(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.HostXonline)
}

func (me *XdrReader) HttpUserAgent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.UserAgent)
}

func (me *XdrReader) HttpContent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Content)
}

func (me *XdrReader) HttpRefer(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Refer)
}

func (me *XdrReader) HttpCookie(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Cookie)
}

func (me *XdrReader) HttpLocation(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Location)
}
