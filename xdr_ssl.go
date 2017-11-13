package libxdr

type XdrCert struct {
	File XdrFile

	Version  byte
	_        byte
	KeyUsage uint16

	NotBefore uint64
	NotAfter  uint64

	Domain               XdrString
	SerialNumber         XdrString
	CountryName          XdrString
	OrganizationName     XdrString
	OrganizationUnitName XdrString
	CommonName           XdrString
}

type XdrSsl struct {
	Reason         byte
	Verfy          byte
	VerfyFailedIdx byte
	_              byte

	VerfyFailedDesc XdrString

	CertServer XdrArray // xdrCert
	CertClient XdrArray // xdrCert
}

func (me *XdrMemFile) Ssl(xdr *Xdr) *XdrSsl {
	return (*XdrSsl)(me.xdrObject(xdr, xdr.OffsetofL6))
}

func (me *XdrMemFile) SslVerfyFailedDesc(xdr *Xdr, obj *XdrSsl) []byte {
	return me.xdrString(xdr, obj.VerfyFailedDesc)
}

func (me *XdrMemFile) sslCerts(xdr *Xdr, obj XdrArray) []*XdrCert {
	if 0 == obj.Count || 0 == obj.Offset || 0 == obj.Size {
		return nil
	}

	body := me.xdrArrayBody(xdr, obj)
	if nil == body {
		return nil
	}

	count := int(obj.Count)
	certs := make([]*XdrCert, 0, count)
	for i := 0; i < count; i++ {
		entry := me.xdrArrayEntry(xdr, body, int(obj.Size), i)
		cert := *(*XdrCert)(entry)

		certs = append(certs, &cert)
	}

	return certs
}

// safe, needn't copy
func (me *XdrMemFile) SslServerCerts(xdr *Xdr, obj *XdrSsl) []*XdrCert {
	return me.sslCerts(xdr, obj.CertServer)
}

// safe, needn't copy
func (me *XdrMemFile) SslClientCerts(xdr *Xdr, obj *XdrSsl) []*XdrCert {
	return me.sslCerts(xdr, obj.CertClient)
}
