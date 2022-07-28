package misc

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\misc\\disassembler.cpp"
var disassemblerBuf string

type (
	Disassembler interface {
		//Fn() (ok bool)
	}
	disassembler struct{}
)

func Newdisassembler() Disassembler { return &disassembler{} }
