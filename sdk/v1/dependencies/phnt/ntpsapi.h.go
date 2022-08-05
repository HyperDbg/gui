package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpsapi.h.back

const(
    ProcessBasicInformation  = 1  //col:3
    ProcessQuotaLimits  = 2  //col:4
    ProcessIoCounters  = 3  //col:5
    ProcessVmCounters  = 4  //col:6
    ProcessTimes  = 5  //col:7
    ProcessBasePriority  = 6  //col:8
    ProcessRaisePriority  = 7  //col:9
    ProcessDebugPort  = 8  //col:10
    ProcessExceptionPort  = 9  //col:11
    ProcessAccessToken  = 10  //col:12
    ProcessLdtInformation  = 11  //col:13
    ProcessLdtSize  = 12  //col:14
    ProcessDefaultHardErrorMode  = 13  //col:15
    ProcessIoPortHandlers  = 14  //col:16
    ProcessPooledUsageAndLimits  = 15  //col:17
    ProcessWorkingSetWatch  = 16  //col:18
    ProcessUserModeIOPL  = 17  //col:19
    ProcessEnableAlignmentFaultFixup  = 18  //col:20
    ProcessPriorityClass  = 19  //col:21
    ProcessWx86Information  = 20  //col:22
    ProcessHandleCount  = 21  //col:23
    ProcessAffinityMask  = 22  //col:24
    ProcessPriorityBoost  = 23  //col:25
    ProcessDeviceMap  = 24  //col:26
    ProcessSessionInformation  = 25  //col:27
    ProcessForegroundInformation  = 26  //col:28
    ProcessWow64Information  = 27  //col:29
    ProcessImageFileName  = 28  //col:30
    ProcessLUIDDeviceMapsEnabled  = 29  //col:31
    ProcessBreakOnTermination  = 30  //col:32
    ProcessDebugObjectHandle  = 31  //col:33
    ProcessDebugFlags  = 32  //col:34
    ProcessHandleTracing  = 33  //col:35
    ProcessIoPriority  = 34  //col:36
    ProcessExecuteFlags  = 35  //col:37
    ProcessTlsInformation  = 36  //col:38
    ProcessCookie  = 37  //col:39
    ProcessImageInformation  = 38  //col:40
    ProcessCycleTime  = 39  //col:41
    ProcessPagePriority  = 40  //col:42
    ProcessInstrumentationCallback  = 41  //col:43
    ProcessThreadStackAllocation  = 42  //col:44
    ProcessWorkingSetWatchEx  = 43  //col:45
    ProcessImageFileNameWin32  = 44  //col:46
    ProcessImageFileMapping  = 45  //col:47
    ProcessAffinityUpdateMode  = 46  //col:48
    ProcessMemoryAllocationMode  = 47  //col:49
    ProcessGroupInformation  = 48  //col:50
    ProcessTokenVirtualizationEnabled  = 49  //col:51
    ProcessConsoleHostProcess  = 50  //col:52
    ProcessWindowInformation  = 51  //col:53
    ProcessHandleInformation  = 52  //col:54
    ProcessMitigationPolicy  = 53  //col:55
    ProcessDynamicFunctionTableInformation = 54  //col:56
    ProcessHandleCheckingMode  = 55  //col:57
    ProcessKeepAliveCount  = 56  //col:58
    ProcessRevokeFileHandles  = 57  //col:59
    ProcessWorkingSetControl  = 58  //col:60
    ProcessHandleTable  = 59  //col:61
    ProcessCheckStackExtentsMode  = 60  //col:62
    ProcessCommandLineInformation  = 61  //col:63
    ProcessProtectionInformation  = 62  //col:64
    ProcessMemoryExhaustion  = 63  //col:65
    ProcessFaultInformation  = 64  //col:66
    ProcessTelemetryIdInformation  = 65  //col:67
    ProcessCommitReleaseInformation  = 66  //col:68
    ProcessDefaultCpuSetsInformation  = 67  //col:69
    ProcessAllowedCpuSetsInformation  = 68  //col:70
    ProcessSubsystemProcess = 69  //col:71
    ProcessJobMemoryInformation  = 70  //col:72
    ProcessInPrivate  = 71  //col:73
    ProcessRaiseUMExceptionOnInvalidHandleClose  = 72  //col:74
    ProcessIumChallengeResponse = 73  //col:75
    ProcessChildProcessInformation  = 74  //col:76
    ProcessHighGraphicsPriorityInformation  = 75  //col:77
    ProcessSubsystemInformation  = 76  //col:78
    ProcessEnergyValues  = 77  //col:79
    ProcessPowerThrottlingState  = 78  //col:80
    ProcessReserved3Information  = 79  //col:81
    ProcessWin32kSyscallFilterInformation  = 80  //col:82
    ProcessDisableSystemAllowedCpuSets  = 81  //col:83
    ProcessWakeInformation  = 82  //col:84
    ProcessEnergyTrackingState  = 83  //col:85
    ProcessManageWritesToExecutableMemory  = 84  //col:86
    ProcessCaptureTrustletLiveDump = 85  //col:87
    ProcessTelemetryCoverage = 86  //col:88
    ProcessEnclaveInformation = 87  //col:89
    ProcessEnableReadWriteVmLogging  = 88  //col:90
    ProcessUptimeInformation  = 89  //col:91
    ProcessImageSection  = 90  //col:92
    ProcessDebugAuthInformation  = 91  //col:93
    ProcessSystemResourceManagement  = 92  //col:94
    ProcessSequenceNumber  = 93  //col:95
    ProcessLoaderDetour  = 94  //col:96
    ProcessSecurityDomainInformation  = 95  //col:97
    ProcessCombineSecurityDomainsInformation  = 96  //col:98
    ProcessEnableLogging  = 97  //col:99
    ProcessLeapSecondInformation  = 98  //col:100
    ProcessFiberShadowStackAllocation  = 99  //col:101
    ProcessFreeFiberShadowStackAllocation  = 100  //col:102
    ProcessAltSystemCallInformation  = 101  //col:103
    ProcessDynamicEHContinuationTargets  = 102  //col:104
    ProcessDynamicEnforcedCetCompatibleRanges  = 103  //col:105
    ProcessCreateStateChange  = 104  //col:106
    ProcessApplyStateChange = 105  //col:107
    ProcessEnableOptionalXStateFeatures = 106  //col:108
    ProcessAltPrefetchParam  = 107  //col:109
    ProcessAssignCpuPartitions = 108  //col:110
    ProcessPriorityClassEx = 109  //col:111
    ProcessMembershipInformation = 110  //col:112
    ProcessEffectiveIoPriority = 111  //col:113
    ProcessEffectivePagePriority = 112  //col:114
    MaxProcessInfoClass = 113  //col:115
)


