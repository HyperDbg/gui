package debugger

import (
	"encoding/binary"
	"errors"
)

type Validator interface {
	Validate() error
}

type SizeVerifier interface {
	ExpectedSize() uintptr
	ExpectedSerSize() uintptr
}

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

type DebuggerMsrActionType uint32

const (
	DebuggerMsrActionRead  DebuggerMsrActionType = 0
	DebuggerMsrActionWrite DebuggerMsrActionType = 1
)

type DebuggerReadOrWriteMsr struct {
	Msr        uint32
	_          [4]byte
	CoreNumber uint32
	ActionType DebuggerMsrActionType
	_          [4]byte
	Value      uint64
}

type DebuggerEditMemoryType uint32

const (
	DebuggerEditMemoryVirtual  DebuggerEditMemoryType = 0
	DebuggerEditMemoryPhysical DebuggerEditMemoryType = 1
)

type DebuggerEditMemoryByteSize uint32

const (
	DebuggerEditMemoryByte  DebuggerEditMemoryByteSize = 0
	DebuggerEditMemoryDword DebuggerEditMemoryByteSize = 1
	DebuggerEditMemoryQword DebuggerEditMemoryByteSize = 2
)

type DebuggerEditMemory struct {
	Result             uint32
	_                  [4]byte
	Address            uint64
	ProcessId          uint32
	MemoryType         DebuggerEditMemoryType
	ByteSize           DebuggerEditMemoryByteSize
	CountOf64Chunks    uint32
	FinalStructureSize uint32
}

type DebuggerSearchMemoryType uint32

const (
	DebuggerSearchMemoryPhysical            DebuggerSearchMemoryType = 0
	DebuggerSearchMemoryVirtual             DebuggerSearchMemoryType = 1
	DebuggerSearchMemoryPhysicalFromVirtual DebuggerSearchMemoryType = 2
)

type DebuggerSearchMemoryByteSize uint32

const (
	DebuggerSearchMemoryByte  DebuggerSearchMemoryByteSize = 0
	DebuggerSearchMemoryDword DebuggerSearchMemoryByteSize = 1
	DebuggerSearchMemoryQword DebuggerSearchMemoryByteSize = 2
)

type DebuggerSearchMemory struct {
	Result             uint32
	_                  [4]byte
	Address            uint64
	ProcessId          uint32
	MemoryType         DebuggerSearchMemoryType
	ByteSize           DebuggerSearchMemoryByteSize
	CountOf64Chunks    uint32
	FinalStructureSize uint32
}

type DebuggerPageInRequest struct {
	VirtualAddressFrom uint64
	VirtualAddressTo   uint64
	ProcessId          uint32
	PageFaultErrorCode uint32
	KernelStatus       uint32
	_                  [4]byte
}

type DebuggerPreallocCommandType uint32

const (
	DebuggerPreallocCommandTypeThreadInterception DebuggerPreallocCommandType = 0
	DebuggerPreallocCommandTypeMonitor            DebuggerPreallocCommandType = 1
	DebuggerPreallocCommandTypeEPTHook            DebuggerPreallocCommandType = 2
	DebuggerPreallocCommandTypeEPTHook2           DebuggerPreallocCommandType = 3
	DebuggerPreallocCommandTypeRegularEvent       DebuggerPreallocCommandType = 4
	DebuggerPreallocCommandTypeBigEvent           DebuggerPreallocCommandType = 5
	DebuggerPreallocCommandTypeRegularSafeBuffer  DebuggerPreallocCommandType = 6
	DebuggerPreallocCommandTypeBigSafeBuffer      DebuggerPreallocCommandType = 7
)

type DebuggerPreallocCommand struct {
	Type         DebuggerPreallocCommandType
	Count        uint32
	KernelStatus uint32
}

type DebuggerHideDebuggerRequest struct {
	IsHide       bool
	IsEnabled    bool
	_            [6]byte
	KernelStatus uint32
	_            [4]byte
}

