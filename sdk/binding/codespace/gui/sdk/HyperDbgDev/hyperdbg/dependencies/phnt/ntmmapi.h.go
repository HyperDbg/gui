package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntmmapi.h.back

const(
_NTMMAPI_H =  //col:1
PAGE_NOACCESS = 0x01 //col:2
PAGE_READONLY = 0x02 //col:3
PAGE_READWRITE = 0x04 //col:4
PAGE_WRITECOPY = 0x08 //col:5
PAGE_EXECUTE = 0x10 //col:6
PAGE_EXECUTE_READ = 0x20 //col:7
PAGE_EXECUTE_READWRITE = 0x40 //col:8
PAGE_EXECUTE_WRITECOPY = 0x80 //col:9
PAGE_GUARD = 0x100 //col:10
PAGE_NOCACHE = 0x200 //col:11
PAGE_WRITECOMBINE = 0x400 //col:12
PAGE_REVERT_TO_FILE_MAP =     0x80000000 //col:13
PAGE_ENCLAVE_THREAD_CONTROL = 0x80000000 //col:14
PAGE_TARGETS_NO_UPDATE =      0x40000000 //col:15
PAGE_TARGETS_INVALID =        0x40000000 //col:16
PAGE_ENCLAVE_UNVALIDATED =    0x20000000 //col:17
PAGE_ENCLAVE_NO_CHANGE =      0x20000000 //col:18
PAGE_ENCLAVE_MASK =           0x10000000 //col:19
PAGE_ENCLAVE_DECOMMIT =       (PAGE_ENCLAVE_MASK | 0) //col:20
PAGE_ENCLAVE_SS_FIRST =       (PAGE_ENCLAVE_MASK | 1) //col:21
PAGE_ENCLAVE_SS_REST =        (PAGE_ENCLAVE_MASK | 2) //col:22
MEM_COMMIT = 0x00001000 //col:23
MEM_RESERVE = 0x00002000 //col:24
MEM_DECOMMIT = 0x00004000 //col:25
MEM_RELEASE = 0x00008000 //col:26
MEM_FREE = 0x00010000 //col:27
MEM_PRIVATE = 0x00020000 //col:28
MEM_MAPPED = 0x00040000 //col:29
MEM_RESET = 0x00080000 //col:30
MEM_TOP_DOWN = 0x00100000 //col:31
MEM_WRITE_WATCH = 0x00200000 //col:32
MEM_PHYSICAL = 0x00400000 //col:33
MEM_ROTATE = 0x00800000 //col:34
MEM_DIFFERENT_IMAGE_BASE_OK = 0x00800000 //col:35
MEM_RESET_UNDO = 0x01000000 //col:36
MEM_LARGE_PAGES = 0x20000000 //col:37
MEM_DOS_LIM = 0x40000000 //col:38
MEM_4MB_PAGES = 0x80000000 //col:39
MEM_64K_PAGES = (MEM_LARGE_PAGES | MEM_PHYSICAL) //col:40
MEM_UNMAP_WITH_TRANSIENT_BOOST = 0x00000001 //col:41
MEM_COALESCE_PLACEHOLDERS = 0x00000001 //col:42
MEM_PRESERVE_PLACEHOLDER = 0x00000002 //col:43
MEM_REPLACE_PLACEHOLDER = 0x00004000 //col:44
MEM_RESERVE_PLACEHOLDER = 0x00040000 //col:45
SEC_HUGE_PAGES = 0x00020000 //col:46
SEC_PARTITION_OWNER_HANDLE = 0x00040000 //col:47
SEC_64K_PAGES = 0x00080000 //col:48
SEC_BASED = 0x00200000 //col:49
SEC_NO_CHANGE = 0x00400000 //col:50
SEC_FILE = 0x00800000 //col:51
SEC_IMAGE = 0x01000000 //col:52
SEC_PROTECTED_IMAGE = 0x02000000 //col:53
SEC_RESERVE = 0x04000000 //col:54
SEC_COMMIT = 0x08000000 //col:55
SEC_NOCACHE = 0x10000000 //col:56
SEC_GLOBAL = 0x20000000 //col:57
SEC_WRITECOMBINE = 0x40000000 //col:58
SEC_LARGE_PAGES = 0x80000000 //col:59
SEC_IMAGE_NO_EXECUTE = (SEC_IMAGE | SEC_NOCACHE) //col:60
MEM_IMAGE = SEC_IMAGE //col:61
MemoryBasicInformation = 0x0 //col:62
MemoryWorkingSetInformation = 0x1 //col:63
MemoryMappedFilenameInformation = 0x2 //col:64
MemoryRegionInformation = 0x3 //col:65
MemoryWorkingSetExInformation = 0x4 //col:66
MemorySharedCommitInformation = 0x5 //col:67
MemoryImageInformation = 0x6 //col:68
MemoryRegionInformationEx = 0x7 //col:69
MemoryPrivilegedBasicInformation = 0x8 //col:70
MemoryEnclaveImageInformation = 0x9 //col:71
MemoryBasicInformationCapped = 0xA //col:72
MemoryPhysicalContiguityInformation = 0xB //col:73
MemoryBadInformation = 0xC //col:74
MemoryBadInformationAllProcesses = 0xD //col:75
MMPFNLIST_ZERO = 0 //col:76
MMPFNLIST_FREE = 1 //col:77
MMPFNLIST_STANDBY = 2 //col:78
MMPFNLIST_MODIFIED = 3 //col:79
MMPFNLIST_MODIFIEDNOWRITE = 4 //col:80
MMPFNLIST_BAD = 5 //col:81
MMPFNLIST_ACTIVE = 6 //col:82
MMPFNLIST_TRANSITION = 7 //col:83
MMPFNUSE_PROCESSPRIVATE = 0 //col:84
MMPFNUSE_FILE = 1 //col:85
MMPFNUSE_PAGEFILEMAPPED = 2 //col:86
MMPFNUSE_PAGETABLE = 3 //col:87
MMPFNUSE_PAGEDPOOL = 4 //col:88
MMPFNUSE_NONPAGEDPOOL = 5 //col:89
MMPFNUSE_SYSTEMPTE = 6 //col:90
MMPFNUSE_SESSIONPRIVATE = 7 //col:91
MMPFNUSE_METAFILE = 8 //col:92
MMPFNUSE_AWEPAGE = 9 //col:93
MMPFNUSE_DRIVERLOCKPAGE = 10 //col:94
MMPFNUSE_KERNELSTACK = 11 //col:95
MEM_EXECUTE_OPTION_ENABLE = 0x1 //col:96
MEM_EXECUTE_OPTION_DISABLE = 0x2 //col:97
MEM_EXECUTE_OPTION_DISABLE_THUNK_EMULATION = 0x4 //col:98
MEM_EXECUTE_OPTION_PERMANENT = 0x8 //col:99
MEM_EXECUTE_OPTION_EXECUTE_DISPATCH_ENABLE = 0x10 //col:100
MEM_EXECUTE_OPTION_IMAGE_DISPATCH_ENABLE = 0x20 //col:101
MEM_EXECUTE_OPTION_VALID_FLAGS = 0x3f //col:102
MAP_PROCESS = 1 //col:103
MAP_SYSTEM = 2 //col:104
MEMORY_PARTITION_QUERY_ACCESS = 0x0001 //col:105
MEMORY_PARTITION_MODIFY_ACCESS = 0x0002 //col:106
MEMORY_PARTITION_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | MEMORY_PARTITION_QUERY_ACCESS | MEMORY_PARTITION_MODIFY_ACCESS) //col:107
SystemMemoryPartitionInformation = 0x0 //col:110
SystemMemoryPartitionMoveMemory = 0x1 //col:111
SystemMemoryPartitionAddPagefile = 0x2 //col:112
SystemMemoryPartitionCombineMemory = 0x3 //col:113
SystemMemoryPartitionInitialAddMemory = 0x4 //col:114
SystemMemoryPartitionGetMemoryEvents = 0x5 //col:115
SystemMemoryPartitionSetAttributes = 0x6 //col:116
SystemMemoryPartitionNodeInformation = 0x7 //col:117
SystemMemoryPartitionCreateLargePages = 0x8 //col:118
SystemMemoryPartitionDedicatedMemoryInformation = 0x9 //col:119
SystemMemoryPartitionOpenDedicatedMemory = 0xA //col:120
SystemMemoryPartitionMemoryChargeAttributes = 0xB //col:121
SystemMemoryPartitionClearAttributes = 0xC //col:122
SystemMemoryPartitionSetMemoryThresholds = 0xD //col:123
SystemMemoryPartitionMax = 0xE //col:124
)