const(
    ThreadBasicInformation  = 1  //col:119
    ThreadTimes  = 2  //col:120
    ThreadPriority  = 3  //col:121
    ThreadBasePriority  = 4  //col:122
    ThreadAffinityMask  = 5  //col:123
    ThreadImpersonationToken  = 6  //col:124
    ThreadDescriptorTableEntry  = 7  //col:125
    ThreadEnableAlignmentFaultFixup  = 8  //col:126
    ThreadEventPair = 9  //col:127
    ThreadQuerySetWin32StartAddress  = 10  //col:128
    ThreadZeroTlsCell  = 11  //col:129
    ThreadPerformanceCount  = 12  //col:130
    ThreadAmILastThread  = 13  //col:131
    ThreadIdealProcessor  = 14  //col:132
    ThreadPriorityBoost  = 15  //col:133
    ThreadSetTlsArrayAddress  = 16  //col:134
    ThreadIsIoPending  = 17  //col:135
    ThreadHideFromDebugger  = 18  //col:136
    ThreadBreakOnTermination  = 19  //col:137
    ThreadSwitchLegacyState  = 20  //col:138
    ThreadIsTerminated  = 21  //col:139
    ThreadLastSystemCall  = 22  //col:140
    ThreadIoPriority  = 23  //col:141
    ThreadCycleTime  = 24  //col:142
    ThreadPagePriority  = 25  //col:143
    ThreadActualBasePriority  = 26  //col:144
    ThreadTebInformation  = 27  //col:145
    ThreadCSwitchMon = 28  //col:146
    ThreadCSwitchPmu = 29  //col:147
    ThreadWow64Context  = 30  //col:148
    ThreadGroupInformation  = 31  //col:149
    ThreadUmsInformation  = 32  //col:150
    ThreadCounterProfiling  = 33  //col:151
    ThreadIdealProcessorEx  = 34  //col:152
    ThreadCpuAccountingInformation  = 35  //col:153
    ThreadSuspendCount  = 36  //col:154
    ThreadHeterogeneousCpuPolicy  = 37  //col:155
    ThreadContainerId  = 38  //col:156
    ThreadNameInformation  = 39  //col:157
    ThreadSelectedCpuSets = 40  //col:158
    ThreadSystemThreadInformation  = 41  //col:159
    ThreadActualGroupAffinity  = 42  //col:160
    ThreadDynamicCodePolicyInfo  = 43  //col:161
    ThreadExplicitCaseSensitivity  = 44  //col:162
    ThreadWorkOnBehalfTicket  = 45  //col:163
    ThreadSubsystemInformation  = 46  //col:164
    ThreadDbgkWerReportActive  = 47  //col:165
    ThreadAttachContainer  = 48  //col:166
    ThreadManageWritesToExecutableMemory  = 49  //col:167
    ThreadPowerThrottlingState  = 50  //col:168
    ThreadWorkloadClass  = 51  //col:169
    ThreadCreateStateChange  = 52  //col:170
    ThreadApplyStateChange = 53  //col:171
    ThreadStrongerBadHandleChecks  = 54  //col:172
    ThreadEffectiveIoPriority = 55  //col:173
    ThreadEffectivePagePriority = 56  //col:174
    MaxThreadInfoClass = 57  //col:175
)


const(
    ProcessTlsReplaceIndex = 1  //col:179
    ProcessTlsReplaceVector = 2  //col:180
    MaxProcessTlsOperation = 3  //col:181
)


const(
    ProcessWorkingSetSwap = 1  //col:185
    ProcessWorkingSetEmpty = 2  //col:186
    ProcessWorkingSetOperationMax = 3  //col:187
)


const(
    PsProtectedTypeNone = 1  //col:191
    PsProtectedTypeProtectedLight = 2  //col:192
    PsProtectedTypeProtected = 3  //col:193
    PsProtectedTypeMax = 4  //col:194
)


const(
    PsProtectedSignerNone = 1  //col:198
    PsProtectedSignerAuthenticode = 2  //col:199
    PsProtectedSignerCodeGen = 3  //col:200
    PsProtectedSignerAntimalware = 4  //col:201
    PsProtectedSignerLsa = 5  //col:202
    PsProtectedSignerWindows = 6  //col:203
    PsProtectedSignerWinTcb = 7  //col:204
    PsProtectedSignerWinSystem = 8  //col:205
    PsProtectedSignerApp = 9  //col:206
    PsProtectedSignerMax = 10  //col:207
)


