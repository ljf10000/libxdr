package libxdr

import (
	. "asdf"
)

type XdrString struct {
	Size   uint32
	Offset XdrOffset
}

type XdrBinary = XdrString

func (me *XdrMemFile) xdrString(xdr *Xdr, obj XdrString) []byte {
	if obj.Size > 0 && obj.Offset > 0 {
		size := int(obj.Size)

		return PointerToSlice(me.xdrObject(xdr, obj.Offset), size, size)
	} else {
		return nil
	}
}
