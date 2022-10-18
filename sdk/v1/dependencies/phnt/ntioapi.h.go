package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntioapi.h.back

const(
    FileDirectoryInformation  =  1   //col:3
    FileFullDirectoryInformation  = 2  //col:4
    FileBothDirectoryInformation  = 3  //col:5
    FileBasicInformation  = 4  //col:6
    FileStandardInformation  = 5  //col:7
    FileInternalInformation  = 6  //col:8
    FileEaInformation  = 7  //col:9
    FileAccessInformation  = 8  //col:10
    FileNameInformation  = 9  //col:11
    FileRenameInformation  = 10  //col:12
    FileLinkInformation  = 11  //col:13
    FileNamesInformation  = 12  //col:14
    FileDispositionInformation  = 13  //col:15
    FilePositionInformation  = 14  //col:16
    FileFullEaInformation  = 15  //col:17
    FileModeInformation  = 16  //col:18
    FileAlignmentInformation  = 17  //col:19
    FileAllInformation  = 18  //col:20
    FileAllocationInformation  = 19  //col:21
    FileEndOfFileInformation  = 20  //col:22
    FileAlternateNameInformation  = 21  //col:23
    FileStreamInformation  = 22  //col:24
    FilePipeInformation  = 23  //col:25
    FilePipeLocalInformation  = 24  //col:26
    FilePipeRemoteInformation  = 25  //col:27
    FileMailslotQueryInformation  = 26  //col:28
    FileMailslotSetInformation  = 27  //col:29
    FileCompressionInformation  = 28  //col:30
    FileObjectIdInformation  = 29  //col:31
    FileCompletionInformation  = 30  //col:32
    FileMoveClusterInformation  = 31  //col:33
    FileQuotaInformation  = 32  //col:34
    FileReparsePointInformation  = 33  //col:35
    FileNetworkOpenInformation  = 34  //col:36
    FileAttributeTagInformation  = 35  //col:37
    FileTrackingInformation  = 36  //col:38
    FileIdBothDirectoryInformation  = 37  //col:39
    FileIdFullDirectoryInformation  = 38  //col:40
    FileValidDataLengthInformation  = 39  //col:41
    FileShortNameInformation  = 40  //col:42
    FileIoCompletionNotificationInformation  = 41  //col:43
    FileIoStatusBlockRangeInformation  = 42  //col:44
    FileIoPriorityHintInformation  = 43  //col:45
    FileSfioReserveInformation  = 44  //col:46
    FileSfioVolumeInformation  = 45  //col:47
    FileHardLinkInformation  = 46  //col:48
    FileProcessIdsUsingFileInformation  = 47  //col:49
    FileNormalizedNameInformation  = 48  //col:50
    FileNetworkPhysicalNameInformation  = 49  //col:51
    FileIdGlobalTxDirectoryInformation  = 50  //col:52
    FileIsRemoteDeviceInformation  = 51  //col:53
    FileUnusedInformation = 52  //col:54
    FileNumaNodeInformation  = 53  //col:55
    FileStandardLinkInformation  = 54  //col:56
    FileRemoteProtocolInformation  = 55  //col:57
    FileRenameInformationBypassAccessCheck  = 56  //col:58
    FileLinkInformationBypassAccessCheck  = 57  //col:59
    FileVolumeNameInformation  = 58  //col:60
    FileIdInformation  = 59  //col:61
    FileIdExtdDirectoryInformation  = 60  //col:62
    FileReplaceCompletionInformation  = 61  //col:63
    FileHardLinkFullIdInformation  = 62  //col:64
    FileIdExtdBothDirectoryInformation  = 63  //col:65
    FileDispositionInformationEx  = 64  //col:66
    FileRenameInformationEx  = 65  //col:67
    FileRenameInformationExBypassAccessCheck  = 66  //col:68
    FileDesiredStorageClassInformation  = 67  //col:69
    FileStatInformation  = 68  //col:70
    FileMemoryPartitionInformation  = 69  //col:71
    FileStatLxInformation  = 70  //col:72
    FileCaseSensitiveInformation  = 71  //col:73
    FileLinkInformationEx  = 72  //col:74
    FileLinkInformationExBypassAccessCheck  = 73  //col:75
    FileStorageReserveIdInformation  = 74  //col:76
    FileCaseSensitiveInformationForceAccessCheck  = 75  //col:77
    FileKnownFolderInformation  = 76  //col:78
    FileMaximumInformation = 77  //col:79
)


const(
    IoPriorityVeryLow  =  0   //col:83
    IoPriorityLow  = 2  //col:84
    IoPriorityNormal  = 3  //col:85
    IoPriorityHigh  = 4  //col:86
    IoPriorityCritical  = 5  //col:87
    MaxIoPriorityTypes = 6  //col:88
)


const(
    KnownFolderNone = 1  //col:92
    KnownFolderDesktop = 2  //col:93
    KnownFolderDocuments = 3  //col:94
    KnownFolderDownloads = 4  //col:95
    KnownFolderMusic = 5  //col:96
    KnownFolderPictures = 6  //col:97
    KnownFolderVideos = 7  //col:98
    KnownFolderOther = 8  //col:99
    KnownFolderMax  =  7  //col:100
)