const(
    UmsInformationCommandInvalid = 1  //col:211
    UmsInformationCommandAttach = 2  //col:212
    UmsInformationCommandDetach = 3  //col:213
    UmsInformationCommandQuery = 4  //col:214
)


const(
    SubsystemInformationTypeWin32 = 1  //col:218
    SubsystemInformationTypeWSL = 2  //col:219
    MaxSubsystemInformationType = 3  //col:220
)


const(
    ThreadWorkloadClassDefault = 1  //col:224
    ThreadWorkloadClassGraphics = 2  //col:225
    MaxThreadWorkloadClass = 3  //col:226
)


const(
    ProcessStateChangeSuspend = 1  //col:230
    ProcessStateChangeResume = 2  //col:231
    ProcessStateChangeMax = 3  //col:232
)


const(
    ThreadStateChangeSuspend = 1  //col:236
    ThreadStateChangeResume = 2  //col:237
    ThreadStateChangeMax = 3  //col:238
)


const(
    QUEUE_USER_APC_FLAGS_NONE  =  0x0  //col:242
    QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC  =  0x1  //col:243
)


const(
    SeSafeOpenExperienceNone  =  0x00  //col:246
    SeSafeOpenExperienceCalled  =  0x01  //col:247
    SeSafeOpenExperienceAppRepCalled  =  0x02  //col:248
    SeSafeOpenExperiencePromptDisplayed  =  0x04  //col:249
    SeSafeOpenExperienceUAC  =  0x08  //col:250
    SeSafeOpenExperienceUninstaller  =  0x10  //col:251
    SeSafeOpenExperienceIgnoreUnknownOrBad  =  0x20  //col:252
    SeSafeOpenExperienceDefenderTrustedInstaller  =  0x40  //col:253
    SeSafeOpenExperienceMOTWPresent  =  0x80  //col:254
)


const(
    PsAttributeParentProcess  = 1  //col:258
    PsAttributeDebugObject  = 2  //col:259
    PsAttributeToken  = 3  //col:260
    PsAttributeClientId  = 4  //col:261
    PsAttributeTebAddress  = 5  //col:262
    PsAttributeImageName  = 6  //col:263
    PsAttributeImageInfo  = 7  //col:264
    PsAttributeMemoryReserve  = 8  //col:265
    PsAttributePriorityClass  = 9  //col:266
    PsAttributeErrorMode  = 10  //col:267
    PsAttributeStdHandleInfo  = 11  //col:268
    PsAttributeHandleList  = 12  //col:269
    PsAttributeGroupAffinity  = 13  //col:270
    PsAttributePreferredNode  = 14  //col:271
    PsAttributeIdealProcessor  = 15  //col:272
    PsAttributeUmsThread  = 16  //col:273
    PsAttributeMitigationOptions  = 17  //col:274
    PsAttributeProtectionLevel  = 18  //col:275
    PsAttributeSecureProcess  = 19  //col:276
    PsAttributeJobList  = 20  //col:277
    PsAttributeChildProcessPolicy  = 21  //col:278
    PsAttributeAllApplicationPackagesPolicy  = 22  //col:279
    PsAttributeWin32kFilter  = 23  //col:280
    PsAttributeSafeOpenPromptOriginClaim  = 24  //col:281
    PsAttributeBnoIsolation  = 25  //col:282
    PsAttributeDesktopAppPolicy  = 26  //col:283
    PsAttributeChpe  = 27  //col:284
    PsAttributeMitigationAuditOptions  = 28  //col:285
    PsAttributeMachineType  = 29  //col:286
    PsAttributeComponentFilter = 30  //col:287
    PsAttributeEnableOptionalXStateFeatures  = 31  //col:288
    PsAttributeMax = 32  //col:289
)


const(
    PsNeverDuplicate = 1  //col:293
    PsRequestDuplicate  = 2  //col:294
    PsAlwaysDuplicate  = 3  //col:295
    PsMaxStdHandleStates = 4  //col:296
)


