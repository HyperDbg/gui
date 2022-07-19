#include "pch.h"
BOOLEAN
AttachingInitialize() {
    UNICODE_STRING FunctionName;
    if (g_PsGetProcessPeb == NULL) {
        RtlInitUnicodeString(&FunctionName, L"PsGetProcessPeb");
        g_PsGetProcessPeb = (PsGetProcessPeb)MmGetSystemRoutineAddress(&FunctionName);
        if (g_PsGetProcessPeb == NULL) {
            LogError("Err, cannot resolve PsGetProcessPeb");
        }
    }
    if (g_PsGetProcessWow64Process == NULL) {
        RtlInitUnicodeString(&FunctionName, L"PsGetProcessWow64Process");
        g_PsGetProcessWow64Process = (PsGetProcessWow64Process)MmGetSystemRoutineAddress(&FunctionName);
        if (g_PsGetProcessWow64Process == NULL) {
            LogError("Err, cannot resolve PsGetProcessPeb");
        }
    }
    if (g_ZwQueryInformationProcess == NULL) {
        UNICODE_STRING RoutineName;
        RtlInitUnicodeString(&RoutineName, L"ZwQueryInformationProcess");
        g_ZwQueryInformationProcess = (ZwQueryInformationProcess)MmGetSystemRoutineAddress(&RoutineName);
        if (g_ZwQueryInformationProcess == NULL) {
            LogError("Err, cannot resolve ZwQueryInformationProcess");
        }
    }
    return TRUE;
}

UINT64
AttachingCreateProcessDebuggingDetails(UINT32    ProcessId,
                                       BOOLEAN   Enabled,
                                       BOOLEAN   Is32Bit,
                                       PEPROCESS Eprocess,
                                       UINT64    PebAddressToMonitor,
                                       UINT64    UsermodeReservedBuffer) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = (PUSERMODE_DEBUGGING_PROCESS_DETAILS)
        ExAllocatePoolWithTag(NonPagedPool, sizeof(USERMODE_DEBUGGING_PROCESS_DETAILS), POOLTAG);
    if (!ProcessDebuggingDetail) {
        return NULL;
    }
    RtlZeroMemory(ProcessDebuggingDetail, sizeof(USERMODE_DEBUGGING_PROCESS_DETAILS));
    ProcessDebuggingDetail->Token                  = g_SeedOfUserDebuggingDetails++;
    ProcessDebuggingDetail->ProcessId              = ProcessId;
    ProcessDebuggingDetail->Enabled                = Enabled;
    ProcessDebuggingDetail->Is32Bit                = Is32Bit;
    ProcessDebuggingDetail->Eprocess               = Eprocess;
    ProcessDebuggingDetail->PebAddressToMonitor    = PebAddressToMonitor;
    ProcessDebuggingDetail->UsermodeReservedBuffer = UsermodeReservedBuffer;
    if (!ThreadHolderAssignThreadHolderToProcessDebuggingDetails(ProcessDebuggingDetail)) {
        ExFreePoolWithTag(ProcessDebuggingDetail, POOLTAG);
        return NULL;
    }
    InsertHeadList(&g_ProcessDebuggingDetailsListHead, &(ProcessDebuggingDetail->AttachedProcessList));
    return ProcessDebuggingDetail->Token;
}

PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsByToken(UINT64 Token) {
    LIST_FOR_EACH_LINK(g_ProcessDebuggingDetailsListHead, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList, ProcessDebuggingDetails) {
        if (ProcessDebuggingDetails->Token == Token && ProcessDebuggingDetails->Enabled) {
            return ProcessDebuggingDetails;
        }
    }
    return NULL;
}

PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsByProcessId(UINT32 ProcessId) {
    LIST_FOR_EACH_LINK(g_ProcessDebuggingDetailsListHead, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList, ProcessDebuggingDetails) {
        if (ProcessDebuggingDetails->ProcessId == ProcessId && ProcessDebuggingDetails->Enabled) {
            return ProcessDebuggingDetails;
        }
    }
    return NULL;
}

PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsInStartingPhase() {
    LIST_FOR_EACH_LINK(g_ProcessDebuggingDetailsListHead, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList, ProcessDebuggingDetails) {
        if (ProcessDebuggingDetails->IsOnTheStartingPhase) {
            return ProcessDebuggingDetails;
        }
    }
    return NULL;
}

VOID
AttachingRemoveAndFreeAllProcessDebuggingDetails() {
    LIST_FOR_EACH_LINK(g_ProcessDebuggingDetailsListHead, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList, ProcessDebuggingDetails) {
        ThreadHolderFreeHoldingStructures(ProcessDebuggingDetails);
        RemoveEntryList(&ProcessDebuggingDetails->AttachedProcessList);
        ExFreePoolWithTag(ProcessDebuggingDetails, POOLTAG);
    }
}

BOOLEAN
AttachingRemoveProcessDebuggingDetailsByToken(UINT64 Token) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails;
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByToken(Token);
    if (!ProcessDebuggingDetails) {
        return FALSE;
    }
    ThreadHolderFreeHoldingStructures(ProcessDebuggingDetails);
    RemoveEntryList(&ProcessDebuggingDetails->AttachedProcessList);
    ExFreePoolWithTag(ProcessDebuggingDetails, POOLTAG);
    return TRUE;
}

BOOLEAN
AttachingSetStartingPhaseOfProcessDebuggingDetailsByToken(BOOLEAN Set, UINT64 Token) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails;
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsInStartingPhase();
    if (Set && ProcessDebuggingDetails != NULL) {
        return FALSE;
    }
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByToken(Token);
    if (!ProcessDebuggingDetails) {
        return FALSE;
    }
    ProcessDebuggingDetails->IsOnTheStartingPhase = Set;
    return TRUE;
}

VOID
AttachingReachedToProcessEntrypoint(UINT32 CurrentProcessorIndex, PGUEST_REGS GuestRegs, UINT64 ThreadDebuggingToken) {
    AttachingSetStartingPhaseOfProcessDebuggingDetailsByToken(FALSE, ThreadDebuggingToken);
    if (g_KernelDebuggerState) {
        KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
                                              GuestRegs,
                                              DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED,
                                              NULL);
    } else {
        UdCheckAndHandleBreakpointsAndDebugBreaks(CurrentProcessorIndex,
                                                  GuestRegs,
                                                  DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED,
                                                  NULL);
    }
}

VOID
AttachingHandleEntrypointDebugBreak(UINT32 CurrentProcessorIndex, PGUEST_REGS GuestRegs) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail = NULL;
    g_GuestState[CurrentProcessorIndex].IncrementRip           = FALSE;
    ProcessDebuggingDetail                                     = AttachingFindProcessDebuggingDetailsByProcessId(PsGetCurrentProcessId());
    if (ProcessDebuggingDetail != NULL) {
        if (g_IsWaitingForUserModeModuleEntrypointToBeCalled) {
            g_IsWaitingForUserModeModuleEntrypointToBeCalled = FALSE;
            if (!CheckMemoryAccessSafety(ProcessDebuggingDetail->EntrypointOfMainModule, sizeof(CHAR))) {
                VMEXIT_INTERRUPT_INFORMATION InterruptInfo = {0};
                g_IsWaitingForReturnAndRunFromPageFault    = TRUE;
                InterruptInfo.Vector                       = EXCEPTION_VECTOR_PAGE_FAULT;
                InterruptInfo.InterruptionType             = INTERRUPT_TYPE_HARDWARE_EXCEPTION;
                InterruptInfo.ErrorCodeValid               = TRUE;
                InterruptInfo.NmiUnblocking                = FALSE;
                InterruptInfo.Valid                        = TRUE;
                IdtEmulationHandlePageFaults(CurrentProcessorIndex, InterruptInfo, ProcessDebuggingDetail->EntrypointOfMainModule, 0x14);
                DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_USER_MODE_ENTRY_POINT,
                                  BREAK_ON_INSTRUCTION_FETCH,
                                  FALSE,
                                  ProcessDebuggingDetail->EntrypointOfMainModule);
            } else {
                AttachingReachedToProcessEntrypoint(CurrentProcessorIndex, GuestRegs, ProcessDebuggingDetail->Token);
            }
        } else if (g_IsWaitingForReturnAndRunFromPageFault) {
            g_IsWaitingForReturnAndRunFromPageFault = FALSE;
            AttachingReachedToProcessEntrypoint(CurrentProcessorIndex, GuestRegs, ProcessDebuggingDetail->Token);
        }
    } else {
        ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsInStartingPhase();
        if (ProcessDebuggingDetail != NULL) {
            DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_USER_MODE_ENTRY_POINT,
                              BREAK_ON_INSTRUCTION_FETCH,
                              FALSE,
                              ProcessDebuggingDetail->EntrypointOfMainModule);
        }
    }
}

