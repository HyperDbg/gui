package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\s.cpp"
var sBuf string

type (
	S interface {
		//Fn() (ok bool)
	}
	s struct{}
)

func News() S { return &s{} }
