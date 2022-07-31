package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpsapi.h.back

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
// = JOB_OBJECT_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x1f) // pre-Vista full access //col:53
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

type     ProcessBasicInformation // q: PROCESS_BASIC_INFORMATION PROCESS_EXTENDED_BASIC_INFORMATION uint32
const(
    ProcessBasicInformation // q: PROCESS_BASIC_INFORMATION PROCESS_EXTENDED_BASIC_INFORMATION PROCESSINFOCLASS = 1  //col:116
    ProcessQuotaLimits // qs: QUOTA_LIMITS QUOTA_LIMITS_EX PROCESSINFOCLASS = 2  //col:117
    ProcessIoCounters // q: IO_COUNTERS PROCESSINFOCLASS = 3  //col:118
    ProcessVmCounters // q: VM_COUNTERS VM_COUNTERS_EX VM_COUNTERS_EX2 PROCESSINFOCLASS = 4  //col:119
    ProcessTimes // q: KERNEL_USER_TIMES PROCESSINFOCLASS = 5  //col:120
    ProcessBasePriority // s: KPRIORITY PROCESSINFOCLASS = 6  //col:121
    ProcessRaisePriority // s: ULONG PROCESSINFOCLASS = 7  //col:122
    ProcessDebugPort // q: HANDLE PROCESSINFOCLASS = 8  //col:123
    ProcessExceptionPort // s: PROCESS_EXCEPTION_PORT PROCESSINFOCLASS = 9  //col:124
    ProcessAccessToken // s: PROCESS_ACCESS_TOKEN PROCESSINFOCLASS = 10  //col:125
    ProcessLdtInformation // qs: PROCESS_LDT_INFORMATION // 10 PROCESSINFOCLASS = 11  //col:126
    ProcessLdtSize // s: PROCESS_LDT_SIZE PROCESSINFOCLASS = 12  //col:127
    ProcessDefaultHardErrorMode // qs: ULONG PROCESSINFOCLASS = 13  //col:128
    ProcessIoPortHandlers // (kernel-mode only) // PROCESS_IO_PORT_HANDLER_INFORMATION PROCESSINFOCLASS = 14  //col:129
    ProcessPooledUsageAndLimits // q: POOLED_USAGE_AND_LIMITS PROCESSINFOCLASS = 15  //col:130
    ProcessWorkingSetWatch // q: PROCESS_WS_WATCH_INFORMATION[]; s: void PROCESSINFOCLASS = 16  //col:131
    ProcessUserModeIOPL // qs: ULONG (requires SeTcbPrivilege) PROCESSINFOCLASS = 17  //col:132
    ProcessEnableAlignmentFaultFixup // s: BOOLEAN PROCESSINFOCLASS = 18  //col:133
    ProcessPriorityClass // qs: PROCESS_PRIORITY_CLASS PROCESSINFOCLASS = 19  //col:134
    ProcessWx86Information // qs: ULONG (requires SeTcbPrivilege) (VdmAllowed) PROCESSINFOCLASS = 20  //col:135
    ProcessHandleCount // q: ULONG PROCESS_HANDLE_INFORMATION // 20 PROCESSINFOCLASS = 21  //col:136
    ProcessAffinityMask // qs: KAFFINITY qs: GROUP_AFFINITY PROCESSINFOCLASS = 22  //col:137
    ProcessPriorityBoost // qs: ULONG PROCESSINFOCLASS = 23  //col:138
    ProcessDeviceMap // qs: PROCESS_DEVICEMAP_INFORMATION PROCESS_DEVICEMAP_INFORMATION_EX PROCESSINFOCLASS = 24  //col:139
    ProcessSessionInformation // q: PROCESS_SESSION_INFORMATION PROCESSINFOCLASS = 25  //col:140
    ProcessForegroundInformation // s: PROCESS_FOREGROUND_BACKGROUND PROCESSINFOCLASS = 26  //col:141
    ProcessWow64Information // q: ULONG_PTR PROCESSINFOCLASS = 27  //col:142
    ProcessImageFileName // q: UNICODE_STRING PROCESSINFOCLASS = 28  //col:143
    ProcessLUIDDeviceMapsEnabled // q: ULONG PROCESSINFOCLASS = 29  //col:144
    ProcessBreakOnTermination // qs: ULONG PROCESSINFOCLASS = 30  //col:145
    ProcessDebugObjectHandle // q: HANDLE // 30 PROCESSINFOCLASS = 31  //col:146
    ProcessDebugFlags // qs: ULONG PROCESSINFOCLASS = 32  //col:147
    ProcessHandleTracing // q: PROCESS_HANDLE_TRACING_QUERY; s: size 0 disables otherwise enables PROCESSINFOCLASS = 33  //col:148
    ProcessIoPriority // qs: IO_PRIORITY_HINT PROCESSINFOCLASS = 34  //col:149
    ProcessExecuteFlags // qs: ULONG PROCESSINFOCLASS = 35  //col:150
    ProcessTlsInformation // PROCESS_TLS_INFORMATION // ProcessResourceManagement  PROCESSINFOCLASS = 36  //col:151
    ProcessCookie // q: ULONG PROCESSINFOCLASS = 37  //col:152
    ProcessImageInformation // q: SECTION_IMAGE_INFORMATION PROCESSINFOCLASS = 38  //col:153
    ProcessCycleTime // q: PROCESS_CYCLE_TIME_INFORMATION // since VISTA PROCESSINFOCLASS = 39  //col:154
    ProcessPagePriority // q: PAGE_PRIORITY_INFORMATION PROCESSINFOCLASS = 40  //col:155
    ProcessInstrumentationCallback // s: PVOID or PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION // 40 PROCESSINFOCLASS = 41  //col:156
    ProcessThreadStackAllocation // s: PROCESS_STACK_ALLOCATION_INFORMATION PROCESS_STACK_ALLOCATION_INFORMATION_EX PROCESSINFOCLASS = 42  //col:157
    ProcessWorkingSetWatchEx // q: PROCESS_WS_WATCH_INFORMATION_EX[] PROCESSINFOCLASS = 43  //col:158
    ProcessImageFileNameWin32 // q: UNICODE_STRING PROCESSINFOCLASS = 44  //col:159
    ProcessImageFileMapping // q: HANDLE (input) PROCESSINFOCLASS = 45  //col:160
    ProcessAffinityUpdateMode // qs: PROCESS_AFFINITY_UPDATE_MODE PROCESSINFOCLASS = 46  //col:161
    ProcessMemoryAllocationMode // qs: PROCESS_MEMORY_ALLOCATION_MODE PROCESSINFOCLASS = 47  //col:162
    ProcessGroupInformation // q: USHORT[] PROCESSINFOCLASS = 48  //col:163
    ProcessTokenVirtualizationEnabled // s: ULONG PROCESSINFOCLASS = 49  //col:164
    ProcessConsoleHostProcess // q: ULONG_PTR // ProcessOwnerInformation PROCESSINFOCLASS = 50  //col:165
    ProcessWindowInformation // q: PROCESS_WINDOW_INFORMATION // 50 PROCESSINFOCLASS = 51  //col:166
    ProcessHandleInformation // q: PROCESS_HANDLE_SNAPSHOT_INFORMATION // since WIN8 PROCESSINFOCLASS = 52  //col:167
    ProcessMitigationPolicy // s: PROCESS_MITIGATION_POLICY_INFORMATION PROCESSINFOCLASS = 53  //col:168
    ProcessDynamicFunctionTableInformation PROCESSINFOCLASS = 54  //col:169
    ProcessHandleCheckingMode // qs: ULONG; s: 0 disables otherwise enables PROCESSINFOCLASS = 55  //col:170
    ProcessKeepAliveCount // q: PROCESS_KEEPALIVE_COUNT_INFORMATION PROCESSINFOCLASS = 56  //col:171
    ProcessRevokeFileHandles // s: PROCESS_REVOKE_FILE_HANDLES_INFORMATION PROCESSINFOCLASS = 57  //col:172
    ProcessWorkingSetControl // s: PROCESS_WORKING_SET_CONTROL PROCESSINFOCLASS = 58  //col:173
    ProcessHandleTable // q: ULONG[] // since WINBLUE PROCESSINFOCLASS = 59  //col:174
    ProcessCheckStackExtentsMode // qs: ULONG // KPROCESS->CheckStackExtents (CFG) PROCESSINFOCLASS = 60  //col:175
    ProcessCommandLineInformation // q: UNICODE_STRING // 60 PROCESSINFOCLASS = 61  //col:176
    ProcessProtectionInformation // q: PS_PROTECTION PROCESSINFOCLASS = 62  //col:177
    ProcessMemoryExhaustion // PROCESS_MEMORY_EXHAUSTION_INFO // since THRESHOLD PROCESSINFOCLASS = 63  //col:178
    ProcessFaultInformation // PROCESS_FAULT_INFORMATION PROCESSINFOCLASS = 64  //col:179
    ProcessTelemetryIdInformation // q: PROCESS_TELEMETRY_ID_INFORMATION PROCESSINFOCLASS = 65  //col:180
    ProcessCommitReleaseInformation // PROCESS_COMMIT_RELEASE_INFORMATION PROCESSINFOCLASS = 66  //col:181
    ProcessDefaultCpuSetsInformation // SYSTEM_CPU_SET_INFORMATION[5] PROCESSINFOCLASS = 67  //col:182
    ProcessAllowedCpuSetsInformation // SYSTEM_CPU_SET_INFORMATION[5] PROCESSINFOCLASS = 68  //col:183
    ProcessSubsystemProcess PROCESSINFOCLASS = 69  //col:184
    ProcessJobMemoryInformation // q: PROCESS_JOB_MEMORY_INFO PROCESSINFOCLASS = 70  //col:185
    ProcessInPrivate // s: void // ETW // since THRESHOLD2 // 70 PROCESSINFOCLASS = 71  //col:186
    ProcessRaiseUMExceptionOnInvalidHandleClose // qs: ULONG; s: 0 disables otherwise enables PROCESSINFOCLASS = 72  //col:187
    ProcessIumChallengeResponse PROCESSINFOCLASS = 73  //col:188
    ProcessChildProcessInformation // q: PROCESS_CHILD_PROCESS_INFORMATION PROCESSINFOCLASS = 74  //col:189
    ProcessHighGraphicsPriorityInformation // qs: BOOLEAN (requires SeTcbPrivilege) PROCESSINFOCLASS = 75  //col:190
    ProcessSubsystemInformation // q: SUBSYSTEM_INFORMATION_TYPE // since REDSTONE2 PROCESSINFOCLASS = 76  //col:191
    ProcessEnergyValues // q: PROCESS_ENERGY_VALUES PROCESS_EXTENDED_ENERGY_VALUES PROCESSINFOCLASS = 77  //col:192
    ProcessPowerThrottlingState // qs: POWER_THROTTLING_PROCESS_STATE PROCESSINFOCLASS = 78  //col:193
    ProcessReserved3Information // ProcessActivityThrottlePolicy // PROCESS_ACTIVITY_THROTTLE_POLICY PROCESSINFOCLASS = 79  //col:194
    ProcessWin32kSyscallFilterInformation // q: WIN32K_SYSCALL_FILTER PROCESSINFOCLASS = 80  //col:195
    ProcessDisableSystemAllowedCpuSets // 80 PROCESSINFOCLASS = 81  //col:196
    ProcessWakeInformation // PROCESS_WAKE_INFORMATION PROCESSINFOCLASS = 82  //col:197
    ProcessEnergyTrackingState // PROCESS_ENERGY_TRACKING_STATE PROCESSINFOCLASS = 83  //col:198
    ProcessManageWritesToExecutableMemory // MANAGE_WRITES_TO_EXECUTABLE_MEMORY // since REDSTONE3 PROCESSINFOCLASS = 84  //col:199
    ProcessCaptureTrustletLiveDump PROCESSINFOCLASS = 85  //col:200
    ProcessTelemetryCoverage PROCESSINFOCLASS = 86  //col:201
    ProcessEnclaveInformation PROCESSINFOCLASS = 87  //col:202
    ProcessEnableReadWriteVmLogging // PROCESS_READWRITEVM_LOGGING_INFORMATION PROCESSINFOCLASS = 88  //col:203
    ProcessUptimeInformation // q: PROCESS_UPTIME_INFORMATION PROCESSINFOCLASS = 89  //col:204
    ProcessImageSection // q: HANDLE PROCESSINFOCLASS = 90  //col:205
    ProcessDebugAuthInformation // since REDSTONE4 // 90 PROCESSINFOCLASS = 91  //col:206
    ProcessSystemResourceManagement // PROCESS_SYSTEM_RESOURCE_MANAGEMENT PROCESSINFOCLASS = 92  //col:207
    ProcessSequenceNumber // q: ULONGLONG PROCESSINFOCLASS = 93  //col:208
    ProcessLoaderDetour // since REDSTONE5 PROCESSINFOCLASS = 94  //col:209
    ProcessSecurityDomainInformation // PROCESS_SECURITY_DOMAIN_INFORMATION PROCESSINFOCLASS = 95  //col:210
    ProcessCombineSecurityDomainsInformation // PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION PROCESSINFOCLASS = 96  //col:211
    ProcessEnableLogging // PROCESS_LOGGING_INFORMATION PROCESSINFOCLASS = 97  //col:212
    ProcessLeapSecondInformation // PROCESS_LEAP_SECOND_INFORMATION PROCESSINFOCLASS = 98  //col:213
    ProcessFiberShadowStackAllocation // PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION // since 19H1 PROCESSINFOCLASS = 99  //col:214
    ProcessFreeFiberShadowStackAllocation // PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION PROCESSINFOCLASS = 100  //col:215
    ProcessAltSystemCallInformation // qs: BOOLEAN (kernel-mode only) // INT2E // since 20H1 // 100 PROCESSINFOCLASS = 101  //col:216
    ProcessDynamicEHContinuationTargets // PROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION PROCESSINFOCLASS = 102  //col:217
    ProcessDynamicEnforcedCetCompatibleRanges // PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE_INFORMATION // since 20H2 PROCESSINFOCLASS = 103  //col:218
    ProcessCreateStateChange // since WIN11 PROCESSINFOCLASS = 104  //col:219
    ProcessApplyStateChange PROCESSINFOCLASS = 105  //col:220
    ProcessEnableOptionalXStateFeatures PROCESSINFOCLASS = 106  //col:221
    ProcessAltPrefetchParam // since 22H1 PROCESSINFOCLASS = 107  //col:222
    ProcessAssignCpuPartitions PROCESSINFOCLASS = 108  //col:223
    ProcessPriorityClassEx PROCESSINFOCLASS = 109  //col:224
    ProcessMembershipInformation PROCESSINFOCLASS = 110  //col:225
    ProcessEffectiveIoPriority PROCESSINFOCLASS = 111  //col:226
    ProcessEffectivePagePriority PROCESSINFOCLASS = 112  //col:227
    MaxProcessInfoClass PROCESSINFOCLASS = 113  //col:228
)