const(
    PS_MITIGATION_OPTION_NX = 1  //col:300
    PS_MITIGATION_OPTION_SEHOP = 2  //col:301
    PS_MITIGATION_OPTION_FORCE_RELOCATE_IMAGES = 3  //col:302
    PS_MITIGATION_OPTION_HEAP_TERMINATE = 4  //col:303
    PS_MITIGATION_OPTION_BOTTOM_UP_ASLR = 5  //col:304
    PS_MITIGATION_OPTION_HIGH_ENTROPY_ASLR = 6  //col:305
    PS_MITIGATION_OPTION_STRICT_HANDLE_CHECKS = 7  //col:306
    PS_MITIGATION_OPTION_WIN32K_SYSTEM_CALL_DISABLE = 8  //col:307
    PS_MITIGATION_OPTION_EXTENSION_POINT_DISABLE = 9  //col:308
    PS_MITIGATION_OPTION_PROHIBIT_DYNAMIC_CODE = 10  //col:309
    PS_MITIGATION_OPTION_CONTROL_FLOW_GUARD = 11  //col:310
    PS_MITIGATION_OPTION_BLOCK_NON_MICROSOFT_BINARIES = 12  //col:311
    PS_MITIGATION_OPTION_FONT_DISABLE = 13  //col:312
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_REMOTE = 14  //col:313
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_LOW_LABEL = 15  //col:314
    PS_MITIGATION_OPTION_IMAGE_LOAD_PREFER_SYSTEM32 = 16  //col:315
    PS_MITIGATION_OPTION_RETURN_FLOW_GUARD = 17  //col:316
    PS_MITIGATION_OPTION_LOADER_INTEGRITY_CONTINUITY = 18  //col:317
    PS_MITIGATION_OPTION_STRICT_CONTROL_FLOW_GUARD = 19  //col:318
    PS_MITIGATION_OPTION_RESTRICT_SET_THREAD_CONTEXT = 20  //col:319
    PS_MITIGATION_OPTION_ROP_STACKPIVOT  = 21  //col:320
    PS_MITIGATION_OPTION_ROP_CALLER_CHECK = 22  //col:321
    PS_MITIGATION_OPTION_ROP_SIMEXEC = 23  //col:322
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER = 24  //col:323
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER_PLUS = 25  //col:324
    PS_MITIGATION_OPTION_RESTRICT_CHILD_PROCESS_CREATION = 26  //col:325
    PS_MITIGATION_OPTION_IMPORT_ADDRESS_FILTER = 27  //col:326
    PS_MITIGATION_OPTION_MODULE_TAMPERING_PROTECTION = 28  //col:327
    PS_MITIGATION_OPTION_RESTRICT_INDIRECT_BRANCH_PREDICTION = 29  //col:328
    PS_MITIGATION_OPTION_SPECULATIVE_STORE_BYPASS_DISABLE  = 30  //col:329
    PS_MITIGATION_OPTION_ALLOW_DOWNGRADE_DYNAMIC_CODE_POLICY = 31  //col:330
    PS_MITIGATION_OPTION_CET_USER_SHADOW_STACKS = 32  //col:331
    PS_MITIGATION_OPTION_USER_CET_SET_CONTEXT_IP_VALIDATION  = 33  //col:332
    PS_MITIGATION_OPTION_BLOCK_NON_CET_BINARIES = 34  //col:333
    PS_MITIGATION_OPTION_CET_DYNAMIC_APIS_OUT_OF_PROC_ONLY = 35  //col:334
    PS_MITIGATION_OPTION_REDIRECTION_TRUST  = 36  //col:335
)


const(
    PsCreateInitialState = 1  //col:339
    PsCreateFailOnFileOpen = 2  //col:340
    PsCreateFailOnSectionCreate = 3  //col:341
    PsCreateFailExeFormat = 4  //col:342
    PsCreateFailMachineMismatch = 5  //col:343
    PsCreateFailExeName  = 6  //col:344
    PsCreateSuccess = 7  //col:345
    PsCreateMaximumStates = 8  //col:346
)


const(
    MemoryReserveUserApc = 1  //col:350
    MemoryReserveIoCompletion = 2  //col:351
    MemoryReserveTypeMax = 3  //col:352
)



type  _PEB_LDR_DATA struct{
Length uint32 //col:14
Initialized bool //col:15
SsHandle uintptr //col:16
InLoadOrderModuleList *list.List //col:17
InMemoryOrderModuleList *list.List //col:18
InInitializationOrderModuleList *list.List //col:19
EntryInProgress uintptr //col:20
ShutdownInProgress bool //col:21
ShutdownThreadId uintptr //col:22
}


type  _INITIAL_TEB struct{
Struct struct //col:21
OldStackBase uintptr //col:23
OldStackLimit uintptr //col:24
}


type  _WOW64_PROCESS struct{
Wow64 uintptr //col:29
}


type  _PAGE_PRIORITY_INFORMATION struct{
PagePriority uint32 //col:33
}


type  _PROCESS_BASIC_INFORMATION struct{
ExitStatus NTSTATUS //col:42
PebBaseAddress PPEB //col:43
AffinityMask KAFFINITY //col:44
BasePriority KPRIORITY //col:45
UniqueProcessId uintptr //col:46
InheritedFromUniqueProcessId uintptr //col:47
}


type  _PROCESS_EXTENDED_BASIC_INFORMATION struct{
Size int64 //col:62
BasicInfo PROCESS_BASIC_INFORMATION //col:63
Union union //col:64
Flags uint32 //col:66
Struct struct //col:67
IsProtectedProcess uint32 //col:69
IsWow64Process uint32 //col:70
IsProcessDeleting uint32 //col:71
IsCrossSessionCreate uint32 //col:72
IsFrozen uint32 //col:73
IsBackground uint32 //col:74
IsStronglyNamed uint32 //col:75
IsSecureProcess uint32 //col:76
IsSubsystemProcess uint32 //col:77
SpareBits uint32 //col:78
}


type  _VM_COUNTERS struct{
PeakVirtualSize int64 //col:78
VirtualSize int64 //col:79
PageFaultCount uint32 //col:80
PeakWorkingSetSize int64 //col:81
WorkingSetSize int64 //col:82
QuotaPeakPagedPoolUsage int64 //col:83
QuotaPagedPoolUsage int64 //col:84
QuotaPeakNonPagedPoolUsage int64 //col:85
QuotaNonPagedPoolUsage int64 //col:86
PagefileUsage int64 //col:87
PeakPagefileUsage int64 //col:88
}


type  _VM_COUNTERS_EX struct{
PeakVirtualSize int64 //col:93
VirtualSize int64 //col:94
PageFaultCount uint32 //col:95
PeakWorkingSetSize int64 //col:96
WorkingSetSize int64 //col:97
QuotaPeakPagedPoolUsage int64 //col:98
QuotaPagedPoolUsage int64 //col:99
QuotaPeakNonPagedPoolUsage int64 //col:100
QuotaNonPagedPoolUsage int64 //col:101
PagefileUsage int64 //col:102
PeakPagefileUsage int64 //col:103
PrivateUsage int64 //col:104
}


type  _VM_COUNTERS_EX2 struct{
CountersEx VM_COUNTERS_EX //col:99
PrivateWorkingSetSize int64 //col:100
SharedCommitUsage int64 //col:101
}


