package globals
type (
	GlobalVariableManagement interface {
		GlobalGuestStateAllocateZeroedMemory() (ok bool) //col:46
		VOID GlobalGuestStateFreeMemory()
(ok bool) //col:56
GlobalEventsAllocateZeroedMemory()(ok bool) //col:78
GlobalEventsFreeMemory()(ok bool)           //col:90
}

)
func NewGlobalVariableManagement() { return &globalVariableManagement{} }
func (g *globalVariableManagement) GlobalGuestStateAllocateZeroedMemory() (ok bool) { //col:46
	return true
}

func (g *globalVariableManagement) VOID GlobalGuestStateFreeMemory()(ok bool) { //col:56
	return true
}

func (g *globalVariableManagement) GlobalEventsAllocateZeroedMemory() (ok bool) { //col:78
	return true
}

func (g *globalVariableManagement) GlobalEventsFreeMemory() (ok bool) { //col:90
	return true
}

