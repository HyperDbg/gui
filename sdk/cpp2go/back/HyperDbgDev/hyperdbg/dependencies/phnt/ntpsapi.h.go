package phnt


const(
_NTPSAPI_H =  //col:13
PROCESS_TERMINATE = 0x0001 //col:16
PROCESS_CREATE_THREAD = 0x0002 //col:17
PROCESS_SET_SESSIONID = 0x0004 //col:18
PROCESS_VM_OPERATION = 0x0008 //col:19
PROCESS_VM_READ = 0x0010 //col:20
PROCESS_VM_WRITE = 0x0020 //col:21
PROCESS_CREATE_PROCESS = 0x0080 //col:22
PROCESS_SET_QUOTA = 0x0100 //col:23
PROCESS_SET_INFORMATION = 0x0200 //col:24
PROCESS_QUERY_INFORMATION = 0x0400 //col:25
PROCESS_SET_PORT = 0x0800 //col:26
PROCESS_SUSPEND_RESUME = 0x0800 //col:27
PROCESS_QUERY_LIMITED_INFORMATION = 0x1000 //col:28
PROCESS_SET_PORT = 0x0800 //col:31
THREAD_QUERY_INFORMATION = 0x0040 //col:36
THREAD_SET_THREAD_TOKEN = 0x0080 //col:37
THREAD_IMPERSONATE = 0x0100 //col:38
THREAD_DIRECT_IMPERSONATION = 0x0200 //col:39
THREAD_ALERT = 0x0004 //col:42
JOB_OBJECT_ASSIGN_PROCESS = 0x0001 //col:47
JOB_OBJECT_SET_ATTRIBUTES = 0x0002 //col:48
JOB_OBJECT_QUERY = 0x0004 //col:49
JOB_OBJECT_TERMINATE = 0x0008 //col:50
JOB_OBJECT_SET_SECURITY_ATTRIBUTES = 0x0010 //col:51
JOB_OBJECT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3F) //col:52
GDI_HANDLE_BUFFER_SIZE32 = 34 //col:56
GDI_HANDLE_BUFFER_SIZE64 = 60 //col:57
GDI_HANDLE_BUFFER_SIZE = GDI_HANDLE_BUFFER_SIZE32 //col:60
GDI_HANDLE_BUFFER_SIZE = GDI_HANDLE_BUFFER_SIZE64 //col:62
FLS_MAXIMUM_AVAILABLE = 128 //col:71
TLS_MINIMUM_AVAILABLE = 64 //col:74
TLS_EXPANSION_SLOTS = 1024 //col:77
PROCESS_EXCEPTION_PORT_ALL_STATE_BITS = 0x00000003 //col:400
PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS = ((ULONG_PTR)((1UL << PROCESS_EXCEPTION_PORT_ALL_STATE_BITS) - 1)) //col:401
PROCESS_PRIORITY_CLASS_UNKNOWN = 0 //col:443
PROCESS_PRIORITY_CLASS_IDLE = 1 //col:444
PROCESS_PRIORITY_CLASS_NORMAL = 2 //col:445
PROCESS_PRIORITY_CLASS_HIGH = 3 //col:446
PROCESS_PRIORITY_CLASS_REALTIME = 4 //col:447
PROCESS_PRIORITY_CLASS_BELOW_NORMAL = 5 //col:448
PROCESS_PRIORITY_CLASS_ABOVE_NORMAL = 6 //col:449
PROCESS_LUID_DOSDEVICES_ONLY = 0x00000001 //col:480
PROCESS_HANDLE_EXCEPTIONS_ENABLED = 0x00000001 //col:504
PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_DISABLED = 0x00000000 //col:506
PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_ENABLED = 0x00000001 //col:507
PROCESS_HANDLE_TRACING_MAX_SLOTS = 0x20000 //col:514
PROCESS_HANDLE_TRACING_MAX_STACKS = 16 //col:522
PROCESS_HANDLE_TRACE_TYPE_OPEN = 1 //col:524
PROCESS_HANDLE_TRACE_TYPE_CLOSE = 2 //col:525
PROCESS_HANDLE_TRACE_TYPE_BADREF = 3 //col:526
PS_PROTECTED_SIGNER_MASK = 0xFF //col:753
PS_PROTECTED_AUDIT_MASK = 0x08 //col:754
PS_PROTECTED_TYPE_MASK = 0x07 //col:755
PsProtectedValue(aSigner, aAudit, aType) ( = ((aSigner & PS_PROTECTED_SIGNER_MASK) << 4) | ((aAudit & PS_PROTECTED_AUDIT_MASK) << 3) | (aType & PS_PROTECTED_TYPE_MASK) ) //col:758
InitializePsProtection(aProtectionLevelPtr, aSigner, aAudit, aType) { = (aProtectionLevelPtr)->Signer = aSigner; (aProtectionLevelPtr)->Audit = aAudit; (aProtectionLevelPtr)->Type = aType; } //col:765
POWER_THROTTLING_PROCESS_CURRENT_VERSION = 1 //col:843
POWER_THROTTLING_PROCESS_EXECUTION_SPEED = 0x1 //col:844
POWER_THROTTLING_PROCESS_DELAYTIMERS = 0x2 //col:845
POWER_THROTTLING_PROCESS_VALID_FLAGS = ((POWER_THROTTLING_PROCESS_EXECUTION_SPEED | POWER_THROTTLING_PROCESS_DELAYTIMERS)) //col:846
WIN32K_SYSCALL_FILTER_STATE_ENABLE = 0x1 //col:856
WIN32K_SYSCALL_FILTER_STATE_AUDIT = 0x2 //col:857
POWER_THROTTLING_THREAD_CURRENT_VERSION = 1 //col:891
POWER_THROTTLING_THREAD_EXECUTION_SPEED = 0x1 //col:892
POWER_THROTTLING_THREAD_VALID_FLAGS = (POWER_THROTTLING_THREAD_EXECUTION_SPEED) //col:893
PROCESS_READWRITEVM_LOGGING_ENABLE_READVM = 1 //col:902
PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM = 2 //col:903
PROCESS_READWRITEVM_LOGGING_ENABLE_READVM_V = 1UL //col:904
PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM_V = 2UL //col:905
PROCESS_CREATE_FLAGS_BREAKAWAY = 0x00000001 // NtCreateProcessEx & NtCreateUserProcess //col:1215
PROCESS_CREATE_FLAGS_NO_DEBUG_INHERIT = 0x00000002 // NtCreateProcessEx & NtCreateUserProcess //col:1216
PROCESS_CREATE_FLAGS_INHERIT_HANDLES = 0x00000004 // NtCreateProcessEx & NtCreateUserProcess //col:1217
PROCESS_CREATE_FLAGS_OVERRIDE_ADDRESS_SPACE = 0x00000008 // NtCreateProcessEx only //col:1218
PROCESS_CREATE_FLAGS_LARGE_PAGES = 0x00000010 // NtCreateProcessEx only, requires SeLockMemory //col:1219
PROCESS_CREATE_FLAGS_LARGE_PAGE_SYSTEM_DLL = 0x00000020 // NtCreateProcessEx only, requires SeLockMemory //col:1220
PROCESS_CREATE_FLAGS_PROTECTED_PROCESS = 0x00000040 // NtCreateUserProcess only //col:1221
PROCESS_CREATE_FLAGS_CREATE_SESSION = 0x00000080 // NtCreateProcessEx & NtCreateUserProcess, requires SeLoadDriver //col:1222
PROCESS_CREATE_FLAGS_INHERIT_FROM_PARENT = 0x00000100 // NtCreateProcessEx & NtCreateUserProcess //col:1223
PROCESS_CREATE_FLAGS_SUSPENDED = 0x00000200 // NtCreateProcessEx & NtCreateUserProcess //col:1224
PROCESS_CREATE_FLAGS_FORCE_BREAKAWAY = 0x00000400 // NtCreateProcessEx & NtCreateUserProcess, requires SeTcb //col:1225
PROCESS_CREATE_FLAGS_MINIMAL_PROCESS = 0x00000800 // NtCreateProcessEx only //col:1226
PROCESS_CREATE_FLAGS_RELEASE_SECTION = 0x00001000 // NtCreateProcessEx & NtCreateUserProcess //col:1227
PROCESS_CREATE_FLAGS_CLONE_MINIMAL = 0x00002000 // NtCreateProcessEx only //col:1228
PROCESS_CREATE_FLAGS_CLONE_MINIMAL_REDUCED_COMMIT = 0x00004000 // //col:1229
PROCESS_CREATE_FLAGS_AUXILIARY_PROCESS = 0x00008000 // NtCreateProcessEx & NtCreateUserProcess, requires SeTcb //col:1230
PROCESS_CREATE_FLAGS_CREATE_STORE = 0x00020000 // NtCreateProcessEx & NtCreateUserProcess //col:1231
PROCESS_CREATE_FLAGS_USE_PROTECTED_ENVIRONMENT = 0x00040000 // NtCreateProcessEx & NtCreateUserProcess //col:1232
NtCurrentProcess() = ((HANDLE)(LONG_PTR)-1) //col:1282
ZwCurrentProcess() = NtCurrentProcess() //col:1283
NtCurrentThread() = ((HANDLE)(LONG_PTR)-2) //col:1284
ZwCurrentThread() = NtCurrentThread() //col:1285
NtCurrentSession() = ((HANDLE)(LONG_PTR)-3) //col:1286
ZwCurrentSession() = NtCurrentSession() //col:1287
NtCurrentPeb() = (NtCurrentTeb()->ProcessEnvironmentBlock) //col:1288
NtCurrentProcessToken() = ((HANDLE)(LONG_PTR)-4) // NtOpenProcessToken(NtCurrentProcess()) //col:1291
NtCurrentThreadToken() = ((HANDLE)(LONG_PTR)-5) // NtOpenThreadToken(NtCurrentThread()) //col:1292
NtCurrentThreadEffectiveToken() = ((HANDLE)(LONG_PTR)-6) // NtOpenThreadToken(NtCurrentThread()) + NtOpenProcessToken(NtCurrentProcess()) //col:1293
NtCurrentSilo() = ((HANDLE)(LONG_PTR)-1) //col:1295
NtCurrentProcessId() = (NtCurrentTeb()->ClientId.UniqueProcess) //col:1298
NtCurrentThreadId() = (NtCurrentTeb()->ClientId.UniqueThread) //col:1299
PROCESS_GET_NEXT_FLAGS_PREVIOUS_PROCESS = 0x00000001 //col:1314
STATECHANGE_SET_ATTRIBUTES = 0x0001 //col:1354
Wow64EncodeApcRoutine(ApcRoutine) = ((PVOID)((0 - ((LONG_PTR)(ApcRoutine))) << 2)) //col:1585
Wow64DecodeApcRoutine(ApcRoutine) = ((PVOID)(0 - (((LONG_PTR)(ApcRoutine)) >> 2))) //col:1588
APC_FORCE_THREAD_SIGNAL = ((HANDLE)1) // ReserveHandle //col:1604
ProcThreadAttributeParentProcess = 0 // in HANDLE //col:1674
ProcThreadAttributeExtendedFlags = 1 // in ULONG (EXTENDED_PROCESS_CREATION_FLAG_*) //col:1675
ProcThreadAttributeHandleList = 2 // in HANDLE[] //col:1676
ProcThreadAttributeGroupAffinity = 3 // in GROUP_AFFINITY // since WIN7 //col:1677
ProcThreadAttributePreferredNode = 4 // in USHORT //col:1678
ProcThreadAttributeIdealProcessor = 5 // in PROCESSOR_NUMBER //col:1679
ProcThreadAttributeUmsThread = 6 // in UMS_CREATE_THREAD_ATTRIBUTES //col:1680
ProcThreadAttributeMitigationPolicy = 7 // in ULONG, ULONG64, or ULONG64[2] //col:1681
ProcThreadAttributePackageFullName = 8 // in WCHAR[] // since WIN8 //col:1682
ProcThreadAttributeSecurityCapabilities = 9 // in SECURITY_CAPABILITIES //col:1683
ProcThreadAttributeConsoleReference = 10 // BaseGetConsoleReference (kernelbase.dll) //col:1684
ProcThreadAttributeProtectionLevel = 11 // in ULONG (PROTECTION_LEVEL_*) // since WINBLUE //col:1685
ProcThreadAttributeOsMaxVersionTested = 12 // in MAXVERSIONTESTED_INFO // since THRESHOLD // (from exe.manifest) //col:1686
ProcThreadAttributeJobList = 13 // in HANDLE[] //col:1687
ProcThreadAttributeChildProcessPolicy = 14 // in ULONG (PROCESS_CREATION_CHILD_PROCESS_*) // since THRESHOLD2 //col:1688
ProcThreadAttributeAllApplicationPackagesPolicy = 15 // in ULONG (PROCESS_CREATION_ALL_APPLICATION_PACKAGES_*) // since REDSTONE //col:1689
ProcThreadAttributeWin32kFilter = 16 // in WIN32K_SYSCALL_FILTER //col:1690
ProcThreadAttributeSafeOpenPromptOriginClaim = 17 // in SE_SAFE_OPEN_PROMPT_RESULTS //col:1691
ProcThreadAttributeDesktopAppPolicy = 18 // in ULONG (PROCESS_CREATION_DESKTOP_APP_*) // since RS2 //col:1692
ProcThreadAttributeBnoIsolation = 19 // in PROC_THREAD_BNOISOLATION_ATTRIBUTE //col:1693
ProcThreadAttributePseudoConsole = 22 // in HANDLE (HPCON) // since RS5 //col:1694
ProcThreadAttributeIsolationManifest = 23 // in ISOLATION_MANIFEST_PROPERTIES // rev (diversenok) // since 19H2+ //col:1695
ProcThreadAttributeMitigationAuditPolicy = 24 // in ULONG, ULONG64, or ULONG64[2] // since 21H1 //col:1696
ProcThreadAttributeMachineType = 25 // in USHORT // since 21H2 //col:1697
ProcThreadAttributeComponentFilter = 26 // in ULONG //col:1698
ProcThreadAttributeEnableOptionalXStateFeatures = 27 // in ULONG64 // since WIN11 //col:1699
ProcThreadAttributeCreateStore = 28 // ULONG // rev (diversenok) //col:1700
PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS = ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE) //col:1703
PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME = ProcThreadAttributeValue(ProcThreadAttributePackageFullName, FALSE, TRUE, FALSE) //col:1707
PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE = ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE) //col:1711
PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED = ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE) //col:1715
PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:1719
PROC_THREAD_ATTRIBUTE_BNO_ISOLATION = ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:1723
PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST = ProcThreadAttributeValue(ProcThreadAttributeIsolationManifest, FALSE, TRUE, FALSE) //col:1727
PROC_THREAD_ATTRIBUTE_CREATE_STORE = ProcThreadAttributeValue(ProcThreadAttributeCreateStore, FALSE, TRUE, FALSE) //col:1731
EXTENDED_PROCESS_CREATION_FLAG_ELEVATION_HANDLED = 0x00000001 //col:1753
EXTENDED_PROCESS_CREATION_FLAG_FORCELUA = 0x00000002 //col:1754
EXTENDED_PROCESS_CREATION_FLAG_FORCE_BREAKAWAY = 0x00000004 // requires SeTcbPrivilege // since WINBLUE //col:1755
PROTECTION_LEVEL_WINTCB_LIGHT = 0x00000000 //col:1758
PROTECTION_LEVEL_WINDOWS = 0x00000001 //col:1759
PROTECTION_LEVEL_WINDOWS_LIGHT = 0x00000002 //col:1760
PROTECTION_LEVEL_ANTIMALWARE_LIGHT = 0x00000003 //col:1761
PROTECTION_LEVEL_LSA_LIGHT = 0x00000004 //col:1762
PROTECTION_LEVEL_WINTCB = 0x00000005 //col:1763
PROTECTION_LEVEL_CODEGEN_LIGHT = 0x00000006 //col:1764
PROTECTION_LEVEL_AUTHENTICODE = 0x00000007 //col:1765
PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS = ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE) //col:1801
PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE = ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE) //col:1805
PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED = ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE) //col:1809
PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:1813
PROC_THREAD_ATTRIBUTE_BNO_ISOLATION = ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:1817
PS_ATTRIBUTE_NUMBER_MASK = 0x0000ffff //col:1861
PS_ATTRIBUTE_THREAD = 0x00010000 // may be used with thread creation //col:1862
PS_ATTRIBUTE_INPUT = 0x00020000 // input only //col:1863
PS_ATTRIBUTE_ADDITIVE = 0x00040000 // "accumulated" e.g. bitmasks, counters, etc. //col:1864
PsAttributeValue(Number, Thread, Input, Additive) = (((Number) & PS_ATTRIBUTE_NUMBER_MASK) | ((Thread) ? PS_ATTRIBUTE_THREAD : 0) | ((Input) ? PS_ATTRIBUTE_INPUT : 0) | ((Additive) ? PS_ATTRIBUTE_ADDITIVE : 0)) //col:1868
PS_ATTRIBUTE_PARENT_PROCESS = PsAttributeValue(PsAttributeParentProcess, FALSE, TRUE, TRUE) //col:1874
PS_ATTRIBUTE_DEBUG_OBJECT = PsAttributeValue(PsAttributeDebugObject, FALSE, TRUE, TRUE) //col:1876
PS_ATTRIBUTE_TOKEN = PsAttributeValue(PsAttributeToken, FALSE, TRUE, TRUE) //col:1878
PS_ATTRIBUTE_CLIENT_ID = PsAttributeValue(PsAttributeClientId, TRUE, FALSE, FALSE) //col:1880
PS_ATTRIBUTE_TEB_ADDRESS = PsAttributeValue(PsAttributeTebAddress, TRUE, FALSE, FALSE) //col:1882
PS_ATTRIBUTE_IMAGE_NAME = PsAttributeValue(PsAttributeImageName, FALSE, TRUE, FALSE) //col:1884
PS_ATTRIBUTE_IMAGE_INFO = PsAttributeValue(PsAttributeImageInfo, FALSE, FALSE, FALSE) //col:1886
PS_ATTRIBUTE_MEMORY_RESERVE = PsAttributeValue(PsAttributeMemoryReserve, FALSE, TRUE, FALSE) //col:1888
PS_ATTRIBUTE_PRIORITY_CLASS = PsAttributeValue(PsAttributePriorityClass, FALSE, TRUE, FALSE) //col:1890
PS_ATTRIBUTE_ERROR_MODE = PsAttributeValue(PsAttributeErrorMode, FALSE, TRUE, FALSE) //col:1892
PS_ATTRIBUTE_STD_HANDLE_INFO = PsAttributeValue(PsAttributeStdHandleInfo, FALSE, TRUE, FALSE) //col:1894
PS_ATTRIBUTE_HANDLE_LIST = PsAttributeValue(PsAttributeHandleList, FALSE, TRUE, FALSE) //col:1896
PS_ATTRIBUTE_GROUP_AFFINITY = PsAttributeValue(PsAttributeGroupAffinity, TRUE, TRUE, FALSE) //col:1898
PS_ATTRIBUTE_PREFERRED_NODE = PsAttributeValue(PsAttributePreferredNode, FALSE, TRUE, FALSE) //col:1900
PS_ATTRIBUTE_IDEAL_PROCESSOR = PsAttributeValue(PsAttributeIdealProcessor, TRUE, TRUE, FALSE) //col:1902
PS_ATTRIBUTE_UMS_THREAD = PsAttributeValue(PsAttributeUmsThread, TRUE, TRUE, FALSE) //col:1904
PS_ATTRIBUTE_MITIGATION_OPTIONS = PsAttributeValue(PsAttributeMitigationOptions, FALSE, TRUE, FALSE) //col:1906
PS_ATTRIBUTE_PROTECTION_LEVEL = PsAttributeValue(PsAttributeProtectionLevel, FALSE, TRUE, TRUE) //col:1908
PS_ATTRIBUTE_SECURE_PROCESS = PsAttributeValue(PsAttributeSecureProcess, FALSE, TRUE, FALSE) //col:1910
PS_ATTRIBUTE_JOB_LIST = PsAttributeValue(PsAttributeJobList, FALSE, TRUE, FALSE) //col:1912
PS_ATTRIBUTE_CHILD_PROCESS_POLICY = PsAttributeValue(PsAttributeChildProcessPolicy, FALSE, TRUE, FALSE) //col:1914
PS_ATTRIBUTE_ALL_APPLICATION_PACKAGES_POLICY = PsAttributeValue(PsAttributeAllApplicationPackagesPolicy, FALSE, TRUE, FALSE) //col:1916
PS_ATTRIBUTE_WIN32K_FILTER = PsAttributeValue(PsAttributeWin32kFilter, FALSE, TRUE, FALSE) //col:1918
PS_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM = PsAttributeValue(PsAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE) //col:1920
PS_ATTRIBUTE_BNO_ISOLATION = PsAttributeValue(PsAttributeBnoIsolation, FALSE, TRUE, FALSE) //col:1922
PS_ATTRIBUTE_DESKTOP_APP_POLICY = PsAttributeValue(PsAttributeDesktopAppPolicy, FALSE, TRUE, FALSE) //col:1924
PS_ATTRIBUTE_CHPE = PsAttributeValue(PsAttributeChpe, FALSE, TRUE, TRUE) //col:1926
PS_ATTRIBUTE_MITIGATION_AUDIT_OPTIONS = PsAttributeValue(PsAttributeMitigationAuditOptions, FALSE, TRUE, FALSE) //col:1928
PS_ATTRIBUTE_MACHINE_TYPE = PsAttributeValue(PsAttributeMachineType, FALSE, TRUE, TRUE) //col:1930
PS_ATTRIBUTE_COMPONENT_FILTER = PsAttributeValue(PsAttributeComponentFilter, FALSE, TRUE, FALSE) //col:1932
PS_ATTRIBUTE_ENABLE_OPTIONAL_XSTATE_FEATURES = PsAttributeValue(PsAttributeEnableOptionalXStateFeatures, TRUE, TRUE, FALSE) //col:1934
PS_STD_INPUT_HANDLE = 0x1 //col:1974
PS_STD_OUTPUT_HANDLE = 0x2 //col:1975
PS_STD_ERROR_HANDLE = 0x4 //col:1976
THREAD_CREATE_FLAGS_CREATE_SUSPENDED = 0x00000001 // NtCreateUserProcess & NtCreateThreadEx //col:2197
THREAD_CREATE_FLAGS_SKIP_THREAD_ATTACH = 0x00000002 // NtCreateThreadEx only //col:2198
THREAD_CREATE_FLAGS_HIDE_FROM_DEBUGGER = 0x00000004 // NtCreateThreadEx only //col:2199
THREAD_CREATE_FLAGS_LOADER_WORKER = 0x00000010 // NtCreateThreadEx only //col:2200
THREAD_CREATE_FLAGS_SKIP_LOADER_INIT = 0x00000020 // NtCreateThreadEx only //col:2201
THREAD_CREATE_FLAGS_BYPASS_PROCESS_FREEZE = 0x00000040 // NtCreateThreadEx only //col:2202
THREAD_CREATE_FLAGS_INITIAL_THREAD = 0x00000080 // ? //col:2203
JobObjectBasicAccountingInformation = 1 // JOBOBJECT_BASIC_ACCOUNTING_INFORMATION //col:2233
JobObjectBasicLimitInformation = 2 // JOBOBJECT_BASIC_LIMIT_INFORMATION //col:2234
JobObjectBasicProcessIdList = 3 // JOBOBJECT_BASIC_PROCESS_ID_LIST //col:2235
JobObjectBasicUIRestrictions = 4 // JOBOBJECT_BASIC_UI_RESTRICTIONS //col:2236
JobObjectSecurityLimitInformation = 5 // JOBOBJECT_SECURITY_LIMIT_INFORMATION //col:2237
JobObjectEndOfJobTimeInformation = 6 // JOBOBJECT_END_OF_JOB_TIME_INFORMATION //col:2238
JobObjectAssociateCompletionPortInformation = 7 // JOBOBJECT_ASSOCIATE_COMPLETION_PORT //col:2239
JobObjectBasicAndIoAccountingInformation = 8 // JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION //col:2240
JobObjectExtendedLimitInformation = 9 // JOBOBJECT_EXTENDED_LIMIT_INFORMATION //col:2241
JobObjectJobSetInformation = 10 // JOBOBJECT_JOBSET_INFORMATION //col:2242
JobObjectGroupInformation = 11 // USHORT //col:2243
JobObjectNotificationLimitInformation = 12 // JOBOBJECT_NOTIFICATION_LIMIT_INFORMATION //col:2244
JobObjectLimitViolationInformation = 13 // JOBOBJECT_LIMIT_VIOLATION_INFORMATION //col:2245
JobObjectGroupInformationEx = 14 // GROUP_AFFINITY (ARRAY) //col:2246
JobObjectCpuRateControlInformation = 15 // JOBOBJECT_CPU_RATE_CONTROL_INFORMATION //col:2247
JobObjectCompletionFilter = 16 //col:2248
JobObjectCompletionCounter = 17 //col:2249
JobObjectFreezeInformation = 18 // JOBOBJECT_FREEZE_INFORMATION //col:2250
JobObjectExtendedAccountingInformation = 19 // JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION //col:2251
JobObjectWakeInformation = 20 // JOBOBJECT_WAKE_INFORMATION //col:2252
JobObjectBackgroundInformation = 21 //col:2253
JobObjectSchedulingRankBiasInformation = 22 //col:2254
JobObjectTimerVirtualizationInformation = 23 //col:2255
JobObjectCycleTimeNotification = 24 //col:2256
JobObjectClearEvent = 25 //col:2257
JobObjectInterferenceInformation = 26 // JOBOBJECT_INTERFERENCE_INFORMATION //col:2258
JobObjectClearPeakJobMemoryUsed = 27 //col:2259
JobObjectMemoryUsageInformation = 28 // JOBOBJECT_MEMORY_USAGE_INFORMATION // JOBOBJECT_MEMORY_USAGE_INFORMATION_V2 //col:2260
JobObjectSharedCommit = 29 //col:2261
JobObjectContainerId = 30 //col:2262
JobObjectIoRateControlInformation = 31 //col:2263
JobObjectNetRateControlInformation = 32 // JOBOBJECT_NET_RATE_CONTROL_INFORMATION //col:2264
JobObjectNotificationLimitInformation2 = 33 // JOBOBJECT_NOTIFICATION_LIMIT_INFORMATION_2 //col:2265
JobObjectLimitViolationInformation2 = 34 // JOBOBJECT_LIMIT_VIOLATION_INFORMATION_2 //col:2266
JobObjectCreateSilo = 35 //col:2267
JobObjectSiloBasicInformation = 36 // SILOOBJECT_BASIC_INFORMATION //col:2268
JobObjectSiloRootDirectory = 37 // SILOOBJECT_ROOT_DIRECTORY //col:2269
JobObjectServerSiloBasicInformation = 38 // SERVERSILO_BASIC_INFORMATION //col:2270
JobObjectServerSiloUserSharedData = 39 // SILO_USER_SHARED_DATA //col:2271
JobObjectServerSiloInitialize = 40 //col:2272
JobObjectServerSiloRunningState = 41 //col:2273
JobObjectIoAttribution = 42 //col:2274
JobObjectMemoryPartitionInformation = 43 //col:2275
JobObjectContainerTelemetryId = 44 //col:2276
JobObjectSiloSystemRoot = 45 //col:2277
JobObjectEnergyTrackingState = 46 // JOBOBJECT_ENERGY_TRACKING_STATE //col:2278
JobObjectThreadImpersonationInformation = 47 //col:2279
JobObjectIoPriorityLimit = 48 //col:2280
JobObjectPagePriorityLimit = 49 //col:2281
MaxJobObjectInfoClass = 50 //col:2282
)

