package phnt


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



type BOOT_ENTRY struct{
Version uint32 //col:3
Length uint32 //col:4
Id uint32 //col:5
Attributes uint32 //col:6
FriendlyNameOffset uint32 //col:7
BootFilePathOffset uint32 //col:8
OsOptionsLength uint32 //col:9
OsOptions[1] uint8 //col:10
}



type BOOT_ENTRY_LIST struct{
NextEntryOffset uint32 //col:14
BootEntry BOOT_ENTRY //col:15
}



type BOOT_OPTIONS struct{
Version uint32 //col:19
Length uint32 //col:20
Timeout uint32 //col:21
CurrentBootEntryId uint32 //col:22
NextBootEntryId uint32 //col:23
HeadlessRedirection[1] WCHAR //col:24
}



type FILE_PATH struct{
Version uint32 //col:28
Length uint32 //col:29
Type uint32 //col:30
FilePath[1] uint8 //col:31
}



type EFI_DRIVER_ENTRY struct{
Version uint32 //col:35
Length uint32 //col:36
Id uint32 //col:37
FriendlyNameOffset uint32 //col:38
DriverFilePathOffset uint32 //col:39
}



type EFI_DRIVER_ENTRY_LIST struct{
NextEntryOffset uint32 //col:43
DriverEntry EFI_DRIVER_ENTRY //col:44
}



type EVENT_BASIC_INFORMATION struct{
EventType EVENT_TYPE //col:48
EventState LONG //col:49
}



type MUTANT_BASIC_INFORMATION struct{
CurrentCount LONG //col:53
OwnedByCaller bool //col:54
AbandonedState bool //col:55
}



type MUTANT_OWNER_INFORMATION struct{
ClientId CLIENT_ID //col:59
}



type SEMAPHORE_BASIC_INFORMATION struct{
CurrentCount LONG //col:63
MaximumCount LONG //col:64
}



type TIMER_BASIC_INFORMATION struct{
RemainingTime LARGE_INTEGER //col:68
TimerState bool //col:69
}



type TIMER_SET_COALESCABLE_TIMER_INFO struct{
LARGE_INTEGER _In_ //col:73
PTIMER_APC_ROUTINE _In_opt_ //col:74
PVOID _In_opt_ //col:75
_In_opt_ // //col:76
ULONG _In_opt_ //col:77
ULONG _In_ //col:78
PBOOLEAN _Out_opt_ //col:79
}



type T2_SET_PARAMETERS_V0 struct{
Version uint32 //col:83
Reserved uint32 //col:84
NoWakeTolerance LONGLONG //col:85
}



type WNF_STATE_NAME struct{
Data[2] uint32 //col:89
}



type WNF_TYPE_ID struct{
TypeId GUID //col:93
}



type WNF_DELIVERY_DESCRIPTOR struct{
SubscriptionId ULONGLONG //col:97
StateName WNF_STATE_NAME //col:98
ChangeStamp WNF_CHANGE_STAMP //col:99
StateDataSize uint32 //col:100
EventMask uint32 //col:101
TypeId WNF_TYPE_ID //col:102
StateDataOffset uint32 //col:103
}



type WORKER_FACTORY_BASIC_INFORMATION struct{
Timeout LARGE_INTEGER //col:107
RetryTimeout LARGE_INTEGER //col:108
IdleTimeout LARGE_INTEGER //col:109
Paused bool //col:110
TimerSet bool //col:111
QueuedToExWorker bool //col:112
MayCreate bool //col:113
CreateInProgress bool //col:114
InsertedIntoQueue bool //col:115
Shutdown bool //col:116
BindingCount uint32 //col:117
ThreadMinimum uint32 //col:118
ThreadMaximum uint32 //col:119
PendingWorkerCount uint32 //col:120
WaitingWorkerCount uint32 //col:121
TotalWorkerCount uint32 //col:122
ReleaseCount uint32 //col:123
InfiniteWaitGoal LONGLONG //col:124
StartRoutine PVOID //col:125
StartParameter PVOID //col:126
ProcessId HANDLE //col:127
StackReserve SIZE_T //col:128
StackCommit SIZE_T //col:129
LastThreadCreationStatus NTSTATUS //col:130
}



type WORKER_FACTORY_DEFERRED_WORK struct{
struct // //col:134
AlpcSendMessagePort PVOID //col:135
AlpcSendMessageFlags uint32 //col:136
Flags uint32 //col:137
}



type SYSTEM_BASIC_INFORMATION struct{
Reserved uint32 //col:141
TimerResolution uint32 //col:142
PageSize uint32 //col:143
NumberOfPhysicalPages uint32 //col:144
LowestPhysicalPageNumber uint32 //col:145
HighestPhysicalPageNumber uint32 //col:146
AllocationGranularity uint32 //col:147
MinimumUserModeAddress ULONG_PTR //col:148
MaximumUserModeAddress ULONG_PTR //col:149
ActiveProcessorsAffinityMask KAFFINITY //col:150
NumberOfProcessors CCHAR //col:151
}



type SYSTEM_PROCESSOR_INFORMATION struct{
ProcessorArchitecture USHORT //col:155
ProcessorLevel USHORT //col:156
ProcessorRevision USHORT //col:157
MaximumProcessors USHORT //col:158
ProcessorFeatureBits uint32 //col:159
}



type SYSTEM_PERFORMANCE_INFORMATION struct{
IdleProcessTime LARGE_INTEGER //col:163
IoReadTransferCount LARGE_INTEGER //col:164
IoWriteTransferCount LARGE_INTEGER //col:165
IoOtherTransferCount LARGE_INTEGER //col:166
IoReadOperationCount uint32 //col:167
IoWriteOperationCount uint32 //col:168
IoOtherOperationCount uint32 //col:169
AvailablePages uint32 //col:170
CommittedPages uint32 //col:171
CommitLimit uint32 //col:172
PeakCommitment uint32 //col:173
PageFaultCount uint32 //col:174
CopyOnWriteCount uint32 //col:175
TransitionCount uint32 //col:176
CacheTransitionCount uint32 //col:177
DemandZeroCount uint32 //col:178
PageReadCount uint32 //col:179
PageReadIoCount uint32 //col:180
CacheReadCount uint32 //col:181
CacheIoCount uint32 //col:182
DirtyPagesWriteCount uint32 //col:183
DirtyWriteIoCount uint32 //col:184
MappedPagesWriteCount uint32 //col:185
MappedWriteIoCount uint32 //col:186
PagedPoolPages uint32 //col:187
NonPagedPoolPages uint32 //col:188
PagedPoolAllocs uint32 //col:189
PagedPoolFrees uint32 //col:190
NonPagedPoolAllocs uint32 //col:191
NonPagedPoolFrees uint32 //col:192
FreeSystemPtes uint32 //col:193
ResidentSystemCodePage uint32 //col:194
TotalSystemDriverPages uint32 //col:195
TotalSystemCodePages uint32 //col:196
NonPagedPoolLookasideHits uint32 //col:197
PagedPoolLookasideHits uint32 //col:198
AvailablePagedPoolPages uint32 //col:199
ResidentSystemCachePage uint32 //col:200
ResidentPagedPoolPage uint32 //col:201
ResidentSystemDriverPage uint32 //col:202
CcFastReadNoWait uint32 //col:203
CcFastReadWait uint32 //col:204
CcFastReadResourceMiss uint32 //col:205
CcFastReadNotPossible uint32 //col:206
CcFastMdlReadNoWait uint32 //col:207
CcFastMdlReadWait uint32 //col:208
CcFastMdlReadResourceMiss uint32 //col:209
CcFastMdlReadNotPossible uint32 //col:210
CcMapDataNoWait uint32 //col:211
CcMapDataWait uint32 //col:212
CcMapDataNoWaitMiss uint32 //col:213
CcMapDataWaitMiss uint32 //col:214
CcPinMappedDataCount uint32 //col:215
CcPinReadNoWait uint32 //col:216
CcPinReadWait uint32 //col:217
CcPinReadNoWaitMiss uint32 //col:218
CcPinReadWaitMiss uint32 //col:219
CcCopyReadNoWait uint32 //col:220
CcCopyReadWait uint32 //col:221
CcCopyReadNoWaitMiss uint32 //col:222
CcCopyReadWaitMiss uint32 //col:223
CcMdlReadNoWait uint32 //col:224
CcMdlReadWait uint32 //col:225
CcMdlReadNoWaitMiss uint32 //col:226
CcMdlReadWaitMiss uint32 //col:227
CcReadAheadIos uint32 //col:228
CcLazyWriteIos uint32 //col:229
CcLazyWritePages uint32 //col:230
CcDataFlushes uint32 //col:231
CcDataPages uint32 //col:232
ContextSwitches uint32 //col:233
FirstLevelTbFills uint32 //col:234
SecondLevelTbFills uint32 //col:235
SystemCalls uint32 //col:236
CcTotalDirtyPages ULONGLONG //col:237
CcDirtyPageThreshold ULONGLONG //col:238
ResidentAvailablePages LONGLONG //col:239
SharedCommittedPages ULONGLONG //col:240
}



type SYSTEM_TIMEOFDAY_INFORMATION struct{
BootTime LARGE_INTEGER //col:244
CurrentTime LARGE_INTEGER //col:245
TimeZoneBias LARGE_INTEGER //col:246
TimeZoneId uint32 //col:247
Reserved uint32 //col:248
BootTimeBias ULONGLONG //col:249
SleepTimeBias ULONGLONG //col:250
}



type SYSTEM_THREAD_INFORMATION struct{
KernelTime LARGE_INTEGER //col:254
UserTime LARGE_INTEGER //col:255
CreateTime LARGE_INTEGER //col:256
WaitTime uint32 //col:257
StartAddress PVOID //col:258
ClientId CLIENT_ID //col:259
Priority KPRIORITY //col:260
BasePriority KPRIORITY //col:261
ContextSwitches uint32 //col:262
ThreadState KTHREAD_STATE //col:263
WaitReason KWAIT_REASON //col:264
}



type  struct{
struct //typedef //col:266
ThreadInfo SYSTEM_THREAD_INFORMATION //col:269
StackBase PVOID //col:270
StackLimit PVOID //col:271
Win32StartAddress PVOID //col:272
TebBase PTEB //col:273
Reserved2 ULONG_PTR //col:274
Reserved3 ULONG_PTR //col:275
Reserved4 ULONG_PTR //col:276
}



type SYSTEM_EXTENDED_THREAD_INFORMATION struct{
ThreadInfo SYSTEM_THREAD_INFORMATION //col:269
StackBase PVOID //col:270
StackLimit PVOID //col:271
Win32StartAddress PVOID //col:272
TebBase PTEB //col:273
Reserved2 ULONG_PTR //col:274
Reserved3 ULONG_PTR //col:275
Reserved4 ULONG_PTR //col:276
}



