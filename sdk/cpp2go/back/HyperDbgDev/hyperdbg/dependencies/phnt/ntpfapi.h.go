package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpfapi.h.back

const(
_NTPFAPI_H =  //col:1
PF_BOOT_CONTROL_VERSION = 1 //col:2
PREFETCHER_INFORMATION_VERSION = 23 //col:3
PREFETCHER_INFORMATION_MAGIC = ('kuhC') //col:4
PF_PFN_PRIO_REQUEST_VERSION = 1 //col:5
PF_PFN_PRIO_REQUEST_QUERY_MEMORY_LIST = 0x1 //col:6
PF_PFN_PRIO_REQUEST_VALID_FLAGS = 0x1 //col:7
PF_PRIVSOURCE_QUERY_REQUEST_VERSION = 8 //col:8
PF_SCENARIO_PHASE_INFO_VERSION = 4 //col:9
PF_MEMORY_LIST_INFO_VERSION = 1 //col:10
PF_PHYSICAL_MEMORY_RANGE_INFO_V1_VERSION = 1 //col:11
PF_PHYSICAL_MEMORY_RANGE_INFO_V2_VERSION = 2 //col:12
PF_REPURPOSED_BY_PREFETCH_INFO_VERSION = 1 //col:13
SUPERFETCH_INFORMATION_VERSION = 45 //col:14
SUPERFETCH_INFORMATION_MAGIC = ('kuhC') //col:15
)

const(
    PfKernelInitPhase  =  0  //col:3
    PfBootDriverInitPhase  =  90  //col:4
    PfSystemDriverInitPhase  =  120  //col:5
    PfSessionManagerInitPhase  =  150  //col:6
    PfSMRegistryInitPhase  =  180  //col:7
    PfVideoInitPhase  =  210  //col:8
    PfPostVideoInitPhase  =  240  //col:9
    PfBootAcceptedRegistryInitPhase  =  270  //col:10
    PfUserShellReadyPhase  =  300  //col:11
    PfMaxBootPhaseId  =  900  //col:12
)


const(
    PfSvNotSpecified = 1  //col:16
    PfSvEnabled = 2  //col:17
    PfSvDisabled = 3  //col:18
    PfSvMaxEnableStatus = 4  //col:19
)


const(
    PrefetcherRetrieveTrace  =  1   //col:23
    PrefetcherSystemParameters  = 2  //col:24
    PrefetcherBootPhase  = 3  //col:25
    PrefetcherSpare1  = 4  //col:26
    PrefetcherBootControl  = 5  //col:27
    PrefetcherScenarioPolicyControl = 6  //col:28
    PrefetcherSpare2 = 7  //col:29
    PrefetcherAppLaunchScenarioControl = 8  //col:30
    PrefetcherInformationMax = 9  //col:31
)


const(
    PfsPrivateSourceKernel = 1  //col:35
    PfsPrivateSourceSession = 2  //col:36
    PfsPrivateSourceProcess = 3  //col:37
    PfsPrivateSourceMax = 4  //col:38
)


const(
    PfScenarioTypeNone = 1  //col:42
    PfScenarioTypeStandby = 2  //col:43
    PfScenarioTypeHibernate = 3  //col:44
    PfScenarioTypeFUS = 4  //col:45
    PfScenarioTypeMax = 5  //col:46
)


const(
    SuperfetchRetrieveTrace  =  1   //col:50
    SuperfetchSystemParameters  = 2  //col:51
    SuperfetchLogEvent = 3  //col:52
    SuperfetchGenerateTrace = 4  //col:53
    SuperfetchPrefetch = 5  //col:54
    SuperfetchPfnQuery  = 6  //col:55
    SuperfetchPfnSetPriority = 7  //col:56
    SuperfetchPrivSourceQuery  = 8  //col:57
    SuperfetchSequenceNumberQuery  = 9  //col:58
    SuperfetchScenarioPhase  = 10  //col:59
    SuperfetchWorkerPriority = 11  //col:60
    SuperfetchScenarioQuery  = 12  //col:61
    SuperfetchScenarioPrefetch = 13  //col:62
    SuperfetchRobustnessControl = 14  //col:63
    SuperfetchTimeControl = 15  //col:64
    SuperfetchMemoryListQuery  = 16  //col:65
    SuperfetchMemoryRangesQuery  = 17  //col:66
    SuperfetchTracingControl = 18  //col:67
    SuperfetchTrimWhileAgingControl = 19  //col:68
    SuperfetchRepurposedByPrefetch  = 20  //col:69
    SuperfetchChannelPowerRequest = 21  //col:70
    SuperfetchMovePages = 22  //col:71
    SuperfetchVirtualQuery = 23  //col:72
    SuperfetchCombineStatsQuery = 24  //col:73
    SuperfetchSetMinWsAgeRate = 25  //col:74
    SuperfetchDeprioritizeOldPagesInWs = 26  //col:75
    SuperfetchFileExtentsQuery = 27  //col:76
    SuperfetchGpuUtilizationQuery  = 28  //col:77
    SuperfetchInformationMax = 29  //col:78
)



