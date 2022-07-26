package userlevel

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\user-level\\pe-parser.cpp"
var peParserBuf string

type (
	PeParser interface {
		//Fn() (ok bool)
	}
	peParser struct{}
)

func NewpeParser() PeParser { return &peParser{} }
