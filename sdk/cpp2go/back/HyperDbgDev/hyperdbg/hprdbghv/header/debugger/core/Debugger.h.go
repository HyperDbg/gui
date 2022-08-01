package core
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\core\Debugger.h.back

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
HiddenHookReadAndWriteEventsHead *list.List
HiddenHookReadEventsHead *list.List
HiddenHookWriteEventsHead *list.List
EptHook2sExecDetourEventsHead *list.List
EptHookExecCcEventsHead *list.List
SyscallHooksEferSyscallEventsHead *list.List
SyscallHooksEferSysretEventsHead *list.List
CpuidInstructionExecutionEventsHead *list.List
RdmsrInstructionExecutionEventsHead *list.List
WrmsrInstructionExecutionEventsHead *list.List
ExceptionOccurredEventsHead *list.List
TscInstructionExecutionEventsHead *list.List
PmcInstructionExecutionEventsHead *list.List
InInstructionExecutionEventsHead *list.List
OutInstructionExecutionEventsHead *list.List
DebugRegistersAccessedEventsHead *list.List
ExternalInterruptOccurredEventsHead *list.List
VmcallInstructionExecutionEventsHead *list.List
ControlRegisterModifiedEventsHead *list.List
}


type PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE struct{
Msr uint64
Value uint64
}


type DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE struct{
WaitForInstrumentationStepInMtf bool
CsSel UINT16
}


type DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS struct{
InitialSetProcessChangeEvent bool
InitialSetThreadChangeEvent bool
InitialSetByClockInterrupt bool
CurrentThreadLocationOnGs uint64
DebugRegisterInterceptionState bool
InterceptClockInterruptsForThreadChange bool
IsWatingForMovCr3VmExits bool
InterceptClockInterruptsForProcessChange bool
}


type PROCESSOR_DEBUGGING_STATE struct{
LONG volatile
BOOLEAN volatile
BOOLEAN volatile
BOOLEAN volatile
NMI_BROADCAST_ACTION_TYPE volatile
IgnoreEvent bool
IgnoreOneMtf bool
WaitForStepTrap bool
MsrState PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE
SoftwareBreakpointState PDEBUGGEE_BP_DESCRIPTOR
InstrumentationStepInTrace DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE
EnableExternalInterruptsOnContinue bool
EnableExternalInterruptsOnContinueMtf bool
DisableTrapFlagOnContinue bool
DoNotNmiNotifyOtherCoresByThisCore bool
ThreadOrProcessTracingDetails DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS
BreakStarterCore bool
InstructionLengthHint UINT16
HardwareDebugRegisterForStepping uint64
* uint64
* uint64
}


type DEBUGGER_EVENT_ACTION struct{
Tag uint64
ActionOrderCode uint32
ActionsList *list.List
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM
ImmediatelySendTheResults bool
DebuggerEventActionRunScriptConfiguration DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
ScriptConfiguration byte
DebuggerEventRequestBuffer DEBUGGER_EVENT_REQUEST_BUFFER
RequestedBuffer byte
CustomCodeBufferSize uint32
CustomCodeBufferAddress PVOID
}


type DEBUGGER_EVENT struct{
Tag uint64
EventsOfSameTypeList *list.List
EventType DEBUGGER_EVENT_TYPE_ENUM
Enabled bool
CoreId uint32
Uint32 uint32
ProcessId byte
ActionsListHead *list.List
CountOfActions uint32
OptionalParam1 uint64
OptionalParam2 uint64
OptionalParam3 uint64
OptionalParam4 uint64
ConditionsBufferSize uint32
ConditionBufferAddress PVOID
}




