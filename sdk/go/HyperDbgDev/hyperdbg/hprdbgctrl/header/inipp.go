package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\inipp.h"
var inippBuf string

type (
	Inipp interface {
		//Fn() (ok bool)
	}
	inipp struct{}
)

func Newinipp() Inipp { return &inipp{} }
