package header
//back\HyperDbgDev\hyperdbg\hprdbgctrl\header\ud.h.back

const(
DbgWaitForUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); } while (FALSE); //col:1
DbgReceivedUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE); //col:10
)

type ACTIVE_DEBUGGING_PROCESS struct{
IsActive bool
ProcessDebuggingToken uint64
ProcessId UINT32
ThreadId UINT32
IsPaused bool
Registers GUEST_REGS
Context uint64
Is32Bit bool
}




