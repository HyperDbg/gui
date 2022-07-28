package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\interpreter.cpp"
var interpreterBuf string

type (
	Interpreter interface {
		HyperDbgInterpreter() (ok bool)
		InterpreterRemoveComments() (ok bool)
		HyperDbgShowSignature() (ok bool)
		HyperDbgCheckMultilineCommand() (ok bool)
		HyperDbgContinuePreviousCommand() (ok bool)
		GetCommandAttributes() (ok bool)
		InitializeDebugger() (ok bool)
		InitializeCommandsDictionary() (ok bool) //todo change to map? all command  in here
	}
	interpreter struct{}
)

func (i *interpreter) HyperDbgInterpreter() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InterpreterRemoveComments() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgShowSignature() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgCheckMultilineCommand() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgContinuePreviousCommand() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) GetCommandAttributes() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InitializeDebugger() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InitializeCommandsDictionary() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func Newinterpreter() Interpreter { return &interpreter{} }
