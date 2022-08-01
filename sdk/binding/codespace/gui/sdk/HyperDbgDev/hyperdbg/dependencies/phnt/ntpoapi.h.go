package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpoapi.h.back

const(
_NTPOAPI_H =  //col:1
SystemPowerPolicyAc = 0 //col:2
SystemPowerPolicyDc = 1 //col:3
VerifySystemPolicyAc = 2 //col:4
VerifySystemPolicyDc = 3 //col:5
SystemPowerCapabilities = 4 //col:6
SystemBatteryState = 5 //col:7
SystemPowerStateHandler = 6 //col:8
ProcessorStateHandler = 7 //col:9
SystemPowerPolicyCurrent = 8 //col:10
AdministratorPowerPolicy = 9 //col:11
SystemReserveHiberFile = 10 //col:12
ProcessorInformation = 11 //col:13
SystemPowerInformation = 12 //col:14
ProcessorStateHandler2 = 13 //col:15
LastWakeTime = 14 //col:16
LastSleepTime = 15 //col:17
SystemExecutionState = 16 //col:18
SystemPowerStateNotifyHandler = 17 //col:19
ProcessorPowerPolicyAc = 18 //col:20
ProcessorPowerPolicyDc = 19 //col:21
VerifyProcessorPowerPolicyAc = 20 //col:22
VerifyProcessorPowerPolicyDc = 21 //col:23
ProcessorPowerPolicyCurrent = 22 //col:24
SystemPowerStateLogging = 23 //col:25
SystemPowerLoggingEntry = 24 //col:26
SetPowerSettingValue = 25 //col:27
NotifyUserPowerSetting = 26 //col:28
PowerInformationLevelUnused0 = 27 //col:29
SystemMonitorHiberBootPowerOff = 28 //col:30
SystemVideoState = 29 //col:31
TraceApplicationPowerMessage = 30 //col:32
TraceApplicationPowerMessageEnd = 31 //col:33
ProcessorPerfStates = 32 //col:34
ProcessorIdleStates = 33 //col:35
ProcessorCap = 34 //col:36
SystemWakeSource = 35 //col:37
SystemHiberFileInformation = 36 //col:38
TraceServicePowerMessage = 37 //col:39
ProcessorLoad = 38 //col:40
PowerShutdownNotification = 39 //col:41
MonitorCapabilities = 40 //col:42
SessionPowerInit = 41 //col:43
SessionDisplayState = 42 //col:44
PowerRequestCreate = 43 //col:45
PowerRequestAction = 44 //col:46
GetPowerRequestList = 45 //col:47
ProcessorInformationEx = 46 //col:48
NotifyUserModeLegacyPowerEvent = 47 //col:49
GroupPark = 48 //col:50
ProcessorIdleDomains = 49 //col:51
WakeTimerList = 50 //col:52
SystemHiberFileSize = 51 //col:53
ProcessorIdleStatesHv = 52 //col:54
ProcessorPerfStatesHv = 53 //col:55
ProcessorPerfCapHv = 54 //col:56
ProcessorSetIdle = 55 //col:57
LogicalProcessorIdling = 56 //col:58
UserPresence = 57 //col:59
PowerSettingNotificationName = 58 //col:60
GetPowerSettingValue = 59 //col:61
IdleResiliency = 60 //col:62
SessionRITState = 61 //col:63
SessionConnectNotification = 62 //col:64
SessionPowerCleanup = 63 //col:65
SessionLockState = 64 //col:66
SystemHiberbootState = 65 //col:67
PlatformInformation = 66 //col:68
PdcInvocation = 67 //col:69
MonitorInvocation = 68 //col:70
FirmwareTableInformationRegistered = 69 //col:71
SetShutdownSelectedTime = 70 //col:72
SuspendResumeInvocation = 71 //col:73
PlmPowerRequestCreate = 72 //col:74
ScreenOff = 73 //col:75
CsDeviceNotification = 74 //col:76
PlatformRole = 75 //col:77
LastResumePerformance = 76 //col:78
DisplayBurst = 77 //col:79
ExitLatencySamplingPercentage = 78 //col:80
RegisterSpmPowerSettings = 79 //col:81
PlatformIdleStates = 80 //col:82
ProcessorIdleVeto = 81 //col:83
PlatformIdleVeto = 82 //col:84
SystemBatteryStatePrecise = 83 //col:85
ThermalEvent = 84 //col:86
PowerRequestActionInternal = 85 //col:87
BatteryDeviceState = 86 //col:88
PowerInformationInternal = 87 //col:89
ThermalStandby = 88 //col:90
SystemHiberFileType = 89 //col:91
PhysicalPowerButtonPress = 90 //col:92
QueryPotentialDripsConstraint = 91 //col:93
EnergyTrackerCreate = 92 //col:94
EnergyTrackerQuery = 93 //col:95
UpdateBlackBoxRecorder = 94 //col:96
SessionAllowExternalDmaDevices = 95 //col:97
SendSuspendResumeNotification = 96 //col:98
PowerInformationLevelMaximum = 97 //col:99
POWER_REQUEST_CONTEXT_NOT_SPECIFIED = DIAGNOSTIC_REASON_NOT_SPECIFIED //col:100
POWER_REQUEST_SUPPORTED_TYPES_V1 = 3 //col:101
POWER_REQUEST_SUPPORTED_TYPES_V2 = 9 //col:102
POWER_REQUEST_SUPPORTED_TYPES_V3 = 5 //col:103
POWER_REQUEST_SUPPORTED_TYPES_V4 = 6 //col:104
)

