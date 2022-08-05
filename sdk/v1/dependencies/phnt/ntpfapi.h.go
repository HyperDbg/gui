package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpfapi.h.back

const (
	PfKernelInitPhase               = 0   //col:3
	PfBootDriverInitPhase           = 90  //col:4
	PfSystemDriverInitPhase         = 120 //col:5
	PfSessionManagerInitPhase       = 150 //col:6
	PfSMRegistryInitPhase           = 180 //col:7
	PfVideoInitPhase                = 210 //col:8
	PfPostVideoInitPhase            = 240 //col:9
	PfBootAcceptedRegistryInitPhase = 270 //col:10
	PfUserShellReadyPhase           = 300 //col:11
	PfMaxBootPhaseId                = 900 //col:12
)

const (
	PfSvNotSpecified    = 1 //col:16
	PfSvEnabled         = 2 //col:17
	PfSvDisabled        = 3 //col:18
	PfSvMaxEnableStatus = 4 //col:19
)

const (
	PrefetcherRetrieveTrace            = 1 //col:23
	PrefetcherSystemParameters         = 2 //col:24
	PrefetcherBootPhase                = 3 //col:25
	PrefetcherSpare1                   = 4 //col:26
	PrefetcherBootControl              = 5 //col:27
	PrefetcherScenarioPolicyControl    = 6 //col:28
	PrefetcherSpare2                   = 7 //col:29
	PrefetcherAppLaunchScenarioControl = 8 //col:30
	PrefetcherInformationMax           = 9 //col:31
)

const (
	PfsPrivateSourceKernel  = 1 //col:35
	PfsPrivateSourceSession = 2 //col:36
	PfsPrivateSourceProcess = 3 //col:37
	PfsPrivateSourceMax     = 4 //col:38
)

const (
	PfScenarioTypeNone      = 1 //col:42
	PfScenarioTypeStandby   = 2 //col:43
	PfScenarioTypeHibernate = 3 //col:44
	PfScenarioTypeFUS       = 4 //col:45
	PfScenarioTypeMax       = 5 //col:46
)

const (
	SuperfetchRetrieveTrace            = 1  //col:50
	SuperfetchSystemParameters         = 2  //col:51
	SuperfetchLogEvent                 = 3  //col:52
	SuperfetchGenerateTrace            = 4  //col:53
	SuperfetchPrefetch                 = 5  //col:54
	SuperfetchPfnQuery                 = 6  //col:55
	SuperfetchPfnSetPriority           = 7  //col:56
	SuperfetchPrivSourceQuery          = 8  //col:57
	SuperfetchSequenceNumberQuery      = 9  //col:58
	SuperfetchScenarioPhase            = 10 //col:59
	SuperfetchWorkerPriority           = 11 //col:60
	SuperfetchScenarioQuery            = 12 //col:61
	SuperfetchScenarioPrefetch         = 13 //col:62
	SuperfetchRobustnessControl        = 14 //col:63
	SuperfetchTimeControl              = 15 //col:64
	SuperfetchMemoryListQuery          = 16 //col:65
	SuperfetchMemoryRangesQuery        = 17 //col:66
	SuperfetchTracingControl           = 18 //col:67
	SuperfetchTrimWhileAgingControl    = 19 //col:68
	SuperfetchRepurposedByPrefetch     = 20 //col:69
	SuperfetchChannelPowerRequest      = 21 //col:70
	SuperfetchMovePages                = 22 //col:71
	SuperfetchVirtualQuery             = 23 //col:72
	SuperfetchCombineStatsQuery        = 24 //col:73
	SuperfetchSetMinWsAgeRate          = 25 //col:74
	SuperfetchDeprioritizeOldPagesInWs = 26 //col:75
	SuperfetchFileExtentsQuery         = 27 //col:76
	SuperfetchGpuUtilizationQuery      = 28 //col:77
	SuperfetchInformationMax           = 29 //col:78
)

type _PF_TRACE_LIMITS struct {
	MaxNumPages    uint32   //col:8
	MaxNumSections uint32   //col:9
	TimerPeriod    LONGLONG //col:10
}

type _PF_SYSTEM_PREFETCH_PARAMETERS struct {
	EnableStatus           [2]PF_ENABLE_STATUS //col:17
	TraceLimits            [2]PF_TRACE_LIMITS  //col:18
	MaxNumActiveTraces     uint32              //col:19
	MaxNumSavedTraces      uint32              //col:20
	RootDirPath            [32]WCHAR           //col:21
	HostingApplicationList [128]WCHAR          //col:22
}

