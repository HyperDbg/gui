/*
 * This file is part of the Process Hacker project - https://processhacker.sourceforge.io/
 *
 * You can redistribute this file and/or modify it under the terms of the 
 * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
 */

#ifndef _NTPOAPI_H
#define _NTPOAPI_H

#if (PHNT_MODE != PHNT_MODE_KERNEL)


#define SystemPowerPolicyAc 0 
#define SystemPowerPolicyDc 1 
#define VerifySystemPolicyAc 2 
#define VerifySystemPolicyDc 3 
#define SystemPowerCapabilities 4 
#define SystemBatteryState 5 
#define SystemPowerStateHandler 6 
#define ProcessorStateHandler 7 
#define SystemPowerPolicyCurrent 8 
#define AdministratorPowerPolicy 9 
#define SystemReserveHiberFile 10 
#define ProcessorInformation 11 
#define SystemPowerInformation 12 
#define ProcessorStateHandler2 13 
#define LastWakeTime 14 
#define LastSleepTime 15 
#define SystemExecutionState 16 
#define SystemPowerStateNotifyHandler 17 
#define ProcessorPowerPolicyAc 18 
#define ProcessorPowerPolicyDc 19 
#define VerifyProcessorPowerPolicyAc 20 
#define VerifyProcessorPowerPolicyDc 21 
#define ProcessorPowerPolicyCurrent 22 
#define SystemPowerStateLogging 23 
#define SystemPowerLoggingEntry 24 
#define SetPowerSettingValue 25 
#define NotifyUserPowerSetting 26 
#define PowerInformationLevelUnused0 27 
#define SystemMonitorHiberBootPowerOff 28 
#define SystemVideoState 29 
#define TraceApplicationPowerMessage 30 
#define TraceApplicationPowerMessageEnd 31 
#define ProcessorPerfStates 32 
#define ProcessorIdleStates 33 
#define ProcessorCap 34 
#define SystemWakeSource 35
#define SystemHiberFileInformation 36 
#define TraceServicePowerMessage 37
#define ProcessorLoad 38
#define PowerShutdownNotification 39 
#define MonitorCapabilities 40 
#define SessionPowerInit 41 
#define SessionDisplayState 42 
#define PowerRequestCreate 43 
#define PowerRequestAction 44 
#define GetPowerRequestList 45 
#define ProcessorInformationEx 46 
#define NotifyUserModeLegacyPowerEvent 47 
#define GroupPark 48 
#define ProcessorIdleDomains 49 
#define WakeTimerList 50 
#define SystemHiberFileSize 51 
#define ProcessorIdleStatesHv 52 
#define ProcessorPerfStatesHv 53 
#define ProcessorPerfCapHv 54 
#define ProcessorSetIdle 55 
#define LogicalProcessorIdling 56 
#define UserPresence 57 
#define PowerSettingNotificationName 58
#define GetPowerSettingValue 59 
#define IdleResiliency 60 
#define SessionRITState 61 
#define SessionConnectNotification 62 
#define SessionPowerCleanup 63
#define SessionLockState 64 
#define SystemHiberbootState 65 
#define PlatformInformation 66 
#define PdcInvocation 67 
#define MonitorInvocation 68 
#define FirmwareTableInformationRegistered 69 
#define SetShutdownSelectedTime 70 
#define SuspendResumeInvocation 71 
#define PlmPowerRequestCreate 72 
#define ScreenOff 73 
#define CsDeviceNotification 74 
#define PlatformRole 75 
#define LastResumePerformance 76 
#define DisplayBurst 77 
#define ExitLatencySamplingPercentage 78
#define RegisterSpmPowerSettings 79 
#define PlatformIdleStates 80 
#define ProcessorIdleVeto 81 
#define PlatformIdleVeto 82 
#define SystemBatteryStatePrecise 83 
#define ThermalEvent 84  
#define PowerRequestActionInternal 85 
#define BatteryDeviceState 86
#define PowerInformationInternal 87 
#define ThermalStandby 88 
#define SystemHiberFileType 89 
#define PhysicalPowerButtonPress 90 
#define QueryPotentialDripsConstraint 91 
#define EnergyTrackerCreate 92
#define EnergyTrackerQuery 93
#define UpdateBlackBoxRecorder 94
#define SessionAllowExternalDmaDevices 95
#define SendSuspendResumeNotification 96 
#define PowerInformationLevelMaximum 97
#endif

typedef struct _PROCESSOR_POWER_INFORMATION
{
    ULONG Number;
    ULONG MaxMhz;
    ULONG CurrentMhz;
    ULONG MhzLimit;
    ULONG MaxIdleState;
    ULONG CurrentIdleState;
} PROCESSOR_POWER_INFORMATION, *PPROCESSOR_POWER_INFORMATION;

