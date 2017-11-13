package libxdr

type XdrSession struct {
	Version byte
	Dir     byte
	Proto   byte
	_       byte
	Sport   uint16
	Dport   uint16
}

type XdrSession4 struct {
	XdrSession

	Sip uint32
	Dip uint32
}

type XdrSession6 struct {
	XdrSession

	Sip0 uint32
	Sip1 uint32
	Sip2 uint32
	Sip3 uint32
	Dip0 uint32
	Dip1 uint32
	Dip2 uint32
	Dip3 uint32
}

func (me *XdrMemFile) Session4(xdr *Xdr) *XdrSession4 {
	return (*XdrSession4)(me.xdrObject(xdr, xdr.OffsetofSession))
}

func (me *XdrMemFile) Session6(xdr *Xdr) *XdrSession6 {
	return (*XdrSession6)(me.xdrObject(xdr, xdr.OffsetofSession))
}
