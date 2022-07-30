package user-level
//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\user-level\Ud.c.back

type (
Ud interface{
UdInitializeUserDebugger()(ok bool)//col:68
UdUninitializeUserDebugger()(ok bool)//col:92
UdRestoreToOriginalDirection()(ok bool)//col:108
UdContinueThread()(ok bool)//col:134
UdStepInstructions()(ok bool)//col:191
UdPerformCommand()(ok bool)//col:246
UdCheckForCommand()(ok bool)//col:318
UdDispatchUsermodeCommands()(ok bool)//col:359
UdSpinThreadOnNop()(ok bool)//col:386
UdHandleAfterSteppingReason()(ok bool)//col:415
UdPrePausingReasons()(ok bool)//col:453
UdCheckAndHandleBreakpointsAndDebugBreaks()(ok bool)//col:634
}
)

func NewUd() { return & ud{} }

func (u *ud)UdInitializeUserDebugger()(ok bool){//col:68
/*UdInitializeUserDebugger()
{
    if (g_UserDebuggerState)
    {
        return TRUE;
    }
    if (g_PsGetProcessPeb == NULL || g_PsGetProcessWow64Process == NULL || g_ZwQueryInformationProcess == NULL)
    {
        LogError("Err, unable to find needed functions for user-debugger");
    }
    g_SeedOfUserDebuggingDetails = DebuggerThreadDebuggingTagStartSeed;
    InitializeListHead(&g_ProcessDebuggingDetailsListHead);
    BroadcastEnableDbAndBpExitingAllCores();
    ThreadHolderAllocateThreadHoldingBuffers();
    g_UserDebuggerState = TRUE;
    return TRUE;
}*/
return true
}

func (u *ud)UdUninitializeUserDebugger()(ok bool){//col:92
/*UdUninitializeUserDebugger()
{
    if (g_UserDebuggerState)
    {
        g_UserDebuggerState = FALSE;
        AttachingRemoveAndFreeAllProcessDebuggingDetails();
    }
}*/
return true
}

func (u *ud)UdRestoreToOriginalDirection()(ok bool){//col:108
/*UdRestoreToOriginalDirection(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails)
{
    __vmx_vmwrite(VMCS_GUEST_RIP, ThreadDebuggingDetails->ThreadRip);
}*/
return true
}

func (u *ud)UdContinueThread()(ok bool){//col:134
/*UdContinueThread(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails)
{
    UdRestoreToOriginalDirection(ThreadDebuggingDetails);
    g_GuestState[KeGetCurrentProcessorNumber()].IncrementRip = FALSE;
    ThreadDebuggingDetails->IsPaused = FALSE;
}*/
return true
}

func (u *ud)UdStepInstructions()(ok bool){//col:191
/*UdStepInstructions(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
                   DEBUGGER_REMOTE_STEPPING_REQUEST   SteppingType)
{
    RFLAGS Rflags = {0};
    UdRestoreToOriginalDirection(ThreadDebuggingDetails);
    switch (SteppingType)
    {
    case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN:
        __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
        Rflags.TrapFlag = TRUE;
        __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
        ThreadDebuggingDetails->IsRflagsTrapFlagsSet = TRUE;
        break;
    case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER:
        break;
    default:
        break;
    }
    g_GuestState[KeGetCurrentProcessorNumber()].IncrementRip = FALSE;
    ThreadDebuggingDetails->IsPaused = FALSE;
}*/
return true
}

func (u *ud)UdPerformCommand()(ok bool){//col:246
/*UdPerformCommand(PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
                 DEBUGGER_UD_COMMAND_ACTION_TYPE    UserAction,
                 UINT64                             OptionalParam1,
                 UINT64                             OptionalParam2,
                 UINT64                             OptionalParam3,
                 UINT64                             OptionalParam4)
{
    switch (UserAction)
    {
    case DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE:
        UdContinueThread(ThreadDebuggingDetails);
        break;
    case DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP:
        UdStepInstructions(ThreadDebuggingDetails, (DEBUGGER_REMOTE_STEPPING_REQUEST)OptionalParam1);
        break;
    default:
        return FALSE;
        break;
    }
    return TRUE;
}*/
return true
}