const(
    ProcessBasicInformation // q: PROCESS_BASIC_INFORMATION PROCESS_EXTENDED_BASIC_INFORMATION = 1  //col:116
    ProcessQuotaLimits // qs: QUOTA_LIMITS QUOTA_LIMITS_EX = 2  //col:117
    ProcessIoCounters // q: IO_COUNTERS = 3  //col:118
    ProcessVmCounters // q: VM_COUNTERS VM_COUNTERS_EX VM_COUNTERS_EX2 = 4  //col:119
    ProcessTimes // q: KERNEL_USER_TIMES = 5  //col:120
    ProcessBasePriority // s: KPRIORITY = 6  //col:121
    ProcessRaisePriority // s: ULONG = 7  //col:122
    ProcessDebugPort // q: HANDLE = 8  //col:123
    ProcessExceptionPort // s: PROCESS_EXCEPTION_PORT = 9  //col:124
    ProcessAccessToken // s: PROCESS_ACCESS_TOKEN = 10  //col:125
    ProcessLdtInformation // qs: PROCESS_LDT_INFORMATION // 10 = 11  //col:126
    ProcessLdtSize // s: PROCESS_LDT_SIZE = 12  //col:127
    ProcessDefaultHardErrorMode // qs: ULONG = 13  //col:128
    ProcessIoPortHandlers // (kernel-mode only) // PROCESS_IO_PORT_HANDLER_INFORMATION = 14  //col:129
    ProcessPooledUsageAndLimits // q: POOLED_USAGE_AND_LIMITS = 15  //col:130
    ProcessWorkingSetWatch // q: PROCESS_WS_WATCH_INFORMATION[]; s: void = 16  //col:131
    ProcessUserModeIOPL // qs: ULONG (requires SeTcbPrivilege) = 17  //col:132
    ProcessEnableAlignmentFaultFixup // s: BOOLEAN = 18  //col:133
    ProcessPriorityClass // qs: PROCESS_PRIORITY_CLASS = 19  //col:134
    ProcessWx86Information // qs: ULONG (requires SeTcbPrivilege) (VdmAllowed) = 20  //col:135
    ProcessHandleCount // q: ULONG PROCESS_HANDLE_INFORMATION // 20 = 21  //col:136
    ProcessAffinityMask // qs: KAFFINITY qs: GROUP_AFFINITY = 22  //col:137
    ProcessPriorityBoost // qs: ULONG = 23  //col:138
    ProcessDeviceMap // qs: PROCESS_DEVICEMAP_INFORMATION PROCESS_DEVICEMAP_INFORMATION_EX = 24  //col:139
    ProcessSessionInformation // q: PROCESS_SESSION_INFORMATION = 25  //col:140
    ProcessForegroundInformation // s: PROCESS_FOREGROUND_BACKGROUND = 26  //col:141
    ProcessWow64Information // q: ULONG_PTR = 27  //col:142
    ProcessImageFileName // q: UNICODE_STRING = 28  //col:143
    ProcessLUIDDeviceMapsEnabled // q: ULONG = 29  //col:144
    ProcessBreakOnTermination // qs: ULONG = 30  //col:145
    ProcessDebugObjectHandle // q: HANDLE // 30 = 31  //col:146
    ProcessDebugFlags // qs: ULONG = 32  //col:147
    ProcessHandleTracing // q: PROCESS_HANDLE_TRACING_QUERY; s: size 0 disables otherwise enables = 33  //col:148
    ProcessIoPriority // qs: IO_PRIORITY_HINT = 34  //col:149
    ProcessExecuteFlags // qs: ULONG = 35  //col:150
    ProcessTlsInformation // PROCESS_TLS_INFORMATION // ProcessResourceManagement  = 36  //col:151
    ProcessCookie // q: ULONG = 37  //col:152
    ProcessImageInformation // q: SECTION_IMAGE_INFORMATION = 38  //col:153
    ProcessCycleTime // q: PROCESS_CYCLE_TIME_INFORMATION // since VISTA = 39  //col:154
    ProcessPagePriority // q: PAGE_PRIORITY_INFORMATION = 40  //col:155
    ProcessInstrumentationCallback // s: PVOID or PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION // 40 = 41  //col:156
    ProcessThreadStackAllocation // s: PROCESS_STACK_ALLOCATION_INFORMATION PROCESS_STACK_ALLOCATION_INFORMATION_EX = 42  //col:157
    ProcessWorkingSetWatchEx // q: PROCESS_WS_WATCH_INFORMATION_EX[] = 43  //col:158
    ProcessImageFileNameWin32 // q: UNICODE_STRING = 44  //col:159
    ProcessImageFileMapping // q: HANDLE (input) = 45  //col:160
    ProcessAffinityUpdateMode // qs: PROCESS_AFFINITY_UPDATE_MODE = 46  //col:161
    ProcessMemoryAllocationMode // qs: PROCESS_MEMORY_ALLOCATION_MODE = 47  //col:162
    ProcessGroupInformation // q: USHORT[] = 48  //col:163
    ProcessTokenVirtualizationEnabled // s: ULONG = 49  //col:164
    ProcessConsoleHostProcess // q: ULONG_PTR // ProcessOwnerInformation = 50  //col:165
    ProcessWindowInformation // q: PROCESS_WINDOW_INFORMATION // 50 = 51  //col:166
    ProcessHandleInformation // q: PROCESS_HANDLE_SNAPSHOT_INFORMATION // since WIN8 = 52  //col:167
    ProcessMitigationPolicy // s: PROCESS_MITIGATION_POLICY_INFORMATION = 53  //col:168
    ProcessDynamicFunctionTableInformation = 54  //col:169
    ProcessHandleCheckingMode // qs: ULONG; s: 0 disables otherwise enables = 55  //col:170
    ProcessKeepAliveCount // q: PROCESS_KEEPALIVE_COUNT_INFORMATION = 56  //col:171
    ProcessRevokeFileHandles // s: PROCESS_REVOKE_FILE_HANDLES_INFORMATION = 57  //col:172
    ProcessWorkingSetControl // s: PROCESS_WORKING_SET_CONTROL = 58  //col:173
    ProcessHandleTable // q: ULONG[] // since WINBLUE = 59  //col:174
    ProcessCheckStackExtentsMode // qs: ULONG // KPROCESS->CheckStackExtents (CFG) = 60  //col:175
    ProcessCommandLineInformation // q: UNICODE_STRING // 60 = 61  //col:176
    ProcessProtectionInformation // q: PS_PROTECTION = 62  //col:177
    ProcessMemoryExhaustion // PROCESS_MEMORY_EXHAUSTION_INFO // since THRESHOLD = 63  //col:178
    ProcessFaultInformation // PROCESS_FAULT_INFORMATION = 64  //col:179
    ProcessTelemetryIdInformation // q: PROCESS_TELEMETRY_ID_INFORMATION = 65  //col:180
    ProcessCommitReleaseInformation // PROCESS_COMMIT_RELEASE_INFORMATION = 66  //col:181
    ProcessDefaultCpuSetsInformation // SYSTEM_CPU_SET_INFORMATION[5] = 67  //col:182
    ProcessAllowedCpuSetsInformation // SYSTEM_CPU_SET_INFORMATION[5] = 68  //col:183
    ProcessSubsystemProcess = 69  //col:184
    ProcessJobMemoryInformation // q: PROCESS_JOB_MEMORY_INFO = 70  //col:185
    ProcessInPrivate // s: void // ETW // since THRESHOLD2 // 70 = 71  //col:186
    ProcessRaiseUMExceptionOnInvalidHandleClose // qs: ULONG; s: 0 disables otherwise enables = 72  //col:187
    ProcessIumChallengeResponse = 73  //col:188
    ProcessChildProcessInformation // q: PROCESS_CHILD_PROCESS_INFORMATION = 74  //col:189
    ProcessHighGraphicsPriorityInformation // qs: BOOLEAN (requires SeTcbPrivilege) = 75  //col:190
    ProcessSubsystemInformation // q: SUBSYSTEM_INFORMATION_TYPE // since REDSTONE2 = 76  //col:191
    ProcessEnergyValues // q: PROCESS_ENERGY_VALUES PROCESS_EXTENDED_ENERGY_VALUES = 77  //col:192
    ProcessPowerThrottlingState // qs: POWER_THROTTLING_PROCESS_STATE = 78  //col:193
    ProcessReserved3Information // ProcessActivityThrottlePolicy // PROCESS_ACTIVITY_THROTTLE_POLICY = 79  //col:194
    ProcessWin32kSyscallFilterInformation // q: WIN32K_SYSCALL_FILTER = 80  //col:195
    ProcessDisableSystemAllowedCpuSets // 80 = 81  //col:196
    ProcessWakeInformation // PROCESS_WAKE_INFORMATION = 82  //col:197
    ProcessEnergyTrackingState // PROCESS_ENERGY_TRACKING_STATE = 83  //col:198
    ProcessManageWritesToExecutableMemory // MANAGE_WRITES_TO_EXECUTABLE_MEMORY // since REDSTONE3 = 84  //col:199
    ProcessCaptureTrustletLiveDump = 85  //col:200
    ProcessTelemetryCoverage = 86  //col:201
    ProcessEnclaveInformation = 87  //col:202
    ProcessEnableReadWriteVmLogging // PROCESS_READWRITEVM_LOGGING_INFORMATION = 88  //col:203
    ProcessUptimeInformation // q: PROCESS_UPTIME_INFORMATION = 89  //col:204
    ProcessImageSection // q: HANDLE = 90  //col:205
    ProcessDebugAuthInformation // since REDSTONE4 // 90 = 91  //col:206
    ProcessSystemResourceManagement // PROCESS_SYSTEM_RESOURCE_MANAGEMENT = 92  //col:207
    ProcessSequenceNumber // q: ULONGLONG = 93  //col:208
    ProcessLoaderDetour // since REDSTONE5 = 94  //col:209
    ProcessSecurityDomainInformation // PROCESS_SECURITY_DOMAIN_INFORMATION = 95  //col:210
    ProcessCombineSecurityDomainsInformation // PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION = 96  //col:211
    ProcessEnableLogging // PROCESS_LOGGING_INFORMATION = 97  //col:212
    ProcessLeapSecondInformation // PROCESS_LEAP_SECOND_INFORMATION = 98  //col:213
    ProcessFiberShadowStackAllocation // PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION // since 19H1 = 99  //col:214
    ProcessFreeFiberShadowStackAllocation // PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION = 100  //col:215
    ProcessAltSystemCallInformation // qs: BOOLEAN (kernel-mode only) // INT2E // since 20H1 // 100 = 101  //col:216
    ProcessDynamicEHContinuationTargets // PROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION = 102  //col:217
    ProcessDynamicEnforcedCetCompatibleRanges // PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE_INFORMATION // since 20H2 = 103  //col:218
    ProcessCreateStateChange // since WIN11 = 104  //col:219
    ProcessApplyStateChange = 105  //col:220
    ProcessEnableOptionalXStateFeatures = 106  //col:221
    ProcessAltPrefetchParam // since 22H1 = 107  //col:222
    ProcessAssignCpuPartitions = 108  //col:223
    ProcessPriorityClassEx = 109  //col:224
    ProcessMembershipInformation = 110  //col:225
    ProcessEffectiveIoPriority = 111  //col:226
    ProcessEffectivePagePriority = 112  //col:227
    MaxProcessInfoClass = 113  //col:228
)


