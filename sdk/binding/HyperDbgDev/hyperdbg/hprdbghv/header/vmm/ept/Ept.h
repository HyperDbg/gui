










#pragma once





#define MaximumHiddenBreakpointsOnPage 40









#define PAGE_ATTRIB_READ  0x2
#define PAGE_ATTRIB_WRITE 0x4
#define PAGE_ATTRIB_EXEC  0x8





#define VMM_EPT_PML4E_COUNT 512





#define VMM_EPT_PML3E_COUNT 512






#define VMM_EPT_PML2E_COUNT 512






#define VMM_EPT_PML1E_COUNT 512





#define SIZE_2_MB ((SIZE_T)(512 * PAGE_SIZE))





#define ADDRMASK_EPT_PML1_OFFSET(_VAR_) (_VAR_ & 0xFFFULL)





#define ADDRMASK_EPT_PML1_INDEX(_VAR_) ((_VAR_ & 0x1FF000ULL) >> 12)





#define ADDRMASK_EPT_PML2_INDEX(_VAR_) ((_VAR_ & 0x3FE00000ULL) >> 21)





#define ADDRMASK_EPT_PML3_INDEX(_VAR_) ((_VAR_ & 0x7FC0000000ULL) >> 30)





#define ADDRMASK_EPT_PML4_INDEX(_VAR_) ((_VAR_ & 0xFF8000000000ULL) >> 39)









volatile LONG Pml1ModificationAndInvalidationLock;









typedef EPT_PML4E   EPT_PML4_POINTER, *PEPT_PML4_POINTER;
typedef EPT_PDPTE   EPT_PML3_POINTER, *PEPT_PML3_POINTER;
typedef EPT_PDE_2MB EPT_PML2_ENTRY, *PEPT_PML2_ENTRY;
typedef EPT_PDE     EPT_PML2_POINTER, *PEPT_PML2_POINTER;
typedef EPT_PTE     EPT_PML1_ENTRY, *PEPT_PML1_ENTRY;









typedef struct _VMM_EPT_PAGE_TABLE
{
    


    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML4_POINTER PML4[VMM_EPT_PML4E_COUNT];

    


    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML3_POINTER PML3[VMM_EPT_PML3E_COUNT];

    




    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML2_ENTRY PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT];

} VMM_EPT_PAGE_TABLE, *PVMM_EPT_PAGE_TABLE;





typedef struct _MTRR_RANGE_DESCRIPTOR
{
    SIZE_T PhysicalBaseAddress;
    SIZE_T PhysicalEndAddress;
    UCHAR  MemoryType;
} MTRR_RANGE_DESCRIPTOR, *PMTRR_RANGE_DESCRIPTOR;





#define EPT_MTRR_RANGE_DESCRIPTOR_MAX 0x9
typedef struct _EPT_STATE
{
    LIST_ENTRY            HookedPagesList;                             
    MTRR_RANGE_DESCRIPTOR MemoryRanges[EPT_MTRR_RANGE_DESCRIPTOR_MAX]; 
    ULONG                 NumberOfEnabledMemoryRanges;                 
    EPT_POINTER           EptPointer;                                  
    PVMM_EPT_PAGE_TABLE   EptPageTable;                                

    PVMM_EPT_PAGE_TABLE SecondaryEptPageTable; 

} EPT_STATE, *PEPT_STATE;





typedef struct _VMM_EPT_DYNAMIC_SPLIT
{
    



    DECLSPEC_ALIGN(PAGE_SIZE)
    EPT_PML1_ENTRY PML1[VMM_EPT_PML1E_COUNT];

    



    union
    {
        PEPT_PML2_ENTRY   Entry;
        PEPT_PML2_POINTER Pointer;
    };

    



    LIST_ENTRY DynamicSplitList;

} VMM_EPT_DYNAMIC_SPLIT, *PVMM_EPT_DYNAMIC_SPLIT;





typedef struct _EPT_HOOKED_PAGE_DETAIL
{
    DECLSPEC_ALIGN(PAGE_SIZE)
    CHAR FakePageContents[PAGE_SIZE];

    


    LIST_ENTRY PageHookList;

    


    UINT64 VirtualAddress;

    



    UINT64 AddressOfEptHook2sDetourListEntry;

    



    SIZE_T PhysicalBaseAddress;

    



    SIZE_T PhysicalBaseAddressOfFakePageContents;

    


    PEPT_PML1_ENTRY EntryAddress;

    



    EPT_PML1_ENTRY OriginalEntry;

    


    EPT_PML1_ENTRY ChangedEntry;

    


    PCHAR Trampoline;

    


    BOOLEAN IsExecutionHook;

    



    BOOLEAN IsHiddenBreakpoint;

    



    UINT64 BreakpointAddresses[MaximumHiddenBreakpointsOnPage];

    



    CHAR PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage];

    



    UINT64 CountOfBreakpoints;

} EPT_HOOKED_PAGE_DETAIL, *PEPT_HOOKED_PAGE_DETAIL;













static PEPT_PML2_ENTRY
EptGetPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable, SIZE_T PhysicalAddress);

static VOID
EptSetupPML2Entry(PEPT_PML2_ENTRY NewEntry, SIZE_T PageFrameNumber);

static PVMM_EPT_PAGE_TABLE
EptAllocateAndCreateIdentityPageTable();

static BOOLEAN
EptHandlePageHookExit(_Inout_ PGUEST_REGS                       Regs,
                      _In_ VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
                      _In_ UINT64                               GuestPhysicalAddr);










BOOLEAN
EptCheckFeatures();






BOOLEAN
EptBuildMtrrMap();










BOOLEAN
EptSplitLargePage(PVMM_EPT_PAGE_TABLE EptPageTable,
                  PVOID               PreAllocatedBuffer,
                  SIZE_T              PhysicalAddress,
                  ULONG               CoreIndex);






BOOLEAN
EptLogicalProcessorInitialize();









BOOLEAN
EptHandleEptViolation(_Inout_ PGUEST_REGS Regs,
                      _In_ ULONG          ExitQualification,
                      _In_ UINT64         GuestPhysicalAddr);








PEPT_PML1_ENTRY
EptGetPml1Entry(PVMM_EPT_PAGE_TABLE EptPageTable, SIZE_T PhysicalAddress);







VOID
EptHandleMonitorTrapFlag(PEPT_HOOKED_PAGE_DETAIL HookedEntry);







VOID
EptHandleMisconfiguration(UINT64 GuestAddress);










VOID
EptSetPML1AndInvalidateTLB(_Out_ PEPT_PML1_ENTRY                EntryAddress,
                           _In_ EPT_PML1_ENTRY                  EntryValue,
                           _In_ _Strict_type_match_ INVEPT_TYPE InvalidationType);
