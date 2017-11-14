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

const SizeofXdrArray = 2*SizeofByte + SizeofInt16 + 2*SizeofInt32

type XdrArray struct {
	V      byte
	Type   byte
	Count  uint16
	Size   uint32
	Offset XdrOffset
}

func (me *XdrArray) IsEmpty() bool {
	return 0 == me.Count
}

func (me *XdrArray) IsGood() bool {
	return me.Offset > 0 && me.Size > 0
}

func (me *XdrArray) HaveEntry() bool {
	return me.IsGood() && !me.IsEmpty()
}

func (me *XdrReader) xdrArrayBody(xdr *Xdr, obj *XdrArray) unsafe.Pointer {
	if obj.HaveEntry() {
		return me.xdrMember(xdr, obj.Offset)
	} else {
		return nil
	}
}

func (me *XdrReader) xdrArrayEntry(xdr *Xdr, obj *XdrArray, idx int) unsafe.Pointer {
	if idx < int(obj.Count) {
		body := me.xdrArrayBody(xdr, obj)
		if nil != body {
			offset := XdrAlign(int(obj.Size)) * idx

			return unsafe.Pointer(uintptr(body) + uintptr(offset))
		}
	}

	return nil
}
