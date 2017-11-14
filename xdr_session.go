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

const SizeofXdrSession4 = SizeofXdrSession + 2*SizeofInt32

type XdrSession4 struct {
	XdrSession

	Sip uint32
	Dip uint32
}

const SizeofXdrSession6 = SizeofXdrSession + 2*16

type XdrSession6 struct {
	XdrSession

	Sip [16]byte
	Dip [16]byte
}

func (me *XdrHandle) Session4(xdr *Xdr) *XdrSession4 {
	return (*XdrSession4)(me.xdrMember(xdr, xdr.OffsetofSession))
}

func (me *XdrHandle) Session6(xdr *Xdr) *XdrSession6 {
	return (*XdrSession6)(me.xdrMember(xdr, xdr.OffsetofSession))
}
