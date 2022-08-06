#include "pch.h"
VOID
TerminateExternalInterruptEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->ExternalInterruptOccurredEventsHead) > 1)
    {
        ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
        TempList = &g_Events->ExternalInterruptOccurredEventsHead;
        while (&g_Events->ExternalInterruptOccurredEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandSetExternalInterruptExitingAllCores();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformSetExternalInterruptExitingOnSingleCore, NULL);
                }
            }
        }
    }
    else
    {
        ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
    }
}

VOID
TerminateHiddenHookReadAndWriteEvent(PDEBUGGER_EVENT Event)
{
    UINT64 PagesBytes;
    UINT64 TempOptionalParam1;
    TempOptionalParam1 = Event->OptionalParam3;
    PagesBytes = PAGE_ALIGN(TempOptionalParam1);
    PagesBytes = Event->OptionalParam4 - PagesBytes;
    for (size_t i = 0; i <= PagesBytes / PAGE_SIZE; i++)
    {
        EptHookUnHookSingleAddress((UINT64)TempOptionalParam1 + (i * PAGE_SIZE), NULL, Event->ProcessId);
    }
}

VOID
TerminateHiddenHookExecCcEvent(PDEBUGGER_EVENT Event)
{
    EptHookUnHookSingleAddress(Event->OptionalParam1, NULL, Event->ProcessId);
}

VOID
TerminateHiddenHookExecDetoursEvent(PDEBUGGER_EVENT Event)
{
    EptHookUnHookSingleAddress(NULL, Event->OptionalParam1, Event->ProcessId);
}

VOID
TerminateRdmsrExecutionEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->RdmsrInstructionExecutionEventsHead) > 1)
    {
        ExtensionCommandResetChangeAllMsrBitmapReadAllCores();
        TempList = &g_Events->RdmsrInstructionExecutionEventsHead;
        while (&g_Events->RdmsrInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandChangeAllMsrBitmapReadAllCores(CurrentEvent->OptionalParam1);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeMsrBitmapReadOnSingleCore, CurrentEvent->OptionalParam1);
                }
            }
        }
    }
    else
    {
        ExtensionCommandResetChangeAllMsrBitmapReadAllCores();
    }
}

VOID
TerminateWrmsrExecutionEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->WrmsrInstructionExecutionEventsHead) > 1)
    {
        ExtensionCommandResetAllMsrBitmapWriteAllCores();
        TempList = &g_Events->WrmsrInstructionExecutionEventsHead;
        while (&g_Events->WrmsrInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandChangeAllMsrBitmapWriteAllCores(CurrentEvent->OptionalParam1);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore, CurrentEvent->OptionalParam1);
                }
            }
        }
    }
    else
    {
        ExtensionCommandResetAllMsrBitmapWriteAllCores();
    }
}

VOID
TerminateExceptionEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->ExceptionOccurredEventsHead) > 1)
    {
        ExtensionCommandResetExceptionBitmapAllCores();
        TempList = &g_Events->ExceptionOccurredEventsHead;
        while (&g_Events->ExceptionOccurredEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandSetExceptionBitmapAllCores(CurrentEvent->OptionalParam1);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformSetExceptionBitmapOnSingleCore, CurrentEvent->OptionalParam1);
                }
            }
        }
    }
    else
    {
        ExtensionCommandResetExceptionBitmapAllCores();
    }
}

VOID
TerminateInInstructionExecutionEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) > 1 ||
        DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) >= 1)
    {
        ExtensionCommandIoBitmapResetAllCores();
        TempList = &g_Events->InInstructionExecutionEventsHead;
        while (&g_Events->InInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandIoBitmapChangeAllCores(CurrentEvent->OptionalParam1);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeIoBitmapOnSingleCore, CurrentEvent->OptionalParam1);
                }
            }
        }
    }
    else
    {
        ExtensionCommandIoBitmapResetAllCores();
    }
}

VOID
TerminateOutInstructionExecutionEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) > 1 ||
        DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) >= 1)
    {
        ExtensionCommandIoBitmapResetAllCores();
        TempList = &g_Events->OutInstructionExecutionEventsHead;
        while (&g_Events->OutInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandIoBitmapChangeAllCores(CurrentEvent->OptionalParam1);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeIoBitmapOnSingleCore, CurrentEvent->OptionalParam1);
                }
            }
        }
    }
    else
    {
        ExtensionCommandIoBitmapResetAllCores();
    }
}

