package app

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\app\\dllmain.cpp"
var dllmainBuf string

type (
	Dllmain interface {
		//Fn() (ok bool)
	}
	dllmain struct{}
)

func Newdllmain() Dllmain { return &dllmain{} }
