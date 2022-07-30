package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntmmapi.h.back

const(
_NTMMAPI_H =  //col:13
PAGE_NOACCESS = 0x01 //col:17
PAGE_READONLY = 0x02 //col:18
PAGE_READWRITE = 0x04 //col:19
PAGE_WRITECOPY = 0x08 //col:20
PAGE_EXECUTE = 0x10 //col:21
PAGE_EXECUTE_READ = 0x20 //col:22
PAGE_EXECUTE_READWRITE = 0x40 //col:23
PAGE_EXECUTE_WRITECOPY = 0x80 //col:24
PAGE_GUARD = 0x100 //col:25
PAGE_NOCACHE = 0x200 //col:26
PAGE_WRITECOMBINE = 0x400 //col:27
PAGE_REVERT_TO_FILE_MAP =     0x80000000 //col:29
PAGE_ENCLAVE_THREAD_CONTROL = 0x80000000 //col:30
PAGE_TARGETS_NO_UPDATE =      0x40000000 //col:31
PAGE_TARGETS_INVALID =        0x40000000 //col:32
PAGE_ENCLAVE_UNVALIDATED =    0x20000000 //col:33
PAGE_ENCLAVE_NO_CHANGE =      0x20000000 //col:34
PAGE_ENCLAVE_MASK =           0x10000000 //col:35
PAGE_ENCLAVE_DECOMMIT =       (PAGE_ENCLAVE_MASK | 0) //col:36
PAGE_ENCLAVE_SS_FIRST =       (PAGE_ENCLAVE_MASK | 1) //col:37
PAGE_ENCLAVE_SS_REST =        (PAGE_ENCLAVE_MASK | 2) //col:38
MEM_COMMIT = 0x00001000 //col:42
MEM_RESERVE = 0x00002000 //col:43
MEM_DECOMMIT = 0x00004000 //col:44
MEM_RELEASE = 0x00008000 //col:45
MEM_FREE = 0x00010000 //col:46
MEM_PRIVATE = 0x00020000 //col:47
MEM_MAPPED = 0x00040000 //col:48
MEM_RESET = 0x00080000 //col:49
MEM_TOP_DOWN = 0x00100000 //col:50
MEM_WRITE_WATCH = 0x00200000 //col:51
MEM_PHYSICAL = 0x00400000 //col:52
MEM_ROTATE = 0x00800000 //col:53
MEM_DIFFERENT_IMAGE_BASE_OK = 0x00800000 //col:54
MEM_RESET_UNDO = 0x01000000 //col:55
MEM_LARGE_PAGES = 0x20000000 //col:56
MEM_DOS_LIM = 0x40000000 //col:57
MEM_4MB_PAGES = 0x80000000 //col:58
MEM_64K_PAGES = (MEM_LARGE_PAGES | MEM_PHYSICAL) //col:59
MEM_UNMAP_WITH_TRANSIENT_BOOST = 0x00000001 //col:61
MEM_COALESCE_PLACEHOLDERS = 0x00000001 //col:62
MEM_PRESERVE_PLACEHOLDER = 0x00000002 //col:63
MEM_REPLACE_PLACEHOLDER = 0x00004000 //col:64
MEM_RESERVE_PLACEHOLDER = 0x00040000 //col:65
SEC_HUGE_PAGES = 0x00020000 //col:67
SEC_PARTITION_OWNER_HANDLE = 0x00040000 //col:68
SEC_64K_PAGES = 0x00080000 //col:69
SEC_BASED = 0x00200000 //col:70
SEC_NO_CHANGE = 0x00400000 //col:71
SEC_FILE = 0x00800000 //col:72
SEC_IMAGE = 0x01000000 //col:73
SEC_PROTECTED_IMAGE = 0x02000000 //col:74
SEC_RESERVE = 0x04000000 //col:75
SEC_COMMIT = 0x08000000 //col:76
SEC_NOCACHE = 0x10000000 //col:77
SEC_GLOBAL = 0x20000000 //col:78
SEC_WRITECOMBINE = 0x40000000 //col:79
SEC_LARGE_PAGES = 0x80000000 //col:80
SEC_IMAGE_NO_EXECUTE = (SEC_IMAGE | SEC_NOCACHE) //col:81
MEM_IMAGE = SEC_IMAGE //col:83
MemoryBasicInformation = 0x0 //col:107
MemoryWorkingSetInformation = 0x1 //col:108
MemoryMappedFilenameInformation = 0x2 //col:109
MemoryRegionInformation = 0x3 //col:110
MemoryWorkingSetExInformation = 0x4 //col:111
MemorySharedCommitInformation = 0x5 //col:112
MemoryImageInformation = 0x6 //col:113
MemoryRegionInformationEx = 0x7 //col:114
MemoryPrivilegedBasicInformation = 0x8 //col:115
MemoryEnclaveImageInformation = 0x9 //col:116
MemoryBasicInformationCapped = 0xA //col:117
MemoryPhysicalContiguityInformation = 0xB //col:118
MemoryBadInformation = 0xC //col:119
MemoryBadInformationAllProcesses = 0xD //col:120
MMPFNLIST_ZERO = 0 //col:303
MMPFNLIST_FREE = 1 //col:304
MMPFNLIST_STANDBY = 2 //col:305
MMPFNLIST_MODIFIED = 3 //col:306
MMPFNLIST_MODIFIEDNOWRITE = 4 //col:307
MMPFNLIST_BAD = 5 //col:308
MMPFNLIST_ACTIVE = 6 //col:309
MMPFNLIST_TRANSITION = 7 //col:310
MMPFNUSE_PROCESSPRIVATE = 0 //col:324
MMPFNUSE_FILE = 1 //col:325
MMPFNUSE_PAGEFILEMAPPED = 2 //col:326
MMPFNUSE_PAGETABLE = 3 //col:327
MMPFNUSE_PAGEDPOOL = 4 //col:328
MMPFNUSE_NONPAGEDPOOL = 5 //col:329
MMPFNUSE_SYSTEMPTE = 6 //col:330
MMPFNUSE_SESSIONPRIVATE = 7 //col:331
MMPFNUSE_METAFILE = 8 //col:332
MMPFNUSE_AWEPAGE = 9 //col:333
MMPFNUSE_DRIVERLOCKPAGE = 10 //col:334
MMPFNUSE_KERNELSTACK = 11 //col:335
MEM_EXECUTE_OPTION_ENABLE = 0x1 //col:521
MEM_EXECUTE_OPTION_DISABLE = 0x2 //col:522
MEM_EXECUTE_OPTION_DISABLE_THUNK_EMULATION = 0x4 //col:523
MEM_EXECUTE_OPTION_PERMANENT = 0x8 //col:524
MEM_EXECUTE_OPTION_EXECUTE_DISPATCH_ENABLE = 0x10 //col:525
MEM_EXECUTE_OPTION_IMAGE_DISPATCH_ENABLE = 0x20 //col:526
MEM_EXECUTE_OPTION_VALID_FLAGS = 0x3f //col:527
MAP_PROCESS = 1 //col:693
MAP_SYSTEM = 2 //col:694
MEMORY_PARTITION_QUERY_ACCESS = 0x0001 //col:845
MEMORY_PARTITION_MODIFY_ACCESS = 0x0002 //col:846
MEMORY_PARTITION_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | MEMORY_PARTITION_QUERY_ACCESS | MEMORY_PARTITION_MODIFY_ACCESS) //col:847
SystemMemoryPartitionInformation = 0x0 //col:873
SystemMemoryPartitionMoveMemory = 0x1 //col:874
SystemMemoryPartitionAddPagefile = 0x2 //col:875
SystemMemoryPartitionCombineMemory = 0x3 //col:876
SystemMemoryPartitionInitialAddMemory = 0x4 //col:877
SystemMemoryPartitionGetMemoryEvents = 0x5 //col:878
SystemMemoryPartitionSetAttributes = 0x6 //col:879
SystemMemoryPartitionNodeInformation = 0x7 //col:880
SystemMemoryPartitionCreateLargePages = 0x8 //col:881
SystemMemoryPartitionDedicatedMemoryInformation = 0x9 //col:882
SystemMemoryPartitionOpenDedicatedMemory = 0xA //col:883
SystemMemoryPartitionMemoryChargeAttributes = 0xB //col:884
SystemMemoryPartitionClearAttributes = 0xC //col:885
SystemMemoryPartitionSetMemoryThresholds = 0xD //col:886
SystemMemoryPartitionMax = 0xE //col:887
)