type     ThreadBasicInformation // q: THREAD_BASIC_INFORMATION uint32
const(
    ThreadBasicInformation // q: THREAD_BASIC_INFORMATION THREADINFOCLASS = 1  //col:235
    ThreadTimes // q: KERNEL_USER_TIMES THREADINFOCLASS = 2  //col:236
    ThreadPriority // s: KPRIORITY (requires SeIncreaseBasePriorityPrivilege) THREADINFOCLASS = 3  //col:237
    ThreadBasePriority // s: KPRIORITY THREADINFOCLASS = 4  //col:238
    ThreadAffinityMask // s: KAFFINITY THREADINFOCLASS = 5  //col:239
    ThreadImpersonationToken // s: HANDLE THREADINFOCLASS = 6  //col:240
    ThreadDescriptorTableEntry // q: DESCRIPTOR_TABLE_ENTRY (or WOW64_DESCRIPTOR_TABLE_ENTRY) THREADINFOCLASS = 7  //col:241
    ThreadEnableAlignmentFaultFixup // s: BOOLEAN THREADINFOCLASS = 8  //col:242
    ThreadEventPair THREADINFOCLASS = 9  //col:243
    ThreadQuerySetWin32StartAddress // q: ULONG_PTR THREADINFOCLASS = 10  //col:244
    ThreadZeroTlsCell // s: ULONG // TlsIndex // 10 THREADINFOCLASS = 11  //col:245
    ThreadPerformanceCount // q: LARGE_INTEGER THREADINFOCLASS = 12  //col:246
    ThreadAmILastThread // q: ULONG THREADINFOCLASS = 13  //col:247
    ThreadIdealProcessor // s: ULONG THREADINFOCLASS = 14  //col:248
    ThreadPriorityBoost // qs: ULONG THREADINFOCLASS = 15  //col:249
    ThreadSetTlsArrayAddress // s: ULONG_PTR  THREADINFOCLASS = 16  //col:250
    ThreadIsIoPending // q: ULONG THREADINFOCLASS = 17  //col:251
    ThreadHideFromDebugger // q: BOOLEAN; s: void THREADINFOCLASS = 18  //col:252
    ThreadBreakOnTermination // qs: ULONG THREADINFOCLASS = 19  //col:253
    ThreadSwitchLegacyState // s: void // NtCurrentThread // NPX/FPU THREADINFOCLASS = 20  //col:254
    ThreadIsTerminated // q: ULONG // 20 THREADINFOCLASS = 21  //col:255
    ThreadLastSystemCall // q: THREAD_LAST_SYSCALL_INFORMATION THREADINFOCLASS = 22  //col:256
    ThreadIoPriority // qs: IO_PRIORITY_HINT (requires SeIncreaseBasePriorityPrivilege) THREADINFOCLASS = 23  //col:257
    ThreadCycleTime // q: THREAD_CYCLE_TIME_INFORMATION THREADINFOCLASS = 24  //col:258
    ThreadPagePriority // q: ULONG THREADINFOCLASS = 25  //col:259
    ThreadActualBasePriority // s: LONG (requires SeIncreaseBasePriorityPrivilege) THREADINFOCLASS = 26  //col:260
    ThreadTebInformation // q: THREAD_TEB_INFORMATION (requires THREAD_GET_CONTEXT + THREAD_SET_CONTEXT) THREADINFOCLASS = 27  //col:261
    ThreadCSwitchMon THREADINFOCLASS = 28  //col:262
    ThreadCSwitchPmu THREADINFOCLASS = 29  //col:263
    ThreadWow64Context // qs: WOW64_CONTEXT THREADINFOCLASS = 30  //col:264
    ThreadGroupInformation // qs: GROUP_AFFINITY // 30 THREADINFOCLASS = 31  //col:265
    ThreadUmsInformation // q: THREAD_UMS_INFORMATION THREADINFOCLASS = 32  //col:266
    ThreadCounterProfiling // q: BOOLEAN; s: THREAD_PROFILING_INFORMATION? THREADINFOCLASS = 33  //col:267
    ThreadIdealProcessorEx // qs: PROCESSOR_NUMBER; s: previous PROCESSOR_NUMBER on return THREADINFOCLASS = 34  //col:268
    ThreadCpuAccountingInformation // q: BOOLEAN; s: HANDLE (NtOpenSession) // NtCurrentThread // since WIN8 THREADINFOCLASS = 35  //col:269
    ThreadSuspendCount // q: ULONG // since WINBLUE THREADINFOCLASS = 36  //col:270
    ThreadHeterogeneousCpuPolicy // q: KHETERO_CPU_POLICY // since THRESHOLD THREADINFOCLASS = 37  //col:271
    ThreadContainerId // q: GUID THREADINFOCLASS = 38  //col:272
    ThreadNameInformation // qs: THREAD_NAME_INFORMATION THREADINFOCLASS = 39  //col:273
    ThreadSelectedCpuSets THREADINFOCLASS = 40  //col:274
    ThreadSystemThreadInformation // q: SYSTEM_THREAD_INFORMATION // 40 THREADINFOCLASS = 41  //col:275
    ThreadActualGroupAffinity // q: GROUP_AFFINITY // since THRESHOLD2 THREADINFOCLASS = 42  //col:276
    ThreadDynamicCodePolicyInfo // q: ULONG; s: ULONG (NtCurrentThread) THREADINFOCLASS = 43  //col:277
    ThreadExplicitCaseSensitivity // qs: ULONG; s: 0 disables otherwise enables THREADINFOCLASS = 44  //col:278
    ThreadWorkOnBehalfTicket // RTL_WORK_ON_BEHALF_TICKET_EX THREADINFOCLASS = 45  //col:279
    ThreadSubsystemInformation // q: SUBSYSTEM_INFORMATION_TYPE // since REDSTONE2 THREADINFOCLASS = 46  //col:280
    ThreadDbgkWerReportActive // s: ULONG; s: 0 disables otherwise enables THREADINFOCLASS = 47  //col:281
    ThreadAttachContainer // s: HANDLE (job object) // NtCurrentThread THREADINFOCLASS = 48  //col:282
    ThreadManageWritesToExecutableMemory // MANAGE_WRITES_TO_EXECUTABLE_MEMORY // since REDSTONE3 THREADINFOCLASS = 49  //col:283
    ThreadPowerThrottlingState // POWER_THROTTLING_THREAD_STATE THREADINFOCLASS = 50  //col:284
    ThreadWorkloadClass // THREAD_WORKLOAD_CLASS // since REDSTONE5 // 50 THREADINFOCLASS = 51  //col:285
    ThreadCreateStateChange // since WIN11 THREADINFOCLASS = 52  //col:286
    ThreadApplyStateChange THREADINFOCLASS = 53  //col:287
    ThreadStrongerBadHandleChecks // since 22H1 THREADINFOCLASS = 54  //col:288
    ThreadEffectiveIoPriority THREADINFOCLASS = 55  //col:289
    ThreadEffectivePagePriority THREADINFOCLASS = 56  //col:290
    MaxThreadInfoClass THREADINFOCLASS = 57  //col:291
)


