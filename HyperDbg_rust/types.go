package hyperdbgrust

import "time"

const (
	ProtocolVersion = 1
	DefaultPort     = 9527
	MaxMessageSize  = 4096
	HeaderSize      = 20

	DriverHTTPHost = "127.0.0.1"
	DriverHTTPPort = 50080
	HTTPTimeout    = 10 * time.Second
)

type MessageType uint32

const (
	MsgTypeBreakpointEvent   MessageType = 1
	MsgTypeExceptionEvent    MessageType = 2
	MsgTypeMemoryAccessEvent MessageType = 3
	MsgTypeSyscallEvent      MessageType = 4
	MsgTypeProcessEvent      MessageType = 5
	MsgTypeThreadEvent       MessageType = 6
	MsgTypeModuleEvent       MessageType = 7
	MsgTypeDebugPrintEvent   MessageType = 8
	MsgTypeVmxExitEvent      MessageType = 9
	MsgTypeLogEvent          MessageType = 10

	MsgTypeInitialize        MessageType = 100
	MsgTypeTerminate         MessageType = 101
	MsgTypeConnect           MessageType = 102
	MsgTypeDisconnect        MessageType = 103
	MsgTypeAttachProcess     MessageType = 104
	MsgTypeDetachProcess     MessageType = 105
	MsgTypeStartProcess      MessageType = 106
	MsgTypeKillProcess       MessageType = 107
	MsgTypeSetBreakpoint     MessageType = 108
	MsgTypeRemoveBreakpoint  MessageType = 109
	MsgTypeEnableBreakpoint  MessageType = 110
	MsgTypeDisableBreakpoint MessageType = 111
	MsgTypeListBreakpoints   MessageType = 112
	MsgTypeContinue          MessageType = 113
	MsgTypePause             MessageType = 114
	MsgTypeStepInto          MessageType = 115
	MsgTypeStepOver          MessageType = 116
	MsgTypeStepOut           MessageType = 117
	MsgTypeReadRegisters     MessageType = 118
	MsgTypeWriteRegisters    MessageType = 119
	MsgTypeReadMemory        MessageType = 120
	MsgTypeWriteMemory       MessageType = 121
	MsgTypeGetProcessList    MessageType = 122
	MsgTypeGetThreadList     MessageType = 123
	MsgTypeGetModuleList     MessageType = 124
	MsgTypeGetCallStack      MessageType = 125

	MsgTypeResponse MessageType = 200
	MsgTypeError    MessageType = 201
)

type BreakpointType uint32

const (
	BreakpointSoftware BreakpointType = 0
	BreakpointHardware BreakpointType = 1
	BreakpointHidden   BreakpointType = 2
	BreakpointEpt      BreakpointType = 3
)

type DebugState uint32

const (
	StateRunning    DebugState = 0
	StatePaused     DebugState = 1
	StateStepping   DebugState = 2
	StateTerminated DebugState = 3
)

type MessageHeader struct {
	Type     MessageType
	Length   uint32
	Sequence uint64
	Flags    uint32
}

type Message struct {
	Header  MessageHeader
	Payload []byte
}

type RegisterState struct {
	RAX    uint64 `json:"RAX"`
	RBX    uint64 `json:"RBX"`
	RCX    uint64 `json:"RCX"`
	RDX    uint64 `json:"RDX"`
	RSI    uint64 `json:"RSI"`
	RDI    uint64 `json:"RDI"`
	RBP    uint64 `json:"RBP"`
	RSP    uint64 `json:"RSP"`
	R8     uint64 `json:"R8"`
	R9     uint64 `json:"R9"`
	R10    uint64 `json:"R10"`
	R11    uint64 `json:"R11"`
	R12    uint64 `json:"R12"`
	R13    uint64 `json:"R13"`
	R14    uint64 `json:"R14"`
	R15    uint64 `json:"R15"`
	RIP    uint64 `json:"RIP"`
	RFLAGS uint64 `json:"RFLAGS"`
	CR0    uint64 `json:"CR0"`
	CR2    uint64 `json:"CR2"`
	CR3    uint64 `json:"CR3"`
	CR4    uint64 `json:"CR4"`
	DR0    uint64 `json:"DR0"`
	DR1    uint64 `json:"DR1"`
	DR2    uint64 `json:"DR2"`
	DR3    uint64 `json:"DR3"`
	DR6    uint64 `json:"DR6"`
	DR7    uint64 `json:"DR7"`
	GDTR   uint64 `json:"GDTR"`
	GSBase uint64 `json:"GSBase"`
	FSBase uint64 `json:"FSBase"`
}

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

