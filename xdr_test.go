package libxdr

import (
	. "asdf"
	"reflect"
	"testing"
	"unsafe"
)

func objChecker(t *testing.T, obj ISize, usize uintptr) {
	name := reflect.TypeOf(obj).String()
	size := int(usize)

	if obj.Size() == size {
		t.Logf("%s size = %d", name, obj.Size())
	} else {
		t.Errorf("%s calc-size[%d] != unsafe-size[%d]", name, obj.Size(), size)
	}
}

func TestXdrArray(t *testing.T) {
	obj := &XdrArray{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrString(t *testing.T) {
	obj := &XdrString{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrFile(t *testing.T) {
	obj := &XdrFile{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrCert(t *testing.T) {
	obj := &XdrCert{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrDns(t *testing.T) {
	obj := &XdrDns{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrFtp(t *testing.T) {
	obj := &XdrFtp{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrHttp(t *testing.T) {
	obj := &XdrHttp{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrL7(t *testing.T) {
	obj := &XdrL7{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrMail(t *testing.T) {
	obj := &XdrMail{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrRtsp(t *testing.T) {
	obj := &XdrRtsp{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrSession(t *testing.T) {
	obj4 := &XdrSession4{}
	obj6 := &XdrSession6{}

	objChecker(t, obj4, unsafe.Sizeof(*obj4))
	objChecker(t, obj6, unsafe.Sizeof(*obj6))
}

func TestXdrSip(t *testing.T) {
	obj := &XdrSip{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdrSsl(t *testing.T) {
	obj := &XdrSsl{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestXdr(t *testing.T) {
	obj := &Xdr{}

	objChecker(t, obj, unsafe.Sizeof(*obj))
}

func TestUnmap(t *testing.T) {

}