type  _KERNEL_USER_TIMES struct{
CreateTime LARGE_INTEGER //col:106
ExitTime LARGE_INTEGER //col:107
KernelTime LARGE_INTEGER //col:108
UserTime LARGE_INTEGER //col:109
}


type  _POOLED_USAGE_AND_LIMITS struct{
PeakPagedPoolUsage int64 //col:118
PagedPoolUsage int64 //col:119
PagedPoolLimit int64 //col:120
PeakNonPagedPoolUsage int64 //col:121
NonPagedPoolUsage int64 //col:122
NonPagedPoolLimit int64 //col:123
PeakPagefileUsage int64 //col:124
PagefileUsage int64 //col:125
PagefileLimit int64 //col:126
}


type  _PROCESS_EXCEPTION_PORT  struct{
HANDLE _In_ //col:123
ULONG _Inout_ //col:124
}


type  _PROCESS_ACCESS_TOKEN struct{
Token uintptr //col:128
Thread uintptr //col:129
}


type  _PROCESS_LDT_INFORMATION struct{
Start uint32 //col:134
Length uint32 //col:135
LdtEntries[1] LDT_ENTRY //col:136
}


type  _PROCESS_LDT_SIZE struct{
Length uint32 //col:138
}


type  _PROCESS_WS_WATCH_INFORMATION struct{
FaultingPc uintptr //col:143
FaultingVa uintptr //col:144
}


type  _PROCESS_WS_WATCH_INFORMATION_EX struct{
BasicInfo PROCESS_WS_WATCH_INFORMATION //col:149
FaultingThreadId ULONG_PTR //col:150
Flags ULONG_PTR //col:151
}


type  _PROCESS_PRIORITY_CLASS struct{
Foreground bool //col:154
PriorityClass uint8 //col:155
}


type  _PROCESS_FOREGROUND_BACKGROUND struct{
Foreground bool //col:158
}


type  _PROCESS_DEVICEMAP_INFORMATION struct{
Union union //col:166
Struct struct //col:168
DirectoryHandle uintptr //col:170
}


type  _PROCESS_DEVICEMAP_INFORMATION_EX struct{
Union union //col:181
Struct struct //col:183
DirectoryHandle uintptr //col:185
}


type  _PROCESS_SESSION_INFORMATION struct{
SessionId uint32 //col:193
}


type  _PROCESS_HANDLE_TRACING_ENABLE struct{
Flags uint32 //col:197
}


type  _PROCESS_HANDLE_TRACING_ENABLE_EX struct{
Flags uint32 //col:202
TotalSlots uint32 //col:203
}


type  _PROCESS_HANDLE_TRACING_ENTRY struct{
Handle uintptr //col:209
ClientId CLIENT_ID //col:210
Type uint32 //col:211
Stacks[PROCESS_HANDLE_TRACING_MAX_STACKS] uintptr //col:212
}


type  _PROCESS_HANDLE_TRACING_QUERY struct{
Handle uintptr //col:215
TotalTraces uint32 //col:216
HandleTrace[1] PROCESS_HANDLE_TRACING_ENTRY //col:217
}


type  _THREAD_TLS_INFORMATION struct{
Flags uint32 //col:222
NewTlsData uintptr //col:223
OldTlsData uintptr //col:224
ThreadId uintptr //col:225
}


type  _PROCESS_TLS_INFORMATION struct{
Flags uint32 //col:231
OperationType uint32 //col:232
ThreadDataCount uint32 //col:233
TlsIndex uint32 //col:234
PreviousCount uint32 //col:235
ThreadData[1] THREAD_TLS_INFORMATION //col:236
}


type  _PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION struct{
Version uint32 //col:237
Reserved uint32 //col:238
Callback uintptr //col:239
}


type  _PROCESS_STACK_ALLOCATION_INFORMATION struct{
ReserveSize int64 //col:243
ZeroBits int64 //col:244
StackBase uintptr //col:245
}


type  _PROCESS_STACK_ALLOCATION_INFORMATION_EX struct{
PreferredNode uint32 //col:251
Reserved0 uint32 //col:252
Reserved1 uint32 //col:253
Reserved2 uint32 //col:254
AllocInfo PROCESS_STACK_ALLOCATION_INFORMATION //col:255
}


type  _PROCESS_HANDLE_INFORMATION struct{
HandleCount uint32 //col:256
HandleCountHighWatermark uint32 //col:257
}


type  _PROCESS_CYCLE_TIME_INFORMATION struct{
AccumulatedCycles ULONGLONG //col:261
CurrentCycleCount ULONGLONG //col:262
}


type  _PROCESS_WINDOW_INFORMATION struct{
WindowFlags uint32 //col:267
WindowTitleLength uint16 //col:268
WindowTitle[1] WCHAR //col:269
}


type  _PROCESS_HANDLE_TABLE_ENTRY_INFO struct{
HandleValue uintptr //col:277
HandleCount ULONG_PTR //col:278
PointerCount ULONG_PTR //col:279
GrantedAccess uint32 //col:280
ObjectTypeIndex uint32 //col:281
HandleAttributes uint32 //col:282
Reserved uint32 //col:283
}


type  _PROCESS_HANDLE_SNAPSHOT_INFORMATION struct{
NumberOfHandles ULONG_PTR //col:283
Reserved ULONG_PTR //col:284
Handles[1] PROCESS_HANDLE_TABLE_ENTRY_INFO //col:285
}


type  _PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY struct{
Flags uint32 //col:293
EnforceRedirectionTrust uint32 //col:295
AuditRedirectionTrust uint32 //col:296
ReservedFlags uint32 //col:297
}


