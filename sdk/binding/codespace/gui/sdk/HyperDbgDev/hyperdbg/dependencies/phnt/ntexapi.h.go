package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntexapi.h.back

const(
_NTEXAPI_H =  //col:1
EFI_VARIABLE_NON_VOLATILE = 0x00000001 //col:2
EFI_VARIABLE_BOOTSERVICE_ACCESS = 0x00000002 //col:3
EFI_VARIABLE_RUNTIME_ACCESS = 0x00000004 //col:4
EFI_VARIABLE_HARDWARE_ERROR_RECORD = 0x00000008 //col:5
EFI_VARIABLE_AUTHENTICATED_WRITE_ACCESS = 0x00000010 //col:6
EFI_VARIABLE_TIME_BASED_AUTHENTICATED_WRITE_ACCESS = 0x00000020 //col:7
EFI_VARIABLE_APPEND_WRITE = 0x00000040 //col:8
EFI_VARIABLE_ENHANCED_AUTHENTICATED_ACCESS = 0x00000080 //col:9
EVENT_QUERY_STATE = 0x0001 //col:10
EVENT_MODIFY_STATE = 0x0002 //col:11
EVENT_ALL_ACCESS = (EVENT_QUERY_STATE|EVENT_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:12
EVENT_PAIR_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE) //col:13
MUTANT_QUERY_STATE = 0x0001 //col:14
MUTANT_ALL_ACCESS = (MUTANT_QUERY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:15
SEMAPHORE_QUERY_STATE = 0x0001 //col:16
SEMAPHORE_MODIFY_STATE = 0x0002 //col:17
SEMAPHORE_ALL_ACCESS = (SEMAPHORE_QUERY_STATE|SEMAPHORE_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:18
TIMER_QUERY_STATE = 0x0001 //col:19
TIMER_MODIFY_STATE = 0x0002 //col:20
TIMER_ALL_ACCESS = (TIMER_QUERY_STATE|TIMER_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:21
PROFILE_CONTROL = 0x0001 //col:22
PROFILE_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | PROFILE_CONTROL) //col:23
KEYEDEVENT_WAIT = 0x0001 //col:24
KEYEDEVENT_WAKE = 0x0002 //col:25
KEYEDEVENT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | KEYEDEVENT_WAIT | KEYEDEVENT_WAKE) //col:26
WORKER_FACTORY_RELEASE_WORKER = 0x0001 //col:28
WORKER_FACTORY_WAIT = 0x0002 //col:29
WORKER_FACTORY_SET_INFORMATION = 0x0004 //col:30
WORKER_FACTORY_QUERY_INFORMATION = 0x0008 //col:31
WORKER_FACTORY_READY_WORKER = 0x0010 //col:32
WORKER_FACTORY_SHUTDOWN = 0x0020 //col:33
WORKER_FACTORY_ALL_ACCESS ( = STANDARD_RIGHTS_REQUIRED | WORKER_FACTORY_RELEASE_WORKER | WORKER_FACTORY_WAIT | WORKER_FACTORY_SET_INFORMATION | WORKER_FACTORY_QUERY_INFORMATION | WORKER_FACTORY_READY_WORKER | WORKER_FACTORY_SHUTDOWN ) //col:34
MM_WORKING_SET_MAX_HARD_ENABLE = 0x1 //col:43
MM_WORKING_SET_MAX_HARD_DISABLE = 0x2 //col:44
MM_WORKING_SET_MIN_HARD_ENABLE = 0x4 //col:45
MM_WORKING_SET_MIN_HARD_DISABLE = 0x8 //col:46
_TRACEHANDLE_DEFINED =  //col:47
PERF_MASK_INDEX = (0xe0000000) //col:48
PERF_MASK_GROUP = (~PERF_MASK_INDEX) //col:49
PERF_NUM_MASKS = 8 //col:50
PERF_GET_MASK_INDEX(GM) = (((GM) & PERF_MASK_INDEX) >> 29) //col:51
PERF_GET_MASK_GROUP(GM) = ((GM) & PERF_MASK_GROUP) //col:52
PERFINFO_OR_GROUP_WITH_GROUPMASK(Group, pGroupMask) = (pGroupMask)->Masks[PERF_GET_MASK_INDEX(Group)] |= PERF_GET_MASK_GROUP(Group); //col:53
PERF_PROCESS =            EVENT_TRACE_FLAG_PROCESS //col:55
PERF_THREAD =             EVENT_TRACE_FLAG_THREAD //col:56
PERF_PROC_THREAD =        EVENT_TRACE_FLAG_PROCESS | EVENT_TRACE_FLAG_THREAD //col:57
PERF_LOADER =             EVENT_TRACE_FLAG_IMAGE_LOAD //col:58
PERF_PERF_COUNTER =       EVENT_TRACE_FLAG_PROCESS_COUNTERS //col:59
PERF_FILENAME =           EVENT_TRACE_FLAG_DISK_FILE_IO //col:60
PERF_DISK_IO =            EVENT_TRACE_FLAG_DISK_FILE_IO | EVENT_TRACE_FLAG_DISK_IO //col:61
PERF_DISK_IO_INIT =       EVENT_TRACE_FLAG_DISK_IO_INIT //col:62
PERF_ALL_FAULTS =         EVENT_TRACE_FLAG_MEMORY_PAGE_FAULTS //col:63
PERF_HARD_FAULTS =        EVENT_TRACE_FLAG_MEMORY_HARD_FAULTS //col:64
PERF_VAMAP =              EVENT_TRACE_FLAG_VAMAP //col:65
PERF_NETWORK =            EVENT_TRACE_FLAG_NETWORK_TCPIP //col:66
PERF_REGISTRY =           EVENT_TRACE_FLAG_REGISTRY //col:67
PERF_DBGPRINT =           EVENT_TRACE_FLAG_DBGPRINT //col:68
PERF_JOB =                EVENT_TRACE_FLAG_JOB //col:69
PERF_ALPC =               EVENT_TRACE_FLAG_ALPC //col:70
PERF_SPLIT_IO =           EVENT_TRACE_FLAG_SPLIT_IO //col:71
PERF_DEBUG_EVENTS =       EVENT_TRACE_FLAG_DEBUG_EVENTS //col:72
PERF_FILE_IO =            EVENT_TRACE_FLAG_FILE_IO //col:73
PERF_FILE_IO_INIT =       EVENT_TRACE_FLAG_FILE_IO_INIT //col:74
PERF_NO_SYSCONFIG =       EVENT_TRACE_FLAG_NO_SYSCONFIG //col:75
PERF_MEMORY =             0x20000001 //col:76
PERF_PROFILE =            0x20000002 //col:77
PERF_CONTEXT_SWITCH =     0x20000004 //col:78
PERF_FOOTPRINT =          0x20000008 //col:79
PERF_DRIVERS =            0x20000010 //col:80
PERF_REFSET =             0x20000020 //col:81
PERF_POOL =               0x20000040 //col:82
PERF_POOLTRACE =          0x20000041 //col:83
PERF_DPC =                0x20000080 //col:84
PERF_COMPACT_CSWITCH =    0x20000100 //col:85
PERF_DISPATCHER =         0x20000200 //col:86
PERF_PMC_PROFILE =        0x20000400 //col:87
PERF_PROFILING =          0x20000402 //col:88
PERF_PROCESS_INSWAP =     0x20000800 //col:89
PERF_AFFINITY =           0x20001000 //col:90
PERF_PRIORITY =           0x20002000 //col:91
PERF_INTERRUPT =          0x20004000 //col:92
PERF_VIRTUAL_ALLOC =      0x20008000 //col:93
PERF_SPINLOCK =           0x20010000 //col:94
PERF_SYNC_OBJECTS =       0x20020000 //col:95
PERF_DPC_QUEUE =          0x20040000 //col:96
PERF_MEMINFO =            0x20080000 //col:97
PERF_CONTMEM_GEN =        0x20100000 //col:98
PERF_SPINLOCK_CNTRS =     0x20200000 //col:99
PERF_SPININSTR =          0x20210000 //col:100
PERF_SESSION =            0x20400000 //col:101
PERF_PFSECTION =          0x20400000 //col:102
PERF_MEMINFO_WS =         0x20800000 //col:103
PERF_KERNEL_QUEUE =       0x21000000 //col:104
PERF_INTERRUPT_STEER =    0x22000000 //col:105
PERF_SHOULD_YIELD =       0x24000000 //col:106
PERF_WS =                 0x28000000 //col:107
PERF_ANTI_STARVATION =    0x40000001 //col:108
PERF_PROCESS_FREEZE =     0x40000002 //col:109
PERF_PFN_LIST =           0x40000004 //col:110
PERF_WS_DETAIL =          0x40000008 //col:111
PERF_WS_ENTRY =           0x40000010 //col:112
PERF_HEAP =               0x40000020 //col:113
PERF_SYSCALL =            0x40000040 //col:114
PERF_UMS =                0x40000080 //col:115
PERF_BACKTRACE =          0x40000100 //col:116
PERF_VULCAN =             0x40000200 //col:117
PERF_OBJECTS =            0x40000400 //col:118
PERF_EVENTS =             0x40000800 //col:119
PERF_FULLTRACE =          0x40001000 //col:120
PERF_DFSS =               0x40002000 //col:121
PERF_PREFETCH =           0x40004000 //col:122
PERF_PROCESSOR_IDLE =     0x40008000 //col:123
PERF_CPU_CONFIG =         0x40010000 //col:124
PERF_TIMER =              0x40020000 //col:125
PERF_CLOCK_INTERRUPT =    0x40040000 //col:126
PERF_LOAD_BALANCER =      0x40080000 //col:127
PERF_CLOCK_TIMER =        0x40100000 //col:128
PERF_IDLE_SELECTION =     0x40200000 //col:129
PERF_IPI =                0x40400000 //col:130
PERF_IO_TIMER =           0x40800000 //col:131
PERF_REG_HIVE =           0x41000000 //col:132
PERF_REG_NOTIF =          0x42000000 //col:133
PERF_PPM_EXIT_LATENCY =   0x44000000 //col:134
PERF_WORKER_THREAD =      0x48000000 //col:135
PERF_OPTICAL_IO =         0x80000001 //col:136
PERF_OPTICAL_IO_INIT =    0x80000002 //col:137
PERF_DLL_INFO =           0x80000008 //col:138
PERF_DLL_FLUSH_WS =       0x80000010 //col:139
PERF_OB_HANDLE =          0x80000040 //col:140
PERF_OB_OBJECT =          0x80000080 //col:141
PERF_WAKE_DROP =          0x80000200 //col:142
PERF_WAKE_EVENT =         0x80000400 //col:143
PERF_DEBUGGER =           0x80000800 //col:144
PERF_PROC_ATTACH =        0x80001000 //col:145
PERF_WAKE_COUNTER =       0x80002000 //col:146
PERF_POWER =              0x80008000 //col:147
PERF_SOFT_TRIM =          0x80010000 //col:148
PERF_CC =                 0x80020000 //col:149
PERF_FLT_IO_INIT =        0x80080000 //col:150
PERF_FLT_IO =             0x80100000 //col:151
PERF_FLT_FASTIO =         0x80200000 //col:152
PERF_FLT_IO_FAILURE =     0x80400000 //col:153
PERF_HV_PROFILE =         0x80800000 //col:154
PERF_WDF_DPC =            0x81000000 //col:155
PERF_WDF_INTERRUPT =      0x82000000 //col:156
PERF_CACHE_FLUSH =        0x84000000 //col:157
PERF_HIBER_RUNDOWN =      0xA0000001 //col:158
PERF_SYSCFG_SYSTEM =      0xC0000001 //col:159
PERF_SYSCFG_GRAPHICS =    0xC0000002 //col:160
PERF_SYSCFG_STORAGE =     0xC0000004 //col:161
PERF_SYSCFG_NETWORK =     0xC0000008 //col:162
PERF_SYSCFG_SERVICES =    0xC0000010 //col:163
PERF_SYSCFG_PNP =         0xC0000020 //col:164
PERF_SYSCFG_OPTICAL =     0xC0000040 //col:165
PERF_SYSCFG_ALL =         0xDFFFFFFF //col:166
PERF_CLUSTER_OFF =        0xE0000001 //col:167
PERF_MEMORY_CONTROL =     0xE0000002 //col:168
MAXIMUM_NODE_COUNT = 0x40 //col:169
MAXIMUM_NODE_COUNT = 0x10 //col:170
CODEINTEGRITY_OPTION_ENABLED = 0x01 //col:171
CODEINTEGRITY_OPTION_TESTSIGN = 0x02 //col:172
CODEINTEGRITY_OPTION_UMCI_ENABLED = 0x04 //col:173
CODEINTEGRITY_OPTION_UMCI_AUDITMODE_ENABLED = 0x08 //col:174
CODEINTEGRITY_OPTION_UMCI_EXCLUSIONPATHS_ENABLED = 0x10 //col:175
CODEINTEGRITY_OPTION_TEST_BUILD = 0x20 //col:176
CODEINTEGRITY_OPTION_PREPRODUCTION_BUILD = 0x40 //col:177
CODEINTEGRITY_OPTION_DEBUGMODE_ENABLED = 0x80 //col:178
CODEINTEGRITY_OPTION_FLIGHT_BUILD = 0x100 //col:179
CODEINTEGRITY_OPTION_FLIGHTING_ENABLED = 0x200 //col:180
CODEINTEGRITY_OPTION_HVCI_KMCI_ENABLED = 0x400 //col:181
CODEINTEGRITY_OPTION_HVCI_KMCI_AUDITMODE_ENABLED = 0x800 //col:182
CODEINTEGRITY_OPTION_HVCI_KMCI_STRICTMODE_ENABLED = 0x1000 //col:183
CODEINTEGRITY_OPTION_HVCI_IUM_ENABLED = 0x2000 //col:184
CODEINTEGRITY_OPTION_WHQL_ENFORCEMENT_ENABLED = 0x4000 //col:185
CODEINTEGRITY_OPTION_WHQL_AUDITMODE_ENABLED = 0x8000 //col:186
SYSTEM_STORE_INFORMATION_VERSION = 1 //col:187
SYSTEM_STORE_STATS_INFORMATION_VERSION = 2 //col:188
SYSTEM_STORE_CREATE_INFORMATION_VERSION = 6 //col:189
SYSTEM_STORE_DELETE_INFORMATION_VERSION = 1 //col:190
SYSTEM_STORE_LIST_INFORMATION_VERSION = 2 //col:191
SYSTEM_CACHE_LIST_INFORMATION_VERSION = 2 //col:192
SYSTEM_CACHE_CREATE_INFORMATION_VERSION = 3 //col:193
SYSTEM_CACHE_DELETE_INFORMATION_VERSION = 1 //col:194
SYSTEM_CACHE_STORE_CREATE_INFORMATION_VERSION = 2 //col:195
SYSTEM_CACHE_STORE_DELETE_INFORMATION_VERSION = 1 //col:196
SYSTEM_CACHE_STATS_INFORMATION_VERSION = 3 //col:197
SYSTEM_STORE_REGISTRATION_INFORMATION_VERSION = 2 //col:198
SYSTEM_STORE_RESIZE_INFORMATION_VERSION = 6 //col:199
SYSTEM_CACHE_STORE_RESIZE_INFORMATION_VERSION = 1 //col:200
SYSTEM_STORE_CONFIG_INFORMATION_VERSION = 4 //col:201
SYSTEM_STORE_HIGH_MEM_PRIORITY_INFORMATION_VERSION = 1 //col:202
SYSTEM_STORE_TRIM_INFORMATION_VERSION = 1 //col:203
SYSTEM_STORE_COMPRESSION_INFORMATION_VERSION = 3 //col:204
MEMORY_COMBINE_FLAGS_COMMON_PAGES_ONLY = 0x4 //col:205
SYSDBG_LIVEDUMP_CONTROL_VERSION = 1 //col:206
SYSDBG_LIVEDUMP_CONTROL_VERSION_WIN11 = 2 //col:207
HARDERROR_OVERRIDE_ERRORMODE = 0x10000000 //col:208
PROCESSOR_FEATURE_MAX = 64 //col:209
MAX_WOW64_SHARED_ENTRIES = 16 //col:210
NX_SUPPORT_POLICY_ALWAYSOFF = 0 //col:211
NX_SUPPORT_POLICY_ALWAYSON = 1 //col:212
NX_SUPPORT_POLICY_OPTIN = 2 //col:213
NX_SUPPORT_POLICY_OPTOUT = 3 //col:214
USER_SHARED_DATA = ((KUSER_SHARED_DATA * const)0x7ffe0000) //col:215
ATOM_FLAG_GLOBAL = 0x2 //col:216
FLG_STOP_ON_EXCEPTION = 0x00000001 //col:217
FLG_SHOW_LDR_SNAPS = 0x00000002 //col:218
FLG_DEBUG_INITIAL_COMMAND = 0x00000004 //col:219
FLG_STOP_ON_HUNG_GUI = 0x00000008 //col:220
FLG_HEAP_ENABLE_TAIL_CHECK = 0x00000010 //col:221
FLG_HEAP_ENABLE_FREE_CHECK = 0x00000020 //col:222
FLG_HEAP_VALIDATE_PARAMETERS = 0x00000040 //col:223
FLG_HEAP_VALIDATE_ALL = 0x00000080 //col:224
FLG_APPLICATION_VERIFIER = 0x00000100 //col:225
FLG_MONITOR_SILENT_PROCESS_EXIT = 0x00000200 //col:226
FLG_POOL_ENABLE_TAGGING = 0x00000400 //col:227
FLG_HEAP_ENABLE_TAGGING = 0x00000800 //col:228
FLG_USER_STACK_TRACE_DB = 0x00001000 //col:229
FLG_KERNEL_STACK_TRACE_DB = 0x00002000 //col:230
FLG_MAINTAIN_OBJECT_TYPELIST = 0x00004000 //col:231
FLG_HEAP_ENABLE_TAG_BY_DLL = 0x00008000 //col:232
FLG_DISABLE_STACK_EXTENSION = 0x00010000 //col:233
FLG_ENABLE_CSRDEBUG = 0x00020000 //col:234
FLG_ENABLE_KDEBUG_SYMBOL_LOAD = 0x00040000 //col:235
FLG_DISABLE_PAGE_KERNEL_STACKS = 0x00080000 //col:236
FLG_ENABLE_SYSTEM_CRIT_BREAKS = 0x00100000 //col:237
FLG_HEAP_DISABLE_COALESCING = 0x00200000 //col:238
FLG_ENABLE_CLOSE_EXCEPTIONS = 0x00400000 //col:239
FLG_ENABLE_EXCEPTION_LOGGING = 0x00800000 //col:240
FLG_ENABLE_HANDLE_TYPE_TAGGING = 0x01000000 //col:241
FLG_HEAP_PAGE_ALLOCS = 0x02000000 //col:242
FLG_DEBUG_INITIAL_COMMAND_EX = 0x04000000 //col:243
FLG_DISABLE_DBGPRINT = 0x08000000 //col:244
FLG_CRITSEC_EVENT_CREATION = 0x10000000 //col:245
FLG_STOP_ON_UNHANDLED_EXCEPTION = 0x20000000 //col:246
FLG_ENABLE_HANDLE_EXCEPTIONS = 0x40000000 //col:247
FLG_DISABLE_PROTDLLS = 0x80000000 //col:248
FLG_VALID_BITS = 0xfffffdff //col:249
FLG_USERMODE_VALID_BITS (FLG_STOP_ON_EXCEPTION | = FLG_SHOW_LDR_SNAPS | FLG_HEAP_ENABLE_TAIL_CHECK | FLG_HEAP_ENABLE_FREE_CHECK | FLG_HEAP_VALIDATE_PARAMETERS | FLG_HEAP_VALIDATE_ALL | FLG_APPLICATION_VERIFIER | FLG_HEAP_ENABLE_TAGGING | FLG_USER_STACK_TRACE_DB | FLG_HEAP_ENABLE_TAG_BY_DLL | FLG_DISABLE_STACK_EXTENSION | FLG_ENABLE_SYSTEM_CRIT_BREAKS | FLG_HEAP_DISABLE_COALESCING | FLG_DISABLE_PROTDLLS | FLG_HEAP_PAGE_ALLOCS | FLG_CRITSEC_EVENT_CREATION | FLG_LDR_TOP_DOWN) //col:250
FLG_BOOTONLY_VALID_BITS (FLG_KERNEL_STACK_TRACE_DB | = FLG_MAINTAIN_OBJECT_TYPELIST | FLG_ENABLE_CSRDEBUG | FLG_DEBUG_INITIAL_COMMAND | FLG_DEBUG_INITIAL_COMMAND_EX | FLG_DISABLE_PAGE_KERNEL_STACKS) //col:267
FLG_KERNELMODE_VALID_BITS (FLG_STOP_ON_EXCEPTION | = FLG_SHOW_LDR_SNAPS | FLG_STOP_ON_HUNG_GUI | FLG_POOL_ENABLE_TAGGING | FLG_ENABLE_KDEBUG_SYMBOL_LOAD | FLG_ENABLE_CLOSE_EXCEPTIONS | FLG_ENABLE_EXCEPTION_LOGGING | FLG_ENABLE_HANDLE_TYPE_TAGGING | FLG_DISABLE_DBGPRINT | FLG_ENABLE_HANDLE_EXCEPTIONS) //col:273
)

const(
    FilterBootOptionOperationOpenSystemStore = 1  //col:3
    FilterBootOptionOperationSetElement = 2  //col:4
    FilterBootOptionOperationDeleteElement = 3  //col:5
    FilterBootOptionOperationMax = 4  //col:6
)


const(
    EventBasicInformation = 1  //col:10
)


const(
    MutantBasicInformation  = 1  //col:14
    MutantOwnerInformation  = 2  //col:15
)


const(
    SemaphoreBasicInformation = 1  //col:19
)


const(
    TimerBasicInformation  = 1  //col:23
)


const(
    TimerSetCoalescableTimer  = 1  //col:27
    MaxTimerInfoClass = 2  //col:28
)


const(
    WnfWellKnownStateName = 1  //col:32
    WnfPermanentStateName = 2  //col:33
    WnfPersistentStateName = 3  //col:34
    WnfTemporaryStateName = 4  //col:35
)


const(
    WnfInfoStateNameExist = 1  //col:39
    WnfInfoSubscribersPresent = 2  //col:40
    WnfInfoIsQuiescent = 3  //col:41
)


const(
    WnfDataScopeSystem = 1  //col:45
    WnfDataScopeSession = 2  //col:46
    WnfDataScopeUser = 3  //col:47
    WnfDataScopeProcess = 4  //col:48
    WnfDataScopeMachine  = 5  //col:49
    WnfDataScopePhysicalMachine  = 6  //col:50
)


const(
    WorkerFactoryTimeout  = 1  //col:54
    WorkerFactoryRetryTimeout  = 2  //col:55
    WorkerFactoryIdleTimeout  = 3  //col:56
    WorkerFactoryBindingCount  = 4  //col:57
    WorkerFactoryThreadMinimum  = 5  //col:58
    WorkerFactoryThreadMaximum  = 6  //col:59
    WorkerFactoryPaused  = 7  //col:60
    WorkerFactoryBasicInformation  = 8  //col:61
    WorkerFactoryAdjustThreadGoal = 9  //col:62
    WorkerFactoryCallbackType = 10  //col:63
    WorkerFactoryStackInformation  = 11  //col:64
    WorkerFactoryThreadBasePriority  = 12  //col:65
    WorkerFactoryTimeoutWaiters  = 13  //col:66
    WorkerFactoryFlags  = 14  //col:67
    WorkerFactoryThreadSoftMaximum  = 15  //col:68
    WorkerFactoryThreadCpuSets  = 16  //col:69
    MaxWorkerFactoryInfoClass = 17  //col:70
)


const(
    SystemBasicInformation  = 1  //col:74
    SystemProcessorInformation  = 2  //col:75
    SystemPerformanceInformation  = 3  //col:76
    SystemTimeOfDayInformation  = 4  //col:77
    SystemPathInformation  = 5  //col:78
    SystemProcessInformation  = 6  //col:79
    SystemCallCountInformation  = 7  //col:80
    SystemDeviceInformation  = 8  //col:81
    SystemProcessorPerformanceInformation  = 9  //col:82
    SystemFlagsInformation  = 10  //col:83
    SystemCallTimeInformation  = 11  //col:84
    SystemModuleInformation  = 12  //col:85
    SystemLocksInformation  = 13  //col:86
    SystemStackTraceInformation  = 14  //col:87
    SystemPagedPoolInformation  = 15  //col:88
    SystemNonPagedPoolInformation  = 16  //col:89
    SystemHandleInformation  = 17  //col:90
    SystemObjectInformation  = 18  //col:91
    SystemPageFileInformation  = 19  //col:92
    SystemVdmInstemulInformation  = 20  //col:93
    SystemVdmBopInformation  = 21  //col:94
    SystemFileCacheInformation  = 22  //col:95
    SystemPoolTagInformation  = 23  //col:96
    SystemInterruptInformation  = 24  //col:97
    SystemDpcBehaviorInformation  = 25  //col:98
    SystemFullMemoryInformation  = 26  //col:99
    SystemLoadGdiDriverInformation  = 27  //col:100
    SystemUnloadGdiDriverInformation  = 28  //col:101
    SystemTimeAdjustmentInformation  = 29  //col:102
    SystemSummaryMemoryInformation  = 30  //col:103
    SystemMirrorMemoryInformation  = 31  //col:104
    SystemPerformanceTraceInformation  = 32  //col:105
    SystemObsolete0  = 33  //col:106
    SystemExceptionInformation  = 34  //col:107
    SystemCrashDumpStateInformation  = 35  //col:108
    SystemKernelDebuggerInformation  = 36  //col:109
    SystemContextSwitchInformation  = 37  //col:110
    SystemRegistryQuotaInformation  = 38  //col:111
    SystemExtendServiceTableInformation  = 39  //col:112
    SystemPrioritySeperation  = 40  //col:113
    SystemVerifierAddDriverInformation  = 41  //col:114
    SystemVerifierRemoveDriverInformation  = 42  //col:115
    SystemProcessorIdleInformation  = 43  //col:116
    SystemLegacyDriverInformation  = 44  //col:117
    SystemCurrentTimeZoneInformation  = 45  //col:118
    SystemLookasideInformation  = 46  //col:119
    SystemTimeSlipNotification  = 47  //col:120
    SystemSessionCreate  = 48  //col:121
    SystemSessionDetach  = 49  //col:122
    SystemSessionInformation  = 50  //col:123
    SystemRangeStartInformation  = 51  //col:124
    SystemVerifierInformation  = 52  //col:125
    SystemVerifierThunkExtend  = 53  //col:126
    SystemSessionProcessInformation  = 54  //col:127
    SystemLoadGdiDriverInSystemSpace  = 55  //col:128
    SystemNumaProcessorMap  = 56  //col:129
    SystemPrefetcherInformation  = 57  //col:130
    SystemExtendedProcessInformation  = 58  //col:131
    SystemRecommendedSharedDataAlignment  = 59  //col:132
    SystemComPlusPackage  = 60  //col:133
    SystemNumaAvailableMemory  = 61  //col:134
    SystemProcessorPowerInformation  = 62  //col:135
    SystemEmulationBasicInformation  = 63  //col:136
    SystemEmulationProcessorInformation  = 64  //col:137
    SystemExtendedHandleInformation  = 65  //col:138
    SystemLostDelayedWriteInformation  = 66  //col:139
    SystemBigPoolInformation  = 67  //col:140
    SystemSessionPoolTagInformation  = 68  //col:141
    SystemSessionMappedViewInformation  = 69  //col:142
    SystemHotpatchInformation  = 70  //col:143
    SystemObjectSecurityMode  = 71  //col:144
    SystemWatchdogTimerHandler  = 72  //col:145
    SystemWatchdogTimerInformation  = 73  //col:146
    SystemLogicalProcessorInformation  = 74  //col:147
    SystemWow64SharedInformationObsolete  = 75  //col:148
    SystemRegisterFirmwareTableInformationHandler  = 76  //col:149
    SystemFirmwareTableInformation  = 77  //col:150
    SystemModuleInformationEx  = 78  //col:151
    SystemVerifierTriageInformation  = 79  //col:152
    SystemSuperfetchInformation  = 80  //col:153
    SystemMemoryListInformation  = 81  //col:154
    SystemFileCacheInformationEx  = 82  //col:155
    SystemThreadPriorityClientIdInformation  = 83  //col:156
    SystemProcessorIdleCycleTimeInformation  = 84  //col:157
    SystemVerifierCancellationInformation  = 85  //col:158
    SystemProcessorPowerInformationEx  = 86  //col:159
    SystemRefTraceInformation  = 87  //col:160
    SystemSpecialPoolInformation  = 88  //col:161
    SystemProcessIdInformation  = 89  //col:162
    SystemErrorPortInformation  = 90  //col:163
    SystemBootEnvironmentInformation  = 91  //col:164
    SystemHypervisorInformation  = 92  //col:165
    SystemVerifierInformationEx  = 93  //col:166
    SystemTimeZoneInformation  = 94  //col:167
    SystemImageFileExecutionOptionsInformation  = 95  //col:168
    SystemCoverageInformation  = 96  //col:169
    SystemPrefetchPatchInformation  = 97  //col:170
    SystemVerifierFaultsInformation  = 98  //col:171
    SystemSystemPartitionInformation  = 99  //col:172
    SystemSystemDiskInformation  = 100  //col:173
    SystemProcessorPerformanceDistribution  = 101  //col:174
    SystemNumaProximityNodeInformation  = 102  //col:175
    SystemDynamicTimeZoneInformation  = 103  //col:176
    SystemCodeIntegrityInformation  = 104  //col:177
    SystemProcessorMicrocodeUpdateInformation  = 105  //col:178
    SystemProcessorBrandString  = 106  //col:179
    SystemVirtualAddressInformation  = 107  //col:180
    SystemLogicalProcessorAndGroupInformation  = 108  //col:181
    SystemProcessorCycleTimeInformation  = 109  //col:182
    SystemStoreInformation  = 110  //col:183
    SystemRegistryAppendString  = 111  //col:184
    SystemAitSamplingValue  = 112  //col:185
    SystemVhdBootInformation  = 113  //col:186
    SystemCpuQuotaInformation  = 114  //col:187
    SystemNativeBasicInformation  = 115  //col:188
    SystemErrorPortTimeouts  = 116  //col:189
    SystemLowPriorityIoInformation  = 117  //col:190
    SystemTpmBootEntropyInformation  = 118  //col:191
    SystemVerifierCountersInformation  = 119  //col:192
    SystemPagedPoolInformationEx  = 120  //col:193
    SystemSystemPtesInformationEx  = 121  //col:194
    SystemNodeDistanceInformation  = 122  //col:195
    SystemAcpiAuditInformation  = 123  //col:196
    SystemBasicPerformanceInformation  = 124  //col:197
    SystemQueryPerformanceCounterInformation  = 125  //col:198
    SystemSessionBigPoolInformation  = 126  //col:199
    SystemBootGraphicsInformation  = 127  //col:200
    SystemScrubPhysicalMemoryInformation  = 128  //col:201
    SystemBadPageInformation = 129  //col:202
    SystemProcessorProfileControlArea  = 130  //col:203
    SystemCombinePhysicalMemoryInformation  = 131  //col:204
    SystemEntropyInterruptTimingInformation  = 132  //col:205
    SystemConsoleInformation  = 133  //col:206
    SystemPlatformBinaryInformation  = 134  //col:207
    SystemPolicyInformation  = 135  //col:208
    SystemHypervisorProcessorCountInformation  = 136  //col:209
    SystemDeviceDataInformation  = 137  //col:210
    SystemDeviceDataEnumerationInformation  = 138  //col:211
    SystemMemoryTopologyInformation  = 139  //col:212
    SystemMemoryChannelInformation  = 140  //col:213
    SystemBootLogoInformation  = 141  //col:214
    SystemProcessorPerformanceInformationEx  = 142  //col:215
    SystemCriticalProcessErrorLogInformation = 143  //col:216
    SystemSecureBootPolicyInformation  = 144  //col:217
    SystemPageFileInformationEx  = 145  //col:218
    SystemSecureBootInformation  = 146  //col:219
    SystemEntropyInterruptTimingRawInformation = 147  //col:220
    SystemPortableWorkspaceEfiLauncherInformation  = 148  //col:221
    SystemFullProcessInformation  = 149  //col:222
    SystemKernelDebuggerInformationEx  = 150  //col:223
    SystemBootMetadataInformation  = 151  //col:224
    SystemSoftRebootInformation  = 152  //col:225
    SystemElamCertificateInformation  = 153  //col:226
    SystemOfflineDumpConfigInformation  = 154  //col:227
    SystemProcessorFeaturesInformation  = 155  //col:228
    SystemRegistryReconciliationInformation  = 156  //col:229
    SystemEdidInformation  = 157  //col:230
    SystemManufacturingInformation  = 158  //col:231
    SystemEnergyEstimationConfigInformation  = 159  //col:232
    SystemHypervisorDetailInformation  = 160  //col:233
    SystemProcessorCycleStatsInformation  = 161  //col:234
    SystemVmGenerationCountInformation = 162  //col:235
    SystemTrustedPlatformModuleInformation  = 163  //col:236
    SystemKernelDebuggerFlags  = 164  //col:237
    SystemCodeIntegrityPolicyInformation  = 165  //col:238
    SystemIsolatedUserModeInformation  = 166  //col:239
    SystemHardwareSecurityTestInterfaceResultsInformation = 167  //col:240
    SystemSingleModuleInformation  = 168  //col:241
    SystemAllowedCpuSetsInformation = 169  //col:242
    SystemVsmProtectionInformation  = 170  //col:243
    SystemInterruptCpuSetsInformation  = 171  //col:244
    SystemSecureBootPolicyFullInformation  = 172  //col:245
    SystemCodeIntegrityPolicyFullInformation = 173  //col:246
    SystemAffinitizedInterruptProcessorInformation  = 174  //col:247
    SystemRootSiloInformation  = 175  //col:248
    SystemCpuSetInformation  = 176  //col:249
    SystemCpuSetTagInformation  = 177  //col:250
    SystemWin32WerStartCallout = 178  //col:251
    SystemSecureKernelProfileInformation  = 179  //col:252
    SystemCodeIntegrityPlatformManifestInformation  = 180  //col:253
    SystemInterruptSteeringInformation  = 181  //col:254
    SystemSupportedProcessorArchitectures  = 182  //col:255
    SystemMemoryUsageInformation  = 183  //col:256
    SystemCodeIntegrityCertificateInformation  = 184  //col:257
    SystemPhysicalMemoryInformation  = 185  //col:258
    SystemControlFlowTransition = 186  //col:259
    SystemKernelDebuggingAllowed  = 187  //col:260
    SystemActivityModerationExeState  = 188  //col:261
    SystemActivityModerationUserSettings  = 189  //col:262
    SystemCodeIntegrityPoliciesFullInformation = 190  //col:263
    SystemCodeIntegrityUnlockInformation  = 191  //col:264
    SystemIntegrityQuotaInformation = 192  //col:265
    SystemFlushInformation  = 193  //col:266
    SystemProcessorIdleMaskInformation  = 194  //col:267
    SystemSecureDumpEncryptionInformation = 195  //col:268
    SystemWriteConstraintInformation  = 196  //col:269
    SystemKernelVaShadowInformation  = 197  //col:270
    SystemHypervisorSharedPageInformation  = 198  //col:271
    SystemFirmwareBootPerformanceInformation = 199  //col:272
    SystemCodeIntegrityVerificationInformation  = 200  //col:273
    SystemFirmwarePartitionInformation  = 201  //col:274
    SystemSpeculationControlInformation  = 202  //col:275
    SystemDmaGuardPolicyInformation  = 203  //col:276
    SystemEnclaveLaunchControlInformation  = 204  //col:277
    SystemWorkloadAllowedCpuSetsInformation  = 205  //col:278
    SystemCodeIntegrityUnlockModeInformation = 206  //col:279
    SystemLeapSecondInformation  = 207  //col:280
    SystemFlags2Information  = 208  //col:281
    SystemSecurityModelInformation  = 209  //col:282
    SystemCodeIntegritySyntheticCacheInformation = 210  //col:283
    SystemFeatureConfigurationInformation  = 211  //col:284
    SystemFeatureConfigurationSectionInformation  = 212  //col:285
    SystemFeatureUsageSubscriptionInformation  = 213  //col:286
    SystemSecureSpeculationControlInformation  = 214  //col:287
    SystemSpacesBootInformation  = 215  //col:288
    SystemFwRamdiskInformation  = 216  //col:289
    SystemWheaIpmiHardwareInformation = 217  //col:290
    SystemDifSetRuleClassInformation = 218  //col:291
    SystemDifClearRuleClassInformation = 219  //col:292
    SystemDifApplyPluginVerificationOnDriver = 220  //col:293
    SystemDifRemovePluginVerificationOnDriver  = 221  //col:294
    SystemShadowStackInformation  = 222  //col:295
    SystemBuildVersionInformation  = 223  //col:296
    SystemPoolLimitInformation  = 224  //col:297
    SystemCodeIntegrityAddDynamicStore = 225  //col:298
    SystemCodeIntegrityClearDynamicStores = 226  //col:299
    SystemDifPoolTrackingInformation = 227  //col:300
    SystemPoolZeroingInformation  = 228  //col:301
    SystemDpcWatchdogInformation = 229  //col:302
    SystemDpcWatchdogInformation2 = 230  //col:303
    SystemSupportedProcessorArchitectures2  = 231  //col:304
    SystemSingleProcessorRelationshipInformation  = 232  //col:305
    SystemXfgCheckFailureInformation = 233  //col:306
    SystemIommuStateInformation  = 234  //col:307
    SystemHypervisorMinrootInformation  = 235  //col:308
    SystemHypervisorBootPagesInformation  = 236  //col:309
    SystemPointerAuthInformation  = 237  //col:310
    SystemSecureKernelDebuggerInformation = 238  //col:311
    SystemOriginalImageFeatureInformation = 239  //col:312
    MaxSystemInfoClass = 240  //col:313
)


const(
    EventTraceKernelVersionInformation  = 1  //col:317
    EventTraceGroupMaskInformation  = 2  //col:318
    EventTracePerformanceInformation  = 3  //col:319
    EventTraceTimeProfileInformation  = 4  //col:320
    EventTraceSessionSecurityInformation  = 5  //col:321
    EventTraceSpinlockInformation  = 6  //col:322
    EventTraceStackTracingInformation  = 7  //col:323
    EventTraceExecutiveResourceInformation  = 8  //col:324
    EventTraceHeapTracingInformation  = 9  //col:325
    EventTraceHeapSummaryTracingInformation  = 10  //col:326
    EventTracePoolTagFilterInformation  = 11  //col:327
    EventTracePebsTracingInformation  = 12  //col:328
    EventTraceProfileConfigInformation  = 13  //col:329
    EventTraceProfileSourceListInformation  = 14  //col:330
    EventTraceProfileEventListInformation  = 15  //col:331
    EventTraceProfileCounterListInformation  = 16  //col:332
    EventTraceStackCachingInformation  = 17  //col:333
    EventTraceObjectTypeFilterInformation  = 18  //col:334
    EventTraceSoftRestartInformation  = 19  //col:335
    EventTraceLastBranchConfigurationInformation  = 20  //col:336
    EventTraceLastBranchEventListInformation = 21  //col:337
    EventTraceProfileSourceAddInformation  = 22  //col:338
    EventTraceProfileSourceRemoveInformation  = 23  //col:339
    EventTraceProcessorTraceConfigurationInformation = 24  //col:340
    EventTraceProcessorTraceEventListInformation = 25  //col:341
    EventTraceCoverageSamplerInformation  = 26  //col:342
    EventTraceUnifiedStackCachingInformation  = 27  //col:343
    MaxEventTraceInfoClass = 28  //col:344
)


const(
    SystemCrashDumpDisable = 1  //col:348
    SystemCrashDumpReconfigure = 2  //col:349
    SystemCrashDumpInitializationComplete = 3  //col:350
)


const(
    WdActionSetTimeoutValue = 1  //col:354
    WdActionQueryTimeoutValue = 2  //col:355
    WdActionResetTimer = 3  //col:356
    WdActionStopTimer = 4  //col:357
    WdActionStartTimer = 5  //col:358
    WdActionSetTriggerAction = 6  //col:359
    WdActionQueryTriggerAction = 7  //col:360
    WdActionQueryState = 8  //col:361
)


const(
    WdInfoTimeoutValue  =  0  //col:365
    WdInfoResetTimer  =  1  //col:366
    WdInfoStopTimer  =  2  //col:367
    WdInfoStartTimer  =  3  //col:368
    WdInfoTriggerAction  =  4  //col:369
    WdInfoState  =  5  //col:370
    WdInfoTriggerReset  =  6  //col:371
    WdInfoNop  =  7  //col:372
    WdInfoGeneratedLastReset  =  8  //col:373
    WdInfoInvalid  =  9  //col:374
)


const(
    SystemFirmwareTableEnumerate = 1  //col:378
    SystemFirmwareTableGet = 2  //col:379
    SystemFirmwareTableMax = 3  //col:380
)


const(
    MemoryCaptureAccessedBits = 1  //col:384
    MemoryCaptureAndResetAccessedBits = 2  //col:385
    MemoryEmptyWorkingSets = 3  //col:386
    MemoryFlushModifiedList = 4  //col:387
    MemoryPurgeStandbyList = 5  //col:388
    MemoryPurgeLowPriorityStandbyList = 6  //col:389
    MemoryCommandMax = 7  //col:390
)


const(
    CoverageAllModules  =  0  //col:394
    CoverageSearchByHash  =  1  //col:395
    CoverageSearchByName  =  2  //col:396
)


const(
    SystemVaTypeAll = 1  //col:400
    SystemVaTypeNonPagedPool = 2  //col:401
    SystemVaTypePagedPool = 3  //col:402
    SystemVaTypeSystemCache = 4  //col:403
    SystemVaTypeSystemPtes = 5  //col:404
    SystemVaTypeSessionSpace = 6  //col:405
    SystemVaTypeMax = 7  //col:406
)


const(
    StorePageRequest  =  1  //col:410
    StoreStatsRequest  =  2   //col:411
    StoreCreateRequest  =  3   //col:412
    StoreDeleteRequest  =  4   //col:413
    StoreListRequest  =  5   //col:414
    Available1  =  6  //col:415
    StoreEmptyRequest  =  7  //col:416
    CacheListRequest  =  8   //col:417
    CacheCreateRequest  =  9   //col:418
    CacheDeleteRequest  =  10   //col:419
    CacheStoreCreateRequest  =  11   //col:420
    CacheStoreDeleteRequest  =  12   //col:421
    CacheStatsRequest  =  13   //col:422
    Available2  =  14  //col:423
    RegistrationRequest  =  15   //col:424
    GlobalCacheStatsRequest  =  16  //col:425
    StoreResizeRequest  =  17   //col:426
    CacheStoreResizeRequest  =  18   //col:427
    SmConfigRequest  =  19   //col:428
    StoreHighMemoryPriorityRequest  =  20   //col:429
    SystemStoreTrimRequest  =  21   //col:430
    MemCompressionInfoRequest  =  22    //col:431
    ProcessStoreInfoRequest  =  23   //col:432
    StoreInformationMax = 24  //col:433
)


const(
    StStatsLevelBasic  =  0  //col:437
    StStatsLevelIoStats  =  1  //col:438
    StStatsLevelRegionSpace  =  2   //col:439
    StStatsLevelSpaceBitmap  =  3   //col:440
    StStatsLevelMax  =  4  //col:441
)


const(
    StoreTypeInMemory = 0  //col:445
    StoreTypeFile = 1  //col:446
    StoreTypeMax = 2  //col:447
)


const(
    SmStoreManagerTypePhysical = 0  //col:451
    SmStoreManagerTypeVirtual = 1  //col:452
    SmStoreManagerTypeMax = 2  //col:453
)


const(
    SmConfigDirtyPageCompression  =  0  //col:457
    SmConfigAsyncInswap  =  1  //col:458
    SmConfigPrefetchSeekThreshold  =  2  //col:459
    SmConfigTypeMax  =  3  //col:460
)


const(
    TpmBootEntropyStructureUninitialized = 1  //col:464
    TpmBootEntropyDisabledByPolicy = 2  //col:465
    TpmBootEntropyNoTpmFound = 3  //col:466
    TpmBootEntropyTpmError = 4  //col:467
    TpmBootEntropySuccess = 5  //col:468
)


const(
    SystemPixelFormatUnknown = 1  //col:472
    SystemPixelFormatR8G8B8 = 2  //col:473
    SystemPixelFormatR8G8B8X8 = 3  //col:474
    SystemPixelFormatB8G8R8 = 4  //col:475
    SystemPixelFormatB8G8R8X8 = 5  //col:476
)


const(
    SystemProcessClassificationNormal = 1  //col:480
    SystemProcessClassificationSystem = 2  //col:481
    SystemProcessClassificationSecureSystem = 3  //col:482
    SystemProcessClassificationMemCompression = 4  //col:483
    SystemProcessClassificationRegistry  = 5  //col:484
    SystemProcessClassificationMaximum = 6  //col:485
)


const(
    SystemActivityModerationStateSystemManaged = 1  //col:489
    SystemActivityModerationStateUserManagedAllowThrottling = 2  //col:490
    SystemActivityModerationStateUserManagedDisableThrottling = 3  //col:491
    MaxSystemActivityModerationState = 4  //col:492
)


const(
    SystemActivityModerationAppTypeClassic = 1  //col:496
    SystemActivityModerationAppTypePackaged = 2  //col:497
    MaxSystemActivityModerationAppType = 3  //col:498
)


const(
    IommuStateBlock = 1  //col:502
    IommuStateUnblock = 2  //col:503
)


const(
    SysDbgQueryModuleInformation = 1  //col:507
    SysDbgQueryTraceInformation = 2  //col:508
    SysDbgSetTracepoint = 3  //col:509
    SysDbgSetSpecialCall  = 4  //col:510
    SysDbgClearSpecialCalls  = 5  //col:511
    SysDbgQuerySpecialCalls = 6  //col:512
    SysDbgBreakPoint = 7  //col:513
    SysDbgQueryVersion  = 8  //col:514
    SysDbgReadVirtual  = 9  //col:515
    SysDbgWriteVirtual  = 10  //col:516
    SysDbgReadPhysical  = 11  //col:517
    SysDbgWritePhysical  = 12  //col:518
    SysDbgReadControlSpace  = 13  //col:519
    SysDbgWriteControlSpace  = 14  //col:520
    SysDbgReadIoSpace  = 15  //col:521
    SysDbgWriteIoSpace  = 16  //col:522
    SysDbgReadMsr  = 17  //col:523
    SysDbgWriteMsr  = 18  //col:524
    SysDbgReadBusData  = 19  //col:525
    SysDbgWriteBusData  = 20  //col:526
    SysDbgCheckLowMemory  = 21  //col:527
    SysDbgEnableKernelDebugger = 22  //col:528
    SysDbgDisableKernelDebugger = 23  //col:529
    SysDbgGetAutoKdEnable = 24  //col:530
    SysDbgSetAutoKdEnable = 25  //col:531
    SysDbgGetPrintBufferSize = 26  //col:532
    SysDbgSetPrintBufferSize = 27  //col:533
    SysDbgGetKdUmExceptionEnable = 28  //col:534
    SysDbgSetKdUmExceptionEnable = 29  //col:535
    SysDbgGetTriageDump  = 30  //col:536
    SysDbgGetKdBlockEnable  = 31  //col:537
    SysDbgSetKdBlockEnable = 32  //col:538
    SysDbgRegisterForUmBreakInfo = 33  //col:539
    SysDbgGetUmBreakPid = 34  //col:540
    SysDbgClearUmBreakPid = 35  //col:541
    SysDbgGetUmAttachPid = 36  //col:542
    SysDbgClearUmAttachPid = 37  //col:543
    SysDbgGetLiveKernelDump  = 38  //col:544
    SysDbgKdPullRemoteFile  = 39  //col:545
    SysDbgMaxInfoClass = 40  //col:546
)


const(
    OptionAbortRetryIgnore = 1  //col:550
    OptionOk = 2  //col:551
    OptionOkCancel = 3  //col:552
    OptionRetryCancel = 4  //col:553
    OptionYesNo = 5  //col:554
    OptionYesNoCancel = 6  //col:555
    OptionShutdownSystem = 7  //col:556
    OptionOkNoWait = 8  //col:557
    OptionCancelTryContinue = 9  //col:558
)


const(
    ResponseReturnToCaller = 1  //col:562
    ResponseNotHandled = 2  //col:563
    ResponseAbort = 3  //col:564
    ResponseCancel = 4  //col:565
    ResponseIgnore = 5  //col:566
    ResponseNo = 6  //col:567
    ResponseOk = 7  //col:568
    ResponseRetry = 8  //col:569
    ResponseYes = 9  //col:570
    ResponseTryAgain = 10  //col:571
    ResponseContinue = 11  //col:572
)


const(
    StandardDesign = 1  //col:576
    NEC98x86 = 2  //col:577
    EndAlternatives = 3  //col:578
)


const(
    AtomBasicInformation = 1  //col:582
    AtomTableInformation = 2  //col:583
)


const(
    ShutdownNoReboot = 1  //col:587
    ShutdownReboot = 2  //col:588
    ShutdownPowerOff = 3  //col:589
    ShutdownRebootForRecovery  = 4  //col:590
)



type BOOT_ENTRY struct{
Version ULONG
Length ULONG
Id ULONG
Attributes ULONG
FriendlyNameOffset ULONG
BootFilePathOffset ULONG
OsOptionsLength ULONG
OsOptions[1] UCHAR
}


type BOOT_ENTRY_LIST struct{
NextEntryOffset ULONG
BootEntry BOOT_ENTRY
}


type BOOT_OPTIONS struct{
Version ULONG
Length ULONG
Timeout ULONG
CurrentBootEntryId ULONG
NextBootEntryId ULONG
HeadlessRedirection[1] WCHAR
}


type FILE_PATH struct{
Version ULONG
Length ULONG
Type ULONG
FilePath[1] UCHAR
}


type EFI_DRIVER_ENTRY struct{
Version ULONG
Length ULONG
Id ULONG
FriendlyNameOffset ULONG
DriverFilePathOffset ULONG
}


type EFI_DRIVER_ENTRY_LIST struct{
NextEntryOffset ULONG
DriverEntry EFI_DRIVER_ENTRY
}


type EVENT_BASIC_INFORMATION struct{
EventType EVENT_TYPE
EventState LONG
}


type MUTANT_BASIC_INFORMATION struct{
CurrentCount LONG
OwnedByCaller bool
AbandonedState bool
}


type MUTANT_OWNER_INFORMATION struct{
ClientId CLIENT_ID
}


type SEMAPHORE_BASIC_INFORMATION struct{
CurrentCount LONG
MaximumCount LONG
}


type TIMER_BASIC_INFORMATION struct{
RemainingTime LARGE_INTEGER
TimerState bool
}


type TIMER_SET_COALESCABLE_TIMER_INFO struct{
LARGE_INTEGER _In_
PTIMER_APC_ROUTINE _In_opt_
PVOID _In_opt_
struct _In_opt_
ULONG _In_opt_
ULONG _In_
PBOOLEAN _Out_opt_
}


type T2_SET_PARAMETERS_V0 struct{
Version ULONG
Reserved ULONG
NoWakeTolerance LONGLONG
}


type WNF_STATE_NAME struct{
Data[2] ULONG
}


type WNF_TYPE_ID struct{
TypeId GUID
}


type WNF_DELIVERY_DESCRIPTOR struct{
SubscriptionId ULONGLONG
StateName WNF_STATE_NAME
ChangeStamp WNF_CHANGE_STAMP
StateDataSize ULONG
EventMask ULONG
TypeId WNF_TYPE_ID
StateDataOffset ULONG
}


type WORKER_FACTORY_BASIC_INFORMATION struct{
Timeout LARGE_INTEGER
RetryTimeout LARGE_INTEGER
IdleTimeout LARGE_INTEGER
Paused bool
TimerSet bool
QueuedToExWorker bool
MayCreate bool
CreateInProgress bool
InsertedIntoQueue bool
Shutdown bool
BindingCount ULONG
ThreadMinimum ULONG
ThreadMaximum ULONG
PendingWorkerCount ULONG
WaitingWorkerCount ULONG
TotalWorkerCount ULONG
ReleaseCount ULONG
InfiniteWaitGoal LONGLONG
StartRoutine PVOID
StartParameter PVOID
ProcessId HANDLE
StackReserve SIZE_T
StackCommit SIZE_T
LastThreadCreationStatus NTSTATUS
}


type WORKER_FACTORY_DEFERRED_WORK struct{
_PORT_MESSAGE struct
AlpcSendMessagePort PVOID
AlpcSendMessageFlags ULONG
Flags ULONG
}


type SYSTEM_BASIC_INFORMATION struct{
Reserved ULONG
TimerResolution ULONG
PageSize ULONG
NumberOfPhysicalPages ULONG
LowestPhysicalPageNumber ULONG
HighestPhysicalPageNumber ULONG
AllocationGranularity ULONG
MinimumUserModeAddress ULONG_PTR
MaximumUserModeAddress ULONG_PTR
ActiveProcessorsAffinityMask KAFFINITY
NumberOfProcessors CCHAR
}


type SYSTEM_PROCESSOR_INFORMATION struct{
ProcessorArchitecture USHORT
ProcessorLevel USHORT
ProcessorRevision USHORT
MaximumProcessors USHORT
ProcessorFeatureBits ULONG
}


type SYSTEM_PERFORMANCE_INFORMATION struct{
IdleProcessTime LARGE_INTEGER
IoReadTransferCount LARGE_INTEGER
IoWriteTransferCount LARGE_INTEGER
IoOtherTransferCount LARGE_INTEGER
IoReadOperationCount ULONG
IoWriteOperationCount ULONG
IoOtherOperationCount ULONG
AvailablePages ULONG
CommittedPages ULONG
CommitLimit ULONG
PeakCommitment ULONG
PageFaultCount ULONG
CopyOnWriteCount ULONG
TransitionCount ULONG
CacheTransitionCount ULONG
DemandZeroCount ULONG
PageReadCount ULONG
PageReadIoCount ULONG
CacheReadCount ULONG
CacheIoCount ULONG
DirtyPagesWriteCount ULONG
DirtyWriteIoCount ULONG
MappedPagesWriteCount ULONG
MappedWriteIoCount ULONG
PagedPoolPages ULONG
NonPagedPoolPages ULONG
PagedPoolAllocs ULONG
PagedPoolFrees ULONG
NonPagedPoolAllocs ULONG
NonPagedPoolFrees ULONG
FreeSystemPtes ULONG
ResidentSystemCodePage ULONG
TotalSystemDriverPages ULONG
TotalSystemCodePages ULONG
NonPagedPoolLookasideHits ULONG
PagedPoolLookasideHits ULONG
AvailablePagedPoolPages ULONG
ResidentSystemCachePage ULONG
ResidentPagedPoolPage ULONG
ResidentSystemDriverPage ULONG
CcFastReadNoWait ULONG
CcFastReadWait ULONG
CcFastReadResourceMiss ULONG
CcFastReadNotPossible ULONG
CcFastMdlReadNoWait ULONG
CcFastMdlReadWait ULONG
CcFastMdlReadResourceMiss ULONG
CcFastMdlReadNotPossible ULONG
CcMapDataNoWait ULONG
CcMapDataWait ULONG
CcMapDataNoWaitMiss ULONG
CcMapDataWaitMiss ULONG
CcPinMappedDataCount ULONG
CcPinReadNoWait ULONG
CcPinReadWait ULONG
CcPinReadNoWaitMiss ULONG
CcPinReadWaitMiss ULONG
CcCopyReadNoWait ULONG
CcCopyReadWait ULONG
CcCopyReadNoWaitMiss ULONG
CcCopyReadWaitMiss ULONG
CcMdlReadNoWait ULONG
CcMdlReadWait ULONG
CcMdlReadNoWaitMiss ULONG
CcMdlReadWaitMiss ULONG
CcReadAheadIos ULONG
CcLazyWriteIos ULONG
CcLazyWritePages ULONG
CcDataFlushes ULONG
CcDataPages ULONG
ContextSwitches ULONG
FirstLevelTbFills ULONG
SecondLevelTbFills ULONG
SystemCalls ULONG
CcTotalDirtyPages ULONGLONG
CcDirtyPageThreshold ULONGLONG
ResidentAvailablePages LONGLONG
SharedCommittedPages ULONGLONG
}


type SYSTEM_TIMEOFDAY_INFORMATION struct{
BootTime LARGE_INTEGER
CurrentTime LARGE_INTEGER
TimeZoneBias LARGE_INTEGER
TimeZoneId ULONG
Reserved ULONG
BootTimeBias ULONGLONG
SleepTimeBias ULONGLONG
}


type SYSTEM_THREAD_INFORMATION struct{
KernelTime LARGE_INTEGER
UserTime LARGE_INTEGER
CreateTime LARGE_INTEGER
WaitTime ULONG
StartAddress PVOID
ClientId CLIENT_ID
Priority KPRIORITY
BasePriority KPRIORITY
ContextSwitches ULONG
ThreadState KTHREAD_STATE
WaitReason KWAIT_REASON
}


type SYSTEM_EXTENDED_THREAD_INFORMATION struct{
ThreadInfo SYSTEM_THREAD_INFORMATION
StackBase PVOID
StackLimit PVOID
Win32StartAddress PVOID
TebBase PTEB
Reserved2 ULONG_PTR
Reserved3 ULONG_PTR
Reserved4 ULONG_PTR
}


type SYSTEM_EXTENDED_THREAD_INFORMATION struct{
ThreadInfo SYSTEM_THREAD_INFORMATION
StackBase PVOID
StackLimit PVOID
Win32StartAddress PVOID
TebBase PTEB
Reserved2 ULONG_PTR
Reserved3 ULONG_PTR
Reserved4 ULONG_PTR
}


type SYSTEM_PROCESS_INFORMATION struct{
NextEntryOffset ULONG
NumberOfThreads ULONG
WorkingSetPrivateSize LARGE_INTEGER
HardFaultCount ULONG
NumberOfThreadsHighWatermark ULONG
CycleTime ULONGLONG
CreateTime LARGE_INTEGER
UserTime LARGE_INTEGER
KernelTime LARGE_INTEGER
ImageName UNICODE_STRING
BasePriority KPRIORITY
UniqueProcessId HANDLE
InheritedFromUniqueProcessId HANDLE
HandleCount ULONG
SessionId ULONG
UniqueProcessKey ULONG_PTR
PeakVirtualSize SIZE_T
VirtualSize SIZE_T
PageFaultCount ULONG
PeakWorkingSetSize SIZE_T
WorkingSetSize SIZE_T
QuotaPeakPagedPoolUsage SIZE_T
QuotaPagedPoolUsage SIZE_T
QuotaPeakNonPagedPoolUsage SIZE_T
QuotaNonPagedPoolUsage SIZE_T
PagefileUsage SIZE_T
PeakPagefileUsage SIZE_T
PrivatePageCount SIZE_T
ReadOperationCount LARGE_INTEGER
WriteOperationCount LARGE_INTEGER
OtherOperationCount LARGE_INTEGER
ReadTransferCount LARGE_INTEGER
WriteTransferCount LARGE_INTEGER
OtherTransferCount LARGE_INTEGER
Threads[1] SYSTEM_THREAD_INFORMATION
}


type SYSTEM_CALL_COUNT_INFORMATION struct{
Length ULONG
NumberOfTables ULONG
}


type SYSTEM_DEVICE_INFORMATION struct{
NumberOfDisks ULONG
NumberOfFloppies ULONG
NumberOfCdRoms ULONG
NumberOfTapes ULONG
NumberOfSerialPorts ULONG
NumberOfParallelPorts ULONG
}


type SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION struct{
IdleTime LARGE_INTEGER
KernelTime LARGE_INTEGER
UserTime LARGE_INTEGER
DpcTime LARGE_INTEGER
InterruptTime LARGE_INTEGER
InterruptCount ULONG
}


type SYSTEM_FLAGS_INFORMATION struct{
Flags ULONG
}


type SYSTEM_CALL_TIME_INFORMATION struct{
Length ULONG
TotalCalls ULONG
TimeOfCalls[1] LARGE_INTEGER
}


type RTL_PROCESS_LOCK_INFORMATION struct{
Address PVOID
Type USHORT
CreatorBackTraceIndex USHORT
OwningThread HANDLE
LockCount LONG
ContentionCount ULONG
EntryCount ULONG
RecursionCount LONG
NumberOfWaitingShared ULONG
NumberOfWaitingExclusive ULONG
}


type RTL_PROCESS_LOCKS struct{
NumberOfLocks ULONG
Locks[1] RTL_PROCESS_LOCK_INFORMATION
}


type RTL_PROCESS_BACKTRACE_INFORMATION struct{
SymbolicBackTrace PCHAR
TraceCount ULONG
Index USHORT
Depth USHORT
BackTrace[32] PVOID
}


type RTL_PROCESS_BACKTRACES struct{
CommittedMemory ULONG
ReservedMemory ULONG
NumberOfBackTraceLookups ULONG
NumberOfBackTraces ULONG
BackTraces[1] RTL_PROCESS_BACKTRACE_INFORMATION
}


type SYSTEM_HANDLE_TABLE_ENTRY_INFO struct{
UniqueProcessId USHORT
CreatorBackTraceIndex USHORT
ObjectTypeIndex UCHAR
HandleAttributes UCHAR
HandleValue USHORT
Object PVOID
GrantedAccess ULONG
}


type SYSTEM_HANDLE_INFORMATION struct{
NumberOfHandles ULONG
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO
}


type SYSTEM_OBJECTTYPE_INFORMATION struct{
NextEntryOffset ULONG
NumberOfObjects ULONG
NumberOfHandles ULONG
TypeIndex ULONG
InvalidAttributes ULONG
GenericMapping GENERIC_MAPPING
ValidAccessMask ULONG
PoolType ULONG
SecurityRequired bool
WaitableObject bool
TypeName UNICODE_STRING
}


type SYSTEM_OBJECT_INFORMATION struct{
NextEntryOffset ULONG
Object PVOID
CreatorUniqueProcess HANDLE
CreatorBackTraceIndex USHORT
Flags USHORT
PointerCount LONG
HandleCount LONG
PagedPoolCharge ULONG
NonPagedPoolCharge ULONG
ExclusiveProcessId HANDLE
SecurityDescriptor PVOID
NameInfo UNICODE_STRING
}


type SYSTEM_PAGEFILE_INFORMATION struct{
NextEntryOffset ULONG
TotalSize ULONG
TotalInUse ULONG
PeakUsage ULONG
PageFileName UNICODE_STRING
}


type SYSTEM_VDM_INSTEMUL_INFO struct{
SegmentNotPresent ULONG
VdmOpcode0F ULONG
OpcodeESPrefix ULONG
OpcodeCSPrefix ULONG
OpcodeSSPrefix ULONG
OpcodeDSPrefix ULONG
OpcodeFSPrefix ULONG
OpcodeGSPrefix ULONG
OpcodeOPER32Prefix ULONG
OpcodeADDR32Prefix ULONG
OpcodeINSB ULONG
OpcodeINSW ULONG
OpcodeOUTSB ULONG
OpcodeOUTSW ULONG
OpcodePUSHF ULONG
OpcodePOPF ULONG
OpcodeINTnn ULONG
OpcodeINTO ULONG
OpcodeIRET ULONG
OpcodeINBimm ULONG
OpcodeINWimm ULONG
OpcodeOUTBimm ULONG
OpcodeOUTWimm ULONG
OpcodeINB ULONG
OpcodeINW ULONG
OpcodeOUTB ULONG
OpcodeOUTW ULONG
OpcodeLOCKPrefix ULONG
OpcodeREPNEPrefix ULONG
OpcodeREPPrefix ULONG
OpcodeHLT ULONG
OpcodeCLI ULONG
OpcodeSTI ULONG
BopCount ULONG
}


type SYSTEM_FILECACHE_INFORMATION struct{
CurrentSize SIZE_T
PeakSize SIZE_T
PageFaultCount ULONG
MinimumWorkingSet SIZE_T
MaximumWorkingSet SIZE_T
CurrentSizeIncludingTransitionInPages SIZE_T
PeakSizeIncludingTransitionInPages SIZE_T
TransitionRePurposeCount ULONG
Flags ULONG
}


type SYSTEM_BASIC_WORKING_SET_INFORMATION struct{
CurrentSize SIZE_T
PeakSize SIZE_T
PageFaultCount ULONG
}


type SYSTEM_POOLTAG struct{
Union union
Tag[4] UCHAR
TagUlong ULONG
}


type SYSTEM_POOLTAG_INFORMATION struct{
Count ULONG
TagInfo[1] SYSTEM_POOLTAG
}


type SYSTEM_INTERRUPT_INFORMATION struct{
ContextSwitches ULONG
DpcCount ULONG
DpcRate ULONG
TimeIncrement ULONG
DpcBypassCount ULONG
ApcBypassCount ULONG
}


type SYSTEM_DPC_BEHAVIOR_INFORMATION struct{
Spare ULONG
DpcQueueDepth ULONG
MinimumDpcRate ULONG
AdjustDpcThreshold ULONG
IdealDpcRate ULONG
}


type SYSTEM_QUERY_TIME_ADJUST_INFORMATION struct{
TimeAdjustment ULONG
TimeIncrement ULONG
Enable bool
}


type SYSTEM_QUERY_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG
TimeIncrement ULONGLONG
Enable bool
}


