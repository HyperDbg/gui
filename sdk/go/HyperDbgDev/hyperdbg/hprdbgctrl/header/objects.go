package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\objects.h"
var objectsBuf string

type (
	Objects interface {
		//Fn() (ok bool)
	}
	objects struct{}
)

func Newobjects() Objects { return &objects{} }