typedef struct _SYSTEM_POWER_INFORMATION
{
    ULONG MaxIdlenessAllowed;
    ULONG Idleness;
    ULONG TimeRemaining;
    UCHAR CoolingMode;
} SYSTEM_POWER_INFORMATION, *PSYSTEM_POWER_INFORMATION;

typedef struct _SYSTEM_HIBERFILE_INFORMATION
{
    ULONG NumberOfMcbPairs;
    LARGE_INTEGER Mcb[1];
} SYSTEM_HIBERFILE_INFORMATION, *PSYSTEM_HIBERFILE_INFORMATION;

#define POWER_REQUEST_CONTEXT_NOT_SPECIFIED DIAGNOSTIC_REASON_NOT_SPECIFIED


typedef struct _COUNTED_REASON_CONTEXT
{
    ULONG Version;
    ULONG Flags;
    union
    {
        struct
        {
            UNICODE_STRING ResourceFileName;
            USHORT ResourceReasonId;
            ULONG StringCount;
            _Field_size_(StringCount) PUNICODE_STRING ReasonStrings;
        };
        UNICODE_STRING SimpleString;
    };
} COUNTED_REASON_CONTEXT, *PCOUNTED_REASON_CONTEXT;

typedef enum _POWER_REQUEST_TYPE_INTERNAL 
{
    PowerRequestDisplayRequiredInternal,
    PowerRequestSystemRequiredInternal,
    PowerRequestAwayModeRequiredInternal,
    PowerRequestExecutionRequiredInternal, 
    PowerRequestPerfBoostRequiredInternal, 
    PowerRequestActiveLockScreenInternal, 
    
    PowerRequestInternalInvalid,
    PowerRequestInternalUnknown,
    PowerRequestFullScreenVideoRequired  
} POWER_REQUEST_TYPE_INTERNAL;

typedef struct _POWER_REQUEST_ACTION
{
    HANDLE PowerRequestHandle;
    POWER_REQUEST_TYPE_INTERNAL RequestType;
    BOOLEAN SetAction;
    HANDLE ProcessHandle; 
} POWER_REQUEST_ACTION, *PPOWER_REQUEST_ACTION;

typedef union _POWER_STATE
{
    SYSTEM_POWER_STATE SystemState;
    DEVICE_POWER_STATE DeviceState;
} POWER_STATE, *PPOWER_STATE;

typedef enum _POWER_STATE_TYPE
{
    SystemPowerState = 0,
    DevicePowerState
} POWER_STATE_TYPE, *PPOWER_STATE_TYPE;


typedef struct _SYSTEM_POWER_STATE_CONTEXT
{
    union
    {
        struct
        {
            ULONG Reserved1 : 8;
            ULONG TargetSystemState : 4;
            ULONG EffectiveSystemState : 4;
            ULONG CurrentSystemState : 4;
            ULONG IgnoreHibernationPath : 1;
            ULONG PseudoTransition : 1;
            ULONG Reserved2 : 10;
        };
        ULONG ContextAsUlong;
    };
} SYSTEM_POWER_STATE_CONTEXT, *PSYSTEM_POWER_STATE_CONTEXT;

typedef enum _REQUESTER_TYPE
{
    KernelRequester = 0,
    UserProcessRequester = 1,
    UserSharedServiceRequester = 2
} REQUESTER_TYPE;

typedef struct _COUNTED_REASON_CONTEXT_RELATIVE
{
    ULONG Flags;
    union
    {
        struct
        {
            ULONG_PTR ResourceFileNameOffset;
            USHORT ResourceReasonId;
            ULONG StringCount;
            ULONG_PTR SubstitutionStringsOffset;
        };
        ULONG_PTR SimpleStringOffset;
    };
} COUNTED_REASON_CONTEXT_RELATIVE, *PCOUNTED_REASON_CONTEXT_RELATIVE;

typedef struct _DIAGNOSTIC_BUFFER
{
    SIZE_T Size;
    REQUESTER_TYPE CallerType;
    union
    {
        struct
        {
            ULONG_PTR ProcessImageNameOffset; 
            ULONG ProcessId;
            ULONG ServiceTag;
        };
        struct
        {
            ULONG_PTR DeviceDescriptionOffset; 
            ULONG_PTR DevicePathOffset; 
        };
    };    
    ULONG_PTR ReasonOffset; 
} DIAGNOSTIC_BUFFER, *PDIAGNOSTIC_BUFFER;