type SYSTEM_PROCESS_INFORMATION struct{
NextEntryOffset uint32 //col:280
NumberOfThreads uint32 //col:281
WorkingSetPrivateSize LARGE_INTEGER //col:282
HardFaultCount uint32 //col:283
NumberOfThreadsHighWatermark uint32 //col:284
CycleTime ULONGLONG //col:285
CreateTime LARGE_INTEGER //col:286
UserTime LARGE_INTEGER //col:287
KernelTime LARGE_INTEGER //col:288
ImageName UNICODE_STRING //col:289
BasePriority KPRIORITY //col:290
UniqueProcessId HANDLE //col:291
InheritedFromUniqueProcessId HANDLE //col:292
HandleCount uint32 //col:293
SessionId uint32 //col:294
UniqueProcessKey ULONG_PTR //col:295
PeakVirtualSize SIZE_T //col:296
VirtualSize SIZE_T //col:297
PageFaultCount uint32 //col:298
PeakWorkingSetSize SIZE_T //col:299
WorkingSetSize SIZE_T //col:300
QuotaPeakPagedPoolUsage SIZE_T //col:301
QuotaPagedPoolUsage SIZE_T //col:302
QuotaPeakNonPagedPoolUsage SIZE_T //col:303
QuotaNonPagedPoolUsage SIZE_T //col:304
PagefileUsage SIZE_T //col:305
PeakPagefileUsage SIZE_T //col:306
PrivatePageCount SIZE_T //col:307
ReadOperationCount LARGE_INTEGER //col:308
WriteOperationCount LARGE_INTEGER //col:309
OtherOperationCount LARGE_INTEGER //col:310
ReadTransferCount LARGE_INTEGER //col:311
WriteTransferCount LARGE_INTEGER //col:312
OtherTransferCount LARGE_INTEGER //col:313
Threads[1] SYSTEM_THREAD_INFORMATION //col:314
}



type SYSTEM_CALL_COUNT_INFORMATION struct{
Length uint32 //col:318
NumberOfTables uint32 //col:319
}



type SYSTEM_DEVICE_INFORMATION struct{
NumberOfDisks uint32 //col:323
NumberOfFloppies uint32 //col:324
NumberOfCdRoms uint32 //col:325
NumberOfTapes uint32 //col:326
NumberOfSerialPorts uint32 //col:327
NumberOfParallelPorts uint32 //col:328
}



type SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION struct{
IdleTime LARGE_INTEGER //col:332
KernelTime LARGE_INTEGER //col:333
UserTime LARGE_INTEGER //col:334
DpcTime LARGE_INTEGER //col:335
InterruptTime LARGE_INTEGER //col:336
InterruptCount uint32 //col:337
}



type SYSTEM_FLAGS_INFORMATION struct{
Flags uint32 //col:341
}



type SYSTEM_CALL_TIME_INFORMATION struct{
Length uint32 //col:345
TotalCalls uint32 //col:346
TimeOfCalls[1] LARGE_INTEGER //col:347
}



type RTL_PROCESS_LOCK_INFORMATION struct{
Address PVOID //col:351
Type USHORT //col:352
CreatorBackTraceIndex USHORT //col:353
OwningThread HANDLE //col:354
LockCount LONG //col:355
ContentionCount uint32 //col:356
EntryCount uint32 //col:357
RecursionCount LONG //col:358
NumberOfWaitingShared uint32 //col:359
NumberOfWaitingExclusive uint32 //col:360
}



type RTL_PROCESS_LOCKS struct{
NumberOfLocks uint32 //col:364
Locks[1] RTL_PROCESS_LOCK_INFORMATION //col:365
}



type RTL_PROCESS_BACKTRACE_INFORMATION struct{
SymbolicBackTrace PCHAR //col:369
TraceCount uint32 //col:370
Index USHORT //col:371
Depth USHORT //col:372
BackTrace[32] PVOID //col:373
}



type RTL_PROCESS_BACKTRACES struct{
CommittedMemory uint32 //col:377
ReservedMemory uint32 //col:378
NumberOfBackTraceLookups uint32 //col:379
NumberOfBackTraces uint32 //col:380
BackTraces[1] RTL_PROCESS_BACKTRACE_INFORMATION //col:381
}



type SYSTEM_HANDLE_TABLE_ENTRY_INFO struct{
UniqueProcessId USHORT //col:385
CreatorBackTraceIndex USHORT //col:386
ObjectTypeIndex uint8 //col:387
HandleAttributes uint8 //col:388
HandleValue USHORT //col:389
Object PVOID //col:390
GrantedAccess uint32 //col:391
}



type SYSTEM_HANDLE_INFORMATION struct{
NumberOfHandles uint32 //col:395
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO //col:396
}



type SYSTEM_OBJECTTYPE_INFORMATION struct{
NextEntryOffset uint32 //col:400
NumberOfObjects uint32 //col:401
NumberOfHandles uint32 //col:402
TypeIndex uint32 //col:403
InvalidAttributes uint32 //col:404
GenericMapping GENERIC_MAPPING //col:405
ValidAccessMask uint32 //col:406
PoolType uint32 //col:407
SecurityRequired bool //col:408
WaitableObject bool //col:409
TypeName UNICODE_STRING //col:410
}



type SYSTEM_OBJECT_INFORMATION struct{
NextEntryOffset uint32 //col:414
Object PVOID //col:415
CreatorUniqueProcess HANDLE //col:416
CreatorBackTraceIndex USHORT //col:417
Flags USHORT //col:418
PointerCount LONG //col:419
HandleCount LONG //col:420
PagedPoolCharge uint32 //col:421
NonPagedPoolCharge uint32 //col:422
ExclusiveProcessId HANDLE //col:423
SecurityDescriptor PVOID //col:424
NameInfo UNICODE_STRING //col:425
}



type SYSTEM_PAGEFILE_INFORMATION struct{
NextEntryOffset uint32 //col:429
TotalSize uint32 //col:430
TotalInUse uint32 //col:431
PeakUsage uint32 //col:432
PageFileName UNICODE_STRING //col:433
}



type SYSTEM_VDM_INSTEMUL_INFO struct{
SegmentNotPresent uint32 //col:437
VdmOpcode0F uint32 //col:438
OpcodeESPrefix uint32 //col:439
OpcodeCSPrefix uint32 //col:440
OpcodeSSPrefix uint32 //col:441
OpcodeDSPrefix uint32 //col:442
OpcodeFSPrefix uint32 //col:443
OpcodeGSPrefix uint32 //col:444
OpcodeOPER32Prefix uint32 //col:445
OpcodeADDR32Prefix uint32 //col:446
OpcodeINSB uint32 //col:447
OpcodeINSW uint32 //col:448
OpcodeOUTSB uint32 //col:449
OpcodeOUTSW uint32 //col:450
OpcodePUSHF uint32 //col:451
OpcodePOPF uint32 //col:452
OpcodeINTnn uint32 //col:453
OpcodeINTO uint32 //col:454
OpcodeIRET uint32 //col:455
OpcodeINBimm uint32 //col:456
OpcodeINWimm uint32 //col:457
OpcodeOUTBimm uint32 //col:458
OpcodeOUTWimm uint32 //col:459
OpcodeINB uint32 //col:460
OpcodeINW uint32 //col:461
OpcodeOUTB uint32 //col:462
OpcodeOUTW uint32 //col:463
OpcodeLOCKPrefix uint32 //col:464
OpcodeREPNEPrefix uint32 //col:465
OpcodeREPPrefix uint32 //col:466
OpcodeHLT uint32 //col:467
OpcodeCLI uint32 //col:468
OpcodeSTI uint32 //col:469
BopCount uint32 //col:470
}



type SYSTEM_FILECACHE_INFORMATION struct{
CurrentSize SIZE_T //col:474
PeakSize SIZE_T //col:475
PageFaultCount uint32 //col:476
MinimumWorkingSet SIZE_T //col:477
MaximumWorkingSet SIZE_T //col:478
CurrentSizeIncludingTransitionInPages SIZE_T //col:479
PeakSizeIncludingTransitionInPages SIZE_T //col:480
TransitionRePurposeCount uint32 //col:481
Flags uint32 //col:482
}



type SYSTEM_BASIC_WORKING_SET_INFORMATION struct{
CurrentSize SIZE_T //col:486
PeakSize SIZE_T //col:487
PageFaultCount uint32 //col:488
}



type SYSTEM_POOLTAG struct{
Union union //col:492
Tag[4] uint8 //col:494
TagUlong uint32 //col:495
}



type SYSTEM_POOLTAG_INFORMATION struct{
Count uint32 //col:506
TagInfo[1] SYSTEM_POOLTAG //col:507
}



type SYSTEM_INTERRUPT_INFORMATION struct{
ContextSwitches uint32 //col:511
DpcCount uint32 //col:512
DpcRate uint32 //col:513
TimeIncrement uint32 //col:514
DpcBypassCount uint32 //col:515
ApcBypassCount uint32 //col:516
}



type SYSTEM_DPC_BEHAVIOR_INFORMATION struct{
Spare uint32 //col:520
DpcQueueDepth uint32 //col:521
MinimumDpcRate uint32 //col:522
AdjustDpcThreshold uint32 //col:523
IdealDpcRate uint32 //col:524
}



type SYSTEM_QUERY_TIME_ADJUST_INFORMATION struct{
TimeAdjustment uint32 //col:528
TimeIncrement uint32 //col:529
Enable bool //col:530
}



type SYSTEM_QUERY_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG //col:534
TimeIncrement ULONGLONG //col:535
Enable bool //col:536
}



type SYSTEM_SET_TIME_ADJUST_INFORMATION struct{
TimeAdjustment uint32 //col:540
Enable bool //col:541
}



type SYSTEM_SET_TIME_ADJUST_INFORMATION_PRECISE struct{
TimeAdjustment ULONGLONG //col:545
Enable bool //col:546
}



type EVENT_TRACE_VERSION_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:550
EventTraceKernelVersion uint32 //col:551
}



type PERFINFO_GROUPMASK struct{
Masks[PERF_NUM_MASKS] uint32 //col:555
}



type EVENT_TRACE_GROUPMASK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:559
TraceHandle TRACEHANDLE //col:560
EventTraceGroupMasks PERFINFO_GROUPMASK //col:561
}



type EVENT_TRACE_PERFORMANCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:565
LogfileBytesWritten LARGE_INTEGER //col:566
}



type EVENT_TRACE_TIME_PROFILE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:570
ProfileInterval uint32 //col:571
}



type EVENT_TRACE_SESSION_SECURITY_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:575
SecurityInformation uint32 //col:576
TraceHandle TRACEHANDLE //col:577
SecurityDescriptor[1] uint8 //col:578
}



