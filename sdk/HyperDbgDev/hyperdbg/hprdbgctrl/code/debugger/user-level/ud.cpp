#include "pch.h"
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
extern BOOLEAN                  g_IsUserDebuggerInitialized;
extern DEBUGGER_SYNCRONIZATION_EVENTS_STATE
    g_UserSyncronizationObjectsHandleTable[DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS];
VOID
UdInitializeUserDebugger() {
    if (!g_IsUserDebuggerInitialized) {
        for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS; i++) {
            g_UserSyncronizationObjectsHandleTable[i].IsOnWaitingState = FALSE;
            g_UserSyncronizationObjectsHandleTable[i].EventHandle =
                CreateEvent(NULL, FALSE, FALSE, NULL);
        }
        g_IsUserDebuggerInitialized = TRUE;
    }
}

VOID
UdUninitializeUserDebugger() {
    if (g_IsUserDebuggerInitialized) {
        UdRemoveActiveDebuggingProcess(TRUE);
        for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS; i++) {
            if (g_UserSyncronizationObjectsHandleTable[i].EventHandle != NULL) {
                if (g_UserSyncronizationObjectsHandleTable[i].IsOnWaitingState) {
                    DbgReceivedUserResponse(i);
                }
                CloseHandle(g_UserSyncronizationObjectsHandleTable[i].EventHandle);
                g_UserSyncronizationObjectsHandleTable[i].EventHandle = NULL;
            }
        }
        g_IsUserDebuggerInitialized = FALSE;
    }
}

VOID
UdSetActiveDebuggingProcess(UINT64  DebuggingId,
                            UINT32  ProcessId,
                            UINT32  ThreadId,
                            BOOLEAN Is32Bit,
                            BOOLEAN IsPaused) {
    g_ActiveProcessDebuggingState.ProcessId             = ProcessId;
    g_ActiveProcessDebuggingState.ThreadId              = ThreadId;
    g_ActiveProcessDebuggingState.Is32Bit               = Is32Bit;
    g_ActiveProcessDebuggingState.ProcessDebuggingToken = DebuggingId;
    g_ActiveProcessDebuggingState.IsPaused              = IsPaused;
    g_ActiveProcessDebuggingState.IsActive              = TRUE;
}

VOID
UdRemoveActiveDebuggingProcess(BOOLEAN DontSwitchToNewProcess) {
    g_ActiveProcessDebuggingState.IsActive = FALSE;
}

VOID
UdPrintError() {
    DWORD   ErrNum;
    TCHAR   SysMsg[256];
    TCHAR * p;
    ErrNum = GetLastError();
    FormatMessage(FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_IGNORE_INSERTS,
                  NULL,
                  ErrNum,
                  MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT), // Default language
                  SysMsg,
                  256,
                  NULL);
    p = SysMsg;
    while ((*p > 31) || (*p == 9))
        ++p;
    do {
        *p-- = 0;
    } while ((p >= SysMsg) && ((*p == '.') || (*p < 33)));
    ShowMessages("\n  WARNING: Thread32First failed with error (%x:%s)", ErrNum, SysMsg);
}

BOOL
UdListProcessThreads(DWORD OwnerPID) {
    HANDLE        ThreadSnap = INVALID_HANDLE_VALUE;
    THREADENTRY32 Te32;
    ThreadSnap = CreateToolhelp32Snapshot(TH32CS_SNAPTHREAD, 0);
    if (ThreadSnap == INVALID_HANDLE_VALUE)
        return FALSE;
    Te32.dwSize = sizeof(THREADENTRY32);
    if (!Thread32First(ThreadSnap, &Te32)) {
        UdPrintError();          // Show cause of failure
        CloseHandle(ThreadSnap); // Must clean up the snapshot object!
        return FALSE;
    }
    ShowMessages("\nThread's of pid\t= 0x%08X\n", OwnerPID);
    do {
        if (Te32.th32OwnerProcessID == OwnerPID) {
            ShowMessages("\n     Thread Id\t= 0x%08X", Te32.th32ThreadID);
        }
    } while (Thread32Next(ThreadSnap, &Te32));
    ShowMessages("\n");
    CloseHandle(ThreadSnap);
    return TRUE;
}

