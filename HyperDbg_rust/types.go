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

type BreakpointEvent struct {
	ProcessID    uint32        `json:"process_id"`
	ThreadID     uint32        `json:"thread_id"`
	CoreID       uint32        `json:"core_id"`
	Address      string        `json:"address"`
	BreakpointID uint64        `json:"breakpoint_id"`
	Registers    RegisterState `json:"registers"`
}

type ExceptionEvent struct {
	ProcessID     uint32        `json:"process_id"`
	ThreadID      uint32        `json:"thread_id"`
	CoreID        uint32        `json:"core_id"`
	ExceptionCode uint32        `json:"exception_code"`
	Address       string        `json:"address"`
	ErrorCode     uint64        `json:"error_code"`
	Registers     RegisterState `json:"registers"`
}

type DebugPrintEvent struct {
	ProcessID uint32 `json:"process_id"`
	ThreadID  uint32 `json:"thread_id"`
	CoreID    uint32 `json:"core_id"`
	Message   string `json:"message"`
	Level     uint32 `json:"level"`
}

type SyscallEvent struct {
	ProcessID     uint32 `json:"process_id"`
	ThreadID      uint32 `json:"thread_id"`
	CoreID        uint32 `json:"core_id"`
	SyscallNumber uint64 `json:"syscall_number"`
	RAX           uint64 `json:"rax"`
	RCX           uint64 `json:"rcx"`
	RDX           uint64 `json:"rdx"`
	R8            uint64 `json:"r8"`
	R9            uint64 `json:"r9"`
	R10           uint64 `json:"r10"`
	RSP           uint64 `json:"rsp"`
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
