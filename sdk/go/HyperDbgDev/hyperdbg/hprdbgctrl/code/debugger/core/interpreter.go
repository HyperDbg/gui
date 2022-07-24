package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\interpreter.cpp"
var interpreterBuf string

type (
	Interpreter interface {
		//Fn() (ok bool)
	}
	interpreter struct{}
)

func Newinterpreter() Interpreter { return &interpreter{} }
