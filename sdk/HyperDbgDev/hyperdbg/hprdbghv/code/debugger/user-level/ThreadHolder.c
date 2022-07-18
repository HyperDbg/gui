#include "pch.h"
VOID
ThreadHolderAllocateThreadHoldingBuffers() {
    PoolManagerRequestAllocation(sizeof(USERMODE_DEBUGGING_THREAD_HOLDER), 2, PROCESS_THREAD_HOLDER);
    PoolManagerCheckAndPerformAllocationAndDeallocation();
}

BOOLEAN
ThreadHolderAssignThreadHolderToProcessDebuggingDetails(PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail) {
    PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder;
    InitializeListHead(&ProcessDebuggingDetail->ThreadsListHead);
    ThreadHolder = (USERMODE_DEBUGGING_THREAD_HOLDER *)
        PoolManagerRequestPool(PROCESS_THREAD_HOLDER, TRUE, sizeof(USERMODE_DEBUGGING_THREAD_HOLDER));
    if (!ThreadHolder) {
        return FALSE;
    }
    InsertHeadList(&ProcessDebuggingDetail->ThreadsListHead, &(ThreadHolder->ThreadHolderList));
    return TRUE;
}

BOOLEAN
ThreadHolderIsAnyPausedThreadInProcess(PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail) {
    PLIST_ENTRY TempList = 0;
    TempList             = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
            if (ThreadHolder->Threads[i].ThreadId != NULL && ThreadHolder->Threads[i].IsPaused) {
                return TRUE;
            }
        }
    }
    return FALSE;
}

PUSERMODE_DEBUGGING_THREAD_DETAILS
ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId(UINT32 ProcessId, UINT32 ThreadId) {
    PLIST_ENTRY                         TempList = 0;
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(ProcessId);
    if (ProcessDebuggingDetail == NULL) {
        return NULL;
    }
    TempList = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
            if (ThreadHolder->Threads[i].ThreadId == ThreadId) {
                return &ThreadHolder->Threads[i];
            }
        }
    }
    return NULL;
}

PUSERMODE_DEBUGGING_THREAD_DETAILS
ThreadHolderGetProcessFirstThreadDetailsByProcessId(UINT32 ProcessId) {
    PLIST_ENTRY                         TempList = 0;
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
    ProcessDebuggingDetail = AttachingFindProcessDebuggingDetailsByProcessId(ProcessId);
    if (ProcessDebuggingDetail == NULL) {
        return NULL;
    }
    TempList = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
            if (ThreadHolder->Threads[i].ThreadId != NULL) {
                return &ThreadHolder->Threads[i];
            }
        }
    }
    return NULL;
}

PUSERMODE_DEBUGGING_PROCESS_DETAILS
ThreadHolderGetProcessDebuggingDetailsByThreadId(UINT32 ThreadId) {
    PLIST_ENTRY TempList  = 0;
    PLIST_ENTRY TempList2 = 0;
    TempList              = &g_ProcessDebuggingDetailsListHead;
    while (&g_ProcessDebuggingDetailsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList);
        TempList2 = &ProcessDebuggingDetails->ThreadsListHead;
        while (&ProcessDebuggingDetails->ThreadsListHead != TempList2->Flink) {
            TempList2 = TempList2->Flink;
            PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
                CONTAINING_RECORD(TempList2, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
            for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
                if (ThreadHolder->Threads[i].ThreadId == ThreadId) {
                    return ProcessDebuggingDetails;
                }
            }
        }
    }
    return NULL;
}

PUSERMODE_DEBUGGING_THREAD_DETAILS
ThreadHolderFindOrCreateThreadDebuggingDetail(UINT32 ThreadId, PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail) {
    PLIST_ENTRY TempList = 0;
    TempList             = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
            if (ThreadHolder->Threads[i].ThreadId == ThreadId) {
                return &ThreadHolder->Threads[i];
            }
        }
    }
    SpinlockLock(&VmxRootThreadHoldingLock);
    TempList = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
            if (ThreadHolder->Threads[i].ThreadId == NULL) {
                ThreadHolder->Threads[i].ThreadId = ThreadId;
                SpinlockUnlock(&VmxRootThreadHoldingLock);
                return &ThreadHolder->Threads[i];
            }
        }
    }
    PUSERMODE_DEBUGGING_THREAD_HOLDER NewThreadHolder = (USERMODE_DEBUGGING_THREAD_HOLDER *)
        PoolManagerRequestPool(PROCESS_THREAD_HOLDER, TRUE, sizeof(USERMODE_DEBUGGING_THREAD_HOLDER));
    if (NewThreadHolder == NULL) {
        LogError("Err, enable to find a place to save the threads data, "
                 "please use 'prealloc' command to allocate more pre-allocated "
                 "buffer for the thread holder");
        SpinlockUnlock(&VmxRootThreadHoldingLock);
        return NULL;
    }
    NewThreadHolder->Threads[0].ThreadId = ThreadId;
    InsertHeadList(&ProcessDebuggingDetail->ThreadsListHead, &(NewThreadHolder->ThreadHolderList));
    SpinlockUnlock(&VmxRootThreadHoldingLock);
    return &NewThreadHolder->Threads[0];
}

