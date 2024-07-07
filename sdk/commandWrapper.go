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

func MonitorControlRegisterModification() {
	InterpreterEx("!crwrite")
}

func MonitorDebugRegisters() {
	InterpreterEx("!dr")
}

func HiddenHookEPTBreakpoints() {
	InterpreterEx("!epthook")
}

func HiddenHookEPTDetours() {
	InterpreterEx("!epthook2")
}

func MonitorIDTEntries() {
	InterpreterEx("!exception")
}

func HideHyperDbg() {
	InterpreterEx("!hide")
}

func MonitorExternalInterrupts() {
	InterpreterEx("!interrupt")
}

func DetectIOInstructionsIn() {
	InterpreterEx("!ioin")
}

func DetectIOInstructionsOut() {
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

func DetectRdmsrExecution() {
	InterpreterEx("!msrread")
}

func DetectWrmsrExecution() {
	InterpreterEx("!msrwrite")
}

func ConvertPhysicalToVirtual() {
	InterpreterEx("!pa2va")
}

func MonitorRdpmcExecution() {
	InterpreterEx("!pmc")
}

func FindVirtualAddressPagingLevels() {
	InterpreterEx("!pte")
}

func UseReversingMachineModule() {
	InterpreterEx("!rev")
}

func MonitorSyscallExecution() {
	InterpreterEx("!syscall")
}

func MonitorSysretExecution() {
	InterpreterEx("!sysret")
}

func TraceExecution() {
	InterpreterEx("!trace")
}

func TrackModeTransitionInstructions() {
	InterpreterEx("!track")
}

func MonitorRdtscInstructions() {
	InterpreterEx("!tsc")
}

func RevertHide() {
	InterpreterEx("!unhide")
}

func ConvertVirtualToPhysical() {
	InterpreterEx("!va2pa")
}

func MonitorVmcallExecution() {
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

func SaveMemoryContext() {
	InterpreterEx(".dump")
}

func ShowValueInDifferentFormats() {
	InterpreterEx(".formats")
}

func ShowCommandHelp() {
	InterpreterEx(".help")
}

func TerminateProcess() {
	InterpreterEx(".kill")
}

func ListenForClientConnection() {
	InterpreterEx(".listen")
}

func CloseLog() {
	InterpreterEx(".logclose")
}

func OpenLog() {
	InterpreterEx(".logopen")
}

func MakePageAvailableInRAM() {
	InterpreterEx(".pagein")
}

func ParseExecutableFiles() {
	InterpreterEx(".pe")
}

func ShowChangeProcesses() {
	InterpreterEx(".process")
}

func RestartProcess() {
	InterpreterEx(".restart")
}

func RunScript() {
	InterpreterEx(".script")
}

func RunsAUser_modeProcess() {
	InterpreterEx(".start")
}

func GetsTheStatusOf() {
	InterpreterEx(".status")
}

func ShowsTheListOf() {
	InterpreterEx(".switch")
}

func PerformsTheSymbolActions() {
	InterpreterEx(".sym")
}

func ShowsAndSetsThe() {
	InterpreterEx(".sympath")
}

func ShowsAndChangesThe() {
	InterpreterEx(".thread")
}
