package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\wrmsr.cpp"
var wrmsrBuf string

type (
	Wrmsr interface {
		//Fn() (ok bool)
	}
	wrmsr struct{}
)

func Newwrmsr() Wrmsr { return &wrmsr{} }