const(
    ThreadBasicInformation // q: THREAD_BASIC_INFORMATION = 1  //col:235
    ThreadTimes // q: KERNEL_USER_TIMES = 2  //col:236
    ThreadPriority // s: KPRIORITY (requires SeIncreaseBasePriorityPrivilege) = 3  //col:237
    ThreadBasePriority // s: KPRIORITY = 4  //col:238
    ThreadAffinityMask // s: KAFFINITY = 5  //col:239
    ThreadImpersonationToken // s: HANDLE = 6  //col:240
    ThreadDescriptorTableEntry // q: DESCRIPTOR_TABLE_ENTRY (or WOW64_DESCRIPTOR_TABLE_ENTRY) = 7  //col:241
    ThreadEnableAlignmentFaultFixup // s: BOOLEAN = 8  //col:242
    ThreadEventPair = 9  //col:243
    ThreadQuerySetWin32StartAddress // q: ULONG_PTR = 10  //col:244
    ThreadZeroTlsCell // s: ULONG // TlsIndex // 10 = 11  //col:245
    ThreadPerformanceCount // q: LARGE_INTEGER = 12  //col:246
    ThreadAmILastThread // q: ULONG = 13  //col:247
    ThreadIdealProcessor // s: ULONG = 14  //col:248
    ThreadPriorityBoost // qs: ULONG = 15  //col:249
    ThreadSetTlsArrayAddress // s: ULONG_PTR  = 16  //col:250
    ThreadIsIoPending // q: ULONG = 17  //col:251
    ThreadHideFromDebugger // q: BOOLEAN; s: void = 18  //col:252
    ThreadBreakOnTermination // qs: ULONG = 19  //col:253
    ThreadSwitchLegacyState // s: void // NtCurrentThread // NPX/FPU = 20  //col:254
    ThreadIsTerminated // q: ULONG // 20 = 21  //col:255
    ThreadLastSystemCall // q: THREAD_LAST_SYSCALL_INFORMATION = 22  //col:256
    ThreadIoPriority // qs: IO_PRIORITY_HINT (requires SeIncreaseBasePriorityPrivilege) = 23  //col:257
    ThreadCycleTime // q: THREAD_CYCLE_TIME_INFORMATION = 24  //col:258
    ThreadPagePriority // q: ULONG = 25  //col:259
    ThreadActualBasePriority // s: LONG (requires SeIncreaseBasePriorityPrivilege) = 26  //col:260
    ThreadTebInformation // q: THREAD_TEB_INFORMATION (requires THREAD_GET_CONTEXT + THREAD_SET_CONTEXT) = 27  //col:261
    ThreadCSwitchMon = 28  //col:262
    ThreadCSwitchPmu = 29  //col:263
    ThreadWow64Context // qs: WOW64_CONTEXT = 30  //col:264
    ThreadGroupInformation // qs: GROUP_AFFINITY // 30 = 31  //col:265
    ThreadUmsInformation // q: THREAD_UMS_INFORMATION = 32  //col:266
    ThreadCounterProfiling // q: BOOLEAN; s: THREAD_PROFILING_INFORMATION? = 33  //col:267
    ThreadIdealProcessorEx // qs: PROCESSOR_NUMBER; s: previous PROCESSOR_NUMBER on return = 34  //col:268
    ThreadCpuAccountingInformation // q: BOOLEAN; s: HANDLE (NtOpenSession) // NtCurrentThread // since WIN8 = 35  //col:269
    ThreadSuspendCount // q: ULONG // since WINBLUE = 36  //col:270
    ThreadHeterogeneousCpuPolicy // q: KHETERO_CPU_POLICY // since THRESHOLD = 37  //col:271
    ThreadContainerId // q: GUID = 38  //col:272
    ThreadNameInformation // qs: THREAD_NAME_INFORMATION = 39  //col:273
    ThreadSelectedCpuSets = 40  //col:274
    ThreadSystemThreadInformation // q: SYSTEM_THREAD_INFORMATION // 40 = 41  //col:275
    ThreadActualGroupAffinity // q: GROUP_AFFINITY // since THRESHOLD2 = 42  //col:276
    ThreadDynamicCodePolicyInfo // q: ULONG; s: ULONG (NtCurrentThread) = 43  //col:277
    ThreadExplicitCaseSensitivity // qs: ULONG; s: 0 disables otherwise enables = 44  //col:278
    ThreadWorkOnBehalfTicket // RTL_WORK_ON_BEHALF_TICKET_EX = 45  //col:279
    ThreadSubsystemInformation // q: SUBSYSTEM_INFORMATION_TYPE // since REDSTONE2 = 46  //col:280
    ThreadDbgkWerReportActive // s: ULONG; s: 0 disables otherwise enables = 47  //col:281
    ThreadAttachContainer // s: HANDLE (job object) // NtCurrentThread = 48  //col:282
    ThreadManageWritesToExecutableMemory // MANAGE_WRITES_TO_EXECUTABLE_MEMORY // since REDSTONE3 = 49  //col:283
    ThreadPowerThrottlingState // POWER_THROTTLING_THREAD_STATE = 50  //col:284
    ThreadWorkloadClass // THREAD_WORKLOAD_CLASS // since REDSTONE5 // 50 = 51  //col:285
    ThreadCreateStateChange // since WIN11 = 52  //col:286
    ThreadApplyStateChange = 53  //col:287
    ThreadStrongerBadHandleChecks // since 22H1 = 54  //col:288
    ThreadEffectiveIoPriority = 55  //col:289
    ThreadEffectivePagePriority = 56  //col:290
    MaxThreadInfoClass = 57  //col:291
)


