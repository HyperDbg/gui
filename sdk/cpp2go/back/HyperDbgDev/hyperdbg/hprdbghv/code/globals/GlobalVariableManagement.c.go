package globals


type (
GlobalVariableManagement interface{
GlobalGuestStateAllocateZeroedMemory()(ok bool)//col:46
VOID GlobalGuestStateFreeMemory()(ok bool)//col:57
GlobalEventsAllocateZeroedMemory()(ok bool)//col:80
GlobalEventsFreeMemory()(ok bool)//col:93
}
















)

func NewGlobalVariableManagement() { return & globalVariableManagement{} }

func (g *globalVariableManagement)GlobalGuestStateAllocateZeroedMemory()(ok bool){//col:46













return true
}

















func (g *globalVariableManagement)VOID GlobalGuestStateFreeMemory()(ok bool){//col:57




return true
}

















func (g *globalVariableManagement)GlobalEventsAllocateZeroedMemory()(ok bool){//col:80












return true
}

















func (g *globalVariableManagement)GlobalEventsFreeMemory()(ok bool){//col:93









return true
}



















