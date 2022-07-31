package API
type (
Memory interface{
  Zyan Core Library ()(ok bool)//col:59
ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool)//col:75
ZyanStatus ZyanMemoryVirtualProtect()(ok bool)//col:102
ZyanStatus ZyanMemoryVirtualFree()(ok bool)//col:124
}

)
func NewMemory() { return & memory{} }
func (m *memory)  Zyan Core Library ()(ok bool){//col:59
return true
}

func (m *memory)ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool){//col:75
return true
}

func (m *memory)ZyanStatus ZyanMemoryVirtualProtect()(ok bool){//col:102
return true
}

func (m *memory)ZyanStatus ZyanMemoryVirtualFree()(ok bool){//col:124
return true
}