type     MemoryBasicInformation // MEMORY_BASIC_INFORMATION uint32
const(
    MemoryBasicInformation // MEMORY_BASIC_INFORMATION MEMORY_INFORMATION_CLASS = 1  //col:90
    MemoryWorkingSetInformation // MEMORY_WORKING_SET_INFORMATION MEMORY_INFORMATION_CLASS = 2  //col:91
    MemoryMappedFilenameInformation // UNICODE_STRING MEMORY_INFORMATION_CLASS = 3  //col:92
    MemoryRegionInformation // MEMORY_REGION_INFORMATION MEMORY_INFORMATION_CLASS = 4  //col:93
    MemoryWorkingSetExInformation // MEMORY_WORKING_SET_EX_INFORMATION // since VISTA MEMORY_INFORMATION_CLASS = 5  //col:94
    MemorySharedCommitInformation // MEMORY_SHARED_COMMIT_INFORMATION // since WIN8 MEMORY_INFORMATION_CLASS = 6  //col:95
    MemoryImageInformation // MEMORY_IMAGE_INFORMATION MEMORY_INFORMATION_CLASS = 7  //col:96
    MemoryRegionInformationEx // MEMORY_REGION_INFORMATION MEMORY_INFORMATION_CLASS = 8  //col:97
    MemoryPrivilegedBasicInformation MEMORY_INFORMATION_CLASS = 9  //col:98
    MemoryEnclaveImageInformation // MEMORY_ENCLAVE_IMAGE_INFORMATION // since REDSTONE3 MEMORY_INFORMATION_CLASS = 10  //col:99
    MemoryBasicInformationCapped // 10 MEMORY_INFORMATION_CLASS = 11  //col:100
    MemoryPhysicalContiguityInformation // MEMORY_PHYSICAL_CONTIGUITY_INFORMATION // since 20H1 MEMORY_INFORMATION_CLASS = 12  //col:101
    MemoryBadInformation // since WIN11 MEMORY_INFORMATION_CLASS = 13  //col:102
    MemoryBadInformationAllProcesses // since 22H1 MEMORY_INFORMATION_CLASS = 14  //col:103
    MaxMemoryInfoClass MEMORY_INFORMATION_CLASS = 15  //col:104
)


