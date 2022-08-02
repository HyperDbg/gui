package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpoapi.h.back

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



type  _PROCESSOR_POWER_INFORMATION struct{
Number uint32 //col:11
MaxMhz uint32 //col:12
CurrentMhz uint32 //col:13
MhzLimit uint32 //col:14
MaxIdleState uint32 //col:15
CurrentIdleState uint32 //col:16
}


type  _SYSTEM_POWER_INFORMATION struct{
MaxIdlenessAllowed uint32 //col:18
Idleness uint32 //col:19
TimeRemaining uint32 //col:20
CoolingMode uint8 //col:21
}


type  _SYSTEM_HIBERFILE_INFORMATION struct{
NumberOfMcbPairs uint32 //col:23
Mcb[1] LARGE_INTEGER //col:24
}


type  _COUNTED_REASON_CONTEXT struct{
Version uint32 //col:36
Flags uint32 //col:37
Union union //col:38
Struct struct //col:40
ResourceFileName *int16 //col:42
ResourceReasonId uint16 //col:43
StringCount uint32 //col:44
PUNICODE_STRING _Field_size_(StringCount) //col:45
}


type  _POWER_REQUEST_ACTION struct{
PowerRequestHandle uintptr //col:46
RequestType POWER_REQUEST_TYPE_INTERNAL //col:47
SetAction bool //col:48
ProcessHandle uintptr //col:49
}


type  _SYSTEM_POWER_STATE_CONTEXT struct{
Union union //col:60
Struct struct //col:62
Reserved1 uint32 //col:64
TargetSystemState uint32 //col:65
EffectiveSystemState uint32 //col:66
CurrentSystemState uint32 //col:67
IgnoreHibernationPath uint32 //col:68
PseudoTransition uint32 //col:69
Reserved2 uint32 //col:70
}


type  _COUNTED_REASON_CONTEXT_RELATIVE struct{
Flags uint32 //col:75
Union union //col:76
Struct struct //col:78
ResourceFileNameOffset ULONG_PTR //col:80
ResourceReasonId uint16 //col:81
StringCount uint32 //col:82
SubstitutionStringsOffset ULONG_PTR //col:83
}


type  _DIAGNOSTIC_BUFFER struct{
Size int64 //col:90
CallerType REQUESTER_TYPE //col:91
Union union //col:92
Struct struct //col:94
ProcessImageNameOffset ULONG_PTR //col:96
ProcessId uint32 //col:97
ServiceTag uint32 //col:98
}


type  _POWER_REQUEST struct{
Union union //col:108
Struct struct //col:110
SupportedRequestMask uint32 //col:112
PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V1] uint32 //col:113
DiagnosticBuffer DIAGNOSTIC_BUFFER //col:114
}


type  _POWER_REQUEST_LIST struct{
Count ULONG_PTR //col:139
PowerRequestOffsets[ANYSIZE_ARRAY] ULONG_PTR //col:140
}


type  _POWER_STATE_HANDLER struct{
Type POWER_STATE_HANDLER_TYPE //col:147
RtcWake bool //col:148
Spare[3] uint8 //col:149
Handler PENTER_STATE_HANDLER //col:150
Context uintptr //col:151
}


type  _POWER_STATE_NOTIFY_HANDLER struct{
Handler PENTER_STATE_NOTIFY_HANDLER //col:152
Context uintptr //col:153
}


type  _POWER_REQUEST_ACTION_INTERNAL struct{
PowerRequestPointer uintptr //col:158
RequestType POWER_REQUEST_TYPE_INTERNAL //col:159
SetAction bool //col:160
}


type  _POWER_S0_LOW_POWER_IDLE_INFO struct{
DisconnectedReason POWER_S0_DISCONNECTED_REASON //col:170
Union union //col:171
Storage bool //col:173
WiFi bool //col:174
Mbn bool //col:175
Ethernet bool //col:176
Reserved bool //col:177
AsUCHAR uint8 //col:178
}


type  _POWER_INFORMATION_INTERNAL_HEADER struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL //col:183
Version uint32 //col:184
}


type  _POWER_USER_ABSENCE_PREDICTION struct{
Header POWER_INFORMATION_INTERNAL_HEADER //col:188
ReturnTime LARGE_INTEGER //col:189
}


type  _POWER_USER_ABSENCE_PREDICTION_CAPABILITY struct{
AbsencePredictionCapability bool //col:192
}


type  _POWER_PROCESSOR_LATENCY_HINT struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:197
Type uint32 //col:198
}


type  _POWER_STANDBY_NETWORK_REQUEST struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:202
Active bool //col:203
}


type  _POWER_SET_BACKGROUND_TASK_STATE struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:207
Engaged bool //col:208
}


type  POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL //col:212
ProcessorNumber PROCESSOR_NUMBER //col:213
}


type  POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT struct{
Version uint32 //col:217
NominalFrequency uint32 //col:218
}