const(
    FileFsVolumeInformation  =  1   //col:104
    FileFsLabelInformation  = 2  //col:105
    FileFsSizeInformation  = 3  //col:106
    FileFsDeviceInformation  = 4  //col:107
    FileFsAttributeInformation  = 5  //col:108
    FileFsControlInformation  = 6  //col:109
    FileFsFullSizeInformation  = 7  //col:110
    FileFsObjectIdInformation  = 8  //col:111
    FileFsDriverPathInformation  = 9  //col:112
    FileFsVolumeFlagsInformation  = 10  //col:113
    FileFsSectorSizeInformation  = 11  //col:114
    FileFsDataCopyInformation  = 12  //col:115
    FileFsMetadataSizeInformation  = 13  //col:116
    FileFsFullSizeInformationEx  = 14  //col:117
    FileFsMaximumInformation = 15  //col:118
)


const(
    DirectoryNotifyInformation  =  1   //col:122
    DirectoryNotifyExtendedInformation  = 2  //col:123
    DirectoryNotifyFullInformation  = 3  //col:124
    DirectoryNotifyMaximumInformation = 4  //col:125
)


const(
    IoCompletionBasicInformation = 1  //col:129
)


const(
    IoSessionEventIgnore = 1  //col:133
    IoSessionEventCreated = 2  //col:134
    IoSessionEventTerminated = 3  //col:135
    IoSessionEventConnected = 4  //col:136
    IoSessionEventDisconnected = 5  //col:137
    IoSessionEventLogon = 6  //col:138
    IoSessionEventLogoff = 7  //col:139
    IoSessionEventMax = 8  //col:140
)


const(
    IoSessionStateCreated  =  1  //col:144
    IoSessionStateInitialized  =  2  //col:145
    IoSessionStateConnected  =  3  //col:146
    IoSessionStateDisconnected  =  4  //col:147
    IoSessionStateDisconnectedLoggedOn  =  5  //col:148
    IoSessionStateLoggedOn  =  6  //col:149
    IoSessionStateLoggedOff  =  7  //col:150
    IoSessionStateTerminated  =  8  //col:151
    IoSessionStateMax = 9  //col:152
)


const(
    InterfaceTypeUndefined  =  -1  //col:156
    Internal  =  0  //col:157
    Isa  =  1  //col:158
    Eisa  =  2  //col:159
    MicroChannel  =  3  //col:160
    TurboChannel  =  4  //col:161
    PCIBus  =  5  //col:162
    VMEBus  =  6  //col:163
    NuBus  =  7  //col:164
    PCMCIABus  =  8  //col:165
    CBus  =  9  //col:166
    MPIBus  =  10  //col:167
    MPSABus  =  11  //col:168
    ProcessorInternal  =  12  //col:169
    InternalPowerBus  =  13  //col:170
    PNPISABus  =  14  //col:171
    PNPBus  =  15  //col:172
    Vmcs  =  16  //col:173
    ACPIBus  =  17  //col:174
    MaximumInterfaceType = 20  //col:175
)


const(
    Width8Bits = 1  //col:179
    Width16Bits = 2  //col:180
    Width32Bits = 3  //col:181
    Width64Bits = 4  //col:182
    WidthNoWrap = 5  //col:183
    MaximumDmaWidth = 6  //col:184
)


const(
    Compatible = 1  //col:188
    TypeA = 2  //col:189
    TypeB = 3  //col:190
    TypeC = 4  //col:191
    TypeF = 5  //col:192
    MaximumDmaSpeed = 6  //col:193
)


const(
    ConfigurationSpaceUndefined  =  -1  //col:197
    Cmos = 2  //col:198
    EisaConfiguration = 3  //col:199
    Pos = 4  //col:200
    CbusConfiguration = 5  //col:201
    PCIConfiguration = 6  //col:202
    VMEConfiguration = 7  //col:203
    NuBusConfiguration = 8  //col:204
    PCMCIAConfiguration = 9  //col:205
    MPIConfiguration = 10  //col:206
    MPSAConfiguration = 11  //col:207
    PNPISAConfiguration = 12  //col:208
    SgiInternalConfiguration = 13  //col:209
    MaximumBusDataType = 14  //col:210
)



type  _IO_STATUS_BLOCK struct{
Union union //col:9
Status NTSTATUS //col:11
Pointer uintptr //col:12
}


type  _FILE_IO_COMPLETION_INFORMATION struct{
KeyContext uintptr //col:17
ApcContext uintptr //col:18
IoStatusBlock IO_STATUS_BLOCK //col:19
}


type  _FILE_BASIC_INFORMATION struct{
CreationTime LARGE_INTEGER //col:25
LastAccessTime LARGE_INTEGER //col:26
LastWriteTime LARGE_INTEGER //col:27
ChangeTime LARGE_INTEGER //col:28
FileAttributes uint32 //col:29
}


