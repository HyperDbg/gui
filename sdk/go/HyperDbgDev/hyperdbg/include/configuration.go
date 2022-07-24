package include

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\Configuration.h"
var configurationBuf string

type (
	Configuration interface {
		//Fn() (ok bool)
	}
	configuration struct{}
)

func Newconfiguration() Configuration { return &configuration{} }