const(
    ProcessTlsReplaceIndex = 1  //col:557
    ProcessTlsReplaceVector = 2  //col:558
    MaxProcessTlsOperation = 3  //col:559
)


const(
    ProcessWorkingSetSwap = 1  //col:719
    ProcessWorkingSetEmpty = 2  //col:720
    ProcessWorkingSetOperationMax = 3  //col:721
)


const(
    PsProtectedTypeNone = 1  //col:733
    PsProtectedTypeProtectedLight = 2  //col:734
    PsProtectedTypeProtected = 3  //col:735
    PsProtectedTypeMax = 4  //col:736
)


const(
    PsProtectedSignerNone = 1  //col:741
    PsProtectedSignerAuthenticode = 2  //col:742
    PsProtectedSignerCodeGen = 3  //col:743
    PsProtectedSignerAntimalware = 4  //col:744
    PsProtectedSignerLsa = 5  //col:745
    PsProtectedSignerWindows = 6  //col:746
    PsProtectedSignerWinTcb = 7  //col:747
    PsProtectedSignerWinSystem = 8  //col:748
    PsProtectedSignerApp = 9  //col:749
    PsProtectedSignerMax = 10  //col:750
)


const(
    UmsInformationCommandInvalid = 1  //col:1116
    UmsInformationCommandAttach = 2  //col:1117
    UmsInformationCommandDetach = 3  //col:1118
    UmsInformationCommandQuery = 4  //col:1119
)


