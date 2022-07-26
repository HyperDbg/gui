package objects

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\objects\\objects.cpp"
var objectsBuf string

type (
	Objects interface {
		//Fn() (ok bool)
	}
	objects struct{}
)

func Newobjects() Objects { return &objects{} }
