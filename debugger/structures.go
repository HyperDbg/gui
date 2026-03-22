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

type MemoryType int

const (
	MemoryTypePhysical MemoryType = iota
	MemoryTypeVirtual
)

type PoolType int

const (
	PoolTypeThreadInterception PoolType = iota
	PoolTypeMonitor
	PoolTypeEPTHook
	PoolTypeEPTHook2
	PoolTypeRegularEvent
	PoolTypeBigEvent
	PoolTypeRegularSafeBuffer
	PoolTypeBigSafeBuffer
)

type EventMode int

const (
	EventModePre EventMode = iota
	EventModePost
)

type EventModifications struct {
	Enable        *bool
	Disable       *bool
	AddActions    []*DebuggerEventAction
	RemoveActions []uint32
	ModifyActions map[uint32]*DebuggerEventAction
}

type BreakpointType int

const (
	BreakpointTypeSoftware BreakpointType = iota
	BreakpointTypeHardware
	BreakpointTypeEPTHook
)

type PageTableEntries struct {
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
}

type FlushResult struct {
	KernelStatus                               uint32
	CountOfMessagesThatSetAsReadFromVmxRoot    uint32
	CountOfMessagesThatSetAsReadFromVmxNonRoot uint32
}

type ProcessDetails struct {
	ProcessId    uint32
	Name         string
	BaseAddress  uint64
	EntryPoint   uint64
	Is32Bit      bool
	IsDebugged   bool
	KernelStatus uint32
}

type ThreadDetails struct {
	ThreadId     uint32
	ProcessId    uint32
	StartAddress uint64
	IsRunning    bool
	IsSuspended  bool
	KernelStatus uint32
}

type IDTEntry struct {
	Offset     uint64
	Selector   uint16
	Type       uint8
	IST        uint8
	Reserved   uint8
	Attributes uint8
}

type APICInfo struct {
	APICID  uint32
	Version uint32
	EOI     uint32
	TPR     uint32
	IRR     uint32
	ISR     uint32
}

type PCIDevice struct {
	Bus               uint32
	Device            uint32
	Function          uint32
	VendorID          uint16
	DeviceID          uint16
	ClassCode         uint32
	RevisionID        uint8
	HeaderType        uint8
	SubSystemID       uint16
	SubSystemVendorID uint16
}

type PCIDeviceInfo struct {
	Bus           uint32
	Device        uint32
	Function      uint32
	VendorID      uint16
	DeviceID      uint16
	ClassCode     uint32
	RevisionID    uint8
	HeaderType    uint8
	BAR0          uint32
	BAR1          uint32
	BAR2          uint32
	BAR3          uint32
	BAR4          uint32
	BAR5          uint32
	InterruptLine uint8
	InterruptPin  uint8
}

type SMIType int

const (
	SMITypeQuery SMIType = iota
	SMITypeEnable
	SMITypeDisable
	SMITypeConfigure
)

type ExportInfo struct {
	Name        string
	Ordinal     uint16
	Address     uint64
	IsForwarded bool
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelCritical
)

type LogMessage struct {
	Timestamp uint64
	Level     LogLevel
	Message   string
	Source    string
}

type ScriptConfiguration struct {
	ScriptBuffer           []byte
	ScriptLength           uint32
	VariableCount          uint32
	TemporaryVariableCount uint32
}

type ScriptResult struct {
	Success     bool
	Output      string
	Error       string
	ReturnValue uint64
	Variables   map[string]uint64
}

type ReconstructMode int

const (
	ReconstructModeUnknown ReconstructMode = iota
	ReconstructModeUserMode
	ReconstructModeKernelMode
)

type ReconstructType int

const (
	ReconstructTypeUnknown ReconstructType = iota
	ReconstructTypeReconstruct
	ReconstructTypePattern
)

type ReconstructionStatus struct {
	IsInProgress   bool
	Progress       float64
	TotalBytes     uint64
	ProcessedBytes uint64
	Results        []uint64
}

type EPTHookType int

const (
	EPTHookTypeReadAndWriteAndExecute EPTHookType = iota
	EPTHookTypeReadAndWrite
	EPTHookTypeReadAndExecute
	EPTHookTypeWriteAndExecute
	EPTHookTypeRead
	EPTHookTypeWrite
	EPTHookTypeExecute
)

type MonitorType int

const (
	MonitorTypeRead MonitorType = iota
	MonitorTypeWrite
	MonitorTypeExecute
	MonitorTypeReadWrite
	MonitorTypeReadWriteExecute
)

type EPTHook struct {
	ID        uint32
	Address   uint64
	Size      uint32
	HookType  EPTHookType
	ProcessId uint32
	Enabled   bool
	HitCount  uint32
}

type Monitor struct {
	ID          uint32
	Address     uint64
	Size        uint32
	MonitorType MonitorType
	ProcessId   uint32
	Enabled     bool
	HitCount    uint32
}

type CoreInfo struct {
	CoreID             uint32
	CurrentCore        bool
	IsHalted           bool
	InstructionPointer uint64
}

type VersionInfo struct {
	Major   uint32
	Minor   uint32
	Patch   uint32
	Build   uint32
	IsBeta  bool
	IsDebug bool
}

