package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntexapi.h.back

const(
    FilterBootOptionOperationOpenSystemStore = 1  //col:3
    FilterBootOptionOperationSetElement = 2  //col:4
    FilterBootOptionOperationDeleteElement = 3  //col:5
    FilterBootOptionOperationMax = 4  //col:6
)


const(
    EventBasicInformation = 1  //col:10
)


const(
    MutantBasicInformation  = 1  //col:14
    MutantOwnerInformation  = 2  //col:15
)


const(
    SemaphoreBasicInformation = 1  //col:19
)


const(
    TimerBasicInformation  = 1  //col:23
)


const(
    TimerSetCoalescableTimer  = 1  //col:27
    MaxTimerInfoClass = 2  //col:28
)


const(
    WnfWellKnownStateName = 1  //col:32
    WnfPermanentStateName = 2  //col:33
    WnfPersistentStateName = 3  //col:34
    WnfTemporaryStateName = 4  //col:35
)


const(
    WnfInfoStateNameExist = 1  //col:39
    WnfInfoSubscribersPresent = 2  //col:40
    WnfInfoIsQuiescent = 3  //col:41
)


const(
    WnfDataScopeSystem = 1  //col:45
    WnfDataScopeSession = 2  //col:46
    WnfDataScopeUser = 3  //col:47
    WnfDataScopeProcess = 4  //col:48
    WnfDataScopeMachine  = 5  //col:49
    WnfDataScopePhysicalMachine  = 6  //col:50
)


const(
    WorkerFactoryTimeout  = 1  //col:54
    WorkerFactoryRetryTimeout  = 2  //col:55
    WorkerFactoryIdleTimeout  = 3  //col:56
    WorkerFactoryBindingCount  = 4  //col:57
    WorkerFactoryThreadMinimum  = 5  //col:58
    WorkerFactoryThreadMaximum  = 6  //col:59
    WorkerFactoryPaused  = 7  //col:60
    WorkerFactoryBasicInformation  = 8  //col:61
    WorkerFactoryAdjustThreadGoal = 9  //col:62
    WorkerFactoryCallbackType = 10  //col:63
    WorkerFactoryStackInformation  = 11  //col:64
    WorkerFactoryThreadBasePriority  = 12  //col:65
    WorkerFactoryTimeoutWaiters  = 13  //col:66
    WorkerFactoryFlags  = 14  //col:67
    WorkerFactoryThreadSoftMaximum  = 15  //col:68
    WorkerFactoryThreadCpuSets  = 16  //col:69
    MaxWorkerFactoryInfoClass = 17  //col:70
)


const(
    SystemBasicInformation  = 1  //col:74
    SystemProcessorInformation  = 2  //col:75
    SystemPerformanceInformation  = 3  //col:76
    SystemTimeOfDayInformation  = 4  //col:77
    SystemPathInformation  = 5  //col:78
    SystemProcessInformation  = 6  //col:79
    SystemCallCountInformation  = 7  //col:80
    SystemDeviceInformation  = 8  //col:81
    SystemProcessorPerformanceInformation  = 9  //col:82
    SystemFlagsInformation  = 10  //col:83
    SystemCallTimeInformation  = 11  //col:84
    SystemModuleInformation  = 12  //col:85
    SystemLocksInformation  = 13  //col:86
    SystemStackTraceInformation  = 14  //col:87
    SystemPagedPoolInformation  = 15  //col:88
    SystemNonPagedPoolInformation  = 16  //col:89
    SystemHandleInformation  = 17  //col:90
    SystemObjectInformation  = 18  //col:91
    SystemPageFileInformation  = 19  //col:92
    SystemVdmInstemulInformation  = 20  //col:93
    SystemVdmBopInformation  = 21  //col:94
    SystemFileCacheInformation  = 22  //col:95
    SystemPoolTagInformation  = 23  //col:96
    SystemInterruptInformation  = 24  //col:97
    SystemDpcBehaviorInformation  = 25  //col:98
    SystemFullMemoryInformation  = 26  //col:99
    SystemLoadGdiDriverInformation  = 27  //col:100
    SystemUnloadGdiDriverInformation  = 28  //col:101
    SystemTimeAdjustmentInformation  = 29  //col:102
    SystemSummaryMemoryInformation  = 30  //col:103
    SystemMirrorMemoryInformation  = 31  //col:104
    SystemPerformanceTraceInformation  = 32  //col:105
    SystemObsolete0  = 33  //col:106
    SystemExceptionInformation  = 34  //col:107
    SystemCrashDumpStateInformation  = 35  //col:108
    SystemKernelDebuggerInformation  = 36  //col:109
    SystemContextSwitchInformation  = 37  //col:110
    SystemRegistryQuotaInformation  = 38  //col:111
    SystemExtendServiceTableInformation  = 39  //col:112
    SystemPrioritySeperation  = 40  //col:113
    SystemVerifierAddDriverInformation  = 41  //col:114
    SystemVerifierRemoveDriverInformation  = 42  //col:115
    SystemProcessorIdleInformation  = 43  //col:116
    SystemLegacyDriverInformation  = 44  //col:117
    SystemCurrentTimeZoneInformation  = 45  //col:118
    SystemLookasideInformation  = 46  //col:119
    SystemTimeSlipNotification  = 47  //col:120
    SystemSessionCreate  = 48  //col:121
    SystemSessionDetach  = 49  //col:122
    SystemSessionInformation  = 50  //col:123
    SystemRangeStartInformation  = 51  //col:124
    SystemVerifierInformation  = 52  //col:125
    SystemVerifierThunkExtend  = 53  //col:126
    SystemSessionProcessInformation  = 54  //col:127
    SystemLoadGdiDriverInSystemSpace  = 55  //col:128
    SystemNumaProcessorMap  = 56  //col:129
    SystemPrefetcherInformation  = 57  //col:130
    SystemExtendedProcessInformation  = 58  //col:131
    SystemRecommendedSharedDataAlignment  = 59  //col:132
    SystemComPlusPackage  = 60  //col:133
    SystemNumaAvailableMemory  = 61  //col:134
    SystemProcessorPowerInformation  = 62  //col:135
    SystemEmulationBasicInformation  = 63  //col:136
    SystemEmulationProcessorInformation  = 64  //col:137
    SystemExtendedHandleInformation  = 65  //col:138
    SystemLostDelayedWriteInformation  = 66  //col:139
    SystemBigPoolInformation  = 67  //col:140
    SystemSessionPoolTagInformation  = 68  //col:141
    SystemSessionMappedViewInformation  = 69  //col:142
    SystemHotpatchInformation  = 70  //col:143
    SystemObjectSecurityMode  = 71  //col:144
    SystemWatchdogTimerHandler  = 72  //col:145
    SystemWatchdogTimerInformation  = 73  //col:146
    SystemLogicalProcessorInformation  = 74  //col:147
    SystemWow64SharedInformationObsolete  = 75  //col:148
    SystemRegisterFirmwareTableInformationHandler  = 76  //col:149
    SystemFirmwareTableInformation  = 77  //col:150
    SystemModuleInformationEx  = 78  //col:151
    SystemVerifierTriageInformation  = 79  //col:152
    SystemSuperfetchInformation  = 80  //col:153
    SystemMemoryListInformation  = 81  //col:154
    SystemFileCacheInformationEx  = 82  //col:155
    SystemThreadPriorityClientIdInformation  = 83  //col:156
    SystemProcessorIdleCycleTimeInformation  = 84  //col:157
    SystemVerifierCancellationInformation  = 85  //col:158
    SystemProcessorPowerInformationEx  = 86  //col:159
    SystemRefTraceInformation  = 87  //col:160
    SystemSpecialPoolInformation  = 88  //col:161
    SystemProcessIdInformation  = 89  //col:162
    SystemErrorPortInformation  = 90  //col:163
    SystemBootEnvironmentInformation  = 91  //col:164
    SystemHypervisorInformation  = 92  //col:165
    SystemVerifierInformationEx  = 93  //col:166
    SystemTimeZoneInformation  = 94  //col:167
    SystemImageFileExecutionOptionsInformation  = 95  //col:168
    SystemCoverageInformation  = 96  //col:169
    SystemPrefetchPatchInformation  = 97  //col:170
    SystemVerifierFaultsInformation  = 98  //col:171
    SystemSystemPartitionInformation  = 99  //col:172
    SystemSystemDiskInformation  = 100  //col:173
    SystemProcessorPerformanceDistribution  = 101  //col:174
    SystemNumaProximityNodeInformation  = 102  //col:175
    SystemDynamicTimeZoneInformation  = 103  //col:176
    SystemCodeIntegrityInformation  = 104  //col:177
    SystemProcessorMicrocodeUpdateInformation  = 105  //col:178
    SystemProcessorBrandString  = 106  //col:179
    SystemVirtualAddressInformation  = 107  //col:180
    SystemLogicalProcessorAndGroupInformation  = 108  //col:181
    SystemProcessorCycleTimeInformation  = 109  //col:182
    SystemStoreInformation  = 110  //col:183
    SystemRegistryAppendString  = 111  //col:184
    SystemAitSamplingValue  = 112  //col:185
    SystemVhdBootInformation  = 113  //col:186
    SystemCpuQuotaInformation  = 114  //col:187
    SystemNativeBasicInformation  = 115  //col:188
    SystemErrorPortTimeouts  = 116  //col:189
    SystemLowPriorityIoInformation  = 117  //col:190
    SystemTpmBootEntropyInformation  = 118  //col:191
    SystemVerifierCountersInformation  = 119  //col:192
    SystemPagedPoolInformationEx  = 120  //col:193
    SystemSystemPtesInformationEx  = 121  //col:194
    SystemNodeDistanceInformation  = 122  //col:195
    SystemAcpiAuditInformation  = 123  //col:196
    SystemBasicPerformanceInformation  = 124  //col:197
    SystemQueryPerformanceCounterInformation  = 125  //col:198
    SystemSessionBigPoolInformation  = 126  //col:199
    SystemBootGraphicsInformation  = 127  //col:200
    SystemScrubPhysicalMemoryInformation  = 128  //col:201
    SystemBadPageInformation = 129  //col:202
    SystemProcessorProfileControlArea  = 130  //col:203
    SystemCombinePhysicalMemoryInformation  = 131  //col:204
    SystemEntropyInterruptTimingInformation  = 132  //col:205
    SystemConsoleInformation  = 133  //col:206
    SystemPlatformBinaryInformation  = 134  //col:207
    SystemPolicyInformation  = 135  //col:208
    SystemHypervisorProcessorCountInformation  = 136  //col:209
    SystemDeviceDataInformation  = 137  //col:210
    SystemDeviceDataEnumerationInformation  = 138  //col:211
    SystemMemoryTopologyInformation  = 139  //col:212
    SystemMemoryChannelInformation  = 140  //col:213
    SystemBootLogoInformation  = 141  //col:214
    SystemProcessorPerformanceInformationEx  = 142  //col:215
    SystemCriticalProcessErrorLogInformation = 143  //col:216
    SystemSecureBootPolicyInformation  = 144  //col:217
    SystemPageFileInformationEx  = 145  //col:218
    SystemSecureBootInformation  = 146  //col:219
    SystemEntropyInterruptTimingRawInformation = 147  //col:220
    SystemPortableWorkspaceEfiLauncherInformation  = 148  //col:221
    SystemFullProcessInformation  = 149  //col:222
    SystemKernelDebuggerInformationEx  = 150  //col:223
    SystemBootMetadataInformation  = 151  //col:224
    SystemSoftRebootInformation  = 152  //col:225
    SystemElamCertificateInformation  = 153  //col:226
    SystemOfflineDumpConfigInformation  = 154  //col:227
    SystemProcessorFeaturesInformation  = 155  //col:228
    SystemRegistryReconciliationInformation  = 156  //col:229
    SystemEdidInformation  = 157  //col:230
    SystemManufacturingInformation  = 158  //col:231
    SystemEnergyEstimationConfigInformation  = 159  //col:232
    SystemHypervisorDetailInformation  = 160  //col:233
    SystemProcessorCycleStatsInformation  = 161  //col:234
    SystemVmGenerationCountInformation = 162  //col:235
    SystemTrustedPlatformModuleInformation  = 163  //col:236
    SystemKernelDebuggerFlags  = 164  //col:237
    SystemCodeIntegrityPolicyInformation  = 165  //col:238
    SystemIsolatedUserModeInformation  = 166  //col:239
    SystemHardwareSecurityTestInterfaceResultsInformation = 167  //col:240
    SystemSingleModuleInformation  = 168  //col:241
    SystemAllowedCpuSetsInformation = 169  //col:242
    SystemVsmProtectionInformation  = 170  //col:243
    SystemInterruptCpuSetsInformation  = 171  //col:244
    SystemSecureBootPolicyFullInformation  = 172  //col:245
    SystemCodeIntegrityPolicyFullInformation = 173  //col:246
    SystemAffinitizedInterruptProcessorInformation  = 174  //col:247
    SystemRootSiloInformation  = 175  //col:248
    SystemCpuSetInformation  = 176  //col:249
    SystemCpuSetTagInformation  = 177  //col:250
    SystemWin32WerStartCallout = 178  //col:251
    SystemSecureKernelProfileInformation  = 179  //col:252
    SystemCodeIntegrityPlatformManifestInformation  = 180  //col:253
    SystemInterruptSteeringInformation  = 181  //col:254
    SystemSupportedProcessorArchitectures  = 182  //col:255
    SystemMemoryUsageInformation  = 183  //col:256
    SystemCodeIntegrityCertificateInformation  = 184  //col:257
    SystemPhysicalMemoryInformation  = 185  //col:258
    SystemControlFlowTransition = 186  //col:259
    SystemKernelDebuggingAllowed  = 187  //col:260
    SystemActivityModerationExeState  = 188  //col:261
    SystemActivityModerationUserSettings  = 189  //col:262
    SystemCodeIntegrityPoliciesFullInformation = 190  //col:263
    SystemCodeIntegrityUnlockInformation  = 191  //col:264
    SystemIntegrityQuotaInformation = 192  //col:265
    SystemFlushInformation  = 193  //col:266
    SystemProcessorIdleMaskInformation  = 194  //col:267
    SystemSecureDumpEncryptionInformation = 195  //col:268
    SystemWriteConstraintInformation  = 196  //col:269
    SystemKernelVaShadowInformation  = 197  //col:270
    SystemHypervisorSharedPageInformation  = 198  //col:271
    SystemFirmwareBootPerformanceInformation = 199  //col:272
    SystemCodeIntegrityVerificationInformation  = 200  //col:273
    SystemFirmwarePartitionInformation  = 201  //col:274
    SystemSpeculationControlInformation  = 202  //col:275
    SystemDmaGuardPolicyInformation  = 203  //col:276
    SystemEnclaveLaunchControlInformation  = 204  //col:277
    SystemWorkloadAllowedCpuSetsInformation  = 205  //col:278
    SystemCodeIntegrityUnlockModeInformation = 206  //col:279
    SystemLeapSecondInformation  = 207  //col:280
    SystemFlags2Information  = 208  //col:281
    SystemSecurityModelInformation  = 209  //col:282
    SystemCodeIntegritySyntheticCacheInformation = 210  //col:283
    SystemFeatureConfigurationInformation  = 211  //col:284
    SystemFeatureConfigurationSectionInformation  = 212  //col:285
    SystemFeatureUsageSubscriptionInformation  = 213  //col:286
    SystemSecureSpeculationControlInformation  = 214  //col:287
    SystemSpacesBootInformation  = 215  //col:288
    SystemFwRamdiskInformation  = 216  //col:289
    SystemWheaIpmiHardwareInformation = 217  //col:290
    SystemDifSetRuleClassInformation = 218  //col:291
    SystemDifClearRuleClassInformation = 219  //col:292
    SystemDifApplyPluginVerificationOnDriver = 220  //col:293
    SystemDifRemovePluginVerificationOnDriver  = 221  //col:294
    SystemShadowStackInformation  = 222  //col:295
    SystemBuildVersionInformation  = 223  //col:296
    SystemPoolLimitInformation  = 224  //col:297
    SystemCodeIntegrityAddDynamicStore = 225  //col:298
    SystemCodeIntegrityClearDynamicStores = 226  //col:299
    SystemDifPoolTrackingInformation = 227  //col:300
    SystemPoolZeroingInformation  = 228  //col:301
    SystemDpcWatchdogInformation = 229  //col:302
    SystemDpcWatchdogInformation2 = 230  //col:303
    SystemSupportedProcessorArchitectures2  = 231  //col:304
    SystemSingleProcessorRelationshipInformation  = 232  //col:305
    SystemXfgCheckFailureInformation = 233  //col:306
    SystemIommuStateInformation  = 234  //col:307
    SystemHypervisorMinrootInformation  = 235  //col:308
    SystemHypervisorBootPagesInformation  = 236  //col:309
    SystemPointerAuthInformation  = 237  //col:310
    SystemSecureKernelDebuggerInformation = 238  //col:311
    SystemOriginalImageFeatureInformation = 239  //col:312
    MaxSystemInfoClass = 240  //col:313
)