type     ProcessTlsReplaceIndex uint32
const(
    ProcessTlsReplaceIndex PROCESS_TLS_INFORMATION_TYPE = 1  //col:557
    ProcessTlsReplaceVector PROCESS_TLS_INFORMATION_TYPE = 2  //col:558
    MaxProcessTlsOperation PROCESS_TLS_INFORMATION_TYPE = 3  //col:559
)


type     ProcessWorkingSetSwap uint32
const(
    ProcessWorkingSetSwap PROCESS_WORKING_SET_OPERATION = 1  //col:719
    ProcessWorkingSetEmpty PROCESS_WORKING_SET_OPERATION = 2  //col:720
    ProcessWorkingSetOperationMax PROCESS_WORKING_SET_OPERATION = 3  //col:721
)


type     PsProtectedTypeNone uint32
const(
    PsProtectedTypeNone PS_PROTECTED_TYPE = 1  //col:733
    PsProtectedTypeProtectedLight PS_PROTECTED_TYPE = 2  //col:734
    PsProtectedTypeProtected PS_PROTECTED_TYPE = 3  //col:735
    PsProtectedTypeMax PS_PROTECTED_TYPE = 4  //col:736
)


type     PsProtectedSignerNone uint32
const(
    PsProtectedSignerNone PS_PROTECTED_SIGNER = 1  //col:741
    PsProtectedSignerAuthenticode PS_PROTECTED_SIGNER = 2  //col:742
    PsProtectedSignerCodeGen PS_PROTECTED_SIGNER = 3  //col:743
    PsProtectedSignerAntimalware PS_PROTECTED_SIGNER = 4  //col:744
    PsProtectedSignerLsa PS_PROTECTED_SIGNER = 5  //col:745
    PsProtectedSignerWindows PS_PROTECTED_SIGNER = 6  //col:746
    PsProtectedSignerWinTcb PS_PROTECTED_SIGNER = 7  //col:747
    PsProtectedSignerWinSystem PS_PROTECTED_SIGNER = 8  //col:748
    PsProtectedSignerApp PS_PROTECTED_SIGNER = 9  //col:749
    PsProtectedSignerMax PS_PROTECTED_SIGNER = 10  //col:750
)


