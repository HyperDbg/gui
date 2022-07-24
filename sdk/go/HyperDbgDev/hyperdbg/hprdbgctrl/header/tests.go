package ud

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\tests.h"
var testsBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)

func New() Interface { return &object{} }