type  _PROCESS_MITIGATION_POLICY_INFORMATION struct{
Policy PROCESS_MITIGATION_POLICY //col:315
Union union //col:316
ASLRPolicy PROCESS_MITIGATION_ASLR_POLICY //col:318
StrictHandleCheckPolicy PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY //col:319
SystemCallDisablePolicy PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY //col:320
ExtensionPointDisablePolicy PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY //col:321
DynamicCodePolicy PROCESS_MITIGATION_DYNAMIC_CODE_POLICY //col:322
ControlFlowGuardPolicy PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY //col:323
SignaturePolicy PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY //col:324
FontDisablePolicy PROCESS_MITIGATION_FONT_DISABLE_POLICY //col:325
ImageLoadPolicy PROCESS_MITIGATION_IMAGE_LOAD_POLICY //col:326
SystemCallFilterPolicy PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY //col:327
PayloadRestrictionPolicy PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY //col:328
ChildProcessPolicy PROCESS_MITIGATION_CHILD_PROCESS_POLICY //col:329
SideChannelIsolationPolicy PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY //col:330
UserShadowStackPolicy PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY //col:331
RedirectionTrustPolicy PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY //col:332
}


type  _PROCESS_KEEPALIVE_COUNT_INFORMATION struct{
WakeCount uint32 //col:321
NoWakeCount uint32 //col:322
}


type  _PROCESS_REVOKE_FILE_HANDLES_INFORMATION struct{
TargetDevicePath *int16 //col:325
}


type  _PROCESS_WORKING_SET_CONTROL struct{
Version uint32 //col:331
Operation PROCESS_WORKING_SET_OPERATION //col:332
Flags uint32 //col:333
}


type  _PS_PROTECTION struct{
Union union //col:342
Level uint8 //col:344
Struct struct //col:345
Type uint8 //col:347
Audit uint8 //col:348
Signer uint8 //col:349
}


type  _PROCESS_FAULT_INFORMATION struct{
FaultFlags uint32 //col:349
AdditionalInfo uint32 //col:350
}


type  _PROCESS_TELEMETRY_ID_INFORMATION struct{
HeaderSize uint32 //col:369
ProcessId uint32 //col:370
ProcessStartKey ULONGLONG //col:371
CreateTime ULONGLONG //col:372
CreateInterruptTime ULONGLONG //col:373
CreateUnbiasedInterruptTime ULONGLONG //col:374
ProcessSequenceNumber ULONGLONG //col:375
SessionCreateTime ULONGLONG //col:376
SessionId uint32 //col:377
BootId uint32 //col:378
ImageChecksum uint32 //col:379
ImageTimeDateStamp uint32 //col:380
UserSidOffset uint32 //col:381
ImagePathOffset uint32 //col:382
PackageNameOffset uint32 //col:383
RelativeAppNameOffset uint32 //col:384
CommandLineOffset uint32 //col:385
}


type  _PROCESS_COMMIT_RELEASE_INFORMATION struct{
Version uint32 //col:379
Struct struct //col:380
Eligible uint32 //col:382
ReleaseRepurposedMemResetCommit uint32 //col:383
ForceReleaseMemResetCommit uint32 //col:384
Spare uint32 //col:385
}


type  _PROCESS_JOB_MEMORY_INFO struct{
SharedCommitUsage ULONGLONG //col:391
PrivateCommitUsage ULONGLONG //col:392
PeakPrivateCommitUsage ULONGLONG //col:393
PrivateCommitLimit ULONGLONG //col:394
TotalCommitLimit ULONGLONG //col:395
}


type  _PROCESS_CHILD_PROCESS_INFORMATION struct{
ProhibitChildProcesses bool //col:397
AlwaysAllowSecureChildProcess bool //col:398
AuditProhibitChildProcesses bool //col:399
}


type  _POWER_THROTTLING_PROCESS_STATE struct{
Version uint32 //col:403
ControlMask uint32 //col:404
StateMask uint32 //col:405
}


type  _WIN32K_SYSCALL_FILTER struct{
FilterState uint32 //col:408
FilterSet uint32 //col:409
}


type  _PROCESS_WAKE_INFORMATION struct{
NotificationChannel ULONGLONG //col:414
WakeCounters[7] uint32 //col:415
_JOBOBJECT_WAKE_FILTER* struct //col:416
}


type  _PROCESS_ENERGY_TRACKING_STATE struct{
StateUpdateMask uint32 //col:422
StateDesiredValue uint32 //col:423
StateSequence uint32 //col:424
UpdateTag uint32 //col:425
Tag[64] WCHAR //col:426
}


type  _MANAGE_WRITES_TO_EXECUTABLE_MEMORY struct{
Version uint32 //col:430
ProcessEnableWriteExceptions uint32 //col:431
ThreadAllowWrites uint32 //col:432
Spare uint32 //col:433
KernelWriteToExecutableSignal uintptr //col:434
}


type  _POWER_THROTTLING_THREAD_STATE struct{
Version uint32 //col:436
ControlMask uint32 //col:437
StateMask uint32 //col:438
}


type  _PROCESS_UPTIME_INFORMATION struct{
QueryInterruptTime ULONGLONG //col:451
QueryUnbiasedTime ULONGLONG //col:452
EndInterruptTime ULONGLONG //col:453
TimeSinceCreation ULONGLONG //col:454
Uptime ULONGLONG //col:455
SuspendedTime ULONGLONG //col:456
Union union //col:457
HangCount uint32 //col:459
GhostCount uint32 //col:460
Crashed uint32 //col:461
Terminated uint32 //col:462
}