type EVENT_TRACE_SPINLOCK_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:582
SpinLockSpinThreshold uint32 //col:583
SpinLockAcquireSampleRate uint32 //col:584
SpinLockContentionSampleRate uint32 //col:585
SpinLockHoldThreshold uint32 //col:586
}



type EVENT_TRACE_SYSTEM_EVENT_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:590
TraceHandle TRACEHANDLE //col:591
HookId[1] uint32 //col:592
}



type EVENT_TRACE_EXECUTIVE_RESOURCE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:596
ReleaseSamplingRate uint32 //col:597
ContentionSamplingRate uint32 //col:598
NumberOfExcessiveTimeouts uint32 //col:599
}



type EVENT_TRACE_HEAP_TRACING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:603
ProcessId uint32 //col:604
}



type EVENT_TRACE_TAG_FILTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:608
TraceHandle TRACEHANDLE //col:609
Filter[1] uint32 //col:610
}



type EVENT_TRACE_PROFILE_COUNTER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:614
TraceHandle TRACEHANDLE //col:615
ProfileSource[1] uint32 //col:616
}



type EVENT_TRACE_PROFILE_LIST_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:620
Spare uint32 //col:621
struct // //col:622
}



type EVENT_TRACE_STACK_CACHING_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:626
TraceHandle TRACEHANDLE //col:627
Enabled bool //col:628
Reserved[3] uint8 //col:629
CacheSize uint32 //col:630
BucketCount uint32 //col:631
}



type EVENT_TRACE_SOFT_RESTART_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:635
TraceHandle TRACEHANDLE //col:636
PersistTraceBuffers bool //col:637
FileName[1] WCHAR //col:638
}



type EVENT_TRACE_PROFILE_ADD_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:642
PerfEvtEventSelect bool //col:643
PerfEvtUnitSelect bool //col:644
PerfEvtType uint32 //col:645
CpuInfoHierarchy[0x3] uint32 //col:646
InitialInterval uint32 //col:647
AllowsHalt bool //col:648
Persist bool //col:649
ProfileSourceDescription[0x1] WCHAR //col:650
}



type EVENT_TRACE_PROFILE_REMOVE_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:654
ProfileSource KPROFILE_SOURCE //col:655
CpuInfoHierarchy[0x3] uint32 //col:656
}



type EVENT_TRACE_COVERAGE_SAMPLER_INFORMATION struct{
EventTraceInformationClass EVENT_TRACE_INFORMATION_CLASS //col:660
CoverageSamplerInformationClass bool //col:661
MajorVersion uint8 //col:662
MinorVersion uint8 //col:663
Reserved uint8 //col:664
SamplerHandle HANDLE //col:665
}



type SYSTEM_EXCEPTION_INFORMATION struct{
AlignmentFixupCount uint32 //col:669
ExceptionDispatchCount uint32 //col:670
FloatingEmulationCount uint32 //col:671
ByteWordEmulationCount uint32 //col:672
}



type SYSTEM_CRASH_DUMP_STATE_INFORMATION struct{
CrashDumpConfigurationClass SYSTEM_CRASH_DUMP_CONFIGURATION_CLASS //col:676
}



type SYSTEM_KERNEL_DEBUGGER_INFORMATION struct{
KernelDebuggerEnabled bool //col:680
KernelDebuggerNotPresent bool //col:681
}



type SYSTEM_CONTEXT_SWITCH_INFORMATION struct{
ContextSwitches uint32 //col:685
FindAny uint32 //col:686
FindLast uint32 //col:687
FindIdeal uint32 //col:688
IdleAny uint32 //col:689
IdleCurrent uint32 //col:690
IdleLast uint32 //col:691
IdleIdeal uint32 //col:692
PreemptAny uint32 //col:693
PreemptCurrent uint32 //col:694
PreemptLast uint32 //col:695
SwitchToIdle uint32 //col:696
}



type SYSTEM_REGISTRY_QUOTA_INFORMATION struct{
RegistryQuotaAllowed uint32 //col:700
RegistryQuotaUsed uint32 //col:701
PagedPoolSize SIZE_T //col:702
}



type SYSTEM_PROCESSOR_IDLE_INFORMATION struct{
IdleTime ULONGLONG //col:706
C1Time ULONGLONG //col:707
C2Time ULONGLONG //col:708
C3Time ULONGLONG //col:709
C1Transitions uint32 //col:710
C2Transitions uint32 //col:711
C3Transitions uint32 //col:712
Padding uint32 //col:713
}



type SYSTEM_LEGACY_DRIVER_INFORMATION struct{
VetoType uint32 //col:717
VetoList UNICODE_STRING //col:718
}



type SYSTEM_LOOKASIDE_INFORMATION struct{
CurrentDepth USHORT //col:722
MaximumDepth USHORT //col:723
TotalAllocates uint32 //col:724
AllocateMisses uint32 //col:725
TotalFrees uint32 //col:726
FreeMisses uint32 //col:727
Type uint32 //col:728
Tag uint32 //col:729
Size uint32 //col:730
}



type SYSTEM_RANGE_START_INFORMATION struct{
SystemRangeStart ULONG_PTR //col:734
}



type SYSTEM_VERIFIER_INFORMATION_LEGACY  struct{
NextEntryOffset uint32 //col:738
Level uint32 //col:739
DriverName UNICODE_STRING //col:740
RaiseIrqls uint32 //col:741
AcquireSpinLocks uint32 //col:742
SynchronizeExecutions uint32 //col:743
AllocationsAttempted uint32 //col:744
AllocationsSucceeded uint32 //col:745
AllocationsSucceededSpecialPool uint32 //col:746
AllocationsWithNoTag uint32 //col:747
TrimRequests uint32 //col:748
Trims uint32 //col:749
AllocationsFailed uint32 //col:750
AllocationsFailedDeliberately uint32 //col:751
Loads uint32 //col:752
Unloads uint32 //col:753
UnTrackedPool uint32 //col:754
CurrentPagedPoolAllocations uint32 //col:755
CurrentNonPagedPoolAllocations uint32 //col:756
PeakPagedPoolAllocations uint32 //col:757
PeakNonPagedPoolAllocations uint32 //col:758
PagedPoolUsageInBytes SIZE_T //col:759
NonPagedPoolUsageInBytes SIZE_T //col:760
PeakPagedPoolUsageInBytes SIZE_T //col:761
PeakNonPagedPoolUsageInBytes SIZE_T //col:762
}



type SYSTEM_VERIFIER_INFORMATION struct{
NextEntryOffset uint32 //col:766
Level uint32 //col:767
RuleClasses[2] uint32 //col:768
TriageContext uint32 //col:769
AreAllDriversBeingVerified uint32 //col:770
DriverName UNICODE_STRING //col:771
RaiseIrqls uint32 //col:772
AcquireSpinLocks uint32 //col:773
SynchronizeExecutions uint32 //col:774
AllocationsAttempted uint32 //col:775
AllocationsSucceeded uint32 //col:776
AllocationsSucceededSpecialPool uint32 //col:777
AllocationsWithNoTag uint32 //col:778
TrimRequests uint32 //col:779
Trims uint32 //col:780
AllocationsFailed uint32 //col:781
AllocationsFailedDeliberately uint32 //col:782
Loads uint32 //col:783
Unloads uint32 //col:784
UnTrackedPool uint32 //col:785
CurrentPagedPoolAllocations uint32 //col:786
CurrentNonPagedPoolAllocations uint32 //col:787
PeakPagedPoolAllocations uint32 //col:788
PeakNonPagedPoolAllocations uint32 //col:789
PagedPoolUsageInBytes SIZE_T //col:790
NonPagedPoolUsageInBytes SIZE_T //col:791
PeakPagedPoolUsageInBytes SIZE_T //col:792
PeakNonPagedPoolUsageInBytes SIZE_T //col:793
}



type SYSTEM_SESSION_PROCESS_INFORMATION struct{
SessionId uint32 //col:797
SizeOfBuf uint32 //col:798
Buffer PVOID //col:799
}



type SYSTEM_GDI_DRIVER_INFORMATION struct{
DriverName UNICODE_STRING //col:803
ImageAddress PVOID //col:804
SectionPointer PVOID //col:805
EntryPoint PVOID //col:806
struct // //col:807
ImageLength uint32 //col:808
}



