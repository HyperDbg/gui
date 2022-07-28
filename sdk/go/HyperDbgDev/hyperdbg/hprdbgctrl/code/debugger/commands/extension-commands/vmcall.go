package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\vmcall.cpp"
var vmcallBuf string

type (
	Vmcall interface {
		//Fn() (ok bool)
	}
	vmcall struct{}
)

func Newvmcall() Vmcall { return &vmcall{} }
