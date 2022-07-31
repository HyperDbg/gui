package phnt


const(
_NTEXAPI_H =  //col:13
EFI_VARIABLE_NON_VOLATILE = 0x00000001 //col:49
EFI_VARIABLE_BOOTSERVICE_ACCESS = 0x00000002 //col:50
EFI_VARIABLE_RUNTIME_ACCESS = 0x00000004 //col:51
EFI_VARIABLE_HARDWARE_ERROR_RECORD = 0x00000008 //col:52
EFI_VARIABLE_AUTHENTICATED_WRITE_ACCESS = 0x00000010 //col:53
EFI_VARIABLE_TIME_BASED_AUTHENTICATED_WRITE_ACCESS = 0x00000020 //col:54
EFI_VARIABLE_APPEND_WRITE = 0x00000040 //col:55
EFI_VARIABLE_ENHANCED_AUTHENTICATED_ACCESS = 0x00000080 //col:56
EVENT_QUERY_STATE = 0x0001 //col:296
EVENT_MODIFY_STATE = 0x0002 //col:300
EVENT_ALL_ACCESS = (EVENT_QUERY_STATE|EVENT_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:304
EVENT_PAIR_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE) //col:389
MUTANT_QUERY_STATE = 0x0001 //col:454
MUTANT_ALL_ACCESS = (MUTANT_QUERY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:458
SEMAPHORE_QUERY_STATE = 0x0001 //col:520
SEMAPHORE_MODIFY_STATE = 0x0002 //col:524
SEMAPHORE_ALL_ACCESS = (SEMAPHORE_QUERY_STATE|SEMAPHORE_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:528
TIMER_QUERY_STATE = 0x0001 //col:585
TIMER_MODIFY_STATE = 0x0002 //col:589
TIMER_ALL_ACCESS = (TIMER_QUERY_STATE|TIMER_MODIFY_STATE|STANDARD_RIGHTS_REQUIRED|SYNCHRONIZE) //col:593
PROFILE_CONTROL = 0x0001 //col:761
PROFILE_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | PROFILE_CONTROL) //col:762
KEYEDEVENT_WAIT = 0x0001 //col:829
KEYEDEVENT_WAKE = 0x0002 //col:830
KEYEDEVENT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | KEYEDEVENT_WAIT | KEYEDEVENT_WAKE) //col:831
WORKER_FACTORY_RELEASE_WORKER = 0x0001 //col:1055
WORKER_FACTORY_WAIT = 0x0002 //col:1056
WORKER_FACTORY_SET_INFORMATION = 0x0004 //col:1057
WORKER_FACTORY_QUERY_INFORMATION = 0x0008 //col:1058
WORKER_FACTORY_READY_WORKER = 0x0010 //col:1059
WORKER_FACTORY_SHUTDOWN = 0x0020 //col:1060
WORKER_FACTORY_ALL_ACCESS ( = STANDARD_RIGHTS_REQUIRED | WORKER_FACTORY_RELEASE_WORKER | WORKER_FACTORY_WAIT | WORKER_FACTORY_SET_INFORMATION | WORKER_FACTORY_QUERY_INFORMATION | WORKER_FACTORY_READY_WORKER | WORKER_FACTORY_SHUTDOWN ) //col:1062
MM_WORKING_SET_MAX_HARD_ENABLE = 0x1 //col:1912
MM_WORKING_SET_MAX_HARD_DISABLE = 0x2 //col:1913
MM_WORKING_SET_MIN_HARD_ENABLE = 0x4 //col:1914
MM_WORKING_SET_MIN_HARD_DISABLE = 0x8 //col:1915
_TRACEHANDLE_DEFINED =  //col:2005
PERF_MASK_INDEX = (0xe0000000) //col:2047
PERF_MASK_GROUP = (~PERF_MASK_INDEX) //col:2048
PERF_NUM_MASKS = 8 //col:2049
PERF_GET_MASK_INDEX(GM) = (((GM) & PERF_MASK_INDEX) >> 29) //col:2051
PERF_GET_MASK_GROUP(GM) = ((GM) & PERF_MASK_GROUP) //col:2052
PERFINFO_OR_GROUP_WITH_GROUPMASK(Group, pGroupMask) = (pGroupMask)->Masks[PERF_GET_MASK_INDEX(Group)] |= PERF_GET_MASK_GROUP(Group); //col:2053
PERF_PROCESS =            EVENT_TRACE_FLAG_PROCESS //col:2057
PERF_THREAD =             EVENT_TRACE_FLAG_THREAD //col:2058
PERF_PROC_THREAD =        EVENT_TRACE_FLAG_PROCESS | EVENT_TRACE_FLAG_THREAD //col:2059
PERF_LOADER =             EVENT_TRACE_FLAG_IMAGE_LOAD //col:2060
PERF_PERF_COUNTER =       EVENT_TRACE_FLAG_PROCESS_COUNTERS //col:2061
PERF_FILENAME =           EVENT_TRACE_FLAG_DISK_FILE_IO //col:2062
PERF_DISK_IO =            EVENT_TRACE_FLAG_DISK_FILE_IO | EVENT_TRACE_FLAG_DISK_IO //col:2063
PERF_DISK_IO_INIT =       EVENT_TRACE_FLAG_DISK_IO_INIT //col:2064
PERF_ALL_FAULTS =         EVENT_TRACE_FLAG_MEMORY_PAGE_FAULTS //col:2065
PERF_HARD_FAULTS =        EVENT_TRACE_FLAG_MEMORY_HARD_FAULTS //col:2066
PERF_VAMAP =              EVENT_TRACE_FLAG_VAMAP //col:2067
PERF_NETWORK =            EVENT_TRACE_FLAG_NETWORK_TCPIP //col:2068
PERF_REGISTRY =           EVENT_TRACE_FLAG_REGISTRY //col:2069
PERF_DBGPRINT =           EVENT_TRACE_FLAG_DBGPRINT //col:2070
PERF_JOB =                EVENT_TRACE_FLAG_JOB //col:2071
PERF_ALPC =               EVENT_TRACE_FLAG_ALPC //col:2072
PERF_SPLIT_IO =           EVENT_TRACE_FLAG_SPLIT_IO //col:2073
PERF_DEBUG_EVENTS =       EVENT_TRACE_FLAG_DEBUG_EVENTS //col:2074
PERF_FILE_IO =            EVENT_TRACE_FLAG_FILE_IO //col:2075
PERF_FILE_IO_INIT =       EVENT_TRACE_FLAG_FILE_IO_INIT //col:2076
PERF_NO_SYSCONFIG =       EVENT_TRACE_FLAG_NO_SYSCONFIG //col:2077
PERF_MEMORY =             0x20000001 //col:2080
PERF_PROFILE =            0x20000002  // equivalent to EVENT_TRACE_FLAG_PROFILE //col:2081
PERF_CONTEXT_SWITCH =     0x20000004  // equivalent to EVENT_TRACE_FLAG_CSWITCH //col:2082
PERF_FOOTPRINT =          0x20000008 //col:2083
PERF_DRIVERS =            0x20000010  // equivalent to EVENT_TRACE_FLAG_DRIVER //col:2084
PERF_REFSET =             0x20000020 //col:2085
PERF_POOL =               0x20000040 //col:2086
PERF_POOLTRACE =          0x20000041 //col:2087
PERF_DPC =                0x20000080  // equivalent to EVENT_TRACE_FLAG_DPC //col:2088
PERF_COMPACT_CSWITCH =    0x20000100 //col:2089
PERF_DISPATCHER =         0x20000200  // equivalent to EVENT_TRACE_FLAG_DISPATCHER //col:2090
PERF_PMC_PROFILE =        0x20000400 //col:2091
PERF_PROFILING =          0x20000402 //col:2092
PERF_PROCESS_INSWAP =     0x20000800 //col:2093
PERF_AFFINITY =           0x20001000 //col:2094
PERF_PRIORITY =           0x20002000 //col:2095
PERF_INTERRUPT =          0x20004000  // equivalent to EVENT_TRACE_FLAG_INTERRUPT //col:2096
PERF_VIRTUAL_ALLOC =      0x20008000  // equivalent to EVENT_TRACE_FLAG_VIRTUAL_ALLOC //col:2097
PERF_SPINLOCK =           0x20010000 //col:2098
PERF_SYNC_OBJECTS =       0x20020000 //col:2099
PERF_DPC_QUEUE =          0x20040000 //col:2100
PERF_MEMINFO =            0x20080000 //col:2101
PERF_CONTMEM_GEN =        0x20100000 //col:2102
PERF_SPINLOCK_CNTRS =     0x20200000 //col:2103
PERF_SPININSTR =          0x20210000 //col:2104
PERF_SESSION =            0x20400000 //col:2105
PERF_PFSECTION =          0x20400000 //col:2106
PERF_MEMINFO_WS =         0x20800000 //col:2107
PERF_KERNEL_QUEUE =       0x21000000 //col:2108
PERF_INTERRUPT_STEER =    0x22000000 //col:2109
PERF_SHOULD_YIELD =       0x24000000 //col:2110
PERF_WS =                 0x28000000 //col:2111
PERF_ANTI_STARVATION =    0x40000001 //col:2114
PERF_PROCESS_FREEZE =     0x40000002 //col:2115
PERF_PFN_LIST =           0x40000004 //col:2116
PERF_WS_DETAIL =          0x40000008 //col:2117
PERF_WS_ENTRY =           0x40000010 //col:2118
PERF_HEAP =               0x40000020 //col:2119
PERF_SYSCALL =            0x40000040  // equivalent to EVENT_TRACE_FLAG_SYSTEMCALL //col:2120
PERF_UMS =                0x40000080 //col:2121
PERF_BACKTRACE =          0x40000100 //col:2122
PERF_VULCAN =             0x40000200 //col:2123
PERF_OBJECTS =            0x40000400 //col:2124
PERF_EVENTS =             0x40000800 //col:2125
PERF_FULLTRACE =          0x40001000 //col:2126
PERF_DFSS =               0x40002000 //col:2127
PERF_PREFETCH =           0x40004000 //col:2128
PERF_PROCESSOR_IDLE =     0x40008000 //col:2129
PERF_CPU_CONFIG =         0x40010000 //col:2130
PERF_TIMER =              0x40020000 //col:2131
PERF_CLOCK_INTERRUPT =    0x40040000 //col:2132
PERF_LOAD_BALANCER =      0x40080000 //col:2133
PERF_CLOCK_TIMER =        0x40100000 //col:2134
PERF_IDLE_SELECTION =     0x40200000 //col:2135
PERF_IPI =                0x40400000 //col:2136
PERF_IO_TIMER =           0x40800000 //col:2137
PERF_REG_HIVE =           0x41000000 //col:2138
PERF_REG_NOTIF =          0x42000000 //col:2139
PERF_PPM_EXIT_LATENCY =   0x44000000 //col:2140
PERF_WORKER_THREAD =      0x48000000 //col:2141
PERF_OPTICAL_IO =         0x80000001 //col:2144
PERF_OPTICAL_IO_INIT =    0x80000002 //col:2145
PERF_DLL_INFO =           0x80000008 //col:2146
PERF_DLL_FLUSH_WS =       0x80000010 //col:2147
PERF_OB_HANDLE =          0x80000040 //col:2148
PERF_OB_OBJECT =          0x80000080 //col:2149
PERF_WAKE_DROP =          0x80000200 //col:2150
PERF_WAKE_EVENT =         0x80000400 //col:2151
PERF_DEBUGGER =           0x80000800 //col:2152
PERF_PROC_ATTACH =        0x80001000 //col:2153
PERF_WAKE_COUNTER =       0x80002000 //col:2154
PERF_POWER =              0x80008000 //col:2155
PERF_SOFT_TRIM =          0x80010000 //col:2156
PERF_CC =                 0x80020000 //col:2157
PERF_FLT_IO_INIT =        0x80080000 //col:2158
PERF_FLT_IO =             0x80100000 //col:2159
PERF_FLT_FASTIO =         0x80200000 //col:2160
PERF_FLT_IO_FAILURE =     0x80400000 //col:2161
PERF_HV_PROFILE =         0x80800000 //col:2162
PERF_WDF_DPC =            0x81000000 //col:2163
PERF_WDF_INTERRUPT =      0x82000000 //col:2164
PERF_CACHE_FLUSH =        0x84000000 //col:2165
PERF_HIBER_RUNDOWN =      0xA0000001 //col:2168
PERF_SYSCFG_SYSTEM =      0xC0000001 //col:2171
PERF_SYSCFG_GRAPHICS =    0xC0000002 //col:2172
PERF_SYSCFG_STORAGE =     0xC0000004 //col:2173
PERF_SYSCFG_NETWORK =     0xC0000008 //col:2174
PERF_SYSCFG_SERVICES =    0xC0000010 //col:2175
PERF_SYSCFG_PNP =         0xC0000020 //col:2176
PERF_SYSCFG_OPTICAL =     0xC0000040 //col:2177
PERF_SYSCFG_ALL =         0xDFFFFFFF //col:2178
PERF_CLUSTER_OFF =        0xE0000001 //col:2181
PERF_MEMORY_CONTROL =     0xE0000002 //col:2182
MAXIMUM_NODE_COUNT = 0x40 //col:2508
MAXIMUM_NODE_COUNT = 0x10 //col:2510
CODEINTEGRITY_OPTION_ENABLED = 0x01 //col:2945
CODEINTEGRITY_OPTION_TESTSIGN = 0x02 //col:2946
CODEINTEGRITY_OPTION_UMCI_ENABLED = 0x04 //col:2947
CODEINTEGRITY_OPTION_UMCI_AUDITMODE_ENABLED = 0x08 //col:2948
CODEINTEGRITY_OPTION_UMCI_EXCLUSIONPATHS_ENABLED = 0x10 //col:2949
CODEINTEGRITY_OPTION_TEST_BUILD = 0x20 //col:2950
CODEINTEGRITY_OPTION_PREPRODUCTION_BUILD = 0x40 //col:2951
CODEINTEGRITY_OPTION_DEBUGMODE_ENABLED = 0x80 //col:2952
CODEINTEGRITY_OPTION_FLIGHT_BUILD = 0x100 //col:2953
CODEINTEGRITY_OPTION_FLIGHTING_ENABLED = 0x200 //col:2954
CODEINTEGRITY_OPTION_HVCI_KMCI_ENABLED = 0x400 //col:2955
CODEINTEGRITY_OPTION_HVCI_KMCI_AUDITMODE_ENABLED = 0x800 //col:2956
CODEINTEGRITY_OPTION_HVCI_KMCI_STRICTMODE_ENABLED = 0x1000 //col:2957
CODEINTEGRITY_OPTION_HVCI_IUM_ENABLED = 0x2000 //col:2958
CODEINTEGRITY_OPTION_WHQL_ENFORCEMENT_ENABLED = 0x4000 //col:2959
CODEINTEGRITY_OPTION_WHQL_AUDITMODE_ENABLED = 0x8000 //col:2960
SYSTEM_STORE_INFORMATION_VERSION = 1 //col:3026
SYSTEM_STORE_STATS_INFORMATION_VERSION = 2 //col:3037
SYSTEM_STORE_CREATE_INFORMATION_VERSION = 6 //col:3142
SYSTEM_STORE_DELETE_INFORMATION_VERSION = 1 //col:3224
SYSTEM_STORE_LIST_INFORMATION_VERSION = 2 //col:3233
SYSTEM_CACHE_LIST_INFORMATION_VERSION = 2 //col:3250
SYSTEM_CACHE_CREATE_INFORMATION_VERSION = 3 //col:3260
SYSTEM_CACHE_DELETE_INFORMATION_VERSION = 1 //col:3287
SYSTEM_CACHE_STORE_CREATE_INFORMATION_VERSION = 2 //col:3296
SYSTEM_CACHE_STORE_DELETE_INFORMATION_VERSION = 1 //col:3315
SYSTEM_CACHE_STATS_INFORMATION_VERSION = 3 //col:3326
SYSTEM_STORE_REGISTRATION_INFORMATION_VERSION = 2 //col:3352
SYSTEM_STORE_RESIZE_INFORMATION_VERSION = 6 //col:3366
SYSTEM_CACHE_STORE_RESIZE_INFORMATION_VERSION = 1 //col:3378
SYSTEM_STORE_CONFIG_INFORMATION_VERSION = 4 //col:3391
SYSTEM_STORE_HIGH_MEM_PRIORITY_INFORMATION_VERSION = 1 //col:3409
SYSTEM_STORE_TRIM_INFORMATION_VERSION = 1 //col:3420
SYSTEM_STORE_COMPRESSION_INFORMATION_VERSION = 3 //col:3431
MEMORY_COMBINE_FLAGS_COMMON_PAGES_ONLY = 0x4 //col:3686
SYSDBG_LIVEDUMP_CONTROL_VERSION = 1 //col:4846
SYSDBG_LIVEDUMP_CONTROL_VERSION_WIN11 = 2 //col:4847
HARDERROR_OVERRIDE_ERRORMODE = 0x10000000 //col:4913
PROCESSOR_FEATURE_MAX = 64 //col:4936
MAX_WOW64_SHARED_ENTRIES = 16 //col:4938
NX_SUPPORT_POLICY_ALWAYSOFF = 0 //col:4940
NX_SUPPORT_POLICY_ALWAYSON = 1 //col:4941
NX_SUPPORT_POLICY_OPTIN = 2 //col:4942
NX_SUPPORT_POLICY_OPTOUT = 3 //col:4943
USER_SHARED_DATA = ((KUSER_SHARED_DATA * const)0x7ffe0000) //col:5141
ATOM_FLAG_GLOBAL = 0x2 //col:5373
FLG_STOP_ON_EXCEPTION = 0x00000001 // uk //col:5437
FLG_SHOW_LDR_SNAPS = 0x00000002 // uk //col:5438
FLG_DEBUG_INITIAL_COMMAND = 0x00000004 // k //col:5439
FLG_STOP_ON_HUNG_GUI = 0x00000008 // k //col:5440
FLG_HEAP_ENABLE_TAIL_CHECK = 0x00000010 // u //col:5442
FLG_HEAP_ENABLE_FREE_CHECK = 0x00000020 // u //col:5443
FLG_HEAP_VALIDATE_PARAMETERS = 0x00000040 // u //col:5444
FLG_HEAP_VALIDATE_ALL = 0x00000080 // u //col:5445
FLG_APPLICATION_VERIFIER = 0x00000100 // u //col:5447
FLG_MONITOR_SILENT_PROCESS_EXIT = 0x00000200 // uk //col:5448
FLG_POOL_ENABLE_TAGGING = 0x00000400 // k //col:5449
FLG_HEAP_ENABLE_TAGGING = 0x00000800 // u //col:5450
FLG_USER_STACK_TRACE_DB = 0x00001000 // u,32 //col:5452
FLG_KERNEL_STACK_TRACE_DB = 0x00002000 // k,32 //col:5453
FLG_MAINTAIN_OBJECT_TYPELIST = 0x00004000 // k //col:5454
FLG_HEAP_ENABLE_TAG_BY_DLL = 0x00008000 // u //col:5455
FLG_DISABLE_STACK_EXTENSION = 0x00010000 // u //col:5457
FLG_ENABLE_CSRDEBUG = 0x00020000 // k //col:5458
FLG_ENABLE_KDEBUG_SYMBOL_LOAD = 0x00040000 // k //col:5459
FLG_DISABLE_PAGE_KERNEL_STACKS = 0x00080000 // k //col:5460
FLG_ENABLE_SYSTEM_CRIT_BREAKS = 0x00100000 // u //col:5462
FLG_HEAP_DISABLE_COALESCING = 0x00200000 // u //col:5463
FLG_ENABLE_CLOSE_EXCEPTIONS = 0x00400000 // k //col:5464
FLG_ENABLE_EXCEPTION_LOGGING = 0x00800000 // k //col:5465
FLG_ENABLE_HANDLE_TYPE_TAGGING = 0x01000000 // k //col:5467
FLG_HEAP_PAGE_ALLOCS = 0x02000000 // u //col:5468
FLG_DEBUG_INITIAL_COMMAND_EX = 0x04000000 // k //col:5469
FLG_DISABLE_DBGPRINT = 0x08000000 // k //col:5470
FLG_CRITSEC_EVENT_CREATION = 0x10000000 // u //col:5472
FLG_STOP_ON_UNHANDLED_EXCEPTION = 0x20000000 // u,64 //col:5473
FLG_ENABLE_HANDLE_EXCEPTIONS = 0x40000000 // k //col:5474
FLG_DISABLE_PROTDLLS = 0x80000000 // u //col:5475
FLG_VALID_BITS = 0xfffffdff //col:5477
FLG_USERMODE_VALID_BITS (FLG_STOP_ON_EXCEPTION | = FLG_SHOW_LDR_SNAPS | FLG_HEAP_ENABLE_TAIL_CHECK | FLG_HEAP_ENABLE_FREE_CHECK | FLG_HEAP_VALIDATE_PARAMETERS | FLG_HEAP_VALIDATE_ALL | FLG_APPLICATION_VERIFIER | FLG_HEAP_ENABLE_TAGGING | FLG_USER_STACK_TRACE_DB | FLG_HEAP_ENABLE_TAG_BY_DLL | FLG_DISABLE_STACK_EXTENSION | FLG_ENABLE_SYSTEM_CRIT_BREAKS | FLG_HEAP_DISABLE_COALESCING | FLG_DISABLE_PROTDLLS | FLG_HEAP_PAGE_ALLOCS | FLG_CRITSEC_EVENT_CREATION | FLG_LDR_TOP_DOWN) //col:5479
FLG_BOOTONLY_VALID_BITS (FLG_KERNEL_STACK_TRACE_DB | = FLG_MAINTAIN_OBJECT_TYPELIST | FLG_ENABLE_CSRDEBUG | FLG_DEBUG_INITIAL_COMMAND | FLG_DEBUG_INITIAL_COMMAND_EX | FLG_DISABLE_PAGE_KERNEL_STACKS) //col:5497
FLG_KERNELMODE_VALID_BITS (FLG_STOP_ON_EXCEPTION | = FLG_SHOW_LDR_SNAPS | FLG_STOP_ON_HUNG_GUI | FLG_POOL_ENABLE_TAGGING | FLG_ENABLE_KDEBUG_SYMBOL_LOAD | FLG_ENABLE_CLOSE_EXCEPTIONS | FLG_ENABLE_EXCEPTION_LOGGING | FLG_ENABLE_HANDLE_TYPE_TAGGING | FLG_DISABLE_DBGPRINT | FLG_ENABLE_HANDLE_EXCEPTIONS) //col:5504
)

