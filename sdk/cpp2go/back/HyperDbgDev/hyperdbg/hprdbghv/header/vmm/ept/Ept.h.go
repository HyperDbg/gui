package ept
//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\ept\Ept.h.back

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
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE)
PML4[VMM_EPT_PML4E_COUNT] EPT_PML4_POINTER
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE)
PML3[VMM_EPT_PML3E_COUNT] EPT_PML3_POINTER
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE)
PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT] EPT_PML2_ENTRY
}


type MTRR_RANGE_DESCRIPTOR struct{
PhysicalBaseAddress SIZE_T
PhysicalEndAddress SIZE_T
MemoryType UCHAR
}


type EPT_STATE struct{
HookedPagesList *list.List
MemoryRanges[EPT_MTRR_RANGE_DESCRIPTOR_MAX] MTRR_RANGE_DESCRIPTOR
NumberOfEnabledMemoryRanges ULONG
EptPointer EPT_POINTER
EptPageTable PVMM_EPT_PAGE_TABLE
SecondaryEptPageTable PVMM_EPT_PAGE_TABLE
}


type VMM_EPT_DYNAMIC_SPLIT struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE)
PML1[VMM_EPT_PML1E_COUNT] EPT_PML1_ENTRY
Union union
Entry PEPT_PML2_ENTRY
Pointer PEPT_PML2_POINTER
}


type EPT_HOOKED_PAGE_DETAIL struct{
DeclspecAlign(PageSize) DECLSPEC_ALIGN(PAGE_SIZE)
FakePageContents[PAGE_SIZE] CHAR
PageHookList *list.List
VirtualAddress uint64
AddressOfEptHook2sDetourListEntry uint64
PhysicalBaseAddress SIZE_T
PhysicalBaseAddressOfFakePageContents SIZE_T
EntryAddress PEPT_PML1_ENTRY
OriginalEntry EPT_PML1_ENTRY
ChangedEntry EPT_PML1_ENTRY
Trampoline PCHAR
IsExecutionHook bool
IsHiddenBreakpoint bool
BreakpointAddresses[MaximumHiddenBreakpointsOnPage] uint64
PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage] CHAR
CountOfBreakpoints uint64
}