type  _FILE_STANDARD_INFORMATION struct{
AllocationSize LARGE_INTEGER //col:33
EndOfFile LARGE_INTEGER //col:34
NumberOfLinks uint32 //col:35
DeletePending bool //col:36
Directory bool //col:37
}


type  _FILE_STANDARD_INFORMATION_EX struct{
AllocationSize LARGE_INTEGER //col:43
EndOfFile LARGE_INTEGER //col:44
NumberOfLinks uint32 //col:45
DeletePending bool //col:46
Directory bool //col:47
AlternateStream bool //col:48
MetadataAttribute bool //col:49
}


type  _FILE_INTERNAL_INFORMATION struct{
IndexNumber LARGE_INTEGER //col:47
}


type  _FILE_EA_INFORMATION struct{
EaSize uint32 //col:51
}


type  _FILE_ACCESS_INFORMATION struct{
AccessFlags ACCESS_MASK //col:55
}


type  _FILE_POSITION_INFORMATION struct{
CurrentByteOffset LARGE_INTEGER //col:59
}


type  _FILE_MODE_INFORMATION struct{
Mode uint32 //col:63
}


type  _FILE_ALIGNMENT_INFORMATION struct{
AlignmentRequirement uint32 //col:67
}


type  _FILE_NAME_INFORMATION struct{
FileNameLength uint32 //col:72
FileName[1] WCHAR //col:73
}


type  _FILE_ALL_INFORMATION struct{
BasicInformation FILE_BASIC_INFORMATION //col:84
StandardInformation FILE_STANDARD_INFORMATION //col:85
InternalInformation FILE_INTERNAL_INFORMATION //col:86
EaInformation FILE_EA_INFORMATION //col:87
AccessInformation FILE_ACCESS_INFORMATION //col:88
PositionInformation FILE_POSITION_INFORMATION //col:89
ModeInformation FILE_MODE_INFORMATION //col:90
AlignmentInformation FILE_ALIGNMENT_INFORMATION //col:91
NameInformation FILE_NAME_INFORMATION //col:92
}


type  _FILE_NETWORK_OPEN_INFORMATION struct{
CreationTime LARGE_INTEGER //col:94
LastAccessTime LARGE_INTEGER //col:95
LastWriteTime LARGE_INTEGER //col:96
ChangeTime LARGE_INTEGER //col:97
AllocationSize LARGE_INTEGER //col:98
EndOfFile LARGE_INTEGER //col:99
FileAttributes uint32 //col:100
}


type  _FILE_ATTRIBUTE_TAG_INFORMATION struct{
FileAttributes uint32 //col:99
ReparseTag uint32 //col:100
}


type  _FILE_ALLOCATION_INFORMATION struct{
AllocationSize LARGE_INTEGER //col:103
}


type  _FILE_COMPRESSION_INFORMATION struct{
CompressedFileSize LARGE_INTEGER //col:112
CompressionFormat uint16 //col:113
CompressionUnitShift uint8 //col:114
ChunkShift uint8 //col:115
ClusterShift uint8 //col:116
Reserved[3] uint8 //col:117
}


type  _FILE_DISPOSITION_INFORMATION struct{
DeleteFile bool //col:116
}


type  _FILE_END_OF_FILE_INFORMATION struct{
EndOfFile LARGE_INTEGER //col:120
}


type  _FILE_END_OF_FILE_INFORMATION_EX  struct{
EndOfFile LARGE_INTEGER //col:127
PagingFileSizeInMM LARGE_INTEGER //col:128
PagingFileMaxSize LARGE_INTEGER //col:129
Flags uint32 //col:130
}


type  _FILE_VALID_DATA_LENGTH_INFORMATION struct{
ValidDataLength LARGE_INTEGER //col:131
}


type  _FILE_LINK_INFORMATION struct{
#if(PhntVersion>=PhntRedstone5) #if (PHNT_VERSION >= PHNT_REDSTONE5) //col:139
Union union //col:140
ReplaceIfExists bool //col:142
Flags uint32 //col:143
}


type  _FILE_LINK_INFORMATION_EX struct{
Flags uint32 //col:153
RootDirectory uintptr //col:154
FileNameLength uint32 //col:155
FileName[1] WCHAR //col:156
}


type  _FILE_MOVE_CLUSTER_INFORMATION struct{
ClusterCount uint32 //col:160
RootDirectory uintptr //col:161
FileNameLength uint32 //col:162
FileName[1] WCHAR //col:163
}


type  _FILE_RENAME_INFORMATION struct{
ReplaceIfExists bool //col:167
RootDirectory uintptr //col:168
FileNameLength uint32 //col:169
FileName[1] WCHAR //col:170
}


type  _FILE_RENAME_INFORMATION_EX struct{
Flags uint32 //col:174
RootDirectory uintptr //col:175
FileNameLength uint32 //col:176
FileName[1] WCHAR //col:177
}


type  _FILE_STREAM_INFORMATION struct{
NextEntryOffset uint32 //col:182
StreamNameLength uint32 //col:183
StreamSize LARGE_INTEGER //col:184
StreamAllocationSize LARGE_INTEGER //col:185
StreamName[1] WCHAR //col:186
}


