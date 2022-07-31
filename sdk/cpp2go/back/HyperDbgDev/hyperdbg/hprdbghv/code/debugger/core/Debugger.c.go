package core
type (
	Debugger interface {
		GetRegValue() (ok bool)                                 //col:24
		DebuggerGetLastError() (ok bool)                        //col:35
		DebuggerSetLastError() (ok bool)                        //col:47
		DebuggerInitialize() (ok bool)                          //col:213
		DebuggerUninitialize() (ok bool)                        //col:279
		DebuggerCreateEvent() (ok bool)                         //col:379
		DebuggerAddActionToEvent() (ok bool)                    //col:615
		DebuggerRegisterEvent() (ok bool)                       //col:698
		DebuggerTriggerEvents() (ok bool)                       //col:1133
		DebuggerPerformActions() (ok bool)                      //col:1179
		DebuggerPerformRunScript() (ok bool)                    //col:1277
		DebuggerPerformRunTheCustomCode() (ok bool)             //col:1329
		DebuggerPerformBreakToDebugger() (ok bool)              //col:1369
		DebuggerGetEventByTag() (ok bool)                       //col:1410
		DebuggerEnableOrDisableAllEvents() (ok bool)            //col:1456
		DebuggerTerminateAllEvents() (ok bool)                  //col:1501
		DebuggerRemoveAllEvents() (ok bool)                     //col:1550
		DebuggerEventListCount() (ok bool)                      //col:1581
		DebuggerEventListCountByCore() (ok bool)                //col:1619
		DebuggerExceptionEventBitmapMask() (ok bool)            //col:1653
		DebuggerEnableEvent() (ok bool)                         //col:1685
		DebuggerQueryStateEvent() (ok bool)                     //col:1714
		DebuggerDisableEvent() (ok bool)                        //col:1747
		DebuggerIsTagValid() (ok bool)                          //col:1774
		DebuggerRemoveEventFromEventList() (ok bool)            //col:1824
		DebuggerRemoveAllActionsFromEvent() (ok bool)           //col:1876
		DebuggerRemoveEvent() (ok bool)                         //col:1938
		DebuggerParseEventFromUsermode() (ok bool)              //col:2739
		DebuggerParseActionFromUsermode() (ok bool)             //col:2879
		DebuggerTerminateEvent() (ok bool)                      //col:3073
		DebuggerParseEventsModificationFromUsermode() (ok bool) //col:3247
	}
)
func NewDebugger() { return &debugger{} }
func (d *debugger) return GetRegValue()(ok bool) { //col:24
	return true
}

func (d *debugger) DebuggerGetLastError() (ok bool) { //col:35
	return true
}

func (d *debugger) DebuggerSetLastError() (ok bool) { //col:47
	return true
}

func (d *debugger) DebuggerInitialize() (ok bool) { //col:213
	return true
}

func (d *debugger) DebuggerUninitialize() (ok bool) { //col:279
	return true
}

func (d *debugger) DebuggerCreateEvent() (ok bool) { //col:379
	return true
}

func (d *debugger) DebuggerAddActionToEvent() (ok bool) { //col:615
	return true
}

func (d *debugger) DebuggerRegisterEvent() (ok bool) { //col:698
	return true
}

func (d *debugger) DebuggerTriggerEvents() (ok bool) { //col:1133
	return true
}

func (d *debugger) DebuggerPerformActions() (ok bool) { //col:1179
	return true
}

func (d *debugger) DebuggerPerformRunScript() (ok bool) { //col:1277
	return true
}

func (d *debugger) DebuggerPerformRunTheCustomCode() (ok bool) { //col:1329
	return true
}

func (d *debugger) DebuggerPerformBreakToDebugger() (ok bool) { //col:1369
	return true
}

func (d *debugger) DebuggerGetEventByTag() (ok bool) { //col:1410
	return true
}

func (d *debugger) DebuggerEnableOrDisableAllEvents() (ok bool) { //col:1456
	return true
}

func (d *debugger) DebuggerTerminateAllEvents() (ok bool) { //col:1501
	return true
}

func (d *debugger) DebuggerRemoveAllEvents() (ok bool) { //col:1550
	return true
}

func (d *debugger) DebuggerEventListCount() (ok bool) { //col:1581
	return true
}

func (d *debugger) DebuggerEventListCountByCore() (ok bool) { //col:1619
	return true
}

func (d *debugger) DebuggerExceptionEventBitmapMask() (ok bool) { //col:1653
	return true
}

func (d *debugger) DebuggerEnableEvent() (ok bool) { //col:1685
	return true
}

func (d *debugger) DebuggerQueryStateEvent() (ok bool) { //col:1714
	return true
}

func (d *debugger) DebuggerDisableEvent() (ok bool) { //col:1747
	return true
}

func (d *debugger) DebuggerIsTagValid() (ok bool) { //col:1774
	return true
}

func (d *debugger) DebuggerRemoveEventFromEventList() (ok bool) { //col:1824
	return true
}

func (d *debugger) DebuggerRemoveAllActionsFromEvent() (ok bool) { //col:1876
	return true
}

func (d *debugger) DebuggerRemoveEvent() (ok bool) { //col:1938
	return true
}

func (d *debugger) DebuggerParseEventFromUsermode() (ok bool) { //col:2739
	return true
}

func (d *debugger) DebuggerParseActionFromUsermode() (ok bool) { //col:2879
	return true
}

func (d *debugger) DebuggerTerminateEvent() (ok bool) { //col:3073
	return true
}

func (d *debugger) DebuggerParseEventsModificationFromUsermode() (ok bool) { //col:3247
	return true
}

