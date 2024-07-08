package sdk

// BreakpointClearByID
// Syntax:
// bc [BreakpointId (hex)]
func BreakpointClearByID() { // bc : clears a breakpoint using breakpoint id.
	InterpreterEx("bc")
}

// BreakpointDisableByID
// Syntax:
// bd [BreakpointId (hex)]
func BreakpointDisableByID() { // bd : disables a breakpoint using breakpoint id.
	InterpreterEx("bd")
}

// BreakpointEnableByID
// Syntax:
// be [BreakpointId (hex)]
func BreakpointEnableByID() { // be : enables a breakpoint using breakpoint id.
	InterpreterEx("be")
}

// BreakpointList
// Syntax:
// bl
func BreakpointList() { // bl : lists all the enabled and disabled breakpoints.
	InterpreterEx("bl")
}

// SetBreakpoint0xcc
// Syntax:
// bp [Address (hex)] [pid ProcessId (hex)] [tid ThreadId (hex)] [core CoreId (hex)]
func SetBreakpoint0xcc() { // bp : puts a breakpoint (0xcc).

	InterpreterEx("bp")
}

// CoreOperatingProcessorForShowsAndChanges
// Syntax:
// ~
// ~ [CoreNumber (hex)]
func CoreOperatingProcessorForShowsAndChanges() { //~ : shows and changes the operating processor.
	InterpreterEx("core")
}

// CpuFeaturesForCollectsReport
// Syntax:
// cpu
func CpuFeaturesForCollectsReport() { // cpu : collects a report from cpu features.
	InterpreterEx("cpu")
}

// FlushKernelModeBuffers
// Syntax:
// flush
func FlushKernelModeBuffers() { // removes all the buffer and messages from kernel-mode buffers.
	InterpreterEx("flush")
}

// ContinueDebugger
// Syntax:
// g
func ContinueDebugger() { // continues debuggee or continues processing kernel messages.
	InterpreterEx("g")
}

// StepOut
// Syntax:
// gu
// gu [Count (hex)]
func StepOut() { // executes a single instruction (step-out) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("gu")
}

// StepIn
// Syntax:
// i
// i [Count (hex)]
// ir
// ir [Count (hex)]
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
func CallstackOrThreadView() { // shows the callstack of the thread.
	InterpreterEx("k, kd, kq")
}

// ListModules
// Syntax:
// lm [m Name (string)] [pid ProcessId (hex)] [Filter (string)]
func ListModules() { // lists user/kernel modules' base address, size, name and path.
	InterpreterEx("lm")
}

// LoadDriversAndModules
// Syntax:
// load [ModuleName (string)]
func LoadDriversAndModules() { // installs the drivers and load the modules.
	InterpreterEx("load")
}

// OutputEventForwardingInstance
// Syntax:
// output [create Name (string)] [file|namedpipe|tcp Address (string)]
// output [open|close Name (string)]
func OutputEventForwardingInstance() { // creates an output instance that can be used in event forwarding.
	InterpreterEx("output")
}

// StepExecuteSingleInstruction
// Syntax:
// p
// p [Count (hex)]
// pr
// pr [Count (hex)]
func StepExecuteSingleInstruction() { // executes a single instruction (step) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("p")
}

// PauseKernelEvents
// Syntax:
// pause
func PauseKernelEvents() { // pauses the kernel events.
	InterpreterEx("pause")
}

// PreactivateSpecialFunctionality
// Syntax:
// preactivate  [Type (string)]
func PreactivateSpecialFunctionality() { // preactivates a special functionality.
	InterpreterEx("preactivate")
}

// PreallocateBuffer
// Syntax:
// prealloc [Type (string)] [Count (hex)]
func PreallocateBuffer() { // pre-allocates buffer for special purposes.
	InterpreterEx("prealloc")
}

// PrintEvaluateExpressions
// Syntax:
// print [Expression (string)]
func PrintEvaluateExpressions() { // evaluates expressions.
	InterpreterEx("print")
}

