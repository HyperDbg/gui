package ept_hook

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\features\hooks\ept-hook\EptHook.c.back

const (
	VmxExecutionModeRoot    = 1 //col:3
	VmxExecutionModeNonRoot = 2 //col:4
)

type (
	EptHook interface {
		GetCurrentVmxExecutionMode() (ok bool)        //col:19
		EptHookCalcBreakpointOffset() (ok bool)       //col:116
		ExAllocatePoolWithTagHook() (ok bool)         //col:432
		EptHookUnHookSingleAddressDetours() (ok bool) //col:486
	}
	eptHook struct{}
)

func NewEptHook() EptHook { return &eptHook{} }

func (e *eptHook) GetCurrentVmxExecutionMode() (ok bool) { //col:19
	/*
	   GetCurrentVmxExecutionMode()

	   	{
	   	    ULONG                   CurrentCore    = KeGetCurrentProcessorIndex();

	   _Success_(return == TRUE)
	   static BOOLEAN
	   EptHookFindByPhysAddress(_In_ UINT64 PhysicalBaseAddress,

	   	_Out_opt_ EPT_HOOKED_PAGE_DETAIL * HookedEntry)

	   	{
	   	    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity)
	   	    {
	   	        if (CurrEntity->PhysicalBaseAddress == PhysicalBaseAddress)
	   	        {
	   	            HookedEntry = CurrEntity;
	   	            return TRUE;
	   	        }
	   	    }
	   	    HookedEntry = NULL;
	   	    return FALSE;
	   	}
	*/
	return true
}