const(
    PowerRequestDisplayRequiredInternal = 1  //col:3
    PowerRequestSystemRequiredInternal = 2  //col:4
    PowerRequestAwayModeRequiredInternal = 3  //col:5
    PowerRequestExecutionRequiredInternal  = 4  //col:6
    PowerRequestPerfBoostRequiredInternal  = 5  //col:7
    PowerRequestActiveLockScreenInternal  = 6  //col:8
    PowerRequestInternalInvalid = 7  //col:9
    PowerRequestInternalUnknown = 8  //col:10
    PowerRequestFullScreenVideoRequired   = 9  //col:11
)


const(
    SystemPowerState  =  0  //col:15
    DevicePowerState = 2  //col:16
)


const(
    KernelRequester  =  0  //col:20
    UserProcessRequester  =  1  //col:21
    UserSharedServiceRequester  =  2  //col:22
)


const(
    PowerStateSleeping1  =  0  //col:26
    PowerStateSleeping2  =  1  //col:27
    PowerStateSleeping3  =  2  //col:28
    PowerStateSleeping4  =  3  //col:29
    PowerStateShutdownOff  =  4  //col:30
    PowerStateShutdownReset  =  5  //col:31
    PowerStateSleeping4Firmware  =  6  //col:32
    PowerStateMaximum  =  7  //col:33
)