const(
    FilterBootOptionOperationOpenSystemStore = 1  //col:272
    FilterBootOptionOperationSetElement = 2  //col:273
    FilterBootOptionOperationDeleteElement = 3  //col:274
    FilterBootOptionOperationMax = 4  //col:275
)


const(
    EventBasicInformation = 1  //col:309
)


const(
    MutantBasicInformation // MUTANT_BASIC_INFORMATION = 1  //col:463
    MutantOwnerInformation // MUTANT_OWNER_INFORMATION = 2  //col:464
)


const(
    SemaphoreBasicInformation = 1  //col:533
)


const(
    TimerBasicInformation // TIMER_BASIC_INFORMATION = 1  //col:598
)


const(
    TimerSetCoalescableTimer // TIMER_SET_COALESCABLE_TIMER_INFO = 1  //col:615
    MaxTimerInfoClass = 2  //col:616
)


const(
    WnfWellKnownStateName = 1  //col:897
    WnfPermanentStateName = 2  //col:898
    WnfPersistentStateName = 3  //col:899
    WnfTemporaryStateName = 4  //col:900
)


const(
    WnfInfoStateNameExist = 1  //col:905
    WnfInfoSubscribersPresent = 2  //col:906
    WnfInfoIsQuiescent = 3  //col:907
)