func (u *ud)UdCheckForCommand()(ok bool){//col:318
/*UdCheckForCommand()
{
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails;
    ThreadDebuggingDetails =
        ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId(PsGetCurrentProcessId(),
                                                                  PsGetCurrentThreadId());
    if (!ThreadDebuggingDetails)
    {
        return FALSE;
    }
    if (!ThreadDebuggingDetails->IsPaused)
    {
        return FALSE;
    }
    for (size_t i = 0; i < MAX_USER_ACTIONS_FOR_THREADS; i++)
    {
        if (ThreadDebuggingDetails->UdAction[i].ActionType != DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE)
        {
            UdPerformCommand(ThreadDebuggingDetails,
                             ThreadDebuggingDetails->UdAction[i].ActionType,
                             ThreadDebuggingDetails->UdAction[i].OptionalParam1,
                             ThreadDebuggingDetails->UdAction[i].OptionalParam2,
                             ThreadDebuggingDetails->UdAction[i].OptionalParam3,
                             ThreadDebuggingDetails->UdAction[i].OptionalParam4);
            ThreadDebuggingDetails->UdAction[i].OptionalParam1 = NULL;
            ThreadDebuggingDetails->UdAction[i].OptionalParam2 = NULL;
            ThreadDebuggingDetails->UdAction[i].OptionalParam3 = NULL;
            ThreadDebuggingDetails->UdAction[i].OptionalParam4 = NULL;
            ThreadDebuggingDetails->UdAction[i].ActionType = DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE;
            break;
        }
    }
    return TRUE;
}*/
return true
}

func (u *ud)UdDispatchUsermodeCommands()(ok bool){//col:359
/*UdDispatchUsermodeCommands(PDEBUGGER_UD_COMMAND_PACKET ActionRequest)
{
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails;
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByToken(ActionRequest->ProcessDebuggingDetailToken);
    if (!ProcessDebuggingDetails)
    {
        return FALSE;
    }
    if (ProcessDebuggingDetails->IsOnThreadInterceptingPhase)
    {
        AttachingConfigureInterceptingThreads(ProcessDebuggingDetails->Token, FALSE);
    }
    return ThreadHolderApplyActionToPausedThreads(ProcessDebuggingDetails, ActionRequest);
}*/
return true
}

func (u *ud)UdSpinThreadOnNop()(ok bool){//col:386
/*UdSpinThreadOnNop(PUSERMODE_DEBUGGING_THREAD_DETAILS  ThreadDebuggingDetails,
                  PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails)
{
    __vmx_vmread(VMCS_GUEST_RIP, &ThreadDebuggingDetails->ThreadRip);
    __vmx_vmwrite(VMCS_GUEST_RIP, ProcessDebuggingDetails->UsermodeReservedBuffer);
    ThreadDebuggingDetails->IsPaused = TRUE;
}*/
return true
}

func (u *ud)UdHandleAfterSteppingReason()(ok bool){//col:415
/*UdHandleAfterSteppingReason(UINT32                             CurrentCore,
                            PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails)
{
    RFLAGS Rflags = {0};
    __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
    Rflags.TrapFlag = FALSE;
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
    ThreadDebuggingDetails->IsRflagsTrapFlagsSet = FALSE;
}*/
return true
}

func (u *ud)UdPrePausingReasons()(ok bool){//col:453
/*UdPrePausingReasons(UINT32                             CurrentCore,
                    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
                    PGUEST_REGS                        GuestRegs,
                    DEBUGGEE_PAUSING_REASON            Reason,
                    PDEBUGGER_TRIGGERED_EVENT_DETAILS  EventDetails)
{
    switch (Reason)
    {
    case DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK:
        if (ThreadDebuggingDetails->IsRflagsTrapFlagsSet)
        {
            UdHandleAfterSteppingReason(CurrentCore, ThreadDebuggingDetails);
        }
        break;
    default:
        break;
    }
}*/
return true
}

