package sdk

// BreakpointClearByID
// Description:bc : clears a breakpoint using breakpoint id.
// Syntax:
// bc [BreakpointId (hex)]
// Examples:
// bc 0
// bc 2
func BreakpointClearByID() (status string) {
	return InterpreterEx("bc")
}

// BreakpointDisableByID
// Description:bd : disables a breakpoint using breakpoint id.
// Syntax:
// bd [BreakpointId (hex)]
// Examples:
// bd 0
// bd 2
func BreakpointDisableByID() (status string) {
	return InterpreterEx("bd")
}

// BreakpointEnableByID
// Description:be : enables a breakpoint using breakpoint id.
// Syntax:
// be [BreakpointId (hex)]
// Examples:
// be 0
// be 2
func BreakpointEnableByID() (status string) {
	return InterpreterEx("be")
}

// BreakpointList
// Description:bl : lists all the enabled and disabled breakpoints.
// Syntax:
// bl
func BreakpointList() (status string) {
	return InterpreterEx("bl")
}

// SetBreakpoint0xcc
// Description:bp : puts a breakpoint (0xcc).
// Syntax:
// bp [Address (hex)] [pid ProcessId (hex)] [tid ThreadId (hex)] [core CoreId (hex)]
// Examples:
// bp nt!ExAllocatePoolWithTag
// bp nt!ExAllocatePoolWithTag+5
// bp nt!ExAllocatePoolWithTag+@rcx+rbx
// bp fffff8077356f010
// bp fffff8077356f010 pid 0x4
// bp fffff8077356f010 tid 0x1000
// bp fffff8077356f010 pid 0x4 core 2
// Notes:
// 'bp' is not an event, if you want to use an event version
func SetBreakpoint0xcc() (status string) {
	return InterpreterEx("bp")
}

// CoreOperatingProcessorForShowsAndChanges
// Description:~ : shows and changes the operating processor.
// Syntax:
// ~
// ~ [CoreNumber (hex)]
// Examples:
// ~
// ~ 2
func CoreOperatingProcessorForShowsAndChanges() (status string) {
	return InterpreterEx("core")
}

// CpuFeaturesForCollectsReport
// Description:cpu : collects a report from cpu features.
// Syntax:
// cpu
// Examples:
// dt nt!_EPROCESS
// dt nt!_EPROCESS fffff8077356f010
// dt nt!_EPROCESS $proc
// dt nt!_KPROCESS @rax+@rbx+c0
// !dt nt!_EPROCESS 1f0300
// dt nt!_MY_STRUCT 7ff00040 pid 1420
// dt nt!_EPROCESS $proc inline all
// dt nt!_EPROCESS fffff8077356f010 inline no
func CpuFeaturesForCollectsReport() (status string) {
	return InterpreterEx("cpu")
}

// FlushKernelModeBuffers
// Description:removes all the buffer and messages from kernel-mode buffers.
// Syntax:
// flush
func FlushKernelModeBuffers() (status string) {
	return InterpreterEx("flush")
}

// ContinueDebugger
// Description:continues debuggee or continues processing kernel messages.
// Syntax:
// g
func ContinueDebugger() (status string) {
	return InterpreterEx("g")
}

// StepOut
// Description:executes a single instruction (step-out) and optionally displays the resulting values of all registers and flags.
// Syntax:
// gu
// gu [Count (hex)]
// Examples:
// gu
// gu 10000
func StepOut() (status string) {
	return InterpreterEx("gu")
}

// StepIn
// Description:executes a single instruction (step-in) and guarantees that no other instruction is executed other than the displayed instruction including user to the kernel (syscalls) and kernel to the user (sysrets) and exceptions and page-faults and optionally displays all registers and flags' resulting values.
// Syntax:
// i
// i [Count (hex)]
// ir
// ir [Count (hex)]
// Examples:
// i
// ir
// ir 1f
func StepIn() (status string) {
	return InterpreterEx("i")
}