const(
    PowerInternalAcpiInterfaceRegister = 1  //col:37
    PowerInternalS0LowPowerIdleInfo  = 2  //col:38
    PowerInternalReapplyBrightnessSettings = 3  //col:39
    PowerInternalUserAbsencePrediction  = 4  //col:40
    PowerInternalUserAbsencePredictionCapability  = 5  //col:41
    PowerInternalPoProcessorLatencyHint  = 6  //col:42
    PowerInternalStandbyNetworkRequest  = 7  //col:43
    PowerInternalDirtyTransitionInformation = 8  //col:44
    PowerInternalSetBackgroundTaskState  = 9  //col:45
    PowerInternalTtmOpenTerminal = 10  //col:46
    PowerInternalTtmCreateTerminal  = 11  //col:47
    PowerInternalTtmEvacuateDevices = 12  //col:48
    PowerInternalTtmCreateTerminalEventQueue = 13  //col:49
    PowerInternalTtmGetTerminalEvent = 14  //col:50
    PowerInternalTtmSetDefaultDeviceAssignment = 15  //col:51
    PowerInternalTtmAssignDevice = 16  //col:52
    PowerInternalTtmSetDisplayState = 17  //col:53
    PowerInternalTtmSetDisplayTimeouts = 18  //col:54
    PowerInternalBootSessionStandbyActivationInformation = 19  //col:55
    PowerInternalSessionPowerState = 20  //col:56
    PowerInternalSessionTerminalInput  = 21  //col:57
    PowerInternalSetWatchdog = 22  //col:58
    PowerInternalPhysicalPowerButtonPressInfoAtBoot = 23  //col:59
    PowerInternalExternalMonitorConnected = 24  //col:60
    PowerInternalHighPrecisionBrightnessSettings = 25  //col:61
    PowerInternalWinrtScreenToggle = 26  //col:62
    PowerInternalPpmQosDisable = 27  //col:63
    PowerInternalTransitionCheckpoint = 28  //col:64
    PowerInternalInputControllerState = 29  //col:65
    PowerInternalFirmwareResetReason = 30  //col:66
    PowerInternalPpmSchedulerQosSupport  = 31  //col:67
    PowerInternalBootStatGet = 32  //col:68
    PowerInternalBootStatSet = 33  //col:69
    PowerInternalCallHasNotReturnedWatchdog = 34  //col:70
    PowerInternalBootStatCheckIntegrity = 35  //col:71
    PowerInternalBootStatRestoreDefaults  = 36  //col:72
    PowerInternalHostEsStateUpdate = 37  //col:73
    PowerInternalGetPowerActionState = 38  //col:74
    PowerInternalBootStatUnlock = 39  //col:75
    PowerInternalWakeOnVoiceState = 40  //col:76
    PowerInternalDeepSleepBlock  = 41  //col:77
    PowerInternalIsPoFxDevice = 42  //col:78
    PowerInternalPowerTransitionExtensionAtBoot = 43  //col:79
    PowerInternalProcessorBrandedFrequency  = 44  //col:80
    PowerInternalTimeBrokerExpirationReason = 45  //col:81
    PowerInternalNotifyUserShutdownStatus = 46  //col:82
    PowerInternalPowerRequestTerminalCoreWindow = 47  //col:83
    PowerInternalProcessorIdleVeto = 48  //col:84
    PowerInternalPlatformIdleVeto = 49  //col:85
    PowerInternalIsLongPowerButtonBugcheckEnabled = 50  //col:86
    PowerInternalAutoChkCausedReboot  = 51  //col:87
    PowerInternalSetWakeAlarmOverride = 52  //col:88
    PowerInternalDirectedFxAddTestDevice  =  53  //col:89
    PowerInternalDirectedFxRemoveTestDevice = 54  //col:90
    PowerInternalDirectedFxSetMode  =  56  //col:91
    PowerInternalRegisterPowerPlane = 56  //col:92
    PowerInternalSetDirectedDripsFlags = 57  //col:93
    PowerInternalClearDirectedDripsFlags = 58  //col:94
    PowerInternalRetrieveHiberFileResumeContext  = 59  //col:95
    PowerInternalReadHiberFilePage = 60  //col:96
    PowerInternalLastBootSucceeded  = 61  //col:97
    PowerInternalQuerySleepStudyHelperRoutineBlock = 62  //col:98
    PowerInternalDirectedDripsQueryCapabilities = 63  //col:99
    PowerInternalClearConstraints = 64  //col:100
    PowerInternalSoftParkVelocityEnabled = 65  //col:101
    PowerInternalQueryIntelPepCapabilities = 66  //col:102
    PowerInternalGetSystemIdleLoopEnablement  = 67  //col:103
    PowerInternalGetVmPerfControlSupport = 68  //col:104
    PowerInternalGetVmPerfControlConfig  = 69  //col:105
    PowerInternalSleepDetailedDiagUpdate = 70  //col:106
    PowerInternalProcessorClassFrequencyBandsStats = 71  //col:107
    PowerInternalHostGlobalUserPresenceStateUpdate = 72  //col:108
    PowerInternalCpuNodeIdleIntervalStats = 73  //col:109
    PowerInternalClassIdleIntervalStats = 74  //col:110
    PowerInternalCpuNodeConcurrencyStats = 75  //col:111
    PowerInternalClassConcurrencyStats = 76  //col:112
    PowerInternalQueryProcMeasurementCapabilities = 77  //col:113
    PowerInternalQueryProcMeasurementValues = 78  //col:114
    PowerInternalPrepareForSystemInitiatedReboot  = 79  //col:115
    PowerInternalGetAdaptiveSessionState = 80  //col:116
    PowerInternalSetConsoleLockedState = 81  //col:117
    PowerInternalOverrideSystemInitiatedRebootState = 82  //col:118
    PowerInternalFanImpactStats = 83  //col:119
    PowerInternalFanRpmBuckets = 84  //col:120
    PowerInternalPowerBootAppDiagInfo = 85  //col:121
    PowerInternalUnregisterShutdownNotification  = 86  //col:122
    PowerInternalManageTransitionStateRecord = 87  //col:123
    PowerInformationInternalMaximum = 88  //col:124
)


