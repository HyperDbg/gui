package sdk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

// m := orderedmap.New(InvalidCommandsKind, []Command{})
// m.Set(DebuggingKind, commands)
func LayoutCommands() unison.Paneler {
	return widget.NewTableScroll(Command{}, widget.TableContext[Command]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Command]) (cells []widget.CellData) {
			if node.Container() {
				node.Data.MethodName = node.Sum(node.Data.MethodName)
			}
			return []widget.CellData{
				{Text: node.Data.MethodName},
				{Text: node.Data.Cmd},
				{Text: fmt.Sprint(node.Data.Args)},
				{Text: node.Data.DoFunc},
				{Text: node.Data.Usage},
				{Text: fmt.Sprint(node.Data.Demo)},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Command]) {
			containerNodes := widget.NewContainerNodes[Command](InvalidCommandsKind.Keys())
			root.SetChildren(containerNodes)
			for _, node := range containerNodes {
				switch strings.TrimSuffix(node.Type, widget.ContainerKeyPostfix) {
				case DebuggingKind.String():
					for _, s := range InvalidDebuggingCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case ExtensionKind.String():
					for _, s := range InvalidExtensionCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case HwdbgKind.String():
					for _, s := range InvalidHwdbgCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case MetaKind.String():
					for _, s := range InvalidMetaCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				default:
				}
			}
		},
		JsonName:   "commands",
		IsDocument: false,
	})
}

type (
	Command struct {
		MethodName string
		Cmd        string // todo replace with kind
		Args       []string
		DoFunc     string
		Usage      string
		Demo       []string
	}
)

// now the ui will call these command functions to execute commands

// RunCommand todo generate them into command_gen.go, need rename part of command name
func (k DebuggingCommandsKind) RunCommand() bool {
	switch k {
	case BcKind:
	case BdKind:
	case BeKind:
	case BlKind:
	case BpKind:
		Interpreter(StringToBytePointer("bp"))       // todo remove
		Interpreter(StringToBytePointer(k.String())) // need like this
	case CpuKind:
	case DuKind:
	case DtStructKind:
	case EKind:
	case EvalKind:
	case EventsKind:
	case ExitKind:
	case FlushKind:
	case GKind:
	case GuKind:
	case IKind:
	case KKind:
	case LmKind:
	case LoadKind:
	case OutputKind:
	case PKind:
	case PauseKind:
	case PreactivateKind:
	case PreallocKind:
	case PrintKind:
	case RKind:
	case RdmsrKind:
	case SKind:
	case SettingsKind:
	case SleepKind:
	case TKind:
	case TestKind:
	case UnloadKind:
	case WrmsrKind:
	case XKind:
	case UnknownKind:
	case InvalidDebuggingCommandsKind:
	}
	return false
}

func (k DebuggingCommandsKind) Usage() string {
	switch k {
	case BcKind:
	case BdKind:
	case BeKind:
	case BlKind:
	case BpKind:
	case CpuKind:
	case DuKind:
	case DtStructKind:
	case EKind:
	case EvalKind:
	case EventsKind:
	case ExitKind:
	case FlushKind:
	case GKind:
	case GuKind:
	case IKind:
	case KKind:
	case LmKind:
	case LoadKind:
	case OutputKind:
	case PKind:
	case PauseKind:
	case PreactivateKind:
	case PreallocKind:
	case PrintKind:
	case RKind:
	case RdmsrKind:
	case SKind:
	case SettingsKind:
	case SleepKind:
	case TKind:
	case TestKind:
	case UnloadKind:
	case WrmsrKind:
	case XKind:
	case UnknownKind:
	case InvalidDebuggingCommandsKind:
	}
	return k.String()
}

func (k DebuggingCommandsKind) Demo() []string {
	switch k {
	case BcKind:
	case BdKind:
	case BeKind:
	case BlKind:
	case BpKind:
	case CpuKind:
	case DuKind:
	case DtStructKind:
	case EKind:
	case EvalKind:
	case EventsKind:
	case ExitKind:
	case FlushKind:
	case GKind:
	case GuKind:
	case IKind:
	case KKind:
	case LmKind:
	case LoadKind:
	case OutputKind:
	case PKind:
	case PauseKind:
	case PreactivateKind:
	case PreallocKind:
	case PrintKind:
	case RKind:
	case RdmsrKind:
	case SKind:
	case SettingsKind:
	case SleepKind:
	case TKind:
	case TestKind:
	case UnloadKind:
	case WrmsrKind:
	case XKind:
	case UnknownKind:
	case InvalidDebuggingCommandsKind:
	}
	return nil
}

func (k ExtensionCommandsKind) RunCommand() bool {
	switch k {
	case CpuidKind:
	case CrwriteKind:
	case DrKind:
	case EpthookKind:
	case Epthook2Kind:
	case ExceptionKind:
	case HideKind:
	case InterruptKind:
	case IoinKind:
	case IooutKind:
	case MeasureKind:
	case ModeKind:
	case MonitorKind:
	case MsrreadKind:
	case MsrwriteKind:
	case Pa2vaKind:
	case PmcKind:
	case PteKind:
	case RevKind:
	case SyscallSysretKind:
	case TraceKind:
	case TrackKind:
	case TscKind:
	case UnhideKind:
	case Va2paKind:
	case VmcallKind:
	case InvalidExtensionCommandsKind:
	}
	return false
}

