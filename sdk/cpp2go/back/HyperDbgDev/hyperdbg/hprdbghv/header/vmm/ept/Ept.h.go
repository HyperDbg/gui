package ept
//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\ept\Ept.h.back

const(
MaximumHiddenBreakpointsOnPage = 40 //col:18
PAGE_ATTRIB_READ =  0x2 //col:28
PAGE_ATTRIB_WRITE = 0x4 //col:29
PAGE_ATTRIB_EXEC =  0x8 //col:30
VMM_EPT_PML4E_COUNT = 512 //col:36
VMM_EPT_PML3E_COUNT = 512 //col:42
VMM_EPT_PML2E_COUNT = 512 //col:49
VMM_EPT_PML1E_COUNT = 512 //col:56
SIZE_2_MB = ((SIZE_T)(512 * PAGE_SIZE)) //col:62
ADDRMASK_EPT_PML1_OFFSET(_VAR_) = (_VAR_ & 0xFFFULL) //col:68
ADDRMASK_EPT_PML1_INDEX(_VAR_) = ((_VAR_ & 0x1FF000ULL) >> 12) //col:74
ADDRMASK_EPT_PML2_INDEX(_VAR_) = ((_VAR_ & 0x3FE00000ULL) >> 21) //col:80
ADDRMASK_EPT_PML3_INDEX(_VAR_) = ((_VAR_ & 0x7FC0000000ULL) >> 30) //col:86
ADDRMASK_EPT_PML4_INDEX(_VAR_) = ((_VAR_ & 0xFF8000000000ULL) >> 39) //col:92
EPT_MTRR_RANGE_DESCRIPTOR_MAX = 0x9 //col:165
)

type (
Ept interface{
#define SIZE_2_MB ()(ok bool)//col:148
    DECLSPEC_ALIGN()(ok bool)//col:207
    DECLSPEC_ALIGN()(ok bool)//col:296
}
)

func NewEpt() { return & ept{} }

func (e *ept)#define SIZE_2_MB ()(ok bool){//col:148
/*#define SIZE_2_MB ((SIZE_T)(512 * PAGE_SIZE))
 * 
#define ADDRMASK_EPT_PML1_OFFSET(_VAR_) (_VAR_ & 0xFFFULL)
 * 
#define ADDRMASK_EPT_PML1_INDEX(_VAR_) ((_VAR_ & 0x1FF000ULL) >> 12)
 * 
#define ADDRMASK_EPT_PML2_INDEX(_VAR_) ((_VAR_ & 0x3FE00000ULL) >> 21)
 * 
#define ADDRMASK_EPT_PML3_INDEX(_VAR_) ((_VAR_ & 0x7FC0000000ULL) >> 30)
 * 
#define ADDRMASK_EPT_PML4_INDEX(_VAR_) ((_VAR_ & 0xFF8000000000ULL) >> 39)
 * 
volatile LONG Pml1ModificationAndInvalidationLock;
typedef EPT_PML4E   EPT_PML4_POINTER, *PEPT_PML4_POINTER;
typedef EPT_PDPTE   EPT_PML3_POINTER, *PEPT_PML3_POINTER;
typedef EPT_PDE_2MB EPT_PML2_ENTRY, *PEPT_PML2_ENTRY;
typedef EPT_PDE     EPT_PML2_POINTER, *PEPT_PML2_POINTER;
typedef EPT_PTE     EPT_PML1_ENTRY, *PEPT_PML1_ENTRY;
 * 
typedef struct _VMM_EPT_PAGE_TABLE
{
    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML4_POINTER PML4[VMM_EPT_PML4E_COUNT];
    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML3_POINTER PML3[VMM_EPT_PML3E_COUNT];
     * NOTE: We are using 2MB pages as the smallest paging size in our map, so we do not manage individiual 4096 byte pages.
     * Therefore, we do not allocate any PML1 (4096 byte) paging structures.
    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML2_ENTRY PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT];
} VMM_EPT_PAGE_TABLE, *PVMM_EPT_PAGE_TABLE;*/
return true
}

func (e *ept)    DECLSPEC_ALIGN()(ok bool){//col:207
/*    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML1_ENTRY PML1[VMM_EPT_PML1E_COUNT];
    * 
    union
    {
        PEPT_PML2_ENTRY   Entry;
        PEPT_PML2_POINTER Pointer;
    };
     * 
    LIST_ENTRY DynamicSplitList;
} VMM_EPT_DYNAMIC_SPLIT, *PVMM_EPT_DYNAMIC_SPLIT;*/
return true
}

func (e *ept)    DECLSPEC_ALIGN()(ok bool){//col:296
/*    DECLSPEC_ALIGN(PAGE_SIZE)
    CHAR FakePageContents[PAGE_SIZE];
    LIST_ENTRY PageHookList;
    UINT64 VirtualAddress;
    * this way we can de-allocate the list whenever the hook is finished
    UINT64 AddressOfEptHook2sDetourListEntry;
     * when a hook is hit.
    SIZE_T PhysicalBaseAddress;
    * when a hook is hit.
    SIZE_T PhysicalBaseAddressOfFakePageContents;
    PEPT_PML1_ENTRY EntryAddress;
     * from the page.
    EPT_PML1_ENTRY OriginalEntry;
    EPT_PML1_ENTRY ChangedEntry;
    PCHAR Trampoline;
    BOOLEAN IsExecutionHook;
     * a hidden breakpoint command (not a monitor or hidden detours)
    BOOLEAN IsHiddenBreakpoint;
     * this is only used in hidden breakpoints (not hidden detours)
    UINT64 BreakpointAddresses[MaximumHiddenBreakpointsOnPage];
     * this is only used in hidden breakpoints (not hidden detours)
    CHAR PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage];
     * this is only used in hidden breakpoints (not hidden detours)
    UINT64 CountOfBreakpoints;
} EPT_HOOKED_PAGE_DETAIL, *PEPT_HOOKED_PAGE_DETAIL;*/
return true
}



