package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\transparency.h"
var transparencyBuf string

type (
	Transparency interface {
		//Fn() (ok bool)
	}
	transparency struct{}
)

func Newtransparency() Transparency { return &transparency{} }
