#pragma once
#define DbgWaitForUserResponse(UserSyncObjectId)                          \
    do                                                                    \
    {                                                                     \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =     \
            &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];    \
                                                                          \
        SyncronizationObject->IsOnWaitingState = TRUE;                    \
        WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); \
    } while (FALSE);
#define DbgReceivedUserResponse(UserSyncObjectId)                      \
    do                                                                 \
    {                                                                  \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =  \
            &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId]; \
                                                                       \
        SyncronizationObject->IsOnWaitingState = FALSE;                \
        SetEvent(SyncronizationObject->EventHandle);                   \
    } while (FALSE);
typedef struct _ACTIVE_DEBUGGING_PROCESS
{
    BOOLEAN    IsActive;
    UINT64     ProcessDebuggingToken;
    UINT32     ProcessId;
    UINT32     ThreadId;
    BOOLEAN    IsPaused;
    GUEST_REGS Registers; 
    UINT64     Context;   
    BOOLEAN    Is32Bit;
} ACTIVE_DEBUGGING_PROCESS, *PACTIVE_DEBUGGING_PROCESS;
VOID
UdInitializeUserDebugger();
VOID
UdUninitializeUserDebugger();
VOID
UdRemoveActiveDebuggingProcess(BOOLEAN DontSwitchToNewProcess);
VOID
UdHandleUserDebuggerPausing(PDEBUGGEE_UD_PAUSED_PACKET PausePacket);
VOID
UdContinueDebuggee(UINT64 ProcessDetailToken);
VOID
UdSendStepPacketToDebuggee(UINT64 ThreadDetailToken, UINT32 TargetThreadId, DEBUGGER_REMOTE_STEPPING_REQUEST StepType);
VOID
UdSetActiveDebuggingProcess(UINT64  DebuggingId,
                            UINT32  ProcessId,
                            UINT32  ThreadId,
                            BOOLEAN Is32Bit,
                            BOOLEAN IsPaused);
BOOLEAN
UdSetActiveDebuggingThreadByPidOrTid(UINT32 TargetPidOrTid, BOOLEAN IsTid);
BOOLEAN
UdSetActiveDebuggingThreadByPidOrTid(UINT32 TargetPidOrTid, BOOLEAN IsTid);
BOOLEAN
UdShowListActiveDebuggingProcessesAndThreads();
BOOL
UdListProcessThreads(DWORD OwnerPID);
BOOLEAN
UdCheckThreadByProcessId(DWORD Pid, DWORD Tid);
BOOLEAN
UdAttachToProcess(UINT32        TargetPid,
                  const WCHAR * TargetFileAddress,
                  WCHAR *       CommandLine);
BOOLEAN
UdKillProcess(UINT32 TargetPid);
BOOLEAN
UdDetachProcess(UINT32 TargetPid, UINT64 ProcessDetailToken);
BOOLEAN
UdPauseProcess(UINT64 ProcessDebuggingToken);
