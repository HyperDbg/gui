package constants

import "fmt"

type IoctlKind int

const (
	MonitorIoctlEnableMonitor                  IoctlKind = 0x00120004
	MonitorIoctlDisableMonitor                           = 0x00120008
	IoctlInternalKeyboardConnect                         = 0x000B0203
	IoctlInternalKeyboardDisconnect                      = 0x000B0403
	IoctlInternalKeyboardEnable                          = 0x000B0803
	IoctlInternalKeyboardDisable                         = 0x000B1003
	IoctlInternalMouseConnect                            = 0x000F0203
	IoctlInternalMouseDisconnect                         = 0x000F0403
	IoctlInternalMouseEnable                             = 0x000F0803
	IoctlInternalMouseDisable                            = 0x000F1003
	IoctlDoKernelmodeSamples                             = 0x00222000
	IoctlRegisterCallback                                = 0x00222004
	IoctlUnregisterCallback                              = 0x00222008
	IoctlGetCallbackVersion                              = 0x0022200C
	IoctlGetVersion                                      = 0x00222000
	IoctlReset                                           = 0x00222007
	IoctlReadwriteMailbox                                = 0x00222008
	IoctlMailboxWait                                     = 0x0022200C
	IoctlReadDma                                         = 0x00226012
	IoctlWriteDma                                        = 0x0022A015
	IoctlAcpiAsyncEvalMethod                             = 0x0032C000
	IoctlAcpiEvalMethod                                  = 0x0032C004
	IoctlAcpiAcquireGlobalLock                           = 0x0032C010
	IoctlAcpiReleaseGlobalLock                           = 0x0032C014
	IoctlAcpiEvalMethodEx                                = 0x0032C018
	IoctlAcpiAsyncEvalMethodEx                           = 0x0032C01C
	IoctlAcpiEnumChildren                                = 0x0032C020
	IoctlAvcUpdateVirtualSubunitInfo                     = 0x002A0000
	IoctlAvcRemoveVirtualSubunitInfo                     = 0x002A0004
	IoctlAvcBusReset                                     = 0x002A0008
	IoctlDot4CreateSocket                                = 0x003A2022
	IoctlDot4DestroySocket                               = 0x003A202A
	IoctlDot4WaitForChannel                              = 0x003A2026
	IoctlDot4OpenChannel                                 = 0x003A2006
	IoctlDot4CloseChannel                                = 0x003A2008
	IoctlDot4Read                                        = 0x003A200E
	IoctlDot4Write                                       = 0x003A2011
	IoctlDot4AddActivityBroadcast                        = 0x003A2014
	IoctlDot4RemoveActivityBroadcast                     = 0x003A2018
	IoctlDot4WaitActivityBroadcast                       = 0x003A201E
	IoctlMpdsmRegister                                   = 0x736DC004
	IoctlMpdsmDeregister                                 = 0x736DC008
	IoctlMountdevQueryUniqueId                           = 0x004D0000
	IoctlMountdevQuerySuggestedLinkName                  = 0x004D000C
	IoctlMountdevLinkCreated                             = 0x004DC010
	IoctlMountdevLinkDeleted                             = 0x004DC014
	IoctlMountdevQueryStableGuid                         = 0x004D0018
	IoctlMountmgrCreatePoint                             = 0x006DC000
	IoctlMountmgrDeletePoints                            = 0x006DC004
	IoctlMountmgrQueryPoints                             = 0x006D0008
	IoctlMountmgrDeletePointsDbonly                      = 0x006DC00C
	IoctlMountmgrNextDriveLetter                         = 0x006DC010
	IoctlMountmgrAutoDlAssignments                       = 0x006DC014
	IoctlMountmgrVolumeMountPointCreated                 = 0x006DC018
	IoctlMountmgrVolumeMountPointDeleted                 = 0x006DC01C
	IoctlMountmgrChangeNotify                            = 0x006D4020
	IoctlMountmgrKeepLinksWhenOffline                    = 0x006DC024
	IoctlMountmgrCheckUnprocessedVolumes                 = 0x006D4028
	IoctlMountmgrVolumeArrivalNotification               = 0x006D402C
	IoctlMountdevQueryDeviceName                         = 0x004D0008
	IoctlMountmgrQueryDosVolumePath                      = 0x006D0030
	IoctlMountmgrQueryDosVolumePaths                     = 0x006D0034
	IoctlMountmgrScrubRegistry                           = 0x006DC038
	IoctlMountmgrQueryAutoMount                          = 0x006D003C
	IoctlMountmgrSetAutoMount                            = 0x006DC040
	IoctlMountmgrBootDlAssignment                        = 0x006DC044
	IoctlMountmgrTracelogCache                           = 0x006D4048
	IoctlAvioAllocateStream                              = 0x00000000
	IoctlAvioFreeStream                                  = 0x00000000
	IoctlAvioModifyStream                                = 0x00000000
	NlbIoctlRegisterHook                                 = 0xC0C08048
	NlbPublicIoctlClientStickinessNotify                 = 0xC0C08054
	IoctlGetTupleData                                    = 0x00042EE0
	IoctlSocketInformation                               = 0x00042EF0
	IoctlPcmciaHideDevice                                = 0x0004AF08
	IoctlPcmciaRevealDevice                              = 0x0004AF0C
	IoctlSdSubmitRequest                                 = 0x00043073
	FsctlRequestOplockLevel1                             = 0x00090000
	FsctlRequestOplockLevel2                             = 0x00090004
	FsctlRequestBatchOplock                              = 0x00090008
	FsctlOplockBreakAcknowledge                          = 0x0009000C
	FsctlOpbatchAckClosePending                          = 0x00090010
	FsctlOplockBreakNotify                               = 0x00090014
	FsctlLockVolume                                      = 0x00090018
	FsctlUnlockVolume                                    = 0x0009001C
	FsctlDismountVolume                                  = 0x00090020
	FsctlIsVolumeMounted                                 = 0x00090028
	FsctlIsPathnameValid                                 = 0x0009002C
	FsctlMarkVolumeDirty                                 = 0x00090030
	FsctlQueryRetrievalPointers                          = 0x0009003B
	FsctlGetCompression                                  = 0x0009003C
	FsctlSetCompression                                  = 0x0009C040
	FsctlSetBootloaderAccessed                           = 0x0009004F
	FsctlOplockBreakAckNo2                               = 0x00090050
	FsctlInvalidateVolumes                               = 0x00090054
	FsctlQueryFatBpb                                     = 0x00090058
	FsctlRequestFilterOplock                             = 0x0009005C
	FsctlFilesystemGetStatistics                         = 0x00090060
	FsctlGetNtfsVolumeData                               = 0x00090064
	FsctlGetNtfsFileRecord                               = 0x00090068
	FsctlGetVolumeBitmap                                 = 0x0009006F
	FsctlGetRetrievalPointers                            = 0x00090073
	FsctlMoveFile                                        = 0x00090074
	FsctlIsVolumeDirty                                   = 0x00090078
	FsctlAllowExtendedDasdIo                             = 0x00090083
	FsctlFindFilesBySid                                  = 0x0009008F
	FsctlSetObjectId                                     = 0x00090098
	FsctlGetObjectId                                     = 0x0009009C
	FsctlDeleteObjectId                                  = 0x000900A0
	FsctlSetReparsePoint                                 = 0x000900A4
	FsctlGetReparsePoint                                 = 0x000900A8
	FsctlDeleteReparsePoint                              = 0x000900AC
	FsctlEnumUsnData                                     = 0x000900B3
	FsctlSecurityIdCheck                                 = 0x000940B7
	FsctlReadUsnJournal                                  = 0x000900BB
	FsctlSetObjectIdExtended                             = 0x000900BC
	FsctlCreateOrGetObjectId                             = 0x000900C0
	FsctlSetSparse                                       = 0x000900C4
	FsctlSetZeroData                                     = 0x000980C8
	FsctlQueryAllocatedRanges                            = 0x000940CF
	FsctlEnableUpgrade                                   = 0x000980D0
	FsctlSetEncryption                                   = 0x000900D7
	FsctlEncryptionFsctlIo                               = 0x000900DB
	FsctlWriteRawEncrypted                               = 0x000900DF
	FsctlReadRawEncrypted                                = 0x000900E3
	FsctlCreateUsnJournal                                = 0x000900E7
	FsctlReadFileUsnData                                 = 0x000900EB
	FsctlWriteUsnCloseRecord                             = 0x000900EF
	FsctlExtendVolume                                    = 0x000900F0
	FsctlQueryUsnJournal                                 = 0x000900F4
	FsctlDeleteUsnJournal                                = 0x000900F8
	FsctlMarkHandle                                      = 0x000900FC
	FsctlSisCopyfile                                     = 0x00090100
	FsctlSisLinkFiles                                    = 0x0009C104
	FsctlRecallFile                                      = 0x00090117
	FsctlReadFromPlex                                    = 0x0009411E
	FsctlFilePrefetch                                    = 0x00090120
	FsctlMakeMediaCompatible                             = 0x00098130
	FsctlSetDefectManagement                             = 0x00098134
	FsctlQuerySparingInfo                                = 0x00090138
	FsctlQueryOnDiskVolumeInfo                           = 0x0009013C
	FsctlSetVolumeCompressionState                       = 0x00090140
	FsctlTxfsModifyRm                                    = 0x00098144
	FsctlTxfsQueryRmInformation                          = 0x00094148
	FsctlTxfsRollforwardRedo                             = 0x00098150
	FsctlTxfsRollforwardUndo                             = 0x00098154
	FsctlTxfsStartRm                                     = 0x00098158
	FsctlTxfsShutdownRm                                  = 0x0009815C
	FsctlTxfsReadBackupInformation                       = 0x00094160
	FsctlTxfsWriteBackupInformation                      = 0x00098164
	FsctlTxfsCreateSecondaryRm                           = 0x00098168
	FsctlTxfsGetMetadataInfo                             = 0x0009416C
	FsctlTxfsGetTransactedVersion                        = 0x00094170
	FsctlTxfsSavepointInformation                        = 0x00098178
	FsctlTxfsCreateMiniversion                           = 0x0009817C
	FsctlTxfsTransactionActive                           = 0x0009418C
	FsctlSetZeroOnDeallocation                           = 0x00090194
	FsctlSetRepair                                       = 0x00090198
	FsctlGetRepair                                       = 0x0009019C
	FsctlWaitForRepair                                   = 0x000901A0
	FsctlInitiateRepair                                  = 0x000901A8
	FsctlCscInternal                                     = 0x000901AF
	FsctlShrinkVolume                                    = 0x000901B0
	FsctlSetShortNameBehavior                            = 0x000901B4
	FsctlDfsrSetGhostHandleState                         = 0x000901B8
	FsctlTxfsListTransactions                            = 0x000941E4
	FsctlQueryPagefileEncryption                         = 0x000901E8
	FsctlResetVolumeAllocationHints                      = 0x000901EC
	FsctlQueryDependentVolume                            = 0x000901F0
	FsctlSdGlobalChange                                  = 0x000901F4
	FsctlTxfsReadBackupInformation2                      = 0x000901F8
	FsctlLookupStreamFromCluster                         = 0x000901FC
	FsctlTxfsWriteBackupInformation2                     = 0x00090200
	FsctlFileTypeNotification                            = 0x00090204
	FsctlGetBootAreaInfo                                 = 0x00090230
	FsctlGetRetrievalPointerBase                         = 0x00090234
	FsctlSetPersistentVolumeState                        = 0x00090238
	FsctlQueryPersistentVolumeState                      = 0x0009023C
	FsctlRequestOplock                                   = 0x00090240
	FsctlCsvTunnelRequest                                = 0x00090244
	FsctlIsCsvFile                                       = 0x00090248
	FsctlQueryFileSystemRecognition                      = 0x0009024C
	FsctlCsvGetVolumePathName                            = 0x00090250
	FsctlCsvGetVolumeNameForVolumeMountPoint             = 0x00090254
	FsctlCsvGetVolumePathNamesForVolumeName              = 0x00090258
	FsctlIsFileOnCsvVolume                               = 0x0009025C
	FsctlLmrGetLinkTrackingInformation                   = 0x001400E8
	FsctlLmrSetLinkTrackingInformation                   = 0x001400EC
	IoctlLmrAreFileObjectsOnSameServer                   = 0x001400F0
	FsctlPipeAssignEvent                                 = 0x00110000
	FsctlPipeDisconnect                                  = 0x00110004
	FsctlPipeListen                                      = 0x00110008
	FsctlPipePeek                                        = 0x0011400C
	FsctlPipeQueryEvent                                  = 0x00110010
	FsctlPipeTransceive                                  = 0x0011C017
	FsctlPipeWait                                        = 0x00110018
	FsctlPipeImpersonate                                 = 0x0011001C
	FsctlPipeSetClientProcess                            = 0x00110020
	FsctlPipeQueryClientProcess                          = 0x00110024
	FsctlPipeGetPipeAttribute                            = 0x00110028
	FsctlPipeSetPipeAttribute                            = 0x0011002C
	FsctlPipeGetConnectionAttribute                      = 0x00110030
	FsctlPipeSetConnectionAttribute                      = 0x00110034
	FsctlPipeGetHandleAttribute                          = 0x00110038
	FsctlPipeSetHandleAttribute                          = 0x0011003C
	FsctlPipeFlush                                       = 0x00118040
	FsctlPipeInternalRead                                = 0x00115FF4
	FsctlPipeInternalWrite                               = 0x00119FF8
	FsctlPipeInternalTransceive                          = 0x0011DFFF
	FsctlPipeInternalReadOvflow                          = 0x00116000
	FsctlMailslotPeek                                    = 0x000C4003
	IoctlRedirQueryPath                                  = 0x0014018F
	IoctlRedirQueryPathEx                                = 0x00140193
	IoctlVolsnapFlushAndHoldWrites                       = 0x0053C000
	IoctlInternalParallelPortAllocate                    = 0x0016002C
	IoctlInternalGetParallelPortInfo                     = 0x00160030
	IoctlInternalParallelConnectInterrupt                = 0x00160034
	IoctlInternalParallelDisconnectInterrupt             = 0x00160038
	IoctlInternalReleaseParallelPortInfo                 = 0x0016003C
	IoctlInternalGetMoreParallelPortInfo                 = 0x00160044
	IoctlInternalParchipConnect                          = 0x00160048
	IoctlInternalParallelSetChipMode                     = 0x0016004C
	IoctlInternalParallelClearChipMode                   = 0x00160050
	IoctlInternalGetParallelPnpInfo                      = 0x00160054
	IoctlInternalInit12843Bus                            = 0x00160058
	IoctlInternalSelectDevice                            = 0x0016005C
	IoctlInternalDeselectDevice                          = 0x00160060
	IoctlInternalGetParportFdo                           = 0x00160074
	IoctlInternalParclassConnect                         = 0x00160078
	IoctlInternalParclassDisconnect                      = 0x0016007C
	IoctlInternalDisconnectIdle                          = 0x00160080
	IoctlInternalLockPort                                = 0x00160094
	IoctlInternalUnlockPort                              = 0x00160098
	IoctlInternalParallelPortFree                        = 0x001600A0
	IoctlInternalPardot3Connect                          = 0x001600A4
	IoctlInternalPardot3Disconnect                       = 0x001600A8
	IoctlInternalPardot3Reset                            = 0x001600AC
	IoctlInternalPardot3Signal                           = 0x001600B0
	IoctlInternalRegisterForRemovalRelations             = 0x001600C8
	IoctlInternalUnregisterForRemovalRelations           = 0x001600CC
	IoctlInternalLockPortNoSelect                        = 0x001600D0
	IoctlInternalUnlockPortNoDeselect                    = 0x001600D4
	IoctlInternalDisableEndOfChainBusRescan              = 0x001600D8
	IoctlInternalEnableEndOfChainBusRescan               = 0x001600DC
	IoctlWpdMessageReadwriteAccess                       = 0x0040C108
	IoctlWpdMessageReadAccess                            = 0x00404108
	IoctlScsiscanCmd                                     = 0x00190012
	IoctlScsiscanLockdevice                              = 0x00190016
	IoctlScsiscanUnlockdevice                            = 0x0019001A
	IoctlScsiscanSetTimeout                              = 0x0019001C
	IoctlScsiscanGetInfo                                 = 0x00190022
	IoctlSwenumInstallInterface                          = 0x002A0000
	IoctlSwenumRemoveInterface                           = 0x002A0004
	IoctlSwenumGetBusId                                  = 0x002A400B
	IoctlCancelIo                                        = 0x80002004
	IoctlWaitOnDeviceEvent                               = 0x80002008
	IoctlReadRegisters                                   = 0x8000200C
	IoctlWriteRegisters                                  = 0x80002010
	IoctlGetChannelAlignRqst                             = 0x80002014
	IoctlGetDeviceDescriptor                             = 0x80002018
	IoctlResetPipe                                       = 0x8000201C
	IoctlGetUsbDescriptor                                = 0x80002020
	IoctlSendUsbRequest                                  = 0x80002024
	IoctlGetPipeConfiguration                            = 0x80002028
	IoctlSetTimeout                                      = 0x8000202C
	IoctlEhstorDeviceSetAuthzState                       = 0x002D1404
	IoctlEhstorDeviceGetAuthzState                       = 0x002D1408
	IoctlEhstorDeviceSiloCommand                         = 0x002D140C
	IoctlEhstorDeviceEnumeratePdos                       = 0x002D1410
	IoctlKsProperty                                      = 0x002F0003
	IoctlKsEnableEvent                                   = 0x002F0007
	IoctlKsDisableEvent                                  = 0x002F000B
	IoctlKsMethod                                        = 0x002F000F
	IoctlKsWriteStream                                   = 0x002F8013
	IoctlKsReadStream                                    = 0x002F4017
	IoctlKsResetState                                    = 0x002F001B
	IoctlKsHandshake                                     = 0x002F001F
	IoctlInternalI8042HookKeyboard                       = 0x000B3FC3
	IoctlInternalI8042HookMouse                          = 0x000F3FC3
	IoctlInternalI8042KeyboardWriteBuffer                = 0x000B3FC7
	IoctlInternalI8042MouseWriteBuffer                   = 0x000F3FC7
	IoctlInternalI8042ControllerWriteBuffer              = 0x000B3FCB
	IoctlInternalI8042KeyboardStartInformation           = 0x000B3FCF
	IoctlInternalI8042MouseStartInformation              = 0x000F3FCF
	IoctlBeepSet                                         = 0x00010000
	IoctlCdromUnloadDriver                               = 0x00025008
	IoctlCdromReadToc                                    = 0x00024000
	IoctlCdromSeekAudioMsf                               = 0x00024004
	IoctlCdromStopAudio                                  = 0x00024008
	IoctlCdromPauseAudio                                 = 0x0002400C
	IoctlCdromResumeAudio                                = 0x00024010
	IoctlCdromGetVolume                                  = 0x00024014
	IoctlCdromPlayAudioMsf                               = 0x00024018
	IoctlCdromSetVolume                                  = 0x00024028
	IoctlCdromReadQChannel                               = 0x0002402C
	IoctlCdromGetControl                                 = 0x00024034
	ObsoleteIoctlCdromGetControl                         = 0x00024034
	IoctlCdromGetLastSession                             = 0x00024038
	IoctlCdromRawRead                                    = 0x0002403E
	IoctlCdromDiskType                                   = 0x00020040
	IoctlCdromGetDriveGeometry                           = 0x0002404C
	IoctlCdromGetDriveGeometryEx                         = 0x00024050
	IoctlCdromReadTocEx                                  = 0x00024054
	IoctlCdromGetConfiguration                           = 0x00024058
	IoctlCdromExclusiveAccess                            = 0x0002C05C
	IoctlCdromSetSpeed                                   = 0x00024060
	IoctlCdromGetInquiryData                             = 0x00024064
	IoctlCdromEnableStreaming                            = 0x00024068
	IoctlCdromSendOpcInformation                         = 0x0002C06C
	IoctlCdromGetPerformance                             = 0x00024070
	IoctlCdromCheckVerify                                = 0x00024800
	IoctlCdromMediaRemoval                               = 0x00024804
	IoctlCdromEjectMedia                                 = 0x00024808
	IoctlCdromLoadMedia                                  = 0x0002480C
	IoctlCdromReserve                                    = 0x00024810
	IoctlCdromRelease                                    = 0x00024814
	IoctlCdromFindNewDevices                             = 0x00024818
	IoctlCdromSimbad                                     = 0x0002400C
	IoctlDvdStartSession                                 = 0x00335000
	IoctlDvdReadKey                                      = 0x00335004
	IoctlDvdSendKey                                      = 0x00335008
	IoctlDvdEndSession                                   = 0x0033500C
	IoctlDvdSetReadAhead                                 = 0x00335010
	IoctlDvdGetRegion                                    = 0x00335014
	IoctlDvdSendKey2                                     = 0x0033D018
	IoctlAacsReadMediaKeyBlockSize                       = 0x003350C0
	IoctlAacsReadMediaKeyBlock                           = 0x003350C4
	IoctlAacsStartSession                                = 0x003350C8
	IoctlAacsEndSession                                  = 0x003350CC
	IoctlAacsSendCertificate                             = 0x003350D0
	IoctlAacsGetCertificate                              = 0x003350D4
	IoctlAacsGetChallengeKey                             = 0x003350D8
	IoctlAacsSendChallengeKey                            = 0x003350DC
	IoctlAacsReadVolumeId                                = 0x003350E0
	IoctlAacsReadSerialNumber                            = 0x003350E4
	IoctlAacsReadMediaId                                 = 0x003350E8
	IoctlAacsReadBindingNonce                            = 0x003350EC
	IoctlAacsGenerateBindingNonce                        = 0x0033D0F0
	IoctlDvdReadStructure                                = 0x00335140
	IoctlStorageSetReadAhead                             = 0x002D4400
	IoctlChangerGetParameters                            = 0x00304000
	IoctlChangerGetStatus                                = 0x00304004
	IoctlChangerGetProductData                           = 0x00304008
	IoctlChangerSetAccess                                = 0x0030C010
	IoctlChangerGetElementStatus                         = 0x0030C014
	IoctlChangerInitializeElementStatus                  = 0x00304018
	IoctlChangerSetPosition                              = 0x0030401C
	IoctlChangerExchangeMedium                           = 0x00304020
	IoctlChangerMoveMedium                               = 0x00304024
	IoctlChangerReinitializeTransport                    = 0x00304028
	IoctlChangerQueryVolumeTags                          = 0x0030C02C
	IoctlDiskGetDriveGeometry                            = 0x00070000
	IoctlDiskGetPartitionInfo                            = 0x00074004
	IoctlDiskSetPartitionInfo                            = 0x0007C008
	IoctlDiskGetDriveLayout                              = 0x0007400C
	IoctlDiskSetDriveLayout                              = 0x0007C010
	IoctlDiskVerify                                      = 0x00070014
	IoctlDiskFormatTracks                                = 0x0007C018
	IoctlDiskReassignBlocks                              = 0x0007C01C
	IoctlDiskPerformance                                 = 0x00070020
	IoctlDiskIsWritable                                  = 0x00070024
	IoctlDiskLogging                                     = 0x00070028
	IoctlDiskFormatTracksEx                              = 0x0007C02C
	IoctlDiskHistogramStructure                          = 0x00070030
	IoctlDiskHistogramData                               = 0x00070034
	IoctlDiskHistogramReset                              = 0x00070038
	IoctlDiskRequestStructure                            = 0x0007003C
	IoctlDiskRequestData                                 = 0x00070040
	IoctlDiskPerformanceOff                              = 0x00070060
	IoctlDiskControllerNumber                            = 0x00070044
	SmartGetVersion                                      = 0x00074080
	SmartSendDriveCommand                                = 0x0007C084
	SmartRcvDriveData                                    = 0x0007C088
	IoctlDiskGetPartitionInfoEx                          = 0x00070048
	IoctlDiskSetPartitionInfoEx                          = 0x0007C04C
	IoctlDiskGetDriveLayoutEx                            = 0x00070050
	IoctlDiskSetDriveLayoutEx                            = 0x0007C054
	IoctlDiskCreateDisk                                  = 0x0007C058
	IoctlDiskGetLengthInfo                               = 0x0007405C
	IoctlDiskGetDriveGeometryEx                          = 0x000700A0
	IoctlDiskReassignBlocksEx                            = 0x0007C0A4
	IoctlDiskUpdateDriveSize                             = 0x0007C0C8
	IoctlDiskGrowPartition                               = 0x0007C0D0
	IoctlDiskGetCacheInformation                         = 0x000740D4
	IoctlDiskSetCacheInformation                         = 0x0007C0D8
	IoctlDiskGetWriteCacheState                          = 0x000740DC
	ObsoleteDiskGetWriteCacheState                       = 0x000740DC
	IoctlDiskDeleteDriveLayout                           = 0x0007C100
	IoctlDiskUpdateProperties                            = 0x00070140
	IoctlDiskFormatDrive                                 = 0x0007C3CC
	IoctlDiskSenseDevice                                 = 0x000703E0
	IoctlDiskGetCacheSetting                             = 0x000740E0
	IoctlDiskSetCacheSetting                             = 0x0007C0E4
	IoctlDiskCopyData                                    = 0x0007C064
	IoctlDiskInternalSetVerify                           = 0x00070403
	IoctlDiskInternalClearVerify                         = 0x00070407
	IoctlDiskInternalSetNotify                           = 0x00070408
	IoctlDiskCheckVerify                                 = 0x00074800
	IoctlDiskMediaRemoval                                = 0x00074804
	IoctlDiskEjectMedia                                  = 0x00074808
	IoctlDiskLoadMedia                                   = 0x0007480C
	IoctlDiskReserve                                     = 0x00074810
	IoctlDiskRelease                                     = 0x00074814
	IoctlDiskFindNewDevices                              = 0x00074818
	IoctlDiskGetMediaTypes                               = 0x00070C00
	IoctlDiskGetPartitionAttributes                      = 0x000700E8
	IoctlDiskSetPartitionAttributes                      = 0x0007C0EC
	IoctlDiskGetDiskAttributes                           = 0x000700F0
	IoctlDiskSetDiskAttributes                           = 0x0007C0F4
	IoctlDiskIsClustered                                 = 0x000700F8
	IoctlDiskGetSanSettings                              = 0x00074200
	IoctlDiskSetSanSettings                              = 0x0007C204
	IoctlDiskGetSnapshotInfo                             = 0x00074208
	IoctlDiskSetSnapshotInfo                             = 0x0007C20C
	IoctlDiskResetSnapshotInfo                           = 0x0007C210
	IoctlDiskSimbad                                      = 0x0007D000
	FtSecondaryRead                                      = 0x00664012
	FtPrimaryRead                                        = 0x00664016
	FtBalancedReadMode                                   = 0x0066001B
	FtSyncRedundantCopy                                  = 0x0066001C
	FtInitializeSet                                      = 0x00660000
	FtRegenerate                                         = 0x00660004
	FtConfigure                                          = 0x0066000B
	FtVerify                                             = 0x0066000C
	FtSequentialWriteMode                                = 0x00660023
	FtParallelWriteMode                                  = 0x00660027
	FtQuerySetState                                      = 0x00660028
	FtClusterSetMemberState                              = 0x0066002C
	FtClusterGetMemberState                              = 0x00660030
	FtEnumerateLogicalDisks                              = 0x00674008
	FtCreateLogicalDisk                                  = 0x0067C000
	FtBreakLogicalDisk                                   = 0x0067C004
	FtQueryLogicalDiskInformation                        = 0x0067400C
	FtOrphanLogicalDiskMember                            = 0x0067C010
	FtReplaceLogicalDiskMember                           = 0x0067C014
	FtQueryNtDeviceNameForLogicalDisk                    = 0x00674018
	FtInitializeLogicalDisk                              = 0x0067C01C
	FtQueryDriveLetterForLogicalDisk                     = 0x00674020
	FtCheckIo                                            = 0x00674024
	FtSetDriveLetterForLogicalDisk                       = 0x0067C028
	FtQueryNtDeviceNameForPartition                      = 0x00674030
	FtChangeNotify                                       = 0x00674034
	FtStopSyncOperations                                 = 0x0067C038
	FtQueryLogicalDiskId                                 = 0x00674190
	FtCreatePartitionLogicalDisk                         = 0x0067C194
	IoctlKeyboardQueryAttributes                         = 0x000B0000
	IoctlKeyboardSetTypematic                            = 0x000B0004
	IoctlKeyboardSetIndicators                           = 0x000B0008
	IoctlKeyboardQueryTypematic                          = 0x000B0020
	IoctlKeyboardQueryIndicators                         = 0x000B0040
	IoctlKeyboardQueryIndicatorTranslation               = 0x000B0080
	IoctlKeyboardInsertData                              = 0x000B0100
	IoctlKeyboardQueryImeStatus                          = 0x000B1000
	IoctlKeyboardSetImeStatus                            = 0x000B1004
	IoctlModemGetPassthrough                             = 0x002B0004
	IoctlModemSetPassthrough                             = 0x002B0008
	IoctlModemSetDleMonitoring                           = 0x002B000C
	IoctlModemGetDle                                     = 0x002B0010
	IoctlModemSetDleShielding                            = 0x002B0014
	IoctlModemStopWaveReceive                            = 0x002B0018
	IoctlModemSendMessage                                = 0x002B001C
	IoctlModemGetMessage                                 = 0x002B0020
	IoctlModemSendGetMessage                             = 0x002B0024
	IoctlModemSendLoopbackMessage                        = 0x002B0028
	IoctlModemCheckForModem                              = 0x002B002C
	IoctlModemSetMinPower                                = 0x002B0030
	IoctlModemWatchForResume                             = 0x002B0034
	IoctlCancelGetSendMessage                            = 0x002B0038
	IoctlSetServerState                                  = 0x002B003C
	IoctlMouseQueryAttributes                            = 0x000F0000
	IoctlMouseInsertData                                 = 0x000F0004
	IoctlNdisReserved5                                   = 0x00170034
	IoctlNdisReserved6                                   = 0x00178038
	IoctlParQueryInformation                             = 0x00160004
	IoctlParSetInformation                               = 0x00160008
	IoctlParQueryDeviceId                                = 0x0016000C
	IoctlParQueryDeviceIdSize                            = 0x00160010
	IoctlIeee1284GetMode                                 = 0x00160014
	IoctlIeee1284Negotiate                               = 0x00160018
	IoctlParSetWriteAddress                              = 0x0016001C
	IoctlParSetReadAddress                               = 0x00160020
	IoctlParGetDeviceCaps                                = 0x00160024
	IoctlParGetDefaultModes                              = 0x00160028
	IoctlParPing                                         = 0x0016002C
	IoctlParQueryRawDeviceId                             = 0x00160030
	IoctlParEcpHostRecovery                              = 0x00160034
	IoctlParGetReadAddress                               = 0x00160038
	IoctlParGetWriteAddress                              = 0x0016003C
	IoctlParTest                                         = 0x00160050
	IoctlParIsPortFree                                   = 0x00160054
	IoctlParQueryLocation                                = 0x00160058
	IoctlScsiPassThrough                                 = 0x0004D004
	IoctlScsiMiniport                                    = 0x0004D008
	IoctlScsiGetInquiryData                              = 0x0004100C
	IoctlScsiGetCapabilities                             = 0x00041010
	IoctlScsiPassThroughDirect                           = 0x0004D014
	IoctlScsiGetAddress                                  = 0x00041018
	IoctlScsiRescanBus                                   = 0x0004101C
	IoctlScsiGetDumpPointers                             = 0x00041020
	IoctlScsiFreeDumpPointers                            = 0x00041024
	IoctlIdePassThrough                                  = 0x0004D028
	IoctlAtaPassThrough                                  = 0x0004D02C
	IoctlAtaPassThroughDirect                            = 0x0004D030
	IoctlAtaMiniport                                     = 0x0004D034
	IoctlMiniportProcessServiceIrp                       = 0x0004D038
	IoctlMpioPassThroughPath                             = 0x0004D03C
	IoctlMpioPassThroughPathDirect                       = 0x0004D040
	IoctlSerialSetBaudRate                               = 0x001B0004
	IoctlSerialSetQueueSize                              = 0x001B0008
	IoctlSerialSetLineControl                            = 0x001B000C
	IoctlSerialSetBreakOn                                = 0x001B0010
	IoctlSerialSetBreakOff                               = 0x001B0014
	IoctlSerialImmediateChar                             = 0x001B0018
	IoctlSerialSetTimeouts                               = 0x001B001C
	IoctlSerialGetTimeouts                               = 0x001B0020
	IoctlSerialSetDtr                                    = 0x001B0024
	IoctlSerialClrDtr                                    = 0x001B0028
	IoctlSerialResetDevice                               = 0x001B002C
	IoctlSerialSetRts                                    = 0x001B0030
	IoctlSerialClrRts                                    = 0x001B0034
	IoctlSerialSetXoff                                   = 0x001B0038
	IoctlSerialSetXon                                    = 0x001B003C
	IoctlSerialGetWaitMask                               = 0x001B0040
	IoctlSerialSetWaitMask                               = 0x001B0044
	IoctlSerialWaitOnMask                                = 0x001B0048
	IoctlSerialPurge                                     = 0x001B004C
	IoctlSerialGetBaudRate                               = 0x001B0050
	IoctlSerialGetLineControl                            = 0x001B0054
	IoctlSerialGetChars                                  = 0x001B0058
	IoctlSerialSetChars                                  = 0x001B005C
	IoctlSerialGetHandflow                               = 0x001B0060
	IoctlSerialSetHandflow                               = 0x001B0064
	IoctlSerialGetModemstatus                            = 0x001B0068
	IoctlSerialGetCommstatus                             = 0x001B006C
	IoctlSerialXoffCounter                               = 0x001B0070
	IoctlSerialGetProperties                             = 0x001B0074
	IoctlSerialGetDtrrts                                 = 0x001B0078
	IoctlSerialLsrmstInsert                              = 0x001B007C
	IoctlSerenumExposeHardware                           = 0x00370200
	IoctlSerenumRemoveHardware                           = 0x00370204
	IoctlSerenumPortDesc                                 = 0x00370208
	IoctlSerenumGetPortName                              = 0x0037020C
	IoctlSerialConfigSize                                = 0x001B0080
	IoctlSerialGetCommconfig                             = 0x001B0084
	IoctlSerialSetCommconfig                             = 0x001B0088
	IoctlSerialGetStats                                  = 0x001B008C
	IoctlSerialClearStats                                = 0x001B0090
	IoctlSerialGetModemControl                           = 0x001B0094
	IoctlSerialSetModemControl                           = 0x001B0098
	IoctlSerialSetFifoControl                            = 0x001B009C
	IoctlSerialInternalDoWaitWake                        = 0x001B0004
	IoctlSerialInternalCancelWaitWake                    = 0x001B0008
	IoctlSerialInternalBasicSettings                     = 0x001B000C
	IoctlSerialInternalRestoreSettings                   = 0x001B0010
	IoctlStorageCheckVerify                              = 0x002D4800
	IoctlStorageCheckVerify2                             = 0x002D0800
	IoctlStorageMediaRemoval                             = 0x002D4804
	IoctlStorageEjectMedia                               = 0x002D4808
	IoctlStorageLoadMedia                                = 0x002D480C
	IoctlStorageLoadMedia2                               = 0x002D080C
	IoctlStorageReserve                                  = 0x002D4810
	IoctlStorageRelease                                  = 0x002D4814
	IoctlStorageFindNewDevices                           = 0x002D4818
	IoctlStorageEjectionControl                          = 0x002D0940
	IoctlStorageMcnControl                               = 0x002D0944
	IoctlStorageGetMediaTypes                            = 0x002D0C00
	IoctlStorageGetMediaTypesEx                          = 0x002D0C04
	IoctlStorageGetMediaSerialNumber                     = 0x002D0C10
	IoctlStorageGetHotplugInfo                           = 0x002D0C14
	IoctlStorageSetHotplugInfo                           = 0x002DCC18
	IoctlStorageResetBus                                 = 0x002D5000
	IoctlStorageResetDevice                              = 0x002D5004
	IoctlStorageBreakReservation                         = 0x002D5014
	IoctlStoragePersistentReserveIn                      = 0x002D5018
	IoctlStoragePersistentReserveOut                     = 0x002DD01C
	IoctlStorageGetDeviceNumber                          = 0x002D1080
	IoctlStoragePredictFailure                           = 0x002D1100
	IoctlStorageReadCapacity                             = 0x002D5140
	IoctlStorageQueryProperty                            = 0x002D1400
	IoctlStorageManageDataSetAttributes                  = 0x002D9404
	IoctlStorageGetBcProperties                          = 0x002D5800
	IoctlStorageAllocateBcStream                         = 0x002DD804
	IoctlStorageFreeBcStream                             = 0x002DD808
	IoctlStorageCheckPriorityHintSupport                 = 0x002D1880
	ObsoleteIoctlStorageResetBus                         = 0x002DD000
	ObsoleteIoctlStorageResetDevice                      = 0x002DD004
	IoctlTapeErase                                       = 0x001FC000
	IoctlTapePrepare                                     = 0x001F4004
	IoctlTapeWriteMarks                                  = 0x001FC008
	IoctlTapeGetPosition                                 = 0x001F400C
	IoctlTapeSetPosition                                 = 0x001F4010
	IoctlTapeGetDriveParams                              = 0x001F4014
	IoctlTapeSetDriveParams                              = 0x001FC018
	IoctlTapeGetMediaParams                              = 0x001F401C
	IoctlTapeSetMediaParams                              = 0x001F4020
	IoctlTapeGetStatus                                   = 0x001F4024
	IoctlTapeCreatePartition                             = 0x001FC028
	IoctlTapeMediaRemoval                                = 0x001F4804
	IoctlTapeEjectMedia                                  = 0x001F4808
	IoctlTapeLoadMedia                                   = 0x001F480C
	IoctlTapeReserve                                     = 0x001F4810
	IoctlTapeRelease                                     = 0x001F4814
	IoctlTapeCheckVerify                                 = 0x001F4800
	IoctlTapeFindNewDevices                              = 0x00074818
	IoctlVolumeGetVolumeDiskExtents                      = 0x00560000
	IoctlVolumeOnline                                    = 0x0056C008
	IoctlVolumeOffline                                   = 0x0056C00C
	IoctlVolumeIsClustered                               = 0x00560030
	IoctlVolumeGetGptAttributes                          = 0x00560038
	IoctlVolumeSupportsOnlineOffline                     = 0x00560004
	IoctlVolumeIsOffline                                 = 0x00560010
	IoctlVolumeIsIoCapable                               = 0x00560014
	IoctlVolumeQueryFailoverSet                          = 0x00560018
	IoctlVolumeQueryVolumeNumber                         = 0x0056001C
	IoctlVolumeLogicalToPhysical                         = 0x00560020
	IoctlVolumePhysicalToLogical                         = 0x00560024
	IoctlVolumeIsPartition                               = 0x00560028
	IoctlVolumeReadPlex                                  = 0x0056402E
	IoctlVolumeSetGptAttributes                          = 0x00560034
	IoctlVolumeGetBcProperties                           = 0x0056403C
	IoctlVolumeAllocateBcStream                          = 0x0056C040
	IoctlVolumeFreeBcStream                              = 0x0056C044
	IoctlVolumeIsDynamic                                 = 0x00560048
	IoctlVolumePrepareForCriticalIo                      = 0x0056C04C
	IoctlVolumeQueryAllocationHint                       = 0x00564052
	IoctlVolumeUpdateProperties                          = 0x00560054
	IoctlVolumeQueryMinimumShrinkSize                    = 0x00564058
	IoctlVolumePrepareForShrink                          = 0x0056C05C
	IoctlPmiGetCapabilities                              = 0x00454000
	IoctlPmiGetConfiguration                             = 0x00454004
	IoctlPmiSetConfiguration                             = 0x00458008
	IoctlPmiGetMeasurement                               = 0x0045400C
	IoctlPmiRegisterEventNotify                          = 0x0045C010
	IoctlBiometricVendor                                 = 0x00442000
)

