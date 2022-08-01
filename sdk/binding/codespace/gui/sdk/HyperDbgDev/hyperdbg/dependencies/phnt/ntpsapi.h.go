package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpsapi.h.back

const(
_NTPSAPI_H =  //col:1
PROCESS_TERMINATE = 0x0001 //col:2
PROCESS_CREATE_THREAD = 0x0002 //col:3
PROCESS_SET_SESSIONID = 0x0004 //col:4
PROCESS_VM_OPERATION = 0x0008 //col:5
PROCESS_VM_READ = 0x0010 //col:6
PROCESS_VM_WRITE = 0x0020 //col:7
PROCESS_CREATE_PROCESS = 0x0080 //col:8
PROCESS_SET_QUOTA = 0x0100 //col:9
PROCESS_SET_INFORMATION = 0x0200 //col:10
PROCESS_QUERY_INFORMATION = 0x0400 //col:11
PROCESS_SET_PORT = 0x0800 //col:12
PROCESS_SUSPEND_RESUME = 0x0800 //col:13
PROCESS_QUERY_LIMITED_INFORMATION = 0x1000 //col:14
PROCESS_SET_PORT = 0x0800 //col:15
THREAD_QUERY_INFORMATION = 0x0040 //col:16
THREAD_SET_THREAD_TOKEN = 0x0080 //col:17
THREAD_IMPERSONATE = 0x0100 //col:18
THREAD_DIRECT_IMPERSONATION = 0x0200 //col:19
THREAD_ALERT = 0x0004 //col:20
JOB_OBJECT_ASSIGN_PROCESS = 0x0001 //col:21
JOB_OBJECT_SET_ATTRIBUTES = 0x0002 //col:22
JOB_OBJECT_QUERY = 0x0004 //col:23
JOB_OBJECT_TERMINATE = 0x0008 //col:24
JOB_OBJECT_SET_SECURITY_ATTRIBUTES = 0x0010 //col:25
JOB_OBJECT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3F) //col:26
GDI_HANDLE_BUFFER_SIZE32 = 34 //col:27
GDI_HANDLE_BUFFER_SIZE64 = 60 //col:28
GDI_HANDLE_BUFFER_SIZE = GDI_HANDLE_BUFFER_SIZE32 //col:29
GDI_HANDLE_BUFFER_SIZE = GDI_HANDLE_BUFFER_SIZE64 //col:30
FLS_MAXIMUM_AVAILABLE = 128 //col:31
TLS_MINIMUM_AVAILABLE = 64 //col:32
TLS_EXPANSION_SLOTS = 1024 //col:33
PROCESS_EXCEPTION_PORT_ALL_STATE_BITS = 0x00000003 //col:34
PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS = ((ULONG_PTR)((1UL << PROCESS_EXCEPTION_PORT_ALL_STATE_BITS) - 1)) //col:35
PROCESS_PRIORITY_CLASS_UNKNOWN = 0 //col:36
PROCESS_PRIORITY_CLASS_IDLE = 1 //col:37
PROCESS_PRIORITY_CLASS_NORMAL = 2 //col:38
PROCESS_PRIORITY_CLASS_HIGH = 3 //col:39
PROCESS_PRIORITY_CLASS_REALTIME = 4 //col:40
PROCESS_PRIORITY_CLASS_BELOW_NORMAL = 5 //col:41
PROCESS_PRIORITY_CLASS_ABOVE_NORMAL = 6 //col:42
PROCESS_LUID_DOSDEVICES_ONLY = 0x00000001 //col:43
PROCESS_HANDLE_EXCEPTIONS_ENABLED = 0x00000001 //col:44
PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_DISABLED = 0x00000000 //col:45
PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_ENABLED = 0x00000001 //col:46
PROCESS_HANDLE_TRACING_MAX_SLOTS = 0x20000 //col:47
PROCESS_HANDLE_TRACING_MAX_STACKS = 16 //col:48
PROCESS_HANDLE_TRACE_TYPE_OPEN = 1 //col:49
PROCESS_HANDLE_TRACE_TYPE_CLOSE = 2 //col:50
PROCESS_HANDLE_TRACE_TYPE_BADREF = 3 //col:51
PS_PROTECTED_SIGNER_MASK = 0xFF //col:52
PS_PROTECTED_AUDIT_MASK = 0x08 //col:53
PS_PROTECTED_TYPE_MASK = 0x07 //col:54
PsProtectedValue(aSigner, aAudit, aType) ( = ((aSigner & PS_PROTECTED_SIGNER_MASK) << 4) | ((aAudit & PS_PROTECTED_AUDIT_MASK) << 3) | (aType & PS_PROTECTED_TYPE_MASK) ) //col:55
InitializePsProtection(aProtectionLevelPtr, aSigner, aAudit, aType) { = (aProtectionLevelPtr)->Signer = aSigner; (aProtectionLevelPtr)->Audit = aAudit; (aProtectionLevelPtr)->Type = aType; } //col:60
POWER_THROTTLING_PROCESS_CURRENT_VERSION = 1 //col:65
POWER_THROTTLING_PROCESS_EXECUTION_SPEED = 0x1 //col:66
POWER_THROTTLING_PROCESS_DELAYTIMERS = 0x2 //col:67
POWER_THROTTLING_PROCESS_VALID_FLAGS = ((POWER_THROTTLING_PROCESS_EXECUTION_SPEED | POWER_THROTTLING_PROCESS_DELAYTIMERS)) //col:68
WIN32K_SYSCALL_FILTER_STATE_ENABLE = 0x1 //col:69
WIN32K_SYSCALL_FILTER_STATE_AUDIT = 0x2 //col:70
POWER_THROTTLING_THREAD_CURRENT_VERSION = 1 //col:71
POWER_THROTTLING_THREAD_EXECUTION_SPEED = 0x1 //col:72
POWER_THROTTLING_THREAD_VALID_FLAGS = (POWER_THROTTLING_THREAD_EXECUTION_SPEED) //col:73
PROCESS_READWRITEVM_LOGGING_ENABLE_READVM = 1 //col:74
PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM = 2 //col:75
PROCESS_READWRITEVM_LOGGING_ENABLE_READVM_V = 1UL //col:76
PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM_V = 2UL //col:77
PROCESS_CREATE_FLAGS_BREAKAWAY = 0x00000001 //col:78
PROCESS_CREATE_FLAGS_NO_DEBUG_INHERIT = 0x00000002 //col:79
PROCESS_CREATE_FLAGS_INHERIT_HANDLES = 0x00000004 //col:80
PROCESS_CREATE_FLAGS_OVERRIDE_ADDRESS_SPACE = 0x00000008 //col:81
PROCESS_CREATE_FLAGS_LARGE_PAGES = 0x00000010 //col:82
PROCESS_CREATE_FLAGS_LARGE_PAGE_SYSTEM_DLL = 0x00000020 //col:83
PROCESS_CREATE_FLAGS_PROTECTED_PROCESS = 0x00000040 //col:84
PROCESS_CREATE_FLAGS_CREATE_SESSION = 0x00000080 //col:85
PROCESS_CREATE_FLAGS_INHERIT_FROM_PARENT = 0x00000100 //col:86
PROCESS_CREATE_FLAGS_SUSPENDED = 0x00000200 //col:87
PROCESS_CREATE_FLAGS_FORCE_BREAKAWAY = 0x00000400 //col:88
PROCESS_CREATE_FLAGS_MINIMAL_PROCESS = 0x00000800 //col:89
PROCESS_CREATE_FLAGS_RELEASE_SECTION = 0x00001000 //col:90
PROCESS_CREATE_FLAGS_CLONE_MINIMAL = 0x00002000 //col:91
PROCESS_CREATE_FLAGS_CLONE_MINIMAL_REDUCED_COMMIT = 0x00004000 //col:92
PROCESS_CREATE_FLAGS_AUXILIARY_PROCESS = 0x00008000 //col:93
PROCESS_CREATE_FLAGS_CREATE_STORE = 0x00020000 //col:94
PROCESS_CREATE_FLAGS_USE_PROTECTED_ENVIRONMENT = 0x00040000 //col:95
NtCurrentProcess() = ((HANDLE)(LONG_PTR)-1) //col:96
ZwCurrentProcess() = NtCurrentProcess() //col:97
NtCurrentThread() = ((HANDLE)(LONG_PTR)-2) //col:98
ZwCurrentThread() = NtCurrentThread() //col:99
NtCurrentSession() = ((HANDLE)(LONG_PTR)-3) //col:100
ZwCurrentSession() = NtCurrentSession() //col:101
NtCurrentPeb() = (NtCurrentTeb()->ProcessEnvironmentBlock) //col:102
NtCurrentProcessToken() = ((HANDLE)(LONG_PTR)-4) //col:103
NtCurrentThreadToken() = ((HANDLE)(LONG_PTR)-5) //col:104
NtCurrentThreadEffectiveToken() = ((HANDLE)(LONG_PTR)-6) //col:105
NtCurrentSilo() = ((HANDLE)(LONG_PTR)-1) //col:106
NtCurrentProcessId() = (NtCurrentTeb()->ClientId.UniqueProcess) //col:107
NtCurrentThreadId() = (NtCurrentTeb()->ClientId.UniqueThread) //col:108
PROCESS_GET_NEXT_FLAGS_PREVIOUS_PROCESS = 0x00000001 //col:109
STATECHANGE_SET_ATTRIBUTES = 0x0001 //col:110
Wow64EncodeApcRoutine(ApcRoutine) = ((PVOID)((0 - ((LONG_PTR)(ApcRoutine))) << 2)) //col:111
Wow64DecodeApcRoutine(ApcRoutine) = ((PVOID)(0 - (((LONG_PTR)(ApcRoutine)) >> 2))) //col:113
APC_FORCE_THREAD_SIGNAL = ((HANDLE)1) //col:115
ProcThreadAttributeParentProcess = 0 //col:116
ProcThreadAttributeExtendedFlags = 1 //col:117
ProcThreadAttributeHandleList = 2 //col:118
ProcThreadAttributeGroupAffinity = 3 //col:119
ProcThreadAttributePreferredNode = 4 //col:120
ProcThreadAttributeIdealProcessor = 5 //col:121
ProcThreadAttributeUmsThread = 6 //col:122
ProcThreadAttributeMitigationPolicy = 7 //col:123
ProcThreadAttributePackageFullName = 8 //col:124
ProcThreadAttributeSecurityCapabilities = 9 //col:125
ProcThreadAttributeConsoleReference = 10 //col:126
ProcThreadAttributeProtectionLevel = 11 //col:127
ProcThreadAttributeOsMaxVersionTested = 12 //col:128
ProcThreadAttributeJobList = 13 //col:129
ProcThreadAttributeChildProcessPolicy = 14 //col:130
ProcThreadAttributeAllApplicationPackagesPolicy = 15 //col:131
ProcThreadAttributeWin32kFilter = 16 //col:132
ProcThreadAttributeSafeOpenPromptOriginClaim = 17 //col:133
ProcThreadAttributeDesktopAppPolicy = 18 //col:134
ProcThreadAttributeBnoIsolation = 19 //col:135
ProcThreadAttributePseudoConsole = 22 //col:136
ProcThreadAttributeIsolationManifest = 23 //col:137
ProcThreadAttributeMitigationAuditPolicy = 24 //col:138
ProcThreadAttributeMachineType = 25 //col:139
ProcThreadAttributeComponentFilter = 26 //col:140
ProcThreadAttributeEnableOptionalXStateFeatures = 27 //col:141
ProcThreadAttributeCreateStore = 28 //col:142
PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS = ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE) //col:143
PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME = ProcThreadAttributeValue(ProcThreadAttributePackageFullName, FALSE, TRUE, FALSE) //col:145
PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE = ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE) //col:147
PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED = ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE) //col:149
PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:151
PROC_THREAD_ATTRIBUTE_BNO_ISOLATION = ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:153
PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST = ProcThreadAttributeValue(ProcThreadAttributeIsolationManifest, FALSE, TRUE, FALSE) //col:155
PROC_THREAD_ATTRIBUTE_CREATE_STORE = ProcThreadAttributeValue(ProcThreadAttributeCreateStore, FALSE, TRUE, FALSE) //col:157
EXTENDED_PROCESS_CREATION_FLAG_ELEVATION_HANDLED = 0x00000001 //col:159
EXTENDED_PROCESS_CREATION_FLAG_FORCELUA = 0x00000002 //col:160
EXTENDED_PROCESS_CREATION_FLAG_FORCE_BREAKAWAY = 0x00000004 //col:161
PROTECTION_LEVEL_WINTCB_LIGHT = 0x00000000 //col:162
PROTECTION_LEVEL_WINDOWS = 0x00000001 //col:163
PROTECTION_LEVEL_WINDOWS_LIGHT = 0x00000002 //col:164
PROTECTION_LEVEL_ANTIMALWARE_LIGHT = 0x00000003 //col:165
PROTECTION_LEVEL_LSA_LIGHT = 0x00000004 //col:166
PROTECTION_LEVEL_WINTCB = 0x00000005 //col:167
PROTECTION_LEVEL_CODEGEN_LIGHT = 0x00000006 //col:168
PROTECTION_LEVEL_AUTHENTICODE = 0x00000007 //col:169
PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS = ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE) //col:170
PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE = ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE) //col:172
PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED = ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE) //col:174
PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:176
PROC_THREAD_ATTRIBUTE_BNO_ISOLATION = ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:178
PS_ATTRIBUTE_NUMBER_MASK = 0x0000ffff //col:180
PS_ATTRIBUTE_THREAD = 0x00010000 //col:181
PS_ATTRIBUTE_INPUT = 0x00020000 //col:182
PS_ATTRIBUTE_ADDITIVE = 0x00040000 //col:183
PsAttributeValue(Number, Thread, Input, Additive) = (((Number) & PS_ATTRIBUTE_NUMBER_MASK) | ((Thread) ? PS_ATTRIBUTE_THREAD : 0) | ((Input) ? PS_ATTRIBUTE_INPUT : 0) | ((Additive) ? PS_ATTRIBUTE_ADDITIVE : 0)) //col:184
PS_ATTRIBUTE_PARENT_PROCESS = PsAttributeValue(PsAttributeParentProcess, FALSE, TRUE, TRUE) //col:189
PS_ATTRIBUTE_DEBUG_OBJECT = PsAttributeValue(PsAttributeDebugObject, FALSE, TRUE, TRUE) //col:191
PS_ATTRIBUTE_TOKEN = PsAttributeValue(PsAttributeToken, FALSE, TRUE, TRUE) //col:193
PS_ATTRIBUTE_CLIENT_ID = PsAttributeValue(PsAttributeClientId, TRUE, FALSE, FALSE) //col:195
PS_ATTRIBUTE_TEB_ADDRESS = PsAttributeValue(PsAttributeTebAddress, TRUE, FALSE, FALSE) //col:197
PS_ATTRIBUTE_IMAGE_NAME = PsAttributeValue(PsAttributeImageName, FALSE, TRUE, FALSE) //col:199
PS_ATTRIBUTE_IMAGE_INFO = PsAttributeValue(PsAttributeImageInfo, FALSE, FALSE, FALSE) //col:201
PS_ATTRIBUTE_MEMORY_RESERVE = PsAttributeValue(PsAttributeMemoryReserve, FALSE, TRUE, FALSE) //col:203
PS_ATTRIBUTE_PRIORITY_CLASS = PsAttributeValue(PsAttributePriorityClass, FALSE, TRUE, FALSE) //col:205
PS_ATTRIBUTE_ERROR_MODE = PsAttributeValue(PsAttributeErrorMode, FALSE, TRUE, FALSE) //col:207
PS_ATTRIBUTE_STD_HANDLE_INFO = PsAttributeValue(PsAttributeStdHandleInfo, FALSE, TRUE, FALSE) //col:209
PS_ATTRIBUTE_HANDLE_LIST = PsAttributeValue(PsAttributeHandleList, FALSE, TRUE, FALSE) //col:211
PS_ATTRIBUTE_GROUP_AFFINITY = PsAttributeValue(PsAttributeGroupAffinity, TRUE, TRUE, FALSE) //col:213
PS_ATTRIBUTE_PREFERRED_NODE = PsAttributeValue(PsAttributePreferredNode, FALSE, TRUE, FALSE) //col:215
PS_ATTRIBUTE_IDEAL_PROCESSOR = PsAttributeValue(PsAttributeIdealProcessor, TRUE, TRUE, FALSE) //col:217
PS_ATTRIBUTE_UMS_THREAD = PsAttributeValue(PsAttributeUmsThread, TRUE, TRUE, FALSE) //col:219
PS_ATTRIBUTE_MITIGATION_OPTIONS = PsAttributeValue(PsAttributeMitigationOptions, FALSE, TRUE, FALSE) //col:221
PS_ATTRIBUTE_PROTECTION_LEVEL = PsAttributeValue(PsAttributeProtectionLevel, FALSE, TRUE, TRUE) //col:223
PS_ATTRIBUTE_SECURE_PROCESS = PsAttributeValue(PsAttributeSecureProcess, FALSE, TRUE, FALSE) //col:225
PS_ATTRIBUTE_JOB_LIST = PsAttributeValue(PsAttributeJobList, FALSE, TRUE, FALSE) //col:227
PS_ATTRIBUTE_CHILD_PROCESS_POLICY = PsAttributeValue(PsAttributeChildProcessPolicy, FALSE, TRUE, FALSE) //col:229
PS_ATTRIBUTE_ALL_APPLICATION_PACKAGES_POLICY = PsAttributeValue(PsAttributeAllApplicationPackagesPolicy, FALSE, TRUE, FALSE) //col:231
PS_ATTRIBUTE_WIN32K_FILTER = PsAttributeValue(PsAttributeWin32kFilter, FALSE, TRUE, FALSE) //col:233
PS_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = PsAttributeValue(PsAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:235
PS_ATTRIBUTE_BNO_ISOLATION = PsAttributeValue(PsAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:237
PS_ATTRIBUTE_DESKTOP_APP_POLICY = PsAttributeValue(PsAttributeDesktopAppPolicy, FALSE, TRUE, FALSE) //col:239
PS_ATTRIBUTE_CHPE = PsAttributeValue(PsAttributeChpe, FALSE, TRUE, TRUE) //col:241
PS_ATTRIBUTE_MITIGATION_AUDIT_OPTIONS = PsAttributeValue(PsAttributeMitigationAuditOptions, FALSE, TRUE, FALSE) //col:243
PS_ATTRIBUTE_MACHINE_TYPE = PsAttributeValue(PsAttributeMachineType, FALSE, TRUE, TRUE) //col:245
PS_ATTRIBUTE_COMPONENT_FILTER = PsAttributeValue(PsAttributeComponentFilter, FALSE, TRUE, FALSE) //col:247
PS_ATTRIBUTE_ENABLE_OPTIONAL_XSTATE_FEATURES = PsAttributeValue(PsAttributeEnableOptionalXStateFeatures, TRUE, TRUE, FALSE) //col:249
PS_STD_INPUT_HANDLE = 0x1 //col:251
PS_STD_OUTPUT_HANDLE = 0x2 //col:252
PS_STD_ERROR_HANDLE = 0x4 //col:253
THREAD_CREATE_FLAGS_CREATE_SUSPENDED = 0x00000001 //col:254
THREAD_CREATE_FLAGS_SKIP_THREAD_ATTACH = 0x00000002 //col:255
THREAD_CREATE_FLAGS_HIDE_FROM_DEBUGGER = 0x00000004 //col:256
THREAD_CREATE_FLAGS_LOADER_WORKER = 0x00000010 //col:257
THREAD_CREATE_FLAGS_SKIP_LOADER_INIT = 0x00000020 //col:258
THREAD_CREATE_FLAGS_BYPASS_PROCESS_FREEZE = 0x00000040 //col:259
THREAD_CREATE_FLAGS_INITIAL_THREAD = 0x00000080 //col:260
JobObjectBasicAccountingInformation = 1 //col:261
JobObjectBasicLimitInformation = 2 //col:262
JobObjectBasicProcessIdList = 3 //col:263
JobObjectBasicUIRestrictions = 4 //col:264
JobObjectSecurityLimitInformation = 5 //col:265
JobObjectEndOfJobTimeInformation = 6 //col:266
JobObjectAssociateCompletionPortInformation = 7 //col:267
JobObjectBasicAndIoAccountingInformation = 8 //col:268
JobObjectExtendedLimitInformation = 9 //col:269
JobObjectJobSetInformation = 10 //col:270
JobObjectGroupInformation = 11 //col:271
JobObjectNotificationLimitInformation = 12 //col:272
JobObjectLimitViolationInformation = 13 //col:273
JobObjectGroupInformationEx = 14 //col:274
JobObjectCpuRateControlInformation = 15 //col:275
JobObjectCompletionFilter = 16 //col:276
JobObjectCompletionCounter = 17 //col:277
JobObjectFreezeInformation = 18 //col:278
JobObjectExtendedAccountingInformation = 19 //col:279
JobObjectWakeInformation = 20 //col:280
JobObjectBackgroundInformation = 21 //col:281
JobObjectSchedulingRankBiasInformation = 22 //col:282
JobObjectTimerVirtualizationInformation = 23 //col:283
JobObjectCycleTimeNotification = 24 //col:284
JobObjectClearEvent = 25 //col:285
JobObjectInterferenceInformation = 26 //col:286
JobObjectClearPeakJobMemoryUsed = 27 //col:287
JobObjectMemoryUsageInformation = 28 //col:288
JobObjectSharedCommit = 29 //col:289
JobObjectContainerId = 30 //col:290
JobObjectIoRateControlInformation = 31 //col:291
JobObjectNetRateControlInformation = 32 //col:292
JobObjectNotificationLimitInformation2 = 33 //col:293
JobObjectLimitViolationInformation2 = 34 //col:294
JobObjectCreateSilo = 35 //col:295
JobObjectSiloBasicInformation = 36 //col:296
JobObjectSiloRootDirectory = 37 //col:297
JobObjectServerSiloBasicInformation = 38 //col:298
JobObjectServerSiloUserSharedData = 39 //col:299
JobObjectServerSiloInitialize = 40 //col:300
JobObjectServerSiloRunningState = 41 //col:301
JobObjectIoAttribution = 42 //col:302
JobObjectMemoryPartitionInformation = 43 //col:303
JobObjectContainerTelemetryId = 44 //col:304
JobObjectSiloSystemRoot = 45 //col:305
JobObjectEnergyTrackingState = 46 //col:306
JobObjectThreadImpersonationInformation = 47 //col:307
JobObjectIoPriorityLimit = 48 //col:308
JobObjectPagePriorityLimit = 49 //col:309
MaxJobObjectInfoClass = 50 //col:310
)

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



type PEB_LDR_DATA struct{
Length uint32 //col:3
Initialized bool //col:4
SsHandle HANDLE //col:5
InLoadOrderModuleList *list.List //col:6
InMemoryOrderModuleList *list.List //col:7
InInitializationOrderModuleList *list.List //col:8
EntryInProgress PVOID //col:9
ShutdownInProgress bool //col:10
ShutdownThreadId HANDLE //col:11
}


type INITIAL_TEB struct{
Struct struct //col:15
OldStackBase PVOID //col:17
OldStackLimit PVOID //col:18
}


type WOW64_PROCESS struct{
Wow64 PVOID //col:26
}


type PAGE_PRIORITY_INFORMATION struct{
PagePriority uint32 //col:30
}


type PROCESS_BASIC_INFORMATION struct{
ExitStatus NTSTATUS //col:34
PebBaseAddress PPEB //col:35
AffinityMask KAFFINITY //col:36
BasePriority KPRIORITY //col:37
UniqueProcessId HANDLE //col:38
InheritedFromUniqueProcessId HANDLE //col:39
}


type PROCESS_EXTENDED_BASIC_INFORMATION struct{
Size SIZE_T //col:43
BasicInfo PROCESS_BASIC_INFORMATION //col:44
Union union //col:45
Flags uint32 //col:47
Struct struct //col:48
IsProtectedProcess uint32 //col:50
IsWow64Process uint32 //col:51
IsProcessDeleting uint32 //col:52
IsCrossSessionCreate uint32 //col:53
IsFrozen uint32 //col:54
IsBackground uint32 //col:55
IsStronglyNamed uint32 //col:56
IsSecureProcess uint32 //col:57
IsSubsystemProcess uint32 //col:58
SpareBits uint32 //col:59
}


type VM_COUNTERS struct{
PeakVirtualSize SIZE_T //col:65
VirtualSize SIZE_T //col:66
PageFaultCount uint32 //col:67
PeakWorkingSetSize SIZE_T //col:68
WorkingSetSize SIZE_T //col:69
QuotaPeakPagedPoolUsage SIZE_T //col:70
QuotaPagedPoolUsage SIZE_T //col:71
QuotaPeakNonPagedPoolUsage SIZE_T //col:72
QuotaNonPagedPoolUsage SIZE_T //col:73
PagefileUsage SIZE_T //col:74
PeakPagefileUsage SIZE_T //col:75
}


type VM_COUNTERS_EX struct{
PeakVirtualSize SIZE_T //col:79
VirtualSize SIZE_T //col:80
PageFaultCount uint32 //col:81
PeakWorkingSetSize SIZE_T //col:82
WorkingSetSize SIZE_T //col:83
QuotaPeakPagedPoolUsage SIZE_T //col:84
QuotaPagedPoolUsage SIZE_T //col:85
QuotaPeakNonPagedPoolUsage SIZE_T //col:86
QuotaNonPagedPoolUsage SIZE_T //col:87
PagefileUsage SIZE_T //col:88
PeakPagefileUsage SIZE_T //col:89
PrivateUsage SIZE_T //col:90
}


type VM_COUNTERS_EX2 struct{
CountersEx VM_COUNTERS_EX //col:94
PrivateWorkingSetSize SIZE_T //col:95
SharedCommitUsage SIZE_T //col:96
}


type KERNEL_USER_TIMES struct{
CreateTime LARGE_INTEGER //col:100
ExitTime LARGE_INTEGER //col:101
KernelTime LARGE_INTEGER //col:102
UserTime LARGE_INTEGER //col:103
}


type POOLED_USAGE_AND_LIMITS struct{
PeakPagedPoolUsage SIZE_T //col:107
PagedPoolUsage SIZE_T //col:108
PagedPoolLimit SIZE_T //col:109
PeakNonPagedPoolUsage SIZE_T //col:110
NonPagedPoolUsage SIZE_T //col:111
NonPagedPoolLimit SIZE_T //col:112
PeakPagefileUsage SIZE_T //col:113
PagefileUsage SIZE_T //col:114
PagefileLimit SIZE_T //col:115
}


type PROCESS_EXCEPTION_PORT  struct{
HANDLE _In_ //col:119
ULONG _Inout_ //col:120
}


type PROCESS_ACCESS_TOKEN struct{
Token HANDLE //col:124
Thread HANDLE //col:125
}


type PROCESS_LDT_INFORMATION struct{
Start uint32 //col:129
Length uint32 //col:130
LdtEntries[1] LDT_ENTRY //col:131
}


type PROCESS_LDT_SIZE struct{
Length uint32 //col:135
}


type PROCESS_WS_WATCH_INFORMATION struct{
FaultingPc PVOID //col:139
FaultingVa PVOID //col:140
}


type PROCESS_WS_WATCH_INFORMATION_EX struct{
BasicInfo PROCESS_WS_WATCH_INFORMATION //col:144
FaultingThreadId ULONG_PTR //col:145
Flags ULONG_PTR //col:146
}


type PROCESS_PRIORITY_CLASS struct{
Foreground bool //col:150
PriorityClass uint8 //col:151
}


type PROCESS_FOREGROUND_BACKGROUND struct{
Foreground bool //col:155
}


type PROCESS_DEVICEMAP_INFORMATION struct{
Union union //col:159
Struct struct //col:161
DirectoryHandle HANDLE //col:163
}


type PROCESS_DEVICEMAP_INFORMATION_EX struct{
Union union //col:174
Struct struct //col:176
DirectoryHandle HANDLE //col:178
}


type PROCESS_SESSION_INFORMATION struct{
SessionId uint32 //col:190
}


type PROCESS_HANDLE_TRACING_ENABLE struct{
Flags uint32 //col:194
}


type PROCESS_HANDLE_TRACING_ENABLE_EX struct{
Flags uint32 //col:198
TotalSlots uint32 //col:199
}


type PROCESS_HANDLE_TRACING_ENTRY struct{
Handle HANDLE //col:203
ClientId CLIENT_ID //col:204
Type uint32 //col:205
Stacks[PROCESS_HANDLE_TRACING_MAX_STACKS] PVOID //col:206
}


type PROCESS_HANDLE_TRACING_QUERY struct{
Handle HANDLE //col:210
TotalTraces uint32 //col:211
HandleTrace[1] PROCESS_HANDLE_TRACING_ENTRY //col:212
}


type THREAD_TLS_INFORMATION struct{
Flags uint32 //col:216
NewTlsData PVOID //col:217
OldTlsData PVOID //col:218
ThreadId HANDLE //col:219
}


type PROCESS_TLS_INFORMATION struct{
Flags uint32 //col:223
OperationType uint32 //col:224
ThreadDataCount uint32 //col:225
TlsIndex uint32 //col:226
PreviousCount uint32 //col:227
ThreadData[1] THREAD_TLS_INFORMATION //col:228
}


type PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION struct{
Version uint32 //col:232
Reserved uint32 //col:233
Callback PVOID //col:234
}


type PROCESS_STACK_ALLOCATION_INFORMATION struct{
ReserveSize SIZE_T //col:238
ZeroBits SIZE_T //col:239
StackBase PVOID //col:240
}


type PROCESS_STACK_ALLOCATION_INFORMATION_EX struct{
PreferredNode uint32 //col:244
Reserved0 uint32 //col:245
Reserved1 uint32 //col:246
Reserved2 uint32 //col:247
AllocInfo PROCESS_STACK_ALLOCATION_INFORMATION //col:248
}


type PROCESS_HANDLE_INFORMATION struct{
HandleCount uint32 //col:252
HandleCountHighWatermark uint32 //col:253
}


type PROCESS_CYCLE_TIME_INFORMATION struct{
AccumulatedCycles ULONGLONG //col:257
CurrentCycleCount ULONGLONG //col:258
}


type PROCESS_WINDOW_INFORMATION struct{
WindowFlags uint32 //col:262
WindowTitleLength USHORT //col:263
WindowTitle[1] WCHAR //col:264
}


type PROCESS_HANDLE_TABLE_ENTRY_INFO struct{
HandleValue HANDLE //col:268
HandleCount ULONG_PTR //col:269
PointerCount ULONG_PTR //col:270
GrantedAccess uint32 //col:271
ObjectTypeIndex uint32 //col:272
HandleAttributes uint32 //col:273
Reserved uint32 //col:274
}


type PROCESS_HANDLE_SNAPSHOT_INFORMATION struct{
NumberOfHandles ULONG_PTR //col:278
Reserved ULONG_PTR //col:279
Handles[1] PROCESS_HANDLE_TABLE_ENTRY_INFO //col:280
}


type PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY struct{
Flags uint32 //col:285
EnforceRedirectionTrust uint32 //col:287
AuditRedirectionTrust uint32 //col:288
ReservedFlags uint32 //col:289
}


type PROCESS_MITIGATION_POLICY_INFORMATION struct{
Policy PROCESS_MITIGATION_POLICY //col:295
Union union //col:296
ASLRPolicy PROCESS_MITIGATION_ASLR_POLICY //col:298
StrictHandleCheckPolicy PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY //col:299
SystemCallDisablePolicy PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY //col:300
ExtensionPointDisablePolicy PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY //col:301
DynamicCodePolicy PROCESS_MITIGATION_DYNAMIC_CODE_POLICY //col:302
ControlFlowGuardPolicy PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY //col:303
SignaturePolicy PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY //col:304
FontDisablePolicy PROCESS_MITIGATION_FONT_DISABLE_POLICY //col:305
ImageLoadPolicy PROCESS_MITIGATION_IMAGE_LOAD_POLICY //col:306
SystemCallFilterPolicy PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY //col:307
PayloadRestrictionPolicy PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY //col:308
ChildProcessPolicy PROCESS_MITIGATION_CHILD_PROCESS_POLICY //col:309
SideChannelIsolationPolicy PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY //col:310
UserShadowStackPolicy PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY //col:311
RedirectionTrustPolicy PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY //col:312
}


type PROCESS_KEEPALIVE_COUNT_INFORMATION struct{
WakeCount uint32 //col:317
NoWakeCount uint32 //col:318
}


type PROCESS_REVOKE_FILE_HANDLES_INFORMATION struct{
TargetDevicePath UNICODE_STRING //col:322
}


type PROCESS_WORKING_SET_CONTROL struct{
Version uint32 //col:326
Operation PROCESS_WORKING_SET_OPERATION //col:327
Flags uint32 //col:328
}


type PS_PROTECTION struct{
Union union //col:332
Level uint8 //col:334
Struct struct //col:335
Type uint8 //col:337
Audit uint8 //col:338
Signer uint8 //col:339
}


type PROCESS_FAULT_INFORMATION struct{
FaultFlags uint32 //col:345
AdditionalInfo uint32 //col:346
}


type PROCESS_TELEMETRY_ID_INFORMATION struct{
HeaderSize uint32 //col:350
ProcessId uint32 //col:351
ProcessStartKey ULONGLONG //col:352
CreateTime ULONGLONG //col:353
CreateInterruptTime ULONGLONG //col:354
CreateUnbiasedInterruptTime ULONGLONG //col:355
ProcessSequenceNumber ULONGLONG //col:356
SessionCreateTime ULONGLONG //col:357
SessionId uint32 //col:358
BootId uint32 //col:359
ImageChecksum uint32 //col:360
ImageTimeDateStamp uint32 //col:361
UserSidOffset uint32 //col:362
ImagePathOffset uint32 //col:363
PackageNameOffset uint32 //col:364
RelativeAppNameOffset uint32 //col:365
CommandLineOffset uint32 //col:366
}


type PROCESS_COMMIT_RELEASE_INFORMATION struct{
Version uint32 //col:370
Struct struct //col:371
Eligible uint32 //col:373
ReleaseRepurposedMemResetCommit uint32 //col:374
ForceReleaseMemResetCommit uint32 //col:375
Spare uint32 //col:376
}


type PROCESS_JOB_MEMORY_INFO struct{
SharedCommitUsage ULONGLONG //col:384
PrivateCommitUsage ULONGLONG //col:385
PeakPrivateCommitUsage ULONGLONG //col:386
PrivateCommitLimit ULONGLONG //col:387
TotalCommitLimit ULONGLONG //col:388
}


type PROCESS_CHILD_PROCESS_INFORMATION struct{
ProhibitChildProcesses bool //col:392
AlwaysAllowSecureChildProcess bool //col:393
AuditProhibitChildProcesses bool //col:394
}


type POWER_THROTTLING_PROCESS_STATE struct{
Version uint32 //col:398
ControlMask uint32 //col:399
StateMask uint32 //col:400
}


type WIN32K_SYSCALL_FILTER struct{
FilterState uint32 //col:404
FilterSet uint32 //col:405
}


type PROCESS_WAKE_INFORMATION struct{
NotificationChannel ULONGLONG //col:409
WakeCounters[7] uint32 //col:410
_JOBOBJECT_WAKE_FILTER* struct //col:411
}


type PROCESS_ENERGY_TRACKING_STATE struct{
StateUpdateMask uint32 //col:415
StateDesiredValue uint32 //col:416
StateSequence uint32 //col:417
UpdateTag uint32 //col:418
Tag[64] WCHAR //col:419
}


type MANAGE_WRITES_TO_EXECUTABLE_MEMORY struct{
Version uint32 //col:423
ProcessEnableWriteExceptions uint32 //col:424
ThreadAllowWrites uint32 //col:425
Spare uint32 //col:426
KernelWriteToExecutableSignal PVOID //col:427
}


type POWER_THROTTLING_THREAD_STATE struct{
Version uint32 //col:431
ControlMask uint32 //col:432
StateMask uint32 //col:433
}


type PROCESS_UPTIME_INFORMATION struct{
QueryInterruptTime ULONGLONG //col:437
QueryUnbiasedTime ULONGLONG //col:438
EndInterruptTime ULONGLONG //col:439
TimeSinceCreation ULONGLONG //col:440
Uptime ULONGLONG //col:441
SuspendedTime ULONGLONG //col:442
Union union //col:443
HangCount uint32 //col:445
GhostCount uint32 //col:446
Crashed uint32 //col:447
Terminated uint32 //col:448
}


type PROCESS_SECURITY_DOMAIN_INFORMATION struct{
SecurityDomain ULONGLONG //col:453
}


type PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION struct{
ProcessHandle HANDLE //col:457
}


type PROCESS_LEAP_SECOND_INFORMATION struct{
Flags uint32 //col:461
Reserved uint32 //col:462
}


type PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION struct{
ReserveSize ULONGLONG //col:466
CommitSize ULONGLONG //col:467
PreferredNode uint32 //col:468
Reserved uint32 //col:469
Ssp PVOID //col:470
}


type PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION struct{
Ssp PVOID //col:474
}


type THREAD_BASIC_INFORMATION struct{
ExitStatus NTSTATUS //col:478
TebBaseAddress PTEB //col:479
ClientId CLIENT_ID //col:480
AffinityMask KAFFINITY //col:481
Priority KPRIORITY //col:482
BasePriority KPRIORITY //col:483
}


type THREAD_LAST_SYSCALL_INFORMATION struct{
FirstArgument PVOID //col:487
SystemCallNumber USHORT //col:488
#ifdefWin64 #ifdef WIN64 //col:489
Pad[0x3] USHORT //col:490
#else #else //col:491
Pad[0x1] USHORT //col:492
#endif #endif //col:493
WaitTime ULONG64 //col:494
}


type THREAD_CYCLE_TIME_INFORMATION struct{
AccumulatedCycles ULONGLONG //col:498
CurrentCycleCount ULONGLONG //col:499
}


type THREAD_TEB_INFORMATION struct{
TebInformation PVOID //col:503
TebOffset uint32 //col:504
BytesToRead uint32 //col:505
}


type COUNTER_READING struct{
Type HARDWARE_COUNTER_TYPE //col:509
Index uint32 //col:510
Start ULONG64 //col:511
Total ULONG64 //col:512
}


type THREAD_PERFORMANCE_DATA struct{
Size USHORT //col:516
Version USHORT //col:517
ProcessorNumber PROCESSOR_NUMBER //col:518
ContextSwitches uint32 //col:519
HwCountersCount uint32 //col:520
UpdateCount ULONG64 //col:521
WaitReasonBitMap ULONG64 //col:522
HardwareCounters ULONG64 //col:523
CycleTime COUNTER_READING //col:524
HwCounters[MAX_HW_COUNTERS] COUNTER_READING //col:525
}


type THREAD_PROFILING_INFORMATION struct{
HardwareCounters ULONG64 //col:529
Flags uint32 //col:530
Enable uint32 //col:531
PerformanceData PTHREAD_PERFORMANCE_DATA //col:532
}


type RTL_UMS_CONTEXT struct{
Link SINGLE_LIST_ENTRY //col:536
Context CONTEXT //col:537
Teb PVOID //col:538
UserContext PVOID //col:539
ULONG volatile //col:540
ULONG volatile //col:541
ULONG volatile //col:542
ULONG volatile //col:543
ULONG volatile //col:544
ULONG volatile //col:545
ULONG volatile //col:546
LONG volatile //col:547
ULONG64 volatile //col:548
ULONG64 volatile //col:549
ULONG64 volatile //col:550
_RTL_UMS_CONTEXT* struct //col:551
SwitchCount uint32 //col:552
KernelYieldCount uint32 //col:553
MixedYieldCount uint32 //col:554
YieldCount uint32 //col:555
}


type RTL_UMS_COMPLETION_LIST struct{
ThreadListHead PSINGLE_LIST_ENTRY //col:559
CompletionEvent PVOID //col:560
CompletionFlags uint32 //col:561
InternalListHead SINGLE_LIST_ENTRY //col:562
}


type THREAD_UMS_INFORMATION struct{
Command THREAD_UMS_INFORMATION_COMMAND //col:566
CompletionList PRTL_UMS_COMPLETION_LIST //col:567
UmsContext PRTL_UMS_CONTEXT //col:568
Union union //col:569
Flags uint32 //col:571
Struct struct //col:572
IsUmsSchedulerThread uint32 //col:574
IsUmsWorkerThread uint32 //col:575
SpareBits uint32 //col:576
}


type THREAD_NAME_INFORMATION struct{
ThreadName UNICODE_STRING //col:582
}


type ALPC_WORK_ON_BEHALF_TICKET struct{
ThreadId uint32 //col:586
ThreadCreationTimeLow uint32 //col:587
}


type RTL_WORK_ON_BEHALF_TICKET_EX struct{
Ticket ALPC_WORK_ON_BEHALF_TICKET //col:591
Union union //col:592
Flags uint32 //col:594
Struct struct //col:595
CurrentThread uint32 //col:597
Reserved1 uint32 //col:598
}


type PROC_THREAD_ATTRIBUTE { struct{
Attribute ULONG_PTR //col:604
Size SIZE_T //col:605
Value ULONG_PTR //col:606
}


type PROC_THREAD_ATTRIBUTE_LIST { struct{
PresentFlags uint32 //col:609
AttributeCount uint32 //col:610
LastAttribute uint32 //col:611
SpareUlong0 uint32 //col:612
ExtendedFlagsAttribute PPROC_THREAD_ATTRIBUTE //col:613
Attributes[1] PROC_THREAD_ATTRIBUTE //col:614
}


type SE_SAFE_OPEN_PROMPT_RESULTS { struct{
Results SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS //col:617
Path[MAX_PATH] WCHAR //col:618
}


type PROC_THREAD_BNOISOLATION_ATTRIBUTE struct{
IsolationEnabled BOOL //col:622
IsolationPrefix[0x88] WCHAR //col:623
}


type ISOLATION_MANIFEST_PROPERTIES { struct{
InstancePath UNICODE_STRING //col:626
FriendlyName UNICODE_STRING //col:627
Description UNICODE_STRING //col:628
Level ULONG_PTR //col:629
}


type PS_ATTRIBUTE struct{
Attribute ULONG_PTR //col:633
Size SIZE_T //col:634
Union union //col:635
Value ULONG_PTR //col:637
ValuePtr PVOID //col:638
}


type PS_ATTRIBUTE_LIST struct{
TotalLength SIZE_T //col:644
Attributes[1] PS_ATTRIBUTE //col:645
}


type PS_MEMORY_RESERVE struct{
ReserveAddress PVOID //col:649
ReserveSize SIZE_T //col:650
}


type PS_STD_HANDLE_INFO struct{
Union union //col:654
Flags uint32 //col:656
Struct struct //col:657
StdHandleState uint32 //col:659
PseudoHandleMask uint32 //col:660
}


type PS_TRUSTLET_ATTRIBUTE_TYPE struct{
Union union //col:667
Struct struct //col:669
Version uint8 //col:671
DataCount uint8 //col:672
SemanticType uint8 //col:673
AccessRights PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS //col:674
}


type PS_TRUSTLET_ATTRIBUTE_HEADER struct{
AttributeType PS_TRUSTLET_ATTRIBUTE_TYPE //col:681
InstanceNumber uint32 //col:682
Reserved uint32 //col:683
}


type PS_TRUSTLET_ATTRIBUTE_DATA struct{
Header PS_TRUSTLET_ATTRIBUTE_HEADER //col:687
Data[1] ULONGLONG //col:688
}


type PS_TRUSTLET_CREATE_ATTRIBUTES struct{
TrustletIdentity ULONGLONG //col:692
Attributes[1] PS_TRUSTLET_ATTRIBUTE_DATA //col:693
}


type PS_BNO_ISOLATION_PARAMETERS struct{
IsolationPrefix UNICODE_STRING //col:697
HandleCount uint32 //col:698
*Handles PVOID //col:699
IsolationEnabled bool //col:700
}


type PS_CREATE_INFO struct{
Size SIZE_T //col:704
State PS_CREATE_STATE //col:705
Union union //col:706
Struct struct //col:708
Union union //col:710
InitFlags uint32 //col:712
Struct struct //col:713
WriteOutputOnExit uint8 //col:715
DetectManifest uint8 //col:716
IFEOSkipDebugger uint8 //col:717
IFEODoNotPropagateKeyState uint8 //col:718
SpareBits1 uint8 //col:719
SpareBits2 uint8 //col:720
ProhibitedImageCharacteristics USHORT //col:721
}


type JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION struct{
BasicInfo JOBOBJECT_BASIC_ACCOUNTING_INFORMATION //col:769
IoInfo IO_COUNTERS //col:770
DiskIoInfo PROCESS_DISK_COUNTERS //col:771
ContextSwitches ULONG64 //col:772
TotalCycleTime LARGE_INTEGER //col:773
ReadyTime ULONG64 //col:774
EnergyValues PROCESS_ENERGY_VALUES //col:775
}


type JOBOBJECT_WAKE_INFORMATION struct{
NotificationChannel HANDLE //col:779
WakeCounters[7] ULONG64 //col:780
}


type JOBOBJECT_WAKE_INFORMATION_V1 struct{
NotificationChannel HANDLE //col:784
WakeCounters[4] ULONG64 //col:785
}


type JOBOBJECT_INTERFERENCE_INFORMATION struct{
Count ULONG64 //col:789
}


type JOBOBJECT_WAKE_FILTER struct{
HighEdgeFilter uint32 //col:793
LowEdgeFilter uint32 //col:794
}


type JOBOBJECT_FREEZE_INFORMATION struct{
Union union //col:798
Flags uint32 //col:800
Struct struct //col:801
FreezeOperation uint32 //col:803
FilterOperation uint32 //col:804
SwapOperation uint32 //col:805
Reserved uint32 //col:806
}


type JOBOBJECT_MEMORY_USAGE_INFORMATION struct{
JobMemory ULONG64 //col:816
PeakJobMemoryUsed ULONG64 //col:817
}


type JOBOBJECT_MEMORY_USAGE_INFORMATION_V2 struct{
BasicInfo JOBOBJECT_MEMORY_USAGE_INFORMATION //col:821
JobSharedMemory ULONG64 //col:822
Reserved[2] ULONG64 //col:823
}


type SILO_USER_SHARED_DATA struct{
ServiceSessionId ULONG64 //col:827
ActiveConsoleId uint32 //col:828
ConsoleSessionForegroundProcessId LONGLONG //col:829
NtProductType NT_PRODUCT_TYPE //col:830
SuiteMask uint32 //col:831
SharedUserSessionId uint32 //col:832
IsMultiSessionSku bool //col:833
NtSystemRoot[260] WCHAR //col:834
UserModeGlobalLogger[16] USHORT //col:835
}


type SILOOBJECT_ROOT_DIRECTORY struct{
ControlFlags uint32 //col:839
Path UNICODE_STRING //col:840
}


type JOBOBJECT_ENERGY_TRACKING_STATE struct{
Value ULONG64 //col:844
UpdateMask uint32 //col:845
DesiredState uint32 //col:846
}



type (
Ntpsapi interface{
NtCreateProcess()(ok bool)//col:374
#if_()(ok bool)//col:870
#if_()(ok bool)//col:1352
#if_()(ok bool)//col:1825
#if_()(ok bool)//col:2291
#if_()(ok bool)//col:2751
#if_()(ok bool)//col:3205
#if_()(ok bool)//col:3656
#if_()(ok bool)//col:4099
#if_()(ok bool)//col:4530
#if_()(ok bool)//col:4949
#if_()(ok bool)//col:5365
#if_()(ok bool)//col:5773
#if_()(ok bool)//col:6171
#if_()(ok bool)//col:6556
#if_()(ok bool)//col:6931
#if_()(ok bool)//col:7293
#if_()(ok bool)//col:7642
#if_()(ok bool)//col:7982
#if_()(ok bool)//col:8315
#if_()(ok bool)//col:8641
#if_()(ok bool)//col:8960
#if_()(ok bool)//col:9272
#if_()(ok bool)//col:9577
#if_()(ok bool)//col:9875
#if_()(ok bool)//col:10166
#if_()(ok bool)//col:10454
#if_()(ok bool)//col:10735
#if_()(ok bool)//col:11013
#if_()(ok bool)//col:11285
#if_()(ok bool)//col:11551
#if_()(ok bool)//col:11810
#if_()(ok bool)//col:12063
#if_()(ok bool)//col:12308
#if_()(ok bool)//col:12547
#if_()(ok bool)//col:12778
#if_()(ok bool)//col:13001
#if_()(ok bool)//col:13213
#if_()(ok bool)//col:13410
#if_()(ok bool)//col:13593
#if_()(ok bool)//col:13770
}
ntpsapi struct{}
)

func NewNtpsapi()Ntpsapi{ return & ntpsapi{} }

func (n *ntpsapi)NtCreateProcess()(ok bool){//col:374
/*NtCreateProcess(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ParentProcess,
    _In_ BOOLEAN InheritObjectTable,
    _In_opt_ HANDLE SectionHandle,
    _In_opt_ HANDLE DebugPort,
    _In_opt_ HANDLE TokenHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessEx(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ParentProcess,
    _In_ ULONG Flags, 
    _In_opt_ HANDLE SectionHandle,
    _In_opt_ HANDLE DebugPort,
    _In_opt_ HANDLE TokenHandle,
    _Reserved_ ULONG Reserved 
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenProcess(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:870
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtCreateProcessEx(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ParentProcess,
    _In_ ULONG Flags, 
    _In_opt_ HANDLE SectionHandle,
    _In_opt_ HANDLE DebugPort,
    _In_opt_ HANDLE TokenHandle,
    _Reserved_ ULONG Reserved 
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenProcess(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:1352
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtOpenProcess(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:1825
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtTerminateProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:2291
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSuspendProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:2751
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtResumeProcess(
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:3205
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtQueryInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:3656
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
    _Out_writes_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:4099
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtGetNextProcess(
    _In_opt_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewProcessHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:4530
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtGetNextThread(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Flags,
    _Out_ PHANDLE NewThreadHandle
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:4949
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSetInformationProcess(
    _In_ HANDLE ProcessHandle,
    _In_ PROCESSINFOCLASS ProcessInformationClass,
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:5365
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
    _In_reads_bytes_(ProcessInformationLength) PVOID ProcessInformation,
    _In_ ULONG ProcessInformationLength
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:5773
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtCreateProcessStateChange(
    _Out_ PHANDLE ProcessStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:6171
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtChangeProcessState(
    _In_ HANDLE ProcessStateChangeHandle,
    _In_ HANDLE ProcessHandle,
    _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:6556
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtCreateThreadStateChange(
    _Out_ PHANDLE ThreadStateChangeHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ThreadHandle,
    _In_opt_ ULONG64 Reserved
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:6931
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtChangeThreadState(
    _In_ HANDLE ThreadStateChangeHandle,
    _In_ HANDLE ThreadHandle,
    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
    _In_opt_ PVOID ExtendedInformation,
    _In_opt_ SIZE_T ExtendedInformationLength,
    _In_opt_ ULONG64 Reserved
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:7293
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtCreateThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _Out_ PCLIENT_ID ClientId,
    _In_ PCONTEXT ThreadContext,
    _In_ PINITIAL_TEB InitialTeb,
    _In_ BOOLEAN CreateSuspended
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:7642
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:7982
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtTerminateThread(
    _In_opt_ HANDLE ThreadHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:8315
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSuspendThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:8641
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:8960
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtGetCurrentProcessorNumber(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:9272
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtGetCurrentProcessorNumberEx(
    _Out_opt_ PPROCESSOR_NUMBER ProcessorNumber
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:9577
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtGetContextThread(
    _In_ HANDLE ThreadHandle,
    _Inout_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:9875
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSetContextThread(
    _In_ HANDLE ThreadHandle,
    _In_ PCONTEXT ThreadContext
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:10166
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtQueryInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:10454
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
    _Out_writes_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:10735
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSetInformationThread(
    _In_ HANDLE ThreadHandle,
    _In_ THREADINFOCLASS ThreadInformationClass,
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:11013
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
    _In_reads_bytes_(ThreadInformationLength) PVOID ThreadInformation,
    _In_ ULONG ThreadInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:11285
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtAlertThread(
    _In_ HANDLE ThreadHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:11551
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtAlertResumeThread(
    _In_ HANDLE ThreadHandle,
    _Out_opt_ PULONG PreviousSuspendCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:11810
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtTestAlert(
    VOID
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:12063
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtImpersonateThread(
    _In_ HANDLE ServerThreadHandle,
    _In_ HANDLE ClientThreadHandle,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:12308
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtRegisterThreadTerminatePort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:12547
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtSetLdtEntries(
    _In_ ULONG Selector0,
    _In_ ULONG Entry0Low,
    _In_ ULONG Entry0Hi,
    _In_ ULONG Selector1,
    _In_ ULONG Entry1Low,
    _In_ ULONG Entry1Hi
    );
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:12778
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
typedef VOID (*PPS_APC_ROUTINE)(
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:13001
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtQueueApcThread(
    _In_ HANDLE ThreadHandle,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:13213
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtQueueApcThreadEx(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:13410
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
    _In_opt_ HANDLE ReserveHandle, 
    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:13593
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtAlertThreadByThreadId(
    _In_ HANDLE ThreadId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}

func (n *ntpsapi)#if_()(ok bool){//col:13770
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle,
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess,
    _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags, 
    _In_ ULONG ThreadFlags, 
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ProcessHandle,
    _In_ PVOID StartRoutine, 
    _In_opt_ PVOID Argument,
    _In_ ULONG CreateFlags, 
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(
    _Out_ PHANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(
    _In_ HANDLE JobHandle,
    _In_ HANDLE ProcessHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(
    _In_ HANDLE JobHandle,
    _In_ NTSTATUS ExitStatus
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(
    _In_ HANDLE ProcessHandle,
    _In_opt_ HANDLE JobHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(
    _In_opt_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _Out_writes_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(
    _In_ HANDLE JobHandle,
    _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
    _In_reads_bytes_(JobObjectInformationLength) PVOID JobObjectInformation,
    _In_ ULONG JobObjectInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(
    _In_ ULONG NumJob,
    _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
    _In_ ULONG Flags
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(
    VOID
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(
    _Out_ PHANDLE MemoryReserveHandle,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ MEMORY_RESERVE_TYPE Type
    );
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(
    _Out_ PHANDLE SnapshotHandle,
    _In_ HANDLE ProcessHandle,
    _In_ ULONG CaptureFlags,
    _In_ ULONG ThreadContextFlags
    );
#endif
#endif
#endif
NtWaitForAlertByThreadId(
    _In_ PVOID Address,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#endif
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS
{
    UCHAR Trustlet : 1;
    UCHAR Ntos : 1;
    UCHAR WriteHandle : 1;
    UCHAR ReadHandle : 1;
    UCHAR Reserved : 4;
    UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;*/
return true
}



