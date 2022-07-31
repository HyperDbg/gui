package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpoapi.h.back

const(
_NTPOAPI_H =  //col:13
SystemPowerPolicyAc = 0 // SYSTEM_POWER_POLICY // GET: InputBuffer NULL. SET: InputBuffer not NULL. //col:18
SystemPowerPolicyDc = 1 // SYSTEM_POWER_POLICY //col:19
VerifySystemPolicyAc = 2 // SYSTEM_POWER_POLICY //col:20
VerifySystemPolicyDc = 3 // SYSTEM_POWER_POLICY //col:21
SystemPowerCapabilities = 4 // SYSTEM_POWER_CAPABILITIES //col:22
SystemBatteryState = 5 // SYSTEM_BATTERY_STATE //col:23
SystemPowerStateHandler = 6 // POWER_STATE_HANDLER // (kernel-mode only) //col:24
ProcessorStateHandler = 7 // PROCESSOR_STATE_HANDLER // (kernel-mode only) //col:25
SystemPowerPolicyCurrent = 8 // SYSTEM_POWER_POLICY //col:26
AdministratorPowerPolicy = 9 // ADMINISTRATOR_POWER_POLICY //col:27
SystemReserveHiberFile = 10 // BOOLEAN // (requires SeCreatePagefilePrivilege) // TRUE: hibernation file created. FALSE: hibernation file deleted. //col:28
ProcessorInformation = 11 // PROCESSOR_POWER_INFORMATION //col:29
SystemPowerInformation = 12 // SYSTEM_POWER_INFORMATION //col:30
ProcessorStateHandler2 = 13 // PROCESSOR_STATE_HANDLER2 // not implemented //col:31
LastWakeTime = 14 // ULONGLONG // InterruptTime //col:32
LastSleepTime = 15 // ULONGLONG // InterruptTime //col:33
SystemExecutionState = 16 // EXECUTION_STATE // NtSetThreadExecutionState //col:34
SystemPowerStateNotifyHandler = 17 // POWER_STATE_NOTIFY_HANDLER // (kernel-mode only) //col:35
ProcessorPowerPolicyAc = 18 // PROCESSOR_POWER_POLICY // not implemented //col:36
ProcessorPowerPolicyDc = 19 // PROCESSOR_POWER_POLICY // not implemented //col:37
VerifyProcessorPowerPolicyAc = 20 // PROCESSOR_POWER_POLICY // not implemented //col:38
VerifyProcessorPowerPolicyDc = 21 // PROCESSOR_POWER_POLICY // not implemented //col:39
ProcessorPowerPolicyCurrent = 22 // PROCESSOR_POWER_POLICY // not implemented //col:40
SystemPowerStateLogging = 23 // SYSTEM_POWER_STATE_DISABLE_REASON[] //col:41
SystemPowerLoggingEntry = 24 // SYSTEM_POWER_LOGGING_ENTRY[] // (kernel-mode only) //col:42
SetPowerSettingValue = 25 // (kernel-mode only) //col:43
NotifyUserPowerSetting = 26 // not implemented //col:44
PowerInformationLevelUnused0 = 27 // not implemented //col:45
SystemMonitorHiberBootPowerOff = 28 // NULL (PowerMonitorOff) //col:46
SystemVideoState = 29 // MONITOR_DISPLAY_STATE //col:47
TraceApplicationPowerMessage = 30 // (kernel-mode only) //col:48
TraceApplicationPowerMessageEnd = 31 // (kernel-mode only) //col:49
ProcessorPerfStates = 32 // (kernel-mode only) //col:50
ProcessorIdleStates = 33 // (kernel-mode only) //col:51
ProcessorCap = 34 // (kernel-mode only) //col:52
SystemWakeSource = 35 //col:53
SystemHiberFileInformation = 36 // q: SYSTEM_HIBERFILE_INFORMATION //col:54
TraceServicePowerMessage = 37 //col:55
ProcessorLoad = 38 //col:56
PowerShutdownNotification = 39 // (kernel-mode only) //col:57
MonitorCapabilities = 40 // (kernel-mode only) //col:58
SessionPowerInit = 41 // (kernel-mode only) //col:59
SessionDisplayState = 42 // (kernel-mode only) //col:60
PowerRequestCreate = 43 // in: COUNTED_REASON_CONTEXT, out: HANDLE //col:61
PowerRequestAction = 44 // in: POWER_REQUEST_ACTION //col:62
GetPowerRequestList = 45 // out: POWER_REQUEST_LIST //col:63
ProcessorInformationEx = 46 // in: USHORT ProcessorGroup, out: PROCESSOR_POWER_INFORMATION //col:64
NotifyUserModeLegacyPowerEvent = 47 // (kernel-mode only) //col:65
GroupPark = 48 // (debug-mode boot only) //col:66
ProcessorIdleDomains = 49 // (kernel-mode only) //col:67
WakeTimerList = 50 // powercfg.exe /waketimers //col:68
SystemHiberFileSize = 51 // ULONG //col:69
ProcessorIdleStatesHv = 52 // (kernel-mode only) //col:70
ProcessorPerfStatesHv = 53 // (kernel-mode only) //col:71
ProcessorPerfCapHv = 54 // (kernel-mode only) //col:72
ProcessorSetIdle = 55 // (debug-mode boot only) //col:73
LogicalProcessorIdling = 56 // (kernel-mode only) //col:74
UserPresence = 57 // POWER_USER_PRESENCE // not implemented //col:75
PowerSettingNotificationName = 58 //col:76
GetPowerSettingValue = 59 // GUID //col:77
IdleResiliency = 60 // POWER_IDLE_RESILIENCY //col:78
SessionRITState = 61 // POWER_SESSION_RIT_STATE //col:79
SessionConnectNotification = 62 // POWER_SESSION_WINLOGON //col:80
SessionPowerCleanup = 63 //col:81
SessionLockState = 64 // POWER_SESSION_WINLOGON //col:82
SystemHiberbootState = 65 // BOOLEAN // fast startup supported //col:83
PlatformInformation = 66 // BOOLEAN // connected standby supported //col:84
PdcInvocation = 67 // (kernel-mode only) //col:85
MonitorInvocation = 68 // (kernel-mode only) //col:86
FirmwareTableInformationRegistered = 69 // (kernel-mode only) //col:87
SetShutdownSelectedTime = 70 // NULL //col:88
SuspendResumeInvocation = 71 // (kernel-mode only) //col:89
PlmPowerRequestCreate = 72 // in: COUNTED_REASON_CONTEXT, out: HANDLE //col:90
ScreenOff = 73 // NULL (PowerMonitorOff) //col:91
CsDeviceNotification = 74 // (kernel-mode only) //col:92
PlatformRole = 75 // POWER_PLATFORM_ROLE //col:93
LastResumePerformance = 76 // RESUME_PERFORMANCE //col:94
DisplayBurst = 77 // NULL (PowerMonitorOn) //col:95
ExitLatencySamplingPercentage = 78 //col:96
RegisterSpmPowerSettings = 79 // (kernel-mode only) //col:97
PlatformIdleStates = 80 // (kernel-mode only) //col:98
ProcessorIdleVeto = 81 // (kernel-mode only) // deprecated //col:99
PlatformIdleVeto = 82 // (kernel-mode only) // deprecated //col:100
SystemBatteryStatePrecise = 83 // SYSTEM_BATTERY_STATE //col:101
ThermalEvent = 84  // THERMAL_EVENT // PowerReportThermalEvent //col:102
PowerRequestActionInternal = 85 // POWER_REQUEST_ACTION_INTERNAL //col:103
BatteryDeviceState = 86 //col:104
PowerInformationInternal = 87 // POWER_INFORMATION_LEVEL_INTERNAL // PopPowerInformationInternal //col:105
ThermalStandby = 88 // NULL // shutdown with thermal standby as reason. //col:106
SystemHiberFileType = 89 // ULONG // zero ? reduced : full // powercfg.exe /h /type //col:107
PhysicalPowerButtonPress = 90 // BOOLEAN //col:108
QueryPotentialDripsConstraint = 91 // (kernel-mode only) //col:109
EnergyTrackerCreate = 92 //col:110
EnergyTrackerQuery = 93 //col:111
UpdateBlackBoxRecorder = 94 //col:112
SessionAllowExternalDmaDevices = 95 //col:113
SendSuspendResumeNotification = 96 // since WIN11 //col:114
PowerInformationLevelMaximum = 97 //col:115
POWER_REQUEST_CONTEXT_NOT_SPECIFIED = DIAGNOSTIC_REASON_NOT_SPECIFIED //col:142
POWER_REQUEST_SUPPORTED_TYPES_V1 = 3 // Windows 7 //col:260
POWER_REQUEST_SUPPORTED_TYPES_V2 = 9 // Windows 8 //col:261
POWER_REQUEST_SUPPORTED_TYPES_V3 = 5 // Windows 8.1 and Windows 10 TH1-TH2 //col:262
POWER_REQUEST_SUPPORTED_TYPES_V4 = 6 // Windows 10 RS1+ //col:263
)