const(
    SubsystemInformationTypeWin32 = 1  //col:1182
    SubsystemInformationTypeWSL = 2  //col:1183
    MaxSubsystemInformationType = 3  //col:1184
)


const(
    ThreadWorkloadClassDefault = 1  //col:1191
    ThreadWorkloadClassGraphics = 2  //col:1192
    MaxThreadWorkloadClass = 3  //col:1193
)


const(
    ProcessStateChangeSuspend = 1  //col:1358
    ProcessStateChangeResume = 2  //col:1359
    ProcessStateChangeMax = 3  //col:1360
)


const(
    ThreadStateChangeSuspend = 1  //col:1392
    ThreadStateChangeResume = 2  //col:1393
    ThreadStateChangeMax = 3  //col:1394
)


const(
    QUEUE_USER_APC_FLAGS_NONE  =  0x0  //col:1624
    QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC  =  0x1  //col:1625
)


const(
    SeSafeOpenExperienceNone  =  0x00  //col:1769
    SeSafeOpenExperienceCalled  =  0x01  //col:1770
    SeSafeOpenExperienceAppRepCalled  =  0x02  //col:1771
    SeSafeOpenExperiencePromptDisplayed  =  0x04  //col:1772
    SeSafeOpenExperienceUAC  =  0x08  //col:1773
    SeSafeOpenExperienceUninstaller  =  0x10  //col:1774
    SeSafeOpenExperienceIgnoreUnknownOrBad  =  0x20  //col:1775
    SeSafeOpenExperienceDefenderTrustedInstaller  =  0x40  //col:1776
    SeSafeOpenExperienceMOTWPresent  =  0x80  //col:1777
)