type     MemoryLocationInvalid uint32
const(
    MemoryLocationInvalid MEMORY_WORKING_SET_EX_LOCATION = 1  //col:177
    MemoryLocationResident MEMORY_WORKING_SET_EX_LOCATION = 2  //col:178
    MemoryLocationPagefile MEMORY_WORKING_SET_EX_LOCATION = 3  //col:179
    MemoryLocationReserved MEMORY_WORKING_SET_EX_LOCATION = 4  //col:180
)


type     MemoryNotContiguous uint32
const(
    MemoryNotContiguous MEMORY_PHYSICAL_CONTIGUITY_UNIT_STATE = 1  //col:272
    MemoryAlignedAndContiguous MEMORY_PHYSICAL_CONTIGUITY_UNIT_STATE = 2  //col:273
    MemoryNotResident MEMORY_PHYSICAL_CONTIGUITY_UNIT_STATE = 3  //col:274
    MemoryNotEligibleToMakeContiguous MEMORY_PHYSICAL_CONTIGUITY_UNIT_STATE = 4  //col:275
    MemoryContiguityStateMax MEMORY_PHYSICAL_CONTIGUITY_UNIT_STATE = 5  //col:276
)


type //    ZeroedPageList = 0 uint32
const(
//    ZeroedPageList  //typedef enum _MMLISTS =  0  //col:314
//    FreePageList  //typedef enum _MMLISTS =  1  //col:315
//    StandbyPageList  //typedef enum _MMLISTS =  2  //col:316
//    ModifiedPageList  //typedef enum _MMLISTS =  3  //col:317
//    ModifiedNoWritePageList  //typedef enum _MMLISTS =  4  //col:318
//    BadPageList  //typedef enum _MMLISTS =  5  //col:319
//    ActiveAndValid  //typedef enum _MMLISTS =  6  //col:320
//    TransitionPage  //typedef enum _MMLISTS =  7  //col:321
)


type //    ProcessPrivatePage uint32
const(
//    ProcessPrivatePage //typedef enum _MMPFNUSE = 1  //col:339
//    MemoryMappedFilePage //typedef enum _MMPFNUSE = 2  //col:340
//    PageFileMappedPage //typedef enum _MMPFNUSE = 3  //col:341
//    PageTablePage //typedef enum _MMPFNUSE = 4  //col:342
//    PagedPoolPage //typedef enum _MMPFNUSE = 5  //col:343
//    NonPagedPoolPage //typedef enum _MMPFNUSE = 6  //col:344
//    SystemPTEPage //typedef enum _MMPFNUSE = 7  //col:345
//    SessionPrivatePage //typedef enum _MMPFNUSE = 8  //col:346
//    MetafilePage //typedef enum _MMPFNUSE = 9  //col:347
//    AWEPage //typedef enum _MMPFNUSE = 10  //col:348
//    DriverLockedPage //typedef enum _MMPFNUSE = 11  //col:349
//    KernelStackPage //typedef enum _MMPFNUSE = 12  //col:350
)


