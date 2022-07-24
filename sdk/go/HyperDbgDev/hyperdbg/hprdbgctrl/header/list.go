package ud

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\list.h"
var listBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)

func New() Interface { return &object{} }