#define POWER_REQUEST_SUPPORTED_TYPES_V1 3 
#define POWER_REQUEST_SUPPORTED_TYPES_V2 9 
#define POWER_REQUEST_SUPPORTED_TYPES_V3 5 
#define POWER_REQUEST_SUPPORTED_TYPES_V4 6 

typedef struct _POWER_REQUEST
{
    union
    {
        struct
        {
            ULONG SupportedRequestMask;
            ULONG PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V1];
            DIAGNOSTIC_BUFFER DiagnosticBuffer;
        } V1;
#if (PHNT_VERSION >= PHNT_WIN8)
        struct
        {
            ULONG SupportedRequestMask;
            ULONG PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V2];
            DIAGNOSTIC_BUFFER DiagnosticBuffer;
        } V2;
#endif
#if (PHNT_VERSION >= PHNT_WINBLUE)
        struct
        {
            ULONG SupportedRequestMask;
            ULONG PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V3];
            DIAGNOSTIC_BUFFER DiagnosticBuffer;
        } V3;
#endif
#if (PHNT_VERSION >= PHNT_REDSTONE)
        struct
        {
            ULONG SupportedRequestMask;
            ULONG PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V4];
            DIAGNOSTIC_BUFFER DiagnosticBuffer;
        } V4;
#endif
    };
} POWER_REQUEST, *PPOWER_REQUEST;

typedef struct _POWER_REQUEST_LIST
{
    ULONG_PTR Count;
    ULONG_PTR PowerRequestOffsets[ANYSIZE_ARRAY]; 
} POWER_REQUEST_LIST, *PPOWER_REQUEST_LIST;

typedef enum _POWER_STATE_HANDLER_TYPE
{
    PowerStateSleeping1 = 0,
    PowerStateSleeping2 = 1,
    PowerStateSleeping3 = 2,
    PowerStateSleeping4 = 3,
    PowerStateShutdownOff = 4,
    PowerStateShutdownReset = 5,
    PowerStateSleeping4Firmware = 6,
    PowerStateMaximum = 7
} POWER_STATE_HANDLER_TYPE, *PPOWER_STATE_HANDLER_TYPE;

typedef NTSTATUS (NTAPI *PENTER_STATE_SYSTEM_HANDLER)(
    _In_ PVOID SystemContext
    );

typedef NTSTATUS (NTAPI *PENTER_STATE_HANDLER)(
    _In_ PVOID Context,
    _In_opt_ PENTER_STATE_SYSTEM_HANDLER SystemHandler,
    _In_ PVOID SystemContext,
    _In_ LONG NumberProcessors,
    _In_ LONG volatile *Number
    );

typedef struct _POWER_STATE_HANDLER
{
    POWER_STATE_HANDLER_TYPE Type;
    BOOLEAN RtcWake;
    UCHAR Spare[3];
    PENTER_STATE_HANDLER Handler;
    PVOID Context;
} POWER_STATE_HANDLER, *PPOWER_STATE_HANDLER;

typedef NTSTATUS (NTAPI *PENTER_STATE_NOTIFY_HANDLER)(
    _In_ POWER_STATE_HANDLER_TYPE State,
    _In_ PVOID Context,
    _In_ BOOLEAN Entering
    );

typedef struct _POWER_STATE_NOTIFY_HANDLER
{
    PENTER_STATE_NOTIFY_HANDLER Handler;
    PVOID Context;
} POWER_STATE_NOTIFY_HANDLER, *PPOWER_STATE_NOTIFY_HANDLER;

typedef struct _POWER_REQUEST_ACTION_INTERNAL
{
    PVOID PowerRequestPointer;
    POWER_REQUEST_TYPE_INTERNAL RequestType;
    BOOLEAN SetAction;
} POWER_REQUEST_ACTION_INTERNAL, *PPOWER_REQUEST_ACTION_INTERNAL;

