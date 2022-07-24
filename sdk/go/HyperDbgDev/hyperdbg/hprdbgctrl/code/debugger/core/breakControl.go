package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\break-control.cpp"
var breakControlBuf string

type (
	BreakControl interface {
		//Fn() (ok bool)
	}
	breakControl struct{}
)

func NewbreakControl() BreakControl { return &breakControl{} }