type DebuggerPteRequest struct {
	VirtualAddress uint64
	ProcessId      uint32
	_              [4]byte
}

type DebuggerVa2paRequest struct {
	VirtualAddress     uint64
	ProcessId          uint32
	IsVirtual2Physical bool
	_                  [3]byte
}

type DebuggerPa2vaRequest struct {
	PhysicalAddress    uint64
	ProcessId          uint32
	IsVirtual2Physical bool
	_                  [3]byte
}

type DebuggerReconstructMemoryRequest struct {
	ProcessId    uint32
	Size         uint32
	Mode         ReconstructMode
	Type         ReconstructType
	KernelStatus uint32
}

type DebuggerEventRegisterRequest struct {
	EventType      uint32
	Tag            uint64
	ProcessId      uint32
	CoreId         uint32
	IsEnabled      bool
	EventStage     uint32
	OptionalParam1 uint64
	OptionalParam2 uint64
	OptionalParam3 uint64
	OptionalParam4 uint64
}

type DebuggerEpthookRequest struct {
	Address   uint64
	Size      uint32
	HookType  EPTHookType
	ProcessId uint32
}

type DebuggerMonitorRequest struct {
	Address     uint64
	Size        uint32
	MonitorType MonitorType
	ProcessId   uint32
}

type DebuggerSyscallHookRequest struct {
	SyscallNumber uint32
	HookType      uint32
}

type DebuggerExceptionHookRequest struct {
	ExceptionType uint32
}

type DebuggerInterruptHookRequest struct {
	Vector uint32
}

type DebuggerIoHookRequest struct {
	Port     uint16
	_        [2]byte
	HookType uint32
}

type DebuggerApicRequest struct {
	ApicID uint32
}

type DebuggerPciCamRequest struct {
	Bus      uint32
	Device   uint32
	Function uint32
}

type DebuggerPmcRequest struct {
	PmcNumber uint32
}

type DebuggerTrackMemoryRequest struct {
	Address uint64
	Size    uint32
}

type DebuggerMeasureRequest struct {
	Address uint64
}

type DebuggerInstructionTraceRequest struct {
	Address uint64
}

type DebuggerSwitchThreadRequest struct {
	ThreadId  uint32
	ProcessId uint32
	Action    uint32
}

type DebuggerQueryProcessThreadRequest struct {
	QueryType   uint32
	QueryAction uint32
	ProcessId   uint32
	ThreadId    uint32
	Count       uint32
}

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

type DebuggeeProcessListNeededDetails struct {
	PsActiveProcessHead      uint64
	ImageFileNameOffset      uint32
	UniquePidOffset          uint32
	ActiveProcessLinksOffset uint32
}

const (
	DebuggeeDetailsAndSwitchProcessGetProcessDetails uint32 = iota
	DebuggeeDetailsAndSwitchProcessGetProcessList
	DebuggeeDetailsAndSwitchProcessPerformSwitch
)

