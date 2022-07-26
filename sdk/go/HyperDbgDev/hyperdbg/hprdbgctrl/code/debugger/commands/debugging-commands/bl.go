package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\bl.cpp"
var blBuf string

type (
	Bl interface {
		//Fn() (ok bool)
	}
	bl struct{}
)

func Newbl() Bl { return &bl{} }