type SYSTEM_SET_TIME_ADJUST_INFORMATION struct{
TimeAdjustment ULONG
Enable bool
}


type SYSTEM_SET_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG
Enable bool
}


type EVENT_TRACE_VERSION_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
EventTraceKernelVersion ULONG
}


type PERFINFO_GROUPMASK struct{
Masks[PERF_NUM_MASKS] ULONG
}


type EVENT_TRACE_GROUPMASK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
EventTraceGroupMasks PERFINFO_GROUPMASK
}


type EVENT_TRACE_PERFORMANCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
LogfileBytesWritten LARGE_INTEGER
}


type EVENT_TRACE_TIME_PROFILE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
ProfileInterval ULONG
}


type EVENT_TRACE_SESSION_SECURITY_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
SecurityInformation ULONG
TraceHandle TRACEHANDLE
SecurityDescriptor[1] UCHAR
}


type EVENT_TRACE_SPINLOCK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
SpinLockSpinThreshold ULONG
SpinLockAcquireSampleRate ULONG
SpinLockContentionSampleRate ULONG
SpinLockHoldThreshold ULONG
}


type EVENT_TRACE_SYSTEM_EVENT_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
HookId[1] ULONG
}


type EVENT_TRACE_EXECUTIVE_RESOURCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
ReleaseSamplingRate ULONG
ContentionSamplingRate ULONG
NumberOfExcessiveTimeouts ULONG
}


