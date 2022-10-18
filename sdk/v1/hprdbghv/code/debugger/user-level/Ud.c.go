package user_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\user-level\Ud.c.back

type (
	Ud interface {
		UdInitializeUserDebugger() (ok bool) //col:27
		UdStepInstructions() (ok bool)       //col:40
	}
	ud struct{}
)

func NewUd() Ud { return &ud{} }

func (u *ud) UdInitializeUserDebugger() (ok bool) { //col:27
	/*
	   UdInitializeUserDebugger()

	   	{
	   	    if (g_UserDebuggerState)
	   	    {
	   	        return TRUE;
	   	    }
	   	    if (g_PsGetProcessPeb == NULL || g_PsGetProcessWow64Process == NULL || g_ZwQueryInformationProcess == NULL)
	   	    {
	   	        LogError("Err, unable to find needed functions for user-debugger");
	   	    InitializeListHead(&g_ProcessDebuggingDetailsListHead);
	   	    BroadcastEnableDbAndBpExitingAllCores();
	   	    ThreadHolderAllocateThreadHoldingBuffers();

	   UdUninitializeUserDebugger()

	   	{
	   	    if (g_UserDebuggerState)
	   	    {
	   	        g_UserDebuggerState = FALSE;
	   	        AttachingRemoveAndFreeAllProcessDebuggingDetails();

	   UdRestoreToOriginalDirection(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails)

	   	{
	   	    __vmx_vmwrite(VMCS_GUEST_RIP, ThreadDebuggingDetails->ThreadRip);

	   UdContinueThread(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails)

	   	{
	   	    UdRestoreToOriginalDirection(ThreadDebuggingDetails);
	   	    g_GuestState[KeGetCurrentProcessorNumber()].IncrementRip = FALSE;
	   	    ThreadDebuggingDetails->IsPaused = FALSE;
	   	}
	*/
	return true
}

func (u *ud) UdStepInstructions() (ok bool) { //col:40
	/*
	   UdStepInstructions(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,

	   	DEBUGGER_REMOTE_STEPPING_REQUEST   SteppingType)

	   	{
	   	    RFLAGS Rflags = {0};
	   	    UdRestoreToOriginalDirection(ThreadDebuggingDetails);
	   	    switch (SteppingType)
	   	    {
	   	    case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN:
	   	        __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
	   	        __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
	   	    g_GuestState[KeGetCurrentProcessorNumber()].IncrementRip = FALSE;
	   	    ThreadDebuggingDetails->IsPaused = FALSE;
	   	}
	*/
	return true
}

