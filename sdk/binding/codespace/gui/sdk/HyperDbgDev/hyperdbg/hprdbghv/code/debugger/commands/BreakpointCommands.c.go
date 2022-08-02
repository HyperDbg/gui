package commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\BreakpointCommands.c.back

type (
	BreakpointCommands interface {
		BreakpointCheckAndHandleEptHookBreakpoints() (ok bool) //col:148
		BreakpointGetEntryByAddress() (ok bool)                //col:163
	}
	breakpointCommands struct{}
)

func NewBreakpointCommands() BreakpointCommands { return &breakpointCommands{} }

func (b *breakpointCommands) BreakpointCheckAndHandleEptHookBreakpoints() (ok bool) { //col:148
	/*
	   BreakpointCheckAndHandleEptHookBreakpoints(UINT32 CurrentProcessorIndex, UINT64 GuestRip, PGUEST_REGS GuestRegs)

	   	{
	   	    PLIST_ENTRY             TempList           = 0;
	   	    BOOLEAN                 IsHandledByEptHook = FALSE;
	   	    VIRTUAL_MACHINE_STATE * CurrentVmState     = &g_GuestState[CurrentProcessorIndex];
	   	    TempList = &g_EptState->HookedPagesList;
	   	    while (&g_EptState->HookedPagesList != TempList->Flink)
	   	    {
	   	        TempList                            = TempList->Flink;
	   	        PEPT_HOOKED_PAGE_DETAIL HookedEntry = CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
	   	        if (HookedEntry->IsExecutionHook)
	   	        {
	   	            for (size_t i = 0; i < HookedEntry->CountOfBreakpoints; i++)
	   	            {
	   	                if (HookedEntry->BreakpointAddresses[i] == GuestRip)
	   	                {
	   	                    DebuggerTriggerEvents(HIDDEN_HOOK_EXEC_CC, GuestRegs, GuestRip);
	   	                    EptSetPML1AndInvalidateTLB(HookedEntry->EntryAddress, HookedEntry->OriginalEntry, InveptSingleContext);
	   	                    HvSetMonitorTrapFlag(TRUE);
	   	                    HvSetExternalInterruptExiting(TRUE);
	   	                    HvSetInterruptWindowExiting(FALSE);

	   BreakpointCheckAndHandleDebuggerDefinedBreakpoints(UINT32                  CurrentProcessorIndex,

	   	UINT64                  GuestRip,
	   	DEBUGGEE_PAUSING_REASON Reason,
	   	PGUEST_REGS             GuestRegs,
	   	PBOOLEAN                AvoidUnsetMtf)

	   	{
	   	    CR3_TYPE                         GuestCr3;
	   	    BOOLEAN                          IsHandledByBpRoutines = FALSE;
	   	    PLIST_ENTRY                      TempList              = 0;
	   	    UINT64                           GuestRipPhysical      = NULL;
	   	    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag         = {0};
	   	    RFLAGS                           Rflags                = {0};
	   	    ULONG                            LengthOfExitInstr     = 0;
	   	    BYTE                             InstrByte             = NULL;
	   	    *AvoidUnsetMtf                                         = FALSE;
	   	    VIRTUAL_MACHINE_STATE * CurrentVmState                 = &g_GuestState[CurrentProcessorIndex];
	   	    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	   	    GuestRipPhysical = VirtualAddressToPhysicalAddressByProcessCr3(GuestRip, GuestCr3);
	   	    while (&g_BreakpointsListHead != TempList->Flink)
	   	    {
	   	        TempList                                      = TempList->Flink;
	   	        PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	   	        if (CurrentBreakpointDesc->PhysAddress == GuestRipPhysical)
	   	        {
	   	            IsHandledByBpRoutines = TRUE;
	   	            MemoryMapperWriteMemorySafeByPhysicalAddress(GuestRipPhysical,
	   	                                                         &CurrentBreakpointDesc->PreviousByte,
	   	                                                         sizeof(BYTE));
	   	            if (Reason == DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT)
	   	            {
	   	                ContextAndTag.Tag = CurrentBreakpointDesc->BreakpointId;
	   	            }
	   	            CurrentVmState->DebuggingState.InstructionLengthHint = CurrentBreakpointDesc->InstructionLength;
	   	            if ((CurrentBreakpointDesc->Pid == DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES || CurrentBreakpointDesc->Pid == PsGetCurrentProcessId()) &&
	   	                (CurrentBreakpointDesc->Tid == DEBUGGEE_BP_APPLY_TO_ALL_THREADS || CurrentBreakpointDesc->Tid == PsGetCurrentThreadId()) &&
	   	                (CurrentBreakpointDesc->Core == DEBUGGEE_BP_APPLY_TO_ALL_CORES || CurrentBreakpointDesc->Core == CurrentProcessorIndex))
	   	            {
	   	                KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
	   	                                                      GuestRegs,
	   	                                                      Reason,
	   	                                                      &ContextAndTag);
	   	            if (!CurrentBreakpointDesc->AvoidReApplyBreakpoint)
	   	            {
	   	                CurrentVmState->DebuggingState.SoftwareBreakpointState = CurrentBreakpointDesc;
	   	                HvSetMonitorTrapFlag(TRUE);
	   	                __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
	   	                if (Rflags.InterruptEnableFlag)
	   	                {
	   	                    Rflags.InterruptEnableFlag = FALSE;
	   	                    __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);

	   BreakpointHandleBpTraps(UINT32 CurrentProcessorIndex, PGUEST_REGS GuestRegs)

	   	{
	   	    UINT64                           GuestRip           = NULL;
	   	    BOOLEAN                          IsHandledByEptHook = FALSE;
	   	    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag      = {0};
	   	    BOOLEAN                          AvoidUnsetMtf;
	   	    VIRTUAL_MACHINE_STATE *          CurrentVmState = &g_GuestState[CurrentProcessorIndex];
	   	    __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
	   	    IsHandledByEptHook = BreakpointCheckAndHandleEptHookBreakpoints(CurrentProcessorIndex, GuestRip, GuestRegs);
	   	    if (!IsHandledByEptHook)
	   	    {
	   	        if (g_KernelDebuggerState)
	   	        {
	   	            if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(CurrentProcessorIndex,
	   	                                                                    GuestRip,
	   	                                                                    DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
	   	                                                                    GuestRegs,
	   	                                                                    &AvoidUnsetMtf))
	   	            {
	   	                ContextAndTag.Context = CurrentVmState->LastVmexitRip;
	   	                KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
	   	                                                      GuestRegs,
	   	                                                      DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
	   	                                                      &ContextAndTag);
	   	            EventInjectBreakpoint();

	   BreakpointWrite(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor)

	   	{
	   	    BYTE PreviousByte   = NULL;
	   	    BYTE BreakpointByte = 0xcc;
	   	    if (!CheckMemoryAccessSafety(BreakpointDescriptor->Address, sizeof(BYTE)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    MemoryMapperReadMemorySafeOnTargetProcess(BreakpointDescriptor->Address, &PreviousByte, sizeof(BYTE));
	   	    MemoryMapperWriteMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress,
	   	                                                 &BreakpointByte,
	   	                                                 sizeof(BYTE));

	   BreakpointClear(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor)

	   	{
	   	    BYTE TargetMem = NULL;
	   	    if (!CheckMemoryAccessSafety(BreakpointDescriptor->Address, sizeof(BYTE)))
	   	    {
	   	        MemoryMapperReadMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress, &TargetMem, sizeof(BYTE));
	   	        if (TargetMem != 0xcc)
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    MemoryMapperWriteMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress,
	   	                                                 &BreakpointDescriptor->PreviousByte,
	   	                                                 sizeof(BYTE));

	   BreakpointRemoveAllBreakpoints()

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    TempList = &g_BreakpointsListHead;
	   	    while (&g_BreakpointsListHead != TempList->Flink)
	   	    {
	   	        TempList                                      = TempList->Flink;
	   	        PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	   	        BreakpointClear(CurrentBreakpointDesc);
	   	        RemoveEntryList(&CurrentBreakpointDesc->BreakpointsList);
	   	        PoolManagerFreePool(CurrentBreakpointDesc);

	   BreakpointGetEntryByBreakpointId(UINT64 BreakpointId)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    TempList = &g_BreakpointsListHead;
	   	    while (&g_BreakpointsListHead != TempList->Flink)
	   	    {
	   	        TempList                                      = TempList->Flink;
	   	        PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	   	        if (CurrentBreakpointDesc->BreakpointId == BreakpointId)
	   	        {
	   	            return CurrentBreakpointDesc;
	   	        }
	   	    }
	   	    return NULL;
	   	}
	*/
	return true
}

func (b *breakpointCommands) BreakpointGetEntryByAddress() (ok bool) { //col:163
	/*
	   BreakpointGetEntryByAddress(UINT64 Address)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    TempList = &g_BreakpointsListHead;
	   	    while (&g_BreakpointsListHead != TempList->Flink)
	   	    {
	   	        TempList                                      = TempList->Flink;
	   	        PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	   	        if (CurrentBreakpointDesc->Address == Address)
	   	        {
	   	            return CurrentBreakpointDesc;
	   	        }
	   	    }
	   	    return NULL;
	   	}
	*/
	return true
}

