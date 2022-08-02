package phnt


const(
    TableEmptyTree = 1  //col:3
    TableFoundNode = 2  //col:4
    TableInsertAsLeft = 3  //col:5
    TableInsertAsRight = 4  //col:6
)


const(
    GenericLessThan = 1  //col:10
    GenericGreaterThan = 2  //col:11
    GenericEqual = 3  //col:12
)


const(
    NormOther  =  0x0  //col:16
    NormC  =  0x1  //col:17
    NormD  =  0x2  //col:18
    NormKC  =  0x5  //col:19
    NormKD  =  0x6  //col:20
    NormIdna  =  0xd  //col:21
    DisallowUnassigned  =  0x100  //col:22
    NormCDisallowUnassigned  =  0x101  //col:23
    NormDDisallowUnassigned  =  0x102  //col:24
    NormKCDisallowUnassigned  =  0x105  //col:25
    NormKDDisallowUnassigned  =  0x106  //col:26
    NormIdnaDisallowUnassigned  =  0x10d  //col:27
)


const(
    RF_SORTED = 1  //col:31
    RF_UNSORTED = 2  //col:32
    RF_CALLBACK = 3  //col:33
    RF_KERNEL_DYNAMIC = 4  //col:34
)


const(
    RtlPathTypeUnknown = 1  //col:38
    RtlPathTypeUncAbsolute = 2  //col:39
    RtlPathTypeDriveAbsolute = 3  //col:40
    RtlPathTypeDriveRelative = 4  //col:41
    RtlPathTypeRooted = 5  //col:42
    RtlPathTypeRelative = 6  //col:43
    RtlPathTypeLocalDevice = 7  //col:44
    RtlPathTypeRootLocalDevice = 8  //col:45
)


const(
    HEAP_COMPATIBILITY_STANDARD  =  0  //col:49
    HEAP_COMPATIBILITY_LAL  =  1  //col:50
    HEAP_COMPATIBILITY_LFH  =  2  //col:51
)


const(
    ImageDepPolicy  = 1  //col:55
    ImageAslrPolicy  = 2  //col:56
    ImageDynamicCodePolicy  = 3  //col:57
    ImageStrictHandleCheckPolicy  = 4  //col:58
    ImageSystemCallDisablePolicy  = 5  //col:59
    ImageMitigationOptionsMask = 6  //col:60
    ImageExtensionPointDisablePolicy  = 7  //col:61
    ImageControlFlowGuardPolicy  = 8  //col:62
    ImageSignaturePolicy  = 9  //col:63
    ImageFontDisablePolicy  = 10  //col:64
    ImageImageLoadPolicy  = 11  //col:65
    ImagePayloadRestrictionPolicy  = 12  //col:66
    ImageChildProcessPolicy  = 13  //col:67
    ImageSehopPolicy  = 14  //col:68
    ImageHeapPolicy  = 15  //col:69
    ImageUserShadowStackPolicy  = 16  //col:70
    MaxImageMitigationPolicy = 17  //col:71
)


const(
    RtlMitigationOptionStateNotConfigured = 1  //col:75
    RtlMitigationOptionStateOn = 2  //col:76
    RtlMitigationOptionStateOff = 3  //col:77
    RtlMitigationOptionStateForce = 4  //col:78
    RtlMitigationOptionStateOption = 5  //col:79
)


const(
    NotAppContainerSidType = 1  //col:83
    ChildAppContainerSidType = 2  //col:84
    ParentAppContainerSidType = 3  //col:85
    InvalidAppContainerSidType = 4  //col:86
    MaxAppContainerSidType = 5  //col:87
)


const(
    LocationTypeRegistry = 1  //col:91
    LocationTypeFileSystem = 2  //col:92
    LocationTypeMaximum = 3  //col:93
)


const(
    RtlBsdItemVersionNumber  = 1  //col:97
    RtlBsdItemProductType  = 2  //col:98
    RtlBsdItemAabEnabled  = 3  //col:99
    RtlBsdItemAabTimeout  = 4  //col:100
    RtlBsdItemBootGood  = 5  //col:101
    RtlBsdItemBootShutdown  = 6  //col:102
    RtlBsdSleepInProgress  = 7  //col:103
    RtlBsdPowerTransition  = 8  //col:104
    RtlBsdItemBootAttemptCount  = 9  //col:105
    RtlBsdItemBootCheckpoint  = 10  //col:106
    RtlBsdItemBootId  = 11  //col:107
    RtlBsdItemShutdownBootId  = 12  //col:108
    RtlBsdItemReportedAbnormalShutdownBootId  = 13  //col:109
    RtlBsdItemErrorInfo  = 14  //col:110
    RtlBsdItemPowerButtonPressInfo  = 15  //col:111
    RtlBsdItemChecksum  = 16  //col:112
    RtlBsdPowerTransitionExtension = 17  //col:113
    RtlBsdItemFeatureConfigurationState  = 18  //col:114
    RtlBsdItemMax = 19  //col:115
)


