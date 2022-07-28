package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\globals.h"
var globalsBuf string

type (
	Globals interface {
		//Fn() (ok bool)
	}
	globals struct{}
)

func Newglobals() Globals { return &globals{} }