type DebuggeeDetailsAndSwitchProcessPacket struct {
	ActionType            uint32
	ProcessId             uint32
	Process               uint64
	IsSwitchByClkIntr     bool
	ProcessName           [16]byte
	ProcessListSymDetails DebuggeeProcessListNeededDetails
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

type DebuggerAttachDetachUserModeProcessActionType uint32

const (
	DebuggerAttachDetachUserModeProcessActionAttach                             DebuggerAttachDetachUserModeProcessActionType = 0
	DebuggerAttachDetachUserModeProcessActionDetach                             DebuggerAttachDetachUserModeProcessActionType = 1
	DebuggerAttachDetachUserModeProcessActionRemoveHooks                        DebuggerAttachDetachUserModeProcessActionType = 2
	DebuggerAttachDetachUserModeProcessActionKillProcess                        DebuggerAttachDetachUserModeProcessActionType = 3
	DebuggerAttachDetachUserModeProcessActionContinueProcess                    DebuggerAttachDetachUserModeProcessActionType = 4
	DebuggerAttachDetachUserModeProcessActionPauseProcess                       DebuggerAttachDetachUserModeProcessActionType = 5
	DebuggerAttachDetachUserModeProcessActionSwitchByProcessOrThread            DebuggerAttachDetachUserModeProcessActionType = 6
	DebuggerAttachDetachUserModeProcessActionQueryCountOfActiveDebuggingThreads DebuggerAttachDetachUserModeProcessActionType = 7
)

type DebuggerAttachDetachUserModeProcess struct {
	IsStartingNewProcess                      uint8
	_                                         [3]byte
	ProcessId                                 uint32
	ThreadId                                  uint32
	CheckCallbackAtFirstInstruction           uint8
	Is32Bit                                   uint8
	_                                         [2]byte
	Rip                                       uint64
	InstructionBytesOnRip                     [16]byte
	SizeOfInstruction                         uint32
	IsPaused                                  uint8
	_                                         [3]byte
	Action                                    DebuggerAttachDetachUserModeProcessActionType
	CountOfActiveDebuggingThreadsAndProcesses uint32
	Token                                     uint64
	Result                                    uint64
}

type DebuggeeBpPacket struct {
	Address           uint64
	Pid               uint32
	Tid               uint32
	Core              uint32
	RemoveAfterHit    uint8
	CheckForCallbacks uint8
	_                 [2]byte
	Result            uint32
	_                 [4]byte
}

type DebuggeeBreakpointModificationRequest uint32

const (
	DebuggeeBreakpointModificationRequestListBreakpoints DebuggeeBreakpointModificationRequest = 0
	DebuggeeBreakpointModificationRequestEnable          DebuggeeBreakpointModificationRequest = 1
	DebuggeeBreakpointModificationRequestDisable         DebuggeeBreakpointModificationRequest = 2
	DebuggeeBreakpointModificationRequestClear           DebuggeeBreakpointModificationRequest = 3
)

type DebuggeeBpListOrModifyPacket struct {
	BreakpointId uint64
	Request      DebuggeeBreakpointModificationRequest
	Result       uint32
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

type DebuggerUdCommandActionType uint32

const (
	DebuggerUdCommandActionTypeNone                DebuggerUdCommandActionType = 0
	DebuggerUdCommandActionTypePause               DebuggerUdCommandActionType = 1
	DebuggerUdCommandActionTypeRegularStep         DebuggerUdCommandActionType = 2
	DebuggerUdCommandActionTypeReadRegisters       DebuggerUdCommandActionType = 3
	DebuggerUdCommandActionTypeExecuteScriptBuffer DebuggerUdCommandActionType = 4
	DebuggerUdCommandActionTypeWriteRegisters      DebuggerUdCommandActionType = 5
	DebuggerUdCommandActionTypeCallstack           DebuggerUdCommandActionType = 6
)

type DebuggerUdCommandAction struct {
	ActionType     DebuggerUdCommandActionType
	OptionalParam1 uint64
	OptionalParam2 uint64
	OptionalParam3 uint64
	OptionalParam4 uint64
}

type DebuggerUdCommandPacket struct {
	UdAction                    DebuggerUdCommandAction
	ProcessDebuggingDetailToken uint64
	TargetThreadId              uint32
	ApplyToAllPausedThreads     uint8
	WaitForEventCompletion      uint8
	Result                      uint32
}

type DebuggerSetBreakpointUserDebugger struct {
	ProcessId                   uint32
	ThreadId                    uint32
	BreakpointType              uint32
	_                           [4]byte
	Address                     uint64
	ProcessDebuggingDetailToken uint64
	Result                      uint32
	_                           [4]byte
}

func (s *DebuggerAttachDetachUserModeProcess) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.ThreadId == 0 {
		return errors.New("ThreadId cannot be zero")
	}
	return nil
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSize() uintptr {
	return 72
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSerSize() uintptr {
	return 72
}

func (s *DebuggerAttachDetachUserModeProcess) Serialize() []byte {
	buf := make([]byte, 72)
	buf[0] = s.IsStartingNewProcess
	binary.LittleEndian.PutUint32(buf[4:], s.ProcessId)
	binary.LittleEndian.PutUint32(buf[8:], s.ThreadId)
	buf[12] = s.CheckCallbackAtFirstInstruction
	buf[13] = s.Is32Bit
	binary.LittleEndian.PutUint64(buf[16:], s.Rip)
	copy(buf[24:], s.InstructionBytesOnRip[:])
	binary.LittleEndian.PutUint32(buf[40:], s.SizeOfInstruction)
	buf[44] = s.IsPaused
	binary.LittleEndian.PutUint32(buf[48:], uint32(s.Action))
	binary.LittleEndian.PutUint32(buf[52:], s.CountOfActiveDebuggingThreadsAndProcesses)
	binary.LittleEndian.PutUint64(buf[56:], s.Token)
	binary.LittleEndian.PutUint64(buf[64:], s.Result)
	return buf
}

func (s *DebuggerAttachDetachUserModeProcess) Deserialize(data []byte) {
	if len(data) < 72 {
		return
	}
	s.IsStartingNewProcess = data[0]
	s.ProcessId = binary.LittleEndian.Uint32(data[4:])
	s.ThreadId = binary.LittleEndian.Uint32(data[8:])
	s.CheckCallbackAtFirstInstruction = data[12]
	s.Is32Bit = data[13]
	s.Rip = binary.LittleEndian.Uint64(data[16:])
	copy(s.InstructionBytesOnRip[:], data[24:40])
	s.SizeOfInstruction = binary.LittleEndian.Uint32(data[40:])
	s.IsPaused = data[44]
	s.Action = DebuggerAttachDetachUserModeProcessActionType(binary.LittleEndian.Uint32(data[48:]))
	s.CountOfActiveDebuggingThreadsAndProcesses = binary.LittleEndian.Uint32(data[52:])
	s.Token = binary.LittleEndian.Uint64(data[56:])
	s.Result = binary.LittleEndian.Uint64(data[64:])
}

func (s *DebuggeeBpPacket) Validate() error {
	if s.Pid == 0 {
		return errors.New("Pid cannot be zero")
	}
	if s.Tid == 0 {
		return errors.New("Tid cannot be zero")
	}
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggeeBpPacket) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggeeBpPacket) ExpectedSerSize() uintptr {
	return 32
}

func (s *DebuggeeBpListOrModifyPacket) Validate() error {
	return nil
}

func (s *DebuggeeBpListOrModifyPacket) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggeeBpListOrModifyPacket) ExpectedSerSize() uintptr {
	return 16
}

func (s *DebuggerUdCommandPacket) Validate() error {
	if s.ProcessDebuggingDetailToken == 0 {
		return errors.New("ProcessDebuggingDetailToken cannot be zero")
	}
	return nil
}

func (s *DebuggerUdCommandPacket) ExpectedSize() uintptr {
	return 64
}

func (s *DebuggerUdCommandPacket) ExpectedSerSize() uintptr {
	return 64
}

func (s *DebuggerUdCommandAction) Validate() error {
	return nil
}

func (s *DebuggerUdCommandAction) ExpectedSize() uintptr {
	return 40
}

func (s *DebuggerUdCommandAction) ExpectedSerSize() uintptr {
	return 40
}

func (s *DebuggerReadMemory) Validate() error {
	if s.Pid == 0 {
		return errors.New("Pid cannot be zero")
	}
	if s.Size == 0 {
		return errors.New("Size cannot be zero")
	}
	return nil
}

func (s *DebuggerReadMemory) ExpectedSize() uintptr {
	return 48
}

func (s *DebuggerReadMemory) ExpectedSerSize() uintptr {
	return 48
}

func (s *DebuggerSetBreakpointUserDebugger) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.ThreadId == 0 {
		return errors.New("ThreadId cannot be zero")
	}
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	if s.ProcessDebuggingDetailToken == 0 {
		return errors.New("ProcessDebuggingDetailToken cannot be zero")
	}
	return nil
}

