package objects

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\objects\\objects.cpp"
var objectsBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)

func New() Interface { return &object{} }
