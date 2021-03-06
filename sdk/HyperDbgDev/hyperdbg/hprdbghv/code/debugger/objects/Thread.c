#include "pch.h"
BOOLEAN
ThreadHandleThreadChange(UINT32 CurrentCore, PGUEST_REGS GuestState) {
    if ((g_ThreadSwitch.ThreadId != NULL && g_ThreadSwitch.ThreadId == PsGetCurrentThreadId()) ||
        (g_ThreadSwitch.Thread != NULL && g_ThreadSwitch.Thread == PsGetCurrentThread())) {
        KdHandleBreakpointAndDebugBreakpoints(CurrentCore, GuestState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED, NULL);
        return TRUE;
    }
    return FALSE;
}

BOOLEAN
ThreadSwitch(UINT32 ThreadId, PETHREAD EThread, BOOLEAN CheckByClockInterrupt) {
    ULONG CoreCount         = 0;
    g_ThreadSwitch.Thread   = NULL;
    g_ThreadSwitch.ThreadId = NULL;
    if (ThreadId == NULL && EThread == NULL) {
        return FALSE;
    }
    if (EThread != NULL) {
        if (CheckMemoryAccessSafety(EThread, sizeof(BYTE))) {
            g_ThreadSwitch.Thread = EThread;
        } else {
            return FALSE;
        }
    } else if (ThreadId != NULL) {
        g_ThreadSwitch.ThreadId = ThreadId;
    }
    CoreCount = KeQueryActiveProcessorCount(0);
    for (size_t i = 0; i < CoreCount; i++) {
        g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent = TRUE;
        g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetByClockInterrupt  = CheckByClockInterrupt;
    }
    return TRUE;
}

BOOLEAN
ThreadShowList(PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS               ThreadListSymbolInfo,
               DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
               UINT32 *                                           CountOfThreads,
               PVOID                                              ListSaveBuffer,
               UINT64                                             ListSaveBuffSize) {
    UINT64                              ThreadListHead;
    UINT32                              EnumerationCount   = 0;
    UINT64                              Thread             = NULL;
    LIST_ENTRY                          ThreadLinks        = {0};
    CLIENT_ID                           ThreadCid          = {0};
    UINT32                              MaximumBufferCount = 0;
    PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY SavingEntries      = ListSaveBuffer;
    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
        CountOfThreads == NULL) {
        return FALSE;
    }
    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS &&
        (ListSaveBuffer == NULL || ListSaveBuffSize == 0)) {
        return FALSE;
    }
    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS) {
        MaximumBufferCount = ListSaveBuffSize / sizeof(DEBUGGEE_THREAD_LIST_DETAILS_ENTRY);
    }
    UINT32 ThreadListHeadOffset       = ThreadListSymbolInfo->ThreadListHeadOffset;     // nt!_EPROCESS.ThreadListHead
    UINT32 ThreadListEntryOffset      = ThreadListSymbolInfo->ThreadListEntryOffset;    // nt!_ETHREAD.ThreadListEntry
    UINT32 CidOffset                  = ThreadListSymbolInfo->CidOffset;                // nt!_ETHREAD.Cid
    UINT32 ActiveProcessLinksOffset   = ThreadListSymbolInfo->ActiveProcessLinksOffset; // nt!_EPROCESS.ActiveProcessLinks
    UINT64 PsActiveProcessHeadAddress = ThreadListSymbolInfo->PsActiveProcessHead;      // nt!PsActiveProcessHead
    if (ThreadListHeadOffset == NULL ||
        ThreadListEntryOffset == NULL ||
        CidOffset == NULL ||
        ActiveProcessLinksOffset == NULL ||
        PsActiveProcessHeadAddress == NULL) {
        return FALSE;
    }
    if (ThreadListSymbolInfo->Process == NULL) {
        ThreadListSymbolInfo->Process = PsGetCurrentProcess();
        ThreadListHead                = (UINT64)PsGetCurrentProcess() + ThreadListHeadOffset;
    } else {
        ThreadListHead = (UINT64)ThreadListSymbolInfo->Process + ThreadListHeadOffset;
    }
    if (!CheckMemoryAccessSafety(ThreadListHead, sizeof(BYTE))) {
        return FALSE;
    }
    if (!ProcessCheckIfEprocessIsValid(ThreadListSymbolInfo->Process,
                                       PsActiveProcessHeadAddress,
                                       ActiveProcessLinksOffset)) {
        return FALSE;
    }
    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY) {
        Log("PROCESS\t%llx\tIMAGE\t%s\n",
            ThreadListSymbolInfo->Process,
            GetProcessNameFromEprocess(ThreadListSymbolInfo->Process));
    }
    MemoryMapperReadMemorySafe(ThreadListHead, &ThreadLinks, sizeof(ThreadLinks));
    Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
    do {
        MemoryMapperReadMemorySafe(Thread + CidOffset,
                                   &ThreadCid,
                                   sizeof(ThreadCid));
        switch (QueryAction) {
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
            Log("\tTHREAD\t%llx (%llx.%llx)\n", Thread, ThreadCid.UniqueProcess, ThreadCid.UniqueThread);
            break;
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT:
            EnumerationCount++;
            break;
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS:
            EnumerationCount++;
            if (EnumerationCount == MaximumBufferCount - 1) {
                goto ReturnEnd;
            }
            SavingEntries[EnumerationCount - 1].Eprocess = ThreadListSymbolInfo->Process;
            SavingEntries[EnumerationCount - 1].Pid      = ThreadCid.UniqueProcess;
            SavingEntries[EnumerationCount - 1].Tid      = ThreadCid.UniqueThread;
            SavingEntries[EnumerationCount - 1].Ethread  = Thread;
            RtlCopyMemory(&SavingEntries[EnumerationCount - 1].ImageFileName,
                          GetProcessNameFromEprocess(ThreadListSymbolInfo->Process),
                          15);
            break;
        default:
            break;
        }
        MemoryMapperReadMemorySafe(Thread + ThreadListEntryOffset,
                                   &ThreadLinks,
                                   sizeof(ThreadLinks));
        Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
    } while ((UINT64)ThreadLinks.Flink != ThreadListHead);
