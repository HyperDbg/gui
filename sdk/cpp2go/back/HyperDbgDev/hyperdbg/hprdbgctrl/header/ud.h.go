package header


const(
DbgWaitForUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); } while (FALSE); //col:18
DbgReceivedUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE); //col:28
)

type (
Ud interface{
        WaitForSingleObject()(ok bool)//col:57
}







































)

func NewUd() { return & ud{} }

func (u *ud)        WaitForSingleObject()(ok bool){//col:57




















return true
}










































