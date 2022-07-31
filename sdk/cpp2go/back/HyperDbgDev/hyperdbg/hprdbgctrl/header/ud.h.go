package header
//back\HyperDbgDev\hyperdbg\hprdbgctrl\header\ud.h.back

const(
DbgWaitForUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); } while (FALSE); //col:18
DbgReceivedUserResponse(UserSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE); //col:28
)

type (
Ud interface{
#define DbgWaitForUserResponse()(ok bool)//col:57
}
)

func NewUd() { return & ud{} }

func (u *ud)#define DbgWaitForUserResponse()(ok bool){//col:57
/*#define DbgWaitForUserResponse(UserSyncObjectId)                          \
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
 * state
 *
typedef struct _ACTIVE_DEBUGGING_PROCESS
{
    BOOLEAN    IsActive;
    UINT64     ProcessDebuggingToken;
    UINT32     ProcessId;
    UINT32     ThreadId;
    BOOLEAN    IsPaused;
    BOOLEAN    Is32Bit;
} ACTIVE_DEBUGGING_PROCESS, *PACTIVE_DEBUGGING_PROCESS;*/
return true
}