const(
    RtlFeatureConfigurationBoot = 1  //col:119
    RtlFeatureConfigurationRuntime = 2  //col:120
    RtlFeatureConfigurationCount = 3  //col:121
)



type RTL_BALANCED_LINKS struct{
struct // //col:3
struct // //col:4
struct // //col:5
Balance int8 //col:6
Reserved[3] uint8 //col:7
}



type RTL_AVL_TABLE struct{
BalancedRoot RTL_BALANCED_LINKS //col:11
OrderedPointer PVOID //col:12
WhichOrderedElement uint32 //col:13
NumberGenericTableElements uint32 //col:14
DepthOfTree uint32 //col:15
RestartKey PRTL_BALANCED_LINKS //col:16
DeleteCount uint32 //col:17
CompareRoutine PRTL_AVL_COMPARE_ROUTINE //col:18
AllocateRoutine PRTL_AVL_ALLOCATE_ROUTINE //col:19
FreeRoutine PRTL_AVL_FREE_ROUTINE //col:20
TableContext PVOID //col:21
}



type RTL_SPLAY_LINKS struct{
struct // //col:25
struct // //col:26
struct // //col:27
}



type RTL_GENERIC_TABLE struct{
TableRoot PRTL_SPLAY_LINKS //col:31
InsertOrderList *list.List //col:32
OrderedPointer PLIST_ENTRY //col:33
WhichOrderedElement uint32 //col:34
NumberGenericTableElements uint32 //col:35
CompareRoutine PRTL_GENERIC_COMPARE_ROUTINE //col:36
AllocateRoutine PRTL_GENERIC_ALLOCATE_ROUTINE //col:37
FreeRoutine PRTL_GENERIC_FREE_ROUTINE //col:38
TableContext PVOID //col:39
}



type RTL_RB_TREE struct{
Root PRTL_BALANCED_NODE //col:43
Min PRTL_BALANCED_NODE //col:44
}



type RTL_DYNAMIC_HASH_TABLE_ENTRY struct{
Linkage *list.List //col:48
Signature ULONG_PTR //col:49
}



type RTL_DYNAMIC_HASH_TABLE_CONTEXT struct{
ChainHead PLIST_ENTRY //col:53
PrevLinkage PLIST_ENTRY //col:54
Signature ULONG_PTR //col:55
}



type RTL_DYNAMIC_HASH_TABLE_ENUMERATOR struct{
HashEntry RTL_DYNAMIC_HASH_TABLE_ENTRY //col:59
ChainHead PLIST_ENTRY //col:60
BucketIndex uint32 //col:61
}



type RTL_DYNAMIC_HASH_TABLE struct{
Flags uint32 //col:65
Shift uint32 //col:66
TableSize uint32 //col:67
Pivot uint32 //col:68
DivisorMask uint32 //col:69
NumEntries uint32 //col:70
NonEmptyBuckets uint32 //col:71
NumEnumerators uint32 //col:72
Directory PVOID //col:73
}



type RTL_RESOURCE struct{
CriticalSection RTL_CRITICAL_SECTION //col:77
SharedSemaphore HANDLE //col:78
ULONG volatile //col:79
ExclusiveSemaphore HANDLE //col:80
ULONG volatile //col:81
LONG volatile //col:82
ExclusiveOwnerThread HANDLE //col:83
Flags uint32 //col:84
DebugInfo PRTL_RESOURCE_DEBUG //col:85
}



type PREFIX_TABLE_ENTRY struct{
NodeTypeCode CSHORT //col:89
NameLength CSHORT //col:90
struct // //col:91
Links RTL_SPLAY_LINKS //col:92
Prefix PSTRING //col:93
}



type PREFIX_TABLE struct{
NodeTypeCode CSHORT //col:97
NameLength CSHORT //col:98
NextPrefixTree PPREFIX_TABLE_ENTRY //col:99
}



type UNICODE_PREFIX_TABLE_ENTRY struct{
NodeTypeCode CSHORT //col:103
NameLength CSHORT //col:104
struct // //col:105
struct // //col:106
Links RTL_SPLAY_LINKS //col:107
Prefix PUNICODE_STRING //col:108
}



