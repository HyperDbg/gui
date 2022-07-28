package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\break-control.cpp"
var breakControlBuf string

type (
	BreakControl interface {
		BreakController() (ok bool)
	}
	breakControl struct{}
)

func (b *breakControl) BreakController() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func NewbreakControl() BreakControl { return &breakControl{} }
