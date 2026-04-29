//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

// Information taken from
// - http://www.ntinternals.net/: UserMode / NTDLL / System Information
// - http://www.geoffchappell.com/studies/windows/km/index.htm
// - MinGW headers

package ntdll

/*
func:
NTSTATUS NtQuerySystemInformation(
  _In_  SYSTEM_INFORMATION_CLASS SystemInformationClass,
  _Out_ PVOID                    SystemInformation,
  _In_  ULONG                    SystemInformationLength,
  _Out_ PULONG                   ReturnLength
);
*/

/*
func:
NTSTATUS NtSetSystemInformation(
  _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass,
  _In_ PVOID                    SystemInformation,
  _In_ ULONG                    SystemInformationLength
);
*/

/*
enum:
typedef enum _SYSTEM_INFORMATION_CLASS {
    SystemBasicInformation,
    SystemProcessorInformation,
    SystemPerformanceInformation,
    SystemTimeOfDayInformation,
    SystemPathInformation,
    SystemProcessInformation,
    SystemCallCountInformation,
    SystemDeviceInformation,
    SystemProcessorPerformanceInformation,
    SystemFlagsInformation,
    SystemCallTimeInformation,
    SystemModuleInformation,
    SystemLocksInformation,
    SystemStackTraceInformation,
    SystemPagedPoolInformation,
    SystemNonPagedPoolInformation,
    SystemHandleInformation,
    SystemObjectInformation,
    SystemPageFileInformation,
    SystemVdmInstemulInformation,
    SystemVdmBopInformation,
    SystemFileCacheInformation,
    SystemPoolTagInformation,
    SystemInterruptInformation,
    SystemDpcBehaviorInformation,
    SystemFullMemoryInformation,
    SystemLoadGdiDriverInformation,
    SystemUnloadGdiDriverInformation,
    SystemTimeAdjustmentInformation,
    SystemSummaryMemoryInformation,
    SystemNextEventIdInformation,
    SystemEventIdsInformation,
    SystemCrashDumpInformation,
    SystemExceptionInformation,
    SystemCrashDumpStateInformation,
    SystemKernelDebuggerInformation,
    SystemContextSwitchInformation,
    SystemRegistryQuotaInformation,
    SystemExtendServiceTableInformation,
    SystemPrioritySeperation,
    SystemPlugPlayBusInformation,
    SystemDockInformation,
    SystemPowerInformation,
    SystemProcessorSpeedInformation,
    SystemCurrentTimeZoneInformation,
    SystemLookasideInformation
} SYSTEM_INFORMATION_CLASS, *PSYSTEM_INFORMATION_CLASS;
*/

/*
type:
typedef struct _SYSTEM_MODULE {
  PVOID                Section;
  PVOID                MappedBase;
  PVOID                ImageBase;
  ULONG                ImageSize;
  ULONG                Flags;
  WORD                 LoadOrderIndex;
  WORD                 InitOrderIndex;
  WORD                 LoadCount;
  WORD                 NameOffset;
  BYTE                 Name[256];
} SYSTEM_MODULE, *PSYSTEM_MODULE;
*/

/*
type:
typedef struct _SYSTEM_BASIC_INFORMATION {
  BYTE Reserved1[4];
  ULONG MaximumIncrement;
  ULONG PhysicalPageSize;
  ULONG NumberOfPhysicalPages;
  ULONG LowestPhysicalPage;
  ULONG HighestPhysicalPage;
  ULONG AllocationGranularity;
  ULONG_PTR LowestUserAddress;
  ULONG_PTR HighestUserAddress;
  ULONG_PTR ActiveProcessors;
  CCHAR NumberOfProcessors;
} SYSTEM_BASIC_INFORMATION, *PSYSTEM_BASIC_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_PROCESSOR_INFORMATION {
  USHORT ProcessorArchitecture;
  USHORT ProcessorLevel;
  USHORT ProcessorRevision;
  USHORT Unknown;
  ULONG FeatureBits;
} SYSTEM_PROCESSOR_INFORMATION, *PSYSTEM_PROCESSOR_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION {
    LARGE_INTEGER IdleTime;
    LARGE_INTEGER KernelTime;
    LARGE_INTEGER UserTime;
    LARGE_INTEGER Reserved1[2];
    ULONG Reserved2;
} SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION, *PSYSTEM_PROCESSOR_PERFORMANCE_INFORMATION;
*/