const(
    WnfDataScopeSystem = 1  //col:912
    WnfDataScopeSession = 2  //col:913
    WnfDataScopeUser = 3  //col:914
    WnfDataScopeProcess = 4  //col:915
    WnfDataScopeMachine // REDSTONE3 = 5  //col:916
    WnfDataScopePhysicalMachine // WIN11 = 6  //col:917
)


const(
    WorkerFactoryTimeout // LARGE_INTEGER = 1  //col:1078
    WorkerFactoryRetryTimeout // LARGE_INTEGER = 2  //col:1079
    WorkerFactoryIdleTimeout // s: LARGE_INTEGER = 3  //col:1080
    WorkerFactoryBindingCount // s: ULONG = 4  //col:1081
    WorkerFactoryThreadMinimum // s: ULONG = 5  //col:1082
    WorkerFactoryThreadMaximum // s: ULONG = 6  //col:1083
    WorkerFactoryPaused // ULONG or BOOLEAN = 7  //col:1084
    WorkerFactoryBasicInformation // q: WORKER_FACTORY_BASIC_INFORMATION = 8  //col:1085
    WorkerFactoryAdjustThreadGoal = 9  //col:1086
    WorkerFactoryCallbackType = 10  //col:1087
    WorkerFactoryStackInformation // 10 = 11  //col:1088
    WorkerFactoryThreadBasePriority // s: ULONG = 12  //col:1089
    WorkerFactoryTimeoutWaiters // s: ULONG since THRESHOLD = 13  //col:1090
    WorkerFactoryFlags // s: ULONG = 14  //col:1091
    WorkerFactoryThreadSoftMaximum // s: ULONG = 15  //col:1092
    WorkerFactoryThreadCpuSets // since REDSTONE5 = 16  //col:1093
    MaxWorkerFactoryInfoClass = 17  //col:1094
)


