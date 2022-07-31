package commands
type (
	BreakpointCommands interface {
		BreakpointCheckAndHandleEptHookBreakpoints() (ok bool)         //col:122
		BreakpointCheckAndHandleDebuggerDefinedBreakpoints() (ok bool) //col:285
		BreakpointHandleBpTraps() (ok bool)                            //col:372
		BreakpointWrite() (ok bool)                                    //col:416
		BreakpointClear() (ok bool)                                    //col:460
		BreakpointRemoveAllBreakpoints() (ok bool)                     //col:497
		BreakpointGetEntryByBreakpointId() (ok bool)                   //col:528
		BreakpointGetEntryByAddress() (ok bool)                        //col:559
		BreakpointAddNew() (ok bool)                                   //col:683
		BreakpointListAllBreakpoint() (ok bool)                        //col:734
		BreakpointListOrModify() (ok bool)                             //col:840
	}
)
func NewBreakpointCommands() { return &breakpointCommands{} }
func (b *breakpointCommands) BreakpointCheckAndHandleEptHookBreakpoints() (ok bool) { //col:122
	return true
}

func (b *breakpointCommands) BreakpointCheckAndHandleDebuggerDefinedBreakpoints() (ok bool) { //col:285
	return true
}

func (b *breakpointCommands) BreakpointHandleBpTraps() (ok bool) { //col:372
	return true
}

func (b *breakpointCommands) BreakpointWrite() (ok bool) { //col:416
	return true
}

func (b *breakpointCommands) BreakpointClear() (ok bool) { //col:460
	return true
}

func (b *breakpointCommands) BreakpointRemoveAllBreakpoints() (ok bool) { //col:497
	return true
}

func (b *breakpointCommands) BreakpointGetEntryByBreakpointId() (ok bool) { //col:528
	return true
}

func (b *breakpointCommands) BreakpointGetEntryByAddress() (ok bool) { //col:559
	return true
}

func (b *breakpointCommands) BreakpointAddNew() (ok bool) { //col:683
	return true
}

func (b *breakpointCommands) BreakpointListAllBreakpoint() (ok bool) { //col:734
	return true
}

func (b *breakpointCommands) BreakpointListOrModify() (ok bool) { //col:840
	return true
}

