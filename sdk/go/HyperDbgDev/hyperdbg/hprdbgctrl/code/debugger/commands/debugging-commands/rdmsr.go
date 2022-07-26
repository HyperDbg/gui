package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\rdmsr.cpp"
var rdmsrBuf string

type (
	Rdmsr interface {
		//Fn() (ok bool)
	}
	rdmsr struct{}
)

func Newrdmsr() Rdmsr { return &rdmsr{} }