type Response[T ResponseType] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type Empty = struct{}

type ResponseType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string | ~bool |
		[]byte | *RegisterState | []ProcessInfo | []ThreadInfo | []ModuleInfo | []BreakpointInfo |
		Empty
}

type EventCallback func(event any)

type ProcessId = uint32
type ThreadId = uint32
type Address = uint64
type PhysicalAddress = uint64

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

type LogLevel uint32

const (
	LogLevelTrace LogLevel = 0
	LogLevelDebug LogLevel = 1
	LogLevelInfo  LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelError LogLevel = 4
)

type VmxCapabilities struct {
	VMXSupported            bool   `json:"vmx_supported"`
	EPTSupported            bool   `json:"ept_supported"`
	VPIDSupported           bool   `json:"vpid_supported"`
	MSRBitmapSupported      bool   `json:"msr_bitmap_supported"`
	IOBitmapSupported       bool   `json:"io_bitmap_supported"`
	MaxPhysicalAddressWidth uint8  `json:"max_physical_address_width"`
	ProcessorCount          uint32 `json:"processor_count"`
}

type EventType uint32

const (
	EventTypeBreakpoint      EventType = 1
	EventTypeException       EventType = 2
	EventTypeMemoryAccess    EventType = 3
	EventTypeSyscallEntry    EventType = 4
	EventTypeSyscallExit     EventType = 5
	EventTypeProcessCreate   EventType = 6
	EventTypeProcessExit     EventType = 7
	EventTypeThreadCreate    EventType = 8
	EventTypeThreadExit      EventType = 9
	EventTypeModuleLoad      EventType = 10
	EventTypeModuleUnload    EventType = 11
	EventTypeDebugPrint      EventType = 12
	EventTypeVmxExit         EventType = 13
	EventTypeTrap            EventType = 14
	EventTypeHiddenHookExec  EventType = 15
	EventTypeHiddenHookRead  EventType = 16
	EventTypeHiddenHookWrite EventType = 17
	EventTypeCpuid           EventType = 18
	EventTypeTsc             EventType = 19
	EventTypePmc             EventType = 20
	EventTypeInterrupt       EventType = 21
	EventTypeExceptionBitmap EventType = 22
	EventTypeCrAccess        EventType = 23
	EventTypeDrAccess        EventType = 24
	EventTypeIoPort          EventType = 25
	EventTypeMsrRead         EventType = 26
	EventTypeMsrWrite        EventType = 27
	EventTypeEptViolation    EventType = 28
	EventTypeVmcalled        EventType = 29
)

type EventHeader struct {
	EventType EventType `json:"event_type"`
	ProcessID ProcessId `json:"process_id"`
	ThreadID  ThreadId  `json:"thread_id"`
	CoreID    uint32    `json:"core_id"`
	Timestamp uint64    `json:"timestamp"`
}

type BreakpointEvent struct {
	Header       EventHeader   `json:"header"`
	Address      Address       `json:"address"`
	BreakpointID uint64        `json:"breakpoint_id"`
	Registers    RegisterState `json:"registers"`
}

type ExceptionCode uint32

const (
	ExceptionCodeDivideError        ExceptionCode = 0
	ExceptionCodeDebug              ExceptionCode = 1
	ExceptionCodeNmi                ExceptionCode = 2
	ExceptionCodeBreakpoint         ExceptionCode = 3
	ExceptionCodeOverflow           ExceptionCode = 4
	ExceptionCodeBound              ExceptionCode = 5
	ExceptionCodeInvalidOpcode      ExceptionCode = 6
	ExceptionCodeNoMath             ExceptionCode = 7
	ExceptionCodeDoubleFault        ExceptionCode = 8
	ExceptionCodeCoprocessorSegment ExceptionCode = 9
	ExceptionCodeInvalidTss         ExceptionCode = 10
	ExceptionCodeSegmentNotPresent  ExceptionCode = 11
	ExceptionCodeStackSegment       ExceptionCode = 12
	ExceptionCodeGeneralProtection  ExceptionCode = 13
	ExceptionCodePageFault          ExceptionCode = 14
	ExceptionCodeFloatingPoint      ExceptionCode = 16
	ExceptionCodeAlignmentCheck     ExceptionCode = 17
	ExceptionCodeMachineCheck       ExceptionCode = 18
	ExceptionCodeSimdFloatingPoint  ExceptionCode = 19
)

