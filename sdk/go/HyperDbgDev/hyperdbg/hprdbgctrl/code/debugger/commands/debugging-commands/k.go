package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\k.cpp"
var kBuf string

type (
	K interface {
		//Fn() (ok bool)
	}
	k struct{}
)

func Newk() K { return &k{} }
