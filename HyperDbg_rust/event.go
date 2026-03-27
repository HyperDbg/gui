package hyperdbgrust

type EventType uint32

const (
	EventTypeBreakpoint      EventType = 1
	EventTypeException       EventType = 2
	EventTypeMemoryAccess    EventType = 3
	EventTypeSyscallEntry    EventType = 4
	EventTypeSyscallExit     EventType = 5
	EventTypeProcessCreate   EventType = 6
	EventTypeProcessExit     EventType = 7
	EventTypeThreadCreate    EventType = 8
	EventTypeThreadExit      EventType = 9
	EventTypeModuleLoad      EventType = 10
	EventTypeModuleUnload    EventType = 11
	EventTypeDebugPrint      EventType = 12
	EventTypeVmxExit         EventType = 13
	EventTypeTrap            EventType = 14
	EventTypeHiddenHookExec  EventType = 15
	EventTypeHiddenHookRead  EventType = 16
	EventTypeHiddenHookWrite EventType = 17
	EventTypeCpuid           EventType = 18
	EventTypeTsc             EventType = 19
	EventTypePmc             EventType = 20
	EventTypeInterrupt       EventType = 21
	EventTypeExceptionBitmap EventType = 22
	EventTypeCrAccess        EventType = 23
	EventTypeDrAccess        EventType = 24
	EventTypeIoPort          EventType = 25
	EventTypeMsrRead         EventType = 26
	EventTypeMsrWrite        EventType = 27
	EventTypeEptViolation    EventType = 28
	EventTypeVmcalled        EventType = 29
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
