package hyperdbgrust

type EventType uint32

const (
	EventTypeBreakpoint EventType = iota + 1
	EventTypeException
	EventTypeMemoryAccess
	EventTypeSyscallEntry
	EventTypeSyscallExit
	EventTypeProcessCreate
	EventTypeProcessExit
	EventTypeThreadCreate
	EventTypeThreadExit
	EventTypeModuleLoad
	EventTypeModuleUnload
	EventTypeDebugPrint
	EventTypeVmxExit
	EventTypeTrap
	EventTypeHiddenHookExec
	EventTypeHiddenHookRead
	EventTypeHiddenHookWrite
	EventTypeCpuid
	EventTypeTsc
	EventTypePmc
	EventTypeInterrupt
	EventTypeExceptionBitmap
	EventTypeCrAccess
	EventTypeDrAccess
	EventTypeIoPort
	EventTypeMsrRead
	EventTypeMsrWrite
	EventTypeEptViolation
	EventTypeVmcalled
)

type EventHeader struct {
	EventType EventType `json:"event_type"`
	ProcessID ProcessId `json:"process_id"`
	ThreadID  ThreadId  `json:"thread_id"`
	CoreID    uint32    `json:"core_id"`
	Timestamp uint64    `json:"timestamp"`
}

type DebuggerEventType int

const (
	DebuggerEventBreakpoint DebuggerEventType = iota
	DebuggerEventException
	DebuggerEventMemoryAccess
	DebuggerEventSyscall
	DebuggerEventProcessCreate
	DebuggerEventProcessExit
	DebuggerEventThreadCreate
	DebuggerEventThreadExit
	DebuggerEventModuleLoad
	DebuggerEventModuleUnload
	DebuggerEventDebugPrint
	DebuggerEventVmxExit
	DebuggerEventTrap
	DebuggerEventHiddenHook
	DebuggerEventCpuid
	DebuggerEventTsc
	DebuggerEventCrAccess
	DebuggerEventDrAccess
	DebuggerEventIoPort
	DebuggerEventMsr
	DebuggerEventEptViolation
)

type DebuggerEvent struct {
	Type DebuggerEventType `json:"type"`
	Data any               `json:"data"`
}

func (t DebuggerEventType) String() string {
	if name, ok := DebuggerEventTypeNames[t]; ok {
		return name
	}
	return "Unknown"
}

var DebuggerEventTypeNames = map[DebuggerEventType]string{
	DebuggerEventBreakpoint:    "Breakpoint",
	DebuggerEventException:     "Exception",
	DebuggerEventMemoryAccess:  "MemoryAccess",
	DebuggerEventSyscall:       "Syscall",
	DebuggerEventProcessCreate: "ProcessCreate",
	DebuggerEventProcessExit:   "ProcessExit",
	DebuggerEventThreadCreate:  "ThreadCreate",
	DebuggerEventThreadExit:    "ThreadExit",
	DebuggerEventModuleLoad:    "ModuleLoad",
	DebuggerEventModuleUnload:  "ModuleUnload",
	DebuggerEventDebugPrint:    "DebugPrint",
	DebuggerEventVmxExit:       "VmxExit",
	DebuggerEventTrap:          "Trap",
	DebuggerEventHiddenHook:    "HiddenHook",
	DebuggerEventCpuid:         "Cpuid",
	DebuggerEventTsc:           "Tsc",
	DebuggerEventCrAccess:      "CrAccess",
	DebuggerEventDrAccess:      "DrAccess",
	DebuggerEventIoPort:        "IoPort",
	DebuggerEventMsr:           "Msr",
	DebuggerEventEptViolation:  "EptViolation",
}

var DebuggerEventTypeByName = map[string]DebuggerEventType{
	"Breakpoint":    DebuggerEventBreakpoint,
	"Exception":     DebuggerEventException,
	"MemoryAccess":  DebuggerEventMemoryAccess,
	"Syscall":       DebuggerEventSyscall,
	"ProcessCreate": DebuggerEventProcessCreate,
	"ProcessExit":   DebuggerEventProcessExit,
	"ThreadCreate":  DebuggerEventThreadCreate,
	"ThreadExit":    DebuggerEventThreadExit,
	"ModuleLoad":    DebuggerEventModuleLoad,
	"ModuleUnload":  DebuggerEventModuleUnload,
	"DebugPrint":    DebuggerEventDebugPrint,
	"VmxExit":       DebuggerEventVmxExit,
	"Trap":          DebuggerEventTrap,
	"HiddenHook":    DebuggerEventHiddenHook,
	"Cpuid":         DebuggerEventCpuid,
	"Tsc":           DebuggerEventTsc,
	"CrAccess":      DebuggerEventCrAccess,
	"DrAccess":      DebuggerEventDrAccess,
	"IoPort":        DebuggerEventIoPort,
	"Msr":           DebuggerEventMsr,
	"EptViolation":  DebuggerEventEptViolation,
}
