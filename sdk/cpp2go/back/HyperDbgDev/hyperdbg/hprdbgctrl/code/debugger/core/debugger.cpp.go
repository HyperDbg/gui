package core
type (
Debugger interface{
ShowErrorMessage()(ok bool)//col:370
DebuggerGetNtoskrnlBase()(ok bool)//col:408
 * then when it returns true the debuggee is not paused anymore ()(ok bool)//col:458
IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool)//col:509
IsTagExist()(ok bool)//col:548
 * successful ()(ok bool)//col:1034
 * successful ()(ok bool)//col:1412
 * input that needs to be considered for this event ()(ok bool)//col:1687
SendEventToKernel()(ok bool)//col:1803
RegisterActionToEvent()(ok bool)//col:2005
GetNewDebuggerEventTag()(ok bool)//col:2017
FreeEventsAndActionsMemory()(ok bool)//col:2053
InterpretGeneralEventAndActionsFields()(ok bool)//col:3002
}

)
func NewDebugger() { return & debugger{} }
func (d *debugger)ShowErrorMessage()(ok bool){//col:370
return true
}

func (d *debugger)DebuggerGetNtoskrnlBase()(ok bool){//col:408
return true
}

func (d *debugger) * then when it returns true the debuggee is not paused anymore ()(ok bool){//col:458
return true
}

func (d *debugger)IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool){//col:509
return true
}

func (d *debugger)IsTagExist()(ok bool){//col:548
return true
}

func (d *debugger) * successful ()(ok bool){//col:1034
return true
}

func (d *debugger) * successful ()(ok bool){//col:1412
return true
}

func (d *debugger) * input that needs to be considered for this event ()(ok bool){//col:1687
return true
}

func (d *debugger)SendEventToKernel()(ok bool){//col:1803
return true
}

func (d *debugger)RegisterActionToEvent()(ok bool){//col:2005
return true
}

func (d *debugger)GetNewDebuggerEventTag()(ok bool){//col:2017
return true
}

func (d *debugger)FreeEventsAndActionsMemory()(ok bool){//col:2053
return true
}

func (d *debugger)InterpretGeneralEventAndActionsFields()(ok bool){//col:3002
return true
}

