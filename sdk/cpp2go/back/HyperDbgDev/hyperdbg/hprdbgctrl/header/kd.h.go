package header
//back\HyperDbgDev\hyperdbg\hprdbgctrl\header\kd.h.back

const(
DbgWaitForKernelResponse(KernelSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId];  SyncronizationObject->IsOnWaitingState = TRUE; WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); } while (FALSE); //col:18
DbgReceivedKernelResponse(KernelSyncObjectId) = do { DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject = &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId];  SyncronizationObject->IsOnWaitingState = FALSE; SetEvent(SyncronizationObject->EventHandle); } while (FALSE); //col:28
)

type (
Kd interface{
#define DbgWaitForKernelResponse()(ok bool)//col:63
}
)

func NewKd() { return & kd{} }

func (k *kd)#define DbgWaitForKernelResponse()(ok bool){//col:63
/*#define DbgWaitForKernelResponse(KernelSyncObjectId)                       \
    do                                                                     \
    {                                                                      \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =      \
            &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId]; \
                                                                           \
        SyncronizationObject->IsOnWaitingState = TRUE;                     \
        WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE);  \
    } while (FALSE);
#define DbgReceivedKernelResponse(KernelSyncObjectId)                      \
    do                                                                     \
    {                                                                      \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =      \
            &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId]; \
                                                                           \
        SyncronizationObject->IsOnWaitingState = FALSE;                    \
        SetEvent(SyncronizationObject->EventHandle);                       \
    } while (FALSE);
struct HKeyHolder
{
private:
    HKEY m_Key;
public:
    HKeyHolder() :
        m_Key(nullptr) { }
    HKeyHolder(const HKeyHolder &) = delete;
    HKeyHolder & operator=(const HKeyHolder &) = delete;
    ~HKeyHolder()
    {
        if (m_Key != nullptr)
            RegCloseKey(m_Key);
    }
    operator HKEY() const { return m_Key; }
    HKEY * operator&() { return &m_Key; }
};*/
return true
}



