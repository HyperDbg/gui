package memory
type (
	MemoryManager interface {
		MemoryManagerReadProcessMemoryNormal() (ok bool) //col:129
	}
)
func NewMemoryManager() { return &memoryManager{} }
func (m *memoryManager) MemoryManagerReadProcessMemoryNormal() (ok bool) { //col:129
	return true
}

