package core
type (
	DebuggerEvents interface {
		DebuggerEventEnableEferOnAllProcessors() (ok bool)             //col:23
		DebuggerEventDisableEferOnAllProcessors() (ok bool)            //col:34
		DebuggerEventEnableMovToCr3ExitingOnAllProcessors() (ok bool)  //col:45
		DebuggerEventDisableMovToCr3ExitingOnAllProcessors() (ok bool) //col:56
		DebuggerEventEptHook2GeneralDetourEventHandler() (ok bool)     //col:124
		DebuggerEventEnableMonitorReadAndWriteForAddress() (ok bool)   //col:157
	}
)
func NewDebuggerEvents() { return &debuggerEvents{} }
func (d *debuggerEvents) DebuggerEventEnableEferOnAllProcessors() (ok bool) { //col:23
	return true
}

func (d *debuggerEvents) DebuggerEventDisableEferOnAllProcessors() (ok bool) { //col:34
	return true
}

func (d *debuggerEvents) DebuggerEventEnableMovToCr3ExitingOnAllProcessors() (ok bool) { //col:45
	return true
}

func (d *debuggerEvents) DebuggerEventDisableMovToCr3ExitingOnAllProcessors() (ok bool) { //col:56
	return true
}

func (d *debuggerEvents) DebuggerEventEptHook2GeneralDetourEventHandler() (ok bool) { //col:124
	return true
}

func (d *debuggerEvents) DebuggerEventEnableMonitorReadAndWriteForAddress() (ok bool) { //col:157
	return true
}

