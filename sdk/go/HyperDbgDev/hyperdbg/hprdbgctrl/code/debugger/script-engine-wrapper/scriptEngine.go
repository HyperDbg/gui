package scriptenginewrapper

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\script-engine-wrapper\\script-engine.cpp"
var scriptEngineBuf string

type (
	ScriptEngine interface {
		//Fn() (ok bool)
	}
	scriptEngine struct{}
)

func NewscriptEngine() ScriptEngine { return &scriptEngine{} }
