package phnt


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
Protection ULONG_PTR //col:3
ShareCount ULONG_PTR //col:4
Shared ULONG_PTR //col:5
Node ULONG_PTR //col:6
#ifdefWin64 #ifdef _WIN64 //col:7
VirtualPage ULONG_PTR //col:8
#else #else //col:9
VirtualPage uint32 //col:10
#endif #endif //col:11
}



type MEMORY_WORKING_SET_INFORMATION struct{
NumberOfEntries ULONG_PTR //col:15
WorkingSetInfo[1] MEMORY_WORKING_SET_BLOCK //col:16
}



type MEMORY_REGION_INFORMATION struct{
AllocationBase PVOID //col:20
AllocationProtect uint32 //col:21
Union union //col:22
RegionType uint32 //col:24
Struct struct //col:25
Private uint32 //col:27
MappedDataFile uint32 //col:28
MappedImage uint32 //col:29
MappedPageFile uint32 //col:30
MappedPhysical uint32 //col:31
DirectMapped uint32 //col:32
SoftwareEnclave uint32 //col:33
PageSize64K uint32 //col:34
PlaceholderReservation uint32 //col:35
MappedAwe uint32 //col:36
MappedWriteWatch uint32 //col:37
PageSizeLarge uint32 //col:38
PageSizeHuge uint32 //col:39
Reserved uint32 //col:40
}



type MEMORY_WORKING_SET_EX_BLOCK struct{
Union union //col:50
Struct struct //col:52
Valid ULONG_PTR //col:54
ShareCount ULONG_PTR //col:55
Win32Protection ULONG_PTR //col:56
Shared ULONG_PTR //col:57
Node ULONG_PTR //col:58
Locked ULONG_PTR //col:59
LargePage ULONG_PTR //col:60
Priority ULONG_PTR //col:61
Reserved ULONG_PTR //col:62
SharedOriginal ULONG_PTR //col:63
Bad ULONG_PTR //col:64
Win32GraphicsProtection ULONG_PTR //col:65
#ifdefWin64 #ifdef _WIN64 //col:66
ReservedUlong ULONG_PTR //col:67
#endif #endif //col:68
}



type MEMORY_WORKING_SET_EX_INFORMATION struct{
VirtualAddress PVOID //col:91
Union union //col:92
VirtualAttributes MEMORY_WORKING_SET_EX_BLOCK //col:94
Long ULONG_PTR //col:95
}



type MEMORY_SHARED_COMMIT_INFORMATION struct{
CommitSize SIZE_T //col:100
}



type MEMORY_IMAGE_INFORMATION struct{
ImageBase PVOID //col:104
SizeOfImage SIZE_T //col:105
Union union //col:106
ImageFlags uint32 //col:108
Struct struct //col:109
ImagePartialMap uint32 //col:111
ImageNotExecutable uint32 //col:112
ImageSigningLevel uint32 //col:113
Reserved uint32 //col:114
}



type MEMORY_ENCLAVE_IMAGE_INFORMATION struct{
ImageInfo MEMORY_IMAGE_INFORMATION //col:120
UniqueID[32] uint8 //col:121
AuthorID[32] uint8 //col:122
}



type MEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION struct{
Union union //col:126
Struct struct //col:128
State uint32 //col:130
Reserved uint32 //col:131
}



type MEMORY_PHYSICAL_CONTIGUITY_INFORMATION struct{
VirtualAddress PVOID //col:138
Size ULONG_PTR //col:139
ContiguityUnitSize ULONG_PTR //col:140
Flags uint32 //col:141
ContiguityUnitInformation PMEMORY_PHYSICAL_CONTIGUITY_UNIT_INFORMATION //col:142
}



type MEMORY_FRAME_INFORMATION struct{
UseDescription ULONGLONG //col:146
ListDescription ULONGLONG //col:147
Cold ULONGLONG //col:148
Pinned ULONGLONG //col:149
DontUse ULONGLONG //col:150
Priority ULONGLONG //col:151
NonTradeable ULONGLONG //col:152
Reserved ULONGLONG //col:153
}



type FILEOFFSET_INFORMATION struct{
DontUse ULONGLONG //col:157
Offset ULONGLONG //col:158
Reserved ULONGLONG //col:159
}



type PAGEDIR_INFORMATION struct{
DontUse ULONGLONG //col:163
PageDirectoryBase ULONGLONG //col:164
Reserved ULONGLONG //col:165
}



type UNIQUE_PROCESS_INFORMATION struct{
DontUse ULONGLONG //col:169
UniqueProcessKey ULONGLONG //col:170
Reserved ULONGLONG //col:171
}