// RegistersReadOrModify
// Syntax:
// r
// r [Register (string)] [= Expr (string)]
func RegistersReadOrModify() { // reads or modifies registers.
	InterpreterEx("r")
}

// ReadMsr
// Syntax:
// rdmsr [Msr (hex)] [core CoreNumber (hex)]
func ReadMsr() { // reads a model-specific register (MSR).
	InterpreterEx("rdmsr")
}

// SearchMemoryPattern
// Syntax:
// sb [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
// sd [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
// sq [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]
func SearchMemoryPattern() { // searches a contiguous memory for a special byte pattern.
	InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

// SettingsManagement
// Syntax:
// settings [OptionName (string)]
// settings [OptionName (string)] [Value (hex)]
// settings [OptionName (string)] [Value (string)]
// settings [OptionName (string)] [on|off]
func SettingsManagement() { // queries, sets, or changes a value for a special settings option.
	InterpreterEx("settings")
}

// SleepDebugger
// Syntax:
// sleep [MillisecondsTime (hex)]
func SleepDebugger() { // sleeps the debugger; this command is used in scripts, it doesn't breaks the debugger but the debugger still shows the buffers received from kernel.
	InterpreterEx("sleep")
}

// StepInExecute
// Syntax:
// t
// t [Count (hex)]
// tr
// tr [Count (hex)]
func StepInExecute() { // executes a single instruction (step-in) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("t")
}

// TestHyperDbgFeatures
// Syntax:
// test [Task (string)]
func TestHyperDbgFeatures() { // tests essential features of HyperDbg in current machine.
	InterpreterEx("test")
}

// UnloadKernelModules
// Syntax:
// unload [remove] [ModuleName (string)]
func UnloadKernelModules() { // unloads the kernel modules and uninstalls the drivers.
	InterpreterEx("unload")
}

// SearchSymbols
// Syntax:
// x [Module!Name (wildcard string)]
func SearchSymbols() { // searches and shows symbols (wildcard) and corresponding addresses.
	InterpreterEx("x")
}

// CpuidExecutionMonitor
// Syntax:
// !cpuid [Eax (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func CpuidExecutionMonitor() { // monitors execution of a special cpuid index or all cpuids instructions.
	InterpreterEx("!cpuid")
}

// ControlRegisterModificationMonitor
// Syntax:
// !crwrite [Cr (hex)] [mask Mask (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func ControlRegisterModificationMonitor() { // monitors modification of control registers (CR0 / CR4).
	InterpreterEx("!crwrite")
}

// DebugRegistersMonitor
// Syntax:
// !dr [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func DebugRegistersMonitor() { // monitors any access to debug registers.
	InterpreterEx("!dr")
}

// EptHook
// Syntax:
// !epthook [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func EptHook() { // puts a hidden-hook EPT (hidden breakpoints).
	InterpreterEx("!epthook")
}

// EptHook2
// Syntax:
// !epthook2 [Address (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func EptHook2() { // puts a hidden-hook EPT (detours).
	InterpreterEx("!epthook2")
}

// IdtEntriesMonitor
// Syntax:
// !exception [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func IdtEntriesMonitor() { // monitors the first 32 entry of IDT (starting from zero).
	InterpreterEx("!exception")
}

// HideHyperDbg
// Syntax:
// !hide
// !hide [pid ProcessId (hex)]
// !hide [name ProcessName (string)]
func HideHyperDbg() { // tries to make HyperDbg transparent from anti-debugging and anti-hypervisor methods.
	InterpreterEx("!hide")
}

// InterruptExternalMonitor
// Syntax:
// [IdtIndex (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func InterruptExternalMonitor() { // monitors the external interrupt (IDT >= 32).
	InterpreterEx("!interrupt")
}

// IoInDetect
// Syntax:
// !ioin [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func IoInDetect() { // detects the execution of IN (I/O instructions) instructions.
	InterpreterEx("!ioin")
}

// IoOutDetect
// Syntax:
// !ioout [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func IoOutDetect() { // detects the execution of OUT (I/O instructions) instructions.
	InterpreterEx("!ioout")
}