type CommandResult struct {
	Success       bool
	Output        string
	Error         string
	ExitCode      uint32
	ExecutionTime uint64
}

type CommandEvent struct {
	Type      EventType
	Data      any
	Timestamp uint64
}

type TestResults struct {
	Passed  uint32
	Failed  uint32
	Skipped uint32
	Total   uint32
	Details []TestDetail
}

type TestDetail struct {
	Name        string
	Description string
	Success     bool
	Error       string
	Duration    uint64
}

type TestState int

const (
	TestStatePending TestState = iota
	TestStateRunning
	TestStateCompleted
	TestStateFailed
	TestStateCancelled
)

type TestStatus struct {
	TestID   string
	Status   TestState
	Progress float64
	Output   string
	Error    string
}

type Configuration struct {
	DebugMode      DebugMode
	LogLevel       LogLevel
	MaxBufferSize  uint32
	Timeout        uint32
	SerialPort     string
	SerialBaudrate uint32
	SymbolPath     string
	ScriptPath     string
}

type DebugMode int

const (
	DebugModeKernel DebugMode = iota
	DebugModeUser
	DebugModeHardware
)

type NotificationCallback func(event *Event) error

type SymbolInfo struct {
	Name    string
	Address uint64
	Size    uint32
	Type    string
	Module  string
}

type UserModeProcessDetails struct {
	Token                       uint64
	Enabled                     bool
	ActiveThreadId              uint32
	EntrypointOfMainModule      uint64
	BaseAddressOfMainModule     uint64
	Eprocess                    uint64
	ProcessId                   uint32
	Is32Bit                     bool
	IsOnTheStartingPhase        bool
	IsOnThreadInterceptingPhase bool
}

type HWDBGInstanceInformation struct {
	Version                                    uint32
	MaximumNumberOfStages                      uint32
	ScriptVariableLength                       uint32
	NumberOfSupportedLocalAndGlobalVariables   uint32
	NumberOfSupportedTemporaryVariables        uint32
	MaximumNumberOfSupportedGetScriptOperators uint32
	MaximumNumberOfSupportedSetScriptOperators uint32
	SharedMemorySize                           uint32
	DebuggerAreaOffset                         uint32
	DebuggeeAreaOffset                         uint32
	NumberOfPins                               uint32
	NumberOfPorts                              uint32
	ScriptCapabilities                         HWDBGScriptCapabilities
	BramAddrWidth                              uint32
	BramDataWidth                              uint32
}

type HWDBGScriptCapabilities struct {
	AssignLocalGlobalVar  bool
	AssignRegisters       bool
	AssignPseudoRegisters bool
	ConditionalStatements bool
	ComparisonOperators   bool
	StackAssignments      bool
	FuncOr                bool
	FuncXor               bool
	FuncAnd               bool
	FuncAsr               bool
	FuncAsl               bool
	FuncAdd               bool
	FuncSub               bool
	FuncMul               bool
	FuncDiv               bool
	FuncMod               bool
	FuncGt                bool
	FuncLt                bool
	FuncEgt               bool
	FuncElt               bool
	FuncEqual             bool
	FuncNeq               bool
	FuncJmp               bool
	FuncJz                bool
	FuncJnz               bool
	FuncMov               bool
	FuncPrintf            bool
}

type HWDBGScriptBuffer struct {
	ScriptNumberOfSymbols uint32
	ScriptBuffer          []byte
}

type HWDBGPortInformationItems struct {
	PortSize uint32
}

type PCICamInfo struct {
	Bus      uint32
	Device   uint32
	Function uint32
	VendorID uint16
	DeviceID uint16
	Data     uint32
}

type Breakpoint struct {
	Type           BreakpointType
	Address        uint64
	Enabled        bool
	SingleShot     bool
	Active         bool
	Name           string
	Module         string
	Slot           uint16
	HwType         HardwareBreakpointType
	HwSize         HardwareBreakpointSize
	HitCount       uint32
	FastResume     bool
	Silent         bool
	BreakCondition string
	LogText        string
	LogCondition   string
	CommandText    string
	OriginalByte   byte
}

type HardwareBreakpointType int

const (
	HardwareBreakpointAccess HardwareBreakpointType = iota
	HardwareBreakpointWrite
	HardwareBreakpointExecute
)

type HardwareBreakpointSize int

const (
	HardwareBreakpointByte HardwareBreakpointSize = iota
	HardwareBreakpointWord
	HardwareBreakpointDword
	HardwareBreakpointQword
)

type StackFrame struct {
	FrameNumber  uint32
	ReturnAddr   uint64
	FunctionName string
	Params       []uint64
	LocalVars    []uint64
	Module       string
	FileName     string
	LineNumber   uint32
}

type ModuleInfo struct {
	BaseAddress uint64
	Size        uint32
	Name        string
	Path        string
	Is32Bit     bool
	IsSystem    bool
}

type ProcessInfo struct {
	ProcessId   uint32
	Name        string
	ParentId    uint32
	ThreadCount uint32
	BaseAddress uint64
	Size        uint64
}

type ThreadInfo struct {
	ThreadId     uint32
	ProcessId    uint32
	StartAddress uint64
	Priority     int32
	State        string
}