// CallstackOrThreadView
// Description:shows the callstack of the thread.
// Syntax:
// k
// kd
// kq
// k [base StackAddress (hex)] [l Length (hex)]
// kd [base StackAddress (hex)] [l Length (hex)]
// kq [base StackAddress (hex)] [l Length (hex)]
// Examples:
// k
// k l 100
// kd base 0x77356f010
// kq base fffff8077356f010
// kq base @rbx-10
// kq base fffff8077356f010 l 100
func CallstackOrThreadView() (status string) {
	return InterpreterEx("k, kd, kq")
}

// ListModules
// Description:lists user/kernel modules' base address, size, name and path.
// Syntax:
// lm [m Name (string)] [pid ProcessId (hex)] [Filter (string)]
// Examples:
// lm
// lm km
// lm um
// lm m nt
// lm km m ntos
// lm um m kernel32
// lm um m kernel32 pid 1240
func ListModules() (status string) {
	return InterpreterEx("lm")
}

// LoadDriversAndModules
// Description:installs the drivers and load the modules.
// Syntax:
// load [ModuleName (string)]
// Examples:
// load vmm
func LoadDriversAndModules() (status string) {
	return InterpreterEx("load")
}

// OutputEventForwardingInstance
// Description:creates an output instance that can be used in event forwarding.
// Syntax:
// output [create Name (string)] [file|namedpipe|tcp Address (string)]
// output [open|close Name (string)]
// Examples:
// output create MyOutputName1 file c:\rev\output.txt
// output create MyOutputName2 tcp 192.168.1.10:8080
// output create MyOutputName3 namedpipe \\.\Pipe\HyperDbgOutput
// output create MyOutputName1 module c:\rev\event_forwarding.dll
// output open MyOutputName1
// output close MyOutputName1
func OutputEventForwardingInstance() (status string) {
	return InterpreterEx("output")
}

// StepExecuteSingleInstruction
// Description:executes a single instruction (step) and optionally displays the resulting values of all registers and flags.
// Syntax:
// p
// p [Count (hex)]
// pr
// pr [Count (hex)]
// Examples:
// p
// pr
// pr 1f
func StepExecuteSingleInstruction() (status string) {
	return InterpreterEx("p")
}

// PauseKernelEvents
// Description:pauses the kernel events.
// Syntax:
// pause
func PauseKernelEvents() (status string) {
	return InterpreterEx("pause")
}

// PreactivateSpecialFunctionality
// Description:preactivates a special functionality.
// Syntax:
// preactivate  [Type (string)]
// Examples:
// preactivate mode
// Notes:
// type of activations: mode: used for preactivation of the '!mode' event
func PreactivateSpecialFunctionality() (status string) {
	return InterpreterEx("preactivate")
}

// PreallocateBuffer
// Description:pre-allocates buffer for special purposes.
// Syntax:
// prealloc [Type (string)] [Count (hex)]
// Examples:
// prealloc thread-interception 8
// prealloc monitor 10
// prealloc epthook 5
// prealloc epthook2 3
// prealloc regular-event 12
// prealloc big-safe-buffer 1
// Notes:
// thread-interception: used for pre-allocations of the thread holders for the thread interception mechanism
// monitor: used for pre-allocations of the '!monitor' EPT hooks
// epthook: used for pre-allocations of the '!epthook' EPT hooks
// epthook2: used for pre-allocations of the '!epthook2' EPT hooks
// regular-event: used for pre-allocations of regular instant events
// big-event: used for pre-allocations of big instant events
// regular-safe-buffer: used for pre-allocations of the regular event safe buffers ($buffer) for instant events
// big-safe-buffer: used for pre-allocations of the big event safe buffers ($buffer) for instant events
func PreallocateBuffer() (status string) {
	return InterpreterEx("prealloc")
}

// PrintEvaluateExpressions
// Description:evaluates expressions.
// Syntax:
// print [Expression (string)]
// Examples:
// print dq(poi(@rcx))
func PrintEvaluateExpressions() (status string) {
	return InterpreterEx("print")
}

// RegistersReadOrModify
// Description:reads or modifies registers.
// Syntax:
// r
// r [Register (string)] [= Expr (string)]
// Examples:
// r
// r @rax
// r rax
// r rax = 0x55
// r rax = @rbx + @rcx + 0n10
func RegistersReadOrModify() (status string) {
	return InterpreterEx("r")
}