// MeasureArgumentsForHide
// Syntax:
// !measure [default]
func MeasureArgumentsForHide() { // measures the arguments needs for the '!hide' command.
	InterpreterEx("!measure")
}

// ModeInstructionsTrap
// Syntax:
// !mode [Mode (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func ModeInstructionsTrap() { // traps (and possibly blocks) the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!mode")
}

// MonitorMemoryAccess
// Syntax:
// !monitor [MemoryType (vapa)] [Attribute (string)] [FromAddress (hex)] [ToAddress (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// !monitor [MemoryType (vapa)] [Attribute (string)] [FromAddress (hex)] [l Length (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func MonitorMemoryAccess() { // monitors address range for read and writes.
	InterpreterEx("!monitor")
}

// MsrRead
// Syntax:
// !msrread [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func MsrRead() { // detects the execution of rdmsr instructions.
	InterpreterEx("!msrread")
}

// MsrWrite
// Syntax:
// !msrwrite [Msr (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func MsrWrite() { // detects the execution of wrmsr instructions.
	InterpreterEx("!msrwrite")
}

// Pa2Va
// Syntax:
// !pa2va [PhysicalAddress (hex)] [pid ProcessId (hex)]
func Pa2Va() { // converts virtual address to physical address.
	InterpreterEx("!pa2va")
}

// PmcExecutionMonitor
// Syntax:
// !pmc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func PmcExecutionMonitor() { // monitors execution of rdpmc instructions.
	InterpreterEx("!pmc")
}

// Pte
// Syntax:
// !pte [VirtualAddress (hex)] [pid ProcessId (hex)]
func Pte() { // finds virtual addresses of different paging-levels.
	InterpreterEx("!pte")
}

// ReversingMachineModuleUse
// Syntax:
// !rev [config] [pid ProcessId (hex)]
// !rev [path Path (string)] [Parameters (string)]
func ReversingMachineModuleUse() { // uses the reversing machine module in order to reconstruct the programmer/memory assumptions.
	InterpreterEx("!rev")
}

// Syscall
// Syntax:
// !syscall [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
// !syscall2 [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func Syscall() { // monitors and hooks all execution of syscall instructions (by accessing memory and checking for instructions).
	InterpreterEx("!syscall")
}

// SysRet
// Syntax:
// !sysret [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]
func SysRet() { // monitors and hooks all execution of sysret instructions (by accessing memory and checking for instructions).
	InterpreterEx("!sysret")
}

// TraceExecution
// Syntax:
// !trace [TraceType (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func TraceExecution() { // traces the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!trace")
}

// TrackModeTransitionInstructions
// Syntax:
// !track [tree] [Count (hex)]
// !track [reg] [Count (hex)]
func TrackModeTransitionInstructions() { // tracks instructions from user-mode to kernel-mode or kernel-mode to user-mode to create call tree. Please note that it's highly recommended to configure symbols before using this command as it maps addresses to corresponding function names.
	InterpreterEx("!track")
}

// TscInstructionsMonitor
// Syntax:
// !tsc [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func TscInstructionsMonitor() { // monitors execution of rdtsc/rdtscp instructions.
	InterpreterEx("!tsc")
}

// UnHide
// Syntax:
// !unhide
func UnHide() { // reverts the transparency measures of the '!hide' command.
	InterpreterEx("!unhide")
}

// Va2Pa
// Syntax:
// !va2pa [VirtualAddress (hex)] [pid ProcessId (hex)]
func Va2Pa() { // converts virtual address to physical address.
	InterpreterEx("!va2pa")
}

// VmCall
// Syntax:
// !vmcall [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]
func VmCall() { // monitors execution of VMCALL instruction.
	InterpreterEx("!vmcall")
}

// HardwareClockDebugging
// Syntax:
// !hw_clk  [script { Script (string) }]
func HardwareClockDebugging() { // performs actions related to hwdbg hardware debugging events for each clock cycle.
	InterpreterEx("!hw_clk")
}

// AttachDebugThread
// Syntax:
// .attach [pid ProcessId (hex)]
func AttachDebugThread() { // attaches to debug a thread in VMI Mode.
	InterpreterEx(".attach")
}

