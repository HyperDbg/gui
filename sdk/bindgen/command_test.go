package bindgen

import (
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/HyperDbg/sdk"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/txt"
)

func TestUnmarshalCommandJson(t *testing.T) {
	var generated []sdk.Command
	mylog.Check(json.Unmarshal(stream.NewBuffer("commands.json").Bytes(), &generated))
	// mylog.Struct(generated)
	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	for _, s := range generated {
		fullName := s.FullName
		// fullName = strings.ReplaceAll(fullName, "-", "_")
		// fullName = stream.ToCamelUpper(fullName, false)
		// fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepIn", "StepIn")
		// fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepOut", "StepOut")
		// fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStep", "StepExecuteSingleInstruction") //? not well
		// fullName = strings.ReplaceAll(fullName, "CreateEventForwardingOutputInstance", "OutputEventForwardingInstance")
		// fullName = strings.ReplaceAll(fullName, "ClearsBreakpointBreakpointId", "BreakpointClearByID")
		// fullName = strings.ReplaceAll(fullName, "DisablesBreakpointBreakpointId", "BreakpointDisableByID")
		// fullName = strings.ReplaceAll(fullName, "EnablesBreakpointBreakpointId", "BreakpointEnableByID")
		// fullName = strings.ReplaceAll(fullName, "ListsAllEnabledAndDisabledBreakpoints", "BreakpointList")
		// fullName = strings.ReplaceAll(fullName, "PutsBreakpoint0xcc", "SetBreakpoint0xcc")
		// fullName = strings.ReplaceAll(fullName, "ShowsAndChangesOperatingProcessor", "CoreOperatingProcessorForShowsAndChanges") // too long
		// fullName = strings.ReplaceAll(fullName, "CollectsReportFromCpuFeatures", "CpuFeaturesForCollectsReport")
		// fullName = strings.ReplaceAll(fullName, "RemoveKernelModeBuffers", "FlushKernelModeBuffers")
		// fullName = strings.ReplaceAll(fullName, "ContinueDebuggee", "ContinueDebugger")
		// fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepOut", "StepOut")
		// fullName = strings.ReplaceAll(fullName, "ExecuteSingleInstructionStepIn", "StepIn") // repeated with StepInExecute? "t"
		// fullName = strings.ReplaceAll(fullName, "ShowThreadCallstack", "CallstackOrThreadView")
		// fullName = strings.ReplaceAll(fullName, "InstallDriversAndLoadModules", "LoadDriversAndModules")
		// fullName = strings.ReplaceAll(fullName, "EvaluateExpressions", "PrintEvaluateExpressions")
		// fullName = strings.ReplaceAll(fullName, "ReadModifyRegisters", "RegistersReadOrModify")
		// fullName = strings.ReplaceAll(fullName, "ReadModelSpecificRegister", "ReadMsr") //?? repeated?
		// fullName = strings.ReplaceAll(fullName, "DetectRdmsrExecution", "MsrRead")
		// fullName = strings.ReplaceAll(fullName, "DetectWrmsrExecution", "MsrWrite")
		// fullName = strings.ReplaceAll(fullName, "MonitorSyscallExecution", "Syscall")
		// fullName = strings.ReplaceAll(fullName, "MonitorSysretExecution", "SysRet")
		// fullName = strings.ReplaceAll(fullName, "ShowsAndChangesThe", "Thread")
		// fullName = strings.ReplaceAll(fullName, "ShowsAndSetsThe", "SymbolPath")
		// fullName = strings.ReplaceAll(fullName, "PerformsTheSymbolActions", "SymbolActions")
		// fullName = strings.ReplaceAll(fullName, "ShowsTheListOf", "SwitchThread")
		// fullName = strings.ReplaceAll(fullName, "GetsTheStatusOf", "Status")
		// fullName = strings.ReplaceAll(fullName, "RunScript", "Script")
		// fullName = strings.ReplaceAll(fullName, "ShowChangeProcesses", "ProcessesView")

		// fullName = strings.ReplaceAll(fullName, "OpenLog", "LogOpen")   //?
		// fullName = strings.ReplaceAll(fullName, "CloseLog", "LogClose") //?
		// fullName = strings.ReplaceAll(fullName, "TerminateProcess", "KillProcess")
		// fullName = strings.ReplaceAll(fullName, "ShowCommandHelp", "HelpForCommand")
		// fullName = strings.ReplaceAll(fullName, "ShowValueInDifferentFormats", "FormatsDiff")
		// fullName = strings.ReplaceAll(fullName, "SaveMemoryContext", "DumpMemoryContext")
		// fullName = strings.ReplaceAll(fullName, "MonitorVmcallExecution", "VmCall")
		// fullName = strings.ReplaceAll(fullName, "ConvertVirtualToPhysical", "Va2Pa")
		// fullName = strings.ReplaceAll(fullName, "ConvertPhysicalToVirtual", "Pa2Va")
		// fullName = strings.ReplaceAll(fullName, "RevertHide", "UnHide")
		// fullName = strings.ReplaceAll(fullName, "MonitorRdtscInstructions", "TscInstructionsMonitor")
		// fullName = strings.ReplaceAll(fullName, "UseReversingMachineModule", "ReversingMachineModuleUse")
		// fullName = strings.ReplaceAll(fullName, "FindVirtualAddressPagingLevels", "Pte")
		// fullName = strings.ReplaceAll(fullName, "MonitorRdpmcExecution", "PmcExecutionMonitor")
		// fullName = strings.ReplaceAll(fullName, "MonitorExternalInterrupts", "InterruptExternalMonitor")
		println(fullName)
		switch fullName {
		case "HiddenHookEPTBreakpoints":
			fullName = "EptHook"
			fallthrough
		case "HiddenHookEPTDetours":
			fullName = "EptHook2"
			fallthrough
		case "MonitorIdtEntries":
			fullName = "IdtEntriesMonitor"
			fallthrough
		case "RunsAUserModeProcess":
			fullName = "StartProcess"
			fallthrough
		case "MakePageAvailableInRam":
			fullName = "PageAvailableInRam"
			fallthrough
		case "DetectIoInstructionsIn":
			fullName = "IoInDetect"
		}

		//fullName = strings.ReplaceAll(fullName, "MonitorIdtEntries", "IdtEntriesMonitor")
		//fullName = strings.ReplaceAll(fullName, "RunsAUserModeProcess", "StartProcess")
		//fullName = strings.ReplaceAll(fullName, "MakePageAvailableInRam", "PageAvailableInRam")
		//fullName = strings.ReplaceAll(fullName, "DetectIoInstructionsIn", "IoInDetect")
		//

		//if fullName == "HiddenHookEPTBreakpoints" {
		//	println(fullName)
		//}

		// fullName = strings.ReplaceAll(fullName, "HiddenHookEptDetours", "EptHook2")    // cc
		// fullName = strings.ReplaceAll(fullName, "HiddenHookEptBreakpoints", "EptHook") // inline
		// fullName = strings.ReplaceAll(fullName, "MonitorDebugRegisters", "DebugRegistersMonitor")
		// fullName = strings.ReplaceAll(fullName, "MonitorControlRegisterModification", "ControlRegisterModificationMonitor") // too long
		// fullName = strings.ReplaceAll(fullName, "MonitorCpuidExecution", "CpuidExecutionMonitor")
		// fullName = strings.ReplaceAll(fullName, "ExecuteStepIn", "StepInExecute") // f7
		// fullName = strings.ReplaceAll(fullName, "DebuggerSleep", "SleepDebugger")
		// fullName = strings.ReplaceAll(fullName, "TrapModeInstructions", "ModeInstructionsTrap")
		// fullName = strings.ReplaceAll(fullName, "DetectIoInstructionsOut", "IoOutDetect")

		s.Description = strings.TrimSuffix(s.Description, "\n")
		g.P("//", fullName)
		g.P("//", "Description:", s.Description)
		g.P("//", "Syntax:")
		for _, syntax := range s.Syntax {
			g.P("//", syntax)
		}
		if len(s.Examples) > 0 {
			g.P("//", "Examples:")
			for _, example := range s.Examples {
				g.P("//", example)
			}
		}
		if len(s.Notes) > 0 {
			g.P("//", "Notes:")
			for _, note := range s.Notes {
				g.P("//", note)
			}
		}
		// println(fullName)
		g.P("func ", fullName, "()(status string) {")
		g.P("return ", "InterpreterEx(", strconv.Quote(s.Name), ") ")
		g.P("}")
		g.P()
	}
	stream.WriteGoFile("../commandWrapper.go", g.Bytes())
}

func TestCommandGenerate(t *testing.T) {
	t.Skip("not well")
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").Enum("commands", []string{
		"debugging",
		"extension",
		"hwdbg",
		"meta",
	}, nil)

	commandGenerate("DebuggingCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\debugging-commands")
	commandGenerate("ExtensionCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\extension-commands")
	commandGenerate("HwdbgCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\hwdbg-commands")
	commandGenerate("MetaCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\meta-commands")
}

func commandGenerate(kindName, path string) {
	var commands []sdk.Command
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		name := stream.BaseName(path)
		name = strings.ReplaceAll(name, "-", "_") // not well
		name = strings.ReplaceAll(name, "~", "unknown")
		name = txt.ToCamelCase(name)
		mylog.Trace(name, path)

		command := sdk.Command{
			Name:        "",
			Description: "",
			Syntax:      nil,
			Examples:    nil,
			Notes:       nil,
			FullName:    "",
		}
		commands = append(commands, command)
		return err
	})

	keys := make([]string, 0)
	for _, command := range commands {
		keys = append(keys, command.FullName) // todo this must be use Cmd,then it will return the right command
	}
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").Enum(kindName, keys, nil)
}