func (s *DebuggerSetBreakpointUserDebugger) ExpectedSize() uintptr {
	return 40
}

func (s *DebuggerSetBreakpointUserDebugger) ExpectedSerSize() uintptr {
	return 40
}

func (s *DebuggerVa2paAndPa2vaCommands) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerVa2paAndPa2vaCommands) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggerVa2paAndPa2vaCommands) ExpectedSerSize() uintptr {
	return 32
}

func (s *DebuggerGeneralEventDetails) Validate() error {
	return nil
}

func (s *DebuggerGeneralEventDetails) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerGeneralEventDetails) ExpectedSerSize() uintptr {
	return 24
}

func (s *DebuggerEventAction) Validate() error {
	return nil
}

func (s *DebuggerEventAction) ExpectedSize() uintptr {
	return 56
}

func (s *DebuggerEventAction) ExpectedSerSize() uintptr {
	return 56
}

func (s *DebuggerModifyEventRequest) Validate() error {
	return nil
}

func (s *DebuggerModifyEventRequest) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerModifyEventRequest) ExpectedSerSize() uintptr {
	return 24
}

func (s *DebuggerEventAndActionResult) Validate() error {
	return nil
}

func (s *DebuggerEventAndActionResult) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerEventAndActionResult) ExpectedSerSize() uintptr {
	return 8
}

