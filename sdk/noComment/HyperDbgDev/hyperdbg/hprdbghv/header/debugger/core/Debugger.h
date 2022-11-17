﻿
#pragma once
BOOLEAN g_TriggerEventForVmcalls;
BOOLEAN g_TriggerEventForCpuids;
#define DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER 0
#define DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT 1
#define DEBUGGER_DEBUG_REGISTER_FOR_USER_MODE_ENTRY_POINT 1
typedef struct _DEBUGGER_CORE_EVENTS {
  LIST_ENTRY
  HiddenHookReadAndWriteEventsHead;    
  LIST_ENTRY HiddenHookReadEventsHead; 
  LIST_ENTRY
  HiddenHookWriteEventsHead; 
  LIST_ENTRY
  EptHook2sExecDetourEventsHead; 
  LIST_ENTRY EptHookExecCcEventsHead; 
  LIST_ENTRY
  SyscallHooksEferSyscallEventsHead; 
  LIST_ENTRY
  SyscallHooksEferSysretEventsHead; 
  LIST_ENTRY
  CpuidInstructionExecutionEventsHead; 
  LIST_ENTRY
  RdmsrInstructionExecutionEventsHead; 
  LIST_ENTRY
  WrmsrInstructionExecutionEventsHead; 
  LIST_ENTRY
  ExceptionOccurredEventsHead; 
  LIST_ENTRY
  TscInstructionExecutionEventsHead; 
  LIST_ENTRY
  PmcInstructionExecutionEventsHead; 
  LIST_ENTRY
  InInstructionExecutionEventsHead; 
  LIST_ENTRY
  OutInstructionExecutionEventsHead; 
  LIST_ENTRY
  DebugRegistersAccessedEventsHead; 
  LIST_ENTRY
  ExternalInterruptOccurredEventsHead; 
  LIST_ENTRY
  VmcallInstructionExecutionEventsHead; 
  LIST_ENTRY
  ControlRegisterModifiedEventsHead; 
} DEBUGGER_CORE_EVENTS, *PDEBUGGER_CORE_EVENTS;
typedef struct _PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE {
  UINT64 Msr;   
  UINT64 Value; 
} PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE,
    *PPROCESSOR_DEBUGGING_MSR_READ_OR_WRITE;
typedef struct _DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE {
  BOOLEAN WaitForInstrumentationStepInMtf;
  UINT16 CsSel; 
} DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE,
    *PDEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE;
typedef struct _DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS {
  BOOLEAN InitialSetProcessChangeEvent;
  BOOLEAN InitialSetThreadChangeEvent;
  BOOLEAN InitialSetByClockInterrupt;
  UINT64 CurrentThreadLocationOnGs;
  BOOLEAN DebugRegisterInterceptionState;
  BOOLEAN InterceptClockInterruptsForThreadChange;
  BOOLEAN IsWatingForMovCr3VmExits;
  BOOLEAN InterceptClockInterruptsForProcessChange;
} DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS,
    *PDEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS;
typedef struct _PROCESSOR_DEBUGGING_STATE {
  volatile LONG Lock;
  volatile BOOLEAN WaitingToBeLocked;
  volatile BOOLEAN MainDebuggingCore;
  volatile BOOLEAN NmiCalledInVmxRootRelatedToHaltDebuggee;
  volatile NMI_BROADCAST_ACTION_TYPE NmiBroadcastAction;
  BOOLEAN IgnoreEvent;
  BOOLEAN IgnoreOneMtf;
  BOOLEAN WaitForStepTrap;
  PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE MsrState;
  PDEBUGGEE_BP_DESCRIPTOR SoftwareBreakpointState;
  DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE InstrumentationStepInTrace;
  BOOLEAN EnableExternalInterruptsOnContinue;
  BOOLEAN EnableExternalInterruptsOnContinueMtf;
  BOOLEAN DisableTrapFlagOnContinue;
  BOOLEAN DoNotNmiNotifyOtherCoresByThisCore;
  DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS ThreadOrProcessTracingDetails;
  BOOLEAN BreakStarterCore;
  UINT16 InstructionLengthHint;
  UINT64 HardwareDebugRegisterForStepping;
  UINT64 *ScriptEngineCoreSpecificLocalVariable;
  UINT64 *ScriptEngineCoreSpecificTempVariable;
} PROCESSOR_DEBUGGING_STATE, PPROCESSOR_DEBUGGING_STATE;
typedef struct _DEBUGGER_EVENT_ACTION {
  UINT64 Tag;             
  UINT32 ActionOrderCode; 
  LIST_ENTRY ActionsList; 
  DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType; 
  BOOLEAN ImmediatelySendTheResults; 
  DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
  ScriptConfiguration; 
  DEBUGGER_EVENT_REQUEST_BUFFER
  RequestedBuffer; 
  UINT32 CustomCodeBufferSize;   
  PVOID CustomCodeBufferAddress; 
} DEBUGGER_EVENT_ACTION, *PDEBUGGER_EVENT_ACTION;
typedef struct _DEBUGGER_EVENT {
  UINT64 Tag;
  LIST_ENTRY EventsOfSameTypeList; 
  DEBUGGER_EVENT_TYPE_ENUM EventType;
  BOOLEAN Enabled;
  UINT32 CoreId; 
  UINT32
  ProcessId; 
  LIST_ENTRY ActionsListHead; 
  UINT32 CountOfActions;      
  UINT64 OptionalParam1; 
  UINT64 OptionalParam2; 
  UINT64 OptionalParam3; 
  UINT64 OptionalParam4; 
  UINT32 ConditionsBufferSize;  
  PVOID ConditionBufferAddress; 
} DEBUGGER_EVENT, *PDEBUGGER_EVENT;
typedef enum _DEBUGGER_TRIGGERING_EVENT_STATUS_TYPE {
  DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL = 0,
  DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT = 1,
  DEBUGGER_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED = 2,
  DEBUGGER_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE = 3,
} DEBUGGER_TRIGGERING_EVENT_STATUS_TYPE;
typedef UINT64 DebuggerCheckForCondition(PGUEST_REGS Regs, PVOID Context);
typedef PVOID DebuggerRunCustomCodeFunc(PVOID PreAllocatedBufferAddress,
                                        PGUEST_REGS Regs, PVOID Context);
