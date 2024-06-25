package constants

import "fmt"

type NtstatusKind int

const (
	StatusSuccess                                               NtstatusKind = 0x00000000
	StatusWait1                                                              = 0x00000001
	StatusWait2                                                              = 0x00000002
	StatusWait3                                                              = 0x00000003
	StatusWait63                                                             = 0x0000003F
	StatusAbandoned                                                          = 0x00000080
	StatusAbandonedWait63                                                    = 0x000000BF
	StatusUserApc                                                            = 0x000000C0
	StatusKernelApc                                                          = 0x00000100
	StatusAlerted                                                            = 0x00000101
	StatusTimeout                                                            = 0x00000102
	StatusPending                                                            = 0x00000103
	StatusReparse                                                            = 0x00000104
	StatusMoreEntries                                                        = 0x00000105
	StatusNotAllAssigned                                                     = 0x00000106
	StatusSomeNotMapped                                                      = 0x00000107
	StatusOplockBreakInProgress                                              = 0x00000108
	StatusVolumeMounted                                                      = 0x00000109
	StatusRxactCommitted                                                     = 0x0000010A
	StatusNotifyCleanup                                                      = 0x0000010B
	StatusNotifyEnumDir                                                      = 0x0000010C
	StatusNoQuotasForAccount                                                 = 0x0000010D
	StatusPrimaryTransportConnectFailed                                      = 0x0000010E
	StatusPageFaultTransition                                                = 0x00000110
	StatusPageFaultDemandZero                                                = 0x00000111
	StatusPageFaultCopyOnWrite                                               = 0x00000112
	StatusPageFaultGuardPage                                                 = 0x00000113
	StatusPageFaultPagingFile                                                = 0x00000114
	StatusCachePageLocked                                                    = 0x00000115
	StatusCrashDump                                                          = 0x00000116
	StatusBufferAllZeros                                                     = 0x00000117
	StatusReparseObject                                                      = 0x00000118
	StatusResourceRequirementsChanged                                        = 0x00000119
	StatusTranslationComplete                                                = 0x00000120
	StatusDsMembershipEvaluatedLocally                                       = 0x00000121
	StatusNothingToTerminate                                                 = 0x00000122
	StatusProcessNotInJob                                                    = 0x00000123
	StatusProcessInJob                                                       = 0x00000124
	StatusVolsnapHibernateReady                                              = 0x00000125
	StatusFsfilterOpCompletedSuccessfully                                    = 0x00000126
	StatusInterruptVectorAlreadyConnected                                    = 0x00000127
	StatusInterruptStillConnected                                            = 0x00000128
	StatusProcessCloned                                                      = 0x00000129
	StatusFileLockedWithOnlyReaders                                          = 0x0000012A
	StatusFileLockedWithWriters                                              = 0x0000012B
	StatusResourcemanagerReadOnly                                            = 0x00000202
	StatusRingPreviouslyEmpty                                                = 0x00000210
	StatusRingPreviouslyFull                                                 = 0x00000211
	StatusRingPreviouslyAboveQuota                                           = 0x00000212
	StatusRingNewlyEmpty                                                     = 0x00000213
	StatusRingSignalOppositeEndpoint                                         = 0x00000214
	StatusOplockSwitchedToNewHandle                                          = 0x00000215
	StatusOplockHandleClosed                                                 = 0x00000216
	StatusWaitForOplock                                                      = 0x00000367
	DbgExceptionHandled                                                      = 0x00010001
	DbgContinue                                                              = 0x00010002
	StatusFltIoComplete                                                      = 0x001C0001
	StatusDisAttributeBuilt                                                  = 0x003C0001
	StatusObjectNameExists                                                   = 0x40000000
	StatusThreadWasSuspended                                                 = 0x40000001
	StatusWorkingSetLimitRange                                               = 0x40000002
	StatusImageNotAtBase                                                     = 0x40000003
	StatusRxactStateCreated                                                  = 0x40000004
	StatusSegmentNotification                                                = 0x40000005
	StatusLocalUserSessionKey                                                = 0x40000006
	StatusBadCurrentDirectory                                                = 0x40000007
	StatusSerialMoreWrites                                                   = 0x40000008
	StatusRegistryRecovered                                                  = 0x40000009
	StatusFtReadRecoveryFromBackup                                           = 0x4000000A
	StatusFtWriteRecovery                                                    = 0x4000000B
	StatusSerialCounterTimeout                                               = 0x4000000C
	StatusNullLmPassword                                                     = 0x4000000D
	StatusImageMachineTypeMismatch                                           = 0x4000000E
	StatusReceivePartial                                                     = 0x4000000F
	StatusReceiveExpedited                                                   = 0x40000010
	StatusReceivePartialExpedited                                            = 0x40000011
	StatusEventDone                                                          = 0x40000012
	StatusEventPending                                                       = 0x40000013
	StatusCheckingFileSystem                                                 = 0x40000014
	StatusFatalAppExit                                                       = 0x40000015
	StatusPredefinedHandle                                                   = 0x40000016
	StatusWasUnlocked                                                        = 0x40000017
	StatusServiceNotification                                                = 0x40000018
	StatusWasLocked                                                          = 0x40000019
	StatusLogHardError                                                       = 0x4000001A
	StatusAlreadyWin32                                                       = 0x4000001B
	StatusWx86Unsimulate                                                     = 0x4000001C
	StatusWx86Continue                                                       = 0x4000001D
	StatusWx86SingleStep                                                     = 0x4000001E
	StatusWx86Breakpoint                                                     = 0x4000001F
	StatusWx86ExceptionContinue                                              = 0x40000020
	StatusWx86ExceptionLastchance                                            = 0x40000021
	StatusWx86ExceptionChain                                                 = 0x40000022
	StatusImageMachineTypeMismatchExe                                        = 0x40000023
	StatusNoYieldPerformed                                                   = 0x40000024
	StatusTimerResumeIgnored                                                 = 0x40000025
	StatusArbitrationUnhandled                                               = 0x40000026
	StatusCardbusNotSupported                                                = 0x40000027
	StatusWx86Createwx86Tib                                                  = 0x40000028
	StatusMpProcessorMismatch                                                = 0x40000029
	StatusHibernated                                                         = 0x4000002A
	StatusResumeHibernation                                                  = 0x4000002B
	StatusFirmwareUpdated                                                    = 0x4000002C
	StatusDriversLeakingLockedPages                                          = 0x4000002D
	StatusMessageRetrieved                                                   = 0x4000002E
	StatusSystemPowerstateTransition                                         = 0x4000002F
	StatusAlpcCheckCompletionList                                            = 0x40000030
	StatusSystemPowerstateComplexTransition                                  = 0x40000031
	StatusAccessAuditByPolicy                                                = 0x40000032
	StatusAbandonHiberfile                                                   = 0x40000033
	StatusBizrulesNotEnabled                                                 = 0x40000034
	DbgReplyLater                                                            = 0x40010001
	DbgUnableToProvideHandle                                                 = 0x40010002
	DbgTerminateThread                                                       = 0x40010003
	DbgTerminateProcess                                                      = 0x40010004
	DbgControlC                                                              = 0x40010005
	DbgPrintexceptionC                                                       = 0x40010006
	DbgRipexception                                                          = 0x40010007
	DbgControlBreak                                                          = 0x40010008
	DbgCommandException                                                      = 0x40010009
	StatusHeuristicDamagePossible                                            = 0x40190001
	StatusGuardPageViolation                                                 = 0x80000001
	StatusDatatypeMisalignment                                               = 0x80000002
	StatusBreakpoint                                                         = 0x80000003
	StatusSingleStep                                                         = 0x80000004
	StatusBufferOverflow                                                     = 0x80000005
	StatusNoMoreFiles                                                        = 0x80000006
	StatusWakeSystemDebugger                                                 = 0x80000007
	StatusHandlesClosed                                                      = 0x8000000A
	StatusNoInheritance                                                      = 0x8000000B
	StatusGuidSubstitutionMade                                               = 0x8000000C
	StatusPartialCopy                                                        = 0x8000000D
	StatusDevicePaperEmpty                                                   = 0x8000000E
	StatusDevicePoweredOff                                                   = 0x8000000F
	StatusDeviceOffLine                                                      = 0x80000010
	StatusDeviceBusy                                                         = 0x80000011
	StatusNoMoreEas                                                          = 0x80000012
	StatusInvalidEaName                                                      = 0x80000013
	StatusEaListInconsistent                                                 = 0x80000014
	StatusInvalidEaFlag                                                      = 0x80000015
	StatusVerifyRequired                                                     = 0x80000016
	StatusExtraneousInformation                                              = 0x80000017
	StatusRxactCommitNecessary                                               = 0x80000018
	StatusNoMoreEntries                                                      = 0x8000001A
	StatusFilemarkDetected                                                   = 0x8000001B
	StatusMediaChanged                                                       = 0x8000001C
	StatusBusReset                                                           = 0x8000001D
	StatusEndOfMedia                                                         = 0x8000001E
	StatusBeginningOfMedia                                                   = 0x8000001F
	StatusMediaCheck                                                         = 0x80000020
	StatusSetmarkDetected                                                    = 0x80000021
	StatusNoDataDetected                                                     = 0x80000022
	StatusRedirectorHasOpenHandles                                           = 0x80000023
	StatusServerHasOpenHandles                                               = 0x80000024
	StatusAlreadyDisconnected                                                = 0x80000025
	StatusLongjump                                                           = 0x80000026
	StatusCleanerCartridgeInstalled                                          = 0x80000027
	StatusPlugplayQueryVetoed                                                = 0x80000028
	StatusUnwindConsolidate                                                  = 0x80000029
	StatusRegistryHiveRecovered                                              = 0x8000002A
	StatusDllMightBeInsecure                                                 = 0x8000002B
	StatusDllMightBeIncompatible                                             = 0x8000002C
	StatusStoppedOnSymlink                                                   = 0x8000002D
	StatusCannotGrantRequestedOplock                                         = 0x8000002E
	StatusNoAceCondition                                                     = 0x8000002F
	DbgExceptionNotHandled                                                   = 0x80010001
	StatusClusterNodeAlreadyUp                                               = 0x80130001
	StatusClusterNodeAlreadyDown                                             = 0x80130002
	StatusClusterNetworkAlreadyOnline                                        = 0x80130003
	StatusClusterNetworkAlreadyOffline                                       = 0x80130004
	StatusClusterNodeAlreadyMember                                           = 0x80130005
	StatusFltBufferTooSmall                                                  = 0x801C0001
	StatusFvePartialMetadata                                                 = 0x80210001
	StatusFveTransientState                                                  = 0x80210002
	StatusUnsuccessful                                                       = 0xC0000001
	StatusNotImplemented                                                     = 0xC0000002
	StatusInvalidInfoClass                                                   = 0xC0000003
	StatusInfoLengthMismatch                                                 = 0xC0000004
	StatusAccessViolation                                                    = 0xC0000005
	StatusInPageError                                                        = 0xC0000006
	StatusPagefileQuota                                                      = 0xC0000007
	StatusInvalidHandle                                                      = 0xC0000008
	StatusBadInitialStack                                                    = 0xC0000009
	StatusBadInitialPc                                                       = 0xC000000A
	StatusInvalidCid                                                         = 0xC000000B
	StatusTimerNotCanceled                                                   = 0xC000000C
	StatusInvalidParameter                                                   = 0xC000000D
	StatusNoSuchDevice                                                       = 0xC000000E
	StatusNoSuchFile                                                         = 0xC000000F
	StatusInvalidDeviceRequest                                               = 0xC0000010
	StatusEndOfFile                                                          = 0xC0000011
	StatusWrongVolume                                                        = 0xC0000012
	StatusNoMediaInDevice                                                    = 0xC0000013
	StatusUnrecognizedMedia                                                  = 0xC0000014
	StatusNonexistentSector                                                  = 0xC0000015
	StatusMoreProcessingRequired                                             = 0xC0000016
	StatusNoMemory                                                           = 0xC0000017
	StatusConflictingAddresses                                               = 0xC0000018
	StatusNotMappedView                                                      = 0xC0000019
	StatusUnableToFreeVm                                                     = 0xC000001A
	StatusUnableToDeleteSection                                              = 0xC000001B
	StatusInvalidSystemService                                               = 0xC000001C
	StatusIllegalInstruction                                                 = 0xC000001D
	StatusInvalidLockSequence                                                = 0xC000001E
	StatusInvalidViewSize                                                    = 0xC000001F
	StatusInvalidFileForSection                                              = 0xC0000020
	StatusAlreadyCommitted                                                   = 0xC0000021
	StatusAccessDenied                                                       = 0xC0000022
	StatusBufferTooSmall                                                     = 0xC0000023
	StatusObjectTypeMismatch                                                 = 0xC0000024
	StatusNoncontinuableException                                            = 0xC0000025
	StatusInvalidDisposition                                                 = 0xC0000026
	StatusUnwind                                                             = 0xC0000027
	StatusBadStack                                                           = 0xC0000028
	StatusInvalidUnwindTarget                                                = 0xC0000029
	StatusNotLocked                                                          = 0xC000002A
	StatusParityError                                                        = 0xC000002B
	StatusUnableToDecommitVm                                                 = 0xC000002C
	StatusNotCommitted                                                       = 0xC000002D
	StatusInvalidPortAttributes                                              = 0xC000002E
	StatusPortMessageTooLong                                                 = 0xC000002F
	StatusInvalidParameterMix                                                = 0xC0000030
	StatusInvalidQuotaLower                                                  = 0xC0000031
	StatusDiskCorruptError                                                   = 0xC0000032
	StatusObjectNameInvalid                                                  = 0xC0000033
	StatusObjectNameNotFound                                                 = 0xC0000034
	StatusObjectNameCollision                                                = 0xC0000035
	StatusPortDisconnected                                                   = 0xC0000037
	StatusDeviceAlreadyAttached                                              = 0xC0000038
	StatusObjectPathInvalid                                                  = 0xC0000039
	StatusObjectPathNotFound                                                 = 0xC000003A
	StatusObjectPathSyntaxBad                                                = 0xC000003B
	StatusDataOverrun                                                        = 0xC000003C
	StatusDataLateError                                                      = 0xC000003D
	StatusDataError                                                          = 0xC000003E
	StatusCrcError                                                           = 0xC000003F
	StatusSectionTooBig                                                      = 0xC0000040
	StatusPortConnectionRefused                                              = 0xC0000041
	StatusInvalidPortHandle                                                  = 0xC0000042
	StatusSharingViolation                                                   = 0xC0000043
	StatusQuotaExceeded                                                      = 0xC0000044
	StatusInvalidPageProtection                                              = 0xC0000045
	StatusMutantNotOwned                                                     = 0xC0000046
	StatusSemaphoreLimitExceeded                                             = 0xC0000047
	StatusPortAlreadySet                                                     = 0xC0000048
	StatusSectionNotImage                                                    = 0xC0000049
	StatusSuspendCountExceeded                                               = 0xC000004A
	StatusThreadIsTerminating                                                = 0xC000004B
	StatusBadWorkingSetLimit                                                 = 0xC000004C
	StatusIncompatibleFileMap                                                = 0xC000004D
	StatusSectionProtection                                                  = 0xC000004E
	StatusEasNotSupported                                                    = 0xC000004F
	StatusEaTooLarge                                                         = 0xC0000050
	StatusNonexistentEaEntry                                                 = 0xC0000051
	StatusNoEasOnFile                                                        = 0xC0000052
	StatusEaCorruptError                                                     = 0xC0000053
	StatusFileLockConflict                                                   = 0xC0000054
	StatusLockNotGranted                                                     = 0xC0000055
	StatusDeletePending                                                      = 0xC0000056
	StatusCtlFileNotSupported                                                = 0xC0000057
	StatusUnknownRevision                                                    = 0xC0000058
	StatusRevisionMismatch                                                   = 0xC0000059
	StatusInvalidOwner                                                       = 0xC000005A
	StatusInvalidPrimaryGroup                                                = 0xC000005B
	StatusNoImpersonationToken                                               = 0xC000005C
	StatusCantDisableMandatory                                               = 0xC000005D
	StatusNoLogonServers                                                     = 0xC000005E
	StatusNoSuchLogonSession                                                 = 0xC000005F
	StatusNoSuchPrivilege                                                    = 0xC0000060
	StatusPrivilegeNotHeld                                                   = 0xC0000061
	StatusInvalidAccountName                                                 = 0xC0000062
	StatusUserExists                                                         = 0xC0000063
	StatusNoSuchUser                                                         = 0xC0000064
	StatusGroupExists                                                        = 0xC0000065
	StatusNoSuchGroup                                                        = 0xC0000066
	StatusMemberInGroup                                                      = 0xC0000067
	StatusMemberNotInGroup                                                   = 0xC0000068
	StatusLastAdmin                                                          = 0xC0000069
	StatusWrongPassword                                                      = 0xC000006A
	StatusIllFormedPassword                                                  = 0xC000006B
	StatusPasswordRestriction                                                = 0xC000006C
	StatusLogonFailure                                                       = 0xC000006D
	StatusAccountRestriction                                                 = 0xC000006E
	StatusInvalidLogonHours                                                  = 0xC000006F
	StatusInvalidWorkstation                                                 = 0xC0000070
	StatusPasswordExpired                                                    = 0xC0000071
	StatusAccountDisabled                                                    = 0xC0000072
	StatusNoneMapped                                                         = 0xC0000073
	StatusTooManyLuidsRequested                                              = 0xC0000074
	StatusLuidsExhausted                                                     = 0xC0000075
	StatusInvalidSubAuthority                                                = 0xC0000076
	StatusInvalidAcl                                                         = 0xC0000077
	StatusInvalidSid                                                         = 0xC0000078
	StatusInvalidSecurityDescr                                               = 0xC0000079
	StatusProcedureNotFound                                                  = 0xC000007A
	StatusInvalidImageFormat                                                 = 0xC000007B
	StatusNoToken                                                            = 0xC000007C
	StatusBadInheritanceAcl                                                  = 0xC000007D
	StatusRangeNotLocked                                                     = 0xC000007E
	StatusDiskFull                                                           = 0xC000007F
	StatusServerDisabled                                                     = 0xC0000080
	StatusServerNotDisabled                                                  = 0xC0000081
	StatusTooManyGuidsRequested                                              = 0xC0000082
	StatusGuidsExhausted                                                     = 0xC0000083
	StatusInvalidIdAuthority                                                 = 0xC0000084
	StatusAgentsExhausted                                                    = 0xC0000085
	StatusInvalidVolumeLabel                                                 = 0xC0000086
	StatusSectionNotExtended                                                 = 0xC0000087
	StatusNotMappedData                                                      = 0xC0000088
	StatusResourceDataNotFound                                               = 0xC0000089
	StatusResourceTypeNotFound                                               = 0xC000008A
	StatusResourceNameNotFound                                               = 0xC000008B
	StatusArrayBoundsExceeded                                                = 0xC000008C
	StatusFloatDenormalOperand                                               = 0xC000008D
	StatusFloatDivideByZero                                                  = 0xC000008E
	StatusFloatInexactResult                                                 = 0xC000008F
	StatusFloatInvalidOperation                                              = 0xC0000090
	StatusFloatOverflow                                                      = 0xC0000091
	StatusFloatStackCheck                                                    = 0xC0000092
	StatusFloatUnderflow                                                     = 0xC0000093
	StatusIntegerDivideByZero                                                = 0xC0000094
	StatusIntegerOverflow                                                    = 0xC0000095
	StatusPrivilegedInstruction                                              = 0xC0000096
	StatusTooManyPagingFiles                                                 = 0xC0000097
	StatusFileInvalid                                                        = 0xC0000098
	StatusAllottedSpaceExceeded                                              = 0xC0000099
	StatusInsufficientResources                                              = 0xC000009A
	StatusDfsExitPathFound                                                   = 0xC000009B
	StatusDeviceDataError                                                    = 0xC000009C
	StatusDeviceNotConnected                                                 = 0xC000009D
	StatusDevicePowerFailure                                                 = 0xC000009E
	StatusFreeVmNotAtBase                                                    = 0xC000009F
	StatusMemoryNotAllocated                                                 = 0xC00000A0
	StatusWorkingSetQuota                                                    = 0xC00000A1
	StatusMediaWriteProtected                                                = 0xC00000A2
	StatusDeviceNotReady                                                     = 0xC00000A3
	StatusInvalidGroupAttributes                                             = 0xC00000A4
	StatusBadImpersonationLevel                                              = 0xC00000A5
	StatusCantOpenAnonymous                                                  = 0xC00000A6
	StatusBadValidationClass                                                 = 0xC00000A7
	StatusBadTokenType                                                       = 0xC00000A8
	StatusBadMasterBootRecord                                                = 0xC00000A9
	StatusInstructionMisalignment                                            = 0xC00000AA
	StatusInstanceNotAvailable                                               = 0xC00000AB
	StatusPipeNotAvailable                                                   = 0xC00000AC
	StatusInvalidPipeState                                                   = 0xC00000AD
	StatusPipeBusy                                                           = 0xC00000AE
	StatusIllegalFunction                                                    = 0xC00000AF
	StatusPipeDisconnected                                                   = 0xC00000B0
	StatusPipeClosing                                                        = 0xC00000B1
	StatusPipeConnected                                                      = 0xC00000B2
	StatusPipeListening                                                      = 0xC00000B3
	StatusInvalidReadMode                                                    = 0xC00000B4
	StatusIoTimeout                                                          = 0xC00000B5
	StatusFileForcedClosed                                                   = 0xC00000B6
	StatusProfilingNotStarted                                                = 0xC00000B7
	StatusProfilingNotStopped                                                = 0xC00000B8
	StatusCouldNotInterpret                                                  = 0xC00000B9
	StatusFileIsADirectory                                                   = 0xC00000BA
	StatusNotSupported                                                       = 0xC00000BB
	StatusRemoteNotListening                                                 = 0xC00000BC
	StatusDuplicateName                                                      = 0xC00000BD
	StatusBadNetworkPath                                                     = 0xC00000BE
	StatusNetworkBusy                                                        = 0xC00000BF
	StatusDeviceDoesNotExist                                                 = 0xC00000C0
	StatusTooManyCommands                                                    = 0xC00000C1
	StatusAdapterHardwareError                                               = 0xC00000C2
	StatusInvalidNetworkResponse                                             = 0xC00000C3
	StatusUnexpectedNetworkError                                             = 0xC00000C4
	StatusBadRemoteAdapter                                                   = 0xC00000C5
	StatusPrintQueueFull                                                     = 0xC00000C6
	StatusNoSpoolSpace                                                       = 0xC00000C7
	StatusPrintCancelled                                                     = 0xC00000C8
	StatusNetworkNameDeleted                                                 = 0xC00000C9
	StatusNetworkAccessDenied                                                = 0xC00000CA
	StatusBadDeviceType                                                      = 0xC00000CB
	StatusBadNetworkName                                                     = 0xC00000CC
	StatusTooManyNames                                                       = 0xC00000CD
	StatusTooManySessions                                                    = 0xC00000CE
	StatusSharingPaused                                                      = 0xC00000CF
	StatusRequestNotAccepted                                                 = 0xC00000D0
	StatusRedirectorPaused                                                   = 0xC00000D1
	StatusNetWriteFault                                                      = 0xC00000D2
	StatusProfilingAtLimit                                                   = 0xC00000D3
	StatusNotSameDevice                                                      = 0xC00000D4
	StatusFileRenamed                                                        = 0xC00000D5
	StatusVirtualCircuitClosed                                               = 0xC00000D6
	StatusNoSecurityOnObject                                                 = 0xC00000D7
	StatusCantWait                                                           = 0xC00000D8
	StatusPipeEmpty                                                          = 0xC00000D9
	StatusCantAccessDomainInfo                                               = 0xC00000DA
	StatusCantTerminateSelf                                                  = 0xC00000DB
	StatusInvalidServerState                                                 = 0xC00000DC
	StatusInvalidDomainState                                                 = 0xC00000DD
	StatusInvalidDomainRole                                                  = 0xC00000DE
	StatusNoSuchDomain                                                       = 0xC00000DF
	StatusDomainExists                                                       = 0xC00000E0
	StatusDomainLimitExceeded                                                = 0xC00000E1
	StatusOplockNotGranted                                                   = 0xC00000E2
	StatusInvalidOplockProtocol                                              = 0xC00000E3
	StatusInternalDbCorruption                                               = 0xC00000E4
	StatusInternalError                                                      = 0xC00000E5
	StatusGenericNotMapped                                                   = 0xC00000E6
	StatusBadDescriptorFormat                                                = 0xC00000E7
	StatusInvalidUserBuffer                                                  = 0xC00000E8
	StatusUnexpectedIoError                                                  = 0xC00000E9
	StatusUnexpectedMmCreateErr                                              = 0xC00000EA
	StatusUnexpectedMmMapError                                               = 0xC00000EB
	StatusUnexpectedMmExtendErr                                              = 0xC00000EC
	StatusNotLogonProcess                                                    = 0xC00000ED
	StatusLogonSessionExists                                                 = 0xC00000EE
	StatusInvalidParameter1                                                  = 0xC00000EF
	StatusInvalidParameter2                                                  = 0xC00000F0
	StatusInvalidParameter3                                                  = 0xC00000F1
	StatusInvalidParameter4                                                  = 0xC00000F2
	StatusInvalidParameter5                                                  = 0xC00000F3
	StatusInvalidParameter6                                                  = 0xC00000F4
	StatusInvalidParameter7                                                  = 0xC00000F5
	StatusInvalidParameter8                                                  = 0xC00000F6
	StatusInvalidParameter9                                                  = 0xC00000F7
	StatusInvalidParameter10                                                 = 0xC00000F8
	StatusInvalidParameter11                                                 = 0xC00000F9
	StatusInvalidParameter12                                                 = 0xC00000FA
	StatusRedirectorNotStarted                                               = 0xC00000FB
	StatusRedirectorStarted                                                  = 0xC00000FC
	StatusStackOverflow                                                      = 0xC00000FD
	StatusNoSuchPackage                                                      = 0xC00000FE
	StatusBadFunctionTable                                                   = 0xC00000FF
	StatusVariableNotFound                                                   = 0xC0000100
	StatusDirectoryNotEmpty                                                  = 0xC0000101
	StatusFileCorruptError                                                   = 0xC0000102
	StatusNotADirectory                                                      = 0xC0000103
	StatusBadLogonSessionState                                               = 0xC0000104
	StatusLogonSessionCollision                                              = 0xC0000105
	StatusNameTooLong                                                        = 0xC0000106
	StatusFilesOpen                                                          = 0xC0000107
	StatusConnectionInUse                                                    = 0xC0000108
	StatusMessageNotFound                                                    = 0xC0000109
	StatusProcessIsTerminating                                               = 0xC000010A
	StatusInvalidLogonType                                                   = 0xC000010B
	StatusNoGuidTranslation                                                  = 0xC000010C
	StatusCannotImpersonate                                                  = 0xC000010D
	StatusImageAlreadyLoaded                                                 = 0xC000010E
	StatusAbiosNotPresent                                                    = 0xC000010F
	StatusAbiosLidNotExist                                                   = 0xC0000110
	StatusAbiosLidAlreadyOwned                                               = 0xC0000111
	StatusAbiosNotLidOwner                                                   = 0xC0000112
	StatusAbiosInvalidCommand                                                = 0xC0000113
	StatusAbiosInvalidLid                                                    = 0xC0000114
	StatusAbiosSelectorNotAvailable                                          = 0xC0000115
	StatusAbiosInvalidSelector                                               = 0xC0000116
	StatusNoLdt                                                              = 0xC0000117
	StatusInvalidLdtSize                                                     = 0xC0000118
	StatusInvalidLdtOffset                                                   = 0xC0000119
	StatusInvalidLdtDescriptor                                               = 0xC000011A
	StatusInvalidImageNeFormat                                               = 0xC000011B
	StatusRxactInvalidState                                                  = 0xC000011C
	StatusRxactCommitFailure                                                 = 0xC000011D
	StatusMappedFileSizeZero                                                 = 0xC000011E
	StatusTooManyOpenedFiles                                                 = 0xC000011F
	StatusCancelled                                                          = 0xC0000120
	StatusCannotDelete                                                       = 0xC0000121
	StatusInvalidComputerName                                                = 0xC0000122
	StatusFileDeleted                                                        = 0xC0000123
	StatusSpecialAccount                                                     = 0xC0000124
	StatusSpecialGroup                                                       = 0xC0000125
	StatusSpecialUser                                                        = 0xC0000126
	StatusMembersPrimaryGroup                                                = 0xC0000127
	StatusFileClosed                                                         = 0xC0000128
	StatusTooManyThreads                                                     = 0xC0000129
	StatusThreadNotInProcess                                                 = 0xC000012A
	StatusTokenAlreadyInUse                                                  = 0xC000012B
	StatusPagefileQuotaExceeded                                              = 0xC000012C
	StatusCommitmentLimit                                                    = 0xC000012D
	StatusInvalidImageLeFormat                                               = 0xC000012E
	StatusInvalidImageNotMz                                                  = 0xC000012F
	StatusInvalidImageProtect                                                = 0xC0000130
	StatusInvalidImageWin16                                                  = 0xC0000131
	StatusLogonServerConflict                                                = 0xC0000132
	StatusTimeDifferenceAtDc                                                 = 0xC0000133
	StatusSynchronizationRequired                                            = 0xC0000134
	StatusDllNotFound                                                        = 0xC0000135
	StatusOpenFailed                                                         = 0xC0000136
	StatusIoPrivilegeFailed                                                  = 0xC0000137
	StatusOrdinalNotFound                                                    = 0xC0000138
	StatusEntrypointNotFound                                                 = 0xC0000139
	StatusControlCExit                                                       = 0xC000013A
	StatusLocalDisconnect                                                    = 0xC000013B
	StatusRemoteDisconnect                                                   = 0xC000013C
	StatusRemoteResources                                                    = 0xC000013D
	StatusLinkFailed                                                         = 0xC000013E
	StatusLinkTimeout                                                        = 0xC000013F
	StatusInvalidConnection                                                  = 0xC0000140
	StatusInvalidAddress                                                     = 0xC0000141
	StatusDllInitFailed                                                      = 0xC0000142
	StatusMissingSystemfile                                                  = 0xC0000143
	StatusUnhandledException                                                 = 0xC0000144
	StatusAppInitFailure                                                     = 0xC0000145
	StatusPagefileCreateFailed                                               = 0xC0000146
	StatusNoPagefile                                                         = 0xC0000147
	StatusInvalidLevel                                                       = 0xC0000148
	StatusWrongPasswordCore                                                  = 0xC0000149
	StatusIllegalFloatContext                                                = 0xC000014A
	StatusPipeBroken                                                         = 0xC000014B
	StatusRegistryCorrupt                                                    = 0xC000014C
	StatusRegistryIoFailed                                                   = 0xC000014D
	StatusNoEventPair                                                        = 0xC000014E
	StatusUnrecognizedVolume                                                 = 0xC000014F
	StatusSerialNoDeviceInited                                               = 0xC0000150
	StatusNoSuchAlias                                                        = 0xC0000151
	StatusMemberNotInAlias                                                   = 0xC0000152
	StatusMemberInAlias                                                      = 0xC0000153
	StatusAliasExists                                                        = 0xC0000154
	StatusLogonNotGranted                                                    = 0xC0000155
	StatusTooManySecrets                                                     = 0xC0000156
	StatusSecretTooLong                                                      = 0xC0000157
	StatusInternalDbError                                                    = 0xC0000158
	StatusFullscreenMode                                                     = 0xC0000159
	StatusTooManyContextIds                                                  = 0xC000015A
	StatusLogonTypeNotGranted                                                = 0xC000015B
	StatusNotRegistryFile                                                    = 0xC000015C
	StatusNtCrossEncryptionRequired                                          = 0xC000015D
	StatusDomainCtrlrConfigError                                             = 0xC000015E
	StatusFtMissingMember                                                    = 0xC000015F
	StatusIllFormedServiceEntry                                              = 0xC0000160
	StatusIllegalCharacter                                                   = 0xC0000161
	StatusUnmappableCharacter                                                = 0xC0000162
	StatusUndefinedCharacter                                                 = 0xC0000163
	StatusFloppyVolume                                                       = 0xC0000164
	StatusFloppyIdMarkNotFound                                               = 0xC0000165
	StatusFloppyWrongCylinder                                                = 0xC0000166
	StatusFloppyUnknownError                                                 = 0xC0000167
	StatusFloppyBadRegisters                                                 = 0xC0000168
	StatusDiskRecalibrateFailed                                              = 0xC0000169
	StatusDiskOperationFailed                                                = 0xC000016A
	StatusDiskResetFailed                                                    = 0xC000016B
	StatusSharedIrqBusy                                                      = 0xC000016C
	StatusFtOrphaning                                                        = 0xC000016D
	StatusBiosFailedToConnectInterrupt                                       = 0xC000016E
	StatusPartitionFailure                                                   = 0xC0000172
	StatusInvalidBlockLength                                                 = 0xC0000173
	StatusDeviceNotPartitioned                                               = 0xC0000174
	StatusUnableToLockMedia                                                  = 0xC0000175
	StatusUnableToUnloadMedia                                                = 0xC0000176
	StatusEomOverflow                                                        = 0xC0000177
	StatusNoMedia                                                            = 0xC0000178
	StatusNoSuchMember                                                       = 0xC000017A
	StatusInvalidMember                                                      = 0xC000017B
	StatusKeyDeleted                                                         = 0xC000017C
	StatusNoLogSpace                                                         = 0xC000017D
	StatusTooManySids                                                        = 0xC000017E
	StatusLmCrossEncryptionRequired                                          = 0xC000017F
	StatusKeyHasChildren                                                     = 0xC0000180
	StatusChildMustBeVolatile                                                = 0xC0000181
	StatusDeviceConfigurationError                                           = 0xC0000182
	StatusDriverInternalError                                                = 0xC0000183
	StatusInvalidDeviceState                                                 = 0xC0000184
	StatusIoDeviceError                                                      = 0xC0000185
	StatusDeviceProtocolError                                                = 0xC0000186
	StatusBackupController                                                   = 0xC0000187
	StatusLogFileFull                                                        = 0xC0000188
	StatusTooLate                                                            = 0xC0000189
	StatusNoTrustLsaSecret                                                   = 0xC000018A
	StatusNoTrustSamAccount                                                  = 0xC000018B
	StatusTrustedDomainFailure                                               = 0xC000018C
	StatusTrustedRelationshipFailure                                         = 0xC000018D
	StatusEventlogFileCorrupt                                                = 0xC000018E
	StatusEventlogCantStart                                                  = 0xC000018F
	StatusTrustFailure                                                       = 0xC0000190
	StatusMutantLimitExceeded                                                = 0xC0000191
	StatusNetlogonNotStarted                                                 = 0xC0000192
	StatusAccountExpired                                                     = 0xC0000193
	StatusPossibleDeadlock                                                   = 0xC0000194
	StatusNetworkCredentialConflict                                          = 0xC0000195
	StatusRemoteSessionLimit                                                 = 0xC0000196
	StatusEventlogFileChanged                                                = 0xC0000197
	StatusNologonInterdomainTrustAccount                                     = 0xC0000198
	StatusNologonWorkstationTrustAccount                                     = 0xC0000199
	StatusNologonServerTrustAccount                                          = 0xC000019A
	StatusDomainTrustInconsistent                                            = 0xC000019B
	StatusFsDriverRequired                                                   = 0xC000019C
	StatusImageAlreadyLoadedAsDll                                            = 0xC000019D
	StatusIncompatibleWithGlobalShortNameRegistrySetting                     = 0xC000019E
	StatusShortNamesNotEnabledOnVolume                                       = 0xC000019F
	StatusSecurityStreamIsInconsistent                                       = 0xC00001A0
	StatusInvalidLockRange                                                   = 0xC00001A1
	StatusInvalidAceCondition                                                = 0xC00001A2
	StatusImageSubsystemNotPresent                                           = 0xC00001A3
	StatusNotificationGuidAlreadyDefined                                     = 0xC00001A4
	StatusNetworkOpenRestriction                                             = 0xC0000201
	StatusNoUserSessionKey                                                   = 0xC0000202
	StatusUserSessionDeleted                                                 = 0xC0000203
	StatusResourceLangNotFound                                               = 0xC0000204
	StatusInsuffServerResources                                              = 0xC0000205
	StatusInvalidBufferSize                                                  = 0xC0000206
	StatusInvalidAddressComponent                                            = 0xC0000207
	StatusInvalidAddressWildcard                                             = 0xC0000208
	StatusTooManyAddresses                                                   = 0xC0000209
	StatusAddressAlreadyExists                                               = 0xC000020A
	StatusAddressClosed                                                      = 0xC000020B
	StatusConnectionDisconnected                                             = 0xC000020C
	StatusConnectionReset                                                    = 0xC000020D
	StatusTooManyNodes                                                       = 0xC000020E
	StatusTransactionAborted                                                 = 0xC000020F
	StatusTransactionTimedOut                                                = 0xC0000210
	StatusTransactionNoRelease                                               = 0xC0000211
	StatusTransactionNoMatch                                                 = 0xC0000212
	StatusTransactionResponded                                               = 0xC0000213
	StatusTransactionInvalidId                                               = 0xC0000214
	StatusTransactionInvalidType                                             = 0xC0000215
	StatusNotServerSession                                                   = 0xC0000216
	StatusNotClientSession                                                   = 0xC0000217
	StatusCannotLoadRegistryFile                                             = 0xC0000218
	StatusDebugAttachFailed                                                  = 0xC0000219
	StatusSystemProcessTerminated                                            = 0xC000021A
	StatusDataNotAccepted                                                    = 0xC000021B
	StatusNoBrowserServersFound                                              = 0xC000021C
	StatusVdmHardError                                                       = 0xC000021D
	StatusDriverCancelTimeout                                                = 0xC000021E
	StatusReplyMessageMismatch                                               = 0xC000021F
	StatusMappedAlignment                                                    = 0xC0000220
	StatusImageChecksumMismatch                                              = 0xC0000221
	StatusLostWritebehindData                                                = 0xC0000222
	StatusClientServerParametersInvalid                                      = 0xC0000223
	StatusPasswordMustChange                                                 = 0xC0000224
	StatusNotFound                                                           = 0xC0000225
	StatusNotTinyStream                                                      = 0xC0000226
	StatusRecoveryFailure                                                    = 0xC0000227
	StatusStackOverflowRead                                                  = 0xC0000228
	StatusFailCheck                                                          = 0xC0000229
	StatusDuplicateObjectid                                                  = 0xC000022A
	StatusObjectidExists                                                     = 0xC000022B
	StatusConvertToLarge                                                     = 0xC000022C
	StatusRetry                                                              = 0xC000022D
	StatusFoundOutOfScope                                                    = 0xC000022E
	StatusAllocateBucket                                                     = 0xC000022F
	StatusPropsetNotFound                                                    = 0xC0000230
	StatusMarshallOverflow                                                   = 0xC0000231
	StatusInvalidVariant                                                     = 0xC0000232
	StatusDomainControllerNotFound                                           = 0xC0000233
	StatusAccountLockedOut                                                   = 0xC0000234
	StatusHandleNotClosable                                                  = 0xC0000235
	StatusConnectionRefused                                                  = 0xC0000236
	StatusGracefulDisconnect                                                 = 0xC0000237
	StatusAddressAlreadyAssociated                                           = 0xC0000238
	StatusAddressNotAssociated                                               = 0xC0000239
	StatusConnectionInvalid                                                  = 0xC000023A
	StatusConnectionActive                                                   = 0xC000023B
	StatusNetworkUnreachable                                                 = 0xC000023C
	StatusHostUnreachable                                                    = 0xC000023D
	StatusProtocolUnreachable                                                = 0xC000023E
	StatusPortUnreachable                                                    = 0xC000023F
	StatusRequestAborted                                                     = 0xC0000240
	StatusConnectionAborted                                                  = 0xC0000241
	StatusBadCompressionBuffer                                               = 0xC0000242
	StatusUserMappedFile                                                     = 0xC0000243
	StatusAuditFailed                                                        = 0xC0000244
	StatusTimerResolutionNotSet                                              = 0xC0000245
	StatusConnectionCountLimit                                               = 0xC0000246
	StatusLoginTimeRestriction                                               = 0xC0000247
	StatusLoginWkstaRestriction                                              = 0xC0000248
	StatusImageMpUpMismatch                                                  = 0xC0000249
	StatusInsufficientLogonInfo                                              = 0xC0000250
	StatusBadDllEntrypoint                                                   = 0xC0000251
	StatusBadServiceEntrypoint                                               = 0xC0000252
	StatusLpcReplyLost                                                       = 0xC0000253
	StatusIpAddressConflict1                                                 = 0xC0000254
	StatusIpAddressConflict2                                                 = 0xC0000255
	StatusRegistryQuotaLimit                                                 = 0xC0000256
	StatusPathNotCovered                                                     = 0xC0000257
	StatusNoCallbackActive                                                   = 0xC0000258
	StatusLicenseQuotaExceeded                                               = 0xC0000259
	StatusPwdTooShort                                                        = 0xC000025A
	StatusPwdTooRecent                                                       = 0xC000025B
	StatusPwdHistoryConflict                                                 = 0xC000025C
	StatusPlugplayNoDevice                                                   = 0xC000025E
	StatusUnsupportedCompression                                             = 0xC000025F
	StatusInvalidHwProfile                                                   = 0xC0000260
	StatusInvalidPlugplayDevicePath                                          = 0xC0000261
	StatusDriverOrdinalNotFound                                              = 0xC0000262
	StatusDriverEntrypointNotFound                                           = 0xC0000263
	StatusResourceNotOwned                                                   = 0xC0000264
	StatusTooManyLinks                                                       = 0xC0000265
	StatusQuotaListInconsistent                                              = 0xC0000266
	StatusFileIsOffline                                                      = 0xC0000267
	StatusEvaluationExpiration                                               = 0xC0000268
	StatusIllegalDllRelocation                                               = 0xC0000269
	StatusLicenseViolation                                                   = 0xC000026A
	StatusDllInitFailedLogoff                                                = 0xC000026B
	StatusDriverUnableToLoad                                                 = 0xC000026C
	StatusDfsUnavailable                                                     = 0xC000026D
	StatusVolumeDismounted                                                   = 0xC000026E
	StatusWx86InternalError                                                  = 0xC000026F
	StatusWx86FloatStackCheck                                                = 0xC0000270
	StatusValidateContinue                                                   = 0xC0000271
	StatusNoMatch                                                            = 0xC0000272
	StatusNoMoreMatches                                                      = 0xC0000273
	StatusNotAReparsePoint                                                   = 0xC0000275
	StatusIoReparseTagInvalid                                                = 0xC0000276
	StatusIoReparseTagMismatch                                               = 0xC0000277
	StatusIoReparseDataInvalid                                               = 0xC0000278
	StatusIoReparseTagNotHandled                                             = 0xC0000279
	StatusReparsePointNotResolved                                            = 0xC0000280
	StatusDirectoryIsAReparsePoint                                           = 0xC0000281
	StatusRangeListConflict                                                  = 0xC0000282
	StatusSourceElementEmpty                                                 = 0xC0000283
	StatusDestinationElementFull                                             = 0xC0000284
	StatusIllegalElementAddress                                              = 0xC0000285
	StatusMagazineNotPresent                                                 = 0xC0000286
	StatusReinitializationNeeded                                             = 0xC0000287
	StatusDeviceRequiresCleaning                                             = 0x80000288
	StatusDeviceDoorOpen                                                     = 0x80000289
	StatusEncryptionFailed                                                   = 0xC000028A
	StatusDecryptionFailed                                                   = 0xC000028B
	StatusRangeNotFound                                                      = 0xC000028C
	StatusNoRecoveryPolicy                                                   = 0xC000028D
	StatusNoEfs                                                              = 0xC000028E
	StatusWrongEfs                                                           = 0xC000028F
	StatusNoUserKeys                                                         = 0xC0000290
	StatusFileNotEncrypted                                                   = 0xC0000291
	StatusNotExportFormat                                                    = 0xC0000292
	StatusFileEncrypted                                                      = 0xC0000293
	StatusWakeSystem                                                         = 0x40000294
	StatusWmiGuidNotFound                                                    = 0xC0000295
	StatusWmiInstanceNotFound                                                = 0xC0000296
	StatusWmiItemidNotFound                                                  = 0xC0000297
	StatusWmiTryAgain                                                        = 0xC0000298
	StatusSharedPolicy                                                       = 0xC0000299
	StatusPolicyObjectNotFound                                               = 0xC000029A
	StatusPolicyOnlyInDs                                                     = 0xC000029B
	StatusVolumeNotUpgraded                                                  = 0xC000029C
	StatusRemoteStorageNotActive                                             = 0xC000029D
	StatusRemoteStorageMediaError                                            = 0xC000029E
	StatusNoTrackingService                                                  = 0xC000029F
	StatusServerSidMismatch                                                  = 0xC00002A0
	StatusDsNoAttributeOrValue                                               = 0xC00002A1
	StatusDsInvalidAttributeSyntax                                           = 0xC00002A2
	StatusDsAttributeTypeUndefined                                           = 0xC00002A3
	StatusDsAttributeOrValueExists                                           = 0xC00002A4
	StatusDsBusy                                                             = 0xC00002A5
	StatusDsUnavailable                                                      = 0xC00002A6
	StatusDsNoRidsAllocated                                                  = 0xC00002A7
	StatusDsNoMoreRids                                                       = 0xC00002A8
	StatusDsIncorrectRoleOwner                                               = 0xC00002A9
	StatusDsRidmgrInitError                                                  = 0xC00002AA
	StatusDsObjClassViolation                                                = 0xC00002AB
	StatusDsCantOnNonLeaf                                                    = 0xC00002AC
	StatusDsCantOnRdn                                                        = 0xC00002AD
	StatusDsCantModObjClass                                                  = 0xC00002AE
	StatusDsCrossDomMoveFailed                                               = 0xC00002AF
	StatusDsGcNotAvailable                                                   = 0xC00002B0
	StatusDirectoryServiceRequired                                           = 0xC00002B1
	StatusReparseAttributeConflict                                           = 0xC00002B2
	StatusCantEnableDenyOnly                                                 = 0xC00002B3
	StatusFloatMultipleFaults                                                = 0xC00002B4
	StatusFloatMultipleTraps                                                 = 0xC00002B5
	StatusDeviceRemoved                                                      = 0xC00002B6
	StatusJournalDeleteInProgress                                            = 0xC00002B7
	StatusJournalNotActive                                                   = 0xC00002B8
	StatusNointerface                                                        = 0xC00002B9
	StatusDsAdminLimitExceeded                                               = 0xC00002C1
	StatusDriverFailedSleep                                                  = 0xC00002C2
	StatusMutualAuthenticationFailed                                         = 0xC00002C3
	StatusCorruptSystemFile                                                  = 0xC00002C4
	StatusDatatypeMisalignmentError                                          = 0xC00002C5
	StatusWmiReadOnly                                                        = 0xC00002C6
	StatusWmiSetFailure                                                      = 0xC00002C7
	StatusCommitmentMinimum                                                  = 0xC00002C8
	StatusRegNatConsumption                                                  = 0xC00002C9
	StatusTransportFull                                                      = 0xC00002CA
	StatusDsSamInitFailure                                                   = 0xC00002CB
	StatusOnlyIfConnected                                                    = 0xC00002CC
	StatusDsSensitiveGroupViolation                                          = 0xC00002CD
	StatusPnpRestartEnumeration                                              = 0xC00002CE
	StatusJournalEntryDeleted                                                = 0xC00002CF
	StatusDsCantModPrimarygroupid                                            = 0xC00002D0
	StatusSystemImageBadSignature                                            = 0xC00002D1
	StatusPnpRebootRequired                                                  = 0xC00002D2
	StatusPowerStateInvalid                                                  = 0xC00002D3
	StatusDsInvalidGroupType                                                 = 0xC00002D4
	StatusDsNoNestGlobalgroupInMixeddomain                                   = 0xC00002D5
	StatusDsNoNestLocalgroupInMixeddomain                                    = 0xC00002D6
	StatusDsGlobalCantHaveLocalMember                                        = 0xC00002D7
	StatusDsGlobalCantHaveUniversalMember                                    = 0xC00002D8
	StatusDsUniversalCantHaveLocalMember                                     = 0xC00002D9
	StatusDsGlobalCantHaveCrossdomainMember                                  = 0xC00002DA
	StatusDsLocalCantHaveCrossdomainLocalMember                              = 0xC00002DB
	StatusDsHavePrimaryMembers                                               = 0xC00002DC
	StatusWmiNotSupported                                                    = 0xC00002DD
	StatusInsufficientPower                                                  = 0xC00002DE
	StatusSamNeedBootkeyPassword                                             = 0xC00002DF
	StatusSamNeedBootkeyFloppy                                               = 0xC00002E0
	StatusDsCantStart                                                        = 0xC00002E1
	StatusDsInitFailure                                                      = 0xC00002E2
	StatusSamInitFailure                                                     = 0xC00002E3
	StatusDsGcRequired                                                       = 0xC00002E4
	StatusDsLocalMemberOfLocalOnly                                           = 0xC00002E5
	StatusDsNoFpoInUniversalGroups                                           = 0xC00002E6
	StatusDsMachineAccountQuotaExceeded                                      = 0xC00002E7
	StatusMultipleFaultViolation                                             = 0xC00002E8
	StatusCurrentDomainNotAllowed                                            = 0xC00002E9
	StatusCannotMake                                                         = 0xC00002EA
	StatusSystemShutdown                                                     = 0xC00002EB
	StatusDsInitFailureConsole                                               = 0xC00002EC
	StatusDsSamInitFailureConsole                                            = 0xC00002ED
	StatusUnfinishedContextDeleted                                           = 0xC00002EE
	StatusNoTgtReply                                                         = 0xC00002EF
	StatusObjectidNotFound                                                   = 0xC00002F0
	StatusNoIpAddresses                                                      = 0xC00002F1
	StatusWrongCredentialHandle                                              = 0xC00002F2
	StatusCryptoSystemInvalid                                                = 0xC00002F3
	StatusMaxReferralsExceeded                                               = 0xC00002F4
	StatusMustBeKdc                                                          = 0xC00002F5
	StatusStrongCryptoNotSupported                                           = 0xC00002F6
	StatusTooManyPrincipals                                                  = 0xC00002F7
	StatusNoPaData                                                           = 0xC00002F8
	StatusPkinitNameMismatch                                                 = 0xC00002F9
	StatusSmartcardLogonRequired                                             = 0xC00002FA
	StatusKdcInvalidRequest                                                  = 0xC00002FB
	StatusKdcUnableToRefer                                                   = 0xC00002FC
	StatusKdcUnknownEtype                                                    = 0xC00002FD
	StatusShutdownInProgress                                                 = 0xC00002FE
	StatusServerShutdownInProgress                                           = 0xC00002FF
	StatusNotSupportedOnSbs                                                  = 0xC0000300
	StatusWmiGuidDisconnected                                                = 0xC0000301
	StatusWmiAlreadyDisabled                                                 = 0xC0000302
	StatusWmiAlreadyEnabled                                                  = 0xC0000303
	StatusMftTooFragmented                                                   = 0xC0000304
	StatusCopyProtectionFailure                                              = 0xC0000305
	StatusCssAuthenticationFailure                                           = 0xC0000306
	StatusCssKeyNotPresent                                                   = 0xC0000307
	StatusCssKeyNotEstablished                                               = 0xC0000308
	StatusCssScrambledSector                                                 = 0xC0000309
	StatusCssRegionMismatch                                                  = 0xC000030A
	StatusCssResetsExhausted                                                 = 0xC000030B
	StatusPkinitFailure                                                      = 0xC0000320
	StatusSmartcardSubsystemFailure                                          = 0xC0000321
	StatusNoKerbKey                                                          = 0xC0000322
	StatusHostDown                                                           = 0xC0000350
	StatusUnsupportedPreauth                                                 = 0xC0000351
	StatusEfsAlgBlobTooBig                                                   = 0xC0000352
	StatusPortNotSet                                                         = 0xC0000353
	StatusDebuggerInactive                                                   = 0xC0000354
	StatusDsVersionCheckFailure                                              = 0xC0000355
	StatusAuditingDisabled                                                   = 0xC0000356
	StatusPrent4MachineAccount                                               = 0xC0000357
	StatusDsAgCantHaveUniversalMember                                        = 0xC0000358
	StatusInvalidImageWin32                                                  = 0xC0000359
	StatusInvalidImageWin64                                                  = 0xC000035A
	StatusBadBindings                                                        = 0xC000035B
	StatusNetworkSessionExpired                                              = 0xC000035C
	StatusApphelpBlock                                                       = 0xC000035D
	StatusAllSidsFiltered                                                    = 0xC000035E
	StatusNotSafeModeDriver                                                  = 0xC000035F
	StatusAccessDisabledByPolicyDefault                                      = 0xC0000361
	StatusAccessDisabledByPolicyPath                                         = 0xC0000362
	StatusAccessDisabledByPolicyPublisher                                    = 0xC0000363
	StatusAccessDisabledByPolicyOther                                        = 0xC0000364
	StatusFailedDriverEntry                                                  = 0xC0000365
	StatusDeviceEnumerationError                                             = 0xC0000366
	StatusMountPointNotResolved                                              = 0xC0000368
	StatusInvalidDeviceObjectParameter                                       = 0xC0000369
	StatusMcaOccured                                                         = 0xC000036A
	StatusDriverBlockedCritical                                              = 0xC000036B
	StatusDriverBlocked                                                      = 0xC000036C
	StatusDriverDatabaseError                                                = 0xC000036D
	StatusSystemHiveTooLarge                                                 = 0xC000036E
	StatusInvalidImportOfNonDll                                              = 0xC000036F
	StatusDsShuttingDown                                                     = 0x40000370
	StatusNoSecrets                                                          = 0xC0000371
	StatusAccessDisabledNoSaferUiByPolicy                                    = 0xC0000372
	StatusFailedStackSwitch                                                  = 0xC0000373
	StatusHeapCorruption                                                     = 0xC0000374
	StatusSmartcardWrongPin                                                  = 0xC0000380
	StatusSmartcardCardBlocked                                               = 0xC0000381
	StatusSmartcardCardNotAuthenticated                                      = 0xC0000382
	StatusSmartcardNoCard                                                    = 0xC0000383
	StatusSmartcardNoKeyContainer                                            = 0xC0000384
	StatusSmartcardNoCertificate                                             = 0xC0000385
	StatusSmartcardNoKeyset                                                  = 0xC0000386
	StatusSmartcardIoError                                                   = 0xC0000387
	StatusDowngradeDetected                                                  = 0xC0000388
	StatusSmartcardCertRevoked                                               = 0xC0000389
	StatusIssuingCaUntrusted                                                 = 0xC000038A
	StatusRevocationOfflineC                                                 = 0xC000038B
	StatusPkinitClientFailure                                                = 0xC000038C
	StatusSmartcardCertExpired                                               = 0xC000038D
	StatusDriverFailedPriorUnload                                            = 0xC000038E
	StatusSmartcardSilentContext                                             = 0xC000038F
	StatusPerUserTrustQuotaExceeded                                          = 0xC0000401
	StatusAllUserTrustQuotaExceeded                                          = 0xC0000402
	StatusUserDeleteTrustQuotaExceeded                                       = 0xC0000403
	StatusDsNameNotUnique                                                    = 0xC0000404
	StatusDsDuplicateIdFound                                                 = 0xC0000405
	StatusDsGroupConversionError                                             = 0xC0000406
	StatusVolsnapPrepareHibernate                                            = 0xC0000407
	StatusUser2UserRequired                                                  = 0xC0000408
	StatusStackBufferOverrun                                                 = 0xC0000409
	StatusNoS4UProtSupport                                                   = 0xC000040A
	StatusCrossrealmDelegationFailure                                        = 0xC000040B
	StatusRevocationOfflineKdc                                               = 0xC000040C
	StatusIssuingCaUntrustedKdc                                              = 0xC000040D
	StatusKdcCertExpired                                                     = 0xC000040E
	StatusKdcCertRevoked                                                     = 0xC000040F
	StatusParameterQuotaExceeded                                             = 0xC0000410
	StatusHibernationFailure                                                 = 0xC0000411
	StatusDelayLoadFailed                                                    = 0xC0000412
	StatusAuthenticationFirewallFailed                                       = 0xC0000413
	StatusVdmDisallowed                                                      = 0xC0000414
	StatusHungDisplayDriverThread                                            = 0xC0000415
	StatusInsufficientResourceForSpecifiedSharedSectionSize                  = 0xC0000416
	StatusInvalidCruntimeParameter                                           = 0xC0000417
	StatusNtlmBlocked                                                        = 0xC0000418
	StatusDsSrcSidExistsInForest                                             = 0xC0000419
	StatusDsDomainNameExistsInForest                                         = 0xC000041A
	StatusDsFlatNameExistsInForest                                           = 0xC000041B
	StatusInvalidUserPrincipalName                                           = 0xC000041C
	StatusFatalUserCallbackException                                         = 0xC000041D
	StatusAssertionFailure                                                   = 0xC0000420
	StatusVerifierStop                                                       = 0xC0000421
	StatusCallbackPopStack                                                   = 0xC0000423
	StatusIncompatibleDriverBlocked                                          = 0xC0000424
	StatusHiveUnloaded                                                       = 0xC0000425
	StatusCompressionDisabled                                                = 0xC0000426
	StatusFileSystemLimitation                                               = 0xC0000427
	StatusInvalidImageHash                                                   = 0xC0000428
	StatusNotCapable                                                         = 0xC0000429
	StatusRequestOutOfSequence                                               = 0xC000042A
	StatusImplementationLimit                                                = 0xC000042B
	StatusElevationRequired                                                  = 0xC000042C
	StatusNoSecurityContext                                                  = 0xC000042D
	StatusPku2UCertFailure                                                   = 0xC000042F
	StatusBeyondVdl                                                          = 0xC0000432
	StatusEncounteredWriteInProgress                                         = 0xC0000433
	StatusPteChanged                                                         = 0xC0000434
	StatusPurgeFailed                                                        = 0xC0000435
	StatusCredRequiresConfirmation                                           = 0xC0000440
	StatusCsEncryptionInvalidServerResponse                                  = 0xC0000441
	StatusCsEncryptionUnsupportedServer                                      = 0xC0000442
	StatusCsEncryptionExistingEncryptedFile                                  = 0xC0000443
	StatusCsEncryptionNewEncryptedFile                                       = 0xC0000444
	StatusCsEncryptionFileNotCse                                             = 0xC0000445
	StatusInvalidLabel                                                       = 0xC0000446
	StatusDriverProcessTerminated                                            = 0xC0000450
	StatusAmbiguousSystemDevice                                              = 0xC0000451
	StatusSystemDeviceNotFound                                               = 0xC0000452
	StatusRestartBootApplication                                             = 0xC0000453
	StatusInsufficientNvramResources                                         = 0xC0000454
	StatusInvalidTaskName                                                    = 0xC0000500
	StatusInvalidTaskIndex                                                   = 0xC0000501
	StatusThreadAlreadyInTask                                                = 0xC0000502
	StatusCallbackBypass                                                     = 0xC0000503
	StatusFailFastException                                                  = 0xC0000602
	StatusImageCertRevoked                                                   = 0xC0000603
	StatusPortClosed                                                         = 0xC0000700
	StatusMessageLost                                                        = 0xC0000701
	StatusInvalidMessage                                                     = 0xC0000702
	StatusRequestCanceled                                                    = 0xC0000703
	StatusRecursiveDispatch                                                  = 0xC0000704
	StatusLpcReceiveBufferExpected                                           = 0xC0000705
	StatusLpcInvalidConnectionUsage                                          = 0xC0000706
	StatusLpcRequestsNotAllowed                                              = 0xC0000707
	StatusResourceInUse                                                      = 0xC0000708
	StatusHardwareMemoryError                                                = 0xC0000709
	StatusThreadpoolHandleException                                          = 0xC000070A
	StatusThreadpoolSetEventOnCompletionFailed                               = 0xC000070B
	StatusThreadpoolReleaseSemaphoreOnCompletionFailed                       = 0xC000070C
	StatusThreadpoolReleaseMutexOnCompletionFailed                           = 0xC000070D
	StatusThreadpoolFreeLibraryOnCompletionFailed                            = 0xC000070E
	StatusThreadpoolReleasedDuringOperation                                  = 0xC000070F
	StatusCallbackReturnedWhileImpersonating                                 = 0xC0000710
	StatusApcReturnedWhileImpersonating                                      = 0xC0000711
	StatusProcessIsProtected                                                 = 0xC0000712
	StatusMcaException                                                       = 0xC0000713
	StatusCertificateMappingNotUnique                                        = 0xC0000714
	StatusSymlinkClassDisabled                                               = 0xC0000715
	StatusInvalidIdnNormalization                                            = 0xC0000716
	StatusNoUnicodeTranslation                                               = 0xC0000717
	StatusAlreadyRegistered                                                  = 0xC0000718
	StatusContextMismatch                                                    = 0xC0000719
	StatusPortAlreadyHasCompletionList                                       = 0xC000071A
	StatusCallbackReturnedThreadPriority                                     = 0xC000071B
	StatusInvalidThread                                                      = 0xC000071C
	StatusCallbackReturnedTransaction                                        = 0xC000071D
	StatusCallbackReturnedLdrLock                                            = 0xC000071E
	StatusCallbackReturnedLang                                               = 0xC000071F
	StatusCallbackReturnedPriBack                                            = 0xC0000720
	StatusCallbackReturnedThreadAffinity                                     = 0xC0000721
	StatusDiskRepairDisabled                                                 = 0xC0000800
	StatusDsDomainRenameInProgress                                           = 0xC0000801
	StatusDiskQuotaExceeded                                                  = 0xC0000802
	StatusDataLostRepair                                                     = 0x80000803
	StatusContentBlocked                                                     = 0xC0000804
	StatusBadClusters                                                        = 0xC0000805
	StatusVolumeDirty                                                        = 0xC0000806
	StatusFileCheckedOut                                                     = 0xC0000901
	StatusCheckoutRequired                                                   = 0xC0000902
	StatusBadFileType                                                        = 0xC0000903
	StatusFileTooLarge                                                       = 0xC0000904
	StatusFormsAuthRequired                                                  = 0xC0000905
	StatusVirusInfected                                                      = 0xC0000906
	StatusVirusDeleted                                                       = 0xC0000907
	StatusBadMcfgTable                                                       = 0xC0000908
	StatusCannotBreakOplock                                                  = 0xC0000909
	StatusWowAssertion                                                       = 0xC0009898
	StatusInvalidSignature                                                   = 0xC000A000
	StatusHmacNotSupported                                                   = 0xC000A001
	StatusAuthTagMismatch                                                    = 0xC000A002
	StatusIpsecQueueOverflow                                                 = 0xC000A010
	StatusNdQueueOverflow                                                    = 0xC000A011
	StatusHoplimitExceeded                                                   = 0xC000A012
	StatusProtocolNotSupported                                               = 0xC000A013
	StatusFastpathRejected                                                   = 0xC000A014
	StatusLostWritebehindDataNetworkDisconnected                             = 0xC000A080
	StatusLostWritebehindDataNetworkServerError                              = 0xC000A081
	StatusLostWritebehindDataLocalDiskError                                  = 0xC000A082
	StatusXmlParseError                                                      = 0xC000A083
	StatusXmldsigError                                                       = 0xC000A084
	StatusWrongCompartment                                                   = 0xC000A085
	StatusAuthipFailure                                                      = 0xC000A086
	StatusDsOidMappedGroupCantHaveMembers                                    = 0xC000A087
	StatusDsOidNotFound                                                      = 0xC000A088
	StatusHashNotSupported                                                   = 0xC000A100
	StatusHashNotPresent                                                     = 0xC000A101
	DbgNoStateChange                                                         = 0xC0010001
	DbgAppNotIdle                                                            = 0xC0010002
	RpcNtInvalidStringBinding                                                = 0xC0020001
	RpcNtWrongKindOfBinding                                                  = 0xC0020002
	RpcNtInvalidBinding                                                      = 0xC0020003
	RpcNtProtseqNotSupported                                                 = 0xC0020004
	RpcNtInvalidRpcProtseq                                                   = 0xC0020005
	RpcNtInvalidStringUuid                                                   = 0xC0020006
	RpcNtInvalidEndpointFormat                                               = 0xC0020007
	RpcNtInvalidNetAddr                                                      = 0xC0020008
	RpcNtNoEndpointFound                                                     = 0xC0020009
	RpcNtInvalidTimeout                                                      = 0xC002000A
	RpcNtObjectNotFound                                                      = 0xC002000B
	RpcNtAlreadyRegistered                                                   = 0xC002000C
	RpcNtTypeAlreadyRegistered                                               = 0xC002000D
	RpcNtAlreadyListening                                                    = 0xC002000E
	RpcNtNoProtseqsRegistered                                                = 0xC002000F
	RpcNtNotListening                                                        = 0xC0020010
	RpcNtUnknownMgrType                                                      = 0xC0020011
	RpcNtUnknownIf                                                           = 0xC0020012
	RpcNtNoBindings                                                          = 0xC0020013
	RpcNtNoProtseqs                                                          = 0xC0020014
	RpcNtCantCreateEndpoint                                                  = 0xC0020015
	RpcNtOutOfResources                                                      = 0xC0020016
	RpcNtServerUnavailable                                                   = 0xC0020017
	RpcNtServerTooBusy                                                       = 0xC0020018
	RpcNtInvalidNetworkOptions                                               = 0xC0020019
	RpcNtNoCallActive                                                        = 0xC002001A
	RpcNtCallFailed                                                          = 0xC002001B
	RpcNtCallFailedDne                                                       = 0xC002001C
	RpcNtProtocolError                                                       = 0xC002001D
	RpcNtUnsupportedTransSyn                                                 = 0xC002001F
	RpcNtUnsupportedType                                                     = 0xC0020021
	RpcNtInvalidTag                                                          = 0xC0020022
	RpcNtInvalidBound                                                        = 0xC0020023
	RpcNtNoEntryName                                                         = 0xC0020024
	RpcNtInvalidNameSyntax                                                   = 0xC0020025
	RpcNtUnsupportedNameSyntax                                               = 0xC0020026
	RpcNtUuidNoAddress                                                       = 0xC0020028
	RpcNtDuplicateEndpoint                                                   = 0xC0020029
	RpcNtUnknownAuthnType                                                    = 0xC002002A
	RpcNtMaxCallsTooSmall                                                    = 0xC002002B
	RpcNtStringTooLong                                                       = 0xC002002C
	RpcNtProtseqNotFound                                                     = 0xC002002D
	RpcNtProcnumOutOfRange                                                   = 0xC002002E
	RpcNtBindingHasNoAuth                                                    = 0xC002002F
	RpcNtUnknownAuthnService                                                 = 0xC0020030
	RpcNtUnknownAuthnLevel                                                   = 0xC0020031
	RpcNtInvalidAuthIdentity                                                 = 0xC0020032
	RpcNtUnknownAuthzService                                                 = 0xC0020033
	EptNtInvalidEntry                                                        = 0xC0020034
	EptNtCantPerformOp                                                       = 0xC0020035
	EptNtNotRegistered                                                       = 0xC0020036
	RpcNtNothingToExport                                                     = 0xC0020037
	RpcNtIncompleteName                                                      = 0xC0020038
	RpcNtInvalidVersOption                                                   = 0xC0020039
	RpcNtNoMoreMembers                                                       = 0xC002003A
	RpcNtNotAllObjsUnexported                                                = 0xC002003B
	RpcNtInterfaceNotFound                                                   = 0xC002003C
	RpcNtEntryAlreadyExists                                                  = 0xC002003D
	RpcNtEntryNotFound                                                       = 0xC002003E
	RpcNtNameServiceUnavailable                                              = 0xC002003F
	RpcNtInvalidNafId                                                        = 0xC0020040
	RpcNtCannotSupport                                                       = 0xC0020041
	RpcNtNoContextAvailable                                                  = 0xC0020042
	RpcNtInternalError                                                       = 0xC0020043
	RpcNtZeroDivide                                                          = 0xC0020044
	RpcNtAddressError                                                        = 0xC0020045
	RpcNtFpDivZero                                                           = 0xC0020046
	RpcNtFpUnderflow                                                         = 0xC0020047
	RpcNtFpOverflow                                                          = 0xC0020048
	RpcNtNoMoreEntries                                                       = 0xC0030001
	RpcNtSsCharTransOpenFail                                                 = 0xC0030002
	RpcNtSsCharTransShortFile                                                = 0xC0030003
	RpcNtSsInNullContext                                                     = 0xC0030004
	RpcNtSsContextMismatch                                                   = 0xC0030005
	RpcNtSsContextDamaged                                                    = 0xC0030006
	RpcNtSsHandlesMismatch                                                   = 0xC0030007
	RpcNtSsCannotGetCallHandle                                               = 0xC0030008
	RpcNtNullRefPointer                                                      = 0xC0030009
	RpcNtEnumValueOutOfRange                                                 = 0xC003000A
	RpcNtByteCountTooSmall                                                   = 0xC003000B
	RpcNtBadStubData                                                         = 0xC003000C
	RpcNtCallInProgress                                                      = 0xC0020049
	RpcNtNoMoreBindings                                                      = 0xC002004A
	RpcNtGroupMemberNotFound                                                 = 0xC002004B
	EptNtCantCreate                                                          = 0xC002004C
	RpcNtInvalidObject                                                       = 0xC002004D
	RpcNtNoInterfaces                                                        = 0xC002004F
	RpcNtCallCancelled                                                       = 0xC0020050
	RpcNtBindingIncomplete                                                   = 0xC0020051
	RpcNtCommFailure                                                         = 0xC0020052
	RpcNtUnsupportedAuthnLevel                                               = 0xC0020053
	RpcNtNoPrincName                                                         = 0xC0020054
	RpcNtNotRpcError                                                         = 0xC0020055
	RpcNtUuidLocalOnly                                                       = 0x40020056
	RpcNtSecPkgError                                                         = 0xC0020057
	RpcNtNotCancelled                                                        = 0xC0020058
	RpcNtInvalidEsAction                                                     = 0xC0030059
	RpcNtWrongEsVersion                                                      = 0xC003005A
	RpcNtWrongStubVersion                                                    = 0xC003005B
	RpcNtInvalidPipeObject                                                   = 0xC003005C
	RpcNtInvalidPipeOperation                                                = 0xC003005D
	RpcNtWrongPipeVersion                                                    = 0xC003005E
	RpcNtPipeClosed                                                          = 0xC003005F
	RpcNtPipeDisciplineError                                                 = 0xC0030060
	RpcNtPipeEmpty                                                           = 0xC0030061
	RpcNtInvalidAsyncHandle                                                  = 0xC0020062
	RpcNtInvalidAsyncCall                                                    = 0xC0020063
	RpcNtProxyAccessDenied                                                   = 0xC0020064
	RpcNtCookieAuthFailed                                                    = 0xC0020065
	RpcNtSendIncomplete                                                      = 0x400200AF
	StatusAcpiInvalidOpcode                                                  = 0xC0140001
	StatusAcpiStackOverflow                                                  = 0xC0140002
	StatusAcpiAssertFailed                                                   = 0xC0140003
	StatusAcpiInvalidIndex                                                   = 0xC0140004
	StatusAcpiInvalidArgument                                                = 0xC0140005
	StatusAcpiFatal                                                          = 0xC0140006
	StatusAcpiInvalidSupername                                               = 0xC0140007
	StatusAcpiInvalidArgtype                                                 = 0xC0140008
	StatusAcpiInvalidObjtype                                                 = 0xC0140009
	StatusAcpiInvalidTargettype                                              = 0xC014000A
	StatusAcpiIncorrectArgumentCount                                         = 0xC014000B
	StatusAcpiAddressNotMapped                                               = 0xC014000C
	StatusAcpiInvalidEventtype                                               = 0xC014000D
	StatusAcpiHandlerCollision                                               = 0xC014000E
	StatusAcpiInvalidData                                                    = 0xC014000F
	StatusAcpiInvalidRegion                                                  = 0xC0140010
	StatusAcpiInvalidAccessSize                                              = 0xC0140011
	StatusAcpiAcquireGlobalLock                                              = 0xC0140012
	StatusAcpiAlreadyInitialized                                             = 0xC0140013
	StatusAcpiNotInitialized                                                 = 0xC0140014
	StatusAcpiInvalidMutexLevel                                              = 0xC0140015
	StatusAcpiMutexNotOwned                                                  = 0xC0140016
	StatusAcpiMutexNotOwner                                                  = 0xC0140017
	StatusAcpiRsAccess                                                       = 0xC0140018
	StatusAcpiInvalidTable                                                   = 0xC0140019
	StatusAcpiRegHandlerFailed                                               = 0xC0140020
	StatusAcpiPowerRequestFailed                                             = 0xC0140021
	StatusCtxWinstationNameInvalid                                           = 0xC00A0001
	StatusCtxInvalidPd                                                       = 0xC00A0002
	StatusCtxPdNotFound                                                      = 0xC00A0003
	StatusCtxCdmConnect                                                      = 0x400A0004
	StatusCtxCdmDisconnect                                                   = 0x400A0005
	StatusCtxClosePending                                                    = 0xC00A0006
	StatusCtxNoOutbuf                                                        = 0xC00A0007
	StatusCtxModemInfNotFound                                                = 0xC00A0008
	StatusCtxInvalidModemname                                                = 0xC00A0009
	StatusCtxResponseError                                                   = 0xC00A000A
	StatusCtxModemResponseTimeout                                            = 0xC00A000B
	StatusCtxModemResponseNoCarrier                                          = 0xC00A000C
	StatusCtxModemResponseNoDialtone                                         = 0xC00A000D
	StatusCtxModemResponseBusy                                               = 0xC00A000E
	StatusCtxModemResponseVoice                                              = 0xC00A000F
	StatusCtxTdError                                                         = 0xC00A0010
	StatusCtxLicenseClientInvalid                                            = 0xC00A0012
	StatusCtxLicenseNotAvailable                                             = 0xC00A0013
	StatusCtxLicenseExpired                                                  = 0xC00A0014
	StatusCtxWinstationNotFound                                              = 0xC00A0015
	StatusCtxWinstationNameCollision                                         = 0xC00A0016
	StatusCtxWinstationBusy                                                  = 0xC00A0017
	StatusCtxBadVideoMode                                                    = 0xC00A0018
	StatusCtxGraphicsInvalid                                                 = 0xC00A0022
	StatusCtxNotConsole                                                      = 0xC00A0024
	StatusCtxClientQueryTimeout                                              = 0xC00A0026
	StatusCtxConsoleDisconnect                                               = 0xC00A0027
	StatusCtxConsoleConnect                                                  = 0xC00A0028
	StatusCtxShadowDenied                                                    = 0xC00A002A
	StatusCtxWinstationAccessDenied                                          = 0xC00A002B
	StatusCtxInvalidWd                                                       = 0xC00A002E
	StatusCtxWdNotFound                                                      = 0xC00A002F
	StatusCtxShadowInvalid                                                   = 0xC00A0030
	StatusCtxShadowDisabled                                                  = 0xC00A0031
	StatusRdpProtocolError                                                   = 0xC00A0032
	StatusCtxClientLicenseNotSet                                             = 0xC00A0033
	StatusCtxClientLicenseInUse                                              = 0xC00A0034
	StatusCtxShadowEndedByModeChange                                         = 0xC00A0035
	StatusCtxShadowNotRunning                                                = 0xC00A0036
	StatusCtxLogonDisabled                                                   = 0xC00A0037
	StatusCtxSecurityLayerError                                              = 0xC00A0038
	StatusTsIncompatibleSessions                                             = 0xC00A0039
	StatusTsVideoSubsystemError                                              = 0xC00A003A
	StatusPnpBadMpsTable                                                     = 0xC0040035
	StatusPnpTranslationFailed                                               = 0xC0040036
	StatusPnpIrqTranslationFailed                                            = 0xC0040037
	StatusPnpInvalidId                                                       = 0xC0040038
	StatusIoReissueAsCached                                                  = 0xC0040039
	StatusMuiFileNotFound                                                    = 0xC00B0001
	StatusMuiInvalidFile                                                     = 0xC00B0002
	StatusMuiInvalidRcConfig                                                 = 0xC00B0003
	StatusMuiInvalidLocaleName                                               = 0xC00B0004
	StatusMuiInvalidUltimatefallbackName                                     = 0xC00B0005
	StatusMuiFileNotLoaded                                                   = 0xC00B0006
	StatusResourceEnumUserStop                                               = 0xC00B0007
	StatusFltNoHandlerDefined                                                = 0xC01C0001
	StatusFltContextAlreadyDefined                                           = 0xC01C0002
	StatusFltInvalidAsynchronousRequest                                      = 0xC01C0003
	StatusFltDisallowFastIo                                                  = 0xC01C0004
	StatusFltInvalidNameRequest                                              = 0xC01C0005
	StatusFltNotSafeToPostOperation                                          = 0xC01C0006
	StatusFltNotInitialized                                                  = 0xC01C0007
	StatusFltFilterNotReady                                                  = 0xC01C0008
	StatusFltPostOperationCleanup                                            = 0xC01C0009
	StatusFltInternalError                                                   = 0xC01C000A
	StatusFltDeletingObject                                                  = 0xC01C000B
	StatusFltMustBeNonpagedPool                                              = 0xC01C000C
	StatusFltDuplicateEntry                                                  = 0xC01C000D
	StatusFltCbdqDisabled                                                    = 0xC01C000E
	StatusFltDoNotAttach                                                     = 0xC01C000F
	StatusFltDoNotDetach                                                     = 0xC01C0010
	StatusFltInstanceAltitudeCollision                                       = 0xC01C0011
	StatusFltInstanceNameCollision                                           = 0xC01C0012
	StatusFltFilterNotFound                                                  = 0xC01C0013
	StatusFltVolumeNotFound                                                  = 0xC01C0014
	StatusFltInstanceNotFound                                                = 0xC01C0015
	StatusFltContextAllocationNotFound                                       = 0xC01C0016
	StatusFltInvalidContextRegistration                                      = 0xC01C0017
	StatusFltNameCacheMiss                                                   = 0xC01C0018
	StatusFltNoDeviceObject                                                  = 0xC01C0019
	StatusFltVolumeAlreadyMounted                                            = 0xC01C001A
	StatusFltAlreadyEnlisted                                                 = 0xC01C001B
	StatusFltContextAlreadyLinked                                            = 0xC01C001C
	StatusFltNoWaiterForReply                                                = 0xC01C0020
	StatusSxsSectionNotFound                                                 = 0xC0150001
	StatusSxsCantGenActctx                                                   = 0xC0150002
	StatusSxsInvalidActctxdataFormat                                         = 0xC0150003
	StatusSxsAssemblyNotFound                                                = 0xC0150004
	StatusSxsManifestFormatError                                             = 0xC0150005
	StatusSxsManifestParseError                                              = 0xC0150006
	StatusSxsActivationContextDisabled                                       = 0xC0150007
	StatusSxsKeyNotFound                                                     = 0xC0150008
	StatusSxsVersionConflict                                                 = 0xC0150009
	StatusSxsWrongSectionType                                                = 0xC015000A
	StatusSxsThreadQueriesDisabled                                           = 0xC015000B
	StatusSxsAssemblyMissing                                                 = 0xC015000C
	StatusSxsReleaseActivationContext                                        = 0x4015000D
	StatusSxsProcessDefaultAlreadySet                                        = 0xC015000E
	StatusSxsEarlyDeactivation                                               = 0xC015000F
	StatusSxsInvalidDeactivation                                             = 0xC0150010
	StatusSxsMultipleDeactivation                                            = 0xC0150011
	StatusSxsSystemDefaultActivationContextEmpty                             = 0xC0150012
	StatusSxsProcessTerminationRequested                                     = 0xC0150013
	StatusSxsCorruptActivationStack                                          = 0xC0150014
	StatusSxsCorruption                                                      = 0xC0150015
	StatusSxsInvalidIdentityAttributeValue                                   = 0xC0150016
	StatusSxsInvalidIdentityAttributeName                                    = 0xC0150017
	StatusSxsIdentityDuplicateAttribute                                      = 0xC0150018
	StatusSxsIdentityParseError                                              = 0xC0150019
	StatusSxsComponentStoreCorrupt                                           = 0xC015001A
	StatusSxsFileHashMismatch                                                = 0xC015001B
	StatusSxsManifestIdentitySameButContentsDifferent                        = 0xC015001C
	StatusSxsIdentitiesDifferent                                             = 0xC015001D
	StatusSxsAssemblyIsNotADeployment                                        = 0xC015001E
	StatusSxsFileNotPartOfAssembly                                           = 0xC015001F
	StatusAdvancedInstallerFailed                                            = 0xC0150020
	StatusXmlEncodingMismatch                                                = 0xC0150021
	StatusSxsManifestTooBig                                                  = 0xC0150022
	StatusSxsSettingNotRegistered                                            = 0xC0150023
	StatusSxsTransactionClosureIncomplete                                    = 0xC0150024
	StatusSmiPrimitiveInstallerFailed                                        = 0xC0150025
	StatusGenericCommandFailed                                               = 0xC0150026
	StatusSxsFileHashMissing                                                 = 0xC0150027
	StatusClusterInvalidNode                                                 = 0xC0130001
	StatusClusterNodeExists                                                  = 0xC0130002
	StatusClusterJoinInProgress                                              = 0xC0130003
	StatusClusterNodeNotFound                                                = 0xC0130004
	StatusClusterLocalNodeNotFound                                           = 0xC0130005
	StatusClusterNetworkExists                                               = 0xC0130006
	StatusClusterNetworkNotFound                                             = 0xC0130007
	StatusClusterNetinterfaceExists                                          = 0xC0130008
	StatusClusterNetinterfaceNotFound                                        = 0xC0130009
	StatusClusterInvalidRequest                                              = 0xC013000A
	StatusClusterInvalidNetworkProvider                                      = 0xC013000B
	StatusClusterNodeDown                                                    = 0xC013000C
	StatusClusterNodeUnreachable                                             = 0xC013000D
	StatusClusterNodeNotMember                                               = 0xC013000E
	StatusClusterJoinNotInProgress                                           = 0xC013000F
	StatusClusterInvalidNetwork                                              = 0xC0130010
	StatusClusterNoNetAdapters                                               = 0xC0130011
	StatusClusterNodeUp                                                      = 0xC0130012
	StatusClusterNodePaused                                                  = 0xC0130013
	StatusClusterNodeNotPaused                                               = 0xC0130014
	StatusClusterNoSecurityContext                                           = 0xC0130015
	StatusClusterNetworkNotInternal                                          = 0xC0130016
	StatusClusterPoisoned                                                    = 0xC0130017
	StatusClusterNonCsvPath                                                  = 0xC0130018
	StatusClusterCsvVolumeNotLocal                                           = 0xC0130019
	StatusTransactionalConflict                                              = 0xC0190001
	StatusInvalidTransaction                                                 = 0xC0190002
	StatusTransactionNotActive                                               = 0xC0190003
	StatusTmInitializationFailed                                             = 0xC0190004
	StatusRmNotActive                                                        = 0xC0190005
	StatusRmMetadataCorrupt                                                  = 0xC0190006
	StatusTransactionNotJoined                                               = 0xC0190007
	StatusDirectoryNotRm                                                     = 0xC0190008
	StatusCouldNotResizeLog                                                  = 0x80190009
	StatusTransactionsUnsupportedRemote                                      = 0xC019000A
	StatusLogResizeInvalidSize                                               = 0xC019000B
	StatusRemoteFileVersionMismatch                                          = 0xC019000C
	StatusCrmProtocolAlreadyExists                                           = 0xC019000F
	StatusTransactionPropagationFailed                                       = 0xC0190010
	StatusCrmProtocolNotFound                                                = 0xC0190011
	StatusTransactionSuperiorExists                                          = 0xC0190012
	StatusTransactionRequestNotValid                                         = 0xC0190013
	StatusTransactionNotRequested                                            = 0xC0190014
	StatusTransactionAlreadyAborted                                          = 0xC0190015
	StatusTransactionAlreadyCommitted                                        = 0xC0190016
	StatusTransactionInvalidMarshallBuffer                                   = 0xC0190017
	StatusCurrentTransactionNotValid                                         = 0xC0190018
	StatusLogGrowthFailed                                                    = 0xC0190019
	StatusObjectNoLongerExists                                               = 0xC0190021
	StatusStreamMiniversionNotFound                                          = 0xC0190022
	StatusStreamMiniversionNotValid                                          = 0xC0190023
	StatusMiniversionInaccessibleFromSpecifiedTransaction                    = 0xC0190024
	StatusCantOpenMiniversionWithModifyIntent                                = 0xC0190025
	StatusCantCreateMoreStreamMiniversions                                   = 0xC0190026
	StatusHandleNoLongerValid                                                = 0xC0190028
	StatusNoTxfMetadata                                                      = 0x80190029
	StatusLogCorruptionDetected                                              = 0xC0190030
	StatusCantRecoverWithHandleOpen                                          = 0x80190031
	StatusRmDisconnected                                                     = 0xC0190032
	StatusEnlistmentNotSuperior                                              = 0xC0190033
	StatusRecoveryNotNeeded                                                  = 0x40190034
	StatusRmAlreadyStarted                                                   = 0x40190035
	StatusFileIdentityNotPersistent                                          = 0xC0190036
	StatusCantBreakTransactionalDependency                                   = 0xC0190037
	StatusCantCrossRmBoundary                                                = 0xC0190038
	StatusTxfDirNotEmpty                                                     = 0xC0190039
	StatusIndoubtTransactionsExist                                           = 0xC019003A
	StatusTmVolatile                                                         = 0xC019003B
	StatusRollbackTimerExpired                                               = 0xC019003C
	StatusTxfAttributeCorrupt                                                = 0xC019003D
	StatusEfsNotAllowedInTransaction                                         = 0xC019003E
	StatusTransactionalOpenNotAllowed                                        = 0xC019003F
	StatusTransactedMappingUnsupportedRemote                                 = 0xC0190040
	StatusTxfMetadataAlreadyPresent                                          = 0x80190041
	StatusTransactionScopeCallbacksNotSet                                    = 0x80190042
	StatusTransactionRequiredPromotion                                       = 0xC0190043
	StatusCannotExecuteFileInTransaction                                     = 0xC0190044
	StatusTransactionsNotFrozen                                              = 0xC0190045
	StatusTransactionFreezeInProgress                                        = 0xC0190046
	StatusNotSnapshotVolume                                                  = 0xC0190047
	StatusNoSavepointWithOpenFiles                                           = 0xC0190048
	StatusSparseNotAllowedInTransaction                                      = 0xC0190049
	StatusTmIdentityMismatch                                                 = 0xC019004A
	StatusFloatedSection                                                     = 0xC019004B
	StatusCannotAcceptTransactedWork                                         = 0xC019004C
	StatusCannotAbortTransactions                                            = 0xC019004D
	StatusTransactionNotFound                                                = 0xC019004E
	StatusResourcemanagerNotFound                                            = 0xC019004F
	StatusEnlistmentNotFound                                                 = 0xC0190050
	StatusTransactionmanagerNotFound                                         = 0xC0190051
	StatusTransactionmanagerNotOnline                                        = 0xC0190052
	StatusTransactionmanagerRecoveryNameCollision                            = 0xC0190053
	StatusTransactionNotRoot                                                 = 0xC0190054
	StatusTransactionObjectExpired                                           = 0xC0190055
	StatusCompressionNotAllowedInTransaction                                 = 0xC0190056
	StatusTransactionResponseNotEnlisted                                     = 0xC0190057
	StatusTransactionRecordTooLong                                           = 0xC0190058
	StatusNoLinkTrackingInTransaction                                        = 0xC0190059
	StatusOperationNotSupportedInTransaction                                 = 0xC019005A
	StatusTransactionIntegrityViolated                                       = 0xC019005B
	StatusTransactionmanagerIdentityMismatch                                 = 0xC019005C
	StatusRmCannotBeFrozenForSnapshot                                        = 0xC019005D
	StatusTransactionMustWritethrough                                        = 0xC019005E
	StatusTransactionNoSuperior                                              = 0xC019005F
	StatusExpiredHandle                                                      = 0xC0190060
	StatusTransactionNotEnlisted                                             = 0xC0190061
	StatusLogSectorInvalid                                                   = 0xC01A0001
	StatusLogSectorParityInvalid                                             = 0xC01A0002
	StatusLogSectorRemapped                                                  = 0xC01A0003
	StatusLogBlockIncomplete                                                 = 0xC01A0004
	StatusLogInvalidRange                                                    = 0xC01A0005
	StatusLogBlocksExhausted                                                 = 0xC01A0006
	StatusLogReadContextInvalid                                              = 0xC01A0007
	StatusLogRestartInvalid                                                  = 0xC01A0008
	StatusLogBlockVersion                                                    = 0xC01A0009
	StatusLogBlockInvalid                                                    = 0xC01A000A
	StatusLogReadModeInvalid                                                 = 0xC01A000B
	StatusLogNoRestart                                                       = 0x401A000C
	StatusLogMetadataCorrupt                                                 = 0xC01A000D
	StatusLogMetadataInvalid                                                 = 0xC01A000E
	StatusLogMetadataInconsistent                                            = 0xC01A000F
	StatusLogReservationInvalid                                              = 0xC01A0010
	StatusLogCantDelete                                                      = 0xC01A0011
	StatusLogContainerLimitExceeded                                          = 0xC01A0012
	StatusLogStartOfLog                                                      = 0xC01A0013
	StatusLogPolicyAlreadyInstalled                                          = 0xC01A0014
	StatusLogPolicyNotInstalled                                              = 0xC01A0015
	StatusLogPolicyInvalid                                                   = 0xC01A0016
	StatusLogPolicyConflict                                                  = 0xC01A0017
	StatusLogPinnedArchiveTail                                               = 0xC01A0018
	StatusLogRecordNonexistent                                               = 0xC01A0019
	StatusLogRecordsReservedInvalid                                          = 0xC01A001A
	StatusLogSpaceReservedInvalid                                            = 0xC01A001B
	StatusLogTailInvalid                                                     = 0xC01A001C
	StatusLogFull                                                            = 0xC01A001D
	StatusLogMultiplexed                                                     = 0xC01A001E
	StatusLogDedicated                                                       = 0xC01A001F
	StatusLogArchiveNotInProgress                                            = 0xC01A0020
	StatusLogArchiveInProgress                                               = 0xC01A0021
	StatusLogEphemeral                                                       = 0xC01A0022
	StatusLogNotEnoughContainers                                             = 0xC01A0023
	StatusLogClientAlreadyRegistered                                         = 0xC01A0024
	StatusLogClientNotRegistered                                             = 0xC01A0025
	StatusLogFullHandlerInProgress                                           = 0xC01A0026
	StatusLogContainerReadFailed                                             = 0xC01A0027
	StatusLogContainerWriteFailed                                            = 0xC01A0028
	StatusLogContainerOpenFailed                                             = 0xC01A0029
	StatusLogContainerStateInvalid                                           = 0xC01A002A
	StatusLogStateInvalid                                                    = 0xC01A002B
	StatusLogPinned                                                          = 0xC01A002C
	StatusLogMetadataFlushFailed                                             = 0xC01A002D
	StatusLogInconsistentSecurity                                            = 0xC01A002E
	StatusLogAppendedFlushFailed                                             = 0xC01A002F
	StatusLogPinnedReservation                                               = 0xC01A0030
	StatusVideoHungDisplayDriverThread                                       = 0xC01B00EA
	StatusVideoHungDisplayDriverThreadRecovered                              = 0x801B00EB
	StatusVideoDriverDebugReportRequest                                      = 0x401B00EC
	StatusMonitorNoDescriptor                                                = 0xC01D0001
	StatusMonitorUnknownDescriptorFormat                                     = 0xC01D0002
	StatusMonitorInvalidDescriptorChecksum                                   = 0xC01D0003
	StatusMonitorInvalidStandardTimingBlock                                  = 0xC01D0004
	StatusMonitorWmiDatablockRegistrationFailed                              = 0xC01D0005
	StatusMonitorInvalidSerialNumberMondscBlock                              = 0xC01D0006
	StatusMonitorInvalidUserFriendlyMondscBlock                              = 0xC01D0007
	StatusMonitorNoMoreDescriptorData                                        = 0xC01D0008
	StatusMonitorInvalidDetailedTimingBlock                                  = 0xC01D0009
	StatusMonitorInvalidManufactureDate                                      = 0xC01D000A
	StatusGraphicsNotExclusiveModeOwner                                      = 0xC01E0000
	StatusGraphicsInsufficientDmaBuffer                                      = 0xC01E0001
	StatusGraphicsInvalidDisplayAdapter                                      = 0xC01E0002
	StatusGraphicsAdapterWasReset                                            = 0xC01E0003
	StatusGraphicsInvalidDriverModel                                         = 0xC01E0004
	StatusGraphicsPresentModeChanged                                         = 0xC01E0005
	StatusGraphicsPresentOccluded                                            = 0xC01E0006
	StatusGraphicsPresentDenied                                              = 0xC01E0007
	StatusGraphicsCannotcolorconvert                                         = 0xC01E0008
	StatusGraphicsDriverMismatch                                             = 0xC01E0009
	StatusGraphicsPartialDataPopulated                                       = 0x401E000A
	StatusGraphicsPresentRedirectionDisabled                                 = 0xC01E000B
	StatusGraphicsPresentUnoccluded                                          = 0xC01E000C
	StatusGraphicsNoVideoMemory                                              = 0xC01E0100
	StatusGraphicsCantLockMemory                                             = 0xC01E0101
	StatusGraphicsAllocationBusy                                             = 0xC01E0102
	StatusGraphicsTooManyReferences                                          = 0xC01E0103
	StatusGraphicsTryAgainLater                                              = 0xC01E0104
	StatusGraphicsTryAgainNow                                                = 0xC01E0105
	StatusGraphicsAllocationInvalid                                          = 0xC01E0106
	StatusGraphicsUnswizzlingApertureUnavailable                             = 0xC01E0107
	StatusGraphicsUnswizzlingApertureUnsupported                             = 0xC01E0108
	StatusGraphicsCantEvictPinnedAllocation                                  = 0xC01E0109
	StatusGraphicsInvalidAllocationUsage                                     = 0xC01E0110
	StatusGraphicsCantRenderLockedAllocation                                 = 0xC01E0111
	StatusGraphicsAllocationClosed                                           = 0xC01E0112
	StatusGraphicsInvalidAllocationInstance                                  = 0xC01E0113
	StatusGraphicsInvalidAllocationHandle                                    = 0xC01E0114
	StatusGraphicsWrongAllocationDevice                                      = 0xC01E0115
	StatusGraphicsAllocationContentLost                                      = 0xC01E0116
	StatusGraphicsGpuExceptionOnDevice                                       = 0xC01E0200
	StatusGraphicsInvalidVidpnTopology                                       = 0xC01E0300
	StatusGraphicsVidpnTopologyNotSupported                                  = 0xC01E0301
	StatusGraphicsVidpnTopologyCurrentlyNotSupported                         = 0xC01E0302
	StatusGraphicsInvalidVidpn                                               = 0xC01E0303
	StatusGraphicsInvalidVideoPresentSource                                  = 0xC01E0304
	StatusGraphicsInvalidVideoPresentTarget                                  = 0xC01E0305
	StatusGraphicsVidpnModalityNotSupported                                  = 0xC01E0306
	StatusGraphicsModeNotPinned                                              = 0x401E0307
	StatusGraphicsInvalidVidpnSourcemodeset                                  = 0xC01E0308
	StatusGraphicsInvalidVidpnTargetmodeset                                  = 0xC01E0309
	StatusGraphicsInvalidFrequency                                           = 0xC01E030A
	StatusGraphicsInvalidActiveRegion                                        = 0xC01E030B
	StatusGraphicsInvalidTotalRegion                                         = 0xC01E030C
	StatusGraphicsInvalidVideoPresentSourceMode                              = 0xC01E0310
	StatusGraphicsInvalidVideoPresentTargetMode                              = 0xC01E0311
	StatusGraphicsPinnedModeMustRemainInSet                                  = 0xC01E0312
	StatusGraphicsPathAlreadyInTopology                                      = 0xC01E0313
	StatusGraphicsModeAlreadyInModeset                                       = 0xC01E0314
	StatusGraphicsInvalidVideopresentsourceset                               = 0xC01E0315
	StatusGraphicsInvalidVideopresenttargetset                               = 0xC01E0316
	StatusGraphicsSourceAlreadyInSet                                         = 0xC01E0317
	StatusGraphicsTargetAlreadyInSet                                         = 0xC01E0318
	StatusGraphicsInvalidVidpnPresentPath                                    = 0xC01E0319
	StatusGraphicsNoRecommendedVidpnTopology                                 = 0xC01E031A
	StatusGraphicsInvalidMonitorFrequencyrangeset                            = 0xC01E031B
	StatusGraphicsInvalidMonitorFrequencyrange                               = 0xC01E031C
	StatusGraphicsFrequencyrangeNotInSet                                     = 0xC01E031D
	StatusGraphicsNoPreferredMode                                            = 0x401E031E
	StatusGraphicsFrequencyrangeAlreadyInSet                                 = 0xC01E031F
	StatusGraphicsStaleModeset                                               = 0xC01E0320
	StatusGraphicsInvalidMonitorSourcemodeset                                = 0xC01E0321
	StatusGraphicsInvalidMonitorSourceMode                                   = 0xC01E0322
	StatusGraphicsNoRecommendedFunctionalVidpn                               = 0xC01E0323
	StatusGraphicsModeIdMustBeUnique                                         = 0xC01E0324
	StatusGraphicsEmptyAdapterMonitorModeSupportIntersection                 = 0xC01E0325
	StatusGraphicsVideoPresentTargetsLessThanSources                         = 0xC01E0326
	StatusGraphicsPathNotInTopology                                          = 0xC01E0327
	StatusGraphicsAdapterMustHaveAtLeastOneSource                            = 0xC01E0328
	StatusGraphicsAdapterMustHaveAtLeastOneTarget                            = 0xC01E0329
	StatusGraphicsInvalidMonitordescriptorset                                = 0xC01E032A
	StatusGraphicsInvalidMonitordescriptor                                   = 0xC01E032B
	StatusGraphicsMonitordescriptorNotInSet                                  = 0xC01E032C
	StatusGraphicsMonitordescriptorAlreadyInSet                              = 0xC01E032D
	StatusGraphicsMonitordescriptorIdMustBeUnique                            = 0xC01E032E
	StatusGraphicsInvalidVidpnTargetSubsetType                               = 0xC01E032F
	StatusGraphicsResourcesNotRelated                                        = 0xC01E0330
	StatusGraphicsSourceIdMustBeUnique                                       = 0xC01E0331
	StatusGraphicsTargetIdMustBeUnique                                       = 0xC01E0332
	StatusGraphicsNoAvailableVidpnTarget                                     = 0xC01E0333
	StatusGraphicsMonitorCouldNotBeAssociatedWithAdapter                     = 0xC01E0334
	StatusGraphicsNoVidpnmgr                                                 = 0xC01E0335
	StatusGraphicsNoActiveVidpn                                              = 0xC01E0336
	StatusGraphicsStaleVidpnTopology                                         = 0xC01E0337
	StatusGraphicsMonitorNotConnected                                        = 0xC01E0338
	StatusGraphicsSourceNotInTopology                                        = 0xC01E0339
	StatusGraphicsInvalidPrimarysurfaceSize                                  = 0xC01E033A
	StatusGraphicsInvalidVisibleregionSize                                   = 0xC01E033B
	StatusGraphicsInvalidStride                                              = 0xC01E033C
	StatusGraphicsInvalidPixelformat                                         = 0xC01E033D
	StatusGraphicsInvalidColorbasis                                          = 0xC01E033E
	StatusGraphicsInvalidPixelvalueaccessmode                                = 0xC01E033F
	StatusGraphicsTargetNotInTopology                                        = 0xC01E0340
	StatusGraphicsNoDisplayModeManagementSupport                             = 0xC01E0341
	StatusGraphicsVidpnSourceInUse                                           = 0xC01E0342
	StatusGraphicsCantAccessActiveVidpn                                      = 0xC01E0343
	StatusGraphicsInvalidPathImportanceOrdinal                               = 0xC01E0344
	StatusGraphicsInvalidPathContentGeometryTransformation                   = 0xC01E0345
	StatusGraphicsPathContentGeometryTransformationNotSupported              = 0xC01E0346
	StatusGraphicsInvalidGammaRamp                                           = 0xC01E0347
	StatusGraphicsGammaRampNotSupported                                      = 0xC01E0348
	StatusGraphicsMultisamplingNotSupported                                  = 0xC01E0349
	StatusGraphicsModeNotInModeset                                           = 0xC01E034A
	StatusGraphicsDatasetIsEmpty                                             = 0x401E034B
	StatusGraphicsNoMoreElementsInDataset                                    = 0x401E034C
	StatusGraphicsInvalidVidpnTopologyRecommendationReason                   = 0xC01E034D
	StatusGraphicsInvalidPathContentType                                     = 0xC01E034E
	StatusGraphicsInvalidCopyprotectionType                                  = 0xC01E034F
	StatusGraphicsUnassignedModesetAlreadyExists                             = 0xC01E0350
	StatusGraphicsPathContentGeometryTransformationNotPinned                 = 0x401E0351
	StatusGraphicsInvalidScanlineOrdering                                    = 0xC01E0352
	StatusGraphicsTopologyChangesNotAllowed                                  = 0xC01E0353
	StatusGraphicsNoAvailableImportanceOrdinals                              = 0xC01E0354
	StatusGraphicsIncompatiblePrivateFormat                                  = 0xC01E0355
	StatusGraphicsInvalidModePruningAlgorithm                                = 0xC01E0356
	StatusGraphicsInvalidMonitorCapabilityOrigin                             = 0xC01E0357
	StatusGraphicsInvalidMonitorFrequencyrangeConstraint                     = 0xC01E0358
	StatusGraphicsMaxNumPathsReached                                         = 0xC01E0359
	StatusGraphicsCancelVidpnTopologyAugmentation                            = 0xC01E035A
	StatusGraphicsInvalidClientType                                          = 0xC01E035B
	StatusGraphicsClientvidpnNotSet                                          = 0xC01E035C
	StatusGraphicsSpecifiedChildAlreadyConnected                             = 0xC01E0400
	StatusGraphicsChildDescriptorNotSupported                                = 0xC01E0401
	StatusGraphicsUnknownChildStatus                                         = 0x401E042F
	StatusGraphicsNotALinkedAdapter                                          = 0xC01E0430
	StatusGraphicsLeadlinkNotEnumerated                                      = 0xC01E0431
	StatusGraphicsChainlinksNotEnumerated                                    = 0xC01E0432
	StatusGraphicsAdapterChainNotReady                                       = 0xC01E0433
	StatusGraphicsChainlinksNotStarted                                       = 0xC01E0434
	StatusGraphicsChainlinksNotPoweredOn                                     = 0xC01E0435
	StatusGraphicsInconsistentDeviceLinkState                                = 0xC01E0436
	StatusGraphicsLeadlinkStartDeferred                                      = 0x401E0437
	StatusGraphicsNotPostDeviceDriver                                        = 0xC01E0438
	StatusGraphicsPollingTooFrequently                                       = 0x401E0439
	StatusGraphicsStartDeferred                                              = 0x401E043A
	StatusGraphicsAdapterAccessNotExcluded                                   = 0xC01E043B
	StatusGraphicsOpmNotSupported                                            = 0xC01E0500
	StatusGraphicsCoppNotSupported                                           = 0xC01E0501
	StatusGraphicsUabNotSupported                                            = 0xC01E0502
	StatusGraphicsOpmInvalidEncryptedParameters                              = 0xC01E0503
	StatusGraphicsOpmNoProtectedOutputsExist                                 = 0xC01E0505
	StatusGraphicsOpmInternalError                                           = 0xC01E050B
	StatusGraphicsOpmInvalidHandle                                           = 0xC01E050C
	StatusGraphicsPvpInvalidCertificateLength                                = 0xC01E050E
	StatusGraphicsOpmSpanningModeEnabled                                     = 0xC01E050F
	StatusGraphicsOpmTheaterModeEnabled                                      = 0xC01E0510
	StatusGraphicsPvpHfsFailed                                               = 0xC01E0511
	StatusGraphicsOpmInvalidSrm                                              = 0xC01E0512
	StatusGraphicsOpmOutputDoesNotSupportHdcp                                = 0xC01E0513
	StatusGraphicsOpmOutputDoesNotSupportAcp                                 = 0xC01E0514
	StatusGraphicsOpmOutputDoesNotSupportCgmsa                               = 0xC01E0515
	StatusGraphicsOpmHdcpSrmNeverSet                                         = 0xC01E0516
	StatusGraphicsOpmResolutionTooHigh                                       = 0xC01E0517
	StatusGraphicsOpmAllHdcpHardwareAlreadyInUse                             = 0xC01E0518
	StatusGraphicsOpmProtectedOutputNoLongerExists                           = 0xC01E051A
	StatusGraphicsOpmProtectedOutputDoesNotHaveCoppSemantics                 = 0xC01E051C
	StatusGraphicsOpmInvalidInformationRequest                               = 0xC01E051D
	StatusGraphicsOpmDriverInternalError                                     = 0xC01E051E
	StatusGraphicsOpmProtectedOutputDoesNotHaveOpmSemantics                  = 0xC01E051F
	StatusGraphicsOpmSignalingNotSupported                                   = 0xC01E0520
	StatusGraphicsOpmInvalidConfigurationRequest                             = 0xC01E0521
	StatusGraphicsI2CNotSupported                                            = 0xC01E0580
	StatusGraphicsI2CDeviceDoesNotExist                                      = 0xC01E0581
	StatusGraphicsI2CErrorTransmittingData                                   = 0xC01E0582
	StatusGraphicsI2CErrorReceivingData                                      = 0xC01E0583
	StatusGraphicsDdcciVcpNotSupported                                       = 0xC01E0584
	StatusGraphicsDdcciInvalidData                                           = 0xC01E0585
	StatusGraphicsDdcciMonitorReturnedInvalidTimingStatusByte                = 0xC01E0586
	StatusGraphicsDdcciInvalidCapabilitiesString                             = 0xC01E0587
	StatusGraphicsMcaInternalError                                           = 0xC01E0588
	StatusGraphicsDdcciInvalidMessageCommand                                 = 0xC01E0589
	StatusGraphicsDdcciInvalidMessageLength                                  = 0xC01E058A
	StatusGraphicsDdcciInvalidMessageChecksum                                = 0xC01E058B
	StatusGraphicsInvalidPhysicalMonitorHandle                               = 0xC01E058C
	StatusGraphicsMonitorNoLongerExists                                      = 0xC01E058D
	StatusGraphicsOnlyConsoleSessionSupported                                = 0xC01E05E0
	StatusGraphicsNoDisplayDeviceCorrespondsToName                           = 0xC01E05E1
	StatusGraphicsDisplayDeviceNotAttachedToDesktop                          = 0xC01E05E2
	StatusGraphicsMirroringDevicesNotSupported                               = 0xC01E05E3
	StatusGraphicsInvalidPointer                                             = 0xC01E05E4
	StatusGraphicsNoMonitorsCorrespondToDisplayDevice                        = 0xC01E05E5
	StatusGraphicsParameterArrayTooSmall                                     = 0xC01E05E6
	StatusGraphicsInternalError                                              = 0xC01E05E7
	StatusGraphicsSessionTypeChangeInProgress                                = 0xC01E05E8
	StatusFveLockedVolume                                                    = 0xC0210000
	StatusFveNotEncrypted                                                    = 0xC0210001
	StatusFveBadInformation                                                  = 0xC0210002
	StatusFveTooSmall                                                        = 0xC0210003
	StatusFveFailedWrongFs                                                   = 0xC0210004
	StatusFveBadPartitionSize                                                = 0xC0210005
	StatusFveFsNotExtended                                                   = 0xC0210006
	StatusFveFsMounted                                                       = 0xC0210007
	StatusFveNoLicense                                                       = 0xC0210008
	StatusFveActionNotAllowed                                                = 0xC0210009
	StatusFveBadData                                                         = 0xC021000A
	StatusFveVolumeNotBound                                                  = 0xC021000B
	StatusFveNotDataVolume                                                   = 0xC021000C
	StatusFveConvReadError                                                   = 0xC021000D
	StatusFveConvWriteError                                                  = 0xC021000E
	StatusFveOverlappedUpdate                                                = 0xC021000F
	StatusFveFailedSectorSize                                                = 0xC0210010
	StatusFveFailedAuthentication                                            = 0xC0210011
	StatusFveNotOsVolume                                                     = 0xC0210012
	StatusFveKeyfileNotFound                                                 = 0xC0210013
	StatusFveKeyfileInvalid                                                  = 0xC0210014
	StatusFveKeyfileNoVmk                                                    = 0xC0210015
	StatusFveTpmDisabled                                                     = 0xC0210016
	StatusFveTpmSrkAuthNotZero                                               = 0xC0210017
	StatusFveTpmInvalidPcr                                                   = 0xC0210018
	StatusFveTpmNoVmk                                                        = 0xC0210019
	StatusFvePinInvalid                                                      = 0xC021001A
	StatusFveAuthInvalidApplication                                          = 0xC021001B
	StatusFveAuthInvalidConfig                                               = 0xC021001C
	StatusFveDebuggerEnabled                                                 = 0xC021001D
	StatusFveDryRunFailed                                                    = 0xC021001E
	StatusFveBadMetadataPointer                                              = 0xC021001F
	StatusFveOldMetadataCopy                                                 = 0xC0210020
	StatusFveRebootRequired                                                  = 0xC0210021
	StatusFveRawAccess                                                       = 0xC0210022
	StatusFveRawBlocked                                                      = 0xC0210023
	StatusFveNoAutounlockMasterKey                                           = 0xC0210024
	StatusFveMorFailed                                                       = 0xC0210025
	StatusFveNoFeatureLicense                                                = 0xC0210026
	StatusFvePolicyUserDisableRdvNotAllowed                                  = 0xC0210027
	StatusFveConvRecoveryFailed                                              = 0xC0210028
	StatusFveVirtualizedSpaceTooBig                                          = 0xC0210029
	StatusFveInvalidDatumType                                                = 0xC021002A
	StatusFveVolumeTooSmall                                                  = 0xC0210030
	StatusFveEnhPinInvalid                                                   = 0xC0210031
	StatusFwpCalloutNotFound                                                 = 0xC0220001
	StatusFwpConditionNotFound                                               = 0xC0220002
	StatusFwpFilterNotFound                                                  = 0xC0220003
	StatusFwpLayerNotFound                                                   = 0xC0220004
	StatusFwpProviderNotFound                                                = 0xC0220005
	StatusFwpProviderContextNotFound                                         = 0xC0220006
	StatusFwpSublayerNotFound                                                = 0xC0220007
	StatusFwpNotFound                                                        = 0xC0220008
	StatusFwpAlreadyExists                                                   = 0xC0220009
	StatusFwpInUse                                                           = 0xC022000A
	StatusFwpDynamicSessionInProgress                                        = 0xC022000B
	StatusFwpWrongSession                                                    = 0xC022000C
	StatusFwpNoTxnInProgress                                                 = 0xC022000D
	StatusFwpTxnInProgress                                                   = 0xC022000E
	StatusFwpTxnAborted                                                      = 0xC022000F
	StatusFwpSessionAborted                                                  = 0xC0220010
	StatusFwpIncompatibleTxn                                                 = 0xC0220011
	StatusFwpTimeout                                                         = 0xC0220012
	StatusFwpNetEventsDisabled                                               = 0xC0220013
	StatusFwpIncompatibleLayer                                               = 0xC0220014
	StatusFwpKmClientsOnly                                                   = 0xC0220015
	StatusFwpLifetimeMismatch                                                = 0xC0220016
	StatusFwpBuiltinObject                                                   = 0xC0220017
	StatusFwpTooManyCallouts                                                 = 0xC0220018
	StatusFwpNotificationDropped                                             = 0xC0220019
	StatusFwpTrafficMismatch                                                 = 0xC022001A
	StatusFwpIncompatibleSaState                                             = 0xC022001B
	StatusFwpNullPointer                                                     = 0xC022001C
	StatusFwpInvalidEnumerator                                               = 0xC022001D
	StatusFwpInvalidFlags                                                    = 0xC022001E
	StatusFwpInvalidNetMask                                                  = 0xC022001F
	StatusFwpInvalidRange                                                    = 0xC0220020
	StatusFwpInvalidInterval                                                 = 0xC0220021
	StatusFwpZeroLengthArray                                                 = 0xC0220022
	StatusFwpNullDisplayName                                                 = 0xC0220023
	StatusFwpInvalidActionType                                               = 0xC0220024
	StatusFwpInvalidWeight                                                   = 0xC0220025
	StatusFwpMatchTypeMismatch                                               = 0xC0220026
	StatusFwpTypeMismatch                                                    = 0xC0220027
	StatusFwpOutOfBounds                                                     = 0xC0220028
	StatusFwpReserved                                                        = 0xC0220029
	StatusFwpDuplicateCondition                                              = 0xC022002A
	StatusFwpDuplicateKeymod                                                 = 0xC022002B
	StatusFwpActionIncompatibleWithLayer                                     = 0xC022002C
	StatusFwpActionIncompatibleWithSublayer                                  = 0xC022002D
	StatusFwpContextIncompatibleWithLayer                                    = 0xC022002E
	StatusFwpContextIncompatibleWithCallout                                  = 0xC022002F
	StatusFwpIncompatibleAuthMethod                                          = 0xC0220030
	StatusFwpIncompatibleDhGroup                                             = 0xC0220031
	StatusFwpEmNotSupported                                                  = 0xC0220032
	StatusFwpNeverMatch                                                      = 0xC0220033
	StatusFwpProviderContextMismatch                                         = 0xC0220034
	StatusFwpInvalidParameter                                                = 0xC0220035
	StatusFwpTooManySublayers                                                = 0xC0220036
	StatusFwpCalloutNotificationFailed                                       = 0xC0220037
	StatusFwpInvalidAuthTransform                                            = 0xC0220038
	StatusFwpInvalidCipherTransform                                          = 0xC0220039
	StatusFwpIncompatibleCipherTransform                                     = 0xC022003A
	StatusFwpInvalidTransformCombination                                     = 0xC022003B
	StatusFwpDuplicateAuthMethod                                             = 0xC022003C
	StatusFwpTcpipNotReady                                                   = 0xC0220100
	StatusFwpInjectHandleClosing                                             = 0xC0220101
	StatusFwpInjectHandleStale                                               = 0xC0220102
	StatusFwpCannotPend                                                      = 0xC0220103
	StatusFwpDropNoicmp                                                      = 0xC0220104
	StatusNdisClosing                                                        = 0xC0230002
	StatusNdisBadVersion                                                     = 0xC0230004
	StatusNdisBadCharacteristics                                             = 0xC0230005
	StatusNdisAdapterNotFound                                                = 0xC0230006
	StatusNdisOpenFailed                                                     = 0xC0230007
	StatusNdisDeviceFailed                                                   = 0xC0230008
	StatusNdisMulticastFull                                                  = 0xC0230009
	StatusNdisMulticastExists                                                = 0xC023000A
	StatusNdisMulticastNotFound                                              = 0xC023000B
	StatusNdisRequestAborted                                                 = 0xC023000C
	StatusNdisResetInProgress                                                = 0xC023000D
	StatusNdisNotSupported                                                   = 0xC02300BB
	StatusNdisInvalidPacket                                                  = 0xC023000F
	StatusNdisAdapterNotReady                                                = 0xC0230011
	StatusNdisInvalidLength                                                  = 0xC0230014
	StatusNdisInvalidData                                                    = 0xC0230015
	StatusNdisBufferTooShort                                                 = 0xC0230016
	StatusNdisInvalidOid                                                     = 0xC0230017
	StatusNdisAdapterRemoved                                                 = 0xC0230018
	StatusNdisUnsupportedMedia                                               = 0xC0230019
	StatusNdisGroupAddressInUse                                              = 0xC023001A
	StatusNdisFileNotFound                                                   = 0xC023001B
	StatusNdisErrorReadingFile                                               = 0xC023001C
	StatusNdisAlreadyMapped                                                  = 0xC023001D
	StatusNdisResourceConflict                                               = 0xC023001E
	StatusNdisMediaDisconnected                                              = 0xC023001F
	StatusNdisInvalidAddress                                                 = 0xC0230022
	StatusNdisInvalidDeviceRequest                                           = 0xC0230010
	StatusNdisPaused                                                         = 0xC023002A
	StatusNdisInterfaceNotFound                                              = 0xC023002B
	StatusNdisUnsupportedRevision                                            = 0xC023002C
	StatusNdisInvalidPort                                                    = 0xC023002D
	StatusNdisInvalidPortState                                               = 0xC023002E
	StatusNdisLowPowerState                                                  = 0xC023002F
	StatusNdisDot11AutoConfigEnabled                                         = 0xC0232000
	StatusNdisDot11MediaInUse                                                = 0xC0232001
	StatusNdisDot11PowerStateInvalid                                         = 0xC0232002
	StatusNdisPmWolPatternListFull                                           = 0xC0232003
	StatusNdisPmProtocolOffloadListFull                                      = 0xC0232004
	StatusNdisIndicationRequired                                             = 0x40230001
	StatusNdisOffloadPolicy                                                  = 0xC023100F
	StatusNdisOffloadConnectionRejected                                      = 0xC0231012
	StatusNdisOffloadPathRejected                                            = 0xC0231013
	StatusHvInvalidHypercallCode                                             = 0xC0350002
	StatusHvInvalidHypercallInput                                            = 0xC0350003
	StatusHvInvalidAlignment                                                 = 0xC0350004
	StatusHvInvalidParameter                                                 = 0xC0350005
	StatusHvAccessDenied                                                     = 0xC0350006
	StatusHvInvalidPartitionState                                            = 0xC0350007
	StatusHvOperationDenied                                                  = 0xC0350008
	StatusHvUnknownProperty                                                  = 0xC0350009
	StatusHvPropertyValueOutOfRange                                          = 0xC035000A
	StatusHvInsufficientMemory                                               = 0xC035000B
	StatusHvPartitionTooDeep                                                 = 0xC035000C
	StatusHvInvalidPartitionId                                               = 0xC035000D
	StatusHvInvalidVpIndex                                                   = 0xC035000E
	StatusHvInvalidPortId                                                    = 0xC0350011
	StatusHvInvalidConnectionId                                              = 0xC0350012
	StatusHvInsufficientBuffers                                              = 0xC0350013
	StatusHvNotAcknowledged                                                  = 0xC0350014
	StatusHvAcknowledged                                                     = 0xC0350016
	StatusHvInvalidSaveRestoreState                                          = 0xC0350017
	StatusHvInvalidSynicState                                                = 0xC0350018
	StatusHvObjectInUse                                                      = 0xC0350019
	StatusHvInvalidProximityDomainInfo                                       = 0xC035001A
	StatusHvNoData                                                           = 0xC035001B
	StatusHvInactive                                                         = 0xC035001C
	StatusHvNoResources                                                      = 0xC035001D
	StatusHvFeatureUnavailable                                               = 0xC035001E
	StatusHvNotPresent                                                       = 0xC0351000
	StatusVidDuplicateHandler                                                = 0xC0370001
	StatusVidTooManyHandlers                                                 = 0xC0370002
	StatusVidQueueFull                                                       = 0xC0370003
	StatusVidHandlerNotPresent                                               = 0xC0370004
	StatusVidInvalidObjectName                                               = 0xC0370005
	StatusVidPartitionNameTooLong                                            = 0xC0370006
	StatusVidMessageQueueNameTooLong                                         = 0xC0370007
	StatusVidPartitionAlreadyExists                                          = 0xC0370008
	StatusVidPartitionDoesNotExist                                           = 0xC0370009
	StatusVidPartitionNameNotFound                                           = 0xC037000A
	StatusVidMessageQueueAlreadyExists                                       = 0xC037000B
	StatusVidExceededMbpEntryMapLimit                                        = 0xC037000C
	StatusVidMbStillReferenced                                               = 0xC037000D
	StatusVidChildGpaPageSetCorrupted                                        = 0xC037000E
	StatusVidInvalidNumaSettings                                             = 0xC037000F
	StatusVidInvalidNumaNodeIndex                                            = 0xC0370010
	StatusVidNotificationQueueAlreadyAssociated                              = 0xC0370011
	StatusVidInvalidMemoryBlockHandle                                        = 0xC0370012
	StatusVidPageRangeOverflow                                               = 0xC0370013
	StatusVidInvalidMessageQueueHandle                                       = 0xC0370014
	StatusVidInvalidGpaRangeHandle                                           = 0xC0370015
	StatusVidNoMemoryBlockNotificationQueue                                  = 0xC0370016
	StatusVidMemoryBlockLockCountExceeded                                    = 0xC0370017
	StatusVidInvalidPpmHandle                                                = 0xC0370018
	StatusVidMbpsAreLocked                                                   = 0xC0370019
	StatusVidMessageQueueClosed                                              = 0xC037001A
	StatusVidVirtualProcessorLimitExceeded                                   = 0xC037001B
	StatusVidStopPending                                                     = 0xC037001C
	StatusVidInvalidProcessorState                                           = 0xC037001D
	StatusVidExceededKmContextCountLimit                                     = 0xC037001E
	StatusVidKmInterfaceAlreadyInitialized                                   = 0xC037001F
	StatusVidMbPropertyAlreadySetReset                                       = 0xC0370020
	StatusVidMmioRangeDestroyed                                              = 0xC0370021
	StatusVidInvalidChildGpaPageSet                                          = 0xC0370022
	StatusVidReservePageSetIsBeingUsed                                       = 0xC0370023
	StatusVidReservePageSetTooSmall                                          = 0xC0370024
	StatusVidMbpAlreadyLockedUsingReservedPage                               = 0xC0370025
	StatusVidMbpCountExceededLimit                                           = 0xC0370026
	StatusVidSavedStateCorrupt                                               = 0xC0370027
	StatusVidSavedStateUnrecognizedItem                                      = 0xC0370028
	StatusVidSavedStateIncompatible                                          = 0xC0370029
	StatusVidRemoteNodeParentGpaPagesUsed                                    = 0x80370001
	StatusIpsecBadSpi                                                        = 0xC0360001
	StatusIpsecSaLifetimeExpired                                             = 0xC0360002
	StatusIpsecWrongSa                                                       = 0xC0360003
	StatusIpsecReplayCheckFailed                                             = 0xC0360004
	StatusIpsecInvalidPacket                                                 = 0xC0360005
	StatusIpsecIntegrityCheckFailed                                          = 0xC0360006
	StatusIpsecClearTextDrop                                                 = 0xC0360007
	StatusIpsecAuthFirewallDrop                                              = 0xC0360008
	StatusIpsecThrottleDrop                                                  = 0xC0360009
	StatusIpsecDospBlock                                                     = 0xC0368000
	StatusIpsecDospReceivedMulticast                                         = 0xC0368001
	StatusIpsecDospInvalidPacket                                             = 0xC0368002
	StatusIpsecDospStateLookupFailed                                         = 0xC0368003
	StatusIpsecDospMaxEntries                                                = 0xC0368004
	StatusIpsecDospKeymodNotAllowed                                          = 0xC0368005
	StatusIpsecDospMaxPerIpRatelimitQueues                                   = 0xC0368006
	StatusVolmgrIncompleteRegeneration                                       = 0x80380001
	StatusVolmgrIncompleteDiskMigration                                      = 0x80380002
	StatusVolmgrDatabaseFull                                                 = 0xC0380001
	StatusVolmgrDiskConfigurationCorrupted                                   = 0xC0380002
	StatusVolmgrDiskConfigurationNotInSync                                   = 0xC0380003
	StatusVolmgrPackConfigUpdateFailed                                       = 0xC0380004
	StatusVolmgrDiskContainsNonSimpleVolume                                  = 0xC0380005
	StatusVolmgrDiskDuplicate                                                = 0xC0380006
	StatusVolmgrDiskDynamic                                                  = 0xC0380007
	StatusVolmgrDiskIdInvalid                                                = 0xC0380008
	StatusVolmgrDiskInvalid                                                  = 0xC0380009
	StatusVolmgrDiskLastVoter                                                = 0xC038000A
	StatusVolmgrDiskLayoutInvalid                                            = 0xC038000B
	StatusVolmgrDiskLayoutNonBasicBetweenBasicPartitions                     = 0xC038000C
	StatusVolmgrDiskLayoutNotCylinderAligned                                 = 0xC038000D
	StatusVolmgrDiskLayoutPartitionsTooSmall                                 = 0xC038000E
	StatusVolmgrDiskLayoutPrimaryBetweenLogicalPartitions                    = 0xC038000F
	StatusVolmgrDiskLayoutTooManyPartitions                                  = 0xC0380010
	StatusVolmgrDiskMissing                                                  = 0xC0380011
	StatusVolmgrDiskNotEmpty                                                 = 0xC0380012
	StatusVolmgrDiskNotEnoughSpace                                           = 0xC0380013
	StatusVolmgrDiskRevectoringFailed                                        = 0xC0380014
	StatusVolmgrDiskSectorSizeInvalid                                        = 0xC0380015
	StatusVolmgrDiskSetNotContained                                          = 0xC0380016
	StatusVolmgrDiskUsedByMultipleMembers                                    = 0xC0380017
	StatusVolmgrDiskUsedByMultiplePlexes                                     = 0xC0380018
	StatusVolmgrDynamicDiskNotSupported                                      = 0xC0380019
	StatusVolmgrExtentAlreadyUsed                                            = 0xC038001A
	StatusVolmgrExtentNotContiguous                                          = 0xC038001B
	StatusVolmgrExtentNotInPublicRegion                                      = 0xC038001C
	StatusVolmgrExtentNotSectorAligned                                       = 0xC038001D
	StatusVolmgrExtentOverlapsEbrPartition                                   = 0xC038001E
	StatusVolmgrExtentVolumeLengthsDoNotMatch                                = 0xC038001F
	StatusVolmgrFaultTolerantNotSupported                                    = 0xC0380020
	StatusVolmgrInterleaveLengthInvalid                                      = 0xC0380021
	StatusVolmgrMaximumRegisteredUsers                                       = 0xC0380022
	StatusVolmgrMemberInSync                                                 = 0xC0380023
	StatusVolmgrMemberIndexDuplicate                                         = 0xC0380024
	StatusVolmgrMemberIndexInvalid                                           = 0xC0380025
	StatusVolmgrMemberMissing                                                = 0xC0380026
	StatusVolmgrMemberNotDetached                                            = 0xC0380027
	StatusVolmgrMemberRegenerating                                           = 0xC0380028
	StatusVolmgrAllDisksFailed                                               = 0xC0380029
	StatusVolmgrNoRegisteredUsers                                            = 0xC038002A
	StatusVolmgrNoSuchUser                                                   = 0xC038002B
	StatusVolmgrNotificationReset                                            = 0xC038002C
	StatusVolmgrNumberOfMembersInvalid                                       = 0xC038002D
	StatusVolmgrNumberOfPlexesInvalid                                        = 0xC038002E
	StatusVolmgrPackDuplicate                                                = 0xC038002F
	StatusVolmgrPackIdInvalid                                                = 0xC0380030
	StatusVolmgrPackInvalid                                                  = 0xC0380031
	StatusVolmgrPackNameInvalid                                              = 0xC0380032
	StatusVolmgrPackOffline                                                  = 0xC0380033
	StatusVolmgrPackHasQuorum                                                = 0xC0380034
	StatusVolmgrPackWithoutQuorum                                            = 0xC0380035
	StatusVolmgrPartitionStyleInvalid                                        = 0xC0380036
	StatusVolmgrPartitionUpdateFailed                                        = 0xC0380037
	StatusVolmgrPlexInSync                                                   = 0xC0380038
	StatusVolmgrPlexIndexDuplicate                                           = 0xC0380039
	StatusVolmgrPlexIndexInvalid                                             = 0xC038003A
	StatusVolmgrPlexLastActive                                               = 0xC038003B
	StatusVolmgrPlexMissing                                                  = 0xC038003C
	StatusVolmgrPlexRegenerating                                             = 0xC038003D
	StatusVolmgrPlexTypeInvalid                                              = 0xC038003E
	StatusVolmgrPlexNotRaid5                                                 = 0xC038003F
	StatusVolmgrPlexNotSimple                                                = 0xC0380040
	StatusVolmgrStructureSizeInvalid                                         = 0xC0380041
	StatusVolmgrTooManyNotificationRequests                                  = 0xC0380042
	StatusVolmgrTransactionInProgress                                        = 0xC0380043
	StatusVolmgrUnexpectedDiskLayoutChange                                   = 0xC0380044
	StatusVolmgrVolumeContainsMissingDisk                                    = 0xC0380045
	StatusVolmgrVolumeIdInvalid                                              = 0xC0380046
	StatusVolmgrVolumeLengthInvalid                                          = 0xC0380047
	StatusVolmgrVolumeLengthNotSectorSizeMultiple                            = 0xC0380048
	StatusVolmgrVolumeNotMirrored                                            = 0xC0380049
	StatusVolmgrVolumeNotRetained                                            = 0xC038004A
	StatusVolmgrVolumeOffline                                                = 0xC038004B
	StatusVolmgrVolumeRetained                                               = 0xC038004C
	StatusVolmgrNumberOfExtentsInvalid                                       = 0xC038004D
	StatusVolmgrDifferentSectorSize                                          = 0xC038004E
	StatusVolmgrBadBootDisk                                                  = 0xC038004F
	StatusVolmgrPackConfigOffline                                            = 0xC0380050
	StatusVolmgrPackConfigOnline                                             = 0xC0380051
	StatusVolmgrNotPrimaryPack                                               = 0xC0380052
	StatusVolmgrPackLogUpdateFailed                                          = 0xC0380053
	StatusVolmgrNumberOfDisksInPlexInvalid                                   = 0xC0380054
	StatusVolmgrNumberOfDisksInMemberInvalid                                 = 0xC0380055
	StatusVolmgrVolumeMirrored                                               = 0xC0380056
	StatusVolmgrPlexNotSimpleSpanned                                         = 0xC0380057
	StatusVolmgrNoValidLogCopies                                             = 0xC0380058
	StatusVolmgrPrimaryPackPresent                                           = 0xC0380059
	StatusVolmgrNumberOfDisksInvalid                                         = 0xC038005A
	StatusVolmgrMirrorNotSupported                                           = 0xC038005B
	StatusVolmgrRaid5NotSupported                                            = 0xC038005C
	StatusBcdNotAllEntriesImported                                           = 0x80390001
	StatusBcdTooManyElements                                                 = 0xC0390002
	StatusBcdNotAllEntriesSynchronized                                       = 0x80390003
	StatusVhdDriveFooterMissing                                              = 0xC03A0001
	StatusVhdDriveFooterChecksumMismatch                                     = 0xC03A0002
	StatusVhdDriveFooterCorrupt                                              = 0xC03A0003
	StatusVhdFormatUnknown                                                   = 0xC03A0004
	StatusVhdFormatUnsupportedVersion                                        = 0xC03A0005
	StatusVhdSparseHeaderChecksumMismatch                                    = 0xC03A0006
	StatusVhdSparseHeaderUnsupportedVersion                                  = 0xC03A0007
	StatusVhdSparseHeaderCorrupt                                             = 0xC03A0008
	StatusVhdBlockAllocationFailure                                          = 0xC03A0009
	StatusVhdBlockAllocationTableCorrupt                                     = 0xC03A000A
	StatusVhdInvalidBlockSize                                                = 0xC03A000B
	StatusVhdBitmapMismatch                                                  = 0xC03A000C
	StatusVhdParentVhdNotFound                                               = 0xC03A000D
	StatusVhdChildParentIdMismatch                                           = 0xC03A000E
	StatusVhdChildParentTimestampMismatch                                    = 0xC03A000F
	StatusVhdMetadataReadFailure                                             = 0xC03A0010
	StatusVhdMetadataWriteFailure                                            = 0xC03A0011
	StatusVhdInvalidSize                                                     = 0xC03A0012
	StatusVhdInvalidFileSize                                                 = 0xC03A0013
	StatusVirtdiskProviderNotFound                                           = 0xC03A0014
	StatusVirtdiskNotVirtualDisk                                             = 0xC03A0015
	StatusVhdParentVhdAccessDenied                                           = 0xC03A0016
	StatusVhdChildParentSizeMismatch                                         = 0xC03A0017
	StatusVhdDifferencingChainCycleDetected                                  = 0xC03A0018
	StatusVhdDifferencingChainErrorInParent                                  = 0xC03A0019
	StatusVirtualDiskLimitation                                              = 0xC03A001A
	StatusVhdInvalidType                                                     = 0xC03A001B
	StatusVhdInvalidState                                                    = 0xC03A001C
	StatusVirtdiskUnsupportedDiskSectorSize                                  = 0xC03A001D
	StatusQueryStorageError                                                  = 0x803A0001
	StatusDisNotPresent                                                      = 0xC03C0001
	StatusDisAttributeNotFound                                               = 0xC03C0002
	StatusDisUnrecognizedAttribute                                           = 0xC03C0003
	StatusDisPartialData                                                     = 0xC03C0004
)

func (k NtstatusKind) String() string {
	switch k {
	case 0x00000000:
		return "StatusSuccess"
	case 0x00000001:
		return "StatusWait1"
	case 0x00000002:
		return "StatusWait2"
	case 0x00000003:
		return "StatusWait3"
	case 0x0000003F:
		return "StatusWait63"
	case 0x00000080:
		return "StatusAbandoned"
	case 0x000000BF:
		return "StatusAbandonedWait63"
	case 0x000000C0:
		return "StatusUserApc"
	case 0x00000100:
		return "StatusKernelApc"
	case 0x00000101:
		return "StatusAlerted"
	case 0x00000102:
		return "StatusTimeout"
	case 0x00000103:
		return "StatusPending"
	case 0x00000104:
		return "StatusReparse"
	case 0x00000105:
		return "StatusMoreEntries"
	case 0x00000106:
		return "StatusNotAllAssigned"
	case 0x00000107:
		return "StatusSomeNotMapped"
	case 0x00000108:
		return "StatusOplockBreakInProgress"
	case 0x00000109:
		return "StatusVolumeMounted"
	case 0x0000010A:
		return "StatusRxactCommitted"
	case 0x0000010B:
		return "StatusNotifyCleanup"
	case 0x0000010C:
		return "StatusNotifyEnumDir"
	case 0x0000010D:
		return "StatusNoQuotasForAccount"
	case 0x0000010E:
		return "StatusPrimaryTransportConnectFailed"
	case 0x00000110:
		return "StatusPageFaultTransition"
	case 0x00000111:
		return "StatusPageFaultDemandZero"
	case 0x00000112:
		return "StatusPageFaultCopyOnWrite"
	case 0x00000113:
		return "StatusPageFaultGuardPage"
	case 0x00000114:
		return "StatusPageFaultPagingFile"
	case 0x00000115:
		return "StatusCachePageLocked"
	case 0x00000116:
		return "StatusCrashDump"
	case 0x00000117:
		return "StatusBufferAllZeros"
	case 0x00000118:
		return "StatusReparseObject"
	case 0x00000119:
		return "StatusResourceRequirementsChanged"
	case 0x00000120:
		return "StatusTranslationComplete"
	case 0x00000121:
		return "StatusDsMembershipEvaluatedLocally"
	case 0x00000122:
		return "StatusNothingToTerminate"
	case 0x00000123:
		return "StatusProcessNotInJob"
	case 0x00000124:
		return "StatusProcessInJob"
	case 0x00000125:
		return "StatusVolsnapHibernateReady"
	case 0x00000126:
		return "StatusFsfilterOpCompletedSuccessfully"
	case 0x00000127:
		return "StatusInterruptVectorAlreadyConnected"
	case 0x00000128:
		return "StatusInterruptStillConnected"
	case 0x00000129:
		return "StatusProcessCloned"
	case 0x0000012A:
		return "StatusFileLockedWithOnlyReaders"
	case 0x0000012B:
		return "StatusFileLockedWithWriters"
	case 0x00000202:
		return "StatusResourcemanagerReadOnly"
	case 0x00000210:
		return "StatusRingPreviouslyEmpty"
	case 0x00000211:
		return "StatusRingPreviouslyFull"
	case 0x00000212:
		return "StatusRingPreviouslyAboveQuota"
	case 0x00000213:
		return "StatusRingNewlyEmpty"
	case 0x00000214:
		return "StatusRingSignalOppositeEndpoint"
	case 0x00000215:
		return "StatusOplockSwitchedToNewHandle"
	case 0x00000216:
		return "StatusOplockHandleClosed"
	case 0x00000367:
		return "StatusWaitForOplock"
	case 0x00010001:
		return "DbgExceptionHandled"
	case 0x00010002:
		return "DbgContinue"
	case 0x001C0001:
		return "StatusFltIoComplete"
	case 0x003C0001:
		return "StatusDisAttributeBuilt"
	case 0x40000000:
		return "StatusObjectNameExists"
	case 0x40000001:
		return "StatusThreadWasSuspended"
	case 0x40000002:
		return "StatusWorkingSetLimitRange"
	case 0x40000003:
		return "StatusImageNotAtBase"
	case 0x40000004:
		return "StatusRxactStateCreated"
	case 0x40000005:
		return "StatusSegmentNotification"
	case 0x40000006:
		return "StatusLocalUserSessionKey"
	case 0x40000007:
		return "StatusBadCurrentDirectory"
	case 0x40000008:
		return "StatusSerialMoreWrites"
	case 0x40000009:
		return "StatusRegistryRecovered"
	case 0x4000000A:
		return "StatusFtReadRecoveryFromBackup"
	case 0x4000000B:
		return "StatusFtWriteRecovery"
	case 0x4000000C:
		return "StatusSerialCounterTimeout"
	case 0x4000000D:
		return "StatusNullLmPassword"
	case 0x4000000E:
		return "StatusImageMachineTypeMismatch"
	case 0x4000000F:
		return "StatusReceivePartial"
	case 0x40000010:
		return "StatusReceiveExpedited"
	case 0x40000011:
		return "StatusReceivePartialExpedited"
	case 0x40000012:
		return "StatusEventDone"
	case 0x40000013:
		return "StatusEventPending"
	case 0x40000014:
		return "StatusCheckingFileSystem"
	case 0x40000015:
		return "StatusFatalAppExit"
	case 0x40000016:
		return "StatusPredefinedHandle"
	case 0x40000017:
		return "StatusWasUnlocked"
	case 0x40000018:
		return "StatusServiceNotification"
	case 0x40000019:
		return "StatusWasLocked"
	case 0x4000001A:
		return "StatusLogHardError"
	case 0x4000001B:
		return "StatusAlreadyWin32"
	case 0x4000001C:
		return "StatusWx86Unsimulate"
	case 0x4000001D:
		return "StatusWx86Continue"
	case 0x4000001E:
		return "StatusWx86SingleStep"
	case 0x4000001F:
		return "StatusWx86Breakpoint"
	case 0x40000020:
		return "StatusWx86ExceptionContinue"
	case 0x40000021:
		return "StatusWx86ExceptionLastchance"
	case 0x40000022:
		return "StatusWx86ExceptionChain"
	case 0x40000023:
		return "StatusImageMachineTypeMismatchExe"
	case 0x40000024:
		return "StatusNoYieldPerformed"
	case 0x40000025:
		return "StatusTimerResumeIgnored"
	case 0x40000026:
		return "StatusArbitrationUnhandled"
	case 0x40000027:
		return "StatusCardbusNotSupported"
	case 0x40000028:
		return "StatusWx86Createwx86Tib"
	case 0x40000029:
		return "StatusMpProcessorMismatch"
	case 0x4000002A:
		return "StatusHibernated"
	case 0x4000002B:
		return "StatusResumeHibernation"
	case 0x4000002C:
		return "StatusFirmwareUpdated"
	case 0x4000002D:
		return "StatusDriversLeakingLockedPages"
	case 0x4000002E:
		return "StatusMessageRetrieved"
	case 0x4000002F:
		return "StatusSystemPowerstateTransition"
	case 0x40000030:
		return "StatusAlpcCheckCompletionList"
	case 0x40000031:
		return "StatusSystemPowerstateComplexTransition"
	case 0x40000032:
		return "StatusAccessAuditByPolicy"
	case 0x40000033:
		return "StatusAbandonHiberfile"
	case 0x40000034:
		return "StatusBizrulesNotEnabled"
	case 0x40010001:
		return "DbgReplyLater"
	case 0x40010002:
		return "DbgUnableToProvideHandle"
	case 0x40010003:
		return "DbgTerminateThread"
	case 0x40010004:
		return "DbgTerminateProcess"
	case 0x40010005:
		return "DbgControlC"
	case 0x40010006:
		return "DbgPrintexceptionC"
	case 0x40010007:
		return "DbgRipexception"
	case 0x40010008:
		return "DbgControlBreak"
	case 0x40010009:
		return "DbgCommandException"
	case 0x40190001:
		return "StatusHeuristicDamagePossible"
	case 0x80000001:
		return "StatusGuardPageViolation"
	case 0x80000002:
		return "StatusDatatypeMisalignment"
	case 0x80000003:
		return "StatusBreakpoint"
	case 0x80000004:
		return "StatusSingleStep"
	case 0x80000005:
		return "StatusBufferOverflow"
	case 0x80000006:
		return "StatusNoMoreFiles"
	case 0x80000007:
		return "StatusWakeSystemDebugger"
	case 0x8000000A:
		return "StatusHandlesClosed"
	case 0x8000000B:
		return "StatusNoInheritance"
	case 0x8000000C:
		return "StatusGuidSubstitutionMade"
	case 0x8000000D:
		return "StatusPartialCopy"
	case 0x8000000E:
		return "StatusDevicePaperEmpty"
	case 0x8000000F:
		return "StatusDevicePoweredOff"
	case 0x80000010:
		return "StatusDeviceOffLine"
	case 0x80000011:
		return "StatusDeviceBusy"
	case 0x80000012:
		return "StatusNoMoreEas"
	case 0x80000013:
		return "StatusInvalidEaName"
	case 0x80000014:
		return "StatusEaListInconsistent"
	case 0x80000015:
		return "StatusInvalidEaFlag"
	case 0x80000016:
		return "StatusVerifyRequired"
	case 0x80000017:
		return "StatusExtraneousInformation"
	case 0x80000018:
		return "StatusRxactCommitNecessary"
	case 0x8000001A:
		return "StatusNoMoreEntries"
	case 0x8000001B:
		return "StatusFilemarkDetected"
	case 0x8000001C:
		return "StatusMediaChanged"
	case 0x8000001D:
		return "StatusBusReset"
	case 0x8000001E:
		return "StatusEndOfMedia"
	case 0x8000001F:
		return "StatusBeginningOfMedia"
	case 0x80000020:
		return "StatusMediaCheck"
	case 0x80000021:
		return "StatusSetmarkDetected"
	case 0x80000022:
		return "StatusNoDataDetected"
	case 0x80000023:
		return "StatusRedirectorHasOpenHandles"
	case 0x80000024:
		return "StatusServerHasOpenHandles"
	case 0x80000025:
		return "StatusAlreadyDisconnected"
	case 0x80000026:
		return "StatusLongjump"
	case 0x80000027:
		return "StatusCleanerCartridgeInstalled"
	case 0x80000028:
		return "StatusPlugplayQueryVetoed"
	case 0x80000029:
		return "StatusUnwindConsolidate"
	case 0x8000002A:
		return "StatusRegistryHiveRecovered"
	case 0x8000002B:
		return "StatusDllMightBeInsecure"
	case 0x8000002C:
		return "StatusDllMightBeIncompatible"
	case 0x8000002D:
		return "StatusStoppedOnSymlink"
	case 0x8000002E:
		return "StatusCannotGrantRequestedOplock"
	case 0x8000002F:
		return "StatusNoAceCondition"
	case 0x80010001:
		return "DbgExceptionNotHandled"
	case 0x80130001:
		return "StatusClusterNodeAlreadyUp"
	case 0x80130002:
		return "StatusClusterNodeAlreadyDown"
	case 0x80130003:
		return "StatusClusterNetworkAlreadyOnline"
	case 0x80130004:
		return "StatusClusterNetworkAlreadyOffline"
	case 0x80130005:
		return "StatusClusterNodeAlreadyMember"
	case 0x801C0001:
		return "StatusFltBufferTooSmall"
	case 0x80210001:
		return "StatusFvePartialMetadata"
	case 0x80210002:
		return "StatusFveTransientState"
	case 0xC0000001:
		return "StatusUnsuccessful"
	case 0xC0000002:
		return "StatusNotImplemented"
	case 0xC0000003:
		return "StatusInvalidInfoClass"
	case 0xC0000004:
		return "StatusInfoLengthMismatch"
	case 0xC0000005:
		return "StatusAccessViolation"
	case 0xC0000006:
		return "StatusInPageError"
	case 0xC0000007:
		return "StatusPagefileQuota"
	case 0xC0000008:
		return "StatusInvalidHandle"
	case 0xC0000009:
		return "StatusBadInitialStack"
	case 0xC000000A:
		return "StatusBadInitialPc"
	case 0xC000000B:
		return "StatusInvalidCid"
	case 0xC000000C:
		return "StatusTimerNotCanceled"
	case 0xC000000D:
		return "StatusInvalidParameter"
	case 0xC000000E:
		return "StatusNoSuchDevice"
	case 0xC000000F:
		return "StatusNoSuchFile"
	case 0xC0000010:
		return "StatusInvalidDeviceRequest"
	case 0xC0000011:
		return "StatusEndOfFile"
	case 0xC0000012:
		return "StatusWrongVolume"
	case 0xC0000013:
		return "StatusNoMediaInDevice"
	case 0xC0000014:
		return "StatusUnrecognizedMedia"
	case 0xC0000015:
		return "StatusNonexistentSector"
	case 0xC0000016:
		return "StatusMoreProcessingRequired"
	case 0xC0000017:
		return "StatusNoMemory"
	case 0xC0000018:
		return "StatusConflictingAddresses"
	case 0xC0000019:
		return "StatusNotMappedView"
	case 0xC000001A:
		return "StatusUnableToFreeVm"
	case 0xC000001B:
		return "StatusUnableToDeleteSection"
	case 0xC000001C:
		return "StatusInvalidSystemService"
	case 0xC000001D:
		return "StatusIllegalInstruction"
	case 0xC000001E:
		return "StatusInvalidLockSequence"
	case 0xC000001F:
		return "StatusInvalidViewSize"
	case 0xC0000020:
		return "StatusInvalidFileForSection"
	case 0xC0000021:
		return "StatusAlreadyCommitted"
	case 0xC0000022:
		return "StatusAccessDenied"
	case 0xC0000023:
		return "StatusBufferTooSmall"
	case 0xC0000024:
		return "StatusObjectTypeMismatch"
	case 0xC0000025:
		return "StatusNoncontinuableException"
	case 0xC0000026:
		return "StatusInvalidDisposition"
	case 0xC0000027:
		return "StatusUnwind"
	case 0xC0000028:
		return "StatusBadStack"
	case 0xC0000029:
		return "StatusInvalidUnwindTarget"
	case 0xC000002A:
		return "StatusNotLocked"
	case 0xC000002B:
		return "StatusParityError"
	case 0xC000002C:
		return "StatusUnableToDecommitVm"
	case 0xC000002D:
		return "StatusNotCommitted"
	case 0xC000002E:
		return "StatusInvalidPortAttributes"
	case 0xC000002F:
		return "StatusPortMessageTooLong"
	case 0xC0000030:
		return "StatusInvalidParameterMix"
	case 0xC0000031:
		return "StatusInvalidQuotaLower"
	case 0xC0000032:
		return "StatusDiskCorruptError"
	case 0xC0000033:
		return "StatusObjectNameInvalid"
	case 0xC0000034:
		return "StatusObjectNameNotFound"
	case 0xC0000035:
		return "StatusObjectNameCollision"
	case 0xC0000037:
		return "StatusPortDisconnected"
	case 0xC0000038:
		return "StatusDeviceAlreadyAttached"
	case 0xC0000039:
		return "StatusObjectPathInvalid"
	case 0xC000003A:
		return "StatusObjectPathNotFound"
	case 0xC000003B:
		return "StatusObjectPathSyntaxBad"
	case 0xC000003C:
		return "StatusDataOverrun"
	case 0xC000003D:
		return "StatusDataLateError"
	case 0xC000003E:
		return "StatusDataError"
	case 0xC000003F:
		return "StatusCrcError"
	case 0xC0000040:
		return "StatusSectionTooBig"
	case 0xC0000041:
		return "StatusPortConnectionRefused"
	case 0xC0000042:
		return "StatusInvalidPortHandle"
	case 0xC0000043:
		return "StatusSharingViolation"
	case 0xC0000044:
		return "StatusQuotaExceeded"
	case 0xC0000045:
		return "StatusInvalidPageProtection"
	case 0xC0000046:
		return "StatusMutantNotOwned"
	case 0xC0000047:
		return "StatusSemaphoreLimitExceeded"
	case 0xC0000048:
		return "StatusPortAlreadySet"
	case 0xC0000049:
		return "StatusSectionNotImage"
	case 0xC000004A:
		return "StatusSuspendCountExceeded"
	case 0xC000004B:
		return "StatusThreadIsTerminating"
	case 0xC000004C:
		return "StatusBadWorkingSetLimit"
	case 0xC000004D:
		return "StatusIncompatibleFileMap"
	case 0xC000004E:
		return "StatusSectionProtection"
	case 0xC000004F:
		return "StatusEasNotSupported"
	case 0xC0000050:
		return "StatusEaTooLarge"
	case 0xC0000051:
		return "StatusNonexistentEaEntry"
	case 0xC0000052:
		return "StatusNoEasOnFile"
	case 0xC0000053:
		return "StatusEaCorruptError"
	case 0xC0000054:
		return "StatusFileLockConflict"
	case 0xC0000055:
		return "StatusLockNotGranted"
	case 0xC0000056:
		return "StatusDeletePending"
	case 0xC0000057:
		return "StatusCtlFileNotSupported"
	case 0xC0000058:
		return "StatusUnknownRevision"
	case 0xC0000059:
		return "StatusRevisionMismatch"
	case 0xC000005A:
		return "StatusInvalidOwner"
	case 0xC000005B:
		return "StatusInvalidPrimaryGroup"
	case 0xC000005C:
		return "StatusNoImpersonationToken"
	case 0xC000005D:
		return "StatusCantDisableMandatory"
	case 0xC000005E:
		return "StatusNoLogonServers"
	case 0xC000005F:
		return "StatusNoSuchLogonSession"
	case 0xC0000060:
		return "StatusNoSuchPrivilege"
	case 0xC0000061:
		return "StatusPrivilegeNotHeld"
	case 0xC0000062:
		return "StatusInvalidAccountName"
	case 0xC0000063:
		return "StatusUserExists"
	case 0xC0000064:
		return "StatusNoSuchUser"
	case 0xC0000065:
		return "StatusGroupExists"
	case 0xC0000066:
		return "StatusNoSuchGroup"
	case 0xC0000067:
		return "StatusMemberInGroup"
	case 0xC0000068:
		return "StatusMemberNotInGroup"
	case 0xC0000069:
		return "StatusLastAdmin"
	case 0xC000006A:
		return "StatusWrongPassword"
	case 0xC000006B:
		return "StatusIllFormedPassword"
	case 0xC000006C:
		return "StatusPasswordRestriction"
	case 0xC000006D:
		return "StatusLogonFailure"
	case 0xC000006E:
		return "StatusAccountRestriction"
	case 0xC000006F:
		return "StatusInvalidLogonHours"
	case 0xC0000070:
		return "StatusInvalidWorkstation"
	case 0xC0000071:
		return "StatusPasswordExpired"
	case 0xC0000072:
		return "StatusAccountDisabled"
	case 0xC0000073:
		return "StatusNoneMapped"
	case 0xC0000074:
		return "StatusTooManyLuidsRequested"
	case 0xC0000075:
		return "StatusLuidsExhausted"
	case 0xC0000076:
		return "StatusInvalidSubAuthority"
	case 0xC0000077:
		return "StatusInvalidAcl"
	case 0xC0000078:
		return "StatusInvalidSid"
	case 0xC0000079:
		return "StatusInvalidSecurityDescr"
	case 0xC000007A:
		return "StatusProcedureNotFound"
	case 0xC000007B:
		return "StatusInvalidImageFormat"
	case 0xC000007C:
		return "StatusNoToken"
	case 0xC000007D:
		return "StatusBadInheritanceAcl"
	case 0xC000007E:
		return "StatusRangeNotLocked"
	case 0xC000007F:
		return "StatusDiskFull"
	case 0xC0000080:
		return "StatusServerDisabled"
	case 0xC0000081:
		return "StatusServerNotDisabled"
	case 0xC0000082:
		return "StatusTooManyGuidsRequested"
	case 0xC0000083:
		return "StatusGuidsExhausted"
	case 0xC0000084:
		return "StatusInvalidIdAuthority"
	case 0xC0000085:
		return "StatusAgentsExhausted"
	case 0xC0000086:
		return "StatusInvalidVolumeLabel"
	case 0xC0000087:
		return "StatusSectionNotExtended"
	case 0xC0000088:
		return "StatusNotMappedData"
	case 0xC0000089:
		return "StatusResourceDataNotFound"
	case 0xC000008A:
		return "StatusResourceTypeNotFound"
	case 0xC000008B:
		return "StatusResourceNameNotFound"
	case 0xC000008C:
		return "StatusArrayBoundsExceeded"
	case 0xC000008D:
		return "StatusFloatDenormalOperand"
	case 0xC000008E:
		return "StatusFloatDivideByZero"
	case 0xC000008F:
		return "StatusFloatInexactResult"
	case 0xC0000090:
		return "StatusFloatInvalidOperation"
	case 0xC0000091:
		return "StatusFloatOverflow"
	case 0xC0000092:
		return "StatusFloatStackCheck"
	case 0xC0000093:
		return "StatusFloatUnderflow"
	case 0xC0000094:
		return "StatusIntegerDivideByZero"
	case 0xC0000095:
		return "StatusIntegerOverflow"
	case 0xC0000096:
		return "StatusPrivilegedInstruction"
	case 0xC0000097:
		return "StatusTooManyPagingFiles"
	case 0xC0000098:
		return "StatusFileInvalid"
	case 0xC0000099:
		return "StatusAllottedSpaceExceeded"
	case 0xC000009A:
		return "StatusInsufficientResources"
	case 0xC000009B:
		return "StatusDfsExitPathFound"
	case 0xC000009C:
		return "StatusDeviceDataError"
	case 0xC000009D:
		return "StatusDeviceNotConnected"
	case 0xC000009E:
		return "StatusDevicePowerFailure"
	case 0xC000009F:
		return "StatusFreeVmNotAtBase"
	case 0xC00000A0:
		return "StatusMemoryNotAllocated"
	case 0xC00000A1:
		return "StatusWorkingSetQuota"
	case 0xC00000A2:
		return "StatusMediaWriteProtected"
	case 0xC00000A3:
		return "StatusDeviceNotReady"
	case 0xC00000A4:
		return "StatusInvalidGroupAttributes"
	case 0xC00000A5:
		return "StatusBadImpersonationLevel"
	case 0xC00000A6:
		return "StatusCantOpenAnonymous"
	case 0xC00000A7:
		return "StatusBadValidationClass"
	case 0xC00000A8:
		return "StatusBadTokenType"
	case 0xC00000A9:
		return "StatusBadMasterBootRecord"
	case 0xC00000AA:
		return "StatusInstructionMisalignment"
	case 0xC00000AB:
		return "StatusInstanceNotAvailable"
	case 0xC00000AC:
		return "StatusPipeNotAvailable"
	case 0xC00000AD:
		return "StatusInvalidPipeState"
	case 0xC00000AE:
		return "StatusPipeBusy"
	case 0xC00000AF:
		return "StatusIllegalFunction"
	case 0xC00000B0:
		return "StatusPipeDisconnected"
	case 0xC00000B1:
		return "StatusPipeClosing"
	case 0xC00000B2:
		return "StatusPipeConnected"
	case 0xC00000B3:
		return "StatusPipeListening"
	case 0xC00000B4:
		return "StatusInvalidReadMode"
	case 0xC00000B5:
		return "StatusIoTimeout"
	case 0xC00000B6:
		return "StatusFileForcedClosed"
	case 0xC00000B7:
		return "StatusProfilingNotStarted"
	case 0xC00000B8:
		return "StatusProfilingNotStopped"
	case 0xC00000B9:
		return "StatusCouldNotInterpret"
	case 0xC00000BA:
		return "StatusFileIsADirectory"
	case 0xC00000BB:
		return "StatusNotSupported"
	case 0xC00000BC:
		return "StatusRemoteNotListening"
	case 0xC00000BD:
		return "StatusDuplicateName"
	case 0xC00000BE:
		return "StatusBadNetworkPath"
	case 0xC00000BF:
		return "StatusNetworkBusy"
	case 0xC00000C0:
		return "StatusDeviceDoesNotExist"
	case 0xC00000C1:
		return "StatusTooManyCommands"
	case 0xC00000C2:
		return "StatusAdapterHardwareError"
	case 0xC00000C3:
		return "StatusInvalidNetworkResponse"
	case 0xC00000C4:
		return "StatusUnexpectedNetworkError"
	case 0xC00000C5:
		return "StatusBadRemoteAdapter"
	case 0xC00000C6:
		return "StatusPrintQueueFull"
	case 0xC00000C7:
		return "StatusNoSpoolSpace"
	case 0xC00000C8:
		return "StatusPrintCancelled"
	case 0xC00000C9:
		return "StatusNetworkNameDeleted"
	case 0xC00000CA:
		return "StatusNetworkAccessDenied"
	case 0xC00000CB:
		return "StatusBadDeviceType"
	case 0xC00000CC:
		return "StatusBadNetworkName"
	case 0xC00000CD:
		return "StatusTooManyNames"
	case 0xC00000CE:
		return "StatusTooManySessions"
	case 0xC00000CF:
		return "StatusSharingPaused"
	case 0xC00000D0:
		return "StatusRequestNotAccepted"
	case 0xC00000D1:
		return "StatusRedirectorPaused"
	case 0xC00000D2:
		return "StatusNetWriteFault"
	case 0xC00000D3:
		return "StatusProfilingAtLimit"
	case 0xC00000D4:
		return "StatusNotSameDevice"
	case 0xC00000D5:
		return "StatusFileRenamed"
	case 0xC00000D6:
		return "StatusVirtualCircuitClosed"
	case 0xC00000D7:
		return "StatusNoSecurityOnObject"
	case 0xC00000D8:
		return "StatusCantWait"
	case 0xC00000D9:
		return "StatusPipeEmpty"
	case 0xC00000DA:
		return "StatusCantAccessDomainInfo"
	case 0xC00000DB:
		return "StatusCantTerminateSelf"
	case 0xC00000DC:
		return "StatusInvalidServerState"
	case 0xC00000DD:
		return "StatusInvalidDomainState"
	case 0xC00000DE:
		return "StatusInvalidDomainRole"
	case 0xC00000DF:
		return "StatusNoSuchDomain"
	case 0xC00000E0:
		return "StatusDomainExists"
	case 0xC00000E1:
		return "StatusDomainLimitExceeded"
	case 0xC00000E2:
		return "StatusOplockNotGranted"
	case 0xC00000E3:
		return "StatusInvalidOplockProtocol"
	case 0xC00000E4:
		return "StatusInternalDbCorruption"
	case 0xC00000E5:
		return "StatusInternalError"
	case 0xC00000E6:
		return "StatusGenericNotMapped"
	case 0xC00000E7:
		return "StatusBadDescriptorFormat"
	case 0xC00000E8:
		return "StatusInvalidUserBuffer"
	case 0xC00000E9:
		return "StatusUnexpectedIoError"
	case 0xC00000EA:
		return "StatusUnexpectedMmCreateErr"
	case 0xC00000EB:
		return "StatusUnexpectedMmMapError"
	case 0xC00000EC:
		return "StatusUnexpectedMmExtendErr"
	case 0xC00000ED:
		return "StatusNotLogonProcess"
	case 0xC00000EE:
		return "StatusLogonSessionExists"
	case 0xC00000EF:
		return "StatusInvalidParameter1"
	case 0xC00000F0:
		return "StatusInvalidParameter2"
	case 0xC00000F1:
		return "StatusInvalidParameter3"
	case 0xC00000F2:
		return "StatusInvalidParameter4"
	case 0xC00000F3:
		return "StatusInvalidParameter5"
	case 0xC00000F4:
		return "StatusInvalidParameter6"
	case 0xC00000F5:
		return "StatusInvalidParameter7"
	case 0xC00000F6:
		return "StatusInvalidParameter8"
	case 0xC00000F7:
		return "StatusInvalidParameter9"
	case 0xC00000F8:
		return "StatusInvalidParameter10"
	case 0xC00000F9:
		return "StatusInvalidParameter11"
	case 0xC00000FA:
		return "StatusInvalidParameter12"
	case 0xC00000FB:
		return "StatusRedirectorNotStarted"
	case 0xC00000FC:
		return "StatusRedirectorStarted"
	case 0xC00000FD:
		return "StatusStackOverflow"
	case 0xC00000FE:
		return "StatusNoSuchPackage"
	case 0xC00000FF:
		return "StatusBadFunctionTable"
	case 0xC0000100:
		return "StatusVariableNotFound"
	case 0xC0000101:
		return "StatusDirectoryNotEmpty"
	case 0xC0000102:
		return "StatusFileCorruptError"
	case 0xC0000103:
		return "StatusNotADirectory"
	case 0xC0000104:
		return "StatusBadLogonSessionState"
	case 0xC0000105:
		return "StatusLogonSessionCollision"
	case 0xC0000106:
		return "StatusNameTooLong"
	case 0xC0000107:
		return "StatusFilesOpen"
	case 0xC0000108:
		return "StatusConnectionInUse"
	case 0xC0000109:
		return "StatusMessageNotFound"
	case 0xC000010A:
		return "StatusProcessIsTerminating"
	case 0xC000010B:
		return "StatusInvalidLogonType"
	case 0xC000010C:
		return "StatusNoGuidTranslation"
	case 0xC000010D:
		return "StatusCannotImpersonate"
	case 0xC000010E:
		return "StatusImageAlreadyLoaded"
	case 0xC000010F:
		return "StatusAbiosNotPresent"
	case 0xC0000110:
		return "StatusAbiosLidNotExist"
	case 0xC0000111:
		return "StatusAbiosLidAlreadyOwned"
	case 0xC0000112:
		return "StatusAbiosNotLidOwner"
	case 0xC0000113:
		return "StatusAbiosInvalidCommand"
	case 0xC0000114:
		return "StatusAbiosInvalidLid"
	case 0xC0000115:
		return "StatusAbiosSelectorNotAvailable"
	case 0xC0000116:
		return "StatusAbiosInvalidSelector"
	case 0xC0000117:
		return "StatusNoLdt"
	case 0xC0000118:
		return "StatusInvalidLdtSize"
	case 0xC0000119:
		return "StatusInvalidLdtOffset"
	case 0xC000011A:
		return "StatusInvalidLdtDescriptor"
	case 0xC000011B:
		return "StatusInvalidImageNeFormat"
	case 0xC000011C:
		return "StatusRxactInvalidState"
	case 0xC000011D:
		return "StatusRxactCommitFailure"
	case 0xC000011E:
		return "StatusMappedFileSizeZero"
	case 0xC000011F:
		return "StatusTooManyOpenedFiles"
	case 0xC0000120:
		return "StatusCancelled"
	case 0xC0000121:
		return "StatusCannotDelete"
	case 0xC0000122:
		return "StatusInvalidComputerName"
	case 0xC0000123:
		return "StatusFileDeleted"
	case 0xC0000124:
		return "StatusSpecialAccount"
	case 0xC0000125:
		return "StatusSpecialGroup"
	case 0xC0000126:
		return "StatusSpecialUser"
	case 0xC0000127:
		return "StatusMembersPrimaryGroup"
	case 0xC0000128:
		return "StatusFileClosed"
	case 0xC0000129:
		return "StatusTooManyThreads"
	case 0xC000012A:
		return "StatusThreadNotInProcess"
	case 0xC000012B:
		return "StatusTokenAlreadyInUse"
	case 0xC000012C:
		return "StatusPagefileQuotaExceeded"
	case 0xC000012D:
		return "StatusCommitmentLimit"
	case 0xC000012E:
		return "StatusInvalidImageLeFormat"
	case 0xC000012F:
		return "StatusInvalidImageNotMz"
	case 0xC0000130:
		return "StatusInvalidImageProtect"
	case 0xC0000131:
		return "StatusInvalidImageWin16"
	case 0xC0000132:
		return "StatusLogonServerConflict"
	case 0xC0000133:
		return "StatusTimeDifferenceAtDc"
	case 0xC0000134:
		return "StatusSynchronizationRequired"
	case 0xC0000135:
		return "StatusDllNotFound"
	case 0xC0000136:
		return "StatusOpenFailed"
	case 0xC0000137:
		return "StatusIoPrivilegeFailed"
	case 0xC0000138:
		return "StatusOrdinalNotFound"
	case 0xC0000139:
		return "StatusEntrypointNotFound"
	case 0xC000013A:
		return "StatusControlCExit"
	case 0xC000013B:
		return "StatusLocalDisconnect"
	case 0xC000013C:
		return "StatusRemoteDisconnect"
	case 0xC000013D:
		return "StatusRemoteResources"
	case 0xC000013E:
		return "StatusLinkFailed"
	case 0xC000013F:
		return "StatusLinkTimeout"
	case 0xC0000140:
		return "StatusInvalidConnection"
	case 0xC0000141:
		return "StatusInvalidAddress"
	case 0xC0000142:
		return "StatusDllInitFailed"
	case 0xC0000143:
		return "StatusMissingSystemfile"
	case 0xC0000144:
		return "StatusUnhandledException"
	case 0xC0000145:
		return "StatusAppInitFailure"
	case 0xC0000146:
		return "StatusPagefileCreateFailed"
	case 0xC0000147:
		return "StatusNoPagefile"
	case 0xC0000148:
		return "StatusInvalidLevel"
	case 0xC0000149:
		return "StatusWrongPasswordCore"
	case 0xC000014A:
		return "StatusIllegalFloatContext"
	case 0xC000014B:
		return "StatusPipeBroken"
	case 0xC000014C:
		return "StatusRegistryCorrupt"
	case 0xC000014D:
		return "StatusRegistryIoFailed"
	case 0xC000014E:
		return "StatusNoEventPair"
	case 0xC000014F:
		return "StatusUnrecognizedVolume"
	case 0xC0000150:
		return "StatusSerialNoDeviceInited"
	case 0xC0000151:
		return "StatusNoSuchAlias"
	case 0xC0000152:
		return "StatusMemberNotInAlias"
	case 0xC0000153:
		return "StatusMemberInAlias"
	case 0xC0000154:
		return "StatusAliasExists"
	case 0xC0000155:
		return "StatusLogonNotGranted"
	case 0xC0000156:
		return "StatusTooManySecrets"
	case 0xC0000157:
		return "StatusSecretTooLong"
	case 0xC0000158:
		return "StatusInternalDbError"
	case 0xC0000159:
		return "StatusFullscreenMode"
	case 0xC000015A:
		return "StatusTooManyContextIds"
	case 0xC000015B:
		return "StatusLogonTypeNotGranted"
	case 0xC000015C:
		return "StatusNotRegistryFile"
	case 0xC000015D:
		return "StatusNtCrossEncryptionRequired"
	case 0xC000015E:
		return "StatusDomainCtrlrConfigError"
	case 0xC000015F:
		return "StatusFtMissingMember"
	case 0xC0000160:
		return "StatusIllFormedServiceEntry"
	case 0xC0000161:
		return "StatusIllegalCharacter"
	case 0xC0000162:
		return "StatusUnmappableCharacter"
	case 0xC0000163:
		return "StatusUndefinedCharacter"
	case 0xC0000164:
		return "StatusFloppyVolume"
	case 0xC0000165:
		return "StatusFloppyIdMarkNotFound"
	case 0xC0000166:
		return "StatusFloppyWrongCylinder"
	case 0xC0000167:
		return "StatusFloppyUnknownError"
	case 0xC0000168:
		return "StatusFloppyBadRegisters"
	case 0xC0000169:
		return "StatusDiskRecalibrateFailed"
	case 0xC000016A:
		return "StatusDiskOperationFailed"
	case 0xC000016B:
		return "StatusDiskResetFailed"
	case 0xC000016C:
		return "StatusSharedIrqBusy"
	case 0xC000016D:
		return "StatusFtOrphaning"
	case 0xC000016E:
		return "StatusBiosFailedToConnectInterrupt"
	case 0xC0000172:
		return "StatusPartitionFailure"
	case 0xC0000173:
		return "StatusInvalidBlockLength"
	case 0xC0000174:
		return "StatusDeviceNotPartitioned"
	case 0xC0000175:
		return "StatusUnableToLockMedia"
	case 0xC0000176:
		return "StatusUnableToUnloadMedia"
	case 0xC0000177:
		return "StatusEomOverflow"
	case 0xC0000178:
		return "StatusNoMedia"
	case 0xC000017A:
		return "StatusNoSuchMember"
	case 0xC000017B:
		return "StatusInvalidMember"
	case 0xC000017C:
		return "StatusKeyDeleted"
	case 0xC000017D:
		return "StatusNoLogSpace"
	case 0xC000017E:
		return "StatusTooManySids"
	case 0xC000017F:
		return "StatusLmCrossEncryptionRequired"
	case 0xC0000180:
		return "StatusKeyHasChildren"
	case 0xC0000181:
		return "StatusChildMustBeVolatile"
	case 0xC0000182:
		return "StatusDeviceConfigurationError"
	case 0xC0000183:
		return "StatusDriverInternalError"
	case 0xC0000184:
		return "StatusInvalidDeviceState"
	case 0xC0000185:
		return "StatusIoDeviceError"
	case 0xC0000186:
		return "StatusDeviceProtocolError"
	case 0xC0000187:
		return "StatusBackupController"
	case 0xC0000188:
		return "StatusLogFileFull"
	case 0xC0000189:
		return "StatusTooLate"
	case 0xC000018A:
		return "StatusNoTrustLsaSecret"
	case 0xC000018B:
		return "StatusNoTrustSamAccount"
	case 0xC000018C:
		return "StatusTrustedDomainFailure"
	case 0xC000018D:
		return "StatusTrustedRelationshipFailure"
	case 0xC000018E:
		return "StatusEventlogFileCorrupt"
	case 0xC000018F:
		return "StatusEventlogCantStart"
	case 0xC0000190:
		return "StatusTrustFailure"
	case 0xC0000191:
		return "StatusMutantLimitExceeded"
	case 0xC0000192:
		return "StatusNetlogonNotStarted"
	case 0xC0000193:
		return "StatusAccountExpired"
	case 0xC0000194:
		return "StatusPossibleDeadlock"
	case 0xC0000195:
		return "StatusNetworkCredentialConflict"
	case 0xC0000196:
		return "StatusRemoteSessionLimit"
	case 0xC0000197:
		return "StatusEventlogFileChanged"
	case 0xC0000198:
		return "StatusNologonInterdomainTrustAccount"
	case 0xC0000199:
		return "StatusNologonWorkstationTrustAccount"
	case 0xC000019A:
		return "StatusNologonServerTrustAccount"
	case 0xC000019B:
		return "StatusDomainTrustInconsistent"
	case 0xC000019C:
		return "StatusFsDriverRequired"
	case 0xC000019D:
		return "StatusImageAlreadyLoadedAsDll"
	case 0xC000019E:
		return "StatusIncompatibleWithGlobalShortNameRegistrySetting"
	case 0xC000019F:
		return "StatusShortNamesNotEnabledOnVolume"
	case 0xC00001A0:
		return "StatusSecurityStreamIsInconsistent"
	case 0xC00001A1:
		return "StatusInvalidLockRange"
	case 0xC00001A2:
		return "StatusInvalidAceCondition"
	case 0xC00001A3:
		return "StatusImageSubsystemNotPresent"
	case 0xC00001A4:
		return "StatusNotificationGuidAlreadyDefined"
	case 0xC0000201:
		return "StatusNetworkOpenRestriction"
	case 0xC0000202:
		return "StatusNoUserSessionKey"
	case 0xC0000203:
		return "StatusUserSessionDeleted"
	case 0xC0000204:
		return "StatusResourceLangNotFound"
	case 0xC0000205:
		return "StatusInsuffServerResources"
	case 0xC0000206:
		return "StatusInvalidBufferSize"
	case 0xC0000207:
		return "StatusInvalidAddressComponent"
	case 0xC0000208:
		return "StatusInvalidAddressWildcard"
	case 0xC0000209:
		return "StatusTooManyAddresses"
	case 0xC000020A:
		return "StatusAddressAlreadyExists"
	case 0xC000020B:
		return "StatusAddressClosed"
	case 0xC000020C:
		return "StatusConnectionDisconnected"
	case 0xC000020D:
		return "StatusConnectionReset"
	case 0xC000020E:
		return "StatusTooManyNodes"
	case 0xC000020F:
		return "StatusTransactionAborted"
	case 0xC0000210:
		return "StatusTransactionTimedOut"
	case 0xC0000211:
		return "StatusTransactionNoRelease"
	case 0xC0000212:
		return "StatusTransactionNoMatch"
	case 0xC0000213:
		return "StatusTransactionResponded"
	case 0xC0000214:
		return "StatusTransactionInvalidId"
	case 0xC0000215:
		return "StatusTransactionInvalidType"
	case 0xC0000216:
		return "StatusNotServerSession"
	case 0xC0000217:
		return "StatusNotClientSession"
	case 0xC0000218:
		return "StatusCannotLoadRegistryFile"
	case 0xC0000219:
		return "StatusDebugAttachFailed"
	case 0xC000021A:
		return "StatusSystemProcessTerminated"
	case 0xC000021B:
		return "StatusDataNotAccepted"
	case 0xC000021C:
		return "StatusNoBrowserServersFound"
	case 0xC000021D:
		return "StatusVdmHardError"
	case 0xC000021E:
		return "StatusDriverCancelTimeout"
	case 0xC000021F:
		return "StatusReplyMessageMismatch"
	case 0xC0000220:
		return "StatusMappedAlignment"
	case 0xC0000221:
		return "StatusImageChecksumMismatch"
	case 0xC0000222:
		return "StatusLostWritebehindData"
	case 0xC0000223:
		return "StatusClientServerParametersInvalid"
	case 0xC0000224:
		return "StatusPasswordMustChange"
	case 0xC0000225:
		return "StatusNotFound"
	case 0xC0000226:
		return "StatusNotTinyStream"
	case 0xC0000227:
		return "StatusRecoveryFailure"
	case 0xC0000228:
		return "StatusStackOverflowRead"
	case 0xC0000229:
		return "StatusFailCheck"
	case 0xC000022A:
		return "StatusDuplicateObjectid"
	case 0xC000022B:
		return "StatusObjectidExists"
	case 0xC000022C:
		return "StatusConvertToLarge"
	case 0xC000022D:
		return "StatusRetry"
	case 0xC000022E:
		return "StatusFoundOutOfScope"
	case 0xC000022F:
		return "StatusAllocateBucket"
	case 0xC0000230:
		return "StatusPropsetNotFound"
	case 0xC0000231:
		return "StatusMarshallOverflow"
	case 0xC0000232:
		return "StatusInvalidVariant"
	case 0xC0000233:
		return "StatusDomainControllerNotFound"
	case 0xC0000234:
		return "StatusAccountLockedOut"
	case 0xC0000235:
		return "StatusHandleNotClosable"
	case 0xC0000236:
		return "StatusConnectionRefused"
	case 0xC0000237:
		return "StatusGracefulDisconnect"
	case 0xC0000238:
		return "StatusAddressAlreadyAssociated"
	case 0xC0000239:
		return "StatusAddressNotAssociated"
	case 0xC000023A:
		return "StatusConnectionInvalid"
	case 0xC000023B:
		return "StatusConnectionActive"
	case 0xC000023C:
		return "StatusNetworkUnreachable"
	case 0xC000023D:
		return "StatusHostUnreachable"
	case 0xC000023E:
		return "StatusProtocolUnreachable"
	case 0xC000023F:
		return "StatusPortUnreachable"
	case 0xC0000240:
		return "StatusRequestAborted"
	case 0xC0000241:
		return "StatusConnectionAborted"
	case 0xC0000242:
		return "StatusBadCompressionBuffer"
	case 0xC0000243:
		return "StatusUserMappedFile"
	case 0xC0000244:
		return "StatusAuditFailed"
	case 0xC0000245:
		return "StatusTimerResolutionNotSet"
	case 0xC0000246:
		return "StatusConnectionCountLimit"
	case 0xC0000247:
		return "StatusLoginTimeRestriction"
	case 0xC0000248:
		return "StatusLoginWkstaRestriction"
	case 0xC0000249:
		return "StatusImageMpUpMismatch"
	case 0xC0000250:
		return "StatusInsufficientLogonInfo"
	case 0xC0000251:
		return "StatusBadDllEntrypoint"
	case 0xC0000252:
		return "StatusBadServiceEntrypoint"
	case 0xC0000253:
		return "StatusLpcReplyLost"
	case 0xC0000254:
		return "StatusIpAddressConflict1"
	case 0xC0000255:
		return "StatusIpAddressConflict2"
	case 0xC0000256:
		return "StatusRegistryQuotaLimit"
	case 0xC0000257:
		return "StatusPathNotCovered"
	case 0xC0000258:
		return "StatusNoCallbackActive"
	case 0xC0000259:
		return "StatusLicenseQuotaExceeded"
	case 0xC000025A:
		return "StatusPwdTooShort"
	case 0xC000025B:
		return "StatusPwdTooRecent"
	case 0xC000025C:
		return "StatusPwdHistoryConflict"
	case 0xC000025E:
		return "StatusPlugplayNoDevice"
	case 0xC000025F:
		return "StatusUnsupportedCompression"
	case 0xC0000260:
		return "StatusInvalidHwProfile"
	case 0xC0000261:
		return "StatusInvalidPlugplayDevicePath"
	case 0xC0000262:
		return "StatusDriverOrdinalNotFound"
	case 0xC0000263:
		return "StatusDriverEntrypointNotFound"
	case 0xC0000264:
		return "StatusResourceNotOwned"
	case 0xC0000265:
		return "StatusTooManyLinks"
	case 0xC0000266:
		return "StatusQuotaListInconsistent"
	case 0xC0000267:
		return "StatusFileIsOffline"
	case 0xC0000268:
		return "StatusEvaluationExpiration"
	case 0xC0000269:
		return "StatusIllegalDllRelocation"
	case 0xC000026A:
		return "StatusLicenseViolation"
	case 0xC000026B:
		return "StatusDllInitFailedLogoff"
	case 0xC000026C:
		return "StatusDriverUnableToLoad"
	case 0xC000026D:
		return "StatusDfsUnavailable"
	case 0xC000026E:
		return "StatusVolumeDismounted"
	case 0xC000026F:
		return "StatusWx86InternalError"
	case 0xC0000270:
		return "StatusWx86FloatStackCheck"
	case 0xC0000271:
		return "StatusValidateContinue"
	case 0xC0000272:
		return "StatusNoMatch"
	case 0xC0000273:
		return "StatusNoMoreMatches"
	case 0xC0000275:
		return "StatusNotAReparsePoint"
	case 0xC0000276:
		return "StatusIoReparseTagInvalid"
	case 0xC0000277:
		return "StatusIoReparseTagMismatch"
	case 0xC0000278:
		return "StatusIoReparseDataInvalid"
	case 0xC0000279:
		return "StatusIoReparseTagNotHandled"
	case 0xC0000280:
		return "StatusReparsePointNotResolved"
	case 0xC0000281:
		return "StatusDirectoryIsAReparsePoint"
	case 0xC0000282:
		return "StatusRangeListConflict"
	case 0xC0000283:
		return "StatusSourceElementEmpty"
	case 0xC0000284:
		return "StatusDestinationElementFull"
	case 0xC0000285:
		return "StatusIllegalElementAddress"
	case 0xC0000286:
		return "StatusMagazineNotPresent"
	case 0xC0000287:
		return "StatusReinitializationNeeded"
	case 0x80000288:
		return "StatusDeviceRequiresCleaning"
	case 0x80000289:
		return "StatusDeviceDoorOpen"
	case 0xC000028A:
		return "StatusEncryptionFailed"
	case 0xC000028B:
		return "StatusDecryptionFailed"
	case 0xC000028C:
		return "StatusRangeNotFound"
	case 0xC000028D:
		return "StatusNoRecoveryPolicy"
	case 0xC000028E:
		return "StatusNoEfs"
	case 0xC000028F:
		return "StatusWrongEfs"
	case 0xC0000290:
		return "StatusNoUserKeys"
	case 0xC0000291:
		return "StatusFileNotEncrypted"
	case 0xC0000292:
		return "StatusNotExportFormat"
	case 0xC0000293:
		return "StatusFileEncrypted"
	case 0x40000294:
		return "StatusWakeSystem"
	case 0xC0000295:
		return "StatusWmiGuidNotFound"
	case 0xC0000296:
		return "StatusWmiInstanceNotFound"
	case 0xC0000297:
		return "StatusWmiItemidNotFound"
	case 0xC0000298:
		return "StatusWmiTryAgain"
	case 0xC0000299:
		return "StatusSharedPolicy"
	case 0xC000029A:
		return "StatusPolicyObjectNotFound"
	case 0xC000029B:
		return "StatusPolicyOnlyInDs"
	case 0xC000029C:
		return "StatusVolumeNotUpgraded"
	case 0xC000029D:
		return "StatusRemoteStorageNotActive"
	case 0xC000029E:
		return "StatusRemoteStorageMediaError"
	case 0xC000029F:
		return "StatusNoTrackingService"
	case 0xC00002A0:
		return "StatusServerSidMismatch"
	case 0xC00002A1:
		return "StatusDsNoAttributeOrValue"
	case 0xC00002A2:
		return "StatusDsInvalidAttributeSyntax"
	case 0xC00002A3:
		return "StatusDsAttributeTypeUndefined"
	case 0xC00002A4:
		return "StatusDsAttributeOrValueExists"
	case 0xC00002A5:
		return "StatusDsBusy"
	case 0xC00002A6:
		return "StatusDsUnavailable"
	case 0xC00002A7:
		return "StatusDsNoRidsAllocated"
	case 0xC00002A8:
		return "StatusDsNoMoreRids"
	case 0xC00002A9:
		return "StatusDsIncorrectRoleOwner"
	case 0xC00002AA:
		return "StatusDsRidmgrInitError"
	case 0xC00002AB:
		return "StatusDsObjClassViolation"
	case 0xC00002AC:
		return "StatusDsCantOnNonLeaf"
	case 0xC00002AD:
		return "StatusDsCantOnRdn"
	case 0xC00002AE:
		return "StatusDsCantModObjClass"
	case 0xC00002AF:
		return "StatusDsCrossDomMoveFailed"
	case 0xC00002B0:
		return "StatusDsGcNotAvailable"
	case 0xC00002B1:
		return "StatusDirectoryServiceRequired"
	case 0xC00002B2:
		return "StatusReparseAttributeConflict"
	case 0xC00002B3:
		return "StatusCantEnableDenyOnly"
	case 0xC00002B4:
		return "StatusFloatMultipleFaults"
	case 0xC00002B5:
		return "StatusFloatMultipleTraps"
	case 0xC00002B6:
		return "StatusDeviceRemoved"
	case 0xC00002B7:
		return "StatusJournalDeleteInProgress"
	case 0xC00002B8:
		return "StatusJournalNotActive"
	case 0xC00002B9:
		return "StatusNointerface"
	case 0xC00002C1:
		return "StatusDsAdminLimitExceeded"
	case 0xC00002C2:
		return "StatusDriverFailedSleep"
	case 0xC00002C3:
		return "StatusMutualAuthenticationFailed"
	case 0xC00002C4:
		return "StatusCorruptSystemFile"
	case 0xC00002C5:
		return "StatusDatatypeMisalignmentError"
	case 0xC00002C6:
		return "StatusWmiReadOnly"
	case 0xC00002C7:
		return "StatusWmiSetFailure"
	case 0xC00002C8:
		return "StatusCommitmentMinimum"
	case 0xC00002C9:
		return "StatusRegNatConsumption"
	case 0xC00002CA:
		return "StatusTransportFull"
	case 0xC00002CB:
		return "StatusDsSamInitFailure"
	case 0xC00002CC:
		return "StatusOnlyIfConnected"
	case 0xC00002CD:
		return "StatusDsSensitiveGroupViolation"
	case 0xC00002CE:
		return "StatusPnpRestartEnumeration"
	case 0xC00002CF:
		return "StatusJournalEntryDeleted"
	case 0xC00002D0:
		return "StatusDsCantModPrimarygroupid"
	case 0xC00002D1:
		return "StatusSystemImageBadSignature"
	case 0xC00002D2:
		return "StatusPnpRebootRequired"
	case 0xC00002D3:
		return "StatusPowerStateInvalid"
	case 0xC00002D4:
		return "StatusDsInvalidGroupType"
	case 0xC00002D5:
		return "StatusDsNoNestGlobalgroupInMixeddomain"
	case 0xC00002D6:
		return "StatusDsNoNestLocalgroupInMixeddomain"
	case 0xC00002D7:
		return "StatusDsGlobalCantHaveLocalMember"
	case 0xC00002D8:
		return "StatusDsGlobalCantHaveUniversalMember"
	case 0xC00002D9:
		return "StatusDsUniversalCantHaveLocalMember"
	case 0xC00002DA:
		return "StatusDsGlobalCantHaveCrossdomainMember"
	case 0xC00002DB:
		return "StatusDsLocalCantHaveCrossdomainLocalMember"
	case 0xC00002DC:
		return "StatusDsHavePrimaryMembers"
	case 0xC00002DD:
		return "StatusWmiNotSupported"
	case 0xC00002DE:
		return "StatusInsufficientPower"
	case 0xC00002DF:
		return "StatusSamNeedBootkeyPassword"
	case 0xC00002E0:
		return "StatusSamNeedBootkeyFloppy"
	case 0xC00002E1:
		return "StatusDsCantStart"
	case 0xC00002E2:
		return "StatusDsInitFailure"
	case 0xC00002E3:
		return "StatusSamInitFailure"
	case 0xC00002E4:
		return "StatusDsGcRequired"
	case 0xC00002E5:
		return "StatusDsLocalMemberOfLocalOnly"
	case 0xC00002E6:
		return "StatusDsNoFpoInUniversalGroups"
	case 0xC00002E7:
		return "StatusDsMachineAccountQuotaExceeded"
	case 0xC00002E8:
		return "StatusMultipleFaultViolation"
	case 0xC00002E9:
		return "StatusCurrentDomainNotAllowed"
	case 0xC00002EA:
		return "StatusCannotMake"
	case 0xC00002EB:
		return "StatusSystemShutdown"
	case 0xC00002EC:
		return "StatusDsInitFailureConsole"
	case 0xC00002ED:
		return "StatusDsSamInitFailureConsole"
	case 0xC00002EE:
		return "StatusUnfinishedContextDeleted"
	case 0xC00002EF:
		return "StatusNoTgtReply"
	case 0xC00002F0:
		return "StatusObjectidNotFound"
	case 0xC00002F1:
		return "StatusNoIpAddresses"
	case 0xC00002F2:
		return "StatusWrongCredentialHandle"
	case 0xC00002F3:
		return "StatusCryptoSystemInvalid"
	case 0xC00002F4:
		return "StatusMaxReferralsExceeded"
	case 0xC00002F5:
		return "StatusMustBeKdc"
	case 0xC00002F6:
		return "StatusStrongCryptoNotSupported"
	case 0xC00002F7:
		return "StatusTooManyPrincipals"
	case 0xC00002F8:
		return "StatusNoPaData"
	case 0xC00002F9:
		return "StatusPkinitNameMismatch"
	case 0xC00002FA:
		return "StatusSmartcardLogonRequired"
	case 0xC00002FB:
		return "StatusKdcInvalidRequest"
	case 0xC00002FC:
		return "StatusKdcUnableToRefer"
	case 0xC00002FD:
		return "StatusKdcUnknownEtype"
	case 0xC00002FE:
		return "StatusShutdownInProgress"
	case 0xC00002FF:
		return "StatusServerShutdownInProgress"
	case 0xC0000300:
		return "StatusNotSupportedOnSbs"
	case 0xC0000301:
		return "StatusWmiGuidDisconnected"
	case 0xC0000302:
		return "StatusWmiAlreadyDisabled"
	case 0xC0000303:
		return "StatusWmiAlreadyEnabled"
	case 0xC0000304:
		return "StatusMftTooFragmented"
	case 0xC0000305:
		return "StatusCopyProtectionFailure"
	case 0xC0000306:
		return "StatusCssAuthenticationFailure"
	case 0xC0000307:
		return "StatusCssKeyNotPresent"
	case 0xC0000308:
		return "StatusCssKeyNotEstablished"
	case 0xC0000309:
		return "StatusCssScrambledSector"
	case 0xC000030A:
		return "StatusCssRegionMismatch"
	case 0xC000030B:
		return "StatusCssResetsExhausted"
	case 0xC0000320:
		return "StatusPkinitFailure"
	case 0xC0000321:
		return "StatusSmartcardSubsystemFailure"
	case 0xC0000322:
		return "StatusNoKerbKey"
	case 0xC0000350:
		return "StatusHostDown"
	case 0xC0000351:
		return "StatusUnsupportedPreauth"
	case 0xC0000352:
		return "StatusEfsAlgBlobTooBig"
	case 0xC0000353:
		return "StatusPortNotSet"
	case 0xC0000354:
		return "StatusDebuggerInactive"
	case 0xC0000355:
		return "StatusDsVersionCheckFailure"
	case 0xC0000356:
		return "StatusAuditingDisabled"
	case 0xC0000357:
		return "StatusPrent4MachineAccount"
	case 0xC0000358:
		return "StatusDsAgCantHaveUniversalMember"
	case 0xC0000359:
		return "StatusInvalidImageWin32"
	case 0xC000035A:
		return "StatusInvalidImageWin64"
	case 0xC000035B:
		return "StatusBadBindings"
	case 0xC000035C:
		return "StatusNetworkSessionExpired"
	case 0xC000035D:
		return "StatusApphelpBlock"
	case 0xC000035E:
		return "StatusAllSidsFiltered"
	case 0xC000035F:
		return "StatusNotSafeModeDriver"
	case 0xC0000361:
		return "StatusAccessDisabledByPolicyDefault"
	case 0xC0000362:
		return "StatusAccessDisabledByPolicyPath"
	case 0xC0000363:
		return "StatusAccessDisabledByPolicyPublisher"
	case 0xC0000364:
		return "StatusAccessDisabledByPolicyOther"
	case 0xC0000365:
		return "StatusFailedDriverEntry"
	case 0xC0000366:
		return "StatusDeviceEnumerationError"
	case 0xC0000368:
		return "StatusMountPointNotResolved"
	case 0xC0000369:
		return "StatusInvalidDeviceObjectParameter"
	case 0xC000036A:
		return "StatusMcaOccured"
	case 0xC000036B:
		return "StatusDriverBlockedCritical"
	case 0xC000036C:
		return "StatusDriverBlocked"
	case 0xC000036D:
		return "StatusDriverDatabaseError"
	case 0xC000036E:
		return "StatusSystemHiveTooLarge"
	case 0xC000036F:
		return "StatusInvalidImportOfNonDll"
	case 0x40000370:
		return "StatusDsShuttingDown"
	case 0xC0000371:
		return "StatusNoSecrets"
	case 0xC0000372:
		return "StatusAccessDisabledNoSaferUiByPolicy"
	case 0xC0000373:
		return "StatusFailedStackSwitch"
	case 0xC0000374:
		return "StatusHeapCorruption"
	case 0xC0000380:
		return "StatusSmartcardWrongPin"
	case 0xC0000381:
		return "StatusSmartcardCardBlocked"
	case 0xC0000382:
		return "StatusSmartcardCardNotAuthenticated"
	case 0xC0000383:
		return "StatusSmartcardNoCard"
	case 0xC0000384:
		return "StatusSmartcardNoKeyContainer"
	case 0xC0000385:
		return "StatusSmartcardNoCertificate"
	case 0xC0000386:
		return "StatusSmartcardNoKeyset"
	case 0xC0000387:
		return "StatusSmartcardIoError"
	case 0xC0000388:
		return "StatusDowngradeDetected"
	case 0xC0000389:
		return "StatusSmartcardCertRevoked"
	case 0xC000038A:
		return "StatusIssuingCaUntrusted"
	case 0xC000038B:
		return "StatusRevocationOfflineC"
	case 0xC000038C:
		return "StatusPkinitClientFailure"
	case 0xC000038D:
		return "StatusSmartcardCertExpired"
	case 0xC000038E:
		return "StatusDriverFailedPriorUnload"
	case 0xC000038F:
		return "StatusSmartcardSilentContext"
	case 0xC0000401:
		return "StatusPerUserTrustQuotaExceeded"
	case 0xC0000402:
		return "StatusAllUserTrustQuotaExceeded"
	case 0xC0000403:
		return "StatusUserDeleteTrustQuotaExceeded"
	case 0xC0000404:
		return "StatusDsNameNotUnique"
	case 0xC0000405:
		return "StatusDsDuplicateIdFound"
	case 0xC0000406:
		return "StatusDsGroupConversionError"
	case 0xC0000407:
		return "StatusVolsnapPrepareHibernate"
	case 0xC0000408:
		return "StatusUser2UserRequired"
	case 0xC0000409:
		return "StatusStackBufferOverrun"
	case 0xC000040A:
		return "StatusNoS4UProtSupport"
	case 0xC000040B:
		return "StatusCrossrealmDelegationFailure"
	case 0xC000040C:
		return "StatusRevocationOfflineKdc"
	case 0xC000040D:
		return "StatusIssuingCaUntrustedKdc"
	case 0xC000040E:
		return "StatusKdcCertExpired"
	case 0xC000040F:
		return "StatusKdcCertRevoked"
	case 0xC0000410:
		return "StatusParameterQuotaExceeded"
	case 0xC0000411:
		return "StatusHibernationFailure"
	case 0xC0000412:
		return "StatusDelayLoadFailed"
	case 0xC0000413:
		return "StatusAuthenticationFirewallFailed"
	case 0xC0000414:
		return "StatusVdmDisallowed"
	case 0xC0000415:
		return "StatusHungDisplayDriverThread"
	case 0xC0000416:
		return "StatusInsufficientResourceForSpecifiedSharedSectionSize"
	case 0xC0000417:
		return "StatusInvalidCruntimeParameter"
	case 0xC0000418:
		return "StatusNtlmBlocked"
	case 0xC0000419:
		return "StatusDsSrcSidExistsInForest"
	case 0xC000041A:
		return "StatusDsDomainNameExistsInForest"
	case 0xC000041B:
		return "StatusDsFlatNameExistsInForest"
	case 0xC000041C:
		return "StatusInvalidUserPrincipalName"
	case 0xC000041D:
		return "StatusFatalUserCallbackException"
	case 0xC0000420:
		return "StatusAssertionFailure"
	case 0xC0000421:
		return "StatusVerifierStop"
	case 0xC0000423:
		return "StatusCallbackPopStack"
	case 0xC0000424:
		return "StatusIncompatibleDriverBlocked"
	case 0xC0000425:
		return "StatusHiveUnloaded"
	case 0xC0000426:
		return "StatusCompressionDisabled"
	case 0xC0000427:
		return "StatusFileSystemLimitation"
	case 0xC0000428:
		return "StatusInvalidImageHash"
	case 0xC0000429:
		return "StatusNotCapable"
	case 0xC000042A:
		return "StatusRequestOutOfSequence"
	case 0xC000042B:
		return "StatusImplementationLimit"
	case 0xC000042C:
		return "StatusElevationRequired"
	case 0xC000042D:
		return "StatusNoSecurityContext"
	case 0xC000042F:
		return "StatusPku2UCertFailure"
	case 0xC0000432:
		return "StatusBeyondVdl"
	case 0xC0000433:
		return "StatusEncounteredWriteInProgress"
	case 0xC0000434:
		return "StatusPteChanged"
	case 0xC0000435:
		return "StatusPurgeFailed"
	case 0xC0000440:
		return "StatusCredRequiresConfirmation"
	case 0xC0000441:
		return "StatusCsEncryptionInvalidServerResponse"
	case 0xC0000442:
		return "StatusCsEncryptionUnsupportedServer"
	case 0xC0000443:
		return "StatusCsEncryptionExistingEncryptedFile"
	case 0xC0000444:
		return "StatusCsEncryptionNewEncryptedFile"
	case 0xC0000445:
		return "StatusCsEncryptionFileNotCse"
	case 0xC0000446:
		return "StatusInvalidLabel"
	case 0xC0000450:
		return "StatusDriverProcessTerminated"
	case 0xC0000451:
		return "StatusAmbiguousSystemDevice"
	case 0xC0000452:
		return "StatusSystemDeviceNotFound"
	case 0xC0000453:
		return "StatusRestartBootApplication"
	case 0xC0000454:
		return "StatusInsufficientNvramResources"
	case 0xC0000500:
		return "StatusInvalidTaskName"
	case 0xC0000501:
		return "StatusInvalidTaskIndex"
	case 0xC0000502:
		return "StatusThreadAlreadyInTask"
	case 0xC0000503:
		return "StatusCallbackBypass"
	case 0xC0000602:
		return "StatusFailFastException"
	case 0xC0000603:
		return "StatusImageCertRevoked"
	case 0xC0000700:
		return "StatusPortClosed"
	case 0xC0000701:
		return "StatusMessageLost"
	case 0xC0000702:
		return "StatusInvalidMessage"
	case 0xC0000703:
		return "StatusRequestCanceled"
	case 0xC0000704:
		return "StatusRecursiveDispatch"
	case 0xC0000705:
		return "StatusLpcReceiveBufferExpected"
	case 0xC0000706:
		return "StatusLpcInvalidConnectionUsage"
	case 0xC0000707:
		return "StatusLpcRequestsNotAllowed"
	case 0xC0000708:
		return "StatusResourceInUse"
	case 0xC0000709:
		return "StatusHardwareMemoryError"
	case 0xC000070A:
		return "StatusThreadpoolHandleException"
	case 0xC000070B:
		return "StatusThreadpoolSetEventOnCompletionFailed"
	case 0xC000070C:
		return "StatusThreadpoolReleaseSemaphoreOnCompletionFailed"
	case 0xC000070D:
		return "StatusThreadpoolReleaseMutexOnCompletionFailed"
	case 0xC000070E:
		return "StatusThreadpoolFreeLibraryOnCompletionFailed"
	case 0xC000070F:
		return "StatusThreadpoolReleasedDuringOperation"
	case 0xC0000710:
		return "StatusCallbackReturnedWhileImpersonating"
	case 0xC0000711:
		return "StatusApcReturnedWhileImpersonating"
	case 0xC0000712:
		return "StatusProcessIsProtected"
	case 0xC0000713:
		return "StatusMcaException"
	case 0xC0000714:
		return "StatusCertificateMappingNotUnique"
	case 0xC0000715:
		return "StatusSymlinkClassDisabled"
	case 0xC0000716:
		return "StatusInvalidIdnNormalization"
	case 0xC0000717:
		return "StatusNoUnicodeTranslation"
	case 0xC0000718:
		return "StatusAlreadyRegistered"
	case 0xC0000719:
		return "StatusContextMismatch"
	case 0xC000071A:
		return "StatusPortAlreadyHasCompletionList"
	case 0xC000071B:
		return "StatusCallbackReturnedThreadPriority"
	case 0xC000071C:
		return "StatusInvalidThread"
	case 0xC000071D:
		return "StatusCallbackReturnedTransaction"
	case 0xC000071E:
		return "StatusCallbackReturnedLdrLock"
	case 0xC000071F:
		return "StatusCallbackReturnedLang"
	case 0xC0000720:
		return "StatusCallbackReturnedPriBack"
	case 0xC0000721:
		return "StatusCallbackReturnedThreadAffinity"
	case 0xC0000800:
		return "StatusDiskRepairDisabled"
	case 0xC0000801:
		return "StatusDsDomainRenameInProgress"
	case 0xC0000802:
		return "StatusDiskQuotaExceeded"
	case 0x80000803:
		return "StatusDataLostRepair"
	case 0xC0000804:
		return "StatusContentBlocked"
	case 0xC0000805:
		return "StatusBadClusters"
	case 0xC0000806:
		return "StatusVolumeDirty"
	case 0xC0000901:
		return "StatusFileCheckedOut"
	case 0xC0000902:
		return "StatusCheckoutRequired"
	case 0xC0000903:
		return "StatusBadFileType"
	case 0xC0000904:
		return "StatusFileTooLarge"
	case 0xC0000905:
		return "StatusFormsAuthRequired"
	case 0xC0000906:
		return "StatusVirusInfected"
	case 0xC0000907:
		return "StatusVirusDeleted"
	case 0xC0000908:
		return "StatusBadMcfgTable"
	case 0xC0000909:
		return "StatusCannotBreakOplock"
	case 0xC0009898:
		return "StatusWowAssertion"
	case 0xC000A000:
		return "StatusInvalidSignature"
	case 0xC000A001:
		return "StatusHmacNotSupported"
	case 0xC000A002:
		return "StatusAuthTagMismatch"
	case 0xC000A010:
		return "StatusIpsecQueueOverflow"
	case 0xC000A011:
		return "StatusNdQueueOverflow"
	case 0xC000A012:
		return "StatusHoplimitExceeded"
	case 0xC000A013:
		return "StatusProtocolNotSupported"
	case 0xC000A014:
		return "StatusFastpathRejected"
	case 0xC000A080:
		return "StatusLostWritebehindDataNetworkDisconnected"
	case 0xC000A081:
		return "StatusLostWritebehindDataNetworkServerError"
	case 0xC000A082:
		return "StatusLostWritebehindDataLocalDiskError"
	case 0xC000A083:
		return "StatusXmlParseError"
	case 0xC000A084:
		return "StatusXmldsigError"
	case 0xC000A085:
		return "StatusWrongCompartment"
	case 0xC000A086:
		return "StatusAuthipFailure"
	case 0xC000A087:
		return "StatusDsOidMappedGroupCantHaveMembers"
	case 0xC000A088:
		return "StatusDsOidNotFound"
	case 0xC000A100:
		return "StatusHashNotSupported"
	case 0xC000A101:
		return "StatusHashNotPresent"
	case 0xC0010001:
		return "DbgNoStateChange"
	case 0xC0010002:
		return "DbgAppNotIdle"
	case 0xC0020001:
		return "RpcNtInvalidStringBinding"
	case 0xC0020002:
		return "RpcNtWrongKindOfBinding"
	case 0xC0020003:
		return "RpcNtInvalidBinding"
	case 0xC0020004:
		return "RpcNtProtseqNotSupported"
	case 0xC0020005:
		return "RpcNtInvalidRpcProtseq"
	case 0xC0020006:
		return "RpcNtInvalidStringUuid"
	case 0xC0020007:
		return "RpcNtInvalidEndpointFormat"
	case 0xC0020008:
		return "RpcNtInvalidNetAddr"
	case 0xC0020009:
		return "RpcNtNoEndpointFound"
	case 0xC002000A:
		return "RpcNtInvalidTimeout"
	case 0xC002000B:
		return "RpcNtObjectNotFound"
	case 0xC002000C:
		return "RpcNtAlreadyRegistered"
	case 0xC002000D:
		return "RpcNtTypeAlreadyRegistered"
	case 0xC002000E:
		return "RpcNtAlreadyListening"
	case 0xC002000F:
		return "RpcNtNoProtseqsRegistered"
	case 0xC0020010:
		return "RpcNtNotListening"
	case 0xC0020011:
		return "RpcNtUnknownMgrType"
	case 0xC0020012:
		return "RpcNtUnknownIf"
	case 0xC0020013:
		return "RpcNtNoBindings"
	case 0xC0020014:
		return "RpcNtNoProtseqs"
	case 0xC0020015:
		return "RpcNtCantCreateEndpoint"
	case 0xC0020016:
		return "RpcNtOutOfResources"
	case 0xC0020017:
		return "RpcNtServerUnavailable"
	case 0xC0020018:
		return "RpcNtServerTooBusy"
	case 0xC0020019:
		return "RpcNtInvalidNetworkOptions"
	case 0xC002001A:
		return "RpcNtNoCallActive"
	case 0xC002001B:
		return "RpcNtCallFailed"
	case 0xC002001C:
		return "RpcNtCallFailedDne"
	case 0xC002001D:
		return "RpcNtProtocolError"
	case 0xC002001F:
		return "RpcNtUnsupportedTransSyn"
	case 0xC0020021:
		return "RpcNtUnsupportedType"
	case 0xC0020022:
		return "RpcNtInvalidTag"
	case 0xC0020023:
		return "RpcNtInvalidBound"
	case 0xC0020024:
		return "RpcNtNoEntryName"
	case 0xC0020025:
		return "RpcNtInvalidNameSyntax"
	case 0xC0020026:
		return "RpcNtUnsupportedNameSyntax"
	case 0xC0020028:
		return "RpcNtUuidNoAddress"
	case 0xC0020029:
		return "RpcNtDuplicateEndpoint"
	case 0xC002002A:
		return "RpcNtUnknownAuthnType"
	case 0xC002002B:
		return "RpcNtMaxCallsTooSmall"
	case 0xC002002C:
		return "RpcNtStringTooLong"
	case 0xC002002D:
		return "RpcNtProtseqNotFound"
	case 0xC002002E:
		return "RpcNtProcnumOutOfRange"
	case 0xC002002F:
		return "RpcNtBindingHasNoAuth"
	case 0xC0020030:
		return "RpcNtUnknownAuthnService"
	case 0xC0020031:
		return "RpcNtUnknownAuthnLevel"
	case 0xC0020032:
		return "RpcNtInvalidAuthIdentity"
	case 0xC0020033:
		return "RpcNtUnknownAuthzService"
	case 0xC0020034:
		return "EptNtInvalidEntry"
	case 0xC0020035:
		return "EptNtCantPerformOp"
	case 0xC0020036:
		return "EptNtNotRegistered"
	case 0xC0020037:
		return "RpcNtNothingToExport"
	case 0xC0020038:
		return "RpcNtIncompleteName"
	case 0xC0020039:
		return "RpcNtInvalidVersOption"
	case 0xC002003A:
		return "RpcNtNoMoreMembers"
	case 0xC002003B:
		return "RpcNtNotAllObjsUnexported"
	case 0xC002003C:
		return "RpcNtInterfaceNotFound"
	case 0xC002003D:
		return "RpcNtEntryAlreadyExists"
	case 0xC002003E:
		return "RpcNtEntryNotFound"
	case 0xC002003F:
		return "RpcNtNameServiceUnavailable"
	case 0xC0020040:
		return "RpcNtInvalidNafId"
	case 0xC0020041:
		return "RpcNtCannotSupport"
	case 0xC0020042:
		return "RpcNtNoContextAvailable"
	case 0xC0020043:
		return "RpcNtInternalError"
	case 0xC0020044:
		return "RpcNtZeroDivide"
	case 0xC0020045:
		return "RpcNtAddressError"
	case 0xC0020046:
		return "RpcNtFpDivZero"
	case 0xC0020047:
		return "RpcNtFpUnderflow"
	case 0xC0020048:
		return "RpcNtFpOverflow"
	case 0xC0030001:
		return "RpcNtNoMoreEntries"
	case 0xC0030002:
		return "RpcNtSsCharTransOpenFail"
	case 0xC0030003:
		return "RpcNtSsCharTransShortFile"
	case 0xC0030004:
		return "RpcNtSsInNullContext"
	case 0xC0030005:
		return "RpcNtSsContextMismatch"
	case 0xC0030006:
		return "RpcNtSsContextDamaged"
	case 0xC0030007:
		return "RpcNtSsHandlesMismatch"
	case 0xC0030008:
		return "RpcNtSsCannotGetCallHandle"
	case 0xC0030009:
		return "RpcNtNullRefPointer"
	case 0xC003000A:
		return "RpcNtEnumValueOutOfRange"
	case 0xC003000B:
		return "RpcNtByteCountTooSmall"
	case 0xC003000C:
		return "RpcNtBadStubData"
	case 0xC0020049:
		return "RpcNtCallInProgress"
	case 0xC002004A:
		return "RpcNtNoMoreBindings"
	case 0xC002004B:
		return "RpcNtGroupMemberNotFound"
	case 0xC002004C:
		return "EptNtCantCreate"
	case 0xC002004D:
		return "RpcNtInvalidObject"
	case 0xC002004F:
		return "RpcNtNoInterfaces"
	case 0xC0020050:
		return "RpcNtCallCancelled"
	case 0xC0020051:
		return "RpcNtBindingIncomplete"
	case 0xC0020052:
		return "RpcNtCommFailure"
	case 0xC0020053:
		return "RpcNtUnsupportedAuthnLevel"
	case 0xC0020054:
		return "RpcNtNoPrincName"
	case 0xC0020055:
		return "RpcNtNotRpcError"
	case 0x40020056:
		return "RpcNtUuidLocalOnly"
	case 0xC0020057:
		return "RpcNtSecPkgError"
	case 0xC0020058:
		return "RpcNtNotCancelled"
	case 0xC0030059:
		return "RpcNtInvalidEsAction"
	case 0xC003005A:
		return "RpcNtWrongEsVersion"
	case 0xC003005B:
		return "RpcNtWrongStubVersion"
	case 0xC003005C:
		return "RpcNtInvalidPipeObject"
	case 0xC003005D:
		return "RpcNtInvalidPipeOperation"
	case 0xC003005E:
		return "RpcNtWrongPipeVersion"
	case 0xC003005F:
		return "RpcNtPipeClosed"
	case 0xC0030060:
		return "RpcNtPipeDisciplineError"
	case 0xC0030061:
		return "RpcNtPipeEmpty"
	case 0xC0020062:
		return "RpcNtInvalidAsyncHandle"
	case 0xC0020063:
		return "RpcNtInvalidAsyncCall"
	case 0xC0020064:
		return "RpcNtProxyAccessDenied"
	case 0xC0020065:
		return "RpcNtCookieAuthFailed"
	case 0x400200AF:
		return "RpcNtSendIncomplete"
	case 0xC0140001:
		return "StatusAcpiInvalidOpcode"
	case 0xC0140002:
		return "StatusAcpiStackOverflow"
	case 0xC0140003:
		return "StatusAcpiAssertFailed"
	case 0xC0140004:
		return "StatusAcpiInvalidIndex"
	case 0xC0140005:
		return "StatusAcpiInvalidArgument"
	case 0xC0140006:
		return "StatusAcpiFatal"
	case 0xC0140007:
		return "StatusAcpiInvalidSupername"
	case 0xC0140008:
		return "StatusAcpiInvalidArgtype"
	case 0xC0140009:
		return "StatusAcpiInvalidObjtype"
	case 0xC014000A:
		return "StatusAcpiInvalidTargettype"
	case 0xC014000B:
		return "StatusAcpiIncorrectArgumentCount"
	case 0xC014000C:
		return "StatusAcpiAddressNotMapped"
	case 0xC014000D:
		return "StatusAcpiInvalidEventtype"
	case 0xC014000E:
		return "StatusAcpiHandlerCollision"
	case 0xC014000F:
		return "StatusAcpiInvalidData"
	case 0xC0140010:
		return "StatusAcpiInvalidRegion"
	case 0xC0140011:
		return "StatusAcpiInvalidAccessSize"
	case 0xC0140012:
		return "StatusAcpiAcquireGlobalLock"
	case 0xC0140013:
		return "StatusAcpiAlreadyInitialized"
	case 0xC0140014:
		return "StatusAcpiNotInitialized"
	case 0xC0140015:
		return "StatusAcpiInvalidMutexLevel"
	case 0xC0140016:
		return "StatusAcpiMutexNotOwned"
	case 0xC0140017:
		return "StatusAcpiMutexNotOwner"
	case 0xC0140018:
		return "StatusAcpiRsAccess"
	case 0xC0140019:
		return "StatusAcpiInvalidTable"
	case 0xC0140020:
		return "StatusAcpiRegHandlerFailed"
	case 0xC0140021:
		return "StatusAcpiPowerRequestFailed"
	case 0xC00A0001:
		return "StatusCtxWinstationNameInvalid"
	case 0xC00A0002:
		return "StatusCtxInvalidPd"
	case 0xC00A0003:
		return "StatusCtxPdNotFound"
	case 0x400A0004:
		return "StatusCtxCdmConnect"
	case 0x400A0005:
		return "StatusCtxCdmDisconnect"
	case 0xC00A0006:
		return "StatusCtxClosePending"
	case 0xC00A0007:
		return "StatusCtxNoOutbuf"
	case 0xC00A0008:
		return "StatusCtxModemInfNotFound"
	case 0xC00A0009:
		return "StatusCtxInvalidModemname"
	case 0xC00A000A:
		return "StatusCtxResponseError"
	case 0xC00A000B:
		return "StatusCtxModemResponseTimeout"
	case 0xC00A000C:
		return "StatusCtxModemResponseNoCarrier"
	case 0xC00A000D:
		return "StatusCtxModemResponseNoDialtone"
	case 0xC00A000E:
		return "StatusCtxModemResponseBusy"
	case 0xC00A000F:
		return "StatusCtxModemResponseVoice"
	case 0xC00A0010:
		return "StatusCtxTdError"
	case 0xC00A0012:
		return "StatusCtxLicenseClientInvalid"
	case 0xC00A0013:
		return "StatusCtxLicenseNotAvailable"
	case 0xC00A0014:
		return "StatusCtxLicenseExpired"
	case 0xC00A0015:
		return "StatusCtxWinstationNotFound"
	case 0xC00A0016:
		return "StatusCtxWinstationNameCollision"
	case 0xC00A0017:
		return "StatusCtxWinstationBusy"
	case 0xC00A0018:
		return "StatusCtxBadVideoMode"
	case 0xC00A0022:
		return "StatusCtxGraphicsInvalid"
	case 0xC00A0024:
		return "StatusCtxNotConsole"
	case 0xC00A0026:
		return "StatusCtxClientQueryTimeout"
	case 0xC00A0027:
		return "StatusCtxConsoleDisconnect"
	case 0xC00A0028:
		return "StatusCtxConsoleConnect"
	case 0xC00A002A:
		return "StatusCtxShadowDenied"
	case 0xC00A002B:
		return "StatusCtxWinstationAccessDenied"
	case 0xC00A002E:
		return "StatusCtxInvalidWd"
	case 0xC00A002F:
		return "StatusCtxWdNotFound"
	case 0xC00A0030:
		return "StatusCtxShadowInvalid"
	case 0xC00A0031:
		return "StatusCtxShadowDisabled"
	case 0xC00A0032:
		return "StatusRdpProtocolError"
	case 0xC00A0033:
		return "StatusCtxClientLicenseNotSet"
	case 0xC00A0034:
		return "StatusCtxClientLicenseInUse"
	case 0xC00A0035:
		return "StatusCtxShadowEndedByModeChange"
	case 0xC00A0036:
		return "StatusCtxShadowNotRunning"
	case 0xC00A0037:
		return "StatusCtxLogonDisabled"
	case 0xC00A0038:
		return "StatusCtxSecurityLayerError"
	case 0xC00A0039:
		return "StatusTsIncompatibleSessions"
	case 0xC00A003A:
		return "StatusTsVideoSubsystemError"
	case 0xC0040035:
		return "StatusPnpBadMpsTable"
	case 0xC0040036:
		return "StatusPnpTranslationFailed"
	case 0xC0040037:
		return "StatusPnpIrqTranslationFailed"
	case 0xC0040038:
		return "StatusPnpInvalidId"
	case 0xC0040039:
		return "StatusIoReissueAsCached"
	case 0xC00B0001:
		return "StatusMuiFileNotFound"
	case 0xC00B0002:
		return "StatusMuiInvalidFile"
	case 0xC00B0003:
		return "StatusMuiInvalidRcConfig"
	case 0xC00B0004:
		return "StatusMuiInvalidLocaleName"
	case 0xC00B0005:
		return "StatusMuiInvalidUltimatefallbackName"
	case 0xC00B0006:
		return "StatusMuiFileNotLoaded"
	case 0xC00B0007:
		return "StatusResourceEnumUserStop"
	case 0xC01C0001:
		return "StatusFltNoHandlerDefined"
	case 0xC01C0002:
		return "StatusFltContextAlreadyDefined"
	case 0xC01C0003:
		return "StatusFltInvalidAsynchronousRequest"
	case 0xC01C0004:
		return "StatusFltDisallowFastIo"
	case 0xC01C0005:
		return "StatusFltInvalidNameRequest"
	case 0xC01C0006:
		return "StatusFltNotSafeToPostOperation"
	case 0xC01C0007:
		return "StatusFltNotInitialized"
	case 0xC01C0008:
		return "StatusFltFilterNotReady"
	case 0xC01C0009:
		return "StatusFltPostOperationCleanup"
	case 0xC01C000A:
		return "StatusFltInternalError"
	case 0xC01C000B:
		return "StatusFltDeletingObject"
	case 0xC01C000C:
		return "StatusFltMustBeNonpagedPool"
	case 0xC01C000D:
		return "StatusFltDuplicateEntry"
	case 0xC01C000E:
		return "StatusFltCbdqDisabled"
	case 0xC01C000F:
		return "StatusFltDoNotAttach"
	case 0xC01C0010:
		return "StatusFltDoNotDetach"
	case 0xC01C0011:
		return "StatusFltInstanceAltitudeCollision"
	case 0xC01C0012:
		return "StatusFltInstanceNameCollision"
	case 0xC01C0013:
		return "StatusFltFilterNotFound"
	case 0xC01C0014:
		return "StatusFltVolumeNotFound"
	case 0xC01C0015:
		return "StatusFltInstanceNotFound"
	case 0xC01C0016:
		return "StatusFltContextAllocationNotFound"
	case 0xC01C0017:
		return "StatusFltInvalidContextRegistration"
	case 0xC01C0018:
		return "StatusFltNameCacheMiss"
	case 0xC01C0019:
		return "StatusFltNoDeviceObject"
	case 0xC01C001A:
		return "StatusFltVolumeAlreadyMounted"
	case 0xC01C001B:
		return "StatusFltAlreadyEnlisted"
	case 0xC01C001C:
		return "StatusFltContextAlreadyLinked"
	case 0xC01C0020:
		return "StatusFltNoWaiterForReply"
	case 0xC0150001:
		return "StatusSxsSectionNotFound"
	case 0xC0150002:
		return "StatusSxsCantGenActctx"
	case 0xC0150003:
		return "StatusSxsInvalidActctxdataFormat"
	case 0xC0150004:
		return "StatusSxsAssemblyNotFound"
	case 0xC0150005:
		return "StatusSxsManifestFormatError"
	case 0xC0150006:
		return "StatusSxsManifestParseError"
	case 0xC0150007:
		return "StatusSxsActivationContextDisabled"
	case 0xC0150008:
		return "StatusSxsKeyNotFound"
	case 0xC0150009:
		return "StatusSxsVersionConflict"
	case 0xC015000A:
		return "StatusSxsWrongSectionType"
	case 0xC015000B:
		return "StatusSxsThreadQueriesDisabled"
	case 0xC015000C:
		return "StatusSxsAssemblyMissing"
	case 0x4015000D:
		return "StatusSxsReleaseActivationContext"
	case 0xC015000E:
		return "StatusSxsProcessDefaultAlreadySet"
	case 0xC015000F:
		return "StatusSxsEarlyDeactivation"
	case 0xC0150010:
		return "StatusSxsInvalidDeactivation"
	case 0xC0150011:
		return "StatusSxsMultipleDeactivation"
	case 0xC0150012:
		return "StatusSxsSystemDefaultActivationContextEmpty"
	case 0xC0150013:
		return "StatusSxsProcessTerminationRequested"
	case 0xC0150014:
		return "StatusSxsCorruptActivationStack"
	case 0xC0150015:
		return "StatusSxsCorruption"
	case 0xC0150016:
		return "StatusSxsInvalidIdentityAttributeValue"
	case 0xC0150017:
		return "StatusSxsInvalidIdentityAttributeName"
	case 0xC0150018:
		return "StatusSxsIdentityDuplicateAttribute"
	case 0xC0150019:
		return "StatusSxsIdentityParseError"
	case 0xC015001A:
		return "StatusSxsComponentStoreCorrupt"
	case 0xC015001B:
		return "StatusSxsFileHashMismatch"
	case 0xC015001C:
		return "StatusSxsManifestIdentitySameButContentsDifferent"
	case 0xC015001D:
		return "StatusSxsIdentitiesDifferent"
	case 0xC015001E:
		return "StatusSxsAssemblyIsNotADeployment"
	case 0xC015001F:
		return "StatusSxsFileNotPartOfAssembly"
	case 0xC0150020:
		return "StatusAdvancedInstallerFailed"
	case 0xC0150021:
		return "StatusXmlEncodingMismatch"
	case 0xC0150022:
		return "StatusSxsManifestTooBig"
	case 0xC0150023:
		return "StatusSxsSettingNotRegistered"
	case 0xC0150024:
		return "StatusSxsTransactionClosureIncomplete"
	case 0xC0150025:
		return "StatusSmiPrimitiveInstallerFailed"
	case 0xC0150026:
		return "StatusGenericCommandFailed"
	case 0xC0150027:
		return "StatusSxsFileHashMissing"
	case 0xC0130001:
		return "StatusClusterInvalidNode"
	case 0xC0130002:
		return "StatusClusterNodeExists"
	case 0xC0130003:
		return "StatusClusterJoinInProgress"
	case 0xC0130004:
		return "StatusClusterNodeNotFound"
	case 0xC0130005:
		return "StatusClusterLocalNodeNotFound"
	case 0xC0130006:
		return "StatusClusterNetworkExists"
	case 0xC0130007:
		return "StatusClusterNetworkNotFound"
	case 0xC0130008:
		return "StatusClusterNetinterfaceExists"
	case 0xC0130009:
		return "StatusClusterNetinterfaceNotFound"
	case 0xC013000A:
		return "StatusClusterInvalidRequest"
	case 0xC013000B:
		return "StatusClusterInvalidNetworkProvider"
	case 0xC013000C:
		return "StatusClusterNodeDown"
	case 0xC013000D:
		return "StatusClusterNodeUnreachable"
	case 0xC013000E:
		return "StatusClusterNodeNotMember"
	case 0xC013000F:
		return "StatusClusterJoinNotInProgress"
	case 0xC0130010:
		return "StatusClusterInvalidNetwork"
	case 0xC0130011:
		return "StatusClusterNoNetAdapters"
	case 0xC0130012:
		return "StatusClusterNodeUp"
	case 0xC0130013:
		return "StatusClusterNodePaused"
	case 0xC0130014:
		return "StatusClusterNodeNotPaused"
	case 0xC0130015:
		return "StatusClusterNoSecurityContext"
	case 0xC0130016:
		return "StatusClusterNetworkNotInternal"
	case 0xC0130017:
		return "StatusClusterPoisoned"
	case 0xC0130018:
		return "StatusClusterNonCsvPath"
	case 0xC0130019:
		return "StatusClusterCsvVolumeNotLocal"
	case 0xC0190001:
		return "StatusTransactionalConflict"
	case 0xC0190002:
		return "StatusInvalidTransaction"
	case 0xC0190003:
		return "StatusTransactionNotActive"
	case 0xC0190004:
		return "StatusTmInitializationFailed"
	case 0xC0190005:
		return "StatusRmNotActive"
	case 0xC0190006:
		return "StatusRmMetadataCorrupt"
	case 0xC0190007:
		return "StatusTransactionNotJoined"
	case 0xC0190008:
		return "StatusDirectoryNotRm"
	case 0x80190009:
		return "StatusCouldNotResizeLog"
	case 0xC019000A:
		return "StatusTransactionsUnsupportedRemote"
	case 0xC019000B:
		return "StatusLogResizeInvalidSize"
	case 0xC019000C:
		return "StatusRemoteFileVersionMismatch"
	case 0xC019000F:
		return "StatusCrmProtocolAlreadyExists"
	case 0xC0190010:
		return "StatusTransactionPropagationFailed"
	case 0xC0190011:
		return "StatusCrmProtocolNotFound"
	case 0xC0190012:
		return "StatusTransactionSuperiorExists"
	case 0xC0190013:
		return "StatusTransactionRequestNotValid"
	case 0xC0190014:
		return "StatusTransactionNotRequested"
	case 0xC0190015:
		return "StatusTransactionAlreadyAborted"
	case 0xC0190016:
		return "StatusTransactionAlreadyCommitted"
	case 0xC0190017:
		return "StatusTransactionInvalidMarshallBuffer"
	case 0xC0190018:
		return "StatusCurrentTransactionNotValid"
	case 0xC0190019:
		return "StatusLogGrowthFailed"
	case 0xC0190021:
		return "StatusObjectNoLongerExists"
	case 0xC0190022:
		return "StatusStreamMiniversionNotFound"
	case 0xC0190023:
		return "StatusStreamMiniversionNotValid"
	case 0xC0190024:
		return "StatusMiniversionInaccessibleFromSpecifiedTransaction"
	case 0xC0190025:
		return "StatusCantOpenMiniversionWithModifyIntent"
	case 0xC0190026:
		return "StatusCantCreateMoreStreamMiniversions"
	case 0xC0190028:
		return "StatusHandleNoLongerValid"
	case 0x80190029:
		return "StatusNoTxfMetadata"
	case 0xC0190030:
		return "StatusLogCorruptionDetected"
	case 0x80190031:
		return "StatusCantRecoverWithHandleOpen"
	case 0xC0190032:
		return "StatusRmDisconnected"
	case 0xC0190033:
		return "StatusEnlistmentNotSuperior"
	case 0x40190034:
		return "StatusRecoveryNotNeeded"
	case 0x40190035:
		return "StatusRmAlreadyStarted"
	case 0xC0190036:
		return "StatusFileIdentityNotPersistent"
	case 0xC0190037:
		return "StatusCantBreakTransactionalDependency"
	case 0xC0190038:
		return "StatusCantCrossRmBoundary"
	case 0xC0190039:
		return "StatusTxfDirNotEmpty"
	case 0xC019003A:
		return "StatusIndoubtTransactionsExist"
	case 0xC019003B:
		return "StatusTmVolatile"
	case 0xC019003C:
		return "StatusRollbackTimerExpired"
	case 0xC019003D:
		return "StatusTxfAttributeCorrupt"
	case 0xC019003E:
		return "StatusEfsNotAllowedInTransaction"
	case 0xC019003F:
		return "StatusTransactionalOpenNotAllowed"
	case 0xC0190040:
		return "StatusTransactedMappingUnsupportedRemote"
	case 0x80190041:
		return "StatusTxfMetadataAlreadyPresent"
	case 0x80190042:
		return "StatusTransactionScopeCallbacksNotSet"
	case 0xC0190043:
		return "StatusTransactionRequiredPromotion"
	case 0xC0190044:
		return "StatusCannotExecuteFileInTransaction"
	case 0xC0190045:
		return "StatusTransactionsNotFrozen"
	case 0xC0190046:
		return "StatusTransactionFreezeInProgress"
	case 0xC0190047:
		return "StatusNotSnapshotVolume"
	case 0xC0190048:
		return "StatusNoSavepointWithOpenFiles"
	case 0xC0190049:
		return "StatusSparseNotAllowedInTransaction"
	case 0xC019004A:
		return "StatusTmIdentityMismatch"
	case 0xC019004B:
		return "StatusFloatedSection"
	case 0xC019004C:
		return "StatusCannotAcceptTransactedWork"
	case 0xC019004D:
		return "StatusCannotAbortTransactions"
	case 0xC019004E:
		return "StatusTransactionNotFound"
	case 0xC019004F:
		return "StatusResourcemanagerNotFound"
	case 0xC0190050:
		return "StatusEnlistmentNotFound"
	case 0xC0190051:
		return "StatusTransactionmanagerNotFound"
	case 0xC0190052:
		return "StatusTransactionmanagerNotOnline"
	case 0xC0190053:
		return "StatusTransactionmanagerRecoveryNameCollision"
	case 0xC0190054:
		return "StatusTransactionNotRoot"
	case 0xC0190055:
		return "StatusTransactionObjectExpired"
	case 0xC0190056:
		return "StatusCompressionNotAllowedInTransaction"
	case 0xC0190057:
		return "StatusTransactionResponseNotEnlisted"
	case 0xC0190058:
		return "StatusTransactionRecordTooLong"
	case 0xC0190059:
		return "StatusNoLinkTrackingInTransaction"
	case 0xC019005A:
		return "StatusOperationNotSupportedInTransaction"
	case 0xC019005B:
		return "StatusTransactionIntegrityViolated"
	case 0xC019005C:
		return "StatusTransactionmanagerIdentityMismatch"
	case 0xC019005D:
		return "StatusRmCannotBeFrozenForSnapshot"
	case 0xC019005E:
		return "StatusTransactionMustWritethrough"
	case 0xC019005F:
		return "StatusTransactionNoSuperior"
	case 0xC0190060:
		return "StatusExpiredHandle"
	case 0xC0190061:
		return "StatusTransactionNotEnlisted"
	case 0xC01A0001:
		return "StatusLogSectorInvalid"
	case 0xC01A0002:
		return "StatusLogSectorParityInvalid"
	case 0xC01A0003:
		return "StatusLogSectorRemapped"
	case 0xC01A0004:
		return "StatusLogBlockIncomplete"
	case 0xC01A0005:
		return "StatusLogInvalidRange"
	case 0xC01A0006:
		return "StatusLogBlocksExhausted"
	case 0xC01A0007:
		return "StatusLogReadContextInvalid"
	case 0xC01A0008:
		return "StatusLogRestartInvalid"
	case 0xC01A0009:
		return "StatusLogBlockVersion"
	case 0xC01A000A:
		return "StatusLogBlockInvalid"
	case 0xC01A000B:
		return "StatusLogReadModeInvalid"
	case 0x401A000C:
		return "StatusLogNoRestart"
	case 0xC01A000D:
		return "StatusLogMetadataCorrupt"
	case 0xC01A000E:
		return "StatusLogMetadataInvalid"
	case 0xC01A000F:
		return "StatusLogMetadataInconsistent"
	case 0xC01A0010:
		return "StatusLogReservationInvalid"
	case 0xC01A0011:
		return "StatusLogCantDelete"
	case 0xC01A0012:
		return "StatusLogContainerLimitExceeded"
	case 0xC01A0013:
		return "StatusLogStartOfLog"
	case 0xC01A0014:
		return "StatusLogPolicyAlreadyInstalled"
	case 0xC01A0015:
		return "StatusLogPolicyNotInstalled"
	case 0xC01A0016:
		return "StatusLogPolicyInvalid"
	case 0xC01A0017:
		return "StatusLogPolicyConflict"
	case 0xC01A0018:
		return "StatusLogPinnedArchiveTail"
	case 0xC01A0019:
		return "StatusLogRecordNonexistent"
	case 0xC01A001A:
		return "StatusLogRecordsReservedInvalid"
	case 0xC01A001B:
		return "StatusLogSpaceReservedInvalid"
	case 0xC01A001C:
		return "StatusLogTailInvalid"
	case 0xC01A001D:
		return "StatusLogFull"
	case 0xC01A001E:
		return "StatusLogMultiplexed"
	case 0xC01A001F:
		return "StatusLogDedicated"
	case 0xC01A0020:
		return "StatusLogArchiveNotInProgress"
	case 0xC01A0021:
		return "StatusLogArchiveInProgress"
	case 0xC01A0022:
		return "StatusLogEphemeral"
	case 0xC01A0023:
		return "StatusLogNotEnoughContainers"
	case 0xC01A0024:
		return "StatusLogClientAlreadyRegistered"
	case 0xC01A0025:
		return "StatusLogClientNotRegistered"
	case 0xC01A0026:
		return "StatusLogFullHandlerInProgress"
	case 0xC01A0027:
		return "StatusLogContainerReadFailed"
	case 0xC01A0028:
		return "StatusLogContainerWriteFailed"
	case 0xC01A0029:
		return "StatusLogContainerOpenFailed"
	case 0xC01A002A:
		return "StatusLogContainerStateInvalid"
	case 0xC01A002B:
		return "StatusLogStateInvalid"
	case 0xC01A002C:
		return "StatusLogPinned"
	case 0xC01A002D:
		return "StatusLogMetadataFlushFailed"
	case 0xC01A002E:
		return "StatusLogInconsistentSecurity"
	case 0xC01A002F:
		return "StatusLogAppendedFlushFailed"
	case 0xC01A0030:
		return "StatusLogPinnedReservation"
	case 0xC01B00EA:
		return "StatusVideoHungDisplayDriverThread"
	case 0x801B00EB:
		return "StatusVideoHungDisplayDriverThreadRecovered"
	case 0x401B00EC:
		return "StatusVideoDriverDebugReportRequest"
	case 0xC01D0001:
		return "StatusMonitorNoDescriptor"
	case 0xC01D0002:
		return "StatusMonitorUnknownDescriptorFormat"
	case 0xC01D0003:
		return "StatusMonitorInvalidDescriptorChecksum"
	case 0xC01D0004:
		return "StatusMonitorInvalidStandardTimingBlock"
	case 0xC01D0005:
		return "StatusMonitorWmiDatablockRegistrationFailed"
	case 0xC01D0006:
		return "StatusMonitorInvalidSerialNumberMondscBlock"
	case 0xC01D0007:
		return "StatusMonitorInvalidUserFriendlyMondscBlock"
	case 0xC01D0008:
		return "StatusMonitorNoMoreDescriptorData"
	case 0xC01D0009:
		return "StatusMonitorInvalidDetailedTimingBlock"
	case 0xC01D000A:
		return "StatusMonitorInvalidManufactureDate"
	case 0xC01E0000:
		return "StatusGraphicsNotExclusiveModeOwner"
	case 0xC01E0001:
		return "StatusGraphicsInsufficientDmaBuffer"
	case 0xC01E0002:
		return "StatusGraphicsInvalidDisplayAdapter"
	case 0xC01E0003:
		return "StatusGraphicsAdapterWasReset"
	case 0xC01E0004:
		return "StatusGraphicsInvalidDriverModel"
	case 0xC01E0005:
		return "StatusGraphicsPresentModeChanged"
	case 0xC01E0006:
		return "StatusGraphicsPresentOccluded"
	case 0xC01E0007:
		return "StatusGraphicsPresentDenied"
	case 0xC01E0008:
		return "StatusGraphicsCannotcolorconvert"
	case 0xC01E0009:
		return "StatusGraphicsDriverMismatch"
	case 0x401E000A:
		return "StatusGraphicsPartialDataPopulated"
	case 0xC01E000B:
		return "StatusGraphicsPresentRedirectionDisabled"
	case 0xC01E000C:
		return "StatusGraphicsPresentUnoccluded"
	case 0xC01E0100:
		return "StatusGraphicsNoVideoMemory"
	case 0xC01E0101:
		return "StatusGraphicsCantLockMemory"
	case 0xC01E0102:
		return "StatusGraphicsAllocationBusy"
	case 0xC01E0103:
		return "StatusGraphicsTooManyReferences"
	case 0xC01E0104:
		return "StatusGraphicsTryAgainLater"
	case 0xC01E0105:
		return "StatusGraphicsTryAgainNow"
	case 0xC01E0106:
		return "StatusGraphicsAllocationInvalid"
	case 0xC01E0107:
		return "StatusGraphicsUnswizzlingApertureUnavailable"
	case 0xC01E0108:
		return "StatusGraphicsUnswizzlingApertureUnsupported"
	case 0xC01E0109:
		return "StatusGraphicsCantEvictPinnedAllocation"
	case 0xC01E0110:
		return "StatusGraphicsInvalidAllocationUsage"
	case 0xC01E0111:
		return "StatusGraphicsCantRenderLockedAllocation"
	case 0xC01E0112:
		return "StatusGraphicsAllocationClosed"
	case 0xC01E0113:
		return "StatusGraphicsInvalidAllocationInstance"
	case 0xC01E0114:
		return "StatusGraphicsInvalidAllocationHandle"
	case 0xC01E0115:
		return "StatusGraphicsWrongAllocationDevice"
	case 0xC01E0116:
		return "StatusGraphicsAllocationContentLost"
	case 0xC01E0200:
		return "StatusGraphicsGpuExceptionOnDevice"
	case 0xC01E0300:
		return "StatusGraphicsInvalidVidpnTopology"
	case 0xC01E0301:
		return "StatusGraphicsVidpnTopologyNotSupported"
	case 0xC01E0302:
		return "StatusGraphicsVidpnTopologyCurrentlyNotSupported"
	case 0xC01E0303:
		return "StatusGraphicsInvalidVidpn"
	case 0xC01E0304:
		return "StatusGraphicsInvalidVideoPresentSource"
	case 0xC01E0305:
		return "StatusGraphicsInvalidVideoPresentTarget"
	case 0xC01E0306:
		return "StatusGraphicsVidpnModalityNotSupported"
	case 0x401E0307:
		return "StatusGraphicsModeNotPinned"
	case 0xC01E0308:
		return "StatusGraphicsInvalidVidpnSourcemodeset"
	case 0xC01E0309:
		return "StatusGraphicsInvalidVidpnTargetmodeset"
	case 0xC01E030A:
		return "StatusGraphicsInvalidFrequency"
	case 0xC01E030B:
		return "StatusGraphicsInvalidActiveRegion"
	case 0xC01E030C:
		return "StatusGraphicsInvalidTotalRegion"
	case 0xC01E0310:
		return "StatusGraphicsInvalidVideoPresentSourceMode"
	case 0xC01E0311:
		return "StatusGraphicsInvalidVideoPresentTargetMode"
	case 0xC01E0312:
		return "StatusGraphicsPinnedModeMustRemainInSet"
	case 0xC01E0313:
		return "StatusGraphicsPathAlreadyInTopology"
	case 0xC01E0314:
		return "StatusGraphicsModeAlreadyInModeset"
	case 0xC01E0315:
		return "StatusGraphicsInvalidVideopresentsourceset"
	case 0xC01E0316:
		return "StatusGraphicsInvalidVideopresenttargetset"
	case 0xC01E0317:
		return "StatusGraphicsSourceAlreadyInSet"
	case 0xC01E0318:
		return "StatusGraphicsTargetAlreadyInSet"
	case 0xC01E0319:
		return "StatusGraphicsInvalidVidpnPresentPath"
	case 0xC01E031A:
		return "StatusGraphicsNoRecommendedVidpnTopology"
	case 0xC01E031B:
		return "StatusGraphicsInvalidMonitorFrequencyrangeset"
	case 0xC01E031C:
		return "StatusGraphicsInvalidMonitorFrequencyrange"
	case 0xC01E031D:
		return "StatusGraphicsFrequencyrangeNotInSet"
	case 0x401E031E:
		return "StatusGraphicsNoPreferredMode"
	case 0xC01E031F:
		return "StatusGraphicsFrequencyrangeAlreadyInSet"
	case 0xC01E0320:
		return "StatusGraphicsStaleModeset"
	case 0xC01E0321:
		return "StatusGraphicsInvalidMonitorSourcemodeset"
	case 0xC01E0322:
		return "StatusGraphicsInvalidMonitorSourceMode"
	case 0xC01E0323:
		return "StatusGraphicsNoRecommendedFunctionalVidpn"
	case 0xC01E0324:
		return "StatusGraphicsModeIdMustBeUnique"
	case 0xC01E0325:
		return "StatusGraphicsEmptyAdapterMonitorModeSupportIntersection"
	case 0xC01E0326:
		return "StatusGraphicsVideoPresentTargetsLessThanSources"
	case 0xC01E0327:
		return "StatusGraphicsPathNotInTopology"
	case 0xC01E0328:
		return "StatusGraphicsAdapterMustHaveAtLeastOneSource"
	case 0xC01E0329:
		return "StatusGraphicsAdapterMustHaveAtLeastOneTarget"
	case 0xC01E032A:
		return "StatusGraphicsInvalidMonitordescriptorset"
	case 0xC01E032B:
		return "StatusGraphicsInvalidMonitordescriptor"
	case 0xC01E032C:
		return "StatusGraphicsMonitordescriptorNotInSet"
	case 0xC01E032D:
		return "StatusGraphicsMonitordescriptorAlreadyInSet"
	case 0xC01E032E:
		return "StatusGraphicsMonitordescriptorIdMustBeUnique"
	case 0xC01E032F:
		return "StatusGraphicsInvalidVidpnTargetSubsetType"
	case 0xC01E0330:
		return "StatusGraphicsResourcesNotRelated"
	case 0xC01E0331:
		return "StatusGraphicsSourceIdMustBeUnique"
	case 0xC01E0332:
		return "StatusGraphicsTargetIdMustBeUnique"
	case 0xC01E0333:
		return "StatusGraphicsNoAvailableVidpnTarget"
	case 0xC01E0334:
		return "StatusGraphicsMonitorCouldNotBeAssociatedWithAdapter"
	case 0xC01E0335:
		return "StatusGraphicsNoVidpnmgr"
	case 0xC01E0336:
		return "StatusGraphicsNoActiveVidpn"
	case 0xC01E0337:
		return "StatusGraphicsStaleVidpnTopology"
	case 0xC01E0338:
		return "StatusGraphicsMonitorNotConnected"
	case 0xC01E0339:
		return "StatusGraphicsSourceNotInTopology"
	case 0xC01E033A:
		return "StatusGraphicsInvalidPrimarysurfaceSize"
	case 0xC01E033B:
		return "StatusGraphicsInvalidVisibleregionSize"
	case 0xC01E033C:
		return "StatusGraphicsInvalidStride"
	case 0xC01E033D:
		return "StatusGraphicsInvalidPixelformat"
	case 0xC01E033E:
		return "StatusGraphicsInvalidColorbasis"
	case 0xC01E033F:
		return "StatusGraphicsInvalidPixelvalueaccessmode"
	case 0xC01E0340:
		return "StatusGraphicsTargetNotInTopology"
	case 0xC01E0341:
		return "StatusGraphicsNoDisplayModeManagementSupport"
	case 0xC01E0342:
		return "StatusGraphicsVidpnSourceInUse"
	case 0xC01E0343:
		return "StatusGraphicsCantAccessActiveVidpn"
	case 0xC01E0344:
		return "StatusGraphicsInvalidPathImportanceOrdinal"
	case 0xC01E0345:
		return "StatusGraphicsInvalidPathContentGeometryTransformation"
	case 0xC01E0346:
		return "StatusGraphicsPathContentGeometryTransformationNotSupported"
	case 0xC01E0347:
		return "StatusGraphicsInvalidGammaRamp"
	case 0xC01E0348:
		return "StatusGraphicsGammaRampNotSupported"
	case 0xC01E0349:
		return "StatusGraphicsMultisamplingNotSupported"
	case 0xC01E034A:
		return "StatusGraphicsModeNotInModeset"
	case 0x401E034B:
		return "StatusGraphicsDatasetIsEmpty"
	case 0x401E034C:
		return "StatusGraphicsNoMoreElementsInDataset"
	case 0xC01E034D:
		return "StatusGraphicsInvalidVidpnTopologyRecommendationReason"
	case 0xC01E034E:
		return "StatusGraphicsInvalidPathContentType"
	case 0xC01E034F:
		return "StatusGraphicsInvalidCopyprotectionType"
	case 0xC01E0350:
		return "StatusGraphicsUnassignedModesetAlreadyExists"
	case 0x401E0351:
		return "StatusGraphicsPathContentGeometryTransformationNotPinned"
	case 0xC01E0352:
		return "StatusGraphicsInvalidScanlineOrdering"
	case 0xC01E0353:
		return "StatusGraphicsTopologyChangesNotAllowed"
	case 0xC01E0354:
		return "StatusGraphicsNoAvailableImportanceOrdinals"
	case 0xC01E0355:
		return "StatusGraphicsIncompatiblePrivateFormat"
	case 0xC01E0356:
		return "StatusGraphicsInvalidModePruningAlgorithm"
	case 0xC01E0357:
		return "StatusGraphicsInvalidMonitorCapabilityOrigin"
	case 0xC01E0358:
		return "StatusGraphicsInvalidMonitorFrequencyrangeConstraint"
	case 0xC01E0359:
		return "StatusGraphicsMaxNumPathsReached"
	case 0xC01E035A:
		return "StatusGraphicsCancelVidpnTopologyAugmentation"
	case 0xC01E035B:
		return "StatusGraphicsInvalidClientType"
	case 0xC01E035C:
		return "StatusGraphicsClientvidpnNotSet"
	case 0xC01E0400:
		return "StatusGraphicsSpecifiedChildAlreadyConnected"
	case 0xC01E0401:
		return "StatusGraphicsChildDescriptorNotSupported"
	case 0x401E042F:
		return "StatusGraphicsUnknownChildStatus"
	case 0xC01E0430:
		return "StatusGraphicsNotALinkedAdapter"
	case 0xC01E0431:
		return "StatusGraphicsLeadlinkNotEnumerated"
	case 0xC01E0432:
		return "StatusGraphicsChainlinksNotEnumerated"
	case 0xC01E0433:
		return "StatusGraphicsAdapterChainNotReady"
	case 0xC01E0434:
		return "StatusGraphicsChainlinksNotStarted"
	case 0xC01E0435:
		return "StatusGraphicsChainlinksNotPoweredOn"
	case 0xC01E0436:
		return "StatusGraphicsInconsistentDeviceLinkState"
	case 0x401E0437:
		return "StatusGraphicsLeadlinkStartDeferred"
	case 0xC01E0438:
		return "StatusGraphicsNotPostDeviceDriver"
	case 0x401E0439:
		return "StatusGraphicsPollingTooFrequently"
	case 0x401E043A:
		return "StatusGraphicsStartDeferred"
	case 0xC01E043B:
		return "StatusGraphicsAdapterAccessNotExcluded"
	case 0xC01E0500:
		return "StatusGraphicsOpmNotSupported"
	case 0xC01E0501:
		return "StatusGraphicsCoppNotSupported"
	case 0xC01E0502:
		return "StatusGraphicsUabNotSupported"
	case 0xC01E0503:
		return "StatusGraphicsOpmInvalidEncryptedParameters"
	case 0xC01E0505:
		return "StatusGraphicsOpmNoProtectedOutputsExist"
	case 0xC01E050B:
		return "StatusGraphicsOpmInternalError"
	case 0xC01E050C:
		return "StatusGraphicsOpmInvalidHandle"
	case 0xC01E050E:
		return "StatusGraphicsPvpInvalidCertificateLength"
	case 0xC01E050F:
		return "StatusGraphicsOpmSpanningModeEnabled"
	case 0xC01E0510:
		return "StatusGraphicsOpmTheaterModeEnabled"
	case 0xC01E0511:
		return "StatusGraphicsPvpHfsFailed"
	case 0xC01E0512:
		return "StatusGraphicsOpmInvalidSrm"
	case 0xC01E0513:
		return "StatusGraphicsOpmOutputDoesNotSupportHdcp"
	case 0xC01E0514:
		return "StatusGraphicsOpmOutputDoesNotSupportAcp"
	case 0xC01E0515:
		return "StatusGraphicsOpmOutputDoesNotSupportCgmsa"
	case 0xC01E0516:
		return "StatusGraphicsOpmHdcpSrmNeverSet"
	case 0xC01E0517:
		return "StatusGraphicsOpmResolutionTooHigh"
	case 0xC01E0518:
		return "StatusGraphicsOpmAllHdcpHardwareAlreadyInUse"
	case 0xC01E051A:
		return "StatusGraphicsOpmProtectedOutputNoLongerExists"
	case 0xC01E051C:
		return "StatusGraphicsOpmProtectedOutputDoesNotHaveCoppSemantics"
	case 0xC01E051D:
		return "StatusGraphicsOpmInvalidInformationRequest"
	case 0xC01E051E:
		return "StatusGraphicsOpmDriverInternalError"
	case 0xC01E051F:
		return "StatusGraphicsOpmProtectedOutputDoesNotHaveOpmSemantics"
	case 0xC01E0520:
		return "StatusGraphicsOpmSignalingNotSupported"
	case 0xC01E0521:
		return "StatusGraphicsOpmInvalidConfigurationRequest"
	case 0xC01E0580:
		return "StatusGraphicsI2CNotSupported"
	case 0xC01E0581:
		return "StatusGraphicsI2CDeviceDoesNotExist"
	case 0xC01E0582:
		return "StatusGraphicsI2CErrorTransmittingData"
	case 0xC01E0583:
		return "StatusGraphicsI2CErrorReceivingData"
	case 0xC01E0584:
		return "StatusGraphicsDdcciVcpNotSupported"
	case 0xC01E0585:
		return "StatusGraphicsDdcciInvalidData"
	case 0xC01E0586:
		return "StatusGraphicsDdcciMonitorReturnedInvalidTimingStatusByte"
	case 0xC01E0587:
		return "StatusGraphicsDdcciInvalidCapabilitiesString"
	case 0xC01E0588:
		return "StatusGraphicsMcaInternalError"
	case 0xC01E0589:
		return "StatusGraphicsDdcciInvalidMessageCommand"
	case 0xC01E058A:
		return "StatusGraphicsDdcciInvalidMessageLength"
	case 0xC01E058B:
		return "StatusGraphicsDdcciInvalidMessageChecksum"
	case 0xC01E058C:
		return "StatusGraphicsInvalidPhysicalMonitorHandle"
	case 0xC01E058D:
		return "StatusGraphicsMonitorNoLongerExists"
	case 0xC01E05E0:
		return "StatusGraphicsOnlyConsoleSessionSupported"
	case 0xC01E05E1:
		return "StatusGraphicsNoDisplayDeviceCorrespondsToName"
	case 0xC01E05E2:
		return "StatusGraphicsDisplayDeviceNotAttachedToDesktop"
	case 0xC01E05E3:
		return "StatusGraphicsMirroringDevicesNotSupported"
	case 0xC01E05E4:
		return "StatusGraphicsInvalidPointer"
	case 0xC01E05E5:
		return "StatusGraphicsNoMonitorsCorrespondToDisplayDevice"
	case 0xC01E05E6:
		return "StatusGraphicsParameterArrayTooSmall"
	case 0xC01E05E7:
		return "StatusGraphicsInternalError"
	case 0xC01E05E8:
		return "StatusGraphicsSessionTypeChangeInProgress"
	case 0xC0210000:
		return "StatusFveLockedVolume"
	case 0xC0210001:
		return "StatusFveNotEncrypted"
	case 0xC0210002:
		return "StatusFveBadInformation"
	case 0xC0210003:
		return "StatusFveTooSmall"
	case 0xC0210004:
		return "StatusFveFailedWrongFs"
	case 0xC0210005:
		return "StatusFveBadPartitionSize"
	case 0xC0210006:
		return "StatusFveFsNotExtended"
	case 0xC0210007:
		return "StatusFveFsMounted"
	case 0xC0210008:
		return "StatusFveNoLicense"
	case 0xC0210009:
		return "StatusFveActionNotAllowed"
	case 0xC021000A:
		return "StatusFveBadData"
	case 0xC021000B:
		return "StatusFveVolumeNotBound"
	case 0xC021000C:
		return "StatusFveNotDataVolume"
	case 0xC021000D:
		return "StatusFveConvReadError"
	case 0xC021000E:
		return "StatusFveConvWriteError"
	case 0xC021000F:
		return "StatusFveOverlappedUpdate"
	case 0xC0210010:
		return "StatusFveFailedSectorSize"
	case 0xC0210011:
		return "StatusFveFailedAuthentication"
	case 0xC0210012:
		return "StatusFveNotOsVolume"
	case 0xC0210013:
		return "StatusFveKeyfileNotFound"
	case 0xC0210014:
		return "StatusFveKeyfileInvalid"
	case 0xC0210015:
		return "StatusFveKeyfileNoVmk"
	case 0xC0210016:
		return "StatusFveTpmDisabled"
	case 0xC0210017:
		return "StatusFveTpmSrkAuthNotZero"
	case 0xC0210018:
		return "StatusFveTpmInvalidPcr"
	case 0xC0210019:
		return "StatusFveTpmNoVmk"
	case 0xC021001A:
		return "StatusFvePinInvalid"
	case 0xC021001B:
		return "StatusFveAuthInvalidApplication"
	case 0xC021001C:
		return "StatusFveAuthInvalidConfig"
	case 0xC021001D:
		return "StatusFveDebuggerEnabled"
	case 0xC021001E:
		return "StatusFveDryRunFailed"
	case 0xC021001F:
		return "StatusFveBadMetadataPointer"
	case 0xC0210020:
		return "StatusFveOldMetadataCopy"
	case 0xC0210021:
		return "StatusFveRebootRequired"
	case 0xC0210022:
		return "StatusFveRawAccess"
	case 0xC0210023:
		return "StatusFveRawBlocked"
	case 0xC0210024:
		return "StatusFveNoAutounlockMasterKey"
	case 0xC0210025:
		return "StatusFveMorFailed"
	case 0xC0210026:
		return "StatusFveNoFeatureLicense"
	case 0xC0210027:
		return "StatusFvePolicyUserDisableRdvNotAllowed"
	case 0xC0210028:
		return "StatusFveConvRecoveryFailed"
	case 0xC0210029:
		return "StatusFveVirtualizedSpaceTooBig"
	case 0xC021002A:
		return "StatusFveInvalidDatumType"
	case 0xC0210030:
		return "StatusFveVolumeTooSmall"
	case 0xC0210031:
		return "StatusFveEnhPinInvalid"
	case 0xC0220001:
		return "StatusFwpCalloutNotFound"
	case 0xC0220002:
		return "StatusFwpConditionNotFound"
	case 0xC0220003:
		return "StatusFwpFilterNotFound"
	case 0xC0220004:
		return "StatusFwpLayerNotFound"
	case 0xC0220005:
		return "StatusFwpProviderNotFound"
	case 0xC0220006:
		return "StatusFwpProviderContextNotFound"
	case 0xC0220007:
		return "StatusFwpSublayerNotFound"
	case 0xC0220008:
		return "StatusFwpNotFound"
	case 0xC0220009:
		return "StatusFwpAlreadyExists"
	case 0xC022000A:
		return "StatusFwpInUse"
	case 0xC022000B:
		return "StatusFwpDynamicSessionInProgress"
	case 0xC022000C:
		return "StatusFwpWrongSession"
	case 0xC022000D:
		return "StatusFwpNoTxnInProgress"
	case 0xC022000E:
		return "StatusFwpTxnInProgress"
	case 0xC022000F:
		return "StatusFwpTxnAborted"
	case 0xC0220010:
		return "StatusFwpSessionAborted"
	case 0xC0220011:
		return "StatusFwpIncompatibleTxn"
	case 0xC0220012:
		return "StatusFwpTimeout"
	case 0xC0220013:
		return "StatusFwpNetEventsDisabled"
	case 0xC0220014:
		return "StatusFwpIncompatibleLayer"
	case 0xC0220015:
		return "StatusFwpKmClientsOnly"
	case 0xC0220016:
		return "StatusFwpLifetimeMismatch"
	case 0xC0220017:
		return "StatusFwpBuiltinObject"
	case 0xC0220018:
		return "StatusFwpTooManyCallouts"
	case 0xC0220019:
		return "StatusFwpNotificationDropped"
	case 0xC022001A:
		return "StatusFwpTrafficMismatch"
	case 0xC022001B:
		return "StatusFwpIncompatibleSaState"
	case 0xC022001C:
		return "StatusFwpNullPointer"
	case 0xC022001D:
		return "StatusFwpInvalidEnumerator"
	case 0xC022001E:
		return "StatusFwpInvalidFlags"
	case 0xC022001F:
		return "StatusFwpInvalidNetMask"
	case 0xC0220020:
		return "StatusFwpInvalidRange"
	case 0xC0220021:
		return "StatusFwpInvalidInterval"
	case 0xC0220022:
		return "StatusFwpZeroLengthArray"
	case 0xC0220023:
		return "StatusFwpNullDisplayName"
	case 0xC0220024:
		return "StatusFwpInvalidActionType"
	case 0xC0220025:
		return "StatusFwpInvalidWeight"
	case 0xC0220026:
		return "StatusFwpMatchTypeMismatch"
	case 0xC0220027:
		return "StatusFwpTypeMismatch"
	case 0xC0220028:
		return "StatusFwpOutOfBounds"
	case 0xC0220029:
		return "StatusFwpReserved"
	case 0xC022002A:
		return "StatusFwpDuplicateCondition"
	case 0xC022002B:
		return "StatusFwpDuplicateKeymod"
	case 0xC022002C:
		return "StatusFwpActionIncompatibleWithLayer"
	case 0xC022002D:
		return "StatusFwpActionIncompatibleWithSublayer"
	case 0xC022002E:
		return "StatusFwpContextIncompatibleWithLayer"
	case 0xC022002F:
		return "StatusFwpContextIncompatibleWithCallout"
	case 0xC0220030:
		return "StatusFwpIncompatibleAuthMethod"
	case 0xC0220031:
		return "StatusFwpIncompatibleDhGroup"
	case 0xC0220032:
		return "StatusFwpEmNotSupported"
	case 0xC0220033:
		return "StatusFwpNeverMatch"
	case 0xC0220034:
		return "StatusFwpProviderContextMismatch"
	case 0xC0220035:
		return "StatusFwpInvalidParameter"
	case 0xC0220036:
		return "StatusFwpTooManySublayers"
	case 0xC0220037:
		return "StatusFwpCalloutNotificationFailed"
	case 0xC0220038:
		return "StatusFwpInvalidAuthTransform"
	case 0xC0220039:
		return "StatusFwpInvalidCipherTransform"
	case 0xC022003A:
		return "StatusFwpIncompatibleCipherTransform"
	case 0xC022003B:
		return "StatusFwpInvalidTransformCombination"
	case 0xC022003C:
		return "StatusFwpDuplicateAuthMethod"
	case 0xC0220100:
		return "StatusFwpTcpipNotReady"
	case 0xC0220101:
		return "StatusFwpInjectHandleClosing"
	case 0xC0220102:
		return "StatusFwpInjectHandleStale"
	case 0xC0220103:
		return "StatusFwpCannotPend"
	case 0xC0220104:
		return "StatusFwpDropNoicmp"
	case 0xC0230002:
		return "StatusNdisClosing"
	case 0xC0230004:
		return "StatusNdisBadVersion"
	case 0xC0230005:
		return "StatusNdisBadCharacteristics"
	case 0xC0230006:
		return "StatusNdisAdapterNotFound"
	case 0xC0230007:
		return "StatusNdisOpenFailed"
	case 0xC0230008:
		return "StatusNdisDeviceFailed"
	case 0xC0230009:
		return "StatusNdisMulticastFull"
	case 0xC023000A:
		return "StatusNdisMulticastExists"
	case 0xC023000B:
		return "StatusNdisMulticastNotFound"
	case 0xC023000C:
		return "StatusNdisRequestAborted"
	case 0xC023000D:
		return "StatusNdisResetInProgress"
	case 0xC02300BB:
		return "StatusNdisNotSupported"
	case 0xC023000F:
		return "StatusNdisInvalidPacket"
	case 0xC0230011:
		return "StatusNdisAdapterNotReady"
	case 0xC0230014:
		return "StatusNdisInvalidLength"
	case 0xC0230015:
		return "StatusNdisInvalidData"
	case 0xC0230016:
		return "StatusNdisBufferTooShort"
	case 0xC0230017:
		return "StatusNdisInvalidOid"
	case 0xC0230018:
		return "StatusNdisAdapterRemoved"
	case 0xC0230019:
		return "StatusNdisUnsupportedMedia"
	case 0xC023001A:
		return "StatusNdisGroupAddressInUse"
	case 0xC023001B:
		return "StatusNdisFileNotFound"
	case 0xC023001C:
		return "StatusNdisErrorReadingFile"
	case 0xC023001D:
		return "StatusNdisAlreadyMapped"
	case 0xC023001E:
		return "StatusNdisResourceConflict"
	case 0xC023001F:
		return "StatusNdisMediaDisconnected"
	case 0xC0230022:
		return "StatusNdisInvalidAddress"
	case 0xC0230010:
		return "StatusNdisInvalidDeviceRequest"
	case 0xC023002A:
		return "StatusNdisPaused"
	case 0xC023002B:
		return "StatusNdisInterfaceNotFound"
	case 0xC023002C:
		return "StatusNdisUnsupportedRevision"
	case 0xC023002D:
		return "StatusNdisInvalidPort"
	case 0xC023002E:
		return "StatusNdisInvalidPortState"
	case 0xC023002F:
		return "StatusNdisLowPowerState"
	case 0xC0232000:
		return "StatusNdisDot11AutoConfigEnabled"
	case 0xC0232001:
		return "StatusNdisDot11MediaInUse"
	case 0xC0232002:
		return "StatusNdisDot11PowerStateInvalid"
	case 0xC0232003:
		return "StatusNdisPmWolPatternListFull"
	case 0xC0232004:
		return "StatusNdisPmProtocolOffloadListFull"
	case 0x40230001:
		return "StatusNdisIndicationRequired"
	case 0xC023100F:
		return "StatusNdisOffloadPolicy"
	case 0xC0231012:
		return "StatusNdisOffloadConnectionRejected"
	case 0xC0231013:
		return "StatusNdisOffloadPathRejected"
	case 0xC0350002:
		return "StatusHvInvalidHypercallCode"
	case 0xC0350003:
		return "StatusHvInvalidHypercallInput"
	case 0xC0350004:
		return "StatusHvInvalidAlignment"
	case 0xC0350005:
		return "StatusHvInvalidParameter"
	case 0xC0350006:
		return "StatusHvAccessDenied"
	case 0xC0350007:
		return "StatusHvInvalidPartitionState"
	case 0xC0350008:
		return "StatusHvOperationDenied"
	case 0xC0350009:
		return "StatusHvUnknownProperty"
	case 0xC035000A:
		return "StatusHvPropertyValueOutOfRange"
	case 0xC035000B:
		return "StatusHvInsufficientMemory"
	case 0xC035000C:
		return "StatusHvPartitionTooDeep"
	case 0xC035000D:
		return "StatusHvInvalidPartitionId"
	case 0xC035000E:
		return "StatusHvInvalidVpIndex"
	case 0xC0350011:
		return "StatusHvInvalidPortId"
	case 0xC0350012:
		return "StatusHvInvalidConnectionId"
	case 0xC0350013:
		return "StatusHvInsufficientBuffers"
	case 0xC0350014:
		return "StatusHvNotAcknowledged"
	case 0xC0350016:
		return "StatusHvAcknowledged"
	case 0xC0350017:
		return "StatusHvInvalidSaveRestoreState"
	case 0xC0350018:
		return "StatusHvInvalidSynicState"
	case 0xC0350019:
		return "StatusHvObjectInUse"
	case 0xC035001A:
		return "StatusHvInvalidProximityDomainInfo"
	case 0xC035001B:
		return "StatusHvNoData"
	case 0xC035001C:
		return "StatusHvInactive"
	case 0xC035001D:
		return "StatusHvNoResources"
	case 0xC035001E:
		return "StatusHvFeatureUnavailable"
	case 0xC0351000:
		return "StatusHvNotPresent"
	case 0xC0370001:
		return "StatusVidDuplicateHandler"
	case 0xC0370002:
		return "StatusVidTooManyHandlers"
	case 0xC0370003:
		return "StatusVidQueueFull"
	case 0xC0370004:
		return "StatusVidHandlerNotPresent"
	case 0xC0370005:
		return "StatusVidInvalidObjectName"
	case 0xC0370006:
		return "StatusVidPartitionNameTooLong"
	case 0xC0370007:
		return "StatusVidMessageQueueNameTooLong"
	case 0xC0370008:
		return "StatusVidPartitionAlreadyExists"
	case 0xC0370009:
		return "StatusVidPartitionDoesNotExist"
	case 0xC037000A:
		return "StatusVidPartitionNameNotFound"
	case 0xC037000B:
		return "StatusVidMessageQueueAlreadyExists"
	case 0xC037000C:
		return "StatusVidExceededMbpEntryMapLimit"
	case 0xC037000D:
		return "StatusVidMbStillReferenced"
	case 0xC037000E:
		return "StatusVidChildGpaPageSetCorrupted"
	case 0xC037000F:
		return "StatusVidInvalidNumaSettings"
	case 0xC0370010:
		return "StatusVidInvalidNumaNodeIndex"
	case 0xC0370011:
		return "StatusVidNotificationQueueAlreadyAssociated"
	case 0xC0370012:
		return "StatusVidInvalidMemoryBlockHandle"
	case 0xC0370013:
		return "StatusVidPageRangeOverflow"
	case 0xC0370014:
		return "StatusVidInvalidMessageQueueHandle"
	case 0xC0370015:
		return "StatusVidInvalidGpaRangeHandle"
	case 0xC0370016:
		return "StatusVidNoMemoryBlockNotificationQueue"
	case 0xC0370017:
		return "StatusVidMemoryBlockLockCountExceeded"
	case 0xC0370018:
		return "StatusVidInvalidPpmHandle"
	case 0xC0370019:
		return "StatusVidMbpsAreLocked"
	case 0xC037001A:
		return "StatusVidMessageQueueClosed"
	case 0xC037001B:
		return "StatusVidVirtualProcessorLimitExceeded"
	case 0xC037001C:
		return "StatusVidStopPending"
	case 0xC037001D:
		return "StatusVidInvalidProcessorState"
	case 0xC037001E:
		return "StatusVidExceededKmContextCountLimit"
	case 0xC037001F:
		return "StatusVidKmInterfaceAlreadyInitialized"
	case 0xC0370020:
		return "StatusVidMbPropertyAlreadySetReset"
	case 0xC0370021:
		return "StatusVidMmioRangeDestroyed"
	case 0xC0370022:
		return "StatusVidInvalidChildGpaPageSet"
	case 0xC0370023:
		return "StatusVidReservePageSetIsBeingUsed"
	case 0xC0370024:
		return "StatusVidReservePageSetTooSmall"
	case 0xC0370025:
		return "StatusVidMbpAlreadyLockedUsingReservedPage"
	case 0xC0370026:
		return "StatusVidMbpCountExceededLimit"
	case 0xC0370027:
		return "StatusVidSavedStateCorrupt"
	case 0xC0370028:
		return "StatusVidSavedStateUnrecognizedItem"
	case 0xC0370029:
		return "StatusVidSavedStateIncompatible"
	case 0x80370001:
		return "StatusVidRemoteNodeParentGpaPagesUsed"
	case 0xC0360001:
		return "StatusIpsecBadSpi"
	case 0xC0360002:
		return "StatusIpsecSaLifetimeExpired"
	case 0xC0360003:
		return "StatusIpsecWrongSa"
	case 0xC0360004:
		return "StatusIpsecReplayCheckFailed"
	case 0xC0360005:
		return "StatusIpsecInvalidPacket"
	case 0xC0360006:
		return "StatusIpsecIntegrityCheckFailed"
	case 0xC0360007:
		return "StatusIpsecClearTextDrop"
	case 0xC0360008:
		return "StatusIpsecAuthFirewallDrop"
	case 0xC0360009:
		return "StatusIpsecThrottleDrop"
	case 0xC0368000:
		return "StatusIpsecDospBlock"
	case 0xC0368001:
		return "StatusIpsecDospReceivedMulticast"
	case 0xC0368002:
		return "StatusIpsecDospInvalidPacket"
	case 0xC0368003:
		return "StatusIpsecDospStateLookupFailed"
	case 0xC0368004:
		return "StatusIpsecDospMaxEntries"
	case 0xC0368005:
		return "StatusIpsecDospKeymodNotAllowed"
	case 0xC0368006:
		return "StatusIpsecDospMaxPerIpRatelimitQueues"
	case 0x80380001:
		return "StatusVolmgrIncompleteRegeneration"
	case 0x80380002:
		return "StatusVolmgrIncompleteDiskMigration"
	case 0xC0380001:
		return "StatusVolmgrDatabaseFull"
	case 0xC0380002:
		return "StatusVolmgrDiskConfigurationCorrupted"
	case 0xC0380003:
		return "StatusVolmgrDiskConfigurationNotInSync"
	case 0xC0380004:
		return "StatusVolmgrPackConfigUpdateFailed"
	case 0xC0380005:
		return "StatusVolmgrDiskContainsNonSimpleVolume"
	case 0xC0380006:
		return "StatusVolmgrDiskDuplicate"
	case 0xC0380007:
		return "StatusVolmgrDiskDynamic"
	case 0xC0380008:
		return "StatusVolmgrDiskIdInvalid"
	case 0xC0380009:
		return "StatusVolmgrDiskInvalid"
	case 0xC038000A:
		return "StatusVolmgrDiskLastVoter"
	case 0xC038000B:
		return "StatusVolmgrDiskLayoutInvalid"
	case 0xC038000C:
		return "StatusVolmgrDiskLayoutNonBasicBetweenBasicPartitions"
	case 0xC038000D:
		return "StatusVolmgrDiskLayoutNotCylinderAligned"
	case 0xC038000E:
		return "StatusVolmgrDiskLayoutPartitionsTooSmall"
	case 0xC038000F:
		return "StatusVolmgrDiskLayoutPrimaryBetweenLogicalPartitions"
	case 0xC0380010:
		return "StatusVolmgrDiskLayoutTooManyPartitions"
	case 0xC0380011:
		return "StatusVolmgrDiskMissing"
	case 0xC0380012:
		return "StatusVolmgrDiskNotEmpty"
	case 0xC0380013:
		return "StatusVolmgrDiskNotEnoughSpace"
	case 0xC0380014:
		return "StatusVolmgrDiskRevectoringFailed"
	case 0xC0380015:
		return "StatusVolmgrDiskSectorSizeInvalid"
	case 0xC0380016:
		return "StatusVolmgrDiskSetNotContained"
	case 0xC0380017:
		return "StatusVolmgrDiskUsedByMultipleMembers"
	case 0xC0380018:
		return "StatusVolmgrDiskUsedByMultiplePlexes"
	case 0xC0380019:
		return "StatusVolmgrDynamicDiskNotSupported"
	case 0xC038001A:
		return "StatusVolmgrExtentAlreadyUsed"
	case 0xC038001B:
		return "StatusVolmgrExtentNotContiguous"
	case 0xC038001C:
		return "StatusVolmgrExtentNotInPublicRegion"
	case 0xC038001D:
		return "StatusVolmgrExtentNotSectorAligned"
	case 0xC038001E:
		return "StatusVolmgrExtentOverlapsEbrPartition"
	case 0xC038001F:
		return "StatusVolmgrExtentVolumeLengthsDoNotMatch"
	case 0xC0380020:
		return "StatusVolmgrFaultTolerantNotSupported"
	case 0xC0380021:
		return "StatusVolmgrInterleaveLengthInvalid"
	case 0xC0380022:
		return "StatusVolmgrMaximumRegisteredUsers"
	case 0xC0380023:
		return "StatusVolmgrMemberInSync"
	case 0xC0380024:
		return "StatusVolmgrMemberIndexDuplicate"
	case 0xC0380025:
		return "StatusVolmgrMemberIndexInvalid"
	case 0xC0380026:
		return "StatusVolmgrMemberMissing"
	case 0xC0380027:
		return "StatusVolmgrMemberNotDetached"
	case 0xC0380028:
		return "StatusVolmgrMemberRegenerating"
	case 0xC0380029:
		return "StatusVolmgrAllDisksFailed"
	case 0xC038002A:
		return "StatusVolmgrNoRegisteredUsers"
	case 0xC038002B:
		return "StatusVolmgrNoSuchUser"
	case 0xC038002C:
		return "StatusVolmgrNotificationReset"
	case 0xC038002D:
		return "StatusVolmgrNumberOfMembersInvalid"
	case 0xC038002E:
		return "StatusVolmgrNumberOfPlexesInvalid"
	case 0xC038002F:
		return "StatusVolmgrPackDuplicate"
	case 0xC0380030:
		return "StatusVolmgrPackIdInvalid"
	case 0xC0380031:
		return "StatusVolmgrPackInvalid"
	case 0xC0380032:
		return "StatusVolmgrPackNameInvalid"
	case 0xC0380033:
		return "StatusVolmgrPackOffline"
	case 0xC0380034:
		return "StatusVolmgrPackHasQuorum"
	case 0xC0380035:
		return "StatusVolmgrPackWithoutQuorum"
	case 0xC0380036:
		return "StatusVolmgrPartitionStyleInvalid"
	case 0xC0380037:
		return "StatusVolmgrPartitionUpdateFailed"
	case 0xC0380038:
		return "StatusVolmgrPlexInSync"
	case 0xC0380039:
		return "StatusVolmgrPlexIndexDuplicate"
	case 0xC038003A:
		return "StatusVolmgrPlexIndexInvalid"
	case 0xC038003B:
		return "StatusVolmgrPlexLastActive"
	case 0xC038003C:
		return "StatusVolmgrPlexMissing"
	case 0xC038003D:
		return "StatusVolmgrPlexRegenerating"
	case 0xC038003E:
		return "StatusVolmgrPlexTypeInvalid"
	case 0xC038003F:
		return "StatusVolmgrPlexNotRaid5"
	case 0xC0380040:
		return "StatusVolmgrPlexNotSimple"
	case 0xC0380041:
		return "StatusVolmgrStructureSizeInvalid"
	case 0xC0380042:
		return "StatusVolmgrTooManyNotificationRequests"
	case 0xC0380043:
		return "StatusVolmgrTransactionInProgress"
	case 0xC0380044:
		return "StatusVolmgrUnexpectedDiskLayoutChange"
	case 0xC0380045:
		return "StatusVolmgrVolumeContainsMissingDisk"
	case 0xC0380046:
		return "StatusVolmgrVolumeIdInvalid"
	case 0xC0380047:
		return "StatusVolmgrVolumeLengthInvalid"
	case 0xC0380048:
		return "StatusVolmgrVolumeLengthNotSectorSizeMultiple"
	case 0xC0380049:
		return "StatusVolmgrVolumeNotMirrored"
	case 0xC038004A:
		return "StatusVolmgrVolumeNotRetained"
	case 0xC038004B:
		return "StatusVolmgrVolumeOffline"
	case 0xC038004C:
		return "StatusVolmgrVolumeRetained"
	case 0xC038004D:
		return "StatusVolmgrNumberOfExtentsInvalid"
	case 0xC038004E:
		return "StatusVolmgrDifferentSectorSize"
	case 0xC038004F:
		return "StatusVolmgrBadBootDisk"
	case 0xC0380050:
		return "StatusVolmgrPackConfigOffline"
	case 0xC0380051:
		return "StatusVolmgrPackConfigOnline"
	case 0xC0380052:
		return "StatusVolmgrNotPrimaryPack"
	case 0xC0380053:
		return "StatusVolmgrPackLogUpdateFailed"
	case 0xC0380054:
		return "StatusVolmgrNumberOfDisksInPlexInvalid"
	case 0xC0380055:
		return "StatusVolmgrNumberOfDisksInMemberInvalid"
	case 0xC0380056:
		return "StatusVolmgrVolumeMirrored"
	case 0xC0380057:
		return "StatusVolmgrPlexNotSimpleSpanned"
	case 0xC0380058:
		return "StatusVolmgrNoValidLogCopies"
	case 0xC0380059:
		return "StatusVolmgrPrimaryPackPresent"
	case 0xC038005A:
		return "StatusVolmgrNumberOfDisksInvalid"
	case 0xC038005B:
		return "StatusVolmgrMirrorNotSupported"
	case 0xC038005C:
		return "StatusVolmgrRaid5NotSupported"
	case 0x80390001:
		return "StatusBcdNotAllEntriesImported"
	case 0xC0390002:
		return "StatusBcdTooManyElements"
	case 0x80390003:
		return "StatusBcdNotAllEntriesSynchronized"
	case 0xC03A0001:
		return "StatusVhdDriveFooterMissing"
	case 0xC03A0002:
		return "StatusVhdDriveFooterChecksumMismatch"
	case 0xC03A0003:
		return "StatusVhdDriveFooterCorrupt"
	case 0xC03A0004:
		return "StatusVhdFormatUnknown"
	case 0xC03A0005:
		return "StatusVhdFormatUnsupportedVersion"
	case 0xC03A0006:
		return "StatusVhdSparseHeaderChecksumMismatch"
	case 0xC03A0007:
		return "StatusVhdSparseHeaderUnsupportedVersion"
	case 0xC03A0008:
		return "StatusVhdSparseHeaderCorrupt"
	case 0xC03A0009:
		return "StatusVhdBlockAllocationFailure"
	case 0xC03A000A:
		return "StatusVhdBlockAllocationTableCorrupt"
	case 0xC03A000B:
		return "StatusVhdInvalidBlockSize"
	case 0xC03A000C:
		return "StatusVhdBitmapMismatch"
	case 0xC03A000D:
		return "StatusVhdParentVhdNotFound"
	case 0xC03A000E:
		return "StatusVhdChildParentIdMismatch"
	case 0xC03A000F:
		return "StatusVhdChildParentTimestampMismatch"
	case 0xC03A0010:
		return "StatusVhdMetadataReadFailure"
	case 0xC03A0011:
		return "StatusVhdMetadataWriteFailure"
	case 0xC03A0012:
		return "StatusVhdInvalidSize"
	case 0xC03A0013:
		return "StatusVhdInvalidFileSize"
	case 0xC03A0014:
		return "StatusVirtdiskProviderNotFound"
	case 0xC03A0015:
		return "StatusVirtdiskNotVirtualDisk"
	case 0xC03A0016:
		return "StatusVhdParentVhdAccessDenied"
	case 0xC03A0017:
		return "StatusVhdChildParentSizeMismatch"
	case 0xC03A0018:
		return "StatusVhdDifferencingChainCycleDetected"
	case 0xC03A0019:
		return "StatusVhdDifferencingChainErrorInParent"
	case 0xC03A001A:
		return "StatusVirtualDiskLimitation"
	case 0xC03A001B:
		return "StatusVhdInvalidType"
	case 0xC03A001C:
		return "StatusVhdInvalidState"
	case 0xC03A001D:
		return "StatusVirtdiskUnsupportedDiskSectorSize"
	case 0x803A0001:
		return "StatusQueryStorageError"
	case 0xC03C0001:
		return "StatusDisNotPresent"
	case 0xC03C0002:
		return "StatusDisAttributeNotFound"
	case 0xC03C0003:
		return "StatusDisUnrecognizedAttribute"
	case 0xC03C0004:
		return "StatusDisPartialData"
	default:
		return "unknown NtstatusKind " + fmt.Sprint(k)
	}
}

func (k NtstatusKind) Elements() []NtstatusKind {
	return []NtstatusKind{
		StatusSuccess,
		StatusWait1,
		StatusWait2,
		StatusWait3,
		StatusWait63,
		StatusAbandoned,
		StatusAbandonedWait63,
		StatusUserApc,
		StatusKernelApc,
		StatusAlerted,
		StatusTimeout,
		StatusPending,
		StatusReparse,
		StatusMoreEntries,
		StatusNotAllAssigned,
		StatusSomeNotMapped,
		StatusOplockBreakInProgress,
		StatusVolumeMounted,
		StatusRxactCommitted,
		StatusNotifyCleanup,
		StatusNotifyEnumDir,
		StatusNoQuotasForAccount,
		StatusPrimaryTransportConnectFailed,
		StatusPageFaultTransition,
		StatusPageFaultDemandZero,
		StatusPageFaultCopyOnWrite,
		StatusPageFaultGuardPage,
		StatusPageFaultPagingFile,
		StatusCachePageLocked,
		StatusCrashDump,
		StatusBufferAllZeros,
		StatusReparseObject,
		StatusResourceRequirementsChanged,
		StatusTranslationComplete,
		StatusDsMembershipEvaluatedLocally,
		StatusNothingToTerminate,
		StatusProcessNotInJob,
		StatusProcessInJob,
		StatusVolsnapHibernateReady,
		StatusFsfilterOpCompletedSuccessfully,
		StatusInterruptVectorAlreadyConnected,
		StatusInterruptStillConnected,
		StatusProcessCloned,
		StatusFileLockedWithOnlyReaders,
		StatusFileLockedWithWriters,
		StatusResourcemanagerReadOnly,
		StatusRingPreviouslyEmpty,
		StatusRingPreviouslyFull,
		StatusRingPreviouslyAboveQuota,
		StatusRingNewlyEmpty,
		StatusRingSignalOppositeEndpoint,
		StatusOplockSwitchedToNewHandle,
		StatusOplockHandleClosed,
		StatusWaitForOplock,
		DbgExceptionHandled,
		DbgContinue,
		StatusFltIoComplete,
		StatusDisAttributeBuilt,
		StatusObjectNameExists,
		StatusThreadWasSuspended,
		StatusWorkingSetLimitRange,
		StatusImageNotAtBase,
		StatusRxactStateCreated,
		StatusSegmentNotification,
		StatusLocalUserSessionKey,
		StatusBadCurrentDirectory,
		StatusSerialMoreWrites,
		StatusRegistryRecovered,
		StatusFtReadRecoveryFromBackup,
		StatusFtWriteRecovery,
		StatusSerialCounterTimeout,
		StatusNullLmPassword,
		StatusImageMachineTypeMismatch,
		StatusReceivePartial,
		StatusReceiveExpedited,
		StatusReceivePartialExpedited,
		StatusEventDone,
		StatusEventPending,
		StatusCheckingFileSystem,
		StatusFatalAppExit,
		StatusPredefinedHandle,
		StatusWasUnlocked,
		StatusServiceNotification,
		StatusWasLocked,
		StatusLogHardError,
		StatusAlreadyWin32,
		StatusWx86Unsimulate,
		StatusWx86Continue,
		StatusWx86SingleStep,
		StatusWx86Breakpoint,
		StatusWx86ExceptionContinue,
		StatusWx86ExceptionLastchance,
		StatusWx86ExceptionChain,
		StatusImageMachineTypeMismatchExe,
		StatusNoYieldPerformed,
		StatusTimerResumeIgnored,
		StatusArbitrationUnhandled,
		StatusCardbusNotSupported,
		StatusWx86Createwx86Tib,
		StatusMpProcessorMismatch,
		StatusHibernated,
		StatusResumeHibernation,
		StatusFirmwareUpdated,
		StatusDriversLeakingLockedPages,
		StatusMessageRetrieved,
		StatusSystemPowerstateTransition,
		StatusAlpcCheckCompletionList,
		StatusSystemPowerstateComplexTransition,
		StatusAccessAuditByPolicy,
		StatusAbandonHiberfile,
		StatusBizrulesNotEnabled,
		DbgReplyLater,
		DbgUnableToProvideHandle,
		DbgTerminateThread,
		DbgTerminateProcess,
		DbgControlC,
		DbgPrintexceptionC,
		DbgRipexception,
		DbgControlBreak,
		DbgCommandException,
		StatusHeuristicDamagePossible,
		StatusGuardPageViolation,
		StatusDatatypeMisalignment,
		StatusBreakpoint,
		StatusSingleStep,
		StatusBufferOverflow,
		StatusNoMoreFiles,
		StatusWakeSystemDebugger,
		StatusHandlesClosed,
		StatusNoInheritance,
		StatusGuidSubstitutionMade,
		StatusPartialCopy,
		StatusDevicePaperEmpty,
		StatusDevicePoweredOff,
		StatusDeviceOffLine,
		StatusDeviceBusy,
		StatusNoMoreEas,
		StatusInvalidEaName,
		StatusEaListInconsistent,
		StatusInvalidEaFlag,
		StatusVerifyRequired,
		StatusExtraneousInformation,
		StatusRxactCommitNecessary,
		StatusNoMoreEntries,
		StatusFilemarkDetected,
		StatusMediaChanged,
		StatusBusReset,
		StatusEndOfMedia,
		StatusBeginningOfMedia,
		StatusMediaCheck,
		StatusSetmarkDetected,
		StatusNoDataDetected,
		StatusRedirectorHasOpenHandles,
		StatusServerHasOpenHandles,
		StatusAlreadyDisconnected,
		StatusLongjump,
		StatusCleanerCartridgeInstalled,
		StatusPlugplayQueryVetoed,
		StatusUnwindConsolidate,
		StatusRegistryHiveRecovered,
		StatusDllMightBeInsecure,
		StatusDllMightBeIncompatible,
		StatusStoppedOnSymlink,
		StatusCannotGrantRequestedOplock,
		StatusNoAceCondition,
		DbgExceptionNotHandled,
		StatusClusterNodeAlreadyUp,
		StatusClusterNodeAlreadyDown,
		StatusClusterNetworkAlreadyOnline,
		StatusClusterNetworkAlreadyOffline,
		StatusClusterNodeAlreadyMember,
		StatusFltBufferTooSmall,
		StatusFvePartialMetadata,
		StatusFveTransientState,
		StatusUnsuccessful,
		StatusNotImplemented,
		StatusInvalidInfoClass,
		StatusInfoLengthMismatch,
		StatusAccessViolation,
		StatusInPageError,
		StatusPagefileQuota,
		StatusInvalidHandle,
		StatusBadInitialStack,
		StatusBadInitialPc,
		StatusInvalidCid,
		StatusTimerNotCanceled,
		StatusInvalidParameter,
		StatusNoSuchDevice,
		StatusNoSuchFile,
		StatusInvalidDeviceRequest,
		StatusEndOfFile,
		StatusWrongVolume,
		StatusNoMediaInDevice,
		StatusUnrecognizedMedia,
		StatusNonexistentSector,
		StatusMoreProcessingRequired,
		StatusNoMemory,
		StatusConflictingAddresses,
		StatusNotMappedView,
		StatusUnableToFreeVm,
		StatusUnableToDeleteSection,
		StatusInvalidSystemService,
		StatusIllegalInstruction,
		StatusInvalidLockSequence,
		StatusInvalidViewSize,
		StatusInvalidFileForSection,
		StatusAlreadyCommitted,
		StatusAccessDenied,
		StatusBufferTooSmall,
		StatusObjectTypeMismatch,
		StatusNoncontinuableException,
		StatusInvalidDisposition,
		StatusUnwind,
		StatusBadStack,
		StatusInvalidUnwindTarget,
		StatusNotLocked,
		StatusParityError,
		StatusUnableToDecommitVm,
		StatusNotCommitted,
		StatusInvalidPortAttributes,
		StatusPortMessageTooLong,
		StatusInvalidParameterMix,
		StatusInvalidQuotaLower,
		StatusDiskCorruptError,
		StatusObjectNameInvalid,
		StatusObjectNameNotFound,
		StatusObjectNameCollision,
		StatusPortDisconnected,
		StatusDeviceAlreadyAttached,
		StatusObjectPathInvalid,
		StatusObjectPathNotFound,
		StatusObjectPathSyntaxBad,
		StatusDataOverrun,
		StatusDataLateError,
		StatusDataError,
		StatusCrcError,
		StatusSectionTooBig,
		StatusPortConnectionRefused,
		StatusInvalidPortHandle,
		StatusSharingViolation,
		StatusQuotaExceeded,
		StatusInvalidPageProtection,
		StatusMutantNotOwned,
		StatusSemaphoreLimitExceeded,
		StatusPortAlreadySet,
		StatusSectionNotImage,
		StatusSuspendCountExceeded,
		StatusThreadIsTerminating,
		StatusBadWorkingSetLimit,
		StatusIncompatibleFileMap,
		StatusSectionProtection,
		StatusEasNotSupported,
		StatusEaTooLarge,
		StatusNonexistentEaEntry,
		StatusNoEasOnFile,
		StatusEaCorruptError,
		StatusFileLockConflict,
		StatusLockNotGranted,
		StatusDeletePending,
		StatusCtlFileNotSupported,
		StatusUnknownRevision,
		StatusRevisionMismatch,
		StatusInvalidOwner,
		StatusInvalidPrimaryGroup,
		StatusNoImpersonationToken,
		StatusCantDisableMandatory,
		StatusNoLogonServers,
		StatusNoSuchLogonSession,
		StatusNoSuchPrivilege,
		StatusPrivilegeNotHeld,
		StatusInvalidAccountName,
		StatusUserExists,
		StatusNoSuchUser,
		StatusGroupExists,
		StatusNoSuchGroup,
		StatusMemberInGroup,
		StatusMemberNotInGroup,
		StatusLastAdmin,
		StatusWrongPassword,
		StatusIllFormedPassword,
		StatusPasswordRestriction,
		StatusLogonFailure,
		StatusAccountRestriction,
		StatusInvalidLogonHours,
		StatusInvalidWorkstation,
		StatusPasswordExpired,
		StatusAccountDisabled,
		StatusNoneMapped,
		StatusTooManyLuidsRequested,
		StatusLuidsExhausted,
		StatusInvalidSubAuthority,
		StatusInvalidAcl,
		StatusInvalidSid,
		StatusInvalidSecurityDescr,
		StatusProcedureNotFound,
		StatusInvalidImageFormat,
		StatusNoToken,
		StatusBadInheritanceAcl,
		StatusRangeNotLocked,
		StatusDiskFull,
		StatusServerDisabled,
		StatusServerNotDisabled,
		StatusTooManyGuidsRequested,
		StatusGuidsExhausted,
		StatusInvalidIdAuthority,
		StatusAgentsExhausted,
		StatusInvalidVolumeLabel,
		StatusSectionNotExtended,
		StatusNotMappedData,
		StatusResourceDataNotFound,
		StatusResourceTypeNotFound,
		StatusResourceNameNotFound,
		StatusArrayBoundsExceeded,
		StatusFloatDenormalOperand,
		StatusFloatDivideByZero,
		StatusFloatInexactResult,
		StatusFloatInvalidOperation,
		StatusFloatOverflow,
		StatusFloatStackCheck,
		StatusFloatUnderflow,
		StatusIntegerDivideByZero,
		StatusIntegerOverflow,
		StatusPrivilegedInstruction,
		StatusTooManyPagingFiles,
		StatusFileInvalid,
		StatusAllottedSpaceExceeded,
		StatusInsufficientResources,
		StatusDfsExitPathFound,
		StatusDeviceDataError,
		StatusDeviceNotConnected,
		StatusDevicePowerFailure,
		StatusFreeVmNotAtBase,
		StatusMemoryNotAllocated,
		StatusWorkingSetQuota,
		StatusMediaWriteProtected,
		StatusDeviceNotReady,
		StatusInvalidGroupAttributes,
		StatusBadImpersonationLevel,
		StatusCantOpenAnonymous,
		StatusBadValidationClass,
		StatusBadTokenType,
		StatusBadMasterBootRecord,
		StatusInstructionMisalignment,
		StatusInstanceNotAvailable,
		StatusPipeNotAvailable,
		StatusInvalidPipeState,
		StatusPipeBusy,
		StatusIllegalFunction,
		StatusPipeDisconnected,
		StatusPipeClosing,
		StatusPipeConnected,
		StatusPipeListening,
		StatusInvalidReadMode,
		StatusIoTimeout,
		StatusFileForcedClosed,
		StatusProfilingNotStarted,
		StatusProfilingNotStopped,
		StatusCouldNotInterpret,
		StatusFileIsADirectory,
		StatusNotSupported,
		StatusRemoteNotListening,
		StatusDuplicateName,
		StatusBadNetworkPath,
		StatusNetworkBusy,
		StatusDeviceDoesNotExist,
		StatusTooManyCommands,
		StatusAdapterHardwareError,
		StatusInvalidNetworkResponse,
		StatusUnexpectedNetworkError,
		StatusBadRemoteAdapter,
		StatusPrintQueueFull,
		StatusNoSpoolSpace,
		StatusPrintCancelled,
		StatusNetworkNameDeleted,
		StatusNetworkAccessDenied,
		StatusBadDeviceType,
		StatusBadNetworkName,
		StatusTooManyNames,
		StatusTooManySessions,
		StatusSharingPaused,
		StatusRequestNotAccepted,
		StatusRedirectorPaused,
		StatusNetWriteFault,
		StatusProfilingAtLimit,
		StatusNotSameDevice,
		StatusFileRenamed,
		StatusVirtualCircuitClosed,
		StatusNoSecurityOnObject,
		StatusCantWait,
		StatusPipeEmpty,
		StatusCantAccessDomainInfo,
		StatusCantTerminateSelf,
		StatusInvalidServerState,
		StatusInvalidDomainState,
		StatusInvalidDomainRole,
		StatusNoSuchDomain,
		StatusDomainExists,
		StatusDomainLimitExceeded,
		StatusOplockNotGranted,
		StatusInvalidOplockProtocol,
		StatusInternalDbCorruption,
		StatusInternalError,
		StatusGenericNotMapped,
		StatusBadDescriptorFormat,
		StatusInvalidUserBuffer,
		StatusUnexpectedIoError,
		StatusUnexpectedMmCreateErr,
		StatusUnexpectedMmMapError,
		StatusUnexpectedMmExtendErr,
		StatusNotLogonProcess,
		StatusLogonSessionExists,
		StatusInvalidParameter1,
		StatusInvalidParameter2,
		StatusInvalidParameter3,
		StatusInvalidParameter4,
		StatusInvalidParameter5,
		StatusInvalidParameter6,
		StatusInvalidParameter7,
		StatusInvalidParameter8,
		StatusInvalidParameter9,
		StatusInvalidParameter10,
		StatusInvalidParameter11,
		StatusInvalidParameter12,
		StatusRedirectorNotStarted,
		StatusRedirectorStarted,
		StatusStackOverflow,
		StatusNoSuchPackage,
		StatusBadFunctionTable,
		StatusVariableNotFound,
		StatusDirectoryNotEmpty,
		StatusFileCorruptError,
		StatusNotADirectory,
		StatusBadLogonSessionState,
		StatusLogonSessionCollision,
		StatusNameTooLong,
		StatusFilesOpen,
		StatusConnectionInUse,
		StatusMessageNotFound,
		StatusProcessIsTerminating,
		StatusInvalidLogonType,
		StatusNoGuidTranslation,
		StatusCannotImpersonate,
		StatusImageAlreadyLoaded,
		StatusAbiosNotPresent,
		StatusAbiosLidNotExist,
		StatusAbiosLidAlreadyOwned,
		StatusAbiosNotLidOwner,
		StatusAbiosInvalidCommand,
		StatusAbiosInvalidLid,
		StatusAbiosSelectorNotAvailable,
		StatusAbiosInvalidSelector,
		StatusNoLdt,
		StatusInvalidLdtSize,
		StatusInvalidLdtOffset,
		StatusInvalidLdtDescriptor,
		StatusInvalidImageNeFormat,
		StatusRxactInvalidState,
		StatusRxactCommitFailure,
		StatusMappedFileSizeZero,
		StatusTooManyOpenedFiles,
		StatusCancelled,
		StatusCannotDelete,
		StatusInvalidComputerName,
		StatusFileDeleted,
		StatusSpecialAccount,
		StatusSpecialGroup,
		StatusSpecialUser,
		StatusMembersPrimaryGroup,
		StatusFileClosed,
		StatusTooManyThreads,
		StatusThreadNotInProcess,
		StatusTokenAlreadyInUse,
		StatusPagefileQuotaExceeded,
		StatusCommitmentLimit,
		StatusInvalidImageLeFormat,
		StatusInvalidImageNotMz,
		StatusInvalidImageProtect,
		StatusInvalidImageWin16,
		StatusLogonServerConflict,
		StatusTimeDifferenceAtDc,
		StatusSynchronizationRequired,
		StatusDllNotFound,
		StatusOpenFailed,
		StatusIoPrivilegeFailed,
		StatusOrdinalNotFound,
		StatusEntrypointNotFound,
		StatusControlCExit,
		StatusLocalDisconnect,
		StatusRemoteDisconnect,
		StatusRemoteResources,
		StatusLinkFailed,
		StatusLinkTimeout,
		StatusInvalidConnection,
		StatusInvalidAddress,
		StatusDllInitFailed,
		StatusMissingSystemfile,
		StatusUnhandledException,
		StatusAppInitFailure,
		StatusPagefileCreateFailed,
		StatusNoPagefile,
		StatusInvalidLevel,
		StatusWrongPasswordCore,
		StatusIllegalFloatContext,
		StatusPipeBroken,
		StatusRegistryCorrupt,
		StatusRegistryIoFailed,
		StatusNoEventPair,
		StatusUnrecognizedVolume,
		StatusSerialNoDeviceInited,
		StatusNoSuchAlias,
		StatusMemberNotInAlias,
		StatusMemberInAlias,
		StatusAliasExists,
		StatusLogonNotGranted,
		StatusTooManySecrets,
		StatusSecretTooLong,
		StatusInternalDbError,
		StatusFullscreenMode,
		StatusTooManyContextIds,
		StatusLogonTypeNotGranted,
		StatusNotRegistryFile,
		StatusNtCrossEncryptionRequired,
		StatusDomainCtrlrConfigError,
		StatusFtMissingMember,
		StatusIllFormedServiceEntry,
		StatusIllegalCharacter,
		StatusUnmappableCharacter,
		StatusUndefinedCharacter,
		StatusFloppyVolume,
		StatusFloppyIdMarkNotFound,
		StatusFloppyWrongCylinder,
		StatusFloppyUnknownError,
		StatusFloppyBadRegisters,
		StatusDiskRecalibrateFailed,
		StatusDiskOperationFailed,
		StatusDiskResetFailed,
		StatusSharedIrqBusy,
		StatusFtOrphaning,
		StatusBiosFailedToConnectInterrupt,
		StatusPartitionFailure,
		StatusInvalidBlockLength,
		StatusDeviceNotPartitioned,
		StatusUnableToLockMedia,
		StatusUnableToUnloadMedia,
		StatusEomOverflow,
		StatusNoMedia,
		StatusNoSuchMember,
		StatusInvalidMember,
		StatusKeyDeleted,
		StatusNoLogSpace,
		StatusTooManySids,
		StatusLmCrossEncryptionRequired,
		StatusKeyHasChildren,
		StatusChildMustBeVolatile,
		StatusDeviceConfigurationError,
		StatusDriverInternalError,
		StatusInvalidDeviceState,
		StatusIoDeviceError,
		StatusDeviceProtocolError,
		StatusBackupController,
		StatusLogFileFull,
		StatusTooLate,
		StatusNoTrustLsaSecret,
		StatusNoTrustSamAccount,
		StatusTrustedDomainFailure,
		StatusTrustedRelationshipFailure,
		StatusEventlogFileCorrupt,
		StatusEventlogCantStart,
		StatusTrustFailure,
		StatusMutantLimitExceeded,
		StatusNetlogonNotStarted,
		StatusAccountExpired,
		StatusPossibleDeadlock,
		StatusNetworkCredentialConflict,
		StatusRemoteSessionLimit,
		StatusEventlogFileChanged,
		StatusNologonInterdomainTrustAccount,
		StatusNologonWorkstationTrustAccount,
		StatusNologonServerTrustAccount,
		StatusDomainTrustInconsistent,
		StatusFsDriverRequired,
		StatusImageAlreadyLoadedAsDll,
		StatusIncompatibleWithGlobalShortNameRegistrySetting,
		StatusShortNamesNotEnabledOnVolume,
		StatusSecurityStreamIsInconsistent,
		StatusInvalidLockRange,
		StatusInvalidAceCondition,
		StatusImageSubsystemNotPresent,
		StatusNotificationGuidAlreadyDefined,
		StatusNetworkOpenRestriction,
		StatusNoUserSessionKey,
		StatusUserSessionDeleted,
		StatusResourceLangNotFound,
		StatusInsuffServerResources,
		StatusInvalidBufferSize,
		StatusInvalidAddressComponent,
		StatusInvalidAddressWildcard,
		StatusTooManyAddresses,
		StatusAddressAlreadyExists,
		StatusAddressClosed,
		StatusConnectionDisconnected,
		StatusConnectionReset,
		StatusTooManyNodes,
		StatusTransactionAborted,
		StatusTransactionTimedOut,
		StatusTransactionNoRelease,
		StatusTransactionNoMatch,
		StatusTransactionResponded,
		StatusTransactionInvalidId,
		StatusTransactionInvalidType,
		StatusNotServerSession,
		StatusNotClientSession,
		StatusCannotLoadRegistryFile,
		StatusDebugAttachFailed,
		StatusSystemProcessTerminated,
		StatusDataNotAccepted,
		StatusNoBrowserServersFound,
		StatusVdmHardError,
		StatusDriverCancelTimeout,
		StatusReplyMessageMismatch,
		StatusMappedAlignment,
		StatusImageChecksumMismatch,
		StatusLostWritebehindData,
		StatusClientServerParametersInvalid,
		StatusPasswordMustChange,
		StatusNotFound,
		StatusNotTinyStream,
		StatusRecoveryFailure,
		StatusStackOverflowRead,
		StatusFailCheck,
		StatusDuplicateObjectid,
		StatusObjectidExists,
		StatusConvertToLarge,
		StatusRetry,
		StatusFoundOutOfScope,
		StatusAllocateBucket,
		StatusPropsetNotFound,
		StatusMarshallOverflow,
		StatusInvalidVariant,
		StatusDomainControllerNotFound,
		StatusAccountLockedOut,
		StatusHandleNotClosable,
		StatusConnectionRefused,
		StatusGracefulDisconnect,
		StatusAddressAlreadyAssociated,
		StatusAddressNotAssociated,
		StatusConnectionInvalid,
		StatusConnectionActive,
		StatusNetworkUnreachable,
		StatusHostUnreachable,
		StatusProtocolUnreachable,
		StatusPortUnreachable,
		StatusRequestAborted,
		StatusConnectionAborted,
		StatusBadCompressionBuffer,
		StatusUserMappedFile,
		StatusAuditFailed,
		StatusTimerResolutionNotSet,
		StatusConnectionCountLimit,
		StatusLoginTimeRestriction,
		StatusLoginWkstaRestriction,
		StatusImageMpUpMismatch,
		StatusInsufficientLogonInfo,
		StatusBadDllEntrypoint,
		StatusBadServiceEntrypoint,
		StatusLpcReplyLost,
		StatusIpAddressConflict1,
		StatusIpAddressConflict2,
		StatusRegistryQuotaLimit,
		StatusPathNotCovered,
		StatusNoCallbackActive,
		StatusLicenseQuotaExceeded,
		StatusPwdTooShort,
		StatusPwdTooRecent,
		StatusPwdHistoryConflict,
		StatusPlugplayNoDevice,
		StatusUnsupportedCompression,
		StatusInvalidHwProfile,
		StatusInvalidPlugplayDevicePath,
		StatusDriverOrdinalNotFound,
		StatusDriverEntrypointNotFound,
		StatusResourceNotOwned,
		StatusTooManyLinks,
		StatusQuotaListInconsistent,
		StatusFileIsOffline,
		StatusEvaluationExpiration,
		StatusIllegalDllRelocation,
		StatusLicenseViolation,
		StatusDllInitFailedLogoff,
		StatusDriverUnableToLoad,
		StatusDfsUnavailable,
		StatusVolumeDismounted,
		StatusWx86InternalError,
		StatusWx86FloatStackCheck,
		StatusValidateContinue,
		StatusNoMatch,
		StatusNoMoreMatches,
		StatusNotAReparsePoint,
		StatusIoReparseTagInvalid,
		StatusIoReparseTagMismatch,
		StatusIoReparseDataInvalid,
		StatusIoReparseTagNotHandled,
		StatusReparsePointNotResolved,
		StatusDirectoryIsAReparsePoint,
		StatusRangeListConflict,
		StatusSourceElementEmpty,
		StatusDestinationElementFull,
		StatusIllegalElementAddress,
		StatusMagazineNotPresent,
		StatusReinitializationNeeded,
		StatusDeviceRequiresCleaning,
		StatusDeviceDoorOpen,
		StatusEncryptionFailed,
		StatusDecryptionFailed,
		StatusRangeNotFound,
		StatusNoRecoveryPolicy,
		StatusNoEfs,
		StatusWrongEfs,
		StatusNoUserKeys,
		StatusFileNotEncrypted,
		StatusNotExportFormat,
		StatusFileEncrypted,
		StatusWakeSystem,
		StatusWmiGuidNotFound,
		StatusWmiInstanceNotFound,
		StatusWmiItemidNotFound,
		StatusWmiTryAgain,
		StatusSharedPolicy,
		StatusPolicyObjectNotFound,
		StatusPolicyOnlyInDs,
		StatusVolumeNotUpgraded,
		StatusRemoteStorageNotActive,
		StatusRemoteStorageMediaError,
		StatusNoTrackingService,
		StatusServerSidMismatch,
		StatusDsNoAttributeOrValue,
		StatusDsInvalidAttributeSyntax,
		StatusDsAttributeTypeUndefined,
		StatusDsAttributeOrValueExists,
		StatusDsBusy,
		StatusDsUnavailable,
		StatusDsNoRidsAllocated,
		StatusDsNoMoreRids,
		StatusDsIncorrectRoleOwner,
		StatusDsRidmgrInitError,
		StatusDsObjClassViolation,
		StatusDsCantOnNonLeaf,
		StatusDsCantOnRdn,
		StatusDsCantModObjClass,
		StatusDsCrossDomMoveFailed,
		StatusDsGcNotAvailable,
		StatusDirectoryServiceRequired,
		StatusReparseAttributeConflict,
		StatusCantEnableDenyOnly,
		StatusFloatMultipleFaults,
		StatusFloatMultipleTraps,
		StatusDeviceRemoved,
		StatusJournalDeleteInProgress,
		StatusJournalNotActive,
		StatusNointerface,
		StatusDsAdminLimitExceeded,
		StatusDriverFailedSleep,
		StatusMutualAuthenticationFailed,
		StatusCorruptSystemFile,
		StatusDatatypeMisalignmentError,
		StatusWmiReadOnly,
		StatusWmiSetFailure,
		StatusCommitmentMinimum,
		StatusRegNatConsumption,
		StatusTransportFull,
		StatusDsSamInitFailure,
		StatusOnlyIfConnected,
		StatusDsSensitiveGroupViolation,
		StatusPnpRestartEnumeration,
		StatusJournalEntryDeleted,
		StatusDsCantModPrimarygroupid,
		StatusSystemImageBadSignature,
		StatusPnpRebootRequired,
		StatusPowerStateInvalid,
		StatusDsInvalidGroupType,
		StatusDsNoNestGlobalgroupInMixeddomain,
		StatusDsNoNestLocalgroupInMixeddomain,
		StatusDsGlobalCantHaveLocalMember,
		StatusDsGlobalCantHaveUniversalMember,
		StatusDsUniversalCantHaveLocalMember,
		StatusDsGlobalCantHaveCrossdomainMember,
		StatusDsLocalCantHaveCrossdomainLocalMember,
		StatusDsHavePrimaryMembers,
		StatusWmiNotSupported,
		StatusInsufficientPower,
		StatusSamNeedBootkeyPassword,
		StatusSamNeedBootkeyFloppy,
		StatusDsCantStart,
		StatusDsInitFailure,
		StatusSamInitFailure,
		StatusDsGcRequired,
		StatusDsLocalMemberOfLocalOnly,
		StatusDsNoFpoInUniversalGroups,
		StatusDsMachineAccountQuotaExceeded,
		StatusMultipleFaultViolation,
		StatusCurrentDomainNotAllowed,
		StatusCannotMake,
		StatusSystemShutdown,
		StatusDsInitFailureConsole,
		StatusDsSamInitFailureConsole,
		StatusUnfinishedContextDeleted,
		StatusNoTgtReply,
		StatusObjectidNotFound,
		StatusNoIpAddresses,
		StatusWrongCredentialHandle,
		StatusCryptoSystemInvalid,
		StatusMaxReferralsExceeded,
		StatusMustBeKdc,
		StatusStrongCryptoNotSupported,
		StatusTooManyPrincipals,
		StatusNoPaData,
		StatusPkinitNameMismatch,
		StatusSmartcardLogonRequired,
		StatusKdcInvalidRequest,
		StatusKdcUnableToRefer,
		StatusKdcUnknownEtype,
		StatusShutdownInProgress,
		StatusServerShutdownInProgress,
		StatusNotSupportedOnSbs,
		StatusWmiGuidDisconnected,
		StatusWmiAlreadyDisabled,
		StatusWmiAlreadyEnabled,
		StatusMftTooFragmented,
		StatusCopyProtectionFailure,
		StatusCssAuthenticationFailure,
		StatusCssKeyNotPresent,
		StatusCssKeyNotEstablished,
		StatusCssScrambledSector,
		StatusCssRegionMismatch,
		StatusCssResetsExhausted,
		StatusPkinitFailure,
		StatusSmartcardSubsystemFailure,
		StatusNoKerbKey,
		StatusHostDown,
		StatusUnsupportedPreauth,
		StatusEfsAlgBlobTooBig,
		StatusPortNotSet,
		StatusDebuggerInactive,
		StatusDsVersionCheckFailure,
		StatusAuditingDisabled,
		StatusPrent4MachineAccount,
		StatusDsAgCantHaveUniversalMember,
		StatusInvalidImageWin32,
		StatusInvalidImageWin64,
		StatusBadBindings,
		StatusNetworkSessionExpired,
		StatusApphelpBlock,
		StatusAllSidsFiltered,
		StatusNotSafeModeDriver,
		StatusAccessDisabledByPolicyDefault,
		StatusAccessDisabledByPolicyPath,
		StatusAccessDisabledByPolicyPublisher,
		StatusAccessDisabledByPolicyOther,
		StatusFailedDriverEntry,
		StatusDeviceEnumerationError,
		StatusMountPointNotResolved,
		StatusInvalidDeviceObjectParameter,
		StatusMcaOccured,
		StatusDriverBlockedCritical,
		StatusDriverBlocked,
		StatusDriverDatabaseError,
		StatusSystemHiveTooLarge,
		StatusInvalidImportOfNonDll,
		StatusDsShuttingDown,
		StatusNoSecrets,
		StatusAccessDisabledNoSaferUiByPolicy,
		StatusFailedStackSwitch,
		StatusHeapCorruption,
		StatusSmartcardWrongPin,
		StatusSmartcardCardBlocked,
		StatusSmartcardCardNotAuthenticated,
		StatusSmartcardNoCard,
		StatusSmartcardNoKeyContainer,
		StatusSmartcardNoCertificate,
		StatusSmartcardNoKeyset,
		StatusSmartcardIoError,
		StatusDowngradeDetected,
		StatusSmartcardCertRevoked,
		StatusIssuingCaUntrusted,
		StatusRevocationOfflineC,
		StatusPkinitClientFailure,
		StatusSmartcardCertExpired,
		StatusDriverFailedPriorUnload,
		StatusSmartcardSilentContext,
		StatusPerUserTrustQuotaExceeded,
		StatusAllUserTrustQuotaExceeded,
		StatusUserDeleteTrustQuotaExceeded,
		StatusDsNameNotUnique,
		StatusDsDuplicateIdFound,
		StatusDsGroupConversionError,
		StatusVolsnapPrepareHibernate,
		StatusUser2UserRequired,
		StatusStackBufferOverrun,
		StatusNoS4UProtSupport,
		StatusCrossrealmDelegationFailure,
		StatusRevocationOfflineKdc,
		StatusIssuingCaUntrustedKdc,
		StatusKdcCertExpired,
		StatusKdcCertRevoked,
		StatusParameterQuotaExceeded,
		StatusHibernationFailure,
		StatusDelayLoadFailed,
		StatusAuthenticationFirewallFailed,
		StatusVdmDisallowed,
		StatusHungDisplayDriverThread,
		StatusInsufficientResourceForSpecifiedSharedSectionSize,
		StatusInvalidCruntimeParameter,
		StatusNtlmBlocked,
		StatusDsSrcSidExistsInForest,
		StatusDsDomainNameExistsInForest,
		StatusDsFlatNameExistsInForest,
		StatusInvalidUserPrincipalName,
		StatusFatalUserCallbackException,
		StatusAssertionFailure,
		StatusVerifierStop,
		StatusCallbackPopStack,
		StatusIncompatibleDriverBlocked,
		StatusHiveUnloaded,
		StatusCompressionDisabled,
		StatusFileSystemLimitation,
		StatusInvalidImageHash,
		StatusNotCapable,
		StatusRequestOutOfSequence,
		StatusImplementationLimit,
		StatusElevationRequired,
		StatusNoSecurityContext,
		StatusPku2UCertFailure,
		StatusBeyondVdl,
		StatusEncounteredWriteInProgress,
		StatusPteChanged,
		StatusPurgeFailed,
		StatusCredRequiresConfirmation,
		StatusCsEncryptionInvalidServerResponse,
		StatusCsEncryptionUnsupportedServer,
		StatusCsEncryptionExistingEncryptedFile,
		StatusCsEncryptionNewEncryptedFile,
		StatusCsEncryptionFileNotCse,
		StatusInvalidLabel,
		StatusDriverProcessTerminated,
		StatusAmbiguousSystemDevice,
		StatusSystemDeviceNotFound,
		StatusRestartBootApplication,
		StatusInsufficientNvramResources,
		StatusInvalidTaskName,
		StatusInvalidTaskIndex,
		StatusThreadAlreadyInTask,
		StatusCallbackBypass,
		StatusFailFastException,
		StatusImageCertRevoked,
		StatusPortClosed,
		StatusMessageLost,
		StatusInvalidMessage,
		StatusRequestCanceled,
		StatusRecursiveDispatch,
		StatusLpcReceiveBufferExpected,
		StatusLpcInvalidConnectionUsage,
		StatusLpcRequestsNotAllowed,
		StatusResourceInUse,
		StatusHardwareMemoryError,
		StatusThreadpoolHandleException,
		StatusThreadpoolSetEventOnCompletionFailed,
		StatusThreadpoolReleaseSemaphoreOnCompletionFailed,
		StatusThreadpoolReleaseMutexOnCompletionFailed,
		StatusThreadpoolFreeLibraryOnCompletionFailed,
		StatusThreadpoolReleasedDuringOperation,
		StatusCallbackReturnedWhileImpersonating,
		StatusApcReturnedWhileImpersonating,
		StatusProcessIsProtected,
		StatusMcaException,
		StatusCertificateMappingNotUnique,
		StatusSymlinkClassDisabled,
		StatusInvalidIdnNormalization,
		StatusNoUnicodeTranslation,
		StatusAlreadyRegistered,
		StatusContextMismatch,
		StatusPortAlreadyHasCompletionList,
		StatusCallbackReturnedThreadPriority,
		StatusInvalidThread,
		StatusCallbackReturnedTransaction,
		StatusCallbackReturnedLdrLock,
		StatusCallbackReturnedLang,
		StatusCallbackReturnedPriBack,
		StatusCallbackReturnedThreadAffinity,
		StatusDiskRepairDisabled,
		StatusDsDomainRenameInProgress,
		StatusDiskQuotaExceeded,
		StatusDataLostRepair,
		StatusContentBlocked,
		StatusBadClusters,
		StatusVolumeDirty,
		StatusFileCheckedOut,
		StatusCheckoutRequired,
		StatusBadFileType,
		StatusFileTooLarge,
		StatusFormsAuthRequired,
		StatusVirusInfected,
		StatusVirusDeleted,
		StatusBadMcfgTable,
		StatusCannotBreakOplock,
		StatusWowAssertion,
		StatusInvalidSignature,
		StatusHmacNotSupported,
		StatusAuthTagMismatch,
		StatusIpsecQueueOverflow,
		StatusNdQueueOverflow,
		StatusHoplimitExceeded,
		StatusProtocolNotSupported,
		StatusFastpathRejected,
		StatusLostWritebehindDataNetworkDisconnected,
		StatusLostWritebehindDataNetworkServerError,
		StatusLostWritebehindDataLocalDiskError,
		StatusXmlParseError,
		StatusXmldsigError,
		StatusWrongCompartment,
		StatusAuthipFailure,
		StatusDsOidMappedGroupCantHaveMembers,
		StatusDsOidNotFound,
		StatusHashNotSupported,
		StatusHashNotPresent,
		DbgNoStateChange,
		DbgAppNotIdle,
		RpcNtInvalidStringBinding,
		RpcNtWrongKindOfBinding,
		RpcNtInvalidBinding,
		RpcNtProtseqNotSupported,
		RpcNtInvalidRpcProtseq,
		RpcNtInvalidStringUuid,
		RpcNtInvalidEndpointFormat,
		RpcNtInvalidNetAddr,
		RpcNtNoEndpointFound,
		RpcNtInvalidTimeout,
		RpcNtObjectNotFound,
		RpcNtAlreadyRegistered,
		RpcNtTypeAlreadyRegistered,
		RpcNtAlreadyListening,
		RpcNtNoProtseqsRegistered,
		RpcNtNotListening,
		RpcNtUnknownMgrType,
		RpcNtUnknownIf,
		RpcNtNoBindings,
		RpcNtNoProtseqs,
		RpcNtCantCreateEndpoint,
		RpcNtOutOfResources,
		RpcNtServerUnavailable,
		RpcNtServerTooBusy,
		RpcNtInvalidNetworkOptions,
		RpcNtNoCallActive,
		RpcNtCallFailed,
		RpcNtCallFailedDne,
		RpcNtProtocolError,
		RpcNtUnsupportedTransSyn,
		RpcNtUnsupportedType,
		RpcNtInvalidTag,
		RpcNtInvalidBound,
		RpcNtNoEntryName,
		RpcNtInvalidNameSyntax,
		RpcNtUnsupportedNameSyntax,
		RpcNtUuidNoAddress,
		RpcNtDuplicateEndpoint,
		RpcNtUnknownAuthnType,
		RpcNtMaxCallsTooSmall,
		RpcNtStringTooLong,
		RpcNtProtseqNotFound,
		RpcNtProcnumOutOfRange,
		RpcNtBindingHasNoAuth,
		RpcNtUnknownAuthnService,
		RpcNtUnknownAuthnLevel,
		RpcNtInvalidAuthIdentity,
		RpcNtUnknownAuthzService,
		EptNtInvalidEntry,
		EptNtCantPerformOp,
		EptNtNotRegistered,
		RpcNtNothingToExport,
		RpcNtIncompleteName,
		RpcNtInvalidVersOption,
		RpcNtNoMoreMembers,
		RpcNtNotAllObjsUnexported,
		RpcNtInterfaceNotFound,
		RpcNtEntryAlreadyExists,
		RpcNtEntryNotFound,
		RpcNtNameServiceUnavailable,
		RpcNtInvalidNafId,
		RpcNtCannotSupport,
		RpcNtNoContextAvailable,
		RpcNtInternalError,
		RpcNtZeroDivide,
		RpcNtAddressError,
		RpcNtFpDivZero,
		RpcNtFpUnderflow,
		RpcNtFpOverflow,
		RpcNtNoMoreEntries,
		RpcNtSsCharTransOpenFail,
		RpcNtSsCharTransShortFile,
		RpcNtSsInNullContext,
		RpcNtSsContextMismatch,
		RpcNtSsContextDamaged,
		RpcNtSsHandlesMismatch,
		RpcNtSsCannotGetCallHandle,
		RpcNtNullRefPointer,
		RpcNtEnumValueOutOfRange,
		RpcNtByteCountTooSmall,
		RpcNtBadStubData,
		RpcNtCallInProgress,
		RpcNtNoMoreBindings,
		RpcNtGroupMemberNotFound,
		EptNtCantCreate,
		RpcNtInvalidObject,
		RpcNtNoInterfaces,
		RpcNtCallCancelled,
		RpcNtBindingIncomplete,
		RpcNtCommFailure,
		RpcNtUnsupportedAuthnLevel,
		RpcNtNoPrincName,
		RpcNtNotRpcError,
		RpcNtUuidLocalOnly,
		RpcNtSecPkgError,
		RpcNtNotCancelled,
		RpcNtInvalidEsAction,
		RpcNtWrongEsVersion,
		RpcNtWrongStubVersion,
		RpcNtInvalidPipeObject,
		RpcNtInvalidPipeOperation,
		RpcNtWrongPipeVersion,
		RpcNtPipeClosed,
		RpcNtPipeDisciplineError,
		RpcNtPipeEmpty,
		RpcNtInvalidAsyncHandle,
		RpcNtInvalidAsyncCall,
		RpcNtProxyAccessDenied,
		RpcNtCookieAuthFailed,
		RpcNtSendIncomplete,
		StatusAcpiInvalidOpcode,
		StatusAcpiStackOverflow,
		StatusAcpiAssertFailed,
		StatusAcpiInvalidIndex,
		StatusAcpiInvalidArgument,
		StatusAcpiFatal,
		StatusAcpiInvalidSupername,
		StatusAcpiInvalidArgtype,
		StatusAcpiInvalidObjtype,
		StatusAcpiInvalidTargettype,
		StatusAcpiIncorrectArgumentCount,
		StatusAcpiAddressNotMapped,
		StatusAcpiInvalidEventtype,
		StatusAcpiHandlerCollision,
		StatusAcpiInvalidData,
		StatusAcpiInvalidRegion,
		StatusAcpiInvalidAccessSize,
		StatusAcpiAcquireGlobalLock,
		StatusAcpiAlreadyInitialized,
		StatusAcpiNotInitialized,
		StatusAcpiInvalidMutexLevel,
		StatusAcpiMutexNotOwned,
		StatusAcpiMutexNotOwner,
		StatusAcpiRsAccess,
		StatusAcpiInvalidTable,
		StatusAcpiRegHandlerFailed,
		StatusAcpiPowerRequestFailed,
		StatusCtxWinstationNameInvalid,
		StatusCtxInvalidPd,
		StatusCtxPdNotFound,
		StatusCtxCdmConnect,
		StatusCtxCdmDisconnect,
		StatusCtxClosePending,
		StatusCtxNoOutbuf,
		StatusCtxModemInfNotFound,
		StatusCtxInvalidModemname,
		StatusCtxResponseError,
		StatusCtxModemResponseTimeout,
		StatusCtxModemResponseNoCarrier,
		StatusCtxModemResponseNoDialtone,
		StatusCtxModemResponseBusy,
		StatusCtxModemResponseVoice,
		StatusCtxTdError,
		StatusCtxLicenseClientInvalid,
		StatusCtxLicenseNotAvailable,
		StatusCtxLicenseExpired,
		StatusCtxWinstationNotFound,
		StatusCtxWinstationNameCollision,
		StatusCtxWinstationBusy,
		StatusCtxBadVideoMode,
		StatusCtxGraphicsInvalid,
		StatusCtxNotConsole,
		StatusCtxClientQueryTimeout,
		StatusCtxConsoleDisconnect,
		StatusCtxConsoleConnect,
		StatusCtxShadowDenied,
		StatusCtxWinstationAccessDenied,
		StatusCtxInvalidWd,
		StatusCtxWdNotFound,
		StatusCtxShadowInvalid,
		StatusCtxShadowDisabled,
		StatusRdpProtocolError,
		StatusCtxClientLicenseNotSet,
		StatusCtxClientLicenseInUse,
		StatusCtxShadowEndedByModeChange,
		StatusCtxShadowNotRunning,
		StatusCtxLogonDisabled,
		StatusCtxSecurityLayerError,
		StatusTsIncompatibleSessions,
		StatusTsVideoSubsystemError,
		StatusPnpBadMpsTable,
		StatusPnpTranslationFailed,
		StatusPnpIrqTranslationFailed,
		StatusPnpInvalidId,
		StatusIoReissueAsCached,
		StatusMuiFileNotFound,
		StatusMuiInvalidFile,
		StatusMuiInvalidRcConfig,
		StatusMuiInvalidLocaleName,
		StatusMuiInvalidUltimatefallbackName,
		StatusMuiFileNotLoaded,
		StatusResourceEnumUserStop,
		StatusFltNoHandlerDefined,
		StatusFltContextAlreadyDefined,
		StatusFltInvalidAsynchronousRequest,
		StatusFltDisallowFastIo,
		StatusFltInvalidNameRequest,
		StatusFltNotSafeToPostOperation,
		StatusFltNotInitialized,
		StatusFltFilterNotReady,
		StatusFltPostOperationCleanup,
		StatusFltInternalError,
		StatusFltDeletingObject,
		StatusFltMustBeNonpagedPool,
		StatusFltDuplicateEntry,
		StatusFltCbdqDisabled,
		StatusFltDoNotAttach,
		StatusFltDoNotDetach,
		StatusFltInstanceAltitudeCollision,
		StatusFltInstanceNameCollision,
		StatusFltFilterNotFound,
		StatusFltVolumeNotFound,
		StatusFltInstanceNotFound,
		StatusFltContextAllocationNotFound,
		StatusFltInvalidContextRegistration,
		StatusFltNameCacheMiss,
		StatusFltNoDeviceObject,
		StatusFltVolumeAlreadyMounted,
		StatusFltAlreadyEnlisted,
		StatusFltContextAlreadyLinked,
		StatusFltNoWaiterForReply,
		StatusSxsSectionNotFound,
		StatusSxsCantGenActctx,
		StatusSxsInvalidActctxdataFormat,
		StatusSxsAssemblyNotFound,
		StatusSxsManifestFormatError,
		StatusSxsManifestParseError,
		StatusSxsActivationContextDisabled,
		StatusSxsKeyNotFound,
		StatusSxsVersionConflict,
		StatusSxsWrongSectionType,
		StatusSxsThreadQueriesDisabled,
		StatusSxsAssemblyMissing,
		StatusSxsReleaseActivationContext,
		StatusSxsProcessDefaultAlreadySet,
		StatusSxsEarlyDeactivation,
		StatusSxsInvalidDeactivation,
		StatusSxsMultipleDeactivation,
		StatusSxsSystemDefaultActivationContextEmpty,
		StatusSxsProcessTerminationRequested,
		StatusSxsCorruptActivationStack,
		StatusSxsCorruption,
		StatusSxsInvalidIdentityAttributeValue,
		StatusSxsInvalidIdentityAttributeName,
		StatusSxsIdentityDuplicateAttribute,
		StatusSxsIdentityParseError,
		StatusSxsComponentStoreCorrupt,
		StatusSxsFileHashMismatch,
		StatusSxsManifestIdentitySameButContentsDifferent,
		StatusSxsIdentitiesDifferent,
		StatusSxsAssemblyIsNotADeployment,
		StatusSxsFileNotPartOfAssembly,
		StatusAdvancedInstallerFailed,
		StatusXmlEncodingMismatch,
		StatusSxsManifestTooBig,
		StatusSxsSettingNotRegistered,
		StatusSxsTransactionClosureIncomplete,
		StatusSmiPrimitiveInstallerFailed,
		StatusGenericCommandFailed,
		StatusSxsFileHashMissing,
		StatusClusterInvalidNode,
		StatusClusterNodeExists,
		StatusClusterJoinInProgress,
		StatusClusterNodeNotFound,
		StatusClusterLocalNodeNotFound,
		StatusClusterNetworkExists,
		StatusClusterNetworkNotFound,
		StatusClusterNetinterfaceExists,
		StatusClusterNetinterfaceNotFound,
		StatusClusterInvalidRequest,
		StatusClusterInvalidNetworkProvider,
		StatusClusterNodeDown,
		StatusClusterNodeUnreachable,
		StatusClusterNodeNotMember,
		StatusClusterJoinNotInProgress,
		StatusClusterInvalidNetwork,
		StatusClusterNoNetAdapters,
		StatusClusterNodeUp,
		StatusClusterNodePaused,
		StatusClusterNodeNotPaused,
		StatusClusterNoSecurityContext,
		StatusClusterNetworkNotInternal,
		StatusClusterPoisoned,
		StatusClusterNonCsvPath,
		StatusClusterCsvVolumeNotLocal,
		StatusTransactionalConflict,
		StatusInvalidTransaction,
		StatusTransactionNotActive,
		StatusTmInitializationFailed,
		StatusRmNotActive,
		StatusRmMetadataCorrupt,
		StatusTransactionNotJoined,
		StatusDirectoryNotRm,
		StatusCouldNotResizeLog,
		StatusTransactionsUnsupportedRemote,
		StatusLogResizeInvalidSize,
		StatusRemoteFileVersionMismatch,
		StatusCrmProtocolAlreadyExists,
		StatusTransactionPropagationFailed,
		StatusCrmProtocolNotFound,
		StatusTransactionSuperiorExists,
		StatusTransactionRequestNotValid,
		StatusTransactionNotRequested,
		StatusTransactionAlreadyAborted,
		StatusTransactionAlreadyCommitted,
		StatusTransactionInvalidMarshallBuffer,
		StatusCurrentTransactionNotValid,
		StatusLogGrowthFailed,
		StatusObjectNoLongerExists,
		StatusStreamMiniversionNotFound,
		StatusStreamMiniversionNotValid,
		StatusMiniversionInaccessibleFromSpecifiedTransaction,
		StatusCantOpenMiniversionWithModifyIntent,
		StatusCantCreateMoreStreamMiniversions,
		StatusHandleNoLongerValid,
		StatusNoTxfMetadata,
		StatusLogCorruptionDetected,
		StatusCantRecoverWithHandleOpen,
		StatusRmDisconnected,
		StatusEnlistmentNotSuperior,
		StatusRecoveryNotNeeded,
		StatusRmAlreadyStarted,
		StatusFileIdentityNotPersistent,
		StatusCantBreakTransactionalDependency,
		StatusCantCrossRmBoundary,
		StatusTxfDirNotEmpty,
		StatusIndoubtTransactionsExist,
		StatusTmVolatile,
		StatusRollbackTimerExpired,
		StatusTxfAttributeCorrupt,
		StatusEfsNotAllowedInTransaction,
		StatusTransactionalOpenNotAllowed,
		StatusTransactedMappingUnsupportedRemote,
		StatusTxfMetadataAlreadyPresent,
		StatusTransactionScopeCallbacksNotSet,
		StatusTransactionRequiredPromotion,
		StatusCannotExecuteFileInTransaction,
		StatusTransactionsNotFrozen,
		StatusTransactionFreezeInProgress,
		StatusNotSnapshotVolume,
		StatusNoSavepointWithOpenFiles,
		StatusSparseNotAllowedInTransaction,
		StatusTmIdentityMismatch,
		StatusFloatedSection,
		StatusCannotAcceptTransactedWork,
		StatusCannotAbortTransactions,
		StatusTransactionNotFound,
		StatusResourcemanagerNotFound,
		StatusEnlistmentNotFound,
		StatusTransactionmanagerNotFound,
		StatusTransactionmanagerNotOnline,
		StatusTransactionmanagerRecoveryNameCollision,
		StatusTransactionNotRoot,
		StatusTransactionObjectExpired,
		StatusCompressionNotAllowedInTransaction,
		StatusTransactionResponseNotEnlisted,
		StatusTransactionRecordTooLong,
		StatusNoLinkTrackingInTransaction,
		StatusOperationNotSupportedInTransaction,
		StatusTransactionIntegrityViolated,
		StatusTransactionmanagerIdentityMismatch,
		StatusRmCannotBeFrozenForSnapshot,
		StatusTransactionMustWritethrough,
		StatusTransactionNoSuperior,
		StatusExpiredHandle,
		StatusTransactionNotEnlisted,
		StatusLogSectorInvalid,
		StatusLogSectorParityInvalid,
		StatusLogSectorRemapped,
		StatusLogBlockIncomplete,
		StatusLogInvalidRange,
		StatusLogBlocksExhausted,
		StatusLogReadContextInvalid,
		StatusLogRestartInvalid,
		StatusLogBlockVersion,
		StatusLogBlockInvalid,
		StatusLogReadModeInvalid,
		StatusLogNoRestart,
		StatusLogMetadataCorrupt,
		StatusLogMetadataInvalid,
		StatusLogMetadataInconsistent,
		StatusLogReservationInvalid,
		StatusLogCantDelete,
		StatusLogContainerLimitExceeded,
		StatusLogStartOfLog,
		StatusLogPolicyAlreadyInstalled,
		StatusLogPolicyNotInstalled,
		StatusLogPolicyInvalid,
		StatusLogPolicyConflict,
		StatusLogPinnedArchiveTail,
		StatusLogRecordNonexistent,
		StatusLogRecordsReservedInvalid,
		StatusLogSpaceReservedInvalid,
		StatusLogTailInvalid,
		StatusLogFull,
		StatusLogMultiplexed,
		StatusLogDedicated,
		StatusLogArchiveNotInProgress,
		StatusLogArchiveInProgress,
		StatusLogEphemeral,
		StatusLogNotEnoughContainers,
		StatusLogClientAlreadyRegistered,
		StatusLogClientNotRegistered,
		StatusLogFullHandlerInProgress,
		StatusLogContainerReadFailed,
		StatusLogContainerWriteFailed,
		StatusLogContainerOpenFailed,
		StatusLogContainerStateInvalid,
		StatusLogStateInvalid,
		StatusLogPinned,
		StatusLogMetadataFlushFailed,
		StatusLogInconsistentSecurity,
		StatusLogAppendedFlushFailed,
		StatusLogPinnedReservation,
		StatusVideoHungDisplayDriverThread,
		StatusVideoHungDisplayDriverThreadRecovered,
		StatusVideoDriverDebugReportRequest,
		StatusMonitorNoDescriptor,
		StatusMonitorUnknownDescriptorFormat,
		StatusMonitorInvalidDescriptorChecksum,
		StatusMonitorInvalidStandardTimingBlock,
		StatusMonitorWmiDatablockRegistrationFailed,
		StatusMonitorInvalidSerialNumberMondscBlock,
		StatusMonitorInvalidUserFriendlyMondscBlock,
		StatusMonitorNoMoreDescriptorData,
		StatusMonitorInvalidDetailedTimingBlock,
		StatusMonitorInvalidManufactureDate,
		StatusGraphicsNotExclusiveModeOwner,
		StatusGraphicsInsufficientDmaBuffer,
		StatusGraphicsInvalidDisplayAdapter,
		StatusGraphicsAdapterWasReset,
		StatusGraphicsInvalidDriverModel,
		StatusGraphicsPresentModeChanged,
		StatusGraphicsPresentOccluded,
		StatusGraphicsPresentDenied,
		StatusGraphicsCannotcolorconvert,
		StatusGraphicsDriverMismatch,
		StatusGraphicsPartialDataPopulated,
		StatusGraphicsPresentRedirectionDisabled,
		StatusGraphicsPresentUnoccluded,
		StatusGraphicsNoVideoMemory,
		StatusGraphicsCantLockMemory,
		StatusGraphicsAllocationBusy,
		StatusGraphicsTooManyReferences,
		StatusGraphicsTryAgainLater,
		StatusGraphicsTryAgainNow,
		StatusGraphicsAllocationInvalid,
		StatusGraphicsUnswizzlingApertureUnavailable,
		StatusGraphicsUnswizzlingApertureUnsupported,
		StatusGraphicsCantEvictPinnedAllocation,
		StatusGraphicsInvalidAllocationUsage,
		StatusGraphicsCantRenderLockedAllocation,
		StatusGraphicsAllocationClosed,
		StatusGraphicsInvalidAllocationInstance,
		StatusGraphicsInvalidAllocationHandle,
		StatusGraphicsWrongAllocationDevice,
		StatusGraphicsAllocationContentLost,
		StatusGraphicsGpuExceptionOnDevice,
		StatusGraphicsInvalidVidpnTopology,
		StatusGraphicsVidpnTopologyNotSupported,
		StatusGraphicsVidpnTopologyCurrentlyNotSupported,
		StatusGraphicsInvalidVidpn,
		StatusGraphicsInvalidVideoPresentSource,
		StatusGraphicsInvalidVideoPresentTarget,
		StatusGraphicsVidpnModalityNotSupported,
		StatusGraphicsModeNotPinned,
		StatusGraphicsInvalidVidpnSourcemodeset,
		StatusGraphicsInvalidVidpnTargetmodeset,
		StatusGraphicsInvalidFrequency,
		StatusGraphicsInvalidActiveRegion,
		StatusGraphicsInvalidTotalRegion,
		StatusGraphicsInvalidVideoPresentSourceMode,
		StatusGraphicsInvalidVideoPresentTargetMode,
		StatusGraphicsPinnedModeMustRemainInSet,
		StatusGraphicsPathAlreadyInTopology,
		StatusGraphicsModeAlreadyInModeset,
		StatusGraphicsInvalidVideopresentsourceset,
		StatusGraphicsInvalidVideopresenttargetset,
		StatusGraphicsSourceAlreadyInSet,
		StatusGraphicsTargetAlreadyInSet,
		StatusGraphicsInvalidVidpnPresentPath,
		StatusGraphicsNoRecommendedVidpnTopology,
		StatusGraphicsInvalidMonitorFrequencyrangeset,
		StatusGraphicsInvalidMonitorFrequencyrange,
		StatusGraphicsFrequencyrangeNotInSet,
		StatusGraphicsNoPreferredMode,
		StatusGraphicsFrequencyrangeAlreadyInSet,
		StatusGraphicsStaleModeset,
		StatusGraphicsInvalidMonitorSourcemodeset,
		StatusGraphicsInvalidMonitorSourceMode,
		StatusGraphicsNoRecommendedFunctionalVidpn,
		StatusGraphicsModeIdMustBeUnique,
		StatusGraphicsEmptyAdapterMonitorModeSupportIntersection,
		StatusGraphicsVideoPresentTargetsLessThanSources,
		StatusGraphicsPathNotInTopology,
		StatusGraphicsAdapterMustHaveAtLeastOneSource,
		StatusGraphicsAdapterMustHaveAtLeastOneTarget,
		StatusGraphicsInvalidMonitordescriptorset,
		StatusGraphicsInvalidMonitordescriptor,
		StatusGraphicsMonitordescriptorNotInSet,
		StatusGraphicsMonitordescriptorAlreadyInSet,
		StatusGraphicsMonitordescriptorIdMustBeUnique,
		StatusGraphicsInvalidVidpnTargetSubsetType,
		StatusGraphicsResourcesNotRelated,
		StatusGraphicsSourceIdMustBeUnique,
		StatusGraphicsTargetIdMustBeUnique,
		StatusGraphicsNoAvailableVidpnTarget,
		StatusGraphicsMonitorCouldNotBeAssociatedWithAdapter,
		StatusGraphicsNoVidpnmgr,
		StatusGraphicsNoActiveVidpn,
		StatusGraphicsStaleVidpnTopology,
		StatusGraphicsMonitorNotConnected,
		StatusGraphicsSourceNotInTopology,
		StatusGraphicsInvalidPrimarysurfaceSize,
		StatusGraphicsInvalidVisibleregionSize,
		StatusGraphicsInvalidStride,
		StatusGraphicsInvalidPixelformat,
		StatusGraphicsInvalidColorbasis,
		StatusGraphicsInvalidPixelvalueaccessmode,
		StatusGraphicsTargetNotInTopology,
		StatusGraphicsNoDisplayModeManagementSupport,
		StatusGraphicsVidpnSourceInUse,
		StatusGraphicsCantAccessActiveVidpn,
		StatusGraphicsInvalidPathImportanceOrdinal,
		StatusGraphicsInvalidPathContentGeometryTransformation,
		StatusGraphicsPathContentGeometryTransformationNotSupported,
		StatusGraphicsInvalidGammaRamp,
		StatusGraphicsGammaRampNotSupported,
		StatusGraphicsMultisamplingNotSupported,
		StatusGraphicsModeNotInModeset,
		StatusGraphicsDatasetIsEmpty,
		StatusGraphicsNoMoreElementsInDataset,
		StatusGraphicsInvalidVidpnTopologyRecommendationReason,
		StatusGraphicsInvalidPathContentType,
		StatusGraphicsInvalidCopyprotectionType,
		StatusGraphicsUnassignedModesetAlreadyExists,
		StatusGraphicsPathContentGeometryTransformationNotPinned,
		StatusGraphicsInvalidScanlineOrdering,
		StatusGraphicsTopologyChangesNotAllowed,
		StatusGraphicsNoAvailableImportanceOrdinals,
		StatusGraphicsIncompatiblePrivateFormat,
		StatusGraphicsInvalidModePruningAlgorithm,
		StatusGraphicsInvalidMonitorCapabilityOrigin,
		StatusGraphicsInvalidMonitorFrequencyrangeConstraint,
		StatusGraphicsMaxNumPathsReached,
		StatusGraphicsCancelVidpnTopologyAugmentation,
		StatusGraphicsInvalidClientType,
		StatusGraphicsClientvidpnNotSet,
		StatusGraphicsSpecifiedChildAlreadyConnected,
		StatusGraphicsChildDescriptorNotSupported,
		StatusGraphicsUnknownChildStatus,
		StatusGraphicsNotALinkedAdapter,
		StatusGraphicsLeadlinkNotEnumerated,
		StatusGraphicsChainlinksNotEnumerated,
		StatusGraphicsAdapterChainNotReady,
		StatusGraphicsChainlinksNotStarted,
		StatusGraphicsChainlinksNotPoweredOn,
		StatusGraphicsInconsistentDeviceLinkState,
		StatusGraphicsLeadlinkStartDeferred,
		StatusGraphicsNotPostDeviceDriver,
		StatusGraphicsPollingTooFrequently,
		StatusGraphicsStartDeferred,
		StatusGraphicsAdapterAccessNotExcluded,
		StatusGraphicsOpmNotSupported,
		StatusGraphicsCoppNotSupported,
		StatusGraphicsUabNotSupported,
		StatusGraphicsOpmInvalidEncryptedParameters,
		StatusGraphicsOpmNoProtectedOutputsExist,
		StatusGraphicsOpmInternalError,
		StatusGraphicsOpmInvalidHandle,
		StatusGraphicsPvpInvalidCertificateLength,
		StatusGraphicsOpmSpanningModeEnabled,
		StatusGraphicsOpmTheaterModeEnabled,
		StatusGraphicsPvpHfsFailed,
		StatusGraphicsOpmInvalidSrm,
		StatusGraphicsOpmOutputDoesNotSupportHdcp,
		StatusGraphicsOpmOutputDoesNotSupportAcp,
		StatusGraphicsOpmOutputDoesNotSupportCgmsa,
		StatusGraphicsOpmHdcpSrmNeverSet,
		StatusGraphicsOpmResolutionTooHigh,
		StatusGraphicsOpmAllHdcpHardwareAlreadyInUse,
		StatusGraphicsOpmProtectedOutputNoLongerExists,
		StatusGraphicsOpmProtectedOutputDoesNotHaveCoppSemantics,
		StatusGraphicsOpmInvalidInformationRequest,
		StatusGraphicsOpmDriverInternalError,
		StatusGraphicsOpmProtectedOutputDoesNotHaveOpmSemantics,
		StatusGraphicsOpmSignalingNotSupported,
		StatusGraphicsOpmInvalidConfigurationRequest,
		StatusGraphicsI2CNotSupported,
		StatusGraphicsI2CDeviceDoesNotExist,
		StatusGraphicsI2CErrorTransmittingData,
		StatusGraphicsI2CErrorReceivingData,
		StatusGraphicsDdcciVcpNotSupported,
		StatusGraphicsDdcciInvalidData,
		StatusGraphicsDdcciMonitorReturnedInvalidTimingStatusByte,
		StatusGraphicsDdcciInvalidCapabilitiesString,
		StatusGraphicsMcaInternalError,
		StatusGraphicsDdcciInvalidMessageCommand,
		StatusGraphicsDdcciInvalidMessageLength,
		StatusGraphicsDdcciInvalidMessageChecksum,
		StatusGraphicsInvalidPhysicalMonitorHandle,
		StatusGraphicsMonitorNoLongerExists,
		StatusGraphicsOnlyConsoleSessionSupported,
		StatusGraphicsNoDisplayDeviceCorrespondsToName,
		StatusGraphicsDisplayDeviceNotAttachedToDesktop,
		StatusGraphicsMirroringDevicesNotSupported,
		StatusGraphicsInvalidPointer,
		StatusGraphicsNoMonitorsCorrespondToDisplayDevice,
		StatusGraphicsParameterArrayTooSmall,
		StatusGraphicsInternalError,
		StatusGraphicsSessionTypeChangeInProgress,
		StatusFveLockedVolume,
		StatusFveNotEncrypted,
		StatusFveBadInformation,
		StatusFveTooSmall,
		StatusFveFailedWrongFs,
		StatusFveBadPartitionSize,
		StatusFveFsNotExtended,
		StatusFveFsMounted,
		StatusFveNoLicense,
		StatusFveActionNotAllowed,
		StatusFveBadData,
		StatusFveVolumeNotBound,
		StatusFveNotDataVolume,
		StatusFveConvReadError,
		StatusFveConvWriteError,
		StatusFveOverlappedUpdate,
		StatusFveFailedSectorSize,
		StatusFveFailedAuthentication,
		StatusFveNotOsVolume,
		StatusFveKeyfileNotFound,
		StatusFveKeyfileInvalid,
		StatusFveKeyfileNoVmk,
		StatusFveTpmDisabled,
		StatusFveTpmSrkAuthNotZero,
		StatusFveTpmInvalidPcr,
		StatusFveTpmNoVmk,
		StatusFvePinInvalid,
		StatusFveAuthInvalidApplication,
		StatusFveAuthInvalidConfig,
		StatusFveDebuggerEnabled,
		StatusFveDryRunFailed,
		StatusFveBadMetadataPointer,
		StatusFveOldMetadataCopy,
		StatusFveRebootRequired,
		StatusFveRawAccess,
		StatusFveRawBlocked,
		StatusFveNoAutounlockMasterKey,
		StatusFveMorFailed,
		StatusFveNoFeatureLicense,
		StatusFvePolicyUserDisableRdvNotAllowed,
		StatusFveConvRecoveryFailed,
		StatusFveVirtualizedSpaceTooBig,
		StatusFveInvalidDatumType,
		StatusFveVolumeTooSmall,
		StatusFveEnhPinInvalid,
		StatusFwpCalloutNotFound,
		StatusFwpConditionNotFound,
		StatusFwpFilterNotFound,
		StatusFwpLayerNotFound,
		StatusFwpProviderNotFound,
		StatusFwpProviderContextNotFound,
		StatusFwpSublayerNotFound,
		StatusFwpNotFound,
		StatusFwpAlreadyExists,
		StatusFwpInUse,
		StatusFwpDynamicSessionInProgress,
		StatusFwpWrongSession,
		StatusFwpNoTxnInProgress,
		StatusFwpTxnInProgress,
		StatusFwpTxnAborted,
		StatusFwpSessionAborted,
		StatusFwpIncompatibleTxn,
		StatusFwpTimeout,
		StatusFwpNetEventsDisabled,
		StatusFwpIncompatibleLayer,
		StatusFwpKmClientsOnly,
		StatusFwpLifetimeMismatch,
		StatusFwpBuiltinObject,
		StatusFwpTooManyCallouts,
		StatusFwpNotificationDropped,
		StatusFwpTrafficMismatch,
		StatusFwpIncompatibleSaState,
		StatusFwpNullPointer,
		StatusFwpInvalidEnumerator,
		StatusFwpInvalidFlags,
		StatusFwpInvalidNetMask,
		StatusFwpInvalidRange,
		StatusFwpInvalidInterval,
		StatusFwpZeroLengthArray,
		StatusFwpNullDisplayName,
		StatusFwpInvalidActionType,
		StatusFwpInvalidWeight,
		StatusFwpMatchTypeMismatch,
		StatusFwpTypeMismatch,
		StatusFwpOutOfBounds,
		StatusFwpReserved,
		StatusFwpDuplicateCondition,
		StatusFwpDuplicateKeymod,
		StatusFwpActionIncompatibleWithLayer,
		StatusFwpActionIncompatibleWithSublayer,
		StatusFwpContextIncompatibleWithLayer,
		StatusFwpContextIncompatibleWithCallout,
		StatusFwpIncompatibleAuthMethod,
		StatusFwpIncompatibleDhGroup,
		StatusFwpEmNotSupported,
		StatusFwpNeverMatch,
		StatusFwpProviderContextMismatch,
		StatusFwpInvalidParameter,
		StatusFwpTooManySublayers,
		StatusFwpCalloutNotificationFailed,
		StatusFwpInvalidAuthTransform,
		StatusFwpInvalidCipherTransform,
		StatusFwpIncompatibleCipherTransform,
		StatusFwpInvalidTransformCombination,
		StatusFwpDuplicateAuthMethod,
		StatusFwpTcpipNotReady,
		StatusFwpInjectHandleClosing,
		StatusFwpInjectHandleStale,
		StatusFwpCannotPend,
		StatusFwpDropNoicmp,
		StatusNdisClosing,
		StatusNdisBadVersion,
		StatusNdisBadCharacteristics,
		StatusNdisAdapterNotFound,
		StatusNdisOpenFailed,
		StatusNdisDeviceFailed,
		StatusNdisMulticastFull,
		StatusNdisMulticastExists,
		StatusNdisMulticastNotFound,
		StatusNdisRequestAborted,
		StatusNdisResetInProgress,
		StatusNdisNotSupported,
		StatusNdisInvalidPacket,
		StatusNdisAdapterNotReady,
		StatusNdisInvalidLength,
		StatusNdisInvalidData,
		StatusNdisBufferTooShort,
		StatusNdisInvalidOid,
		StatusNdisAdapterRemoved,
		StatusNdisUnsupportedMedia,
		StatusNdisGroupAddressInUse,
		StatusNdisFileNotFound,
		StatusNdisErrorReadingFile,
		StatusNdisAlreadyMapped,
		StatusNdisResourceConflict,
		StatusNdisMediaDisconnected,
		StatusNdisInvalidAddress,
		StatusNdisInvalidDeviceRequest,
		StatusNdisPaused,
		StatusNdisInterfaceNotFound,
		StatusNdisUnsupportedRevision,
		StatusNdisInvalidPort,
		StatusNdisInvalidPortState,
		StatusNdisLowPowerState,
		StatusNdisDot11AutoConfigEnabled,
		StatusNdisDot11MediaInUse,
		StatusNdisDot11PowerStateInvalid,
		StatusNdisPmWolPatternListFull,
		StatusNdisPmProtocolOffloadListFull,
		StatusNdisIndicationRequired,
		StatusNdisOffloadPolicy,
		StatusNdisOffloadConnectionRejected,
		StatusNdisOffloadPathRejected,
		StatusHvInvalidHypercallCode,
		StatusHvInvalidHypercallInput,
		StatusHvInvalidAlignment,
		StatusHvInvalidParameter,
		StatusHvAccessDenied,
		StatusHvInvalidPartitionState,
		StatusHvOperationDenied,
		StatusHvUnknownProperty,
		StatusHvPropertyValueOutOfRange,
		StatusHvInsufficientMemory,
		StatusHvPartitionTooDeep,
		StatusHvInvalidPartitionId,
		StatusHvInvalidVpIndex,
		StatusHvInvalidPortId,
		StatusHvInvalidConnectionId,
		StatusHvInsufficientBuffers,
		StatusHvNotAcknowledged,
		StatusHvAcknowledged,
		StatusHvInvalidSaveRestoreState,
		StatusHvInvalidSynicState,
		StatusHvObjectInUse,
		StatusHvInvalidProximityDomainInfo,
		StatusHvNoData,
		StatusHvInactive,
		StatusHvNoResources,
		StatusHvFeatureUnavailable,
		StatusHvNotPresent,
		StatusVidDuplicateHandler,
		StatusVidTooManyHandlers,
		StatusVidQueueFull,
		StatusVidHandlerNotPresent,
		StatusVidInvalidObjectName,
		StatusVidPartitionNameTooLong,
		StatusVidMessageQueueNameTooLong,
		StatusVidPartitionAlreadyExists,
		StatusVidPartitionDoesNotExist,
		StatusVidPartitionNameNotFound,
		StatusVidMessageQueueAlreadyExists,
		StatusVidExceededMbpEntryMapLimit,
		StatusVidMbStillReferenced,
		StatusVidChildGpaPageSetCorrupted,
		StatusVidInvalidNumaSettings,
		StatusVidInvalidNumaNodeIndex,
		StatusVidNotificationQueueAlreadyAssociated,
		StatusVidInvalidMemoryBlockHandle,
		StatusVidPageRangeOverflow,
		StatusVidInvalidMessageQueueHandle,
		StatusVidInvalidGpaRangeHandle,
		StatusVidNoMemoryBlockNotificationQueue,
		StatusVidMemoryBlockLockCountExceeded,
		StatusVidInvalidPpmHandle,
		StatusVidMbpsAreLocked,
		StatusVidMessageQueueClosed,
		StatusVidVirtualProcessorLimitExceeded,
		StatusVidStopPending,
		StatusVidInvalidProcessorState,
		StatusVidExceededKmContextCountLimit,
		StatusVidKmInterfaceAlreadyInitialized,
		StatusVidMbPropertyAlreadySetReset,
		StatusVidMmioRangeDestroyed,
		StatusVidInvalidChildGpaPageSet,
		StatusVidReservePageSetIsBeingUsed,
		StatusVidReservePageSetTooSmall,
		StatusVidMbpAlreadyLockedUsingReservedPage,
		StatusVidMbpCountExceededLimit,
		StatusVidSavedStateCorrupt,
		StatusVidSavedStateUnrecognizedItem,
		StatusVidSavedStateIncompatible,
		StatusVidRemoteNodeParentGpaPagesUsed,
		StatusIpsecBadSpi,
		StatusIpsecSaLifetimeExpired,
		StatusIpsecWrongSa,
		StatusIpsecReplayCheckFailed,
		StatusIpsecInvalidPacket,
		StatusIpsecIntegrityCheckFailed,
		StatusIpsecClearTextDrop,
		StatusIpsecAuthFirewallDrop,
		StatusIpsecThrottleDrop,
		StatusIpsecDospBlock,
		StatusIpsecDospReceivedMulticast,
		StatusIpsecDospInvalidPacket,
		StatusIpsecDospStateLookupFailed,
		StatusIpsecDospMaxEntries,
		StatusIpsecDospKeymodNotAllowed,
		StatusIpsecDospMaxPerIpRatelimitQueues,
		StatusVolmgrIncompleteRegeneration,
		StatusVolmgrIncompleteDiskMigration,
		StatusVolmgrDatabaseFull,
		StatusVolmgrDiskConfigurationCorrupted,
		StatusVolmgrDiskConfigurationNotInSync,
		StatusVolmgrPackConfigUpdateFailed,
		StatusVolmgrDiskContainsNonSimpleVolume,
		StatusVolmgrDiskDuplicate,
		StatusVolmgrDiskDynamic,
		StatusVolmgrDiskIdInvalid,
		StatusVolmgrDiskInvalid,
		StatusVolmgrDiskLastVoter,
		StatusVolmgrDiskLayoutInvalid,
		StatusVolmgrDiskLayoutNonBasicBetweenBasicPartitions,
		StatusVolmgrDiskLayoutNotCylinderAligned,
		StatusVolmgrDiskLayoutPartitionsTooSmall,
		StatusVolmgrDiskLayoutPrimaryBetweenLogicalPartitions,
		StatusVolmgrDiskLayoutTooManyPartitions,
		StatusVolmgrDiskMissing,
		StatusVolmgrDiskNotEmpty,
		StatusVolmgrDiskNotEnoughSpace,
		StatusVolmgrDiskRevectoringFailed,
		StatusVolmgrDiskSectorSizeInvalid,
		StatusVolmgrDiskSetNotContained,
		StatusVolmgrDiskUsedByMultipleMembers,
		StatusVolmgrDiskUsedByMultiplePlexes,
		StatusVolmgrDynamicDiskNotSupported,
		StatusVolmgrExtentAlreadyUsed,
		StatusVolmgrExtentNotContiguous,
		StatusVolmgrExtentNotInPublicRegion,
		StatusVolmgrExtentNotSectorAligned,
		StatusVolmgrExtentOverlapsEbrPartition,
		StatusVolmgrExtentVolumeLengthsDoNotMatch,
		StatusVolmgrFaultTolerantNotSupported,
		StatusVolmgrInterleaveLengthInvalid,
		StatusVolmgrMaximumRegisteredUsers,
		StatusVolmgrMemberInSync,
		StatusVolmgrMemberIndexDuplicate,
		StatusVolmgrMemberIndexInvalid,
		StatusVolmgrMemberMissing,
		StatusVolmgrMemberNotDetached,
		StatusVolmgrMemberRegenerating,
		StatusVolmgrAllDisksFailed,
		StatusVolmgrNoRegisteredUsers,
		StatusVolmgrNoSuchUser,
		StatusVolmgrNotificationReset,
		StatusVolmgrNumberOfMembersInvalid,
		StatusVolmgrNumberOfPlexesInvalid,
		StatusVolmgrPackDuplicate,
		StatusVolmgrPackIdInvalid,
		StatusVolmgrPackInvalid,
		StatusVolmgrPackNameInvalid,
		StatusVolmgrPackOffline,
		StatusVolmgrPackHasQuorum,
		StatusVolmgrPackWithoutQuorum,
		StatusVolmgrPartitionStyleInvalid,
		StatusVolmgrPartitionUpdateFailed,
		StatusVolmgrPlexInSync,
		StatusVolmgrPlexIndexDuplicate,
		StatusVolmgrPlexIndexInvalid,
		StatusVolmgrPlexLastActive,
		StatusVolmgrPlexMissing,
		StatusVolmgrPlexRegenerating,
		StatusVolmgrPlexTypeInvalid,
		StatusVolmgrPlexNotRaid5,
		StatusVolmgrPlexNotSimple,
		StatusVolmgrStructureSizeInvalid,
		StatusVolmgrTooManyNotificationRequests,
		StatusVolmgrTransactionInProgress,
		StatusVolmgrUnexpectedDiskLayoutChange,
		StatusVolmgrVolumeContainsMissingDisk,
		StatusVolmgrVolumeIdInvalid,
		StatusVolmgrVolumeLengthInvalid,
		StatusVolmgrVolumeLengthNotSectorSizeMultiple,
		StatusVolmgrVolumeNotMirrored,
		StatusVolmgrVolumeNotRetained,
		StatusVolmgrVolumeOffline,
		StatusVolmgrVolumeRetained,
		StatusVolmgrNumberOfExtentsInvalid,
		StatusVolmgrDifferentSectorSize,
		StatusVolmgrBadBootDisk,
		StatusVolmgrPackConfigOffline,
		StatusVolmgrPackConfigOnline,
		StatusVolmgrNotPrimaryPack,
		StatusVolmgrPackLogUpdateFailed,
		StatusVolmgrNumberOfDisksInPlexInvalid,
		StatusVolmgrNumberOfDisksInMemberInvalid,
		StatusVolmgrVolumeMirrored,
		StatusVolmgrPlexNotSimpleSpanned,
		StatusVolmgrNoValidLogCopies,
		StatusVolmgrPrimaryPackPresent,
		StatusVolmgrNumberOfDisksInvalid,
		StatusVolmgrMirrorNotSupported,
		StatusVolmgrRaid5NotSupported,
		StatusBcdNotAllEntriesImported,
		StatusBcdTooManyElements,
		StatusBcdNotAllEntriesSynchronized,
		StatusVhdDriveFooterMissing,
		StatusVhdDriveFooterChecksumMismatch,
		StatusVhdDriveFooterCorrupt,
		StatusVhdFormatUnknown,
		StatusVhdFormatUnsupportedVersion,
		StatusVhdSparseHeaderChecksumMismatch,
		StatusVhdSparseHeaderUnsupportedVersion,
		StatusVhdSparseHeaderCorrupt,
		StatusVhdBlockAllocationFailure,
		StatusVhdBlockAllocationTableCorrupt,
		StatusVhdInvalidBlockSize,
		StatusVhdBitmapMismatch,
		StatusVhdParentVhdNotFound,
		StatusVhdChildParentIdMismatch,
		StatusVhdChildParentTimestampMismatch,
		StatusVhdMetadataReadFailure,
		StatusVhdMetadataWriteFailure,
		StatusVhdInvalidSize,
		StatusVhdInvalidFileSize,
		StatusVirtdiskProviderNotFound,
		StatusVirtdiskNotVirtualDisk,
		StatusVhdParentVhdAccessDenied,
		StatusVhdChildParentSizeMismatch,
		StatusVhdDifferencingChainCycleDetected,
		StatusVhdDifferencingChainErrorInParent,
		StatusVirtualDiskLimitation,
		StatusVhdInvalidType,
		StatusVhdInvalidState,
		StatusVirtdiskUnsupportedDiskSectorSize,
		StatusQueryStorageError,
		StatusDisNotPresent,
		StatusDisAttributeNotFound,
		StatusDisUnrecognizedAttribute,
		StatusDisPartialData,
	}
}
