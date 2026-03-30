#![allow(non_snake_case)]
#![allow(dead_code)]

// EPT Hook Database: 1251 functions
// EPT hooks provide stealth by using Extended Page Tables

use alloc::string::String;
use alloc::vec::Vec;
use crate::hyperkd::hyperhv::hooks::hooks::*;
use crate::hyperkd::hyperhv::ProcessId;
use super::*;

pub struct EptHookDb {
    pub name: &'static str,
    pub address: u64,
    pub enabled: bool,
}

pub static EPT_HOOK_DATABASE: &[EptHookDb] = &[
    EptHookDb { name: "DbgBreakPointWithStatus", address: 0, enabled: false },
    EptHookDb { name: "DbgPrint", address: 0, enabled: false },
    EptHookDb { name: "DbgPrintEx", address: 0, enabled: false },
    EptHookDb { name: "DbgPrintReturnControlC", address: 0, enabled: false },
    EptHookDb { name: "DbgPrompt", address: 0, enabled: false },
    EptHookDb { name: "DbgQueryDebugFilterState", address: 0, enabled: false },
    EptHookDb { name: "DbgSetDebugFilterState", address: 0, enabled: false },
    EptHookDb { name: "DbgSetDebugPrintCallback", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireFastMutex", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireFastMutexUnsafe", address: 0, enabled: false },
    EptHookDb { name: "ExAcquirePushLockExclusiveEx", address: 0, enabled: false },
    EptHookDb { name: "ExAcquirePushLockSharedEx", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireResourceExclusiveLite", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireResourceSharedLite", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireRundownProtectionCacheAwareEx", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireRundownProtectionEx", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSharedStarveExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSharedWaitForExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSpinLockExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSpinLockExclusiveAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSpinLockShared", address: 0, enabled: false },
    EptHookDb { name: "ExAcquireSpinLockSharedAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExAdjustLookasideDepth", address: 0, enabled: false },
    EptHookDb { name: "ExAllocateCacheAwareRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExAllocateFromLookasideListEx", address: 0, enabled: false },
    EptHookDb { name: "ExAllocateFromNPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExAllocateFromPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExAllocatePool2", address: 0, enabled: false },
    EptHookDb { name: "ExAllocatePool3", address: 0, enabled: false },
    EptHookDb { name: "ExAllocatePoolWithQuota", address: 0, enabled: false },
    EptHookDb { name: "ExAllocateTimer", address: 0, enabled: false },
    EptHookDb { name: "ExCancelTimer", address: 0, enabled: false },
    EptHookDb { name: "ExCleanupRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExConvertExclusiveToSharedLite", address: 0, enabled: false },
    EptHookDb { name: "ExCreateCallback", address: 0, enabled: false },
    EptHookDb { name: "ExCreatePool", address: 0, enabled: false },
    EptHookDb { name: "ExDeleteLookasideListEx", address: 0, enabled: false },
    EptHookDb { name: "ExDeleteNPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExDeletePagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExDeleteResourceLite", address: 0, enabled: false },
    EptHookDb { name: "ExDeleteTimer", address: 0, enabled: false },
    EptHookDb { name: "ExDestroyPool", address: 0, enabled: false },
    EptHookDb { name: "ExDisableResourceBoostLite", address: 0, enabled: false },
    EptHookDb { name: "ExEnterCriticalRegionAndAcquireResourceExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExEnterCriticalRegionAndAcquireResourceShared", address: 0, enabled: false },
    EptHookDb { name: "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExEnumerateSystemFirmwareTables", address: 0, enabled: false },
    EptHookDb { name: "ExExtendZone", address: 0, enabled: false },
    EptHookDb { name: "ExFlushLookasideListEx", address: 0, enabled: false },
    EptHookDb { name: "ExFreeCacheAwareRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExFreePool", address: 0, enabled: false },
    EptHookDb { name: "ExFreePool2", address: 0, enabled: false },
    EptHookDb { name: "ExFreePoolWithTag", address: 0, enabled: false },
    EptHookDb { name: "ExFreeToLookasideListEx", address: 0, enabled: false },
    EptHookDb { name: "ExFreeToNPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExFreeToPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExGetExclusiveWaiterCount", address: 0, enabled: false },
    EptHookDb { name: "ExGetFirmwareEnvironmentVariable", address: 0, enabled: false },
    EptHookDb { name: "ExGetFirmwareType", address: 0, enabled: false },
    EptHookDb { name: "ExGetPreviousMode", address: 0, enabled: false },
    EptHookDb { name: "ExGetSharedWaiterCount", address: 0, enabled: false },
    EptHookDb { name: "ExGetSystemFirmwareTable", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeDeviceAts", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeLookasideListEx", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeNPagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExInitializePagedLookasideList", address: 0, enabled: false },
    EptHookDb { name: "ExInitializePushLock", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeResourceLite", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeRundownProtectionCacheAwareEx", address: 0, enabled: false },
    EptHookDb { name: "ExInitializeZone", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedAddLargeInteger", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedAddUlong", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedExtendZone", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedInsertHeadList", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedInsertTailList", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedPopEntryList", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedPushEntryList", address: 0, enabled: false },
    EptHookDb { name: "ExInterlockedRemoveHeadList", address: 0, enabled: false },
    EptHookDb { name: "ExIsManufacturingModeEnabled", address: 0, enabled: false },
    EptHookDb { name: "ExIsProcessorFeaturePresent", address: 0, enabled: false },
    EptHookDb { name: "ExIsResourceAcquiredExclusiveLite", address: 0, enabled: false },
    EptHookDb { name: "ExIsResourceAcquiredSharedLite", address: 0, enabled: false },
    EptHookDb { name: "ExIsSoftBoot", address: 0, enabled: false },
    EptHookDb { name: "ExLocalTimeToSystemTime", address: 0, enabled: false },
    EptHookDb { name: "ExNotifyCallback", address: 0, enabled: false },
    EptHookDb { name: "ExQueryDepthSList", address: 0, enabled: false },
    EptHookDb { name: "ExQueryPoolBlockSize", address: 0, enabled: false },
    EptHookDb { name: "ExQueryTimerResolution", address: 0, enabled: false },
    EptHookDb { name: "ExQueueWorkItem", address: 0, enabled: false },
    EptHookDb { name: "ExRaiseAccessViolation", address: 0, enabled: false },
    EptHookDb { name: "ExRaiseDatatypeMisalignment", address: 0, enabled: false },
    EptHookDb { name: "ExRaiseStatus", address: 0, enabled: false },
    EptHookDb { name: "ExRcuFreePool", address: 0, enabled: false },
    EptHookDb { name: "ExReInitializeRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExReInitializeRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExRegisterCallback", address: 0, enabled: false },
    EptHookDb { name: "ExReinitializeResourceLite", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseFastMutex", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseFastMutexUnsafe", address: 0, enabled: false },
    EptHookDb { name: "ExReleasePushLockExclusiveEx", address: 0, enabled: false },
    EptHookDb { name: "ExReleasePushLockSharedEx", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseResourceAndLeaveCriticalRegion", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseResourceForThreadLite", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseResourceLite", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseRundownProtection", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseRundownProtectionCacheAwareEx", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseRundownProtectionEx", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseSpinLockExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseSpinLockExclusiveFromDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseSpinLockShared", address: 0, enabled: false },
    EptHookDb { name: "ExReleaseSpinLockSharedFromDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExRundownCompleted", address: 0, enabled: false },
    EptHookDb { name: "ExRundownCompletedCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExSecurePoolUpdate", address: 0, enabled: false },
    EptHookDb { name: "ExSecurePoolValidate", address: 0, enabled: false },
    EptHookDb { name: "ExSetFirmwareEnvironmentVariable", address: 0, enabled: false },
    EptHookDb { name: "ExSetResourceOwnerPointer", address: 0, enabled: false },
    EptHookDb { name: "ExSetResourceOwnerPointerEx", address: 0, enabled: false },
    EptHookDb { name: "ExSetTimer", address: 0, enabled: false },
    EptHookDb { name: "ExSetTimerResolution", address: 0, enabled: false },
    EptHookDb { name: "ExSizeOfRundownProtectionCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExSystemTimeToLocalTime", address: 0, enabled: false },
    EptHookDb { name: "ExTryAcquirePushLockExclusiveEx", address: 0, enabled: false },
    EptHookDb { name: "ExTryAcquirePushLockSharedEx", address: 0, enabled: false },
    EptHookDb { name: "ExTryAcquireSpinLockExclusiveAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExTryAcquireSpinLockSharedAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "ExTryConvertSharedSpinLockExclusive", address: 0, enabled: false },
    EptHookDb { name: "ExTryToAcquireFastMutex", address: 0, enabled: false },
    EptHookDb { name: "ExUnregisterCallback", address: 0, enabled: false },
    EptHookDb { name: "ExUuidCreate", address: 0, enabled: false },
    EptHookDb { name: "ExVerifySuite", address: 0, enabled: false },
    EptHookDb { name: "ExWaitForRundownProtectionRelease", address: 0, enabled: false },
    EptHookDb { name: "ExWaitForRundownProtectionReleaseCacheAware", address: 0, enabled: false },
    EptHookDb { name: "ExpInterlockedFlushSList", address: 0, enabled: false },
    EptHookDb { name: "ExpInterlockedPopEntrySList", address: 0, enabled: false },
    EptHookDb { name: "ExpInterlockedPushEntrySList", address: 0, enabled: false },
    EptHookDb { name: "ExportSecurityContext", address: 0, enabled: false },
    EptHookDb { name: "IoAcquireCancelSpinLock", address: 0, enabled: false },
    EptHookDb { name: "IoAcquireKsrPersistentMemory", address: 0, enabled: false },
    EptHookDb { name: "IoAcquireKsrPersistentMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "IoAcquireRemoveLockEx", address: 0, enabled: false },
    EptHookDb { name: "IoAcquireVpbSpinLock", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateController", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateDriverObjectExtension", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateErrorLogEntry", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateIrp", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateIrpEx", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateMdl", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateMiniCompletionPacket", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateSfioStreamIdentifier", address: 0, enabled: false },
    EptHookDb { name: "IoAllocateWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoApplyPriorityInfoThread", address: 0, enabled: false },
    EptHookDb { name: "IoAssignResources", address: 0, enabled: false },
    EptHookDb { name: "IoAttachDevice", address: 0, enabled: false },
    EptHookDb { name: "IoAttachDeviceByPointer", address: 0, enabled: false },
    EptHookDb { name: "IoAttachDeviceToDeviceStack", address: 0, enabled: false },
    EptHookDb { name: "IoAttachDeviceToDeviceStackSafe", address: 0, enabled: false },
    EptHookDb { name: "IoBuildAsynchronousFsdRequest", address: 0, enabled: false },
    EptHookDb { name: "IoBuildDeviceIoControlRequest", address: 0, enabled: false },
    EptHookDb { name: "IoBuildPartialMdl", address: 0, enabled: false },
    EptHookDb { name: "IoBuildSynchronousFsdRequest", address: 0, enabled: false },
    EptHookDb { name: "IoCancelFileOpen", address: 0, enabled: false },
    EptHookDb { name: "IoCancelIrp", address: 0, enabled: false },
    EptHookDb { name: "IoCheckDesiredAccess", address: 0, enabled: false },
    EptHookDb { name: "IoCheckEaBufferValidity", address: 0, enabled: false },
    EptHookDb { name: "IoCheckFileObjectOpenedAsCopyDestination", address: 0, enabled: false },
    EptHookDb { name: "IoCheckFileObjectOpenedAsCopySource", address: 0, enabled: false },
    EptHookDb { name: "IoCheckFunctionAccess", address: 0, enabled: false },
    EptHookDb { name: "IoCheckLinkShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoCheckQuerySetFileInformation", address: 0, enabled: false },
    EptHookDb { name: "IoCheckQuerySetVolumeInformation", address: 0, enabled: false },
    EptHookDb { name: "IoCheckQuotaBufferValidity", address: 0, enabled: false },
    EptHookDb { name: "IoCheckShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoCheckShareAccessEx", address: 0, enabled: false },
    EptHookDb { name: "IoCleanupIrp", address: 0, enabled: false },
    EptHookDb { name: "IoClearActivityIdThread", address: 0, enabled: false },
    EptHookDb { name: "IoClearFsTrackOffsetState", address: 0, enabled: false },
    EptHookDb { name: "IoClearIrpExtraCreateParameter", address: 0, enabled: false },
    EptHookDb { name: "IoConnectInterrupt", address: 0, enabled: false },
    EptHookDb { name: "IoConnectInterruptEx", address: 0, enabled: false },
    EptHookDb { name: "IoCreateController", address: 0, enabled: false },
    EptHookDb { name: "IoCreateDevice", address: 0, enabled: false },
    EptHookDb { name: "IoCreateDisk", address: 0, enabled: false },
    EptHookDb { name: "IoCreateDriverProxyExtension", address: 0, enabled: false },
    EptHookDb { name: "IoCreateFile", address: 0, enabled: false },
    EptHookDb { name: "IoCreateFileEx", address: 0, enabled: false },
    EptHookDb { name: "IoCreateFileSpecifyDeviceObjectHint", address: 0, enabled: false },
    EptHookDb { name: "IoCreateNotificationEvent", address: 0, enabled: false },
    EptHookDb { name: "IoCreateStreamFileObject", address: 0, enabled: false },
    EptHookDb { name: "IoCreateStreamFileObjectEx", address: 0, enabled: false },
    EptHookDb { name: "IoCreateStreamFileObjectEx2", address: 0, enabled: false },
    EptHookDb { name: "IoCreateStreamFileObjectLite", address: 0, enabled: false },
    EptHookDb { name: "IoCreateSymbolicLink", address: 0, enabled: false },
    EptHookDb { name: "IoCreateSynchronizationEvent", address: 0, enabled: false },
    EptHookDb { name: "IoCreateSystemThread", address: 0, enabled: false },
    EptHookDb { name: "IoCreateUnprotectedSymbolicLink", address: 0, enabled: false },
    EptHookDb { name: "IoCsqInitialize", address: 0, enabled: false },
    EptHookDb { name: "IoCsqInitializeEx", address: 0, enabled: false },
    EptHookDb { name: "IoCsqInsertIrp", address: 0, enabled: false },
    EptHookDb { name: "IoCsqInsertIrpEx", address: 0, enabled: false },
    EptHookDb { name: "IoCsqRemoveIrp", address: 0, enabled: false },
    EptHookDb { name: "IoCsqRemoveNextIrp", address: 0, enabled: false },
    EptHookDb { name: "IoDecrementKeepAliveCount", address: 0, enabled: false },
    EptHookDb { name: "IoDeleteController", address: 0, enabled: false },
    EptHookDb { name: "IoDeleteDevice", address: 0, enabled: false },
    EptHookDb { name: "IoDeleteSymbolicLink", address: 0, enabled: false },
    EptHookDb { name: "IoDetachDevice", address: 0, enabled: false },
    EptHookDb { name: "IoDisconnectInterrupt", address: 0, enabled: false },
    EptHookDb { name: "IoDisconnectInterruptEx", address: 0, enabled: false },
    EptHookDb { name: "IoDriverProxyCreateHotSwappableWorkerThread", address: 0, enabled: false },
    EptHookDb { name: "IoEnumerateDeviceObjectList", address: 0, enabled: false },
    EptHookDb { name: "IoEnumerateKsrPersistentMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "IoEnumerateRegisteredFiltersList", address: 0, enabled: false },
    EptHookDb { name: "IoFastQueryNetworkAttributes", address: 0, enabled: false },
    EptHookDb { name: "IoForwardIrpSynchronously", address: 0, enabled: false },
    EptHookDb { name: "IoFreeController", address: 0, enabled: false },
    EptHookDb { name: "IoFreeErrorLogEntry", address: 0, enabled: false },
    EptHookDb { name: "IoFreeIrp", address: 0, enabled: false },
    EptHookDb { name: "IoFreeKsrPersistentMemory", address: 0, enabled: false },
    EptHookDb { name: "IoFreeMdl", address: 0, enabled: false },
    EptHookDb { name: "IoFreeMiniCompletionPacket", address: 0, enabled: false },
    EptHookDb { name: "IoFreeSfioStreamIdentifier", address: 0, enabled: false },
    EptHookDb { name: "IoFreeWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoGetActivityIdIrp", address: 0, enabled: false },
    EptHookDb { name: "IoGetActivityIdThread", address: 0, enabled: false },
    EptHookDb { name: "IoGetAffinityInterrupt", address: 0, enabled: false },
    EptHookDb { name: "IoGetAttachedDevice", address: 0, enabled: false },
    EptHookDb { name: "IoGetAttachedDeviceReference", address: 0, enabled: false },
    EptHookDb { name: "IoGetBaseFileSystemDeviceObject", address: 0, enabled: false },
    EptHookDb { name: "IoGetBootDiskInformation", address: 0, enabled: false },
    EptHookDb { name: "IoGetBootDiskInformationLite", address: 0, enabled: false },
    EptHookDb { name: "IoGetConfigurationInformation", address: 0, enabled: false },
    EptHookDb { name: "IoGetContainerInformation", address: 0, enabled: false },
    EptHookDb { name: "IoGetCopyInformationExtension", address: 0, enabled: false },
    EptHookDb { name: "IoGetCurrentProcess", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceAttachmentBaseRef", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceDirectory", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceInterfaceAlias", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceInterfacePropertyData", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceInterfaces", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceNumaNode", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceObjectPointer", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceProperty", address: 0, enabled: false },
    EptHookDb { name: "IoGetDevicePropertyData", address: 0, enabled: false },
    EptHookDb { name: "IoGetDeviceToVerify", address: 0, enabled: false },
    EptHookDb { name: "IoGetDiskDeviceObject", address: 0, enabled: false },
    EptHookDb { name: "IoGetDmaAdapter", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverDirectory", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverObjectExtension", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverProxyEndpointWrapper", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverProxyExtensionFromDriverObject", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverProxyExtensionVersion", address: 0, enabled: false },
    EptHookDb { name: "IoGetDriverProxyFeatures", address: 0, enabled: false },
    EptHookDb { name: "IoGetFileObjectGenericMapping", address: 0, enabled: false },
    EptHookDb { name: "IoGetFsTrackOffsetState", address: 0, enabled: false },
    EptHookDb { name: "IoGetFsZeroingOffset", address: 0, enabled: false },
    EptHookDb { name: "IoGetInitialStack", address: 0, enabled: false },
    EptHookDb { name: "IoGetInitiatorProcess", address: 0, enabled: false },
    EptHookDb { name: "IoGetIoAttributionHandle", address: 0, enabled: false },
    EptHookDb { name: "IoGetIoPriorityHint", address: 0, enabled: false },
    EptHookDb { name: "IoGetIommuInterface", address: 0, enabled: false },
    EptHookDb { name: "IoGetIommuInterfaceEx", address: 0, enabled: false },
    EptHookDb { name: "IoGetIrpExtraCreateParameter", address: 0, enabled: false },
    EptHookDb { name: "IoGetKsrPersistentMemoryBuffer", address: 0, enabled: false },
    EptHookDb { name: "IoGetLowerDeviceObject", address: 0, enabled: false },
    EptHookDb { name: "IoGetOplockKeyContext", address: 0, enabled: false },
    EptHookDb { name: "IoGetOplockKeyContextEx", address: 0, enabled: false },
    EptHookDb { name: "IoGetPagingIoPriority", address: 0, enabled: false },
    EptHookDb { name: "IoGetRelatedDeviceObject", address: 0, enabled: false },
    EptHookDb { name: "IoGetRequestorProcess", address: 0, enabled: false },
    EptHookDb { name: "IoGetRequestorProcessId", address: 0, enabled: false },
    EptHookDb { name: "IoGetRequestorSessionId", address: 0, enabled: false },
    EptHookDb { name: "IoGetSfioStreamIdentifier", address: 0, enabled: false },
    EptHookDb { name: "IoGetShadowFileInformation", address: 0, enabled: false },
    EptHookDb { name: "IoGetSilo", address: 0, enabled: false },
    EptHookDb { name: "IoGetSiloParameters", address: 0, enabled: false },
    EptHookDb { name: "IoGetStackLimits", address: 0, enabled: false },
    EptHookDb { name: "IoGetTopLevelIrp", address: 0, enabled: false },
    EptHookDb { name: "IoGetTransactionParameterBlock", address: 0, enabled: false },
    EptHookDb { name: "IoHotSwapDriverProxyEndpoints", address: 0, enabled: false },
    EptHookDb { name: "IoIncrementKeepAliveCount", address: 0, enabled: false },
    EptHookDb { name: "IoInitializeIrp", address: 0, enabled: false },
    EptHookDb { name: "IoInitializeIrpEx", address: 0, enabled: false },
    EptHookDb { name: "IoInitializeRemoveLockEx", address: 0, enabled: false },
    EptHookDb { name: "IoInitializeTimer", address: 0, enabled: false },
    EptHookDb { name: "IoInitializeWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoInvalidateDeviceRelations", address: 0, enabled: false },
    EptHookDb { name: "IoInvalidateDeviceState", address: 0, enabled: false },
    EptHookDb { name: "IoIrpHasFsTrackOffsetExtensionType", address: 0, enabled: false },
    EptHookDb { name: "IoIs32bitProcess", address: 0, enabled: false },
    EptHookDb { name: "IoIsFileObjectIgnoringSharing", address: 0, enabled: false },
    EptHookDb { name: "IoIsFileOriginRemote", address: 0, enabled: false },
    EptHookDb { name: "IoIsInitiator32bitProcess", address: 0, enabled: false },
    EptHookDb { name: "IoIsOperationSynchronous", address: 0, enabled: false },
    EptHookDb { name: "IoIsSystemThread", address: 0, enabled: false },
    EptHookDb { name: "IoIsValidIrpStatus", address: 0, enabled: false },
    EptHookDb { name: "IoIsValidNameGraftingBuffer", address: 0, enabled: false },
    EptHookDb { name: "IoIsWdmVersionAvailable", address: 0, enabled: false },
    EptHookDb { name: "IoMakeAssociatedIrp", address: 0, enabled: false },
    EptHookDb { name: "IoMakeAssociatedIrpEx", address: 0, enabled: false },
    EptHookDb { name: "IoMapKsrPersistentMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "IoOpenDeviceInterfaceRegistryKey", address: 0, enabled: false },
    EptHookDb { name: "IoOpenDeviceRegistryKey", address: 0, enabled: false },
    EptHookDb { name: "IoOpenDriverRegistryKey", address: 0, enabled: false },
    EptHookDb { name: "IoPageRead", address: 0, enabled: false },
    EptHookDb { name: "IoPropagateActivityIdToThread", address: 0, enabled: false },
    EptHookDb { name: "IoQueryDeviceDescription", address: 0, enabled: false },
    EptHookDb { name: "IoQueryDmaFeatureSupport", address: 0, enabled: false },
    EptHookDb { name: "IoQueryFileDosDeviceName", address: 0, enabled: false },
    EptHookDb { name: "IoQueryFileInformation", address: 0, enabled: false },
    EptHookDb { name: "IoQueryFullDriverPath", address: 0, enabled: false },
    EptHookDb { name: "IoQueryInformationByName", address: 0, enabled: false },
    EptHookDb { name: "IoQueryKsrPersistentMemorySize", address: 0, enabled: false },
    EptHookDb { name: "IoQueryKsrPersistentMemorySizeEx", address: 0, enabled: false },
    EptHookDb { name: "IoQueryVolumeInformation", address: 0, enabled: false },
    EptHookDb { name: "IoQueueThreadIrp", address: 0, enabled: false },
    EptHookDb { name: "IoQueueWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoQueueWorkItemEx", address: 0, enabled: false },
    EptHookDb { name: "IoRaiseHardError", address: 0, enabled: false },
    EptHookDb { name: "IoRaiseInformationalHardError", address: 0, enabled: false },
    EptHookDb { name: "IoReadDiskSignature", address: 0, enabled: false },
    EptHookDb { name: "IoReadPartitionTable", address: 0, enabled: false },
    EptHookDb { name: "IoReadPartitionTableEx", address: 0, enabled: false },
    EptHookDb { name: "IoRecordIoAttribution", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterBootDriverCallback", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterBootDriverReinitialization", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterContainerNotification", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterDeviceInterface", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterDriverProxyEndpoints", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterDriverReinitialization", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterFileSystem", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterFsRegistrationChange", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterFsRegistrationChangeMountAware", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterLastChanceShutdownNotification", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterPlugPlayNotification", address: 0, enabled: false },
    EptHookDb { name: "IoRegisterShutdownNotification", address: 0, enabled: false },
    EptHookDb { name: "IoReleaseCancelSpinLock", address: 0, enabled: false },
    EptHookDb { name: "IoReleaseRemoveLockAndWaitEx", address: 0, enabled: false },
    EptHookDb { name: "IoReleaseRemoveLockEx", address: 0, enabled: false },
    EptHookDb { name: "IoReleaseVpbSpinLock", address: 0, enabled: false },
    EptHookDb { name: "IoRemoveIoCompletion", address: 0, enabled: false },
    EptHookDb { name: "IoRemoveLinkShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoRemoveLinkShareAccessEx", address: 0, enabled: false },
    EptHookDb { name: "IoRemoveShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoReplaceFileObjectName", address: 0, enabled: false },
    EptHookDb { name: "IoReplacePartitionUnit", address: 0, enabled: false },
    EptHookDb { name: "IoReportDetectedDevice", address: 0, enabled: false },
    EptHookDb { name: "IoReportInterruptActive", address: 0, enabled: false },
    EptHookDb { name: "IoReportInterruptInactive", address: 0, enabled: false },
    EptHookDb { name: "IoReportResourceForDetection", address: 0, enabled: false },
    EptHookDb { name: "IoReportResourceUsage", address: 0, enabled: false },
    EptHookDb { name: "IoReportRootDevice", address: 0, enabled: false },
    EptHookDb { name: "IoReportTargetDeviceChange", address: 0, enabled: false },
    EptHookDb { name: "IoReportTargetDeviceChangeAsynchronous", address: 0, enabled: false },
    EptHookDb { name: "IoRequestDeviceEject", address: 0, enabled: false },
    EptHookDb { name: "IoRequestDeviceEjectEx", address: 0, enabled: false },
    EptHookDb { name: "IoRequestDeviceRemovalForReset", address: 0, enabled: false },
    EptHookDb { name: "IoReserveKsrPersistentMemory", address: 0, enabled: false },
    EptHookDb { name: "IoReserveKsrPersistentMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "IoRetrievePriorityInfo", address: 0, enabled: false },
    EptHookDb { name: "IoReuseIrp", address: 0, enabled: false },
    EptHookDb { name: "IoSetActivityIdIrp", address: 0, enabled: false },
    EptHookDb { name: "IoSetActivityIdThread", address: 0, enabled: false },
    EptHookDb { name: "IoSetCompletionRoutineEx", address: 0, enabled: false },
    EptHookDb { name: "IoSetDeviceInterfacePropertyData", address: 0, enabled: false },
    EptHookDb { name: "IoSetDeviceInterfaceState", address: 0, enabled: false },
    EptHookDb { name: "IoSetDevicePropertyData", address: 0, enabled: false },
    EptHookDb { name: "IoSetDeviceToVerify", address: 0, enabled: false },
    EptHookDb { name: "IoSetFileObjectIgnoreSharing", address: 0, enabled: false },
    EptHookDb { name: "IoSetFileOrigin", address: 0, enabled: false },
    EptHookDb { name: "IoSetFsTrackOffsetState", address: 0, enabled: false },
    EptHookDb { name: "IoSetFsZeroingOffset", address: 0, enabled: false },
    EptHookDb { name: "IoSetFsZeroingOffsetRequired", address: 0, enabled: false },
    EptHookDb { name: "IoSetHardErrorOrVerifyDevice", address: 0, enabled: false },
    EptHookDb { name: "IoSetInformation", address: 0, enabled: false },
    EptHookDb { name: "IoSetIoAttributionIrp", address: 0, enabled: false },
    EptHookDb { name: "IoSetIoCompletionEx", address: 0, enabled: false },
    EptHookDb { name: "IoSetIoPriorityHint", address: 0, enabled: false },
    EptHookDb { name: "IoSetIrpExtraCreateParameter", address: 0, enabled: false },
    EptHookDb { name: "IoSetLinkShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoSetMasterIrpStatus", address: 0, enabled: false },
    EptHookDb { name: "IoSetPartitionInformation", address: 0, enabled: false },
    EptHookDb { name: "IoSetPartitionInformationEx", address: 0, enabled: false },
    EptHookDb { name: "IoSetShadowFileInformation", address: 0, enabled: false },
    EptHookDb { name: "IoSetShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoSetShareAccessEx", address: 0, enabled: false },
    EptHookDb { name: "IoSetStartIoAttributes", address: 0, enabled: false },
    EptHookDb { name: "IoSetSystemPartition", address: 0, enabled: false },
    EptHookDb { name: "IoSetThreadHardErrorMode", address: 0, enabled: false },
    EptHookDb { name: "IoSetTopLevelIrp", address: 0, enabled: false },
    EptHookDb { name: "IoSizeOfIrpEx", address: 0, enabled: false },
    EptHookDb { name: "IoSizeofWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoStartNextPacket", address: 0, enabled: false },
    EptHookDb { name: "IoStartNextPacketByKey", address: 0, enabled: false },
    EptHookDb { name: "IoStartPacket", address: 0, enabled: false },
    EptHookDb { name: "IoStartTimer", address: 0, enabled: false },
    EptHookDb { name: "IoStopTimer", address: 0, enabled: false },
    EptHookDb { name: "IoSynchronousCallDriver", address: 0, enabled: false },
    EptHookDb { name: "IoSynchronousPageWrite", address: 0, enabled: false },
    EptHookDb { name: "IoThreadToProcess", address: 0, enabled: false },
    EptHookDb { name: "IoTransferActivityId", address: 0, enabled: false },
    EptHookDb { name: "IoTranslateBusAddress", address: 0, enabled: false },
    EptHookDb { name: "IoTryQueueWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoUninitializeWorkItem", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterBootDriverCallback", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterContainerNotification", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterFileSystem", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterFsRegistrationChange", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterPlugPlayNotification", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterPlugPlayNotificationEx", address: 0, enabled: false },
    EptHookDb { name: "IoUnregisterShutdownNotification", address: 0, enabled: false },
    EptHookDb { name: "IoUpdateLinkShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoUpdateLinkShareAccessEx", address: 0, enabled: false },
    EptHookDb { name: "IoUpdateShareAccess", address: 0, enabled: false },
    EptHookDb { name: "IoValidateDeviceIoControlAccess", address: 0, enabled: false },
    EptHookDb { name: "IoVerifyPartitionTable", address: 0, enabled: false },
    EptHookDb { name: "IoVerifyVolume", address: 0, enabled: false },
    EptHookDb { name: "IoVolumeDeviceNameToGuid", address: 0, enabled: false },
    EptHookDb { name: "IoVolumeDeviceNameToGuidPath", address: 0, enabled: false },
    EptHookDb { name: "IoVolumeDeviceToDosName", address: 0, enabled: false },
    EptHookDb { name: "IoVolumeDeviceToGuid", address: 0, enabled: false },
    EptHookDb { name: "IoVolumeDeviceToGuidPath", address: 0, enabled: false },
    EptHookDb { name: "IoWMIAllocateInstanceIds", address: 0, enabled: false },
    EptHookDb { name: "IoWMIDeviceObjectToInstanceName", address: 0, enabled: false },
    EptHookDb { name: "IoWMIDeviceObjectToProviderId", address: 0, enabled: false },
    EptHookDb { name: "IoWMIExecuteMethod", address: 0, enabled: false },
    EptHookDb { name: "IoWMIHandleToInstanceName", address: 0, enabled: false },
    EptHookDb { name: "IoWMIOpenBlock", address: 0, enabled: false },
    EptHookDb { name: "IoWMIQueryAllData", address: 0, enabled: false },
    EptHookDb { name: "IoWMIQueryAllDataMultiple", address: 0, enabled: false },
    EptHookDb { name: "IoWMIQuerySingleInstance", address: 0, enabled: false },
    EptHookDb { name: "IoWMIQuerySingleInstanceMultiple", address: 0, enabled: false },
    EptHookDb { name: "IoWMIRegistrationControl", address: 0, enabled: false },
    EptHookDb { name: "IoWMISetNotificationCallback", address: 0, enabled: false },
    EptHookDb { name: "IoWMISetSingleInstance", address: 0, enabled: false },
    EptHookDb { name: "IoWMISetSingleItem", address: 0, enabled: false },
    EptHookDb { name: "IoWMISuggestInstanceName", address: 0, enabled: false },
    EptHookDb { name: "IoWMIWriteEvent", address: 0, enabled: false },
    EptHookDb { name: "IoWithinStackLimits", address: 0, enabled: false },
    EptHookDb { name: "IoWriteErrorLogEntry", address: 0, enabled: false },
    EptHookDb { name: "IoWriteKsrPersistentMemory", address: 0, enabled: false },
    EptHookDb { name: "IoWritePartitionTable", address: 0, enabled: false },
    EptHookDb { name: "IoWritePartitionTableEx", address: 0, enabled: false },
    EptHookDb { name: "IofCallDriver", address: 0, enabled: false },
    EptHookDb { name: "IofCompleteRequest", address: 0, enabled: false },
    EptHookDb { name: "IofGetDriverProxyWrapperFromEndpoint", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireGuardedMutex", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireGuardedMutexUnsafe", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireInStackQueuedSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireInStackQueuedSpinLockAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireInStackQueuedSpinLockForDpc", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireInterruptSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireQueuedSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireSpinLockAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireSpinLockForDpc", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireSpinLockRaiseToDpc", address: 0, enabled: false },
    EptHookDb { name: "KeAcquireSpinLockRaiseToSynch", address: 0, enabled: false },
    EptHookDb { name: "KeAddTriageDumpDataBlock", address: 0, enabled: false },
    EptHookDb { name: "KeAreAllApcsDisabled", address: 0, enabled: false },
    EptHookDb { name: "KeAreApcsDisabled", address: 0, enabled: false },
    EptHookDb { name: "KeAttachProcess", address: 0, enabled: false },
    EptHookDb { name: "KeBugCheck", address: 0, enabled: false },
    EptHookDb { name: "KeBugCheckEx", address: 0, enabled: false },
    EptHookDb { name: "KeCancelTimer", address: 0, enabled: false },
    EptHookDb { name: "KeClearEvent", address: 0, enabled: false },
    EptHookDb { name: "KeConvertAuxiliaryCounterToPerformanceCounter", address: 0, enabled: false },
    EptHookDb { name: "KeConvertPerformanceCounterToAuxiliaryCounter", address: 0, enabled: false },
    EptHookDb { name: "KeDelayExecutionThread", address: 0, enabled: false },
    EptHookDb { name: "KeDeregisterBoundCallback", address: 0, enabled: false },
    EptHookDb { name: "KeDeregisterBugCheckCallback", address: 0, enabled: false },
    EptHookDb { name: "KeDeregisterBugCheckReasonCallback", address: 0, enabled: false },
    EptHookDb { name: "KeDeregisterNmiCallback", address: 0, enabled: false },
    EptHookDb { name: "KeDeregisterProcessorChangeCallback", address: 0, enabled: false },
    EptHookDb { name: "KeDetachProcess", address: 0, enabled: false },
    EptHookDb { name: "KeEnterCriticalRegion", address: 0, enabled: false },
    EptHookDb { name: "KeEnterGuardedRegion", address: 0, enabled: false },
    EptHookDb { name: "KeExpandKernelStackAndCallout", address: 0, enabled: false },
    EptHookDb { name: "KeExpandKernelStackAndCalloutEx", address: 0, enabled: false },
    EptHookDb { name: "KeFlushIoBuffers", address: 0, enabled: false },
    EptHookDb { name: "KeFlushQueuedDpcs", address: 0, enabled: false },
    EptHookDb { name: "KeFlushWriteBuffer", address: 0, enabled: false },
    EptHookDb { name: "KeGetCurrentIrql", address: 0, enabled: false },
    EptHookDb { name: "KeGetCurrentNodeNumber", address: 0, enabled: false },
    EptHookDb { name: "KeGetCurrentProcessorNumberEx", address: 0, enabled: false },
    EptHookDb { name: "KeGetProcessorIndexFromNumber", address: 0, enabled: false },
    EptHookDb { name: "KeGetProcessorNumberFromIndex", address: 0, enabled: false },
    EptHookDb { name: "KeGetRecommendedSharedDataAlignment", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeCrashDumpHeader", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeDpc", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeEvent", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeGuardedMutex", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeMutant", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeMutex", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeSemaphore", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeThreadedDpc", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeTimer", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeTimerEx", address: 0, enabled: false },
    EptHookDb { name: "KeInitializeTriageDumpDataArray", address: 0, enabled: false },
    EptHookDb { name: "KeInsertByKeyDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInsertDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInsertHeadQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInsertQueue", address: 0, enabled: false },
    EptHookDb { name: "KeInsertQueueDpc", address: 0, enabled: false },
    EptHookDb { name: "KeInvalidateAllCaches", address: 0, enabled: false },
    EptHookDb { name: "KeInvalidateRangeAllCaches", address: 0, enabled: false },
    EptHookDb { name: "KeIpiGenericCall", address: 0, enabled: false },
    EptHookDb { name: "KeIsExecutingDpc", address: 0, enabled: false },
    EptHookDb { name: "KeLeaveCriticalRegion", address: 0, enabled: false },
    EptHookDb { name: "KeLeaveGuardedRegion", address: 0, enabled: false },
    EptHookDb { name: "KeLowerIrql", address: 0, enabled: false },
    EptHookDb { name: "KePulseEvent", address: 0, enabled: false },
    EptHookDb { name: "KeQueryActiveGroupCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryActiveProcessorCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryActiveProcessorCountEx", address: 0, enabled: false },
    EptHookDb { name: "KeQueryActiveProcessors", address: 0, enabled: false },
    EptHookDb { name: "KeQueryAuxiliaryCounterFrequency", address: 0, enabled: false },
    EptHookDb { name: "KeQueryDpcWatchdogInformation", address: 0, enabled: false },
    EptHookDb { name: "KeQueryGroupAffinity", address: 0, enabled: false },
    EptHookDb { name: "KeQueryHardwareCounterConfiguration", address: 0, enabled: false },
    EptHookDb { name: "KeQueryHighestNodeNumber", address: 0, enabled: false },
    EptHookDb { name: "KeQueryInterruptTimePrecise", address: 0, enabled: false },
    EptHookDb { name: "KeQueryLogicalProcessorRelationship", address: 0, enabled: false },
    EptHookDb { name: "KeQueryMaximumGroupCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryMaximumProcessorCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryMaximumProcessorCountEx", address: 0, enabled: false },
    EptHookDb { name: "KeQueryNodeActiveAffinity", address: 0, enabled: false },
    EptHookDb { name: "KeQueryNodeActiveAffinity2", address: 0, enabled: false },
    EptHookDb { name: "KeQueryNodeActiveProcessorCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryNodeMaximumProcessorCount", address: 0, enabled: false },
    EptHookDb { name: "KeQueryOwnerMutant", address: 0, enabled: false },
    EptHookDb { name: "KeQueryPerformanceCounter", address: 0, enabled: false },
    EptHookDb { name: "KeQueryPriorityThread", address: 0, enabled: false },
    EptHookDb { name: "KeQueryRuntimeThread", address: 0, enabled: false },
    EptHookDb { name: "KeQuerySystemTimePrecise", address: 0, enabled: false },
    EptHookDb { name: "KeQueryTimeIncrement", address: 0, enabled: false },
    EptHookDb { name: "KeQueryTotalCycleTimeThread", address: 0, enabled: false },
    EptHookDb { name: "KeQueryUnbiasedInterruptTime", address: 0, enabled: false },
    EptHookDb { name: "KeQueryUnbiasedInterruptTimePrecise", address: 0, enabled: false },
    EptHookDb { name: "KeRcuReadLock", address: 0, enabled: false },
    EptHookDb { name: "KeRcuReadLockAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeRcuReadUnlock", address: 0, enabled: false },
    EptHookDb { name: "KeRcuSynchronize", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateEvent", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateMutant", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateMutex", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateQueue", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateSemaphore", address: 0, enabled: false },
    EptHookDb { name: "KeReadStateTimer", address: 0, enabled: false },
    EptHookDb { name: "KeRegisterBoundCallback", address: 0, enabled: false },
    EptHookDb { name: "KeRegisterBugCheckCallback", address: 0, enabled: false },
    EptHookDb { name: "KeRegisterBugCheckReasonCallback", address: 0, enabled: false },
    EptHookDb { name: "KeRegisterNmiCallback", address: 0, enabled: false },
    EptHookDb { name: "KeRegisterProcessorChangeCallback", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseGuardedMutex", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseGuardedMutexUnsafe", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseInStackQueuedSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseInStackQueuedSpinLockForDpc", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseInStackQueuedSpinLockFromDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseInterruptSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseMutant", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseMutex", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseQueuedSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseSemaphore", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseSpinLockForDpc", address: 0, enabled: false },
    EptHookDb { name: "KeReleaseSpinLockFromDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveByKeyDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveByKeyDeviceQueueIfBusy", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveEntryDeviceQueue", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveQueue", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveQueueDpc", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveQueueDpcEx", address: 0, enabled: false },
    EptHookDb { name: "KeRemoveQueueEx", address: 0, enabled: false },
    EptHookDb { name: "KeResetEvent", address: 0, enabled: false },
    EptHookDb { name: "KeRestoreExtendedProcessorState", address: 0, enabled: false },
    EptHookDb { name: "KeRevertToUserAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "KeRevertToUserAffinityThreadEx", address: 0, enabled: false },
    EptHookDb { name: "KeRevertToUserGroupAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "KeRundownQueue", address: 0, enabled: false },
    EptHookDb { name: "KeSaveExtendedProcessorState", address: 0, enabled: false },
    EptHookDb { name: "KeSetBasePriorityThread", address: 0, enabled: false },
    EptHookDb { name: "KeSetCoalescableTimer", address: 0, enabled: false },
    EptHookDb { name: "KeSetEvent", address: 0, enabled: false },
    EptHookDb { name: "KeSetHardwareCounterConfiguration", address: 0, enabled: false },
    EptHookDb { name: "KeSetIdealProcessorThread", address: 0, enabled: false },
    EptHookDb { name: "KeSetImportanceDpc", address: 0, enabled: false },
    EptHookDb { name: "KeSetKernelStackSwapEnable", address: 0, enabled: false },
    EptHookDb { name: "KeSetPriorityThread", address: 0, enabled: false },
    EptHookDb { name: "KeSetSystemAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "KeSetSystemAffinityThreadEx", address: 0, enabled: false },
    EptHookDb { name: "KeSetSystemGroupAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "KeSetTargetProcessorDpc", address: 0, enabled: false },
    EptHookDb { name: "KeSetTargetProcessorDpcEx", address: 0, enabled: false },
    EptHookDb { name: "KeSetTimer", address: 0, enabled: false },
    EptHookDb { name: "KeSetTimerEx", address: 0, enabled: false },
    EptHookDb { name: "KeShouldYieldProcessor", address: 0, enabled: false },
    EptHookDb { name: "KeSrcuAllocate", address: 0, enabled: false },
    EptHookDb { name: "KeSrcuFree", address: 0, enabled: false },
    EptHookDb { name: "KeSrcuReadLock", address: 0, enabled: false },
    EptHookDb { name: "KeSrcuReadUnlock", address: 0, enabled: false },
    EptHookDb { name: "KeSrcuSynchronize", address: 0, enabled: false },
    EptHookDb { name: "KeStackAttachProcess", address: 0, enabled: false },
    EptHookDb { name: "KeStallExecutionProcessor", address: 0, enabled: false },
    EptHookDb { name: "KeSynchronizeExecution", address: 0, enabled: false },
    EptHookDb { name: "KeTestSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeTryToAcquireGuardedMutex", address: 0, enabled: false },
    EptHookDb { name: "KeTryToAcquireQueuedSpinLock", address: 0, enabled: false },
    EptHookDb { name: "KeTryToAcquireSpinLockAtDpcLevel", address: 0, enabled: false },
    EptHookDb { name: "KeUnstackDetachProcess", address: 0, enabled: false },
    EptHookDb { name: "KeWaitForMultipleObjects", address: 0, enabled: false },
    EptHookDb { name: "KeWaitForSingleObject", address: 0, enabled: false },
    EptHookDb { name: "MmAddPhysicalMemory", address: 0, enabled: false },
    EptHookDb { name: "MmAddVerifierSpecialThunks", address: 0, enabled: false },
    EptHookDb { name: "MmAddVerifierThunks", address: 0, enabled: false },
    EptHookDb { name: "MmAdvanceMdl", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateContiguousMemory", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateContiguousMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateContiguousMemorySpecifyCache", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateContiguousMemorySpecifyCacheNode", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateContiguousNodeMemory", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateMappingAddress", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateMappingAddressEx", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateMdlForIoSpace", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateNodePagesForMdlEx", address: 0, enabled: false },
    EptHookDb { name: "MmAllocateNonCachedMemory", address: 0, enabled: false },
    EptHookDb { name: "MmAllocatePagesForMdl", address: 0, enabled: false },
    EptHookDb { name: "MmAllocatePagesForMdlEx", address: 0, enabled: false },
    EptHookDb { name: "MmAllocatePartitionNodePagesForMdlEx", address: 0, enabled: false },
    EptHookDb { name: "MmAreMdlPagesCached", address: 0, enabled: false },
    EptHookDb { name: "MmBuildMdlForNonPagedPool", address: 0, enabled: false },
    EptHookDb { name: "MmCanFileBeTruncated", address: 0, enabled: false },
    EptHookDb { name: "MmCopyMemory", address: 0, enabled: false },
    EptHookDb { name: "MmCreateMdl", address: 0, enabled: false },
    EptHookDb { name: "MmCreateMirror", address: 0, enabled: false },
    EptHookDb { name: "MmDoesFileHaveUserWritableReferences", address: 0, enabled: false },
    EptHookDb { name: "MmFlushImageSection", address: 0, enabled: false },
    EptHookDb { name: "MmForceSectionClosed", address: 0, enabled: false },
    EptHookDb { name: "MmForceSectionClosedEx", address: 0, enabled: false },
    EptHookDb { name: "MmFreeContiguousMemory", address: 0, enabled: false },
    EptHookDb { name: "MmFreeContiguousMemorySpecifyCache", address: 0, enabled: false },
    EptHookDb { name: "MmFreeMappingAddress", address: 0, enabled: false },
    EptHookDb { name: "MmFreeNonCachedMemory", address: 0, enabled: false },
    EptHookDb { name: "MmFreePagesFromMdl", address: 0, enabled: false },
    EptHookDb { name: "MmFreePagesFromMdlEx", address: 0, enabled: false },
    EptHookDb { name: "MmGetCacheAttribute", address: 0, enabled: false },
    EptHookDb { name: "MmGetCacheAttributeEx", address: 0, enabled: false },
    EptHookDb { name: "MmGetMaximumFileSectionSize", address: 0, enabled: false },
    EptHookDb { name: "MmGetModuleRoutineAddress", address: 0, enabled: false },
    EptHookDb { name: "MmGetPhysicalAddress", address: 0, enabled: false },
    EptHookDb { name: "MmGetPhysicalMemoryRanges", address: 0, enabled: false },
    EptHookDb { name: "MmGetPhysicalMemoryRangesEx", address: 0, enabled: false },
    EptHookDb { name: "MmGetPhysicalMemoryRangesEx2", address: 0, enabled: false },
    EptHookDb { name: "MmGetSystemRoutineAddress", address: 0, enabled: false },
    EptHookDb { name: "MmGetVirtualForPhysical", address: 0, enabled: false },
    EptHookDb { name: "MmIsAddressValid", address: 0, enabled: false },
    EptHookDb { name: "MmIsDriverSuspectForVerifier", address: 0, enabled: false },
    EptHookDb { name: "MmIsDriverVerifying", address: 0, enabled: false },
    EptHookDb { name: "MmIsDriverVerifyingByAddress", address: 0, enabled: false },
    EptHookDb { name: "MmIsFileSectionActive", address: 0, enabled: false },
    EptHookDb { name: "MmIsIoSpaceActive", address: 0, enabled: false },
    EptHookDb { name: "MmIsKernelAddress", address: 0, enabled: false },
    EptHookDb { name: "MmIsNonPagedSystemAddressValid", address: 0, enabled: false },
    EptHookDb { name: "MmIsRecursiveIoFault", address: 0, enabled: false },
    EptHookDb { name: "MmIsThisAnNtAsSystem", address: 0, enabled: false },
    EptHookDb { name: "MmIsUserAddress", address: 0, enabled: false },
    EptHookDb { name: "MmIsVerifierEnabled", address: 0, enabled: false },
    EptHookDb { name: "MmLockPagableDataSection", address: 0, enabled: false },
    EptHookDb { name: "MmLockPagableSectionByHandle", address: 0, enabled: false },
    EptHookDb { name: "MmMapIoSpace", address: 0, enabled: false },
    EptHookDb { name: "MmMapIoSpaceEx", address: 0, enabled: false },
    EptHookDb { name: "MmMapLockedPages", address: 0, enabled: false },
    EptHookDb { name: "MmMapLockedPagesSpecifyCache", address: 0, enabled: false },
    EptHookDb { name: "MmMapLockedPagesWithReservedMapping", address: 0, enabled: false },
    EptHookDb { name: "MmMapMdl", address: 0, enabled: false },
    EptHookDb { name: "MmMapMemoryDumpMdlEx", address: 0, enabled: false },
    EptHookDb { name: "MmMapUserAddressesToPage", address: 0, enabled: false },
    EptHookDb { name: "MmMapVideoDisplay", address: 0, enabled: false },
    EptHookDb { name: "MmMapViewInSessionSpace", address: 0, enabled: false },
    EptHookDb { name: "MmMapViewInSessionSpaceEx", address: 0, enabled: false },
    EptHookDb { name: "MmMapViewInSystemSpace", address: 0, enabled: false },
    EptHookDb { name: "MmMapViewInSystemSpaceEx", address: 0, enabled: false },
    EptHookDb { name: "MmMdlPageContentsState", address: 0, enabled: false },
    EptHookDb { name: "MmMdlPagesAreZero", address: 0, enabled: false },
    EptHookDb { name: "MmPageEntireDriver", address: 0, enabled: false },
    EptHookDb { name: "MmPrefetchPages", address: 0, enabled: false },
    EptHookDb { name: "MmProbeAndLockPages", address: 0, enabled: false },
    EptHookDb { name: "MmProbeAndLockPagesEx", address: 0, enabled: false },
    EptHookDb { name: "MmProbeAndLockProcessPages", address: 0, enabled: false },
    EptHookDb { name: "MmProbeAndLockSelectedPages", address: 0, enabled: false },
    EptHookDb { name: "MmProtectDriverSection", address: 0, enabled: false },
    EptHookDb { name: "MmProtectMdlSystemAddress", address: 0, enabled: false },
    EptHookDb { name: "MmQuerySystemSize", address: 0, enabled: false },
    EptHookDb { name: "MmRemovePhysicalMemory", address: 0, enabled: false },
    EptHookDb { name: "MmResetDriverPaging", address: 0, enabled: false },
    EptHookDb { name: "MmRotatePhysicalView", address: 0, enabled: false },
    EptHookDb { name: "MmSecureVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "MmSecureVirtualMemoryEx", address: 0, enabled: false },
    EptHookDb { name: "MmSetAddressRangeModified", address: 0, enabled: false },
    EptHookDb { name: "MmSetPermanentCacheAttribute", address: 0, enabled: false },
    EptHookDb { name: "MmSizeOfMdl", address: 0, enabled: false },
    EptHookDb { name: "MmUnlockPagableImageSection", address: 0, enabled: false },
    EptHookDb { name: "MmUnlockPages", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapIoSpace", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapLockedPages", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapReservedMapping", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapVideoDisplay", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapViewInSessionSpace", address: 0, enabled: false },
    EptHookDb { name: "MmUnmapViewInSystemSpace", address: 0, enabled: false },
    EptHookDb { name: "MmUnsecureVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "NtAccessCheckAndAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtAccessCheckByTypeAndAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtAccessCheckByTypeResultListAndAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtAccessCheckByTypeResultListAndAuditAlarmByHandle", address: 0, enabled: false },
    EptHookDb { name: "NtAdjustGroupsToken", address: 0, enabled: false },
    EptHookDb { name: "NtAdjustPrivilegesToken", address: 0, enabled: false },
    EptHookDb { name: "NtAllocateVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "NtClose", address: 0, enabled: false },
    EptHookDb { name: "NtCloseObjectAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtCommitComplete", address: 0, enabled: false },
    EptHookDb { name: "NtCommitEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtCommitTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtCopyFileChunk", address: 0, enabled: false },
    EptHookDb { name: "NtCreateEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtCreateFile", address: 0, enabled: false },
    EptHookDb { name: "NtCreateResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtCreateSection", address: 0, enabled: false },
    EptHookDb { name: "NtCreateSectionEx", address: 0, enabled: false },
    EptHookDb { name: "NtCreateTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtCreateTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtDeleteObjectAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtDeviceIoControlFile", address: 0, enabled: false },
    EptHookDb { name: "NtDuplicateToken", address: 0, enabled: false },
    EptHookDb { name: "NtEnumerateTransactionObject", address: 0, enabled: false },
    EptHookDb { name: "NtFilterToken", address: 0, enabled: false },
    EptHookDb { name: "NtFlushBuffersFileEx", address: 0, enabled: false },
    EptHookDb { name: "NtFreeVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "NtFsControlFile", address: 0, enabled: false },
    EptHookDb { name: "NtGetNotificationResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtImpersonateAnonymousToken", address: 0, enabled: false },
    EptHookDb { name: "NtLockFile", address: 0, enabled: false },
    EptHookDb { name: "NtManagePartition", address: 0, enabled: false },
    EptHookDb { name: "NtOpenEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtOpenFile", address: 0, enabled: false },
    EptHookDb { name: "NtOpenJobObjectToken", address: 0, enabled: false },
    EptHookDb { name: "NtOpenObjectAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtOpenProcess", address: 0, enabled: false },
    EptHookDb { name: "NtOpenProcessToken", address: 0, enabled: false },
    EptHookDb { name: "NtOpenProcessTokenEx", address: 0, enabled: false },
    EptHookDb { name: "NtOpenRegistryTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtOpenResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtOpenThreadToken", address: 0, enabled: false },
    EptHookDb { name: "NtOpenThreadTokenEx", address: 0, enabled: false },
    EptHookDb { name: "NtOpenTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtOpenTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtPowerInformation", address: 0, enabled: false },
    EptHookDb { name: "NtPrePrepareComplete", address: 0, enabled: false },
    EptHookDb { name: "NtPrePrepareEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtPrepareComplete", address: 0, enabled: false },
    EptHookDb { name: "NtPrepareEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtPrivilegeCheck", address: 0, enabled: false },
    EptHookDb { name: "NtPrivilegeObjectAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtPrivilegedServiceAuditAlarm", address: 0, enabled: false },
    EptHookDb { name: "NtPropagationComplete", address: 0, enabled: false },
    EptHookDb { name: "NtPropagationFailed", address: 0, enabled: false },
    EptHookDb { name: "NtQueryDirectoryFile", address: 0, enabled: false },
    EptHookDb { name: "NtQueryDirectoryFileEx", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationByName", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationToken", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtQueryInformationTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtQueryObject", address: 0, enabled: false },
    EptHookDb { name: "NtQueryQuotaInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtQuerySecurityObject", address: 0, enabled: false },
    EptHookDb { name: "NtQueryVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "NtQueryVolumeInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtReadFile", address: 0, enabled: false },
    EptHookDb { name: "NtReadOnlyEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtRecoverEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtRecoverResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtRecoverTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtRegisterProtocolAddressInformation", address: 0, enabled: false },
    EptHookDb { name: "NtRenameTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtRollbackComplete", address: 0, enabled: false },
    EptHookDb { name: "NtRollbackEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtRollbackRegistryTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtRollbackTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtRollforwardTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationEnlistment", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationResourceManager", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationThread", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationToken", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationTransaction", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationTransactionManager", address: 0, enabled: false },
    EptHookDb { name: "NtSetInformationVirtualMemory", address: 0, enabled: false },
    EptHookDb { name: "NtSetQuotaInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtSetSecurityObject", address: 0, enabled: false },
    EptHookDb { name: "NtSetVolumeInformationFile", address: 0, enabled: false },
    EptHookDb { name: "NtSinglePhaseReject", address: 0, enabled: false },
    EptHookDb { name: "NtUnlockFile", address: 0, enabled: false },
    EptHookDb { name: "NtWriteFile", address: 0, enabled: false },
    EptHookDb { name: "ObCloseHandle", address: 0, enabled: false },
    EptHookDb { name: "ObCloseHandleWithResult", address: 0, enabled: false },
    EptHookDb { name: "ObDereferenceObjectDeferDelete", address: 0, enabled: false },
    EptHookDb { name: "ObDereferenceObjectDeferDeleteWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObGetFilterVersion", address: 0, enabled: false },
    EptHookDb { name: "ObGetObjectSecurity", address: 0, enabled: false },
    EptHookDb { name: "ObInsertObject", address: 0, enabled: false },
    EptHookDb { name: "ObIsKernelHandle", address: 0, enabled: false },
    EptHookDb { name: "ObMakeTemporaryObject", address: 0, enabled: false },
    EptHookDb { name: "ObOpenObjectByPointer", address: 0, enabled: false },
    EptHookDb { name: "ObOpenObjectByPointerWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObQueryNameString", address: 0, enabled: false },
    EptHookDb { name: "ObQueryObjectAuditingByHandle", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectByHandle", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectByHandleWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectByPointer", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectByPointerWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectSafe", address: 0, enabled: false },
    EptHookDb { name: "ObReferenceObjectSafeWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObRegisterCallbacks", address: 0, enabled: false },
    EptHookDb { name: "ObReleaseObjectSecurity", address: 0, enabled: false },
    EptHookDb { name: "ObUnRegisterCallbacks", address: 0, enabled: false },
    EptHookDb { name: "ObfDereferenceObject", address: 0, enabled: false },
    EptHookDb { name: "ObfDereferenceObjectWithTag", address: 0, enabled: false },
    EptHookDb { name: "ObfReferenceObject", address: 0, enabled: false },
    EptHookDb { name: "ObfReferenceObjectWithTag", address: 0, enabled: false },
    EptHookDb { name: "PsAcquireSiloHardReference", address: 0, enabled: false },
    EptHookDb { name: "PsAllocSiloContextSlot", address: 0, enabled: false },
    EptHookDb { name: "PsAllocateAffinityToken", address: 0, enabled: false },
    EptHookDb { name: "PsAssignImpersonationToken", address: 0, enabled: false },
    EptHookDb { name: "PsAttachSiloToCurrentThread", address: 0, enabled: false },
    EptHookDb { name: "PsChargePoolQuota", address: 0, enabled: false },
    EptHookDb { name: "PsChargeProcessPoolQuota", address: 0, enabled: false },
    EptHookDb { name: "PsCreateSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsCreateSystemThread", address: 0, enabled: false },
    EptHookDb { name: "PsDereferenceImpersonationToken", address: 0, enabled: false },
    EptHookDb { name: "PsDereferencePrimaryToken", address: 0, enabled: false },
    EptHookDb { name: "PsDereferenceSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsDetachSiloFromCurrentThread", address: 0, enabled: false },
    EptHookDb { name: "PsDisableImpersonation", address: 0, enabled: false },
    EptHookDb { name: "PsFreeAffinityToken", address: 0, enabled: false },
    EptHookDb { name: "PsFreeSiloContextSlot", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentProcessId", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentServerSiloName", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentThreadId", address: 0, enabled: false },
    EptHookDb { name: "PsGetCurrentThreadTeb", address: 0, enabled: false },
    EptHookDb { name: "PsGetEffectiveServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetHostSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetJobServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetJobSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetParentSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetPermanentSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsGetProcessCreateTimeQuadPart", address: 0, enabled: false },
    EptHookDb { name: "PsGetProcessExitStatus", address: 0, enabled: false },
    EptHookDb { name: "PsGetProcessExitTime", address: 0, enabled: false },
    EptHookDb { name: "PsGetProcessId", address: 0, enabled: false },
    EptHookDb { name: "PsGetProcessStartKey", address: 0, enabled: false },
    EptHookDb { name: "PsGetServerSiloActiveConsoleId", address: 0, enabled: false },
    EptHookDb { name: "PsGetServerSiloServiceSessionId", address: 0, enabled: false },
    EptHookDb { name: "PsGetSiloContainerId", address: 0, enabled: false },
    EptHookDb { name: "PsGetSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsGetSiloMonitorContextSlot", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadCreateTime", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadExitStatus", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadId", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadProcess", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadProcessId", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadProperty", address: 0, enabled: false },
    EptHookDb { name: "PsGetThreadServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsGetVersion", address: 0, enabled: false },
    EptHookDb { name: "PsImpersonateClient", address: 0, enabled: false },
    EptHookDb { name: "PsInsertPermanentSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsInsertSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsIsCurrentThreadInServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsIsCurrentThreadPrefetching", address: 0, enabled: false },
    EptHookDb { name: "PsIsDiskCountersEnabled", address: 0, enabled: false },
    EptHookDb { name: "PsIsHostSilo", address: 0, enabled: false },
    EptHookDb { name: "PsIsSystemThread", address: 0, enabled: false },
    EptHookDb { name: "PsIsThreadAttachedToSpecificSilo", address: 0, enabled: false },
    EptHookDb { name: "PsIsThreadTerminating", address: 0, enabled: false },
    EptHookDb { name: "PsLookupProcessByProcessId", address: 0, enabled: false },
    EptHookDb { name: "PsLookupThreadByThreadId", address: 0, enabled: false },
    EptHookDb { name: "PsMakeSiloContextPermanent", address: 0, enabled: false },
    EptHookDb { name: "PsQueryProcessAvailableCpus", address: 0, enabled: false },
    EptHookDb { name: "PsQueryProcessAvailableCpusCount", address: 0, enabled: false },
    EptHookDb { name: "PsQuerySystemAvailableCpus", address: 0, enabled: false },
    EptHookDb { name: "PsQuerySystemAvailableCpusCount", address: 0, enabled: false },
    EptHookDb { name: "PsQueryTotalCycleTimeProcess", address: 0, enabled: false },
    EptHookDb { name: "PsReferenceImpersonationToken", address: 0, enabled: false },
    EptHookDb { name: "PsReferencePrimaryToken", address: 0, enabled: false },
    EptHookDb { name: "PsReferenceSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsRegisterProcessAvailableCpusChangeNotification", address: 0, enabled: false },
    EptHookDb { name: "PsRegisterSiloMonitor", address: 0, enabled: false },
    EptHookDb { name: "PsRegisterSystemAvailableCpusChangeNotification", address: 0, enabled: false },
    EptHookDb { name: "PsReleaseSiloHardReference", address: 0, enabled: false },
    EptHookDb { name: "PsRemoveCreateThreadNotifyRoutine", address: 0, enabled: false },
    EptHookDb { name: "PsRemoveLoadImageNotifyRoutine", address: 0, enabled: false },
    EptHookDb { name: "PsRemoveSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsReplaceSiloContext", address: 0, enabled: false },
    EptHookDb { name: "PsRestoreImpersonation", address: 0, enabled: false },
    EptHookDb { name: "PsReturnPoolQuota", address: 0, enabled: false },
    EptHookDb { name: "PsRevertToSelf", address: 0, enabled: false },
    EptHookDb { name: "PsRevertToUserMultipleGroupAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "PsSetCreateProcessNotifyRoutine", address: 0, enabled: false },
    EptHookDb { name: "PsSetCreateProcessNotifyRoutineEx", address: 0, enabled: false },
    EptHookDb { name: "PsSetCreateProcessNotifyRoutineEx2", address: 0, enabled: false },
    EptHookDb { name: "PsSetCreateThreadNotifyRoutine", address: 0, enabled: false },
    EptHookDb { name: "PsSetCreateThreadNotifyRoutineEx", address: 0, enabled: false },
    EptHookDb { name: "PsSetCurrentThreadPrefetching", address: 0, enabled: false },
    EptHookDb { name: "PsSetLoadImageNotifyRoutine", address: 0, enabled: false },
    EptHookDb { name: "PsSetLoadImageNotifyRoutineEx", address: 0, enabled: false },
    EptHookDb { name: "PsSetSystemMultipleGroupAffinityThread", address: 0, enabled: false },
    EptHookDb { name: "PsStartSiloMonitor", address: 0, enabled: false },
    EptHookDb { name: "PsTerminateServerSilo", address: 0, enabled: false },
    EptHookDb { name: "PsTerminateSystemThread", address: 0, enabled: false },
    EptHookDb { name: "PsUnregisterAvailableCpusChangeNotification", address: 0, enabled: false },
    EptHookDb { name: "PsUnregisterSiloMonitor", address: 0, enabled: false },
    EptHookDb { name: "PsUpdateDiskCounters", address: 0, enabled: false },
    EptHookDb { name: "PsWrapApcWow64Thread", address: 0, enabled: false },
    EptHookDb { name: "PshedAllocateMemory", address: 0, enabled: false },
    EptHookDb { name: "PshedFreeMemory", address: 0, enabled: false },
    EptHookDb { name: "PshedIsSystemWheaEnabled", address: 0, enabled: false },
    EptHookDb { name: "PshedRegisterPlugin", address: 0, enabled: false },
    EptHookDb { name: "PshedSynchronizeExecution", address: 0, enabled: false },
    EptHookDb { name: "PshedUnregisterPlugin", address: 0, enabled: false },
    EptHookDb { name: "RtlAbsoluteToSelfRelativeSD", address: 0, enabled: false },
    EptHookDb { name: "RtlAddAccessAllowedAce", address: 0, enabled: false },
    EptHookDb { name: "RtlAddAccessAllowedAceEx", address: 0, enabled: false },
    EptHookDb { name: "RtlAddAce", address: 0, enabled: false },
    EptHookDb { name: "RtlAllocateAndInitializeSid", address: 0, enabled: false },
    EptHookDb { name: "RtlAllocateAndInitializeSidEx", address: 0, enabled: false },
    EptHookDb { name: "RtlAllocateHeap", address: 0, enabled: false },
    EptHookDb { name: "RtlAnsiStringToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlAppendStringToString", address: 0, enabled: false },
    EptHookDb { name: "RtlAppendUnicodeStringToString", address: 0, enabled: false },
    EptHookDb { name: "RtlAppendUnicodeToString", address: 0, enabled: false },
    EptHookDb { name: "RtlAreBitsClear", address: 0, enabled: false },
    EptHookDb { name: "RtlAreBitsSet", address: 0, enabled: false },
    EptHookDb { name: "RtlAssert", address: 0, enabled: false },
    EptHookDb { name: "RtlCaptureContext", address: 0, enabled: false },
    EptHookDb { name: "RtlCaptureContext2", address: 0, enabled: false },
    EptHookDb { name: "RtlCaptureStackBackTrace", address: 0, enabled: false },
    EptHookDb { name: "RtlCharToInteger", address: 0, enabled: false },
    EptHookDb { name: "RtlCheckRegistryKey", address: 0, enabled: false },
    EptHookDb { name: "RtlClearAllBits", address: 0, enabled: false },
    EptHookDb { name: "RtlClearBit", address: 0, enabled: false },
    EptHookDb { name: "RtlClearBits", address: 0, enabled: false },
    EptHookDb { name: "RtlCmDecodeMemIoResource", address: 0, enabled: false },
    EptHookDb { name: "RtlCmEncodeMemIoResource", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareAltitudes", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareMemoryUlong", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareString", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlCompareUnicodeStrings", address: 0, enabled: false },
    EptHookDb { name: "RtlCompressBuffer", address: 0, enabled: false },
    EptHookDb { name: "RtlCompressChunks", address: 0, enabled: false },
    EptHookDb { name: "RtlContractHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlConvertSidToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyBitMap", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyDeviceMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyLuid", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyMemoryNonTemporal", address: 0, enabled: false },
    EptHookDb { name: "RtlCopySid", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyString", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlCopyVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlCrc32", address: 0, enabled: false },
    EptHookDb { name: "RtlCrc64", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateAcl", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateHashTableEx", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateHeap", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateRegistryKey", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateSecurityDescriptorRelative", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateServiceSid", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateSystemVolumeInformationFolder", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlCreateVirtualAccountSid", address: 0, enabled: false },
    EptHookDb { name: "RtlCustomCPToUnicodeN", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressBuffer", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressBufferEx", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressBufferEx2", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressChunks", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressFragment", address: 0, enabled: false },
    EptHookDb { name: "RtlDecompressFragmentEx", address: 0, enabled: false },
    EptHookDb { name: "RtlDelete", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteAce", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteElementGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteElementGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteElementGenericTableAvlEx", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteNoSplay", address: 0, enabled: false },
    EptHookDb { name: "RtlDeleteRegistryValue", address: 0, enabled: false },
    EptHookDb { name: "RtlDescribeChunk", address: 0, enabled: false },
    EptHookDb { name: "RtlDestroyHeap", address: 0, enabled: false },
    EptHookDb { name: "RtlDowncaseUnicodeChar", address: 0, enabled: false },
    EptHookDb { name: "RtlDowncaseUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlDrainNonVolatileFlush", address: 0, enabled: false },
    EptHookDb { name: "RtlDuplicateUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlEndEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlEndStrongEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlEndWeakEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateGenericTableLikeADirectory", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateGenericTableWithoutSplaying", address: 0, enabled: false },
    EptHookDb { name: "RtlEnumerateGenericTableWithoutSplayingAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlEqualPrefixSid", address: 0, enabled: false },
    EptHookDb { name: "RtlEqualSid", address: 0, enabled: false },
    EptHookDb { name: "RtlEqualString", address: 0, enabled: false },
    EptHookDb { name: "RtlEqualUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlExpandHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlExtendCorrelationVector", address: 0, enabled: false },
    EptHookDb { name: "RtlExtractBitMap", address: 0, enabled: false },
    EptHookDb { name: "RtlFillMemoryNonTemporal", address: 0, enabled: false },
    EptHookDb { name: "RtlFillNonVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlFindClearBits", address: 0, enabled: false },
    EptHookDb { name: "RtlFindClearBitsAndSet", address: 0, enabled: false },
    EptHookDb { name: "RtlFindClearRuns", address: 0, enabled: false },
    EptHookDb { name: "RtlFindClosestEncodableLength", address: 0, enabled: false },
    EptHookDb { name: "RtlFindFirstRunClear", address: 0, enabled: false },
    EptHookDb { name: "RtlFindLastBackwardRunClear", address: 0, enabled: false },
    EptHookDb { name: "RtlFindLeastSignificantBit", address: 0, enabled: false },
    EptHookDb { name: "RtlFindLongestRunClear", address: 0, enabled: false },
    EptHookDb { name: "RtlFindMostSignificantBit", address: 0, enabled: false },
    EptHookDb { name: "RtlFindNextForwardRunClear", address: 0, enabled: false },
    EptHookDb { name: "RtlFindSetBits", address: 0, enabled: false },
    EptHookDb { name: "RtlFindSetBitsAndClear", address: 0, enabled: false },
    EptHookDb { name: "RtlFindUnicodePrefix", address: 0, enabled: false },
    EptHookDb { name: "RtlFlushNonVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlFlushNonVolatileMemoryRanges", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeAnsiString", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeHeap", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeNonVolatileToken", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeOemString", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeSid", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeUTF8String", address: 0, enabled: false },
    EptHookDb { name: "RtlFreeUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlGUIDFromString", address: 0, enabled: false },
    EptHookDb { name: "RtlGenerate8dot3Name", address: 0, enabled: false },
    EptHookDb { name: "RtlGenerateClass5Guid", address: 0, enabled: false },
    EptHookDb { name: "RtlGetAce", address: 0, enabled: false },
    EptHookDb { name: "RtlGetAcesBufferSize", address: 0, enabled: false },
    EptHookDb { name: "RtlGetActiveConsoleId", address: 0, enabled: false },
    EptHookDb { name: "RtlGetCallersAddress", address: 0, enabled: false },
    EptHookDb { name: "RtlGetCompressionWorkSpaceSize", address: 0, enabled: false },
    EptHookDb { name: "RtlGetConsoleSessionForegroundProcessId", address: 0, enabled: false },
    EptHookDb { name: "RtlGetDaclSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlGetElementGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlGetElementGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlGetEnabledExtendedFeatures", address: 0, enabled: false },
    EptHookDb { name: "RtlGetGroupSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlGetNextEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlGetNonVolatileToken", address: 0, enabled: false },
    EptHookDb { name: "RtlGetNtProductType", address: 0, enabled: false },
    EptHookDb { name: "RtlGetNtSystemRoot", address: 0, enabled: false },
    EptHookDb { name: "RtlGetOwnerSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlGetPersistedStateLocation", address: 0, enabled: false },
    EptHookDb { name: "RtlGetProductInfo", address: 0, enabled: false },
    EptHookDb { name: "RtlGetSaclSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlGetSuiteMask", address: 0, enabled: false },
    EptHookDb { name: "RtlGetSystemGlobalData", address: 0, enabled: false },
    EptHookDb { name: "RtlGetVersion", address: 0, enabled: false },
    EptHookDb { name: "RtlHashUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlIdentifierAuthoritySid", address: 0, enabled: false },
    EptHookDb { name: "RtlIdnToAscii", address: 0, enabled: false },
    EptHookDb { name: "RtlIdnToNameprepUnicode", address: 0, enabled: false },
    EptHookDb { name: "RtlIdnToUnicode", address: 0, enabled: false },
    EptHookDb { name: "RtlIncrementCorrelationVector", address: 0, enabled: false },
    EptHookDb { name: "RtlInitAnsiString", address: 0, enabled: false },
    EptHookDb { name: "RtlInitAnsiStringEx", address: 0, enabled: false },
    EptHookDb { name: "RtlInitCodePageTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInitEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInitString", address: 0, enabled: false },
    EptHookDb { name: "RtlInitStringEx", address: 0, enabled: false },
    EptHookDb { name: "RtlInitStrongEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInitUTF8String", address: 0, enabled: false },
    EptHookDb { name: "RtlInitUTF8StringEx", address: 0, enabled: false },
    EptHookDb { name: "RtlInitUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlInitUnicodeStringEx", address: 0, enabled: false },
    EptHookDb { name: "RtlInitWeakEnumerationHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeBitMap", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeCorrelationVector", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeSid", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeSidEx", address: 0, enabled: false },
    EptHookDb { name: "RtlInitializeUnicodePrefix", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertElementGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertElementGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertElementGenericTableFull", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertElementGenericTableFullAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlInsertUnicodePrefix", address: 0, enabled: false },
    EptHookDb { name: "RtlInt64ToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlIntegerToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlIoDecodeMemIoResource", address: 0, enabled: false },
    EptHookDb { name: "RtlIoEncodeMemIoResource", address: 0, enabled: false },
    EptHookDb { name: "RtlIsApiSetImplemented", address: 0, enabled: false },
    EptHookDb { name: "RtlIsCloudFilesPlaceholder", address: 0, enabled: false },
    EptHookDb { name: "RtlIsGenericTableEmpty", address: 0, enabled: false },
    EptHookDb { name: "RtlIsGenericTableEmptyAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlIsMultiSessionSku", address: 0, enabled: false },
    EptHookDb { name: "RtlIsMultiUsersInSessionSku", address: 0, enabled: false },
    EptHookDb { name: "RtlIsNameLegalDOS8Dot3", address: 0, enabled: false },
    EptHookDb { name: "RtlIsNonEmptyDirectoryReparsePointAllowed", address: 0, enabled: false },
    EptHookDb { name: "RtlIsNormalizedString", address: 0, enabled: false },
    EptHookDb { name: "RtlIsNtDdiVersionAvailable", address: 0, enabled: false },
    EptHookDb { name: "RtlIsPartialPlaceholder", address: 0, enabled: false },
    EptHookDb { name: "RtlIsPartialPlaceholderFileHandle", address: 0, enabled: false },
    EptHookDb { name: "RtlIsPartialPlaceholderFileInfo", address: 0, enabled: false },
    EptHookDb { name: "RtlIsSandboxedToken", address: 0, enabled: false },
    EptHookDb { name: "RtlIsSandboxedTokenHandle", address: 0, enabled: false },
    EptHookDb { name: "RtlIsServicePackVersionInstalled", address: 0, enabled: false },
    EptHookDb { name: "RtlIsStateSeparationEnabled", address: 0, enabled: false },
    EptHookDb { name: "RtlIsUntrustedObject", address: 0, enabled: false },
    EptHookDb { name: "RtlIsValidOemCharacter", address: 0, enabled: false },
    EptHookDb { name: "RtlIsZeroMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlLengthRequiredSid", address: 0, enabled: false },
    EptHookDb { name: "RtlLengthSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlLengthSid", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupElementGenericTable", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupElementGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupElementGenericTableFull", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupElementGenericTableFullAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlLookupFirstMatchingElementGenericTableAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlMapGenericMask", address: 0, enabled: false },
    EptHookDb { name: "RtlMoveVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlMultiByteToUnicodeN", address: 0, enabled: false },
    EptHookDb { name: "RtlMultiByteToUnicodeSize", address: 0, enabled: false },
    EptHookDb { name: "RtlNextUnicodePrefix", address: 0, enabled: false },
    EptHookDb { name: "RtlNormalizeSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlNormalizeString", address: 0, enabled: false },
    EptHookDb { name: "RtlNtStatusToDosError", address: 0, enabled: false },
    EptHookDb { name: "RtlNtStatusToDosErrorNoTeb", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberGenericTableElements", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberGenericTableElementsAvl", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberOfClearBits", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberOfClearBitsInRange", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberOfSetBits", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberOfSetBitsInRange", address: 0, enabled: false },
    EptHookDb { name: "RtlNumberOfSetBitsUlongPtr", address: 0, enabled: false },
    EptHookDb { name: "RtlOemStringToCountedUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlOemStringToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlOemToUnicodeN", address: 0, enabled: false },
    EptHookDb { name: "RtlOsDeploymentState", address: 0, enabled: false },
    EptHookDb { name: "RtlPrefetchMemoryNonTemporal", address: 0, enabled: false },
    EptHookDb { name: "RtlPrefixString", address: 0, enabled: false },
    EptHookDb { name: "RtlPrefixUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryPackageIdentity", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryPackageIdentityEx", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryProcessPlaceholderCompatibilityMode", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryRegistryValueWithFallback", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryRegistryValues", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryThreadPlaceholderCompatibilityMode", address: 0, enabled: false },
    EptHookDb { name: "RtlQueryValidationRunlevel", address: 0, enabled: false },
    EptHookDb { name: "RtlRaiseCustomSystemEventTrigger", address: 0, enabled: false },
    EptHookDb { name: "RtlRandom", address: 0, enabled: false },
    EptHookDb { name: "RtlRandomEx", address: 0, enabled: false },
    EptHookDb { name: "RtlRealPredecessor", address: 0, enabled: false },
    EptHookDb { name: "RtlRealSuccessor", address: 0, enabled: false },
    EptHookDb { name: "RtlRemoveEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlRemoveUnicodePrefix", address: 0, enabled: false },
    EptHookDb { name: "RtlReplaceSidInSd", address: 0, enabled: false },
    EptHookDb { name: "RtlReserveChunk", address: 0, enabled: false },
    EptHookDb { name: "RtlRunOnceBeginInitialize", address: 0, enabled: false },
    EptHookDb { name: "RtlRunOnceComplete", address: 0, enabled: false },
    EptHookDb { name: "RtlRunOnceExecuteOnce", address: 0, enabled: false },
    EptHookDb { name: "RtlRunOnceInitialize", address: 0, enabled: false },
    EptHookDb { name: "RtlSecondsSince1970ToTime", address: 0, enabled: false },
    EptHookDb { name: "RtlSecondsSince1980ToTime", address: 0, enabled: false },
    EptHookDb { name: "RtlSelfRelativeToAbsoluteSD", address: 0, enabled: false },
    EptHookDb { name: "RtlSetAllBits", address: 0, enabled: false },
    EptHookDb { name: "RtlSetBit", address: 0, enabled: false },
    EptHookDb { name: "RtlSetBits", address: 0, enabled: false },
    EptHookDb { name: "RtlSetDaclSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlSetGroupSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlSetOwnerSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlSetProcessPlaceholderCompatibilityMode", address: 0, enabled: false },
    EptHookDb { name: "RtlSetSystemGlobalData", address: 0, enabled: false },
    EptHookDb { name: "RtlSetThreadPlaceholderCompatibilityMode", address: 0, enabled: false },
    EptHookDb { name: "RtlSetVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlSplay", address: 0, enabled: false },
    EptHookDb { name: "RtlStringFromGUID", address: 0, enabled: false },
    EptHookDb { name: "RtlStronglyEnumerateEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlSubAuthorityCountSid", address: 0, enabled: false },
    EptHookDb { name: "RtlSubAuthoritySid", address: 0, enabled: false },
    EptHookDb { name: "RtlSubtreePredecessor", address: 0, enabled: false },
    EptHookDb { name: "RtlSubtreeSuccessor", address: 0, enabled: false },
    EptHookDb { name: "RtlSuffixUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlTestBit", address: 0, enabled: false },
    EptHookDb { name: "RtlTimeFieldsToTime", address: 0, enabled: false },
    EptHookDb { name: "RtlTimeToSecondsSince1970", address: 0, enabled: false },
    EptHookDb { name: "RtlTimeToSecondsSince1980", address: 0, enabled: false },
    EptHookDb { name: "RtlTimeToTimeFields", address: 0, enabled: false },
    EptHookDb { name: "RtlUTF8StringToUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlUTF8ToUnicodeN", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToAnsiString", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToCountedOemString", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToInt64", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToInteger", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToOemString", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeStringToUTF8String", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeToCustomCPN", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeToMultiByteN", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeToMultiByteSize", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeToOemN", address: 0, enabled: false },
    EptHookDb { name: "RtlUnicodeToUTF8N", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeChar", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeStringToCountedOemString", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeStringToOemString", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeToCustomCPN", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeToMultiByteN", address: 0, enabled: false },
    EptHookDb { name: "RtlUpcaseUnicodeToOemN", address: 0, enabled: false },
    EptHookDb { name: "RtlUpperChar", address: 0, enabled: false },
    EptHookDb { name: "RtlUpperString", address: 0, enabled: false },
    EptHookDb { name: "RtlValidRelativeSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlValidSecurityDescriptor", address: 0, enabled: false },
    EptHookDb { name: "RtlValidSid", address: 0, enabled: false },
    EptHookDb { name: "RtlValidateCorrelationVector", address: 0, enabled: false },
    EptHookDb { name: "RtlValidateUnicodeString", address: 0, enabled: false },
    EptHookDb { name: "RtlVerifyVersionInfo", address: 0, enabled: false },
    EptHookDb { name: "RtlVolumeDeviceToDosName", address: 0, enabled: false },
    EptHookDb { name: "RtlWalkFrameChain", address: 0, enabled: false },
    EptHookDb { name: "RtlWeaklyEnumerateEntryHashTable", address: 0, enabled: false },
    EptHookDb { name: "RtlWriteNonVolatileMemory", address: 0, enabled: false },
    EptHookDb { name: "RtlWriteRegistryValue", address: 0, enabled: false },
    EptHookDb { name: "RtlxAnsiStringToUnicodeSize", address: 0, enabled: false },
    EptHookDb { name: "RtlxOemStringToUnicodeSize", address: 0, enabled: false },
    EptHookDb { name: "RtlxUnicodeStringToAnsiSize", address: 0, enabled: false },
    EptHookDb { name: "RtlxUnicodeStringToOemSize", address: 0, enabled: false },
];

pub fn install_ept_hook_by_name(name: &str, process_id: ProcessId) -> Result<(), HookError> {
    match name {
        "DbgBreakPointWithStatus" => {
            let addr = DbgBreakPointWithStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgPrint" => {
            let addr = DbgPrint as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgPrintEx" => {
            let addr = DbgPrintEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgPrintReturnControlC" => {
            let addr = DbgPrintReturnControlC as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgPrompt" => {
            let addr = DbgPrompt as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgQueryDebugFilterState" => {
            let addr = DbgQueryDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgSetDebugFilterState" => {
            let addr = DbgSetDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "DbgSetDebugPrintCallback" => {
            let addr = DbgSetDebugPrintCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireFastMutex" => {
            let addr = ExAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireFastMutexUnsafe" => {
            let addr = ExAcquireFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquirePushLockExclusiveEx" => {
            let addr = ExAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquirePushLockSharedEx" => {
            let addr = ExAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireResourceExclusiveLite" => {
            let addr = ExAcquireResourceExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireResourceSharedLite" => {
            let addr = ExAcquireResourceSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireRundownProtection" => {
            let addr = ExAcquireRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireRundownProtectionCacheAware" => {
            let addr = ExAcquireRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireRundownProtectionCacheAwareEx" => {
            let addr = ExAcquireRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireRundownProtectionEx" => {
            let addr = ExAcquireRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSharedStarveExclusive" => {
            let addr = ExAcquireSharedStarveExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSharedWaitForExclusive" => {
            let addr = ExAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSpinLockExclusive" => {
            let addr = ExAcquireSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSpinLockShared" => {
            let addr = ExAcquireSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAdjustLookasideDepth" => {
            let addr = ExAdjustLookasideDepth as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocateCacheAwareRundownProtection" => {
            let addr = ExAllocateCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocateFromLookasideListEx" => {
            let addr = ExAllocateFromLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocateFromNPagedLookasideList" => {
            let addr = ExAllocateFromNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocateFromPagedLookasideList" => {
            let addr = ExAllocateFromPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocatePool2" => {
            let addr = ExAllocatePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocatePool3" => {
            let addr = ExAllocatePool3 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocatePoolWithQuota" => {
            let addr = ExAllocatePoolWithQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExAllocateTimer" => {
            let addr = ExAllocateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExCancelTimer" => {
            let addr = ExCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExCleanupRundownProtectionCacheAware" => {
            let addr = ExCleanupRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExConvertExclusiveToSharedLite" => {
            let addr = ExConvertExclusiveToSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExCreateCallback" => {
            let addr = ExCreateCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExCreatePool" => {
            let addr = ExCreatePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDeleteLookasideListEx" => {
            let addr = ExDeleteLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDeleteNPagedLookasideList" => {
            let addr = ExDeleteNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDeletePagedLookasideList" => {
            let addr = ExDeletePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDeleteResourceLite" => {
            let addr = ExDeleteResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDeleteTimer" => {
            let addr = ExDeleteTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDestroyPool" => {
            let addr = ExDestroyPool as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExDisableResourceBoostLite" => {
            let addr = ExDisableResourceBoostLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExEnterCriticalRegionAndAcquireResourceExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExEnterCriticalRegionAndAcquireResourceShared" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExEnumerateSystemFirmwareTables" => {
            let addr = ExEnumerateSystemFirmwareTables as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExExtendZone" => {
            let addr = ExExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFlushLookasideListEx" => {
            let addr = ExFlushLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreeCacheAwareRundownProtection" => {
            let addr = ExFreeCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreePool" => {
            let addr = ExFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreePool2" => {
            let addr = ExFreePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreePoolWithTag" => {
            let addr = ExFreePoolWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreeToLookasideListEx" => {
            let addr = ExFreeToLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreeToNPagedLookasideList" => {
            let addr = ExFreeToNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExFreeToPagedLookasideList" => {
            let addr = ExFreeToPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetExclusiveWaiterCount" => {
            let addr = ExGetExclusiveWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetFirmwareEnvironmentVariable" => {
            let addr = ExGetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetFirmwareType" => {
            let addr = ExGetFirmwareType as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetPreviousMode" => {
            let addr = ExGetPreviousMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetSharedWaiterCount" => {
            let addr = ExGetSharedWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExGetSystemFirmwareTable" => {
            let addr = ExGetSystemFirmwareTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeDeviceAts" => {
            let addr = ExInitializeDeviceAts as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeLookasideListEx" => {
            let addr = ExInitializeLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeNPagedLookasideList" => {
            let addr = ExInitializeNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializePagedLookasideList" => {
            let addr = ExInitializePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializePushLock" => {
            let addr = ExInitializePushLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeResourceLite" => {
            let addr = ExInitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeRundownProtection" => {
            let addr = ExInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeRundownProtectionCacheAware" => {
            let addr = ExInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeRundownProtectionCacheAwareEx" => {
            let addr = ExInitializeRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInitializeZone" => {
            let addr = ExInitializeZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedAddLargeInteger" => {
            let addr = ExInterlockedAddLargeInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedAddUlong" => {
            let addr = ExInterlockedAddUlong as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedExtendZone" => {
            let addr = ExInterlockedExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedInsertHeadList" => {
            let addr = ExInterlockedInsertHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedInsertTailList" => {
            let addr = ExInterlockedInsertTailList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedPopEntryList" => {
            let addr = ExInterlockedPopEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedPushEntryList" => {
            let addr = ExInterlockedPushEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExInterlockedRemoveHeadList" => {
            let addr = ExInterlockedRemoveHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExIsManufacturingModeEnabled" => {
            let addr = ExIsManufacturingModeEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExIsProcessorFeaturePresent" => {
            let addr = ExIsProcessorFeaturePresent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExIsResourceAcquiredExclusiveLite" => {
            let addr = ExIsResourceAcquiredExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExIsResourceAcquiredSharedLite" => {
            let addr = ExIsResourceAcquiredSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExIsSoftBoot" => {
            let addr = ExIsSoftBoot as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExLocalTimeToSystemTime" => {
            let addr = ExLocalTimeToSystemTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExNotifyCallback" => {
            let addr = ExNotifyCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExQueryDepthSList" => {
            let addr = ExQueryDepthSList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExQueryPoolBlockSize" => {
            let addr = ExQueryPoolBlockSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExQueryTimerResolution" => {
            let addr = ExQueryTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExQueueWorkItem" => {
            let addr = ExQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRaiseAccessViolation" => {
            let addr = ExRaiseAccessViolation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRaiseDatatypeMisalignment" => {
            let addr = ExRaiseDatatypeMisalignment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRaiseStatus" => {
            let addr = ExRaiseStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRcuFreePool" => {
            let addr = ExRcuFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReInitializeRundownProtection" => {
            let addr = ExReInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReInitializeRundownProtectionCacheAware" => {
            let addr = ExReInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRegisterCallback" => {
            let addr = ExRegisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReinitializeResourceLite" => {
            let addr = ExReinitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseFastMutex" => {
            let addr = ExReleaseFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseFastMutexUnsafe" => {
            let addr = ExReleaseFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleasePushLockExclusiveEx" => {
            let addr = ExReleasePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleasePushLockSharedEx" => {
            let addr = ExReleasePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseResourceAndLeaveCriticalRegion" => {
            let addr = ExReleaseResourceAndLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseResourceForThreadLite" => {
            let addr = ExReleaseResourceForThreadLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseResourceLite" => {
            let addr = ExReleaseResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseRundownProtection" => {
            let addr = ExReleaseRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseRundownProtectionCacheAware" => {
            let addr = ExReleaseRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseRundownProtectionCacheAwareEx" => {
            let addr = ExReleaseRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseRundownProtectionEx" => {
            let addr = ExReleaseRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseSpinLockExclusive" => {
            let addr = ExReleaseSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseSpinLockExclusiveFromDpcLevel" => {
            let addr = ExReleaseSpinLockExclusiveFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseSpinLockShared" => {
            let addr = ExReleaseSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExReleaseSpinLockSharedFromDpcLevel" => {
            let addr = ExReleaseSpinLockSharedFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRundownCompleted" => {
            let addr = ExRundownCompleted as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExRundownCompletedCacheAware" => {
            let addr = ExRundownCompletedCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSecurePoolUpdate" => {
            let addr = ExSecurePoolUpdate as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSecurePoolValidate" => {
            let addr = ExSecurePoolValidate as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSetFirmwareEnvironmentVariable" => {
            let addr = ExSetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSetResourceOwnerPointer" => {
            let addr = ExSetResourceOwnerPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSetResourceOwnerPointerEx" => {
            let addr = ExSetResourceOwnerPointerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSetTimer" => {
            let addr = ExSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSetTimerResolution" => {
            let addr = ExSetTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSizeOfRundownProtectionCacheAware" => {
            let addr = ExSizeOfRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExSystemTimeToLocalTime" => {
            let addr = ExSystemTimeToLocalTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryAcquirePushLockExclusiveEx" => {
            let addr = ExTryAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryAcquirePushLockSharedEx" => {
            let addr = ExTryAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryConvertSharedSpinLockExclusive" => {
            let addr = ExTryConvertSharedSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExTryToAcquireFastMutex" => {
            let addr = ExTryToAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExUnregisterCallback" => {
            let addr = ExUnregisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExUuidCreate" => {
            let addr = ExUuidCreate as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExVerifySuite" => {
            let addr = ExVerifySuite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExWaitForRundownProtectionRelease" => {
            let addr = ExWaitForRundownProtectionRelease as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExWaitForRundownProtectionReleaseCacheAware" => {
            let addr = ExWaitForRundownProtectionReleaseCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExpInterlockedFlushSList" => {
            let addr = ExpInterlockedFlushSList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExpInterlockedPopEntrySList" => {
            let addr = ExpInterlockedPopEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExpInterlockedPushEntrySList" => {
            let addr = ExpInterlockedPushEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ExportSecurityContext" => {
            let addr = ExportSecurityContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAcquireCancelSpinLock" => {
            let addr = IoAcquireCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAcquireKsrPersistentMemory" => {
            let addr = IoAcquireKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAcquireKsrPersistentMemoryEx" => {
            let addr = IoAcquireKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAcquireRemoveLockEx" => {
            let addr = IoAcquireRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAcquireVpbSpinLock" => {
            let addr = IoAcquireVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateController" => {
            let addr = IoAllocateController as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateDriverObjectExtension" => {
            let addr = IoAllocateDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateErrorLogEntry" => {
            let addr = IoAllocateErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateIrp" => {
            let addr = IoAllocateIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateIrpEx" => {
            let addr = IoAllocateIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateMdl" => {
            let addr = IoAllocateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateMiniCompletionPacket" => {
            let addr = IoAllocateMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateSfioStreamIdentifier" => {
            let addr = IoAllocateSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAllocateWorkItem" => {
            let addr = IoAllocateWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoApplyPriorityInfoThread" => {
            let addr = IoApplyPriorityInfoThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAssignResources" => {
            let addr = IoAssignResources as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAttachDevice" => {
            let addr = IoAttachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAttachDeviceByPointer" => {
            let addr = IoAttachDeviceByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAttachDeviceToDeviceStack" => {
            let addr = IoAttachDeviceToDeviceStack as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoAttachDeviceToDeviceStackSafe" => {
            let addr = IoAttachDeviceToDeviceStackSafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoBuildAsynchronousFsdRequest" => {
            let addr = IoBuildAsynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoBuildDeviceIoControlRequest" => {
            let addr = IoBuildDeviceIoControlRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoBuildPartialMdl" => {
            let addr = IoBuildPartialMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoBuildSynchronousFsdRequest" => {
            let addr = IoBuildSynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCancelFileOpen" => {
            let addr = IoCancelFileOpen as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCancelIrp" => {
            let addr = IoCancelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckDesiredAccess" => {
            let addr = IoCheckDesiredAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckEaBufferValidity" => {
            let addr = IoCheckEaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckFileObjectOpenedAsCopyDestination" => {
            let addr = IoCheckFileObjectOpenedAsCopyDestination as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckFileObjectOpenedAsCopySource" => {
            let addr = IoCheckFileObjectOpenedAsCopySource as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckFunctionAccess" => {
            let addr = IoCheckFunctionAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckLinkShareAccess" => {
            let addr = IoCheckLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckQuerySetFileInformation" => {
            let addr = IoCheckQuerySetFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckQuerySetVolumeInformation" => {
            let addr = IoCheckQuerySetVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckQuotaBufferValidity" => {
            let addr = IoCheckQuotaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckShareAccess" => {
            let addr = IoCheckShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCheckShareAccessEx" => {
            let addr = IoCheckShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCleanupIrp" => {
            let addr = IoCleanupIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoClearActivityIdThread" => {
            let addr = IoClearActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoClearFsTrackOffsetState" => {
            let addr = IoClearFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoClearIrpExtraCreateParameter" => {
            let addr = IoClearIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoConnectInterrupt" => {
            let addr = IoConnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoConnectInterruptEx" => {
            let addr = IoConnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateController" => {
            let addr = IoCreateController as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateDevice" => {
            let addr = IoCreateDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateDisk" => {
            let addr = IoCreateDisk as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateDriverProxyExtension" => {
            let addr = IoCreateDriverProxyExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateFile" => {
            let addr = IoCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateFileEx" => {
            let addr = IoCreateFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateFileSpecifyDeviceObjectHint" => {
            let addr = IoCreateFileSpecifyDeviceObjectHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateNotificationEvent" => {
            let addr = IoCreateNotificationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateStreamFileObject" => {
            let addr = IoCreateStreamFileObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateStreamFileObjectEx" => {
            let addr = IoCreateStreamFileObjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateStreamFileObjectEx2" => {
            let addr = IoCreateStreamFileObjectEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateStreamFileObjectLite" => {
            let addr = IoCreateStreamFileObjectLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateSymbolicLink" => {
            let addr = IoCreateSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateSynchronizationEvent" => {
            let addr = IoCreateSynchronizationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateSystemThread" => {
            let addr = IoCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCreateUnprotectedSymbolicLink" => {
            let addr = IoCreateUnprotectedSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqInitialize" => {
            let addr = IoCsqInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqInitializeEx" => {
            let addr = IoCsqInitializeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqInsertIrp" => {
            let addr = IoCsqInsertIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqInsertIrpEx" => {
            let addr = IoCsqInsertIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqRemoveIrp" => {
            let addr = IoCsqRemoveIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoCsqRemoveNextIrp" => {
            let addr = IoCsqRemoveNextIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDecrementKeepAliveCount" => {
            let addr = IoDecrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDeleteController" => {
            let addr = IoDeleteController as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDeleteDevice" => {
            let addr = IoDeleteDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDeleteSymbolicLink" => {
            let addr = IoDeleteSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDetachDevice" => {
            let addr = IoDetachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDisconnectInterrupt" => {
            let addr = IoDisconnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDisconnectInterruptEx" => {
            let addr = IoDisconnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoDriverProxyCreateHotSwappableWorkerThread" => {
            let addr = IoDriverProxyCreateHotSwappableWorkerThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoEnumerateDeviceObjectList" => {
            let addr = IoEnumerateDeviceObjectList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoEnumerateKsrPersistentMemoryEx" => {
            let addr = IoEnumerateKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoEnumerateRegisteredFiltersList" => {
            let addr = IoEnumerateRegisteredFiltersList as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFastQueryNetworkAttributes" => {
            let addr = IoFastQueryNetworkAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoForwardIrpSynchronously" => {
            let addr = IoForwardIrpSynchronously as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeController" => {
            let addr = IoFreeController as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeErrorLogEntry" => {
            let addr = IoFreeErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeIrp" => {
            let addr = IoFreeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeKsrPersistentMemory" => {
            let addr = IoFreeKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeMdl" => {
            let addr = IoFreeMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeMiniCompletionPacket" => {
            let addr = IoFreeMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeSfioStreamIdentifier" => {
            let addr = IoFreeSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoFreeWorkItem" => {
            let addr = IoFreeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetActivityIdIrp" => {
            let addr = IoGetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetActivityIdThread" => {
            let addr = IoGetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetAffinityInterrupt" => {
            let addr = IoGetAffinityInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetAttachedDevice" => {
            let addr = IoGetAttachedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetAttachedDeviceReference" => {
            let addr = IoGetAttachedDeviceReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetBaseFileSystemDeviceObject" => {
            let addr = IoGetBaseFileSystemDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetBootDiskInformation" => {
            let addr = IoGetBootDiskInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetBootDiskInformationLite" => {
            let addr = IoGetBootDiskInformationLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetConfigurationInformation" => {
            let addr = IoGetConfigurationInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetContainerInformation" => {
            let addr = IoGetContainerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetCopyInformationExtension" => {
            let addr = IoGetCopyInformationExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetCurrentProcess" => {
            let addr = IoGetCurrentProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceAttachmentBaseRef" => {
            let addr = IoGetDeviceAttachmentBaseRef as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceDirectory" => {
            let addr = IoGetDeviceDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceInterfaceAlias" => {
            let addr = IoGetDeviceInterfaceAlias as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceInterfacePropertyData" => {
            let addr = IoGetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceInterfaces" => {
            let addr = IoGetDeviceInterfaces as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceNumaNode" => {
            let addr = IoGetDeviceNumaNode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceObjectPointer" => {
            let addr = IoGetDeviceObjectPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceProperty" => {
            let addr = IoGetDeviceProperty as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDevicePropertyData" => {
            let addr = IoGetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDeviceToVerify" => {
            let addr = IoGetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDiskDeviceObject" => {
            let addr = IoGetDiskDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDmaAdapter" => {
            let addr = IoGetDmaAdapter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverDirectory" => {
            let addr = IoGetDriverDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverObjectExtension" => {
            let addr = IoGetDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverProxyEndpointWrapper" => {
            let addr = IoGetDriverProxyEndpointWrapper as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverProxyExtensionFromDriverObject" => {
            let addr = IoGetDriverProxyExtensionFromDriverObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverProxyExtensionVersion" => {
            let addr = IoGetDriverProxyExtensionVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetDriverProxyFeatures" => {
            let addr = IoGetDriverProxyFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetFileObjectGenericMapping" => {
            let addr = IoGetFileObjectGenericMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetFsTrackOffsetState" => {
            let addr = IoGetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetFsZeroingOffset" => {
            let addr = IoGetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetInitialStack" => {
            let addr = IoGetInitialStack as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetInitiatorProcess" => {
            let addr = IoGetInitiatorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetIoAttributionHandle" => {
            let addr = IoGetIoAttributionHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetIoPriorityHint" => {
            let addr = IoGetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetIommuInterface" => {
            let addr = IoGetIommuInterface as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetIommuInterfaceEx" => {
            let addr = IoGetIommuInterfaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetIrpExtraCreateParameter" => {
            let addr = IoGetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetKsrPersistentMemoryBuffer" => {
            let addr = IoGetKsrPersistentMemoryBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetLowerDeviceObject" => {
            let addr = IoGetLowerDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetOplockKeyContext" => {
            let addr = IoGetOplockKeyContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetOplockKeyContextEx" => {
            let addr = IoGetOplockKeyContextEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetPagingIoPriority" => {
            let addr = IoGetPagingIoPriority as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetRelatedDeviceObject" => {
            let addr = IoGetRelatedDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetRequestorProcess" => {
            let addr = IoGetRequestorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetRequestorProcessId" => {
            let addr = IoGetRequestorProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetRequestorSessionId" => {
            let addr = IoGetRequestorSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetSfioStreamIdentifier" => {
            let addr = IoGetSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetShadowFileInformation" => {
            let addr = IoGetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetSilo" => {
            let addr = IoGetSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetSiloParameters" => {
            let addr = IoGetSiloParameters as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetStackLimits" => {
            let addr = IoGetStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetTopLevelIrp" => {
            let addr = IoGetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoGetTransactionParameterBlock" => {
            let addr = IoGetTransactionParameterBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoHotSwapDriverProxyEndpoints" => {
            let addr = IoHotSwapDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIncrementKeepAliveCount" => {
            let addr = IoIncrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInitializeIrp" => {
            let addr = IoInitializeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInitializeIrpEx" => {
            let addr = IoInitializeIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInitializeRemoveLockEx" => {
            let addr = IoInitializeRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInitializeTimer" => {
            let addr = IoInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInitializeWorkItem" => {
            let addr = IoInitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInvalidateDeviceRelations" => {
            let addr = IoInvalidateDeviceRelations as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoInvalidateDeviceState" => {
            let addr = IoInvalidateDeviceState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIrpHasFsTrackOffsetExtensionType" => {
            let addr = IoIrpHasFsTrackOffsetExtensionType as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIs32bitProcess" => {
            let addr = IoIs32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsFileObjectIgnoringSharing" => {
            let addr = IoIsFileObjectIgnoringSharing as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsFileOriginRemote" => {
            let addr = IoIsFileOriginRemote as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsInitiator32bitProcess" => {
            let addr = IoIsInitiator32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsOperationSynchronous" => {
            let addr = IoIsOperationSynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsSystemThread" => {
            let addr = IoIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsValidIrpStatus" => {
            let addr = IoIsValidIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsValidNameGraftingBuffer" => {
            let addr = IoIsValidNameGraftingBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoIsWdmVersionAvailable" => {
            let addr = IoIsWdmVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoMakeAssociatedIrp" => {
            let addr = IoMakeAssociatedIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoMakeAssociatedIrpEx" => {
            let addr = IoMakeAssociatedIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoMapKsrPersistentMemoryEx" => {
            let addr = IoMapKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoOpenDeviceInterfaceRegistryKey" => {
            let addr = IoOpenDeviceInterfaceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoOpenDeviceRegistryKey" => {
            let addr = IoOpenDeviceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoOpenDriverRegistryKey" => {
            let addr = IoOpenDriverRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoPageRead" => {
            let addr = IoPageRead as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoPropagateActivityIdToThread" => {
            let addr = IoPropagateActivityIdToThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryDeviceDescription" => {
            let addr = IoQueryDeviceDescription as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryDmaFeatureSupport" => {
            let addr = IoQueryDmaFeatureSupport as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryFileDosDeviceName" => {
            let addr = IoQueryFileDosDeviceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryFileInformation" => {
            let addr = IoQueryFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryFullDriverPath" => {
            let addr = IoQueryFullDriverPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryInformationByName" => {
            let addr = IoQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryKsrPersistentMemorySize" => {
            let addr = IoQueryKsrPersistentMemorySize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryKsrPersistentMemorySizeEx" => {
            let addr = IoQueryKsrPersistentMemorySizeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueryVolumeInformation" => {
            let addr = IoQueryVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueueThreadIrp" => {
            let addr = IoQueueThreadIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueueWorkItem" => {
            let addr = IoQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoQueueWorkItemEx" => {
            let addr = IoQueueWorkItemEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRaiseHardError" => {
            let addr = IoRaiseHardError as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRaiseInformationalHardError" => {
            let addr = IoRaiseInformationalHardError as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReadDiskSignature" => {
            let addr = IoReadDiskSignature as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReadPartitionTable" => {
            let addr = IoReadPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReadPartitionTableEx" => {
            let addr = IoReadPartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRecordIoAttribution" => {
            let addr = IoRecordIoAttribution as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterBootDriverCallback" => {
            let addr = IoRegisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterBootDriverReinitialization" => {
            let addr = IoRegisterBootDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterContainerNotification" => {
            let addr = IoRegisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterDeviceInterface" => {
            let addr = IoRegisterDeviceInterface as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterDriverProxyEndpoints" => {
            let addr = IoRegisterDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterDriverReinitialization" => {
            let addr = IoRegisterDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterFileSystem" => {
            let addr = IoRegisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterFsRegistrationChange" => {
            let addr = IoRegisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterFsRegistrationChangeMountAware" => {
            let addr = IoRegisterFsRegistrationChangeMountAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterLastChanceShutdownNotification" => {
            let addr = IoRegisterLastChanceShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterPlugPlayNotification" => {
            let addr = IoRegisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRegisterShutdownNotification" => {
            let addr = IoRegisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReleaseCancelSpinLock" => {
            let addr = IoReleaseCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReleaseRemoveLockAndWaitEx" => {
            let addr = IoReleaseRemoveLockAndWaitEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReleaseRemoveLockEx" => {
            let addr = IoReleaseRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReleaseVpbSpinLock" => {
            let addr = IoReleaseVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRemoveIoCompletion" => {
            let addr = IoRemoveIoCompletion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRemoveLinkShareAccess" => {
            let addr = IoRemoveLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRemoveLinkShareAccessEx" => {
            let addr = IoRemoveLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRemoveShareAccess" => {
            let addr = IoRemoveShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReplaceFileObjectName" => {
            let addr = IoReplaceFileObjectName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReplacePartitionUnit" => {
            let addr = IoReplacePartitionUnit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportDetectedDevice" => {
            let addr = IoReportDetectedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportInterruptActive" => {
            let addr = IoReportInterruptActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportInterruptInactive" => {
            let addr = IoReportInterruptInactive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportResourceForDetection" => {
            let addr = IoReportResourceForDetection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportResourceUsage" => {
            let addr = IoReportResourceUsage as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportRootDevice" => {
            let addr = IoReportRootDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportTargetDeviceChange" => {
            let addr = IoReportTargetDeviceChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReportTargetDeviceChangeAsynchronous" => {
            let addr = IoReportTargetDeviceChangeAsynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRequestDeviceEject" => {
            let addr = IoRequestDeviceEject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRequestDeviceEjectEx" => {
            let addr = IoRequestDeviceEjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRequestDeviceRemovalForReset" => {
            let addr = IoRequestDeviceRemovalForReset as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReserveKsrPersistentMemory" => {
            let addr = IoReserveKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReserveKsrPersistentMemoryEx" => {
            let addr = IoReserveKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoRetrievePriorityInfo" => {
            let addr = IoRetrievePriorityInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoReuseIrp" => {
            let addr = IoReuseIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetActivityIdIrp" => {
            let addr = IoSetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetActivityIdThread" => {
            let addr = IoSetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetCompletionRoutineEx" => {
            let addr = IoSetCompletionRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetDeviceInterfacePropertyData" => {
            let addr = IoSetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetDeviceInterfaceState" => {
            let addr = IoSetDeviceInterfaceState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetDevicePropertyData" => {
            let addr = IoSetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetDeviceToVerify" => {
            let addr = IoSetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetFileObjectIgnoreSharing" => {
            let addr = IoSetFileObjectIgnoreSharing as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetFileOrigin" => {
            let addr = IoSetFileOrigin as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetFsTrackOffsetState" => {
            let addr = IoSetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetFsZeroingOffset" => {
            let addr = IoSetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetFsZeroingOffsetRequired" => {
            let addr = IoSetFsZeroingOffsetRequired as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetHardErrorOrVerifyDevice" => {
            let addr = IoSetHardErrorOrVerifyDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetInformation" => {
            let addr = IoSetInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetIoAttributionIrp" => {
            let addr = IoSetIoAttributionIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetIoCompletionEx" => {
            let addr = IoSetIoCompletionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetIoPriorityHint" => {
            let addr = IoSetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetIrpExtraCreateParameter" => {
            let addr = IoSetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetLinkShareAccess" => {
            let addr = IoSetLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetMasterIrpStatus" => {
            let addr = IoSetMasterIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetPartitionInformation" => {
            let addr = IoSetPartitionInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetPartitionInformationEx" => {
            let addr = IoSetPartitionInformationEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetShadowFileInformation" => {
            let addr = IoSetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetShareAccess" => {
            let addr = IoSetShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetShareAccessEx" => {
            let addr = IoSetShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetStartIoAttributes" => {
            let addr = IoSetStartIoAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetSystemPartition" => {
            let addr = IoSetSystemPartition as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetThreadHardErrorMode" => {
            let addr = IoSetThreadHardErrorMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSetTopLevelIrp" => {
            let addr = IoSetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSizeOfIrpEx" => {
            let addr = IoSizeOfIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSizeofWorkItem" => {
            let addr = IoSizeofWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoStartNextPacket" => {
            let addr = IoStartNextPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoStartNextPacketByKey" => {
            let addr = IoStartNextPacketByKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoStartPacket" => {
            let addr = IoStartPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoStartTimer" => {
            let addr = IoStartTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoStopTimer" => {
            let addr = IoStopTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSynchronousCallDriver" => {
            let addr = IoSynchronousCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoSynchronousPageWrite" => {
            let addr = IoSynchronousPageWrite as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoThreadToProcess" => {
            let addr = IoThreadToProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoTransferActivityId" => {
            let addr = IoTransferActivityId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoTranslateBusAddress" => {
            let addr = IoTranslateBusAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoTryQueueWorkItem" => {
            let addr = IoTryQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUninitializeWorkItem" => {
            let addr = IoUninitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterBootDriverCallback" => {
            let addr = IoUnregisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterContainerNotification" => {
            let addr = IoUnregisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterFileSystem" => {
            let addr = IoUnregisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterFsRegistrationChange" => {
            let addr = IoUnregisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterPlugPlayNotification" => {
            let addr = IoUnregisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterPlugPlayNotificationEx" => {
            let addr = IoUnregisterPlugPlayNotificationEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUnregisterShutdownNotification" => {
            let addr = IoUnregisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUpdateLinkShareAccess" => {
            let addr = IoUpdateLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUpdateLinkShareAccessEx" => {
            let addr = IoUpdateLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoUpdateShareAccess" => {
            let addr = IoUpdateShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoValidateDeviceIoControlAccess" => {
            let addr = IoValidateDeviceIoControlAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVerifyPartitionTable" => {
            let addr = IoVerifyPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVerifyVolume" => {
            let addr = IoVerifyVolume as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVolumeDeviceNameToGuid" => {
            let addr = IoVolumeDeviceNameToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVolumeDeviceNameToGuidPath" => {
            let addr = IoVolumeDeviceNameToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVolumeDeviceToDosName" => {
            let addr = IoVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVolumeDeviceToGuid" => {
            let addr = IoVolumeDeviceToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoVolumeDeviceToGuidPath" => {
            let addr = IoVolumeDeviceToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIAllocateInstanceIds" => {
            let addr = IoWMIAllocateInstanceIds as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIDeviceObjectToInstanceName" => {
            let addr = IoWMIDeviceObjectToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIDeviceObjectToProviderId" => {
            let addr = IoWMIDeviceObjectToProviderId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIExecuteMethod" => {
            let addr = IoWMIExecuteMethod as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIHandleToInstanceName" => {
            let addr = IoWMIHandleToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIOpenBlock" => {
            let addr = IoWMIOpenBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIQueryAllData" => {
            let addr = IoWMIQueryAllData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIQueryAllDataMultiple" => {
            let addr = IoWMIQueryAllDataMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIQuerySingleInstance" => {
            let addr = IoWMIQuerySingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIQuerySingleInstanceMultiple" => {
            let addr = IoWMIQuerySingleInstanceMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIRegistrationControl" => {
            let addr = IoWMIRegistrationControl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMISetNotificationCallback" => {
            let addr = IoWMISetNotificationCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMISetSingleInstance" => {
            let addr = IoWMISetSingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMISetSingleItem" => {
            let addr = IoWMISetSingleItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMISuggestInstanceName" => {
            let addr = IoWMISuggestInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWMIWriteEvent" => {
            let addr = IoWMIWriteEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWithinStackLimits" => {
            let addr = IoWithinStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWriteErrorLogEntry" => {
            let addr = IoWriteErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWriteKsrPersistentMemory" => {
            let addr = IoWriteKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWritePartitionTable" => {
            let addr = IoWritePartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IoWritePartitionTableEx" => {
            let addr = IoWritePartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IofCallDriver" => {
            let addr = IofCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IofCompleteRequest" => {
            let addr = IofCompleteRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "IofGetDriverProxyWrapperFromEndpoint" => {
            let addr = IofGetDriverProxyWrapperFromEndpoint as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireGuardedMutex" => {
            let addr = KeAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireGuardedMutexUnsafe" => {
            let addr = KeAcquireGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireInStackQueuedSpinLock" => {
            let addr = KeAcquireInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireInStackQueuedSpinLockAtDpcLevel" => {
            let addr = KeAcquireInStackQueuedSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireInStackQueuedSpinLockForDpc" => {
            let addr = KeAcquireInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireInterruptSpinLock" => {
            let addr = KeAcquireInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireQueuedSpinLock" => {
            let addr = KeAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireSpinLockAtDpcLevel" => {
            let addr = KeAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireSpinLockForDpc" => {
            let addr = KeAcquireSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireSpinLockRaiseToDpc" => {
            let addr = KeAcquireSpinLockRaiseToDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAcquireSpinLockRaiseToSynch" => {
            let addr = KeAcquireSpinLockRaiseToSynch as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAddTriageDumpDataBlock" => {
            let addr = KeAddTriageDumpDataBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAreAllApcsDisabled" => {
            let addr = KeAreAllApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAreApcsDisabled" => {
            let addr = KeAreApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeAttachProcess" => {
            let addr = KeAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeBugCheck" => {
            let addr = KeBugCheck as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeBugCheckEx" => {
            let addr = KeBugCheckEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeCancelTimer" => {
            let addr = KeCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeClearEvent" => {
            let addr = KeClearEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeConvertAuxiliaryCounterToPerformanceCounter" => {
            let addr = KeConvertAuxiliaryCounterToPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeConvertPerformanceCounterToAuxiliaryCounter" => {
            let addr = KeConvertPerformanceCounterToAuxiliaryCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDelayExecutionThread" => {
            let addr = KeDelayExecutionThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDeregisterBoundCallback" => {
            let addr = KeDeregisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDeregisterBugCheckCallback" => {
            let addr = KeDeregisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDeregisterBugCheckReasonCallback" => {
            let addr = KeDeregisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDeregisterNmiCallback" => {
            let addr = KeDeregisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDeregisterProcessorChangeCallback" => {
            let addr = KeDeregisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeDetachProcess" => {
            let addr = KeDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeEnterCriticalRegion" => {
            let addr = KeEnterCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeEnterGuardedRegion" => {
            let addr = KeEnterGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeExpandKernelStackAndCallout" => {
            let addr = KeExpandKernelStackAndCallout as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeExpandKernelStackAndCalloutEx" => {
            let addr = KeExpandKernelStackAndCalloutEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeFlushIoBuffers" => {
            let addr = KeFlushIoBuffers as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeFlushQueuedDpcs" => {
            let addr = KeFlushQueuedDpcs as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeFlushWriteBuffer" => {
            let addr = KeFlushWriteBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetCurrentIrql" => {
            let addr = KeGetCurrentIrql as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetCurrentNodeNumber" => {
            let addr = KeGetCurrentNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetCurrentProcessorNumberEx" => {
            let addr = KeGetCurrentProcessorNumberEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetProcessorIndexFromNumber" => {
            let addr = KeGetProcessorIndexFromNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetProcessorNumberFromIndex" => {
            let addr = KeGetProcessorNumberFromIndex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeGetRecommendedSharedDataAlignment" => {
            let addr = KeGetRecommendedSharedDataAlignment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeCrashDumpHeader" => {
            let addr = KeInitializeCrashDumpHeader as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeDeviceQueue" => {
            let addr = KeInitializeDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeDpc" => {
            let addr = KeInitializeDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeEvent" => {
            let addr = KeInitializeEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeGuardedMutex" => {
            let addr = KeInitializeGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeMutant" => {
            let addr = KeInitializeMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeMutex" => {
            let addr = KeInitializeMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeQueue" => {
            let addr = KeInitializeQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeSemaphore" => {
            let addr = KeInitializeSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeSpinLock" => {
            let addr = KeInitializeSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeThreadedDpc" => {
            let addr = KeInitializeThreadedDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeTimer" => {
            let addr = KeInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeTimerEx" => {
            let addr = KeInitializeTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInitializeTriageDumpDataArray" => {
            let addr = KeInitializeTriageDumpDataArray as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInsertByKeyDeviceQueue" => {
            let addr = KeInsertByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInsertDeviceQueue" => {
            let addr = KeInsertDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInsertHeadQueue" => {
            let addr = KeInsertHeadQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInsertQueue" => {
            let addr = KeInsertQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInsertQueueDpc" => {
            let addr = KeInsertQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInvalidateAllCaches" => {
            let addr = KeInvalidateAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeInvalidateRangeAllCaches" => {
            let addr = KeInvalidateRangeAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeIpiGenericCall" => {
            let addr = KeIpiGenericCall as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeIsExecutingDpc" => {
            let addr = KeIsExecutingDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeLeaveCriticalRegion" => {
            let addr = KeLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeLeaveGuardedRegion" => {
            let addr = KeLeaveGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeLowerIrql" => {
            let addr = KeLowerIrql as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KePulseEvent" => {
            let addr = KePulseEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryActiveGroupCount" => {
            let addr = KeQueryActiveGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryActiveProcessorCount" => {
            let addr = KeQueryActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryActiveProcessorCountEx" => {
            let addr = KeQueryActiveProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryActiveProcessors" => {
            let addr = KeQueryActiveProcessors as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryAuxiliaryCounterFrequency" => {
            let addr = KeQueryAuxiliaryCounterFrequency as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryDpcWatchdogInformation" => {
            let addr = KeQueryDpcWatchdogInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryGroupAffinity" => {
            let addr = KeQueryGroupAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryHardwareCounterConfiguration" => {
            let addr = KeQueryHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryHighestNodeNumber" => {
            let addr = KeQueryHighestNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryInterruptTimePrecise" => {
            let addr = KeQueryInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryLogicalProcessorRelationship" => {
            let addr = KeQueryLogicalProcessorRelationship as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryMaximumGroupCount" => {
            let addr = KeQueryMaximumGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryMaximumProcessorCount" => {
            let addr = KeQueryMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryMaximumProcessorCountEx" => {
            let addr = KeQueryMaximumProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryNodeActiveAffinity" => {
            let addr = KeQueryNodeActiveAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryNodeActiveAffinity2" => {
            let addr = KeQueryNodeActiveAffinity2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryNodeActiveProcessorCount" => {
            let addr = KeQueryNodeActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryNodeMaximumProcessorCount" => {
            let addr = KeQueryNodeMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryOwnerMutant" => {
            let addr = KeQueryOwnerMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryPerformanceCounter" => {
            let addr = KeQueryPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryPriorityThread" => {
            let addr = KeQueryPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryRuntimeThread" => {
            let addr = KeQueryRuntimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQuerySystemTimePrecise" => {
            let addr = KeQuerySystemTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryTimeIncrement" => {
            let addr = KeQueryTimeIncrement as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryTotalCycleTimeThread" => {
            let addr = KeQueryTotalCycleTimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryUnbiasedInterruptTime" => {
            let addr = KeQueryUnbiasedInterruptTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeQueryUnbiasedInterruptTimePrecise" => {
            let addr = KeQueryUnbiasedInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRcuReadLock" => {
            let addr = KeRcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRcuReadLockAtDpcLevel" => {
            let addr = KeRcuReadLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRcuReadUnlock" => {
            let addr = KeRcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRcuSynchronize" => {
            let addr = KeRcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateEvent" => {
            let addr = KeReadStateEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateMutant" => {
            let addr = KeReadStateMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateMutex" => {
            let addr = KeReadStateMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateQueue" => {
            let addr = KeReadStateQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateSemaphore" => {
            let addr = KeReadStateSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReadStateTimer" => {
            let addr = KeReadStateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRegisterBoundCallback" => {
            let addr = KeRegisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRegisterBugCheckCallback" => {
            let addr = KeRegisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRegisterBugCheckReasonCallback" => {
            let addr = KeRegisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRegisterNmiCallback" => {
            let addr = KeRegisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRegisterProcessorChangeCallback" => {
            let addr = KeRegisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseGuardedMutex" => {
            let addr = KeReleaseGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseGuardedMutexUnsafe" => {
            let addr = KeReleaseGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseInStackQueuedSpinLock" => {
            let addr = KeReleaseInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseInStackQueuedSpinLockForDpc" => {
            let addr = KeReleaseInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseInStackQueuedSpinLockFromDpcLevel" => {
            let addr = KeReleaseInStackQueuedSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseInterruptSpinLock" => {
            let addr = KeReleaseInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseMutant" => {
            let addr = KeReleaseMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseMutex" => {
            let addr = KeReleaseMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseQueuedSpinLock" => {
            let addr = KeReleaseQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseSemaphore" => {
            let addr = KeReleaseSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseSpinLock" => {
            let addr = KeReleaseSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseSpinLockForDpc" => {
            let addr = KeReleaseSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeReleaseSpinLockFromDpcLevel" => {
            let addr = KeReleaseSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveByKeyDeviceQueue" => {
            let addr = KeRemoveByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveByKeyDeviceQueueIfBusy" => {
            let addr = KeRemoveByKeyDeviceQueueIfBusy as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveDeviceQueue" => {
            let addr = KeRemoveDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveEntryDeviceQueue" => {
            let addr = KeRemoveEntryDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveQueue" => {
            let addr = KeRemoveQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveQueueDpc" => {
            let addr = KeRemoveQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveQueueDpcEx" => {
            let addr = KeRemoveQueueDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRemoveQueueEx" => {
            let addr = KeRemoveQueueEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeResetEvent" => {
            let addr = KeResetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRestoreExtendedProcessorState" => {
            let addr = KeRestoreExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRevertToUserAffinityThread" => {
            let addr = KeRevertToUserAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRevertToUserAffinityThreadEx" => {
            let addr = KeRevertToUserAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRevertToUserGroupAffinityThread" => {
            let addr = KeRevertToUserGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeRundownQueue" => {
            let addr = KeRundownQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSaveExtendedProcessorState" => {
            let addr = KeSaveExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetBasePriorityThread" => {
            let addr = KeSetBasePriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetCoalescableTimer" => {
            let addr = KeSetCoalescableTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetEvent" => {
            let addr = KeSetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetHardwareCounterConfiguration" => {
            let addr = KeSetHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetIdealProcessorThread" => {
            let addr = KeSetIdealProcessorThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetImportanceDpc" => {
            let addr = KeSetImportanceDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetKernelStackSwapEnable" => {
            let addr = KeSetKernelStackSwapEnable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetPriorityThread" => {
            let addr = KeSetPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetSystemAffinityThread" => {
            let addr = KeSetSystemAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetSystemAffinityThreadEx" => {
            let addr = KeSetSystemAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetSystemGroupAffinityThread" => {
            let addr = KeSetSystemGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetTargetProcessorDpc" => {
            let addr = KeSetTargetProcessorDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetTargetProcessorDpcEx" => {
            let addr = KeSetTargetProcessorDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetTimer" => {
            let addr = KeSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSetTimerEx" => {
            let addr = KeSetTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeShouldYieldProcessor" => {
            let addr = KeShouldYieldProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSrcuAllocate" => {
            let addr = KeSrcuAllocate as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSrcuFree" => {
            let addr = KeSrcuFree as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSrcuReadLock" => {
            let addr = KeSrcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSrcuReadUnlock" => {
            let addr = KeSrcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSrcuSynchronize" => {
            let addr = KeSrcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeStackAttachProcess" => {
            let addr = KeStackAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeStallExecutionProcessor" => {
            let addr = KeStallExecutionProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeSynchronizeExecution" => {
            let addr = KeSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeTestSpinLock" => {
            let addr = KeTestSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeTryToAcquireGuardedMutex" => {
            let addr = KeTryToAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeTryToAcquireQueuedSpinLock" => {
            let addr = KeTryToAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeTryToAcquireSpinLockAtDpcLevel" => {
            let addr = KeTryToAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeUnstackDetachProcess" => {
            let addr = KeUnstackDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeWaitForMultipleObjects" => {
            let addr = KeWaitForMultipleObjects as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "KeWaitForSingleObject" => {
            let addr = KeWaitForSingleObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAddPhysicalMemory" => {
            let addr = MmAddPhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAddVerifierSpecialThunks" => {
            let addr = MmAddVerifierSpecialThunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAddVerifierThunks" => {
            let addr = MmAddVerifierThunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAdvanceMdl" => {
            let addr = MmAdvanceMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateContiguousMemory" => {
            let addr = MmAllocateContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateContiguousMemoryEx" => {
            let addr = MmAllocateContiguousMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateContiguousMemorySpecifyCache" => {
            let addr = MmAllocateContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateContiguousMemorySpecifyCacheNode" => {
            let addr = MmAllocateContiguousMemorySpecifyCacheNode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateContiguousNodeMemory" => {
            let addr = MmAllocateContiguousNodeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateMappingAddress" => {
            let addr = MmAllocateMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateMappingAddressEx" => {
            let addr = MmAllocateMappingAddressEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateMdlForIoSpace" => {
            let addr = MmAllocateMdlForIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateNodePagesForMdlEx" => {
            let addr = MmAllocateNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocateNonCachedMemory" => {
            let addr = MmAllocateNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocatePagesForMdl" => {
            let addr = MmAllocatePagesForMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocatePagesForMdlEx" => {
            let addr = MmAllocatePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAllocatePartitionNodePagesForMdlEx" => {
            let addr = MmAllocatePartitionNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmAreMdlPagesCached" => {
            let addr = MmAreMdlPagesCached as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmBuildMdlForNonPagedPool" => {
            let addr = MmBuildMdlForNonPagedPool as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmCanFileBeTruncated" => {
            let addr = MmCanFileBeTruncated as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmCopyMemory" => {
            let addr = MmCopyMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmCreateMdl" => {
            let addr = MmCreateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmCreateMirror" => {
            let addr = MmCreateMirror as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmDoesFileHaveUserWritableReferences" => {
            let addr = MmDoesFileHaveUserWritableReferences as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFlushImageSection" => {
            let addr = MmFlushImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmForceSectionClosed" => {
            let addr = MmForceSectionClosed as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmForceSectionClosedEx" => {
            let addr = MmForceSectionClosedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreeContiguousMemory" => {
            let addr = MmFreeContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreeContiguousMemorySpecifyCache" => {
            let addr = MmFreeContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreeMappingAddress" => {
            let addr = MmFreeMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreeNonCachedMemory" => {
            let addr = MmFreeNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreePagesFromMdl" => {
            let addr = MmFreePagesFromMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmFreePagesFromMdlEx" => {
            let addr = MmFreePagesFromMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetCacheAttribute" => {
            let addr = MmGetCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetCacheAttributeEx" => {
            let addr = MmGetCacheAttributeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetMaximumFileSectionSize" => {
            let addr = MmGetMaximumFileSectionSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetModuleRoutineAddress" => {
            let addr = MmGetModuleRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetPhysicalAddress" => {
            let addr = MmGetPhysicalAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetPhysicalMemoryRanges" => {
            let addr = MmGetPhysicalMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetPhysicalMemoryRangesEx" => {
            let addr = MmGetPhysicalMemoryRangesEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetPhysicalMemoryRangesEx2" => {
            let addr = MmGetPhysicalMemoryRangesEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetSystemRoutineAddress" => {
            let addr = MmGetSystemRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmGetVirtualForPhysical" => {
            let addr = MmGetVirtualForPhysical as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsAddressValid" => {
            let addr = MmIsAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsDriverSuspectForVerifier" => {
            let addr = MmIsDriverSuspectForVerifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsDriverVerifying" => {
            let addr = MmIsDriverVerifying as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsDriverVerifyingByAddress" => {
            let addr = MmIsDriverVerifyingByAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsFileSectionActive" => {
            let addr = MmIsFileSectionActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsIoSpaceActive" => {
            let addr = MmIsIoSpaceActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsKernelAddress" => {
            let addr = MmIsKernelAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsNonPagedSystemAddressValid" => {
            let addr = MmIsNonPagedSystemAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsRecursiveIoFault" => {
            let addr = MmIsRecursiveIoFault as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsThisAnNtAsSystem" => {
            let addr = MmIsThisAnNtAsSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsUserAddress" => {
            let addr = MmIsUserAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmIsVerifierEnabled" => {
            let addr = MmIsVerifierEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmLockPagableDataSection" => {
            let addr = MmLockPagableDataSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmLockPagableSectionByHandle" => {
            let addr = MmLockPagableSectionByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapIoSpace" => {
            let addr = MmMapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapIoSpaceEx" => {
            let addr = MmMapIoSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapLockedPages" => {
            let addr = MmMapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapLockedPagesSpecifyCache" => {
            let addr = MmMapLockedPagesSpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapLockedPagesWithReservedMapping" => {
            let addr = MmMapLockedPagesWithReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapMdl" => {
            let addr = MmMapMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapMemoryDumpMdlEx" => {
            let addr = MmMapMemoryDumpMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapUserAddressesToPage" => {
            let addr = MmMapUserAddressesToPage as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapVideoDisplay" => {
            let addr = MmMapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapViewInSessionSpace" => {
            let addr = MmMapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapViewInSessionSpaceEx" => {
            let addr = MmMapViewInSessionSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapViewInSystemSpace" => {
            let addr = MmMapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMapViewInSystemSpaceEx" => {
            let addr = MmMapViewInSystemSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMdlPageContentsState" => {
            let addr = MmMdlPageContentsState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmMdlPagesAreZero" => {
            let addr = MmMdlPagesAreZero as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmPageEntireDriver" => {
            let addr = MmPageEntireDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmPrefetchPages" => {
            let addr = MmPrefetchPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProbeAndLockPages" => {
            let addr = MmProbeAndLockPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProbeAndLockPagesEx" => {
            let addr = MmProbeAndLockPagesEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProbeAndLockProcessPages" => {
            let addr = MmProbeAndLockProcessPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProbeAndLockSelectedPages" => {
            let addr = MmProbeAndLockSelectedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProtectDriverSection" => {
            let addr = MmProtectDriverSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmProtectMdlSystemAddress" => {
            let addr = MmProtectMdlSystemAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmQuerySystemSize" => {
            let addr = MmQuerySystemSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmRemovePhysicalMemory" => {
            let addr = MmRemovePhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmResetDriverPaging" => {
            let addr = MmResetDriverPaging as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmRotatePhysicalView" => {
            let addr = MmRotatePhysicalView as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmSecureVirtualMemory" => {
            let addr = MmSecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmSecureVirtualMemoryEx" => {
            let addr = MmSecureVirtualMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmSetAddressRangeModified" => {
            let addr = MmSetAddressRangeModified as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmSetPermanentCacheAttribute" => {
            let addr = MmSetPermanentCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmSizeOfMdl" => {
            let addr = MmSizeOfMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnlockPagableImageSection" => {
            let addr = MmUnlockPagableImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnlockPages" => {
            let addr = MmUnlockPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapIoSpace" => {
            let addr = MmUnmapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapLockedPages" => {
            let addr = MmUnmapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapReservedMapping" => {
            let addr = MmUnmapReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapVideoDisplay" => {
            let addr = MmUnmapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapViewInSessionSpace" => {
            let addr = MmUnmapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnmapViewInSystemSpace" => {
            let addr = MmUnmapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "MmUnsecureVirtualMemory" => {
            let addr = MmUnsecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAccessCheckAndAuditAlarm" => {
            let addr = NtAccessCheckAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAccessCheckByTypeAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarmByHandle" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarmByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAdjustGroupsToken" => {
            let addr = NtAdjustGroupsToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAdjustPrivilegesToken" => {
            let addr = NtAdjustPrivilegesToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtAllocateVirtualMemory" => {
            let addr = NtAllocateVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtClose" => {
            let addr = NtClose as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCloseObjectAuditAlarm" => {
            let addr = NtCloseObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCommitComplete" => {
            let addr = NtCommitComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCommitEnlistment" => {
            let addr = NtCommitEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCommitTransaction" => {
            let addr = NtCommitTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCopyFileChunk" => {
            let addr = NtCopyFileChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateEnlistment" => {
            let addr = NtCreateEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateFile" => {
            let addr = NtCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateResourceManager" => {
            let addr = NtCreateResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateSection" => {
            let addr = NtCreateSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateSectionEx" => {
            let addr = NtCreateSectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateTransaction" => {
            let addr = NtCreateTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtCreateTransactionManager" => {
            let addr = NtCreateTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtDeleteObjectAuditAlarm" => {
            let addr = NtDeleteObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtDeviceIoControlFile" => {
            let addr = NtDeviceIoControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtDuplicateToken" => {
            let addr = NtDuplicateToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtEnumerateTransactionObject" => {
            let addr = NtEnumerateTransactionObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtFilterToken" => {
            let addr = NtFilterToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtFlushBuffersFileEx" => {
            let addr = NtFlushBuffersFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtFreeVirtualMemory" => {
            let addr = NtFreeVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtFsControlFile" => {
            let addr = NtFsControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtGetNotificationResourceManager" => {
            let addr = NtGetNotificationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtImpersonateAnonymousToken" => {
            let addr = NtImpersonateAnonymousToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtLockFile" => {
            let addr = NtLockFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtManagePartition" => {
            let addr = NtManagePartition as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenEnlistment" => {
            let addr = NtOpenEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenFile" => {
            let addr = NtOpenFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenJobObjectToken" => {
            let addr = NtOpenJobObjectToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenObjectAuditAlarm" => {
            let addr = NtOpenObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenProcess" => {
            let addr = NtOpenProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenProcessToken" => {
            let addr = NtOpenProcessToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenProcessTokenEx" => {
            let addr = NtOpenProcessTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenRegistryTransaction" => {
            let addr = NtOpenRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenResourceManager" => {
            let addr = NtOpenResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenThreadToken" => {
            let addr = NtOpenThreadToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenThreadTokenEx" => {
            let addr = NtOpenThreadTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenTransaction" => {
            let addr = NtOpenTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtOpenTransactionManager" => {
            let addr = NtOpenTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPowerInformation" => {
            let addr = NtPowerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrePrepareComplete" => {
            let addr = NtPrePrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrePrepareEnlistment" => {
            let addr = NtPrePrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrepareComplete" => {
            let addr = NtPrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrepareEnlistment" => {
            let addr = NtPrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrivilegeCheck" => {
            let addr = NtPrivilegeCheck as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrivilegeObjectAuditAlarm" => {
            let addr = NtPrivilegeObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPrivilegedServiceAuditAlarm" => {
            let addr = NtPrivilegedServiceAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPropagationComplete" => {
            let addr = NtPropagationComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtPropagationFailed" => {
            let addr = NtPropagationFailed as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryDirectoryFile" => {
            let addr = NtQueryDirectoryFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryDirectoryFileEx" => {
            let addr = NtQueryDirectoryFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationByName" => {
            let addr = NtQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationEnlistment" => {
            let addr = NtQueryInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationFile" => {
            let addr = NtQueryInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationResourceManager" => {
            let addr = NtQueryInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationToken" => {
            let addr = NtQueryInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationTransaction" => {
            let addr = NtQueryInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryInformationTransactionManager" => {
            let addr = NtQueryInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryObject" => {
            let addr = NtQueryObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryQuotaInformationFile" => {
            let addr = NtQueryQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQuerySecurityObject" => {
            let addr = NtQuerySecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryVirtualMemory" => {
            let addr = NtQueryVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtQueryVolumeInformationFile" => {
            let addr = NtQueryVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtReadFile" => {
            let addr = NtReadFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtReadOnlyEnlistment" => {
            let addr = NtReadOnlyEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRecoverEnlistment" => {
            let addr = NtRecoverEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRecoverResourceManager" => {
            let addr = NtRecoverResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRecoverTransactionManager" => {
            let addr = NtRecoverTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRegisterProtocolAddressInformation" => {
            let addr = NtRegisterProtocolAddressInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRenameTransactionManager" => {
            let addr = NtRenameTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRollbackComplete" => {
            let addr = NtRollbackComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRollbackEnlistment" => {
            let addr = NtRollbackEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRollbackRegistryTransaction" => {
            let addr = NtRollbackRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRollbackTransaction" => {
            let addr = NtRollbackTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtRollforwardTransactionManager" => {
            let addr = NtRollforwardTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationEnlistment" => {
            let addr = NtSetInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationFile" => {
            let addr = NtSetInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationResourceManager" => {
            let addr = NtSetInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationThread" => {
            let addr = NtSetInformationThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationToken" => {
            let addr = NtSetInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationTransaction" => {
            let addr = NtSetInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationTransactionManager" => {
            let addr = NtSetInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetInformationVirtualMemory" => {
            let addr = NtSetInformationVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetQuotaInformationFile" => {
            let addr = NtSetQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetSecurityObject" => {
            let addr = NtSetSecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSetVolumeInformationFile" => {
            let addr = NtSetVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtSinglePhaseReject" => {
            let addr = NtSinglePhaseReject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtUnlockFile" => {
            let addr = NtUnlockFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "NtWriteFile" => {
            let addr = NtWriteFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObCloseHandle" => {
            let addr = ObCloseHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObCloseHandleWithResult" => {
            let addr = ObCloseHandleWithResult as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObDereferenceObjectDeferDelete" => {
            let addr = ObDereferenceObjectDeferDelete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObDereferenceObjectDeferDeleteWithTag" => {
            let addr = ObDereferenceObjectDeferDeleteWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObGetFilterVersion" => {
            let addr = ObGetFilterVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObGetObjectSecurity" => {
            let addr = ObGetObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObInsertObject" => {
            let addr = ObInsertObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObIsKernelHandle" => {
            let addr = ObIsKernelHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObMakeTemporaryObject" => {
            let addr = ObMakeTemporaryObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObOpenObjectByPointer" => {
            let addr = ObOpenObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObOpenObjectByPointerWithTag" => {
            let addr = ObOpenObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObQueryNameString" => {
            let addr = ObQueryNameString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObQueryObjectAuditingByHandle" => {
            let addr = ObQueryObjectAuditingByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectByHandle" => {
            let addr = ObReferenceObjectByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectByHandleWithTag" => {
            let addr = ObReferenceObjectByHandleWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectByPointer" => {
            let addr = ObReferenceObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectByPointerWithTag" => {
            let addr = ObReferenceObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectSafe" => {
            let addr = ObReferenceObjectSafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReferenceObjectSafeWithTag" => {
            let addr = ObReferenceObjectSafeWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObRegisterCallbacks" => {
            let addr = ObRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObReleaseObjectSecurity" => {
            let addr = ObReleaseObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObUnRegisterCallbacks" => {
            let addr = ObUnRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObfDereferenceObject" => {
            let addr = ObfDereferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObfDereferenceObjectWithTag" => {
            let addr = ObfDereferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObfReferenceObject" => {
            let addr = ObfReferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "ObfReferenceObjectWithTag" => {
            let addr = ObfReferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsAcquireSiloHardReference" => {
            let addr = PsAcquireSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsAllocSiloContextSlot" => {
            let addr = PsAllocSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsAllocateAffinityToken" => {
            let addr = PsAllocateAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsAssignImpersonationToken" => {
            let addr = PsAssignImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsAttachSiloToCurrentThread" => {
            let addr = PsAttachSiloToCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsChargePoolQuota" => {
            let addr = PsChargePoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsChargeProcessPoolQuota" => {
            let addr = PsChargeProcessPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsCreateSiloContext" => {
            let addr = PsCreateSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsCreateSystemThread" => {
            let addr = PsCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsDereferenceImpersonationToken" => {
            let addr = PsDereferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsDereferencePrimaryToken" => {
            let addr = PsDereferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsDereferenceSiloContext" => {
            let addr = PsDereferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsDetachSiloFromCurrentThread" => {
            let addr = PsDetachSiloFromCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsDisableImpersonation" => {
            let addr = PsDisableImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsFreeAffinityToken" => {
            let addr = PsFreeAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsFreeSiloContextSlot" => {
            let addr = PsFreeSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentProcessId" => {
            let addr = PsGetCurrentProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentServerSilo" => {
            let addr = PsGetCurrentServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentServerSiloName" => {
            let addr = PsGetCurrentServerSiloName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentSilo" => {
            let addr = PsGetCurrentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentThreadId" => {
            let addr = PsGetCurrentThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetCurrentThreadTeb" => {
            let addr = PsGetCurrentThreadTeb as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetEffectiveServerSilo" => {
            let addr = PsGetEffectiveServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetHostSilo" => {
            let addr = PsGetHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetJobServerSilo" => {
            let addr = PsGetJobServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetJobSilo" => {
            let addr = PsGetJobSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetParentSilo" => {
            let addr = PsGetParentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetPermanentSiloContext" => {
            let addr = PsGetPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetProcessCreateTimeQuadPart" => {
            let addr = PsGetProcessCreateTimeQuadPart as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetProcessExitStatus" => {
            let addr = PsGetProcessExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetProcessExitTime" => {
            let addr = PsGetProcessExitTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetProcessId" => {
            let addr = PsGetProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetProcessStartKey" => {
            let addr = PsGetProcessStartKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetServerSiloActiveConsoleId" => {
            let addr = PsGetServerSiloActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetServerSiloServiceSessionId" => {
            let addr = PsGetServerSiloServiceSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetSiloContainerId" => {
            let addr = PsGetSiloContainerId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetSiloContext" => {
            let addr = PsGetSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetSiloMonitorContextSlot" => {
            let addr = PsGetSiloMonitorContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadCreateTime" => {
            let addr = PsGetThreadCreateTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadExitStatus" => {
            let addr = PsGetThreadExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadId" => {
            let addr = PsGetThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadProcess" => {
            let addr = PsGetThreadProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadProcessId" => {
            let addr = PsGetThreadProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadProperty" => {
            let addr = PsGetThreadProperty as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetThreadServerSilo" => {
            let addr = PsGetThreadServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsGetVersion" => {
            let addr = PsGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsImpersonateClient" => {
            let addr = PsImpersonateClient as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsInsertPermanentSiloContext" => {
            let addr = PsInsertPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsInsertSiloContext" => {
            let addr = PsInsertSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsCurrentThreadInServerSilo" => {
            let addr = PsIsCurrentThreadInServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsCurrentThreadPrefetching" => {
            let addr = PsIsCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsDiskCountersEnabled" => {
            let addr = PsIsDiskCountersEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsHostSilo" => {
            let addr = PsIsHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsSystemThread" => {
            let addr = PsIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsThreadAttachedToSpecificSilo" => {
            let addr = PsIsThreadAttachedToSpecificSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsIsThreadTerminating" => {
            let addr = PsIsThreadTerminating as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsLookupProcessByProcessId" => {
            let addr = PsLookupProcessByProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsLookupThreadByThreadId" => {
            let addr = PsLookupThreadByThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsMakeSiloContextPermanent" => {
            let addr = PsMakeSiloContextPermanent as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsQueryProcessAvailableCpus" => {
            let addr = PsQueryProcessAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsQueryProcessAvailableCpusCount" => {
            let addr = PsQueryProcessAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsQuerySystemAvailableCpus" => {
            let addr = PsQuerySystemAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsQuerySystemAvailableCpusCount" => {
            let addr = PsQuerySystemAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsQueryTotalCycleTimeProcess" => {
            let addr = PsQueryTotalCycleTimeProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReferenceImpersonationToken" => {
            let addr = PsReferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReferencePrimaryToken" => {
            let addr = PsReferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReferenceSiloContext" => {
            let addr = PsReferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRegisterProcessAvailableCpusChangeNotification" => {
            let addr = PsRegisterProcessAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRegisterSiloMonitor" => {
            let addr = PsRegisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRegisterSystemAvailableCpusChangeNotification" => {
            let addr = PsRegisterSystemAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReleaseSiloHardReference" => {
            let addr = PsReleaseSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRemoveCreateThreadNotifyRoutine" => {
            let addr = PsRemoveCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRemoveLoadImageNotifyRoutine" => {
            let addr = PsRemoveLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRemoveSiloContext" => {
            let addr = PsRemoveSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReplaceSiloContext" => {
            let addr = PsReplaceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRestoreImpersonation" => {
            let addr = PsRestoreImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsReturnPoolQuota" => {
            let addr = PsReturnPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRevertToSelf" => {
            let addr = PsRevertToSelf as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsRevertToUserMultipleGroupAffinityThread" => {
            let addr = PsRevertToUserMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCreateProcessNotifyRoutine" => {
            let addr = PsSetCreateProcessNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCreateProcessNotifyRoutineEx" => {
            let addr = PsSetCreateProcessNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCreateProcessNotifyRoutineEx2" => {
            let addr = PsSetCreateProcessNotifyRoutineEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCreateThreadNotifyRoutine" => {
            let addr = PsSetCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCreateThreadNotifyRoutineEx" => {
            let addr = PsSetCreateThreadNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetCurrentThreadPrefetching" => {
            let addr = PsSetCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetLoadImageNotifyRoutine" => {
            let addr = PsSetLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetLoadImageNotifyRoutineEx" => {
            let addr = PsSetLoadImageNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsSetSystemMultipleGroupAffinityThread" => {
            let addr = PsSetSystemMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsStartSiloMonitor" => {
            let addr = PsStartSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsTerminateServerSilo" => {
            let addr = PsTerminateServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsTerminateSystemThread" => {
            let addr = PsTerminateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsUnregisterAvailableCpusChangeNotification" => {
            let addr = PsUnregisterAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsUnregisterSiloMonitor" => {
            let addr = PsUnregisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsUpdateDiskCounters" => {
            let addr = PsUpdateDiskCounters as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PsWrapApcWow64Thread" => {
            let addr = PsWrapApcWow64Thread as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedAllocateMemory" => {
            let addr = PshedAllocateMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedFreeMemory" => {
            let addr = PshedFreeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedIsSystemWheaEnabled" => {
            let addr = PshedIsSystemWheaEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedRegisterPlugin" => {
            let addr = PshedRegisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedSynchronizeExecution" => {
            let addr = PshedSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "PshedUnregisterPlugin" => {
            let addr = PshedUnregisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAbsoluteToSelfRelativeSD" => {
            let addr = RtlAbsoluteToSelfRelativeSD as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAddAccessAllowedAce" => {
            let addr = RtlAddAccessAllowedAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAddAccessAllowedAceEx" => {
            let addr = RtlAddAccessAllowedAceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAddAce" => {
            let addr = RtlAddAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAllocateAndInitializeSid" => {
            let addr = RtlAllocateAndInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAllocateAndInitializeSidEx" => {
            let addr = RtlAllocateAndInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAllocateHeap" => {
            let addr = RtlAllocateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAnsiStringToUnicodeString" => {
            let addr = RtlAnsiStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAppendStringToString" => {
            let addr = RtlAppendStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAppendUnicodeStringToString" => {
            let addr = RtlAppendUnicodeStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAppendUnicodeToString" => {
            let addr = RtlAppendUnicodeToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAreBitsClear" => {
            let addr = RtlAreBitsClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAreBitsSet" => {
            let addr = RtlAreBitsSet as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlAssert" => {
            let addr = RtlAssert as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCaptureContext" => {
            let addr = RtlCaptureContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCaptureContext2" => {
            let addr = RtlCaptureContext2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCaptureStackBackTrace" => {
            let addr = RtlCaptureStackBackTrace as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCharToInteger" => {
            let addr = RtlCharToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCheckRegistryKey" => {
            let addr = RtlCheckRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlClearAllBits" => {
            let addr = RtlClearAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlClearBit" => {
            let addr = RtlClearBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlClearBits" => {
            let addr = RtlClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCmDecodeMemIoResource" => {
            let addr = RtlCmDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCmEncodeMemIoResource" => {
            let addr = RtlCmEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareAltitudes" => {
            let addr = RtlCompareAltitudes as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareMemory" => {
            let addr = RtlCompareMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareMemoryUlong" => {
            let addr = RtlCompareMemoryUlong as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareString" => {
            let addr = RtlCompareString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareUnicodeString" => {
            let addr = RtlCompareUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompareUnicodeStrings" => {
            let addr = RtlCompareUnicodeStrings as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompressBuffer" => {
            let addr = RtlCompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCompressChunks" => {
            let addr = RtlCompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlContractHashTable" => {
            let addr = RtlContractHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlConvertSidToUnicodeString" => {
            let addr = RtlConvertSidToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyBitMap" => {
            let addr = RtlCopyBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyDeviceMemory" => {
            let addr = RtlCopyDeviceMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyLuid" => {
            let addr = RtlCopyLuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyMemoryNonTemporal" => {
            let addr = RtlCopyMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopySid" => {
            let addr = RtlCopySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyString" => {
            let addr = RtlCopyString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyUnicodeString" => {
            let addr = RtlCopyUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCopyVolatileMemory" => {
            let addr = RtlCopyVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCrc32" => {
            let addr = RtlCrc32 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCrc64" => {
            let addr = RtlCrc64 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateAcl" => {
            let addr = RtlCreateAcl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateHashTable" => {
            let addr = RtlCreateHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateHashTableEx" => {
            let addr = RtlCreateHashTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateHeap" => {
            let addr = RtlCreateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateRegistryKey" => {
            let addr = RtlCreateRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateSecurityDescriptor" => {
            let addr = RtlCreateSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateSecurityDescriptorRelative" => {
            let addr = RtlCreateSecurityDescriptorRelative as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateServiceSid" => {
            let addr = RtlCreateServiceSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateSystemVolumeInformationFolder" => {
            let addr = RtlCreateSystemVolumeInformationFolder as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateUnicodeString" => {
            let addr = RtlCreateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCreateVirtualAccountSid" => {
            let addr = RtlCreateVirtualAccountSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlCustomCPToUnicodeN" => {
            let addr = RtlCustomCPToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressBuffer" => {
            let addr = RtlDecompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressBufferEx" => {
            let addr = RtlDecompressBufferEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressBufferEx2" => {
            let addr = RtlDecompressBufferEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressChunks" => {
            let addr = RtlDecompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressFragment" => {
            let addr = RtlDecompressFragment as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDecompressFragmentEx" => {
            let addr = RtlDecompressFragmentEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDelete" => {
            let addr = RtlDelete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteAce" => {
            let addr = RtlDeleteAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteElementGenericTable" => {
            let addr = RtlDeleteElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteElementGenericTableAvl" => {
            let addr = RtlDeleteElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteElementGenericTableAvlEx" => {
            let addr = RtlDeleteElementGenericTableAvlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteHashTable" => {
            let addr = RtlDeleteHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteNoSplay" => {
            let addr = RtlDeleteNoSplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDeleteRegistryValue" => {
            let addr = RtlDeleteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDescribeChunk" => {
            let addr = RtlDescribeChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDestroyHeap" => {
            let addr = RtlDestroyHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDowncaseUnicodeChar" => {
            let addr = RtlDowncaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDowncaseUnicodeString" => {
            let addr = RtlDowncaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDrainNonVolatileFlush" => {
            let addr = RtlDrainNonVolatileFlush as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlDuplicateUnicodeString" => {
            let addr = RtlDuplicateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEndEnumerationHashTable" => {
            let addr = RtlEndEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEndStrongEnumerationHashTable" => {
            let addr = RtlEndStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEndWeakEnumerationHashTable" => {
            let addr = RtlEndWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateEntryHashTable" => {
            let addr = RtlEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateGenericTable" => {
            let addr = RtlEnumerateGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateGenericTableAvl" => {
            let addr = RtlEnumerateGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateGenericTableLikeADirectory" => {
            let addr = RtlEnumerateGenericTableLikeADirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateGenericTableWithoutSplaying" => {
            let addr = RtlEnumerateGenericTableWithoutSplaying as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEnumerateGenericTableWithoutSplayingAvl" => {
            let addr = RtlEnumerateGenericTableWithoutSplayingAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEqualPrefixSid" => {
            let addr = RtlEqualPrefixSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEqualSid" => {
            let addr = RtlEqualSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEqualString" => {
            let addr = RtlEqualString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlEqualUnicodeString" => {
            let addr = RtlEqualUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlExpandHashTable" => {
            let addr = RtlExpandHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlExtendCorrelationVector" => {
            let addr = RtlExtendCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlExtractBitMap" => {
            let addr = RtlExtractBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFillMemoryNonTemporal" => {
            let addr = RtlFillMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFillNonVolatileMemory" => {
            let addr = RtlFillNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindClearBits" => {
            let addr = RtlFindClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindClearBitsAndSet" => {
            let addr = RtlFindClearBitsAndSet as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindClearRuns" => {
            let addr = RtlFindClearRuns as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindClosestEncodableLength" => {
            let addr = RtlFindClosestEncodableLength as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindFirstRunClear" => {
            let addr = RtlFindFirstRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindLastBackwardRunClear" => {
            let addr = RtlFindLastBackwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindLeastSignificantBit" => {
            let addr = RtlFindLeastSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindLongestRunClear" => {
            let addr = RtlFindLongestRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindMostSignificantBit" => {
            let addr = RtlFindMostSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindNextForwardRunClear" => {
            let addr = RtlFindNextForwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindSetBits" => {
            let addr = RtlFindSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindSetBitsAndClear" => {
            let addr = RtlFindSetBitsAndClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFindUnicodePrefix" => {
            let addr = RtlFindUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFlushNonVolatileMemory" => {
            let addr = RtlFlushNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFlushNonVolatileMemoryRanges" => {
            let addr = RtlFlushNonVolatileMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeAnsiString" => {
            let addr = RtlFreeAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeHeap" => {
            let addr = RtlFreeHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeNonVolatileToken" => {
            let addr = RtlFreeNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeOemString" => {
            let addr = RtlFreeOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeSid" => {
            let addr = RtlFreeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeUTF8String" => {
            let addr = RtlFreeUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlFreeUnicodeString" => {
            let addr = RtlFreeUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGUIDFromString" => {
            let addr = RtlGUIDFromString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGenerate8dot3Name" => {
            let addr = RtlGenerate8dot3Name as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGenerateClass5Guid" => {
            let addr = RtlGenerateClass5Guid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetAce" => {
            let addr = RtlGetAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetAcesBufferSize" => {
            let addr = RtlGetAcesBufferSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetActiveConsoleId" => {
            let addr = RtlGetActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetCallersAddress" => {
            let addr = RtlGetCallersAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetCompressionWorkSpaceSize" => {
            let addr = RtlGetCompressionWorkSpaceSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetConsoleSessionForegroundProcessId" => {
            let addr = RtlGetConsoleSessionForegroundProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetDaclSecurityDescriptor" => {
            let addr = RtlGetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetElementGenericTable" => {
            let addr = RtlGetElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetElementGenericTableAvl" => {
            let addr = RtlGetElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetEnabledExtendedFeatures" => {
            let addr = RtlGetEnabledExtendedFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetGroupSecurityDescriptor" => {
            let addr = RtlGetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetNextEntryHashTable" => {
            let addr = RtlGetNextEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetNonVolatileToken" => {
            let addr = RtlGetNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetNtProductType" => {
            let addr = RtlGetNtProductType as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetNtSystemRoot" => {
            let addr = RtlGetNtSystemRoot as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetOwnerSecurityDescriptor" => {
            let addr = RtlGetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetPersistedStateLocation" => {
            let addr = RtlGetPersistedStateLocation as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetProductInfo" => {
            let addr = RtlGetProductInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetSaclSecurityDescriptor" => {
            let addr = RtlGetSaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetSuiteMask" => {
            let addr = RtlGetSuiteMask as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetSystemGlobalData" => {
            let addr = RtlGetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlGetVersion" => {
            let addr = RtlGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlHashUnicodeString" => {
            let addr = RtlHashUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIdentifierAuthoritySid" => {
            let addr = RtlIdentifierAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIdnToAscii" => {
            let addr = RtlIdnToAscii as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIdnToNameprepUnicode" => {
            let addr = RtlIdnToNameprepUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIdnToUnicode" => {
            let addr = RtlIdnToUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIncrementCorrelationVector" => {
            let addr = RtlIncrementCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitAnsiString" => {
            let addr = RtlInitAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitAnsiStringEx" => {
            let addr = RtlInitAnsiStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitCodePageTable" => {
            let addr = RtlInitCodePageTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitEnumerationHashTable" => {
            let addr = RtlInitEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitString" => {
            let addr = RtlInitString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitStringEx" => {
            let addr = RtlInitStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitStrongEnumerationHashTable" => {
            let addr = RtlInitStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitUTF8String" => {
            let addr = RtlInitUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitUTF8StringEx" => {
            let addr = RtlInitUTF8StringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitUnicodeString" => {
            let addr = RtlInitUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitUnicodeStringEx" => {
            let addr = RtlInitUnicodeStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitWeakEnumerationHashTable" => {
            let addr = RtlInitWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeBitMap" => {
            let addr = RtlInitializeBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeCorrelationVector" => {
            let addr = RtlInitializeCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeGenericTable" => {
            let addr = RtlInitializeGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeGenericTableAvl" => {
            let addr = RtlInitializeGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeSid" => {
            let addr = RtlInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeSidEx" => {
            let addr = RtlInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInitializeUnicodePrefix" => {
            let addr = RtlInitializeUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertElementGenericTable" => {
            let addr = RtlInsertElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertElementGenericTableAvl" => {
            let addr = RtlInsertElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertElementGenericTableFull" => {
            let addr = RtlInsertElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertElementGenericTableFullAvl" => {
            let addr = RtlInsertElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertEntryHashTable" => {
            let addr = RtlInsertEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInsertUnicodePrefix" => {
            let addr = RtlInsertUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlInt64ToUnicodeString" => {
            let addr = RtlInt64ToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIntegerToUnicodeString" => {
            let addr = RtlIntegerToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIoDecodeMemIoResource" => {
            let addr = RtlIoDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIoEncodeMemIoResource" => {
            let addr = RtlIoEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsApiSetImplemented" => {
            let addr = RtlIsApiSetImplemented as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsCloudFilesPlaceholder" => {
            let addr = RtlIsCloudFilesPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsGenericTableEmpty" => {
            let addr = RtlIsGenericTableEmpty as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsGenericTableEmptyAvl" => {
            let addr = RtlIsGenericTableEmptyAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsMultiSessionSku" => {
            let addr = RtlIsMultiSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsMultiUsersInSessionSku" => {
            let addr = RtlIsMultiUsersInSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsNameLegalDOS8Dot3" => {
            let addr = RtlIsNameLegalDOS8Dot3 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsNonEmptyDirectoryReparsePointAllowed" => {
            let addr = RtlIsNonEmptyDirectoryReparsePointAllowed as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsNormalizedString" => {
            let addr = RtlIsNormalizedString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsNtDdiVersionAvailable" => {
            let addr = RtlIsNtDdiVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsPartialPlaceholder" => {
            let addr = RtlIsPartialPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsPartialPlaceholderFileHandle" => {
            let addr = RtlIsPartialPlaceholderFileHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsPartialPlaceholderFileInfo" => {
            let addr = RtlIsPartialPlaceholderFileInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsSandboxedToken" => {
            let addr = RtlIsSandboxedToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsSandboxedTokenHandle" => {
            let addr = RtlIsSandboxedTokenHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsServicePackVersionInstalled" => {
            let addr = RtlIsServicePackVersionInstalled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsStateSeparationEnabled" => {
            let addr = RtlIsStateSeparationEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsUntrustedObject" => {
            let addr = RtlIsUntrustedObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsValidOemCharacter" => {
            let addr = RtlIsValidOemCharacter as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlIsZeroMemory" => {
            let addr = RtlIsZeroMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLengthRequiredSid" => {
            let addr = RtlLengthRequiredSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLengthSecurityDescriptor" => {
            let addr = RtlLengthSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLengthSid" => {
            let addr = RtlLengthSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupElementGenericTable" => {
            let addr = RtlLookupElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupElementGenericTableAvl" => {
            let addr = RtlLookupElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupElementGenericTableFull" => {
            let addr = RtlLookupElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupElementGenericTableFullAvl" => {
            let addr = RtlLookupElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupEntryHashTable" => {
            let addr = RtlLookupEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlLookupFirstMatchingElementGenericTableAvl" => {
            let addr = RtlLookupFirstMatchingElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlMapGenericMask" => {
            let addr = RtlMapGenericMask as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlMoveVolatileMemory" => {
            let addr = RtlMoveVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlMultiByteToUnicodeN" => {
            let addr = RtlMultiByteToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlMultiByteToUnicodeSize" => {
            let addr = RtlMultiByteToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNextUnicodePrefix" => {
            let addr = RtlNextUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNormalizeSecurityDescriptor" => {
            let addr = RtlNormalizeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNormalizeString" => {
            let addr = RtlNormalizeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNtStatusToDosError" => {
            let addr = RtlNtStatusToDosError as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNtStatusToDosErrorNoTeb" => {
            let addr = RtlNtStatusToDosErrorNoTeb as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberGenericTableElements" => {
            let addr = RtlNumberGenericTableElements as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberGenericTableElementsAvl" => {
            let addr = RtlNumberGenericTableElementsAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberOfClearBits" => {
            let addr = RtlNumberOfClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberOfClearBitsInRange" => {
            let addr = RtlNumberOfClearBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberOfSetBits" => {
            let addr = RtlNumberOfSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberOfSetBitsInRange" => {
            let addr = RtlNumberOfSetBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlNumberOfSetBitsUlongPtr" => {
            let addr = RtlNumberOfSetBitsUlongPtr as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlOemStringToCountedUnicodeString" => {
            let addr = RtlOemStringToCountedUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlOemStringToUnicodeString" => {
            let addr = RtlOemStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlOemToUnicodeN" => {
            let addr = RtlOemToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlOsDeploymentState" => {
            let addr = RtlOsDeploymentState as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlPrefetchMemoryNonTemporal" => {
            let addr = RtlPrefetchMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlPrefixString" => {
            let addr = RtlPrefixString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlPrefixUnicodeString" => {
            let addr = RtlPrefixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryPackageIdentity" => {
            let addr = RtlQueryPackageIdentity as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryPackageIdentityEx" => {
            let addr = RtlQueryPackageIdentityEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryProcessPlaceholderCompatibilityMode" => {
            let addr = RtlQueryProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryRegistryValueWithFallback" => {
            let addr = RtlQueryRegistryValueWithFallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryRegistryValues" => {
            let addr = RtlQueryRegistryValues as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryThreadPlaceholderCompatibilityMode" => {
            let addr = RtlQueryThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlQueryValidationRunlevel" => {
            let addr = RtlQueryValidationRunlevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRaiseCustomSystemEventTrigger" => {
            let addr = RtlRaiseCustomSystemEventTrigger as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRandom" => {
            let addr = RtlRandom as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRandomEx" => {
            let addr = RtlRandomEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRealPredecessor" => {
            let addr = RtlRealPredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRealSuccessor" => {
            let addr = RtlRealSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRemoveEntryHashTable" => {
            let addr = RtlRemoveEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRemoveUnicodePrefix" => {
            let addr = RtlRemoveUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlReplaceSidInSd" => {
            let addr = RtlReplaceSidInSd as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlReserveChunk" => {
            let addr = RtlReserveChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRunOnceBeginInitialize" => {
            let addr = RtlRunOnceBeginInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRunOnceComplete" => {
            let addr = RtlRunOnceComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRunOnceExecuteOnce" => {
            let addr = RtlRunOnceExecuteOnce as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlRunOnceInitialize" => {
            let addr = RtlRunOnceInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSecondsSince1970ToTime" => {
            let addr = RtlSecondsSince1970ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSecondsSince1980ToTime" => {
            let addr = RtlSecondsSince1980ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSelfRelativeToAbsoluteSD" => {
            let addr = RtlSelfRelativeToAbsoluteSD as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetAllBits" => {
            let addr = RtlSetAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetBit" => {
            let addr = RtlSetBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetBits" => {
            let addr = RtlSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetDaclSecurityDescriptor" => {
            let addr = RtlSetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetGroupSecurityDescriptor" => {
            let addr = RtlSetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetOwnerSecurityDescriptor" => {
            let addr = RtlSetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetProcessPlaceholderCompatibilityMode" => {
            let addr = RtlSetProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetSystemGlobalData" => {
            let addr = RtlSetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetThreadPlaceholderCompatibilityMode" => {
            let addr = RtlSetThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSetVolatileMemory" => {
            let addr = RtlSetVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSplay" => {
            let addr = RtlSplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlStringFromGUID" => {
            let addr = RtlStringFromGUID as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlStronglyEnumerateEntryHashTable" => {
            let addr = RtlStronglyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSubAuthorityCountSid" => {
            let addr = RtlSubAuthorityCountSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSubAuthoritySid" => {
            let addr = RtlSubAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSubtreePredecessor" => {
            let addr = RtlSubtreePredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSubtreeSuccessor" => {
            let addr = RtlSubtreeSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlSuffixUnicodeString" => {
            let addr = RtlSuffixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlTestBit" => {
            let addr = RtlTestBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlTimeFieldsToTime" => {
            let addr = RtlTimeFieldsToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlTimeToSecondsSince1970" => {
            let addr = RtlTimeToSecondsSince1970 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlTimeToSecondsSince1980" => {
            let addr = RtlTimeToSecondsSince1980 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlTimeToTimeFields" => {
            let addr = RtlTimeToTimeFields as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUTF8StringToUnicodeString" => {
            let addr = RtlUTF8StringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUTF8ToUnicodeN" => {
            let addr = RtlUTF8ToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToAnsiString" => {
            let addr = RtlUnicodeStringToAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToCountedOemString" => {
            let addr = RtlUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToInt64" => {
            let addr = RtlUnicodeStringToInt64 as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToInteger" => {
            let addr = RtlUnicodeStringToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToOemString" => {
            let addr = RtlUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeStringToUTF8String" => {
            let addr = RtlUnicodeStringToUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeToCustomCPN" => {
            let addr = RtlUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeToMultiByteN" => {
            let addr = RtlUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeToMultiByteSize" => {
            let addr = RtlUnicodeToMultiByteSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeToOemN" => {
            let addr = RtlUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUnicodeToUTF8N" => {
            let addr = RtlUnicodeToUTF8N as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeChar" => {
            let addr = RtlUpcaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeString" => {
            let addr = RtlUpcaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeStringToCountedOemString" => {
            let addr = RtlUpcaseUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeStringToOemString" => {
            let addr = RtlUpcaseUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeToCustomCPN" => {
            let addr = RtlUpcaseUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeToMultiByteN" => {
            let addr = RtlUpcaseUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpcaseUnicodeToOemN" => {
            let addr = RtlUpcaseUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpperChar" => {
            let addr = RtlUpperChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlUpperString" => {
            let addr = RtlUpperString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlValidRelativeSecurityDescriptor" => {
            let addr = RtlValidRelativeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlValidSecurityDescriptor" => {
            let addr = RtlValidSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlValidSid" => {
            let addr = RtlValidSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlValidateCorrelationVector" => {
            let addr = RtlValidateCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlValidateUnicodeString" => {
            let addr = RtlValidateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlVerifyVersionInfo" => {
            let addr = RtlVerifyVersionInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlVolumeDeviceToDosName" => {
            let addr = RtlVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlWalkFrameChain" => {
            let addr = RtlWalkFrameChain as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlWeaklyEnumerateEntryHashTable" => {
            let addr = RtlWeaklyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlWriteNonVolatileMemory" => {
            let addr = RtlWriteNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlWriteRegistryValue" => {
            let addr = RtlWriteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlxAnsiStringToUnicodeSize" => {
            let addr = RtlxAnsiStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlxOemStringToUnicodeSize" => {
            let addr = RtlxOemStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlxUnicodeStringToAnsiSize" => {
            let addr = RtlxUnicodeStringToAnsiSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        "RtlxUnicodeStringToOemSize" => {
            let addr = RtlxUnicodeStringToOemSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }
        }
        _ => Err(HookError::InvalidAddress),
    }
}

pub fn remove_ept_hook_by_name(name: &str) -> Result<(), HookError> {
    match name {
        "DbgBreakPointWithStatus" => {
            let addr = DbgBreakPointWithStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgPrint" => {
            let addr = DbgPrint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgPrintEx" => {
            let addr = DbgPrintEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgPrintReturnControlC" => {
            let addr = DbgPrintReturnControlC as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgPrompt" => {
            let addr = DbgPrompt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgQueryDebugFilterState" => {
            let addr = DbgQueryDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgSetDebugFilterState" => {
            let addr = DbgSetDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "DbgSetDebugPrintCallback" => {
            let addr = DbgSetDebugPrintCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireFastMutex" => {
            let addr = ExAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireFastMutexUnsafe" => {
            let addr = ExAcquireFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquirePushLockExclusiveEx" => {
            let addr = ExAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquirePushLockSharedEx" => {
            let addr = ExAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireResourceExclusiveLite" => {
            let addr = ExAcquireResourceExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireResourceSharedLite" => {
            let addr = ExAcquireResourceSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireRundownProtection" => {
            let addr = ExAcquireRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireRundownProtectionCacheAware" => {
            let addr = ExAcquireRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireRundownProtectionCacheAwareEx" => {
            let addr = ExAcquireRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireRundownProtectionEx" => {
            let addr = ExAcquireRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSharedStarveExclusive" => {
            let addr = ExAcquireSharedStarveExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSharedWaitForExclusive" => {
            let addr = ExAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSpinLockExclusive" => {
            let addr = ExAcquireSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSpinLockShared" => {
            let addr = ExAcquireSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAdjustLookasideDepth" => {
            let addr = ExAdjustLookasideDepth as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocateCacheAwareRundownProtection" => {
            let addr = ExAllocateCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocateFromLookasideListEx" => {
            let addr = ExAllocateFromLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocateFromNPagedLookasideList" => {
            let addr = ExAllocateFromNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocateFromPagedLookasideList" => {
            let addr = ExAllocateFromPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocatePool2" => {
            let addr = ExAllocatePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocatePool3" => {
            let addr = ExAllocatePool3 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocatePoolWithQuota" => {
            let addr = ExAllocatePoolWithQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExAllocateTimer" => {
            let addr = ExAllocateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExCancelTimer" => {
            let addr = ExCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExCleanupRundownProtectionCacheAware" => {
            let addr = ExCleanupRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExConvertExclusiveToSharedLite" => {
            let addr = ExConvertExclusiveToSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExCreateCallback" => {
            let addr = ExCreateCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExCreatePool" => {
            let addr = ExCreatePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDeleteLookasideListEx" => {
            let addr = ExDeleteLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDeleteNPagedLookasideList" => {
            let addr = ExDeleteNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDeletePagedLookasideList" => {
            let addr = ExDeletePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDeleteResourceLite" => {
            let addr = ExDeleteResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDeleteTimer" => {
            let addr = ExDeleteTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDestroyPool" => {
            let addr = ExDestroyPool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExDisableResourceBoostLite" => {
            let addr = ExDisableResourceBoostLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireResourceExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireResourceShared" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExEnumerateSystemFirmwareTables" => {
            let addr = ExEnumerateSystemFirmwareTables as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExExtendZone" => {
            let addr = ExExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFlushLookasideListEx" => {
            let addr = ExFlushLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreeCacheAwareRundownProtection" => {
            let addr = ExFreeCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreePool" => {
            let addr = ExFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreePool2" => {
            let addr = ExFreePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreePoolWithTag" => {
            let addr = ExFreePoolWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreeToLookasideListEx" => {
            let addr = ExFreeToLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreeToNPagedLookasideList" => {
            let addr = ExFreeToNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExFreeToPagedLookasideList" => {
            let addr = ExFreeToPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetExclusiveWaiterCount" => {
            let addr = ExGetExclusiveWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetFirmwareEnvironmentVariable" => {
            let addr = ExGetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetFirmwareType" => {
            let addr = ExGetFirmwareType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetPreviousMode" => {
            let addr = ExGetPreviousMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetSharedWaiterCount" => {
            let addr = ExGetSharedWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExGetSystemFirmwareTable" => {
            let addr = ExGetSystemFirmwareTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeDeviceAts" => {
            let addr = ExInitializeDeviceAts as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeLookasideListEx" => {
            let addr = ExInitializeLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeNPagedLookasideList" => {
            let addr = ExInitializeNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializePagedLookasideList" => {
            let addr = ExInitializePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializePushLock" => {
            let addr = ExInitializePushLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeResourceLite" => {
            let addr = ExInitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeRundownProtection" => {
            let addr = ExInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeRundownProtectionCacheAware" => {
            let addr = ExInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeRundownProtectionCacheAwareEx" => {
            let addr = ExInitializeRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInitializeZone" => {
            let addr = ExInitializeZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedAddLargeInteger" => {
            let addr = ExInterlockedAddLargeInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedAddUlong" => {
            let addr = ExInterlockedAddUlong as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedExtendZone" => {
            let addr = ExInterlockedExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedInsertHeadList" => {
            let addr = ExInterlockedInsertHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedInsertTailList" => {
            let addr = ExInterlockedInsertTailList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedPopEntryList" => {
            let addr = ExInterlockedPopEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedPushEntryList" => {
            let addr = ExInterlockedPushEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExInterlockedRemoveHeadList" => {
            let addr = ExInterlockedRemoveHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExIsManufacturingModeEnabled" => {
            let addr = ExIsManufacturingModeEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExIsProcessorFeaturePresent" => {
            let addr = ExIsProcessorFeaturePresent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExIsResourceAcquiredExclusiveLite" => {
            let addr = ExIsResourceAcquiredExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExIsResourceAcquiredSharedLite" => {
            let addr = ExIsResourceAcquiredSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExIsSoftBoot" => {
            let addr = ExIsSoftBoot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExLocalTimeToSystemTime" => {
            let addr = ExLocalTimeToSystemTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExNotifyCallback" => {
            let addr = ExNotifyCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExQueryDepthSList" => {
            let addr = ExQueryDepthSList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExQueryPoolBlockSize" => {
            let addr = ExQueryPoolBlockSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExQueryTimerResolution" => {
            let addr = ExQueryTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExQueueWorkItem" => {
            let addr = ExQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRaiseAccessViolation" => {
            let addr = ExRaiseAccessViolation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRaiseDatatypeMisalignment" => {
            let addr = ExRaiseDatatypeMisalignment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRaiseStatus" => {
            let addr = ExRaiseStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRcuFreePool" => {
            let addr = ExRcuFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReInitializeRundownProtection" => {
            let addr = ExReInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReInitializeRundownProtectionCacheAware" => {
            let addr = ExReInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRegisterCallback" => {
            let addr = ExRegisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReinitializeResourceLite" => {
            let addr = ExReinitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseFastMutex" => {
            let addr = ExReleaseFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseFastMutexUnsafe" => {
            let addr = ExReleaseFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleasePushLockExclusiveEx" => {
            let addr = ExReleasePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleasePushLockSharedEx" => {
            let addr = ExReleasePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseResourceAndLeaveCriticalRegion" => {
            let addr = ExReleaseResourceAndLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseResourceForThreadLite" => {
            let addr = ExReleaseResourceForThreadLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseResourceLite" => {
            let addr = ExReleaseResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseRundownProtection" => {
            let addr = ExReleaseRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseRundownProtectionCacheAware" => {
            let addr = ExReleaseRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseRundownProtectionCacheAwareEx" => {
            let addr = ExReleaseRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseRundownProtectionEx" => {
            let addr = ExReleaseRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseSpinLockExclusive" => {
            let addr = ExReleaseSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseSpinLockExclusiveFromDpcLevel" => {
            let addr = ExReleaseSpinLockExclusiveFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseSpinLockShared" => {
            let addr = ExReleaseSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExReleaseSpinLockSharedFromDpcLevel" => {
            let addr = ExReleaseSpinLockSharedFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRundownCompleted" => {
            let addr = ExRundownCompleted as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExRundownCompletedCacheAware" => {
            let addr = ExRundownCompletedCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSecurePoolUpdate" => {
            let addr = ExSecurePoolUpdate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSecurePoolValidate" => {
            let addr = ExSecurePoolValidate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSetFirmwareEnvironmentVariable" => {
            let addr = ExSetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSetResourceOwnerPointer" => {
            let addr = ExSetResourceOwnerPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSetResourceOwnerPointerEx" => {
            let addr = ExSetResourceOwnerPointerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSetTimer" => {
            let addr = ExSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSetTimerResolution" => {
            let addr = ExSetTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSizeOfRundownProtectionCacheAware" => {
            let addr = ExSizeOfRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExSystemTimeToLocalTime" => {
            let addr = ExSystemTimeToLocalTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryAcquirePushLockExclusiveEx" => {
            let addr = ExTryAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryAcquirePushLockSharedEx" => {
            let addr = ExTryAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryConvertSharedSpinLockExclusive" => {
            let addr = ExTryConvertSharedSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExTryToAcquireFastMutex" => {
            let addr = ExTryToAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExUnregisterCallback" => {
            let addr = ExUnregisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExUuidCreate" => {
            let addr = ExUuidCreate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExVerifySuite" => {
            let addr = ExVerifySuite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExWaitForRundownProtectionRelease" => {
            let addr = ExWaitForRundownProtectionRelease as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExWaitForRundownProtectionReleaseCacheAware" => {
            let addr = ExWaitForRundownProtectionReleaseCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExpInterlockedFlushSList" => {
            let addr = ExpInterlockedFlushSList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExpInterlockedPopEntrySList" => {
            let addr = ExpInterlockedPopEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExpInterlockedPushEntrySList" => {
            let addr = ExpInterlockedPushEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ExportSecurityContext" => {
            let addr = ExportSecurityContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAcquireCancelSpinLock" => {
            let addr = IoAcquireCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAcquireKsrPersistentMemory" => {
            let addr = IoAcquireKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAcquireKsrPersistentMemoryEx" => {
            let addr = IoAcquireKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAcquireRemoveLockEx" => {
            let addr = IoAcquireRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAcquireVpbSpinLock" => {
            let addr = IoAcquireVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateController" => {
            let addr = IoAllocateController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateDriverObjectExtension" => {
            let addr = IoAllocateDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateErrorLogEntry" => {
            let addr = IoAllocateErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateIrp" => {
            let addr = IoAllocateIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateIrpEx" => {
            let addr = IoAllocateIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateMdl" => {
            let addr = IoAllocateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateMiniCompletionPacket" => {
            let addr = IoAllocateMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateSfioStreamIdentifier" => {
            let addr = IoAllocateSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAllocateWorkItem" => {
            let addr = IoAllocateWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoApplyPriorityInfoThread" => {
            let addr = IoApplyPriorityInfoThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAssignResources" => {
            let addr = IoAssignResources as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAttachDevice" => {
            let addr = IoAttachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAttachDeviceByPointer" => {
            let addr = IoAttachDeviceByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAttachDeviceToDeviceStack" => {
            let addr = IoAttachDeviceToDeviceStack as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoAttachDeviceToDeviceStackSafe" => {
            let addr = IoAttachDeviceToDeviceStackSafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoBuildAsynchronousFsdRequest" => {
            let addr = IoBuildAsynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoBuildDeviceIoControlRequest" => {
            let addr = IoBuildDeviceIoControlRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoBuildPartialMdl" => {
            let addr = IoBuildPartialMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoBuildSynchronousFsdRequest" => {
            let addr = IoBuildSynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCancelFileOpen" => {
            let addr = IoCancelFileOpen as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCancelIrp" => {
            let addr = IoCancelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckDesiredAccess" => {
            let addr = IoCheckDesiredAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckEaBufferValidity" => {
            let addr = IoCheckEaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckFileObjectOpenedAsCopyDestination" => {
            let addr = IoCheckFileObjectOpenedAsCopyDestination as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckFileObjectOpenedAsCopySource" => {
            let addr = IoCheckFileObjectOpenedAsCopySource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckFunctionAccess" => {
            let addr = IoCheckFunctionAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckLinkShareAccess" => {
            let addr = IoCheckLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckQuerySetFileInformation" => {
            let addr = IoCheckQuerySetFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckQuerySetVolumeInformation" => {
            let addr = IoCheckQuerySetVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckQuotaBufferValidity" => {
            let addr = IoCheckQuotaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckShareAccess" => {
            let addr = IoCheckShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCheckShareAccessEx" => {
            let addr = IoCheckShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCleanupIrp" => {
            let addr = IoCleanupIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoClearActivityIdThread" => {
            let addr = IoClearActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoClearFsTrackOffsetState" => {
            let addr = IoClearFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoClearIrpExtraCreateParameter" => {
            let addr = IoClearIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoConnectInterrupt" => {
            let addr = IoConnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoConnectInterruptEx" => {
            let addr = IoConnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateController" => {
            let addr = IoCreateController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateDevice" => {
            let addr = IoCreateDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateDisk" => {
            let addr = IoCreateDisk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateDriverProxyExtension" => {
            let addr = IoCreateDriverProxyExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateFile" => {
            let addr = IoCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateFileEx" => {
            let addr = IoCreateFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateFileSpecifyDeviceObjectHint" => {
            let addr = IoCreateFileSpecifyDeviceObjectHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateNotificationEvent" => {
            let addr = IoCreateNotificationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateStreamFileObject" => {
            let addr = IoCreateStreamFileObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateStreamFileObjectEx" => {
            let addr = IoCreateStreamFileObjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateStreamFileObjectEx2" => {
            let addr = IoCreateStreamFileObjectEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateStreamFileObjectLite" => {
            let addr = IoCreateStreamFileObjectLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateSymbolicLink" => {
            let addr = IoCreateSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateSynchronizationEvent" => {
            let addr = IoCreateSynchronizationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateSystemThread" => {
            let addr = IoCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCreateUnprotectedSymbolicLink" => {
            let addr = IoCreateUnprotectedSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqInitialize" => {
            let addr = IoCsqInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqInitializeEx" => {
            let addr = IoCsqInitializeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqInsertIrp" => {
            let addr = IoCsqInsertIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqInsertIrpEx" => {
            let addr = IoCsqInsertIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqRemoveIrp" => {
            let addr = IoCsqRemoveIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoCsqRemoveNextIrp" => {
            let addr = IoCsqRemoveNextIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDecrementKeepAliveCount" => {
            let addr = IoDecrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDeleteController" => {
            let addr = IoDeleteController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDeleteDevice" => {
            let addr = IoDeleteDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDeleteSymbolicLink" => {
            let addr = IoDeleteSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDetachDevice" => {
            let addr = IoDetachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDisconnectInterrupt" => {
            let addr = IoDisconnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDisconnectInterruptEx" => {
            let addr = IoDisconnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoDriverProxyCreateHotSwappableWorkerThread" => {
            let addr = IoDriverProxyCreateHotSwappableWorkerThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoEnumerateDeviceObjectList" => {
            let addr = IoEnumerateDeviceObjectList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoEnumerateKsrPersistentMemoryEx" => {
            let addr = IoEnumerateKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoEnumerateRegisteredFiltersList" => {
            let addr = IoEnumerateRegisteredFiltersList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFastQueryNetworkAttributes" => {
            let addr = IoFastQueryNetworkAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoForwardIrpSynchronously" => {
            let addr = IoForwardIrpSynchronously as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeController" => {
            let addr = IoFreeController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeErrorLogEntry" => {
            let addr = IoFreeErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeIrp" => {
            let addr = IoFreeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeKsrPersistentMemory" => {
            let addr = IoFreeKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeMdl" => {
            let addr = IoFreeMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeMiniCompletionPacket" => {
            let addr = IoFreeMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeSfioStreamIdentifier" => {
            let addr = IoFreeSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoFreeWorkItem" => {
            let addr = IoFreeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetActivityIdIrp" => {
            let addr = IoGetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetActivityIdThread" => {
            let addr = IoGetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetAffinityInterrupt" => {
            let addr = IoGetAffinityInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetAttachedDevice" => {
            let addr = IoGetAttachedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetAttachedDeviceReference" => {
            let addr = IoGetAttachedDeviceReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetBaseFileSystemDeviceObject" => {
            let addr = IoGetBaseFileSystemDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetBootDiskInformation" => {
            let addr = IoGetBootDiskInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetBootDiskInformationLite" => {
            let addr = IoGetBootDiskInformationLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetConfigurationInformation" => {
            let addr = IoGetConfigurationInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetContainerInformation" => {
            let addr = IoGetContainerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetCopyInformationExtension" => {
            let addr = IoGetCopyInformationExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetCurrentProcess" => {
            let addr = IoGetCurrentProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceAttachmentBaseRef" => {
            let addr = IoGetDeviceAttachmentBaseRef as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceDirectory" => {
            let addr = IoGetDeviceDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceInterfaceAlias" => {
            let addr = IoGetDeviceInterfaceAlias as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceInterfacePropertyData" => {
            let addr = IoGetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceInterfaces" => {
            let addr = IoGetDeviceInterfaces as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceNumaNode" => {
            let addr = IoGetDeviceNumaNode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceObjectPointer" => {
            let addr = IoGetDeviceObjectPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceProperty" => {
            let addr = IoGetDeviceProperty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDevicePropertyData" => {
            let addr = IoGetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDeviceToVerify" => {
            let addr = IoGetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDiskDeviceObject" => {
            let addr = IoGetDiskDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDmaAdapter" => {
            let addr = IoGetDmaAdapter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverDirectory" => {
            let addr = IoGetDriverDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverObjectExtension" => {
            let addr = IoGetDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverProxyEndpointWrapper" => {
            let addr = IoGetDriverProxyEndpointWrapper as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverProxyExtensionFromDriverObject" => {
            let addr = IoGetDriverProxyExtensionFromDriverObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverProxyExtensionVersion" => {
            let addr = IoGetDriverProxyExtensionVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetDriverProxyFeatures" => {
            let addr = IoGetDriverProxyFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetFileObjectGenericMapping" => {
            let addr = IoGetFileObjectGenericMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetFsTrackOffsetState" => {
            let addr = IoGetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetFsZeroingOffset" => {
            let addr = IoGetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetInitialStack" => {
            let addr = IoGetInitialStack as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetInitiatorProcess" => {
            let addr = IoGetInitiatorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetIoAttributionHandle" => {
            let addr = IoGetIoAttributionHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetIoPriorityHint" => {
            let addr = IoGetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetIommuInterface" => {
            let addr = IoGetIommuInterface as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetIommuInterfaceEx" => {
            let addr = IoGetIommuInterfaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetIrpExtraCreateParameter" => {
            let addr = IoGetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetKsrPersistentMemoryBuffer" => {
            let addr = IoGetKsrPersistentMemoryBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetLowerDeviceObject" => {
            let addr = IoGetLowerDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetOplockKeyContext" => {
            let addr = IoGetOplockKeyContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetOplockKeyContextEx" => {
            let addr = IoGetOplockKeyContextEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetPagingIoPriority" => {
            let addr = IoGetPagingIoPriority as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetRelatedDeviceObject" => {
            let addr = IoGetRelatedDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetRequestorProcess" => {
            let addr = IoGetRequestorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetRequestorProcessId" => {
            let addr = IoGetRequestorProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetRequestorSessionId" => {
            let addr = IoGetRequestorSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetSfioStreamIdentifier" => {
            let addr = IoGetSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetShadowFileInformation" => {
            let addr = IoGetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetSilo" => {
            let addr = IoGetSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetSiloParameters" => {
            let addr = IoGetSiloParameters as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetStackLimits" => {
            let addr = IoGetStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetTopLevelIrp" => {
            let addr = IoGetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoGetTransactionParameterBlock" => {
            let addr = IoGetTransactionParameterBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoHotSwapDriverProxyEndpoints" => {
            let addr = IoHotSwapDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIncrementKeepAliveCount" => {
            let addr = IoIncrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInitializeIrp" => {
            let addr = IoInitializeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInitializeIrpEx" => {
            let addr = IoInitializeIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInitializeRemoveLockEx" => {
            let addr = IoInitializeRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInitializeTimer" => {
            let addr = IoInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInitializeWorkItem" => {
            let addr = IoInitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInvalidateDeviceRelations" => {
            let addr = IoInvalidateDeviceRelations as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoInvalidateDeviceState" => {
            let addr = IoInvalidateDeviceState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIrpHasFsTrackOffsetExtensionType" => {
            let addr = IoIrpHasFsTrackOffsetExtensionType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIs32bitProcess" => {
            let addr = IoIs32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsFileObjectIgnoringSharing" => {
            let addr = IoIsFileObjectIgnoringSharing as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsFileOriginRemote" => {
            let addr = IoIsFileOriginRemote as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsInitiator32bitProcess" => {
            let addr = IoIsInitiator32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsOperationSynchronous" => {
            let addr = IoIsOperationSynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsSystemThread" => {
            let addr = IoIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsValidIrpStatus" => {
            let addr = IoIsValidIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsValidNameGraftingBuffer" => {
            let addr = IoIsValidNameGraftingBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoIsWdmVersionAvailable" => {
            let addr = IoIsWdmVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoMakeAssociatedIrp" => {
            let addr = IoMakeAssociatedIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoMakeAssociatedIrpEx" => {
            let addr = IoMakeAssociatedIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoMapKsrPersistentMemoryEx" => {
            let addr = IoMapKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoOpenDeviceInterfaceRegistryKey" => {
            let addr = IoOpenDeviceInterfaceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoOpenDeviceRegistryKey" => {
            let addr = IoOpenDeviceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoOpenDriverRegistryKey" => {
            let addr = IoOpenDriverRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoPageRead" => {
            let addr = IoPageRead as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoPropagateActivityIdToThread" => {
            let addr = IoPropagateActivityIdToThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryDeviceDescription" => {
            let addr = IoQueryDeviceDescription as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryDmaFeatureSupport" => {
            let addr = IoQueryDmaFeatureSupport as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryFileDosDeviceName" => {
            let addr = IoQueryFileDosDeviceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryFileInformation" => {
            let addr = IoQueryFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryFullDriverPath" => {
            let addr = IoQueryFullDriverPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryInformationByName" => {
            let addr = IoQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryKsrPersistentMemorySize" => {
            let addr = IoQueryKsrPersistentMemorySize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryKsrPersistentMemorySizeEx" => {
            let addr = IoQueryKsrPersistentMemorySizeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueryVolumeInformation" => {
            let addr = IoQueryVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueueThreadIrp" => {
            let addr = IoQueueThreadIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueueWorkItem" => {
            let addr = IoQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoQueueWorkItemEx" => {
            let addr = IoQueueWorkItemEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRaiseHardError" => {
            let addr = IoRaiseHardError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRaiseInformationalHardError" => {
            let addr = IoRaiseInformationalHardError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReadDiskSignature" => {
            let addr = IoReadDiskSignature as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReadPartitionTable" => {
            let addr = IoReadPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReadPartitionTableEx" => {
            let addr = IoReadPartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRecordIoAttribution" => {
            let addr = IoRecordIoAttribution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterBootDriverCallback" => {
            let addr = IoRegisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterBootDriverReinitialization" => {
            let addr = IoRegisterBootDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterContainerNotification" => {
            let addr = IoRegisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterDeviceInterface" => {
            let addr = IoRegisterDeviceInterface as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterDriverProxyEndpoints" => {
            let addr = IoRegisterDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterDriverReinitialization" => {
            let addr = IoRegisterDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterFileSystem" => {
            let addr = IoRegisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterFsRegistrationChange" => {
            let addr = IoRegisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterFsRegistrationChangeMountAware" => {
            let addr = IoRegisterFsRegistrationChangeMountAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterLastChanceShutdownNotification" => {
            let addr = IoRegisterLastChanceShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterPlugPlayNotification" => {
            let addr = IoRegisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRegisterShutdownNotification" => {
            let addr = IoRegisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReleaseCancelSpinLock" => {
            let addr = IoReleaseCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReleaseRemoveLockAndWaitEx" => {
            let addr = IoReleaseRemoveLockAndWaitEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReleaseRemoveLockEx" => {
            let addr = IoReleaseRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReleaseVpbSpinLock" => {
            let addr = IoReleaseVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRemoveIoCompletion" => {
            let addr = IoRemoveIoCompletion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRemoveLinkShareAccess" => {
            let addr = IoRemoveLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRemoveLinkShareAccessEx" => {
            let addr = IoRemoveLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRemoveShareAccess" => {
            let addr = IoRemoveShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReplaceFileObjectName" => {
            let addr = IoReplaceFileObjectName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReplacePartitionUnit" => {
            let addr = IoReplacePartitionUnit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportDetectedDevice" => {
            let addr = IoReportDetectedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportInterruptActive" => {
            let addr = IoReportInterruptActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportInterruptInactive" => {
            let addr = IoReportInterruptInactive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportResourceForDetection" => {
            let addr = IoReportResourceForDetection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportResourceUsage" => {
            let addr = IoReportResourceUsage as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportRootDevice" => {
            let addr = IoReportRootDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportTargetDeviceChange" => {
            let addr = IoReportTargetDeviceChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReportTargetDeviceChangeAsynchronous" => {
            let addr = IoReportTargetDeviceChangeAsynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRequestDeviceEject" => {
            let addr = IoRequestDeviceEject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRequestDeviceEjectEx" => {
            let addr = IoRequestDeviceEjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRequestDeviceRemovalForReset" => {
            let addr = IoRequestDeviceRemovalForReset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReserveKsrPersistentMemory" => {
            let addr = IoReserveKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReserveKsrPersistentMemoryEx" => {
            let addr = IoReserveKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoRetrievePriorityInfo" => {
            let addr = IoRetrievePriorityInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoReuseIrp" => {
            let addr = IoReuseIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetActivityIdIrp" => {
            let addr = IoSetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetActivityIdThread" => {
            let addr = IoSetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetCompletionRoutineEx" => {
            let addr = IoSetCompletionRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetDeviceInterfacePropertyData" => {
            let addr = IoSetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetDeviceInterfaceState" => {
            let addr = IoSetDeviceInterfaceState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetDevicePropertyData" => {
            let addr = IoSetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetDeviceToVerify" => {
            let addr = IoSetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetFileObjectIgnoreSharing" => {
            let addr = IoSetFileObjectIgnoreSharing as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetFileOrigin" => {
            let addr = IoSetFileOrigin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetFsTrackOffsetState" => {
            let addr = IoSetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetFsZeroingOffset" => {
            let addr = IoSetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetFsZeroingOffsetRequired" => {
            let addr = IoSetFsZeroingOffsetRequired as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetHardErrorOrVerifyDevice" => {
            let addr = IoSetHardErrorOrVerifyDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetInformation" => {
            let addr = IoSetInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetIoAttributionIrp" => {
            let addr = IoSetIoAttributionIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetIoCompletionEx" => {
            let addr = IoSetIoCompletionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetIoPriorityHint" => {
            let addr = IoSetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetIrpExtraCreateParameter" => {
            let addr = IoSetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetLinkShareAccess" => {
            let addr = IoSetLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetMasterIrpStatus" => {
            let addr = IoSetMasterIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetPartitionInformation" => {
            let addr = IoSetPartitionInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetPartitionInformationEx" => {
            let addr = IoSetPartitionInformationEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetShadowFileInformation" => {
            let addr = IoSetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetShareAccess" => {
            let addr = IoSetShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetShareAccessEx" => {
            let addr = IoSetShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetStartIoAttributes" => {
            let addr = IoSetStartIoAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetSystemPartition" => {
            let addr = IoSetSystemPartition as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetThreadHardErrorMode" => {
            let addr = IoSetThreadHardErrorMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSetTopLevelIrp" => {
            let addr = IoSetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSizeOfIrpEx" => {
            let addr = IoSizeOfIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSizeofWorkItem" => {
            let addr = IoSizeofWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoStartNextPacket" => {
            let addr = IoStartNextPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoStartNextPacketByKey" => {
            let addr = IoStartNextPacketByKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoStartPacket" => {
            let addr = IoStartPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoStartTimer" => {
            let addr = IoStartTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoStopTimer" => {
            let addr = IoStopTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSynchronousCallDriver" => {
            let addr = IoSynchronousCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoSynchronousPageWrite" => {
            let addr = IoSynchronousPageWrite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoThreadToProcess" => {
            let addr = IoThreadToProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoTransferActivityId" => {
            let addr = IoTransferActivityId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoTranslateBusAddress" => {
            let addr = IoTranslateBusAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoTryQueueWorkItem" => {
            let addr = IoTryQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUninitializeWorkItem" => {
            let addr = IoUninitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterBootDriverCallback" => {
            let addr = IoUnregisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterContainerNotification" => {
            let addr = IoUnregisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterFileSystem" => {
            let addr = IoUnregisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterFsRegistrationChange" => {
            let addr = IoUnregisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterPlugPlayNotification" => {
            let addr = IoUnregisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterPlugPlayNotificationEx" => {
            let addr = IoUnregisterPlugPlayNotificationEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUnregisterShutdownNotification" => {
            let addr = IoUnregisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUpdateLinkShareAccess" => {
            let addr = IoUpdateLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUpdateLinkShareAccessEx" => {
            let addr = IoUpdateLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoUpdateShareAccess" => {
            let addr = IoUpdateShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoValidateDeviceIoControlAccess" => {
            let addr = IoValidateDeviceIoControlAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVerifyPartitionTable" => {
            let addr = IoVerifyPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVerifyVolume" => {
            let addr = IoVerifyVolume as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVolumeDeviceNameToGuid" => {
            let addr = IoVolumeDeviceNameToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVolumeDeviceNameToGuidPath" => {
            let addr = IoVolumeDeviceNameToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVolumeDeviceToDosName" => {
            let addr = IoVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVolumeDeviceToGuid" => {
            let addr = IoVolumeDeviceToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoVolumeDeviceToGuidPath" => {
            let addr = IoVolumeDeviceToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIAllocateInstanceIds" => {
            let addr = IoWMIAllocateInstanceIds as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIDeviceObjectToInstanceName" => {
            let addr = IoWMIDeviceObjectToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIDeviceObjectToProviderId" => {
            let addr = IoWMIDeviceObjectToProviderId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIExecuteMethod" => {
            let addr = IoWMIExecuteMethod as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIHandleToInstanceName" => {
            let addr = IoWMIHandleToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIOpenBlock" => {
            let addr = IoWMIOpenBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIQueryAllData" => {
            let addr = IoWMIQueryAllData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIQueryAllDataMultiple" => {
            let addr = IoWMIQueryAllDataMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIQuerySingleInstance" => {
            let addr = IoWMIQuerySingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIQuerySingleInstanceMultiple" => {
            let addr = IoWMIQuerySingleInstanceMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIRegistrationControl" => {
            let addr = IoWMIRegistrationControl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMISetNotificationCallback" => {
            let addr = IoWMISetNotificationCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMISetSingleInstance" => {
            let addr = IoWMISetSingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMISetSingleItem" => {
            let addr = IoWMISetSingleItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMISuggestInstanceName" => {
            let addr = IoWMISuggestInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWMIWriteEvent" => {
            let addr = IoWMIWriteEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWithinStackLimits" => {
            let addr = IoWithinStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWriteErrorLogEntry" => {
            let addr = IoWriteErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWriteKsrPersistentMemory" => {
            let addr = IoWriteKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWritePartitionTable" => {
            let addr = IoWritePartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IoWritePartitionTableEx" => {
            let addr = IoWritePartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IofCallDriver" => {
            let addr = IofCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IofCompleteRequest" => {
            let addr = IofCompleteRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "IofGetDriverProxyWrapperFromEndpoint" => {
            let addr = IofGetDriverProxyWrapperFromEndpoint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireGuardedMutex" => {
            let addr = KeAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireGuardedMutexUnsafe" => {
            let addr = KeAcquireGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLock" => {
            let addr = KeAcquireInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLockAtDpcLevel" => {
            let addr = KeAcquireInStackQueuedSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLockForDpc" => {
            let addr = KeAcquireInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireInterruptSpinLock" => {
            let addr = KeAcquireInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireQueuedSpinLock" => {
            let addr = KeAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireSpinLockAtDpcLevel" => {
            let addr = KeAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireSpinLockForDpc" => {
            let addr = KeAcquireSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireSpinLockRaiseToDpc" => {
            let addr = KeAcquireSpinLockRaiseToDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAcquireSpinLockRaiseToSynch" => {
            let addr = KeAcquireSpinLockRaiseToSynch as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAddTriageDumpDataBlock" => {
            let addr = KeAddTriageDumpDataBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAreAllApcsDisabled" => {
            let addr = KeAreAllApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAreApcsDisabled" => {
            let addr = KeAreApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeAttachProcess" => {
            let addr = KeAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeBugCheck" => {
            let addr = KeBugCheck as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeBugCheckEx" => {
            let addr = KeBugCheckEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeCancelTimer" => {
            let addr = KeCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeClearEvent" => {
            let addr = KeClearEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeConvertAuxiliaryCounterToPerformanceCounter" => {
            let addr = KeConvertAuxiliaryCounterToPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeConvertPerformanceCounterToAuxiliaryCounter" => {
            let addr = KeConvertPerformanceCounterToAuxiliaryCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDelayExecutionThread" => {
            let addr = KeDelayExecutionThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDeregisterBoundCallback" => {
            let addr = KeDeregisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDeregisterBugCheckCallback" => {
            let addr = KeDeregisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDeregisterBugCheckReasonCallback" => {
            let addr = KeDeregisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDeregisterNmiCallback" => {
            let addr = KeDeregisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDeregisterProcessorChangeCallback" => {
            let addr = KeDeregisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeDetachProcess" => {
            let addr = KeDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeEnterCriticalRegion" => {
            let addr = KeEnterCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeEnterGuardedRegion" => {
            let addr = KeEnterGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeExpandKernelStackAndCallout" => {
            let addr = KeExpandKernelStackAndCallout as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeExpandKernelStackAndCalloutEx" => {
            let addr = KeExpandKernelStackAndCalloutEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeFlushIoBuffers" => {
            let addr = KeFlushIoBuffers as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeFlushQueuedDpcs" => {
            let addr = KeFlushQueuedDpcs as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeFlushWriteBuffer" => {
            let addr = KeFlushWriteBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetCurrentIrql" => {
            let addr = KeGetCurrentIrql as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetCurrentNodeNumber" => {
            let addr = KeGetCurrentNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetCurrentProcessorNumberEx" => {
            let addr = KeGetCurrentProcessorNumberEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetProcessorIndexFromNumber" => {
            let addr = KeGetProcessorIndexFromNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetProcessorNumberFromIndex" => {
            let addr = KeGetProcessorNumberFromIndex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeGetRecommendedSharedDataAlignment" => {
            let addr = KeGetRecommendedSharedDataAlignment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeCrashDumpHeader" => {
            let addr = KeInitializeCrashDumpHeader as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeDeviceQueue" => {
            let addr = KeInitializeDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeDpc" => {
            let addr = KeInitializeDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeEvent" => {
            let addr = KeInitializeEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeGuardedMutex" => {
            let addr = KeInitializeGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeMutant" => {
            let addr = KeInitializeMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeMutex" => {
            let addr = KeInitializeMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeQueue" => {
            let addr = KeInitializeQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeSemaphore" => {
            let addr = KeInitializeSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeSpinLock" => {
            let addr = KeInitializeSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeThreadedDpc" => {
            let addr = KeInitializeThreadedDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeTimer" => {
            let addr = KeInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeTimerEx" => {
            let addr = KeInitializeTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInitializeTriageDumpDataArray" => {
            let addr = KeInitializeTriageDumpDataArray as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInsertByKeyDeviceQueue" => {
            let addr = KeInsertByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInsertDeviceQueue" => {
            let addr = KeInsertDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInsertHeadQueue" => {
            let addr = KeInsertHeadQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInsertQueue" => {
            let addr = KeInsertQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInsertQueueDpc" => {
            let addr = KeInsertQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInvalidateAllCaches" => {
            let addr = KeInvalidateAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeInvalidateRangeAllCaches" => {
            let addr = KeInvalidateRangeAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeIpiGenericCall" => {
            let addr = KeIpiGenericCall as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeIsExecutingDpc" => {
            let addr = KeIsExecutingDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeLeaveCriticalRegion" => {
            let addr = KeLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeLeaveGuardedRegion" => {
            let addr = KeLeaveGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeLowerIrql" => {
            let addr = KeLowerIrql as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KePulseEvent" => {
            let addr = KePulseEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryActiveGroupCount" => {
            let addr = KeQueryActiveGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryActiveProcessorCount" => {
            let addr = KeQueryActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryActiveProcessorCountEx" => {
            let addr = KeQueryActiveProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryActiveProcessors" => {
            let addr = KeQueryActiveProcessors as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryAuxiliaryCounterFrequency" => {
            let addr = KeQueryAuxiliaryCounterFrequency as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryDpcWatchdogInformation" => {
            let addr = KeQueryDpcWatchdogInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryGroupAffinity" => {
            let addr = KeQueryGroupAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryHardwareCounterConfiguration" => {
            let addr = KeQueryHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryHighestNodeNumber" => {
            let addr = KeQueryHighestNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryInterruptTimePrecise" => {
            let addr = KeQueryInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryLogicalProcessorRelationship" => {
            let addr = KeQueryLogicalProcessorRelationship as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryMaximumGroupCount" => {
            let addr = KeQueryMaximumGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryMaximumProcessorCount" => {
            let addr = KeQueryMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryMaximumProcessorCountEx" => {
            let addr = KeQueryMaximumProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryNodeActiveAffinity" => {
            let addr = KeQueryNodeActiveAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryNodeActiveAffinity2" => {
            let addr = KeQueryNodeActiveAffinity2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryNodeActiveProcessorCount" => {
            let addr = KeQueryNodeActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryNodeMaximumProcessorCount" => {
            let addr = KeQueryNodeMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryOwnerMutant" => {
            let addr = KeQueryOwnerMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryPerformanceCounter" => {
            let addr = KeQueryPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryPriorityThread" => {
            let addr = KeQueryPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryRuntimeThread" => {
            let addr = KeQueryRuntimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQuerySystemTimePrecise" => {
            let addr = KeQuerySystemTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryTimeIncrement" => {
            let addr = KeQueryTimeIncrement as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryTotalCycleTimeThread" => {
            let addr = KeQueryTotalCycleTimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryUnbiasedInterruptTime" => {
            let addr = KeQueryUnbiasedInterruptTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeQueryUnbiasedInterruptTimePrecise" => {
            let addr = KeQueryUnbiasedInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRcuReadLock" => {
            let addr = KeRcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRcuReadLockAtDpcLevel" => {
            let addr = KeRcuReadLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRcuReadUnlock" => {
            let addr = KeRcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRcuSynchronize" => {
            let addr = KeRcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateEvent" => {
            let addr = KeReadStateEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateMutant" => {
            let addr = KeReadStateMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateMutex" => {
            let addr = KeReadStateMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateQueue" => {
            let addr = KeReadStateQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateSemaphore" => {
            let addr = KeReadStateSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReadStateTimer" => {
            let addr = KeReadStateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRegisterBoundCallback" => {
            let addr = KeRegisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRegisterBugCheckCallback" => {
            let addr = KeRegisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRegisterBugCheckReasonCallback" => {
            let addr = KeRegisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRegisterNmiCallback" => {
            let addr = KeRegisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRegisterProcessorChangeCallback" => {
            let addr = KeRegisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseGuardedMutex" => {
            let addr = KeReleaseGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseGuardedMutexUnsafe" => {
            let addr = KeReleaseGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLock" => {
            let addr = KeReleaseInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLockForDpc" => {
            let addr = KeReleaseInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLockFromDpcLevel" => {
            let addr = KeReleaseInStackQueuedSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseInterruptSpinLock" => {
            let addr = KeReleaseInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseMutant" => {
            let addr = KeReleaseMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseMutex" => {
            let addr = KeReleaseMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseQueuedSpinLock" => {
            let addr = KeReleaseQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseSemaphore" => {
            let addr = KeReleaseSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseSpinLock" => {
            let addr = KeReleaseSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseSpinLockForDpc" => {
            let addr = KeReleaseSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeReleaseSpinLockFromDpcLevel" => {
            let addr = KeReleaseSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveByKeyDeviceQueue" => {
            let addr = KeRemoveByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveByKeyDeviceQueueIfBusy" => {
            let addr = KeRemoveByKeyDeviceQueueIfBusy as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveDeviceQueue" => {
            let addr = KeRemoveDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveEntryDeviceQueue" => {
            let addr = KeRemoveEntryDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveQueue" => {
            let addr = KeRemoveQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveQueueDpc" => {
            let addr = KeRemoveQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveQueueDpcEx" => {
            let addr = KeRemoveQueueDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRemoveQueueEx" => {
            let addr = KeRemoveQueueEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeResetEvent" => {
            let addr = KeResetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRestoreExtendedProcessorState" => {
            let addr = KeRestoreExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRevertToUserAffinityThread" => {
            let addr = KeRevertToUserAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRevertToUserAffinityThreadEx" => {
            let addr = KeRevertToUserAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRevertToUserGroupAffinityThread" => {
            let addr = KeRevertToUserGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeRundownQueue" => {
            let addr = KeRundownQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSaveExtendedProcessorState" => {
            let addr = KeSaveExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetBasePriorityThread" => {
            let addr = KeSetBasePriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetCoalescableTimer" => {
            let addr = KeSetCoalescableTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetEvent" => {
            let addr = KeSetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetHardwareCounterConfiguration" => {
            let addr = KeSetHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetIdealProcessorThread" => {
            let addr = KeSetIdealProcessorThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetImportanceDpc" => {
            let addr = KeSetImportanceDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetKernelStackSwapEnable" => {
            let addr = KeSetKernelStackSwapEnable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetPriorityThread" => {
            let addr = KeSetPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetSystemAffinityThread" => {
            let addr = KeSetSystemAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetSystemAffinityThreadEx" => {
            let addr = KeSetSystemAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetSystemGroupAffinityThread" => {
            let addr = KeSetSystemGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetTargetProcessorDpc" => {
            let addr = KeSetTargetProcessorDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetTargetProcessorDpcEx" => {
            let addr = KeSetTargetProcessorDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetTimer" => {
            let addr = KeSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSetTimerEx" => {
            let addr = KeSetTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeShouldYieldProcessor" => {
            let addr = KeShouldYieldProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSrcuAllocate" => {
            let addr = KeSrcuAllocate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSrcuFree" => {
            let addr = KeSrcuFree as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSrcuReadLock" => {
            let addr = KeSrcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSrcuReadUnlock" => {
            let addr = KeSrcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSrcuSynchronize" => {
            let addr = KeSrcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeStackAttachProcess" => {
            let addr = KeStackAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeStallExecutionProcessor" => {
            let addr = KeStallExecutionProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeSynchronizeExecution" => {
            let addr = KeSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeTestSpinLock" => {
            let addr = KeTestSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeTryToAcquireGuardedMutex" => {
            let addr = KeTryToAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeTryToAcquireQueuedSpinLock" => {
            let addr = KeTryToAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeTryToAcquireSpinLockAtDpcLevel" => {
            let addr = KeTryToAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeUnstackDetachProcess" => {
            let addr = KeUnstackDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeWaitForMultipleObjects" => {
            let addr = KeWaitForMultipleObjects as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "KeWaitForSingleObject" => {
            let addr = KeWaitForSingleObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAddPhysicalMemory" => {
            let addr = MmAddPhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAddVerifierSpecialThunks" => {
            let addr = MmAddVerifierSpecialThunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAddVerifierThunks" => {
            let addr = MmAddVerifierThunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAdvanceMdl" => {
            let addr = MmAdvanceMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateContiguousMemory" => {
            let addr = MmAllocateContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateContiguousMemoryEx" => {
            let addr = MmAllocateContiguousMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateContiguousMemorySpecifyCache" => {
            let addr = MmAllocateContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateContiguousMemorySpecifyCacheNode" => {
            let addr = MmAllocateContiguousMemorySpecifyCacheNode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateContiguousNodeMemory" => {
            let addr = MmAllocateContiguousNodeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateMappingAddress" => {
            let addr = MmAllocateMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateMappingAddressEx" => {
            let addr = MmAllocateMappingAddressEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateMdlForIoSpace" => {
            let addr = MmAllocateMdlForIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateNodePagesForMdlEx" => {
            let addr = MmAllocateNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocateNonCachedMemory" => {
            let addr = MmAllocateNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocatePagesForMdl" => {
            let addr = MmAllocatePagesForMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocatePagesForMdlEx" => {
            let addr = MmAllocatePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAllocatePartitionNodePagesForMdlEx" => {
            let addr = MmAllocatePartitionNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmAreMdlPagesCached" => {
            let addr = MmAreMdlPagesCached as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmBuildMdlForNonPagedPool" => {
            let addr = MmBuildMdlForNonPagedPool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmCanFileBeTruncated" => {
            let addr = MmCanFileBeTruncated as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmCopyMemory" => {
            let addr = MmCopyMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmCreateMdl" => {
            let addr = MmCreateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmCreateMirror" => {
            let addr = MmCreateMirror as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmDoesFileHaveUserWritableReferences" => {
            let addr = MmDoesFileHaveUserWritableReferences as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFlushImageSection" => {
            let addr = MmFlushImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmForceSectionClosed" => {
            let addr = MmForceSectionClosed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmForceSectionClosedEx" => {
            let addr = MmForceSectionClosedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreeContiguousMemory" => {
            let addr = MmFreeContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreeContiguousMemorySpecifyCache" => {
            let addr = MmFreeContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreeMappingAddress" => {
            let addr = MmFreeMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreeNonCachedMemory" => {
            let addr = MmFreeNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreePagesFromMdl" => {
            let addr = MmFreePagesFromMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmFreePagesFromMdlEx" => {
            let addr = MmFreePagesFromMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetCacheAttribute" => {
            let addr = MmGetCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetCacheAttributeEx" => {
            let addr = MmGetCacheAttributeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetMaximumFileSectionSize" => {
            let addr = MmGetMaximumFileSectionSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetModuleRoutineAddress" => {
            let addr = MmGetModuleRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetPhysicalAddress" => {
            let addr = MmGetPhysicalAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetPhysicalMemoryRanges" => {
            let addr = MmGetPhysicalMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetPhysicalMemoryRangesEx" => {
            let addr = MmGetPhysicalMemoryRangesEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetPhysicalMemoryRangesEx2" => {
            let addr = MmGetPhysicalMemoryRangesEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetSystemRoutineAddress" => {
            let addr = MmGetSystemRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmGetVirtualForPhysical" => {
            let addr = MmGetVirtualForPhysical as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsAddressValid" => {
            let addr = MmIsAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsDriverSuspectForVerifier" => {
            let addr = MmIsDriverSuspectForVerifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsDriverVerifying" => {
            let addr = MmIsDriverVerifying as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsDriverVerifyingByAddress" => {
            let addr = MmIsDriverVerifyingByAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsFileSectionActive" => {
            let addr = MmIsFileSectionActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsIoSpaceActive" => {
            let addr = MmIsIoSpaceActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsKernelAddress" => {
            let addr = MmIsKernelAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsNonPagedSystemAddressValid" => {
            let addr = MmIsNonPagedSystemAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsRecursiveIoFault" => {
            let addr = MmIsRecursiveIoFault as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsThisAnNtAsSystem" => {
            let addr = MmIsThisAnNtAsSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsUserAddress" => {
            let addr = MmIsUserAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmIsVerifierEnabled" => {
            let addr = MmIsVerifierEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmLockPagableDataSection" => {
            let addr = MmLockPagableDataSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmLockPagableSectionByHandle" => {
            let addr = MmLockPagableSectionByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapIoSpace" => {
            let addr = MmMapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapIoSpaceEx" => {
            let addr = MmMapIoSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapLockedPages" => {
            let addr = MmMapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapLockedPagesSpecifyCache" => {
            let addr = MmMapLockedPagesSpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapLockedPagesWithReservedMapping" => {
            let addr = MmMapLockedPagesWithReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapMdl" => {
            let addr = MmMapMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapMemoryDumpMdlEx" => {
            let addr = MmMapMemoryDumpMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapUserAddressesToPage" => {
            let addr = MmMapUserAddressesToPage as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapVideoDisplay" => {
            let addr = MmMapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapViewInSessionSpace" => {
            let addr = MmMapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapViewInSessionSpaceEx" => {
            let addr = MmMapViewInSessionSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapViewInSystemSpace" => {
            let addr = MmMapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMapViewInSystemSpaceEx" => {
            let addr = MmMapViewInSystemSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMdlPageContentsState" => {
            let addr = MmMdlPageContentsState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmMdlPagesAreZero" => {
            let addr = MmMdlPagesAreZero as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmPageEntireDriver" => {
            let addr = MmPageEntireDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmPrefetchPages" => {
            let addr = MmPrefetchPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProbeAndLockPages" => {
            let addr = MmProbeAndLockPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProbeAndLockPagesEx" => {
            let addr = MmProbeAndLockPagesEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProbeAndLockProcessPages" => {
            let addr = MmProbeAndLockProcessPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProbeAndLockSelectedPages" => {
            let addr = MmProbeAndLockSelectedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProtectDriverSection" => {
            let addr = MmProtectDriverSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmProtectMdlSystemAddress" => {
            let addr = MmProtectMdlSystemAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmQuerySystemSize" => {
            let addr = MmQuerySystemSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmRemovePhysicalMemory" => {
            let addr = MmRemovePhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmResetDriverPaging" => {
            let addr = MmResetDriverPaging as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmRotatePhysicalView" => {
            let addr = MmRotatePhysicalView as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmSecureVirtualMemory" => {
            let addr = MmSecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmSecureVirtualMemoryEx" => {
            let addr = MmSecureVirtualMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmSetAddressRangeModified" => {
            let addr = MmSetAddressRangeModified as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmSetPermanentCacheAttribute" => {
            let addr = MmSetPermanentCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmSizeOfMdl" => {
            let addr = MmSizeOfMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnlockPagableImageSection" => {
            let addr = MmUnlockPagableImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnlockPages" => {
            let addr = MmUnlockPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapIoSpace" => {
            let addr = MmUnmapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapLockedPages" => {
            let addr = MmUnmapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapReservedMapping" => {
            let addr = MmUnmapReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapVideoDisplay" => {
            let addr = MmUnmapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapViewInSessionSpace" => {
            let addr = MmUnmapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnmapViewInSystemSpace" => {
            let addr = MmUnmapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "MmUnsecureVirtualMemory" => {
            let addr = MmUnsecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAccessCheckAndAuditAlarm" => {
            let addr = NtAccessCheckAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAccessCheckByTypeAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarmByHandle" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarmByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAdjustGroupsToken" => {
            let addr = NtAdjustGroupsToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAdjustPrivilegesToken" => {
            let addr = NtAdjustPrivilegesToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtAllocateVirtualMemory" => {
            let addr = NtAllocateVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtClose" => {
            let addr = NtClose as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCloseObjectAuditAlarm" => {
            let addr = NtCloseObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCommitComplete" => {
            let addr = NtCommitComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCommitEnlistment" => {
            let addr = NtCommitEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCommitTransaction" => {
            let addr = NtCommitTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCopyFileChunk" => {
            let addr = NtCopyFileChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateEnlistment" => {
            let addr = NtCreateEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateFile" => {
            let addr = NtCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateResourceManager" => {
            let addr = NtCreateResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateSection" => {
            let addr = NtCreateSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateSectionEx" => {
            let addr = NtCreateSectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateTransaction" => {
            let addr = NtCreateTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtCreateTransactionManager" => {
            let addr = NtCreateTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtDeleteObjectAuditAlarm" => {
            let addr = NtDeleteObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtDeviceIoControlFile" => {
            let addr = NtDeviceIoControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtDuplicateToken" => {
            let addr = NtDuplicateToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtEnumerateTransactionObject" => {
            let addr = NtEnumerateTransactionObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtFilterToken" => {
            let addr = NtFilterToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtFlushBuffersFileEx" => {
            let addr = NtFlushBuffersFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtFreeVirtualMemory" => {
            let addr = NtFreeVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtFsControlFile" => {
            let addr = NtFsControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtGetNotificationResourceManager" => {
            let addr = NtGetNotificationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtImpersonateAnonymousToken" => {
            let addr = NtImpersonateAnonymousToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtLockFile" => {
            let addr = NtLockFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtManagePartition" => {
            let addr = NtManagePartition as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenEnlistment" => {
            let addr = NtOpenEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenFile" => {
            let addr = NtOpenFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenJobObjectToken" => {
            let addr = NtOpenJobObjectToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenObjectAuditAlarm" => {
            let addr = NtOpenObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenProcess" => {
            let addr = NtOpenProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenProcessToken" => {
            let addr = NtOpenProcessToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenProcessTokenEx" => {
            let addr = NtOpenProcessTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenRegistryTransaction" => {
            let addr = NtOpenRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenResourceManager" => {
            let addr = NtOpenResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenThreadToken" => {
            let addr = NtOpenThreadToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenThreadTokenEx" => {
            let addr = NtOpenThreadTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenTransaction" => {
            let addr = NtOpenTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtOpenTransactionManager" => {
            let addr = NtOpenTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPowerInformation" => {
            let addr = NtPowerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrePrepareComplete" => {
            let addr = NtPrePrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrePrepareEnlistment" => {
            let addr = NtPrePrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrepareComplete" => {
            let addr = NtPrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrepareEnlistment" => {
            let addr = NtPrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrivilegeCheck" => {
            let addr = NtPrivilegeCheck as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrivilegeObjectAuditAlarm" => {
            let addr = NtPrivilegeObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPrivilegedServiceAuditAlarm" => {
            let addr = NtPrivilegedServiceAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPropagationComplete" => {
            let addr = NtPropagationComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtPropagationFailed" => {
            let addr = NtPropagationFailed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryDirectoryFile" => {
            let addr = NtQueryDirectoryFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryDirectoryFileEx" => {
            let addr = NtQueryDirectoryFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationByName" => {
            let addr = NtQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationEnlistment" => {
            let addr = NtQueryInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationFile" => {
            let addr = NtQueryInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationResourceManager" => {
            let addr = NtQueryInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationToken" => {
            let addr = NtQueryInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationTransaction" => {
            let addr = NtQueryInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryInformationTransactionManager" => {
            let addr = NtQueryInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryObject" => {
            let addr = NtQueryObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryQuotaInformationFile" => {
            let addr = NtQueryQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQuerySecurityObject" => {
            let addr = NtQuerySecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryVirtualMemory" => {
            let addr = NtQueryVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtQueryVolumeInformationFile" => {
            let addr = NtQueryVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtReadFile" => {
            let addr = NtReadFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtReadOnlyEnlistment" => {
            let addr = NtReadOnlyEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRecoverEnlistment" => {
            let addr = NtRecoverEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRecoverResourceManager" => {
            let addr = NtRecoverResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRecoverTransactionManager" => {
            let addr = NtRecoverTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRegisterProtocolAddressInformation" => {
            let addr = NtRegisterProtocolAddressInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRenameTransactionManager" => {
            let addr = NtRenameTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRollbackComplete" => {
            let addr = NtRollbackComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRollbackEnlistment" => {
            let addr = NtRollbackEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRollbackRegistryTransaction" => {
            let addr = NtRollbackRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRollbackTransaction" => {
            let addr = NtRollbackTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtRollforwardTransactionManager" => {
            let addr = NtRollforwardTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationEnlistment" => {
            let addr = NtSetInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationFile" => {
            let addr = NtSetInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationResourceManager" => {
            let addr = NtSetInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationThread" => {
            let addr = NtSetInformationThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationToken" => {
            let addr = NtSetInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationTransaction" => {
            let addr = NtSetInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationTransactionManager" => {
            let addr = NtSetInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetInformationVirtualMemory" => {
            let addr = NtSetInformationVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetQuotaInformationFile" => {
            let addr = NtSetQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetSecurityObject" => {
            let addr = NtSetSecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSetVolumeInformationFile" => {
            let addr = NtSetVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtSinglePhaseReject" => {
            let addr = NtSinglePhaseReject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtUnlockFile" => {
            let addr = NtUnlockFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "NtWriteFile" => {
            let addr = NtWriteFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObCloseHandle" => {
            let addr = ObCloseHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObCloseHandleWithResult" => {
            let addr = ObCloseHandleWithResult as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObDereferenceObjectDeferDelete" => {
            let addr = ObDereferenceObjectDeferDelete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObDereferenceObjectDeferDeleteWithTag" => {
            let addr = ObDereferenceObjectDeferDeleteWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObGetFilterVersion" => {
            let addr = ObGetFilterVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObGetObjectSecurity" => {
            let addr = ObGetObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObInsertObject" => {
            let addr = ObInsertObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObIsKernelHandle" => {
            let addr = ObIsKernelHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObMakeTemporaryObject" => {
            let addr = ObMakeTemporaryObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObOpenObjectByPointer" => {
            let addr = ObOpenObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObOpenObjectByPointerWithTag" => {
            let addr = ObOpenObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObQueryNameString" => {
            let addr = ObQueryNameString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObQueryObjectAuditingByHandle" => {
            let addr = ObQueryObjectAuditingByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectByHandle" => {
            let addr = ObReferenceObjectByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectByHandleWithTag" => {
            let addr = ObReferenceObjectByHandleWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectByPointer" => {
            let addr = ObReferenceObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectByPointerWithTag" => {
            let addr = ObReferenceObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectSafe" => {
            let addr = ObReferenceObjectSafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReferenceObjectSafeWithTag" => {
            let addr = ObReferenceObjectSafeWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObRegisterCallbacks" => {
            let addr = ObRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObReleaseObjectSecurity" => {
            let addr = ObReleaseObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObUnRegisterCallbacks" => {
            let addr = ObUnRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObfDereferenceObject" => {
            let addr = ObfDereferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObfDereferenceObjectWithTag" => {
            let addr = ObfDereferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObfReferenceObject" => {
            let addr = ObfReferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "ObfReferenceObjectWithTag" => {
            let addr = ObfReferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsAcquireSiloHardReference" => {
            let addr = PsAcquireSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsAllocSiloContextSlot" => {
            let addr = PsAllocSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsAllocateAffinityToken" => {
            let addr = PsAllocateAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsAssignImpersonationToken" => {
            let addr = PsAssignImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsAttachSiloToCurrentThread" => {
            let addr = PsAttachSiloToCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsChargePoolQuota" => {
            let addr = PsChargePoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsChargeProcessPoolQuota" => {
            let addr = PsChargeProcessPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsCreateSiloContext" => {
            let addr = PsCreateSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsCreateSystemThread" => {
            let addr = PsCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsDereferenceImpersonationToken" => {
            let addr = PsDereferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsDereferencePrimaryToken" => {
            let addr = PsDereferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsDereferenceSiloContext" => {
            let addr = PsDereferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsDetachSiloFromCurrentThread" => {
            let addr = PsDetachSiloFromCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsDisableImpersonation" => {
            let addr = PsDisableImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsFreeAffinityToken" => {
            let addr = PsFreeAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsFreeSiloContextSlot" => {
            let addr = PsFreeSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentProcessId" => {
            let addr = PsGetCurrentProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentServerSilo" => {
            let addr = PsGetCurrentServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentServerSiloName" => {
            let addr = PsGetCurrentServerSiloName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentSilo" => {
            let addr = PsGetCurrentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentThreadId" => {
            let addr = PsGetCurrentThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetCurrentThreadTeb" => {
            let addr = PsGetCurrentThreadTeb as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetEffectiveServerSilo" => {
            let addr = PsGetEffectiveServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetHostSilo" => {
            let addr = PsGetHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetJobServerSilo" => {
            let addr = PsGetJobServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetJobSilo" => {
            let addr = PsGetJobSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetParentSilo" => {
            let addr = PsGetParentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetPermanentSiloContext" => {
            let addr = PsGetPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetProcessCreateTimeQuadPart" => {
            let addr = PsGetProcessCreateTimeQuadPart as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetProcessExitStatus" => {
            let addr = PsGetProcessExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetProcessExitTime" => {
            let addr = PsGetProcessExitTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetProcessId" => {
            let addr = PsGetProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetProcessStartKey" => {
            let addr = PsGetProcessStartKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetServerSiloActiveConsoleId" => {
            let addr = PsGetServerSiloActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetServerSiloServiceSessionId" => {
            let addr = PsGetServerSiloServiceSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetSiloContainerId" => {
            let addr = PsGetSiloContainerId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetSiloContext" => {
            let addr = PsGetSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetSiloMonitorContextSlot" => {
            let addr = PsGetSiloMonitorContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadCreateTime" => {
            let addr = PsGetThreadCreateTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadExitStatus" => {
            let addr = PsGetThreadExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadId" => {
            let addr = PsGetThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadProcess" => {
            let addr = PsGetThreadProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadProcessId" => {
            let addr = PsGetThreadProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadProperty" => {
            let addr = PsGetThreadProperty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetThreadServerSilo" => {
            let addr = PsGetThreadServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsGetVersion" => {
            let addr = PsGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsImpersonateClient" => {
            let addr = PsImpersonateClient as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsInsertPermanentSiloContext" => {
            let addr = PsInsertPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsInsertSiloContext" => {
            let addr = PsInsertSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsCurrentThreadInServerSilo" => {
            let addr = PsIsCurrentThreadInServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsCurrentThreadPrefetching" => {
            let addr = PsIsCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsDiskCountersEnabled" => {
            let addr = PsIsDiskCountersEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsHostSilo" => {
            let addr = PsIsHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsSystemThread" => {
            let addr = PsIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsThreadAttachedToSpecificSilo" => {
            let addr = PsIsThreadAttachedToSpecificSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsIsThreadTerminating" => {
            let addr = PsIsThreadTerminating as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsLookupProcessByProcessId" => {
            let addr = PsLookupProcessByProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsLookupThreadByThreadId" => {
            let addr = PsLookupThreadByThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsMakeSiloContextPermanent" => {
            let addr = PsMakeSiloContextPermanent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsQueryProcessAvailableCpus" => {
            let addr = PsQueryProcessAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsQueryProcessAvailableCpusCount" => {
            let addr = PsQueryProcessAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsQuerySystemAvailableCpus" => {
            let addr = PsQuerySystemAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsQuerySystemAvailableCpusCount" => {
            let addr = PsQuerySystemAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsQueryTotalCycleTimeProcess" => {
            let addr = PsQueryTotalCycleTimeProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReferenceImpersonationToken" => {
            let addr = PsReferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReferencePrimaryToken" => {
            let addr = PsReferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReferenceSiloContext" => {
            let addr = PsReferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRegisterProcessAvailableCpusChangeNotification" => {
            let addr = PsRegisterProcessAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRegisterSiloMonitor" => {
            let addr = PsRegisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRegisterSystemAvailableCpusChangeNotification" => {
            let addr = PsRegisterSystemAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReleaseSiloHardReference" => {
            let addr = PsReleaseSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRemoveCreateThreadNotifyRoutine" => {
            let addr = PsRemoveCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRemoveLoadImageNotifyRoutine" => {
            let addr = PsRemoveLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRemoveSiloContext" => {
            let addr = PsRemoveSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReplaceSiloContext" => {
            let addr = PsReplaceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRestoreImpersonation" => {
            let addr = PsRestoreImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsReturnPoolQuota" => {
            let addr = PsReturnPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRevertToSelf" => {
            let addr = PsRevertToSelf as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsRevertToUserMultipleGroupAffinityThread" => {
            let addr = PsRevertToUserMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutine" => {
            let addr = PsSetCreateProcessNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutineEx" => {
            let addr = PsSetCreateProcessNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutineEx2" => {
            let addr = PsSetCreateProcessNotifyRoutineEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCreateThreadNotifyRoutine" => {
            let addr = PsSetCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCreateThreadNotifyRoutineEx" => {
            let addr = PsSetCreateThreadNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetCurrentThreadPrefetching" => {
            let addr = PsSetCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetLoadImageNotifyRoutine" => {
            let addr = PsSetLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetLoadImageNotifyRoutineEx" => {
            let addr = PsSetLoadImageNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsSetSystemMultipleGroupAffinityThread" => {
            let addr = PsSetSystemMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsStartSiloMonitor" => {
            let addr = PsStartSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsTerminateServerSilo" => {
            let addr = PsTerminateServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsTerminateSystemThread" => {
            let addr = PsTerminateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsUnregisterAvailableCpusChangeNotification" => {
            let addr = PsUnregisterAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsUnregisterSiloMonitor" => {
            let addr = PsUnregisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsUpdateDiskCounters" => {
            let addr = PsUpdateDiskCounters as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PsWrapApcWow64Thread" => {
            let addr = PsWrapApcWow64Thread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedAllocateMemory" => {
            let addr = PshedAllocateMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedFreeMemory" => {
            let addr = PshedFreeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedIsSystemWheaEnabled" => {
            let addr = PshedIsSystemWheaEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedRegisterPlugin" => {
            let addr = PshedRegisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedSynchronizeExecution" => {
            let addr = PshedSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "PshedUnregisterPlugin" => {
            let addr = PshedUnregisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAbsoluteToSelfRelativeSD" => {
            let addr = RtlAbsoluteToSelfRelativeSD as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAddAccessAllowedAce" => {
            let addr = RtlAddAccessAllowedAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAddAccessAllowedAceEx" => {
            let addr = RtlAddAccessAllowedAceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAddAce" => {
            let addr = RtlAddAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAllocateAndInitializeSid" => {
            let addr = RtlAllocateAndInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAllocateAndInitializeSidEx" => {
            let addr = RtlAllocateAndInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAllocateHeap" => {
            let addr = RtlAllocateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAnsiStringToUnicodeString" => {
            let addr = RtlAnsiStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAppendStringToString" => {
            let addr = RtlAppendStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAppendUnicodeStringToString" => {
            let addr = RtlAppendUnicodeStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAppendUnicodeToString" => {
            let addr = RtlAppendUnicodeToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAreBitsClear" => {
            let addr = RtlAreBitsClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAreBitsSet" => {
            let addr = RtlAreBitsSet as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlAssert" => {
            let addr = RtlAssert as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCaptureContext" => {
            let addr = RtlCaptureContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCaptureContext2" => {
            let addr = RtlCaptureContext2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCaptureStackBackTrace" => {
            let addr = RtlCaptureStackBackTrace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCharToInteger" => {
            let addr = RtlCharToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCheckRegistryKey" => {
            let addr = RtlCheckRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlClearAllBits" => {
            let addr = RtlClearAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlClearBit" => {
            let addr = RtlClearBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlClearBits" => {
            let addr = RtlClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCmDecodeMemIoResource" => {
            let addr = RtlCmDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCmEncodeMemIoResource" => {
            let addr = RtlCmEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareAltitudes" => {
            let addr = RtlCompareAltitudes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareMemory" => {
            let addr = RtlCompareMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareMemoryUlong" => {
            let addr = RtlCompareMemoryUlong as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareString" => {
            let addr = RtlCompareString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareUnicodeString" => {
            let addr = RtlCompareUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompareUnicodeStrings" => {
            let addr = RtlCompareUnicodeStrings as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompressBuffer" => {
            let addr = RtlCompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCompressChunks" => {
            let addr = RtlCompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlContractHashTable" => {
            let addr = RtlContractHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlConvertSidToUnicodeString" => {
            let addr = RtlConvertSidToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyBitMap" => {
            let addr = RtlCopyBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyDeviceMemory" => {
            let addr = RtlCopyDeviceMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyLuid" => {
            let addr = RtlCopyLuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyMemoryNonTemporal" => {
            let addr = RtlCopyMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopySid" => {
            let addr = RtlCopySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyString" => {
            let addr = RtlCopyString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyUnicodeString" => {
            let addr = RtlCopyUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCopyVolatileMemory" => {
            let addr = RtlCopyVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCrc32" => {
            let addr = RtlCrc32 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCrc64" => {
            let addr = RtlCrc64 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateAcl" => {
            let addr = RtlCreateAcl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateHashTable" => {
            let addr = RtlCreateHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateHashTableEx" => {
            let addr = RtlCreateHashTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateHeap" => {
            let addr = RtlCreateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateRegistryKey" => {
            let addr = RtlCreateRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateSecurityDescriptor" => {
            let addr = RtlCreateSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateSecurityDescriptorRelative" => {
            let addr = RtlCreateSecurityDescriptorRelative as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateServiceSid" => {
            let addr = RtlCreateServiceSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateSystemVolumeInformationFolder" => {
            let addr = RtlCreateSystemVolumeInformationFolder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateUnicodeString" => {
            let addr = RtlCreateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCreateVirtualAccountSid" => {
            let addr = RtlCreateVirtualAccountSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlCustomCPToUnicodeN" => {
            let addr = RtlCustomCPToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressBuffer" => {
            let addr = RtlDecompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressBufferEx" => {
            let addr = RtlDecompressBufferEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressBufferEx2" => {
            let addr = RtlDecompressBufferEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressChunks" => {
            let addr = RtlDecompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressFragment" => {
            let addr = RtlDecompressFragment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDecompressFragmentEx" => {
            let addr = RtlDecompressFragmentEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDelete" => {
            let addr = RtlDelete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteAce" => {
            let addr = RtlDeleteAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteElementGenericTable" => {
            let addr = RtlDeleteElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteElementGenericTableAvl" => {
            let addr = RtlDeleteElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteElementGenericTableAvlEx" => {
            let addr = RtlDeleteElementGenericTableAvlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteHashTable" => {
            let addr = RtlDeleteHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteNoSplay" => {
            let addr = RtlDeleteNoSplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDeleteRegistryValue" => {
            let addr = RtlDeleteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDescribeChunk" => {
            let addr = RtlDescribeChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDestroyHeap" => {
            let addr = RtlDestroyHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDowncaseUnicodeChar" => {
            let addr = RtlDowncaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDowncaseUnicodeString" => {
            let addr = RtlDowncaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDrainNonVolatileFlush" => {
            let addr = RtlDrainNonVolatileFlush as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlDuplicateUnicodeString" => {
            let addr = RtlDuplicateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEndEnumerationHashTable" => {
            let addr = RtlEndEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEndStrongEnumerationHashTable" => {
            let addr = RtlEndStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEndWeakEnumerationHashTable" => {
            let addr = RtlEndWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateEntryHashTable" => {
            let addr = RtlEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateGenericTable" => {
            let addr = RtlEnumerateGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateGenericTableAvl" => {
            let addr = RtlEnumerateGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateGenericTableLikeADirectory" => {
            let addr = RtlEnumerateGenericTableLikeADirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateGenericTableWithoutSplaying" => {
            let addr = RtlEnumerateGenericTableWithoutSplaying as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEnumerateGenericTableWithoutSplayingAvl" => {
            let addr = RtlEnumerateGenericTableWithoutSplayingAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEqualPrefixSid" => {
            let addr = RtlEqualPrefixSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEqualSid" => {
            let addr = RtlEqualSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEqualString" => {
            let addr = RtlEqualString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlEqualUnicodeString" => {
            let addr = RtlEqualUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlExpandHashTable" => {
            let addr = RtlExpandHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlExtendCorrelationVector" => {
            let addr = RtlExtendCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlExtractBitMap" => {
            let addr = RtlExtractBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFillMemoryNonTemporal" => {
            let addr = RtlFillMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFillNonVolatileMemory" => {
            let addr = RtlFillNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindClearBits" => {
            let addr = RtlFindClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindClearBitsAndSet" => {
            let addr = RtlFindClearBitsAndSet as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindClearRuns" => {
            let addr = RtlFindClearRuns as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindClosestEncodableLength" => {
            let addr = RtlFindClosestEncodableLength as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindFirstRunClear" => {
            let addr = RtlFindFirstRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindLastBackwardRunClear" => {
            let addr = RtlFindLastBackwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindLeastSignificantBit" => {
            let addr = RtlFindLeastSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindLongestRunClear" => {
            let addr = RtlFindLongestRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindMostSignificantBit" => {
            let addr = RtlFindMostSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindNextForwardRunClear" => {
            let addr = RtlFindNextForwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindSetBits" => {
            let addr = RtlFindSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindSetBitsAndClear" => {
            let addr = RtlFindSetBitsAndClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFindUnicodePrefix" => {
            let addr = RtlFindUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFlushNonVolatileMemory" => {
            let addr = RtlFlushNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFlushNonVolatileMemoryRanges" => {
            let addr = RtlFlushNonVolatileMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeAnsiString" => {
            let addr = RtlFreeAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeHeap" => {
            let addr = RtlFreeHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeNonVolatileToken" => {
            let addr = RtlFreeNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeOemString" => {
            let addr = RtlFreeOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeSid" => {
            let addr = RtlFreeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeUTF8String" => {
            let addr = RtlFreeUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlFreeUnicodeString" => {
            let addr = RtlFreeUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGUIDFromString" => {
            let addr = RtlGUIDFromString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGenerate8dot3Name" => {
            let addr = RtlGenerate8dot3Name as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGenerateClass5Guid" => {
            let addr = RtlGenerateClass5Guid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetAce" => {
            let addr = RtlGetAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetAcesBufferSize" => {
            let addr = RtlGetAcesBufferSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetActiveConsoleId" => {
            let addr = RtlGetActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetCallersAddress" => {
            let addr = RtlGetCallersAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetCompressionWorkSpaceSize" => {
            let addr = RtlGetCompressionWorkSpaceSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetConsoleSessionForegroundProcessId" => {
            let addr = RtlGetConsoleSessionForegroundProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetDaclSecurityDescriptor" => {
            let addr = RtlGetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetElementGenericTable" => {
            let addr = RtlGetElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetElementGenericTableAvl" => {
            let addr = RtlGetElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetEnabledExtendedFeatures" => {
            let addr = RtlGetEnabledExtendedFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetGroupSecurityDescriptor" => {
            let addr = RtlGetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetNextEntryHashTable" => {
            let addr = RtlGetNextEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetNonVolatileToken" => {
            let addr = RtlGetNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetNtProductType" => {
            let addr = RtlGetNtProductType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetNtSystemRoot" => {
            let addr = RtlGetNtSystemRoot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetOwnerSecurityDescriptor" => {
            let addr = RtlGetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetPersistedStateLocation" => {
            let addr = RtlGetPersistedStateLocation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetProductInfo" => {
            let addr = RtlGetProductInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetSaclSecurityDescriptor" => {
            let addr = RtlGetSaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetSuiteMask" => {
            let addr = RtlGetSuiteMask as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetSystemGlobalData" => {
            let addr = RtlGetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlGetVersion" => {
            let addr = RtlGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlHashUnicodeString" => {
            let addr = RtlHashUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIdentifierAuthoritySid" => {
            let addr = RtlIdentifierAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIdnToAscii" => {
            let addr = RtlIdnToAscii as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIdnToNameprepUnicode" => {
            let addr = RtlIdnToNameprepUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIdnToUnicode" => {
            let addr = RtlIdnToUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIncrementCorrelationVector" => {
            let addr = RtlIncrementCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitAnsiString" => {
            let addr = RtlInitAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitAnsiStringEx" => {
            let addr = RtlInitAnsiStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitCodePageTable" => {
            let addr = RtlInitCodePageTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitEnumerationHashTable" => {
            let addr = RtlInitEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitString" => {
            let addr = RtlInitString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitStringEx" => {
            let addr = RtlInitStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitStrongEnumerationHashTable" => {
            let addr = RtlInitStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitUTF8String" => {
            let addr = RtlInitUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitUTF8StringEx" => {
            let addr = RtlInitUTF8StringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitUnicodeString" => {
            let addr = RtlInitUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitUnicodeStringEx" => {
            let addr = RtlInitUnicodeStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitWeakEnumerationHashTable" => {
            let addr = RtlInitWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeBitMap" => {
            let addr = RtlInitializeBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeCorrelationVector" => {
            let addr = RtlInitializeCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeGenericTable" => {
            let addr = RtlInitializeGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeGenericTableAvl" => {
            let addr = RtlInitializeGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeSid" => {
            let addr = RtlInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeSidEx" => {
            let addr = RtlInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInitializeUnicodePrefix" => {
            let addr = RtlInitializeUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertElementGenericTable" => {
            let addr = RtlInsertElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertElementGenericTableAvl" => {
            let addr = RtlInsertElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertElementGenericTableFull" => {
            let addr = RtlInsertElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertElementGenericTableFullAvl" => {
            let addr = RtlInsertElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertEntryHashTable" => {
            let addr = RtlInsertEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInsertUnicodePrefix" => {
            let addr = RtlInsertUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlInt64ToUnicodeString" => {
            let addr = RtlInt64ToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIntegerToUnicodeString" => {
            let addr = RtlIntegerToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIoDecodeMemIoResource" => {
            let addr = RtlIoDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIoEncodeMemIoResource" => {
            let addr = RtlIoEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsApiSetImplemented" => {
            let addr = RtlIsApiSetImplemented as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsCloudFilesPlaceholder" => {
            let addr = RtlIsCloudFilesPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsGenericTableEmpty" => {
            let addr = RtlIsGenericTableEmpty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsGenericTableEmptyAvl" => {
            let addr = RtlIsGenericTableEmptyAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsMultiSessionSku" => {
            let addr = RtlIsMultiSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsMultiUsersInSessionSku" => {
            let addr = RtlIsMultiUsersInSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsNameLegalDOS8Dot3" => {
            let addr = RtlIsNameLegalDOS8Dot3 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsNonEmptyDirectoryReparsePointAllowed" => {
            let addr = RtlIsNonEmptyDirectoryReparsePointAllowed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsNormalizedString" => {
            let addr = RtlIsNormalizedString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsNtDdiVersionAvailable" => {
            let addr = RtlIsNtDdiVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsPartialPlaceholder" => {
            let addr = RtlIsPartialPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsPartialPlaceholderFileHandle" => {
            let addr = RtlIsPartialPlaceholderFileHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsPartialPlaceholderFileInfo" => {
            let addr = RtlIsPartialPlaceholderFileInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsSandboxedToken" => {
            let addr = RtlIsSandboxedToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsSandboxedTokenHandle" => {
            let addr = RtlIsSandboxedTokenHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsServicePackVersionInstalled" => {
            let addr = RtlIsServicePackVersionInstalled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsStateSeparationEnabled" => {
            let addr = RtlIsStateSeparationEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsUntrustedObject" => {
            let addr = RtlIsUntrustedObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsValidOemCharacter" => {
            let addr = RtlIsValidOemCharacter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlIsZeroMemory" => {
            let addr = RtlIsZeroMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLengthRequiredSid" => {
            let addr = RtlLengthRequiredSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLengthSecurityDescriptor" => {
            let addr = RtlLengthSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLengthSid" => {
            let addr = RtlLengthSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupElementGenericTable" => {
            let addr = RtlLookupElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupElementGenericTableAvl" => {
            let addr = RtlLookupElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupElementGenericTableFull" => {
            let addr = RtlLookupElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupElementGenericTableFullAvl" => {
            let addr = RtlLookupElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupEntryHashTable" => {
            let addr = RtlLookupEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlLookupFirstMatchingElementGenericTableAvl" => {
            let addr = RtlLookupFirstMatchingElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlMapGenericMask" => {
            let addr = RtlMapGenericMask as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlMoveVolatileMemory" => {
            let addr = RtlMoveVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlMultiByteToUnicodeN" => {
            let addr = RtlMultiByteToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlMultiByteToUnicodeSize" => {
            let addr = RtlMultiByteToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNextUnicodePrefix" => {
            let addr = RtlNextUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNormalizeSecurityDescriptor" => {
            let addr = RtlNormalizeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNormalizeString" => {
            let addr = RtlNormalizeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNtStatusToDosError" => {
            let addr = RtlNtStatusToDosError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNtStatusToDosErrorNoTeb" => {
            let addr = RtlNtStatusToDosErrorNoTeb as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberGenericTableElements" => {
            let addr = RtlNumberGenericTableElements as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberGenericTableElementsAvl" => {
            let addr = RtlNumberGenericTableElementsAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberOfClearBits" => {
            let addr = RtlNumberOfClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberOfClearBitsInRange" => {
            let addr = RtlNumberOfClearBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberOfSetBits" => {
            let addr = RtlNumberOfSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberOfSetBitsInRange" => {
            let addr = RtlNumberOfSetBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlNumberOfSetBitsUlongPtr" => {
            let addr = RtlNumberOfSetBitsUlongPtr as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlOemStringToCountedUnicodeString" => {
            let addr = RtlOemStringToCountedUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlOemStringToUnicodeString" => {
            let addr = RtlOemStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlOemToUnicodeN" => {
            let addr = RtlOemToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlOsDeploymentState" => {
            let addr = RtlOsDeploymentState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlPrefetchMemoryNonTemporal" => {
            let addr = RtlPrefetchMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlPrefixString" => {
            let addr = RtlPrefixString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlPrefixUnicodeString" => {
            let addr = RtlPrefixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryPackageIdentity" => {
            let addr = RtlQueryPackageIdentity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryPackageIdentityEx" => {
            let addr = RtlQueryPackageIdentityEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryProcessPlaceholderCompatibilityMode" => {
            let addr = RtlQueryProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryRegistryValueWithFallback" => {
            let addr = RtlQueryRegistryValueWithFallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryRegistryValues" => {
            let addr = RtlQueryRegistryValues as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryThreadPlaceholderCompatibilityMode" => {
            let addr = RtlQueryThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlQueryValidationRunlevel" => {
            let addr = RtlQueryValidationRunlevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRaiseCustomSystemEventTrigger" => {
            let addr = RtlRaiseCustomSystemEventTrigger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRandom" => {
            let addr = RtlRandom as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRandomEx" => {
            let addr = RtlRandomEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRealPredecessor" => {
            let addr = RtlRealPredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRealSuccessor" => {
            let addr = RtlRealSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRemoveEntryHashTable" => {
            let addr = RtlRemoveEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRemoveUnicodePrefix" => {
            let addr = RtlRemoveUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlReplaceSidInSd" => {
            let addr = RtlReplaceSidInSd as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlReserveChunk" => {
            let addr = RtlReserveChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRunOnceBeginInitialize" => {
            let addr = RtlRunOnceBeginInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRunOnceComplete" => {
            let addr = RtlRunOnceComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRunOnceExecuteOnce" => {
            let addr = RtlRunOnceExecuteOnce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlRunOnceInitialize" => {
            let addr = RtlRunOnceInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSecondsSince1970ToTime" => {
            let addr = RtlSecondsSince1970ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSecondsSince1980ToTime" => {
            let addr = RtlSecondsSince1980ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSelfRelativeToAbsoluteSD" => {
            let addr = RtlSelfRelativeToAbsoluteSD as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetAllBits" => {
            let addr = RtlSetAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetBit" => {
            let addr = RtlSetBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetBits" => {
            let addr = RtlSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetDaclSecurityDescriptor" => {
            let addr = RtlSetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetGroupSecurityDescriptor" => {
            let addr = RtlSetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetOwnerSecurityDescriptor" => {
            let addr = RtlSetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetProcessPlaceholderCompatibilityMode" => {
            let addr = RtlSetProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetSystemGlobalData" => {
            let addr = RtlSetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetThreadPlaceholderCompatibilityMode" => {
            let addr = RtlSetThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSetVolatileMemory" => {
            let addr = RtlSetVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSplay" => {
            let addr = RtlSplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlStringFromGUID" => {
            let addr = RtlStringFromGUID as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlStronglyEnumerateEntryHashTable" => {
            let addr = RtlStronglyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSubAuthorityCountSid" => {
            let addr = RtlSubAuthorityCountSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSubAuthoritySid" => {
            let addr = RtlSubAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSubtreePredecessor" => {
            let addr = RtlSubtreePredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSubtreeSuccessor" => {
            let addr = RtlSubtreeSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlSuffixUnicodeString" => {
            let addr = RtlSuffixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlTestBit" => {
            let addr = RtlTestBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlTimeFieldsToTime" => {
            let addr = RtlTimeFieldsToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlTimeToSecondsSince1970" => {
            let addr = RtlTimeToSecondsSince1970 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlTimeToSecondsSince1980" => {
            let addr = RtlTimeToSecondsSince1980 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlTimeToTimeFields" => {
            let addr = RtlTimeToTimeFields as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUTF8StringToUnicodeString" => {
            let addr = RtlUTF8StringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUTF8ToUnicodeN" => {
            let addr = RtlUTF8ToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToAnsiString" => {
            let addr = RtlUnicodeStringToAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToCountedOemString" => {
            let addr = RtlUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToInt64" => {
            let addr = RtlUnicodeStringToInt64 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToInteger" => {
            let addr = RtlUnicodeStringToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToOemString" => {
            let addr = RtlUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeStringToUTF8String" => {
            let addr = RtlUnicodeStringToUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeToCustomCPN" => {
            let addr = RtlUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeToMultiByteN" => {
            let addr = RtlUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeToMultiByteSize" => {
            let addr = RtlUnicodeToMultiByteSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeToOemN" => {
            let addr = RtlUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUnicodeToUTF8N" => {
            let addr = RtlUnicodeToUTF8N as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeChar" => {
            let addr = RtlUpcaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeString" => {
            let addr = RtlUpcaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeStringToCountedOemString" => {
            let addr = RtlUpcaseUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeStringToOemString" => {
            let addr = RtlUpcaseUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeToCustomCPN" => {
            let addr = RtlUpcaseUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeToMultiByteN" => {
            let addr = RtlUpcaseUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpcaseUnicodeToOemN" => {
            let addr = RtlUpcaseUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpperChar" => {
            let addr = RtlUpperChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlUpperString" => {
            let addr = RtlUpperString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlValidRelativeSecurityDescriptor" => {
            let addr = RtlValidRelativeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlValidSecurityDescriptor" => {
            let addr = RtlValidSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlValidSid" => {
            let addr = RtlValidSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlValidateCorrelationVector" => {
            let addr = RtlValidateCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlValidateUnicodeString" => {
            let addr = RtlValidateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlVerifyVersionInfo" => {
            let addr = RtlVerifyVersionInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlVolumeDeviceToDosName" => {
            let addr = RtlVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlWalkFrameChain" => {
            let addr = RtlWalkFrameChain as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlWeaklyEnumerateEntryHashTable" => {
            let addr = RtlWeaklyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlWriteNonVolatileMemory" => {
            let addr = RtlWriteNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlWriteRegistryValue" => {
            let addr = RtlWriteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlxAnsiStringToUnicodeSize" => {
            let addr = RtlxAnsiStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlxOemStringToUnicodeSize" => {
            let addr = RtlxOemStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlxUnicodeStringToAnsiSize" => {
            let addr = RtlxUnicodeStringToAnsiSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        "RtlxUnicodeStringToOemSize" => {
            let addr = RtlxUnicodeStringToOemSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }
        }
        _ => Err(HookError::NotHooked),
    }
}
