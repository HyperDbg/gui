#include "pch.h"
typedef enum _VMX_EXECUTION_MODE {
    VmxExecutionModeRoot,
    VmxExecutionModeNonRoot,
} VMX_EXECUTION_MODE;
_Must_inspect_result_
inline static VMX_EXECUTION_MODE
GetCurrentVmxExecutionMode() {
    ULONG                   CurrentCore    = KeGetCurrentProcessorIndex();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
    return CurrentVmState->IsOnVmxRootMode ? VmxExecutionModeRoot : VmxExecutionModeNonRoot;
}

_Must_inspect_result_
_Success_(return == TRUE)
static BOOLEAN
EptHookFindByPhysAddress(_In_ UINT64                        PhysicalBaseAddress,
                         _Out_opt_ EPT_HOOKED_PAGE_DETAIL * HookedEntry) {
    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity) {
        if (CurrEntity->PhysicalBaseAddress == PhysicalBaseAddress) {
            HookedEntry = CurrEntity;
            return TRUE;
        }
    }
    HookedEntry = NULL;
    return FALSE;
}

static UINT64
EptHookCalcBreakpointOffset(_In_ PVOID                    TargetAddress,
                            _In_ EPT_HOOKED_PAGE_DETAIL * HookedEntry) {
    UINT64 TargetAddressInFakePageContent;
    UINT64 PageOffset;
    TargetAddressInFakePageContent = &HookedEntry->FakePageContents;
    TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
    PageOffset                     = PAGE_OFFSET(TargetAddress);
    TargetAddressInFakePageContent = TargetAddressInFakePageContent + PageOffset;
    return TargetAddressInFakePageContent;
}

