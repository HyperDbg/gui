package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtQueryInformationProcess(
  HANDLE           ProcessHandle,
  PROCESS_INFORMATION_CLASS ProcessInformationClass,
  PVOID            ProcessInformation,
  ULONG            ProcessInformationLength,
  PULONG           ReturnLength
);
*/

/*
func:
NTSTATUS NtOpenProcess(
    _Out_ PHANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
);
*/

/*
func:
NTSTATUS NtOpenThread(
    _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_opt_ PCLIENT_ID ClientId
);
*/

/*
enum:
typedef enum _PROCESS_INFORMATION_CLASS {
  ProcessBasicInformation,
  ProcessQuotaLimits,
  ProcessIoCounters,
  ProcessVmCounters,
  ProcessTimes,
  ProcessBasePriority,
  ProcessRaisePriority,
  ProcessDebugPort,
  ProcessExceptionPort,
  ProcessAccessToken,
  ProcessLdtInformation,
  ProcessLdtSize,
  ProcessDefaultHardErrorMode,
  ProcessIoPortHandlers,
  ProcessPooledUsageAndLimits,
  ProcessWorkingSetWatch,
  ProcessUserModeIOPL,
  ProcessEnableAlignmentFaultFixup,
  ProcessPriorityClass,
  ProcessWx86Information,
  ProcessHandleCount,
  ProcessAffinityMask,
  ProcessPriorityBoost,
  ProcessDeviceMap,
  ProcessSessionInformation,
  ProcessForegroundInformation,
  ProcessWow64Information,
  ProcessImageFileName,
  ProcessLUIDDeviceMapsEnabled,
  ProcessBreakOnTermination,
  ProcessDebugObjectHandle,
  ProcessDebugFlags,
  ProcessHandleTracing,
  ProcessIoPriority,
  ProcessExecuteFlags,
  ProcessTlsInformation,
  ProcessCookie,
  ProcessImageInformation,
  ProcessCycleTime,
  ProcessPagePriority,
  ProcessInstrumentationCallback,
  ProcessThreadStackAllocation,
  ProcessWorkingSetWatchEx,
  ProcessImageFileNameWin32,
  ProcessImageFileMapping,
  ProcessAffinityUpdateMode,
  ProcessMemoryAllocationMode,
  ProcessGroupInformation,
  ProcessTokenVirtualizationEnabled,
  ProcessConsoleHostProcess,
  ProcessWindowInformation,
  ProcessHandleInformation,
  ProcessMitigationPolicy,
  ProcessDynamicFunctionTableInformation,
  ProcessHandleCheckingMode,
  ProcessKeepAliveCount,
  ProcessRevokeFileHandles,
  ProcessWorkingSetControl,
  ProcessHandleTable,
  ProcessCheckStackExtentsMode,
  ProcessCommandLineInformation,
  ProcessProtectionInformation,
  ProcessMemoryExhaustion,
  ProcessFaultInformation,
  ProcessTelemetryIdInformation,
  ProcessCommitReleaseInformation,
  ProcessDefaultCpuSetsInformation,
  ProcessAllowedCpuSetsInformation,
  ProcessSubsystemProcess,
  ProcessJobMemoryInformation,
  ProcessInPrivate,
  ProcessRaiseUMExceptionOnInvalidHandleClose,
  ProcessIumChallengeResponse,
  ProcessChildProcessInformation,
  ProcessHighGraphicsPriorityInformation,
  ProcessSubsystemInformation,
  ProcessEnergyValues,
  ProcessPowerThrottlingState,
  ProcessReserved3Information,
  ProcessWin32kSyscallFilterInformation,
  ProcessDisableSystemAllowedCpuSets,
  ProcessWakeInformation,
  ProcessEnergyTrackingState,
  ProcessManageWritesToExecutableMemory,
  ProcessCaptureTrustletLiveDump,
  ProcessTelemetryCoverage,
  ProcessEnclaveInformation,
  ProcessEnableReadWriteVmLogging,
  ProcessUptimeInformation,
  ProcessImageSection,
  ProcessDebugAuthInformation,
  ProcessSystemResourceManagement,
  ProcessSequenceNumber,
  ProcessLoaderDetour,
  ProcessSecurityDomainInformation,
  ProcessCombineSecurityDomainsInformation,
  ProcessEnableLogging,
  ProcessLeapSecondInformation,
  ProcessFiberShadowStackAllocation,
  ProcessFreeFiberShadowStackAllocation,
  ProcessAltSystemCallInformation,
  ProcessDynamicEHContinuationTargets,
  ProcessDynamicEnforcedCetCompatibleRanges,
  ProcessCreateStateChange,
  ProcessApplyStateChange,
  ProcessEnableOptionalXStateFeatures,
  ProcessAltPrefetchParam,
  ProcessAssignCpuPartitions,
  ProcessPriorityClassEx,
  ProcessMembershipInformation,
  ProcessEffectiveIoPriority,
  ProcessEffectivePagePriority,
} PROCESS_INFORMATION_CLASS;
*/

