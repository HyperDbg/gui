/*
 * This file is part of the Process Hacker project -
 * https:processhacker.sourceforge.io/
 *
 * You can redistribute this file and/or modify it under the terms of the
 * Attribution 4.0 International (CC BY 4.0) license.
 *
 * You must give appropriate credit, provide a link to the license, and
 * indicate if changes were made. You may do so in any reasonable manner, but
 * not in any way that suggests the licensor endorses you or your use.
 */
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
typedef struct _PEB_LDR_DATA {
  ULONG Length;
  BOOLEAN Initialized;
  HANDLE SsHandle;
  LIST_ENTRY InLoadOrderModuleList;
  LIST_ENTRY InMemoryOrderModuleList;
  LIST_ENTRY InInitializationOrderModuleList;
  PVOID EntryInProgress;
  BOOLEAN ShutdownInProgress;
  HANDLE ShutdownThreadId;
} PEB_LDR_DATA, *PPEB_LDR_DATA;
typedef struct _INITIAL_TEB {
  struct {
    PVOID OldStackBase;
    PVOID OldStackLimit;
  } OldInitialTeb;
  PVOID StackBase;
  PVOID StackLimit;
  PVOID StackAllocationBase;
} INITIAL_TEB, *PINITIAL_TEB;
typedef struct _WOW64_PROCESS {
  PVOID Wow64;
} WOW64_PROCESS, *PWOW64_PROCESS;
#include <ntpebteb.h>
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _PROCESSINFOCLASS {
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
  MaxProcessInfoClass
} PROCESSINFOCLASS;
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _THREADINFOCLASS {
  ThreadBasicInformation, 
  ThreadTimes,            
  ThreadPriority,     
  ThreadBasePriority, 
  ThreadAffinityMask, 
  ThreadImpersonationToken,        
  ThreadDescriptorTableEntry,      
  ThreadEnableAlignmentFaultFixup, 
  ThreadEventPair,
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
  ThreadCpuAccountingInformation, 
  ThreadSuspendCount,             
  ThreadHeterogeneousCpuPolicy,   
  ThreadContainerId,              
  ThreadNameInformation,          
  ThreadSelectedCpuSets,
  ThreadSystemThreadInformation, 
  ThreadActualGroupAffinity,     
  ThreadDynamicCodePolicyInfo,   
  ThreadExplicitCaseSensitivity, 
  ThreadWorkOnBehalfTicket,      
  ThreadSubsystemInformation,    
  ThreadDbgkWerReportActive,     
  ThreadAttachContainer,         
  ThreadManageWritesToExecutableMemory, 
  ThreadPowerThrottlingState,           
  ThreadWorkloadClass,     
  ThreadCreateStateChange, 
  ThreadApplyStateChange,
  ThreadStrongerBadHandleChecks, 
  ThreadEffectiveIoPriority,
  ThreadEffectivePagePriority,
  MaxThreadInfoClass
} THREADINFOCLASS;
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PAGE_PRIORITY_INFORMATION {
  ULONG PagePriority;
} PAGE_PRIORITY_INFORMATION, *PPAGE_PRIORITY_INFORMATION;
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PROCESS_BASIC_INFORMATION {
  NTSTATUS ExitStatus;
  PPEB PebBaseAddress;
  KAFFINITY AffinityMask;
  KPRIORITY BasePriority;
  HANDLE UniqueProcessId;
  HANDLE InheritedFromUniqueProcessId;
} PROCESS_BASIC_INFORMATION, *PPROCESS_BASIC_INFORMATION;
typedef struct _PROCESS_EXTENDED_BASIC_INFORMATION {
  SIZE_T Size; 
  PROCESS_BASIC_INFORMATION BasicInfo;
  union {
    ULONG Flags;
    struct {
      ULONG IsProtectedProcess : 1;
      ULONG IsWow64Process : 1;
      ULONG IsProcessDeleting : 1;
      ULONG IsCrossSessionCreate : 1;
      ULONG IsFrozen : 1;
      ULONG IsBackground : 1;
      ULONG IsStronglyNamed : 1;
      ULONG IsSecureProcess : 1;
      ULONG IsSubsystemProcess : 1;
      ULONG SpareBits : 23;
    };
  };
} PROCESS_EXTENDED_BASIC_INFORMATION, *PPROCESS_EXTENDED_BASIC_INFORMATION;
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
typedef struct _VM_COUNTERS_EX {
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
  SIZE_T PrivateUsage;
} VM_COUNTERS_EX, *PVM_COUNTERS_EX;
typedef struct _VM_COUNTERS_EX2 {
  VM_COUNTERS_EX CountersEx;
  SIZE_T PrivateWorkingSetSize;
  SIZE_T SharedCommitUsage;
} VM_COUNTERS_EX2, *PVM_COUNTERS_EX2;
typedef struct _KERNEL_USER_TIMES {
  LARGE_INTEGER CreateTime;
  LARGE_INTEGER ExitTime;
  LARGE_INTEGER KernelTime;
  LARGE_INTEGER UserTime;
} KERNEL_USER_TIMES, *PKERNEL_USER_TIMES;
typedef struct _POOLED_USAGE_AND_LIMITS {
  SIZE_T PeakPagedPoolUsage;
  SIZE_T PagedPoolUsage;
  SIZE_T PagedPoolLimit;
  SIZE_T PeakNonPagedPoolUsage;
  SIZE_T NonPagedPoolUsage;
  SIZE_T NonPagedPoolLimit;
  SIZE_T PeakPagefileUsage;
  SIZE_T PagefileUsage;
  SIZE_T PagefileLimit;
} POOLED_USAGE_AND_LIMITS, *PPOOLED_USAGE_AND_LIMITS;
#define PROCESS_EXCEPTION_PORT_ALL_STATE_BITS 0x00000003
#define PROCESS_EXCEPTION_PORT_ALL_STATE_FLAGS                                 \
  ((ULONG_PTR)((1UL << PROCESS_EXCEPTION_PORT_ALL_STATE_BITS) - 1))
typedef struct _PROCESS_EXCEPTION_PORT {
  _In_ HANDLE ExceptionPortHandle; 
  _Inout_ ULONG StateFlags; 
} PROCESS_EXCEPTION_PORT, *PPROCESS_EXCEPTION_PORT;
typedef struct _PROCESS_ACCESS_TOKEN {
  HANDLE Token;  
  HANDLE Thread; 
} PROCESS_ACCESS_TOKEN, *PPROCESS_ACCESS_TOKEN;
typedef struct _PROCESS_LDT_INFORMATION {
  ULONG Start;
  ULONG Length;
  LDT_ENTRY LdtEntries[1];
} PROCESS_LDT_INFORMATION, *PPROCESS_LDT_INFORMATION;
typedef struct _PROCESS_LDT_SIZE {
  ULONG Length;
} PROCESS_LDT_SIZE, *PPROCESS_LDT_SIZE;
typedef struct _PROCESS_WS_WATCH_INFORMATION {
  PVOID FaultingPc;
  PVOID FaultingVa;
} PROCESS_WS_WATCH_INFORMATION, *PPROCESS_WS_WATCH_INFORMATION;
#endif
typedef struct _PROCESS_WS_WATCH_INFORMATION_EX {
  PROCESS_WS_WATCH_INFORMATION BasicInfo;
  ULONG_PTR FaultingThreadId;
  ULONG_PTR Flags;
} PROCESS_WS_WATCH_INFORMATION_EX, *PPROCESS_WS_WATCH_INFORMATION_EX;
#define PROCESS_PRIORITY_CLASS_UNKNOWN 0
#define PROCESS_PRIORITY_CLASS_IDLE 1
#define PROCESS_PRIORITY_CLASS_NORMAL 2
#define PROCESS_PRIORITY_CLASS_HIGH 3
#define PROCESS_PRIORITY_CLASS_REALTIME 4
#define PROCESS_PRIORITY_CLASS_BELOW_NORMAL 5
#define PROCESS_PRIORITY_CLASS_ABOVE_NORMAL 6
typedef struct _PROCESS_PRIORITY_CLASS {
  BOOLEAN Foreground;
  UCHAR PriorityClass;
} PROCESS_PRIORITY_CLASS, *PPROCESS_PRIORITY_CLASS;
typedef struct _PROCESS_FOREGROUND_BACKGROUND {
  BOOLEAN Foreground;
} PROCESS_FOREGROUND_BACKGROUND, *PPROCESS_FOREGROUND_BACKGROUND;
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _PROCESS_DEVICEMAP_INFORMATION {
  union {
    struct {
      HANDLE DirectoryHandle;
    } Set;
    struct {
      ULONG DriveMap;
      UCHAR DriveType[32];
    } Query;
  };
} PROCESS_DEVICEMAP_INFORMATION, *PPROCESS_DEVICEMAP_INFORMATION;
#define PROCESS_LUID_DOSDEVICES_ONLY 0x00000001
typedef struct _PROCESS_DEVICEMAP_INFORMATION_EX {
  union {
    struct {
      HANDLE DirectoryHandle;
    } Set;
    struct {
      ULONG DriveMap;
      UCHAR DriveType[32];
    } Query;
  };
  ULONG Flags; 
} PROCESS_DEVICEMAP_INFORMATION_EX, *PPROCESS_DEVICEMAP_INFORMATION_EX;
typedef struct _PROCESS_SESSION_INFORMATION {
  ULONG SessionId;
} PROCESS_SESSION_INFORMATION, *PPROCESS_SESSION_INFORMATION;
#define PROCESS_HANDLE_EXCEPTIONS_ENABLED 0x00000001
#define PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_DISABLED        \
  0x00000000
#define PROCESS_HANDLE_RAISE_EXCEPTION_ON_INVALID_HANDLE_CLOSE_ENABLED         \
  0x00000001