type  _FILE_TRACKING_INFORMATION struct{
DestinationFile uintptr //col:188
ObjectInformationLength uint32 //col:189
ObjectInformation[1] int8 //col:190
}


type  _FILE_COMPLETION_INFORMATION struct{
Port uintptr //col:193
Key uintptr //col:194
}


type  _FILE_PIPE_INFORMATION struct{
ReadMode uint32 //col:198
CompletionMode uint32 //col:199
}


type  _FILE_PIPE_LOCAL_INFORMATION struct{
NamedPipeType uint32 //col:211
NamedPipeConfiguration uint32 //col:212
MaximumInstances uint32 //col:213
CurrentInstances uint32 //col:214
InboundQuota uint32 //col:215
ReadDataAvailable uint32 //col:216
OutboundQuota uint32 //col:217
WriteQuotaAvailable uint32 //col:218
NamedPipeState uint32 //col:219
NamedPipeEnd uint32 //col:220
}


type  _FILE_PIPE_REMOTE_INFORMATION struct{
CollectDataTime LARGE_INTEGER //col:216
MaximumCollectionCount uint32 //col:217
}


type  _FILE_MAILSLOT_QUERY_INFORMATION struct{
MaximumMessageSize uint32 //col:224
MailslotQuota uint32 //col:225
NextMessageSize uint32 //col:226
MessagesAvailable uint32 //col:227
ReadTimeout LARGE_INTEGER //col:228
}


type  _FILE_MAILSLOT_SET_INFORMATION struct{
ReadTimeout PLARGE_INTEGER //col:228
}


type  _FILE_REPARSE_POINT_INFORMATION struct{
FileReference LONGLONG //col:233
Tag uint32 //col:234
}


type  _FILE_LINK_ENTRY_INFORMATION struct{
NextEntryOffset uint32 //col:240
ParentFileId LONGLONG //col:241
FileNameLength uint32 //col:242
FileName[1] WCHAR //col:243
}


type  _FILE_LINKS_INFORMATION struct{
BytesNeeded uint32 //col:246
EntriesReturned uint32 //col:247
Entry FILE_LINK_ENTRY_INFORMATION //col:248
}


type  _FILE_NETWORK_PHYSICAL_NAME_INFORMATION struct{
FileNameLength uint32 //col:251
FileName[1] WCHAR //col:252
}


type  _FILE_STANDARD_LINK_INFORMATION struct{
NumberOfAccessibleLinks uint32 //col:258
TotalNumberOfLinks uint32 //col:259
DeletePending bool //col:260
Directory bool //col:261
}


type  _FILE_SFIO_RESERVE_INFORMATION struct{
RequestsPerPeriod uint32 //col:267
Period uint32 //col:268
RetryFailures bool //col:269
Discardable bool //col:270
RequestSize uint32 //col:271
NumOutstandingRequests uint32 //col:272
}


type  _FILE_SFIO_VOLUME_INFORMATION struct{
MaximumRequestsPerPeriod uint32 //col:273
MinimumPeriod uint32 //col:274
MinimumTransferSize uint32 //col:275
}


type  _FILE_IO_PRIORITY_HINT_INFORMATION_EX struct{
PriorityHint IO_PRIORITY_HINT //col:278
BoostOutstanding bool //col:279
}


type  _FILE_IO_COMPLETION_NOTIFICATION_INFORMATION struct{
Flags uint32 //col:282
}


type  _FILE_PROCESS_IDS_USING_FILE_INFORMATION struct{
NumberOfProcessIdsInList uint32 //col:287
ProcessIdList[1] ULONG_PTR //col:288
}


type  _FILE_IS_REMOTE_DEVICE_INFORMATION struct{
IsRemote bool //col:291
}


type  _FILE_NUMA_NODE_INFORMATION struct{
NodeNumber uint16 //col:295
}


type  _FILE_IOSTATUSBLOCK_RANGE_INFORMATION struct{
IoStatusBlockRange PUCHAR //col:300
Length uint32 //col:301
}


type  _FILE_REMOTE_PROTOCOL_INFORMATION struct{
StructureVersion uint16 //col:314
StructureSize uint16 //col:315
Protocol uint32 //col:316
ProtocolMajorVersion uint16 //col:317
ProtocolMinorVersion uint16 //col:318
ProtocolRevision uint16 //col:319
Reserved uint16 //col:320
Flags uint32 //col:321
Struct struct //col:322
Reserved[8] uint32 //col:324
}


type  _FILE_INTEGRITY_STREAM_INFORMATION struct{
ChecksumAlgorithm uint16 //col:355
ChecksumChunkShift uint8 //col:356
ClusterShift uint8 //col:357
Flags uint32 //col:358
}


type  _FILE_VOLUME_NAME_INFORMATION struct{
DeviceNameLength uint32 //col:360
DeviceName[1] WCHAR //col:361
}


