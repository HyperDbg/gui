package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\test.cpp"
var testBuf string

type (
	Test interface {
		//Fn() (ok bool)
	}
	test struct{}
)

func Newtest() Test { return &test{} }