typedef struct _PROCESS_HANDLE_TRACING_ENABLE {
  ULONG Flags;
} PROCESS_HANDLE_TRACING_ENABLE, *PPROCESS_HANDLE_TRACING_ENABLE;
#define PROCESS_HANDLE_TRACING_MAX_SLOTS 0x20000
typedef struct _PROCESS_HANDLE_TRACING_ENABLE_EX {
  ULONG Flags;
  ULONG TotalSlots;
} PROCESS_HANDLE_TRACING_ENABLE_EX, *PPROCESS_HANDLE_TRACING_ENABLE_EX;
#define PROCESS_HANDLE_TRACING_MAX_STACKS 16
#define PROCESS_HANDLE_TRACE_TYPE_OPEN 1
#define PROCESS_HANDLE_TRACE_TYPE_CLOSE 2
#define PROCESS_HANDLE_TRACE_TYPE_BADREF 3
typedef struct _PROCESS_HANDLE_TRACING_ENTRY {
  HANDLE Handle;
  CLIENT_ID ClientId;
  ULONG Type;
  PVOID Stacks[PROCESS_HANDLE_TRACING_MAX_STACKS];
} PROCESS_HANDLE_TRACING_ENTRY, *PPROCESS_HANDLE_TRACING_ENTRY;
typedef struct _PROCESS_HANDLE_TRACING_QUERY {
  HANDLE Handle;
  ULONG TotalTraces;
  PROCESS_HANDLE_TRACING_ENTRY HandleTrace[1];
} PROCESS_HANDLE_TRACING_QUERY, *PPROCESS_HANDLE_TRACING_QUERY;
#endif
typedef struct _THREAD_TLS_INFORMATION {
  ULONG Flags;
  PVOID NewTlsData;
  PVOID OldTlsData;
  HANDLE ThreadId;
} THREAD_TLS_INFORMATION, *PTHREAD_TLS_INFORMATION;
typedef enum _PROCESS_TLS_INFORMATION_TYPE {
  ProcessTlsReplaceIndex,
  ProcessTlsReplaceVector,
  MaxProcessTlsOperation
} PROCESS_TLS_INFORMATION_TYPE,
    *PPROCESS_TLS_INFORMATION_TYPE;
typedef struct _PROCESS_TLS_INFORMATION {
  ULONG Flags;
  ULONG OperationType;
  ULONG ThreadDataCount;
  ULONG TlsIndex;
  ULONG PreviousCount;
  THREAD_TLS_INFORMATION ThreadData[1];
} PROCESS_TLS_INFORMATION, *PPROCESS_TLS_INFORMATION;
typedef struct _PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION {
  ULONG Version;
  ULONG Reserved;
  PVOID Callback;
} PROCESS_INSTRUMENTATION_CALLBACK_INFORMATION,
    *PPROCESS_INSTRUMENTATION_CALLBACK_INFORMATION;
typedef struct _PROCESS_STACK_ALLOCATION_INFORMATION {
  SIZE_T ReserveSize;
  SIZE_T ZeroBits;
  PVOID StackBase;
} PROCESS_STACK_ALLOCATION_INFORMATION, *PPROCESS_STACK_ALLOCATION_INFORMATION;
typedef struct _PROCESS_STACK_ALLOCATION_INFORMATION_EX {
  ULONG PreferredNode;
  ULONG Reserved0;
  ULONG Reserved1;
  ULONG Reserved2;
  PROCESS_STACK_ALLOCATION_INFORMATION AllocInfo;
} PROCESS_STACK_ALLOCATION_INFORMATION_EX,
    *PPROCESS_STACK_ALLOCATION_INFORMATION_EX;
typedef union _PROCESS_AFFINITY_UPDATE_MODE {
  ULONG Flags;
  struct {
    ULONG EnableAutoUpdate : 1;
    ULONG Permanent : 1;
    ULONG Reserved : 30;
  };
} PROCESS_AFFINITY_UPDATE_MODE, *PPROCESS_AFFINITY_UPDATE_MODE;
typedef union _PROCESS_MEMORY_ALLOCATION_MODE {
  ULONG Flags;
  struct {
    ULONG TopDown : 1;
    ULONG Reserved : 31;
  };
} PROCESS_MEMORY_ALLOCATION_MODE, *PPROCESS_MEMORY_ALLOCATION_MODE;
typedef struct _PROCESS_HANDLE_INFORMATION {
  ULONG HandleCount;
  ULONG HandleCountHighWatermark;
} PROCESS_HANDLE_INFORMATION, *PPROCESS_HANDLE_INFORMATION;
typedef struct _PROCESS_CYCLE_TIME_INFORMATION {
  ULONGLONG AccumulatedCycles;
  ULONGLONG CurrentCycleCount;
} PROCESS_CYCLE_TIME_INFORMATION, *PPROCESS_CYCLE_TIME_INFORMATION;
typedef struct _PROCESS_WINDOW_INFORMATION {
  ULONG WindowFlags;
  USHORT WindowTitleLength;
  WCHAR WindowTitle[1];
} PROCESS_WINDOW_INFORMATION, *PPROCESS_WINDOW_INFORMATION;
typedef struct _PROCESS_HANDLE_TABLE_ENTRY_INFO {
  HANDLE HandleValue;
  ULONG_PTR HandleCount;
  ULONG_PTR PointerCount;
  ULONG GrantedAccess;
  ULONG ObjectTypeIndex;
  ULONG HandleAttributes;
  ULONG Reserved;
} PROCESS_HANDLE_TABLE_ENTRY_INFO, *PPROCESS_HANDLE_TABLE_ENTRY_INFO;
typedef struct _PROCESS_HANDLE_SNAPSHOT_INFORMATION {
  ULONG_PTR NumberOfHandles;
  ULONG_PTR Reserved;
  PROCESS_HANDLE_TABLE_ENTRY_INFO Handles[1];
} PROCESS_HANDLE_SNAPSHOT_INFORMATION, *PPROCESS_HANDLE_SNAPSHOT_INFORMATION;
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if !defined(NTDDI_WIN10_CO) ||                                                \
    (NTDDI_VERSION < NTDDI_WIN10_CO) && !PHNT_PATCH_FOR_HYPERDBG
typedef struct _PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY {
  union {
    ULONG Flags;
    struct {
      ULONG EnforceRedirectionTrust : 1;
      ULONG AuditRedirectionTrust : 1;
      ULONG ReservedFlags : 30;
    };
  };
} PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY,
    *PPROCESS_MITIGATION_REDIRECTION_TRUST_POLICY;
#endif
typedef struct _PROCESS_MITIGATION_POLICY_INFORMATION {
  PROCESS_MITIGATION_POLICY Policy;
  union {
    PROCESS_MITIGATION_ASLR_POLICY ASLRPolicy;
    PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY StrictHandleCheckPolicy;
    PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY SystemCallDisablePolicy;
    PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY
    ExtensionPointDisablePolicy;
    PROCESS_MITIGATION_DYNAMIC_CODE_POLICY DynamicCodePolicy;
    PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY ControlFlowGuardPolicy;
    PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY SignaturePolicy;
    PROCESS_MITIGATION_FONT_DISABLE_POLICY FontDisablePolicy;
    PROCESS_MITIGATION_IMAGE_LOAD_POLICY ImageLoadPolicy;
    PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY SystemCallFilterPolicy;
    PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY PayloadRestrictionPolicy;
    PROCESS_MITIGATION_CHILD_PROCESS_POLICY ChildProcessPolicy;
    PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY SideChannelIsolationPolicy;
    PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY UserShadowStackPolicy;
    PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY RedirectionTrustPolicy;
  };
} PROCESS_MITIGATION_POLICY_INFORMATION,
    *PPROCESS_MITIGATION_POLICY_INFORMATION;
typedef struct _PROCESS_KEEPALIVE_COUNT_INFORMATION {
  ULONG WakeCount;
  ULONG NoWakeCount;
} PROCESS_KEEPALIVE_COUNT_INFORMATION, *PPROCESS_KEEPALIVE_COUNT_INFORMATION;
typedef struct _PROCESS_REVOKE_FILE_HANDLES_INFORMATION {
  UNICODE_STRING TargetDevicePath;
} PROCESS_REVOKE_FILE_HANDLES_INFORMATION,
    *PPROCESS_REVOKE_FILE_HANDLES_INFORMATION;
typedef enum _PROCESS_WORKING_SET_OPERATION {
  ProcessWorkingSetSwap,
  ProcessWorkingSetEmpty,
  ProcessWorkingSetOperationMax
} PROCESS_WORKING_SET_OPERATION;
typedef struct _PROCESS_WORKING_SET_CONTROL {
  ULONG Version;
  PROCESS_WORKING_SET_OPERATION Operation;
  ULONG Flags;
} PROCESS_WORKING_SET_CONTROL, *PPROCESS_WORKING_SET_CONTROL;
typedef enum _PS_PROTECTED_TYPE {
  PsProtectedTypeNone,
  PsProtectedTypeProtectedLight,
  PsProtectedTypeProtected,
  PsProtectedTypeMax
} PS_PROTECTED_TYPE;
typedef enum _PS_PROTECTED_SIGNER {
  PsProtectedSignerNone,
  PsProtectedSignerAuthenticode,
  PsProtectedSignerCodeGen,
  PsProtectedSignerAntimalware,
  PsProtectedSignerLsa,
  PsProtectedSignerWindows,
  PsProtectedSignerWinTcb,
  PsProtectedSignerWinSystem,
  PsProtectedSignerApp,
  PsProtectedSignerMax
} PS_PROTECTED_SIGNER;
#define PS_PROTECTED_SIGNER_MASK 0xFF
#define PS_PROTECTED_AUDIT_MASK 0x08
#define PS_PROTECTED_TYPE_MASK 0x07
#define PsProtectedValue(aSigner, aAudit, aType)                               \
  (((aSigner & PS_PROTECTED_SIGNER_MASK) << 4) |                               \
   ((aAudit & PS_PROTECTED_AUDIT_MASK) << 3) |                                 \
   (aType & PS_PROTECTED_TYPE_MASK))
#define InitializePsProtection(aProtectionLevelPtr, aSigner, aAudit, aType)    \
  {                                                                            \
    (aProtectionLevelPtr)->Signer = aSigner;                                   \
    (aProtectionLevelPtr)->Audit = aAudit;                                     \
    (aProtectionLevelPtr)->Type = aType;                                       \
  }