BOOLEAN
AttachingAdjustNopSledBuffer(UINT64 ReservedBuffAddress, UINT32 ProcessId) {
    PEPROCESS  SourceProcess;
    KAPC_STATE State = {0};
    if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS) {
        return FALSE;
    }
    __try {
        KeStackAttachProcess(SourceProcess, &State);
        memset(ReservedBuffAddress, 0x90, PAGE_SIZE);
        *(UINT16 *)(ReservedBuffAddress + PAGE_SIZE - 4) = 0xa20f;
        *(UINT16 *)(ReservedBuffAddress + PAGE_SIZE - 2) = 0xf4eb;
        KeUnstackDetachProcess(&State);
        ObDereferenceObject(SourceProcess);
    } __except (EXCEPTION_EXECUTE_HANDLER) {
        KeUnstackDetachProcess(&State);
        return NULL;
    }
    return TRUE;
}

BOOLEAN
AttachingCheckPageFaultsWithUserDebugger(UINT32                       CurrentProcessorIndex,
                                         PGUEST_REGS                  GuestRegs,
                                         VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                         UINT64                       Address,
                                         ULONG                        ErrorCode) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    VIRTUAL_MACHINE_STATE *             CurrentVmState = &g_GuestState[CurrentProcessorIndex];
    if (CurrentVmState->LastVmexitRip & 0xf000000000000000) {
        return FALSE;
    }
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(PsGetCurrentProcessId());
    if (!ProcessDebuggingDetail) {
        return FALSE;
    }
    if (!ProcessDebuggingDetail->IsOnThreadInterceptingPhase) {
        return FALSE;
    }
    UdCheckAndHandleBreakpointsAndDebugBreaks(CurrentProcessorIndex,
                                              GuestRegs,
                                              DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED,
                                              NULL);
    CurrentVmState->IncrementRip = FALSE;
    return TRUE;
}

BOOLEAN
AttachingConfigureInterceptingThreads(UINT64 ProcessDebuggingToken, BOOLEAN Enable) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByToken(ProcessDebuggingToken);
    if (!ProcessDebuggingDetail) {
        return FALSE;
    }
    if (Enable && ProcessDebuggingDetail->IsOnThreadInterceptingPhase) {
        LogError("Err, thread is already in intercepting phase");
        return FALSE;
    }
    if (!Enable && !ProcessDebuggingDetail->IsOnThreadInterceptingPhase) {
        return FALSE;
    }
    g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger  = Enable;
    ProcessDebuggingDetail->IsOnThreadInterceptingPhase = Enable;
    if (Enable) {
        DebuggerEventEnableMovToCr3ExitingOnAllProcessors();
    } else {
        DebuggerEventDisableMovToCr3ExitingOnAllProcessors();
    }
    if (!Enable) {
        for (size_t i = 0; i < MAX_CR3_IN_A_PROCESS; i++) {
            if (ProcessDebuggingDetail->InterceptedCr3[i].Flags != NULL) {
                if (!MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(NULL,
                                                                       TRUE,
                                                                       PagingLevelPageMapLevel4,
                                                                       ProcessDebuggingDetail->InterceptedCr3[i])) {
                    return FALSE;
                }
            }
        }
    }
    return TRUE;
}

