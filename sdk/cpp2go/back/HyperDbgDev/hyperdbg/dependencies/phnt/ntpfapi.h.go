package phnt


const(
_NTPFAPI_H =  //col:13
PF_BOOT_CONTROL_VERSION = 1 //col:58
PREFETCHER_INFORMATION_VERSION = 23 // rev //col:79
PREFETCHER_INFORMATION_MAGIC = ('kuhC') // rev //col:80
PF_PFN_PRIO_REQUEST_VERSION = 1 //col:104
PF_PFN_PRIO_REQUEST_QUERY_MEMORY_LIST = 0x1 //col:105
PF_PFN_PRIO_REQUEST_VALID_FLAGS = 0x1 //col:106
PF_PRIVSOURCE_QUERY_REQUEST_VERSION = 8 //col:159
PF_SCENARIO_PHASE_INFO_VERSION = 4 //col:178
PF_MEMORY_LIST_INFO_VERSION = 1 //col:201
PF_PHYSICAL_MEMORY_RANGE_INFO_V1_VERSION = 1 //col:217
PF_PHYSICAL_MEMORY_RANGE_INFO_V2_VERSION = 2 //col:226
PF_REPURPOSED_BY_PREFETCH_INFO_VERSION = 1 //col:238
SUPERFETCH_INFORMATION_VERSION = 45 // rev //col:281
SUPERFETCH_INFORMATION_MAGIC = ('kuhC') // rev //col:282
)

const(
    PfKernelInitPhase  =  0  //col:21
    PfBootDriverInitPhase  =  90  //col:22
    PfSystemDriverInitPhase  =  120  //col:23
    PfSessionManagerInitPhase  =  150  //col:24
    PfSMRegistryInitPhase  =  180  //col:25
    PfVideoInitPhase  =  210  //col:26
    PfPostVideoInitPhase  =  240  //col:27
    PfBootAcceptedRegistryInitPhase  =  270  //col:28
    PfUserShellReadyPhase  =  300  //col:29
    PfMaxBootPhaseId  =  900  //col:30
)


const(
    PfSvNotSpecified = 1  //col:35
    PfSvEnabled = 2  //col:36
    PfSvDisabled = 3  //col:37
    PfSvMaxEnableStatus = 4  //col:38
)


const(
    PrefetcherRetrieveTrace  =  1 // q: CHAR[]  //col:68
    PrefetcherSystemParameters // q: PF_SYSTEM_PREFETCH_PARAMETERS = 2  //col:69
    PrefetcherBootPhase // s: PF_BOOT_PHASE_ID = 3  //col:70
    PrefetcherSpare1 // PrefetcherRetrieveBootLoaderTrace // q: CHAR[] = 4  //col:71
    PrefetcherBootControl // s: PF_BOOT_CONTROL = 5  //col:72
    PrefetcherScenarioPolicyControl = 6  //col:73
    PrefetcherSpare2 = 7  //col:74
    PrefetcherAppLaunchScenarioControl = 8  //col:75
    PrefetcherInformationMax = 9  //col:76
)


const(
    PfsPrivateSourceKernel = 1  //col:119
    PfsPrivateSourceSession = 2  //col:120
    PfsPrivateSourceProcess = 3  //col:121
    PfsPrivateSourceMax = 4  //col:122
)


const(
    PfScenarioTypeNone = 1  //col:171
    PfScenarioTypeStandby = 2  //col:172
    PfScenarioTypeHibernate = 3  //col:173
    PfScenarioTypeFUS = 4  //col:174
    PfScenarioTypeMax = 5  //col:175
)


const(
    SuperfetchRetrieveTrace  =  1 // q: CHAR[]  //col:250
    SuperfetchSystemParameters // q: PF_SYSTEM_SUPERFETCH_PARAMETERS = 2  //col:251
    SuperfetchLogEvent = 3  //col:252
    SuperfetchGenerateTrace = 4  //col:253
    SuperfetchPrefetch = 5  //col:254
    SuperfetchPfnQuery // q: PF_PFN_PRIO_REQUEST = 6  //col:255
    SuperfetchPfnSetPriority = 7  //col:256
    SuperfetchPrivSourceQuery // q: PF_PRIVSOURCE_QUERY_REQUEST = 8  //col:257
    SuperfetchSequenceNumberQuery // q: ULONG = 9  //col:258
    SuperfetchScenarioPhase // 10 = 10  //col:259
    SuperfetchWorkerPriority = 11  //col:260
    SuperfetchScenarioQuery // q: PF_SCENARIO_PHASE_INFO = 12  //col:261
    SuperfetchScenarioPrefetch = 13  //col:262
    SuperfetchRobustnessControl = 14  //col:263
    SuperfetchTimeControl = 15  //col:264
    SuperfetchMemoryListQuery // q: PF_MEMORY_LIST_INFO = 16  //col:265
    SuperfetchMemoryRangesQuery // q: PF_PHYSICAL_MEMORY_RANGE_INFO = 17  //col:266
    SuperfetchTracingControl = 18  //col:267
    SuperfetchTrimWhileAgingControl = 19  //col:268
    SuperfetchRepurposedByPrefetch // q: PF_REPURPOSED_BY_PREFETCH_INFO // rev = 20  //col:269
    SuperfetchChannelPowerRequest = 21  //col:270
    SuperfetchMovePages = 22  //col:271
    SuperfetchVirtualQuery = 23  //col:272
    SuperfetchCombineStatsQuery = 24  //col:273
    SuperfetchSetMinWsAgeRate = 25  //col:274
    SuperfetchDeprioritizeOldPagesInWs = 26  //col:275
    SuperfetchFileExtentsQuery = 27  //col:276
    SuperfetchGpuUtilizationQuery // PF_GPU_UTILIZATION_INFO = 28  //col:277
    SuperfetchInformationMax = 29  //col:278
)




