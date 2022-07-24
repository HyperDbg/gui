package ud

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\transparency.h"
var transparencyBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)

func New() Interface { return &object{} }