BOOLEAN
UdCheckThreadByProcessId(DWORD Pid, DWORD Tid) {
    HANDLE        ThreadSnap = INVALID_HANDLE_VALUE;
    THREADENTRY32 Te32;
    BOOLEAN       Result = FALSE;
    ThreadSnap           = CreateToolhelp32Snapshot(TH32CS_SNAPTHREAD, 0);
    if (ThreadSnap == INVALID_HANDLE_VALUE)
        return FALSE;
    Te32.dwSize = sizeof(THREADENTRY32);
    if (!Thread32First(ThreadSnap, &Te32)) {
        UdPrintError();          // Show cause of failure
        CloseHandle(ThreadSnap); // Must clean up the snapshot object!
        return FALSE;
    }
    do {
        if (Te32.th32OwnerProcessID == Pid) {
            if (Te32.th32ThreadID == Tid) {
                Result = TRUE;
                break;
            }
        }
    } while (Thread32Next(ThreadSnap, &Te32));
    CloseHandle(ThreadSnap);
    return Result;
}

BOOLEAN
UdCreateSuspendedProcess(const WCHAR * FileName, WCHAR * CommandLine, PPROCESS_INFORMATION ProcessInformation) {
    STARTUPINFOW StartupInfo;
    BOOL         CreateProcessResult;
    memset(&StartupInfo, 0, sizeof(StartupInfo));
    StartupInfo.cb      = sizeof(STARTUPINFOA);
    CreateProcessResult = CreateProcessW(FileName,
                                         CommandLine,
                                         NULL,
                                         NULL,
                                         FALSE,
                                         CREATE_SUSPENDED,
                                         NULL,
                                         NULL,
                                         &StartupInfo,
                                         ProcessInformation);
    if (!CreateProcessResult) {
        ShowMessages("err, start process failed (%x)", GetLastError());
        return FALSE;
    }
    return TRUE;
}

BOOLEAN
UdAttachToProcess(UINT32        TargetPid,
                  const WCHAR * TargetFileAddress,
                  WCHAR *       CommandLine) {
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS AttachRequest = {0};
    PROCESS_INFORMATION                      ProcInfo      = {0};
    UdInitializeUserDebugger();
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    if (TargetFileAddress == NULL) {
        AttachRequest.IsStartingNewProcess = FALSE;
    } else {
        AttachRequest.IsStartingNewProcess = TRUE;
    }
    AttachRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH;
    if (AttachRequest.IsStartingNewProcess) {
        if (!IsFileExistW(TargetFileAddress)) {
            ShowMessages("err, unable to start, file not found\n");
            return FALSE;
        }
        UdCreateSuspendedProcess(TargetFileAddress, CommandLine, &ProcInfo);
        AttachRequest.ProcessId = ProcInfo.dwProcessId;
        AttachRequest.ThreadId  = ProcInfo.dwThreadId;
    } else {
        AttachRequest.ProcessId = TargetPid;
    }
    Status = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &AttachRequest,                                  // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &AttachRequest,                                  // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (AttachRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        if (!AttachRequest.IsStartingNewProcess) {
            ShowMessages("successfully attached to the target process!\n"
                         "please keep interacting with the process until all the "
                         "threads are intercepted and halted; whenever you execute "
                         "the first command, the thread interception will be stopped\n");
            return TRUE;
        }
        ResumeThread(ProcInfo.hThread);
        while (TRUE) {
            AttachRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS;
            Status               = DeviceIoControl(
                g_DeviceHandle,                                  // Handle to device
                IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
                &AttachRequest,                                  // Input Buffer to driver.
                SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
                &AttachRequest,                                  // Output Buffer from driver.
                SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
                &ReturnedLength,                                 // Bytes placed in buffer.
                NULL                                             // synchronous call
            );
            if (!Status) {
                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
                return FALSE;
            }
            if (AttachRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                break;
            } else if (AttachRequest.Result == DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED) {
                Sleep(1000);
                continue;
            } else {
                ShowErrorMessage(AttachRequest.Result);
                return FALSE;
            }
        }
        return TRUE;
    } else {
        ShowErrorMessage(AttachRequest.Result);
        return FALSE;
    }
    return FALSE;
}

BOOLEAN
UdKillProcess(UINT32 TargetPid) {
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS KillRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    KillRequest.Action    = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS;
    KillRequest.ProcessId = TargetPid;
    Status                = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &KillRequest,                                    // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &KillRequest,                                    // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (KillRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        UdRemoveActiveDebuggingProcess(FALSE);
        return TRUE;
    } else {
        ShowErrorMessage(KillRequest.Result);
        return FALSE;
    }
}