type  _PROCESS_SECURITY_DOMAIN_INFORMATION struct{
SecurityDomain ULONGLONG //col:456
}


type  _PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION struct{
ProcessHandle uintptr //col:460
}


type  _PROCESS_LEAP_SECOND_INFORMATION struct{
Flags uint32 //col:465
Reserved uint32 //col:466
}


type  _PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION struct{
ReserveSize ULONGLONG //col:473
CommitSize ULONGLONG //col:474
PreferredNode uint32 //col:475
Reserved uint32 //col:476
Ssp uintptr //col:477
}


type  _PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION struct{
Ssp uintptr //col:477
}


type  _THREAD_BASIC_INFORMATION struct{
ExitStatus NTSTATUS //col:486
TebBaseAddress PTEB //col:487
ClientId CLIENT_ID //col:488
AffinityMask KAFFINITY //col:489
Priority KPRIORITY //col:490
BasePriority KPRIORITY //col:491
}


type  _THREAD_LAST_SYSCALL_INFORMATION struct{
FirstArgument uintptr //col:497
SystemCallNumber uint16 //col:498
#ifdefWin64 #ifdef WIN64 //col:499
Pad[0x3] uint16 //col:500
#else #else //col:501
Pad[0x1] uint16 //col:502
#endif #endif //col:503
WaitTime ULONG64 //col:504
}


type  _THREAD_CYCLE_TIME_INFORMATION struct{
AccumulatedCycles ULONGLONG //col:502
CurrentCycleCount ULONGLONG //col:503
}


type  _THREAD_TEB_INFORMATION struct{
TebInformation uintptr //col:508
TebOffset uint32 //col:509
BytesToRead uint32 //col:510
}


type  _COUNTER_READING struct{
Type HARDWARE_COUNTER_TYPE //col:515
Index uint32 //col:516
Start ULONG64 //col:517
Total ULONG64 //col:518
}


type  _THREAD_PERFORMANCE_DATA struct{
Size uint16 //col:528
Version uint16 //col:529
ProcessorNumber PROCESSOR_NUMBER //col:530
ContextSwitches uint32 //col:531
HwCountersCount uint32 //col:532
UpdateCount ULONG64 //col:533
WaitReasonBitMap ULONG64 //col:534
HardwareCounters ULONG64 //col:535
CycleTime COUNTER_READING //col:536
HwCounters[MAX_HW_COUNTERS] COUNTER_READING //col:537
}


type  _THREAD_PROFILING_INFORMATION struct{
HardwareCounters ULONG64 //col:535
Flags uint32 //col:536
Enable uint32 //col:537
PerformanceData PTHREAD_PERFORMANCE_DATA //col:538
}


type  _RTL_UMS_CONTEXT struct{
Link SINGLE_LIST_ENTRY //col:558
Context CONTEXT //col:559
Teb uintptr //col:560
UserContext uintptr //col:561
ScheduledThread uint32 //col:562
Suspended uint32 //col:563
VolatileContext uint32 //col:564
Terminated uint32 //col:565
DebugActive uint32 //col:566
RunningOnSelfThread uint32 //col:567
DenyRunningOnSelfThread uint32 //col:568
Flags int32 //col:569
KernelUpdateLock ULONG64 //col:570
PrimaryClientID ULONG64 //col:571
ContextLock ULONG64 //col:572
_RTL_UMS_CONTEXT* struct //col:573
SwitchCount uint32 //col:574
KernelYieldCount uint32 //col:575
MixedYieldCount uint32 //col:576
YieldCount uint32 //col:577
}


type  _RTL_UMS_COMPLETION_LIST struct{
ThreadListHead PSINGLE_LIST_ENTRY //col:565
CompletionEvent uintptr //col:566
CompletionFlags uint32 //col:567
InternalListHead SINGLE_LIST_ENTRY //col:568
}


type  _THREAD_UMS_INFORMATION struct{
Command THREAD_UMS_INFORMATION_COMMAND //col:579
CompletionList PRTL_UMS_COMPLETION_LIST //col:580
UmsContext PRTL_UMS_CONTEXT //col:581
Union union //col:582
Flags uint32 //col:584
Struct struct //col:585
IsUmsSchedulerThread uint32 //col:587
IsUmsWorkerThread uint32 //col:588
SpareBits uint32 //col:589
}


type  _THREAD_NAME_INFORMATION struct{
ThreadName *int16 //col:585
}


type  _ALPC_WORK_ON_BEHALF_TICKET struct{
ThreadId uint32 //col:590
ThreadCreationTimeLow uint32 //col:591
}


type  _RTL_WORK_ON_BEHALF_TICKET_EX struct{
Ticket ALPC_WORK_ON_BEHALF_TICKET //col:601
Union union //col:602
Flags uint32 //col:604
Struct struct //col:605
CurrentThread uint32 //col:607
Reserved1 uint32 //col:608
}


type  _PROC_THREAD_ATTRIBUTE  struct{
Attribute ULONG_PTR //col:608
Size int64 //col:609
Value ULONG_PTR //col:610
}


type  _PROC_THREAD_ATTRIBUTE_LIST  struct{
PresentFlags uint32 //col:616
AttributeCount uint32 //col:617
LastAttribute uint32 //col:618
SpareUlong0 uint32 //col:619
ExtendedFlagsAttribute PPROC_THREAD_ATTRIBUTE //col:620
Attributes[1] PROC_THREAD_ATTRIBUTE //col:621
}


type  _SE_SAFE_OPEN_PROMPT_RESULTS  struct{
Results SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS //col:620
Path[MAX_PATH] WCHAR //col:621
}