type EVENT_TRACE_HEAP_TRACING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
ProcessId ULONG
}


type EVENT_TRACE_TAG_FILTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
Filter[1] ULONG
}


type EVENT_TRACE_PROFILE_COUNTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
ProfileSource[1] ULONG
}


type EVENT_TRACE_PROFILE_LIST_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
Spare ULONG
_PROFILE_SOURCE_INFO* struct
}


type EVENT_TRACE_STACK_CACHING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
Enabled bool
Reserved[3] UCHAR
CacheSize ULONG
BucketCount ULONG
}


type EVENT_TRACE_SOFT_RESTART_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
TraceHandle TRACEHANDLE
PersistTraceBuffers bool
FileName[1] WCHAR
}


type EVENT_TRACE_PROFILE_ADD_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
PerfEvtEventSelect bool
PerfEvtUnitSelect bool
PerfEvtType ULONG
CpuInfoHierarchy[0x3] ULONG
InitialInterval ULONG
AllowsHalt bool
Persist bool
ProfileSourceDescription[0x1] WCHAR
}


type EVENT_TRACE_PROFILE_REMOVE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
ProfileSource KPROFILE_SOURCE
CpuInfoHierarchy[0x3] ULONG
}


type EVENT_TRACE_COVERAGE_SAMPLER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS
CoverageSamplerInformationClass bool
MajorVersion UCHAR
MinorVersion UCHAR
Reserved UCHAR
SamplerHandle HANDLE
}


