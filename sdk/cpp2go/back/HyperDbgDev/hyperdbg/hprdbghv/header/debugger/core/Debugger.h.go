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
HiddenHookReadAndWriteEventsHead LIST_ENTRY
HiddenHookReadEventsHead LIST_ENTRY
HiddenHookWriteEventsHead LIST_ENTRY
EptHook2sExecDetourEventsHead LIST_ENTRY
EptHookExecCcEventsHead LIST_ENTRY
SyscallHooksEferSyscallEventsHead LIST_ENTRY
SyscallHooksEferSysretEventsHead LIST_ENTRY
CpuidInstructionExecutionEventsHead LIST_ENTRY
RdmsrInstructionExecutionEventsHead LIST_ENTRY
WrmsrInstructionExecutionEventsHead LIST_ENTRY
ExceptionOccurredEventsHead LIST_ENTRY
TscInstructionExecutionEventsHead LIST_ENTRY
PmcInstructionExecutionEventsHead LIST_ENTRY
InInstructionExecutionEventsHead LIST_ENTRY
OutInstructionExecutionEventsHead LIST_ENTRY
DebugRegistersAccessedEventsHead LIST_ENTRY
ExternalInterruptOccurredEventsHead LIST_ENTRY
VmcallInstructionExecutionEventsHead LIST_ENTRY
ControlRegisterModifiedEventsHead LIST_ENTRY
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
ActionOrderCode UINT32
ActionsList LIST_ENTRY
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM
ImmediatelySendTheResults bool
DebuggerEventActionRunScriptConfiguration DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
ScriptConfiguration; byte
DebuggerEventRequestBuffer DEBUGGER_EVENT_REQUEST_BUFFER
RequestedBuffer; byte
CustomCodeBufferSize UINT32
CustomCodeBufferAddress PVOID
}


type DEBUGGER_EVENT struct{
Tag uint64
EventsOfSameTypeList LIST_ENTRY
EventType DEBUGGER_EVENT_TYPE_ENUM
Enabled bool
CoreId UINT32
Uint32 UINT32
ProcessId; byte
ActionsListHead LIST_ENTRY
CountOfActions UINT32
OptionalParam1 uint64
OptionalParam2 uint64
OptionalParam3 uint64
OptionalParam4 uint64
ConditionsBufferSize UINT32
ConditionBufferAddress PVOID
}




