package scriptenginewrapper

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\script-engine-wrapper\\symbol.cpp"
var symbolBuf string

type (
	Symbol interface {
		//Fn() (ok bool)
	}
	symbol struct{}
)

func Newsymbol() Symbol { return &symbol{} }
