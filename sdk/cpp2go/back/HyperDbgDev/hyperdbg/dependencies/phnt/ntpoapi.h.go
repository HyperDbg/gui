package phnt


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

const(
    PowerRequestDisplayRequiredInternal = 1  //col:164
    PowerRequestSystemRequiredInternal = 2  //col:165
    PowerRequestAwayModeRequiredInternal = 3  //col:166
    PowerRequestExecutionRequiredInternal // Windows 8+ = 4  //col:167
    PowerRequestPerfBoostRequiredInternal // Windows 8+ = 5  //col:168
    PowerRequestActiveLockScreenInternal // Windows 10 RS1+ (reserved on Windows 8) = 6  //col:169
    PowerRequestInternalInvalid = 7  //col:171
    PowerRequestInternalUnknown = 8  //col:172
    PowerRequestFullScreenVideoRequired  // Windows 8 only = 9  //col:173
)


const(
    SystemPowerState  =  0  //col:192
    DevicePowerState = 2  //col:193
)


const(
    KernelRequester  =  0  //col:217
    UserProcessRequester  =  1  //col:218
    UserSharedServiceRequester  =  2  //col:219
)


const(
    PowerStateSleeping1  =  0  //col:310
    PowerStateSleeping2  =  1  //col:311
    PowerStateSleeping3  =  2  //col:312
    PowerStateSleeping4  =  3  //col:313
    PowerStateShutdownOff  =  4  //col:314
    PowerStateShutdownReset  =  5  //col:315
    PowerStateSleeping4Firmware  =  6  //col:316
    PowerStateMaximum  =  7  //col:317
)