BOOLEAN
UdDetachProcess(UINT32 TargetPid, UINT64 ProcessDetailToken) {
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS DetachRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    UdContinueDebuggee(ProcessDetailToken);
    DetachRequest.Action    = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH;
    DetachRequest.ProcessId = TargetPid;
    Status                  = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &DetachRequest,                                  // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &DetachRequest,                                  // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (DetachRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        UdRemoveActiveDebuggingProcess(FALSE);
        return TRUE;
    } else {
        ShowErrorMessage(DetachRequest.Result);
        return FALSE;
    }
}

BOOLEAN
UdPauseProcess(UINT64 ProcessDebuggingToken) {
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS PauseRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    PauseRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS;
    PauseRequest.Token  = ProcessDebuggingToken;
    Status              = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &PauseRequest,                                   // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &PauseRequest,                                   // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (PauseRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        return TRUE;
    } else {
        ShowErrorMessage(PauseRequest.Result);
        return FALSE;
    }
}

VOID
UdSendCommand(UINT64                          ProcessDetailToken,
              UINT32                          ThreadId,
              DEBUGGER_UD_COMMAND_ACTION_TYPE ActionType,
              BOOLEAN                         ApplyToAllPausedThreads,
              UINT64                          OptionalParam1,
              UINT64                          OptionalParam2,
              UINT64                          OptionalParam3,
              UINT64                          OptionalParam4) {
    BOOL                       Status;
    ULONG                      ReturnedLength;
    DEBUGGER_UD_COMMAND_PACKET CommandPacket;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    RtlZeroMemory(&CommandPacket, sizeof(DEBUGGER_UD_COMMAND_PACKET));
    CommandPacket.ProcessDebuggingDetailToken = ProcessDetailToken;
    CommandPacket.ApplyToAllPausedThreads     = ApplyToAllPausedThreads;
    CommandPacket.TargetThreadId              = ThreadId;
    CommandPacket.UdAction.ActionType         = ActionType;
    CommandPacket.UdAction.OptionalParam1     = OptionalParam1;
    CommandPacket.UdAction.OptionalParam2     = OptionalParam2;
    CommandPacket.UdAction.OptionalParam3     = OptionalParam3;
    CommandPacket.UdAction.OptionalParam4     = OptionalParam4;
    Status =
        DeviceIoControl(g_DeviceHandle,                     // Handle to device
                        IOCTL_SEND_USER_DEBUGGER_COMMANDS,  // IO Control code
                        &CommandPacket,                     // Input Buffer to driver.
                        sizeof(DEBUGGER_UD_COMMAND_PACKET), // Input buffer length
                        &CommandPacket,                     // Output Buffer from driver.
                        sizeof(DEBUGGER_UD_COMMAND_PACKET), // Length of output buffer in bytes.
                        &ReturnedLength,                    // Bytes placed in buffer.
                        NULL                                // synchronous call
        );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
}

VOID
UdContinueDebuggee(UINT64 ProcessDetailToken) {
    UdSendCommand(ProcessDetailToken,
                  NULL,
                  DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE,
                  TRUE,
                  NULL,
                  NULL,
                  NULL,
                  NULL);
}

VOID
UdSendStepPacketToDebuggee(UINT64 ProcessDetailToken, UINT32 TargetThreadId, DEBUGGER_REMOTE_STEPPING_REQUEST StepType) {
    g_UserSyncronizationObjectsHandleTable
        [DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING]
            .IsOnWaitingState = TRUE;
    UdSendCommand(ProcessDetailToken,
                  TargetThreadId,
                  DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP,
                  FALSE,
                  StepType,
                  NULL,
                  NULL,
                  NULL);
    WaitForSingleObject(
        g_UserSyncronizationObjectsHandleTable
            [DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING]
                .EventHandle,
        INFINITE);
}

