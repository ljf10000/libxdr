package libxdr

import (
	. "asdf"
	"unsafe"
)

const (
	XDR_F_FILE            = 0x0002
	XDR_F_HTTP_REQUEST    = 0x0004
	XDR_F_HTTP_RESPONSE   = 0x0008
	XDR_F_SSL_SERVER_CERT = 0x0010
	XDR_F_SSL_CLIENT_CERT = 0x0020
)

const (
	XDR_CLASS_COMMON = 100
	XDR_CLASS_DNS    = 101
	XDR_CLASS_MMS    = 102
	XDR_CLASS_HTTP   = 103
	XDR_CLASS_FTP    = 104
	XDR_CLASS_MAIL   = 105
	XDR_CLASS_VOIP   = 106
	XDR_CLASS_RTSP   = 107
	XDR_CLASS_P2P    = 108
	XDR_CLASS_VIDEO  = 109
	XDR_CLASS_IM     = 110
	XDR_CLASS_SSL    = 111
)

func XdrAlign(v int) int {
	return AlignI(v, 4)
}

type XdrOffset uint32
type XdrIp4Addr uint32
type XdrIp6Addr []byte

const SizeofXdr = 8*SizeofByte +
	3*SizeofInt64 +
	1*SizeofXdrL7 +
	15*SizeofInt32

type Xdr struct {
	Version byte
	_       byte
	_       byte
	_       byte

	Appid        byte
	IpProto      byte
	IpVersion    byte
	SessionState byte

	SessionTimeCreate uint64
	SessionTimeStart  uint64
	SessionTimeStop   uint64

	L7                  XdrL7
	Bkdr                Bkdr   // session bkdr
	Time                uint32 // time of analysis xdr
	Seq                 uint32
	Flag                uint32 // XDR_F_XXX
	Total               uint32 // total size
	FirstResponseDelay  uint32
	OffsetofL4          XdrOffset // tcp
	OffsetofL5          XdrOffset // http/sip/rtsp/ftp/mail/dns
	OffsetofL6          XdrOffset // ssl
	OffsetofSession     XdrOffset
	OffsetofSessionSt   XdrOffset
	OffsetofServiceSt   XdrOffset
	OffsetofFileContent XdrOffset // xdr_file_t
	OffsetofAlert       XdrOffset
	_                   XdrOffset // padding for align 8
}

func (me *Xdr) Size() int {
	return SizeofXdr
}

func (me *XdrReader) dump(xdr *Xdr) {
	dump("xdr header:")
	dump(Tab+"version:%d", xdr.Version)
	dump(Tab+"total:%d", xdr.Total)
	dump(Tab+"appid:%d", xdr.Appid)
	dump(Tab+"ip-proto:%d", xdr.IpProto)
	dump(Tab+"ip-version:%d", xdr.SessionState)
	dump(Tab+"session-state:%d", xdr.SessionState)
	dump(Tab+"session-time-create:%d", xdr.SessionTimeCreate)
	dump(Tab+"session-time-start:%d", xdr.SessionTimeStart)
	dump(Tab+"session-time-stop:%d", xdr.SessionTimeStop)
	dump(Tab+"bkdr:0x%x", xdr.Bkdr)
	dump(Tab+"time:%d", xdr.Time)
	dump(Tab+"seq:%d", xdr.Seq)
	dump(Tab+"flag:%d", xdr.Flag)
	dump(Tab+"frist-response-delay:%d", xdr.FirstResponseDelay)
	dump(Tab+"offsetof-session:%d", xdr.OffsetofSession)
	dump(Tab+"offsetof-session-st:%d", xdr.OffsetofSessionSt)
	dump(Tab+"offsetof-service-st:%d", xdr.OffsetofServiceSt)
	dump(Tab+"offsetof-allert:%d", xdr.OffsetofAlert)
	dump(Tab+"offsetof-file-content:%d", xdr.OffsetofFileContent)
	dump(Tab+"offsetof-L4:%d", xdr.OffsetofL4)
	dump(Tab+"offsetof-L5:%d", xdr.OffsetofL5)
	dump(Tab+"offsetof-L6:%d", xdr.OffsetofL6)
}

func (me *XdrReader) walk(xdr *Xdr, left uint32, walk XdrWalker) error {
	for left > 0 {
		if xdr.Total < SizeofXdr {
			return ErrBadProto
		} else if left < xdr.Total {
			return ErrTooShortBuffer
		}

		if nil != walk {
			err := walk(me, xdr)
			if nil != err {
				return err
			}
		}

		left -= xdr.Total
		xdr = me.xdrNext(xdr)
	}

	return nil
}

func (me *XdrReader) check(xdr *Xdr, left uint32) error {
	return me.walk(xdr, left, nil)
}

func (me *XdrReader) xdrNext(xdr *Xdr) *Xdr {
	return (*Xdr)(me.object(me.xdrOffset(xdr) + XdrOffset(xdr.Total)))
}

func (me *XdrReader) xdrOffset(xdr *Xdr) XdrOffset {
	return me.offset(unsafe.Pointer(xdr))
}

func (me *XdrReader) xdrMember(xdr *Xdr, offset XdrOffset) unsafe.Pointer {
	if offset > 0 {
		return me.object(me.xdrOffset(xdr) + offset)
	} else {
		return nil
	}
}
