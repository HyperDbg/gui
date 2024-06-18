package symbol

type (
	InterfaceNtdll interface {
		NtAcceptConnectPort() (ok bool)
		NtAccessCheck() (ok bool)
		NtAccessCheckAndAuditAlarm() (ok bool)
		NtAccessCheckByType() (ok bool)
		NtAccessCheckByTypeAndAuditAlarm() (ok bool)
		NtAccessCheckByTypeResultList() (ok bool)
		NtAccessCheckByTypeResultListAndAuditAlarm() (ok bool)
		NtAccessCheckByTypeResultListAndAuditAlarmByHandle() (ok bool)
		NtAcquireCrossVmMutant() (ok bool)
		NtAcquireProcessActivityReference() (ok bool)
		NtAddAtom() (ok bool)
		NtAddAtomEx() (ok bool)
		NtAddBootEntry() (ok bool)
		NtAddDriverEntry() (ok bool)
		NtAdjustGroupsToken() (ok bool)
		NtAdjustPrivilegesToken() (ok bool)
		NtAdjustTokenClaimsAndDeviceGroups() (ok bool)
		NtAlertResumeThread() (ok bool)
		NtAlertThread() (ok bool)
		NtAlertThreadByThreadId() (ok bool)
		NtAllocateLocallyUniqueId() (ok bool)
		NtAllocateReserveObject() (ok bool)
		NtAllocateUserPhysicalPages() (ok bool)
		NtAllocateUserPhysicalPagesEx() (ok bool)
		NtAllocateUuids() (ok bool)
		NtAllocateVirtualMemory() (ok bool)
		NtAllocateVirtualMemoryEx() (ok bool)
		NtAlpcAcceptConnectPort() (ok bool)
		NtAlpcCancelMessage() (ok bool)
		NtAlpcConnectPort() (ok bool)
		NtAlpcConnectPortEx() (ok bool)
		NtAlpcCreatePort() (ok bool)
		NtAlpcCreatePortSection() (ok bool)
		NtAlpcCreateResourceReserve() (ok bool)
		NtAlpcCreateSectionView() (ok bool)
		NtAlpcCreateSecurityContext() (ok bool)
		NtAlpcDeletePortSection() (ok bool)
		NtAlpcDeleteResourceReserve() (ok bool)
		NtAlpcDeleteSectionView() (ok bool)
		NtAlpcDeleteSecurityContext() (ok bool)
		NtAlpcDisconnectPort() (ok bool)
		NtAlpcImpersonateClientContainerOfPort() (ok bool)
		NtAlpcImpersonateClientOfPort() (ok bool)
		NtAlpcOpenSenderProcess() (ok bool)
		NtAlpcOpenSenderThread() (ok bool)
		NtAlpcQueryInformation() (ok bool)
		NtAlpcQueryInformationMessage() (ok bool)
		NtAlpcRevokeSecurityContext() (ok bool)
		NtAlpcSendWaitReceivePort() (ok bool)
		NtAlpcSetInformation() (ok bool)
		NtApphelpCacheControl() (ok bool)
		NtAreMappedFilesTheSame() (ok bool)
		NtAssignProcessToJobObject() (ok bool)
		NtAssociateWaitCompletionPacket() (ok bool)
		NtCallEnclave() (ok bool)
		NtCallbackReturn() (ok bool)
		NtCancelIoFile() (ok bool)
		NtCancelIoFileEx() (ok bool)
		NtCancelSynchronousIoFile() (ok bool)
		NtCancelTimer() (ok bool)
		NtCancelTimer2() (ok bool)
		NtCancelWaitCompletionPacket() (ok bool)
		NtChangeProcessState() (ok bool)
		NtChangeThreadState() (ok bool)
		NtClearEvent() (ok bool)
		NtClose() (ok bool)
		NtCloseObjectAuditAlarm() (ok bool)
		NtCommitComplete() (ok bool)
		NtCommitEnlistment() (ok bool)
		NtCommitRegistryTransaction() (ok bool)
		NtCommitTransaction() (ok bool)
		NtCompactKeys() (ok bool)
		NtCompareObjects() (ok bool)
		NtCompareSigningLevels() (ok bool)
		NtCompareTokens() (ok bool)
		NtCompleteConnectPort() (ok bool)
		NtCompressKey() (ok bool)
		NtConnectPort() (ok bool)
		NtContinue() (ok bool)
		NtContinueEx() (ok bool)
		NtConvertBetweenAuxiliaryCounterAndPerformanceCounter() (ok bool)
		NtCopyFileChunk() (ok bool)
		NtCreateCpuPartition() (ok bool)
		NtCreateCrossVmEvent() (ok bool)
		NtCreateCrossVmMutant() (ok bool)
		NtCreateDebugObject() (ok bool)
		NtCreateDirectoryObject() (ok bool)
		NtCreateDirectoryObjectEx() (ok bool)
		NtCreateEnclave() (ok bool)
		NtCreateEnlistment() (ok bool)
		NtCreateEvent() (ok bool)
		NtCreateEventPair() (ok bool)
		NtCreateFile() (ok bool)
		NtCreateIRTimer() (ok bool)
		NtCreateIoCompletion() (ok bool)
		NtCreateIoRing() (ok bool)
		NtCreateJobObject() (ok bool)
		NtCreateJobSet() (ok bool)
		NtCreateKey() (ok bool)
		NtCreateKeyTransacted() (ok bool)
		NtCreateKeyedEvent() (ok bool)
		NtCreateLowBoxToken() (ok bool)
		NtCreateMailslotFile() (ok bool)
		NtCreateMutant() (ok bool)
		NtCreateNamedPipeFile() (ok bool)
		NtCreatePagingFile() (ok bool)
		NtCreatePartition() (ok bool)
		NtCreatePort() (ok bool)
		NtCreatePrivateNamespace() (ok bool)
		NtCreateProcess() (ok bool)
		NtCreateProcessEx() (ok bool)
		NtCreateProcessStateChange() (ok bool)
		NtCreateProfile() (ok bool)
		NtCreateProfileEx() (ok bool)
		NtCreateRegistryTransaction() (ok bool)
		NtCreateResourceManager() (ok bool)
		NtCreateSection() (ok bool)
		NtCreateSectionEx() (ok bool)
		NtCreateSemaphore() (ok bool)
		NtCreateSymbolicLinkObject() (ok bool)
		NtCreateThread() (ok bool)
		NtCreateThreadEx() (ok bool)
		NtCreateThreadStateChange() (ok bool)
		NtCreateTimer() (ok bool)
		NtCreateTimer2() (ok bool)
		NtCreateToken() (ok bool)
		NtCreateTokenEx() (ok bool)
		NtCreateTransaction() (ok bool)
		NtCreateTransactionManager() (ok bool)
		NtCreateUserProcess() (ok bool)
		NtCreateWaitCompletionPacket() (ok bool)
		NtCreateWaitablePort() (ok bool)
		NtCreateWnfStateName() (ok bool)
		NtCreateWorkerFactory() (ok bool)
		NtDebugActiveProcess() (ok bool)
		NtDebugContinue() (ok bool)
		NtDelayExecution() (ok bool)
		NtDeleteAtom() (ok bool)
		NtDeleteBootEntry() (ok bool)
		NtDeleteDriverEntry() (ok bool)
		NtDeleteFile() (ok bool)
		NtDeleteKey() (ok bool)
		NtDeleteObjectAuditAlarm() (ok bool)
		NtDeletePrivateNamespace() (ok bool)
		NtDeleteValueKey() (ok bool)
		NtDeleteWnfStateData() (ok bool)
		NtDeleteWnfStateName() (ok bool)
		NtDeviceIoControlFile() (ok bool)
		NtDirectGraphicsCall() (ok bool)
		NtDisableLastKnownGood() (ok bool)
		NtDisplayString() (ok bool)
		NtDrawText() (ok bool)
		NtDuplicateObject() (ok bool)
		NtDuplicateToken() (ok bool)
		NtEnableLastKnownGood() (ok bool)
		NtEnumerateBootEntries() (ok bool)
		NtEnumerateDriverEntries() (ok bool)
		NtEnumerateKey() (ok bool)
		NtEnumerateSystemEnvironmentValuesEx() (ok bool)
		NtEnumerateTransactionObject() (ok bool)
		NtEnumerateValueKey() (ok bool)
		NtExtendSection() (ok bool)
		NtFilterBootOption() (ok bool)
		NtFilterToken() (ok bool)
		NtFilterTokenEx() (ok bool)
		NtFindAtom() (ok bool)
		NtFlushBuffersFile() (ok bool)
		NtFlushBuffersFileEx() (ok bool)
		NtFlushInstallUILanguage() (ok bool)
		NtFlushInstructionCache() (ok bool)
		NtFlushKey() (ok bool)
		NtFlushProcessWriteBuffers() (ok bool)
		NtFlushVirtualMemory() (ok bool)
		NtFlushWriteBuffer() (ok bool)
		NtFreeUserPhysicalPages() (ok bool)
		NtFreeVirtualMemory() (ok bool)
		NtFreezeRegistry() (ok bool)
		NtFreezeTransactions() (ok bool)
		NtFsControlFile() (ok bool)
		NtGetCachedSigningLevel() (ok bool)
		NtGetCompleteWnfStateSubscription() (ok bool)
		NtGetContextThread() (ok bool)
		NtGetCurrentProcessorNumber() (ok bool)
		NtGetCurrentProcessorNumberEx() (ok bool)
		NtGetDevicePowerState() (ok bool)
		NtGetMUIRegistryInfo() (ok bool)
		NtGetNextProcess() (ok bool)
		NtGetNextThread() (ok bool)
		NtGetNlsSectionPtr() (ok bool)
		NtGetNotificationResourceManager() (ok bool)
		NtGetWriteWatch() (ok bool)
		NtImpersonateAnonymousToken() (ok bool)
		NtImpersonateClientOfPort() (ok bool)
		NtImpersonateThread() (ok bool)
		NtInitializeEnclave() (ok bool)
		NtInitializeNlsFiles() (ok bool)
		NtInitializeRegistry() (ok bool)
		NtInitiatePowerAction() (ok bool)
		NtIsProcessInJob() (ok bool)
		NtIsSystemResumeAutomatic() (ok bool)
		NtIsUILanguageComitted() (ok bool)
		NtListenPort() (ok bool)
		NtLoadDriver() (ok bool)
		NtLoadEnclaveData() (ok bool)
		NtLoadKey() (ok bool)
		NtLoadKey2() (ok bool)
		NtLoadKey3() (ok bool)
		NtLoadKeyEx() (ok bool)
		NtLockFile() (ok bool)
		NtLockProductActivationKeys() (ok bool)
		NtLockRegistryKey() (ok bool)
		NtLockVirtualMemory() (ok bool)
		NtMakePermanentObject() (ok bool)
		NtMakeTemporaryObject() (ok bool)
		NtManageHotPatch() (ok bool)
		NtManagePartition() (ok bool)
		NtMapCMFModule() (ok bool)
		NtMapUserPhysicalPages() (ok bool)
		NtMapUserPhysicalPagesScatter() (ok bool)
		NtMapViewOfSection() (ok bool)
		NtMapViewOfSectionEx() (ok bool)
		NtModifyBootEntry() (ok bool)
		NtModifyDriverEntry() (ok bool)
		NtNotifyChangeDirectoryFile() (ok bool)
		NtNotifyChangeDirectoryFileEx() (ok bool)
		NtNotifyChangeKey() (ok bool)
		NtNotifyChangeMultipleKeys() (ok bool)
		NtNotifyChangeSession() (ok bool)
		NtOpenCpuPartition() (ok bool)
		NtOpenDirectoryObject() (ok bool)
		NtOpenEnlistment() (ok bool)
		NtOpenEvent() (ok bool)
		NtOpenEventPair() (ok bool)
		NtOpenFile() (ok bool)
		NtOpenIoCompletion() (ok bool)
		NtOpenJobObject() (ok bool)
		NtOpenKey() (ok bool)
		NtOpenKeyEx() (ok bool)
		NtOpenKeyTransacted() (ok bool)
		NtOpenKeyTransactedEx() (ok bool)
		NtOpenKeyedEvent() (ok bool)
		NtOpenMutant() (ok bool)
		NtOpenObjectAuditAlarm() (ok bool)
		NtOpenPartition() (ok bool)
		NtOpenPrivateNamespace() (ok bool)
		NtOpenProcess() (ok bool)
		NtOpenProcessToken() (ok bool)
		NtOpenProcessTokenEx() (ok bool)
		NtOpenRegistryTransaction() (ok bool)
		NtOpenResourceManager() (ok bool)
		NtOpenSection() (ok bool)
		NtOpenSemaphore() (ok bool)
		NtOpenSession() (ok bool)
		NtOpenSymbolicLinkObject() (ok bool)
		NtOpenThread() (ok bool)
		NtOpenThreadToken() (ok bool)
		NtOpenThreadTokenEx() (ok bool)
		NtOpenTimer() (ok bool)
		NtOpenTransaction() (ok bool)
		NtOpenTransactionManager() (ok bool)
		NtPlugPlayControl() (ok bool)
		NtPowerInformation() (ok bool)
		NtPrePrepareComplete() (ok bool)
		NtPrePrepareEnlistment() (ok bool)
		NtPrepareComplete() (ok bool)
		NtPrepareEnlistment() (ok bool)
		NtPrivilegeCheck() (ok bool)
		NtPrivilegeObjectAuditAlarm() (ok bool)
		NtPrivilegedServiceAuditAlarm() (ok bool)
		NtPropagationComplete() (ok bool)
		NtPropagationFailed() (ok bool)
		NtProtectVirtualMemory() (ok bool)
		NtPssCaptureVaSpaceBulk() (ok bool)
		NtPulseEvent() (ok bool)
		NtQueryAttributesFile() (ok bool)
		NtQueryAuxiliaryCounterFrequency() (ok bool)
		NtQueryBootEntryOrder() (ok bool)
		NtQueryBootOptions() (ok bool)
		NtQueryDebugFilterState() (ok bool)
		NtQueryDefaultLocale() (ok bool)
		NtQueryDefaultUILanguage() (ok bool)
		NtQueryDirectoryFile() (ok bool)
		NtQueryDirectoryFileEx() (ok bool)
		NtQueryDirectoryObject() (ok bool)
		NtQueryDriverEntryOrder() (ok bool)
		NtQueryEaFile() (ok bool)
		NtQueryEvent() (ok bool)
		NtQueryFullAttributesFile() (ok bool)
		NtQueryInformationAtom() (ok bool)
		NtQueryInformationByName() (ok bool)
		NtQueryInformationCpuPartition() (ok bool)
		NtQueryInformationEnlistment() (ok bool)
		NtQueryInformationFile() (ok bool)
		NtQueryInformationJobObject() (ok bool)
		NtQueryInformationPort() (ok bool)
		NtQueryInformationProcess() (ok bool)
		NtQueryInformationResourceManager() (ok bool)
		NtQueryInformationThread() (ok bool)
		NtQueryInformationToken() (ok bool)
		NtQueryInformationTransaction() (ok bool)
		NtQueryInformationTransactionManager() (ok bool)
		NtQueryInformationWorkerFactory() (ok bool)
		NtQueryInstallUILanguage() (ok bool)
		NtQueryIntervalProfile() (ok bool)
		NtQueryIoCompletion() (ok bool)
		NtQueryIoRingCapabilities() (ok bool)
		NtQueryKey() (ok bool)
		NtQueryLicenseValue() (ok bool)
		NtQueryMultipleValueKey() (ok bool)
		NtQueryMutant() (ok bool)
		NtQueryObject() (ok bool)
		NtQueryOpenSubKeys() (ok bool)
		NtQueryOpenSubKeysEx() (ok bool)
		NtQueryPerformanceCounter() (ok bool)
		NtQueryPortInformationProcess() (ok bool)
		NtQueryQuotaInformationFile() (ok bool)
		NtQuerySection() (ok bool)
		NtQuerySecurityAttributesToken() (ok bool)
		NtQuerySecurityObject() (ok bool)
		NtQuerySecurityPolicy() (ok bool)
		NtQuerySemaphore() (ok bool)
		NtQuerySymbolicLinkObject() (ok bool)
		NtQuerySystemEnvironmentValue() (ok bool)
		NtQuerySystemEnvironmentValueEx() (ok bool)
		NtQuerySystemInformation() (ok bool)
		NtQuerySystemInformationEx() (ok bool)
		NtQueryTimer() (ok bool)
		NtQueryTimerResolution() (ok bool)
		NtQueryValueKey() (ok bool)
		NtQueryVirtualMemory() (ok bool)
		NtQueryVolumeInformationFile() (ok bool)
		NtQueryWnfStateData() (ok bool)
		NtQueryWnfStateNameInformation() (ok bool)
		NtQueueApcThread() (ok bool)
		NtQueueApcThreadEx() (ok bool)
		NtQueueApcThreadEx2() (ok bool)
		NtRaiseException() (ok bool)
		NtRaiseHardError() (ok bool)
		NtReadFile() (ok bool)
		NtReadFileScatter() (ok bool)
		NtReadOnlyEnlistment() (ok bool)
		NtReadRequestData() (ok bool)
		NtReadVirtualMemory() (ok bool)
		NtReadVirtualMemoryEx() (ok bool)
		NtRecoverEnlistment() (ok bool)
		NtRecoverResourceManager() (ok bool)
		NtRecoverTransactionManager() (ok bool)
		NtRegisterProtocolAddressInformation() (ok bool)
		NtRegisterThreadTerminatePort() (ok bool)
		NtReleaseKeyedEvent() (ok bool)
		NtReleaseMutant() (ok bool)
		NtReleaseSemaphore() (ok bool)
		NtReleaseWorkerFactoryWorker() (ok bool)
		NtRemoveIoCompletion() (ok bool)
		NtRemoveIoCompletionEx() (ok bool)
		NtRemoveProcessDebug() (ok bool)
		NtRenameKey() (ok bool)
		NtRenameTransactionManager() (ok bool)
		NtReplaceKey() (ok bool)
		NtReplacePartitionUnit() (ok bool)
		NtReplyPort() (ok bool)
		NtReplyWaitReceivePort() (ok bool)
		NtReplyWaitReceivePortEx() (ok bool)
		NtReplyWaitReplyPort() (ok bool)
		NtRequestPort() (ok bool)
		NtRequestWaitReplyPort() (ok bool)
		NtResetEvent() (ok bool)
		NtResetWriteWatch() (ok bool)
		NtRestoreKey() (ok bool)
		NtResumeProcess() (ok bool)
		NtResumeThread() (ok bool)
		NtRevertContainerImpersonation() (ok bool)
		NtRollbackComplete() (ok bool)
		NtRollbackEnlistment() (ok bool)
		NtRollbackRegistryTransaction() (ok bool)
		NtRollbackTransaction() (ok bool)
		NtRollforwardTransactionManager() (ok bool)
		NtSaveKey() (ok bool)
		NtSaveKeyEx() (ok bool)
		NtSaveMergedKeys() (ok bool)
		NtSecureConnectPort() (ok bool)
		NtSerializeBoot() (ok bool)
		NtSetBootEntryOrder() (ok bool)
		NtSetBootOptions() (ok bool)
		NtSetCachedSigningLevel() (ok bool)
		NtSetCachedSigningLevel2() (ok bool)
		NtSetContextThread() (ok bool)
		NtSetDebugFilterState() (ok bool)
		NtSetDefaultHardErrorPort() (ok bool)
		NtSetDefaultLocale() (ok bool)
		NtSetDefaultUILanguage() (ok bool)
		NtSetDriverEntryOrder() (ok bool)
		NtSetEaFile() (ok bool)
		NtSetEvent() (ok bool)
		NtSetEventBoostPriority() (ok bool)
		NtSetHighEventPair() (ok bool)
		NtSetHighWaitLowEventPair() (ok bool)
		NtSetIRTimer() (ok bool)
		NtSetInformationCpuPartition() (ok bool)
		NtSetInformationDebugObject() (ok bool)
		NtSetInformationEnlistment() (ok bool)
		NtSetInformationFile() (ok bool)
		NtSetInformationIoRing() (ok bool)
		NtSetInformationJobObject() (ok bool)
		NtSetInformationKey() (ok bool)
		NtSetInformationObject() (ok bool)
		NtSetInformationProcess() (ok bool)
		NtSetInformationResourceManager() (ok bool)
		NtSetInformationSymbolicLink() (ok bool)
		NtSetInformationThread() (ok bool)
		NtSetInformationToken() (ok bool)
		NtSetInformationTransaction() (ok bool)
		NtSetInformationTransactionManager() (ok bool)
		NtSetInformationVirtualMemory() (ok bool)
		NtSetInformationWorkerFactory() (ok bool)
		NtSetIntervalProfile() (ok bool)
		NtSetIoCompletion() (ok bool)
		NtSetIoCompletionEx() (ok bool)
		NtSetLdtEntries() (ok bool)
		NtSetLowEventPair() (ok bool)
		NtSetLowWaitHighEventPair() (ok bool)
		NtSetQuotaInformationFile() (ok bool)
		NtSetSecurityObject() (ok bool)
		NtSetSystemEnvironmentValue() (ok bool)
		NtSetSystemEnvironmentValueEx() (ok bool)
		NtSetSystemInformation() (ok bool)
		NtSetSystemPowerState() (ok bool)
		NtSetSystemTime() (ok bool)
		NtSetThreadExecutionState() (ok bool)
		NtSetTimer() (ok bool)
		NtSetTimer2() (ok bool)
		NtSetTimerEx() (ok bool)
		NtSetTimerResolution() (ok bool)
		NtSetUuidSeed() (ok bool)
		NtSetValueKey() (ok bool)
		NtSetVolumeInformationFile() (ok bool)
		NtSetWnfProcessNotificationEvent() (ok bool)
		NtShutdownSystem() (ok bool)
		NtShutdownWorkerFactory() (ok bool)
		NtSignalAndWaitForSingleObject() (ok bool)
		NtSinglePhaseReject() (ok bool)
		NtStartProfile() (ok bool)
		NtStopProfile() (ok bool)
		NtSubmitIoRing() (ok bool)
		NtSubscribeWnfStateChange() (ok bool)
		NtSuspendProcess() (ok bool)
		NtSuspendThread() (ok bool)
		NtSystemDebugControl() (ok bool)
		NtTerminateEnclave() (ok bool)
		NtTerminateJobObject() (ok bool)
		NtTerminateProcess() (ok bool)
		NtTerminateThread() (ok bool)
		NtTestAlert() (ok bool)
		NtThawRegistry() (ok bool)
		NtThawTransactions() (ok bool)
		NtTraceControl() (ok bool)
		NtTraceEvent() (ok bool)
		NtTranslateFilePath() (ok bool)
		NtUmsThreadYield() (ok bool)
		NtUnloadDriver() (ok bool)
		NtUnloadKey() (ok bool)
		NtUnloadKey2() (ok bool)
		NtUnloadKeyEx() (ok bool)
		NtUnlockFile() (ok bool)
		NtUnlockVirtualMemory() (ok bool)
		NtUnmapViewOfSection() (ok bool)
		NtUnmapViewOfSectionEx() (ok bool)
		NtUnsubscribeWnfStateChange() (ok bool)
		NtUpdateWnfStateData() (ok bool)
		NtVdmControl() (ok bool)
		NtWaitForAlertByThreadId() (ok bool)
		NtWaitForDebugEvent() (ok bool)
		NtWaitForKeyedEvent() (ok bool)
		NtWaitForMultipleObjects() (ok bool)
		NtWaitForMultipleObjects32() (ok bool)
		NtWaitForSingleObject() (ok bool)
		NtWaitForWorkViaWorkerFactory() (ok bool)
		NtWaitHighEventPair() (ok bool)
		NtWaitLowEventPair() (ok bool)
		NtWorkerFactoryWorkerReady() (ok bool)
		NtWriteFile() (ok bool)
		NtWriteFileGather() (ok bool)
		NtWriteRequestData() (ok bool)
		NtWriteVirtualMemory() (ok bool)
		NtYieldExecution() (ok bool)
	}
	objectNtdll struct{}
)