const(
    PowerInternalAcpiInterfaceRegister = 1  //col:362
    PowerInternalS0LowPowerIdleInfo // POWER_S0_LOW_POWER_IDLE_INFO = 2  //col:363
    PowerInternalReapplyBrightnessSettings = 3  //col:364
    PowerInternalUserAbsencePrediction // POWER_USER_ABSENCE_PREDICTION = 4  //col:365
    PowerInternalUserAbsencePredictionCapability // POWER_USER_ABSENCE_PREDICTION_CAPABILITY = 5  //col:366
    PowerInternalPoProcessorLatencyHint // POWER_PROCESSOR_LATENCY_HINT = 6  //col:367
    PowerInternalStandbyNetworkRequest // POWER_STANDBY_NETWORK_REQUEST = 7  //col:368
    PowerInternalDirtyTransitionInformation = 8  //col:369
    PowerInternalSetBackgroundTaskState // POWER_SET_BACKGROUND_TASK_STATE = 9  //col:370
    PowerInternalTtmOpenTerminal = 10  //col:371
    PowerInternalTtmCreateTerminal // 10 = 11  //col:372
    PowerInternalTtmEvacuateDevices = 12  //col:373
    PowerInternalTtmCreateTerminalEventQueue = 13  //col:374
    PowerInternalTtmGetTerminalEvent = 14  //col:375
    PowerInternalTtmSetDefaultDeviceAssignment = 15  //col:376
    PowerInternalTtmAssignDevice = 16  //col:377
    PowerInternalTtmSetDisplayState = 17  //col:378
    PowerInternalTtmSetDisplayTimeouts = 18  //col:379
    PowerInternalBootSessionStandbyActivationInformation = 19  //col:380
    PowerInternalSessionPowerState = 20  //col:381
    PowerInternalSessionTerminalInput // 20 = 21  //col:382
    PowerInternalSetWatchdog = 22  //col:383
    PowerInternalPhysicalPowerButtonPressInfoAtBoot = 23  //col:384
    PowerInternalExternalMonitorConnected = 24  //col:385
    PowerInternalHighPrecisionBrightnessSettings = 25  //col:386
    PowerInternalWinrtScreenToggle = 26  //col:387
    PowerInternalPpmQosDisable = 27  //col:388
    PowerInternalTransitionCheckpoint = 28  //col:389
    PowerInternalInputControllerState = 29  //col:390
    PowerInternalFirmwareResetReason = 30  //col:391
    PowerInternalPpmSchedulerQosSupport // 30 = 31  //col:392
    PowerInternalBootStatGet = 32  //col:393
    PowerInternalBootStatSet = 33  //col:394
    PowerInternalCallHasNotReturnedWatchdog = 34  //col:395
    PowerInternalBootStatCheckIntegrity = 35  //col:396
    PowerInternalBootStatRestoreDefaults // in: void = 36  //col:397
    PowerInternalHostEsStateUpdate = 37  //col:398
    PowerInternalGetPowerActionState = 38  //col:399
    PowerInternalBootStatUnlock = 39  //col:400
    PowerInternalWakeOnVoiceState = 40  //col:401
    PowerInternalDeepSleepBlock // 40 = 41  //col:402
    PowerInternalIsPoFxDevice = 42  //col:403
    PowerInternalPowerTransitionExtensionAtBoot = 43  //col:404
    PowerInternalProcessorBrandedFrequency // in: POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_INPUT out: POWER_INTERNAL_PROCESSOR_BRANDED_FREQENCY_OUTPUT = 44  //col:405
    PowerInternalTimeBrokerExpirationReason = 45  //col:406
    PowerInternalNotifyUserShutdownStatus = 46  //col:407
    PowerInternalPowerRequestTerminalCoreWindow = 47  //col:408
    PowerInternalProcessorIdleVeto = 48  //col:409
    PowerInternalPlatformIdleVeto = 49  //col:410
    PowerInternalIsLongPowerButtonBugcheckEnabled = 50  //col:411
    PowerInternalAutoChkCausedReboot // 50 = 51  //col:412
    PowerInternalSetWakeAlarmOverride = 52  //col:413
    PowerInternalDirectedFxAddTestDevice  =  53  //col:415
    PowerInternalDirectedFxRemoveTestDevice = 54  //col:416
    PowerInternalDirectedFxSetMode  =  56  //col:418
    PowerInternalRegisterPowerPlane = 56  //col:419
    PowerInternalSetDirectedDripsFlags = 57  //col:420
    PowerInternalClearDirectedDripsFlags = 58  //col:421
    PowerInternalRetrieveHiberFileResumeContext // 60 = 59  //col:422
    PowerInternalReadHiberFilePage = 60  //col:423
    PowerInternalLastBootSucceeded // out: BOOLEAN = 61  //col:424
    PowerInternalQuerySleepStudyHelperRoutineBlock = 62  //col:425
    PowerInternalDirectedDripsQueryCapabilities = 63  //col:426
    PowerInternalClearConstraints = 64  //col:427
    PowerInternalSoftParkVelocityEnabled = 65  //col:428
    PowerInternalQueryIntelPepCapabilities = 66  //col:429
    PowerInternalGetSystemIdleLoopEnablement // since WIN11 = 67  //col:430
    PowerInternalGetVmPerfControlSupport = 68  //col:431
    PowerInternalGetVmPerfControlConfig // 70 = 69  //col:432
    PowerInternalSleepDetailedDiagUpdate = 70  //col:433
    PowerInternalProcessorClassFrequencyBandsStats = 71  //col:434
    PowerInternalHostGlobalUserPresenceStateUpdate = 72  //col:435
    PowerInternalCpuNodeIdleIntervalStats = 73  //col:436
    PowerInternalClassIdleIntervalStats = 74  //col:437
    PowerInternalCpuNodeConcurrencyStats = 75  //col:438
    PowerInternalClassConcurrencyStats = 76  //col:439
    PowerInternalQueryProcMeasurementCapabilities = 77  //col:440
    PowerInternalQueryProcMeasurementValues = 78  //col:441
    PowerInternalPrepareForSystemInitiatedReboot // 80 = 79  //col:442
    PowerInternalGetAdaptiveSessionState = 80  //col:443
    PowerInternalSetConsoleLockedState = 81  //col:444
    PowerInternalOverrideSystemInitiatedRebootState = 82  //col:445
    PowerInternalFanImpactStats = 83  //col:446
    PowerInternalFanRpmBuckets = 84  //col:447
    PowerInternalPowerBootAppDiagInfo = 85  //col:448
    PowerInternalUnregisterShutdownNotification // since 22H1 = 86  //col:449
    PowerInternalManageTransitionStateRecord = 87  //col:450
    PowerInformationInternalMaximum = 88  //col:451
)


const(
    PoS0DisconnectedReasonNone = 1  //col:456
    PoS0DisconnectedReasonNonCompliantNic = 2  //col:457
    PoS0DisconnectedReasonSettingPolicy = 3  //col:458
    PoS0DisconnectedReasonEnforceDsPolicy = 4  //col:459
    PoS0DisconnectedReasonCsChecksFailed = 5  //col:460
    PoS0DisconnectedReasonSmartStandby = 6  //col:461
    PoS0DisconnectedReasonMaximum = 7  //col:462
)



type (
Ntpoapi interface{
#if ()(ok bool)//col:126
            _Field_size_()(ok bool)//col:160
#if ()(ok bool)//col:300
typedef NTSTATUS ()(ok bool)//col:339
typedef NTSTATUS ()(ok bool)//col:351
}

)

func NewNtpoapi() { return & ntpoapi{} }

func (n *ntpoapi)#if ()(ok bool){//col:126























return true
}


func (n *ntpoapi)            _Field_size_()(ok bool){//col:160





return true
}


func (n *ntpoapi)#if ()(ok bool){//col:300


























return true
}


func (n *ntpoapi)typedef NTSTATUS ()(ok bool){//col:339


















return true
}


func (n *ntpoapi)typedef NTSTATUS ()(ok bool){//col:351










return true
}