BOOLEAN
AttachingPerformAttachToProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS AttachRequest, BOOLEAN IsAttachingToEntrypoint) {
    PEPROCESS                           SourceProcess;
    UINT64                              ProcessDebuggingToken;
    UINT64                              PebAddressToMonitor;
    UINT64                              UsermodeReservedBuffer;
    BOOLEAN                             ResultOfApplyingEvent;
    BOOLEAN                             Is32Bit;
    PUSERMODE_DEBUGGING_PROCESS_DETAILS TempProcessDebuggingDetail;
    if (g_PsGetProcessWow64Process == NULL || g_PsGetProcessPeb == NULL) {
        AttachRequest->Result = DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED;
        return FALSE;
    }
    if (PsLookupProcessByProcessId(AttachRequest->ProcessId, &SourceProcess) != STATUS_SUCCESS) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
        return FALSE;
    }
    ObDereferenceObject(SourceProcess);
    TempProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(AttachRequest->ProcessId);
    if (TempProcessDebuggingDetail != NULL &&
        TempProcessDebuggingDetail->Eprocess == SourceProcess) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS;
        return FALSE;
    }
    if (!UserAccessIsWow64Process(AttachRequest->ProcessId, &Is32Bit)) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS;
        return FALSE;
    }
    AttachRequest->Is32Bit = Is32Bit;
    if (Is32Bit) {
        PebAddressToMonitor = (PPEB32)g_PsGetProcessWow64Process(SourceProcess);
    } else {
        PebAddressToMonitor = (PPEB)g_PsGetProcessPeb(SourceProcess);
    }
    if (PebAddressToMonitor == NULL) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
        return FALSE;
    }
    UsermodeReservedBuffer = MemoryMapperReserveUsermodeAddressInTargetProcess(AttachRequest->ProcessId, TRUE);
    if (UsermodeReservedBuffer == NULL) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
        return FALSE;
    }
    if (!AttachingAdjustNopSledBuffer(UsermodeReservedBuffer,
                                      AttachRequest->ProcessId)) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
        return FALSE;
    }
    ProcessDebuggingToken = AttachingCreateProcessDebuggingDetails(AttachRequest->ProcessId,
                                                                   TRUE,
                                                                   Is32Bit,
                                                                   SourceProcess,
                                                                   PebAddressToMonitor,
                                                                   UsermodeReservedBuffer);
    if (ProcessDebuggingToken == NULL) {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
        return FALSE;
    }
    if (IsAttachingToEntrypoint) {
        g_IsWaitingForUserModeModuleEntrypointToBeCalled = TRUE;
        if (!AttachingSetStartingPhaseOfProcessDebuggingDetailsByToken(TRUE, ProcessDebuggingToken)) {
            AttachingRemoveProcessDebuggingDetailsByToken(ProcessDebuggingToken);
            g_IsWaitingForUserModeModuleEntrypointToBeCalled = FALSE;
            AttachRequest->Result                            = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
            return FALSE;
        }
        ResultOfApplyingEvent = DebuggerEventEnableMonitorReadAndWriteForAddress(
            PebAddressToMonitor,
            AttachRequest->ProcessId,
            TRUE,
            TRUE);
        if (!ResultOfApplyingEvent) {
            AttachingRemoveProcessDebuggingDetailsByToken(ProcessDebuggingToken);
            g_IsWaitingForUserModeModuleEntrypointToBeCalled = FALSE;
            AttachRequest->Result                            = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
            return FALSE;
        }
        AttachRequest->Token  = ProcessDebuggingToken;
        AttachRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        return TRUE;
    } else {
        if (!AttachingConfigureInterceptingThreads(ProcessDebuggingToken, TRUE)) {
            AttachingRemoveProcessDebuggingDetailsByToken(ProcessDebuggingToken);
            AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS;
            return FALSE;
        }
        AttachRequest->Token  = ProcessDebuggingToken;
        AttachRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        return TRUE;
    }
}

