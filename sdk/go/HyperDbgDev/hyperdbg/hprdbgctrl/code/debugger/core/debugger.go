package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\debugger.cpp"
var debuggerBuf string

type (
	Debugger interface {
		ShowErrorMessage() (ok bool)
		DebuggerGetNtoskrnlBase() (ok bool)
		DebuggerPauseDebuggee() (ok bool)
		IsConnectedToAnyInstanceOfDebuggerOrDebuggee() (ok bool)
		IsTagExist() (ok bool)
		InterpretScript() (ok bool)
		InterpretConditionsAndCodes() (ok bool)
		InterpretOutput() (ok bool)
		SendEventToKernel() (ok bool)
		RegisterActionToEvent() (ok bool)
		GetNewDebuggerEventTag() (ok bool)
		FreeEventsAndActionsMemory() (ok bool)
		InterpretGeneralEventAndActionsFields() (ok bool)
	}
	debugger struct{}
)

func (d *debugger) ShowErrorMessage() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) DebuggerGetNtoskrnlBase() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) DebuggerPauseDebuggee() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) IsConnectedToAnyInstanceOfDebuggerOrDebuggee() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) IsTagExist() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) InterpretScript() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) InterpretConditionsAndCodes() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) InterpretOutput() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) SendEventToKernel() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) RegisterActionToEvent() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) GetNewDebuggerEventTag() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) FreeEventsAndActionsMemory() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (d *debugger) InterpretGeneralEventAndActionsFields() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func Newdebugger() Debugger { return &debugger{} }