func (e *eptHook) EptHookCalcBreakpointOffset() (ok bool) { //col:116
	/*
	   EptHookCalcBreakpointOffset(_In_ PVOID TargetAddress,

	   	_In_ EPT_HOOKED_PAGE_DETAIL * HookedEntry)

	   	{
	   	    UINT64 TargetAddressInFakePageContent;
	   	    UINT64 PageOffset;
	   	    TargetAddressInFakePageContent = &HookedEntry->FakePageContents;
	   	    TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
	   	    PageOffset                     = PAGE_OFFSET(TargetAddress);

	   EptHookCreateHookPage(_In_ PVOID    TargetAddress,

	   	_In_ CR3_TYPE ProcessCr3)

	   	{
	   	    EPT_PML1_ENTRY          ChangedEntry;
	   	    INVEPT_DESCRIPTOR       Descriptor;
	   	    SIZE_T                  PhysicalBaseAddress;
	   	    PVOID                   VirtualTarget;
	   	    PVOID                   TargetBuffer;
	   	    UINT64                  TargetAddressInFakePageContent;
	   	    UINT64                  PageOffset;
	   	    PEPT_PML1_ENTRY         TargetPage;
	   	    PEPT_HOOKED_PAGE_DETAIL HookedPage;
	   	    CR3_TYPE Cr3OfCurrentProcess;
	   	    BYTE     OriginalByte;
	   	    BOOLEAN  HookedEntryFound = FALSE;
	   	    ULONG                   CurrentCore    = KeGetCurrentProcessorIndex();
	   	    if (CurrentVmState->IsOnVmxRootMode && !CurrentVmState->HasLaunched)
	   	    {
	   	        return FALSE;
	   	    }
	   	    VirtualTarget = PAGE_ALIGN(TargetAddress);
	   	    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
	   	    if (!PhysicalBaseAddress)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
	   	    TargetBuffer = PoolManagerRequestPool(SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
	   	    if (!TargetBuffer)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
	   	    if (!EptSplitLargePage(g_EptState->EptPageTable, TargetBuffer, PhysicalBaseAddress, CurrentCore))
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        LogDebugInfo("Err, could not split page for the address : 0x%llx", PhysicalBaseAddress);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
	   	    TargetPage = EptGetPml1Entry(g_EptState->EptPageTable, PhysicalBaseAddress);
	   	    if (!TargetPage)
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
	   	    HookedPage = PoolManagerRequestPool(TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
	   	    if (!HookedPage)
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
	   	    HookedPage->PhysicalBaseAddressOfFakePageContents = (SIZE_T)VirtualAddressToPhysicalAddress(&HookedPage->FakePageContents[0]) / PAGE_SIZE;
	   	    HookedPage->EntryAddress = TargetPage;
	   	    HookedPage->OriginalEntry = *TargetPage;
	   	    HookedPage->IsExecutionHook = TRUE;
	   	    HookedPage->BreakpointAddresses[0] = TargetAddress;
	   	    HookedPage->CountOfBreakpoints = 1;
	   	    ChangedEntry.ReadAccess    = 0;
	   	    ChangedEntry.WriteAccess   = 0;
	   	    ChangedEntry.ExecuteAccess = 1;
	   	    ChangedEntry.PageFrameNumber = HookedPage->PhysicalBaseAddressOfFakePageContents;
	   	    TargetAddressInFakePageContent = &HookedPage->FakePageContents;
	   	    TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
	   	    PageOffset                     = PAGE_OFFSET(TargetAddress);
	   	    Cr3OfCurrentProcess = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
	   	    MemoryMapperReadMemorySafe(VirtualTarget, &HookedPage->FakePageContents, PAGE_SIZE);
	   	    *(BYTE *)TargetAddressInFakePageContent = 0xcc;
	   	    RestoreToPreviousProcess(Cr3OfCurrentProcess);
	   	    InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
	   	    if (!CurrentVmState->HasLaunched)
	   	    {
	   	        TargetPage->AsUInt = ChangedEntry.AsUInt;
	   	    }
	   	    else
	   	    {
	   	        EptSetPML1AndInvalidateTLB(TargetPage, ChangedEntry, InveptSingleContext);

	   EptHookUpdateHookPage(_In_ PVOID TargetAddress,

	   	_Inout_ EPT_HOOKED_PAGE_DETAIL * HookedEntry)

	   	{
	   	    UINT64 TargetAddressInFakePageContent;
	   	    UINT64 PageOffset;
	   	    BYTE   OriginalByte;
	   	    if (HookedEntry == NULL)
	   	        return FALSE;
	   	    if (HookedEntry->CountOfBreakpoints >= MaximumHiddenBreakpointsOnPage)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT);
	   	    TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
	   	    PageOffset                     = PAGE_OFFSET(TargetAddress);
	   	    OriginalByte = *(BYTE *)TargetAddressInFakePageContent;
	   	    *(BYTE *)TargetAddressInFakePageContent = 0xcc;
	   	    HookedEntry->BreakpointAddresses[HookedEntry->CountOfBreakpoints] = TargetAddress;
	   	    HookedEntry->PreviousBytesOnBreakpointAddresses[HookedEntry->CountOfBreakpoints] = OriginalByte;
	   	    HookedEntry->CountOfBreakpoints = HookedEntry->CountOfBreakpoints + 1;
	   	    return TRUE;
	   	}
	*/
	return true
}