BOOLEAN
ThreadHolderApplyActionToPausedThreads(PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails,
                                       PDEBUGGER_UD_COMMAND_PACKET         ActionRequest) {
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails;
    BOOLEAN                            CommandApplied = FALSE;
    if (!ActionRequest->ApplyToAllPausedThreads) {
        ThreadDebuggingDetails = ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId(ProcessDebuggingDetails->ProcessId,
                                                                                           ActionRequest->TargetThreadId);
        for (size_t i = 0; i < MAX_USER_ACTIONS_FOR_THREADS; i++) {
            if (ThreadDebuggingDetails->UdAction[i].ActionType == DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE) {
                ThreadDebuggingDetails->UdAction[i].OptionalParam1 = ActionRequest->UdAction.OptionalParam1;
                ThreadDebuggingDetails->UdAction[i].OptionalParam2 = ActionRequest->UdAction.OptionalParam2;
                ThreadDebuggingDetails->UdAction[i].OptionalParam3 = ActionRequest->UdAction.OptionalParam3;
                ThreadDebuggingDetails->UdAction[i].OptionalParam4 = ActionRequest->UdAction.OptionalParam4;
                ThreadDebuggingDetails->UdAction[i].ActionType     = ActionRequest->UdAction.ActionType;
                return TRUE;
            }
        }
    } else {
        PLIST_ENTRY TempList = 0;
        TempList             = &ProcessDebuggingDetails->ThreadsListHead;
        while (&ProcessDebuggingDetails->ThreadsListHead != TempList->Flink) {
            TempList = TempList->Flink;
            PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
                CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
            for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
                if (ThreadHolder->Threads[i].ThreadId != NULL &&
                    ThreadHolder->Threads[i].IsPaused) {
                    for (size_t j = 0; j < MAX_USER_ACTIONS_FOR_THREADS; j++) {
                        if (ThreadHolder->Threads[i].UdAction[j].ActionType == DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE) {
                            ThreadHolder->Threads[i].UdAction[j].OptionalParam1 = ActionRequest->UdAction.OptionalParam1;
                            ThreadHolder->Threads[i].UdAction[j].OptionalParam2 = ActionRequest->UdAction.OptionalParam2;
                            ThreadHolder->Threads[i].UdAction[j].OptionalParam3 = ActionRequest->UdAction.OptionalParam3;
                            ThreadHolder->Threads[i].UdAction[j].OptionalParam4 = ActionRequest->UdAction.OptionalParam4;
                            ThreadHolder->Threads[i].UdAction[j].ActionType     = ActionRequest->UdAction.ActionType;
                            CommandApplied                                      = TRUE;
                            break;
                        }
                    }
                }
            }
        }
    }
    return CommandApplied;
}

VOID
ThreadHolderFreeHoldingStructures(PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail) {
    PLIST_ENTRY TempList = 0;
    TempList             = &ProcessDebuggingDetail->ThreadsListHead;
    while (&ProcessDebuggingDetail->ThreadsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
        PoolManagerFreePool(ThreadHolder);
    }
}

UINT32
ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses() {
    PLIST_ENTRY TempList                   = 0;
    PLIST_ENTRY TempList2                  = 0;
    UINT32      CountOfThreadsAndProcesses = 0;
    TempList                               = &g_ProcessDebuggingDetailsListHead;
    while (&g_ProcessDebuggingDetailsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList);
        CountOfThreadsAndProcesses++;
        TempList2 = &ProcessDebuggingDetails->ThreadsListHead;
        while (&ProcessDebuggingDetails->ThreadsListHead != TempList2->Flink) {
            TempList2 = TempList2->Flink;
            PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
                CONTAINING_RECORD(TempList2, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
            for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
                if (ThreadHolder->Threads[i].IsPaused) {
                    CountOfThreadsAndProcesses++;
                }
            }
        }
    }
    return CountOfThreadsAndProcesses;
}

VOID
ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses(
    USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS * BufferToStoreDetails,
    UINT32                                               MaxCount) {
    PLIST_ENTRY TempList     = 0;
    PLIST_ENTRY TempList2    = 0;
    UINT32      CurrentIndex = 0;
    if (MaxCount == 0) {
        return;
    }
    TempList = &g_ProcessDebuggingDetailsListHead;
    while (&g_ProcessDebuggingDetailsListHead != TempList->Flink) {
        TempList = TempList->Flink;
        PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails =
            CONTAINING_RECORD(TempList, USERMODE_DEBUGGING_PROCESS_DETAILS, AttachedProcessList);
        BufferToStoreDetails[CurrentIndex].IsProcess = TRUE;
        BufferToStoreDetails[CurrentIndex].ProcessId = ProcessDebuggingDetails->ProcessId;
        CurrentIndex++;
        if (MaxCount == CurrentIndex) {
            return;
        }
        TempList2 = &ProcessDebuggingDetails->ThreadsListHead;
        while (&ProcessDebuggingDetails->ThreadsListHead != TempList2->Flink) {
            TempList2 = TempList2->Flink;
            PUSERMODE_DEBUGGING_THREAD_HOLDER ThreadHolder =
                CONTAINING_RECORD(TempList2, USERMODE_DEBUGGING_THREAD_HOLDER, ThreadHolderList);
            for (size_t i = 0; i < MAX_THREADS_IN_A_PROCESS_HOLDER; i++) {
                if (ThreadHolder->Threads[i].IsPaused) {
                    BufferToStoreDetails[CurrentIndex].IsProcess = FALSE;
                    BufferToStoreDetails[CurrentIndex].ProcessId = ProcessDebuggingDetails->ProcessId;
                    BufferToStoreDetails[CurrentIndex].ThreadId  = ThreadHolder->Threads[i].ThreadId;
                    CurrentIndex++;
                    if (MaxCount == CurrentIndex) {
                        return;
                    }
                }
            }
        }
    }
}