const(
    SystemBasicInformation // q: SYSTEM_BASIC_INFORMATION = 1  //col:1306
    SystemProcessorInformation // q: SYSTEM_PROCESSOR_INFORMATION = 2  //col:1307
    SystemPerformanceInformation // q: SYSTEM_PERFORMANCE_INFORMATION = 3  //col:1308
    SystemTimeOfDayInformation // q: SYSTEM_TIMEOFDAY_INFORMATION = 4  //col:1309
    SystemPathInformation // not implemented = 5  //col:1310
    SystemProcessInformation // q: SYSTEM_PROCESS_INFORMATION = 6  //col:1311
    SystemCallCountInformation // q: SYSTEM_CALL_COUNT_INFORMATION = 7  //col:1312
    SystemDeviceInformation // q: SYSTEM_DEVICE_INFORMATION = 8  //col:1313
    SystemProcessorPerformanceInformation // q: SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION (EX in: USHORT ProcessorGroup) = 9  //col:1314
    SystemFlagsInformation // q: SYSTEM_FLAGS_INFORMATION = 10  //col:1315
    SystemCallTimeInformation // not implemented // SYSTEM_CALL_TIME_INFORMATION // 10 = 11  //col:1316
    SystemModuleInformation // q: RTL_PROCESS_MODULES = 12  //col:1317
    SystemLocksInformation // q: RTL_PROCESS_LOCKS = 13  //col:1318
    SystemStackTraceInformation // q: RTL_PROCESS_BACKTRACES = 14  //col:1319
    SystemPagedPoolInformation // not implemented = 15  //col:1320
    SystemNonPagedPoolInformation // not implemented = 16  //col:1321
    SystemHandleInformation // q: SYSTEM_HANDLE_INFORMATION = 17  //col:1322
    SystemObjectInformation // q: SYSTEM_OBJECTTYPE_INFORMATION mixed with SYSTEM_OBJECT_INFORMATION = 18  //col:1323
    SystemPageFileInformation // q: SYSTEM_PAGEFILE_INFORMATION = 19  //col:1324
    SystemVdmInstemulInformation // q: SYSTEM_VDM_INSTEMUL_INFO = 20  //col:1325
    SystemVdmBopInformation // not implemented // 20 = 21  //col:1326
    SystemFileCacheInformation // q: SYSTEM_FILECACHE_INFORMATION; s (requires SeIncreaseQuotaPrivilege) (info for WorkingSetTypeSystemCache) = 22  //col:1327
    SystemPoolTagInformation // q: SYSTEM_POOLTAG_INFORMATION = 23  //col:1328
    SystemInterruptInformation // q: SYSTEM_INTERRUPT_INFORMATION (EX in: USHORT ProcessorGroup) = 24  //col:1329
    SystemDpcBehaviorInformation // q: SYSTEM_DPC_BEHAVIOR_INFORMATION; s: SYSTEM_DPC_BEHAVIOR_INFORMATION (requires SeLoadDriverPrivilege) = 25  //col:1330
    SystemFullMemoryInformation // not implemented // SYSTEM_MEMORY_USAGE_INFORMATION = 26  //col:1331
    SystemLoadGdiDriverInformation // s (kernel-mode only) = 27  //col:1332
    SystemUnloadGdiDriverInformation // s (kernel-mode only) = 28  //col:1333
    SystemTimeAdjustmentInformation // q: SYSTEM_QUERY_TIME_ADJUST_INFORMATION; s: SYSTEM_SET_TIME_ADJUST_INFORMATION (requires SeSystemtimePrivilege) = 29  //col:1334
    SystemSummaryMemoryInformation // not implemented // SYSTEM_MEMORY_USAGE_INFORMATION = 30  //col:1335
    SystemMirrorMemoryInformation // s (requires license value "Kernel-MemoryMirroringSupported") (requires SeShutdownPrivilege) // 30 = 31  //col:1336
    SystemPerformanceTraceInformation // q; s: (type depends on EVENT_TRACE_INFORMATION_CLASS) = 32  //col:1337
    SystemObsolete0 // not implemented = 33  //col:1338
    SystemExceptionInformation // q: SYSTEM_EXCEPTION_INFORMATION = 34  //col:1339
    SystemCrashDumpStateInformation // s: SYSTEM_CRASH_DUMP_STATE_INFORMATION (requires SeDebugPrivilege) = 35  //col:1340
    SystemKernelDebuggerInformation // q: SYSTEM_KERNEL_DEBUGGER_INFORMATION = 36  //col:1341
    SystemContextSwitchInformation // q: SYSTEM_CONTEXT_SWITCH_INFORMATION = 37  //col:1342
    SystemRegistryQuotaInformation // q: SYSTEM_REGISTRY_QUOTA_INFORMATION; s (requires SeIncreaseQuotaPrivilege) = 38  //col:1343
    SystemExtendServiceTableInformation // s (requires SeLoadDriverPrivilege) // loads win32k only = 39  //col:1344
    SystemPrioritySeperation // s (requires SeTcbPrivilege) = 40  //col:1345
    SystemVerifierAddDriverInformation // s (requires SeDebugPrivilege) // 40 = 41  //col:1346
    SystemVerifierRemoveDriverInformation // s (requires SeDebugPrivilege) = 42  //col:1347
    SystemProcessorIdleInformation // q: SYSTEM_PROCESSOR_IDLE_INFORMATION (EX in: USHORT ProcessorGroup) = 43  //col:1348
    SystemLegacyDriverInformation // q: SYSTEM_LEGACY_DRIVER_INFORMATION = 44  //col:1349
    SystemCurrentTimeZoneInformation // q; s: RTL_TIME_ZONE_INFORMATION = 45  //col:1350
    SystemLookasideInformation // q: SYSTEM_LOOKASIDE_INFORMATION = 46  //col:1351
    SystemTimeSlipNotification // s: HANDLE (NtCreateEvent) (requires SeSystemtimePrivilege) = 47  //col:1352
    SystemSessionCreate // not implemented = 48  //col:1353
    SystemSessionDetach // not implemented = 49  //col:1354
    SystemSessionInformation // not implemented (SYSTEM_SESSION_INFORMATION) = 50  //col:1355
    SystemRangeStartInformation // q: SYSTEM_RANGE_START_INFORMATION // 50 = 51  //col:1356
    SystemVerifierInformation // q: SYSTEM_VERIFIER_INFORMATION; s (requires SeDebugPrivilege) = 52  //col:1357
    SystemVerifierThunkExtend // s (kernel-mode only) = 53  //col:1358
    SystemSessionProcessInformation // q: SYSTEM_SESSION_PROCESS_INFORMATION = 54  //col:1359
    SystemLoadGdiDriverInSystemSpace // s: SYSTEM_GDI_DRIVER_INFORMATION (kernel-mode only) (same as SystemLoadGdiDriverInformation) = 55  //col:1360
    SystemNumaProcessorMap // q: SYSTEM_NUMA_INFORMATION = 56  //col:1361
    SystemPrefetcherInformation // q; s: PREFETCHER_INFORMATION // PfSnQueryPrefetcherInformation = 57  //col:1362
    SystemExtendedProcessInformation // q: SYSTEM_PROCESS_INFORMATION = 58  //col:1363
    SystemRecommendedSharedDataAlignment // q: ULONG // KeGetRecommendedSharedDataAlignment = 59  //col:1364
    SystemComPlusPackage // q; s: ULONG = 60  //col:1365
    SystemNumaAvailableMemory // q: SYSTEM_NUMA_INFORMATION // 60 = 61  //col:1366
    SystemProcessorPowerInformation // q: SYSTEM_PROCESSOR_POWER_INFORMATION (EX in: USHORT ProcessorGroup) = 62  //col:1367
    SystemEmulationBasicInformation // q: SYSTEM_BASIC_INFORMATION = 63  //col:1368
    SystemEmulationProcessorInformation // q: SYSTEM_PROCESSOR_INFORMATION = 64  //col:1369
    SystemExtendedHandleInformation // q: SYSTEM_HANDLE_INFORMATION_EX = 65  //col:1370
    SystemLostDelayedWriteInformation // q: ULONG = 66  //col:1371
    SystemBigPoolInformation // q: SYSTEM_BIGPOOL_INFORMATION = 67  //col:1372
    SystemSessionPoolTagInformation // q: SYSTEM_SESSION_POOLTAG_INFORMATION = 68  //col:1373
    SystemSessionMappedViewInformation // q: SYSTEM_SESSION_MAPPED_VIEW_INFORMATION = 69  //col:1374
    SystemHotpatchInformation // q; s: SYSTEM_HOTPATCH_CODE_INFORMATION = 70  //col:1375
    SystemObjectSecurityMode // q: ULONG // 70 = 71  //col:1376
    SystemWatchdogTimerHandler // s: SYSTEM_WATCHDOG_HANDLER_INFORMATION // (kernel-mode only) = 72  //col:1377
    SystemWatchdogTimerInformation // q: SYSTEM_WATCHDOG_TIMER_INFORMATION // (kernel-mode only) = 73  //col:1378
    SystemLogicalProcessorInformation // q: SYSTEM_LOGICAL_PROCESSOR_INFORMATION (EX in: USHORT ProcessorGroup) = 74  //col:1379
    SystemWow64SharedInformationObsolete // not implemented = 75  //col:1380
    SystemRegisterFirmwareTableInformationHandler // s: SYSTEM_FIRMWARE_TABLE_HANDLER // (kernel-mode only) = 76  //col:1381
    SystemFirmwareTableInformation // SYSTEM_FIRMWARE_TABLE_INFORMATION = 77  //col:1382
    SystemModuleInformationEx // q: RTL_PROCESS_MODULE_INFORMATION_EX = 78  //col:1383
    SystemVerifierTriageInformation // not implemented = 79  //col:1384
    SystemSuperfetchInformation // q; s: SUPERFETCH_INFORMATION // PfQuerySuperfetchInformation = 80  //col:1385
    SystemMemoryListInformation // q: SYSTEM_MEMORY_LIST_INFORMATION; s: SYSTEM_MEMORY_LIST_COMMAND (requires SeProfileSingleProcessPrivilege) // 80 = 81  //col:1386
    SystemFileCacheInformationEx // q: SYSTEM_FILECACHE_INFORMATION; s (requires SeIncreaseQuotaPrivilege) (same as SystemFileCacheInformation) = 82  //col:1387
    SystemThreadPriorityClientIdInformation // s: SYSTEM_THREAD_CID_PRIORITY_INFORMATION (requires SeIncreaseBasePriorityPrivilege) = 83  //col:1388
    SystemProcessorIdleCycleTimeInformation // q: SYSTEM_PROCESSOR_IDLE_CYCLE_TIME_INFORMATION[] (EX in: USHORT ProcessorGroup) = 84  //col:1389
    SystemVerifierCancellationInformation // SYSTEM_VERIFIER_CANCELLATION_INFORMATION // name:wow64:whNT32QuerySystemVerifierCancellationInformation = 85  //col:1390
    SystemProcessorPowerInformationEx // not implemented = 86  //col:1391
    SystemRefTraceInformation // q; s: SYSTEM_REF_TRACE_INFORMATION // ObQueryRefTraceInformation = 87  //col:1392
    SystemSpecialPoolInformation // q; s: SYSTEM_SPECIAL_POOL_INFORMATION (requires SeDebugPrivilege) // MmSpecialPoolTag then MmSpecialPoolCatchOverruns ! =  0  //col:1393
    SystemProcessIdInformation // q: SYSTEM_PROCESS_ID_INFORMATION = 89  //col:1394
    SystemErrorPortInformation // s (requires SeTcbPrivilege) = 90  //col:1395
    SystemBootEnvironmentInformation // q: SYSTEM_BOOT_ENVIRONMENT_INFORMATION // 90 = 91  //col:1396
    SystemHypervisorInformation // q: SYSTEM_HYPERVISOR_QUERY_INFORMATION = 92  //col:1397
    SystemVerifierInformationEx // q; s: SYSTEM_VERIFIER_INFORMATION_EX = 93  //col:1398
    SystemTimeZoneInformation // q; s: RTL_TIME_ZONE_INFORMATION (requires SeTimeZonePrivilege) = 94  //col:1399
    SystemImageFileExecutionOptionsInformation // s: SYSTEM_IMAGE_FILE_EXECUTION_OPTIONS_INFORMATION (requires SeTcbPrivilege) = 95  //col:1400
    SystemCoverageInformation // q: COVERAGE_MODULES s: COVERAGE_MODULE_REQUEST // ExpCovQueryInformation (requires SeDebugPrivilege) = 96  //col:1401
    SystemPrefetchPatchInformation // SYSTEM_PREFETCH_PATCH_INFORMATION = 97  //col:1402
    SystemVerifierFaultsInformation // s: SYSTEM_VERIFIER_FAULTS_INFORMATION (requires SeDebugPrivilege) = 98  //col:1403
    SystemSystemPartitionInformation // q: SYSTEM_SYSTEM_PARTITION_INFORMATION = 99  //col:1404
    SystemSystemDiskInformation // q: SYSTEM_SYSTEM_DISK_INFORMATION = 100  //col:1405
    SystemProcessorPerformanceDistribution // q: SYSTEM_PROCESSOR_PERFORMANCE_DISTRIBUTION (EX in: USHORT ProcessorGroup) // 100 = 101  //col:1406
    SystemNumaProximityNodeInformation // q; s: SYSTEM_NUMA_PROXIMITY_MAP = 102  //col:1407
    SystemDynamicTimeZoneInformation // q; s: RTL_DYNAMIC_TIME_ZONE_INFORMATION (requires SeTimeZonePrivilege) = 103  //col:1408
    SystemCodeIntegrityInformation // q: SYSTEM_CODEINTEGRITY_INFORMATION // SeCodeIntegrityQueryInformation = 104  //col:1409
    SystemProcessorMicrocodeUpdateInformation // s: SYSTEM_PROCESSOR_MICROCODE_UPDATE_INFORMATION = 105  //col:1410
    SystemProcessorBrandString // q: CHAR[] // HaliQuerySystemInformation -> HalpGetProcessorBrandString info class 23 = 106  //col:1411
    SystemVirtualAddressInformation // q: SYSTEM_VA_LIST_INFORMATION[]; s: SYSTEM_VA_LIST_INFORMATION[] (requires SeIncreaseQuotaPrivilege) // MmQuerySystemVaInformation = 107  //col:1412
    SystemLogicalProcessorAndGroupInformation // q: SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX (EX in: LOGICAL_PROCESSOR_RELATIONSHIP RelationshipType) // since WIN7 // KeQueryLogicalProcessorRelationship = 108  //col:1413
    SystemProcessorCycleTimeInformation // q: SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION[] (EX in: USHORT ProcessorGroup) = 109  //col:1414
    SystemStoreInformation // q; s: SYSTEM_STORE_INFORMATION (requires SeProfileSingleProcessPrivilege) // SmQueryStoreInformation = 110  //col:1415
    SystemRegistryAppendString // s: SYSTEM_REGISTRY_APPEND_STRING_PARAMETERS // 110 = 111  //col:1416
    SystemAitSamplingValue // s: ULONG (requires SeProfileSingleProcessPrivilege) = 112  //col:1417
    SystemVhdBootInformation // q: SYSTEM_VHD_BOOT_INFORMATION = 113  //col:1418
    SystemCpuQuotaInformation // q; s: PS_CPU_QUOTA_QUERY_INFORMATION = 114  //col:1419
    SystemNativeBasicInformation // q: SYSTEM_BASIC_INFORMATION = 115  //col:1420
    SystemErrorPortTimeouts // SYSTEM_ERROR_PORT_TIMEOUTS = 116  //col:1421
    SystemLowPriorityIoInformation // q: SYSTEM_LOW_PRIORITY_IO_INFORMATION = 117  //col:1422
    SystemTpmBootEntropyInformation // q: TPM_BOOT_ENTROPY_NT_RESULT // ExQueryTpmBootEntropyInformation = 118  //col:1423
    SystemVerifierCountersInformation // q: SYSTEM_VERIFIER_COUNTERS_INFORMATION = 119  //col:1424
    SystemPagedPoolInformationEx // q: SYSTEM_FILECACHE_INFORMATION; s (requires SeIncreaseQuotaPrivilege) (info for WorkingSetTypePagedPool) = 120  //col:1425
    SystemSystemPtesInformationEx // q: SYSTEM_FILECACHE_INFORMATION; s (requires SeIncreaseQuotaPrivilege) (info for WorkingSetTypeSystemPtes) // 120 = 121  //col:1426
    SystemNodeDistanceInformation // q: USHORT[4*NumaNodes] // (EX in: USHORT NodeNumber) = 122  //col:1427
    SystemAcpiAuditInformation // q: SYSTEM_ACPI_AUDIT_INFORMATION // HaliQuerySystemInformation -> HalpAuditQueryResults info class 26 = 123  //col:1428
    SystemBasicPerformanceInformation // q: SYSTEM_BASIC_PERFORMANCE_INFORMATION // name:wow64:whNtQuerySystemInformation_SystemBasicPerformanceInformation = 124  //col:1429
    SystemQueryPerformanceCounterInformation // q: SYSTEM_QUERY_PERFORMANCE_COUNTER_INFORMATION // since WIN7 SP1 = 125  //col:1430
    SystemSessionBigPoolInformation // q: SYSTEM_SESSION_POOLTAG_INFORMATION // since WIN8 = 126  //col:1431
    SystemBootGraphicsInformation // q; s: SYSTEM_BOOT_GRAPHICS_INFORMATION (kernel-mode only) = 127  //col:1432
    SystemScrubPhysicalMemoryInformation // q; s: MEMORY_SCRUB_INFORMATION = 128  //col:1433
    SystemBadPageInformation = 129  //col:1434
    SystemProcessorProfileControlArea // q; s: SYSTEM_PROCESSOR_PROFILE_CONTROL_AREA = 130  //col:1435
    SystemCombinePhysicalMemoryInformation // s: MEMORY_COMBINE_INFORMATION MEMORY_COMBINE_INFORMATION_EX MEMORY_COMBINE_INFORMATION_EX2 // 130 = 131  //col:1436
    SystemEntropyInterruptTimingInformation // q; s: SYSTEM_ENTROPY_TIMING_INFORMATION = 132  //col:1437
    SystemConsoleInformation // q: SYSTEM_CONSOLE_INFORMATION = 133  //col:1438
    SystemPlatformBinaryInformation // q: SYSTEM_PLATFORM_BINARY_INFORMATION (requires SeTcbPrivilege) = 134  //col:1439
    SystemPolicyInformation // q: SYSTEM_POLICY_INFORMATION = 135  //col:1440
    SystemHypervisorProcessorCountInformation // q: SYSTEM_HYPERVISOR_PROCESSOR_COUNT_INFORMATION = 136  //col:1441
    SystemDeviceDataInformation // q: SYSTEM_DEVICE_DATA_INFORMATION = 137  //col:1442
    SystemDeviceDataEnumerationInformation // q: SYSTEM_DEVICE_DATA_INFORMATION = 138  //col:1443
    SystemMemoryTopologyInformation // q: SYSTEM_MEMORY_TOPOLOGY_INFORMATION = 139  //col:1444
    SystemMemoryChannelInformation // q: SYSTEM_MEMORY_CHANNEL_INFORMATION = 140  //col:1445
    SystemBootLogoInformation // q: SYSTEM_BOOT_LOGO_INFORMATION // 140 = 141  //col:1446
    SystemProcessorPerformanceInformationEx // q: SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION_EX // (EX in: USHORT ProcessorGroup) // since WINBLUE = 142  //col:1447
    SystemCriticalProcessErrorLogInformation = 143  //col:1448
    SystemSecureBootPolicyInformation // q: SYSTEM_SECUREBOOT_POLICY_INFORMATION = 144  //col:1449
    SystemPageFileInformationEx // q: SYSTEM_PAGEFILE_INFORMATION_EX = 145  //col:1450
    SystemSecureBootInformation // q: SYSTEM_SECUREBOOT_INFORMATION = 146  //col:1451
    SystemEntropyInterruptTimingRawInformation = 147  //col:1452
    SystemPortableWorkspaceEfiLauncherInformation // q: SYSTEM_PORTABLE_WORKSPACE_EFI_LAUNCHER_INFORMATION = 148  //col:1453
    SystemFullProcessInformation // q: SYSTEM_PROCESS_INFORMATION with SYSTEM_PROCESS_INFORMATION_EXTENSION (requires admin) = 149  //col:1454
    SystemKernelDebuggerInformationEx // q: SYSTEM_KERNEL_DEBUGGER_INFORMATION_EX = 150  //col:1455
    SystemBootMetadataInformation // 150 = 151  //col:1456
    SystemSoftRebootInformation // q: ULONG = 152  //col:1457
    SystemElamCertificateInformation // s: SYSTEM_ELAM_CERTIFICATE_INFORMATION = 153  //col:1458
    SystemOfflineDumpConfigInformation // q: OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V2 = 154  //col:1459
    SystemProcessorFeaturesInformation // q: SYSTEM_PROCESSOR_FEATURES_INFORMATION = 155  //col:1460
    SystemRegistryReconciliationInformation // s: NULL (requires admin) (flushes registry hives) = 156  //col:1461
    SystemEdidInformation // q: SYSTEM_EDID_INFORMATION = 157  //col:1462
    SystemManufacturingInformation // q: SYSTEM_MANUFACTURING_INFORMATION // since THRESHOLD = 158  //col:1463
    SystemEnergyEstimationConfigInformation // q: SYSTEM_ENERGY_ESTIMATION_CONFIG_INFORMATION = 159  //col:1464
    SystemHypervisorDetailInformation // q: SYSTEM_HYPERVISOR_DETAIL_INFORMATION = 160  //col:1465
    SystemProcessorCycleStatsInformation // q: SYSTEM_PROCESSOR_CYCLE_STATS_INFORMATION (EX in: USHORT ProcessorGroup) // 160 = 161  //col:1466
    SystemVmGenerationCountInformation = 162  //col:1467
    SystemTrustedPlatformModuleInformation // q: SYSTEM_TPM_INFORMATION = 163  //col:1468
    SystemKernelDebuggerFlags // SYSTEM_KERNEL_DEBUGGER_FLAGS = 164  //col:1469
    SystemCodeIntegrityPolicyInformation // q: SYSTEM_CODEINTEGRITYPOLICY_INFORMATION = 165  //col:1470
    SystemIsolatedUserModeInformation // q: SYSTEM_ISOLATED_USER_MODE_INFORMATION = 166  //col:1471
    SystemHardwareSecurityTestInterfaceResultsInformation = 167  //col:1472
    SystemSingleModuleInformation // q: SYSTEM_SINGLE_MODULE_INFORMATION = 168  //col:1473
    SystemAllowedCpuSetsInformation = 169  //col:1474
    SystemVsmProtectionInformation // q: SYSTEM_VSM_PROTECTION_INFORMATION (previously SystemDmaProtectionInformation) = 170  //col:1475
    SystemInterruptCpuSetsInformation // q: SYSTEM_INTERRUPT_CPU_SET_INFORMATION // 170 = 171  //col:1476
    SystemSecureBootPolicyFullInformation // q: SYSTEM_SECUREBOOT_POLICY_FULL_INFORMATION = 172  //col:1477
    SystemCodeIntegrityPolicyFullInformation = 173  //col:1478
    SystemAffinitizedInterruptProcessorInformation // (requires SeIncreaseBasePriorityPrivilege) = 174  //col:1479
    SystemRootSiloInformation // q: SYSTEM_ROOT_SILO_INFORMATION = 175  //col:1480
    SystemCpuSetInformation // q: SYSTEM_CPU_SET_INFORMATION // since THRESHOLD2 = 176  //col:1481
    SystemCpuSetTagInformation // q: SYSTEM_CPU_SET_TAG_INFORMATION = 177  //col:1482
    SystemWin32WerStartCallout = 178  //col:1483
    SystemSecureKernelProfileInformation // q: SYSTEM_SECURE_KERNEL_HYPERGUARD_PROFILE_INFORMATION = 179  //col:1484
    SystemCodeIntegrityPlatformManifestInformation // q: SYSTEM_SECUREBOOT_PLATFORM_MANIFEST_INFORMATION // since REDSTONE = 180  //col:1485
    SystemInterruptSteeringInformation // SYSTEM_INTERRUPT_STEERING_INFORMATION_INPUT // 180 = 181  //col:1486
    SystemSupportedProcessorArchitectures // p: in opt: HANDLE out: SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION[] // NtQuerySystemInformationEx = 182  //col:1487
    SystemMemoryUsageInformation // q: SYSTEM_MEMORY_USAGE_INFORMATION = 183  //col:1488
    SystemCodeIntegrityCertificateInformation // q: SYSTEM_CODEINTEGRITY_CERTIFICATE_INFORMATION = 184  //col:1489
    SystemPhysicalMemoryInformation // q: SYSTEM_PHYSICAL_MEMORY_INFORMATION // since REDSTONE2 = 185  //col:1490
    SystemControlFlowTransition = 186  //col:1491
    SystemKernelDebuggingAllowed // s: ULONG = 187  //col:1492
    SystemActivityModerationExeState // SYSTEM_ACTIVITY_MODERATION_EXE_STATE = 188  //col:1493
    SystemActivityModerationUserSettings // SYSTEM_ACTIVITY_MODERATION_USER_SETTINGS = 189  //col:1494
    SystemCodeIntegrityPoliciesFullInformation = 190  //col:1495
    SystemCodeIntegrityUnlockInformation // SYSTEM_CODEINTEGRITY_UNLOCK_INFORMATION // 190 = 191  //col:1496
    SystemIntegrityQuotaInformation = 192  //col:1497
    SystemFlushInformation // q: SYSTEM_FLUSH_INFORMATION = 193  //col:1498
    SystemProcessorIdleMaskInformation // q: ULONG_PTR[ActiveGroupCount] // since REDSTONE3 = 194  //col:1499
    SystemSecureDumpEncryptionInformation = 195  //col:1500
    SystemWriteConstraintInformation // SYSTEM_WRITE_CONSTRAINT_INFORMATION = 196  //col:1501
    SystemKernelVaShadowInformation // SYSTEM_KERNEL_VA_SHADOW_INFORMATION = 197  //col:1502
    SystemHypervisorSharedPageInformation // SYSTEM_HYPERVISOR_SHARED_PAGE_INFORMATION // since REDSTONE4 = 198  //col:1503
    SystemFirmwareBootPerformanceInformation = 199  //col:1504
    SystemCodeIntegrityVerificationInformation // SYSTEM_CODEINTEGRITYVERIFICATION_INFORMATION = 200  //col:1505
    SystemFirmwarePartitionInformation // SYSTEM_FIRMWARE_PARTITION_INFORMATION // 200 = 201  //col:1506
    SystemSpeculationControlInformation // SYSTEM_SPECULATION_CONTROL_INFORMATION // (CVE-2017-5715) REDSTONE3 and above. = 202  //col:1507
    SystemDmaGuardPolicyInformation // SYSTEM_DMA_GUARD_POLICY_INFORMATION = 203  //col:1508
    SystemEnclaveLaunchControlInformation // SYSTEM_ENCLAVE_LAUNCH_CONTROL_INFORMATION = 204  //col:1509
    SystemWorkloadAllowedCpuSetsInformation // SYSTEM_WORKLOAD_ALLOWED_CPU_SET_INFORMATION // since REDSTONE5 = 205  //col:1510
    SystemCodeIntegrityUnlockModeInformation = 206  //col:1511
    SystemLeapSecondInformation // SYSTEM_LEAP_SECOND_INFORMATION = 207  //col:1512
    SystemFlags2Information // q: SYSTEM_FLAGS_INFORMATION = 208  //col:1513
    SystemSecurityModelInformation // SYSTEM_SECURITY_MODEL_INFORMATION // since 19H1 = 209  //col:1514
    SystemCodeIntegritySyntheticCacheInformation = 210  //col:1515
    SystemFeatureConfigurationInformation // SYSTEM_FEATURE_CONFIGURATION_INFORMATION // since 20H1 // 210 = 211  //col:1516
    SystemFeatureConfigurationSectionInformation // SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION = 212  //col:1517
    SystemFeatureUsageSubscriptionInformation // SYSTEM_FEATURE_USAGE_SUBSCRIPTION_DETAILS = 213  //col:1518
    SystemSecureSpeculationControlInformation // SECURE_SPECULATION_CONTROL_INFORMATION = 214  //col:1519
    SystemSpacesBootInformation // since 20H2 = 215  //col:1520
    SystemFwRamdiskInformation // SYSTEM_FIRMWARE_RAMDISK_INFORMATION = 216  //col:1521
    SystemWheaIpmiHardwareInformation = 217  //col:1522
    SystemDifSetRuleClassInformation = 218  //col:1523
    SystemDifClearRuleClassInformation = 219  //col:1524
    SystemDifApplyPluginVerificationOnDriver = 220  //col:1525
    SystemDifRemovePluginVerificationOnDriver // 220 = 221  //col:1526
    SystemShadowStackInformation // SYSTEM_SHADOW_STACK_INFORMATION = 222  //col:1527
    SystemBuildVersionInformation // SYSTEM_BUILD_VERSION_INFORMATION = 223  //col:1528
    SystemPoolLimitInformation // SYSTEM_POOL_LIMIT_INFORMATION = 224  //col:1529
    SystemCodeIntegrityAddDynamicStore = 225  //col:1530
    SystemCodeIntegrityClearDynamicStores = 226  //col:1531
    SystemDifPoolTrackingInformation = 227  //col:1532
    SystemPoolZeroingInformation // SYSTEM_POOL_ZEROING_INFORMATION = 228  //col:1533
    SystemDpcWatchdogInformation = 229  //col:1534
    SystemDpcWatchdogInformation2 = 230  //col:1535
    SystemSupportedProcessorArchitectures2 // q: in opt: HANDLE out: SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION[] // NtQuerySystemInformationEx  // 230 = 231  //col:1536
    SystemSingleProcessorRelationshipInformation // q: SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX // (EX in: PROCESSOR_NUMBER Processor) = 232  //col:1537
    SystemXfgCheckFailureInformation = 233  //col:1538
    SystemIommuStateInformation // SYSTEM_IOMMU_STATE_INFORMATION // since 22H1 = 234  //col:1539
    SystemHypervisorMinrootInformation // SYSTEM_HYPERVISOR_MINROOT_INFORMATION = 235  //col:1540
    SystemHypervisorBootPagesInformation // SYSTEM_HYPERVISOR_BOOT_PAGES_INFORMATION = 236  //col:1541
    SystemPointerAuthInformation // SYSTEM_POINTER_AUTH_INFORMATION = 237  //col:1542
    SystemSecureKernelDebuggerInformation = 238  //col:1543
    SystemOriginalImageFeatureInformation = 239  //col:1544
    MaxSystemInfoClass = 240  //col:1545
)


