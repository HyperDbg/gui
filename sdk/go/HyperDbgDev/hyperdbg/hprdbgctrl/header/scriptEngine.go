package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\script-engine.h"
var scriptEngineBuf string

type (
	ScriptEngine interface {
		//Fn() (ok bool)
	}
	scriptEngine struct{}
)

func NewscriptEngine() ScriptEngine { return &scriptEngine{} }
