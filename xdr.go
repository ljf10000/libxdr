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

type Xdr struct {
	Version byte
	_       byte
	_       byte
	_       byte

	Appid        byte
	IpProto      byte
	SessionState byte
	IpVersion    byte

	Bkdr               Bkdr
	Time               uint32
	Seq                uint32
	Flag               uint32 // XDR_F_XXX
	Total              uint32 // total size
	FirstResponseDelay uint32

	SessionTimeCreate uint64
	SessionTimeStart  uint64
	SessionTimeStop   uint64

	OffsetofSession     XdrOffset
	OffsetofSessionSt   XdrOffset
	OffsetofServiceSt   XdrOffset
	OffsetofAlert       XdrOffset
	OffsetofFileContent XdrOffset // xdr_file_t

	OffsetofL4 XdrOffset // tcp
	OffsetofL5 XdrOffset // http/sip/rtsp/ftp/mail/dns
	OffsetofL6 XdrOffset // ssl

	L7 XdrL7
}

type XdrMemFile = mmap

type XdrWalker func(mm *XdrMemFile, xdrs *Xdr) error

func (me *XdrMemFile) Walk(filename string, walk XdrWalker) error {
	err := me.open(filename)
	if nil != err {
		return err
	}
	defer me.close()

	var xdr *Xdr

	left := me.size
	for left > 0 {
		xdr = (*Xdr)(me.addr)

		err = walk(me, xdr)
		if nil != err {
			return err
		}
	}

	return nil
}

func (me *XdrMemFile) xdrOffset(xdr *Xdr) XdrOffset {
	return me.offset(unsafe.Pointer(xdr))
}

func (me *XdrMemFile) xdrObject(xdr *Xdr, offset XdrOffset) unsafe.Pointer {
	if offset > 0 {
		return me.object(me.xdrOffset(xdr) + offset)
	} else {
		return nil
	}
}
