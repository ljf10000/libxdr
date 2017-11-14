package libxdr

import (
	. "asdf"
)

const SizeofXdrString = 2 * SizeofInt32

type XdrString struct {
	Size   uint32
	Offset XdrOffset
}

type XdrBinary = XdrString

func (me *XdrReader) xdrString(xdr *Xdr, xstr XdrString) []byte {
	if 0 == xstr.Size || 0 == xstr.Offset {
		return nil
	}

	return ObjToSlice(me.xdrMember(xdr, xstr.Offset), int(xstr.Size))
}