type     UmsInformationCommandInvalid uint32
const(
    UmsInformationCommandInvalid THREAD_UMS_INFORMATION_COMMAND = 1  //col:1116
    UmsInformationCommandAttach THREAD_UMS_INFORMATION_COMMAND = 2  //col:1117
    UmsInformationCommandDetach THREAD_UMS_INFORMATION_COMMAND = 3  //col:1118
    UmsInformationCommandQuery THREAD_UMS_INFORMATION_COMMAND = 4  //col:1119
)


type     SubsystemInformationTypeWin32 uint32
const(
    SubsystemInformationTypeWin32 SUBSYSTEM_INFORMATION_TYPE  = 1  //col:1182
    SubsystemInformationTypeWSL SUBSYSTEM_INFORMATION_TYPE  = 2  //col:1183
    MaxSubsystemInformationType SUBSYSTEM_INFORMATION_TYPE  = 3  //col:1184
)


type     ThreadWorkloadClassDefault uint32
const(
    ThreadWorkloadClassDefault THREAD_WORKLOAD_CLASS = 1  //col:1191
    ThreadWorkloadClassGraphics THREAD_WORKLOAD_CLASS = 2  //col:1192
    MaxThreadWorkloadClass THREAD_WORKLOAD_CLASS = 3  //col:1193
)


type     ProcessStateChangeSuspend uint32
const(
    ProcessStateChangeSuspend PROCESS_STATE_CHANGE_TYPE = 1  //col:1358
    ProcessStateChangeResume PROCESS_STATE_CHANGE_TYPE = 2  //col:1359
    ProcessStateChangeMax PROCESS_STATE_CHANGE_TYPE = 3  //col:1360
)


type     ThreadStateChangeSuspend uint32
const(
    ThreadStateChangeSuspend THREAD_STATE_CHANGE_TYPE = 1  //col:1392
    ThreadStateChangeResume THREAD_STATE_CHANGE_TYPE = 2  //col:1393
    ThreadStateChangeMax THREAD_STATE_CHANGE_TYPE = 3  //col:1394
)


type     QUEUE_USER_APC_FLAGS_NONE = 0x0 uint32
const(
    QUEUE_USER_APC_FLAGS_NONE  QUEUE_USER_APC_FLAGS =  0x0  //col:1624
    QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC  QUEUE_USER_APC_FLAGS =  0x1  //col:1625
)