type     SectionBasicInformation // q; SECTION_BASIC_INFORMATION uint32
const(
    SectionBasicInformation // q; SECTION_BASIC_INFORMATION SECTION_INFORMATION_CLASS = 1  //col:427
    SectionImageInformation // q; SECTION_IMAGE_INFORMATION SECTION_INFORMATION_CLASS = 2  //col:428
    SectionRelocationInformation // q; PVOID RelocationAddress // name:wow64:whNtQuerySection_SectionRelocationInformation // since WIN7 SECTION_INFORMATION_CLASS = 3  //col:429
    SectionOriginalBaseInformation // PVOID BaseAddress SECTION_INFORMATION_CLASS = 4  //col:430
    SectionInternalImageInformation // SECTION_INTERNAL_IMAGE_INFORMATION // since REDSTONE2 SECTION_INFORMATION_CLASS = 5  //col:431
    MaxSectionInfoClass SECTION_INFORMATION_CLASS = 6  //col:432
)


type     ViewShare = 1 uint32
const(
    ViewShare  SECTION_INHERIT =  1  //col:516
    ViewUnmap  SECTION_INHERIT =  2  //col:517
)


type     VmPrefetchInformation // ULONG uint32
const(
    VmPrefetchInformation // ULONG VIRTUAL_MEMORY_INFORMATION_CLASS = 1  //col:646
    VmPagePriorityInformation // OFFER_PRIORITY VIRTUAL_MEMORY_INFORMATION_CLASS = 2  //col:647
    VmCfgCallTargetInformation // CFG_CALL_TARGET_LIST_INFORMATION // REDSTONE2 VIRTUAL_MEMORY_INFORMATION_CLASS = 3  //col:648
    VmPageDirtyStateInformation // REDSTONE3 VIRTUAL_MEMORY_INFORMATION_CLASS = 4  //col:649
    VmImageHotPatchInformation // 19H1 VIRTUAL_MEMORY_INFORMATION_CLASS = 5  //col:650
    VmPhysicalContiguityInformation // 20H1 VIRTUAL_MEMORY_INFORMATION_CLASS = 6  //col:651
    VmVirtualMachinePrepopulateInformation VIRTUAL_MEMORY_INFORMATION_CLASS = 7  //col:652
    VmRemoveFromWorkingSetInformation VIRTUAL_MEMORY_INFORMATION_CLASS = 8  //col:653
    MaxVmInfoClass VIRTUAL_MEMORY_INFORMATION_CLASS = 9  //col:654
)


type     SystemMemoryPartitionInformation // q: MEMORY_PARTITION_CONFIGURATION_INFORMATION uint32
const(
    SystemMemoryPartitionInformation // q: MEMORY_PARTITION_CONFIGURATION_INFORMATION PARTITION_INFORMATION_CLASS = 1  //col:856
    SystemMemoryPartitionMoveMemory // s: MEMORY_PARTITION_TRANSFER_INFORMATION PARTITION_INFORMATION_CLASS = 2  //col:857
    SystemMemoryPartitionAddPagefile // s: MEMORY_PARTITION_PAGEFILE_INFORMATION PARTITION_INFORMATION_CLASS = 3  //col:858
    SystemMemoryPartitionCombineMemory // q; s: MEMORY_PARTITION_PAGE_COMBINE_INFORMATION PARTITION_INFORMATION_CLASS = 4  //col:859
    SystemMemoryPartitionInitialAddMemory // q; s: MEMORY_PARTITION_INITIAL_ADD_INFORMATION PARTITION_INFORMATION_CLASS = 5  //col:860
    SystemMemoryPartitionGetMemoryEvents // MEMORY_PARTITION_MEMORY_EVENTS_INFORMATION // since REDSTONE2 PARTITION_INFORMATION_CLASS = 6  //col:861
    SystemMemoryPartitionSetAttributes PARTITION_INFORMATION_CLASS = 7  //col:862
    SystemMemoryPartitionNodeInformation PARTITION_INFORMATION_CLASS = 8  //col:863
    SystemMemoryPartitionCreateLargePages PARTITION_INFORMATION_CLASS = 9  //col:864
    SystemMemoryPartitionDedicatedMemoryInformation PARTITION_INFORMATION_CLASS = 10  //col:865
    SystemMemoryPartitionOpenDedicatedMemory // 10 PARTITION_INFORMATION_CLASS = 11  //col:866
    SystemMemoryPartitionMemoryChargeAttributes PARTITION_INFORMATION_CLASS = 12  //col:867
    SystemMemoryPartitionClearAttributes PARTITION_INFORMATION_CLASS = 13  //col:868
    SystemMemoryPartitionSetMemoryThresholds // since WIN11 PARTITION_INFORMATION_CLASS = 14  //col:869
    SystemMemoryPartitionMax PARTITION_INFORMATION_CLASS = 15  //col:870
)