typedef struct _PS_PROTECTION {
  union {
    UCHAR Level;
    struct {
      UCHAR Type : 3;
      UCHAR Audit : 1;
      UCHAR Signer : 4;
    };
  };
} PS_PROTECTION, *PPS_PROTECTION;
typedef struct _PROCESS_FAULT_INFORMATION {
  ULONG FaultFlags;
  ULONG AdditionalInfo;
} PROCESS_FAULT_INFORMATION, *PPROCESS_FAULT_INFORMATION;
typedef struct _PROCESS_TELEMETRY_ID_INFORMATION {
  ULONG HeaderSize;
  ULONG ProcessId;
  ULONGLONG ProcessStartKey;
  ULONGLONG CreateTime;
  ULONGLONG CreateInterruptTime;
  ULONGLONG CreateUnbiasedInterruptTime;
  ULONGLONG ProcessSequenceNumber;
  ULONGLONG SessionCreateTime;
  ULONG SessionId;
  ULONG BootId;
  ULONG ImageChecksum;
  ULONG ImageTimeDateStamp;
  ULONG UserSidOffset;
  ULONG ImagePathOffset;
  ULONG PackageNameOffset;
  ULONG RelativeAppNameOffset;
  ULONG CommandLineOffset;
} PROCESS_TELEMETRY_ID_INFORMATION, *PPROCESS_TELEMETRY_ID_INFORMATION;
typedef struct _PROCESS_COMMIT_RELEASE_INFORMATION {
  ULONG Version;
  struct {
    ULONG Eligible : 1;
    ULONG ReleaseRepurposedMemResetCommit : 1;
    ULONG ForceReleaseMemResetCommit : 1;
    ULONG Spare : 29;
  };
  SIZE_T CommitDebt;
  SIZE_T CommittedMemResetSize;
  SIZE_T RepurposedMemResetSize;
} PROCESS_COMMIT_RELEASE_INFORMATION, *PPROCESS_COMMIT_RELEASE_INFORMATION;
typedef struct _PROCESS_JOB_MEMORY_INFO {
  ULONGLONG SharedCommitUsage;
  ULONGLONG PrivateCommitUsage;
  ULONGLONG PeakPrivateCommitUsage;
  ULONGLONG PrivateCommitLimit;
  ULONGLONG TotalCommitLimit;
} PROCESS_JOB_MEMORY_INFO, *PPROCESS_JOB_MEMORY_INFO;
typedef struct _PROCESS_CHILD_PROCESS_INFORMATION {
  BOOLEAN ProhibitChildProcesses;
  BOOLEAN AlwaysAllowSecureChildProcess; 
  BOOLEAN AuditProhibitChildProcesses;
} PROCESS_CHILD_PROCESS_INFORMATION, *PPROCESS_CHILD_PROCESS_INFORMATION;
#define POWER_THROTTLING_PROCESS_CURRENT_VERSION 1
#define POWER_THROTTLING_PROCESS_EXECUTION_SPEED 0x1
#define POWER_THROTTLING_PROCESS_DELAYTIMERS 0x2
#define POWER_THROTTLING_PROCESS_VALID_FLAGS                                   \
  ((POWER_THROTTLING_PROCESS_EXECUTION_SPEED |                                 \
    POWER_THROTTLING_PROCESS_DELAYTIMERS))
typedef struct _POWER_THROTTLING_PROCESS_STATE {
  ULONG Version;
  ULONG ControlMask;
  ULONG StateMask;
} POWER_THROTTLING_PROCESS_STATE, *PPOWER_THROTTLING_PROCESS_STATE;
#define WIN32K_SYSCALL_FILTER_STATE_ENABLE 0x1
#define WIN32K_SYSCALL_FILTER_STATE_AUDIT 0x2
typedef struct _WIN32K_SYSCALL_FILTER {
  ULONG FilterState;
  ULONG FilterSet;
} WIN32K_SYSCALL_FILTER, *PWIN32K_SYSCALL_FILTER;
typedef struct _PROCESS_WAKE_INFORMATION {
  ULONGLONG NotificationChannel;
  ULONG WakeCounters[7];
  struct _JOBOBJECT_WAKE_FILTER *WakeFilter;
} PROCESS_WAKE_INFORMATION, *PPROCESS_WAKE_INFORMATION;
typedef struct _PROCESS_ENERGY_TRACKING_STATE {
  ULONG StateUpdateMask;
  ULONG StateDesiredValue;
  ULONG StateSequence;
  ULONG UpdateTag : 1;
  WCHAR Tag[64];
} PROCESS_ENERGY_TRACKING_STATE, *PPROCESS_ENERGY_TRACKING_STATE;
typedef struct _MANAGE_WRITES_TO_EXECUTABLE_MEMORY {
  ULONG Version : 8;
  ULONG ProcessEnableWriteExceptions : 1;
  ULONG ThreadAllowWrites : 1;
  ULONG Spare : 22;
  PVOID KernelWriteToExecutableSignal; 
} MANAGE_WRITES_TO_EXECUTABLE_MEMORY, *PMANAGE_WRITES_TO_EXECUTABLE_MEMORY;
#define POWER_THROTTLING_THREAD_CURRENT_VERSION 1
#define POWER_THROTTLING_THREAD_EXECUTION_SPEED 0x1
#define POWER_THROTTLING_THREAD_VALID_FLAGS                                    \
  (POWER_THROTTLING_THREAD_EXECUTION_SPEED)
typedef struct _POWER_THROTTLING_THREAD_STATE {
  ULONG Version;
  ULONG ControlMask;
  ULONG StateMask;
} POWER_THROTTLING_THREAD_STATE, *PPOWER_THROTTLING_THREAD_STATE;
#define PROCESS_READWRITEVM_LOGGING_ENABLE_READVM 1
#define PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM 2
#define PROCESS_READWRITEVM_LOGGING_ENABLE_READVM_V 1UL
#define PROCESS_READWRITEVM_LOGGING_ENABLE_WRITEVM_V 2UL
typedef union _PROCESS_READWRITEVM_LOGGING_INFORMATION {
  UCHAR Flags;
  struct {
    UCHAR EnableReadVmLogging : 1;
    UCHAR EnableWriteVmLogging : 1;
    UCHAR Unused : 6;
  };
} PROCESS_READWRITEVM_LOGGING_INFORMATION,
    *PPROCESS_READWRITEVM_LOGGING_INFORMATION;
typedef struct _PROCESS_UPTIME_INFORMATION {
  ULONGLONG QueryInterruptTime;
  ULONGLONG QueryUnbiasedTime;
  ULONGLONG EndInterruptTime;
  ULONGLONG TimeSinceCreation;
  ULONGLONG Uptime;
  ULONGLONG SuspendedTime;
  union {
    ULONG HangCount : 4;
    ULONG GhostCount : 4;
    ULONG Crashed : 1;
    ULONG Terminated : 1;
  };
} PROCESS_UPTIME_INFORMATION, *PPROCESS_UPTIME_INFORMATION;
typedef union _PROCESS_SYSTEM_RESOURCE_MANAGEMENT {
  ULONG Flags;
  struct {
    ULONG Foreground : 1;
    ULONG Reserved : 31;
  };
} PROCESS_SYSTEM_RESOURCE_MANAGEMENT, *PPROCESS_SYSTEM_RESOURCE_MANAGEMENT;
typedef struct _PROCESS_SECURITY_DOMAIN_INFORMATION {
  ULONGLONG SecurityDomain;
} PROCESS_SECURITY_DOMAIN_INFORMATION, *PPROCESS_SECURITY_DOMAIN_INFORMATION;
typedef struct _PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION {
  HANDLE ProcessHandle;
} PROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION,
    *PPROCESS_COMBINE_SECURITY_DOMAINS_INFORMATION;
typedef union _PROCESS_LOGGING_INFORMATION {
  ULONG Flags;
  struct {
    ULONG EnableReadVmLogging : 1;
    ULONG EnableWriteVmLogging : 1;
    ULONG EnableProcessSuspendResumeLogging : 1;
    ULONG EnableThreadSuspendResumeLogging : 1;
    ULONG Reserved : 28;
  };
} PROCESS_LOGGING_INFORMATION, *PPROCESS_LOGGING_INFORMATION;
typedef struct _PROCESS_LEAP_SECOND_INFORMATION {
  ULONG Flags;
  ULONG Reserved;
} PROCESS_LEAP_SECOND_INFORMATION, *PPROCESS_LEAP_SECOND_INFORMATION;
typedef struct _PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION {
  ULONGLONG ReserveSize;
  ULONGLONG CommitSize;
  ULONG PreferredNode;
  ULONG Reserved;
  PVOID Ssp;
} PROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION,
    *PPROCESS_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION;
typedef struct _PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION {
  PVOID Ssp;
} PROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION,
    *PPROCESS_FREE_FIBER_SHADOW_STACK_ALLOCATION_INFORMATION;