VOID
TerminateVmcallExecutionEvent(PDEBUGGER_EVENT Event)
{
    if (DebuggerEventListCount(&g_Events->VmcallInstructionExecutionEventsHead) > 1)
    {
        return;
    }
    else
    {
        g_TriggerEventForVmcalls = FALSE;
    }
}

VOID
TerminateCpuidExecutionEvent(PDEBUGGER_EVENT Event)
{
    if (DebuggerEventListCount(&g_Events->CpuidInstructionExecutionEventsHead) > 1)
    {
        return;
    }
    else
    {
        g_TriggerEventForCpuids = FALSE;
    }
}

VOID
TerminateTscEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->TscInstructionExecutionEventsHead) > 1)
    {
        ExtensionCommandDisableRdtscExitingForClearingEventsAllCores();
        TempList = &g_Events->TscInstructionExecutionEventsHead;
        while (&g_Events->TscInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandEnableRdtscExitingAllCores();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableRdtscExitingOnSingleCore, NULL);
                }
            }
        }
    }
    else
    {
        ExtensionCommandDisableRdtscExitingForClearingEventsAllCores();
    }
}

VOID
TerminatePmcEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->PmcInstructionExecutionEventsHead) > 1)
    {
        ExtensionCommandDisableRdpmcExitingAllCores();
        TempList = &g_Events->PmcInstructionExecutionEventsHead;
        while (&g_Events->PmcInstructionExecutionEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandEnableRdpmcExitingAllCores();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableRdpmcExitingOnSingleCore, NULL);
                }
            }
        }
    }
    else
    {
        ExtensionCommandDisableRdpmcExitingAllCores();
    }
}

VOID
TerminateControlRegistersEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->ControlRegisterModifiedEventsHead) > 1)
    {
        ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(Event);
        TempList = &g_Events->ControlRegisterModifiedEventsHead;
        while (&g_Events->ControlRegisterModifiedEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandEnableMovControlRegisterExitingAllCores(CurrentEvent);
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableMovToControlRegisterExiting, CurrentEvent);
                }
            }
        }
    }
    else
    {
        ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(Event);
    }
}

VOID
TerminateDebugRegistersEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->DebugRegistersAccessedEventsHead) > 1)
    {
        ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores();
        TempList = &g_Events->DebugRegistersAccessedEventsHead;
        while (&g_Events->DebugRegistersAccessedEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    ExtensionCommandEnableMovDebugRegistersExitingAllCores();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableMovToDebugRegistersExiting, NULL);
                }
            }
        }
    }
    else
    {
        ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores();
    }
}

VOID
TerminateSyscallHookEferEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->SyscallHooksEferSyscallEventsHead) > 1 ||
        DebuggerEventListCount(&g_Events->SyscallHooksEferSysretEventsHead) >= 1)
    {
        DebuggerEventDisableEferOnAllProcessors();
        TempList = &g_Events->SyscallHooksEferSyscallEventsHead;
        while (&g_Events->SyscallHooksEferSyscallEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    DebuggerEventEnableEferOnAllProcessors();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableEferSyscallHookOnSingleCore, NULL);
                }
            }
        }
    }
    else
    {
        DebuggerEventDisableEferOnAllProcessors();
    }
}

VOID
TerminateSysretHookEferEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList = 0;
    if (DebuggerEventListCount(&g_Events->SyscallHooksEferSysretEventsHead) > 1 ||
        DebuggerEventListCount(&g_Events->SyscallHooksEferSyscallEventsHead) > 1)
    {
        DebuggerEventDisableEferOnAllProcessors();
        TempList = &g_Events->SyscallHooksEferSysretEventsHead;
        while (&g_Events->SyscallHooksEferSysretEventsHead != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag != Event->Tag)
            {
                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
                {
                    DebuggerEventEnableEferOnAllProcessors();
                }
                else
                {
                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformEnableEferSyscallHookOnSingleCore, NULL);
                }
            }
        }
    }
    else
    {
        DebuggerEventDisableEferOnAllProcessors();
    }
}