type     PowerRequestDisplayRequiredInternal uint32
const(
    PowerRequestDisplayRequiredInternal POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 1  //col:164
    PowerRequestSystemRequiredInternal POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 2  //col:165
    PowerRequestAwayModeRequiredInternal POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 3  //col:166
    PowerRequestExecutionRequiredInternal // Windows 8+ POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 4  //col:167
    PowerRequestPerfBoostRequiredInternal // Windows 8+ POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 5  //col:168
    PowerRequestActiveLockScreenInternal // Windows 10 RS1+ (reserved on Windows 8) POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 6  //col:169
    // Values 6 and 7 are reserved for Windows 8 only POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 7  //col:170
    PowerRequestInternalInvalid POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 8  //col:171
    PowerRequestInternalUnknown POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 9  //col:172
    PowerRequestFullScreenVideoRequired  // Windows 8 only POWER_REQUEST_TYPE_INTERNAL // POWER_REQUEST_TYPE = 10  //col:173
)


type     SystemPowerState = 0 uint32
const(
    SystemPowerState  POWER_STATE_TYPE =  0  //col:192
    DevicePowerState POWER_STATE_TYPE = 2  //col:193
)


type     KernelRequester = 0 uint32
const(
    KernelRequester  REQUESTER_TYPE =  0  //col:217
    UserProcessRequester  REQUESTER_TYPE =  1  //col:218
    UserSharedServiceRequester  REQUESTER_TYPE =  2  //col:219
)


