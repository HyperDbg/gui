package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\symbol.h"
var symbolBuf string

type (
	Symbol interface {
		//Fn() (ok bool)
	}
	symbol struct{}
)

func Newsymbol() Symbol { return &symbol{} }
