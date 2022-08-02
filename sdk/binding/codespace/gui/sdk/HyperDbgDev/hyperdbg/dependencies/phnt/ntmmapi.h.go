package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntmmapi.h.back

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



type  _MEMORY_WORKING_SET_BLOCK struct{
Protection ULONG_PTR //col:14
ShareCount ULONG_PTR //col:15
Shared ULONG_PTR //col:16
Node ULONG_PTR //col:17
#ifdefWin64 #ifdef _WIN64 //col:18
VirtualPage ULONG_PTR //col:19
#else #else //col:20
VirtualPage uint32 //col:21
#endif #endif //col:22
}


type  _MEMORY_WORKING_SET_INFORMATION struct{
NumberOfEntries ULONG_PTR //col:19
WorkingSetInfo[1] MEMORY_WORKING_SET_BLOCK //col:20
}


type  _MEMORY_REGION_INFORMATION struct{
AllocationBase uintptr //col:43
AllocationProtect uint32 //col:44
Union union //col:45
RegionType uint32 //col:47
Struct struct //col:48
Private uint32 //col:50
MappedDataFile uint32 //col:51
MappedImage uint32 //col:52
MappedPageFile uint32 //col:53
MappedPhysical uint32 //col:54
DirectMapped uint32 //col:55
SoftwareEnclave uint32 //col:56
PageSize64K uint32 //col:57
PlaceholderReservation uint32 //col:58
MappedAwe uint32 //col:59
MappedWriteWatch uint32 //col:60
PageSizeLarge uint32 //col:61
PageSizeHuge uint32 //col:62
Reserved uint32 //col:63
}


type  _MEMORY_WORKING_SET_EX_BLOCK struct{
Union union //col:71
Struct struct //col:73
Valid ULONG_PTR //col:75
ShareCount ULONG_PTR //col:76
Win32Protection ULONG_PTR //col:77
Shared ULONG_PTR //col:78
Node ULONG_PTR //col:79
Locked ULONG_PTR //col:80
LargePage ULONG_PTR //col:81
Priority ULONG_PTR //col:82
Reserved ULONG_PTR //col:83
SharedOriginal ULONG_PTR //col:84
Bad ULONG_PTR //col:85
Win32GraphicsProtection ULONG_PTR //col:86
#ifdefWin64 #ifdef _WIN64 //col:87
ReservedUlong ULONG_PTR //col:88
#endif #endif //col:89
}


type  _MEMORY_WORKING_SET_EX_INFORMATION struct{
VirtualAddress uintptr //col:98
Union union //col:99
VirtualAttributes MEMORY_WORKING_SET_EX_BLOCK //col:101
Long ULONG_PTR //col:102
}


type  _MEMORY_SHARED_COMMIT_INFORMATION struct{
CommitSize int64 //col:103
}


type  _MEMORY_IMAGE_INFORMATION struct{
ImageBase uintptr //col:117
SizeOfImage int64 //col:118
Union union //col:119
ImageFlags uint32 //col:121
Struct struct //col:122
ImagePartialMap uint32 //col:124
ImageNotExecutable uint32 //col:125
ImageSigningLevel uint32 //col:126
Reserved uint32 //col:127
}


type  _MEMORY_ENCLAVE_IMAGE_INFORMATION struct{
ImageInfo MEMORY_IMAGE_INFORMATION //col:125
UniqueID[32] uint8 //col:126
AuthorID[32] uint8 //col:127
}


type  _MEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION struct{
Union union //col:134
Struct struct //col:136
State uint32 //col:138
Reserved uint32 //col:139
}


type  _MEMORY_PHYSICAL_CONTIGUITY_INFORMATION struct{
VirtualAddress uintptr //col:145
Size ULONG_PTR //col:146
ContiguityUnitSize ULONG_PTR //col:147
Flags uint32 //col:148
ContiguityUnitInformation PMEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION //col:149
}


type  _MEMORY_FRAME_INFORMATION struct{
UseDescription ULONGLONG //col:156
ListDescription ULONGLONG //col:157
Cold ULONGLONG //col:158
Pinned ULONGLONG //col:159
DontUse ULONGLONG //col:160
Priority ULONGLONG //col:161
NonTradeable ULONGLONG //col:162
Reserved ULONGLONG //col:163
}


type  _FILEOFFSET_INFORMATION struct{
DontUse ULONGLONG //col:162
Offset ULONGLONG //col:163
Reserved ULONGLONG //col:164
}


type  _PAGEDIR_INFORMATION struct{
DontUse ULONGLONG //col:168
PageDirectoryBase ULONGLONG //col:169
Reserved ULONGLONG //col:170
}


type  _UNIQUE_PROCESS_INFORMATION struct{
DontUse ULONGLONG //col:174
UniqueProcessKey ULONGLONG //col:175
Reserved ULONGLONG //col:176
}


