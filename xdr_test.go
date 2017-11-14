package libxdr

import (
	"reflect"
	"testing"
	"unsafe"
)

func sizeChecker(t *testing.T, csize int, osize uintptr, Type reflect.Type) {
	if csize == int(osize) {
		t.Logf("%s size == %d", Type.String(), csize)
	} else {
		t.Errorf("%s size != %d", Type.String(), csize)
	}
}

func TestXdrArray(t *testing.T) {
	obj := XdrArray{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrString(t *testing.T) {
	obj := XdrString{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrFile(t *testing.T) {
	obj := XdrFile{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrCert(t *testing.T) {
	obj := XdrCert{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrDns(t *testing.T) {
	obj := XdrDns{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrFtp(t *testing.T) {
	obj := XdrFtp{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrHttp(t *testing.T) {
	obj := XdrHttp{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrL7(t *testing.T) {
	obj := XdrL7{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrMail(t *testing.T) {
	obj := XdrMail{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrRtsp(t *testing.T) {
	obj := XdrRtsp{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrSession(t *testing.T) {
	obj4 := XdrSession4{}
	obj6 := XdrSession6{}

	sizeChecker(t, obj4.Size(), unsafe.Sizeof(obj4), reflect.TypeOf(obj4))
	sizeChecker(t, obj6.Size(), unsafe.Sizeof(obj6), reflect.TypeOf(obj6))
}

func TestXdrSip(t *testing.T) {
	obj := XdrSip{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdrSsl(t *testing.T) {
	obj := XdrSsl{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestXdr(t *testing.T) {
	obj := Xdr{}

	sizeChecker(t, obj.Size(), unsafe.Sizeof(obj), reflect.TypeOf(obj))
}

func TestUnmap(t *testing.T) {

}