const(
    EventTraceKernelVersionInformation  = 1  //col:317
    EventTraceGroupMaskInformation  = 2  //col:318
    EventTracePerformanceInformation  = 3  //col:319
    EventTraceTimeProfileInformation  = 4  //col:320
    EventTraceSessionSecurityInformation  = 5  //col:321
    EventTraceSpinlockInformation  = 6  //col:322
    EventTraceStackTracingInformation  = 7  //col:323
    EventTraceExecutiveResourceInformation  = 8  //col:324
    EventTraceHeapTracingInformation  = 9  //col:325
    EventTraceHeapSummaryTracingInformation  = 10  //col:326
    EventTracePoolTagFilterInformation  = 11  //col:327
    EventTracePebsTracingInformation  = 12  //col:328
    EventTraceProfileConfigInformation  = 13  //col:329
    EventTraceProfileSourceListInformation  = 14  //col:330
    EventTraceProfileEventListInformation  = 15  //col:331
    EventTraceProfileCounterListInformation  = 16  //col:332
    EventTraceStackCachingInformation  = 17  //col:333
    EventTraceObjectTypeFilterInformation  = 18  //col:334
    EventTraceSoftRestartInformation  = 19  //col:335
    EventTraceLastBranchConfigurationInformation  = 20  //col:336
    EventTraceLastBranchEventListInformation = 21  //col:337
    EventTraceProfileSourceAddInformation  = 22  //col:338
    EventTraceProfileSourceRemoveInformation  = 23  //col:339
    EventTraceProcessorTraceConfigurationInformation = 24  //col:340
    EventTraceProcessorTraceEventListInformation = 25  //col:341
    EventTraceCoverageSamplerInformation  = 26  //col:342
    EventTraceUnifiedStackCachingInformation  = 27  //col:343
    MaxEventTraceInfoClass = 28  //col:344
)


const(
    SystemCrashDumpDisable = 1  //col:348
    SystemCrashDumpReconfigure = 2  //col:349
    SystemCrashDumpInitializationComplete = 3  //col:350
)


const(
    WdActionSetTimeoutValue = 1  //col:354
    WdActionQueryTimeoutValue = 2  //col:355
    WdActionResetTimer = 3  //col:356
    WdActionStopTimer = 4  //col:357
    WdActionStartTimer = 5  //col:358
    WdActionSetTriggerAction = 6  //col:359
    WdActionQueryTriggerAction = 7  //col:360
    WdActionQueryState = 8  //col:361
)


const(
    WdInfoTimeoutValue  =  0  //col:365
    WdInfoResetTimer  =  1  //col:366
    WdInfoStopTimer  =  2  //col:367
    WdInfoStartTimer  =  3  //col:368
    WdInfoTriggerAction  =  4  //col:369
    WdInfoState  =  5  //col:370
    WdInfoTriggerReset  =  6  //col:371
    WdInfoNop  =  7  //col:372
    WdInfoGeneratedLastReset  =  8  //col:373
    WdInfoInvalid  =  9  //col:374
)


const(
    SystemFirmwareTableEnumerate = 1  //col:378
    SystemFirmwareTableGet = 2  //col:379
    SystemFirmwareTableMax = 3  //col:380
)


const(
    MemoryCaptureAccessedBits = 1  //col:384
    MemoryCaptureAndResetAccessedBits = 2  //col:385
    MemoryEmptyWorkingSets = 3  //col:386
    MemoryFlushModifiedList = 4  //col:387
    MemoryPurgeStandbyList = 5  //col:388
    MemoryPurgeLowPriorityStandbyList = 6  //col:389
    MemoryCommandMax = 7  //col:390
)


const(
    CoverageAllModules  =  0  //col:394
    CoverageSearchByHash  =  1  //col:395
    CoverageSearchByName  =  2  //col:396
)


const(
    SystemVaTypeAll = 1  //col:400
    SystemVaTypeNonPagedPool = 2  //col:401
    SystemVaTypePagedPool = 3  //col:402
    SystemVaTypeSystemCache = 4  //col:403
    SystemVaTypeSystemPtes = 5  //col:404
    SystemVaTypeSessionSpace = 6  //col:405
    SystemVaTypeMax = 7  //col:406
)


const(
    StorePageRequest  =  1  //col:410
    StoreStatsRequest  =  2   //col:411
    StoreCreateRequest  =  3   //col:412
    StoreDeleteRequest  =  4   //col:413
    StoreListRequest  =  5   //col:414
    Available1  =  6  //col:415
    StoreEmptyRequest  =  7  //col:416
    CacheListRequest  =  8   //col:417
    CacheCreateRequest  =  9   //col:418
    CacheDeleteRequest  =  10   //col:419
    CacheStoreCreateRequest  =  11   //col:420
    CacheStoreDeleteRequest  =  12   //col:421
    CacheStatsRequest  =  13   //col:422
    Available2  =  14  //col:423
    RegistrationRequest  =  15   //col:424
    GlobalCacheStatsRequest  =  16  //col:425
    StoreResizeRequest  =  17   //col:426
    CacheStoreResizeRequest  =  18   //col:427
    SmConfigRequest  =  19   //col:428
    StoreHighMemoryPriorityRequest  =  20   //col:429
    SystemStoreTrimRequest  =  21   //col:430
    MemCompressionInfoRequest  =  22    //col:431
    ProcessStoreInfoRequest  =  23   //col:432
    StoreInformationMax = 24  //col:433
)


const(
    StStatsLevelBasic  =  0  //col:437
    StStatsLevelIoStats  =  1  //col:438
    StStatsLevelRegionSpace  =  2   //col:439
    StStatsLevelSpaceBitmap  =  3   //col:440
    StStatsLevelMax  =  4  //col:441
)


const(
    StoreTypeInMemory = 0  //col:445
    StoreTypeFile = 1  //col:446
    StoreTypeMax = 2  //col:447
)


const(
    SmStoreManagerTypePhysical = 0  //col:451
    SmStoreManagerTypeVirtual = 1  //col:452
    SmStoreManagerTypeMax = 2  //col:453
)


const(
    SmConfigDirtyPageCompression  =  0  //col:457
    SmConfigAsyncInswap  =  1  //col:458
    SmConfigPrefetchSeekThreshold  =  2  //col:459
    SmConfigTypeMax  =  3  //col:460
)


const(
    TpmBootEntropyStructureUninitialized = 1  //col:464
    TpmBootEntropyDisabledByPolicy = 2  //col:465
    TpmBootEntropyNoTpmFound = 3  //col:466
    TpmBootEntropyTpmError = 4  //col:467
    TpmBootEntropySuccess = 5  //col:468
)


const(
    SystemPixelFormatUnknown = 1  //col:472
    SystemPixelFormatR8G8B8 = 2  //col:473
    SystemPixelFormatR8G8B8X8 = 3  //col:474
    SystemPixelFormatB8G8R8 = 4  //col:475
    SystemPixelFormatB8G8R8X8 = 5  //col:476
)


const(
    SystemProcessClassificationNormal = 1  //col:480
    SystemProcessClassificationSystem = 2  //col:481
    SystemProcessClassificationSecureSystem = 3  //col:482
    SystemProcessClassificationMemCompression = 4  //col:483
    SystemProcessClassificationRegistry  = 5  //col:484
    SystemProcessClassificationMaximum = 6  //col:485
)


const(
    SystemActivityModerationStateSystemManaged = 1  //col:489
    SystemActivityModerationStateUserManagedAllowThrottling = 2  //col:490
    SystemActivityModerationStateUserManagedDisableThrottling = 3  //col:491
    MaxSystemActivityModerationState = 4  //col:492
)


const(
    SystemActivityModerationAppTypeClassic = 1  //col:496
    SystemActivityModerationAppTypePackaged = 2  //col:497
    MaxSystemActivityModerationAppType = 3  //col:498
)


const(
    IommuStateBlock = 1  //col:502
    IommuStateUnblock = 2  //col:503
)


const(
    SysDbgQueryModuleInformation = 1  //col:507
    SysDbgQueryTraceInformation = 2  //col:508
    SysDbgSetTracepoint = 3  //col:509
    SysDbgSetSpecialCall  = 4  //col:510
    SysDbgClearSpecialCalls  = 5  //col:511
    SysDbgQuerySpecialCalls = 6  //col:512
    SysDbgBreakPoint = 7  //col:513
    SysDbgQueryVersion  = 8  //col:514
    SysDbgReadVirtual  = 9  //col:515
    SysDbgWriteVirtual  = 10  //col:516
    SysDbgReadPhysical  = 11  //col:517
    SysDbgWritePhysical  = 12  //col:518
    SysDbgReadControlSpace  = 13  //col:519
    SysDbgWriteControlSpace  = 14  //col:520
    SysDbgReadIoSpace  = 15  //col:521
    SysDbgWriteIoSpace  = 16  //col:522
    SysDbgReadMsr  = 17  //col:523
    SysDbgWriteMsr  = 18  //col:524
    SysDbgReadBusData  = 19  //col:525
    SysDbgWriteBusData  = 20  //col:526
    SysDbgCheckLowMemory  = 21  //col:527
    SysDbgEnableKernelDebugger = 22  //col:528
    SysDbgDisableKernelDebugger = 23  //col:529
    SysDbgGetAutoKdEnable = 24  //col:530
    SysDbgSetAutoKdEnable = 25  //col:531
    SysDbgGetPrintBufferSize = 26  //col:532
    SysDbgSetPrintBufferSize = 27  //col:533
    SysDbgGetKdUmExceptionEnable = 28  //col:534
    SysDbgSetKdUmExceptionEnable = 29  //col:535
    SysDbgGetTriageDump  = 30  //col:536
    SysDbgGetKdBlockEnable  = 31  //col:537
    SysDbgSetKdBlockEnable = 32  //col:538
    SysDbgRegisterForUmBreakInfo = 33  //col:539
    SysDbgGetUmBreakPid = 34  //col:540
    SysDbgClearUmBreakPid = 35  //col:541
    SysDbgGetUmAttachPid = 36  //col:542
    SysDbgClearUmAttachPid = 37  //col:543
    SysDbgGetLiveKernelDump  = 38  //col:544
    SysDbgKdPullRemoteFile  = 39  //col:545
    SysDbgMaxInfoClass = 40  //col:546
)


const(
    OptionAbortRetryIgnore = 1  //col:550
    OptionOk = 2  //col:551
    OptionOkCancel = 3  //col:552
    OptionRetryCancel = 4  //col:553
    OptionYesNo = 5  //col:554
    OptionYesNoCancel = 6  //col:555
    OptionShutdownSystem = 7  //col:556
    OptionOkNoWait = 8  //col:557
    OptionCancelTryContinue = 9  //col:558
)


const(
    ResponseReturnToCaller = 1  //col:562
    ResponseNotHandled = 2  //col:563
    ResponseAbort = 3  //col:564
    ResponseCancel = 4  //col:565
    ResponseIgnore = 5  //col:566
    ResponseNo = 6  //col:567
    ResponseOk = 7  //col:568
    ResponseRetry = 8  //col:569
    ResponseYes = 9  //col:570
    ResponseTryAgain = 10  //col:571
    ResponseContinue = 11  //col:572
)


const(
    StandardDesign = 1  //col:576
    NEC98x86 = 2  //col:577
    EndAlternatives = 3  //col:578
)


const(
    AtomBasicInformation = 1  //col:582
    AtomTableInformation = 2  //col:583
)