type  _FILE_ID_INFORMATION struct{
VolumeSerialNumber ULONGLONG //col:365
FileId FILE_ID_128 //col:366
}


type  _FILE_ID_EXTD_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:382
FileIndex uint32 //col:383
CreationTime LARGE_INTEGER //col:384
LastAccessTime LARGE_INTEGER //col:385
LastWriteTime LARGE_INTEGER //col:386
ChangeTime LARGE_INTEGER //col:387
EndOfFile LARGE_INTEGER //col:388
AllocationSize LARGE_INTEGER //col:389
FileAttributes uint32 //col:390
FileNameLength uint32 //col:391
EaSize uint32 //col:392
ReparsePointTag uint32 //col:393
FileId FILE_ID_128 //col:394
FileName[1] WCHAR //col:395
}


type  _FILE_LINK_ENTRY_FULL_ID_INFORMATION struct{
NextEntryOffset uint32 //col:389
ParentFileId FILE_ID_128 //col:390
FileNameLength uint32 //col:391
FileName[1] WCHAR //col:392
}


type  _FILE_LINKS_FULL_ID_INFORMATION  struct{
BytesNeeded uint32 //col:395
EntriesReturned uint32 //col:396
Entry FILE_LINK_ENTRY_FULL_ID_INFORMATION //col:397
}


type  _FILE_ID_EXTD_BOTH_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:414
FileIndex uint32 //col:415
CreationTime LARGE_INTEGER //col:416
LastAccessTime LARGE_INTEGER //col:417
LastWriteTime LARGE_INTEGER //col:418
ChangeTime LARGE_INTEGER //col:419
EndOfFile LARGE_INTEGER //col:420
AllocationSize LARGE_INTEGER //col:421
FileAttributes uint32 //col:422
FileNameLength uint32 //col:423
EaSize uint32 //col:424
ReparsePointTag uint32 //col:425
FileId FILE_ID_128 //col:426
ShortNameLength CCHAR //col:427
ShortName[12] WCHAR //col:428
FileName[1] WCHAR //col:429
}


type  _FILE_STAT_INFORMATION struct{
FileId LARGE_INTEGER //col:428
CreationTime LARGE_INTEGER //col:429
LastAccessTime LARGE_INTEGER //col:430
LastWriteTime LARGE_INTEGER //col:431
ChangeTime LARGE_INTEGER //col:432
AllocationSize LARGE_INTEGER //col:433
EndOfFile LARGE_INTEGER //col:434
FileAttributes uint32 //col:435
ReparseTag uint32 //col:436
NumberOfLinks uint32 //col:437
EffectiveAccess ACCESS_MASK //col:438
}


type  _FILE_MEMORY_PARTITION_INFORMATION struct{
OwnerPartitionHandle uintptr //col:438
Union union //col:439
Struct struct //col:441
NoCrossPartitionAccess uint8 //col:443
Spare[3] uint8 //col:444
}


type  _FILE_STAT_LX_INFORMATION struct{
FileId LARGE_INTEGER //col:461
CreationTime LARGE_INTEGER //col:462
LastAccessTime LARGE_INTEGER //col:463
LastWriteTime LARGE_INTEGER //col:464
ChangeTime LARGE_INTEGER //col:465
AllocationSize LARGE_INTEGER //col:466
EndOfFile LARGE_INTEGER //col:467
FileAttributes uint32 //col:468
ReparseTag uint32 //col:469
NumberOfLinks uint32 //col:470
EffectiveAccess ACCESS_MASK //col:471
LxFlags uint32 //col:472
LxUid uint32 //col:473
LxGid uint32 //col:474
LxMode uint32 //col:475
LxDeviceIdMajor uint32 //col:476
LxDeviceIdMinor uint32 //col:477
}


type  _FILE_CASE_SENSITIVE_INFORMATION struct{
Flags uint32 //col:465
}


type  _FILE_KNOWN_FOLDER_INFORMATION struct{
Type FILE_KNOWN_FOLDER_TYPE //col:469
}


type  _FILE_DIRECTORY_INFORMATION struct{
NextEntryOffset uint32 //col:483
FileIndex uint32 //col:484
CreationTime LARGE_INTEGER //col:485
LastAccessTime LARGE_INTEGER //col:486
LastWriteTime LARGE_INTEGER //col:487
ChangeTime LARGE_INTEGER //col:488
EndOfFile LARGE_INTEGER //col:489
AllocationSize LARGE_INTEGER //col:490
FileAttributes uint32 //col:491
FileNameLength uint32 //col:492
FileName[1] WCHAR //col:493
}


type  _FILE_FULL_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:498
FileIndex uint32 //col:499
CreationTime LARGE_INTEGER //col:500
LastAccessTime LARGE_INTEGER //col:501
LastWriteTime LARGE_INTEGER //col:502
ChangeTime LARGE_INTEGER //col:503
EndOfFile LARGE_INTEGER //col:504
AllocationSize LARGE_INTEGER //col:505
FileAttributes uint32 //col:506
FileNameLength uint32 //col:507
EaSize uint32 //col:508
FileName[1] WCHAR //col:509
}