func (s *DebuggeeKdPausedPacket) Validate() error {
	return nil
}

func (s *DebuggeeKdPausedPacket) ExpectedSize() uintptr {
	return 72
}

func (s *DebuggeeKdPausedPacket) ExpectedSerSize() uintptr {
	return 72
}

func (s *DebuggeeDetailsAndSwitchProcessPacket) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggeeDetailsAndSwitchProcessPacket) ExpectedSize() uintptr {
	return 72
}

func (s *DebuggeeDetailsAndSwitchProcessPacket) ExpectedSerSize() uintptr {
	return 72
}

func (s *DebuggeeDetailsAndSwitchThreadPacket) Validate() error {
	if s.ThreadId == 0 {
		return errors.New("ThreadId cannot be zero")
	}
	return nil
}

func (s *DebuggeeDetailsAndSwitchThreadPacket) ExpectedSize() uintptr {
	return 48
}

func (s *DebuggeeDetailsAndSwitchThreadPacket) ExpectedSerSize() uintptr {
	return 48
}

func (s *DebuggeeRegisterReadDescription) Validate() error {
	return nil
}

func (s *DebuggeeRegisterReadDescription) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggeeRegisterReadDescription) ExpectedSerSize() uintptr {
	return 24
}

func (s *PageTableEntries) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *PageTableEntries) ExpectedSize() uintptr {
	return 80
}

func (s *PageTableEntries) ExpectedSerSize() uintptr {
	return 80
}

func (s *DebuggerReadPageTableEntriesDetails) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerReadPageTableEntriesDetails) ExpectedSize() uintptr {
	return 88
}

func (s *DebuggerReadPageTableEntriesDetails) ExpectedSerSize() uintptr {
	return 88
}

func (s *CoreInfo) Validate() error {
	return nil
}

func (s *CoreInfo) ExpectedSize() uintptr {
	return 16
}

func (s *CoreInfo) ExpectedSerSize() uintptr {
	return 16
}

func (s *VersionInfo) Validate() error {
	return nil
}

func (s *VersionInfo) ExpectedSize() uintptr {
	return 20
}

func (s *VersionInfo) ExpectedSerSize() uintptr {
	return 20
}

func (s *UserModeProcessDetails) Validate() error {
	if s.Token == 0 {
		return errors.New("Token cannot be zero")
	}
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *UserModeProcessDetails) ExpectedSize() uintptr {
	return 48
}

func (s *UserModeProcessDetails) ExpectedSerSize() uintptr {
	return 48
}

func (s *DebuggerReadOrWriteMsr) Validate() error {
	return nil
}

func (s *DebuggerReadOrWriteMsr) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggerReadOrWriteMsr) ExpectedSerSize() uintptr {
	return 32
}

