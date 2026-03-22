package pages

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/structview"
)

type EventRow struct {
	Tag       string
	Type      string
	ProcessId string
	ThreadId  string
	CoreId    string
	EIP       string
	Enabled   string
	Actions   string
}

type EventsPage struct {
	dbg    debugger.Eventer
	table  *structview.StructView
	events []debugger.Event
}

func NewEvents(dbg debugger.Eventer) *EventsPage {
	p := &EventsPage{
		dbg: dbg,
	}

	p.table = structview.New(EventRow{}, *button.NewButtonsLayout(button.LayoutButton{Button: button.Text(), Label: ""}))
	p.refreshEvents()
	return p
}

func (p *EventsPage) Self() any {
	return p
}

func (p *EventsPage) refreshEvents() {
	events := p.dbg.Events()

	p.events = events

	rows := make([]EventRow, len(events))
	for i, event := range events {
		rows[i] = EventRow{
			Tag:       fmt.Sprintf("0x%X", event.Address),
			Type:      p.getEventTypeName(event.Type),
			ProcessId: fmt.Sprintf("%d", event.Pid),
			ThreadId:  "",
			CoreId:    "",
			EIP:       fmt.Sprintf("0x%X", event.Address),
			Enabled:   "Yes",
			Actions:   "",
		}
	}

	if len(rows) > 0 {
		p.table.SetData(rows[0])
	}
}

func (p *EventsPage) getEventTypeName(eventType debugger.EventType) string {
	switch eventType {
	case debugger.EventStep:
		return "Step"
	case debugger.EventBreakpoint:
		return "Breakpoint"
	case debugger.EventException:
		return "Exception"
	case debugger.EventProcessCreated:
		return "Process Created"
	case debugger.EventProcessExited:
		return "Process Exited"
	case debugger.EventThreadCreated:
		return "Thread Created"
	case debugger.EventThreadExited:
		return "Thread Exited"
	case debugger.EventModuleLoaded:
		return "Module Loaded"
	case debugger.EventModuleUnloaded:
		return "Module Unloaded"
	case debugger.EptHookReadWriteAndExecute:
		return "EPT Hook R/W/X"
	case debugger.EptHookReadWrite:
		return "EPT Hook R/W"
	case debugger.EptHookReadAndExecute:
		return "EPT Hook R/X"
	case debugger.EptHookWriteAndExecute:
		return "EPT Hook W/X"
	case debugger.EptHookRead:
		return "EPT Hook R"
	case debugger.EptHookWrite:
		return "EPT Hook W"
	case debugger.EptHookExecute:
		return "EPT Hook X"
	case debugger.EptHookExecDetours:
		return "EPT Hook Detours"
	case debugger.EptHookExecCC:
		return "EPT Hook CC"
	case debugger.SyscallHookEferSyscall:
		return "Syscall Hook SYSCALL"
	case debugger.SyscallHookEferSysret:
		return "Syscall Hook SYSRET"
	case debugger.CpuidInstructionExecution:
		return "CPUID Execution"
	case debugger.RdmsrInstructionExecution:
		return "RDMSR Execution"
	case debugger.WrmsrInstructionExecution:
		return "WRMSR Execution"
	case debugger.InInstructionExecution:
		return "IN Execution"
	case debugger.OutInstructionExecution:
		return "OUT Execution"
	case debugger.ExceptionOccurred:
		return "Exception Occurred"
	case debugger.ExternalInterruptOccurred:
		return "External Interrupt"
	case debugger.DebugRegistersAccessed:
		return "Debug Registers"
	case debugger.TscInstructionExecution:
		return "TSC Execution"
	case debugger.PmcInstructionExecution:
		return "PMC Execution"
	case debugger.VmcallInstructionExecution:
		return "VMCALL Execution"
	case debugger.ControlRegisterModified:
		return "CR Modified"
	case debugger.ControlRegisterRead:
		return "CR Read"
	case debugger.ControlRegister3Modified:
		return "CR3 Modified"
	case debugger.TrapExecutionModeChanged:
		return "Trap Mode Changed"
	case debugger.TrapExecutionInstructionTrace:
		return "Trap Instruction Trace"
	case debugger.XsetbvInstructionExecution:
		return "XSETBV Execution"
	default:
		return "Unknown"
	}
}

func (p *EventsPage) Layout() layout.Widget {
	return p.table.Layout
}
