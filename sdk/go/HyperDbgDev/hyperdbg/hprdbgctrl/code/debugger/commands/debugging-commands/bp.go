package bp

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\bp.cpp"
var bpBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)
func New() Interface { return &object{} }