type MMPFN_IDENTITY struct{
Union union //col:175
e1 MEMORY_FRAME_INFORMATION //col:177
e2 FILEOFFSET_INFORMATION //col:178
e3 PAGEDIR_INFORMATION //col:179
e4 UNIQUE_PROCESS_INFORMATION //col:180
}



type MMPFN_MEMSNAP_INFORMATION struct{
InitialPageFrameIndex ULONG_PTR //col:202
Count ULONG_PTR //col:203
}



type SECTION_BASIC_INFORMATION struct{
BaseAddress PVOID //col:207
AllocationAttributes uint32 //col:208
MaximumSize LARGE_INTEGER //col:209
}



type SECTION_IMAGE_INFORMATION struct{
TransferAddress PVOID //col:213
ZeroBits uint32 //col:214
MaximumStackSize SIZE_T //col:215
CommittedStackSize SIZE_T //col:216
SubSystemType uint32 //col:217
Union union //col:218
Struct struct //col:220
SubSystemMinorVersion USHORT //col:222
SubSystemMajorVersion USHORT //col:223
}



type SECTION_INTERNAL_IMAGE_INFORMATION struct{
SectionInformation SECTION_IMAGE_INFORMATION //col:260
Union union //col:261
ExtendedFlags uint32 //col:263
Struct struct //col:264
ImageExportSuppressionEnabled uint32 //col:266
ImageCetShadowStacksReady uint32 //col:267
ImageXfgEnabled uint32 //col:268
ImageCetShadowStacksStrictMode uint32 //col:269
ImageCetSetContextIpValidationRelaxedMode uint32 //col:270
ImageCetDynamicApisAllowInProc uint32 //col:271
ImageCetDowngradeReserved1 uint32 //col:272
ImageCetDowngradeReserved2 uint32 //col:273
Reserved uint32 //col:274
}



type MEMORY_RANGE_ENTRY struct{
VirtualAddress PVOID //col:280
NumberOfBytes SIZE_T //col:281
}



type CFG_CALL_TARGET_LIST_INFORMATION struct{
NumberOfEntries uint32 //col:285
Reserved uint32 //col:286
NumberOfEntriesProcessed PULONG //col:287
CallTargetInfo PCFG_CALL_TARGET_INFO //col:288
Section PVOID //col:289
FileOffset ULONGLONG //col:290
}



type MEMORY_PARTITION_CONFIGURATION_INFORMATION struct{
Flags uint32 //col:294
NumaNode uint32 //col:295
Channel uint32 //col:296
NumberOfNumaNodes uint32 //col:297
ResidentAvailablePages ULONG_PTR //col:298
CommittedPages ULONG_PTR //col:299
CommitLimit ULONG_PTR //col:300
PeakCommitment ULONG_PTR //col:301
TotalNumberOfPages ULONG_PTR //col:302
AvailablePages ULONG_PTR //col:303
ZeroPages ULONG_PTR //col:304
FreePages ULONG_PTR //col:305
StandbyPages ULONG_PTR //col:306
StandbyPageCountByPriority[8] ULONG_PTR //col:307
RepurposedPagesByPriority[8] ULONG_PTR //col:308
MaximumCommitLimit ULONG_PTR //col:309
Reserved ULONG_PTR //col:310
PartitionId uint32 //col:311
}



type MEMORY_PARTITION_TRANSFER_INFORMATION struct{
NumberOfPages ULONG_PTR //col:315
NumaNode uint32 //col:316
Flags uint32 //col:317
}



type MEMORY_PARTITION_PAGEFILE_INFORMATION struct{
PageFileName UNICODE_STRING //col:321
MinimumSize LARGE_INTEGER //col:322
MaximumSize LARGE_INTEGER //col:323
Flags uint32 //col:324
}



type MEMORY_PARTITION_PAGE_COMBINE_INFORMATION struct{
StopHandle HANDLE //col:328
Flags uint32 //col:329
TotalNumberOfPages ULONG_PTR //col:330
}



type MEMORY_PARTITION_PAGE_RANGE struct{
StartPage ULONG_PTR //col:334
NumberOfPages ULONG_PTR //col:335
}



type MEMORY_PARTITION_INITIAL_ADD_INFORMATION struct{
Flags uint32 //col:339
NumberOfRanges uint32 //col:340
NumberOfPagesAdded ULONG_PTR //col:341
PartitionRanges[1] MEMORY_PARTITION_PAGE_RANGE //col:342
}



type MEMORY_PARTITION_MEMORY_EVENTS_INFORMATION struct{
Union union //col:346
Struct struct //col:348
CommitEvents uint32 //col:350
Spare uint32 //col:351
}