#endif
typedef struct _THREAD_BASIC_INFORMATION {
  NTSTATUS ExitStatus;
  PTEB TebBaseAddress;
  CLIENT_ID ClientId;
  KAFFINITY AffinityMask;
  KPRIORITY Priority;
  KPRIORITY BasePriority;
} THREAD_BASIC_INFORMATION, *PTHREAD_BASIC_INFORMATION;
typedef struct _THREAD_LAST_SYSCALL_INFORMATION {
  PVOID FirstArgument;
  USHORT SystemCallNumber;
#ifdef WIN64
  USHORT Pad[0x3]; 
#else
  USHORT Pad[0x1]; 
#endif
  ULONG64 WaitTime;
} THREAD_LAST_SYSCALL_INFORMATION, *PTHREAD_LAST_SYSCALL_INFORMATION;
typedef struct _THREAD_CYCLE_TIME_INFORMATION {
  ULONGLONG AccumulatedCycles;
  ULONGLONG CurrentCycleCount;
} THREAD_CYCLE_TIME_INFORMATION, *PTHREAD_CYCLE_TIME_INFORMATION;
typedef struct _THREAD_TEB_INFORMATION {
  PVOID TebInformation; 
  ULONG TebOffset;      
  ULONG BytesToRead;    
} THREAD_TEB_INFORMATION, *PTHREAD_TEB_INFORMATION;
typedef struct _COUNTER_READING {
  HARDWARE_COUNTER_TYPE Type;
  ULONG Index;
  ULONG64 Start;
  ULONG64 Total;
} COUNTER_READING, *PCOUNTER_READING;
typedef struct _THREAD_PERFORMANCE_DATA {
  USHORT Size;
  USHORT Version;
  PROCESSOR_NUMBER ProcessorNumber;
  ULONG ContextSwitches;
  ULONG HwCountersCount;
  ULONG64 UpdateCount;
  ULONG64 WaitReasonBitMap;
  ULONG64 HardwareCounters;
  COUNTER_READING CycleTime;
  COUNTER_READING HwCounters[MAX_HW_COUNTERS];
} THREAD_PERFORMANCE_DATA, *PTHREAD_PERFORMANCE_DATA;
typedef struct _THREAD_PROFILING_INFORMATION {
  ULONG64 HardwareCounters;
  ULONG Flags;
  ULONG Enable;
  PTHREAD_PERFORMANCE_DATA PerformanceData;
} THREAD_PROFILING_INFORMATION, *PTHREAD_PROFILING_INFORMATION;
typedef struct _RTL_UMS_CONTEXT {
  SINGLE_LIST_ENTRY Link;
  CONTEXT Context;
  PVOID Teb;
  PVOID UserContext;
  volatile ULONG ScheduledThread : 1;
  volatile ULONG Suspended : 1;
  volatile ULONG VolatileContext : 1;
  volatile ULONG Terminated : 1;
  volatile ULONG DebugActive : 1;
  volatile ULONG RunningOnSelfThread : 1;
  volatile ULONG DenyRunningOnSelfThread : 1;
  volatile LONG Flags;
  volatile ULONG64 KernelUpdateLock : 2;
  volatile ULONG64 PrimaryClientID : 62;
  volatile ULONG64 ContextLock;
  struct _RTL_UMS_CONTEXT *PrimaryUmsContext;
  ULONG SwitchCount;
  ULONG KernelYieldCount;
  ULONG MixedYieldCount;
  ULONG YieldCount;
} RTL_UMS_CONTEXT, *PRTL_UMS_CONTEXT;
typedef enum _THREAD_UMS_INFORMATION_COMMAND {
  UmsInformationCommandInvalid,
  UmsInformationCommandAttach,
  UmsInformationCommandDetach,
  UmsInformationCommandQuery
} THREAD_UMS_INFORMATION_COMMAND;
typedef struct _RTL_UMS_COMPLETION_LIST {
  PSINGLE_LIST_ENTRY ThreadListHead;
  PVOID CompletionEvent;
  ULONG CompletionFlags;
  SINGLE_LIST_ENTRY InternalListHead;
} RTL_UMS_COMPLETION_LIST, *PRTL_UMS_COMPLETION_LIST;
typedef struct _THREAD_UMS_INFORMATION {
  THREAD_UMS_INFORMATION_COMMAND Command;
  PRTL_UMS_COMPLETION_LIST CompletionList;
  PRTL_UMS_CONTEXT UmsContext;
  union {
    ULONG Flags;
    struct {
      ULONG IsUmsSchedulerThread : 1;
      ULONG IsUmsWorkerThread : 1;
      ULONG SpareBits : 30;
    };
  };
} THREAD_UMS_INFORMATION, *PTHREAD_UMS_INFORMATION;
typedef struct _THREAD_NAME_INFORMATION {
  UNICODE_STRING ThreadName;
} THREAD_NAME_INFORMATION, *PTHREAD_NAME_INFORMATION;
typedef struct _ALPC_WORK_ON_BEHALF_TICKET {
  ULONG ThreadId;
  ULONG ThreadCreationTimeLow;
} ALPC_WORK_ON_BEHALF_TICKET, *PALPC_WORK_ON_BEHALF_TICKET;
typedef struct _RTL_WORK_ON_BEHALF_TICKET_EX {
  ALPC_WORK_ON_BEHALF_TICKET Ticket;
  union {
    ULONG Flags;
    struct {
      ULONG CurrentThread : 1;
      ULONG Reserved1 : 31;
    };
  };
  ULONG Reserved2;
} RTL_WORK_ON_BEHALF_TICKET_EX, *PRTL_WORK_ON_BEHALF_TICKET_EX;
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _SUBSYSTEM_INFORMATION_TYPE {
  SubsystemInformationTypeWin32,
  SubsystemInformationTypeWSL,
  MaxSubsystemInformationType
} SUBSYSTEM_INFORMATION_TYPE;
#endif
typedef enum _THREAD_WORKLOAD_CLASS {
  ThreadWorkloadClassDefault,
  ThreadWorkloadClassGraphics,
  MaxThreadWorkloadClass
} THREAD_WORKLOAD_CLASS;
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcess(_Out_ PHANDLE ProcessHandle, _In_ ACCESS_MASK DesiredAccess,
                _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
                _In_ HANDLE ParentProcess, _In_ BOOLEAN InheritObjectTable,
                _In_opt_ HANDLE SectionHandle, _In_opt_ HANDLE DebugPort,
                _In_opt_ HANDLE TokenHandle);
#define PROCESS_CREATE_FLAGS_BREAKAWAY                                         \
  0x00000001 
#define PROCESS_CREATE_FLAGS_NO_DEBUG_INHERIT                                  \
  0x00000002 
#define PROCESS_CREATE_FLAGS_INHERIT_HANDLES                                   \
  0x00000004 
#define PROCESS_CREATE_FLAGS_OVERRIDE_ADDRESS_SPACE                            \
  0x00000008 
#define PROCESS_CREATE_FLAGS_LARGE_PAGES                                       \
  0x00000010 
#define PROCESS_CREATE_FLAGS_LARGE_PAGE_SYSTEM_DLL                             \
  0x00000020 
#define PROCESS_CREATE_FLAGS_PROTECTED_PROCESS                                 \
  0x00000040 
#define PROCESS_CREATE_FLAGS_CREATE_SESSION                                    \
  0x00000080 
#define PROCESS_CREATE_FLAGS_INHERIT_FROM_PARENT                               \
  0x00000100 
#define PROCESS_CREATE_FLAGS_SUSPENDED                                         \
  0x00000200 
#define PROCESS_CREATE_FLAGS_FORCE_BREAKAWAY                                   \
  0x00000400 
#define PROCESS_CREATE_FLAGS_MINIMAL_PROCESS                                   \
  0x00000800 
#define PROCESS_CREATE_FLAGS_RELEASE_SECTION                                   \
  0x00001000 
#define PROCESS_CREATE_FLAGS_CLONE_MINIMAL 0x00002000 
#define PROCESS_CREATE_FLAGS_CLONE_MINIMAL_REDUCED_COMMIT 0x00004000 //
#define PROCESS_CREATE_FLAGS_AUXILIARY_PROCESS                                 \
  0x00008000 
#define PROCESS_CREATE_FLAGS_CREATE_STORE                                      \
  0x00020000 
#define PROCESS_CREATE_FLAGS_USE_PROTECTED_ENVIRONMENT                         \
  0x00040000 
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessEx(_Out_ PHANDLE ProcessHandle, _In_ ACCESS_MASK DesiredAccess,
                  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
                  _In_ HANDLE ParentProcess,
                  _In_ ULONG Flags, 
                  _In_opt_ HANDLE SectionHandle, _In_opt_ HANDLE DebugPort,
                  _In_opt_ HANDLE TokenHandle,
                  _Reserved_ ULONG Reserved 
);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenProcess(_Out_ PHANDLE ProcessHandle, _In_ ACCESS_MASK DesiredAccess,
              _In_ POBJECT_ATTRIBUTES ObjectAttributes,
              _In_opt_ PCLIENT_ID ClientId);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateProcess(_In_opt_ HANDLE ProcessHandle, _In_ NTSTATUS ExitStatus);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendProcess(_In_ HANDLE ProcessHandle);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeProcess(_In_ HANDLE ProcessHandle);
#define NtCurrentProcess() ((HANDLE)(LONG_PTR)-1)
#define ZwCurrentProcess() NtCurrentProcess()
#define NtCurrentThread() ((HANDLE)(LONG_PTR)-2)
#define ZwCurrentThread() NtCurrentThread()
#define NtCurrentSession() ((HANDLE)(LONG_PTR)-3)
#define ZwCurrentSession() NtCurrentSession()
#define NtCurrentPeb() (NtCurrentTeb()->ProcessEnvironmentBlock)
#define NtCurrentProcessToken()                                                \
  ((HANDLE)(LONG_PTR)-4) 
#define NtCurrentThreadToken()                                                 \
  ((HANDLE)(LONG_PTR)-5) 
#define NtCurrentThreadEffectiveToken()                                        \
  ((HANDLE)(LONG_PTR)-6) 
#define NtCurrentSilo() ((HANDLE)(LONG_PTR)-1)
#define NtCurrentProcessId() (NtCurrentTeb()->ClientId.UniqueProcess)
#define NtCurrentThreadId() (NtCurrentTeb()->ClientId.UniqueThread)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationProcess(_In_ HANDLE ProcessHandle,
                          _In_ PROCESSINFOCLASS ProcessInformationClass,
                          _Out_writes_bytes_(ProcessInformationLength)
                              PVOID ProcessInformation,
                          _In_ ULONG ProcessInformationLength,
                          _Out_opt_ PULONG ReturnLength);
#if (PHNT_VERSION >= PHNT_WS03)
#define PROCESS_GET_NEXT_FLAGS_PREVIOUS_PROCESS 0x00000001
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextProcess(_In_opt_ HANDLE ProcessHandle, _In_ ACCESS_MASK DesiredAccess,
                 _In_ ULONG HandleAttributes, _In_ ULONG Flags,
                 _Out_ PHANDLE NewProcessHandle);