type UNICODE_PREFIX_TABLE struct{
NodeTypeCode CSHORT //col:112
NameLength CSHORT //col:113
NextPrefixTree PUNICODE_PREFIX_TABLE_ENTRY //col:114
LastNextEntry PUNICODE_PREFIX_TABLE_ENTRY //col:115
}



type COMPRESSED_DATA_INFO struct{
CompressionFormatAndEngine USHORT //col:119
CompressionUnitShift uint8 //col:120
ChunkShift uint8 //col:121
ClusterShift uint8 //col:122
Reserved uint8 //col:123
NumberOfChunks USHORT //col:124
CompressedChunkSizes[1] uint32 //col:125
}



type CURDIR struct{
DosPath UNICODE_STRING //col:129
Handle HANDLE //col:130
}



type RTL_DRIVE_LETTER_CURDIR struct{
Flags USHORT //col:134
Length USHORT //col:135
TimeStamp uint32 //col:136
DosPath STRING //col:137
}



type RTL_USER_PROCESS_PARAMETERS struct{
MaximumLength uint32 //col:141
Length uint32 //col:142
Flags uint32 //col:143
DebugFlags uint32 //col:144
ConsoleHandle HANDLE //col:145
ConsoleFlags uint32 //col:146
StandardInput HANDLE //col:147
StandardOutput HANDLE //col:148
StandardError HANDLE //col:149
CurrentDirectory CURDIR //col:150
DllPath UNICODE_STRING //col:151
ImagePathName UNICODE_STRING //col:152
CommandLine UNICODE_STRING //col:153
Environment PVOID //col:154
StartingX uint32 //col:155
StartingY uint32 //col:156
CountX uint32 //col:157
CountY uint32 //col:158
CountCharsX uint32 //col:159
CountCharsY uint32 //col:160
FillAttribute uint32 //col:161
WindowFlags uint32 //col:162
ShowWindowFlags uint32 //col:163
WindowTitle UNICODE_STRING //col:164
DesktopInfo UNICODE_STRING //col:165
ShellInfo UNICODE_STRING //col:166
RuntimeData UNICODE_STRING //col:167
CurrentDirectories[RTL_MAX_DRIVE_LETTERS] RTL_DRIVE_LETTER_CURDIR //col:168
EnvironmentSize ULONG_PTR //col:169
EnvironmentVersion ULONG_PTR //col:170
PackageDependencyData PVOID //col:171
ProcessGroupId uint32 //col:172
LoaderThreads uint32 //col:173
RedirectionDllName UNICODE_STRING //col:174
HeapPartitionName UNICODE_STRING //col:175
DefaultThreadpoolCpuSetMasks ULONG_PTR //col:176
DefaultThreadpoolCpuSetMaskCount uint32 //col:177
}



type RTL_USER_PROCESS_INFORMATION struct{
Length uint32 //col:181
ProcessHandle HANDLE //col:182
ThreadHandle HANDLE //col:183
ClientId CLIENT_ID //col:184
ImageInformation SECTION_IMAGE_INFORMATION //col:185
}



type RTL_USER_PROCESS_EXTENDED_PARAMETERS struct{
Version USHORT //col:189
NodeNumber USHORT //col:190
ProcessSecurityDescriptor PSECURITY_DESCRIPTOR //col:191
ThreadSecurityDescriptor PSECURITY_DESCRIPTOR //col:192
ParentProcess HANDLE //col:193
DebugPort HANDLE //col:194
TokenHandle HANDLE //col:195
JobHandle HANDLE //col:196
}



type RTLP_PROCESS_REFLECTION_REFLECTION_INFORMATION struct{
ReflectionProcessHandle HANDLE //col:200
ReflectionThreadHandle HANDLE //col:201
ReflectionClientId CLIENT_ID //col:202
}



type CONTEXT_CHUNK  struct{
Offset LONG //col:206
Length uint32 //col:207
}



type CONTEXT_EX  struct{
All CONTEXT_CHUNK //col:211
Legacy CONTEXT_CHUNK //col:212
XState CONTEXT_CHUNK //col:213
}



