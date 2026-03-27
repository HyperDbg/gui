package hyperdbgrust

type MemoryAccessType uint32

const (
	MemoryAccessTypeRead MemoryAccessType = iota
	MemoryAccessTypeWrite
	MemoryAccessTypeExecute
)

type MemoryAccessEvent struct {
	Header          EventHeader      `json:"header"`
	VirtualAddress  Address          `json:"virtual_address"`
	PhysicalAddress PhysicalAddress  `json:"physical_address"`
	AccessType      MemoryAccessType `json:"access_type"`
	Size            uint32           `json:"size"`
	Value           uint64           `json:"value"`
}
