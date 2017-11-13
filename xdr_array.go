package libxdr

import (
	. "asdf"
	"unsafe"
)

const (
	XDR_ARRAY_string = 0
	XDR_ARRAY_ip4    = 1
	XDR_ARRAY_ip6    = 2
	XDR_ARRAY_cert   = 3
	XDR_ARRAY_END    = 4
)

type XdrArray struct {
	Size   uint32
	Offset XdrOffset
	Count  uint16
	Type   byte
	_      byte
}

func (me *XdrMemFile) xdrArrayBody(xdr *Xdr, obj XdrArray) unsafe.Pointer {
	if obj.Offset > 0 && obj.Size > 0 && obj.Count > 0 {
		return me.xdrObject(xdr, obj.Offset)
	} else {
		return nil
	}
}

func (me *XdrMemFile) xdrArrayEntry(xdr *Xdr, body unsafe.Pointer, size, idx int) unsafe.Pointer {
	offset := XdrAlign(size) * idx

	return unsafe.Pointer(uintptr(body) + uintptr(offset))
}

func (me *XdrMemFile) xdrArrayEntrySlice(xdr *Xdr, body unsafe.Pointer, size, idx int) []byte {
	entry := me.xdrArrayEntry(xdr, body, size, idx)

	return PointerToSlice(entry, size, size)
}
