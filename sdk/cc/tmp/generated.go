package sdk

func ClearsBreakpointBreakpointId() {
Interpreter(StringToBytePointer("bc")) 
}

func DisablesBreakpointBreakpointId() {
Interpreter(StringToBytePointer("bd")) 
}

func EnablesBreakpointBreakpointId() {
Interpreter(StringToBytePointer("be")) 
}

func ListsAllEnabledAndDisabledBreakpoints() {
Interpreter(StringToBytePointer("bl")) 
}

func PutsBreakpoint0xcc() {
Interpreter(StringToBytePointer("bp")) 
}

func ShowsAndChangesOperatingProcessor() {
Interpreter(StringToBytePointer("core")) 
}

func CollectsReportFromCpuFeatures() {
Interpreter(StringToBytePointer("cpu")) 
}

func RemoveKernelModeBuffers() {
Interpreter(StringToBytePointer("flush")) 
}

func ContinueDebuggee() {
Interpreter(StringToBytePointer("g")) 
}

func ExecuteSingleInstructionStepOut() {
Interpreter(StringToBytePointer("gu")) 
}

func ExecuteSingleInstructionStepIn() {
Interpreter(StringToBytePointer("i")) 
}

func ShowThreadCallstack() {
Interpreter(StringToBytePointer("k, kd, kq")) 
}

func ListModules() {
Interpreter(StringToBytePointer("lm")) 
}

func InstallDriversAndLoadModules() {
Interpreter(StringToBytePointer("load")) 
}

func CreateEventForwardingOutputInstance() {
Interpreter(StringToBytePointer("output")) 
}

func ExecuteSingleInstructionStep() {
Interpreter(StringToBytePointer("p")) 
}

func PauseKernelEvents() {
Interpreter(StringToBytePointer("pause")) 
}

func PreactivateSpecialFunctionality() {
Interpreter(StringToBytePointer("preactivate")) 
}

func PreallocateBuffer() {
Interpreter(StringToBytePointer("prealloc")) 
}

func EvaluateExpressions() {
Interpreter(StringToBytePointer("print")) 
}

func ReadModifyRegisters() {
Interpreter(StringToBytePointer("r")) 
}

func ReadModelSpecificRegister() {
Interpreter(StringToBytePointer("rdmsr")) 
}

func SearchMemoryPattern() {
Interpreter(StringToBytePointer("sb, !sb, sd, !sd, sq, !sq")) 
}

func SettingsManagement() {
Interpreter(StringToBytePointer("settings")) 
}

func DebuggerSleep() {
Interpreter(StringToBytePointer("sleep")) 
}

func ExecuteStepIn() {
Interpreter(StringToBytePointer("t")) 
}

func TestHyperDbgFeatures() {
Interpreter(StringToBytePointer("test")) 
}

func UnloadKernelModules() {
Interpreter(StringToBytePointer("unload")) 
}

func SearchSymbols() {
Interpreter(StringToBytePointer("x")) 
}

func MonitorCpuidExecution() {
Interpreter(StringToBytePointer("!cpuid")) 
}

func MonitorControlRegisterModification() {
Interpreter(StringToBytePointer("!crwrite")) 
}

func MonitorDebugRegisters() {
Interpreter(StringToBytePointer("!dr")) 
}

func HiddenHookEPTBreakpoints() {
Interpreter(StringToBytePointer("!epthook")) 
}

func HiddenHookEPTDetours() {
Interpreter(StringToBytePointer("!epthook2")) 
}

func MonitorIDTEntries() {
Interpreter(StringToBytePointer("!exception")) 
}

func HideHyperDbg() {
Interpreter(StringToBytePointer("!hide")) 
}

func MonitorExternalInterrupts() {
Interpreter(StringToBytePointer("!interrupt")) 
}

func DetectIOInstructionsIn() {
Interpreter(StringToBytePointer("!ioin")) 
}

func DetectIOInstructionsOut() {
Interpreter(StringToBytePointer("!ioout")) 
}

