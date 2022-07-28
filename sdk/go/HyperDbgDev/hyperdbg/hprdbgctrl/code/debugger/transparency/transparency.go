package transparency

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\transparency\\transparency.cpp"
var transparencyBuf string

type (
	Transparency interface {
		//Fn() (ok bool)
	}
	transparency struct{}
)

func Newtransparency() Transparency { return &transparency{} }