const(
    ShutdownNoReboot = 1  //col:587
    ShutdownReboot = 2  //col:588
    ShutdownPowerOff = 3  //col:589
    ShutdownRebootForRecovery  = 4  //col:590
)



type  _BOOT_ENTRY struct{
Version uint32 //col:13
Length uint32 //col:14
Id uint32 //col:15
Attributes uint32 //col:16
FriendlyNameOffset uint32 //col:17
BootFilePathOffset uint32 //col:18
OsOptionsLength uint32 //col:19
OsOptions[1] uint8 //col:20
}


type  _BOOT_ENTRY_LIST struct{
NextEntryOffset uint32 //col:18
BootEntry BOOT_ENTRY //col:19
}


type  _BOOT_OPTIONS struct{
Version uint32 //col:27
Length uint32 //col:28
Timeout uint32 //col:29
CurrentBootEntryId uint32 //col:30
NextBootEntryId uint32 //col:31
HeadlessRedirection[1] WCHAR //col:32
}


type  _FILE_PATH struct{
Version uint32 //col:34
Length uint32 //col:35
Type uint32 //col:36
FilePath[1] uint8 //col:37
}


type  _EFI_DRIVER_ENTRY struct{
Version uint32 //col:42
Length uint32 //col:43
Id uint32 //col:44
FriendlyNameOffset uint32 //col:45
DriverFilePathOffset uint32 //col:46
}


type  _EFI_DRIVER_ENTRY_LIST struct{
NextEntryOffset uint32 //col:47
DriverEntry EFI_DRIVER_ENTRY //col:48
}


type  _EVENT_BASIC_INFORMATION struct{
EventType EVENT_TYPE //col:52
EventState int32 //col:53
}


type  _MUTANT_BASIC_INFORMATION struct{
CurrentCount int32 //col:58
OwnedByCaller bool //col:59
AbandonedState bool //col:60
}


type  _MUTANT_OWNER_INFORMATION struct{
ClientId CLIENT_ID //col:62
}


type  _SEMAPHORE_BASIC_INFORMATION struct{
CurrentCount int32 //col:67
MaximumCount int32 //col:68
}


type  _TIMER_BASIC_INFORMATION struct{
RemainingTime LARGE_INTEGER //col:72
TimerState bool //col:73
}


type  _TIMER_SET_COALESCABLE_TIMER_INFO struct{
LARGE_INTEGER _In_ //col:82
PTIMER_APC_ROUTINE _In_opt_ //col:83
PVOID _In_opt_ //col:84
struct _In_opt_ //col:85
ULONG _In_opt_ //col:86
ULONG _In_ //col:87
PBOOLEAN _Out_opt_ //col:88
}


type  _T2_SET_PARAMETERS_V0 struct{
Version uint32 //col:88
Reserved uint32 //col:89
NoWakeTolerance LONGLONG //col:90
}


type  _WNF_STATE_NAME struct{
Data[2] uint32 //col:92
}


type  _WNF_TYPE_ID struct{
TypeId GUID //col:96
}


type  _WNF_DELIVERY_DESCRIPTOR struct{
SubscriptionId ULONGLONG //col:106
StateName WNF_STATE_NAME //col:107
ChangeStamp WNF_CHANGE_STAMP //col:108
StateDataSize uint32 //col:109
EventMask uint32 //col:110
TypeId WNF_TYPE_ID //col:111
StateDataOffset uint32 //col:112
}


type  _WORKER_FACTORY_BASIC_INFORMATION struct{
Timeout LARGE_INTEGER //col:133
RetryTimeout LARGE_INTEGER //col:134
IdleTimeout LARGE_INTEGER //col:135
Paused bool //col:136
TimerSet bool //col:137
QueuedToExWorker bool //col:138
MayCreate bool //col:139
CreateInProgress bool //col:140
InsertedIntoQueue bool //col:141
Shutdown bool //col:142
BindingCount uint32 //col:143
ThreadMinimum uint32 //col:144
ThreadMaximum uint32 //col:145
PendingWorkerCount uint32 //col:146
WaitingWorkerCount uint32 //col:147
TotalWorkerCount uint32 //col:148
ReleaseCount uint32 //col:149
InfiniteWaitGoal LONGLONG //col:150
StartRoutine uintptr //col:151
StartParameter uintptr //col:152
ProcessId uintptr //col:153
StackReserve int64 //col:154
StackCommit int64 //col:155
LastThreadCreationStatus NTSTATUS //col:156
}


type  _WORKER_FACTORY_DEFERRED_WORK struct{
_PORT_MESSAGE struct //col:140
AlpcSendMessagePort uintptr //col:141
AlpcSendMessageFlags uint32 //col:142
Flags uint32 //col:143
}


type  _SYSTEM_BASIC_INFORMATION struct{
Reserved uint32 //col:154
TimerResolution uint32 //col:155
PageSize uint32 //col:156
NumberOfPhysicalPages uint32 //col:157
LowestPhysicalPageNumber uint32 //col:158
HighestPhysicalPageNumber uint32 //col:159
AllocationGranularity uint32 //col:160
MinimumUserModeAddress ULONG_PTR //col:161
MaximumUserModeAddress ULONG_PTR //col:162
ActiveProcessorsAffinityMask KAFFINITY //col:163
NumberOfProcessors CCHAR //col:164
}


type  _SYSTEM_PROCESSOR_INFORMATION struct{
ProcessorArchitecture uint16 //col:162
ProcessorLevel uint16 //col:163
ProcessorRevision uint16 //col:164
MaximumProcessors uint16 //col:165
ProcessorFeatureBits uint32 //col:166
}


type  _SYSTEM_PERFORMANCE_INFORMATION struct{
IdleProcessTime LARGE_INTEGER //col:243
IoReadTransferCount LARGE_INTEGER //col:244
IoWriteTransferCount LARGE_INTEGER //col:245
IoOtherTransferCount LARGE_INTEGER //col:246
IoReadOperationCount uint32 //col:247
IoWriteOperationCount uint32 //col:248
IoOtherOperationCount uint32 //col:249
AvailablePages uint32 //col:250
CommittedPages uint32 //col:251
CommitLimit uint32 //col:252
PeakCommitment uint32 //col:253
PageFaultCount uint32 //col:254
CopyOnWriteCount uint32 //col:255
TransitionCount uint32 //col:256
CacheTransitionCount uint32 //col:257
DemandZeroCount uint32 //col:258
PageReadCount uint32 //col:259
PageReadIoCount uint32 //col:260
CacheReadCount uint32 //col:261
CacheIoCount uint32 //col:262
DirtyPagesWriteCount uint32 //col:263
DirtyWriteIoCount uint32 //col:264
MappedPagesWriteCount uint32 //col:265
MappedWriteIoCount uint32 //col:266
PagedPoolPages uint32 //col:267
NonPagedPoolPages uint32 //col:268
PagedPoolAllocs uint32 //col:269
PagedPoolFrees uint32 //col:270
NonPagedPoolAllocs uint32 //col:271
NonPagedPoolFrees uint32 //col:272
FreeSystemPtes uint32 //col:273
ResidentSystemCodePage uint32 //col:274
TotalSystemDriverPages uint32 //col:275
TotalSystemCodePages uint32 //col:276
NonPagedPoolLookasideHits uint32 //col:277
PagedPoolLookasideHits uint32 //col:278
AvailablePagedPoolPages uint32 //col:279
ResidentSystemCachePage uint32 //col:280
ResidentPagedPoolPage uint32 //col:281
ResidentSystemDriverPage uint32 //col:282
CcFastReadNoWait uint32 //col:283
CcFastReadWait uint32 //col:284
CcFastReadResourceMiss uint32 //col:285
CcFastReadNotPossible uint32 //col:286
CcFastMdlReadNoWait uint32 //col:287
CcFastMdlReadWait uint32 //col:288
CcFastMdlReadResourceMiss uint32 //col:289
CcFastMdlReadNotPossible uint32 //col:290
CcMapDataNoWait uint32 //col:291
CcMapDataWait uint32 //col:292
CcMapDataNoWaitMiss uint32 //col:293
CcMapDataWaitMiss uint32 //col:294
CcPinMappedDataCount uint32 //col:295
CcPinReadNoWait uint32 //col:296
CcPinReadWait uint32 //col:297
CcPinReadNoWaitMiss uint32 //col:298
CcPinReadWaitMiss uint32 //col:299
CcCopyReadNoWait uint32 //col:300
CcCopyReadWait uint32 //col:301
CcCopyReadNoWaitMiss uint32 //col:302
CcCopyReadWaitMiss uint32 //col:303
CcMdlReadNoWait uint32 //col:304
CcMdlReadWait uint32 //col:305
CcMdlReadNoWaitMiss uint32 //col:306
CcMdlReadWaitMiss uint32 //col:307
CcReadAheadIos uint32 //col:308
CcLazyWriteIos uint32 //col:309
CcLazyWritePages uint32 //col:310
CcDataFlushes uint32 //col:311
CcDataPages uint32 //col:312
ContextSwitches uint32 //col:313
FirstLevelTbFills uint32 //col:314
SecondLevelTbFills uint32 //col:315
SystemCalls uint32 //col:316
CcTotalDirtyPages ULONGLONG //col:317
CcDirtyPageThreshold ULONGLONG //col:318
ResidentAvailablePages LONGLONG //col:319
SharedCommittedPages ULONGLONG //col:320
}


type  _SYSTEM_TIMEOFDAY_INFORMATION struct{
BootTime LARGE_INTEGER //col:253
CurrentTime LARGE_INTEGER //col:254
TimeZoneBias LARGE_INTEGER //col:255
TimeZoneId uint32 //col:256
Reserved uint32 //col:257
BootTimeBias ULONGLONG //col:258
SleepTimeBias ULONGLONG //col:259
}


type  _SYSTEM_THREAD_INFORMATION struct{
KernelTime LARGE_INTEGER //col:267
UserTime LARGE_INTEGER //col:268
CreateTime LARGE_INTEGER //col:269
WaitTime uint32 //col:270
StartAddress uintptr //col:271
ClientId CLIENT_ID //col:272
Priority KPRIORITY //col:273
BasePriority KPRIORITY //col:274
ContextSwitches uint32 //col:275
ThreadState KTHREAD_STATE //col:276
WaitReason KWAIT_REASON //col:277
}


type  struct{
struct typedef //col:277
ThreadInfo SYSTEM_THREAD_INFORMATION //col:280
StackBase uintptr //col:281
StackLimit uintptr //col:282
Win32StartAddress uintptr //col:283
TebBase PTEB //col:284
Reserved2 ULONG_PTR //col:285
Reserved3 ULONG_PTR //col:286
Reserved4 ULONG_PTR //col:287
}


type  _SYSTEM_PROCESS_INFORMATION struct{
NextEntryOffset uint32 //col:317
NumberOfThreads uint32 //col:318
WorkingSetPrivateSize LARGE_INTEGER //col:319
HardFaultCount uint32 //col:320
NumberOfThreadsHighWatermark uint32 //col:321
CycleTime ULONGLONG //col:322
CreateTime LARGE_INTEGER //col:323
UserTime LARGE_INTEGER //col:324
KernelTime LARGE_INTEGER //col:325
ImageName *int16 //col:326
BasePriority KPRIORITY //col:327
UniqueProcessId uintptr //col:328
InheritedFromUniqueProcessId uintptr //col:329
HandleCount uint32 //col:330
SessionId uint32 //col:331
UniqueProcessKey ULONG_PTR //col:332
PeakVirtualSize int64 //col:333
VirtualSize int64 //col:334
PageFaultCount uint32 //col:335
PeakWorkingSetSize int64 //col:336
WorkingSetSize int64 //col:337
QuotaPeakPagedPoolUsage int64 //col:338
QuotaPagedPoolUsage int64 //col:339
QuotaPeakNonPagedPoolUsage int64 //col:340
QuotaNonPagedPoolUsage int64 //col:341
PagefileUsage int64 //col:342
PeakPagefileUsage int64 //col:343
PrivatePageCount int64 //col:344
ReadOperationCount LARGE_INTEGER //col:345
WriteOperationCount LARGE_INTEGER //col:346
OtherOperationCount LARGE_INTEGER //col:347
ReadTransferCount LARGE_INTEGER //col:348
WriteTransferCount LARGE_INTEGER //col:349
OtherTransferCount LARGE_INTEGER //col:350
Threads[1] SYSTEM_THREAD_INFORMATION //col:351
}


type  _SYSTEM_CALL_COUNT_INFORMATION struct{
Length uint32 //col:322
NumberOfTables uint32 //col:323
}


