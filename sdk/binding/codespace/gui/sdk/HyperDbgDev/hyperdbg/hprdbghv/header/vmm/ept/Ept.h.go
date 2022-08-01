package ept
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\ept\Ept.h.back

const(
MaximumHiddenBreakpointsOnPage = 40 //col:1
PAGE_ATTRIB_READ =  0x2 //col:2
PAGE_ATTRIB_WRITE = 0x4 //col:3
PAGE_ATTRIB_EXEC =  0x8 //col:4
VMM_EPT_PML4E_COUNT = 512 //col:5
VMM_EPT_PML3E_COUNT = 512 //col:6
VMM_EPT_PML2E_COUNT = 512 //col:7
VMM_EPT_PML1E_COUNT = 512 //col:8
SIZE_2_MB = ((SIZE_T)(512 * PAGE_SIZE)) //col:9
ADDRMASK_EPT_PML1_OFFSET(_VAR_) = (_VAR_ & 0xFFFULL) //col:10
ADDRMASK_EPT_PML1_INDEX(_VAR_) = ((_VAR_ & 0x1FF000ULL) >> 12) //col:11
ADDRMASK_EPT_PML2_INDEX(_VAR_) = ((_VAR_ & 0x3FE00000ULL) >> 21) //col:12
ADDRMASK_EPT_PML3_INDEX(_VAR_) = ((_VAR_ & 0x7FC0000000ULL) >> 30) //col:13
ADDRMASK_EPT_PML4_INDEX(_VAR_) = ((_VAR_ & 0xFF8000000000ULL) >> 39) //col:14
EPT_MTRR_RANGE_DESCRIPTOR_MAX = 0x9 //col:15
)

type VMM_EPT_PAGE_TABLE struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:3
PML4[VMM_EPT_PML4E_COUNT] EPT_PML4_POINTER //col:4
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:5
PML3[VMM_EPT_PML3E_COUNT] EPT_PML3_POINTER //col:6
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:7
PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT] EPT_PML2_ENTRY //col:8
}


type MTRR_RANGE_DESCRIPTOR struct{
PhysicalBaseAddress SIZE_T //col:12
PhysicalEndAddress SIZE_T //col:13
MemoryType uint8 //col:14
}


type EPT_STATE struct{
HookedPagesList *list.List //col:18
MemoryRanges[EPT_MTRR_RANGE_DESCRIPTOR_MAX] MTRR_RANGE_DESCRIPTOR //col:19
NumberOfEnabledMemoryRanges uint32 //col:20
EptPointer EPT_POINTER //col:21
EptPageTable PVMM_EPT_PAGE_TABLE //col:22
SecondaryEptPageTable PVMM_EPT_PAGE_TABLE //col:23
}


type VMM_EPT_DYNAMIC_SPLIT struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:27
PML1[VMM_EPT_PML1E_COUNT] EPT_PML1_ENTRY //col:28
Union union //col:29
Entry PEPT_PML2_ENTRY //col:31
Pointer PEPT_PML2_POINTER //col:32
}


type EPT_HOOKED_PAGE_DETAIL struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE) //col:38
FakePageContents[PAGE_SIZE] int8 //col:39
PageHookList *list.List //col:40
VirtualAddress uint64 //col:41
AddressOfEptHook2sDetourListEntry uint64 //col:42
PhysicalBaseAddress SIZE_T //col:43
PhysicalBaseAddressOfFakePageContents SIZE_T //col:44
EntryAddress PEPT_PML1_ENTRY //col:45
OriginalEntry EPT_PML1_ENTRY //col:46
ChangedEntry EPT_PML1_ENTRY //col:47
Trampoline PCHAR //col:48
IsExecutionHook bool //col:49
IsHiddenBreakpoint bool //col:50
BreakpointAddresses[MaximumHiddenBreakpointsOnPage] uint64 //col:51
PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage] int8 //col:52
CountOfBreakpoints uint64 //col:53
}