const(
    EventTraceKernelVersionInformation // EVENT_TRACE_VERSION_INFORMATION = 1  //col:2011
    EventTraceGroupMaskInformation // EVENT_TRACE_GROUPMASK_INFORMATION = 2  //col:2012
    EventTracePerformanceInformation // EVENT_TRACE_PERFORMANCE_INFORMATION = 3  //col:2013
    EventTraceTimeProfileInformation // EVENT_TRACE_TIME_PROFILE_INFORMATION = 4  //col:2014
    EventTraceSessionSecurityInformation // EVENT_TRACE_SESSION_SECURITY_INFORMATION = 5  //col:2015
    EventTraceSpinlockInformation // EVENT_TRACE_SPINLOCK_INFORMATION = 6  //col:2016
    EventTraceStackTracingInformation // EVENT_TRACE_SYSTEM_EVENT_INFORMATION = 7  //col:2017
    EventTraceExecutiveResourceInformation // EVENT_TRACE_EXECUTIVE_RESOURCE_INFORMATION = 8  //col:2018
    EventTraceHeapTracingInformation // EVENT_TRACE_HEAP_TRACING_INFORMATION = 9  //col:2019
    EventTraceHeapSummaryTracingInformation // EVENT_TRACE_HEAP_TRACING_INFORMATION = 10  //col:2020
    EventTracePoolTagFilterInformation // EVENT_TRACE_TAG_FILTER_INFORMATION = 11  //col:2021
    EventTracePebsTracingInformation // EVENT_TRACE_SYSTEM_EVENT_INFORMATION = 12  //col:2022
    EventTraceProfileConfigInformation // EVENT_TRACE_PROFILE_COUNTER_INFORMATION = 13  //col:2023
    EventTraceProfileSourceListInformation // EVENT_TRACE_PROFILE_LIST_INFORMATION = 14  //col:2024
    EventTraceProfileEventListInformation // EVENT_TRACE_SYSTEM_EVENT_INFORMATION = 15  //col:2025
    EventTraceProfileCounterListInformation // EVENT_TRACE_PROFILE_COUNTER_INFORMATION = 16  //col:2026
    EventTraceStackCachingInformation // EVENT_TRACE_STACK_CACHING_INFORMATION = 17  //col:2027
    EventTraceObjectTypeFilterInformation // EVENT_TRACE_TAG_FILTER_INFORMATION = 18  //col:2028
    EventTraceSoftRestartInformation // EVENT_TRACE_SOFT_RESTART_INFORMATION = 19  //col:2029
    EventTraceLastBranchConfigurationInformation // REDSTONE3 = 20  //col:2030
    EventTraceLastBranchEventListInformation = 21  //col:2031
    EventTraceProfileSourceAddInformation // EVENT_TRACE_PROFILE_ADD_INFORMATION // REDSTONE4 = 22  //col:2032
    EventTraceProfileSourceRemoveInformation // EVENT_TRACE_PROFILE_REMOVE_INFORMATION = 23  //col:2033
    EventTraceProcessorTraceConfigurationInformation = 24  //col:2034
    EventTraceProcessorTraceEventListInformation = 25  //col:2035
    EventTraceCoverageSamplerInformation // EVENT_TRACE_COVERAGE_SAMPLER_INFORMATION = 26  //col:2036
    EventTraceUnifiedStackCachingInformation // sicne 21H1 = 27  //col:2037
    MaxEventTraceInfoClass = 28  //col:2038
)


