package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\output.cpp"
var outputBuf string

type (
	Output interface {
		//Fn() (ok bool)
	}
	output struct{}
)

func Newoutput() Output { return &output{} }