func (s *DebuggerEditMemory) Validate() error {
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerEditMemory) ExpectedSize() uintptr {
	return 40
}

func (s *DebuggerEditMemory) ExpectedSerSize() uintptr {
	return 40
}

func (s *DebuggerSearchMemory) Validate() error {
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerSearchMemory) ExpectedSize() uintptr {
	return 40
}

func (s *DebuggerSearchMemory) ExpectedSerSize() uintptr {
	return 40
}

func (s *DebuggerPageInRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerPageInRequest) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggerPageInRequest) ExpectedSerSize() uintptr {
	return 32
}

func (s *DebuggerPreallocCommand) Validate() error {
	return nil
}

func (s *DebuggerPreallocCommand) ExpectedSize() uintptr {
	return 12
}

func (s *DebuggerPreallocCommand) ExpectedSerSize() uintptr {
	return 12
}

func (s *DebuggerHideDebuggerRequest) Validate() error {
	return nil
}

func (s *DebuggerHideDebuggerRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerHideDebuggerRequest) ExpectedSerSize() uintptr {
	return 16
}

func (s *DebuggerPteRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerPteRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerPteRequest) ExpectedSerSize() uintptr {
	return 16
}

func (s *DebuggerVa2paRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerVa2paRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerVa2paRequest) ExpectedSerSize() uintptr {
	return 16
}

func (s *DebuggerPa2vaRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerPa2vaRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerPa2vaRequest) ExpectedSerSize() uintptr {
	return 16
}

func (s *DebuggerReconstructMemoryRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerReconstructMemoryRequest) ExpectedSize() uintptr {
	return 20
}

func (s *DebuggerReconstructMemoryRequest) ExpectedSerSize() uintptr {
	return 20
}

func (s *DebuggerEventRegisterRequest) Validate() error {
	return nil
}

func (s *DebuggerEventRegisterRequest) ExpectedSize() uintptr {
	return 56
}

func (s *DebuggerEventRegisterRequest) ExpectedSerSize() uintptr {
	return 56
}

func (s *DebuggerEpthookRequest) Validate() error {
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerEpthookRequest) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerEpthookRequest) ExpectedSerSize() uintptr {
	return 24
}

func (s *DebuggerMonitorRequest) Validate() error {
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerMonitorRequest) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerMonitorRequest) ExpectedSerSize() uintptr {
	return 24
}

func (s *DebuggerSyscallHookRequest) Validate() error {
	return nil
}

func (s *DebuggerSyscallHookRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerSyscallHookRequest) ExpectedSerSize() uintptr {
	return 8
}

func (s *DebuggerExceptionHookRequest) Validate() error {
	return nil
}

func (s *DebuggerExceptionHookRequest) ExpectedSize() uintptr {
	return 4
}

func (s *DebuggerExceptionHookRequest) ExpectedSerSize() uintptr {
	return 4
}

func (s *DebuggerInterruptHookRequest) Validate() error {
	return nil
}

func (s *DebuggerInterruptHookRequest) ExpectedSize() uintptr {
	return 4
}

func (s *DebuggerInterruptHookRequest) ExpectedSerSize() uintptr {
	return 4
}

func (s *DebuggerIoHookRequest) Validate() error {
	return nil
}

func (s *DebuggerIoHookRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerIoHookRequest) ExpectedSerSize() uintptr {
	return 8
}

func (s *DebuggerApicRequest) Validate() error {
	return nil
}

func (s *DebuggerApicRequest) ExpectedSize() uintptr {
	return 4
}

func (s *DebuggerApicRequest) ExpectedSerSize() uintptr {
	return 4
}

func (s *DebuggerPciCamRequest) Validate() error {
	return nil
}

func (s *DebuggerPciCamRequest) ExpectedSize() uintptr {
	return 12
}

func (s *DebuggerPciCamRequest) ExpectedSerSize() uintptr {
	return 12
}

func (s *DebuggerPmcRequest) Validate() error {
	return nil
}

