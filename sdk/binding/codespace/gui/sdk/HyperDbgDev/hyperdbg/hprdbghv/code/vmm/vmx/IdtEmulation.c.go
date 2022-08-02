package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\IdtEmulation.c.back

type (
	IdtEmulation interface {
		IdtEmulationReInjectInterruptOrException() (ok bool) //col:114
	}
	idtEmulation struct{}
)

func NewIdtEmulation() IdtEmulation { return &idtEmulation{} }

func (i *idtEmulation) IdtEmulationReInjectInterruptOrException() (ok bool) { //col:114
	/*
	   IdtEmulationReInjectInterruptOrException(_In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit)

	   	{
	   	    ULONG ErrorCode = 0;
	   	    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, InterruptExit.AsUInt);
	   	    if (InterruptExit.ErrorCodeValid)
	   	    {
	   	        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
	   	        __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);

	   IdtEmulationHandlePageFaults(_In_ UINT32                       CurrentProcessorIndex,

	   	_In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
	   	_In_ UINT64                       Address,
	   	_In_ ULONG                        ErrorCode)

	   	{
	   	    PAGE_FAULT_ERROR_CODE   PageFaultCode     = {0};
	   	    VIRTUAL_MACHINE_STATE * CurrentGuestState = &g_GuestState[CurrentProcessorIndex];
	   	    __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &PageFaultCode);
	   	    if (Address == NULL)
	   	    {
	   	        UINT64 PageFaultAddress = 0;
	   	        __vmx_vmread(VMCS_EXIT_QUALIFICATION, &PageFaultAddress);
	   	        __writecr2(PageFaultAddress);
	   	        __writecr2(Address);
	   	    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, InterruptExit.AsUInt);
	   	    if (InterruptExit.ErrorCodeValid)
	   	    {
	   	        __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);

	   IdtEmulationHandleExceptionAndNmi(_In_ UINT32                          CurrentProcessorIndex,

	   	_Inout_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
	   	_Inout_ PGUEST_REGS                  GuestRegs)

	   	{
	   	    ULONG                       ErrorCode            = 0;
	   	    VIRTUAL_MACHINE_STATE *     CurrentGuestState    = &g_GuestState[CurrentProcessorIndex];
	   	    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = &g_GuestState[CurrentProcessorIndex].DebuggingState;
	   	    if (InterruptExit.InterruptionType == INTERRUPT_TYPE_NMI &&
	   	        InterruptExit.Vector == EXCEPTION_VECTOR_NMI)
	   	    {
	   	        if (!CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf &&
	   	            VmxBroadcastNmiHandler(CurrentProcessorIndex, GuestRegs, FALSE))
	   	        {
	   	            return;
	   	        }
	   	    }
	   	    if (CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf)
	   	    {
	   	        return;
	   	    }
	   	    DebuggerTriggerEvents(EXCEPTION_OCCURRED, GuestRegs, InterruptExit.Vector);
	   	    if (CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf)
	   	    {
	   	        return;
	   	    }
	   	    switch (InterruptExit.Vector)
	   	    {
	   	    case EXCEPTION_VECTOR_BREAKPOINT:
	   	        BreakpointHandleBpTraps(CurrentProcessorIndex, GuestRegs);
	   	        if (!SyscallHookHandleUD(GuestRegs, CurrentProcessorIndex))
	   	        {
	   	            EventInjectUndefinedOpcode(CurrentProcessorIndex);
	   	        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
	   	        if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger &&
	   	            AttachingCheckPageFaultsWithUserDebugger(CurrentProcessorIndex,
	   	                                                     GuestRegs,
	   	                                                     InterruptExit,
	   	                                                     NULL,
	   	                                                     ErrorCode))
	   	        {
	   	        }
	   	        else
	   	        {
	   	            IdtEmulationHandlePageFaults(CurrentProcessorIndex, InterruptExit, NULL, ErrorCode);
	   	        if (CurrentDebuggerState->ThreadOrProcessTracingDetails.DebugRegisterInterceptionState)
	   	        {
	   	            ThreadHandleThreadChange(CurrentProcessorIndex, GuestRegs);
	   	        else if (g_UserDebuggerState == TRUE &&
	   	                 (g_IsWaitingForUserModeModuleEntrypointToBeCalled || g_IsWaitingForReturnAndRunFromPageFault))
	   	        {
	   	            AttachingHandleEntrypointDebugBreak(CurrentProcessorIndex, GuestRegs);
	   	        else if (g_KernelDebuggerState == TRUE)
	   	        {
	   	            KdHandleDebugEventsWhenKernelDebuggerIsAttached(CurrentProcessorIndex, GuestRegs);
	   	        else if (UdCheckAndHandleBreakpointsAndDebugBreaks(CurrentProcessorIndex,
	   	                                                           GuestRegs,
	   	                                                           DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK,
	   	                                                           NULL))
	   	        {
	   	        }
	   	        else
	   	        {
	   	            IdtEmulationReInjectInterruptOrException(InterruptExit);
	   	        if (CurrentDebuggerState->EnableExternalInterruptsOnContinue ||
	   	            CurrentDebuggerState->EnableExternalInterruptsOnContinueMtf ||
	   	            CurrentDebuggerState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf)
	   	        {
	   	        }
	   	        else
	   	        {
	   	            IdtEmulationReInjectInterruptOrException(InterruptExit);
	   	        IdtEmulationReInjectInterruptOrException(InterruptExit);

	   IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(_In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,

	   	_In_ UINT32                       CurrentProcessorIndex)

	   	{
	   	    BOOLEAN                 FoundAPlaceForFutureInjection = FALSE;
	   	    VIRTUAL_MACHINE_STATE * CurrentGuestState             = &g_GuestState[CurrentProcessorIndex];
	   	    for (size_t i = 0; i < PENDING_INTERRUPTS_BUFFER_CAPACITY; i++)
	   	    {
	   	        if (CurrentGuestState->PendingExternalInterrupts[i] == NULL)
	   	        {
	   	            CurrentGuestState->PendingExternalInterrupts[i] = InterruptExit.AsUInt;
	   	            FoundAPlaceForFutureInjection                   = TRUE;
	   	            break;
	   	        }
	   	    }
	   	    return FoundAPlaceForFutureInjection;
	   	}
	*/
	return true
}