type DYNAMIC_FUNCTION_TABLE struct{
ListEntry *list.List //col:217
FunctionTable PRUNTIME_FUNCTION //col:218
TimeStamp LARGE_INTEGER //col:219
MinimumAddress ULONG64 //col:220
MaximumAddress ULONG64 //col:221
BaseAddress ULONG64 //col:222
Callback PGET_RUNTIME_FUNCTION_CALLBACK //col:223
Context PVOID //col:224
OutOfProcessCallbackDll PWSTR //col:225
Type FUNCTION_TABLE_TYPE //col:226
EntryCount uint32 //col:227
TreeNodeMin RTL_BALANCED_NODE //col:228
TreeNodeMax RTL_BALANCED_NODE //col:229
}



type RTLP_CURDIR_REF struct{
ReferenceCount LONG //col:233
DirectoryHandle HANDLE //col:234
}



type RTL_RELATIVE_NAME_U struct{
RelativeName UNICODE_STRING //col:238
ContainingDirectory HANDLE //col:239
CurDirRef PRTLP_CURDIR_REF //col:240
}



type GENERATE_NAME_CONTEXT struct{
Checksum USHORT //col:244
CheckSumInserted bool //col:245
NameLength uint8 //col:246
NameBuffer[8] WCHAR //col:247
ExtensionLength uint32 //col:248
ExtensionBuffer[4] WCHAR //col:249
LastIndexValue uint32 //col:250
}



type RTL_HEAP_ENTRY struct{
Size SIZE_T //col:254
Flags USHORT //col:255
AllocatorBackTraceIndex USHORT //col:256
Union union //col:257
Struct struct //col:259
Settable SIZE_T //col:261
Tag uint32 //col:262
}



type RTL_HEAP_TAG struct{
NumberOfAllocations uint32 //col:273
NumberOfFrees uint32 //col:274
BytesAllocated SIZE_T //col:275
TagIndex USHORT //col:276
CreatorBackTraceIndex USHORT //col:277
TagName[24] WCHAR //col:278
}



type RTL_HEAP_INFORMATION struct{
BaseAddress PVOID //col:282
Flags uint32 //col:283
EntryOverhead USHORT //col:284
CreatorBackTraceIndex USHORT //col:285
BytesAllocated SIZE_T //col:286
BytesCommitted SIZE_T //col:287
NumberOfTags uint32 //col:288
NumberOfEntries uint32 //col:289
NumberOfPseudoTags uint32 //col:290
PseudoTagGranularity uint32 //col:291
Reserved[5] uint32 //col:292
Tags PRTL_HEAP_TAG //col:293
Entries PRTL_HEAP_ENTRY //col:294
HeapTag ULONG64 //col:295
}



type RTL_PROCESS_HEAPS struct{
NumberOfHeaps uint32 //col:299
Heaps[1] RTL_HEAP_INFORMATION //col:300
}



type RTL_HEAP_PARAMETERS struct{
Length uint32 //col:304
SegmentReserve SIZE_T //col:305
SegmentCommit SIZE_T //col:306
DeCommitFreeBlockThreshold SIZE_T //col:307
DeCommitTotalFreeThreshold SIZE_T //col:308
MaximumAllocationSize SIZE_T //col:309
VirtualMemoryThreshold SIZE_T //col:310
InitialCommit SIZE_T //col:311
InitialReserve SIZE_T //col:312
CommitRoutine PRTL_HEAP_COMMIT_ROUTINE //col:313
Reserved[2] SIZE_T //col:314
}



type RTL_HEAP_TAG_INFO struct{
NumberOfAllocations uint32 //col:318
NumberOfFrees uint32 //col:319
BytesAllocated SIZE_T //col:320
}



type RTL_HEAP_USAGE_ENTRY struct{
struct // //col:324
Address PVOID //col:325
Size SIZE_T //col:326
AllocatorBackTraceIndex USHORT //col:327
TagIndex USHORT //col:328
}



type RTL_HEAP_USAGE struct{
Length uint32 //col:332
BytesAllocated SIZE_T //col:333
BytesCommitted SIZE_T //col:334
BytesReserved SIZE_T //col:335
BytesReservedMaximum SIZE_T //col:336
Entries PRTL_HEAP_USAGE_ENTRY //col:337
AddedEntries PRTL_HEAP_USAGE_ENTRY //col:338
RemovedEntries PRTL_HEAP_USAGE_ENTRY //col:339
Reserved[8] ULONG_PTR //col:340
}



type RTL_HEAP_WALK_ENTRY struct{
DataAddress PVOID //col:344
DataSize SIZE_T //col:345
OverheadBytes uint8 //col:346
SegmentIndex uint8 //col:347
Flags USHORT //col:348
Union union //col:349
Struct struct //col:351
Settable SIZE_T //col:353
TagIndex USHORT //col:354
AllocatorBackTraceIndex USHORT //col:355
Reserved[2] uint32 //col:356
}



