package libxdr

type XdrDumper func(format string, v ...interface{})

type XdrWalker func(r *XdrReader, xdrs *Xdr) error

func SetupDumper(f XdrDumper) {
	dump = f
}

func Walk(r *XdrReader, filename string, walk XdrWalker) error {
	if err := r.open(filename); nil != err {
		return err
	}
	defer r.close()

	xdr := (*Xdr)(r.addr)
	left := r.size

	if err := r.check(xdr, left); nil != err {
		return err
	}

	return r.walk(xdr, left, walk)
}

func Dump(r *XdrReader, xdr *Xdr) {
	r.dump(xdr)
}
