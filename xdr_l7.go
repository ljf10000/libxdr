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