func (o *objectNtdll) NtAcceptConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheck() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckAndAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckByType() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckByTypeAndAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckByTypeResultList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckByTypeResultListAndAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAccessCheckByTypeResultListAndAuditAlarmByHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAcquireCrossVmMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAcquireProcessActivityReference() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAddAtom() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAddAtomEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAddBootEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAddDriverEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAdjustGroupsToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAdjustPrivilegesToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAdjustTokenClaimsAndDeviceGroups() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlertResumeThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlertThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlertThreadByThreadId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateLocallyUniqueId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateReserveObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateUserPhysicalPages() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateUserPhysicalPagesEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateUuids() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAllocateVirtualMemoryEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcAcceptConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCancelMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcConnectPortEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCreatePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCreatePortSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCreateResourceReserve() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCreateSectionView() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcCreateSecurityContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcDeletePortSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcDeleteResourceReserve() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcDeleteSectionView() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcDeleteSecurityContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcDisconnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcImpersonateClientContainerOfPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcImpersonateClientOfPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcOpenSenderProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcOpenSenderThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcQueryInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcQueryInformationMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcRevokeSecurityContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcSendWaitReceivePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAlpcSetInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtApphelpCacheControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAreMappedFilesTheSame() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAssignProcessToJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtAssociateWaitCompletionPacket() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCallEnclave() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCallbackReturn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelIoFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelIoFileEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelSynchronousIoFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelTimer2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCancelWaitCompletionPacket() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtChangeProcessState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtChangeThreadState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtClearEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtClose() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCloseObjectAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCommitComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCommitEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCommitRegistryTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCommitTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompactKeys() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompareObjects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompareSigningLevels() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompareTokens() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompleteConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCompressKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtContinue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtContinueEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtConvertBetweenAuxiliaryCounterAndPerformanceCounter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCopyFileChunk() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateCpuPartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateCrossVmEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateCrossVmMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateDebugObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateDirectoryObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateDirectoryObjectEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateEnclave() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateIRTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateIoCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateIoRing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateJobSet() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateKeyTransacted() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateKeyedEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateLowBoxToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateMailslotFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateNamedPipeFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreatePagingFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreatePartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreatePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreatePrivateNamespace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateProcessEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateProcessStateChange() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateProfile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateProfileEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateRegistryTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateSectionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateSemaphore() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateSymbolicLinkObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateThreadEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateThreadStateChange() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateTimer2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateTokenEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateUserProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateWaitCompletionPacket() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateWaitablePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateWnfStateName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtCreateWorkerFactory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDebugActiveProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDebugContinue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDelayExecution() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteAtom() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteBootEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteDriverEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteObjectAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeletePrivateNamespace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteValueKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteWnfStateData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeleteWnfStateName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDeviceIoControlFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDirectGraphicsCall() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDisableLastKnownGood() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDisplayString() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDrawText() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDuplicateObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtDuplicateToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnableLastKnownGood() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateBootEntries() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateDriverEntries() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateSystemEnvironmentValuesEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateTransactionObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtEnumerateValueKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtExtendSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFilterBootOption() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFilterToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFilterTokenEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFindAtom() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushBuffersFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushBuffersFileEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushInstallUILanguage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushInstructionCache() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushProcessWriteBuffers() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFlushWriteBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFreeUserPhysicalPages() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFreeVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFreezeRegistry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFreezeTransactions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtFsControlFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetCachedSigningLevel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetCompleteWnfStateSubscription() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetContextThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetCurrentProcessorNumber() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetCurrentProcessorNumberEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetDevicePowerState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetMUIRegistryInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetNextProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetNextThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetNlsSectionPtr() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetNotificationResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtGetWriteWatch() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtImpersonateAnonymousToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtImpersonateClientOfPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtImpersonateThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtInitializeEnclave() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtInitializeNlsFiles() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtInitializeRegistry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtInitiatePowerAction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtIsProcessInJob() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtIsSystemResumeAutomatic() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtIsUILanguageComitted() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtListenPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadDriver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadEnclaveData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadKey2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadKey3() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLoadKeyEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLockFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLockProductActivationKeys() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLockRegistryKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtLockVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMakePermanentObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMakeTemporaryObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtManageHotPatch() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtManagePartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMapCMFModule() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMapUserPhysicalPages() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMapUserPhysicalPagesScatter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMapViewOfSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtMapViewOfSectionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtModifyBootEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtModifyDriverEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtNotifyChangeDirectoryFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtNotifyChangeDirectoryFileEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtNotifyChangeKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtNotifyChangeMultipleKeys() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtNotifyChangeSession() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenCpuPartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenDirectoryObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenIoCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenKeyEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenKeyTransacted() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenKeyTransactedEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenKeyedEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenObjectAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenPartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenPrivateNamespace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenProcessToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenProcessTokenEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenRegistryTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenSemaphore() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenSession() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenSymbolicLinkObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenThreadToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenThreadTokenEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtOpenTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPlugPlayControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPowerInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrePrepareComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrePrepareEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrepareComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrepareEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrivilegeCheck() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrivilegeObjectAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPrivilegedServiceAuditAlarm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPropagationComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPropagationFailed() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtProtectVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPssCaptureVaSpaceBulk() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtPulseEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryAttributesFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryAuxiliaryCounterFrequency() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryBootEntryOrder() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryBootOptions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDebugFilterState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDefaultLocale() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDefaultUILanguage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDirectoryFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDirectoryFileEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDirectoryObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryDriverEntryOrder() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryEaFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryFullAttributesFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationAtom() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationByName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationCpuPartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInformationWorkerFactory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryInstallUILanguage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryIntervalProfile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryIoCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryIoRingCapabilities() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryLicenseValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryMultipleValueKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryOpenSubKeys() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryOpenSubKeysEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryPerformanceCounter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryPortInformationProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryQuotaInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySecurityAttributesToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySecurityObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySecurityPolicy() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySemaphore() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySymbolicLinkObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySystemEnvironmentValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySystemEnvironmentValueEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySystemInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQuerySystemInformationEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryTimerResolution() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryValueKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryVolumeInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryWnfStateData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueryWnfStateNameInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueueApcThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueueApcThreadEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtQueueApcThreadEx2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRaiseException() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRaiseHardError() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadFileScatter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadOnlyEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadRequestData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReadVirtualMemoryEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRecoverEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRecoverResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRecoverTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRegisterProtocolAddressInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRegisterThreadTerminatePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReleaseKeyedEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReleaseMutant() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReleaseSemaphore() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReleaseWorkerFactoryWorker() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRemoveIoCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRemoveIoCompletionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRemoveProcessDebug() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRenameKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRenameTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplaceKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplacePartitionUnit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplyPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplyWaitReceivePort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplyWaitReceivePortEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtReplyWaitReplyPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRequestPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRequestWaitReplyPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtResetEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtResetWriteWatch() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRestoreKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtResumeProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtResumeThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRevertContainerImpersonation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRollbackComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRollbackEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRollbackRegistryTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRollbackTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtRollforwardTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSaveKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSaveKeyEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSaveMergedKeys() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSecureConnectPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSerializeBoot() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetBootEntryOrder() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetBootOptions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetCachedSigningLevel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetCachedSigningLevel2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetContextThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetDebugFilterState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetDefaultHardErrorPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetDefaultLocale() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetDefaultUILanguage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetDriverEntryOrder() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetEaFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetEventBoostPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetHighEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetHighWaitLowEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetIRTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationCpuPartition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationDebugObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationEnlistment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationIoRing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationResourceManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationSymbolicLink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationTransaction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationTransactionManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetInformationWorkerFactory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetIntervalProfile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetIoCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetIoCompletionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetLdtEntries() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetLowEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetLowWaitHighEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetQuotaInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSecurityObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSystemEnvironmentValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSystemEnvironmentValueEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSystemInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSystemPowerState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetSystemTime() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetThreadExecutionState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetTimer2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetTimerEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetTimerResolution() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetUuidSeed() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetValueKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetVolumeInformationFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSetWnfProcessNotificationEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtShutdownSystem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtShutdownWorkerFactory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSignalAndWaitForSingleObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSinglePhaseReject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtStartProfile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtStopProfile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSubmitIoRing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSubscribeWnfStateChange() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSuspendProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSuspendThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtSystemDebugControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTerminateEnclave() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTerminateJobObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTerminateProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTerminateThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTestAlert() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtThawRegistry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtThawTransactions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTraceControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTraceEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtTranslateFilePath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUmsThreadYield() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnloadDriver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnloadKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnloadKey2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnloadKeyEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnlockFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnlockVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnmapViewOfSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnmapViewOfSectionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUnsubscribeWnfStateChange() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtUpdateWnfStateData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtVdmControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForAlertByThreadId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForDebugEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForKeyedEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForMultipleObjects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForMultipleObjects32() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForSingleObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitForWorkViaWorkerFactory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitHighEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWaitLowEventPair() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWorkerFactoryWorkerReady() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWriteFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWriteFileGather() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWriteRequestData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtWriteVirtualMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectNtdll) NtYieldExecution() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func NewNtdll() InterfaceNtdll { return &objectNtdll{} }