BOOLEAN
UdSetActiveDebuggingThreadByPidOrTid(UINT32 TargetPidOrTid, BOOLEAN IsTid) {
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS SwitchRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    SwitchRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD;
    if (IsTid) {
        SwitchRequest.ThreadId = TargetPidOrTid;
    } else {
        SwitchRequest.ProcessId = TargetPidOrTid;
    }
    Status = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &SwitchRequest,                                  // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &SwitchRequest,                                  // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (SwitchRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        UdSetActiveDebuggingProcess(SwitchRequest.Token,
                                    SwitchRequest.ProcessId,
                                    SwitchRequest.ThreadId,
                                    SwitchRequest.Is32Bit,
                                    SwitchRequest.IsPaused);
        return TRUE;
    } else {
        ShowErrorMessage(SwitchRequest.Result);
        return FALSE;
    }
}

BOOLEAN
UdShowListActiveDebuggingProcessesAndThreads() {
    BOOLEAN                                              Status;
    BOOLEAN                                              CheckCurrentProcessOrThread;
    ULONG                                                ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS             QueryCountOfActiveThreadsRequest        = {0};
    USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS * AddressOfThreadsAndProcessDetails       = NULL;
    UINT32                                               SizeOfBufferForThreadsAndProcessDetails = NULL;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    if (!g_IsUserDebuggerInitialized) {
        ShowMessages("user debugger is not initialized. Did you use the '.attach' or the '.start' "
                     "command before?\n");
        return FALSE;
    }
    QueryCountOfActiveThreadsRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS;
    Status                                  = DeviceIoControl(
        g_DeviceHandle,                                  // Handle to device
        IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // IO Control
        &QueryCountOfActiveThreadsRequest,               // Input Buffer to driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Input buffer length
        &QueryCountOfActiveThreadsRequest,               // Output Buffer from driver.
        SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS, // Length of output
        &ReturnedLength,                                 // Bytes placed in buffer.
        NULL                                             // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (QueryCountOfActiveThreadsRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        if (QueryCountOfActiveThreadsRequest.CountOfActiveDebuggingThreadsAndProcesses == 0) {
            ShowMessages("no active debugging threads!\n");
        } else {
            SizeOfBufferForThreadsAndProcessDetails = QueryCountOfActiveThreadsRequest.CountOfActiveDebuggingThreadsAndProcesses * SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS;
            AddressOfThreadsAndProcessDetails       = (USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS *)
                malloc(SizeOfBufferForThreadsAndProcessDetails);
            RtlZeroMemory(AddressOfThreadsAndProcessDetails, SizeOfBufferForThreadsAndProcessDetails);
            Status = DeviceIoControl(
                g_DeviceHandle,                                   // Handle to device
                IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES, // IO Control
                NULL,                                             // Input Buffer to driver.
                0,                                                // Input buffer length.
                AddressOfThreadsAndProcessDetails,                // Output Buffer from driver.
                SizeOfBufferForThreadsAndProcessDetails,          // Length of output buffer in bytes.
                &ReturnedLength,                                  // Bytes placed in buffer.
                NULL                                              // synchronous call
            );
            if (!Status) {
                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
                return FALSE;
            }
            for (size_t i = 0; i < QueryCountOfActiveThreadsRequest.CountOfActiveDebuggingThreadsAndProcesses; i++) {
                if (AddressOfThreadsAndProcessDetails[i].IsProcess) {
                    CheckCurrentProcessOrThread = FALSE;
                    if (g_ActiveProcessDebuggingState.IsActive &&
                        AddressOfThreadsAndProcessDetails[i].ProcessId == g_ActiveProcessDebuggingState.ProcessId) {
                        CheckCurrentProcessOrThread = TRUE;
                    }
                    ShowMessages("%s%04x (process)\n",
                                 CheckCurrentProcessOrThread ? "*" : "",
                                 AddressOfThreadsAndProcessDetails[i].ProcessId);
                } else {
                    CheckCurrentProcessOrThread = FALSE;
                    if (g_ActiveProcessDebuggingState.IsActive &&
                        AddressOfThreadsAndProcessDetails[i].ThreadId == g_ActiveProcessDebuggingState.ThreadId) {
                        CheckCurrentProcessOrThread = TRUE;
                    }
                    ShowMessages("\t%s %04x (thread)\n",
                                 CheckCurrentProcessOrThread ? "->" : "  ",
                                 AddressOfThreadsAndProcessDetails[i].ThreadId);
                }
            }
        }
        return TRUE;
    } else {
        ShowErrorMessage(QueryCountOfActiveThreadsRequest.Result);
        return FALSE;
    }
}