#endif
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetNextThread(_In_ HANDLE ProcessHandle, _In_opt_ HANDLE ThreadHandle,
                _In_ ACCESS_MASK DesiredAccess, _In_ ULONG HandleAttributes,
                _In_ ULONG Flags, _Out_ PHANDLE NewThreadHandle);
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationProcess(_In_ HANDLE ProcessHandle,
                        _In_ PROCESSINFOCLASS ProcessInformationClass,
                        _In_reads_bytes_(ProcessInformationLength)
                            PVOID ProcessInformation,
                        _In_ ULONG ProcessInformationLength);
#endif
#define STATECHANGE_SET_ATTRIBUTES 0x0001
typedef enum _PROCESS_STATE_CHANGE_TYPE {
  ProcessStateChangeSuspend,
  ProcessStateChangeResume,
  ProcessStateChangeMax,
} PROCESS_STATE_CHANGE_TYPE,
    *PPROCESS_STATE_CHANGE_TYPE;
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateProcessStateChange(_Out_ PHANDLE ProcessStateChangeHandle,
                           _In_ ACCESS_MASK DesiredAccess,
                           _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
                           _In_ HANDLE ProcessHandle,
                           _In_opt_ ULONG64 Reserved);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeProcessState(_In_ HANDLE ProcessStateChangeHandle,
                     _In_ HANDLE ProcessHandle,
                     _In_ PROCESS_STATE_CHANGE_TYPE StateChangeType,
                     _In_opt_ PVOID ExtendedInformation,
                     _In_opt_ SIZE_T ExtendedInformationLength,
                     _In_opt_ ULONG64 Reserved);
#endif
typedef enum _THREAD_STATE_CHANGE_TYPE {
  ThreadStateChangeSuspend,
  ThreadStateChangeResume,
  ThreadStateChangeMax,
} THREAD_STATE_CHANGE_TYPE,
    *PTHREAD_STATE_CHANGE_TYPE;
