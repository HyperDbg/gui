#include "pch.h"
BOOLEAN
IdtEmulationReInjectInterruptOrException(_In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit) {
    ULONG ErrorCode = 0;
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, InterruptExit.AsUInt);
    if (InterruptExit.ErrorCodeValid) {
        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
    }
}

BOOLEAN
IdtEmulationHandlePageFaults(_In_ UINT32                       CurrentProcessorIndex,
                             _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                             _In_ UINT64                       Address,
                             _In_ ULONG                        ErrorCode) {
    PAGE_FAULT_ERROR_CODE   PageFaultCode     = {0};
    VIRTUAL_MACHINE_STATE * CurrentGuestState = &g_GuestState[CurrentProcessorIndex];
    __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &PageFaultCode);
    if (Address == NULL) {
        UINT64 PageFaultAddress = 0;
        __vmx_vmread(VMCS_EXIT_QUALIFICATION, &PageFaultAddress);
        __writecr2(PageFaultAddress);
    } else {
        __writecr2(Address);
    }
    CurrentGuestState->IncrementRip = FALSE;
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, InterruptExit.AsUInt);
    if (InterruptExit.ErrorCodeValid) {
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
    }
}

VOID
IdtEmulationHandleExceptionAndNmi(_In_ UINT32                          CurrentProcessorIndex,
                                  _Inout_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                  _Inout_ PGUEST_REGS                  GuestRegs) {
    ULONG                       ErrorCode            = 0;
    VIRTUAL_MACHINE_STATE *     CurrentGuestState    = &g_GuestState[CurrentProcessorIndex];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = &g_GuestState[CurrentProcessorIndex].DebuggingState;
    if (InterruptExit.InterruptionType == INTERRUPT_TYPE_NMI &&
        InterruptExit.Vector == EXCEPTION_VECTOR_NMI) {
        if (!CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf &&
            VmxBroadcastNmiHandler(CurrentProcessorIndex, GuestRegs, FALSE)) {
            return;
        }
    }
    if (CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf) {
        return;
    }
    DebuggerTriggerEvents(EXCEPTION_OCCURRED, GuestRegs, InterruptExit.Vector);
    if (CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf) {
        return;
    }
    switch (InterruptExit.Vector) {
    case EXCEPTION_VECTOR_BREAKPOINT:
        BreakpointHandleBpTraps(CurrentProcessorIndex, GuestRegs);
        break;
    case EXCEPTION_VECTOR_UNDEFINED_OPCODE:
        if (!SyscallHookHandleUD(GuestRegs, CurrentProcessorIndex)) {
            EventInjectUndefinedOpcode(CurrentProcessorIndex);
        }
        break;
    case EXCEPTION_VECTOR_PAGE_FAULT:
        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
        if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger &&
            AttachingCheckPageFaultsWithUserDebugger(CurrentProcessorIndex,
                                                     GuestRegs,
                                                     InterruptExit,
                                                     NULL,
                                                     ErrorCode)) {
        } else {
            IdtEmulationHandlePageFaults(CurrentProcessorIndex, InterruptExit, NULL, ErrorCode);
        }
        break;
    case EXCEPTION_VECTOR_DEBUG_BREAKPOINT:
        if (CurrentDebuggerState->ThreadOrProcessTracingDetails.DebugRegisterInterceptionState) {
            ThreadHandleThreadChange(CurrentProcessorIndex, GuestRegs);
        } else if (g_UserDebuggerState == TRUE &&
                   (g_IsWaitingForUserModeModuleEntrypointToBeCalled || g_IsWaitingForReturnAndRunFromPageFault)) {
            AttachingHandleEntrypointDebugBreak(CurrentProcessorIndex, GuestRegs);
        } else if (g_KernelDebuggerState == TRUE) {
            KdHandleDebugEventsWhenKernelDebuggerIsAttached(CurrentProcessorIndex, GuestRegs);
        } else if (UdCheckAndHandleBreakpointsAndDebugBreaks(CurrentProcessorIndex,
                                                             GuestRegs,
                                                             DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK,
                                                             NULL)) {
        } else {
            IdtEmulationReInjectInterruptOrException(InterruptExit);
        }
        break;
    case EXCEPTION_VECTOR_NMI:
        if (CurrentDebuggerState->EnableExternalInterruptsOnContinue ||
            CurrentDebuggerState->EnableExternalInterruptsOnContinueMtf ||
            CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf) {
        } else {
            IdtEmulationReInjectInterruptOrException(InterruptExit);
        }
        break;
    default:
        IdtEmulationReInjectInterruptOrException(InterruptExit);
        break;
    }
}

BOOLEAN
IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(_In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                                     _In_ UINT32                       CurrentProcessorIndex) {
    BOOLEAN                 FoundAPlaceForFutureInjection = FALSE;
    VIRTUAL_MACHINE_STATE * CurrentGuestState             = &g_GuestState[CurrentProcessorIndex];
    for (size_t i = 0; i < PENDING_INTERRUPTS_BUFFER_CAPACITY; i++) {
        if (CurrentGuestState->PendingExternalInterrupts[i] == NULL) {
            CurrentGuestState->PendingExternalInterrupts[i] = InterruptExit.AsUInt;
            FoundAPlaceForFutureInjection                   = TRUE;
            break;
        }
    }
    return FoundAPlaceForFutureInjection;
}