type MEMORY_WORKING_SET_BLOCK struct{
Protection ULONG_PTR
ShareCount ULONG_PTR
Shared ULONG_PTR
Node ULONG_PTR
#ifdefWin64 #ifdef _WIN64
VirtualPage ULONG_PTR
#else #else
VirtualPage ULONG
#endif #endif
}


type MEMORY_WORKING_SET_INFORMATION struct{
NumberOfEntries ULONG_PTR
WorkingSetInfo[1] MEMORY_WORKING_SET_BLOCK
}


type MEMORY_REGION_INFORMATION struct{
AllocationBase PVOID
AllocationProtect ULONG
Union union
RegionType ULONG
Struct struct
Private ULONG
MappedDataFile ULONG
MappedImage ULONG
MappedPageFile ULONG
MappedPhysical ULONG
DirectMapped ULONG
PageSize64K ULONG
MappedWriteWatch ULONG
PageSizeLarge ULONG
PageSizeHuge ULONG
Reserved ULONG
}


type MEMORY_WORKING_SET_EX_BLOCK struct{
Union union
Struct struct
Valid ULONG_PTR
ShareCount ULONG_PTR
Win32Protection ULONG_PTR
Shared ULONG_PTR
Node ULONG_PTR
Locked ULONG_PTR
LargePage ULONG_PTR
Priority ULONG_PTR
Reserved ULONG_PTR
SharedOriginal ULONG_PTR
Bad ULONG_PTR
#ifdefWin64 #ifdef _WIN64
ReservedUlong ULONG_PTR
#endif #endif
}


type MEMORY_WORKING_SET_EX_INFORMATION struct{
VirtualAddress PVOID
Union union
VirtualAttributes MEMORY_WORKING_SET_EX_BLOCK
Long ULONG_PTR
}


type MEMORY_SHARED_COMMIT_INFORMATION struct{
CommitSize SIZE_T
}


type MEMORY_IMAGE_INFORMATION struct{
ImageBase PVOID
SizeOfImage SIZE_T
Union union
ImageFlags ULONG
Struct struct
ImagePartialMap ULONG
ImageNotExecutable ULONG
Reserved ULONG
}


type MEMORY_ENCLAVE_IMAGE_INFORMATION struct{
ImageInfo MEMORY_IMAGE_INFORMATION
UniqueID[32] UCHAR
AuthorID[32] UCHAR
}


type MEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION struct{
Union union
Struct struct
State ULONG
Reserved ULONG
}


type MEMORY_PHYSICAL_CONTIGUITY_INFORMATION struct{
VirtualAddress PVOID
Size ULONG_PTR
ContiguityUnitSize ULONG_PTR
Flags ULONG
ContiguityUnitInformation PMEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION
}


type MEMORY_FRAME_INFORMATION struct{
Priority ULONGLONG
NonTradeable ULONGLONG
Reserved ULONGLONG
}


type FILEOFFSET_INFORMATION struct{
Reserved ULONGLONG
}


type PAGEDIR_INFORMATION struct{
Reserved ULONGLONG
}


type UNIQUE_PROCESS_INFORMATION struct{
Reserved ULONGLONG
}


type MMPFN_IDENTITY struct{
Union union
}


type MMPFN_MEMSNAP_INFORMATION struct{
InitialPageFrameIndex ULONG_PTR
Count ULONG_PTR
}


type SECTION_BASIC_INFORMATION struct{
BaseAddress PVOID
AllocationAttributes ULONG
MaximumSize LARGE_INTEGER
}


type SECTION_IMAGE_INFORMATION struct{
TransferAddress PVOID
ZeroBits ULONG
MaximumStackSize SIZE_T
CommittedStackSize SIZE_T
SubSystemType ULONG
Union union
Struct struct
SubSystemMinorVersion USHORT
SubSystemMajorVersion USHORT
}


type SECTION_INTERNAL_IMAGE_INFORMATION struct{
SectionInformation SECTION_IMAGE_INFORMATION
Union union
ExtendedFlags ULONG
Struct struct
ImageExportSuppressionEnabled ULONG
ImageCetShadowStacksStrictMode ULONG
ImageCetSetContextIpValidationRelaxedMode ULONG
ImageCetDynamicApisAllowInProc ULONG
ImageCetDowngradeReserved1 ULONG
ImageCetDowngradeReserved2 ULONG
Reserved ULONG
}


