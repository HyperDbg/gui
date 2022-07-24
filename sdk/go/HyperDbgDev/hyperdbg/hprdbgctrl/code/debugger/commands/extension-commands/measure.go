package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\measure.cpp"
var measureBuf string

type (
	Measure interface {
		//Fn() (ok bool)
	}
	measure struct{}
)

func Newmeasure() Measure { return &measure{} }
