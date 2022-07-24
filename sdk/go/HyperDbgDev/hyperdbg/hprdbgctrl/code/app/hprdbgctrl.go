package app

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\app\\hprdbgctrl.cpp"
var hprdbgctrlBuf string

type (
	Hprdbgctrl interface {
		//Fn() (ok bool)
	}
	hprdbgctrl struct{}
)

func Newhprdbgctrl() Hprdbgctrl { return &hprdbgctrl{} }