type SYSTEM_EXCEPTION_INFORMATION struct{
AlignmentFixupCount ULONG
ExceptionDispatchCount ULONG
FloatingEmulationCount ULONG
ByteWordEmulationCount ULONG
}


type SYSTEM_CRASH_DUMP_STATE_INFORMATION struct{
CrashDumpConfigurationClass SYSTEM_CRASH_DUMP_CONFIGURATION_CLASS
}


type SYSTEM_KERNEL_DEBUGGER_INFORMATION struct{
KernelDebuggerEnabled bool
KernelDebuggerNotPresent bool
}


type SYSTEM_CONTEXT_SWITCH_INFORMATION struct{
ContextSwitches ULONG
FindAny ULONG
FindLast ULONG
FindIdeal ULONG
IdleAny ULONG
IdleCurrent ULONG
IdleLast ULONG
IdleIdeal ULONG
PreemptAny ULONG
PreemptCurrent ULONG
PreemptLast ULONG
SwitchToIdle ULONG
}


type SYSTEM_REGISTRY_QUOTA_INFORMATION struct{
RegistryQuotaAllowed ULONG
RegistryQuotaUsed ULONG
PagedPoolSize SIZE_T
}


type SYSTEM_PROCESSOR_IDLE_INFORMATION struct{
IdleTime ULONGLONG
C1Time ULONGLONG
C2Time ULONGLONG
C3Time ULONGLONG
C1Transitions ULONG
C2Transitions ULONG
C3Transitions ULONG
Padding ULONG
}


type SYSTEM_LEGACY_DRIVER_INFORMATION struct{
VetoType ULONG
VetoList UNICODE_STRING
}


type SYSTEM_LOOKASIDE_INFORMATION struct{
CurrentDepth USHORT
MaximumDepth USHORT
TotalAllocates ULONG
AllocateMisses ULONG
TotalFrees ULONG
FreeMisses ULONG
Type ULONG
Tag ULONG
Size ULONG
}


type SYSTEM_RANGE_START_INFORMATION struct{
SystemRangeStart ULONG_PTR
}


type SYSTEM_VERIFIER_INFORMATION_LEGACY  struct{
NextEntryOffset ULONG
Level ULONG
DriverName UNICODE_STRING
RaiseIrqls ULONG
AcquireSpinLocks ULONG
SynchronizeExecutions ULONG
AllocationsAttempted ULONG
AllocationsSucceeded ULONG
AllocationsSucceededSpecialPool ULONG
AllocationsWithNoTag ULONG
TrimRequests ULONG
Trims ULONG
AllocationsFailed ULONG
AllocationsFailedDeliberately ULONG
Loads ULONG
Unloads ULONG
UnTrackedPool ULONG
CurrentPagedPoolAllocations ULONG
CurrentNonPagedPoolAllocations ULONG
PeakPagedPoolAllocations ULONG
PeakNonPagedPoolAllocations ULONG
PagedPoolUsageInBytes SIZE_T
NonPagedPoolUsageInBytes SIZE_T
PeakPagedPoolUsageInBytes SIZE_T
PeakNonPagedPoolUsageInBytes SIZE_T
}


