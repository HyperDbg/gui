package header


const(
DbgWaitForKernelResponse(KernelSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId];  SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); } while (FALSE); //col:18
DbgReceivedKernelResponse(KernelSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE); //col:28
)

type (
Kd interface{
        WaitForSingleObject()(ok bool)//col:63
}







































)

func NewKd() { return & kd{} }

func (k *kd)        WaitForSingleObject()(ok bool){//col:63




























return true
}










































