package vmx
type (
	Events interface {
		EventInjectInterruption() (ok bool)      //col:37
		EventInjectBreakpoint() (ok bool)        //col:51
		EventInjectGeneralProtection() (ok bool) //col:65
		EventInjectUndefinedOpcode() (ok bool)   //col:81
		EventInjectDebugBreakpoint() (ok bool)   //col:92
		EventInjectPageFault() (ok bool)         //col:123
	}
)
func NewEvents() { return &events{} }
func (e *events) EventInjectInterruption() (ok bool) { //col:37
	return true
}

func (e *events) EventInjectBreakpoint() (ok bool) { //col:51
	return true
}

func (e *events) EventInjectGeneralProtection() (ok bool) { //col:65
	return true
}

func (e *events) EventInjectUndefinedOpcode() (ok bool) { //col:81
	return true
}

func (e *events) EventInjectDebugBreakpoint() (ok bool) { //col:92
	return true
}

func (e *events) EventInjectPageFault() (ok bool) { //col:123
	return true
}