const(
    MemoryBasicInformation  = 1  //col:3
    MemoryWorkingSetInformation  = 2  //col:4
    MemoryMappedFilenameInformation  = 3  //col:5
    MemoryRegionInformation  = 4  //col:6
    MemoryWorkingSetExInformation  = 5  //col:7
    MemorySharedCommitInformation  = 6  //col:8
    MemoryImageInformation  = 7  //col:9
    MemoryRegionInformationEx  = 8  //col:10
    MemoryPrivilegedBasicInformation = 9  //col:11
    MemoryEnclaveImageInformation  = 10  //col:12
    MemoryBasicInformationCapped  = 11  //col:13
    MemoryPhysicalContiguityInformation  = 12  //col:14
    MemoryBadInformation  = 13  //col:15
    MemoryBadInformationAllProcesses  = 14  //col:16
    MaxMemoryInfoClass = 15  //col:17
)


const(
    MemoryLocationInvalid = 1  //col:21
    MemoryLocationResident = 2  //col:22
    MemoryLocationPagefile = 3  //col:23
    MemoryLocationReserved = 4  //col:24
)


const(
    MemoryNotContiguous = 1  //col:28
    MemoryAlignedAndContiguous = 2  //col:29
    MemoryNotResident = 3  //col:30
    MemoryNotEligibleToMakeContiguous = 4  //col:31
    MemoryContiguityStateMax = 5  //col:32
)