#if (PHNT_VERSION >= PHNT_WIN11)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadStateChange(_Out_ PHANDLE ThreadStateChangeHandle,
                          _In_ ACCESS_MASK DesiredAccess,
                          _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
                          _In_ HANDLE ThreadHandle, _In_opt_ ULONG64 Reserved);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtChangeThreadState(_In_ HANDLE ThreadStateChangeHandle,
                    _In_ HANDLE ThreadHandle,
                    _In_ THREAD_STATE_CHANGE_TYPE StateChangeType,
                    _In_opt_ PVOID ExtendedInformation,
                    _In_opt_ SIZE_T ExtendedInformationLength,
                    _In_opt_ ULONG64 Reserved);
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThread(_Out_ PHANDLE ThreadHandle, _In_ ACCESS_MASK DesiredAccess,
               _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
               _In_ HANDLE ProcessHandle, _Out_ PCLIENT_ID ClientId,
               _In_ PCONTEXT ThreadContext, _In_ PINITIAL_TEB InitialTeb,
               _In_ BOOLEAN CreateSuspended);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenThread(_Out_ PHANDLE ThreadHandle, _In_ ACCESS_MASK DesiredAccess,
             _In_ POBJECT_ATTRIBUTES ObjectAttributes,
             _In_opt_ PCLIENT_ID ClientId);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateThread(_In_opt_ HANDLE ThreadHandle, _In_ NTSTATUS ExitStatus);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSuspendThread(_In_ HANDLE ThreadHandle,
                _Out_opt_ PULONG PreviousSuspendCount);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtResumeThread(_In_ HANDLE ThreadHandle, _Out_opt_ PULONG PreviousSuspendCount);
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumber(VOID);
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
ULONG
NTAPI
NtGetCurrentProcessorNumberEx(_Out_opt_ PPROCESSOR_NUMBER ProcessorNumber);
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetContextThread(_In_ HANDLE ThreadHandle, _Inout_ PCONTEXT ThreadContext);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetContextThread(_In_ HANDLE ThreadHandle, _In_ PCONTEXT ThreadContext);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationThread(_In_ HANDLE ThreadHandle,
                         _In_ THREADINFOCLASS ThreadInformationClass,
                         _Out_writes_bytes_(ThreadInformationLength)
                             PVOID ThreadInformation,
                         _In_ ULONG ThreadInformationLength,
                         _Out_opt_ PULONG ReturnLength);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationThread(_In_ HANDLE ThreadHandle,
                       _In_ THREADINFOCLASS ThreadInformationClass,
                       _In_reads_bytes_(ThreadInformationLength)
                           PVOID ThreadInformation,
                       _In_ ULONG ThreadInformationLength);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThread(_In_ HANDLE ThreadHandle);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertResumeThread(_In_ HANDLE ThreadHandle,
                    _Out_opt_ PULONG PreviousSuspendCount);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTestAlert(VOID);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateThread(_In_ HANDLE ServerThreadHandle,
                    _In_ HANDLE ClientThreadHandle,
                    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRegisterThreadTerminatePort(_In_ HANDLE PortHandle);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetLdtEntries(_In_ ULONG Selector0, _In_ ULONG Entry0Low, _In_ ULONG Entry0Hi,
                _In_ ULONG Selector1, _In_ ULONG Entry1Low,
                _In_ ULONG Entry1Hi);
typedef VOID (*PPS_APC_ROUTINE)(_In_opt_ PVOID ApcArgument1,
                                _In_opt_ PVOID ApcArgument2,
                                _In_opt_ PVOID ApcArgument3);
#define Wow64EncodeApcRoutine(ApcRoutine)                                      \
  ((PVOID)((0 - ((LONG_PTR)(ApcRoutine))) << 2))
#define Wow64DecodeApcRoutine(ApcRoutine)                                      \
  ((PVOID)(0 - (((LONG_PTR)(ApcRoutine)) >> 2)))
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThread(_In_ HANDLE ThreadHandle, _In_ PPS_APC_ROUTINE ApcRoutine,
                 _In_opt_ PVOID ApcArgument1, _In_opt_ PVOID ApcArgument2,
                 _In_opt_ PVOID ApcArgument3);
#if (PHNT_VERSION >= PHNT_WIN7)
#define APC_FORCE_THREAD_SIGNAL ((HANDLE)1) 
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx(_In_ HANDLE ThreadHandle,
                   _In_opt_ HANDLE ReserveHandle, 
                   _In_ PPS_APC_ROUTINE ApcRoutine, _In_opt_ PVOID ApcArgument1,
                   _In_opt_ PVOID ApcArgument2, _In_opt_ PVOID ApcArgument3);
#endif
#if (PHNT_VERSION >= PHNT_WIN11)
#if !defined(NTDDI_WIN10_CO) || (NTDDI_VERSION < NTDDI_WIN10_CO)
typedef enum _QUEUE_USER_APC_FLAGS {
  QUEUE_USER_APC_FLAGS_NONE = 0x0,
  QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC = 0x1,
} QUEUE_USER_APC_FLAGS;
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueueApcThreadEx2(_In_ HANDLE ThreadHandle,
                    _In_opt_ HANDLE ReserveHandle, 
                    _In_ QUEUE_USER_APC_FLAGS ApcFlags,
                    _In_ PPS_APC_ROUTINE ApcRoutine,
                    _In_opt_ PVOID ApcArgument1, _In_opt_ PVOID ApcArgument2,
                    _In_opt_ PVOID ApcArgument3);
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAlertThreadByThreadId(_In_ HANDLE ThreadId);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForAlertByThreadId(_In_ PVOID Address, _In_opt_ PLARGE_INTEGER Timeout);
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define ProcThreadAttributeParentProcess 0 
#define ProcThreadAttributeExtendedFlags                                       \
  1 
#define ProcThreadAttributeHandleList 2     
#define ProcThreadAttributeGroupAffinity 3  
#define ProcThreadAttributePreferredNode 4  
#define ProcThreadAttributeIdealProcessor 5 
#define ProcThreadAttributeUmsThread 6      
#define ProcThreadAttributeMitigationPolicy                                    \
  7                                          
#define ProcThreadAttributePackageFullName 8 
#define ProcThreadAttributeSecurityCapabilities 9 
#define ProcThreadAttributeConsoleReference                                    \
  10 
#define ProcThreadAttributeProtectionLevel                                     \
  11 
#define ProcThreadAttributeOsMaxVersionTested                                  \
  12 
#define ProcThreadAttributeJobList 13 
#define ProcThreadAttributeChildProcessPolicy                                  \
  14 
#define ProcThreadAttributeAllApplicationPackagesPolicy                        \
  15 
#define ProcThreadAttributeWin32kFilter 16 
#define ProcThreadAttributeSafeOpenPromptOriginClaim                           \
  17 
#define ProcThreadAttributeDesktopAppPolicy                                    \
  18 
#define ProcThreadAttributeBnoIsolation                                        \
  19 
#define ProcThreadAttributePseudoConsole 22 
#define ProcThreadAttributeIsolationManifest                                   \
  23 
#define ProcThreadAttributeMitigationAuditPolicy                               \
  24 
#define ProcThreadAttributeMachineType 25     
#define ProcThreadAttributeComponentFilter 26 
#define ProcThreadAttributeEnableOptionalXStateFeatures                        \
  27                                      
#define ProcThreadAttributeCreateStore 28 
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#define PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS                                   \
  ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME
#define PROC_THREAD_ATTRIBUTE_PACKAGE_FULL_NAME                                \
  ProcThreadAttributeValue(ProcThreadAttributePackageFullName, FALSE, TRUE,    \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#define PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE                                \
  ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE,   \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#define PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED                               \
  ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#define PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM                    \
  ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim,       \
                           FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#define PROC_THREAD_ATTRIBUTE_BNO_ISOLATION                                    \
  ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST
#define PROC_THREAD_ATTRIBUTE_ISOLATION_MANIFEST                               \
  ProcThreadAttributeValue(ProcThreadAttributeIsolationManifest, FALSE, TRUE,  \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CREATE_STORE
#define PROC_THREAD_ATTRIBUTE_CREATE_STORE                                     \
  ProcThreadAttributeValue(ProcThreadAttributeCreateStore, FALSE, TRUE, FALSE)
#endif
typedef struct _PROC_THREAD_ATTRIBUTE {
  ULONG_PTR Attribute;
  SIZE_T Size;
  ULONG_PTR Value;
} PROC_THREAD_ATTRIBUTE, *PPROC_THREAD_ATTRIBUTE;
typedef struct _PROC_THREAD_ATTRIBUTE_LIST {
  ULONG PresentFlags;
  ULONG AttributeCount;
  ULONG LastAttribute;
  ULONG SpareUlong0;
  PPROC_THREAD_ATTRIBUTE ExtendedFlagsAttribute;
  PROC_THREAD_ATTRIBUTE Attributes[1];
} PROC_THREAD_ATTRIBUTE_LIST;
#define EXTENDED_PROCESS_CREATION_FLAG_ELEVATION_HANDLED 0x00000001
#define EXTENDED_PROCESS_CREATION_FLAG_FORCELUA 0x00000002
#define EXTENDED_PROCESS_CREATION_FLAG_FORCE_BREAKAWAY                         \
  0x00000004 
#define PROTECTION_LEVEL_WINTCB_LIGHT 0x00000000
#define PROTECTION_LEVEL_WINDOWS 0x00000001
#define PROTECTION_LEVEL_WINDOWS_LIGHT 0x00000002
#define PROTECTION_LEVEL_ANTIMALWARE_LIGHT 0x00000003
#define PROTECTION_LEVEL_LSA_LIGHT 0x00000004
#define PROTECTION_LEVEL_WINTCB 0x00000005
#define PROTECTION_LEVEL_CODEGEN_LIGHT 0x00000006
#define PROTECTION_LEVEL_AUTHENTICODE 0x00000007
typedef enum _SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS {
  SeSafeOpenExperienceNone = 0x00,
  SeSafeOpenExperienceCalled = 0x01,
  SeSafeOpenExperienceAppRepCalled = 0x02,
  SeSafeOpenExperiencePromptDisplayed = 0x04,
  SeSafeOpenExperienceUAC = 0x08,
  SeSafeOpenExperienceUninstaller = 0x10,
  SeSafeOpenExperienceIgnoreUnknownOrBad = 0x20,
  SeSafeOpenExperienceDefenderTrustedInstaller = 0x40,
  SeSafeOpenExperienceMOTWPresent = 0x80
} SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS;
typedef struct _SE_SAFE_OPEN_PROMPT_RESULTS {
  SE_SAFE_OPEN_PROMPT_EXPERIENCE_RESULTS Results;
  WCHAR Path[MAX_PATH];
} SE_SAFE_OPEN_PROMPT_RESULTS, *PSE_SAFE_OPEN_PROMPT_RESULTS;
typedef struct _PROC_THREAD_BNOISOLATION_ATTRIBUTE {
  BOOL IsolationEnabled;
  WCHAR IsolationPrefix[0x88];
} PROC_THREAD_BNOISOLATION_ATTRIBUTE, *PPROC_THREAD_BNOISOLATION_ATTRIBUTE;
typedef struct _ISOLATION_MANIFEST_PROPERTIES {
  UNICODE_STRING InstancePath;
  UNICODE_STRING FriendlyName;
  UNICODE_STRING Description;
  ULONG_PTR Level;
} ISOLATION_MANIFEST_PROPERTIES, *PISOLATION_MANIFEST_PROPERTIES;
#ifndef PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS
#define PROC_THREAD_ATTRIBUTE_EXTENDED_FLAGS                                   \
  ProcThreadAttributeValue(ProcThreadAttributeExtendedFlags, FALSE, TRUE, TRUE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE
#define PROC_THREAD_ATTRIBUTE_CONSOLE_REFERENCE                                \
  ProcThreadAttributeValue(ProcThreadAttributeConsoleReference, FALSE, TRUE,   \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED
#define PROC_THREAD_ATTRIBUTE_OSMAXVERSIONTESTED                               \
  ProcThreadAttributeValue(ProcThreadAttributeOsMaxVersionTested, FALSE, TRUE, \
                           FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM
#define PROC_THREAD_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM                    \
  ProcThreadAttributeValue(ProcThreadAttributeSafeOpenPromptOriginClaim,       \
                           FALSE, TRUE, FALSE)
#endif
#ifndef PROC_THREAD_ATTRIBUTE_BNO_ISOLATION
#define PROC_THREAD_ATTRIBUTE_BNO_ISOLATION                                    \
  ProcThreadAttributeValue(ProcThreadAttributeBnoIsolation, FALSE, TRUE, FALSE)
#endif
typedef enum _PS_ATTRIBUTE_NUM {
  PsAttributeParentProcess,      
  PsAttributeDebugObject,        
  PsAttributeToken,              
  PsAttributeClientId,           
  PsAttributeTebAddress,         
  PsAttributeImageName,          
  PsAttributeImageInfo,          
  PsAttributeMemoryReserve,      
  PsAttributePriorityClass,      
  PsAttributeErrorMode,          
  PsAttributeStdHandleInfo,      
  PsAttributeHandleList,         
  PsAttributeGroupAffinity,      
  PsAttributePreferredNode,      
  PsAttributeIdealProcessor,     
  PsAttributeUmsThread,          
  PsAttributeMitigationOptions,  
  PsAttributeProtectionLevel,    
  PsAttributeSecureProcess,      
  PsAttributeJobList,            
  PsAttributeChildProcessPolicy, 
  PsAttributeAllApplicationPackagesPolicy, 
  PsAttributeWin32kFilter,              
  PsAttributeSafeOpenPromptOriginClaim, 
  PsAttributeBnoIsolation, 
  PsAttributeDesktopAppPolicy, 
  PsAttributeChpe,             
  PsAttributeMitigationAuditOptions, 
  PsAttributeMachineType, 
  PsAttributeComponentFilter,
  PsAttributeEnableOptionalXStateFeatures, 
  PsAttributeMax
} PS_ATTRIBUTE_NUM;
#define PS_ATTRIBUTE_NUMBER_MASK 0x0000ffff
#define PS_ATTRIBUTE_THREAD 0x00010000 
#define PS_ATTRIBUTE_INPUT 0x00020000  
#define PS_ATTRIBUTE_ADDITIVE                                                  \
  0x00040000 
#define PsAttributeValue(Number, Thread, Input, Additive)                      \
  (((Number)&PS_ATTRIBUTE_NUMBER_MASK) |                                       \
   ((Thread) ? PS_ATTRIBUTE_THREAD : 0) | ((Input) ? PS_ATTRIBUTE_INPUT : 0) | \
   ((Additive) ? PS_ATTRIBUTE_ADDITIVE : 0))
#define PS_ATTRIBUTE_PARENT_PROCESS                                            \
  PsAttributeValue(PsAttributeParentProcess, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_DEBUG_OBJECT                                              \
  PsAttributeValue(PsAttributeDebugObject, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_TOKEN PsAttributeValue(PsAttributeToken, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_CLIENT_ID                                                 \
  PsAttributeValue(PsAttributeClientId, TRUE, FALSE, FALSE)
#define PS_ATTRIBUTE_TEB_ADDRESS                                               \
  PsAttributeValue(PsAttributeTebAddress, TRUE, FALSE, FALSE)
#define PS_ATTRIBUTE_IMAGE_NAME                                                \
  PsAttributeValue(PsAttributeImageName, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_IMAGE_INFO                                                \
  PsAttributeValue(PsAttributeImageInfo, FALSE, FALSE, FALSE)
#define PS_ATTRIBUTE_MEMORY_RESERVE                                            \
  PsAttributeValue(PsAttributeMemoryReserve, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_PRIORITY_CLASS                                            \
  PsAttributeValue(PsAttributePriorityClass, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ERROR_MODE                                                \
  PsAttributeValue(PsAttributeErrorMode, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_STD_HANDLE_INFO                                           \
  PsAttributeValue(PsAttributeStdHandleInfo, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_HANDLE_LIST                                               \
  PsAttributeValue(PsAttributeHandleList, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_GROUP_AFFINITY                                            \
  PsAttributeValue(PsAttributeGroupAffinity, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_PREFERRED_NODE                                            \
  PsAttributeValue(PsAttributePreferredNode, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_IDEAL_PROCESSOR                                           \
  PsAttributeValue(PsAttributeIdealProcessor, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_UMS_THREAD                                                \
  PsAttributeValue(PsAttributeUmsThread, TRUE, TRUE, FALSE)
#define PS_ATTRIBUTE_MITIGATION_OPTIONS                                        \
  PsAttributeValue(PsAttributeMitigationOptions, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_PROTECTION_LEVEL                                          \
  PsAttributeValue(PsAttributeProtectionLevel, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_SECURE_PROCESS                                            \
  PsAttributeValue(PsAttributeSecureProcess, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_JOB_LIST                                                  \
  PsAttributeValue(PsAttributeJobList, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_CHILD_PROCESS_POLICY                                      \
  PsAttributeValue(PsAttributeChildProcessPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ALL_APPLICATION_PACKAGES_POLICY                           \
  PsAttributeValue(PsAttributeAllApplicationPackagesPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_WIN32K_FILTER                                             \
  PsAttributeValue(PsAttributeWin32kFilter, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_SAFE_OPEN_PROMPT_ORIGIN_CLAIM                             \
  PsAttributeValue(PsAttributeSafeOpenPromptOriginClaim, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_BNO_ISOLATION                                             \
  PsAttributeValue(PsAttributeBnoIsolation, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_DESKTOP_APP_POLICY                                        \
  PsAttributeValue(PsAttributeDesktopAppPolicy, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_CHPE PsAttributeValue(PsAttributeChpe, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_MITIGATION_AUDIT_OPTIONS                                  \
  PsAttributeValue(PsAttributeMitigationAuditOptions, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_MACHINE_TYPE                                              \
  PsAttributeValue(PsAttributeMachineType, FALSE, TRUE, TRUE)
#define PS_ATTRIBUTE_COMPONENT_FILTER                                          \
  PsAttributeValue(PsAttributeComponentFilter, FALSE, TRUE, FALSE)
#define PS_ATTRIBUTE_ENABLE_OPTIONAL_XSTATE_FEATURES                           \
  PsAttributeValue(PsAttributeEnableOptionalXStateFeatures, TRUE, TRUE, FALSE)
typedef struct _PS_ATTRIBUTE {
  ULONG_PTR Attribute;
  SIZE_T Size;
  union {
    ULONG_PTR Value;
    PVOID ValuePtr;
  };
  PSIZE_T ReturnLength;
} PS_ATTRIBUTE, *PPS_ATTRIBUTE;
typedef struct _PS_ATTRIBUTE_LIST {
  SIZE_T TotalLength;
  PS_ATTRIBUTE Attributes[1];
} PS_ATTRIBUTE_LIST, *PPS_ATTRIBUTE_LIST;
typedef struct _PS_MEMORY_RESERVE {
  PVOID ReserveAddress;
  SIZE_T ReserveSize;
} PS_MEMORY_RESERVE, *PPS_MEMORY_RESERVE;
typedef enum _PS_STD_HANDLE_STATE {
  PsNeverDuplicate,
  PsRequestDuplicate, 
  PsAlwaysDuplicate,  
  PsMaxStdHandleStates
} PS_STD_HANDLE_STATE;
#define PS_STD_INPUT_HANDLE 0x1
#define PS_STD_OUTPUT_HANDLE 0x2
#define PS_STD_ERROR_HANDLE 0x4
typedef struct _PS_STD_HANDLE_INFO {
  union {
    ULONG Flags;
    struct {
      ULONG StdHandleState : 2;   
      ULONG PseudoHandleMask : 3; 
    };
  };
  ULONG StdHandleSubsystemType;
} PS_STD_HANDLE_INFO, *PPS_STD_HANDLE_INFO;
typedef union _PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS {
  UCHAR Trustlet : 1;
  UCHAR Ntos : 1;
  UCHAR WriteHandle : 1;
  UCHAR ReadHandle : 1;
  UCHAR Reserved : 4;
  UCHAR AccessRights;
} PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS, *PPS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS;
typedef struct _PS_TRUSTLET_ATTRIBUTE_TYPE {
  union {
    struct {
      UCHAR Version;
      UCHAR DataCount;
      UCHAR SemanticType;
      PS_TRUSTLET_ATTRIBUTE_ACCESSRIGHTS AccessRights;
    };
    ULONG AttributeType;
  };
} PS_TRUSTLET_ATTRIBUTE_TYPE, *PPS_TRUSTLET_ATTRIBUTE_TYPE;
typedef struct _PS_TRUSTLET_ATTRIBUTE_HEADER {
  PS_TRUSTLET_ATTRIBUTE_TYPE AttributeType;
  ULONG InstanceNumber : 8;
  ULONG Reserved : 24;
} PS_TRUSTLET_ATTRIBUTE_HEADER, *PPS_TRUSTLET_ATTRIBUTE_HEADER;
typedef struct _PS_TRUSTLET_ATTRIBUTE_DATA {
  PS_TRUSTLET_ATTRIBUTE_HEADER Header;
  ULONGLONG Data[1];
} PS_TRUSTLET_ATTRIBUTE_DATA, *PPS_TRUSTLET_ATTRIBUTE_DATA;
typedef struct _PS_TRUSTLET_CREATE_ATTRIBUTES {
  ULONGLONG TrustletIdentity;
  PS_TRUSTLET_ATTRIBUTE_DATA Attributes[1];
} PS_TRUSTLET_CREATE_ATTRIBUTES, *PPS_TRUSTLET_CREATE_ATTRIBUTES;
typedef struct _PS_BNO_ISOLATION_PARAMETERS {
  UNICODE_STRING IsolationPrefix;
  ULONG HandleCount;
  PVOID *Handles;
  BOOLEAN IsolationEnabled;
} PS_BNO_ISOLATION_PARAMETERS, *PPS_BNO_ISOLATION_PARAMETERS;
typedef enum _PS_MITIGATION_OPTION {
  PS_MITIGATION_OPTION_NX,
  PS_MITIGATION_OPTION_SEHOP,
  PS_MITIGATION_OPTION_FORCE_RELOCATE_IMAGES,
  PS_MITIGATION_OPTION_HEAP_TERMINATE,
  PS_MITIGATION_OPTION_BOTTOM_UP_ASLR,
  PS_MITIGATION_OPTION_HIGH_ENTROPY_ASLR,
  PS_MITIGATION_OPTION_STRICT_HANDLE_CHECKS,
  PS_MITIGATION_OPTION_WIN32K_SYSTEM_CALL_DISABLE,
  PS_MITIGATION_OPTION_EXTENSION_POINT_DISABLE,
  PS_MITIGATION_OPTION_PROHIBIT_DYNAMIC_CODE,
  PS_MITIGATION_OPTION_CONTROL_FLOW_GUARD,
  PS_MITIGATION_OPTION_BLOCK_NON_MICROSOFT_BINARIES,
  PS_MITIGATION_OPTION_FONT_DISABLE,
  PS_MITIGATION_OPTION_IMAGE_LOAD_NO_REMOTE,
  PS_MITIGATION_OPTION_IMAGE_LOAD_NO_LOW_LABEL,
  PS_MITIGATION_OPTION_IMAGE_LOAD_PREFER_SYSTEM32,
  PS_MITIGATION_OPTION_RETURN_FLOW_GUARD,
  PS_MITIGATION_OPTION_LOADER_INTEGRITY_CONTINUITY,
  PS_MITIGATION_OPTION_STRICT_CONTROL_FLOW_GUARD,
  PS_MITIGATION_OPTION_RESTRICT_SET_THREAD_CONTEXT,
  PS_MITIGATION_OPTION_ROP_STACKPIVOT, 
  PS_MITIGATION_OPTION_ROP_CALLER_CHECK,
  PS_MITIGATION_OPTION_ROP_SIMEXEC,
  PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER,
  PS_MITIGATION_OPTION_EXPORT_ADDRESS_FILTER_PLUS,
  PS_MITIGATION_OPTION_RESTRICT_CHILD_PROCESS_CREATION,
  PS_MITIGATION_OPTION_IMPORT_ADDRESS_FILTER,
  PS_MITIGATION_OPTION_MODULE_TAMPERING_PROTECTION,
  PS_MITIGATION_OPTION_RESTRICT_INDIRECT_BRANCH_PREDICTION,
  PS_MITIGATION_OPTION_SPECULATIVE_STORE_BYPASS_DISABLE, 
  PS_MITIGATION_OPTION_ALLOW_DOWNGRADE_DYNAMIC_CODE_POLICY,
  PS_MITIGATION_OPTION_CET_USER_SHADOW_STACKS,
  PS_MITIGATION_OPTION_USER_CET_SET_CONTEXT_IP_VALIDATION, 
  PS_MITIGATION_OPTION_BLOCK_NON_CET_BINARIES,
  PS_MITIGATION_OPTION_CET_DYNAMIC_APIS_OUT_OF_PROC_ONLY,
  PS_MITIGATION_OPTION_REDIRECTION_TRUST, 
} PS_MITIGATION_OPTION;
typedef enum _PS_CREATE_STATE {
  PsCreateInitialState,
  PsCreateFailOnFileOpen,
  PsCreateFailOnSectionCreate,
  PsCreateFailExeFormat,
  PsCreateFailMachineMismatch,
  PsCreateFailExeName, 
  PsCreateSuccess,
  PsCreateMaximumStates
} PS_CREATE_STATE;
typedef struct _PS_CREATE_INFO {
  SIZE_T Size;
  PS_CREATE_STATE State;
  union {
    struct {
      union {
        ULONG InitFlags;
        struct {
          UCHAR WriteOutputOnExit : 1;
          UCHAR DetectManifest : 1;
          UCHAR IFEOSkipDebugger : 1;
          UCHAR IFEODoNotPropagateKeyState : 1;
          UCHAR SpareBits1 : 4;
          UCHAR SpareBits2 : 8;
          USHORT ProhibitedImageCharacteristics : 16;
        };
      };
      ACCESS_MASK AdditionalFileAccess;
    } InitState;
    struct {
      HANDLE FileHandle;
    } FailSection;
    struct {
      USHORT DllCharacteristics;
    } ExeFormat;
    struct {
      HANDLE IFEOKey;
    } ExeName;
    struct {
      union {
        ULONG OutputFlags;
        struct {
          UCHAR ProtectedProcess : 1;
          UCHAR AddressSpaceOverride : 1;
          UCHAR DevOverrideEnabled : 1; 
          UCHAR ManifestDetected : 1;
          UCHAR ProtectedProcessLight : 1;
          UCHAR SpareBits1 : 3;
          UCHAR SpareBits2 : 8;
          USHORT SpareBits3 : 16;
        };
      };
      HANDLE FileHandle;
      HANDLE SectionHandle;
      ULONGLONG UserProcessParametersNative;
      ULONG UserProcessParametersWow64;
      ULONG CurrentParameterFlags;
      ULONGLONG PebAddressNative;
      ULONG PebAddressWow64;
      ULONGLONG ManifestAddress;
      ULONG ManifestSize;
    } SuccessState;
  };
} PS_CREATE_INFO, *PPS_CREATE_INFO;
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateUserProcess(
    _Out_ PHANDLE ProcessHandle, _Out_ PHANDLE ThreadHandle,
    _In_ ACCESS_MASK ProcessDesiredAccess, _In_ ACCESS_MASK ThreadDesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ProcessObjectAttributes,
    _In_opt_ POBJECT_ATTRIBUTES ThreadObjectAttributes,
    _In_ ULONG ProcessFlags,          
    _In_ ULONG ThreadFlags,           
    _In_opt_ PVOID ProcessParameters, 
    _Inout_ PPS_CREATE_INFO CreateInfo,
    _In_opt_ PPS_ATTRIBUTE_LIST AttributeList);
#endif
#define THREAD_CREATE_FLAGS_CREATE_SUSPENDED                                   \
  0x00000001 
#define THREAD_CREATE_FLAGS_SKIP_THREAD_ATTACH                                 \
  0x00000002 
#define THREAD_CREATE_FLAGS_HIDE_FROM_DEBUGGER                                 \
  0x00000004                                            
#define THREAD_CREATE_FLAGS_LOADER_WORKER 0x00000010    
#define THREAD_CREATE_FLAGS_SKIP_LOADER_INIT 0x00000020 
#define THREAD_CREATE_FLAGS_BYPASS_PROCESS_FREEZE                              \
  0x00000040                                          
#define THREAD_CREATE_FLAGS_INITIAL_THREAD 0x00000080 
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateThreadEx(_Out_ PHANDLE ThreadHandle, _In_ ACCESS_MASK DesiredAccess,
                 _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
                 _In_ HANDLE ProcessHandle,
                 _In_ PVOID StartRoutine, 
                 _In_opt_ PVOID Argument,
                 _In_ ULONG CreateFlags, 
                 _In_ SIZE_T ZeroBits, _In_ SIZE_T StackSize,
                 _In_ SIZE_T MaximumStackSize,
                 _In_opt_ PPS_ATTRIBUTE_LIST AttributeList);
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define JobObjectBasicAccountingInformation                                    \
  1 
#define JobObjectBasicLimitInformation 2 
#define JobObjectBasicProcessIdList 3    
#define JobObjectBasicUIRestrictions 4   
#define JobObjectSecurityLimitInformation                                      \
  5 
#define JobObjectEndOfJobTimeInformation                                       \
  6 
#define JobObjectAssociateCompletionPortInformation                            \
  7 
#define JobObjectBasicAndIoAccountingInformation                               \
  8 
#define JobObjectExtendedLimitInformation                                      \
  9                                   
#define JobObjectJobSetInformation 10 
#define JobObjectGroupInformation 11  
#define JobObjectNotificationLimitInformation                                  \
  12 
#define JobObjectLimitViolationInformation                                     \
  13                                   
#define JobObjectGroupInformationEx 14 
#define JobObjectCpuRateControlInformation                                     \
  15 
#define JobObjectCompletionFilter 16
#define JobObjectCompletionCounter 17
#define JobObjectFreezeInformation 18 
#define JobObjectExtendedAccountingInformation                                 \
  19                                
#define JobObjectWakeInformation 20 
#define JobObjectBackgroundInformation 21
#define JobObjectSchedulingRankBiasInformation 22
#define JobObjectTimerVirtualizationInformation 23
#define JobObjectCycleTimeNotification 24
#define JobObjectClearEvent 25
#define JobObjectInterferenceInformation                                       \
  26 
#define JobObjectClearPeakJobMemoryUsed 27
#define JobObjectMemoryUsageInformation                                        \
  28 
#define JobObjectSharedCommit 29
#define JobObjectContainerId 30
#define JobObjectIoRateControlInformation 31
#define JobObjectNetRateControlInformation                                     \
  32 
#define JobObjectNotificationLimitInformation2                                 \
  33 
#define JobObjectLimitViolationInformation2                                    \
  34 
#define JobObjectCreateSilo 35
#define JobObjectSiloBasicInformation 36       
#define JobObjectSiloRootDirectory 37          
#define JobObjectServerSiloBasicInformation 38 
#define JobObjectServerSiloUserSharedData 39   
#define JobObjectServerSiloInitialize 40
#define JobObjectServerSiloRunningState 41
#define JobObjectIoAttribution 42
#define JobObjectMemoryPartitionInformation 43
#define JobObjectContainerTelemetryId 44
#define JobObjectSiloSystemRoot 45
#define JobObjectEnergyTrackingState 46 
#define JobObjectThreadImpersonationInformation 47
#define JobObjectIoPriorityLimit 48
#define JobObjectPagePriorityLimit 49
#define MaxJobObjectInfoClass 50
typedef struct _JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION {
  JOBOBJECT_BASIC_ACCOUNTING_INFORMATION BasicInfo;
  IO_COUNTERS IoInfo;
  PROCESS_DISK_COUNTERS DiskIoInfo;
  ULONG64 ContextSwitches;
  LARGE_INTEGER TotalCycleTime;
  ULONG64 ReadyTime;
  PROCESS_ENERGY_VALUES EnergyValues;
} JOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION,
    *PJOBOBJECT_EXTENDED_ACCOUNTING_INFORMATION;
typedef struct _JOBOBJECT_WAKE_INFORMATION {
  HANDLE NotificationChannel;
  ULONG64 WakeCounters[7];
} JOBOBJECT_WAKE_INFORMATION, *PJOBOBJECT_WAKE_INFORMATION;
typedef struct _JOBOBJECT_WAKE_INFORMATION_V1 {
  HANDLE NotificationChannel;
  ULONG64 WakeCounters[4];
} JOBOBJECT_WAKE_INFORMATION_V1, *PJOBOBJECT_WAKE_INFORMATION_V1;
typedef struct _JOBOBJECT_INTERFERENCE_INFORMATION {
  ULONG64 Count;
} JOBOBJECT_INTERFERENCE_INFORMATION, *PJOBOBJECT_INTERFERENCE_INFORMATION;
typedef struct _JOBOBJECT_WAKE_FILTER {
  ULONG HighEdgeFilter;
  ULONG LowEdgeFilter;
} JOBOBJECT_WAKE_FILTER, *PJOBOBJECT_WAKE_FILTER;
typedef struct _JOBOBJECT_FREEZE_INFORMATION {
  union {
    ULONG Flags;
    struct {
      ULONG FreezeOperation : 1;
      ULONG FilterOperation : 1;
      ULONG SwapOperation : 1;
      ULONG Reserved : 29;
    };
  };
  BOOLEAN Freeze;
  BOOLEAN Swap;
  UCHAR Reserved0[2];
  JOBOBJECT_WAKE_FILTER WakeFilter;
} JOBOBJECT_FREEZE_INFORMATION, *PJOBOBJECT_FREEZE_INFORMATION;
typedef struct _JOBOBJECT_MEMORY_USAGE_INFORMATION {
  ULONG64 JobMemory;
  ULONG64 PeakJobMemoryUsed;
} JOBOBJECT_MEMORY_USAGE_INFORMATION, *PJOBOBJECT_MEMORY_USAGE_INFORMATION;
typedef struct _JOBOBJECT_MEMORY_USAGE_INFORMATION_V2 {
  JOBOBJECT_MEMORY_USAGE_INFORMATION BasicInfo;
  ULONG64 JobSharedMemory;
  ULONG64 Reserved[2];
} JOBOBJECT_MEMORY_USAGE_INFORMATION_V2,
    *PJOBOBJECT_MEMORY_USAGE_INFORMATION_V2;
typedef struct _SILO_USER_SHARED_DATA {
  ULONG64 ServiceSessionId;
  ULONG ActiveConsoleId;
  LONGLONG ConsoleSessionForegroundProcessId;
  NT_PRODUCT_TYPE NtProductType;
  ULONG SuiteMask;
  ULONG SharedUserSessionId;
  BOOLEAN IsMultiSessionSku;
  WCHAR NtSystemRoot[260];
  USHORT UserModeGlobalLogger[16];
} SILO_USER_SHARED_DATA, *PSILO_USER_SHARED_DATA;
typedef struct _SILOOBJECT_ROOT_DIRECTORY {
  ULONG ControlFlags;
  UNICODE_STRING Path;
} SILOOBJECT_ROOT_DIRECTORY, *PSILOOBJECT_ROOT_DIRECTORY;
typedef struct _JOBOBJECT_ENERGY_TRACKING_STATE {
  ULONG64 Value;
  ULONG UpdateMask;
  ULONG DesiredState;
} JOBOBJECT_ENERGY_TRACKING_STATE, *PJOBOBJECT_ENERGY_TRACKING_STATE;
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobObject(_Out_ PHANDLE JobHandle, _In_ ACCESS_MASK DesiredAccess,
                  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenJobObject(_Out_ PHANDLE JobHandle, _In_ ACCESS_MASK DesiredAccess,
                _In_ POBJECT_ATTRIBUTES ObjectAttributes);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAssignProcessToJobObject(_In_ HANDLE JobHandle, _In_ HANDLE ProcessHandle);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTerminateJobObject(_In_ HANDLE JobHandle, _In_ NTSTATUS ExitStatus);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtIsProcessInJob(_In_ HANDLE ProcessHandle, _In_opt_ HANDLE JobHandle);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryInformationJobObject(_In_opt_ HANDLE JobHandle,
                            _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
                            _Out_writes_bytes_(JobObjectInformationLength)
                                PVOID JobObjectInformation,
                            _In_ ULONG JobObjectInformationLength,
                            _Out_opt_ PULONG ReturnLength);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationJobObject(_In_ HANDLE JobHandle,
                          _In_ JOBOBJECTINFOCLASS JobObjectInformationClass,
                          _In_reads_bytes_(JobObjectInformationLength)
                              PVOID JobObjectInformation,
                          _In_ ULONG JobObjectInformationLength);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateJobSet(_In_ ULONG NumJob, _In_reads_(NumJob) PJOB_SET_ARRAY UserJobSet,
               _In_ ULONG Flags);
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRevertContainerImpersonation(VOID);
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _MEMORY_RESERVE_TYPE {
  MemoryReserveUserApc,
  MemoryReserveIoCompletion,
  MemoryReserveTypeMax
} MEMORY_RESERVE_TYPE;
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAllocateReserveObject(_Out_ PHANDLE MemoryReserveHandle,
                        _In_ POBJECT_ATTRIBUTES ObjectAttributes,
                        _In_ MEMORY_RESERVE_TYPE Type);
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSCALLAPI
NTSTATUS
NTAPI
PssNtCaptureSnapshot(_Out_ PHANDLE SnapshotHandle, _In_ HANDLE ProcessHandle,
                     _In_ ULONG CaptureFlags, _In_ ULONG ThreadContextFlags);
#endif
#endif
#endif