func (e *eptHook) ExAllocatePoolWithTagHook() (ok bool) { //col:432
	/*
	   ExAllocatePoolWithTagHook(

	   	POOL_TYPE PoolType,
	   	SIZE_T    NumberOfBytes,
	   	ULONG     Tag)

	   	{
	   	    LogInfo("ExAllocatePoolWithTag Called with : Tag = 0x%x , Number Of Bytes = 0x%x , Pool Type = 0x%x ", Tag, NumberOfBytes, PoolType);
	   	    return ExAllocatePoolWithTagOrig(PoolType, NumberOfBytes, Tag);

	   EptHookPerformPageHook(PVOID TargetAddress, CR3_TYPE ProcessCr3)

	   	{
	   	    EPT_PML1_ENTRY          ChangedEntry;
	   	    INVEPT_DESCRIPTOR       Descriptor;
	   	    SIZE_T                  PhysicalBaseAddress;
	   	    PVOID                   VirtualTarget;
	   	    PVOID                   TargetBuffer;
	   	    UINT64                  TargetAddressInFakePageContent;
	   	    UINT64                  PageOffset;
	   	    PEPT_PML1_ENTRY         TargetPage;
	   	    PEPT_HOOKED_PAGE_DETAIL HookedPage;
	   	    ULONG    CurrentCore;
	   	    CR3_TYPE Cr3OfCurrentProcess;
	   	    BYTE     OriginalByte;
	   	    BOOLEAN  HookedEntryFound = FALSE;
	   	    CurrentCore                            = KeGetCurrentProcessorIndex();
	   	    if (CurrentVmState->IsOnVmxRootMode && !CurrentVmState->HasLaunched)
	   	    {
	   	        return FALSE;
	   	    }
	   	    VirtualTarget = PAGE_ALIGN(TargetAddress);
	   	    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
	   	    if (!PhysicalBaseAddress)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
	   	    if (EptHookFindByPhysAddress(PhysicalBaseAddress, HookedEntry) == TRUE && HookedEntry != NULL)
	   	    {
	   	        return EptHookUpdateHookPage(TargetAddress, HookedEntry);
	   	        return EptHookCreateHookPage(TargetAddress, ProcessCr3);

	   EptHook(PVOID TargetAddress, UINT32 ProcessId)

	   	{
	   	    ULONG                   CurrentCore    = KeGetCurrentProcessorIndex();
	   	    if (CurrentVmState->HasLaunched == FALSE)
	   	        return FALSE;
	   	    BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores();
	   	    if (AsmVmxVmcall(VMCALL_SET_HIDDEN_CC_BREAKPOINT, TargetAddress, GetCr3FromProcessId(ProcessId).Flags, NULL) == STATUS_SUCCESS)
	   	    {
	   	        LogDebugInfo("Hidden breakpoint hook applied from VMX Root Mode");
	   	        if (!CurrentVmState->IsOnVmxRootMode)
	   	        {
	   	            BroadcastNotifyAllToInvalidateEptAllCores();
	   	            LogError("Err, unable to notify all cores to invalidate their TLB caches as you called hook on vmx-root mode");

	   EptHookRestoreSingleHookToOrginalEntry(SIZE_T PhysicalAddress)

	   	{
	   	    if (!g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode)
	   	    {
	   	        return FALSE;
	   	    }
	   	    EPT_HOOKED_PAGE_DETAIL * HookedEntry = {0};
	   	    if (EptHookFindByPhysAddress(PAGE_ALIGN(PhysicalAddress), HookedEntry) == TRUE && HookedEntry != NULL)
	   	    {
	   	        EptSetPML1AndInvalidateTLB(HookedEntry->EntryAddress, HookedEntry->OriginalEntry, InveptSingleContext);

	   EptHookRestoreAllHooksToOrginalEntry()

	   	{
	   	    if (GetCurrentVmxExecutionMode() == VmxExecutionModeNonRoot)
	   	        return;
	   	    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity)
	   	    {
	   	        EptSetPML1AndInvalidateTLB(CurrEntity->EntryAddress, CurrEntity->OriginalEntry, InveptSingleContext);

	   EptHookWriteAbsoluteJump(PCHAR TargetBuffer, SIZE_T TargetAddress)

	   	{
	   	    TargetBuffer[0] = 0xe8;
	   	    TargetBuffer[1] = 0x00;
	   	    TargetBuffer[2] = 0x00;
	   	    TargetBuffer[3] = 0x00;
	   	    TargetBuffer[4] = 0x00;
	   	    TargetBuffer[5] = 0x68;
	   	    *((PUINT32)&TargetBuffer[6]) = (UINT32)TargetAddress;
	   	    TargetBuffer[10] = 0xC7;
	   	    TargetBuffer[11] = 0x44;
	   	    TargetBuffer[12] = 0x24;
	   	    TargetBuffer[13] = 0x04;
	   	    *((PUINT32)&TargetBuffer[14]) = (UINT32)(TargetAddress >> 32);

	   EptHookWriteAbsoluteJump2(PCHAR TargetBuffer, SIZE_T TargetAddress)

	   	{
	   	    TargetBuffer[0] = 0x68;
	   	    *((PUINT32)&TargetBuffer[1]) = (UINT32)TargetAddress;
	   	    TargetBuffer[5] = 0xC7;
	   	    TargetBuffer[6] = 0x44;
	   	    TargetBuffer[7] = 0x24;
	   	    TargetBuffer[8] = 0x04;
	   	    *((PUINT32)&TargetBuffer[9]) = (UINT32)(TargetAddress >> 32);

	   EptHookInstructionMemory(PEPT_HOOKED_PAGE_DETAIL Hook,

	   	CR3_TYPE                ProcessCr3,
	   	PVOID                   TargetFunction,
	   	PVOID                   TargetFunctionInSafeMemory,
	   	PVOID                   HookFunction)

	   	{
	   	    PHIDDEN_HOOKS_DETOUR_DETAILS DetourHookDetails;
	   	    SIZE_T                       SizeOfHookedInstructions;
	   	    SIZE_T                       OffsetIntoPage;
	   	    CR3_TYPE                     Cr3OfCurrentProcess;
	   	    OffsetIntoPage = ADDRMASK_EPT_PML1_OFFSET((SIZE_T)TargetFunction);
	   	    if ((OffsetIntoPage + 19) > PAGE_SIZE - 1)
	   	    {
	   	        LogError("Err, function extends past a page boundary");
	   	    for (SizeOfHookedInstructions = 0;
	   	         SizeOfHookedInstructions < 19;
	   	         SizeOfHookedInstructions += ldisasm(((UINT64)TargetFunctionInSafeMemory + SizeOfHookedInstructions), TRUE))
	   	    {
	   	    }
	   	    Hook->Trampoline = PoolManagerRequestPool(EXEC_TRAMPOLINE, TRUE, MAX_EXEC_TRAMPOLINE_SIZE);
	   	    if (!Hook->Trampoline)
	   	    {
	   	        LogError("Err, could not allocate trampoline function buffer");
	   	    Cr3OfCurrentProcess = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
	   	    MemoryMapperReadMemorySafe(TargetFunction, Hook->Trampoline, SizeOfHookedInstructions);
	   	    RestoreToPreviousProcess(Cr3OfCurrentProcess);
	   	    EptHookWriteAbsoluteJump2(&Hook->Trampoline[SizeOfHookedInstructions], (SIZE_T)TargetFunction + SizeOfHookedInstructions);
	   	    DetourHookDetails                        = PoolManagerRequestPool(DETOUR_HOOK_DETAILS, TRUE, sizeof(HIDDEN_HOOKS_DETOUR_DETAILS));
	   	    InsertHeadList(&g_EptHook2sDetourListHead, &(DetourHookDetails->OtherHooksList));
	   	    EptHookWriteAbsoluteJump(&Hook->FakePageContents[OffsetIntoPage], (SIZE_T)HookFunction);

	   EptHookPerformPageHook2(PVOID    TargetAddress,

	   	PVOID    HookFunction,
	   	CR3_TYPE ProcessCr3,
	   	BOOLEAN  UnsetRead,
	   	BOOLEAN  UnsetWrite,
	   	BOOLEAN  UnsetExecute)

	   	{
	   	    EPT_PML1_ENTRY          ChangedEntry;
	   	    INVEPT_DESCRIPTOR       Descriptor;
	   	    SIZE_T                  PhysicalBaseAddress;
	   	    PVOID                   VirtualTarget;
	   	    PVOID                   TargetBuffer;
	   	    UINT64                  TargetAddressInSafeMemory;
	   	    UINT64                  PageOffset;
	   	    PEPT_PML1_ENTRY         TargetPage;
	   	    PEPT_HOOKED_PAGE_DETAIL HookedPage;
	   	    ULONG                   LogicalCoreIndex;
	   	    CR3_TYPE                Cr3OfCurrentProcess;
	   	    PLIST_ENTRY             TempList    = 0;
	   	    PEPT_HOOKED_PAGE_DETAIL HookedEntry = NULL;
	   	    LogicalCoreIndex = KeGetCurrentProcessorIndex();
	   	    if (g_GuestState[LogicalCoreIndex].IsOnVmxRootMode && !g_GuestState[LogicalCoreIndex].HasLaunched)
	   	    {
	   	        return FALSE;
	   	    }
	   	    VirtualTarget = PAGE_ALIGN(TargetAddress);
	   	    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
	   	    if (!PhysicalBaseAddress)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
	   	    while (&g_EptState->HookedPagesList != TempList->Flink)
	   	    {
	   	        TempList    = TempList->Flink;
	   	        HookedEntry = CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
	   	        if (HookedEntry->PhysicalBaseAddress == PhysicalBaseAddress)
	   	        {
	   	            DebuggerSetLastError(DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE);
	   	    TargetBuffer = PoolManagerRequestPool(SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
	   	    if (!TargetBuffer)
	   	    {
	   	        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
	   	    if (!EptSplitLargePage(g_EptState->EptPageTable, TargetBuffer, PhysicalBaseAddress, LogicalCoreIndex))
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        LogDebugInfo("Err, could not split page for the address : 0x%llx", PhysicalBaseAddress);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
	   	    TargetPage = EptGetPml1Entry(g_EptState->EptPageTable, PhysicalBaseAddress);
	   	    if (!TargetPage)
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
	   	    if (UnsetRead)
	   	        ChangedEntry.ReadAccess = 0;
	   	    else
	   	        ChangedEntry.ReadAccess = 1;
	   	    if (UnsetWrite)
	   	        ChangedEntry.WriteAccess = 0;
	   	    else
	   	        ChangedEntry.WriteAccess = 1;
	   	    HookedPage = PoolManagerRequestPool(TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
	   	    if (!HookedPage)
	   	    {
	   	        PoolManagerFreePool(TargetBuffer);
	   	        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
	   	    HookedPage->PhysicalBaseAddressOfFakePageContents = (SIZE_T)VirtualAddressToPhysicalAddress(&HookedPage->FakePageContents[0]) / PAGE_SIZE;
	   	    HookedPage->EntryAddress = TargetPage;
	   	    HookedPage->OriginalEntry = *TargetPage;
	   	    if (UnsetExecute)
	   	    {
	   	        HookedPage->IsExecutionHook = TRUE;
	   	        ChangedEntry.ReadAccess    = 0;
	   	        ChangedEntry.WriteAccess   = 0;
	   	        ChangedEntry.ExecuteAccess = 1;
	   	        ChangedEntry.PageFrameNumber = HookedPage->PhysicalBaseAddressOfFakePageContents;
	   	        Cr3OfCurrentProcess = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
	   	        MemoryMapperReadMemorySafe(VirtualTarget, &HookedPage->FakePageContents, PAGE_SIZE);
	   	        RestoreToPreviousProcess(Cr3OfCurrentProcess);
	   	        TargetAddressInSafeMemory = PAGE_ALIGN(TargetAddressInSafeMemory);
	   	        PageOffset                = PAGE_OFFSET(TargetAddress);
	   	        if (!EptHookInstructionMemory(HookedPage, ProcessCr3, TargetAddress, TargetAddressInSafeMemory, HookFunction))
	   	        {
	   	            PoolManagerFreePool(TargetBuffer);
	   	            PoolManagerFreePool(HookedPage);
	   	            DebuggerSetLastError(DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK);
	   	    InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
	   	    if (!g_GuestState[LogicalCoreIndex].HasLaunched)
	   	    {
	   	        TargetPage->AsUInt = ChangedEntry.AsUInt;
	   	    }
	   	    else
	   	    {
	   	        EptSetPML1AndInvalidateTLB(TargetPage, ChangedEntry, InveptSingleContext);

	   EptHook2(PVOID TargetAddress, PVOID HookFunction, UINT32 ProcessId, BOOLEAN SetHookForRead, BOOLEAN SetHookForWrite, BOOLEAN SetHookForExec)

	   	{
	   	    UINT32 PageHookMask = 0;
	   	    ULONG  LogicalCoreIndex;
	   	    if (SetHookForExec && !g_ExecuteOnlySupport)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!SetHookForWrite && SetHookForRead)
	   	    {
	   	        return FALSE;
	   	    }
	   	    LogicalCoreIndex = KeGetCurrentProcessorIndex();
	   	    if (SetHookForRead)
	   	    {
	   	        PageHookMask |= PAGE_ATTRIB_READ;
	   	    }
	   	    if (SetHookForWrite)
	   	    {
	   	        PageHookMask |= PAGE_ATTRIB_WRITE;
	   	    }
	   	    if (SetHookForExec)
	   	    {
	   	        PageHookMask |= PAGE_ATTRIB_EXEC;
	   	    }
	   	    if (PageHookMask == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (g_GuestState[LogicalCoreIndex].HasLaunched)
	   	    {
	   	        UINT64 VmcallNumber = ((UINT64)PageHookMask) << 32 | VMCALL_CHANGE_PAGE_ATTRIB;
	   	        if (AsmVmxVmcall(VmcallNumber, TargetAddress, HookFunction, GetCr3FromProcessId(ProcessId).Flags) == STATUS_SUCCESS)
	   	        {
	   	            if (!g_GuestState[LogicalCoreIndex].IsOnVmxRootMode)
	   	            {
	   	                BroadcastNotifyAllToInvalidateEptAllCores();
	   	                LogInfo("Err, unable to notify all cores to invalidate their TLB "
	   	                        "caches as you called hook on vmx-root mode, however, the "
	   	                        "hook is still works");
	   	        if (EptHookPerformPageHook2(TargetAddress,
	   	                                    HookFunction,
	   	                                    GetCr3FromProcessId(ProcessId),
	   	                                    SetHookForRead,
	   	                                    SetHookForWrite,
	   	                                    SetHookForExec) == TRUE)
	   	        {
	   	            LogInfo("Hook applied (VM has not launched)");
	   	    LogWarning("Err, hook was not applied");

	   EptHookHandleHookedPage(PGUEST_REGS                          Regs,

	   	EPT_HOOKED_PAGE_DETAIL *             HookedEntryDetails,
	   	VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
	   	SIZE_T                               PhysicalAddress)

	   	{
	   	    UINT64                      GuestRip;
	   	    UINT64                      ExactAccessedVirtualAddress;
	   	    UINT64                      AlignedVirtualAddress;
	   	    UINT64                      AlignedPhysicalAddress;
	   	    EPT_HOOKS_TEMPORARY_CONTEXT TemporaryContext = {0};
	   	    AlignedVirtualAddress  = PAGE_ALIGN(HookedEntryDetails->VirtualAddress);
	   	    AlignedPhysicalAddress = PAGE_ALIGN(PhysicalAddress);
	   	    if (!ViolationQualification.EptExecutable && ViolationQualification.ExecuteAccess)
	   	    {
	   	        __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
	   	        LogError("Err, Guest RIP : 0x%llx tries to execute the page at : 0x%llx", GuestRip, ExactAccessedVirtualAddress);
	   	    else if (!ViolationQualification.EptWriteable && ViolationQualification.WriteAccess)
	   	    {
	   	        DebuggerTriggerEvents(HIDDEN_HOOK_WRITE, Regs, &TemporaryContext);
	   	        DebuggerTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE, Regs, &TemporaryContext);
	   	    else if (!ViolationQualification.EptReadable && ViolationQualification.ReadAccess)
	   	    {
	   	        DebuggerTriggerEvents(HIDDEN_HOOK_READ, Regs, &TemporaryContext);
	   	        DebuggerTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE, Regs, &TemporaryContext);

	   EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(UINT64 Address)

	   	{
	   	    LIST_FOR_EACH_LINK(g_EptHook2sDetourListHead, HIDDEN_HOOKS_DETOUR_DETAILS, OtherHooksList, CurrentHookedDetails)
	   	    {
	   	        if (CurrentHookedDetails->HookedFunctionAddress == Address)
	   	        {
	   	            RemoveEntryList(&CurrentHookedDetails->OtherHooksList);
	   	            if (!PoolManagerFreePool(CurrentHookedDetails))
	   	            {
	   	                LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");

	   EptHookGetCountOfEpthooks(BOOLEAN IsEptHook2)

	   	{
	   	    UINT32 Count = 0;
	   	    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, HookedEntry)
	   	    {
	   	        if (IsEptHook2)
	   	        {
	   	            if (HookedEntry->IsHiddenBreakpoint == FALSE)
	   	            {
	   	                Count++;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            if (HookedEntry->IsHiddenBreakpoint == TRUE)
	   	            {
	   	                Count++;
	   	            }
	   	        }
	   	    }
	   	    return Count;
	   	}
	*/
	return true
}

