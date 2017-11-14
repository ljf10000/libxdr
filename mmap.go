package libxdr

import (
	"os"
	"reflect"
	"unsafe"

	mgo "github.com/edsrzf/mmap-go"
)

type XdrReader struct {
	filename string
	f        *os.File
	mm       mgo.MMap
	addr     unsafe.Pointer
	size     uint32
}

func (me *XdrReader) open(filename string) error {
	f, err := os.Open(filename)
	if nil != err {
		return err
	}

	mm, err := mgo.Map(f, mgo.RDONLY, 0)
	if nil != err {
		return err
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&mm))
	me.addr = unsafe.Pointer(hdr.Data)
	me.size = uint32(hdr.Len)

	me.filename = filename
	me.mm = mm
	me.f = f

	return nil
}

func (me *XdrReader) close() {
	if len(me.mm) > 0 {
		me.mm.Unmap()
	}

	if nil != me.f {
		me.f.Close()
	}
}

func (me *XdrReader) object(offset XdrOffset) unsafe.Pointer {
	return unsafe.Pointer(uintptr(me.addr) + uintptr(offset))
}

func (me *XdrReader) offset(obj unsafe.Pointer) XdrOffset {
	return XdrOffset(uintptr(obj) - uintptr(me.addr))
}
