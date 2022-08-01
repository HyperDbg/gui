package core
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\core\Debugger.h.back

const(
DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER = 0 //col:1
DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT = 1 //col:2
DEBUGGER_DEBUG_REGISTER_FOR_USER_MODE_ENTRY_POINT = 1 //col:3
)

const(
    DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL               =  0  //col:3
    DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT  =  1  //col:4
    DEBUGGER_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED     =  2  //col:5
    DEBUGGER_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE       =  3  //col:6
)



type DEBUGGER_CORE_EVENTS struct{
HiddenHookReadAndWriteEventsHead *list.List //col:3
HiddenHookReadEventsHead *list.List //col:4
HiddenHookWriteEventsHead *list.List //col:5
EptHook2sExecDetourEventsHead *list.List //col:6
EptHookExecCcEventsHead *list.List //col:7
SyscallHooksEferSyscallEventsHead *list.List //col:8
SyscallHooksEferSysretEventsHead *list.List //col:9
CpuidInstructionExecutionEventsHead *list.List //col:10
RdmsrInstructionExecutionEventsHead *list.List //col:11
WrmsrInstructionExecutionEventsHead *list.List //col:12
ExceptionOccurredEventsHead *list.List //col:13
TscInstructionExecutionEventsHead *list.List //col:14
PmcInstructionExecutionEventsHead *list.List //col:15
InInstructionExecutionEventsHead *list.List //col:16
OutInstructionExecutionEventsHead *list.List //col:17
DebugRegistersAccessedEventsHead *list.List //col:18
ExternalInterruptOccurredEventsHead *list.List //col:19
VmcallInstructionExecutionEventsHead *list.List //col:20
ControlRegisterModifiedEventsHead *list.List //col:21
}


type PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE struct{
Msr uint64 //col:25
Value uint64 //col:26
}


type DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE struct{
WaitForInstrumentationStepInMtf bool //col:30
CsSel uint16 //col:31
}


type DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS struct{
InitialSetProcessChangeEvent bool //col:35
InitialSetThreadChangeEvent bool //col:36
InitialSetByClockInterrupt bool //col:37
CurrentThreadLocationOnGs uint64 //col:38
DebugRegisterInterceptionState bool //col:39
InterceptClockInterruptsForThreadChange bool //col:40
IsWatingForMovCr3VmExits bool //col:41
InterceptClockInterruptsForProcessChange bool //col:42
}


type PROCESSOR_DEBUGGING_STATE struct{
LONG volatile //col:46
BOOLEAN volatile //col:47
BOOLEAN volatile //col:48
BOOLEAN volatile //col:49
NMI_BROADCAST_ACTION_TYPE volatile //col:50
IgnoreEvent bool //col:51
IgnoreOneMtf bool //col:52
WaitForStepTrap bool //col:53
MsrState PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE //col:54
SoftwareBreakpointState PDEBUGGEE_BP_DESCRIPTOR //col:55
InstrumentationStepInTrace DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE //col:56
EnableExternalInterruptsOnContinue bool //col:57
EnableExternalInterruptsOnContinueMtf bool //col:58
DisableTrapFlagOnContinue bool //col:59
DoNotNmiNotifyOtherCoresByThisCore bool //col:60
ThreadOrProcessTracingDetails DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS //col:61
BreakStarterCore bool //col:62
InstructionLengthHint uint16 //col:63
HardwareDebugRegisterForStepping uint64 //col:64
* uint64 //col:65
* uint64 //col:66
}


type DEBUGGER_EVENT_ACTION struct{
Tag uint64 //col:70
ActionOrderCode uint32 //col:71
ActionsList *list.List //col:72
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM //col:73
ImmediatelySendTheResults bool //col:74
DebuggerEventActionRunScriptConfiguration DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION //col:75
ScriptConfiguration byte //col:76
DebuggerEventRequestBuffer DEBUGGER_EVENT_REQUEST_BUFFER //col:77
RequestedBuffer byte //col:78
CustomCodeBufferSize uint32 //col:79
CustomCodeBufferAddress PVOID //col:80
}


type DEBUGGER_EVENT struct{
Tag uint64 //col:84
EventsOfSameTypeList *list.List //col:85
EventType DEBUGGER_EVENT_TYPE_ENUM //col:86
Enabled bool //col:87
CoreId uint32 //col:88
Uint32 uint32 //col:89
ProcessId byte //col:90
ActionsListHead *list.List //col:91
CountOfActions uint32 //col:92
OptionalParam1 uint64 //col:93
OptionalParam2 uint64 //col:94
OptionalParam3 uint64 //col:95
OptionalParam4 uint64 //col:96
ConditionsBufferSize uint32 //col:97
ConditionBufferAddress PVOID //col:98
}