type     SeSafeOpenExperienceNone = 0x00 uint32
const(
    SeSafeOpenExperienceNone  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x00  //col:1769
    SeSafeOpenExperienceCalled  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x01  //col:1770
    SeSafeOpenExperienceAppRepCalled  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x02  //col:1771
    SeSafeOpenExperiencePromptDisplayed  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x04  //col:1772
    SeSafeOpenExperienceUAC  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x08  //col:1773
    SeSafeOpenExperienceUninstaller  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x10  //col:1774
    SeSafeOpenExperienceIgnoreUnknownOrBad  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x20  //col:1775
    SeSafeOpenExperienceDefenderTrustedInstaller  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x40  //col:1776
    SeSafeOpenExperienceMOTWPresent  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS { =  0x80  //col:1777
)


type     PsAttributeParentProcess // in HANDLE uint32
const(
    PsAttributeParentProcess // in HANDLE PS_ATTRIBUTE_NUM = 1  //col:1826
    PsAttributeDebugObject // in HANDLE PS_ATTRIBUTE_NUM = 2  //col:1827
    PsAttributeToken // in HANDLE PS_ATTRIBUTE_NUM = 3  //col:1828
    PsAttributeClientId // out PCLIENT_ID PS_ATTRIBUTE_NUM = 4  //col:1829
    PsAttributeTebAddress // out PTEB * PS_ATTRIBUTE_NUM = 5  //col:1830
    PsAttributeImageName // in PWSTR PS_ATTRIBUTE_NUM = 6  //col:1831
    PsAttributeImageInfo // out PSECTION_IMAGE_INFORMATION PS_ATTRIBUTE_NUM = 7  //col:1832
    PsAttributeMemoryReserve // in PPS_MEMORY_RESERVE PS_ATTRIBUTE_NUM = 8  //col:1833
    PsAttributePriorityClass // in UCHAR PS_ATTRIBUTE_NUM = 9  //col:1834
    PsAttributeErrorMode // in ULONG PS_ATTRIBUTE_NUM = 10  //col:1835
    PsAttributeStdHandleInfo // 10 in PPS_STD_HANDLE_INFO PS_ATTRIBUTE_NUM = 11  //col:1836
    PsAttributeHandleList // in HANDLE[] PS_ATTRIBUTE_NUM = 12  //col:1837
    PsAttributeGroupAffinity // in PGROUP_AFFINITY PS_ATTRIBUTE_NUM = 13  //col:1838
    PsAttributePreferredNode // in PUSHORT PS_ATTRIBUTE_NUM = 14  //col:1839
    PsAttributeIdealProcessor // in PPROCESSOR_NUMBER PS_ATTRIBUTE_NUM = 15  //col:1840
    PsAttributeUmsThread // ? in PUMS_CREATE_THREAD_ATTRIBUTES PS_ATTRIBUTE_NUM = 16  //col:1841
    PsAttributeMitigationOptions // in PPS_MITIGATION_OPTIONS_MAP (PROCESS_CREATION_MITIGATION_POLICY_*) // since WIN8 PS_ATTRIBUTE_NUM = 17  //col:1842
    PsAttributeProtectionLevel // in PS_PROTECTION // since WINBLUE PS_ATTRIBUTE_NUM = 18  //col:1843
    PsAttributeSecureProcess // in PPS_TRUSTLET_CREATE_ATTRIBUTES since THRESHOLD PS_ATTRIBUTE_NUM = 19  //col:1844
    PsAttributeJobList // in HANDLE[] PS_ATTRIBUTE_NUM = 20  //col:1845
    PsAttributeChildProcessPolicy // 20 in PULONG (PROCESS_CREATION_CHILD_PROCESS_*) // since THRESHOLD2 PS_ATTRIBUTE_NUM = 21  //col:1846
    PsAttributeAllApplicationPackagesPolicy // in PULONG (PROCESS_CREATION_ALL_APPLICATION_PACKAGES_*) // since REDSTONE PS_ATTRIBUTE_NUM = 22  //col:1847
    PsAttributeWin32kFilter // in PWIN32K_SYSCALL_FILTER PS_ATTRIBUTE_NUM = 23  //col:1848
    PsAttributeSafeOpenPromptOriginClaim // in PS_ATTRIBUTE_NUM = 24  //col:1849
    PsAttributeBnoIsolation // in PPS_BNO_ISOLATION_PARAMETERS // since REDSTONE2 PS_ATTRIBUTE_NUM = 25  //col:1850
    PsAttributeDesktopAppPolicy // in PULONG (PROCESS_CREATION_DESKTOP_APP_*) PS_ATTRIBUTE_NUM = 26  //col:1851
    PsAttributeChpe // in BOOLEAN // since REDSTONE3 PS_ATTRIBUTE_NUM = 27  //col:1852
    PsAttributeMitigationAuditOptions // in PPS_MITIGATION_AUDIT_OPTIONS_MAP (PROCESS_CREATION_MITIGATION_AUDIT_POLICY_*) // since 21H1 PS_ATTRIBUTE_NUM = 28  //col:1853
    PsAttributeMachineType // in WORD // since 21H2 PS_ATTRIBUTE_NUM = 29  //col:1854
    PsAttributeComponentFilter PS_ATTRIBUTE_NUM = 30  //col:1855
    PsAttributeEnableOptionalXStateFeatures // since WIN11 PS_ATTRIBUTE_NUM = 31  //col:1856
    PsAttributeMax PS_ATTRIBUTE_NUM = 32  //col:1857
)


type     PsNeverDuplicate uint32
const(
    PsNeverDuplicate PS_STD_HANDLE_STATE = 1  //col:1967
    PsRequestDuplicate // duplicate standard handles specified by PseudoHandleMask and only if StdHandleSubsystemType matches the image subsystem PS_STD_HANDLE_STATE = 2  //col:1968
    PsAlwaysDuplicate // always duplicate standard handles PS_STD_HANDLE_STATE = 3  //col:1969
    PsMaxStdHandleStates PS_STD_HANDLE_STATE = 4  //col:1970
)


type     PS_MITIGATION_OPTION_NX uint32
const(
    PS_MITIGATION_OPTION_NX PS_MITIGATION_OPTION = 1  //col:2049
    PS_MITIGATION_OPTION_SEHOP PS_MITIGATION_OPTION = 2  //col:2050
    PS_MITIGATION_OPTION_FORCE_RELOCATE_IMAGES PS_MITIGATION_OPTION = 3  //col:2051
    PS_MITIGATION_OPTION_HEAP_TERMINATE PS_MITIGATION_OPTION = 4  //col:2052
    PS_MITIGATION_OPTION_BOTTOM_UP_ASLR PS_MITIGATION_OPTION = 5  //col:2053
    PS_MITIGATION_OPTION_HIGH_ENTROPY_ASLR PS_MITIGATION_OPTION = 6  //col:2054
    PS_MITIGATION_OPTION_STRICT_HANDLE_CHECKS PS_MITIGATION_OPTION = 7  //col:2055
    PS_MITIGATION_OPTION_WIN32K_SYSTEM_CALL_DISABLE PS_MITIGATION_OPTION = 8  //col:2056
    PS_MITIGATION_OPTION_EXTENSION_POINT_DISABLE PS_MITIGATION_OPTION = 9  //col:2057
    PS_MITIGATION_OPTION_PROHIBIT_DYNAMIC_CODE PS_MITIGATION_OPTION = 10  //col:2058
    PS_MITIGATION_OPTION_CONTROL_FLOW_GUARD PS_MITIGATION_OPTION = 11  //col:2059
    PS_MITIGATION_OPTION_BLOCK_NON_MICROSOFT_BINARIES PS_MITIGATION_OPTION = 12  //col:2060
    PS_MITIGATION_OPTION_FONT_DISABLE PS_MITIGATION_OPTION = 13  //col:2061
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_REMOTE PS_MITIGATION_OPTION = 14  //col:2062
    PS_MITIGATION_OPTION_IMAGE_LOAD_NO_LOW_LABEL PS_MITIGATION_OPTION = 15  //col:2063
    PS_MITIGATION_OPTION_IMAGE_LOAD_PREFER_SYSTEM32 PS_MITIGATION_OPTION = 16  //col:2064
    PS_MITIGATION_OPTION_RETURN_FLOW_GUARD PS_MITIGATION_OPTION = 17  //col:2065
    PS_MITIGATION_OPTION_LOADER_INTEGRITY_CONTINUITY PS_MITIGATION_OPTION = 18  //col:2066
    PS_MITIGATION_OPTION_STRICT_CONTROL_FLOW_GUARD PS_MITIGATION_OPTION = 19  //col:2067
    PS_MITIGATION_OPTION_RESTRICT_SET_THREAD_CONTEXT PS_MITIGATION_OPTION = 20  //col:2068
    PS_MITIGATION_OPTION_ROP_STACKPIVOT // since REDSTONE3 PS_MITIGATION_OPTION = 21  //col:2069
    PS_MITIGATION_OPTION_ROP_CALLER_CHECK PS_MITIGATION_OPTION = 22  //col:2070
    PS_MITIGATION_OPTION_ROP_SIMEXEC PS_MITIGATION_OPTION = 23  //col:2071
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER PS_MITIGATION_OPTION = 24  //col:2072
    PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER_PLUS PS_MITIGATION_OPTION = 25  //col:2073
    PS_MITIGATION_OPTION_RESTRICT_CHILD_PROCESS_CREATION PS_MITIGATION_OPTION = 26  //col:2074
    PS_MITIGATION_OPTION_IMPORT_ADDRESS_FILTER PS_MITIGATION_OPTION = 27  //col:2075
    PS_MITIGATION_OPTION_MODULE_TAMPERING_PROTECTION PS_MITIGATION_OPTION = 28  //col:2076
    PS_MITIGATION_OPTION_RESTRICT_INDIRECT_BRANCH_PREDICTION PS_MITIGATION_OPTION = 29  //col:2077
    PS_MITIGATION_OPTION_SPECULATIVE_STORE_BYPASS_DISABLE // since REDSTONE5 PS_MITIGATION_OPTION = 30  //col:2078
    PS_MITIGATION_OPTION_ALLOW_DOWNGRADE_DYNAMIC_CODE_POLICY PS_MITIGATION_OPTION = 31  //col:2079
    PS_MITIGATION_OPTION_CET_USER_SHADOW_STACKS PS_MITIGATION_OPTION = 32  //col:2080
    PS_MITIGATION_OPTION_USER_CET_SET_CONTEXT_IP_VALIDATION // since 21H1 PS_MITIGATION_OPTION = 33  //col:2081
    PS_MITIGATION_OPTION_BLOCK_NON_CET_BINARIES PS_MITIGATION_OPTION = 34  //col:2082
    PS_MITIGATION_OPTION_CET_DYNAMIC_APIS_OUT_OF_PROC_ONLY PS_MITIGATION_OPTION = 35  //col:2083
    PS_MITIGATION_OPTION_REDIRECTION_TRUST // since 22H1 PS_MITIGATION_OPTION = 36  //col:2084
)


type     PsCreateInitialState uint32
const(
    PsCreateInitialState PS_CREATE_STATE = 1  //col:2090
    PsCreateFailOnFileOpen PS_CREATE_STATE = 2  //col:2091
    PsCreateFailOnSectionCreate PS_CREATE_STATE = 3  //col:2092
    PsCreateFailExeFormat PS_CREATE_STATE = 4  //col:2093
    PsCreateFailMachineMismatch PS_CREATE_STATE = 5  //col:2094
    PsCreateFailExeName // Debugger specified PS_CREATE_STATE = 6  //col:2095
    PsCreateSuccess PS_CREATE_STATE = 7  //col:2096
    PsCreateMaximumStates PS_CREATE_STATE = 8  //col:2097
)


type     MemoryReserveUserApc uint32
const(
    MemoryReserveUserApc MEMORY_RESERVE_TYPE = 1  //col:2477
    MemoryReserveIoCompletion MEMORY_RESERVE_TYPE = 2  //col:2478
    MemoryReserveTypeMax MEMORY_RESERVE_TYPE = 3  //col:2479
)



type (
Ntpsapi interface{
 * Attribution 4.0 International ()(ok bool)//col:92
#if ()(ok bool)//col:229
#if ()(ok bool)//col:292
#if ()(ok bool)//col:300
#if ()(ok bool)//col:315
#define PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS ()(ok bool)//col:407
#if ()(ok bool)//col:478
#if ()(ok bool)//col:677
#define PsProtectedValue()(ok bool)//col:783
#define POWER_THROTTLING_PROCESS_VALID_FLAGS ()(ok bool)//col:853
#define POWER_THROTTLING_THREAD_VALID_FLAGS ()(ok bool)//col:900
#if ()(ok bool)//col:1185
#if ()(ok bool)//col:1361
#if ()(ok bool)//col:1395
#if ()(ok bool)//col:1626
NtQueueApcThreadEx2()(ok bool)//col:1740
    ProcThreadAttributeValue()(ok bool)//col:1858
#define PsAttributeValue()(ok bool)//col:1951
#if ()(ok bool)//col:2294
NtCreateJobObject()(ok bool)//col:2480
}
)

func NewNtpsapi() { return & ntpsapi{} }

func (n *ntpsapi) * Attribution 4.0 International ()(ok bool){//col:92
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTPSAPI_H
#define _NTPSAPI_H
#if (PHNT_MODE == PHNT_MODE_KERNEL)
#define PROCESS_TERMINATE 0x0001
#define PROCESS_CREATE_THREAD 0x0002
#define PROCESS_SET_SESSIONID 0x0004
#define PROCESS_VM_OPERATION 0x0008
#define PROCESS_VM_READ 0x0010
#define PROCESS_VM_WRITE 0x0020
#define PROCESS_CREATE_PROCESS 0x0080
#define PROCESS_SET_QUOTA 0x0100
#define PROCESS_SET_INFORMATION 0x0200
#define PROCESS_QUERY_INFORMATION 0x0400
#define PROCESS_SET_PORT 0x0800
#define PROCESS_SUSPEND_RESUME 0x0800
#define PROCESS_QUERY_LIMITED_INFORMATION 0x1000
#else
#ifndef PROCESS_SET_PORT
#define PROCESS_SET_PORT 0x0800
#endif
#endif
#if (PHNT_MODE == PHNT_MODE_KERNEL)
#define THREAD_QUERY_INFORMATION 0x0040
#define THREAD_SET_THREAD_TOKEN 0x0080
#define THREAD_IMPERSONATE 0x0100
#define THREAD_DIRECT_IMPERSONATION 0x0200
#else
#ifndef THREAD_ALERT
#define THREAD_ALERT 0x0004
#endif
#endif
#if (PHNT_MODE == PHNT_MODE_KERNEL)
#define JOB_OBJECT_ASSIGN_PROCESS 0x0001
#define JOB_OBJECT_SET_ATTRIBUTES 0x0002
#define JOB_OBJECT_QUERY 0x0004
#define JOB_OBJECT_TERMINATE 0x0008
#define JOB_OBJECT_SET_SECURITY_ATTRIBUTES 0x0010
#define JOB_OBJECT_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3F)
#endif
#define GDI_HANDLE_BUFFER_SIZE32 34
#define GDI_HANDLE_BUFFER_SIZE64 60
#ifndef _WIN64
#define GDI_HANDLE_BUFFER_SIZE GDI_HANDLE_BUFFER_SIZE32
#else
#define GDI_HANDLE_BUFFER_SIZE GDI_HANDLE_BUFFER_SIZE64
#endif
typedef ULONG GDI_HANDLE_BUFFER[GDI_HANDLE_BUFFER_SIZE];
typedef ULONG GDI_HANDLE_BUFFER32[GDI_HANDLE_BUFFER_SIZE32];
typedef ULONG GDI_HANDLE_BUFFER64[GDI_HANDLE_BUFFER_SIZE64];
#ifndef FLS_MAXIMUM_AVAILABLE
#define FLS_MAXIMUM_AVAILABLE 128
#endif
#ifndef TLS_MINIMUM_AVAILABLE
#define TLS_MINIMUM_AVAILABLE 64
#endif
#ifndef TLS_EXPANSION_SLOTS
#define TLS_EXPANSION_SLOTS 1024
#endif
typedef struct _PEB_LDR_DATA
{
    ULONG Length;
    BOOLEAN Initialized;
    HANDLE SsHandle;
    LIST_ENTRY InLoadOrderModuleList;
    LIST_ENTRY InMemoryOrderModuleList;
    LIST_ENTRY InInitializationOrderModuleList;
    PVOID EntryInProgress;
    BOOLEAN ShutdownInProgress;
    HANDLE ShutdownThreadId;
} PEB_LDR_DATA, *PPEB_LDR_DATA;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:229
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _PROCESSINFOCLASS
{
    ProcessDynamicFunctionTableInformation,
    ProcessSubsystemProcess,
    ProcessIumChallengeResponse,
    ProcessCaptureTrustletLiveDump,
    ProcessTelemetryCoverage,
    ProcessEnclaveInformation,
    ProcessApplyStateChange,
    ProcessEnableOptionalXStateFeatures,
    ProcessAssignCpuPartitions,
    ProcessPriorityClassEx,
    ProcessMembershipInformation,
    ProcessEffectiveIoPriority,
    ProcessEffectivePagePriority,
    MaxProcessInfoClass
} PROCESSINFOCLASS;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:292
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _THREADINFOCLASS
{
    ThreadEventPair,
    ThreadCSwitchMon,
    ThreadCSwitchPmu,
    ThreadSelectedCpuSets,
    ThreadApplyStateChange,
    ThreadEffectiveIoPriority,
    ThreadEffectivePagePriority,
    MaxThreadInfoClass
} THREADINFOCLASS;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:300
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PAGE_PRIORITY_INFORMATION
{
    ULONG PagePriority;
} PAGE_PRIORITY_INFORMATION, *PPAGE_PRIORITY_INFORMATION;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:315
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PROCESS_BASIC_INFORMATION
{
    NTSTATUS ExitStatus;
    PPEB PebBaseAddress;
    KAFFINITY AffinityMask;
    KPRIORITY BasePriority;
    HANDLE UniqueProcessId;
    HANDLE InheritedFromUniqueProcessId;
} PROCESS_BASIC_INFORMATION, *PPROCESS_BASIC_INFORMATION;*/
return true
}

func (n *ntpsapi)#define PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS ()(ok bool){//col:407
/*#define PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS ((ULONG_PTR)((1UL << PROCESS_EXCEPTION_PORT_ALL_STATE_BITS) - 1))
typedef struct _PROCESS_EXCEPTION_PORT 
{
} PROCESS_EXCEPTION_PORT, *PPROCESS_EXCEPTION_PORT;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:478
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PROCESS_DEVICEMAP_INFORMATION
{
    union
    {
        struct
        {
            HANDLE DirectoryHandle;
        } Set;
        struct
        {
            ULONG DriveMap;
            UCHAR DriveType[32];
        } Query;
    };
} PROCESS_DEVICEMAP_INFORMATION, *PPROCESS_DEVICEMAP_INFORMATION;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:677
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO) && !PHNT_PATCH_FOR_HYPERDBG
typedef struct _PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY
{
    union {
        ULONG Flags;
        struct {
            ULONG EnforceRedirectionTrust : 1;
            ULONG AuditRedirectionTrust : 1;
            ULONG ReservedFlags : 30;
        };
    };
} PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY, * PPROCESS_MITIGATION_REDIRECTION_TRUST_POLICY;*/
return true
}