func (k ExtensionCommandsKind) Usage() {
	switch k {
	case CpuidKind:
	case CrwriteKind:
	case DrKind:
	case EpthookKind:
	case Epthook2Kind:
	case ExceptionKind:
	case HideKind:
	case InterruptKind:
	case IoinKind:
	case IooutKind:
	case MeasureKind:
	case ModeKind:
	case MonitorKind:
	case MsrreadKind:
	case MsrwriteKind:
	case Pa2vaKind:
	case PmcKind:
	case PteKind:
	case RevKind:
	case SyscallSysretKind:
	case TraceKind:
	case TrackKind:
	case TscKind:
	case UnhideKind:
	case Va2paKind:
	case VmcallKind:
	case InvalidExtensionCommandsKind:
	}
}

func (k ExtensionCommandsKind) Demo() {
	switch k {
	case CpuidKind:
	case CrwriteKind:
	case DrKind:
	case EpthookKind:
	case Epthook2Kind:
	case ExceptionKind:
	case HideKind:
	case InterruptKind:
	case IoinKind:
	case IooutKind:
	case MeasureKind:
	case ModeKind:
	case MonitorKind:
	case MsrreadKind:
	case MsrwriteKind:
	case Pa2vaKind:
	case PmcKind:
	case PteKind:
	case RevKind:
	case SyscallSysretKind:
	case TraceKind:
	case TrackKind:
	case TscKind:
	case UnhideKind:
	case Va2paKind:
	case VmcallKind:
	case InvalidExtensionCommandsKind:
	}
}

func (k HwdbgCommandsKind) RunCommand() bool {
	switch k {
	case HwClkKind:
	case InvalidHwdbgCommandsKind:
	}
	return false
}

func (k HwdbgCommandsKind) Usage() {
	switch k {
	case HwClkKind:
	case InvalidHwdbgCommandsKind:
	}
}

func (k HwdbgCommandsKind) Demo() {
	switch k {
	case HwClkKind:
	case InvalidHwdbgCommandsKind:
	}
}

func (k MetaCommandsKind) RunCommand() bool {
	switch k {
	case AttachKind:
	case ClsKind:
	case ConnectKind:
	case DebugKind:
	case DetachKind:
	case DisconnectKind:
	case DumpKind:
	case FormatsKind:
	case HelpKind:
	case KillKind:
	case ListenKind:
	case LogcloseKind:
	case LogopenKind:
	case PageinKind:
	case PeKind:
	case ProcessKind:
	case RestartKind:
	case ScriptKind:
	case StartKind:
	case StatusKind:
	case SwitchKind:
	case SymKind:
	case SympathKind:
	case ThreadKind:
	case InvalidMetaCommandsKind:
	}
	return false
}

func (k MetaCommandsKind) Usage() {
	switch k {
	case AttachKind:
	case ClsKind:
	case ConnectKind:
	case DebugKind:
	case DetachKind:
	case DisconnectKind:
	case DumpKind:
	case FormatsKind:
	case HelpKind:
	case KillKind:
	case ListenKind:
	case LogcloseKind:
	case LogopenKind:
	case PageinKind:
	case PeKind:
	case ProcessKind:
	case RestartKind:
	case ScriptKind:
	case StartKind:
	case StatusKind:
	case SwitchKind:
	case SymKind:
	case SympathKind:
	case ThreadKind:
	case InvalidMetaCommandsKind:
	}
}

func (k MetaCommandsKind) Demo() {
	switch k {
	case AttachKind:
	case ClsKind:
	case ConnectKind:
	case DebugKind:
	case DetachKind:
	case DisconnectKind:
	case DumpKind:
	case FormatsKind:
	case HelpKind:
	case KillKind:
	case ListenKind:
	case LogcloseKind:
	case LogopenKind:
	case PageinKind:
	case PeKind:
	case ProcessKind:
	case RestartKind:
	case ScriptKind:
	case StartKind:
	case StatusKind:
	case SwitchKind:
	case SymKind:
	case SympathKind:
	case ThreadKind:
	case InvalidMetaCommandsKind:
	}
}

/*

	CommandBpHelp()
	{
	    ShowMessages("bp : puts a breakpoint (0xcc).\n");

	    ShowMessages(
	        "Note : 'bp' is not an event, if you want to use an event version "
	        "of breakpoints use !epthook or !epthook2 instead. See "
	        "documentation for more information.\n\n");

	    ShowMessages("syntax : \tbp [Address (hex)] [pid ProcessId (hex)] [tid ThreadId (hex)] [core CoreId (hex)]\n");

	    ShowMessages("\n");
	    ShowMessages("\t\te.g : bp nt!ExAllocatePoolWithTag\n");
	    ShowMessages("\t\te.g : bp nt!ExAllocatePoolWithTag+5\n");
	    ShowMessages("\t\te.g : bp nt!ExAllocatePoolWithTag+@rcx+rbx\n");
	    ShowMessages("\t\te.g : bp fffff8077356f010\n");
	    ShowMessages("\t\te.g : bp fffff8077356f010 pid 0x4\n");
	    ShowMessages("\t\te.g : bp fffff8077356f010 tid 0x1000\n");
	    ShowMessages("\t\te.g : bp fffff8077356f010 pid 0x4 core 2\n");
	}

*/
