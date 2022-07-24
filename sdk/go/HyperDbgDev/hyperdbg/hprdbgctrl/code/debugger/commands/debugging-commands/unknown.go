package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\~.cpp"
var unknownBuf string

type (
	Unknown interface {
		//Fn() (ok bool)
	}
	unknown struct{}
)

func Newunknown() Unknown { return &unknown{} }