func (n *ntpsapi)#define PsProtectedValue()(ok bool){//col:783
/*#define PsProtectedValue(aSigner, aAudit, aType) ( \
    ((aSigner & PS_PROTECTED_SIGNER_MASK) << 4) | \
    ((aAudit & PS_PROTECTED_AUDIT_MASK) << 3) | \
    (aType & PS_PROTECTED_TYPE_MASK)\
    )
#define InitializePsProtection(aProtectionLevelPtr, aSigner, aAudit, aType) { \
    (aProtectionLevelPtr)->Signer = aSigner; \
    (aProtectionLevelPtr)->Audit = aAudit; \
    (aProtectionLevelPtr)->Type = aType; \
    }
typedef struct _PS_PROTECTION
{
    union
    {
        UCHAR Level;
        struct
        {
            UCHAR Type : 3;
            UCHAR Audit : 1;
            UCHAR Signer : 4;
        };
    };
} PS_PROTECTION, *PPS_PROTECTION;*/
return true
}

func (n *ntpsapi)#define POWER_THROTTLING_PROCESS_VALID_FLAGS ()(ok bool){//col:853
/*#define POWER_THROTTLING_PROCESS_VALID_FLAGS ((POWER_THROTTLING_PROCESS_EXECUTION_SPEED | POWER_THROTTLING_PROCESS_DELAYTIMERS))
typedef struct _POWER_THROTTLING_PROCESS_STATE
{
    ULONG Version;
    ULONG ControlMask;
    ULONG StateMask;
} POWER_THROTTLING_PROCESS_STATE, *PPOWER_THROTTLING_PROCESS_STATE;*/
return true
}

