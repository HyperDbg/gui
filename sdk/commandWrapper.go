package sdk

// BreakpointClearByID
// Syntax:
// bc [BreakpointId (hex)]
// Examples:
// bc 0
// bc 2
func BreakpointClearByID() { // bc : clears a breakpoint using breakpoint id.
	InterpreterEx("bc")
}

// BreakpointDisableByID
// Syntax:
// bd [BreakpointId (hex)]
// Examples:
// bd 0
// bd 2
func BreakpointDisableByID() { // bd : disables a breakpoint using breakpoint id.
	InterpreterEx("bd")
}

// BreakpointEnableByID
// Syntax:
// be [BreakpointId (hex)]
// Examples:
// be 0
// be 2
func BreakpointEnableByID() { // be : enables a breakpoint using breakpoint id.
	InterpreterEx("be")
}

// BreakpointList
// Syntax:
// bl
// Examples:
func BreakpointList() { // bl : lists all the enabled and disabled breakpoints.
	InterpreterEx("bl")
}

// SetBreakpoint0xcc
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
func SetBreakpoint0xcc() { // bp : puts a breakpoint (0xcc).

	InterpreterEx("bp")
}

// CoreOperatingProcessorForShowsAndChanges
// Syntax:
// ~
// ~ [CoreNumber (hex)]
// Examples:
// ~
// ~ 2
func CoreOperatingProcessorForShowsAndChanges() { //~ : shows and changes the operating processor.
	InterpreterEx("core")
}

// CpuFeaturesForCollectsReport
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
func CpuFeaturesForCollectsReport() { // cpu : collects a report from cpu features.
	InterpreterEx("cpu")
}

// FlushKernelModeBuffers
// Syntax:
// flush
// Examples:
func FlushKernelModeBuffers() { // removes all the buffer and messages from kernel-mode buffers.
	InterpreterEx("flush")
}

// ContinueDebugger
// Syntax:
// g
// Examples:
func ContinueDebugger() { // continues debuggee or continues processing kernel messages.
	InterpreterEx("g")
}

// StepOut
// Syntax:
// gu
// gu [Count (hex)]
// Examples:
// gu
// gu 10000
func StepOut() { // executes a single instruction (step-out) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("gu")
}

// StepIn
// Syntax:
// i
// i [Count (hex)]
// ir
// ir [Count (hex)]
// Examples:
// i
// ir
// ir 1f
func StepIn() { // executes a single instruction (step-in) and guarantees that no other instruction is executed other than the displayed instruction including user to the kernel (syscalls) and kernel to the user (sysrets) and exceptions and page-faults and optionally displays all registers and flags' resulting values.
	InterpreterEx("i")
}

// CallstackOrThreadView
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
func CallstackOrThreadView() { // shows the callstack of the thread.
	InterpreterEx("k, kd, kq")
}

// ListModules
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
func ListModules() { // lists user/kernel modules' base address, size, name and path.
	InterpreterEx("lm")
}

// LoadDriversAndModules
// Syntax:
// load [ModuleName (string)]
// Examples:
// load vmm
func LoadDriversAndModules() { // installs the drivers and load the modules.
	InterpreterEx("load")
}

// OutputEventForwardingInstance
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
func OutputEventForwardingInstance() { // creates an output instance that can be used in event forwarding.
	InterpreterEx("output")
}

// StepExecuteSingleInstruction
// Syntax:
// p
// p [Count (hex)]
// pr
// pr [Count (hex)]
// Examples:
// p
// pr
// pr 1f
func StepExecuteSingleInstruction() { // executes a single instruction (step) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("p")
}

// PauseKernelEvents
// Syntax:
// pause
// Examples:
func PauseKernelEvents() { // pauses the kernel events.
	InterpreterEx("pause")
}

// PreactivateSpecialFunctionality
// Syntax:
// preactivate  [Type (string)]
// Examples:
// preactivate mode
// Notes:
// type of activations: mode: used for preactivation of the '!mode' event
func PreactivateSpecialFunctionality() { // preactivates a special functionality.
	InterpreterEx("preactivate")
}

// PreallocateBuffer
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
func PreallocateBuffer() { // pre-allocates buffer for special purposes.
	InterpreterEx("prealloc")
}