type ExceptionEvent struct {
	Header        EventHeader   `json:"header"`
	ExceptionCode ExceptionCode `json:"exception_code"`
	Address       Address       `json:"address"`
	ErrorCode     uint64        `json:"error_code"`
	Registers     RegisterState `json:"registers"`
}

type MemoryAccessType uint32

const (
	MemoryAccessTypeRead    MemoryAccessType = 0
	MemoryAccessTypeWrite   MemoryAccessType = 1
	MemoryAccessTypeExecute MemoryAccessType = 2
)

type MemoryAccessEvent struct {
	Header          EventHeader      `json:"header"`
	VirtualAddress  Address          `json:"virtual_address"`
	PhysicalAddress PhysicalAddress  `json:"physical_address"`
	AccessType      MemoryAccessType `json:"access_type"`
	Size            uint32           `json:"size"`
	Value           uint64           `json:"value"`
}

type SyscallEvent struct {
	Header        EventHeader `json:"header"`
	SyscallNumber uint32      `json:"syscall_number"`
	Arguments     [8]uint64   `json:"arguments"`
	ReturnValue   uint64      `json:"return_value"`
	IsEntry       bool        `json:"is_entry"`
}

type ProcessEvent struct {
	Header          EventHeader `json:"header"`
	ProcessInfo     ProcessInfo `json:"process_info"`
	ParentProcessID ProcessId   `json:"parent_process_id"`
	IsCreate        bool        `json:"is_create"`
}

type ThreadEvent struct {
	Header     EventHeader `json:"header"`
	ThreadInfo ThreadInfo  `json:"thread_info"`
	IsCreate   bool        `json:"is_create"`
}

type ModuleEvent struct {
	Header     EventHeader `json:"header"`
	ModuleInfo ModuleInfo  `json:"module_info"`
	IsLoad     bool        `json:"is_load"`
}

type DebugPrintEvent struct {
	Header  EventHeader `json:"header"`
	Message string      `json:"message"`
	Level   LogLevel    `json:"level"`
}

type VmxExitReason uint32

