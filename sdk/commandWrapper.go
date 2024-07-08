package sdk

func BreakpointClearByID() { // bc : clears a breakpoint using breakpoint id.
	InterpreterEx("bc")
}

func BreakpointDisableByID() { // bd : disables a breakpoint using breakpoint id.
	InterpreterEx("bd")
}

func BreakpointEnableByID() { // be : enables a breakpoint using breakpoint id.
	InterpreterEx("be")
}

func BreakpointList() { // bl : lists all the enabled and disabled breakpoints.
	InterpreterEx("bl")
}

func SetBreakpoint0xcc() { // bp : puts a breakpoint (0xcc).

	InterpreterEx("bp")
}

func CoreOperatingProcessorForShowsAndChanges() { //~ : shows and changes the operating processor.
	InterpreterEx("core")
}

func CpuFeaturesForCollectsReport() { // cpu : collects a report from cpu features.
	InterpreterEx("cpu")
}

func FlushKernelModeBuffers() { // removes all the buffer and messages from kernel-mode buffers.
	InterpreterEx("flush")
}

func ContinueDebugger() { // continues debuggee or continues processing kernel messages.
	InterpreterEx("g")
}

func StepOut() { // executes a single instruction (step-out) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("gu")
}

func StepIn() { // executes a single instruction (step-in) and guarantees that no other instruction is executed other than the displayed instruction including user to the kernel (syscalls) and kernel to the user (sysrets) and exceptions and page-faults and optionally displays all registers and flags' resulting values.
	InterpreterEx("i")
}

func CallstackOrThreadView() { // shows the callstack of the thread.
	InterpreterEx("k, kd, kq")
}

func ListModules() { // lists user/kernel modules' base address, size, name and path.
	InterpreterEx("lm")
}

func LoadDriversAndModules() { // installs the drivers and load the modules.
	InterpreterEx("load")
}

func OutputEventForwardingInstance() { // creates an output instance that can be used in event forwarding.
	InterpreterEx("output")
}

func StepExecuteSingleInstruction() { // executes a single instruction (step) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("p")
}

func PauseKernelEvents() { // pauses the kernel events.
	InterpreterEx("pause")
}

func PreactivateSpecialFunctionality() { // preactivates a special functionality.
	InterpreterEx("preactivate")
}

func PreallocateBuffer() { // pre-allocates buffer for special purposes.
	InterpreterEx("prealloc")
}

func PrintEvaluateExpressions() { // evaluates expressions.
	InterpreterEx("print")
}

func RegistersReadOrModify() { // reads or modifies registers.
	InterpreterEx("r")
}

func ReadMsr() { // reads a model-specific register (MSR).
	InterpreterEx("rdmsr")
}