const(
    PsAttributeParentProcess // in HANDLE = 1  //col:1826
    PsAttributeDebugObject // in HANDLE = 2  //col:1827
    PsAttributeToken // in HANDLE = 3  //col:1828
    PsAttributeClientId // out PCLIENT_ID = 4  //col:1829
    PsAttributeTebAddress // out PTEB * = 5  //col:1830
    PsAttributeImageName // in PWSTR = 6  //col:1831
    PsAttributeImageInfo // out PSECTION_IMAGE_INFORMATION = 7  //col:1832
    PsAttributeMemoryReserve // in PPS_MEMORY_RESERVE = 8  //col:1833
    PsAttributePriorityClass // in UCHAR = 9  //col:1834
    PsAttributeErrorMode // in ULONG = 10  //col:1835
    PsAttributeStdHandleInfo // 10 in PPS_STD_HANDLE_INFO = 11  //col:1836
    PsAttributeHandleList // in HANDLE[] = 12  //col:1837
    PsAttributeGroupAffinity // in PGROUP_AFFINITY = 13  //col:1838
    PsAttributePreferredNode // in PUSHORT = 14  //col:1839
    PsAttributeIdealProcessor // in PPROCESSOR_NUMBER = 15  //col:1840
    PsAttributeUmsThread // ? in PUMS_CREATE_THREAD_ATTRIBUTES = 16  //col:1841
    PsAttributeMitigationOptions // in PPS_MITIGATION_OPTIONS_MAP (PROCESS_CREATION_MITIGATION_POLICY_*) // since WIN8 = 17  //col:1842
    PsAttributeProtectionLevel // in PS_PROTECTION // since WINBLUE = 18  //col:1843
    PsAttributeSecureProcess // in PPS_TRUSTLET_CREATE_ATTRIBUTES since THRESHOLD = 19  //col:1844
    PsAttributeJobList // in HANDLE[] = 20  //col:1845
    PsAttributeChildProcessPolicy // 20 in PULONG (PROCESS_CREATION_CHILD_PROCESS_*) // since THRESHOLD2 = 21  //col:1846
    PsAttributeAllApplicationPackagesPolicy // in PULONG (PROCESS_CREATION_ALL_APPLICATION_PACKAGES_*) // since REDSTONE = 22  //col:1847
    PsAttributeWin32kFilter // in PWIN32K_SYSCALL_FILTER = 23  //col:1848
    PsAttributeSafeOpenPromptOriginClaim // in = 24  //col:1849
    PsAttributeBnoIsolation // in PPS_BNO_ISOLATION_PARAMETERS // since REDSTONE2 = 25  //col:1850
    PsAttributeDesktopAppPolicy // in PULONG (PROCESS_CREATION_DESKTOP_APP_*) = 26  //col:1851
    PsAttributeChpe // in BOOLEAN // since REDSTONE3 = 27  //col:1852
    PsAttributeMitigationAuditOptions // in PPS_MITIGATION_AUDIT_OPTIONS_MAP (PROCESS_CREATION_MITIGATION_AUDIT_POLICY_*) // since 21H1 = 28  //col:1853
    PsAttributeMachineType // in WORD // since 21H2 = 29  //col:1854
    PsAttributeComponentFilter = 30  //col:1855
    PsAttributeEnableOptionalXStateFeatures // since WIN11 = 31  //col:1856
    PsAttributeMax = 32  //col:1857
)


