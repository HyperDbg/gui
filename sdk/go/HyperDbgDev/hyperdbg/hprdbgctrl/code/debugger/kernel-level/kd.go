package kernellevel

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\kernel-level\\kd.cpp"
var kdBuf string

type (
	Kd interface {
		//Fn() (ok bool)
	}
	kd struct{}
)

func Newkd() Kd { return &kd{} }