func (u *ud)UdCheckAndHandleBreakpointsAndDebugBreaks()(ok bool){//col:634
/*UdCheckAndHandleBreakpointsAndDebugBreaks(UINT32                            CurrentCore,
                                          PGUEST_REGS                       GuestRegs,
                                          DEBUGGEE_PAUSING_REASON           Reason,
                                          PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails)
{
    DEBUGGEE_UD_PAUSED_PACKET           PausePacket;
    ULONG                               ExitInstructionLength   = 0;
    UINT64                              SizeOfSafeBufferToRead  = 0;
    RFLAGS                              Rflags                  = {0};
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails = NULL;
    PUSERMODE_DEBUGGING_THREAD_DETAILS  ThreadDebuggingDetails  = NULL;
    if (!g_UserDebuggerState && !g_GuestState[CurrentCore].IsOnVmxRootMode)
    {
        return FALSE;
    }
    ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByProcessId(PsGetCurrentProcessId());
    if (!ProcessDebuggingDetails)
    {
        return FALSE;
    }
    ThreadDebuggingDetails = ThreadHolderFindOrCreateThreadDebuggingDetail(PsGetCurrentThreadId(), ProcessDebuggingDetails);
    if (!ThreadDebuggingDetails)
    {
        return FALSE;
    }
    ProcessDebuggingDetails->ActiveThreadId = ThreadDebuggingDetails->ThreadId;
    UdPrePausingReasons(CurrentCore, ThreadDebuggingDetails, GuestRegs, Reason, EventDetails);
    RtlZeroMemory(&PausePacket, sizeof(DEBUGGEE_UD_PAUSED_PACKET));
    PausePacket.PausingReason = Reason;
    PausePacket.ProcessId             = PsGetCurrentProcessId();
    PausePacket.ThreadId              = PsGetCurrentThreadId();
    PausePacket.ProcessDebuggingToken = ProcessDebuggingDetails->Token;
    PausePacket.Rip     = g_GuestState[CurrentCore].LastVmexitRip;
    PausePacket.Is32Bit = KdIsGuestOnUsermode32Bit();
    __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
    PausePacket.Rflags = Rflags.AsUInt;
    if (EventDetails != NULL)
    {
        PausePacket.EventTag = EventDetails->Tag;
    }
    if (g_GuestState[CurrentCore].DebuggingState.InstructionLengthHint != 0)
    {
        ExitInstructionLength = g_GuestState[CurrentCore].DebuggingState.InstructionLengthHint;
    }
    else
    {
        SizeOfSafeBufferToRead = g_GuestState[CurrentCore].LastVmexitRip & 0xfff;
        SizeOfSafeBufferToRead += MAXIMUM_INSTR_SIZE;
        if (SizeOfSafeBufferToRead >= PAGE_SIZE)
        {
            SizeOfSafeBufferToRead = SizeOfSafeBufferToRead - PAGE_SIZE;
            SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE - SizeOfSafeBufferToRead;
        }
        else
        {
            SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE;
        }
        ExitInstructionLength = SizeOfSafeBufferToRead;
    }
    PausePacket.ReadInstructionLen = ExitInstructionLength;
    MemoryMapperReadMemorySafeOnTargetProcess(g_GuestState[CurrentCore].LastVmexitRip,
                                              &PausePacket.InstructionBytesOnRip,
                                              ExitInstructionLength);
    RtlCopyMemory(&PausePacket.GuestRegs, GuestRegs, sizeof(GUEST_REGS));
    LogSendBuffer(OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE,
                  &PausePacket,
                  sizeof(DEBUGGEE_UD_PAUSED_PACKET),
                  TRUE);
    UdSpinThreadOnNop(ThreadDebuggingDetails, ProcessDebuggingDetails);
    return TRUE;
}*/
return true
}



