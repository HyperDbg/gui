package cc

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

type Command struct {
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Syntax      []string `json:"Syntax"`
	Examples    []string `json:"Examples"`
	Notes       string   `json:"Notes"`
}

type Commands struct {
	Debugging []Command
	Extension []Command
	Hwdbg     []Command
	Meta      []Command
}

func TestUnmarshalCommandJson(t *testing.T) {
	type AutoGenerated struct {
		Name        string   `json:"Name"`
		Description string   `json:"Description"`
		Syntax      []string `json:"Syntax"`
		Examples    []string `json:"Examples"`
		Notes       []any    `json:"Notes"`
		FullName    string   `json:"FullName"`
	}
	var generated []AutoGenerated
	mylog.Check(json.Unmarshal(stream.NewBuffer("sina2.json").Bytes(), &generated))
	// mylog.Struct(generated)
	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	for _, s := range generated {
		fullName := s.FullName
		fullName = strings.ReplaceAll(fullName, "-", "_")
		fullName = stream.ToCamelUpper(fullName, false)
		fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepIn", "StepIn")
		fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepOut", "StepOut")
		fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStep", "StepExecuteSingleInstruction") //? not well
		fullName = strings.ReplaceAll(fullName, "CreateEventForwardingOutputInstance", "OutputEventForwardingInstance")
		fullName = strings.ReplaceAll(fullName, "ClearsBreakpointBreakpointId", "BreakpointClearByID")
		fullName = strings.ReplaceAll(fullName, "DisablesBreakpointBreakpointId", "BreakpointDisableByID")
		fullName = strings.ReplaceAll(fullName, "EnablesBreakpointBreakpointId", "BreakpointEnableByID")
		fullName = strings.ReplaceAll(fullName, "ListsAllEnabledAndDisabledBreakpoints", "BreakpointList")
		fullName = strings.ReplaceAll(fullName, "PutsBreakpoint0xcc", "SetBreakpoint0xcc")
		fullName = strings.ReplaceAll(fullName, "ShowsAndChangesOperatingProcessor", "CoreOperatingProcessorForShowsAndChanges") // too long
		fullName = strings.ReplaceAll(fullName, "CollectsReportFromCpuFeatures", "CpuFeaturesForCollectsReport")
		fullName = strings.ReplaceAll(fullName, "RemoveKernelModeBuffers", "FlushKernelModeBuffers")
		fullName = strings.ReplaceAll(fullName, "ContinueDebuggee", "ContinueDebugger")
		fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepOut", "StepOut")
		fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepIn", "StepIn") // repeated with StepInExecute? "t"
		fullName = strings.ReplaceAll(fullName, "ShowThreadCallstack", "CallstackOrThreadView")
		fullName = strings.ReplaceAll(fullName, "InstallDriversAndLoadModules", "LoadDriversAndModules")
		fullName = strings.ReplaceAll(fullName, "EvaluateExpressions", "PrintEvaluateExpressions")
		fullName = strings.ReplaceAll(fullName, "ReadModifyRegisters", "RegistersReadOrModify")
		fullName = strings.ReplaceAll(fullName, "ReadModelSpecificRegister", "ReadMsr") //?? repeated?
		fullName = strings.ReplaceAll(fullName, "DetectRdmsrExecution", "MsrRead")
		fullName = strings.ReplaceAll(fullName, "DetectWrmsrExecution", "MsrWrite")
		fullName = strings.ReplaceAll(fullName, "MonitorSyscallExecution", "Syscall")
		fullName = strings.ReplaceAll(fullName, "MonitorSysretExecution", "SysRet")
		fullName = strings.ReplaceAll(fullName, "ShowsAndChangesThe", "Thread")
		fullName = strings.ReplaceAll(fullName, "ShowsAndSetsThe", "SymbolPath")
		fullName = strings.ReplaceAll(fullName, "PerformsTheSymbolActions", "Symbol")
		fullName = strings.ReplaceAll(fullName, "ShowsTheListOf", "SwitchThread")
		fullName = strings.ReplaceAll(fullName, "GetsTheStatusOf", "Status")
		fullName = strings.ReplaceAll(fullName, "RunsAUserModeProcess", "StartProcess")
		fullName = strings.ReplaceAll(fullName, "RunScript", "Script")
		fullName = strings.ReplaceAll(fullName, "ShowChangeProcesses", "ProcessesView")
		fullName = strings.ReplaceAll(fullName, "MakePageAvailableInRam", "PageAvailableInRam")
		fullName = strings.ReplaceAll(fullName, "OpenLog", "LogOpen")   //?
		fullName = strings.ReplaceAll(fullName, "CloseLog", "LogClose") //?
		fullName = strings.ReplaceAll(fullName, "TerminateProcess", "KillProcess")
		fullName = strings.ReplaceAll(fullName, "ShowCommandHelp", "HelpForCommand")
		fullName = strings.ReplaceAll(fullName, "ShowValueInDifferentFormats", "FormatsDiff")
		fullName = strings.ReplaceAll(fullName, "SaveMemoryContext", "DumpMemoryContext")
		fullName = strings.ReplaceAll(fullName, "MonitorVmcallExecution", "VmCall")
		fullName = strings.ReplaceAll(fullName, "ConvertVirtualToPhysical", "Va2Pa")
		fullName = strings.ReplaceAll(fullName, "ConvertPhysicalToVirtual", "Pa2Va")
		fullName = strings.ReplaceAll(fullName, "RevertHide", "UnHide")
		fullName = strings.ReplaceAll(fullName, "MonitorRdtscInstructions", "TscInstructionsMonitor")
		fullName = strings.ReplaceAll(fullName, "UseReversingMachineModule", "ReversingMachineModuleUse")
		fullName = strings.ReplaceAll(fullName, "FindVirtualAddressPagingLevels", "Pte")
		fullName = strings.ReplaceAll(fullName, "MonitorRdpmcExecution", "PmcExecutionMonitor")
		fullName = strings.ReplaceAll(fullName, "DetectIoInstructionsIn", "IoInDetect")
		fullName = strings.ReplaceAll(fullName, "MonitorExternalInterrupts", "InterruptExternalMonitor")
		fullName = strings.ReplaceAll(fullName, "MonitorIdtEntries", "IdtEntriesMonitor")
		fullName = strings.ReplaceAll(fullName, "HiddenHookEptDetours", "EptHook2")    // cc
		fullName = strings.ReplaceAll(fullName, "HiddenHookEptBreakpoints", "EptHook") // inline
		fullName = strings.ReplaceAll(fullName, "MonitorDebugRegisters", "DebugRegistersMonitor")
		fullName = strings.ReplaceAll(fullName, "MonitorControlRegisterModification", "ControlRegisterModificationMonitor") // too long
		fullName = strings.ReplaceAll(fullName, "MonitorCpuidExecution", "CpuidExecutionMonitor")
		fullName = strings.ReplaceAll(fullName, "ExecuteStepIn", "StepInExecute") // f7
		fullName = strings.ReplaceAll(fullName, "DebuggerSleep", "SleepDebugger")
		fullName = strings.ReplaceAll(fullName, "TrapModeInstructions", "ModeInstructionsTrap")
		fullName = strings.ReplaceAll(fullName, "DetectIoInstructionsOut", "IoOutDetect")

		g.P("//", fullName)
		g.P("//", "Syntax:")
		for _, syntax := range s.Syntax {
			g.P("//", syntax)
		}

		g.P("func ", fullName, "() { //", s.Description)
		g.P("InterpreterEx(", strconv.Quote(s.Name), ") ")
		g.P("}")
		g.P()
	}
	stream.WriteGoFile("../commandWrapper.go", g.Bytes())
}