type MEMORY_RANGE_ENTRY struct{
VirtualAddress PVOID
NumberOfBytes SIZE_T
}


type CFG_CALL_TARGET_LIST_INFORMATION struct{
NumberOfEntries ULONG
Reserved ULONG
NumberOfEntriesProcessed PULONG
CallTargetInfo PCFG_CALL_TARGET_INFO
FileOffset ULONGLONG
}


type MEMORY_PARTITION_CONFIGURATION_INFORMATION struct{
Flags ULONG
NumaNode ULONG
Channel ULONG
NumberOfNumaNodes ULONG
ResidentAvailablePages ULONG_PTR
CommittedPages ULONG_PTR
CommitLimit ULONG_PTR
PeakCommitment ULONG_PTR
TotalNumberOfPages ULONG_PTR
AvailablePages ULONG_PTR
ZeroPages ULONG_PTR
FreePages ULONG_PTR
StandbyPages ULONG_PTR
RepurposedPagesByPriority[8] ULONG_PTR
MaximumCommitLimit ULONG_PTR
}


type MEMORY_PARTITION_TRANSFER_INFORMATION struct{
NumberOfPages ULONG_PTR
NumaNode ULONG
Flags ULONG
}


type MEMORY_PARTITION_PAGEFILE_INFORMATION struct{
PageFileName UNICODE_STRING
MinimumSize LARGE_INTEGER
MaximumSize LARGE_INTEGER
Flags ULONG
}


type MEMORY_PARTITION_PAGE_COMBINE_INFORMATION struct{
StopHandle HANDLE
Flags ULONG
TotalNumberOfPages ULONG_PTR
}


type MEMORY_PARTITION_PAGE_RANGE struct{
StartPage ULONG_PTR
NumberOfPages ULONG_PTR
}


type MEMORY_PARTITION_INITIAL_ADD_INFORMATION struct{
Flags ULONG
NumberOfRanges ULONG
NumberOfPagesAdded ULONG_PTR
PartitionRanges[1] MEMORY_PARTITION_PAGE_RANGE
}


type MEMORY_PARTITION_MEMORY_EVENTS_INFORMATION struct{
Union union
Struct struct
CommitEvents ULONG
Spare ULONG
}



type (
Ntmmapi interface{
 * Attribution 4.0 International ()(ok bool)//col:105
#if ()(ok bool)//col:518
#if ()(ok bool)//col:655
#if ()(ok bool)//col:871
}
)

func NewNtmmapi() { return & ntmmapi{} }