type  _FILE_ID_FULL_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:514
FileIndex uint32 //col:515
CreationTime LARGE_INTEGER //col:516
LastAccessTime LARGE_INTEGER //col:517
LastWriteTime LARGE_INTEGER //col:518
ChangeTime LARGE_INTEGER //col:519
EndOfFile LARGE_INTEGER //col:520
AllocationSize LARGE_INTEGER //col:521
FileAttributes uint32 //col:522
FileNameLength uint32 //col:523
EaSize uint32 //col:524
FileId LARGE_INTEGER //col:525
FileName[1] WCHAR //col:526
}


type  _FILE_BOTH_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:531
FileIndex uint32 //col:532
CreationTime LARGE_INTEGER //col:533
LastAccessTime LARGE_INTEGER //col:534
LastWriteTime LARGE_INTEGER //col:535
ChangeTime LARGE_INTEGER //col:536
EndOfFile LARGE_INTEGER //col:537
AllocationSize LARGE_INTEGER //col:538
FileAttributes uint32 //col:539
FileNameLength uint32 //col:540
EaSize uint32 //col:541
ShortNameLength CCHAR //col:542
ShortName[12] WCHAR //col:543
FileName[1] WCHAR //col:544
}


type  _FILE_ID_BOTH_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:549
FileIndex uint32 //col:550
CreationTime LARGE_INTEGER //col:551
LastAccessTime LARGE_INTEGER //col:552
LastWriteTime LARGE_INTEGER //col:553
ChangeTime LARGE_INTEGER //col:554
EndOfFile LARGE_INTEGER //col:555
AllocationSize LARGE_INTEGER //col:556
FileAttributes uint32 //col:557
FileNameLength uint32 //col:558
EaSize uint32 //col:559
ShortNameLength CCHAR //col:560
ShortName[12] WCHAR //col:561
FileId LARGE_INTEGER //col:562
FileName[1] WCHAR //col:563
}


type  _FILE_NAMES_INFORMATION struct{
NextEntryOffset uint32 //col:556
FileIndex uint32 //col:557
FileNameLength uint32 //col:558
FileName[1] WCHAR //col:559
}


type  _FILE_ID_GLOBAL_TX_DIR_INFORMATION struct{
NextEntryOffset uint32 //col:573
FileIndex uint32 //col:574
CreationTime LARGE_INTEGER //col:575
LastAccessTime LARGE_INTEGER //col:576
LastWriteTime LARGE_INTEGER //col:577
ChangeTime LARGE_INTEGER //col:578
EndOfFile LARGE_INTEGER //col:579
AllocationSize LARGE_INTEGER //col:580
FileAttributes uint32 //col:581
FileNameLength uint32 //col:582
FileId LARGE_INTEGER //col:583
LockingTransactionId GUID //col:584
TxInfoFlags uint32 //col:585
FileName[1] WCHAR //col:586
}


type  _FILE_OBJECTID_INFORMATION struct{
FileReference LONGLONG //col:585
ObjectId[16] uint8 //col:586
Union union //col:587
Struct struct //col:589
BirthVolumeId[16] uint8 //col:591
BirthObjectId[16] uint8 //col:592
DomainId[16] uint8 //col:593
}


type  _FILE_FULL_EA_INFORMATION struct{
NextEntryOffset uint32 //col:596
Flags uint8 //col:597
EaNameLength uint8 //col:598
EaValueLength uint16 //col:599
EaName[1] int8 //col:600
}


type  _FILE_GET_EA_INFORMATION struct{
NextEntryOffset uint32 //col:602
EaNameLength uint8 //col:603
EaName[1] int8 //col:604
}


type  _FILE_GET_QUOTA_INFORMATION struct{
NextEntryOffset uint32 //col:608
SidLength uint32 //col:609
Sid SID //col:610
}


type  _FILE_QUOTA_INFORMATION struct{
NextEntryOffset uint32 //col:618
SidLength uint32 //col:619
ChangeTime LARGE_INTEGER //col:620
QuotaUsed LARGE_INTEGER //col:621
QuotaThreshold LARGE_INTEGER //col:622
QuotaLimit LARGE_INTEGER //col:623
Sid SID //col:624
}


type  _FILE_FS_VOLUME_INFORMATION struct{
VolumeCreationTime LARGE_INTEGER //col:626
VolumeSerialNumber uint32 //col:627
VolumeLabelLength uint32 //col:628
SupportsObjects bool //col:629
VolumeLabel[1] WCHAR //col:630
}


type  _FILE_FS_LABEL_INFORMATION struct{
VolumeLabelLength uint32 //col:631
VolumeLabel[1] WCHAR //col:632
}


type  _FILE_FS_SIZE_INFORMATION struct{
TotalAllocationUnits LARGE_INTEGER //col:638
AvailableAllocationUnits LARGE_INTEGER //col:639
SectorsPerAllocationUnit uint32 //col:640
BytesPerSector uint32 //col:641
}


