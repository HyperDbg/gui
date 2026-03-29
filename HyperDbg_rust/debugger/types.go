package debugger

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
	MsgTypeHookEvent         MessageType = 11
	MsgTypeCpuidEvent       MessageType = 12

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

	MsgTypeResponse     MessageType = 200
	MsgTypeError        MessageType = 201
	MsgTypeHookResponse MessageType = 202
)

type DebugState uint32

const (
	StateRunning DebugState = iota
	StatePaused
	StateStepping
	StateTerminated
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