/*
type:
typedef struct _IO_COUNTERS {
  ULONGLONG ReadOperationCount;
  ULONGLONG WriteOperationCount;
  ULONGLONG OtherOperationCount;
  ULONGLONG ReadTransferCount;
  ULONGLONG WriteTransferCount;
  ULONGLONG OtherTransferCount;
} IO_COUNTERS, *PIO_COUNTERS;
*/

/*
type:
typedef struct _VM_COUNTERS {
  SIZE_T PeakVirtualSize;
  SIZE_T VirtualSize;
  ULONG PageFaultCount;
  SIZE_T PeakWorkingSetSize;
  SIZE_T WorkingSetSize;
  SIZE_T QuotaPeakPagedPoolUsage;
  SIZE_T QuotaPagedPoolUsage;
  SIZE_T QuotaPeakNonPagedPoolUsage;
  SIZE_T QuotaNonPagedPoolUsage;
  SIZE_T PagefileUsage;
  SIZE_T PeakPagefileUsage;
} VM_COUNTERS, *PVM_COUNTERS;
*/

/*
type:
typedef struct _SYSTEM_PROCESS_INFORMATION {
  ULONG NextEntryOffset;
  ULONG NumberOfThreads;
  LARGE_INTEGER Reserved[3];
  LARGE_INTEGER CreateTime;
  LARGE_INTEGER UserTime;
  LARGE_INTEGER KernelTime;
  UNICODE_STRING ImageName;
  KPRIORITY BasePriority;
  HANDLE UniqueProcessId;
  HANDLE InheritedFromUniqueProcessId;
  ULONG HandleCount;
  ULONG SessionId;
  ULONG PageDirectoryBase;
  VM_COUNTERS VirtualMemoryCounters;
  SIZE_T PrivatePageCount;
  IO_COUNTERS IoCounters;
} SYSTEM_PROCESS_INFORMATION, *PSYSTEM_PROCESS_INFORMATION;
*/

type Kpriority uint64

/*
type:
typedef struct _CLIENT_ID {
  HANDLE UniqueProcess;
  HANDLE UniqueThread;
} CLIENT_ID, *PCLIENT_ID;
*/

