package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\pe-parser.h"
var peParserBuf string

type (
	PeParser interface {
		//Fn() (ok bool)
	}
	peParser struct{}
)

func NewpeParser() PeParser { return &peParser{} }
