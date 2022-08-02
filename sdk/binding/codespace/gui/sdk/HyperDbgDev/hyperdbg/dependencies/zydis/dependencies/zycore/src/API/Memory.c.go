package API

type (
	Memory interface {
		ZyanU32_ZyanMemoryGetSystemPageSize() (ok bool) //col:29
		ZyanStatus_ZyanMemoryVirtualFree() (ok bool)    //col:45
	}
	memory struct{}
)

func NewMemory() Memory { return &memory{} }

func (m *memory) ZyanU32_ZyanMemoryGetSystemPageSize() (ok bool) { //col:29

	return true
}

func (m *memory) ZyanStatus_ZyanMemoryVirtualFree() (ok bool) { //col:45

	return true
}