BOOLEAN
AttachingHandleCr3VmexitsForThreadInterception(UINT32 CurrentCoreIndex, CR3_TYPE NewCr3) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(PsGetCurrentProcessId());
    if (!ProcessDebuggingDetail || !ProcessDebuggingDetail->IsOnThreadInterceptingPhase) {
        HvUnsetExceptionBitmap(EXCEPTION_VECTOR_PAGE_FAULT);
        return FALSE;
    }
    for (size_t i = 0; i < MAX_CR3_IN_A_PROCESS; i++) {
        if (ProcessDebuggingDetail->InterceptedCr3[i].Flags == NewCr3.Flags) {
            break;
        }
        if (ProcessDebuggingDetail->InterceptedCr3[i].Flags == NULL) {
            ProcessDebuggingDetail->InterceptedCr3[i].Flags = NewCr3.Flags;
            break;
        }
    }
    if (!MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(NULL, FALSE, PagingLevelPageMapLevel4, NewCr3)) {
        return FALSE;
    }
    HvSetExceptionBitmap(EXCEPTION_VECTOR_PAGE_FAULT);
    return TRUE;
}

BOOLEAN
AttachingRemoveHooks(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS AttachRequest) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails;
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByToken(AttachRequest->Token);
    if (!ProcessDebuggingDetails) {
        AttachRequest->Result = DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN;
        return FALSE;
    }
    if (!g_IsWaitingForUserModeModuleEntrypointToBeCalled) {
        if (!EptHookUnHookSingleAddress(ProcessDebuggingDetails->PebAddressToMonitor,
                                        NULL,
                                        ProcessDebuggingDetails->ProcessId)) {
            AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS;
            return FALSE;
        } else {
            AttachRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
            return TRUE;
        }
    } else {
        AttachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED;
        return FALSE;
    }
}

BOOLEAN
AttachingPauseProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS PauseRequest) {
    if (AttachingConfigureInterceptingThreads(PauseRequest->Token, TRUE)) {
        PauseRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        return TRUE;
    } else {
        PauseRequest->Result = DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS;
        return FALSE;
    }
}

BOOLEAN
AttachingKillProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS KillRequest) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    BOOLEAN                             WasKilled = FALSE;
    ProcessDebuggingDetail                        = AttachingFindProcessDebuggingDetailsByProcessId(KillRequest->ProcessId);
    if (!ProcessDebuggingDetail) {
        KillRequest->Result = DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS;
        return FALSE;
    }
    WasKilled = KillProcess(KillRequest->ProcessId, PROCESS_KILL_METHOD_1);
    if (WasKilled) {
        goto Success;
    }
    WasKilled = KillProcess(KillRequest->ProcessId, PROCESS_KILL_METHOD_2);
    if (WasKilled) {
        goto Success;
    }
    WasKilled = KillProcess(KillRequest->ProcessId, PROCESS_KILL_METHOD_3);
    if (WasKilled) {
        goto Success;
    }
    KillRequest->Result = DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS;
    return FALSE;
Success:
    AttachingRemoveProcessDebuggingDetailsByToken(ProcessDebuggingDetail->Token);
    KillRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}

