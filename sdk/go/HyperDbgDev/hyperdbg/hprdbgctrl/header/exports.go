package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\exports.h"
var exportsBuf string

type (
	Exports interface {
		//Fn() (ok bool)
	}
	exports struct{}
)

func Newexports() Exports { return &exports{} }
