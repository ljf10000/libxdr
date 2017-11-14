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

	OffsetofRequest  XdrOffset
	OffsetofResponse XdrOffset

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
	dump(TabN(tab) + "ftp:")

	tab += 1
	dump(TabN(tab)+"time-request:%d", obj.TimeRequest)
	dump(TabN(tab)+"time-response:%d", obj.TimeFirstResponse)
	dump(TabN(tab)+"time-last-content:%d", obj.TimeLastContent)
	dump(TabN(tab)+"service-delay:%d", obj.ServiceDelay)
	dump(TabN(tab)+"content-length:%d", obj.ContentLength)
	dump(TabN(tab)+"status-code:%d", obj.StatusCode)
	dump(TabN(tab)+"method:%d", obj.Method)
	dump(TabN(tab)+"version:%d", obj.Version)
	dump(TabN(tab)+"flags:%d", obj.Flags)
	dump(TabN(tab)+"ie:%d", obj.Ie)
	dump(TabN(tab)+"portal:%d", obj.Portal)

	dump(TabN(tab)+"host:%s", string(me.HttpHost(xdr, obj)))
	dump(TabN(tab)+"url:%s", string(me.HttpUrl(xdr, obj)))
	dump(TabN(tab)+"host-xonline:%s", string(me.HttpHostXonline(xdr, obj)))
	dump(TabN(tab)+"user-agent:%s", string(me.HttpUserAgent(xdr, obj)))
	dump(TabN(tab)+"content:%s", string(me.HttpContent(xdr, obj)))
	dump(TabN(tab)+"refer:%s", string(me.HttpRefer(xdr, obj)))
	dump(TabN(tab)+"cookie:%s", string(me.HttpCookie(xdr, obj)))
	dump(TabN(tab)+"location:%s", string(me.HttpLocation(xdr, obj)))

	if obj.OffsetofRequest > 0 {
		file := me.HttpRequest(xdr, obj)

		dump(TabN(tab) + "request:")
		me.dumpXdrFile(xdr, file, 1+tab)
	}

	if obj.OffsetofResponse > 0 {
		file := me.HttpResponse(xdr, obj)

		dump(TabN(tab) + "response:")
		me.dumpXdrFile(xdr, file, 1+tab)
	}
}

func (me *XdrReader) Http(xdr *Xdr) *XdrHttp {
	return (*XdrHttp)(me.xdrMember(xdr, xdr.OffsetofL5))
}

func (me *XdrReader) HttpRequest(xdr *Xdr, obj *XdrHttp) *XdrFile {
	return me.xdrFile(xdr, obj.OffsetofRequest)
}

func (me *XdrReader) HttpResponse(xdr *Xdr, obj *XdrHttp) *XdrFile {
	return me.xdrFile(xdr, obj.OffsetofResponse)
}

func (me *XdrReader) HttpHost(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Host)
}

func (me *XdrReader) HttpUrl(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Url)
}

func (me *XdrReader) HttpHostXonline(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.HostXonline)
}

func (me *XdrReader) HttpUserAgent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.UserAgent)
}

func (me *XdrReader) HttpContent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Content)
}

func (me *XdrReader) HttpRefer(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Refer)
}

func (me *XdrReader) HttpCookie(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Cookie)
}

func (me *XdrReader) HttpLocation(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrBinary(xdr, &obj.Location)
}
