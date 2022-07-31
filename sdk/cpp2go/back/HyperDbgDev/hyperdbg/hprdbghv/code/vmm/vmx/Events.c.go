package vmx


type (
Events interface{
EventInjectInterruption()(ok bool)//col:37
EventInjectBreakpoint()(ok bool)//col:52
EventInjectGeneralProtection()(ok bool)//col:67
EventInjectUndefinedOpcode()(ok bool)//col:84
EventInjectDebugBreakpoint()(ok bool)//col:96
EventInjectPageFault()(ok bool)//col:128
}

)

func NewEvents() { return & events{} }

func (e *events)EventInjectInterruption()(ok bool){//col:37













return true
}


func (e *events)EventInjectBreakpoint()(ok bool){//col:52







return true
}


func (e *events)EventInjectGeneralProtection()(ok bool){//col:67







return true
}


func (e *events)EventInjectUndefinedOpcode()(ok bool){//col:84





return true
}


func (e *events)EventInjectDebugBreakpoint()(ok bool){//col:96




return true
}


func (e *events)EventInjectPageFault()(ok bool){//col:128











return true
}




