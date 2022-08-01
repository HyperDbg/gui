package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\ud.h.back

const (
	DbgWaitForUserResponse (UserSyncObjectId) = do{DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId]; SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE)} while (
FALSE
)                                                                                                                                                                                                                                                                                       //col:1
DbgReceivedUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE) //col:10
)
type ACTIVE_DEBUGGING_PROCESS struct {
	IsActive              bool       //col:3
	ProcessDebuggingToken uint64     //col:4
	ProcessId             uint32     //col:5
	ThreadId              uint32     //col:6
	IsPaused              bool       //col:7
	Registers             GUEST_REGS //col:8
	Context               uint64     //col:9
	Is32Bit               bool       //col:10
}
