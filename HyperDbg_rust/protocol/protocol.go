package protocol

import (
	"encoding/binary"
	"io"
)

const (
	ProtocolVersion  = 1
	DefaultPort      = 9527
	MaxMessageSize   = 4096
	HeaderSize       = 20
)

type MessageType uint32

const (
	// 驱动 -> Go 事件
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

	// Go -> 驱动命令
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

	// 响应
	MsgTypeResponse          MessageType = 200
	MsgTypeError             MessageType = 201
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

func NewMessage(msgType MessageType, payload []byte) *Message {
	return &Message{
		Header: MessageHeader{
			Type:     msgType,
			Length:   uint32(len(payload)),
			Sequence: 0,
			Flags:    0,
		},
		Payload: payload,
	}
}

func (m *Message) Serialize() []byte {
	buf := make([]byte, HeaderSize+len(m.Payload))
	binary.LittleEndian.PutUint32(buf[0:4], uint32(m.Header.Type))
	binary.LittleEndian.PutUint32(buf[4:8], m.Header.Length)
	binary.LittleEndian.PutUint64(buf[8:16], m.Header.Sequence)
	binary.LittleEndian.PutUint32(buf[16:20], m.Header.Flags)
	copy(buf[HeaderSize:], m.Payload)
	return buf
}

func ReadMessage(r io.Reader) (*Message, error) {
	headerBuf := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, headerBuf); err != nil {
		return nil, err
	}

	msg := &Message{
		Header: MessageHeader{
			Type:     MessageType(binary.LittleEndian.Uint32(headerBuf[0:4])),
			Length:   binary.LittleEndian.Uint32(headerBuf[4:8]),
			Sequence: binary.LittleEndian.Uint64(headerBuf[8:16]),
			Flags:    binary.LittleEndian.Uint32(headerBuf[16:20]),
		},
	}

	if msg.Header.Length > 0 {
		msg.Payload = make([]byte, msg.Header.Length)
		if _, err := io.ReadFull(r, msg.Payload); err != nil {
			return nil, err
		}
	}

	return msg, nil
}

func WriteMessage(w io.Writer, msg *Message) error {
	_, err := w.Write(msg.Serialize())
	return err
}

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

type RegisterState struct {
	RAX, RBX, RCX, RDX uint64
	RSI, RDI, RBP, RSP uint64
	R8, R9, R10, R11   uint64
	R12, R13, R14, R15 uint64
	RIP, RFLAGS        uint64
	CR0, CR2, CR3, CR4 uint64
	DR0, DR1, DR2, DR3 uint64
	DR6, DR7           uint64
	GDTR, GSBase, FSBase uint64
}

type ProcessInfo struct {
	ProcessID   uint32
	ImageName   string
	BaseAddress uint64
	Size        uint64
	CR3         uint64
}

type ThreadInfo struct {
	ThreadID     uint32
	ProcessID    uint32
	StartAddress uint64
	TEB          uint64
	State        DebugState
}

type ModuleInfo struct {
	BaseAddress uint64
	Size        uint64
	Name        string
	Path        string
}

type BreakpointInfo struct {
	ID         uint64
	Address    uint64
	Type       BreakpointType
	ProcessID  uint32
	Enabled    bool
	HitCount   uint64
}

type BreakpointEvent struct {
	ProcessID    uint32
	ThreadID     uint32
	CoreID       uint32
	Address      uint64
	BreakpointID uint64
	Registers    RegisterState
}

type ExceptionEvent struct {
	ProcessID     uint32
	ThreadID      uint32
	CoreID        uint32
	ExceptionCode uint32
	Address       uint64
	ErrorCode     uint64
	Registers     RegisterState
}

type DebugPrintEvent struct {
	ProcessID uint32
	ThreadID  uint32
	CoreID    uint32
	Message   string
	Level     uint32
}

type SetBreakpointCommand struct {
	Address   uint64
	Type      BreakpointType
	ProcessID uint32
}

type ReadMemoryCommand struct {
	Address   uint64
	Size      uint32
	ProcessID uint32
}

type WriteMemoryCommand struct {
	Address   uint64
	Data      []byte
	ProcessID uint32
}

type Response struct {
	Success bool
	Error   string
	Data    []byte
}

type SyscallEvent struct {
	ProcessID     uint32
	ThreadID      uint32
	CoreID        uint32
	SyscallNumber uint64
	RAX           uint64
	RCX           uint64
	RDX           uint64
	R8            uint64
	R9            uint64
	R10           uint64
	RSP           uint64
	Options       DebugEventOptions
}

type DebugEventOptions struct {
	OptionalParam1 uint64
	OptionalParam2 uint64
	OptionalParam3 uint64
	OptionalParam4 uint64
}
