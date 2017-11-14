package libxdr

import (
	. "asdf"
)

const SizeofXdrSession = 4*SizeofByte + 2*SizeofInt16

type XdrSession struct {
	Version byte
	Dir     byte
	Proto   byte
	_       byte

	Sport uint16
	Dport uint16
}

func (me *XdrSession) Size() int {
	return SizeofXdrSession
}

func (me *XdrReader) dumpXdrSession(xdr *Xdr, obj *XdrSession, tab int) {

}

const SizeofXdrSession4 = SizeofXdrSession + 2*SizeofInt32

type XdrSession4 struct {
	XdrSession

	Sip uint32
	Dip uint32
}

func (me *XdrSession4) Size() int {
	return SizeofXdrSession4
}

func (me *XdrReader) dumpSession4(xdr *Xdr, obj *XdrSession4, tab int) {

}

const SizeofXdrSession6 = SizeofXdrSession + 2*16

type XdrSession6 struct {
	XdrSession

	Sip [16]byte
	Dip [16]byte
}

func (me *XdrSession6) Size() int {
	return SizeofXdrSession6
}

func (me *XdrReader) dumpSession6(xdr *Xdr, obj *XdrSession6, tab int) {

}

func (me *XdrReader) Session4(xdr *Xdr) *XdrSession4 {
	return (*XdrSession4)(me.xdrMember(xdr, xdr.OffsetofSession))
}

func (me *XdrReader) Session6(xdr *Xdr) *XdrSession6 {
	return (*XdrSession6)(me.xdrMember(xdr, xdr.OffsetofSession))
}
