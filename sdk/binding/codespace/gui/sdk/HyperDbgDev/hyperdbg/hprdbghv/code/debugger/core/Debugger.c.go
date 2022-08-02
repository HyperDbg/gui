package core



type (
	Debugger interface {
		DebuggerGetRegValueWrapper() (ok bool)       //col:7
		DebuggerSetLastError() (ok bool)             //col:11
		DebuggerInitialize() (ok bool)               //col:74
		DebuggerUninitialize() (ok bool)             //col:433
		DebuggerPerformActions() (ok bool)           //col:530
		DebuggerEnableOrDisableAllEvents() (ok bool) //col:550
		DebuggerTerminateAllEvents() (ok bool)       //col:608
		DebuggerExceptionEventBitmapMask() (ok bool) //col:625
		DebuggerEnableEvent() (ok bool)              //col:636
		DebuggerQueryStateEvent() (ok bool)          //col:646
		DebuggerDisableEvent() (ok bool)             //col:657
		DebuggerIsTagValid() (ok bool)               //col:667
		DebuggerRemoveEventFromEventList() (ok bool) //col:1050
	}
	debugger struct{}
)

func NewDebugger() Debugger { return &debugger{} }

func (d *debugger) DebuggerGetRegValueWrapper() (ok bool) { //col:7












	return true
}


func (d *debugger) DebuggerSetLastError() (ok bool) { //col:11







	return true
}


func (d *debugger) DebuggerInitialize() (ok bool) { //col:74


































































	return true
}


func (d *debugger) DebuggerUninitialize() (ok bool) { //col:433



















































































































































































































































































































































































	return true
}


func (d *debugger) DebuggerPerformActions() (ok bool) { //col:530













































































































	return true
}


func (d *debugger) DebuggerEnableOrDisableAllEvents() (ok bool) { //col:550























	return true
}


func (d *debugger) DebuggerTerminateAllEvents() (ok bool) { //col:608



































































	return true
}


func (d *debugger) DebuggerExceptionEventBitmapMask() (ok bool) { //col:625




















	return true
}


func (d *debugger) DebuggerEnableEvent() (ok bool) { //col:636














	return true
}


func (d *debugger) DebuggerQueryStateEvent() (ok bool) { //col:646













	return true
}


func (d *debugger) DebuggerDisableEvent() (ok bool) { //col:657














	return true
}


func (d *debugger) DebuggerIsTagValid() (ok bool) { //col:667













	return true
}


func (d *debugger) DebuggerRemoveEventFromEventList() (ok bool) { //col:1050














































































































































































































































































































































































































	return true
}