typedef enum _POWER_INFORMATION_LEVEL_INTERNAL
{
    PowerInternalAcpiInterfaceRegister,
    PowerInternalS0LowPowerIdleInfo, 
    PowerInternalReapplyBrightnessSettings,
    PowerInternalUserAbsencePrediction, 
    PowerInternalUserAbsencePredictionCapability, 
    PowerInternalPoProcessorLatencyHint, 
    PowerInternalStandbyNetworkRequest, 
    PowerInternalDirtyTransitionInformation,
    PowerInternalSetBackgroundTaskState, 
    PowerInternalTtmOpenTerminal,
    PowerInternalTtmCreateTerminal, 
    PowerInternalTtmEvacuateDevices,
    PowerInternalTtmCreateTerminalEventQueue,
    PowerInternalTtmGetTerminalEvent,
    PowerInternalTtmSetDefaultDeviceAssignment,
    PowerInternalTtmAssignDevice,
    PowerInternalTtmSetDisplayState,
    PowerInternalTtmSetDisplayTimeouts,
    PowerInternalBootSessionStandbyActivationInformation,
    PowerInternalSessionPowerState,
    PowerInternalSessionTerminalInput, 
    PowerInternalSetWatchdog,
    PowerInternalPhysicalPowerButtonPressInfoAtBoot,
    PowerInternalExternalMonitorConnected,
    PowerInternalHighPrecisionBrightnessSettings,
    PowerInternalWinrtScreenToggle,
    PowerInternalPpmQosDisable,
    PowerInternalTransitionCheckpoint,
    PowerInternalInputControllerState,
    PowerInternalFirmwareResetReason,
    PowerInternalPpmSchedulerQosSupport, 
    PowerInternalBootStatGet,
    PowerInternalBootStatSet,
    PowerInternalCallHasNotReturnedWatchdog,
    PowerInternalBootStatCheckIntegrity,
    PowerInternalBootStatRestoreDefaults, 
    PowerInternalHostEsStateUpdate,
    PowerInternalGetPowerActionState,
    PowerInternalBootStatUnlock,
    PowerInternalWakeOnVoiceState,
    PowerInternalDeepSleepBlock, 
    PowerInternalIsPoFxDevice,
    PowerInternalPowerTransitionExtensionAtBoot,
    PowerInternalProcessorBrandedFrequency, 
    PowerInternalTimeBrokerExpirationReason,
    PowerInternalNotifyUserShutdownStatus,
    PowerInternalPowerRequestTerminalCoreWindow,
    PowerInternalProcessorIdleVeto,
    PowerInternalPlatformIdleVeto,
    PowerInternalIsLongPowerButtonBugcheckEnabled,
    PowerInternalAutoChkCausedReboot, 
    PowerInternalSetWakeAlarmOverride,

    PowerInternalDirectedFxAddTestDevice = 53,
    PowerInternalDirectedFxRemoveTestDevice,

    PowerInternalDirectedFxSetMode = 56,
    PowerInternalRegisterPowerPlane,
    PowerInternalSetDirectedDripsFlags,
    PowerInternalClearDirectedDripsFlags,
    PowerInternalRetrieveHiberFileResumeContext, 
    PowerInternalReadHiberFilePage,
    PowerInternalLastBootSucceeded, 
    PowerInternalQuerySleepStudyHelperRoutineBlock,
    PowerInternalDirectedDripsQueryCapabilities,
    PowerInternalClearConstraints,
    PowerInternalSoftParkVelocityEnabled,
    PowerInternalQueryIntelPepCapabilities,
    PowerInternalGetSystemIdleLoopEnablement, 
    PowerInternalGetVmPerfControlSupport,
    PowerInternalGetVmPerfControlConfig, 
    PowerInternalSleepDetailedDiagUpdate,
    PowerInternalProcessorClassFrequencyBandsStats,
    PowerInternalHostGlobalUserPresenceStateUpdate,
    PowerInternalCpuNodeIdleIntervalStats,
    PowerInternalClassIdleIntervalStats,
    PowerInternalCpuNodeConcurrencyStats,
    PowerInternalClassConcurrencyStats,
    PowerInternalQueryProcMeasurementCapabilities,
    PowerInternalQueryProcMeasurementValues,
    PowerInternalPrepareForSystemInitiatedReboot, 
    PowerInternalGetAdaptiveSessionState,
    PowerInternalSetConsoleLockedState,
    PowerInternalOverrideSystemInitiatedRebootState,
    PowerInternalFanImpactStats,
    PowerInternalFanRpmBuckets,
    PowerInternalPowerBootAppDiagInfo,
    PowerInternalUnregisterShutdownNotification, 
    PowerInternalManageTransitionStateRecord,
    PowerInformationInternalMaximum
} POWER_INFORMATION_LEVEL_INTERNAL;

typedef enum _POWER_S0_DISCONNECTED_REASON
{
    PoS0DisconnectedReasonNone,
    PoS0DisconnectedReasonNonCompliantNic,
    PoS0DisconnectedReasonSettingPolicy,
    PoS0DisconnectedReasonEnforceDsPolicy,
    PoS0DisconnectedReasonCsChecksFailed,
    PoS0DisconnectedReasonSmartStandby,
    PoS0DisconnectedReasonMaximum
} POWER_S0_DISCONNECTED_REASON;