// ReadMsr
// Description:reads a model-specific register (MSR).
// Syntax:
// rdmsr [Msr (hex)] [core CoreNumber (hex)]
// Examples:
// rdmsr c0000082
// rdmsr c0000082 core 2
func ReadMsr() (status string) {
	return InterpreterEx("rdmsr")
}

// SearchMemoryPattern
// Description:searches a contiguous memory for a special byte pattern.
// Syntax:
// sb [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
// sd [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
// sq [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
// Examples:
// sb nt!ExAllocatePoolWithTag 90 85 95 l ffff
// sb nt!ExAllocatePoolWithTag+5 90 85 95 l ffff
// sb @rcx+5 90 85 95 l ffff
// sb fffff8077356f010 90 85 95 l ffff
// sd fffff8077356f010 90423580 l ffff pid 1c0
// !sq 100000 9090909090909090 l ffff
// !sq @rdx+r12 9090909090909090 l ffff
// !sq 100000 9090909090909090 9090909090909090 9090909090909090 l ffffff
// Notes:
// If you want to search in physical (address) memory then add '!' at the start of the command
func SearchMemoryPattern() (status string) {
	return InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

// SettingsManagement
// Description:queries, sets, or changes a value for a special settings option.
// Syntax:
// settings [OptionName (string)]
// settings [OptionName (string)] [Value (hex)]
// settings [OptionName (string)] [Value (string)]
// settings [OptionName (string)] [on|off]
// Examples:
// settings autounpause
// settings autounpause on
// settings autounpause off
// settings addressconversion on
// settings addressconversion off
// settings autoflush on
// settings autoflush off
// settings syntax intel
// settings syntax att
// settings syntax masm
func SettingsManagement() (status string) {
	return InterpreterEx("settings")
}

// SleepDebugger
// Description:sleeps the debugger; this command is used in scripts, it doesn't breaks the debugger but the debugger still shows the buffers received from kernel.
// Syntax:
// sleep [MillisecondsTime (hex)]
func SleepDebugger() (status string) {
	return InterpreterEx("sleep")
}

// StepInExecute
// Description:executes a single instruction (step-in) and optionally displays the resulting values of all registers and flags.
// Syntax:
// t
// t [Count (hex)]
// tr
// tr [Count (hex)]
// Examples:
// t
// tr
// tr 1f
func StepInExecute() (status string) {
	return InterpreterEx("t")
}

// TestHyperDbgFeatures
// Description:tests essential features of HyperDbg in current machine.
// Syntax:
// test [Task (string)]
// Examples:
// test
// test query
// test trap-status
// test pool
// test query
// test breakpoint on
// test breakpoint off
// test trap on
// test trap off
func TestHyperDbgFeatures() (status string) {
	return InterpreterEx("test")
}

// UnloadKernelModules
// Description:unloads the kernel modules and uninstalls the drivers.
// Syntax:
// unload [remove] [ModuleName (string)]
// Examples:
// unload vmm
// unload remove vmm
func UnloadKernelModules() (status string) {
	return InterpreterEx("unload")
}

// SearchSymbols
// Description:searches and shows symbols (wildcard) and corresponding addresses.
// Syntax:
// x [Module!Name (wildcard string)]
// Examples:
// x nt!ExAllocatePoolWithTag
// x nt!ExAllocatePool*
// x nt!*
func SearchSymbols() (status string) {
	return InterpreterEx("x")
}

// CpuidExecutionMonitor
// Description:monitors execution of a special cpuid index or all cpuids instructions.
// Syntax:
// !cpuid [Eax (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !cpuid
// !cpuid 1
// !cpuid pid 400
// !cpuid core 2 pid 400
func CpuidExecutionMonitor() (status string) {
	return InterpreterEx("!cpuid")
}

// ControlRegisterModificationMonitor
// Description:monitors modification of control registers (CR0 / CR4).
// Syntax:
// !crwrite [Cr (hex)] [mask Mask (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !crwrite 0
// !crwrite 0 0x10000
// !crwrite 4 pid 400
// !crwrite 4 core 2 pid 400
func ControlRegisterModificationMonitor() (status string) {
	return InterpreterEx("!crwrite")
}

// DebugRegistersMonitor
// Description:monitors any access to debug registers.
// Syntax:
// !dr [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !dr
// !dr pid 400
// !dr core 2 pid 400
func DebugRegistersMonitor() (status string) {
	return InterpreterEx("!dr")
}

// EptHook
// Description:puts a hidden-hook EPT (hidden breakpoints).
// Syntax:
// !epthook [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !epthook nt!ExAllocatePoolWithTag
// !epthook nt!ExAllocatePoolWithTag+5
// !epthook fffff801deadb000
// !epthook fffff801deadb000 pid 400
// !epthook fffff801deadb000 core 2 pid 400
func EptHook() (status string) {
	return InterpreterEx("!epthook")
}

// EptHook2
// Description:puts a hidden-hook EPT (detours).
// Syntax:
// !epthook2 [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !epthook2 nt!ExAllocatePoolWithTag
// !epthook2 nt!ExAllocatePoolWithTag+5
// !epthook2 fffff801deadb000
// !epthook2 fffff801deadb000 pid 400
// !epthook2 fffff801deadb000 core 2 pid 400
func EptHook2() (status string) {
	return InterpreterEx("!epthook2")
}

// MonitorIDTEntries
// Description:monitors the first 32 entry of IDT (starting from zero).
// Syntax:
// !exception [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !exception
// !exception 0xe
// !exception pid 400
// !exception core 2 pid 400
// Notes:
// monitoring page-faults (entry 0xe) is implemented differently.
func MonitorIDTEntries() (status string) {
	return InterpreterEx("!exception")
}

// HideHyperDbg
// Description:tries to make HyperDbg transparent from anti-debugging and anti-hypervisor methods.
// Syntax:
// !hide
// !hide [pid ProcessId (hex)]
// !hide [name ProcessName (string)]
// Examples:
// !hide
// !hide pid b60
// !hide name procexp.exe
// Notes:
// process names are case sensitive and you can use this command multiple times.
func HideHyperDbg() (status string) {
	return InterpreterEx("!hide")
}

// InterruptExternalMonitor
// Description:monitors the external interrupt (IDT >= 32).
// Syntax:
// [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !interrupt 0x2f
// !interrupt 0x2f pid 400
// !interrupt 0x2f core 2 pid 400
// Notes:
// The index should be greater than 0x20 (32) and less than 0xFF (255) - starting from zero.
func InterruptExternalMonitor() (status string) {
	return InterpreterEx("!interrupt")
}

// DetectIOInstructionsIn
// Description:detects the execution of IN (I/O instructions) instructions.
// Syntax:
// !ioin [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !ioin
// !ioin 0x64
// !ioin pid 400
// !ioin core 2 pid 400
func DetectIOInstructionsIn() (status string) {
	return InterpreterEx("!ioin")
}

// DetectIOInstructionsOut
// Description:detects the execution of OUT (I/O instructions) instructions.
// Syntax:
// !ioout [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !ioout
// !ioout 0x64
// !ioout pid 400
// !ioout core 2 pid 400
func DetectIOInstructionsOut() (status string) {
	return InterpreterEx("!ioout")
}

// MeasureArgumentsForHide
// Description:measures the arguments needs for the '!hide' command.
// Syntax:
// !measure [default]
// Examples:
// !measure
// !measure default
func MeasureArgumentsForHide() (status string) {
	return InterpreterEx("!measure")
}

// ModeInstructionsTrap
// Description:traps (and possibly blocks) the execution of user-mode/kernel-mode instructions.
// Syntax:
// !mode [Mode (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !mode u pid 1c0
// !mode k pid 1c0
// !mode ku pid 1c0
// !mode ku core 2 pid 400
// Notes:
// this event applies to the target process; thus, you need to specify the process id
func ModeInstructionsTrap() (status string) {
	return InterpreterEx("!mode")
}

// MonitorMemoryAccess
// Description:monitors address range for read and writes.
// Syntax:
// !monitor [MemoryType (vapa)] [Attribute (string)] [FromAddress (hex)] [ToAddress (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// !monitor [MemoryType (vapa)] [Attribute (string)] [FromAddress (hex)] [l Length (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !monitor rw fffff801deadb000 fffff801deadbfff
// !monitor rw fffff801deadb000 l 1000
// !monitor pa rw c01000 l 1000
// !monitor rwx fffff801deadb000 fffff801deadbfff
// !monitor rwx fffff801deadb000 l 230d0
// !monitor rw nt!Kd_DEFAULT_Mask Kd_DEFAULT_Mask+5
// !monitor r fffff801deadb000 fffff801deadbfff pid 400
// !monitor w fffff801deadb000 fffff801deadbfff core 2 pid 400
// !monitor w c01000 c01000+2500 core 2 pid 400
// !monitor x fffff801deadb000 fffff801deadbfff core 2 pid 400
// !monitor x fffff801deadb000 l 500 core 2 pid 400
// !monitor wx fffff801deadb000 fffff801deadbfff core 2 pid 400
func MonitorMemoryAccess() (status string) {
	return InterpreterEx("!monitor")
}

// MsrRead
// Description:detects the execution of rdmsr instructions.
// Syntax:
// !msrread [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !msrread
// !msrread 0xc0000082
// !msread pid 400
// !msrread core 2 pid 400
func MsrRead() (status string) {
	return InterpreterEx("!msrread")
}

// MsrWrite
// Description:detects the execution of wrmsr instructions.
// Syntax:
// !msrwrite [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !msrwrite
// !msrwrite 0xc0000082
// !msrwrite pid 400
// !msrwrite core 2 pid 400
func MsrWrite() (status string) {
	return InterpreterEx("!msrwrite")
}

// Pa2Va
// Description:converts virtual address to physical address.
// Syntax:
// !pa2va [PhysicalAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !pa2va nt!ExAllocatePoolWithTag
// !pa2va nt!ExAllocatePoolWithTag+5
// !pa2va @rax+5
// !pa2va fffff801deadbeef
// !pa2va fffff801deadbeef pid 0xc8
func Pa2Va() (status string) {
	return InterpreterEx("!pa2va")
}

// PmcExecutionMonitor
// Description:monitors execution of rdpmc instructions.
// Syntax:
// !pmc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !pmc
// !pmc pid 400
// !pmc core 2 pid 400
func PmcExecutionMonitor() (status string) {
	return InterpreterEx("!pmc")
}

// Pte
// Description:finds virtual addresses of different paging-levels.
// Syntax:
// !pte [VirtualAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !pte nt!ExAllocatePoolWithTag
// !pte nt!ExAllocatePoolWithTag+5
// !pte fffff801deadbeef
// !pte 0x400021000 pid 1c0
func Pte() (status string) {
	return InterpreterEx("!pte")
}

// ReversingMachineModuleUse
// Description:uses the reversing machine module in order to reconstruct the programmer/memory assumptions.
// Syntax:
// !rev [config] [pid ProcessId (hex)]
// !rev [path Path (string)] [Parameters (string)]
// Examples:
// !rev path c:\reverse eng\my_file.exe
// !rev pattern
// !rev reconstruct
// !rev pattern pid 1c0
// !rev reconstruct pid 1c0
func ReversingMachineModuleUse() (status string) {
	return InterpreterEx("!rev")
}

// Syscall
// Description:monitors and hooks all execution of syscall instructions (by accessing memory and checking for instructions).
// Syntax:
// !syscall [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// !syscall2 [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !syscall
// !syscall2
// !syscall 0x55
// !syscall2 0x55
// !syscall 0x55 pid 400
// !syscall 0x55 core 2 pid 400
// !syscall2 0x55 core 2 pid 400
func Syscall() (status string) {
	return InterpreterEx("!syscall")
}

// SysRet
// Description:monitors and hooks all execution of sysret instructions (by accessing memory and checking for instructions).
// Syntax:
// !sysret [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]
// Examples:
// !sysret
// !sysret2
// !sysret pid 400
// !sysret2 pid 400
// !sysret core 2 pid 400
// !sysret2 core 2 pid 400
func SysRet() (status string) {
	return InterpreterEx("!sysret")
}

// TraceExecution
// Description:traces the execution of user-mode/kernel-mode instructions.
// Syntax:
// !trace [TraceType (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !trace step-out
// !trace step-in pid 1c0
// !trace instrument-step core 2 pid 400
// Notes:
// valid trace types:
// step-in : single step-in (regular)
// step-out : single step-out (regular)
// instrument-step : single step-in (instrumentation)
func TraceExecution() (status string) {
	return InterpreterEx("!trace")
}

// TrackModeTransitionInstructions
// Description:tracks instructions from user-mode to kernel-mode or kernel-mode to user-mode to create call tree. Please note that it's highly recommended to configure symbols before using this command as it maps addresses to corresponding function names.
// Syntax:
// !track [tree] [Count (hex)]
// !track [reg] [Count (hex)]
// Examples:
// !track tree 10000
// !track reg 10000
func TrackModeTransitionInstructions() (status string) {
	return InterpreterEx("!track")
}

// TscInstructionsMonitor
// Description:monitors execution of rdtsc/rdtscp instructions.
// Syntax:
// !tsc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !tsc
// !tsc pid 400
// !tsc core 2 pid 400
func TscInstructionsMonitor() (status string) {
	return InterpreterEx("!tsc")
}

// UnHide
// Description:reverts the transparency measures of the '!hide' command.
// Syntax:
// !unhide
// Examples:
// !unhide
func UnHide() (status string) {
	return InterpreterEx("!unhide")
}

// Va2Pa
// Description:converts virtual address to physical address.
// Syntax:
// !va2pa [VirtualAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !va2pa nt!ExAllocatePoolWithTag
// !va2pa nt!ExAllocatePoolWithTag+5
// !va2pa @rcx
// !va2pa @rcx+5
// !va2pa fffff801deadbeef
// !va2pa fffff801deadbeef pid 0xc8
func Va2Pa() (status string) {
	return InterpreterEx("!va2pa")
}

// VmCall
// Description:monitors execution of VMCALL instruction.
// Syntax:
// !vmcall [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !vmcall
// !vmcall pid 400
// !vmcall core 2 pid 400
func VmCall() (status string) {
	return InterpreterEx("!vmcall")
}

// HardwareClockDebugging
// Description:performs actions related to hwdbg hardware debugging events for each clock cycle.
// Syntax:
// !hw_clk  [script { Script (string) }]
// Examples:
// !hw_clk script { @hw_pin1 = 0; }
func HardwareClockDebugging() (status string) {
	return InterpreterEx("!hw_clk")
}

// AttachDebugThread
// Description:attaches to debug a thread in VMI Mode.
// Syntax:
// .attach [pid ProcessId (hex)]
// Examples:
// .attach pid b60
func AttachDebugThread() (status string) {
	return InterpreterEx(".attach")
}

// ClearScreen
// Description:clears the screen.
// Syntax:
// .cls
func ClearScreen() (status string) {
	return InterpreterEx(".cls")
}

// ConnectToMachine
// Description:connects to a remote or local machine to start debugging.
// Syntax:
// .connect [local]
// .connect [Ip (string)] [Port (decimal)]
// Examples:
// .connect local
// .connect 192.168.1.5 50000
func ConnectToMachine() (status string) {
	return InterpreterEx(".connect")
}

// DebugMachine
// Description:debugs a target machine or makes this machine a debuggee.
// Syntax:
// .debug [remote] [serial|namedpipe] [Baudrate (decimal)] [Address (string)]
// .debug [prepare] [serial] [Baudrate (decimal)] [Address (string)]
// .debug [close]
// Examples:
// .debug remote serial 115200 com2
// .debug remote namedpipe \\\.\pipe\HyperDbgPipe
// .debug prepare serial 115200 com1
// .debug prepare serial 115200 com2
// .debug close
// Notes:
// valid baud rates (decimal): 110, 300, 600, 1200, 2400, 4800, 9600, 14400, 19200, 38400, 56000, 57600, 115200, 128000, 256000
// valid COM ports: COM1, COM2, COM3, COM4
func DebugMachine() (status string) {
	return InterpreterEx(".debug")
}

// DetachDebugging
// Description:detaches from debugging a user-mode process.
// Syntax:
// .detach
func DetachDebugging() (status string) {
	return InterpreterEx(".detach")
}

// DisconnectSession
// Description:disconnects from a debugging session (it won't unload the modules).
// Syntax:
// .disconnect
func DisconnectSession() (status string) {
	return InterpreterEx(".disconnect")
}

// DumpMemoryContext
// Description:saves memory context into a file.
// Syntax:
// .dump [FromAddress (hex)] [ToAddress (hex)] [pid ProcessId (hex)] [path Path (string)]
// Examples:
// .dump 401000 40b000 path c:\\rev\\dump1.dmp
// .dump 401000 40b000 pid 1c0 path c:\\rev\\desktop\\dump2.dmp
// .dump fffff801deadb000 fffff801deade054 path c:\\rev\\dump3.dmp
// .dump fffff801deadb000 fffff801deade054 path c:\\rev\\dump4.dmp
// .dump 00007ff8349f2000 00007ff8349f8000 path c:\\rev\\dump5.dmp
// .dump @rax+@rcx @rax+@rcx+1000 path c:\\rev\\dump6.dmp
// !dump 1000 2100 path c:\\rev\\dump7.dmp
// Notes:
// If you want to dump physical memory then add '!' at the start of the command
func DumpMemoryContext() (status string) {
	return InterpreterEx(".dump")
}

// FormatsDiff
// Description:shows a value or register in different formats
// Syntax:
// .formats [Expression (string)]
// Examples:
// .formats nt!ExAllocatePoolWithTag
// .formats nt!Kd_DEFAULT_Mask
// .formats nt!Kd_DEFAULT_Mask+5
// .formats 55
// .formats @rax
// .formats @rbx+@rcx
// .formats $pid
func FormatsDiff() (status string) {
	return InterpreterEx(".formats")
}

// HelpForCommand
// Description:shows help and example(s) of a specific command
// Syntax:
// .help [Command (string)]
// Examples:
// .help !monitor
func HelpForCommand() (status string) {
	return InterpreterEx(".help")
}

// KillProcess
// Description:terminates the current running process
// Syntax:
// .kill
func KillProcess() (status string) {
	return InterpreterEx(".kill")
}

// ListenForClientConnection
// Description:listens for a client to connect to HyperDbg (works as a guest server)
// Syntax:
// .listen [Port (decimal)]
// Examples:
// .listen
// .listen 50000
// Notes:
// if you don't specify port then HyperDbg uses the default port
func ListenForClientConnection() (status string) {
	return InterpreterEx(".listen")
}

// LogClose
// Description:closes the previously opened log
// Syntax:
// .logclose
func LogClose() (status string) {
	return InterpreterEx(".logclose")
}

// LogOpen
// Description:saves commands and results in a file
// Syntax:
// .logopen [FilePath (string)]
func LogOpen() (status string) {
	return InterpreterEx(".logopen")
}

// MakePageAvailableInRAM
// Description:brings the page in, making it available in the RAM
// Syntax:
// .pagein [Mode (string)] [l Length (hex)]
// .pagein [Mode (string)] [VirtualAddress (hex)] [l Length (hex)]
// Examples:
// .pagein fffff801deadbeef
// .pagein 00007ff8349f2224 l 1a000
// .pagein u 00007ff8349f2224
// .pagein w 00007ff8349f2224
// .pagein f 00007ff8349f2224
// .pagein pw 00007ff8349f2224
// .pagein wu 00007ff8349f2224
// .pagein wu 00007ff8349f2224 l 6000
// .pagein pf @rax
// .pagein uf @rip+@rcx
// .pagein pwu @rax+5
// .pagein pwu @rax l 2000
// Notes:
// valid mode formats: present (p), write (w), user (u), fetch (f), protection key (k), shadow stack (s), hlat (h), sgx (g)
// common page-fault codes: 0x0 (default), 0x2 (w), 0x3 (pw), 0x4 (u), 0x6 (wu), 0x7 (pwu), 0x10 (f), 0x11 (pf), 0x14 (uf)
func MakePageAvailableInRAM() (status string) {
	return InterpreterEx(".pagein")
}

// ParseExecutableFiles
// Description:parses portable executable (PE) files and dump sections
// Syntax:
// .pe [header] [FilePath (string)]
// .pe [section] [SectionName (string)] [FilePath (string)]
// Examples:
// .pe header c:\reverse files\myfile.exe
// .pe section .text c:\reverse files\myfile.exe
// .pe section .rdata c:\reverse files\myfile.exe
func ParseExecutableFiles() (status string) {
	return InterpreterEx(".pe")
}

// ProcessesView
// Description:shows and changes the processes. This command needs public symbols for ntoskrnl.exe if you want to see the processes list. Please visit the documentation to know about the difference between '.process' and '.process2'.
// Syntax:
// .process
// .process [list]
// .process [pid ProcessId (hex)]
// .process [process Eprocess (hex)]
// .process2 [pid ProcessId (hex)]
// .process2 [process Eprocess (hex)]
// Examples:
// .process
// .process list
// .process pid 4
// .process2 pid 4
// .process process ffff948cc2349280
func ProcessesView() (status string) {
	return InterpreterEx(".process")
}

// RestartProcess
// Description:restarts the previously started process using '.start' command.
// Syntax:
// .restart
func RestartProcess() (status string) {
	return InterpreterEx(".restart")
}

// Script
// Description:runs a HyperDbg script.
// Syntax:
// .script [FilePath (string)] [Args (string)]
// Examples:
// .script C:\scripts\script.ds
// .script C:\scripts\script.ds 95 85 @rsp
// .script "C:\scripts\hello world.ds"
// .script "C:\scripts\hello world.ds" @rax
// .script "C:\scripts\hello world.ds" @rax @rcx+55 $pid
// .script "C:\scripts\hello world.ds" 12 55 @rip
func Script() (status string) {
	return InterpreterEx(".script")
}

// RunsAUser_modeProcess
// Description:runs a user-mode process
// Syntax:
// .start [path Path (string)] [Parameters (string)]
// Examples:
// .start path c:\reverse eng\my_file.exe
func RunsAUser_modeProcess() (status string) {
	return InterpreterEx(".start")
}

// Status
// Description:gets the status of current debugger in local system
// Syntax:
// .status
// status
// Notes:
// If connected to a remote system, '.status' shows the state of current debugger, while 'status' shows the state of remote debuggee.
func Status() (status string) {
	return InterpreterEx(".status")
}

// SwitchThread
// Description:shows the list of active debugging threads and switches between processes and threads in VMI Mode
// Syntax:
// .switch
// .switch [pid ProcessId (hex)]
// .switch [tid ThreadId (hex)]
// Examples:
// .switch list
// .switch pid b60
// .switch tid b60
func SwitchThread() (status string) {
	return InterpreterEx(".switch")
}

// SymbolActions
// Description:performs the symbol actions
// Syntax:
// .sym [table]
// .sym [reload] [pid ProcessId (hex)]
// .sym [download]
// .sym [load]
// .sym [unload]
// .sym [add] [base Address (hex)] [path Path (string)]
// Examples:
// .sym table
// .sym reload
// .sym load
// .sym download
// .sym add base fffff8077356000 path c:\symbols\my_dll.pdb
// .sym unload
func SymbolActions() (status string) {
	return InterpreterEx(".sym")
}

// SymbolPath
// Description:shows and sets the symbol server and path
// Syntax:
// .sympath
// .sympath [SymServer (string)]
// Examples:
// .sympath
// .sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbols
func SymbolPath() (status string) {
	return InterpreterEx(".sympath")
}

// Thread
// Description:shows and changes the threads
// Syntax:
// .thread
// .thread [list] [process Eprocess (hex)]
// .thread [tid ThreadId (hex)]
// .thread [thread Ethread (hex)]
// .thread2 [tid ThreadId (hex)]
// .thread2 [thread Ethread (hex)]
// Examples:
// .thread
// .thread tid 48a4
// .thread2 tid 48a4
// .thread thread ffff948c`c8970200
// .thread list
// .thread list process ffff948c`a1279880
// Notes:
// This command needs public symbols for 'ntoskrnl.exe' if you want to see the threads list. Please visit the documentation to know about the difference between '.thread' and '.thread2'.
func Thread() (status string) {
	return InterpreterEx(".thread")
}