type SYSTEM_VERIFIER_INFORMATION struct{
NextEntryOffset ULONG
Level ULONG
RuleClasses[2] ULONG
TriageContext ULONG
AreAllDriversBeingVerified ULONG
DriverName UNICODE_STRING
RaiseIrqls ULONG
AcquireSpinLocks ULONG
SynchronizeExecutions ULONG
AllocationsAttempted ULONG
AllocationsSucceeded ULONG
AllocationsSucceededSpecialPool ULONG
AllocationsWithNoTag ULONG
TrimRequests ULONG
Trims ULONG
AllocationsFailed ULONG
AllocationsFailedDeliberately ULONG
Loads ULONG
Unloads ULONG
UnTrackedPool ULONG
CurrentPagedPoolAllocations ULONG
CurrentNonPagedPoolAllocations ULONG
PeakPagedPoolAllocations ULONG
PeakNonPagedPoolAllocations ULONG
PagedPoolUsageInBytes SIZE_T
NonPagedPoolUsageInBytes SIZE_T
PeakPagedPoolUsageInBytes SIZE_T
PeakNonPagedPoolUsageInBytes SIZE_T
}


type SYSTEM_SESSION_PROCESS_INFORMATION struct{
SessionId ULONG
SizeOfBuf ULONG
Buffer PVOID
}


type SYSTEM_GDI_DRIVER_INFORMATION struct{
DriverName UNICODE_STRING
ImageAddress PVOID
SectionPointer PVOID
EntryPoint PVOID
_IMAGE_EXPORT_DIRECTORY* struct
ImageLength ULONG
}


