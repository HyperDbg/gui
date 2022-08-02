package core

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\core\Debugger.h.back

const (
	DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL              = 0 //col:3
	DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT = 1 //col:4
	DEBUGGER_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED    = 2 //col:5
	DEBUGGER_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE      = 3 //col:6
)

type _DEBUGGER_CORE_EVENTS struct {
	HiddenHookReadAndWriteEventsHead     *list.List //col:24
	HiddenHookReadEventsHead             *list.List //col:25
	HiddenHookWriteEventsHead            *list.List //col:26
	EptHook2sExecDetourEventsHead        *list.List //col:27
	EptHookExecCcEventsHead              *list.List //col:28
	SyscallHooksEferSyscallEventsHead    *list.List //col:29
	SyscallHooksEferSysretEventsHead     *list.List //col:30
	CpuidInstructionExecutionEventsHead  *list.List //col:31
	RdmsrInstructionExecutionEventsHead  *list.List //col:32
	WrmsrInstructionExecutionEventsHead  *list.List //col:33
	ExceptionOccurredEventsHead          *list.List //col:34
	TscInstructionExecutionEventsHead    *list.List //col:35
	PmcInstructionExecutionEventsHead    *list.List //col:36
	InInstructionExecutionEventsHead     *list.List //col:37
	OutInstructionExecutionEventsHead    *list.List //col:38
	DebugRegistersAccessedEventsHead     *list.List //col:39
	ExternalInterruptOccurredEventsHead  *list.List //col:40
	VmcallInstructionExecutionEventsHead *list.List //col:41
	ControlRegisterModifiedEventsHead    *list.List //col:42
}

type _PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE struct {
	Msr   uint64 //col:29
	Value uint64 //col:30
}

type _DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE struct {
	WaitForInstrumentationStepInMtf bool   //col:34
	CsSel                           uint16 //col:35
}

type _DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS struct {
	InitialSetProcessChangeEvent             bool   //col:45
	InitialSetThreadChangeEvent              bool   //col:46
	InitialSetByClockInterrupt               bool   //col:47
	CurrentThreadLocationOnGs                uint64 //col:48
	DebugRegisterInterceptionState           bool   //col:49
	InterceptClockInterruptsForThreadChange  bool   //col:50
	IsWatingForMovCr3VmExits                 bool   //col:51
	InterceptClockInterruptsForProcessChange bool   //col:52
}

type _PROCESSOR_DEBUGGING_STATE struct {
	Lock                                    int32                                      //col:69
	WaitingToBeLocked                       bool                                       //col:70
	MainDebuggingCore                       bool                                       //col:71
	NmiCalledInVmxRootRelatedToHaltDebuggee bool                                       //col:72
	NmiBroadcastAction                      NMI_BROADCAST_ACTION_TYPE                  //col:73
	IgnoreEvent                             bool                                       //col:74
	IgnoreOneMtf                            bool                                       //col:75
	WaitForStepTrap                         bool                                       //col:76
	MsrState                                PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE      //col:77
	SoftwareBreakpointState                 PDEBUGGEE_BP_DESCRIPTOR                    //col:78
	InstrumentationStepInTrace              DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE     //col:79
	EnableExternalInterruptsOnContinue      bool                                       //col:80
	EnableExternalInterruptsOnContinueMtf   bool                                       //col:81
	DisableTrapFlagOnContinue               bool                                       //col:82
	DoNotNmiNotifyOtherCoresByThisCore      bool                                       //col:83
	ThreadOrProcessTracingDetails           DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS //col:84
	BreakStarterCore                        bool                                       //col:85
	InstructionLengthHint                   uint16                                     //col:86
	HardwareDebugRegisterForStepping        uint64                                     //col:87
	ScriptEngineCoreSpecificLocalVariable   *uint64                                    //col:88
	ScriptEngineCoreSpecificTempVariable    *uint64                                    //col:89
}

type _DEBUGGER_EVENT_ACTION struct {
	Tag                                       uint64                                         //col:83
	ActionOrderCode                           uint32                                         //col:84
	ActionsList                               *list.List                                     //col:85
	ActionType                                DEBUGGER_EVENT_ACTION_TYPE_ENUM                //col:86
	ImmediatelySendTheResults                 bool                                           //col:87
	DebuggerEventActionRunScriptConfiguration DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION //col:88
	ScriptConfiguration                       byte                                           //col:89
	DebuggerEventRequestBuffer                DEBUGGER_EVENT_REQUEST_BUFFER                  //col:90
	RequestedBuffer                           byte                                           //col:91
	CustomCodeBufferSize                      uint32                                         //col:92
	CustomCodeBufferAddress                   uintptr                                        //col:93
}

type _DEBUGGER_EVENT struct {
	Tag                    uint64                   //col:101
	EventsOfSameTypeList   *list.List               //col:102
	EventType              DEBUGGER_EVENT_TYPE_ENUM //col:103
	Enabled                bool                     //col:104
	CoreId                 uint32                   //col:105
	Uint32                 uint32                   //col:106
	ProcessId              byte                     //col:107
	ActionsListHead        *list.List               //col:108
	CountOfActions         uint32                   //col:109
	OptionalParam1         uint64                   //col:110
	OptionalParam2         uint64                   //col:111
	OptionalParam3         uint64                   //col:112
	OptionalParam4         uint64                   //col:113
	ConditionsBufferSize   uint32                   //col:114
	ConditionBufferAddress uintptr                  //col:115
}