typedef struct _POWER_S0_LOW_POWER_IDLE_INFO
{
    POWER_S0_DISCONNECTED_REASON DisconnectedReason;
    union
    {
        BOOLEAN Storage : 1;
        BOOLEAN WiFi : 1;
        BOOLEAN Mbn : 1;
        BOOLEAN Ethernet : 1;
        BOOLEAN Reserved : 4;
        UCHAR AsUCHAR;
    } CsDeviceCompliance;
    union
    {
        BOOLEAN DisconnectInStandby : 1;
        BOOLEAN EnforceDs : 1;
        BOOLEAN Reserved : 6;
        UCHAR AsUCHAR;
    } Policy;
} POWER_S0_LOW_POWER_IDLE_INFO, *PPOWER_S0_LOW_POWER_IDLE_INFO;

typedef struct _POWER_INFORMATION_INTERNAL_HEADER
{
    POWER_INFORMATION_LEVEL_INTERNAL InternalType;
    ULONG Version;
} POWER_INFORMATION_INTERNAL_HEADER, *PPOWER_INFORMATION_INTERNAL_HEADER;

typedef struct _POWER_USER_ABSENCE_PREDICTION
{
    POWER_INFORMATION_INTERNAL_HEADER Header;
    LARGE_INTEGER ReturnTime;
} POWER_USER_ABSENCE_PREDICTION, *PPOWER_USER_ABSENCE_PREDICTION;

typedef struct _POWER_USER_ABSENCE_PREDICTION_CAPABILITY
{
    BOOLEAN AbsencePredictionCapability;
} POWER_USER_ABSENCE_PREDICTION_CAPABILITY, *PPOWER_USER_ABSENCE_PREDICTION_CAPABILITY;

typedef struct _POWER_PROCESSOR_LATENCY_HINT
{
    POWER_INFORMATION_INTERNAL_HEADER PowerInformationInternalHeader;
    ULONG Type;
} POWER_PROCESSOR_LATENCY_HINT, *PPO_PROCESSOR_LATENCY_HINT;

typedef struct _POWER_STANDBY_NETWORK_REQUEST
{
    POWER_INFORMATION_INTERNAL_HEADER PowerInformationInternalHeader;
    BOOLEAN Active;
} POWER_STANDBY_NETWORK_REQUEST, *PPOWER_STANDBY_NETWORK_REQUEST;

typedef struct _POWER_SET_BACKGROUND_TASK_STATE
{
    POWER_INFORMATION_INTERNAL_HEADER PowerInformationInternalHeader;
    BOOLEAN Engaged;
} POWER_SET_BACKGROUND_TASK_STATE, *PPOWER_SET_BACKGROUND_TASK_STATE;

typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT
{
    POWER_INFORMATION_LEVEL_INTERNAL InternalType;
    PROCESSOR_NUMBER ProcessorNumber; 
} POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT, *PPOWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT;

typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT
{
    ULONG Version;
    ULONG NominalFrequency; 
} POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT, *PPOWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT;

NTSYSCALLAPI
NTSTATUS
NTAPI
NtPowerInformation(
    _In_ POWER_INFORMATION_LEVEL InformationLevel,
    _In_reads_bytes_opt_(InputBufferLength) PVOID InputBuffer,
    _In_ ULONG InputBufferLength,
    _Out_writes_bytes_opt_(OutputBufferLength) PVOID OutputBuffer,
    _In_ ULONG OutputBufferLength
    );

NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetThreadExecutionState(
    _In_ EXECUTION_STATE NewFlags, 
    _Out_ EXECUTION_STATE *PreviousFlags
    );

#if (PHNT_VERSION < PHNT_WIN7)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRequestWakeupLatency(
    _In_ LATENCY_TIME latency
    );
#endif

NTSYSCALLAPI
NTSTATUS
NTAPI
NtInitiatePowerAction(
    _In_ POWER_ACTION SystemAction,
    _In_ SYSTEM_POWER_STATE LightestSystemState,
    _In_ ULONG Flags, 
    _In_ BOOLEAN Asynchronous
    );

NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSystemPowerState(
    _In_ POWER_ACTION SystemAction,
    _In_ SYSTEM_POWER_STATE LightestSystemState,
    _In_ ULONG Flags 
    );

NTSYSCALLAPI
NTSTATUS
NTAPI
NtGetDevicePowerState(
    _In_ HANDLE Device,
    _Out_ PDEVICE_POWER_STATE State
    );

NTSYSCALLAPI
BOOLEAN
NTAPI
NtIsSystemResumeAutomatic(
    VOID
    );

#endif
