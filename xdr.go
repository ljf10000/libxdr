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
	SessionState byte
	IpVersion    byte

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
	OffsetofSession     XdrOffset
	OffsetofSessionSt   XdrOffset
	OffsetofServiceSt   XdrOffset
	OffsetofAlert       XdrOffset
	OffsetofFileContent XdrOffset // xdr_file_t
	OffsetofL4          XdrOffset // tcp
	OffsetofL5          XdrOffset // http/sip/rtsp/ftp/mail/dns
	OffsetofL6          XdrOffset // ssl
	_                   XdrOffset // padding for align 8
}

type XdrHandle = mmap

type XdrWalker func(mm *XdrHandle, xdrs *Xdr) error

func (me *XdrHandle) walk(xdr *Xdr, left uint32, walk XdrWalker) error {
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

func (me *XdrHandle) check(xdr *Xdr, left uint32) error {
	return me.walk(xdr, left, nil)
}

func (me *XdrHandle) Walk(filename string, walk XdrWalker) error {
	err := me.open(filename)
	if nil != err {
		return err
	}
	defer me.close()

	xdr := (*Xdr)(me.addr)
	left := me.size

	if err := me.check(xdr, left); nil != err {
		return err
	}

	return me.walk(xdr, left, walk)
}

func (me *XdrHandle) xdrNext(xdr *Xdr) *Xdr {
	return (*Xdr)(me.object(me.xdrOffset(xdr) + XdrOffset(xdr.Total)))
}

func (me *XdrHandle) xdrOffset(xdr *Xdr) XdrOffset {
	return me.offset(unsafe.Pointer(xdr))
}

func (me *XdrHandle) xdrMember(xdr *Xdr, offset XdrOffset) unsafe.Pointer {
	if offset > 0 {
		return me.object(me.xdrOffset(xdr) + offset)
	} else {
		return nil
	}
}
