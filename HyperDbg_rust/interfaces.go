package hyperdbgrust

import (
	"context"
	"time"
)

type Debugger interface {
	Start() error
	Stop()
	IsConnected() bool
	GetState() DebugState
	Ping() error
	Status() (string, error)
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
	GetProcessList() ([]ProcessInfo, error)
	GetThreadList(processID uint32) ([]ThreadInfo, error)
	GetModuleList(processID uint32) ([]ModuleInfo, error)
	RegisterCallback(msgType MessageType, cb EventCallback)
	GetEvent() any
	WaitForEvent(timeout time.Duration) *Message
	GetConnectedDrivers() []uint64
	WaitForDriver(timeout time.Duration) error
	ExecuteScript(script string) (string, error)
	ExecuteScriptWithContext(ctx context.Context, script string) (string, error)
}
