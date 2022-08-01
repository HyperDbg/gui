package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Mtf.c.back

type (
Mtf interface{
MtfHandleVmexit()(ok bool)//col:102
}
mtf struct{}
)

func NewMtf()Mtf{ return & mtf{} }

func (m *mtf)MtfHandleVmexit()(ok bool){//col:102
/*MtfHandleVmexit(ULONG CurrentProcessorIndex, PGUEST_REGS GuestRegs)
{
    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag = {0};
    BOOLEAN                          AvoidUnsetMtf;
    UINT64                      CsSel                 = NULL;
    BOOLEAN                     IsMtfHandled          = FALSE;
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentProcessorIndex];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    CurrentVmState->IncrementRip = FALSE;
    CurrentVmState->IgnoreMtfUnset = FALSE;
    if (CurrentDebuggingState->SoftwareBreakpointState != NULL)
    {
        BYTE BreakpointByte = 0xcc;
        IsMtfHandled = TRUE;
        MemoryMapperWriteMemorySafeByPhysicalAddress(
            CurrentDebuggingState->SoftwareBreakpointState->PhysAddress,
            &BreakpointByte,
            sizeof(BYTE));
        if (CurrentDebuggingState->SoftwareBreakpointState->SetRflagsIFBitOnMtf)
        {
            RFLAGS Rflags = {0};
            __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
            Rflags.InterruptEnableFlag = TRUE;
            __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
            CurrentDebuggingState->SoftwareBreakpointState->SetRflagsIFBitOnMtf = FALSE;
        }
        CurrentDebuggingState->SoftwareBreakpointState = NULL;
    }
    if (CurrentVmState->MtfEptHookRestorePoint)
    {
        IsMtfHandled = TRUE;
        if (g_IsWaitingForUserModeModuleEntrypointToBeCalled)
        {
            UserAccessCheckForLoadedModuleDetails();
        }
        EptHandleMonitorTrapFlag(CurrentVmState->MtfEptHookRestorePoint);
        CurrentVmState->MtfEptHookRestorePoint = NULL;
        if (CurrentDebuggingState->EnableExternalInterruptsOnContinueMtf)
        {
            HvSetExternalInterruptExiting(FALSE);
            if (CurrentVmState->PendingExternalInterrupts[0] != NULL)
            {
                HvSetInterruptWindowExiting(TRUE);
            }
            CurrentDebuggingState->EnableExternalInterruptsOnContinueMtf = FALSE;
        }
    }
    if (CurrentDebuggingState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf)
    {
        IsMtfHandled = TRUE;
        __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &CsSel);
        KdCheckGuestOperatingModeChanges(CurrentDebuggingState->InstrumentationStepInTrace.CsSel,
                                         (UINT16)CsSel);
        CurrentDebuggingState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf = FALSE;
        CurrentDebuggingState->InstrumentationStepInTrace.CsSel                           = 0;
        if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(CurrentProcessorIndex,
                                                                CurrentVmState->LastVmexitRip,
                                                                DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
                                                                GuestRegs,
                                                                &AvoidUnsetMtf))
        {
            ContextAndTag.Context = CurrentVmState->LastVmexitRip;
            KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
                                                  GuestRegs,
                                                  DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
                                                  &ContextAndTag);
        }
        else
        {
            CurrentVmState->IgnoreMtfUnset = AvoidUnsetMtf;
        }
    }
    else if (CurrentDebuggingState->WaitingToBeLocked)
    {
        IsMtfHandled = TRUE;
        if (CurrentDebuggingState->NmiCalledInVmxRootRelatedToHaltDebuggee)
        {
            KdHandleHaltsWhenNmiReceivedFromVmxRoot(CurrentProcessorIndex, GuestRegs);
        }
        else
        {
            KdHandleNmi(CurrentProcessorIndex, GuestRegs);
        }
    }
    if (CurrentDebuggingState->IgnoreOneMtf)
    {
        IsMtfHandled = TRUE;
        CurrentDebuggingState->IgnoreOneMtf = FALSE;
    }
    if (!IsMtfHandled)
    {
        LogError("Err, why MTF occurred?!");
    }
    if (!CurrentVmState->IgnoreMtfUnset)
    {
        HvSetMonitorTrapFlag(FALSE);
    }
    else
    {
        CurrentVmState->IgnoreMtfUnset = FALSE;
    }
}*/
return true
}