func (s *DebuggerPmcRequest) ExpectedSize() uintptr {
	return 4
}

func (s *DebuggerPmcRequest) ExpectedSerSize() uintptr {
	return 4
}

func (s *DebuggerTrackMemoryRequest) Validate() error {
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerTrackMemoryRequest) ExpectedSize() uintptr {
	return 12
}

func (s *DebuggerTrackMemoryRequest) ExpectedSerSize() uintptr {
	return 12
}

func (s *DebuggerMeasureRequest) Validate() error {
	return nil
}

func (s *DebuggerMeasureRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerMeasureRequest) ExpectedSerSize() uintptr {
	return 8
}

func (s *DebuggerInstructionTraceRequest) Validate() error {
	return nil
}

func (s *DebuggerInstructionTraceRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerInstructionTraceRequest) ExpectedSerSize() uintptr {
	return 8
}

func (s *DebuggerSwitchThreadRequest) Validate() error {
	if s.ThreadId == 0 {
		return errors.New("ThreadId cannot be zero")
	}
	return nil
}

func (s *DebuggerSwitchThreadRequest) ExpectedSize() uintptr {
	return 12
}

func (s *DebuggerSwitchThreadRequest) ExpectedSerSize() uintptr {
	return 12
}

func (s *DebuggerQueryProcessThreadRequest) Validate() error {
	return nil
}

func (s *DebuggerQueryProcessThreadRequest) ExpectedSize() uintptr {
	return 20
}

func (s *DebuggerQueryProcessThreadRequest) ExpectedSerSize() uintptr {
	return 20
}

type DebuggerSearchMemoryPatternRequest struct {
	ProcessId   uint32
	PatternSize uint32
	Mode        ReconstructMode
	_           [4]byte
}

func (s *DebuggerSearchMemoryPatternRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.PatternSize == 0 {
		return errors.New("PatternSize cannot be zero")
	}
	return nil
}

func (s *DebuggerSearchMemoryPatternRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerSearchMemoryPatternRequest) ExpectedSerSize() uintptr {
	return 16
}

type DebuggerQueryProcessRequest struct {
	ProcessId uint32
	_         [4]byte
}

func (s *DebuggerQueryProcessRequest) Validate() error {
	return nil
}

func (s *DebuggerQueryProcessRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerQueryProcessRequest) ExpectedSerSize() uintptr {
	return 8
}

type DebuggerQueryThreadRequest struct {
	ThreadId uint32
	_        [4]byte
}

func (s *DebuggerQueryThreadRequest) Validate() error {
	return nil
}

func (s *DebuggerQueryThreadRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerQueryThreadRequest) ExpectedSerSize() uintptr {
	return 8
}

type DebuggerQueryThreadsRequest struct {
	ProcessId uint32
	_         [4]byte
}

func (s *DebuggerQueryThreadsRequest) Validate() error {
	return nil
}

func (s *DebuggerQueryThreadsRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerQueryThreadsRequest) ExpectedSerSize() uintptr {
	return 8
}

type DebuggerSmiRequest struct {
	Operation SMIType
}

func (s *DebuggerSmiRequest) Validate() error {
	return nil
}

func (s *DebuggerSmiRequest) ExpectedSize() uintptr {
	return 4
}

func (s *DebuggerSmiRequest) ExpectedSerSize() uintptr {
	return 4
}

type DebuggerHideRequest struct {
	IsHide bool
}

func (s *DebuggerHideRequest) Validate() error {
	return nil
}

func (s *DebuggerHideRequest) ExpectedSize() uintptr {
	return 1
}

func (s *DebuggerHideRequest) ExpectedSerSize() uintptr {
	return 1
}

type DebuggerBringPagesInRequest struct {
	FromAddress uint64
	ToAddress   uint64
	ProcessId   uint32
	_           [4]byte
}

func (s *DebuggerBringPagesInRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerBringPagesInRequest) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerBringPagesInRequest) ExpectedSerSize() uintptr {
	return 24
}