const(
    SystemCrashDumpDisable = 1  //col:2337
    SystemCrashDumpReconfigure = 2  //col:2338
    SystemCrashDumpInitializationComplete = 3  //col:2339
)


const(
    WdActionSetTimeoutValue = 1  //col:2634
    WdActionQueryTimeoutValue = 2  //col:2635
    WdActionResetTimer = 3  //col:2636
    WdActionStopTimer = 4  //col:2637
    WdActionStartTimer = 5  //col:2638
    WdActionSetTriggerAction = 6  //col:2639
    WdActionQueryTriggerAction = 7  //col:2640
    WdActionQueryState = 8  //col:2641
)


const(
    WdInfoTimeoutValue  =  0  //col:2655
    WdInfoResetTimer  =  1  //col:2656
    WdInfoStopTimer  =  2  //col:2657
    WdInfoStartTimer  =  3  //col:2658
    WdInfoTriggerAction  =  4  //col:2659
    WdInfoState  =  5  //col:2660
    WdInfoTriggerReset  =  6  //col:2661
    WdInfoNop  =  7  //col:2662
    WdInfoGeneratedLastReset  =  8  //col:2663
    WdInfoInvalid  =  9  //col:2664
)


const(
    SystemFirmwareTableEnumerate = 1  //col:2678
    SystemFirmwareTableGet = 2  //col:2679
    SystemFirmwareTableMax = 3  //col:2680
)