type PROCESS_HEAP_INFORMATION struct{
ReserveSize ULONG_PTR //col:369
CommitSize ULONG_PTR //col:370
NumberOfHeaps uint32 //col:371
FirstHeapInformationOffset ULONG_PTR //col:372
}



type HEAP_INFORMATION struct{
Address ULONG_PTR //col:376
Mode uint32 //col:377
ReserveSize ULONG_PTR //col:378
CommitSize ULONG_PTR //col:379
FirstRegionInformationOffset ULONG_PTR //col:380
NextHeapInformationOffset ULONG_PTR //col:381
}



type HEAP_EXTENDED_INFORMATION struct{
Process HANDLE //col:385
Heap ULONG_PTR //col:386
Level uint32 //col:387
CallbackRoutine PVOID //col:388
CallbackContext PVOID //col:389
Union union //col:390
ProcessHeapInformation PROCESS_HEAP_INFORMATION //col:392
HeapInformation HEAP_INFORMATION //col:393
}



type HEAP_DEBUGGING_INFORMATION struct{
InterceptorFunction PVOID //col:398
InterceptorValue USHORT //col:399
ExtendedOptions uint32 //col:400
StackTraceDepth uint32 //col:401
MinTotalBlockSize SIZE_T //col:402
MaxTotalBlockSize SIZE_T //col:403
HeapLeakEnumerationRoutine PRTL_HEAP_LEAK_ENUMERATION_ROUTINE //col:404
}



type RTL_MEMORY_ZONE_SEGMENT struct{
struct // //col:408
Size SIZE_T //col:409
Next PVOID //col:410
Limit PVOID //col:411
}



type RTL_MEMORY_ZONE struct{
Segment RTL_MEMORY_ZONE_SEGMENT //col:415
Lock RTL_SRWLOCK //col:416
LockCount uint32 //col:417
FirstSegment PRTL_MEMORY_ZONE_SEGMENT //col:418
}



type RTL_PROCESS_VERIFIER_OPTIONS struct{
SizeStruct uint32 //col:422
Option uint32 //col:423
OptionData[1] uint8 //col:424
}



type RTL_DEBUG_INFORMATION struct{
SectionHandleClient HANDLE //col:428
ViewBaseClient PVOID //col:429
ViewBaseTarget PVOID //col:430
ViewBaseDelta ULONG_PTR //col:431
EventPairClient HANDLE //col:432
EventPairTarget HANDLE //col:433
TargetProcessId HANDLE //col:434
TargetThreadHandle HANDLE //col:435
Flags uint32 //col:436
OffsetFree SIZE_T //col:437
CommitSize SIZE_T //col:438
ViewSize SIZE_T //col:439
Union union //col:440
struct // //col:442
struct // //col:443
}



type PARSE_MESSAGE_CONTEXT struct{
fFlags uint32 //col:458
cwSavColumn uint32 //col:459
iwSrc SIZE_T //col:460
iwDst SIZE_T //col:461
iwDstSpace SIZE_T //col:462
lpvArgStart va_list //col:463
}



type TIME_FIELDS struct{
Year CSHORT //col:467
Month CSHORT //col:468
Day CSHORT //col:469
Hour CSHORT //col:470
Minute CSHORT //col:471
Second CSHORT //col:472
Milliseconds CSHORT //col:473
Weekday CSHORT //col:474
}



type RTL_TIME_ZONE_INFORMATION struct{
Bias LONG //col:478
StandardName[32] WCHAR //col:479
StandardStart TIME_FIELDS //col:480
StandardBias LONG //col:481
DaylightName[32] WCHAR //col:482
DaylightStart TIME_FIELDS //col:483
DaylightBias LONG //col:484
}



type RTL_BITMAP struct{
SizeOfBitMap uint32 //col:488
Buffer PULONG //col:489
}



type RTL_BITMAP_RUN struct{
StartingIndex uint32 //col:493
NumberOfBits uint32 //col:494
}



type RTL_BITMAP_EX struct{
SizeOfBitMap ULONG64 //col:498
Buffer PULONG64 //col:499
}



type RTL_HANDLE_TABLE_ENTRY struct{
Union union //col:503
Flags uint32 //col:505
struct // //col:506
}