type DebuggerEditMemoryRequest struct {
	ProcessId      uint32
	_              [4]byte
	Address        uint64
	ByteCount      uint32
	MemoryType     MemoryType
	IsDebuggee     bool
	IsPhysical     bool
	Is32BitProcess bool
	_              [1]byte
}

func (s *DebuggerEditMemoryRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerEditMemoryRequest) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggerEditMemoryRequest) ExpectedSerSize() uintptr {
	return 32
}

type DebuggerReadMemoryRequest struct {
	ProcessId      uint32
	_              [4]byte
	Address        uint64
	ByteCount      uint32
	MemoryType     MemoryType
	IsDebuggee     bool
	IsPhysical     bool
	Is32BitProcess bool
	_              [1]byte
}

func (s *DebuggerReadMemoryRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.Address == 0 {
		return errors.New("Address cannot be zero")
	}
	return nil
}

func (s *DebuggerReadMemoryRequest) ExpectedSize() uintptr {
	return 32
}

func (s *DebuggerReadMemoryRequest) ExpectedSerSize() uintptr {
	return 32
}

type DebuggerSearchMemoryRequest struct {
	ProcessId   uint32
	_           [4]byte
	Address     uint64
	Size        uint32
	PatternSize uint32
}

func (s *DebuggerSearchMemoryRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	if s.PatternSize == 0 {
		return errors.New("PatternSize cannot be zero")
	}
	return nil
}

func (s *DebuggerSearchMemoryRequest) ExpectedSize() uintptr {
	return 24
}

func (s *DebuggerSearchMemoryRequest) ExpectedSerSize() uintptr {
	return 24
}

type DebuggerSwitchProcessRequest struct {
	ProcessId uint32
	_         [4]byte
}

func (s *DebuggerSwitchProcessRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerSwitchProcessRequest) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerSwitchProcessRequest) ExpectedSerSize() uintptr {
	return 8
}

type DebuggerAttachDetachUserModeProcessRequest struct {
	ProcessId uint32
	Action    DebuggerAttachDetachUserModeProcessActionType
	_         [4]byte
}

func (s *DebuggerAttachDetachUserModeProcessRequest) Validate() error {
	if s.ProcessId == 0 {
		return errors.New("ProcessId cannot be zero")
	}
	return nil
}

func (s *DebuggerAttachDetachUserModeProcessRequest) ExpectedSize() uintptr {
	return 12
}

func (s *DebuggerAttachDetachUserModeProcessRequest) ExpectedSerSize() uintptr {
	return 12
}

type DebuggerAttachDetachUserModeProcessResponse struct {
	KernelStatus uint32
	_            [4]byte
}

func (s *DebuggerAttachDetachUserModeProcessResponse) Validate() error {
	return nil
}

func (s *DebuggerAttachDetachUserModeProcessResponse) ExpectedSize() uintptr {
	return 8
}

func (s *DebuggerAttachDetachUserModeProcessResponse) ExpectedSerSize() uintptr {
	return 8
}

type DebuggerGetUserModeModuleDetailsRequest struct {
	ProcessDebuggingDetailToken uint64
	_                           [8]byte
}

func (s *DebuggerGetUserModeModuleDetailsRequest) Validate() error {
	if s.ProcessDebuggingDetailToken == 0 {
		return errors.New("ProcessDebuggingDetailToken cannot be zero")
	}
	return nil
}

func (s *DebuggerGetUserModeModuleDetailsRequest) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerGetUserModeModuleDetailsRequest) ExpectedSerSize() uintptr {
	return 16
}

type DebuggerModuleDetailResponse struct {
	BaseAddress uint64
	Size        uint32
	Is32Bit     uint8
	_           [3]byte
}

func (s *DebuggerModuleDetailResponse) Validate() error {
	return nil
}

func (s *DebuggerModuleDetailResponse) ExpectedSize() uintptr {
	return 16
}

func (s *DebuggerModuleDetailResponse) ExpectedSerSize() uintptr {
	return 16
}