/*
type:
typedef struct _SYSTEM_THREAD_INFORMATION {
    LARGE_INTEGER Reserved1[3];
    ULONG Reserved2;
    PVOID StartAddress;
    CLIENT_ID ClientId;
    KPRIORITY Priority;
    LONG BasePriority;
    ULONG Reserved3;
    ULONG ThreadState;
    ULONG WaitReason;
} SYSTEM_THREAD_INFORMATION, *PSYSTEM_THREAD_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_REGISTRY_QUOTA_INFORMATION {
    ULONG RegistryQuotaAllowed;
    ULONG RegistryQuotaUsed;
    PVOID Reserved1;
} SYSTEM_REGISTRY_QUOTA_INFORMATION, *PSYSTEM_REGISTRY_QUOTA_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_TIMEOFDAY_INFORMATION {
  LARGE_INTEGER BootTime;
  LARGE_INTEGER CurrentTime;
  LARGE_INTEGER TimeZoneBias;
  ULONG CurrentTimeZoneId;
  BYTE Reserved1[20];
} SYSTEM_TIMEOFDAY_INFORMATION,*PSYSTEM_TIMEOFDAY_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_PERFORMANCE_INFORMATION {
  LARGE_INTEGER IdleTime;
  LARGE_INTEGER ReadTransferCount;
  LARGE_INTEGER WriteTransferCount;
  LARGE_INTEGER OtherTransferCount;
  ULONG ReadOperationCount;
  ULONG WriteOperationCount;
  ULONG OtherOperationCount;
  ULONG AvailablePages;
  ULONG TotalCommittedPages;
  ULONG TotalCommitLimit;
  ULONG PeakCommitment;
  ULONG PageFaults;
  ULONG WriteCopyFaults;
  ULONG TransitionFaults;
  ULONG CacheTransitionFaults;
  ULONG DemandZeroFaults;
  ULONG PagesRead;
  ULONG PageReadIos;
  ULONG CacheReads;
  ULONG CacheIos;
  ULONG PagefilePagesWritten;
  ULONG PagefilePageWriteIos;
  ULONG MappedFilePagesWritten;
  ULONG MappedFilePageWriteIos;
  ULONG PagedPoolUsage;
  ULONG NonPagedPoolUsage;
  ULONG PagedPoolAllocs;
  ULONG PagedPoolFrees;
  ULONG NonPagedPoolAllocs;
  ULONG NonPagedPoolFrees;
  ULONG TotalFreeSystemPtes;
  ULONG SystemCodePage;
  ULONG TotalSystemDriverPages;
  ULONG TotalSystemCodePages;
  ULONG SmallNonPagedLookasideListAllocateHits;
  ULONG SmallPagedLookasideListAllocateHits;
  ULONG Reserved3;
  ULONG MmSystemCachePage;
  ULONG PagedPoolPage;
  ULONG SystemDriverPage;
  ULONG FastReadNoWait;
  ULONG FastReadWait;
  ULONG FastReadResourceMiss;
  ULONG FastReadNotPossible;
  ULONG FastMdlReadNoWait;
  ULONG FastMdlReadWait;
  ULONG FastMdlReadResourceMiss;
  ULONG FastMdlReadNotPossible;
  ULONG MapDataNoWait;
  ULONG MapDataWait;
  ULONG MapDataNoWaitMiss;
  ULONG MapDataWaitMiss;
  ULONG PinMappedDataCount;
  ULONG PinReadNoWait;
  ULONG PinReadWait;
  ULONG PinReadNoWaitMiss;
  ULONG PinReadWaitMiss;
  ULONG CopyReadNoWait;
  ULONG CopyReadWait;
  ULONG CopyReadNoWaitMiss;
  ULONG CopyReadWaitMiss;
  ULONG MdlReadNoWait;
  ULONG MdlReadWait;
  ULONG MdlReadNoWaitMiss;
  ULONG MdlReadWaitMiss;
  ULONG ReadAheadIos;
  ULONG LazyWriteIos;
  ULONG LazyWritePages;
  ULONG DataFlushes;
  ULONG DataPages;
  ULONG ContextSwitches;
  ULONG FirstLevelTbFills;
  ULONG SecondLevelTbFills;
  ULONG SystemCalls;
} SYSTEM_PERFORMANCE_INFORMATION, *PSYSTEM_PERFORMANCE_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_EXCEPTION_INFORMATION {
    BYTE Reserved1[16];
} SYSTEM_EXCEPTION_INFORMATION, *PSYSTEM_EXCEPTION_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_LOOKASIDE_INFORMATION {
    BYTE Reserved1[32];
} SYSTEM_LOOKASIDE_INFORMATION, *PSYSTEM_LOOKASIDE_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_INTERRUPT_INFORMATION {
    BYTE Reserved1[24];
} SYSTEM_INTERRUPT_INFORMATION, *PSYSTEM_INTERRUPT_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_POLICY_INFORMATION {
    PVOID Reserved1[2];
    ULONG Reserved2[3];
} SYSTEM_POLICY_INFORMATION, *PSYSTEM_POLICY_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_HANDLE_ENTRY {
  ULONG OwnerPid;
  BYTE ObjectType;
  BYTE HandleFlags;
  USHORT HandleValue;
  PVOID ObjectPointer;
  ULONG AccessMask;
} SYSTEM_HANDLE_ENTRY, *PSYSTEM_HANDLE_ENTRY;
*/

/*
type:
typedef struct _SYSTEM_HANDLE_INFORMATION {
  ULONG Count;
  SYSTEM_HANDLE_ENTRY Handles[1];
} SYSTEM_HANDLE_INFORMATION, *PSYSTEM_HANDLE_INFORMATION;
*/

func (mi *SystemHandleInformationT) GetHandles() []SystemHandleEntry {
	return mi.HandlesSlice(int(mi.Count))
}

/*
type:
typedef struct _SYSTEM_PAGEFILE_INFORMATION {
  ULONG NextEntryOffset;
  ULONG CurrentSize;
  ULONG TotalUsed;
  ULONG PeakUsed;
  UNICODE_STRING FileName;
} SYSTEM_PAGEFILE_INFORMATION, *PSYSTEM_PAGEFILE_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_MODULE_INFORMATION {
  ULONG                ModulesCount;
  SYSTEM_MODULE        Modules[ANYSIZE_ARRAY];
} SYSTEM_MODULE_INFORMATION, *PSYSTEM_MODULE_INFORMATION;
*/

func (mi *SystemModuleInformationT) GetModules() []SystemModule {
	return mi.ModulesSlice(int(mi.ModulesCount))
}