type RTL_HANDLE_TABLE struct{
MaximumNumberOfHandles uint32 //col:511
SizeOfHandleTableEntry uint32 //col:512
Reserved[2] uint32 //col:513
FreeHandles PRTL_HANDLE_TABLE_ENTRY //col:514
CommittedHandles PRTL_HANDLE_TABLE_ENTRY //col:515
UnCommittedHandles PRTL_HANDLE_TABLE_ENTRY //col:516
MaxReservedHandles PRTL_HANDLE_TABLE_ENTRY //col:517
}



type RTL_ACE_DATA struct{
AceType uint8 //col:521
InheritFlags uint8 //col:522
AceFlags uint8 //col:523
AccessMask ACCESS_MASK //col:524
Sid PSID* //col:525
}



type RTL_QUERY_REGISTRY_TABLE struct{
QueryRoutine PRTL_QUERY_REGISTRY_ROUTINE //col:529
Flags uint32 //col:530
Name PWSTR //col:531
EntryContext PVOID //col:532
DefaultType uint32 //col:533
DefaultData PVOID //col:534
DefaultLength uint32 //col:535
}



type RTL_UNLOAD_EVENT_TRACE struct{
BaseAddress PVOID //col:539
SizeOfImage SIZE_T //col:540
Sequence uint32 //col:541
TimeDateStamp uint32 //col:542
CheckSum uint32 //col:543
ImageName[32] WCHAR //col:544
Version[2] uint32 //col:545
}



type RTL_UNLOAD_EVENT_TRACE32  struct{
BaseAddress uint32 //col:549
SizeOfImage uint32 //col:550
Sequence uint32 //col:551
TimeDateStamp uint32 //col:552
CheckSum uint32 //col:553
ImageName[32] WCHAR //col:554
Version[2] uint32 //col:555
}



type RTL_IMAGE_MITIGATION_DEP_POLICY struct{
Dep RTL_IMAGE_MITIGATION_POLICY //col:559
}



type RTL_IMAGE_MITIGATION_ASLR_POLICY struct{
ForceRelocateImages RTL_IMAGE_MITIGATION_POLICY //col:563
BottomUpRandomization RTL_IMAGE_MITIGATION_POLICY //col:564
HighEntropyRandomization RTL_IMAGE_MITIGATION_POLICY //col:565
}



type RTL_IMAGE_MITIGATION_DYNAMIC_CODE_POLICY struct{
BlockDynamicCode RTL_IMAGE_MITIGATION_POLICY //col:569
}



type RTL_IMAGE_MITIGATION_STRICT_HANDLE_CHECK_POLICY struct{
StrictHandleChecks RTL_IMAGE_MITIGATION_POLICY //col:573
}



type RTL_IMAGE_MITIGATION_SYSTEM_CALL_DISABLE_POLICY struct{
BlockWin32kSystemCalls RTL_IMAGE_MITIGATION_POLICY //col:577
}



type RTL_IMAGE_MITIGATION_EXTENSION_POINT_DISABLE_POLICY struct{
DisableExtensionPoints RTL_IMAGE_MITIGATION_POLICY //col:581
}



type RTL_IMAGE_MITIGATION_CONTROL_FLOW_GUARD_POLICY struct{
ControlFlowGuard RTL_IMAGE_MITIGATION_POLICY //col:585
StrictControlFlowGuard RTL_IMAGE_MITIGATION_POLICY //col:586
}



type RTL_IMAGE_MITIGATION_BINARY_SIGNATURE_POLICY struct{
BlockNonMicrosoftSignedBinaries RTL_IMAGE_MITIGATION_POLICY //col:590
EnforceSigningOnModuleDependencies RTL_IMAGE_MITIGATION_POLICY //col:591
}



type RTL_IMAGE_MITIGATION_FONT_DISABLE_POLICY struct{
DisableNonSystemFonts RTL_IMAGE_MITIGATION_POLICY //col:595
}



type RTL_IMAGE_MITIGATION_IMAGE_LOAD_POLICY struct{
BlockRemoteImageLoads RTL_IMAGE_MITIGATION_POLICY //col:599
BlockLowLabelImageLoads RTL_IMAGE_MITIGATION_POLICY //col:600
PreferSystem32 RTL_IMAGE_MITIGATION_POLICY //col:601
}



type RTL_IMAGE_MITIGATION_PAYLOAD_RESTRICTION_POLICY struct{
EnableExportAddressFilter RTL_IMAGE_MITIGATION_POLICY //col:605
EnableExportAddressFilterPlus RTL_IMAGE_MITIGATION_POLICY //col:606
EnableImportAddressFilter RTL_IMAGE_MITIGATION_POLICY //col:607
EnableRopStackPivot RTL_IMAGE_MITIGATION_POLICY //col:608
EnableRopCallerCheck RTL_IMAGE_MITIGATION_POLICY //col:609
EnableRopSimExec RTL_IMAGE_MITIGATION_POLICY //col:610
EafPlusModuleList[512] WCHAR //col:611
}