func MeasureArgumentsForHide() {
Interpreter(StringToBytePointer("!measure")) 
}

func TrapModeInstructions() {
Interpreter(StringToBytePointer("!mode")) 
}

func MonitorMemoryAccess() {
Interpreter(StringToBytePointer("!monitor")) 
}

func DetectRdmsrExecution() {
Interpreter(StringToBytePointer("!msrread")) 
}

func DetectWrmsrExecution() {
Interpreter(StringToBytePointer("!msrwrite")) 
}

func ConvertPhysicalToVirtual() {
Interpreter(StringToBytePointer("!pa2va")) 
}

func MonitorRdpmcExecution() {
Interpreter(StringToBytePointer("!pmc")) 
}

func FindVirtualAddressPagingLevels() {
Interpreter(StringToBytePointer("!pte")) 
}

func UseReversingMachineModule() {
Interpreter(StringToBytePointer("!rev")) 
}

func MonitorSyscallExecution() {
Interpreter(StringToBytePointer("!syscall")) 
}

func MonitorSysretExecution() {
Interpreter(StringToBytePointer("!sysret")) 
}

func TraceExecution() {
Interpreter(StringToBytePointer("!trace")) 
}

func TrackModeTransitionInstructions() {
Interpreter(StringToBytePointer("!track")) 
}

func MonitorRdtscInstructions() {
Interpreter(StringToBytePointer("!tsc")) 
}

func RevertHide() {
Interpreter(StringToBytePointer("!unhide")) 
}

func ConvertVirtualToPhysical() {
Interpreter(StringToBytePointer("!va2pa")) 
}

func MonitorVmcallExecution() {
Interpreter(StringToBytePointer("!vmcall")) 
}

func HardwareClockDebugging() {
Interpreter(StringToBytePointer("!hw_clk")) 
}

func AttachDebugThread() {
Interpreter(StringToBytePointer(".attach")) 
}

func ClearScreen() {
Interpreter(StringToBytePointer(".cls")) 
}

func ConnectToMachine() {
Interpreter(StringToBytePointer(".connect")) 
}

func DebugMachine() {
Interpreter(StringToBytePointer(".debug")) 
}

func DetachDebugging() {
Interpreter(StringToBytePointer(".detach")) 
}

func DisconnectSession() {
Interpreter(StringToBytePointer(".disconnect")) 
}

func SaveMemoryContext() {
Interpreter(StringToBytePointer(".dump")) 
}

func ShowValueInDifferentFormats() {
Interpreter(StringToBytePointer(".formats")) 
}

func ShowCommandHelp() {
Interpreter(StringToBytePointer(".help")) 
}

func TerminateProcess() {
Interpreter(StringToBytePointer(".kill")) 
}

func ListenForClientConnection() {
Interpreter(StringToBytePointer(".listen")) 
}

func CloseLog() {
Interpreter(StringToBytePointer(".logclose")) 
}

func OpenLog() {
Interpreter(StringToBytePointer(".logopen")) 
}

func MakePageAvailableInRAM() {
Interpreter(StringToBytePointer(".pagein")) 
}

func ParseExecutableFiles() {
Interpreter(StringToBytePointer(".pe")) 
}

func ShowChangeProcesses() {
Interpreter(StringToBytePointer(".process")) 
}

func RestartProcess() {
Interpreter(StringToBytePointer(".restart")) 
}

func RunScript() {
Interpreter(StringToBytePointer(".script")) 
}

func RunsAUser-modeProcess() {
Interpreter(StringToBytePointer(".start")) 
}

func GetsTheStatusOf() {
Interpreter(StringToBytePointer(".status")) 
}

func ShowsTheListOf() {
Interpreter(StringToBytePointer(".switch")) 
}

func PerformsTheSymbolActions() {
Interpreter(StringToBytePointer(".sym")) 
}

func ShowsAndSetsThe() {
Interpreter(StringToBytePointer(".sympath")) 
}

func ShowsAndChangesThe() {
Interpreter(StringToBytePointer(".thread")) 
}