type  _PROC_THREAD_BNOISOLATION_ATTRIBUTE struct{
IsolationEnabled BOOL //col:626
IsolationPrefix[0x88] WCHAR //col:627
}


type  _ISOLATION_MANIFEST_PROPERTIES  struct{
InstancePath *int16 //col:631
FriendlyName *int16 //col:632
Description *int16 //col:633
Level ULONG_PTR //col:634
}


type  _PS_ATTRIBUTE struct{
Attribute ULONG_PTR //col:641
Size int64 //col:642
Union union //col:643
Value ULONG_PTR //col:645
ValuePtr uintptr //col:646
}


type  _PS_ATTRIBUTE_LIST struct{
TotalLength int64 //col:648
Attributes[1] PS_ATTRIBUTE //col:649
}


type  _PS_MEMORY_RESERVE struct{
ReserveAddress uintptr //col:653
ReserveSize int64 //col:654
}


type  _PS_STD_HANDLE_INFO struct{
Union union //col:663
Flags uint32 //col:665
Struct struct //col:666
StdHandleState uint32 //col:668
PseudoHandleMask uint32 //col:669
}


type  _PS_TRUSTLET_ATTRIBUTE_TYPE struct{
Union union //col:677
Struct struct //col:679
Version uint8 //col:681
DataCount uint8 //col:682
SemanticType uint8 //col:683
AccessRights PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS //col:684
}


type  _PS_TRUSTLET_ATTRIBUTE_HEADER struct{
AttributeType PS_TRUSTLET_ATTRIBUTE_TYPE //col:686
InstanceNumber uint32 //col:687
Reserved uint32 //col:688
}


type  _PS_TRUSTLET_ATTRIBUTE_DATA struct{
Header PS_TRUSTLET_ATTRIBUTE_HEADER //col:691
Data[1] ULONGLONG //col:692
}


type  _PS_TRUSTLET_CREATE_ATTRIBUTES struct{
TrustletIdentity ULONGLONG //col:696
Attributes[1] PS_TRUSTLET_ATTRIBUTE_DATA //col:697
}


type  _PS_BNO_ISOLATION_PARAMETERS struct{
IsolationPrefix *int16 //col:703
HandleCount uint32 //col:704
*Handles uintptr //col:705
IsolationEnabled bool //col:706
}


type  _PS_CREATE_INFO struct{
Size int64 //col:724
State PS_CREATE_STATE //col:725
Union union //col:726
Struct struct //col:728
Union union //col:730
InitFlags uint32 //col:732
Struct struct //col:733
WriteOutputOnExit uint8 //col:735
DetectManifest uint8 //col:736
IFEOSkipDebugger uint8 //col:737
IFEODoNotPropagateKeyState uint8 //col:738
SpareBits1 uint8 //col:739
SpareBits2 uint8 //col:740
ProhibitedImageCharacteristics uint16 //col:741
}


type  _JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION struct{
BasicInfo JOBOBJECT_BASIC_ACCOUNTING_INFORMATION //col:778
IoInfo IO_COUNTERS //col:779
DiskIoInfo PROCESS_DISK_COUNTERS //col:780
ContextSwitches ULONG64 //col:781
TotalCycleTime LARGE_INTEGER //col:782
ReadyTime ULONG64 //col:783
EnergyValues PROCESS_ENERGY_VALUES //col:784
}


type  _JOBOBJECT_WAKE_INFORMATION struct{
NotificationChannel uintptr //col:783
WakeCounters[7] ULONG64 //col:784
}


type  _JOBOBJECT_WAKE_INFORMATION_V1 struct{
NotificationChannel uintptr //col:788
WakeCounters[4] ULONG64 //col:789
}


type  _JOBOBJECT_INTERFERENCE_INFORMATION struct{
Count ULONG64 //col:792
}


type  _JOBOBJECT_WAKE_FILTER struct{
HighEdgeFilter uint32 //col:797
LowEdgeFilter uint32 //col:798
}


type  _JOBOBJECT_FREEZE_INFORMATION struct{
Union union //col:809
Flags uint32 //col:811
Struct struct //col:812
FreezeOperation uint32 //col:814
FilterOperation uint32 //col:815
SwapOperation uint32 //col:816
Reserved uint32 //col:817
}


type  _JOBOBJECT_MEMORY_USAGE_INFORMATION struct{
JobMemory ULONG64 //col:820
PeakJobMemoryUsed ULONG64 //col:821
}


type  _JOBOBJECT_MEMORY_USAGE_INFORMATION_V2 struct{
BasicInfo JOBOBJECT_MEMORY_USAGE_INFORMATION //col:826
JobSharedMemory ULONG64 //col:827
Reserved[2] ULONG64 //col:828
}


type  _SILO_USER_SHARED_DATA struct{
ServiceSessionId ULONG64 //col:838
ActiveConsoleId uint32 //col:839
ConsoleSessionForegroundProcessId LONGLONG //col:840
NtProductType NT_PRODUCT_TYPE //col:841
SuiteMask uint32 //col:842
SharedUserSessionId uint32 //col:843
IsMultiSessionSku bool //col:844
NtSystemRoot[260] WCHAR //col:845
UserModeGlobalLogger[16] uint16 //col:846
}


type  _SILOOBJECT_ROOT_DIRECTORY struct{
ControlFlags uint32 //col:843
Path *int16 //col:844
}


type  _JOBOBJECT_ENERGY_TRACKING_STATE struct{
Value ULONG64 //col:849
UpdateMask uint32 //col:850
DesiredState uint32 //col:851
}