func (k IoctlKind) String() string {
	switch k {
	case 0x00120004:
		return "MonitorIoctlEnableMonitor"
	case 0x00120008:
		return "MonitorIoctlDisableMonitor"
	case 0x000B0203:
		return "IoctlInternalKeyboardConnect"
	case 0x000B0403:
		return "IoctlInternalKeyboardDisconnect"
	case 0x000B0803:
		return "IoctlInternalKeyboardEnable"
	case 0x000B1003:
		return "IoctlInternalKeyboardDisable"
	case 0x000F0203:
		return "IoctlInternalMouseConnect"
	case 0x000F0403:
		return "IoctlInternalMouseDisconnect"
	case 0x000F0803:
		return "IoctlInternalMouseEnable"
	case 0x000F1003:
		return "IoctlInternalMouseDisable"
	case 0x00222000:
		return "IoctlDoKernelmodeSamples"
	case 0x00222004:
		return "IoctlRegisterCallback"
	case 0x00222008:
		return "IoctlUnregisterCallback"
	case 0x0022200C:
		return "IoctlGetCallbackVersion"
	case 0x00222000:
		return "IoctlGetVersion"
	case 0x00222007:
		return "IoctlReset"
	case 0x00222008:
		return "IoctlReadwriteMailbox"
	case 0x0022200C:
		return "IoctlMailboxWait"
	case 0x00226012:
		return "IoctlReadDma"
	case 0x0022A015:
		return "IoctlWriteDma"
	case 0x0032C000:
		return "IoctlAcpiAsyncEvalMethod"
	case 0x0032C004:
		return "IoctlAcpiEvalMethod"
	case 0x0032C010:
		return "IoctlAcpiAcquireGlobalLock"
	case 0x0032C014:
		return "IoctlAcpiReleaseGlobalLock"
	case 0x0032C018:
		return "IoctlAcpiEvalMethodEx"
	case 0x0032C01C:
		return "IoctlAcpiAsyncEvalMethodEx"
	case 0x0032C020:
		return "IoctlAcpiEnumChildren"
	case 0x002A0000:
		return "IoctlAvcUpdateVirtualSubunitInfo"
	case 0x002A0004:
		return "IoctlAvcRemoveVirtualSubunitInfo"
	case 0x002A0008:
		return "IoctlAvcBusReset"
	case 0x003A2022:
		return "IoctlDot4CreateSocket"
	case 0x003A202A:
		return "IoctlDot4DestroySocket"
	case 0x003A2026:
		return "IoctlDot4WaitForChannel"
	case 0x003A2006:
		return "IoctlDot4OpenChannel"
	case 0x003A2008:
		return "IoctlDot4CloseChannel"
	case 0x003A200E:
		return "IoctlDot4Read"
	case 0x003A2011:
		return "IoctlDot4Write"
	case 0x003A2014:
		return "IoctlDot4AddActivityBroadcast"
	case 0x003A2018:
		return "IoctlDot4RemoveActivityBroadcast"
	case 0x003A201E:
		return "IoctlDot4WaitActivityBroadcast"
	case 0x736DC004:
		return "IoctlMpdsmRegister"
	case 0x736DC008:
		return "IoctlMpdsmDeregister"
	case 0x004D0000:
		return "IoctlMountdevQueryUniqueId"
	case 0x004D000C:
		return "IoctlMountdevQuerySuggestedLinkName"
	case 0x004DC010:
		return "IoctlMountdevLinkCreated"
	case 0x004DC014:
		return "IoctlMountdevLinkDeleted"
	case 0x004D0018:
		return "IoctlMountdevQueryStableGuid"
	case 0x006DC000:
		return "IoctlMountmgrCreatePoint"
	case 0x006DC004:
		return "IoctlMountmgrDeletePoints"
	case 0x006D0008:
		return "IoctlMountmgrQueryPoints"
	case 0x006DC00C:
		return "IoctlMountmgrDeletePointsDbonly"
	case 0x006DC010:
		return "IoctlMountmgrNextDriveLetter"
	case 0x006DC014:
		return "IoctlMountmgrAutoDlAssignments"
	case 0x006DC018:
		return "IoctlMountmgrVolumeMountPointCreated"
	case 0x006DC01C:
		return "IoctlMountmgrVolumeMountPointDeleted"
	case 0x006D4020:
		return "IoctlMountmgrChangeNotify"
	case 0x006DC024:
		return "IoctlMountmgrKeepLinksWhenOffline"
	case 0x006D4028:
		return "IoctlMountmgrCheckUnprocessedVolumes"
	case 0x006D402C:
		return "IoctlMountmgrVolumeArrivalNotification"
	case 0x004D0008:
		return "IoctlMountdevQueryDeviceName"
	case 0x006D0030:
		return "IoctlMountmgrQueryDosVolumePath"
	case 0x006D0034:
		return "IoctlMountmgrQueryDosVolumePaths"
	case 0x006DC038:
		return "IoctlMountmgrScrubRegistry"
	case 0x006D003C:
		return "IoctlMountmgrQueryAutoMount"
	case 0x006DC040:
		return "IoctlMountmgrSetAutoMount"
	case 0x006DC044:
		return "IoctlMountmgrBootDlAssignment"
	case 0x006D4048:
		return "IoctlMountmgrTracelogCache"
	case 0x00000000:
		return "IoctlAvioAllocateStream"
	case 0x00000000:
		return "IoctlAvioFreeStream"
	case 0x00000000:
		return "IoctlAvioModifyStream"
	case 0xC0C08048:
		return "NlbIoctlRegisterHook"
	case 0xC0C08054:
		return "NlbPublicIoctlClientStickinessNotify"
	case 0x00042EE0:
		return "IoctlGetTupleData"
	case 0x00042EF0:
		return "IoctlSocketInformation"
	case 0x0004AF08:
		return "IoctlPcmciaHideDevice"
	case 0x0004AF0C:
		return "IoctlPcmciaRevealDevice"
	case 0x00043073:
		return "IoctlSdSubmitRequest"
	case 0x00090000:
		return "FsctlRequestOplockLevel1"
	case 0x00090004:
		return "FsctlRequestOplockLevel2"
	case 0x00090008:
		return "FsctlRequestBatchOplock"
	case 0x0009000C:
		return "FsctlOplockBreakAcknowledge"
	case 0x00090010:
		return "FsctlOpbatchAckClosePending"
	case 0x00090014:
		return "FsctlOplockBreakNotify"
	case 0x00090018:
		return "FsctlLockVolume"
	case 0x0009001C:
		return "FsctlUnlockVolume"
	case 0x00090020:
		return "FsctlDismountVolume"
	case 0x00090028:
		return "FsctlIsVolumeMounted"
	case 0x0009002C:
		return "FsctlIsPathnameValid"
	case 0x00090030:
		return "FsctlMarkVolumeDirty"
	case 0x0009003B:
		return "FsctlQueryRetrievalPointers"
	case 0x0009003C:
		return "FsctlGetCompression"
	case 0x0009C040:
		return "FsctlSetCompression"
	case 0x0009004F:
		return "FsctlSetBootloaderAccessed"
	case 0x00090050:
		return "FsctlOplockBreakAckNo2"
	case 0x00090054:
		return "FsctlInvalidateVolumes"
	case 0x00090058:
		return "FsctlQueryFatBpb"
	case 0x0009005C:
		return "FsctlRequestFilterOplock"
	case 0x00090060:
		return "FsctlFilesystemGetStatistics"
	case 0x00090064:
		return "FsctlGetNtfsVolumeData"
	case 0x00090068:
		return "FsctlGetNtfsFileRecord"
	case 0x0009006F:
		return "FsctlGetVolumeBitmap"
	case 0x00090073:
		return "FsctlGetRetrievalPointers"
	case 0x00090074:
		return "FsctlMoveFile"
	case 0x00090078:
		return "FsctlIsVolumeDirty"
	case 0x00090083:
		return "FsctlAllowExtendedDasdIo"
	case 0x0009008F:
		return "FsctlFindFilesBySid"
	case 0x00090098:
		return "FsctlSetObjectId"
	case 0x0009009C:
		return "FsctlGetObjectId"
	case 0x000900A0:
		return "FsctlDeleteObjectId"
	case 0x000900A4:
		return "FsctlSetReparsePoint"
	case 0x000900A8:
		return "FsctlGetReparsePoint"
	case 0x000900AC:
		return "FsctlDeleteReparsePoint"
	case 0x000900B3:
		return "FsctlEnumUsnData"
	case 0x000940B7:
		return "FsctlSecurityIdCheck"
	case 0x000900BB:
		return "FsctlReadUsnJournal"
	case 0x000900BC:
		return "FsctlSetObjectIdExtended"
	case 0x000900C0:
		return "FsctlCreateOrGetObjectId"
	case 0x000900C4:
		return "FsctlSetSparse"
	case 0x000980C8:
		return "FsctlSetZeroData"
	case 0x000940CF:
		return "FsctlQueryAllocatedRanges"
	case 0x000980D0:
		return "FsctlEnableUpgrade"
	case 0x000900D7:
		return "FsctlSetEncryption"
	case 0x000900DB:
		return "FsctlEncryptionFsctlIo"
	case 0x000900DF:
		return "FsctlWriteRawEncrypted"
	case 0x000900E3:
		return "FsctlReadRawEncrypted"
	case 0x000900E7:
		return "FsctlCreateUsnJournal"
	case 0x000900EB:
		return "FsctlReadFileUsnData"
	case 0x000900EF:
		return "FsctlWriteUsnCloseRecord"
	case 0x000900F0:
		return "FsctlExtendVolume"
	case 0x000900F4:
		return "FsctlQueryUsnJournal"
	case 0x000900F8:
		return "FsctlDeleteUsnJournal"
	case 0x000900FC:
		return "FsctlMarkHandle"
	case 0x00090100:
		return "FsctlSisCopyfile"
	case 0x0009C104:
		return "FsctlSisLinkFiles"
	case 0x00090117:
		return "FsctlRecallFile"
	case 0x0009411E:
		return "FsctlReadFromPlex"
	case 0x00090120:
		return "FsctlFilePrefetch"
	case 0x00098130:
		return "FsctlMakeMediaCompatible"
	case 0x00098134:
		return "FsctlSetDefectManagement"
	case 0x00090138:
		return "FsctlQuerySparingInfo"
	case 0x0009013C:
		return "FsctlQueryOnDiskVolumeInfo"
	case 0x00090140:
		return "FsctlSetVolumeCompressionState"
	case 0x00098144:
		return "FsctlTxfsModifyRm"
	case 0x00094148:
		return "FsctlTxfsQueryRmInformation"
	case 0x00098150:
		return "FsctlTxfsRollforwardRedo"
	case 0x00098154:
		return "FsctlTxfsRollforwardUndo"
	case 0x00098158:
		return "FsctlTxfsStartRm"
	case 0x0009815C:
		return "FsctlTxfsShutdownRm"
	case 0x00094160:
		return "FsctlTxfsReadBackupInformation"
	case 0x00098164:
		return "FsctlTxfsWriteBackupInformation"
	case 0x00098168:
		return "FsctlTxfsCreateSecondaryRm"
	case 0x0009416C:
		return "FsctlTxfsGetMetadataInfo"
	case 0x00094170:
		return "FsctlTxfsGetTransactedVersion"
	case 0x00098178:
		return "FsctlTxfsSavepointInformation"
	case 0x0009817C:
		return "FsctlTxfsCreateMiniversion"
	case 0x0009418C:
		return "FsctlTxfsTransactionActive"
	case 0x00090194:
		return "FsctlSetZeroOnDeallocation"
	case 0x00090198:
		return "FsctlSetRepair"
	case 0x0009019C:
		return "FsctlGetRepair"
	case 0x000901A0:
		return "FsctlWaitForRepair"
	case 0x000901A8:
		return "FsctlInitiateRepair"
	case 0x000901AF:
		return "FsctlCscInternal"
	case 0x000901B0:
		return "FsctlShrinkVolume"
	case 0x000901B4:
		return "FsctlSetShortNameBehavior"
	case 0x000901B8:
		return "FsctlDfsrSetGhostHandleState"
	case 0x000941E4:
		return "FsctlTxfsListTransactions"
	case 0x000901E8:
		return "FsctlQueryPagefileEncryption"
	case 0x000901EC:
		return "FsctlResetVolumeAllocationHints"
	case 0x000901F0:
		return "FsctlQueryDependentVolume"
	case 0x000901F4:
		return "FsctlSdGlobalChange"
	case 0x000901F8:
		return "FsctlTxfsReadBackupInformation2"
	case 0x000901FC:
		return "FsctlLookupStreamFromCluster"
	case 0x00090200:
		return "FsctlTxfsWriteBackupInformation2"
	case 0x00090204:
		return "FsctlFileTypeNotification"
	case 0x00090230:
		return "FsctlGetBootAreaInfo"
	case 0x00090234:
		return "FsctlGetRetrievalPointerBase"
	case 0x00090238:
		return "FsctlSetPersistentVolumeState"
	case 0x0009023C:
		return "FsctlQueryPersistentVolumeState"
	case 0x00090240:
		return "FsctlRequestOplock"
	case 0x00090244:
		return "FsctlCsvTunnelRequest"
	case 0x00090248:
		return "FsctlIsCsvFile"
	case 0x0009024C:
		return "FsctlQueryFileSystemRecognition"
	case 0x00090250:
		return "FsctlCsvGetVolumePathName"
	case 0x00090254:
		return "FsctlCsvGetVolumeNameForVolumeMountPoint"
	case 0x00090258:
		return "FsctlCsvGetVolumePathNamesForVolumeName"
	case 0x0009025C:
		return "FsctlIsFileOnCsvVolume"
	case 0x001400E8:
		return "FsctlLmrGetLinkTrackingInformation"
	case 0x001400EC:
		return "FsctlLmrSetLinkTrackingInformation"
	case 0x001400F0:
		return "IoctlLmrAreFileObjectsOnSameServer"
	case 0x00110000:
		return "FsctlPipeAssignEvent"
	case 0x00110004:
		return "FsctlPipeDisconnect"
	case 0x00110008:
		return "FsctlPipeListen"
	case 0x0011400C:
		return "FsctlPipePeek"
	case 0x00110010:
		return "FsctlPipeQueryEvent"
	case 0x0011C017:
		return "FsctlPipeTransceive"
	case 0x00110018:
		return "FsctlPipeWait"
	case 0x0011001C:
		return "FsctlPipeImpersonate"
	case 0x00110020:
		return "FsctlPipeSetClientProcess"
	case 0x00110024:
		return "FsctlPipeQueryClientProcess"
	case 0x00110028:
		return "FsctlPipeGetPipeAttribute"
	case 0x0011002C:
		return "FsctlPipeSetPipeAttribute"
	case 0x00110030:
		return "FsctlPipeGetConnectionAttribute"
	case 0x00110034:
		return "FsctlPipeSetConnectionAttribute"
	case 0x00110038:
		return "FsctlPipeGetHandleAttribute"
	case 0x0011003C:
		return "FsctlPipeSetHandleAttribute"
	case 0x00118040:
		return "FsctlPipeFlush"
	case 0x00115FF4:
		return "FsctlPipeInternalRead"
	case 0x00119FF8:
		return "FsctlPipeInternalWrite"
	case 0x0011DFFF:
		return "FsctlPipeInternalTransceive"
	case 0x00116000:
		return "FsctlPipeInternalReadOvflow"
	case 0x000C4003:
		return "FsctlMailslotPeek"
	case 0x0014018F:
		return "IoctlRedirQueryPath"
	case 0x00140193:
		return "IoctlRedirQueryPathEx"
	case 0x0053C000:
		return "IoctlVolsnapFlushAndHoldWrites"
	case 0x0016002C:
		return "IoctlInternalParallelPortAllocate"
	case 0x00160030:
		return "IoctlInternalGetParallelPortInfo"
	case 0x00160034:
		return "IoctlInternalParallelConnectInterrupt"
	case 0x00160038:
		return "IoctlInternalParallelDisconnectInterrupt"
	case 0x0016003C:
		return "IoctlInternalReleaseParallelPortInfo"
	case 0x00160044:
		return "IoctlInternalGetMoreParallelPortInfo"
	case 0x00160048:
		return "IoctlInternalParchipConnect"
	case 0x0016004C:
		return "IoctlInternalParallelSetChipMode"
	case 0x00160050:
		return "IoctlInternalParallelClearChipMode"
	case 0x00160054:
		return "IoctlInternalGetParallelPnpInfo"
	case 0x00160058:
		return "IoctlInternalInit12843Bus"
	case 0x0016005C:
		return "IoctlInternalSelectDevice"
	case 0x00160060:
		return "IoctlInternalDeselectDevice"
	case 0x00160074:
		return "IoctlInternalGetParportFdo"
	case 0x00160078:
		return "IoctlInternalParclassConnect"
	case 0x0016007C:
		return "IoctlInternalParclassDisconnect"
	case 0x00160080:
		return "IoctlInternalDisconnectIdle"
	case 0x00160094:
		return "IoctlInternalLockPort"
	case 0x00160098:
		return "IoctlInternalUnlockPort"
	case 0x001600A0:
		return "IoctlInternalParallelPortFree"
	case 0x001600A4:
		return "IoctlInternalPardot3Connect"
	case 0x001600A8:
		return "IoctlInternalPardot3Disconnect"
	case 0x001600AC:
		return "IoctlInternalPardot3Reset"
	case 0x001600B0:
		return "IoctlInternalPardot3Signal"
	case 0x001600C8:
		return "IoctlInternalRegisterForRemovalRelations"
	case 0x001600CC:
		return "IoctlInternalUnregisterForRemovalRelations"
	case 0x001600D0:
		return "IoctlInternalLockPortNoSelect"
	case 0x001600D4:
		return "IoctlInternalUnlockPortNoDeselect"
	case 0x001600D8:
		return "IoctlInternalDisableEndOfChainBusRescan"
	case 0x001600DC:
		return "IoctlInternalEnableEndOfChainBusRescan"
	case 0x0040C108:
		return "IoctlWpdMessageReadwriteAccess"
	case 0x00404108:
		return "IoctlWpdMessageReadAccess"
	case 0x00190012:
		return "IoctlScsiscanCmd"
	case 0x00190016:
		return "IoctlScsiscanLockdevice"
	case 0x0019001A:
		return "IoctlScsiscanUnlockdevice"
	case 0x0019001C:
		return "IoctlScsiscanSetTimeout"
	case 0x00190022:
		return "IoctlScsiscanGetInfo"
	case 0x002A0000:
		return "IoctlSwenumInstallInterface"
	case 0x002A0004:
		return "IoctlSwenumRemoveInterface"
	case 0x002A400B:
		return "IoctlSwenumGetBusId"
	case 0x80002004:
		return "IoctlCancelIo"
	case 0x80002008:
		return "IoctlWaitOnDeviceEvent"
	case 0x8000200C:
		return "IoctlReadRegisters"
	case 0x80002010:
		return "IoctlWriteRegisters"
	case 0x80002014:
		return "IoctlGetChannelAlignRqst"
	case 0x80002018:
		return "IoctlGetDeviceDescriptor"
	case 0x8000201C:
		return "IoctlResetPipe"
	case 0x80002020:
		return "IoctlGetUsbDescriptor"
	case 0x80002024:
		return "IoctlSendUsbRequest"
	case 0x80002028:
		return "IoctlGetPipeConfiguration"
	case 0x8000202C:
		return "IoctlSetTimeout"
	case 0x002D1404:
		return "IoctlEhstorDeviceSetAuthzState"
	case 0x002D1408:
		return "IoctlEhstorDeviceGetAuthzState"
	case 0x002D140C:
		return "IoctlEhstorDeviceSiloCommand"
	case 0x002D1410:
		return "IoctlEhstorDeviceEnumeratePdos"
	case 0x002F0003:
		return "IoctlKsProperty"
	case 0x002F0007:
		return "IoctlKsEnableEvent"
	case 0x002F000B:
		return "IoctlKsDisableEvent"
	case 0x002F000F:
		return "IoctlKsMethod"
	case 0x002F8013:
		return "IoctlKsWriteStream"
	case 0x002F4017:
		return "IoctlKsReadStream"
	case 0x002F001B:
		return "IoctlKsResetState"
	case 0x002F001F:
		return "IoctlKsHandshake"
	case 0x000B3FC3:
		return "IoctlInternalI8042HookKeyboard"
	case 0x000F3FC3:
		return "IoctlInternalI8042HookMouse"
	case 0x000B3FC7:
		return "IoctlInternalI8042KeyboardWriteBuffer"
	case 0x000F3FC7:
		return "IoctlInternalI8042MouseWriteBuffer"
	case 0x000B3FCB:
		return "IoctlInternalI8042ControllerWriteBuffer"
	case 0x000B3FCF:
		return "IoctlInternalI8042KeyboardStartInformation"
	case 0x000F3FCF:
		return "IoctlInternalI8042MouseStartInformation"
	case 0x00010000:
		return "IoctlBeepSet"
	case 0x00025008:
		return "IoctlCdromUnloadDriver"
	case 0x00024000:
		return "IoctlCdromReadToc"
	case 0x00024004:
		return "IoctlCdromSeekAudioMsf"
	case 0x00024008:
		return "IoctlCdromStopAudio"
	case 0x0002400C:
		return "IoctlCdromPauseAudio"
	case 0x00024010:
		return "IoctlCdromResumeAudio"
	case 0x00024014:
		return "IoctlCdromGetVolume"
	case 0x00024018:
		return "IoctlCdromPlayAudioMsf"
	case 0x00024028:
		return "IoctlCdromSetVolume"
	case 0x0002402C:
		return "IoctlCdromReadQChannel"
	case 0x00024034:
		return "IoctlCdromGetControl"
	case 0x00024034:
		return "ObsoleteIoctlCdromGetControl"
	case 0x00024038:
		return "IoctlCdromGetLastSession"
	case 0x0002403E:
		return "IoctlCdromRawRead"
	case 0x00020040:
		return "IoctlCdromDiskType"
	case 0x0002404C:
		return "IoctlCdromGetDriveGeometry"
	case 0x00024050:
		return "IoctlCdromGetDriveGeometryEx"
	case 0x00024054:
		return "IoctlCdromReadTocEx"
	case 0x00024058:
		return "IoctlCdromGetConfiguration"
	case 0x0002C05C:
		return "IoctlCdromExclusiveAccess"
	case 0x00024060:
		return "IoctlCdromSetSpeed"
	case 0x00024064:
		return "IoctlCdromGetInquiryData"
	case 0x00024068:
		return "IoctlCdromEnableStreaming"
	case 0x0002C06C:
		return "IoctlCdromSendOpcInformation"
	case 0x00024070:
		return "IoctlCdromGetPerformance"
	case 0x00024800:
		return "IoctlCdromCheckVerify"
	case 0x00024804:
		return "IoctlCdromMediaRemoval"
	case 0x00024808:
		return "IoctlCdromEjectMedia"
	case 0x0002480C:
		return "IoctlCdromLoadMedia"
	case 0x00024810:
		return "IoctlCdromReserve"
	case 0x00024814:
		return "IoctlCdromRelease"
	case 0x00024818:
		return "IoctlCdromFindNewDevices"
	case 0x0002400C:
		return "IoctlCdromSimbad"
	case 0x00335000:
		return "IoctlDvdStartSession"
	case 0x00335004:
		return "IoctlDvdReadKey"
	case 0x00335008:
		return "IoctlDvdSendKey"
	case 0x0033500C:
		return "IoctlDvdEndSession"
	case 0x00335010:
		return "IoctlDvdSetReadAhead"
	case 0x00335014:
		return "IoctlDvdGetRegion"
	case 0x0033D018:
		return "IoctlDvdSendKey2"
	case 0x003350C0:
		return "IoctlAacsReadMediaKeyBlockSize"
	case 0x003350C4:
		return "IoctlAacsReadMediaKeyBlock"
	case 0x003350C8:
		return "IoctlAacsStartSession"
	case 0x003350CC:
		return "IoctlAacsEndSession"
	case 0x003350D0:
		return "IoctlAacsSendCertificate"
	case 0x003350D4:
		return "IoctlAacsGetCertificate"
	case 0x003350D8:
		return "IoctlAacsGetChallengeKey"
	case 0x003350DC:
		return "IoctlAacsSendChallengeKey"
	case 0x003350E0:
		return "IoctlAacsReadVolumeId"
	case 0x003350E4:
		return "IoctlAacsReadSerialNumber"
	case 0x003350E8:
		return "IoctlAacsReadMediaId"
	case 0x003350EC:
		return "IoctlAacsReadBindingNonce"
	case 0x0033D0F0:
		return "IoctlAacsGenerateBindingNonce"
	case 0x00335140:
		return "IoctlDvdReadStructure"
	case 0x002D4400:
		return "IoctlStorageSetReadAhead"
	case 0x00304000:
		return "IoctlChangerGetParameters"
	case 0x00304004:
		return "IoctlChangerGetStatus"
	case 0x00304008:
		return "IoctlChangerGetProductData"
	case 0x0030C010:
		return "IoctlChangerSetAccess"
	case 0x0030C014:
		return "IoctlChangerGetElementStatus"
	case 0x00304018:
		return "IoctlChangerInitializeElementStatus"
	case 0x0030401C:
		return "IoctlChangerSetPosition"
	case 0x00304020:
		return "IoctlChangerExchangeMedium"
	case 0x00304024:
		return "IoctlChangerMoveMedium"
	case 0x00304028:
		return "IoctlChangerReinitializeTransport"
	case 0x0030C02C:
		return "IoctlChangerQueryVolumeTags"
	case 0x00070000:
		return "IoctlDiskGetDriveGeometry"
	case 0x00074004:
		return "IoctlDiskGetPartitionInfo"
	case 0x0007C008:
		return "IoctlDiskSetPartitionInfo"
	case 0x0007400C:
		return "IoctlDiskGetDriveLayout"
	case 0x0007C010:
		return "IoctlDiskSetDriveLayout"
	case 0x00070014:
		return "IoctlDiskVerify"
	case 0x0007C018:
		return "IoctlDiskFormatTracks"
	case 0x0007C01C:
		return "IoctlDiskReassignBlocks"
	case 0x00070020:
		return "IoctlDiskPerformance"
	case 0x00070024:
		return "IoctlDiskIsWritable"
	case 0x00070028:
		return "IoctlDiskLogging"
	case 0x0007C02C:
		return "IoctlDiskFormatTracksEx"
	case 0x00070030:
		return "IoctlDiskHistogramStructure"
	case 0x00070034:
		return "IoctlDiskHistogramData"
	case 0x00070038:
		return "IoctlDiskHistogramReset"
	case 0x0007003C:
		return "IoctlDiskRequestStructure"
	case 0x00070040:
		return "IoctlDiskRequestData"
	case 0x00070060:
		return "IoctlDiskPerformanceOff"
	case 0x00070044:
		return "IoctlDiskControllerNumber"
	case 0x00074080:
		return "SmartGetVersion"
	case 0x0007C084:
		return "SmartSendDriveCommand"
	case 0x0007C088:
		return "SmartRcvDriveData"
	case 0x00070048:
		return "IoctlDiskGetPartitionInfoEx"
	case 0x0007C04C:
		return "IoctlDiskSetPartitionInfoEx"
	case 0x00070050:
		return "IoctlDiskGetDriveLayoutEx"
	case 0x0007C054:
		return "IoctlDiskSetDriveLayoutEx"
	case 0x0007C058:
		return "IoctlDiskCreateDisk"
	case 0x0007405C:
		return "IoctlDiskGetLengthInfo"
	case 0x000700A0:
		return "IoctlDiskGetDriveGeometryEx"
	case 0x0007C0A4:
		return "IoctlDiskReassignBlocksEx"
	case 0x0007C0C8:
		return "IoctlDiskUpdateDriveSize"
	case 0x0007C0D0:
		return "IoctlDiskGrowPartition"
	case 0x000740D4:
		return "IoctlDiskGetCacheInformation"
	case 0x0007C0D8:
		return "IoctlDiskSetCacheInformation"
	case 0x000740DC:
		return "IoctlDiskGetWriteCacheState"
	case 0x000740DC:
		return "ObsoleteDiskGetWriteCacheState"
	case 0x0007C100:
		return "IoctlDiskDeleteDriveLayout"
	case 0x00070140:
		return "IoctlDiskUpdateProperties"
	case 0x0007C3CC:
		return "IoctlDiskFormatDrive"
	case 0x000703E0:
		return "IoctlDiskSenseDevice"
	case 0x000740E0:
		return "IoctlDiskGetCacheSetting"
	case 0x0007C0E4:
		return "IoctlDiskSetCacheSetting"
	case 0x0007C064:
		return "IoctlDiskCopyData"
	case 0x00070403:
		return "IoctlDiskInternalSetVerify"
	case 0x00070407:
		return "IoctlDiskInternalClearVerify"
	case 0x00070408:
		return "IoctlDiskInternalSetNotify"
	case 0x00074800:
		return "IoctlDiskCheckVerify"
	case 0x00074804:
		return "IoctlDiskMediaRemoval"
	case 0x00074808:
		return "IoctlDiskEjectMedia"
	case 0x0007480C:
		return "IoctlDiskLoadMedia"
	case 0x00074810:
		return "IoctlDiskReserve"
	case 0x00074814:
		return "IoctlDiskRelease"
	case 0x00074818:
		return "IoctlDiskFindNewDevices"
	case 0x00070C00:
		return "IoctlDiskGetMediaTypes"
	case 0x000700E8:
		return "IoctlDiskGetPartitionAttributes"
	case 0x0007C0EC:
		return "IoctlDiskSetPartitionAttributes"
	case 0x000700F0:
		return "IoctlDiskGetDiskAttributes"
	case 0x0007C0F4:
		return "IoctlDiskSetDiskAttributes"
	case 0x000700F8:
		return "IoctlDiskIsClustered"
	case 0x00074200:
		return "IoctlDiskGetSanSettings"
	case 0x0007C204:
		return "IoctlDiskSetSanSettings"
	case 0x00074208:
		return "IoctlDiskGetSnapshotInfo"
	case 0x0007C20C:
		return "IoctlDiskSetSnapshotInfo"
	case 0x0007C210:
		return "IoctlDiskResetSnapshotInfo"
	case 0x0007D000:
		return "IoctlDiskSimbad"
	case 0x00664012:
		return "FtSecondaryRead"
	case 0x00664016:
		return "FtPrimaryRead"
	case 0x0066001B:
		return "FtBalancedReadMode"
	case 0x0066001C:
		return "FtSyncRedundantCopy"
	case 0x00660000:
		return "FtInitializeSet"
	case 0x00660004:
		return "FtRegenerate"
	case 0x0066000B:
		return "FtConfigure"
	case 0x0066000C:
		return "FtVerify"
	case 0x00660023:
		return "FtSequentialWriteMode"
	case 0x00660027:
		return "FtParallelWriteMode"
	case 0x00660028:
		return "FtQuerySetState"
	case 0x0066002C:
		return "FtClusterSetMemberState"
	case 0x00660030:
		return "FtClusterGetMemberState"
	case 0x00674008:
		return "FtEnumerateLogicalDisks"
	case 0x0067C000:
		return "FtCreateLogicalDisk"
	case 0x0067C004:
		return "FtBreakLogicalDisk"
	case 0x0067400C:
		return "FtQueryLogicalDiskInformation"
	case 0x0067C010:
		return "FtOrphanLogicalDiskMember"
	case 0x0067C014:
		return "FtReplaceLogicalDiskMember"
	case 0x00674018:
		return "FtQueryNtDeviceNameForLogicalDisk"
	case 0x0067C01C:
		return "FtInitializeLogicalDisk"
	case 0x00674020:
		return "FtQueryDriveLetterForLogicalDisk"
	case 0x00674024:
		return "FtCheckIo"
	case 0x0067C028:
		return "FtSetDriveLetterForLogicalDisk"
	case 0x00674030:
		return "FtQueryNtDeviceNameForPartition"
	case 0x00674034:
		return "FtChangeNotify"
	case 0x0067C038:
		return "FtStopSyncOperations"
	case 0x00674190:
		return "FtQueryLogicalDiskId"
	case 0x0067C194:
		return "FtCreatePartitionLogicalDisk"
	case 0x000B0000:
		return "IoctlKeyboardQueryAttributes"
	case 0x000B0004:
		return "IoctlKeyboardSetTypematic"
	case 0x000B0008:
		return "IoctlKeyboardSetIndicators"
	case 0x000B0020:
		return "IoctlKeyboardQueryTypematic"
	case 0x000B0040:
		return "IoctlKeyboardQueryIndicators"
	case 0x000B0080:
		return "IoctlKeyboardQueryIndicatorTranslation"
	case 0x000B0100:
		return "IoctlKeyboardInsertData"
	case 0x000B1000:
		return "IoctlKeyboardQueryImeStatus"
	case 0x000B1004:
		return "IoctlKeyboardSetImeStatus"
	case 0x002B0004:
		return "IoctlModemGetPassthrough"
	case 0x002B0008:
		return "IoctlModemSetPassthrough"
	case 0x002B000C:
		return "IoctlModemSetDleMonitoring"
	case 0x002B0010:
		return "IoctlModemGetDle"
	case 0x002B0014:
		return "IoctlModemSetDleShielding"
	case 0x002B0018:
		return "IoctlModemStopWaveReceive"
	case 0x002B001C:
		return "IoctlModemSendMessage"
	case 0x002B0020:
		return "IoctlModemGetMessage"
	case 0x002B0024:
		return "IoctlModemSendGetMessage"
	case 0x002B0028:
		return "IoctlModemSendLoopbackMessage"
	case 0x002B002C:
		return "IoctlModemCheckForModem"
	case 0x002B0030:
		return "IoctlModemSetMinPower"
	case 0x002B0034:
		return "IoctlModemWatchForResume"
	case 0x002B0038:
		return "IoctlCancelGetSendMessage"
	case 0x002B003C:
		return "IoctlSetServerState"
	case 0x000F0000:
		return "IoctlMouseQueryAttributes"
	case 0x000F0004:
		return "IoctlMouseInsertData"
	case 0x00170034:
		return "IoctlNdisReserved5"
	case 0x00178038:
		return "IoctlNdisReserved6"
	case 0x00160004:
		return "IoctlParQueryInformation"
	case 0x00160008:
		return "IoctlParSetInformation"
	case 0x0016000C:
		return "IoctlParQueryDeviceId"
	case 0x00160010:
		return "IoctlParQueryDeviceIdSize"
	case 0x00160014:
		return "IoctlIeee1284GetMode"
	case 0x00160018:
		return "IoctlIeee1284Negotiate"
	case 0x0016001C:
		return "IoctlParSetWriteAddress"
	case 0x00160020:
		return "IoctlParSetReadAddress"
	case 0x00160024:
		return "IoctlParGetDeviceCaps"
	case 0x00160028:
		return "IoctlParGetDefaultModes"
	case 0x0016002C:
		return "IoctlParPing"
	case 0x00160030:
		return "IoctlParQueryRawDeviceId"
	case 0x00160034:
		return "IoctlParEcpHostRecovery"
	case 0x00160038:
		return "IoctlParGetReadAddress"
	case 0x0016003C:
		return "IoctlParGetWriteAddress"
	case 0x00160050:
		return "IoctlParTest"
	case 0x00160054:
		return "IoctlParIsPortFree"
	case 0x00160058:
		return "IoctlParQueryLocation"
	case 0x0004D004:
		return "IoctlScsiPassThrough"
	case 0x0004D008:
		return "IoctlScsiMiniport"
	case 0x0004100C:
		return "IoctlScsiGetInquiryData"
	case 0x00041010:
		return "IoctlScsiGetCapabilities"
	case 0x0004D014:
		return "IoctlScsiPassThroughDirect"
	case 0x00041018:
		return "IoctlScsiGetAddress"
	case 0x0004101C:
		return "IoctlScsiRescanBus"
	case 0x00041020:
		return "IoctlScsiGetDumpPointers"
	case 0x00041024:
		return "IoctlScsiFreeDumpPointers"
	case 0x0004D028:
		return "IoctlIdePassThrough"
	case 0x0004D02C:
		return "IoctlAtaPassThrough"
	case 0x0004D030:
		return "IoctlAtaPassThroughDirect"
	case 0x0004D034:
		return "IoctlAtaMiniport"
	case 0x0004D038:
		return "IoctlMiniportProcessServiceIrp"
	case 0x0004D03C:
		return "IoctlMpioPassThroughPath"
	case 0x0004D040:
		return "IoctlMpioPassThroughPathDirect"
	case 0x001B0004:
		return "IoctlSerialSetBaudRate"
	case 0x001B0008:
		return "IoctlSerialSetQueueSize"
	case 0x001B000C:
		return "IoctlSerialSetLineControl"
	case 0x001B0010:
		return "IoctlSerialSetBreakOn"
	case 0x001B0014:
		return "IoctlSerialSetBreakOff"
	case 0x001B0018:
		return "IoctlSerialImmediateChar"
	case 0x001B001C:
		return "IoctlSerialSetTimeouts"
	case 0x001B0020:
		return "IoctlSerialGetTimeouts"
	case 0x001B0024:
		return "IoctlSerialSetDtr"
	case 0x001B0028:
		return "IoctlSerialClrDtr"
	case 0x001B002C:
		return "IoctlSerialResetDevice"
	case 0x001B0030:
		return "IoctlSerialSetRts"
	case 0x001B0034:
		return "IoctlSerialClrRts"
	case 0x001B0038:
		return "IoctlSerialSetXoff"
	case 0x001B003C:
		return "IoctlSerialSetXon"
	case 0x001B0040:
		return "IoctlSerialGetWaitMask"
	case 0x001B0044:
		return "IoctlSerialSetWaitMask"
	case 0x001B0048:
		return "IoctlSerialWaitOnMask"
	case 0x001B004C:
		return "IoctlSerialPurge"
	case 0x001B0050:
		return "IoctlSerialGetBaudRate"
	case 0x001B0054:
		return "IoctlSerialGetLineControl"
	case 0x001B0058:
		return "IoctlSerialGetChars"
	case 0x001B005C:
		return "IoctlSerialSetChars"
	case 0x001B0060:
		return "IoctlSerialGetHandflow"
	case 0x001B0064:
		return "IoctlSerialSetHandflow"
	case 0x001B0068:
		return "IoctlSerialGetModemstatus"
	case 0x001B006C:
		return "IoctlSerialGetCommstatus"
	case 0x001B0070:
		return "IoctlSerialXoffCounter"
	case 0x001B0074:
		return "IoctlSerialGetProperties"
	case 0x001B0078:
		return "IoctlSerialGetDtrrts"
	case 0x001B007C:
		return "IoctlSerialLsrmstInsert"
	case 0x00370200:
		return "IoctlSerenumExposeHardware"
	case 0x00370204:
		return "IoctlSerenumRemoveHardware"
	case 0x00370208:
		return "IoctlSerenumPortDesc"
	case 0x0037020C:
		return "IoctlSerenumGetPortName"
	case 0x001B0080:
		return "IoctlSerialConfigSize"
	case 0x001B0084:
		return "IoctlSerialGetCommconfig"
	case 0x001B0088:
		return "IoctlSerialSetCommconfig"
	case 0x001B008C:
		return "IoctlSerialGetStats"
	case 0x001B0090:
		return "IoctlSerialClearStats"
	case 0x001B0094:
		return "IoctlSerialGetModemControl"
	case 0x001B0098:
		return "IoctlSerialSetModemControl"
	case 0x001B009C:
		return "IoctlSerialSetFifoControl"
	case 0x001B0004:
		return "IoctlSerialInternalDoWaitWake"
	case 0x001B0008:
		return "IoctlSerialInternalCancelWaitWake"
	case 0x001B000C:
		return "IoctlSerialInternalBasicSettings"
	case 0x001B0010:
		return "IoctlSerialInternalRestoreSettings"
	case 0x002D4800:
		return "IoctlStorageCheckVerify"
	case 0x002D0800:
		return "IoctlStorageCheckVerify2"
	case 0x002D4804:
		return "IoctlStorageMediaRemoval"
	case 0x002D4808:
		return "IoctlStorageEjectMedia"
	case 0x002D480C:
		return "IoctlStorageLoadMedia"
	case 0x002D080C:
		return "IoctlStorageLoadMedia2"
	case 0x002D4810:
		return "IoctlStorageReserve"
	case 0x002D4814:
		return "IoctlStorageRelease"
	case 0x002D4818:
		return "IoctlStorageFindNewDevices"
	case 0x002D0940:
		return "IoctlStorageEjectionControl"
	case 0x002D0944:
		return "IoctlStorageMcnControl"
	case 0x002D0C00:
		return "IoctlStorageGetMediaTypes"
	case 0x002D0C04:
		return "IoctlStorageGetMediaTypesEx"
	case 0x002D0C10:
		return "IoctlStorageGetMediaSerialNumber"
	case 0x002D0C14:
		return "IoctlStorageGetHotplugInfo"
	case 0x002DCC18:
		return "IoctlStorageSetHotplugInfo"
	case 0x002D5000:
		return "IoctlStorageResetBus"
	case 0x002D5004:
		return "IoctlStorageResetDevice"
	case 0x002D5014:
		return "IoctlStorageBreakReservation"
	case 0x002D5018:
		return "IoctlStoragePersistentReserveIn"
	case 0x002DD01C:
		return "IoctlStoragePersistentReserveOut"
	case 0x002D1080:
		return "IoctlStorageGetDeviceNumber"
	case 0x002D1100:
		return "IoctlStoragePredictFailure"
	case 0x002D5140:
		return "IoctlStorageReadCapacity"
	case 0x002D1400:
		return "IoctlStorageQueryProperty"
	case 0x002D9404:
		return "IoctlStorageManageDataSetAttributes"
	case 0x002D5800:
		return "IoctlStorageGetBcProperties"
	case 0x002DD804:
		return "IoctlStorageAllocateBcStream"
	case 0x002DD808:
		return "IoctlStorageFreeBcStream"
	case 0x002D1880:
		return "IoctlStorageCheckPriorityHintSupport"
	case 0x002DD000:
		return "ObsoleteIoctlStorageResetBus"
	case 0x002DD004:
		return "ObsoleteIoctlStorageResetDevice"
	case 0x001FC000:
		return "IoctlTapeErase"
	case 0x001F4004:
		return "IoctlTapePrepare"
	case 0x001FC008:
		return "IoctlTapeWriteMarks"
	case 0x001F400C:
		return "IoctlTapeGetPosition"
	case 0x001F4010:
		return "IoctlTapeSetPosition"
	case 0x001F4014:
		return "IoctlTapeGetDriveParams"
	case 0x001FC018:
		return "IoctlTapeSetDriveParams"
	case 0x001F401C:
		return "IoctlTapeGetMediaParams"
	case 0x001F4020:
		return "IoctlTapeSetMediaParams"
	case 0x001F4024:
		return "IoctlTapeGetStatus"
	case 0x001FC028:
		return "IoctlTapeCreatePartition"
	case 0x001F4804:
		return "IoctlTapeMediaRemoval"
	case 0x001F4808:
		return "IoctlTapeEjectMedia"
	case 0x001F480C:
		return "IoctlTapeLoadMedia"
	case 0x001F4810:
		return "IoctlTapeReserve"
	case 0x001F4814:
		return "IoctlTapeRelease"
	case 0x001F4800:
		return "IoctlTapeCheckVerify"
	case 0x00074818:
		return "IoctlTapeFindNewDevices"
	case 0x00560000:
		return "IoctlVolumeGetVolumeDiskExtents"
	case 0x0056C008:
		return "IoctlVolumeOnline"
	case 0x0056C00C:
		return "IoctlVolumeOffline"
	case 0x00560030:
		return "IoctlVolumeIsClustered"
	case 0x00560038:
		return "IoctlVolumeGetGptAttributes"
	case 0x00560004:
		return "IoctlVolumeSupportsOnlineOffline"
	case 0x00560010:
		return "IoctlVolumeIsOffline"
	case 0x00560014:
		return "IoctlVolumeIsIoCapable"
	case 0x00560018:
		return "IoctlVolumeQueryFailoverSet"
	case 0x0056001C:
		return "IoctlVolumeQueryVolumeNumber"
	case 0x00560020:
		return "IoctlVolumeLogicalToPhysical"
	case 0x00560024:
		return "IoctlVolumePhysicalToLogical"
	case 0x00560028:
		return "IoctlVolumeIsPartition"
	case 0x0056402E:
		return "IoctlVolumeReadPlex"
	case 0x00560034:
		return "IoctlVolumeSetGptAttributes"
	case 0x0056403C:
		return "IoctlVolumeGetBcProperties"
	case 0x0056C040:
		return "IoctlVolumeAllocateBcStream"
	case 0x0056C044:
		return "IoctlVolumeFreeBcStream"
	case 0x00560048:
		return "IoctlVolumeIsDynamic"
	case 0x0056C04C:
		return "IoctlVolumePrepareForCriticalIo"
	case 0x00564052:
		return "IoctlVolumeQueryAllocationHint"
	case 0x00560054:
		return "IoctlVolumeUpdateProperties"
	case 0x00564058:
		return "IoctlVolumeQueryMinimumShrinkSize"
	case 0x0056C05C:
		return "IoctlVolumePrepareForShrink"
	case 0x00454000:
		return "IoctlPmiGetCapabilities"
	case 0x00454004:
		return "IoctlPmiGetConfiguration"
	case 0x00458008:
		return "IoctlPmiSetConfiguration"
	case 0x0045400C:
		return "IoctlPmiGetMeasurement"
	case 0x0045C010:
		return "IoctlPmiRegisterEventNotify"
	case 0x00442000:
		return "IoctlBiometricVendor"
	default:
		return "unknown IoctlKind " + fmt.Sprint(k)
	}
}

