package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Constants.h"
var constantsBuf string

type (
	Constants interface {
		//Fn() (ok bool)
	}
	constants struct{}
)

func Newconstants() Constants { return &constants{} }
