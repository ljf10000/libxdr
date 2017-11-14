package libxdr

import (
	. "asdf"
)

var dump XdrDumper

func init() {
	dump = Log.Info
}
