package debugger

type DebuggerReadMemoryAddressMode uint32

const (
	DebuggerReadMemoryAddressModePhysical DebuggerReadMemoryAddressMode = 0
	DebuggerReadMemoryAddressModeVirtual  DebuggerReadMemoryAddressMode = 1
	DebuggerReadMemoryAddressModeVirtual2 DebuggerReadMemoryAddressMode = 2
)

type DebuggerReadMemoryType uint32

const (
	DEBUGGER_READ_MEMORY_TYPE_UNKNOWN          DebuggerReadMemoryType = 0
	DEBUGGER_READ_MEMORY_TYPE_READ_FROM_MEMORY DebuggerReadMemoryType = 1
	DEBUGGER_READ_MEMORY_TYPE_CACHE            DebuggerReadMemoryType = 2
	DEBUGGER_READ_MEMORY_TYPE_TLB              DebuggerReadMemoryType = 3
)

type DebuggerReadReadingType uint32

const (
	DebuggerReadReadingTypeStandard DebuggerReadReadingType = 0
	DebuggerReadReadingTypePte      DebuggerReadReadingType = 1
	DebuggerReadReadingTypePdpte    DebuggerReadReadingType = 2
	DebuggerReadReadingTypePdpte2   DebuggerReadReadingType = 3
	DebuggerReadReadingTypePml4e    DebuggerReadReadingType = 4
)

type DebuggerReadMemory struct {
	Pid            uint32
	Address        uint64
	Size           uint32
	GetAddressMode bool
	AddressMode    DebuggerReadMemoryAddressMode
	MemoryType     DebuggerReadMemoryType
	ReadingType    DebuggerReadReadingType
	ReturnLength   uint32
	KernelStatus   uint32
}

type DebuggerReadPageTableEntriesDetails struct {
	VirtualAddress      uint64
	ProcessId           uint32
	Pml4eVirtualAddress uint64
	Pml4eValue          uint64
	PdpteVirtualAddress uint64
	PdpteValue          uint64
	PdeVirtualAddress   uint64
	PdeValue            uint64
	PteVirtualAddress   uint64
	PteValue            uint64
	KernelStatus        uint32
}

type DebuggerVa2paAndPa2vaCommands struct {
	VirtualAddress     uint64
	PhysicalAddress    uint64
	ProcessId          uint32
	IsVirtual2Physical bool
	KernelStatus       uint32
}

type DebuggerGeneralEventDetails struct {
	Tag       uint64
	ProcessId uint32
	IsEnabled bool
	Type      uint32
}

type DebuggerEventAction struct {
	ActionType                uint32
	ActionTag                 uint64
	SendTheResultsImmediately bool
	OptionalParam1            uint64
	OptionalParam2            uint64
	OptionalParam3            uint64
	OptionalParam4            uint64
}

type DebuggerModifyEventRequest struct {
	Tag       uint64
	IsEnabled bool
	Type      uint32
	ProcessId uint32
}

type DebuggerEventAndActionResult struct {
	IsSuccessful bool
	Error        uint32
}

type DebuggerWriteMemory struct {
	Pid          uint32
	Address      uint64
	Buffer       []byte
	BufferSize   uint32
	KernelStatus uint32
}

type EventAndActionResult struct {
	IsSuccessful bool
	Error        uint32
}

type DebuggeeKdPausedPacket struct {
	Rip                    uint64
	IsProcessorOn32BitMode bool
	IgnoreDisassembling    bool
	PausingReason          uint32
	CurrentCore            uint32
	EventTag               uint64
	EventCallingStage      uint32
	Rflags                 uint64
	InstructionBytesOnRip  [15]byte
	ReadInstructionLen     uint16
}

type DebuggeeDetailsAndSwitchProcessPacket struct {
	ActionType            uint32
	ProcessId             uint32
	Process               uint64
	IsSwitchByClkIntr     bool
	ProcessName           [16]byte
	ProcessListSymDetails uint64
	Result                uint32
}

type DebuggeeDetailsAndSwitchThreadPacket struct {
	ActionType  uint32
	ThreadId    uint32
	Thread      uint64
	ProcessId   uint32
	Process     uint64
	ProcessName [16]byte
	Result      uint32
}

type DebuggeeRegisterReadDescription struct {
	RegisterId   uint32
	Value        uint64
	KernelStatus uint32
}