type  _SYSTEM_DEVICE_INFORMATION struct{
NumberOfDisks uint32 //col:331
NumberOfFloppies uint32 //col:332
NumberOfCdRoms uint32 //col:333
NumberOfTapes uint32 //col:334
NumberOfSerialPorts uint32 //col:335
NumberOfParallelPorts uint32 //col:336
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION struct{
IdleTime LARGE_INTEGER //col:340
KernelTime LARGE_INTEGER //col:341
UserTime LARGE_INTEGER //col:342
DpcTime LARGE_INTEGER //col:343
InterruptTime LARGE_INTEGER //col:344
InterruptCount uint32 //col:345
}


type  _SYSTEM_FLAGS_INFORMATION struct{
Flags uint32 //col:344
}


type  _SYSTEM_CALL_TIME_INFORMATION struct{
Length uint32 //col:350
TotalCalls uint32 //col:351
TimeOfCalls[1] LARGE_INTEGER //col:352
}


type  _RTL_PROCESS_LOCK_INFORMATION struct{
Address uintptr //col:363
Type uint16 //col:364
CreatorBackTraceIndex uint16 //col:365
OwningThread uintptr //col:366
LockCount int32 //col:367
ContentionCount uint32 //col:368
EntryCount uint32 //col:369
RecursionCount int32 //col:370
NumberOfWaitingShared uint32 //col:371
NumberOfWaitingExclusive uint32 //col:372
}


type  _RTL_PROCESS_LOCKS struct{
NumberOfLocks uint32 //col:368
Locks[1] RTL_PROCESS_LOCK_INFORMATION //col:369
}


type  _RTL_PROCESS_BACKTRACE_INFORMATION struct{
SymbolicBackTrace PCHAR //col:376
TraceCount uint32 //col:377
Index uint16 //col:378
Depth uint16 //col:379
BackTrace[32] uintptr //col:380
}


type  _RTL_PROCESS_BACKTRACES struct{
CommittedMemory uint32 //col:384
ReservedMemory uint32 //col:385
NumberOfBackTraceLookups uint32 //col:386
NumberOfBackTraces uint32 //col:387
BackTraces[1] RTL_PROCESS_BACKTRACE_INFORMATION //col:388
}


type  _SYSTEM_HANDLE_TABLE_ENTRY_INFO struct{
UniqueProcessId uint16 //col:394
CreatorBackTraceIndex uint16 //col:395
ObjectTypeIndex uint8 //col:396
HandleAttributes uint8 //col:397
HandleValue uint16 //col:398
Object uintptr //col:399
GrantedAccess uint32 //col:400
}


type  _SYSTEM_HANDLE_INFORMATION struct{
NumberOfHandles uint32 //col:399
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO //col:400
}


type  _SYSTEM_OBJECTTYPE_INFORMATION struct{
NextEntryOffset uint32 //col:413
NumberOfObjects uint32 //col:414
NumberOfHandles uint32 //col:415
TypeIndex uint32 //col:416
InvalidAttributes uint32 //col:417
GenericMapping GENERIC_MAPPING //col:418
ValidAccessMask uint32 //col:419
PoolType uint32 //col:420
SecurityRequired bool //col:421
WaitableObject bool //col:422
TypeName *int16 //col:423
}


type  _SYSTEM_OBJECT_INFORMATION struct{
NextEntryOffset uint32 //col:428
Object uintptr //col:429
CreatorUniqueProcess uintptr //col:430
CreatorBackTraceIndex uint16 //col:431
Flags uint16 //col:432
PointerCount int32 //col:433
HandleCount int32 //col:434
PagedPoolCharge uint32 //col:435
NonPagedPoolCharge uint32 //col:436
ExclusiveProcessId uintptr //col:437
SecurityDescriptor uintptr //col:438
NameInfo *int16 //col:439
}


type  _SYSTEM_PAGEFILE_INFORMATION struct{
NextEntryOffset uint32 //col:436
TotalSize uint32 //col:437
TotalInUse uint32 //col:438
PeakUsage uint32 //col:439
PageFileName *int16 //col:440
}


type  _SYSTEM_VDM_INSTEMUL_INFO struct{
SegmentNotPresent uint32 //col:473
VdmOpcode0F uint32 //col:474
OpcodeESPrefix uint32 //col:475
OpcodeCSPrefix uint32 //col:476
OpcodeSSPrefix uint32 //col:477
OpcodeDSPrefix uint32 //col:478
OpcodeFSPrefix uint32 //col:479
OpcodeGSPrefix uint32 //col:480
OpcodeOPER32Prefix uint32 //col:481
OpcodeADDR32Prefix uint32 //col:482
OpcodeINSB uint32 //col:483
OpcodeINSW uint32 //col:484
OpcodeOUTSB uint32 //col:485
OpcodeOUTSW uint32 //col:486
OpcodePUSHF uint32 //col:487
OpcodePOPF uint32 //col:488
OpcodeINTnn uint32 //col:489
OpcodeINTO uint32 //col:490
OpcodeIRET uint32 //col:491
OpcodeINBimm uint32 //col:492
OpcodeINWimm uint32 //col:493
OpcodeOUTBimm uint32 //col:494
OpcodeOUTWimm uint32 //col:495
OpcodeINB uint32 //col:496
OpcodeINW uint32 //col:497
OpcodeOUTB uint32 //col:498
OpcodeOUTW uint32 //col:499
OpcodeLOCKPrefix uint32 //col:500
OpcodeREPNEPrefix uint32 //col:501
OpcodeREPPrefix uint32 //col:502
OpcodeHLT uint32 //col:503
OpcodeCLI uint32 //col:504
OpcodeSTI uint32 //col:505
BopCount uint32 //col:506
}


type  _SYSTEM_FILECACHE_INFORMATION struct{
CurrentSize int64 //col:485
PeakSize int64 //col:486
PageFaultCount uint32 //col:487
MinimumWorkingSet int64 //col:488
MaximumWorkingSet int64 //col:489
CurrentSizeIncludingTransitionInPages int64 //col:490
PeakSizeIncludingTransitionInPages int64 //col:491
TransitionRePurposeCount uint32 //col:492
Flags uint32 //col:493
}


type  _SYSTEM_BASIC_WORKING_SET_INFORMATION struct{
CurrentSize int64 //col:491
PeakSize int64 //col:492
PageFaultCount uint32 //col:493
}


type  _SYSTEM_POOLTAG struct{
Union union //col:498
Tag[4] uint8 //col:500
TagUlong uint32 //col:501
}


type  _SYSTEM_POOLTAG_INFORMATION struct{
Count uint32 //col:510
TagInfo[1] SYSTEM_POOLTAG //col:511
}


type  _SYSTEM_INTERRUPT_INFORMATION struct{
ContextSwitches uint32 //col:519
DpcCount uint32 //col:520
DpcRate uint32 //col:521
TimeIncrement uint32 //col:522
DpcBypassCount uint32 //col:523
ApcBypassCount uint32 //col:524
}


type  _SYSTEM_DPC_BEHAVIOR_INFORMATION struct{
Spare uint32 //col:527
DpcQueueDepth uint32 //col:528
MinimumDpcRate uint32 //col:529
AdjustDpcThreshold uint32 //col:530
IdealDpcRate uint32 //col:531
}


type  _SYSTEM_QUERY_TIME_ADJUST_INFORMATION struct{
TimeAdjustment uint32 //col:533
TimeIncrement uint32 //col:534
Enable bool //col:535
}


type  _SYSTEM_QUERY_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG //col:539
TimeIncrement ULONGLONG //col:540
Enable bool //col:541
}


type  _SYSTEM_SET_TIME_ADJUST_INFORMATION struct{
TimeAdjustment uint32 //col:544
Enable bool //col:545
}


type  _SYSTEM_SET_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG //col:549
Enable bool //col:550
}


type  _EVENT_TRACE_VERSION_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:554
EventTraceKernelVersion uint32 //col:555
}


type  _PERFINFO_GROUPMASK struct{
Masks[PERF_NUM_MASKS] uint32 //col:558
}


type  _EVENT_TRACE_GROUPMASK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:564
TraceHandle TRACEHANDLE //col:565
EventTraceGroupMasks PERFINFO_GROUPMASK //col:566
}


type  _EVENT_TRACE_PERFORMANCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:569
LogfileBytesWritten LARGE_INTEGER //col:570
}


type  _EVENT_TRACE_TIME_PROFILE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:574
ProfileInterval uint32 //col:575
}


type  _EVENT_TRACE_SESSION_SECURITY_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:581
SecurityInformation uint32 //col:582
TraceHandle TRACEHANDLE //col:583
SecurityDescriptor[1] uint8 //col:584
}


type  _EVENT_TRACE_SPINLOCK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:589
SpinLockSpinThreshold uint32 //col:590
SpinLockAcquireSampleRate uint32 //col:591
SpinLockContentionSampleRate uint32 //col:592
SpinLockHoldThreshold uint32 //col:593
}


type  _EVENT_TRACE_SYSTEM_EVENT_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:595
TraceHandle TRACEHANDLE //col:596
HookId[1] uint32 //col:597
}


type  _EVENT_TRACE_EXECUTIVE_RESOURCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:602
ReleaseSamplingRate uint32 //col:603
ContentionSamplingRate uint32 //col:604
NumberOfExcessiveTimeouts uint32 //col:605
}


type  _EVENT_TRACE_HEAP_TRACING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:607
ProcessId uint32 //col:608
}


type  _EVENT_TRACE_TAG_FILTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:613
TraceHandle TRACEHANDLE //col:614
Filter[1] uint32 //col:615
}


type  _EVENT_TRACE_PROFILE_COUNTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:619
TraceHandle TRACEHANDLE //col:620
ProfileSource[1] uint32 //col:621
}


type  _EVENT_TRACE_PROFILE_LIST_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:625
Spare uint32 //col:626
_PROFILE_SOURCE_INFO* struct //col:627
}


type  _EVENT_TRACE_STACK_CACHING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:634
TraceHandle TRACEHANDLE //col:635
Enabled bool //col:636
Reserved[3] uint8 //col:637
CacheSize uint32 //col:638
BucketCount uint32 //col:639
}


type  _EVENT_TRACE_SOFT_RESTART_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:641
TraceHandle TRACEHANDLE //col:642
PersistTraceBuffers bool //col:643
FileName[1] WCHAR //col:644
}


type  _EVENT_TRACE_PROFILE_ADD_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:653
PerfEvtEventSelect bool //col:654
PerfEvtUnitSelect bool //col:655
PerfEvtType uint32 //col:656
CpuInfoHierarchy[0x3] uint32 //col:657
InitialInterval uint32 //col:658
AllowsHalt bool //col:659
Persist bool //col:660
ProfileSourceDescription[0x1] WCHAR //col:661
}


type  _EVENT_TRACE_PROFILE_REMOVE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:659
ProfileSource KPROFILE_SOURCE //col:660
CpuInfoHierarchy[0x3] uint32 //col:661
}


type  _EVENT_TRACE_COVERAGE_SAMPLER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:668
CoverageSamplerInformationClass bool //col:669
MajorVersion uint8 //col:670
MinorVersion uint8 //col:671
Reserved uint8 //col:672
SamplerHandle uintptr //col:673
}


type  _SYSTEM_EXCEPTION_INFORMATION struct{
AlignmentFixupCount uint32 //col:675
ExceptionDispatchCount uint32 //col:676
FloatingEmulationCount uint32 //col:677
ByteWordEmulationCount uint32 //col:678
}


type  _SYSTEM_CRASH_DUMP_STATE_INFORMATION struct{
CrashDumpConfigurationClass SYSTEM_CRASH_DUMP_CONFIGURATION_CLASS //col:679
}


type  _SYSTEM_KERNEL_DEBUGGER_INFORMATION struct{
KernelDebuggerEnabled bool //col:684
KernelDebuggerNotPresent bool //col:685
}


type  _SYSTEM_CONTEXT_SWITCH_INFORMATION struct{
ContextSwitches uint32 //col:699
FindAny uint32 //col:700
FindLast uint32 //col:701
FindIdeal uint32 //col:702
IdleAny uint32 //col:703
IdleCurrent uint32 //col:704
IdleLast uint32 //col:705
IdleIdeal uint32 //col:706
PreemptAny uint32 //col:707
PreemptCurrent uint32 //col:708
PreemptLast uint32 //col:709
SwitchToIdle uint32 //col:710
}


type  _SYSTEM_REGISTRY_QUOTA_INFORMATION struct{
RegistryQuotaAllowed uint32 //col:705
RegistryQuotaUsed uint32 //col:706
PagedPoolSize int64 //col:707
}


type  _SYSTEM_PROCESSOR_IDLE_INFORMATION struct{
IdleTime ULONGLONG //col:716
C1Time ULONGLONG //col:717
C2Time ULONGLONG //col:718
C3Time ULONGLONG //col:719
C1Transitions uint32 //col:720
C2Transitions uint32 //col:721
C3Transitions uint32 //col:722
Padding uint32 //col:723
}


type  _SYSTEM_LEGACY_DRIVER_INFORMATION struct{
VetoType uint32 //col:721
VetoList *int16 //col:722
}


type  _SYSTEM_LOOKASIDE_INFORMATION struct{
CurrentDepth uint16 //col:733
MaximumDepth uint16 //col:734
TotalAllocates uint32 //col:735
AllocateMisses uint32 //col:736
TotalFrees uint32 //col:737
FreeMisses uint32 //col:738
Type uint32 //col:739
Tag uint32 //col:740
Size uint32 //col:741
}


type  _SYSTEM_RANGE_START_INFORMATION struct{
SystemRangeStart ULONG_PTR //col:737
}


type  _SYSTEM_VERIFIER_INFORMATION_LEGACY  struct{
NextEntryOffset uint32 //col:765
Level uint32 //col:766
DriverName *int16 //col:767
RaiseIrqls uint32 //col:768
AcquireSpinLocks uint32 //col:769
SynchronizeExecutions uint32 //col:770
AllocationsAttempted uint32 //col:771
AllocationsSucceeded uint32 //col:772
AllocationsSucceededSpecialPool uint32 //col:773
AllocationsWithNoTag uint32 //col:774
TrimRequests uint32 //col:775
Trims uint32 //col:776
AllocationsFailed uint32 //col:777
AllocationsFailedDeliberately uint32 //col:778
Loads uint32 //col:779
Unloads uint32 //col:780
UnTrackedPool uint32 //col:781
CurrentPagedPoolAllocations uint32 //col:782
CurrentNonPagedPoolAllocations uint32 //col:783
PeakPagedPoolAllocations uint32 //col:784
PeakNonPagedPoolAllocations uint32 //col:785
PagedPoolUsageInBytes int64 //col:786
NonPagedPoolUsageInBytes int64 //col:787
PeakPagedPoolUsageInBytes int64 //col:788
PeakNonPagedPoolUsageInBytes int64 //col:789
}


type  _SYSTEM_VERIFIER_INFORMATION struct{
NextEntryOffset uint32 //col:796
Level uint32 //col:797
RuleClasses[2] uint32 //col:798
TriageContext uint32 //col:799
AreAllDriversBeingVerified uint32 //col:800
DriverName *int16 //col:801
RaiseIrqls uint32 //col:802
AcquireSpinLocks uint32 //col:803
SynchronizeExecutions uint32 //col:804
AllocationsAttempted uint32 //col:805
AllocationsSucceeded uint32 //col:806
AllocationsSucceededSpecialPool uint32 //col:807
AllocationsWithNoTag uint32 //col:808
TrimRequests uint32 //col:809
Trims uint32 //col:810
AllocationsFailed uint32 //col:811
AllocationsFailedDeliberately uint32 //col:812
Loads uint32 //col:813
Unloads uint32 //col:814
UnTrackedPool uint32 //col:815
CurrentPagedPoolAllocations uint32 //col:816
CurrentNonPagedPoolAllocations uint32 //col:817
PeakPagedPoolAllocations uint32 //col:818
PeakNonPagedPoolAllocations uint32 //col:819
PagedPoolUsageInBytes int64 //col:820
NonPagedPoolUsageInBytes int64 //col:821
PeakPagedPoolUsageInBytes int64 //col:822
PeakNonPagedPoolUsageInBytes int64 //col:823
}