type     PowerStateSleeping1 = 0 uint32
const(
    PowerStateSleeping1  POWER_STATE_HANDLER_TYPE =  0  //col:310
    PowerStateSleeping2  POWER_STATE_HANDLER_TYPE =  1  //col:311
    PowerStateSleeping3  POWER_STATE_HANDLER_TYPE =  2  //col:312
    PowerStateSleeping4  POWER_STATE_HANDLER_TYPE =  3  //col:313
    PowerStateShutdownOff  POWER_STATE_HANDLER_TYPE =  4  //col:314
    PowerStateShutdownReset  POWER_STATE_HANDLER_TYPE =  5  //col:315
    PowerStateSleeping4Firmware  POWER_STATE_HANDLER_TYPE =  6  //col:316
    PowerStateMaximum  POWER_STATE_HANDLER_TYPE =  7  //col:317
)


type     PowerInternalAcpiInterfaceRegister uint32
const(
    PowerInternalAcpiInterfaceRegister POWER_INFORMATION_LEVEL_INTERNAL = 1  //col:362
    PowerInternalS0LowPowerIdleInfo // POWER_S0_LOW_POWER_IDLE_INFO POWER_INFORMATION_LEVEL_INTERNAL = 2  //col:363
    PowerInternalReapplyBrightnessSettings POWER_INFORMATION_LEVEL_INTERNAL = 3  //col:364
    PowerInternalUserAbsencePrediction // POWER_USER_ABSENCE_PREDICTION POWER_INFORMATION_LEVEL_INTERNAL = 4  //col:365
    PowerInternalUserAbsencePredictionCapability // POWER_USER_ABSENCE_PREDICTION_CAPABILITY POWER_INFORMATION_LEVEL_INTERNAL = 5  //col:366
    PowerInternalPoProcessorLatencyHint // POWER_PROCESSOR_LATENCY_HINT POWER_INFORMATION_LEVEL_INTERNAL = 6  //col:367
    PowerInternalStandbyNetworkRequest // POWER_STANDBY_NETWORK_REQUEST POWER_INFORMATION_LEVEL_INTERNAL = 7  //col:368
    PowerInternalDirtyTransitionInformation POWER_INFORMATION_LEVEL_INTERNAL = 8  //col:369
    PowerInternalSetBackgroundTaskState // POWER_SET_BACKGROUND_TASK_STATE POWER_INFORMATION_LEVEL_INTERNAL = 9  //col:370
    PowerInternalTtmOpenTerminal POWER_INFORMATION_LEVEL_INTERNAL = 10  //col:371
    PowerInternalTtmCreateTerminal // 10 POWER_INFORMATION_LEVEL_INTERNAL = 11  //col:372
    PowerInternalTtmEvacuateDevices POWER_INFORMATION_LEVEL_INTERNAL = 12  //col:373
    PowerInternalTtmCreateTerminalEventQueue POWER_INFORMATION_LEVEL_INTERNAL = 13  //col:374
    PowerInternalTtmGetTerminalEvent POWER_INFORMATION_LEVEL_INTERNAL = 14  //col:375
    PowerInternalTtmSetDefaultDeviceAssignment POWER_INFORMATION_LEVEL_INTERNAL = 15  //col:376
    PowerInternalTtmAssignDevice POWER_INFORMATION_LEVEL_INTERNAL = 16  //col:377
    PowerInternalTtmSetDisplayState POWER_INFORMATION_LEVEL_INTERNAL = 17  //col:378
    PowerInternalTtmSetDisplayTimeouts POWER_INFORMATION_LEVEL_INTERNAL = 18  //col:379
    PowerInternalBootSessionStandbyActivationInformation POWER_INFORMATION_LEVEL_INTERNAL = 19  //col:380
    PowerInternalSessionPowerState POWER_INFORMATION_LEVEL_INTERNAL = 20  //col:381
    PowerInternalSessionTerminalInput // 20 POWER_INFORMATION_LEVEL_INTERNAL = 21  //col:382
    PowerInternalSetWatchdog POWER_INFORMATION_LEVEL_INTERNAL = 22  //col:383
    PowerInternalPhysicalPowerButtonPressInfoAtBoot POWER_INFORMATION_LEVEL_INTERNAL = 23  //col:384
    PowerInternalExternalMonitorConnected POWER_INFORMATION_LEVEL_INTERNAL = 24  //col:385
    PowerInternalHighPrecisionBrightnessSettings POWER_INFORMATION_LEVEL_INTERNAL = 25  //col:386
    PowerInternalWinrtScreenToggle POWER_INFORMATION_LEVEL_INTERNAL = 26  //col:387
    PowerInternalPpmQosDisable POWER_INFORMATION_LEVEL_INTERNAL = 27  //col:388
    PowerInternalTransitionCheckpoint POWER_INFORMATION_LEVEL_INTERNAL = 28  //col:389
    PowerInternalInputControllerState POWER_INFORMATION_LEVEL_INTERNAL = 29  //col:390
    PowerInternalFirmwareResetReason POWER_INFORMATION_LEVEL_INTERNAL = 30  //col:391
    PowerInternalPpmSchedulerQosSupport // 30 POWER_INFORMATION_LEVEL_INTERNAL = 31  //col:392
    PowerInternalBootStatGet POWER_INFORMATION_LEVEL_INTERNAL = 32  //col:393
    PowerInternalBootStatSet POWER_INFORMATION_LEVEL_INTERNAL = 33  //col:394
    PowerInternalCallHasNotReturnedWatchdog POWER_INFORMATION_LEVEL_INTERNAL = 34  //col:395
    PowerInternalBootStatCheckIntegrity POWER_INFORMATION_LEVEL_INTERNAL = 35  //col:396
    PowerInternalBootStatRestoreDefaults // in: void POWER_INFORMATION_LEVEL_INTERNAL = 36  //col:397
    PowerInternalHostEsStateUpdate POWER_INFORMATION_LEVEL_INTERNAL = 37  //col:398
    PowerInternalGetPowerActionState POWER_INFORMATION_LEVEL_INTERNAL = 38  //col:399
    PowerInternalBootStatUnlock POWER_INFORMATION_LEVEL_INTERNAL = 39  //col:400
    PowerInternalWakeOnVoiceState POWER_INFORMATION_LEVEL_INTERNAL = 40  //col:401
    PowerInternalDeepSleepBlock // 40 POWER_INFORMATION_LEVEL_INTERNAL = 41  //col:402
    PowerInternalIsPoFxDevice POWER_INFORMATION_LEVEL_INTERNAL = 42  //col:403
    PowerInternalPowerTransitionExtensionAtBoot POWER_INFORMATION_LEVEL_INTERNAL = 43  //col:404
    PowerInternalProcessorBrandedFrequency // in: POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT out: POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT POWER_INFORMATION_LEVEL_INTERNAL = 44  //col:405
    PowerInternalTimeBrokerExpirationReason POWER_INFORMATION_LEVEL_INTERNAL = 45  //col:406
    PowerInternalNotifyUserShutdownStatus POWER_INFORMATION_LEVEL_INTERNAL = 46  //col:407
    PowerInternalPowerRequestTerminalCoreWindow POWER_INFORMATION_LEVEL_INTERNAL = 47  //col:408
    PowerInternalProcessorIdleVeto POWER_INFORMATION_LEVEL_INTERNAL = 48  //col:409
    PowerInternalPlatformIdleVeto POWER_INFORMATION_LEVEL_INTERNAL = 49  //col:410
    PowerInternalIsLongPowerButtonBugcheckEnabled POWER_INFORMATION_LEVEL_INTERNAL = 50  //col:411
    PowerInternalAutoChkCausedReboot // 50 POWER_INFORMATION_LEVEL_INTERNAL = 51  //col:412
    PowerInternalSetWakeAlarmOverride POWER_INFORMATION_LEVEL_INTERNAL = 52  //col:413
    PowerInternalDirectedFxAddTestDevice  POWER_INFORMATION_LEVEL_INTERNAL =  53  //col:415
    PowerInternalDirectedFxRemoveTestDevice POWER_INFORMATION_LEVEL_INTERNAL = 54  //col:416
    PowerInternalDirectedFxSetMode  POWER_INFORMATION_LEVEL_INTERNAL =  56  //col:418
    PowerInternalRegisterPowerPlane POWER_INFORMATION_LEVEL_INTERNAL = 56  //col:419
    PowerInternalSetDirectedDripsFlags POWER_INFORMATION_LEVEL_INTERNAL = 57  //col:420
    PowerInternalClearDirectedDripsFlags POWER_INFORMATION_LEVEL_INTERNAL = 58  //col:421
    PowerInternalRetrieveHiberFileResumeContext // 60 POWER_INFORMATION_LEVEL_INTERNAL = 59  //col:422
    PowerInternalReadHiberFilePage POWER_INFORMATION_LEVEL_INTERNAL = 60  //col:423
    PowerInternalLastBootSucceeded // out: BOOLEAN POWER_INFORMATION_LEVEL_INTERNAL = 61  //col:424
    PowerInternalQuerySleepStudyHelperRoutineBlock POWER_INFORMATION_LEVEL_INTERNAL = 62  //col:425
    PowerInternalDirectedDripsQueryCapabilities POWER_INFORMATION_LEVEL_INTERNAL = 63  //col:426
    PowerInternalClearConstraints POWER_INFORMATION_LEVEL_INTERNAL = 64  //col:427
    PowerInternalSoftParkVelocityEnabled POWER_INFORMATION_LEVEL_INTERNAL = 65  //col:428
    PowerInternalQueryIntelPepCapabilities POWER_INFORMATION_LEVEL_INTERNAL = 66  //col:429
    PowerInternalGetSystemIdleLoopEnablement // since WIN11 POWER_INFORMATION_LEVEL_INTERNAL = 67  //col:430
    PowerInternalGetVmPerfControlSupport POWER_INFORMATION_LEVEL_INTERNAL = 68  //col:431
    PowerInternalGetVmPerfControlConfig // 70 POWER_INFORMATION_LEVEL_INTERNAL = 69  //col:432
    PowerInternalSleepDetailedDiagUpdate POWER_INFORMATION_LEVEL_INTERNAL = 70  //col:433
    PowerInternalProcessorClassFrequencyBandsStats POWER_INFORMATION_LEVEL_INTERNAL = 71  //col:434
    PowerInternalHostGlobalUserPresenceStateUpdate POWER_INFORMATION_LEVEL_INTERNAL = 72  //col:435
    PowerInternalCpuNodeIdleIntervalStats POWER_INFORMATION_LEVEL_INTERNAL = 73  //col:436
    PowerInternalClassIdleIntervalStats POWER_INFORMATION_LEVEL_INTERNAL = 74  //col:437
    PowerInternalCpuNodeConcurrencyStats POWER_INFORMATION_LEVEL_INTERNAL = 75  //col:438
    PowerInternalClassConcurrencyStats POWER_INFORMATION_LEVEL_INTERNAL = 76  //col:439
    PowerInternalQueryProcMeasurementCapabilities POWER_INFORMATION_LEVEL_INTERNAL = 77  //col:440
    PowerInternalQueryProcMeasurementValues POWER_INFORMATION_LEVEL_INTERNAL = 78  //col:441
    PowerInternalPrepareForSystemInitiatedReboot // 80 POWER_INFORMATION_LEVEL_INTERNAL = 79  //col:442
    PowerInternalGetAdaptiveSessionState POWER_INFORMATION_LEVEL_INTERNAL = 80  //col:443
    PowerInternalSetConsoleLockedState POWER_INFORMATION_LEVEL_INTERNAL = 81  //col:444
    PowerInternalOverrideSystemInitiatedRebootState POWER_INFORMATION_LEVEL_INTERNAL = 82  //col:445
    PowerInternalFanImpactStats POWER_INFORMATION_LEVEL_INTERNAL = 83  //col:446
    PowerInternalFanRpmBuckets POWER_INFORMATION_LEVEL_INTERNAL = 84  //col:447
    PowerInternalPowerBootAppDiagInfo POWER_INFORMATION_LEVEL_INTERNAL = 85  //col:448
    PowerInternalUnregisterShutdownNotification // since 22H1 POWER_INFORMATION_LEVEL_INTERNAL = 86  //col:449
    PowerInternalManageTransitionStateRecord POWER_INFORMATION_LEVEL_INTERNAL = 87  //col:450
    PowerInformationInternalMaximum POWER_INFORMATION_LEVEL_INTERNAL = 88  //col:451
)