const(
    SectionBasicInformation  = 1  //col:36
    SectionImageInformation  = 2  //col:37
    SectionRelocationInformation  = 3  //col:38
    SectionOriginalBaseInformation  = 4  //col:39
    SectionInternalImageInformation  = 5  //col:40
    MaxSectionInfoClass = 6  //col:41
)


const(
    ViewShare  =  1  //col:45
    ViewUnmap  =  2  //col:46
)


const(
    VmPrefetchInformation  = 1  //col:50
    VmPagePriorityInformation  = 2  //col:51
    VmCfgCallTargetInformation  = 3  //col:52
    VmPageDirtyStateInformation  = 4  //col:53
    VmImageHotPatchInformation  = 5  //col:54
    VmPhysicalContiguityInformation  = 6  //col:55
    VmVirtualMachinePrepopulateInformation = 7  //col:56
    VmRemoveFromWorkingSetInformation = 8  //col:57
    MaxVmInfoClass = 9  //col:58
)


const(
    SystemMemoryPartitionInformation  = 1  //col:62
    SystemMemoryPartitionMoveMemory  = 2  //col:63
    SystemMemoryPartitionAddPagefile  = 3  //col:64
    SystemMemoryPartitionCombineMemory  = 4  //col:65
    SystemMemoryPartitionInitialAddMemory  = 5  //col:66
    SystemMemoryPartitionGetMemoryEvents  = 6  //col:67
    SystemMemoryPartitionSetAttributes = 7  //col:68
    SystemMemoryPartitionNodeInformation = 8  //col:69
    SystemMemoryPartitionCreateLargePages = 9  //col:70
    SystemMemoryPartitionDedicatedMemoryInformation = 10  //col:71
    SystemMemoryPartitionOpenDedicatedMemory  = 11  //col:72
    SystemMemoryPartitionMemoryChargeAttributes = 12  //col:73
    SystemMemoryPartitionClearAttributes = 13  //col:74
    SystemMemoryPartitionSetMemoryThresholds  = 14  //col:75
    SystemMemoryPartitionMax = 15  //col:76
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
SoftwareEnclave ULONG
PageSize64K ULONG
PlaceholderReservation ULONG
MappedAwe ULONG
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
Win32GraphicsProtection ULONG_PTR
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
ImageSigningLevel ULONG
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
UseDescription ULONGLONG
ListDescription ULONGLONG
Cold ULONGLONG
Pinned ULONGLONG
DontUse ULONGLONG
Priority ULONGLONG
NonTradeable ULONGLONG
Reserved ULONGLONG
}


type FILEOFFSET_INFORMATION struct{
DontUse ULONGLONG
Offset ULONGLONG
Reserved ULONGLONG
}


type PAGEDIR_INFORMATION struct{
DontUse ULONGLONG
PageDirectoryBase ULONGLONG
Reserved ULONGLONG
}


type UNIQUE_PROCESS_INFORMATION struct{
DontUse ULONGLONG
UniqueProcessKey ULONGLONG
Reserved ULONGLONG
}


type MMPFN_IDENTITY struct{
Union union
e1 MEMORY_FRAME_INFORMATION
e2 FILEOFFSET_INFORMATION
e3 PAGEDIR_INFORMATION
e4 UNIQUE_PROCESS_INFORMATION
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
ImageCetShadowStacksReady ULONG
ImageXfgEnabled ULONG
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
Section PVOID
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
StandbyPageCountByPriority[8] ULONG_PTR
RepurposedPagesByPriority[8] ULONG_PTR
MaximumCommitLimit ULONG_PTR
Reserved ULONG_PTR
PartitionId ULONG
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