// PrintEvaluateExpressions
// Syntax:
// print [Expression (string)]
// Examples:
// print dq(poi(@rcx))
func PrintEvaluateExpressions() { // evaluates expressions.
	InterpreterEx("print")
}

// RegistersReadOrModify
// Syntax:
// r
// r [Register (string)] [= Expr (string)]
// Examples:
// r
// r @rax
// r rax
// r rax = 0x55
// r rax = @rbx + @rcx + 0n10
func RegistersReadOrModify() { // reads or modifies registers.
	InterpreterEx("r")
}

// ReadMsr
// Syntax:
// rdmsr [Msr (hex)] [core CoreNumber (hex)]
// Examples:
// rdmsr c0000082
// rdmsr c0000082 core 2
func ReadMsr() { // reads a model-specific register (MSR).
	InterpreterEx("rdmsr")
}

// SearchMemoryPattern
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
func SearchMemoryPattern() { // searches a contiguous memory for a special byte pattern.
	InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

// SettingsManagement
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
func SettingsManagement() { // queries, sets, or changes a value for a special settings option.
	InterpreterEx("settings")
}

// SleepDebugger
// Syntax:
// sleep [MillisecondsTime (hex)]
// Examples:
func SleepDebugger() { // sleeps the debugger; this command is used in scripts, it doesn't breaks the debugger but the debugger still shows the buffers received from kernel.
	InterpreterEx("sleep")
}

// StepInExecute
// Syntax:
// t
// t [Count (hex)]
// tr
// tr [Count (hex)]
// Examples:
// t
// tr
// tr 1f
func StepInExecute() { // executes a single instruction (step-in) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("t")
}

// TestHyperDbgFeatures
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
func TestHyperDbgFeatures() { // tests essential features of HyperDbg in current machine.
	InterpreterEx("test")
}

// UnloadKernelModules
// Syntax:
// unload [remove] [ModuleName (string)]
// Examples:
// unload vmm
// unload remove vmm
func UnloadKernelModules() { // unloads the kernel modules and uninstalls the drivers.
	InterpreterEx("unload")
}

// SearchSymbols
// Syntax:
// x [Module!Name (wildcard string)]
// Examples:
// x nt!ExAllocatePoolWithTag
// x nt!ExAllocatePool*
// x nt!*
func SearchSymbols() { // searches and shows symbols (wildcard) and corresponding addresses.
	InterpreterEx("x")
}

// CpuidExecutionMonitor
// Syntax:
// !cpuid [Eax (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !cpuid
// !cpuid 1
// !cpuid pid 400
// !cpuid core 2 pid 400
func CpuidExecutionMonitor() { // monitors execution of a special cpuid index or all cpuids instructions.
	InterpreterEx("!cpuid")
}

// ControlRegisterModificationMonitor
// Syntax:
// !crwrite [Cr (hex)] [mask Mask (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !crwrite 0
// !crwrite 0 0x10000
// !crwrite 4 pid 400
// !crwrite 4 core 2 pid 400
func ControlRegisterModificationMonitor() { // monitors modification of control registers (CR0 / CR4).
	InterpreterEx("!crwrite")
}

// DebugRegistersMonitor
// Syntax:
// !dr [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !dr
// !dr pid 400
// !dr core 2 pid 400
func DebugRegistersMonitor() { // monitors any access to debug registers.
	InterpreterEx("!dr")
}

// EptHook
// Syntax:
// !epthook [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !epthook nt!ExAllocatePoolWithTag
// !epthook nt!ExAllocatePoolWithTag+5
// !epthook fffff801deadb000
// !epthook fffff801deadb000 pid 400
// !epthook fffff801deadb000 core 2 pid 400
func EptHook() { // puts a hidden-hook EPT (hidden breakpoints).
	InterpreterEx("!epthook")
}

// EptHook2
// Syntax:
// !epthook2 [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !epthook2 nt!ExAllocatePoolWithTag
// !epthook2 nt!ExAllocatePoolWithTag+5
// !epthook2 fffff801deadb000
// !epthook2 fffff801deadb000 pid 400
// !epthook2 fffff801deadb000 core 2 pid 400
func EptHook2() { // puts a hidden-hook EPT (detours).
	InterpreterEx("!epthook2")
}

// IdtEntriesMonitor
// Syntax:
// !exception [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !exception
// !exception 0xe
// !exception pid 400
// !exception core 2 pid 400
// Notes:
// monitoring page-faults (entry 0xe) is implemented differently.
func IdtEntriesMonitor() { // monitors the first 32 entry of IDT (starting from zero).
	InterpreterEx("!exception")
}

// HideHyperDbg
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
func HideHyperDbg() { // tries to make HyperDbg transparent from anti-debugging and anti-hypervisor methods.
	InterpreterEx("!hide")
}

// InterruptExternalMonitor
// Syntax:
// [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !interrupt 0x2f
// !interrupt 0x2f pid 400
// !interrupt 0x2f core 2 pid 400
// Notes:
// The index should be greater than 0x20 (32) and less than 0xFF (255) - starting from zero.
func InterruptExternalMonitor() { // monitors the external interrupt (IDT >= 32).
	InterpreterEx("!interrupt")
}

// IoInDetect
// Syntax:
// !ioin [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !ioin
// !ioin 0x64
// !ioin pid 400
// !ioin core 2 pid 400
func IoInDetect() { // detects the execution of IN (I/O instructions) instructions.
	InterpreterEx("!ioin")
}

// IoOutDetect
// Syntax:
// !ioout [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !ioout
// !ioout 0x64
// !ioout pid 400
// !ioout core 2 pid 400
func IoOutDetect() { // detects the execution of OUT (I/O instructions) instructions.
	InterpreterEx("!ioout")
}

// MeasureArgumentsForHide
// Syntax:
// !measure [default]
// Examples:
// !measure
// !measure default
func MeasureArgumentsForHide() { // measures the arguments needs for the '!hide' command.
	InterpreterEx("!measure")
}

// ModeInstructionsTrap
// Syntax:
// !mode [Mode (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !mode u pid 1c0
// !mode k pid 1c0
// !mode ku pid 1c0
// !mode ku core 2 pid 400
// Notes:
// this event applies to the target process; thus, you need to specify the process id
func ModeInstructionsTrap() { // traps (and possibly blocks) the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!mode")
}

// MonitorMemoryAccess
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
func MonitorMemoryAccess() { // monitors address range for read and writes.
	InterpreterEx("!monitor")
}

// MsrRead
// Syntax:
// !msrread [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !msrread
// !msrread 0xc0000082
// !msread pid 400
// !msrread core 2 pid 400
func MsrRead() { // detects the execution of rdmsr instructions.
	InterpreterEx("!msrread")
}

// MsrWrite
// Syntax:
// !msrwrite [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !msrwrite
// !msrwrite 0xc0000082
// !msrwrite pid 400
// !msrwrite core 2 pid 400
func MsrWrite() { // detects the execution of wrmsr instructions.
	InterpreterEx("!msrwrite")
}

// Pa2Va
// Syntax:
// !pa2va [PhysicalAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !pa2va nt!ExAllocatePoolWithTag
// !pa2va nt!ExAllocatePoolWithTag+5
// !pa2va @rax+5
// !pa2va fffff801deadbeef
// !pa2va fffff801deadbeef pid 0xc8
func Pa2Va() { // converts virtual address to physical address.
	InterpreterEx("!pa2va")
}

// PmcExecutionMonitor
// Syntax:
// !pmc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !pmc
// !pmc pid 400
// !pmc core 2 pid 400
func PmcExecutionMonitor() { // monitors execution of rdpmc instructions.
	InterpreterEx("!pmc")
}

// Pte
// Syntax:
// !pte [VirtualAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !pte nt!ExAllocatePoolWithTag
// !pte nt!ExAllocatePoolWithTag+5
// !pte fffff801deadbeef
// !pte 0x400021000 pid 1c0
func Pte() { // finds virtual addresses of different paging-levels.
	InterpreterEx("!pte")
}

// ReversingMachineModuleUse
// Syntax:
// !rev [config] [pid ProcessId (hex)]
// !rev [path Path (string)] [Parameters (string)]
// Examples:
// !rev path c:\reverse eng\my_file.exe
// !rev pattern
// !rev reconstruct
// !rev pattern pid 1c0
// !rev reconstruct pid 1c0
func ReversingMachineModuleUse() { // uses the reversing machine module in order to reconstruct the programmer/memory assumptions.
	InterpreterEx("!rev")
}

// Syscall
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
func Syscall() { // monitors and hooks all execution of syscall instructions (by accessing memory and checking for instructions).
	InterpreterEx("!syscall")
}

// SysRet
// Syntax:
// !sysret [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]
// Examples:
// !sysret
// !sysret2
// !sysret pid 400
// !sysret2 pid 400
// !sysret core 2 pid 400
// !sysret2 core 2 pid 400
func SysRet() { // monitors and hooks all execution of sysret instructions (by accessing memory and checking for instructions).
	InterpreterEx("!sysret")
}

// TraceExecution
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
func TraceExecution() { // traces the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!trace")
}

// TrackModeTransitionInstructions
// Syntax:
// !track [tree] [Count (hex)]
// !track [reg] [Count (hex)]
// Examples:
// !track tree 10000
// !track reg 10000
func TrackModeTransitionInstructions() { // tracks instructions from user-mode to kernel-mode or kernel-mode to user-mode to create call tree. Please note that it's highly recommended to configure symbols before using this command as it maps addresses to corresponding function names.
	InterpreterEx("!track")
}

// TscInstructionsMonitor
// Syntax:
// !tsc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !tsc
// !tsc pid 400
// !tsc core 2 pid 400
func TscInstructionsMonitor() { // monitors execution of rdtsc/rdtscp instructions.
	InterpreterEx("!tsc")
}

// UnHide
// Syntax:
// !unhide
// Examples:
// !unhide
func UnHide() { // reverts the transparency measures of the '!hide' command.
	InterpreterEx("!unhide")
}

// Va2Pa
// Syntax:
// !va2pa [VirtualAddress (hex)] [pid ProcessId (hex)]
// Examples:
// !va2pa nt!ExAllocatePoolWithTag
// !va2pa nt!ExAllocatePoolWithTag+5
// !va2pa @rcx
// !va2pa @rcx+5
// !va2pa fffff801deadbeef
// !va2pa fffff801deadbeef pid 0xc8
func Va2Pa() { // converts virtual address to physical address.
	InterpreterEx("!va2pa")
}

// VmCall
// Syntax:
// !vmcall [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// Examples:
// !vmcall
// !vmcall pid 400
// !vmcall core 2 pid 400
func VmCall() { // monitors execution of VMCALL instruction.
	InterpreterEx("!vmcall")
}

// HardwareClockDebugging
// Syntax:
// !hw_clk  [script { Script (string) }]
// Examples:
// !hw_clk script { @hw_pin1 = 0; }
func HardwareClockDebugging() { // performs actions related to hwdbg hardware debugging events for each clock cycle.
	InterpreterEx("!hw_clk")
}

// AttachDebugThread
// Syntax:
// .attach [pid ProcessId (hex)]
// Examples:
// .attach pid b60
func AttachDebugThread() { // attaches to debug a thread in VMI Mode.
	InterpreterEx(".attach")
}

// ClearScreen
// Syntax:
// .cls
// Examples:
func ClearScreen() { // clears the screen.
	InterpreterEx(".cls")
}

// ConnectToMachine
// Syntax:
// .connect [local]
// .connect [Ip (string)] [Port (decimal)]
// Examples:
// .connect local
// .connect 192.168.1.5 50000
func ConnectToMachine() { // connects to a remote or local machine to start debugging.
	InterpreterEx(".connect")
}

// DebugMachine
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
func DebugMachine() { // debugs a target machine or makes this machine a debuggee.
	InterpreterEx(".debug")
}

// DetachDebugging
// Syntax:
// .detach
// Examples:
func DetachDebugging() { // detaches from debugging a user-mode process.
	InterpreterEx(".detach")
}

// DisconnectSession
// Syntax:
// .disconnect
// Examples:
func DisconnectSession() { // disconnects from a debugging session (it won't unload the modules).
	InterpreterEx(".disconnect")
}

// DumpMemoryContext
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
func DumpMemoryContext() { // saves memory context into a file.
	InterpreterEx(".dump")
}

// FormatsDiff
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
func FormatsDiff() { // shows a value or register in different formats
	InterpreterEx(".formats")
}

// HelpForCommand
// Syntax:
// .help [Command (string)]
// Examples:
// .help !monitor
func HelpForCommand() { // shows help and example(s) of a specific command
	InterpreterEx(".help")
}

// KillProcess
// Syntax:
// .kill
// Examples:
func KillProcess() { // terminates the current running process
	InterpreterEx(".kill")
}

// ListenForClientConnection
// Syntax:
// .listen [Port (decimal)]
// Examples:
// .listen
// .listen 50000
// Notes:
// if you don't specify port then HyperDbg uses the default port
func ListenForClientConnection() { // listens for a client to connect to HyperDbg (works as a guest server)
	InterpreterEx(".listen")
}

// LogClose
// Syntax:
// .logclose
// Examples:
func LogClose() { // closes the previously opened log
	InterpreterEx(".logclose")
}

// LogOpen
// Syntax:
// .logopen [FilePath (string)]
// Examples:
func LogOpen() { // saves commands and results in a file
	InterpreterEx(".logopen")
}

// PageAvailableInRam
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
func PageAvailableInRam() { // brings the page in, making it available in the RAM
	InterpreterEx(".pagein")
}

// ParseExecutableFiles
// Syntax:
// .pe [header] [FilePath (string)]
// .pe [section] [SectionName (string)] [FilePath (string)]
// Examples:
// .pe header c:\reverse files\myfile.exe
// .pe section .text c:\reverse files\myfile.exe
// .pe section .rdata c:\reverse files\myfile.exe
func ParseExecutableFiles() { // parses portable executable (PE) files and dump sections
	InterpreterEx(".pe")
}

// ProcessesView
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
func ProcessesView() { // shows and changes the processes. This command needs public symbols for ntoskrnl.exe if you want to see the processes list. Please visit the documentation to know about the difference between '.process' and '.process2'.
	InterpreterEx(".process")
}

// RestartProcess
// Syntax:
// .restart
// Examples:
func RestartProcess() { // restarts the previously started process using '.start' command.
	InterpreterEx(".restart")
}

// Script
// Syntax:
// .script [FilePath (string)] [Args (string)]
// Examples:
// .script C:\scripts\script.ds
// .script C:\scripts\script.ds 95 85 @rsp
// .script "C:\scripts\hello world.ds"
// .script "C:\scripts\hello world.ds" @rax
// .script "C:\scripts\hello world.ds" @rax @rcx+55 $pid
// .script "C:\scripts\hello world.ds" 12 55 @rip
func Script() { // runs a HyperDbg script.
	InterpreterEx(".script")
}

// StartProcess
// Syntax:
// .start [path Path (string)] [Parameters (string)]
// Examples:
// .start path c:\reverse eng\my_file.exe
func StartProcess() { // runs a user-mode process
	InterpreterEx(".start")
}

// Status
// Syntax:
// .status
// status
// Examples:
// Notes:
// If connected to a remote system, '.status' shows the state of current debugger, while 'status' shows the state of remote debuggee.
func Status() { // gets the status of current debugger in local system
	InterpreterEx(".status")
}

// SwitchThread
// Syntax:
// .switch
// .switch [pid ProcessId (hex)]
// .switch [tid ThreadId (hex)]
// Examples:
// .switch list
// .switch pid b60
// .switch tid b60
func SwitchThread() { // shows the list of active debugging threads and switches between processes and threads in VMI Mode
	InterpreterEx(".switch")
}

// Symbol
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
func Symbol() { // performs the symbol actions
	InterpreterEx(".sym")
}

// SymbolPath
// Syntax:
// .sympath
// .sympath [SymServer (string)]
// Examples:
// .sympath
// .sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbols
func SymbolPath() { // shows and sets the symbol server and path
	InterpreterEx(".sympath")
}

// Thread
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
func Thread() { // shows and changes the threads
	InterpreterEx(".thread")
}