BOOLEAN
IdtEmulationCheckProcessOrThreadChange(_In_ UINT32                       CurrentProcessorIndex,
                                       _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                       _Inout_ PGUEST_REGS               GuestRegs) {
    VIRTUAL_MACHINE_STATE *     CurrentGuestState    = &g_GuestState[CurrentProcessorIndex];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = &g_GuestState[CurrentProcessorIndex].DebuggingState;
    if ((CurrentDebuggerState->ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange ||
         CurrentDebuggerState->ThreadOrProcessTracingDetails.InterceptClockInterruptsForProcessChange) &&
        ((CurrentProcessorIndex == 0 && InterruptExit.Vector == CLOCK_INTERRUPT) ||
         (CurrentProcessorIndex != 0 && InterruptExit.Vector == IPI_INTERRUPT))) {
        if (CurrentDebuggerState->ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange) {
            return ThreadHandleThreadChange(CurrentProcessorIndex, GuestRegs);
        } else {
            return ProcessHandleProcessChange(CurrentProcessorIndex, GuestRegs);
        }
    }
    return FALSE;
}

VOID
IdtEmulationHandleExternalInterrupt(_In_ UINT32                          CurrentProcessorIndex,
                                    _Inout_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                    _Inout_ PGUEST_REGS                  GuestRegs) {
    BOOLEAN                     Interruptible         = TRUE;
    VMX_INTERRUPTIBILITY_STATE  InterruptibilityState = {0};
    RFLAGS                      GuestRflags           = {0};
    ULONG                       ErrorCode             = 0;
    VIRTUAL_MACHINE_STATE *     CurrentGuestState     = &g_GuestState[CurrentProcessorIndex];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState  = &g_GuestState[CurrentProcessorIndex].DebuggingState;
    if (CurrentGuestState->WaitForImmediateVmexit &&
        InterruptExit.Vector == IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI) {
        HvSetExternalInterruptExiting(FALSE);
        CurrentGuestState->IncrementRip = FALSE;
        VmxMechanismHandleImmediateVmexit(CurrentProcessorIndex, GuestRegs);
        return;
    }
    IdtEmulationCheckProcessOrThreadChange(CurrentProcessorIndex, InterruptExit, GuestRegs);
    if ((CurrentDebuggerState->EnableExternalInterruptsOnContinue ||
         CurrentDebuggerState->EnableExternalInterruptsOnContinueMtf)) {
        IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(InterruptExit, CurrentProcessorIndex);
        CurrentGuestState->IncrementRip = FALSE;
    } else if (InterruptExit.Valid && InterruptExit.InterruptionType == INTERRUPT_TYPE_EXTERNAL_INTERRUPT) {
        __vmx_vmread(VMCS_GUEST_RFLAGS, &GuestRflags);
        __vmx_vmread(VMCS_GUEST_INTERRUPTIBILITY_STATE, &InterruptibilityState);
        Interruptible = GuestRflags.InterruptEnableFlag && !InterruptibilityState.BlockingByMovSs;
        if (Interruptible) {
            IdtEmulationReInjectInterruptOrException(InterruptExit);
        } else {
            IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(InterruptExit, CurrentProcessorIndex);
            HvSetInterruptWindowExiting(TRUE);
        }
        CurrentGuestState->IncrementRip = FALSE;
    } else {
        Interruptible = FALSE;
        LogError("Err, why we are here ? it's a vm-exit due to the external"
                 "interrupt and its type is not external interrupt? weird!");
    }
    DebuggerTriggerEvents(EXTERNAL_INTERRUPT_OCCURRED, GuestRegs, InterruptExit.Vector);
}

VOID
IdtEmulationHandleNmiWindowExiting(_In_ UINT32 CurrentProcessorIndex, _Inout_ PGUEST_REGS GuestRegs) {
    LogError("Why NMI-window exiting happens?");
}

VOID
IdtEmulationHandleInterruptWindowExiting(_In_ UINT32 CurrentProcessorIndex) {
    VMEXIT_INTERRUPT_INFORMATION InterruptExit     = {0};
    ULONG                        ErrorCode         = 0;
    VIRTUAL_MACHINE_STATE *      CurrentGuestState = &g_GuestState[CurrentProcessorIndex];
    for (size_t i = 0; i < PENDING_INTERRUPTS_BUFFER_CAPACITY; i++) {
        if (CurrentGuestState->PendingExternalInterrupts[i] != NULL) {
            InterruptExit.AsUInt                            = CurrentGuestState->PendingExternalInterrupts[i];
            CurrentGuestState->PendingExternalInterrupts[i] = NULL;
            break;
        }
    }
    if (InterruptExit.AsUInt == 0) {
        HvSetInterruptWindowExiting(FALSE);
    } else {
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, InterruptExit.AsUInt);
        if (InterruptExit.ErrorCodeValid) {
            __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
            __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
        }
    }
    CurrentGuestState->IncrementRip = FALSE;
}
