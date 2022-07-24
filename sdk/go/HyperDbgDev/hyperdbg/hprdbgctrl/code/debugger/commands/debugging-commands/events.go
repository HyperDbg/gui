package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\events.cpp"
var eventsBuf string

type (
	Events interface {
		//Fn() (ok bool)
	}
	events struct{}
)

func Newevents() Events { return &events{} }