type RTL_IMAGE_MITIGATION_CHILD_PROCESS_POLICY struct{
DisallowChildProcessCreation RTL_IMAGE_MITIGATION_POLICY //col:615
}



type RTL_IMAGE_MITIGATION_SEHOP_POLICY struct{
Sehop RTL_IMAGE_MITIGATION_POLICY //col:619
}



type RTL_IMAGE_MITIGATION_HEAP_POLICY struct{
TerminateOnHeapErrors RTL_IMAGE_MITIGATION_POLICY //col:623
}



type RTL_IMAGE_MITIGATION_USER_SHADOW_STACK_POLICY struct{
UserShadowStack RTL_IMAGE_MITIGATION_POLICY //col:627
SetContextIpValidation RTL_IMAGE_MITIGATION_POLICY //col:628
BlockNonCetBinaries RTL_IMAGE_MITIGATION_POLICY //col:629
}



type PS_PKG_CLAIM struct{
Flags uint32 //col:633
Origin uint32 //col:634
}



type RTL_BSD_DATA_POWER_TRANSITION struct{
PowerButtonTimestamp LARGE_INTEGER //col:638
Struct struct //col:639
SystemRunning bool //col:641
ConnectedStandbyInProgress bool //col:642
UserShutdownInProgress bool //col:643
SystemShutdownInProgress bool //col:644
SleepInProgress bool //col:645
}



type RTL_BSD_DATA_ERROR_INFO struct{
BootId uint32 //col:657
RepeatCount uint32 //col:658
OtherErrorCount uint32 //col:659
Code uint32 //col:660
OtherErrorCount2 uint32 //col:661
}



type RTL_BSD_POWER_BUTTON_PRESS_INFO struct{
LastPressTime LARGE_INTEGER //col:665
CumulativePressCount uint32 //col:666
LastPressBootId USHORT //col:667
LastPowerWatchdogStage uint8 //col:668
Struct struct //col:669
WatchdogArmed uint8 //col:671
ShutdownInProgress uint8 //col:672
}



type RTL_BSD_ITEM struct{
Type RTL_BSD_ITEM_TYPE //col:685
DataBuffer PVOID //col:686
DataLength uint32 //col:687
}



type _RTL_FEATURE_USAGE_REPORT struct{
FeatureId uint32 //col:691
ReportingKind USHORT //col:692
ReportingOptions USHORT //col:693
}



type RTL_FEATURE_CONFIGURATION struct{
FeatureId uint32 //col:697
Union union //col:698
Flags uint32 //col:700
Struct struct //col:701
Priority uint32 //col:703
EnabledState uint32 //col:704
IsWexpConfiguration uint32 //col:705
HasSubscriptions uint32 //col:706
Variant uint32 //col:707
VariantPayloadKind uint32 //col:708
Reserved uint32 //col:709
}




type (
Ntrtl interface{
FORCEINLINE_VOID_InitializeListHead()(ok bool)//col:6
_Check_return__FORCEINLINE_BOOLEAN_IsListEmpty()(ok bool)//col:12
FORCEINLINE_BOOLEAN_RemoveEntryList()(ok bool)//col:24
FORCEINLINE_PLIST_ENTRY_RemoveHeadList()(ok bool)//col:36
FORCEINLINE_PLIST_ENTRY_RemoveTailList()(ok bool)//col:48
FORCEINLINE_VOID_InsertTailList()(ok bool)//col:60
FORCEINLINE_VOID_InsertHeadList()(ok bool)//col:72
FORCEINLINE_VOID_AppendTailList()(ok bool)//col:83
FORCEINLINE_PSINGLE_LIST_ENTRY_PopEntryList()(ok bool)//col:93
FORCEINLINE_VOID_PushEntryList()(ok bool)//col:101
typedef_RTL_GENERIC_COMPARE_RESULTS_()(ok bool)//col:291
RtlInitHashTableContextFromEnumerator()(ok bool)//col:299
RtlReleaseHashTableContext()(ok bool)//col:310
RtlNonEmptyBucketsHashTable()(ok bool)//col:316
RtlEmptyBucketsHashTable()(ok bool)//col:322
RtlTotalEntriesHashTable()(ok bool)//col:328
RtlActiveEnumeratorsHashTable()(ok bool)//col:334
RtlCreateHashTable()(ok bool)//col:553
RtlInitString()(ok bool)//col:570
RtlInitAnsiString()(ok bool)//col:635
RtlInitEmptyUnicodeString()(ok bool)//col:651
RtlInitUnicodeString()(ok bool)//col:1526
RtlFillMemoryUlong()(ok bool)//col:2001
FORCEINLINE_BOOLEAN_RtlIsZeroLuid()(ok bool)//col:2007
FORCEINLINE_LUID_RtlConvertLongToLuid()(ok bool)//col:2018
FORCEINLINE_LUID_RtlConvertUlongToLuid()(ok bool)//col:2027
RtlCopyLuid()(ok bool)//col:2430
RtlNumberOfClearBits()(ok bool)//col:2839
RtlAreAnyAccessesGranted()(ok bool)//col:2846
}

ntrtl struct{}
)