const(
    MemoryCaptureAccessedBits = 1  //col:2726
    MemoryCaptureAndResetAccessedBits = 2  //col:2727
    MemoryEmptyWorkingSets = 3  //col:2728
    MemoryFlushModifiedList = 4  //col:2729
    MemoryPurgeStandbyList = 5  //col:2730
    MemoryPurgeLowPriorityStandbyList = 6  //col:2731
    MemoryCommandMax = 7  //col:2732
)


const(
    CoverageAllModules  =  0  //col:2832
    CoverageSearchByHash  =  1  //col:2833
    CoverageSearchByName  =  2  //col:2834
)


const(
    SystemVaTypeAll = 1  //col:2978
    SystemVaTypeNonPagedPool = 2  //col:2979
    SystemVaTypePagedPool = 3  //col:2980
    SystemVaTypeSystemCache = 4  //col:2981
    SystemVaTypeSystemPtes = 5  //col:2982
    SystemVaTypeSessionSpace = 6  //col:2983
    SystemVaTypeMax = 7  //col:2984
)


const(
    StorePageRequest  =  1  //col:2999
    StoreStatsRequest  =  2 // q: SM_STATS_REQUEST // SmProcessStatsRequest  //col:3000
    StoreCreateRequest  =  3 // s: SM_CREATE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3001
    StoreDeleteRequest  =  4 // s: SM_DELETE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3002
    StoreListRequest  =  5 // q: SM_STORE_LIST_REQUEST / SM_STORE_LIST_REQUEST_EX // SmProcessListRequest  //col:3003
    Available1  =  6  //col:3004
    StoreEmptyRequest  =  7  //col:3005
    CacheListRequest  =  8 // q: SMC_CACHE_LIST_REQUEST // SmcProcessListRequest  //col:3006
    CacheCreateRequest  =  9 // s: SMC_CACHE_CREATE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3007
    CacheDeleteRequest  =  10 // s: SMC_CACHE_DELETE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3008
    CacheStoreCreateRequest  =  11 // s: SMC_STORE_CREATE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3009
    CacheStoreDeleteRequest  =  12 // s: SMC_STORE_DELETE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3010
    CacheStatsRequest  =  13 // q: SMC_CACHE_STATS_REQUEST // SmcProcessStatsRequest  //col:3011
    Available2  =  14  //col:3012
    RegistrationRequest  =  15 // q: SM_REGISTRATION_REQUEST (requires SeProfileSingleProcessPrivilege) // SmProcessRegistrationRequest  //col:3013
    GlobalCacheStatsRequest  =  16  //col:3014
    StoreResizeRequest  =  17 // s: SM_STORE_RESIZE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3015
    CacheStoreResizeRequest  =  18 // s: SMC_STORE_RESIZE_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3016
    SmConfigRequest  =  19 // s: SM_CONFIG_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3017
    StoreHighMemoryPriorityRequest  =  20 // s: SM_STORE_HIGH_MEM_PRIORITY_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3018
    SystemStoreTrimRequest  =  21 // s: SM_SYSTEM_STORE_TRIM_REQUEST (requires SeProfileSingleProcessPrivilege)  //col:3019
    MemCompressionInfoRequest  =  22  // q: SM_MEM_COMPRESSION_INFO_REQUEST // SmProcessCompressionInfoRequest  //col:3020
    ProcessStoreInfoRequest  =  23 // SmProcessProcessStoreInfoRequest  //col:3021
    StoreInformationMax = 24  //col:3022
)


const(
    StStatsLevelBasic  =  0  //col:3041
    StStatsLevelIoStats  =  1  //col:3042
    StStatsLevelRegionSpace  =  2 // requires SeProfileSingleProcessPrivilege  //col:3043
    StStatsLevelSpaceBitmap  =  3 // requires SeProfileSingleProcessPrivilege  //col:3044
    StStatsLevelMax  =  4  //col:3045
)


const(
    StoreTypeInMemory = 0  //col:3146
    StoreTypeFile = 1  //col:3147
    StoreTypeMax = 2  //col:3148
)


const(
    SmStoreManagerTypePhysical = 0  //col:3300
    SmStoreManagerTypeVirtual = 1  //col:3301
    SmStoreManagerTypeMax = 2  //col:3302
)


const(
    SmConfigDirtyPageCompression  =  0  //col:3395
    SmConfigAsyncInswap  =  1  //col:3396
    SmConfigPrefetchSeekThreshold  =  2  //col:3397
    SmConfigTypeMax  =  3  //col:3398
)


const(
    TpmBootEntropyStructureUninitialized = 1  //col:3507
    TpmBootEntropyDisabledByPolicy = 2  //col:3508
    TpmBootEntropyNoTpmFound = 3  //col:3509
    TpmBootEntropyTpmError = 4  //col:3510
    TpmBootEntropySuccess = 5  //col:3511
)


const(
    SystemPixelFormatUnknown = 1  //col:3602
    SystemPixelFormatR8G8B8 = 2  //col:3603
    SystemPixelFormatR8G8B8X8 = 3  //col:3604
    SystemPixelFormatB8G8R8 = 4  //col:3605
    SystemPixelFormatB8G8R8X8 = 5  //col:3606
)