const (
	VmxExitReasonExceptionNmi          VmxExitReason = 0
	VmxExitReasonExternalInterrupt     VmxExitReason = 1
	VmxExitReasonTripleFault           VmxExitReason = 2
	VmxExitReasonInitSignal            VmxExitReason = 3
	VmxExitReasonStartupIpi            VmxExitReason = 4
	VmxExitReasonIoSmi                 VmxExitReason = 5
	VmxExitReasonOtherSmi              VmxExitReason = 6
	VmxExitReasonInterruptWindow       VmxExitReason = 7
	VmxExitReasonNmiWindow             VmxExitReason = 8
	VmxExitReasonTaskSwitch            VmxExitReason = 9
	VmxExitReasonCpuid                 VmxExitReason = 10
	VmxExitReasonGetsec                VmxExitReason = 11
	VmxExitReasonHlt                   VmxExitReason = 12
	VmxExitReasonInvd                  VmxExitReason = 13
	VmxExitReasonInvlpg                VmxExitReason = 14
	VmxExitReasonRdpmc                 VmxExitReason = 15
	VmxExitReasonRdtsc                 VmxExitReason = 16
	VmxExitReasonRsm                   VmxExitReason = 17
	VmxExitReasonVmcall                VmxExitReason = 18
	VmxExitReasonVmclear               VmxExitReason = 19
	VmxExitReasonVmlaunch              VmxExitReason = 20
	VmxExitReasonVmptrld               VmxExitReason = 21
	VmxExitReasonVmptrst               VmxExitReason = 22
	VmxExitReasonVmread                VmxExitReason = 23
	VmxExitReasonVmresume              VmxExitReason = 24
	VmxExitReasonVmwrite               VmxExitReason = 25
	VmxExitReasonVmxoff                VmxExitReason = 26
	VmxExitReasonVmxon                 VmxExitReason = 27
	VmxExitReasonCrAccess              VmxExitReason = 28
	VmxExitReasonMovDr                 VmxExitReason = 29
	VmxExitReasonIoInstruction         VmxExitReason = 30
	VmxExitReasonRdmsr                 VmxExitReason = 31
	VmxExitReasonWrmsr                 VmxExitReason = 32
	VmxExitReasonEntryFailGuestState   VmxExitReason = 33
	VmxExitReasonEntryFailMsrLoad      VmxExitReason = 34
	VmxExitReasonMwait                 VmxExitReason = 36
	VmxExitReasonMtf                   VmxExitReason = 37
	VmxExitReasonMonitor               VmxExitReason = 39
	VmxExitReasonPause                 VmxExitReason = 40
	VmxExitReasonEntryFailMachineCheck VmxExitReason = 41
	VmxExitReasonTprBelowThreshold     VmxExitReason = 43
	VmxExitReasonApicAccess            VmxExitReason = 44
	VmxExitReasonVirtualizedEoi        VmxExitReason = 45
	VmxExitReasonGdtrIdtrAccess        VmxExitReason = 46
	VmxExitReasonLdtrTrAccess          VmxExitReason = 47
	VmxExitReasonEptViolation          VmxExitReason = 48
	VmxExitReasonEptMisconfig          VmxExitReason = 49
	VmxExitReasonInvept                VmxExitReason = 50
	VmxExitReasonRdtscp                VmxExitReason = 51
	VmxExitReasonVmxPreemptTimer       VmxExitReason = 52
	VmxExitReasonInvvpid               VmxExitReason = 53
	VmxExitReasonWbinvd                VmxExitReason = 54
	VmxExitReasonXsetbv                VmxExitReason = 55
	VmxExitReasonApicWrite             VmxExitReason = 56
	VmxExitReasonRdrand                VmxExitReason = 57
	VmxExitReasonInvpcid               VmxExitReason = 58
	VmxExitReasonVmfFunc               VmxExitReason = 59
	VmxExitReasonEncls                 VmxExitReason = 60
	VmxExitReasonRdseed                VmxExitReason = 61
	VmxExitReasonPmlFull               VmxExitReason = 62
	VmxExitReasonXsaves                VmxExitReason = 63
	VmxExitReasonXrstors               VmxExitReason = 64
)

type VmxExitEvent struct {
	Header            EventHeader   `json:"header"`
	ExitReason        VmxExitReason `json:"exit_reason"`
	ExitQualification uint64        `json:"exit_qualification"`
	GuestRIP          Address       `json:"guest_rip"`
	GuestRSP          Address       `json:"guest_rsp"`
	InstructionLength uint32        `json:"instruction_length"`
}

type TrapEvent struct {
	Header     EventHeader `json:"header"`
	TrapNumber uint32      `json:"trap_number"`
	ErrorCode  uint64      `json:"error_code"`
	Address    Address     `json:"address"`
}

type HiddenHookEvent struct {
	Header      EventHeader      `json:"header"`
	HookAddress Address          `json:"hook_address"`
	HookType    MemoryAccessType `json:"hook_type"`
	Data        []byte           `json:"data"`
}

type CpuidEvent struct {
	Header  EventHeader `json:"header"`
	Leaf    uint32      `json:"leaf"`
	SubLeaf uint32      `json:"sub_leaf"`
	EAX     uint32      `json:"eax"`
	EBX     uint32      `json:"ebx"`
	ECX     uint32      `json:"ecx"`
	EDX     uint32      `json:"edx"`
}

type TscEvent struct {
	Header    EventHeader `json:"header"`
	TscValue  uint64      `json:"tsc_value"`
	RdtscExit bool        `json:"rdtsc_exit"`
}

type CrAccessEvent struct {
	Header   EventHeader `json:"header"`
	CrNumber uint32      `json:"cr_number"`
	IsWrite  bool        `json:"is_write"`
	OldValue uint64      `json:"old_value"`
	NewValue uint64      `json:"new_value"`
}

type DrAccessEvent struct {
	Header   EventHeader `json:"header"`
	DrNumber uint32      `json:"dr_number"`
	IsWrite  bool        `json:"is_write"`
	Value    uint64      `json:"value"`
}

type IoPortEvent struct {
	Header  EventHeader `json:"header"`
	Port    uint16      `json:"port"`
	Size    uint32      `json:"size"`
	IsWrite bool        `json:"is_write"`
	Value   uint32      `json:"value"`
}

