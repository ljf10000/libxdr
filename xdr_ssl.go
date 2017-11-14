package libxdr

import (
	. "asdf"
)

const SizeofXdrCert = SizeofXdrFile +
	2*SizeofByte +
	1*SizeofInt16 +
	1*SizeofPad4 +
	2*SizeofInt64 +
	6*SizeofXdrString

type XdrCert struct {
	File XdrFile

	V        byte
	Version  byte
	KeyUsage uint16
	_        Pad4

	NotBefore uint64
	NotAfter  uint64

	Domain               XdrString
	SerialNumber         XdrString
	CountryName          XdrString
	OrganizationName     XdrString
	OrganizationUnitName XdrString
	CommonName           XdrString
}

func (me *XdrCert) Size() int {
	return SizeofXdrCert
}

const SizeofXdrSsl = 4*SizeofByte + 2*SizeofXdrArray

type XdrSsl struct {
	V      byte
	Reason byte
	_      byte
	_      byte

	CertServer XdrArray // xdrCert
	CertClient XdrArray // xdrCert
}

func (me *XdrSsl) Size() int {
	return SizeofXdrSsl
}

func (me *XdrReader) Ssl(xdr *Xdr) *XdrSsl {
	return (*XdrSsl)(me.xdrMember(xdr, xdr.OffsetofL6))
}

func (me *XdrReader) sslCert(xdr *Xdr, obj *XdrArray, idx int) *XdrCert {
	entry := me.xdrArrayEntry(xdr, obj, idx)

	return (*XdrCert)(entry)
}

func (me *XdrReader) SslServerCert(xdr *Xdr, obj *XdrSsl, idx int) *XdrCert {
	return me.sslCert(xdr, &obj.CertServer, idx)
}

func (me *XdrReader) SslClientCert(xdr *Xdr, obj *XdrSsl, idx int) *XdrCert {
	return me.sslCert(xdr, &obj.CertClient, idx)
}