type PF_TRACE_LIMITS struct{
MaxNumPages ULONG
MaxNumSections ULONG
TimerPeriod LONGLONG
}


type PF_SYSTEM_PREFETCH_PARAMETERS struct{
EnableStatus[2] PF_ENABLE_STATUS
TraceLimits[2] PF_TRACE_LIMITS
MaxNumActiveTraces ULONG
MaxNumSavedTraces ULONG
RootDirPath[32] WCHAR
HostingApplicationList[128] WCHAR
}


type PF_BOOT_CONTROL struct{
Version ULONG
DisableBootPrefetching ULONG
}


type PREFETCHER_INFORMATION struct{
ULONG _In_
ULONG _In_
PREFETCHER_INFORMATION_CLASS _In_
PVOID _Inout_
ULONG _Inout_
}


type PF_SYSTEM_SUPERFETCH_PARAMETERS struct{
EnabledComponents ULONG
BootID ULONG
SavedSectInfoTracesMax ULONG
SavedPageAccessTracesMax ULONG
ScenarioPrefetchTimeoutStandby ULONG
ScenarioPrefetchTimeoutHibernate ULONG
ScenarioPrefetchTimeoutHiberBoot ULONG
}


type PF_PFN_PRIO_REQUEST struct{
Version ULONG
RequestFlags ULONG
PfnCount ULONG_PTR
MemInfo SYSTEM_MEMORY_LIST_INFORMATION
PageData[256] MMPFN_IDENTITY
}


type PFS_PRIVATE_PAGE_SOURCE struct{
Type PFS_PRIVATE_PAGE_SOURCE_TYPE
Union union
SessionId ULONG
ProcessId ULONG
}


type PF_PRIVSOURCE_INFO struct{
DbInfo PFS_PRIVATE_PAGE_SOURCE
EProcess PVOID
WsPrivatePages SIZE_T
TotalPrivatePages SIZE_T
SessionID ULONG
ImageName[16] CHAR
WsSwapPages ULONG_PTR
SessionPagedPoolPages ULONG_PTR
StoreSizePages ULONG_PTR
}


type PF_PRIVSOURCE_QUERY_REQUEST struct{
Version ULONG
Flags ULONG
InfoCount ULONG
InfoArray[1] PF_PRIVSOURCE_INFO
}


type PF_SCENARIO_PHASE_INFO struct{
Version ULONG
ScenType PF_PHASED_SCENARIO_TYPE
PhaseId ULONG
SequenceNumber ULONG
Flags ULONG
FUSUserId ULONG
}


type PF_MEMORY_LIST_NODE struct{
Node ULONGLONG
Spare ULONGLONG
StandbyLowPageCount ULONGLONG
StandbyMediumPageCount ULONGLONG
StandbyHighPageCount ULONGLONG
FreePageCount ULONGLONG
ModifiedPageCount ULONGLONG
}


type PF_MEMORY_LIST_INFO struct{
Version ULONG
Size ULONG
NodeCount ULONG
Nodes[1] PF_MEMORY_LIST_NODE
}


type PF_PHYSICAL_MEMORY_RANGE struct{
BasePfn ULONG_PTR
PageCount ULONG_PTR
}


type PF_PHYSICAL_MEMORY_RANGE_INFO_V1 struct{
Version ULONG
RangeCount ULONG
Ranges[1] PF_PHYSICAL_MEMORY_RANGE
}


type PF_PHYSICAL_MEMORY_RANGE_INFO_V2 struct{
Version ULONG
Flags ULONG
RangeCount ULONG
Ranges[ANYSIZE_ARRAY] PF_PHYSICAL_MEMORY_RANGE
}


type PF_REPURPOSED_BY_PREFETCH_INFO struct{
Version ULONG
RepurposedByPrefetch ULONG
}


type SUPERFETCH_INFORMATION struct{
ULONG _In_
ULONG _In_
SUPERFETCH_INFORMATION_CLASS _In_
PVOID _Inout_
ULONG _Inout_
}




