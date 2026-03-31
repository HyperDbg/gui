package debugger

import (
	"context"
	"time"
)

type Debugger interface {
	Initialize() error
	Terminate() error
	StartNetworkServer(addr string) error
	StopNetworkServer() error
	KernelDebuggerInitialize() error
	KernelDebuggerUninitialize() error
	Connect() error
	Disconnect()
	IsConnected() bool
	GetState() DebugState
	Status() (string, error)
	Ping() error
	LoadVmm() error
	UnloadVmm() error
	AttachProcess(processID uint32) error
	DetachProcess() error
	StartProcess(exePath string) (uint32, error)
	KillProcess(processID uint32) error
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
	RegisterCpuidCallback(cb EventCallback)
	GetEvent() any
	WaitForEvent(timeout time.Duration) *Message
	ExecuteScript(script string) (string, error)
	ExecuteScriptWithContext(ctx context.Context, script string) (string, error)
	Disassemble(address uint64, bytes []byte, maxInstructions uint32) ([]Instruction, error)
	LoadSymbols(pdbPath string) error
	UnloadSymbols()
	GetSymbolByName(name string) (SymbolInfo, error)
	GetSymbolByAddress(address uint64) (SymbolInfo, error)
	GetFunctionByAddress(address uint64) (FunctionInfo, error)
	InstallHookScript(script *HookScript) error
}