func NewNtrtl()Ntrtl{ return & ntrtl{} }

func (n *ntrtl)FORCEINLINE_VOID_InitializeListHead()(ok bool){//col:6






return true
}


func (n *ntrtl)_Check_return__FORCEINLINE_BOOLEAN_IsListEmpty()(ok bool){//col:12






return true
}


func (n *ntrtl)FORCEINLINE_BOOLEAN_RemoveEntryList()(ok bool){//col:24












return true
}


func (n *ntrtl)FORCEINLINE_PLIST_ENTRY_RemoveHeadList()(ok bool){//col:36












return true
}


func (n *ntrtl)FORCEINLINE_PLIST_ENTRY_RemoveTailList()(ok bool){//col:48












return true
}


func (n *ntrtl)FORCEINLINE_VOID_InsertTailList()(ok bool){//col:60












return true
}


func (n *ntrtl)FORCEINLINE_VOID_InsertHeadList()(ok bool){//col:72












return true
}


func (n *ntrtl)FORCEINLINE_VOID_AppendTailList()(ok bool){//col:83











return true
}


func (n *ntrtl)FORCEINLINE_PSINGLE_LIST_ENTRY_PopEntryList()(ok bool){//col:93










return true
}


func (n *ntrtl)FORCEINLINE_VOID_PushEntryList()(ok bool){//col:101








return true
}


func (n *ntrtl)typedef_RTL_GENERIC_COMPARE_RESULTS_()(ok bool){//col:291






























































































































































































return true
}


func (n *ntrtl)RtlInitHashTableContextFromEnumerator()(ok bool){//col:299








return true
}


func (n *ntrtl)RtlReleaseHashTableContext()(ok bool){//col:310











return true
}


func (n *ntrtl)RtlNonEmptyBucketsHashTable()(ok bool){//col:316






return true
}


func (n *ntrtl)RtlEmptyBucketsHashTable()(ok bool){//col:322






return true
}


func (n *ntrtl)RtlTotalEntriesHashTable()(ok bool){//col:328






return true
}


func (n *ntrtl)RtlActiveEnumeratorsHashTable()(ok bool){//col:334






return true
}


func (n *ntrtl)RtlCreateHashTable()(ok bool){//col:553



























































































































































































































return true
}


func (n *ntrtl)RtlInitString()(ok bool){//col:570

















return true
}


func (n *ntrtl)RtlInitAnsiString()(ok bool){//col:635

































































return true
}


func (n *ntrtl)RtlInitEmptyUnicodeString()(ok bool){//col:651
















return true
}


func (n *ntrtl)RtlInitUnicodeString()(ok bool){//col:1526











































































































































































































































































































































































































































































































































































































































































































































































































































































































return true
}


func (n *ntrtl)RtlFillMemoryUlong()(ok bool){//col:2001



























































































































































































































































































































































































































































































return true
}


func (n *ntrtl)FORCEINLINE_BOOLEAN_RtlIsZeroLuid()(ok bool){//col:2007






return true
}


func (n *ntrtl)FORCEINLINE_LUID_RtlConvertLongToLuid()(ok bool){//col:2018











return true
}


func (n *ntrtl)FORCEINLINE_LUID_RtlConvertUlongToLuid()(ok bool){//col:2027









return true
}


func (n *ntrtl)RtlCopyLuid()(ok bool){//col:2430



















































































































































































































































































































































































































return true
}


func (n *ntrtl)RtlNumberOfClearBits()(ok bool){//col:2839

























































































































































































































































































































































































































return true
}


func (n *ntrtl)RtlAreAnyAccessesGranted()(ok bool){//col:2846







return true
}




