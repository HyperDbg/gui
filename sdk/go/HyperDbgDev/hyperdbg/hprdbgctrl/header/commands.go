package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\commands.h"
var commandsBuf string

type (
	Commands interface {
		//Fn() (ok bool)
	}
	commands struct{}
)

func Newcommands() Commands { return &commands{} }