func (n *ntpsapi)#define POWER_THROTTLING_THREAD_VALID_FLAGS ()(ok bool){//col:900
/*#define POWER_THROTTLING_THREAD_VALID_FLAGS (POWER_THROTTLING_THREAD_EXECUTION_SPEED)
typedef struct _POWER_THROTTLING_THREAD_STATE
{
    ULONG Version;
    ULONG ControlMask;
    ULONG StateMask;
} POWER_THROTTLING_THREAD_STATE, *PPOWER_THROTTLING_THREAD_STATE;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:1185
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _SUBSYSTEM_INFORMATION_TYPE 
{
    SubsystemInformationTypeWin32,
    SubsystemInformationTypeWSL,
    MaxSubsystemInformationType
} SUBSYSTEM_INFORMATION_TYPE;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:1361
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcess(
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
    _In_opt_ HANDLE SectionHandle,
    _In_opt_ HANDLE DebugPort,
    _In_opt_ HANDLE TokenHandle,
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
#define NtCurrentProcess() ((HANDLE)(LONG_PTR)-1)
#define ZwCurrentProcess() NtCurrentProcess()
#define NtCurrentThread() ((HANDLE)(LONG_PTR)-2)
#define ZwCurrentThread() NtCurrentThread()
#define NtCurrentSession() ((HANDLE)(LONG_PTR)-3)
#define ZwCurrentSession() NtCurrentSession()
#define NtCurrentPeb() (NtCurrentTeb()->ProcessEnvironmentBlock)
#define NtCurrentSilo() ((HANDLE)(LONG_PTR)-1)
#define NtCurrentProcessId() (NtCurrentTeb()->ClientId.UniqueProcess)
#define NtCurrentThreadId() (NtCurrentTeb()->ClientId.UniqueThread)
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
#define PROCESS_GET_NEXT_FLAGS_PREVIOUS_PROCESS 0x00000001
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
#define STATECHANGE_SET_ATTRIBUTES 0x0001
typedef enum _PROCESS_STATE_CHANGE_TYPE
{
    ProcessStateChangeSuspend,
    ProcessStateChangeResume,
    ProcessStateChangeMax,
} PROCESS_STATE_CHANGE_TYPE, *PPROCESS_STATE_CHANGE_TYPE;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:1395
/*#if (PHNT_VERSION >= PHNT_WIN11)
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
typedef enum _THREAD_STATE_CHANGE_TYPE
{
    ThreadStateChangeSuspend,
    ThreadStateChangeResume,
    ThreadStateChangeMax,
} THREAD_STATE_CHANGE_TYPE, *PTHREAD_STATE_CHANGE_TYPE;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:1626
/*#if (PHNT_VERSION >= PHNT_WIN11)
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
#define Wow64EncodeApcRoutine(ApcRoutine) \
    ((PVOID)((0 - ((LONG_PTR)(ApcRoutine))) << 2))
#define Wow64DecodeApcRoutine(ApcRoutine) \
    ((PVOID)(0 - (((LONG_PTR)(ApcRoutine)) >> 2)))
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
    _In_ PPS_APC_ROUTINE ApcRoutine,
    _In_opt_ PVOID ApcArgument1,
    _In_opt_ PVOID ApcArgument2,
    _In_opt_ PVOID ApcArgument3
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
typedef enum _QUEUE_USER_APC_FLAGS
{
    QUEUE_USER_APC_FLAGS_NONE = 0x0,
    QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC = 0x1,
} QUEUE_USER_APC_FLAGS;*/
return true
}