type  _FILE_FS_CONTROL_INFORMATION struct{
FreeSpaceStartFiltering LARGE_INTEGER //col:647
FreeSpaceThreshold LARGE_INTEGER //col:648
FreeSpaceStopFiltering LARGE_INTEGER //col:649
DefaultQuotaThreshold LARGE_INTEGER //col:650
DefaultQuotaLimit LARGE_INTEGER //col:651
FileSystemControlFlags uint32 //col:652
}


type  _FILE_FS_FULL_SIZE_INFORMATION struct{
TotalAllocationUnits LARGE_INTEGER //col:655
CallerAvailableAllocationUnits LARGE_INTEGER //col:656
ActualAvailableAllocationUnits LARGE_INTEGER //col:657
SectorsPerAllocationUnit uint32 //col:658
BytesPerSector uint32 //col:659
}


type  _FILE_FS_OBJECTID_INFORMATION struct{
ObjectId[16] uint8 //col:660
ExtendedInfo[48] uint8 //col:661
}


type  _FILE_FS_DEVICE_INFORMATION struct{
DeviceType DEVICE_TYPE //col:665
Characteristics uint32 //col:666
}


type  _FILE_FS_ATTRIBUTE_INFORMATION struct{
FileSystemAttributes uint32 //col:672
MaximumComponentNameLength int32 //col:673
FileSystemNameLength uint32 //col:674
FileSystemName[1] WCHAR //col:675
}


type  _FILE_FS_DRIVER_PATH_INFORMATION struct{
DriverInPath bool //col:678
DriverNameLength uint32 //col:679
DriverName[1] WCHAR //col:680
}


type  _FILE_FS_VOLUME_FLAGS_INFORMATION struct{
Flags uint32 //col:682
}


type  _FILE_FS_SECTOR_SIZE_INFORMATION struct{
LogicalBytesPerSector uint32 //col:692
PhysicalBytesPerSectorForAtomicity uint32 //col:693
PhysicalBytesPerSectorForPerformance uint32 //col:694
FileSystemEffectivePhysicalBytesPerSectorForAtomicity uint32 //col:695
Flags uint32 //col:696
ByteOffsetForSectorAlignment uint32 //col:697
ByteOffsetForPartitionAlignment uint32 //col:698
}


type  _FILE_FS_DATA_COPY_INFORMATION struct{
NumberOfCopies uint32 //col:696
}


type  _FILE_FS_METADATA_SIZE_INFORMATION struct{
TotalMetadataAllocationUnits LARGE_INTEGER //col:702
SectorsPerAllocationUnit uint32 //col:703
BytesPerSector uint32 //col:704
}


type  _FILE_FS_FULL_SIZE_INFORMATION_EX struct{
ActualTotalAllocationUnits ULONGLONG //col:718
ActualAvailableAllocationUnits ULONGLONG //col:719
ActualPoolUnavailableAllocationUnits ULONGLONG //col:720
CallerTotalAllocationUnits ULONGLONG //col:721
CallerAvailableAllocationUnits ULONGLONG //col:722
CallerPoolUnavailableAllocationUnits ULONGLONG //col:723
UsedAllocationUnits ULONGLONG //col:724
TotalReservedAllocationUnits ULONGLONG //col:725
VolumeStorageReserveAllocationUnits ULONGLONG //col:726
AvailableCommittedAllocationUnits ULONGLONG //col:727
PoolAvailableAllocationUnits ULONGLONG //col:728
SectorsPerAllocationUnit uint32 //col:729
BytesPerSector uint32 //col:730
}


type  _IO_COMPLETION_BASIC_INFORMATION struct{
Depth int32 //col:722
}


type  _REPARSE_DATA_BUFFER struct{
ReparseTag uint32 //col:739
ReparseDataLength uint16 //col:740
Reserved uint16 //col:741
FieldSizeBytes(ReparseDataLength) _Field_size_bytes_(ReparseDataLength) //col:742
Union union //col:743
Struct struct //col:745
SubstituteNameOffset uint16 //col:747
SubstituteNameLength uint16 //col:748
PrintNameOffset uint16 //col:749
PrintNameLength uint16 //col:750
Flags uint32 //col:751
PathBuffer[1] WCHAR //col:752
}


type  _FILE_PIPE_ASSIGN_EVENT_BUFFER struct{
EventHandle uintptr //col:758
KeyValue uint32 //col:759
}


type  _FILE_PIPE_PEEK_BUFFER struct{
NamedPipeState uint32 //col:766
ReadDataAvailable uint32 //col:767
NumberOfMessages uint32 //col:768
MessageLength uint32 //col:769
Data[1] int8 //col:770
}


type  _FILE_PIPE_EVENT_BUFFER struct{
NamedPipeState uint32 //col:774
EntryType uint32 //col:775
ByteCount uint32 //col:776
KeyValue uint32 //col:777
NumberRequests uint32 //col:778
}


type  _FILE_PIPE_WAIT_FOR_BUFFER struct{
Timeout LARGE_INTEGER //col:781
NameLength uint32 //col:782
TimeoutSpecified bool //col:783
Name[1] WCHAR //col:784
}


