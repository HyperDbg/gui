package debugger

import (
	"context"
	"time"
)

type Debugger interface {
	Start(port int) error
	Stop()
	WaitForDriver(timeout time.Duration) error
	Initialize() error
	Terminate() error
	AttachProcess(processID uint32) error
	DetachProcess() error
	SetBreakpoint(address uint64, bpType BreakpointType) error
	RemoveBreakpoint(breakpointID uint64) error
	Continue() error
	Pause() error
	StepInto() error
	StepOver() error
	StepOut() error
	ReadMemory(address uint64, size uint32) ([]byte, error)
	WriteMemory(address uint64, data []byte) error
	ReadRegisters() (*RegisterState, error)
	WriteRegisters(regs *RegisterState) error
	GetState() DebugState
	WaitForEvent(timeout time.Duration) *Message
	ExecuteScript(script string) (string, error)
	ExecuteScriptWithContext(ctx context.Context, script string) (string, error)
}

type BreakpointType int

const (
	BreakpointTypeSoftware BreakpointType = iota
	BreakpointTypeHardware
	BreakpointTypeEPT
)

type DebugState int

const (
	StateTerminated DebugState = iota
	StateRunning
	StatePaused
	StateStepping
)

type RegisterState struct {
	RAX    uint64
	RBX    uint64
	RCX    uint64
	RDX    uint64
	RSI    uint64
	RDI    uint64
	RBP    uint64
	RSP    uint64
	R8     uint64
	R9     uint64
	R10    uint64
	R11    uint64
	R12    uint64
	R13    uint64
	R14    uint64
	R15    uint64
	RIP    uint64
	RFLAGS uint64
}

type Message struct {
	Type    uint32
	Payload []byte
}

type BreakpointEvent struct {
	ProcessID    uint32
	ThreadID     uint32
	CoreID       uint32
	Address      uint64
	BreakpointID uint64
}

type ExceptionEvent struct {
	ProcessID     uint32
	ThreadID      uint32
	CoreID        uint32
	ExceptionCode uint32
	Address       uint64
	ErrorCode     uint64
}

type DebugPrintEvent struct {
	ProcessID uint32
	ThreadID  uint32
	CoreID    uint32
	Level     uint32
	Message   string
}