func (e *eptHook) EptHookUnHookSingleAddressDetours() (ok bool) { //col:486
	/*
	   EptHookUnHookSingleAddressDetours(PEPT_HOOKED_PAGE_DETAIL HookedEntry)

	   	{
	   	    KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores, HookedEntry->PhysicalBaseAddress);
	   	    EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(HookedEntry->VirtualAddress);
	   	    RemoveEntryList(&HookedEntry->PageHookList);
	   	    if (!PoolManagerFreePool(HookedEntry))
	   	    {
	   	        LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");

	   EptHookUnHookSingleAddressHiddenBreakpoint(PEPT_HOOKED_PAGE_DETAIL HookedEntry, UINT64 VirtualAddress)

	   	{
	   	    UINT64 TargetAddressInFakePageContent;
	   	    UINT64 PageOffset;
	   	    UINT32 CountOfEntriesWithSameAddr = 0;
	   	    for (size_t i = HookedEntry->CountOfBreakpoints; i-- > 0;)
	   	    {
	   	        if (HookedEntry->BreakpointAddresses[i] == VirtualAddress)
	   	        {
	   	            if (HookedEntry->CountOfBreakpoints == 1)
	   	            {
	   	                KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores, HookedEntry->PhysicalBaseAddress);
	   	                RemoveEntryList(&HookedEntry->PageHookList);
	   	                if (!PoolManagerFreePool(HookedEntry))
	   	                {
	   	                    LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");
	   	                if (EptHookGetCountOfEpthooks(FALSE) == 0)
	   	                {
	   	                    BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores();
	   	                TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
	   	                PageOffset                     = PAGE_OFFSET(VirtualAddress);
	   	                for (size_t j = 0; j < HookedEntry->CountOfBreakpoints; j++)
	   	                {
	   	                    if (HookedEntry->BreakpointAddresses[i] == VirtualAddress)
	   	                    {
	   	                        CountOfEntriesWithSameAddr++;
	   	                    }
	   	                }
	   	                if (CountOfEntriesWithSameAddr == 1)
	   	                {
	   	                    *(BYTE *)TargetAddressInFakePageContent = HookedEntry->PreviousBytesOnBreakpointAddresses[i];
	   	                }
	   	                HookedEntry->BreakpointAddresses[i]                = NULL;
	   	                HookedEntry->PreviousBytesOnBreakpointAddresses[i] = 0x0;
	   	                for (size_t j = i; j < MaximumHiddenBreakpointsOnPage - 1; j++)
	   	                {
	   	                    HookedEntry->BreakpointAddresses[j]                = HookedEntry->BreakpointAddresses[j + 1];
	   	                    HookedEntry->PreviousBytesOnBreakpointAddresses[j] = HookedEntry->PreviousBytesOnBreakpointAddresses[j + 1];
	   	                }
	   	                HookedEntry->CountOfBreakpoints = HookedEntry->CountOfBreakpoints - 1;
	   	                return TRUE;
	   	            }
	   	        }
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

