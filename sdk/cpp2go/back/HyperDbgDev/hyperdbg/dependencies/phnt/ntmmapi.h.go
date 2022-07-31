package phnt


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

const(
    MemoryBasicInformation // MEMORY_BASIC_INFORMATION = 1  //col:90
    MemoryWorkingSetInformation // MEMORY_WORKING_SET_INFORMATION = 2  //col:91
    MemoryMappedFilenameInformation // UNICODE_STRING = 3  //col:92
    MemoryRegionInformation // MEMORY_REGION_INFORMATION = 4  //col:93
    MemoryWorkingSetExInformation // MEMORY_WORKING_SET_EX_INFORMATION // since VISTA = 5  //col:94
    MemorySharedCommitInformation // MEMORY_SHARED_COMMIT_INFORMATION // since WIN8 = 6  //col:95
    MemoryImageInformation // MEMORY_IMAGE_INFORMATION = 7  //col:96
    MemoryRegionInformationEx // MEMORY_REGION_INFORMATION = 8  //col:97
    MemoryPrivilegedBasicInformation = 9  //col:98
    MemoryEnclaveImageInformation // MEMORY_ENCLAVE_IMAGE_INFORMATION // since REDSTONE3 = 10  //col:99
    MemoryBasicInformationCapped // 10 = 11  //col:100
    MemoryPhysicalContiguityInformation // MEMORY_PHYSICAL_CONTIGUITY_INFORMATION // since 20H1 = 12  //col:101
    MemoryBadInformation // since WIN11 = 13  //col:102
    MemoryBadInformationAllProcesses // since 22H1 = 14  //col:103
    MaxMemoryInfoClass = 15  //col:104
)


const(
    MemoryLocationInvalid = 1  //col:177
    MemoryLocationResident = 2  //col:178
    MemoryLocationPagefile = 3  //col:179
    MemoryLocationReserved = 4  //col:180
)


const(
    MemoryNotContiguous = 1  //col:272
    MemoryAlignedAndContiguous = 2  //col:273
    MemoryNotResident = 3  //col:274
    MemoryNotEligibleToMakeContiguous = 4  //col:275
    MemoryContiguityStateMax = 5  //col:276
)


const(
    SectionBasicInformation // q; SECTION_BASIC_INFORMATION = 1  //col:427
    SectionImageInformation // q; SECTION_IMAGE_INFORMATION = 2  //col:428
    SectionRelocationInformation // q; PVOID RelocationAddress // name:wow64:whNtQuerySection_SectionRelocationInformation // since WIN7 = 3  //col:429
    SectionOriginalBaseInformation // PVOID BaseAddress = 4  //col:430
    SectionInternalImageInformation // SECTION_INTERNAL_IMAGE_INFORMATION // since REDSTONE2 = 5  //col:431
    MaxSectionInfoClass = 6  //col:432
)


const(
    ViewShare  =  1  //col:516
    ViewUnmap  =  2  //col:517
)


const(
    VmPrefetchInformation // ULONG = 1  //col:646
    VmPagePriorityInformation // OFFER_PRIORITY = 2  //col:647
    VmCfgCallTargetInformation // CFG_CALL_TARGET_LIST_INFORMATION // REDSTONE2 = 3  //col:648
    VmPageDirtyStateInformation // REDSTONE3 = 4  //col:649
    VmImageHotPatchInformation // 19H1 = 5  //col:650
    VmPhysicalContiguityInformation // 20H1 = 6  //col:651
    VmVirtualMachinePrepopulateInformation = 7  //col:652
    VmRemoveFromWorkingSetInformation = 8  //col:653
    MaxVmInfoClass = 9  //col:654
)


const(
    SystemMemoryPartitionInformation // q: MEMORY_PARTITION_CONFIGURATION_INFORMATION = 1  //col:856
    SystemMemoryPartitionMoveMemory // s: MEMORY_PARTITION_TRANSFER_INFORMATION = 2  //col:857
    SystemMemoryPartitionAddPagefile // s: MEMORY_PARTITION_PAGEFILE_INFORMATION = 3  //col:858
    SystemMemoryPartitionCombineMemory // q; s: MEMORY_PARTITION_PAGE_COMBINE_INFORMATION = 4  //col:859
    SystemMemoryPartitionInitialAddMemory // q; s: MEMORY_PARTITION_INITIAL_ADD_INFORMATION = 5  //col:860
    SystemMemoryPartitionGetMemoryEvents // MEMORY_PARTITION_MEMORY_EVENTS_INFORMATION // since REDSTONE2 = 6  //col:861
    SystemMemoryPartitionSetAttributes = 7  //col:862
    SystemMemoryPartitionNodeInformation = 8  //col:863
    SystemMemoryPartitionCreateLargePages = 9  //col:864
    SystemMemoryPartitionDedicatedMemoryInformation = 10  //col:865
    SystemMemoryPartitionOpenDedicatedMemory // 10 = 11  //col:866
    SystemMemoryPartitionMemoryChargeAttributes = 12  //col:867
    SystemMemoryPartitionClearAttributes = 13  //col:868
    SystemMemoryPartitionSetMemoryThresholds // since WIN11 = 14  //col:869
    SystemMemoryPartitionMax = 15  //col:870
)



type (
Ntmmapi interface{
#if ()(ok bool)//col:105
#if ()(ok bool)//col:518
#if ()(ok bool)//col:655
#if ()(ok bool)//col:871
}






)

func NewNtmmapi() { return & ntmmapi{} }

func (n *ntmmapi)#if ()(ok bool){//col:105









return true
}







func (n *ntmmapi)#if ()(ok bool){//col:518






return true
}







func (n *ntmmapi)#if ()(ok bool){//col:655










































































































return true
}







func (n *ntmmapi)#if ()(ok bool){//col:871



































































































































































return true
}









