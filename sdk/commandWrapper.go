package sdk

func ClearsBreakpointBreakpointId() {
	InterpreterEx("bc")
}

func DisablesBreakpointBreakpointId() {
	InterpreterEx("bd")
}

func EnablesBreakpointBreakpointId() {
	InterpreterEx("be")
}

func ListsAllEnabledAndDisabledBreakpoints() {
	InterpreterEx("bl")
}

func PutsBreakpoint0xcc() {
	InterpreterEx("bp")
}

func ShowsAndChangesOperatingProcessor() {
	InterpreterEx("core")
}

func CollectsReportFromCpuFeatures() {
	InterpreterEx("cpu")
}

func RemoveKernelModeBuffers() {
	InterpreterEx("flush")
}

func ContinueDebuggee_() {
	InterpreterEx("g")
}

func ExecuteSingleInstructionStepOut() {
	InterpreterEx("gu")
}

func ExecuteSingleInstructionStepIn() {
	InterpreterEx("i")
}

func ShowThreadCallstack() {
	InterpreterEx("k, kd, kq")
}

func ListModules() {
	InterpreterEx("lm")
}

func InstallDriversAndLoadModules() {
	InterpreterEx("load")
}

func CreateEventForwardingOutputInstance() {
	InterpreterEx("output")
}

func ExecuteSingleInstructionStep() {
	InterpreterEx("p")
}

func PauseKernelEvents() {
	InterpreterEx("pause")
}

func PreactivateSpecialFunctionality() {
	InterpreterEx("preactivate")
}

func PreallocateBuffer() {
	InterpreterEx("prealloc")
}

func EvaluateExpressions() {
	InterpreterEx("print")
}

func ReadModifyRegisters() {
	InterpreterEx("r")
}

func ReadModelSpecificRegister() {
	InterpreterEx("rdmsr")
}

func SearchMemoryPattern() {
	InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

func SettingsManagement() {
	InterpreterEx("settings")
}

func DebuggerSleep() {
	InterpreterEx("sleep")
}

func ExecuteStepIn() {
	InterpreterEx("t")
}

func TestHyperDbgFeatures() {
	InterpreterEx("test")
}

func UnloadKernelModules() {
	InterpreterEx("unload")
}

func SearchSymbols() {
	InterpreterEx("x")
}

func MonitorCpuidExecution() {
	InterpreterEx("!cpuid")
}

func ControlRegisterModificationMonitor() {
	InterpreterEx("!crwrite")
}

func DebugRegistersMonitor() {
	InterpreterEx("!dr")
}

func EptHook() {
	InterpreterEx("!epthook")
}

func EptHook2() {
	InterpreterEx("!epthook2")
}

func IdtEntriesMonitor() {
	InterpreterEx("!exception")
}

func HideHyperDbg() {
	InterpreterEx("!hide")
}

func ExternalInterruptsMonitor() {
	InterpreterEx("!interrupt")
}

func IoInstructionsInDetect() {
	InterpreterEx("!ioin")
}

func DetectIoInstructionsOut() {
	InterpreterEx("!ioout")
}

func MeasureArgumentsForHide() {
	InterpreterEx("!measure")
}

func TrapModeInstructions() {
	InterpreterEx("!mode")
}

func MonitorMemoryAccess() {
	InterpreterEx("!monitor")
}

func MsrRead() {
	InterpreterEx("!msrread")
}

func MsrWrite() {
	InterpreterEx("!msrwrite")
}

func Pa2Va() {
	InterpreterEx("!pa2va")
}

func PmcExecutionMonitor() {
	InterpreterEx("!pmc")
}

func Pte() {
	InterpreterEx("!pte")
}

func ReversingMachineModuleUse() {
	InterpreterEx("!rev")
}

func Syscall() {
	InterpreterEx("!syscall")
}

func SysRet() {
	InterpreterEx("!sysret")
}

func TraceExecution() {
	InterpreterEx("!trace")
}

func TrackModeTransitionInstructions() {
	InterpreterEx("!track")
}

func TscInstructionsMonitor() {
	InterpreterEx("!tsc")
}

func UnHide() {
	InterpreterEx("!unhide")
}

func Va2Pa() {
	InterpreterEx("!va2pa")
}

func VmCall() {
	InterpreterEx("!vmcall")
}

func HardwareClockDebugging() {
	InterpreterEx("!hw_clk")
}

func AttachDebugThread() {
	InterpreterEx(".attach")
}

func ClearScreen() {
	InterpreterEx(".cls")
}

func ConnectToMachine() {
	InterpreterEx(".connect")
}

func DebugMachine() {
	InterpreterEx(".debug")
}

func DetachDebugging() {
	InterpreterEx(".detach")
}

func DisconnectSession() {
	InterpreterEx(".disconnect")
}

func DumpMemoryContext() {
	InterpreterEx(".dump")
}

func FormatsDiff() {
	InterpreterEx(".formats")
}

func HelpForCommand() {
	InterpreterEx(".help")
}

func KillProcess() {
	InterpreterEx(".kill")
}

func ListenForClientConnection() {
	InterpreterEx(".listen")
}

func LogClose() {
	InterpreterEx(".logclose")
}

func LogOpen() {
	InterpreterEx(".logopen")
}

func PageAvailableInRam() {
	InterpreterEx(".pagein")
}

func ParseExecutableFiles() {
	InterpreterEx(".pe")
}

func ProcessesView() {
	InterpreterEx(".process")
}

func RestartProcess() {
	InterpreterEx(".restart")
}

func Script() {
	InterpreterEx(".script")
}

func StartProcess() {
	InterpreterEx(".start")
}

func Status() {
	InterpreterEx(".status")
}

func SwitchThread() {
	InterpreterEx(".switch")
}

func Symbol() {
	InterpreterEx(".sym")
}

func SymbolPath() {
	InterpreterEx(".sympath")
}

func Thread() {
	InterpreterEx(".thread")
}