type SYSTEM_NUMA_INFORMATION struct{
HighestNodeNumber ULONG
Reserved ULONG
Union union
ActiveProcessorsGroupAffinity[MAXIMUM_NODE_COUNT] GROUP_AFFINITY
AvailableMemory[MAXIMUM_NODE_COUNT] ULONGLONG
Pad[MAXIMUM_NODE_COUNT ULONGLONG
}


type SYSTEM_PROCESSOR_POWER_INFORMATION struct{
CurrentFrequency UCHAR
ThermalLimitFrequency UCHAR
ConstantThrottleFrequency UCHAR
DegradedThrottleFrequency UCHAR
LastBusyFrequency UCHAR
LastC3Frequency UCHAR
LastAdjustedBusyFrequency UCHAR
ProcessorMinThrottle UCHAR
ProcessorMaxThrottle UCHAR
NumberOfFrequencies ULONG
PromotionCount ULONG
DemotionCount ULONG
ErrorCount ULONG
RetryCount ULONG
CurrentFrequencyTime ULONGLONG
CurrentProcessorTime ULONGLONG
CurrentProcessorIdleTime ULONGLONG
LastProcessorTime ULONGLONG
LastProcessorIdleTime ULONGLONG
Energy ULONGLONG
}


type SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX struct{
Object PVOID
UniqueProcessId ULONG_PTR
HandleValue ULONG_PTR
GrantedAccess ULONG
CreatorBackTraceIndex USHORT
ObjectTypeIndex USHORT
HandleAttributes ULONG
Reserved ULONG
}


type SYSTEM_HANDLE_INFORMATION_EX struct{
NumberOfHandles ULONG_PTR
Reserved ULONG_PTR
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX
}


type SYSTEM_BIGPOOL_ENTRY struct{
Union union
VirtualAddress PVOID
NonPaged ULONG_PTR
}


type SYSTEM_BIGPOOL_INFORMATION struct{
Count ULONG
AllocatedInfo[1] SYSTEM_BIGPOOL_ENTRY
}


type SYSTEM_POOL_ENTRY struct{
Allocated bool
Spare0 bool
AllocatorBackTraceIndex USHORT
Size ULONG
Union union
Tag[4] UCHAR
TagUlong ULONG
ProcessChargedQuota PVOID
}


type SYSTEM_POOL_INFORMATION struct{
TotalSize SIZE_T
FirstEntry PVOID
EntryOverhead USHORT
PoolTagPresent bool
Spare0 bool
NumberOfEntries ULONG
Entries[1] SYSTEM_POOL_ENTRY
}


type SYSTEM_SESSION_POOLTAG_INFORMATION struct{
NextEntryOffset SIZE_T
SessionId ULONG
Count ULONG
TagInfo[1] SYSTEM_POOLTAG
}


type SYSTEM_SESSION_MAPPED_VIEW_INFORMATION struct{
NextEntryOffset SIZE_T
SessionId ULONG
ViewFailures ULONG
NumberOfBytesAvailable SIZE_T
NumberOfBytesAvailableContiguous SIZE_T
}


type SYSTEM_WATCHDOG_HANDLER_INFORMATION  struct{
WdHandler PSYSTEM_WATCHDOG_HANDLER
Context PVOID
}


type SYSTEM_WATCHDOG_TIMER_INFORMATION struct{
WdInfoClass WATCHDOG_INFORMATION_CLASS
DataValue ULONG
}


type SYSTEM_FIRMWARE_TABLE_INFORMATION struct{
ProviderSignature ULONG
Action SYSTEM_FIRMWARE_TABLE_ACTION
TableID ULONG
TableBufferLength ULONG
TableBuffer[1] UCHAR
}


type SYSTEM_FIRMWARE_TABLE_HANDLER struct{
ProviderSignature ULONG
Register bool
FirmwareTableHandler PFNFTH
DriverObject PVOID
}


type SYSTEM_MEMORY_LIST_INFORMATION struct{
ZeroPageCount ULONG_PTR
FreePageCount ULONG_PTR
ModifiedPageCount ULONG_PTR
ModifiedNoWritePageCount ULONG_PTR
BadPageCount ULONG_PTR
PageCountByPriority[8] ULONG_PTR
RepurposedPagesByPriority[8] ULONG_PTR
ModifiedPageCountPageFile ULONG_PTR
}


type SYSTEM_THREAD_CID_PRIORITY_INFORMATION struct{
ClientId CLIENT_ID
Priority KPRIORITY
}


type SYSTEM_PROCESSOR_IDLE_CYCLE_TIME_INFORMATION struct{
CycleTime ULONGLONG
}


type SYSTEM_VERIFIER_ISSUE struct{
IssueType ULONGLONG
Address PVOID
Parameters[2] ULONGLONG
}


type SYSTEM_VERIFIER_CANCELLATION_INFORMATION struct{
CancelProbability ULONG
CancelThreshold ULONG
CompletionThreshold ULONG
CancellationVerifierDisabled ULONG
AvailableIssues ULONG
Issues[128] SYSTEM_VERIFIER_ISSUE
}


type SYSTEM_REF_TRACE_INFORMATION struct{
TraceEnable bool
TracePermanent bool
TraceProcessName UNICODE_STRING
TracePoolTags UNICODE_STRING
}


type SYSTEM_SPECIAL_POOL_INFORMATION struct{
PoolTag ULONG
Flags ULONG
}


type SYSTEM_PROCESS_ID_INFORMATION struct{
ProcessId HANDLE
ImageName UNICODE_STRING
}


type SYSTEM_HYPERVISOR_QUERY_INFORMATION struct{
HypervisorConnected bool
HypervisorDebuggingEnabled bool
HypervisorPresent bool
Spare0[5] bool
EnabledEnlightenments ULONGLONG
}


type SYSTEM_BOOT_ENVIRONMENT_INFORMATION struct{
BootIdentifier GUID
FirmwareType FIRMWARE_TYPE
Union union
BootFlags ULONGLONG
Struct struct
DbgMenuOsSelection ULONGLONG
DbgHiberBoot ULONGLONG
DbgSoftBoot ULONGLONG
DbgMeasuredLaunch ULONGLONG
DbgMeasuredLaunchCapable ULONGLONG
DbgSystemHiveReplace ULONGLONG
DbgMeasuredLaunchSmmProtections ULONGLONG
DbgMeasuredLaunchSmmLevel ULONGLONG
}


type SYSTEM_IMAGE_FILE_EXECUTION_OPTIONS_INFORMATION struct{
FlagsToEnable ULONG
FlagsToDisable ULONG
}


type COVERAGE_MODULE_REQUEST struct{
RequestType COVERAGE_REQUEST_CODES
Union union
MD5Hash[16] UCHAR
ModuleName UNICODE_STRING
}


type COVERAGE_MODULE_INFO struct{
ModuleInfoSize ULONG
IsBinaryLoaded ULONG
ModulePathName UNICODE_STRING
CoverageSectionSize ULONG
CoverageSection[1] UCHAR
}


type COVERAGE_MODULES struct{
ListAndReset ULONG
NumberOfModules ULONG
ModuleRequestInfo COVERAGE_MODULE_REQUEST
Modules[1] COVERAGE_MODULE_INFO
}


type SYSTEM_PREFETCH_PATCH_INFORMATION struct{
PrefetchPatchCount ULONG
}


type SYSTEM_VERIFIER_FAULTS_INFORMATION struct{
Probability ULONG
MaxProbability ULONG
PoolTags UNICODE_STRING
Applications UNICODE_STRING
}


type SYSTEM_VERIFIER_INFORMATION_EX struct{
VerifyMode ULONG
OptionChanges ULONG
PreviousBucketName UNICODE_STRING
IrpCancelTimeoutMsec ULONG
VerifierExtensionEnabled ULONG
#ifdefWin64 #ifdef _WIN64
Reserved[1] ULONG
#else #else
Reserved[3] ULONG
#endif #endif
}


type SYSTEM_SYSTEM_PARTITION_INFORMATION struct{
SystemPartition UNICODE_STRING
}


type SYSTEM_SYSTEM_DISK_INFORMATION struct{
SystemDisk UNICODE_STRING
}


type SYSTEM_NUMA_PROXIMITY_MAP struct{
NodeProximityId ULONG
NodeNumber USHORT
}


type SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT struct{
Hits ULONGLONG
PercentFrequency UCHAR
}


type SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT_WIN8 struct{
Hits ULONG
PercentFrequency UCHAR
}


type SYSTEM_PROCESSOR_PERFORMANCE_STATE_DISTRIBUTION struct{
ProcessorNumber ULONG
StateCount ULONG
States[1] SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT
}


type SYSTEM_PROCESSOR_PERFORMANCE_DISTRIBUTION struct{
ProcessorCount ULONG
Offsets[1] ULONG
}


type SYSTEM_CODEINTEGRITY_INFORMATION struct{
Length ULONG
CodeIntegrityOptions ULONG
}


type SYSTEM_PROCESSOR_MICROCODE_UPDATE_INFORMATION struct{
Operation ULONG
}


type SYSTEM_VA_LIST_INFORMATION struct{
VirtualSize SIZE_T
VirtualPeak SIZE_T
VirtualLimit SIZE_T
AllocationFailures SIZE_T
}


type SYSTEM_STORE_INFORMATION struct{
ULONG _In_
STORE_INFORMATION_CLASS _In_
PVOID _Inout_
ULONG _Inout_
}


type SM_STATS_REQUEST struct{
Version ULONG
DetailLevel ULONG
StoreId ULONG
BufferSize ULONG
Buffer PVOID
}


type ST_DATA_MGR_STATS struct{
RegionCount ULONG
PagesStored ULONG
UniquePagesStored ULONG
LazyCleanupRegionCount ULONG
RegionsInUse ULONG
SpaceUsed ULONG
}


type ST_IO_STATS_PERIOD struct{
PageCounts[5] ULONG
}


type ST_IO_STATS struct{
PeriodCount ULONG
Periods[64] ST_IO_STATS_PERIOD
}


type ST_READ_LATENCY_BUCKET struct{
LatencyUs ULONG
Count ULONG
}


type ST_READ_LATENCY_STATS struct{
Buckets[8] ST_READ_LATENCY_BUCKET
}


type ST_STATS_REGION_INFO struct{
SpaceUsed USHORT
Priority UCHAR
Spare UCHAR
}


type ST_STATS_SPACE_BITMAP struct{
CompressedBytes SIZE_T
BytesPerBit ULONG
StoreBitmap[1] UCHAR
}


type ST_STATS struct{
Version ULONG
Level ULONG
StoreType ULONG
NoDuplication ULONG
NoCompression ULONG
EncryptionStrength ULONG
VirtualRegions ULONG
Spare0 ULONG
Size ULONG
CompressionFormat USHORT
Spare USHORT
Struct struct
RegionSize ULONG
RegionCount ULONG
RegionCountMax ULONG
Granularity ULONG
UserData ST_DATA_MGR_STATS
Metadata ST_DATA_MGR_STATS
}


type SM_STORE_BASIC_PARAMS struct{
Union union
Struct struct
StoreType ULONG
NoDuplication ULONG
FailNoCompression ULONG
NoCompression ULONG
NoEncryption ULONG
NoEvictOnAdd ULONG
PerformsFileIo ULONG
VdlNotSet ULONG
UseIntermediateAddBuffer ULONG
CompressNoHuff ULONG
LockActiveRegions ULONG
VirtualRegions ULONG
Spare ULONG
}


type SMKM_REGION_EXTENT struct{
RegionCount ULONG
ByteOffset SIZE_T
}


type SMKM_FILE_INFO struct{
FileHandle HANDLE
_FILE_OBJECT struct
_FILE_OBJECT struct
_DEVICE_OBJECT struct
VolumePnpHandle HANDLE
_IRP struct
Extents PSMKM_REGION_EXTENT
ExtentCount ULONG
}


type SM_STORE_CACHE_BACKED_PARAMS struct{
SectorSize ULONG
EncryptionKey PCHAR
EncryptionKeySize ULONG
FileInfo PSMKM_FILE_INFO
EtaContext PVOID
_RTL_BITMAP struct
}


type SM_STORE_PARAMETERS struct{
Store SM_STORE_BASIC_PARAMS
Priority ULONG
Flags ULONG
CacheBacked SM_STORE_CACHE_BACKED_PARAMS
}


type SM_CREATE_REQUEST struct{
Version ULONG
AcquireReference ULONG
KeyedStore ULONG
Spare ULONG
Params SM_STORE_PARAMETERS
StoreId ULONG
}


type SM_DELETE_REQUEST struct{
Version ULONG
Spare ULONG
StoreId ULONG
}


type SM_STORE_LIST_REQUEST struct{
Version ULONG
StoreCount ULONG
ExtendedRequest ULONG
Spare ULONG
StoreId[32] ULONG
}


type SM_STORE_LIST_REQUEST_EX struct{
Request SM_STORE_LIST_REQUEST
NameBuffer[32][64] WCHAR
}


type SMC_CACHE_LIST_REQUEST struct{
Version ULONG
CacheCount ULONG
Spare ULONG
CacheId[16] ULONG
}


type SMC_CACHE_PARAMETERS struct{
CacheFileSize SIZE_T
StoreAlignment ULONG
PerformsFileIo ULONG
VdlNotSet ULONG
Spare ULONG
CacheFlags ULONG
Priority ULONG
}


type SMC_CACHE_CREATE_PARAMETERS struct{
CacheParameters SMC_CACHE_PARAMETERS
TemplateFilePath[512] WCHAR
}


type SMC_CACHE_CREATE_REQUEST struct{
Version ULONG
Spare ULONG
CacheId ULONG
CacheCreateParams SMC_CACHE_CREATE_PARAMETERS
}


type SMC_CACHE_DELETE_REQUEST struct{
Version ULONG
Spare ULONG
CacheId ULONG
}


type SMC_STORE_CREATE_REQUEST struct{
Version ULONG
Spare ULONG
StoreParams SM_STORE_BASIC_PARAMS
CacheId ULONG
StoreManagerType SM_STORE_MANAGER_TYPE
StoreId ULONG
}


type SMC_STORE_DELETE_REQUEST struct{
Version ULONG
Spare ULONG
CacheId ULONG
StoreManagerType SM_STORE_MANAGER_TYPE
StoreId ULONG
}


type SMC_CACHE_STATS struct{
TotalFileSize SIZE_T
StoreCount ULONG
RegionCount ULONG
RegionSizeBytes ULONG
FileCount ULONG
PerformsFileIo ULONG
Spare ULONG
StoreIds[16] ULONG
PhysicalStoreBitmap ULONG
Priority ULONG
TemplateFilePath[512] WCHAR
}


type SMC_CACHE_STATS_REQUEST struct{
Version ULONG
NoFilePath ULONG
Spare ULONG
CacheId ULONG
CacheStats SMC_CACHE_STATS
}


type SM_REGISTRATION_INFO struct{
CachesUpdatedEvent HANDLE
}


type SM_REGISTRATION_REQUEST struct{
Version ULONG
Spare ULONG
RegInfo SM_REGISTRATION_INFO
}


type SM_STORE_RESIZE_REQUEST struct{
Version ULONG
AddRegions ULONG
Spare ULONG
StoreId ULONG
NumberOfRegions ULONG
_RTL_BITMAP struct
}


type SMC_STORE_RESIZE_REQUEST struct{
Version ULONG
AddRegions ULONG
Spare ULONG
CacheId ULONG
StoreId ULONG
StoreManagerType SM_STORE_MANAGER_TYPE
RegionCount ULONG
}


type SM_CONFIG_REQUEST struct{
Version ULONG
Spare ULONG
ConfigType ULONG
ConfigValue ULONG
}


type SM_STORE_HIGH_MEM_PRIORITY_REQUEST struct{
Version ULONG
SetHighMemoryPriority ULONG
Spare ULONG
ProcessHandle HANDLE
}


type SM_SYSTEM_STORE_TRIM_REQUEST struct{
Version ULONG
Spare ULONG
PagesToTrim SIZE_T
}


type SM_MEM_COMPRESSION_INFO_REQUEST struct{
Version ULONG
Spare ULONG
CompressionPid ULONG
WorkingSetSize ULONG
TotalDataCompressed SIZE_T
TotalCompressedSize SIZE_T
TotalUniqueDataCompressed SIZE_T
}


type SYSTEM_REGISTRY_APPEND_STRING_PARAMETERS struct{
KeyHandle HANDLE
ValueNamePointer PUNICODE_STRING
RequiredLengthPointer PULONG
Buffer PUCHAR
BufferLength ULONG
Type ULONG
AppendBuffer PUCHAR
AppendBufferLength ULONG
CreateIfDoesntExist bool
TruncateExistingValue bool
}


type SYSTEM_VHD_BOOT_INFORMATION struct{
OsDiskIsVhd bool
OsVhdFilePathOffset ULONG
OsVhdParentVolume[1] WCHAR
}


type PS_CPU_QUOTA_QUERY_ENTRY struct{
SessionId ULONG
Weight ULONG
}


type PS_CPU_QUOTA_QUERY_INFORMATION struct{
SessionCount ULONG
SessionInformation[1] PS_CPU_QUOTA_QUERY_ENTRY
}


type SYSTEM_ERROR_PORT_TIMEOUTS struct{
StartTimeout ULONG
CommTimeout ULONG
}


type SYSTEM_LOW_PRIORITY_IO_INFORMATION struct{
LowPriReadOperations ULONG
LowPriWriteOperations ULONG
KernelBumpedToNormalOperations ULONG
LowPriPagingReadOperations ULONG
KernelPagingReadsBumpedToNormal ULONG
LowPriPagingWriteOperations ULONG
KernelPagingWritesBumpedToNormal ULONG
BoostedIrpCount ULONG
BoostedPagingIrpCount ULONG
BlanketBoostCount ULONG
}


type TPM_BOOT_ENTROPY_NT_RESULT struct{
Policy ULONGLONG
ResultCode TPM_BOOT_ENTROPY_RESULT_CODE
ResultStatus NTSTATUS
Time ULONGLONG
EntropyLength ULONG
EntropyData[40] UCHAR
}


type SYSTEM_VERIFIER_COUNTERS_INFORMATION struct{
Legacy SYSTEM_VERIFIER_INFORMATION
RaiseIrqls ULONG
AcquireSpinLocks ULONG
SynchronizeExecutions ULONG
AllocationsWithNoTag ULONG
AllocationsFailed ULONG
AllocationsFailedDeliberately ULONG
LockedBytes SIZE_T
PeakLockedBytes SIZE_T
MappedLockedBytes SIZE_T
PeakMappedLockedBytes SIZE_T
MappedIoSpaceBytes SIZE_T
PeakMappedIoSpaceBytes SIZE_T
PagesForMdlBytes SIZE_T
PeakPagesForMdlBytes SIZE_T
ContiguousMemoryBytes SIZE_T
PeakContiguousMemoryBytes SIZE_T
ExecutePoolTypes ULONG
ExecutePageProtections ULONG
ExecutePageMappings ULONG
ExecuteWriteSections ULONG
SectionAlignmentFailures ULONG
UnsupportedRelocs ULONG
IATInExecutableSection ULONG
}


type SYSTEM_ACPI_AUDIT_INFORMATION struct{
RsdpCount ULONG
SameRsdt ULONG
SlicPresent ULONG
SlicDifferent ULONG
}


type SYSTEM_BASIC_PERFORMANCE_INFORMATION struct{
AvailablePages SIZE_T
CommittedPages SIZE_T
CommitLimit SIZE_T
PeakCommitment SIZE_T
}


type QUERY_PERFORMANCE_COUNTER_FLAGS struct{
Union union
Struct struct
KernelTransition ULONG
Reserved ULONG
}


type SYSTEM_QUERY_PERFORMANCE_COUNTER_INFORMATION struct{
Version ULONG
Flags QUERY_PERFORMANCE_COUNTER_FLAGS
ValidFlags QUERY_PERFORMANCE_COUNTER_FLAGS
}


type SYSTEM_BOOT_GRAPHICS_INFORMATION struct{
FrameBuffer LARGE_INTEGER
Width ULONG
Height ULONG
PixelStride ULONG
Flags ULONG
Format SYSTEM_PIXEL_FORMAT
DisplayRotation ULONG
}


type MEMORY_SCRUB_INFORMATION struct{
Handle HANDLE
PagesScrubbed ULONG
}


type PEBS_DS_SAVE_AREA32 struct{
BtsBufferBase ULONG
BtsIndex ULONG
BtsAbsoluteMaximum ULONG
BtsInterruptThreshold ULONG
PebsBufferBase ULONG
PebsIndex ULONG
PebsAbsoluteMaximum ULONG
PebsInterruptThreshold ULONG
PebsGpCounterReset[8] ULONG
PebsFixedCounterReset[4] ULONG
}


type PEBS_DS_SAVE_AREA64 struct{
BtsBufferBase ULONGLONG
BtsIndex ULONGLONG
BtsAbsoluteMaximum ULONGLONG
BtsInterruptThreshold ULONGLONG
PebsBufferBase ULONGLONG
PebsIndex ULONGLONG
PebsAbsoluteMaximum ULONGLONG
PebsInterruptThreshold ULONGLONG
PebsGpCounterReset[8] ULONGLONG
PebsFixedCounterReset[4] ULONGLONG
}


type PROCESSOR_PROFILE_CONTROL_AREA struct{
PebsDsSaveArea PEBS_DS_SAVE_AREA
}


type SYSTEM_PROCESSOR_PROFILE_CONTROL_AREA struct{
ProcessorProfileControlArea PROCESSOR_PROFILE_CONTROL_AREA
Allocate bool
}


type MEMORY_COMBINE_INFORMATION struct{
Handle HANDLE
PagesCombined ULONG_PTR
}


type MEMORY_COMBINE_INFORMATION_EX struct{
Handle HANDLE
PagesCombined ULONG_PTR
Flags ULONG
}


type MEMORY_COMBINE_INFORMATION_EX2 struct{
Handle HANDLE
PagesCombined ULONG_PTR
Flags ULONG
ProcessHandle HANDLE
}


type SYSTEM_ENTROPY_TIMING_INFORMATION struct{
(NTAPI VOID
(NTAPI VOID
InitializationContext PVOID
}


type SYSTEM_CONSOLE_INFORMATION struct{
DriverLoaded ULONG
Spare ULONG
}


type SYSTEM_PLATFORM_BINARY_INFORMATION struct{
PhysicalAddress ULONG64
HandoffBuffer PVOID
CommandLineBuffer PVOID
HandoffBufferSize ULONG
CommandLineBufferSize ULONG
}


type SYSTEM_POLICY_INFORMATION struct{
InputData PVOID
OutputData PVOID
InputDataSize ULONG
OutputDataSize ULONG
Version ULONG
}


type SYSTEM_HYPERVISOR_PROCESSOR_COUNT_INFORMATION struct{
NumberOfLogicalProcessors ULONG
NumberOfCores ULONG
}


type SYSTEM_DEVICE_DATA_INFORMATION struct{
DeviceId UNICODE_STRING
DataName UNICODE_STRING
DataType ULONG
DataBufferLength ULONG
DataBuffer PVOID
}


type PHYSICAL_CHANNEL_RUN struct{
NodeNumber ULONG
ChannelNumber ULONG
BasePage ULONGLONG
PageCount ULONGLONG
Flags ULONG
}


type SYSTEM_MEMORY_TOPOLOGY_INFORMATION struct{
NumberOfRuns ULONGLONG
NumberOfNodes ULONG
NumberOfChannels ULONG
Run[1] PHYSICAL_CHANNEL_RUN
}


type SYSTEM_MEMORY_CHANNEL_INFORMATION struct{
ChannelNumber ULONG
ChannelHeatIndex ULONG
TotalPageCount ULONGLONG
ZeroPageCount ULONGLONG
FreePageCount ULONGLONG
StandbyPageCount ULONGLONG
}


type SYSTEM_BOOT_LOGO_INFORMATION struct{
Flags ULONG
BitmapOffset ULONG
}


type SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION_EX struct{
IdleTime LARGE_INTEGER
KernelTime LARGE_INTEGER
UserTime LARGE_INTEGER
DpcTime LARGE_INTEGER
InterruptTime LARGE_INTEGER
InterruptCount ULONG
Spare0 ULONG
AvailableTime LARGE_INTEGER
Spare1 LARGE_INTEGER
Spare2 LARGE_INTEGER
}


type SYSTEM_SECUREBOOT_POLICY_INFORMATION struct{
PolicyPublisher GUID
PolicyVersion ULONG
PolicyOptions ULONG
}


type SYSTEM_PAGEFILE_INFORMATION_EX struct{
Union union
Info SYSTEM_PAGEFILE_INFORMATION
Struct struct
NextEntryOffset ULONG
TotalSize ULONG
TotalInUse ULONG
PeakUsage ULONG
PageFileName UNICODE_STRING
}


type SYSTEM_SECUREBOOT_INFORMATION struct{
SecureBootEnabled bool
SecureBootCapable bool
}


type PROCESS_DISK_COUNTERS struct{
BytesRead ULONGLONG
BytesWritten ULONGLONG
ReadOperationCount ULONGLONG
WriteOperationCount ULONGLONG
FlushOperationCount ULONGLONG
}


type PROCESS_ENERGY_VALUES struct{
Cycles[4][2] ULONGLONG
DiskEnergy ULONGLONG
NetworkTailEnergy ULONGLONG
MBBTailEnergy ULONGLONG
NetworkTxRxBytes ULONGLONG
MBBTxRxBytes ULONGLONG
Union union
Durations[3] ENERGY_STATE_DURATION
Struct struct
ForegroundDuration ENERGY_STATE_DURATION
DesktopVisibleDuration ENERGY_STATE_DURATION
PSMForegroundDuration ENERGY_STATE_DURATION
}


type PROCESS_ENERGY_VALUES_EXTENSION struct{
Union union
Timelines[14] TIMELINE_BITMAP
Struct struct
CpuTimeline TIMELINE_BITMAP
DiskTimeline TIMELINE_BITMAP
NetworkTimeline TIMELINE_BITMAP
MBBTimeline TIMELINE_BITMAP
ForegroundTimeline TIMELINE_BITMAP
DesktopVisibleTimeline TIMELINE_BITMAP
CompositionRenderedTimeline TIMELINE_BITMAP
CompositionDirtyGeneratedTimeline TIMELINE_BITMAP
CompositionDirtyPropagatedTimeline TIMELINE_BITMAP
InputTimeline TIMELINE_BITMAP
AudioInTimeline TIMELINE_BITMAP
AudioOutTimeline TIMELINE_BITMAP
DisplayRequiredTimeline TIMELINE_BITMAP
KeyboardInputTimeline TIMELINE_BITMAP
}


type PROCESS_EXTENDED_ENERGY_VALUES struct{
Base PROCESS_ENERGY_VALUES
Extension PROCESS_ENERGY_VALUES_EXTENSION
}


type SYSTEM_PROCESS_INFORMATION_EXTENSION struct{
DiskCounters PROCESS_DISK_COUNTERS
ContextSwitches ULONGLONG
Union union
Flags ULONG
Struct struct
HasStrongId ULONG
Classification ULONG
BackgroundActivityModerated ULONG
Spare ULONG
}


type SYSTEM_PORTABLE_WORKSPACE_EFI_LAUNCHER_INFORMATION struct{
EfiLauncherEnabled bool
}


type SYSTEM_KERNEL_DEBUGGER_INFORMATION_EX struct{
DebuggerAllowed bool
DebuggerEnabled bool
DebuggerPresent bool
}


type SYSTEM_ELAM_CERTIFICATE_INFORMATION struct{
ElamDriverFile HANDLE
}


type OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V2 struct{
Version ULONG
AbnormalResetOccurred ULONG
OfflineMemoryDumpCapable ULONG
ResetDataAddress LARGE_INTEGER
ResetDataSize ULONG
}


type OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V1 struct{
Version ULONG
AbnormalResetOccurred ULONG
OfflineMemoryDumpCapable ULONG
}


type SYSTEM_PROCESSOR_FEATURES_INFORMATION struct{
ProcessorFeatureBits ULONGLONG
Reserved[3] ULONGLONG
}


type SYSTEM_EDID_INFORMATION struct{
Edid[128] UCHAR
}


type SYSTEM_MANUFACTURING_INFORMATION struct{
Options ULONG
ProfileName UNICODE_STRING
}


type SYSTEM_ENERGY_ESTIMATION_CONFIG_INFORMATION struct{
Enabled bool
}


type HV_DETAILS struct{
Data[4] ULONG
}


type SYSTEM_HYPERVISOR_DETAIL_INFORMATION struct{
HvVendorAndMaxFunction HV_DETAILS
HypervisorInterface HV_DETAILS
HypervisorVersion HV_DETAILS
HvFeatures HV_DETAILS
HwFeatures HV_DETAILS
EnlightenmentInfo HV_DETAILS
ImplementationLimits HV_DETAILS
}


type SYSTEM_PROCESSOR_CYCLE_STATS_INFORMATION struct{
Cycles[4][2] ULONGLONG
}


type SYSTEM_TPM_INFORMATION struct{
Flags ULONG
}


type SYSTEM_VSM_PROTECTION_INFORMATION struct{
DmaProtectionsAvailable bool
DmaProtectionsInUse bool
HardwareMbecAvailable bool
ApicVirtualizationAvailable bool
}


type SYSTEM_KERNEL_DEBUGGER_FLAGS struct{
KernelDebuggerIgnoreUmExceptions bool
}


type SYSTEM_CODEINTEGRITYPOLICY_INFORMATION struct{
Options ULONG
HVCIOptions ULONG
Version ULONGLONG
PolicyGuid GUID
}


type SYSTEM_ISOLATED_USER_MODE_INFORMATION struct{
SecureKernelRunning bool
HvciEnabled bool
HvciStrictMode bool
DebugEnabled bool
FirmwarePageProtection bool
EncryptionKeyAvailable bool
SpareFlags bool
TrustletRunning bool
HvciDisableAllowed bool
SpareFlags2 bool
Spare0[6] bool
Spare1 ULONGLONG
}


type SYSTEM_SINGLE_MODULE_INFORMATION struct{
TargetModuleAddress PVOID
ExInfo RTL_PROCESS_MODULE_INFORMATION_EX
}


type SYSTEM_INTERRUPT_CPU_SET_INFORMATION struct{
Gsiv ULONG
Group USHORT
CpuSets ULONGLONG
}


type SYSTEM_SECUREBOOT_POLICY_FULL_INFORMATION struct{
PolicyInformation SYSTEM_SECUREBOOT_POLICY_INFORMATION
PolicySize ULONG
Policy[1] UCHAR
}


type SYSTEM_ROOT_SILO_INFORMATION struct{
NumberOfSilos ULONG
SiloIdList[1] ULONG
}


type SYSTEM_CPU_SET_TAG_INFORMATION struct{
Tag ULONGLONG
CpuSets[1] ULONGLONG
}


type SYSTEM_SECURE_KERNEL_HYPERGUARD_PROFILE_INFORMATION struct{
ExtentCount ULONG
ValidStructureSize ULONG
NextExtentIndex ULONG
ExtentRestart ULONG
CycleCount ULONG
TimeoutCount ULONG
CycleTime ULONGLONG
CycleTimeMax ULONGLONG
ExtentTime ULONGLONG
ExtentTimeIndex ULONG
ExtentTimeMaxIndex ULONG
ExtentTimeMax ULONGLONG
HyperFlushTimeMax ULONGLONG
TranslateVaTimeMax ULONGLONG
DebugExemptionCount ULONGLONG
TbHitCount ULONGLONG
TbMissCount ULONGLONG
VinaPendingYield ULONGLONG
HashCycles ULONGLONG
HistogramOffset ULONG
HistogramBuckets ULONG
HistogramShift ULONG
Reserved1 ULONG
PageNotPresentCount ULONGLONG
}


type SYSTEM_SECUREBOOT_PLATFORM_MANIFEST_INFORMATION struct{
PlatformManifestSize ULONG
PlatformManifest[1] UCHAR
}


type SYSTEM_INTERRUPT_STEERING_INFORMATION_INPUT struct{
Gsiv ULONG
ControllerInterrupt UCHAR
EdgeInterrupt UCHAR
IsPrimaryInterrupt UCHAR
TargetAffinity GROUP_AFFINITY
}


type SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION struct{
Machine ULONG
KernelMode ULONG
UserMode ULONG
Native ULONG
Process ULONG
WoW64Container ULONG
ReservedZero0 ULONG
}


type SYSTEM_MEMORY_USAGE_INFORMATION struct{
TotalPhysicalBytes ULONGLONG
AvailableBytes ULONGLONG
ResidentAvailableBytes LONGLONG
CommittedBytes ULONGLONG
SharedCommittedBytes ULONGLONG
CommitLimitBytes ULONGLONG
PeakCommitmentBytes ULONGLONG
}


type SYSTEM_CODEINTEGRITY_CERTIFICATE_INFORMATION struct{
ImageFile HANDLE
Type ULONG
}


type SYSTEM_PHYSICAL_MEMORY_INFORMATION struct{
TotalPhysicalBytes ULONGLONG
LowestPhysicalAddress ULONGLONG
HighestPhysicalAddress ULONGLONG
}


type SYSTEM_ACTIVITY_MODERATION_EXE_STATE  struct{
ExePathNt UNICODE_STRING
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE
}


type SYSTEM_ACTIVITY_MODERATION_INFO struct{
Identifier UNICODE_STRING
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE
AppType SYSTEM_ACTIVITY_MODERATION_APP_TYPE
}


type SYSTEM_ACTIVITY_MODERATION_USER_SETTINGS struct{
UserKeyHandle HANDLE
}


type SYSTEM_CODEINTEGRITY_UNLOCK_INFORMATION struct{
Union union
Flags ULONG
Struct struct
Locked ULONG
UnlockApplied ULONG
UnlockIdValid ULONG
Reserved ULONG
}


type SYSTEM_FLUSH_INFORMATION struct{
SupportedFlushMethods ULONG
ProcessorCacheFlushSize ULONG
SystemFlushCapabilities ULONGLONG
Reserved[2] ULONGLONG
}


type SYSTEM_WRITE_CONSTRAINT_INFORMATION struct{
WriteConstraintPolicy ULONG
Reserved ULONG
}


type SYSTEM_KERNEL_VA_SHADOW_INFORMATION struct{
Union union
KvaShadowFlags ULONG
Struct struct
KvaShadowEnabled ULONG
KvaShadowUserGlobal ULONG
KvaShadowPcid ULONG
KvaShadowInvpcid ULONG
KvaShadowRequired ULONG
KvaShadowRequiredAvailable ULONG
InvalidPteBit ULONG
L1DataCacheFlushSupported ULONG
L1TerminalFaultMitigationPresent ULONG
Reserved ULONG
}


type SYSTEM_CODEINTEGRITYVERIFICATION_INFORMATION struct{
FileHandle HANDLE
ImageSize ULONG
Image PVOID
}


type SYSTEM_HYPERVISOR_SHARED_PAGE_INFORMATION struct{
HypervisorSharedUserVa PVOID
}


type SYSTEM_FIRMWARE_PARTITION_INFORMATION struct{
FirmwarePartition UNICODE_STRING
}


type SYSTEM_SPECULATION_CONTROL_INFORMATION struct{
Union union
Flags ULONG
Struct struct
BpbEnabled ULONG
BpbDisabledSystemPolicy ULONG
BpbDisabledNoHardwareSupport ULONG
SpecCtrlEnumerated ULONG
SpecCmdEnumerated ULONG
IbrsPresent ULONG
StibpPresent ULONG
SmepPresent ULONG
SpeculativeStoreBypassDisableAvailable ULONG
SpeculativeStoreBypassDisableSupported ULONG
SpeculativeStoreBypassDisabledSystemWide ULONG
SpeculativeStoreBypassDisabledKernel ULONG
SpeculativeStoreBypassDisableRequired ULONG
BpbDisabledKernelToUser ULONG
SpecCtrlRetpolineEnabled ULONG
SpecCtrlImportOptimizationEnabled ULONG
EnhancedIbrs ULONG
HvL1tfStatusAvailable ULONG
HvL1tfProcessorNotAffected ULONG
HvL1tfMigitationEnabled ULONG
HvL1tfMigitationNotEnabled_Hardware ULONG
HvL1tfMigitationNotEnabled_LoadOption ULONG
HvL1tfMigitationNotEnabled_CoreScheduler ULONG
EnhancedIbrsReported ULONG
MdsHardwareProtected ULONG
MbClearEnabled ULONG
MbClearReported ULONG
Reserved ULONG
}


type SYSTEM_DMA_GUARD_POLICY_INFORMATION struct{
DmaGuardPolicyEnabled bool
}


type SYSTEM_ENCLAVE_LAUNCH_CONTROL_INFORMATION struct{
EnclaveLaunchSigner[32] UCHAR
}


type SYSTEM_WORKLOAD_ALLOWED_CPU_SET_INFORMATION struct{
WorkloadClass ULONGLONG
CpuSets[1] ULONGLONG
}


type SYSTEM_SECURITY_MODEL_INFORMATION struct{
Union union
SecurityModelFlags ULONG
Struct struct
SModeAdminlessEnabled ULONG
AllowDeviceOwnerProtectionDowngrade ULONG
Reserved ULONG
}


type SYSTEM_FEATURE_CONFIGURATION_INFORMATION struct{
ChangeStamp ULONGLONG
_RTL_FEATURE_CONFIGURATION* struct
}


type SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY struct{
ChangeStamp ULONGLONG
Section PVOID
Size ULONGLONG
}


type SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION struct{
OverallChangeStamp ULONGLONG
Descriptors[3] SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY
}


type RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET struct{
Data[2] ULONG
}


type SYSTEM_FEATURE_USAGE_SUBSCRIPTION_DETAILS struct{
FeatureId ULONG
ReportingKind USHORT
ReportingOptions USHORT
ReportingTarget RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET
}


type SYSTEM_FIRMWARE_RAMDISK_INFORMATION struct{
Version ULONG
BlockSize ULONG
BaseAddress ULONG_PTR
Size SIZE_T
}


type SYSTEM_SHADOW_STACK_INFORMATION struct{
Union union
Flags ULONG
Struct struct
CetCapable ULONG
UserCetAllowed ULONG
ReservedForUserCet ULONG
KernelCetEnabled ULONG
KernelCetAuditModeEnabled ULONG
ReservedForKernelCet ULONG
Reserved ULONG
}


type SYSTEM_BUILD_VERSION_INFORMATION struct{
LayerNumber USHORT
LayerCount USHORT
OsMajorVersion ULONG
OsMinorVersion ULONG
NtBuildNumber ULONG
NtBuildQfe ULONG
LayerName[128] UCHAR
NtBuildBranch[128] UCHAR
NtBuildLab[128] UCHAR
NtBuildLabEx[128] UCHAR
NtBuildStamp[26] UCHAR
NtBuildArch[16] UCHAR
Flags SYSTEM_BUILD_VERSION_INFORMATION_FLAGS
}


type SYSTEM_POOL_LIMIT_MEM_INFO struct{
MemoryLimit ULONGLONG
NotificationLimit ULONGLONG
}


type SYSTEM_POOL_LIMIT_INFO struct{
PoolTag ULONG
MemLimits[2] SYSTEM_POOL_LIMIT_MEM_INFO
NotificationHandle WNF_STATE_NAME
}


type SYSTEM_POOL_LIMIT_INFORMATION struct{
Version ULONG
EntryCount ULONG
LimitEntries[1] SYSTEM_POOL_LIMIT_INFO
}


type HV_MINROOT_NUMA_LPS struct{
NodeIndex ULONG
Mask[16] ULONG_PTR
}


type SYSTEM_IOMMU_STATE_INFORMATION struct{
State SYSTEM_IOMMU_STATE
Pdo PVOID
}


type SYSTEM_HYPERVISOR_MINROOT_INFORMATION struct{
NumProc ULONG
RootProc ULONG
RootProcNumaNodesSpecified ULONG
RootProcNumaNodes[64] USHORT
RootProcPerCore ULONG
RootProcPerNode ULONG
RootProcNumaNodesLpsSpecified ULONG
RootProcNumaNodeLps[64] HV_MINROOT_NUMA_LPS
}


type SYSTEM_HYPERVISOR_BOOT_PAGES_INFORMATION struct{
RangeCount ULONG
RangeArray[1] ULONG_PTR
}


type SYSTEM_POINTER_AUTH_INFORMATION struct{
Union union
SupportedFlags USHORT
Struct struct
AddressAuthSupported USHORT
AddressAuthQarma USHORT
GenericAuthSupported USHORT
GenericAuthQarma USHORT
SupportedReserved USHORT
}


type SYSDBG_VIRTUAL struct{
Address PVOID
Buffer PVOID
Request ULONG
}


type SYSDBG_PHYSICAL struct{
Address PHYSICAL_ADDRESS
Buffer PVOID
Request ULONG
}


type SYSDBG_CONTROL_SPACE struct{
Address ULONG64
Buffer PVOID
Request ULONG
Processor ULONG
}


type SYSDBG_IO_SPACE struct{
Address ULONG64
Buffer PVOID
Request ULONG
_INTERFACE_TYPE enum
BusNumber ULONG
AddressSpace ULONG
}


type SYSDBG_MSR struct{
Msr ULONG
Data ULONG64
}


type SYSDBG_BUS_DATA struct{
Address ULONG
Buffer PVOID
Request ULONG
_BUS_DATA_TYPE enum
BusNumber ULONG
SlotNumber ULONG
}


type SYSDBG_TRIAGE_DUMP struct{
Flags ULONG
BugCheckCode ULONG
BugCheckParam1 ULONG_PTR
BugCheckParam2 ULONG_PTR
BugCheckParam3 ULONG_PTR
BugCheckParam4 ULONG_PTR
ProcessHandles ULONG
ThreadHandles ULONG
Handles PHANDLE
}


type SYSDBG_LIVEDUMP_SELECTIVE_CONTROL struct{
Version ULONG
Size ULONG
Union union
Flags ULONGLONG
Struct struct
ThreadKernelStacks ULONGLONG
ReservedFlags ULONGLONG
}


type SYSDBG_LIVEDUMP_CONTROL struct{
Version ULONG
BugCheckCode ULONG
BugCheckParam1 ULONG_PTR
BugCheckParam2 ULONG_PTR
BugCheckParam3 ULONG_PTR
BugCheckParam4 ULONG_PTR
DumpFileHandle HANDLE
CancelEventHandle HANDLE
Flags SYSDBG_LIVEDUMP_CONTROL_FLAGS
AddPagesControl SYSDBG_LIVEDUMP_CONTROL_ADDPAGES
SelectiveControl PSYSDBG_LIVEDUMP_SELECTIVE_CONTROL
}


type SYSDBG_KD_PULL_REMOTE_FILE struct{
ImageFileName UNICODE_STRING
}


type KUSER_SHARED_DATA struct{
TickCountLowDeprecated ULONG
TickCountMultiplier ULONG
KSYSTEM_TIME volatile
KSYSTEM_TIME volatile
KSYSTEM_TIME volatile
ImageNumberLow USHORT
ImageNumberHigh USHORT
NtSystemRoot[260] WCHAR
MaxStackTraceDepth ULONG
CryptoExponent ULONG
TimeZoneId ULONG
LargePageMinimum ULONG
AitSamplingValue ULONG
AppCompatFlag ULONG
RNGSeedVersion ULONGLONG
GlobalValidationRunlevel ULONG
TimeZoneBiasStamp LONG
NtBuildNumber ULONG
NtProductType NT_PRODUCT_TYPE
ProductTypeIsValid bool
Reserved0[1] UCHAR
NativeProcessorArchitecture USHORT
NtMajorVersion ULONG
NtMinorVersion ULONG
ProcessorFeatures[PROCESSOR_FEATURE_MAX] bool
Reserved1 ULONG
Reserved3 ULONG
ULONG volatile
AlternativeArchitecture ALTERNATIVE_ARCHITECTURE_TYPE
BootId ULONG
SystemExpirationDate LARGE_INTEGER
SuiteMask ULONG
KdDebuggerEnabled bool
Union union
MitigationPolicies UCHAR
Struct struct
NXSupportPolicy UCHAR
SEHValidationPolicy UCHAR
CurDirDevicesSkippedForDlls UCHAR
Reserved UCHAR
}


type ATOM_BASIC_INFORMATION struct{
UsageCount USHORT
Flags USHORT
NameLength USHORT
Name[1] WCHAR
}


type ATOM_TABLE_INFORMATION struct{
NumberOfAtoms ULONG
Atoms[1] RTL_ATOM
}



type (
Ntexapi interface{
NtDelayExecution()(ok bool)//col:831
#if_!defined()(ok bool)//col:879
#if_()(ok bool)//col:934
NtSystemDebugControl()(ok bool)//col:1005
FORCEINLINE_ULONG_NtGetTickCount()(ok bool)//col:1023
FORCEINLINE_ULONGLONG_NtGetTickCount64()(ok bool)//col:1027
FORCEINLINE_ULONG_NtGetTickCount()(ok bool)//col:1031
}
)

func NewNtexapi() { return & ntexapi{} }

func (n *ntexapi)NtDelayExecution()(ok bool){//col:831
/*NtDelayExecution(
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER DelayInterval
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySystemEnvironmentValue(
    _In_ PUNICODE_STRING VariableName,
    _Out_writes_bytes_(ValueLength) PWSTR VariableValue,
    _In_ USHORT ValueLength,
    _Out_opt_ PUSHORT ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSystemEnvironmentValue(
    _In_ PUNICODE_STRING VariableName,
    _In_ PUNICODE_STRING VariableValue
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySystemEnvironmentValueEx(
    _In_ PUNICODE_STRING VariableName,
    _In_ LPGUID VendorGuid,
    _Out_writes_bytes_opt_(*ValueLength) PVOID Value,
    _Inout_ PULONG ValueLength,
    _Out_opt_ PULONG Attributes 
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSystemEnvironmentValueEx(
    _In_ PUNICODE_STRING VariableName,
    _In_ LPGUID VendorGuid,
    _In_reads_bytes_opt_(ValueLength) PVOID Value,
    _In_ ULONG ValueLength, 
    _In_ ULONG Attributes 
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtEnumerateSystemEnvironmentValuesEx(
    _In_ ULONG InformationClass,
    _Out_ PVOID Buffer,
    _Inout_ PULONG BufferLength
    );
#if (PHNT_VERSION >= PHNT_WINXP)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAddBootEntry(
    _In_ PBOOT_ENTRY BootEntry,
    _Out_opt_ PULONG Id
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDeleteBootEntry(
    _In_ ULONG Id
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtModifyBootEntry(
    _In_ PBOOT_ENTRY BootEntry
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtEnumerateBootEntries(
    _Out_writes_bytes_opt_(*BufferLength) PVOID Buffer,
    _Inout_ PULONG BufferLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryBootEntryOrder(
    _Out_writes_opt_(*Count) PULONG Ids,
    _Inout_ PULONG Count
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetBootEntryOrder(
    _In_reads_(Count) PULONG Ids,
    _In_ ULONG Count
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryBootOptions(
    _Out_writes_bytes_opt_(*BootOptionsLength) PBOOT_OPTIONS BootOptions,
    _Inout_ PULONG BootOptionsLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetBootOptions(
    _In_ PBOOT_OPTIONS BootOptions,
    _In_ ULONG FieldsToChange
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTranslateFilePath(
    _In_ PFILE_PATH InputFilePath,
    _In_ ULONG OutputType,
    _Out_writes_bytes_opt_(*OutputFilePathLength) PFILE_PATH OutputFilePath,
    _Inout_opt_ PULONG OutputFilePathLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAddDriverEntry(
    _In_ PEFI_DRIVER_ENTRY DriverEntry,
    _Out_opt_ PULONG Id
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDeleteDriverEntry(
    _In_ ULONG Id
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtModifyDriverEntry(
    _In_ PEFI_DRIVER_ENTRY DriverEntry
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtEnumerateDriverEntries(
    _Out_writes_bytes_opt_(*BufferLength) PVOID Buffer,
    _Inout_ PULONG BufferLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryDriverEntryOrder(
    _Out_writes_opt_(*Count) PULONG Ids,
    _Inout_ PULONG Count
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetDriverEntryOrder(
    _In_reads_(Count) PULONG Ids,
    _In_ ULONG Count
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtFilterBootOption(
    _In_ FILTER_BOOT_OPTION_OPERATION FilterOperation,
    _In_ ULONG ObjectType,
    _In_ ULONG ElementType,
    _In_reads_bytes_opt_(DataSize) PVOID Data,
    _In_ ULONG DataSize
    );
#endif
#ifndef EVENT_QUERY_STATE
#endif
#ifndef EVENT_MODIFY_STATE
#endif
#ifndef EVENT_ALL_ACCESS
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateEvent(
    _Out_ PHANDLE EventHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ EVENT_TYPE EventType,
    _In_ BOOLEAN InitialState
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenEvent(
    _Out_ PHANDLE EventHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetEvent(
    _In_ HANDLE EventHandle,
    _Out_opt_ PLONG PreviousState
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetEventBoostPriority(
    _In_ HANDLE EventHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtClearEvent(
    _In_ HANDLE EventHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResetEvent(
    _In_ HANDLE EventHandle,
    _Out_opt_ PLONG PreviousState
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtPulseEvent(
    _In_ HANDLE EventHandle,
    _Out_opt_ PLONG PreviousState
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryEvent(
    _In_ HANDLE EventHandle,
    _In_ EVENT_INFORMATION_CLASS EventInformationClass,
    _Out_writes_bytes_(EventInformationLength) PVOID EventInformation,
    _In_ ULONG EventInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateEventPair(
    _Out_ PHANDLE EventPairHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenEventPair(
    _Out_ PHANDLE EventPairHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLowEventPair(
    _In_ HANDLE EventPairHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetHighEventPair(
    _In_ HANDLE EventPairHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitLowEventPair(
    _In_ HANDLE EventPairHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitHighEventPair(
    _In_ HANDLE EventPairHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLowWaitHighEventPair(
    _In_ HANDLE EventPairHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetHighWaitLowEventPair(
    _In_ HANDLE EventPairHandle
    );
#ifndef MUTANT_QUERY_STATE
#endif
#ifndef MUTANT_ALL_ACCESS
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateMutant(
    _Out_ PHANDLE MutantHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ BOOLEAN InitialOwner
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenMutant(
    _Out_ PHANDLE MutantHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReleaseMutant(
    _In_ HANDLE MutantHandle,
    _Out_opt_ PLONG PreviousCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryMutant(
    _In_ HANDLE MutantHandle,
    _In_ MUTANT_INFORMATION_CLASS MutantInformationClass,
    _Out_writes_bytes_(MutantInformationLength) PVOID MutantInformation,
    _In_ ULONG MutantInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#ifndef SEMAPHORE_QUERY_STATE
#endif
#ifndef SEMAPHORE_MODIFY_STATE
#endif
#ifndef SEMAPHORE_ALL_ACCESS
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateSemaphore(
    _Out_ PHANDLE SemaphoreHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ LONG InitialCount,
    _In_ LONG MaximumCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenSemaphore(
    _Out_ PHANDLE SemaphoreHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReleaseSemaphore(
    _In_ HANDLE SemaphoreHandle,
    _In_ LONG ReleaseCount,
    _Out_opt_ PLONG PreviousCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySemaphore(
    _In_ HANDLE SemaphoreHandle,
    _In_ SEMAPHORE_INFORMATION_CLASS SemaphoreInformationClass,
    _Out_writes_bytes_(SemaphoreInformationLength) PVOID SemaphoreInformation,
    _In_ ULONG SemaphoreInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#ifndef TIMER_QUERY_STATE
#endif
#ifndef TIMER_MODIFY_STATE
#endif
#ifndef TIMER_ALL_ACCESS
#endif
typedef VOID (NTAPI *PTIMER_APC_ROUTINE)(
    _In_ PVOID TimerContext,
    _In_ ULONG TimerLowValue,
    _In_ LONG TimerHighValue
    );
#if (PHNT_VERSION >= PHNT_WIN7)
struct _COUNTED_REASON_CONTEXT;
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateTimer(
    _Out_ PHANDLE TimerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ TIMER_TYPE TimerType
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenTimer(
    _Out_ PHANDLE TimerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetTimer(
    _In_ HANDLE TimerHandle,
    _In_ PLARGE_INTEGER DueTime,
    _In_opt_ PTIMER_APC_ROUTINE TimerApcRoutine,
    _In_opt_ PVOID TimerContext,
    _In_ BOOLEAN ResumeTimer,
    _In_opt_ LONG Period,
    _Out_opt_ PBOOLEAN PreviousState
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetTimerEx(
    _In_ HANDLE TimerHandle,
    _In_ TIMER_SET_INFORMATION_CLASS TimerSetInformationClass,
    _Inout_updates_bytes_opt_(TimerSetInformationLength) PVOID TimerSetInformation,
    _In_ ULONG TimerSetInformationLength
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCancelTimer(
    _In_ HANDLE TimerHandle,
    _Out_opt_ PBOOLEAN CurrentState
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryTimer(
    _In_ HANDLE TimerHandle,
    _In_ TIMER_INFORMATION_CLASS TimerInformationClass,
    _Out_writes_bytes_(TimerInformationLength) PVOID TimerInformation,
    _In_ ULONG TimerInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateIRTimer(
    _Out_ PHANDLE TimerHandle,
    _In_ ACCESS_MASK DesiredAccess
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetIRTimer(
    _In_ HANDLE TimerHandle,
    _In_opt_ PLARGE_INTEGER DueTime
    );
#endif
typedef PVOID PT2_CANCEL_PARAMETERS;
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateTimer2(
    _Out_ PHANDLE TimerHandle,
    _In_opt_ PVOID Reserved1,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ ULONG Attributes,
    _In_ ACCESS_MASK DesiredAccess
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetTimer2(
    _In_ HANDLE TimerHandle,
    _In_ PLARGE_INTEGER DueTime,
    _In_opt_ PLARGE_INTEGER Period,
    _In_ PT2_SET_PARAMETERS Parameters
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCancelTimer2(
    _In_ HANDLE TimerHandle,
    _In_ PT2_CANCEL_PARAMETERS Parameters
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProfile(
    _Out_ PHANDLE ProfileHandle,
    _In_opt_ HANDLE Process,
    _In_ PVOID ProfileBase,
    _In_ SIZE_T ProfileSize,
    _In_ ULONG BucketSize,
    _In_reads_bytes_(BufferSize) PULONG Buffer,
    _In_ ULONG BufferSize,
    _In_ KPROFILE_SOURCE ProfileSource,
    _In_ KAFFINITY Affinity
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProfileEx(
    _Out_ PHANDLE ProfileHandle,
    _In_opt_ HANDLE Process,
    _In_ PVOID ProfileBase,
    _In_ SIZE_T ProfileSize,
    _In_ ULONG BucketSize,
    _In_reads_bytes_(BufferSize) PULONG Buffer,
    _In_ ULONG BufferSize,
    _In_ KPROFILE_SOURCE ProfileSource,
    _In_ USHORT GroupCount,
    _In_reads_(GroupCount) PGROUP_AFFINITY GroupAffinity
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtStartProfile(
    _In_ HANDLE ProfileHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtStopProfile(
    _In_ HANDLE ProfileHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryIntervalProfile(
    _In_ KPROFILE_SOURCE ProfileSource,
    _Out_ PULONG Interval
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetIntervalProfile(
    _In_ ULONG Interval,
    _In_ KPROFILE_SOURCE Source
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateKeyedEvent(
    _Out_ PHANDLE KeyedEventHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ ULONG Flags
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenKeyedEvent(
    _Out_ PHANDLE KeyedEventHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReleaseKeyedEvent(
    _In_ HANDLE KeyedEventHandle,
    _In_ PVOID KeyValue,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForKeyedEvent(
    _In_ HANDLE KeyedEventHandle,
    _In_ PVOID KeyValue,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUmsThreadYield(
    _In_ PVOID SchedulerParam
    );
#endif
typedef const WNF_STATE_NAME *PCWNF_STATE_NAME;
typedef const WNF_TYPE_ID *PCWNF_TYPE_ID;
typedef ULONG WNF_CHANGE_STAMP, *PWNF_CHANGE_STAMP;
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateWnfStateName(
    _Out_ PWNF_STATE_NAME StateName,
    _In_ WNF_STATE_NAME_LIFETIME NameLifetime,
    _In_ WNF_DATA_SCOPE DataScope,
    _In_ BOOLEAN PersistData,
    _In_opt_ PCWNF_TYPE_ID TypeId,
    _In_ ULONG MaximumStateSize,
    _In_ PSECURITY_DESCRIPTOR SecurityDescriptor
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDeleteWnfStateName(
    _In_ PCWNF_STATE_NAME StateName
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUpdateWnfStateData(
    _In_ PCWNF_STATE_NAME StateName,
    _In_reads_bytes_opt_(Length) const VOID *Buffer,
    _In_opt_ ULONG Length,
    _In_opt_ PCWNF_TYPE_ID TypeId,
    _In_opt_ const VOID *ExplicitScope,
    _In_ WNF_CHANGE_STAMP MatchingChangeStamp,
    _In_ LOGICAL CheckStamp
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDeleteWnfStateData(
    _In_ PCWNF_STATE_NAME StateName,
    _In_opt_ const VOID *ExplicitScope
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryWnfStateData(
    _In_ PCWNF_STATE_NAME StateName,
    _In_opt_ PCWNF_TYPE_ID TypeId,
    _In_opt_ const VOID *ExplicitScope,
    _Out_ PWNF_CHANGE_STAMP ChangeStamp,
    _Out_writes_bytes_to_opt_(*BufferSize, *BufferSize) PVOID Buffer,
    _Inout_ PULONG BufferSize
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryWnfStateNameInformation(
    _In_ PCWNF_STATE_NAME StateName,
    _In_ WNF_STATE_NAME_INFORMATION NameInfoClass,
    _In_opt_ const VOID *ExplicitScope,
    _Out_writes_bytes_(InfoBufferSize) PVOID InfoBuffer,
    _In_ ULONG InfoBufferSize
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSubscribeWnfStateChange(
    _In_ PCWNF_STATE_NAME StateName,
    _In_opt_ WNF_CHANGE_STAMP ChangeStamp,
    _In_ ULONG EventMask,
    _Out_opt_ PULONG64 SubscriptionId
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtUnsubscribeWnfStateChange(
    _In_ PCWNF_STATE_NAME StateName
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetCompleteWnfStateSubscription(
    _In_opt_ PWNF_STATE_NAME OldDescriptorStateName,
    _In_opt_ ULONG64 *OldSubscriptionId,
    _In_opt_ ULONG OldDescriptorEventMask,
    _In_opt_ ULONG OldDescriptorStatus,
    _Out_writes_bytes_(DescriptorSize) PWNF_DELIVERY_DESCRIPTOR NewDeliveryDescriptor,
    _In_ ULONG DescriptorSize
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetWnfProcessNotificationEvent(
    _In_ HANDLE NotificationEvent
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateWorkerFactory(
    _Out_ PHANDLE WorkerFactoryHandleReturn,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE CompletionPortHandle,
    _In_ HANDLE WorkerProcessHandle,
    _In_ PVOID StartRoutine,
    _In_opt_ PVOID StartParameter,
    _In_opt_ ULONG MaxThreadCount,
    _In_opt_ SIZE_T StackReserve,
    _In_opt_ SIZE_T StackCommit
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationWorkerFactory(
    _In_ HANDLE WorkerFactoryHandle,
    _In_ WORKERFACTORYINFOCLASS WorkerFactoryInformationClass,
    _Out_writes_bytes_(WorkerFactoryInformationLength) PVOID WorkerFactoryInformation,
    _In_ ULONG WorkerFactoryInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationWorkerFactory(
    _In_ HANDLE WorkerFactoryHandle,
    _In_ WORKERFACTORYINFOCLASS WorkerFactoryInformationClass,
    _In_reads_bytes_(WorkerFactoryInformationLength) PVOID WorkerFactoryInformation,
    _In_ ULONG WorkerFactoryInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtShutdownWorkerFactory(
    _In_ HANDLE WorkerFactoryHandle,
    _Inout_ volatile LONG *PendingWorkerCount
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReleaseWorkerFactoryWorker(
    _In_ HANDLE WorkerFactoryHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWorkerFactoryWorkerReady(
    _In_ HANDLE WorkerFactoryHandle
    );
struct _FILE_IO_COMPLETION_INFORMATION;
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForWorkViaWorkerFactory(
    _In_ HANDLE WorkerFactoryHandle,
    _Out_writes_to_(Count, *PacketsReturned) struct _FILE_IO_COMPLETION_INFORMATION *MiniPackets,
    _In_ ULONG Count,
    _Out_ PULONG PacketsReturned,
    _In_ PWORKER_FACTORY_DEFERRED_WORK DeferredWork
    );
#else
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForWorkViaWorkerFactory(
    _In_ HANDLE WorkerFactoryHandle,
    _Out_ struct _FILE_IO_COMPLETION_INFORMATION *MiniPacket
    );
#endif
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySystemTime(
    _Out_ PLARGE_INTEGER SystemTime
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSystemTime(
    _In_opt_ PLARGE_INTEGER SystemTime,
    _Out_opt_ PLARGE_INTEGER PreviousTime
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryTimerResolution(
    _Out_ PULONG MaximumTime,
    _Out_ PULONG MinimumTime,
    _Out_ PULONG CurrentTime
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetTimerResolution(
    _In_ ULONG DesiredTime,
    _In_ BOOLEAN SetResolution,
    _Out_ PULONG ActualTime
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryPerformanceCounter(
    _Out_ PLARGE_INTEGER PerformanceCounter,
    _Out_opt_ PLARGE_INTEGER PerformanceFrequency
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateLocallyUniqueId(
    _Out_ PLUID Luid
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetUuidSeed(
    _In_ PCHAR Seed
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateUuids(
    _Out_ PULARGE_INTEGER Time,
    _Out_ PULONG Range,
    _Out_ PULONG Sequence,
    _Out_ PCHAR Seed
    );
#endif 
#ifndef _TRACEHANDLE_DEFINED
typedef ULONG64 TRACEHANDLE, *PTRACEHANDLE;
#endif
typedef ULONG PERFINFO_MASK;
#ifdef _WIN64
#else
#endif
typedef NTSTATUS (*PSYSTEM_WATCHDOG_HANDLER)(_In_ WATCHDOG_HANDLER_ACTION Action, _In_ PVOID Context, _Inout_ PULONG DataValue, _In_ BOOLEAN NoLocks);
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef NTSTATUS (__cdecl* PFNFTH)(
    _Inout_ PSYSTEM_FIRMWARE_TABLE_INFORMATION SystemFirmwareTableInfo
    );
#endif
typedef union _PEBS_DS_SAVE_AREA
{
    PEBS_DS_SAVE_AREA32 As32Bit;
    PEBS_DS_SAVE_AREA64 As64Bit;
} PEBS_DS_SAVE_AREA, *PPEBS_DS_SAVE_AREA;*/
return true
}

func (n *ntexapi)#if_!defined()(ok bool){//col:879
/*#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO) && !PHNT_PATCH_FOR_HYPERDBG
#endif
typedef union _SECURE_SPECULATION_CONTROL_INFORMATION
{
    ULONG KvaShadowSupported : 1;
    ULONG KvaShadowEnabled : 1;
    ULONG KvaShadowUserGlobal : 1;
    ULONG KvaShadowPcid : 1;
    ULONG MbClearEnabled : 1;
    ULONG L1TFMitigated : 1; 
    ULONG BpbEnabled : 1;
    ULONG IbrsPresent : 1;
    ULONG EnhancedIbrs : 1;
    ULONG StibpPresent : 1;
    ULONG SsbdSupported : 1;
    ULONG SsbdRequired : 1;
    ULONG BpbKernelToUser : 1;
    ULONG BpbUserToKernel : 1;
    ULONG Reserved : 18;
} SECURE_SPECULATION_CONTROL_INFORMATION, *PSECURE_SPECULATION_CONTROL_INFORMATION;*/
return true
}

func (n *ntexapi)#if_()(ok bool){//col:934
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySystemInformation(
    _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass,
    _Out_writes_bytes_opt_(SystemInformationLength) PVOID SystemInformation,
    _In_ ULONG SystemInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySystemInformationEx(
    _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass,
    _In_reads_bytes_(InputBufferLength) PVOID InputBuffer,
    _In_ ULONG InputBufferLength,
    _Out_writes_bytes_opt_(SystemInformationLength) PVOID SystemInformation,
    _In_ ULONG SystemInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSystemInformation(
    _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass,
    _In_reads_bytes_opt_(SystemInformationLength) PVOID SystemInformation,
    _In_ ULONG SystemInformationLength
    );
enum _INTERFACE_TYPE;
enum _BUS_DATA_TYPE;
typedef union _SYSDBG_LIVEDUMP_CONTROL_FLAGS
{
    struct
    {
        ULONG UseDumpStorageStack : 1;
        ULONG CompressMemoryPagesData : 1;
        ULONG IncludeUserSpaceMemoryPages : 1;
        ULONG AbortIfMemoryPressure : 1; 
        ULONG SelectiveDump : 1; 
        ULONG Reserved : 27;
    };
    ULONG AsUlong;
} SYSDBG_LIVEDUMP_CONTROL_FLAGS, *PSYSDBG_LIVEDUMP_CONTROL_FLAGS;*/
return true
}

func (n *ntexapi)NtSystemDebugControl()(ok bool){//col:1005
/*NtSystemDebugControl(
    _In_ SYSDBG_COMMAND Command,
    _Inout_updates_bytes_opt_(InputBufferLength) PVOID InputBuffer,
    _In_ ULONG InputBufferLength,
    _Out_writes_bytes_opt_(OutputBufferLength) PVOID OutputBuffer,
    _In_ ULONG OutputBufferLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRaiseHardError(
    _In_ NTSTATUS ErrorStatus,
    _In_ ULONG NumberOfParameters,
    _In_ ULONG UnicodeStringParameterMask,
    _In_reads_(NumberOfParameters) PULONG_PTR Parameters,
    _In_ ULONG ValidResponseOptions,
    _Out_ PULONG Response
    );
#include <pshpack4.h>
#include <poppack.h>
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, TickCountMultiplier) == 0x4);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, InterruptTime) == 0x8);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, SystemTime) == 0x14);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, NtSystemRoot) == 0x30);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, LargePageMinimum) == 0x244);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, NtProductType) == 0x264);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, NtMajorVersion) == 0x26c);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, NtMinorVersion) == 0x270);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, ProcessorFeatures) == 0x274);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, KdDebuggerEnabled) == 0x2d4);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, ActiveConsoleId) == 0x2d8);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, NumberOfPhysicalPages) == 0x2e8);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, SafeBootMode) == 0x2ec);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, TickCount) == 0x320);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, TickCountQuad) == 0x320);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, ActiveProcessorCount) == 0x3c0);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, ActiveGroupCount) == 0x3c4);
C_ASSERT(FIELD_OFFSET(KUSER_SHARED_DATA, XState) == 0x3d8);
#if (PHNT_VERSION >= PHNT_WS03)
FORCEINLINE ULONGLONG NtGetTickCount64()
{
    ULARGE_INTEGER tickCount;
#ifdef _WIN64
    tickCount.QuadPart = USER_SHARED_DATA->TickCountQuad;
#else
    while (TRUE)
    {
        tickCount.HighPart = (ULONG)USER_SHARED_DATA->TickCount.High1Time;
        tickCount.LowPart = USER_SHARED_DATA->TickCount.LowPart;
        if (tickCount.HighPart == (ULONG)USER_SHARED_DATA->TickCount.High2Time)
            break;
        YieldProcessor();
    }
#endif
    return (UInt32x32To64(tickCount.LowPart, USER_SHARED_DATA->TickCountMultiplier) >> 24) +
        (UInt32x32To64(tickCount.HighPart, USER_SHARED_DATA->TickCountMultiplier) << 8);
}*/
return true
}

func (n *ntexapi)FORCEINLINE_ULONG_NtGetTickCount()(ok bool){//col:1023
/*FORCEINLINE ULONG NtGetTickCount()
{
#ifdef _WIN64
    return (ULONG)((USER_SHARED_DATA->TickCountQuad * USER_SHARED_DATA->TickCountMultiplier) >> 24);
#else
    ULARGE_INTEGER tickCount;
    while (TRUE)
    {
        tickCount.HighPart = (ULONG)USER_SHARED_DATA->TickCount.High1Time;
        tickCount.LowPart = USER_SHARED_DATA->TickCount.LowPart;
        if (tickCount.HighPart == (ULONG)USER_SHARED_DATA->TickCount.High2Time)
            break;
        YieldProcessor();
    }
    return (ULONG)((UInt32x32To64(tickCount.LowPart, USER_SHARED_DATA->TickCountMultiplier) >> 24) +
        UInt32x32To64((tickCount.HighPart << 8) & 0xffffffff, USER_SHARED_DATA->TickCountMultiplier));
#endif
}*/
return true
}

func (n *ntexapi)FORCEINLINE_ULONGLONG_NtGetTickCount64()(ok bool){//col:1027
/*FORCEINLINE ULONGLONG NtGetTickCount64()
{
    return GetTickCount(); 
}*/
return true
}

func (n *ntexapi)FORCEINLINE_ULONG_NtGetTickCount()(ok bool){//col:1031
/*FORCEINLINE ULONG NtGetTickCount()
{
    return GetTickCount();
}*/
return true
}