/*
type:
typedef struct _PEB {
  BYTE Reserved1[2];
  BYTE BeingDebugged;
  BYTE Reserved2[1];
  PVOID Reserved3[2];
  PPEB_LDR_DATA Ldr;
  PRTL_USER_PROCESS_PARAMETERS ProcessParameters;
  PVOID Reserved4[3];
  PVOID AtlThunkSListPtr;
  PVOID Reserved5;
  ULONG Reserved6;
  PVOID Reserved7;
  ULONG Reserved8;
  ULONG AtlThunkSListPtr32;
  PVOID Reserved9[45];
  BYTE Reserved10[96];
  PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
  BYTE Reserved11[128];
  PVOID Reserved12[1];
  ULONG SessionId;
} PEB,*PPEB;
*/

/*
type:
typedef struct _PEB_LDR_DATA {
  BYTE Reserved1[8];
  PVOID Reserved2[3];
  LIST_ENTRY InMemoryOrderModuleList;
} PEB_LDR_DATA,*PPEB_LDR_DATA;
*/

/*
type:
typedef struct _RTL_USER_PROCESS_PARAMETERS {
  BYTE Reserved1[16];
  PVOID Reserved2[10];
  UNICODE_STRING ImagePathName;
  UNICODE_STRING CommandLine;
} RTL_USER_PROCESS_PARAMETERS,*PRTL_USER_PROCESS_PARAMETERS;
*/

type PsPostProcessInitRoutine struct{} // FIXME

/*
type:
typedef struct _PROCESS_BASIC_INFORMATION
{
    NTSTATUS ExitStatus;
    PPEB PebBaseAddress;
    ULONG_PTR AffinityMask;
    KPRIORITY BasePriority;
    HANDLE UniqueProcessId;
    HANDLE InheritedFromUniqueProcessId;
} PROCESS_BASIC_INFORMATION, *PPROCESS_BASIC_INFORMATION;
*/

/*
func:
NTSTATUS NtQueryInformationThread(
  HANDLE          ThreadHandle,
  THREAD_INFORMATION_CLASS ThreadInformationClass,
  PVOID           ThreadInformation,
  ULONG           ThreadInformationLength,
  PULONG          ReturnLength
);
*/

/*
func:
NTSTATUS NtSetInformationThread(
  HANDLE          ThreadHandle,
  THREAD_INFORMATION_CLASS ThreadInformationClass,
  PVOID           ThreadInformation,
  ULONG           ThreadInformationLength
);
*/

/*
func:
NTSTATUS NtSetInformationProcess (
  HANDLE           ProcessHandle,
  PROCESS_INFORMATION_CLASS ProcessInformationClass,
  PVOID            ProcessInformation,
  ULONG            ProcessInformationLength
);
*/

/*
enum:
typedef enum _THREAD_INFORMATION_CLASS {
  ThreadBasicInformation,
  ThreadTimes,
  ThreadPriority,
  ThreadBasePriority,
  ThreadAffinityMask,
  ThreadImpersonationToken,
  ThreadDescriptorTableEntry,
  ThreadEnableAlignmentFaultFixup,
  ThreadEventPair_Reusable,
  ThreadQuerySetWin32StartAddress,
  ThreadZeroTlsCell,
  ThreadPerformanceCount,
  ThreadAmILastThread,
  ThreadIdealProcessor,
  ThreadPriorityBoost,
  ThreadSetTlsArrayAddress,
  ThreadIsIoPending,
  ThreadHideFromDebugger,
  ThreadBreakOnTermination,
  ThreadSwitchLegacyState,
  ThreadIsTerminated,
  ThreadLastSystemCall,
  ThreadIoPriority,
  ThreadCycleTime,
  ThreadPagePriority,
  ThreadActualBasePriority,
  ThreadTebInformation,
  ThreadCSwitchMon,
  ThreadCSwitchPmu,
  ThreadWow64Context,
  ThreadGroupInformation,
  ThreadUmsInformation,
  ThreadCounterProfiling,
  ThreadIdealProcessorEx,
  MaxThreadInfoClass
} THREAD_INFORMATION_CLASS;
*/

/*
func:
NTSTATUS
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
*/

/*
func:
NTSTATUS NtCreateProcessEx(
    _Out_ PHANDLE     ProcessHandle,
    _In_ ACCESS_MASK  DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE   ParentProcess,
    _In_ ULONG    Flags,
    _In_opt_ HANDLE SectionHandle,
    _In_opt_ HANDLE DebugPort,
    _In_opt_ HANDLE ExceptionPort,
    _In_ BOOLEAN  InJob
);
*/

const (
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_SET_SESSIONID             = 0x0004
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_ALL_ACCESS                = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xffff)
)

const (
	THREAD_TERMINATE                 = 0x0001
	THREAD_SUSPEND_RESUME            = 0x0002
	THREAD_GET_CONTEXT               = 0x0008
	THREAD_SET_CONTEXT               = 0x0010
	THREAD_SET_INFORMATION           = 0x0020
	THREAD_QUERY_INFORMATION         = 0x0040
	THREAD_SET_THREAD_TOKEN          = 0x0080
	THREAD_IMPERSONATE               = 0x0100
	THREAD_DIRECT_IMPERSONATION      = 0x0200
	THREAD_SET_LIMITED_INFORMATION   = 0x0400
	THREAD_QUERY_LIMITED_INFORMATION = 0x0800
)

const (
	NtCurrentProcess = Handle(^uintptr(0))
	NtCurrentThread  = Handle(^uintptr(0) - 1)
)
