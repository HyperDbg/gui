package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\ud.h"
var udBuf string

type (
	Ud interface {
		//Fn() (ok bool)
	}
	ud struct{}
)

func Newud() Ud { return &ud{} }
