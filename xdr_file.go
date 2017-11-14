package libxdr

import (
	. "asdf"
	"encoding/hex"
)

const SizeofXdrDigest = 8 * SizeofInt32

type XdrDigest []byte

const SizeofXdrFile = 2*SizeofInt32 + 1*DigestSize

type XdrFile struct {
	Len    uint32
	Bkdr   Bkdr
	Digest [DigestSize]byte
}

func (me *XdrFile) Size() int {
	return SizeofXdrFile
}

func (me *XdrReader) dumpXdrFile(xdr *Xdr, obj *XdrFile, tab int) {
	dump(TabN(tab)+"len:%d", obj.Len)
	dump(Tab2+"bkdr:0x%x", obj.Bkdr)
	dump(Tab2+"digest:%s", hex.EncodeToString(obj.Digest[:]))
}

func (me *XdrReader) xdrFile(xdr *Xdr, offset XdrOffset) *XdrFile {
	entry := me.xdrMember(xdr, offset)

	return (*XdrFile)(entry)
}
