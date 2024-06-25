package constants

import "fmt"

type ErrorsKind int

const (
	ErrorSuccess                                                 ErrorsKind = 0
	ErrorInvalidFunction                                                    = 1
	ErrorFileNotFound                                                       = 2
	ErrorPathNotFound                                                       = 3
	ErrorTooManyOpenFiles                                                   = 4
	ErrorAccessDenied                                                       = 5
	ErrorInvalidHandle                                                      = 6
	ErrorArenaTrashed                                                       = 7
	ErrorNotEnoughMemory                                                    = 8
	ErrorInvalidBlock                                                       = 9
	ErrorBadEnvironment                                                     = 10
	ErrorBadFormat                                                          = 11
	ErrorInvalidAccess                                                      = 12
	ErrorInvalidData                                                        = 13
	ErrorOutofmemory                                                        = 14
	ErrorInvalidDrive                                                       = 15
	ErrorCurrentDirectory                                                   = 16
	ErrorNotSameDevice                                                      = 17
	ErrorNoMoreFiles                                                        = 18
	ErrorWriteProtect                                                       = 19
	ErrorBadUnit                                                            = 20
	ErrorNotReady                                                           = 21
	ErrorBadCommand                                                         = 22
	ErrorCrc                                                                = 23
	ErrorBadLength                                                          = 24
	ErrorSeek                                                               = 25
	ErrorNotDosDisk                                                         = 26
	ErrorSectorNotFound                                                     = 27
	ErrorOutOfPaper                                                         = 28
	ErrorWriteFault                                                         = 29
	ErrorReadFault                                                          = 30
	ErrorGenFailure                                                         = 31
	ErrorSharingViolation                                                   = 32
	ErrorLockViolation                                                      = 33
	ErrorWrongDisk                                                          = 34
	ErrorSharingBufferExceeded                                              = 36
	ErrorHandleEof                                                          = 38
	ErrorHandleDiskFull                                                     = 39
	ErrorNotSupported                                                       = 50
	ErrorRemNotList                                                         = 51
	ErrorDupName                                                            = 52
	ErrorBadNetpath                                                         = 53
	ErrorNetworkBusy                                                        = 54
	ErrorDevNotExist                                                        = 55
	ErrorTooManyCmds                                                        = 56
	ErrorAdapHdwErr                                                         = 57
	ErrorBadNetResp                                                         = 58
	ErrorUnexpNetErr                                                        = 59
	ErrorBadRemAdap                                                         = 60
	ErrorPrintqFull                                                         = 61
	ErrorNoSpoolSpace                                                       = 62
	ErrorPrintCancelled                                                     = 63
	ErrorNetnameDeleted                                                     = 64
	ErrorNetworkAccessDenied                                                = 65
	ErrorBadDevType                                                         = 66
	ErrorBadNetName                                                         = 67
	ErrorTooManyNames                                                       = 68
	ErrorTooManySess                                                        = 69
	ErrorSharingPaused                                                      = 70
	ErrorReqNotAccep                                                        = 71
	ErrorRedirPaused                                                        = 72
	ErrorFileExists                                                         = 80
	ErrorCannotMake                                                         = 82
	ErrorFailI24                                                            = 83
	ErrorOutOfStructures                                                    = 84
	ErrorAlreadyAssigned                                                    = 85
	ErrorInvalidPassword                                                    = 86
	ErrorInvalidParameter                                                   = 87
	ErrorNetWriteFault                                                      = 88
	ErrorNoProcSlots                                                        = 89
	ErrorTooManySemaphores                                                  = 100
	ErrorExclSemAlreadyOwned                                                = 101
	ErrorSemIsSet                                                           = 102
	ErrorTooManySemRequests                                                 = 103
	ErrorInvalidAtInterruptTime                                             = 104
	ErrorSemOwnerDied                                                       = 105
	ErrorSemUserLimit                                                       = 106
	ErrorDiskChange                                                         = 107
	ErrorDriveLocked                                                        = 108
	ErrorBrokenPipe                                                         = 109
	ErrorOpenFailed                                                         = 110
	ErrorBufferOverflow                                                     = 111
	ErrorDiskFull                                                           = 112
	ErrorNoMoreSearchHandles                                                = 113
	ErrorInvalidTargetHandle                                                = 114
	ErrorInvalidCategory                                                    = 117
	ErrorInvalidVerifySwitch                                                = 118
	ErrorBadDriverLevel                                                     = 119
	ErrorCallNotImplemented                                                 = 120
	ErrorSemTimeout                                                         = 121
	ErrorInsufficientBuffer                                                 = 122
	ErrorInvalidName                                                        = 123
	ErrorInvalidLevel                                                       = 124
	ErrorNoVolumeLabel                                                      = 125
	ErrorModNotFound                                                        = 126
	ErrorProcNotFound                                                       = 127
	ErrorWaitNoChildren                                                     = 128
	ErrorChildNotComplete                                                   = 129
	ErrorDirectAccessHandle                                                 = 130
	ErrorNegativeSeek                                                       = 131
	ErrorSeekOnDevice                                                       = 132
	ErrorIsJoinTarget                                                       = 133
	ErrorIsJoined                                                           = 134
	ErrorIsSubsted                                                          = 135
	ErrorNotJoined                                                          = 136
	ErrorNotSubsted                                                         = 137
	ErrorJoinToJoin                                                         = 138
	ErrorSubstToSubst                                                       = 139
	ErrorJoinToSubst                                                        = 140
	ErrorSubstToJoin                                                        = 141
	ErrorBusyDrive                                                          = 142
	ErrorSameDrive                                                          = 143
	ErrorDirNotRoot                                                         = 144
	ErrorDirNotEmpty                                                        = 145
	ErrorIsSubstPath                                                        = 146
	ErrorIsJoinPath                                                         = 147
	ErrorPathBusy                                                           = 148
	ErrorIsSubstTarget                                                      = 149
	ErrorSystemTrace                                                        = 150
	ErrorInvalidEventCount                                                  = 151
	ErrorTooManyMuxwaiters                                                  = 152
	ErrorInvalidListFormat                                                  = 153
	ErrorLabelTooLong                                                       = 154
	ErrorTooManyTcbs                                                        = 155
	ErrorSignalRefused                                                      = 156
	ErrorDiscarded                                                          = 157
	ErrorNotLocked                                                          = 158
	ErrorBadThreadidAddr                                                    = 159
	ErrorBadArguments                                                       = 160
	ErrorBadPathname                                                        = 161
	ErrorSignalPending                                                      = 162
	ErrorMaxThrdsReached                                                    = 164
	ErrorLockFailed                                                         = 167
	ErrorBusy                                                               = 170
	ErrorCancelViolation                                                    = 173
	ErrorAtomicLocksNotSupported                                            = 174
	ErrorInvalidSegmentNumber                                               = 180
	ErrorInvalidOrdinal                                                     = 182
	ErrorAlreadyExists                                                      = 183
	ErrorInvalidFlagNumber                                                  = 186
	ErrorSemNotFound                                                        = 187
	ErrorInvalidStartingCodeseg                                             = 188
	ErrorInvalidStackseg                                                    = 189
	ErrorInvalidModuletype                                                  = 190
	ErrorInvalidExeSignature                                                = 191
	ErrorExeMarkedInvalid                                                   = 192
	ErrorBadExeFormat                                                       = 193
	ErrorIteratedDataExceeds64k                                             = 194
	ErrorInvalidMinallocsize                                                = 195
	ErrorDynlinkFromInvalidRing                                             = 196
	ErrorIoplNotEnabled                                                     = 197
	ErrorInvalidSegdpl                                                      = 198
	ErrorAutodatasegExceeds64k                                              = 199
	ErrorRing2SegMustBeMovable                                              = 200
	ErrorRelocChainXeedsSeglim                                              = 201
	ErrorInfloopInRelocChain                                                = 202
	ErrorEnvvarNotFound                                                     = 203
	ErrorNoSignalSent                                                       = 205
	ErrorFilenameExcedRange                                                 = 206
	ErrorRing2StackInUse                                                    = 207
	ErrorMetaExpansionTooLong                                               = 208
	ErrorInvalidSignalNumber                                                = 209
	ErrorThread1Inactive                                                    = 210
	ErrorLocked                                                             = 212
	ErrorTooManyModules                                                     = 214
	ErrorNestingNotAllowed                                                  = 215
	ErrorExeMachineTypeMismatch                                             = 216
	ErrorExeCannotModifySignedBinary                                        = 217
	ErrorExeCannotModifyStrongSignedBinary                                  = 218
	ErrorFileCheckedOut                                                     = 220
	ErrorCheckoutRequired                                                   = 221
	ErrorBadFileType                                                        = 222
	ErrorFileTooLarge                                                       = 223
	ErrorFormsAuthRequired                                                  = 224
	ErrorVirusInfected                                                      = 225
	ErrorVirusDeleted                                                       = 226
	ErrorPipeLocal                                                          = 229
	ErrorBadPipe                                                            = 230
	ErrorPipeBusy                                                           = 231
	ErrorNoData                                                             = 232
	ErrorPipeNotConnected                                                   = 233
	ErrorMoreData                                                           = 234
	ErrorVcDisconnected                                                     = 240
	ErrorInvalidEaName                                                      = 254
	ErrorEaListInconsistent                                                 = 255
	WaitTimeout                                                             = 258
	ErrorNoMoreItems                                                        = 259
	ErrorCannotCopy                                                         = 266
	ErrorDirectory                                                          = 267
	ErrorEasDidntFit                                                        = 275
	ErrorEaFileCorrupt                                                      = 276
	ErrorEaTableFull                                                        = 277
	ErrorInvalidEaHandle                                                    = 278
	ErrorEasNotSupported                                                    = 282
	ErrorNotOwner                                                           = 288
	ErrorTooManyPosts                                                       = 298
	ErrorPartialCopy                                                        = 299
	ErrorOplockNotGranted                                                   = 300
	ErrorInvalidOplockProtocol                                              = 301
	ErrorDiskTooFragmented                                                  = 302
	ErrorDeletePending                                                      = 303
	ErrorIncompatibleWithGlobalShortNameRegistrySetting                     = 304
	ErrorShortNamesNotEnabledOnVolume                                       = 305
	ErrorSecurityStreamIsInconsistent                                       = 306
	ErrorInvalidLockRange                                                   = 307
	ErrorImageSubsystemNotPresent                                           = 308
	ErrorNotificationGuidAlreadyDefined                                     = 309
	ErrorMrMidNotFound                                                      = 317
	ErrorScopeNotFound                                                      = 318
	ErrorFailNoactionReboot                                                 = 350
	ErrorFailShutdown                                                       = 351
	ErrorFailRestart                                                        = 352
	ErrorMaxSessionsReached                                                 = 353
	ErrorThreadModeAlreadyBackground                                        = 400
	ErrorThreadModeNotBackground                                            = 401
	ErrorProcessModeAlreadyBackground                                       = 402
	ErrorProcessModeNotBackground                                           = 403
	ErrorInvalidAddress                                                     = 487
	ErrorUserProfileLoad                                                    = 500
	ErrorArithmeticOverflow                                                 = 534
	ErrorPipeConnected                                                      = 535
	ErrorPipeListening                                                      = 536
	ErrorVerifierStop                                                       = 537
	ErrorAbiosError                                                         = 538
	ErrorWx86Warning                                                        = 539
	ErrorWx86Error                                                          = 540
	ErrorTimerNotCanceled                                                   = 541
	ErrorUnwind                                                             = 542
	ErrorBadStack                                                           = 543
	ErrorInvalidUnwindTarget                                                = 544
	ErrorInvalidPortAttributes                                              = 545
	ErrorPortMessageTooLong                                                 = 546
	ErrorInvalidQuotaLower                                                  = 547
	ErrorDeviceAlreadyAttached                                              = 548
	ErrorInstructionMisalignment                                            = 549
	ErrorProfilingNotStarted                                                = 550
	ErrorProfilingNotStopped                                                = 551
	ErrorCouldNotInterpret                                                  = 552
	ErrorProfilingAtLimit                                                   = 553
	ErrorCantWait                                                           = 554
	ErrorCantTerminateSelf                                                  = 555
	ErrorUnexpectedMmCreateErr                                              = 556
	ErrorUnexpectedMmMapError                                               = 557
	ErrorUnexpectedMmExtendErr                                              = 558
	ErrorBadFunctionTable                                                   = 559
	ErrorNoGuidTranslation                                                  = 560
	ErrorInvalidLdtSize                                                     = 561
	ErrorInvalidLdtOffset                                                   = 563
	ErrorInvalidLdtDescriptor                                               = 564
	ErrorTooManyThreads                                                     = 565
	ErrorThreadNotInProcess                                                 = 566
	ErrorPagefileQuotaExceeded                                              = 567
	ErrorLogonServerConflict                                                = 568
	ErrorSynchronizationRequired                                            = 569
	ErrorNetOpenFailed                                                      = 570
	ErrorIoPrivilegeFailed                                                  = 571
	ErrorControlCExit                                                       = 572
	ErrorMissingSystemfile                                                  = 573
	ErrorUnhandledException                                                 = 574
	ErrorAppInitFailure                                                     = 575
	ErrorPagefileCreateFailed                                               = 576
	ErrorInvalidImageHash                                                   = 577
	ErrorNoPagefile                                                         = 578
	ErrorIllegalFloatContext                                                = 579
	ErrorNoEventPair                                                        = 580
	ErrorDomainCtrlrConfigError                                             = 581
	ErrorIllegalCharacter                                                   = 582
	ErrorUndefinedCharacter                                                 = 583
	ErrorFloppyVolume                                                       = 584
	ErrorBiosFailedToConnectInterrupt                                       = 585
	ErrorBackupController                                                   = 586
	ErrorMutantLimitExceeded                                                = 587
	ErrorFsDriverRequired                                                   = 588
	ErrorCannotLoadRegistryFile                                             = 589
	ErrorDebugAttachFailed                                                  = 590
	ErrorSystemProcessTerminated                                            = 591
	ErrorDataNotAccepted                                                    = 592
	ErrorVdmHardError                                                       = 593
	ErrorDriverCancelTimeout                                                = 594
	ErrorReplyMessageMismatch                                               = 595
	ErrorLostWritebehindData                                                = 596
	ErrorClientServerParametersInvalid                                      = 597
	ErrorNotTinyStream                                                      = 598
	ErrorStackOverflowRead                                                  = 599
	ErrorConvertToLarge                                                     = 600
	ErrorFoundOutOfScope                                                    = 601
	ErrorAllocateBucket                                                     = 602
	ErrorMarshallOverflow                                                   = 603
	ErrorInvalidVariant                                                     = 604
	ErrorBadCompressionBuffer                                               = 605
	ErrorAuditFailed                                                        = 606
	ErrorTimerResolutionNotSet                                              = 607
	ErrorInsufficientLogonInfo                                              = 608
	ErrorBadDllEntrypoint                                                   = 609
	ErrorBadServiceEntrypoint                                               = 610
	ErrorIpAddressConflict1                                                 = 611
	ErrorIpAddressConflict2                                                 = 612
	ErrorRegistryQuotaLimit                                                 = 613
	ErrorNoCallbackActive                                                   = 614
	ErrorPwdTooShort                                                        = 615
	ErrorPwdTooRecent                                                       = 616
	ErrorPwdHistoryConflict                                                 = 617
	ErrorUnsupportedCompression                                             = 618
	ErrorInvalidHwProfile                                                   = 619
	ErrorInvalidPlugplayDevicePath                                          = 620
	ErrorQuotaListInconsistent                                              = 621
	ErrorEvaluationExpiration                                               = 622
	ErrorIllegalDllRelocation                                               = 623
	ErrorDllInitFailedLogoff                                                = 624
	ErrorValidateContinue                                                   = 625
	ErrorNoMoreMatches                                                      = 626
	ErrorRangeListConflict                                                  = 627
	ErrorServerSidMismatch                                                  = 628
	ErrorCantEnableDenyOnly                                                 = 629
	ErrorFloatMultipleFaults                                                = 630
	ErrorFloatMultipleTraps                                                 = 631
	ErrorNointerface                                                        = 632
	ErrorDriverFailedSleep                                                  = 633
	ErrorCorruptSystemFile                                                  = 634
	ErrorCommitmentMinimum                                                  = 635
	ErrorPnpRestartEnumeration                                              = 636
	ErrorSystemImageBadSignature                                            = 637
	ErrorPnpRebootRequired                                                  = 638
	ErrorInsufficientPower                                                  = 639
	ErrorMultipleFaultViolation                                             = 640
	ErrorSystemShutdown                                                     = 641
	ErrorPortNotSet                                                         = 642
	ErrorDsVersionCheckFailure                                              = 643
	ErrorRangeNotFound                                                      = 644
	ErrorNotSafeModeDriver                                                  = 646
	ErrorFailedDriverEntry                                                  = 647
	ErrorDeviceEnumerationError                                             = 648
	ErrorMountPointNotResolved                                              = 649
	ErrorInvalidDeviceObjectParameter                                       = 650
	ErrorMcaOccured                                                         = 651
	ErrorDriverDatabaseError                                                = 652
	ErrorSystemHiveTooLarge                                                 = 653
	ErrorDriverFailedPriorUnload                                            = 654
	ErrorVolsnapPrepareHibernate                                            = 655
	ErrorHibernationFailure                                                 = 656
	ErrorFileSystemLimitation                                               = 665
	ErrorAssertionFailure                                                   = 668
	ErrorAcpiError                                                          = 669
	ErrorWowAssertion                                                       = 670
	ErrorPnpBadMpsTable                                                     = 671
	ErrorPnpTranslationFailed                                               = 672
	ErrorPnpIrqTranslationFailed                                            = 673
	ErrorPnpInvalidId                                                       = 674
	ErrorWakeSystemDebugger                                                 = 675
	ErrorHandlesClosed                                                      = 676
	ErrorExtraneousInformation                                              = 677
	ErrorRxactCommitNecessary                                               = 678
	ErrorMediaCheck                                                         = 679
	ErrorGuidSubstitutionMade                                               = 680
	ErrorStoppedOnSymlink                                                   = 681
	ErrorLongjump                                                           = 682
	ErrorPlugplayQueryVetoed                                                = 683
	ErrorUnwindConsolidate                                                  = 684
	ErrorRegistryHiveRecovered                                              = 685
	ErrorDllMightBeInsecure                                                 = 686
	ErrorDllMightBeIncompatible                                             = 687
	ErrorDbgExceptionNotHandled                                             = 688
	ErrorDbgReplyLater                                                      = 689
	ErrorDbgUnableToProvideHandle                                           = 690
	ErrorDbgTerminateThread                                                 = 691
	ErrorDbgTerminateProcess                                                = 692
	ErrorDbgControlC                                                        = 693
	ErrorDbgPrintexceptionC                                                 = 694
	ErrorDbgRipexception                                                    = 695
	ErrorDbgControlBreak                                                    = 696
	ErrorDbgCommandException                                                = 697
	ErrorObjectNameExists                                                   = 698
	ErrorThreadWasSuspended                                                 = 699
	ErrorImageNotAtBase                                                     = 700
	ErrorRxactStateCreated                                                  = 701
	ErrorSegmentNotification                                                = 702
	ErrorBadCurrentDirectory                                                = 703
	ErrorFtReadRecoveryFromBackup                                           = 704
	ErrorFtWriteRecovery                                                    = 705
	ErrorImageMachineTypeMismatch                                           = 706
	ErrorReceivePartial                                                     = 707
	ErrorReceiveExpedited                                                   = 708
	ErrorReceivePartialExpedited                                            = 709
	ErrorEventDone                                                          = 710
	ErrorEventPending                                                       = 711
	ErrorCheckingFileSystem                                                 = 712
	ErrorFatalAppExit                                                       = 713
	ErrorPredefinedHandle                                                   = 714
	ErrorWasUnlocked                                                        = 715
	ErrorServiceNotification                                                = 716
	ErrorWasLocked                                                          = 717
	ErrorLogHardError                                                       = 718
	ErrorAlreadyWin32                                                       = 719
	ErrorImageMachineTypeMismatchExe                                        = 720
	ErrorNoYieldPerformed                                                   = 721
	ErrorTimerResumeIgnored                                                 = 722
	ErrorArbitrationUnhandled                                               = 723
	ErrorCardbusNotSupported                                                = 724
	ErrorMpProcessorMismatch                                                = 725
	ErrorHibernated                                                         = 726
	ErrorResumeHibernation                                                  = 727
	ErrorFirmwareUpdated                                                    = 728
	ErrorDriversLeakingLockedPages                                          = 729
	ErrorWakeSystem                                                         = 730
	ErrorWait1                                                              = 731
	ErrorWait2                                                              = 732
	ErrorWait3                                                              = 733
	ErrorWait63                                                             = 734
	ErrorAbandonedWait0                                                     = 735
	ErrorAbandonedWait63                                                    = 736
	ErrorUserApc                                                            = 737
	ErrorKernelApc                                                          = 738
	ErrorAlerted                                                            = 739
	ErrorElevationRequired                                                  = 740
	ErrorReparse                                                            = 741
	ErrorOplockBreakInProgress                                              = 742
	ErrorVolumeMounted                                                      = 743
	ErrorRxactCommitted                                                     = 744
	ErrorNotifyCleanup                                                      = 745
	ErrorPrimaryTransportConnectFailed                                      = 746
	ErrorPageFaultTransition                                                = 747
	ErrorPageFaultDemandZero                                                = 748
	ErrorPageFaultCopyOnWrite                                               = 749
	ErrorPageFaultGuardPage                                                 = 750
	ErrorPageFaultPagingFile                                                = 751
	ErrorCachePageLocked                                                    = 752
	ErrorCrashDump                                                          = 753
	ErrorBufferAllZeros                                                     = 754
	ErrorReparseObject                                                      = 755
	ErrorResourceRequirementsChanged                                        = 756
	ErrorTranslationComplete                                                = 757
	ErrorNothingToTerminate                                                 = 758
	ErrorProcessNotInJob                                                    = 759
	ErrorProcessInJob                                                       = 760
	ErrorVolsnapHibernateReady                                              = 761
	ErrorFsfilterOpCompletedSuccessfully                                    = 762
	ErrorInterruptVectorAlreadyConnected                                    = 763
	ErrorInterruptStillConnected                                            = 764
	ErrorWaitForOplock                                                      = 765
	ErrorDbgExceptionHandled                                                = 766
	ErrorDbgContinue                                                        = 767
	ErrorCallbackPopStack                                                   = 768
	ErrorCompressionDisabled                                                = 769
	ErrorCantfetchbackwards                                                 = 770
	ErrorCantscrollbackwards                                                = 771
	ErrorRowsnotreleased                                                    = 772
	ErrorBadAccessorFlags                                                   = 773
	ErrorErrorsEncountered                                                  = 774
	ErrorNotCapable                                                         = 775
	ErrorRequestOutOfSequence                                               = 776
	ErrorVersionParseError                                                  = 777
	ErrorBadstartposition                                                   = 778
	ErrorMemoryHardware                                                     = 779
	ErrorDiskRepairDisabled                                                 = 780
	ErrorInsufficientResourceForSpecifiedSharedSectionSize                  = 781
	ErrorSystemPowerstateTransition                                         = 782
	ErrorSystemPowerstateComplexTransition                                  = 783
	ErrorMcaException                                                       = 784
	ErrorAccessAuditByPolicy                                                = 785
	ErrorAccessDisabledNoSaferUiByPolicy                                    = 786
	ErrorAbandonHiberfile                                                   = 787
	ErrorLostWritebehindDataNetworkDisconnected                             = 788
	ErrorLostWritebehindDataNetworkServerError                              = 789
	ErrorLostWritebehindDataLocalDiskError                                  = 790
	ErrorBadMcfgTable                                                       = 791
	ErrorOplockSwitchedToNewHandle                                          = 800
	ErrorCannotGrantRequestedOplock                                         = 801
	ErrorCannotBreakOplock                                                  = 802
	ErrorOplockHandleClosed                                                 = 803
	ErrorNoAceCondition                                                     = 804
	ErrorInvalidAceCondition                                                = 805
	ErrorEaAccessDenied                                                     = 994
	ErrorOperationAborted                                                   = 995
	ErrorIoIncomplete                                                       = 996
	ErrorIoPending                                                          = 997
	ErrorNoaccess                                                           = 998
	ErrorSwaperror                                                          = 999
	ErrorStackOverflow                                                      = 1001
	ErrorInvalidMessage                                                     = 1002
	ErrorCanNotComplete                                                     = 1003
	ErrorInvalidFlags                                                       = 1004
	ErrorUnrecognizedVolume                                                 = 1005
	ErrorFileInvalid                                                        = 1006
	ErrorFullscreenMode                                                     = 1007
	ErrorNoToken                                                            = 1008
	ErrorBaddb                                                              = 1009
	ErrorBadkey                                                             = 1010
	ErrorCantopen                                                           = 1011
	ErrorCantread                                                           = 1012
	ErrorCantwrite                                                          = 1013
	ErrorRegistryRecovered                                                  = 1014
	ErrorRegistryCorrupt                                                    = 1015
	ErrorRegistryIoFailed                                                   = 1016
	ErrorNotRegistryFile                                                    = 1017
	ErrorKeyDeleted                                                         = 1018
	ErrorNoLogSpace                                                         = 1019
	ErrorKeyHasChildren                                                     = 1020
	ErrorChildMustBeVolatile                                                = 1021
	ErrorNotifyEnumDir                                                      = 1022
	ErrorDependentServicesRunning                                           = 1051
	ErrorInvalidServiceControl                                              = 1052
	ErrorServiceRequestTimeout                                              = 1053
	ErrorServiceNoThread                                                    = 1054
	ErrorServiceDatabaseLocked                                              = 1055
	ErrorServiceAlreadyRunning                                              = 1056
	ErrorInvalidServiceAccount                                              = 1057
	ErrorServiceDisabled                                                    = 1058
	ErrorCircularDependency                                                 = 1059
	ErrorServiceDoesNotExist                                                = 1060
	ErrorServiceCannotAcceptCtrl                                            = 1061
	ErrorServiceNotActive                                                   = 1062
	ErrorFailedServiceControllerConnect                                     = 1063
	ErrorExceptionInService                                                 = 1064
	ErrorDatabaseDoesNotExist                                               = 1065
	ErrorServiceSpecificError                                               = 1066
	ErrorProcessAborted                                                     = 1067
	ErrorServiceDependencyFail                                              = 1068
	ErrorServiceLogonFailed                                                 = 1069
	ErrorServiceStartHang                                                   = 1070
	ErrorInvalidServiceLock                                                 = 1071
	ErrorServiceMarkedForDelete                                             = 1072
	ErrorServiceExists                                                      = 1073
	ErrorAlreadyRunningLkg                                                  = 1074
	ErrorServiceDependencyDeleted                                           = 1075
	ErrorBootAlreadyAccepted                                                = 1076
	ErrorServiceNeverStarted                                                = 1077
	ErrorDuplicateServiceName                                               = 1078
	ErrorDifferentServiceAccount                                            = 1079
	ErrorCannotDetectDriverFailure                                          = 1080
	ErrorCannotDetectProcessAbort                                           = 1081
	ErrorNoRecoveryProgram                                                  = 1082
	ErrorServiceNotInExe                                                    = 1083
	ErrorNotSafebootService                                                 = 1084
	ErrorEndOfMedia                                                         = 1100
	ErrorFilemarkDetected                                                   = 1101
	ErrorBeginningOfMedia                                                   = 1102
	ErrorSetmarkDetected                                                    = 1103
	ErrorNoDataDetected                                                     = 1104
	ErrorPartitionFailure                                                   = 1105
	ErrorInvalidBlockLength                                                 = 1106
	ErrorDeviceNotPartitioned                                               = 1107
	ErrorUnableToLockMedia                                                  = 1108
	ErrorUnableToUnloadMedia                                                = 1109
	ErrorMediaChanged                                                       = 1110
	ErrorBusReset                                                           = 1111
	ErrorNoMediaInDrive                                                     = 1112
	ErrorNoUnicodeTranslation                                               = 1113
	ErrorDllInitFailed                                                      = 1114
	ErrorShutdownInProgress                                                 = 1115
	ErrorNoShutdownInProgress                                               = 1116
	ErrorIoDevice                                                           = 1117
	ErrorSerialNoDevice                                                     = 1118
	ErrorIrqBusy                                                            = 1119
	ErrorMoreWrites                                                         = 1120
	ErrorCounterTimeout                                                     = 1121
	ErrorFloppyIdMarkNotFound                                               = 1122
	ErrorFloppyWrongCylinder                                                = 1123
	ErrorFloppyUnknownError                                                 = 1124
	ErrorFloppyBadRegisters                                                 = 1125
	ErrorDiskRecalibrateFailed                                              = 1126
	ErrorDiskOperationFailed                                                = 1127
	ErrorDiskResetFailed                                                    = 1128
	ErrorEomOverflow                                                        = 1129
	ErrorNotEnoughServerMemory                                              = 1130
	ErrorPossibleDeadlock                                                   = 1131
	ErrorMappedAlignment                                                    = 1132
	ErrorSetPowerStateVetoed                                                = 1140
	ErrorSetPowerStateFailed                                                = 1141
	ErrorTooManyLinks                                                       = 1142
	ErrorOldWinVersion                                                      = 1150
	ErrorAppWrongOs                                                         = 1151
	ErrorSingleInstanceApp                                                  = 1152
	ErrorRmodeApp                                                           = 1153
	ErrorInvalidDll                                                         = 1154
	ErrorNoAssociation                                                      = 1155
	ErrorDdeFail                                                            = 1156
	ErrorDllNotFound                                                        = 1157
	ErrorNoMoreUserHandles                                                  = 1158
	ErrorMessageSyncOnly                                                    = 1159
	ErrorSourceElementEmpty                                                 = 1160
	ErrorDestinationElementFull                                             = 1161
	ErrorIllegalElementAddress                                              = 1162
	ErrorMagazineNotPresent                                                 = 1163
	ErrorDeviceReinitializationNeeded                                       = 1164
	ErrorDeviceRequiresCleaning                                             = 1165
	ErrorDeviceDoorOpen                                                     = 1166
	ErrorDeviceNotConnected                                                 = 1167
	ErrorNotFound                                                           = 1168
	ErrorNoMatch                                                            = 1169
	ErrorSetNotFound                                                        = 1170
	ErrorPointNotFound                                                      = 1171
	ErrorNoTrackingService                                                  = 1172
	ErrorNoVolumeId                                                         = 1173
	ErrorUnableToRemoveReplaced                                             = 1175
	ErrorUnableToMoveReplacement                                            = 1176
	ErrorUnableToMoveReplacement2                                           = 1177
	ErrorJournalDeleteInProgress                                            = 1178
	ErrorJournalNotActive                                                   = 1179
	ErrorPotentialFileFound                                                 = 1180
	ErrorJournalEntryDeleted                                                = 1181
	ErrorShutdownIsScheduled                                                = 1190
	ErrorShutdownUsersLoggedOn                                              = 1191
	ErrorBadDevice                                                          = 1200
	ErrorConnectionUnavail                                                  = 1201
	ErrorDeviceAlreadyRemembered                                            = 1202
	ErrorNoNetOrBadPath                                                     = 1203
	ErrorBadProvider                                                        = 1204
	ErrorCannotOpenProfile                                                  = 1205
	ErrorBadProfile                                                         = 1206
	ErrorNotContainer                                                       = 1207
	ErrorExtendedError                                                      = 1208
	ErrorInvalidGroupname                                                   = 1209
	ErrorInvalidComputername                                                = 1210
	ErrorInvalidEventname                                                   = 1211
	ErrorInvalidDomainname                                                  = 1212
	ErrorInvalidServicename                                                 = 1213
	ErrorInvalidNetname                                                     = 1214
	ErrorInvalidSharename                                                   = 1215
	ErrorInvalidPasswordname                                                = 1216
	ErrorInvalidMessagename                                                 = 1217
	ErrorInvalidMessagedest                                                 = 1218
	ErrorSessionCredentialConflict                                          = 1219
	ErrorRemoteSessionLimitExceeded                                         = 1220
	ErrorDupDomainname                                                      = 1221
	ErrorNoNetwork                                                          = 1222
	ErrorCancelled                                                          = 1223
	ErrorUserMappedFile                                                     = 1224
	ErrorConnectionRefused                                                  = 1225
	ErrorGracefulDisconnect                                                 = 1226
	ErrorAddressAlreadyAssociated                                           = 1227
	ErrorAddressNotAssociated                                               = 1228
	ErrorConnectionInvalid                                                  = 1229
	ErrorConnectionActive                                                   = 1230
	ErrorNetworkUnreachable                                                 = 1231
	ErrorHostUnreachable                                                    = 1232
	ErrorProtocolUnreachable                                                = 1233
	ErrorPortUnreachable                                                    = 1234
	ErrorRequestAborted                                                     = 1235
	ErrorConnectionAborted                                                  = 1236
	ErrorRetry                                                              = 1237
	ErrorConnectionCountLimit                                               = 1238
	ErrorLoginTimeRestriction                                               = 1239
	ErrorLoginWkstaRestriction                                              = 1240
	ErrorIncorrectAddress                                                   = 1241
	ErrorAlreadyRegistered                                                  = 1242
	ErrorServiceNotFound                                                    = 1243
	ErrorNotAuthenticated                                                   = 1244
	ErrorNotLoggedOn                                                        = 1245
	ErrorContinue                                                           = 1246
	ErrorAlreadyInitialized                                                 = 1247
	ErrorNoMoreDevices                                                      = 1248
	ErrorNoSuchSite                                                         = 1249
	ErrorDomainControllerExists                                             = 1250
	ErrorOnlyIfConnected                                                    = 1251
	ErrorOverrideNochanges                                                  = 1252
	ErrorBadUserProfile                                                     = 1253
	ErrorNotSupportedOnSbs                                                  = 1254
	ErrorServerShutdownInProgress                                           = 1255
	ErrorHostDown                                                           = 1256
	ErrorNonAccountSid                                                      = 1257
	ErrorNonDomainSid                                                       = 1258
	ErrorApphelpBlock                                                       = 1259
	ErrorAccessDisabledByPolicy                                             = 1260
	ErrorRegNatConsumption                                                  = 1261
	ErrorCscshareOffline                                                    = 1262
	ErrorPkinitFailure                                                      = 1263
	ErrorSmartcardSubsystemFailure                                          = 1264
	ErrorDowngradeDetected                                                  = 1265
	ErrorMachineLocked                                                      = 1271
	ErrorCallbackSuppliedInvalidData                                        = 1273
	ErrorSyncForegroundRefreshRequired                                      = 1274
	ErrorDriverBlocked                                                      = 1275
	ErrorInvalidImportOfNonDll                                              = 1276
	ErrorAccessDisabledWebblade                                             = 1277
	ErrorAccessDisabledWebbladeTamper                                       = 1278
	ErrorRecoveryFailure                                                    = 1279
	ErrorAlreadyFiber                                                       = 1280
	ErrorAlreadyThread                                                      = 1281
	ErrorStackBufferOverrun                                                 = 1282
	ErrorParameterQuotaExceeded                                             = 1283
	ErrorDebuggerInactive                                                   = 1284
	ErrorDelayLoadFailed                                                    = 1285
	ErrorVdmDisallowed                                                      = 1286
	ErrorUnidentifiedError                                                  = 1287
	ErrorInvalidCruntimeParameter                                           = 1288
	ErrorBeyondVdl                                                          = 1289
	ErrorIncompatibleServiceSidType                                         = 1290
	ErrorDriverProcessTerminated                                            = 1291
	ErrorImplementationLimit                                                = 1292
	ErrorProcessIsProtected                                                 = 1293
	ErrorServiceNotifyClientLagging                                         = 1294
	ErrorDiskQuotaExceeded                                                  = 1295
	ErrorContentBlocked                                                     = 1296
	ErrorIncompatibleServicePrivilege                                       = 1297
	ErrorAppHang                                                            = 1298
	ErrorInvalidLabel                                                       = 1299
	ErrorNotAllAssigned                                                     = 1300
	ErrorSomeNotMapped                                                      = 1301
	ErrorNoQuotasForAccount                                                 = 1302
	ErrorLocalUserSessionKey                                                = 1303
	ErrorNullLmPassword                                                     = 1304
	ErrorUnknownRevision                                                    = 1305
	ErrorRevisionMismatch                                                   = 1306
	ErrorInvalidOwner                                                       = 1307
	ErrorInvalidPrimaryGroup                                                = 1308
	ErrorNoImpersonationToken                                               = 1309
	ErrorCantDisableMandatory                                               = 1310
	ErrorNoLogonServers                                                     = 1311
	ErrorNoSuchLogonSession                                                 = 1312
	ErrorNoSuchPrivilege                                                    = 1313
	ErrorPrivilegeNotHeld                                                   = 1314
	ErrorInvalidAccountName                                                 = 1315
	ErrorUserExists                                                         = 1316
	ErrorNoSuchUser                                                         = 1317
	ErrorGroupExists                                                        = 1318
	ErrorNoSuchGroup                                                        = 1319
	ErrorMemberInGroup                                                      = 1320
	ErrorMemberNotInGroup                                                   = 1321
	ErrorLastAdmin                                                          = 1322
	ErrorWrongPassword                                                      = 1323
	ErrorIllFormedPassword                                                  = 1324
	ErrorPasswordRestriction                                                = 1325
	ErrorLogonFailure                                                       = 1326
	ErrorAccountRestriction                                                 = 1327
	ErrorInvalidLogonHours                                                  = 1328
	ErrorInvalidWorkstation                                                 = 1329
	ErrorPasswordExpired                                                    = 1330
	ErrorAccountDisabled                                                    = 1331
	ErrorNoneMapped                                                         = 1332
	ErrorTooManyLuidsRequested                                              = 1333
	ErrorLuidsExhausted                                                     = 1334
	ErrorInvalidSubAuthority                                                = 1335
	ErrorInvalidAcl                                                         = 1336
	ErrorInvalidSid                                                         = 1337
	ErrorInvalidSecurityDescr                                               = 1338
	ErrorBadInheritanceAcl                                                  = 1340
	ErrorServerDisabled                                                     = 1341
	ErrorServerNotDisabled                                                  = 1342
	ErrorInvalidIdAuthority                                                 = 1343
	ErrorAllottedSpaceExceeded                                              = 1344
	ErrorInvalidGroupAttributes                                             = 1345
	ErrorBadImpersonationLevel                                              = 1346
	ErrorCantOpenAnonymous                                                  = 1347
	ErrorBadValidationClass                                                 = 1348
	ErrorBadTokenType                                                       = 1349
	ErrorNoSecurityOnObject                                                 = 1350
	ErrorCantAccessDomainInfo                                               = 1351
	ErrorInvalidServerState                                                 = 1352
	ErrorInvalidDomainState                                                 = 1353
	ErrorInvalidDomainRole                                                  = 1354
	ErrorNoSuchDomain                                                       = 1355
	ErrorDomainExists                                                       = 1356
	ErrorDomainLimitExceeded                                                = 1357
	ErrorInternalDbCorruption                                               = 1358
	ErrorInternalError                                                      = 1359
	ErrorGenericNotMapped                                                   = 1360
	ErrorBadDescriptorFormat                                                = 1361
	ErrorNotLogonProcess                                                    = 1362
	ErrorLogonSessionExists                                                 = 1363
	ErrorNoSuchPackage                                                      = 1364
	ErrorBadLogonSessionState                                               = 1365
	ErrorLogonSessionCollision                                              = 1366
	ErrorInvalidLogonType                                                   = 1367
	ErrorCannotImpersonate                                                  = 1368
	ErrorRxactInvalidState                                                  = 1369
	ErrorRxactCommitFailure                                                 = 1370
	ErrorSpecialAccount                                                     = 1371
	ErrorSpecialGroup                                                       = 1372
	ErrorSpecialUser                                                        = 1373
	ErrorMembersPrimaryGroup                                                = 1374
	ErrorTokenAlreadyInUse                                                  = 1375
	ErrorNoSuchAlias                                                        = 1376
	ErrorMemberNotInAlias                                                   = 1377
	ErrorMemberInAlias                                                      = 1378
	ErrorAliasExists                                                        = 1379
	ErrorLogonNotGranted                                                    = 1380
	ErrorTooManySecrets                                                     = 1381
	ErrorSecretTooLong                                                      = 1382
	ErrorInternalDbError                                                    = 1383
	ErrorTooManyContextIds                                                  = 1384
	ErrorLogonTypeNotGranted                                                = 1385
	ErrorNtCrossEncryptionRequired                                          = 1386
	ErrorNoSuchMember                                                       = 1387
	ErrorInvalidMember                                                      = 1388
	ErrorTooManySids                                                        = 1389
	ErrorLmCrossEncryptionRequired                                          = 1390
	ErrorNoInheritance                                                      = 1391
	ErrorFileCorrupt                                                        = 1392
	ErrorDiskCorrupt                                                        = 1393
	ErrorNoUserSessionKey                                                   = 1394
	ErrorLicenseQuotaExceeded                                               = 1395
	ErrorWrongTargetName                                                    = 1396
	ErrorMutualAuthFailed                                                   = 1397
	ErrorTimeSkew                                                           = 1398
	ErrorCurrentDomainNotAllowed                                            = 1399
	ErrorInvalidWindowHandle                                                = 1400
	ErrorInvalidMenuHandle                                                  = 1401
	ErrorInvalidCursorHandle                                                = 1402
	ErrorInvalidAccelHandle                                                 = 1403
	ErrorInvalidHookHandle                                                  = 1404
	ErrorInvalidDwpHandle                                                   = 1405
	ErrorTlwWithWschild                                                     = 1406
	ErrorCannotFindWndClass                                                 = 1407
	ErrorWindowOfOtherThread                                                = 1408
	ErrorHotkeyAlreadyRegistered                                            = 1409
	ErrorClassAlreadyExists                                                 = 1410
	ErrorClassDoesNotExist                                                  = 1411
	ErrorClassHasWindows                                                    = 1412
	ErrorInvalidIndex                                                       = 1413
	ErrorInvalidIconHandle                                                  = 1414
	ErrorPrivateDialogIndex                                                 = 1415
	ErrorListboxIdNotFound                                                  = 1416
	ErrorNoWildcardCharacters                                               = 1417
	ErrorClipboardNotOpen                                                   = 1418
	ErrorHotkeyNotRegistered                                                = 1419
	ErrorWindowNotDialog                                                    = 1420
	ErrorControlIdNotFound                                                  = 1421
	ErrorInvalidComboboxMessage                                             = 1422
	ErrorWindowNotCombobox                                                  = 1423
	ErrorInvalidEditHeight                                                  = 1424
	ErrorDcNotFound                                                         = 1425
	ErrorInvalidHookFilter                                                  = 1426
	ErrorInvalidFilterProc                                                  = 1427
	ErrorHookNeedsHmod                                                      = 1428
	ErrorGlobalOnlyHook                                                     = 1429
	ErrorJournalHookSet                                                     = 1430
	ErrorHookNotInstalled                                                   = 1431
	ErrorInvalidLbMessage                                                   = 1432
	ErrorSetcountOnBadLb                                                    = 1433
	ErrorLbWithoutTabstops                                                  = 1434
	ErrorDestroyObjectOfOtherThread                                         = 1435
	ErrorChildWindowMenu                                                    = 1436
	ErrorNoSystemMenu                                                       = 1437
	ErrorInvalidMsgboxStyle                                                 = 1438
	ErrorInvalidSpiValue                                                    = 1439
	ErrorScreenAlreadyLocked                                                = 1440
	ErrorHwndsHaveDiffParent                                                = 1441
	ErrorNotChildWindow                                                     = 1442
	ErrorInvalidGwCommand                                                   = 1443
	ErrorInvalidThreadId                                                    = 1444
	ErrorNonMdichildWindow                                                  = 1445
	ErrorPopupAlreadyActive                                                 = 1446
	ErrorNoScrollbars                                                       = 1447
	ErrorInvalidScrollbarRange                                              = 1448
	ErrorInvalidShowwinCommand                                              = 1449
	ErrorNoSystemResources                                                  = 1450
	ErrorNonpagedSystemResources                                            = 1451
	ErrorPagedSystemResources                                               = 1452
	ErrorWorkingSetQuota                                                    = 1453
	ErrorPagefileQuota                                                      = 1454
	ErrorCommitmentLimit                                                    = 1455
	ErrorMenuItemNotFound                                                   = 1456
	ErrorInvalidKeyboardHandle                                              = 1457
	ErrorHookTypeNotAllowed                                                 = 1458
	ErrorRequiresInteractiveWindowstation                                   = 1459
	ErrorTimeout                                                            = 1460
	ErrorInvalidMonitorHandle                                               = 1461
	ErrorIncorrectSize                                                      = 1462
	ErrorSymlinkClassDisabled                                               = 1463
	ErrorSymlinkNotSupported                                                = 1464
	ErrorXmlParseError                                                      = 1465
	ErrorXmldsigError                                                       = 1466
	ErrorRestartApplication                                                 = 1467
	ErrorWrongCompartment                                                   = 1468
	ErrorAuthipFailure                                                      = 1469
	ErrorNoNvramResources                                                   = 1470
	ErrorEventlogFileCorrupt                                                = 1500
	ErrorEventlogCantStart                                                  = 1501
	ErrorLogFileFull                                                        = 1502
	ErrorEventlogFileChanged                                                = 1503
	ErrorInvalidTaskName                                                    = 1550
	ErrorInvalidTaskIndex                                                   = 1551
	ErrorThreadAlreadyInTask                                                = 1552
	ErrorInstallServiceFailure                                              = 1601
	ErrorInstallUserexit                                                    = 1602
	ErrorInstallFailure                                                     = 1603
	ErrorInstallSuspend                                                     = 1604
	ErrorUnknownProduct                                                     = 1605
	ErrorUnknownFeature                                                     = 1606
	ErrorUnknownComponent                                                   = 1607
	ErrorUnknownProperty                                                    = 1608
	ErrorInvalidHandleState                                                 = 1609
	ErrorBadConfiguration                                                   = 1610
	ErrorIndexAbsent                                                        = 1611
	ErrorInstallSourceAbsent                                                = 1612
	ErrorInstallPackageVersion                                              = 1613
	ErrorProductUninstalled                                                 = 1614
	ErrorBadQuerySyntax                                                     = 1615
	ErrorInvalidField                                                       = 1616
	ErrorDeviceRemoved                                                      = 1617
	ErrorInstallAlreadyRunning                                              = 1618
	ErrorInstallPackageOpenFailed                                           = 1619
	ErrorInstallPackageInvalid                                              = 1620
	ErrorInstallUiFailure                                                   = 1621
	ErrorInstallLogFailure                                                  = 1622
	ErrorInstallLanguageUnsupported                                         = 1623
	ErrorInstallTransformFailure                                            = 1624
	ErrorInstallPackageRejected                                             = 1625
	ErrorFunctionNotCalled                                                  = 1626
	ErrorFunctionFailed                                                     = 1627
	ErrorInvalidTable                                                       = 1628
	ErrorDatatypeMismatch                                                   = 1629
	ErrorUnsupportedType                                                    = 1630
	ErrorCreateFailed                                                       = 1631
	ErrorInstallTempUnwritable                                              = 1632
	ErrorInstallPlatformUnsupported                                         = 1633
	ErrorInstallNotused                                                     = 1634
	ErrorPatchPackageOpenFailed                                             = 1635
	ErrorPatchPackageInvalid                                                = 1636
	ErrorPatchPackageUnsupported                                            = 1637
	ErrorProductVersion                                                     = 1638
	ErrorInvalidCommandLine                                                 = 1639
	ErrorInstallRemoteDisallowed                                            = 1640
	ErrorSuccessRebootInitiated                                             = 1641
	ErrorPatchTargetNotFound                                                = 1642
	ErrorPatchPackageRejected                                               = 1643
	ErrorInstallTransformRejected                                           = 1644
	ErrorInstallRemoteProhibited                                            = 1645
	ErrorPatchRemovalUnsupported                                            = 1646
	ErrorUnknownPatch                                                       = 1647
	ErrorPatchNoSequence                                                    = 1648
	ErrorPatchRemovalDisallowed                                             = 1649
	ErrorInvalidPatchXml                                                    = 1650
	ErrorPatchManagedAdvertisedProduct                                      = 1651
	ErrorInstallServiceSafeboot                                             = 1652
	ErrorFailFastException                                                  = 1653
	RpcSInvalidStringBinding                                                = 1700
	RpcSWrongKindOfBinding                                                  = 1701
	RpcSInvalidBinding                                                      = 1702
	RpcSProtseqNotSupported                                                 = 1703
	RpcSInvalidRpcProtseq                                                   = 1704
	RpcSInvalidStringUuid                                                   = 1705
	RpcSInvalidEndpointFormat                                               = 1706
	RpcSInvalidNetAddr                                                      = 1707
	RpcSNoEndpointFound                                                     = 1708
	RpcSInvalidTimeout                                                      = 1709
	RpcSObjectNotFound                                                      = 1710
	RpcSAlreadyRegistered                                                   = 1711
	RpcSTypeAlreadyRegistered                                               = 1712
	RpcSAlreadyListening                                                    = 1713
	RpcSNoProtseqsRegistered                                                = 1714
	RpcSNotListening                                                        = 1715
	RpcSUnknownMgrType                                                      = 1716
	RpcSUnknownIf                                                           = 1717
	RpcSNoBindings                                                          = 1718
	RpcSNoProtseqs                                                          = 1719
	RpcSCantCreateEndpoint                                                  = 1720
	RpcSOutOfResources                                                      = 1721
	RpcSServerUnavailable                                                   = 1722
	RpcSServerTooBusy                                                       = 1723
	RpcSInvalidNetworkOptions                                               = 1724
	RpcSNoCallActive                                                        = 1725
	RpcSCallFailed                                                          = 1726
	RpcSCallFailedDne                                                       = 1727
	RpcSProtocolError                                                       = 1728
	RpcSProxyAccessDenied                                                   = 1729
	RpcSUnsupportedTransSyn                                                 = 1730
	RpcSUnsupportedType                                                     = 1732
	RpcSInvalidTag                                                          = 1733
	RpcSInvalidBound                                                        = 1734
	RpcSNoEntryName                                                         = 1735
	RpcSInvalidNameSyntax                                                   = 1736
	RpcSUnsupportedNameSyntax                                               = 1737
	RpcSUuidNoAddress                                                       = 1739
	RpcSDuplicateEndpoint                                                   = 1740
	RpcSUnknownAuthnType                                                    = 1741
	RpcSMaxCallsTooSmall                                                    = 1742
	RpcSStringTooLong                                                       = 1743
	RpcSProtseqNotFound                                                     = 1744
	RpcSProcnumOutOfRange                                                   = 1745
	RpcSBindingHasNoAuth                                                    = 1746
	RpcSUnknownAuthnService                                                 = 1747
	RpcSUnknownAuthnLevel                                                   = 1748
	RpcSInvalidAuthIdentity                                                 = 1749
	RpcSUnknownAuthzService                                                 = 1750
	EptSInvalidEntry                                                        = 1751
	EptSCantPerformOp                                                       = 1752
	EptSNotRegistered                                                       = 1753
	RpcSNothingToExport                                                     = 1754
	RpcSIncompleteName                                                      = 1755
	RpcSInvalidVersOption                                                   = 1756
	RpcSNoMoreMembers                                                       = 1757
	RpcSNotAllObjsUnexported                                                = 1758
	RpcSInterfaceNotFound                                                   = 1759
	RpcSEntryAlreadyExists                                                  = 1760
	RpcSEntryNotFound                                                       = 1761
	RpcSNameServiceUnavailable                                              = 1762
	RpcSInvalidNafId                                                        = 1763
	RpcSCannotSupport                                                       = 1764
	RpcSNoContextAvailable                                                  = 1765
	RpcSInternalError                                                       = 1766
	RpcSZeroDivide                                                          = 1767
	RpcSAddressError                                                        = 1768
	RpcSFpDivZero                                                           = 1769
	RpcSFpUnderflow                                                         = 1770
	RpcSFpOverflow                                                          = 1771
	RpcXNoMoreEntries                                                       = 1772
	RpcXSsCharTransOpenFail                                                 = 1773
	RpcXSsCharTransShortFile                                                = 1774
	RpcXSsInNullContext                                                     = 1775
	RpcXSsContextDamaged                                                    = 1777
	RpcXSsHandlesMismatch                                                   = 1778
	RpcXSsCannotGetCallHandle                                               = 1779
	RpcXNullRefPointer                                                      = 1780
	RpcXEnumValueOutOfRange                                                 = 1781
	RpcXByteCountTooSmall                                                   = 1782
	RpcXBadStubData                                                         = 1783
	ErrorInvalidUserBuffer                                                  = 1784
	ErrorUnrecognizedMedia                                                  = 1785
	ErrorNoTrustLsaSecret                                                   = 1786
	ErrorNoTrustSamAccount                                                  = 1787
	ErrorTrustedDomainFailure                                               = 1788
	ErrorTrustedRelationshipFailure                                         = 1789
	ErrorTrustFailure                                                       = 1790
	RpcSCallInProgress                                                      = 1791
	ErrorNetlogonNotStarted                                                 = 1792
	ErrorAccountExpired                                                     = 1793
	ErrorRedirectorHasOpenHandles                                           = 1794
	ErrorPrinterDriverAlreadyInstalled                                      = 1795
	ErrorUnknownPort                                                        = 1796
	ErrorUnknownPrinterDriver                                               = 1797
	ErrorUnknownPrintprocessor                                              = 1798
	ErrorInvalidSeparatorFile                                               = 1799
	ErrorInvalidPriority                                                    = 1800
	ErrorInvalidPrinterName                                                 = 1801
	ErrorPrinterAlreadyExists                                               = 1802
	ErrorInvalidPrinterCommand                                              = 1803
	ErrorInvalidDatatype                                                    = 1804
	ErrorInvalidEnvironment                                                 = 1805
	RpcSNoMoreBindings                                                      = 1806
	ErrorNologonInterdomainTrustAccount                                     = 1807
	ErrorNologonWorkstationTrustAccount                                     = 1808
	ErrorNologonServerTrustAccount                                          = 1809
	ErrorDomainTrustInconsistent                                            = 1810
	ErrorServerHasOpenHandles                                               = 1811
	ErrorResourceDataNotFound                                               = 1812
	ErrorResourceTypeNotFound                                               = 1813
	ErrorResourceNameNotFound                                               = 1814
	ErrorResourceLangNotFound                                               = 1815
	ErrorNotEnoughQuota                                                     = 1816
	RpcSNoInterfaces                                                        = 1817
	RpcSCallCancelled                                                       = 1818
	RpcSBindingIncomplete                                                   = 1819
	RpcSCommFailure                                                         = 1820
	RpcSUnsupportedAuthnLevel                                               = 1821
	RpcSNoPrincName                                                         = 1822
	RpcSNotRpcError                                                         = 1823
	RpcSUuidLocalOnly                                                       = 1824
	RpcSSecPkgError                                                         = 1825
	RpcSNotCancelled                                                        = 1826
	RpcXInvalidEsAction                                                     = 1827
	RpcXWrongEsVersion                                                      = 1828
	RpcXWrongStubVersion                                                    = 1829
	RpcXInvalidPipeObject                                                   = 1830
	RpcXWrongPipeOrder                                                      = 1831
	RpcXWrongPipeVersion                                                    = 1832
	RpcSCookieAuthFailed                                                    = 1833
	RpcSGroupMemberNotFound                                                 = 1898
	EptSCantCreate                                                          = 1899
	RpcSInvalidObject                                                       = 1900
	ErrorInvalidTime                                                        = 1901
	ErrorInvalidFormName                                                    = 1902
	ErrorInvalidFormSize                                                    = 1903
	ErrorAlreadyWaiting                                                     = 1904
	ErrorPrinterDeleted                                                     = 1905
	ErrorInvalidPrinterState                                                = 1906
	ErrorPasswordMustChange                                                 = 1907
	ErrorDomainControllerNotFound                                           = 1908
	ErrorAccountLockedOut                                                   = 1909
	OrInvalidOxid                                                           = 1910
	OrInvalidOid                                                            = 1911
	OrInvalidSet                                                            = 1912
	RpcSSendIncomplete                                                      = 1913
	RpcSInvalidAsyncHandle                                                  = 1914
	RpcSInvalidAsyncCall                                                    = 1915
	RpcXPipeClosed                                                          = 1916
	RpcXPipeDisciplineError                                                 = 1917
	RpcXPipeEmpty                                                           = 1918
	ErrorNoSitename                                                         = 1919
	ErrorCantAccessFile                                                     = 1920
	ErrorCantResolveFilename                                                = 1921
	RpcSEntryTypeMismatch                                                   = 1922
	RpcSNotAllObjsExported                                                  = 1923
	RpcSInterfaceNotExported                                                = 1924
	RpcSProfileNotAdded                                                     = 1925
	RpcSPrfEltNotAdded                                                      = 1926
	RpcSPrfEltNotRemoved                                                    = 1927
	RpcSGrpEltNotAdded                                                      = 1928
	RpcSGrpEltNotRemoved                                                    = 1929
	ErrorKmDriverBlocked                                                    = 1930
	ErrorContextExpired                                                     = 1931
	ErrorPerUserTrustQuotaExceeded                                          = 1932
	ErrorAllUserTrustQuotaExceeded                                          = 1933
	ErrorUserDeleteTrustQuotaExceeded                                       = 1934
	ErrorAuthenticationFirewallFailed                                       = 1935
	ErrorRemotePrintConnectionsBlocked                                      = 1936
	ErrorNtlmBlocked                                                        = 1937
	ErrorInvalidPixelFormat                                                 = 2000
	ErrorBadDriver                                                          = 2001
	ErrorInvalidWindowStyle                                                 = 2002
	ErrorMetafileNotSupported                                               = 2003
	ErrorTransformNotSupported                                              = 2004
	ErrorClippingNotSupported                                               = 2005
	ErrorInvalidCmm                                                         = 2010
	ErrorInvalidProfile                                                     = 2011
	ErrorTagNotFound                                                        = 2012
	ErrorTagNotPresent                                                      = 2013
	ErrorDuplicateTag                                                       = 2014
	ErrorProfileNotAssociatedWithDevice                                     = 2015
	ErrorProfileNotFound                                                    = 2016
	ErrorInvalidColorspace                                                  = 2017
	ErrorIcmNotEnabled                                                      = 2018
	ErrorDeletingIcmXform                                                   = 2019
	ErrorInvalidTransform                                                   = 2020
	ErrorColorspaceMismatch                                                 = 2021
	ErrorInvalidColorindex                                                  = 2022
	ErrorProfileDoesNotMatchDevice                                          = 2023
	ErrorConnectedOtherPassword                                             = 2108
	ErrorConnectedOtherPasswordDefault                                      = 2109
	ErrorBadUsername                                                        = 2202
	ErrorNotConnected                                                       = 2250
	ErrorOpenFiles                                                          = 2401
	ErrorActiveConnections                                                  = 2402
	ErrorDeviceInUse                                                        = 2404
	ErrorUnknownPrintMonitor                                                = 3000
	ErrorPrinterDriverInUse                                                 = 3001
	ErrorSpoolFileNotFound                                                  = 3002
	ErrorSplNoStartdoc                                                      = 3003
	ErrorSplNoAddjob                                                        = 3004
	ErrorPrintProcessorAlreadyInstalled                                     = 3005
	ErrorPrintMonitorAlreadyInstalled                                       = 3006
	ErrorInvalidPrintMonitor                                                = 3007
	ErrorPrintMonitorInUse                                                  = 3008
	ErrorPrinterHasJobsQueued                                               = 3009
	ErrorSuccessRebootRequired                                              = 3010
	ErrorSuccessRestartRequired                                             = 3011
	ErrorPrinterNotFound                                                    = 3012
	ErrorPrinterDriverWarned                                                = 3013
	ErrorPrinterDriverBlocked                                               = 3014
	ErrorPrinterDriverPackageInUse                                          = 3015
	ErrorCoreDriverPackageNotFound                                          = 3016
	ErrorFailRebootRequired                                                 = 3017
	ErrorFailRebootInitiated                                                = 3018
	ErrorPrinterDriverDownloadNeeded                                        = 3019
	ErrorPrintJobRestartRequired                                            = 3020
	ErrorIoReissueAsCached                                                  = 3950
	ErrorWinsInternal                                                       = 4000
	ErrorCanNotDelLocalWins                                                 = 4001
	ErrorStaticInit                                                         = 4002
	ErrorIncBackup                                                          = 4003
	ErrorFullBackup                                                         = 4004
	ErrorRecNonExistent                                                     = 4005
	ErrorRplNotAllowed                                                      = 4006
	PeerdistErrorContentinfoVersionUnsupported                              = 4050
	PeerdistErrorCannotParseContentinfo                                     = 4051
	PeerdistErrorMissingData                                                = 4052
	PeerdistErrorNoMore                                                     = 4053
	PeerdistErrorNotInitialized                                             = 4054
	PeerdistErrorAlreadyInitialized                                         = 4055
	PeerdistErrorShutdownInProgress                                         = 4056
	PeerdistErrorInvalidated                                                = 4057
	PeerdistErrorAlreadyExists                                              = 4058
	PeerdistErrorOperationNotfound                                          = 4059
	PeerdistErrorAlreadyCompleted                                           = 4060
	PeerdistErrorOutOfBounds                                                = 4061
	PeerdistErrorVersionUnsupported                                         = 4062
	PeerdistErrorInvalidConfiguration                                       = 4063
	PeerdistErrorNotLicensed                                                = 4064
	PeerdistErrorServiceUnavailable                                         = 4065
	ErrorDhcpAddressConflict                                                = 4100
	ErrorWmiGuidNotFound                                                    = 4200
	ErrorWmiInstanceNotFound                                                = 4201
	ErrorWmiItemidNotFound                                                  = 4202
	ErrorWmiTryAgain                                                        = 4203
	ErrorWmiDpNotFound                                                      = 4204
	ErrorWmiUnresolvedInstanceRef                                           = 4205
	ErrorWmiAlreadyEnabled                                                  = 4206
	ErrorWmiGuidDisconnected                                                = 4207
	ErrorWmiServerUnavailable                                               = 4208
	ErrorWmiDpFailed                                                        = 4209
	ErrorWmiInvalidMof                                                      = 4210
	ErrorWmiInvalidReginfo                                                  = 4211
	ErrorWmiAlreadyDisabled                                                 = 4212
	ErrorWmiReadOnly                                                        = 4213
	ErrorWmiSetFailure                                                      = 4214
	ErrorInvalidMedia                                                       = 4300
	ErrorInvalidLibrary                                                     = 4301
	ErrorInvalidMediaPool                                                   = 4302
	ErrorDriveMediaMismatch                                                 = 4303
	ErrorMediaOffline                                                       = 4304
	ErrorLibraryOffline                                                     = 4305
	ErrorEmpty                                                              = 4306
	ErrorNotEmpty                                                           = 4307
	ErrorMediaUnavailable                                                   = 4308
	ErrorResourceDisabled                                                   = 4309
	ErrorInvalidCleaner                                                     = 4310
	ErrorUnableToClean                                                      = 4311
	ErrorObjectNotFound                                                     = 4312
	ErrorDatabaseFailure                                                    = 4313
	ErrorDatabaseFull                                                       = 4314
	ErrorMediaIncompatible                                                  = 4315
	ErrorResourceNotPresent                                                 = 4316
	ErrorInvalidOperation                                                   = 4317
	ErrorMediaNotAvailable                                                  = 4318
	ErrorDeviceNotAvailable                                                 = 4319
	ErrorRequestRefused                                                     = 4320
	ErrorInvalidDriveObject                                                 = 4321
	ErrorLibraryFull                                                        = 4322
	ErrorMediumNotAccessible                                                = 4323
	ErrorUnableToLoadMedium                                                 = 4324
	ErrorUnableToInventoryDrive                                             = 4325
	ErrorUnableToInventorySlot                                              = 4326
	ErrorUnableToInventoryTransport                                         = 4327
	ErrorTransportFull                                                      = 4328
	ErrorControllingIeport                                                  = 4329
	ErrorUnableToEjectMountedMedia                                          = 4330
	ErrorCleanerSlotSet                                                     = 4331
	ErrorCleanerSlotNotSet                                                  = 4332
	ErrorCleanerCartridgeSpent                                              = 4333
	ErrorUnexpectedOmid                                                     = 4334
	ErrorCantDeleteLastItem                                                 = 4335
	ErrorMessageExceedsMaxSize                                              = 4336
	ErrorVolumeContainsSysFiles                                             = 4337
	ErrorIndigenousType                                                     = 4338
	ErrorNoSupportingDrives                                                 = 4339
	ErrorCleanerCartridgeInstalled                                          = 4340
	ErrorIeportFull                                                         = 4341
	ErrorFileOffline                                                        = 4350
	ErrorRemoteStorageNotActive                                             = 4351
	ErrorRemoteStorageMediaError                                            = 4352
	ErrorNotAReparsePoint                                                   = 4390
	ErrorReparseAttributeConflict                                           = 4391
	ErrorInvalidReparseData                                                 = 4392
	ErrorReparseTagInvalid                                                  = 4393
	ErrorReparseTagMismatch                                                 = 4394
	ErrorVolumeNotSisEnabled                                                = 4500
	ErrorDependentResourceExists                                            = 5001
	ErrorDependencyNotFound                                                 = 5002
	ErrorDependencyAlreadyExists                                            = 5003
	ErrorResourceNotOnline                                                  = 5004
	ErrorHostNodeNotAvailable                                               = 5005
	ErrorResourceNotAvailable                                               = 5006
	ErrorResourceNotFound                                                   = 5007
	ErrorShutdownCluster                                                    = 5008
	ErrorCantEvictActiveNode                                                = 5009
	ErrorObjectAlreadyExists                                                = 5010
	ErrorObjectInList                                                       = 5011
	ErrorGroupNotAvailable                                                  = 5012
	ErrorGroupNotFound                                                      = 5013
	ErrorGroupNotOnline                                                     = 5014
	ErrorHostNodeNotResourceOwner                                           = 5015
	ErrorHostNodeNotGroupOwner                                              = 5016
	ErrorResmonCreateFailed                                                 = 5017
	ErrorResmonOnlineFailed                                                 = 5018
	ErrorResourceOnline                                                     = 5019
	ErrorQuorumResource                                                     = 5020
	ErrorNotQuorumCapable                                                   = 5021
	ErrorClusterShuttingDown                                                = 5022
	ErrorInvalidState                                                       = 5023
	ErrorResourcePropertiesStored                                           = 5024
	ErrorNotQuorumClass                                                     = 5025
	ErrorCoreResource                                                       = 5026
	ErrorQuorumResourceOnlineFailed                                         = 5027
	ErrorQuorumlogOpenFailed                                                = 5028
	ErrorClusterlogCorrupt                                                  = 5029
	ErrorClusterlogRecordExceedsMaxsize                                     = 5030
	ErrorClusterlogExceedsMaxsize                                           = 5031
	ErrorClusterlogChkpointNotFound                                         = 5032
	ErrorClusterlogNotEnoughSpace                                           = 5033
	ErrorQuorumOwnerAlive                                                   = 5034
	ErrorNetworkNotAvailable                                                = 5035
	ErrorNodeNotAvailable                                                   = 5036
	ErrorAllNodesNotAvailable                                               = 5037
	ErrorResourceFailed                                                     = 5038
	ErrorClusterInvalidNode                                                 = 5039
	ErrorClusterNodeExists                                                  = 5040
	ErrorClusterJoinInProgress                                              = 5041
	ErrorClusterNodeNotFound                                                = 5042
	ErrorClusterLocalNodeNotFound                                           = 5043
	ErrorClusterNetworkExists                                               = 5044
	ErrorClusterNetworkNotFound                                             = 5045
	ErrorClusterNetinterfaceExists                                          = 5046
	ErrorClusterNetinterfaceNotFound                                        = 5047
	ErrorClusterInvalidRequest                                              = 5048
	ErrorClusterInvalidNetworkProvider                                      = 5049
	ErrorClusterNodeDown                                                    = 5050
	ErrorClusterNodeUnreachable                                             = 5051
	ErrorClusterNodeNotMember                                               = 5052
	ErrorClusterJoinNotInProgress                                           = 5053
	ErrorClusterInvalidNetwork                                              = 5054
	ErrorClusterNodeUp                                                      = 5056
	ErrorClusterIpaddrInUse                                                 = 5057
	ErrorClusterNodeNotPaused                                               = 5058
	ErrorClusterNoSecurityContext                                           = 5059
	ErrorClusterNetworkNotInternal                                          = 5060
	ErrorClusterNodeAlreadyUp                                               = 5061
	ErrorClusterNodeAlreadyDown                                             = 5062
	ErrorClusterNetworkAlreadyOnline                                        = 5063
	ErrorClusterNetworkAlreadyOffline                                       = 5064
	ErrorClusterNodeAlreadyMember                                           = 5065
	ErrorClusterLastInternalNetwork                                         = 5066
	ErrorClusterNetworkHasDependents                                        = 5067
	ErrorInvalidOperationOnQuorum                                           = 5068
	ErrorDependencyNotAllowed                                               = 5069
	ErrorClusterNodePaused                                                  = 5070
	ErrorNodeCantHostResource                                               = 5071
	ErrorClusterNodeNotReady                                                = 5072
	ErrorClusterNodeShuttingDown                                            = 5073
	ErrorClusterJoinAborted                                                 = 5074
	ErrorClusterIncompatibleVersions                                        = 5075
	ErrorClusterMaxnumOfResourcesExceeded                                   = 5076
	ErrorClusterSystemConfigChanged                                         = 5077
	ErrorClusterResourceTypeNotFound                                        = 5078
	ErrorClusterRestypeNotSupported                                         = 5079
	ErrorClusterResnameNotFound                                             = 5080
	ErrorClusterNoRpcPackagesRegistered                                     = 5081
	ErrorClusterOwnerNotInPreflist                                          = 5082
	ErrorClusterDatabaseSeqmismatch                                         = 5083
	ErrorResmonInvalidState                                                 = 5084
	ErrorClusterGumNotLocker                                                = 5085
	ErrorQuorumDiskNotFound                                                 = 5086
	ErrorDatabaseBackupCorrupt                                              = 5087
	ErrorClusterNodeAlreadyHasDfsRoot                                       = 5088
	ErrorResourcePropertyUnchangeable                                       = 5089
	ErrorClusterMembershipInvalidState                                      = 5890
	ErrorClusterQuorumlogNotFound                                           = 5891
	ErrorClusterMembershipHalt                                              = 5892
	ErrorClusterInstanceIdMismatch                                          = 5893
	ErrorClusterNetworkNotFoundForIp                                        = 5894
	ErrorClusterPropertyDataTypeMismatch                                    = 5895
	ErrorClusterEvictWithoutCleanup                                         = 5896
	ErrorClusterParameterMismatch                                           = 5897
	ErrorNodeCannotBeClustered                                              = 5898
	ErrorClusterWrongOsVersion                                              = 5899
	ErrorClusterCantCreateDupClusterName                                    = 5900
	ErrorCluscfgAlreadyCommitted                                            = 5901
	ErrorCluscfgRollbackFailed                                              = 5902
	ErrorCluscfgSystemDiskDriveLetterConflict                               = 5903
	ErrorClusterOldVersion                                                  = 5904
	ErrorClusterMismatchedComputerAcctName                                  = 5905
	ErrorClusterNoNetAdapters                                               = 5906
	ErrorClusterPoisoned                                                    = 5907
	ErrorClusterGroupMoving                                                 = 5908
	ErrorClusterResourceTypeBusy                                            = 5909
	ErrorResourceCallTimedOut                                               = 5910
	ErrorInvalidClusterIpv6Address                                          = 5911
	ErrorClusterInternalInvalidFunction                                     = 5912
	ErrorClusterParameterOutOfBounds                                        = 5913
	ErrorClusterPartialSend                                                 = 5914
	ErrorClusterRegistryInvalidFunction                                     = 5915
	ErrorClusterInvalidStringTermination                                    = 5916
	ErrorClusterInvalidStringFormat                                         = 5917
	ErrorClusterDatabaseTransactionInProgress                               = 5918
	ErrorClusterDatabaseTransactionNotInProgress                            = 5919
	ErrorClusterNullData                                                    = 5920
	ErrorClusterPartialRead                                                 = 5921
	ErrorClusterPartialWrite                                                = 5922
	ErrorClusterCantDeserializeData                                         = 5923
	ErrorDependentResourcePropertyConflict                                  = 5924
	ErrorClusterNoQuorum                                                    = 5925
	ErrorClusterInvalidIpv6Network                                          = 5926
	ErrorClusterInvalidIpv6TunnelNetwork                                    = 5927
	ErrorQuorumNotAllowedInThisGroup                                        = 5928
	ErrorDependencyTreeTooComplex                                           = 5929
	ErrorExceptionInResourceCall                                            = 5930
	ErrorClusterRhsFailedInitialization                                     = 5931
	ErrorClusterNotInstalled                                                = 5932
	ErrorClusterResourcesMustBeOnlineOnTheSameNode                          = 5933
	ErrorClusterMaxNodesInCluster                                           = 5934
	ErrorClusterTooManyNodes                                                = 5935
	ErrorClusterObjectAlreadyUsed                                           = 5936
	ErrorNoncoreGroupsFound                                                 = 5937
	ErrorFileShareResourceConflict                                          = 5938
	ErrorClusterEvictInvalidRequest                                         = 5939
	ErrorClusterSingletonResource                                           = 5940
	ErrorClusterGroupSingletonResource                                      = 5941
	ErrorClusterResourceProviderFailed                                      = 5942
	ErrorClusterResourceConfigurationError                                  = 5943
	ErrorClusterGroupBusy                                                   = 5944
	ErrorClusterNotSharedVolume                                             = 5945
	ErrorClusterInvalidSecurityDescriptor                                   = 5946
	ErrorClusterSharedVolumesInUse                                          = 5947
	ErrorClusterUseSharedVolumesApi                                         = 5948
	ErrorClusterBackupInProgress                                            = 5949
	ErrorNonCsvPath                                                         = 5950
	ErrorCsvVolumeNotLocal                                                  = 5951
	ErrorClusterWatchdogTerminating                                         = 5952
	ErrorEncryptionFailed                                                   = 6000
	ErrorDecryptionFailed                                                   = 6001
	ErrorFileEncrypted                                                      = 6002
	ErrorNoRecoveryPolicy                                                   = 6003
	ErrorNoEfs                                                              = 6004
	ErrorWrongEfs                                                           = 6005
	ErrorNoUserKeys                                                         = 6006
	ErrorFileNotEncrypted                                                   = 6007
	ErrorNotExportFormat                                                    = 6008
	ErrorFileReadOnly                                                       = 6009
	ErrorDirEfsDisallowed                                                   = 6010
	ErrorEfsServerNotTrusted                                                = 6011
	ErrorBadRecoveryPolicy                                                  = 6012
	ErrorEfsAlgBlobTooBig                                                   = 6013
	ErrorVolumeNotSupportEfs                                                = 6014
	ErrorEfsDisabled                                                        = 6015
	ErrorEfsVersionNotSupport                                               = 6016
	ErrorCsEncryptionInvalidServerResponse                                  = 6017
	ErrorCsEncryptionUnsupportedServer                                      = 6018
	ErrorCsEncryptionExistingEncryptedFile                                  = 6019
	ErrorCsEncryptionNewEncryptedFile                                       = 6020
	ErrorCsEncryptionFileNotCse                                             = 6021
	ErrorEncryptionPolicyDeniesOperation                                    = 6022
	ErrorNoBrowserServersFound                                              = 6118
	SchedEServiceNotLocalsystem                                             = 6200
	ErrorLogSectorInvalid                                                   = 6600
	ErrorLogSectorParityInvalid                                             = 6601
	ErrorLogSectorRemapped                                                  = 6602
	ErrorLogBlockIncomplete                                                 = 6603
	ErrorLogInvalidRange                                                    = 6604
	ErrorLogBlocksExhausted                                                 = 6605
	ErrorLogReadContextInvalid                                              = 6606
	ErrorLogRestartInvalid                                                  = 6607
	ErrorLogBlockVersion                                                    = 6608
	ErrorLogBlockInvalid                                                    = 6609
	ErrorLogReadModeInvalid                                                 = 6610
	ErrorLogNoRestart                                                       = 6611
	ErrorLogMetadataCorrupt                                                 = 6612
	ErrorLogMetadataInvalid                                                 = 6613
	ErrorLogMetadataInconsistent                                            = 6614
	ErrorLogReservationInvalid                                              = 6615
	ErrorLogCantDelete                                                      = 6616
	ErrorLogContainerLimitExceeded                                          = 6617
	ErrorLogStartOfLog                                                      = 6618
	ErrorLogPolicyAlreadyInstalled                                          = 6619
	ErrorLogPolicyNotInstalled                                              = 6620
	ErrorLogPolicyInvalid                                                   = 6621
	ErrorLogPolicyConflict                                                  = 6622
	ErrorLogPinnedArchiveTail                                               = 6623
	ErrorLogRecordNonexistent                                               = 6624
	ErrorLogRecordsReservedInvalid                                          = 6625
	ErrorLogSpaceReservedInvalid                                            = 6626
	ErrorLogTailInvalid                                                     = 6627
	ErrorLogFull                                                            = 6628
	ErrorCouldNotResizeLog                                                  = 6629
	ErrorLogMultiplexed                                                     = 6630
	ErrorLogDedicated                                                       = 6631
	ErrorLogArchiveNotInProgress                                            = 6632
	ErrorLogArchiveInProgress                                               = 6633
	ErrorLogEphemeral                                                       = 6634
	ErrorLogNotEnoughContainers                                             = 6635
	ErrorLogClientAlreadyRegistered                                         = 6636
	ErrorLogClientNotRegistered                                             = 6637
	ErrorLogFullHandlerInProgress                                           = 6638
	ErrorLogContainerReadFailed                                             = 6639
	ErrorLogContainerWriteFailed                                            = 6640
	ErrorLogContainerOpenFailed                                             = 6641
	ErrorLogContainerStateInvalid                                           = 6642
	ErrorLogStateInvalid                                                    = 6643
	ErrorLogPinned                                                          = 6644
	ErrorLogMetadataFlushFailed                                             = 6645
	ErrorLogInconsistentSecurity                                            = 6646
	ErrorLogAppendedFlushFailed                                             = 6647
	ErrorLogPinnedReservation                                               = 6648
	ErrorInvalidTransaction                                                 = 6700
	ErrorTransactionNotActive                                               = 6701
	ErrorTransactionRequestNotValid                                         = 6702
	ErrorTransactionNotRequested                                            = 6703
	ErrorTransactionAlreadyAborted                                          = 6704
	ErrorTransactionAlreadyCommitted                                        = 6705
	ErrorTmInitializationFailed                                             = 6706
	ErrorResourcemanagerReadOnly                                            = 6707
	ErrorTransactionNotJoined                                               = 6708
	ErrorTransactionSuperiorExists                                          = 6709
	ErrorCrmProtocolAlreadyExists                                           = 6710
	ErrorTransactionPropagationFailed                                       = 6711
	ErrorCrmProtocolNotFound                                                = 6712
	ErrorTransactionInvalidMarshallBuffer                                   = 6713
	ErrorCurrentTransactionNotValid                                         = 6714
	ErrorTransactionNotFound                                                = 6715
	ErrorResourcemanagerNotFound                                            = 6716
	ErrorEnlistmentNotFound                                                 = 6717
	ErrorTransactionmanagerNotFound                                         = 6718
	ErrorTransactionmanagerNotOnline                                        = 6719
	ErrorTransactionmanagerRecoveryNameCollision                            = 6720
	ErrorTransactionNotRoot                                                 = 6721
	ErrorTransactionObjectExpired                                           = 6722
	ErrorTransactionResponseNotEnlisted                                     = 6723
	ErrorTransactionRecordTooLong                                           = 6724
	ErrorImplicitTransactionNotSupported                                    = 6725
	ErrorTransactionIntegrityViolated                                       = 6726
	ErrorTransactionmanagerIdentityMismatch                                 = 6727
	ErrorRmCannotBeFrozenForSnapshot                                        = 6728
	ErrorTransactionMustWritethrough                                        = 6729
	ErrorTransactionNoSuperior                                              = 6730
	ErrorHeuristicDamagePossible                                            = 6731
	ErrorTransactionalConflict                                              = 6800
	ErrorRmNotActive                                                        = 6801
	ErrorRmMetadataCorrupt                                                  = 6802
	ErrorDirectoryNotRm                                                     = 6803
	ErrorTransactionsUnsupportedRemote                                      = 6805
	ErrorLogResizeInvalidSize                                               = 6806
	ErrorObjectNoLongerExists                                               = 6807
	ErrorStreamMiniversionNotFound                                          = 6808
	ErrorStreamMiniversionNotValid                                          = 6809
	ErrorMiniversionInaccessibleFromSpecifiedTransaction                    = 6810
	ErrorCantOpenMiniversionWithModifyIntent                                = 6811
	ErrorCantCreateMoreStreamMiniversions                                   = 6812
	ErrorRemoteFileVersionMismatch                                          = 6814
	ErrorHandleNoLongerValid                                                = 6815
	ErrorNoTxfMetadata                                                      = 6816
	ErrorLogCorruptionDetected                                              = 6817
	ErrorCantRecoverWithHandleOpen                                          = 6818
	ErrorRmDisconnected                                                     = 6819
	ErrorEnlistmentNotSuperior                                              = 6820
	ErrorRecoveryNotNeeded                                                  = 6821
	ErrorRmAlreadyStarted                                                   = 6822
	ErrorFileIdentityNotPersistent                                          = 6823
	ErrorCantBreakTransactionalDependency                                   = 6824
	ErrorCantCrossRmBoundary                                                = 6825
	ErrorTxfDirNotEmpty                                                     = 6826
	ErrorIndoubtTransactionsExist                                           = 6827
	ErrorTmVolatile                                                         = 6828
	ErrorRollbackTimerExpired                                               = 6829
	ErrorTxfAttributeCorrupt                                                = 6830
	ErrorEfsNotAllowedInTransaction                                         = 6831
	ErrorTransactionalOpenNotAllowed                                        = 6832
	ErrorLogGrowthFailed                                                    = 6833
	ErrorTransactedMappingUnsupportedRemote                                 = 6834
	ErrorTxfMetadataAlreadyPresent                                          = 6835
	ErrorTransactionScopeCallbacksNotSet                                    = 6836
	ErrorTransactionRequiredPromotion                                       = 6837
	ErrorCannotExecuteFileInTransaction                                     = 6838
	ErrorTransactionsNotFrozen                                              = 6839
	ErrorTransactionFreezeInProgress                                        = 6840
	ErrorNotSnapshotVolume                                                  = 6841
	ErrorNoSavepointWithOpenFiles                                           = 6842
	ErrorDataLostRepair                                                     = 6843
	ErrorSparseNotAllowedInTransaction                                      = 6844
	ErrorTmIdentityMismatch                                                 = 6845
	ErrorFloatedSection                                                     = 6846
	ErrorCannotAcceptTransactedWork                                         = 6847
	ErrorCannotAbortTransactions                                            = 6848
	ErrorBadClusters                                                        = 6849
	ErrorCompressionNotAllowedInTransaction                                 = 6850
	ErrorVolumeDirty                                                        = 6851
	ErrorNoLinkTrackingInTransaction                                        = 6852
	ErrorOperationNotSupportedInTransaction                                 = 6853
	ErrorExpiredHandle                                                      = 6854
	ErrorTransactionNotEnlisted                                             = 6855
	ErrorCtxWinstationNameInvalid                                           = 7001
	ErrorCtxInvalidPd                                                       = 7002
	ErrorCtxPdNotFound                                                      = 7003
	ErrorCtxWdNotFound                                                      = 7004
	ErrorCtxCannotMakeEventlogEntry                                         = 7005
	ErrorCtxServiceNameCollision                                            = 7006
	ErrorCtxClosePending                                                    = 7007
	ErrorCtxNoOutbuf                                                        = 7008
	ErrorCtxModemInfNotFound                                                = 7009
	ErrorCtxInvalidModemname                                                = 7010
	ErrorCtxModemResponseError                                              = 7011
	ErrorCtxModemResponseTimeout                                            = 7012
	ErrorCtxModemResponseNoCarrier                                          = 7013
	ErrorCtxModemResponseNoDialtone                                         = 7014
	ErrorCtxModemResponseBusy                                               = 7015
	ErrorCtxModemResponseVoice                                              = 7016
	ErrorCtxTdError                                                         = 7017
	ErrorCtxWinstationNotFound                                              = 7022
	ErrorCtxWinstationAlreadyExists                                         = 7023
	ErrorCtxWinstationBusy                                                  = 7024
	ErrorCtxBadVideoMode                                                    = 7025
	ErrorCtxGraphicsInvalid                                                 = 7035
	ErrorCtxLogonDisabled                                                   = 7037
	ErrorCtxNotConsole                                                      = 7038
	ErrorCtxClientQueryTimeout                                              = 7040
	ErrorCtxConsoleDisconnect                                               = 7041
	ErrorCtxConsoleConnect                                                  = 7042
	ErrorCtxShadowDenied                                                    = 7044
	ErrorCtxWinstationAccessDenied                                          = 7045
	ErrorCtxInvalidWd                                                       = 7049
	ErrorCtxShadowInvalid                                                   = 7050
	ErrorCtxShadowDisabled                                                  = 7051
	ErrorCtxClientLicenseInUse                                              = 7052
	ErrorCtxClientLicenseNotSet                                             = 7053
	ErrorCtxLicenseNotAvailable                                             = 7054
	ErrorCtxLicenseClientInvalid                                            = 7055
	ErrorCtxLicenseExpired                                                  = 7056
	ErrorCtxShadowNotRunning                                                = 7057
	ErrorCtxShadowEndedByModeChange                                         = 7058
	ErrorActivationCountExceeded                                            = 7059
	ErrorCtxWinstationsDisabled                                             = 7060
	ErrorCtxEncryptionLevelRequired                                         = 7061
	ErrorCtxSessionInUse                                                    = 7062
	ErrorCtxNoForceLogoff                                                   = 7063
	ErrorCtxAccountRestriction                                              = 7064
	ErrorRdpProtocolError                                                   = 7065
	ErrorCtxCdmConnect                                                      = 7066
	ErrorCtxCdmDisconnect                                                   = 7067
	ErrorCtxSecurityLayerError                                              = 7068
	ErrorTsIncompatibleSessions                                             = 7069
	ErrorTsVideoSubsystemError                                              = 7070
	FrsErrInvalidApiSequence                                                = 8001
	FrsErrStartingService                                                   = 8002
	FrsErrStoppingService                                                   = 8003
	FrsErrInternalApi                                                       = 8004
	FrsErrInternal                                                          = 8005
	FrsErrServiceComm                                                       = 8006
	FrsErrInsufficientPriv                                                  = 8007
	FrsErrAuthentication                                                    = 8008
	FrsErrParentInsufficientPriv                                            = 8009
	FrsErrParentAuthentication                                              = 8010
	FrsErrChildToParentComm                                                 = 8011
	FrsErrParentToChildComm                                                 = 8012
	FrsErrSysvolPopulate                                                    = 8013
	FrsErrSysvolPopulateTimeout                                             = 8014
	FrsErrSysvolIsBusy                                                      = 8015
	FrsErrSysvolDemote                                                      = 8016
	FrsErrInvalidServiceParameter                                           = 8017
	ErrorDsNotInstalled                                                     = 8200
	ErrorDsMembershipEvaluatedLocally                                       = 8201
	ErrorDsNoAttributeOrValue                                               = 8202
	ErrorDsInvalidAttributeSyntax                                           = 8203
	ErrorDsAttributeTypeUndefined                                           = 8204
	ErrorDsAttributeOrValueExists                                           = 8205
	ErrorDsBusy                                                             = 8206
	ErrorDsUnavailable                                                      = 8207
	ErrorDsNoRidsAllocated                                                  = 8208
	ErrorDsNoMoreRids                                                       = 8209
	ErrorDsIncorrectRoleOwner                                               = 8210
	ErrorDsRidmgrInitError                                                  = 8211
	ErrorDsObjClassViolation                                                = 8212
	ErrorDsCantOnNonLeaf                                                    = 8213
	ErrorDsCantOnRdn                                                        = 8214
	ErrorDsCantModObjClass                                                  = 8215
	ErrorDsCrossDomMoveError                                                = 8216
	ErrorDsGcNotAvailable                                                   = 8217
	ErrorSharedPolicy                                                       = 8218
	ErrorPolicyObjectNotFound                                               = 8219
	ErrorPolicyOnlyInDs                                                     = 8220
	ErrorPromotionActive                                                    = 8221
	ErrorNoPromotionActive                                                  = 8222
	ErrorDsOperationsError                                                  = 8224
	ErrorDsProtocolError                                                    = 8225
	ErrorDsTimelimitExceeded                                                = 8226
	ErrorDsSizelimitExceeded                                                = 8227
	ErrorDsAdminLimitExceeded                                               = 8228
	ErrorDsCompareFalse                                                     = 8229
	ErrorDsCompareTrue                                                      = 8230
	ErrorDsAuthMethodNotSupported                                           = 8231
	ErrorDsStrongAuthRequired                                               = 8232
	ErrorDsInappropriateAuth                                                = 8233
	ErrorDsAuthUnknown                                                      = 8234
	ErrorDsReferral                                                         = 8235
	ErrorDsUnavailableCritExtension                                         = 8236
	ErrorDsConfidentialityRequired                                          = 8237
	ErrorDsInappropriateMatching                                            = 8238
	ErrorDsConstraintViolation                                              = 8239
	ErrorDsNoSuchObject                                                     = 8240
	ErrorDsAliasProblem                                                     = 8241
	ErrorDsInvalidDnSyntax                                                  = 8242
	ErrorDsIsLeaf                                                           = 8243
	ErrorDsAliasDerefProblem                                                = 8244
	ErrorDsUnwillingToPerform                                               = 8245
	ErrorDsLoopDetect                                                       = 8246
	ErrorDsNamingViolation                                                  = 8247
	ErrorDsObjectResultsTooLarge                                            = 8248
	ErrorDsAffectsMultipleDsas                                              = 8249
	ErrorDsServerDown                                                       = 8250
	ErrorDsLocalError                                                       = 8251
	ErrorDsEncodingError                                                    = 8252
	ErrorDsDecodingError                                                    = 8253
	ErrorDsFilterUnknown                                                    = 8254
	ErrorDsParamError                                                       = 8255
	ErrorDsNotSupported                                                     = 8256
	ErrorDsNoResultsReturned                                                = 8257
	ErrorDsControlNotFound                                                  = 8258
	ErrorDsClientLoop                                                       = 8259
	ErrorDsReferralLimitExceeded                                            = 8260
	ErrorDsSortControlMissing                                               = 8261
	ErrorDsOffsetRangeError                                                 = 8262
	ErrorDsRootMustBeNc                                                     = 8301
	ErrorDsAddReplicaInhibited                                              = 8302
	ErrorDsAttNotDefInSchema                                                = 8303
	ErrorDsMaxObjSizeExceeded                                               = 8304
	ErrorDsObjStringNameExists                                              = 8305
	ErrorDsNoRdnDefinedInSchema                                             = 8306
	ErrorDsRdnDoesntMatchSchema                                             = 8307
	ErrorDsNoRequestedAttsFound                                             = 8308
	ErrorDsUserBufferToSmall                                                = 8309
	ErrorDsAttIsNotOnObj                                                    = 8310
	ErrorDsIllegalModOperation                                              = 8311
	ErrorDsObjTooLarge                                                      = 8312
	ErrorDsBadInstanceType                                                  = 8313
	ErrorDsMasterdsaRequired                                                = 8314
	ErrorDsObjectClassRequired                                              = 8315
	ErrorDsMissingRequiredAtt                                               = 8316
	ErrorDsAttNotDefForClass                                                = 8317
	ErrorDsAttAlreadyExists                                                 = 8318
	ErrorDsCantAddAttValues                                                 = 8320
	ErrorDsSingleValueConstraint                                            = 8321
	ErrorDsRangeConstraint                                                  = 8322
	ErrorDsAttValAlreadyExists                                              = 8323
	ErrorDsCantRemMissingAtt                                                = 8324
	ErrorDsCantRemMissingAttVal                                             = 8325
	ErrorDsRootCantBeSubref                                                 = 8326
	ErrorDsNoChaining                                                       = 8327
	ErrorDsNoChainedEval                                                    = 8328
	ErrorDsNoParentObject                                                   = 8329
	ErrorDsParentIsAnAlias                                                  = 8330
	ErrorDsCantMixMasterAndReps                                             = 8331
	ErrorDsChildrenExist                                                    = 8332
	ErrorDsObjNotFound                                                      = 8333
	ErrorDsAliasedObjMissing                                                = 8334
	ErrorDsBadNameSyntax                                                    = 8335
	ErrorDsAliasPointsToAlias                                               = 8336
	ErrorDsCantDerefAlias                                                   = 8337
	ErrorDsOutOfScope                                                       = 8338
	ErrorDsObjectBeingRemoved                                               = 8339
	ErrorDsCantDeleteDsaObj                                                 = 8340
	ErrorDsGenericError                                                     = 8341
	ErrorDsDsaMustBeIntMaster                                               = 8342
	ErrorDsClassNotDsa                                                      = 8343
	ErrorDsInsuffAccessRights                                               = 8344
	ErrorDsIllegalSuperior                                                  = 8345
	ErrorDsAttributeOwnedBySam                                              = 8346
	ErrorDsNameTooManyParts                                                 = 8347
	ErrorDsNameTooLong                                                      = 8348
	ErrorDsNameValueTooLong                                                 = 8349
	ErrorDsNameUnparseable                                                  = 8350
	ErrorDsNameTypeUnknown                                                  = 8351
	ErrorDsNotAnObject                                                      = 8352
	ErrorDsSecDescTooShort                                                  = 8353
	ErrorDsSecDescInvalid                                                   = 8354
	ErrorDsNoDeletedName                                                    = 8355
	ErrorDsSubrefMustHaveParent                                             = 8356
	ErrorDsNcnameMustBeNc                                                   = 8357
	ErrorDsCantAddSystemOnly                                                = 8358
	ErrorDsClassMustBeConcrete                                              = 8359
	ErrorDsInvalidDmd                                                       = 8360
	ErrorDsObjGuidExists                                                    = 8361
	ErrorDsNotOnBacklink                                                    = 8362
	ErrorDsNoCrossrefForNc                                                  = 8363
	ErrorDsShuttingDown                                                     = 8364
	ErrorDsUnknownOperation                                                 = 8365
	ErrorDsInvalidRoleOwner                                                 = 8366
	ErrorDsCouldntContactFsmo                                               = 8367
	ErrorDsCrossNcDnRename                                                  = 8368
	ErrorDsCantModSystemOnly                                                = 8369
	ErrorDsReplicatorOnly                                                   = 8370
	ErrorDsObjClassNotDefined                                               = 8371
	ErrorDsObjClassNotSubclass                                              = 8372
	ErrorDsNameReferenceInvalid                                             = 8373
	ErrorDsCrossRefExists                                                   = 8374
	ErrorDsCantDelMasterCrossref                                            = 8375
	ErrorDsSubtreeNotifyNotNcHead                                           = 8376
	ErrorDsNotifyFilterTooComplex                                           = 8377
	ErrorDsDupRdn                                                           = 8378
	ErrorDsDupOid                                                           = 8379
	ErrorDsDupMapiId                                                        = 8380
	ErrorDsDupSchemaIdGuid                                                  = 8381
	ErrorDsDupLdapDisplayName                                               = 8382
	ErrorDsSemanticAttTest                                                  = 8383
	ErrorDsSyntaxMismatch                                                   = 8384
	ErrorDsExistsInMustHave                                                 = 8385
	ErrorDsExistsInMayHave                                                  = 8386
	ErrorDsNonexistentMayHave                                               = 8387
	ErrorDsNonexistentMustHave                                              = 8388
	ErrorDsAuxClsTestFail                                                   = 8389
	ErrorDsNonexistentPossSup                                               = 8390
	ErrorDsSubClsTestFail                                                   = 8391
	ErrorDsBadRdnAttIdSyntax                                                = 8392
	ErrorDsExistsInAuxCls                                                   = 8393
	ErrorDsExistsInSubCls                                                   = 8394
	ErrorDsExistsInPossSup                                                  = 8395
	ErrorDsRecalcschemaFailed                                               = 8396
	ErrorDsTreeDeleteNotFinished                                            = 8397
	ErrorDsCantDelete                                                       = 8398
	ErrorDsAttSchemaReqId                                                   = 8399
	ErrorDsBadAttSchemaSyntax                                               = 8400
	ErrorDsCantCacheAtt                                                     = 8401
	ErrorDsCantCacheClass                                                   = 8402
	ErrorDsCantRemoveAttCache                                               = 8403
	ErrorDsCantRemoveClassCache                                             = 8404
	ErrorDsCantRetrieveDn                                                   = 8405
	ErrorDsMissingSupref                                                    = 8406
	ErrorDsCantRetrieveInstance                                             = 8407
	ErrorDsCodeInconsistency                                                = 8408
	ErrorDsDatabaseError                                                    = 8409
	ErrorDsGovernsidMissing                                                 = 8410
	ErrorDsMissingExpectedAtt                                               = 8411
	ErrorDsNcnameMissingCrRef                                               = 8412
	ErrorDsSecurityCheckingError                                            = 8413
	ErrorDsSchemaNotLoaded                                                  = 8414
	ErrorDsSchemaAllocFailed                                                = 8415
	ErrorDsAttSchemaReqSyntax                                               = 8416
	ErrorDsGcverifyError                                                    = 8417
	ErrorDsDraSchemaMismatch                                                = 8418
	ErrorDsCantFindDsaObj                                                   = 8419
	ErrorDsCantFindExpectedNc                                               = 8420
	ErrorDsCantFindNcInCache                                                = 8421
	ErrorDsCantRetrieveChild                                                = 8422
	ErrorDsSecurityIllegalModify                                            = 8423
	ErrorDsCantReplaceHiddenRec                                             = 8424
	ErrorDsBadHierarchyFile                                                 = 8425
	ErrorDsBuildHierarchyTableFailed                                        = 8426
	ErrorDsConfigParamMissing                                               = 8427
	ErrorDsCountingAbIndicesFailed                                          = 8428
	ErrorDsHierarchyTableMallocFailed                                       = 8429
	ErrorDsInternalFailure                                                  = 8430
	ErrorDsUnknownError                                                     = 8431
	ErrorDsRootRequiresClassTop                                             = 8432
	ErrorDsRefusingFsmoRoles                                                = 8433
	ErrorDsMissingFsmoSettings                                              = 8434
	ErrorDsUnableToSurrenderRoles                                           = 8435
	ErrorDsDraGeneric                                                       = 8436
	ErrorDsDraInvalidParameter                                              = 8437
	ErrorDsDraBusy                                                          = 8438
	ErrorDsDraBadDn                                                         = 8439
	ErrorDsDraBadNc                                                         = 8440
	ErrorDsDraDnExists                                                      = 8441
	ErrorDsDraInternalError                                                 = 8442
	ErrorDsDraInconsistentDit                                               = 8443
	ErrorDsDraConnectionFailed                                              = 8444
	ErrorDsDraBadInstanceType                                               = 8445
	ErrorDsDraOutOfMem                                                      = 8446
	ErrorDsDraMailProblem                                                   = 8447
	ErrorDsDraRefAlreadyExists                                              = 8448
	ErrorDsDraRefNotFound                                                   = 8449
	ErrorDsDraObjIsRepSource                                                = 8450
	ErrorDsDraDbError                                                       = 8451
	ErrorDsDraNoReplica                                                     = 8452
	ErrorDsDraAccessDenied                                                  = 8453
	ErrorDsDraNotSupported                                                  = 8454
	ErrorDsDraRpcCancelled                                                  = 8455
	ErrorDsDraSourceDisabled                                                = 8456
	ErrorDsDraSinkDisabled                                                  = 8457
	ErrorDsDraNameCollision                                                 = 8458
	ErrorDsDraSourceReinstalled                                             = 8459
	ErrorDsDraMissingParent                                                 = 8460
	ErrorDsDraPreempted                                                     = 8461
	ErrorDsDraAbandonSync                                                   = 8462
	ErrorDsDraShutdown                                                      = 8463
	ErrorDsDraIncompatiblePartialSet                                        = 8464
	ErrorDsDraSourceIsPartialReplica                                        = 8465
	ErrorDsDraExtnConnectionFailed                                          = 8466
	ErrorDsInstallSchemaMismatch                                            = 8467
	ErrorDsDupLinkId                                                        = 8468
	ErrorDsNameErrorResolving                                               = 8469
	ErrorDsNameErrorNotFound                                                = 8470
	ErrorDsNameErrorNotUnique                                               = 8471
	ErrorDsNameErrorNoMapping                                               = 8472
	ErrorDsNameErrorDomainOnly                                              = 8473
	ErrorDsNameErrorNoSyntacticalMapping                                    = 8474
	ErrorDsConstructedAttMod                                                = 8475
	ErrorDsWrongOmObjClass                                                  = 8476
	ErrorDsDraReplPending                                                   = 8477
	ErrorDsDsRequired                                                       = 8478
	ErrorDsInvalidLdapDisplayName                                           = 8479
	ErrorDsNonBaseSearch                                                    = 8480
	ErrorDsCantRetrieveAtts                                                 = 8481
	ErrorDsBacklinkWithoutLink                                              = 8482
	ErrorDsEpochMismatch                                                    = 8483
	ErrorDsSrcNameMismatch                                                  = 8484
	ErrorDsSrcAndDstNcIdentical                                             = 8485
	ErrorDsDstNcMismatch                                                    = 8486
	ErrorDsNotAuthoritiveForDstNc                                           = 8487
	ErrorDsSrcGuidMismatch                                                  = 8488
	ErrorDsCantMoveDeletedObject                                            = 8489
	ErrorDsPdcOperationInProgress                                           = 8490
	ErrorDsCrossDomainCleanupReqd                                           = 8491
	ErrorDsIllegalXdomMoveOperation                                         = 8492
	ErrorDsCantWithAcctGroupMembershps                                      = 8493
	ErrorDsNcMustHaveNcParent                                               = 8494
	ErrorDsCrImpossibleToValidate                                           = 8495
	ErrorDsDstDomainNotNative                                               = 8496
	ErrorDsMissingInfrastructureContainer                                   = 8497
	ErrorDsCantMoveAccountGroup                                             = 8498
	ErrorDsCantMoveResourceGroup                                            = 8499
	ErrorDsInvalidSearchFlag                                                = 8500
	ErrorDsNoTreeDeleteAboveNc                                              = 8501
	ErrorDsCouldntLockTreeForDelete                                         = 8502
	ErrorDsCouldntIdentifyObjectsForTreeDelete                              = 8503
	ErrorDsSamInitFailure                                                   = 8504
	ErrorDsSensitiveGroupViolation                                          = 8505
	ErrorDsCantModPrimarygroupid                                            = 8506
	ErrorDsIllegalBaseSchemaMod                                             = 8507
	ErrorDsNonsafeSchemaChange                                              = 8508
	ErrorDsSchemaUpdateDisallowed                                           = 8509
	ErrorDsCantCreateUnderSchema                                            = 8510
	ErrorDsInstallNoSrcSchVersion                                           = 8511
	ErrorDsInstallNoSchVersionInInifile                                     = 8512
	ErrorDsInvalidGroupType                                                 = 8513
	ErrorDsNoNestGlobalgroupInMixeddomain                                   = 8514
	ErrorDsNoNestLocalgroupInMixeddomain                                    = 8515
	ErrorDsGlobalCantHaveLocalMember                                        = 8516
	ErrorDsGlobalCantHaveUniversalMember                                    = 8517
	ErrorDsUniversalCantHaveLocalMember                                     = 8518
	ErrorDsGlobalCantHaveCrossdomainMember                                  = 8519
	ErrorDsLocalCantHaveCrossdomainLocalMember                              = 8520
	ErrorDsHavePrimaryMembers                                               = 8521
	ErrorDsStringSdConversionFailed                                         = 8522
	ErrorDsNamingMasterGc                                                   = 8523
	ErrorDsDnsLookupFailure                                                 = 8524
	ErrorDsCouldntUpdateSpns                                                = 8525
	ErrorDsCantRetrieveSd                                                   = 8526
	ErrorDsKeyNotUnique                                                     = 8527
	ErrorDsWrongLinkedAttSyntax                                             = 8528
	ErrorDsSamNeedBootkeyPassword                                           = 8529
	ErrorDsSamNeedBootkeyFloppy                                             = 8530
	ErrorDsCantStart                                                        = 8531
	ErrorDsInitFailure                                                      = 8532
	ErrorDsNoPktPrivacyOnConnection                                         = 8533
	ErrorDsSourceDomainInForest                                             = 8534
	ErrorDsDestinationDomainNotInForest                                     = 8535
	ErrorDsDestinationAuditingNotEnabled                                    = 8536
	ErrorDsCantFindDcForSrcDomain                                           = 8537
	ErrorDsSrcObjNotGroupOrUser                                             = 8538
	ErrorDsSrcSidExistsInForest                                             = 8539
	ErrorDsSrcAndDstObjectClassMismatch                                     = 8540
	ErrorSamInitFailure                                                     = 8541
	ErrorDsDraSchemaInfoShip                                                = 8542
	ErrorDsDraSchemaConflict                                                = 8543
	ErrorDsDraEarlierSchemaConflict                                         = 8544
	ErrorDsDraObjNcMismatch                                                 = 8545
	ErrorDsNcStillHasDsas                                                   = 8546
	ErrorDsGcRequired                                                       = 8547
	ErrorDsLocalMemberOfLocalOnly                                           = 8548
	ErrorDsNoFpoInUniversalGroups                                           = 8549
	ErrorDsCantAddToGc                                                      = 8550
	ErrorDsNoCheckpointWithPdc                                              = 8551
	ErrorDsSourceAuditingNotEnabled                                         = 8552
	ErrorDsCantCreateInNondomainNc                                          = 8553
	ErrorDsInvalidNameForSpn                                                = 8554
	ErrorDsFilterUsesContructedAttrs                                        = 8555
	ErrorDsUnicodepwdNotInQuotes                                            = 8556
	ErrorDsMachineAccountQuotaExceeded                                      = 8557
	ErrorDsMustBeRunOnDstDc                                                 = 8558
	ErrorDsSrcDcMustBeSp4OrGreater                                          = 8559
	ErrorDsCantTreeDeleteCriticalObj                                        = 8560
	ErrorDsInitFailureConsole                                               = 8561
	ErrorDsSamInitFailureConsole                                            = 8562
	ErrorDsForestVersionTooHigh                                             = 8563
	ErrorDsDomainVersionTooHigh                                             = 8564
	ErrorDsForestVersionTooLow                                              = 8565
	ErrorDsDomainVersionTooLow                                              = 8566
	ErrorDsIncompatibleVersion                                              = 8567
	ErrorDsLowDsaVersion                                                    = 8568
	ErrorDsNoBehaviorVersionInMixeddomain                                   = 8569
	ErrorDsNotSupportedSortOrder                                            = 8570
	ErrorDsNameNotUnique                                                    = 8571
	ErrorDsMachineAccountCreatedPrent4                                      = 8572
	ErrorDsOutOfVersionStore                                                = 8573
	ErrorDsIncompatibleControlsUsed                                         = 8574
	ErrorDsNoRefDomain                                                      = 8575
	ErrorDsReservedLinkId                                                   = 8576
	ErrorDsLinkIdNotAvailable                                               = 8577
	ErrorDsAgCantHaveUniversalMember                                        = 8578
	ErrorDsModifydnDisallowedByInstanceType                                 = 8579
	ErrorDsNoObjectMoveInSchemaNc                                           = 8580
	ErrorDsModifydnDisallowedByFlag                                         = 8581
	ErrorDsModifydnWrongGrandparent                                         = 8582
	ErrorDsNameErrorTrustReferral                                           = 8583
	ErrorNotSupportedOnStandardServer                                       = 8584
	ErrorDsCantAccessRemotePartOfAd                                         = 8585
	ErrorDsCrImpossibleToValidateV2                                         = 8586
	ErrorDsThreadLimitExceeded                                              = 8587
	ErrorDsNotClosest                                                       = 8588
	ErrorDsCantDeriveSpnWithoutServerRef                                    = 8589
	ErrorDsSingleUserModeFailed                                             = 8590
	ErrorDsNtdscriptSyntaxError                                             = 8591
	ErrorDsNtdscriptProcessError                                            = 8592
	ErrorDsDifferentReplEpochs                                              = 8593
	ErrorDsDrsExtensionsChanged                                             = 8594
	ErrorDsReplicaSetChangeNotAllowedOnDisabledCr                           = 8595
	ErrorDsNoMsdsIntid                                                      = 8596
	ErrorDsDupMsdsIntid                                                     = 8597
	ErrorDsExistsInRdnattid                                                 = 8598
	ErrorDsAuthorizationFailed                                              = 8599
	ErrorDsInvalidScript                                                    = 8600
	ErrorDsRemoteCrossrefOpFailed                                           = 8601
	ErrorDsCrossRefBusy                                                     = 8602
	ErrorDsCantDeriveSpnForDeletedDomain                                    = 8603
	ErrorDsCantDemoteWithWriteableNc                                        = 8604
	ErrorDsDuplicateIdFound                                                 = 8605
	ErrorDsInsufficientAttrToCreateObject                                   = 8606
	ErrorDsGroupConversionError                                             = 8607
	ErrorDsCantMoveAppBasicGroup                                            = 8608
	ErrorDsCantMoveAppQueryGroup                                            = 8609
	ErrorDsRoleNotVerified                                                  = 8610
	ErrorDsWkoContainerCannotBeSpecial                                      = 8611
	ErrorDsDomainRenameInProgress                                           = 8612
	ErrorDsExistingAdChildNc                                                = 8613
	ErrorDsReplLifetimeExceeded                                             = 8614
	ErrorDsDisallowedInSystemContainer                                      = 8615
	ErrorDsLdapSendQueueFull                                                = 8616
	ErrorDsDraOutScheduleWindow                                             = 8617
	ErrorDsPolicyNotKnown                                                   = 8618
	ErrorNoSiteSettingsObject                                               = 8619
	ErrorNoSecrets                                                          = 8620
	ErrorNoWritableDcFound                                                  = 8621
	ErrorDsNoServerObject                                                   = 8622
	ErrorDsNoNtdsaObject                                                    = 8623
	ErrorDsNonAsqSearch                                                     = 8624
	ErrorDsAuditFailure                                                     = 8625
	ErrorDsInvalidSearchFlagSubtree                                         = 8626
	ErrorDsInvalidSearchFlagTuple                                           = 8627
	ErrorDsHierarchyTableTooDeep                                            = 8628
	ErrorDsDraCorruptUtdVector                                              = 8629
	ErrorDsDraSecretsDenied                                                 = 8630
	ErrorDsReservedMapiId                                                   = 8631
	ErrorDsMapiIdNotAvailable                                               = 8632
	ErrorDsDraMissingKrbtgtSecret                                           = 8633
	ErrorDsDomainNameExistsInForest                                         = 8634
	ErrorDsFlatNameExistsInForest                                           = 8635
	ErrorInvalidUserPrincipalName                                           = 8636
	ErrorDsOidMappedGroupCantHaveMembers                                    = 8637
	ErrorDsOidNotFound                                                      = 8638
	ErrorDsDraRecycledTarget                                                = 8639
	DnsErrorResponseCodesBase                                               = 9000
	DnsErrorMask                                                            = 0x00002328
	DnsErrorRcodeFormatError                                                = 9001
	DnsErrorRcodeServerFailure                                              = 9002
	DnsErrorRcodeNameError                                                  = 9003
	DnsErrorRcodeNotImplemented                                             = 9004
	DnsErrorRcodeRefused                                                    = 9005
	DnsErrorRcodeYxdomain                                                   = 9006
	DnsErrorRcodeYxrrset                                                    = 9007
	DnsErrorRcodeNxrrset                                                    = 9008
	DnsErrorRcodeNotauth                                                    = 9009
	DnsErrorRcodeNotzone                                                    = 9010
	DnsErrorRcodeBadsig                                                     = 9016
	DnsErrorRcodeBadkey                                                     = 9017
	DnsErrorRcodeBadtime                                                    = 9018
	DnsErrorPacketFmtBase                                                   = 9500
	DnsInfoNoRecords                                                        = 9501
	DnsErrorBadPacket                                                       = 9502
	DnsErrorNoPacket                                                        = 9503
	DnsErrorRcode                                                           = 9504
	DnsErrorUnsecurePacket                                                  = 9505
	DnsErrorGeneralApiBase                                                  = 9550
	DnsErrorInvalidType                                                     = 9551
	DnsErrorInvalidIpAddress                                                = 9552
	DnsErrorInvalidProperty                                                 = 9553
	DnsErrorTryAgainLater                                                   = 9554
	DnsErrorNotUnique                                                       = 9555
	DnsErrorNonRfcName                                                      = 9556
	DnsStatusFqdn                                                           = 9557
	DnsStatusDottedName                                                     = 9558
	DnsStatusSinglePartName                                                 = 9559
	DnsErrorInvalidNameChar                                                 = 9560
	DnsErrorNumericName                                                     = 9561
	DnsErrorNotAllowedOnRootServer                                          = 9562
	DnsErrorNotAllowedUnderDelegation                                       = 9563
	DnsErrorCannotFindRootHints                                             = 9564
	DnsErrorInconsistentRootHints                                           = 9565
	DnsErrorDwordValueTooSmall                                              = 9566
	DnsErrorDwordValueTooLarge                                              = 9567
	DnsErrorBackgroundLoading                                               = 9568
	DnsErrorNotAllowedOnRodc                                                = 9569
	DnsErrorNotAllowedUnderDname                                            = 9570
	DnsErrorDelegationRequired                                              = 9571
	DnsErrorInvalidPolicyTable                                              = 9572
	DnsErrorZoneBase                                                        = 9600
	DnsErrorZoneDoesNotExist                                                = 9601
	DnsErrorNoZoneInfo                                                      = 9602
	DnsErrorInvalidZoneOperation                                            = 9603
	DnsErrorZoneConfigurationError                                          = 9604
	DnsErrorZoneHasNoSoaRecord                                              = 9605
	DnsErrorZoneHasNoNsRecords                                              = 9606
	DnsErrorZoneLocked                                                      = 9607
	DnsErrorZoneCreationFailed                                              = 9608
	DnsErrorZoneAlreadyExists                                               = 9609
	DnsErrorAutozoneAlreadyExists                                           = 9610
	DnsErrorInvalidZoneType                                                 = 9611
	DnsErrorSecondaryRequiresMasterIp                                       = 9612
	DnsErrorZoneNotSecondary                                                = 9613
	DnsErrorNeedSecondaryAddresses                                          = 9614
	DnsErrorWinsInitFailed                                                  = 9615
	DnsErrorNeedWinsServers                                                 = 9616
	DnsErrorNbstatInitFailed                                                = 9617
	DnsErrorSoaDeleteInvalid                                                = 9618
	DnsErrorForwarderAlreadyExists                                          = 9619
	DnsErrorZoneRequiresMasterIp                                            = 9620
	DnsErrorZoneIsShutdown                                                  = 9621
	DnsErrorDatafileBase                                                    = 9650
	DnsErrorPrimaryRequiresDatafile                                         = 9651
	DnsErrorInvalidDatafileName                                             = 9652
	DnsErrorDatafileOpenFailure                                             = 9653
	DnsErrorFileWritebackFailed                                             = 9654
	DnsErrorDatafileParsing                                                 = 9655
	DnsErrorDatabaseBase                                                    = 9700
	DnsErrorRecordDoesNotExist                                              = 9701
	DnsErrorRecordFormat                                                    = 9702
	DnsErrorNodeCreationFailed                                              = 9703
	DnsErrorUnknownRecordType                                               = 9704
	DnsErrorRecordTimedOut                                                  = 9705
	DnsErrorNameNotInZone                                                   = 9706
	DnsErrorCnameLoop                                                       = 9707
	DnsErrorNodeIsCname                                                     = 9708
	DnsErrorCnameCollision                                                  = 9709
	DnsErrorRecordOnlyAtZoneRoot                                            = 9710
	DnsErrorRecordAlreadyExists                                             = 9711
	DnsErrorSecondaryData                                                   = 9712
	DnsErrorNoCreateCacheData                                               = 9713
	DnsErrorNameDoesNotExist                                                = 9714
	DnsWarningPtrCreateFailed                                               = 9715
	DnsWarningDomainUndeleted                                               = 9716
	DnsErrorDsUnavailable                                                   = 9717
	DnsErrorDsZoneAlreadyExists                                             = 9718
	DnsErrorNoBootfileIfDsZone                                              = 9719
	DnsErrorNodeIsDname                                                     = 9720
	DnsErrorDnameCollision                                                  = 9721
	DnsErrorAliasLoop                                                       = 9722
	DnsErrorOperationBase                                                   = 9750
	DnsInfoAxfrComplete                                                     = 9751
	DnsErrorAxfr                                                            = 9752
	DnsInfoAddedLocalWins                                                   = 9753
	DnsErrorSecureBase                                                      = 9800
	DnsStatusContinueNeeded                                                 = 9801
	DnsErrorSetupBase                                                       = 9850
	DnsErrorNoTcpip                                                         = 9851
	DnsErrorNoDnsServers                                                    = 9852
	DnsErrorDpBase                                                          = 9900
	DnsErrorDpDoesNotExist                                                  = 9901
	DnsErrorDpAlreadyExists                                                 = 9902
	DnsErrorDpNotEnlisted                                                   = 9903
	DnsErrorDpAlreadyEnlisted                                               = 9904
	DnsErrorDpNotAvailable                                                  = 9905
	DnsErrorDpFsmoError                                                     = 9906
	Wsabaseerr                                                              = 10000
	Wsaeintr                                                                = 10004
	Wsaebadf                                                                = 10009
	Wsaeacces                                                               = 10013
	Wsaefault                                                               = 10014
	Wsaeinval                                                               = 10022
	Wsaemfile                                                               = 10024
	Wsaewouldblock                                                          = 10035
	Wsaeinprogress                                                          = 10036
	Wsaealready                                                             = 10037
	Wsaenotsock                                                             = 10038
	Wsaedestaddrreq                                                         = 10039
	Wsaemsgsize                                                             = 10040
	Wsaeprototype                                                           = 10041
	Wsaenoprotoopt                                                          = 10042
	Wsaeprotonosupport                                                      = 10043
	Wsaesocktnosupport                                                      = 10044
	Wsaeopnotsupp                                                           = 10045
	Wsaepfnosupport                                                         = 10046
	Wsaeafnosupport                                                         = 10047
	Wsaeaddrinuse                                                           = 10048
	Wsaeaddrnotavail                                                        = 10049
	Wsaenetdown                                                             = 10050
	Wsaenetunreach                                                          = 10051
	Wsaenetreset                                                            = 10052
	Wsaeconnaborted                                                         = 10053
	Wsaeconnreset                                                           = 10054
	Wsaenobufs                                                              = 10055
	Wsaeisconn                                                              = 10056
	Wsaenotconn                                                             = 10057
	Wsaeshutdown                                                            = 10058
	Wsaetoomanyrefs                                                         = 10059
	Wsaetimedout                                                            = 10060
	Wsaeconnrefused                                                         = 10061
	Wsaeloop                                                                = 10062
	Wsaenametoolong                                                         = 10063
	Wsaehostdown                                                            = 10064
	Wsaehostunreach                                                         = 10065
	Wsaenotempty                                                            = 10066
	Wsaeproclim                                                             = 10067
	Wsaeusers                                                               = 10068
	Wsaedquot                                                               = 10069
	Wsaestale                                                               = 10070
	Wsaeremote                                                              = 10071
	Wsasysnotready                                                          = 10091
	Wsavernotsupported                                                      = 10092
	Wsanotinitialised                                                       = 10093
	Wsaediscon                                                              = 10101
	Wsaenomore                                                              = 10102
	Wsaecancelled                                                           = 10103
	Wsaeinvalidproctable                                                    = 10104
	Wsaeinvalidprovider                                                     = 10105
	Wsaeproviderfailedinit                                                  = 10106
	Wsasyscallfailure                                                       = 10107
	WsaserviceNotFound                                                      = 10108
	WsatypeNotFound                                                         = 10109
	WsaENoMore                                                              = 10110
	WsaECancelled                                                           = 10111
	Wsaerefused                                                             = 10112
	WsahostNotFound                                                         = 11001
	WsatryAgain                                                             = 11002
	WsanoRecovery                                                           = 11003
	WsanoData                                                               = 11004
	WsaQosReceivers                                                         = 11005
	WsaQosSenders                                                           = 11006
	WsaQosNoSenders                                                         = 11007
	WsaQosNoReceivers                                                       = 11008
	WsaQosRequestConfirmed                                                  = 11009
	WsaQosAdmissionFailure                                                  = 11010
	WsaQosPolicyFailure                                                     = 11011
	WsaQosBadStyle                                                          = 11012
	WsaQosBadObject                                                         = 11013
	WsaQosTrafficCtrlError                                                  = 11014
	WsaQosGenericError                                                      = 11015
	WsaQosEservicetype                                                      = 11016
	WsaQosEflowspec                                                         = 11017
	WsaQosEprovspecbuf                                                      = 11018
	WsaQosEfilterstyle                                                      = 11019
	WsaQosEfiltertype                                                       = 11020
	WsaQosEfiltercount                                                      = 11021
	WsaQosEobjlength                                                        = 11022
	WsaQosEflowcount                                                        = 11023
	WsaQosEunkownpsobj                                                      = 11024
	WsaQosEpolicyobj                                                        = 11025
	WsaQosEflowdesc                                                         = 11026
	WsaQosEpsflowspec                                                       = 11027
	WsaQosEpsfilterspec                                                     = 11028
	WsaQosEsdmodeobj                                                        = 11029
	WsaQosEshaperateobj                                                     = 11030
	WsaQosReservedPetype                                                    = 11031
	WsaSecureHostNotFound                                                   = 11032
	WsaIpsecNamePolicyError                                                 = 11033
	ErrorIpsecQmPolicyExists                                                = 13000
	ErrorIpsecQmPolicyNotFound                                              = 13001
	ErrorIpsecQmPolicyInUse                                                 = 13002
	ErrorIpsecMmPolicyExists                                                = 13003
	ErrorIpsecMmPolicyNotFound                                              = 13004
	ErrorIpsecMmPolicyInUse                                                 = 13005
	ErrorIpsecMmFilterExists                                                = 13006
	ErrorIpsecMmFilterNotFound                                              = 13007
	ErrorIpsecTransportFilterExists                                         = 13008
	ErrorIpsecTransportFilterNotFound                                       = 13009
	ErrorIpsecMmAuthExists                                                  = 13010
	ErrorIpsecMmAuthNotFound                                                = 13011
	ErrorIpsecMmAuthInUse                                                   = 13012
	ErrorIpsecDefaultMmPolicyNotFound                                       = 13013
	ErrorIpsecDefaultMmAuthNotFound                                         = 13014
	ErrorIpsecDefaultQmPolicyNotFound                                       = 13015
	ErrorIpsecTunnelFilterExists                                            = 13016
	ErrorIpsecTunnelFilterNotFound                                          = 13017
	ErrorIpsecMmFilterPendingDeletion                                       = 13018
	ErrorIpsecTransportFilterPendingDeletion                                = 13019
	ErrorIpsecTunnelFilterPendingDeletion                                   = 13020
	ErrorIpsecMmPolicyPendingDeletion                                       = 13021
	ErrorIpsecMmAuthPendingDeletion                                         = 13022
	ErrorIpsecQmPolicyPendingDeletion                                       = 13023
	WarningIpsecMmPolicyPruned                                              = 13024
	WarningIpsecQmPolicyPruned                                              = 13025
	ErrorIpsecIkeNegStatusBegin                                             = 13800
	ErrorIpsecIkeAuthFail                                                   = 13801
	ErrorIpsecIkeAttribFail                                                 = 13802
	ErrorIpsecIkeNegotiationPending                                         = 13803
	ErrorIpsecIkeGeneralProcessingError                                     = 13804
	ErrorIpsecIkeTimedOut                                                   = 13805
	ErrorIpsecIkeNoCert                                                     = 13806
	ErrorIpsecIkeSaDeleted                                                  = 13807
	ErrorIpsecIkeSaReaped                                                   = 13808
	ErrorIpsecIkeMmAcquireDrop                                              = 13809
	ErrorIpsecIkeQmAcquireDrop                                              = 13810
	ErrorIpsecIkeQueueDropMm                                                = 13811
	ErrorIpsecIkeQueueDropNoMm                                              = 13812
	ErrorIpsecIkeDropNoResponse                                             = 13813
	ErrorIpsecIkeMmDelayDrop                                                = 13814
	ErrorIpsecIkeQmDelayDrop                                                = 13815
	ErrorIpsecIkeError                                                      = 13816
	ErrorIpsecIkeCrlFailed                                                  = 13817
	ErrorIpsecIkeInvalidKeyUsage                                            = 13818
	ErrorIpsecIkeInvalidCertType                                            = 13819
	ErrorIpsecIkeNoPrivateKey                                               = 13820
	ErrorIpsecIkeSimultaneousRekey                                          = 13821
	ErrorIpsecIkeDhFail                                                     = 13822
	ErrorIpsecIkeCriticalPayloadNotRecognized                               = 13823
	ErrorIpsecIkeInvalidHeader                                              = 13824
	ErrorIpsecIkeNoPolicy                                                   = 13825
	ErrorIpsecIkeInvalidSignature                                           = 13826
	ErrorIpsecIkeKerberosError                                              = 13827
	ErrorIpsecIkeNoPublicKey                                                = 13828
	ErrorIpsecIkeProcessErr                                                 = 13829
	ErrorIpsecIkeProcessErrSa                                               = 13830
	ErrorIpsecIkeProcessErrProp                                             = 13831
	ErrorIpsecIkeProcessErrTrans                                            = 13832
	ErrorIpsecIkeProcessErrKe                                               = 13833
	ErrorIpsecIkeProcessErrId                                               = 13834
	ErrorIpsecIkeProcessErrCert                                             = 13835
	ErrorIpsecIkeProcessErrCertReq                                          = 13836
	ErrorIpsecIkeProcessErrHash                                             = 13837
	ErrorIpsecIkeProcessErrSig                                              = 13838
	ErrorIpsecIkeProcessErrNonce                                            = 13839
	ErrorIpsecIkeProcessErrNotify                                           = 13840
	ErrorIpsecIkeProcessErrDelete                                           = 13841
	ErrorIpsecIkeProcessErrVendor                                           = 13842
	ErrorIpsecIkeInvalidPayload                                             = 13843
	ErrorIpsecIkeLoadSoftSa                                                 = 13844
	ErrorIpsecIkeSoftSaTornDown                                             = 13845
	ErrorIpsecIkeInvalidCookie                                              = 13846
	ErrorIpsecIkeNoPeerCert                                                 = 13847
	ErrorIpsecIkePeerCrlFailed                                              = 13848
	ErrorIpsecIkePolicyChange                                               = 13849
	ErrorIpsecIkeNoMmPolicy                                                 = 13850
	ErrorIpsecIkeNotcbpriv                                                  = 13851
	ErrorIpsecIkeSecloadfail                                                = 13852
	ErrorIpsecIkeFailsspinit                                                = 13853
	ErrorIpsecIkeFailqueryssp                                               = 13854
	ErrorIpsecIkeSrvacqfail                                                 = 13855
	ErrorIpsecIkeSrvquerycred                                               = 13856
	ErrorIpsecIkeGetspifail                                                 = 13857
	ErrorIpsecIkeInvalidFilter                                              = 13858
	ErrorIpsecIkeOutOfMemory                                                = 13859
	ErrorIpsecIkeAddUpdateKeyFailed                                         = 13860
	ErrorIpsecIkeInvalidPolicy                                              = 13861
	ErrorIpsecIkeUnknownDoi                                                 = 13862
	ErrorIpsecIkeInvalidSituation                                           = 13863
	ErrorIpsecIkeDhFailure                                                  = 13864
	ErrorIpsecIkeInvalidGroup                                               = 13865
	ErrorIpsecIkeEncrypt                                                    = 13866
	ErrorIpsecIkeDecrypt                                                    = 13867
	ErrorIpsecIkePolicyMatch                                                = 13868
	ErrorIpsecIkeUnsupportedId                                              = 13869
	ErrorIpsecIkeInvalidHash                                                = 13870
	ErrorIpsecIkeInvalidHashAlg                                             = 13871
	ErrorIpsecIkeInvalidHashSize                                            = 13872
	ErrorIpsecIkeInvalidEncryptAlg                                          = 13873
	ErrorIpsecIkeInvalidAuthAlg                                             = 13874
	ErrorIpsecIkeInvalidSig                                                 = 13875
	ErrorIpsecIkeLoadFailed                                                 = 13876
	ErrorIpsecIkeRpcDelete                                                  = 13877
	ErrorIpsecIkeBenignReinit                                               = 13878
	ErrorIpsecIkeInvalidResponderLifetimeNotify                             = 13879
	ErrorIpsecIkeInvalidMajorVersion                                        = 13880
	ErrorIpsecIkeInvalidCertKeylen                                          = 13881
	ErrorIpsecIkeMmLimit                                                    = 13882
	ErrorIpsecIkeNegotiationDisabled                                        = 13883
	ErrorIpsecIkeQmLimit                                                    = 13884
	ErrorIpsecIkeMmExpired                                                  = 13885
	ErrorIpsecIkePeerMmAssumedInvalid                                       = 13886
	ErrorIpsecIkeCertChainPolicyMismatch                                    = 13887
	ErrorIpsecIkeUnexpectedMessageId                                        = 13888
	ErrorIpsecIkeInvalidAuthPayload                                         = 13889
	ErrorIpsecIkeDosCookieSent                                              = 13890
	ErrorIpsecIkeShuttingDown                                               = 13891
	ErrorIpsecIkeCgaAuthFailed                                              = 13892
	ErrorIpsecIkeProcessErrNatoa                                            = 13893
	ErrorIpsecIkeInvalidMmForQm                                             = 13894
	ErrorIpsecIkeQmExpired                                                  = 13895
	ErrorIpsecIkeTooManyFilters                                             = 13896
	ErrorIpsecIkeNegStatusEnd                                               = 13897
	ErrorIpsecIkeKillDummyNapTunnel                                         = 13898
	ErrorIpsecIkeInnerIpAssignmentFailure                                   = 13899
	ErrorIpsecIkeRequireCpPayloadMissing                                    = 13900
	ErrorIpsecKeyModuleImpersonationNegotiationPending                      = 13901
	ErrorIpsecIkeCoexistenceSuppress                                        = 13902
	ErrorIpsecIkeRatelimitDrop                                              = 13903
	ErrorIpsecIkePeerDoesntSupportMobike                                    = 13904
	ErrorIpsecIkeAuthorizationFailure                                       = 13905
	ErrorIpsecIkeStrongCredAuthorizationFailure                             = 13906
	ErrorIpsecIkeAuthorizationFailureWithOptionalRetry                      = 13907
	ErrorIpsecIkeStrongCredAuthorizationAndCertmapFailure                   = 13908
	ErrorIpsecIkeNegStatusExtendedEnd                                       = 13909
	ErrorIpsecBadSpi                                                        = 13910
	ErrorIpsecSaLifetimeExpired                                             = 13911
	ErrorIpsecWrongSa                                                       = 13912
	ErrorIpsecReplayCheckFailed                                             = 13913
	ErrorIpsecInvalidPacket                                                 = 13914
	ErrorIpsecIntegrityCheckFailed                                          = 13915
	ErrorIpsecClearTextDrop                                                 = 13916
	ErrorIpsecAuthFirewallDrop                                              = 13917
	ErrorIpsecThrottleDrop                                                  = 13918
	ErrorIpsecDospBlock                                                     = 13925
	ErrorIpsecDospReceivedMulticast                                         = 13926
	ErrorIpsecDospInvalidPacket                                             = 13927
	ErrorIpsecDospStateLookupFailed                                         = 13928
	ErrorIpsecDospMaxEntries                                                = 13929
	ErrorIpsecDospKeymodNotAllowed                                          = 13930
	ErrorIpsecDospNotInstalled                                              = 13931
	ErrorIpsecDospMaxPerIpRatelimitQueues                                   = 13932
	ErrorSxsSectionNotFound                                                 = 14000
	ErrorSxsCantGenActctx                                                   = 14001
	ErrorSxsInvalidActctxdataFormat                                         = 14002
	ErrorSxsAssemblyNotFound                                                = 14003
	ErrorSxsManifestFormatError                                             = 14004
	ErrorSxsManifestParseError                                              = 14005
	ErrorSxsActivationContextDisabled                                       = 14006
	ErrorSxsKeyNotFound                                                     = 14007
	ErrorSxsVersionConflict                                                 = 14008
	ErrorSxsWrongSectionType                                                = 14009
	ErrorSxsThreadQueriesDisabled                                           = 14010
	ErrorSxsProcessDefaultAlreadySet                                        = 14011
	ErrorSxsUnknownEncodingGroup                                            = 14012
	ErrorSxsUnknownEncoding                                                 = 14013
	ErrorSxsInvalidXmlNamespaceUri                                          = 14014
	ErrorSxsRootManifestDependencyNotInstalled                              = 14015
	ErrorSxsLeafManifestDependencyNotInstalled                              = 14016
	ErrorSxsInvalidAssemblyIdentityAttribute                                = 14017
	ErrorSxsManifestMissingRequiredDefaultNamespace                         = 14018
	ErrorSxsManifestInvalidRequiredDefaultNamespace                         = 14019
	ErrorSxsPrivateManifestCrossPathWithReparsePoint                        = 14020
	ErrorSxsDuplicateDllName                                                = 14021
	ErrorSxsDuplicateWindowclassName                                        = 14022
	ErrorSxsDuplicateClsid                                                  = 14023
	ErrorSxsDuplicateIid                                                    = 14024
	ErrorSxsDuplicateTlbid                                                  = 14025
	ErrorSxsDuplicateProgid                                                 = 14026
	ErrorSxsDuplicateAssemblyName                                           = 14027
	ErrorSxsFileHashMismatch                                                = 14028
	ErrorSxsPolicyParseError                                                = 14029
	ErrorSxsXmlEMissingquote                                                = 14030
	ErrorSxsXmlECommentsyntax                                               = 14031
	ErrorSxsXmlEBadstartnamechar                                            = 14032
	ErrorSxsXmlEBadnamechar                                                 = 14033
	ErrorSxsXmlEBadcharinstring                                             = 14034
	ErrorSxsXmlEXmldeclsyntax                                               = 14035
	ErrorSxsXmlEBadchardata                                                 = 14036
	ErrorSxsXmlEMissingwhitespace                                           = 14037
	ErrorSxsXmlEExpectingtagend                                             = 14038
	ErrorSxsXmlEMissingsemicolon                                            = 14039
	ErrorSxsXmlEUnbalancedparen                                             = 14040
	ErrorSxsXmlEInternalerror                                               = 14041
	ErrorSxsXmlEUnexpectedWhitespace                                        = 14042
	ErrorSxsXmlEIncompleteEncoding                                          = 14043
	ErrorSxsXmlEMissingParen                                                = 14044
	ErrorSxsXmlEExpectingclosequote                                         = 14045
	ErrorSxsXmlEMultipleColons                                              = 14046
	ErrorSxsXmlEInvalidDecimal                                              = 14047
	ErrorSxsXmlEInvalidHexidecimal                                          = 14048
	ErrorSxsXmlEInvalidUnicode                                              = 14049
	ErrorSxsXmlEWhitespaceorquestionmark                                    = 14050
	ErrorSxsXmlEUnexpectedendtag                                            = 14051
	ErrorSxsXmlEUnclosedtag                                                 = 14052
	ErrorSxsXmlEDuplicateattribute                                          = 14053
	ErrorSxsXmlEMultipleroots                                               = 14054
	ErrorSxsXmlEInvalidatrootlevel                                          = 14055
	ErrorSxsXmlEBadxmldecl                                                  = 14056
	ErrorSxsXmlEMissingroot                                                 = 14057
	ErrorSxsXmlEUnexpectedeof                                               = 14058
	ErrorSxsXmlEBadperefinsubset                                            = 14059
	ErrorSxsXmlEUnclosedstarttag                                            = 14060
	ErrorSxsXmlEUnclosedendtag                                              = 14061
	ErrorSxsXmlEUnclosedstring                                              = 14062
	ErrorSxsXmlEUnclosedcomment                                             = 14063
	ErrorSxsXmlEUncloseddecl                                                = 14064
	ErrorSxsXmlEUnclosedcdata                                               = 14065
	ErrorSxsXmlEReservednamespace                                           = 14066
	ErrorSxsXmlEInvalidencoding                                             = 14067
	ErrorSxsXmlEInvalidswitch                                               = 14068
	ErrorSxsXmlEBadxmlcase                                                  = 14069
	ErrorSxsXmlEInvalidStandalone                                           = 14070
	ErrorSxsXmlEUnexpectedStandalone                                        = 14071
	ErrorSxsXmlEInvalidVersion                                              = 14072
	ErrorSxsXmlEMissingequals                                               = 14073
	ErrorSxsProtectionRecoveryFailed                                        = 14074
	ErrorSxsProtectionPublicKeyTooShort                                     = 14075
	ErrorSxsProtectionCatalogNotValid                                       = 14076
	ErrorSxsUntranslatableHresult                                           = 14077
	ErrorSxsProtectionCatalogFileMissing                                    = 14078
	ErrorSxsMissingAssemblyIdentityAttribute                                = 14079
	ErrorSxsInvalidAssemblyIdentityAttributeName                            = 14080
	ErrorSxsAssemblyMissing                                                 = 14081
	ErrorSxsCorruptActivationStack                                          = 14082
	ErrorSxsCorruption                                                      = 14083
	ErrorSxsEarlyDeactivation                                               = 14084
	ErrorSxsInvalidDeactivation                                             = 14085
	ErrorSxsMultipleDeactivation                                            = 14086
	ErrorSxsProcessTerminationRequested                                     = 14087
	ErrorSxsReleaseActivationContext                                        = 14088
	ErrorSxsSystemDefaultActivationContextEmpty                             = 14089
	ErrorSxsInvalidIdentityAttributeValue                                   = 14090
	ErrorSxsInvalidIdentityAttributeName                                    = 14091
	ErrorSxsIdentityDuplicateAttribute                                      = 14092
	ErrorSxsIdentityParseError                                              = 14093
	ErrorMalformedSubstitutionString                                        = 14094
	ErrorSxsIncorrectPublicKeyToken                                         = 14095
	ErrorUnmappedSubstitutionString                                         = 14096
	ErrorSxsAssemblyNotLocked                                               = 14097
	ErrorSxsComponentStoreCorrupt                                           = 14098
	ErrorAdvancedInstallerFailed                                            = 14099
	ErrorXmlEncodingMismatch                                                = 14100
	ErrorSxsManifestIdentitySameButContentsDifferent                        = 14101
	ErrorSxsIdentitiesDifferent                                             = 14102
	ErrorSxsAssemblyIsNotADeployment                                        = 14103
	ErrorSxsFileNotPartOfAssembly                                           = 14104
	ErrorSxsManifestTooBig                                                  = 14105
	ErrorSxsSettingNotRegistered                                            = 14106
	ErrorSxsTransactionClosureIncomplete                                    = 14107
	ErrorSmiPrimitiveInstallerFailed                                        = 14108
	ErrorGenericCommandFailed                                               = 14109
	ErrorSxsFileHashMissing                                                 = 14110
	ErrorEvtInvalidChannelPath                                              = 15000
	ErrorEvtInvalidQuery                                                    = 15001
	ErrorEvtPublisherMetadataNotFound                                       = 15002
	ErrorEvtEventTemplateNotFound                                           = 15003
	ErrorEvtInvalidPublisherName                                            = 15004
	ErrorEvtInvalidEventData                                                = 15005
	ErrorEvtChannelNotFound                                                 = 15007
	ErrorEvtMalformedXmlText                                                = 15008
	ErrorEvtSubscriptionToDirectChannel                                     = 15009
	ErrorEvtConfigurationError                                              = 15010
	ErrorEvtQueryResultStale                                                = 15011
	ErrorEvtQueryResultInvalidPosition                                      = 15012
	ErrorEvtNonValidatingMsxml                                              = 15013
	ErrorEvtFilterAlreadyscoped                                             = 15014
	ErrorEvtFilterNoteltset                                                 = 15015
	ErrorEvtFilterInvarg                                                    = 15016
	ErrorEvtFilterInvtest                                                   = 15017
	ErrorEvtFilterInvtype                                                   = 15018
	ErrorEvtFilterParseerr                                                  = 15019
	ErrorEvtFilterUnsupportedop                                             = 15020
	ErrorEvtFilterUnexpectedtoken                                           = 15021
	ErrorEvtInvalidOperationOverEnabledDirectChannel                        = 15022
	ErrorEvtInvalidChannelPropertyValue                                     = 15023
	ErrorEvtInvalidPublisherPropertyValue                                   = 15024
	ErrorEvtChannelCannotActivate                                           = 15025
	ErrorEvtFilterTooComplex                                                = 15026
	ErrorEvtMessageNotFound                                                 = 15027
	ErrorEvtMessageIdNotFound                                               = 15028
	ErrorEvtUnresolvedValueInsert                                           = 15029
	ErrorEvtUnresolvedParameterInsert                                       = 15030
	ErrorEvtMaxInsertsReached                                               = 15031
	ErrorEvtEventDefinitionNotFound                                         = 15032
	ErrorEvtMessageLocaleNotFound                                           = 15033
	ErrorEvtVersionTooOld                                                   = 15034
	ErrorEvtVersionTooNew                                                   = 15035
	ErrorEvtCannotOpenChannelOfQuery                                        = 15036
	ErrorEvtPublisherDisabled                                               = 15037
	ErrorEvtFilterOutOfRange                                                = 15038
	ErrorEcSubscriptionCannotActivate                                       = 15080
	ErrorEcLogDisabled                                                      = 15081
	ErrorEcCircularForwarding                                               = 15082
	ErrorEcCredstoreFull                                                    = 15083
	ErrorEcCredNotFound                                                     = 15084
	ErrorEcNoActiveChannel                                                  = 15085
	ErrorMuiFileNotFound                                                    = 15100
	ErrorMuiInvalidFile                                                     = 15101
	ErrorMuiInvalidRcConfig                                                 = 15102
	ErrorMuiInvalidLocaleName                                               = 15103
	ErrorMuiInvalidUltimatefallbackName                                     = 15104
	ErrorMuiFileNotLoaded                                                   = 15105
	ErrorResourceEnumUserStop                                               = 15106
	ErrorMuiIntlsettingsUilangNotInstalled                                  = 15107
	ErrorMuiIntlsettingsInvalidLocaleName                                   = 15108
	ErrorMcaInvalidCapabilitiesString                                       = 15200
	ErrorMcaInvalidVcpVersion                                               = 15201
	ErrorMcaMonitorViolatesMccsSpecification                                = 15202
	ErrorMcaMccsVersionMismatch                                             = 15203
	ErrorMcaUnsupportedMccsVersion                                          = 15204
	ErrorMcaInternalError                                                   = 15205
	ErrorMcaInvalidTechnologyTypeReturned                                   = 15206
	ErrorMcaUnsupportedColorTemperature                                     = 15207
	ErrorAmbiguousSystemDevice                                              = 15250
	ErrorSystemDeviceNotFound                                               = 15299
	ErrorHashNotSupported                                                   = 15300
	ErrorHashNotPresent                                                     = 15301
	SeveritySuccess                                                         = 0
	SeverityError                                                           = 1
	FacilityNtBit                                                           = 0x10000000
	EUnexpected                                                             = 0x8000FFFF
	ENotimpl                                                                = 0x80004001
	EOutofmemory                                                            = 0x8007000E
	EInvalidarg                                                             = 0x80070057
	ENointerface                                                            = 0x80004002
	EPointer                                                                = 0x80004003
	EHandle                                                                 = 0x80070006
	EAbort                                                                  = 0x80004004
	EFail                                                                   = 0x80004005
	EAccessdenied                                                           = 0x80070005
	EPending                                                                = 0x8000000A
	CoEInitTls                                                              = 0x80004006
	CoEInitSharedAllocator                                                  = 0x80004007
	CoEInitMemoryAllocator                                                  = 0x80004008
	CoEInitClassCache                                                       = 0x80004009
	CoEInitRpcChannel                                                       = 0x8000400A
	CoEInitTlsSetChannelControl                                             = 0x8000400B
	CoEInitTlsChannelControl                                                = 0x8000400C
	CoEInitUnacceptedUserAllocator                                          = 0x8000400D
	CoEInitScmMutexExists                                                   = 0x8000400E
	CoEInitScmFileMappingExists                                             = 0x8000400F
	CoEInitScmMapViewOfFile                                                 = 0x80004010
	CoEInitScmExecFailure                                                   = 0x80004011
	CoEInitOnlySingleThreaded                                               = 0x80004012
	CoECantRemote                                                           = 0x80004013
	CoEBadServerName                                                        = 0x80004014
	CoEWrongServerIdentity                                                  = 0x80004015
	CoEOle1DdeDisabled                                                      = 0x80004016
	CoERunasSyntax                                                          = 0x80004017
	CoECreateprocessFailure                                                 = 0x80004018
	CoERunasCreateprocessFailure                                            = 0x80004019
	CoERunasLogonFailure                                                    = 0x8000401A
	CoELaunchPermssionDenied                                                = 0x8000401B
	CoEStartServiceFailure                                                  = 0x8000401C
	CoERemoteCommunicationFailure                                           = 0x8000401D
	CoEServerStartTimeout                                                   = 0x8000401E
	CoEClsregInconsistent                                                   = 0x8000401F
	CoEIidregInconsistent                                                   = 0x80004020
	CoENotSupported                                                         = 0x80004021
	CoEReloadDll                                                            = 0x80004022
	CoEMsiError                                                             = 0x80004023
	CoEAttemptToCreateOutsideClientContext                                  = 0x80004024
	CoEServerPaused                                                         = 0x80004025
	CoEServerNotPaused                                                      = 0x80004026
	CoEClassDisabled                                                        = 0x80004027
	CoEClrnotavailable                                                      = 0x80004028
	CoEAsyncWorkRejected                                                    = 0x80004029
	CoEServerInitTimeout                                                    = 0x8000402A
	CoENoSecctxInActivate                                                   = 0x8000402B
	CoETrackerConfig                                                        = 0x80004030
	CoEThreadpoolConfig                                                     = 0x80004031
	CoESxsConfig                                                            = 0x80004032
	CoEMalformedSpn                                                         = 0x80004033
	SOk                                                                     = 0
	SFalse                                                                  = 1
	OleEFirst                                                               = 0x80040000
	OleELast                                                                = 0x800400FF
	OleSFirst                                                               = 0x00040000
	OleSLast                                                                = 0x000400FF
	OleEOleverb                                                             = 0x80040000
	OleEAdvf                                                                = 0x80040001
	OleEEnumNomore                                                          = 0x80040002
	OleEAdvisenotsupported                                                  = 0x80040003
	OleENoconnection                                                        = 0x80040004
	OleENotrunning                                                          = 0x80040005
	OleENocache                                                             = 0x80040006
	OleEBlank                                                               = 0x80040007
	OleEClassdiff                                                           = 0x80040008
	OleECantGetmoniker                                                      = 0x80040009
	OleECantBindtosource                                                    = 0x8004000A
	OleEStatic                                                              = 0x8004000B
	OleEPromptsavecancelled                                                 = 0x8004000C
	OleEInvalidrect                                                         = 0x8004000D
	OleEWrongcompobj                                                        = 0x8004000E
	OleEInvalidhwnd                                                         = 0x8004000F
	OleENotInplaceactive                                                    = 0x80040010
	OleECantconvert                                                         = 0x80040011
	OleENostorage                                                           = 0x80040012
	DvEFormatetc                                                            = 0x80040064
	DvEDvtargetdevice                                                       = 0x80040065
	DvEStgmedium                                                            = 0x80040066
	DvEStatdata                                                             = 0x80040067
	DvELindex                                                               = 0x80040068
	DvETymed                                                                = 0x80040069
	DvEClipformat                                                           = 0x8004006A
	DvEDvaspect                                                             = 0x8004006B
	DvEDvtargetdeviceSize                                                   = 0x8004006C
	DvENoiviewobject                                                        = 0x8004006D
	DragdropEFirst                                                          = 0x80040100
	DragdropELast                                                           = 0x8004010F
	DragdropSFirst                                                          = 0x00040100
	DragdropSLast                                                           = 0x0004010F
	DragdropENotregistered                                                  = 0x80040100
	DragdropEAlreadyregistered                                              = 0x80040101
	DragdropEInvalidhwnd                                                    = 0x80040102
	ClassfactoryEFirst                                                      = 0x80040110
	ClassfactoryELast                                                       = 0x8004011F
	ClassfactorySFirst                                                      = 0x00040110
	ClassfactorySLast                                                       = 0x0004011F
	ClassENoaggregation                                                     = 0x80040110
	ClassEClassnotavailable                                                 = 0x80040111
	ClassENotlicensed                                                       = 0x80040112
	MarshalEFirst                                                           = 0x80040120
	MarshalELast                                                            = 0x8004012F
	MarshalSFirst                                                           = 0x00040120
	MarshalSLast                                                            = 0x0004012F
	DataEFirst                                                              = 0x80040130
	DataELast                                                               = 0x8004013F
	DataSFirst                                                              = 0x00040130
	DataSLast                                                               = 0x0004013F
	ViewEFirst                                                              = 0x80040140
	ViewELast                                                               = 0x8004014F
	ViewSFirst                                                              = 0x00040140
	ViewSLast                                                               = 0x0004014F
	ViewEDraw                                                               = 0x80040140
	RegdbEFirst                                                             = 0x80040150
	RegdbELast                                                              = 0x8004015F
	RegdbSFirst                                                             = 0x00040150
	RegdbSLast                                                              = 0x0004015F
	RegdbEReadregdb                                                         = 0x80040150
	RegdbEWriteregdb                                                        = 0x80040151
	RegdbEKeymissing                                                        = 0x80040152
	RegdbEInvalidvalue                                                      = 0x80040153
	RegdbEClassnotreg                                                       = 0x80040154
	RegdbEIidnotreg                                                         = 0x80040155
	RegdbEBadthreadingmodel                                                 = 0x80040156
	CatEFirst                                                               = 0x80040160
	CatELast                                                                = 0x80040161
	CatECatidnoexist                                                        = 0x80040160
	CatENodescription                                                       = 0x80040161
	CsEFirst                                                                = 0x80040164
	CsELast                                                                 = 0x8004016F
	CsEPackageNotfound                                                      = 0x80040164
	CsENotDeletable                                                         = 0x80040165
	CsEClassNotfound                                                        = 0x80040166
	CsEInvalidVersion                                                       = 0x80040167
	CsENoClassstore                                                         = 0x80040168
	CsEObjectNotfound                                                       = 0x80040169
	CsEObjectAlreadyExists                                                  = 0x8004016A
	CsEInvalidPath                                                          = 0x8004016B
	CsENetworkError                                                         = 0x8004016C
	CsEAdminLimitExceeded                                                   = 0x8004016D
	CsESchemaMismatch                                                       = 0x8004016E
	CsEInternalError                                                        = 0x8004016F
	CacheEFirst                                                             = 0x80040170
	CacheELast                                                              = 0x8004017F
	CacheSFirst                                                             = 0x00040170
	CacheSLast                                                              = 0x0004017F
	CacheENocacheUpdated                                                    = 0x80040170
	OleobjEFirst                                                            = 0x80040180
	OleobjELast                                                             = 0x8004018F
	OleobjSFirst                                                            = 0x00040180
	OleobjSLast                                                             = 0x0004018F
	OleobjENoverbs                                                          = 0x80040180
	OleobjEInvalidverb                                                      = 0x80040181
	ClientsiteEFirst                                                        = 0x80040190
	ClientsiteELast                                                         = 0x8004019F
	ClientsiteSFirst                                                        = 0x00040190
	ClientsiteSLast                                                         = 0x0004019F
	InplaceENotundoable                                                     = 0x800401A0
	InplaceENotoolspace                                                     = 0x800401A1
	InplaceEFirst                                                           = 0x800401A0
	InplaceELast                                                            = 0x800401AF
	InplaceSFirst                                                           = 0x000401A0
	InplaceSLast                                                            = 0x000401AF
	EnumEFirst                                                              = 0x800401B0
	EnumELast                                                               = 0x800401BF
	EnumSFirst                                                              = 0x000401B0
	EnumSLast                                                               = 0x000401BF
	Convert10EFirst                                                         = 0x800401C0
	Convert10ELast                                                          = 0x800401CF
	Convert10SFirst                                                         = 0x000401C0
	Convert10SLast                                                          = 0x000401CF
	Convert10EOlestreamGet                                                  = 0x800401C0
	Convert10EOlestreamPut                                                  = 0x800401C1
	Convert10EOlestreamFmt                                                  = 0x800401C2
	Convert10EOlestreamBitmapToDib                                          = 0x800401C3
	Convert10EStgFmt                                                        = 0x800401C4
	Convert10EStgNoStdStream                                                = 0x800401C5
	Convert10EStgDibToBitmap                                                = 0x800401C6
	ClipbrdEFirst                                                           = 0x800401D0
	ClipbrdELast                                                            = 0x800401DF
	ClipbrdSFirst                                                           = 0x000401D0
	ClipbrdSLast                                                            = 0x000401DF
	ClipbrdECantOpen                                                        = 0x800401D0
	ClipbrdECantEmpty                                                       = 0x800401D1
	ClipbrdECantSet                                                         = 0x800401D2
	ClipbrdEBadData                                                         = 0x800401D3
	ClipbrdECantClose                                                       = 0x800401D4
	MkEFirst                                                                = 0x800401E0
	MkELast                                                                 = 0x800401EF
	MkSFirst                                                                = 0x000401E0
	MkSLast                                                                 = 0x000401EF
	MkEConnectmanually                                                      = 0x800401E0
	MkEExceededdeadline                                                     = 0x800401E1
	MkENeedgeneric                                                          = 0x800401E2
	MkEUnavailable                                                          = 0x800401E3
	MkESyntax                                                               = 0x800401E4
	MkENoobject                                                             = 0x800401E5
	MkEInvalidextension                                                     = 0x800401E6
	MkEIntermediateinterfacenotsupported                                    = 0x800401E7
	MkENotbindable                                                          = 0x800401E8
	MkENotbound                                                             = 0x800401E9
	MkECantopenfile                                                         = 0x800401EA
	MkEMustbotheruser                                                       = 0x800401EB
	MkENoinverse                                                            = 0x800401EC
	MkENostorage                                                            = 0x800401ED
	MkENoprefix                                                             = 0x800401EE
	MkEEnumerationFailed                                                    = 0x800401EF
	CoEFirst                                                                = 0x800401F0
	CoELast                                                                 = 0x800401FF
	CoSFirst                                                                = 0x000401F0
	CoSLast                                                                 = 0x000401FF
	CoENotinitialized                                                       = 0x800401F0
	CoEAlreadyinitialized                                                   = 0x800401F1
	CoECantdetermineclass                                                   = 0x800401F2
	CoEClassstring                                                          = 0x800401F3
	CoEIidstring                                                            = 0x800401F4
	CoEAppnotfound                                                          = 0x800401F5
	CoEAppsingleuse                                                         = 0x800401F6
	CoEErrorinapp                                                           = 0x800401F7
	CoEDllnotfound                                                          = 0x800401F8
	CoEErrorindll                                                           = 0x800401F9
	CoEWrongosforapp                                                        = 0x800401FA
	CoEObjnotreg                                                            = 0x800401FB
	CoEObjisreg                                                             = 0x800401FC
	CoEObjnotconnected                                                      = 0x800401FD
	CoEAppdidntreg                                                          = 0x800401FE
	CoEReleased                                                             = 0x800401FF
	EventEFirst                                                             = 0x80040200
	EventELast                                                              = 0x8004021F
	EventSFirst                                                             = 0x00040200
	EventSLast                                                              = 0x0004021F
	EventSSomeSubscribersFailed                                             = 0x00040200
	EventEAllSubscribersFailed                                              = 0x80040201
	EventSNosubscribers                                                     = 0x00040202
	EventEQuerysyntax                                                       = 0x80040203
	EventEQueryfield                                                        = 0x80040204
	EventEInternalexception                                                 = 0x80040205
	EventEInternalerror                                                     = 0x80040206
	EventEInvalidPerUserSid                                                 = 0x80040207
	EventEUserException                                                     = 0x80040208
	EventETooManyMethods                                                    = 0x80040209
	EventEMissingEventclass                                                 = 0x8004020A
	EventENotAllRemoved                                                     = 0x8004020B
	EventEComplusNotInstalled                                               = 0x8004020C
	EventECantModifyOrDeleteUnconfiguredObject                              = 0x8004020D
	EventECantModifyOrDeleteConfiguredObject                                = 0x8004020E
	EventEInvalidEventClassPartition                                        = 0x8004020F
	EventEPerUserSidNotLoggedOn                                             = 0x80040210
	XactEFirst                                                              = 0x8004D000
	XactELast                                                               = 0x8004D02B
	XactSFirst                                                              = 0x0004D000
	XactSLast                                                               = 0x0004D010
	XactEAlreadyothersinglephase                                            = 0x8004D000
	XactECantretain                                                         = 0x8004D001
	XactECommitfailed                                                       = 0x8004D002
	XactECommitprevented                                                    = 0x8004D003
	XactEHeuristicabort                                                     = 0x8004D004
	XactEHeuristiccommit                                                    = 0x8004D005
	XactEHeuristicdamage                                                    = 0x8004D006
	XactEHeuristicdanger                                                    = 0x8004D007
	XactEIsolationlevel                                                     = 0x8004D008
	XactENoasync                                                            = 0x8004D009
	XactENoenlist                                                           = 0x8004D00A
	XactENoisoretain                                                        = 0x8004D00B
	XactENoresource                                                         = 0x8004D00C
	XactENotcurrent                                                         = 0x8004D00D
	XactENotransaction                                                      = 0x8004D00E
	XactENotsupported                                                       = 0x8004D00F
	XactEUnknownrmgrid                                                      = 0x8004D010
	XactEWrongstate                                                         = 0x8004D011
	XactEWronguow                                                           = 0x8004D012
	XactEXtionexists                                                        = 0x8004D013
	XactENoimportobject                                                     = 0x8004D014
	XactEInvalidcookie                                                      = 0x8004D015
	XactEIndoubt                                                            = 0x8004D016
	XactENotimeout                                                          = 0x8004D017
	XactEAlreadyinprogress                                                  = 0x8004D018
	XactEAborted                                                            = 0x8004D019
	XactELogfull                                                            = 0x8004D01A
	XactETmnotavailable                                                     = 0x8004D01B
	XactEConnectionDown                                                     = 0x8004D01C
	XactEConnectionDenied                                                   = 0x8004D01D
	XactEReenlisttimeout                                                    = 0x8004D01E
	XactETipConnectFailed                                                   = 0x8004D01F
	XactETipProtocolError                                                   = 0x8004D020
	XactETipPullFailed                                                      = 0x8004D021
	XactEDestTmnotavailable                                                 = 0x8004D022
	XactETipDisabled                                                        = 0x8004D023
	XactENetworkTxDisabled                                                  = 0x8004D024
	XactEPartnerNetworkTxDisabled                                           = 0x8004D025
	XactEXaTxDisabled                                                       = 0x8004D026
	XactEUnableToReadDtcConfig                                              = 0x8004D027
	XactEUnableToLoadDtcProxy                                               = 0x8004D028
	XactEAborting                                                           = 0x8004D029
	XactEPushCommFailure                                                    = 0x8004D02A
	XactEPullCommFailure                                                    = 0x8004D02B
	XactELuTxDisabled                                                       = 0x8004D02C
	XactEClerknotfound                                                      = 0x8004D080
	XactEClerkexists                                                        = 0x8004D081
	XactERecoveryinprogress                                                 = 0x8004D082
	XactETransactionclosed                                                  = 0x8004D083
	XactEInvalidlsn                                                         = 0x8004D084
	XactEReplayrequest                                                      = 0x8004D085
	XactSAsync                                                              = 0x0004D000
	XactSDefect                                                             = 0x0004D001
	XactSReadonly                                                           = 0x0004D002
	XactSSomenoretain                                                       = 0x0004D003
	XactSOkinform                                                           = 0x0004D004
	XactSMadechangescontent                                                 = 0x0004D005
	XactSMadechangesinform                                                  = 0x0004D006
	XactSAllnoretain                                                        = 0x0004D007
	XactSAborting                                                           = 0x0004D008
	XactSSinglephase                                                        = 0x0004D009
	XactSLocallyOk                                                          = 0x0004D00A
	XactSLastresourcemanager                                                = 0x0004D010
	ContextEFirst                                                           = 0x8004E000
	ContextELast                                                            = 0x8004E02F
	ContextSFirst                                                           = 0x0004E000
	ContextSLast                                                            = 0x0004E02F
	ContextEAborted                                                         = 0x8004E002
	ContextEAborting                                                        = 0x8004E003
	ContextENocontext                                                       = 0x8004E004
	ContextEWouldDeadlock                                                   = 0x8004E005
	ContextESynchTimeout                                                    = 0x8004E006
	ContextEOldref                                                          = 0x8004E007
	ContextERolenotfound                                                    = 0x8004E00C
	ContextETmnotavailable                                                  = 0x8004E00F
	CoEActivationfailed                                                     = 0x8004E021
	CoEActivationfailedEventlogged                                          = 0x8004E022
	CoEActivationfailedCatalogerror                                         = 0x8004E023
	CoEActivationfailedTimeout                                              = 0x8004E024
	CoEInitializationfailed                                                 = 0x8004E025
	ContextENojit                                                           = 0x8004E026
	ContextENotransaction                                                   = 0x8004E027
	CoEThreadingmodelChanged                                                = 0x8004E028
	CoENoiisintrinsics                                                      = 0x8004E029
	CoENocookies                                                            = 0x8004E02A
	CoEDberror                                                              = 0x8004E02B
	CoENotpooled                                                            = 0x8004E02C
	CoENotconstructed                                                       = 0x8004E02D
	CoENosynchronization                                                    = 0x8004E02E
	CoEIsolevelmismatch                                                     = 0x8004E02F
	CoECallOutOfTxScopeNotAllowed                                           = 0x8004E030
	CoEExitTransactionScopeNotCalled                                        = 0x8004E031
	OleSUsereg                                                              = 0x00040000
	OleSStatic                                                              = 0x00040001
	OleSMacClipformat                                                       = 0x00040002
	DragdropSDrop                                                           = 0x00040100
	DragdropSCancel                                                         = 0x00040101
	DragdropSUsedefaultcursors                                              = 0x00040102
	DataSSameformatetc                                                      = 0x00040130
	ViewSAlreadyFrozen                                                      = 0x00040140
	CacheSFormatetcNotsupported                                             = 0x00040170
	CacheSSamecache                                                         = 0x00040171
	CacheSSomecachesNotupdated                                              = 0x00040172
	OleobjSInvalidverb                                                      = 0x00040180
	OleobjSCannotDoverbNow                                                  = 0x00040181
	OleobjSInvalidhwnd                                                      = 0x00040182
	InplaceSTruncated                                                       = 0x000401A0
	Convert10SNoPresentation                                                = 0x000401C0
	MkSReducedToSelf                                                        = 0x000401E2
	MkSMe                                                                   = 0x000401E4
	MkSHim                                                                  = 0x000401E5
	MkSUs                                                                   = 0x000401E6
	MkSMonikeralreadyregistered                                             = 0x000401E7
	SchedSTaskReady                                                         = 0x00041300
	SchedSTaskRunning                                                       = 0x00041301
	SchedSTaskDisabled                                                      = 0x00041302
	SchedSTaskHasNotRun                                                     = 0x00041303
	SchedSTaskNoMoreRuns                                                    = 0x00041304
	SchedSTaskNotScheduled                                                  = 0x00041305
	SchedSTaskTerminated                                                    = 0x00041306
	SchedSTaskNoValidTriggers                                               = 0x00041307
	SchedSEventTrigger                                                      = 0x00041308
	SchedETriggerNotFound                                                   = 0x80041309
	SchedETaskNotReady                                                      = 0x8004130A
	SchedETaskNotRunning                                                    = 0x8004130B
	SchedEServiceNotInstalled                                               = 0x8004130C
	SchedECannotOpenTask                                                    = 0x8004130D
	SchedEInvalidTask                                                       = 0x8004130E
	SchedEAccountInformationNotSet                                          = 0x8004130F
	SchedEAccountNameNotFound                                               = 0x80041310
	SchedEAccountDbaseCorrupt                                               = 0x80041311
	SchedENoSecurityServices                                                = 0x80041312
	SchedEUnknownObjectVersion                                              = 0x80041313
	SchedEUnsupportedAccountOption                                          = 0x80041314
	SchedEServiceNotRunning                                                 = 0x80041315
	SchedEUnexpectednode                                                    = 0x80041316
	SchedENamespace                                                         = 0x80041317
	SchedEInvalidvalue                                                      = 0x80041318
	SchedEMissingnode                                                       = 0x80041319
	SchedEMalformedxml                                                      = 0x8004131A
	SchedSSomeTriggersFailed                                                = 0x0004131B
	SchedSBatchLogonProblem                                                 = 0x0004131C
	SchedETooManyNodes                                                      = 0x8004131D
	SchedEPastEndBoundary                                                   = 0x8004131E
	SchedEAlreadyRunning                                                    = 0x8004131F
	SchedEUserNotLoggedOn                                                   = 0x80041320
	SchedEInvalidTaskHash                                                   = 0x80041321
	SchedEServiceNotAvailable                                               = 0x80041322
	SchedEServiceTooBusy                                                    = 0x80041323
	SchedETaskAttempted                                                     = 0x80041324
	SchedSTaskQueued                                                        = 0x00041325
	SchedETaskDisabled                                                      = 0x80041326
	SchedETaskNotV1Compat                                                   = 0x80041327
	SchedEStartOnDemand                                                     = 0x80041328
	CoEClassCreateFailed                                                    = 0x80080001
	CoEScmError                                                             = 0x80080002
	CoEScmRpcFailure                                                        = 0x80080003
	CoEBadPath                                                              = 0x80080004
	CoEServerExecFailure                                                    = 0x80080005
	CoEObjsrvRpcFailure                                                     = 0x80080006
	MkENoNormalized                                                         = 0x80080007
	CoEServerStopping                                                       = 0x80080008
	MemEInvalidRoot                                                         = 0x80080009
	MemEInvalidLink                                                         = 0x80080010
	MemEInvalidSize                                                         = 0x80080011
	CoSNotallinterfaces                                                     = 0x00080012
	CoSMachinenamenotfound                                                  = 0x00080013
	CoEMissingDisplayname                                                   = 0x80080015
	CoERunasValueMustBeAaa                                                  = 0x80080016
	CoEElevationDisabled                                                    = 0x80080017
	DispEUnknowninterface                                                   = 0x80020001
	DispEMembernotfound                                                     = 0x80020003
	DispEParamnotfound                                                      = 0x80020004
	DispETypemismatch                                                       = 0x80020005
	DispEUnknownname                                                        = 0x80020006
	DispENonamedargs                                                        = 0x80020007
	DispEBadvartype                                                         = 0x80020008
	DispEException                                                          = 0x80020009
	DispEOverflow                                                           = 0x8002000A
	DispEBadindex                                                           = 0x8002000B
	DispEUnknownlcid                                                        = 0x8002000C
	DispEArrayislocked                                                      = 0x8002000D
	DispEBadparamcount                                                      = 0x8002000E
	DispEParamnotoptional                                                   = 0x8002000F
	DispEBadcallee                                                          = 0x80020010
	DispENotacollection                                                     = 0x80020011
	DispEDivbyzero                                                          = 0x80020012
	DispEBuffertoosmall                                                     = 0x80020013
	TypeEBuffertoosmall                                                     = 0x80028016
	TypeEFieldnotfound                                                      = 0x80028017
	TypeEInvdataread                                                        = 0x80028018
	TypeEUnsupformat                                                        = 0x80028019
	TypeERegistryaccess                                                     = 0x8002801C
	TypeELibnotregistered                                                   = 0x8002801D
	TypeEUndefinedtype                                                      = 0x80028027
	TypeEQualifiednamedisallowed                                            = 0x80028028
	TypeEInvalidstate                                                       = 0x80028029
	TypeEWrongtypekind                                                      = 0x8002802A
	TypeEElementnotfound                                                    = 0x8002802B
	TypeEAmbiguousname                                                      = 0x8002802C
	TypeENameconflict                                                       = 0x8002802D
	TypeEUnknownlcid                                                        = 0x8002802E
	TypeEDllfunctionnotfound                                                = 0x8002802F
	TypeEBadmodulekind                                                      = 0x800288BD
	TypeESizetoobig                                                         = 0x800288C5
	TypeEDuplicateid                                                        = 0x800288C6
	TypeEInvalidid                                                          = 0x800288CF
	TypeETypemismatch                                                       = 0x80028CA0
	TypeEOutofbounds                                                        = 0x80028CA1
	TypeEIoerror                                                            = 0x80028CA2
	TypeECantcreatetmpfile                                                  = 0x80028CA3
	TypeECantloadlibrary                                                    = 0x80029C4A
	TypeEInconsistentpropfuncs                                              = 0x80029C83
	TypeECirculartype                                                       = 0x80029C84
	StgEInvalidfunction                                                     = 0x80030001
	StgEFilenotfound                                                        = 0x80030002
	StgEPathnotfound                                                        = 0x80030003
	StgEToomanyopenfiles                                                    = 0x80030004
	StgEAccessdenied                                                        = 0x80030005
	StgEInvalidhandle                                                       = 0x80030006
	StgEInsufficientmemory                                                  = 0x80030008
	StgEInvalidpointer                                                      = 0x80030009
	StgENomorefiles                                                         = 0x80030012
	StgEDiskiswriteprotected                                                = 0x80030013
	StgESeekerror                                                           = 0x80030019
	StgEWritefault                                                          = 0x8003001D
	StgEReadfault                                                           = 0x8003001E
	StgEShareviolation                                                      = 0x80030020
	StgELockviolation                                                       = 0x80030021
	StgEFilealreadyexists                                                   = 0x80030050
	StgEInvalidparameter                                                    = 0x80030057
	StgEMediumfull                                                          = 0x80030070
	StgEPropsetmismatched                                                   = 0x800300F0
	StgEAbnormalapiexit                                                     = 0x800300FA
	StgEInvalidheader                                                       = 0x800300FB
	StgEInvalidname                                                         = 0x800300FC
	StgEUnknown                                                             = 0x800300FD
	StgEUnimplementedfunction                                               = 0x800300FE
	StgEInvalidflag                                                         = 0x800300FF
	StgEInuse                                                               = 0x80030100
	StgENotcurrent                                                          = 0x80030101
	StgEReverted                                                            = 0x80030102
	StgECantsave                                                            = 0x80030103
	StgEOldformat                                                           = 0x80030104
	StgEOlddll                                                              = 0x80030105
	StgESharerequired                                                       = 0x80030106
	StgENotfilebasedstorage                                                 = 0x80030107
	StgEExtantmarshallings                                                  = 0x80030108
	StgEDocfilecorrupt                                                      = 0x80030109
	StgEBadbaseaddress                                                      = 0x80030110
	StgEDocfiletoolarge                                                     = 0x80030111
	StgENotsimpleformat                                                     = 0x80030112
	StgEIncomplete                                                          = 0x80030201
	StgETerminated                                                          = 0x80030202
	StgSConverted                                                           = 0x00030200
	StgSBlock                                                               = 0x00030201
	StgSRetrynow                                                            = 0x00030202
	StgSMonitoring                                                          = 0x00030203
	StgSMultipleopens                                                       = 0x00030204
	StgSConsolidationfailed                                                 = 0x00030205
	StgSCannotconsolidate                                                   = 0x00030206
	StgEStatusCopyProtectionFailure                                         = 0x80030305
	StgECssAuthenticationFailure                                            = 0x80030306
	StgECssKeyNotPresent                                                    = 0x80030307
	StgECssKeyNotEstablished                                                = 0x80030308
	StgECssScrambledSector                                                  = 0x80030309
	StgECssRegionMismatch                                                   = 0x8003030A
	StgEResetsExhausted                                                     = 0x8003030B
	RpcECallRejected                                                        = 0x80010001
	RpcECallCanceled                                                        = 0x80010002
	RpcECantpostInsendcall                                                  = 0x80010003
	RpcECantcalloutInasynccall                                              = 0x80010004
	RpcECantcalloutInexternalcall                                           = 0x80010005
	RpcEConnectionTerminated                                                = 0x80010006
	RpcEServerDied                                                          = 0x80010007
	RpcEClientDied                                                          = 0x80010008
	RpcEInvalidDatapacket                                                   = 0x80010009
	RpcECanttransmitCall                                                    = 0x8001000A
	RpcEClientCantmarshalData                                               = 0x8001000B
	RpcEClientCantunmarshalData                                             = 0x8001000C
	RpcEServerCantmarshalData                                               = 0x8001000D
	RpcEServerCantunmarshalData                                             = 0x8001000E
	RpcEInvalidData                                                         = 0x8001000F
	RpcEInvalidParameter                                                    = 0x80010010
	RpcECantcalloutAgain                                                    = 0x80010011
	RpcEServerDiedDne                                                       = 0x80010012
	RpcESysCallFailed                                                       = 0x80010100
	RpcEOutOfResources                                                      = 0x80010101
	RpcEAttemptedMultithread                                                = 0x80010102
	RpcENotRegistered                                                       = 0x80010103
	RpcEFault                                                               = 0x80010104
	RpcEServerfault                                                         = 0x80010105
	RpcEChangedMode                                                         = 0x80010106
	RpcEInvalidmethod                                                       = 0x80010107
	RpcEDisconnected                                                        = 0x80010108
	RpcERetry                                                               = 0x80010109
	RpcEServercallRetrylater                                                = 0x8001010A
	RpcEServercallRejected                                                  = 0x8001010B
	RpcEInvalidCalldata                                                     = 0x8001010C
	RpcECantcalloutIninputsynccall                                          = 0x8001010D
	RpcEWrongThread                                                         = 0x8001010E
	RpcEThreadNotInit                                                       = 0x8001010F
	RpcEVersionMismatch                                                     = 0x80010110
	RpcEInvalidHeader                                                       = 0x80010111
	RpcEInvalidExtension                                                    = 0x80010112
	RpcEInvalidIpid                                                         = 0x80010113
	RpcEInvalidObject                                                       = 0x80010114
	RpcSCallpending                                                         = 0x80010115
	RpcSWaitontimer                                                         = 0x80010116
	RpcECallComplete                                                        = 0x80010117
	RpcEUnsecureCall                                                        = 0x80010118
	RpcETooLate                                                             = 0x80010119
	RpcENoGoodSecurityPackages                                              = 0x8001011A
	RpcEAccessDenied                                                        = 0x8001011B
	RpcERemoteDisabled                                                      = 0x8001011C
	RpcEInvalidObjref                                                       = 0x8001011D
	RpcENoContext                                                           = 0x8001011E
	RpcETimeout                                                             = 0x8001011F
	RpcENoSync                                                              = 0x80010120
	RpcEFullsicRequired                                                     = 0x80010121
	RpcEInvalidStdName                                                      = 0x80010122
	CoEFailedtoimpersonate                                                  = 0x80010123
	CoEFailedtogetsecctx                                                    = 0x80010124
	CoEFailedtoopenthreadtoken                                              = 0x80010125
	CoEFailedtogettokeninfo                                                 = 0x80010126
	CoETrusteedoesntmatchclient                                             = 0x80010127
	CoEFailedtoqueryclientblanket                                           = 0x80010128
	CoEFailedtosetdacl                                                      = 0x80010129
	CoEAccesscheckfailed                                                    = 0x8001012A
	CoENetaccessapifailed                                                   = 0x8001012B
	CoEWrongtrusteenamesyntax                                               = 0x8001012C
	CoEInvalidsid                                                           = 0x8001012D
	CoEConversionfailed                                                     = 0x8001012E
	CoENomatchingsidfound                                                   = 0x8001012F
	CoELookupaccsidfailed                                                   = 0x80010130
	CoENomatchingnamefound                                                  = 0x80010131
	CoELookupaccnamefailed                                                  = 0x80010132
	CoESetserlhndlfailed                                                    = 0x80010133
	CoEFailedtogetwindir                                                    = 0x80010134
	CoEPathtoolong                                                          = 0x80010135
	CoEFailedtogenuuid                                                      = 0x80010136
	CoEFailedtocreatefile                                                   = 0x80010137
	CoEFailedtoclosehandle                                                  = 0x80010138
	CoEExceedsysacllimit                                                    = 0x80010139
	CoEAcesinwrongorder                                                     = 0x8001013A
	CoEIncompatiblestreamversion                                            = 0x8001013B
	CoEFailedtoopenprocesstoken                                             = 0x8001013C
	CoEDecodefailed                                                         = 0x8001013D
	CoEAcnotinitialized                                                     = 0x8001013F
	CoECancelDisabled                                                       = 0x80010140
	RpcEUnexpected                                                          = 0x8001FFFF
	ErrorAuditingDisabled                                                   = 0xC0090001
	ErrorAllSidsFiltered                                                    = 0xC0090002
	ErrorBizrulesNotEnabled                                                 = 0xC0090003
	NteBadUid                                                               = 0x80090001
	NteBadHash                                                              = 0x80090002
	NteBadKey                                                               = 0x80090003
	NteBadLen                                                               = 0x80090004
	NteBadData                                                              = 0x80090005
	NteBadSignature                                                         = 0x80090006
	NteBadVer                                                               = 0x80090007
	NteBadAlgid                                                             = 0x80090008
	NteBadFlags                                                             = 0x80090009
	NteBadType                                                              = 0x8009000A
	NteBadKeyState                                                          = 0x8009000B
	NteBadHashState                                                         = 0x8009000C
	NteNoKey                                                                = 0x8009000D
	NteNoMemory                                                             = 0x8009000E
	NteExists                                                               = 0x8009000F
	NtePerm                                                                 = 0x80090010
	NteNotFound                                                             = 0x80090011
	NteDoubleEncrypt                                                        = 0x80090012
	NteBadProvider                                                          = 0x80090013
	NteBadProvType                                                          = 0x80090014
	NteBadPublicKey                                                         = 0x80090015
	NteBadKeyset                                                            = 0x80090016
	NteProvTypeNotDef                                                       = 0x80090017
	NteProvTypeEntryBad                                                     = 0x80090018
	NteKeysetNotDef                                                         = 0x80090019
	NteKeysetEntryBad                                                       = 0x8009001A
	NteProvTypeNoMatch                                                      = 0x8009001B
	NteSignatureFileBad                                                     = 0x8009001C
	NteProviderDllFail                                                      = 0x8009001D
	NteProvDllNotFound                                                      = 0x8009001E
	NteBadKeysetParam                                                       = 0x8009001F
	NteFail                                                                 = 0x80090020
	NteSysErr                                                               = 0x80090021
	NteSilentContext                                                        = 0x80090022
	NteTokenKeysetStorageFull                                               = 0x80090023
	NteTemporaryProfile                                                     = 0x80090024
	NteFixedparameter                                                       = 0x80090025
	NteInvalidHandle                                                        = 0x80090026
	NteInvalidParameter                                                     = 0x80090027
	NteBufferTooSmall                                                       = 0x80090028
	NteNotSupported                                                         = 0x80090029
	NteNoMoreItems                                                          = 0x8009002A
	NteBuffersOverlap                                                       = 0x8009002B
	NteDecryptionFailure                                                    = 0x8009002C
	NteInternalError                                                        = 0x8009002D
	NteUiRequired                                                           = 0x8009002E
	NteHmacNotSupported                                                     = 0x8009002F
	SecEInsufficientMemory                                                  = 0x80090300
	SecEInvalidHandle                                                       = 0x80090301
	SecEUnsupportedFunction                                                 = 0x80090302
	SecETargetUnknown                                                       = 0x80090303
	SecEInternalError                                                       = 0x80090304
	SecESecpkgNotFound                                                      = 0x80090305
	SecENotOwner                                                            = 0x80090306
	SecECannotInstall                                                       = 0x80090307
	SecEInvalidToken                                                        = 0x80090308
	SecECannotPack                                                          = 0x80090309
	SecEQopNotSupported                                                     = 0x8009030A
	SecENoImpersonation                                                     = 0x8009030B
	SecELogonDenied                                                         = 0x8009030C
	SecEUnknownCredentials                                                  = 0x8009030D
	SecENoCredentials                                                       = 0x8009030E
	SecEMessageAltered                                                      = 0x8009030F
	SecEOutOfSequence                                                       = 0x80090310
	SecENoAuthenticatingAuthority                                           = 0x80090311
	SecIContinueNeeded                                                      = 0x00090312
	SecICompleteNeeded                                                      = 0x00090313
	SecICompleteAndContinue                                                 = 0x00090314
	SecILocalLogon                                                          = 0x00090315
	SecEBadPkgid                                                            = 0x80090316
	SecEContextExpired                                                      = 0x80090317
	SecIContextExpired                                                      = 0x00090317
	SecEIncompleteMessage                                                   = 0x80090318
	SecEIncompleteCredentials                                               = 0x80090320
	SecEBufferTooSmall                                                      = 0x80090321
	SecIIncompleteCredentials                                               = 0x00090320
	SecIRenegotiate                                                         = 0x00090321
	SecEWrongPrincipal                                                      = 0x80090322
	SecINoLsaContext                                                        = 0x00090323
	SecETimeSkew                                                            = 0x80090324
	SecEUntrustedRoot                                                       = 0x80090325
	SecEIllegalMessage                                                      = 0x80090326
	SecECertUnknown                                                         = 0x80090327
	SecECertExpired                                                         = 0x80090328
	SecEEncryptFailure                                                      = 0x80090329
	SecEDecryptFailure                                                      = 0x80090330
	SecEAlgorithmMismatch                                                   = 0x80090331
	SecESecurityQosFailed                                                   = 0x80090332
	SecEUnfinishedContextDeleted                                            = 0x80090333
	SecENoTgtReply                                                          = 0x80090334
	SecENoIpAddresses                                                       = 0x80090335
	SecEWrongCredentialHandle                                               = 0x80090336
	SecECryptoSystemInvalid                                                 = 0x80090337
	SecEMaxReferralsExceeded                                                = 0x80090338
	SecEMustBeKdc                                                           = 0x80090339
	SecEStrongCryptoNotSupported                                            = 0x8009033A
	SecETooManyPrincipals                                                   = 0x8009033B
	SecENoPaData                                                            = 0x8009033C
	SecEPkinitNameMismatch                                                  = 0x8009033D
	SecESmartcardLogonRequired                                              = 0x8009033E
	SecEShutdownInProgress                                                  = 0x8009033F
	SecEKdcInvalidRequest                                                   = 0x80090340
	SecEKdcUnableToRefer                                                    = 0x80090341
	SecEKdcUnknownEtype                                                     = 0x80090342
	SecEUnsupportedPreauth                                                  = 0x80090343
	SecEDelegationRequired                                                  = 0x80090345
	SecEBadBindings                                                         = 0x80090346
	SecEMultipleAccounts                                                    = 0x80090347
	SecENoKerbKey                                                           = 0x80090348
	SecECertWrongUsage                                                      = 0x80090349
	SecEDowngradeDetected                                                   = 0x80090350
	SecESmartcardCertRevoked                                                = 0x80090351
	SecEIssuingCaUntrusted                                                  = 0x80090352
	SecERevocationOfflineC                                                  = 0x80090353
	SecEPkinitClientFailure                                                 = 0x80090354
	SecESmartcardCertExpired                                                = 0x80090355
	SecENoS4UProtSupport                                                    = 0x80090356
	SecECrossrealmDelegationFailure                                         = 0x80090357
	SecERevocationOfflineKdc                                                = 0x80090358
	SecEIssuingCaUntrustedKdc                                               = 0x80090359
	SecEKdcCertExpired                                                      = 0x8009035A
	SecEKdcCertRevoked                                                      = 0x8009035B
	SecISignatureNeeded                                                     = 0x0009035C
	SecEInvalidParameter                                                    = 0x8009035D
	SecEDelegationPolicy                                                    = 0x8009035E
	SecEPolicyNltmOnly                                                      = 0x8009035F
	SecINoRenegotiation                                                     = 0x00090360
	SecENoContext                                                           = 0x80090361
	SecEPku2UCertFailure                                                    = 0x80090362
	SecEMutualAuthFailed                                                    = 0x80090363
	CryptEMsgError                                                          = 0x80091001
	CryptEUnknownAlgo                                                       = 0x80091002
	CryptEOidFormat                                                         = 0x80091003
	CryptEInvalidMsgType                                                    = 0x80091004
	CryptEUnexpectedEncoding                                                = 0x80091005
	CryptEAuthAttrMissing                                                   = 0x80091006
	CryptEHashValue                                                         = 0x80091007
	CryptEInvalidIndex                                                      = 0x80091008
	CryptEAlreadyDecrypted                                                  = 0x80091009
	CryptENotDecrypted                                                      = 0x8009100A
	CryptERecipientNotFound                                                 = 0x8009100B
	CryptEControlType                                                       = 0x8009100C
	CryptEIssuerSerialnumber                                                = 0x8009100D
	CryptESignerNotFound                                                    = 0x8009100E
	CryptEAttributesMissing                                                 = 0x8009100F
	CryptEStreamMsgNotReady                                                 = 0x80091010
	CryptEStreamInsufficientData                                            = 0x80091011
	CryptINewProtectionRequired                                             = 0x00091012
	CryptEBadLen                                                            = 0x80092001
	CryptEBadEncode                                                         = 0x80092002
	CryptEFileError                                                         = 0x80092003
	CryptENotFound                                                          = 0x80092004
	CryptEExists                                                            = 0x80092005
	CryptENoProvider                                                        = 0x80092006
	CryptESelfSigned                                                        = 0x80092007
	CryptEDeletedPrev                                                       = 0x80092008
	CryptENoMatch                                                           = 0x80092009
	CryptEUnexpectedMsgType                                                 = 0x8009200A
	CryptENoKeyProperty                                                     = 0x8009200B
	CryptENoDecryptCert                                                     = 0x8009200C
	CryptEBadMsg                                                            = 0x8009200D
	CryptENoSigner                                                          = 0x8009200E
	CryptEPendingClose                                                      = 0x8009200F
	CryptERevoked                                                           = 0x80092010
	CryptENoRevocationDll                                                   = 0x80092011
	CryptENoRevocationCheck                                                 = 0x80092012
	CryptERevocationOffline                                                 = 0x80092013
	CryptENotInRevocationDatabase                                           = 0x80092014
	CryptEInvalidNumericString                                              = 0x80092020
	CryptEInvalidPrintableString                                            = 0x80092021
	CryptEInvalidIa5String                                                  = 0x80092022
	CryptEInvalidX500String                                                 = 0x80092023
	CryptENotCharString                                                     = 0x80092024
	CryptEFileresized                                                       = 0x80092025
	CryptESecuritySettings                                                  = 0x80092026
	CryptENoVerifyUsageDll                                                  = 0x80092027
	CryptENoVerifyUsageCheck                                                = 0x80092028
	CryptEVerifyUsageOffline                                                = 0x80092029
	CryptENotInCtl                                                          = 0x8009202A
	CryptENoTrustedSigner                                                   = 0x8009202B
	CryptEMissingPubkeyPara                                                 = 0x8009202C
	CryptEOssError                                                          = 0x80093000
	OssMoreBuf                                                              = 0x80093001
	OssNegativeUinteger                                                     = 0x80093002
	OssPduRange                                                             = 0x80093003
	OssMoreInput                                                            = 0x80093004
	OssDataError                                                            = 0x80093005
	OssBadArg                                                               = 0x80093006
	OssBadVersion                                                           = 0x80093007
	OssOutMemory                                                            = 0x80093008
	OssPduMismatch                                                          = 0x80093009
	OssLimited                                                              = 0x8009300A
	OssBadPtr                                                               = 0x8009300B
	OssBadTime                                                              = 0x8009300C
	OssIndefiniteNotSupported                                               = 0x8009300D
	OssMemError                                                             = 0x8009300E
	OssBadTable                                                             = 0x8009300F
	OssTooLong                                                              = 0x80093010
	OssConstraintViolated                                                   = 0x80093011
	OssFatalError                                                           = 0x80093012
	OssAccessSerializationError                                             = 0x80093013
	OssNullTbl                                                              = 0x80093014
	OssNullFcn                                                              = 0x80093015
	OssBadEncrules                                                          = 0x80093016
	OssUnavailEncrules                                                      = 0x80093017
	OssCantOpenTraceWindow                                                  = 0x80093018
	OssUnimplemented                                                        = 0x80093019
	OssOidDllNotLinked                                                      = 0x8009301A
	OssCantOpenTraceFile                                                    = 0x8009301B
	OssTraceFileAlreadyOpen                                                 = 0x8009301C
	OssTableMismatch                                                        = 0x8009301D
	OssTypeNotSupported                                                     = 0x8009301E
	OssRealDllNotLinked                                                     = 0x8009301F
	OssRealCodeNotLinked                                                    = 0x80093020
	OssOutOfRange                                                           = 0x80093021
	OssCopierDllNotLinked                                                   = 0x80093022
	OssConstraintDllNotLinked                                               = 0x80093023
	OssComparatorDllNotLinked                                               = 0x80093024
	OssComparatorCodeNotLinked                                              = 0x80093025
	OssMemMgrDllNotLinked                                                   = 0x80093026
	OssPdvDllNotLinked                                                      = 0x80093027
	OssPdvCodeNotLinked                                                     = 0x80093028
	OssApiDllNotLinked                                                      = 0x80093029
	OssBerderDllNotLinked                                                   = 0x8009302A
	OssPerDllNotLinked                                                      = 0x8009302B
	OssOpenTypeError                                                        = 0x8009302C
	OssMutexNotCreated                                                      = 0x8009302D
	OssCantCloseTraceFile                                                   = 0x8009302E
	CryptEAsn1Error                                                         = 0x80093100
	CryptEAsn1Internal                                                      = 0x80093101
	CryptEAsn1Eod                                                           = 0x80093102
	CryptEAsn1Corrupt                                                       = 0x80093103
	CryptEAsn1Large                                                         = 0x80093104
	CryptEAsn1Constraint                                                    = 0x80093105
	CryptEAsn1Memory                                                        = 0x80093106
	CryptEAsn1Overflow                                                      = 0x80093107
	CryptEAsn1Badpdu                                                        = 0x80093108
	CryptEAsn1Badargs                                                       = 0x80093109
	CryptEAsn1Badreal                                                       = 0x8009310A
	CryptEAsn1Badtag                                                        = 0x8009310B
	CryptEAsn1Choice                                                        = 0x8009310C
	CryptEAsn1Rule                                                          = 0x8009310D
	CryptEAsn1Utf8                                                          = 0x8009310E
	CryptEAsn1PduType                                                       = 0x80093133
	CryptEAsn1Nyi                                                           = 0x80093134
	CryptEAsn1Extended                                                      = 0x80093201
	CryptEAsn1Noeod                                                         = 0x80093202
	CertsrvEBadRequestsubject                                               = 0x80094001
	CertsrvENoRequest                                                       = 0x80094002
	CertsrvEBadRequeststatus                                                = 0x80094003
	CertsrvEPropertyEmpty                                                   = 0x80094004
	CertsrvEInvalidCaCertificate                                            = 0x80094005
	CertsrvEServerSuspended                                                 = 0x80094006
	CertsrvEEncodingLength                                                  = 0x80094007
	CertsrvERoleconflict                                                    = 0x80094008
	CertsrvERestrictedofficer                                               = 0x80094009
	CertsrvEKeyArchivalNotConfigured                                        = 0x8009400A
	CertsrvENoValidKra                                                      = 0x8009400B
	CertsrvEBadRequestKeyArchival                                           = 0x8009400C
	CertsrvENoCaadminDefined                                                = 0x8009400D
	CertsrvEBadRenewalCertAttribute                                         = 0x8009400E
	CertsrvENoDbSessions                                                    = 0x8009400F
	CertsrvEAlignmentFault                                                  = 0x80094010
	CertsrvEEnrollDenied                                                    = 0x80094011
	CertsrvETemplateDenied                                                  = 0x80094012
	CertsrvEDownlevelDcSslOrUpgrade                                         = 0x80094013
	CertsrvEAdminDeniedRequest                                              = 0x80094014
	CertsrvENoPolicyServer                                                  = 0x80094015
	CertsrvEUnsupportedCertType                                             = 0x80094800
	CertsrvENoCertType                                                      = 0x80094801
	CertsrvETemplateConflict                                                = 0x80094802
	CertsrvESubjectAltNameRequired                                          = 0x80094803
	CertsrvEArchivedKeyRequired                                             = 0x80094804
	CertsrvESmimeRequired                                                   = 0x80094805
	CertsrvEBadRenewalSubject                                               = 0x80094806
	CertsrvEBadTemplateVersion                                              = 0x80094807
	CertsrvETemplatePolicyRequired                                          = 0x80094808
	CertsrvESignaturePolicyRequired                                         = 0x80094809
	CertsrvESignatureCount                                                  = 0x8009480A
	CertsrvESignatureRejected                                               = 0x8009480B
	CertsrvEIssuancePolicyRequired                                          = 0x8009480C
	CertsrvESubjectUpnRequired                                              = 0x8009480D
	CertsrvESubjectDirectoryGuidRequired                                    = 0x8009480E
	CertsrvESubjectDnsRequired                                              = 0x8009480F
	CertsrvEArchivedKeyUnexpected                                           = 0x80094810
	CertsrvEKeyLength                                                       = 0x80094811
	CertsrvESubjectEmailRequired                                            = 0x80094812
	CertsrvEUnknownCertType                                                 = 0x80094813
	CertsrvECertTypeOverlap                                                 = 0x80094814
	CertsrvETooManySignatures                                               = 0x80094815
	XenrollEKeyNotExportable                                                = 0x80095000
	XenrollECannotAddRootCert                                               = 0x80095001
	XenrollEResponseKaHashNotFound                                          = 0x80095002
	XenrollEResponseUnexpectedKaHash                                        = 0x80095003
	XenrollEResponseKaHashMismatch                                          = 0x80095004
	XenrollEKeyspecSmimeMismatch                                            = 0x80095005
	TrustESystemError                                                       = 0x80096001
	TrustENoSignerCert                                                      = 0x80096002
	TrustECounterSigner                                                     = 0x80096003
	TrustECertSignature                                                     = 0x80096004
	TrustETimeStamp                                                         = 0x80096005
	TrustEBadDigest                                                         = 0x80096010
	TrustEBasicConstraints                                                  = 0x80096019
	TrustEFinancialCriteria                                                 = 0x8009601E
	MssipotfEOutofmemrange                                                  = 0x80097001
	MssipotfECantgetobject                                                  = 0x80097002
	MssipotfENoheadtable                                                    = 0x80097003
	MssipotfEBadMagicnumber                                                 = 0x80097004
	MssipotfEBadOffsetTable                                                 = 0x80097005
	MssipotfETableTagorder                                                  = 0x80097006
	MssipotfETableLongword                                                  = 0x80097007
	MssipotfEBadFirstTablePlacement                                         = 0x80097008
	MssipotfETablesOverlap                                                  = 0x80097009
	MssipotfETablePadbytes                                                  = 0x8009700A
	MssipotfEFiletoosmall                                                   = 0x8009700B
	MssipotfETableChecksum                                                  = 0x8009700C
	MssipotfEFileChecksum                                                   = 0x8009700D
	MssipotfEFailedPolicy                                                   = 0x80097010
	MssipotfEFailedHintsCheck                                               = 0x80097011
	MssipotfENotOpentype                                                    = 0x80097012
	MssipotfEFile                                                           = 0x80097013
	MssipotfECrypt                                                          = 0x80097014
	MssipotfEBadversion                                                     = 0x80097015
	MssipotfEDsigStructure                                                  = 0x80097016
	MssipotfEPconstCheck                                                    = 0x80097017
	MssipotfEStructure                                                      = 0x80097018
	ErrorCredRequiresConfirmation                                           = 0x80097019
	NteOpOk                                                                 = 0
	TrustEProviderUnknown                                                   = 0x800B0001
	TrustEActionUnknown                                                     = 0x800B0002
	TrustESubjectFormUnknown                                                = 0x800B0003
	TrustESubjectNotTrusted                                                 = 0x800B0004
	DigsigEEncode                                                           = 0x800B0005
	DigsigEDecode                                                           = 0x800B0006
	DigsigEExtensibility                                                    = 0x800B0007
	DigsigECrypto                                                           = 0x800B0008
	PersistESizedefinite                                                    = 0x800B0009
	PersistESizeindefinite                                                  = 0x800B000A
	PersistENotselfsizing                                                   = 0x800B000B
	TrustENosignature                                                       = 0x800B0100
	CertEExpired                                                            = 0x800B0101
	CertEValidityperiodnesting                                              = 0x800B0102
	CertERole                                                               = 0x800B0103
	CertEPathlenconst                                                       = 0x800B0104
	CertECritical                                                           = 0x800B0105
	CertEPurpose                                                            = 0x800B0106
	CertEIssuerchaining                                                     = 0x800B0107
	CertEMalformed                                                          = 0x800B0108
	CertEUntrustedroot                                                      = 0x800B0109
	CertEChaining                                                           = 0x800B010A
	TrustEFail                                                              = 0x800B010B
	CertERevoked                                                            = 0x800B010C
	CertEUntrustedtestroot                                                  = 0x800B010D
	CertERevocationFailure                                                  = 0x800B010E
	CertECnNoMatch                                                          = 0x800B010F
	CertEWrongUsage                                                         = 0x800B0110
	TrustEExplicitDistrust                                                  = 0x800B0111
	CertEUntrustedca                                                        = 0x800B0112
	CertEInvalidPolicy                                                      = 0x800B0113
	CertEInvalidName                                                        = 0x800B0114
	SpapiEExpectedSectionName                                               = 0x800F0000
	SpapiEBadSectionNameLine                                                = 0x800F0001
	SpapiESectionNameTooLong                                                = 0x800F0002
	SpapiEGeneralSyntax                                                     = 0x800F0003
	SpapiEWrongInfStyle                                                     = 0x800F0100
	SpapiESectionNotFound                                                   = 0x800F0101
	SpapiELineNotFound                                                      = 0x800F0102
	SpapiENoBackup                                                          = 0x800F0103
	SpapiENoAssociatedClass                                                 = 0x800F0200
	SpapiEClassMismatch                                                     = 0x800F0201
	SpapiEDuplicateFound                                                    = 0x800F0202
	SpapiENoDriverSelected                                                  = 0x800F0203
	SpapiEKeyDoesNotExist                                                   = 0x800F0204
	SpapiEInvalidDevinstName                                                = 0x800F0205
	SpapiEInvalidClass                                                      = 0x800F0206
	SpapiEDevinstAlreadyExists                                              = 0x800F0207
	SpapiEDevinfoNotRegistered                                              = 0x800F0208
	SpapiEInvalidRegProperty                                                = 0x800F0209
	SpapiENoInf                                                             = 0x800F020A
	SpapiENoSuchDevinst                                                     = 0x800F020B
	SpapiECantLoadClassIcon                                                 = 0x800F020C
	SpapiEInvalidClassInstaller                                             = 0x800F020D
	SpapiEDiDoDefault                                                       = 0x800F020E
	SpapiEDiNofilecopy                                                      = 0x800F020F
	SpapiEInvalidHwprofile                                                  = 0x800F0210
	SpapiENoDeviceSelected                                                  = 0x800F0211
	SpapiEDevinfoListLocked                                                 = 0x800F0212
	SpapiEDevinfoDataLocked                                                 = 0x800F0213
	SpapiEDiBadPath                                                         = 0x800F0214
	SpapiENoClassinstallParams                                              = 0x800F0215
	SpapiEFilequeueLocked                                                   = 0x800F0216
	SpapiEBadServiceInstallsect                                             = 0x800F0217
	SpapiENoClassDriverList                                                 = 0x800F0218
	SpapiENoAssociatedService                                               = 0x800F0219
	SpapiENoDefaultDeviceInterface                                          = 0x800F021A
	SpapiEDeviceInterfaceActive                                             = 0x800F021B
	SpapiEDeviceInterfaceRemoved                                            = 0x800F021C
	SpapiEBadInterfaceInstallsect                                           = 0x800F021D
	SpapiENoSuchInterfaceClass                                              = 0x800F021E
	SpapiEInvalidReferenceString                                            = 0x800F021F
	SpapiEInvalidMachinename                                                = 0x800F0220
	SpapiERemoteCommFailure                                                 = 0x800F0221
	SpapiEMachineUnavailable                                                = 0x800F0222
	SpapiENoConfigmgrServices                                               = 0x800F0223
	SpapiEInvalidProppageProvider                                           = 0x800F0224
	SpapiENoSuchDeviceInterface                                             = 0x800F0225
	SpapiEDiPostprocessingRequired                                          = 0x800F0226
	SpapiEInvalidCoinstaller                                                = 0x800F0227
	SpapiENoCompatDrivers                                                   = 0x800F0228
	SpapiENoDeviceIcon                                                      = 0x800F0229
	SpapiEInvalidInfLogconfig                                               = 0x800F022A
	SpapiEDiDontInstall                                                     = 0x800F022B
	SpapiEInvalidFilterDriver                                               = 0x800F022C
	SpapiENonWindowsNtDriver                                                = 0x800F022D
	SpapiENonWindowsDriver                                                  = 0x800F022E
	SpapiENoCatalogForOemInf                                                = 0x800F022F
	SpapiEDevinstallQueueNonnative                                          = 0x800F0230
	SpapiENotDisableable                                                    = 0x800F0231
	SpapiECantRemoveDevinst                                                 = 0x800F0232
	SpapiEInvalidTarget                                                     = 0x800F0233
	SpapiEDriverNonnative                                                   = 0x800F0234
	SpapiEInWow64                                                           = 0x800F0235
	SpapiESetSystemRestorePoint                                             = 0x800F0236
	SpapiEIncorrectlyCopiedInf                                              = 0x800F0237
	SpapiESceDisabled                                                       = 0x800F0238
	SpapiEUnknownException                                                  = 0x800F0239
	SpapiEPnpRegistryError                                                  = 0x800F023A
	SpapiERemoteRequestUnsupported                                          = 0x800F023B
	SpapiENotAnInstalledOemInf                                              = 0x800F023C
	SpapiEInfInUseByDevices                                                 = 0x800F023D
	SpapiEDiFunctionObsolete                                                = 0x800F023E
	SpapiENoAuthenticodeCatalog                                             = 0x800F023F
	SpapiEAuthenticodeDisallowed                                            = 0x800F0240
	SpapiEAuthenticodeTrustedPublisher                                      = 0x800F0241
	SpapiEAuthenticodeTrustNotEstablished                                   = 0x800F0242
	SpapiEAuthenticodePublisherNotTrusted                                   = 0x800F0243
	SpapiESignatureOsattributeMismatch                                      = 0x800F0244
	SpapiEOnlyValidateViaAuthenticode                                       = 0x800F0245
	SpapiEDeviceInstallerNotReady                                           = 0x800F0246
	SpapiEDriverStoreAddFailed                                              = 0x800F0247
	SpapiEDeviceInstallBlocked                                              = 0x800F0248
	SpapiEDriverInstallBlocked                                              = 0x800F0249
	SpapiEWrongInfType                                                      = 0x800F024A
	SpapiEFileHashNotInCatalog                                              = 0x800F024B
	SpapiEDriverStoreDeleteFailed                                           = 0x800F024C
	SpapiEUnrecoverableStackOverflow                                        = 0x800F0300
	SpapiEErrorNotInstalled                                                 = 0x800F1000
	ScardFInternalError                                                     = 0x80100001
	ScardECancelled                                                         = 0x80100002
	ScardEInvalidHandle                                                     = 0x80100003
	ScardEInvalidParameter                                                  = 0x80100004
	ScardEInvalidTarget                                                     = 0x80100005
	ScardENoMemory                                                          = 0x80100006
	ScardFWaitedTooLong                                                     = 0x80100007
	ScardEInsufficientBuffer                                                = 0x80100008
	ScardEUnknownReader                                                     = 0x80100009
	ScardETimeout                                                           = 0x8010000A
	ScardESharingViolation                                                  = 0x8010000B
	ScardENoSmartcard                                                       = 0x8010000C
	ScardEUnknownCard                                                       = 0x8010000D
	ScardECantDispose                                                       = 0x8010000E
	ScardEProtoMismatch                                                     = 0x8010000F
	ScardENotReady                                                          = 0x80100010
	ScardEInvalidValue                                                      = 0x80100011
	ScardESystemCancelled                                                   = 0x80100012
	ScardFCommError                                                         = 0x80100013
	ScardFUnknownError                                                      = 0x80100014
	ScardEInvalidAtr                                                        = 0x80100015
	ScardENotTransacted                                                     = 0x80100016
	ScardEReaderUnavailable                                                 = 0x80100017
	ScardPShutdown                                                          = 0x80100018
	ScardEPciTooSmall                                                       = 0x80100019
	ScardEReaderUnsupported                                                 = 0x8010001A
	ScardEDuplicateReader                                                   = 0x8010001B
	ScardECardUnsupported                                                   = 0x8010001C
	ScardENoService                                                         = 0x8010001D
	ScardEServiceStopped                                                    = 0x8010001E
	ScardEUnexpected                                                        = 0x8010001F
	ScardEIccInstallation                                                   = 0x80100020
	ScardEIccCreateorder                                                    = 0x80100021
	ScardEUnsupportedFeature                                                = 0x80100022
	ScardEDirNotFound                                                       = 0x80100023
	ScardEFileNotFound                                                      = 0x80100024
	ScardENoDir                                                             = 0x80100025
	ScardENoFile                                                            = 0x80100026
	ScardENoAccess                                                          = 0x80100027
	ScardEWriteTooMany                                                      = 0x80100028
	ScardEBadSeek                                                           = 0x80100029
	ScardEInvalidChv                                                        = 0x8010002A
	ScardEUnknownResMng                                                     = 0x8010002B
	ScardENoSuchCertificate                                                 = 0x8010002C
	ScardECertificateUnavailable                                            = 0x8010002D
	ScardENoReadersAvailable                                                = 0x8010002E
	ScardECommDataLost                                                      = 0x8010002F
	ScardENoKeyContainer                                                    = 0x80100030
	ScardEServerTooBusy                                                     = 0x80100031
	ScardEPinCacheExpired                                                   = 0x80100032
	ScardENoPinCache                                                        = 0x80100033
	ScardEReadOnlyCard                                                      = 0x80100034
	ScardWUnsupportedCard                                                   = 0x80100065
	ScardWUnresponsiveCard                                                  = 0x80100066
	ScardWUnpoweredCard                                                     = 0x80100067
	ScardWResetCard                                                         = 0x80100068
	ScardWRemovedCard                                                       = 0x80100069
	ScardWSecurityViolation                                                 = 0x8010006A
	ScardWWrongChv                                                          = 0x8010006B
	ScardWChvBlocked                                                        = 0x8010006C
	ScardWEof                                                               = 0x8010006D
	ScardWCancelledByUser                                                   = 0x8010006E
	ScardWCardNotAuthenticated                                              = 0x8010006F
	ScardWCacheItemNotFound                                                 = 0x80100070
	ScardWCacheItemStale                                                    = 0x80100071
	ScardWCacheItemTooBig                                                   = 0x80100072
	ComadminEObjecterrors                                                   = 0x80110401
	ComadminEObjectinvalid                                                  = 0x80110402
	ComadminEKeymissing                                                     = 0x80110403
	ComadminEAlreadyinstalled                                               = 0x80110404
	ComadminEAppFileWritefail                                               = 0x80110407
	ComadminEAppFileReadfail                                                = 0x80110408
	ComadminEAppFileVersion                                                 = 0x80110409
	ComadminEBadpath                                                        = 0x8011040A
	ComadminEApplicationexists                                              = 0x8011040B
	ComadminERoleexists                                                     = 0x8011040C
	ComadminECantcopyfile                                                   = 0x8011040D
	ComadminENouser                                                         = 0x8011040F
	ComadminEInvaliduserids                                                 = 0x80110410
	ComadminENoregistryclsid                                                = 0x80110411
	ComadminEBadregistryprogid                                              = 0x80110412
	ComadminEAuthenticationlevel                                            = 0x80110413
	ComadminEUserpasswdnotvalid                                             = 0x80110414
	ComadminEClsidoriidmismatch                                             = 0x80110418
	ComadminERemoteinterface                                                = 0x80110419
	ComadminEDllregisterserver                                              = 0x8011041A
	ComadminENoservershare                                                  = 0x8011041B
	ComadminEDllloadfailed                                                  = 0x8011041D
	ComadminEBadregistrylibid                                               = 0x8011041E
	ComadminEAppdirnotfound                                                 = 0x8011041F
	ComadminERegistrarfailed                                                = 0x80110423
	ComadminECompfileDoesnotexist                                           = 0x80110424
	ComadminECompfileLoaddllfail                                            = 0x80110425
	ComadminECompfileGetclassobj                                            = 0x80110426
	ComadminECompfileClassnotavail                                          = 0x80110427
	ComadminECompfileBadtlb                                                 = 0x80110428
	ComadminECompfileNotinstallable                                         = 0x80110429
	ComadminENotchangeable                                                  = 0x8011042A
	ComadminENotdeleteable                                                  = 0x8011042B
	ComadminESession                                                        = 0x8011042C
	ComadminECompMoveLocked                                                 = 0x8011042D
	ComadminECompMoveBadDest                                                = 0x8011042E
	ComadminERegistertlb                                                    = 0x80110430
	ComadminESystemapp                                                      = 0x80110433
	ComadminECompfileNoregistrar                                            = 0x80110434
	ComadminECoreqcompinstalled                                             = 0x80110435
	ComadminEServicenotinstalled                                            = 0x80110436
	ComadminEPropertysavefailed                                             = 0x80110437
	ComadminEObjectexists                                                   = 0x80110438
	ComadminEComponentexists                                                = 0x80110439
	ComadminERegfileCorrupt                                                 = 0x8011043B
	ComadminEPropertyOverflow                                               = 0x8011043C
	ComadminENotinregistry                                                  = 0x8011043E
	ComadminEObjectnotpoolable                                              = 0x8011043F
	ComadminEApplidMatchesClsid                                             = 0x80110446
	ComadminERoleDoesNotExist                                               = 0x80110447
	ComadminEStartAppNeedsComponents                                        = 0x80110448
	ComadminERequiresDifferentPlatform                                      = 0x80110449
	ComadminECanNotExportAppProxy                                           = 0x8011044A
	ComadminECanNotStartApp                                                 = 0x8011044B
	ComadminECanNotExportSysApp                                             = 0x8011044C
	ComadminECantSubscribeToComponent                                       = 0x8011044D
	ComadminEEventclassCantBeSubscriber                                     = 0x8011044E
	ComadminELibAppProxyIncompatible                                        = 0x8011044F
	ComadminEBasePartitionOnly                                              = 0x80110450
	ComadminEStartAppDisabled                                               = 0x80110451
	ComadminECatDuplicatePartitionName                                      = 0x80110457
	ComadminECatInvalidPartitionName                                        = 0x80110458
	ComadminECatPartitionInUse                                              = 0x80110459
	ComadminEFilePartitionDuplicateFiles                                    = 0x8011045A
	ComadminECatImportedComponentsNotAllowed                                = 0x8011045B
	ComadminEAmbiguousApplicationName                                       = 0x8011045C
	ComadminEAmbiguousPartitionName                                         = 0x8011045D
	ComadminERegdbNotinitialized                                            = 0x80110472
	ComadminERegdbNotopen                                                   = 0x80110473
	ComadminERegdbSystemerr                                                 = 0x80110474
	ComadminERegdbAlreadyrunning                                            = 0x80110475
	ComadminEMigVersionnotsupported                                         = 0x80110480
	ComadminEMigSchemanotfound                                              = 0x80110481
	ComadminECatBitnessmismatch                                             = 0x80110482
	ComadminECatUnacceptablebitness                                         = 0x80110483
	ComadminECatWrongappbitness                                             = 0x80110484
	ComadminECatPauseResumeNotSupported                                     = 0x80110485
	ComadminECatServerfault                                                 = 0x80110486
	ComqcEApplicationNotQueued                                              = 0x80110600
	ComqcENoQueueableInterfaces                                             = 0x80110601
	ComqcEQueuingServiceNotAvailable                                        = 0x80110602
	ComqcENoIpersiststream                                                  = 0x80110603
	ComqcEBadMessage                                                        = 0x80110604
	ComqcEUnauthenticated                                                   = 0x80110605
	ComqcEUntrustedEnqueuer                                                 = 0x80110606
	MsdtcEDuplicateResource                                                 = 0x80110701
	ComadminEObjectParentMissing                                            = 0x80110808
	ComadminEObjectDoesNotExist                                             = 0x80110809
	ComadminEAppNotRunning                                                  = 0x8011080A
	ComadminEInvalidPartition                                               = 0x8011080B
	ComadminESvcappNotPoolableOrRecyclable                                  = 0x8011080D
	ComadminEUserInSet                                                      = 0x8011080E
	ComadminECantrecyclelibraryapps                                         = 0x8011080F
	ComadminECantrecycleserviceapps                                         = 0x80110811
	ComadminEProcessalreadyrecycled                                         = 0x80110812
	ComadminEPausedprocessmaynotberecycled                                  = 0x80110813
	ComadminECantmakeinprocservice                                          = 0x80110814
	ComadminEProgidinusebyclsid                                             = 0x80110815
	ComadminEDefaultPartitionNotInSet                                       = 0x80110816
	ComadminERecycledprocessmaynotbepaused                                  = 0x80110817
	ComadminEPartitionAccessdenied                                          = 0x80110818
	ComadminEPartitionMsiOnly                                               = 0x80110819
	ComadminELegacycompsNotAllowedIn10Format                                = 0x8011081A
	ComadminELegacycompsNotAllowedInNonbasePartitions                       = 0x8011081B
	ComadminECompMoveSource                                                 = 0x8011081C
	ComadminECompMoveDest                                                   = 0x8011081D
	ComadminECompMovePrivate                                                = 0x8011081E
	ComadminEBasepartitionRequiredInSet                                     = 0x8011081F
	ComadminECannotAliasEventclass                                          = 0x80110820
	ComadminEPrivateAccessdenied                                            = 0x80110821
	ComadminESaferinvalid                                                   = 0x80110822
	ComadminERegistryAccessdenied                                           = 0x80110823
	ComadminEPartitionsDisabled                                             = 0x80110824
	ErrorFltIoComplete                                                      = 0x001F0001
	ErrorFltNoHandlerDefined                                                = 0x801F0001
	ErrorFltContextAlreadyDefined                                           = 0x801F0002
	ErrorFltInvalidAsynchronousRequest                                      = 0x801F0003
	ErrorFltDisallowFastIo                                                  = 0x801F0004
	ErrorFltInvalidNameRequest                                              = 0x801F0005
	ErrorFltNotSafeToPostOperation                                          = 0x801F0006
	ErrorFltNotInitialized                                                  = 0x801F0007
	ErrorFltFilterNotReady                                                  = 0x801F0008
	ErrorFltPostOperationCleanup                                            = 0x801F0009
	ErrorFltInternalError                                                   = 0x801F000A
	ErrorFltDeletingObject                                                  = 0x801F000B
	ErrorFltMustBeNonpagedPool                                              = 0x801F000C
	ErrorFltDuplicateEntry                                                  = 0x801F000D
	ErrorFltCbdqDisabled                                                    = 0x801F000E
	ErrorFltDoNotAttach                                                     = 0x801F000F
	ErrorFltDoNotDetach                                                     = 0x801F0010
	ErrorFltInstanceAltitudeCollision                                       = 0x801F0011
	ErrorFltInstanceNameCollision                                           = 0x801F0012
	ErrorFltFilterNotFound                                                  = 0x801F0013
	ErrorFltVolumeNotFound                                                  = 0x801F0014
	ErrorFltInstanceNotFound                                                = 0x801F0015
	ErrorFltContextAllocationNotFound                                       = 0x801F0016
	ErrorFltInvalidContextRegistration                                      = 0x801F0017
	ErrorFltNameCacheMiss                                                   = 0x801F0018
	ErrorFltNoDeviceObject                                                  = 0x801F0019
	ErrorFltVolumeAlreadyMounted                                            = 0x801F001A
	ErrorFltAlreadyEnlisted                                                 = 0x801F001B
	ErrorFltContextAlreadyLinked                                            = 0x801F001C
	ErrorFltNoWaiterForReply                                                = 0x801F0020
	ErrorHungDisplayDriverThread                                            = 0x80260001
	DwmECompositiondisabled                                                 = 0x80263001
	DwmERemotingNotSupported                                                = 0x80263002
	DwmENoRedirectionSurfaceAvailable                                       = 0x80263003
	DwmENotQueuingPresents                                                  = 0x80263004
	DwmEAdapterNotFound                                                     = 0x80263005
	DwmSGdiRedirectionSurface                                               = 0x00263005
	ErrorMonitorNoDescriptor                                                = 0x00261001
	ErrorMonitorUnknownDescriptorFormat                                     = 0x00261002
	ErrorMonitorInvalidDescriptorChecksum                                   = 0xC0261003
	ErrorMonitorInvalidStandardTimingBlock                                  = 0xC0261004
	ErrorMonitorWmiDatablockRegistrationFailed                              = 0xC0261005
	ErrorMonitorInvalidSerialNumberMondscBlock                              = 0xC0261006
	ErrorMonitorInvalidUserFriendlyMondscBlock                              = 0xC0261007
	ErrorMonitorNoMoreDescriptorData                                        = 0xC0261008
	ErrorMonitorInvalidDetailedTimingBlock                                  = 0xC0261009
	ErrorMonitorInvalidManufactureDate                                      = 0xC026100A
	ErrorGraphicsNotExclusiveModeOwner                                      = 0xC0262000
	ErrorGraphicsInsufficientDmaBuffer                                      = 0xC0262001
	ErrorGraphicsInvalidDisplayAdapter                                      = 0xC0262002
	ErrorGraphicsAdapterWasReset                                            = 0xC0262003
	ErrorGraphicsInvalidDriverModel                                         = 0xC0262004
	ErrorGraphicsPresentModeChanged                                         = 0xC0262005
	ErrorGraphicsPresentOccluded                                            = 0xC0262006
	ErrorGraphicsPresentDenied                                              = 0xC0262007
	ErrorGraphicsCannotcolorconvert                                         = 0xC0262008
	ErrorGraphicsDriverMismatch                                             = 0xC0262009
	ErrorGraphicsPartialDataPopulated                                       = 0x4026200A
	ErrorGraphicsPresentRedirectionDisabled                                 = 0xC026200B
	ErrorGraphicsPresentUnoccluded                                          = 0xC026200C
	ErrorGraphicsNoVideoMemory                                              = 0xC0262100
	ErrorGraphicsCantLockMemory                                             = 0xC0262101
	ErrorGraphicsAllocationBusy                                             = 0xC0262102
	ErrorGraphicsTooManyReferences                                          = 0xC0262103
	ErrorGraphicsTryAgainLater                                              = 0xC0262104
	ErrorGraphicsTryAgainNow                                                = 0xC0262105
	ErrorGraphicsAllocationInvalid                                          = 0xC0262106
	ErrorGraphicsUnswizzlingApertureUnavailable                             = 0xC0262107
	ErrorGraphicsUnswizzlingApertureUnsupported                             = 0xC0262108
	ErrorGraphicsCantEvictPinnedAllocation                                  = 0xC0262109
	ErrorGraphicsInvalidAllocationUsage                                     = 0xC0262110
	ErrorGraphicsCantRenderLockedAllocation                                 = 0xC0262111
	ErrorGraphicsAllocationClosed                                           = 0xC0262112
	ErrorGraphicsInvalidAllocationInstance                                  = 0xC0262113
	ErrorGraphicsInvalidAllocationHandle                                    = 0xC0262114
	ErrorGraphicsWrongAllocationDevice                                      = 0xC0262115
	ErrorGraphicsAllocationContentLost                                      = 0xC0262116
	ErrorGraphicsGpuExceptionOnDevice                                       = 0xC0262200
	ErrorGraphicsInvalidVidpnTopology                                       = 0xC0262300
	ErrorGraphicsVidpnTopologyNotSupported                                  = 0xC0262301
	ErrorGraphicsVidpnTopologyCurrentlyNotSupported                         = 0xC0262302
	ErrorGraphicsInvalidVidpn                                               = 0xC0262303
	ErrorGraphicsInvalidVideoPresentSource                                  = 0xC0262304
	ErrorGraphicsInvalidVideoPresentTarget                                  = 0xC0262305
	ErrorGraphicsVidpnModalityNotSupported                                  = 0xC0262306
	ErrorGraphicsModeNotPinned                                              = 0x00262307
	ErrorGraphicsInvalidVidpnSourcemodeset                                  = 0xC0262308
	ErrorGraphicsInvalidVidpnTargetmodeset                                  = 0xC0262309
	ErrorGraphicsInvalidFrequency                                           = 0xC026230A
	ErrorGraphicsInvalidActiveRegion                                        = 0xC026230B
	ErrorGraphicsInvalidTotalRegion                                         = 0xC026230C
	ErrorGraphicsInvalidVideoPresentSourceMode                              = 0xC0262310
	ErrorGraphicsInvalidVideoPresentTargetMode                              = 0xC0262311
	ErrorGraphicsPinnedModeMustRemainInSet                                  = 0xC0262312
	ErrorGraphicsPathAlreadyInTopology                                      = 0xC0262313
	ErrorGraphicsModeAlreadyInModeset                                       = 0xC0262314
	ErrorGraphicsInvalidVideopresentsourceset                               = 0xC0262315
	ErrorGraphicsInvalidVideopresenttargetset                               = 0xC0262316
	ErrorGraphicsSourceAlreadyInSet                                         = 0xC0262317
	ErrorGraphicsTargetAlreadyInSet                                         = 0xC0262318
	ErrorGraphicsInvalidVidpnPresentPath                                    = 0xC0262319
	ErrorGraphicsNoRecommendedVidpnTopology                                 = 0xC026231A
	ErrorGraphicsInvalidMonitorFrequencyrangeset                            = 0xC026231B
	ErrorGraphicsInvalidMonitorFrequencyrange                               = 0xC026231C
	ErrorGraphicsFrequencyrangeNotInSet                                     = 0xC026231D
	ErrorGraphicsNoPreferredMode                                            = 0x0026231E
	ErrorGraphicsFrequencyrangeAlreadyInSet                                 = 0xC026231F
	ErrorGraphicsStaleModeset                                               = 0xC0262320
	ErrorGraphicsInvalidMonitorSourcemodeset                                = 0xC0262321
	ErrorGraphicsInvalidMonitorSourceMode                                   = 0xC0262322
	ErrorGraphicsNoRecommendedFunctionalVidpn                               = 0xC0262323
	ErrorGraphicsModeIdMustBeUnique                                         = 0xC0262324
	ErrorGraphicsEmptyAdapterMonitorModeSupportIntersection                 = 0xC0262325
	ErrorGraphicsVideoPresentTargetsLessThanSources                         = 0xC0262326
	ErrorGraphicsPathNotInTopology                                          = 0xC0262327
	ErrorGraphicsAdapterMustHaveAtLeastOneSource                            = 0xC0262328
	ErrorGraphicsAdapterMustHaveAtLeastOneTarget                            = 0xC0262329
	ErrorGraphicsInvalidMonitordescriptorset                                = 0xC026232A
	ErrorGraphicsInvalidMonitordescriptor                                   = 0xC026232B
	ErrorGraphicsMonitordescriptorNotInSet                                  = 0xC026232C
	ErrorGraphicsMonitordescriptorAlreadyInSet                              = 0xC026232D
	ErrorGraphicsMonitordescriptorIdMustBeUnique                            = 0xC026232E
	ErrorGraphicsInvalidVidpnTargetSubsetType                               = 0xC026232F
	ErrorGraphicsResourcesNotRelated                                        = 0xC0262330
	ErrorGraphicsSourceIdMustBeUnique                                       = 0xC0262331
	ErrorGraphicsTargetIdMustBeUnique                                       = 0xC0262332
	ErrorGraphicsNoAvailableVidpnTarget                                     = 0xC0262333
	ErrorGraphicsMonitorCouldNotBeAssociatedWithAdapter                     = 0xC0262334
	ErrorGraphicsNoVidpnmgr                                                 = 0xC0262335
	ErrorGraphicsNoActiveVidpn                                              = 0xC0262336
	ErrorGraphicsStaleVidpnTopology                                         = 0xC0262337
	ErrorGraphicsMonitorNotConnected                                        = 0xC0262338
	ErrorGraphicsSourceNotInTopology                                        = 0xC0262339
	ErrorGraphicsInvalidPrimarysurfaceSize                                  = 0xC026233A
	ErrorGraphicsInvalidVisibleregionSize                                   = 0xC026233B
	ErrorGraphicsInvalidStride                                              = 0xC026233C
	ErrorGraphicsInvalidPixelformat                                         = 0xC026233D
	ErrorGraphicsInvalidColorbasis                                          = 0xC026233E
	ErrorGraphicsInvalidPixelvalueaccessmode                                = 0xC026233F
	ErrorGraphicsTargetNotInTopology                                        = 0xC0262340
	ErrorGraphicsNoDisplayModeManagementSupport                             = 0xC0262341
	ErrorGraphicsVidpnSourceInUse                                           = 0xC0262342
	ErrorGraphicsCantAccessActiveVidpn                                      = 0xC0262343
	ErrorGraphicsInvalidPathImportanceOrdinal                               = 0xC0262344
	ErrorGraphicsInvalidPathContentGeometryTransformation                   = 0xC0262345
	ErrorGraphicsPathContentGeometryTransformationNotSupported              = 0xC0262346
	ErrorGraphicsInvalidGammaRamp                                           = 0xC0262347
	ErrorGraphicsGammaRampNotSupported                                      = 0xC0262348
	ErrorGraphicsMultisamplingNotSupported                                  = 0xC0262349
	ErrorGraphicsModeNotInModeset                                           = 0xC026234A
	ErrorGraphicsDatasetIsEmpty                                             = 0x0026234B
	ErrorGraphicsNoMoreElementsInDataset                                    = 0x0026234C
	ErrorGraphicsInvalidVidpnTopologyRecommendationReason                   = 0xC026234D
	ErrorGraphicsInvalidPathContentType                                     = 0xC026234E
	ErrorGraphicsInvalidCopyprotectionType                                  = 0xC026234F
	ErrorGraphicsUnassignedModesetAlreadyExists                             = 0xC0262350
	ErrorGraphicsPathContentGeometryTransformationNotPinned                 = 0x00262351
	ErrorGraphicsInvalidScanlineOrdering                                    = 0xC0262352
	ErrorGraphicsTopologyChangesNotAllowed                                  = 0xC0262353
	ErrorGraphicsNoAvailableImportanceOrdinals                              = 0xC0262354
	ErrorGraphicsIncompatiblePrivateFormat                                  = 0xC0262355
	ErrorGraphicsInvalidModePruningAlgorithm                                = 0xC0262356
	ErrorGraphicsInvalidMonitorCapabilityOrigin                             = 0xC0262357
	ErrorGraphicsInvalidMonitorFrequencyrangeConstraint                     = 0xC0262358
	ErrorGraphicsMaxNumPathsReached                                         = 0xC0262359
	ErrorGraphicsCancelVidpnTopologyAugmentation                            = 0xC026235A
	ErrorGraphicsInvalidClientType                                          = 0xC026235B
	ErrorGraphicsClientvidpnNotSet                                          = 0xC026235C
	ErrorGraphicsSpecifiedChildAlreadyConnected                             = 0xC0262400
	ErrorGraphicsChildDescriptorNotSupported                                = 0xC0262401
	ErrorGraphicsUnknownChildStatus                                         = 0x4026242F
	ErrorGraphicsNotALinkedAdapter                                          = 0xC0262430
	ErrorGraphicsLeadlinkNotEnumerated                                      = 0xC0262431
	ErrorGraphicsChainlinksNotEnumerated                                    = 0xC0262432
	ErrorGraphicsAdapterChainNotReady                                       = 0xC0262433
	ErrorGraphicsChainlinksNotStarted                                       = 0xC0262434
	ErrorGraphicsChainlinksNotPoweredOn                                     = 0xC0262435
	ErrorGraphicsInconsistentDeviceLinkState                                = 0xC0262436
	ErrorGraphicsLeadlinkStartDeferred                                      = 0x40262437
	ErrorGraphicsNotPostDeviceDriver                                        = 0xC0262438
	ErrorGraphicsPollingTooFrequently                                       = 0x40262439
	ErrorGraphicsStartDeferred                                              = 0x4026243A
	ErrorGraphicsAdapterAccessNotExcluded                                   = 0xC026243B
	ErrorGraphicsOpmNotSupported                                            = 0xC0262500
	ErrorGraphicsCoppNotSupported                                           = 0xC0262501
	ErrorGraphicsUabNotSupported                                            = 0xC0262502
	ErrorGraphicsOpmInvalidEncryptedParameters                              = 0xC0262503
	ErrorGraphicsOpmNoVideoOutputsExist                                     = 0xC0262505
	ErrorGraphicsOpmInternalError                                           = 0xC026250B
	ErrorGraphicsOpmInvalidHandle                                           = 0xC026250C
	ErrorGraphicsPvpInvalidCertificateLength                                = 0xC026250E
	ErrorGraphicsOpmSpanningModeEnabled                                     = 0xC026250F
	ErrorGraphicsOpmTheaterModeEnabled                                      = 0xC0262510
	ErrorGraphicsPvpHfsFailed                                               = 0xC0262511
	ErrorGraphicsOpmInvalidSrm                                              = 0xC0262512
	ErrorGraphicsOpmOutputDoesNotSupportHdcp                                = 0xC0262513
	ErrorGraphicsOpmOutputDoesNotSupportAcp                                 = 0xC0262514
	ErrorGraphicsOpmOutputDoesNotSupportCgmsa                               = 0xC0262515
	ErrorGraphicsOpmHdcpSrmNeverSet                                         = 0xC0262516
	ErrorGraphicsOpmResolutionTooHigh                                       = 0xC0262517
	ErrorGraphicsOpmAllHdcpHardwareAlreadyInUse                             = 0xC0262518
	ErrorGraphicsOpmVideoOutputNoLongerExists                               = 0xC026251A
	ErrorGraphicsOpmSessionTypeChangeInProgress                             = 0xC026251B
	ErrorGraphicsOpmVideoOutputDoesNotHaveCoppSemantics                     = 0xC026251C
	ErrorGraphicsOpmInvalidInformationRequest                               = 0xC026251D
	ErrorGraphicsOpmDriverInternalError                                     = 0xC026251E
	ErrorGraphicsOpmVideoOutputDoesNotHaveOpmSemantics                      = 0xC026251F
	ErrorGraphicsOpmSignalingNotSupported                                   = 0xC0262520
	ErrorGraphicsOpmInvalidConfigurationRequest                             = 0xC0262521
	ErrorGraphicsI2CNotSupported                                            = 0xC0262580
	ErrorGraphicsI2CDeviceDoesNotExist                                      = 0xC0262581
	ErrorGraphicsI2CErrorTransmittingData                                   = 0xC0262582
	ErrorGraphicsI2CErrorReceivingData                                      = 0xC0262583
	ErrorGraphicsDdcciVcpNotSupported                                       = 0xC0262584
	ErrorGraphicsDdcciInvalidData                                           = 0xC0262585
	ErrorGraphicsDdcciMonitorReturnedInvalidTimingStatusByte                = 0xC0262586
	ErrorGraphicsMcaInvalidCapabilitiesString                               = 0xC0262587
	ErrorGraphicsMcaInternalError                                           = 0xC0262588
	ErrorGraphicsDdcciInvalidMessageCommand                                 = 0xC0262589
	ErrorGraphicsDdcciInvalidMessageLength                                  = 0xC026258A
	ErrorGraphicsDdcciInvalidMessageChecksum                                = 0xC026258B
	ErrorGraphicsInvalidPhysicalMonitorHandle                               = 0xC026258C
	ErrorGraphicsMonitorNoLongerExists                                      = 0xC026258D
	ErrorGraphicsDdcciCurrentCurrentValueGreaterThanMaximumValue            = 0xC02625D8
	ErrorGraphicsMcaInvalidVcpVersion                                       = 0xC02625D9
	ErrorGraphicsMcaMonitorViolatesMccsSpecification                        = 0xC02625DA
	ErrorGraphicsMcaMccsVersionMismatch                                     = 0xC02625DB
	ErrorGraphicsMcaUnsupportedMccsVersion                                  = 0xC02625DC
	ErrorGraphicsMcaInvalidTechnologyTypeReturned                           = 0xC02625DE
	ErrorGraphicsMcaUnsupportedColorTemperature                             = 0xC02625DF
	ErrorGraphicsOnlyConsoleSessionSupported                                = 0xC02625E0
	ErrorGraphicsNoDisplayDeviceCorrespondsToName                           = 0xC02625E1
	ErrorGraphicsDisplayDeviceNotAttachedToDesktop                          = 0xC02625E2
	ErrorGraphicsMirroringDevicesNotSupported                               = 0xC02625E3
	ErrorGraphicsInvalidPointer                                             = 0xC02625E4
	ErrorGraphicsNoMonitorsCorrespondToDisplayDevice                        = 0xC02625E5
	ErrorGraphicsParameterArrayTooSmall                                     = 0xC02625E6
	ErrorGraphicsInternalError                                              = 0xC02625E7
	ErrorGraphicsSessionTypeChangeInProgress                                = 0xC02605E8
	TpmEErrorMask                                                           = 0x80280000
	TpmEAuthfail                                                            = 0x80280001
	TpmEBadindex                                                            = 0x80280002
	TpmEBadParameter                                                        = 0x80280003
	TpmEAuditfailure                                                        = 0x80280004
	TpmEClearDisabled                                                       = 0x80280005
	TpmEDeactivated                                                         = 0x80280006
	TpmEDisabled                                                            = 0x80280007
	TpmEDisabledCmd                                                         = 0x80280008
	TpmEFail                                                                = 0x80280009
	TpmEBadOrdinal                                                          = 0x8028000A
	TpmEInstallDisabled                                                     = 0x8028000B
	TpmEInvalidKeyhandle                                                    = 0x8028000C
	TpmEKeynotfound                                                         = 0x8028000D
	TpmEInappropriateEnc                                                    = 0x8028000E
	TpmEMigratefail                                                         = 0x8028000F
	TpmEInvalidPcrInfo                                                      = 0x80280010
	TpmENospace                                                             = 0x80280011
	TpmENosrk                                                               = 0x80280012
	TpmENotsealedBlob                                                       = 0x80280013
	TpmEOwnerSet                                                            = 0x80280014
	TpmEResources                                                           = 0x80280015
	TpmEShortrandom                                                         = 0x80280016
	TpmESize                                                                = 0x80280017
	TpmEWrongpcrval                                                         = 0x80280018
	TpmEBadParamSize                                                        = 0x80280019
	TpmEShaThread                                                           = 0x8028001A
	TpmEShaError                                                            = 0x8028001B
	TpmEFailedselftest                                                      = 0x8028001C
	TpmEAuth2Fail                                                           = 0x8028001D
	TpmEBadtag                                                              = 0x8028001E
	TpmEIoerror                                                             = 0x8028001F
	TpmEEncryptError                                                        = 0x80280020
	TpmEDecryptError                                                        = 0x80280021
	TpmEInvalidAuthhandle                                                   = 0x80280022
	TpmENoEndorsement                                                       = 0x80280023
	TpmEInvalidKeyusage                                                     = 0x80280024
	TpmEWrongEntitytype                                                     = 0x80280025
	TpmEInvalidPostinit                                                     = 0x80280026
	TpmEInappropriateSig                                                    = 0x80280027
	TpmEBadKeyProperty                                                      = 0x80280028
	TpmEBadMigration                                                        = 0x80280029
	TpmEBadScheme                                                           = 0x8028002A
	TpmEBadDatasize                                                         = 0x8028002B
	TpmEBadMode                                                             = 0x8028002C
	TpmEBadPresence                                                         = 0x8028002D
	TpmEBadVersion                                                          = 0x8028002E
	TpmENoWrapTransport                                                     = 0x8028002F
	TpmEAuditfailUnsuccessful                                               = 0x80280030
	TpmEAuditfailSuccessful                                                 = 0x80280031
	TpmENotresetable                                                        = 0x80280032
	TpmENotlocal                                                            = 0x80280033
	TpmEBadType                                                             = 0x80280034
	TpmEInvalidResource                                                     = 0x80280035
	TpmENotfips                                                             = 0x80280036
	TpmEInvalidFamily                                                       = 0x80280037
	TpmENoNvPermission                                                      = 0x80280038
	TpmERequiresSign                                                        = 0x80280039
	TpmEKeyNotsupported                                                     = 0x8028003A
	TpmEAuthConflict                                                        = 0x8028003B
	TpmEAreaLocked                                                          = 0x8028003C
	TpmEBadLocality                                                         = 0x8028003D
	TpmEReadOnly                                                            = 0x8028003E
	TpmEPerNowrite                                                          = 0x8028003F
	TpmEFamilycount                                                         = 0x80280040
	TpmEWriteLocked                                                         = 0x80280041
	TpmEBadAttributes                                                       = 0x80280042
	TpmEInvalidStructure                                                    = 0x80280043
	TpmEKeyOwnerControl                                                     = 0x80280044
	TpmEBadCounter                                                          = 0x80280045
	TpmENotFullwrite                                                        = 0x80280046
	TpmEContextGap                                                          = 0x80280047
	TpmEMaxnvwrites                                                         = 0x80280048
	TpmENooperator                                                          = 0x80280049
	TpmEResourcemissing                                                     = 0x8028004A
	TpmEDelegateLock                                                        = 0x8028004B
	TpmEDelegateFamily                                                      = 0x8028004C
	TpmEDelegateAdmin                                                       = 0x8028004D
	TpmETransportNotexclusive                                               = 0x8028004E
	TpmEOwnerControl                                                        = 0x8028004F
	TpmEDaaResources                                                        = 0x80280050
	TpmEDaaInputData0                                                       = 0x80280051
	TpmEDaaInputData1                                                       = 0x80280052
	TpmEDaaIssuerSettings                                                   = 0x80280053
	TpmEDaaTpmSettings                                                      = 0x80280054
	TpmEDaaStage                                                            = 0x80280055
	TpmEDaaIssuerValidity                                                   = 0x80280056
	TpmEDaaWrongW                                                           = 0x80280057
	TpmEBadHandle                                                           = 0x80280058
	TpmEBadDelegate                                                         = 0x80280059
	TpmEBadcontext                                                          = 0x8028005A
	TpmEToomanycontexts                                                     = 0x8028005B
	TpmEMaTicketSignature                                                   = 0x8028005C
	TpmEMaDestination                                                       = 0x8028005D
	TpmEMaSource                                                            = 0x8028005E
	TpmEMaAuthority                                                         = 0x8028005F
	TpmEPermanentek                                                         = 0x80280061
	TpmEBadSignature                                                        = 0x80280062
	TpmENocontextspace                                                      = 0x80280063
	TpmECommandBlocked                                                      = 0x80280400
	TpmEInvalidHandle                                                       = 0x80280401
	TpmEDuplicateVhandle                                                    = 0x80280402
	TpmEEmbeddedCommandBlocked                                              = 0x80280403
	TpmEEmbeddedCommandUnsupported                                          = 0x80280404
	TpmERetry                                                               = 0x80280800
	TpmENeedsSelftest                                                       = 0x80280801
	TpmEDoingSelftest                                                       = 0x80280802
	TpmEDefendLockRunning                                                   = 0x80280803
	TbsEInternalError                                                       = 0x80284001
	TbsEBadParameter                                                        = 0x80284002
	TbsEInvalidOutputPointer                                                = 0x80284003
	TbsEInvalidContext                                                      = 0x80284004
	TbsEInsufficientBuffer                                                  = 0x80284005
	TbsEIoerror                                                             = 0x80284006
	TbsEInvalidContextParam                                                 = 0x80284007
	TbsEServiceNotRunning                                                   = 0x80284008
	TbsETooManyTbsContexts                                                  = 0x80284009
	TbsETooManyResources                                                    = 0x8028400A
	TbsEServiceStartPending                                                 = 0x8028400B
	TbsEPpiNotSupported                                                     = 0x8028400C
	TbsECommandCanceled                                                     = 0x8028400D
	TbsEBufferTooLarge                                                      = 0x8028400E
	TbsETpmNotFound                                                         = 0x8028400F
	TbsEServiceDisabled                                                     = 0x80284010
	TbsENoEventLog                                                          = 0x80284011
	TpmapiEInvalidState                                                     = 0x80290100
	TpmapiENotEnoughData                                                    = 0x80290101
	TpmapiETooMuchData                                                      = 0x80290102
	TpmapiEInvalidOutputPointer                                             = 0x80290103
	TpmapiEInvalidParameter                                                 = 0x80290104
	TpmapiEOutOfMemory                                                      = 0x80290105
	TpmapiEBufferTooSmall                                                   = 0x80290106
	TpmapiEInternalError                                                    = 0x80290107
	TpmapiEAccessDenied                                                     = 0x80290108
	TpmapiEAuthorizationFailed                                              = 0x80290109
	TpmapiEInvalidContextHandle                                             = 0x8029010A
	TpmapiETbsCommunicationError                                            = 0x8029010B
	TpmapiETpmCommandError                                                  = 0x8029010C
	TpmapiEMessageTooLarge                                                  = 0x8029010D
	TpmapiEInvalidEncoding                                                  = 0x8029010E
	TpmapiEInvalidKeySize                                                   = 0x8029010F
	TpmapiEEncryptionFailed                                                 = 0x80290110
	TpmapiEInvalidKeyParams                                                 = 0x80290111
	TpmapiEInvalidMigrationAuthorizationBlob                                = 0x80290112
	TpmapiEInvalidPcrIndex                                                  = 0x80290113
	TpmapiEInvalidDelegateBlob                                              = 0x80290114
	TpmapiEInvalidContextParams                                             = 0x80290115
	TpmapiEInvalidKeyBlob                                                   = 0x80290116
	TpmapiEInvalidPcrData                                                   = 0x80290117
	TpmapiEInvalidOwnerAuth                                                 = 0x80290118
	TpmapiEFipsRngCheckFailed                                               = 0x80290119
	TpmapiEEmptyTcgLog                                                      = 0x8029011A
	TpmapiEInvalidTcgLogEntry                                               = 0x8029011B
	TpmapiETcgSeparatorAbsent                                               = 0x8029011C
	TpmapiETcgInvalidDigestEntry                                            = 0x8029011D
	TbsimpEBufferTooSmall                                                   = 0x80290200
	TbsimpECleanupFailed                                                    = 0x80290201
	TbsimpEInvalidContextHandle                                             = 0x80290202
	TbsimpEInvalidContextParam                                              = 0x80290203
	TbsimpETpmError                                                         = 0x80290204
	TbsimpEHashBadKey                                                       = 0x80290205
	TbsimpEDuplicateVhandle                                                 = 0x80290206
	TbsimpEInvalidOutputPointer                                             = 0x80290207
	TbsimpEInvalidParameter                                                 = 0x80290208
	TbsimpERpcInitFailed                                                    = 0x80290209
	TbsimpESchedulerNotRunning                                              = 0x8029020A
	TbsimpECommandCanceled                                                  = 0x8029020B
	TbsimpEOutOfMemory                                                      = 0x8029020C
	TbsimpEListNoMoreItems                                                  = 0x8029020D
	TbsimpEListNotFound                                                     = 0x8029020E
	TbsimpENotEnoughSpace                                                   = 0x8029020F
	TbsimpENotEnoughTpmContexts                                             = 0x80290210
	TbsimpECommandFailed                                                    = 0x80290211
	TbsimpEUnknownOrdinal                                                   = 0x80290212
	TbsimpEResourceExpired                                                  = 0x80290213
	TbsimpEInvalidResource                                                  = 0x80290214
	TbsimpENothingToUnload                                                  = 0x80290215
	TbsimpEHashTableFull                                                    = 0x80290216
	TbsimpETooManyTbsContexts                                               = 0x80290217
	TbsimpETooManyResources                                                 = 0x80290218
	TbsimpEPpiNotSupported                                                  = 0x80290219
	TbsimpETpmIncompatible                                                  = 0x8029021A
	TbsimpENoEventLog                                                       = 0x8029021B
	TpmEPpiAcpiFailure                                                      = 0x80290300
	TpmEPpiUserAbort                                                        = 0x80290301
	TpmEPpiBiosFailure                                                      = 0x80290302
	TpmEPpiNotSupported                                                     = 0x80290303
	PlaEDcsNotFound                                                         = 0x80300002
	PlaEDcsInUse                                                            = 0x803000AA
	PlaETooManyFolders                                                      = 0x80300045
	PlaENoMinDisk                                                           = 0x80300070
	PlaEDcsAlreadyExists                                                    = 0x803000B7
	PlaSPropertyIgnored                                                     = 0x00300100
	PlaEPropertyConflict                                                    = 0x80300101
	PlaEDcsSingletonRequired                                                = 0x80300102
	PlaECredentialsRequired                                                 = 0x80300103
	PlaEDcsNotRunning                                                       = 0x80300104
	PlaEConflictInclExclApi                                                 = 0x80300105
	PlaENetworkExeNotValid                                                  = 0x80300106
	PlaEExeAlreadyConfigured                                                = 0x80300107
	PlaEExePathNotValid                                                     = 0x80300108
	PlaEDcAlreadyExists                                                     = 0x80300109
	PlaEDcsStartWaitTimeout                                                 = 0x8030010A
	PlaEDcStartWaitTimeout                                                  = 0x8030010B
	PlaEReportWaitTimeout                                                   = 0x8030010C
	PlaENoDuplicates                                                        = 0x8030010D
	PlaEExeFullPathRequired                                                 = 0x8030010E
	PlaEInvalidSessionName                                                  = 0x8030010F
	PlaEPlaChannelNotEnabled                                                = 0x80300110
	PlaETaskschedChannelNotEnabled                                          = 0x80300111
	PlaERulesManagerFailed                                                  = 0x80300112
	PlaECabapiFailure                                                       = 0x80300113
	FveELockedVolume                                                        = 0x80310000
	FveENotEncrypted                                                        = 0x80310001
	FveENoTpmBios                                                           = 0x80310002
	FveENoMbrMetric                                                         = 0x80310003
	FveENoBootsectorMetric                                                  = 0x80310004
	FveENoBootmgrMetric                                                     = 0x80310005
	FveEWrongBootmgr                                                        = 0x80310006
	FveESecureKeyRequired                                                   = 0x80310007
	FveENotActivated                                                        = 0x80310008
	FveEActionNotAllowed                                                    = 0x80310009
	FveEAdSchemaNotInstalled                                                = 0x8031000A
	FveEAdInvalidDatatype                                                   = 0x8031000B
	FveEAdInvalidDatasize                                                   = 0x8031000C
	FveEAdNoValues                                                          = 0x8031000D
	FveEAdAttrNotSet                                                        = 0x8031000E
	FveEAdGuidNotFound                                                      = 0x8031000F
	FveEBadInformation                                                      = 0x80310010
	FveETooSmall                                                            = 0x80310011
	FveESystemVolume                                                        = 0x80310012
	FveEFailedWrongFs                                                       = 0x80310013
	FveEBadPartitionSize                                                    = 0x80310014
	FveENotSupported                                                        = 0x80310015
	FveEBadData                                                             = 0x80310016
	FveEVolumeNotBound                                                      = 0x80310017
	FveETpmNotOwned                                                         = 0x80310018
	FveENotDataVolume                                                       = 0x80310019
	FveEAdInsufficientBuffer                                                = 0x8031001A
	FveEConvRead                                                            = 0x8031001B
	FveEConvWrite                                                           = 0x8031001C
	FveEKeyRequired                                                         = 0x8031001D
	FveEClusteringNotSupported                                              = 0x8031001E
	FveEVolumeBoundAlready                                                  = 0x8031001F
	FveEOsNotProtected                                                      = 0x80310020
	FveEProtectionDisabled                                                  = 0x80310021
	FveERecoveryKeyRequired                                                 = 0x80310022
	FveEForeignVolume                                                       = 0x80310023
	FveEOverlappedUpdate                                                    = 0x80310024
	FveETpmSrkAuthNotZero                                                   = 0x80310025
	FveEFailedSectorSize                                                    = 0x80310026
	FveEFailedAuthentication                                                = 0x80310027
	FveENotOsVolume                                                         = 0x80310028
	FveEAutounlockEnabled                                                   = 0x80310029
	FveEWrongBootsector                                                     = 0x8031002A
	FveEWrongSystemFs                                                       = 0x8031002B
	FveEPolicyPasswordRequired                                              = 0x8031002C
	FveECannotSetFvekEncrypted                                              = 0x8031002D
	FveECannotEncryptNoKey                                                  = 0x8031002E
	FveEBootableCddvd                                                       = 0x80310030
	FveEProtectorExists                                                     = 0x80310031
	FveERelativePath                                                        = 0x80310032
	FveEProtectorNotFound                                                   = 0x80310033
	FveEInvalidKeyFormat                                                    = 0x80310034
	FveEInvalidPasswordFormat                                               = 0x80310035
	FveEFipsRngCheckFailed                                                  = 0x80310036
	FveEFipsPreventsRecoveryPassword                                        = 0x80310037
	FveEFipsPreventsExternalKeyExport                                       = 0x80310038
	FveENotDecrypted                                                        = 0x80310039
	FveEInvalidProtectorType                                                = 0x8031003A
	FveENoProtectorsToTest                                                  = 0x8031003B
	FveEKeyfileNotFound                                                     = 0x8031003C
	FveEKeyfileInvalid                                                      = 0x8031003D
	FveEKeyfileNoVmk                                                        = 0x8031003E
	FveETpmDisabled                                                         = 0x8031003F
	FveENotAllowedInSafeMode                                                = 0x80310040
	FveETpmInvalidPcr                                                       = 0x80310041
	FveETpmNoVmk                                                            = 0x80310042
	FveEPinInvalid                                                          = 0x80310043
	FveEAuthInvalidApplication                                              = 0x80310044
	FveEAuthInvalidConfig                                                   = 0x80310045
	FveEFipsDisableProtectionNotAllowed                                     = 0x80310046
	FveEFsNotExtended                                                       = 0x80310047
	FveEFirmwareTypeNotSupported                                            = 0x80310048
	FveENoLicense                                                           = 0x80310049
	FveENotOnStack                                                          = 0x8031004A
	FveEFsMounted                                                           = 0x8031004B
	FveETokenNotImpersonated                                                = 0x8031004C
	FveEDryRunFailed                                                        = 0x8031004D
	FveERebootRequired                                                      = 0x8031004E
	FveEDebuggerEnabled                                                     = 0x8031004F
	FveERawAccess                                                           = 0x80310050
	FveERawBlocked                                                          = 0x80310051
	FveEBcdApplicationsPathIncorrect                                        = 0x80310052
	FveENotAllowedInVersion                                                 = 0x80310053
	FveENoAutounlockMasterKey                                               = 0x80310054
	FveEMorFailed                                                           = 0x80310055
	FveEHiddenVolume                                                        = 0x80310056
	FveETransientState                                                      = 0x80310057
	FveEPubkeyNotAllowed                                                    = 0x80310058
	FveEVolumeHandleOpen                                                    = 0x80310059
	FveENoFeatureLicense                                                    = 0x8031005A
	FveEInvalidStartupOptions                                               = 0x8031005B
	FveEPolicyRecoveryPasswordNotAllowed                                    = 0x8031005C
	FveEPolicyRecoveryPasswordRequired                                      = 0x8031005D
	FveEPolicyRecoveryKeyNotAllowed                                         = 0x8031005E
	FveEPolicyRecoveryKeyRequired                                           = 0x8031005F
	FveEPolicyStartupPinNotAllowed                                          = 0x80310060
	FveEPolicyStartupPinRequired                                            = 0x80310061
	FveEPolicyStartupKeyNotAllowed                                          = 0x80310062
	FveEPolicyStartupKeyRequired                                            = 0x80310063
	FveEPolicyStartupPinKeyNotAllowed                                       = 0x80310064
	FveEPolicyStartupPinKeyRequired                                         = 0x80310065
	FveEPolicyStartupTpmNotAllowed                                          = 0x80310066
	FveEPolicyStartupTpmRequired                                            = 0x80310067
	FveEPolicyInvalidPinLength                                              = 0x80310068
	FveEKeyProtectorNotSupported                                            = 0x80310069
	FveEPolicyPassphraseNotAllowed                                          = 0x8031006A
	FveEPolicyPassphraseRequired                                            = 0x8031006B
	FveEFipsPreventsPassphrase                                              = 0x8031006C
	FveEOsVolumePassphraseNotAllowed                                        = 0x8031006D
	FveEInvalidBitlockerOid                                                 = 0x8031006E
	FveEVolumeTooSmall                                                      = 0x8031006F
	FveEDvNotSupportedOnFs                                                  = 0x80310070
	FveEDvNotAllowedByGp                                                    = 0x80310071
	FveEPolicyUserCertificateNotAllowed                                     = 0x80310072
	FveEPolicyUserCertificateRequired                                       = 0x80310073
	FveEPolicyUserCertMustBeHw                                              = 0x80310074
	FveEPolicyUserConfigureFdvAutounlockNotAllowed                          = 0x80310075
	FveEPolicyUserConfigureRdvAutounlockNotAllowed                          = 0x80310076
	FveEPolicyUserConfigureRdvNotAllowed                                    = 0x80310077
	FveEPolicyUserEnableRdvNotAllowed                                       = 0x80310078
	FveEPolicyUserDisableRdvNotAllowed                                      = 0x80310079
	FveEPolicyInvalidPassphraseLength                                       = 0x80310080
	FveEPolicyPassphraseTooSimple                                           = 0x80310081
	FveERecoveryPartition                                                   = 0x80310082
	FveEPolicyConflictFdvRkOffAukOn                                         = 0x80310083
	FveEPolicyConflictRdvRkOffAukOn                                         = 0x80310084
	FveENonBitlockerOid                                                     = 0x80310085
	FveEPolicyProhibitsSelfsigned                                           = 0x80310086
	FveEPolicyConflictRoAndStartupKeyRequired                               = 0x80310087
	FveEConvRecoveryFailed                                                  = 0x80310088
	FveEVirtualizedSpaceTooBig                                              = 0x80310089
	FveEPolicyConflictOsvRpOffAdbOn                                         = 0x80310090
	FveEPolicyConflictFdvRpOffAdbOn                                         = 0x80310091
	FveEPolicyConflictRdvRpOffAdbOn                                         = 0x80310092
	FveENonBitlockerKu                                                      = 0x80310093
	FveEPrivatekeyAuthFailed                                                = 0x80310094
	FveERemovalOfDraFailed                                                  = 0x80310095
	FveEOperationNotSupportedOnVistaVolume                                  = 0x80310096
	FveECantLockAutounlockEnabledVolume                                     = 0x80310097
	FveEFipsHashKdfNotAllowed                                               = 0x80310098
	FveEEnhPinInvalid                                                       = 0x80310099
	FveEInvalidPinChars                                                     = 0x8031009A
	FveEInvalidDatumType                                                    = 0x8031009B
	FwpECalloutNotFound                                                     = 0x80320001
	FwpEConditionNotFound                                                   = 0x80320002
	FwpEFilterNotFound                                                      = 0x80320003
	FwpELayerNotFound                                                       = 0x80320004
	FwpEProviderNotFound                                                    = 0x80320005
	FwpEProviderContextNotFound                                             = 0x80320006
	FwpESublayerNotFound                                                    = 0x80320007
	FwpENotFound                                                            = 0x80320008
	FwpEAlreadyExists                                                       = 0x80320009
	FwpEInUse                                                               = 0x8032000A
	FwpEDynamicSessionInProgress                                            = 0x8032000B
	FwpEWrongSession                                                        = 0x8032000C
	FwpENoTxnInProgress                                                     = 0x8032000D
	FwpETxnInProgress                                                       = 0x8032000E
	FwpETxnAborted                                                          = 0x8032000F
	FwpESessionAborted                                                      = 0x80320010
	FwpEIncompatibleTxn                                                     = 0x80320011
	FwpETimeout                                                             = 0x80320012
	FwpENetEventsDisabled                                                   = 0x80320013
	FwpEIncompatibleLayer                                                   = 0x80320014
	FwpEKmClientsOnly                                                       = 0x80320015
	FwpELifetimeMismatch                                                    = 0x80320016
	FwpEBuiltinObject                                                       = 0x80320017
	FwpETooManyCallouts                                                     = 0x80320018
	FwpENotificationDropped                                                 = 0x80320019
	FwpETrafficMismatch                                                     = 0x8032001A
	FwpEIncompatibleSaState                                                 = 0x8032001B
	FwpENullPointer                                                         = 0x8032001C
	FwpEInvalidEnumerator                                                   = 0x8032001D
	FwpEInvalidFlags                                                        = 0x8032001E
	FwpEInvalidNetMask                                                      = 0x8032001F
	FwpEInvalidRange                                                        = 0x80320020
	FwpEInvalidInterval                                                     = 0x80320021
	FwpEZeroLengthArray                                                     = 0x80320022
	FwpENullDisplayName                                                     = 0x80320023
	FwpEInvalidActionType                                                   = 0x80320024
	FwpEInvalidWeight                                                       = 0x80320025
	FwpEMatchTypeMismatch                                                   = 0x80320026
	FwpETypeMismatch                                                        = 0x80320027
	FwpEOutOfBounds                                                         = 0x80320028
	FwpEReserved                                                            = 0x80320029
	FwpEDuplicateCondition                                                  = 0x8032002A
	FwpEDuplicateKeymod                                                     = 0x8032002B
	FwpEActionIncompatibleWithLayer                                         = 0x8032002C
	FwpEActionIncompatibleWithSublayer                                      = 0x8032002D
	FwpEContextIncompatibleWithLayer                                        = 0x8032002E
	FwpEContextIncompatibleWithCallout                                      = 0x8032002F
	FwpEIncompatibleAuthMethod                                              = 0x80320030
	FwpEIncompatibleDhGroup                                                 = 0x80320031
	FwpEEmNotSupported                                                      = 0x80320032
	FwpENeverMatch                                                          = 0x80320033
	FwpEProviderContextMismatch                                             = 0x80320034
	FwpEInvalidParameter                                                    = 0x80320035
	FwpETooManySublayers                                                    = 0x80320036
	FwpECalloutNotificationFailed                                           = 0x80320037
	FwpEInvalidAuthTransform                                                = 0x80320038
	FwpEInvalidCipherTransform                                              = 0x80320039
	FwpEDropNoicmp                                                          = 0x80320104
	FwpEIncompatibleCipherTransform                                         = 0x8032003A
	FwpEInvalidTransformCombination                                         = 0x8032003B
	FwpEDuplicateAuthMethod                                                 = 0x8032003C
	WsSAsync                                                                = 0x003D0000
	WsSEnd                                                                  = 0x003D0001
	WsEInvalidFormat                                                        = 0x803D0000
	WsEObjectFaulted                                                        = 0x803D0001
	WsENumericOverflow                                                      = 0x803D0002
	WsEInvalidOperation                                                     = 0x803D0003
	WsEOperationAborted                                                     = 0x803D0004
	WsEEndpointAccessDenied                                                 = 0x803D0005
	WsEOperationTimedOut                                                    = 0x803D0006
	WsEOperationAbandoned                                                   = 0x803D0007
	WsEQuotaExceeded                                                        = 0x803D0008
	WsENoTranslationAvailable                                               = 0x803D0009
	WsESecurityVerificationFailure                                          = 0x803D000A
	WsEAddressInUse                                                         = 0x803D000B
	WsEAddressNotAvailable                                                  = 0x803D000C
	WsEEndpointNotFound                                                     = 0x803D000D
	WsEEndpointNotAvailable                                                 = 0x803D000E
	WsEEndpointFailure                                                      = 0x803D000F
	WsEEndpointUnreachable                                                  = 0x803D0010
	WsEEndpointActionNotSupported                                           = 0x803D0011
	WsEEndpointTooBusy                                                      = 0x803D0012
	WsEEndpointFaultReceived                                                = 0x803D0013
	WsEEndpointDisconnected                                                 = 0x803D0014
	WsEProxyFailure                                                         = 0x803D0015
	WsEProxyAccessDenied                                                    = 0x803D0016
	WsENotSupported                                                         = 0x803D0017
	WsEProxyRequiresBasicAuth                                               = 0x803D0018
	WsEProxyRequiresDigestAuth                                              = 0x803D0019
	WsEProxyRequiresNtlmAuth                                                = 0x803D001A
	WsEProxyRequiresNegotiateAuth                                           = 0x803D001B
	WsEServerRequiresBasicAuth                                              = 0x803D001C
	WsEServerRequiresDigestAuth                                             = 0x803D001D
	WsEServerRequiresNtlmAuth                                               = 0x803D001E
	WsEServerRequiresNegotiateAuth                                          = 0x803D001F
	WsEInvalidEndpointUrl                                                   = 0x803D0020
	WsEOther                                                                = 0x803D0021
	WsESecurityTokenExpired                                                 = 0x803D0022
	WsESecuritySystemFailure                                                = 0x803D0023
	ErrorNdisInterfaceClosing                                               = 0x80340002
	ErrorNdisBadVersion                                                     = 0x80340004
	ErrorNdisBadCharacteristics                                             = 0x80340005
	ErrorNdisAdapterNotFound                                                = 0x80340006
	ErrorNdisOpenFailed                                                     = 0x80340007
	ErrorNdisDeviceFailed                                                   = 0x80340008
	ErrorNdisMulticastFull                                                  = 0x80340009
	ErrorNdisMulticastExists                                                = 0x8034000A
	ErrorNdisMulticastNotFound                                              = 0x8034000B
	ErrorNdisRequestAborted                                                 = 0x8034000C
	ErrorNdisResetInProgress                                                = 0x8034000D
	ErrorNdisNotSupported                                                   = 0x803400BB
	ErrorNdisInvalidPacket                                                  = 0x8034000F
	ErrorNdisAdapterNotReady                                                = 0x80340011
	ErrorNdisInvalidLength                                                  = 0x80340014
	ErrorNdisInvalidData                                                    = 0x80340015
	ErrorNdisBufferTooShort                                                 = 0x80340016
	ErrorNdisInvalidOid                                                     = 0x80340017
	ErrorNdisAdapterRemoved                                                 = 0x80340018
	ErrorNdisUnsupportedMedia                                               = 0x80340019
	ErrorNdisGroupAddressInUse                                              = 0x8034001A
	ErrorNdisFileNotFound                                                   = 0x8034001B
	ErrorNdisErrorReadingFile                                               = 0x8034001C
	ErrorNdisAlreadyMapped                                                  = 0x8034001D
	ErrorNdisResourceConflict                                               = 0x8034001E
	ErrorNdisMediaDisconnected                                              = 0x8034001F
	ErrorNdisInvalidAddress                                                 = 0x80340022
	ErrorNdisInvalidDeviceRequest                                           = 0x80340010
	ErrorNdisPaused                                                         = 0x8034002A
	ErrorNdisInterfaceNotFound                                              = 0x8034002B
	ErrorNdisUnsupportedRevision                                            = 0x8034002C
	ErrorNdisInvalidPort                                                    = 0x8034002D
	ErrorNdisInvalidPortState                                               = 0x8034002E
	ErrorNdisLowPowerState                                                  = 0x8034002F
	ErrorNdisDot11AutoConfigEnabled                                         = 0x80342000
	ErrorNdisDot11MediaInUse                                                = 0x80342001
	ErrorNdisDot11PowerStateInvalid                                         = 0x80342002
	ErrorNdisPmWolPatternListFull                                           = 0x80342003
	ErrorNdisPmProtocolOffloadListFull                                      = 0x80342004
	ErrorNdisIndicationRequired                                             = 0x00340001
	ErrorNdisOffloadPolicy                                                  = 0xC034100F
	ErrorNdisOffloadConnectionRejected                                      = 0xC0341012
	ErrorNdisOffloadPathRejected                                            = 0xC0341013
	ErrorHvInvalidHypercallCode                                             = 0xC0350002
	ErrorHvInvalidHypercallInput                                            = 0xC0350003
	ErrorHvInvalidAlignment                                                 = 0xC0350004
	ErrorHvInvalidParameter                                                 = 0xC0350005
	ErrorHvAccessDenied                                                     = 0xC0350006
	ErrorHvInvalidPartitionState                                            = 0xC0350007
	ErrorHvOperationDenied                                                  = 0xC0350008
	ErrorHvUnknownProperty                                                  = 0xC0350009
	ErrorHvPropertyValueOutOfRange                                          = 0xC035000A
	ErrorHvInsufficientMemory                                               = 0xC035000B
	ErrorHvPartitionTooDeep                                                 = 0xC035000C
	ErrorHvInvalidPartitionId                                               = 0xC035000D
	ErrorHvInvalidVpIndex                                                   = 0xC035000E
	ErrorHvInvalidPortId                                                    = 0xC0350011
	ErrorHvInvalidConnectionId                                              = 0xC0350012
	ErrorHvInsufficientBuffers                                              = 0xC0350013
	ErrorHvNotAcknowledged                                                  = 0xC0350014
	ErrorHvAcknowledged                                                     = 0xC0350016
	ErrorHvInvalidSaveRestoreState                                          = 0xC0350017
	ErrorHvInvalidSynicState                                                = 0xC0350018
	ErrorHvObjectInUse                                                      = 0xC0350019
	ErrorHvInvalidProximityDomainInfo                                       = 0xC035001A
	ErrorHvNoData                                                           = 0xC035001B
	ErrorHvInactive                                                         = 0xC035001C
	ErrorHvNoResources                                                      = 0xC035001D
	ErrorHvFeatureUnavailable                                               = 0xC035001E
	ErrorHvNotPresent                                                       = 0xC0351000
	ErrorVidDuplicateHandler                                                = 0xC0370001
	ErrorVidTooManyHandlers                                                 = 0xC0370002
	ErrorVidQueueFull                                                       = 0xC0370003
	ErrorVidHandlerNotPresent                                               = 0xC0370004
	ErrorVidInvalidObjectName                                               = 0xC0370005
	ErrorVidPartitionNameTooLong                                            = 0xC0370006
	ErrorVidMessageQueueNameTooLong                                         = 0xC0370007
	ErrorVidPartitionAlreadyExists                                          = 0xC0370008
	ErrorVidPartitionDoesNotExist                                           = 0xC0370009
	ErrorVidPartitionNameNotFound                                           = 0xC037000A
	ErrorVidMessageQueueAlreadyExists                                       = 0xC037000B
	ErrorVidExceededMbpEntryMapLimit                                        = 0xC037000C
	ErrorVidMbStillReferenced                                               = 0xC037000D
	ErrorVidChildGpaPageSetCorrupted                                        = 0xC037000E
	ErrorVidInvalidNumaSettings                                             = 0xC037000F
	ErrorVidInvalidNumaNodeIndex                                            = 0xC0370010
	ErrorVidNotificationQueueAlreadyAssociated                              = 0xC0370011
	ErrorVidInvalidMemoryBlockHandle                                        = 0xC0370012
	ErrorVidPageRangeOverflow                                               = 0xC0370013
	ErrorVidInvalidMessageQueueHandle                                       = 0xC0370014
	ErrorVidInvalidGpaRangeHandle                                           = 0xC0370015
	ErrorVidNoMemoryBlockNotificationQueue                                  = 0xC0370016
	ErrorVidMemoryBlockLockCountExceeded                                    = 0xC0370017
	ErrorVidInvalidPpmHandle                                                = 0xC0370018
	ErrorVidMbpsAreLocked                                                   = 0xC0370019
	ErrorVidMessageQueueClosed                                              = 0xC037001A
	ErrorVidVirtualProcessorLimitExceeded                                   = 0xC037001B
	ErrorVidStopPending                                                     = 0xC037001C
	ErrorVidInvalidProcessorState                                           = 0xC037001D
	ErrorVidExceededKmContextCountLimit                                     = 0xC037001E
	ErrorVidKmInterfaceAlreadyInitialized                                   = 0xC037001F
	ErrorVidMbPropertyAlreadySetReset                                       = 0xC0370020
	ErrorVidMmioRangeDestroyed                                              = 0xC0370021
	ErrorVidInvalidChildGpaPageSet                                          = 0xC0370022
	ErrorVidReservePageSetIsBeingUsed                                       = 0xC0370023
	ErrorVidReservePageSetTooSmall                                          = 0xC0370024
	ErrorVidMbpAlreadyLockedUsingReservedPage                               = 0xC0370025
	ErrorVidMbpCountExceededLimit                                           = 0xC0370026
	ErrorVidSavedStateCorrupt                                               = 0xC0370027
	ErrorVidSavedStateUnrecognizedItem                                      = 0xC0370028
	ErrorVidSavedStateIncompatible                                          = 0xC0370029
	ErrorVidRemoteNodeParentGpaPagesUsed                                    = 0x80370001
	ErrorVolmgrIncompleteRegeneration                                       = 0x80380001
	ErrorVolmgrIncompleteDiskMigration                                      = 0x80380002
	ErrorVolmgrDatabaseFull                                                 = 0xC0380001
	ErrorVolmgrDiskConfigurationCorrupted                                   = 0xC0380002
	ErrorVolmgrDiskConfigurationNotInSync                                   = 0xC0380003
	ErrorVolmgrPackConfigUpdateFailed                                       = 0xC0380004
	ErrorVolmgrDiskContainsNonSimpleVolume                                  = 0xC0380005
	ErrorVolmgrDiskDuplicate                                                = 0xC0380006
	ErrorVolmgrDiskDynamic                                                  = 0xC0380007
	ErrorVolmgrDiskIdInvalid                                                = 0xC0380008
	ErrorVolmgrDiskInvalid                                                  = 0xC0380009
	ErrorVolmgrDiskLastVoter                                                = 0xC038000A
	ErrorVolmgrDiskLayoutInvalid                                            = 0xC038000B
	ErrorVolmgrDiskLayoutNonBasicBetweenBasicPartitions                     = 0xC038000C
	ErrorVolmgrDiskLayoutNotCylinderAligned                                 = 0xC038000D
	ErrorVolmgrDiskLayoutPartitionsTooSmall                                 = 0xC038000E
	ErrorVolmgrDiskLayoutPrimaryBetweenLogicalPartitions                    = 0xC038000F
	ErrorVolmgrDiskLayoutTooManyPartitions                                  = 0xC0380010
	ErrorVolmgrDiskMissing                                                  = 0xC0380011
	ErrorVolmgrDiskNotEmpty                                                 = 0xC0380012
	ErrorVolmgrDiskNotEnoughSpace                                           = 0xC0380013
	ErrorVolmgrDiskRevectoringFailed                                        = 0xC0380014
	ErrorVolmgrDiskSectorSizeInvalid                                        = 0xC0380015
	ErrorVolmgrDiskSetNotContained                                          = 0xC0380016
	ErrorVolmgrDiskUsedByMultipleMembers                                    = 0xC0380017
	ErrorVolmgrDiskUsedByMultiplePlexes                                     = 0xC0380018
	ErrorVolmgrDynamicDiskNotSupported                                      = 0xC0380019
	ErrorVolmgrExtentAlreadyUsed                                            = 0xC038001A
	ErrorVolmgrExtentNotContiguous                                          = 0xC038001B
	ErrorVolmgrExtentNotInPublicRegion                                      = 0xC038001C
	ErrorVolmgrExtentNotSectorAligned                                       = 0xC038001D
	ErrorVolmgrExtentOverlapsEbrPartition                                   = 0xC038001E
	ErrorVolmgrExtentVolumeLengthsDoNotMatch                                = 0xC038001F
	ErrorVolmgrFaultTolerantNotSupported                                    = 0xC0380020
	ErrorVolmgrInterleaveLengthInvalid                                      = 0xC0380021
	ErrorVolmgrMaximumRegisteredUsers                                       = 0xC0380022
	ErrorVolmgrMemberInSync                                                 = 0xC0380023
	ErrorVolmgrMemberIndexDuplicate                                         = 0xC0380024
	ErrorVolmgrMemberIndexInvalid                                           = 0xC0380025
	ErrorVolmgrMemberMissing                                                = 0xC0380026
	ErrorVolmgrMemberNotDetached                                            = 0xC0380027
	ErrorVolmgrMemberRegenerating                                           = 0xC0380028
	ErrorVolmgrAllDisksFailed                                               = 0xC0380029
	ErrorVolmgrNoRegisteredUsers                                            = 0xC038002A
	ErrorVolmgrNoSuchUser                                                   = 0xC038002B
	ErrorVolmgrNotificationReset                                            = 0xC038002C
	ErrorVolmgrNumberOfMembersInvalid                                       = 0xC038002D
	ErrorVolmgrNumberOfPlexesInvalid                                        = 0xC038002E
	ErrorVolmgrPackDuplicate                                                = 0xC038002F
	ErrorVolmgrPackIdInvalid                                                = 0xC0380030
	ErrorVolmgrPackInvalid                                                  = 0xC0380031
	ErrorVolmgrPackNameInvalid                                              = 0xC0380032
	ErrorVolmgrPackOffline                                                  = 0xC0380033
	ErrorVolmgrPackHasQuorum                                                = 0xC0380034
	ErrorVolmgrPackWithoutQuorum                                            = 0xC0380035
	ErrorVolmgrPartitionStyleInvalid                                        = 0xC0380036
	ErrorVolmgrPartitionUpdateFailed                                        = 0xC0380037
	ErrorVolmgrPlexInSync                                                   = 0xC0380038
	ErrorVolmgrPlexIndexDuplicate                                           = 0xC0380039
	ErrorVolmgrPlexIndexInvalid                                             = 0xC038003A
	ErrorVolmgrPlexLastActive                                               = 0xC038003B
	ErrorVolmgrPlexMissing                                                  = 0xC038003C
	ErrorVolmgrPlexRegenerating                                             = 0xC038003D
	ErrorVolmgrPlexTypeInvalid                                              = 0xC038003E
	ErrorVolmgrPlexNotRaid5                                                 = 0xC038003F
	ErrorVolmgrPlexNotSimple                                                = 0xC0380040
	ErrorVolmgrStructureSizeInvalid                                         = 0xC0380041
	ErrorVolmgrTooManyNotificationRequests                                  = 0xC0380042
	ErrorVolmgrTransactionInProgress                                        = 0xC0380043
	ErrorVolmgrUnexpectedDiskLayoutChange                                   = 0xC0380044
	ErrorVolmgrVolumeContainsMissingDisk                                    = 0xC0380045
	ErrorVolmgrVolumeIdInvalid                                              = 0xC0380046
	ErrorVolmgrVolumeLengthInvalid                                          = 0xC0380047
	ErrorVolmgrVolumeLengthNotSectorSizeMultiple                            = 0xC0380048
	ErrorVolmgrVolumeNotMirrored                                            = 0xC0380049
	ErrorVolmgrVolumeNotRetained                                            = 0xC038004A
	ErrorVolmgrVolumeOffline                                                = 0xC038004B
	ErrorVolmgrVolumeRetained                                               = 0xC038004C
	ErrorVolmgrNumberOfExtentsInvalid                                       = 0xC038004D
	ErrorVolmgrDifferentSectorSize                                          = 0xC038004E
	ErrorVolmgrBadBootDisk                                                  = 0xC038004F
	ErrorVolmgrPackConfigOffline                                            = 0xC0380050
	ErrorVolmgrPackConfigOnline                                             = 0xC0380051
	ErrorVolmgrNotPrimaryPack                                               = 0xC0380052
	ErrorVolmgrPackLogUpdateFailed                                          = 0xC0380053
	ErrorVolmgrNumberOfDisksInPlexInvalid                                   = 0xC0380054
	ErrorVolmgrNumberOfDisksInMemberInvalid                                 = 0xC0380055
	ErrorVolmgrVolumeMirrored                                               = 0xC0380056
	ErrorVolmgrPlexNotSimpleSpanned                                         = 0xC0380057
	ErrorVolmgrNoValidLogCopies                                             = 0xC0380058
	ErrorVolmgrPrimaryPackPresent                                           = 0xC0380059
	ErrorVolmgrNumberOfDisksInvalid                                         = 0xC038005A
	ErrorVolmgrMirrorNotSupported                                           = 0xC038005B
	ErrorVolmgrRaid5NotSupported                                            = 0xC038005C
	ErrorBcdNotAllEntriesImported                                           = 0x80390001
	ErrorBcdTooManyElements                                                 = 0xC0390002
	ErrorBcdNotAllEntriesSynchronized                                       = 0x80390003
	ErrorVhdDriveFooterMissing                                              = 0xC03A0001
	ErrorVhdDriveFooterChecksumMismatch                                     = 0xC03A0002
	ErrorVhdDriveFooterCorrupt                                              = 0xC03A0003
	ErrorVhdFormatUnknown                                                   = 0xC03A0004
	ErrorVhdFormatUnsupportedVersion                                        = 0xC03A0005
	ErrorVhdSparseHeaderChecksumMismatch                                    = 0xC03A0006
	ErrorVhdSparseHeaderUnsupportedVersion                                  = 0xC03A0007
	ErrorVhdSparseHeaderCorrupt                                             = 0xC03A0008
	ErrorVhdBlockAllocationFailure                                          = 0xC03A0009
	ErrorVhdBlockAllocationTableCorrupt                                     = 0xC03A000A
	ErrorVhdInvalidBlockSize                                                = 0xC03A000B
	ErrorVhdBitmapMismatch                                                  = 0xC03A000C
	ErrorVhdParentVhdNotFound                                               = 0xC03A000D
	ErrorVhdChildParentIdMismatch                                           = 0xC03A000E
	ErrorVhdChildParentTimestampMismatch                                    = 0xC03A000F
	ErrorVhdMetadataReadFailure                                             = 0xC03A0010
	ErrorVhdMetadataWriteFailure                                            = 0xC03A0011
	ErrorVhdInvalidSize                                                     = 0xC03A0012
	ErrorVhdInvalidFileSize                                                 = 0xC03A0013
	ErrorVirtdiskProviderNotFound                                           = 0xC03A0014
	ErrorVirtdiskNotVirtualDisk                                             = 0xC03A0015
	ErrorVhdParentVhdAccessDenied                                           = 0xC03A0016
	ErrorVhdChildParentSizeMismatch                                         = 0xC03A0017
	ErrorVhdDifferencingChainCycleDetected                                  = 0xC03A0018
	ErrorVhdDifferencingChainErrorInParent                                  = 0xC03A0019
	ErrorVirtualDiskLimitation                                              = 0xC03A001A
	ErrorVhdInvalidType                                                     = 0xC03A001B
	ErrorVhdInvalidState                                                    = 0xC03A001C
	ErrorVirtdiskUnsupportedDiskSectorSize                                  = 0xC03A001D
	ErrorQueryStorageError                                                  = 0x803A0001
	SdiagECancelled                                                         = 0x803C0100
	SdiagEScript                                                            = 0x803C0101
	SdiagEPowershell                                                        = 0x803C0102
	SdiagEManagedhost                                                       = 0x803C0103
	SdiagENoverifier                                                        = 0x803C0104
	SdiagSCannotrun                                                         = 0x003C0105
	SdiagEDisabled                                                          = 0x803C0106
	SdiagETrust                                                             = 0x803C0107
	SdiagECannotrun                                                         = 0x803C0108
	SdiagEVersion                                                           = 0x803C0109
	SdiagEResource                                                          = 0x803C010A
	SdiagERootcause                                                         = 0x803C010B
	EMbnContextNotActivated                                                 = 0x80548201
	EMbnBadSim                                                              = 0x80548202
	EMbnDataClassNotAvailable                                               = 0x80548203
	EMbnInvalidAccessString                                                 = 0x80548204
	EMbnMaxActivatedContexts                                                = 0x80548205
	EMbnPacketSvcDetached                                                   = 0x80548206
	EMbnProviderNotVisible                                                  = 0x80548207
	EMbnRadioPowerOff                                                       = 0x80548208
	EMbnServiceNotActivated                                                 = 0x80548209
	EMbnSimNotInserted                                                      = 0x8054820A
	EMbnVoiceCallInProgress                                                 = 0x8054820B
	EMbnInvalidCache                                                        = 0x8054820C
	EMbnNotRegistered                                                       = 0x8054820D
	EMbnProvidersNotFound                                                   = 0x8054820E
	EMbnPinNotSupported                                                     = 0x8054820F
	EMbnPinRequired                                                         = 0x80548210
	EMbnPinDisabled                                                         = 0x80548211
	EMbnFailure                                                             = 0x80548212
	EMbnInvalidProfile                                                      = 0x80548218
	EMbnDefaultProfileExist                                                 = 0x80548219
	EMbnSmsEncodingNotSupported                                             = 0x80548220
	EMbnSmsFilterNotSupported                                               = 0x80548221
	EMbnSmsInvalidMemoryIndex                                               = 0x80548222
	EMbnSmsLangNotSupported                                                 = 0x80548223
	EMbnSmsMemoryFailure                                                    = 0x80548224
	EMbnSmsNetworkTimeout                                                   = 0x80548225
	EMbnSmsUnknownSmscAddress                                               = 0x80548226
	EMbnSmsFormatNotSupported                                               = 0x80548227
	EMbnSmsOperationNotAllowed                                              = 0x80548228
	EMbnSmsMemoryFull                                                       = 0x80548229
	UiECreateFailed                                                         = 0x802A0001
	UiEShutdownCalled                                                       = 0x802A0002
	UiEIllegalReentrancy                                                    = 0x802A0003
	UiEObjectSealed                                                         = 0x802A0004
	UiEValueNotSet                                                          = 0x802A0005
	UiEValueNotDetermined                                                   = 0x802A0006
	UiEInvalidOutput                                                        = 0x802A0007
	UiEBooleanExpected                                                      = 0x802A0008
	UiEDifferentOwner                                                       = 0x802A0009
	UiEAmbiguousMatch                                                       = 0x802A000A
	UiEFpOverflow                                                           = 0x802A000B
	UiEWrongThread                                                          = 0x802A000C
	UiEStoryboardActive                                                     = 0x802A0101
	UiEStoryboardNotPlaying                                                 = 0x802A0102
	UiEStartKeyframeAfterEnd                                                = 0x802A0103
	UiEEndKeyframeNotDetermined                                             = 0x802A0104
	UiELoopsOverlap                                                         = 0x802A0105
	UiETransitionAlreadyUsed                                                = 0x802A0106
	UiETransitionNotInStoryboard                                            = 0x802A0107
	UiETransitionEclipsed                                                   = 0x802A0108
	UiETimeBeforeLastUpdate                                                 = 0x802A0109
	UiETimerClientAlreadyConnected                                          = 0x802A010A
)

func (k ErrorsKind) String() string {
	switch k {
	case 0:
		return "ErrorSuccess"
	case 1:
		return "ErrorInvalidFunction"
	case 2:
		return "ErrorFileNotFound"
	case 3:
		return "ErrorPathNotFound"
	case 4:
		return "ErrorTooManyOpenFiles"
	case 5:
		return "ErrorAccessDenied"
	case 6:
		return "ErrorInvalidHandle"
	case 7:
		return "ErrorArenaTrashed"
	case 8:
		return "ErrorNotEnoughMemory"
	case 9:
		return "ErrorInvalidBlock"
	case 10:
		return "ErrorBadEnvironment"
	case 11:
		return "ErrorBadFormat"
	case 12:
		return "ErrorInvalidAccess"
	case 13:
		return "ErrorInvalidData"
	case 14:
		return "ErrorOutofmemory"
	case 15:
		return "ErrorInvalidDrive"
	case 16:
		return "ErrorCurrentDirectory"
	case 17:
		return "ErrorNotSameDevice"
	case 18:
		return "ErrorNoMoreFiles"
	case 19:
		return "ErrorWriteProtect"
	case 20:
		return "ErrorBadUnit"
	case 21:
		return "ErrorNotReady"
	case 22:
		return "ErrorBadCommand"
	case 23:
		return "ErrorCrc"
	case 24:
		return "ErrorBadLength"
	case 25:
		return "ErrorSeek"
	case 26:
		return "ErrorNotDosDisk"
	case 27:
		return "ErrorSectorNotFound"
	case 28:
		return "ErrorOutOfPaper"
	case 29:
		return "ErrorWriteFault"
	case 30:
		return "ErrorReadFault"
	case 31:
		return "ErrorGenFailure"
	case 32:
		return "ErrorSharingViolation"
	case 33:
		return "ErrorLockViolation"
	case 34:
		return "ErrorWrongDisk"
	case 36:
		return "ErrorSharingBufferExceeded"
	case 38:
		return "ErrorHandleEof"
	case 39:
		return "ErrorHandleDiskFull"
	case 50:
		return "ErrorNotSupported"
	case 51:
		return "ErrorRemNotList"
	case 52:
		return "ErrorDupName"
	case 53:
		return "ErrorBadNetpath"
	case 54:
		return "ErrorNetworkBusy"
	case 55:
		return "ErrorDevNotExist"
	case 56:
		return "ErrorTooManyCmds"
	case 57:
		return "ErrorAdapHdwErr"
	case 58:
		return "ErrorBadNetResp"
	case 59:
		return "ErrorUnexpNetErr"
	case 60:
		return "ErrorBadRemAdap"
	case 61:
		return "ErrorPrintqFull"
	case 62:
		return "ErrorNoSpoolSpace"
	case 63:
		return "ErrorPrintCancelled"
	case 64:
		return "ErrorNetnameDeleted"
	case 65:
		return "ErrorNetworkAccessDenied"
	case 66:
		return "ErrorBadDevType"
	case 67:
		return "ErrorBadNetName"
	case 68:
		return "ErrorTooManyNames"
	case 69:
		return "ErrorTooManySess"
	case 70:
		return "ErrorSharingPaused"
	case 71:
		return "ErrorReqNotAccep"
	case 72:
		return "ErrorRedirPaused"
	case 80:
		return "ErrorFileExists"
	case 82:
		return "ErrorCannotMake"
	case 83:
		return "ErrorFailI24"
	case 84:
		return "ErrorOutOfStructures"
	case 85:
		return "ErrorAlreadyAssigned"
	case 86:
		return "ErrorInvalidPassword"
	case 87:
		return "ErrorInvalidParameter"
	case 88:
		return "ErrorNetWriteFault"
	case 89:
		return "ErrorNoProcSlots"
	case 100:
		return "ErrorTooManySemaphores"
	case 101:
		return "ErrorExclSemAlreadyOwned"
	case 102:
		return "ErrorSemIsSet"
	case 103:
		return "ErrorTooManySemRequests"
	case 104:
		return "ErrorInvalidAtInterruptTime"
	case 105:
		return "ErrorSemOwnerDied"
	case 106:
		return "ErrorSemUserLimit"
	case 107:
		return "ErrorDiskChange"
	case 108:
		return "ErrorDriveLocked"
	case 109:
		return "ErrorBrokenPipe"
	case 110:
		return "ErrorOpenFailed"
	case 111:
		return "ErrorBufferOverflow"
	case 112:
		return "ErrorDiskFull"
	case 113:
		return "ErrorNoMoreSearchHandles"
	case 114:
		return "ErrorInvalidTargetHandle"
	case 117:
		return "ErrorInvalidCategory"
	case 118:
		return "ErrorInvalidVerifySwitch"
	case 119:
		return "ErrorBadDriverLevel"
	case 120:
		return "ErrorCallNotImplemented"
	case 121:
		return "ErrorSemTimeout"
	case 122:
		return "ErrorInsufficientBuffer"
	case 123:
		return "ErrorInvalidName"
	case 124:
		return "ErrorInvalidLevel"
	case 125:
		return "ErrorNoVolumeLabel"
	case 126:
		return "ErrorModNotFound"
	case 127:
		return "ErrorProcNotFound"
	case 128:
		return "ErrorWaitNoChildren"
	case 129:
		return "ErrorChildNotComplete"
	case 130:
		return "ErrorDirectAccessHandle"
	case 131:
		return "ErrorNegativeSeek"
	case 132:
		return "ErrorSeekOnDevice"
	case 133:
		return "ErrorIsJoinTarget"
	case 134:
		return "ErrorIsJoined"
	case 135:
		return "ErrorIsSubsted"
	case 136:
		return "ErrorNotJoined"
	case 137:
		return "ErrorNotSubsted"
	case 138:
		return "ErrorJoinToJoin"
	case 139:
		return "ErrorSubstToSubst"
	case 140:
		return "ErrorJoinToSubst"
	case 141:
		return "ErrorSubstToJoin"
	case 142:
		return "ErrorBusyDrive"
	case 143:
		return "ErrorSameDrive"
	case 144:
		return "ErrorDirNotRoot"
	case 145:
		return "ErrorDirNotEmpty"
	case 146:
		return "ErrorIsSubstPath"
	case 147:
		return "ErrorIsJoinPath"
	case 148:
		return "ErrorPathBusy"
	case 149:
		return "ErrorIsSubstTarget"
	case 150:
		return "ErrorSystemTrace"
	case 151:
		return "ErrorInvalidEventCount"
	case 152:
		return "ErrorTooManyMuxwaiters"
	case 153:
		return "ErrorInvalidListFormat"
	case 154:
		return "ErrorLabelTooLong"
	case 155:
		return "ErrorTooManyTcbs"
	case 156:
		return "ErrorSignalRefused"
	case 157:
		return "ErrorDiscarded"
	case 158:
		return "ErrorNotLocked"
	case 159:
		return "ErrorBadThreadidAddr"
	case 160:
		return "ErrorBadArguments"
	case 161:
		return "ErrorBadPathname"
	case 162:
		return "ErrorSignalPending"
	case 164:
		return "ErrorMaxThrdsReached"
	case 167:
		return "ErrorLockFailed"
	case 170:
		return "ErrorBusy"
	case 173:
		return "ErrorCancelViolation"
	case 174:
		return "ErrorAtomicLocksNotSupported"
	case 180:
		return "ErrorInvalidSegmentNumber"
	case 182:
		return "ErrorInvalidOrdinal"
	case 183:
		return "ErrorAlreadyExists"
	case 186:
		return "ErrorInvalidFlagNumber"
	case 187:
		return "ErrorSemNotFound"
	case 188:
		return "ErrorInvalidStartingCodeseg"
	case 189:
		return "ErrorInvalidStackseg"
	case 190:
		return "ErrorInvalidModuletype"
	case 191:
		return "ErrorInvalidExeSignature"
	case 192:
		return "ErrorExeMarkedInvalid"
	case 193:
		return "ErrorBadExeFormat"
	case 194:
		return "ErrorIteratedDataExceeds64k"
	case 195:
		return "ErrorInvalidMinallocsize"
	case 196:
		return "ErrorDynlinkFromInvalidRing"
	case 197:
		return "ErrorIoplNotEnabled"
	case 198:
		return "ErrorInvalidSegdpl"
	case 199:
		return "ErrorAutodatasegExceeds64k"
	case 200:
		return "ErrorRing2SegMustBeMovable"
	case 201:
		return "ErrorRelocChainXeedsSeglim"
	case 202:
		return "ErrorInfloopInRelocChain"
	case 203:
		return "ErrorEnvvarNotFound"
	case 205:
		return "ErrorNoSignalSent"
	case 206:
		return "ErrorFilenameExcedRange"
	case 207:
		return "ErrorRing2StackInUse"
	case 208:
		return "ErrorMetaExpansionTooLong"
	case 209:
		return "ErrorInvalidSignalNumber"
	case 210:
		return "ErrorThread1Inactive"
	case 212:
		return "ErrorLocked"
	case 214:
		return "ErrorTooManyModules"
	case 215:
		return "ErrorNestingNotAllowed"
	case 216:
		return "ErrorExeMachineTypeMismatch"
	case 217:
		return "ErrorExeCannotModifySignedBinary"
	case 218:
		return "ErrorExeCannotModifyStrongSignedBinary"
	case 220:
		return "ErrorFileCheckedOut"
	case 221:
		return "ErrorCheckoutRequired"
	case 222:
		return "ErrorBadFileType"
	case 223:
		return "ErrorFileTooLarge"
	case 224:
		return "ErrorFormsAuthRequired"
	case 225:
		return "ErrorVirusInfected"
	case 226:
		return "ErrorVirusDeleted"
	case 229:
		return "ErrorPipeLocal"
	case 230:
		return "ErrorBadPipe"
	case 231:
		return "ErrorPipeBusy"
	case 232:
		return "ErrorNoData"
	case 233:
		return "ErrorPipeNotConnected"
	case 234:
		return "ErrorMoreData"
	case 240:
		return "ErrorVcDisconnected"
	case 254:
		return "ErrorInvalidEaName"
	case 255:
		return "ErrorEaListInconsistent"
	case 258:
		return "WaitTimeout"
	case 259:
		return "ErrorNoMoreItems"
	case 266:
		return "ErrorCannotCopy"
	case 267:
		return "ErrorDirectory"
	case 275:
		return "ErrorEasDidntFit"
	case 276:
		return "ErrorEaFileCorrupt"
	case 277:
		return "ErrorEaTableFull"
	case 278:
		return "ErrorInvalidEaHandle"
	case 282:
		return "ErrorEasNotSupported"
	case 288:
		return "ErrorNotOwner"
	case 298:
		return "ErrorTooManyPosts"
	case 299:
		return "ErrorPartialCopy"
	case 300:
		return "ErrorOplockNotGranted"
	case 301:
		return "ErrorInvalidOplockProtocol"
	case 302:
		return "ErrorDiskTooFragmented"
	case 303:
		return "ErrorDeletePending"
	case 304:
		return "ErrorIncompatibleWithGlobalShortNameRegistrySetting"
	case 305:
		return "ErrorShortNamesNotEnabledOnVolume"
	case 306:
		return "ErrorSecurityStreamIsInconsistent"
	case 307:
		return "ErrorInvalidLockRange"
	case 308:
		return "ErrorImageSubsystemNotPresent"
	case 309:
		return "ErrorNotificationGuidAlreadyDefined"
	case 317:
		return "ErrorMrMidNotFound"
	case 318:
		return "ErrorScopeNotFound"
	case 350:
		return "ErrorFailNoactionReboot"
	case 351:
		return "ErrorFailShutdown"
	case 352:
		return "ErrorFailRestart"
	case 353:
		return "ErrorMaxSessionsReached"
	case 400:
		return "ErrorThreadModeAlreadyBackground"
	case 401:
		return "ErrorThreadModeNotBackground"
	case 402:
		return "ErrorProcessModeAlreadyBackground"
	case 403:
		return "ErrorProcessModeNotBackground"
	case 487:
		return "ErrorInvalidAddress"
	case 500:
		return "ErrorUserProfileLoad"
	case 534:
		return "ErrorArithmeticOverflow"
	case 535:
		return "ErrorPipeConnected"
	case 536:
		return "ErrorPipeListening"
	case 537:
		return "ErrorVerifierStop"
	case 538:
		return "ErrorAbiosError"
	case 539:
		return "ErrorWx86Warning"
	case 540:
		return "ErrorWx86Error"
	case 541:
		return "ErrorTimerNotCanceled"
	case 542:
		return "ErrorUnwind"
	case 543:
		return "ErrorBadStack"
	case 544:
		return "ErrorInvalidUnwindTarget"
	case 545:
		return "ErrorInvalidPortAttributes"
	case 546:
		return "ErrorPortMessageTooLong"
	case 547:
		return "ErrorInvalidQuotaLower"
	case 548:
		return "ErrorDeviceAlreadyAttached"
	case 549:
		return "ErrorInstructionMisalignment"
	case 550:
		return "ErrorProfilingNotStarted"
	case 551:
		return "ErrorProfilingNotStopped"
	case 552:
		return "ErrorCouldNotInterpret"
	case 553:
		return "ErrorProfilingAtLimit"
	case 554:
		return "ErrorCantWait"
	case 555:
		return "ErrorCantTerminateSelf"
	case 556:
		return "ErrorUnexpectedMmCreateErr"
	case 557:
		return "ErrorUnexpectedMmMapError"
	case 558:
		return "ErrorUnexpectedMmExtendErr"
	case 559:
		return "ErrorBadFunctionTable"
	case 560:
		return "ErrorNoGuidTranslation"
	case 561:
		return "ErrorInvalidLdtSize"
	case 563:
		return "ErrorInvalidLdtOffset"
	case 564:
		return "ErrorInvalidLdtDescriptor"
	case 565:
		return "ErrorTooManyThreads"
	case 566:
		return "ErrorThreadNotInProcess"
	case 567:
		return "ErrorPagefileQuotaExceeded"
	case 568:
		return "ErrorLogonServerConflict"
	case 569:
		return "ErrorSynchronizationRequired"
	case 570:
		return "ErrorNetOpenFailed"
	case 571:
		return "ErrorIoPrivilegeFailed"
	case 572:
		return "ErrorControlCExit"
	case 573:
		return "ErrorMissingSystemfile"
	case 574:
		return "ErrorUnhandledException"
	case 575:
		return "ErrorAppInitFailure"
	case 576:
		return "ErrorPagefileCreateFailed"
	case 577:
		return "ErrorInvalidImageHash"
	case 578:
		return "ErrorNoPagefile"
	case 579:
		return "ErrorIllegalFloatContext"
	case 580:
		return "ErrorNoEventPair"
	case 581:
		return "ErrorDomainCtrlrConfigError"
	case 582:
		return "ErrorIllegalCharacter"
	case 583:
		return "ErrorUndefinedCharacter"
	case 584:
		return "ErrorFloppyVolume"
	case 585:
		return "ErrorBiosFailedToConnectInterrupt"
	case 586:
		return "ErrorBackupController"
	case 587:
		return "ErrorMutantLimitExceeded"
	case 588:
		return "ErrorFsDriverRequired"
	case 589:
		return "ErrorCannotLoadRegistryFile"
	case 590:
		return "ErrorDebugAttachFailed"
	case 591:
		return "ErrorSystemProcessTerminated"
	case 592:
		return "ErrorDataNotAccepted"
	case 593:
		return "ErrorVdmHardError"
	case 594:
		return "ErrorDriverCancelTimeout"
	case 595:
		return "ErrorReplyMessageMismatch"
	case 596:
		return "ErrorLostWritebehindData"
	case 597:
		return "ErrorClientServerParametersInvalid"
	case 598:
		return "ErrorNotTinyStream"
	case 599:
		return "ErrorStackOverflowRead"
	case 600:
		return "ErrorConvertToLarge"
	case 601:
		return "ErrorFoundOutOfScope"
	case 602:
		return "ErrorAllocateBucket"
	case 603:
		return "ErrorMarshallOverflow"
	case 604:
		return "ErrorInvalidVariant"
	case 605:
		return "ErrorBadCompressionBuffer"
	case 606:
		return "ErrorAuditFailed"
	case 607:
		return "ErrorTimerResolutionNotSet"
	case 608:
		return "ErrorInsufficientLogonInfo"
	case 609:
		return "ErrorBadDllEntrypoint"
	case 610:
		return "ErrorBadServiceEntrypoint"
	case 611:
		return "ErrorIpAddressConflict1"
	case 612:
		return "ErrorIpAddressConflict2"
	case 613:
		return "ErrorRegistryQuotaLimit"
	case 614:
		return "ErrorNoCallbackActive"
	case 615:
		return "ErrorPwdTooShort"
	case 616:
		return "ErrorPwdTooRecent"
	case 617:
		return "ErrorPwdHistoryConflict"
	case 618:
		return "ErrorUnsupportedCompression"
	case 619:
		return "ErrorInvalidHwProfile"
	case 620:
		return "ErrorInvalidPlugplayDevicePath"
	case 621:
		return "ErrorQuotaListInconsistent"
	case 622:
		return "ErrorEvaluationExpiration"
	case 623:
		return "ErrorIllegalDllRelocation"
	case 624:
		return "ErrorDllInitFailedLogoff"
	case 625:
		return "ErrorValidateContinue"
	case 626:
		return "ErrorNoMoreMatches"
	case 627:
		return "ErrorRangeListConflict"
	case 628:
		return "ErrorServerSidMismatch"
	case 629:
		return "ErrorCantEnableDenyOnly"
	case 630:
		return "ErrorFloatMultipleFaults"
	case 631:
		return "ErrorFloatMultipleTraps"
	case 632:
		return "ErrorNointerface"
	case 633:
		return "ErrorDriverFailedSleep"
	case 634:
		return "ErrorCorruptSystemFile"
	case 635:
		return "ErrorCommitmentMinimum"
	case 636:
		return "ErrorPnpRestartEnumeration"
	case 637:
		return "ErrorSystemImageBadSignature"
	case 638:
		return "ErrorPnpRebootRequired"
	case 639:
		return "ErrorInsufficientPower"
	case 640:
		return "ErrorMultipleFaultViolation"
	case 641:
		return "ErrorSystemShutdown"
	case 642:
		return "ErrorPortNotSet"
	case 643:
		return "ErrorDsVersionCheckFailure"
	case 644:
		return "ErrorRangeNotFound"
	case 646:
		return "ErrorNotSafeModeDriver"
	case 647:
		return "ErrorFailedDriverEntry"
	case 648:
		return "ErrorDeviceEnumerationError"
	case 649:
		return "ErrorMountPointNotResolved"
	case 650:
		return "ErrorInvalidDeviceObjectParameter"
	case 651:
		return "ErrorMcaOccured"
	case 652:
		return "ErrorDriverDatabaseError"
	case 653:
		return "ErrorSystemHiveTooLarge"
	case 654:
		return "ErrorDriverFailedPriorUnload"
	case 655:
		return "ErrorVolsnapPrepareHibernate"
	case 656:
		return "ErrorHibernationFailure"
	case 665:
		return "ErrorFileSystemLimitation"
	case 668:
		return "ErrorAssertionFailure"
	case 669:
		return "ErrorAcpiError"
	case 670:
		return "ErrorWowAssertion"
	case 671:
		return "ErrorPnpBadMpsTable"
	case 672:
		return "ErrorPnpTranslationFailed"
	case 673:
		return "ErrorPnpIrqTranslationFailed"
	case 674:
		return "ErrorPnpInvalidId"
	case 675:
		return "ErrorWakeSystemDebugger"
	case 676:
		return "ErrorHandlesClosed"
	case 677:
		return "ErrorExtraneousInformation"
	case 678:
		return "ErrorRxactCommitNecessary"
	case 679:
		return "ErrorMediaCheck"
	case 680:
		return "ErrorGuidSubstitutionMade"
	case 681:
		return "ErrorStoppedOnSymlink"
	case 682:
		return "ErrorLongjump"
	case 683:
		return "ErrorPlugplayQueryVetoed"
	case 684:
		return "ErrorUnwindConsolidate"
	case 685:
		return "ErrorRegistryHiveRecovered"
	case 686:
		return "ErrorDllMightBeInsecure"
	case 687:
		return "ErrorDllMightBeIncompatible"
	case 688:
		return "ErrorDbgExceptionNotHandled"
	case 689:
		return "ErrorDbgReplyLater"
	case 690:
		return "ErrorDbgUnableToProvideHandle"
	case 691:
		return "ErrorDbgTerminateThread"
	case 692:
		return "ErrorDbgTerminateProcess"
	case 693:
		return "ErrorDbgControlC"
	case 694:
		return "ErrorDbgPrintexceptionC"
	case 695:
		return "ErrorDbgRipexception"
	case 696:
		return "ErrorDbgControlBreak"
	case 697:
		return "ErrorDbgCommandException"
	case 698:
		return "ErrorObjectNameExists"
	case 699:
		return "ErrorThreadWasSuspended"
	case 700:
		return "ErrorImageNotAtBase"
	case 701:
		return "ErrorRxactStateCreated"
	case 702:
		return "ErrorSegmentNotification"
	case 703:
		return "ErrorBadCurrentDirectory"
	case 704:
		return "ErrorFtReadRecoveryFromBackup"
	case 705:
		return "ErrorFtWriteRecovery"
	case 706:
		return "ErrorImageMachineTypeMismatch"
	case 707:
		return "ErrorReceivePartial"
	case 708:
		return "ErrorReceiveExpedited"
	case 709:
		return "ErrorReceivePartialExpedited"
	case 710:
		return "ErrorEventDone"
	case 711:
		return "ErrorEventPending"
	case 712:
		return "ErrorCheckingFileSystem"
	case 713:
		return "ErrorFatalAppExit"
	case 714:
		return "ErrorPredefinedHandle"
	case 715:
		return "ErrorWasUnlocked"
	case 716:
		return "ErrorServiceNotification"
	case 717:
		return "ErrorWasLocked"
	case 718:
		return "ErrorLogHardError"
	case 719:
		return "ErrorAlreadyWin32"
	case 720:
		return "ErrorImageMachineTypeMismatchExe"
	case 721:
		return "ErrorNoYieldPerformed"
	case 722:
		return "ErrorTimerResumeIgnored"
	case 723:
		return "ErrorArbitrationUnhandled"
	case 724:
		return "ErrorCardbusNotSupported"
	case 725:
		return "ErrorMpProcessorMismatch"
	case 726:
		return "ErrorHibernated"
	case 727:
		return "ErrorResumeHibernation"
	case 728:
		return "ErrorFirmwareUpdated"
	case 729:
		return "ErrorDriversLeakingLockedPages"
	case 730:
		return "ErrorWakeSystem"
	case 731:
		return "ErrorWait1"
	case 732:
		return "ErrorWait2"
	case 733:
		return "ErrorWait3"
	case 734:
		return "ErrorWait63"
	case 735:
		return "ErrorAbandonedWait0"
	case 736:
		return "ErrorAbandonedWait63"
	case 737:
		return "ErrorUserApc"
	case 738:
		return "ErrorKernelApc"
	case 739:
		return "ErrorAlerted"
	case 740:
		return "ErrorElevationRequired"
	case 741:
		return "ErrorReparse"
	case 742:
		return "ErrorOplockBreakInProgress"
	case 743:
		return "ErrorVolumeMounted"
	case 744:
		return "ErrorRxactCommitted"
	case 745:
		return "ErrorNotifyCleanup"
	case 746:
		return "ErrorPrimaryTransportConnectFailed"
	case 747:
		return "ErrorPageFaultTransition"
	case 748:
		return "ErrorPageFaultDemandZero"
	case 749:
		return "ErrorPageFaultCopyOnWrite"
	case 750:
		return "ErrorPageFaultGuardPage"
	case 751:
		return "ErrorPageFaultPagingFile"
	case 752:
		return "ErrorCachePageLocked"
	case 753:
		return "ErrorCrashDump"
	case 754:
		return "ErrorBufferAllZeros"
	case 755:
		return "ErrorReparseObject"
	case 756:
		return "ErrorResourceRequirementsChanged"
	case 757:
		return "ErrorTranslationComplete"
	case 758:
		return "ErrorNothingToTerminate"
	case 759:
		return "ErrorProcessNotInJob"
	case 760:
		return "ErrorProcessInJob"
	case 761:
		return "ErrorVolsnapHibernateReady"
	case 762:
		return "ErrorFsfilterOpCompletedSuccessfully"
	case 763:
		return "ErrorInterruptVectorAlreadyConnected"
	case 764:
		return "ErrorInterruptStillConnected"
	case 765:
		return "ErrorWaitForOplock"
	case 766:
		return "ErrorDbgExceptionHandled"
	case 767:
		return "ErrorDbgContinue"
	case 768:
		return "ErrorCallbackPopStack"
	case 769:
		return "ErrorCompressionDisabled"
	case 770:
		return "ErrorCantfetchbackwards"
	case 771:
		return "ErrorCantscrollbackwards"
	case 772:
		return "ErrorRowsnotreleased"
	case 773:
		return "ErrorBadAccessorFlags"
	case 774:
		return "ErrorErrorsEncountered"
	case 775:
		return "ErrorNotCapable"
	case 776:
		return "ErrorRequestOutOfSequence"
	case 777:
		return "ErrorVersionParseError"
	case 778:
		return "ErrorBadstartposition"
	case 779:
		return "ErrorMemoryHardware"
	case 780:
		return "ErrorDiskRepairDisabled"
	case 781:
		return "ErrorInsufficientResourceForSpecifiedSharedSectionSize"
	case 782:
		return "ErrorSystemPowerstateTransition"
	case 783:
		return "ErrorSystemPowerstateComplexTransition"
	case 784:
		return "ErrorMcaException"
	case 785:
		return "ErrorAccessAuditByPolicy"
	case 786:
		return "ErrorAccessDisabledNoSaferUiByPolicy"
	case 787:
		return "ErrorAbandonHiberfile"
	case 788:
		return "ErrorLostWritebehindDataNetworkDisconnected"
	case 789:
		return "ErrorLostWritebehindDataNetworkServerError"
	case 790:
		return "ErrorLostWritebehindDataLocalDiskError"
	case 791:
		return "ErrorBadMcfgTable"
	case 800:
		return "ErrorOplockSwitchedToNewHandle"
	case 801:
		return "ErrorCannotGrantRequestedOplock"
	case 802:
		return "ErrorCannotBreakOplock"
	case 803:
		return "ErrorOplockHandleClosed"
	case 804:
		return "ErrorNoAceCondition"
	case 805:
		return "ErrorInvalidAceCondition"
	case 994:
		return "ErrorEaAccessDenied"
	case 995:
		return "ErrorOperationAborted"
	case 996:
		return "ErrorIoIncomplete"
	case 997:
		return "ErrorIoPending"
	case 998:
		return "ErrorNoaccess"
	case 999:
		return "ErrorSwaperror"
	case 1001:
		return "ErrorStackOverflow"
	case 1002:
		return "ErrorInvalidMessage"
	case 1003:
		return "ErrorCanNotComplete"
	case 1004:
		return "ErrorInvalidFlags"
	case 1005:
		return "ErrorUnrecognizedVolume"
	case 1006:
		return "ErrorFileInvalid"
	case 1007:
		return "ErrorFullscreenMode"
	case 1008:
		return "ErrorNoToken"
	case 1009:
		return "ErrorBaddb"
	case 1010:
		return "ErrorBadkey"
	case 1011:
		return "ErrorCantopen"
	case 1012:
		return "ErrorCantread"
	case 1013:
		return "ErrorCantwrite"
	case 1014:
		return "ErrorRegistryRecovered"
	case 1015:
		return "ErrorRegistryCorrupt"
	case 1016:
		return "ErrorRegistryIoFailed"
	case 1017:
		return "ErrorNotRegistryFile"
	case 1018:
		return "ErrorKeyDeleted"
	case 1019:
		return "ErrorNoLogSpace"
	case 1020:
		return "ErrorKeyHasChildren"
	case 1021:
		return "ErrorChildMustBeVolatile"
	case 1022:
		return "ErrorNotifyEnumDir"
	case 1051:
		return "ErrorDependentServicesRunning"
	case 1052:
		return "ErrorInvalidServiceControl"
	case 1053:
		return "ErrorServiceRequestTimeout"
	case 1054:
		return "ErrorServiceNoThread"
	case 1055:
		return "ErrorServiceDatabaseLocked"
	case 1056:
		return "ErrorServiceAlreadyRunning"
	case 1057:
		return "ErrorInvalidServiceAccount"
	case 1058:
		return "ErrorServiceDisabled"
	case 1059:
		return "ErrorCircularDependency"
	case 1060:
		return "ErrorServiceDoesNotExist"
	case 1061:
		return "ErrorServiceCannotAcceptCtrl"
	case 1062:
		return "ErrorServiceNotActive"
	case 1063:
		return "ErrorFailedServiceControllerConnect"
	case 1064:
		return "ErrorExceptionInService"
	case 1065:
		return "ErrorDatabaseDoesNotExist"
	case 1066:
		return "ErrorServiceSpecificError"
	case 1067:
		return "ErrorProcessAborted"
	case 1068:
		return "ErrorServiceDependencyFail"
	case 1069:
		return "ErrorServiceLogonFailed"
	case 1070:
		return "ErrorServiceStartHang"
	case 1071:
		return "ErrorInvalidServiceLock"
	case 1072:
		return "ErrorServiceMarkedForDelete"
	case 1073:
		return "ErrorServiceExists"
	case 1074:
		return "ErrorAlreadyRunningLkg"
	case 1075:
		return "ErrorServiceDependencyDeleted"
	case 1076:
		return "ErrorBootAlreadyAccepted"
	case 1077:
		return "ErrorServiceNeverStarted"
	case 1078:
		return "ErrorDuplicateServiceName"
	case 1079:
		return "ErrorDifferentServiceAccount"
	case 1080:
		return "ErrorCannotDetectDriverFailure"
	case 1081:
		return "ErrorCannotDetectProcessAbort"
	case 1082:
		return "ErrorNoRecoveryProgram"
	case 1083:
		return "ErrorServiceNotInExe"
	case 1084:
		return "ErrorNotSafebootService"
	case 1100:
		return "ErrorEndOfMedia"
	case 1101:
		return "ErrorFilemarkDetected"
	case 1102:
		return "ErrorBeginningOfMedia"
	case 1103:
		return "ErrorSetmarkDetected"
	case 1104:
		return "ErrorNoDataDetected"
	case 1105:
		return "ErrorPartitionFailure"
	case 1106:
		return "ErrorInvalidBlockLength"
	case 1107:
		return "ErrorDeviceNotPartitioned"
	case 1108:
		return "ErrorUnableToLockMedia"
	case 1109:
		return "ErrorUnableToUnloadMedia"
	case 1110:
		return "ErrorMediaChanged"
	case 1111:
		return "ErrorBusReset"
	case 1112:
		return "ErrorNoMediaInDrive"
	case 1113:
		return "ErrorNoUnicodeTranslation"
	case 1114:
		return "ErrorDllInitFailed"
	case 1115:
		return "ErrorShutdownInProgress"
	case 1116:
		return "ErrorNoShutdownInProgress"
	case 1117:
		return "ErrorIoDevice"
	case 1118:
		return "ErrorSerialNoDevice"
	case 1119:
		return "ErrorIrqBusy"
	case 1120:
		return "ErrorMoreWrites"
	case 1121:
		return "ErrorCounterTimeout"
	case 1122:
		return "ErrorFloppyIdMarkNotFound"
	case 1123:
		return "ErrorFloppyWrongCylinder"
	case 1124:
		return "ErrorFloppyUnknownError"
	case 1125:
		return "ErrorFloppyBadRegisters"
	case 1126:
		return "ErrorDiskRecalibrateFailed"
	case 1127:
		return "ErrorDiskOperationFailed"
	case 1128:
		return "ErrorDiskResetFailed"
	case 1129:
		return "ErrorEomOverflow"
	case 1130:
		return "ErrorNotEnoughServerMemory"
	case 1131:
		return "ErrorPossibleDeadlock"
	case 1132:
		return "ErrorMappedAlignment"
	case 1140:
		return "ErrorSetPowerStateVetoed"
	case 1141:
		return "ErrorSetPowerStateFailed"
	case 1142:
		return "ErrorTooManyLinks"
	case 1150:
		return "ErrorOldWinVersion"
	case 1151:
		return "ErrorAppWrongOs"
	case 1152:
		return "ErrorSingleInstanceApp"
	case 1153:
		return "ErrorRmodeApp"
	case 1154:
		return "ErrorInvalidDll"
	case 1155:
		return "ErrorNoAssociation"
	case 1156:
		return "ErrorDdeFail"
	case 1157:
		return "ErrorDllNotFound"
	case 1158:
		return "ErrorNoMoreUserHandles"
	case 1159:
		return "ErrorMessageSyncOnly"
	case 1160:
		return "ErrorSourceElementEmpty"
	case 1161:
		return "ErrorDestinationElementFull"
	case 1162:
		return "ErrorIllegalElementAddress"
	case 1163:
		return "ErrorMagazineNotPresent"
	case 1164:
		return "ErrorDeviceReinitializationNeeded"
	case 1165:
		return "ErrorDeviceRequiresCleaning"
	case 1166:
		return "ErrorDeviceDoorOpen"
	case 1167:
		return "ErrorDeviceNotConnected"
	case 1168:
		return "ErrorNotFound"
	case 1169:
		return "ErrorNoMatch"
	case 1170:
		return "ErrorSetNotFound"
	case 1171:
		return "ErrorPointNotFound"
	case 1172:
		return "ErrorNoTrackingService"
	case 1173:
		return "ErrorNoVolumeId"
	case 1175:
		return "ErrorUnableToRemoveReplaced"
	case 1176:
		return "ErrorUnableToMoveReplacement"
	case 1177:
		return "ErrorUnableToMoveReplacement2"
	case 1178:
		return "ErrorJournalDeleteInProgress"
	case 1179:
		return "ErrorJournalNotActive"
	case 1180:
		return "ErrorPotentialFileFound"
	case 1181:
		return "ErrorJournalEntryDeleted"
	case 1190:
		return "ErrorShutdownIsScheduled"
	case 1191:
		return "ErrorShutdownUsersLoggedOn"
	case 1200:
		return "ErrorBadDevice"
	case 1201:
		return "ErrorConnectionUnavail"
	case 1202:
		return "ErrorDeviceAlreadyRemembered"
	case 1203:
		return "ErrorNoNetOrBadPath"
	case 1204:
		return "ErrorBadProvider"
	case 1205:
		return "ErrorCannotOpenProfile"
	case 1206:
		return "ErrorBadProfile"
	case 1207:
		return "ErrorNotContainer"
	case 1208:
		return "ErrorExtendedError"
	case 1209:
		return "ErrorInvalidGroupname"
	case 1210:
		return "ErrorInvalidComputername"
	case 1211:
		return "ErrorInvalidEventname"
	case 1212:
		return "ErrorInvalidDomainname"
	case 1213:
		return "ErrorInvalidServicename"
	case 1214:
		return "ErrorInvalidNetname"
	case 1215:
		return "ErrorInvalidSharename"
	case 1216:
		return "ErrorInvalidPasswordname"
	case 1217:
		return "ErrorInvalidMessagename"
	case 1218:
		return "ErrorInvalidMessagedest"
	case 1219:
		return "ErrorSessionCredentialConflict"
	case 1220:
		return "ErrorRemoteSessionLimitExceeded"
	case 1221:
		return "ErrorDupDomainname"
	case 1222:
		return "ErrorNoNetwork"
	case 1223:
		return "ErrorCancelled"
	case 1224:
		return "ErrorUserMappedFile"
	case 1225:
		return "ErrorConnectionRefused"
	case 1226:
		return "ErrorGracefulDisconnect"
	case 1227:
		return "ErrorAddressAlreadyAssociated"
	case 1228:
		return "ErrorAddressNotAssociated"
	case 1229:
		return "ErrorConnectionInvalid"
	case 1230:
		return "ErrorConnectionActive"
	case 1231:
		return "ErrorNetworkUnreachable"
	case 1232:
		return "ErrorHostUnreachable"
	case 1233:
		return "ErrorProtocolUnreachable"
	case 1234:
		return "ErrorPortUnreachable"
	case 1235:
		return "ErrorRequestAborted"
	case 1236:
		return "ErrorConnectionAborted"
	case 1237:
		return "ErrorRetry"
	case 1238:
		return "ErrorConnectionCountLimit"
	case 1239:
		return "ErrorLoginTimeRestriction"
	case 1240:
		return "ErrorLoginWkstaRestriction"
	case 1241:
		return "ErrorIncorrectAddress"
	case 1242:
		return "ErrorAlreadyRegistered"
	case 1243:
		return "ErrorServiceNotFound"
	case 1244:
		return "ErrorNotAuthenticated"
	case 1245:
		return "ErrorNotLoggedOn"
	case 1246:
		return "ErrorContinue"
	case 1247:
		return "ErrorAlreadyInitialized"
	case 1248:
		return "ErrorNoMoreDevices"
	case 1249:
		return "ErrorNoSuchSite"
	case 1250:
		return "ErrorDomainControllerExists"
	case 1251:
		return "ErrorOnlyIfConnected"
	case 1252:
		return "ErrorOverrideNochanges"
	case 1253:
		return "ErrorBadUserProfile"
	case 1254:
		return "ErrorNotSupportedOnSbs"
	case 1255:
		return "ErrorServerShutdownInProgress"
	case 1256:
		return "ErrorHostDown"
	case 1257:
		return "ErrorNonAccountSid"
	case 1258:
		return "ErrorNonDomainSid"
	case 1259:
		return "ErrorApphelpBlock"
	case 1260:
		return "ErrorAccessDisabledByPolicy"
	case 1261:
		return "ErrorRegNatConsumption"
	case 1262:
		return "ErrorCscshareOffline"
	case 1263:
		return "ErrorPkinitFailure"
	case 1264:
		return "ErrorSmartcardSubsystemFailure"
	case 1265:
		return "ErrorDowngradeDetected"
	case 1271:
		return "ErrorMachineLocked"
	case 1273:
		return "ErrorCallbackSuppliedInvalidData"
	case 1274:
		return "ErrorSyncForegroundRefreshRequired"
	case 1275:
		return "ErrorDriverBlocked"
	case 1276:
		return "ErrorInvalidImportOfNonDll"
	case 1277:
		return "ErrorAccessDisabledWebblade"
	case 1278:
		return "ErrorAccessDisabledWebbladeTamper"
	case 1279:
		return "ErrorRecoveryFailure"
	case 1280:
		return "ErrorAlreadyFiber"
	case 1281:
		return "ErrorAlreadyThread"
	case 1282:
		return "ErrorStackBufferOverrun"
	case 1283:
		return "ErrorParameterQuotaExceeded"
	case 1284:
		return "ErrorDebuggerInactive"
	case 1285:
		return "ErrorDelayLoadFailed"
	case 1286:
		return "ErrorVdmDisallowed"
	case 1287:
		return "ErrorUnidentifiedError"
	case 1288:
		return "ErrorInvalidCruntimeParameter"
	case 1289:
		return "ErrorBeyondVdl"
	case 1290:
		return "ErrorIncompatibleServiceSidType"
	case 1291:
		return "ErrorDriverProcessTerminated"
	case 1292:
		return "ErrorImplementationLimit"
	case 1293:
		return "ErrorProcessIsProtected"
	case 1294:
		return "ErrorServiceNotifyClientLagging"
	case 1295:
		return "ErrorDiskQuotaExceeded"
	case 1296:
		return "ErrorContentBlocked"
	case 1297:
		return "ErrorIncompatibleServicePrivilege"
	case 1298:
		return "ErrorAppHang"
	case 1299:
		return "ErrorInvalidLabel"
	case 1300:
		return "ErrorNotAllAssigned"
	case 1301:
		return "ErrorSomeNotMapped"
	case 1302:
		return "ErrorNoQuotasForAccount"
	case 1303:
		return "ErrorLocalUserSessionKey"
	case 1304:
		return "ErrorNullLmPassword"
	case 1305:
		return "ErrorUnknownRevision"
	case 1306:
		return "ErrorRevisionMismatch"
	case 1307:
		return "ErrorInvalidOwner"
	case 1308:
		return "ErrorInvalidPrimaryGroup"
	case 1309:
		return "ErrorNoImpersonationToken"
	case 1310:
		return "ErrorCantDisableMandatory"
	case 1311:
		return "ErrorNoLogonServers"
	case 1312:
		return "ErrorNoSuchLogonSession"
	case 1313:
		return "ErrorNoSuchPrivilege"
	case 1314:
		return "ErrorPrivilegeNotHeld"
	case 1315:
		return "ErrorInvalidAccountName"
	case 1316:
		return "ErrorUserExists"
	case 1317:
		return "ErrorNoSuchUser"
	case 1318:
		return "ErrorGroupExists"
	case 1319:
		return "ErrorNoSuchGroup"
	case 1320:
		return "ErrorMemberInGroup"
	case 1321:
		return "ErrorMemberNotInGroup"
	case 1322:
		return "ErrorLastAdmin"
	case 1323:
		return "ErrorWrongPassword"
	case 1324:
		return "ErrorIllFormedPassword"
	case 1325:
		return "ErrorPasswordRestriction"
	case 1326:
		return "ErrorLogonFailure"
	case 1327:
		return "ErrorAccountRestriction"
	case 1328:
		return "ErrorInvalidLogonHours"
	case 1329:
		return "ErrorInvalidWorkstation"
	case 1330:
		return "ErrorPasswordExpired"
	case 1331:
		return "ErrorAccountDisabled"
	case 1332:
		return "ErrorNoneMapped"
	case 1333:
		return "ErrorTooManyLuidsRequested"
	case 1334:
		return "ErrorLuidsExhausted"
	case 1335:
		return "ErrorInvalidSubAuthority"
	case 1336:
		return "ErrorInvalidAcl"
	case 1337:
		return "ErrorInvalidSid"
	case 1338:
		return "ErrorInvalidSecurityDescr"
	case 1340:
		return "ErrorBadInheritanceAcl"
	case 1341:
		return "ErrorServerDisabled"
	case 1342:
		return "ErrorServerNotDisabled"
	case 1343:
		return "ErrorInvalidIdAuthority"
	case 1344:
		return "ErrorAllottedSpaceExceeded"
	case 1345:
		return "ErrorInvalidGroupAttributes"
	case 1346:
		return "ErrorBadImpersonationLevel"
	case 1347:
		return "ErrorCantOpenAnonymous"
	case 1348:
		return "ErrorBadValidationClass"
	case 1349:
		return "ErrorBadTokenType"
	case 1350:
		return "ErrorNoSecurityOnObject"
	case 1351:
		return "ErrorCantAccessDomainInfo"
	case 1352:
		return "ErrorInvalidServerState"
	case 1353:
		return "ErrorInvalidDomainState"
	case 1354:
		return "ErrorInvalidDomainRole"
	case 1355:
		return "ErrorNoSuchDomain"
	case 1356:
		return "ErrorDomainExists"
	case 1357:
		return "ErrorDomainLimitExceeded"
	case 1358:
		return "ErrorInternalDbCorruption"
	case 1359:
		return "ErrorInternalError"
	case 1360:
		return "ErrorGenericNotMapped"
	case 1361:
		return "ErrorBadDescriptorFormat"
	case 1362:
		return "ErrorNotLogonProcess"
	case 1363:
		return "ErrorLogonSessionExists"
	case 1364:
		return "ErrorNoSuchPackage"
	case 1365:
		return "ErrorBadLogonSessionState"
	case 1366:
		return "ErrorLogonSessionCollision"
	case 1367:
		return "ErrorInvalidLogonType"
	case 1368:
		return "ErrorCannotImpersonate"
	case 1369:
		return "ErrorRxactInvalidState"
	case 1370:
		return "ErrorRxactCommitFailure"
	case 1371:
		return "ErrorSpecialAccount"
	case 1372:
		return "ErrorSpecialGroup"
	case 1373:
		return "ErrorSpecialUser"
	case 1374:
		return "ErrorMembersPrimaryGroup"
	case 1375:
		return "ErrorTokenAlreadyInUse"
	case 1376:
		return "ErrorNoSuchAlias"
	case 1377:
		return "ErrorMemberNotInAlias"
	case 1378:
		return "ErrorMemberInAlias"
	case 1379:
		return "ErrorAliasExists"
	case 1380:
		return "ErrorLogonNotGranted"
	case 1381:
		return "ErrorTooManySecrets"
	case 1382:
		return "ErrorSecretTooLong"
	case 1383:
		return "ErrorInternalDbError"
	case 1384:
		return "ErrorTooManyContextIds"
	case 1385:
		return "ErrorLogonTypeNotGranted"
	case 1386:
		return "ErrorNtCrossEncryptionRequired"
	case 1387:
		return "ErrorNoSuchMember"
	case 1388:
		return "ErrorInvalidMember"
	case 1389:
		return "ErrorTooManySids"
	case 1390:
		return "ErrorLmCrossEncryptionRequired"
	case 1391:
		return "ErrorNoInheritance"
	case 1392:
		return "ErrorFileCorrupt"
	case 1393:
		return "ErrorDiskCorrupt"
	case 1394:
		return "ErrorNoUserSessionKey"
	case 1395:
		return "ErrorLicenseQuotaExceeded"
	case 1396:
		return "ErrorWrongTargetName"
	case 1397:
		return "ErrorMutualAuthFailed"
	case 1398:
		return "ErrorTimeSkew"
	case 1399:
		return "ErrorCurrentDomainNotAllowed"
	case 1400:
		return "ErrorInvalidWindowHandle"
	case 1401:
		return "ErrorInvalidMenuHandle"
	case 1402:
		return "ErrorInvalidCursorHandle"
	case 1403:
		return "ErrorInvalidAccelHandle"
	case 1404:
		return "ErrorInvalidHookHandle"
	case 1405:
		return "ErrorInvalidDwpHandle"
	case 1406:
		return "ErrorTlwWithWschild"
	case 1407:
		return "ErrorCannotFindWndClass"
	case 1408:
		return "ErrorWindowOfOtherThread"
	case 1409:
		return "ErrorHotkeyAlreadyRegistered"
	case 1410:
		return "ErrorClassAlreadyExists"
	case 1411:
		return "ErrorClassDoesNotExist"
	case 1412:
		return "ErrorClassHasWindows"
	case 1413:
		return "ErrorInvalidIndex"
	case 1414:
		return "ErrorInvalidIconHandle"
	case 1415:
		return "ErrorPrivateDialogIndex"
	case 1416:
		return "ErrorListboxIdNotFound"
	case 1417:
		return "ErrorNoWildcardCharacters"
	case 1418:
		return "ErrorClipboardNotOpen"
	case 1419:
		return "ErrorHotkeyNotRegistered"
	case 1420:
		return "ErrorWindowNotDialog"
	case 1421:
		return "ErrorControlIdNotFound"
	case 1422:
		return "ErrorInvalidComboboxMessage"
	case 1423:
		return "ErrorWindowNotCombobox"
	case 1424:
		return "ErrorInvalidEditHeight"
	case 1425:
		return "ErrorDcNotFound"
	case 1426:
		return "ErrorInvalidHookFilter"
	case 1427:
		return "ErrorInvalidFilterProc"
	case 1428:
		return "ErrorHookNeedsHmod"
	case 1429:
		return "ErrorGlobalOnlyHook"
	case 1430:
		return "ErrorJournalHookSet"
	case 1431:
		return "ErrorHookNotInstalled"
	case 1432:
		return "ErrorInvalidLbMessage"
	case 1433:
		return "ErrorSetcountOnBadLb"
	case 1434:
		return "ErrorLbWithoutTabstops"
	case 1435:
		return "ErrorDestroyObjectOfOtherThread"
	case 1436:
		return "ErrorChildWindowMenu"
	case 1437:
		return "ErrorNoSystemMenu"
	case 1438:
		return "ErrorInvalidMsgboxStyle"
	case 1439:
		return "ErrorInvalidSpiValue"
	case 1440:
		return "ErrorScreenAlreadyLocked"
	case 1441:
		return "ErrorHwndsHaveDiffParent"
	case 1442:
		return "ErrorNotChildWindow"
	case 1443:
		return "ErrorInvalidGwCommand"
	case 1444:
		return "ErrorInvalidThreadId"
	case 1445:
		return "ErrorNonMdichildWindow"
	case 1446:
		return "ErrorPopupAlreadyActive"
	case 1447:
		return "ErrorNoScrollbars"
	case 1448:
		return "ErrorInvalidScrollbarRange"
	case 1449:
		return "ErrorInvalidShowwinCommand"
	case 1450:
		return "ErrorNoSystemResources"
	case 1451:
		return "ErrorNonpagedSystemResources"
	case 1452:
		return "ErrorPagedSystemResources"
	case 1453:
		return "ErrorWorkingSetQuota"
	case 1454:
		return "ErrorPagefileQuota"
	case 1455:
		return "ErrorCommitmentLimit"
	case 1456:
		return "ErrorMenuItemNotFound"
	case 1457:
		return "ErrorInvalidKeyboardHandle"
	case 1458:
		return "ErrorHookTypeNotAllowed"
	case 1459:
		return "ErrorRequiresInteractiveWindowstation"
	case 1460:
		return "ErrorTimeout"
	case 1461:
		return "ErrorInvalidMonitorHandle"
	case 1462:
		return "ErrorIncorrectSize"
	case 1463:
		return "ErrorSymlinkClassDisabled"
	case 1464:
		return "ErrorSymlinkNotSupported"
	case 1465:
		return "ErrorXmlParseError"
	case 1466:
		return "ErrorXmldsigError"
	case 1467:
		return "ErrorRestartApplication"
	case 1468:
		return "ErrorWrongCompartment"
	case 1469:
		return "ErrorAuthipFailure"
	case 1470:
		return "ErrorNoNvramResources"
	case 1500:
		return "ErrorEventlogFileCorrupt"
	case 1501:
		return "ErrorEventlogCantStart"
	case 1502:
		return "ErrorLogFileFull"
	case 1503:
		return "ErrorEventlogFileChanged"
	case 1550:
		return "ErrorInvalidTaskName"
	case 1551:
		return "ErrorInvalidTaskIndex"
	case 1552:
		return "ErrorThreadAlreadyInTask"
	case 1601:
		return "ErrorInstallServiceFailure"
	case 1602:
		return "ErrorInstallUserexit"
	case 1603:
		return "ErrorInstallFailure"
	case 1604:
		return "ErrorInstallSuspend"
	case 1605:
		return "ErrorUnknownProduct"
	case 1606:
		return "ErrorUnknownFeature"
	case 1607:
		return "ErrorUnknownComponent"
	case 1608:
		return "ErrorUnknownProperty"
	case 1609:
		return "ErrorInvalidHandleState"
	case 1610:
		return "ErrorBadConfiguration"
	case 1611:
		return "ErrorIndexAbsent"
	case 1612:
		return "ErrorInstallSourceAbsent"
	case 1613:
		return "ErrorInstallPackageVersion"
	case 1614:
		return "ErrorProductUninstalled"
	case 1615:
		return "ErrorBadQuerySyntax"
	case 1616:
		return "ErrorInvalidField"
	case 1617:
		return "ErrorDeviceRemoved"
	case 1618:
		return "ErrorInstallAlreadyRunning"
	case 1619:
		return "ErrorInstallPackageOpenFailed"
	case 1620:
		return "ErrorInstallPackageInvalid"
	case 1621:
		return "ErrorInstallUiFailure"
	case 1622:
		return "ErrorInstallLogFailure"
	case 1623:
		return "ErrorInstallLanguageUnsupported"
	case 1624:
		return "ErrorInstallTransformFailure"
	case 1625:
		return "ErrorInstallPackageRejected"
	case 1626:
		return "ErrorFunctionNotCalled"
	case 1627:
		return "ErrorFunctionFailed"
	case 1628:
		return "ErrorInvalidTable"
	case 1629:
		return "ErrorDatatypeMismatch"
	case 1630:
		return "ErrorUnsupportedType"
	case 1631:
		return "ErrorCreateFailed"
	case 1632:
		return "ErrorInstallTempUnwritable"
	case 1633:
		return "ErrorInstallPlatformUnsupported"
	case 1634:
		return "ErrorInstallNotused"
	case 1635:
		return "ErrorPatchPackageOpenFailed"
	case 1636:
		return "ErrorPatchPackageInvalid"
	case 1637:
		return "ErrorPatchPackageUnsupported"
	case 1638:
		return "ErrorProductVersion"
	case 1639:
		return "ErrorInvalidCommandLine"
	case 1640:
		return "ErrorInstallRemoteDisallowed"
	case 1641:
		return "ErrorSuccessRebootInitiated"
	case 1642:
		return "ErrorPatchTargetNotFound"
	case 1643:
		return "ErrorPatchPackageRejected"
	case 1644:
		return "ErrorInstallTransformRejected"
	case 1645:
		return "ErrorInstallRemoteProhibited"
	case 1646:
		return "ErrorPatchRemovalUnsupported"
	case 1647:
		return "ErrorUnknownPatch"
	case 1648:
		return "ErrorPatchNoSequence"
	case 1649:
		return "ErrorPatchRemovalDisallowed"
	case 1650:
		return "ErrorInvalidPatchXml"
	case 1651:
		return "ErrorPatchManagedAdvertisedProduct"
	case 1652:
		return "ErrorInstallServiceSafeboot"
	case 1653:
		return "ErrorFailFastException"
	case 1700:
		return "RpcSInvalidStringBinding"
	case 1701:
		return "RpcSWrongKindOfBinding"
	case 1702:
		return "RpcSInvalidBinding"
	case 1703:
		return "RpcSProtseqNotSupported"
	case 1704:
		return "RpcSInvalidRpcProtseq"
	case 1705:
		return "RpcSInvalidStringUuid"
	case 1706:
		return "RpcSInvalidEndpointFormat"
	case 1707:
		return "RpcSInvalidNetAddr"
	case 1708:
		return "RpcSNoEndpointFound"
	case 1709:
		return "RpcSInvalidTimeout"
	case 1710:
		return "RpcSObjectNotFound"
	case 1711:
		return "RpcSAlreadyRegistered"
	case 1712:
		return "RpcSTypeAlreadyRegistered"
	case 1713:
		return "RpcSAlreadyListening"
	case 1714:
		return "RpcSNoProtseqsRegistered"
	case 1715:
		return "RpcSNotListening"
	case 1716:
		return "RpcSUnknownMgrType"
	case 1717:
		return "RpcSUnknownIf"
	case 1718:
		return "RpcSNoBindings"
	case 1719:
		return "RpcSNoProtseqs"
	case 1720:
		return "RpcSCantCreateEndpoint"
	case 1721:
		return "RpcSOutOfResources"
	case 1722:
		return "RpcSServerUnavailable"
	case 1723:
		return "RpcSServerTooBusy"
	case 1724:
		return "RpcSInvalidNetworkOptions"
	case 1725:
		return "RpcSNoCallActive"
	case 1726:
		return "RpcSCallFailed"
	case 1727:
		return "RpcSCallFailedDne"
	case 1728:
		return "RpcSProtocolError"
	case 1729:
		return "RpcSProxyAccessDenied"
	case 1730:
		return "RpcSUnsupportedTransSyn"
	case 1732:
		return "RpcSUnsupportedType"
	case 1733:
		return "RpcSInvalidTag"
	case 1734:
		return "RpcSInvalidBound"
	case 1735:
		return "RpcSNoEntryName"
	case 1736:
		return "RpcSInvalidNameSyntax"
	case 1737:
		return "RpcSUnsupportedNameSyntax"
	case 1739:
		return "RpcSUuidNoAddress"
	case 1740:
		return "RpcSDuplicateEndpoint"
	case 1741:
		return "RpcSUnknownAuthnType"
	case 1742:
		return "RpcSMaxCallsTooSmall"
	case 1743:
		return "RpcSStringTooLong"
	case 1744:
		return "RpcSProtseqNotFound"
	case 1745:
		return "RpcSProcnumOutOfRange"
	case 1746:
		return "RpcSBindingHasNoAuth"
	case 1747:
		return "RpcSUnknownAuthnService"
	case 1748:
		return "RpcSUnknownAuthnLevel"
	case 1749:
		return "RpcSInvalidAuthIdentity"
	case 1750:
		return "RpcSUnknownAuthzService"
	case 1751:
		return "EptSInvalidEntry"
	case 1752:
		return "EptSCantPerformOp"
	case 1753:
		return "EptSNotRegistered"
	case 1754:
		return "RpcSNothingToExport"
	case 1755:
		return "RpcSIncompleteName"
	case 1756:
		return "RpcSInvalidVersOption"
	case 1757:
		return "RpcSNoMoreMembers"
	case 1758:
		return "RpcSNotAllObjsUnexported"
	case 1759:
		return "RpcSInterfaceNotFound"
	case 1760:
		return "RpcSEntryAlreadyExists"
	case 1761:
		return "RpcSEntryNotFound"
	case 1762:
		return "RpcSNameServiceUnavailable"
	case 1763:
		return "RpcSInvalidNafId"
	case 1764:
		return "RpcSCannotSupport"
	case 1765:
		return "RpcSNoContextAvailable"
	case 1766:
		return "RpcSInternalError"
	case 1767:
		return "RpcSZeroDivide"
	case 1768:
		return "RpcSAddressError"
	case 1769:
		return "RpcSFpDivZero"
	case 1770:
		return "RpcSFpUnderflow"
	case 1771:
		return "RpcSFpOverflow"
	case 1772:
		return "RpcXNoMoreEntries"
	case 1773:
		return "RpcXSsCharTransOpenFail"
	case 1774:
		return "RpcXSsCharTransShortFile"
	case 1775:
		return "RpcXSsInNullContext"
	case 1777:
		return "RpcXSsContextDamaged"
	case 1778:
		return "RpcXSsHandlesMismatch"
	case 1779:
		return "RpcXSsCannotGetCallHandle"
	case 1780:
		return "RpcXNullRefPointer"
	case 1781:
		return "RpcXEnumValueOutOfRange"
	case 1782:
		return "RpcXByteCountTooSmall"
	case 1783:
		return "RpcXBadStubData"
	case 1784:
		return "ErrorInvalidUserBuffer"
	case 1785:
		return "ErrorUnrecognizedMedia"
	case 1786:
		return "ErrorNoTrustLsaSecret"
	case 1787:
		return "ErrorNoTrustSamAccount"
	case 1788:
		return "ErrorTrustedDomainFailure"
	case 1789:
		return "ErrorTrustedRelationshipFailure"
	case 1790:
		return "ErrorTrustFailure"
	case 1791:
		return "RpcSCallInProgress"
	case 1792:
		return "ErrorNetlogonNotStarted"
	case 1793:
		return "ErrorAccountExpired"
	case 1794:
		return "ErrorRedirectorHasOpenHandles"
	case 1795:
		return "ErrorPrinterDriverAlreadyInstalled"
	case 1796:
		return "ErrorUnknownPort"
	case 1797:
		return "ErrorUnknownPrinterDriver"
	case 1798:
		return "ErrorUnknownPrintprocessor"
	case 1799:
		return "ErrorInvalidSeparatorFile"
	case 1800:
		return "ErrorInvalidPriority"
	case 1801:
		return "ErrorInvalidPrinterName"
	case 1802:
		return "ErrorPrinterAlreadyExists"
	case 1803:
		return "ErrorInvalidPrinterCommand"
	case 1804:
		return "ErrorInvalidDatatype"
	case 1805:
		return "ErrorInvalidEnvironment"
	case 1806:
		return "RpcSNoMoreBindings"
	case 1807:
		return "ErrorNologonInterdomainTrustAccount"
	case 1808:
		return "ErrorNologonWorkstationTrustAccount"
	case 1809:
		return "ErrorNologonServerTrustAccount"
	case 1810:
		return "ErrorDomainTrustInconsistent"
	case 1811:
		return "ErrorServerHasOpenHandles"
	case 1812:
		return "ErrorResourceDataNotFound"
	case 1813:
		return "ErrorResourceTypeNotFound"
	case 1814:
		return "ErrorResourceNameNotFound"
	case 1815:
		return "ErrorResourceLangNotFound"
	case 1816:
		return "ErrorNotEnoughQuota"
	case 1817:
		return "RpcSNoInterfaces"
	case 1818:
		return "RpcSCallCancelled"
	case 1819:
		return "RpcSBindingIncomplete"
	case 1820:
		return "RpcSCommFailure"
	case 1821:
		return "RpcSUnsupportedAuthnLevel"
	case 1822:
		return "RpcSNoPrincName"
	case 1823:
		return "RpcSNotRpcError"
	case 1824:
		return "RpcSUuidLocalOnly"
	case 1825:
		return "RpcSSecPkgError"
	case 1826:
		return "RpcSNotCancelled"
	case 1827:
		return "RpcXInvalidEsAction"
	case 1828:
		return "RpcXWrongEsVersion"
	case 1829:
		return "RpcXWrongStubVersion"
	case 1830:
		return "RpcXInvalidPipeObject"
	case 1831:
		return "RpcXWrongPipeOrder"
	case 1832:
		return "RpcXWrongPipeVersion"
	case 1833:
		return "RpcSCookieAuthFailed"
	case 1898:
		return "RpcSGroupMemberNotFound"
	case 1899:
		return "EptSCantCreate"
	case 1900:
		return "RpcSInvalidObject"
	case 1901:
		return "ErrorInvalidTime"
	case 1902:
		return "ErrorInvalidFormName"
	case 1903:
		return "ErrorInvalidFormSize"
	case 1904:
		return "ErrorAlreadyWaiting"
	case 1905:
		return "ErrorPrinterDeleted"
	case 1906:
		return "ErrorInvalidPrinterState"
	case 1907:
		return "ErrorPasswordMustChange"
	case 1908:
		return "ErrorDomainControllerNotFound"
	case 1909:
		return "ErrorAccountLockedOut"
	case 1910:
		return "OrInvalidOxid"
	case 1911:
		return "OrInvalidOid"
	case 1912:
		return "OrInvalidSet"
	case 1913:
		return "RpcSSendIncomplete"
	case 1914:
		return "RpcSInvalidAsyncHandle"
	case 1915:
		return "RpcSInvalidAsyncCall"
	case 1916:
		return "RpcXPipeClosed"
	case 1917:
		return "RpcXPipeDisciplineError"
	case 1918:
		return "RpcXPipeEmpty"
	case 1919:
		return "ErrorNoSitename"
	case 1920:
		return "ErrorCantAccessFile"
	case 1921:
		return "ErrorCantResolveFilename"
	case 1922:
		return "RpcSEntryTypeMismatch"
	case 1923:
		return "RpcSNotAllObjsExported"
	case 1924:
		return "RpcSInterfaceNotExported"
	case 1925:
		return "RpcSProfileNotAdded"
	case 1926:
		return "RpcSPrfEltNotAdded"
	case 1927:
		return "RpcSPrfEltNotRemoved"
	case 1928:
		return "RpcSGrpEltNotAdded"
	case 1929:
		return "RpcSGrpEltNotRemoved"
	case 1930:
		return "ErrorKmDriverBlocked"
	case 1931:
		return "ErrorContextExpired"
	case 1932:
		return "ErrorPerUserTrustQuotaExceeded"
	case 1933:
		return "ErrorAllUserTrustQuotaExceeded"
	case 1934:
		return "ErrorUserDeleteTrustQuotaExceeded"
	case 1935:
		return "ErrorAuthenticationFirewallFailed"
	case 1936:
		return "ErrorRemotePrintConnectionsBlocked"
	case 1937:
		return "ErrorNtlmBlocked"
	case 2000:
		return "ErrorInvalidPixelFormat"
	case 2001:
		return "ErrorBadDriver"
	case 2002:
		return "ErrorInvalidWindowStyle"
	case 2003:
		return "ErrorMetafileNotSupported"
	case 2004:
		return "ErrorTransformNotSupported"
	case 2005:
		return "ErrorClippingNotSupported"
	case 2010:
		return "ErrorInvalidCmm"
	case 2011:
		return "ErrorInvalidProfile"
	case 2012:
		return "ErrorTagNotFound"
	case 2013:
		return "ErrorTagNotPresent"
	case 2014:
		return "ErrorDuplicateTag"
	case 2015:
		return "ErrorProfileNotAssociatedWithDevice"
	case 2016:
		return "ErrorProfileNotFound"
	case 2017:
		return "ErrorInvalidColorspace"
	case 2018:
		return "ErrorIcmNotEnabled"
	case 2019:
		return "ErrorDeletingIcmXform"
	case 2020:
		return "ErrorInvalidTransform"
	case 2021:
		return "ErrorColorspaceMismatch"
	case 2022:
		return "ErrorInvalidColorindex"
	case 2023:
		return "ErrorProfileDoesNotMatchDevice"
	case 2108:
		return "ErrorConnectedOtherPassword"
	case 2109:
		return "ErrorConnectedOtherPasswordDefault"
	case 2202:
		return "ErrorBadUsername"
	case 2250:
		return "ErrorNotConnected"
	case 2401:
		return "ErrorOpenFiles"
	case 2402:
		return "ErrorActiveConnections"
	case 2404:
		return "ErrorDeviceInUse"
	case 3000:
		return "ErrorUnknownPrintMonitor"
	case 3001:
		return "ErrorPrinterDriverInUse"
	case 3002:
		return "ErrorSpoolFileNotFound"
	case 3003:
		return "ErrorSplNoStartdoc"
	case 3004:
		return "ErrorSplNoAddjob"
	case 3005:
		return "ErrorPrintProcessorAlreadyInstalled"
	case 3006:
		return "ErrorPrintMonitorAlreadyInstalled"
	case 3007:
		return "ErrorInvalidPrintMonitor"
	case 3008:
		return "ErrorPrintMonitorInUse"
	case 3009:
		return "ErrorPrinterHasJobsQueued"
	case 3010:
		return "ErrorSuccessRebootRequired"
	case 3011:
		return "ErrorSuccessRestartRequired"
	case 3012:
		return "ErrorPrinterNotFound"
	case 3013:
		return "ErrorPrinterDriverWarned"
	case 3014:
		return "ErrorPrinterDriverBlocked"
	case 3015:
		return "ErrorPrinterDriverPackageInUse"
	case 3016:
		return "ErrorCoreDriverPackageNotFound"
	case 3017:
		return "ErrorFailRebootRequired"
	case 3018:
		return "ErrorFailRebootInitiated"
	case 3019:
		return "ErrorPrinterDriverDownloadNeeded"
	case 3020:
		return "ErrorPrintJobRestartRequired"
	case 3950:
		return "ErrorIoReissueAsCached"
	case 4000:
		return "ErrorWinsInternal"
	case 4001:
		return "ErrorCanNotDelLocalWins"
	case 4002:
		return "ErrorStaticInit"
	case 4003:
		return "ErrorIncBackup"
	case 4004:
		return "ErrorFullBackup"
	case 4005:
		return "ErrorRecNonExistent"
	case 4006:
		return "ErrorRplNotAllowed"
	case 4050:
		return "PeerdistErrorContentinfoVersionUnsupported"
	case 4051:
		return "PeerdistErrorCannotParseContentinfo"
	case 4052:
		return "PeerdistErrorMissingData"
	case 4053:
		return "PeerdistErrorNoMore"
	case 4054:
		return "PeerdistErrorNotInitialized"
	case 4055:
		return "PeerdistErrorAlreadyInitialized"
	case 4056:
		return "PeerdistErrorShutdownInProgress"
	case 4057:
		return "PeerdistErrorInvalidated"
	case 4058:
		return "PeerdistErrorAlreadyExists"
	case 4059:
		return "PeerdistErrorOperationNotfound"
	case 4060:
		return "PeerdistErrorAlreadyCompleted"
	case 4061:
		return "PeerdistErrorOutOfBounds"
	case 4062:
		return "PeerdistErrorVersionUnsupported"
	case 4063:
		return "PeerdistErrorInvalidConfiguration"
	case 4064:
		return "PeerdistErrorNotLicensed"
	case 4065:
		return "PeerdistErrorServiceUnavailable"
	case 4100:
		return "ErrorDhcpAddressConflict"
	case 4200:
		return "ErrorWmiGuidNotFound"
	case 4201:
		return "ErrorWmiInstanceNotFound"
	case 4202:
		return "ErrorWmiItemidNotFound"
	case 4203:
		return "ErrorWmiTryAgain"
	case 4204:
		return "ErrorWmiDpNotFound"
	case 4205:
		return "ErrorWmiUnresolvedInstanceRef"
	case 4206:
		return "ErrorWmiAlreadyEnabled"
	case 4207:
		return "ErrorWmiGuidDisconnected"
	case 4208:
		return "ErrorWmiServerUnavailable"
	case 4209:
		return "ErrorWmiDpFailed"
	case 4210:
		return "ErrorWmiInvalidMof"
	case 4211:
		return "ErrorWmiInvalidReginfo"
	case 4212:
		return "ErrorWmiAlreadyDisabled"
	case 4213:
		return "ErrorWmiReadOnly"
	case 4214:
		return "ErrorWmiSetFailure"
	case 4300:
		return "ErrorInvalidMedia"
	case 4301:
		return "ErrorInvalidLibrary"
	case 4302:
		return "ErrorInvalidMediaPool"
	case 4303:
		return "ErrorDriveMediaMismatch"
	case 4304:
		return "ErrorMediaOffline"
	case 4305:
		return "ErrorLibraryOffline"
	case 4306:
		return "ErrorEmpty"
	case 4307:
		return "ErrorNotEmpty"
	case 4308:
		return "ErrorMediaUnavailable"
	case 4309:
		return "ErrorResourceDisabled"
	case 4310:
		return "ErrorInvalidCleaner"
	case 4311:
		return "ErrorUnableToClean"
	case 4312:
		return "ErrorObjectNotFound"
	case 4313:
		return "ErrorDatabaseFailure"
	case 4314:
		return "ErrorDatabaseFull"
	case 4315:
		return "ErrorMediaIncompatible"
	case 4316:
		return "ErrorResourceNotPresent"
	case 4317:
		return "ErrorInvalidOperation"
	case 4318:
		return "ErrorMediaNotAvailable"
	case 4319:
		return "ErrorDeviceNotAvailable"
	case 4320:
		return "ErrorRequestRefused"
	case 4321:
		return "ErrorInvalidDriveObject"
	case 4322:
		return "ErrorLibraryFull"
	case 4323:
		return "ErrorMediumNotAccessible"
	case 4324:
		return "ErrorUnableToLoadMedium"
	case 4325:
		return "ErrorUnableToInventoryDrive"
	case 4326:
		return "ErrorUnableToInventorySlot"
	case 4327:
		return "ErrorUnableToInventoryTransport"
	case 4328:
		return "ErrorTransportFull"
	case 4329:
		return "ErrorControllingIeport"
	case 4330:
		return "ErrorUnableToEjectMountedMedia"
	case 4331:
		return "ErrorCleanerSlotSet"
	case 4332:
		return "ErrorCleanerSlotNotSet"
	case 4333:
		return "ErrorCleanerCartridgeSpent"
	case 4334:
		return "ErrorUnexpectedOmid"
	case 4335:
		return "ErrorCantDeleteLastItem"
	case 4336:
		return "ErrorMessageExceedsMaxSize"
	case 4337:
		return "ErrorVolumeContainsSysFiles"
	case 4338:
		return "ErrorIndigenousType"
	case 4339:
		return "ErrorNoSupportingDrives"
	case 4340:
		return "ErrorCleanerCartridgeInstalled"
	case 4341:
		return "ErrorIeportFull"
	case 4350:
		return "ErrorFileOffline"
	case 4351:
		return "ErrorRemoteStorageNotActive"
	case 4352:
		return "ErrorRemoteStorageMediaError"
	case 4390:
		return "ErrorNotAReparsePoint"
	case 4391:
		return "ErrorReparseAttributeConflict"
	case 4392:
		return "ErrorInvalidReparseData"
	case 4393:
		return "ErrorReparseTagInvalid"
	case 4394:
		return "ErrorReparseTagMismatch"
	case 4500:
		return "ErrorVolumeNotSisEnabled"
	case 5001:
		return "ErrorDependentResourceExists"
	case 5002:
		return "ErrorDependencyNotFound"
	case 5003:
		return "ErrorDependencyAlreadyExists"
	case 5004:
		return "ErrorResourceNotOnline"
	case 5005:
		return "ErrorHostNodeNotAvailable"
	case 5006:
		return "ErrorResourceNotAvailable"
	case 5007:
		return "ErrorResourceNotFound"
	case 5008:
		return "ErrorShutdownCluster"
	case 5009:
		return "ErrorCantEvictActiveNode"
	case 5010:
		return "ErrorObjectAlreadyExists"
	case 5011:
		return "ErrorObjectInList"
	case 5012:
		return "ErrorGroupNotAvailable"
	case 5013:
		return "ErrorGroupNotFound"
	case 5014:
		return "ErrorGroupNotOnline"
	case 5015:
		return "ErrorHostNodeNotResourceOwner"
	case 5016:
		return "ErrorHostNodeNotGroupOwner"
	case 5017:
		return "ErrorResmonCreateFailed"
	case 5018:
		return "ErrorResmonOnlineFailed"
	case 5019:
		return "ErrorResourceOnline"
	case 5020:
		return "ErrorQuorumResource"
	case 5021:
		return "ErrorNotQuorumCapable"
	case 5022:
		return "ErrorClusterShuttingDown"
	case 5023:
		return "ErrorInvalidState"
	case 5024:
		return "ErrorResourcePropertiesStored"
	case 5025:
		return "ErrorNotQuorumClass"
	case 5026:
		return "ErrorCoreResource"
	case 5027:
		return "ErrorQuorumResourceOnlineFailed"
	case 5028:
		return "ErrorQuorumlogOpenFailed"
	case 5029:
		return "ErrorClusterlogCorrupt"
	case 5030:
		return "ErrorClusterlogRecordExceedsMaxsize"
	case 5031:
		return "ErrorClusterlogExceedsMaxsize"
	case 5032:
		return "ErrorClusterlogChkpointNotFound"
	case 5033:
		return "ErrorClusterlogNotEnoughSpace"
	case 5034:
		return "ErrorQuorumOwnerAlive"
	case 5035:
		return "ErrorNetworkNotAvailable"
	case 5036:
		return "ErrorNodeNotAvailable"
	case 5037:
		return "ErrorAllNodesNotAvailable"
	case 5038:
		return "ErrorResourceFailed"
	case 5039:
		return "ErrorClusterInvalidNode"
	case 5040:
		return "ErrorClusterNodeExists"
	case 5041:
		return "ErrorClusterJoinInProgress"
	case 5042:
		return "ErrorClusterNodeNotFound"
	case 5043:
		return "ErrorClusterLocalNodeNotFound"
	case 5044:
		return "ErrorClusterNetworkExists"
	case 5045:
		return "ErrorClusterNetworkNotFound"
	case 5046:
		return "ErrorClusterNetinterfaceExists"
	case 5047:
		return "ErrorClusterNetinterfaceNotFound"
	case 5048:
		return "ErrorClusterInvalidRequest"
	case 5049:
		return "ErrorClusterInvalidNetworkProvider"
	case 5050:
		return "ErrorClusterNodeDown"
	case 5051:
		return "ErrorClusterNodeUnreachable"
	case 5052:
		return "ErrorClusterNodeNotMember"
	case 5053:
		return "ErrorClusterJoinNotInProgress"
	case 5054:
		return "ErrorClusterInvalidNetwork"
	case 5056:
		return "ErrorClusterNodeUp"
	case 5057:
		return "ErrorClusterIpaddrInUse"
	case 5058:
		return "ErrorClusterNodeNotPaused"
	case 5059:
		return "ErrorClusterNoSecurityContext"
	case 5060:
		return "ErrorClusterNetworkNotInternal"
	case 5061:
		return "ErrorClusterNodeAlreadyUp"
	case 5062:
		return "ErrorClusterNodeAlreadyDown"
	case 5063:
		return "ErrorClusterNetworkAlreadyOnline"
	case 5064:
		return "ErrorClusterNetworkAlreadyOffline"
	case 5065:
		return "ErrorClusterNodeAlreadyMember"
	case 5066:
		return "ErrorClusterLastInternalNetwork"
	case 5067:
		return "ErrorClusterNetworkHasDependents"
	case 5068:
		return "ErrorInvalidOperationOnQuorum"
	case 5069:
		return "ErrorDependencyNotAllowed"
	case 5070:
		return "ErrorClusterNodePaused"
	case 5071:
		return "ErrorNodeCantHostResource"
	case 5072:
		return "ErrorClusterNodeNotReady"
	case 5073:
		return "ErrorClusterNodeShuttingDown"
	case 5074:
		return "ErrorClusterJoinAborted"
	case 5075:
		return "ErrorClusterIncompatibleVersions"
	case 5076:
		return "ErrorClusterMaxnumOfResourcesExceeded"
	case 5077:
		return "ErrorClusterSystemConfigChanged"
	case 5078:
		return "ErrorClusterResourceTypeNotFound"
	case 5079:
		return "ErrorClusterRestypeNotSupported"
	case 5080:
		return "ErrorClusterResnameNotFound"
	case 5081:
		return "ErrorClusterNoRpcPackagesRegistered"
	case 5082:
		return "ErrorClusterOwnerNotInPreflist"
	case 5083:
		return "ErrorClusterDatabaseSeqmismatch"
	case 5084:
		return "ErrorResmonInvalidState"
	case 5085:
		return "ErrorClusterGumNotLocker"
	case 5086:
		return "ErrorQuorumDiskNotFound"
	case 5087:
		return "ErrorDatabaseBackupCorrupt"
	case 5088:
		return "ErrorClusterNodeAlreadyHasDfsRoot"
	case 5089:
		return "ErrorResourcePropertyUnchangeable"
	case 5890:
		return "ErrorClusterMembershipInvalidState"
	case 5891:
		return "ErrorClusterQuorumlogNotFound"
	case 5892:
		return "ErrorClusterMembershipHalt"
	case 5893:
		return "ErrorClusterInstanceIdMismatch"
	case 5894:
		return "ErrorClusterNetworkNotFoundForIp"
	case 5895:
		return "ErrorClusterPropertyDataTypeMismatch"
	case 5896:
		return "ErrorClusterEvictWithoutCleanup"
	case 5897:
		return "ErrorClusterParameterMismatch"
	case 5898:
		return "ErrorNodeCannotBeClustered"
	case 5899:
		return "ErrorClusterWrongOsVersion"
	case 5900:
		return "ErrorClusterCantCreateDupClusterName"
	case 5901:
		return "ErrorCluscfgAlreadyCommitted"
	case 5902:
		return "ErrorCluscfgRollbackFailed"
	case 5903:
		return "ErrorCluscfgSystemDiskDriveLetterConflict"
	case 5904:
		return "ErrorClusterOldVersion"
	case 5905:
		return "ErrorClusterMismatchedComputerAcctName"
	case 5906:
		return "ErrorClusterNoNetAdapters"
	case 5907:
		return "ErrorClusterPoisoned"
	case 5908:
		return "ErrorClusterGroupMoving"
	case 5909:
		return "ErrorClusterResourceTypeBusy"
	case 5910:
		return "ErrorResourceCallTimedOut"
	case 5911:
		return "ErrorInvalidClusterIpv6Address"
	case 5912:
		return "ErrorClusterInternalInvalidFunction"
	case 5913:
		return "ErrorClusterParameterOutOfBounds"
	case 5914:
		return "ErrorClusterPartialSend"
	case 5915:
		return "ErrorClusterRegistryInvalidFunction"
	case 5916:
		return "ErrorClusterInvalidStringTermination"
	case 5917:
		return "ErrorClusterInvalidStringFormat"
	case 5918:
		return "ErrorClusterDatabaseTransactionInProgress"
	case 5919:
		return "ErrorClusterDatabaseTransactionNotInProgress"
	case 5920:
		return "ErrorClusterNullData"
	case 5921:
		return "ErrorClusterPartialRead"
	case 5922:
		return "ErrorClusterPartialWrite"
	case 5923:
		return "ErrorClusterCantDeserializeData"
	case 5924:
		return "ErrorDependentResourcePropertyConflict"
	case 5925:
		return "ErrorClusterNoQuorum"
	case 5926:
		return "ErrorClusterInvalidIpv6Network"
	case 5927:
		return "ErrorClusterInvalidIpv6TunnelNetwork"
	case 5928:
		return "ErrorQuorumNotAllowedInThisGroup"
	case 5929:
		return "ErrorDependencyTreeTooComplex"
	case 5930:
		return "ErrorExceptionInResourceCall"
	case 5931:
		return "ErrorClusterRhsFailedInitialization"
	case 5932:
		return "ErrorClusterNotInstalled"
	case 5933:
		return "ErrorClusterResourcesMustBeOnlineOnTheSameNode"
	case 5934:
		return "ErrorClusterMaxNodesInCluster"
	case 5935:
		return "ErrorClusterTooManyNodes"
	case 5936:
		return "ErrorClusterObjectAlreadyUsed"
	case 5937:
		return "ErrorNoncoreGroupsFound"
	case 5938:
		return "ErrorFileShareResourceConflict"
	case 5939:
		return "ErrorClusterEvictInvalidRequest"
	case 5940:
		return "ErrorClusterSingletonResource"
	case 5941:
		return "ErrorClusterGroupSingletonResource"
	case 5942:
		return "ErrorClusterResourceProviderFailed"
	case 5943:
		return "ErrorClusterResourceConfigurationError"
	case 5944:
		return "ErrorClusterGroupBusy"
	case 5945:
		return "ErrorClusterNotSharedVolume"
	case 5946:
		return "ErrorClusterInvalidSecurityDescriptor"
	case 5947:
		return "ErrorClusterSharedVolumesInUse"
	case 5948:
		return "ErrorClusterUseSharedVolumesApi"
	case 5949:
		return "ErrorClusterBackupInProgress"
	case 5950:
		return "ErrorNonCsvPath"
	case 5951:
		return "ErrorCsvVolumeNotLocal"
	case 5952:
		return "ErrorClusterWatchdogTerminating"
	case 6000:
		return "ErrorEncryptionFailed"
	case 6001:
		return "ErrorDecryptionFailed"
	case 6002:
		return "ErrorFileEncrypted"
	case 6003:
		return "ErrorNoRecoveryPolicy"
	case 6004:
		return "ErrorNoEfs"
	case 6005:
		return "ErrorWrongEfs"
	case 6006:
		return "ErrorNoUserKeys"
	case 6007:
		return "ErrorFileNotEncrypted"
	case 6008:
		return "ErrorNotExportFormat"
	case 6009:
		return "ErrorFileReadOnly"
	case 6010:
		return "ErrorDirEfsDisallowed"
	case 6011:
		return "ErrorEfsServerNotTrusted"
	case 6012:
		return "ErrorBadRecoveryPolicy"
	case 6013:
		return "ErrorEfsAlgBlobTooBig"
	case 6014:
		return "ErrorVolumeNotSupportEfs"
	case 6015:
		return "ErrorEfsDisabled"
	case 6016:
		return "ErrorEfsVersionNotSupport"
	case 6017:
		return "ErrorCsEncryptionInvalidServerResponse"
	case 6018:
		return "ErrorCsEncryptionUnsupportedServer"
	case 6019:
		return "ErrorCsEncryptionExistingEncryptedFile"
	case 6020:
		return "ErrorCsEncryptionNewEncryptedFile"
	case 6021:
		return "ErrorCsEncryptionFileNotCse"
	case 6022:
		return "ErrorEncryptionPolicyDeniesOperation"
	case 6118:
		return "ErrorNoBrowserServersFound"
	case 6200:
		return "SchedEServiceNotLocalsystem"
	case 6600:
		return "ErrorLogSectorInvalid"
	case 6601:
		return "ErrorLogSectorParityInvalid"
	case 6602:
		return "ErrorLogSectorRemapped"
	case 6603:
		return "ErrorLogBlockIncomplete"
	case 6604:
		return "ErrorLogInvalidRange"
	case 6605:
		return "ErrorLogBlocksExhausted"
	case 6606:
		return "ErrorLogReadContextInvalid"
	case 6607:
		return "ErrorLogRestartInvalid"
	case 6608:
		return "ErrorLogBlockVersion"
	case 6609:
		return "ErrorLogBlockInvalid"
	case 6610:
		return "ErrorLogReadModeInvalid"
	case 6611:
		return "ErrorLogNoRestart"
	case 6612:
		return "ErrorLogMetadataCorrupt"
	case 6613:
		return "ErrorLogMetadataInvalid"
	case 6614:
		return "ErrorLogMetadataInconsistent"
	case 6615:
		return "ErrorLogReservationInvalid"
	case 6616:
		return "ErrorLogCantDelete"
	case 6617:
		return "ErrorLogContainerLimitExceeded"
	case 6618:
		return "ErrorLogStartOfLog"
	case 6619:
		return "ErrorLogPolicyAlreadyInstalled"
	case 6620:
		return "ErrorLogPolicyNotInstalled"
	case 6621:
		return "ErrorLogPolicyInvalid"
	case 6622:
		return "ErrorLogPolicyConflict"
	case 6623:
		return "ErrorLogPinnedArchiveTail"
	case 6624:
		return "ErrorLogRecordNonexistent"
	case 6625:
		return "ErrorLogRecordsReservedInvalid"
	case 6626:
		return "ErrorLogSpaceReservedInvalid"
	case 6627:
		return "ErrorLogTailInvalid"
	case 6628:
		return "ErrorLogFull"
	case 6629:
		return "ErrorCouldNotResizeLog"
	case 6630:
		return "ErrorLogMultiplexed"
	case 6631:
		return "ErrorLogDedicated"
	case 6632:
		return "ErrorLogArchiveNotInProgress"
	case 6633:
		return "ErrorLogArchiveInProgress"
	case 6634:
		return "ErrorLogEphemeral"
	case 6635:
		return "ErrorLogNotEnoughContainers"
	case 6636:
		return "ErrorLogClientAlreadyRegistered"
	case 6637:
		return "ErrorLogClientNotRegistered"
	case 6638:
		return "ErrorLogFullHandlerInProgress"
	case 6639:
		return "ErrorLogContainerReadFailed"
	case 6640:
		return "ErrorLogContainerWriteFailed"
	case 6641:
		return "ErrorLogContainerOpenFailed"
	case 6642:
		return "ErrorLogContainerStateInvalid"
	case 6643:
		return "ErrorLogStateInvalid"
	case 6644:
		return "ErrorLogPinned"
	case 6645:
		return "ErrorLogMetadataFlushFailed"
	case 6646:
		return "ErrorLogInconsistentSecurity"
	case 6647:
		return "ErrorLogAppendedFlushFailed"
	case 6648:
		return "ErrorLogPinnedReservation"
	case 6700:
		return "ErrorInvalidTransaction"
	case 6701:
		return "ErrorTransactionNotActive"
	case 6702:
		return "ErrorTransactionRequestNotValid"
	case 6703:
		return "ErrorTransactionNotRequested"
	case 6704:
		return "ErrorTransactionAlreadyAborted"
	case 6705:
		return "ErrorTransactionAlreadyCommitted"
	case 6706:
		return "ErrorTmInitializationFailed"
	case 6707:
		return "ErrorResourcemanagerReadOnly"
	case 6708:
		return "ErrorTransactionNotJoined"
	case 6709:
		return "ErrorTransactionSuperiorExists"
	case 6710:
		return "ErrorCrmProtocolAlreadyExists"
	case 6711:
		return "ErrorTransactionPropagationFailed"
	case 6712:
		return "ErrorCrmProtocolNotFound"
	case 6713:
		return "ErrorTransactionInvalidMarshallBuffer"
	case 6714:
		return "ErrorCurrentTransactionNotValid"
	case 6715:
		return "ErrorTransactionNotFound"
	case 6716:
		return "ErrorResourcemanagerNotFound"
	case 6717:
		return "ErrorEnlistmentNotFound"
	case 6718:
		return "ErrorTransactionmanagerNotFound"
	case 6719:
		return "ErrorTransactionmanagerNotOnline"
	case 6720:
		return "ErrorTransactionmanagerRecoveryNameCollision"
	case 6721:
		return "ErrorTransactionNotRoot"
	case 6722:
		return "ErrorTransactionObjectExpired"
	case 6723:
		return "ErrorTransactionResponseNotEnlisted"
	case 6724:
		return "ErrorTransactionRecordTooLong"
	case 6725:
		return "ErrorImplicitTransactionNotSupported"
	case 6726:
		return "ErrorTransactionIntegrityViolated"
	case 6727:
		return "ErrorTransactionmanagerIdentityMismatch"
	case 6728:
		return "ErrorRmCannotBeFrozenForSnapshot"
	case 6729:
		return "ErrorTransactionMustWritethrough"
	case 6730:
		return "ErrorTransactionNoSuperior"
	case 6731:
		return "ErrorHeuristicDamagePossible"
	case 6800:
		return "ErrorTransactionalConflict"
	case 6801:
		return "ErrorRmNotActive"
	case 6802:
		return "ErrorRmMetadataCorrupt"
	case 6803:
		return "ErrorDirectoryNotRm"
	case 6805:
		return "ErrorTransactionsUnsupportedRemote"
	case 6806:
		return "ErrorLogResizeInvalidSize"
	case 6807:
		return "ErrorObjectNoLongerExists"
	case 6808:
		return "ErrorStreamMiniversionNotFound"
	case 6809:
		return "ErrorStreamMiniversionNotValid"
	case 6810:
		return "ErrorMiniversionInaccessibleFromSpecifiedTransaction"
	case 6811:
		return "ErrorCantOpenMiniversionWithModifyIntent"
	case 6812:
		return "ErrorCantCreateMoreStreamMiniversions"
	case 6814:
		return "ErrorRemoteFileVersionMismatch"
	case 6815:
		return "ErrorHandleNoLongerValid"
	case 6816:
		return "ErrorNoTxfMetadata"
	case 6817:
		return "ErrorLogCorruptionDetected"
	case 6818:
		return "ErrorCantRecoverWithHandleOpen"
	case 6819:
		return "ErrorRmDisconnected"
	case 6820:
		return "ErrorEnlistmentNotSuperior"
	case 6821:
		return "ErrorRecoveryNotNeeded"
	case 6822:
		return "ErrorRmAlreadyStarted"
	case 6823:
		return "ErrorFileIdentityNotPersistent"
	case 6824:
		return "ErrorCantBreakTransactionalDependency"
	case 6825:
		return "ErrorCantCrossRmBoundary"
	case 6826:
		return "ErrorTxfDirNotEmpty"
	case 6827:
		return "ErrorIndoubtTransactionsExist"
	case 6828:
		return "ErrorTmVolatile"
	case 6829:
		return "ErrorRollbackTimerExpired"
	case 6830:
		return "ErrorTxfAttributeCorrupt"
	case 6831:
		return "ErrorEfsNotAllowedInTransaction"
	case 6832:
		return "ErrorTransactionalOpenNotAllowed"
	case 6833:
		return "ErrorLogGrowthFailed"
	case 6834:
		return "ErrorTransactedMappingUnsupportedRemote"
	case 6835:
		return "ErrorTxfMetadataAlreadyPresent"
	case 6836:
		return "ErrorTransactionScopeCallbacksNotSet"
	case 6837:
		return "ErrorTransactionRequiredPromotion"
	case 6838:
		return "ErrorCannotExecuteFileInTransaction"
	case 6839:
		return "ErrorTransactionsNotFrozen"
	case 6840:
		return "ErrorTransactionFreezeInProgress"
	case 6841:
		return "ErrorNotSnapshotVolume"
	case 6842:
		return "ErrorNoSavepointWithOpenFiles"
	case 6843:
		return "ErrorDataLostRepair"
	case 6844:
		return "ErrorSparseNotAllowedInTransaction"
	case 6845:
		return "ErrorTmIdentityMismatch"
	case 6846:
		return "ErrorFloatedSection"
	case 6847:
		return "ErrorCannotAcceptTransactedWork"
	case 6848:
		return "ErrorCannotAbortTransactions"
	case 6849:
		return "ErrorBadClusters"
	case 6850:
		return "ErrorCompressionNotAllowedInTransaction"
	case 6851:
		return "ErrorVolumeDirty"
	case 6852:
		return "ErrorNoLinkTrackingInTransaction"
	case 6853:
		return "ErrorOperationNotSupportedInTransaction"
	case 6854:
		return "ErrorExpiredHandle"
	case 6855:
		return "ErrorTransactionNotEnlisted"
	case 7001:
		return "ErrorCtxWinstationNameInvalid"
	case 7002:
		return "ErrorCtxInvalidPd"
	case 7003:
		return "ErrorCtxPdNotFound"
	case 7004:
		return "ErrorCtxWdNotFound"
	case 7005:
		return "ErrorCtxCannotMakeEventlogEntry"
	case 7006:
		return "ErrorCtxServiceNameCollision"
	case 7007:
		return "ErrorCtxClosePending"
	case 7008:
		return "ErrorCtxNoOutbuf"
	case 7009:
		return "ErrorCtxModemInfNotFound"
	case 7010:
		return "ErrorCtxInvalidModemname"
	case 7011:
		return "ErrorCtxModemResponseError"
	case 7012:
		return "ErrorCtxModemResponseTimeout"
	case 7013:
		return "ErrorCtxModemResponseNoCarrier"
	case 7014:
		return "ErrorCtxModemResponseNoDialtone"
	case 7015:
		return "ErrorCtxModemResponseBusy"
	case 7016:
		return "ErrorCtxModemResponseVoice"
	case 7017:
		return "ErrorCtxTdError"
	case 7022:
		return "ErrorCtxWinstationNotFound"
	case 7023:
		return "ErrorCtxWinstationAlreadyExists"
	case 7024:
		return "ErrorCtxWinstationBusy"
	case 7025:
		return "ErrorCtxBadVideoMode"
	case 7035:
		return "ErrorCtxGraphicsInvalid"
	case 7037:
		return "ErrorCtxLogonDisabled"
	case 7038:
		return "ErrorCtxNotConsole"
	case 7040:
		return "ErrorCtxClientQueryTimeout"
	case 7041:
		return "ErrorCtxConsoleDisconnect"
	case 7042:
		return "ErrorCtxConsoleConnect"
	case 7044:
		return "ErrorCtxShadowDenied"
	case 7045:
		return "ErrorCtxWinstationAccessDenied"
	case 7049:
		return "ErrorCtxInvalidWd"
	case 7050:
		return "ErrorCtxShadowInvalid"
	case 7051:
		return "ErrorCtxShadowDisabled"
	case 7052:
		return "ErrorCtxClientLicenseInUse"
	case 7053:
		return "ErrorCtxClientLicenseNotSet"
	case 7054:
		return "ErrorCtxLicenseNotAvailable"
	case 7055:
		return "ErrorCtxLicenseClientInvalid"
	case 7056:
		return "ErrorCtxLicenseExpired"
	case 7057:
		return "ErrorCtxShadowNotRunning"
	case 7058:
		return "ErrorCtxShadowEndedByModeChange"
	case 7059:
		return "ErrorActivationCountExceeded"
	case 7060:
		return "ErrorCtxWinstationsDisabled"
	case 7061:
		return "ErrorCtxEncryptionLevelRequired"
	case 7062:
		return "ErrorCtxSessionInUse"
	case 7063:
		return "ErrorCtxNoForceLogoff"
	case 7064:
		return "ErrorCtxAccountRestriction"
	case 7065:
		return "ErrorRdpProtocolError"
	case 7066:
		return "ErrorCtxCdmConnect"
	case 7067:
		return "ErrorCtxCdmDisconnect"
	case 7068:
		return "ErrorCtxSecurityLayerError"
	case 7069:
		return "ErrorTsIncompatibleSessions"
	case 7070:
		return "ErrorTsVideoSubsystemError"
	case 8001:
		return "FrsErrInvalidApiSequence"
	case 8002:
		return "FrsErrStartingService"
	case 8003:
		return "FrsErrStoppingService"
	case 8004:
		return "FrsErrInternalApi"
	case 8005:
		return "FrsErrInternal"
	case 8006:
		return "FrsErrServiceComm"
	case 8007:
		return "FrsErrInsufficientPriv"
	case 8008:
		return "FrsErrAuthentication"
	case 8009:
		return "FrsErrParentInsufficientPriv"
	case 8010:
		return "FrsErrParentAuthentication"
	case 8011:
		return "FrsErrChildToParentComm"
	case 8012:
		return "FrsErrParentToChildComm"
	case 8013:
		return "FrsErrSysvolPopulate"
	case 8014:
		return "FrsErrSysvolPopulateTimeout"
	case 8015:
		return "FrsErrSysvolIsBusy"
	case 8016:
		return "FrsErrSysvolDemote"
	case 8017:
		return "FrsErrInvalidServiceParameter"
	case 8200:
		return "ErrorDsNotInstalled"
	case 8201:
		return "ErrorDsMembershipEvaluatedLocally"
	case 8202:
		return "ErrorDsNoAttributeOrValue"
	case 8203:
		return "ErrorDsInvalidAttributeSyntax"
	case 8204:
		return "ErrorDsAttributeTypeUndefined"
	case 8205:
		return "ErrorDsAttributeOrValueExists"
	case 8206:
		return "ErrorDsBusy"
	case 8207:
		return "ErrorDsUnavailable"
	case 8208:
		return "ErrorDsNoRidsAllocated"
	case 8209:
		return "ErrorDsNoMoreRids"
	case 8210:
		return "ErrorDsIncorrectRoleOwner"
	case 8211:
		return "ErrorDsRidmgrInitError"
	case 8212:
		return "ErrorDsObjClassViolation"
	case 8213:
		return "ErrorDsCantOnNonLeaf"
	case 8214:
		return "ErrorDsCantOnRdn"
	case 8215:
		return "ErrorDsCantModObjClass"
	case 8216:
		return "ErrorDsCrossDomMoveError"
	case 8217:
		return "ErrorDsGcNotAvailable"
	case 8218:
		return "ErrorSharedPolicy"
	case 8219:
		return "ErrorPolicyObjectNotFound"
	case 8220:
		return "ErrorPolicyOnlyInDs"
	case 8221:
		return "ErrorPromotionActive"
	case 8222:
		return "ErrorNoPromotionActive"
	case 8224:
		return "ErrorDsOperationsError"
	case 8225:
		return "ErrorDsProtocolError"
	case 8226:
		return "ErrorDsTimelimitExceeded"
	case 8227:
		return "ErrorDsSizelimitExceeded"
	case 8228:
		return "ErrorDsAdminLimitExceeded"
	case 8229:
		return "ErrorDsCompareFalse"
	case 8230:
		return "ErrorDsCompareTrue"
	case 8231:
		return "ErrorDsAuthMethodNotSupported"
	case 8232:
		return "ErrorDsStrongAuthRequired"
	case 8233:
		return "ErrorDsInappropriateAuth"
	case 8234:
		return "ErrorDsAuthUnknown"
	case 8235:
		return "ErrorDsReferral"
	case 8236:
		return "ErrorDsUnavailableCritExtension"
	case 8237:
		return "ErrorDsConfidentialityRequired"
	case 8238:
		return "ErrorDsInappropriateMatching"
	case 8239:
		return "ErrorDsConstraintViolation"
	case 8240:
		return "ErrorDsNoSuchObject"
	case 8241:
		return "ErrorDsAliasProblem"
	case 8242:
		return "ErrorDsInvalidDnSyntax"
	case 8243:
		return "ErrorDsIsLeaf"
	case 8244:
		return "ErrorDsAliasDerefProblem"
	case 8245:
		return "ErrorDsUnwillingToPerform"
	case 8246:
		return "ErrorDsLoopDetect"
	case 8247:
		return "ErrorDsNamingViolation"
	case 8248:
		return "ErrorDsObjectResultsTooLarge"
	case 8249:
		return "ErrorDsAffectsMultipleDsas"
	case 8250:
		return "ErrorDsServerDown"
	case 8251:
		return "ErrorDsLocalError"
	case 8252:
		return "ErrorDsEncodingError"
	case 8253:
		return "ErrorDsDecodingError"
	case 8254:
		return "ErrorDsFilterUnknown"
	case 8255:
		return "ErrorDsParamError"
	case 8256:
		return "ErrorDsNotSupported"
	case 8257:
		return "ErrorDsNoResultsReturned"
	case 8258:
		return "ErrorDsControlNotFound"
	case 8259:
		return "ErrorDsClientLoop"
	case 8260:
		return "ErrorDsReferralLimitExceeded"
	case 8261:
		return "ErrorDsSortControlMissing"
	case 8262:
		return "ErrorDsOffsetRangeError"
	case 8301:
		return "ErrorDsRootMustBeNc"
	case 8302:
		return "ErrorDsAddReplicaInhibited"
	case 8303:
		return "ErrorDsAttNotDefInSchema"
	case 8304:
		return "ErrorDsMaxObjSizeExceeded"
	case 8305:
		return "ErrorDsObjStringNameExists"
	case 8306:
		return "ErrorDsNoRdnDefinedInSchema"
	case 8307:
		return "ErrorDsRdnDoesntMatchSchema"
	case 8308:
		return "ErrorDsNoRequestedAttsFound"
	case 8309:
		return "ErrorDsUserBufferToSmall"
	case 8310:
		return "ErrorDsAttIsNotOnObj"
	case 8311:
		return "ErrorDsIllegalModOperation"
	case 8312:
		return "ErrorDsObjTooLarge"
	case 8313:
		return "ErrorDsBadInstanceType"
	case 8314:
		return "ErrorDsMasterdsaRequired"
	case 8315:
		return "ErrorDsObjectClassRequired"
	case 8316:
		return "ErrorDsMissingRequiredAtt"
	case 8317:
		return "ErrorDsAttNotDefForClass"
	case 8318:
		return "ErrorDsAttAlreadyExists"
	case 8320:
		return "ErrorDsCantAddAttValues"
	case 8321:
		return "ErrorDsSingleValueConstraint"
	case 8322:
		return "ErrorDsRangeConstraint"
	case 8323:
		return "ErrorDsAttValAlreadyExists"
	case 8324:
		return "ErrorDsCantRemMissingAtt"
	case 8325:
		return "ErrorDsCantRemMissingAttVal"
	case 8326:
		return "ErrorDsRootCantBeSubref"
	case 8327:
		return "ErrorDsNoChaining"
	case 8328:
		return "ErrorDsNoChainedEval"
	case 8329:
		return "ErrorDsNoParentObject"
	case 8330:
		return "ErrorDsParentIsAnAlias"
	case 8331:
		return "ErrorDsCantMixMasterAndReps"
	case 8332:
		return "ErrorDsChildrenExist"
	case 8333:
		return "ErrorDsObjNotFound"
	case 8334:
		return "ErrorDsAliasedObjMissing"
	case 8335:
		return "ErrorDsBadNameSyntax"
	case 8336:
		return "ErrorDsAliasPointsToAlias"
	case 8337:
		return "ErrorDsCantDerefAlias"
	case 8338:
		return "ErrorDsOutOfScope"
	case 8339:
		return "ErrorDsObjectBeingRemoved"
	case 8340:
		return "ErrorDsCantDeleteDsaObj"
	case 8341:
		return "ErrorDsGenericError"
	case 8342:
		return "ErrorDsDsaMustBeIntMaster"
	case 8343:
		return "ErrorDsClassNotDsa"
	case 8344:
		return "ErrorDsInsuffAccessRights"
	case 8345:
		return "ErrorDsIllegalSuperior"
	case 8346:
		return "ErrorDsAttributeOwnedBySam"
	case 8347:
		return "ErrorDsNameTooManyParts"
	case 8348:
		return "ErrorDsNameTooLong"
	case 8349:
		return "ErrorDsNameValueTooLong"
	case 8350:
		return "ErrorDsNameUnparseable"
	case 8351:
		return "ErrorDsNameTypeUnknown"
	case 8352:
		return "ErrorDsNotAnObject"
	case 8353:
		return "ErrorDsSecDescTooShort"
	case 8354:
		return "ErrorDsSecDescInvalid"
	case 8355:
		return "ErrorDsNoDeletedName"
	case 8356:
		return "ErrorDsSubrefMustHaveParent"
	case 8357:
		return "ErrorDsNcnameMustBeNc"
	case 8358:
		return "ErrorDsCantAddSystemOnly"
	case 8359:
		return "ErrorDsClassMustBeConcrete"
	case 8360:
		return "ErrorDsInvalidDmd"
	case 8361:
		return "ErrorDsObjGuidExists"
	case 8362:
		return "ErrorDsNotOnBacklink"
	case 8363:
		return "ErrorDsNoCrossrefForNc"
	case 8364:
		return "ErrorDsShuttingDown"
	case 8365:
		return "ErrorDsUnknownOperation"
	case 8366:
		return "ErrorDsInvalidRoleOwner"
	case 8367:
		return "ErrorDsCouldntContactFsmo"
	case 8368:
		return "ErrorDsCrossNcDnRename"
	case 8369:
		return "ErrorDsCantModSystemOnly"
	case 8370:
		return "ErrorDsReplicatorOnly"
	case 8371:
		return "ErrorDsObjClassNotDefined"
	case 8372:
		return "ErrorDsObjClassNotSubclass"
	case 8373:
		return "ErrorDsNameReferenceInvalid"
	case 8374:
		return "ErrorDsCrossRefExists"
	case 8375:
		return "ErrorDsCantDelMasterCrossref"
	case 8376:
		return "ErrorDsSubtreeNotifyNotNcHead"
	case 8377:
		return "ErrorDsNotifyFilterTooComplex"
	case 8378:
		return "ErrorDsDupRdn"
	case 8379:
		return "ErrorDsDupOid"
	case 8380:
		return "ErrorDsDupMapiId"
	case 8381:
		return "ErrorDsDupSchemaIdGuid"
	case 8382:
		return "ErrorDsDupLdapDisplayName"
	case 8383:
		return "ErrorDsSemanticAttTest"
	case 8384:
		return "ErrorDsSyntaxMismatch"
	case 8385:
		return "ErrorDsExistsInMustHave"
	case 8386:
		return "ErrorDsExistsInMayHave"
	case 8387:
		return "ErrorDsNonexistentMayHave"
	case 8388:
		return "ErrorDsNonexistentMustHave"
	case 8389:
		return "ErrorDsAuxClsTestFail"
	case 8390:
		return "ErrorDsNonexistentPossSup"
	case 8391:
		return "ErrorDsSubClsTestFail"
	case 8392:
		return "ErrorDsBadRdnAttIdSyntax"
	case 8393:
		return "ErrorDsExistsInAuxCls"
	case 8394:
		return "ErrorDsExistsInSubCls"
	case 8395:
		return "ErrorDsExistsInPossSup"
	case 8396:
		return "ErrorDsRecalcschemaFailed"
	case 8397:
		return "ErrorDsTreeDeleteNotFinished"
	case 8398:
		return "ErrorDsCantDelete"
	case 8399:
		return "ErrorDsAttSchemaReqId"
	case 8400:
		return "ErrorDsBadAttSchemaSyntax"
	case 8401:
		return "ErrorDsCantCacheAtt"
	case 8402:
		return "ErrorDsCantCacheClass"
	case 8403:
		return "ErrorDsCantRemoveAttCache"
	case 8404:
		return "ErrorDsCantRemoveClassCache"
	case 8405:
		return "ErrorDsCantRetrieveDn"
	case 8406:
		return "ErrorDsMissingSupref"
	case 8407:
		return "ErrorDsCantRetrieveInstance"
	case 8408:
		return "ErrorDsCodeInconsistency"
	case 8409:
		return "ErrorDsDatabaseError"
	case 8410:
		return "ErrorDsGovernsidMissing"
	case 8411:
		return "ErrorDsMissingExpectedAtt"
	case 8412:
		return "ErrorDsNcnameMissingCrRef"
	case 8413:
		return "ErrorDsSecurityCheckingError"
	case 8414:
		return "ErrorDsSchemaNotLoaded"
	case 8415:
		return "ErrorDsSchemaAllocFailed"
	case 8416:
		return "ErrorDsAttSchemaReqSyntax"
	case 8417:
		return "ErrorDsGcverifyError"
	case 8418:
		return "ErrorDsDraSchemaMismatch"
	case 8419:
		return "ErrorDsCantFindDsaObj"
	case 8420:
		return "ErrorDsCantFindExpectedNc"
	case 8421:
		return "ErrorDsCantFindNcInCache"
	case 8422:
		return "ErrorDsCantRetrieveChild"
	case 8423:
		return "ErrorDsSecurityIllegalModify"
	case 8424:
		return "ErrorDsCantReplaceHiddenRec"
	case 8425:
		return "ErrorDsBadHierarchyFile"
	case 8426:
		return "ErrorDsBuildHierarchyTableFailed"
	case 8427:
		return "ErrorDsConfigParamMissing"
	case 8428:
		return "ErrorDsCountingAbIndicesFailed"
	case 8429:
		return "ErrorDsHierarchyTableMallocFailed"
	case 8430:
		return "ErrorDsInternalFailure"
	case 8431:
		return "ErrorDsUnknownError"
	case 8432:
		return "ErrorDsRootRequiresClassTop"
	case 8433:
		return "ErrorDsRefusingFsmoRoles"
	case 8434:
		return "ErrorDsMissingFsmoSettings"
	case 8435:
		return "ErrorDsUnableToSurrenderRoles"
	case 8436:
		return "ErrorDsDraGeneric"
	case 8437:
		return "ErrorDsDraInvalidParameter"
	case 8438:
		return "ErrorDsDraBusy"
	case 8439:
		return "ErrorDsDraBadDn"
	case 8440:
		return "ErrorDsDraBadNc"
	case 8441:
		return "ErrorDsDraDnExists"
	case 8442:
		return "ErrorDsDraInternalError"
	case 8443:
		return "ErrorDsDraInconsistentDit"
	case 8444:
		return "ErrorDsDraConnectionFailed"
	case 8445:
		return "ErrorDsDraBadInstanceType"
	case 8446:
		return "ErrorDsDraOutOfMem"
	case 8447:
		return "ErrorDsDraMailProblem"
	case 8448:
		return "ErrorDsDraRefAlreadyExists"
	case 8449:
		return "ErrorDsDraRefNotFound"
	case 8450:
		return "ErrorDsDraObjIsRepSource"
	case 8451:
		return "ErrorDsDraDbError"
	case 8452:
		return "ErrorDsDraNoReplica"
	case 8453:
		return "ErrorDsDraAccessDenied"
	case 8454:
		return "ErrorDsDraNotSupported"
	case 8455:
		return "ErrorDsDraRpcCancelled"
	case 8456:
		return "ErrorDsDraSourceDisabled"
	case 8457:
		return "ErrorDsDraSinkDisabled"
	case 8458:
		return "ErrorDsDraNameCollision"
	case 8459:
		return "ErrorDsDraSourceReinstalled"
	case 8460:
		return "ErrorDsDraMissingParent"
	case 8461:
		return "ErrorDsDraPreempted"
	case 8462:
		return "ErrorDsDraAbandonSync"
	case 8463:
		return "ErrorDsDraShutdown"
	case 8464:
		return "ErrorDsDraIncompatiblePartialSet"
	case 8465:
		return "ErrorDsDraSourceIsPartialReplica"
	case 8466:
		return "ErrorDsDraExtnConnectionFailed"
	case 8467:
		return "ErrorDsInstallSchemaMismatch"
	case 8468:
		return "ErrorDsDupLinkId"
	case 8469:
		return "ErrorDsNameErrorResolving"
	case 8470:
		return "ErrorDsNameErrorNotFound"
	case 8471:
		return "ErrorDsNameErrorNotUnique"
	case 8472:
		return "ErrorDsNameErrorNoMapping"
	case 8473:
		return "ErrorDsNameErrorDomainOnly"
	case 8474:
		return "ErrorDsNameErrorNoSyntacticalMapping"
	case 8475:
		return "ErrorDsConstructedAttMod"
	case 8476:
		return "ErrorDsWrongOmObjClass"
	case 8477:
		return "ErrorDsDraReplPending"
	case 8478:
		return "ErrorDsDsRequired"
	case 8479:
		return "ErrorDsInvalidLdapDisplayName"
	case 8480:
		return "ErrorDsNonBaseSearch"
	case 8481:
		return "ErrorDsCantRetrieveAtts"
	case 8482:
		return "ErrorDsBacklinkWithoutLink"
	case 8483:
		return "ErrorDsEpochMismatch"
	case 8484:
		return "ErrorDsSrcNameMismatch"
	case 8485:
		return "ErrorDsSrcAndDstNcIdentical"
	case 8486:
		return "ErrorDsDstNcMismatch"
	case 8487:
		return "ErrorDsNotAuthoritiveForDstNc"
	case 8488:
		return "ErrorDsSrcGuidMismatch"
	case 8489:
		return "ErrorDsCantMoveDeletedObject"
	case 8490:
		return "ErrorDsPdcOperationInProgress"
	case 8491:
		return "ErrorDsCrossDomainCleanupReqd"
	case 8492:
		return "ErrorDsIllegalXdomMoveOperation"
	case 8493:
		return "ErrorDsCantWithAcctGroupMembershps"
	case 8494:
		return "ErrorDsNcMustHaveNcParent"
	case 8495:
		return "ErrorDsCrImpossibleToValidate"
	case 8496:
		return "ErrorDsDstDomainNotNative"
	case 8497:
		return "ErrorDsMissingInfrastructureContainer"
	case 8498:
		return "ErrorDsCantMoveAccountGroup"
	case 8499:
		return "ErrorDsCantMoveResourceGroup"
	case 8500:
		return "ErrorDsInvalidSearchFlag"
	case 8501:
		return "ErrorDsNoTreeDeleteAboveNc"
	case 8502:
		return "ErrorDsCouldntLockTreeForDelete"
	case 8503:
		return "ErrorDsCouldntIdentifyObjectsForTreeDelete"
	case 8504:
		return "ErrorDsSamInitFailure"
	case 8505:
		return "ErrorDsSensitiveGroupViolation"
	case 8506:
		return "ErrorDsCantModPrimarygroupid"
	case 8507:
		return "ErrorDsIllegalBaseSchemaMod"
	case 8508:
		return "ErrorDsNonsafeSchemaChange"
	case 8509:
		return "ErrorDsSchemaUpdateDisallowed"
	case 8510:
		return "ErrorDsCantCreateUnderSchema"
	case 8511:
		return "ErrorDsInstallNoSrcSchVersion"
	case 8512:
		return "ErrorDsInstallNoSchVersionInInifile"
	case 8513:
		return "ErrorDsInvalidGroupType"
	case 8514:
		return "ErrorDsNoNestGlobalgroupInMixeddomain"
	case 8515:
		return "ErrorDsNoNestLocalgroupInMixeddomain"
	case 8516:
		return "ErrorDsGlobalCantHaveLocalMember"
	case 8517:
		return "ErrorDsGlobalCantHaveUniversalMember"
	case 8518:
		return "ErrorDsUniversalCantHaveLocalMember"
	case 8519:
		return "ErrorDsGlobalCantHaveCrossdomainMember"
	case 8520:
		return "ErrorDsLocalCantHaveCrossdomainLocalMember"
	case 8521:
		return "ErrorDsHavePrimaryMembers"
	case 8522:
		return "ErrorDsStringSdConversionFailed"
	case 8523:
		return "ErrorDsNamingMasterGc"
	case 8524:
		return "ErrorDsDnsLookupFailure"
	case 8525:
		return "ErrorDsCouldntUpdateSpns"
	case 8526:
		return "ErrorDsCantRetrieveSd"
	case 8527:
		return "ErrorDsKeyNotUnique"
	case 8528:
		return "ErrorDsWrongLinkedAttSyntax"
	case 8529:
		return "ErrorDsSamNeedBootkeyPassword"
	case 8530:
		return "ErrorDsSamNeedBootkeyFloppy"
	case 8531:
		return "ErrorDsCantStart"
	case 8532:
		return "ErrorDsInitFailure"
	case 8533:
		return "ErrorDsNoPktPrivacyOnConnection"
	case 8534:
		return "ErrorDsSourceDomainInForest"
	case 8535:
		return "ErrorDsDestinationDomainNotInForest"
	case 8536:
		return "ErrorDsDestinationAuditingNotEnabled"
	case 8537:
		return "ErrorDsCantFindDcForSrcDomain"
	case 8538:
		return "ErrorDsSrcObjNotGroupOrUser"
	case 8539:
		return "ErrorDsSrcSidExistsInForest"
	case 8540:
		return "ErrorDsSrcAndDstObjectClassMismatch"
	case 8541:
		return "ErrorSamInitFailure"
	case 8542:
		return "ErrorDsDraSchemaInfoShip"
	case 8543:
		return "ErrorDsDraSchemaConflict"
	case 8544:
		return "ErrorDsDraEarlierSchemaConflict"
	case 8545:
		return "ErrorDsDraObjNcMismatch"
	case 8546:
		return "ErrorDsNcStillHasDsas"
	case 8547:
		return "ErrorDsGcRequired"
	case 8548:
		return "ErrorDsLocalMemberOfLocalOnly"
	case 8549:
		return "ErrorDsNoFpoInUniversalGroups"
	case 8550:
		return "ErrorDsCantAddToGc"
	case 8551:
		return "ErrorDsNoCheckpointWithPdc"
	case 8552:
		return "ErrorDsSourceAuditingNotEnabled"
	case 8553:
		return "ErrorDsCantCreateInNondomainNc"
	case 8554:
		return "ErrorDsInvalidNameForSpn"
	case 8555:
		return "ErrorDsFilterUsesContructedAttrs"
	case 8556:
		return "ErrorDsUnicodepwdNotInQuotes"
	case 8557:
		return "ErrorDsMachineAccountQuotaExceeded"
	case 8558:
		return "ErrorDsMustBeRunOnDstDc"
	case 8559:
		return "ErrorDsSrcDcMustBeSp4OrGreater"
	case 8560:
		return "ErrorDsCantTreeDeleteCriticalObj"
	case 8561:
		return "ErrorDsInitFailureConsole"
	case 8562:
		return "ErrorDsSamInitFailureConsole"
	case 8563:
		return "ErrorDsForestVersionTooHigh"
	case 8564:
		return "ErrorDsDomainVersionTooHigh"
	case 8565:
		return "ErrorDsForestVersionTooLow"
	case 8566:
		return "ErrorDsDomainVersionTooLow"
	case 8567:
		return "ErrorDsIncompatibleVersion"
	case 8568:
		return "ErrorDsLowDsaVersion"
	case 8569:
		return "ErrorDsNoBehaviorVersionInMixeddomain"
	case 8570:
		return "ErrorDsNotSupportedSortOrder"
	case 8571:
		return "ErrorDsNameNotUnique"
	case 8572:
		return "ErrorDsMachineAccountCreatedPrent4"
	case 8573:
		return "ErrorDsOutOfVersionStore"
	case 8574:
		return "ErrorDsIncompatibleControlsUsed"
	case 8575:
		return "ErrorDsNoRefDomain"
	case 8576:
		return "ErrorDsReservedLinkId"
	case 8577:
		return "ErrorDsLinkIdNotAvailable"
	case 8578:
		return "ErrorDsAgCantHaveUniversalMember"
	case 8579:
		return "ErrorDsModifydnDisallowedByInstanceType"
	case 8580:
		return "ErrorDsNoObjectMoveInSchemaNc"
	case 8581:
		return "ErrorDsModifydnDisallowedByFlag"
	case 8582:
		return "ErrorDsModifydnWrongGrandparent"
	case 8583:
		return "ErrorDsNameErrorTrustReferral"
	case 8584:
		return "ErrorNotSupportedOnStandardServer"
	case 8585:
		return "ErrorDsCantAccessRemotePartOfAd"
	case 8586:
		return "ErrorDsCrImpossibleToValidateV2"
	case 8587:
		return "ErrorDsThreadLimitExceeded"
	case 8588:
		return "ErrorDsNotClosest"
	case 8589:
		return "ErrorDsCantDeriveSpnWithoutServerRef"
	case 8590:
		return "ErrorDsSingleUserModeFailed"
	case 8591:
		return "ErrorDsNtdscriptSyntaxError"
	case 8592:
		return "ErrorDsNtdscriptProcessError"
	case 8593:
		return "ErrorDsDifferentReplEpochs"
	case 8594:
		return "ErrorDsDrsExtensionsChanged"
	case 8595:
		return "ErrorDsReplicaSetChangeNotAllowedOnDisabledCr"
	case 8596:
		return "ErrorDsNoMsdsIntid"
	case 8597:
		return "ErrorDsDupMsdsIntid"
	case 8598:
		return "ErrorDsExistsInRdnattid"
	case 8599:
		return "ErrorDsAuthorizationFailed"
	case 8600:
		return "ErrorDsInvalidScript"
	case 8601:
		return "ErrorDsRemoteCrossrefOpFailed"
	case 8602:
		return "ErrorDsCrossRefBusy"
	case 8603:
		return "ErrorDsCantDeriveSpnForDeletedDomain"
	case 8604:
		return "ErrorDsCantDemoteWithWriteableNc"
	case 8605:
		return "ErrorDsDuplicateIdFound"
	case 8606:
		return "ErrorDsInsufficientAttrToCreateObject"
	case 8607:
		return "ErrorDsGroupConversionError"
	case 8608:
		return "ErrorDsCantMoveAppBasicGroup"
	case 8609:
		return "ErrorDsCantMoveAppQueryGroup"
	case 8610:
		return "ErrorDsRoleNotVerified"
	case 8611:
		return "ErrorDsWkoContainerCannotBeSpecial"
	case 8612:
		return "ErrorDsDomainRenameInProgress"
	case 8613:
		return "ErrorDsExistingAdChildNc"
	case 8614:
		return "ErrorDsReplLifetimeExceeded"
	case 8615:
		return "ErrorDsDisallowedInSystemContainer"
	case 8616:
		return "ErrorDsLdapSendQueueFull"
	case 8617:
		return "ErrorDsDraOutScheduleWindow"
	case 8618:
		return "ErrorDsPolicyNotKnown"
	case 8619:
		return "ErrorNoSiteSettingsObject"
	case 8620:
		return "ErrorNoSecrets"
	case 8621:
		return "ErrorNoWritableDcFound"
	case 8622:
		return "ErrorDsNoServerObject"
	case 8623:
		return "ErrorDsNoNtdsaObject"
	case 8624:
		return "ErrorDsNonAsqSearch"
	case 8625:
		return "ErrorDsAuditFailure"
	case 8626:
		return "ErrorDsInvalidSearchFlagSubtree"
	case 8627:
		return "ErrorDsInvalidSearchFlagTuple"
	case 8628:
		return "ErrorDsHierarchyTableTooDeep"
	case 8629:
		return "ErrorDsDraCorruptUtdVector"
	case 8630:
		return "ErrorDsDraSecretsDenied"
	case 8631:
		return "ErrorDsReservedMapiId"
	case 8632:
		return "ErrorDsMapiIdNotAvailable"
	case 8633:
		return "ErrorDsDraMissingKrbtgtSecret"
	case 8634:
		return "ErrorDsDomainNameExistsInForest"
	case 8635:
		return "ErrorDsFlatNameExistsInForest"
	case 8636:
		return "ErrorInvalidUserPrincipalName"
	case 8637:
		return "ErrorDsOidMappedGroupCantHaveMembers"
	case 8638:
		return "ErrorDsOidNotFound"
	case 8639:
		return "ErrorDsDraRecycledTarget"
	case 9000:
		return "DnsErrorResponseCodesBase"
	case 0x00002328:
		return "DnsErrorMask"
	case 9001:
		return "DnsErrorRcodeFormatError"
	case 9002:
		return "DnsErrorRcodeServerFailure"
	case 9003:
		return "DnsErrorRcodeNameError"
	case 9004:
		return "DnsErrorRcodeNotImplemented"
	case 9005:
		return "DnsErrorRcodeRefused"
	case 9006:
		return "DnsErrorRcodeYxdomain"
	case 9007:
		return "DnsErrorRcodeYxrrset"
	case 9008:
		return "DnsErrorRcodeNxrrset"
	case 9009:
		return "DnsErrorRcodeNotauth"
	case 9010:
		return "DnsErrorRcodeNotzone"
	case 9016:
		return "DnsErrorRcodeBadsig"
	case 9017:
		return "DnsErrorRcodeBadkey"
	case 9018:
		return "DnsErrorRcodeBadtime"
	case 9500:
		return "DnsErrorPacketFmtBase"
	case 9501:
		return "DnsInfoNoRecords"
	case 9502:
		return "DnsErrorBadPacket"
	case 9503:
		return "DnsErrorNoPacket"
	case 9504:
		return "DnsErrorRcode"
	case 9505:
		return "DnsErrorUnsecurePacket"
	case 9550:
		return "DnsErrorGeneralApiBase"
	case 9551:
		return "DnsErrorInvalidType"
	case 9552:
		return "DnsErrorInvalidIpAddress"
	case 9553:
		return "DnsErrorInvalidProperty"
	case 9554:
		return "DnsErrorTryAgainLater"
	case 9555:
		return "DnsErrorNotUnique"
	case 9556:
		return "DnsErrorNonRfcName"
	case 9557:
		return "DnsStatusFqdn"
	case 9558:
		return "DnsStatusDottedName"
	case 9559:
		return "DnsStatusSinglePartName"
	case 9560:
		return "DnsErrorInvalidNameChar"
	case 9561:
		return "DnsErrorNumericName"
	case 9562:
		return "DnsErrorNotAllowedOnRootServer"
	case 9563:
		return "DnsErrorNotAllowedUnderDelegation"
	case 9564:
		return "DnsErrorCannotFindRootHints"
	case 9565:
		return "DnsErrorInconsistentRootHints"
	case 9566:
		return "DnsErrorDwordValueTooSmall"
	case 9567:
		return "DnsErrorDwordValueTooLarge"
	case 9568:
		return "DnsErrorBackgroundLoading"
	case 9569:
		return "DnsErrorNotAllowedOnRodc"
	case 9570:
		return "DnsErrorNotAllowedUnderDname"
	case 9571:
		return "DnsErrorDelegationRequired"
	case 9572:
		return "DnsErrorInvalidPolicyTable"
	case 9600:
		return "DnsErrorZoneBase"
	case 9601:
		return "DnsErrorZoneDoesNotExist"
	case 9602:
		return "DnsErrorNoZoneInfo"
	case 9603:
		return "DnsErrorInvalidZoneOperation"
	case 9604:
		return "DnsErrorZoneConfigurationError"
	case 9605:
		return "DnsErrorZoneHasNoSoaRecord"
	case 9606:
		return "DnsErrorZoneHasNoNsRecords"
	case 9607:
		return "DnsErrorZoneLocked"
	case 9608:
		return "DnsErrorZoneCreationFailed"
	case 9609:
		return "DnsErrorZoneAlreadyExists"
	case 9610:
		return "DnsErrorAutozoneAlreadyExists"
	case 9611:
		return "DnsErrorInvalidZoneType"
	case 9612:
		return "DnsErrorSecondaryRequiresMasterIp"
	case 9613:
		return "DnsErrorZoneNotSecondary"
	case 9614:
		return "DnsErrorNeedSecondaryAddresses"
	case 9615:
		return "DnsErrorWinsInitFailed"
	case 9616:
		return "DnsErrorNeedWinsServers"
	case 9617:
		return "DnsErrorNbstatInitFailed"
	case 9618:
		return "DnsErrorSoaDeleteInvalid"
	case 9619:
		return "DnsErrorForwarderAlreadyExists"
	case 9620:
		return "DnsErrorZoneRequiresMasterIp"
	case 9621:
		return "DnsErrorZoneIsShutdown"
	case 9650:
		return "DnsErrorDatafileBase"
	case 9651:
		return "DnsErrorPrimaryRequiresDatafile"
	case 9652:
		return "DnsErrorInvalidDatafileName"
	case 9653:
		return "DnsErrorDatafileOpenFailure"
	case 9654:
		return "DnsErrorFileWritebackFailed"
	case 9655:
		return "DnsErrorDatafileParsing"
	case 9700:
		return "DnsErrorDatabaseBase"
	case 9701:
		return "DnsErrorRecordDoesNotExist"
	case 9702:
		return "DnsErrorRecordFormat"
	case 9703:
		return "DnsErrorNodeCreationFailed"
	case 9704:
		return "DnsErrorUnknownRecordType"
	case 9705:
		return "DnsErrorRecordTimedOut"
	case 9706:
		return "DnsErrorNameNotInZone"
	case 9707:
		return "DnsErrorCnameLoop"
	case 9708:
		return "DnsErrorNodeIsCname"
	case 9709:
		return "DnsErrorCnameCollision"
	case 9710:
		return "DnsErrorRecordOnlyAtZoneRoot"
	case 9711:
		return "DnsErrorRecordAlreadyExists"
	case 9712:
		return "DnsErrorSecondaryData"
	case 9713:
		return "DnsErrorNoCreateCacheData"
	case 9714:
		return "DnsErrorNameDoesNotExist"
	case 9715:
		return "DnsWarningPtrCreateFailed"
	case 9716:
		return "DnsWarningDomainUndeleted"
	case 9717:
		return "DnsErrorDsUnavailable"
	case 9718:
		return "DnsErrorDsZoneAlreadyExists"
	case 9719:
		return "DnsErrorNoBootfileIfDsZone"
	case 9720:
		return "DnsErrorNodeIsDname"
	case 9721:
		return "DnsErrorDnameCollision"
	case 9722:
		return "DnsErrorAliasLoop"
	case 9750:
		return "DnsErrorOperationBase"
	case 9751:
		return "DnsInfoAxfrComplete"
	case 9752:
		return "DnsErrorAxfr"
	case 9753:
		return "DnsInfoAddedLocalWins"
	case 9800:
		return "DnsErrorSecureBase"
	case 9801:
		return "DnsStatusContinueNeeded"
	case 9850:
		return "DnsErrorSetupBase"
	case 9851:
		return "DnsErrorNoTcpip"
	case 9852:
		return "DnsErrorNoDnsServers"
	case 9900:
		return "DnsErrorDpBase"
	case 9901:
		return "DnsErrorDpDoesNotExist"
	case 9902:
		return "DnsErrorDpAlreadyExists"
	case 9903:
		return "DnsErrorDpNotEnlisted"
	case 9904:
		return "DnsErrorDpAlreadyEnlisted"
	case 9905:
		return "DnsErrorDpNotAvailable"
	case 9906:
		return "DnsErrorDpFsmoError"
	case 10000:
		return "Wsabaseerr"
	case 10004:
		return "Wsaeintr"
	case 10009:
		return "Wsaebadf"
	case 10013:
		return "Wsaeacces"
	case 10014:
		return "Wsaefault"
	case 10022:
		return "Wsaeinval"
	case 10024:
		return "Wsaemfile"
	case 10035:
		return "Wsaewouldblock"
	case 10036:
		return "Wsaeinprogress"
	case 10037:
		return "Wsaealready"
	case 10038:
		return "Wsaenotsock"
	case 10039:
		return "Wsaedestaddrreq"
	case 10040:
		return "Wsaemsgsize"
	case 10041:
		return "Wsaeprototype"
	case 10042:
		return "Wsaenoprotoopt"
	case 10043:
		return "Wsaeprotonosupport"
	case 10044:
		return "Wsaesocktnosupport"
	case 10045:
		return "Wsaeopnotsupp"
	case 10046:
		return "Wsaepfnosupport"
	case 10047:
		return "Wsaeafnosupport"
	case 10048:
		return "Wsaeaddrinuse"
	case 10049:
		return "Wsaeaddrnotavail"
	case 10050:
		return "Wsaenetdown"
	case 10051:
		return "Wsaenetunreach"
	case 10052:
		return "Wsaenetreset"
	case 10053:
		return "Wsaeconnaborted"
	case 10054:
		return "Wsaeconnreset"
	case 10055:
		return "Wsaenobufs"
	case 10056:
		return "Wsaeisconn"
	case 10057:
		return "Wsaenotconn"
	case 10058:
		return "Wsaeshutdown"
	case 10059:
		return "Wsaetoomanyrefs"
	case 10060:
		return "Wsaetimedout"
	case 10061:
		return "Wsaeconnrefused"
	case 10062:
		return "Wsaeloop"
	case 10063:
		return "Wsaenametoolong"
	case 10064:
		return "Wsaehostdown"
	case 10065:
		return "Wsaehostunreach"
	case 10066:
		return "Wsaenotempty"
	case 10067:
		return "Wsaeproclim"
	case 10068:
		return "Wsaeusers"
	case 10069:
		return "Wsaedquot"
	case 10070:
		return "Wsaestale"
	case 10071:
		return "Wsaeremote"
	case 10091:
		return "Wsasysnotready"
	case 10092:
		return "Wsavernotsupported"
	case 10093:
		return "Wsanotinitialised"
	case 10101:
		return "Wsaediscon"
	case 10102:
		return "Wsaenomore"
	case 10103:
		return "Wsaecancelled"
	case 10104:
		return "Wsaeinvalidproctable"
	case 10105:
		return "Wsaeinvalidprovider"
	case 10106:
		return "Wsaeproviderfailedinit"
	case 10107:
		return "Wsasyscallfailure"
	case 10108:
		return "WsaserviceNotFound"
	case 10109:
		return "WsatypeNotFound"
	case 10110:
		return "WsaENoMore"
	case 10111:
		return "WsaECancelled"
	case 10112:
		return "Wsaerefused"
	case 11001:
		return "WsahostNotFound"
	case 11002:
		return "WsatryAgain"
	case 11003:
		return "WsanoRecovery"
	case 11004:
		return "WsanoData"
	case 11005:
		return "WsaQosReceivers"
	case 11006:
		return "WsaQosSenders"
	case 11007:
		return "WsaQosNoSenders"
	case 11008:
		return "WsaQosNoReceivers"
	case 11009:
		return "WsaQosRequestConfirmed"
	case 11010:
		return "WsaQosAdmissionFailure"
	case 11011:
		return "WsaQosPolicyFailure"
	case 11012:
		return "WsaQosBadStyle"
	case 11013:
		return "WsaQosBadObject"
	case 11014:
		return "WsaQosTrafficCtrlError"
	case 11015:
		return "WsaQosGenericError"
	case 11016:
		return "WsaQosEservicetype"
	case 11017:
		return "WsaQosEflowspec"
	case 11018:
		return "WsaQosEprovspecbuf"
	case 11019:
		return "WsaQosEfilterstyle"
	case 11020:
		return "WsaQosEfiltertype"
	case 11021:
		return "WsaQosEfiltercount"
	case 11022:
		return "WsaQosEobjlength"
	case 11023:
		return "WsaQosEflowcount"
	case 11024:
		return "WsaQosEunkownpsobj"
	case 11025:
		return "WsaQosEpolicyobj"
	case 11026:
		return "WsaQosEflowdesc"
	case 11027:
		return "WsaQosEpsflowspec"
	case 11028:
		return "WsaQosEpsfilterspec"
	case 11029:
		return "WsaQosEsdmodeobj"
	case 11030:
		return "WsaQosEshaperateobj"
	case 11031:
		return "WsaQosReservedPetype"
	case 11032:
		return "WsaSecureHostNotFound"
	case 11033:
		return "WsaIpsecNamePolicyError"
	case 13000:
		return "ErrorIpsecQmPolicyExists"
	case 13001:
		return "ErrorIpsecQmPolicyNotFound"
	case 13002:
		return "ErrorIpsecQmPolicyInUse"
	case 13003:
		return "ErrorIpsecMmPolicyExists"
	case 13004:
		return "ErrorIpsecMmPolicyNotFound"
	case 13005:
		return "ErrorIpsecMmPolicyInUse"
	case 13006:
		return "ErrorIpsecMmFilterExists"
	case 13007:
		return "ErrorIpsecMmFilterNotFound"
	case 13008:
		return "ErrorIpsecTransportFilterExists"
	case 13009:
		return "ErrorIpsecTransportFilterNotFound"
	case 13010:
		return "ErrorIpsecMmAuthExists"
	case 13011:
		return "ErrorIpsecMmAuthNotFound"
	case 13012:
		return "ErrorIpsecMmAuthInUse"
	case 13013:
		return "ErrorIpsecDefaultMmPolicyNotFound"
	case 13014:
		return "ErrorIpsecDefaultMmAuthNotFound"
	case 13015:
		return "ErrorIpsecDefaultQmPolicyNotFound"
	case 13016:
		return "ErrorIpsecTunnelFilterExists"
	case 13017:
		return "ErrorIpsecTunnelFilterNotFound"
	case 13018:
		return "ErrorIpsecMmFilterPendingDeletion"
	case 13019:
		return "ErrorIpsecTransportFilterPendingDeletion"
	case 13020:
		return "ErrorIpsecTunnelFilterPendingDeletion"
	case 13021:
		return "ErrorIpsecMmPolicyPendingDeletion"
	case 13022:
		return "ErrorIpsecMmAuthPendingDeletion"
	case 13023:
		return "ErrorIpsecQmPolicyPendingDeletion"
	case 13024:
		return "WarningIpsecMmPolicyPruned"
	case 13025:
		return "WarningIpsecQmPolicyPruned"
	case 13800:
		return "ErrorIpsecIkeNegStatusBegin"
	case 13801:
		return "ErrorIpsecIkeAuthFail"
	case 13802:
		return "ErrorIpsecIkeAttribFail"
	case 13803:
		return "ErrorIpsecIkeNegotiationPending"
	case 13804:
		return "ErrorIpsecIkeGeneralProcessingError"
	case 13805:
		return "ErrorIpsecIkeTimedOut"
	case 13806:
		return "ErrorIpsecIkeNoCert"
	case 13807:
		return "ErrorIpsecIkeSaDeleted"
	case 13808:
		return "ErrorIpsecIkeSaReaped"
	case 13809:
		return "ErrorIpsecIkeMmAcquireDrop"
	case 13810:
		return "ErrorIpsecIkeQmAcquireDrop"
	case 13811:
		return "ErrorIpsecIkeQueueDropMm"
	case 13812:
		return "ErrorIpsecIkeQueueDropNoMm"
	case 13813:
		return "ErrorIpsecIkeDropNoResponse"
	case 13814:
		return "ErrorIpsecIkeMmDelayDrop"
	case 13815:
		return "ErrorIpsecIkeQmDelayDrop"
	case 13816:
		return "ErrorIpsecIkeError"
	case 13817:
		return "ErrorIpsecIkeCrlFailed"
	case 13818:
		return "ErrorIpsecIkeInvalidKeyUsage"
	case 13819:
		return "ErrorIpsecIkeInvalidCertType"
	case 13820:
		return "ErrorIpsecIkeNoPrivateKey"
	case 13821:
		return "ErrorIpsecIkeSimultaneousRekey"
	case 13822:
		return "ErrorIpsecIkeDhFail"
	case 13823:
		return "ErrorIpsecIkeCriticalPayloadNotRecognized"
	case 13824:
		return "ErrorIpsecIkeInvalidHeader"
	case 13825:
		return "ErrorIpsecIkeNoPolicy"
	case 13826:
		return "ErrorIpsecIkeInvalidSignature"
	case 13827:
		return "ErrorIpsecIkeKerberosError"
	case 13828:
		return "ErrorIpsecIkeNoPublicKey"
	case 13829:
		return "ErrorIpsecIkeProcessErr"
	case 13830:
		return "ErrorIpsecIkeProcessErrSa"
	case 13831:
		return "ErrorIpsecIkeProcessErrProp"
	case 13832:
		return "ErrorIpsecIkeProcessErrTrans"
	case 13833:
		return "ErrorIpsecIkeProcessErrKe"
	case 13834:
		return "ErrorIpsecIkeProcessErrId"
	case 13835:
		return "ErrorIpsecIkeProcessErrCert"
	case 13836:
		return "ErrorIpsecIkeProcessErrCertReq"
	case 13837:
		return "ErrorIpsecIkeProcessErrHash"
	case 13838:
		return "ErrorIpsecIkeProcessErrSig"
	case 13839:
		return "ErrorIpsecIkeProcessErrNonce"
	case 13840:
		return "ErrorIpsecIkeProcessErrNotify"
	case 13841:
		return "ErrorIpsecIkeProcessErrDelete"
	case 13842:
		return "ErrorIpsecIkeProcessErrVendor"
	case 13843:
		return "ErrorIpsecIkeInvalidPayload"
	case 13844:
		return "ErrorIpsecIkeLoadSoftSa"
	case 13845:
		return "ErrorIpsecIkeSoftSaTornDown"
	case 13846:
		return "ErrorIpsecIkeInvalidCookie"
	case 13847:
		return "ErrorIpsecIkeNoPeerCert"
	case 13848:
		return "ErrorIpsecIkePeerCrlFailed"
	case 13849:
		return "ErrorIpsecIkePolicyChange"
	case 13850:
		return "ErrorIpsecIkeNoMmPolicy"
	case 13851:
		return "ErrorIpsecIkeNotcbpriv"
	case 13852:
		return "ErrorIpsecIkeSecloadfail"
	case 13853:
		return "ErrorIpsecIkeFailsspinit"
	case 13854:
		return "ErrorIpsecIkeFailqueryssp"
	case 13855:
		return "ErrorIpsecIkeSrvacqfail"
	case 13856:
		return "ErrorIpsecIkeSrvquerycred"
	case 13857:
		return "ErrorIpsecIkeGetspifail"
	case 13858:
		return "ErrorIpsecIkeInvalidFilter"
	case 13859:
		return "ErrorIpsecIkeOutOfMemory"
	case 13860:
		return "ErrorIpsecIkeAddUpdateKeyFailed"
	case 13861:
		return "ErrorIpsecIkeInvalidPolicy"
	case 13862:
		return "ErrorIpsecIkeUnknownDoi"
	case 13863:
		return "ErrorIpsecIkeInvalidSituation"
	case 13864:
		return "ErrorIpsecIkeDhFailure"
	case 13865:
		return "ErrorIpsecIkeInvalidGroup"
	case 13866:
		return "ErrorIpsecIkeEncrypt"
	case 13867:
		return "ErrorIpsecIkeDecrypt"
	case 13868:
		return "ErrorIpsecIkePolicyMatch"
	case 13869:
		return "ErrorIpsecIkeUnsupportedId"
	case 13870:
		return "ErrorIpsecIkeInvalidHash"
	case 13871:
		return "ErrorIpsecIkeInvalidHashAlg"
	case 13872:
		return "ErrorIpsecIkeInvalidHashSize"
	case 13873:
		return "ErrorIpsecIkeInvalidEncryptAlg"
	case 13874:
		return "ErrorIpsecIkeInvalidAuthAlg"
	case 13875:
		return "ErrorIpsecIkeInvalidSig"
	case 13876:
		return "ErrorIpsecIkeLoadFailed"
	case 13877:
		return "ErrorIpsecIkeRpcDelete"
	case 13878:
		return "ErrorIpsecIkeBenignReinit"
	case 13879:
		return "ErrorIpsecIkeInvalidResponderLifetimeNotify"
	case 13880:
		return "ErrorIpsecIkeInvalidMajorVersion"
	case 13881:
		return "ErrorIpsecIkeInvalidCertKeylen"
	case 13882:
		return "ErrorIpsecIkeMmLimit"
	case 13883:
		return "ErrorIpsecIkeNegotiationDisabled"
	case 13884:
		return "ErrorIpsecIkeQmLimit"
	case 13885:
		return "ErrorIpsecIkeMmExpired"
	case 13886:
		return "ErrorIpsecIkePeerMmAssumedInvalid"
	case 13887:
		return "ErrorIpsecIkeCertChainPolicyMismatch"
	case 13888:
		return "ErrorIpsecIkeUnexpectedMessageId"
	case 13889:
		return "ErrorIpsecIkeInvalidAuthPayload"
	case 13890:
		return "ErrorIpsecIkeDosCookieSent"
	case 13891:
		return "ErrorIpsecIkeShuttingDown"
	case 13892:
		return "ErrorIpsecIkeCgaAuthFailed"
	case 13893:
		return "ErrorIpsecIkeProcessErrNatoa"
	case 13894:
		return "ErrorIpsecIkeInvalidMmForQm"
	case 13895:
		return "ErrorIpsecIkeQmExpired"
	case 13896:
		return "ErrorIpsecIkeTooManyFilters"
	case 13897:
		return "ErrorIpsecIkeNegStatusEnd"
	case 13898:
		return "ErrorIpsecIkeKillDummyNapTunnel"
	case 13899:
		return "ErrorIpsecIkeInnerIpAssignmentFailure"
	case 13900:
		return "ErrorIpsecIkeRequireCpPayloadMissing"
	case 13901:
		return "ErrorIpsecKeyModuleImpersonationNegotiationPending"
	case 13902:
		return "ErrorIpsecIkeCoexistenceSuppress"
	case 13903:
		return "ErrorIpsecIkeRatelimitDrop"
	case 13904:
		return "ErrorIpsecIkePeerDoesntSupportMobike"
	case 13905:
		return "ErrorIpsecIkeAuthorizationFailure"
	case 13906:
		return "ErrorIpsecIkeStrongCredAuthorizationFailure"
	case 13907:
		return "ErrorIpsecIkeAuthorizationFailureWithOptionalRetry"
	case 13908:
		return "ErrorIpsecIkeStrongCredAuthorizationAndCertmapFailure"
	case 13909:
		return "ErrorIpsecIkeNegStatusExtendedEnd"
	case 13910:
		return "ErrorIpsecBadSpi"
	case 13911:
		return "ErrorIpsecSaLifetimeExpired"
	case 13912:
		return "ErrorIpsecWrongSa"
	case 13913:
		return "ErrorIpsecReplayCheckFailed"
	case 13914:
		return "ErrorIpsecInvalidPacket"
	case 13915:
		return "ErrorIpsecIntegrityCheckFailed"
	case 13916:
		return "ErrorIpsecClearTextDrop"
	case 13917:
		return "ErrorIpsecAuthFirewallDrop"
	case 13918:
		return "ErrorIpsecThrottleDrop"
	case 13925:
		return "ErrorIpsecDospBlock"
	case 13926:
		return "ErrorIpsecDospReceivedMulticast"
	case 13927:
		return "ErrorIpsecDospInvalidPacket"
	case 13928:
		return "ErrorIpsecDospStateLookupFailed"
	case 13929:
		return "ErrorIpsecDospMaxEntries"
	case 13930:
		return "ErrorIpsecDospKeymodNotAllowed"
	case 13931:
		return "ErrorIpsecDospNotInstalled"
	case 13932:
		return "ErrorIpsecDospMaxPerIpRatelimitQueues"
	case 14000:
		return "ErrorSxsSectionNotFound"
	case 14001:
		return "ErrorSxsCantGenActctx"
	case 14002:
		return "ErrorSxsInvalidActctxdataFormat"
	case 14003:
		return "ErrorSxsAssemblyNotFound"
	case 14004:
		return "ErrorSxsManifestFormatError"
	case 14005:
		return "ErrorSxsManifestParseError"
	case 14006:
		return "ErrorSxsActivationContextDisabled"
	case 14007:
		return "ErrorSxsKeyNotFound"
	case 14008:
		return "ErrorSxsVersionConflict"
	case 14009:
		return "ErrorSxsWrongSectionType"
	case 14010:
		return "ErrorSxsThreadQueriesDisabled"
	case 14011:
		return "ErrorSxsProcessDefaultAlreadySet"
	case 14012:
		return "ErrorSxsUnknownEncodingGroup"
	case 14013:
		return "ErrorSxsUnknownEncoding"
	case 14014:
		return "ErrorSxsInvalidXmlNamespaceUri"
	case 14015:
		return "ErrorSxsRootManifestDependencyNotInstalled"
	case 14016:
		return "ErrorSxsLeafManifestDependencyNotInstalled"
	case 14017:
		return "ErrorSxsInvalidAssemblyIdentityAttribute"
	case 14018:
		return "ErrorSxsManifestMissingRequiredDefaultNamespace"
	case 14019:
		return "ErrorSxsManifestInvalidRequiredDefaultNamespace"
	case 14020:
		return "ErrorSxsPrivateManifestCrossPathWithReparsePoint"
	case 14021:
		return "ErrorSxsDuplicateDllName"
	case 14022:
		return "ErrorSxsDuplicateWindowclassName"
	case 14023:
		return "ErrorSxsDuplicateClsid"
	case 14024:
		return "ErrorSxsDuplicateIid"
	case 14025:
		return "ErrorSxsDuplicateTlbid"
	case 14026:
		return "ErrorSxsDuplicateProgid"
	case 14027:
		return "ErrorSxsDuplicateAssemblyName"
	case 14028:
		return "ErrorSxsFileHashMismatch"
	case 14029:
		return "ErrorSxsPolicyParseError"
	case 14030:
		return "ErrorSxsXmlEMissingquote"
	case 14031:
		return "ErrorSxsXmlECommentsyntax"
	case 14032:
		return "ErrorSxsXmlEBadstartnamechar"
	case 14033:
		return "ErrorSxsXmlEBadnamechar"
	case 14034:
		return "ErrorSxsXmlEBadcharinstring"
	case 14035:
		return "ErrorSxsXmlEXmldeclsyntax"
	case 14036:
		return "ErrorSxsXmlEBadchardata"
	case 14037:
		return "ErrorSxsXmlEMissingwhitespace"
	case 14038:
		return "ErrorSxsXmlEExpectingtagend"
	case 14039:
		return "ErrorSxsXmlEMissingsemicolon"
	case 14040:
		return "ErrorSxsXmlEUnbalancedparen"
	case 14041:
		return "ErrorSxsXmlEInternalerror"
	case 14042:
		return "ErrorSxsXmlEUnexpectedWhitespace"
	case 14043:
		return "ErrorSxsXmlEIncompleteEncoding"
	case 14044:
		return "ErrorSxsXmlEMissingParen"
	case 14045:
		return "ErrorSxsXmlEExpectingclosequote"
	case 14046:
		return "ErrorSxsXmlEMultipleColons"
	case 14047:
		return "ErrorSxsXmlEInvalidDecimal"
	case 14048:
		return "ErrorSxsXmlEInvalidHexidecimal"
	case 14049:
		return "ErrorSxsXmlEInvalidUnicode"
	case 14050:
		return "ErrorSxsXmlEWhitespaceorquestionmark"
	case 14051:
		return "ErrorSxsXmlEUnexpectedendtag"
	case 14052:
		return "ErrorSxsXmlEUnclosedtag"
	case 14053:
		return "ErrorSxsXmlEDuplicateattribute"
	case 14054:
		return "ErrorSxsXmlEMultipleroots"
	case 14055:
		return "ErrorSxsXmlEInvalidatrootlevel"
	case 14056:
		return "ErrorSxsXmlEBadxmldecl"
	case 14057:
		return "ErrorSxsXmlEMissingroot"
	case 14058:
		return "ErrorSxsXmlEUnexpectedeof"
	case 14059:
		return "ErrorSxsXmlEBadperefinsubset"
	case 14060:
		return "ErrorSxsXmlEUnclosedstarttag"
	case 14061:
		return "ErrorSxsXmlEUnclosedendtag"
	case 14062:
		return "ErrorSxsXmlEUnclosedstring"
	case 14063:
		return "ErrorSxsXmlEUnclosedcomment"
	case 14064:
		return "ErrorSxsXmlEUncloseddecl"
	case 14065:
		return "ErrorSxsXmlEUnclosedcdata"
	case 14066:
		return "ErrorSxsXmlEReservednamespace"
	case 14067:
		return "ErrorSxsXmlEInvalidencoding"
	case 14068:
		return "ErrorSxsXmlEInvalidswitch"
	case 14069:
		return "ErrorSxsXmlEBadxmlcase"
	case 14070:
		return "ErrorSxsXmlEInvalidStandalone"
	case 14071:
		return "ErrorSxsXmlEUnexpectedStandalone"
	case 14072:
		return "ErrorSxsXmlEInvalidVersion"
	case 14073:
		return "ErrorSxsXmlEMissingequals"
	case 14074:
		return "ErrorSxsProtectionRecoveryFailed"
	case 14075:
		return "ErrorSxsProtectionPublicKeyTooShort"
	case 14076:
		return "ErrorSxsProtectionCatalogNotValid"
	case 14077:
		return "ErrorSxsUntranslatableHresult"
	case 14078:
		return "ErrorSxsProtectionCatalogFileMissing"
	case 14079:
		return "ErrorSxsMissingAssemblyIdentityAttribute"
	case 14080:
		return "ErrorSxsInvalidAssemblyIdentityAttributeName"
	case 14081:
		return "ErrorSxsAssemblyMissing"
	case 14082:
		return "ErrorSxsCorruptActivationStack"
	case 14083:
		return "ErrorSxsCorruption"
	case 14084:
		return "ErrorSxsEarlyDeactivation"
	case 14085:
		return "ErrorSxsInvalidDeactivation"
	case 14086:
		return "ErrorSxsMultipleDeactivation"
	case 14087:
		return "ErrorSxsProcessTerminationRequested"
	case 14088:
		return "ErrorSxsReleaseActivationContext"
	case 14089:
		return "ErrorSxsSystemDefaultActivationContextEmpty"
	case 14090:
		return "ErrorSxsInvalidIdentityAttributeValue"
	case 14091:
		return "ErrorSxsInvalidIdentityAttributeName"
	case 14092:
		return "ErrorSxsIdentityDuplicateAttribute"
	case 14093:
		return "ErrorSxsIdentityParseError"
	case 14094:
		return "ErrorMalformedSubstitutionString"
	case 14095:
		return "ErrorSxsIncorrectPublicKeyToken"
	case 14096:
		return "ErrorUnmappedSubstitutionString"
	case 14097:
		return "ErrorSxsAssemblyNotLocked"
	case 14098:
		return "ErrorSxsComponentStoreCorrupt"
	case 14099:
		return "ErrorAdvancedInstallerFailed"
	case 14100:
		return "ErrorXmlEncodingMismatch"
	case 14101:
		return "ErrorSxsManifestIdentitySameButContentsDifferent"
	case 14102:
		return "ErrorSxsIdentitiesDifferent"
	case 14103:
		return "ErrorSxsAssemblyIsNotADeployment"
	case 14104:
		return "ErrorSxsFileNotPartOfAssembly"
	case 14105:
		return "ErrorSxsManifestTooBig"
	case 14106:
		return "ErrorSxsSettingNotRegistered"
	case 14107:
		return "ErrorSxsTransactionClosureIncomplete"
	case 14108:
		return "ErrorSmiPrimitiveInstallerFailed"
	case 14109:
		return "ErrorGenericCommandFailed"
	case 14110:
		return "ErrorSxsFileHashMissing"
	case 15000:
		return "ErrorEvtInvalidChannelPath"
	case 15001:
		return "ErrorEvtInvalidQuery"
	case 15002:
		return "ErrorEvtPublisherMetadataNotFound"
	case 15003:
		return "ErrorEvtEventTemplateNotFound"
	case 15004:
		return "ErrorEvtInvalidPublisherName"
	case 15005:
		return "ErrorEvtInvalidEventData"
	case 15007:
		return "ErrorEvtChannelNotFound"
	case 15008:
		return "ErrorEvtMalformedXmlText"
	case 15009:
		return "ErrorEvtSubscriptionToDirectChannel"
	case 15010:
		return "ErrorEvtConfigurationError"
	case 15011:
		return "ErrorEvtQueryResultStale"
	case 15012:
		return "ErrorEvtQueryResultInvalidPosition"
	case 15013:
		return "ErrorEvtNonValidatingMsxml"
	case 15014:
		return "ErrorEvtFilterAlreadyscoped"
	case 15015:
		return "ErrorEvtFilterNoteltset"
	case 15016:
		return "ErrorEvtFilterInvarg"
	case 15017:
		return "ErrorEvtFilterInvtest"
	case 15018:
		return "ErrorEvtFilterInvtype"
	case 15019:
		return "ErrorEvtFilterParseerr"
	case 15020:
		return "ErrorEvtFilterUnsupportedop"
	case 15021:
		return "ErrorEvtFilterUnexpectedtoken"
	case 15022:
		return "ErrorEvtInvalidOperationOverEnabledDirectChannel"
	case 15023:
		return "ErrorEvtInvalidChannelPropertyValue"
	case 15024:
		return "ErrorEvtInvalidPublisherPropertyValue"
	case 15025:
		return "ErrorEvtChannelCannotActivate"
	case 15026:
		return "ErrorEvtFilterTooComplex"
	case 15027:
		return "ErrorEvtMessageNotFound"
	case 15028:
		return "ErrorEvtMessageIdNotFound"
	case 15029:
		return "ErrorEvtUnresolvedValueInsert"
	case 15030:
		return "ErrorEvtUnresolvedParameterInsert"
	case 15031:
		return "ErrorEvtMaxInsertsReached"
	case 15032:
		return "ErrorEvtEventDefinitionNotFound"
	case 15033:
		return "ErrorEvtMessageLocaleNotFound"
	case 15034:
		return "ErrorEvtVersionTooOld"
	case 15035:
		return "ErrorEvtVersionTooNew"
	case 15036:
		return "ErrorEvtCannotOpenChannelOfQuery"
	case 15037:
		return "ErrorEvtPublisherDisabled"
	case 15038:
		return "ErrorEvtFilterOutOfRange"
	case 15080:
		return "ErrorEcSubscriptionCannotActivate"
	case 15081:
		return "ErrorEcLogDisabled"
	case 15082:
		return "ErrorEcCircularForwarding"
	case 15083:
		return "ErrorEcCredstoreFull"
	case 15084:
		return "ErrorEcCredNotFound"
	case 15085:
		return "ErrorEcNoActiveChannel"
	case 15100:
		return "ErrorMuiFileNotFound"
	case 15101:
		return "ErrorMuiInvalidFile"
	case 15102:
		return "ErrorMuiInvalidRcConfig"
	case 15103:
		return "ErrorMuiInvalidLocaleName"
	case 15104:
		return "ErrorMuiInvalidUltimatefallbackName"
	case 15105:
		return "ErrorMuiFileNotLoaded"
	case 15106:
		return "ErrorResourceEnumUserStop"
	case 15107:
		return "ErrorMuiIntlsettingsUilangNotInstalled"
	case 15108:
		return "ErrorMuiIntlsettingsInvalidLocaleName"
	case 15200:
		return "ErrorMcaInvalidCapabilitiesString"
	case 15201:
		return "ErrorMcaInvalidVcpVersion"
	case 15202:
		return "ErrorMcaMonitorViolatesMccsSpecification"
	case 15203:
		return "ErrorMcaMccsVersionMismatch"
	case 15204:
		return "ErrorMcaUnsupportedMccsVersion"
	case 15205:
		return "ErrorMcaInternalError"
	case 15206:
		return "ErrorMcaInvalidTechnologyTypeReturned"
	case 15207:
		return "ErrorMcaUnsupportedColorTemperature"
	case 15250:
		return "ErrorAmbiguousSystemDevice"
	case 15299:
		return "ErrorSystemDeviceNotFound"
	case 15300:
		return "ErrorHashNotSupported"
	case 15301:
		return "ErrorHashNotPresent"
	case 0:
		return "SeveritySuccess"
	case 1:
		return "SeverityError"
	case 0x10000000:
		return "FacilityNtBit"
	case 0x8000FFFF:
		return "EUnexpected"
	case 0x80004001:
		return "ENotimpl"
	case 0x8007000E:
		return "EOutofmemory"
	case 0x80070057:
		return "EInvalidarg"
	case 0x80004002:
		return "ENointerface"
	case 0x80004003:
		return "EPointer"
	case 0x80070006:
		return "EHandle"
	case 0x80004004:
		return "EAbort"
	case 0x80004005:
		return "EFail"
	case 0x80070005:
		return "EAccessdenied"
	case 0x8000000A:
		return "EPending"
	case 0x80004006:
		return "CoEInitTls"
	case 0x80004007:
		return "CoEInitSharedAllocator"
	case 0x80004008:
		return "CoEInitMemoryAllocator"
	case 0x80004009:
		return "CoEInitClassCache"
	case 0x8000400A:
		return "CoEInitRpcChannel"
	case 0x8000400B:
		return "CoEInitTlsSetChannelControl"
	case 0x8000400C:
		return "CoEInitTlsChannelControl"
	case 0x8000400D:
		return "CoEInitUnacceptedUserAllocator"
	case 0x8000400E:
		return "CoEInitScmMutexExists"
	case 0x8000400F:
		return "CoEInitScmFileMappingExists"
	case 0x80004010:
		return "CoEInitScmMapViewOfFile"
	case 0x80004011:
		return "CoEInitScmExecFailure"
	case 0x80004012:
		return "CoEInitOnlySingleThreaded"
	case 0x80004013:
		return "CoECantRemote"
	case 0x80004014:
		return "CoEBadServerName"
	case 0x80004015:
		return "CoEWrongServerIdentity"
	case 0x80004016:
		return "CoEOle1DdeDisabled"
	case 0x80004017:
		return "CoERunasSyntax"
	case 0x80004018:
		return "CoECreateprocessFailure"
	case 0x80004019:
		return "CoERunasCreateprocessFailure"
	case 0x8000401A:
		return "CoERunasLogonFailure"
	case 0x8000401B:
		return "CoELaunchPermssionDenied"
	case 0x8000401C:
		return "CoEStartServiceFailure"
	case 0x8000401D:
		return "CoERemoteCommunicationFailure"
	case 0x8000401E:
		return "CoEServerStartTimeout"
	case 0x8000401F:
		return "CoEClsregInconsistent"
	case 0x80004020:
		return "CoEIidregInconsistent"
	case 0x80004021:
		return "CoENotSupported"
	case 0x80004022:
		return "CoEReloadDll"
	case 0x80004023:
		return "CoEMsiError"
	case 0x80004024:
		return "CoEAttemptToCreateOutsideClientContext"
	case 0x80004025:
		return "CoEServerPaused"
	case 0x80004026:
		return "CoEServerNotPaused"
	case 0x80004027:
		return "CoEClassDisabled"
	case 0x80004028:
		return "CoEClrnotavailable"
	case 0x80004029:
		return "CoEAsyncWorkRejected"
	case 0x8000402A:
		return "CoEServerInitTimeout"
	case 0x8000402B:
		return "CoENoSecctxInActivate"
	case 0x80004030:
		return "CoETrackerConfig"
	case 0x80004031:
		return "CoEThreadpoolConfig"
	case 0x80004032:
		return "CoESxsConfig"
	case 0x80004033:
		return "CoEMalformedSpn"
	case 0:
		return "SOk"
	case 1:
		return "SFalse"
	case 0x80040000:
		return "OleEFirst"
	case 0x800400FF:
		return "OleELast"
	case 0x00040000:
		return "OleSFirst"
	case 0x000400FF:
		return "OleSLast"
	case 0x80040000:
		return "OleEOleverb"
	case 0x80040001:
		return "OleEAdvf"
	case 0x80040002:
		return "OleEEnumNomore"
	case 0x80040003:
		return "OleEAdvisenotsupported"
	case 0x80040004:
		return "OleENoconnection"
	case 0x80040005:
		return "OleENotrunning"
	case 0x80040006:
		return "OleENocache"
	case 0x80040007:
		return "OleEBlank"
	case 0x80040008:
		return "OleEClassdiff"
	case 0x80040009:
		return "OleECantGetmoniker"
	case 0x8004000A:
		return "OleECantBindtosource"
	case 0x8004000B:
		return "OleEStatic"
	case 0x8004000C:
		return "OleEPromptsavecancelled"
	case 0x8004000D:
		return "OleEInvalidrect"
	case 0x8004000E:
		return "OleEWrongcompobj"
	case 0x8004000F:
		return "OleEInvalidhwnd"
	case 0x80040010:
		return "OleENotInplaceactive"
	case 0x80040011:
		return "OleECantconvert"
	case 0x80040012:
		return "OleENostorage"
	case 0x80040064:
		return "DvEFormatetc"
	case 0x80040065:
		return "DvEDvtargetdevice"
	case 0x80040066:
		return "DvEStgmedium"
	case 0x80040067:
		return "DvEStatdata"
	case 0x80040068:
		return "DvELindex"
	case 0x80040069:
		return "DvETymed"
	case 0x8004006A:
		return "DvEClipformat"
	case 0x8004006B:
		return "DvEDvaspect"
	case 0x8004006C:
		return "DvEDvtargetdeviceSize"
	case 0x8004006D:
		return "DvENoiviewobject"
	case 0x80040100:
		return "DragdropEFirst"
	case 0x8004010F:
		return "DragdropELast"
	case 0x00040100:
		return "DragdropSFirst"
	case 0x0004010F:
		return "DragdropSLast"
	case 0x80040100:
		return "DragdropENotregistered"
	case 0x80040101:
		return "DragdropEAlreadyregistered"
	case 0x80040102:
		return "DragdropEInvalidhwnd"
	case 0x80040110:
		return "ClassfactoryEFirst"
	case 0x8004011F:
		return "ClassfactoryELast"
	case 0x00040110:
		return "ClassfactorySFirst"
	case 0x0004011F:
		return "ClassfactorySLast"
	case 0x80040110:
		return "ClassENoaggregation"
	case 0x80040111:
		return "ClassEClassnotavailable"
	case 0x80040112:
		return "ClassENotlicensed"
	case 0x80040120:
		return "MarshalEFirst"
	case 0x8004012F:
		return "MarshalELast"
	case 0x00040120:
		return "MarshalSFirst"
	case 0x0004012F:
		return "MarshalSLast"
	case 0x80040130:
		return "DataEFirst"
	case 0x8004013F:
		return "DataELast"
	case 0x00040130:
		return "DataSFirst"
	case 0x0004013F:
		return "DataSLast"
	case 0x80040140:
		return "ViewEFirst"
	case 0x8004014F:
		return "ViewELast"
	case 0x00040140:
		return "ViewSFirst"
	case 0x0004014F:
		return "ViewSLast"
	case 0x80040140:
		return "ViewEDraw"
	case 0x80040150:
		return "RegdbEFirst"
	case 0x8004015F:
		return "RegdbELast"
	case 0x00040150:
		return "RegdbSFirst"
	case 0x0004015F:
		return "RegdbSLast"
	case 0x80040150:
		return "RegdbEReadregdb"
	case 0x80040151:
		return "RegdbEWriteregdb"
	case 0x80040152:
		return "RegdbEKeymissing"
	case 0x80040153:
		return "RegdbEInvalidvalue"
	case 0x80040154:
		return "RegdbEClassnotreg"
	case 0x80040155:
		return "RegdbEIidnotreg"
	case 0x80040156:
		return "RegdbEBadthreadingmodel"
	case 0x80040160:
		return "CatEFirst"
	case 0x80040161:
		return "CatELast"
	case 0x80040160:
		return "CatECatidnoexist"
	case 0x80040161:
		return "CatENodescription"
	case 0x80040164:
		return "CsEFirst"
	case 0x8004016F:
		return "CsELast"
	case 0x80040164:
		return "CsEPackageNotfound"
	case 0x80040165:
		return "CsENotDeletable"
	case 0x80040166:
		return "CsEClassNotfound"
	case 0x80040167:
		return "CsEInvalidVersion"
	case 0x80040168:
		return "CsENoClassstore"
	case 0x80040169:
		return "CsEObjectNotfound"
	case 0x8004016A:
		return "CsEObjectAlreadyExists"
	case 0x8004016B:
		return "CsEInvalidPath"
	case 0x8004016C:
		return "CsENetworkError"
	case 0x8004016D:
		return "CsEAdminLimitExceeded"
	case 0x8004016E:
		return "CsESchemaMismatch"
	case 0x8004016F:
		return "CsEInternalError"
	case 0x80040170:
		return "CacheEFirst"
	case 0x8004017F:
		return "CacheELast"
	case 0x00040170:
		return "CacheSFirst"
	case 0x0004017F:
		return "CacheSLast"
	case 0x80040170:
		return "CacheENocacheUpdated"
	case 0x80040180:
		return "OleobjEFirst"
	case 0x8004018F:
		return "OleobjELast"
	case 0x00040180:
		return "OleobjSFirst"
	case 0x0004018F:
		return "OleobjSLast"
	case 0x80040180:
		return "OleobjENoverbs"
	case 0x80040181:
		return "OleobjEInvalidverb"
	case 0x80040190:
		return "ClientsiteEFirst"
	case 0x8004019F:
		return "ClientsiteELast"
	case 0x00040190:
		return "ClientsiteSFirst"
	case 0x0004019F:
		return "ClientsiteSLast"
	case 0x800401A0:
		return "InplaceENotundoable"
	case 0x800401A1:
		return "InplaceENotoolspace"
	case 0x800401A0:
		return "InplaceEFirst"
	case 0x800401AF:
		return "InplaceELast"
	case 0x000401A0:
		return "InplaceSFirst"
	case 0x000401AF:
		return "InplaceSLast"
	case 0x800401B0:
		return "EnumEFirst"
	case 0x800401BF:
		return "EnumELast"
	case 0x000401B0:
		return "EnumSFirst"
	case 0x000401BF:
		return "EnumSLast"
	case 0x800401C0:
		return "Convert10EFirst"
	case 0x800401CF:
		return "Convert10ELast"
	case 0x000401C0:
		return "Convert10SFirst"
	case 0x000401CF:
		return "Convert10SLast"
	case 0x800401C0:
		return "Convert10EOlestreamGet"
	case 0x800401C1:
		return "Convert10EOlestreamPut"
	case 0x800401C2:
		return "Convert10EOlestreamFmt"
	case 0x800401C3:
		return "Convert10EOlestreamBitmapToDib"
	case 0x800401C4:
		return "Convert10EStgFmt"
	case 0x800401C5:
		return "Convert10EStgNoStdStream"
	case 0x800401C6:
		return "Convert10EStgDibToBitmap"
	case 0x800401D0:
		return "ClipbrdEFirst"
	case 0x800401DF:
		return "ClipbrdELast"
	case 0x000401D0:
		return "ClipbrdSFirst"
	case 0x000401DF:
		return "ClipbrdSLast"
	case 0x800401D0:
		return "ClipbrdECantOpen"
	case 0x800401D1:
		return "ClipbrdECantEmpty"
	case 0x800401D2:
		return "ClipbrdECantSet"
	case 0x800401D3:
		return "ClipbrdEBadData"
	case 0x800401D4:
		return "ClipbrdECantClose"
	case 0x800401E0:
		return "MkEFirst"
	case 0x800401EF:
		return "MkELast"
	case 0x000401E0:
		return "MkSFirst"
	case 0x000401EF:
		return "MkSLast"
	case 0x800401E0:
		return "MkEConnectmanually"
	case 0x800401E1:
		return "MkEExceededdeadline"
	case 0x800401E2:
		return "MkENeedgeneric"
	case 0x800401E3:
		return "MkEUnavailable"
	case 0x800401E4:
		return "MkESyntax"
	case 0x800401E5:
		return "MkENoobject"
	case 0x800401E6:
		return "MkEInvalidextension"
	case 0x800401E7:
		return "MkEIntermediateinterfacenotsupported"
	case 0x800401E8:
		return "MkENotbindable"
	case 0x800401E9:
		return "MkENotbound"
	case 0x800401EA:
		return "MkECantopenfile"
	case 0x800401EB:
		return "MkEMustbotheruser"
	case 0x800401EC:
		return "MkENoinverse"
	case 0x800401ED:
		return "MkENostorage"
	case 0x800401EE:
		return "MkENoprefix"
	case 0x800401EF:
		return "MkEEnumerationFailed"
	case 0x800401F0:
		return "CoEFirst"
	case 0x800401FF:
		return "CoELast"
	case 0x000401F0:
		return "CoSFirst"
	case 0x000401FF:
		return "CoSLast"
	case 0x800401F0:
		return "CoENotinitialized"
	case 0x800401F1:
		return "CoEAlreadyinitialized"
	case 0x800401F2:
		return "CoECantdetermineclass"
	case 0x800401F3:
		return "CoEClassstring"
	case 0x800401F4:
		return "CoEIidstring"
	case 0x800401F5:
		return "CoEAppnotfound"
	case 0x800401F6:
		return "CoEAppsingleuse"
	case 0x800401F7:
		return "CoEErrorinapp"
	case 0x800401F8:
		return "CoEDllnotfound"
	case 0x800401F9:
		return "CoEErrorindll"
	case 0x800401FA:
		return "CoEWrongosforapp"
	case 0x800401FB:
		return "CoEObjnotreg"
	case 0x800401FC:
		return "CoEObjisreg"
	case 0x800401FD:
		return "CoEObjnotconnected"
	case 0x800401FE:
		return "CoEAppdidntreg"
	case 0x800401FF:
		return "CoEReleased"
	case 0x80040200:
		return "EventEFirst"
	case 0x8004021F:
		return "EventELast"
	case 0x00040200:
		return "EventSFirst"
	case 0x0004021F:
		return "EventSLast"
	case 0x00040200:
		return "EventSSomeSubscribersFailed"
	case 0x80040201:
		return "EventEAllSubscribersFailed"
	case 0x00040202:
		return "EventSNosubscribers"
	case 0x80040203:
		return "EventEQuerysyntax"
	case 0x80040204:
		return "EventEQueryfield"
	case 0x80040205:
		return "EventEInternalexception"
	case 0x80040206:
		return "EventEInternalerror"
	case 0x80040207:
		return "EventEInvalidPerUserSid"
	case 0x80040208:
		return "EventEUserException"
	case 0x80040209:
		return "EventETooManyMethods"
	case 0x8004020A:
		return "EventEMissingEventclass"
	case 0x8004020B:
		return "EventENotAllRemoved"
	case 0x8004020C:
		return "EventEComplusNotInstalled"
	case 0x8004020D:
		return "EventECantModifyOrDeleteUnconfiguredObject"
	case 0x8004020E:
		return "EventECantModifyOrDeleteConfiguredObject"
	case 0x8004020F:
		return "EventEInvalidEventClassPartition"
	case 0x80040210:
		return "EventEPerUserSidNotLoggedOn"
	case 0x8004D000:
		return "XactEFirst"
	case 0x8004D02B:
		return "XactELast"
	case 0x0004D000:
		return "XactSFirst"
	case 0x0004D010:
		return "XactSLast"
	case 0x8004D000:
		return "XactEAlreadyothersinglephase"
	case 0x8004D001:
		return "XactECantretain"
	case 0x8004D002:
		return "XactECommitfailed"
	case 0x8004D003:
		return "XactECommitprevented"
	case 0x8004D004:
		return "XactEHeuristicabort"
	case 0x8004D005:
		return "XactEHeuristiccommit"
	case 0x8004D006:
		return "XactEHeuristicdamage"
	case 0x8004D007:
		return "XactEHeuristicdanger"
	case 0x8004D008:
		return "XactEIsolationlevel"
	case 0x8004D009:
		return "XactENoasync"
	case 0x8004D00A:
		return "XactENoenlist"
	case 0x8004D00B:
		return "XactENoisoretain"
	case 0x8004D00C:
		return "XactENoresource"
	case 0x8004D00D:
		return "XactENotcurrent"
	case 0x8004D00E:
		return "XactENotransaction"
	case 0x8004D00F:
		return "XactENotsupported"
	case 0x8004D010:
		return "XactEUnknownrmgrid"
	case 0x8004D011:
		return "XactEWrongstate"
	case 0x8004D012:
		return "XactEWronguow"
	case 0x8004D013:
		return "XactEXtionexists"
	case 0x8004D014:
		return "XactENoimportobject"
	case 0x8004D015:
		return "XactEInvalidcookie"
	case 0x8004D016:
		return "XactEIndoubt"
	case 0x8004D017:
		return "XactENotimeout"
	case 0x8004D018:
		return "XactEAlreadyinprogress"
	case 0x8004D019:
		return "XactEAborted"
	case 0x8004D01A:
		return "XactELogfull"
	case 0x8004D01B:
		return "XactETmnotavailable"
	case 0x8004D01C:
		return "XactEConnectionDown"
	case 0x8004D01D:
		return "XactEConnectionDenied"
	case 0x8004D01E:
		return "XactEReenlisttimeout"
	case 0x8004D01F:
		return "XactETipConnectFailed"
	case 0x8004D020:
		return "XactETipProtocolError"
	case 0x8004D021:
		return "XactETipPullFailed"
	case 0x8004D022:
		return "XactEDestTmnotavailable"
	case 0x8004D023:
		return "XactETipDisabled"
	case 0x8004D024:
		return "XactENetworkTxDisabled"
	case 0x8004D025:
		return "XactEPartnerNetworkTxDisabled"
	case 0x8004D026:
		return "XactEXaTxDisabled"
	case 0x8004D027:
		return "XactEUnableToReadDtcConfig"
	case 0x8004D028:
		return "XactEUnableToLoadDtcProxy"
	case 0x8004D029:
		return "XactEAborting"
	case 0x8004D02A:
		return "XactEPushCommFailure"
	case 0x8004D02B:
		return "XactEPullCommFailure"
	case 0x8004D02C:
		return "XactELuTxDisabled"
	case 0x8004D080:
		return "XactEClerknotfound"
	case 0x8004D081:
		return "XactEClerkexists"
	case 0x8004D082:
		return "XactERecoveryinprogress"
	case 0x8004D083:
		return "XactETransactionclosed"
	case 0x8004D084:
		return "XactEInvalidlsn"
	case 0x8004D085:
		return "XactEReplayrequest"
	case 0x0004D000:
		return "XactSAsync"
	case 0x0004D001:
		return "XactSDefect"
	case 0x0004D002:
		return "XactSReadonly"
	case 0x0004D003:
		return "XactSSomenoretain"
	case 0x0004D004:
		return "XactSOkinform"
	case 0x0004D005:
		return "XactSMadechangescontent"
	case 0x0004D006:
		return "XactSMadechangesinform"
	case 0x0004D007:
		return "XactSAllnoretain"
	case 0x0004D008:
		return "XactSAborting"
	case 0x0004D009:
		return "XactSSinglephase"
	case 0x0004D00A:
		return "XactSLocallyOk"
	case 0x0004D010:
		return "XactSLastresourcemanager"
	case 0x8004E000:
		return "ContextEFirst"
	case 0x8004E02F:
		return "ContextELast"
	case 0x0004E000:
		return "ContextSFirst"
	case 0x0004E02F:
		return "ContextSLast"
	case 0x8004E002:
		return "ContextEAborted"
	case 0x8004E003:
		return "ContextEAborting"
	case 0x8004E004:
		return "ContextENocontext"
	case 0x8004E005:
		return "ContextEWouldDeadlock"
	case 0x8004E006:
		return "ContextESynchTimeout"
	case 0x8004E007:
		return "ContextEOldref"
	case 0x8004E00C:
		return "ContextERolenotfound"
	case 0x8004E00F:
		return "ContextETmnotavailable"
	case 0x8004E021:
		return "CoEActivationfailed"
	case 0x8004E022:
		return "CoEActivationfailedEventlogged"
	case 0x8004E023:
		return "CoEActivationfailedCatalogerror"
	case 0x8004E024:
		return "CoEActivationfailedTimeout"
	case 0x8004E025:
		return "CoEInitializationfailed"
	case 0x8004E026:
		return "ContextENojit"
	case 0x8004E027:
		return "ContextENotransaction"
	case 0x8004E028:
		return "CoEThreadingmodelChanged"
	case 0x8004E029:
		return "CoENoiisintrinsics"
	case 0x8004E02A:
		return "CoENocookies"
	case 0x8004E02B:
		return "CoEDberror"
	case 0x8004E02C:
		return "CoENotpooled"
	case 0x8004E02D:
		return "CoENotconstructed"
	case 0x8004E02E:
		return "CoENosynchronization"
	case 0x8004E02F:
		return "CoEIsolevelmismatch"
	case 0x8004E030:
		return "CoECallOutOfTxScopeNotAllowed"
	case 0x8004E031:
		return "CoEExitTransactionScopeNotCalled"
	case 0x00040000:
		return "OleSUsereg"
	case 0x00040001:
		return "OleSStatic"
	case 0x00040002:
		return "OleSMacClipformat"
	case 0x00040100:
		return "DragdropSDrop"
	case 0x00040101:
		return "DragdropSCancel"
	case 0x00040102:
		return "DragdropSUsedefaultcursors"
	case 0x00040130:
		return "DataSSameformatetc"
	case 0x00040140:
		return "ViewSAlreadyFrozen"
	case 0x00040170:
		return "CacheSFormatetcNotsupported"
	case 0x00040171:
		return "CacheSSamecache"
	case 0x00040172:
		return "CacheSSomecachesNotupdated"
	case 0x00040180:
		return "OleobjSInvalidverb"
	case 0x00040181:
		return "OleobjSCannotDoverbNow"
	case 0x00040182:
		return "OleobjSInvalidhwnd"
	case 0x000401A0:
		return "InplaceSTruncated"
	case 0x000401C0:
		return "Convert10SNoPresentation"
	case 0x000401E2:
		return "MkSReducedToSelf"
	case 0x000401E4:
		return "MkSMe"
	case 0x000401E5:
		return "MkSHim"
	case 0x000401E6:
		return "MkSUs"
	case 0x000401E7:
		return "MkSMonikeralreadyregistered"
	case 0x00041300:
		return "SchedSTaskReady"
	case 0x00041301:
		return "SchedSTaskRunning"
	case 0x00041302:
		return "SchedSTaskDisabled"
	case 0x00041303:
		return "SchedSTaskHasNotRun"
	case 0x00041304:
		return "SchedSTaskNoMoreRuns"
	case 0x00041305:
		return "SchedSTaskNotScheduled"
	case 0x00041306:
		return "SchedSTaskTerminated"
	case 0x00041307:
		return "SchedSTaskNoValidTriggers"
	case 0x00041308:
		return "SchedSEventTrigger"
	case 0x80041309:
		return "SchedETriggerNotFound"
	case 0x8004130A:
		return "SchedETaskNotReady"
	case 0x8004130B:
		return "SchedETaskNotRunning"
	case 0x8004130C:
		return "SchedEServiceNotInstalled"
	case 0x8004130D:
		return "SchedECannotOpenTask"
	case 0x8004130E:
		return "SchedEInvalidTask"
	case 0x8004130F:
		return "SchedEAccountInformationNotSet"
	case 0x80041310:
		return "SchedEAccountNameNotFound"
	case 0x80041311:
		return "SchedEAccountDbaseCorrupt"
	case 0x80041312:
		return "SchedENoSecurityServices"
	case 0x80041313:
		return "SchedEUnknownObjectVersion"
	case 0x80041314:
		return "SchedEUnsupportedAccountOption"
	case 0x80041315:
		return "SchedEServiceNotRunning"
	case 0x80041316:
		return "SchedEUnexpectednode"
	case 0x80041317:
		return "SchedENamespace"
	case 0x80041318:
		return "SchedEInvalidvalue"
	case 0x80041319:
		return "SchedEMissingnode"
	case 0x8004131A:
		return "SchedEMalformedxml"
	case 0x0004131B:
		return "SchedSSomeTriggersFailed"
	case 0x0004131C:
		return "SchedSBatchLogonProblem"
	case 0x8004131D:
		return "SchedETooManyNodes"
	case 0x8004131E:
		return "SchedEPastEndBoundary"
	case 0x8004131F:
		return "SchedEAlreadyRunning"
	case 0x80041320:
		return "SchedEUserNotLoggedOn"
	case 0x80041321:
		return "SchedEInvalidTaskHash"
	case 0x80041322:
		return "SchedEServiceNotAvailable"
	case 0x80041323:
		return "SchedEServiceTooBusy"
	case 0x80041324:
		return "SchedETaskAttempted"
	case 0x00041325:
		return "SchedSTaskQueued"
	case 0x80041326:
		return "SchedETaskDisabled"
	case 0x80041327:
		return "SchedETaskNotV1Compat"
	case 0x80041328:
		return "SchedEStartOnDemand"
	case 0x80080001:
		return "CoEClassCreateFailed"
	case 0x80080002:
		return "CoEScmError"
	case 0x80080003:
		return "CoEScmRpcFailure"
	case 0x80080004:
		return "CoEBadPath"
	case 0x80080005:
		return "CoEServerExecFailure"
	case 0x80080006:
		return "CoEObjsrvRpcFailure"
	case 0x80080007:
		return "MkENoNormalized"
	case 0x80080008:
		return "CoEServerStopping"
	case 0x80080009:
		return "MemEInvalidRoot"
	case 0x80080010:
		return "MemEInvalidLink"
	case 0x80080011:
		return "MemEInvalidSize"
	case 0x00080012:
		return "CoSNotallinterfaces"
	case 0x00080013:
		return "CoSMachinenamenotfound"
	case 0x80080015:
		return "CoEMissingDisplayname"
	case 0x80080016:
		return "CoERunasValueMustBeAaa"
	case 0x80080017:
		return "CoEElevationDisabled"
	case 0x80020001:
		return "DispEUnknowninterface"
	case 0x80020003:
		return "DispEMembernotfound"
	case 0x80020004:
		return "DispEParamnotfound"
	case 0x80020005:
		return "DispETypemismatch"
	case 0x80020006:
		return "DispEUnknownname"
	case 0x80020007:
		return "DispENonamedargs"
	case 0x80020008:
		return "DispEBadvartype"
	case 0x80020009:
		return "DispEException"
	case 0x8002000A:
		return "DispEOverflow"
	case 0x8002000B:
		return "DispEBadindex"
	case 0x8002000C:
		return "DispEUnknownlcid"
	case 0x8002000D:
		return "DispEArrayislocked"
	case 0x8002000E:
		return "DispEBadparamcount"
	case 0x8002000F:
		return "DispEParamnotoptional"
	case 0x80020010:
		return "DispEBadcallee"
	case 0x80020011:
		return "DispENotacollection"
	case 0x80020012:
		return "DispEDivbyzero"
	case 0x80020013:
		return "DispEBuffertoosmall"
	case 0x80028016:
		return "TypeEBuffertoosmall"
	case 0x80028017:
		return "TypeEFieldnotfound"
	case 0x80028018:
		return "TypeEInvdataread"
	case 0x80028019:
		return "TypeEUnsupformat"
	case 0x8002801C:
		return "TypeERegistryaccess"
	case 0x8002801D:
		return "TypeELibnotregistered"
	case 0x80028027:
		return "TypeEUndefinedtype"
	case 0x80028028:
		return "TypeEQualifiednamedisallowed"
	case 0x80028029:
		return "TypeEInvalidstate"
	case 0x8002802A:
		return "TypeEWrongtypekind"
	case 0x8002802B:
		return "TypeEElementnotfound"
	case 0x8002802C:
		return "TypeEAmbiguousname"
	case 0x8002802D:
		return "TypeENameconflict"
	case 0x8002802E:
		return "TypeEUnknownlcid"
	case 0x8002802F:
		return "TypeEDllfunctionnotfound"
	case 0x800288BD:
		return "TypeEBadmodulekind"
	case 0x800288C5:
		return "TypeESizetoobig"
	case 0x800288C6:
		return "TypeEDuplicateid"
	case 0x800288CF:
		return "TypeEInvalidid"
	case 0x80028CA0:
		return "TypeETypemismatch"
	case 0x80028CA1:
		return "TypeEOutofbounds"
	case 0x80028CA2:
		return "TypeEIoerror"
	case 0x80028CA3:
		return "TypeECantcreatetmpfile"
	case 0x80029C4A:
		return "TypeECantloadlibrary"
	case 0x80029C83:
		return "TypeEInconsistentpropfuncs"
	case 0x80029C84:
		return "TypeECirculartype"
	case 0x80030001:
		return "StgEInvalidfunction"
	case 0x80030002:
		return "StgEFilenotfound"
	case 0x80030003:
		return "StgEPathnotfound"
	case 0x80030004:
		return "StgEToomanyopenfiles"
	case 0x80030005:
		return "StgEAccessdenied"
	case 0x80030006:
		return "StgEInvalidhandle"
	case 0x80030008:
		return "StgEInsufficientmemory"
	case 0x80030009:
		return "StgEInvalidpointer"
	case 0x80030012:
		return "StgENomorefiles"
	case 0x80030013:
		return "StgEDiskiswriteprotected"
	case 0x80030019:
		return "StgESeekerror"
	case 0x8003001D:
		return "StgEWritefault"
	case 0x8003001E:
		return "StgEReadfault"
	case 0x80030020:
		return "StgEShareviolation"
	case 0x80030021:
		return "StgELockviolation"
	case 0x80030050:
		return "StgEFilealreadyexists"
	case 0x80030057:
		return "StgEInvalidparameter"
	case 0x80030070:
		return "StgEMediumfull"
	case 0x800300F0:
		return "StgEPropsetmismatched"
	case 0x800300FA:
		return "StgEAbnormalapiexit"
	case 0x800300FB:
		return "StgEInvalidheader"
	case 0x800300FC:
		return "StgEInvalidname"
	case 0x800300FD:
		return "StgEUnknown"
	case 0x800300FE:
		return "StgEUnimplementedfunction"
	case 0x800300FF:
		return "StgEInvalidflag"
	case 0x80030100:
		return "StgEInuse"
	case 0x80030101:
		return "StgENotcurrent"
	case 0x80030102:
		return "StgEReverted"
	case 0x80030103:
		return "StgECantsave"
	case 0x80030104:
		return "StgEOldformat"
	case 0x80030105:
		return "StgEOlddll"
	case 0x80030106:
		return "StgESharerequired"
	case 0x80030107:
		return "StgENotfilebasedstorage"
	case 0x80030108:
		return "StgEExtantmarshallings"
	case 0x80030109:
		return "StgEDocfilecorrupt"
	case 0x80030110:
		return "StgEBadbaseaddress"
	case 0x80030111:
		return "StgEDocfiletoolarge"
	case 0x80030112:
		return "StgENotsimpleformat"
	case 0x80030201:
		return "StgEIncomplete"
	case 0x80030202:
		return "StgETerminated"
	case 0x00030200:
		return "StgSConverted"
	case 0x00030201:
		return "StgSBlock"
	case 0x00030202:
		return "StgSRetrynow"
	case 0x00030203:
		return "StgSMonitoring"
	case 0x00030204:
		return "StgSMultipleopens"
	case 0x00030205:
		return "StgSConsolidationfailed"
	case 0x00030206:
		return "StgSCannotconsolidate"
	case 0x80030305:
		return "StgEStatusCopyProtectionFailure"
	case 0x80030306:
		return "StgECssAuthenticationFailure"
	case 0x80030307:
		return "StgECssKeyNotPresent"
	case 0x80030308:
		return "StgECssKeyNotEstablished"
	case 0x80030309:
		return "StgECssScrambledSector"
	case 0x8003030A:
		return "StgECssRegionMismatch"
	case 0x8003030B:
		return "StgEResetsExhausted"
	case 0x80010001:
		return "RpcECallRejected"
	case 0x80010002:
		return "RpcECallCanceled"
	case 0x80010003:
		return "RpcECantpostInsendcall"
	case 0x80010004:
		return "RpcECantcalloutInasynccall"
	case 0x80010005:
		return "RpcECantcalloutInexternalcall"
	case 0x80010006:
		return "RpcEConnectionTerminated"
	case 0x80010007:
		return "RpcEServerDied"
	case 0x80010008:
		return "RpcEClientDied"
	case 0x80010009:
		return "RpcEInvalidDatapacket"
	case 0x8001000A:
		return "RpcECanttransmitCall"
	case 0x8001000B:
		return "RpcEClientCantmarshalData"
	case 0x8001000C:
		return "RpcEClientCantunmarshalData"
	case 0x8001000D:
		return "RpcEServerCantmarshalData"
	case 0x8001000E:
		return "RpcEServerCantunmarshalData"
	case 0x8001000F:
		return "RpcEInvalidData"
	case 0x80010010:
		return "RpcEInvalidParameter"
	case 0x80010011:
		return "RpcECantcalloutAgain"
	case 0x80010012:
		return "RpcEServerDiedDne"
	case 0x80010100:
		return "RpcESysCallFailed"
	case 0x80010101:
		return "RpcEOutOfResources"
	case 0x80010102:
		return "RpcEAttemptedMultithread"
	case 0x80010103:
		return "RpcENotRegistered"
	case 0x80010104:
		return "RpcEFault"
	case 0x80010105:
		return "RpcEServerfault"
	case 0x80010106:
		return "RpcEChangedMode"
	case 0x80010107:
		return "RpcEInvalidmethod"
	case 0x80010108:
		return "RpcEDisconnected"
	case 0x80010109:
		return "RpcERetry"
	case 0x8001010A:
		return "RpcEServercallRetrylater"
	case 0x8001010B:
		return "RpcEServercallRejected"
	case 0x8001010C:
		return "RpcEInvalidCalldata"
	case 0x8001010D:
		return "RpcECantcalloutIninputsynccall"
	case 0x8001010E:
		return "RpcEWrongThread"
	case 0x8001010F:
		return "RpcEThreadNotInit"
	case 0x80010110:
		return "RpcEVersionMismatch"
	case 0x80010111:
		return "RpcEInvalidHeader"
	case 0x80010112:
		return "RpcEInvalidExtension"
	case 0x80010113:
		return "RpcEInvalidIpid"
	case 0x80010114:
		return "RpcEInvalidObject"
	case 0x80010115:
		return "RpcSCallpending"
	case 0x80010116:
		return "RpcSWaitontimer"
	case 0x80010117:
		return "RpcECallComplete"
	case 0x80010118:
		return "RpcEUnsecureCall"
	case 0x80010119:
		return "RpcETooLate"
	case 0x8001011A:
		return "RpcENoGoodSecurityPackages"
	case 0x8001011B:
		return "RpcEAccessDenied"
	case 0x8001011C:
		return "RpcERemoteDisabled"
	case 0x8001011D:
		return "RpcEInvalidObjref"
	case 0x8001011E:
		return "RpcENoContext"
	case 0x8001011F:
		return "RpcETimeout"
	case 0x80010120:
		return "RpcENoSync"
	case 0x80010121:
		return "RpcEFullsicRequired"
	case 0x80010122:
		return "RpcEInvalidStdName"
	case 0x80010123:
		return "CoEFailedtoimpersonate"
	case 0x80010124:
		return "CoEFailedtogetsecctx"
	case 0x80010125:
		return "CoEFailedtoopenthreadtoken"
	case 0x80010126:
		return "CoEFailedtogettokeninfo"
	case 0x80010127:
		return "CoETrusteedoesntmatchclient"
	case 0x80010128:
		return "CoEFailedtoqueryclientblanket"
	case 0x80010129:
		return "CoEFailedtosetdacl"
	case 0x8001012A:
		return "CoEAccesscheckfailed"
	case 0x8001012B:
		return "CoENetaccessapifailed"
	case 0x8001012C:
		return "CoEWrongtrusteenamesyntax"
	case 0x8001012D:
		return "CoEInvalidsid"
	case 0x8001012E:
		return "CoEConversionfailed"
	case 0x8001012F:
		return "CoENomatchingsidfound"
	case 0x80010130:
		return "CoELookupaccsidfailed"
	case 0x80010131:
		return "CoENomatchingnamefound"
	case 0x80010132:
		return "CoELookupaccnamefailed"
	case 0x80010133:
		return "CoESetserlhndlfailed"
	case 0x80010134:
		return "CoEFailedtogetwindir"
	case 0x80010135:
		return "CoEPathtoolong"
	case 0x80010136:
		return "CoEFailedtogenuuid"
	case 0x80010137:
		return "CoEFailedtocreatefile"
	case 0x80010138:
		return "CoEFailedtoclosehandle"
	case 0x80010139:
		return "CoEExceedsysacllimit"
	case 0x8001013A:
		return "CoEAcesinwrongorder"
	case 0x8001013B:
		return "CoEIncompatiblestreamversion"
	case 0x8001013C:
		return "CoEFailedtoopenprocesstoken"
	case 0x8001013D:
		return "CoEDecodefailed"
	case 0x8001013F:
		return "CoEAcnotinitialized"
	case 0x80010140:
		return "CoECancelDisabled"
	case 0x8001FFFF:
		return "RpcEUnexpected"
	case 0xC0090001:
		return "ErrorAuditingDisabled"
	case 0xC0090002:
		return "ErrorAllSidsFiltered"
	case 0xC0090003:
		return "ErrorBizrulesNotEnabled"
	case 0x80090001:
		return "NteBadUid"
	case 0x80090002:
		return "NteBadHash"
	case 0x80090003:
		return "NteBadKey"
	case 0x80090004:
		return "NteBadLen"
	case 0x80090005:
		return "NteBadData"
	case 0x80090006:
		return "NteBadSignature"
	case 0x80090007:
		return "NteBadVer"
	case 0x80090008:
		return "NteBadAlgid"
	case 0x80090009:
		return "NteBadFlags"
	case 0x8009000A:
		return "NteBadType"
	case 0x8009000B:
		return "NteBadKeyState"
	case 0x8009000C:
		return "NteBadHashState"
	case 0x8009000D:
		return "NteNoKey"
	case 0x8009000E:
		return "NteNoMemory"
	case 0x8009000F:
		return "NteExists"
	case 0x80090010:
		return "NtePerm"
	case 0x80090011:
		return "NteNotFound"
	case 0x80090012:
		return "NteDoubleEncrypt"
	case 0x80090013:
		return "NteBadProvider"
	case 0x80090014:
		return "NteBadProvType"
	case 0x80090015:
		return "NteBadPublicKey"
	case 0x80090016:
		return "NteBadKeyset"
	case 0x80090017:
		return "NteProvTypeNotDef"
	case 0x80090018:
		return "NteProvTypeEntryBad"
	case 0x80090019:
		return "NteKeysetNotDef"
	case 0x8009001A:
		return "NteKeysetEntryBad"
	case 0x8009001B:
		return "NteProvTypeNoMatch"
	case 0x8009001C:
		return "NteSignatureFileBad"
	case 0x8009001D:
		return "NteProviderDllFail"
	case 0x8009001E:
		return "NteProvDllNotFound"
	case 0x8009001F:
		return "NteBadKeysetParam"
	case 0x80090020:
		return "NteFail"
	case 0x80090021:
		return "NteSysErr"
	case 0x80090022:
		return "NteSilentContext"
	case 0x80090023:
		return "NteTokenKeysetStorageFull"
	case 0x80090024:
		return "NteTemporaryProfile"
	case 0x80090025:
		return "NteFixedparameter"
	case 0x80090026:
		return "NteInvalidHandle"
	case 0x80090027:
		return "NteInvalidParameter"
	case 0x80090028:
		return "NteBufferTooSmall"
	case 0x80090029:
		return "NteNotSupported"
	case 0x8009002A:
		return "NteNoMoreItems"
	case 0x8009002B:
		return "NteBuffersOverlap"
	case 0x8009002C:
		return "NteDecryptionFailure"
	case 0x8009002D:
		return "NteInternalError"
	case 0x8009002E:
		return "NteUiRequired"
	case 0x8009002F:
		return "NteHmacNotSupported"
	case 0x80090300:
		return "SecEInsufficientMemory"
	case 0x80090301:
		return "SecEInvalidHandle"
	case 0x80090302:
		return "SecEUnsupportedFunction"
	case 0x80090303:
		return "SecETargetUnknown"
	case 0x80090304:
		return "SecEInternalError"
	case 0x80090305:
		return "SecESecpkgNotFound"
	case 0x80090306:
		return "SecENotOwner"
	case 0x80090307:
		return "SecECannotInstall"
	case 0x80090308:
		return "SecEInvalidToken"
	case 0x80090309:
		return "SecECannotPack"
	case 0x8009030A:
		return "SecEQopNotSupported"
	case 0x8009030B:
		return "SecENoImpersonation"
	case 0x8009030C:
		return "SecELogonDenied"
	case 0x8009030D:
		return "SecEUnknownCredentials"
	case 0x8009030E:
		return "SecENoCredentials"
	case 0x8009030F:
		return "SecEMessageAltered"
	case 0x80090310:
		return "SecEOutOfSequence"
	case 0x80090311:
		return "SecENoAuthenticatingAuthority"
	case 0x00090312:
		return "SecIContinueNeeded"
	case 0x00090313:
		return "SecICompleteNeeded"
	case 0x00090314:
		return "SecICompleteAndContinue"
	case 0x00090315:
		return "SecILocalLogon"
	case 0x80090316:
		return "SecEBadPkgid"
	case 0x80090317:
		return "SecEContextExpired"
	case 0x00090317:
		return "SecIContextExpired"
	case 0x80090318:
		return "SecEIncompleteMessage"
	case 0x80090320:
		return "SecEIncompleteCredentials"
	case 0x80090321:
		return "SecEBufferTooSmall"
	case 0x00090320:
		return "SecIIncompleteCredentials"
	case 0x00090321:
		return "SecIRenegotiate"
	case 0x80090322:
		return "SecEWrongPrincipal"
	case 0x00090323:
		return "SecINoLsaContext"
	case 0x80090324:
		return "SecETimeSkew"
	case 0x80090325:
		return "SecEUntrustedRoot"
	case 0x80090326:
		return "SecEIllegalMessage"
	case 0x80090327:
		return "SecECertUnknown"
	case 0x80090328:
		return "SecECertExpired"
	case 0x80090329:
		return "SecEEncryptFailure"
	case 0x80090330:
		return "SecEDecryptFailure"
	case 0x80090331:
		return "SecEAlgorithmMismatch"
	case 0x80090332:
		return "SecESecurityQosFailed"
	case 0x80090333:
		return "SecEUnfinishedContextDeleted"
	case 0x80090334:
		return "SecENoTgtReply"
	case 0x80090335:
		return "SecENoIpAddresses"
	case 0x80090336:
		return "SecEWrongCredentialHandle"
	case 0x80090337:
		return "SecECryptoSystemInvalid"
	case 0x80090338:
		return "SecEMaxReferralsExceeded"
	case 0x80090339:
		return "SecEMustBeKdc"
	case 0x8009033A:
		return "SecEStrongCryptoNotSupported"
	case 0x8009033B:
		return "SecETooManyPrincipals"
	case 0x8009033C:
		return "SecENoPaData"
	case 0x8009033D:
		return "SecEPkinitNameMismatch"
	case 0x8009033E:
		return "SecESmartcardLogonRequired"
	case 0x8009033F:
		return "SecEShutdownInProgress"
	case 0x80090340:
		return "SecEKdcInvalidRequest"
	case 0x80090341:
		return "SecEKdcUnableToRefer"
	case 0x80090342:
		return "SecEKdcUnknownEtype"
	case 0x80090343:
		return "SecEUnsupportedPreauth"
	case 0x80090345:
		return "SecEDelegationRequired"
	case 0x80090346:
		return "SecEBadBindings"
	case 0x80090347:
		return "SecEMultipleAccounts"
	case 0x80090348:
		return "SecENoKerbKey"
	case 0x80090349:
		return "SecECertWrongUsage"
	case 0x80090350:
		return "SecEDowngradeDetected"
	case 0x80090351:
		return "SecESmartcardCertRevoked"
	case 0x80090352:
		return "SecEIssuingCaUntrusted"
	case 0x80090353:
		return "SecERevocationOfflineC"
	case 0x80090354:
		return "SecEPkinitClientFailure"
	case 0x80090355:
		return "SecESmartcardCertExpired"
	case 0x80090356:
		return "SecENoS4UProtSupport"
	case 0x80090357:
		return "SecECrossrealmDelegationFailure"
	case 0x80090358:
		return "SecERevocationOfflineKdc"
	case 0x80090359:
		return "SecEIssuingCaUntrustedKdc"
	case 0x8009035A:
		return "SecEKdcCertExpired"
	case 0x8009035B:
		return "SecEKdcCertRevoked"
	case 0x0009035C:
		return "SecISignatureNeeded"
	case 0x8009035D:
		return "SecEInvalidParameter"
	case 0x8009035E:
		return "SecEDelegationPolicy"
	case 0x8009035F:
		return "SecEPolicyNltmOnly"
	case 0x00090360:
		return "SecINoRenegotiation"
	case 0x80090361:
		return "SecENoContext"
	case 0x80090362:
		return "SecEPku2UCertFailure"
	case 0x80090363:
		return "SecEMutualAuthFailed"
	case 0x80091001:
		return "CryptEMsgError"
	case 0x80091002:
		return "CryptEUnknownAlgo"
	case 0x80091003:
		return "CryptEOidFormat"
	case 0x80091004:
		return "CryptEInvalidMsgType"
	case 0x80091005:
		return "CryptEUnexpectedEncoding"
	case 0x80091006:
		return "CryptEAuthAttrMissing"
	case 0x80091007:
		return "CryptEHashValue"
	case 0x80091008:
		return "CryptEInvalidIndex"
	case 0x80091009:
		return "CryptEAlreadyDecrypted"
	case 0x8009100A:
		return "CryptENotDecrypted"
	case 0x8009100B:
		return "CryptERecipientNotFound"
	case 0x8009100C:
		return "CryptEControlType"
	case 0x8009100D:
		return "CryptEIssuerSerialnumber"
	case 0x8009100E:
		return "CryptESignerNotFound"
	case 0x8009100F:
		return "CryptEAttributesMissing"
	case 0x80091010:
		return "CryptEStreamMsgNotReady"
	case 0x80091011:
		return "CryptEStreamInsufficientData"
	case 0x00091012:
		return "CryptINewProtectionRequired"
	case 0x80092001:
		return "CryptEBadLen"
	case 0x80092002:
		return "CryptEBadEncode"
	case 0x80092003:
		return "CryptEFileError"
	case 0x80092004:
		return "CryptENotFound"
	case 0x80092005:
		return "CryptEExists"
	case 0x80092006:
		return "CryptENoProvider"
	case 0x80092007:
		return "CryptESelfSigned"
	case 0x80092008:
		return "CryptEDeletedPrev"
	case 0x80092009:
		return "CryptENoMatch"
	case 0x8009200A:
		return "CryptEUnexpectedMsgType"
	case 0x8009200B:
		return "CryptENoKeyProperty"
	case 0x8009200C:
		return "CryptENoDecryptCert"
	case 0x8009200D:
		return "CryptEBadMsg"
	case 0x8009200E:
		return "CryptENoSigner"
	case 0x8009200F:
		return "CryptEPendingClose"
	case 0x80092010:
		return "CryptERevoked"
	case 0x80092011:
		return "CryptENoRevocationDll"
	case 0x80092012:
		return "CryptENoRevocationCheck"
	case 0x80092013:
		return "CryptERevocationOffline"
	case 0x80092014:
		return "CryptENotInRevocationDatabase"
	case 0x80092020:
		return "CryptEInvalidNumericString"
	case 0x80092021:
		return "CryptEInvalidPrintableString"
	case 0x80092022:
		return "CryptEInvalidIa5String"
	case 0x80092023:
		return "CryptEInvalidX500String"
	case 0x80092024:
		return "CryptENotCharString"
	case 0x80092025:
		return "CryptEFileresized"
	case 0x80092026:
		return "CryptESecuritySettings"
	case 0x80092027:
		return "CryptENoVerifyUsageDll"
	case 0x80092028:
		return "CryptENoVerifyUsageCheck"
	case 0x80092029:
		return "CryptEVerifyUsageOffline"
	case 0x8009202A:
		return "CryptENotInCtl"
	case 0x8009202B:
		return "CryptENoTrustedSigner"
	case 0x8009202C:
		return "CryptEMissingPubkeyPara"
	case 0x80093000:
		return "CryptEOssError"
	case 0x80093001:
		return "OssMoreBuf"
	case 0x80093002:
		return "OssNegativeUinteger"
	case 0x80093003:
		return "OssPduRange"
	case 0x80093004:
		return "OssMoreInput"
	case 0x80093005:
		return "OssDataError"
	case 0x80093006:
		return "OssBadArg"
	case 0x80093007:
		return "OssBadVersion"
	case 0x80093008:
		return "OssOutMemory"
	case 0x80093009:
		return "OssPduMismatch"
	case 0x8009300A:
		return "OssLimited"
	case 0x8009300B:
		return "OssBadPtr"
	case 0x8009300C:
		return "OssBadTime"
	case 0x8009300D:
		return "OssIndefiniteNotSupported"
	case 0x8009300E:
		return "OssMemError"
	case 0x8009300F:
		return "OssBadTable"
	case 0x80093010:
		return "OssTooLong"
	case 0x80093011:
		return "OssConstraintViolated"
	case 0x80093012:
		return "OssFatalError"
	case 0x80093013:
		return "OssAccessSerializationError"
	case 0x80093014:
		return "OssNullTbl"
	case 0x80093015:
		return "OssNullFcn"
	case 0x80093016:
		return "OssBadEncrules"
	case 0x80093017:
		return "OssUnavailEncrules"
	case 0x80093018:
		return "OssCantOpenTraceWindow"
	case 0x80093019:
		return "OssUnimplemented"
	case 0x8009301A:
		return "OssOidDllNotLinked"
	case 0x8009301B:
		return "OssCantOpenTraceFile"
	case 0x8009301C:
		return "OssTraceFileAlreadyOpen"
	case 0x8009301D:
		return "OssTableMismatch"
	case 0x8009301E:
		return "OssTypeNotSupported"
	case 0x8009301F:
		return "OssRealDllNotLinked"
	case 0x80093020:
		return "OssRealCodeNotLinked"
	case 0x80093021:
		return "OssOutOfRange"
	case 0x80093022:
		return "OssCopierDllNotLinked"
	case 0x80093023:
		return "OssConstraintDllNotLinked"
	case 0x80093024:
		return "OssComparatorDllNotLinked"
	case 0x80093025:
		return "OssComparatorCodeNotLinked"
	case 0x80093026:
		return "OssMemMgrDllNotLinked"
	case 0x80093027:
		return "OssPdvDllNotLinked"
	case 0x80093028:
		return "OssPdvCodeNotLinked"
	case 0x80093029:
		return "OssApiDllNotLinked"
	case 0x8009302A:
		return "OssBerderDllNotLinked"
	case 0x8009302B:
		return "OssPerDllNotLinked"
	case 0x8009302C:
		return "OssOpenTypeError"
	case 0x8009302D:
		return "OssMutexNotCreated"
	case 0x8009302E:
		return "OssCantCloseTraceFile"
	case 0x80093100:
		return "CryptEAsn1Error"
	case 0x80093101:
		return "CryptEAsn1Internal"
	case 0x80093102:
		return "CryptEAsn1Eod"
	case 0x80093103:
		return "CryptEAsn1Corrupt"
	case 0x80093104:
		return "CryptEAsn1Large"
	case 0x80093105:
		return "CryptEAsn1Constraint"
	case 0x80093106:
		return "CryptEAsn1Memory"
	case 0x80093107:
		return "CryptEAsn1Overflow"
	case 0x80093108:
		return "CryptEAsn1Badpdu"
	case 0x80093109:
		return "CryptEAsn1Badargs"
	case 0x8009310A:
		return "CryptEAsn1Badreal"
	case 0x8009310B:
		return "CryptEAsn1Badtag"
	case 0x8009310C:
		return "CryptEAsn1Choice"
	case 0x8009310D:
		return "CryptEAsn1Rule"
	case 0x8009310E:
		return "CryptEAsn1Utf8"
	case 0x80093133:
		return "CryptEAsn1PduType"
	case 0x80093134:
		return "CryptEAsn1Nyi"
	case 0x80093201:
		return "CryptEAsn1Extended"
	case 0x80093202:
		return "CryptEAsn1Noeod"
	case 0x80094001:
		return "CertsrvEBadRequestsubject"
	case 0x80094002:
		return "CertsrvENoRequest"
	case 0x80094003:
		return "CertsrvEBadRequeststatus"
	case 0x80094004:
		return "CertsrvEPropertyEmpty"
	case 0x80094005:
		return "CertsrvEInvalidCaCertificate"
	case 0x80094006:
		return "CertsrvEServerSuspended"
	case 0x80094007:
		return "CertsrvEEncodingLength"
	case 0x80094008:
		return "CertsrvERoleconflict"
	case 0x80094009:
		return "CertsrvERestrictedofficer"
	case 0x8009400A:
		return "CertsrvEKeyArchivalNotConfigured"
	case 0x8009400B:
		return "CertsrvENoValidKra"
	case 0x8009400C:
		return "CertsrvEBadRequestKeyArchival"
	case 0x8009400D:
		return "CertsrvENoCaadminDefined"
	case 0x8009400E:
		return "CertsrvEBadRenewalCertAttribute"
	case 0x8009400F:
		return "CertsrvENoDbSessions"
	case 0x80094010:
		return "CertsrvEAlignmentFault"
	case 0x80094011:
		return "CertsrvEEnrollDenied"
	case 0x80094012:
		return "CertsrvETemplateDenied"
	case 0x80094013:
		return "CertsrvEDownlevelDcSslOrUpgrade"
	case 0x80094014:
		return "CertsrvEAdminDeniedRequest"
	case 0x80094015:
		return "CertsrvENoPolicyServer"
	case 0x80094800:
		return "CertsrvEUnsupportedCertType"
	case 0x80094801:
		return "CertsrvENoCertType"
	case 0x80094802:
		return "CertsrvETemplateConflict"
	case 0x80094803:
		return "CertsrvESubjectAltNameRequired"
	case 0x80094804:
		return "CertsrvEArchivedKeyRequired"
	case 0x80094805:
		return "CertsrvESmimeRequired"
	case 0x80094806:
		return "CertsrvEBadRenewalSubject"
	case 0x80094807:
		return "CertsrvEBadTemplateVersion"
	case 0x80094808:
		return "CertsrvETemplatePolicyRequired"
	case 0x80094809:
		return "CertsrvESignaturePolicyRequired"
	case 0x8009480A:
		return "CertsrvESignatureCount"
	case 0x8009480B:
		return "CertsrvESignatureRejected"
	case 0x8009480C:
		return "CertsrvEIssuancePolicyRequired"
	case 0x8009480D:
		return "CertsrvESubjectUpnRequired"
	case 0x8009480E:
		return "CertsrvESubjectDirectoryGuidRequired"
	case 0x8009480F:
		return "CertsrvESubjectDnsRequired"
	case 0x80094810:
		return "CertsrvEArchivedKeyUnexpected"
	case 0x80094811:
		return "CertsrvEKeyLength"
	case 0x80094812:
		return "CertsrvESubjectEmailRequired"
	case 0x80094813:
		return "CertsrvEUnknownCertType"
	case 0x80094814:
		return "CertsrvECertTypeOverlap"
	case 0x80094815:
		return "CertsrvETooManySignatures"
	case 0x80095000:
		return "XenrollEKeyNotExportable"
	case 0x80095001:
		return "XenrollECannotAddRootCert"
	case 0x80095002:
		return "XenrollEResponseKaHashNotFound"
	case 0x80095003:
		return "XenrollEResponseUnexpectedKaHash"
	case 0x80095004:
		return "XenrollEResponseKaHashMismatch"
	case 0x80095005:
		return "XenrollEKeyspecSmimeMismatch"
	case 0x80096001:
		return "TrustESystemError"
	case 0x80096002:
		return "TrustENoSignerCert"
	case 0x80096003:
		return "TrustECounterSigner"
	case 0x80096004:
		return "TrustECertSignature"
	case 0x80096005:
		return "TrustETimeStamp"
	case 0x80096010:
		return "TrustEBadDigest"
	case 0x80096019:
		return "TrustEBasicConstraints"
	case 0x8009601E:
		return "TrustEFinancialCriteria"
	case 0x80097001:
		return "MssipotfEOutofmemrange"
	case 0x80097002:
		return "MssipotfECantgetobject"
	case 0x80097003:
		return "MssipotfENoheadtable"
	case 0x80097004:
		return "MssipotfEBadMagicnumber"
	case 0x80097005:
		return "MssipotfEBadOffsetTable"
	case 0x80097006:
		return "MssipotfETableTagorder"
	case 0x80097007:
		return "MssipotfETableLongword"
	case 0x80097008:
		return "MssipotfEBadFirstTablePlacement"
	case 0x80097009:
		return "MssipotfETablesOverlap"
	case 0x8009700A:
		return "MssipotfETablePadbytes"
	case 0x8009700B:
		return "MssipotfEFiletoosmall"
	case 0x8009700C:
		return "MssipotfETableChecksum"
	case 0x8009700D:
		return "MssipotfEFileChecksum"
	case 0x80097010:
		return "MssipotfEFailedPolicy"
	case 0x80097011:
		return "MssipotfEFailedHintsCheck"
	case 0x80097012:
		return "MssipotfENotOpentype"
	case 0x80097013:
		return "MssipotfEFile"
	case 0x80097014:
		return "MssipotfECrypt"
	case 0x80097015:
		return "MssipotfEBadversion"
	case 0x80097016:
		return "MssipotfEDsigStructure"
	case 0x80097017:
		return "MssipotfEPconstCheck"
	case 0x80097018:
		return "MssipotfEStructure"
	case 0x80097019:
		return "ErrorCredRequiresConfirmation"
	case 0:
		return "NteOpOk"
	case 0x800B0001:
		return "TrustEProviderUnknown"
	case 0x800B0002:
		return "TrustEActionUnknown"
	case 0x800B0003:
		return "TrustESubjectFormUnknown"
	case 0x800B0004:
		return "TrustESubjectNotTrusted"
	case 0x800B0005:
		return "DigsigEEncode"
	case 0x800B0006:
		return "DigsigEDecode"
	case 0x800B0007:
		return "DigsigEExtensibility"
	case 0x800B0008:
		return "DigsigECrypto"
	case 0x800B0009:
		return "PersistESizedefinite"
	case 0x800B000A:
		return "PersistESizeindefinite"
	case 0x800B000B:
		return "PersistENotselfsizing"
	case 0x800B0100:
		return "TrustENosignature"
	case 0x800B0101:
		return "CertEExpired"
	case 0x800B0102:
		return "CertEValidityperiodnesting"
	case 0x800B0103:
		return "CertERole"
	case 0x800B0104:
		return "CertEPathlenconst"
	case 0x800B0105:
		return "CertECritical"
	case 0x800B0106:
		return "CertEPurpose"
	case 0x800B0107:
		return "CertEIssuerchaining"
	case 0x800B0108:
		return "CertEMalformed"
	case 0x800B0109:
		return "CertEUntrustedroot"
	case 0x800B010A:
		return "CertEChaining"
	case 0x800B010B:
		return "TrustEFail"
	case 0x800B010C:
		return "CertERevoked"
	case 0x800B010D:
		return "CertEUntrustedtestroot"
	case 0x800B010E:
		return "CertERevocationFailure"
	case 0x800B010F:
		return "CertECnNoMatch"
	case 0x800B0110:
		return "CertEWrongUsage"
	case 0x800B0111:
		return "TrustEExplicitDistrust"
	case 0x800B0112:
		return "CertEUntrustedca"
	case 0x800B0113:
		return "CertEInvalidPolicy"
	case 0x800B0114:
		return "CertEInvalidName"
	case 0x800F0000:
		return "SpapiEExpectedSectionName"
	case 0x800F0001:
		return "SpapiEBadSectionNameLine"
	case 0x800F0002:
		return "SpapiESectionNameTooLong"
	case 0x800F0003:
		return "SpapiEGeneralSyntax"
	case 0x800F0100:
		return "SpapiEWrongInfStyle"
	case 0x800F0101:
		return "SpapiESectionNotFound"
	case 0x800F0102:
		return "SpapiELineNotFound"
	case 0x800F0103:
		return "SpapiENoBackup"
	case 0x800F0200:
		return "SpapiENoAssociatedClass"
	case 0x800F0201:
		return "SpapiEClassMismatch"
	case 0x800F0202:
		return "SpapiEDuplicateFound"
	case 0x800F0203:
		return "SpapiENoDriverSelected"
	case 0x800F0204:
		return "SpapiEKeyDoesNotExist"
	case 0x800F0205:
		return "SpapiEInvalidDevinstName"
	case 0x800F0206:
		return "SpapiEInvalidClass"
	case 0x800F0207:
		return "SpapiEDevinstAlreadyExists"
	case 0x800F0208:
		return "SpapiEDevinfoNotRegistered"
	case 0x800F0209:
		return "SpapiEInvalidRegProperty"
	case 0x800F020A:
		return "SpapiENoInf"
	case 0x800F020B:
		return "SpapiENoSuchDevinst"
	case 0x800F020C:
		return "SpapiECantLoadClassIcon"
	case 0x800F020D:
		return "SpapiEInvalidClassInstaller"
	case 0x800F020E:
		return "SpapiEDiDoDefault"
	case 0x800F020F:
		return "SpapiEDiNofilecopy"
	case 0x800F0210:
		return "SpapiEInvalidHwprofile"
	case 0x800F0211:
		return "SpapiENoDeviceSelected"
	case 0x800F0212:
		return "SpapiEDevinfoListLocked"
	case 0x800F0213:
		return "SpapiEDevinfoDataLocked"
	case 0x800F0214:
		return "SpapiEDiBadPath"
	case 0x800F0215:
		return "SpapiENoClassinstallParams"
	case 0x800F0216:
		return "SpapiEFilequeueLocked"
	case 0x800F0217:
		return "SpapiEBadServiceInstallsect"
	case 0x800F0218:
		return "SpapiENoClassDriverList"
	case 0x800F0219:
		return "SpapiENoAssociatedService"
	case 0x800F021A:
		return "SpapiENoDefaultDeviceInterface"
	case 0x800F021B:
		return "SpapiEDeviceInterfaceActive"
	case 0x800F021C:
		return "SpapiEDeviceInterfaceRemoved"
	case 0x800F021D:
		return "SpapiEBadInterfaceInstallsect"
	case 0x800F021E:
		return "SpapiENoSuchInterfaceClass"
	case 0x800F021F:
		return "SpapiEInvalidReferenceString"
	case 0x800F0220:
		return "SpapiEInvalidMachinename"
	case 0x800F0221:
		return "SpapiERemoteCommFailure"
	case 0x800F0222:
		return "SpapiEMachineUnavailable"
	case 0x800F0223:
		return "SpapiENoConfigmgrServices"
	case 0x800F0224:
		return "SpapiEInvalidProppageProvider"
	case 0x800F0225:
		return "SpapiENoSuchDeviceInterface"
	case 0x800F0226:
		return "SpapiEDiPostprocessingRequired"
	case 0x800F0227:
		return "SpapiEInvalidCoinstaller"
	case 0x800F0228:
		return "SpapiENoCompatDrivers"
	case 0x800F0229:
		return "SpapiENoDeviceIcon"
	case 0x800F022A:
		return "SpapiEInvalidInfLogconfig"
	case 0x800F022B:
		return "SpapiEDiDontInstall"
	case 0x800F022C:
		return "SpapiEInvalidFilterDriver"
	case 0x800F022D:
		return "SpapiENonWindowsNtDriver"
	case 0x800F022E:
		return "SpapiENonWindowsDriver"
	case 0x800F022F:
		return "SpapiENoCatalogForOemInf"
	case 0x800F0230:
		return "SpapiEDevinstallQueueNonnative"
	case 0x800F0231:
		return "SpapiENotDisableable"
	case 0x800F0232:
		return "SpapiECantRemoveDevinst"
	case 0x800F0233:
		return "SpapiEInvalidTarget"
	case 0x800F0234:
		return "SpapiEDriverNonnative"
	case 0x800F0235:
		return "SpapiEInWow64"
	case 0x800F0236:
		return "SpapiESetSystemRestorePoint"
	case 0x800F0237:
		return "SpapiEIncorrectlyCopiedInf"
	case 0x800F0238:
		return "SpapiESceDisabled"
	case 0x800F0239:
		return "SpapiEUnknownException"
	case 0x800F023A:
		return "SpapiEPnpRegistryError"
	case 0x800F023B:
		return "SpapiERemoteRequestUnsupported"
	case 0x800F023C:
		return "SpapiENotAnInstalledOemInf"
	case 0x800F023D:
		return "SpapiEInfInUseByDevices"
	case 0x800F023E:
		return "SpapiEDiFunctionObsolete"
	case 0x800F023F:
		return "SpapiENoAuthenticodeCatalog"
	case 0x800F0240:
		return "SpapiEAuthenticodeDisallowed"
	case 0x800F0241:
		return "SpapiEAuthenticodeTrustedPublisher"
	case 0x800F0242:
		return "SpapiEAuthenticodeTrustNotEstablished"
	case 0x800F0243:
		return "SpapiEAuthenticodePublisherNotTrusted"
	case 0x800F0244:
		return "SpapiESignatureOsattributeMismatch"
	case 0x800F0245:
		return "SpapiEOnlyValidateViaAuthenticode"
	case 0x800F0246:
		return "SpapiEDeviceInstallerNotReady"
	case 0x800F0247:
		return "SpapiEDriverStoreAddFailed"
	case 0x800F0248:
		return "SpapiEDeviceInstallBlocked"
	case 0x800F0249:
		return "SpapiEDriverInstallBlocked"
	case 0x800F024A:
		return "SpapiEWrongInfType"
	case 0x800F024B:
		return "SpapiEFileHashNotInCatalog"
	case 0x800F024C:
		return "SpapiEDriverStoreDeleteFailed"
	case 0x800F0300:
		return "SpapiEUnrecoverableStackOverflow"
	case 0x800F1000:
		return "SpapiEErrorNotInstalled"
	case 0x80100001:
		return "ScardFInternalError"
	case 0x80100002:
		return "ScardECancelled"
	case 0x80100003:
		return "ScardEInvalidHandle"
	case 0x80100004:
		return "ScardEInvalidParameter"
	case 0x80100005:
		return "ScardEInvalidTarget"
	case 0x80100006:
		return "ScardENoMemory"
	case 0x80100007:
		return "ScardFWaitedTooLong"
	case 0x80100008:
		return "ScardEInsufficientBuffer"
	case 0x80100009:
		return "ScardEUnknownReader"
	case 0x8010000A:
		return "ScardETimeout"
	case 0x8010000B:
		return "ScardESharingViolation"
	case 0x8010000C:
		return "ScardENoSmartcard"
	case 0x8010000D:
		return "ScardEUnknownCard"
	case 0x8010000E:
		return "ScardECantDispose"
	case 0x8010000F:
		return "ScardEProtoMismatch"
	case 0x80100010:
		return "ScardENotReady"
	case 0x80100011:
		return "ScardEInvalidValue"
	case 0x80100012:
		return "ScardESystemCancelled"
	case 0x80100013:
		return "ScardFCommError"
	case 0x80100014:
		return "ScardFUnknownError"
	case 0x80100015:
		return "ScardEInvalidAtr"
	case 0x80100016:
		return "ScardENotTransacted"
	case 0x80100017:
		return "ScardEReaderUnavailable"
	case 0x80100018:
		return "ScardPShutdown"
	case 0x80100019:
		return "ScardEPciTooSmall"
	case 0x8010001A:
		return "ScardEReaderUnsupported"
	case 0x8010001B:
		return "ScardEDuplicateReader"
	case 0x8010001C:
		return "ScardECardUnsupported"
	case 0x8010001D:
		return "ScardENoService"
	case 0x8010001E:
		return "ScardEServiceStopped"
	case 0x8010001F:
		return "ScardEUnexpected"
	case 0x80100020:
		return "ScardEIccInstallation"
	case 0x80100021:
		return "ScardEIccCreateorder"
	case 0x80100022:
		return "ScardEUnsupportedFeature"
	case 0x80100023:
		return "ScardEDirNotFound"
	case 0x80100024:
		return "ScardEFileNotFound"
	case 0x80100025:
		return "ScardENoDir"
	case 0x80100026:
		return "ScardENoFile"
	case 0x80100027:
		return "ScardENoAccess"
	case 0x80100028:
		return "ScardEWriteTooMany"
	case 0x80100029:
		return "ScardEBadSeek"
	case 0x8010002A:
		return "ScardEInvalidChv"
	case 0x8010002B:
		return "ScardEUnknownResMng"
	case 0x8010002C:
		return "ScardENoSuchCertificate"
	case 0x8010002D:
		return "ScardECertificateUnavailable"
	case 0x8010002E:
		return "ScardENoReadersAvailable"
	case 0x8010002F:
		return "ScardECommDataLost"
	case 0x80100030:
		return "ScardENoKeyContainer"
	case 0x80100031:
		return "ScardEServerTooBusy"
	case 0x80100032:
		return "ScardEPinCacheExpired"
	case 0x80100033:
		return "ScardENoPinCache"
	case 0x80100034:
		return "ScardEReadOnlyCard"
	case 0x80100065:
		return "ScardWUnsupportedCard"
	case 0x80100066:
		return "ScardWUnresponsiveCard"
	case 0x80100067:
		return "ScardWUnpoweredCard"
	case 0x80100068:
		return "ScardWResetCard"
	case 0x80100069:
		return "ScardWRemovedCard"
	case 0x8010006A:
		return "ScardWSecurityViolation"
	case 0x8010006B:
		return "ScardWWrongChv"
	case 0x8010006C:
		return "ScardWChvBlocked"
	case 0x8010006D:
		return "ScardWEof"
	case 0x8010006E:
		return "ScardWCancelledByUser"
	case 0x8010006F:
		return "ScardWCardNotAuthenticated"
	case 0x80100070:
		return "ScardWCacheItemNotFound"
	case 0x80100071:
		return "ScardWCacheItemStale"
	case 0x80100072:
		return "ScardWCacheItemTooBig"
	case 0x80110401:
		return "ComadminEObjecterrors"
	case 0x80110402:
		return "ComadminEObjectinvalid"
	case 0x80110403:
		return "ComadminEKeymissing"
	case 0x80110404:
		return "ComadminEAlreadyinstalled"
	case 0x80110407:
		return "ComadminEAppFileWritefail"
	case 0x80110408:
		return "ComadminEAppFileReadfail"
	case 0x80110409:
		return "ComadminEAppFileVersion"
	case 0x8011040A:
		return "ComadminEBadpath"
	case 0x8011040B:
		return "ComadminEApplicationexists"
	case 0x8011040C:
		return "ComadminERoleexists"
	case 0x8011040D:
		return "ComadminECantcopyfile"
	case 0x8011040F:
		return "ComadminENouser"
	case 0x80110410:
		return "ComadminEInvaliduserids"
	case 0x80110411:
		return "ComadminENoregistryclsid"
	case 0x80110412:
		return "ComadminEBadregistryprogid"
	case 0x80110413:
		return "ComadminEAuthenticationlevel"
	case 0x80110414:
		return "ComadminEUserpasswdnotvalid"
	case 0x80110418:
		return "ComadminEClsidoriidmismatch"
	case 0x80110419:
		return "ComadminERemoteinterface"
	case 0x8011041A:
		return "ComadminEDllregisterserver"
	case 0x8011041B:
		return "ComadminENoservershare"
	case 0x8011041D:
		return "ComadminEDllloadfailed"
	case 0x8011041E:
		return "ComadminEBadregistrylibid"
	case 0x8011041F:
		return "ComadminEAppdirnotfound"
	case 0x80110423:
		return "ComadminERegistrarfailed"
	case 0x80110424:
		return "ComadminECompfileDoesnotexist"
	case 0x80110425:
		return "ComadminECompfileLoaddllfail"
	case 0x80110426:
		return "ComadminECompfileGetclassobj"
	case 0x80110427:
		return "ComadminECompfileClassnotavail"
	case 0x80110428:
		return "ComadminECompfileBadtlb"
	case 0x80110429:
		return "ComadminECompfileNotinstallable"
	case 0x8011042A:
		return "ComadminENotchangeable"
	case 0x8011042B:
		return "ComadminENotdeleteable"
	case 0x8011042C:
		return "ComadminESession"
	case 0x8011042D:
		return "ComadminECompMoveLocked"
	case 0x8011042E:
		return "ComadminECompMoveBadDest"
	case 0x80110430:
		return "ComadminERegistertlb"
	case 0x80110433:
		return "ComadminESystemapp"
	case 0x80110434:
		return "ComadminECompfileNoregistrar"
	case 0x80110435:
		return "ComadminECoreqcompinstalled"
	case 0x80110436:
		return "ComadminEServicenotinstalled"
	case 0x80110437:
		return "ComadminEPropertysavefailed"
	case 0x80110438:
		return "ComadminEObjectexists"
	case 0x80110439:
		return "ComadminEComponentexists"
	case 0x8011043B:
		return "ComadminERegfileCorrupt"
	case 0x8011043C:
		return "ComadminEPropertyOverflow"
	case 0x8011043E:
		return "ComadminENotinregistry"
	case 0x8011043F:
		return "ComadminEObjectnotpoolable"
	case 0x80110446:
		return "ComadminEApplidMatchesClsid"
	case 0x80110447:
		return "ComadminERoleDoesNotExist"
	case 0x80110448:
		return "ComadminEStartAppNeedsComponents"
	case 0x80110449:
		return "ComadminERequiresDifferentPlatform"
	case 0x8011044A:
		return "ComadminECanNotExportAppProxy"
	case 0x8011044B:
		return "ComadminECanNotStartApp"
	case 0x8011044C:
		return "ComadminECanNotExportSysApp"
	case 0x8011044D:
		return "ComadminECantSubscribeToComponent"
	case 0x8011044E:
		return "ComadminEEventclassCantBeSubscriber"
	case 0x8011044F:
		return "ComadminELibAppProxyIncompatible"
	case 0x80110450:
		return "ComadminEBasePartitionOnly"
	case 0x80110451:
		return "ComadminEStartAppDisabled"
	case 0x80110457:
		return "ComadminECatDuplicatePartitionName"
	case 0x80110458:
		return "ComadminECatInvalidPartitionName"
	case 0x80110459:
		return "ComadminECatPartitionInUse"
	case 0x8011045A:
		return "ComadminEFilePartitionDuplicateFiles"
	case 0x8011045B:
		return "ComadminECatImportedComponentsNotAllowed"
	case 0x8011045C:
		return "ComadminEAmbiguousApplicationName"
	case 0x8011045D:
		return "ComadminEAmbiguousPartitionName"
	case 0x80110472:
		return "ComadminERegdbNotinitialized"
	case 0x80110473:
		return "ComadminERegdbNotopen"
	case 0x80110474:
		return "ComadminERegdbSystemerr"
	case 0x80110475:
		return "ComadminERegdbAlreadyrunning"
	case 0x80110480:
		return "ComadminEMigVersionnotsupported"
	case 0x80110481:
		return "ComadminEMigSchemanotfound"
	case 0x80110482:
		return "ComadminECatBitnessmismatch"
	case 0x80110483:
		return "ComadminECatUnacceptablebitness"
	case 0x80110484:
		return "ComadminECatWrongappbitness"
	case 0x80110485:
		return "ComadminECatPauseResumeNotSupported"
	case 0x80110486:
		return "ComadminECatServerfault"
	case 0x80110600:
		return "ComqcEApplicationNotQueued"
	case 0x80110601:
		return "ComqcENoQueueableInterfaces"
	case 0x80110602:
		return "ComqcEQueuingServiceNotAvailable"
	case 0x80110603:
		return "ComqcENoIpersiststream"
	case 0x80110604:
		return "ComqcEBadMessage"
	case 0x80110605:
		return "ComqcEUnauthenticated"
	case 0x80110606:
		return "ComqcEUntrustedEnqueuer"
	case 0x80110701:
		return "MsdtcEDuplicateResource"
	case 0x80110808:
		return "ComadminEObjectParentMissing"
	case 0x80110809:
		return "ComadminEObjectDoesNotExist"
	case 0x8011080A:
		return "ComadminEAppNotRunning"
	case 0x8011080B:
		return "ComadminEInvalidPartition"
	case 0x8011080D:
		return "ComadminESvcappNotPoolableOrRecyclable"
	case 0x8011080E:
		return "ComadminEUserInSet"
	case 0x8011080F:
		return "ComadminECantrecyclelibraryapps"
	case 0x80110811:
		return "ComadminECantrecycleserviceapps"
	case 0x80110812:
		return "ComadminEProcessalreadyrecycled"
	case 0x80110813:
		return "ComadminEPausedprocessmaynotberecycled"
	case 0x80110814:
		return "ComadminECantmakeinprocservice"
	case 0x80110815:
		return "ComadminEProgidinusebyclsid"
	case 0x80110816:
		return "ComadminEDefaultPartitionNotInSet"
	case 0x80110817:
		return "ComadminERecycledprocessmaynotbepaused"
	case 0x80110818:
		return "ComadminEPartitionAccessdenied"
	case 0x80110819:
		return "ComadminEPartitionMsiOnly"
	case 0x8011081A:
		return "ComadminELegacycompsNotAllowedIn10Format"
	case 0x8011081B:
		return "ComadminELegacycompsNotAllowedInNonbasePartitions"
	case 0x8011081C:
		return "ComadminECompMoveSource"
	case 0x8011081D:
		return "ComadminECompMoveDest"
	case 0x8011081E:
		return "ComadminECompMovePrivate"
	case 0x8011081F:
		return "ComadminEBasepartitionRequiredInSet"
	case 0x80110820:
		return "ComadminECannotAliasEventclass"
	case 0x80110821:
		return "ComadminEPrivateAccessdenied"
	case 0x80110822:
		return "ComadminESaferinvalid"
	case 0x80110823:
		return "ComadminERegistryAccessdenied"
	case 0x80110824:
		return "ComadminEPartitionsDisabled"
	case 0x001F0001:
		return "ErrorFltIoComplete"
	case 0x801F0001:
		return "ErrorFltNoHandlerDefined"
	case 0x801F0002:
		return "ErrorFltContextAlreadyDefined"
	case 0x801F0003:
		return "ErrorFltInvalidAsynchronousRequest"
	case 0x801F0004:
		return "ErrorFltDisallowFastIo"
	case 0x801F0005:
		return "ErrorFltInvalidNameRequest"
	case 0x801F0006:
		return "ErrorFltNotSafeToPostOperation"
	case 0x801F0007:
		return "ErrorFltNotInitialized"
	case 0x801F0008:
		return "ErrorFltFilterNotReady"
	case 0x801F0009:
		return "ErrorFltPostOperationCleanup"
	case 0x801F000A:
		return "ErrorFltInternalError"
	case 0x801F000B:
		return "ErrorFltDeletingObject"
	case 0x801F000C:
		return "ErrorFltMustBeNonpagedPool"
	case 0x801F000D:
		return "ErrorFltDuplicateEntry"
	case 0x801F000E:
		return "ErrorFltCbdqDisabled"
	case 0x801F000F:
		return "ErrorFltDoNotAttach"
	case 0x801F0010:
		return "ErrorFltDoNotDetach"
	case 0x801F0011:
		return "ErrorFltInstanceAltitudeCollision"
	case 0x801F0012:
		return "ErrorFltInstanceNameCollision"
	case 0x801F0013:
		return "ErrorFltFilterNotFound"
	case 0x801F0014:
		return "ErrorFltVolumeNotFound"
	case 0x801F0015:
		return "ErrorFltInstanceNotFound"
	case 0x801F0016:
		return "ErrorFltContextAllocationNotFound"
	case 0x801F0017:
		return "ErrorFltInvalidContextRegistration"
	case 0x801F0018:
		return "ErrorFltNameCacheMiss"
	case 0x801F0019:
		return "ErrorFltNoDeviceObject"
	case 0x801F001A:
		return "ErrorFltVolumeAlreadyMounted"
	case 0x801F001B:
		return "ErrorFltAlreadyEnlisted"
	case 0x801F001C:
		return "ErrorFltContextAlreadyLinked"
	case 0x801F0020:
		return "ErrorFltNoWaiterForReply"
	case 0x80260001:
		return "ErrorHungDisplayDriverThread"
	case 0x80263001:
		return "DwmECompositiondisabled"
	case 0x80263002:
		return "DwmERemotingNotSupported"
	case 0x80263003:
		return "DwmENoRedirectionSurfaceAvailable"
	case 0x80263004:
		return "DwmENotQueuingPresents"
	case 0x80263005:
		return "DwmEAdapterNotFound"
	case 0x00263005:
		return "DwmSGdiRedirectionSurface"
	case 0x00261001:
		return "ErrorMonitorNoDescriptor"
	case 0x00261002:
		return "ErrorMonitorUnknownDescriptorFormat"
	case 0xC0261003:
		return "ErrorMonitorInvalidDescriptorChecksum"
	case 0xC0261004:
		return "ErrorMonitorInvalidStandardTimingBlock"
	case 0xC0261005:
		return "ErrorMonitorWmiDatablockRegistrationFailed"
	case 0xC0261006:
		return "ErrorMonitorInvalidSerialNumberMondscBlock"
	case 0xC0261007:
		return "ErrorMonitorInvalidUserFriendlyMondscBlock"
	case 0xC0261008:
		return "ErrorMonitorNoMoreDescriptorData"
	case 0xC0261009:
		return "ErrorMonitorInvalidDetailedTimingBlock"
	case 0xC026100A:
		return "ErrorMonitorInvalidManufactureDate"
	case 0xC0262000:
		return "ErrorGraphicsNotExclusiveModeOwner"
	case 0xC0262001:
		return "ErrorGraphicsInsufficientDmaBuffer"
	case 0xC0262002:
		return "ErrorGraphicsInvalidDisplayAdapter"
	case 0xC0262003:
		return "ErrorGraphicsAdapterWasReset"
	case 0xC0262004:
		return "ErrorGraphicsInvalidDriverModel"
	case 0xC0262005:
		return "ErrorGraphicsPresentModeChanged"
	case 0xC0262006:
		return "ErrorGraphicsPresentOccluded"
	case 0xC0262007:
		return "ErrorGraphicsPresentDenied"
	case 0xC0262008:
		return "ErrorGraphicsCannotcolorconvert"
	case 0xC0262009:
		return "ErrorGraphicsDriverMismatch"
	case 0x4026200A:
		return "ErrorGraphicsPartialDataPopulated"
	case 0xC026200B:
		return "ErrorGraphicsPresentRedirectionDisabled"
	case 0xC026200C:
		return "ErrorGraphicsPresentUnoccluded"
	case 0xC0262100:
		return "ErrorGraphicsNoVideoMemory"
	case 0xC0262101:
		return "ErrorGraphicsCantLockMemory"
	case 0xC0262102:
		return "ErrorGraphicsAllocationBusy"
	case 0xC0262103:
		return "ErrorGraphicsTooManyReferences"
	case 0xC0262104:
		return "ErrorGraphicsTryAgainLater"
	case 0xC0262105:
		return "ErrorGraphicsTryAgainNow"
	case 0xC0262106:
		return "ErrorGraphicsAllocationInvalid"
	case 0xC0262107:
		return "ErrorGraphicsUnswizzlingApertureUnavailable"
	case 0xC0262108:
		return "ErrorGraphicsUnswizzlingApertureUnsupported"
	case 0xC0262109:
		return "ErrorGraphicsCantEvictPinnedAllocation"
	case 0xC0262110:
		return "ErrorGraphicsInvalidAllocationUsage"
	case 0xC0262111:
		return "ErrorGraphicsCantRenderLockedAllocation"
	case 0xC0262112:
		return "ErrorGraphicsAllocationClosed"
	case 0xC0262113:
		return "ErrorGraphicsInvalidAllocationInstance"
	case 0xC0262114:
		return "ErrorGraphicsInvalidAllocationHandle"
	case 0xC0262115:
		return "ErrorGraphicsWrongAllocationDevice"
	case 0xC0262116:
		return "ErrorGraphicsAllocationContentLost"
	case 0xC0262200:
		return "ErrorGraphicsGpuExceptionOnDevice"
	case 0xC0262300:
		return "ErrorGraphicsInvalidVidpnTopology"
	case 0xC0262301:
		return "ErrorGraphicsVidpnTopologyNotSupported"
	case 0xC0262302:
		return "ErrorGraphicsVidpnTopologyCurrentlyNotSupported"
	case 0xC0262303:
		return "ErrorGraphicsInvalidVidpn"
	case 0xC0262304:
		return "ErrorGraphicsInvalidVideoPresentSource"
	case 0xC0262305:
		return "ErrorGraphicsInvalidVideoPresentTarget"
	case 0xC0262306:
		return "ErrorGraphicsVidpnModalityNotSupported"
	case 0x00262307:
		return "ErrorGraphicsModeNotPinned"
	case 0xC0262308:
		return "ErrorGraphicsInvalidVidpnSourcemodeset"
	case 0xC0262309:
		return "ErrorGraphicsInvalidVidpnTargetmodeset"
	case 0xC026230A:
		return "ErrorGraphicsInvalidFrequency"
	case 0xC026230B:
		return "ErrorGraphicsInvalidActiveRegion"
	case 0xC026230C:
		return "ErrorGraphicsInvalidTotalRegion"
	case 0xC0262310:
		return "ErrorGraphicsInvalidVideoPresentSourceMode"
	case 0xC0262311:
		return "ErrorGraphicsInvalidVideoPresentTargetMode"
	case 0xC0262312:
		return "ErrorGraphicsPinnedModeMustRemainInSet"
	case 0xC0262313:
		return "ErrorGraphicsPathAlreadyInTopology"
	case 0xC0262314:
		return "ErrorGraphicsModeAlreadyInModeset"
	case 0xC0262315:
		return "ErrorGraphicsInvalidVideopresentsourceset"
	case 0xC0262316:
		return "ErrorGraphicsInvalidVideopresenttargetset"
	case 0xC0262317:
		return "ErrorGraphicsSourceAlreadyInSet"
	case 0xC0262318:
		return "ErrorGraphicsTargetAlreadyInSet"
	case 0xC0262319:
		return "ErrorGraphicsInvalidVidpnPresentPath"
	case 0xC026231A:
		return "ErrorGraphicsNoRecommendedVidpnTopology"
	case 0xC026231B:
		return "ErrorGraphicsInvalidMonitorFrequencyrangeset"
	case 0xC026231C:
		return "ErrorGraphicsInvalidMonitorFrequencyrange"
	case 0xC026231D:
		return "ErrorGraphicsFrequencyrangeNotInSet"
	case 0x0026231E:
		return "ErrorGraphicsNoPreferredMode"
	case 0xC026231F:
		return "ErrorGraphicsFrequencyrangeAlreadyInSet"
	case 0xC0262320:
		return "ErrorGraphicsStaleModeset"
	case 0xC0262321:
		return "ErrorGraphicsInvalidMonitorSourcemodeset"
	case 0xC0262322:
		return "ErrorGraphicsInvalidMonitorSourceMode"
	case 0xC0262323:
		return "ErrorGraphicsNoRecommendedFunctionalVidpn"
	case 0xC0262324:
		return "ErrorGraphicsModeIdMustBeUnique"
	case 0xC0262325:
		return "ErrorGraphicsEmptyAdapterMonitorModeSupportIntersection"
	case 0xC0262326:
		return "ErrorGraphicsVideoPresentTargetsLessThanSources"
	case 0xC0262327:
		return "ErrorGraphicsPathNotInTopology"
	case 0xC0262328:
		return "ErrorGraphicsAdapterMustHaveAtLeastOneSource"
	case 0xC0262329:
		return "ErrorGraphicsAdapterMustHaveAtLeastOneTarget"
	case 0xC026232A:
		return "ErrorGraphicsInvalidMonitordescriptorset"
	case 0xC026232B:
		return "ErrorGraphicsInvalidMonitordescriptor"
	case 0xC026232C:
		return "ErrorGraphicsMonitordescriptorNotInSet"
	case 0xC026232D:
		return "ErrorGraphicsMonitordescriptorAlreadyInSet"
	case 0xC026232E:
		return "ErrorGraphicsMonitordescriptorIdMustBeUnique"
	case 0xC026232F:
		return "ErrorGraphicsInvalidVidpnTargetSubsetType"
	case 0xC0262330:
		return "ErrorGraphicsResourcesNotRelated"
	case 0xC0262331:
		return "ErrorGraphicsSourceIdMustBeUnique"
	case 0xC0262332:
		return "ErrorGraphicsTargetIdMustBeUnique"
	case 0xC0262333:
		return "ErrorGraphicsNoAvailableVidpnTarget"
	case 0xC0262334:
		return "ErrorGraphicsMonitorCouldNotBeAssociatedWithAdapter"
	case 0xC0262335:
		return "ErrorGraphicsNoVidpnmgr"
	case 0xC0262336:
		return "ErrorGraphicsNoActiveVidpn"
	case 0xC0262337:
		return "ErrorGraphicsStaleVidpnTopology"
	case 0xC0262338:
		return "ErrorGraphicsMonitorNotConnected"
	case 0xC0262339:
		return "ErrorGraphicsSourceNotInTopology"
	case 0xC026233A:
		return "ErrorGraphicsInvalidPrimarysurfaceSize"
	case 0xC026233B:
		return "ErrorGraphicsInvalidVisibleregionSize"
	case 0xC026233C:
		return "ErrorGraphicsInvalidStride"
	case 0xC026233D:
		return "ErrorGraphicsInvalidPixelformat"
	case 0xC026233E:
		return "ErrorGraphicsInvalidColorbasis"
	case 0xC026233F:
		return "ErrorGraphicsInvalidPixelvalueaccessmode"
	case 0xC0262340:
		return "ErrorGraphicsTargetNotInTopology"
	case 0xC0262341:
		return "ErrorGraphicsNoDisplayModeManagementSupport"
	case 0xC0262342:
		return "ErrorGraphicsVidpnSourceInUse"
	case 0xC0262343:
		return "ErrorGraphicsCantAccessActiveVidpn"
	case 0xC0262344:
		return "ErrorGraphicsInvalidPathImportanceOrdinal"
	case 0xC0262345:
		return "ErrorGraphicsInvalidPathContentGeometryTransformation"
	case 0xC0262346:
		return "ErrorGraphicsPathContentGeometryTransformationNotSupported"
	case 0xC0262347:
		return "ErrorGraphicsInvalidGammaRamp"
	case 0xC0262348:
		return "ErrorGraphicsGammaRampNotSupported"
	case 0xC0262349:
		return "ErrorGraphicsMultisamplingNotSupported"
	case 0xC026234A:
		return "ErrorGraphicsModeNotInModeset"
	case 0x0026234B:
		return "ErrorGraphicsDatasetIsEmpty"
	case 0x0026234C:
		return "ErrorGraphicsNoMoreElementsInDataset"
	case 0xC026234D:
		return "ErrorGraphicsInvalidVidpnTopologyRecommendationReason"
	case 0xC026234E:
		return "ErrorGraphicsInvalidPathContentType"
	case 0xC026234F:
		return "ErrorGraphicsInvalidCopyprotectionType"
	case 0xC0262350:
		return "ErrorGraphicsUnassignedModesetAlreadyExists"
	case 0x00262351:
		return "ErrorGraphicsPathContentGeometryTransformationNotPinned"
	case 0xC0262352:
		return "ErrorGraphicsInvalidScanlineOrdering"
	case 0xC0262353:
		return "ErrorGraphicsTopologyChangesNotAllowed"
	case 0xC0262354:
		return "ErrorGraphicsNoAvailableImportanceOrdinals"
	case 0xC0262355:
		return "ErrorGraphicsIncompatiblePrivateFormat"
	case 0xC0262356:
		return "ErrorGraphicsInvalidModePruningAlgorithm"
	case 0xC0262357:
		return "ErrorGraphicsInvalidMonitorCapabilityOrigin"
	case 0xC0262358:
		return "ErrorGraphicsInvalidMonitorFrequencyrangeConstraint"
	case 0xC0262359:
		return "ErrorGraphicsMaxNumPathsReached"
	case 0xC026235A:
		return "ErrorGraphicsCancelVidpnTopologyAugmentation"
	case 0xC026235B:
		return "ErrorGraphicsInvalidClientType"
	case 0xC026235C:
		return "ErrorGraphicsClientvidpnNotSet"
	case 0xC0262400:
		return "ErrorGraphicsSpecifiedChildAlreadyConnected"
	case 0xC0262401:
		return "ErrorGraphicsChildDescriptorNotSupported"
	case 0x4026242F:
		return "ErrorGraphicsUnknownChildStatus"
	case 0xC0262430:
		return "ErrorGraphicsNotALinkedAdapter"
	case 0xC0262431:
		return "ErrorGraphicsLeadlinkNotEnumerated"
	case 0xC0262432:
		return "ErrorGraphicsChainlinksNotEnumerated"
	case 0xC0262433:
		return "ErrorGraphicsAdapterChainNotReady"
	case 0xC0262434:
		return "ErrorGraphicsChainlinksNotStarted"
	case 0xC0262435:
		return "ErrorGraphicsChainlinksNotPoweredOn"
	case 0xC0262436:
		return "ErrorGraphicsInconsistentDeviceLinkState"
	case 0x40262437:
		return "ErrorGraphicsLeadlinkStartDeferred"
	case 0xC0262438:
		return "ErrorGraphicsNotPostDeviceDriver"
	case 0x40262439:
		return "ErrorGraphicsPollingTooFrequently"
	case 0x4026243A:
		return "ErrorGraphicsStartDeferred"
	case 0xC026243B:
		return "ErrorGraphicsAdapterAccessNotExcluded"
	case 0xC0262500:
		return "ErrorGraphicsOpmNotSupported"
	case 0xC0262501:
		return "ErrorGraphicsCoppNotSupported"
	case 0xC0262502:
		return "ErrorGraphicsUabNotSupported"
	case 0xC0262503:
		return "ErrorGraphicsOpmInvalidEncryptedParameters"
	case 0xC0262505:
		return "ErrorGraphicsOpmNoVideoOutputsExist"
	case 0xC026250B:
		return "ErrorGraphicsOpmInternalError"
	case 0xC026250C:
		return "ErrorGraphicsOpmInvalidHandle"
	case 0xC026250E:
		return "ErrorGraphicsPvpInvalidCertificateLength"
	case 0xC026250F:
		return "ErrorGraphicsOpmSpanningModeEnabled"
	case 0xC0262510:
		return "ErrorGraphicsOpmTheaterModeEnabled"
	case 0xC0262511:
		return "ErrorGraphicsPvpHfsFailed"
	case 0xC0262512:
		return "ErrorGraphicsOpmInvalidSrm"
	case 0xC0262513:
		return "ErrorGraphicsOpmOutputDoesNotSupportHdcp"
	case 0xC0262514:
		return "ErrorGraphicsOpmOutputDoesNotSupportAcp"
	case 0xC0262515:
		return "ErrorGraphicsOpmOutputDoesNotSupportCgmsa"
	case 0xC0262516:
		return "ErrorGraphicsOpmHdcpSrmNeverSet"
	case 0xC0262517:
		return "ErrorGraphicsOpmResolutionTooHigh"
	case 0xC0262518:
		return "ErrorGraphicsOpmAllHdcpHardwareAlreadyInUse"
	case 0xC026251A:
		return "ErrorGraphicsOpmVideoOutputNoLongerExists"
	case 0xC026251B:
		return "ErrorGraphicsOpmSessionTypeChangeInProgress"
	case 0xC026251C:
		return "ErrorGraphicsOpmVideoOutputDoesNotHaveCoppSemantics"
	case 0xC026251D:
		return "ErrorGraphicsOpmInvalidInformationRequest"
	case 0xC026251E:
		return "ErrorGraphicsOpmDriverInternalError"
	case 0xC026251F:
		return "ErrorGraphicsOpmVideoOutputDoesNotHaveOpmSemantics"
	case 0xC0262520:
		return "ErrorGraphicsOpmSignalingNotSupported"
	case 0xC0262521:
		return "ErrorGraphicsOpmInvalidConfigurationRequest"
	case 0xC0262580:
		return "ErrorGraphicsI2CNotSupported"
	case 0xC0262581:
		return "ErrorGraphicsI2CDeviceDoesNotExist"
	case 0xC0262582:
		return "ErrorGraphicsI2CErrorTransmittingData"
	case 0xC0262583:
		return "ErrorGraphicsI2CErrorReceivingData"
	case 0xC0262584:
		return "ErrorGraphicsDdcciVcpNotSupported"
	case 0xC0262585:
		return "ErrorGraphicsDdcciInvalidData"
	case 0xC0262586:
		return "ErrorGraphicsDdcciMonitorReturnedInvalidTimingStatusByte"
	case 0xC0262587:
		return "ErrorGraphicsMcaInvalidCapabilitiesString"
	case 0xC0262588:
		return "ErrorGraphicsMcaInternalError"
	case 0xC0262589:
		return "ErrorGraphicsDdcciInvalidMessageCommand"
	case 0xC026258A:
		return "ErrorGraphicsDdcciInvalidMessageLength"
	case 0xC026258B:
		return "ErrorGraphicsDdcciInvalidMessageChecksum"
	case 0xC026258C:
		return "ErrorGraphicsInvalidPhysicalMonitorHandle"
	case 0xC026258D:
		return "ErrorGraphicsMonitorNoLongerExists"
	case 0xC02625D8:
		return "ErrorGraphicsDdcciCurrentCurrentValueGreaterThanMaximumValue"
	case 0xC02625D9:
		return "ErrorGraphicsMcaInvalidVcpVersion"
	case 0xC02625DA:
		return "ErrorGraphicsMcaMonitorViolatesMccsSpecification"
	case 0xC02625DB:
		return "ErrorGraphicsMcaMccsVersionMismatch"
	case 0xC02625DC:
		return "ErrorGraphicsMcaUnsupportedMccsVersion"
	case 0xC02625DE:
		return "ErrorGraphicsMcaInvalidTechnologyTypeReturned"
	case 0xC02625DF:
		return "ErrorGraphicsMcaUnsupportedColorTemperature"
	case 0xC02625E0:
		return "ErrorGraphicsOnlyConsoleSessionSupported"
	case 0xC02625E1:
		return "ErrorGraphicsNoDisplayDeviceCorrespondsToName"
	case 0xC02625E2:
		return "ErrorGraphicsDisplayDeviceNotAttachedToDesktop"
	case 0xC02625E3:
		return "ErrorGraphicsMirroringDevicesNotSupported"
	case 0xC02625E4:
		return "ErrorGraphicsInvalidPointer"
	case 0xC02625E5:
		return "ErrorGraphicsNoMonitorsCorrespondToDisplayDevice"
	case 0xC02625E6:
		return "ErrorGraphicsParameterArrayTooSmall"
	case 0xC02625E7:
		return "ErrorGraphicsInternalError"
	case 0xC02605E8:
		return "ErrorGraphicsSessionTypeChangeInProgress"
	case 0x80280000:
		return "TpmEErrorMask"
	case 0x80280001:
		return "TpmEAuthfail"
	case 0x80280002:
		return "TpmEBadindex"
	case 0x80280003:
		return "TpmEBadParameter"
	case 0x80280004:
		return "TpmEAuditfailure"
	case 0x80280005:
		return "TpmEClearDisabled"
	case 0x80280006:
		return "TpmEDeactivated"
	case 0x80280007:
		return "TpmEDisabled"
	case 0x80280008:
		return "TpmEDisabledCmd"
	case 0x80280009:
		return "TpmEFail"
	case 0x8028000A:
		return "TpmEBadOrdinal"
	case 0x8028000B:
		return "TpmEInstallDisabled"
	case 0x8028000C:
		return "TpmEInvalidKeyhandle"
	case 0x8028000D:
		return "TpmEKeynotfound"
	case 0x8028000E:
		return "TpmEInappropriateEnc"
	case 0x8028000F:
		return "TpmEMigratefail"
	case 0x80280010:
		return "TpmEInvalidPcrInfo"
	case 0x80280011:
		return "TpmENospace"
	case 0x80280012:
		return "TpmENosrk"
	case 0x80280013:
		return "TpmENotsealedBlob"
	case 0x80280014:
		return "TpmEOwnerSet"
	case 0x80280015:
		return "TpmEResources"
	case 0x80280016:
		return "TpmEShortrandom"
	case 0x80280017:
		return "TpmESize"
	case 0x80280018:
		return "TpmEWrongpcrval"
	case 0x80280019:
		return "TpmEBadParamSize"
	case 0x8028001A:
		return "TpmEShaThread"
	case 0x8028001B:
		return "TpmEShaError"
	case 0x8028001C:
		return "TpmEFailedselftest"
	case 0x8028001D:
		return "TpmEAuth2Fail"
	case 0x8028001E:
		return "TpmEBadtag"
	case 0x8028001F:
		return "TpmEIoerror"
	case 0x80280020:
		return "TpmEEncryptError"
	case 0x80280021:
		return "TpmEDecryptError"
	case 0x80280022:
		return "TpmEInvalidAuthhandle"
	case 0x80280023:
		return "TpmENoEndorsement"
	case 0x80280024:
		return "TpmEInvalidKeyusage"
	case 0x80280025:
		return "TpmEWrongEntitytype"
	case 0x80280026:
		return "TpmEInvalidPostinit"
	case 0x80280027:
		return "TpmEInappropriateSig"
	case 0x80280028:
		return "TpmEBadKeyProperty"
	case 0x80280029:
		return "TpmEBadMigration"
	case 0x8028002A:
		return "TpmEBadScheme"
	case 0x8028002B:
		return "TpmEBadDatasize"
	case 0x8028002C:
		return "TpmEBadMode"
	case 0x8028002D:
		return "TpmEBadPresence"
	case 0x8028002E:
		return "TpmEBadVersion"
	case 0x8028002F:
		return "TpmENoWrapTransport"
	case 0x80280030:
		return "TpmEAuditfailUnsuccessful"
	case 0x80280031:
		return "TpmEAuditfailSuccessful"
	case 0x80280032:
		return "TpmENotresetable"
	case 0x80280033:
		return "TpmENotlocal"
	case 0x80280034:
		return "TpmEBadType"
	case 0x80280035:
		return "TpmEInvalidResource"
	case 0x80280036:
		return "TpmENotfips"
	case 0x80280037:
		return "TpmEInvalidFamily"
	case 0x80280038:
		return "TpmENoNvPermission"
	case 0x80280039:
		return "TpmERequiresSign"
	case 0x8028003A:
		return "TpmEKeyNotsupported"
	case 0x8028003B:
		return "TpmEAuthConflict"
	case 0x8028003C:
		return "TpmEAreaLocked"
	case 0x8028003D:
		return "TpmEBadLocality"
	case 0x8028003E:
		return "TpmEReadOnly"
	case 0x8028003F:
		return "TpmEPerNowrite"
	case 0x80280040:
		return "TpmEFamilycount"
	case 0x80280041:
		return "TpmEWriteLocked"
	case 0x80280042:
		return "TpmEBadAttributes"
	case 0x80280043:
		return "TpmEInvalidStructure"
	case 0x80280044:
		return "TpmEKeyOwnerControl"
	case 0x80280045:
		return "TpmEBadCounter"
	case 0x80280046:
		return "TpmENotFullwrite"
	case 0x80280047:
		return "TpmEContextGap"
	case 0x80280048:
		return "TpmEMaxnvwrites"
	case 0x80280049:
		return "TpmENooperator"
	case 0x8028004A:
		return "TpmEResourcemissing"
	case 0x8028004B:
		return "TpmEDelegateLock"
	case 0x8028004C:
		return "TpmEDelegateFamily"
	case 0x8028004D:
		return "TpmEDelegateAdmin"
	case 0x8028004E:
		return "TpmETransportNotexclusive"
	case 0x8028004F:
		return "TpmEOwnerControl"
	case 0x80280050:
		return "TpmEDaaResources"
	case 0x80280051:
		return "TpmEDaaInputData0"
	case 0x80280052:
		return "TpmEDaaInputData1"
	case 0x80280053:
		return "TpmEDaaIssuerSettings"
	case 0x80280054:
		return "TpmEDaaTpmSettings"
	case 0x80280055:
		return "TpmEDaaStage"
	case 0x80280056:
		return "TpmEDaaIssuerValidity"
	case 0x80280057:
		return "TpmEDaaWrongW"
	case 0x80280058:
		return "TpmEBadHandle"
	case 0x80280059:
		return "TpmEBadDelegate"
	case 0x8028005A:
		return "TpmEBadcontext"
	case 0x8028005B:
		return "TpmEToomanycontexts"
	case 0x8028005C:
		return "TpmEMaTicketSignature"
	case 0x8028005D:
		return "TpmEMaDestination"
	case 0x8028005E:
		return "TpmEMaSource"
	case 0x8028005F:
		return "TpmEMaAuthority"
	case 0x80280061:
		return "TpmEPermanentek"
	case 0x80280062:
		return "TpmEBadSignature"
	case 0x80280063:
		return "TpmENocontextspace"
	case 0x80280400:
		return "TpmECommandBlocked"
	case 0x80280401:
		return "TpmEInvalidHandle"
	case 0x80280402:
		return "TpmEDuplicateVhandle"
	case 0x80280403:
		return "TpmEEmbeddedCommandBlocked"
	case 0x80280404:
		return "TpmEEmbeddedCommandUnsupported"
	case 0x80280800:
		return "TpmERetry"
	case 0x80280801:
		return "TpmENeedsSelftest"
	case 0x80280802:
		return "TpmEDoingSelftest"
	case 0x80280803:
		return "TpmEDefendLockRunning"
	case 0x80284001:
		return "TbsEInternalError"
	case 0x80284002:
		return "TbsEBadParameter"
	case 0x80284003:
		return "TbsEInvalidOutputPointer"
	case 0x80284004:
		return "TbsEInvalidContext"
	case 0x80284005:
		return "TbsEInsufficientBuffer"
	case 0x80284006:
		return "TbsEIoerror"
	case 0x80284007:
		return "TbsEInvalidContextParam"
	case 0x80284008:
		return "TbsEServiceNotRunning"
	case 0x80284009:
		return "TbsETooManyTbsContexts"
	case 0x8028400A:
		return "TbsETooManyResources"
	case 0x8028400B:
		return "TbsEServiceStartPending"
	case 0x8028400C:
		return "TbsEPpiNotSupported"
	case 0x8028400D:
		return "TbsECommandCanceled"
	case 0x8028400E:
		return "TbsEBufferTooLarge"
	case 0x8028400F:
		return "TbsETpmNotFound"
	case 0x80284010:
		return "TbsEServiceDisabled"
	case 0x80284011:
		return "TbsENoEventLog"
	case 0x80290100:
		return "TpmapiEInvalidState"
	case 0x80290101:
		return "TpmapiENotEnoughData"
	case 0x80290102:
		return "TpmapiETooMuchData"
	case 0x80290103:
		return "TpmapiEInvalidOutputPointer"
	case 0x80290104:
		return "TpmapiEInvalidParameter"
	case 0x80290105:
		return "TpmapiEOutOfMemory"
	case 0x80290106:
		return "TpmapiEBufferTooSmall"
	case 0x80290107:
		return "TpmapiEInternalError"
	case 0x80290108:
		return "TpmapiEAccessDenied"
	case 0x80290109:
		return "TpmapiEAuthorizationFailed"
	case 0x8029010A:
		return "TpmapiEInvalidContextHandle"
	case 0x8029010B:
		return "TpmapiETbsCommunicationError"
	case 0x8029010C:
		return "TpmapiETpmCommandError"
	case 0x8029010D:
		return "TpmapiEMessageTooLarge"
	case 0x8029010E:
		return "TpmapiEInvalidEncoding"
	case 0x8029010F:
		return "TpmapiEInvalidKeySize"
	case 0x80290110:
		return "TpmapiEEncryptionFailed"
	case 0x80290111:
		return "TpmapiEInvalidKeyParams"
	case 0x80290112:
		return "TpmapiEInvalidMigrationAuthorizationBlob"
	case 0x80290113:
		return "TpmapiEInvalidPcrIndex"
	case 0x80290114:
		return "TpmapiEInvalidDelegateBlob"
	case 0x80290115:
		return "TpmapiEInvalidContextParams"
	case 0x80290116:
		return "TpmapiEInvalidKeyBlob"
	case 0x80290117:
		return "TpmapiEInvalidPcrData"
	case 0x80290118:
		return "TpmapiEInvalidOwnerAuth"
	case 0x80290119:
		return "TpmapiEFipsRngCheckFailed"
	case 0x8029011A:
		return "TpmapiEEmptyTcgLog"
	case 0x8029011B:
		return "TpmapiEInvalidTcgLogEntry"
	case 0x8029011C:
		return "TpmapiETcgSeparatorAbsent"
	case 0x8029011D:
		return "TpmapiETcgInvalidDigestEntry"
	case 0x80290200:
		return "TbsimpEBufferTooSmall"
	case 0x80290201:
		return "TbsimpECleanupFailed"
	case 0x80290202:
		return "TbsimpEInvalidContextHandle"
	case 0x80290203:
		return "TbsimpEInvalidContextParam"
	case 0x80290204:
		return "TbsimpETpmError"
	case 0x80290205:
		return "TbsimpEHashBadKey"
	case 0x80290206:
		return "TbsimpEDuplicateVhandle"
	case 0x80290207:
		return "TbsimpEInvalidOutputPointer"
	case 0x80290208:
		return "TbsimpEInvalidParameter"
	case 0x80290209:
		return "TbsimpERpcInitFailed"
	case 0x8029020A:
		return "TbsimpESchedulerNotRunning"
	case 0x8029020B:
		return "TbsimpECommandCanceled"
	case 0x8029020C:
		return "TbsimpEOutOfMemory"
	case 0x8029020D:
		return "TbsimpEListNoMoreItems"
	case 0x8029020E:
		return "TbsimpEListNotFound"
	case 0x8029020F:
		return "TbsimpENotEnoughSpace"
	case 0x80290210:
		return "TbsimpENotEnoughTpmContexts"
	case 0x80290211:
		return "TbsimpECommandFailed"
	case 0x80290212:
		return "TbsimpEUnknownOrdinal"
	case 0x80290213:
		return "TbsimpEResourceExpired"
	case 0x80290214:
		return "TbsimpEInvalidResource"
	case 0x80290215:
		return "TbsimpENothingToUnload"
	case 0x80290216:
		return "TbsimpEHashTableFull"
	case 0x80290217:
		return "TbsimpETooManyTbsContexts"
	case 0x80290218:
		return "TbsimpETooManyResources"
	case 0x80290219:
		return "TbsimpEPpiNotSupported"
	case 0x8029021A:
		return "TbsimpETpmIncompatible"
	case 0x8029021B:
		return "TbsimpENoEventLog"
	case 0x80290300:
		return "TpmEPpiAcpiFailure"
	case 0x80290301:
		return "TpmEPpiUserAbort"
	case 0x80290302:
		return "TpmEPpiBiosFailure"
	case 0x80290303:
		return "TpmEPpiNotSupported"
	case 0x80300002:
		return "PlaEDcsNotFound"
	case 0x803000AA:
		return "PlaEDcsInUse"
	case 0x80300045:
		return "PlaETooManyFolders"
	case 0x80300070:
		return "PlaENoMinDisk"
	case 0x803000B7:
		return "PlaEDcsAlreadyExists"
	case 0x00300100:
		return "PlaSPropertyIgnored"
	case 0x80300101:
		return "PlaEPropertyConflict"
	case 0x80300102:
		return "PlaEDcsSingletonRequired"
	case 0x80300103:
		return "PlaECredentialsRequired"
	case 0x80300104:
		return "PlaEDcsNotRunning"
	case 0x80300105:
		return "PlaEConflictInclExclApi"
	case 0x80300106:
		return "PlaENetworkExeNotValid"
	case 0x80300107:
		return "PlaEExeAlreadyConfigured"
	case 0x80300108:
		return "PlaEExePathNotValid"
	case 0x80300109:
		return "PlaEDcAlreadyExists"
	case 0x8030010A:
		return "PlaEDcsStartWaitTimeout"
	case 0x8030010B:
		return "PlaEDcStartWaitTimeout"
	case 0x8030010C:
		return "PlaEReportWaitTimeout"
	case 0x8030010D:
		return "PlaENoDuplicates"
	case 0x8030010E:
		return "PlaEExeFullPathRequired"
	case 0x8030010F:
		return "PlaEInvalidSessionName"
	case 0x80300110:
		return "PlaEPlaChannelNotEnabled"
	case 0x80300111:
		return "PlaETaskschedChannelNotEnabled"
	case 0x80300112:
		return "PlaERulesManagerFailed"
	case 0x80300113:
		return "PlaECabapiFailure"
	case 0x80310000:
		return "FveELockedVolume"
	case 0x80310001:
		return "FveENotEncrypted"
	case 0x80310002:
		return "FveENoTpmBios"
	case 0x80310003:
		return "FveENoMbrMetric"
	case 0x80310004:
		return "FveENoBootsectorMetric"
	case 0x80310005:
		return "FveENoBootmgrMetric"
	case 0x80310006:
		return "FveEWrongBootmgr"
	case 0x80310007:
		return "FveESecureKeyRequired"
	case 0x80310008:
		return "FveENotActivated"
	case 0x80310009:
		return "FveEActionNotAllowed"
	case 0x8031000A:
		return "FveEAdSchemaNotInstalled"
	case 0x8031000B:
		return "FveEAdInvalidDatatype"
	case 0x8031000C:
		return "FveEAdInvalidDatasize"
	case 0x8031000D:
		return "FveEAdNoValues"
	case 0x8031000E:
		return "FveEAdAttrNotSet"
	case 0x8031000F:
		return "FveEAdGuidNotFound"
	case 0x80310010:
		return "FveEBadInformation"
	case 0x80310011:
		return "FveETooSmall"
	case 0x80310012:
		return "FveESystemVolume"
	case 0x80310013:
		return "FveEFailedWrongFs"
	case 0x80310014:
		return "FveEBadPartitionSize"
	case 0x80310015:
		return "FveENotSupported"
	case 0x80310016:
		return "FveEBadData"
	case 0x80310017:
		return "FveEVolumeNotBound"
	case 0x80310018:
		return "FveETpmNotOwned"
	case 0x80310019:
		return "FveENotDataVolume"
	case 0x8031001A:
		return "FveEAdInsufficientBuffer"
	case 0x8031001B:
		return "FveEConvRead"
	case 0x8031001C:
		return "FveEConvWrite"
	case 0x8031001D:
		return "FveEKeyRequired"
	case 0x8031001E:
		return "FveEClusteringNotSupported"
	case 0x8031001F:
		return "FveEVolumeBoundAlready"
	case 0x80310020:
		return "FveEOsNotProtected"
	case 0x80310021:
		return "FveEProtectionDisabled"
	case 0x80310022:
		return "FveERecoveryKeyRequired"
	case 0x80310023:
		return "FveEForeignVolume"
	case 0x80310024:
		return "FveEOverlappedUpdate"
	case 0x80310025:
		return "FveETpmSrkAuthNotZero"
	case 0x80310026:
		return "FveEFailedSectorSize"
	case 0x80310027:
		return "FveEFailedAuthentication"
	case 0x80310028:
		return "FveENotOsVolume"
	case 0x80310029:
		return "FveEAutounlockEnabled"
	case 0x8031002A:
		return "FveEWrongBootsector"
	case 0x8031002B:
		return "FveEWrongSystemFs"
	case 0x8031002C:
		return "FveEPolicyPasswordRequired"
	case 0x8031002D:
		return "FveECannotSetFvekEncrypted"
	case 0x8031002E:
		return "FveECannotEncryptNoKey"
	case 0x80310030:
		return "FveEBootableCddvd"
	case 0x80310031:
		return "FveEProtectorExists"
	case 0x80310032:
		return "FveERelativePath"
	case 0x80310033:
		return "FveEProtectorNotFound"
	case 0x80310034:
		return "FveEInvalidKeyFormat"
	case 0x80310035:
		return "FveEInvalidPasswordFormat"
	case 0x80310036:
		return "FveEFipsRngCheckFailed"
	case 0x80310037:
		return "FveEFipsPreventsRecoveryPassword"
	case 0x80310038:
		return "FveEFipsPreventsExternalKeyExport"
	case 0x80310039:
		return "FveENotDecrypted"
	case 0x8031003A:
		return "FveEInvalidProtectorType"
	case 0x8031003B:
		return "FveENoProtectorsToTest"
	case 0x8031003C:
		return "FveEKeyfileNotFound"
	case 0x8031003D:
		return "FveEKeyfileInvalid"
	case 0x8031003E:
		return "FveEKeyfileNoVmk"
	case 0x8031003F:
		return "FveETpmDisabled"
	case 0x80310040:
		return "FveENotAllowedInSafeMode"
	case 0x80310041:
		return "FveETpmInvalidPcr"
	case 0x80310042:
		return "FveETpmNoVmk"
	case 0x80310043:
		return "FveEPinInvalid"
	case 0x80310044:
		return "FveEAuthInvalidApplication"
	case 0x80310045:
		return "FveEAuthInvalidConfig"
	case 0x80310046:
		return "FveEFipsDisableProtectionNotAllowed"
	case 0x80310047:
		return "FveEFsNotExtended"
	case 0x80310048:
		return "FveEFirmwareTypeNotSupported"
	case 0x80310049:
		return "FveENoLicense"
	case 0x8031004A:
		return "FveENotOnStack"
	case 0x8031004B:
		return "FveEFsMounted"
	case 0x8031004C:
		return "FveETokenNotImpersonated"
	case 0x8031004D:
		return "FveEDryRunFailed"
	case 0x8031004E:
		return "FveERebootRequired"
	case 0x8031004F:
		return "FveEDebuggerEnabled"
	case 0x80310050:
		return "FveERawAccess"
	case 0x80310051:
		return "FveERawBlocked"
	case 0x80310052:
		return "FveEBcdApplicationsPathIncorrect"
	case 0x80310053:
		return "FveENotAllowedInVersion"
	case 0x80310054:
		return "FveENoAutounlockMasterKey"
	case 0x80310055:
		return "FveEMorFailed"
	case 0x80310056:
		return "FveEHiddenVolume"
	case 0x80310057:
		return "FveETransientState"
	case 0x80310058:
		return "FveEPubkeyNotAllowed"
	case 0x80310059:
		return "FveEVolumeHandleOpen"
	case 0x8031005A:
		return "FveENoFeatureLicense"
	case 0x8031005B:
		return "FveEInvalidStartupOptions"
	case 0x8031005C:
		return "FveEPolicyRecoveryPasswordNotAllowed"
	case 0x8031005D:
		return "FveEPolicyRecoveryPasswordRequired"
	case 0x8031005E:
		return "FveEPolicyRecoveryKeyNotAllowed"
	case 0x8031005F:
		return "FveEPolicyRecoveryKeyRequired"
	case 0x80310060:
		return "FveEPolicyStartupPinNotAllowed"
	case 0x80310061:
		return "FveEPolicyStartupPinRequired"
	case 0x80310062:
		return "FveEPolicyStartupKeyNotAllowed"
	case 0x80310063:
		return "FveEPolicyStartupKeyRequired"
	case 0x80310064:
		return "FveEPolicyStartupPinKeyNotAllowed"
	case 0x80310065:
		return "FveEPolicyStartupPinKeyRequired"
	case 0x80310066:
		return "FveEPolicyStartupTpmNotAllowed"
	case 0x80310067:
		return "FveEPolicyStartupTpmRequired"
	case 0x80310068:
		return "FveEPolicyInvalidPinLength"
	case 0x80310069:
		return "FveEKeyProtectorNotSupported"
	case 0x8031006A:
		return "FveEPolicyPassphraseNotAllowed"
	case 0x8031006B:
		return "FveEPolicyPassphraseRequired"
	case 0x8031006C:
		return "FveEFipsPreventsPassphrase"
	case 0x8031006D:
		return "FveEOsVolumePassphraseNotAllowed"
	case 0x8031006E:
		return "FveEInvalidBitlockerOid"
	case 0x8031006F:
		return "FveEVolumeTooSmall"
	case 0x80310070:
		return "FveEDvNotSupportedOnFs"
	case 0x80310071:
		return "FveEDvNotAllowedByGp"
	case 0x80310072:
		return "FveEPolicyUserCertificateNotAllowed"
	case 0x80310073:
		return "FveEPolicyUserCertificateRequired"
	case 0x80310074:
		return "FveEPolicyUserCertMustBeHw"
	case 0x80310075:
		return "FveEPolicyUserConfigureFdvAutounlockNotAllowed"
	case 0x80310076:
		return "FveEPolicyUserConfigureRdvAutounlockNotAllowed"
	case 0x80310077:
		return "FveEPolicyUserConfigureRdvNotAllowed"
	case 0x80310078:
		return "FveEPolicyUserEnableRdvNotAllowed"
	case 0x80310079:
		return "FveEPolicyUserDisableRdvNotAllowed"
	case 0x80310080:
		return "FveEPolicyInvalidPassphraseLength"
	case 0x80310081:
		return "FveEPolicyPassphraseTooSimple"
	case 0x80310082:
		return "FveERecoveryPartition"
	case 0x80310083:
		return "FveEPolicyConflictFdvRkOffAukOn"
	case 0x80310084:
		return "FveEPolicyConflictRdvRkOffAukOn"
	case 0x80310085:
		return "FveENonBitlockerOid"
	case 0x80310086:
		return "FveEPolicyProhibitsSelfsigned"
	case 0x80310087:
		return "FveEPolicyConflictRoAndStartupKeyRequired"
	case 0x80310088:
		return "FveEConvRecoveryFailed"
	case 0x80310089:
		return "FveEVirtualizedSpaceTooBig"
	case 0x80310090:
		return "FveEPolicyConflictOsvRpOffAdbOn"
	case 0x80310091:
		return "FveEPolicyConflictFdvRpOffAdbOn"
	case 0x80310092:
		return "FveEPolicyConflictRdvRpOffAdbOn"
	case 0x80310093:
		return "FveENonBitlockerKu"
	case 0x80310094:
		return "FveEPrivatekeyAuthFailed"
	case 0x80310095:
		return "FveERemovalOfDraFailed"
	case 0x80310096:
		return "FveEOperationNotSupportedOnVistaVolume"
	case 0x80310097:
		return "FveECantLockAutounlockEnabledVolume"
	case 0x80310098:
		return "FveEFipsHashKdfNotAllowed"
	case 0x80310099:
		return "FveEEnhPinInvalid"
	case 0x8031009A:
		return "FveEInvalidPinChars"
	case 0x8031009B:
		return "FveEInvalidDatumType"
	case 0x80320001:
		return "FwpECalloutNotFound"
	case 0x80320002:
		return "FwpEConditionNotFound"
	case 0x80320003:
		return "FwpEFilterNotFound"
	case 0x80320004:
		return "FwpELayerNotFound"
	case 0x80320005:
		return "FwpEProviderNotFound"
	case 0x80320006:
		return "FwpEProviderContextNotFound"
	case 0x80320007:
		return "FwpESublayerNotFound"
	case 0x80320008:
		return "FwpENotFound"
	case 0x80320009:
		return "FwpEAlreadyExists"
	case 0x8032000A:
		return "FwpEInUse"
	case 0x8032000B:
		return "FwpEDynamicSessionInProgress"
	case 0x8032000C:
		return "FwpEWrongSession"
	case 0x8032000D:
		return "FwpENoTxnInProgress"
	case 0x8032000E:
		return "FwpETxnInProgress"
	case 0x8032000F:
		return "FwpETxnAborted"
	case 0x80320010:
		return "FwpESessionAborted"
	case 0x80320011:
		return "FwpEIncompatibleTxn"
	case 0x80320012:
		return "FwpETimeout"
	case 0x80320013:
		return "FwpENetEventsDisabled"
	case 0x80320014:
		return "FwpEIncompatibleLayer"
	case 0x80320015:
		return "FwpEKmClientsOnly"
	case 0x80320016:
		return "FwpELifetimeMismatch"
	case 0x80320017:
		return "FwpEBuiltinObject"
	case 0x80320018:
		return "FwpETooManyCallouts"
	case 0x80320019:
		return "FwpENotificationDropped"
	case 0x8032001A:
		return "FwpETrafficMismatch"
	case 0x8032001B:
		return "FwpEIncompatibleSaState"
	case 0x8032001C:
		return "FwpENullPointer"
	case 0x8032001D:
		return "FwpEInvalidEnumerator"
	case 0x8032001E:
		return "FwpEInvalidFlags"
	case 0x8032001F:
		return "FwpEInvalidNetMask"
	case 0x80320020:
		return "FwpEInvalidRange"
	case 0x80320021:
		return "FwpEInvalidInterval"
	case 0x80320022:
		return "FwpEZeroLengthArray"
	case 0x80320023:
		return "FwpENullDisplayName"
	case 0x80320024:
		return "FwpEInvalidActionType"
	case 0x80320025:
		return "FwpEInvalidWeight"
	case 0x80320026:
		return "FwpEMatchTypeMismatch"
	case 0x80320027:
		return "FwpETypeMismatch"
	case 0x80320028:
		return "FwpEOutOfBounds"
	case 0x80320029:
		return "FwpEReserved"
	case 0x8032002A:
		return "FwpEDuplicateCondition"
	case 0x8032002B:
		return "FwpEDuplicateKeymod"
	case 0x8032002C:
		return "FwpEActionIncompatibleWithLayer"
	case 0x8032002D:
		return "FwpEActionIncompatibleWithSublayer"
	case 0x8032002E:
		return "FwpEContextIncompatibleWithLayer"
	case 0x8032002F:
		return "FwpEContextIncompatibleWithCallout"
	case 0x80320030:
		return "FwpEIncompatibleAuthMethod"
	case 0x80320031:
		return "FwpEIncompatibleDhGroup"
	case 0x80320032:
		return "FwpEEmNotSupported"
	case 0x80320033:
		return "FwpENeverMatch"
	case 0x80320034:
		return "FwpEProviderContextMismatch"
	case 0x80320035:
		return "FwpEInvalidParameter"
	case 0x80320036:
		return "FwpETooManySublayers"
	case 0x80320037:
		return "FwpECalloutNotificationFailed"
	case 0x80320038:
		return "FwpEInvalidAuthTransform"
	case 0x80320039:
		return "FwpEInvalidCipherTransform"
	case 0x80320104:
		return "FwpEDropNoicmp"
	case 0x8032003A:
		return "FwpEIncompatibleCipherTransform"
	case 0x8032003B:
		return "FwpEInvalidTransformCombination"
	case 0x8032003C:
		return "FwpEDuplicateAuthMethod"
	case 0x003D0000:
		return "WsSAsync"
	case 0x003D0001:
		return "WsSEnd"
	case 0x803D0000:
		return "WsEInvalidFormat"
	case 0x803D0001:
		return "WsEObjectFaulted"
	case 0x803D0002:
		return "WsENumericOverflow"
	case 0x803D0003:
		return "WsEInvalidOperation"
	case 0x803D0004:
		return "WsEOperationAborted"
	case 0x803D0005:
		return "WsEEndpointAccessDenied"
	case 0x803D0006:
		return "WsEOperationTimedOut"
	case 0x803D0007:
		return "WsEOperationAbandoned"
	case 0x803D0008:
		return "WsEQuotaExceeded"
	case 0x803D0009:
		return "WsENoTranslationAvailable"
	case 0x803D000A:
		return "WsESecurityVerificationFailure"
	case 0x803D000B:
		return "WsEAddressInUse"
	case 0x803D000C:
		return "WsEAddressNotAvailable"
	case 0x803D000D:
		return "WsEEndpointNotFound"
	case 0x803D000E:
		return "WsEEndpointNotAvailable"
	case 0x803D000F:
		return "WsEEndpointFailure"
	case 0x803D0010:
		return "WsEEndpointUnreachable"
	case 0x803D0011:
		return "WsEEndpointActionNotSupported"
	case 0x803D0012:
		return "WsEEndpointTooBusy"
	case 0x803D0013:
		return "WsEEndpointFaultReceived"
	case 0x803D0014:
		return "WsEEndpointDisconnected"
	case 0x803D0015:
		return "WsEProxyFailure"
	case 0x803D0016:
		return "WsEProxyAccessDenied"
	case 0x803D0017:
		return "WsENotSupported"
	case 0x803D0018:
		return "WsEProxyRequiresBasicAuth"
	case 0x803D0019:
		return "WsEProxyRequiresDigestAuth"
	case 0x803D001A:
		return "WsEProxyRequiresNtlmAuth"
	case 0x803D001B:
		return "WsEProxyRequiresNegotiateAuth"
	case 0x803D001C:
		return "WsEServerRequiresBasicAuth"
	case 0x803D001D:
		return "WsEServerRequiresDigestAuth"
	case 0x803D001E:
		return "WsEServerRequiresNtlmAuth"
	case 0x803D001F:
		return "WsEServerRequiresNegotiateAuth"
	case 0x803D0020:
		return "WsEInvalidEndpointUrl"
	case 0x803D0021:
		return "WsEOther"
	case 0x803D0022:
		return "WsESecurityTokenExpired"
	case 0x803D0023:
		return "WsESecuritySystemFailure"
	case 0x80340002:
		return "ErrorNdisInterfaceClosing"
	case 0x80340004:
		return "ErrorNdisBadVersion"
	case 0x80340005:
		return "ErrorNdisBadCharacteristics"
	case 0x80340006:
		return "ErrorNdisAdapterNotFound"
	case 0x80340007:
		return "ErrorNdisOpenFailed"
	case 0x80340008:
		return "ErrorNdisDeviceFailed"
	case 0x80340009:
		return "ErrorNdisMulticastFull"
	case 0x8034000A:
		return "ErrorNdisMulticastExists"
	case 0x8034000B:
		return "ErrorNdisMulticastNotFound"
	case 0x8034000C:
		return "ErrorNdisRequestAborted"
	case 0x8034000D:
		return "ErrorNdisResetInProgress"
	case 0x803400BB:
		return "ErrorNdisNotSupported"
	case 0x8034000F:
		return "ErrorNdisInvalidPacket"
	case 0x80340011:
		return "ErrorNdisAdapterNotReady"
	case 0x80340014:
		return "ErrorNdisInvalidLength"
	case 0x80340015:
		return "ErrorNdisInvalidData"
	case 0x80340016:
		return "ErrorNdisBufferTooShort"
	case 0x80340017:
		return "ErrorNdisInvalidOid"
	case 0x80340018:
		return "ErrorNdisAdapterRemoved"
	case 0x80340019:
		return "ErrorNdisUnsupportedMedia"
	case 0x8034001A:
		return "ErrorNdisGroupAddressInUse"
	case 0x8034001B:
		return "ErrorNdisFileNotFound"
	case 0x8034001C:
		return "ErrorNdisErrorReadingFile"
	case 0x8034001D:
		return "ErrorNdisAlreadyMapped"
	case 0x8034001E:
		return "ErrorNdisResourceConflict"
	case 0x8034001F:
		return "ErrorNdisMediaDisconnected"
	case 0x80340022:
		return "ErrorNdisInvalidAddress"
	case 0x80340010:
		return "ErrorNdisInvalidDeviceRequest"
	case 0x8034002A:
		return "ErrorNdisPaused"
	case 0x8034002B:
		return "ErrorNdisInterfaceNotFound"
	case 0x8034002C:
		return "ErrorNdisUnsupportedRevision"
	case 0x8034002D:
		return "ErrorNdisInvalidPort"
	case 0x8034002E:
		return "ErrorNdisInvalidPortState"
	case 0x8034002F:
		return "ErrorNdisLowPowerState"
	case 0x80342000:
		return "ErrorNdisDot11AutoConfigEnabled"
	case 0x80342001:
		return "ErrorNdisDot11MediaInUse"
	case 0x80342002:
		return "ErrorNdisDot11PowerStateInvalid"
	case 0x80342003:
		return "ErrorNdisPmWolPatternListFull"
	case 0x80342004:
		return "ErrorNdisPmProtocolOffloadListFull"
	case 0x00340001:
		return "ErrorNdisIndicationRequired"
	case 0xC034100F:
		return "ErrorNdisOffloadPolicy"
	case 0xC0341012:
		return "ErrorNdisOffloadConnectionRejected"
	case 0xC0341013:
		return "ErrorNdisOffloadPathRejected"
	case 0xC0350002:
		return "ErrorHvInvalidHypercallCode"
	case 0xC0350003:
		return "ErrorHvInvalidHypercallInput"
	case 0xC0350004:
		return "ErrorHvInvalidAlignment"
	case 0xC0350005:
		return "ErrorHvInvalidParameter"
	case 0xC0350006:
		return "ErrorHvAccessDenied"
	case 0xC0350007:
		return "ErrorHvInvalidPartitionState"
	case 0xC0350008:
		return "ErrorHvOperationDenied"
	case 0xC0350009:
		return "ErrorHvUnknownProperty"
	case 0xC035000A:
		return "ErrorHvPropertyValueOutOfRange"
	case 0xC035000B:
		return "ErrorHvInsufficientMemory"
	case 0xC035000C:
		return "ErrorHvPartitionTooDeep"
	case 0xC035000D:
		return "ErrorHvInvalidPartitionId"
	case 0xC035000E:
		return "ErrorHvInvalidVpIndex"
	case 0xC0350011:
		return "ErrorHvInvalidPortId"
	case 0xC0350012:
		return "ErrorHvInvalidConnectionId"
	case 0xC0350013:
		return "ErrorHvInsufficientBuffers"
	case 0xC0350014:
		return "ErrorHvNotAcknowledged"
	case 0xC0350016:
		return "ErrorHvAcknowledged"
	case 0xC0350017:
		return "ErrorHvInvalidSaveRestoreState"
	case 0xC0350018:
		return "ErrorHvInvalidSynicState"
	case 0xC0350019:
		return "ErrorHvObjectInUse"
	case 0xC035001A:
		return "ErrorHvInvalidProximityDomainInfo"
	case 0xC035001B:
		return "ErrorHvNoData"
	case 0xC035001C:
		return "ErrorHvInactive"
	case 0xC035001D:
		return "ErrorHvNoResources"
	case 0xC035001E:
		return "ErrorHvFeatureUnavailable"
	case 0xC0351000:
		return "ErrorHvNotPresent"
	case 0xC0370001:
		return "ErrorVidDuplicateHandler"
	case 0xC0370002:
		return "ErrorVidTooManyHandlers"
	case 0xC0370003:
		return "ErrorVidQueueFull"
	case 0xC0370004:
		return "ErrorVidHandlerNotPresent"
	case 0xC0370005:
		return "ErrorVidInvalidObjectName"
	case 0xC0370006:
		return "ErrorVidPartitionNameTooLong"
	case 0xC0370007:
		return "ErrorVidMessageQueueNameTooLong"
	case 0xC0370008:
		return "ErrorVidPartitionAlreadyExists"
	case 0xC0370009:
		return "ErrorVidPartitionDoesNotExist"
	case 0xC037000A:
		return "ErrorVidPartitionNameNotFound"
	case 0xC037000B:
		return "ErrorVidMessageQueueAlreadyExists"
	case 0xC037000C:
		return "ErrorVidExceededMbpEntryMapLimit"
	case 0xC037000D:
		return "ErrorVidMbStillReferenced"
	case 0xC037000E:
		return "ErrorVidChildGpaPageSetCorrupted"
	case 0xC037000F:
		return "ErrorVidInvalidNumaSettings"
	case 0xC0370010:
		return "ErrorVidInvalidNumaNodeIndex"
	case 0xC0370011:
		return "ErrorVidNotificationQueueAlreadyAssociated"
	case 0xC0370012:
		return "ErrorVidInvalidMemoryBlockHandle"
	case 0xC0370013:
		return "ErrorVidPageRangeOverflow"
	case 0xC0370014:
		return "ErrorVidInvalidMessageQueueHandle"
	case 0xC0370015:
		return "ErrorVidInvalidGpaRangeHandle"
	case 0xC0370016:
		return "ErrorVidNoMemoryBlockNotificationQueue"
	case 0xC0370017:
		return "ErrorVidMemoryBlockLockCountExceeded"
	case 0xC0370018:
		return "ErrorVidInvalidPpmHandle"
	case 0xC0370019:
		return "ErrorVidMbpsAreLocked"
	case 0xC037001A:
		return "ErrorVidMessageQueueClosed"
	case 0xC037001B:
		return "ErrorVidVirtualProcessorLimitExceeded"
	case 0xC037001C:
		return "ErrorVidStopPending"
	case 0xC037001D:
		return "ErrorVidInvalidProcessorState"
	case 0xC037001E:
		return "ErrorVidExceededKmContextCountLimit"
	case 0xC037001F:
		return "ErrorVidKmInterfaceAlreadyInitialized"
	case 0xC0370020:
		return "ErrorVidMbPropertyAlreadySetReset"
	case 0xC0370021:
		return "ErrorVidMmioRangeDestroyed"
	case 0xC0370022:
		return "ErrorVidInvalidChildGpaPageSet"
	case 0xC0370023:
		return "ErrorVidReservePageSetIsBeingUsed"
	case 0xC0370024:
		return "ErrorVidReservePageSetTooSmall"
	case 0xC0370025:
		return "ErrorVidMbpAlreadyLockedUsingReservedPage"
	case 0xC0370026:
		return "ErrorVidMbpCountExceededLimit"
	case 0xC0370027:
		return "ErrorVidSavedStateCorrupt"
	case 0xC0370028:
		return "ErrorVidSavedStateUnrecognizedItem"
	case 0xC0370029:
		return "ErrorVidSavedStateIncompatible"
	case 0x80370001:
		return "ErrorVidRemoteNodeParentGpaPagesUsed"
	case 0x80380001:
		return "ErrorVolmgrIncompleteRegeneration"
	case 0x80380002:
		return "ErrorVolmgrIncompleteDiskMigration"
	case 0xC0380001:
		return "ErrorVolmgrDatabaseFull"
	case 0xC0380002:
		return "ErrorVolmgrDiskConfigurationCorrupted"
	case 0xC0380003:
		return "ErrorVolmgrDiskConfigurationNotInSync"
	case 0xC0380004:
		return "ErrorVolmgrPackConfigUpdateFailed"
	case 0xC0380005:
		return "ErrorVolmgrDiskContainsNonSimpleVolume"
	case 0xC0380006:
		return "ErrorVolmgrDiskDuplicate"
	case 0xC0380007:
		return "ErrorVolmgrDiskDynamic"
	case 0xC0380008:
		return "ErrorVolmgrDiskIdInvalid"
	case 0xC0380009:
		return "ErrorVolmgrDiskInvalid"
	case 0xC038000A:
		return "ErrorVolmgrDiskLastVoter"
	case 0xC038000B:
		return "ErrorVolmgrDiskLayoutInvalid"
	case 0xC038000C:
		return "ErrorVolmgrDiskLayoutNonBasicBetweenBasicPartitions"
	case 0xC038000D:
		return "ErrorVolmgrDiskLayoutNotCylinderAligned"
	case 0xC038000E:
		return "ErrorVolmgrDiskLayoutPartitionsTooSmall"
	case 0xC038000F:
		return "ErrorVolmgrDiskLayoutPrimaryBetweenLogicalPartitions"
	case 0xC0380010:
		return "ErrorVolmgrDiskLayoutTooManyPartitions"
	case 0xC0380011:
		return "ErrorVolmgrDiskMissing"
	case 0xC0380012:
		return "ErrorVolmgrDiskNotEmpty"
	case 0xC0380013:
		return "ErrorVolmgrDiskNotEnoughSpace"
	case 0xC0380014:
		return "ErrorVolmgrDiskRevectoringFailed"
	case 0xC0380015:
		return "ErrorVolmgrDiskSectorSizeInvalid"
	case 0xC0380016:
		return "ErrorVolmgrDiskSetNotContained"
	case 0xC0380017:
		return "ErrorVolmgrDiskUsedByMultipleMembers"
	case 0xC0380018:
		return "ErrorVolmgrDiskUsedByMultiplePlexes"
	case 0xC0380019:
		return "ErrorVolmgrDynamicDiskNotSupported"
	case 0xC038001A:
		return "ErrorVolmgrExtentAlreadyUsed"
	case 0xC038001B:
		return "ErrorVolmgrExtentNotContiguous"
	case 0xC038001C:
		return "ErrorVolmgrExtentNotInPublicRegion"
	case 0xC038001D:
		return "ErrorVolmgrExtentNotSectorAligned"
	case 0xC038001E:
		return "ErrorVolmgrExtentOverlapsEbrPartition"
	case 0xC038001F:
		return "ErrorVolmgrExtentVolumeLengthsDoNotMatch"
	case 0xC0380020:
		return "ErrorVolmgrFaultTolerantNotSupported"
	case 0xC0380021:
		return "ErrorVolmgrInterleaveLengthInvalid"
	case 0xC0380022:
		return "ErrorVolmgrMaximumRegisteredUsers"
	case 0xC0380023:
		return "ErrorVolmgrMemberInSync"
	case 0xC0380024:
		return "ErrorVolmgrMemberIndexDuplicate"
	case 0xC0380025:
		return "ErrorVolmgrMemberIndexInvalid"
	case 0xC0380026:
		return "ErrorVolmgrMemberMissing"
	case 0xC0380027:
		return "ErrorVolmgrMemberNotDetached"
	case 0xC0380028:
		return "ErrorVolmgrMemberRegenerating"
	case 0xC0380029:
		return "ErrorVolmgrAllDisksFailed"
	case 0xC038002A:
		return "ErrorVolmgrNoRegisteredUsers"
	case 0xC038002B:
		return "ErrorVolmgrNoSuchUser"
	case 0xC038002C:
		return "ErrorVolmgrNotificationReset"
	case 0xC038002D:
		return "ErrorVolmgrNumberOfMembersInvalid"
	case 0xC038002E:
		return "ErrorVolmgrNumberOfPlexesInvalid"
	case 0xC038002F:
		return "ErrorVolmgrPackDuplicate"
	case 0xC0380030:
		return "ErrorVolmgrPackIdInvalid"
	case 0xC0380031:
		return "ErrorVolmgrPackInvalid"
	case 0xC0380032:
		return "ErrorVolmgrPackNameInvalid"
	case 0xC0380033:
		return "ErrorVolmgrPackOffline"
	case 0xC0380034:
		return "ErrorVolmgrPackHasQuorum"
	case 0xC0380035:
		return "ErrorVolmgrPackWithoutQuorum"
	case 0xC0380036:
		return "ErrorVolmgrPartitionStyleInvalid"
	case 0xC0380037:
		return "ErrorVolmgrPartitionUpdateFailed"
	case 0xC0380038:
		return "ErrorVolmgrPlexInSync"
	case 0xC0380039:
		return "ErrorVolmgrPlexIndexDuplicate"
	case 0xC038003A:
		return "ErrorVolmgrPlexIndexInvalid"
	case 0xC038003B:
		return "ErrorVolmgrPlexLastActive"
	case 0xC038003C:
		return "ErrorVolmgrPlexMissing"
	case 0xC038003D:
		return "ErrorVolmgrPlexRegenerating"
	case 0xC038003E:
		return "ErrorVolmgrPlexTypeInvalid"
	case 0xC038003F:
		return "ErrorVolmgrPlexNotRaid5"
	case 0xC0380040:
		return "ErrorVolmgrPlexNotSimple"
	case 0xC0380041:
		return "ErrorVolmgrStructureSizeInvalid"
	case 0xC0380042:
		return "ErrorVolmgrTooManyNotificationRequests"
	case 0xC0380043:
		return "ErrorVolmgrTransactionInProgress"
	case 0xC0380044:
		return "ErrorVolmgrUnexpectedDiskLayoutChange"
	case 0xC0380045:
		return "ErrorVolmgrVolumeContainsMissingDisk"
	case 0xC0380046:
		return "ErrorVolmgrVolumeIdInvalid"
	case 0xC0380047:
		return "ErrorVolmgrVolumeLengthInvalid"
	case 0xC0380048:
		return "ErrorVolmgrVolumeLengthNotSectorSizeMultiple"
	case 0xC0380049:
		return "ErrorVolmgrVolumeNotMirrored"
	case 0xC038004A:
		return "ErrorVolmgrVolumeNotRetained"
	case 0xC038004B:
		return "ErrorVolmgrVolumeOffline"
	case 0xC038004C:
		return "ErrorVolmgrVolumeRetained"
	case 0xC038004D:
		return "ErrorVolmgrNumberOfExtentsInvalid"
	case 0xC038004E:
		return "ErrorVolmgrDifferentSectorSize"
	case 0xC038004F:
		return "ErrorVolmgrBadBootDisk"
	case 0xC0380050:
		return "ErrorVolmgrPackConfigOffline"
	case 0xC0380051:
		return "ErrorVolmgrPackConfigOnline"
	case 0xC0380052:
		return "ErrorVolmgrNotPrimaryPack"
	case 0xC0380053:
		return "ErrorVolmgrPackLogUpdateFailed"
	case 0xC0380054:
		return "ErrorVolmgrNumberOfDisksInPlexInvalid"
	case 0xC0380055:
		return "ErrorVolmgrNumberOfDisksInMemberInvalid"
	case 0xC0380056:
		return "ErrorVolmgrVolumeMirrored"
	case 0xC0380057:
		return "ErrorVolmgrPlexNotSimpleSpanned"
	case 0xC0380058:
		return "ErrorVolmgrNoValidLogCopies"
	case 0xC0380059:
		return "ErrorVolmgrPrimaryPackPresent"
	case 0xC038005A:
		return "ErrorVolmgrNumberOfDisksInvalid"
	case 0xC038005B:
		return "ErrorVolmgrMirrorNotSupported"
	case 0xC038005C:
		return "ErrorVolmgrRaid5NotSupported"
	case 0x80390001:
		return "ErrorBcdNotAllEntriesImported"
	case 0xC0390002:
		return "ErrorBcdTooManyElements"
	case 0x80390003:
		return "ErrorBcdNotAllEntriesSynchronized"
	case 0xC03A0001:
		return "ErrorVhdDriveFooterMissing"
	case 0xC03A0002:
		return "ErrorVhdDriveFooterChecksumMismatch"
	case 0xC03A0003:
		return "ErrorVhdDriveFooterCorrupt"
	case 0xC03A0004:
		return "ErrorVhdFormatUnknown"
	case 0xC03A0005:
		return "ErrorVhdFormatUnsupportedVersion"
	case 0xC03A0006:
		return "ErrorVhdSparseHeaderChecksumMismatch"
	case 0xC03A0007:
		return "ErrorVhdSparseHeaderUnsupportedVersion"
	case 0xC03A0008:
		return "ErrorVhdSparseHeaderCorrupt"
	case 0xC03A0009:
		return "ErrorVhdBlockAllocationFailure"
	case 0xC03A000A:
		return "ErrorVhdBlockAllocationTableCorrupt"
	case 0xC03A000B:
		return "ErrorVhdInvalidBlockSize"
	case 0xC03A000C:
		return "ErrorVhdBitmapMismatch"
	case 0xC03A000D:
		return "ErrorVhdParentVhdNotFound"
	case 0xC03A000E:
		return "ErrorVhdChildParentIdMismatch"
	case 0xC03A000F:
		return "ErrorVhdChildParentTimestampMismatch"
	case 0xC03A0010:
		return "ErrorVhdMetadataReadFailure"
	case 0xC03A0011:
		return "ErrorVhdMetadataWriteFailure"
	case 0xC03A0012:
		return "ErrorVhdInvalidSize"
	case 0xC03A0013:
		return "ErrorVhdInvalidFileSize"
	case 0xC03A0014:
		return "ErrorVirtdiskProviderNotFound"
	case 0xC03A0015:
		return "ErrorVirtdiskNotVirtualDisk"
	case 0xC03A0016:
		return "ErrorVhdParentVhdAccessDenied"
	case 0xC03A0017:
		return "ErrorVhdChildParentSizeMismatch"
	case 0xC03A0018:
		return "ErrorVhdDifferencingChainCycleDetected"
	case 0xC03A0019:
		return "ErrorVhdDifferencingChainErrorInParent"
	case 0xC03A001A:
		return "ErrorVirtualDiskLimitation"
	case 0xC03A001B:
		return "ErrorVhdInvalidType"
	case 0xC03A001C:
		return "ErrorVhdInvalidState"
	case 0xC03A001D:
		return "ErrorVirtdiskUnsupportedDiskSectorSize"
	case 0x803A0001:
		return "ErrorQueryStorageError"
	case 0x803C0100:
		return "SdiagECancelled"
	case 0x803C0101:
		return "SdiagEScript"
	case 0x803C0102:
		return "SdiagEPowershell"
	case 0x803C0103:
		return "SdiagEManagedhost"
	case 0x803C0104:
		return "SdiagENoverifier"
	case 0x003C0105:
		return "SdiagSCannotrun"
	case 0x803C0106:
		return "SdiagEDisabled"
	case 0x803C0107:
		return "SdiagETrust"
	case 0x803C0108:
		return "SdiagECannotrun"
	case 0x803C0109:
		return "SdiagEVersion"
	case 0x803C010A:
		return "SdiagEResource"
	case 0x803C010B:
		return "SdiagERootcause"
	case 0x80548201:
		return "EMbnContextNotActivated"
	case 0x80548202:
		return "EMbnBadSim"
	case 0x80548203:
		return "EMbnDataClassNotAvailable"
	case 0x80548204:
		return "EMbnInvalidAccessString"
	case 0x80548205:
		return "EMbnMaxActivatedContexts"
	case 0x80548206:
		return "EMbnPacketSvcDetached"
	case 0x80548207:
		return "EMbnProviderNotVisible"
	case 0x80548208:
		return "EMbnRadioPowerOff"
	case 0x80548209:
		return "EMbnServiceNotActivated"
	case 0x8054820A:
		return "EMbnSimNotInserted"
	case 0x8054820B:
		return "EMbnVoiceCallInProgress"
	case 0x8054820C:
		return "EMbnInvalidCache"
	case 0x8054820D:
		return "EMbnNotRegistered"
	case 0x8054820E:
		return "EMbnProvidersNotFound"
	case 0x8054820F:
		return "EMbnPinNotSupported"
	case 0x80548210:
		return "EMbnPinRequired"
	case 0x80548211:
		return "EMbnPinDisabled"
	case 0x80548212:
		return "EMbnFailure"
	case 0x80548218:
		return "EMbnInvalidProfile"
	case 0x80548219:
		return "EMbnDefaultProfileExist"
	case 0x80548220:
		return "EMbnSmsEncodingNotSupported"
	case 0x80548221:
		return "EMbnSmsFilterNotSupported"
	case 0x80548222:
		return "EMbnSmsInvalidMemoryIndex"
	case 0x80548223:
		return "EMbnSmsLangNotSupported"
	case 0x80548224:
		return "EMbnSmsMemoryFailure"
	case 0x80548225:
		return "EMbnSmsNetworkTimeout"
	case 0x80548226:
		return "EMbnSmsUnknownSmscAddress"
	case 0x80548227:
		return "EMbnSmsFormatNotSupported"
	case 0x80548228:
		return "EMbnSmsOperationNotAllowed"
	case 0x80548229:
		return "EMbnSmsMemoryFull"
	case 0x802A0001:
		return "UiECreateFailed"
	case 0x802A0002:
		return "UiEShutdownCalled"
	case 0x802A0003:
		return "UiEIllegalReentrancy"
	case 0x802A0004:
		return "UiEObjectSealed"
	case 0x802A0005:
		return "UiEValueNotSet"
	case 0x802A0006:
		return "UiEValueNotDetermined"
	case 0x802A0007:
		return "UiEInvalidOutput"
	case 0x802A0008:
		return "UiEBooleanExpected"
	case 0x802A0009:
		return "UiEDifferentOwner"
	case 0x802A000A:
		return "UiEAmbiguousMatch"
	case 0x802A000B:
		return "UiEFpOverflow"
	case 0x802A000C:
		return "UiEWrongThread"
	case 0x802A0101:
		return "UiEStoryboardActive"
	case 0x802A0102:
		return "UiEStoryboardNotPlaying"
	case 0x802A0103:
		return "UiEStartKeyframeAfterEnd"
	case 0x802A0104:
		return "UiEEndKeyframeNotDetermined"
	case 0x802A0105:
		return "UiELoopsOverlap"
	case 0x802A0106:
		return "UiETransitionAlreadyUsed"
	case 0x802A0107:
		return "UiETransitionNotInStoryboard"
	case 0x802A0108:
		return "UiETransitionEclipsed"
	case 0x802A0109:
		return "UiETimeBeforeLastUpdate"
	case 0x802A010A:
		return "UiETimerClientAlreadyConnected"
	default:
		return "unknown ErrorsKind " + fmt.Sprint(k)
	}
}

func (k ErrorsKind) Elements() []ErrorsKind {
	return []ErrorsKind{
		ErrorSuccess,
		ErrorInvalidFunction,
		ErrorFileNotFound,
		ErrorPathNotFound,
		ErrorTooManyOpenFiles,
		ErrorAccessDenied,
		ErrorInvalidHandle,
		ErrorArenaTrashed,
		ErrorNotEnoughMemory,
		ErrorInvalidBlock,
		ErrorBadEnvironment,
		ErrorBadFormat,
		ErrorInvalidAccess,
		ErrorInvalidData,
		ErrorOutofmemory,
		ErrorInvalidDrive,
		ErrorCurrentDirectory,
		ErrorNotSameDevice,
		ErrorNoMoreFiles,
		ErrorWriteProtect,
		ErrorBadUnit,
		ErrorNotReady,
		ErrorBadCommand,
		ErrorCrc,
		ErrorBadLength,
		ErrorSeek,
		ErrorNotDosDisk,
		ErrorSectorNotFound,
		ErrorOutOfPaper,
		ErrorWriteFault,
		ErrorReadFault,
		ErrorGenFailure,
		ErrorSharingViolation,
		ErrorLockViolation,
		ErrorWrongDisk,
		ErrorSharingBufferExceeded,
		ErrorHandleEof,
		ErrorHandleDiskFull,
		ErrorNotSupported,
		ErrorRemNotList,
		ErrorDupName,
		ErrorBadNetpath,
		ErrorNetworkBusy,
		ErrorDevNotExist,
		ErrorTooManyCmds,
		ErrorAdapHdwErr,
		ErrorBadNetResp,
		ErrorUnexpNetErr,
		ErrorBadRemAdap,
		ErrorPrintqFull,
		ErrorNoSpoolSpace,
		ErrorPrintCancelled,
		ErrorNetnameDeleted,
		ErrorNetworkAccessDenied,
		ErrorBadDevType,
		ErrorBadNetName,
		ErrorTooManyNames,
		ErrorTooManySess,
		ErrorSharingPaused,
		ErrorReqNotAccep,
		ErrorRedirPaused,
		ErrorFileExists,
		ErrorCannotMake,
		ErrorFailI24,
		ErrorOutOfStructures,
		ErrorAlreadyAssigned,
		ErrorInvalidPassword,
		ErrorInvalidParameter,
		ErrorNetWriteFault,
		ErrorNoProcSlots,
		ErrorTooManySemaphores,
		ErrorExclSemAlreadyOwned,
		ErrorSemIsSet,
		ErrorTooManySemRequests,
		ErrorInvalidAtInterruptTime,
		ErrorSemOwnerDied,
		ErrorSemUserLimit,
		ErrorDiskChange,
		ErrorDriveLocked,
		ErrorBrokenPipe,
		ErrorOpenFailed,
		ErrorBufferOverflow,
		ErrorDiskFull,
		ErrorNoMoreSearchHandles,
		ErrorInvalidTargetHandle,
		ErrorInvalidCategory,
		ErrorInvalidVerifySwitch,
		ErrorBadDriverLevel,
		ErrorCallNotImplemented,
		ErrorSemTimeout,
		ErrorInsufficientBuffer,
		ErrorInvalidName,
		ErrorInvalidLevel,
		ErrorNoVolumeLabel,
		ErrorModNotFound,
		ErrorProcNotFound,
		ErrorWaitNoChildren,
		ErrorChildNotComplete,
		ErrorDirectAccessHandle,
		ErrorNegativeSeek,
		ErrorSeekOnDevice,
		ErrorIsJoinTarget,
		ErrorIsJoined,
		ErrorIsSubsted,
		ErrorNotJoined,
		ErrorNotSubsted,
		ErrorJoinToJoin,
		ErrorSubstToSubst,
		ErrorJoinToSubst,
		ErrorSubstToJoin,
		ErrorBusyDrive,
		ErrorSameDrive,
		ErrorDirNotRoot,
		ErrorDirNotEmpty,
		ErrorIsSubstPath,
		ErrorIsJoinPath,
		ErrorPathBusy,
		ErrorIsSubstTarget,
		ErrorSystemTrace,
		ErrorInvalidEventCount,
		ErrorTooManyMuxwaiters,
		ErrorInvalidListFormat,
		ErrorLabelTooLong,
		ErrorTooManyTcbs,
		ErrorSignalRefused,
		ErrorDiscarded,
		ErrorNotLocked,
		ErrorBadThreadidAddr,
		ErrorBadArguments,
		ErrorBadPathname,
		ErrorSignalPending,
		ErrorMaxThrdsReached,
		ErrorLockFailed,
		ErrorBusy,
		ErrorCancelViolation,
		ErrorAtomicLocksNotSupported,
		ErrorInvalidSegmentNumber,
		ErrorInvalidOrdinal,
		ErrorAlreadyExists,
		ErrorInvalidFlagNumber,
		ErrorSemNotFound,
		ErrorInvalidStartingCodeseg,
		ErrorInvalidStackseg,
		ErrorInvalidModuletype,
		ErrorInvalidExeSignature,
		ErrorExeMarkedInvalid,
		ErrorBadExeFormat,
		ErrorIteratedDataExceeds64k,
		ErrorInvalidMinallocsize,
		ErrorDynlinkFromInvalidRing,
		ErrorIoplNotEnabled,
		ErrorInvalidSegdpl,
		ErrorAutodatasegExceeds64k,
		ErrorRing2SegMustBeMovable,
		ErrorRelocChainXeedsSeglim,
		ErrorInfloopInRelocChain,
		ErrorEnvvarNotFound,
		ErrorNoSignalSent,
		ErrorFilenameExcedRange,
		ErrorRing2StackInUse,
		ErrorMetaExpansionTooLong,
		ErrorInvalidSignalNumber,
		ErrorThread1Inactive,
		ErrorLocked,
		ErrorTooManyModules,
		ErrorNestingNotAllowed,
		ErrorExeMachineTypeMismatch,
		ErrorExeCannotModifySignedBinary,
		ErrorExeCannotModifyStrongSignedBinary,
		ErrorFileCheckedOut,
		ErrorCheckoutRequired,
		ErrorBadFileType,
		ErrorFileTooLarge,
		ErrorFormsAuthRequired,
		ErrorVirusInfected,
		ErrorVirusDeleted,
		ErrorPipeLocal,
		ErrorBadPipe,
		ErrorPipeBusy,
		ErrorNoData,
		ErrorPipeNotConnected,
		ErrorMoreData,
		ErrorVcDisconnected,
		ErrorInvalidEaName,
		ErrorEaListInconsistent,
		WaitTimeout,
		ErrorNoMoreItems,
		ErrorCannotCopy,
		ErrorDirectory,
		ErrorEasDidntFit,
		ErrorEaFileCorrupt,
		ErrorEaTableFull,
		ErrorInvalidEaHandle,
		ErrorEasNotSupported,
		ErrorNotOwner,
		ErrorTooManyPosts,
		ErrorPartialCopy,
		ErrorOplockNotGranted,
		ErrorInvalidOplockProtocol,
		ErrorDiskTooFragmented,
		ErrorDeletePending,
		ErrorIncompatibleWithGlobalShortNameRegistrySetting,
		ErrorShortNamesNotEnabledOnVolume,
		ErrorSecurityStreamIsInconsistent,
		ErrorInvalidLockRange,
		ErrorImageSubsystemNotPresent,
		ErrorNotificationGuidAlreadyDefined,
		ErrorMrMidNotFound,
		ErrorScopeNotFound,
		ErrorFailNoactionReboot,
		ErrorFailShutdown,
		ErrorFailRestart,
		ErrorMaxSessionsReached,
		ErrorThreadModeAlreadyBackground,
		ErrorThreadModeNotBackground,
		ErrorProcessModeAlreadyBackground,
		ErrorProcessModeNotBackground,
		ErrorInvalidAddress,
		ErrorUserProfileLoad,
		ErrorArithmeticOverflow,
		ErrorPipeConnected,
		ErrorPipeListening,
		ErrorVerifierStop,
		ErrorAbiosError,
		ErrorWx86Warning,
		ErrorWx86Error,
		ErrorTimerNotCanceled,
		ErrorUnwind,
		ErrorBadStack,
		ErrorInvalidUnwindTarget,
		ErrorInvalidPortAttributes,
		ErrorPortMessageTooLong,
		ErrorInvalidQuotaLower,
		ErrorDeviceAlreadyAttached,
		ErrorInstructionMisalignment,
		ErrorProfilingNotStarted,
		ErrorProfilingNotStopped,
		ErrorCouldNotInterpret,
		ErrorProfilingAtLimit,
		ErrorCantWait,
		ErrorCantTerminateSelf,
		ErrorUnexpectedMmCreateErr,
		ErrorUnexpectedMmMapError,
		ErrorUnexpectedMmExtendErr,
		ErrorBadFunctionTable,
		ErrorNoGuidTranslation,
		ErrorInvalidLdtSize,
		ErrorInvalidLdtOffset,
		ErrorInvalidLdtDescriptor,
		ErrorTooManyThreads,
		ErrorThreadNotInProcess,
		ErrorPagefileQuotaExceeded,
		ErrorLogonServerConflict,
		ErrorSynchronizationRequired,
		ErrorNetOpenFailed,
		ErrorIoPrivilegeFailed,
		ErrorControlCExit,
		ErrorMissingSystemfile,
		ErrorUnhandledException,
		ErrorAppInitFailure,
		ErrorPagefileCreateFailed,
		ErrorInvalidImageHash,
		ErrorNoPagefile,
		ErrorIllegalFloatContext,
		ErrorNoEventPair,
		ErrorDomainCtrlrConfigError,
		ErrorIllegalCharacter,
		ErrorUndefinedCharacter,
		ErrorFloppyVolume,
		ErrorBiosFailedToConnectInterrupt,
		ErrorBackupController,
		ErrorMutantLimitExceeded,
		ErrorFsDriverRequired,
		ErrorCannotLoadRegistryFile,
		ErrorDebugAttachFailed,
		ErrorSystemProcessTerminated,
		ErrorDataNotAccepted,
		ErrorVdmHardError,
		ErrorDriverCancelTimeout,
		ErrorReplyMessageMismatch,
		ErrorLostWritebehindData,
		ErrorClientServerParametersInvalid,
		ErrorNotTinyStream,
		ErrorStackOverflowRead,
		ErrorConvertToLarge,
		ErrorFoundOutOfScope,
		ErrorAllocateBucket,
		ErrorMarshallOverflow,
		ErrorInvalidVariant,
		ErrorBadCompressionBuffer,
		ErrorAuditFailed,
		ErrorTimerResolutionNotSet,
		ErrorInsufficientLogonInfo,
		ErrorBadDllEntrypoint,
		ErrorBadServiceEntrypoint,
		ErrorIpAddressConflict1,
		ErrorIpAddressConflict2,
		ErrorRegistryQuotaLimit,
		ErrorNoCallbackActive,
		ErrorPwdTooShort,
		ErrorPwdTooRecent,
		ErrorPwdHistoryConflict,
		ErrorUnsupportedCompression,
		ErrorInvalidHwProfile,
		ErrorInvalidPlugplayDevicePath,
		ErrorQuotaListInconsistent,
		ErrorEvaluationExpiration,
		ErrorIllegalDllRelocation,
		ErrorDllInitFailedLogoff,
		ErrorValidateContinue,
		ErrorNoMoreMatches,
		ErrorRangeListConflict,
		ErrorServerSidMismatch,
		ErrorCantEnableDenyOnly,
		ErrorFloatMultipleFaults,
		ErrorFloatMultipleTraps,
		ErrorNointerface,
		ErrorDriverFailedSleep,
		ErrorCorruptSystemFile,
		ErrorCommitmentMinimum,
		ErrorPnpRestartEnumeration,
		ErrorSystemImageBadSignature,
		ErrorPnpRebootRequired,
		ErrorInsufficientPower,
		ErrorMultipleFaultViolation,
		ErrorSystemShutdown,
		ErrorPortNotSet,
		ErrorDsVersionCheckFailure,
		ErrorRangeNotFound,
		ErrorNotSafeModeDriver,
		ErrorFailedDriverEntry,
		ErrorDeviceEnumerationError,
		ErrorMountPointNotResolved,
		ErrorInvalidDeviceObjectParameter,
		ErrorMcaOccured,
		ErrorDriverDatabaseError,
		ErrorSystemHiveTooLarge,
		ErrorDriverFailedPriorUnload,
		ErrorVolsnapPrepareHibernate,
		ErrorHibernationFailure,
		ErrorFileSystemLimitation,
		ErrorAssertionFailure,
		ErrorAcpiError,
		ErrorWowAssertion,
		ErrorPnpBadMpsTable,
		ErrorPnpTranslationFailed,
		ErrorPnpIrqTranslationFailed,
		ErrorPnpInvalidId,
		ErrorWakeSystemDebugger,
		ErrorHandlesClosed,
		ErrorExtraneousInformation,
		ErrorRxactCommitNecessary,
		ErrorMediaCheck,
		ErrorGuidSubstitutionMade,
		ErrorStoppedOnSymlink,
		ErrorLongjump,
		ErrorPlugplayQueryVetoed,
		ErrorUnwindConsolidate,
		ErrorRegistryHiveRecovered,
		ErrorDllMightBeInsecure,
		ErrorDllMightBeIncompatible,
		ErrorDbgExceptionNotHandled,
		ErrorDbgReplyLater,
		ErrorDbgUnableToProvideHandle,
		ErrorDbgTerminateThread,
		ErrorDbgTerminateProcess,
		ErrorDbgControlC,
		ErrorDbgPrintexceptionC,
		ErrorDbgRipexception,
		ErrorDbgControlBreak,
		ErrorDbgCommandException,
		ErrorObjectNameExists,
		ErrorThreadWasSuspended,
		ErrorImageNotAtBase,
		ErrorRxactStateCreated,
		ErrorSegmentNotification,
		ErrorBadCurrentDirectory,
		ErrorFtReadRecoveryFromBackup,
		ErrorFtWriteRecovery,
		ErrorImageMachineTypeMismatch,
		ErrorReceivePartial,
		ErrorReceiveExpedited,
		ErrorReceivePartialExpedited,
		ErrorEventDone,
		ErrorEventPending,
		ErrorCheckingFileSystem,
		ErrorFatalAppExit,
		ErrorPredefinedHandle,
		ErrorWasUnlocked,
		ErrorServiceNotification,
		ErrorWasLocked,
		ErrorLogHardError,
		ErrorAlreadyWin32,
		ErrorImageMachineTypeMismatchExe,
		ErrorNoYieldPerformed,
		ErrorTimerResumeIgnored,
		ErrorArbitrationUnhandled,
		ErrorCardbusNotSupported,
		ErrorMpProcessorMismatch,
		ErrorHibernated,
		ErrorResumeHibernation,
		ErrorFirmwareUpdated,
		ErrorDriversLeakingLockedPages,
		ErrorWakeSystem,
		ErrorWait1,
		ErrorWait2,
		ErrorWait3,
		ErrorWait63,
		ErrorAbandonedWait0,
		ErrorAbandonedWait63,
		ErrorUserApc,
		ErrorKernelApc,
		ErrorAlerted,
		ErrorElevationRequired,
		ErrorReparse,
		ErrorOplockBreakInProgress,
		ErrorVolumeMounted,
		ErrorRxactCommitted,
		ErrorNotifyCleanup,
		ErrorPrimaryTransportConnectFailed,
		ErrorPageFaultTransition,
		ErrorPageFaultDemandZero,
		ErrorPageFaultCopyOnWrite,
		ErrorPageFaultGuardPage,
		ErrorPageFaultPagingFile,
		ErrorCachePageLocked,
		ErrorCrashDump,
		ErrorBufferAllZeros,
		ErrorReparseObject,
		ErrorResourceRequirementsChanged,
		ErrorTranslationComplete,
		ErrorNothingToTerminate,
		ErrorProcessNotInJob,
		ErrorProcessInJob,
		ErrorVolsnapHibernateReady,
		ErrorFsfilterOpCompletedSuccessfully,
		ErrorInterruptVectorAlreadyConnected,
		ErrorInterruptStillConnected,
		ErrorWaitForOplock,
		ErrorDbgExceptionHandled,
		ErrorDbgContinue,
		ErrorCallbackPopStack,
		ErrorCompressionDisabled,
		ErrorCantfetchbackwards,
		ErrorCantscrollbackwards,
		ErrorRowsnotreleased,
		ErrorBadAccessorFlags,
		ErrorErrorsEncountered,
		ErrorNotCapable,
		ErrorRequestOutOfSequence,
		ErrorVersionParseError,
		ErrorBadstartposition,
		ErrorMemoryHardware,
		ErrorDiskRepairDisabled,
		ErrorInsufficientResourceForSpecifiedSharedSectionSize,
		ErrorSystemPowerstateTransition,
		ErrorSystemPowerstateComplexTransition,
		ErrorMcaException,
		ErrorAccessAuditByPolicy,
		ErrorAccessDisabledNoSaferUiByPolicy,
		ErrorAbandonHiberfile,
		ErrorLostWritebehindDataNetworkDisconnected,
		ErrorLostWritebehindDataNetworkServerError,
		ErrorLostWritebehindDataLocalDiskError,
		ErrorBadMcfgTable,
		ErrorOplockSwitchedToNewHandle,
		ErrorCannotGrantRequestedOplock,
		ErrorCannotBreakOplock,
		ErrorOplockHandleClosed,
		ErrorNoAceCondition,
		ErrorInvalidAceCondition,
		ErrorEaAccessDenied,
		ErrorOperationAborted,
		ErrorIoIncomplete,
		ErrorIoPending,
		ErrorNoaccess,
		ErrorSwaperror,
		ErrorStackOverflow,
		ErrorInvalidMessage,
		ErrorCanNotComplete,
		ErrorInvalidFlags,
		ErrorUnrecognizedVolume,
		ErrorFileInvalid,
		ErrorFullscreenMode,
		ErrorNoToken,
		ErrorBaddb,
		ErrorBadkey,
		ErrorCantopen,
		ErrorCantread,
		ErrorCantwrite,
		ErrorRegistryRecovered,
		ErrorRegistryCorrupt,
		ErrorRegistryIoFailed,
		ErrorNotRegistryFile,
		ErrorKeyDeleted,
		ErrorNoLogSpace,
		ErrorKeyHasChildren,
		ErrorChildMustBeVolatile,
		ErrorNotifyEnumDir,
		ErrorDependentServicesRunning,
		ErrorInvalidServiceControl,
		ErrorServiceRequestTimeout,
		ErrorServiceNoThread,
		ErrorServiceDatabaseLocked,
		ErrorServiceAlreadyRunning,
		ErrorInvalidServiceAccount,
		ErrorServiceDisabled,
		ErrorCircularDependency,
		ErrorServiceDoesNotExist,
		ErrorServiceCannotAcceptCtrl,
		ErrorServiceNotActive,
		ErrorFailedServiceControllerConnect,
		ErrorExceptionInService,
		ErrorDatabaseDoesNotExist,
		ErrorServiceSpecificError,
		ErrorProcessAborted,
		ErrorServiceDependencyFail,
		ErrorServiceLogonFailed,
		ErrorServiceStartHang,
		ErrorInvalidServiceLock,
		ErrorServiceMarkedForDelete,
		ErrorServiceExists,
		ErrorAlreadyRunningLkg,
		ErrorServiceDependencyDeleted,
		ErrorBootAlreadyAccepted,
		ErrorServiceNeverStarted,
		ErrorDuplicateServiceName,
		ErrorDifferentServiceAccount,
		ErrorCannotDetectDriverFailure,
		ErrorCannotDetectProcessAbort,
		ErrorNoRecoveryProgram,
		ErrorServiceNotInExe,
		ErrorNotSafebootService,
		ErrorEndOfMedia,
		ErrorFilemarkDetected,
		ErrorBeginningOfMedia,
		ErrorSetmarkDetected,
		ErrorNoDataDetected,
		ErrorPartitionFailure,
		ErrorInvalidBlockLength,
		ErrorDeviceNotPartitioned,
		ErrorUnableToLockMedia,
		ErrorUnableToUnloadMedia,
		ErrorMediaChanged,
		ErrorBusReset,
		ErrorNoMediaInDrive,
		ErrorNoUnicodeTranslation,
		ErrorDllInitFailed,
		ErrorShutdownInProgress,
		ErrorNoShutdownInProgress,
		ErrorIoDevice,
		ErrorSerialNoDevice,
		ErrorIrqBusy,
		ErrorMoreWrites,
		ErrorCounterTimeout,
		ErrorFloppyIdMarkNotFound,
		ErrorFloppyWrongCylinder,
		ErrorFloppyUnknownError,
		ErrorFloppyBadRegisters,
		ErrorDiskRecalibrateFailed,
		ErrorDiskOperationFailed,
		ErrorDiskResetFailed,
		ErrorEomOverflow,
		ErrorNotEnoughServerMemory,
		ErrorPossibleDeadlock,
		ErrorMappedAlignment,
		ErrorSetPowerStateVetoed,
		ErrorSetPowerStateFailed,
		ErrorTooManyLinks,
		ErrorOldWinVersion,
		ErrorAppWrongOs,
		ErrorSingleInstanceApp,
		ErrorRmodeApp,
		ErrorInvalidDll,
		ErrorNoAssociation,
		ErrorDdeFail,
		ErrorDllNotFound,
		ErrorNoMoreUserHandles,
		ErrorMessageSyncOnly,
		ErrorSourceElementEmpty,
		ErrorDestinationElementFull,
		ErrorIllegalElementAddress,
		ErrorMagazineNotPresent,
		ErrorDeviceReinitializationNeeded,
		ErrorDeviceRequiresCleaning,
		ErrorDeviceDoorOpen,
		ErrorDeviceNotConnected,
		ErrorNotFound,
		ErrorNoMatch,
		ErrorSetNotFound,
		ErrorPointNotFound,
		ErrorNoTrackingService,
		ErrorNoVolumeId,
		ErrorUnableToRemoveReplaced,
		ErrorUnableToMoveReplacement,
		ErrorUnableToMoveReplacement2,
		ErrorJournalDeleteInProgress,
		ErrorJournalNotActive,
		ErrorPotentialFileFound,
		ErrorJournalEntryDeleted,
		ErrorShutdownIsScheduled,
		ErrorShutdownUsersLoggedOn,
		ErrorBadDevice,
		ErrorConnectionUnavail,
		ErrorDeviceAlreadyRemembered,
		ErrorNoNetOrBadPath,
		ErrorBadProvider,
		ErrorCannotOpenProfile,
		ErrorBadProfile,
		ErrorNotContainer,
		ErrorExtendedError,
		ErrorInvalidGroupname,
		ErrorInvalidComputername,
		ErrorInvalidEventname,
		ErrorInvalidDomainname,
		ErrorInvalidServicename,
		ErrorInvalidNetname,
		ErrorInvalidSharename,
		ErrorInvalidPasswordname,
		ErrorInvalidMessagename,
		ErrorInvalidMessagedest,
		ErrorSessionCredentialConflict,
		ErrorRemoteSessionLimitExceeded,
		ErrorDupDomainname,
		ErrorNoNetwork,
		ErrorCancelled,
		ErrorUserMappedFile,
		ErrorConnectionRefused,
		ErrorGracefulDisconnect,
		ErrorAddressAlreadyAssociated,
		ErrorAddressNotAssociated,
		ErrorConnectionInvalid,
		ErrorConnectionActive,
		ErrorNetworkUnreachable,
		ErrorHostUnreachable,
		ErrorProtocolUnreachable,
		ErrorPortUnreachable,
		ErrorRequestAborted,
		ErrorConnectionAborted,
		ErrorRetry,
		ErrorConnectionCountLimit,
		ErrorLoginTimeRestriction,
		ErrorLoginWkstaRestriction,
		ErrorIncorrectAddress,
		ErrorAlreadyRegistered,
		ErrorServiceNotFound,
		ErrorNotAuthenticated,
		ErrorNotLoggedOn,
		ErrorContinue,
		ErrorAlreadyInitialized,
		ErrorNoMoreDevices,
		ErrorNoSuchSite,
		ErrorDomainControllerExists,
		ErrorOnlyIfConnected,
		ErrorOverrideNochanges,
		ErrorBadUserProfile,
		ErrorNotSupportedOnSbs,
		ErrorServerShutdownInProgress,
		ErrorHostDown,
		ErrorNonAccountSid,
		ErrorNonDomainSid,
		ErrorApphelpBlock,
		ErrorAccessDisabledByPolicy,
		ErrorRegNatConsumption,
		ErrorCscshareOffline,
		ErrorPkinitFailure,
		ErrorSmartcardSubsystemFailure,
		ErrorDowngradeDetected,
		ErrorMachineLocked,
		ErrorCallbackSuppliedInvalidData,
		ErrorSyncForegroundRefreshRequired,
		ErrorDriverBlocked,
		ErrorInvalidImportOfNonDll,
		ErrorAccessDisabledWebblade,
		ErrorAccessDisabledWebbladeTamper,
		ErrorRecoveryFailure,
		ErrorAlreadyFiber,
		ErrorAlreadyThread,
		ErrorStackBufferOverrun,
		ErrorParameterQuotaExceeded,
		ErrorDebuggerInactive,
		ErrorDelayLoadFailed,
		ErrorVdmDisallowed,
		ErrorUnidentifiedError,
		ErrorInvalidCruntimeParameter,
		ErrorBeyondVdl,
		ErrorIncompatibleServiceSidType,
		ErrorDriverProcessTerminated,
		ErrorImplementationLimit,
		ErrorProcessIsProtected,
		ErrorServiceNotifyClientLagging,
		ErrorDiskQuotaExceeded,
		ErrorContentBlocked,
		ErrorIncompatibleServicePrivilege,
		ErrorAppHang,
		ErrorInvalidLabel,
		ErrorNotAllAssigned,
		ErrorSomeNotMapped,
		ErrorNoQuotasForAccount,
		ErrorLocalUserSessionKey,
		ErrorNullLmPassword,
		ErrorUnknownRevision,
		ErrorRevisionMismatch,
		ErrorInvalidOwner,
		ErrorInvalidPrimaryGroup,
		ErrorNoImpersonationToken,
		ErrorCantDisableMandatory,
		ErrorNoLogonServers,
		ErrorNoSuchLogonSession,
		ErrorNoSuchPrivilege,
		ErrorPrivilegeNotHeld,
		ErrorInvalidAccountName,
		ErrorUserExists,
		ErrorNoSuchUser,
		ErrorGroupExists,
		ErrorNoSuchGroup,
		ErrorMemberInGroup,
		ErrorMemberNotInGroup,
		ErrorLastAdmin,
		ErrorWrongPassword,
		ErrorIllFormedPassword,
		ErrorPasswordRestriction,
		ErrorLogonFailure,
		ErrorAccountRestriction,
		ErrorInvalidLogonHours,
		ErrorInvalidWorkstation,
		ErrorPasswordExpired,
		ErrorAccountDisabled,
		ErrorNoneMapped,
		ErrorTooManyLuidsRequested,
		ErrorLuidsExhausted,
		ErrorInvalidSubAuthority,
		ErrorInvalidAcl,
		ErrorInvalidSid,
		ErrorInvalidSecurityDescr,
		ErrorBadInheritanceAcl,
		ErrorServerDisabled,
		ErrorServerNotDisabled,
		ErrorInvalidIdAuthority,
		ErrorAllottedSpaceExceeded,
		ErrorInvalidGroupAttributes,
		ErrorBadImpersonationLevel,
		ErrorCantOpenAnonymous,
		ErrorBadValidationClass,
		ErrorBadTokenType,
		ErrorNoSecurityOnObject,
		ErrorCantAccessDomainInfo,
		ErrorInvalidServerState,
		ErrorInvalidDomainState,
		ErrorInvalidDomainRole,
		ErrorNoSuchDomain,
		ErrorDomainExists,
		ErrorDomainLimitExceeded,
		ErrorInternalDbCorruption,
		ErrorInternalError,
		ErrorGenericNotMapped,
		ErrorBadDescriptorFormat,
		ErrorNotLogonProcess,
		ErrorLogonSessionExists,
		ErrorNoSuchPackage,
		ErrorBadLogonSessionState,
		ErrorLogonSessionCollision,
		ErrorInvalidLogonType,
		ErrorCannotImpersonate,
		ErrorRxactInvalidState,
		ErrorRxactCommitFailure,
		ErrorSpecialAccount,
		ErrorSpecialGroup,
		ErrorSpecialUser,
		ErrorMembersPrimaryGroup,
		ErrorTokenAlreadyInUse,
		ErrorNoSuchAlias,
		ErrorMemberNotInAlias,
		ErrorMemberInAlias,
		ErrorAliasExists,
		ErrorLogonNotGranted,
		ErrorTooManySecrets,
		ErrorSecretTooLong,
		ErrorInternalDbError,
		ErrorTooManyContextIds,
		ErrorLogonTypeNotGranted,
		ErrorNtCrossEncryptionRequired,
		ErrorNoSuchMember,
		ErrorInvalidMember,
		ErrorTooManySids,
		ErrorLmCrossEncryptionRequired,
		ErrorNoInheritance,
		ErrorFileCorrupt,
		ErrorDiskCorrupt,
		ErrorNoUserSessionKey,
		ErrorLicenseQuotaExceeded,
		ErrorWrongTargetName,
		ErrorMutualAuthFailed,
		ErrorTimeSkew,
		ErrorCurrentDomainNotAllowed,
		ErrorInvalidWindowHandle,
		ErrorInvalidMenuHandle,
		ErrorInvalidCursorHandle,
		ErrorInvalidAccelHandle,
		ErrorInvalidHookHandle,
		ErrorInvalidDwpHandle,
		ErrorTlwWithWschild,
		ErrorCannotFindWndClass,
		ErrorWindowOfOtherThread,
		ErrorHotkeyAlreadyRegistered,
		ErrorClassAlreadyExists,
		ErrorClassDoesNotExist,
		ErrorClassHasWindows,
		ErrorInvalidIndex,
		ErrorInvalidIconHandle,
		ErrorPrivateDialogIndex,
		ErrorListboxIdNotFound,
		ErrorNoWildcardCharacters,
		ErrorClipboardNotOpen,
		ErrorHotkeyNotRegistered,
		ErrorWindowNotDialog,
		ErrorControlIdNotFound,
		ErrorInvalidComboboxMessage,
		ErrorWindowNotCombobox,
		ErrorInvalidEditHeight,
		ErrorDcNotFound,
		ErrorInvalidHookFilter,
		ErrorInvalidFilterProc,
		ErrorHookNeedsHmod,
		ErrorGlobalOnlyHook,
		ErrorJournalHookSet,
		ErrorHookNotInstalled,
		ErrorInvalidLbMessage,
		ErrorSetcountOnBadLb,
		ErrorLbWithoutTabstops,
		ErrorDestroyObjectOfOtherThread,
		ErrorChildWindowMenu,
		ErrorNoSystemMenu,
		ErrorInvalidMsgboxStyle,
		ErrorInvalidSpiValue,
		ErrorScreenAlreadyLocked,
		ErrorHwndsHaveDiffParent,
		ErrorNotChildWindow,
		ErrorInvalidGwCommand,
		ErrorInvalidThreadId,
		ErrorNonMdichildWindow,
		ErrorPopupAlreadyActive,
		ErrorNoScrollbars,
		ErrorInvalidScrollbarRange,
		ErrorInvalidShowwinCommand,
		ErrorNoSystemResources,
		ErrorNonpagedSystemResources,
		ErrorPagedSystemResources,
		ErrorWorkingSetQuota,
		ErrorPagefileQuota,
		ErrorCommitmentLimit,
		ErrorMenuItemNotFound,
		ErrorInvalidKeyboardHandle,
		ErrorHookTypeNotAllowed,
		ErrorRequiresInteractiveWindowstation,
		ErrorTimeout,
		ErrorInvalidMonitorHandle,
		ErrorIncorrectSize,
		ErrorSymlinkClassDisabled,
		ErrorSymlinkNotSupported,
		ErrorXmlParseError,
		ErrorXmldsigError,
		ErrorRestartApplication,
		ErrorWrongCompartment,
		ErrorAuthipFailure,
		ErrorNoNvramResources,
		ErrorEventlogFileCorrupt,
		ErrorEventlogCantStart,
		ErrorLogFileFull,
		ErrorEventlogFileChanged,
		ErrorInvalidTaskName,
		ErrorInvalidTaskIndex,
		ErrorThreadAlreadyInTask,
		ErrorInstallServiceFailure,
		ErrorInstallUserexit,
		ErrorInstallFailure,
		ErrorInstallSuspend,
		ErrorUnknownProduct,
		ErrorUnknownFeature,
		ErrorUnknownComponent,
		ErrorUnknownProperty,
		ErrorInvalidHandleState,
		ErrorBadConfiguration,
		ErrorIndexAbsent,
		ErrorInstallSourceAbsent,
		ErrorInstallPackageVersion,
		ErrorProductUninstalled,
		ErrorBadQuerySyntax,
		ErrorInvalidField,
		ErrorDeviceRemoved,
		ErrorInstallAlreadyRunning,
		ErrorInstallPackageOpenFailed,
		ErrorInstallPackageInvalid,
		ErrorInstallUiFailure,
		ErrorInstallLogFailure,
		ErrorInstallLanguageUnsupported,
		ErrorInstallTransformFailure,
		ErrorInstallPackageRejected,
		ErrorFunctionNotCalled,
		ErrorFunctionFailed,
		ErrorInvalidTable,
		ErrorDatatypeMismatch,
		ErrorUnsupportedType,
		ErrorCreateFailed,
		ErrorInstallTempUnwritable,
		ErrorInstallPlatformUnsupported,
		ErrorInstallNotused,
		ErrorPatchPackageOpenFailed,
		ErrorPatchPackageInvalid,
		ErrorPatchPackageUnsupported,
		ErrorProductVersion,
		ErrorInvalidCommandLine,
		ErrorInstallRemoteDisallowed,
		ErrorSuccessRebootInitiated,
		ErrorPatchTargetNotFound,
		ErrorPatchPackageRejected,
		ErrorInstallTransformRejected,
		ErrorInstallRemoteProhibited,
		ErrorPatchRemovalUnsupported,
		ErrorUnknownPatch,
		ErrorPatchNoSequence,
		ErrorPatchRemovalDisallowed,
		ErrorInvalidPatchXml,
		ErrorPatchManagedAdvertisedProduct,
		ErrorInstallServiceSafeboot,
		ErrorFailFastException,
		RpcSInvalidStringBinding,
		RpcSWrongKindOfBinding,
		RpcSInvalidBinding,
		RpcSProtseqNotSupported,
		RpcSInvalidRpcProtseq,
		RpcSInvalidStringUuid,
		RpcSInvalidEndpointFormat,
		RpcSInvalidNetAddr,
		RpcSNoEndpointFound,
		RpcSInvalidTimeout,
		RpcSObjectNotFound,
		RpcSAlreadyRegistered,
		RpcSTypeAlreadyRegistered,
		RpcSAlreadyListening,
		RpcSNoProtseqsRegistered,
		RpcSNotListening,
		RpcSUnknownMgrType,
		RpcSUnknownIf,
		RpcSNoBindings,
		RpcSNoProtseqs,
		RpcSCantCreateEndpoint,
		RpcSOutOfResources,
		RpcSServerUnavailable,
		RpcSServerTooBusy,
		RpcSInvalidNetworkOptions,
		RpcSNoCallActive,
		RpcSCallFailed,
		RpcSCallFailedDne,
		RpcSProtocolError,
		RpcSProxyAccessDenied,
		RpcSUnsupportedTransSyn,
		RpcSUnsupportedType,
		RpcSInvalidTag,
		RpcSInvalidBound,
		RpcSNoEntryName,
		RpcSInvalidNameSyntax,
		RpcSUnsupportedNameSyntax,
		RpcSUuidNoAddress,
		RpcSDuplicateEndpoint,
		RpcSUnknownAuthnType,
		RpcSMaxCallsTooSmall,
		RpcSStringTooLong,
		RpcSProtseqNotFound,
		RpcSProcnumOutOfRange,
		RpcSBindingHasNoAuth,
		RpcSUnknownAuthnService,
		RpcSUnknownAuthnLevel,
		RpcSInvalidAuthIdentity,
		RpcSUnknownAuthzService,
		EptSInvalidEntry,
		EptSCantPerformOp,
		EptSNotRegistered,
		RpcSNothingToExport,
		RpcSIncompleteName,
		RpcSInvalidVersOption,
		RpcSNoMoreMembers,
		RpcSNotAllObjsUnexported,
		RpcSInterfaceNotFound,
		RpcSEntryAlreadyExists,
		RpcSEntryNotFound,
		RpcSNameServiceUnavailable,
		RpcSInvalidNafId,
		RpcSCannotSupport,
		RpcSNoContextAvailable,
		RpcSInternalError,
		RpcSZeroDivide,
		RpcSAddressError,
		RpcSFpDivZero,
		RpcSFpUnderflow,
		RpcSFpOverflow,
		RpcXNoMoreEntries,
		RpcXSsCharTransOpenFail,
		RpcXSsCharTransShortFile,
		RpcXSsInNullContext,
		RpcXSsContextDamaged,
		RpcXSsHandlesMismatch,
		RpcXSsCannotGetCallHandle,
		RpcXNullRefPointer,
		RpcXEnumValueOutOfRange,
		RpcXByteCountTooSmall,
		RpcXBadStubData,
		ErrorInvalidUserBuffer,
		ErrorUnrecognizedMedia,
		ErrorNoTrustLsaSecret,
		ErrorNoTrustSamAccount,
		ErrorTrustedDomainFailure,
		ErrorTrustedRelationshipFailure,
		ErrorTrustFailure,
		RpcSCallInProgress,
		ErrorNetlogonNotStarted,
		ErrorAccountExpired,
		ErrorRedirectorHasOpenHandles,
		ErrorPrinterDriverAlreadyInstalled,
		ErrorUnknownPort,
		ErrorUnknownPrinterDriver,
		ErrorUnknownPrintprocessor,
		ErrorInvalidSeparatorFile,
		ErrorInvalidPriority,
		ErrorInvalidPrinterName,
		ErrorPrinterAlreadyExists,
		ErrorInvalidPrinterCommand,
		ErrorInvalidDatatype,
		ErrorInvalidEnvironment,
		RpcSNoMoreBindings,
		ErrorNologonInterdomainTrustAccount,
		ErrorNologonWorkstationTrustAccount,
		ErrorNologonServerTrustAccount,
		ErrorDomainTrustInconsistent,
		ErrorServerHasOpenHandles,
		ErrorResourceDataNotFound,
		ErrorResourceTypeNotFound,
		ErrorResourceNameNotFound,
		ErrorResourceLangNotFound,
		ErrorNotEnoughQuota,
		RpcSNoInterfaces,
		RpcSCallCancelled,
		RpcSBindingIncomplete,
		RpcSCommFailure,
		RpcSUnsupportedAuthnLevel,
		RpcSNoPrincName,
		RpcSNotRpcError,
		RpcSUuidLocalOnly,
		RpcSSecPkgError,
		RpcSNotCancelled,
		RpcXInvalidEsAction,
		RpcXWrongEsVersion,
		RpcXWrongStubVersion,
		RpcXInvalidPipeObject,
		RpcXWrongPipeOrder,
		RpcXWrongPipeVersion,
		RpcSCookieAuthFailed,
		RpcSGroupMemberNotFound,
		EptSCantCreate,
		RpcSInvalidObject,
		ErrorInvalidTime,
		ErrorInvalidFormName,
		ErrorInvalidFormSize,
		ErrorAlreadyWaiting,
		ErrorPrinterDeleted,
		ErrorInvalidPrinterState,
		ErrorPasswordMustChange,
		ErrorDomainControllerNotFound,
		ErrorAccountLockedOut,
		OrInvalidOxid,
		OrInvalidOid,
		OrInvalidSet,
		RpcSSendIncomplete,
		RpcSInvalidAsyncHandle,
		RpcSInvalidAsyncCall,
		RpcXPipeClosed,
		RpcXPipeDisciplineError,
		RpcXPipeEmpty,
		ErrorNoSitename,
		ErrorCantAccessFile,
		ErrorCantResolveFilename,
		RpcSEntryTypeMismatch,
		RpcSNotAllObjsExported,
		RpcSInterfaceNotExported,
		RpcSProfileNotAdded,
		RpcSPrfEltNotAdded,
		RpcSPrfEltNotRemoved,
		RpcSGrpEltNotAdded,
		RpcSGrpEltNotRemoved,
		ErrorKmDriverBlocked,
		ErrorContextExpired,
		ErrorPerUserTrustQuotaExceeded,
		ErrorAllUserTrustQuotaExceeded,
		ErrorUserDeleteTrustQuotaExceeded,
		ErrorAuthenticationFirewallFailed,
		ErrorRemotePrintConnectionsBlocked,
		ErrorNtlmBlocked,
		ErrorInvalidPixelFormat,
		ErrorBadDriver,
		ErrorInvalidWindowStyle,
		ErrorMetafileNotSupported,
		ErrorTransformNotSupported,
		ErrorClippingNotSupported,
		ErrorInvalidCmm,
		ErrorInvalidProfile,
		ErrorTagNotFound,
		ErrorTagNotPresent,
		ErrorDuplicateTag,
		ErrorProfileNotAssociatedWithDevice,
		ErrorProfileNotFound,
		ErrorInvalidColorspace,
		ErrorIcmNotEnabled,
		ErrorDeletingIcmXform,
		ErrorInvalidTransform,
		ErrorColorspaceMismatch,
		ErrorInvalidColorindex,
		ErrorProfileDoesNotMatchDevice,
		ErrorConnectedOtherPassword,
		ErrorConnectedOtherPasswordDefault,
		ErrorBadUsername,
		ErrorNotConnected,
		ErrorOpenFiles,
		ErrorActiveConnections,
		ErrorDeviceInUse,
		ErrorUnknownPrintMonitor,
		ErrorPrinterDriverInUse,
		ErrorSpoolFileNotFound,
		ErrorSplNoStartdoc,
		ErrorSplNoAddjob,
		ErrorPrintProcessorAlreadyInstalled,
		ErrorPrintMonitorAlreadyInstalled,
		ErrorInvalidPrintMonitor,
		ErrorPrintMonitorInUse,
		ErrorPrinterHasJobsQueued,
		ErrorSuccessRebootRequired,
		ErrorSuccessRestartRequired,
		ErrorPrinterNotFound,
		ErrorPrinterDriverWarned,
		ErrorPrinterDriverBlocked,
		ErrorPrinterDriverPackageInUse,
		ErrorCoreDriverPackageNotFound,
		ErrorFailRebootRequired,
		ErrorFailRebootInitiated,
		ErrorPrinterDriverDownloadNeeded,
		ErrorPrintJobRestartRequired,
		ErrorIoReissueAsCached,
		ErrorWinsInternal,
		ErrorCanNotDelLocalWins,
		ErrorStaticInit,
		ErrorIncBackup,
		ErrorFullBackup,
		ErrorRecNonExistent,
		ErrorRplNotAllowed,
		PeerdistErrorContentinfoVersionUnsupported,
		PeerdistErrorCannotParseContentinfo,
		PeerdistErrorMissingData,
		PeerdistErrorNoMore,
		PeerdistErrorNotInitialized,
		PeerdistErrorAlreadyInitialized,
		PeerdistErrorShutdownInProgress,
		PeerdistErrorInvalidated,
		PeerdistErrorAlreadyExists,
		PeerdistErrorOperationNotfound,
		PeerdistErrorAlreadyCompleted,
		PeerdistErrorOutOfBounds,
		PeerdistErrorVersionUnsupported,
		PeerdistErrorInvalidConfiguration,
		PeerdistErrorNotLicensed,
		PeerdistErrorServiceUnavailable,
		ErrorDhcpAddressConflict,
		ErrorWmiGuidNotFound,
		ErrorWmiInstanceNotFound,
		ErrorWmiItemidNotFound,
		ErrorWmiTryAgain,
		ErrorWmiDpNotFound,
		ErrorWmiUnresolvedInstanceRef,
		ErrorWmiAlreadyEnabled,
		ErrorWmiGuidDisconnected,
		ErrorWmiServerUnavailable,
		ErrorWmiDpFailed,
		ErrorWmiInvalidMof,
		ErrorWmiInvalidReginfo,
		ErrorWmiAlreadyDisabled,
		ErrorWmiReadOnly,
		ErrorWmiSetFailure,
		ErrorInvalidMedia,
		ErrorInvalidLibrary,
		ErrorInvalidMediaPool,
		ErrorDriveMediaMismatch,
		ErrorMediaOffline,
		ErrorLibraryOffline,
		ErrorEmpty,
		ErrorNotEmpty,
		ErrorMediaUnavailable,
		ErrorResourceDisabled,
		ErrorInvalidCleaner,
		ErrorUnableToClean,
		ErrorObjectNotFound,
		ErrorDatabaseFailure,
		ErrorDatabaseFull,
		ErrorMediaIncompatible,
		ErrorResourceNotPresent,
		ErrorInvalidOperation,
		ErrorMediaNotAvailable,
		ErrorDeviceNotAvailable,
		ErrorRequestRefused,
		ErrorInvalidDriveObject,
		ErrorLibraryFull,
		ErrorMediumNotAccessible,
		ErrorUnableToLoadMedium,
		ErrorUnableToInventoryDrive,
		ErrorUnableToInventorySlot,
		ErrorUnableToInventoryTransport,
		ErrorTransportFull,
		ErrorControllingIeport,
		ErrorUnableToEjectMountedMedia,
		ErrorCleanerSlotSet,
		ErrorCleanerSlotNotSet,
		ErrorCleanerCartridgeSpent,
		ErrorUnexpectedOmid,
		ErrorCantDeleteLastItem,
		ErrorMessageExceedsMaxSize,
		ErrorVolumeContainsSysFiles,
		ErrorIndigenousType,
		ErrorNoSupportingDrives,
		ErrorCleanerCartridgeInstalled,
		ErrorIeportFull,
		ErrorFileOffline,
		ErrorRemoteStorageNotActive,
		ErrorRemoteStorageMediaError,
		ErrorNotAReparsePoint,
		ErrorReparseAttributeConflict,
		ErrorInvalidReparseData,
		ErrorReparseTagInvalid,
		ErrorReparseTagMismatch,
		ErrorVolumeNotSisEnabled,
		ErrorDependentResourceExists,
		ErrorDependencyNotFound,
		ErrorDependencyAlreadyExists,
		ErrorResourceNotOnline,
		ErrorHostNodeNotAvailable,
		ErrorResourceNotAvailable,
		ErrorResourceNotFound,
		ErrorShutdownCluster,
		ErrorCantEvictActiveNode,
		ErrorObjectAlreadyExists,
		ErrorObjectInList,
		ErrorGroupNotAvailable,
		ErrorGroupNotFound,
		ErrorGroupNotOnline,
		ErrorHostNodeNotResourceOwner,
		ErrorHostNodeNotGroupOwner,
		ErrorResmonCreateFailed,
		ErrorResmonOnlineFailed,
		ErrorResourceOnline,
		ErrorQuorumResource,
		ErrorNotQuorumCapable,
		ErrorClusterShuttingDown,
		ErrorInvalidState,
		ErrorResourcePropertiesStored,
		ErrorNotQuorumClass,
		ErrorCoreResource,
		ErrorQuorumResourceOnlineFailed,
		ErrorQuorumlogOpenFailed,
		ErrorClusterlogCorrupt,
		ErrorClusterlogRecordExceedsMaxsize,
		ErrorClusterlogExceedsMaxsize,
		ErrorClusterlogChkpointNotFound,
		ErrorClusterlogNotEnoughSpace,
		ErrorQuorumOwnerAlive,
		ErrorNetworkNotAvailable,
		ErrorNodeNotAvailable,
		ErrorAllNodesNotAvailable,
		ErrorResourceFailed,
		ErrorClusterInvalidNode,
		ErrorClusterNodeExists,
		ErrorClusterJoinInProgress,
		ErrorClusterNodeNotFound,
		ErrorClusterLocalNodeNotFound,
		ErrorClusterNetworkExists,
		ErrorClusterNetworkNotFound,
		ErrorClusterNetinterfaceExists,
		ErrorClusterNetinterfaceNotFound,
		ErrorClusterInvalidRequest,
		ErrorClusterInvalidNetworkProvider,
		ErrorClusterNodeDown,
		ErrorClusterNodeUnreachable,
		ErrorClusterNodeNotMember,
		ErrorClusterJoinNotInProgress,
		ErrorClusterInvalidNetwork,
		ErrorClusterNodeUp,
		ErrorClusterIpaddrInUse,
		ErrorClusterNodeNotPaused,
		ErrorClusterNoSecurityContext,
		ErrorClusterNetworkNotInternal,
		ErrorClusterNodeAlreadyUp,
		ErrorClusterNodeAlreadyDown,
		ErrorClusterNetworkAlreadyOnline,
		ErrorClusterNetworkAlreadyOffline,
		ErrorClusterNodeAlreadyMember,
		ErrorClusterLastInternalNetwork,
		ErrorClusterNetworkHasDependents,
		ErrorInvalidOperationOnQuorum,
		ErrorDependencyNotAllowed,
		ErrorClusterNodePaused,
		ErrorNodeCantHostResource,
		ErrorClusterNodeNotReady,
		ErrorClusterNodeShuttingDown,
		ErrorClusterJoinAborted,
		ErrorClusterIncompatibleVersions,
		ErrorClusterMaxnumOfResourcesExceeded,
		ErrorClusterSystemConfigChanged,
		ErrorClusterResourceTypeNotFound,
		ErrorClusterRestypeNotSupported,
		ErrorClusterResnameNotFound,
		ErrorClusterNoRpcPackagesRegistered,
		ErrorClusterOwnerNotInPreflist,
		ErrorClusterDatabaseSeqmismatch,
		ErrorResmonInvalidState,
		ErrorClusterGumNotLocker,
		ErrorQuorumDiskNotFound,
		ErrorDatabaseBackupCorrupt,
		ErrorClusterNodeAlreadyHasDfsRoot,
		ErrorResourcePropertyUnchangeable,
		ErrorClusterMembershipInvalidState,
		ErrorClusterQuorumlogNotFound,
		ErrorClusterMembershipHalt,
		ErrorClusterInstanceIdMismatch,
		ErrorClusterNetworkNotFoundForIp,
		ErrorClusterPropertyDataTypeMismatch,
		ErrorClusterEvictWithoutCleanup,
		ErrorClusterParameterMismatch,
		ErrorNodeCannotBeClustered,
		ErrorClusterWrongOsVersion,
		ErrorClusterCantCreateDupClusterName,
		ErrorCluscfgAlreadyCommitted,
		ErrorCluscfgRollbackFailed,
		ErrorCluscfgSystemDiskDriveLetterConflict,
		ErrorClusterOldVersion,
		ErrorClusterMismatchedComputerAcctName,
		ErrorClusterNoNetAdapters,
		ErrorClusterPoisoned,
		ErrorClusterGroupMoving,
		ErrorClusterResourceTypeBusy,
		ErrorResourceCallTimedOut,
		ErrorInvalidClusterIpv6Address,
		ErrorClusterInternalInvalidFunction,
		ErrorClusterParameterOutOfBounds,
		ErrorClusterPartialSend,
		ErrorClusterRegistryInvalidFunction,
		ErrorClusterInvalidStringTermination,
		ErrorClusterInvalidStringFormat,
		ErrorClusterDatabaseTransactionInProgress,
		ErrorClusterDatabaseTransactionNotInProgress,
		ErrorClusterNullData,
		ErrorClusterPartialRead,
		ErrorClusterPartialWrite,
		ErrorClusterCantDeserializeData,
		ErrorDependentResourcePropertyConflict,
		ErrorClusterNoQuorum,
		ErrorClusterInvalidIpv6Network,
		ErrorClusterInvalidIpv6TunnelNetwork,
		ErrorQuorumNotAllowedInThisGroup,
		ErrorDependencyTreeTooComplex,
		ErrorExceptionInResourceCall,
		ErrorClusterRhsFailedInitialization,
		ErrorClusterNotInstalled,
		ErrorClusterResourcesMustBeOnlineOnTheSameNode,
		ErrorClusterMaxNodesInCluster,
		ErrorClusterTooManyNodes,
		ErrorClusterObjectAlreadyUsed,
		ErrorNoncoreGroupsFound,
		ErrorFileShareResourceConflict,
		ErrorClusterEvictInvalidRequest,
		ErrorClusterSingletonResource,
		ErrorClusterGroupSingletonResource,
		ErrorClusterResourceProviderFailed,
		ErrorClusterResourceConfigurationError,
		ErrorClusterGroupBusy,
		ErrorClusterNotSharedVolume,
		ErrorClusterInvalidSecurityDescriptor,
		ErrorClusterSharedVolumesInUse,
		ErrorClusterUseSharedVolumesApi,
		ErrorClusterBackupInProgress,
		ErrorNonCsvPath,
		ErrorCsvVolumeNotLocal,
		ErrorClusterWatchdogTerminating,
		ErrorEncryptionFailed,
		ErrorDecryptionFailed,
		ErrorFileEncrypted,
		ErrorNoRecoveryPolicy,
		ErrorNoEfs,
		ErrorWrongEfs,
		ErrorNoUserKeys,
		ErrorFileNotEncrypted,
		ErrorNotExportFormat,
		ErrorFileReadOnly,
		ErrorDirEfsDisallowed,
		ErrorEfsServerNotTrusted,
		ErrorBadRecoveryPolicy,
		ErrorEfsAlgBlobTooBig,
		ErrorVolumeNotSupportEfs,
		ErrorEfsDisabled,
		ErrorEfsVersionNotSupport,
		ErrorCsEncryptionInvalidServerResponse,
		ErrorCsEncryptionUnsupportedServer,
		ErrorCsEncryptionExistingEncryptedFile,
		ErrorCsEncryptionNewEncryptedFile,
		ErrorCsEncryptionFileNotCse,
		ErrorEncryptionPolicyDeniesOperation,
		ErrorNoBrowserServersFound,
		SchedEServiceNotLocalsystem,
		ErrorLogSectorInvalid,
		ErrorLogSectorParityInvalid,
		ErrorLogSectorRemapped,
		ErrorLogBlockIncomplete,
		ErrorLogInvalidRange,
		ErrorLogBlocksExhausted,
		ErrorLogReadContextInvalid,
		ErrorLogRestartInvalid,
		ErrorLogBlockVersion,
		ErrorLogBlockInvalid,
		ErrorLogReadModeInvalid,
		ErrorLogNoRestart,
		ErrorLogMetadataCorrupt,
		ErrorLogMetadataInvalid,
		ErrorLogMetadataInconsistent,
		ErrorLogReservationInvalid,
		ErrorLogCantDelete,
		ErrorLogContainerLimitExceeded,
		ErrorLogStartOfLog,
		ErrorLogPolicyAlreadyInstalled,
		ErrorLogPolicyNotInstalled,
		ErrorLogPolicyInvalid,
		ErrorLogPolicyConflict,
		ErrorLogPinnedArchiveTail,
		ErrorLogRecordNonexistent,
		ErrorLogRecordsReservedInvalid,
		ErrorLogSpaceReservedInvalid,
		ErrorLogTailInvalid,
		ErrorLogFull,
		ErrorCouldNotResizeLog,
		ErrorLogMultiplexed,
		ErrorLogDedicated,
		ErrorLogArchiveNotInProgress,
		ErrorLogArchiveInProgress,
		ErrorLogEphemeral,
		ErrorLogNotEnoughContainers,
		ErrorLogClientAlreadyRegistered,
		ErrorLogClientNotRegistered,
		ErrorLogFullHandlerInProgress,
		ErrorLogContainerReadFailed,
		ErrorLogContainerWriteFailed,
		ErrorLogContainerOpenFailed,
		ErrorLogContainerStateInvalid,
		ErrorLogStateInvalid,
		ErrorLogPinned,
		ErrorLogMetadataFlushFailed,
		ErrorLogInconsistentSecurity,
		ErrorLogAppendedFlushFailed,
		ErrorLogPinnedReservation,
		ErrorInvalidTransaction,
		ErrorTransactionNotActive,
		ErrorTransactionRequestNotValid,
		ErrorTransactionNotRequested,
		ErrorTransactionAlreadyAborted,
		ErrorTransactionAlreadyCommitted,
		ErrorTmInitializationFailed,
		ErrorResourcemanagerReadOnly,
		ErrorTransactionNotJoined,
		ErrorTransactionSuperiorExists,
		ErrorCrmProtocolAlreadyExists,
		ErrorTransactionPropagationFailed,
		ErrorCrmProtocolNotFound,
		ErrorTransactionInvalidMarshallBuffer,
		ErrorCurrentTransactionNotValid,
		ErrorTransactionNotFound,
		ErrorResourcemanagerNotFound,
		ErrorEnlistmentNotFound,
		ErrorTransactionmanagerNotFound,
		ErrorTransactionmanagerNotOnline,
		ErrorTransactionmanagerRecoveryNameCollision,
		ErrorTransactionNotRoot,
		ErrorTransactionObjectExpired,
		ErrorTransactionResponseNotEnlisted,
		ErrorTransactionRecordTooLong,
		ErrorImplicitTransactionNotSupported,
		ErrorTransactionIntegrityViolated,
		ErrorTransactionmanagerIdentityMismatch,
		ErrorRmCannotBeFrozenForSnapshot,
		ErrorTransactionMustWritethrough,
		ErrorTransactionNoSuperior,
		ErrorHeuristicDamagePossible,
		ErrorTransactionalConflict,
		ErrorRmNotActive,
		ErrorRmMetadataCorrupt,
		ErrorDirectoryNotRm,
		ErrorTransactionsUnsupportedRemote,
		ErrorLogResizeInvalidSize,
		ErrorObjectNoLongerExists,
		ErrorStreamMiniversionNotFound,
		ErrorStreamMiniversionNotValid,
		ErrorMiniversionInaccessibleFromSpecifiedTransaction,
		ErrorCantOpenMiniversionWithModifyIntent,
		ErrorCantCreateMoreStreamMiniversions,
		ErrorRemoteFileVersionMismatch,
		ErrorHandleNoLongerValid,
		ErrorNoTxfMetadata,
		ErrorLogCorruptionDetected,
		ErrorCantRecoverWithHandleOpen,
		ErrorRmDisconnected,
		ErrorEnlistmentNotSuperior,
		ErrorRecoveryNotNeeded,
		ErrorRmAlreadyStarted,
		ErrorFileIdentityNotPersistent,
		ErrorCantBreakTransactionalDependency,
		ErrorCantCrossRmBoundary,
		ErrorTxfDirNotEmpty,
		ErrorIndoubtTransactionsExist,
		ErrorTmVolatile,
		ErrorRollbackTimerExpired,
		ErrorTxfAttributeCorrupt,
		ErrorEfsNotAllowedInTransaction,
		ErrorTransactionalOpenNotAllowed,
		ErrorLogGrowthFailed,
		ErrorTransactedMappingUnsupportedRemote,
		ErrorTxfMetadataAlreadyPresent,
		ErrorTransactionScopeCallbacksNotSet,
		ErrorTransactionRequiredPromotion,
		ErrorCannotExecuteFileInTransaction,
		ErrorTransactionsNotFrozen,
		ErrorTransactionFreezeInProgress,
		ErrorNotSnapshotVolume,
		ErrorNoSavepointWithOpenFiles,
		ErrorDataLostRepair,
		ErrorSparseNotAllowedInTransaction,
		ErrorTmIdentityMismatch,
		ErrorFloatedSection,
		ErrorCannotAcceptTransactedWork,
		ErrorCannotAbortTransactions,
		ErrorBadClusters,
		ErrorCompressionNotAllowedInTransaction,
		ErrorVolumeDirty,
		ErrorNoLinkTrackingInTransaction,
		ErrorOperationNotSupportedInTransaction,
		ErrorExpiredHandle,
		ErrorTransactionNotEnlisted,
		ErrorCtxWinstationNameInvalid,
		ErrorCtxInvalidPd,
		ErrorCtxPdNotFound,
		ErrorCtxWdNotFound,
		ErrorCtxCannotMakeEventlogEntry,
		ErrorCtxServiceNameCollision,
		ErrorCtxClosePending,
		ErrorCtxNoOutbuf,
		ErrorCtxModemInfNotFound,
		ErrorCtxInvalidModemname,
		ErrorCtxModemResponseError,
		ErrorCtxModemResponseTimeout,
		ErrorCtxModemResponseNoCarrier,
		ErrorCtxModemResponseNoDialtone,
		ErrorCtxModemResponseBusy,
		ErrorCtxModemResponseVoice,
		ErrorCtxTdError,
		ErrorCtxWinstationNotFound,
		ErrorCtxWinstationAlreadyExists,
		ErrorCtxWinstationBusy,
		ErrorCtxBadVideoMode,
		ErrorCtxGraphicsInvalid,
		ErrorCtxLogonDisabled,
		ErrorCtxNotConsole,
		ErrorCtxClientQueryTimeout,
		ErrorCtxConsoleDisconnect,
		ErrorCtxConsoleConnect,
		ErrorCtxShadowDenied,
		ErrorCtxWinstationAccessDenied,
		ErrorCtxInvalidWd,
		ErrorCtxShadowInvalid,
		ErrorCtxShadowDisabled,
		ErrorCtxClientLicenseInUse,
		ErrorCtxClientLicenseNotSet,
		ErrorCtxLicenseNotAvailable,
		ErrorCtxLicenseClientInvalid,
		ErrorCtxLicenseExpired,
		ErrorCtxShadowNotRunning,
		ErrorCtxShadowEndedByModeChange,
		ErrorActivationCountExceeded,
		ErrorCtxWinstationsDisabled,
		ErrorCtxEncryptionLevelRequired,
		ErrorCtxSessionInUse,
		ErrorCtxNoForceLogoff,
		ErrorCtxAccountRestriction,
		ErrorRdpProtocolError,
		ErrorCtxCdmConnect,
		ErrorCtxCdmDisconnect,
		ErrorCtxSecurityLayerError,
		ErrorTsIncompatibleSessions,
		ErrorTsVideoSubsystemError,
		FrsErrInvalidApiSequence,
		FrsErrStartingService,
		FrsErrStoppingService,
		FrsErrInternalApi,
		FrsErrInternal,
		FrsErrServiceComm,
		FrsErrInsufficientPriv,
		FrsErrAuthentication,
		FrsErrParentInsufficientPriv,
		FrsErrParentAuthentication,
		FrsErrChildToParentComm,
		FrsErrParentToChildComm,
		FrsErrSysvolPopulate,
		FrsErrSysvolPopulateTimeout,
		FrsErrSysvolIsBusy,
		FrsErrSysvolDemote,
		FrsErrInvalidServiceParameter,
		ErrorDsNotInstalled,
		ErrorDsMembershipEvaluatedLocally,
		ErrorDsNoAttributeOrValue,
		ErrorDsInvalidAttributeSyntax,
		ErrorDsAttributeTypeUndefined,
		ErrorDsAttributeOrValueExists,
		ErrorDsBusy,
		ErrorDsUnavailable,
		ErrorDsNoRidsAllocated,
		ErrorDsNoMoreRids,
		ErrorDsIncorrectRoleOwner,
		ErrorDsRidmgrInitError,
		ErrorDsObjClassViolation,
		ErrorDsCantOnNonLeaf,
		ErrorDsCantOnRdn,
		ErrorDsCantModObjClass,
		ErrorDsCrossDomMoveError,
		ErrorDsGcNotAvailable,
		ErrorSharedPolicy,
		ErrorPolicyObjectNotFound,
		ErrorPolicyOnlyInDs,
		ErrorPromotionActive,
		ErrorNoPromotionActive,
		ErrorDsOperationsError,
		ErrorDsProtocolError,
		ErrorDsTimelimitExceeded,
		ErrorDsSizelimitExceeded,
		ErrorDsAdminLimitExceeded,
		ErrorDsCompareFalse,
		ErrorDsCompareTrue,
		ErrorDsAuthMethodNotSupported,
		ErrorDsStrongAuthRequired,
		ErrorDsInappropriateAuth,
		ErrorDsAuthUnknown,
		ErrorDsReferral,
		ErrorDsUnavailableCritExtension,
		ErrorDsConfidentialityRequired,
		ErrorDsInappropriateMatching,
		ErrorDsConstraintViolation,
		ErrorDsNoSuchObject,
		ErrorDsAliasProblem,
		ErrorDsInvalidDnSyntax,
		ErrorDsIsLeaf,
		ErrorDsAliasDerefProblem,
		ErrorDsUnwillingToPerform,
		ErrorDsLoopDetect,
		ErrorDsNamingViolation,
		ErrorDsObjectResultsTooLarge,
		ErrorDsAffectsMultipleDsas,
		ErrorDsServerDown,
		ErrorDsLocalError,
		ErrorDsEncodingError,
		ErrorDsDecodingError,
		ErrorDsFilterUnknown,
		ErrorDsParamError,
		ErrorDsNotSupported,
		ErrorDsNoResultsReturned,
		ErrorDsControlNotFound,
		ErrorDsClientLoop,
		ErrorDsReferralLimitExceeded,
		ErrorDsSortControlMissing,
		ErrorDsOffsetRangeError,
		ErrorDsRootMustBeNc,
		ErrorDsAddReplicaInhibited,
		ErrorDsAttNotDefInSchema,
		ErrorDsMaxObjSizeExceeded,
		ErrorDsObjStringNameExists,
		ErrorDsNoRdnDefinedInSchema,
		ErrorDsRdnDoesntMatchSchema,
		ErrorDsNoRequestedAttsFound,
		ErrorDsUserBufferToSmall,
		ErrorDsAttIsNotOnObj,
		ErrorDsIllegalModOperation,
		ErrorDsObjTooLarge,
		ErrorDsBadInstanceType,
		ErrorDsMasterdsaRequired,
		ErrorDsObjectClassRequired,
		ErrorDsMissingRequiredAtt,
		ErrorDsAttNotDefForClass,
		ErrorDsAttAlreadyExists,
		ErrorDsCantAddAttValues,
		ErrorDsSingleValueConstraint,
		ErrorDsRangeConstraint,
		ErrorDsAttValAlreadyExists,
		ErrorDsCantRemMissingAtt,
		ErrorDsCantRemMissingAttVal,
		ErrorDsRootCantBeSubref,
		ErrorDsNoChaining,
		ErrorDsNoChainedEval,
		ErrorDsNoParentObject,
		ErrorDsParentIsAnAlias,
		ErrorDsCantMixMasterAndReps,
		ErrorDsChildrenExist,
		ErrorDsObjNotFound,
		ErrorDsAliasedObjMissing,
		ErrorDsBadNameSyntax,
		ErrorDsAliasPointsToAlias,
		ErrorDsCantDerefAlias,
		ErrorDsOutOfScope,
		ErrorDsObjectBeingRemoved,
		ErrorDsCantDeleteDsaObj,
		ErrorDsGenericError,
		ErrorDsDsaMustBeIntMaster,
		ErrorDsClassNotDsa,
		ErrorDsInsuffAccessRights,
		ErrorDsIllegalSuperior,
		ErrorDsAttributeOwnedBySam,
		ErrorDsNameTooManyParts,
		ErrorDsNameTooLong,
		ErrorDsNameValueTooLong,
		ErrorDsNameUnparseable,
		ErrorDsNameTypeUnknown,
		ErrorDsNotAnObject,
		ErrorDsSecDescTooShort,
		ErrorDsSecDescInvalid,
		ErrorDsNoDeletedName,
		ErrorDsSubrefMustHaveParent,
		ErrorDsNcnameMustBeNc,
		ErrorDsCantAddSystemOnly,
		ErrorDsClassMustBeConcrete,
		ErrorDsInvalidDmd,
		ErrorDsObjGuidExists,
		ErrorDsNotOnBacklink,
		ErrorDsNoCrossrefForNc,
		ErrorDsShuttingDown,
		ErrorDsUnknownOperation,
		ErrorDsInvalidRoleOwner,
		ErrorDsCouldntContactFsmo,
		ErrorDsCrossNcDnRename,
		ErrorDsCantModSystemOnly,
		ErrorDsReplicatorOnly,
		ErrorDsObjClassNotDefined,
		ErrorDsObjClassNotSubclass,
		ErrorDsNameReferenceInvalid,
		ErrorDsCrossRefExists,
		ErrorDsCantDelMasterCrossref,
		ErrorDsSubtreeNotifyNotNcHead,
		ErrorDsNotifyFilterTooComplex,
		ErrorDsDupRdn,
		ErrorDsDupOid,
		ErrorDsDupMapiId,
		ErrorDsDupSchemaIdGuid,
		ErrorDsDupLdapDisplayName,
		ErrorDsSemanticAttTest,
		ErrorDsSyntaxMismatch,
		ErrorDsExistsInMustHave,
		ErrorDsExistsInMayHave,
		ErrorDsNonexistentMayHave,
		ErrorDsNonexistentMustHave,
		ErrorDsAuxClsTestFail,
		ErrorDsNonexistentPossSup,
		ErrorDsSubClsTestFail,
		ErrorDsBadRdnAttIdSyntax,
		ErrorDsExistsInAuxCls,
		ErrorDsExistsInSubCls,
		ErrorDsExistsInPossSup,
		ErrorDsRecalcschemaFailed,
		ErrorDsTreeDeleteNotFinished,
		ErrorDsCantDelete,
		ErrorDsAttSchemaReqId,
		ErrorDsBadAttSchemaSyntax,
		ErrorDsCantCacheAtt,
		ErrorDsCantCacheClass,
		ErrorDsCantRemoveAttCache,
		ErrorDsCantRemoveClassCache,
		ErrorDsCantRetrieveDn,
		ErrorDsMissingSupref,
		ErrorDsCantRetrieveInstance,
		ErrorDsCodeInconsistency,
		ErrorDsDatabaseError,
		ErrorDsGovernsidMissing,
		ErrorDsMissingExpectedAtt,
		ErrorDsNcnameMissingCrRef,
		ErrorDsSecurityCheckingError,
		ErrorDsSchemaNotLoaded,
		ErrorDsSchemaAllocFailed,
		ErrorDsAttSchemaReqSyntax,
		ErrorDsGcverifyError,
		ErrorDsDraSchemaMismatch,
		ErrorDsCantFindDsaObj,
		ErrorDsCantFindExpectedNc,
		ErrorDsCantFindNcInCache,
		ErrorDsCantRetrieveChild,
		ErrorDsSecurityIllegalModify,
		ErrorDsCantReplaceHiddenRec,
		ErrorDsBadHierarchyFile,
		ErrorDsBuildHierarchyTableFailed,
		ErrorDsConfigParamMissing,
		ErrorDsCountingAbIndicesFailed,
		ErrorDsHierarchyTableMallocFailed,
		ErrorDsInternalFailure,
		ErrorDsUnknownError,
		ErrorDsRootRequiresClassTop,
		ErrorDsRefusingFsmoRoles,
		ErrorDsMissingFsmoSettings,
		ErrorDsUnableToSurrenderRoles,
		ErrorDsDraGeneric,
		ErrorDsDraInvalidParameter,
		ErrorDsDraBusy,
		ErrorDsDraBadDn,
		ErrorDsDraBadNc,
		ErrorDsDraDnExists,
		ErrorDsDraInternalError,
		ErrorDsDraInconsistentDit,
		ErrorDsDraConnectionFailed,
		ErrorDsDraBadInstanceType,
		ErrorDsDraOutOfMem,
		ErrorDsDraMailProblem,
		ErrorDsDraRefAlreadyExists,
		ErrorDsDraRefNotFound,
		ErrorDsDraObjIsRepSource,
		ErrorDsDraDbError,
		ErrorDsDraNoReplica,
		ErrorDsDraAccessDenied,
		ErrorDsDraNotSupported,
		ErrorDsDraRpcCancelled,
		ErrorDsDraSourceDisabled,
		ErrorDsDraSinkDisabled,
		ErrorDsDraNameCollision,
		ErrorDsDraSourceReinstalled,
		ErrorDsDraMissingParent,
		ErrorDsDraPreempted,
		ErrorDsDraAbandonSync,
		ErrorDsDraShutdown,
		ErrorDsDraIncompatiblePartialSet,
		ErrorDsDraSourceIsPartialReplica,
		ErrorDsDraExtnConnectionFailed,
		ErrorDsInstallSchemaMismatch,
		ErrorDsDupLinkId,
		ErrorDsNameErrorResolving,
		ErrorDsNameErrorNotFound,
		ErrorDsNameErrorNotUnique,
		ErrorDsNameErrorNoMapping,
		ErrorDsNameErrorDomainOnly,
		ErrorDsNameErrorNoSyntacticalMapping,
		ErrorDsConstructedAttMod,
		ErrorDsWrongOmObjClass,
		ErrorDsDraReplPending,
		ErrorDsDsRequired,
		ErrorDsInvalidLdapDisplayName,
		ErrorDsNonBaseSearch,
		ErrorDsCantRetrieveAtts,
		ErrorDsBacklinkWithoutLink,
		ErrorDsEpochMismatch,
		ErrorDsSrcNameMismatch,
		ErrorDsSrcAndDstNcIdentical,
		ErrorDsDstNcMismatch,
		ErrorDsNotAuthoritiveForDstNc,
		ErrorDsSrcGuidMismatch,
		ErrorDsCantMoveDeletedObject,
		ErrorDsPdcOperationInProgress,
		ErrorDsCrossDomainCleanupReqd,
		ErrorDsIllegalXdomMoveOperation,
		ErrorDsCantWithAcctGroupMembershps,
		ErrorDsNcMustHaveNcParent,
		ErrorDsCrImpossibleToValidate,
		ErrorDsDstDomainNotNative,
		ErrorDsMissingInfrastructureContainer,
		ErrorDsCantMoveAccountGroup,
		ErrorDsCantMoveResourceGroup,
		ErrorDsInvalidSearchFlag,
		ErrorDsNoTreeDeleteAboveNc,
		ErrorDsCouldntLockTreeForDelete,
		ErrorDsCouldntIdentifyObjectsForTreeDelete,
		ErrorDsSamInitFailure,
		ErrorDsSensitiveGroupViolation,
		ErrorDsCantModPrimarygroupid,
		ErrorDsIllegalBaseSchemaMod,
		ErrorDsNonsafeSchemaChange,
		ErrorDsSchemaUpdateDisallowed,
		ErrorDsCantCreateUnderSchema,
		ErrorDsInstallNoSrcSchVersion,
		ErrorDsInstallNoSchVersionInInifile,
		ErrorDsInvalidGroupType,
		ErrorDsNoNestGlobalgroupInMixeddomain,
		ErrorDsNoNestLocalgroupInMixeddomain,
		ErrorDsGlobalCantHaveLocalMember,
		ErrorDsGlobalCantHaveUniversalMember,
		ErrorDsUniversalCantHaveLocalMember,
		ErrorDsGlobalCantHaveCrossdomainMember,
		ErrorDsLocalCantHaveCrossdomainLocalMember,
		ErrorDsHavePrimaryMembers,
		ErrorDsStringSdConversionFailed,
		ErrorDsNamingMasterGc,
		ErrorDsDnsLookupFailure,
		ErrorDsCouldntUpdateSpns,
		ErrorDsCantRetrieveSd,
		ErrorDsKeyNotUnique,
		ErrorDsWrongLinkedAttSyntax,
		ErrorDsSamNeedBootkeyPassword,
		ErrorDsSamNeedBootkeyFloppy,
		ErrorDsCantStart,
		ErrorDsInitFailure,
		ErrorDsNoPktPrivacyOnConnection,
		ErrorDsSourceDomainInForest,
		ErrorDsDestinationDomainNotInForest,
		ErrorDsDestinationAuditingNotEnabled,
		ErrorDsCantFindDcForSrcDomain,
		ErrorDsSrcObjNotGroupOrUser,
		ErrorDsSrcSidExistsInForest,
		ErrorDsSrcAndDstObjectClassMismatch,
		ErrorSamInitFailure,
		ErrorDsDraSchemaInfoShip,
		ErrorDsDraSchemaConflict,
		ErrorDsDraEarlierSchemaConflict,
		ErrorDsDraObjNcMismatch,
		ErrorDsNcStillHasDsas,
		ErrorDsGcRequired,
		ErrorDsLocalMemberOfLocalOnly,
		ErrorDsNoFpoInUniversalGroups,
		ErrorDsCantAddToGc,
		ErrorDsNoCheckpointWithPdc,
		ErrorDsSourceAuditingNotEnabled,
		ErrorDsCantCreateInNondomainNc,
		ErrorDsInvalidNameForSpn,
		ErrorDsFilterUsesContructedAttrs,
		ErrorDsUnicodepwdNotInQuotes,
		ErrorDsMachineAccountQuotaExceeded,
		ErrorDsMustBeRunOnDstDc,
		ErrorDsSrcDcMustBeSp4OrGreater,
		ErrorDsCantTreeDeleteCriticalObj,
		ErrorDsInitFailureConsole,
		ErrorDsSamInitFailureConsole,
		ErrorDsForestVersionTooHigh,
		ErrorDsDomainVersionTooHigh,
		ErrorDsForestVersionTooLow,
		ErrorDsDomainVersionTooLow,
		ErrorDsIncompatibleVersion,
		ErrorDsLowDsaVersion,
		ErrorDsNoBehaviorVersionInMixeddomain,
		ErrorDsNotSupportedSortOrder,
		ErrorDsNameNotUnique,
		ErrorDsMachineAccountCreatedPrent4,
		ErrorDsOutOfVersionStore,
		ErrorDsIncompatibleControlsUsed,
		ErrorDsNoRefDomain,
		ErrorDsReservedLinkId,
		ErrorDsLinkIdNotAvailable,
		ErrorDsAgCantHaveUniversalMember,
		ErrorDsModifydnDisallowedByInstanceType,
		ErrorDsNoObjectMoveInSchemaNc,
		ErrorDsModifydnDisallowedByFlag,
		ErrorDsModifydnWrongGrandparent,
		ErrorDsNameErrorTrustReferral,
		ErrorNotSupportedOnStandardServer,
		ErrorDsCantAccessRemotePartOfAd,
		ErrorDsCrImpossibleToValidateV2,
		ErrorDsThreadLimitExceeded,
		ErrorDsNotClosest,
		ErrorDsCantDeriveSpnWithoutServerRef,
		ErrorDsSingleUserModeFailed,
		ErrorDsNtdscriptSyntaxError,
		ErrorDsNtdscriptProcessError,
		ErrorDsDifferentReplEpochs,
		ErrorDsDrsExtensionsChanged,
		ErrorDsReplicaSetChangeNotAllowedOnDisabledCr,
		ErrorDsNoMsdsIntid,
		ErrorDsDupMsdsIntid,
		ErrorDsExistsInRdnattid,
		ErrorDsAuthorizationFailed,
		ErrorDsInvalidScript,
		ErrorDsRemoteCrossrefOpFailed,
		ErrorDsCrossRefBusy,
		ErrorDsCantDeriveSpnForDeletedDomain,
		ErrorDsCantDemoteWithWriteableNc,
		ErrorDsDuplicateIdFound,
		ErrorDsInsufficientAttrToCreateObject,
		ErrorDsGroupConversionError,
		ErrorDsCantMoveAppBasicGroup,
		ErrorDsCantMoveAppQueryGroup,
		ErrorDsRoleNotVerified,
		ErrorDsWkoContainerCannotBeSpecial,
		ErrorDsDomainRenameInProgress,
		ErrorDsExistingAdChildNc,
		ErrorDsReplLifetimeExceeded,
		ErrorDsDisallowedInSystemContainer,
		ErrorDsLdapSendQueueFull,
		ErrorDsDraOutScheduleWindow,
		ErrorDsPolicyNotKnown,
		ErrorNoSiteSettingsObject,
		ErrorNoSecrets,
		ErrorNoWritableDcFound,
		ErrorDsNoServerObject,
		ErrorDsNoNtdsaObject,
		ErrorDsNonAsqSearch,
		ErrorDsAuditFailure,
		ErrorDsInvalidSearchFlagSubtree,
		ErrorDsInvalidSearchFlagTuple,
		ErrorDsHierarchyTableTooDeep,
		ErrorDsDraCorruptUtdVector,
		ErrorDsDraSecretsDenied,
		ErrorDsReservedMapiId,
		ErrorDsMapiIdNotAvailable,
		ErrorDsDraMissingKrbtgtSecret,
		ErrorDsDomainNameExistsInForest,
		ErrorDsFlatNameExistsInForest,
		ErrorInvalidUserPrincipalName,
		ErrorDsOidMappedGroupCantHaveMembers,
		ErrorDsOidNotFound,
		ErrorDsDraRecycledTarget,
		DnsErrorResponseCodesBase,
		DnsErrorMask,
		DnsErrorRcodeFormatError,
		DnsErrorRcodeServerFailure,
		DnsErrorRcodeNameError,
		DnsErrorRcodeNotImplemented,
		DnsErrorRcodeRefused,
		DnsErrorRcodeYxdomain,
		DnsErrorRcodeYxrrset,
		DnsErrorRcodeNxrrset,
		DnsErrorRcodeNotauth,
		DnsErrorRcodeNotzone,
		DnsErrorRcodeBadsig,
		DnsErrorRcodeBadkey,
		DnsErrorRcodeBadtime,
		DnsErrorPacketFmtBase,
		DnsInfoNoRecords,
		DnsErrorBadPacket,
		DnsErrorNoPacket,
		DnsErrorRcode,
		DnsErrorUnsecurePacket,
		DnsErrorGeneralApiBase,
		DnsErrorInvalidType,
		DnsErrorInvalidIpAddress,
		DnsErrorInvalidProperty,
		DnsErrorTryAgainLater,
		DnsErrorNotUnique,
		DnsErrorNonRfcName,
		DnsStatusFqdn,
		DnsStatusDottedName,
		DnsStatusSinglePartName,
		DnsErrorInvalidNameChar,
		DnsErrorNumericName,
		DnsErrorNotAllowedOnRootServer,
		DnsErrorNotAllowedUnderDelegation,
		DnsErrorCannotFindRootHints,
		DnsErrorInconsistentRootHints,
		DnsErrorDwordValueTooSmall,
		DnsErrorDwordValueTooLarge,
		DnsErrorBackgroundLoading,
		DnsErrorNotAllowedOnRodc,
		DnsErrorNotAllowedUnderDname,
		DnsErrorDelegationRequired,
		DnsErrorInvalidPolicyTable,
		DnsErrorZoneBase,
		DnsErrorZoneDoesNotExist,
		DnsErrorNoZoneInfo,
		DnsErrorInvalidZoneOperation,
		DnsErrorZoneConfigurationError,
		DnsErrorZoneHasNoSoaRecord,
		DnsErrorZoneHasNoNsRecords,
		DnsErrorZoneLocked,
		DnsErrorZoneCreationFailed,
		DnsErrorZoneAlreadyExists,
		DnsErrorAutozoneAlreadyExists,
		DnsErrorInvalidZoneType,
		DnsErrorSecondaryRequiresMasterIp,
		DnsErrorZoneNotSecondary,
		DnsErrorNeedSecondaryAddresses,
		DnsErrorWinsInitFailed,
		DnsErrorNeedWinsServers,
		DnsErrorNbstatInitFailed,
		DnsErrorSoaDeleteInvalid,
		DnsErrorForwarderAlreadyExists,
		DnsErrorZoneRequiresMasterIp,
		DnsErrorZoneIsShutdown,
		DnsErrorDatafileBase,
		DnsErrorPrimaryRequiresDatafile,
		DnsErrorInvalidDatafileName,
		DnsErrorDatafileOpenFailure,
		DnsErrorFileWritebackFailed,
		DnsErrorDatafileParsing,
		DnsErrorDatabaseBase,
		DnsErrorRecordDoesNotExist,
		DnsErrorRecordFormat,
		DnsErrorNodeCreationFailed,
		DnsErrorUnknownRecordType,
		DnsErrorRecordTimedOut,
		DnsErrorNameNotInZone,
		DnsErrorCnameLoop,
		DnsErrorNodeIsCname,
		DnsErrorCnameCollision,
		DnsErrorRecordOnlyAtZoneRoot,
		DnsErrorRecordAlreadyExists,
		DnsErrorSecondaryData,
		DnsErrorNoCreateCacheData,
		DnsErrorNameDoesNotExist,
		DnsWarningPtrCreateFailed,
		DnsWarningDomainUndeleted,
		DnsErrorDsUnavailable,
		DnsErrorDsZoneAlreadyExists,
		DnsErrorNoBootfileIfDsZone,
		DnsErrorNodeIsDname,
		DnsErrorDnameCollision,
		DnsErrorAliasLoop,
		DnsErrorOperationBase,
		DnsInfoAxfrComplete,
		DnsErrorAxfr,
		DnsInfoAddedLocalWins,
		DnsErrorSecureBase,
		DnsStatusContinueNeeded,
		DnsErrorSetupBase,
		DnsErrorNoTcpip,
		DnsErrorNoDnsServers,
		DnsErrorDpBase,
		DnsErrorDpDoesNotExist,
		DnsErrorDpAlreadyExists,
		DnsErrorDpNotEnlisted,
		DnsErrorDpAlreadyEnlisted,
		DnsErrorDpNotAvailable,
		DnsErrorDpFsmoError,
		Wsabaseerr,
		Wsaeintr,
		Wsaebadf,
		Wsaeacces,
		Wsaefault,
		Wsaeinval,
		Wsaemfile,
		Wsaewouldblock,
		Wsaeinprogress,
		Wsaealready,
		Wsaenotsock,
		Wsaedestaddrreq,
		Wsaemsgsize,
		Wsaeprototype,
		Wsaenoprotoopt,
		Wsaeprotonosupport,
		Wsaesocktnosupport,
		Wsaeopnotsupp,
		Wsaepfnosupport,
		Wsaeafnosupport,
		Wsaeaddrinuse,
		Wsaeaddrnotavail,
		Wsaenetdown,
		Wsaenetunreach,
		Wsaenetreset,
		Wsaeconnaborted,
		Wsaeconnreset,
		Wsaenobufs,
		Wsaeisconn,
		Wsaenotconn,
		Wsaeshutdown,
		Wsaetoomanyrefs,
		Wsaetimedout,
		Wsaeconnrefused,
		Wsaeloop,
		Wsaenametoolong,
		Wsaehostdown,
		Wsaehostunreach,
		Wsaenotempty,
		Wsaeproclim,
		Wsaeusers,
		Wsaedquot,
		Wsaestale,
		Wsaeremote,
		Wsasysnotready,
		Wsavernotsupported,
		Wsanotinitialised,
		Wsaediscon,
		Wsaenomore,
		Wsaecancelled,
		Wsaeinvalidproctable,
		Wsaeinvalidprovider,
		Wsaeproviderfailedinit,
		Wsasyscallfailure,
		WsaserviceNotFound,
		WsatypeNotFound,
		WsaENoMore,
		WsaECancelled,
		Wsaerefused,
		WsahostNotFound,
		WsatryAgain,
		WsanoRecovery,
		WsanoData,
		WsaQosReceivers,
		WsaQosSenders,
		WsaQosNoSenders,
		WsaQosNoReceivers,
		WsaQosRequestConfirmed,
		WsaQosAdmissionFailure,
		WsaQosPolicyFailure,
		WsaQosBadStyle,
		WsaQosBadObject,
		WsaQosTrafficCtrlError,
		WsaQosGenericError,
		WsaQosEservicetype,
		WsaQosEflowspec,
		WsaQosEprovspecbuf,
		WsaQosEfilterstyle,
		WsaQosEfiltertype,
		WsaQosEfiltercount,
		WsaQosEobjlength,
		WsaQosEflowcount,
		WsaQosEunkownpsobj,
		WsaQosEpolicyobj,
		WsaQosEflowdesc,
		WsaQosEpsflowspec,
		WsaQosEpsfilterspec,
		WsaQosEsdmodeobj,
		WsaQosEshaperateobj,
		WsaQosReservedPetype,
		WsaSecureHostNotFound,
		WsaIpsecNamePolicyError,
		ErrorIpsecQmPolicyExists,
		ErrorIpsecQmPolicyNotFound,
		ErrorIpsecQmPolicyInUse,
		ErrorIpsecMmPolicyExists,
		ErrorIpsecMmPolicyNotFound,
		ErrorIpsecMmPolicyInUse,
		ErrorIpsecMmFilterExists,
		ErrorIpsecMmFilterNotFound,
		ErrorIpsecTransportFilterExists,
		ErrorIpsecTransportFilterNotFound,
		ErrorIpsecMmAuthExists,
		ErrorIpsecMmAuthNotFound,
		ErrorIpsecMmAuthInUse,
		ErrorIpsecDefaultMmPolicyNotFound,
		ErrorIpsecDefaultMmAuthNotFound,
		ErrorIpsecDefaultQmPolicyNotFound,
		ErrorIpsecTunnelFilterExists,
		ErrorIpsecTunnelFilterNotFound,
		ErrorIpsecMmFilterPendingDeletion,
		ErrorIpsecTransportFilterPendingDeletion,
		ErrorIpsecTunnelFilterPendingDeletion,
		ErrorIpsecMmPolicyPendingDeletion,
		ErrorIpsecMmAuthPendingDeletion,
		ErrorIpsecQmPolicyPendingDeletion,
		WarningIpsecMmPolicyPruned,
		WarningIpsecQmPolicyPruned,
		ErrorIpsecIkeNegStatusBegin,
		ErrorIpsecIkeAuthFail,
		ErrorIpsecIkeAttribFail,
		ErrorIpsecIkeNegotiationPending,
		ErrorIpsecIkeGeneralProcessingError,
		ErrorIpsecIkeTimedOut,
		ErrorIpsecIkeNoCert,
		ErrorIpsecIkeSaDeleted,
		ErrorIpsecIkeSaReaped,
		ErrorIpsecIkeMmAcquireDrop,
		ErrorIpsecIkeQmAcquireDrop,
		ErrorIpsecIkeQueueDropMm,
		ErrorIpsecIkeQueueDropNoMm,
		ErrorIpsecIkeDropNoResponse,
		ErrorIpsecIkeMmDelayDrop,
		ErrorIpsecIkeQmDelayDrop,
		ErrorIpsecIkeError,
		ErrorIpsecIkeCrlFailed,
		ErrorIpsecIkeInvalidKeyUsage,
		ErrorIpsecIkeInvalidCertType,
		ErrorIpsecIkeNoPrivateKey,
		ErrorIpsecIkeSimultaneousRekey,
		ErrorIpsecIkeDhFail,
		ErrorIpsecIkeCriticalPayloadNotRecognized,
		ErrorIpsecIkeInvalidHeader,
		ErrorIpsecIkeNoPolicy,
		ErrorIpsecIkeInvalidSignature,
		ErrorIpsecIkeKerberosError,
		ErrorIpsecIkeNoPublicKey,
		ErrorIpsecIkeProcessErr,
		ErrorIpsecIkeProcessErrSa,
		ErrorIpsecIkeProcessErrProp,
		ErrorIpsecIkeProcessErrTrans,
		ErrorIpsecIkeProcessErrKe,
		ErrorIpsecIkeProcessErrId,
		ErrorIpsecIkeProcessErrCert,
		ErrorIpsecIkeProcessErrCertReq,
		ErrorIpsecIkeProcessErrHash,
		ErrorIpsecIkeProcessErrSig,
		ErrorIpsecIkeProcessErrNonce,
		ErrorIpsecIkeProcessErrNotify,
		ErrorIpsecIkeProcessErrDelete,
		ErrorIpsecIkeProcessErrVendor,
		ErrorIpsecIkeInvalidPayload,
		ErrorIpsecIkeLoadSoftSa,
		ErrorIpsecIkeSoftSaTornDown,
		ErrorIpsecIkeInvalidCookie,
		ErrorIpsecIkeNoPeerCert,
		ErrorIpsecIkePeerCrlFailed,
		ErrorIpsecIkePolicyChange,
		ErrorIpsecIkeNoMmPolicy,
		ErrorIpsecIkeNotcbpriv,
		ErrorIpsecIkeSecloadfail,
		ErrorIpsecIkeFailsspinit,
		ErrorIpsecIkeFailqueryssp,
		ErrorIpsecIkeSrvacqfail,
		ErrorIpsecIkeSrvquerycred,
		ErrorIpsecIkeGetspifail,
		ErrorIpsecIkeInvalidFilter,
		ErrorIpsecIkeOutOfMemory,
		ErrorIpsecIkeAddUpdateKeyFailed,
		ErrorIpsecIkeInvalidPolicy,
		ErrorIpsecIkeUnknownDoi,
		ErrorIpsecIkeInvalidSituation,
		ErrorIpsecIkeDhFailure,
		ErrorIpsecIkeInvalidGroup,
		ErrorIpsecIkeEncrypt,
		ErrorIpsecIkeDecrypt,
		ErrorIpsecIkePolicyMatch,
		ErrorIpsecIkeUnsupportedId,
		ErrorIpsecIkeInvalidHash,
		ErrorIpsecIkeInvalidHashAlg,
		ErrorIpsecIkeInvalidHashSize,
		ErrorIpsecIkeInvalidEncryptAlg,
		ErrorIpsecIkeInvalidAuthAlg,
		ErrorIpsecIkeInvalidSig,
		ErrorIpsecIkeLoadFailed,
		ErrorIpsecIkeRpcDelete,
		ErrorIpsecIkeBenignReinit,
		ErrorIpsecIkeInvalidResponderLifetimeNotify,
		ErrorIpsecIkeInvalidMajorVersion,
		ErrorIpsecIkeInvalidCertKeylen,
		ErrorIpsecIkeMmLimit,
		ErrorIpsecIkeNegotiationDisabled,
		ErrorIpsecIkeQmLimit,
		ErrorIpsecIkeMmExpired,
		ErrorIpsecIkePeerMmAssumedInvalid,
		ErrorIpsecIkeCertChainPolicyMismatch,
		ErrorIpsecIkeUnexpectedMessageId,
		ErrorIpsecIkeInvalidAuthPayload,
		ErrorIpsecIkeDosCookieSent,
		ErrorIpsecIkeShuttingDown,
		ErrorIpsecIkeCgaAuthFailed,
		ErrorIpsecIkeProcessErrNatoa,
		ErrorIpsecIkeInvalidMmForQm,
		ErrorIpsecIkeQmExpired,
		ErrorIpsecIkeTooManyFilters,
		ErrorIpsecIkeNegStatusEnd,
		ErrorIpsecIkeKillDummyNapTunnel,
		ErrorIpsecIkeInnerIpAssignmentFailure,
		ErrorIpsecIkeRequireCpPayloadMissing,
		ErrorIpsecKeyModuleImpersonationNegotiationPending,
		ErrorIpsecIkeCoexistenceSuppress,
		ErrorIpsecIkeRatelimitDrop,
		ErrorIpsecIkePeerDoesntSupportMobike,
		ErrorIpsecIkeAuthorizationFailure,
		ErrorIpsecIkeStrongCredAuthorizationFailure,
		ErrorIpsecIkeAuthorizationFailureWithOptionalRetry,
		ErrorIpsecIkeStrongCredAuthorizationAndCertmapFailure,
		ErrorIpsecIkeNegStatusExtendedEnd,
		ErrorIpsecBadSpi,
		ErrorIpsecSaLifetimeExpired,
		ErrorIpsecWrongSa,
		ErrorIpsecReplayCheckFailed,
		ErrorIpsecInvalidPacket,
		ErrorIpsecIntegrityCheckFailed,
		ErrorIpsecClearTextDrop,
		ErrorIpsecAuthFirewallDrop,
		ErrorIpsecThrottleDrop,
		ErrorIpsecDospBlock,
		ErrorIpsecDospReceivedMulticast,
		ErrorIpsecDospInvalidPacket,
		ErrorIpsecDospStateLookupFailed,
		ErrorIpsecDospMaxEntries,
		ErrorIpsecDospKeymodNotAllowed,
		ErrorIpsecDospNotInstalled,
		ErrorIpsecDospMaxPerIpRatelimitQueues,
		ErrorSxsSectionNotFound,
		ErrorSxsCantGenActctx,
		ErrorSxsInvalidActctxdataFormat,
		ErrorSxsAssemblyNotFound,
		ErrorSxsManifestFormatError,
		ErrorSxsManifestParseError,
		ErrorSxsActivationContextDisabled,
		ErrorSxsKeyNotFound,
		ErrorSxsVersionConflict,
		ErrorSxsWrongSectionType,
		ErrorSxsThreadQueriesDisabled,
		ErrorSxsProcessDefaultAlreadySet,
		ErrorSxsUnknownEncodingGroup,
		ErrorSxsUnknownEncoding,
		ErrorSxsInvalidXmlNamespaceUri,
		ErrorSxsRootManifestDependencyNotInstalled,
		ErrorSxsLeafManifestDependencyNotInstalled,
		ErrorSxsInvalidAssemblyIdentityAttribute,
		ErrorSxsManifestMissingRequiredDefaultNamespace,
		ErrorSxsManifestInvalidRequiredDefaultNamespace,
		ErrorSxsPrivateManifestCrossPathWithReparsePoint,
		ErrorSxsDuplicateDllName,
		ErrorSxsDuplicateWindowclassName,
		ErrorSxsDuplicateClsid,
		ErrorSxsDuplicateIid,
		ErrorSxsDuplicateTlbid,
		ErrorSxsDuplicateProgid,
		ErrorSxsDuplicateAssemblyName,
		ErrorSxsFileHashMismatch,
		ErrorSxsPolicyParseError,
		ErrorSxsXmlEMissingquote,
		ErrorSxsXmlECommentsyntax,
		ErrorSxsXmlEBadstartnamechar,
		ErrorSxsXmlEBadnamechar,
		ErrorSxsXmlEBadcharinstring,
		ErrorSxsXmlEXmldeclsyntax,
		ErrorSxsXmlEBadchardata,
		ErrorSxsXmlEMissingwhitespace,
		ErrorSxsXmlEExpectingtagend,
		ErrorSxsXmlEMissingsemicolon,
		ErrorSxsXmlEUnbalancedparen,
		ErrorSxsXmlEInternalerror,
		ErrorSxsXmlEUnexpectedWhitespace,
		ErrorSxsXmlEIncompleteEncoding,
		ErrorSxsXmlEMissingParen,
		ErrorSxsXmlEExpectingclosequote,
		ErrorSxsXmlEMultipleColons,
		ErrorSxsXmlEInvalidDecimal,
		ErrorSxsXmlEInvalidHexidecimal,
		ErrorSxsXmlEInvalidUnicode,
		ErrorSxsXmlEWhitespaceorquestionmark,
		ErrorSxsXmlEUnexpectedendtag,
		ErrorSxsXmlEUnclosedtag,
		ErrorSxsXmlEDuplicateattribute,
		ErrorSxsXmlEMultipleroots,
		ErrorSxsXmlEInvalidatrootlevel,
		ErrorSxsXmlEBadxmldecl,
		ErrorSxsXmlEMissingroot,
		ErrorSxsXmlEUnexpectedeof,
		ErrorSxsXmlEBadperefinsubset,
		ErrorSxsXmlEUnclosedstarttag,
		ErrorSxsXmlEUnclosedendtag,
		ErrorSxsXmlEUnclosedstring,
		ErrorSxsXmlEUnclosedcomment,
		ErrorSxsXmlEUncloseddecl,
		ErrorSxsXmlEUnclosedcdata,
		ErrorSxsXmlEReservednamespace,
		ErrorSxsXmlEInvalidencoding,
		ErrorSxsXmlEInvalidswitch,
		ErrorSxsXmlEBadxmlcase,
		ErrorSxsXmlEInvalidStandalone,
		ErrorSxsXmlEUnexpectedStandalone,
		ErrorSxsXmlEInvalidVersion,
		ErrorSxsXmlEMissingequals,
		ErrorSxsProtectionRecoveryFailed,
		ErrorSxsProtectionPublicKeyTooShort,
		ErrorSxsProtectionCatalogNotValid,
		ErrorSxsUntranslatableHresult,
		ErrorSxsProtectionCatalogFileMissing,
		ErrorSxsMissingAssemblyIdentityAttribute,
		ErrorSxsInvalidAssemblyIdentityAttributeName,
		ErrorSxsAssemblyMissing,
		ErrorSxsCorruptActivationStack,
		ErrorSxsCorruption,
		ErrorSxsEarlyDeactivation,
		ErrorSxsInvalidDeactivation,
		ErrorSxsMultipleDeactivation,
		ErrorSxsProcessTerminationRequested,
		ErrorSxsReleaseActivationContext,
		ErrorSxsSystemDefaultActivationContextEmpty,
		ErrorSxsInvalidIdentityAttributeValue,
		ErrorSxsInvalidIdentityAttributeName,
		ErrorSxsIdentityDuplicateAttribute,
		ErrorSxsIdentityParseError,
		ErrorMalformedSubstitutionString,
		ErrorSxsIncorrectPublicKeyToken,
		ErrorUnmappedSubstitutionString,
		ErrorSxsAssemblyNotLocked,
		ErrorSxsComponentStoreCorrupt,
		ErrorAdvancedInstallerFailed,
		ErrorXmlEncodingMismatch,
		ErrorSxsManifestIdentitySameButContentsDifferent,
		ErrorSxsIdentitiesDifferent,
		ErrorSxsAssemblyIsNotADeployment,
		ErrorSxsFileNotPartOfAssembly,
		ErrorSxsManifestTooBig,
		ErrorSxsSettingNotRegistered,
		ErrorSxsTransactionClosureIncomplete,
		ErrorSmiPrimitiveInstallerFailed,
		ErrorGenericCommandFailed,
		ErrorSxsFileHashMissing,
		ErrorEvtInvalidChannelPath,
		ErrorEvtInvalidQuery,
		ErrorEvtPublisherMetadataNotFound,
		ErrorEvtEventTemplateNotFound,
		ErrorEvtInvalidPublisherName,
		ErrorEvtInvalidEventData,
		ErrorEvtChannelNotFound,
		ErrorEvtMalformedXmlText,
		ErrorEvtSubscriptionToDirectChannel,
		ErrorEvtConfigurationError,
		ErrorEvtQueryResultStale,
		ErrorEvtQueryResultInvalidPosition,
		ErrorEvtNonValidatingMsxml,
		ErrorEvtFilterAlreadyscoped,
		ErrorEvtFilterNoteltset,
		ErrorEvtFilterInvarg,
		ErrorEvtFilterInvtest,
		ErrorEvtFilterInvtype,
		ErrorEvtFilterParseerr,
		ErrorEvtFilterUnsupportedop,
		ErrorEvtFilterUnexpectedtoken,
		ErrorEvtInvalidOperationOverEnabledDirectChannel,
		ErrorEvtInvalidChannelPropertyValue,
		ErrorEvtInvalidPublisherPropertyValue,
		ErrorEvtChannelCannotActivate,
		ErrorEvtFilterTooComplex,
		ErrorEvtMessageNotFound,
		ErrorEvtMessageIdNotFound,
		ErrorEvtUnresolvedValueInsert,
		ErrorEvtUnresolvedParameterInsert,
		ErrorEvtMaxInsertsReached,
		ErrorEvtEventDefinitionNotFound,
		ErrorEvtMessageLocaleNotFound,
		ErrorEvtVersionTooOld,
		ErrorEvtVersionTooNew,
		ErrorEvtCannotOpenChannelOfQuery,
		ErrorEvtPublisherDisabled,
		ErrorEvtFilterOutOfRange,
		ErrorEcSubscriptionCannotActivate,
		ErrorEcLogDisabled,
		ErrorEcCircularForwarding,
		ErrorEcCredstoreFull,
		ErrorEcCredNotFound,
		ErrorEcNoActiveChannel,
		ErrorMuiFileNotFound,
		ErrorMuiInvalidFile,
		ErrorMuiInvalidRcConfig,
		ErrorMuiInvalidLocaleName,
		ErrorMuiInvalidUltimatefallbackName,
		ErrorMuiFileNotLoaded,
		ErrorResourceEnumUserStop,
		ErrorMuiIntlsettingsUilangNotInstalled,
		ErrorMuiIntlsettingsInvalidLocaleName,
		ErrorMcaInvalidCapabilitiesString,
		ErrorMcaInvalidVcpVersion,
		ErrorMcaMonitorViolatesMccsSpecification,
		ErrorMcaMccsVersionMismatch,
		ErrorMcaUnsupportedMccsVersion,
		ErrorMcaInternalError,
		ErrorMcaInvalidTechnologyTypeReturned,
		ErrorMcaUnsupportedColorTemperature,
		ErrorAmbiguousSystemDevice,
		ErrorSystemDeviceNotFound,
		ErrorHashNotSupported,
		ErrorHashNotPresent,
		SeveritySuccess,
		SeverityError,
		FacilityNtBit,
		EUnexpected,
		ENotimpl,
		EOutofmemory,
		EInvalidarg,
		ENointerface,
		EPointer,
		EHandle,
		EAbort,
		EFail,
		EAccessdenied,
		EPending,
		CoEInitTls,
		CoEInitSharedAllocator,
		CoEInitMemoryAllocator,
		CoEInitClassCache,
		CoEInitRpcChannel,
		CoEInitTlsSetChannelControl,
		CoEInitTlsChannelControl,
		CoEInitUnacceptedUserAllocator,
		CoEInitScmMutexExists,
		CoEInitScmFileMappingExists,
		CoEInitScmMapViewOfFile,
		CoEInitScmExecFailure,
		CoEInitOnlySingleThreaded,
		CoECantRemote,
		CoEBadServerName,
		CoEWrongServerIdentity,
		CoEOle1DdeDisabled,
		CoERunasSyntax,
		CoECreateprocessFailure,
		CoERunasCreateprocessFailure,
		CoERunasLogonFailure,
		CoELaunchPermssionDenied,
		CoEStartServiceFailure,
		CoERemoteCommunicationFailure,
		CoEServerStartTimeout,
		CoEClsregInconsistent,
		CoEIidregInconsistent,
		CoENotSupported,
		CoEReloadDll,
		CoEMsiError,
		CoEAttemptToCreateOutsideClientContext,
		CoEServerPaused,
		CoEServerNotPaused,
		CoEClassDisabled,
		CoEClrnotavailable,
		CoEAsyncWorkRejected,
		CoEServerInitTimeout,
		CoENoSecctxInActivate,
		CoETrackerConfig,
		CoEThreadpoolConfig,
		CoESxsConfig,
		CoEMalformedSpn,
		SOk,
		SFalse,
		OleEFirst,
		OleELast,
		OleSFirst,
		OleSLast,
		OleEOleverb,
		OleEAdvf,
		OleEEnumNomore,
		OleEAdvisenotsupported,
		OleENoconnection,
		OleENotrunning,
		OleENocache,
		OleEBlank,
		OleEClassdiff,
		OleECantGetmoniker,
		OleECantBindtosource,
		OleEStatic,
		OleEPromptsavecancelled,
		OleEInvalidrect,
		OleEWrongcompobj,
		OleEInvalidhwnd,
		OleENotInplaceactive,
		OleECantconvert,
		OleENostorage,
		DvEFormatetc,
		DvEDvtargetdevice,
		DvEStgmedium,
		DvEStatdata,
		DvELindex,
		DvETymed,
		DvEClipformat,
		DvEDvaspect,
		DvEDvtargetdeviceSize,
		DvENoiviewobject,
		DragdropEFirst,
		DragdropELast,
		DragdropSFirst,
		DragdropSLast,
		DragdropENotregistered,
		DragdropEAlreadyregistered,
		DragdropEInvalidhwnd,
		ClassfactoryEFirst,
		ClassfactoryELast,
		ClassfactorySFirst,
		ClassfactorySLast,
		ClassENoaggregation,
		ClassEClassnotavailable,
		ClassENotlicensed,
		MarshalEFirst,
		MarshalELast,
		MarshalSFirst,
		MarshalSLast,
		DataEFirst,
		DataELast,
		DataSFirst,
		DataSLast,
		ViewEFirst,
		ViewELast,
		ViewSFirst,
		ViewSLast,
		ViewEDraw,
		RegdbEFirst,
		RegdbELast,
		RegdbSFirst,
		RegdbSLast,
		RegdbEReadregdb,
		RegdbEWriteregdb,
		RegdbEKeymissing,
		RegdbEInvalidvalue,
		RegdbEClassnotreg,
		RegdbEIidnotreg,
		RegdbEBadthreadingmodel,
		CatEFirst,
		CatELast,
		CatECatidnoexist,
		CatENodescription,
		CsEFirst,
		CsELast,
		CsEPackageNotfound,
		CsENotDeletable,
		CsEClassNotfound,
		CsEInvalidVersion,
		CsENoClassstore,
		CsEObjectNotfound,
		CsEObjectAlreadyExists,
		CsEInvalidPath,
		CsENetworkError,
		CsEAdminLimitExceeded,
		CsESchemaMismatch,
		CsEInternalError,
		CacheEFirst,
		CacheELast,
		CacheSFirst,
		CacheSLast,
		CacheENocacheUpdated,
		OleobjEFirst,
		OleobjELast,
		OleobjSFirst,
		OleobjSLast,
		OleobjENoverbs,
		OleobjEInvalidverb,
		ClientsiteEFirst,
		ClientsiteELast,
		ClientsiteSFirst,
		ClientsiteSLast,
		InplaceENotundoable,
		InplaceENotoolspace,
		InplaceEFirst,
		InplaceELast,
		InplaceSFirst,
		InplaceSLast,
		EnumEFirst,
		EnumELast,
		EnumSFirst,
		EnumSLast,
		Convert10EFirst,
		Convert10ELast,
		Convert10SFirst,
		Convert10SLast,
		Convert10EOlestreamGet,
		Convert10EOlestreamPut,
		Convert10EOlestreamFmt,
		Convert10EOlestreamBitmapToDib,
		Convert10EStgFmt,
		Convert10EStgNoStdStream,
		Convert10EStgDibToBitmap,
		ClipbrdEFirst,
		ClipbrdELast,
		ClipbrdSFirst,
		ClipbrdSLast,
		ClipbrdECantOpen,
		ClipbrdECantEmpty,
		ClipbrdECantSet,
		ClipbrdEBadData,
		ClipbrdECantClose,
		MkEFirst,
		MkELast,
		MkSFirst,
		MkSLast,
		MkEConnectmanually,
		MkEExceededdeadline,
		MkENeedgeneric,
		MkEUnavailable,
		MkESyntax,
		MkENoobject,
		MkEInvalidextension,
		MkEIntermediateinterfacenotsupported,
		MkENotbindable,
		MkENotbound,
		MkECantopenfile,
		MkEMustbotheruser,
		MkENoinverse,
		MkENostorage,
		MkENoprefix,
		MkEEnumerationFailed,
		CoEFirst,
		CoELast,
		CoSFirst,
		CoSLast,
		CoENotinitialized,
		CoEAlreadyinitialized,
		CoECantdetermineclass,
		CoEClassstring,
		CoEIidstring,
		CoEAppnotfound,
		CoEAppsingleuse,
		CoEErrorinapp,
		CoEDllnotfound,
		CoEErrorindll,
		CoEWrongosforapp,
		CoEObjnotreg,
		CoEObjisreg,
		CoEObjnotconnected,
		CoEAppdidntreg,
		CoEReleased,
		EventEFirst,
		EventELast,
		EventSFirst,
		EventSLast,
		EventSSomeSubscribersFailed,
		EventEAllSubscribersFailed,
		EventSNosubscribers,
		EventEQuerysyntax,
		EventEQueryfield,
		EventEInternalexception,
		EventEInternalerror,
		EventEInvalidPerUserSid,
		EventEUserException,
		EventETooManyMethods,
		EventEMissingEventclass,
		EventENotAllRemoved,
		EventEComplusNotInstalled,
		EventECantModifyOrDeleteUnconfiguredObject,
		EventECantModifyOrDeleteConfiguredObject,
		EventEInvalidEventClassPartition,
		EventEPerUserSidNotLoggedOn,
		XactEFirst,
		XactELast,
		XactSFirst,
		XactSLast,
		XactEAlreadyothersinglephase,
		XactECantretain,
		XactECommitfailed,
		XactECommitprevented,
		XactEHeuristicabort,
		XactEHeuristiccommit,
		XactEHeuristicdamage,
		XactEHeuristicdanger,
		XactEIsolationlevel,
		XactENoasync,
		XactENoenlist,
		XactENoisoretain,
		XactENoresource,
		XactENotcurrent,
		XactENotransaction,
		XactENotsupported,
		XactEUnknownrmgrid,
		XactEWrongstate,
		XactEWronguow,
		XactEXtionexists,
		XactENoimportobject,
		XactEInvalidcookie,
		XactEIndoubt,
		XactENotimeout,
		XactEAlreadyinprogress,
		XactEAborted,
		XactELogfull,
		XactETmnotavailable,
		XactEConnectionDown,
		XactEConnectionDenied,
		XactEReenlisttimeout,
		XactETipConnectFailed,
		XactETipProtocolError,
		XactETipPullFailed,
		XactEDestTmnotavailable,
		XactETipDisabled,
		XactENetworkTxDisabled,
		XactEPartnerNetworkTxDisabled,
		XactEXaTxDisabled,
		XactEUnableToReadDtcConfig,
		XactEUnableToLoadDtcProxy,
		XactEAborting,
		XactEPushCommFailure,
		XactEPullCommFailure,
		XactELuTxDisabled,
		XactEClerknotfound,
		XactEClerkexists,
		XactERecoveryinprogress,
		XactETransactionclosed,
		XactEInvalidlsn,
		XactEReplayrequest,
		XactSAsync,
		XactSDefect,
		XactSReadonly,
		XactSSomenoretain,
		XactSOkinform,
		XactSMadechangescontent,
		XactSMadechangesinform,
		XactSAllnoretain,
		XactSAborting,
		XactSSinglephase,
		XactSLocallyOk,
		XactSLastresourcemanager,
		ContextEFirst,
		ContextELast,
		ContextSFirst,
		ContextSLast,
		ContextEAborted,
		ContextEAborting,
		ContextENocontext,
		ContextEWouldDeadlock,
		ContextESynchTimeout,
		ContextEOldref,
		ContextERolenotfound,
		ContextETmnotavailable,
		CoEActivationfailed,
		CoEActivationfailedEventlogged,
		CoEActivationfailedCatalogerror,
		CoEActivationfailedTimeout,
		CoEInitializationfailed,
		ContextENojit,
		ContextENotransaction,
		CoEThreadingmodelChanged,
		CoENoiisintrinsics,
		CoENocookies,
		CoEDberror,
		CoENotpooled,
		CoENotconstructed,
		CoENosynchronization,
		CoEIsolevelmismatch,
		CoECallOutOfTxScopeNotAllowed,
		CoEExitTransactionScopeNotCalled,
		OleSUsereg,
		OleSStatic,
		OleSMacClipformat,
		DragdropSDrop,
		DragdropSCancel,
		DragdropSUsedefaultcursors,
		DataSSameformatetc,
		ViewSAlreadyFrozen,
		CacheSFormatetcNotsupported,
		CacheSSamecache,
		CacheSSomecachesNotupdated,
		OleobjSInvalidverb,
		OleobjSCannotDoverbNow,
		OleobjSInvalidhwnd,
		InplaceSTruncated,
		Convert10SNoPresentation,
		MkSReducedToSelf,
		MkSMe,
		MkSHim,
		MkSUs,
		MkSMonikeralreadyregistered,
		SchedSTaskReady,
		SchedSTaskRunning,
		SchedSTaskDisabled,
		SchedSTaskHasNotRun,
		SchedSTaskNoMoreRuns,
		SchedSTaskNotScheduled,
		SchedSTaskTerminated,
		SchedSTaskNoValidTriggers,
		SchedSEventTrigger,
		SchedETriggerNotFound,
		SchedETaskNotReady,
		SchedETaskNotRunning,
		SchedEServiceNotInstalled,
		SchedECannotOpenTask,
		SchedEInvalidTask,
		SchedEAccountInformationNotSet,
		SchedEAccountNameNotFound,
		SchedEAccountDbaseCorrupt,
		SchedENoSecurityServices,
		SchedEUnknownObjectVersion,
		SchedEUnsupportedAccountOption,
		SchedEServiceNotRunning,
		SchedEUnexpectednode,
		SchedENamespace,
		SchedEInvalidvalue,
		SchedEMissingnode,
		SchedEMalformedxml,
		SchedSSomeTriggersFailed,
		SchedSBatchLogonProblem,
		SchedETooManyNodes,
		SchedEPastEndBoundary,
		SchedEAlreadyRunning,
		SchedEUserNotLoggedOn,
		SchedEInvalidTaskHash,
		SchedEServiceNotAvailable,
		SchedEServiceTooBusy,
		SchedETaskAttempted,
		SchedSTaskQueued,
		SchedETaskDisabled,
		SchedETaskNotV1Compat,
		SchedEStartOnDemand,
		CoEClassCreateFailed,
		CoEScmError,
		CoEScmRpcFailure,
		CoEBadPath,
		CoEServerExecFailure,
		CoEObjsrvRpcFailure,
		MkENoNormalized,
		CoEServerStopping,
		MemEInvalidRoot,
		MemEInvalidLink,
		MemEInvalidSize,
		CoSNotallinterfaces,
		CoSMachinenamenotfound,
		CoEMissingDisplayname,
		CoERunasValueMustBeAaa,
		CoEElevationDisabled,
		DispEUnknowninterface,
		DispEMembernotfound,
		DispEParamnotfound,
		DispETypemismatch,
		DispEUnknownname,
		DispENonamedargs,
		DispEBadvartype,
		DispEException,
		DispEOverflow,
		DispEBadindex,
		DispEUnknownlcid,
		DispEArrayislocked,
		DispEBadparamcount,
		DispEParamnotoptional,
		DispEBadcallee,
		DispENotacollection,
		DispEDivbyzero,
		DispEBuffertoosmall,
		TypeEBuffertoosmall,
		TypeEFieldnotfound,
		TypeEInvdataread,
		TypeEUnsupformat,
		TypeERegistryaccess,
		TypeELibnotregistered,
		TypeEUndefinedtype,
		TypeEQualifiednamedisallowed,
		TypeEInvalidstate,
		TypeEWrongtypekind,
		TypeEElementnotfound,
		TypeEAmbiguousname,
		TypeENameconflict,
		TypeEUnknownlcid,
		TypeEDllfunctionnotfound,
		TypeEBadmodulekind,
		TypeESizetoobig,
		TypeEDuplicateid,
		TypeEInvalidid,
		TypeETypemismatch,
		TypeEOutofbounds,
		TypeEIoerror,
		TypeECantcreatetmpfile,
		TypeECantloadlibrary,
		TypeEInconsistentpropfuncs,
		TypeECirculartype,
		StgEInvalidfunction,
		StgEFilenotfound,
		StgEPathnotfound,
		StgEToomanyopenfiles,
		StgEAccessdenied,
		StgEInvalidhandle,
		StgEInsufficientmemory,
		StgEInvalidpointer,
		StgENomorefiles,
		StgEDiskiswriteprotected,
		StgESeekerror,
		StgEWritefault,
		StgEReadfault,
		StgEShareviolation,
		StgELockviolation,
		StgEFilealreadyexists,
		StgEInvalidparameter,
		StgEMediumfull,
		StgEPropsetmismatched,
		StgEAbnormalapiexit,
		StgEInvalidheader,
		StgEInvalidname,
		StgEUnknown,
		StgEUnimplementedfunction,
		StgEInvalidflag,
		StgEInuse,
		StgENotcurrent,
		StgEReverted,
		StgECantsave,
		StgEOldformat,
		StgEOlddll,
		StgESharerequired,
		StgENotfilebasedstorage,
		StgEExtantmarshallings,
		StgEDocfilecorrupt,
		StgEBadbaseaddress,
		StgEDocfiletoolarge,
		StgENotsimpleformat,
		StgEIncomplete,
		StgETerminated,
		StgSConverted,
		StgSBlock,
		StgSRetrynow,
		StgSMonitoring,
		StgSMultipleopens,
		StgSConsolidationfailed,
		StgSCannotconsolidate,
		StgEStatusCopyProtectionFailure,
		StgECssAuthenticationFailure,
		StgECssKeyNotPresent,
		StgECssKeyNotEstablished,
		StgECssScrambledSector,
		StgECssRegionMismatch,
		StgEResetsExhausted,
		RpcECallRejected,
		RpcECallCanceled,
		RpcECantpostInsendcall,
		RpcECantcalloutInasynccall,
		RpcECantcalloutInexternalcall,
		RpcEConnectionTerminated,
		RpcEServerDied,
		RpcEClientDied,
		RpcEInvalidDatapacket,
		RpcECanttransmitCall,
		RpcEClientCantmarshalData,
		RpcEClientCantunmarshalData,
		RpcEServerCantmarshalData,
		RpcEServerCantunmarshalData,
		RpcEInvalidData,
		RpcEInvalidParameter,
		RpcECantcalloutAgain,
		RpcEServerDiedDne,
		RpcESysCallFailed,
		RpcEOutOfResources,
		RpcEAttemptedMultithread,
		RpcENotRegistered,
		RpcEFault,
		RpcEServerfault,
		RpcEChangedMode,
		RpcEInvalidmethod,
		RpcEDisconnected,
		RpcERetry,
		RpcEServercallRetrylater,
		RpcEServercallRejected,
		RpcEInvalidCalldata,
		RpcECantcalloutIninputsynccall,
		RpcEWrongThread,
		RpcEThreadNotInit,
		RpcEVersionMismatch,
		RpcEInvalidHeader,
		RpcEInvalidExtension,
		RpcEInvalidIpid,
		RpcEInvalidObject,
		RpcSCallpending,
		RpcSWaitontimer,
		RpcECallComplete,
		RpcEUnsecureCall,
		RpcETooLate,
		RpcENoGoodSecurityPackages,
		RpcEAccessDenied,
		RpcERemoteDisabled,
		RpcEInvalidObjref,
		RpcENoContext,
		RpcETimeout,
		RpcENoSync,
		RpcEFullsicRequired,
		RpcEInvalidStdName,
		CoEFailedtoimpersonate,
		CoEFailedtogetsecctx,
		CoEFailedtoopenthreadtoken,
		CoEFailedtogettokeninfo,
		CoETrusteedoesntmatchclient,
		CoEFailedtoqueryclientblanket,
		CoEFailedtosetdacl,
		CoEAccesscheckfailed,
		CoENetaccessapifailed,
		CoEWrongtrusteenamesyntax,
		CoEInvalidsid,
		CoEConversionfailed,
		CoENomatchingsidfound,
		CoELookupaccsidfailed,
		CoENomatchingnamefound,
		CoELookupaccnamefailed,
		CoESetserlhndlfailed,
		CoEFailedtogetwindir,
		CoEPathtoolong,
		CoEFailedtogenuuid,
		CoEFailedtocreatefile,
		CoEFailedtoclosehandle,
		CoEExceedsysacllimit,
		CoEAcesinwrongorder,
		CoEIncompatiblestreamversion,
		CoEFailedtoopenprocesstoken,
		CoEDecodefailed,
		CoEAcnotinitialized,
		CoECancelDisabled,
		RpcEUnexpected,
		ErrorAuditingDisabled,
		ErrorAllSidsFiltered,
		ErrorBizrulesNotEnabled,
		NteBadUid,
		NteBadHash,
		NteBadKey,
		NteBadLen,
		NteBadData,
		NteBadSignature,
		NteBadVer,
		NteBadAlgid,
		NteBadFlags,
		NteBadType,
		NteBadKeyState,
		NteBadHashState,
		NteNoKey,
		NteNoMemory,
		NteExists,
		NtePerm,
		NteNotFound,
		NteDoubleEncrypt,
		NteBadProvider,
		NteBadProvType,
		NteBadPublicKey,
		NteBadKeyset,
		NteProvTypeNotDef,
		NteProvTypeEntryBad,
		NteKeysetNotDef,
		NteKeysetEntryBad,
		NteProvTypeNoMatch,
		NteSignatureFileBad,
		NteProviderDllFail,
		NteProvDllNotFound,
		NteBadKeysetParam,
		NteFail,
		NteSysErr,
		NteSilentContext,
		NteTokenKeysetStorageFull,
		NteTemporaryProfile,
		NteFixedparameter,
		NteInvalidHandle,
		NteInvalidParameter,
		NteBufferTooSmall,
		NteNotSupported,
		NteNoMoreItems,
		NteBuffersOverlap,
		NteDecryptionFailure,
		NteInternalError,
		NteUiRequired,
		NteHmacNotSupported,
		SecEInsufficientMemory,
		SecEInvalidHandle,
		SecEUnsupportedFunction,
		SecETargetUnknown,
		SecEInternalError,
		SecESecpkgNotFound,
		SecENotOwner,
		SecECannotInstall,
		SecEInvalidToken,
		SecECannotPack,
		SecEQopNotSupported,
		SecENoImpersonation,
		SecELogonDenied,
		SecEUnknownCredentials,
		SecENoCredentials,
		SecEMessageAltered,
		SecEOutOfSequence,
		SecENoAuthenticatingAuthority,
		SecIContinueNeeded,
		SecICompleteNeeded,
		SecICompleteAndContinue,
		SecILocalLogon,
		SecEBadPkgid,
		SecEContextExpired,
		SecIContextExpired,
		SecEIncompleteMessage,
		SecEIncompleteCredentials,
		SecEBufferTooSmall,
		SecIIncompleteCredentials,
		SecIRenegotiate,
		SecEWrongPrincipal,
		SecINoLsaContext,
		SecETimeSkew,
		SecEUntrustedRoot,
		SecEIllegalMessage,
		SecECertUnknown,
		SecECertExpired,
		SecEEncryptFailure,
		SecEDecryptFailure,
		SecEAlgorithmMismatch,
		SecESecurityQosFailed,
		SecEUnfinishedContextDeleted,
		SecENoTgtReply,
		SecENoIpAddresses,
		SecEWrongCredentialHandle,
		SecECryptoSystemInvalid,
		SecEMaxReferralsExceeded,
		SecEMustBeKdc,
		SecEStrongCryptoNotSupported,
		SecETooManyPrincipals,
		SecENoPaData,
		SecEPkinitNameMismatch,
		SecESmartcardLogonRequired,
		SecEShutdownInProgress,
		SecEKdcInvalidRequest,
		SecEKdcUnableToRefer,
		SecEKdcUnknownEtype,
		SecEUnsupportedPreauth,
		SecEDelegationRequired,
		SecEBadBindings,
		SecEMultipleAccounts,
		SecENoKerbKey,
		SecECertWrongUsage,
		SecEDowngradeDetected,
		SecESmartcardCertRevoked,
		SecEIssuingCaUntrusted,
		SecERevocationOfflineC,
		SecEPkinitClientFailure,
		SecESmartcardCertExpired,
		SecENoS4UProtSupport,
		SecECrossrealmDelegationFailure,
		SecERevocationOfflineKdc,
		SecEIssuingCaUntrustedKdc,
		SecEKdcCertExpired,
		SecEKdcCertRevoked,
		SecISignatureNeeded,
		SecEInvalidParameter,
		SecEDelegationPolicy,
		SecEPolicyNltmOnly,
		SecINoRenegotiation,
		SecENoContext,
		SecEPku2UCertFailure,
		SecEMutualAuthFailed,
		CryptEMsgError,
		CryptEUnknownAlgo,
		CryptEOidFormat,
		CryptEInvalidMsgType,
		CryptEUnexpectedEncoding,
		CryptEAuthAttrMissing,
		CryptEHashValue,
		CryptEInvalidIndex,
		CryptEAlreadyDecrypted,
		CryptENotDecrypted,
		CryptERecipientNotFound,
		CryptEControlType,
		CryptEIssuerSerialnumber,
		CryptESignerNotFound,
		CryptEAttributesMissing,
		CryptEStreamMsgNotReady,
		CryptEStreamInsufficientData,
		CryptINewProtectionRequired,
		CryptEBadLen,
		CryptEBadEncode,
		CryptEFileError,
		CryptENotFound,
		CryptEExists,
		CryptENoProvider,
		CryptESelfSigned,
		CryptEDeletedPrev,
		CryptENoMatch,
		CryptEUnexpectedMsgType,
		CryptENoKeyProperty,
		CryptENoDecryptCert,
		CryptEBadMsg,
		CryptENoSigner,
		CryptEPendingClose,
		CryptERevoked,
		CryptENoRevocationDll,
		CryptENoRevocationCheck,
		CryptERevocationOffline,
		CryptENotInRevocationDatabase,
		CryptEInvalidNumericString,
		CryptEInvalidPrintableString,
		CryptEInvalidIa5String,
		CryptEInvalidX500String,
		CryptENotCharString,
		CryptEFileresized,
		CryptESecuritySettings,
		CryptENoVerifyUsageDll,
		CryptENoVerifyUsageCheck,
		CryptEVerifyUsageOffline,
		CryptENotInCtl,
		CryptENoTrustedSigner,
		CryptEMissingPubkeyPara,
		CryptEOssError,
		OssMoreBuf,
		OssNegativeUinteger,
		OssPduRange,
		OssMoreInput,
		OssDataError,
		OssBadArg,
		OssBadVersion,
		OssOutMemory,
		OssPduMismatch,
		OssLimited,
		OssBadPtr,
		OssBadTime,
		OssIndefiniteNotSupported,
		OssMemError,
		OssBadTable,
		OssTooLong,
		OssConstraintViolated,
		OssFatalError,
		OssAccessSerializationError,
		OssNullTbl,
		OssNullFcn,
		OssBadEncrules,
		OssUnavailEncrules,
		OssCantOpenTraceWindow,
		OssUnimplemented,
		OssOidDllNotLinked,
		OssCantOpenTraceFile,
		OssTraceFileAlreadyOpen,
		OssTableMismatch,
		OssTypeNotSupported,
		OssRealDllNotLinked,
		OssRealCodeNotLinked,
		OssOutOfRange,
		OssCopierDllNotLinked,
		OssConstraintDllNotLinked,
		OssComparatorDllNotLinked,
		OssComparatorCodeNotLinked,
		OssMemMgrDllNotLinked,
		OssPdvDllNotLinked,
		OssPdvCodeNotLinked,
		OssApiDllNotLinked,
		OssBerderDllNotLinked,
		OssPerDllNotLinked,
		OssOpenTypeError,
		OssMutexNotCreated,
		OssCantCloseTraceFile,
		CryptEAsn1Error,
		CryptEAsn1Internal,
		CryptEAsn1Eod,
		CryptEAsn1Corrupt,
		CryptEAsn1Large,
		CryptEAsn1Constraint,
		CryptEAsn1Memory,
		CryptEAsn1Overflow,
		CryptEAsn1Badpdu,
		CryptEAsn1Badargs,
		CryptEAsn1Badreal,
		CryptEAsn1Badtag,
		CryptEAsn1Choice,
		CryptEAsn1Rule,
		CryptEAsn1Utf8,
		CryptEAsn1PduType,
		CryptEAsn1Nyi,
		CryptEAsn1Extended,
		CryptEAsn1Noeod,
		CertsrvEBadRequestsubject,
		CertsrvENoRequest,
		CertsrvEBadRequeststatus,
		CertsrvEPropertyEmpty,
		CertsrvEInvalidCaCertificate,
		CertsrvEServerSuspended,
		CertsrvEEncodingLength,
		CertsrvERoleconflict,
		CertsrvERestrictedofficer,
		CertsrvEKeyArchivalNotConfigured,
		CertsrvENoValidKra,
		CertsrvEBadRequestKeyArchival,
		CertsrvENoCaadminDefined,
		CertsrvEBadRenewalCertAttribute,
		CertsrvENoDbSessions,
		CertsrvEAlignmentFault,
		CertsrvEEnrollDenied,
		CertsrvETemplateDenied,
		CertsrvEDownlevelDcSslOrUpgrade,
		CertsrvEAdminDeniedRequest,
		CertsrvENoPolicyServer,
		CertsrvEUnsupportedCertType,
		CertsrvENoCertType,
		CertsrvETemplateConflict,
		CertsrvESubjectAltNameRequired,
		CertsrvEArchivedKeyRequired,
		CertsrvESmimeRequired,
		CertsrvEBadRenewalSubject,
		CertsrvEBadTemplateVersion,
		CertsrvETemplatePolicyRequired,
		CertsrvESignaturePolicyRequired,
		CertsrvESignatureCount,
		CertsrvESignatureRejected,
		CertsrvEIssuancePolicyRequired,
		CertsrvESubjectUpnRequired,
		CertsrvESubjectDirectoryGuidRequired,
		CertsrvESubjectDnsRequired,
		CertsrvEArchivedKeyUnexpected,
		CertsrvEKeyLength,
		CertsrvESubjectEmailRequired,
		CertsrvEUnknownCertType,
		CertsrvECertTypeOverlap,
		CertsrvETooManySignatures,
		XenrollEKeyNotExportable,
		XenrollECannotAddRootCert,
		XenrollEResponseKaHashNotFound,
		XenrollEResponseUnexpectedKaHash,
		XenrollEResponseKaHashMismatch,
		XenrollEKeyspecSmimeMismatch,
		TrustESystemError,
		TrustENoSignerCert,
		TrustECounterSigner,
		TrustECertSignature,
		TrustETimeStamp,
		TrustEBadDigest,
		TrustEBasicConstraints,
		TrustEFinancialCriteria,
		MssipotfEOutofmemrange,
		MssipotfECantgetobject,
		MssipotfENoheadtable,
		MssipotfEBadMagicnumber,
		MssipotfEBadOffsetTable,
		MssipotfETableTagorder,
		MssipotfETableLongword,
		MssipotfEBadFirstTablePlacement,
		MssipotfETablesOverlap,
		MssipotfETablePadbytes,
		MssipotfEFiletoosmall,
		MssipotfETableChecksum,
		MssipotfEFileChecksum,
		MssipotfEFailedPolicy,
		MssipotfEFailedHintsCheck,
		MssipotfENotOpentype,
		MssipotfEFile,
		MssipotfECrypt,
		MssipotfEBadversion,
		MssipotfEDsigStructure,
		MssipotfEPconstCheck,
		MssipotfEStructure,
		ErrorCredRequiresConfirmation,
		NteOpOk,
		TrustEProviderUnknown,
		TrustEActionUnknown,
		TrustESubjectFormUnknown,
		TrustESubjectNotTrusted,
		DigsigEEncode,
		DigsigEDecode,
		DigsigEExtensibility,
		DigsigECrypto,
		PersistESizedefinite,
		PersistESizeindefinite,
		PersistENotselfsizing,
		TrustENosignature,
		CertEExpired,
		CertEValidityperiodnesting,
		CertERole,
		CertEPathlenconst,
		CertECritical,
		CertEPurpose,
		CertEIssuerchaining,
		CertEMalformed,
		CertEUntrustedroot,
		CertEChaining,
		TrustEFail,
		CertERevoked,
		CertEUntrustedtestroot,
		CertERevocationFailure,
		CertECnNoMatch,
		CertEWrongUsage,
		TrustEExplicitDistrust,
		CertEUntrustedca,
		CertEInvalidPolicy,
		CertEInvalidName,
		SpapiEExpectedSectionName,
		SpapiEBadSectionNameLine,
		SpapiESectionNameTooLong,
		SpapiEGeneralSyntax,
		SpapiEWrongInfStyle,
		SpapiESectionNotFound,
		SpapiELineNotFound,
		SpapiENoBackup,
		SpapiENoAssociatedClass,
		SpapiEClassMismatch,
		SpapiEDuplicateFound,
		SpapiENoDriverSelected,
		SpapiEKeyDoesNotExist,
		SpapiEInvalidDevinstName,
		SpapiEInvalidClass,
		SpapiEDevinstAlreadyExists,
		SpapiEDevinfoNotRegistered,
		SpapiEInvalidRegProperty,
		SpapiENoInf,
		SpapiENoSuchDevinst,
		SpapiECantLoadClassIcon,
		SpapiEInvalidClassInstaller,
		SpapiEDiDoDefault,
		SpapiEDiNofilecopy,
		SpapiEInvalidHwprofile,
		SpapiENoDeviceSelected,
		SpapiEDevinfoListLocked,
		SpapiEDevinfoDataLocked,
		SpapiEDiBadPath,
		SpapiENoClassinstallParams,
		SpapiEFilequeueLocked,
		SpapiEBadServiceInstallsect,
		SpapiENoClassDriverList,
		SpapiENoAssociatedService,
		SpapiENoDefaultDeviceInterface,
		SpapiEDeviceInterfaceActive,
		SpapiEDeviceInterfaceRemoved,
		SpapiEBadInterfaceInstallsect,
		SpapiENoSuchInterfaceClass,
		SpapiEInvalidReferenceString,
		SpapiEInvalidMachinename,
		SpapiERemoteCommFailure,
		SpapiEMachineUnavailable,
		SpapiENoConfigmgrServices,
		SpapiEInvalidProppageProvider,
		SpapiENoSuchDeviceInterface,
		SpapiEDiPostprocessingRequired,
		SpapiEInvalidCoinstaller,
		SpapiENoCompatDrivers,
		SpapiENoDeviceIcon,
		SpapiEInvalidInfLogconfig,
		SpapiEDiDontInstall,
		SpapiEInvalidFilterDriver,
		SpapiENonWindowsNtDriver,
		SpapiENonWindowsDriver,
		SpapiENoCatalogForOemInf,
		SpapiEDevinstallQueueNonnative,
		SpapiENotDisableable,
		SpapiECantRemoveDevinst,
		SpapiEInvalidTarget,
		SpapiEDriverNonnative,
		SpapiEInWow64,
		SpapiESetSystemRestorePoint,
		SpapiEIncorrectlyCopiedInf,
		SpapiESceDisabled,
		SpapiEUnknownException,
		SpapiEPnpRegistryError,
		SpapiERemoteRequestUnsupported,
		SpapiENotAnInstalledOemInf,
		SpapiEInfInUseByDevices,
		SpapiEDiFunctionObsolete,
		SpapiENoAuthenticodeCatalog,
		SpapiEAuthenticodeDisallowed,
		SpapiEAuthenticodeTrustedPublisher,
		SpapiEAuthenticodeTrustNotEstablished,
		SpapiEAuthenticodePublisherNotTrusted,
		SpapiESignatureOsattributeMismatch,
		SpapiEOnlyValidateViaAuthenticode,
		SpapiEDeviceInstallerNotReady,
		SpapiEDriverStoreAddFailed,
		SpapiEDeviceInstallBlocked,
		SpapiEDriverInstallBlocked,
		SpapiEWrongInfType,
		SpapiEFileHashNotInCatalog,
		SpapiEDriverStoreDeleteFailed,
		SpapiEUnrecoverableStackOverflow,
		SpapiEErrorNotInstalled,
		ScardFInternalError,
		ScardECancelled,
		ScardEInvalidHandle,
		ScardEInvalidParameter,
		ScardEInvalidTarget,
		ScardENoMemory,
		ScardFWaitedTooLong,
		ScardEInsufficientBuffer,
		ScardEUnknownReader,
		ScardETimeout,
		ScardESharingViolation,
		ScardENoSmartcard,
		ScardEUnknownCard,
		ScardECantDispose,
		ScardEProtoMismatch,
		ScardENotReady,
		ScardEInvalidValue,
		ScardESystemCancelled,
		ScardFCommError,
		ScardFUnknownError,
		ScardEInvalidAtr,
		ScardENotTransacted,
		ScardEReaderUnavailable,
		ScardPShutdown,
		ScardEPciTooSmall,
		ScardEReaderUnsupported,
		ScardEDuplicateReader,
		ScardECardUnsupported,
		ScardENoService,
		ScardEServiceStopped,
		ScardEUnexpected,
		ScardEIccInstallation,
		ScardEIccCreateorder,
		ScardEUnsupportedFeature,
		ScardEDirNotFound,
		ScardEFileNotFound,
		ScardENoDir,
		ScardENoFile,
		ScardENoAccess,
		ScardEWriteTooMany,
		ScardEBadSeek,
		ScardEInvalidChv,
		ScardEUnknownResMng,
		ScardENoSuchCertificate,
		ScardECertificateUnavailable,
		ScardENoReadersAvailable,
		ScardECommDataLost,
		ScardENoKeyContainer,
		ScardEServerTooBusy,
		ScardEPinCacheExpired,
		ScardENoPinCache,
		ScardEReadOnlyCard,
		ScardWUnsupportedCard,
		ScardWUnresponsiveCard,
		ScardWUnpoweredCard,
		ScardWResetCard,
		ScardWRemovedCard,
		ScardWSecurityViolation,
		ScardWWrongChv,
		ScardWChvBlocked,
		ScardWEof,
		ScardWCancelledByUser,
		ScardWCardNotAuthenticated,
		ScardWCacheItemNotFound,
		ScardWCacheItemStale,
		ScardWCacheItemTooBig,
		ComadminEObjecterrors,
		ComadminEObjectinvalid,
		ComadminEKeymissing,
		ComadminEAlreadyinstalled,
		ComadminEAppFileWritefail,
		ComadminEAppFileReadfail,
		ComadminEAppFileVersion,
		ComadminEBadpath,
		ComadminEApplicationexists,
		ComadminERoleexists,
		ComadminECantcopyfile,
		ComadminENouser,
		ComadminEInvaliduserids,
		ComadminENoregistryclsid,
		ComadminEBadregistryprogid,
		ComadminEAuthenticationlevel,
		ComadminEUserpasswdnotvalid,
		ComadminEClsidoriidmismatch,
		ComadminERemoteinterface,
		ComadminEDllregisterserver,
		ComadminENoservershare,
		ComadminEDllloadfailed,
		ComadminEBadregistrylibid,
		ComadminEAppdirnotfound,
		ComadminERegistrarfailed,
		ComadminECompfileDoesnotexist,
		ComadminECompfileLoaddllfail,
		ComadminECompfileGetclassobj,
		ComadminECompfileClassnotavail,
		ComadminECompfileBadtlb,
		ComadminECompfileNotinstallable,
		ComadminENotchangeable,
		ComadminENotdeleteable,
		ComadminESession,
		ComadminECompMoveLocked,
		ComadminECompMoveBadDest,
		ComadminERegistertlb,
		ComadminESystemapp,
		ComadminECompfileNoregistrar,
		ComadminECoreqcompinstalled,
		ComadminEServicenotinstalled,
		ComadminEPropertysavefailed,
		ComadminEObjectexists,
		ComadminEComponentexists,
		ComadminERegfileCorrupt,
		ComadminEPropertyOverflow,
		ComadminENotinregistry,
		ComadminEObjectnotpoolable,
		ComadminEApplidMatchesClsid,
		ComadminERoleDoesNotExist,
		ComadminEStartAppNeedsComponents,
		ComadminERequiresDifferentPlatform,
		ComadminECanNotExportAppProxy,
		ComadminECanNotStartApp,
		ComadminECanNotExportSysApp,
		ComadminECantSubscribeToComponent,
		ComadminEEventclassCantBeSubscriber,
		ComadminELibAppProxyIncompatible,
		ComadminEBasePartitionOnly,
		ComadminEStartAppDisabled,
		ComadminECatDuplicatePartitionName,
		ComadminECatInvalidPartitionName,
		ComadminECatPartitionInUse,
		ComadminEFilePartitionDuplicateFiles,
		ComadminECatImportedComponentsNotAllowed,
		ComadminEAmbiguousApplicationName,
		ComadminEAmbiguousPartitionName,
		ComadminERegdbNotinitialized,
		ComadminERegdbNotopen,
		ComadminERegdbSystemerr,
		ComadminERegdbAlreadyrunning,
		ComadminEMigVersionnotsupported,
		ComadminEMigSchemanotfound,
		ComadminECatBitnessmismatch,
		ComadminECatUnacceptablebitness,
		ComadminECatWrongappbitness,
		ComadminECatPauseResumeNotSupported,
		ComadminECatServerfault,
		ComqcEApplicationNotQueued,
		ComqcENoQueueableInterfaces,
		ComqcEQueuingServiceNotAvailable,
		ComqcENoIpersiststream,
		ComqcEBadMessage,
		ComqcEUnauthenticated,
		ComqcEUntrustedEnqueuer,
		MsdtcEDuplicateResource,
		ComadminEObjectParentMissing,
		ComadminEObjectDoesNotExist,
		ComadminEAppNotRunning,
		ComadminEInvalidPartition,
		ComadminESvcappNotPoolableOrRecyclable,
		ComadminEUserInSet,
		ComadminECantrecyclelibraryapps,
		ComadminECantrecycleserviceapps,
		ComadminEProcessalreadyrecycled,
		ComadminEPausedprocessmaynotberecycled,
		ComadminECantmakeinprocservice,
		ComadminEProgidinusebyclsid,
		ComadminEDefaultPartitionNotInSet,
		ComadminERecycledprocessmaynotbepaused,
		ComadminEPartitionAccessdenied,
		ComadminEPartitionMsiOnly,
		ComadminELegacycompsNotAllowedIn10Format,
		ComadminELegacycompsNotAllowedInNonbasePartitions,
		ComadminECompMoveSource,
		ComadminECompMoveDest,
		ComadminECompMovePrivate,
		ComadminEBasepartitionRequiredInSet,
		ComadminECannotAliasEventclass,
		ComadminEPrivateAccessdenied,
		ComadminESaferinvalid,
		ComadminERegistryAccessdenied,
		ComadminEPartitionsDisabled,
		ErrorFltIoComplete,
		ErrorFltNoHandlerDefined,
		ErrorFltContextAlreadyDefined,
		ErrorFltInvalidAsynchronousRequest,
		ErrorFltDisallowFastIo,
		ErrorFltInvalidNameRequest,
		ErrorFltNotSafeToPostOperation,
		ErrorFltNotInitialized,
		ErrorFltFilterNotReady,
		ErrorFltPostOperationCleanup,
		ErrorFltInternalError,
		ErrorFltDeletingObject,
		ErrorFltMustBeNonpagedPool,
		ErrorFltDuplicateEntry,
		ErrorFltCbdqDisabled,
		ErrorFltDoNotAttach,
		ErrorFltDoNotDetach,
		ErrorFltInstanceAltitudeCollision,
		ErrorFltInstanceNameCollision,
		ErrorFltFilterNotFound,
		ErrorFltVolumeNotFound,
		ErrorFltInstanceNotFound,
		ErrorFltContextAllocationNotFound,
		ErrorFltInvalidContextRegistration,
		ErrorFltNameCacheMiss,
		ErrorFltNoDeviceObject,
		ErrorFltVolumeAlreadyMounted,
		ErrorFltAlreadyEnlisted,
		ErrorFltContextAlreadyLinked,
		ErrorFltNoWaiterForReply,
		ErrorHungDisplayDriverThread,
		DwmECompositiondisabled,
		DwmERemotingNotSupported,
		DwmENoRedirectionSurfaceAvailable,
		DwmENotQueuingPresents,
		DwmEAdapterNotFound,
		DwmSGdiRedirectionSurface,
		ErrorMonitorNoDescriptor,
		ErrorMonitorUnknownDescriptorFormat,
		ErrorMonitorInvalidDescriptorChecksum,
		ErrorMonitorInvalidStandardTimingBlock,
		ErrorMonitorWmiDatablockRegistrationFailed,
		ErrorMonitorInvalidSerialNumberMondscBlock,
		ErrorMonitorInvalidUserFriendlyMondscBlock,
		ErrorMonitorNoMoreDescriptorData,
		ErrorMonitorInvalidDetailedTimingBlock,
		ErrorMonitorInvalidManufactureDate,
		ErrorGraphicsNotExclusiveModeOwner,
		ErrorGraphicsInsufficientDmaBuffer,
		ErrorGraphicsInvalidDisplayAdapter,
		ErrorGraphicsAdapterWasReset,
		ErrorGraphicsInvalidDriverModel,
		ErrorGraphicsPresentModeChanged,
		ErrorGraphicsPresentOccluded,
		ErrorGraphicsPresentDenied,
		ErrorGraphicsCannotcolorconvert,
		ErrorGraphicsDriverMismatch,
		ErrorGraphicsPartialDataPopulated,
		ErrorGraphicsPresentRedirectionDisabled,
		ErrorGraphicsPresentUnoccluded,
		ErrorGraphicsNoVideoMemory,
		ErrorGraphicsCantLockMemory,
		ErrorGraphicsAllocationBusy,
		ErrorGraphicsTooManyReferences,
		ErrorGraphicsTryAgainLater,
		ErrorGraphicsTryAgainNow,
		ErrorGraphicsAllocationInvalid,
		ErrorGraphicsUnswizzlingApertureUnavailable,
		ErrorGraphicsUnswizzlingApertureUnsupported,
		ErrorGraphicsCantEvictPinnedAllocation,
		ErrorGraphicsInvalidAllocationUsage,
		ErrorGraphicsCantRenderLockedAllocation,
		ErrorGraphicsAllocationClosed,
		ErrorGraphicsInvalidAllocationInstance,
		ErrorGraphicsInvalidAllocationHandle,
		ErrorGraphicsWrongAllocationDevice,
		ErrorGraphicsAllocationContentLost,
		ErrorGraphicsGpuExceptionOnDevice,
		ErrorGraphicsInvalidVidpnTopology,
		ErrorGraphicsVidpnTopologyNotSupported,
		ErrorGraphicsVidpnTopologyCurrentlyNotSupported,
		ErrorGraphicsInvalidVidpn,
		ErrorGraphicsInvalidVideoPresentSource,
		ErrorGraphicsInvalidVideoPresentTarget,
		ErrorGraphicsVidpnModalityNotSupported,
		ErrorGraphicsModeNotPinned,
		ErrorGraphicsInvalidVidpnSourcemodeset,
		ErrorGraphicsInvalidVidpnTargetmodeset,
		ErrorGraphicsInvalidFrequency,
		ErrorGraphicsInvalidActiveRegion,
		ErrorGraphicsInvalidTotalRegion,
		ErrorGraphicsInvalidVideoPresentSourceMode,
		ErrorGraphicsInvalidVideoPresentTargetMode,
		ErrorGraphicsPinnedModeMustRemainInSet,
		ErrorGraphicsPathAlreadyInTopology,
		ErrorGraphicsModeAlreadyInModeset,
		ErrorGraphicsInvalidVideopresentsourceset,
		ErrorGraphicsInvalidVideopresenttargetset,
		ErrorGraphicsSourceAlreadyInSet,
		ErrorGraphicsTargetAlreadyInSet,
		ErrorGraphicsInvalidVidpnPresentPath,
		ErrorGraphicsNoRecommendedVidpnTopology,
		ErrorGraphicsInvalidMonitorFrequencyrangeset,
		ErrorGraphicsInvalidMonitorFrequencyrange,
		ErrorGraphicsFrequencyrangeNotInSet,
		ErrorGraphicsNoPreferredMode,
		ErrorGraphicsFrequencyrangeAlreadyInSet,
		ErrorGraphicsStaleModeset,
		ErrorGraphicsInvalidMonitorSourcemodeset,
		ErrorGraphicsInvalidMonitorSourceMode,
		ErrorGraphicsNoRecommendedFunctionalVidpn,
		ErrorGraphicsModeIdMustBeUnique,
		ErrorGraphicsEmptyAdapterMonitorModeSupportIntersection,
		ErrorGraphicsVideoPresentTargetsLessThanSources,
		ErrorGraphicsPathNotInTopology,
		ErrorGraphicsAdapterMustHaveAtLeastOneSource,
		ErrorGraphicsAdapterMustHaveAtLeastOneTarget,
		ErrorGraphicsInvalidMonitordescriptorset,
		ErrorGraphicsInvalidMonitordescriptor,
		ErrorGraphicsMonitordescriptorNotInSet,
		ErrorGraphicsMonitordescriptorAlreadyInSet,
		ErrorGraphicsMonitordescriptorIdMustBeUnique,
		ErrorGraphicsInvalidVidpnTargetSubsetType,
		ErrorGraphicsResourcesNotRelated,
		ErrorGraphicsSourceIdMustBeUnique,
		ErrorGraphicsTargetIdMustBeUnique,
		ErrorGraphicsNoAvailableVidpnTarget,
		ErrorGraphicsMonitorCouldNotBeAssociatedWithAdapter,
		ErrorGraphicsNoVidpnmgr,
		ErrorGraphicsNoActiveVidpn,
		ErrorGraphicsStaleVidpnTopology,
		ErrorGraphicsMonitorNotConnected,
		ErrorGraphicsSourceNotInTopology,
		ErrorGraphicsInvalidPrimarysurfaceSize,
		ErrorGraphicsInvalidVisibleregionSize,
		ErrorGraphicsInvalidStride,
		ErrorGraphicsInvalidPixelformat,
		ErrorGraphicsInvalidColorbasis,
		ErrorGraphicsInvalidPixelvalueaccessmode,
		ErrorGraphicsTargetNotInTopology,
		ErrorGraphicsNoDisplayModeManagementSupport,
		ErrorGraphicsVidpnSourceInUse,
		ErrorGraphicsCantAccessActiveVidpn,
		ErrorGraphicsInvalidPathImportanceOrdinal,
		ErrorGraphicsInvalidPathContentGeometryTransformation,
		ErrorGraphicsPathContentGeometryTransformationNotSupported,
		ErrorGraphicsInvalidGammaRamp,
		ErrorGraphicsGammaRampNotSupported,
		ErrorGraphicsMultisamplingNotSupported,
		ErrorGraphicsModeNotInModeset,
		ErrorGraphicsDatasetIsEmpty,
		ErrorGraphicsNoMoreElementsInDataset,
		ErrorGraphicsInvalidVidpnTopologyRecommendationReason,
		ErrorGraphicsInvalidPathContentType,
		ErrorGraphicsInvalidCopyprotectionType,
		ErrorGraphicsUnassignedModesetAlreadyExists,
		ErrorGraphicsPathContentGeometryTransformationNotPinned,
		ErrorGraphicsInvalidScanlineOrdering,
		ErrorGraphicsTopologyChangesNotAllowed,
		ErrorGraphicsNoAvailableImportanceOrdinals,
		ErrorGraphicsIncompatiblePrivateFormat,
		ErrorGraphicsInvalidModePruningAlgorithm,
		ErrorGraphicsInvalidMonitorCapabilityOrigin,
		ErrorGraphicsInvalidMonitorFrequencyrangeConstraint,
		ErrorGraphicsMaxNumPathsReached,
		ErrorGraphicsCancelVidpnTopologyAugmentation,
		ErrorGraphicsInvalidClientType,
		ErrorGraphicsClientvidpnNotSet,
		ErrorGraphicsSpecifiedChildAlreadyConnected,
		ErrorGraphicsChildDescriptorNotSupported,
		ErrorGraphicsUnknownChildStatus,
		ErrorGraphicsNotALinkedAdapter,
		ErrorGraphicsLeadlinkNotEnumerated,
		ErrorGraphicsChainlinksNotEnumerated,
		ErrorGraphicsAdapterChainNotReady,
		ErrorGraphicsChainlinksNotStarted,
		ErrorGraphicsChainlinksNotPoweredOn,
		ErrorGraphicsInconsistentDeviceLinkState,
		ErrorGraphicsLeadlinkStartDeferred,
		ErrorGraphicsNotPostDeviceDriver,
		ErrorGraphicsPollingTooFrequently,
		ErrorGraphicsStartDeferred,
		ErrorGraphicsAdapterAccessNotExcluded,
		ErrorGraphicsOpmNotSupported,
		ErrorGraphicsCoppNotSupported,
		ErrorGraphicsUabNotSupported,
		ErrorGraphicsOpmInvalidEncryptedParameters,
		ErrorGraphicsOpmNoVideoOutputsExist,
		ErrorGraphicsOpmInternalError,
		ErrorGraphicsOpmInvalidHandle,
		ErrorGraphicsPvpInvalidCertificateLength,
		ErrorGraphicsOpmSpanningModeEnabled,
		ErrorGraphicsOpmTheaterModeEnabled,
		ErrorGraphicsPvpHfsFailed,
		ErrorGraphicsOpmInvalidSrm,
		ErrorGraphicsOpmOutputDoesNotSupportHdcp,
		ErrorGraphicsOpmOutputDoesNotSupportAcp,
		ErrorGraphicsOpmOutputDoesNotSupportCgmsa,
		ErrorGraphicsOpmHdcpSrmNeverSet,
		ErrorGraphicsOpmResolutionTooHigh,
		ErrorGraphicsOpmAllHdcpHardwareAlreadyInUse,
		ErrorGraphicsOpmVideoOutputNoLongerExists,
		ErrorGraphicsOpmSessionTypeChangeInProgress,
		ErrorGraphicsOpmVideoOutputDoesNotHaveCoppSemantics,
		ErrorGraphicsOpmInvalidInformationRequest,
		ErrorGraphicsOpmDriverInternalError,
		ErrorGraphicsOpmVideoOutputDoesNotHaveOpmSemantics,
		ErrorGraphicsOpmSignalingNotSupported,
		ErrorGraphicsOpmInvalidConfigurationRequest,
		ErrorGraphicsI2CNotSupported,
		ErrorGraphicsI2CDeviceDoesNotExist,
		ErrorGraphicsI2CErrorTransmittingData,
		ErrorGraphicsI2CErrorReceivingData,
		ErrorGraphicsDdcciVcpNotSupported,
		ErrorGraphicsDdcciInvalidData,
		ErrorGraphicsDdcciMonitorReturnedInvalidTimingStatusByte,
		ErrorGraphicsMcaInvalidCapabilitiesString,
		ErrorGraphicsMcaInternalError,
		ErrorGraphicsDdcciInvalidMessageCommand,
		ErrorGraphicsDdcciInvalidMessageLength,
		ErrorGraphicsDdcciInvalidMessageChecksum,
		ErrorGraphicsInvalidPhysicalMonitorHandle,
		ErrorGraphicsMonitorNoLongerExists,
		ErrorGraphicsDdcciCurrentCurrentValueGreaterThanMaximumValue,
		ErrorGraphicsMcaInvalidVcpVersion,
		ErrorGraphicsMcaMonitorViolatesMccsSpecification,
		ErrorGraphicsMcaMccsVersionMismatch,
		ErrorGraphicsMcaUnsupportedMccsVersion,
		ErrorGraphicsMcaInvalidTechnologyTypeReturned,
		ErrorGraphicsMcaUnsupportedColorTemperature,
		ErrorGraphicsOnlyConsoleSessionSupported,
		ErrorGraphicsNoDisplayDeviceCorrespondsToName,
		ErrorGraphicsDisplayDeviceNotAttachedToDesktop,
		ErrorGraphicsMirroringDevicesNotSupported,
		ErrorGraphicsInvalidPointer,
		ErrorGraphicsNoMonitorsCorrespondToDisplayDevice,
		ErrorGraphicsParameterArrayTooSmall,
		ErrorGraphicsInternalError,
		ErrorGraphicsSessionTypeChangeInProgress,
		TpmEErrorMask,
		TpmEAuthfail,
		TpmEBadindex,
		TpmEBadParameter,
		TpmEAuditfailure,
		TpmEClearDisabled,
		TpmEDeactivated,
		TpmEDisabled,
		TpmEDisabledCmd,
		TpmEFail,
		TpmEBadOrdinal,
		TpmEInstallDisabled,
		TpmEInvalidKeyhandle,
		TpmEKeynotfound,
		TpmEInappropriateEnc,
		TpmEMigratefail,
		TpmEInvalidPcrInfo,
		TpmENospace,
		TpmENosrk,
		TpmENotsealedBlob,
		TpmEOwnerSet,
		TpmEResources,
		TpmEShortrandom,
		TpmESize,
		TpmEWrongpcrval,
		TpmEBadParamSize,
		TpmEShaThread,
		TpmEShaError,
		TpmEFailedselftest,
		TpmEAuth2Fail,
		TpmEBadtag,
		TpmEIoerror,
		TpmEEncryptError,
		TpmEDecryptError,
		TpmEInvalidAuthhandle,
		TpmENoEndorsement,
		TpmEInvalidKeyusage,
		TpmEWrongEntitytype,
		TpmEInvalidPostinit,
		TpmEInappropriateSig,
		TpmEBadKeyProperty,
		TpmEBadMigration,
		TpmEBadScheme,
		TpmEBadDatasize,
		TpmEBadMode,
		TpmEBadPresence,
		TpmEBadVersion,
		TpmENoWrapTransport,
		TpmEAuditfailUnsuccessful,
		TpmEAuditfailSuccessful,
		TpmENotresetable,
		TpmENotlocal,
		TpmEBadType,
		TpmEInvalidResource,
		TpmENotfips,
		TpmEInvalidFamily,
		TpmENoNvPermission,
		TpmERequiresSign,
		TpmEKeyNotsupported,
		TpmEAuthConflict,
		TpmEAreaLocked,
		TpmEBadLocality,
		TpmEReadOnly,
		TpmEPerNowrite,
		TpmEFamilycount,
		TpmEWriteLocked,
		TpmEBadAttributes,
		TpmEInvalidStructure,
		TpmEKeyOwnerControl,
		TpmEBadCounter,
		TpmENotFullwrite,
		TpmEContextGap,
		TpmEMaxnvwrites,
		TpmENooperator,
		TpmEResourcemissing,
		TpmEDelegateLock,
		TpmEDelegateFamily,
		TpmEDelegateAdmin,
		TpmETransportNotexclusive,
		TpmEOwnerControl,
		TpmEDaaResources,
		TpmEDaaInputData0,
		TpmEDaaInputData1,
		TpmEDaaIssuerSettings,
		TpmEDaaTpmSettings,
		TpmEDaaStage,
		TpmEDaaIssuerValidity,
		TpmEDaaWrongW,
		TpmEBadHandle,
		TpmEBadDelegate,
		TpmEBadcontext,
		TpmEToomanycontexts,
		TpmEMaTicketSignature,
		TpmEMaDestination,
		TpmEMaSource,
		TpmEMaAuthority,
		TpmEPermanentek,
		TpmEBadSignature,
		TpmENocontextspace,
		TpmECommandBlocked,
		TpmEInvalidHandle,
		TpmEDuplicateVhandle,
		TpmEEmbeddedCommandBlocked,
		TpmEEmbeddedCommandUnsupported,
		TpmERetry,
		TpmENeedsSelftest,
		TpmEDoingSelftest,
		TpmEDefendLockRunning,
		TbsEInternalError,
		TbsEBadParameter,
		TbsEInvalidOutputPointer,
		TbsEInvalidContext,
		TbsEInsufficientBuffer,
		TbsEIoerror,
		TbsEInvalidContextParam,
		TbsEServiceNotRunning,
		TbsETooManyTbsContexts,
		TbsETooManyResources,
		TbsEServiceStartPending,
		TbsEPpiNotSupported,
		TbsECommandCanceled,
		TbsEBufferTooLarge,
		TbsETpmNotFound,
		TbsEServiceDisabled,
		TbsENoEventLog,
		TpmapiEInvalidState,
		TpmapiENotEnoughData,
		TpmapiETooMuchData,
		TpmapiEInvalidOutputPointer,
		TpmapiEInvalidParameter,
		TpmapiEOutOfMemory,
		TpmapiEBufferTooSmall,
		TpmapiEInternalError,
		TpmapiEAccessDenied,
		TpmapiEAuthorizationFailed,
		TpmapiEInvalidContextHandle,
		TpmapiETbsCommunicationError,
		TpmapiETpmCommandError,
		TpmapiEMessageTooLarge,
		TpmapiEInvalidEncoding,
		TpmapiEInvalidKeySize,
		TpmapiEEncryptionFailed,
		TpmapiEInvalidKeyParams,
		TpmapiEInvalidMigrationAuthorizationBlob,
		TpmapiEInvalidPcrIndex,
		TpmapiEInvalidDelegateBlob,
		TpmapiEInvalidContextParams,
		TpmapiEInvalidKeyBlob,
		TpmapiEInvalidPcrData,
		TpmapiEInvalidOwnerAuth,
		TpmapiEFipsRngCheckFailed,
		TpmapiEEmptyTcgLog,
		TpmapiEInvalidTcgLogEntry,
		TpmapiETcgSeparatorAbsent,
		TpmapiETcgInvalidDigestEntry,
		TbsimpEBufferTooSmall,
		TbsimpECleanupFailed,
		TbsimpEInvalidContextHandle,
		TbsimpEInvalidContextParam,
		TbsimpETpmError,
		TbsimpEHashBadKey,
		TbsimpEDuplicateVhandle,
		TbsimpEInvalidOutputPointer,
		TbsimpEInvalidParameter,
		TbsimpERpcInitFailed,
		TbsimpESchedulerNotRunning,
		TbsimpECommandCanceled,
		TbsimpEOutOfMemory,
		TbsimpEListNoMoreItems,
		TbsimpEListNotFound,
		TbsimpENotEnoughSpace,
		TbsimpENotEnoughTpmContexts,
		TbsimpECommandFailed,
		TbsimpEUnknownOrdinal,
		TbsimpEResourceExpired,
		TbsimpEInvalidResource,
		TbsimpENothingToUnload,
		TbsimpEHashTableFull,
		TbsimpETooManyTbsContexts,
		TbsimpETooManyResources,
		TbsimpEPpiNotSupported,
		TbsimpETpmIncompatible,
		TbsimpENoEventLog,
		TpmEPpiAcpiFailure,
		TpmEPpiUserAbort,
		TpmEPpiBiosFailure,
		TpmEPpiNotSupported,
		PlaEDcsNotFound,
		PlaEDcsInUse,
		PlaETooManyFolders,
		PlaENoMinDisk,
		PlaEDcsAlreadyExists,
		PlaSPropertyIgnored,
		PlaEPropertyConflict,
		PlaEDcsSingletonRequired,
		PlaECredentialsRequired,
		PlaEDcsNotRunning,
		PlaEConflictInclExclApi,
		PlaENetworkExeNotValid,
		PlaEExeAlreadyConfigured,
		PlaEExePathNotValid,
		PlaEDcAlreadyExists,
		PlaEDcsStartWaitTimeout,
		PlaEDcStartWaitTimeout,
		PlaEReportWaitTimeout,
		PlaENoDuplicates,
		PlaEExeFullPathRequired,
		PlaEInvalidSessionName,
		PlaEPlaChannelNotEnabled,
		PlaETaskschedChannelNotEnabled,
		PlaERulesManagerFailed,
		PlaECabapiFailure,
		FveELockedVolume,
		FveENotEncrypted,
		FveENoTpmBios,
		FveENoMbrMetric,
		FveENoBootsectorMetric,
		FveENoBootmgrMetric,
		FveEWrongBootmgr,
		FveESecureKeyRequired,
		FveENotActivated,
		FveEActionNotAllowed,
		FveEAdSchemaNotInstalled,
		FveEAdInvalidDatatype,
		FveEAdInvalidDatasize,
		FveEAdNoValues,
		FveEAdAttrNotSet,
		FveEAdGuidNotFound,
		FveEBadInformation,
		FveETooSmall,
		FveESystemVolume,
		FveEFailedWrongFs,
		FveEBadPartitionSize,
		FveENotSupported,
		FveEBadData,
		FveEVolumeNotBound,
		FveETpmNotOwned,
		FveENotDataVolume,
		FveEAdInsufficientBuffer,
		FveEConvRead,
		FveEConvWrite,
		FveEKeyRequired,
		FveEClusteringNotSupported,
		FveEVolumeBoundAlready,
		FveEOsNotProtected,
		FveEProtectionDisabled,
		FveERecoveryKeyRequired,
		FveEForeignVolume,
		FveEOverlappedUpdate,
		FveETpmSrkAuthNotZero,
		FveEFailedSectorSize,
		FveEFailedAuthentication,
		FveENotOsVolume,
		FveEAutounlockEnabled,
		FveEWrongBootsector,
		FveEWrongSystemFs,
		FveEPolicyPasswordRequired,
		FveECannotSetFvekEncrypted,
		FveECannotEncryptNoKey,
		FveEBootableCddvd,
		FveEProtectorExists,
		FveERelativePath,
		FveEProtectorNotFound,
		FveEInvalidKeyFormat,
		FveEInvalidPasswordFormat,
		FveEFipsRngCheckFailed,
		FveEFipsPreventsRecoveryPassword,
		FveEFipsPreventsExternalKeyExport,
		FveENotDecrypted,
		FveEInvalidProtectorType,
		FveENoProtectorsToTest,
		FveEKeyfileNotFound,
		FveEKeyfileInvalid,
		FveEKeyfileNoVmk,
		FveETpmDisabled,
		FveENotAllowedInSafeMode,
		FveETpmInvalidPcr,
		FveETpmNoVmk,
		FveEPinInvalid,
		FveEAuthInvalidApplication,
		FveEAuthInvalidConfig,
		FveEFipsDisableProtectionNotAllowed,
		FveEFsNotExtended,
		FveEFirmwareTypeNotSupported,
		FveENoLicense,
		FveENotOnStack,
		FveEFsMounted,
		FveETokenNotImpersonated,
		FveEDryRunFailed,
		FveERebootRequired,
		FveEDebuggerEnabled,
		FveERawAccess,
		FveERawBlocked,
		FveEBcdApplicationsPathIncorrect,
		FveENotAllowedInVersion,
		FveENoAutounlockMasterKey,
		FveEMorFailed,
		FveEHiddenVolume,
		FveETransientState,
		FveEPubkeyNotAllowed,
		FveEVolumeHandleOpen,
		FveENoFeatureLicense,
		FveEInvalidStartupOptions,
		FveEPolicyRecoveryPasswordNotAllowed,
		FveEPolicyRecoveryPasswordRequired,
		FveEPolicyRecoveryKeyNotAllowed,
		FveEPolicyRecoveryKeyRequired,
		FveEPolicyStartupPinNotAllowed,
		FveEPolicyStartupPinRequired,
		FveEPolicyStartupKeyNotAllowed,
		FveEPolicyStartupKeyRequired,
		FveEPolicyStartupPinKeyNotAllowed,
		FveEPolicyStartupPinKeyRequired,
		FveEPolicyStartupTpmNotAllowed,
		FveEPolicyStartupTpmRequired,
		FveEPolicyInvalidPinLength,
		FveEKeyProtectorNotSupported,
		FveEPolicyPassphraseNotAllowed,
		FveEPolicyPassphraseRequired,
		FveEFipsPreventsPassphrase,
		FveEOsVolumePassphraseNotAllowed,
		FveEInvalidBitlockerOid,
		FveEVolumeTooSmall,
		FveEDvNotSupportedOnFs,
		FveEDvNotAllowedByGp,
		FveEPolicyUserCertificateNotAllowed,
		FveEPolicyUserCertificateRequired,
		FveEPolicyUserCertMustBeHw,
		FveEPolicyUserConfigureFdvAutounlockNotAllowed,
		FveEPolicyUserConfigureRdvAutounlockNotAllowed,
		FveEPolicyUserConfigureRdvNotAllowed,
		FveEPolicyUserEnableRdvNotAllowed,
		FveEPolicyUserDisableRdvNotAllowed,
		FveEPolicyInvalidPassphraseLength,
		FveEPolicyPassphraseTooSimple,
		FveERecoveryPartition,
		FveEPolicyConflictFdvRkOffAukOn,
		FveEPolicyConflictRdvRkOffAukOn,
		FveENonBitlockerOid,
		FveEPolicyProhibitsSelfsigned,
		FveEPolicyConflictRoAndStartupKeyRequired,
		FveEConvRecoveryFailed,
		FveEVirtualizedSpaceTooBig,
		FveEPolicyConflictOsvRpOffAdbOn,
		FveEPolicyConflictFdvRpOffAdbOn,
		FveEPolicyConflictRdvRpOffAdbOn,
		FveENonBitlockerKu,
		FveEPrivatekeyAuthFailed,
		FveERemovalOfDraFailed,
		FveEOperationNotSupportedOnVistaVolume,
		FveECantLockAutounlockEnabledVolume,
		FveEFipsHashKdfNotAllowed,
		FveEEnhPinInvalid,
		FveEInvalidPinChars,
		FveEInvalidDatumType,
		FwpECalloutNotFound,
		FwpEConditionNotFound,
		FwpEFilterNotFound,
		FwpELayerNotFound,
		FwpEProviderNotFound,
		FwpEProviderContextNotFound,
		FwpESublayerNotFound,
		FwpENotFound,
		FwpEAlreadyExists,
		FwpEInUse,
		FwpEDynamicSessionInProgress,
		FwpEWrongSession,
		FwpENoTxnInProgress,
		FwpETxnInProgress,
		FwpETxnAborted,
		FwpESessionAborted,
		FwpEIncompatibleTxn,
		FwpETimeout,
		FwpENetEventsDisabled,
		FwpEIncompatibleLayer,
		FwpEKmClientsOnly,
		FwpELifetimeMismatch,
		FwpEBuiltinObject,
		FwpETooManyCallouts,
		FwpENotificationDropped,
		FwpETrafficMismatch,
		FwpEIncompatibleSaState,
		FwpENullPointer,
		FwpEInvalidEnumerator,
		FwpEInvalidFlags,
		FwpEInvalidNetMask,
		FwpEInvalidRange,
		FwpEInvalidInterval,
		FwpEZeroLengthArray,
		FwpENullDisplayName,
		FwpEInvalidActionType,
		FwpEInvalidWeight,
		FwpEMatchTypeMismatch,
		FwpETypeMismatch,
		FwpEOutOfBounds,
		FwpEReserved,
		FwpEDuplicateCondition,
		FwpEDuplicateKeymod,
		FwpEActionIncompatibleWithLayer,
		FwpEActionIncompatibleWithSublayer,
		FwpEContextIncompatibleWithLayer,
		FwpEContextIncompatibleWithCallout,
		FwpEIncompatibleAuthMethod,
		FwpEIncompatibleDhGroup,
		FwpEEmNotSupported,
		FwpENeverMatch,
		FwpEProviderContextMismatch,
		FwpEInvalidParameter,
		FwpETooManySublayers,
		FwpECalloutNotificationFailed,
		FwpEInvalidAuthTransform,
		FwpEInvalidCipherTransform,
		FwpEDropNoicmp,
		FwpEIncompatibleCipherTransform,
		FwpEInvalidTransformCombination,
		FwpEDuplicateAuthMethod,
		WsSAsync,
		WsSEnd,
		WsEInvalidFormat,
		WsEObjectFaulted,
		WsENumericOverflow,
		WsEInvalidOperation,
		WsEOperationAborted,
		WsEEndpointAccessDenied,
		WsEOperationTimedOut,
		WsEOperationAbandoned,
		WsEQuotaExceeded,
		WsENoTranslationAvailable,
		WsESecurityVerificationFailure,
		WsEAddressInUse,
		WsEAddressNotAvailable,
		WsEEndpointNotFound,
		WsEEndpointNotAvailable,
		WsEEndpointFailure,
		WsEEndpointUnreachable,
		WsEEndpointActionNotSupported,
		WsEEndpointTooBusy,
		WsEEndpointFaultReceived,
		WsEEndpointDisconnected,
		WsEProxyFailure,
		WsEProxyAccessDenied,
		WsENotSupported,
		WsEProxyRequiresBasicAuth,
		WsEProxyRequiresDigestAuth,
		WsEProxyRequiresNtlmAuth,
		WsEProxyRequiresNegotiateAuth,
		WsEServerRequiresBasicAuth,
		WsEServerRequiresDigestAuth,
		WsEServerRequiresNtlmAuth,
		WsEServerRequiresNegotiateAuth,
		WsEInvalidEndpointUrl,
		WsEOther,
		WsESecurityTokenExpired,
		WsESecuritySystemFailure,
		ErrorNdisInterfaceClosing,
		ErrorNdisBadVersion,
		ErrorNdisBadCharacteristics,
		ErrorNdisAdapterNotFound,
		ErrorNdisOpenFailed,
		ErrorNdisDeviceFailed,
		ErrorNdisMulticastFull,
		ErrorNdisMulticastExists,
		ErrorNdisMulticastNotFound,
		ErrorNdisRequestAborted,
		ErrorNdisResetInProgress,
		ErrorNdisNotSupported,
		ErrorNdisInvalidPacket,
		ErrorNdisAdapterNotReady,
		ErrorNdisInvalidLength,
		ErrorNdisInvalidData,
		ErrorNdisBufferTooShort,
		ErrorNdisInvalidOid,
		ErrorNdisAdapterRemoved,
		ErrorNdisUnsupportedMedia,
		ErrorNdisGroupAddressInUse,
		ErrorNdisFileNotFound,
		ErrorNdisErrorReadingFile,
		ErrorNdisAlreadyMapped,
		ErrorNdisResourceConflict,
		ErrorNdisMediaDisconnected,
		ErrorNdisInvalidAddress,
		ErrorNdisInvalidDeviceRequest,
		ErrorNdisPaused,
		ErrorNdisInterfaceNotFound,
		ErrorNdisUnsupportedRevision,
		ErrorNdisInvalidPort,
		ErrorNdisInvalidPortState,
		ErrorNdisLowPowerState,
		ErrorNdisDot11AutoConfigEnabled,
		ErrorNdisDot11MediaInUse,
		ErrorNdisDot11PowerStateInvalid,
		ErrorNdisPmWolPatternListFull,
		ErrorNdisPmProtocolOffloadListFull,
		ErrorNdisIndicationRequired,
		ErrorNdisOffloadPolicy,
		ErrorNdisOffloadConnectionRejected,
		ErrorNdisOffloadPathRejected,
		ErrorHvInvalidHypercallCode,
		ErrorHvInvalidHypercallInput,
		ErrorHvInvalidAlignment,
		ErrorHvInvalidParameter,
		ErrorHvAccessDenied,
		ErrorHvInvalidPartitionState,
		ErrorHvOperationDenied,
		ErrorHvUnknownProperty,
		ErrorHvPropertyValueOutOfRange,
		ErrorHvInsufficientMemory,
		ErrorHvPartitionTooDeep,
		ErrorHvInvalidPartitionId,
		ErrorHvInvalidVpIndex,
		ErrorHvInvalidPortId,
		ErrorHvInvalidConnectionId,
		ErrorHvInsufficientBuffers,
		ErrorHvNotAcknowledged,
		ErrorHvAcknowledged,
		ErrorHvInvalidSaveRestoreState,
		ErrorHvInvalidSynicState,
		ErrorHvObjectInUse,
		ErrorHvInvalidProximityDomainInfo,
		ErrorHvNoData,
		ErrorHvInactive,
		ErrorHvNoResources,
		ErrorHvFeatureUnavailable,
		ErrorHvNotPresent,
		ErrorVidDuplicateHandler,
		ErrorVidTooManyHandlers,
		ErrorVidQueueFull,
		ErrorVidHandlerNotPresent,
		ErrorVidInvalidObjectName,
		ErrorVidPartitionNameTooLong,
		ErrorVidMessageQueueNameTooLong,
		ErrorVidPartitionAlreadyExists,
		ErrorVidPartitionDoesNotExist,
		ErrorVidPartitionNameNotFound,
		ErrorVidMessageQueueAlreadyExists,
		ErrorVidExceededMbpEntryMapLimit,
		ErrorVidMbStillReferenced,
		ErrorVidChildGpaPageSetCorrupted,
		ErrorVidInvalidNumaSettings,
		ErrorVidInvalidNumaNodeIndex,
		ErrorVidNotificationQueueAlreadyAssociated,
		ErrorVidInvalidMemoryBlockHandle,
		ErrorVidPageRangeOverflow,
		ErrorVidInvalidMessageQueueHandle,
		ErrorVidInvalidGpaRangeHandle,
		ErrorVidNoMemoryBlockNotificationQueue,
		ErrorVidMemoryBlockLockCountExceeded,
		ErrorVidInvalidPpmHandle,
		ErrorVidMbpsAreLocked,
		ErrorVidMessageQueueClosed,
		ErrorVidVirtualProcessorLimitExceeded,
		ErrorVidStopPending,
		ErrorVidInvalidProcessorState,
		ErrorVidExceededKmContextCountLimit,
		ErrorVidKmInterfaceAlreadyInitialized,
		ErrorVidMbPropertyAlreadySetReset,
		ErrorVidMmioRangeDestroyed,
		ErrorVidInvalidChildGpaPageSet,
		ErrorVidReservePageSetIsBeingUsed,
		ErrorVidReservePageSetTooSmall,
		ErrorVidMbpAlreadyLockedUsingReservedPage,
		ErrorVidMbpCountExceededLimit,
		ErrorVidSavedStateCorrupt,
		ErrorVidSavedStateUnrecognizedItem,
		ErrorVidSavedStateIncompatible,
		ErrorVidRemoteNodeParentGpaPagesUsed,
		ErrorVolmgrIncompleteRegeneration,
		ErrorVolmgrIncompleteDiskMigration,
		ErrorVolmgrDatabaseFull,
		ErrorVolmgrDiskConfigurationCorrupted,
		ErrorVolmgrDiskConfigurationNotInSync,
		ErrorVolmgrPackConfigUpdateFailed,
		ErrorVolmgrDiskContainsNonSimpleVolume,
		ErrorVolmgrDiskDuplicate,
		ErrorVolmgrDiskDynamic,
		ErrorVolmgrDiskIdInvalid,
		ErrorVolmgrDiskInvalid,
		ErrorVolmgrDiskLastVoter,
		ErrorVolmgrDiskLayoutInvalid,
		ErrorVolmgrDiskLayoutNonBasicBetweenBasicPartitions,
		ErrorVolmgrDiskLayoutNotCylinderAligned,
		ErrorVolmgrDiskLayoutPartitionsTooSmall,
		ErrorVolmgrDiskLayoutPrimaryBetweenLogicalPartitions,
		ErrorVolmgrDiskLayoutTooManyPartitions,
		ErrorVolmgrDiskMissing,
		ErrorVolmgrDiskNotEmpty,
		ErrorVolmgrDiskNotEnoughSpace,
		ErrorVolmgrDiskRevectoringFailed,
		ErrorVolmgrDiskSectorSizeInvalid,
		ErrorVolmgrDiskSetNotContained,
		ErrorVolmgrDiskUsedByMultipleMembers,
		ErrorVolmgrDiskUsedByMultiplePlexes,
		ErrorVolmgrDynamicDiskNotSupported,
		ErrorVolmgrExtentAlreadyUsed,
		ErrorVolmgrExtentNotContiguous,
		ErrorVolmgrExtentNotInPublicRegion,
		ErrorVolmgrExtentNotSectorAligned,
		ErrorVolmgrExtentOverlapsEbrPartition,
		ErrorVolmgrExtentVolumeLengthsDoNotMatch,
		ErrorVolmgrFaultTolerantNotSupported,
		ErrorVolmgrInterleaveLengthInvalid,
		ErrorVolmgrMaximumRegisteredUsers,
		ErrorVolmgrMemberInSync,
		ErrorVolmgrMemberIndexDuplicate,
		ErrorVolmgrMemberIndexInvalid,
		ErrorVolmgrMemberMissing,
		ErrorVolmgrMemberNotDetached,
		ErrorVolmgrMemberRegenerating,
		ErrorVolmgrAllDisksFailed,
		ErrorVolmgrNoRegisteredUsers,
		ErrorVolmgrNoSuchUser,
		ErrorVolmgrNotificationReset,
		ErrorVolmgrNumberOfMembersInvalid,
		ErrorVolmgrNumberOfPlexesInvalid,
		ErrorVolmgrPackDuplicate,
		ErrorVolmgrPackIdInvalid,
		ErrorVolmgrPackInvalid,
		ErrorVolmgrPackNameInvalid,
		ErrorVolmgrPackOffline,
		ErrorVolmgrPackHasQuorum,
		ErrorVolmgrPackWithoutQuorum,
		ErrorVolmgrPartitionStyleInvalid,
		ErrorVolmgrPartitionUpdateFailed,
		ErrorVolmgrPlexInSync,
		ErrorVolmgrPlexIndexDuplicate,
		ErrorVolmgrPlexIndexInvalid,
		ErrorVolmgrPlexLastActive,
		ErrorVolmgrPlexMissing,
		ErrorVolmgrPlexRegenerating,
		ErrorVolmgrPlexTypeInvalid,
		ErrorVolmgrPlexNotRaid5,
		ErrorVolmgrPlexNotSimple,
		ErrorVolmgrStructureSizeInvalid,
		ErrorVolmgrTooManyNotificationRequests,
		ErrorVolmgrTransactionInProgress,
		ErrorVolmgrUnexpectedDiskLayoutChange,
		ErrorVolmgrVolumeContainsMissingDisk,
		ErrorVolmgrVolumeIdInvalid,
		ErrorVolmgrVolumeLengthInvalid,
		ErrorVolmgrVolumeLengthNotSectorSizeMultiple,
		ErrorVolmgrVolumeNotMirrored,
		ErrorVolmgrVolumeNotRetained,
		ErrorVolmgrVolumeOffline,
		ErrorVolmgrVolumeRetained,
		ErrorVolmgrNumberOfExtentsInvalid,
		ErrorVolmgrDifferentSectorSize,
		ErrorVolmgrBadBootDisk,
		ErrorVolmgrPackConfigOffline,
		ErrorVolmgrPackConfigOnline,
		ErrorVolmgrNotPrimaryPack,
		ErrorVolmgrPackLogUpdateFailed,
		ErrorVolmgrNumberOfDisksInPlexInvalid,
		ErrorVolmgrNumberOfDisksInMemberInvalid,
		ErrorVolmgrVolumeMirrored,
		ErrorVolmgrPlexNotSimpleSpanned,
		ErrorVolmgrNoValidLogCopies,
		ErrorVolmgrPrimaryPackPresent,
		ErrorVolmgrNumberOfDisksInvalid,
		ErrorVolmgrMirrorNotSupported,
		ErrorVolmgrRaid5NotSupported,
		ErrorBcdNotAllEntriesImported,
		ErrorBcdTooManyElements,
		ErrorBcdNotAllEntriesSynchronized,
		ErrorVhdDriveFooterMissing,
		ErrorVhdDriveFooterChecksumMismatch,
		ErrorVhdDriveFooterCorrupt,
		ErrorVhdFormatUnknown,
		ErrorVhdFormatUnsupportedVersion,
		ErrorVhdSparseHeaderChecksumMismatch,
		ErrorVhdSparseHeaderUnsupportedVersion,
		ErrorVhdSparseHeaderCorrupt,
		ErrorVhdBlockAllocationFailure,
		ErrorVhdBlockAllocationTableCorrupt,
		ErrorVhdInvalidBlockSize,
		ErrorVhdBitmapMismatch,
		ErrorVhdParentVhdNotFound,
		ErrorVhdChildParentIdMismatch,
		ErrorVhdChildParentTimestampMismatch,
		ErrorVhdMetadataReadFailure,
		ErrorVhdMetadataWriteFailure,
		ErrorVhdInvalidSize,
		ErrorVhdInvalidFileSize,
		ErrorVirtdiskProviderNotFound,
		ErrorVirtdiskNotVirtualDisk,
		ErrorVhdParentVhdAccessDenied,
		ErrorVhdChildParentSizeMismatch,
		ErrorVhdDifferencingChainCycleDetected,
		ErrorVhdDifferencingChainErrorInParent,
		ErrorVirtualDiskLimitation,
		ErrorVhdInvalidType,
		ErrorVhdInvalidState,
		ErrorVirtdiskUnsupportedDiskSectorSize,
		ErrorQueryStorageError,
		SdiagECancelled,
		SdiagEScript,
		SdiagEPowershell,
		SdiagEManagedhost,
		SdiagENoverifier,
		SdiagSCannotrun,
		SdiagEDisabled,
		SdiagETrust,
		SdiagECannotrun,
		SdiagEVersion,
		SdiagEResource,
		SdiagERootcause,
		EMbnContextNotActivated,
		EMbnBadSim,
		EMbnDataClassNotAvailable,
		EMbnInvalidAccessString,
		EMbnMaxActivatedContexts,
		EMbnPacketSvcDetached,
		EMbnProviderNotVisible,
		EMbnRadioPowerOff,
		EMbnServiceNotActivated,
		EMbnSimNotInserted,
		EMbnVoiceCallInProgress,
		EMbnInvalidCache,
		EMbnNotRegistered,
		EMbnProvidersNotFound,
		EMbnPinNotSupported,
		EMbnPinRequired,
		EMbnPinDisabled,
		EMbnFailure,
		EMbnInvalidProfile,
		EMbnDefaultProfileExist,
		EMbnSmsEncodingNotSupported,
		EMbnSmsFilterNotSupported,
		EMbnSmsInvalidMemoryIndex,
		EMbnSmsLangNotSupported,
		EMbnSmsMemoryFailure,
		EMbnSmsNetworkTimeout,
		EMbnSmsUnknownSmscAddress,
		EMbnSmsFormatNotSupported,
		EMbnSmsOperationNotAllowed,
		EMbnSmsMemoryFull,
		UiECreateFailed,
		UiEShutdownCalled,
		UiEIllegalReentrancy,
		UiEObjectSealed,
		UiEValueNotSet,
		UiEValueNotDetermined,
		UiEInvalidOutput,
		UiEBooleanExpected,
		UiEDifferentOwner,
		UiEAmbiguousMatch,
		UiEFpOverflow,
		UiEWrongThread,
		UiEStoryboardActive,
		UiEStoryboardNotPlaying,
		UiEStartKeyframeAfterEnd,
		UiEEndKeyframeNotDetermined,
		UiELoopsOverlap,
		UiETransitionAlreadyUsed,
		UiETransitionNotInStoryboard,
		UiETransitionEclipsed,
		UiETimeBeforeLastUpdate,
		UiETimerClientAlreadyConnected,
	}
}
