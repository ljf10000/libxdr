package libxdr

import (
	"os"
	"reflect"
	"unsafe"

	mgo "github.com/edsrzf/mmap-go"
)

type mmap struct {
	filename string
	f        *os.File
	mm       mgo.MMap
	addr     unsafe.Pointer
	size     uint32
}

func (me *mmap) open(filename string) error {
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

func (me *mmap) close() {
	if len(me.mm) > 0 {
		me.mm.Unmap()
	}

	if nil != me.f {
		me.f.Close()
	}
}

func (me *mmap) object(offset XdrOffset) unsafe.Pointer {
	return unsafe.Pointer(uintptr(me.addr) + uintptr(offset))
}

func (me *mmap) offset(obj unsafe.Pointer) XdrOffset {
	return XdrOffset(uintptr(obj) - uintptr(me.addr))
}