const(
    SystemProcessClassificationNormal = 1  //col:3952
    SystemProcessClassificationSystem = 2  //col:3953
    SystemProcessClassificationSecureSystem = 3  //col:3954
    SystemProcessClassificationMemCompression = 4  //col:3955
    SystemProcessClassificationRegistry // REDSTONE4 = 5  //col:3956
    SystemProcessClassificationMaximum = 6  //col:3957
)


const(
    SystemActivityModerationStateSystemManaged = 1  //col:4258
    SystemActivityModerationStateUserManagedAllowThrottling = 2  //col:4259
    SystemActivityModerationStateUserManagedDisableThrottling = 3  //col:4260
    MaxSystemActivityModerationState = 4  //col:4261
)


const(
    SystemActivityModerationAppTypeClassic = 1  //col:4273
    SystemActivityModerationAppTypePackaged = 2  //col:4274
    MaxSystemActivityModerationAppType = 3  //col:4275
)


const(
    IommuStateBlock = 1  //col:4594
    IommuStateUnblock = 2  //col:4595
)


const(
    SysDbgQueryModuleInformation = 1  //col:4694
    SysDbgQueryTraceInformation = 2  //col:4695
    SysDbgSetTracepoint = 3  //col:4696
    SysDbgSetSpecialCall // PVOID = 4  //col:4697
    SysDbgClearSpecialCalls // void = 5  //col:4698
    SysDbgQuerySpecialCalls = 6  //col:4699
    SysDbgBreakPoint = 7  //col:4700
    SysDbgQueryVersion // DBGKD_GET_VERSION64 = 8  //col:4701
    SysDbgReadVirtual // SYSDBG_VIRTUAL = 9  //col:4702
    SysDbgWriteVirtual // SYSDBG_VIRTUAL = 10  //col:4703
    SysDbgReadPhysical // SYSDBG_PHYSICAL // 10 = 11  //col:4704
    SysDbgWritePhysical // SYSDBG_PHYSICAL = 12  //col:4705
    SysDbgReadControlSpace // SYSDBG_CONTROL_SPACE = 13  //col:4706
    SysDbgWriteControlSpace // SYSDBG_CONTROL_SPACE = 14  //col:4707
    SysDbgReadIoSpace // SYSDBG_IO_SPACE = 15  //col:4708
    SysDbgWriteIoSpace // SYSDBG_IO_SPACE = 16  //col:4709
    SysDbgReadMsr // SYSDBG_MSR = 17  //col:4710
    SysDbgWriteMsr // SYSDBG_MSR = 18  //col:4711
    SysDbgReadBusData // SYSDBG_BUS_DATA = 19  //col:4712
    SysDbgWriteBusData // SYSDBG_BUS_DATA = 20  //col:4713
    SysDbgCheckLowMemory // 20 = 21  //col:4714
    SysDbgEnableKernelDebugger = 22  //col:4715
    SysDbgDisableKernelDebugger = 23  //col:4716
    SysDbgGetAutoKdEnable = 24  //col:4717
    SysDbgSetAutoKdEnable = 25  //col:4718
    SysDbgGetPrintBufferSize = 26  //col:4719
    SysDbgSetPrintBufferSize = 27  //col:4720
    SysDbgGetKdUmExceptionEnable = 28  //col:4721
    SysDbgSetKdUmExceptionEnable = 29  //col:4722
    SysDbgGetTriageDump // SYSDBG_TRIAGE_DUMP = 30  //col:4723
    SysDbgGetKdBlockEnable // 30 = 31  //col:4724
    SysDbgSetKdBlockEnable = 32  //col:4725
    SysDbgRegisterForUmBreakInfo = 33  //col:4726
    SysDbgGetUmBreakPid = 34  //col:4727
    SysDbgClearUmBreakPid = 35  //col:4728
    SysDbgGetUmAttachPid = 36  //col:4729
    SysDbgClearUmAttachPid = 37  //col:4730
    SysDbgGetLiveKernelDump // SYSDBG_LIVEDUMP_CONTROL = 38  //col:4731
    SysDbgKdPullRemoteFile // SYSDBG_KD_PULL_REMOTE_FILE = 39  //col:4732
    SysDbgMaxInfoClass = 40  //col:4733
)


const(
    OptionAbortRetryIgnore = 1  //col:4887
    OptionOk = 2  //col:4888
    OptionOkCancel = 3  //col:4889
    OptionRetryCancel = 4  //col:4890
    OptionYesNo = 5  //col:4891
    OptionYesNoCancel = 6  //col:4892
    OptionShutdownSystem = 7  //col:4893
    OptionOkNoWait = 8  //col:4894
    OptionCancelTryContinue = 9  //col:4895
)


const(
    ResponseReturnToCaller = 1  //col:4900
    ResponseNotHandled = 2  //col:4901
    ResponseAbort = 3  //col:4902
    ResponseCancel = 4  //col:4903
    ResponseIgnore = 5  //col:4904
    ResponseNo = 6  //col:4905
    ResponseOk = 7  //col:4906
    ResponseRetry = 8  //col:4907
    ResponseYes = 9  //col:4908
    ResponseTryAgain = 10  //col:4909
    ResponseContinue = 11  //col:4910
)


const(
    StandardDesign = 1  //col:4931
    NEC98x86 = 2  //col:4932
    EndAlternatives = 3  //col:4933
)


const(
    AtomBasicInformation = 1  //col:5406
    AtomTableInformation = 2  //col:5407
)


const(
    ShutdownNoReboot = 1  //col:5539
    ShutdownReboot = 2  //col:5540
    ShutdownPowerOff = 3  //col:5541
    ShutdownRebootForRecovery // since WIN11 = 4  //col:5542
)



type (
Ntexapi interface{
#if ()(ok bool)//col:102
#if ()(ok bool)//col:276
#if ()(ok bool)//col:310
NtCreateEvent()(ok bool)//col:465
NtCreateMutant()(ok bool)//col:534
NtCreateSemaphore()(ok bool)//col:599
typedef VOID ()(ok bool)//col:617
#if ()(ok bool)//col:631
NtCreateTimer()(ok bool)//col:722
#if ()(ok bool)//col:891
#if ()(ok bool)//col:1095
#if ()(ok bool)//col:1198
NtWaitForWorkViaWorkerFactory()(ok bool)//col:1546
    ()(ok bool)//col:2189
typedef NTSTATUS ()(ok bool)//col:2651
#if ()(ok bool)//col:2681
#if ()(ok bool)//col:2707
    VOID ()(ok bool)//col:3711
#if !defined()(ok bool)//col:4225
#if ()(ok bool)//col:4734
NtSystemDebugControl()(ok bool)//col:4896
NtRaiseHardError()(ok bool)//col:4934
C_ASSERT()(ok bool)//col:5170
FORCEINLINE ULONG NtGetTickCount()(ok bool)//col:5198
FORCEINLINE ULONGLONG NtGetTickCount64()(ok bool)//col:5206
FORCEINLINE ULONG NtGetTickCount()(ok bool)//col:5212
NtQueryDefaultLocale()(ok bool)//col:5408
NtQueryInformationAtom()(ok bool)//col:5543
}






)

func NewNtexapi() { return & ntexapi{} }

func (n *ntexapi)#if ()(ok bool){//col:102




































































return true
}







func (n *ntexapi)#if ()(ok bool){//col:276
















































































































return true
}







func (n *ntexapi)#if ()(ok bool){//col:310

























return true
}







func (n *ntexapi)NtCreateEvent()(ok bool){//col:465
























































































































return true
}







func (n *ntexapi)NtCreateMutant()(ok bool){//col:534












































return true
}







func (n *ntexapi)NtCreateSemaphore()(ok bool){//col:599













































return true
}







func (n *ntexapi)typedef VOID ()(ok bool){//col:617









return true
}







func (n *ntexapi)#if ()(ok bool){//col:631












return true
}







func (n *ntexapi)NtCreateTimer()(ok bool){//col:722












































































return true
}







func (n *ntexapi)#if ()(ok bool){//col:891










































































































































return true
}







func (n *ntexapi)#if ()(ok bool){//col:1095



















































































































return true
}







func (n *ntexapi)#if ()(ok bool){//col:1198































































return true
}







func (n *ntexapi)NtWaitForWorkViaWorkerFactory()(ok bool){//col:1546









































































































return true
}







func (n *ntexapi)    ()(ok bool){//col:2189
















































































































return true
}







func (n *ntexapi)typedef NTSTATUS ()(ok bool){//col:2651






return true
}







func (n *ntexapi)#if ()(ok bool){//col:2681







return true
}







func (n *ntexapi)#if ()(ok bool){//col:2707











return true
}







func (n *ntexapi)    VOID ()(ok bool){//col:3711




return true
}







func (n *ntexapi)#if !defined()(ok bool){//col:4225











return true
}







func (n *ntexapi)#if ()(ok bool){//col:4734






















































return true
}







func (n *ntexapi)NtSystemDebugControl()(ok bool){//col:4896




















return true
}







func (n *ntexapi)NtRaiseHardError()(ok bool){//col:4934














return true
}







func (n *ntexapi)C_ASSERT()(ok bool){//col:5170






































return true
}







func (n *ntexapi)FORCEINLINE ULONG NtGetTickCount()(ok bool){//col:5198


















return true
}







func (n *ntexapi)FORCEINLINE ULONGLONG NtGetTickCount64()(ok bool){//col:5206



return true
}







func (n *ntexapi)FORCEINLINE ULONG NtGetTickCount()(ok bool){//col:5212




return true
}







func (n *ntexapi)NtQueryDefaultLocale()(ok bool){//col:5408
























































































































































return true
}







func (n *ntexapi)NtQueryInformationAtom()(ok bool){//col:5543































































return true
}









