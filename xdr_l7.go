package libxdr

import (
	. "asdf"
)

const SizeofXdrL7 = 2*SizeofByte + 1*SizeofInt16

type XdrL7 struct {
	Status   byte
	Class    byte
	Protocol uint16
}

func (me *XdrL7) Size() int {
	return SizeofXdrL7
}

func (me *XdrReader) dumpXdrL7(xdr *Xdr, obj *XdrL7, tab int) {
	dump(TabN(tab) + "L7:")

	tab += 1
	dump(TabN(tab)+"status:%d", obj.Status)
	dump(TabN(tab)+"class:%d", obj.Class)
	dump(TabN(tab)+"protocol:%d", obj.Protocol)
}