type MsrEvent struct {
	Header  EventHeader `json:"header"`
	Msr     uint32      `json:"msr"`
	IsWrite bool        `json:"is_write"`
	Value   uint64      `json:"value"`
}

type EptViolationEvent struct {
	Header               EventHeader     `json:"header"`
	GuestPhysicalAddress PhysicalAddress `json:"guest_physical_address"`
	GuestVirtualAddress  Address         `json:"guest_virtual_address"`
	Read                 bool            `json:"read"`
	Write                bool            `json:"write"`
	Execute              bool            `json:"execute"`
	Readable             bool            `json:"readable"`
	Writable             bool            `json:"writable"`
	Executable           bool            `json:"executable"`
}

type DebuggerEventType int

const (
	DebuggerEventBreakpoint DebuggerEventType = iota
	DebuggerEventException
	DebuggerEventMemoryAccess
	DebuggerEventSyscall
	DebuggerEventProcessCreate
	DebuggerEventProcessExit
	DebuggerEventThreadCreate
	DebuggerEventThreadExit
	DebuggerEventModuleLoad
	DebuggerEventModuleUnload
	DebuggerEventDebugPrint
	DebuggerEventVmxExit
	DebuggerEventTrap
	DebuggerEventHiddenHook
	DebuggerEventCpuid
	DebuggerEventTsc
	DebuggerEventCrAccess
	DebuggerEventDrAccess
	DebuggerEventIoPort
	DebuggerEventMsr
	DebuggerEventEptViolation
)

type DebuggerEvent struct {
	Type DebuggerEventType `json:"type"`
	Data any               `json:"data"`
}

func (t DebuggerEventType) String() string {
	if name, ok := DebuggerEventTypeNames[t]; ok {
		return name
	}
	return "Unknown"
}

var DebuggerEventTypeNames = map[DebuggerEventType]string{
	DebuggerEventBreakpoint:    "Breakpoint",
	DebuggerEventException:     "Exception",
	DebuggerEventMemoryAccess:  "MemoryAccess",
	DebuggerEventSyscall:       "Syscall",
	DebuggerEventProcessCreate: "ProcessCreate",
	DebuggerEventProcessExit:   "ProcessExit",
	DebuggerEventThreadCreate:  "ThreadCreate",
	DebuggerEventThreadExit:    "ThreadExit",
	DebuggerEventModuleLoad:    "ModuleLoad",
	DebuggerEventModuleUnload:  "ModuleUnload",
	DebuggerEventDebugPrint:    "DebugPrint",
	DebuggerEventVmxExit:       "VmxExit",
	DebuggerEventTrap:          "Trap",
	DebuggerEventHiddenHook:    "HiddenHook",
	DebuggerEventCpuid:         "Cpuid",
	DebuggerEventTsc:           "Tsc",
	DebuggerEventCrAccess:      "CrAccess",
	DebuggerEventDrAccess:      "DrAccess",
	DebuggerEventIoPort:        "IoPort",
	DebuggerEventMsr:           "Msr",
	DebuggerEventEptViolation:  "EptViolation",
}

var DebuggerEventTypeByName = map[string]DebuggerEventType{
	"Breakpoint":    DebuggerEventBreakpoint,
	"Exception":     DebuggerEventException,
	"MemoryAccess":  DebuggerEventMemoryAccess,
	"Syscall":       DebuggerEventSyscall,
	"ProcessCreate": DebuggerEventProcessCreate,
	"ProcessExit":   DebuggerEventProcessExit,
	"ThreadCreate":  DebuggerEventThreadCreate,
	"ThreadExit":    DebuggerEventThreadExit,
	"ModuleLoad":    DebuggerEventModuleLoad,
	"ModuleUnload":  DebuggerEventModuleUnload,
	"DebugPrint":    DebuggerEventDebugPrint,
	"VmxExit":       DebuggerEventVmxExit,
	"Trap":          DebuggerEventTrap,
	"HiddenHook":    DebuggerEventHiddenHook,
	"Cpuid":         DebuggerEventCpuid,
	"Tsc":           DebuggerEventTsc,
	"CrAccess":      DebuggerEventCrAccess,
	"DrAccess":      DebuggerEventDrAccess,
	"IoPort":        DebuggerEventIoPort,
	"Msr":           DebuggerEventMsr,
	"EptViolation":  DebuggerEventEptViolation,
}
