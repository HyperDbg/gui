package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\lm.cpp"
var lmBuf string

type (
	Lm interface {
		//Fn() (ok bool)
	}
	lm struct{}
)

func Newlm() Lm { return &lm{} }