const(
    PsNeverDuplicate = 1  //col:1967
    PsRequestDuplicate // duplicate standard handles specified by PseudoHandleMask and only if StdHandleSubsystemType matches the image subsystem = 2  //col:1968
    PsAlwaysDuplicate // always duplicate standard handles = 3  //col:1969
    PsMaxStdHandleStates = 4  //col:1970
)


const(
    PS_MITIGATION_OPTION_NX = 1  //col:2049
    PS_MITIGATION_OPTION_SEHOP = 2  //col:2050
    PS_MITIGATION_OPTION_FORCE_RELOCATE_IMAGES = 3  //col:2051
    PS_MITIGATION_OPTION_HEAP_TERMINATE = 4  //col:2052
    PS_MITIGATION_OPTION_BOTTOM_UP_ASLR = 5  //col:2053
    PS_MITIGATION_OPTION_HIGH_ENTROPY_ASLR = 6  //col:2054
    PS_MITIGATION_OPTION_STRICT_HANDLE_CHECKS = 7  //col:2055
    PS_MITIGATION_OPTION_WIN32K_SYSTEM_CALL_DISABLE = 8  //col:2056
    PS_MITIGATION_OPTION_EXTENSION_POINT_DISABLE = 9  //col:2057
    PS_MITIGATION_OPTION_PROHIBIT_DYNAMIC_CODE = 10  //col:2058
    PS_MITIGATION_OPTION_CONTROL_FLOW_GUARD = 11  //col:2059
    PS_MITIGATION_OPTION_BLOCK_NON_MICROSOFT_BINARIES = 12  //col:2060
    PS_MITIGATION_OPTION_FONT_DISABLE = 13  //col:2061
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_REMOTE = 14  //col:2062
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_LOW_LABEL = 15  //col:2063
    PS_MITIGATION_OPTION_IMAGE_LOAD_PREFER_SYSTEM32 = 16  //col:2064
    PS_MITIGATION_OPTION_RETURN_FLOW_GUARD = 17  //col:2065
    PS_MITIGATION_OPTION_LOADER_INTEGRITY_CONTINUITY = 18  //col:2066
    PS_MITIGATION_OPTION_STRICT_CONTROL_FLOW_GUARD = 19  //col:2067
    PS_MITIGATION_OPTION_RESTRICT_SET_THREAD_CONTEXT = 20  //col:2068
    PS_MITIGATION_OPTION_ROP_STACKPIVOT // since REDSTONE3 = 21  //col:2069
    PS_MITIGATION_OPTION_ROP_CALLER_CHECK = 22  //col:2070
    PS_MITIGATION_OPTION_ROP_SIMEXEC = 23  //col:2071
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER = 24  //col:2072
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER_PLUS = 25  //col:2073
    PS_MITIGATION_OPTION_RESTRICT_CHILD_PROCESS_CREATION = 26  //col:2074
    PS_MITIGATION_OPTION_IMPORT_ADDRESS_FILTER = 27  //col:2075
    PS_MITIGATION_OPTION_MODULE_TAMPERING_PROTECTION = 28  //col:2076
    PS_MITIGATION_OPTION_RESTRICT_INDIRECT_BRANCH_PREDICTION = 29  //col:2077
    PS_MITIGATION_OPTION_SPECULATIVE_STORE_BYPASS_DISABLE // since REDSTONE5 = 30  //col:2078
    PS_MITIGATION_OPTION_ALLOW_DOWNGRADE_DYNAMIC_CODE_POLICY = 31  //col:2079
    PS_MITIGATION_OPTION_CET_USER_SHADOW_STACKS = 32  //col:2080
    PS_MITIGATION_OPTION_USER_CET_SET_CONTEXT_IP_VALIDATION // since 21H1 = 33  //col:2081
    PS_MITIGATION_OPTION_BLOCK_NON_CET_BINARIES = 34  //col:2082
    PS_MITIGATION_OPTION_CET_DYNAMIC_APIS_OUT_OF_PROC_ONLY = 35  //col:2083
    PS_MITIGATION_OPTION_REDIRECTION_TRUST // since 22H1 = 36  //col:2084
)


