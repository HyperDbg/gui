package debugger

// EventType 定义事件类型
// 注意：内核事件类型（VMM_EVENT_TYPE_ENUM）必须与 C++ Events.h 中的 VMM_EVENT_TYPE_ENUM 完全匹配
type EventType int

const (
	// ==================== 内部命令事件类型（用于用户层事件循环）====================
	// 这些不是内核事件类型，仅用于内部事件处理
	EventTypeInternalStart EventType = -1000 + iota
	EventReadMemory
	EventWriteMemory
	EventSetBreakpoint
	EventRemoveBreakpoint
	EventSingleStep
	EventContinue
	EventAttachProcess
	EventDetachProcess
	EventRegisterEvent
	EventUnregisterEvent
	EventModifyEvent
	EventStep
	EventBreakpoint
	EventException
	EventProcessCreated
	EventProcessExited
	EventThreadCreated
	EventThreadExited
	EventModuleLoaded
	EventModuleUnloaded
)

// VMM_EVENT_TYPE_ENUM - 内核事件类型（必须与 C++ Events.h 中的 VMM_EVENT_TYPE_ENUM 完全匹配）
const (
	// EPT Memory Monitoring Events
	HiddenHookReadAndWriteAndExecute EventType = iota
	HiddenHookReadWrite
	HiddenHookReadAndExecute
	HiddenHookWriteAndExecute
	HiddenHookRead
	HiddenHookWrite
	HiddenHookExecute

	// EPT Hook Events
	HiddenHookExecDetours
	HiddenHookExecCC

	// System-call Events
	SyscallHookEferSyscall
	SyscallHookEferSysret

	// CPUID Instruction Execution Events
	CpuidInstructionExecution

	// Model-Specific Registers (MSRs) Reads/Modifications Events
	RdmsrInstructionExecution
	WrmsrInstructionExecution

	// PMIO Events
	InInstructionExecution
	OutInstructionExecution

	// Interrupts/Exceptions/Faults Events
	ExceptionOccurred
	ExternalInterruptOccurred

	// Debug Registers Events
	DebugRegistersAccessed

	// Timing & Performance Events
	TscInstructionExecution
	PmcInstructionExecution

	// VMCALL Instruction Execution Events
	VmcallInstructionExecution

	// Control Registers Events
	ControlRegisterModified
	ControlRegisterRead
	ControlRegister3Modified

	// Execution Trap Events
	TrapExecutionModeChanged
	TrapExecutionInstructionTrace

	// XSETBV Instruction Execution Events
	XsetbvInstructionExecution
)

// 别名（向后兼容）
const (
	EptHookReadWriteAndExecute = HiddenHookReadAndWriteAndExecute
	EptHookReadWrite           = HiddenHookReadWrite
	EptHookReadAndExecute      = HiddenHookReadAndExecute
	EptHookWriteAndExecute     = HiddenHookWriteAndExecute
	EptHookRead                = HiddenHookRead
	EptHookWrite               = HiddenHookWrite
	EptHookExecute             = HiddenHookExecute
	EptHookExecDetours         = HiddenHookExecDetours
	EptHookExecCC              = HiddenHookExecCC
)

