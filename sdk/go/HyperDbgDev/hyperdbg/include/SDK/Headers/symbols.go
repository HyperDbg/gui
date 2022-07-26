package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Symbols.h"
var symbolsBuf string

type (
	Symbols interface {
		//Fn() (ok bool)
	}
	symbols struct{}
)

func Newsymbols() Symbols { return &symbols{} }

const ()
