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
type     PfKernelInitPhase = 0 uint32
const(
    PfKernelInitPhase  PF_BOOT_PHASE_ID =  0  //col:21
    PfBootDriverInitPhase  PF_BOOT_PHASE_ID =  90  //col:22
    PfSystemDriverInitPhase  PF_BOOT_PHASE_ID =  120  //col:23
    PfSessionManagerInitPhase  PF_BOOT_PHASE_ID =  150  //col:24
    PfSMRegistryInitPhase  PF_BOOT_PHASE_ID =  180  //col:25
    PfVideoInitPhase  PF_BOOT_PHASE_ID =  210  //col:26
    PfPostVideoInitPhase  PF_BOOT_PHASE_ID =  240  //col:27
    PfBootAcceptedRegistryInitPhase  PF_BOOT_PHASE_ID =  270  //col:28
    PfUserShellReadyPhase  PF_BOOT_PHASE_ID =  300  //col:29
    PfMaxBootPhaseId  PF_BOOT_PHASE_ID =  900  //col:30
)
type     PfSvNotSpecified uint32
const(
    PfSvNotSpecified PF_ENABLE_STATUS = 1  //col:35
    PfSvEnabled PF_ENABLE_STATUS = 2  //col:36
    PfSvDisabled PF_ENABLE_STATUS = 3  //col:37
    PfSvMaxEnableStatus PF_ENABLE_STATUS = 4  //col:38
)
type     PrefetcherRetrieveTrace = 1 // q: CHAR[] uint32
const(
    PrefetcherRetrieveTrace  PREFETCHER_INFORMATION_CLASS =  1 // q: CHAR[]  //col:68
    PrefetcherSystemParameters // q: PF_SYSTEM_PREFETCH_PARAMETERS PREFETCHER_INFORMATION_CLASS = 2  //col:69
    PrefetcherBootPhase // s: PF_BOOT_PHASE_ID PREFETCHER_INFORMATION_CLASS = 3  //col:70
    PrefetcherSpare1 // PrefetcherRetrieveBootLoaderTrace // q: CHAR[] PREFETCHER_INFORMATION_CLASS = 4  //col:71
    PrefetcherBootControl // s: PF_BOOT_CONTROL PREFETCHER_INFORMATION_CLASS = 5  //col:72
    PrefetcherScenarioPolicyControl PREFETCHER_INFORMATION_CLASS = 6  //col:73
    PrefetcherSpare2 PREFETCHER_INFORMATION_CLASS = 7  //col:74
    PrefetcherAppLaunchScenarioControl PREFETCHER_INFORMATION_CLASS = 8  //col:75
    PrefetcherInformationMax PREFETCHER_INFORMATION_CLASS = 9  //col:76
)
type     PfsPrivateSourceKernel uint32
const(
    PfsPrivateSourceKernel PFS_PRIVATE_PAGE_SOURCE_TYPE = 1  //col:119
    PfsPrivateSourceSession PFS_PRIVATE_PAGE_SOURCE_TYPE = 2  //col:120
    PfsPrivateSourceProcess PFS_PRIVATE_PAGE_SOURCE_TYPE = 3  //col:121
    PfsPrivateSourceMax PFS_PRIVATE_PAGE_SOURCE_TYPE = 4  //col:122
)
type     PfScenarioTypeNone uint32
const(
    PfScenarioTypeNone PF_PHASED_SCENARIO_TYPE = 1  //col:171
    PfScenarioTypeStandby PF_PHASED_SCENARIO_TYPE = 2  //col:172
    PfScenarioTypeHibernate PF_PHASED_SCENARIO_TYPE = 3  //col:173
    PfScenarioTypeFUS PF_PHASED_SCENARIO_TYPE = 4  //col:174
    PfScenarioTypeMax PF_PHASED_SCENARIO_TYPE = 5  //col:175
)
type     SuperfetchRetrieveTrace = 1 // q: CHAR[] uint32
const(
    SuperfetchRetrieveTrace  SUPERFETCH_INFORMATION_CLASS =  1 // q: CHAR[]  //col:250
    SuperfetchSystemParameters // q: PF_SYSTEM_SUPERFETCH_PARAMETERS SUPERFETCH_INFORMATION_CLASS = 2  //col:251
    SuperfetchLogEvent SUPERFETCH_INFORMATION_CLASS = 3  //col:252
    SuperfetchGenerateTrace SUPERFETCH_INFORMATION_CLASS = 4  //col:253
    SuperfetchPrefetch SUPERFETCH_INFORMATION_CLASS = 5  //col:254
    SuperfetchPfnQuery // q: PF_PFN_PRIO_REQUEST SUPERFETCH_INFORMATION_CLASS = 6  //col:255
    SuperfetchPfnSetPriority SUPERFETCH_INFORMATION_CLASS = 7  //col:256
    SuperfetchPrivSourceQuery // q: PF_PRIVSOURCE_QUERY_REQUEST SUPERFETCH_INFORMATION_CLASS = 8  //col:257
    SuperfetchSequenceNumberQuery // q: ULONG SUPERFETCH_INFORMATION_CLASS = 9  //col:258
    SuperfetchScenarioPhase // 10 SUPERFETCH_INFORMATION_CLASS = 10  //col:259
    SuperfetchWorkerPriority SUPERFETCH_INFORMATION_CLASS = 11  //col:260
    SuperfetchScenarioQuery // q: PF_SCENARIO_PHASE_INFO SUPERFETCH_INFORMATION_CLASS = 12  //col:261
    SuperfetchScenarioPrefetch SUPERFETCH_INFORMATION_CLASS = 13  //col:262
    SuperfetchRobustnessControl SUPERFETCH_INFORMATION_CLASS = 14  //col:263
    SuperfetchTimeControl SUPERFETCH_INFORMATION_CLASS = 15  //col:264
    SuperfetchMemoryListQuery // q: PF_MEMORY_LIST_INFO SUPERFETCH_INFORMATION_CLASS = 16  //col:265
    SuperfetchMemoryRangesQuery // q: PF_PHYSICAL_MEMORY_RANGE_INFO SUPERFETCH_INFORMATION_CLASS = 17  //col:266
    SuperfetchTracingControl SUPERFETCH_INFORMATION_CLASS = 18  //col:267
    SuperfetchTrimWhileAgingControl SUPERFETCH_INFORMATION_CLASS = 19  //col:268
    SuperfetchRepurposedByPrefetch // q: PF_REPURPOSED_BY_PREFETCH_INFO // rev SUPERFETCH_INFORMATION_CLASS = 20  //col:269
    SuperfetchChannelPowerRequest SUPERFETCH_INFORMATION_CLASS = 21  //col:270
    SuperfetchMovePages SUPERFETCH_INFORMATION_CLASS = 22  //col:271
    SuperfetchVirtualQuery SUPERFETCH_INFORMATION_CLASS = 23  //col:272
    SuperfetchCombineStatsQuery SUPERFETCH_INFORMATION_CLASS = 24  //col:273
    SuperfetchSetMinWsAgeRate SUPERFETCH_INFORMATION_CLASS = 25  //col:274
    SuperfetchDeprioritizeOldPagesInWs SUPERFETCH_INFORMATION_CLASS = 26  //col:275
    SuperfetchFileExtentsQuery SUPERFETCH_INFORMATION_CLASS = 27  //col:276
    SuperfetchGpuUtilizationQuery // PF_GPU_UTILIZATION_INFO SUPERFETCH_INFORMATION_CLASS = 28  //col:277
    SuperfetchInformationMax SUPERFETCH_INFORMATION_CLASS = 29  //col:278
)
type (
Ntpfapi interface{
 * Attribution 4.0 International ()(ok bool)//col:31
}

)
func NewNtpfapi() { return & ntpfapi{} }
func (n *ntpfapi) * Attribution 4.0 International ()(ok bool){//col:31
return true
}

