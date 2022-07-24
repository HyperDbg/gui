package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\dr.cpp"
var drBuf string

type (
	Dr interface {
		//Fn() (ok bool)
	}
	dr struct{}
)

func Newdr() Dr { return &dr{} }