const(
    PoS0DisconnectedReasonNone = 1  //col:128
    PoS0DisconnectedReasonNonCompliantNic = 2  //col:129
    PoS0DisconnectedReasonSettingPolicy = 3  //col:130
    PoS0DisconnectedReasonEnforceDsPolicy = 4  //col:131
    PoS0DisconnectedReasonCsChecksFailed = 5  //col:132
    PoS0DisconnectedReasonSmartStandby = 6  //col:133
    PoS0DisconnectedReasonMaximum = 7  //col:134
)



type PROCESSOR_POWER_INFORMATION struct{
Number ULONG
MaxMhz ULONG
CurrentMhz ULONG
MhzLimit ULONG
MaxIdleState ULONG
CurrentIdleState ULONG
}


type SYSTEM_POWER_INFORMATION struct{
MaxIdlenessAllowed ULONG
Idleness ULONG
TimeRemaining ULONG
CoolingMode UCHAR
}


type SYSTEM_HIBERFILE_INFORMATION struct{
NumberOfMcbPairs ULONG
Mcb[1] LARGE_INTEGER
}


type COUNTED_REASON_CONTEXT struct{
Version ULONG
Flags ULONG
Union union
Struct struct
ResourceFileName UNICODE_STRING
ResourceReasonId USHORT
StringCount ULONG
PUNICODE_STRING _Field_size_(StringCount)
}


type POWER_REQUEST_ACTION struct{
PowerRequestHandle HANDLE
RequestType POWER_REQUEST_TYPE_INTERNAL
SetAction bool
ProcessHandle HANDLE
}


type SYSTEM_POWER_STATE_CONTEXT struct{
Union union
Struct struct
Reserved1 ULONG
TargetSystemState ULONG
EffectiveSystemState ULONG
CurrentSystemState ULONG
IgnoreHibernationPath ULONG
PseudoTransition ULONG
Reserved2 ULONG
}


type COUNTED_REASON_CONTEXT_RELATIVE struct{
Flags ULONG
Union union
Struct struct
ResourceFileNameOffset ULONG_PTR
ResourceReasonId USHORT
StringCount ULONG
SubstitutionStringsOffset ULONG_PTR
}


type DIAGNOSTIC_BUFFER struct{
Size SIZE_T
CallerType REQUESTER_TYPE
Union union
Struct struct
ProcessImageNameOffset ULONG_PTR
ProcessId ULONG
ServiceTag ULONG
}


type POWER_REQUEST struct{
Union union
Struct struct
SupportedRequestMask ULONG
PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V1] ULONG
DiagnosticBuffer DIAGNOSTIC_BUFFER
}


type POWER_REQUEST_LIST struct{
Count ULONG_PTR
PowerRequestOffsets[ANYSIZE_ARRAY] ULONG_PTR
}


type POWER_STATE_HANDLER struct{
Type POWER_STATE_HANDLER_TYPE
RtcWake bool
Spare[3] UCHAR
Handler PENTER_STATE_HANDLER
Context PVOID
}


type POWER_STATE_NOTIFY_HANDLER struct{
Handler PENTER_STATE_NOTIFY_HANDLER
Context PVOID
}


type POWER_REQUEST_ACTION_INTERNAL struct{
PowerRequestPointer PVOID
RequestType POWER_REQUEST_TYPE_INTERNAL
SetAction bool
}


type POWER_S0_LOW_POWER_IDLE_INFO struct{
DisconnectedReason POWER_S0_DISCONNECTED_REASON
Union union
Storage bool
WiFi bool
Mbn bool
Ethernet bool
Reserved bool
AsUCHAR UCHAR
}


type POWER_INFORMATION_INTERNAL_HEADER struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL
Version ULONG
}


type POWER_USER_ABSENCE_PREDICTION struct{
Header POWER_INFORMATION_INTERNAL_HEADER
ReturnTime LARGE_INTEGER
}


type POWER_USER_ABSENCE_PREDICTION_CAPABILITY struct{
AbsencePredictionCapability bool
}


type POWER_PROCESSOR_LATENCY_HINT struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER
Type ULONG
}


type POWER_STANDBY_NETWORK_REQUEST struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER
Active bool
}


type POWER_SET_BACKGROUND_TASK_STATE struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER
Engaged bool
}


type typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL
ProcessorNumber PROCESSOR_NUMBER
}


type typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT struct{
Version ULONG
NominalFrequency ULONG
}