// ClearScreen
// Syntax:
// .cls
func ClearScreen() { // clears the screen.
	InterpreterEx(".cls")
}

// ConnectToMachine
// Syntax:
// .connect [local]
// .connect [Ip (string)] [Port (decimal)]
func ConnectToMachine() { // connects to a remote or local machine to start debugging.
	InterpreterEx(".connect")
}

// DebugMachine
// Syntax:
// .debug [remote] [serial|namedpipe] [Baudrate (decimal)] [Address (string)]
// .debug [prepare] [serial] [Baudrate (decimal)] [Address (string)]
// .debug [close]
func DebugMachine() { // debugs a target machine or makes this machine a debuggee.
	InterpreterEx(".debug")
}

// DetachDebugging
// Syntax:
// .detach
func DetachDebugging() { // detaches from debugging a user-mode process.
	InterpreterEx(".detach")
}

// DisconnectSession
// Syntax:
// .disconnect
func DisconnectSession() { // disconnects from a debugging session (it won't unload the modules).
	InterpreterEx(".disconnect")
}

// DumpMemoryContext
// Syntax:
// .dump [FromAddress (hex)] [ToAddress (hex)] [pid ProcessId (hex)] [path Path (string)]
func DumpMemoryContext() { // saves memory context into a file.
	InterpreterEx(".dump")
}

// FormatsDiff
// Syntax:
// .formats [Expression (string)]
func FormatsDiff() { // shows a value or register in different formats
	InterpreterEx(".formats")
}

// HelpForCommand
// Syntax:
// .help [Command (string)]
func HelpForCommand() { // shows help and example(s) of a specific command
	InterpreterEx(".help")
}

// KillProcess
// Syntax:
// .kill
func KillProcess() { // terminates the current running process
	InterpreterEx(".kill")
}

// ListenForClientConnection
// Syntax:
// .listen [Port (decimal)]
func ListenForClientConnection() { // listens for a client to connect to HyperDbg (works as a guest server)
	InterpreterEx(".listen")
}

// LogClose
// Syntax:
// .logclose
func LogClose() { // closes the previously opened log
	InterpreterEx(".logclose")
}

// LogOpen
// Syntax:
// .logopen [FilePath (string)]
func LogOpen() { // saves commands and results in a file
	InterpreterEx(".logopen")
}

// PageAvailableInRam
// Syntax:
// .pagein [Mode (string)] [l Length (hex)]
// .pagein [Mode (string)] [VirtualAddress (hex)] [l Length (hex)]
func PageAvailableInRam() { // brings the page in, making it available in the RAM
	InterpreterEx(".pagein")
}

// ParseExecutableFiles
// Syntax:
// .pe [header] [FilePath (string)]
// .pe [section] [SectionName (string)] [FilePath (string)]
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
func ProcessesView() { // shows and changes the processes. This command needs public symbols for ntoskrnl.exe if you want to see the processes list. Please visit the documentation to know about the difference between '.process' and '.process2'.
	InterpreterEx(".process")
}

// RestartProcess
// Syntax:
// .restart
func RestartProcess() { // restarts the previously started process using '.start' command.
	InterpreterEx(".restart")
}

// Script
// Syntax:
// .script [FilePath (string)] [Args (string)]
func Script() { // runs a HyperDbg script.
	InterpreterEx(".script")
}

// StartProcess
// Syntax:
// .start [path Path (string)] [Parameters (string)]
func StartProcess() { // runs a user-mode process
	InterpreterEx(".start")
}

// Status
// Syntax:
// .status
// status
func Status() { // gets the status of current debugger in local system
	InterpreterEx(".status")
}

// SwitchThread
// Syntax:
// .switch
// .switch [pid ProcessId (hex)]
// .switch [tid ThreadId (hex)]
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
func Symbol() { // performs the symbol actions
	InterpreterEx(".sym")
}

// SymbolPath
// Syntax:
// .sympath
// .sympath [SymServer (string)]
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
func Thread() { // shows and changes the threads
	InterpreterEx(".thread")
}
