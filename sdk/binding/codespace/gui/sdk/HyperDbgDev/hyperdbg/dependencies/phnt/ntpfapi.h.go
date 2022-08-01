package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpfapi.h.back

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
MaxNumPages uint32 //col:3
MaxNumSections uint32 //col:4
TimerPeriod LONGLONG //col:5
}


type PF_SYSTEM_PREFETCH_PARAMETERS struct{
EnableStatus[2] PF_ENABLE_STATUS //col:9
TraceLimits[2] PF_TRACE_LIMITS //col:10
MaxNumActiveTraces uint32 //col:11
MaxNumSavedTraces uint32 //col:12
RootDirPath[32] WCHAR //col:13
HostingApplicationList[128] WCHAR //col:14
}


type PF_BOOT_CONTROL struct{
Version uint32 //col:18
DisableBootPrefetching uint32 //col:19
}


type PREFETCHER_INFORMATION struct{
ULONG _In_ //col:23
ULONG _In_ //col:24
PREFETCHER_INFORMATION_CLASS _In_ //col:25
PVOID _Inout_ //col:26
ULONG _Inout_ //col:27
}


type PF_SYSTEM_SUPERFETCH_PARAMETERS struct{
EnabledComponents uint32 //col:31
BootID uint32 //col:32
SavedSectInfoTracesMax uint32 //col:33
SavedPageAccessTracesMax uint32 //col:34
ScenarioPrefetchTimeoutStandby uint32 //col:35
ScenarioPrefetchTimeoutHibernate uint32 //col:36
ScenarioPrefetchTimeoutHiberBoot uint32 //col:37
}


type PF_PFN_PRIO_REQUEST struct{
Version uint32 //col:41
RequestFlags uint32 //col:42
PfnCount ULONG_PTR //col:43
MemInfo SYSTEM_MEMORY_LIST_INFORMATION //col:44
PageData[256] MMPFN_IDENTITY //col:45
}


type PFS_PRIVATE_PAGE_SOURCE struct{
Type PFS_PRIVATE_PAGE_SOURCE_TYPE //col:49
Union union //col:50
SessionId uint32 //col:52
ProcessId uint32 //col:53
}


type PF_PRIVSOURCE_INFO struct{
DbInfo PFS_PRIVATE_PAGE_SOURCE //col:60
EProcess PVOID //col:61
WsPrivatePages SIZE_T //col:62
TotalPrivatePages SIZE_T //col:63
SessionID uint32 //col:64
ImageName[16] int8 //col:65
WsSwapPages ULONG_PTR //col:67
SessionPagedPoolPages ULONG_PTR //col:68
StoreSizePages ULONG_PTR //col:69
}


type PF_PRIVSOURCE_QUERY_REQUEST struct{
Version uint32 //col:81
Flags uint32 //col:82
InfoCount uint32 //col:83
InfoArray[1] PF_PRIVSOURCE_INFO //col:84
}


type PF_SCENARIO_PHASE_INFO struct{
Version uint32 //col:88
ScenType PF_PHASED_SCENARIO_TYPE //col:89
PhaseId uint32 //col:90
SequenceNumber uint32 //col:91
Flags uint32 //col:92
FUSUserId uint32 //col:93
}


type PF_MEMORY_LIST_NODE struct{
Node ULONGLONG //col:97
Spare ULONGLONG //col:98
StandbyLowPageCount ULONGLONG //col:99
StandbyMediumPageCount ULONGLONG //col:100
StandbyHighPageCount ULONGLONG //col:101
FreePageCount ULONGLONG //col:102
ModifiedPageCount ULONGLONG //col:103
}


type PF_MEMORY_LIST_INFO struct{
Version uint32 //col:107
Size uint32 //col:108
NodeCount uint32 //col:109
Nodes[1] PF_MEMORY_LIST_NODE //col:110
}


type PF_PHYSICAL_MEMORY_RANGE struct{
BasePfn ULONG_PTR //col:114
PageCount ULONG_PTR //col:115
}


type PF_PHYSICAL_MEMORY_RANGE_INFO_V1 struct{
Version uint32 //col:119
RangeCount uint32 //col:120
Ranges[1] PF_PHYSICAL_MEMORY_RANGE //col:121
}


type PF_PHYSICAL_MEMORY_RANGE_INFO_V2 struct{
Version uint32 //col:125
Flags uint32 //col:126
RangeCount uint32 //col:127
Ranges[ANYSIZE_ARRAY] PF_PHYSICAL_MEMORY_RANGE //col:128
}


type PF_REPURPOSED_BY_PREFETCH_INFO struct{
Version uint32 //col:132
RepurposedByPrefetch uint32 //col:133
}


type SUPERFETCH_INFORMATION struct{
ULONG _In_ //col:137
ULONG _In_ //col:138
SUPERFETCH_INFORMATION_CLASS _In_ //col:139
PVOID _Inout_ //col:140
ULONG _Inout_ //col:141
}




