package libxdr

import (
	. "asdf"
)

const SizeofXdrDigest = 8 * SizeofInt32

type XdrDigest []byte

const SizeofXdrFile = 2*SizeofInt32 + 1*DigestSize

type XdrFile struct {
	Size   uint32
	Bkdr   Bkdr
	Digest [DigestSize]byte
}
