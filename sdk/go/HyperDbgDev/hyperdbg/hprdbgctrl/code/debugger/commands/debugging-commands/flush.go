package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\flush.cpp"
var flushBuf string

type (
	Flush interface {
		//Fn() (ok bool)
	}
	flush struct{}
)

func Newflush() Flush { return &flush{} }
