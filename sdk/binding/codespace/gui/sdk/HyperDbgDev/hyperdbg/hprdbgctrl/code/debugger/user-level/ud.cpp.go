package user_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\user-level\ud.cpp.back

type (
	Ud interface {
		UdInitializeUserDebugger() (ok bool)       //col:35
		UdRemoveActiveDebuggingProcess() (ok bool) //col:39
	}
	ud struct{}
)

func NewUd() Ud { return &ud{} }

func (u *ud) UdInitializeUserDebugger() (ok bool) { //col:35
	/*
	   UdInitializeUserDebugger()

	   	{
	   	    if (!g_IsUserDebuggerInitialized)
	   	    {
	   	        for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS; i++)
	   	        {
	   	            g_UserSyncronizationObjectsHandleTable[i].IsOnWaitingState = FALSE;
	   	            g_UserSyncronizationObjectsHandleTable[i].EventHandle =
	   	                CreateEvent(NULL, FALSE, FALSE, NULL);

	   UdUninitializeUserDebugger()

	   	{
	   	    if (g_IsUserDebuggerInitialized)
	   	    {
	   	        UdRemoveActiveDebuggingProcess(TRUE);
	   	        for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS; i++)
	   	        {
	   	            if (g_UserSyncronizationObjectsHandleTable[i].EventHandle != NULL)
	   	            {
	   	                if (g_UserSyncronizationObjectsHandleTable[i].IsOnWaitingState)
	   	                {
	   	                    DbgReceivedUserResponse(i);
	   	                CloseHandle(g_UserSyncronizationObjectsHandleTable[i].EventHandle);

	   UdSetActiveDebuggingProcess(UINT64  DebuggingId,

	   	UINT32  ProcessId,
	   	UINT32  ThreadId,
	   	BOOLEAN Is32Bit,
	   	BOOLEAN IsPaused)

	   	{
	   	    g_ActiveProcessDebuggingState.ProcessId             = ProcessId;
	   	    g_ActiveProcessDebuggingState.ThreadId              = ThreadId;
	   	    g_ActiveProcessDebuggingState.Is32Bit               = Is32Bit;
	   	    g_ActiveProcessDebuggingState.ProcessDebuggingToken = DebuggingId;
	   	    g_ActiveProcessDebuggingState.IsPaused = IsPaused;
	   	    g_ActiveProcessDebuggingState.IsActive = TRUE;
	   	}
	*/
	return true
}

func (u *ud) UdRemoveActiveDebuggingProcess() (ok bool) { //col:39
	/*
	   UdRemoveActiveDebuggingProcess(BOOLEAN DontSwitchToNewProcess)

	   	{
	   	    g_ActiveProcessDebuggingState.IsActive = FALSE;
	   	}
	*/
	return true
}