const(
    PsCreateInitialState = 1  //col:2090
    PsCreateFailOnFileOpen = 2  //col:2091
    PsCreateFailOnSectionCreate = 3  //col:2092
    PsCreateFailExeFormat = 4  //col:2093
    PsCreateFailMachineMismatch = 5  //col:2094
    PsCreateFailExeName // Debugger specified = 6  //col:2095
    PsCreateSuccess = 7  //col:2096
    PsCreateMaximumStates = 8  //col:2097
)


const(
    MemoryReserveUserApc = 1  //col:2477
    MemoryReserveIoCompletion = 2  //col:2478
    MemoryReserveTypeMax = 3  //col:2479
)



type (
Ntpsapi interface{
#if ()(ok bool)//col:92
#if ()(ok bool)//col:229
#if ()(ok bool)//col:292
#if ()(ok bool)//col:300
#if ()(ok bool)//col:315
#if ()(ok bool)//col:478
#if ()(ok bool)//col:677
    ()(ok bool)//col:783
#if ()(ok bool)//col:1185
#if ()(ok bool)//col:1361
#if ()(ok bool)//col:1395
#if ()(ok bool)//col:1626
NtQueueApcThreadEx2()(ok bool)//col:1740
    ProcThreadAttributeValue()(ok bool)//col:1858
    ()(ok bool)//col:1951
#if ()(ok bool)//col:2294
NtCreateJobObject()(ok bool)//col:2480
}

















































)

func NewNtpsapi() { return & ntpsapi{} }

func (n *ntpsapi)#if ()(ok bool){//col:92




































































return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:229


















return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:292












return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:300





return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:315










return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:478
















return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:677













return true
}


















































func (n *ntpsapi)    ()(ok bool){//col:783






















return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:1185







return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:1361





















































































































return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:1395





























return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:1626




































































































































































































return true
}


















































func (n *ntpsapi)NtQueueApcThreadEx2()(ok bool){//col:1740































































return true
}


















































func (n *ntpsapi)    ProcThreadAttributeValue()(ok bool){//col:1858























return true
}


















































func (n *ntpsapi)    ()(ok bool){//col:1951













































































return true
}


















































func (n *ntpsapi)#if ()(ok bool){//col:2294

































































return true
}


















































func (n *ntpsapi)NtCreateJobObject()(ok bool){//col:2480













































































return true
}




















































