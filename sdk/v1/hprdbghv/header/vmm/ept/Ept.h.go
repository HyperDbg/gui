package ept
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\ept\Ept.h.back

type  _VMM_EPT_PAGE_TABLE struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:11
PML4[VMM_EPT_PML4E_COUNT] EPT_PML4_POINTER //col:12
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:13
PML3[VMM_EPT_PML3E_COUNT] EPT_PML3_POINTER //col:14
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:15
PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT] EPT_PML2_ENTRY //col:16
}


type  _MTRR_RANGE_DESCRIPTOR struct{
PhysicalBaseAddress int64 //col:17
PhysicalEndAddress int64 //col:18
MemoryType uint8 //col:19
}


type  _EPT_STATE struct{
HookedPagesList *list.List //col:26
MemoryRanges[EPT_MTRR_RANGE_DESCRIPTOR_MAX] MTRR_RANGE_DESCRIPTOR //col:27
NumberOfEnabledMemoryRanges uint32 //col:28
EptPointer EPT_POINTER //col:29
EptPageTable PVMM_EPT_PAGE_TABLE //col:30
SecondaryEptPageTable PVMM_EPT_PAGE_TABLE //col:31
}


type  _VMM_EPT_DYNAMIC_SPLIT struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:35
PML1[VMM_EPT_PML1E_COUNT] EPT_PML1_ENTRY //col:36
Union union //col:37
Entry PEPT_PML2_ENTRY //col:39
Pointer PEPT_PML2_POINTER //col:40
}


type  _EPT_HOOKED_PAGE_DETAIL struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:56
FakePageContents[PAGE_SIZE] int8 //col:57
PageHookList *list.List //col:58
VirtualAddress uint64 //col:59
AddressOfEptHook2sDetourListEntry uint64 //col:60
PhysicalBaseAddress int64 //col:61
PhysicalBaseAddressOfFakePageContents int64 //col:62
EntryAddress PEPT_PML1_ENTRY //col:63
OriginalEntry EPT_PML1_ENTRY //col:64
ChangedEntry EPT_PML1_ENTRY //col:65
Trampoline PCHAR //col:66
IsExecutionHook bool //col:67
IsHiddenBreakpoint bool //col:68
BreakpointAddresses[MaximumHiddenBreakpointsOnPage] uint64 //col:69
PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage] int8 //col:70
CountOfBreakpoints uint64 //col:71
}




