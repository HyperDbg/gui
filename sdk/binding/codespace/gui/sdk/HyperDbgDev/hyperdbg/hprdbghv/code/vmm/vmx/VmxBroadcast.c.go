package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxBroadcast.c.back

type (
VmxBroadcast interface{
VmxBroadcastHandleNmiCallback()(ok bool)//col:13
VmxBroadcastNmi()(ok bool)//col:35
VmxBroadcastHandleKdDebugBreaks()(ok bool)//col:49
VmxBroadcastNmiHandler()(ok bool)//col:77
}
)

func NewVmxBroadcast() { return & vmxBroadcast{} }

func (v *vmxBroadcast)VmxBroadcastHandleNmiCallback()(ok bool){//col:13
/*VmxBroadcastHandleNmiCallback(PVOID Context, BOOLEAN Handled)
{
    ULONG CurrentCoreIndex;
    CurrentCoreIndex = KeGetCurrentProcessorNumber();
    if (!VmxBroadcastNmiHandler(CurrentCoreIndex, NULL, TRUE))
    {
        return Handled;
    }
    else
    {
        return TRUE;
    }
}*/
return true
}

func (v *vmxBroadcast)VmxBroadcastNmi()(ok bool){//col:35
/*VmxBroadcastNmi(UINT32 CurrentCoreIndex, NMI_BROADCAST_ACTION_TYPE VmxBroadcastAction)
{
    ULONG CoreCount;
    if (!g_NmiBroadcastingInitialized)
    {
        return FALSE;
    }
    CoreCount = KeQueryActiveProcessorCount(0);
    SpinlockLock(&DebuggerResponseLock);
    for (size_t i = 0; i < CoreCount; i++)
    {
        if (i != CurrentCoreIndex)
        {
            SpinlockInterlockedCompareExchange((volatile LONG *)&g_GuestState[i].DebuggingState.NmiBroadcastAction,
                                               VmxBroadcastAction,
                                               NMI_BROADCAST_ACTION_NONE);
        }
    }
    ApicTriggerGenericNmi();
    SpinlockUnlock(&DebuggerResponseLock);
    return TRUE;
}*/
return true
}

func (v *vmxBroadcast)VmxBroadcastHandleKdDebugBreaks()(ok bool){//col:49
/*VmxBroadcastHandleKdDebugBreaks(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs, BOOLEAN IsOnVmxNmiHandler)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCoreIndex];
    CurrentVmState->DebuggingState.WaitingToBeLocked = TRUE;
    if (IsOnVmxNmiHandler)
    {
        CurrentVmState->DebuggingState.NmiCalledInVmxRootRelatedToHaltDebuggee = TRUE;
        HvSetMonitorTrapFlag(TRUE);
    }
    else
    {
        KdHandleNmi(CurrentCoreIndex, GuestRegs);
    }
}*/
return true
}

func (v *vmxBroadcast)VmxBroadcastNmiHandler()(ok bool){//col:77
/*VmxBroadcastNmiHandler(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs, BOOLEAN IsOnVmxNmiHandler)
{
    NMI_BROADCAST_ACTION_TYPE   BroadcastAction;
    BOOLEAN                     IsHandled            = FALSE;
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = &g_GuestState[CurrentCoreIndex].DebuggingState;
    BroadcastAction = InterlockedExchange((volatile LONG *)&CurrentDebuggerState->NmiBroadcastAction, NMI_BROADCAST_ACTION_NONE);
    if (BroadcastAction == NMI_BROADCAST_ACTION_NONE)
    {
        IsHandled = FALSE;
        goto ReturnIsHandled;
    }
    switch (BroadcastAction)
    {
    case NMI_BROADCAST_ACTION_TEST:
        IsHandled = TRUE;
        break;
    case NMI_BROADCAST_ACTION_KD_HALT_CORE:
        IsHandled = TRUE;
        VmxBroadcastHandleKdDebugBreaks(CurrentCoreIndex, GuestRegs, IsOnVmxNmiHandler);
        break;
    default:
        IsHandled = FALSE;
        LogError("Err, invalid NMI reason received");
        break;
    }
ReturnIsHandled:
    return IsHandled;
}*/
return true
}