NTSTATUS
MemoryManagerReadProcessMemoryNormal(HANDLE PID, PVOID Address,
                                     DEBUGGER_READ_MEMORY_TYPE MemType,
                                     PVOID UserBuffer, SIZE_T Size,
                                     PSIZE_T ReturnSize);
UINT64
DebuggerGetRegValueWrapper(PGUEST_REGS GuestRegs, UINT32  RegId);
UINT32
DebuggerGetLastError();
VOID DebuggerSetLastError(UINT32 LastError);
BOOLEAN
DebuggerInitialize();
VOID DebuggerUninitialize();
PDEBUGGER_EVENT
DebuggerCreateEvent(BOOLEAN Enabled, UINT32 CoreId, UINT32 ProcessId,
                    DEBUGGER_EVENT_TYPE_ENUM EventType, UINT64 Tag,
                    UINT64 OptionalParam1, UINT64 OptionalParam2,
                    UINT64 OptionalParam3, UINT64 OptionalParam4,
                    UINT32 ConditionsBufferSize, PVOID ConditionBuffer);
PDEBUGGER_EVENT_ACTION
DebuggerAddActionToEvent(
    PDEBUGGER_EVENT Event, DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType,
    BOOLEAN SendTheResultsImmediately,
    PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE InTheCaseOfCustomCode,
    PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION InTheCaseOfRunScript);
BOOLEAN
DebuggerRegisterEvent(PDEBUGGER_EVENT Event);
DEBUGGER_TRIGGERING_EVENT_STATUS_TYPE
DebuggerTriggerEvents(DEBUGGER_EVENT_TYPE_ENUM EventType, PGUEST_REGS Regs,
                      PVOID Context);
PDEBUGGER_EVENT
DebuggerGetEventByTag(UINT64 Tag);
BOOLEAN
DebuggerRemoveEvent(UINT64 Tag);
BOOLEAN
DebuggerParseEventFromUsermode(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails, UINT32 BufferLength,
    PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER ResultsToReturnUsermode);
BOOLEAN
DebuggerParseActionFromUsermode(
    PDEBUGGER_GENERAL_ACTION Action, UINT32 BufferLength,
    PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER ResultsToReturnUsermode);
BOOLEAN
DebuggerParseEventsModificationFromUsermode(
    PDEBUGGER_MODIFY_EVENTS DebuggerEventModificationRequest);
BOOLEAN
DebuggerTerminateEvent(UINT64 Tag);
BOOLEAN
DebuggerEnableOrDisableAllEvents(BOOLEAN IsEnable);
BOOLEAN
DebuggerRemoveAllEvents();
BOOLEAN
DebuggerTerminateAllEvents();
UINT32
DebuggerEventListCount(PLIST_ENTRY TargetEventList);
UINT32
DebuggerEventListCountByCore(PLIST_ENTRY TargetEventList, UINT32 TargetCore);
UINT32
DebuggerExceptionEventBitmapMask(UINT32 CoreIndex);
BOOLEAN
DebuggerIsTagValid(UINT64 Tag);
BOOLEAN
DebuggerEnableEvent(UINT64 Tag);
BOOLEAN
DebuggerQueryStateEvent(UINT64 Tag);
BOOLEAN
DebuggerDisableEvent(UINT64 Tag);
VOID DebuggerPerformActions(PDEBUGGER_EVENT Event, PGUEST_REGS Regs,
                            PVOID Context);
VOID DebuggerPerformBreakToDebugger(UINT64 Tag, PDEBUGGER_EVENT_ACTION Action,
                                    PGUEST_REGS Regs, PVOID Context);
BOOLEAN
DebuggerPerformRunScript(UINT64 Tag, PDEBUGGER_EVENT_ACTION Action,
                         PDEBUGGEE_SCRIPT_PACKET ScriptDetails,
                         PGUEST_REGS Regs, PVOID Context);
VOID DebuggerPerformRunTheCustomCode(UINT64 Tag, PDEBUGGER_EVENT_ACTION Action,
                                     PGUEST_REGS Regs, PVOID Context);