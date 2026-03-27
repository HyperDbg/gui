package debugger

type ProcessId = uint32
type ThreadId = uint32
type Address = uint64
type PhysicalAddress = uint64

type ProcessInfo struct {
	ProcessID   uint32 `json:"process_id"`
	ImageName   string `json:"image_name"`
	BaseAddress string `json:"base_address"`
	Size        uint64 `json:"size"`
	CR3         uint64 `json:"cr3"`
}

type ThreadInfo struct {
	ThreadID     uint32 `json:"thread_id"`
	ProcessID    uint32 `json:"process_id"`
	StartAddress string `json:"start_address"`
	TEB          string `json:"teb"`
	State        uint32 `json:"state"`
}

type ModuleInfo struct {
	BaseAddress string `json:"base_address"`
	Size        uint64 `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
}

type BreakpointInfo struct {
	ID        uint64 `json:"id"`
	Address   string `json:"address"`
	Type      int    `json:"type"`
	ProcessID uint32 `json:"process_id"`
	Enabled   bool   `json:"enabled"`
	HitCount  uint64 `json:"hit_count"`
}

type MemoryRegion struct {
	BaseAddress Address `json:"base_address"`
	Size        uint64  `json:"size"`
	Protection  uint32  `json:"protection"`
	State       uint32  `json:"state"`
	Type        uint32  `json:"type_"`
}

type CallStackFrame struct {
	InstructionPointer Address `json:"instruction_pointer"`
	ReturnAddress      Address `json:"return_address"`
	StackPointer       Address `json:"stack_pointer"`
	FramePointer       Address `json:"frame_pointer"`
	ModuleName         string  `json:"module_name"`
	SymbolName         string  `json:"symbol_name,omitempty"`
}

type SymbolInfo struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Size    uint64  `json:"size"`
	Module  string  `json:"module"`
}

type VmxCapabilities struct {
	VMXSupported            bool   `json:"vmx_supported"`
	EPTSupported            bool   `json:"ept_supported"`
	VPIDSupported           bool   `json:"vpid_supported"`
	MSRBitmapSupported      bool   `json:"msr_bitmap_supported"`
	IOBitmapSupported       bool   `json:"io_bitmap_supported"`
	MaxPhysicalAddressWidth uint8  `json:"max_physical_address_width"`
	ProcessorCount          uint32 `json:"processor_count"`
}
