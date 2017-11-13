package libxdr

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

func (me *XdrMemFile) Http(xdr *Xdr) *XdrHttp {
	return (*XdrHttp)(me.xdrObject(xdr, xdr.OffsetofL5))
}

func (me *XdrMemFile) HttpHost(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Host)
}

func (me *XdrMemFile) HttpUrl(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Url)
}

func (me *XdrMemFile) HttpHostXonline(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.HostXonline)
}

func (me *XdrMemFile) HttpUserAgent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.UserAgent)
}

func (me *XdrMemFile) HttpContent(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Content)
}

func (me *XdrMemFile) HttpRefer(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Refer)
}

func (me *XdrMemFile) HttpCookie(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Cookie)
}

func (me *XdrMemFile) HttpLocation(xdr *Xdr, obj *XdrHttp) []byte {
	return me.xdrString(xdr, obj.Location)
}