func SearchMemoryPattern() { // searches a contiguous memory for a special byte pattern.
	InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

func SettingsManagement() { // queries, sets, or changes a value for a special settings option.
	InterpreterEx("settings")
}

func SleepDebugger() { // sleeps the debugger; this command is used in scripts, it doesn't breaks the debugger but the debugger still shows the buffers received from kernel.
	InterpreterEx("sleep")
}

func StepInExecute() { // executes a single instruction (step-in) and optionally displays the resulting values of all registers and flags.
	InterpreterEx("t")
}

func TestHyperDbgFeatures() { // tests essential features of HyperDbg in current machine.
	InterpreterEx("test")
}

func UnloadKernelModules() { // unloads the kernel modules and uninstalls the drivers.
	InterpreterEx("unload")
}

func SearchSymbols() { // searches and shows symbols (wildcard) and corresponding addresses.
	InterpreterEx("x")
}

func CpuidExecutionMonitor() { // monitors execution of a special cpuid index or all cpuids instructions.
	InterpreterEx("!cpuid")
}

func ControlRegisterModificationMonitor() { // monitors modification of control registers (CR0 / CR4).
	InterpreterEx("!crwrite")
}

func DebugRegistersMonitor() { // monitors any access to debug registers.
	InterpreterEx("!dr")
}

func EptHook() { // puts a hidden-hook EPT (hidden breakpoints).
	InterpreterEx("!epthook")
}

func EptHook2() { // puts a hidden-hook EPT (detours).
	InterpreterEx("!epthook2")
}

func IdtEntriesMonitor() { // monitors the first 32 entry of IDT (starting from zero).
	InterpreterEx("!exception")
}

func HideHyperDbg() { // tries to make HyperDbg transparent from anti-debugging and anti-hypervisor methods.
	InterpreterEx("!hide")
}

func InterruptExternalMonitor() { // monitors the external interrupt (IDT >= 32).
	InterpreterEx("!interrupt")
}

func IoInDetect() { // detects the execution of IN (I/O instructions) instructions.
	InterpreterEx("!ioin")
}

func IoOutDetect() { // detects the execution of OUT (I/O instructions) instructions.
	InterpreterEx("!ioout")
}

func MeasureArgumentsForHide() { // measures the arguments needs for the '!hide' command.
	InterpreterEx("!measure")
}

func ModeInstructionsTrap() { // traps (and possibly blocks) the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!mode")
}

func MonitorMemoryAccess() { // monitors address range for read and writes.
	InterpreterEx("!monitor")
}

func MsrRead() { // detects the execution of rdmsr instructions.
	InterpreterEx("!msrread")
}

func MsrWrite() { // detects the execution of wrmsr instructions.
	InterpreterEx("!msrwrite")
}

func Pa2Va() { // converts virtual address to physical address.
	InterpreterEx("!pa2va")
}

func PmcExecutionMonitor() { // monitors execution of rdpmc instructions.
	InterpreterEx("!pmc")
}

func Pte() { // finds virtual addresses of different paging-levels.
	InterpreterEx("!pte")
}

func ReversingMachineModuleUse() { // uses the reversing machine module in order to reconstruct the programmer/memory assumptions.
	InterpreterEx("!rev")
}

func Syscall() { // monitors and hooks all execution of syscall instructions (by accessing memory and checking for instructions).
	InterpreterEx("!syscall")
}

func SysRet() { // monitors and hooks all execution of sysret instructions (by accessing memory and checking for instructions).
	InterpreterEx("!sysret")
}

func TraceExecution() { // traces the execution of user-mode/kernel-mode instructions.
	InterpreterEx("!trace")
}

func TrackModeTransitionInstructions() { // tracks instructions from user-mode to kernel-mode or kernel-mode to user-mode to create call tree. Please note that it's highly recommended to configure symbols before using this command as it maps addresses to corresponding function names.
	InterpreterEx("!track")
}

func TscInstructionsMonitor() { // monitors execution of rdtsc/rdtscp instructions.
	InterpreterEx("!tsc")
}

func UnHide() { // reverts the transparency measures of the '!hide' command.
	InterpreterEx("!unhide")
}

func Va2Pa() { // converts virtual address to physical address.
	InterpreterEx("!va2pa")
}

func VmCall() { // monitors execution of VMCALL instruction.
	InterpreterEx("!vmcall")
}

func HardwareClockDebugging() { // performs actions related to hwdbg hardware debugging events for each clock cycle.
	InterpreterEx("!hw_clk")
}

func AttachDebugThread() { // attaches to debug a thread in VMI Mode.
	InterpreterEx(".attach")
}

func ClearScreen() { // clears the screen.
	InterpreterEx(".cls")
}

func ConnectToMachine() { // connects to a remote or local machine to start debugging.
	InterpreterEx(".connect")
}

func DebugMachine() { // debugs a target machine or makes this machine a debuggee.
	InterpreterEx(".debug")
}

func DetachDebugging() { // detaches from debugging a user-mode process.
	InterpreterEx(".detach")
}

func DisconnectSession() { // disconnects from a debugging session (it won't unload the modules).
	InterpreterEx(".disconnect")
}

func DumpMemoryContext() { // saves memory context into a file.
	InterpreterEx(".dump")
}

func FormatsDiff() { // shows a value or register in different formats
	InterpreterEx(".formats")
}

func HelpForCommand() { // shows help and example(s) of a specific command
	InterpreterEx(".help")
}

func KillProcess() { // terminates the current running process
	InterpreterEx(".kill")
}

func ListenForClientConnection() { // listens for a client to connect to HyperDbg (works as a guest server)
	InterpreterEx(".listen")
}

func LogClose() { // closes the previously opened log
	InterpreterEx(".logclose")
}

func LogOpen() { // saves commands and results in a file
	InterpreterEx(".logopen")
}

func PageAvailableInRam() { // brings the page in, making it available in the RAM
	InterpreterEx(".pagein")
}

func ParseExecutableFiles() { // parses portable executable (PE) files and dump sections
	InterpreterEx(".pe")
}

func ProcessesView() { // shows and changes the processes. This command needs public symbols for ntoskrnl.exe if you want to see the processes list. Please visit the documentation to know about the difference between '.process' and '.process2'.
	InterpreterEx(".process")
}

func RestartProcess() { // restarts the previously started process using '.start' command.
	InterpreterEx(".restart")
}

func Script() { // runs a HyperDbg script.
	InterpreterEx(".script")
}

func StartProcess() { // runs a user-mode process
	InterpreterEx(".start")
}

func Status() { // gets the status of current debugger in local system
	InterpreterEx(".status")
}

func SwitchThread() { // shows the list of active debugging threads and switches between processes and threads in VMI Mode
	InterpreterEx(".switch")
}

func Symbol() { // performs the symbol actions
	InterpreterEx(".sym")
}

func SymbolPath() { // shows and sets the symbol server and path
	InterpreterEx(".sympath")
}

func Thread() { // shows and changes the threads
	InterpreterEx(".thread")
}
