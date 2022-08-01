package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Vmcall.c.back

type (
	Vmcall interface {
		VmxHypervVmcallHandler() (ok bool) //col:23
		VmxHandleVmcallVmExit() (ok bool)  //col:50
		VmxVmcallHandler() (ok bool)       //col:381
		VmcallTest() (ok bool)             //col:395
	}
	vmcall struct{}
)

func NewVmcall() Vmcall { return &vmcall{} }

func (v *vmcall) VmxHypervVmcallHandler() (ok bool) { //col:23
	/*
	   VmxHypervVmcallHandler(PGUEST_REGS GuestRegs)

	   	{
	   	    UINT64                GuestRsp   = NULL;
	   	    HYPERCALL_INPUT_VALUE InputValue = {.Flags = GuestRegs->rcx};
	   	    switch (InputValue.Fields.CallCode)
	   	    {
	   	    case HvSwitchVirtualAddressSpace:
	   	    case HvFlushVirtualAddressSpace:
	   	    case HvFlushVirtualAddressList:
	   	    case HvCallFlushVirtualAddressSpaceEx:
	   	    case HvCallFlushVirtualAddressListEx:
	   	        VpidInvvpidAllContext();
	   	        break;
	   	    case HvCallFlushGuestPhysicalAddressSpace:
	   	    case HvCallFlushGuestPhysicalAddressList:
	   	        EptInveptSingleContext(g_EptState->EptPointer.AsUInt);
	   	        break;
	   	    }
	   	    GuestRsp = GuestRegs->rsp;
	   	    AsmHypervVmcall(GuestRegs);
	   	    GuestRegs->rsp = GuestRsp;
	   	    return STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (v *vmcall) VmxHandleVmcallVmExit() (ok bool) { //col:50
	/*
	   VmxHandleVmcallVmExit(UINT32      CoreIndex,

	   	PGUEST_REGS GuestRegs)

	   	{
	   	    UINT64  GuestRsp         = NULL;
	   	    BOOLEAN IsHyperdbgVmcall = FALSE;
	   	    if (g_TriggerEventForVmcalls)
	   	    {
	   	        DebuggerTriggerEvents(VMCALL_INSTRUCTION_EXECUTION, GuestRegs, NULL);
	   	    }
	   	    IsHyperdbgVmcall = (GuestRegs->r10 == 0x48564653 &&
	   	                        GuestRegs->r11 == 0x564d43414c4c &&
	   	                        GuestRegs->r12 == 0x4e4f485950455256);
	   	    if (IsHyperdbgVmcall)
	   	    {
	   	        GuestRegs->rax = VmxVmcallHandler(CoreIndex,
	   	                                          GuestRegs->rcx,
	   	                                          GuestRegs->rdx,
	   	                                          GuestRegs->r8,
	   	                                          GuestRegs->r9,
	   	                                          GuestRegs);
	   	    }
	   	    else
	   	    {
	   	        return VmxHypervVmcallHandler(GuestRegs);
	   	    }
	   	    return STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (v *vmcall) VmxVmcallHandler() (ok bool) { //col:381
	/*
	   VmxVmcallHandler(UINT32      CurrentCoreIndex,

	   	UINT64      VmcallNumber,
	   	UINT64      OptionalParam1,
	   	UINT64      OptionalParam2,
	   	UINT64      OptionalParam3,
	   	PGUEST_REGS GuestRegs)

	   	{
	   	    NTSTATUS VmcallStatus = STATUS_UNSUCCESSFUL;
	   	    BOOLEAN  HookResult   = FALSE;
	   	    BOOLEAN  UnsetExec    = FALSE;
	   	    BOOLEAN  UnsetWrite   = FALSE;
	   	    BOOLEAN  UnsetRead    = FALSE;
	   	    switch (VmcallNumber & 0xffffffff)
	   	    {
	   	    case VMCALL_TEST:
	   	    {
	   	        VmcallStatus = VmcallTest(OptionalParam1, OptionalParam2, OptionalParam3);
	   	        break;
	   	    }
	   	    case VMCALL_VMXOFF:
	   	    {
	   	        VmxVmxoff();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_CHANGE_PAGE_ATTRIB:
	   	    {
	   	        UINT32 AttributeMask = (UINT32)((VmcallNumber & 0xFFFFFFFF00000000LL) >> 32);
	   	        UnsetRead  = (AttributeMask & PAGE_ATTRIB_READ) ? TRUE : FALSE;
	   	        UnsetWrite = (AttributeMask & PAGE_ATTRIB_WRITE) ? TRUE : FALSE;
	   	        UnsetExec  = (AttributeMask & PAGE_ATTRIB_EXEC) ? TRUE : FALSE;
	   	        CR3_TYPE ProcCr3 = {.Flags = OptionalParam3};
	   	        HookResult = EptHookPerformPageHook2(OptionalParam1,
	   	                                             OptionalParam2,
	   	                                             ProcCr3,
	   	                                             UnsetRead,
	   	                                             UnsetWrite,
	   	                                             UnsetExec);
	   	        VmcallStatus = (HookResult == TRUE) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
	   	        break;
	   	    }
	   	    case VMCALL_INVEPT_SINGLE_CONTEXT:
	   	    {
	   	        EptInveptSingleContext(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_INVEPT_ALL_CONTEXTS:
	   	    {
	   	        EptInveptAllContexts();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNHOOK_ALL_PAGES:
	   	    {
	   	        EptHookRestoreAllHooksToOrginalEntry();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNHOOK_SINGLE_PAGE:
	   	    {
	   	        if (EptHookRestoreSingleHookToOrginalEntry(OptionalParam1))
	   	            VmcallStatus = STATUS_SUCCESS;
	   	        else
	   	            VmcallStatus = STATUS_UNSUCCESSFUL;
	   	        break;
	   	    }
	   	    case VMCALL_ENABLE_SYSCALL_HOOK_EFER:
	   	    {
	   	        SyscallHookConfigureEFER(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_SYSCALL_HOOK_EFER:
	   	    {
	   	        SyscallHookConfigureEFER(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_CHANGE_MSR_BITMAP_READ:
	   	    {
	   	        MsrHandlePerformMsrBitmapReadChange(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_CHANGE_MSR_BITMAP_WRITE:
	   	    {
	   	        MsrHandlePerformMsrBitmapWriteChange(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_RDTSC_EXITING:
	   	    {
	   	        HvSetRdtscExiting(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_RDPMC_EXITING:
	   	    {
	   	        HvSetPmcVmexit(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_EXCEPTION_BITMAP:
	   	    {
	   	        HvSetExceptionBitmap(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING:
	   	    {
	   	        HvSetMovDebugRegsExiting(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING:
	   	    {
	   	        HvSetExternalInterruptExiting(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_CHANGE_IO_BITMAP:
	   	    {
	   	        IoHandlePerformIoBitmapChange(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_HIDDEN_CC_BREAKPOINT:
	   	    {
	   	        CR3_TYPE ProcCr3 = {.Flags = OptionalParam2};
	   	        HookResult = EptHookPerformPageHook(OptionalParam1,
	   	                                            ProcCr3);
	   	        VmcallStatus = (HookResult == TRUE) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS:
	   	    {
	   	        ProtectedHvExternalInterruptExitingForDisablingInterruptCommands();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_RDTSC_EXITING:
	   	    {
	   	        HvSetRdtscExiting(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_RDPMC_EXITING:
	   	    {
	   	        HvSetPmcVmexit(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING:
	   	    {
	   	        HvSetMovDebugRegsExiting(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_RESET_MSR_BITMAP_READ:
	   	    {
	   	        MsrHandlePerformMsrBitmapReadReset();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_RESET_MSR_BITMAP_WRITE:
	   	    {
	   	        MsrHandlePerformMsrBitmapWriteReset();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS:
	   	    {
	   	        ProtectedHvResetExceptionBitmapToClearEvents();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_RESET_IO_BITMAP:
	   	    {
	   	        IoHandlePerformIoBitmapReset();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING:
	   	    {
	   	        HvSetMovControlRegsExiting(TRUE, OptionalParam1, OptionalParam2);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING:
	   	    {
	   	        HvSetMovControlRegsExiting(FALSE, OptionalParam1, OptionalParam2);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_ENABLE_MOV_TO_CR3_EXITING:
	   	    {
	   	        HvSetMovToCr3Vmexit(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_MOV_TO_CR3_EXITING:
	   	    {
	   	        HvSetMovToCr3Vmexit(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_EXCEPTION_BITMAP:
	   	    {
	   	        HvUnsetExceptionBitmap(OptionalParam1);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_VM_ENTRY_LOAD_DEBUG_CONTROLS:
	   	    {
	   	        HvSetLoadDebugControls(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_VM_ENTRY_LOAD_DEBUG_CONTROLS:
	   	    {
	   	        HvSetLoadDebugControls(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_VM_EXIT_SAVE_DEBUG_CONTROLS:
	   	    {
	   	        HvSetSaveDebugControls(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_VM_EXIT_SAVE_DEBUG_CONTROLS:
	   	    {
	   	        HvSetSaveDebugControls(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_VM_EXIT_HALT_SYSTEM:
	   	    {
	   	        KdHandleBreakpointAndDebugBreakpoints(CurrentCoreIndex,
	   	                                              GuestRegs,
	   	                                              DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER,
	   	                                              NULL);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SET_VM_EXIT_ON_NMIS:
	   	    {
	   	        HvSetNmiExiting(TRUE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_UNSET_VM_EXIT_ON_NMIS:
	   	    {
	   	        HvSetNmiExiting(FALSE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED:
	   	    {
	   	        KdSendCommandFinishedSignal(CurrentCoreIndex, GuestRegs);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SEND_MESSAGES_TO_DEBUGGER:
	   	    {
	   	        KdLoggingResponsePacketToDebugger(
	   	            OptionalParam1,
	   	            OptionalParam2,
	   	            OPERATION_LOG_INFO_MESSAGE);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER:
	   	    {
	   	        PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER DebuggeeBufferRequest =
	   	            OptionalParam1;
	   	        KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                   DebuggeeBufferRequest->RequestedAction,
	   	                                   (UINT64)DebuggeeBufferRequest + (SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER),
	   	                                   DebuggeeBufferRequest->LengthOfBuffer);
	   	        if (DebuggeeBufferRequest->PauseDebuggeeWhenSent)
	   	        {
	   	            KdHandleBreakpointAndDebugBreakpoints(CurrentCoreIndex,
	   	                                                  GuestRegs,
	   	                                                  DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM,
	   	                                                  NULL);
	   	        }
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT:
	   	    {
	   	        DEBUGGER_TRIGGERED_EVENT_DETAILS TriggeredEventDetail = {0};
	   	        TriggeredEventDetail.Context = OptionalParam1;
	   	        TriggeredEventDetail.Tag     = OptionalParam2;
	   	        KdHandleBreakpointAndDebugBreakpoints(CurrentCoreIndex,
	   	                                              OptionalParam3,
	   	                                              DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
	   	                                              &TriggeredEventDetail);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS:
	   	    {
	   	        ProtectedHvDisableRdtscExitingForDisablingTscCommands();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS:
	   	    {
	   	        ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands();
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    case VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS:
	   	    {
	   	        ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands(OptionalParam1, OptionalParam2);
	   	        VmcallStatus = STATUS_SUCCESS;
	   	        break;
	   	    }
	   	    default:
	   	    {
	   	        LogError("Err, unsupported VMCALL");
	   	        VmcallStatus = STATUS_UNSUCCESSFUL;
	   	        break;
	   	    }
	   	    }
	   	    return VmcallStatus;
	   	}
	*/
	return true
}

func (v *vmcall) VmcallTest() (ok bool) { //col:395
	/*
	   VmcallTest(_In_ UINT64 Param1,

	   	_In_ UINT64 Param2,
	   	_In_ UINT64 Param3)

	   	{
	   	                 Param1,
	   	                 Param2,
	   	                 Param3);
	   	    LogSendBuffer(OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED,
	   	                  "$",
	   	                  1,
	   	                  TRUE);
	   	    return STATUS_SUCCESS;
	   	}
	*/
	return true
}

