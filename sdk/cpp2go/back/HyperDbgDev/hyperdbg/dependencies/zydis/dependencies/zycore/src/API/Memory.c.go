package API


type (
Memory interface{
#if   defined()(ok bool)//col:59
ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool)//col:76
ZyanStatus ZyanMemoryVirtualProtect()(ok bool)//col:104
ZyanStatus ZyanMemoryVirtualFree()(ok bool)//col:127
}

















)

func NewMemory() { return & memory{} }

func (m *memory)#if   defined()(ok bool){//col:59
















return true
}


















func (m *memory)ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool){//col:76










return true
}


















func (m *memory)ZyanStatus ZyanMemoryVirtualProtect()(ok bool){//col:104

















return true
}


















func (m *memory)ZyanStatus ZyanMemoryVirtualFree()(ok bool){//col:127
















return true
}




















