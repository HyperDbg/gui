package scriptenginewrapper

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\script-engine-wrapper\\script-engine-wrapper.cpp"
var scriptEngineWrapperBuf string

type (
	ScriptEngineWrapper interface {
		//Fn() (ok bool)
	}
	scriptEngineWrapper struct{}
)

func NewscriptEngineWrapper() ScriptEngineWrapper { return &scriptEngineWrapper{} }
