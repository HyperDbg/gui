package include

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\Definition.h"
var definitionBuf string

type (
	Definition interface {
		//Fn() (ok bool)
	}
	definition struct{}
)

func Newdefinition() Definition { return &definition{} }