type  _SYSTEM_SESSION_PROCESS_INFORMATION struct{
SessionId uint32 //col:802
SizeOfBuf uint32 //col:803
Buffer uintptr //col:804
}


type  _SYSTEM_GDI_DRIVER_INFORMATION struct{
DriverName *int16 //col:811
ImageAddress uintptr //col:812
SectionPointer uintptr //col:813
EntryPoint uintptr //col:814
_IMAGE_EXPORT_DIRECTORY* struct //col:815
ImageLength uint32 //col:816
}


type  _SYSTEM_NUMA_INFORMATION struct{
HighestNodeNumber uint32 //col:821
Reserved uint32 //col:822
Union union //col:823
ActiveProcessorsGroupAffinity[MAXIMUM_NODE_COUNT] GROUP_AFFINITY //col:825
AvailableMemory[MAXIMUM_NODE_COUNT] ULONGLONG //col:826
Pad[MAXIMUM_NODE_COUNT ULONGLONG //col:827
}


type  _SYSTEM_PROCESSOR_POWER_INFORMATION struct{
CurrentFrequency uint8 //col:845
ThermalLimitFrequency uint8 //col:846
ConstantThrottleFrequency uint8 //col:847
DegradedThrottleFrequency uint8 //col:848
LastBusyFrequency uint8 //col:849
LastC3Frequency uint8 //col:850
LastAdjustedBusyFrequency uint8 //col:851
ProcessorMinThrottle uint8 //col:852
ProcessorMaxThrottle uint8 //col:853
NumberOfFrequencies uint32 //col:854
PromotionCount uint32 //col:855
DemotionCount uint32 //col:856
ErrorCount uint32 //col:857
RetryCount uint32 //col:858
CurrentFrequencyTime ULONGLONG //col:859
CurrentProcessorTime ULONGLONG //col:860
CurrentProcessorIdleTime ULONGLONG //col:861
LastProcessorTime ULONGLONG //col:862
LastProcessorIdleTime ULONGLONG //col:863
Energy ULONGLONG //col:864
}


type  _SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX struct{
Object uintptr //col:856
UniqueProcessId ULONG_PTR //col:857
HandleValue ULONG_PTR //col:858
GrantedAccess uint32 //col:859
CreatorBackTraceIndex uint16 //col:860
ObjectTypeIndex uint16 //col:861
HandleAttributes uint32 //col:862
Reserved uint32 //col:863
}


type  _SYSTEM_HANDLE_INFORMATION_EX struct{
NumberOfHandles ULONG_PTR //col:862
Reserved ULONG_PTR //col:863
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX //col:864
}


type  _SYSTEM_BIGPOOL_ENTRY struct{
Union union //col:869
VirtualAddress uintptr //col:871
NonPaged ULONG_PTR //col:872
}


type  _SYSTEM_BIGPOOL_INFORMATION struct{
Count uint32 //col:881
AllocatedInfo[1] SYSTEM_BIGPOOL_ENTRY //col:882
}


type  _SYSTEM_POOL_ENTRY struct{
Allocated bool //col:893
Spare0 bool //col:894
AllocatorBackTraceIndex uint16 //col:895
Size uint32 //col:896
Union union //col:897
Tag[4] uint8 //col:899
TagUlong uint32 //col:900
ProcessChargedQuota uintptr //col:901
}


type  _SYSTEM_POOL_INFORMATION struct{
TotalSize int64 //col:904
FirstEntry uintptr //col:905
EntryOverhead uint16 //col:906
PoolTagPresent bool //col:907
Spare0 bool //col:908
NumberOfEntries uint32 //col:909
Entries[1] SYSTEM_POOL_ENTRY //col:910
}


type  _SYSTEM_SESSION_POOLTAG_INFORMATION struct{
NextEntryOffset int64 //col:911
SessionId uint32 //col:912
Count uint32 //col:913
TagInfo[1] SYSTEM_POOLTAG //col:914
}


type  _SYSTEM_SESSION_MAPPED_VIEW_INFORMATION struct{
NextEntryOffset int64 //col:919
SessionId uint32 //col:920
ViewFailures uint32 //col:921
NumberOfBytesAvailable int64 //col:922
NumberOfBytesAvailableContiguous int64 //col:923
}


type  _SYSTEM_WATCHDOG_HANDLER_INFORMATION  struct{
WdHandler PSYSTEM_WATCHDOG_HANDLER //col:924
Context uintptr //col:925
}


type  _SYSTEM_WATCHDOG_TIMER_INFORMATION struct{
WdInfoClass WATCHDOG_INFORMATION_CLASS //col:929
DataValue uint32 //col:930
}


type  _SYSTEM_FIRMWARE_TABLE_INFORMATION struct{
ProviderSignature uint32 //col:937
Action SYSTEM_FIRMWARE_TABLE_ACTION //col:938
TableID uint32 //col:939
TableBufferLength uint32 //col:940
TableBuffer[1] uint8 //col:941
}


type  _SYSTEM_FIRMWARE_TABLE_HANDLER struct{
ProviderSignature uint32 //col:944
Register bool //col:945
FirmwareTableHandler PFNFTH //col:946
DriverObject uintptr //col:947
}


type  _SYSTEM_MEMORY_LIST_INFORMATION struct{
ZeroPageCount ULONG_PTR //col:955
FreePageCount ULONG_PTR //col:956
ModifiedPageCount ULONG_PTR //col:957
ModifiedNoWritePageCount ULONG_PTR //col:958
BadPageCount ULONG_PTR //col:959
PageCountByPriority[8] ULONG_PTR //col:960
RepurposedPagesByPriority[8] ULONG_PTR //col:961
ModifiedPageCountPageFile ULONG_PTR //col:962
}


type  _SYSTEM_THREAD_CID_PRIORITY_INFORMATION struct{
ClientId CLIENT_ID //col:960
Priority KPRIORITY //col:961
}


type  _SYSTEM_PROCESSOR_IDLE_CYCLE_TIME_INFORMATION struct{
CycleTime ULONGLONG //col:964
}


type  _SYSTEM_VERIFIER_ISSUE struct{
IssueType ULONGLONG //col:970
Address uintptr //col:971
Parameters[2] ULONGLONG //col:972
}


type  _SYSTEM_VERIFIER_CANCELLATION_INFORMATION struct{
CancelProbability uint32 //col:979
CancelThreshold uint32 //col:980
CompletionThreshold uint32 //col:981
CancellationVerifierDisabled uint32 //col:982
AvailableIssues uint32 //col:983
Issues[128] SYSTEM_VERIFIER_ISSUE //col:984
}


type  _SYSTEM_REF_TRACE_INFORMATION struct{
TraceEnable bool //col:986
TracePermanent bool //col:987
TraceProcessName *int16 //col:988
TracePoolTags *int16 //col:989
}


type  _SYSTEM_SPECIAL_POOL_INFORMATION struct{
PoolTag uint32 //col:991
Flags uint32 //col:992
}


type  _SYSTEM_PROCESS_ID_INFORMATION struct{
ProcessId uintptr //col:996
ImageName *int16 //col:997
}


type  _SYSTEM_HYPERVISOR_QUERY_INFORMATION struct{
HypervisorConnected bool //col:1004
HypervisorDebuggingEnabled bool //col:1005
HypervisorPresent bool //col:1006
Spare0[5] bool //col:1007
EnabledEnlightenments ULONGLONG //col:1008
}


type  _SYSTEM_BOOT_ENVIRONMENT_INFORMATION struct{
BootIdentifier GUID //col:1022
FirmwareType FIRMWARE_TYPE //col:1023
Union union //col:1024
BootFlags ULONGLONG //col:1026
Struct struct //col:1027
DbgMenuOsSelection ULONGLONG //col:1029
DbgHiberBoot ULONGLONG //col:1030
DbgSoftBoot ULONGLONG //col:1031
DbgMeasuredLaunch ULONGLONG //col:1032
DbgMeasuredLaunchCapable ULONGLONG //col:1033
DbgSystemHiveReplace ULONGLONG //col:1034
DbgMeasuredLaunchSmmProtections ULONGLONG //col:1035
DbgMeasuredLaunchSmmLevel ULONGLONG //col:1036
}


type  _SYSTEM_IMAGE_FILE_EXECUTION_OPTIONS_INFORMATION struct{
FlagsToEnable uint32 //col:1029
FlagsToDisable uint32 //col:1030
}


type  _COVERAGE_MODULE_REQUEST struct{
RequestType COVERAGE_REQUEST_CODES //col:1037
Union union //col:1038
MD5Hash[16] uint8 //col:1040
ModuleName *int16 //col:1041
}


type  _COVERAGE_MODULE_INFO struct{
ModuleInfoSize uint32 //col:1046
IsBinaryLoaded uint32 //col:1047
ModulePathName *int16 //col:1048
CoverageSectionSize uint32 //col:1049
CoverageSection[1] uint8 //col:1050
}


type  _COVERAGE_MODULES struct{
ListAndReset uint32 //col:1053
NumberOfModules uint32 //col:1054
ModuleRequestInfo COVERAGE_MODULE_REQUEST //col:1055
Modules[1] COVERAGE_MODULE_INFO //col:1056
}


type  _SYSTEM_PREFETCH_PATCH_INFORMATION struct{
PrefetchPatchCount uint32 //col:1057
}


type  _SYSTEM_VERIFIER_FAULTS_INFORMATION struct{
Probability uint32 //col:1064
MaxProbability uint32 //col:1065
PoolTags *int16 //col:1066
Applications *int16 //col:1067
}


type  _SYSTEM_VERIFIER_INFORMATION_EX struct{
VerifyMode uint32 //col:1077
OptionChanges uint32 //col:1078
PreviousBucketName *int16 //col:1079
IrpCancelTimeoutMsec uint32 //col:1080
VerifierExtensionEnabled uint32 //col:1081
#ifdefWin64 #ifdef _WIN64 //col:1082
Reserved[1] uint32 //col:1083
#else #else //col:1084
Reserved[3] uint32 //col:1085
#endif #endif //col:1086
}


type  _SYSTEM_SYSTEM_PARTITION_INFORMATION struct{
SystemPartition *int16 //col:1081
}


type  _SYSTEM_SYSTEM_DISK_INFORMATION struct{
SystemDisk *int16 //col:1085
}


type  _SYSTEM_NUMA_PROXIMITY_MAP struct{
NodeProximityId uint32 //col:1090
NodeNumber uint16 //col:1091
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT struct{
Hits ULONGLONG //col:1095
PercentFrequency uint8 //col:1096
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT_WIN8 struct{
Hits uint32 //col:1100
PercentFrequency uint8 //col:1101
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_STATE_DISTRIBUTION struct{
ProcessorNumber uint32 //col:1106
StateCount uint32 //col:1107
States[1] SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT //col:1108
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_DISTRIBUTION struct{
ProcessorCount uint32 //col:1111
Offsets[1] uint32 //col:1112
}


type  _SYSTEM_CODEINTEGRITY_INFORMATION struct{
Length uint32 //col:1116
CodeIntegrityOptions uint32 //col:1117
}


type  _SYSTEM_PROCESSOR_MICROCODE_UPDATE_INFORMATION struct{
Operation uint32 //col:1120
}


type  _SYSTEM_VA_LIST_INFORMATION struct{
VirtualSize int64 //col:1127
VirtualPeak int64 //col:1128
VirtualLimit int64 //col:1129
AllocationFailures int64 //col:1130
}


type  _SYSTEM_STORE_INFORMATION struct{
ULONG _In_ //col:1134
STORE_INFORMATION_CLASS _In_ //col:1135
PVOID _Inout_ //col:1136
ULONG _Inout_ //col:1137
}


type  _SM_STATS_REQUEST struct{
Version uint32 //col:1142
DetailLevel uint32 //col:1143
StoreId uint32 //col:1144
BufferSize uint32 //col:1145
Buffer uintptr //col:1146
}


type  _ST_DATA_MGR_STATS struct{
RegionCount uint32 //col:1152
PagesStored uint32 //col:1153
UniquePagesStored uint32 //col:1154
LazyCleanupRegionCount uint32 //col:1155
RegionsInUse uint32 //col:1157
SpaceUsed uint32 //col:1158
}


type  _ST_IO_STATS_PERIOD struct{
PageCounts[5] uint32 //col:1157
}


type  _ST_IO_STATS struct{
PeriodCount uint32 //col:1162
Periods[64] ST_IO_STATS_PERIOD //col:1163
}


type  _ST_READ_LATENCY_BUCKET struct{
LatencyUs uint32 //col:1167
Count uint32 //col:1168
}


type  _ST_READ_LATENCY_STATS struct{
Buckets[8] ST_READ_LATENCY_BUCKET //col:1171
}


type  _ST_STATS_REGION_INFO struct{
SpaceUsed uint16 //col:1177
Priority uint8 //col:1178
Spare uint8 //col:1179
}


type  _ST_STATS_SPACE_BITMAP struct{
CompressedBytes int64 //col:1183
BytesPerBit uint32 //col:1184
StoreBitmap[1] uint8 //col:1185
}


type  _ST_STATS struct{
Version uint32 //col:1205
Level uint32 //col:1206
StoreType uint32 //col:1207
NoDuplication uint32 //col:1208
NoCompression uint32 //col:1209
EncryptionStrength uint32 //col:1210
VirtualRegions uint32 //col:1211
Spare0 uint32 //col:1212
Size uint32 //col:1213
CompressionFormat uint16 //col:1214
Spare uint16 //col:1215
Struct struct //col:1216
RegionSize uint32 //col:1218
RegionCount uint32 //col:1219
RegionCountMax uint32 //col:1220
Granularity uint32 //col:1221
UserData ST_DATA_MGR_STATS //col:1222
Metadata ST_DATA_MGR_STATS //col:1223
}


type  _SM_STORE_BASIC_PARAMS struct{
Union union //col:1231
Struct struct //col:1233
StoreType uint32 //col:1235
NoDuplication uint32 //col:1236
FailNoCompression uint32 //col:1237
NoCompression uint32 //col:1238
NoEncryption uint32 //col:1239
NoEvictOnAdd uint32 //col:1240
PerformsFileIo uint32 //col:1241
VdlNotSet uint32 //col:1242
UseIntermediateAddBuffer uint32 //col:1243
CompressNoHuff uint32 //col:1244
LockActiveRegions uint32 //col:1245
VirtualRegions uint32 //col:1246
Spare uint32 //col:1247
}


type  _SMKM_REGION_EXTENT struct{
RegionCount uint32 //col:1242
ByteOffset int64 //col:1243
}


type  _SMKM_FILE_INFO struct{
FileHandle uintptr //col:1253
_FILE_OBJECT struct //col:1254
_FILE_OBJECT struct //col:1255
_DEVICE_OBJECT struct //col:1256
VolumePnpHandle uintptr //col:1257
_IRP struct //col:1258
Extents PSMKM_REGION_EXTENT //col:1259
ExtentCount uint32 //col:1260
}


type  _SM_STORE_CACHE_BACKED_PARAMS struct{
SectorSize uint32 //col:1262
EncryptionKey PCHAR //col:1263
EncryptionKeySize uint32 //col:1264
FileInfo PSMKM_FILE_INFO //col:1265
EtaContext uintptr //col:1266
_RTL_BITMAP struct //col:1267
}


type  _SM_STORE_PARAMETERS struct{
Store SM_STORE_BASIC_PARAMS //col:1269
Priority uint32 //col:1270
Flags uint32 //col:1271
CacheBacked SM_STORE_CACHE_BACKED_PARAMS //col:1272
}


type  _SM_CREATE_REQUEST struct{
Version uint32 //col:1278
AcquireReference uint32 //col:1279
KeyedStore uint32 //col:1280
Spare uint32 //col:1281
Params SM_STORE_PARAMETERS //col:1282
StoreId uint32 //col:1283
}


type  _SM_DELETE_REQUEST struct{
Version uint32 //col:1284
Spare uint32 //col:1285
StoreId uint32 //col:1286
}


type  _SM_STORE_LIST_REQUEST struct{
Version uint32 //col:1292
StoreCount uint32 //col:1293
ExtendedRequest uint32 //col:1294
Spare uint32 //col:1295
StoreId[32] uint32 //col:1296
}


type  _SM_STORE_LIST_REQUEST_EX struct{
Request SM_STORE_LIST_REQUEST //col:1297
NameBuffer[32][64] WCHAR //col:1298
}


type  _SMC_CACHE_LIST_REQUEST struct{
Version uint32 //col:1304
CacheCount uint32 //col:1305
Spare uint32 //col:1306
CacheId[16] uint32 //col:1307
}


type  _SMC_CACHE_PARAMETERS struct{
CacheFileSize int64 //col:1314
StoreAlignment uint32 //col:1315
PerformsFileIo uint32 //col:1316
VdlNotSet uint32 //col:1317
Spare uint32 //col:1318
CacheFlags uint32 //col:1319
Priority uint32 //col:1320
}


type  _SMC_CACHE_CREATE_PARAMETERS struct{
CacheParameters SMC_CACHE_PARAMETERS //col:1319
TemplateFilePath[512] WCHAR //col:1320
}


type  _SMC_CACHE_CREATE_REQUEST struct{
Version uint32 //col:1326
Spare uint32 //col:1327
CacheId uint32 //col:1328
CacheCreateParams SMC_CACHE_CREATE_PARAMETERS //col:1329
}


type  _SMC_CACHE_DELETE_REQUEST struct{
Version uint32 //col:1332
Spare uint32 //col:1333
CacheId uint32 //col:1334
}


type  _SMC_STORE_CREATE_REQUEST struct{
Version uint32 //col:1341
Spare uint32 //col:1342
StoreParams SM_STORE_BASIC_PARAMS //col:1343
CacheId uint32 //col:1344
StoreManagerType SM_STORE_MANAGER_TYPE //col:1345
StoreId uint32 //col:1346
}


type  _SMC_STORE_DELETE_REQUEST struct{
Version uint32 //col:1349
Spare uint32 //col:1350
CacheId uint32 //col:1351
StoreManagerType SM_STORE_MANAGER_TYPE //col:1352
StoreId uint32 //col:1353
}


type  _SMC_CACHE_STATS struct{
TotalFileSize int64 //col:1363
StoreCount uint32 //col:1364
RegionCount uint32 //col:1365
RegionSizeBytes uint32 //col:1366
FileCount uint32 //col:1367
PerformsFileIo uint32 //col:1368
Spare uint32 //col:1369
StoreIds[16] uint32 //col:1370
PhysicalStoreBitmap uint32 //col:1371
Priority uint32 //col:1372
TemplateFilePath[512] WCHAR //col:1373
}


type  _SMC_CACHE_STATS_REQUEST struct{
Version uint32 //col:1371
NoFilePath uint32 //col:1372
Spare uint32 //col:1373
CacheId uint32 //col:1374
CacheStats SMC_CACHE_STATS //col:1375
}


type  _SM_REGISTRATION_INFO struct{
CachesUpdatedEvent uintptr //col:1375
}


type  _SM_REGISTRATION_REQUEST struct{
Version uint32 //col:1381
Spare uint32 //col:1382
RegInfo SM_REGISTRATION_INFO //col:1383
}


type  _SM_STORE_RESIZE_REQUEST struct{
Version uint32 //col:1390
AddRegions uint32 //col:1391
Spare uint32 //col:1392
StoreId uint32 //col:1393
NumberOfRegions uint32 //col:1394
_RTL_BITMAP struct //col:1395
}


type  _SMC_STORE_RESIZE_REQUEST struct{
Version uint32 //col:1400
AddRegions uint32 //col:1401
Spare uint32 //col:1402
CacheId uint32 //col:1403
StoreId uint32 //col:1404
StoreManagerType SM_STORE_MANAGER_TYPE //col:1405
RegionCount uint32 //col:1406
}


type  _SM_CONFIG_REQUEST struct{
Version uint32 //col:1407
Spare uint32 //col:1408
ConfigType uint32 //col:1409
ConfigValue uint32 //col:1410
}


type  _SM_STORE_HIGH_MEM_PRIORITY_REQUEST struct{
Version uint32 //col:1414
SetHighMemoryPriority uint32 //col:1415
Spare uint32 //col:1416
ProcessHandle uintptr //col:1417
}


type  _SM_SYSTEM_STORE_TRIM_REQUEST struct{
Version uint32 //col:1420
Spare uint32 //col:1421
PagesToTrim int64 //col:1422
}


type  _SM_MEM_COMPRESSION_INFO_REQUEST struct{
Version uint32 //col:1430
Spare uint32 //col:1431
CompressionPid uint32 //col:1432
WorkingSetSize uint32 //col:1433
TotalDataCompressed int64 //col:1434
TotalCompressedSize int64 //col:1435
TotalUniqueDataCompressed int64 //col:1436
}


type  _SYSTEM_REGISTRY_APPEND_STRING_PARAMETERS struct{
KeyHandle uintptr //col:1443
ValueNamePointer *uint32 //col:1444
RequiredLengthPointer PULONG //col:1445
Buffer PUCHAR //col:1446
BufferLength uint32 //col:1447
Type uint32 //col:1448
AppendBuffer PUCHAR //col:1449
AppendBufferLength uint32 //col:1450
CreateIfDoesntExist bool //col:1451
TruncateExistingValue bool //col:1452
}


type  _SYSTEM_VHD_BOOT_INFORMATION struct{
OsDiskIsVhd bool //col:1449
OsVhdFilePathOffset uint32 //col:1450
OsVhdParentVolume[1] WCHAR //col:1451
}


type  _PS_CPU_QUOTA_QUERY_ENTRY struct{
SessionId uint32 //col:1454
Weight uint32 //col:1455
}


type  _PS_CPU_QUOTA_QUERY_INFORMATION struct{
SessionCount uint32 //col:1459
SessionInformation[1] PS_CPU_QUOTA_QUERY_ENTRY //col:1460
}


type  _SYSTEM_ERROR_PORT_TIMEOUTS struct{
StartTimeout uint32 //col:1464
CommTimeout uint32 //col:1465
}


type  _SYSTEM_LOW_PRIORITY_IO_INFORMATION struct{
LowPriReadOperations uint32 //col:1477
LowPriWriteOperations uint32 //col:1478
KernelBumpedToNormalOperations uint32 //col:1479
LowPriPagingReadOperations uint32 //col:1480
KernelPagingReadsBumpedToNormal uint32 //col:1481
LowPriPagingWriteOperations uint32 //col:1482
KernelPagingWritesBumpedToNormal uint32 //col:1483
BoostedIrpCount uint32 //col:1484
BoostedPagingIrpCount uint32 //col:1485
BlanketBoostCount uint32 //col:1486
}


type  _TPM_BOOT_ENTROPY_NT_RESULT struct{
Policy ULONGLONG //col:1486
ResultCode TPM_BOOT_ENTROPY_RESULT_CODE //col:1487
ResultStatus NTSTATUS //col:1488
Time ULONGLONG //col:1489
EntropyLength uint32 //col:1490
EntropyData[40] uint8 //col:1491
}


type  _SYSTEM_VERIFIER_COUNTERS_INFORMATION struct{
Legacy SYSTEM_VERIFIER_INFORMATION //col:1513
RaiseIrqls uint32 //col:1514
AcquireSpinLocks uint32 //col:1515
SynchronizeExecutions uint32 //col:1516
AllocationsWithNoTag uint32 //col:1517
AllocationsFailed uint32 //col:1518
AllocationsFailedDeliberately uint32 //col:1519
LockedBytes int64 //col:1520
PeakLockedBytes int64 //col:1521
MappedLockedBytes int64 //col:1522
PeakMappedLockedBytes int64 //col:1523
MappedIoSpaceBytes int64 //col:1524
PeakMappedIoSpaceBytes int64 //col:1525
PagesForMdlBytes int64 //col:1526
PeakPagesForMdlBytes int64 //col:1527
ContiguousMemoryBytes int64 //col:1528
PeakContiguousMemoryBytes int64 //col:1529
ExecutePoolTypes uint32 //col:1530
ExecutePageProtections uint32 //col:1531
ExecutePageMappings uint32 //col:1532
ExecuteWriteSections uint32 //col:1533
SectionAlignmentFailures uint32 //col:1534
UnsupportedRelocs uint32 //col:1535
IATInExecutableSection uint32 //col:1536
}


type  _SYSTEM_ACPI_AUDIT_INFORMATION struct{
RsdpCount uint32 //col:1520
SameRsdt uint32 //col:1521
SlicPresent uint32 //col:1522
SlicDifferent uint32 //col:1523
}


type  _SYSTEM_BASIC_PERFORMANCE_INFORMATION struct{
AvailablePages int64 //col:1527
CommittedPages int64 //col:1528
CommitLimit int64 //col:1529
PeakCommitment int64 //col:1530
}


type  _QUERY_PERFORMANCE_COUNTER_FLAGS struct{
Union union //col:1536
Struct struct //col:1538
KernelTransition uint32 //col:1540
Reserved uint32 //col:1541
}


type  _SYSTEM_QUERY_PERFORMANCE_COUNTER_INFORMATION struct{
Version uint32 //col:1545
Flags QUERY_PERFORMANCE_COUNTER_FLAGS //col:1546
ValidFlags QUERY_PERFORMANCE_COUNTER_FLAGS //col:1547
}


type  _SYSTEM_BOOT_GRAPHICS_INFORMATION struct{
FrameBuffer LARGE_INTEGER //col:1555
Width uint32 //col:1556
Height uint32 //col:1557
PixelStride uint32 //col:1558
Flags uint32 //col:1559
Format SYSTEM_PIXEL_FORMAT //col:1560
DisplayRotation uint32 //col:1561
}


type  _MEMORY_SCRUB_INFORMATION struct{
Handle uintptr //col:1560
PagesScrubbed uint32 //col:1561
}


type  _PEBS_DS_SAVE_AREA32 struct{
BtsBufferBase uint32 //col:1573
BtsIndex uint32 //col:1574
BtsAbsoluteMaximum uint32 //col:1575
BtsInterruptThreshold uint32 //col:1576
PebsBufferBase uint32 //col:1577
PebsIndex uint32 //col:1578
PebsAbsoluteMaximum uint32 //col:1579
PebsInterruptThreshold uint32 //col:1580
PebsGpCounterReset[8] uint32 //col:1581
PebsFixedCounterReset[4] uint32 //col:1582
}


type  _PEBS_DS_SAVE_AREA64 struct{
BtsBufferBase ULONGLONG //col:1586
BtsIndex ULONGLONG //col:1587
BtsAbsoluteMaximum ULONGLONG //col:1588
BtsInterruptThreshold ULONGLONG //col:1589
PebsBufferBase ULONGLONG //col:1590
PebsIndex ULONGLONG //col:1591
PebsAbsoluteMaximum ULONGLONG //col:1592
PebsInterruptThreshold ULONGLONG //col:1593
PebsGpCounterReset[8] ULONGLONG //col:1594
PebsFixedCounterReset[4] ULONGLONG //col:1595
}


type  _PROCESSOR_PROFILE_CONTROL_AREA struct{
PebsDsSaveArea PEBS_DS_SAVE_AREA //col:1590
}


type  _SYSTEM_PROCESSOR_PROFILE_CONTROL_AREA struct{
ProcessorProfileControlArea PROCESSOR_PROFILE_CONTROL_AREA //col:1595
Allocate bool //col:1596
}


type  _MEMORY_COMBINE_INFORMATION struct{
Handle uintptr //col:1600
PagesCombined ULONG_PTR //col:1601
}


type  _MEMORY_COMBINE_INFORMATION_EX struct{
Handle uintptr //col:1606
PagesCombined ULONG_PTR //col:1607
Flags uint32 //col:1608
}


type  _MEMORY_COMBINE_INFORMATION_EX2 struct{
Handle uintptr //col:1613
PagesCombined ULONG_PTR //col:1614
Flags uint32 //col:1615
ProcessHandle uintptr //col:1616
}


type  _SYSTEM_ENTROPY_TIMING_INFORMATION struct{
(NTAPI VOID //col:1619
(NTAPI VOID //col:1620
InitializationContext uintptr //col:1621
}


type  _SYSTEM_CONSOLE_INFORMATION struct{
DriverLoaded uint32 //col:1624
Spare uint32 //col:1625
}


type  _SYSTEM_PLATFORM_BINARY_INFORMATION struct{
PhysicalAddress ULONG64 //col:1632
HandoffBuffer uintptr //col:1633
CommandLineBuffer uintptr //col:1634
HandoffBufferSize uint32 //col:1635
CommandLineBufferSize uint32 //col:1636
}


type  _SYSTEM_POLICY_INFORMATION struct{
InputData uintptr //col:1640
OutputData uintptr //col:1641
InputDataSize uint32 //col:1642
OutputDataSize uint32 //col:1643
Version uint32 //col:1644
}


type  _SYSTEM_HYPERVISOR_PROCESSOR_COUNT_INFORMATION struct{
NumberOfLogicalProcessors uint32 //col:1645
NumberOfCores uint32 //col:1646
}


type  _SYSTEM_DEVICE_DATA_INFORMATION struct{
DeviceId *int16 //col:1653
DataName *int16 //col:1654
DataType uint32 //col:1655
DataBufferLength uint32 //col:1656
DataBuffer uintptr //col:1657
}


type  _PHYSICAL_CHANNEL_RUN struct{
NodeNumber uint32 //col:1661
ChannelNumber uint32 //col:1662
BasePage ULONGLONG //col:1663
PageCount ULONGLONG //col:1664
Flags uint32 //col:1665
}


type  _SYSTEM_MEMORY_TOPOLOGY_INFORMATION struct{
NumberOfRuns ULONGLONG //col:1668
NumberOfNodes uint32 //col:1669
NumberOfChannels uint32 //col:1670
Run[1] PHYSICAL_CHANNEL_RUN //col:1671
}


type  _SYSTEM_MEMORY_CHANNEL_INFORMATION struct{
ChannelNumber uint32 //col:1677
ChannelHeatIndex uint32 //col:1678
TotalPageCount ULONGLONG //col:1679
ZeroPageCount ULONGLONG //col:1680
FreePageCount ULONGLONG //col:1681
StandbyPageCount ULONGLONG //col:1682
}


type  _SYSTEM_BOOT_LOGO_INFORMATION struct{
Flags uint32 //col:1682
BitmapOffset uint32 //col:1683
}


type  _SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION_EX struct{
IdleTime LARGE_INTEGER //col:1695
KernelTime LARGE_INTEGER //col:1696
UserTime LARGE_INTEGER //col:1697
DpcTime LARGE_INTEGER //col:1698
InterruptTime LARGE_INTEGER //col:1699
InterruptCount uint32 //col:1700
Spare0 uint32 //col:1701
AvailableTime LARGE_INTEGER //col:1702
Spare1 LARGE_INTEGER //col:1703
Spare2 LARGE_INTEGER //col:1704
}


type  _SYSTEM_SECUREBOOT_POLICY_INFORMATION struct{
PolicyPublisher GUID //col:1701
PolicyVersion uint32 //col:1702
PolicyOptions uint32 //col:1703
}


type  _SYSTEM_PAGEFILE_INFORMATION_EX struct{
Union union //col:1714
Info SYSTEM_PAGEFILE_INFORMATION //col:1716
Struct struct //col:1717
NextEntryOffset uint32 //col:1719
TotalSize uint32 //col:1720
TotalInUse uint32 //col:1721
PeakUsage uint32 //col:1722
PageFileName *int16 //col:1723
}


type  _SYSTEM_SECUREBOOT_INFORMATION struct{
SecureBootEnabled bool //col:1723
SecureBootCapable bool //col:1724
}


type  _PROCESS_DISK_COUNTERS struct{
BytesRead ULONGLONG //col:1731
BytesWritten ULONGLONG //col:1732
ReadOperationCount ULONGLONG //col:1733
WriteOperationCount ULONGLONG //col:1734
FlushOperationCount ULONGLONG //col:1735
}


type  _PROCESS_ENERGY_VALUES struct{
Cycles[4][2] ULONGLONG //col:1748
DiskEnergy ULONGLONG //col:1749
NetworkTailEnergy ULONGLONG //col:1750
MBBTailEnergy ULONGLONG //col:1751
NetworkTxRxBytes ULONGLONG //col:1752
MBBTxRxBytes ULONGLONG //col:1753
Union union //col:1754
Durations[3] ENERGY_STATE_DURATION //col:1756
Struct struct //col:1757
ForegroundDuration ENERGY_STATE_DURATION //col:1759
DesktopVisibleDuration ENERGY_STATE_DURATION //col:1760
PSMForegroundDuration ENERGY_STATE_DURATION //col:1761
}


type  _PROCESS_ENERGY_VALUES_EXTENSION struct{
Union union //col:1778
Timelines[14] TIMELINE_BITMAP //col:1780
Struct struct //col:1781
CpuTimeline TIMELINE_BITMAP //col:1783
DiskTimeline TIMELINE_BITMAP //col:1784
NetworkTimeline TIMELINE_BITMAP //col:1785
MBBTimeline TIMELINE_BITMAP //col:1786
ForegroundTimeline TIMELINE_BITMAP //col:1787
DesktopVisibleTimeline TIMELINE_BITMAP //col:1788
CompositionRenderedTimeline TIMELINE_BITMAP //col:1789
CompositionDirtyGeneratedTimeline TIMELINE_BITMAP //col:1790
CompositionDirtyPropagatedTimeline TIMELINE_BITMAP //col:1791
InputTimeline TIMELINE_BITMAP //col:1792
AudioInTimeline TIMELINE_BITMAP //col:1793
AudioOutTimeline TIMELINE_BITMAP //col:1794
DisplayRequiredTimeline TIMELINE_BITMAP //col:1795
KeyboardInputTimeline TIMELINE_BITMAP //col:1796
}


type  _PROCESS_EXTENDED_ENERGY_VALUES struct{
Base PROCESS_ENERGY_VALUES //col:1799
Extension PROCESS_ENERGY_VALUES_EXTENSION //col:1800
}


type  _SYSTEM_PROCESS_INFORMATION_EXTENSION struct{
DiskCounters PROCESS_DISK_COUNTERS //col:1813
ContextSwitches ULONGLONG //col:1814
Union union //col:1815
Flags uint32 //col:1817
Struct struct //col:1818
HasStrongId uint32 //col:1820
Classification uint32 //col:1821
BackgroundActivityModerated uint32 //col:1822
Spare uint32 //col:1823
}


type  _SYSTEM_PORTABLE_WORKSPACE_EFI_LAUNCHER_INFORMATION struct{
EfiLauncherEnabled bool //col:1827
}


type  _SYSTEM_KERNEL_DEBUGGER_INFORMATION_EX struct{
DebuggerAllowed bool //col:1833
DebuggerEnabled bool //col:1834
DebuggerPresent bool //col:1835
}


type  _SYSTEM_ELAM_CERTIFICATE_INFORMATION struct{
ElamDriverFile uintptr //col:1837
}


type  _OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V2 struct{
Version uint32 //col:1845
AbnormalResetOccurred uint32 //col:1846
OfflineMemoryDumpCapable uint32 //col:1847
ResetDataAddress LARGE_INTEGER //col:1848
ResetDataSize uint32 //col:1849
}


type  _OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V1 struct{
Version uint32 //col:1851
AbnormalResetOccurred uint32 //col:1852
OfflineMemoryDumpCapable uint32 //col:1853
}


type  _SYSTEM_PROCESSOR_FEATURES_INFORMATION struct{
ProcessorFeatureBits ULONGLONG //col:1856
Reserved[3] ULONGLONG //col:1857
}


type  _SYSTEM_EDID_INFORMATION struct{
Edid[128] uint8 //col:1860
}


type  _SYSTEM_MANUFACTURING_INFORMATION struct{
Options uint32 //col:1865
ProfileName *int16 //col:1866
}


type  _SYSTEM_ENERGY_ESTIMATION_CONFIG_INFORMATION struct{
Enabled bool //col:1869
}


type  _HV_DETAILS struct{
Data[4] uint32 //col:1873
}


type  _SYSTEM_HYPERVISOR_DETAIL_INFORMATION struct{
HvVendorAndMaxFunction HV_DETAILS //col:1883
HypervisorInterface HV_DETAILS //col:1884
HypervisorVersion HV_DETAILS //col:1885
HvFeatures HV_DETAILS //col:1886
HwFeatures HV_DETAILS //col:1887
EnlightenmentInfo HV_DETAILS //col:1888
ImplementationLimits HV_DETAILS //col:1889
}


type  _SYSTEM_PROCESSOR_CYCLE_STATS_INFORMATION struct{
Cycles[4][2] ULONGLONG //col:1887
}


type  _SYSTEM_TPM_INFORMATION struct{
Flags uint32 //col:1891
}


type  _SYSTEM_VSM_PROTECTION_INFORMATION struct{
DmaProtectionsAvailable bool //col:1898
DmaProtectionsInUse bool //col:1899
HardwareMbecAvailable bool //col:1900
ApicVirtualizationAvailable bool //col:1901
}


type  _SYSTEM_KERNEL_DEBUGGER_FLAGS struct{
KernelDebuggerIgnoreUmExceptions bool //col:1902
}


type  _SYSTEM_CODEINTEGRITYPOLICY_INFORMATION struct{
Options uint32 //col:1909
HVCIOptions uint32 //col:1910
Version ULONGLONG //col:1911
PolicyGuid GUID //col:1912
}


type  _SYSTEM_ISOLATED_USER_MODE_INFORMATION struct{
SecureKernelRunning bool //col:1924
HvciEnabled bool //col:1925
HvciStrictMode bool //col:1926
DebugEnabled bool //col:1927
FirmwarePageProtection bool //col:1928
EncryptionKeyAvailable bool //col:1929
SpareFlags bool //col:1930
TrustletRunning bool //col:1931
HvciDisableAllowed bool //col:1932
SpareFlags2 bool //col:1933
Spare0[6] bool //col:1934
Spare1 ULONGLONG //col:1935
}


type  _SYSTEM_SINGLE_MODULE_INFORMATION struct{
TargetModuleAddress uintptr //col:1929
ExInfo RTL_PROCESS_MODULE_INFORMATION_EX //col:1930
}


type  _SYSTEM_INTERRUPT_CPU_SET_INFORMATION struct{
Gsiv uint32 //col:1935
Group uint16 //col:1936
CpuSets ULONGLONG //col:1937
}


type  _SYSTEM_SECUREBOOT_POLICY_FULL_INFORMATION struct{
PolicyInformation SYSTEM_SECUREBOOT_POLICY_INFORMATION //col:1941
PolicySize uint32 //col:1942
Policy[1] uint8 //col:1943
}


type  _SYSTEM_ROOT_SILO_INFORMATION struct{
NumberOfSilos uint32 //col:1946
SiloIdList[1] uint32 //col:1947
}


type  _SYSTEM_CPU_SET_TAG_INFORMATION struct{
Tag ULONGLONG //col:1951
CpuSets[1] ULONGLONG //col:1952
}


type  _SYSTEM_SECURE_KERNEL_HYPERGUARD_PROFILE_INFORMATION struct{
ExtentCount uint32 //col:1978
ValidStructureSize uint32 //col:1979
NextExtentIndex uint32 //col:1980
ExtentRestart uint32 //col:1981
CycleCount uint32 //col:1982
TimeoutCount uint32 //col:1983
CycleTime ULONGLONG //col:1984
CycleTimeMax ULONGLONG //col:1985
ExtentTime ULONGLONG //col:1986
ExtentTimeIndex uint32 //col:1987
ExtentTimeMaxIndex uint32 //col:1988
ExtentTimeMax ULONGLONG //col:1989
HyperFlushTimeMax ULONGLONG //col:1990
TranslateVaTimeMax ULONGLONG //col:1991
DebugExemptionCount ULONGLONG //col:1992
TbHitCount ULONGLONG //col:1993
TbMissCount ULONGLONG //col:1994
VinaPendingYield ULONGLONG //col:1995
HashCycles ULONGLONG //col:1996
HistogramOffset uint32 //col:1997
HistogramBuckets uint32 //col:1998
HistogramShift uint32 //col:1999
Reserved1 uint32 //col:2000
PageNotPresentCount ULONGLONG //col:2001
}


type  _SYSTEM_SECUREBOOT_PLATFORM_MANIFEST_INFORMATION struct{
PlatformManifestSize uint32 //col:1983
PlatformManifest[1] uint8 //col:1984
}


type  _SYSTEM_INTERRUPT_STEERING_INFORMATION_INPUT struct{
Gsiv uint32 //col:1991
ControllerInterrupt uint8 //col:1992
EdgeInterrupt uint8 //col:1993
IsPrimaryInterrupt uint8 //col:1994
TargetAffinity GROUP_AFFINITY //col:1995
}


type  _SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION struct{
Machine uint32 //col:2001
KernelMode uint32 //col:2002
UserMode uint32 //col:2003
Native uint32 //col:2004
Process uint32 //col:2005
WoW64Container uint32 //col:2006
ReservedZero0 uint32 //col:2007
}


type  _SYSTEM_MEMORY_USAGE_INFORMATION struct{
TotalPhysicalBytes ULONGLONG //col:2011
AvailableBytes ULONGLONG //col:2012
ResidentAvailableBytes LONGLONG //col:2013
CommittedBytes ULONGLONG //col:2014
SharedCommittedBytes ULONGLONG //col:2015
CommitLimitBytes ULONGLONG //col:2016
PeakCommitmentBytes ULONGLONG //col:2017
}


type  _SYSTEM_CODEINTEGRITY_CERTIFICATE_INFORMATION struct{
ImageFile uintptr //col:2016
Type uint32 //col:2017
}


type  _SYSTEM_PHYSICAL_MEMORY_INFORMATION struct{
TotalPhysicalBytes ULONGLONG //col:2022
LowestPhysicalAddress ULONGLONG //col:2023
HighestPhysicalAddress ULONGLONG //col:2024
}


type  _SYSTEM_ACTIVITY_MODERATION_EXE_STATE  struct{
ExePathNt *int16 //col:2027
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE //col:2028
}


type  _SYSTEM_ACTIVITY_MODERATION_INFO struct{
Identifier *int16 //col:2033
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE //col:2034
AppType SYSTEM_ACTIVITY_MODERATION_APP_TYPE //col:2035
}


type  _SYSTEM_ACTIVITY_MODERATION_USER_SETTINGS struct{
UserKeyHandle uintptr //col:2037
}


type  _SYSTEM_CODEINTEGRITY_UNLOCK_INFORMATION struct{
Union union //col:2049
Flags uint32 //col:2051
Struct struct //col:2052
Locked uint32 //col:2054
UnlockApplied uint32 //col:2055
UnlockIdValid uint32 //col:2056
Reserved uint32 //col:2057
}


type  _SYSTEM_FLUSH_INFORMATION struct{
SupportedFlushMethods uint32 //col:2059
ProcessorCacheFlushSize uint32 //col:2060
SystemFlushCapabilities ULONGLONG //col:2061
Reserved[2] ULONGLONG //col:2062
}


type  _SYSTEM_WRITE_CONSTRAINT_INFORMATION struct{
WriteConstraintPolicy uint32 //col:2064
Reserved uint32 //col:2065
}


type  _SYSTEM_KERNEL_VA_SHADOW_INFORMATION struct{
Union union //col:2082
KvaShadowFlags uint32 //col:2084
Struct struct //col:2085
KvaShadowEnabled uint32 //col:2087
KvaShadowUserGlobal uint32 //col:2088
KvaShadowPcid uint32 //col:2089
KvaShadowInvpcid uint32 //col:2090
KvaShadowRequired uint32 //col:2091
KvaShadowRequiredAvailable uint32 //col:2092
InvalidPteBit uint32 //col:2093
L1DataCacheFlushSupported uint32 //col:2094
L1TerminalFaultMitigationPresent uint32 //col:2095
Reserved uint32 //col:2096
}


type  _SYSTEM_CODEINTEGRITYVERIFICATION_INFORMATION struct{
FileHandle uintptr //col:2090
ImageSize uint32 //col:2091
Image uintptr //col:2092
}


type  _SYSTEM_HYPERVISOR_SHARED_PAGE_INFORMATION struct{
HypervisorSharedUserVa uintptr //col:2094
}


type  _SYSTEM_FIRMWARE_PARTITION_INFORMATION struct{
FirmwarePartition *int16 //col:2098
}


type  _SYSTEM_SPECULATION_CONTROL_INFORMATION struct{
Union union //col:2134
Flags uint32 //col:2136
Struct struct //col:2137
BpbEnabled uint32 //col:2139
BpbDisabledSystemPolicy uint32 //col:2140
BpbDisabledNoHardwareSupport uint32 //col:2141
SpecCtrlEnumerated uint32 //col:2142
SpecCmdEnumerated uint32 //col:2143
IbrsPresent uint32 //col:2144
StibpPresent uint32 //col:2145
SmepPresent uint32 //col:2146
SpeculativeStoreBypassDisableAvailable uint32 //col:2147
SpeculativeStoreBypassDisableSupported uint32 //col:2148
SpeculativeStoreBypassDisabledSystemWide uint32 //col:2149
SpeculativeStoreBypassDisabledKernel uint32 //col:2150
SpeculativeStoreBypassDisableRequired uint32 //col:2151
BpbDisabledKernelToUser uint32 //col:2152
SpecCtrlRetpolineEnabled uint32 //col:2153
SpecCtrlImportOptimizationEnabled uint32 //col:2154
EnhancedIbrs uint32 //col:2155
HvL1tfStatusAvailable uint32 //col:2156
HvL1tfProcessorNotAffected uint32 //col:2157
HvL1tfMigitationEnabled uint32 //col:2158
HvL1tfMigitationNotEnabled_Hardware uint32 //col:2159
HvL1tfMigitationNotEnabled_LoadOption uint32 //col:2160
HvL1tfMigitationNotEnabled_CoreScheduler uint32 //col:2161
EnhancedIbrsReported uint32 //col:2162
MdsHardwareProtected uint32 //col:2163
MbClearEnabled uint32 //col:2164
MbClearReported uint32 //col:2165
Reserved uint32 //col:2166
}


type  _SYSTEM_DMA_GUARD_POLICY_INFORMATION struct{
DmaGuardPolicyEnabled bool //col:2140
}


type  _SYSTEM_ENCLAVE_LAUNCH_CONTROL_INFORMATION struct{
EnclaveLaunchSigner[32] uint8 //col:2144
}


type  _SYSTEM_WORKLOAD_ALLOWED_CPU_SET_INFORMATION struct{
WorkloadClass ULONGLONG //col:2149
CpuSets[1] ULONGLONG //col:2150
}


type  _SYSTEM_SECURITY_MODEL_INFORMATION struct{
Union union //col:2160
SecurityModelFlags uint32 //col:2162
Struct struct //col:2163
SModeAdminlessEnabled uint32 //col:2165
AllowDeviceOwnerProtectionDowngrade uint32 //col:2166
Reserved uint32 //col:2167
}


type  _SYSTEM_FEATURE_CONFIGURATION_INFORMATION struct{
ChangeStamp ULONGLONG //col:2167
_RTL_FEATURE_CONFIGURATION* struct //col:2168
}


type  _SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY struct{
ChangeStamp ULONGLONG //col:2173
Section uintptr //col:2174
Size ULONGLONG //col:2175
}


type  _SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION struct{
OverallChangeStamp ULONGLONG //col:2178
Descriptors[3] SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY //col:2179
}


type  _RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET struct{
Data[2] uint32 //col:2182
}


type  _SYSTEM_FEATURE_USAGE_SUBSCRIPTION_DETAILS struct{
FeatureId uint32 //col:2189
ReportingKind uint16 //col:2190
ReportingOptions uint16 //col:2191
ReportingTarget RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET //col:2192
}


type  _SYSTEM_FIRMWARE_RAMDISK_INFORMATION struct{
Version uint32 //col:2196
BlockSize uint32 //col:2197
BaseAddress ULONG_PTR //col:2198
Size int64 //col:2199
}


type  _SYSTEM_SHADOW_STACK_INFORMATION struct{
Union union //col:2211
Flags uint32 //col:2213
Struct struct //col:2214
CetCapable uint32 //col:2216
UserCetAllowed uint32 //col:2217
ReservedForUserCet uint32 //col:2218
KernelCetEnabled uint32 //col:2219
KernelCetAuditModeEnabled uint32 //col:2220
ReservedForKernelCet uint32 //col:2221
Reserved uint32 //col:2222
}


type  _SYSTEM_BUILD_VERSION_INFORMATION struct{
LayerNumber uint16 //col:2229
LayerCount uint16 //col:2230
OsMajorVersion uint32 //col:2231
OsMinorVersion uint32 //col:2232
NtBuildNumber uint32 //col:2233
NtBuildQfe uint32 //col:2234
LayerName[128] uint8 //col:2235
NtBuildBranch[128] uint8 //col:2236
NtBuildLab[128] uint8 //col:2237
NtBuildLabEx[128] uint8 //col:2238
NtBuildStamp[26] uint8 //col:2239
NtBuildArch[16] uint8 //col:2240
Flags SYSTEM_BUILD_VERSION_INFORMATION_FLAGS //col:2241
}


type  _SYSTEM_POOL_LIMIT_MEM_INFO struct{
MemoryLimit ULONGLONG //col:2234
NotificationLimit ULONGLONG //col:2235
}


type  _SYSTEM_POOL_LIMIT_INFO struct{
PoolTag uint32 //col:2240
MemLimits[2] SYSTEM_POOL_LIMIT_MEM_INFO //col:2241
NotificationHandle WNF_STATE_NAME //col:2242
}


type  _SYSTEM_POOL_LIMIT_INFORMATION struct{
Version uint32 //col:2246
EntryCount uint32 //col:2247
LimitEntries[1] SYSTEM_POOL_LIMIT_INFO //col:2248
}


type  _HV_MINROOT_NUMA_LPS struct{
NodeIndex uint32 //col:2251
Mask[16] ULONG_PTR //col:2252
}


type  _SYSTEM_IOMMU_STATE_INFORMATION struct{
State SYSTEM_IOMMU_STATE //col:2256
Pdo uintptr //col:2257
}


type  _SYSTEM_HYPERVISOR_MINROOT_INFORMATION struct{
NumProc uint32 //col:2267
RootProc uint32 //col:2268
RootProcNumaNodesSpecified uint32 //col:2269
RootProcNumaNodes[64] uint16 //col:2270
RootProcPerCore uint32 //col:2271
RootProcPerNode uint32 //col:2272
RootProcNumaNodesLpsSpecified uint32 //col:2273
RootProcNumaNodeLps[64] HV_MINROOT_NUMA_LPS //col:2274
}


type  _SYSTEM_HYPERVISOR_BOOT_PAGES_INFORMATION struct{
RangeCount uint32 //col:2272
RangeArray[1] ULONG_PTR //col:2273
}


type  _SYSTEM_POINTER_AUTH_INFORMATION struct{
Union union //col:2285
SupportedFlags uint16 //col:2287
Struct struct //col:2288
AddressAuthSupported uint16 //col:2290
AddressAuthQarma uint16 //col:2291
GenericAuthSupported uint16 //col:2292
GenericAuthQarma uint16 //col:2293
SupportedReserved uint16 //col:2294
}


type  _SYSDBG_VIRTUAL struct{
Address uintptr //col:2305
Buffer uintptr //col:2306
Request uint32 //col:2307
}


type  _SYSDBG_PHYSICAL struct{
Address PHYSICAL_ADDRESS //col:2311
Buffer uintptr //col:2312
Request uint32 //col:2313
}


type  _SYSDBG_CONTROL_SPACE struct{
Address ULONG64 //col:2318
Buffer uintptr //col:2319
Request uint32 //col:2320
Processor uint32 //col:2321
}


type  _SYSDBG_IO_SPACE struct{
Address ULONG64 //col:2327
Buffer uintptr //col:2328
Request uint32 //col:2329
_INTERFACE_TYPE enum //col:2330
BusNumber uint32 //col:2331
AddressSpace uint32 //col:2332
}


type  _SYSDBG_MSR struct{
Msr uint32 //col:2332
Data ULONG64 //col:2333
}


type  _SYSDBG_BUS_DATA struct{
Address uint32 //col:2341
Buffer uintptr //col:2342
Request uint32 //col:2343
_BUS_DATA_TYPE enum //col:2344
BusNumber uint32 //col:2345
SlotNumber uint32 //col:2346
}


type  _SYSDBG_TRIAGE_DUMP struct{
Flags uint32 //col:2353
BugCheckCode uint32 //col:2354
BugCheckParam1 ULONG_PTR //col:2355
BugCheckParam2 ULONG_PTR //col:2356
BugCheckParam3 ULONG_PTR //col:2357
BugCheckParam4 ULONG_PTR //col:2358
ProcessHandles uint32 //col:2359
ThreadHandles uint32 //col:2360
Handles PHANDLE //col:2361
}


type  _SYSDBG_LIVEDUMP_SELECTIVE_CONTROL struct{
Version uint32 //col:2365
Size uint32 //col:2366
Union union //col:2367
Flags ULONGLONG //col:2369
Struct struct //col:2370
ThreadKernelStacks ULONGLONG //col:2372
ReservedFlags ULONGLONG //col:2373
}


type  _SYSDBG_LIVEDUMP_CONTROL struct{
Version uint32 //col:2382
BugCheckCode uint32 //col:2383
BugCheckParam1 ULONG_PTR //col:2384
BugCheckParam2 ULONG_PTR //col:2385
BugCheckParam3 ULONG_PTR //col:2386
BugCheckParam4 ULONG_PTR //col:2387
DumpFileHandle uintptr //col:2388
CancelEventHandle uintptr //col:2389
Flags SYSDBG_LIVEDUMP_CONTROL_FLAGS //col:2390
AddPagesControl SYSDBG_LIVEDUMP_CONTROL_ADDPAGES //col:2391
SelectiveControl PSYSDBG_LIVEDUMP_SELECTIVE_CONTROL //col:2392
}


type  _SYSDBG_KD_PULL_REMOTE_FILE struct{
ImageFileName *int16 //col:2386
}


type  _KUSER_SHARED_DATA struct{
TickCountLowDeprecated uint32 //col:2431
TickCountMultiplier uint32 //col:2432
InterruptTime KSYSTEM_TIME //col:2433
SystemTime KSYSTEM_TIME //col:2434
TimeZoneBias KSYSTEM_TIME //col:2435
ImageNumberLow uint16 //col:2436
ImageNumberHigh uint16 //col:2437
NtSystemRoot[260] WCHAR //col:2438
MaxStackTraceDepth uint32 //col:2439
CryptoExponent uint32 //col:2440
TimeZoneId uint32 //col:2441
LargePageMinimum uint32 //col:2442
AitSamplingValue uint32 //col:2443
AppCompatFlag uint32 //col:2444
RNGSeedVersion ULONGLONG //col:2445
GlobalValidationRunlevel uint32 //col:2446
TimeZoneBiasStamp int32 //col:2447
NtBuildNumber uint32 //col:2448
NtProductType NT_PRODUCT_TYPE //col:2449
ProductTypeIsValid bool //col:2450
Reserved0[1] uint8 //col:2451
NativeProcessorArchitecture uint16 //col:2452
NtMajorVersion uint32 //col:2453
NtMinorVersion uint32 //col:2454
ProcessorFeatures[PROCESSOR_FEATURE_MAX] bool //col:2455
Reserved1 uint32 //col:2456
Reserved3 uint32 //col:2457
TimeSlip uint32 //col:2458
AlternativeArchitecture ALTERNATIVE_ARCHITECTURE_TYPE //col:2459
BootId uint32 //col:2460
SystemExpirationDate LARGE_INTEGER //col:2461
SuiteMask uint32 //col:2462
KdDebuggerEnabled bool //col:2463
Union union //col:2464
MitigationPolicies uint8 //col:2466
Struct struct //col:2467
NXSupportPolicy uint8 //col:2469
SEHValidationPolicy uint8 //col:2470
CurDirDevicesSkippedForDlls uint8 //col:2471
Reserved uint8 //col:2472
}


type  _ATOM_BASIC_INFORMATION struct{
UsageCount uint16 //col:2530
Flags uint16 //col:2531
NameLength uint16 //col:2532
Name[1] WCHAR //col:2533
}


type  _ATOM_TABLE_INFORMATION struct{
NumberOfAtoms uint32 //col:2535
Atoms[1] RTL_ATOM //col:2536
}




