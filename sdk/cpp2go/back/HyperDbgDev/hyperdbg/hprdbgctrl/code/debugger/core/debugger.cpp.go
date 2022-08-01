package core


type (
Debugger interface{
ShowErrorMessage()(ok bool)//col:370
DebuggerGetNtoskrnlBase()(ok bool)//col:409
DebuggerPauseDebuggee()(ok bool)//col:460
IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool)//col:512
IsTagExist()(ok bool)//col:552
InterpretScript()(ok bool)//col:1039
InterpretConditionsAndCodes()(ok bool)//col:1418
InterpretOutput()(ok bool)//col:1694
SendEventToKernel()(ok bool)//col:1811
RegisterActionToEvent()(ok bool)//col:2014
GetNewDebuggerEventTag()(ok bool)//col:2027
FreeEventsAndActionsMemory()(ok bool)//col:2064
InterpretGeneralEventAndActionsFields()(ok bool)//col:3014
}













































)

func NewDebugger() { return & debugger{} }

func (d *debugger)ShowErrorMessage()(ok bool){//col:370



















































































































































































































































































return true
}














































func (d *debugger)DebuggerGetNtoskrnlBase()(ok bool){//col:409






















return true
}














































func (d *debugger)DebuggerPauseDebuggee()(ok bool){//col:460























return true
}














































func (d *debugger)IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool){//col:512








































return true
}














































func (d *debugger)IsTagExist()(ok bool){//col:552





















return true
}














































func (d *debugger)InterpretScript()(ok bool){//col:1039






















































































































































































































return true
}














































func (d *debugger)InterpretConditionsAndCodes()(ok bool){//col:1418


































































































































































































return true
}














































func (d *debugger)InterpretOutput()(ok bool){//col:1694












































































































































return true
}














































func (d *debugger)SendEventToKernel()(ok bool){//col:1811

















































return true
}














































func (d *debugger)RegisterActionToEvent()(ok bool){//col:2014










































































return true
}














































func (d *debugger)GetNewDebuggerEventTag()(ok bool){//col:2027




return true
}














































func (d *debugger)FreeEventsAndActionsMemory()(ok bool){//col:2064


























return true
}














































func (d *debugger)InterpretGeneralEventAndActionsFields()(ok bool){//col:3014






















































































































































































































































































































































































































































































return true
}
















