type  _FILE_PIPE_CLIENT_PROCESS_BUFFER struct{
#if!Defined(BuildWow6432) #if !defined(BUILD_WOW6432) //col:791
ClientSession uintptr //col:792
ClientProcess uintptr //col:793
#else #else //col:794
ClientSession ULONGLONG //col:795
ClientProcess ULONGLONG //col:796
#endif #endif //col:797
}


type  _FILE_PIPE_CLIENT_PROCESS_BUFFER_V2  struct{
ClientSession ULONGLONG //col:800
#if!Defined(BuildWow6432) #if !defined(BUILD_WOW6432) //col:801
ClientProcess uintptr //col:802
#else #else //col:803
ClientProcess ULONGLONG //col:804
#endif #endif //col:805
}


type  _FILE_PIPE_CLIENT_PROCESS_BUFFER_EX struct{
#if!Defined(BuildWow6432) #if !defined(BUILD_WOW6432) //col:812
ClientSession uintptr //col:813
ClientProcess uintptr //col:814
#else #else //col:815
ClientSession ULONGLONG //col:816
ClientProcess ULONGLONG //col:817
#endif #endif //col:818
ClientComputerNameLength uint16 //col:819
ClientComputerBuffer[FILE_PIPE_COMPUTER_NAME_LENGTH WCHAR //col:820
}


type  _FILE_PIPE_SILO_ARRIVAL_INPUT struct{
JobHandle uintptr //col:816
}


type  _FILE_PIPE_CREATE_SYMLINK_INPUT struct{
NameOffset uint16 //col:824
NameLength uint16 //col:825
SubstituteNameOffset uint16 //col:826
SubstituteNameLength uint16 //col:827
Flags uint32 //col:828
}


type  _FILE_PIPE_DELETE_SYMLINK_INPUT struct{
NameOffset uint16 //col:829
NameLength uint16 //col:830
}


type  _FILE_MAILSLOT_PEEK_BUFFER struct{
ReadDataAvailable uint32 //col:835
NumberOfMessages uint32 //col:836
MessageLength uint32 //col:837
}


type  _MOUNTMGR_CREATE_POINT_INPUT struct{
SymbolicLinkNameOffset uint16 //col:842
SymbolicLinkNameLength uint16 //col:843
DeviceNameOffset uint16 //col:844
DeviceNameLength uint16 //col:845
}


type  _MOUNTMGR_MOUNT_POINT struct{
SymbolicLinkNameOffset uint32 //col:854
SymbolicLinkNameLength uint16 //col:855
Reserved1 uint16 //col:856
UniqueIdOffset uint32 //col:857
UniqueIdLength uint16 //col:858
Reserved2 uint16 //col:859
DeviceNameOffset uint32 //col:860
DeviceNameLength uint16 //col:861
Reserved3 uint16 //col:862
}


type  _MOUNTMGR_MOUNT_POINTS struct{
Size uint32 //col:860
NumberOfMountPoints uint32 //col:861
MountPoints[1] MOUNTMGR_MOUNT_POINT //col:862
}


type  _MOUNTMGR_DRIVE_LETTER_TARGET struct{
DeviceNameLength uint16 //col:865
DeviceName[1] WCHAR //col:866
}


type  _MOUNTMGR_DRIVE_LETTER_INFORMATION struct{
DriveLetterWasAssigned bool //col:870
CurrentDriveLetter uint8 //col:871
}


type  _MOUNTMGR_VOLUME_MOUNT_POINT struct{
SourceVolumeNameOffset uint16 //col:877
SourceVolumeNameLength uint16 //col:878
TargetVolumeNameOffset uint16 //col:879
TargetVolumeNameLength uint16 //col:880
}


type  _MOUNTMGR_CHANGE_NOTIFY_INFO struct{
EpicNumber uint32 //col:881
}


type  _MOUNTMGR_TARGET_NAME struct{
DeviceNameLength uint16 //col:886
DeviceName[1] WCHAR //col:887
}


type  _MOUNTDEV_NAME struct{
NameLength uint16 //col:891
Name[1] WCHAR //col:892
}


type  _MOUNTMGR_VOLUME_PATHS struct{
MultiSzLength uint32 //col:896
MultiSz[1] WCHAR //col:897
}



type (
Ntioapi interface{
typedef_VOID_()(ok bool)//col:9
}
ntioapi struct{}
)

func NewNtioapi()Ntioapi{ return & ntioapi{} }

func (n *ntioapi)typedef_VOID_()(ok bool){//col:9
/*typedef VOID (NTAPI *PIO_APC_ROUTINE)(
    _In_ PVOID ApcContext,
    _In_ PIO_STATUS_BLOCK IoStatusBlock,
    _In_ ULONG Reserved
    );
typedef DECLSPEC_ALIGN(8) struct _FILE_IO_PRIORITY_HINT_INFORMATION
{
    IO_PRIORITY_HINT PriorityHint;
} FILE_IO_PRIORITY_HINT_INFORMATION, *PFILE_IO_PRIORITY_HINT_INFORMATION;*/
return true
}



