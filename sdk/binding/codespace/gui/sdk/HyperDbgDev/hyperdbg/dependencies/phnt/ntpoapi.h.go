package phnt


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
Number uint32 //col:3
MaxMhz uint32 //col:4
CurrentMhz uint32 //col:5
MhzLimit uint32 //col:6
MaxIdleState uint32 //col:7
CurrentIdleState uint32 //col:8
}



type SYSTEM_POWER_INFORMATION struct{
MaxIdlenessAllowed uint32 //col:12
Idleness uint32 //col:13
TimeRemaining uint32 //col:14
CoolingMode uint8 //col:15
}



type SYSTEM_HIBERFILE_INFORMATION struct{
NumberOfMcbPairs uint32 //col:19
Mcb[1] LARGE_INTEGER //col:20
}



type COUNTED_REASON_CONTEXT struct{
Version uint32 //col:24
Flags uint32 //col:25
Union union //col:26
Struct struct //col:28
ResourceFileName UNICODE_STRING //col:30
ResourceReasonId USHORT //col:31
StringCount uint32 //col:32
PUNICODE_STRING _Field_size_(StringCount) //col:33
}



type POWER_REQUEST_ACTION struct{
PowerRequestHandle HANDLE //col:40
RequestType POWER_REQUEST_TYPE_INTERNAL //col:41
SetAction bool //col:42
ProcessHandle HANDLE //col:43
}



type SYSTEM_POWER_STATE_CONTEXT struct{
Union union //col:47
Struct struct //col:49
Reserved1 uint32 //col:51
TargetSystemState uint32 //col:52
EffectiveSystemState uint32 //col:53
CurrentSystemState uint32 //col:54
IgnoreHibernationPath uint32 //col:55
PseudoTransition uint32 //col:56
Reserved2 uint32 //col:57
}



type COUNTED_REASON_CONTEXT_RELATIVE struct{
Flags uint32 //col:64
Union union //col:65
Struct struct //col:67
ResourceFileNameOffset ULONG_PTR //col:69
ResourceReasonId USHORT //col:70
StringCount uint32 //col:71
SubstitutionStringsOffset ULONG_PTR //col:72
}



type DIAGNOSTIC_BUFFER struct{
Size SIZE_T //col:79
CallerType REQUESTER_TYPE //col:80
Union union //col:81
Struct struct //col:83
ProcessImageNameOffset ULONG_PTR //col:85
ProcessId uint32 //col:86
ServiceTag uint32 //col:87
}



type POWER_REQUEST struct{
Union union //col:99
Struct struct //col:101
SupportedRequestMask uint32 //col:103
PowerRequestCount[POWER_REQUEST_SUPPORTED_TYPES_V1] uint32 //col:104
DiagnosticBuffer DIAGNOSTIC_BUFFER //col:105
}



type POWER_REQUEST_LIST struct{
Count ULONG_PTR //col:135
PowerRequestOffsets[ANYSIZE_ARRAY] ULONG_PTR //col:136
}



type POWER_STATE_HANDLER struct{
Type POWER_STATE_HANDLER_TYPE //col:140
RtcWake bool //col:141
Spare[3] uint8 //col:142
Handler PENTER_STATE_HANDLER //col:143
Context PVOID //col:144
}



type POWER_STATE_NOTIFY_HANDLER struct{
Handler PENTER_STATE_NOTIFY_HANDLER //col:148
Context PVOID //col:149
}



type POWER_REQUEST_ACTION_INTERNAL struct{
PowerRequestPointer PVOID //col:153
RequestType POWER_REQUEST_TYPE_INTERNAL //col:154
SetAction bool //col:155
}



type POWER_S0_LOW_POWER_IDLE_INFO struct{
DisconnectedReason POWER_S0_DISCONNECTED_REASON //col:159
Union union //col:160
Storage bool //col:162
WiFi bool //col:163
Mbn bool //col:164
Ethernet bool //col:165
Reserved bool //col:166
AsUCHAR uint8 //col:167
}



type POWER_INFORMATION_INTERNAL_HEADER struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL //col:179
Version uint32 //col:180
}



type POWER_USER_ABSENCE_PREDICTION struct{
Header POWER_INFORMATION_INTERNAL_HEADER //col:184
ReturnTime LARGE_INTEGER //col:185
}



type POWER_USER_ABSENCE_PREDICTION_CAPABILITY struct{
AbsencePredictionCapability bool //col:189
}



type POWER_PROCESSOR_LATENCY_HINT struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:193
Type uint32 //col:194
}



type POWER_STANDBY_NETWORK_REQUEST struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:198
Active bool //col:199
}



type POWER_SET_BACKGROUND_TASK_STATE struct{
PowerInformationInternalHeader POWER_INFORMATION_INTERNAL_HEADER //col:203
Engaged bool //col:204
}



type typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT struct{
InternalType POWER_INFORMATION_LEVEL_INTERNAL //col:208
ProcessorNumber PROCESSOR_NUMBER //col:209
}



type typedef struct POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT struct{
Version uint32 //col:213
NominalFrequency uint32 //col:214
}