type SYSTEM_NUMA_INFORMATION struct{
HighestNodeNumber uint32 //col:812
Reserved uint32 //col:813
Union union //col:814
ActiveProcessorsGroupAffinity[MAXIMUM_NODE_COUNT] GROUP_AFFINITY //col:816
AvailableMemory[MAXIMUM_NODE_COUNT] ULONGLONG //col:817
Pad[MAXIMUM_NODE_COUNT ULONGLONG //col:818
}



type SYSTEM_PROCESSOR_POWER_INFORMATION struct{
CurrentFrequency uint8 //col:823
ThermalLimitFrequency uint8 //col:824
ConstantThrottleFrequency uint8 //col:825
DegradedThrottleFrequency uint8 //col:826
LastBusyFrequency uint8 //col:827
LastC3Frequency uint8 //col:828
LastAdjustedBusyFrequency uint8 //col:829
ProcessorMinThrottle uint8 //col:830
ProcessorMaxThrottle uint8 //col:831
NumberOfFrequencies uint32 //col:832
PromotionCount uint32 //col:833
DemotionCount uint32 //col:834
ErrorCount uint32 //col:835
RetryCount uint32 //col:836
CurrentFrequencyTime ULONGLONG //col:837
CurrentProcessorTime ULONGLONG //col:838
CurrentProcessorIdleTime ULONGLONG //col:839
LastProcessorTime ULONGLONG //col:840
LastProcessorIdleTime ULONGLONG //col:841
Energy ULONGLONG //col:842
}



type SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX struct{
Object PVOID //col:846
UniqueProcessId ULONG_PTR //col:847
HandleValue ULONG_PTR //col:848
GrantedAccess uint32 //col:849
CreatorBackTraceIndex USHORT //col:850
ObjectTypeIndex USHORT //col:851
HandleAttributes uint32 //col:852
Reserved uint32 //col:853
}



type SYSTEM_HANDLE_INFORMATION_EX struct{
NumberOfHandles ULONG_PTR //col:857
Reserved ULONG_PTR //col:858
Handles[1] SYSTEM_HANDLE_TABLE_ENTRY_INFO_EX //col:859
}



type SYSTEM_BIGPOOL_ENTRY struct{
Union union //col:863
VirtualAddress PVOID //col:865
NonPaged ULONG_PTR //col:866
}



type SYSTEM_BIGPOOL_INFORMATION struct{
Count uint32 //col:877
AllocatedInfo[1] SYSTEM_BIGPOOL_ENTRY //col:878
}



type SYSTEM_POOL_ENTRY struct{
Allocated bool //col:882
Spare0 bool //col:883
AllocatorBackTraceIndex USHORT //col:884
Size uint32 //col:885
Union union //col:886
Tag[4] uint8 //col:888
TagUlong uint32 //col:889
ProcessChargedQuota PVOID //col:890
}



type SYSTEM_POOL_INFORMATION struct{
TotalSize SIZE_T //col:895
FirstEntry PVOID //col:896
EntryOverhead USHORT //col:897
PoolTagPresent bool //col:898
Spare0 bool //col:899
NumberOfEntries uint32 //col:900
Entries[1] SYSTEM_POOL_ENTRY //col:901
}



type SYSTEM_SESSION_POOLTAG_INFORMATION struct{
NextEntryOffset SIZE_T //col:905
SessionId uint32 //col:906
Count uint32 //col:907
TagInfo[1] SYSTEM_POOLTAG //col:908
}



type SYSTEM_SESSION_MAPPED_VIEW_INFORMATION struct{
NextEntryOffset SIZE_T //col:912
SessionId uint32 //col:913
ViewFailures uint32 //col:914
NumberOfBytesAvailable SIZE_T //col:915
NumberOfBytesAvailableContiguous SIZE_T //col:916
}



type SYSTEM_WATCHDOG_HANDLER_INFORMATION  struct{
WdHandler PSYSTEM_WATCHDOG_HANDLER //col:920
Context PVOID //col:921
}



type SYSTEM_WATCHDOG_TIMER_INFORMATION struct{
WdInfoClass WATCHDOG_INFORMATION_CLASS //col:925
DataValue uint32 //col:926
}



type SYSTEM_FIRMWARE_TABLE_INFORMATION struct{
ProviderSignature uint32 //col:930
Action SYSTEM_FIRMWARE_TABLE_ACTION //col:931
TableID uint32 //col:932
TableBufferLength uint32 //col:933
TableBuffer[1] uint8 //col:934
}



type SYSTEM_FIRMWARE_TABLE_HANDLER struct{
ProviderSignature uint32 //col:938
Register bool //col:939
FirmwareTableHandler PFNFTH //col:940
DriverObject PVOID //col:941
}



type SYSTEM_MEMORY_LIST_INFORMATION struct{
ZeroPageCount ULONG_PTR //col:945
FreePageCount ULONG_PTR //col:946
ModifiedPageCount ULONG_PTR //col:947
ModifiedNoWritePageCount ULONG_PTR //col:948
BadPageCount ULONG_PTR //col:949
PageCountByPriority[8] ULONG_PTR //col:950
RepurposedPagesByPriority[8] ULONG_PTR //col:951
ModifiedPageCountPageFile ULONG_PTR //col:952
}



type SYSTEM_THREAD_CID_PRIORITY_INFORMATION struct{
ClientId CLIENT_ID //col:956
Priority KPRIORITY //col:957
}



type SYSTEM_PROCESSOR_IDLE_CYCLE_TIME_INFORMATION struct{
CycleTime ULONGLONG //col:961
}



type SYSTEM_VERIFIER_ISSUE struct{
IssueType ULONGLONG //col:965
Address PVOID //col:966
Parameters[2] ULONGLONG //col:967
}



type SYSTEM_VERIFIER_CANCELLATION_INFORMATION struct{
CancelProbability uint32 //col:971
CancelThreshold uint32 //col:972
CompletionThreshold uint32 //col:973
CancellationVerifierDisabled uint32 //col:974
AvailableIssues uint32 //col:975
Issues[128] SYSTEM_VERIFIER_ISSUE //col:976
}



type SYSTEM_REF_TRACE_INFORMATION struct{
TraceEnable bool //col:980
TracePermanent bool //col:981
TraceProcessName UNICODE_STRING //col:982
TracePoolTags UNICODE_STRING //col:983
}



type SYSTEM_SPECIAL_POOL_INFORMATION struct{
PoolTag uint32 //col:987
Flags uint32 //col:988
}



type SYSTEM_PROCESS_ID_INFORMATION struct{
ProcessId HANDLE //col:992
ImageName UNICODE_STRING //col:993
}



type SYSTEM_HYPERVISOR_QUERY_INFORMATION struct{
HypervisorConnected bool //col:997
HypervisorDebuggingEnabled bool //col:998
HypervisorPresent bool //col:999
Spare0[5] bool //col:1000
EnabledEnlightenments ULONGLONG //col:1001
}



type SYSTEM_BOOT_ENVIRONMENT_INFORMATION struct{
BootIdentifier GUID //col:1005
FirmwareType FIRMWARE_TYPE //col:1006
Union union //col:1007
BootFlags ULONGLONG //col:1009
Struct struct //col:1010
DbgMenuOsSelection ULONGLONG //col:1012
DbgHiberBoot ULONGLONG //col:1013
DbgSoftBoot ULONGLONG //col:1014
DbgMeasuredLaunch ULONGLONG //col:1015
DbgMeasuredLaunchCapable ULONGLONG //col:1016
DbgSystemHiveReplace ULONGLONG //col:1017
DbgMeasuredLaunchSmmProtections ULONGLONG //col:1018
DbgMeasuredLaunchSmmLevel ULONGLONG //col:1019
}



type SYSTEM_IMAGE_FILE_EXECUTION_OPTIONS_INFORMATION struct{
FlagsToEnable uint32 //col:1025
FlagsToDisable uint32 //col:1026
}



type COVERAGE_MODULE_REQUEST struct{
RequestType COVERAGE_REQUEST_CODES //col:1030
Union union //col:1031
MD5Hash[16] uint8 //col:1033
ModuleName UNICODE_STRING //col:1034
}



type COVERAGE_MODULE_INFO struct{
ModuleInfoSize uint32 //col:1039
IsBinaryLoaded uint32 //col:1040
ModulePathName UNICODE_STRING //col:1041
CoverageSectionSize uint32 //col:1042
CoverageSection[1] uint8 //col:1043
}



type COVERAGE_MODULES struct{
ListAndReset uint32 //col:1047
NumberOfModules uint32 //col:1048
ModuleRequestInfo COVERAGE_MODULE_REQUEST //col:1049
Modules[1] COVERAGE_MODULE_INFO //col:1050
}



type SYSTEM_PREFETCH_PATCH_INFORMATION struct{
PrefetchPatchCount uint32 //col:1054
}



type SYSTEM_VERIFIER_FAULTS_INFORMATION struct{
Probability uint32 //col:1058
MaxProbability uint32 //col:1059
PoolTags UNICODE_STRING //col:1060
Applications UNICODE_STRING //col:1061
}



type SYSTEM_VERIFIER_INFORMATION_EX struct{
VerifyMode uint32 //col:1065
OptionChanges uint32 //col:1066
PreviousBucketName UNICODE_STRING //col:1067
IrpCancelTimeoutMsec uint32 //col:1068
VerifierExtensionEnabled uint32 //col:1069
#ifdefWin64 #ifdef _WIN64 //col:1070
Reserved[1] uint32 //col:1071
#else #else //col:1072
Reserved[3] uint32 //col:1073
#endif #endif //col:1074
}



type SYSTEM_SYSTEM_PARTITION_INFORMATION struct{
SystemPartition UNICODE_STRING //col:1078
}



type SYSTEM_SYSTEM_DISK_INFORMATION struct{
SystemDisk UNICODE_STRING //col:1082
}



type SYSTEM_NUMA_PROXIMITY_MAP struct{
NodeProximityId uint32 //col:1086
NodeNumber USHORT //col:1087
}



type SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT struct{
Hits ULONGLONG //col:1091
PercentFrequency uint8 //col:1092
}



type SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT_WIN8 struct{
Hits uint32 //col:1096
PercentFrequency uint8 //col:1097
}



type SYSTEM_PROCESSOR_PERFORMANCE_STATE_DISTRIBUTION struct{
ProcessorNumber uint32 //col:1101
StateCount uint32 //col:1102
States[1] SYSTEM_PROCESSOR_PERFORMANCE_HITCOUNT //col:1103
}



type SYSTEM_PROCESSOR_PERFORMANCE_DISTRIBUTION struct{
ProcessorCount uint32 //col:1107
Offsets[1] uint32 //col:1108
}



type SYSTEM_CODEINTEGRITY_INFORMATION struct{
Length uint32 //col:1112
CodeIntegrityOptions uint32 //col:1113
}



type SYSTEM_PROCESSOR_MICROCODE_UPDATE_INFORMATION struct{
Operation uint32 //col:1117
}



type SYSTEM_VA_LIST_INFORMATION struct{
VirtualSize SIZE_T //col:1121
VirtualPeak SIZE_T //col:1122
VirtualLimit SIZE_T //col:1123
AllocationFailures SIZE_T //col:1124
}



type SYSTEM_STORE_INFORMATION struct{
ULONG _In_ //col:1128
STORE_INFORMATION_CLASS _In_ //col:1129
PVOID _Inout_ //col:1130
ULONG _Inout_ //col:1131
}



type SM_STATS_REQUEST struct{
Version uint32 //col:1135
DetailLevel uint32 //col:1136
StoreId uint32 //col:1137
BufferSize uint32 //col:1138
Buffer PVOID //col:1139
}



type ST_DATA_MGR_STATS struct{
RegionCount uint32 //col:1143
PagesStored uint32 //col:1144
UniquePagesStored uint32 //col:1145
LazyCleanupRegionCount uint32 //col:1146
RegionsInUse uint32 //col:1148
SpaceUsed uint32 //col:1149
}



type ST_IO_STATS_PERIOD struct{
PageCounts[5] uint32 //col:1154
}



type ST_IO_STATS struct{
PeriodCount uint32 //col:1158
Periods[64] ST_IO_STATS_PERIOD //col:1159
}



type ST_READ_LATENCY_BUCKET struct{
LatencyUs uint32 //col:1163
Count uint32 //col:1164
}



type ST_READ_LATENCY_STATS struct{
Buckets[8] ST_READ_LATENCY_BUCKET //col:1168
}



type ST_STATS_REGION_INFO struct{
SpaceUsed USHORT //col:1172
Priority uint8 //col:1173
Spare uint8 //col:1174
}



type ST_STATS_SPACE_BITMAP struct{
CompressedBytes SIZE_T //col:1178
BytesPerBit uint32 //col:1179
StoreBitmap[1] uint8 //col:1180
}



type ST_STATS struct{
Version uint32 //col:1184
Level uint32 //col:1185
StoreType uint32 //col:1186
NoDuplication uint32 //col:1187
NoCompression uint32 //col:1188
EncryptionStrength uint32 //col:1189
VirtualRegions uint32 //col:1190
Spare0 uint32 //col:1191
Size uint32 //col:1192
CompressionFormat USHORT //col:1193
Spare USHORT //col:1194
Struct struct //col:1195
RegionSize uint32 //col:1197
RegionCount uint32 //col:1198
RegionCountMax uint32 //col:1199
Granularity uint32 //col:1200
UserData ST_DATA_MGR_STATS //col:1201
Metadata ST_DATA_MGR_STATS //col:1202
}



type SM_STORE_BASIC_PARAMS struct{
Union union //col:1212
Struct struct //col:1214
StoreType uint32 //col:1216
NoDuplication uint32 //col:1217
FailNoCompression uint32 //col:1218
NoCompression uint32 //col:1219
NoEncryption uint32 //col:1220
NoEvictOnAdd uint32 //col:1221
PerformsFileIo uint32 //col:1222
VdlNotSet uint32 //col:1223
UseIntermediateAddBuffer uint32 //col:1224
CompressNoHuff uint32 //col:1225
LockActiveRegions uint32 //col:1226
VirtualRegions uint32 //col:1227
Spare uint32 //col:1228
}



type SMKM_REGION_EXTENT struct{
RegionCount uint32 //col:1238
ByteOffset SIZE_T //col:1239
}



type SMKM_FILE_INFO struct{
FileHandle HANDLE //col:1243
struct // //col:1244
struct // //col:1245
struct // //col:1246
VolumePnpHandle HANDLE //col:1247
struct // //col:1248
Extents PSMKM_REGION_EXTENT //col:1249
ExtentCount uint32 //col:1250
}



type SM_STORE_CACHE_BACKED_PARAMS struct{
SectorSize uint32 //col:1254
EncryptionKey PCHAR //col:1255
EncryptionKeySize uint32 //col:1256
FileInfo PSMKM_FILE_INFO //col:1257
EtaContext PVOID //col:1258
struct // //col:1259
}



type SM_STORE_PARAMETERS struct{
Store SM_STORE_BASIC_PARAMS //col:1263
Priority uint32 //col:1264
Flags uint32 //col:1265
CacheBacked SM_STORE_CACHE_BACKED_PARAMS //col:1266
}



type SM_CREATE_REQUEST struct{
Version uint32 //col:1270
AcquireReference uint32 //col:1271
KeyedStore uint32 //col:1272
Spare uint32 //col:1273
Params SM_STORE_PARAMETERS //col:1274
StoreId uint32 //col:1275
}



type SM_DELETE_REQUEST struct{
Version uint32 //col:1279
Spare uint32 //col:1280
StoreId uint32 //col:1281
}



type SM_STORE_LIST_REQUEST struct{
Version uint32 //col:1285
StoreCount uint32 //col:1286
ExtendedRequest uint32 //col:1287
Spare uint32 //col:1288
StoreId[32] uint32 //col:1289
}



type SM_STORE_LIST_REQUEST_EX struct{
Request SM_STORE_LIST_REQUEST //col:1293
NameBuffer[32][64] WCHAR //col:1294
}



type SMC_CACHE_LIST_REQUEST struct{
Version uint32 //col:1298
CacheCount uint32 //col:1299
Spare uint32 //col:1300
CacheId[16] uint32 //col:1301
}



type SMC_CACHE_PARAMETERS struct{
CacheFileSize SIZE_T //col:1305
StoreAlignment uint32 //col:1306
PerformsFileIo uint32 //col:1307
VdlNotSet uint32 //col:1308
Spare uint32 //col:1309
CacheFlags uint32 //col:1310
Priority uint32 //col:1311
}



type SMC_CACHE_CREATE_PARAMETERS struct{
CacheParameters SMC_CACHE_PARAMETERS //col:1315
TemplateFilePath[512] WCHAR //col:1316
}



type SMC_CACHE_CREATE_REQUEST struct{
Version uint32 //col:1320
Spare uint32 //col:1321
CacheId uint32 //col:1322
CacheCreateParams SMC_CACHE_CREATE_PARAMETERS //col:1323
}



type SMC_CACHE_DELETE_REQUEST struct{
Version uint32 //col:1327
Spare uint32 //col:1328
CacheId uint32 //col:1329
}



type SMC_STORE_CREATE_REQUEST struct{
Version uint32 //col:1333
Spare uint32 //col:1334
StoreParams SM_STORE_BASIC_PARAMS //col:1335
CacheId uint32 //col:1336
StoreManagerType SM_STORE_MANAGER_TYPE //col:1337
StoreId uint32 //col:1338
}



type SMC_STORE_DELETE_REQUEST struct{
Version uint32 //col:1342
Spare uint32 //col:1343
CacheId uint32 //col:1344
StoreManagerType SM_STORE_MANAGER_TYPE //col:1345
StoreId uint32 //col:1346
}



type SMC_CACHE_STATS struct{
TotalFileSize SIZE_T //col:1350
StoreCount uint32 //col:1351
RegionCount uint32 //col:1352
RegionSizeBytes uint32 //col:1353
FileCount uint32 //col:1354
PerformsFileIo uint32 //col:1355
Spare uint32 //col:1356
StoreIds[16] uint32 //col:1357
PhysicalStoreBitmap uint32 //col:1358
Priority uint32 //col:1359
TemplateFilePath[512] WCHAR //col:1360
}



type SMC_CACHE_STATS_REQUEST struct{
Version uint32 //col:1364
NoFilePath uint32 //col:1365
Spare uint32 //col:1366
CacheId uint32 //col:1367
CacheStats SMC_CACHE_STATS //col:1368
}



type SM_REGISTRATION_INFO struct{
CachesUpdatedEvent HANDLE //col:1372
}



type SM_REGISTRATION_REQUEST struct{
Version uint32 //col:1376
Spare uint32 //col:1377
RegInfo SM_REGISTRATION_INFO //col:1378
}



type SM_STORE_RESIZE_REQUEST struct{
Version uint32 //col:1382
AddRegions uint32 //col:1383
Spare uint32 //col:1384
StoreId uint32 //col:1385
NumberOfRegions uint32 //col:1386
struct // //col:1387
}



type SMC_STORE_RESIZE_REQUEST struct{
Version uint32 //col:1391
AddRegions uint32 //col:1392
Spare uint32 //col:1393
CacheId uint32 //col:1394
StoreId uint32 //col:1395
StoreManagerType SM_STORE_MANAGER_TYPE //col:1396
RegionCount uint32 //col:1397
}



type SM_CONFIG_REQUEST struct{
Version uint32 //col:1401
Spare uint32 //col:1402
ConfigType uint32 //col:1403
ConfigValue uint32 //col:1404
}



type SM_STORE_HIGH_MEM_PRIORITY_REQUEST struct{
Version uint32 //col:1408
SetHighMemoryPriority uint32 //col:1409
Spare uint32 //col:1410
ProcessHandle HANDLE //col:1411
}



type SM_SYSTEM_STORE_TRIM_REQUEST struct{
Version uint32 //col:1415
Spare uint32 //col:1416
PagesToTrim SIZE_T //col:1417
}



type SM_MEM_COMPRESSION_INFO_REQUEST struct{
Version uint32 //col:1421
Spare uint32 //col:1422
CompressionPid uint32 //col:1423
WorkingSetSize uint32 //col:1424
TotalDataCompressed SIZE_T //col:1425
TotalCompressedSize SIZE_T //col:1426
TotalUniqueDataCompressed SIZE_T //col:1427
}



type SYSTEM_REGISTRY_APPEND_STRING_PARAMETERS struct{
KeyHandle HANDLE //col:1431
ValueNamePointer PUNICODE_STRING //col:1432
RequiredLengthPointer PULONG //col:1433
Buffer PUCHAR //col:1434
BufferLength uint32 //col:1435
Type uint32 //col:1436
AppendBuffer PUCHAR //col:1437
AppendBufferLength uint32 //col:1438
CreateIfDoesntExist bool //col:1439
TruncateExistingValue bool //col:1440
}



type SYSTEM_VHD_BOOT_INFORMATION struct{
OsDiskIsVhd bool //col:1444
OsVhdFilePathOffset uint32 //col:1445
OsVhdParentVolume[1] WCHAR //col:1446
}



type PS_CPU_QUOTA_QUERY_ENTRY struct{
SessionId uint32 //col:1450
Weight uint32 //col:1451
}



type PS_CPU_QUOTA_QUERY_INFORMATION struct{
SessionCount uint32 //col:1455
SessionInformation[1] PS_CPU_QUOTA_QUERY_ENTRY //col:1456
}



type SYSTEM_ERROR_PORT_TIMEOUTS struct{
StartTimeout uint32 //col:1460
CommTimeout uint32 //col:1461
}



type SYSTEM_LOW_PRIORITY_IO_INFORMATION struct{
LowPriReadOperations uint32 //col:1465
LowPriWriteOperations uint32 //col:1466
KernelBumpedToNormalOperations uint32 //col:1467
LowPriPagingReadOperations uint32 //col:1468
KernelPagingReadsBumpedToNormal uint32 //col:1469
LowPriPagingWriteOperations uint32 //col:1470
KernelPagingWritesBumpedToNormal uint32 //col:1471
BoostedIrpCount uint32 //col:1472
BoostedPagingIrpCount uint32 //col:1473
BlanketBoostCount uint32 //col:1474
}



type TPM_BOOT_ENTROPY_NT_RESULT struct{
Policy ULONGLONG //col:1478
ResultCode TPM_BOOT_ENTROPY_RESULT_CODE //col:1479
ResultStatus NTSTATUS //col:1480
Time ULONGLONG //col:1481
EntropyLength uint32 //col:1482
EntropyData[40] uint8 //col:1483
}



type SYSTEM_VERIFIER_COUNTERS_INFORMATION struct{
Legacy SYSTEM_VERIFIER_INFORMATION //col:1487
RaiseIrqls uint32 //col:1488
AcquireSpinLocks uint32 //col:1489
SynchronizeExecutions uint32 //col:1490
AllocationsWithNoTag uint32 //col:1491
AllocationsFailed uint32 //col:1492
AllocationsFailedDeliberately uint32 //col:1493
LockedBytes SIZE_T //col:1494
PeakLockedBytes SIZE_T //col:1495
MappedLockedBytes SIZE_T //col:1496
PeakMappedLockedBytes SIZE_T //col:1497
MappedIoSpaceBytes SIZE_T //col:1498
PeakMappedIoSpaceBytes SIZE_T //col:1499
PagesForMdlBytes SIZE_T //col:1500
PeakPagesForMdlBytes SIZE_T //col:1501
ContiguousMemoryBytes SIZE_T //col:1502
PeakContiguousMemoryBytes SIZE_T //col:1503
ExecutePoolTypes uint32 //col:1504
ExecutePageProtections uint32 //col:1505
ExecutePageMappings uint32 //col:1506
ExecuteWriteSections uint32 //col:1507
SectionAlignmentFailures uint32 //col:1508
UnsupportedRelocs uint32 //col:1509
IATInExecutableSection uint32 //col:1510
}



type SYSTEM_ACPI_AUDIT_INFORMATION struct{
RsdpCount uint32 //col:1514
SameRsdt uint32 //col:1515
SlicPresent uint32 //col:1516
SlicDifferent uint32 //col:1517
}



type SYSTEM_BASIC_PERFORMANCE_INFORMATION struct{
AvailablePages SIZE_T //col:1521
CommittedPages SIZE_T //col:1522
CommitLimit SIZE_T //col:1523
PeakCommitment SIZE_T //col:1524
}



type QUERY_PERFORMANCE_COUNTER_FLAGS struct{
Union union //col:1528
Struct struct //col:1530
KernelTransition uint32 //col:1532
Reserved uint32 //col:1533
}



type SYSTEM_QUERY_PERFORMANCE_COUNTER_INFORMATION struct{
Version uint32 //col:1540
Flags QUERY_PERFORMANCE_COUNTER_FLAGS //col:1541
ValidFlags QUERY_PERFORMANCE_COUNTER_FLAGS //col:1542
}



type SYSTEM_BOOT_GRAPHICS_INFORMATION struct{
FrameBuffer LARGE_INTEGER //col:1546
Width uint32 //col:1547
Height uint32 //col:1548
PixelStride uint32 //col:1549
Flags uint32 //col:1550
Format SYSTEM_PIXEL_FORMAT //col:1551
DisplayRotation uint32 //col:1552
}



type MEMORY_SCRUB_INFORMATION struct{
Handle HANDLE //col:1556
PagesScrubbed uint32 //col:1557
}



type PEBS_DS_SAVE_AREA32 struct{
BtsBufferBase uint32 //col:1561
BtsIndex uint32 //col:1562
BtsAbsoluteMaximum uint32 //col:1563
BtsInterruptThreshold uint32 //col:1564
PebsBufferBase uint32 //col:1565
PebsIndex uint32 //col:1566
PebsAbsoluteMaximum uint32 //col:1567
PebsInterruptThreshold uint32 //col:1568
PebsGpCounterReset[8] uint32 //col:1569
PebsFixedCounterReset[4] uint32 //col:1570
}



type PEBS_DS_SAVE_AREA64 struct{
BtsBufferBase ULONGLONG //col:1574
BtsIndex ULONGLONG //col:1575
BtsAbsoluteMaximum ULONGLONG //col:1576
BtsInterruptThreshold ULONGLONG //col:1577
PebsBufferBase ULONGLONG //col:1578
PebsIndex ULONGLONG //col:1579
PebsAbsoluteMaximum ULONGLONG //col:1580
PebsInterruptThreshold ULONGLONG //col:1581
PebsGpCounterReset[8] ULONGLONG //col:1582
PebsFixedCounterReset[4] ULONGLONG //col:1583
}



type PROCESSOR_PROFILE_CONTROL_AREA struct{
PebsDsSaveArea PEBS_DS_SAVE_AREA //col:1587
}



type SYSTEM_PROCESSOR_PROFILE_CONTROL_AREA struct{
ProcessorProfileControlArea PROCESSOR_PROFILE_CONTROL_AREA //col:1591
Allocate bool //col:1592
}



type MEMORY_COMBINE_INFORMATION struct{
Handle HANDLE //col:1596
PagesCombined ULONG_PTR //col:1597
}



type MEMORY_COMBINE_INFORMATION_EX struct{
Handle HANDLE //col:1601
PagesCombined ULONG_PTR //col:1602
Flags uint32 //col:1603
}



type MEMORY_COMBINE_INFORMATION_EX2 struct{
Handle HANDLE //col:1607
PagesCombined ULONG_PTR //col:1608
Flags uint32 //col:1609
ProcessHandle HANDLE //col:1610
}



type SYSTEM_ENTROPY_TIMING_INFORMATION struct{
(NTAPI VOID //col:1614
(NTAPI VOID //col:1615
InitializationContext PVOID //col:1616
}



type SYSTEM_CONSOLE_INFORMATION struct{
DriverLoaded uint32 //col:1620
Spare uint32 //col:1621
}



type SYSTEM_PLATFORM_BINARY_INFORMATION struct{
PhysicalAddress ULONG64 //col:1625
HandoffBuffer PVOID //col:1626
CommandLineBuffer PVOID //col:1627
HandoffBufferSize uint32 //col:1628
CommandLineBufferSize uint32 //col:1629
}



type SYSTEM_POLICY_INFORMATION struct{
InputData PVOID //col:1633
OutputData PVOID //col:1634
InputDataSize uint32 //col:1635
OutputDataSize uint32 //col:1636
Version uint32 //col:1637
}



type SYSTEM_HYPERVISOR_PROCESSOR_COUNT_INFORMATION struct{
NumberOfLogicalProcessors uint32 //col:1641
NumberOfCores uint32 //col:1642
}



type SYSTEM_DEVICE_DATA_INFORMATION struct{
DeviceId UNICODE_STRING //col:1646
DataName UNICODE_STRING //col:1647
DataType uint32 //col:1648
DataBufferLength uint32 //col:1649
DataBuffer PVOID //col:1650
}



type PHYSICAL_CHANNEL_RUN struct{
NodeNumber uint32 //col:1654
ChannelNumber uint32 //col:1655
BasePage ULONGLONG //col:1656
PageCount ULONGLONG //col:1657
Flags uint32 //col:1658
}



type SYSTEM_MEMORY_TOPOLOGY_INFORMATION struct{
NumberOfRuns ULONGLONG //col:1662
NumberOfNodes uint32 //col:1663
NumberOfChannels uint32 //col:1664
Run[1] PHYSICAL_CHANNEL_RUN //col:1665
}



type SYSTEM_MEMORY_CHANNEL_INFORMATION struct{
ChannelNumber uint32 //col:1669
ChannelHeatIndex uint32 //col:1670
TotalPageCount ULONGLONG //col:1671
ZeroPageCount ULONGLONG //col:1672
FreePageCount ULONGLONG //col:1673
StandbyPageCount ULONGLONG //col:1674
}



type SYSTEM_BOOT_LOGO_INFORMATION struct{
Flags uint32 //col:1678
BitmapOffset uint32 //col:1679
}



type SYSTEM_PROCESSOR_PERFORMANCE_INFORMATION_EX struct{
IdleTime LARGE_INTEGER //col:1683
KernelTime LARGE_INTEGER //col:1684
UserTime LARGE_INTEGER //col:1685
DpcTime LARGE_INTEGER //col:1686
InterruptTime LARGE_INTEGER //col:1687
InterruptCount uint32 //col:1688
Spare0 uint32 //col:1689
AvailableTime LARGE_INTEGER //col:1690
Spare1 LARGE_INTEGER //col:1691
Spare2 LARGE_INTEGER //col:1692
}



type SYSTEM_SECUREBOOT_POLICY_INFORMATION struct{
PolicyPublisher GUID //col:1696
PolicyVersion uint32 //col:1697
PolicyOptions uint32 //col:1698
}



type SYSTEM_PAGEFILE_INFORMATION_EX struct{
Union union //col:1702
Info SYSTEM_PAGEFILE_INFORMATION //col:1704
Struct struct //col:1705
NextEntryOffset uint32 //col:1707
TotalSize uint32 //col:1708
TotalInUse uint32 //col:1709
PeakUsage uint32 //col:1710
PageFileName UNICODE_STRING //col:1711
}



type SYSTEM_SECUREBOOT_INFORMATION struct{
SecureBootEnabled bool //col:1719
SecureBootCapable bool //col:1720
}



type PROCESS_DISK_COUNTERS struct{
BytesRead ULONGLONG //col:1724
BytesWritten ULONGLONG //col:1725
ReadOperationCount ULONGLONG //col:1726
WriteOperationCount ULONGLONG //col:1727
FlushOperationCount ULONGLONG //col:1728
}



type PROCESS_ENERGY_VALUES struct{
Cycles[4][2] ULONGLONG //col:1732
DiskEnergy ULONGLONG //col:1733
NetworkTailEnergy ULONGLONG //col:1734
MBBTailEnergy ULONGLONG //col:1735
NetworkTxRxBytes ULONGLONG //col:1736
MBBTxRxBytes ULONGLONG //col:1737
Union union //col:1738
Durations[3] ENERGY_STATE_DURATION //col:1740
Struct struct //col:1741
ForegroundDuration ENERGY_STATE_DURATION //col:1743
DesktopVisibleDuration ENERGY_STATE_DURATION //col:1744
PSMForegroundDuration ENERGY_STATE_DURATION //col:1745
}



type PROCESS_ENERGY_VALUES_EXTENSION struct{
Union union //col:1757
Timelines[14] TIMELINE_BITMAP //col:1759
Struct struct //col:1760
CpuTimeline TIMELINE_BITMAP //col:1762
DiskTimeline TIMELINE_BITMAP //col:1763
NetworkTimeline TIMELINE_BITMAP //col:1764
MBBTimeline TIMELINE_BITMAP //col:1765
ForegroundTimeline TIMELINE_BITMAP //col:1766
DesktopVisibleTimeline TIMELINE_BITMAP //col:1767
CompositionRenderedTimeline TIMELINE_BITMAP //col:1768
CompositionDirtyGeneratedTimeline TIMELINE_BITMAP //col:1769
CompositionDirtyPropagatedTimeline TIMELINE_BITMAP //col:1770
InputTimeline TIMELINE_BITMAP //col:1771
AudioInTimeline TIMELINE_BITMAP //col:1772
AudioOutTimeline TIMELINE_BITMAP //col:1773
DisplayRequiredTimeline TIMELINE_BITMAP //col:1774
KeyboardInputTimeline TIMELINE_BITMAP //col:1775
}



type PROCESS_EXTENDED_ENERGY_VALUES struct{
Base PROCESS_ENERGY_VALUES //col:1795
Extension PROCESS_ENERGY_VALUES_EXTENSION //col:1796
}



type SYSTEM_PROCESS_INFORMATION_EXTENSION struct{
DiskCounters PROCESS_DISK_COUNTERS //col:1800
ContextSwitches ULONGLONG //col:1801
Union union //col:1802
Flags uint32 //col:1804
Struct struct //col:1805
HasStrongId uint32 //col:1807
Classification uint32 //col:1808
BackgroundActivityModerated uint32 //col:1809
Spare uint32 //col:1810
}



type SYSTEM_PORTABLE_WORKSPACE_EFI_LAUNCHER_INFORMATION struct{
EfiLauncherEnabled bool //col:1824
}



type SYSTEM_KERNEL_DEBUGGER_INFORMATION_EX struct{
DebuggerAllowed bool //col:1828
DebuggerEnabled bool //col:1829
DebuggerPresent bool //col:1830
}



type SYSTEM_ELAM_CERTIFICATE_INFORMATION struct{
ElamDriverFile HANDLE //col:1834
}



type OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V2 struct{
Version uint32 //col:1838
AbnormalResetOccurred uint32 //col:1839
OfflineMemoryDumpCapable uint32 //col:1840
ResetDataAddress LARGE_INTEGER //col:1841
ResetDataSize uint32 //col:1842
}



type OFFLINE_CRASHDUMP_CONFIGURATION_TABLE_V1 struct{
Version uint32 //col:1846
AbnormalResetOccurred uint32 //col:1847
OfflineMemoryDumpCapable uint32 //col:1848
}



type SYSTEM_PROCESSOR_FEATURES_INFORMATION struct{
ProcessorFeatureBits ULONGLONG //col:1852
Reserved[3] ULONGLONG //col:1853
}



type SYSTEM_EDID_INFORMATION struct{
Edid[128] uint8 //col:1857
}



type SYSTEM_MANUFACTURING_INFORMATION struct{
Options uint32 //col:1861
ProfileName UNICODE_STRING //col:1862
}



type SYSTEM_ENERGY_ESTIMATION_CONFIG_INFORMATION struct{
Enabled bool //col:1866
}



type HV_DETAILS struct{
Data[4] uint32 //col:1870
}



type SYSTEM_HYPERVISOR_DETAIL_INFORMATION struct{
HvVendorAndMaxFunction HV_DETAILS //col:1874
HypervisorInterface HV_DETAILS //col:1875
HypervisorVersion HV_DETAILS //col:1876
HvFeatures HV_DETAILS //col:1877
HwFeatures HV_DETAILS //col:1878
EnlightenmentInfo HV_DETAILS //col:1879
ImplementationLimits HV_DETAILS //col:1880
}



type SYSTEM_PROCESSOR_CYCLE_STATS_INFORMATION struct{
Cycles[4][2] ULONGLONG //col:1884
}



type SYSTEM_TPM_INFORMATION struct{
Flags uint32 //col:1888
}



type SYSTEM_VSM_PROTECTION_INFORMATION struct{
DmaProtectionsAvailable bool //col:1892
DmaProtectionsInUse bool //col:1893
HardwareMbecAvailable bool //col:1894
ApicVirtualizationAvailable bool //col:1895
}



type SYSTEM_KERNEL_DEBUGGER_FLAGS struct{
KernelDebuggerIgnoreUmExceptions bool //col:1899
}



type SYSTEM_CODEINTEGRITYPOLICY_INFORMATION struct{
Options uint32 //col:1903
HVCIOptions uint32 //col:1904
Version ULONGLONG //col:1905
PolicyGuid GUID //col:1906
}



type SYSTEM_ISOLATED_USER_MODE_INFORMATION struct{
SecureKernelRunning bool //col:1910
HvciEnabled bool //col:1911
HvciStrictMode bool //col:1912
DebugEnabled bool //col:1913
FirmwarePageProtection bool //col:1914
EncryptionKeyAvailable bool //col:1915
SpareFlags bool //col:1916
TrustletRunning bool //col:1917
HvciDisableAllowed bool //col:1918
SpareFlags2 bool //col:1919
Spare0[6] bool //col:1920
Spare1 ULONGLONG //col:1921
}



type SYSTEM_SINGLE_MODULE_INFORMATION struct{
TargetModuleAddress PVOID //col:1925
ExInfo RTL_PROCESS_MODULE_INFORMATION_EX //col:1926
}



type SYSTEM_INTERRUPT_CPU_SET_INFORMATION struct{
Gsiv uint32 //col:1930
Group USHORT //col:1931
CpuSets ULONGLONG //col:1932
}



type SYSTEM_SECUREBOOT_POLICY_FULL_INFORMATION struct{
PolicyInformation SYSTEM_SECUREBOOT_POLICY_INFORMATION //col:1936
PolicySize uint32 //col:1937
Policy[1] uint8 //col:1938
}



type SYSTEM_ROOT_SILO_INFORMATION struct{
NumberOfSilos uint32 //col:1942
SiloIdList[1] uint32 //col:1943
}



type SYSTEM_CPU_SET_TAG_INFORMATION struct{
Tag ULONGLONG //col:1947
CpuSets[1] ULONGLONG //col:1948
}



type SYSTEM_SECURE_KERNEL_HYPERGUARD_PROFILE_INFORMATION struct{
ExtentCount uint32 //col:1952
ValidStructureSize uint32 //col:1953
NextExtentIndex uint32 //col:1954
ExtentRestart uint32 //col:1955
CycleCount uint32 //col:1956
TimeoutCount uint32 //col:1957
CycleTime ULONGLONG //col:1958
CycleTimeMax ULONGLONG //col:1959
ExtentTime ULONGLONG //col:1960
ExtentTimeIndex uint32 //col:1961
ExtentTimeMaxIndex uint32 //col:1962
ExtentTimeMax ULONGLONG //col:1963
HyperFlushTimeMax ULONGLONG //col:1964
TranslateVaTimeMax ULONGLONG //col:1965
DebugExemptionCount ULONGLONG //col:1966
TbHitCount ULONGLONG //col:1967
TbMissCount ULONGLONG //col:1968
VinaPendingYield ULONGLONG //col:1969
HashCycles ULONGLONG //col:1970
HistogramOffset uint32 //col:1971
HistogramBuckets uint32 //col:1972
HistogramShift uint32 //col:1973
Reserved1 uint32 //col:1974
PageNotPresentCount ULONGLONG //col:1975
}



type SYSTEM_SECUREBOOT_PLATFORM_MANIFEST_INFORMATION struct{
PlatformManifestSize uint32 //col:1979
PlatformManifest[1] uint8 //col:1980
}



type SYSTEM_INTERRUPT_STEERING_INFORMATION_INPUT struct{
Gsiv uint32 //col:1984
ControllerInterrupt uint8 //col:1985
EdgeInterrupt uint8 //col:1986
IsPrimaryInterrupt uint8 //col:1987
TargetAffinity GROUP_AFFINITY //col:1988
}



type SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION struct{
Machine uint32 //col:1992
KernelMode uint32 //col:1993
UserMode uint32 //col:1994
Native uint32 //col:1995
Process uint32 //col:1996
WoW64Container uint32 //col:1997
ReservedZero0 uint32 //col:1998
}



type SYSTEM_MEMORY_USAGE_INFORMATION struct{
TotalPhysicalBytes ULONGLONG //col:2002
AvailableBytes ULONGLONG //col:2003
ResidentAvailableBytes LONGLONG //col:2004
CommittedBytes ULONGLONG //col:2005
SharedCommittedBytes ULONGLONG //col:2006
CommitLimitBytes ULONGLONG //col:2007
PeakCommitmentBytes ULONGLONG //col:2008
}



type SYSTEM_CODEINTEGRITY_CERTIFICATE_INFORMATION struct{
ImageFile HANDLE //col:2012
Type uint32 //col:2013
}



type SYSTEM_PHYSICAL_MEMORY_INFORMATION struct{
TotalPhysicalBytes ULONGLONG //col:2017
LowestPhysicalAddress ULONGLONG //col:2018
HighestPhysicalAddress ULONGLONG //col:2019
}



type SYSTEM_ACTIVITY_MODERATION_EXE_STATE  struct{
ExePathNt UNICODE_STRING //col:2023
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE //col:2024
}



type SYSTEM_ACTIVITY_MODERATION_INFO struct{
Identifier UNICODE_STRING //col:2028
ModerationState SYSTEM_ACTIVITY_MODERATION_STATE //col:2029
AppType SYSTEM_ACTIVITY_MODERATION_APP_TYPE //col:2030
}



type SYSTEM_ACTIVITY_MODERATION_USER_SETTINGS struct{
UserKeyHandle HANDLE //col:2034
}



type SYSTEM_CODEINTEGRITY_UNLOCK_INFORMATION struct{
Union union //col:2038
Flags uint32 //col:2040
Struct struct //col:2041
Locked uint32 //col:2043
UnlockApplied uint32 //col:2044
UnlockIdValid uint32 //col:2045
Reserved uint32 //col:2046
}



type SYSTEM_FLUSH_INFORMATION struct{
SupportedFlushMethods uint32 //col:2053
ProcessorCacheFlushSize uint32 //col:2054
SystemFlushCapabilities ULONGLONG //col:2055
Reserved[2] ULONGLONG //col:2056
}



type SYSTEM_WRITE_CONSTRAINT_INFORMATION struct{
WriteConstraintPolicy uint32 //col:2060
Reserved uint32 //col:2061
}



type SYSTEM_KERNEL_VA_SHADOW_INFORMATION struct{
Union union //col:2065
KvaShadowFlags uint32 //col:2067
Struct struct //col:2068
KvaShadowEnabled uint32 //col:2070
KvaShadowUserGlobal uint32 //col:2071
KvaShadowPcid uint32 //col:2072
KvaShadowInvpcid uint32 //col:2073
KvaShadowRequired uint32 //col:2074
KvaShadowRequiredAvailable uint32 //col:2075
InvalidPteBit uint32 //col:2076
L1DataCacheFlushSupported uint32 //col:2077
L1TerminalFaultMitigationPresent uint32 //col:2078
Reserved uint32 //col:2079
}



type SYSTEM_CODEINTEGRITYVERIFICATION_INFORMATION struct{
FileHandle HANDLE //col:2085
ImageSize uint32 //col:2086
Image PVOID //col:2087
}



type SYSTEM_HYPERVISOR_SHARED_PAGE_INFORMATION struct{
HypervisorSharedUserVa PVOID //col:2091
}



type SYSTEM_FIRMWARE_PARTITION_INFORMATION struct{
FirmwarePartition UNICODE_STRING //col:2095
}



type SYSTEM_SPECULATION_CONTROL_INFORMATION struct{
Union union //col:2099
Flags uint32 //col:2101
Struct struct //col:2102
BpbEnabled uint32 //col:2104
BpbDisabledSystemPolicy uint32 //col:2105
BpbDisabledNoHardwareSupport uint32 //col:2106
SpecCtrlEnumerated uint32 //col:2107
SpecCmdEnumerated uint32 //col:2108
IbrsPresent uint32 //col:2109
StibpPresent uint32 //col:2110
SmepPresent uint32 //col:2111
SpeculativeStoreBypassDisableAvailable uint32 //col:2112
SpeculativeStoreBypassDisableSupported uint32 //col:2113
SpeculativeStoreBypassDisabledSystemWide uint32 //col:2114
SpeculativeStoreBypassDisabledKernel uint32 //col:2115
SpeculativeStoreBypassDisableRequired uint32 //col:2116
BpbDisabledKernelToUser uint32 //col:2117
SpecCtrlRetpolineEnabled uint32 //col:2118
SpecCtrlImportOptimizationEnabled uint32 //col:2119
EnhancedIbrs uint32 //col:2120
HvL1tfStatusAvailable uint32 //col:2121
HvL1tfProcessorNotAffected uint32 //col:2122
HvL1tfMigitationEnabled uint32 //col:2123
HvL1tfMigitationNotEnabled_Hardware uint32 //col:2124
HvL1tfMigitationNotEnabled_LoadOption uint32 //col:2125
HvL1tfMigitationNotEnabled_CoreScheduler uint32 //col:2126
EnhancedIbrsReported uint32 //col:2127
MdsHardwareProtected uint32 //col:2128
MbClearEnabled uint32 //col:2129
MbClearReported uint32 //col:2130
Reserved uint32 //col:2131
}



type SYSTEM_DMA_GUARD_POLICY_INFORMATION struct{
DmaGuardPolicyEnabled bool //col:2137
}



type SYSTEM_ENCLAVE_LAUNCH_CONTROL_INFORMATION struct{
EnclaveLaunchSigner[32] uint8 //col:2141
}



type SYSTEM_WORKLOAD_ALLOWED_CPU_SET_INFORMATION struct{
WorkloadClass ULONGLONG //col:2145
CpuSets[1] ULONGLONG //col:2146
}



type SYSTEM_SECURITY_MODEL_INFORMATION struct{
Union union //col:2150
SecurityModelFlags uint32 //col:2152
Struct struct //col:2153
SModeAdminlessEnabled uint32 //col:2155
AllowDeviceOwnerProtectionDowngrade uint32 //col:2156
Reserved uint32 //col:2157
}



type SYSTEM_FEATURE_CONFIGURATION_INFORMATION struct{
ChangeStamp ULONGLONG //col:2163
struct // //col:2164
}



type SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY struct{
ChangeStamp ULONGLONG //col:2168
Section PVOID //col:2169
Size ULONGLONG //col:2170
}



type SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION struct{
OverallChangeStamp ULONGLONG //col:2174
Descriptors[3] SYSTEM_FEATURE_CONFIGURATION_SECTIONS_INFORMATION_ENTRY //col:2175
}



type RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET struct{
Data[2] uint32 //col:2179
}



type SYSTEM_FEATURE_USAGE_SUBSCRIPTION_DETAILS struct{
FeatureId uint32 //col:2183
ReportingKind USHORT //col:2184
ReportingOptions USHORT //col:2185
ReportingTarget RTL_FEATURE_USAGE_SUBSCRIPTION_TARGET //col:2186
}



type SYSTEM_FIRMWARE_RAMDISK_INFORMATION struct{
Version uint32 //col:2190
BlockSize uint32 //col:2191
BaseAddress ULONG_PTR //col:2192
Size SIZE_T //col:2193
}



type SYSTEM_SHADOW_STACK_INFORMATION struct{
Union union //col:2197
Flags uint32 //col:2199
Struct struct //col:2200
CetCapable uint32 //col:2202
UserCetAllowed uint32 //col:2203
ReservedForUserCet uint32 //col:2204
KernelCetEnabled uint32 //col:2205
KernelCetAuditModeEnabled uint32 //col:2206
ReservedForKernelCet uint32 //col:2207
Reserved uint32 //col:2208
}



type SYSTEM_BUILD_VERSION_INFORMATION struct{
LayerNumber USHORT //col:2214
LayerCount USHORT //col:2215
OsMajorVersion uint32 //col:2216
OsMinorVersion uint32 //col:2217
NtBuildNumber uint32 //col:2218
NtBuildQfe uint32 //col:2219
LayerName[128] uint8 //col:2220
NtBuildBranch[128] uint8 //col:2221
NtBuildLab[128] uint8 //col:2222
NtBuildLabEx[128] uint8 //col:2223
NtBuildStamp[26] uint8 //col:2224
NtBuildArch[16] uint8 //col:2225
Flags SYSTEM_BUILD_VERSION_INFORMATION_FLAGS //col:2226
}



type SYSTEM_POOL_LIMIT_MEM_INFO struct{
MemoryLimit ULONGLONG //col:2230
NotificationLimit ULONGLONG //col:2231
}



type SYSTEM_POOL_LIMIT_INFO struct{
PoolTag uint32 //col:2235
MemLimits[2] SYSTEM_POOL_LIMIT_MEM_INFO //col:2236
NotificationHandle WNF_STATE_NAME //col:2237
}



type SYSTEM_POOL_LIMIT_INFORMATION struct{
Version uint32 //col:2241
EntryCount uint32 //col:2242
LimitEntries[1] SYSTEM_POOL_LIMIT_INFO //col:2243
}



type HV_MINROOT_NUMA_LPS struct{
NodeIndex uint32 //col:2247
Mask[16] ULONG_PTR //col:2248
}



type SYSTEM_IOMMU_STATE_INFORMATION struct{
State SYSTEM_IOMMU_STATE //col:2252
Pdo PVOID //col:2253
}



type SYSTEM_HYPERVISOR_MINROOT_INFORMATION struct{
NumProc uint32 //col:2257
RootProc uint32 //col:2258
RootProcNumaNodesSpecified uint32 //col:2259
RootProcNumaNodes[64] USHORT //col:2260
RootProcPerCore uint32 //col:2261
RootProcPerNode uint32 //col:2262
RootProcNumaNodesLpsSpecified uint32 //col:2263
RootProcNumaNodeLps[64] HV_MINROOT_NUMA_LPS //col:2264
}



type SYSTEM_HYPERVISOR_BOOT_PAGES_INFORMATION struct{
RangeCount uint32 //col:2268
RangeArray[1] ULONG_PTR //col:2269
}



type SYSTEM_POINTER_AUTH_INFORMATION struct{
Union union //col:2273
SupportedFlags USHORT //col:2275
Struct struct //col:2276
AddressAuthSupported USHORT //col:2278
AddressAuthQarma USHORT //col:2279
GenericAuthSupported USHORT //col:2280
GenericAuthQarma USHORT //col:2281
SupportedReserved USHORT //col:2282
}



type SYSDBG_VIRTUAL struct{
Address PVOID //col:2300
Buffer PVOID //col:2301
Request uint32 //col:2302
}



type SYSDBG_PHYSICAL struct{
Address PHYSICAL_ADDRESS //col:2306
Buffer PVOID //col:2307
Request uint32 //col:2308
}



type SYSDBG_CONTROL_SPACE struct{
Address ULONG64 //col:2312
Buffer PVOID //col:2313
Request uint32 //col:2314
Processor uint32 //col:2315
}



type SYSDBG_IO_SPACE struct{
Address ULONG64 //col:2319
Buffer PVOID //col:2320
Request uint32 //col:2321
_INTERFACE_TYPE enum //col:2322
BusNumber uint32 //col:2323
AddressSpace uint32 //col:2324
}



type SYSDBG_MSR struct{
Msr uint32 //col:2328
Data ULONG64 //col:2329
}



type SYSDBG_BUS_DATA struct{
Address uint32 //col:2333
Buffer PVOID //col:2334
Request uint32 //col:2335
_BUS_DATA_TYPE enum //col:2336
BusNumber uint32 //col:2337
SlotNumber uint32 //col:2338
}



type SYSDBG_TRIAGE_DUMP struct{
Flags uint32 //col:2342
BugCheckCode uint32 //col:2343
BugCheckParam1 ULONG_PTR //col:2344
BugCheckParam2 ULONG_PTR //col:2345
BugCheckParam3 ULONG_PTR //col:2346
BugCheckParam4 ULONG_PTR //col:2347
ProcessHandles uint32 //col:2348
ThreadHandles uint32 //col:2349
Handles PHANDLE //col:2350
}



type SYSDBG_LIVEDUMP_SELECTIVE_CONTROL struct{
Version uint32 //col:2354
Size uint32 //col:2355
Union union //col:2356
Flags ULONGLONG //col:2358
Struct struct //col:2359
ThreadKernelStacks ULONGLONG //col:2361
ReservedFlags ULONGLONG //col:2362
}



type SYSDBG_LIVEDUMP_CONTROL struct{
Version uint32 //col:2369
BugCheckCode uint32 //col:2370
BugCheckParam1 ULONG_PTR //col:2371
BugCheckParam2 ULONG_PTR //col:2372
BugCheckParam3 ULONG_PTR //col:2373
BugCheckParam4 ULONG_PTR //col:2374
DumpFileHandle HANDLE //col:2375
CancelEventHandle HANDLE //col:2376
Flags SYSDBG_LIVEDUMP_CONTROL_FLAGS //col:2377
AddPagesControl SYSDBG_LIVEDUMP_CONTROL_ADDPAGES //col:2378
SelectiveControl PSYSDBG_LIVEDUMP_SELECTIVE_CONTROL //col:2379
}



type SYSDBG_KD_PULL_REMOTE_FILE struct{
ImageFileName UNICODE_STRING //col:2383
}



type KUSER_SHARED_DATA struct{
TickCountLowDeprecated uint32 //col:2387
TickCountMultiplier uint32 //col:2388
KSYSTEM_TIME volatile //col:2389
KSYSTEM_TIME volatile //col:2390
KSYSTEM_TIME volatile //col:2391
ImageNumberLow USHORT //col:2392
ImageNumberHigh USHORT //col:2393
NtSystemRoot[260] WCHAR //col:2394
MaxStackTraceDepth uint32 //col:2395
CryptoExponent uint32 //col:2396
TimeZoneId uint32 //col:2397
LargePageMinimum uint32 //col:2398
AitSamplingValue uint32 //col:2399
AppCompatFlag uint32 //col:2400
RNGSeedVersion ULONGLONG //col:2401
GlobalValidationRunlevel uint32 //col:2402
TimeZoneBiasStamp LONG //col:2403
NtBuildNumber uint32 //col:2404
NtProductType NT_PRODUCT_TYPE //col:2405
ProductTypeIsValid bool //col:2406
Reserved0[1] uint8 //col:2407
NativeProcessorArchitecture USHORT //col:2408
NtMajorVersion uint32 //col:2409
NtMinorVersion uint32 //col:2410
ProcessorFeatures[PROCESSOR_FEATURE_MAX] bool //col:2411
Reserved1 uint32 //col:2412
Reserved3 uint32 //col:2413
ULONG volatile //col:2414
AlternativeArchitecture ALTERNATIVE_ARCHITECTURE_TYPE //col:2415
BootId uint32 //col:2416
SystemExpirationDate LARGE_INTEGER //col:2417
SuiteMask uint32 //col:2418
KdDebuggerEnabled bool //col:2419
Union union //col:2420
MitigationPolicies uint8 //col:2422
Struct struct //col:2423
NXSupportPolicy uint8 //col:2425
SEHValidationPolicy uint8 //col:2426
CurDirDevicesSkippedForDlls uint8 //col:2427
Reserved uint8 //col:2428
}



type ATOM_BASIC_INFORMATION struct{
UsageCount USHORT //col:2524
Flags USHORT //col:2525
NameLength USHORT //col:2526
Name[1] WCHAR //col:2527
}



type ATOM_TABLE_INFORMATION struct{
NumberOfAtoms uint32 //col:2531
Atoms[1] RTL_ATOM //col:2532
}





