package core


type (
DebuggerEvents interface{
DebuggerEventEnableEferOnAllProcessors()(ok bool)//col:23
DebuggerEventDisableEferOnAllProcessors()(ok bool)//col:35
DebuggerEventEnableMovToCr3ExitingOnAllProcessors()(ok bool)//col:47
DebuggerEventDisableMovToCr3ExitingOnAllProcessors()(ok bool)//col:59
DebuggerEventEptHook2GeneralDetourEventHandler()(ok bool)//col:128
DebuggerEventEnableMonitorReadAndWriteForAddress()(ok bool)//col:162
}

)

func NewDebuggerEvents() { return & debuggerEvents{} }

func (d *debuggerEvents)DebuggerEventEnableEferOnAllProcessors()(ok bool){//col:23




return true
}


func (d *debuggerEvents)DebuggerEventDisableEferOnAllProcessors()(ok bool){//col:35




return true
}


func (d *debuggerEvents)DebuggerEventEnableMovToCr3ExitingOnAllProcessors()(ok bool){//col:47




return true
}


func (d *debuggerEvents)DebuggerEventDisableMovToCr3ExitingOnAllProcessors()(ok bool){//col:59




return true
}


func (d *debuggerEvents)DebuggerEventEptHook2GeneralDetourEventHandler()(ok bool){//col:128




















return true
}


func (d *debuggerEvents)DebuggerEventEnableMonitorReadAndWriteForAddress()(ok bool){//col:162












return true
}




