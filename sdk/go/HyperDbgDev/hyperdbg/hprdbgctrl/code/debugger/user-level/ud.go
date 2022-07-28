package userlevel

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\user-level\\ud.cpp"
var udBuf string

type (
	Ud interface {
		//Fn() (ok bool)
	}
	ud struct{}
)

func Newud() Ud { return &ud{} }