ReturnEnd:
    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT) {
        *CountOfThreads = EnumerationCount;
    }
    return TRUE;
}

BOOLEAN
ThreadInterpretThread(PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET TidRequest) {
    switch (TidRequest->ActionType) {
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS:
        TidRequest->ProcessId = PsGetCurrentProcessId();
        TidRequest->ThreadId  = PsGetCurrentThreadId();
        TidRequest->Process   = PsGetCurrentProcess();
        TidRequest->Thread    = PsGetCurrentThread();
        MemoryMapperReadMemorySafe(GetProcessNameFromEprocess(PsGetCurrentProcess()), &TidRequest->ProcessName, 16);
        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        break;
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH:
        if (!ThreadSwitch(TidRequest->ThreadId, TidRequest->Thread, TidRequest->CheckByClockInterrupt)) {
            TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
            break;
        }
        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        break;
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST:
        if (!ThreadShowList(&TidRequest->ThreadListSymDetails,
                            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
                            NULL,
                            NULL,
                            NULL)) {
            TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
            break;
        }
        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        break;
    default:
        TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
        break;
    }
    if (TidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        return TRUE;
    } else {
        return FALSE;
    }
}

VOID
ThreadDetectChangeByDebugRegisterOnGs(UINT32  CurrentProcessorIndex,
                                      BOOLEAN Enable) {
    UINT64 MsrGsBase;
    if (Enable) {
        MsrGsBase = __readmsr(IA32_GS_BASE);
        MsrGsBase += 0x188;
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs      = MsrGsBase;
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.DebugRegisterInterceptionState = TRUE;
        HvSetLoadDebugControls(TRUE);
        HvSetSaveDebugControls(TRUE);
        HvSetExceptionBitmap(EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
        DebugRegistersSet(
            DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT,
            BREAK_ON_WRITE_ONLY,
            TRUE,
            g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs);
        HvSetMovDebugRegsExiting(TRUE);
    } else {
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.DebugRegisterInterceptionState = FALSE;
        HvSetMovDebugRegsExiting(FALSE);
        HvSetLoadDebugControls(FALSE);
        HvSetSaveDebugControls(FALSE);
        HvUnsetExceptionBitmap(EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs = NULL;
    }
}

VOID
ThreadDetectChangeByInterceptingClockInterrupts(UINT32  CurrentProcessorIndex,
                                                BOOLEAN Enable) {
    if (Enable) {
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange = TRUE;
        HvSetExternalInterruptExiting(TRUE);
    } else {
        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange = FALSE;
        HvSetExternalInterruptExiting(FALSE);
    }
}

VOID
ThreadEnableOrDisableThreadChangeMonitor(UINT32  CurrentProcessorIndex,
                                         BOOLEAN Enable,
                                         BOOLEAN CheckByClockInterrupts) {
    if (!CheckByClockInterrupts) {
        ThreadDetectChangeByDebugRegisterOnGs(CurrentProcessorIndex, Enable);
    } else {
        ThreadDetectChangeByInterceptingClockInterrupts(CurrentProcessorIndex, Enable);
    }
}

BOOLEAN
ThreadQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS DebuggerUsermodeProcessOrThreadQueryRequest) {
    BOOLEAN Result = FALSE;
    Result         = ThreadShowList(&DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
                            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT,
                            &DebuggerUsermodeProcessOrThreadQueryRequest->Count,
                            NULL,
                            NULL);
    if (Result && DebuggerUsermodeProcessOrThreadQueryRequest->Count != 0) {
        DebuggerUsermodeProcessOrThreadQueryRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        return TRUE;
    }
    DebuggerUsermodeProcessOrThreadQueryRequest->Result = DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS;
    return FALSE;
}

BOOLEAN
ThreadQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS DebuggerUsermodeProcessOrThreadQueryRequest,
                PVOID                                       AddressToSaveDetail,
                UINT32                                      BufferSize) {
    BOOLEAN Result = FALSE;
    Result         = ThreadShowList(&DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
                            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS,
                            NULL,
                            AddressToSaveDetail,
                            BufferSize);
    return Result;
}

BOOLEAN
ThreadQueryDetails(PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThreadRequest) {
    GetInformationThreadRequest->ProcessId = PsGetCurrentProcessId();
    GetInformationThreadRequest->Process   = PsGetCurrentProcess();
    GetInformationThreadRequest->Thread    = PsGetCurrentThread();
    GetInformationThreadRequest->ThreadId  = PsGetCurrentThreadId();
    RtlCopyMemory(&GetInformationThreadRequest->ProcessName,
                  GetProcessNameFromEprocess(PsGetCurrentProcess()),
                  15);
    GetInformationThreadRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}