func (e EventType) String() string {
	switch e {
	// 内部命令事件类型
	case EventReadMemory:
		return "EventReadMemory"
	case EventWriteMemory:
		return "EventWriteMemory"
	case EventSetBreakpoint:
		return "EventSetBreakpoint"
	case EventRemoveBreakpoint:
		return "EventRemoveBreakpoint"
	case EventSingleStep:
		return "EventSingleStep"
	case EventContinue:
		return "EventContinue"
	case EventAttachProcess:
		return "EventAttachProcess"
	case EventDetachProcess:
		return "EventDetachProcess"
	case EventRegisterEvent:
		return "EventRegisterEvent"
	case EventUnregisterEvent:
		return "EventUnregisterEvent"
	case EventModifyEvent:
		return "EventModifyEvent"
	case EventStep:
		return "EventStep"
	case EventBreakpoint:
		return "EventBreakpoint"
	case EventException:
		return "EventException"
	case EventProcessCreated:
		return "EventProcessCreated"
	case EventProcessExited:
		return "EventProcessExited"
	case EventThreadCreated:
		return "EventThreadCreated"
	case EventThreadExited:
		return "EventThreadExited"
	case EventModuleLoaded:
		return "EventModuleLoaded"
	case EventModuleUnloaded:
		return "EventModuleUnloaded"

	// 内核事件类型（VMM_EVENT_TYPE_ENUM）
	case HiddenHookReadAndWriteAndExecute:
		return "HiddenHookReadAndWriteAndExecute"
	case HiddenHookReadWrite:
		return "HiddenHookReadWrite"
	case HiddenHookReadAndExecute:
		return "HiddenHookReadAndExecute"
	case HiddenHookWriteAndExecute:
		return "HiddenHookWriteAndExecute"
	case HiddenHookRead:
		return "HiddenHookRead"
	case HiddenHookWrite:
		return "HiddenHookWrite"
	case HiddenHookExecute:
		return "HiddenHookExecute"
	case HiddenHookExecDetours:
		return "HiddenHookExecDetours"
	case HiddenHookExecCC:
		return "HiddenHookExecCC"
	case SyscallHookEferSyscall:
		return "SyscallHookEferSyscall"
	case SyscallHookEferSysret:
		return "SyscallHookEferSysret"
	case CpuidInstructionExecution:
		return "CpuidInstructionExecution"
	case RdmsrInstructionExecution:
		return "RdmsrInstructionExecution"
	case WrmsrInstructionExecution:
		return "WrmsrInstructionExecution"
	case InInstructionExecution:
		return "InInstructionExecution"
	case OutInstructionExecution:
		return "OutInstructionExecution"
	case ExceptionOccurred:
		return "ExceptionOccurred"
	case ExternalInterruptOccurred:
		return "ExternalInterruptOccurred"
	case DebugRegistersAccessed:
		return "DebugRegistersAccessed"
	case TscInstructionExecution:
		return "TscInstructionExecution"
	case PmcInstructionExecution:
		return "PmcInstructionExecution"
	case VmcallInstructionExecution:
		return "VmcallInstructionExecution"
	case ControlRegisterModified:
		return "ControlRegisterModified"
	case ControlRegisterRead:
		return "ControlRegisterRead"
	case ControlRegister3Modified:
		return "ControlRegister3Modified"
	case TrapExecutionModeChanged:
		return "TrapExecutionModeChanged"
	case TrapExecutionInstructionTrace:
		return "TrapExecutionInstructionTrace"
	case XsetbvInstructionExecution:
		return "XsetbvInstructionExecution"
	default:
		return "Unknown"
	}
}

type EventActionType int

const (
	ActionBreakToDebugger EventActionType = iota
	ActionRunScript
	ActionRunCustomCode
)

type EventModifyType int

const (
	ModifyQueryState EventModifyType = iota
	ModifyEnable
	ModifyDisable
	ModifyClear
)

type EventCallingStage int

const (
	EventStagePre EventCallingStage = iota
	EventStagePost
)

type ShortCircuitingEvent struct {
	IsShortCircuiting bool
	KernelStatus      uint64
}

type EventOptions struct {
	OptionalParam1 uint64
	OptionalParam2 uint64
	OptionalParam3 uint64
	OptionalParam4 uint64
	OptionalParam5 uint64
	OptionalParam6 uint64
}

type EventAction struct {
	EventTag                uint64
	ActionType              EventActionType
	ImmediateMessagePassing bool
	PreAllocatedBuffer      uint32
	CustomCodeBufferSize    uint32
	ScriptBufferSize        uint32
	ScriptBufferPointer     uint32
}

type ModifyEventRequest struct {
	Tag          uint64
	KernelStatus uint64
	TypeOfAction EventModifyType
	IsEnabled    bool
}

type DebugEvent struct {
	Type                  EventType
	ProcessId             uint32
	ThreadId              uint32
	CoreId                uint32
	EIP                   uint64
	RAX                   uint64
	RBX                   uint64
	RCX                   uint64
	RDX                   uint64
	RSI                   uint64
	RDI                   uint64
	RBP                   uint64
	RSP                   uint64
	R8                    uint64
	R9                    uint64
	R10                   uint64
	R11                   uint64
	R12                   uint64
	R13                   uint64
	R14                   uint64
	R15                   uint64
	RFLAGS                uint32
	Timestamp             int64
	IsEnabled             bool
	EnableShortCircuiting bool
	EventStage            EventCallingStage
	HasCustomOutput       bool
	Tag                   uint64
	Options               EventOptions
	CountOfActions        uint32
	CreationTime          int64
}

type DebugEventCallback func(event *DebugEvent)

type Event struct {
	Type     EventType
	Pid      uint32
	Address  uint64
	Size     uint32
	Data     []byte
	Metadata map[string]any
}