type     PoS0DisconnectedReasonNone uint32
const(
    PoS0DisconnectedReasonNone POWER_S0_DISCONNECTED_REASON = 1  //col:456
    PoS0DisconnectedReasonNonCompliantNic POWER_S0_DISCONNECTED_REASON = 2  //col:457
    PoS0DisconnectedReasonSettingPolicy POWER_S0_DISCONNECTED_REASON = 3  //col:458
    PoS0DisconnectedReasonEnforceDsPolicy POWER_S0_DISCONNECTED_REASON = 4  //col:459
    PoS0DisconnectedReasonCsChecksFailed POWER_S0_DISCONNECTED_REASON = 5  //col:460
    PoS0DisconnectedReasonSmartStandby POWER_S0_DISCONNECTED_REASON = 6  //col:461
    PoS0DisconnectedReasonMaximum POWER_S0_DISCONNECTED_REASON = 7  //col:462
)



type (
Ntpoapi interface{
 * Attribution 4.0 International ()(ok bool)//col:126
            _Field_size_()(ok bool)//col:160
#if ()(ok bool)//col:300
typedef NTSTATUS ()(ok bool)//col:339
typedef NTSTATUS ()(ok bool)//col:351
}
)

func NewNtpoapi() { return & ntpoapi{} }

func (n *ntpoapi) * Attribution 4.0 International ()(ok bool){//col:126
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTPOAPI_H
#define _NTPOAPI_H
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define SystemWakeSource 35
#define TraceServicePowerMessage 37
#define ProcessorLoad 38
#define PowerSettingNotificationName 58
#define SessionPowerCleanup 63
#define ExitLatencySamplingPercentage 78
#define BatteryDeviceState 86
#define EnergyTrackerCreate 92
#define EnergyTrackerQuery 93
#define UpdateBlackBoxRecorder 94
#define SessionAllowExternalDmaDevices 95
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
} PROCESSOR_POWER_INFORMATION, *PPROCESSOR_POWER_INFORMATION;*/
return true
}

