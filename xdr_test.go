package libxdr

import (
	"testing"
	"unsafe"
)

func TestXdrArray(t *testing.T) {
	obj := XdrArray{}

	if SizeofXdrArray != unsafe.Sizeof(obj) {
		t.Errorf("XdrArray size must %d", SizeofXdrArray)
	}
}

func TestXdrString(t *testing.T) {
	obj := XdrString{}

	if SizeofXdrString != unsafe.Sizeof(obj) {
		t.Errorf("XdrString size must %d", SizeofXdrString)
	}
}

func TestXdrFile(t *testing.T) {
	obj := XdrFile{}

	if SizeofXdrFile != unsafe.Sizeof(obj) {
		t.Errorf("XdrFile size must %d", SizeofXdrFile)
	}
}

func TestXdrCert(t *testing.T) {
	obj := XdrCert{}

	if SizeofXdrCert != unsafe.Sizeof(obj) {
		t.Errorf("XdrCert size must %d", SizeofXdrCert)
	}
}

func TestXdrDns(t *testing.T) {
	obj := XdrDns{}

	if SizeofXdrDns != unsafe.Sizeof(obj) {
		t.Errorf("XdrDns size must %d", SizeofXdrDns)
	}
}

func TestXdrFtp(t *testing.T) {
	obj := XdrFtp{}

	if SizeofXdrFtp != unsafe.Sizeof(obj) {
		t.Errorf("XdrFtp size must %d", SizeofXdrFtp)
	}
}

func TestXdrHttp(t *testing.T) {
	obj := XdrHttp{}

	if SizeofXdrHttp != unsafe.Sizeof(obj) {
		t.Errorf("XdrHttp size must %d", SizeofXdrHttp)
	}
}

func TestXdrL7(t *testing.T) {
	obj := XdrL7{}

	if SizeofXdrL7 != unsafe.Sizeof(obj) {
		t.Errorf("XdrL7 size must %d", SizeofXdrL7)
	}
}

func TestXdrMail(t *testing.T) {
	obj := XdrMail{}

	if SizeofXdrMail != unsafe.Sizeof(obj) {
		t.Errorf("XdrMail size must %d", SizeofXdrMail)
	}
}

func TestXdrRtsp(t *testing.T) {
	obj := XdrRtsp{}

	if SizeofXdrRtsp != unsafe.Sizeof(obj) {
		t.Errorf("XdrRtsp size must %d", SizeofXdrRtsp)
	}
}

func TestXdrSession(t *testing.T) {
	obj4 := XdrSession4{}
	obj6 := XdrSession6{}

	if SizeofXdrSession4 != unsafe.Sizeof(obj4) {
		t.Errorf("XdrSession4 size must %d", SizeofXdrSession4)
	}

	if SizeofXdrSession6 != unsafe.Sizeof(obj6) {
		t.Errorf("XdrSession6 size must %d", SizeofXdrSession6)
	}
}

func TestXdrSip(t *testing.T) {
	obj := XdrSip{}

	if SizeofXdrSip != unsafe.Sizeof(obj) {
		t.Errorf("XdrSip size must %d", SizeofXdrSip)
	}
}

func TestXdrSsl(t *testing.T) {
	obj := XdrSsl{}

	if SizeofXdrSsl != unsafe.Sizeof(obj) {
		t.Errorf("XdrSsl size must %d", SizeofXdrSsl)
	}
}

func TestXdr(t *testing.T) {
	obj := Xdr{}

	if SizeofXdr != unsafe.Sizeof(obj) {
		t.Errorf("Xdr size must %d", SizeofXdr)
	}
}

func TestUnmap(t *testing.T) {

}