type  _MMPFN_IDENTITY struct{
Union union //col:183
e1 MEMORY_FRAME_INFORMATION //col:185
e2 FILEOFFSET_INFORMATION //col:186
e3 PAGEDIR_INFORMATION //col:187
e4 UNIQUE_PROCESS_INFORMATION //col:188
}


type  _MMPFN_MEMSNAP_INFORMATION struct{
InitialPageFrameIndex ULONG_PTR //col:206
Count ULONG_PTR //col:207
}


type  _SECTION_BASIC_INFORMATION struct{
BaseAddress uintptr //col:212
AllocationAttributes uint32 //col:213
MaximumSize LARGE_INTEGER //col:214
}


type  _SECTION_IMAGE_INFORMATION struct{
TransferAddress uintptr //col:226
ZeroBits uint32 //col:227
MaximumStackSize int64 //col:228
CommittedStackSize int64 //col:229
SubSystemType uint32 //col:230
Union union //col:231
Struct struct //col:233
SubSystemMinorVersion uint16 //col:235
SubSystemMajorVersion uint16 //col:236
}


type  _SECTION_INTERNAL_IMAGE_INFORMATION struct{
SectionInformation SECTION_IMAGE_INFORMATION //col:277
Union union //col:278
ExtendedFlags uint32 //col:280
Struct struct //col:281
ImageExportSuppressionEnabled uint32 //col:283
ImageCetShadowStacksReady uint32 //col:284
ImageXfgEnabled uint32 //col:285
ImageCetShadowStacksStrictMode uint32 //col:286
ImageCetSetContextIpValidationRelaxedMode uint32 //col:287
ImageCetDynamicApisAllowInProc uint32 //col:288
ImageCetDowngradeReserved1 uint32 //col:289
ImageCetDowngradeReserved2 uint32 //col:290
Reserved uint32 //col:291
}


type  _MEMORY_RANGE_ENTRY struct{
VirtualAddress uintptr //col:284
NumberOfBytes int64 //col:285
}


type  _CFG_CALL_TARGET_LIST_INFORMATION struct{
NumberOfEntries uint32 //col:293
Reserved uint32 //col:294
NumberOfEntriesProcessed PULONG //col:295
CallTargetInfo PCFG_CALL_TARGET_INFO //col:296
Section uintptr //col:297
FileOffset ULONGLONG //col:298
}


type  _MEMORY_PARTITION_CONFIGURATION_INFORMATION struct{
Flags uint32 //col:314
NumaNode uint32 //col:315
Channel uint32 //col:316
NumberOfNumaNodes uint32 //col:317
ResidentAvailablePages ULONG_PTR //col:318
CommittedPages ULONG_PTR //col:319
CommitLimit ULONG_PTR //col:320
PeakCommitment ULONG_PTR //col:321
TotalNumberOfPages ULONG_PTR //col:322
AvailablePages ULONG_PTR //col:323
ZeroPages ULONG_PTR //col:324
FreePages ULONG_PTR //col:325
StandbyPages ULONG_PTR //col:326
StandbyPageCountByPriority[8] ULONG_PTR //col:327
RepurposedPagesByPriority[8] ULONG_PTR //col:328
MaximumCommitLimit ULONG_PTR //col:329
Reserved ULONG_PTR //col:330
PartitionId uint32 //col:331
}


type  _MEMORY_PARTITION_TRANSFER_INFORMATION struct{
NumberOfPages ULONG_PTR //col:320
NumaNode uint32 //col:321
Flags uint32 //col:322
}


type  _MEMORY_PARTITION_PAGEFILE_INFORMATION struct{
PageFileName *int16 //col:327
MinimumSize LARGE_INTEGER //col:328
MaximumSize LARGE_INTEGER //col:329
Flags uint32 //col:330
}


type  _MEMORY_PARTITION_PAGE_COMBINE_INFORMATION struct{
StopHandle uintptr //col:333
Flags uint32 //col:334
TotalNumberOfPages ULONG_PTR //col:335
}


type  _MEMORY_PARTITION_PAGE_RANGE struct{
StartPage ULONG_PTR //col:338
NumberOfPages ULONG_PTR //col:339
}


type  _MEMORY_PARTITION_INITIAL_ADD_INFORMATION struct{
Flags uint32 //col:345
NumberOfRanges uint32 //col:346
NumberOfPagesAdded ULONG_PTR //col:347
PartitionRanges[1] MEMORY_PARTITION_PAGE_RANGE //col:348
}


type  _MEMORY_PARTITION_MEMORY_EVENTS_INFORMATION struct{
Union union //col:354
Struct struct //col:356
CommitEvents uint32 //col:358
Spare uint32 //col:359
}