type _PF_BOOT_CONTROL struct {
	Version                uint32 //col:22
	DisableBootPrefetching uint32 //col:23
}

type _PREFETCHER_INFORMATION struct {
	ULONG                        _In_    //col:30
	ULONG                        _In_    //col:31
	PREFETCHER_INFORMATION_CLASS _In_    //col:32
	PVOID                        _Inout_ //col:33
	ULONG                        _Inout_ //col:34
}

type _PF_SYSTEM_SUPERFETCH_PARAMETERS struct {
	EnabledComponents                uint32 //col:40
	BootID                           uint32 //col:41
	SavedSectInfoTracesMax           uint32 //col:42
	SavedPageAccessTracesMax         uint32 //col:43
	ScenarioPrefetchTimeoutStandby   uint32 //col:44
	ScenarioPrefetchTimeoutHibernate uint32 //col:45
	ScenarioPrefetchTimeoutHiberBoot uint32 //col:46
}

type _PF_PFN_PRIO_REQUEST struct {
	Version      uint32                         //col:48
	RequestFlags uint32                         //col:49
	PfnCount     ULONG_PTR                      //col:50
	MemInfo      SYSTEM_MEMORY_LIST_INFORMATION //col:51
	PageData     [256]MMPFN_IDENTITY            //col:52
}

type _PFS_PRIVATE_PAGE_SOURCE struct {
	Type      PFS_PRIVATE_PAGE_SOURCE_TYPE //col:56
	Union     union                        //col:57
	SessionId uint32                       //col:59
	ProcessId uint32                       //col:60
}

type _PF_PRIVSOURCE_INFO struct {
	DbInfo                PFS_PRIVATE_PAGE_SOURCE //col:72
	EProcess              uintptr                 //col:73
	WsPrivatePages        int64                   //col:74
	TotalPrivatePages     int64                   //col:75
	SessionID             uint32                  //col:76
	ImageName             [16]int8                //col:77
	WsSwapPages           ULONG_PTR               //col:79
	SessionPagedPoolPages ULONG_PTR               //col:80
	StoreSizePages        ULONG_PTR               //col:81
}

type _PF_PRIVSOURCE_QUERY_REQUEST struct {
	Version   uint32                //col:87
	Flags     uint32                //col:88
	InfoCount uint32                //col:89
	InfoArray [1]PF_PRIVSOURCE_INFO //col:90
}

type _PF_SCENARIO_PHASE_INFO struct {
	Version        uint32                  //col:96
	ScenType       PF_PHASED_SCENARIO_TYPE //col:97
	PhaseId        uint32                  //col:98
	SequenceNumber uint32                  //col:99
	Flags          uint32                  //col:100
	FUSUserId      uint32                  //col:101
}

type _PF_MEMORY_LIST_NODE struct {
	Node                   ULONGLONG //col:106
	Spare                  ULONGLONG //col:107
	StandbyLowPageCount    ULONGLONG //col:108
	StandbyMediumPageCount ULONGLONG //col:109
	StandbyHighPageCount   ULONGLONG //col:110
	FreePageCount          ULONGLONG //col:111
	ModifiedPageCount      ULONGLONG //col:112
}

type _PF_MEMORY_LIST_INFO struct {
	Version   uint32                 //col:113
	Size      uint32                 //col:114
	NodeCount uint32                 //col:115
	Nodes     [1]PF_MEMORY_LIST_NODE //col:116
}

type _PF_PHYSICAL_MEMORY_RANGE struct {
	BasePfn   ULONG_PTR //col:118
	PageCount ULONG_PTR //col:119
}

type _PF_PHYSICAL_MEMORY_RANGE_INFO_V1 struct {
	Version    uint32                      //col:124
	RangeCount uint32                      //col:125
	Ranges     [1]PF_PHYSICAL_MEMORY_RANGE //col:126
}

type _PF_PHYSICAL_MEMORY_RANGE_INFO_V2 struct {
	Version    uint32                                  //col:131
	Flags      uint32                                  //col:132
	RangeCount uint32                                  //col:133
	Ranges     [ANYSIZE_ARRAY]PF_PHYSICAL_MEMORY_RANGE //col:134
}

type _PF_REPURPOSED_BY_PREFETCH_INFO struct {
	Version              uint32 //col:136
	RepurposedByPrefetch uint32 //col:137
}

type _SUPERFETCH_INFORMATION struct {
	ULONG                        _In_    //col:144
	ULONG                        _In_    //col:145
	SUPERFETCH_INFORMATION_CLASS _In_    //col:146
	PVOID                        _Inout_ //col:147
	ULONG                        _Inout_ //col:148
}

