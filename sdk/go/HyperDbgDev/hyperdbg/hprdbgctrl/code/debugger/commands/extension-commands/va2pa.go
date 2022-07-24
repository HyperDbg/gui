package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\va2pa.cpp"
var va2paBuf string

type (
	Va2pa interface {
		//Fn() (ok bool)
	}
	va2pa struct{}
)

func Newva2pa() Va2pa { return &va2pa{} }
