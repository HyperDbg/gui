package commands


type (
BreakpointCommands interface{
BreakpointCheckAndHandleEptHookBreakpoints()(ok bool)//col:122
BreakpointCheckAndHandleDebuggerDefinedBreakpoints()(ok bool)//col:286
BreakpointHandleBpTraps()(ok bool)//col:374
BreakpointWrite()(ok bool)//col:419
BreakpointClear()(ok bool)//col:464
BreakpointRemoveAllBreakpoints()(ok bool)//col:502
BreakpointGetEntryByBreakpointId()(ok bool)//col:534
BreakpointGetEntryByAddress()(ok bool)//col:566
BreakpointAddNew()(ok bool)//col:691
BreakpointListAllBreakpoint()(ok bool)//col:743
BreakpointListOrModify()(ok bool)//col:850
}

)

func NewBreakpointCommands() { return & breakpointCommands{} }

func (b *breakpointCommands)BreakpointCheckAndHandleEptHookBreakpoints()(ok bool){//col:122































return true
}


func (b *breakpointCommands)BreakpointCheckAndHandleDebuggerDefinedBreakpoints()(ok bool){//col:286
































































return true
}


func (b *breakpointCommands)BreakpointHandleBpTraps()(ok bool){//col:374


































return true
}


func (b *breakpointCommands)BreakpointWrite()(ok bool){//col:419
















return true
}


func (b *breakpointCommands)BreakpointClear()(ok bool){//col:464


















return true
}


func (b *breakpointCommands)BreakpointRemoveAllBreakpoints()(ok bool){//col:502













return true
}


func (b *breakpointCommands)BreakpointGetEntryByBreakpointId()(ok bool){//col:534















return true
}


func (b *breakpointCommands)BreakpointGetEntryByAddress()(ok bool){//col:566















return true
}


func (b *breakpointCommands)BreakpointAddNew()(ok bool){//col:691












































return true
}


func (b *breakpointCommands)BreakpointListAllBreakpoint()(ok bool){//col:743


































return true
}


func (b *breakpointCommands)BreakpointListOrModify()(ok bool){//col:850




















































return true
}