static BOOLEAN
EptHookCreateHookPage(_In_ PVOID    TargetAddress,
                      _In_ CR3_TYPE ProcessCr3) {
    EPT_PML1_ENTRY          ChangedEntry;
    INVEPT_DESCRIPTOR       Descriptor;
    SIZE_T                  PhysicalBaseAddress;
    PVOID                   VirtualTarget;
    PVOID                   TargetBuffer;
    UINT64                  TargetAddressInFakePageContent;
    UINT64                  PageOffset;
    PEPT_PML1_ENTRY         TargetPage;
    PEPT_HOOKED_PAGE_DETAIL HookedPage;
    CR3_TYPE                Cr3OfCurrentProcess;
    BYTE                    OriginalByte;
    BOOLEAN                 HookedEntryFound = FALSE;
    ULONG                   CurrentCore      = KeGetCurrentProcessorIndex();
    VIRTUAL_MACHINE_STATE * CurrentVmState   = &g_GuestState[CurrentCore];
    if (CurrentVmState->IsOnVmxRootMode && !CurrentVmState->HasLaunched) {
        return FALSE;
    }
    VirtualTarget       = PAGE_ALIGN(TargetAddress);
    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
    if (!PhysicalBaseAddress) {
        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
        return FALSE;
    }
    TargetBuffer = PoolManagerRequestPool(SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
    if (!TargetBuffer) {
        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
        return FALSE;
    }
    if (!EptSplitLargePage(g_EptState->EptPageTable, TargetBuffer, PhysicalBaseAddress, CurrentCore)) {
        PoolManagerFreePool(TargetBuffer);
        LogDebugInfo("Err, could not split page for the address : 0x%llx", PhysicalBaseAddress);
        DebuggerSetLastError(DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
        return FALSE;
    }
    TargetPage = EptGetPml1Entry(g_EptState->EptPageTable, PhysicalBaseAddress);
    if (!TargetPage) {
        PoolManagerFreePool(TargetBuffer);
        DebuggerSetLastError(DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
        return FALSE;
    }
    ChangedEntry = *TargetPage;
    HookedPage   = PoolManagerRequestPool(TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
    if (!HookedPage) {
        PoolManagerFreePool(TargetBuffer);
        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
        return FALSE;
    }
    HookedPage->IsHiddenBreakpoint                    = TRUE;
    HookedPage->VirtualAddress                        = TargetAddress;
    HookedPage->PhysicalBaseAddress                   = PhysicalBaseAddress;
    HookedPage->PhysicalBaseAddressOfFakePageContents = (SIZE_T)VirtualAddressToPhysicalAddress(&HookedPage->FakePageContents[0]) / PAGE_SIZE;
    HookedPage->EntryAddress                          = TargetPage;
    HookedPage->OriginalEntry                         = *TargetPage;
    HookedPage->IsExecutionHook                       = TRUE;
    HookedPage->BreakpointAddresses[0]                = TargetAddress;
    HookedPage->CountOfBreakpoints                    = 1;
    ChangedEntry.ReadAccess                           = 0;
    ChangedEntry.WriteAccess                          = 0;
    ChangedEntry.ExecuteAccess                        = 1;
    ChangedEntry.PageFrameNumber                      = HookedPage->PhysicalBaseAddressOfFakePageContents;
    TargetAddressInFakePageContent                    = &HookedPage->FakePageContents;
    TargetAddressInFakePageContent                    = PAGE_ALIGN(TargetAddressInFakePageContent);
    PageOffset                                        = PAGE_OFFSET(TargetAddress);
    TargetAddressInFakePageContent                    = TargetAddressInFakePageContent + PageOffset;
    Cr3OfCurrentProcess                               = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
    MemoryMapperReadMemorySafe(VirtualTarget, &HookedPage->FakePageContents, PAGE_SIZE);
    *(BYTE *)TargetAddressInFakePageContent = 0xcc;
    RestoreToPreviousProcess(Cr3OfCurrentProcess);
    HookedPage->ChangedEntry = ChangedEntry;
    InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
    if (!CurrentVmState->HasLaunched) {
        TargetPage->AsUInt = ChangedEntry.AsUInt;
    } else {
        EptSetPML1AndInvalidateTLB(TargetPage, ChangedEntry, InveptSingleContext);
    }
    return TRUE;
}

static BOOLEAN
EptHookUpdateHookPage(_In_ PVOID                       TargetAddress,
                      _Inout_ EPT_HOOKED_PAGE_DETAIL * HookedEntry) {
    UINT64 TargetAddressInFakePageContent;
    UINT64 PageOffset;
    BYTE   OriginalByte;
    if (HookedEntry == NULL)
        return FALSE;
    if (HookedEntry->CountOfBreakpoints >= MaximumHiddenBreakpointsOnPage) {
        DebuggerSetLastError(DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT);
        return FALSE;
    }
    TargetAddressInFakePageContent                                                   = &HookedEntry->FakePageContents;
    TargetAddressInFakePageContent                                                   = PAGE_ALIGN(TargetAddressInFakePageContent);
    PageOffset                                                                       = PAGE_OFFSET(TargetAddress);
    TargetAddressInFakePageContent                                                   = TargetAddressInFakePageContent + PageOffset;
    OriginalByte                                                                     = *(BYTE *)TargetAddressInFakePageContent;
    *(BYTE *)TargetAddressInFakePageContent                                          = 0xcc;
    HookedEntry->BreakpointAddresses[HookedEntry->CountOfBreakpoints]                = TargetAddress;
    HookedEntry->PreviousBytesOnBreakpointAddresses[HookedEntry->CountOfBreakpoints] = OriginalByte;
    HookedEntry->CountOfBreakpoints                                                  = HookedEntry->CountOfBreakpoints + 1;
    return TRUE;
}

PVOID
ExAllocatePoolWithTagHook(
    POOL_TYPE PoolType,
    SIZE_T    NumberOfBytes,
    ULONG     Tag) {
    LogInfo("ExAllocatePoolWithTag Called with : Tag = 0x%x , Number Of Bytes = 0x%x , Pool Type = 0x%x ", Tag, NumberOfBytes, PoolType);
    return ExAllocatePoolWithTagOrig(PoolType, NumberOfBytes, Tag);
}

BOOLEAN
EptHookPerformPageHook(PVOID TargetAddress, CR3_TYPE ProcessCr3) {
    EPT_PML1_ENTRY          ChangedEntry;
    INVEPT_DESCRIPTOR       Descriptor;
    SIZE_T                  PhysicalBaseAddress;
    PVOID                   VirtualTarget;
    PVOID                   TargetBuffer;
    UINT64                  TargetAddressInFakePageContent;
    UINT64                  PageOffset;
    PEPT_PML1_ENTRY         TargetPage;
    PEPT_HOOKED_PAGE_DETAIL HookedPage;
    ULONG                   CurrentCore;
    CR3_TYPE                Cr3OfCurrentProcess;
    BYTE                    OriginalByte;
    BOOLEAN                 HookedEntryFound = FALSE;
    CurrentCore                              = KeGetCurrentProcessorIndex();
    VIRTUAL_MACHINE_STATE * CurrentVmState   = &g_GuestState[CurrentCore];
    if (CurrentVmState->IsOnVmxRootMode && !CurrentVmState->HasLaunched) {
        return FALSE;
    }
    VirtualTarget       = PAGE_ALIGN(TargetAddress);
    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
    if (!PhysicalBaseAddress) {
        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
        return FALSE;
    }
    EPT_HOOKED_PAGE_DETAIL * HookedEntry = {0};
    if (EptHookFindByPhysAddress(PhysicalBaseAddress, HookedEntry) == TRUE && HookedEntry != NULL) {
        return EptHookUpdateHookPage(TargetAddress, HookedEntry);
    } else {
        return EptHookCreateHookPage(TargetAddress, ProcessCr3);
    }
}

BOOLEAN
EptHook(PVOID TargetAddress, UINT32 ProcessId) {
    ULONG                   CurrentCore    = KeGetCurrentProcessorIndex();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
    if (CurrentVmState->HasLaunched == FALSE)
        return FALSE;
    BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores();
    if (AsmVmxVmcall(VMCALL_SET_HIDDEN_CC_BREAKPOINT, TargetAddress, GetCr3FromProcessId(ProcessId).Flags, NULL) == STATUS_SUCCESS) {
        LogDebugInfo("Hidden breakpoint hook applied from VMX Root Mode");
        if (!CurrentVmState->IsOnVmxRootMode) {
            BroadcastNotifyAllToInvalidateEptAllCores();
        } else {
            LogError("Err, unable to notify all cores to invalidate their TLB caches as you called hook on vmx-root mode");
        }
        return TRUE;
    }
    return FALSE;
}

BOOLEAN
EptHookRestoreSingleHookToOrginalEntry(SIZE_T PhysicalAddress) {
    if (!g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode) {
        return FALSE;
    }
    EPT_HOOKED_PAGE_DETAIL * HookedEntry = {0};
    if (EptHookFindByPhysAddress(PAGE_ALIGN(PhysicalAddress), HookedEntry) == TRUE && HookedEntry != NULL) {
        EptSetPML1AndInvalidateTLB(HookedEntry->EntryAddress, HookedEntry->OriginalEntry, InveptSingleContext);
        return TRUE;
    }
    return FALSE;
}

VOID
EptHookRestoreAllHooksToOrginalEntry() {
    if (GetCurrentVmxExecutionMode() == VmxExecutionModeNonRoot)
        return;
    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity) {
        EptSetPML1AndInvalidateTLB(CurrEntity->EntryAddress, CurrEntity->OriginalEntry, InveptSingleContext);
    }
}

VOID
EptHookWriteAbsoluteJump(PCHAR TargetBuffer, SIZE_T TargetAddress) {
    TargetBuffer[0]               = 0xe8;
    TargetBuffer[1]               = 0x00;
    TargetBuffer[2]               = 0x00;
    TargetBuffer[3]               = 0x00;
    TargetBuffer[4]               = 0x00;
    TargetBuffer[5]               = 0x68;
    *((PUINT32)&TargetBuffer[6])  = (UINT32)TargetAddress;
    TargetBuffer[10]              = 0xC7;
    TargetBuffer[11]              = 0x44;
    TargetBuffer[12]              = 0x24;
    TargetBuffer[13]              = 0x04;
    *((PUINT32)&TargetBuffer[14]) = (UINT32)(TargetAddress >> 32);
    TargetBuffer[18]              = 0xC3;
}

VOID
EptHookWriteAbsoluteJump2(PCHAR TargetBuffer, SIZE_T TargetAddress) {
    TargetBuffer[0]              = 0x68;
    *((PUINT32)&TargetBuffer[1]) = (UINT32)TargetAddress;
    TargetBuffer[5]              = 0xC7;
    TargetBuffer[6]              = 0x44;
    TargetBuffer[7]              = 0x24;
    TargetBuffer[8]              = 0x04;
    *((PUINT32)&TargetBuffer[9]) = (UINT32)(TargetAddress >> 32);
    TargetBuffer[13]             = 0xC3;
}

BOOLEAN
EptHookInstructionMemory(PEPT_HOOKED_PAGE_DETAIL Hook,
                         CR3_TYPE                ProcessCr3,
                         PVOID                   TargetFunction,
                         PVOID                   TargetFunctionInSafeMemory,
                         PVOID                   HookFunction) {
    PHIDDEN_HOOKS_DETOUR_DETAILS DetourHookDetails;
    SIZE_T                       SizeOfHookedInstructions;
    SIZE_T                       OffsetIntoPage;
    CR3_TYPE                     Cr3OfCurrentProcess;
    OffsetIntoPage = ADDRMASK_EPT_PML1_OFFSET((SIZE_T)TargetFunction);
    if ((OffsetIntoPage + 19) > PAGE_SIZE - 1) {
        LogError("Err, function extends past a page boundary");
        return FALSE;
    }
    for (SizeOfHookedInstructions = 0;
         SizeOfHookedInstructions < 19;
         SizeOfHookedInstructions += ldisasm(((UINT64)TargetFunctionInSafeMemory + SizeOfHookedInstructions), TRUE)) {
    }
    Hook->Trampoline = PoolManagerRequestPool(EXEC_TRAMPOLINE, TRUE, MAX_EXEC_TRAMPOLINE_SIZE);
    if (!Hook->Trampoline) {
        LogError("Err, could not allocate trampoline function buffer");
        return FALSE;
    }
    Cr3OfCurrentProcess = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
    MemoryMapperReadMemorySafe(TargetFunction, Hook->Trampoline, SizeOfHookedInstructions);
    RestoreToPreviousProcess(Cr3OfCurrentProcess);
    EptHookWriteAbsoluteJump2(&Hook->Trampoline[SizeOfHookedInstructions], (SIZE_T)TargetFunction + SizeOfHookedInstructions);
    DetourHookDetails                        = PoolManagerRequestPool(DETOUR_HOOK_DETAILS, TRUE, sizeof(HIDDEN_HOOKS_DETOUR_DETAILS));
    DetourHookDetails->HookedFunctionAddress = TargetFunction;
    DetourHookDetails->ReturnAddress         = Hook->Trampoline;
    Hook->AddressOfEptHook2sDetourListEntry  = DetourHookDetails;
    InsertHeadList(&g_EptHook2sDetourListHead, &(DetourHookDetails->OtherHooksList));
    EptHookWriteAbsoluteJump(&Hook->FakePageContents[OffsetIntoPage], (SIZE_T)HookFunction);
    return TRUE;
}

BOOLEAN
EptHookPerformPageHook2(PVOID    TargetAddress,
                        PVOID    HookFunction,
                        CR3_TYPE ProcessCr3,
                        BOOLEAN  UnsetRead,
                        BOOLEAN  UnsetWrite,
                        BOOLEAN  UnsetExecute) {
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
    LogicalCoreIndex                    = KeGetCurrentProcessorIndex();
    if (g_GuestState[LogicalCoreIndex].IsOnVmxRootMode && !g_GuestState[LogicalCoreIndex].HasLaunched) {
        return FALSE;
    }
    VirtualTarget       = PAGE_ALIGN(TargetAddress);
    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(VirtualTarget, ProcessCr3);
    if (!PhysicalBaseAddress) {
        DebuggerSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
        return FALSE;
    }
    TempList = &g_EptState->HookedPagesList;
    while (&g_EptState->HookedPagesList != TempList->Flink) {
        TempList    = TempList->Flink;
        HookedEntry = CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
        if (HookedEntry->PhysicalBaseAddress == PhysicalBaseAddress) {
            DebuggerSetLastError(DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE);
            return FALSE;
        }
    }
    TargetBuffer = PoolManagerRequestPool(SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
    if (!TargetBuffer) {
        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
        return FALSE;
    }
    if (!EptSplitLargePage(g_EptState->EptPageTable, TargetBuffer, PhysicalBaseAddress, LogicalCoreIndex)) {
        PoolManagerFreePool(TargetBuffer);
        LogDebugInfo("Err, could not split page for the address : 0x%llx", PhysicalBaseAddress);
        DebuggerSetLastError(DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
        return FALSE;
    }
    TargetPage = EptGetPml1Entry(g_EptState->EptPageTable, PhysicalBaseAddress);
    if (!TargetPage) {
        PoolManagerFreePool(TargetBuffer);
        DebuggerSetLastError(DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
        return FALSE;
    }
    ChangedEntry = *TargetPage;
    if (UnsetRead)
        ChangedEntry.ReadAccess = 0;
    else
        ChangedEntry.ReadAccess = 1;
    if (UnsetWrite)
        ChangedEntry.WriteAccess = 0;
    else
        ChangedEntry.WriteAccess = 1;
    HookedPage = PoolManagerRequestPool(TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
    if (!HookedPage) {
        PoolManagerFreePool(TargetBuffer);
        DebuggerSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
        return FALSE;
    }
    HookedPage->VirtualAddress                        = TargetAddress;
    HookedPage->PhysicalBaseAddress                   = PhysicalBaseAddress;
    HookedPage->PhysicalBaseAddressOfFakePageContents = (SIZE_T)VirtualAddressToPhysicalAddress(&HookedPage->FakePageContents[0]) / PAGE_SIZE;
    HookedPage->EntryAddress                          = TargetPage;
    HookedPage->OriginalEntry                         = *TargetPage;
    if (UnsetExecute) {
        HookedPage->IsExecutionHook  = TRUE;
        ChangedEntry.ReadAccess      = 0;
        ChangedEntry.WriteAccess     = 0;
        ChangedEntry.ExecuteAccess   = 1;
        ChangedEntry.PageFrameNumber = HookedPage->PhysicalBaseAddressOfFakePageContents;
        Cr3OfCurrentProcess          = SwitchOnAnotherProcessMemoryLayoutByCr3(ProcessCr3);
        MemoryMapperReadMemorySafe(VirtualTarget, &HookedPage->FakePageContents, PAGE_SIZE);
        RestoreToPreviousProcess(Cr3OfCurrentProcess);
        TargetAddressInSafeMemory = &HookedPage->FakePageContents;
        TargetAddressInSafeMemory = PAGE_ALIGN(TargetAddressInSafeMemory);
        PageOffset                = PAGE_OFFSET(TargetAddress);
        TargetAddressInSafeMemory = TargetAddressInSafeMemory + PageOffset;
        if (!EptHookInstructionMemory(HookedPage, ProcessCr3, TargetAddress, TargetAddressInSafeMemory, HookFunction)) {
            PoolManagerFreePool(TargetBuffer);
            PoolManagerFreePool(HookedPage);
            DebuggerSetLastError(DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK);
            return FALSE;
        }
    }
    HookedPage->ChangedEntry = ChangedEntry;
    InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
    if (!g_GuestState[LogicalCoreIndex].HasLaunched) {
        TargetPage->AsUInt = ChangedEntry.AsUInt;
    } else {
        EptSetPML1AndInvalidateTLB(TargetPage, ChangedEntry, InveptSingleContext);
    }
    return TRUE;
}

BOOLEAN
EptHook2(PVOID TargetAddress, PVOID HookFunction, UINT32 ProcessId, BOOLEAN SetHookForRead, BOOLEAN SetHookForWrite, BOOLEAN SetHookForExec) {
    UINT32 PageHookMask = 0;
    ULONG  LogicalCoreIndex;
    if (SetHookForExec && !g_ExecuteOnlySupport) {
        return FALSE;
    }
    if (!SetHookForWrite && SetHookForRead) {
        return FALSE;
    }
    LogicalCoreIndex = KeGetCurrentProcessorIndex();
    if (SetHookForRead) {
        PageHookMask |= PAGE_ATTRIB_READ;
    }
    if (SetHookForWrite) {
        PageHookMask |= PAGE_ATTRIB_WRITE;
    }
    if (SetHookForExec) {
        PageHookMask |= PAGE_ATTRIB_EXEC;
    }
    if (PageHookMask == 0) {
        return FALSE;
    }
    if (g_GuestState[LogicalCoreIndex].HasLaunched) {
        UINT64 VmcallNumber = ((UINT64)PageHookMask) << 32 | VMCALL_CHANGE_PAGE_ATTRIB;
        if (AsmVmxVmcall(VmcallNumber, TargetAddress, HookFunction, GetCr3FromProcessId(ProcessId).Flags) == STATUS_SUCCESS) {
            if (!g_GuestState[LogicalCoreIndex].IsOnVmxRootMode) {
                BroadcastNotifyAllToInvalidateEptAllCores();
            } else {
                LogInfo("Err, unable to notify all cores to invalidate their TLB "
                        "caches as you called hook on vmx-root mode, however, the "
                        "hook is still works");
            }
            return TRUE;
        } else {
            return FALSE;
        }
    } else {
        if (EptHookPerformPageHook2(TargetAddress,
                                    HookFunction,
                                    GetCr3FromProcessId(ProcessId),
                                    SetHookForRead,
                                    SetHookForWrite,
                                    SetHookForExec) == TRUE) {
            LogInfo("Hook applied (VM has not launched)");
            return TRUE;
        }
    }
    LogWarning("Err, hook was not applied");
    return FALSE;
}

BOOLEAN
EptHookHandleHookedPage(PGUEST_REGS                          Regs,
                        EPT_HOOKED_PAGE_DETAIL *             HookedEntryDetails,
                        VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
                        SIZE_T                               PhysicalAddress) {
    UINT64                      GuestRip;
    UINT64                      ExactAccessedVirtualAddress;
    UINT64                      AlignedVirtualAddress;
    UINT64                      AlignedPhysicalAddress;
    EPT_HOOKS_TEMPORARY_CONTEXT TemporaryContext = {0};
    AlignedVirtualAddress                        = PAGE_ALIGN(HookedEntryDetails->VirtualAddress);
    AlignedPhysicalAddress                       = PAGE_ALIGN(PhysicalAddress);
    ExactAccessedVirtualAddress                  = AlignedVirtualAddress + PhysicalAddress - AlignedPhysicalAddress;
    TemporaryContext.PhysicalAddress             = PhysicalAddress;
    TemporaryContext.VirtualAddress              = ExactAccessedVirtualAddress;
    if (!ViolationQualification.EptExecutable && ViolationQualification.ExecuteAccess) {
        __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
        LogError("Err, Guest RIP : 0x%llx tries to execute the page at : 0x%llx", GuestRip, ExactAccessedVirtualAddress);
    } else if (!ViolationQualification.EptWriteable && ViolationQualification.WriteAccess) {
        DebuggerTriggerEvents(HIDDEN_HOOK_WRITE, Regs, &TemporaryContext);
        DebuggerTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE, Regs, &TemporaryContext);
    } else if (!ViolationQualification.EptReadable && ViolationQualification.ReadAccess) {
        DebuggerTriggerEvents(HIDDEN_HOOK_READ, Regs, &TemporaryContext);
        DebuggerTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE, Regs, &TemporaryContext);
    } else {
        return FALSE;
    }
    return TRUE;
}

BOOLEAN
EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(UINT64 Address) {
    LIST_FOR_EACH_LINK(g_EptHook2sDetourListHead, HIDDEN_HOOKS_DETOUR_DETAILS, OtherHooksList, CurrentHookedDetails) {
        if (CurrentHookedDetails->HookedFunctionAddress == Address) {
            RemoveEntryList(&CurrentHookedDetails->OtherHooksList);
            if (!PoolManagerFreePool(CurrentHookedDetails)) {
                LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");
            }
            return TRUE;
        }
    }
    return FALSE;
}

UINT32
EptHookGetCountOfEpthooks(BOOLEAN IsEptHook2) {
    UINT32 Count = 0;
    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, HookedEntry) {
        if (IsEptHook2) {
            if (HookedEntry->IsHiddenBreakpoint == FALSE) {
                Count++;
            }
        } else {
            if (HookedEntry->IsHiddenBreakpoint == TRUE) {
                Count++;
            }
        }
    }
    return Count;
}

BOOLEAN
EptHookUnHookSingleAddressDetours(PEPT_HOOKED_PAGE_DETAIL HookedEntry) {
    KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores, HookedEntry->PhysicalBaseAddress);
    EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(HookedEntry->VirtualAddress);
    RemoveEntryList(&HookedEntry->PageHookList);
    if (!PoolManagerFreePool(HookedEntry)) {
        LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");
        return FALSE;
    }
    return TRUE;
}

BOOLEAN
EptHookUnHookSingleAddressHiddenBreakpoint(PEPT_HOOKED_PAGE_DETAIL HookedEntry, UINT64 VirtualAddress) {
    UINT64 TargetAddressInFakePageContent;
    UINT64 PageOffset;
    UINT32 CountOfEntriesWithSameAddr = 0;
    for (size_t i = HookedEntry->CountOfBreakpoints; i-- > 0;) {
        if (HookedEntry->BreakpointAddresses[i] == VirtualAddress) {
            if (HookedEntry->CountOfBreakpoints == 1) {
                KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores, HookedEntry->PhysicalBaseAddress);
                RemoveEntryList(&HookedEntry->PageHookList);
                if (!PoolManagerFreePool(HookedEntry)) {
                    LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");
                }
                if (EptHookGetCountOfEpthooks(FALSE) == 0) {
                    BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores();
                }
                return TRUE;
            } else {
                TargetAddressInFakePageContent = &HookedEntry->FakePageContents;
                TargetAddressInFakePageContent = PAGE_ALIGN(TargetAddressInFakePageContent);
                PageOffset                     = PAGE_OFFSET(VirtualAddress);
                TargetAddressInFakePageContent = TargetAddressInFakePageContent + PageOffset;
                for (size_t j = 0; j < HookedEntry->CountOfBreakpoints; j++) {
                    if (HookedEntry->BreakpointAddresses[i] == VirtualAddress) {
                        CountOfEntriesWithSameAddr++;
                    }
                }
                if (CountOfEntriesWithSameAddr == 1) {
                    *(BYTE *)TargetAddressInFakePageContent = HookedEntry->PreviousBytesOnBreakpointAddresses[i];
                }
                HookedEntry->BreakpointAddresses[i]                = NULL;
                HookedEntry->PreviousBytesOnBreakpointAddresses[i] = 0x0;
                for (size_t j = i; j < MaximumHiddenBreakpointsOnPage - 1; j++) {
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

BOOLEAN
EptHookUnHookSingleAddress(UINT64 VirtualAddress, UINT64 PhysAddress, UINT32 ProcessId) {
    SIZE_T PhysicalAddress;
    if (GetCurrentVmxExecutionMode() == VmxExecutionModeRoot) {
        return FALSE;
    }
    if (ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES || ProcessId == 0) {
        ProcessId = PsGetCurrentProcessId();
    }
    if (PhysAddress == NULL) {
        PhysicalAddress = PAGE_ALIGN(VirtualAddressToPhysicalAddressByProcessId(VirtualAddress, ProcessId));
    } else {
        PhysicalAddress = PAGE_ALIGN(PhysAddress);
    }
    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity) {
        if (CurrEntity->IsHiddenBreakpoint) {
            return EptHookUnHookSingleAddressHiddenBreakpoint(CurrEntity, VirtualAddress);
        } else {
            if (CurrEntity->PhysicalBaseAddress == PhysicalAddress) {
                return EptHookUnHookSingleAddressDetours(CurrEntity);
            }
        }
    }
    return FALSE;
}

VOID
EptHookUnHookAll() {
    if (GetCurrentVmxExecutionMode() != VmxExecutionModeNonRoot) {
        return;
    }
    KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores, 0x0);
    LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL, PageHookList, CurrEntity) {
        if (!CurrEntity->IsHiddenBreakpoint) {
            EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(CurrEntity->VirtualAddress);
        }
        if (!PoolManagerFreePool(CurrEntity)) {
            LogError("Err, something goes wrong, the pool not found in the list of previously allocated pools by pool manager");
        }
    }
}
