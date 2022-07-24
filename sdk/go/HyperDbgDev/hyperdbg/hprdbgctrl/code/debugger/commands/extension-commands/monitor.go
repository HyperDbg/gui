package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\monitor.cpp"
var monitorBuf string

type (
	Monitor interface {
		//Fn() (ok bool)
	}
	monitor struct{}
)

func Newmonitor() Monitor { return &monitor{} }
