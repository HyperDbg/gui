package memory



type (
	MemoryManager interface {
		MemoryManagerReadProcessMemoryNormal() (ok bool) //col:42
	}
	memoryManager struct{}
)

func NewMemoryManager() MemoryManager { return &memoryManager{} }

func (m *memoryManager) MemoryManagerReadProcessMemoryNormal() (ok bool) { //col:42














































	return true
}