BOOLEAN
AttachingPerformDetach(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS DetachRequest) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(DetachRequest->ProcessId);
    if (!ProcessDebuggingDetail) {
        DetachRequest->Result = DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS;
        return FALSE;
    }
    if (ThreadHolderIsAnyPausedThreadInProcess(ProcessDebuggingDetail)) {
        DetachRequest->Result = DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS;
        return FALSE;
    }
    if (!MemoryMapperFreeMemoryInTargetProcess(DetachRequest->ProcessId,
                                               ProcessDebuggingDetail->UsermodeReservedBuffer)) {
        LogError("Err, cannot deallocate reserved buffer in the detached process");
    }
    AttachingRemoveProcessDebuggingDetailsByToken(ProcessDebuggingDetail->Token);
    DetachRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}

BOOLEAN
AttachingSwitchProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS SwitchRequest) {
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    PUSERMODE_DEBUGGING_THREAD_DETAILS  ThreadDebuggingDetail;
    if (SwitchRequest->ProcessId != NULL) {
        ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(SwitchRequest->ProcessId);
    } else if (SwitchRequest->ThreadId != NULL) {
        ProcessDebuggingDetail = ThreadHolderGetProcessDebuggingDetailsByThreadId(SwitchRequest->ThreadId);
    } else {
        SwitchRequest->Result = DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID;
        return FALSE;
    }
    if (!ProcessDebuggingDetail) {
        SwitchRequest->Result = DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID;
        return FALSE;
    }
    if (SwitchRequest->ThreadId != NULL) {
        ThreadDebuggingDetail = ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId(ProcessDebuggingDetail->ProcessId,
                                                                                          SwitchRequest->ThreadId);
    } else {
        ThreadDebuggingDetail = ThreadHolderGetProcessFirstThreadDetailsByProcessId(ProcessDebuggingDetail->ProcessId);
    }
    if (!ThreadDebuggingDetail) {
        SwitchRequest->Result = DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS;
        return FALSE;
    }
    SwitchRequest->Token     = ProcessDebuggingDetail->Token;
    SwitchRequest->ProcessId = ProcessDebuggingDetail->ProcessId;
    SwitchRequest->ThreadId  = ThreadDebuggingDetail->ThreadId;
    SwitchRequest->Is32Bit   = ProcessDebuggingDetail->Is32Bit;
    SwitchRequest->IsPaused  = ThreadDebuggingDetail->IsPaused;
    SwitchRequest->Result    = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}

BOOLEAN
AttachingQueryCountOfActiveDebuggingThreadsAndProcesses(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS QueryCountOfDebugThreadsRequest) {
    UINT32 CountOfProcessesAndThreads;
    CountOfProcessesAndThreads                                                 = ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses();
    QueryCountOfDebugThreadsRequest->CountOfActiveDebuggingThreadsAndProcesses = CountOfProcessesAndThreads;
    QueryCountOfDebugThreadsRequest->Result                                    = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}

BOOLEAN
AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses(PVOID BufferToStoreDetails, UINT32 BufferSize) {
    UINT32 CountOfProcessesAndThreadsToStore;
    CountOfProcessesAndThreadsToStore = BufferSize / SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS;
    if (CountOfProcessesAndThreadsToStore == 0) {
        return FALSE;
    }
    ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses(BufferToStoreDetails, CountOfProcessesAndThreadsToStore);
    return TRUE;
}

VOID
AttachingTargetProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS Request) {
    UdInitializeUserDebugger();
    switch (Request->Action) {
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH:
        AttachingPerformAttachToProcess(Request, Request->IsStartingNewProcess);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH:
        AttachingPerformDetach(Request);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS:
        AttachingRemoveHooks(Request);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS:
        AttachingKillProcess(Request);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS:
        AttachingPauseProcess(Request);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD:
        AttachingSwitchProcess(Request);
        break;
    case DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS:
        AttachingQueryCountOfActiveDebuggingThreadsAndProcesses(Request);
        break;
    default:
        Request->Result = DEBUGGER_ERROR_INVALID_ACTION_TYPE;
        break;
    }
}