func (n *ntpsapi)NtQueueApcThreadEx2()(ok bool){//col:1740
/*NtQueueApcThreadEx2(
    _In_ HANDLE ThreadHandle,
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
#define PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS \
    ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#define PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME \
    ProcThreadAttributeValue(ProcThreadAttributePackageFullName, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#define PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE \
    ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#define PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED \
    ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#define PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM \
    ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#define PROC_THREAD_ATTRIBUTE_BNO_ISOLATION \
    ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#define PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST \
    ProcThreadAttributeValue(ProcThreadAttributeIsolationManifest, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#define PROC_THREAD_ATTRIBUTE_CREATE_STORE \
    ProcThreadAttributeValue(ProcThreadAttributeCreateStore, FALSE, TRUE, FALSE)
#endif
typedef struct _PROC_THREAD_ATTRIBUTE {
    ULONG_PTR Attribute;
    SIZE_T Size;
    ULONG_PTR Value;
} PROC_THREAD_ATTRIBUTE, *PPROC_THREAD_ATTRIBUTE;*/
return true
}

func (n *ntpsapi)    ProcThreadAttributeValue()(ok bool){//col:1858
/*    ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#define PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE \
    ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#define PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED \
    ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#define PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM \
    ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#define PROC_THREAD_ATTRIBUTE_BNO_ISOLATION \
    ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE)
#endif
typedef enum _PS_ATTRIBUTE_NUM
{
    PsAttributeComponentFilter,
    PsAttributeMax
} PS_ATTRIBUTE_NUM;*/
return true
}

func (n *ntpsapi)#define PsAttributeValue()(ok bool){//col:1951
/*#define PsAttributeValue(Number, Thread, Input, Additive) \
    (((Number) & PS_ATTRIBUTE_NUMBER_MASK) | \
    ((Thread) ? PS_ATTRIBUTE_THREAD : 0) | \
    ((Input) ? PS_ATTRIBUTE_INPUT : 0) | \
    ((Additive) ? PS_ATTRIBUTE_ADDITIVE : 0))
#define PS_ATTRIBUTE_PARENT_PROCESS \
    PsAttributeValue(PsAttributeParentProcess, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_DEBUG_OBJECT \
    PsAttributeValue(PsAttributeDebugObject, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_TOKEN \
    PsAttributeValue(PsAttributeToken, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_CLIENT_ID \
    PsAttributeValue(PsAttributeClientId, TRUE, FALSE, FALSE)
#define PS_ATTRIBUTE_TEB_ADDRESS \
    PsAttributeValue(PsAttributeTebAddress, TRUE, FALSE, FALSE)
#define PS_ATTRIBUTE_IMAGE_NAME \
    PsAttributeValue(PsAttributeImageName, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_IMAGE_INFO \
    PsAttributeValue(PsAttributeImageInfo, FALSE, FALSE, FALSE)
#define PS_ATTRIBUTE_MEMORY_RESERVE \
    PsAttributeValue(PsAttributeMemoryReserve, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_PRIORITY_CLASS \
    PsAttributeValue(PsAttributePriorityClass, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ERROR_MODE \
    PsAttributeValue(PsAttributeErrorMode, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_STD_HANDLE_INFO \
    PsAttributeValue(PsAttributeStdHandleInfo, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_HANDLE_LIST \
    PsAttributeValue(PsAttributeHandleList, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_GROUP_AFFINITY \
    PsAttributeValue(PsAttributeGroupAffinity, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_PREFERRED_NODE \
    PsAttributeValue(PsAttributePreferredNode, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_IDEAL_PROCESSOR \
    PsAttributeValue(PsAttributeIdealProcessor, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_UMS_THREAD \
    PsAttributeValue(PsAttributeUmsThread, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_MITIGATION_OPTIONS \
    PsAttributeValue(PsAttributeMitigationOptions, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_PROTECTION_LEVEL \
    PsAttributeValue(PsAttributeProtectionLevel, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_SECURE_PROCESS \
    PsAttributeValue(PsAttributeSecureProcess, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_JOB_LIST \
    PsAttributeValue(PsAttributeJobList, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_CHILD_PROCESS_POLICY \
    PsAttributeValue(PsAttributeChildProcessPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ALL_APPLICATION_PACKAGES_POLICY \
    PsAttributeValue(PsAttributeAllApplicationPackagesPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_WIN32K_FILTER \
    PsAttributeValue(PsAttributeWin32kFilter, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM \
    PsAttributeValue(PsAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_BNO_ISOLATION \
    PsAttributeValue(PsAttributeBnoIsolation, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_DESKTOP_APP_POLICY \
    PsAttributeValue(PsAttributeDesktopAppPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_CHPE \
    PsAttributeValue(PsAttributeChpe, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_MITIGATION_AUDIT_OPTIONS \
    PsAttributeValue(PsAttributeMitigationAuditOptions, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_MACHINE_TYPE \
    PsAttributeValue(PsAttributeMachineType, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_COMPONENT_FILTER \
    PsAttributeValue(PsAttributeComponentFilter, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ENABLE_OPTIONAL_XSTATE_FEATURES \
    PsAttributeValue(PsAttributeEnableOptionalXStateFeatures, TRUE, TRUE, FALSE)
typedef struct _PS_ATTRIBUTE
{
    ULONG_PTR Attribute;
    SIZE_T Size;
    union
    {
        ULONG_PTR Value;
        PVOID ValuePtr;
    };
    PSIZE_T ReturnLength;
} PS_ATTRIBUTE, *PPS_ATTRIBUTE;*/
return true
}

func (n *ntpsapi)#if ()(ok bool){//col:2294
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
    _In_opt_ PVOID Argument,
    _In_ SIZE_T ZeroBits,
    _In_ SIZE_T StackSize,
    _In_ SIZE_T MaximumStackSize,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define JobObjectCompletionFilter 16
#define JobObjectCompletionCounter 17
#define JobObjectBackgroundInformation 21
#define JobObjectSchedulingRankBiasInformation 22
#define JobObjectTimerVirtualizationInformation 23
#define JobObjectCycleTimeNotification 24
#define JobObjectClearEvent 25
#define JobObjectClearPeakJobMemoryUsed 27
#define JobObjectSharedCommit 29
#define JobObjectContainerId 30
#define JobObjectIoRateControlInformation 31
#define JobObjectCreateSilo 35
#define JobObjectServerSiloInitialize 40
#define JobObjectServerSiloRunningState 41
#define JobObjectIoAttribution 42
#define JobObjectMemoryPartitionInformation 43
#define JobObjectContainerTelemetryId 44
#define JobObjectSiloSystemRoot 45
#define JobObjectThreadImpersonationInformation 47
#define JobObjectIoPriorityLimit 48
#define JobObjectPagePriorityLimit 49
#define MaxJobObjectInfoClass 50
typedef struct _JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION
{
    JOBOBJECT_BASIC_ACCOUNTING_INFORMATION BasicInfo;
    IO_COUNTERS IoInfo;
    PROCESS_DISK_COUNTERS DiskIoInfo;
    ULONG64 ContextSwitches;
    LARGE_INTEGER TotalCycleTime;
    ULONG64 ReadyTime;
    PROCESS_ENERGY_VALUES EnergyValues;
} JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION, *PJOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION;*/
return true
}

func (n *ntpsapi)NtCreateJobObject()(ok bool){//col:2480
/*NtCreateJobObject(
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
typedef enum _MEMORY_RESERVE_TYPE
{
    MemoryReserveUserApc,
    MemoryReserveIoCompletion,
    MemoryReserveTypeMax
} MEMORY_RESERVE_TYPE;*/
return true
}