func (n *ntmmapi) * Attribution 4.0 International ()(ok bool){//col:105
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTMMAPI_H
#define _NTMMAPI_H
#define PAGE_NOACCESS 0x01
#define PAGE_READONLY 0x02
#define PAGE_READWRITE 0x04
#define PAGE_WRITECOPY 0x08
#define PAGE_EXECUTE 0x10
#define PAGE_EXECUTE_READ 0x20
#define PAGE_EXECUTE_READWRITE 0x40
#define PAGE_EXECUTE_WRITECOPY 0x80
#define PAGE_GUARD 0x100
#define PAGE_NOCACHE 0x200
#define PAGE_WRITECOMBINE 0x400
#define PAGE_REVERT_TO_FILE_MAP     0x80000000
#define PAGE_ENCLAVE_THREAD_CONTROL 0x80000000
#define PAGE_TARGETS_NO_UPDATE      0x40000000
#define PAGE_TARGETS_INVALID        0x40000000
#define PAGE_ENCLAVE_UNVALIDATED    0x20000000  
#define PAGE_ENCLAVE_NO_CHANGE      0x20000000
#define PAGE_ENCLAVE_MASK           0x10000000  
#define PAGE_ENCLAVE_DECOMMIT       (PAGE_ENCLAVE_MASK | 0) 
#define PAGE_ENCLAVE_SS_FIRST       (PAGE_ENCLAVE_MASK | 1) 
#define PAGE_ENCLAVE_SS_REST        (PAGE_ENCLAVE_MASK | 2)
#define MEM_COMMIT 0x00001000
#define MEM_RESERVE 0x00002000
#define MEM_DECOMMIT 0x00004000
#define MEM_RELEASE 0x00008000
#define MEM_FREE 0x00010000
#define MEM_PRIVATE 0x00020000
#define MEM_MAPPED 0x00040000
#define MEM_RESET 0x00080000
#define MEM_TOP_DOWN 0x00100000
#define MEM_WRITE_WATCH 0x00200000
#define MEM_PHYSICAL 0x00400000
#define MEM_ROTATE 0x00800000
#define MEM_DIFFERENT_IMAGE_BASE_OK 0x00800000
#define MEM_RESET_UNDO 0x01000000
#define MEM_LARGE_PAGES 0x20000000
#define MEM_DOS_LIM 0x40000000
#define MEM_4MB_PAGES 0x80000000
#define MEM_64K_PAGES (MEM_LARGE_PAGES | MEM_PHYSICAL)
#define MEM_UNMAP_WITH_TRANSIENT_BOOST 0x00000001
#define MEM_COALESCE_PLACEHOLDERS 0x00000001
#define MEM_PRESERVE_PLACEHOLDER 0x00000002
#define MEM_REPLACE_PLACEHOLDER 0x00004000
#define MEM_RESERVE_PLACEHOLDER 0x00040000
#define SEC_HUGE_PAGES 0x00020000  
#define SEC_PARTITION_OWNER_HANDLE 0x00040000 
#define SEC_64K_PAGES 0x00080000
#define SEC_BASED 0x00200000
#define SEC_NO_CHANGE 0x00400000
#define SEC_FILE 0x00800000
#define SEC_IMAGE 0x01000000
#define SEC_PROTECTED_IMAGE 0x02000000
#define SEC_RESERVE 0x04000000
#define SEC_COMMIT 0x08000000
#define SEC_NOCACHE 0x10000000
#define SEC_GLOBAL 0x20000000
#define SEC_WRITECOMBINE 0x40000000
#define SEC_LARGE_PAGES 0x80000000
#define SEC_IMAGE_NO_EXECUTE (SEC_IMAGE | SEC_NOCACHE)
#if (PHNT_MODE == PHNT_MODE_KERNEL)
#define MEM_IMAGE SEC_IMAGE
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _MEMORY_INFORMATION_CLASS
{
    MemoryPrivilegedBasicInformation,
    MaxMemoryInfoClass
} MEMORY_INFORMATION_CLASS;*/
return true
}

func (n *ntmmapi)#if ()(ok bool){//col:518
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _SECTION_INHERIT
{
    ViewShare = 1,
    ViewUnmap = 2
} SECTION_INHERIT;*/
return true
}

func (n *ntmmapi)#if ()(ok bool){//col:655
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ _At_(*BaseAddress, _Readable_bytes_(*RegionSize) _Writable_bytes_(*RegionSize) _Post_readable_byte_size_(*RegionSize)) PVOID *BaseAddress,
    _In_ ULONG_PTR ZeroBits,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG AllocationType,
    _In_ ULONG Protect
    );
#if (PHNT_VERSION >= PHNT_REDSTONE5)
NTSYSAPI
NTSTATUS
NTAPI
NtAllocateVirtualMemoryEx(
    _In_ HANDLE ProcessHandle,
    _Inout_ _At_(*BaseAddress, _Readable_bytes_(*RegionSize) _Writable_bytes_(*RegionSize) _Post_readable_byte_size_(*RegionSize)) PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG AllocationType,
    _In_ ULONG PageProtection,
    _Inout_updates_opt_(ExtendedParameterCount) PMEM_EXTENDED_PARAMETER ExtendedParameters,
    _In_ ULONG ExtendedParameterCount
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtFreeVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG FreeType
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReadVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _Out_writes_bytes_(BufferSize) PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesRead
    );
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSAPI
NTSTATUS
NTAPI
NtReadVirtualMemoryEx(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _Out_writes_bytes_(BufferSize) PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesRead,
    _In_ ULONG Flags
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWriteVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _In_reads_bytes_(BufferSize) PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesWritten
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtProtectVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG NewProtect,
    _Out_ PULONG OldProtect
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _In_ MEMORY_INFORMATION_CLASS MemoryInformationClass,
    _Out_writes_bytes_(MemoryInformationLength) PVOID MemoryInformation,
    _In_ SIZE_T MemoryInformationLength,
    _Out_opt_ PSIZE_T ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtFlushVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _Out_ struct _IO_STATUS_BLOCK* IoStatus
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _VIRTUAL_MEMORY_INFORMATION_CLASS
{
    VmVirtualMachinePrepopulateInformation,
    VmRemoveFromWorkingSetInformation,
    MaxVmInfoClass
} VIRTUAL_MEMORY_INFORMATION_CLASS;*/
return true
}