func (n *ntpoapi)            _Field_size_()(ok bool){//col:160
/*            _Field_size_(StringCount) PUNICODE_STRING ReasonStrings;
        };
        UNICODE_STRING SimpleString;
    };
} COUNTED_REASON_CONTEXT, *PCOUNTED_REASON_CONTEXT;*/
return true
}

func (n *ntpoapi)#if ()(ok bool){//col:300
/*#if (PHNT_VERSION >= PHNT_WIN8)
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
} POWER_REQUEST, *PPOWER_REQUEST;*/
return true
}

func (n *ntpoapi)typedef NTSTATUS ()(ok bool){//col:339
/*typedef NTSTATUS (NTAPI *PENTER_STATE_SYSTEM_HANDLER)(
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
} POWER_STATE_HANDLER, *PPOWER_STATE_HANDLER;*/
return true
}

func (n *ntpoapi)typedef NTSTATUS ()(ok bool){//col:351
/*typedef NTSTATUS (NTAPI *PENTER_STATE_NOTIFY_HANDLER)(
    _In_ POWER_STATE_HANDLER_TYPE State,
    _In_ PVOID Context,
    _In_ BOOLEAN Entering
    );
typedef struct _POWER_STATE_NOTIFY_HANDLER
{
    PENTER_STATE_NOTIFY_HANDLER Handler;
    PVOID Context;
} POWER_STATE_NOTIFY_HANDLER, *PPOWER_STATE_NOTIFY_HANDLER;*/
return true
}



