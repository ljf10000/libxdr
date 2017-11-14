package libxdr

import (
	. "asdf"
)

const SizeofXdrString = 2 * SizeofInt32

type XdrString struct {
	Len    uint32
	Offset XdrOffset
}

func (me *XdrString) Size() int {
	return SizeofXdrString
}

type XdrBinary = XdrString

func (me *XdrReader) xdrString(xdr *Xdr, xstr XdrString) []byte {
	if 0 == xstr.Len || 0 == xstr.Offset {
		return nil
	}

	return ObjToSlice(me.xdrMember(xdr, xstr.Offset), int(xstr.Len))
}