func (n *ntmmapi)#if ()(ok bool){//col:871
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_ VIRTUAL_MEMORY_INFORMATION_CLASS VmInformationClass,
    _In_ ULONG_PTR NumberOfEntries,
    _In_reads_ (NumberOfEntries) PMEMORY_RANGE_ENTRY VirtualAddresses,
    _In_reads_bytes_ (VmInformationLength) PVOID VmInformation,
    _In_ ULONG VmInformationLength
    );
#endif
#define MAP_PROCESS 1
#define MAP_SYSTEM 2
NTSYSCALLAPI
NTSTATUS
NTAPI
NtLockVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG MapType
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUnlockVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG MapType
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateSection(
    _Out_ PHANDLE SectionHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PLARGE_INTEGER MaximumSize,
    _In_ ULONG SectionPageProtection,
    _In_ ULONG AllocationAttributes,
    _In_opt_ HANDLE FileHandle
    );
#if (PHNT_VERSION >= PHNT_REDSTONE5)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateSectionEx(
    _Out_ PHANDLE SectionHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PLARGE_INTEGER MaximumSize,
    _In_ ULONG SectionPageProtection,
    _In_ ULONG AllocationAttributes,
    _In_opt_ HANDLE FileHandle,
    _Inout_updates_opt_(ExtendedParameterCount) PMEM_EXTENDED_PARAMETER ExtendedParameters,
    _In_ ULONG ExtendedParameterCount
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenSection(
    _Out_ PHANDLE SectionHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtMapViewOfSection(
    _In_ HANDLE SectionHandle,
    _In_ HANDLE ProcessHandle,
    _Inout_ _At_(*BaseAddress, _Readable_bytes_(*ViewSize) _Writable_bytes_(*ViewSize) _Post_readable_byte_size_(*ViewSize)) PVOID *BaseAddress,
    _In_ ULONG_PTR ZeroBits,
    _In_ SIZE_T CommitSize,
    _Inout_opt_ PLARGE_INTEGER SectionOffset,
    _Inout_ PSIZE_T ViewSize,
    _In_ SECTION_INHERIT InheritDisposition,
    _In_ ULONG AllocationType,
    _In_ ULONG Win32Protect
    );
#if (PHNT_VERSION >= PHNT_REDSTONE5)
NTSYSAPI
NTSTATUS
NTAPI
NtMapViewOfSectionEx(
    _In_ HANDLE SectionHandle,
    _In_ HANDLE ProcessHandle,
    _Inout_ _At_(*BaseAddress, _Readable_bytes_(*ViewSize) _Writable_bytes_(*ViewSize) _Post_readable_byte_size_(*ViewSize)) PVOID *BaseAddress,
    _Inout_opt_ PLARGE_INTEGER SectionOffset,
    _Inout_ PSIZE_T ViewSize,
    _In_ ULONG AllocationType,
    _In_ ULONG Win32Protect,
    _Inout_updates_opt_(ParameterCount) PMEM_EXTENDED_PARAMETER ExtendedParameters,
    _In_ ULONG ExtendedParameterCount
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUnmapViewOfSection(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUnmapViewOfSectionEx(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _In_ ULONG Flags
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtExtendSection(
    _In_ HANDLE SectionHandle,
    _Inout_ PLARGE_INTEGER NewSectionSize
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySection(
    _In_ HANDLE SectionHandle,
    _In_ SECTION_INFORMATION_CLASS SectionInformationClass,
    _Out_writes_bytes_(SectionInformationLength) PVOID SectionInformation,
    _In_ SIZE_T SectionInformationLength,
    _Out_opt_ PSIZE_T ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAreMappedFilesTheSame(
    _In_ PVOID File1MappedAsAnImage,
    _In_ PVOID File2MappedAsFile
    );
#endif
#ifndef MEMORY_PARTITION_QUERY_ACCESS
#define MEMORY_PARTITION_QUERY_ACCESS 0x0001
#define MEMORY_PARTITION_MODIFY_ACCESS 0x0002
#define MEMORY_PARTITION_ALL_ACCESS \
    (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | \
     MEMORY_PARTITION_QUERY_ACCESS | MEMORY_PARTITION_MODIFY_ACCESS)
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _PARTITION_INFORMATION_CLASS
{
    SystemMemoryPartitionSetAttributes,
    SystemMemoryPartitionNodeInformation,
    SystemMemoryPartitionCreateLargePages,
    SystemMemoryPartitionDedicatedMemoryInformation,
    SystemMemoryPartitionMemoryChargeAttributes,
    SystemMemoryPartitionClearAttributes,
    SystemMemoryPartitionMax
} PARTITION_INFORMATION_CLASS, *PPARTITION_INFORMATION_CLASS;*/
return true
}