func (k IoctlKind) Elements() []IoctlKind {
	return []IoctlKind{
		MonitorIoctlEnableMonitor,
		MonitorIoctlDisableMonitor,
		IoctlInternalKeyboardConnect,
		IoctlInternalKeyboardDisconnect,
		IoctlInternalKeyboardEnable,
		IoctlInternalKeyboardDisable,
		IoctlInternalMouseConnect,
		IoctlInternalMouseDisconnect,
		IoctlInternalMouseEnable,
		IoctlInternalMouseDisable,
		IoctlDoKernelmodeSamples,
		IoctlRegisterCallback,
		IoctlUnregisterCallback,
		IoctlGetCallbackVersion,
		IoctlGetVersion,
		IoctlReset,
		IoctlReadwriteMailbox,
		IoctlMailboxWait,
		IoctlReadDma,
		IoctlWriteDma,
		IoctlAcpiAsyncEvalMethod,
		IoctlAcpiEvalMethod,
		IoctlAcpiAcquireGlobalLock,
		IoctlAcpiReleaseGlobalLock,
		IoctlAcpiEvalMethodEx,
		IoctlAcpiAsyncEvalMethodEx,
		IoctlAcpiEnumChildren,
		IoctlAvcUpdateVirtualSubunitInfo,
		IoctlAvcRemoveVirtualSubunitInfo,
		IoctlAvcBusReset,
		IoctlDot4CreateSocket,
		IoctlDot4DestroySocket,
		IoctlDot4WaitForChannel,
		IoctlDot4OpenChannel,
		IoctlDot4CloseChannel,
		IoctlDot4Read,
		IoctlDot4Write,
		IoctlDot4AddActivityBroadcast,
		IoctlDot4RemoveActivityBroadcast,
		IoctlDot4WaitActivityBroadcast,
		IoctlMpdsmRegister,
		IoctlMpdsmDeregister,
		IoctlMountdevQueryUniqueId,
		IoctlMountdevQuerySuggestedLinkName,
		IoctlMountdevLinkCreated,
		IoctlMountdevLinkDeleted,
		IoctlMountdevQueryStableGuid,
		IoctlMountmgrCreatePoint,
		IoctlMountmgrDeletePoints,
		IoctlMountmgrQueryPoints,
		IoctlMountmgrDeletePointsDbonly,
		IoctlMountmgrNextDriveLetter,
		IoctlMountmgrAutoDlAssignments,
		IoctlMountmgrVolumeMountPointCreated,
		IoctlMountmgrVolumeMountPointDeleted,
		IoctlMountmgrChangeNotify,
		IoctlMountmgrKeepLinksWhenOffline,
		IoctlMountmgrCheckUnprocessedVolumes,
		IoctlMountmgrVolumeArrivalNotification,
		IoctlMountdevQueryDeviceName,
		IoctlMountmgrQueryDosVolumePath,
		IoctlMountmgrQueryDosVolumePaths,
		IoctlMountmgrScrubRegistry,
		IoctlMountmgrQueryAutoMount,
		IoctlMountmgrSetAutoMount,
		IoctlMountmgrBootDlAssignment,
		IoctlMountmgrTracelogCache,
		IoctlAvioAllocateStream,
		IoctlAvioFreeStream,
		IoctlAvioModifyStream,
		NlbIoctlRegisterHook,
		NlbPublicIoctlClientStickinessNotify,
		IoctlGetTupleData,
		IoctlSocketInformation,
		IoctlPcmciaHideDevice,
		IoctlPcmciaRevealDevice,
		IoctlSdSubmitRequest,
		FsctlRequestOplockLevel1,
		FsctlRequestOplockLevel2,
		FsctlRequestBatchOplock,
		FsctlOplockBreakAcknowledge,
		FsctlOpbatchAckClosePending,
		FsctlOplockBreakNotify,
		FsctlLockVolume,
		FsctlUnlockVolume,
		FsctlDismountVolume,
		FsctlIsVolumeMounted,
		FsctlIsPathnameValid,
		FsctlMarkVolumeDirty,
		FsctlQueryRetrievalPointers,
		FsctlGetCompression,
		FsctlSetCompression,
		FsctlSetBootloaderAccessed,
		FsctlOplockBreakAckNo2,
		FsctlInvalidateVolumes,
		FsctlQueryFatBpb,
		FsctlRequestFilterOplock,
		FsctlFilesystemGetStatistics,
		FsctlGetNtfsVolumeData,
		FsctlGetNtfsFileRecord,
		FsctlGetVolumeBitmap,
		FsctlGetRetrievalPointers,
		FsctlMoveFile,
		FsctlIsVolumeDirty,
		FsctlAllowExtendedDasdIo,
		FsctlFindFilesBySid,
		FsctlSetObjectId,
		FsctlGetObjectId,
		FsctlDeleteObjectId,
		FsctlSetReparsePoint,
		FsctlGetReparsePoint,
		FsctlDeleteReparsePoint,
		FsctlEnumUsnData,
		FsctlSecurityIdCheck,
		FsctlReadUsnJournal,
		FsctlSetObjectIdExtended,
		FsctlCreateOrGetObjectId,
		FsctlSetSparse,
		FsctlSetZeroData,
		FsctlQueryAllocatedRanges,
		FsctlEnableUpgrade,
		FsctlSetEncryption,
		FsctlEncryptionFsctlIo,
		FsctlWriteRawEncrypted,
		FsctlReadRawEncrypted,
		FsctlCreateUsnJournal,
		FsctlReadFileUsnData,
		FsctlWriteUsnCloseRecord,
		FsctlExtendVolume,
		FsctlQueryUsnJournal,
		FsctlDeleteUsnJournal,
		FsctlMarkHandle,
		FsctlSisCopyfile,
		FsctlSisLinkFiles,
		FsctlRecallFile,
		FsctlReadFromPlex,
		FsctlFilePrefetch,
		FsctlMakeMediaCompatible,
		FsctlSetDefectManagement,
		FsctlQuerySparingInfo,
		FsctlQueryOnDiskVolumeInfo,
		FsctlSetVolumeCompressionState,
		FsctlTxfsModifyRm,
		FsctlTxfsQueryRmInformation,
		FsctlTxfsRollforwardRedo,
		FsctlTxfsRollforwardUndo,
		FsctlTxfsStartRm,
		FsctlTxfsShutdownRm,
		FsctlTxfsReadBackupInformation,
		FsctlTxfsWriteBackupInformation,
		FsctlTxfsCreateSecondaryRm,
		FsctlTxfsGetMetadataInfo,
		FsctlTxfsGetTransactedVersion,
		FsctlTxfsSavepointInformation,
		FsctlTxfsCreateMiniversion,
		FsctlTxfsTransactionActive,
		FsctlSetZeroOnDeallocation,
		FsctlSetRepair,
		FsctlGetRepair,
		FsctlWaitForRepair,
		FsctlInitiateRepair,
		FsctlCscInternal,
		FsctlShrinkVolume,
		FsctlSetShortNameBehavior,
		FsctlDfsrSetGhostHandleState,
		FsctlTxfsListTransactions,
		FsctlQueryPagefileEncryption,
		FsctlResetVolumeAllocationHints,
		FsctlQueryDependentVolume,
		FsctlSdGlobalChange,
		FsctlTxfsReadBackupInformation2,
		FsctlLookupStreamFromCluster,
		FsctlTxfsWriteBackupInformation2,
		FsctlFileTypeNotification,
		FsctlGetBootAreaInfo,
		FsctlGetRetrievalPointerBase,
		FsctlSetPersistentVolumeState,
		FsctlQueryPersistentVolumeState,
		FsctlRequestOplock,
		FsctlCsvTunnelRequest,
		FsctlIsCsvFile,
		FsctlQueryFileSystemRecognition,
		FsctlCsvGetVolumePathName,
		FsctlCsvGetVolumeNameForVolumeMountPoint,
		FsctlCsvGetVolumePathNamesForVolumeName,
		FsctlIsFileOnCsvVolume,
		FsctlLmrGetLinkTrackingInformation,
		FsctlLmrSetLinkTrackingInformation,
		IoctlLmrAreFileObjectsOnSameServer,
		FsctlPipeAssignEvent,
		FsctlPipeDisconnect,
		FsctlPipeListen,
		FsctlPipePeek,
		FsctlPipeQueryEvent,
		FsctlPipeTransceive,
		FsctlPipeWait,
		FsctlPipeImpersonate,
		FsctlPipeSetClientProcess,
		FsctlPipeQueryClientProcess,
		FsctlPipeGetPipeAttribute,
		FsctlPipeSetPipeAttribute,
		FsctlPipeGetConnectionAttribute,
		FsctlPipeSetConnectionAttribute,
		FsctlPipeGetHandleAttribute,
		FsctlPipeSetHandleAttribute,
		FsctlPipeFlush,
		FsctlPipeInternalRead,
		FsctlPipeInternalWrite,
		FsctlPipeInternalTransceive,
		FsctlPipeInternalReadOvflow,
		FsctlMailslotPeek,
		IoctlRedirQueryPath,
		IoctlRedirQueryPathEx,
		IoctlVolsnapFlushAndHoldWrites,
		IoctlInternalParallelPortAllocate,
		IoctlInternalGetParallelPortInfo,
		IoctlInternalParallelConnectInterrupt,
		IoctlInternalParallelDisconnectInterrupt,
		IoctlInternalReleaseParallelPortInfo,
		IoctlInternalGetMoreParallelPortInfo,
		IoctlInternalParchipConnect,
		IoctlInternalParallelSetChipMode,
		IoctlInternalParallelClearChipMode,
		IoctlInternalGetParallelPnpInfo,
		IoctlInternalInit12843Bus,
		IoctlInternalSelectDevice,
		IoctlInternalDeselectDevice,
		IoctlInternalGetParportFdo,
		IoctlInternalParclassConnect,
		IoctlInternalParclassDisconnect,
		IoctlInternalDisconnectIdle,
		IoctlInternalLockPort,
		IoctlInternalUnlockPort,
		IoctlInternalParallelPortFree,
		IoctlInternalPardot3Connect,
		IoctlInternalPardot3Disconnect,
		IoctlInternalPardot3Reset,
		IoctlInternalPardot3Signal,
		IoctlInternalRegisterForRemovalRelations,
		IoctlInternalUnregisterForRemovalRelations,
		IoctlInternalLockPortNoSelect,
		IoctlInternalUnlockPortNoDeselect,
		IoctlInternalDisableEndOfChainBusRescan,
		IoctlInternalEnableEndOfChainBusRescan,
		IoctlWpdMessageReadwriteAccess,
		IoctlWpdMessageReadAccess,
		IoctlScsiscanCmd,
		IoctlScsiscanLockdevice,
		IoctlScsiscanUnlockdevice,
		IoctlScsiscanSetTimeout,
		IoctlScsiscanGetInfo,
		IoctlSwenumInstallInterface,
		IoctlSwenumRemoveInterface,
		IoctlSwenumGetBusId,
		IoctlCancelIo,
		IoctlWaitOnDeviceEvent,
		IoctlReadRegisters,
		IoctlWriteRegisters,
		IoctlGetChannelAlignRqst,
		IoctlGetDeviceDescriptor,
		IoctlResetPipe,
		IoctlGetUsbDescriptor,
		IoctlSendUsbRequest,
		IoctlGetPipeConfiguration,
		IoctlSetTimeout,
		IoctlEhstorDeviceSetAuthzState,
		IoctlEhstorDeviceGetAuthzState,
		IoctlEhstorDeviceSiloCommand,
		IoctlEhstorDeviceEnumeratePdos,
		IoctlKsProperty,
		IoctlKsEnableEvent,
		IoctlKsDisableEvent,
		IoctlKsMethod,
		IoctlKsWriteStream,
		IoctlKsReadStream,
		IoctlKsResetState,
		IoctlKsHandshake,
		IoctlInternalI8042HookKeyboard,
		IoctlInternalI8042HookMouse,
		IoctlInternalI8042KeyboardWriteBuffer,
		IoctlInternalI8042MouseWriteBuffer,
		IoctlInternalI8042ControllerWriteBuffer,
		IoctlInternalI8042KeyboardStartInformation,
		IoctlInternalI8042MouseStartInformation,
		IoctlBeepSet,
		IoctlCdromUnloadDriver,
		IoctlCdromReadToc,
		IoctlCdromSeekAudioMsf,
		IoctlCdromStopAudio,
		IoctlCdromPauseAudio,
		IoctlCdromResumeAudio,
		IoctlCdromGetVolume,
		IoctlCdromPlayAudioMsf,
		IoctlCdromSetVolume,
		IoctlCdromReadQChannel,
		IoctlCdromGetControl,
		ObsoleteIoctlCdromGetControl,
		IoctlCdromGetLastSession,
		IoctlCdromRawRead,
		IoctlCdromDiskType,
		IoctlCdromGetDriveGeometry,
		IoctlCdromGetDriveGeometryEx,
		IoctlCdromReadTocEx,
		IoctlCdromGetConfiguration,
		IoctlCdromExclusiveAccess,
		IoctlCdromSetSpeed,
		IoctlCdromGetInquiryData,
		IoctlCdromEnableStreaming,
		IoctlCdromSendOpcInformation,
		IoctlCdromGetPerformance,
		IoctlCdromCheckVerify,
		IoctlCdromMediaRemoval,
		IoctlCdromEjectMedia,
		IoctlCdromLoadMedia,
		IoctlCdromReserve,
		IoctlCdromRelease,
		IoctlCdromFindNewDevices,
		IoctlCdromSimbad,
		IoctlDvdStartSession,
		IoctlDvdReadKey,
		IoctlDvdSendKey,
		IoctlDvdEndSession,
		IoctlDvdSetReadAhead,
		IoctlDvdGetRegion,
		IoctlDvdSendKey2,
		IoctlAacsReadMediaKeyBlockSize,
		IoctlAacsReadMediaKeyBlock,
		IoctlAacsStartSession,
		IoctlAacsEndSession,
		IoctlAacsSendCertificate,
		IoctlAacsGetCertificate,
		IoctlAacsGetChallengeKey,
		IoctlAacsSendChallengeKey,
		IoctlAacsReadVolumeId,
		IoctlAacsReadSerialNumber,
		IoctlAacsReadMediaId,
		IoctlAacsReadBindingNonce,
		IoctlAacsGenerateBindingNonce,
		IoctlDvdReadStructure,
		IoctlStorageSetReadAhead,
		IoctlChangerGetParameters,
		IoctlChangerGetStatus,
		IoctlChangerGetProductData,
		IoctlChangerSetAccess,
		IoctlChangerGetElementStatus,
		IoctlChangerInitializeElementStatus,
		IoctlChangerSetPosition,
		IoctlChangerExchangeMedium,
		IoctlChangerMoveMedium,
		IoctlChangerReinitializeTransport,
		IoctlChangerQueryVolumeTags,
		IoctlDiskGetDriveGeometry,
		IoctlDiskGetPartitionInfo,
		IoctlDiskSetPartitionInfo,
		IoctlDiskGetDriveLayout,
		IoctlDiskSetDriveLayout,
		IoctlDiskVerify,
		IoctlDiskFormatTracks,
		IoctlDiskReassignBlocks,
		IoctlDiskPerformance,
		IoctlDiskIsWritable,
		IoctlDiskLogging,
		IoctlDiskFormatTracksEx,
		IoctlDiskHistogramStructure,
		IoctlDiskHistogramData,
		IoctlDiskHistogramReset,
		IoctlDiskRequestStructure,
		IoctlDiskRequestData,
		IoctlDiskPerformanceOff,
		IoctlDiskControllerNumber,
		SmartGetVersion,
		SmartSendDriveCommand,
		SmartRcvDriveData,
		IoctlDiskGetPartitionInfoEx,
		IoctlDiskSetPartitionInfoEx,
		IoctlDiskGetDriveLayoutEx,
		IoctlDiskSetDriveLayoutEx,
		IoctlDiskCreateDisk,
		IoctlDiskGetLengthInfo,
		IoctlDiskGetDriveGeometryEx,
		IoctlDiskReassignBlocksEx,
		IoctlDiskUpdateDriveSize,
		IoctlDiskGrowPartition,
		IoctlDiskGetCacheInformation,
		IoctlDiskSetCacheInformation,
		IoctlDiskGetWriteCacheState,
		ObsoleteDiskGetWriteCacheState,
		IoctlDiskDeleteDriveLayout,
		IoctlDiskUpdateProperties,
		IoctlDiskFormatDrive,
		IoctlDiskSenseDevice,
		IoctlDiskGetCacheSetting,
		IoctlDiskSetCacheSetting,
		IoctlDiskCopyData,
		IoctlDiskInternalSetVerify,
		IoctlDiskInternalClearVerify,
		IoctlDiskInternalSetNotify,
		IoctlDiskCheckVerify,
		IoctlDiskMediaRemoval,
		IoctlDiskEjectMedia,
		IoctlDiskLoadMedia,
		IoctlDiskReserve,
		IoctlDiskRelease,
		IoctlDiskFindNewDevices,
		IoctlDiskGetMediaTypes,
		IoctlDiskGetPartitionAttributes,
		IoctlDiskSetPartitionAttributes,
		IoctlDiskGetDiskAttributes,
		IoctlDiskSetDiskAttributes,
		IoctlDiskIsClustered,
		IoctlDiskGetSanSettings,
		IoctlDiskSetSanSettings,
		IoctlDiskGetSnapshotInfo,
		IoctlDiskSetSnapshotInfo,
		IoctlDiskResetSnapshotInfo,
		IoctlDiskSimbad,
		FtSecondaryRead,
		FtPrimaryRead,
		FtBalancedReadMode,
		FtSyncRedundantCopy,
		FtInitializeSet,
		FtRegenerate,
		FtConfigure,
		FtVerify,
		FtSequentialWriteMode,
		FtParallelWriteMode,
		FtQuerySetState,
		FtClusterSetMemberState,
		FtClusterGetMemberState,
		FtEnumerateLogicalDisks,
		FtCreateLogicalDisk,
		FtBreakLogicalDisk,
		FtQueryLogicalDiskInformation,
		FtOrphanLogicalDiskMember,
		FtReplaceLogicalDiskMember,
		FtQueryNtDeviceNameForLogicalDisk,
		FtInitializeLogicalDisk,
		FtQueryDriveLetterForLogicalDisk,
		FtCheckIo,
		FtSetDriveLetterForLogicalDisk,
		FtQueryNtDeviceNameForPartition,
		FtChangeNotify,
		FtStopSyncOperations,
		FtQueryLogicalDiskId,
		FtCreatePartitionLogicalDisk,
		IoctlKeyboardQueryAttributes,
		IoctlKeyboardSetTypematic,
		IoctlKeyboardSetIndicators,
		IoctlKeyboardQueryTypematic,
		IoctlKeyboardQueryIndicators,
		IoctlKeyboardQueryIndicatorTranslation,
		IoctlKeyboardInsertData,
		IoctlKeyboardQueryImeStatus,
		IoctlKeyboardSetImeStatus,
		IoctlModemGetPassthrough,
		IoctlModemSetPassthrough,
		IoctlModemSetDleMonitoring,
		IoctlModemGetDle,
		IoctlModemSetDleShielding,
		IoctlModemStopWaveReceive,
		IoctlModemSendMessage,
		IoctlModemGetMessage,
		IoctlModemSendGetMessage,
		IoctlModemSendLoopbackMessage,
		IoctlModemCheckForModem,
		IoctlModemSetMinPower,
		IoctlModemWatchForResume,
		IoctlCancelGetSendMessage,
		IoctlSetServerState,
		IoctlMouseQueryAttributes,
		IoctlMouseInsertData,
		IoctlNdisReserved5,
		IoctlNdisReserved6,
		IoctlParQueryInformation,
		IoctlParSetInformation,
		IoctlParQueryDeviceId,
		IoctlParQueryDeviceIdSize,
		IoctlIeee1284GetMode,
		IoctlIeee1284Negotiate,
		IoctlParSetWriteAddress,
		IoctlParSetReadAddress,
		IoctlParGetDeviceCaps,
		IoctlParGetDefaultModes,
		IoctlParPing,
		IoctlParQueryRawDeviceId,
		IoctlParEcpHostRecovery,
		IoctlParGetReadAddress,
		IoctlParGetWriteAddress,
		IoctlParTest,
		IoctlParIsPortFree,
		IoctlParQueryLocation,
		IoctlScsiPassThrough,
		IoctlScsiMiniport,
		IoctlScsiGetInquiryData,
		IoctlScsiGetCapabilities,
		IoctlScsiPassThroughDirect,
		IoctlScsiGetAddress,
		IoctlScsiRescanBus,
		IoctlScsiGetDumpPointers,
		IoctlScsiFreeDumpPointers,
		IoctlIdePassThrough,
		IoctlAtaPassThrough,
		IoctlAtaPassThroughDirect,
		IoctlAtaMiniport,
		IoctlMiniportProcessServiceIrp,
		IoctlMpioPassThroughPath,
		IoctlMpioPassThroughPathDirect,
		IoctlSerialSetBaudRate,
		IoctlSerialSetQueueSize,
		IoctlSerialSetLineControl,
		IoctlSerialSetBreakOn,
		IoctlSerialSetBreakOff,
		IoctlSerialImmediateChar,
		IoctlSerialSetTimeouts,
		IoctlSerialGetTimeouts,
		IoctlSerialSetDtr,
		IoctlSerialClrDtr,
		IoctlSerialResetDevice,
		IoctlSerialSetRts,
		IoctlSerialClrRts,
		IoctlSerialSetXoff,
		IoctlSerialSetXon,
		IoctlSerialGetWaitMask,
		IoctlSerialSetWaitMask,
		IoctlSerialWaitOnMask,
		IoctlSerialPurge,
		IoctlSerialGetBaudRate,
		IoctlSerialGetLineControl,
		IoctlSerialGetChars,
		IoctlSerialSetChars,
		IoctlSerialGetHandflow,
		IoctlSerialSetHandflow,
		IoctlSerialGetModemstatus,
		IoctlSerialGetCommstatus,
		IoctlSerialXoffCounter,
		IoctlSerialGetProperties,
		IoctlSerialGetDtrrts,
		IoctlSerialLsrmstInsert,
		IoctlSerenumExposeHardware,
		IoctlSerenumRemoveHardware,
		IoctlSerenumPortDesc,
		IoctlSerenumGetPortName,
		IoctlSerialConfigSize,
		IoctlSerialGetCommconfig,
		IoctlSerialSetCommconfig,
		IoctlSerialGetStats,
		IoctlSerialClearStats,
		IoctlSerialGetModemControl,
		IoctlSerialSetModemControl,
		IoctlSerialSetFifoControl,
		IoctlSerialInternalDoWaitWake,
		IoctlSerialInternalCancelWaitWake,
		IoctlSerialInternalBasicSettings,
		IoctlSerialInternalRestoreSettings,
		IoctlStorageCheckVerify,
		IoctlStorageCheckVerify2,
		IoctlStorageMediaRemoval,
		IoctlStorageEjectMedia,
		IoctlStorageLoadMedia,
		IoctlStorageLoadMedia2,
		IoctlStorageReserve,
		IoctlStorageRelease,
		IoctlStorageFindNewDevices,
		IoctlStorageEjectionControl,
		IoctlStorageMcnControl,
		IoctlStorageGetMediaTypes,
		IoctlStorageGetMediaTypesEx,
		IoctlStorageGetMediaSerialNumber,
		IoctlStorageGetHotplugInfo,
		IoctlStorageSetHotplugInfo,
		IoctlStorageResetBus,
		IoctlStorageResetDevice,
		IoctlStorageBreakReservation,
		IoctlStoragePersistentReserveIn,
		IoctlStoragePersistentReserveOut,
		IoctlStorageGetDeviceNumber,
		IoctlStoragePredictFailure,
		IoctlStorageReadCapacity,
		IoctlStorageQueryProperty,
		IoctlStorageManageDataSetAttributes,
		IoctlStorageGetBcProperties,
		IoctlStorageAllocateBcStream,
		IoctlStorageFreeBcStream,
		IoctlStorageCheckPriorityHintSupport,
		ObsoleteIoctlStorageResetBus,
		ObsoleteIoctlStorageResetDevice,
		IoctlTapeErase,
		IoctlTapePrepare,
		IoctlTapeWriteMarks,
		IoctlTapeGetPosition,
		IoctlTapeSetPosition,
		IoctlTapeGetDriveParams,
		IoctlTapeSetDriveParams,
		IoctlTapeGetMediaParams,
		IoctlTapeSetMediaParams,
		IoctlTapeGetStatus,
		IoctlTapeCreatePartition,
		IoctlTapeMediaRemoval,
		IoctlTapeEjectMedia,
		IoctlTapeLoadMedia,
		IoctlTapeReserve,
		IoctlTapeRelease,
		IoctlTapeCheckVerify,
		IoctlTapeFindNewDevices,
		IoctlVolumeGetVolumeDiskExtents,
		IoctlVolumeOnline,
		IoctlVolumeOffline,
		IoctlVolumeIsClustered,
		IoctlVolumeGetGptAttributes,
		IoctlVolumeSupportsOnlineOffline,
		IoctlVolumeIsOffline,
		IoctlVolumeIsIoCapable,
		IoctlVolumeQueryFailoverSet,
		IoctlVolumeQueryVolumeNumber,
		IoctlVolumeLogicalToPhysical,
		IoctlVolumePhysicalToLogical,
		IoctlVolumeIsPartition,
		IoctlVolumeReadPlex,
		IoctlVolumeSetGptAttributes,
		IoctlVolumeGetBcProperties,
		IoctlVolumeAllocateBcStream,
		IoctlVolumeFreeBcStream,
		IoctlVolumeIsDynamic,
		IoctlVolumePrepareForCriticalIo,
		IoctlVolumeQueryAllocationHint,
		IoctlVolumeUpdateProperties,
		IoctlVolumeQueryMinimumShrinkSize,
		IoctlVolumePrepareForShrink,
		IoctlPmiGetCapabilities,
		IoctlPmiGetConfiguration,
		IoctlPmiSetConfiguration,
		IoctlPmiGetMeasurement,
		IoctlPmiRegisterEventNotify,
		IoctlBiometricVendor,
	}
}
