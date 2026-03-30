#![allow(non_snake_case)]
#![allow(dead_code)]

// Inline Hook Database: 1251 functions
// Inline hooks modify code directly, faster but detectable

use alloc::string::String;
use crate::hyperkd::hyperhv::hooks::hooks::*;
use crate::ntapi::*;

pub struct InlineHookDb {
    pub name: &'static str,
    pub address: u64,
    pub trampoline: u64,
    pub enabled: bool,
}

pub static INLINE_HOOK_DATABASE: &[InlineHookDb] = &[
    InlineHookDb { name: "DbgBreakPointWithStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgPrint", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgPrintEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgPrintReturnControlC", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgPrompt", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgQueryDebugFilterState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgSetDebugFilterState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "DbgSetDebugPrintCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireFastMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireFastMutexUnsafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquirePushLockExclusiveEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquirePushLockSharedEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireResourceExclusiveLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireResourceSharedLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireRundownProtectionCacheAwareEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireRundownProtectionEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSharedStarveExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSharedWaitForExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSpinLockExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSpinLockExclusiveAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSpinLockShared", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAcquireSpinLockSharedAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAdjustLookasideDepth", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocateCacheAwareRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocateFromLookasideListEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocateFromNPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocateFromPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocatePool2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocatePool3", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocatePoolWithQuota", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExAllocateTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExCancelTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExCleanupRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExConvertExclusiveToSharedLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExCreateCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExCreatePool", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDeleteLookasideListEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDeleteNPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDeletePagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDeleteResourceLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDeleteTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDestroyPool", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExDisableResourceBoostLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExEnterCriticalRegionAndAcquireResourceExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExEnterCriticalRegionAndAcquireResourceShared", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExEnumerateSystemFirmwareTables", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExExtendZone", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFlushLookasideListEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreeCacheAwareRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreePool", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreePool2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreePoolWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreeToLookasideListEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreeToNPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExFreeToPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetExclusiveWaiterCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetFirmwareEnvironmentVariable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetFirmwareType", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetPreviousMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetSharedWaiterCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExGetSystemFirmwareTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeDeviceAts", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeLookasideListEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeNPagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializePagedLookasideList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializePushLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeResourceLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeRundownProtectionCacheAwareEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInitializeZone", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedAddLargeInteger", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedAddUlong", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedExtendZone", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedInsertHeadList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedInsertTailList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedPopEntryList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedPushEntryList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExInterlockedRemoveHeadList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExIsManufacturingModeEnabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExIsProcessorFeaturePresent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExIsResourceAcquiredExclusiveLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExIsResourceAcquiredSharedLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExIsSoftBoot", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExLocalTimeToSystemTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExNotifyCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExQueryDepthSList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExQueryPoolBlockSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExQueryTimerResolution", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExQueueWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRaiseAccessViolation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRaiseDatatypeMisalignment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRaiseStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRcuFreePool", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReInitializeRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReInitializeRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRegisterCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReinitializeResourceLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseFastMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseFastMutexUnsafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleasePushLockExclusiveEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleasePushLockSharedEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseResourceAndLeaveCriticalRegion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseResourceForThreadLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseResourceLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseRundownProtection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseRundownProtectionCacheAwareEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseRundownProtectionEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseSpinLockExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseSpinLockExclusiveFromDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseSpinLockShared", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExReleaseSpinLockSharedFromDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRundownCompleted", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExRundownCompletedCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSecurePoolUpdate", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSecurePoolValidate", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSetFirmwareEnvironmentVariable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSetResourceOwnerPointer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSetResourceOwnerPointerEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSetTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSetTimerResolution", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSizeOfRundownProtectionCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExSystemTimeToLocalTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryAcquirePushLockExclusiveEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryAcquirePushLockSharedEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryAcquireSpinLockExclusiveAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryAcquireSpinLockSharedAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryConvertSharedSpinLockExclusive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExTryToAcquireFastMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExUnregisterCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExUuidCreate", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExVerifySuite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExWaitForRundownProtectionRelease", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExWaitForRundownProtectionReleaseCacheAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExpInterlockedFlushSList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExpInterlockedPopEntrySList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExpInterlockedPushEntrySList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ExportSecurityContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAcquireCancelSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAcquireKsrPersistentMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAcquireKsrPersistentMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAcquireRemoveLockEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAcquireVpbSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateController", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateDriverObjectExtension", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateErrorLogEntry", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateIrpEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateMiniCompletionPacket", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateSfioStreamIdentifier", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAllocateWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoApplyPriorityInfoThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAssignResources", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAttachDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAttachDeviceByPointer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAttachDeviceToDeviceStack", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoAttachDeviceToDeviceStackSafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoBuildAsynchronousFsdRequest", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoBuildDeviceIoControlRequest", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoBuildPartialMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoBuildSynchronousFsdRequest", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCancelFileOpen", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCancelIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckDesiredAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckEaBufferValidity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckFileObjectOpenedAsCopyDestination", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckFileObjectOpenedAsCopySource", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckFunctionAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckLinkShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckQuerySetFileInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckQuerySetVolumeInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckQuotaBufferValidity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCheckShareAccessEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCleanupIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoClearActivityIdThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoClearFsTrackOffsetState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoClearIrpExtraCreateParameter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoConnectInterrupt", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoConnectInterruptEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateController", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateDisk", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateDriverProxyExtension", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateFileEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateFileSpecifyDeviceObjectHint", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateNotificationEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateStreamFileObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateStreamFileObjectEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateStreamFileObjectEx2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateStreamFileObjectLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateSymbolicLink", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateSynchronizationEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateSystemThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCreateUnprotectedSymbolicLink", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqInitialize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqInitializeEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqInsertIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqInsertIrpEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqRemoveIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoCsqRemoveNextIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDecrementKeepAliveCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDeleteController", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDeleteDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDeleteSymbolicLink", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDetachDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDisconnectInterrupt", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDisconnectInterruptEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoDriverProxyCreateHotSwappableWorkerThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoEnumerateDeviceObjectList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoEnumerateKsrPersistentMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoEnumerateRegisteredFiltersList", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFastQueryNetworkAttributes", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoForwardIrpSynchronously", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeController", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeErrorLogEntry", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeKsrPersistentMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeMiniCompletionPacket", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeSfioStreamIdentifier", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoFreeWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetActivityIdIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetActivityIdThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetAffinityInterrupt", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetAttachedDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetAttachedDeviceReference", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetBaseFileSystemDeviceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetBootDiskInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetBootDiskInformationLite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetConfigurationInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetContainerInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetCopyInformationExtension", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetCurrentProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceAttachmentBaseRef", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceDirectory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceInterfaceAlias", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceInterfacePropertyData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceInterfaces", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceNumaNode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceObjectPointer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceProperty", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDevicePropertyData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDeviceToVerify", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDiskDeviceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDmaAdapter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverDirectory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverObjectExtension", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverProxyEndpointWrapper", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverProxyExtensionFromDriverObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverProxyExtensionVersion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetDriverProxyFeatures", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetFileObjectGenericMapping", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetFsTrackOffsetState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetFsZeroingOffset", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetInitialStack", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetInitiatorProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetIoAttributionHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetIoPriorityHint", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetIommuInterface", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetIommuInterfaceEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetIrpExtraCreateParameter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetKsrPersistentMemoryBuffer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetLowerDeviceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetOplockKeyContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetOplockKeyContextEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetPagingIoPriority", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetRelatedDeviceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetRequestorProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetRequestorProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetRequestorSessionId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetSfioStreamIdentifier", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetShadowFileInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetSiloParameters", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetStackLimits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetTopLevelIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoGetTransactionParameterBlock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoHotSwapDriverProxyEndpoints", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIncrementKeepAliveCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInitializeIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInitializeIrpEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInitializeRemoveLockEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInitializeTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInitializeWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInvalidateDeviceRelations", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoInvalidateDeviceState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIrpHasFsTrackOffsetExtensionType", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIs32bitProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsFileObjectIgnoringSharing", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsFileOriginRemote", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsInitiator32bitProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsOperationSynchronous", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsSystemThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsValidIrpStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsValidNameGraftingBuffer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoIsWdmVersionAvailable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoMakeAssociatedIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoMakeAssociatedIrpEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoMapKsrPersistentMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoOpenDeviceInterfaceRegistryKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoOpenDeviceRegistryKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoOpenDriverRegistryKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoPageRead", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoPropagateActivityIdToThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryDeviceDescription", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryDmaFeatureSupport", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryFileDosDeviceName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryFileInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryFullDriverPath", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryInformationByName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryKsrPersistentMemorySize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryKsrPersistentMemorySizeEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueryVolumeInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueueThreadIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueueWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoQueueWorkItemEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRaiseHardError", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRaiseInformationalHardError", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReadDiskSignature", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReadPartitionTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReadPartitionTableEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRecordIoAttribution", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterBootDriverCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterBootDriverReinitialization", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterContainerNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterDeviceInterface", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterDriverProxyEndpoints", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterDriverReinitialization", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterFileSystem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterFsRegistrationChange", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterFsRegistrationChangeMountAware", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterLastChanceShutdownNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterPlugPlayNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRegisterShutdownNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReleaseCancelSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReleaseRemoveLockAndWaitEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReleaseRemoveLockEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReleaseVpbSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRemoveIoCompletion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRemoveLinkShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRemoveLinkShareAccessEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRemoveShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReplaceFileObjectName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReplacePartitionUnit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportDetectedDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportInterruptActive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportInterruptInactive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportResourceForDetection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportResourceUsage", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportRootDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportTargetDeviceChange", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReportTargetDeviceChangeAsynchronous", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRequestDeviceEject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRequestDeviceEjectEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRequestDeviceRemovalForReset", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReserveKsrPersistentMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReserveKsrPersistentMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoRetrievePriorityInfo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoReuseIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetActivityIdIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetActivityIdThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetCompletionRoutineEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetDeviceInterfacePropertyData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetDeviceInterfaceState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetDevicePropertyData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetDeviceToVerify", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetFileObjectIgnoreSharing", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetFileOrigin", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetFsTrackOffsetState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetFsZeroingOffset", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetFsZeroingOffsetRequired", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetHardErrorOrVerifyDevice", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetIoAttributionIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetIoCompletionEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetIoPriorityHint", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetIrpExtraCreateParameter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetLinkShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetMasterIrpStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetPartitionInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetPartitionInformationEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetShadowFileInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetShareAccessEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetStartIoAttributes", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetSystemPartition", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetThreadHardErrorMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSetTopLevelIrp", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSizeOfIrpEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSizeofWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoStartNextPacket", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoStartNextPacketByKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoStartPacket", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoStartTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoStopTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSynchronousCallDriver", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoSynchronousPageWrite", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoThreadToProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoTransferActivityId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoTranslateBusAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoTryQueueWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUninitializeWorkItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterBootDriverCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterContainerNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterFileSystem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterFsRegistrationChange", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterPlugPlayNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterPlugPlayNotificationEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUnregisterShutdownNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUpdateLinkShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUpdateLinkShareAccessEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoUpdateShareAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoValidateDeviceIoControlAccess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVerifyPartitionTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVerifyVolume", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVolumeDeviceNameToGuid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVolumeDeviceNameToGuidPath", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVolumeDeviceToDosName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVolumeDeviceToGuid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoVolumeDeviceToGuidPath", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIAllocateInstanceIds", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIDeviceObjectToInstanceName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIDeviceObjectToProviderId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIExecuteMethod", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIHandleToInstanceName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIOpenBlock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIQueryAllData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIQueryAllDataMultiple", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIQuerySingleInstance", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIQuerySingleInstanceMultiple", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIRegistrationControl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMISetNotificationCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMISetSingleInstance", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMISetSingleItem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMISuggestInstanceName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWMIWriteEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWithinStackLimits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWriteErrorLogEntry", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWriteKsrPersistentMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWritePartitionTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IoWritePartitionTableEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IofCallDriver", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IofCompleteRequest", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "IofGetDriverProxyWrapperFromEndpoint", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireGuardedMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireGuardedMutexUnsafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireInStackQueuedSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireInStackQueuedSpinLockAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireInStackQueuedSpinLockForDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireInterruptSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireQueuedSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireSpinLockAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireSpinLockForDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireSpinLockRaiseToDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAcquireSpinLockRaiseToSynch", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAddTriageDumpDataBlock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAreAllApcsDisabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAreApcsDisabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeAttachProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeBugCheck", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeBugCheckEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeCancelTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeClearEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeConvertAuxiliaryCounterToPerformanceCounter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeConvertPerformanceCounterToAuxiliaryCounter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDelayExecutionThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDeregisterBoundCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDeregisterBugCheckCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDeregisterBugCheckReasonCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDeregisterNmiCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDeregisterProcessorChangeCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeDetachProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeEnterCriticalRegion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeEnterGuardedRegion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeExpandKernelStackAndCallout", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeExpandKernelStackAndCalloutEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeFlushIoBuffers", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeFlushQueuedDpcs", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeFlushWriteBuffer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetCurrentIrql", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetCurrentNodeNumber", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetCurrentProcessorNumberEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetProcessorIndexFromNumber", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetProcessorNumberFromIndex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeGetRecommendedSharedDataAlignment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeCrashDumpHeader", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeGuardedMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeMutant", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeSemaphore", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeThreadedDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeTimerEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInitializeTriageDumpDataArray", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInsertByKeyDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInsertDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInsertHeadQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInsertQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInsertQueueDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInvalidateAllCaches", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeInvalidateRangeAllCaches", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeIpiGenericCall", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeIsExecutingDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeLeaveCriticalRegion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeLeaveGuardedRegion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeLowerIrql", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KePulseEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryActiveGroupCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryActiveProcessorCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryActiveProcessorCountEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryActiveProcessors", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryAuxiliaryCounterFrequency", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryDpcWatchdogInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryGroupAffinity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryHardwareCounterConfiguration", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryHighestNodeNumber", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryInterruptTimePrecise", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryLogicalProcessorRelationship", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryMaximumGroupCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryMaximumProcessorCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryMaximumProcessorCountEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryNodeActiveAffinity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryNodeActiveAffinity2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryNodeActiveProcessorCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryNodeMaximumProcessorCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryOwnerMutant", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryPerformanceCounter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryPriorityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryRuntimeThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQuerySystemTimePrecise", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryTimeIncrement", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryTotalCycleTimeThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryUnbiasedInterruptTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeQueryUnbiasedInterruptTimePrecise", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRcuReadLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRcuReadLockAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRcuReadUnlock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRcuSynchronize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateMutant", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateSemaphore", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReadStateTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRegisterBoundCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRegisterBugCheckCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRegisterBugCheckReasonCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRegisterNmiCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRegisterProcessorChangeCallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseGuardedMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseGuardedMutexUnsafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseInStackQueuedSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseInStackQueuedSpinLockForDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseInStackQueuedSpinLockFromDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseInterruptSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseMutant", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseQueuedSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseSemaphore", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseSpinLockForDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeReleaseSpinLockFromDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveByKeyDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveByKeyDeviceQueueIfBusy", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveEntryDeviceQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveQueueDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveQueueDpcEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRemoveQueueEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeResetEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRestoreExtendedProcessorState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRevertToUserAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRevertToUserAffinityThreadEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRevertToUserGroupAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeRundownQueue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSaveExtendedProcessorState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetBasePriorityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetCoalescableTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetEvent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetHardwareCounterConfiguration", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetIdealProcessorThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetImportanceDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetKernelStackSwapEnable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetPriorityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetSystemAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetSystemAffinityThreadEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetSystemGroupAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetTargetProcessorDpc", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetTargetProcessorDpcEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetTimer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSetTimerEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeShouldYieldProcessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSrcuAllocate", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSrcuFree", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSrcuReadLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSrcuReadUnlock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSrcuSynchronize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeStackAttachProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeStallExecutionProcessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeSynchronizeExecution", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeTestSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeTryToAcquireGuardedMutex", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeTryToAcquireQueuedSpinLock", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeTryToAcquireSpinLockAtDpcLevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeUnstackDetachProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeWaitForMultipleObjects", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "KeWaitForSingleObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAddPhysicalMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAddVerifierSpecialThunks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAddVerifierThunks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAdvanceMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateContiguousMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateContiguousMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateContiguousMemorySpecifyCache", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateContiguousMemorySpecifyCacheNode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateContiguousNodeMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateMappingAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateMappingAddressEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateMdlForIoSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateNodePagesForMdlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocateNonCachedMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocatePagesForMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocatePagesForMdlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAllocatePartitionNodePagesForMdlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmAreMdlPagesCached", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmBuildMdlForNonPagedPool", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmCanFileBeTruncated", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmCopyMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmCreateMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmCreateMirror", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmDoesFileHaveUserWritableReferences", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFlushImageSection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmForceSectionClosed", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmForceSectionClosedEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreeContiguousMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreeContiguousMemorySpecifyCache", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreeMappingAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreeNonCachedMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreePagesFromMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmFreePagesFromMdlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetCacheAttribute", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetCacheAttributeEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetMaximumFileSectionSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetModuleRoutineAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetPhysicalAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetPhysicalMemoryRanges", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetPhysicalMemoryRangesEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetPhysicalMemoryRangesEx2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetSystemRoutineAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmGetVirtualForPhysical", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsAddressValid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsDriverSuspectForVerifier", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsDriverVerifying", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsDriverVerifyingByAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsFileSectionActive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsIoSpaceActive", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsKernelAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsNonPagedSystemAddressValid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsRecursiveIoFault", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsThisAnNtAsSystem", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsUserAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmIsVerifierEnabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmLockPagableDataSection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmLockPagableSectionByHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapIoSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapIoSpaceEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapLockedPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapLockedPagesSpecifyCache", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapLockedPagesWithReservedMapping", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapMemoryDumpMdlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapUserAddressesToPage", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapVideoDisplay", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapViewInSessionSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapViewInSessionSpaceEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapViewInSystemSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMapViewInSystemSpaceEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMdlPageContentsState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmMdlPagesAreZero", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmPageEntireDriver", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmPrefetchPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProbeAndLockPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProbeAndLockPagesEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProbeAndLockProcessPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProbeAndLockSelectedPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProtectDriverSection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmProtectMdlSystemAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmQuerySystemSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmRemovePhysicalMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmResetDriverPaging", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmRotatePhysicalView", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmSecureVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmSecureVirtualMemoryEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmSetAddressRangeModified", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmSetPermanentCacheAttribute", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmSizeOfMdl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnlockPagableImageSection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnlockPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapIoSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapLockedPages", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapReservedMapping", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapVideoDisplay", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapViewInSessionSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnmapViewInSystemSpace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "MmUnsecureVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAccessCheckAndAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAccessCheckByTypeAndAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAccessCheckByTypeResultListAndAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAccessCheckByTypeResultListAndAuditAlarmByHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAdjustGroupsToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAdjustPrivilegesToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtAllocateVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtClose", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCloseObjectAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCommitComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCommitEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCommitTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCopyFileChunk", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateSection", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateSectionEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtCreateTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtDeleteObjectAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtDeviceIoControlFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtDuplicateToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtEnumerateTransactionObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtFilterToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtFlushBuffersFileEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtFreeVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtFsControlFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtGetNotificationResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtImpersonateAnonymousToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtLockFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtManagePartition", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenJobObjectToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenObjectAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenProcessToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenProcessTokenEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenRegistryTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenThreadToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenThreadTokenEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtOpenTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPowerInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrePrepareComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrePrepareEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrepareComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrepareEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrivilegeCheck", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrivilegeObjectAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPrivilegedServiceAuditAlarm", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPropagationComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtPropagationFailed", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryDirectoryFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryDirectoryFileEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationByName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryInformationTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryQuotaInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQuerySecurityObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtQueryVolumeInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtReadFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtReadOnlyEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRecoverEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRecoverResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRecoverTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRegisterProtocolAddressInformation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRenameTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRollbackComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRollbackEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRollbackRegistryTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRollbackTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtRollforwardTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationEnlistment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationResourceManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationTransaction", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationTransactionManager", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetInformationVirtualMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetQuotaInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetSecurityObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSetVolumeInformationFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtSinglePhaseReject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtUnlockFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "NtWriteFile", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObCloseHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObCloseHandleWithResult", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObDereferenceObjectDeferDelete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObDereferenceObjectDeferDeleteWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObGetFilterVersion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObGetObjectSecurity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObInsertObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObIsKernelHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObMakeTemporaryObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObOpenObjectByPointer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObOpenObjectByPointerWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObQueryNameString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObQueryObjectAuditingByHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectByHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectByHandleWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectByPointer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectByPointerWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectSafe", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReferenceObjectSafeWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObRegisterCallbacks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObReleaseObjectSecurity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObUnRegisterCallbacks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObfDereferenceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObfDereferenceObjectWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObfReferenceObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "ObfReferenceObjectWithTag", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsAcquireSiloHardReference", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsAllocSiloContextSlot", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsAllocateAffinityToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsAssignImpersonationToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsAttachSiloToCurrentThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsChargePoolQuota", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsChargeProcessPoolQuota", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsCreateSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsCreateSystemThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsDereferenceImpersonationToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsDereferencePrimaryToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsDereferenceSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsDetachSiloFromCurrentThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsDisableImpersonation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsFreeAffinityToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsFreeSiloContextSlot", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentServerSiloName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentThreadId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetCurrentThreadTeb", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetEffectiveServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetHostSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetJobServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetJobSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetParentSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetPermanentSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetProcessCreateTimeQuadPart", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetProcessExitStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetProcessExitTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetProcessStartKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetServerSiloActiveConsoleId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetServerSiloServiceSessionId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetSiloContainerId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetSiloMonitorContextSlot", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadCreateTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadExitStatus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadProperty", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetThreadServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsGetVersion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsImpersonateClient", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsInsertPermanentSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsInsertSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsCurrentThreadInServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsCurrentThreadPrefetching", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsDiskCountersEnabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsHostSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsSystemThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsThreadAttachedToSpecificSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsIsThreadTerminating", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsLookupProcessByProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsLookupThreadByThreadId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsMakeSiloContextPermanent", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsQueryProcessAvailableCpus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsQueryProcessAvailableCpusCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsQuerySystemAvailableCpus", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsQuerySystemAvailableCpusCount", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsQueryTotalCycleTimeProcess", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReferenceImpersonationToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReferencePrimaryToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReferenceSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRegisterProcessAvailableCpusChangeNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRegisterSiloMonitor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRegisterSystemAvailableCpusChangeNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReleaseSiloHardReference", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRemoveCreateThreadNotifyRoutine", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRemoveLoadImageNotifyRoutine", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRemoveSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReplaceSiloContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRestoreImpersonation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsReturnPoolQuota", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRevertToSelf", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsRevertToUserMultipleGroupAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCreateProcessNotifyRoutine", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCreateProcessNotifyRoutineEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCreateProcessNotifyRoutineEx2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCreateThreadNotifyRoutine", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCreateThreadNotifyRoutineEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetCurrentThreadPrefetching", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetLoadImageNotifyRoutine", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetLoadImageNotifyRoutineEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsSetSystemMultipleGroupAffinityThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsStartSiloMonitor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsTerminateServerSilo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsTerminateSystemThread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsUnregisterAvailableCpusChangeNotification", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsUnregisterSiloMonitor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsUpdateDiskCounters", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PsWrapApcWow64Thread", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedAllocateMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedFreeMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedIsSystemWheaEnabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedRegisterPlugin", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedSynchronizeExecution", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "PshedUnregisterPlugin", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAbsoluteToSelfRelativeSD", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAddAccessAllowedAce", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAddAccessAllowedAceEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAddAce", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAllocateAndInitializeSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAllocateAndInitializeSidEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAllocateHeap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAnsiStringToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAppendStringToString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAppendUnicodeStringToString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAppendUnicodeToString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAreBitsClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAreBitsSet", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlAssert", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCaptureContext", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCaptureContext2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCaptureStackBackTrace", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCharToInteger", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCheckRegistryKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlClearAllBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlClearBit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlClearBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCmDecodeMemIoResource", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCmEncodeMemIoResource", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareAltitudes", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareMemoryUlong", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompareUnicodeStrings", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompressBuffer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCompressChunks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlContractHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlConvertSidToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyBitMap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyDeviceMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyLuid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyMemoryNonTemporal", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopySid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCopyVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCrc32", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCrc64", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateAcl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateHashTableEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateHeap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateRegistryKey", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateSecurityDescriptorRelative", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateServiceSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateSystemVolumeInformationFolder", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCreateVirtualAccountSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlCustomCPToUnicodeN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressBuffer", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressBufferEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressBufferEx2", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressChunks", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressFragment", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDecompressFragmentEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDelete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteAce", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteElementGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteElementGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteElementGenericTableAvlEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteNoSplay", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDeleteRegistryValue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDescribeChunk", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDestroyHeap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDowncaseUnicodeChar", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDowncaseUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDrainNonVolatileFlush", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlDuplicateUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEndEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEndStrongEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEndWeakEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateGenericTableLikeADirectory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateGenericTableWithoutSplaying", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEnumerateGenericTableWithoutSplayingAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEqualPrefixSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEqualSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEqualString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlEqualUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlExpandHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlExtendCorrelationVector", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlExtractBitMap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFillMemoryNonTemporal", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFillNonVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindClearBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindClearBitsAndSet", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindClearRuns", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindClosestEncodableLength", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindFirstRunClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindLastBackwardRunClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindLeastSignificantBit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindLongestRunClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindMostSignificantBit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindNextForwardRunClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindSetBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindSetBitsAndClear", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFindUnicodePrefix", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFlushNonVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFlushNonVolatileMemoryRanges", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeAnsiString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeHeap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeNonVolatileToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeOemString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeUTF8String", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlFreeUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGUIDFromString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGenerate8dot3Name", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGenerateClass5Guid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetAce", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetAcesBufferSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetActiveConsoleId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetCallersAddress", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetCompressionWorkSpaceSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetConsoleSessionForegroundProcessId", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetDaclSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetElementGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetElementGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetEnabledExtendedFeatures", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetGroupSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetNextEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetNonVolatileToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetNtProductType", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetNtSystemRoot", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetOwnerSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetPersistedStateLocation", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetProductInfo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetSaclSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetSuiteMask", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetSystemGlobalData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlGetVersion", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlHashUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIdentifierAuthoritySid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIdnToAscii", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIdnToNameprepUnicode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIdnToUnicode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIncrementCorrelationVector", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitAnsiString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitAnsiStringEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitCodePageTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitStringEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitStrongEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitUTF8String", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitUTF8StringEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitUnicodeStringEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitWeakEnumerationHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeBitMap", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeCorrelationVector", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeSidEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInitializeUnicodePrefix", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertElementGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertElementGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertElementGenericTableFull", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertElementGenericTableFullAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInsertUnicodePrefix", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlInt64ToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIntegerToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIoDecodeMemIoResource", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIoEncodeMemIoResource", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsApiSetImplemented", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsCloudFilesPlaceholder", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsGenericTableEmpty", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsGenericTableEmptyAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsMultiSessionSku", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsMultiUsersInSessionSku", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsNameLegalDOS8Dot3", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsNonEmptyDirectoryReparsePointAllowed", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsNormalizedString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsNtDdiVersionAvailable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsPartialPlaceholder", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsPartialPlaceholderFileHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsPartialPlaceholderFileInfo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsSandboxedToken", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsSandboxedTokenHandle", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsServicePackVersionInstalled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsStateSeparationEnabled", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsUntrustedObject", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsValidOemCharacter", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlIsZeroMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLengthRequiredSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLengthSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLengthSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupElementGenericTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupElementGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupElementGenericTableFull", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupElementGenericTableFullAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlLookupFirstMatchingElementGenericTableAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlMapGenericMask", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlMoveVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlMultiByteToUnicodeN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlMultiByteToUnicodeSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNextUnicodePrefix", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNormalizeSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNormalizeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNtStatusToDosError", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNtStatusToDosErrorNoTeb", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberGenericTableElements", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberGenericTableElementsAvl", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberOfClearBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberOfClearBitsInRange", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberOfSetBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberOfSetBitsInRange", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlNumberOfSetBitsUlongPtr", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlOemStringToCountedUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlOemStringToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlOemToUnicodeN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlOsDeploymentState", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlPrefetchMemoryNonTemporal", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlPrefixString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlPrefixUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryPackageIdentity", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryPackageIdentityEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryProcessPlaceholderCompatibilityMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryRegistryValueWithFallback", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryRegistryValues", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryThreadPlaceholderCompatibilityMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlQueryValidationRunlevel", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRaiseCustomSystemEventTrigger", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRandom", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRandomEx", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRealPredecessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRealSuccessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRemoveEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRemoveUnicodePrefix", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlReplaceSidInSd", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlReserveChunk", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRunOnceBeginInitialize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRunOnceComplete", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRunOnceExecuteOnce", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlRunOnceInitialize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSecondsSince1970ToTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSecondsSince1980ToTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSelfRelativeToAbsoluteSD", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetAllBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetBit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetBits", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetDaclSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetGroupSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetOwnerSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetProcessPlaceholderCompatibilityMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetSystemGlobalData", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetThreadPlaceholderCompatibilityMode", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSetVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSplay", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlStringFromGUID", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlStronglyEnumerateEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSubAuthorityCountSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSubAuthoritySid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSubtreePredecessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSubtreeSuccessor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlSuffixUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlTestBit", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlTimeFieldsToTime", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlTimeToSecondsSince1970", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlTimeToSecondsSince1980", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlTimeToTimeFields", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUTF8StringToUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUTF8ToUnicodeN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToAnsiString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToCountedOemString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToInt64", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToInteger", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToOemString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeStringToUTF8String", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeToCustomCPN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeToMultiByteN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeToMultiByteSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeToOemN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUnicodeToUTF8N", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeChar", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeStringToCountedOemString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeStringToOemString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeToCustomCPN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeToMultiByteN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpcaseUnicodeToOemN", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpperChar", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlUpperString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlValidRelativeSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlValidSecurityDescriptor", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlValidSid", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlValidateCorrelationVector", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlValidateUnicodeString", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlVerifyVersionInfo", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlVolumeDeviceToDosName", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlWalkFrameChain", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlWeaklyEnumerateEntryHashTable", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlWriteNonVolatileMemory", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlWriteRegistryValue", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlxAnsiStringToUnicodeSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlxOemStringToUnicodeSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlxUnicodeStringToAnsiSize", address: 0, trampoline: 0, enabled: false },
    InlineHookDb { name: "RtlxUnicodeStringToOemSize", address: 0, trampoline: 0, enabled: false },
];

pub fn install_inline_hook_by_name(name: &str, hook_handler: u64) -> Result<u64, HookError> {
    match name {
        "DbgBreakPointWithStatus" => {
            let addr = DbgBreakPointWithStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgPrint" => {
            let addr = DbgPrint as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgPrintEx" => {
            let addr = DbgPrintEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgPrintReturnControlC" => {
            let addr = DbgPrintReturnControlC as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgPrompt" => {
            let addr = DbgPrompt as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgQueryDebugFilterState" => {
            let addr = DbgQueryDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgSetDebugFilterState" => {
            let addr = DbgSetDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "DbgSetDebugPrintCallback" => {
            let addr = DbgSetDebugPrintCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireFastMutex" => {
            let addr = ExAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireFastMutexUnsafe" => {
            let addr = ExAcquireFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquirePushLockExclusiveEx" => {
            let addr = ExAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquirePushLockSharedEx" => {
            let addr = ExAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireResourceExclusiveLite" => {
            let addr = ExAcquireResourceExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireResourceSharedLite" => {
            let addr = ExAcquireResourceSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireRundownProtection" => {
            let addr = ExAcquireRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireRundownProtectionCacheAware" => {
            let addr = ExAcquireRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireRundownProtectionCacheAwareEx" => {
            let addr = ExAcquireRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireRundownProtectionEx" => {
            let addr = ExAcquireRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSharedStarveExclusive" => {
            let addr = ExAcquireSharedStarveExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSharedWaitForExclusive" => {
            let addr = ExAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSpinLockExclusive" => {
            let addr = ExAcquireSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSpinLockShared" => {
            let addr = ExAcquireSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAdjustLookasideDepth" => {
            let addr = ExAdjustLookasideDepth as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocateCacheAwareRundownProtection" => {
            let addr = ExAllocateCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocateFromLookasideListEx" => {
            let addr = ExAllocateFromLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocateFromNPagedLookasideList" => {
            let addr = ExAllocateFromNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocateFromPagedLookasideList" => {
            let addr = ExAllocateFromPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocatePool2" => {
            let addr = ExAllocatePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocatePool3" => {
            let addr = ExAllocatePool3 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocatePoolWithQuota" => {
            let addr = ExAllocatePoolWithQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExAllocateTimer" => {
            let addr = ExAllocateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExCancelTimer" => {
            let addr = ExCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExCleanupRundownProtectionCacheAware" => {
            let addr = ExCleanupRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExConvertExclusiveToSharedLite" => {
            let addr = ExConvertExclusiveToSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExCreateCallback" => {
            let addr = ExCreateCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExCreatePool" => {
            let addr = ExCreatePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDeleteLookasideListEx" => {
            let addr = ExDeleteLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDeleteNPagedLookasideList" => {
            let addr = ExDeleteNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDeletePagedLookasideList" => {
            let addr = ExDeletePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDeleteResourceLite" => {
            let addr = ExDeleteResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDeleteTimer" => {
            let addr = ExDeleteTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDestroyPool" => {
            let addr = ExDestroyPool as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExDisableResourceBoostLite" => {
            let addr = ExDisableResourceBoostLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExEnterCriticalRegionAndAcquireResourceExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExEnterCriticalRegionAndAcquireResourceShared" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExEnumerateSystemFirmwareTables" => {
            let addr = ExEnumerateSystemFirmwareTables as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExExtendZone" => {
            let addr = ExExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFlushLookasideListEx" => {
            let addr = ExFlushLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreeCacheAwareRundownProtection" => {
            let addr = ExFreeCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreePool" => {
            let addr = ExFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreePool2" => {
            let addr = ExFreePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreePoolWithTag" => {
            let addr = ExFreePoolWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreeToLookasideListEx" => {
            let addr = ExFreeToLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreeToNPagedLookasideList" => {
            let addr = ExFreeToNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExFreeToPagedLookasideList" => {
            let addr = ExFreeToPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetExclusiveWaiterCount" => {
            let addr = ExGetExclusiveWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetFirmwareEnvironmentVariable" => {
            let addr = ExGetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetFirmwareType" => {
            let addr = ExGetFirmwareType as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetPreviousMode" => {
            let addr = ExGetPreviousMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetSharedWaiterCount" => {
            let addr = ExGetSharedWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExGetSystemFirmwareTable" => {
            let addr = ExGetSystemFirmwareTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeDeviceAts" => {
            let addr = ExInitializeDeviceAts as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeLookasideListEx" => {
            let addr = ExInitializeLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeNPagedLookasideList" => {
            let addr = ExInitializeNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializePagedLookasideList" => {
            let addr = ExInitializePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializePushLock" => {
            let addr = ExInitializePushLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeResourceLite" => {
            let addr = ExInitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeRundownProtection" => {
            let addr = ExInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeRundownProtectionCacheAware" => {
            let addr = ExInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeRundownProtectionCacheAwareEx" => {
            let addr = ExInitializeRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInitializeZone" => {
            let addr = ExInitializeZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedAddLargeInteger" => {
            let addr = ExInterlockedAddLargeInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedAddUlong" => {
            let addr = ExInterlockedAddUlong as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedExtendZone" => {
            let addr = ExInterlockedExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedInsertHeadList" => {
            let addr = ExInterlockedInsertHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedInsertTailList" => {
            let addr = ExInterlockedInsertTailList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedPopEntryList" => {
            let addr = ExInterlockedPopEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedPushEntryList" => {
            let addr = ExInterlockedPushEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExInterlockedRemoveHeadList" => {
            let addr = ExInterlockedRemoveHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExIsManufacturingModeEnabled" => {
            let addr = ExIsManufacturingModeEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExIsProcessorFeaturePresent" => {
            let addr = ExIsProcessorFeaturePresent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExIsResourceAcquiredExclusiveLite" => {
            let addr = ExIsResourceAcquiredExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExIsResourceAcquiredSharedLite" => {
            let addr = ExIsResourceAcquiredSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExIsSoftBoot" => {
            let addr = ExIsSoftBoot as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExLocalTimeToSystemTime" => {
            let addr = ExLocalTimeToSystemTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExNotifyCallback" => {
            let addr = ExNotifyCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExQueryDepthSList" => {
            let addr = ExQueryDepthSList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExQueryPoolBlockSize" => {
            let addr = ExQueryPoolBlockSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExQueryTimerResolution" => {
            let addr = ExQueryTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExQueueWorkItem" => {
            let addr = ExQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRaiseAccessViolation" => {
            let addr = ExRaiseAccessViolation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRaiseDatatypeMisalignment" => {
            let addr = ExRaiseDatatypeMisalignment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRaiseStatus" => {
            let addr = ExRaiseStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRcuFreePool" => {
            let addr = ExRcuFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReInitializeRundownProtection" => {
            let addr = ExReInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReInitializeRundownProtectionCacheAware" => {
            let addr = ExReInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRegisterCallback" => {
            let addr = ExRegisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReinitializeResourceLite" => {
            let addr = ExReinitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseFastMutex" => {
            let addr = ExReleaseFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseFastMutexUnsafe" => {
            let addr = ExReleaseFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleasePushLockExclusiveEx" => {
            let addr = ExReleasePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleasePushLockSharedEx" => {
            let addr = ExReleasePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseResourceAndLeaveCriticalRegion" => {
            let addr = ExReleaseResourceAndLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseResourceForThreadLite" => {
            let addr = ExReleaseResourceForThreadLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseResourceLite" => {
            let addr = ExReleaseResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseRundownProtection" => {
            let addr = ExReleaseRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseRundownProtectionCacheAware" => {
            let addr = ExReleaseRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseRundownProtectionCacheAwareEx" => {
            let addr = ExReleaseRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseRundownProtectionEx" => {
            let addr = ExReleaseRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseSpinLockExclusive" => {
            let addr = ExReleaseSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseSpinLockExclusiveFromDpcLevel" => {
            let addr = ExReleaseSpinLockExclusiveFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseSpinLockShared" => {
            let addr = ExReleaseSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExReleaseSpinLockSharedFromDpcLevel" => {
            let addr = ExReleaseSpinLockSharedFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRundownCompleted" => {
            let addr = ExRundownCompleted as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExRundownCompletedCacheAware" => {
            let addr = ExRundownCompletedCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSecurePoolUpdate" => {
            let addr = ExSecurePoolUpdate as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSecurePoolValidate" => {
            let addr = ExSecurePoolValidate as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSetFirmwareEnvironmentVariable" => {
            let addr = ExSetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSetResourceOwnerPointer" => {
            let addr = ExSetResourceOwnerPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSetResourceOwnerPointerEx" => {
            let addr = ExSetResourceOwnerPointerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSetTimer" => {
            let addr = ExSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSetTimerResolution" => {
            let addr = ExSetTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSizeOfRundownProtectionCacheAware" => {
            let addr = ExSizeOfRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExSystemTimeToLocalTime" => {
            let addr = ExSystemTimeToLocalTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryAcquirePushLockExclusiveEx" => {
            let addr = ExTryAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryAcquirePushLockSharedEx" => {
            let addr = ExTryAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryConvertSharedSpinLockExclusive" => {
            let addr = ExTryConvertSharedSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExTryToAcquireFastMutex" => {
            let addr = ExTryToAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExUnregisterCallback" => {
            let addr = ExUnregisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExUuidCreate" => {
            let addr = ExUuidCreate as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExVerifySuite" => {
            let addr = ExVerifySuite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExWaitForRundownProtectionRelease" => {
            let addr = ExWaitForRundownProtectionRelease as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExWaitForRundownProtectionReleaseCacheAware" => {
            let addr = ExWaitForRundownProtectionReleaseCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExpInterlockedFlushSList" => {
            let addr = ExpInterlockedFlushSList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExpInterlockedPopEntrySList" => {
            let addr = ExpInterlockedPopEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExpInterlockedPushEntrySList" => {
            let addr = ExpInterlockedPushEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ExportSecurityContext" => {
            let addr = ExportSecurityContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAcquireCancelSpinLock" => {
            let addr = IoAcquireCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAcquireKsrPersistentMemory" => {
            let addr = IoAcquireKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAcquireKsrPersistentMemoryEx" => {
            let addr = IoAcquireKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAcquireRemoveLockEx" => {
            let addr = IoAcquireRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAcquireVpbSpinLock" => {
            let addr = IoAcquireVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateController" => {
            let addr = IoAllocateController as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateDriverObjectExtension" => {
            let addr = IoAllocateDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateErrorLogEntry" => {
            let addr = IoAllocateErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateIrp" => {
            let addr = IoAllocateIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateIrpEx" => {
            let addr = IoAllocateIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateMdl" => {
            let addr = IoAllocateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateMiniCompletionPacket" => {
            let addr = IoAllocateMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateSfioStreamIdentifier" => {
            let addr = IoAllocateSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAllocateWorkItem" => {
            let addr = IoAllocateWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoApplyPriorityInfoThread" => {
            let addr = IoApplyPriorityInfoThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAssignResources" => {
            let addr = IoAssignResources as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAttachDevice" => {
            let addr = IoAttachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAttachDeviceByPointer" => {
            let addr = IoAttachDeviceByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAttachDeviceToDeviceStack" => {
            let addr = IoAttachDeviceToDeviceStack as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoAttachDeviceToDeviceStackSafe" => {
            let addr = IoAttachDeviceToDeviceStackSafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoBuildAsynchronousFsdRequest" => {
            let addr = IoBuildAsynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoBuildDeviceIoControlRequest" => {
            let addr = IoBuildDeviceIoControlRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoBuildPartialMdl" => {
            let addr = IoBuildPartialMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoBuildSynchronousFsdRequest" => {
            let addr = IoBuildSynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCancelFileOpen" => {
            let addr = IoCancelFileOpen as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCancelIrp" => {
            let addr = IoCancelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckDesiredAccess" => {
            let addr = IoCheckDesiredAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckEaBufferValidity" => {
            let addr = IoCheckEaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckFileObjectOpenedAsCopyDestination" => {
            let addr = IoCheckFileObjectOpenedAsCopyDestination as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckFileObjectOpenedAsCopySource" => {
            let addr = IoCheckFileObjectOpenedAsCopySource as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckFunctionAccess" => {
            let addr = IoCheckFunctionAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckLinkShareAccess" => {
            let addr = IoCheckLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckQuerySetFileInformation" => {
            let addr = IoCheckQuerySetFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckQuerySetVolumeInformation" => {
            let addr = IoCheckQuerySetVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckQuotaBufferValidity" => {
            let addr = IoCheckQuotaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckShareAccess" => {
            let addr = IoCheckShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCheckShareAccessEx" => {
            let addr = IoCheckShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCleanupIrp" => {
            let addr = IoCleanupIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoClearActivityIdThread" => {
            let addr = IoClearActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoClearFsTrackOffsetState" => {
            let addr = IoClearFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoClearIrpExtraCreateParameter" => {
            let addr = IoClearIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoConnectInterrupt" => {
            let addr = IoConnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoConnectInterruptEx" => {
            let addr = IoConnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateController" => {
            let addr = IoCreateController as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateDevice" => {
            let addr = IoCreateDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateDisk" => {
            let addr = IoCreateDisk as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateDriverProxyExtension" => {
            let addr = IoCreateDriverProxyExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateFile" => {
            let addr = IoCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateFileEx" => {
            let addr = IoCreateFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateFileSpecifyDeviceObjectHint" => {
            let addr = IoCreateFileSpecifyDeviceObjectHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateNotificationEvent" => {
            let addr = IoCreateNotificationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateStreamFileObject" => {
            let addr = IoCreateStreamFileObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateStreamFileObjectEx" => {
            let addr = IoCreateStreamFileObjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateStreamFileObjectEx2" => {
            let addr = IoCreateStreamFileObjectEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateStreamFileObjectLite" => {
            let addr = IoCreateStreamFileObjectLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateSymbolicLink" => {
            let addr = IoCreateSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateSynchronizationEvent" => {
            let addr = IoCreateSynchronizationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateSystemThread" => {
            let addr = IoCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCreateUnprotectedSymbolicLink" => {
            let addr = IoCreateUnprotectedSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqInitialize" => {
            let addr = IoCsqInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqInitializeEx" => {
            let addr = IoCsqInitializeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqInsertIrp" => {
            let addr = IoCsqInsertIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqInsertIrpEx" => {
            let addr = IoCsqInsertIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqRemoveIrp" => {
            let addr = IoCsqRemoveIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoCsqRemoveNextIrp" => {
            let addr = IoCsqRemoveNextIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDecrementKeepAliveCount" => {
            let addr = IoDecrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDeleteController" => {
            let addr = IoDeleteController as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDeleteDevice" => {
            let addr = IoDeleteDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDeleteSymbolicLink" => {
            let addr = IoDeleteSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDetachDevice" => {
            let addr = IoDetachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDisconnectInterrupt" => {
            let addr = IoDisconnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDisconnectInterruptEx" => {
            let addr = IoDisconnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoDriverProxyCreateHotSwappableWorkerThread" => {
            let addr = IoDriverProxyCreateHotSwappableWorkerThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoEnumerateDeviceObjectList" => {
            let addr = IoEnumerateDeviceObjectList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoEnumerateKsrPersistentMemoryEx" => {
            let addr = IoEnumerateKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoEnumerateRegisteredFiltersList" => {
            let addr = IoEnumerateRegisteredFiltersList as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFastQueryNetworkAttributes" => {
            let addr = IoFastQueryNetworkAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoForwardIrpSynchronously" => {
            let addr = IoForwardIrpSynchronously as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeController" => {
            let addr = IoFreeController as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeErrorLogEntry" => {
            let addr = IoFreeErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeIrp" => {
            let addr = IoFreeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeKsrPersistentMemory" => {
            let addr = IoFreeKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeMdl" => {
            let addr = IoFreeMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeMiniCompletionPacket" => {
            let addr = IoFreeMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeSfioStreamIdentifier" => {
            let addr = IoFreeSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoFreeWorkItem" => {
            let addr = IoFreeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetActivityIdIrp" => {
            let addr = IoGetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetActivityIdThread" => {
            let addr = IoGetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetAffinityInterrupt" => {
            let addr = IoGetAffinityInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetAttachedDevice" => {
            let addr = IoGetAttachedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetAttachedDeviceReference" => {
            let addr = IoGetAttachedDeviceReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetBaseFileSystemDeviceObject" => {
            let addr = IoGetBaseFileSystemDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetBootDiskInformation" => {
            let addr = IoGetBootDiskInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetBootDiskInformationLite" => {
            let addr = IoGetBootDiskInformationLite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetConfigurationInformation" => {
            let addr = IoGetConfigurationInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetContainerInformation" => {
            let addr = IoGetContainerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetCopyInformationExtension" => {
            let addr = IoGetCopyInformationExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetCurrentProcess" => {
            let addr = IoGetCurrentProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceAttachmentBaseRef" => {
            let addr = IoGetDeviceAttachmentBaseRef as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceDirectory" => {
            let addr = IoGetDeviceDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceInterfaceAlias" => {
            let addr = IoGetDeviceInterfaceAlias as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceInterfacePropertyData" => {
            let addr = IoGetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceInterfaces" => {
            let addr = IoGetDeviceInterfaces as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceNumaNode" => {
            let addr = IoGetDeviceNumaNode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceObjectPointer" => {
            let addr = IoGetDeviceObjectPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceProperty" => {
            let addr = IoGetDeviceProperty as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDevicePropertyData" => {
            let addr = IoGetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDeviceToVerify" => {
            let addr = IoGetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDiskDeviceObject" => {
            let addr = IoGetDiskDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDmaAdapter" => {
            let addr = IoGetDmaAdapter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverDirectory" => {
            let addr = IoGetDriverDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverObjectExtension" => {
            let addr = IoGetDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverProxyEndpointWrapper" => {
            let addr = IoGetDriverProxyEndpointWrapper as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverProxyExtensionFromDriverObject" => {
            let addr = IoGetDriverProxyExtensionFromDriverObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverProxyExtensionVersion" => {
            let addr = IoGetDriverProxyExtensionVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetDriverProxyFeatures" => {
            let addr = IoGetDriverProxyFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetFileObjectGenericMapping" => {
            let addr = IoGetFileObjectGenericMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetFsTrackOffsetState" => {
            let addr = IoGetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetFsZeroingOffset" => {
            let addr = IoGetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetInitialStack" => {
            let addr = IoGetInitialStack as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetInitiatorProcess" => {
            let addr = IoGetInitiatorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetIoAttributionHandle" => {
            let addr = IoGetIoAttributionHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetIoPriorityHint" => {
            let addr = IoGetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetIommuInterface" => {
            let addr = IoGetIommuInterface as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetIommuInterfaceEx" => {
            let addr = IoGetIommuInterfaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetIrpExtraCreateParameter" => {
            let addr = IoGetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetKsrPersistentMemoryBuffer" => {
            let addr = IoGetKsrPersistentMemoryBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetLowerDeviceObject" => {
            let addr = IoGetLowerDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetOplockKeyContext" => {
            let addr = IoGetOplockKeyContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetOplockKeyContextEx" => {
            let addr = IoGetOplockKeyContextEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetPagingIoPriority" => {
            let addr = IoGetPagingIoPriority as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetRelatedDeviceObject" => {
            let addr = IoGetRelatedDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetRequestorProcess" => {
            let addr = IoGetRequestorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetRequestorProcessId" => {
            let addr = IoGetRequestorProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetRequestorSessionId" => {
            let addr = IoGetRequestorSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetSfioStreamIdentifier" => {
            let addr = IoGetSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetShadowFileInformation" => {
            let addr = IoGetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetSilo" => {
            let addr = IoGetSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetSiloParameters" => {
            let addr = IoGetSiloParameters as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetStackLimits" => {
            let addr = IoGetStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetTopLevelIrp" => {
            let addr = IoGetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoGetTransactionParameterBlock" => {
            let addr = IoGetTransactionParameterBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoHotSwapDriverProxyEndpoints" => {
            let addr = IoHotSwapDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIncrementKeepAliveCount" => {
            let addr = IoIncrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInitializeIrp" => {
            let addr = IoInitializeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInitializeIrpEx" => {
            let addr = IoInitializeIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInitializeRemoveLockEx" => {
            let addr = IoInitializeRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInitializeTimer" => {
            let addr = IoInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInitializeWorkItem" => {
            let addr = IoInitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInvalidateDeviceRelations" => {
            let addr = IoInvalidateDeviceRelations as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoInvalidateDeviceState" => {
            let addr = IoInvalidateDeviceState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIrpHasFsTrackOffsetExtensionType" => {
            let addr = IoIrpHasFsTrackOffsetExtensionType as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIs32bitProcess" => {
            let addr = IoIs32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsFileObjectIgnoringSharing" => {
            let addr = IoIsFileObjectIgnoringSharing as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsFileOriginRemote" => {
            let addr = IoIsFileOriginRemote as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsInitiator32bitProcess" => {
            let addr = IoIsInitiator32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsOperationSynchronous" => {
            let addr = IoIsOperationSynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsSystemThread" => {
            let addr = IoIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsValidIrpStatus" => {
            let addr = IoIsValidIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsValidNameGraftingBuffer" => {
            let addr = IoIsValidNameGraftingBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoIsWdmVersionAvailable" => {
            let addr = IoIsWdmVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoMakeAssociatedIrp" => {
            let addr = IoMakeAssociatedIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoMakeAssociatedIrpEx" => {
            let addr = IoMakeAssociatedIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoMapKsrPersistentMemoryEx" => {
            let addr = IoMapKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoOpenDeviceInterfaceRegistryKey" => {
            let addr = IoOpenDeviceInterfaceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoOpenDeviceRegistryKey" => {
            let addr = IoOpenDeviceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoOpenDriverRegistryKey" => {
            let addr = IoOpenDriverRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoPageRead" => {
            let addr = IoPageRead as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoPropagateActivityIdToThread" => {
            let addr = IoPropagateActivityIdToThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryDeviceDescription" => {
            let addr = IoQueryDeviceDescription as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryDmaFeatureSupport" => {
            let addr = IoQueryDmaFeatureSupport as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryFileDosDeviceName" => {
            let addr = IoQueryFileDosDeviceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryFileInformation" => {
            let addr = IoQueryFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryFullDriverPath" => {
            let addr = IoQueryFullDriverPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryInformationByName" => {
            let addr = IoQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryKsrPersistentMemorySize" => {
            let addr = IoQueryKsrPersistentMemorySize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryKsrPersistentMemorySizeEx" => {
            let addr = IoQueryKsrPersistentMemorySizeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueryVolumeInformation" => {
            let addr = IoQueryVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueueThreadIrp" => {
            let addr = IoQueueThreadIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueueWorkItem" => {
            let addr = IoQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoQueueWorkItemEx" => {
            let addr = IoQueueWorkItemEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRaiseHardError" => {
            let addr = IoRaiseHardError as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRaiseInformationalHardError" => {
            let addr = IoRaiseInformationalHardError as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReadDiskSignature" => {
            let addr = IoReadDiskSignature as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReadPartitionTable" => {
            let addr = IoReadPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReadPartitionTableEx" => {
            let addr = IoReadPartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRecordIoAttribution" => {
            let addr = IoRecordIoAttribution as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterBootDriverCallback" => {
            let addr = IoRegisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterBootDriverReinitialization" => {
            let addr = IoRegisterBootDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterContainerNotification" => {
            let addr = IoRegisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterDeviceInterface" => {
            let addr = IoRegisterDeviceInterface as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterDriverProxyEndpoints" => {
            let addr = IoRegisterDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterDriverReinitialization" => {
            let addr = IoRegisterDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterFileSystem" => {
            let addr = IoRegisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterFsRegistrationChange" => {
            let addr = IoRegisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterFsRegistrationChangeMountAware" => {
            let addr = IoRegisterFsRegistrationChangeMountAware as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterLastChanceShutdownNotification" => {
            let addr = IoRegisterLastChanceShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterPlugPlayNotification" => {
            let addr = IoRegisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRegisterShutdownNotification" => {
            let addr = IoRegisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReleaseCancelSpinLock" => {
            let addr = IoReleaseCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReleaseRemoveLockAndWaitEx" => {
            let addr = IoReleaseRemoveLockAndWaitEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReleaseRemoveLockEx" => {
            let addr = IoReleaseRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReleaseVpbSpinLock" => {
            let addr = IoReleaseVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRemoveIoCompletion" => {
            let addr = IoRemoveIoCompletion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRemoveLinkShareAccess" => {
            let addr = IoRemoveLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRemoveLinkShareAccessEx" => {
            let addr = IoRemoveLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRemoveShareAccess" => {
            let addr = IoRemoveShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReplaceFileObjectName" => {
            let addr = IoReplaceFileObjectName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReplacePartitionUnit" => {
            let addr = IoReplacePartitionUnit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportDetectedDevice" => {
            let addr = IoReportDetectedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportInterruptActive" => {
            let addr = IoReportInterruptActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportInterruptInactive" => {
            let addr = IoReportInterruptInactive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportResourceForDetection" => {
            let addr = IoReportResourceForDetection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportResourceUsage" => {
            let addr = IoReportResourceUsage as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportRootDevice" => {
            let addr = IoReportRootDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportTargetDeviceChange" => {
            let addr = IoReportTargetDeviceChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReportTargetDeviceChangeAsynchronous" => {
            let addr = IoReportTargetDeviceChangeAsynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRequestDeviceEject" => {
            let addr = IoRequestDeviceEject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRequestDeviceEjectEx" => {
            let addr = IoRequestDeviceEjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRequestDeviceRemovalForReset" => {
            let addr = IoRequestDeviceRemovalForReset as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReserveKsrPersistentMemory" => {
            let addr = IoReserveKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReserveKsrPersistentMemoryEx" => {
            let addr = IoReserveKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoRetrievePriorityInfo" => {
            let addr = IoRetrievePriorityInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoReuseIrp" => {
            let addr = IoReuseIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetActivityIdIrp" => {
            let addr = IoSetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetActivityIdThread" => {
            let addr = IoSetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetCompletionRoutineEx" => {
            let addr = IoSetCompletionRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetDeviceInterfacePropertyData" => {
            let addr = IoSetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetDeviceInterfaceState" => {
            let addr = IoSetDeviceInterfaceState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetDevicePropertyData" => {
            let addr = IoSetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetDeviceToVerify" => {
            let addr = IoSetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetFileObjectIgnoreSharing" => {
            let addr = IoSetFileObjectIgnoreSharing as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetFileOrigin" => {
            let addr = IoSetFileOrigin as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetFsTrackOffsetState" => {
            let addr = IoSetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetFsZeroingOffset" => {
            let addr = IoSetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetFsZeroingOffsetRequired" => {
            let addr = IoSetFsZeroingOffsetRequired as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetHardErrorOrVerifyDevice" => {
            let addr = IoSetHardErrorOrVerifyDevice as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetInformation" => {
            let addr = IoSetInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetIoAttributionIrp" => {
            let addr = IoSetIoAttributionIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetIoCompletionEx" => {
            let addr = IoSetIoCompletionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetIoPriorityHint" => {
            let addr = IoSetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetIrpExtraCreateParameter" => {
            let addr = IoSetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetLinkShareAccess" => {
            let addr = IoSetLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetMasterIrpStatus" => {
            let addr = IoSetMasterIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetPartitionInformation" => {
            let addr = IoSetPartitionInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetPartitionInformationEx" => {
            let addr = IoSetPartitionInformationEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetShadowFileInformation" => {
            let addr = IoSetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetShareAccess" => {
            let addr = IoSetShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetShareAccessEx" => {
            let addr = IoSetShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetStartIoAttributes" => {
            let addr = IoSetStartIoAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetSystemPartition" => {
            let addr = IoSetSystemPartition as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetThreadHardErrorMode" => {
            let addr = IoSetThreadHardErrorMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSetTopLevelIrp" => {
            let addr = IoSetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSizeOfIrpEx" => {
            let addr = IoSizeOfIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSizeofWorkItem" => {
            let addr = IoSizeofWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoStartNextPacket" => {
            let addr = IoStartNextPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoStartNextPacketByKey" => {
            let addr = IoStartNextPacketByKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoStartPacket" => {
            let addr = IoStartPacket as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoStartTimer" => {
            let addr = IoStartTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoStopTimer" => {
            let addr = IoStopTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSynchronousCallDriver" => {
            let addr = IoSynchronousCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoSynchronousPageWrite" => {
            let addr = IoSynchronousPageWrite as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoThreadToProcess" => {
            let addr = IoThreadToProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoTransferActivityId" => {
            let addr = IoTransferActivityId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoTranslateBusAddress" => {
            let addr = IoTranslateBusAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoTryQueueWorkItem" => {
            let addr = IoTryQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUninitializeWorkItem" => {
            let addr = IoUninitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterBootDriverCallback" => {
            let addr = IoUnregisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterContainerNotification" => {
            let addr = IoUnregisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterFileSystem" => {
            let addr = IoUnregisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterFsRegistrationChange" => {
            let addr = IoUnregisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterPlugPlayNotification" => {
            let addr = IoUnregisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterPlugPlayNotificationEx" => {
            let addr = IoUnregisterPlugPlayNotificationEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUnregisterShutdownNotification" => {
            let addr = IoUnregisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUpdateLinkShareAccess" => {
            let addr = IoUpdateLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUpdateLinkShareAccessEx" => {
            let addr = IoUpdateLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoUpdateShareAccess" => {
            let addr = IoUpdateShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoValidateDeviceIoControlAccess" => {
            let addr = IoValidateDeviceIoControlAccess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVerifyPartitionTable" => {
            let addr = IoVerifyPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVerifyVolume" => {
            let addr = IoVerifyVolume as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVolumeDeviceNameToGuid" => {
            let addr = IoVolumeDeviceNameToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVolumeDeviceNameToGuidPath" => {
            let addr = IoVolumeDeviceNameToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVolumeDeviceToDosName" => {
            let addr = IoVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVolumeDeviceToGuid" => {
            let addr = IoVolumeDeviceToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoVolumeDeviceToGuidPath" => {
            let addr = IoVolumeDeviceToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIAllocateInstanceIds" => {
            let addr = IoWMIAllocateInstanceIds as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIDeviceObjectToInstanceName" => {
            let addr = IoWMIDeviceObjectToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIDeviceObjectToProviderId" => {
            let addr = IoWMIDeviceObjectToProviderId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIExecuteMethod" => {
            let addr = IoWMIExecuteMethod as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIHandleToInstanceName" => {
            let addr = IoWMIHandleToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIOpenBlock" => {
            let addr = IoWMIOpenBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIQueryAllData" => {
            let addr = IoWMIQueryAllData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIQueryAllDataMultiple" => {
            let addr = IoWMIQueryAllDataMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIQuerySingleInstance" => {
            let addr = IoWMIQuerySingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIQuerySingleInstanceMultiple" => {
            let addr = IoWMIQuerySingleInstanceMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIRegistrationControl" => {
            let addr = IoWMIRegistrationControl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMISetNotificationCallback" => {
            let addr = IoWMISetNotificationCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMISetSingleInstance" => {
            let addr = IoWMISetSingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMISetSingleItem" => {
            let addr = IoWMISetSingleItem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMISuggestInstanceName" => {
            let addr = IoWMISuggestInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWMIWriteEvent" => {
            let addr = IoWMIWriteEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWithinStackLimits" => {
            let addr = IoWithinStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWriteErrorLogEntry" => {
            let addr = IoWriteErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWriteKsrPersistentMemory" => {
            let addr = IoWriteKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWritePartitionTable" => {
            let addr = IoWritePartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IoWritePartitionTableEx" => {
            let addr = IoWritePartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IofCallDriver" => {
            let addr = IofCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IofCompleteRequest" => {
            let addr = IofCompleteRequest as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "IofGetDriverProxyWrapperFromEndpoint" => {
            let addr = IofGetDriverProxyWrapperFromEndpoint as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireGuardedMutex" => {
            let addr = KeAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireGuardedMutexUnsafe" => {
            let addr = KeAcquireGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireInStackQueuedSpinLock" => {
            let addr = KeAcquireInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireInStackQueuedSpinLockAtDpcLevel" => {
            let addr = KeAcquireInStackQueuedSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireInStackQueuedSpinLockForDpc" => {
            let addr = KeAcquireInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireInterruptSpinLock" => {
            let addr = KeAcquireInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireQueuedSpinLock" => {
            let addr = KeAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireSpinLockAtDpcLevel" => {
            let addr = KeAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireSpinLockForDpc" => {
            let addr = KeAcquireSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireSpinLockRaiseToDpc" => {
            let addr = KeAcquireSpinLockRaiseToDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAcquireSpinLockRaiseToSynch" => {
            let addr = KeAcquireSpinLockRaiseToSynch as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAddTriageDumpDataBlock" => {
            let addr = KeAddTriageDumpDataBlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAreAllApcsDisabled" => {
            let addr = KeAreAllApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAreApcsDisabled" => {
            let addr = KeAreApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeAttachProcess" => {
            let addr = KeAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeBugCheck" => {
            let addr = KeBugCheck as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeBugCheckEx" => {
            let addr = KeBugCheckEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeCancelTimer" => {
            let addr = KeCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeClearEvent" => {
            let addr = KeClearEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeConvertAuxiliaryCounterToPerformanceCounter" => {
            let addr = KeConvertAuxiliaryCounterToPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeConvertPerformanceCounterToAuxiliaryCounter" => {
            let addr = KeConvertPerformanceCounterToAuxiliaryCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDelayExecutionThread" => {
            let addr = KeDelayExecutionThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDeregisterBoundCallback" => {
            let addr = KeDeregisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDeregisterBugCheckCallback" => {
            let addr = KeDeregisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDeregisterBugCheckReasonCallback" => {
            let addr = KeDeregisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDeregisterNmiCallback" => {
            let addr = KeDeregisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDeregisterProcessorChangeCallback" => {
            let addr = KeDeregisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeDetachProcess" => {
            let addr = KeDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeEnterCriticalRegion" => {
            let addr = KeEnterCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeEnterGuardedRegion" => {
            let addr = KeEnterGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeExpandKernelStackAndCallout" => {
            let addr = KeExpandKernelStackAndCallout as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeExpandKernelStackAndCalloutEx" => {
            let addr = KeExpandKernelStackAndCalloutEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeFlushIoBuffers" => {
            let addr = KeFlushIoBuffers as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeFlushQueuedDpcs" => {
            let addr = KeFlushQueuedDpcs as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeFlushWriteBuffer" => {
            let addr = KeFlushWriteBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetCurrentIrql" => {
            let addr = KeGetCurrentIrql as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetCurrentNodeNumber" => {
            let addr = KeGetCurrentNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetCurrentProcessorNumberEx" => {
            let addr = KeGetCurrentProcessorNumberEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetProcessorIndexFromNumber" => {
            let addr = KeGetProcessorIndexFromNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetProcessorNumberFromIndex" => {
            let addr = KeGetProcessorNumberFromIndex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeGetRecommendedSharedDataAlignment" => {
            let addr = KeGetRecommendedSharedDataAlignment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeCrashDumpHeader" => {
            let addr = KeInitializeCrashDumpHeader as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeDeviceQueue" => {
            let addr = KeInitializeDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeDpc" => {
            let addr = KeInitializeDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeEvent" => {
            let addr = KeInitializeEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeGuardedMutex" => {
            let addr = KeInitializeGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeMutant" => {
            let addr = KeInitializeMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeMutex" => {
            let addr = KeInitializeMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeQueue" => {
            let addr = KeInitializeQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeSemaphore" => {
            let addr = KeInitializeSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeSpinLock" => {
            let addr = KeInitializeSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeThreadedDpc" => {
            let addr = KeInitializeThreadedDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeTimer" => {
            let addr = KeInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeTimerEx" => {
            let addr = KeInitializeTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInitializeTriageDumpDataArray" => {
            let addr = KeInitializeTriageDumpDataArray as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInsertByKeyDeviceQueue" => {
            let addr = KeInsertByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInsertDeviceQueue" => {
            let addr = KeInsertDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInsertHeadQueue" => {
            let addr = KeInsertHeadQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInsertQueue" => {
            let addr = KeInsertQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInsertQueueDpc" => {
            let addr = KeInsertQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInvalidateAllCaches" => {
            let addr = KeInvalidateAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeInvalidateRangeAllCaches" => {
            let addr = KeInvalidateRangeAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeIpiGenericCall" => {
            let addr = KeIpiGenericCall as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeIsExecutingDpc" => {
            let addr = KeIsExecutingDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeLeaveCriticalRegion" => {
            let addr = KeLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeLeaveGuardedRegion" => {
            let addr = KeLeaveGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeLowerIrql" => {
            let addr = KeLowerIrql as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KePulseEvent" => {
            let addr = KePulseEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryActiveGroupCount" => {
            let addr = KeQueryActiveGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryActiveProcessorCount" => {
            let addr = KeQueryActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryActiveProcessorCountEx" => {
            let addr = KeQueryActiveProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryActiveProcessors" => {
            let addr = KeQueryActiveProcessors as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryAuxiliaryCounterFrequency" => {
            let addr = KeQueryAuxiliaryCounterFrequency as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryDpcWatchdogInformation" => {
            let addr = KeQueryDpcWatchdogInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryGroupAffinity" => {
            let addr = KeQueryGroupAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryHardwareCounterConfiguration" => {
            let addr = KeQueryHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryHighestNodeNumber" => {
            let addr = KeQueryHighestNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryInterruptTimePrecise" => {
            let addr = KeQueryInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryLogicalProcessorRelationship" => {
            let addr = KeQueryLogicalProcessorRelationship as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryMaximumGroupCount" => {
            let addr = KeQueryMaximumGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryMaximumProcessorCount" => {
            let addr = KeQueryMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryMaximumProcessorCountEx" => {
            let addr = KeQueryMaximumProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryNodeActiveAffinity" => {
            let addr = KeQueryNodeActiveAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryNodeActiveAffinity2" => {
            let addr = KeQueryNodeActiveAffinity2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryNodeActiveProcessorCount" => {
            let addr = KeQueryNodeActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryNodeMaximumProcessorCount" => {
            let addr = KeQueryNodeMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryOwnerMutant" => {
            let addr = KeQueryOwnerMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryPerformanceCounter" => {
            let addr = KeQueryPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryPriorityThread" => {
            let addr = KeQueryPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryRuntimeThread" => {
            let addr = KeQueryRuntimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQuerySystemTimePrecise" => {
            let addr = KeQuerySystemTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryTimeIncrement" => {
            let addr = KeQueryTimeIncrement as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryTotalCycleTimeThread" => {
            let addr = KeQueryTotalCycleTimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryUnbiasedInterruptTime" => {
            let addr = KeQueryUnbiasedInterruptTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeQueryUnbiasedInterruptTimePrecise" => {
            let addr = KeQueryUnbiasedInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRcuReadLock" => {
            let addr = KeRcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRcuReadLockAtDpcLevel" => {
            let addr = KeRcuReadLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRcuReadUnlock" => {
            let addr = KeRcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRcuSynchronize" => {
            let addr = KeRcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateEvent" => {
            let addr = KeReadStateEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateMutant" => {
            let addr = KeReadStateMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateMutex" => {
            let addr = KeReadStateMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateQueue" => {
            let addr = KeReadStateQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateSemaphore" => {
            let addr = KeReadStateSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReadStateTimer" => {
            let addr = KeReadStateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRegisterBoundCallback" => {
            let addr = KeRegisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRegisterBugCheckCallback" => {
            let addr = KeRegisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRegisterBugCheckReasonCallback" => {
            let addr = KeRegisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRegisterNmiCallback" => {
            let addr = KeRegisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRegisterProcessorChangeCallback" => {
            let addr = KeRegisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseGuardedMutex" => {
            let addr = KeReleaseGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseGuardedMutexUnsafe" => {
            let addr = KeReleaseGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseInStackQueuedSpinLock" => {
            let addr = KeReleaseInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseInStackQueuedSpinLockForDpc" => {
            let addr = KeReleaseInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseInStackQueuedSpinLockFromDpcLevel" => {
            let addr = KeReleaseInStackQueuedSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseInterruptSpinLock" => {
            let addr = KeReleaseInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseMutant" => {
            let addr = KeReleaseMutant as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseMutex" => {
            let addr = KeReleaseMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseQueuedSpinLock" => {
            let addr = KeReleaseQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseSemaphore" => {
            let addr = KeReleaseSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseSpinLock" => {
            let addr = KeReleaseSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseSpinLockForDpc" => {
            let addr = KeReleaseSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeReleaseSpinLockFromDpcLevel" => {
            let addr = KeReleaseSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveByKeyDeviceQueue" => {
            let addr = KeRemoveByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveByKeyDeviceQueueIfBusy" => {
            let addr = KeRemoveByKeyDeviceQueueIfBusy as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveDeviceQueue" => {
            let addr = KeRemoveDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveEntryDeviceQueue" => {
            let addr = KeRemoveEntryDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveQueue" => {
            let addr = KeRemoveQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveQueueDpc" => {
            let addr = KeRemoveQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveQueueDpcEx" => {
            let addr = KeRemoveQueueDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRemoveQueueEx" => {
            let addr = KeRemoveQueueEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeResetEvent" => {
            let addr = KeResetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRestoreExtendedProcessorState" => {
            let addr = KeRestoreExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRevertToUserAffinityThread" => {
            let addr = KeRevertToUserAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRevertToUserAffinityThreadEx" => {
            let addr = KeRevertToUserAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRevertToUserGroupAffinityThread" => {
            let addr = KeRevertToUserGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeRundownQueue" => {
            let addr = KeRundownQueue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSaveExtendedProcessorState" => {
            let addr = KeSaveExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetBasePriorityThread" => {
            let addr = KeSetBasePriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetCoalescableTimer" => {
            let addr = KeSetCoalescableTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetEvent" => {
            let addr = KeSetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetHardwareCounterConfiguration" => {
            let addr = KeSetHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetIdealProcessorThread" => {
            let addr = KeSetIdealProcessorThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetImportanceDpc" => {
            let addr = KeSetImportanceDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetKernelStackSwapEnable" => {
            let addr = KeSetKernelStackSwapEnable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetPriorityThread" => {
            let addr = KeSetPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetSystemAffinityThread" => {
            let addr = KeSetSystemAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetSystemAffinityThreadEx" => {
            let addr = KeSetSystemAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetSystemGroupAffinityThread" => {
            let addr = KeSetSystemGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetTargetProcessorDpc" => {
            let addr = KeSetTargetProcessorDpc as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetTargetProcessorDpcEx" => {
            let addr = KeSetTargetProcessorDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetTimer" => {
            let addr = KeSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSetTimerEx" => {
            let addr = KeSetTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeShouldYieldProcessor" => {
            let addr = KeShouldYieldProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSrcuAllocate" => {
            let addr = KeSrcuAllocate as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSrcuFree" => {
            let addr = KeSrcuFree as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSrcuReadLock" => {
            let addr = KeSrcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSrcuReadUnlock" => {
            let addr = KeSrcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSrcuSynchronize" => {
            let addr = KeSrcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeStackAttachProcess" => {
            let addr = KeStackAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeStallExecutionProcessor" => {
            let addr = KeStallExecutionProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeSynchronizeExecution" => {
            let addr = KeSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeTestSpinLock" => {
            let addr = KeTestSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeTryToAcquireGuardedMutex" => {
            let addr = KeTryToAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeTryToAcquireQueuedSpinLock" => {
            let addr = KeTryToAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeTryToAcquireSpinLockAtDpcLevel" => {
            let addr = KeTryToAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeUnstackDetachProcess" => {
            let addr = KeUnstackDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeWaitForMultipleObjects" => {
            let addr = KeWaitForMultipleObjects as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "KeWaitForSingleObject" => {
            let addr = KeWaitForSingleObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAddPhysicalMemory" => {
            let addr = MmAddPhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAddVerifierSpecialThunks" => {
            let addr = MmAddVerifierSpecialThunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAddVerifierThunks" => {
            let addr = MmAddVerifierThunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAdvanceMdl" => {
            let addr = MmAdvanceMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateContiguousMemory" => {
            let addr = MmAllocateContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateContiguousMemoryEx" => {
            let addr = MmAllocateContiguousMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateContiguousMemorySpecifyCache" => {
            let addr = MmAllocateContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateContiguousMemorySpecifyCacheNode" => {
            let addr = MmAllocateContiguousMemorySpecifyCacheNode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateContiguousNodeMemory" => {
            let addr = MmAllocateContiguousNodeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateMappingAddress" => {
            let addr = MmAllocateMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateMappingAddressEx" => {
            let addr = MmAllocateMappingAddressEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateMdlForIoSpace" => {
            let addr = MmAllocateMdlForIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateNodePagesForMdlEx" => {
            let addr = MmAllocateNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocateNonCachedMemory" => {
            let addr = MmAllocateNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocatePagesForMdl" => {
            let addr = MmAllocatePagesForMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocatePagesForMdlEx" => {
            let addr = MmAllocatePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAllocatePartitionNodePagesForMdlEx" => {
            let addr = MmAllocatePartitionNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmAreMdlPagesCached" => {
            let addr = MmAreMdlPagesCached as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmBuildMdlForNonPagedPool" => {
            let addr = MmBuildMdlForNonPagedPool as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmCanFileBeTruncated" => {
            let addr = MmCanFileBeTruncated as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmCopyMemory" => {
            let addr = MmCopyMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmCreateMdl" => {
            let addr = MmCreateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmCreateMirror" => {
            let addr = MmCreateMirror as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmDoesFileHaveUserWritableReferences" => {
            let addr = MmDoesFileHaveUserWritableReferences as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFlushImageSection" => {
            let addr = MmFlushImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmForceSectionClosed" => {
            let addr = MmForceSectionClosed as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmForceSectionClosedEx" => {
            let addr = MmForceSectionClosedEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreeContiguousMemory" => {
            let addr = MmFreeContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreeContiguousMemorySpecifyCache" => {
            let addr = MmFreeContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreeMappingAddress" => {
            let addr = MmFreeMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreeNonCachedMemory" => {
            let addr = MmFreeNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreePagesFromMdl" => {
            let addr = MmFreePagesFromMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmFreePagesFromMdlEx" => {
            let addr = MmFreePagesFromMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetCacheAttribute" => {
            let addr = MmGetCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetCacheAttributeEx" => {
            let addr = MmGetCacheAttributeEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetMaximumFileSectionSize" => {
            let addr = MmGetMaximumFileSectionSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetModuleRoutineAddress" => {
            let addr = MmGetModuleRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetPhysicalAddress" => {
            let addr = MmGetPhysicalAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetPhysicalMemoryRanges" => {
            let addr = MmGetPhysicalMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetPhysicalMemoryRangesEx" => {
            let addr = MmGetPhysicalMemoryRangesEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetPhysicalMemoryRangesEx2" => {
            let addr = MmGetPhysicalMemoryRangesEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetSystemRoutineAddress" => {
            let addr = MmGetSystemRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmGetVirtualForPhysical" => {
            let addr = MmGetVirtualForPhysical as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsAddressValid" => {
            let addr = MmIsAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsDriverSuspectForVerifier" => {
            let addr = MmIsDriverSuspectForVerifier as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsDriverVerifying" => {
            let addr = MmIsDriverVerifying as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsDriverVerifyingByAddress" => {
            let addr = MmIsDriverVerifyingByAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsFileSectionActive" => {
            let addr = MmIsFileSectionActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsIoSpaceActive" => {
            let addr = MmIsIoSpaceActive as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsKernelAddress" => {
            let addr = MmIsKernelAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsNonPagedSystemAddressValid" => {
            let addr = MmIsNonPagedSystemAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsRecursiveIoFault" => {
            let addr = MmIsRecursiveIoFault as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsThisAnNtAsSystem" => {
            let addr = MmIsThisAnNtAsSystem as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsUserAddress" => {
            let addr = MmIsUserAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmIsVerifierEnabled" => {
            let addr = MmIsVerifierEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmLockPagableDataSection" => {
            let addr = MmLockPagableDataSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmLockPagableSectionByHandle" => {
            let addr = MmLockPagableSectionByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapIoSpace" => {
            let addr = MmMapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapIoSpaceEx" => {
            let addr = MmMapIoSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapLockedPages" => {
            let addr = MmMapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapLockedPagesSpecifyCache" => {
            let addr = MmMapLockedPagesSpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapLockedPagesWithReservedMapping" => {
            let addr = MmMapLockedPagesWithReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapMdl" => {
            let addr = MmMapMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapMemoryDumpMdlEx" => {
            let addr = MmMapMemoryDumpMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapUserAddressesToPage" => {
            let addr = MmMapUserAddressesToPage as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapVideoDisplay" => {
            let addr = MmMapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapViewInSessionSpace" => {
            let addr = MmMapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapViewInSessionSpaceEx" => {
            let addr = MmMapViewInSessionSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapViewInSystemSpace" => {
            let addr = MmMapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMapViewInSystemSpaceEx" => {
            let addr = MmMapViewInSystemSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMdlPageContentsState" => {
            let addr = MmMdlPageContentsState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmMdlPagesAreZero" => {
            let addr = MmMdlPagesAreZero as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmPageEntireDriver" => {
            let addr = MmPageEntireDriver as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmPrefetchPages" => {
            let addr = MmPrefetchPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProbeAndLockPages" => {
            let addr = MmProbeAndLockPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProbeAndLockPagesEx" => {
            let addr = MmProbeAndLockPagesEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProbeAndLockProcessPages" => {
            let addr = MmProbeAndLockProcessPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProbeAndLockSelectedPages" => {
            let addr = MmProbeAndLockSelectedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProtectDriverSection" => {
            let addr = MmProtectDriverSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmProtectMdlSystemAddress" => {
            let addr = MmProtectMdlSystemAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmQuerySystemSize" => {
            let addr = MmQuerySystemSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmRemovePhysicalMemory" => {
            let addr = MmRemovePhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmResetDriverPaging" => {
            let addr = MmResetDriverPaging as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmRotatePhysicalView" => {
            let addr = MmRotatePhysicalView as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmSecureVirtualMemory" => {
            let addr = MmSecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmSecureVirtualMemoryEx" => {
            let addr = MmSecureVirtualMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmSetAddressRangeModified" => {
            let addr = MmSetAddressRangeModified as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmSetPermanentCacheAttribute" => {
            let addr = MmSetPermanentCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmSizeOfMdl" => {
            let addr = MmSizeOfMdl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnlockPagableImageSection" => {
            let addr = MmUnlockPagableImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnlockPages" => {
            let addr = MmUnlockPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapIoSpace" => {
            let addr = MmUnmapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapLockedPages" => {
            let addr = MmUnmapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapReservedMapping" => {
            let addr = MmUnmapReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapVideoDisplay" => {
            let addr = MmUnmapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapViewInSessionSpace" => {
            let addr = MmUnmapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnmapViewInSystemSpace" => {
            let addr = MmUnmapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "MmUnsecureVirtualMemory" => {
            let addr = MmUnsecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAccessCheckAndAuditAlarm" => {
            let addr = NtAccessCheckAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAccessCheckByTypeAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarmByHandle" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarmByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAdjustGroupsToken" => {
            let addr = NtAdjustGroupsToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAdjustPrivilegesToken" => {
            let addr = NtAdjustPrivilegesToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtAllocateVirtualMemory" => {
            let addr = NtAllocateVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtClose" => {
            let addr = NtClose as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCloseObjectAuditAlarm" => {
            let addr = NtCloseObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCommitComplete" => {
            let addr = NtCommitComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCommitEnlistment" => {
            let addr = NtCommitEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCommitTransaction" => {
            let addr = NtCommitTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCopyFileChunk" => {
            let addr = NtCopyFileChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateEnlistment" => {
            let addr = NtCreateEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateFile" => {
            let addr = NtCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateResourceManager" => {
            let addr = NtCreateResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateSection" => {
            let addr = NtCreateSection as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateSectionEx" => {
            let addr = NtCreateSectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateTransaction" => {
            let addr = NtCreateTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtCreateTransactionManager" => {
            let addr = NtCreateTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtDeleteObjectAuditAlarm" => {
            let addr = NtDeleteObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtDeviceIoControlFile" => {
            let addr = NtDeviceIoControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtDuplicateToken" => {
            let addr = NtDuplicateToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtEnumerateTransactionObject" => {
            let addr = NtEnumerateTransactionObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtFilterToken" => {
            let addr = NtFilterToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtFlushBuffersFileEx" => {
            let addr = NtFlushBuffersFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtFreeVirtualMemory" => {
            let addr = NtFreeVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtFsControlFile" => {
            let addr = NtFsControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtGetNotificationResourceManager" => {
            let addr = NtGetNotificationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtImpersonateAnonymousToken" => {
            let addr = NtImpersonateAnonymousToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtLockFile" => {
            let addr = NtLockFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtManagePartition" => {
            let addr = NtManagePartition as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenEnlistment" => {
            let addr = NtOpenEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenFile" => {
            let addr = NtOpenFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenJobObjectToken" => {
            let addr = NtOpenJobObjectToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenObjectAuditAlarm" => {
            let addr = NtOpenObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenProcess" => {
            let addr = NtOpenProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenProcessToken" => {
            let addr = NtOpenProcessToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenProcessTokenEx" => {
            let addr = NtOpenProcessTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenRegistryTransaction" => {
            let addr = NtOpenRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenResourceManager" => {
            let addr = NtOpenResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenThreadToken" => {
            let addr = NtOpenThreadToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenThreadTokenEx" => {
            let addr = NtOpenThreadTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenTransaction" => {
            let addr = NtOpenTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtOpenTransactionManager" => {
            let addr = NtOpenTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPowerInformation" => {
            let addr = NtPowerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrePrepareComplete" => {
            let addr = NtPrePrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrePrepareEnlistment" => {
            let addr = NtPrePrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrepareComplete" => {
            let addr = NtPrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrepareEnlistment" => {
            let addr = NtPrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrivilegeCheck" => {
            let addr = NtPrivilegeCheck as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrivilegeObjectAuditAlarm" => {
            let addr = NtPrivilegeObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPrivilegedServiceAuditAlarm" => {
            let addr = NtPrivilegedServiceAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPropagationComplete" => {
            let addr = NtPropagationComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtPropagationFailed" => {
            let addr = NtPropagationFailed as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryDirectoryFile" => {
            let addr = NtQueryDirectoryFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryDirectoryFileEx" => {
            let addr = NtQueryDirectoryFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationByName" => {
            let addr = NtQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationEnlistment" => {
            let addr = NtQueryInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationFile" => {
            let addr = NtQueryInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationResourceManager" => {
            let addr = NtQueryInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationToken" => {
            let addr = NtQueryInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationTransaction" => {
            let addr = NtQueryInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryInformationTransactionManager" => {
            let addr = NtQueryInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryObject" => {
            let addr = NtQueryObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryQuotaInformationFile" => {
            let addr = NtQueryQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQuerySecurityObject" => {
            let addr = NtQuerySecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryVirtualMemory" => {
            let addr = NtQueryVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtQueryVolumeInformationFile" => {
            let addr = NtQueryVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtReadFile" => {
            let addr = NtReadFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtReadOnlyEnlistment" => {
            let addr = NtReadOnlyEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRecoverEnlistment" => {
            let addr = NtRecoverEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRecoverResourceManager" => {
            let addr = NtRecoverResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRecoverTransactionManager" => {
            let addr = NtRecoverTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRegisterProtocolAddressInformation" => {
            let addr = NtRegisterProtocolAddressInformation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRenameTransactionManager" => {
            let addr = NtRenameTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRollbackComplete" => {
            let addr = NtRollbackComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRollbackEnlistment" => {
            let addr = NtRollbackEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRollbackRegistryTransaction" => {
            let addr = NtRollbackRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRollbackTransaction" => {
            let addr = NtRollbackTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtRollforwardTransactionManager" => {
            let addr = NtRollforwardTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationEnlistment" => {
            let addr = NtSetInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationFile" => {
            let addr = NtSetInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationResourceManager" => {
            let addr = NtSetInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationThread" => {
            let addr = NtSetInformationThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationToken" => {
            let addr = NtSetInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationTransaction" => {
            let addr = NtSetInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationTransactionManager" => {
            let addr = NtSetInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetInformationVirtualMemory" => {
            let addr = NtSetInformationVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetQuotaInformationFile" => {
            let addr = NtSetQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetSecurityObject" => {
            let addr = NtSetSecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSetVolumeInformationFile" => {
            let addr = NtSetVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtSinglePhaseReject" => {
            let addr = NtSinglePhaseReject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtUnlockFile" => {
            let addr = NtUnlockFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "NtWriteFile" => {
            let addr = NtWriteFile as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObCloseHandle" => {
            let addr = ObCloseHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObCloseHandleWithResult" => {
            let addr = ObCloseHandleWithResult as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObDereferenceObjectDeferDelete" => {
            let addr = ObDereferenceObjectDeferDelete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObDereferenceObjectDeferDeleteWithTag" => {
            let addr = ObDereferenceObjectDeferDeleteWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObGetFilterVersion" => {
            let addr = ObGetFilterVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObGetObjectSecurity" => {
            let addr = ObGetObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObInsertObject" => {
            let addr = ObInsertObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObIsKernelHandle" => {
            let addr = ObIsKernelHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObMakeTemporaryObject" => {
            let addr = ObMakeTemporaryObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObOpenObjectByPointer" => {
            let addr = ObOpenObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObOpenObjectByPointerWithTag" => {
            let addr = ObOpenObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObQueryNameString" => {
            let addr = ObQueryNameString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObQueryObjectAuditingByHandle" => {
            let addr = ObQueryObjectAuditingByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectByHandle" => {
            let addr = ObReferenceObjectByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectByHandleWithTag" => {
            let addr = ObReferenceObjectByHandleWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectByPointer" => {
            let addr = ObReferenceObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectByPointerWithTag" => {
            let addr = ObReferenceObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectSafe" => {
            let addr = ObReferenceObjectSafe as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReferenceObjectSafeWithTag" => {
            let addr = ObReferenceObjectSafeWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObRegisterCallbacks" => {
            let addr = ObRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObReleaseObjectSecurity" => {
            let addr = ObReleaseObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObUnRegisterCallbacks" => {
            let addr = ObUnRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObfDereferenceObject" => {
            let addr = ObfDereferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObfDereferenceObjectWithTag" => {
            let addr = ObfDereferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObfReferenceObject" => {
            let addr = ObfReferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "ObfReferenceObjectWithTag" => {
            let addr = ObfReferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsAcquireSiloHardReference" => {
            let addr = PsAcquireSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsAllocSiloContextSlot" => {
            let addr = PsAllocSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsAllocateAffinityToken" => {
            let addr = PsAllocateAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsAssignImpersonationToken" => {
            let addr = PsAssignImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsAttachSiloToCurrentThread" => {
            let addr = PsAttachSiloToCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsChargePoolQuota" => {
            let addr = PsChargePoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsChargeProcessPoolQuota" => {
            let addr = PsChargeProcessPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsCreateSiloContext" => {
            let addr = PsCreateSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsCreateSystemThread" => {
            let addr = PsCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsDereferenceImpersonationToken" => {
            let addr = PsDereferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsDereferencePrimaryToken" => {
            let addr = PsDereferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsDereferenceSiloContext" => {
            let addr = PsDereferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsDetachSiloFromCurrentThread" => {
            let addr = PsDetachSiloFromCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsDisableImpersonation" => {
            let addr = PsDisableImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsFreeAffinityToken" => {
            let addr = PsFreeAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsFreeSiloContextSlot" => {
            let addr = PsFreeSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentProcessId" => {
            let addr = PsGetCurrentProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentServerSilo" => {
            let addr = PsGetCurrentServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentServerSiloName" => {
            let addr = PsGetCurrentServerSiloName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentSilo" => {
            let addr = PsGetCurrentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentThreadId" => {
            let addr = PsGetCurrentThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetCurrentThreadTeb" => {
            let addr = PsGetCurrentThreadTeb as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetEffectiveServerSilo" => {
            let addr = PsGetEffectiveServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetHostSilo" => {
            let addr = PsGetHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetJobServerSilo" => {
            let addr = PsGetJobServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetJobSilo" => {
            let addr = PsGetJobSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetParentSilo" => {
            let addr = PsGetParentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetPermanentSiloContext" => {
            let addr = PsGetPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetProcessCreateTimeQuadPart" => {
            let addr = PsGetProcessCreateTimeQuadPart as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetProcessExitStatus" => {
            let addr = PsGetProcessExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetProcessExitTime" => {
            let addr = PsGetProcessExitTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetProcessId" => {
            let addr = PsGetProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetProcessStartKey" => {
            let addr = PsGetProcessStartKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetServerSiloActiveConsoleId" => {
            let addr = PsGetServerSiloActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetServerSiloServiceSessionId" => {
            let addr = PsGetServerSiloServiceSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetSiloContainerId" => {
            let addr = PsGetSiloContainerId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetSiloContext" => {
            let addr = PsGetSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetSiloMonitorContextSlot" => {
            let addr = PsGetSiloMonitorContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadCreateTime" => {
            let addr = PsGetThreadCreateTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadExitStatus" => {
            let addr = PsGetThreadExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadId" => {
            let addr = PsGetThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadProcess" => {
            let addr = PsGetThreadProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadProcessId" => {
            let addr = PsGetThreadProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadProperty" => {
            let addr = PsGetThreadProperty as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetThreadServerSilo" => {
            let addr = PsGetThreadServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsGetVersion" => {
            let addr = PsGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsImpersonateClient" => {
            let addr = PsImpersonateClient as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsInsertPermanentSiloContext" => {
            let addr = PsInsertPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsInsertSiloContext" => {
            let addr = PsInsertSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsCurrentThreadInServerSilo" => {
            let addr = PsIsCurrentThreadInServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsCurrentThreadPrefetching" => {
            let addr = PsIsCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsDiskCountersEnabled" => {
            let addr = PsIsDiskCountersEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsHostSilo" => {
            let addr = PsIsHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsSystemThread" => {
            let addr = PsIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsThreadAttachedToSpecificSilo" => {
            let addr = PsIsThreadAttachedToSpecificSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsIsThreadTerminating" => {
            let addr = PsIsThreadTerminating as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsLookupProcessByProcessId" => {
            let addr = PsLookupProcessByProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsLookupThreadByThreadId" => {
            let addr = PsLookupThreadByThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsMakeSiloContextPermanent" => {
            let addr = PsMakeSiloContextPermanent as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsQueryProcessAvailableCpus" => {
            let addr = PsQueryProcessAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsQueryProcessAvailableCpusCount" => {
            let addr = PsQueryProcessAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsQuerySystemAvailableCpus" => {
            let addr = PsQuerySystemAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsQuerySystemAvailableCpusCount" => {
            let addr = PsQuerySystemAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsQueryTotalCycleTimeProcess" => {
            let addr = PsQueryTotalCycleTimeProcess as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReferenceImpersonationToken" => {
            let addr = PsReferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReferencePrimaryToken" => {
            let addr = PsReferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReferenceSiloContext" => {
            let addr = PsReferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRegisterProcessAvailableCpusChangeNotification" => {
            let addr = PsRegisterProcessAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRegisterSiloMonitor" => {
            let addr = PsRegisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRegisterSystemAvailableCpusChangeNotification" => {
            let addr = PsRegisterSystemAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReleaseSiloHardReference" => {
            let addr = PsReleaseSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRemoveCreateThreadNotifyRoutine" => {
            let addr = PsRemoveCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRemoveLoadImageNotifyRoutine" => {
            let addr = PsRemoveLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRemoveSiloContext" => {
            let addr = PsRemoveSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReplaceSiloContext" => {
            let addr = PsReplaceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRestoreImpersonation" => {
            let addr = PsRestoreImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsReturnPoolQuota" => {
            let addr = PsReturnPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRevertToSelf" => {
            let addr = PsRevertToSelf as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsRevertToUserMultipleGroupAffinityThread" => {
            let addr = PsRevertToUserMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCreateProcessNotifyRoutine" => {
            let addr = PsSetCreateProcessNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCreateProcessNotifyRoutineEx" => {
            let addr = PsSetCreateProcessNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCreateProcessNotifyRoutineEx2" => {
            let addr = PsSetCreateProcessNotifyRoutineEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCreateThreadNotifyRoutine" => {
            let addr = PsSetCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCreateThreadNotifyRoutineEx" => {
            let addr = PsSetCreateThreadNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetCurrentThreadPrefetching" => {
            let addr = PsSetCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetLoadImageNotifyRoutine" => {
            let addr = PsSetLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetLoadImageNotifyRoutineEx" => {
            let addr = PsSetLoadImageNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsSetSystemMultipleGroupAffinityThread" => {
            let addr = PsSetSystemMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsStartSiloMonitor" => {
            let addr = PsStartSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsTerminateServerSilo" => {
            let addr = PsTerminateServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsTerminateSystemThread" => {
            let addr = PsTerminateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsUnregisterAvailableCpusChangeNotification" => {
            let addr = PsUnregisterAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsUnregisterSiloMonitor" => {
            let addr = PsUnregisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsUpdateDiskCounters" => {
            let addr = PsUpdateDiskCounters as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PsWrapApcWow64Thread" => {
            let addr = PsWrapApcWow64Thread as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedAllocateMemory" => {
            let addr = PshedAllocateMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedFreeMemory" => {
            let addr = PshedFreeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedIsSystemWheaEnabled" => {
            let addr = PshedIsSystemWheaEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedRegisterPlugin" => {
            let addr = PshedRegisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedSynchronizeExecution" => {
            let addr = PshedSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "PshedUnregisterPlugin" => {
            let addr = PshedUnregisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAbsoluteToSelfRelativeSD" => {
            let addr = RtlAbsoluteToSelfRelativeSD as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAddAccessAllowedAce" => {
            let addr = RtlAddAccessAllowedAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAddAccessAllowedAceEx" => {
            let addr = RtlAddAccessAllowedAceEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAddAce" => {
            let addr = RtlAddAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAllocateAndInitializeSid" => {
            let addr = RtlAllocateAndInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAllocateAndInitializeSidEx" => {
            let addr = RtlAllocateAndInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAllocateHeap" => {
            let addr = RtlAllocateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAnsiStringToUnicodeString" => {
            let addr = RtlAnsiStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAppendStringToString" => {
            let addr = RtlAppendStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAppendUnicodeStringToString" => {
            let addr = RtlAppendUnicodeStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAppendUnicodeToString" => {
            let addr = RtlAppendUnicodeToString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAreBitsClear" => {
            let addr = RtlAreBitsClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAreBitsSet" => {
            let addr = RtlAreBitsSet as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlAssert" => {
            let addr = RtlAssert as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCaptureContext" => {
            let addr = RtlCaptureContext as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCaptureContext2" => {
            let addr = RtlCaptureContext2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCaptureStackBackTrace" => {
            let addr = RtlCaptureStackBackTrace as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCharToInteger" => {
            let addr = RtlCharToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCheckRegistryKey" => {
            let addr = RtlCheckRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlClearAllBits" => {
            let addr = RtlClearAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlClearBit" => {
            let addr = RtlClearBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlClearBits" => {
            let addr = RtlClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCmDecodeMemIoResource" => {
            let addr = RtlCmDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCmEncodeMemIoResource" => {
            let addr = RtlCmEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareAltitudes" => {
            let addr = RtlCompareAltitudes as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareMemory" => {
            let addr = RtlCompareMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareMemoryUlong" => {
            let addr = RtlCompareMemoryUlong as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareString" => {
            let addr = RtlCompareString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareUnicodeString" => {
            let addr = RtlCompareUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompareUnicodeStrings" => {
            let addr = RtlCompareUnicodeStrings as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompressBuffer" => {
            let addr = RtlCompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCompressChunks" => {
            let addr = RtlCompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlContractHashTable" => {
            let addr = RtlContractHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlConvertSidToUnicodeString" => {
            let addr = RtlConvertSidToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyBitMap" => {
            let addr = RtlCopyBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyDeviceMemory" => {
            let addr = RtlCopyDeviceMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyLuid" => {
            let addr = RtlCopyLuid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyMemoryNonTemporal" => {
            let addr = RtlCopyMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopySid" => {
            let addr = RtlCopySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyString" => {
            let addr = RtlCopyString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyUnicodeString" => {
            let addr = RtlCopyUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCopyVolatileMemory" => {
            let addr = RtlCopyVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCrc32" => {
            let addr = RtlCrc32 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCrc64" => {
            let addr = RtlCrc64 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateAcl" => {
            let addr = RtlCreateAcl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateHashTable" => {
            let addr = RtlCreateHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateHashTableEx" => {
            let addr = RtlCreateHashTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateHeap" => {
            let addr = RtlCreateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateRegistryKey" => {
            let addr = RtlCreateRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateSecurityDescriptor" => {
            let addr = RtlCreateSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateSecurityDescriptorRelative" => {
            let addr = RtlCreateSecurityDescriptorRelative as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateServiceSid" => {
            let addr = RtlCreateServiceSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateSystemVolumeInformationFolder" => {
            let addr = RtlCreateSystemVolumeInformationFolder as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateUnicodeString" => {
            let addr = RtlCreateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCreateVirtualAccountSid" => {
            let addr = RtlCreateVirtualAccountSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlCustomCPToUnicodeN" => {
            let addr = RtlCustomCPToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressBuffer" => {
            let addr = RtlDecompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressBufferEx" => {
            let addr = RtlDecompressBufferEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressBufferEx2" => {
            let addr = RtlDecompressBufferEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressChunks" => {
            let addr = RtlDecompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressFragment" => {
            let addr = RtlDecompressFragment as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDecompressFragmentEx" => {
            let addr = RtlDecompressFragmentEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDelete" => {
            let addr = RtlDelete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteAce" => {
            let addr = RtlDeleteAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteElementGenericTable" => {
            let addr = RtlDeleteElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteElementGenericTableAvl" => {
            let addr = RtlDeleteElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteElementGenericTableAvlEx" => {
            let addr = RtlDeleteElementGenericTableAvlEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteHashTable" => {
            let addr = RtlDeleteHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteNoSplay" => {
            let addr = RtlDeleteNoSplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDeleteRegistryValue" => {
            let addr = RtlDeleteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDescribeChunk" => {
            let addr = RtlDescribeChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDestroyHeap" => {
            let addr = RtlDestroyHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDowncaseUnicodeChar" => {
            let addr = RtlDowncaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDowncaseUnicodeString" => {
            let addr = RtlDowncaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDrainNonVolatileFlush" => {
            let addr = RtlDrainNonVolatileFlush as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlDuplicateUnicodeString" => {
            let addr = RtlDuplicateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEndEnumerationHashTable" => {
            let addr = RtlEndEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEndStrongEnumerationHashTable" => {
            let addr = RtlEndStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEndWeakEnumerationHashTable" => {
            let addr = RtlEndWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateEntryHashTable" => {
            let addr = RtlEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateGenericTable" => {
            let addr = RtlEnumerateGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateGenericTableAvl" => {
            let addr = RtlEnumerateGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateGenericTableLikeADirectory" => {
            let addr = RtlEnumerateGenericTableLikeADirectory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateGenericTableWithoutSplaying" => {
            let addr = RtlEnumerateGenericTableWithoutSplaying as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEnumerateGenericTableWithoutSplayingAvl" => {
            let addr = RtlEnumerateGenericTableWithoutSplayingAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEqualPrefixSid" => {
            let addr = RtlEqualPrefixSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEqualSid" => {
            let addr = RtlEqualSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEqualString" => {
            let addr = RtlEqualString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlEqualUnicodeString" => {
            let addr = RtlEqualUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlExpandHashTable" => {
            let addr = RtlExpandHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlExtendCorrelationVector" => {
            let addr = RtlExtendCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlExtractBitMap" => {
            let addr = RtlExtractBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFillMemoryNonTemporal" => {
            let addr = RtlFillMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFillNonVolatileMemory" => {
            let addr = RtlFillNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindClearBits" => {
            let addr = RtlFindClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindClearBitsAndSet" => {
            let addr = RtlFindClearBitsAndSet as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindClearRuns" => {
            let addr = RtlFindClearRuns as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindClosestEncodableLength" => {
            let addr = RtlFindClosestEncodableLength as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindFirstRunClear" => {
            let addr = RtlFindFirstRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindLastBackwardRunClear" => {
            let addr = RtlFindLastBackwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindLeastSignificantBit" => {
            let addr = RtlFindLeastSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindLongestRunClear" => {
            let addr = RtlFindLongestRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindMostSignificantBit" => {
            let addr = RtlFindMostSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindNextForwardRunClear" => {
            let addr = RtlFindNextForwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindSetBits" => {
            let addr = RtlFindSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindSetBitsAndClear" => {
            let addr = RtlFindSetBitsAndClear as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFindUnicodePrefix" => {
            let addr = RtlFindUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFlushNonVolatileMemory" => {
            let addr = RtlFlushNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFlushNonVolatileMemoryRanges" => {
            let addr = RtlFlushNonVolatileMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeAnsiString" => {
            let addr = RtlFreeAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeHeap" => {
            let addr = RtlFreeHeap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeNonVolatileToken" => {
            let addr = RtlFreeNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeOemString" => {
            let addr = RtlFreeOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeSid" => {
            let addr = RtlFreeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeUTF8String" => {
            let addr = RtlFreeUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlFreeUnicodeString" => {
            let addr = RtlFreeUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGUIDFromString" => {
            let addr = RtlGUIDFromString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGenerate8dot3Name" => {
            let addr = RtlGenerate8dot3Name as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGenerateClass5Guid" => {
            let addr = RtlGenerateClass5Guid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetAce" => {
            let addr = RtlGetAce as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetAcesBufferSize" => {
            let addr = RtlGetAcesBufferSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetActiveConsoleId" => {
            let addr = RtlGetActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetCallersAddress" => {
            let addr = RtlGetCallersAddress as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetCompressionWorkSpaceSize" => {
            let addr = RtlGetCompressionWorkSpaceSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetConsoleSessionForegroundProcessId" => {
            let addr = RtlGetConsoleSessionForegroundProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetDaclSecurityDescriptor" => {
            let addr = RtlGetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetElementGenericTable" => {
            let addr = RtlGetElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetElementGenericTableAvl" => {
            let addr = RtlGetElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetEnabledExtendedFeatures" => {
            let addr = RtlGetEnabledExtendedFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetGroupSecurityDescriptor" => {
            let addr = RtlGetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetNextEntryHashTable" => {
            let addr = RtlGetNextEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetNonVolatileToken" => {
            let addr = RtlGetNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetNtProductType" => {
            let addr = RtlGetNtProductType as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetNtSystemRoot" => {
            let addr = RtlGetNtSystemRoot as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetOwnerSecurityDescriptor" => {
            let addr = RtlGetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetPersistedStateLocation" => {
            let addr = RtlGetPersistedStateLocation as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetProductInfo" => {
            let addr = RtlGetProductInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetSaclSecurityDescriptor" => {
            let addr = RtlGetSaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetSuiteMask" => {
            let addr = RtlGetSuiteMask as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetSystemGlobalData" => {
            let addr = RtlGetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlGetVersion" => {
            let addr = RtlGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlHashUnicodeString" => {
            let addr = RtlHashUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIdentifierAuthoritySid" => {
            let addr = RtlIdentifierAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIdnToAscii" => {
            let addr = RtlIdnToAscii as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIdnToNameprepUnicode" => {
            let addr = RtlIdnToNameprepUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIdnToUnicode" => {
            let addr = RtlIdnToUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIncrementCorrelationVector" => {
            let addr = RtlIncrementCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitAnsiString" => {
            let addr = RtlInitAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitAnsiStringEx" => {
            let addr = RtlInitAnsiStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitCodePageTable" => {
            let addr = RtlInitCodePageTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitEnumerationHashTable" => {
            let addr = RtlInitEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitString" => {
            let addr = RtlInitString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitStringEx" => {
            let addr = RtlInitStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitStrongEnumerationHashTable" => {
            let addr = RtlInitStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitUTF8String" => {
            let addr = RtlInitUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitUTF8StringEx" => {
            let addr = RtlInitUTF8StringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitUnicodeString" => {
            let addr = RtlInitUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitUnicodeStringEx" => {
            let addr = RtlInitUnicodeStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitWeakEnumerationHashTable" => {
            let addr = RtlInitWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeBitMap" => {
            let addr = RtlInitializeBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeCorrelationVector" => {
            let addr = RtlInitializeCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeGenericTable" => {
            let addr = RtlInitializeGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeGenericTableAvl" => {
            let addr = RtlInitializeGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeSid" => {
            let addr = RtlInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeSidEx" => {
            let addr = RtlInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInitializeUnicodePrefix" => {
            let addr = RtlInitializeUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertElementGenericTable" => {
            let addr = RtlInsertElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertElementGenericTableAvl" => {
            let addr = RtlInsertElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertElementGenericTableFull" => {
            let addr = RtlInsertElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertElementGenericTableFullAvl" => {
            let addr = RtlInsertElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertEntryHashTable" => {
            let addr = RtlInsertEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInsertUnicodePrefix" => {
            let addr = RtlInsertUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlInt64ToUnicodeString" => {
            let addr = RtlInt64ToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIntegerToUnicodeString" => {
            let addr = RtlIntegerToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIoDecodeMemIoResource" => {
            let addr = RtlIoDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIoEncodeMemIoResource" => {
            let addr = RtlIoEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsApiSetImplemented" => {
            let addr = RtlIsApiSetImplemented as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsCloudFilesPlaceholder" => {
            let addr = RtlIsCloudFilesPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsGenericTableEmpty" => {
            let addr = RtlIsGenericTableEmpty as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsGenericTableEmptyAvl" => {
            let addr = RtlIsGenericTableEmptyAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsMultiSessionSku" => {
            let addr = RtlIsMultiSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsMultiUsersInSessionSku" => {
            let addr = RtlIsMultiUsersInSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsNameLegalDOS8Dot3" => {
            let addr = RtlIsNameLegalDOS8Dot3 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsNonEmptyDirectoryReparsePointAllowed" => {
            let addr = RtlIsNonEmptyDirectoryReparsePointAllowed as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsNormalizedString" => {
            let addr = RtlIsNormalizedString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsNtDdiVersionAvailable" => {
            let addr = RtlIsNtDdiVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsPartialPlaceholder" => {
            let addr = RtlIsPartialPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsPartialPlaceholderFileHandle" => {
            let addr = RtlIsPartialPlaceholderFileHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsPartialPlaceholderFileInfo" => {
            let addr = RtlIsPartialPlaceholderFileInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsSandboxedToken" => {
            let addr = RtlIsSandboxedToken as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsSandboxedTokenHandle" => {
            let addr = RtlIsSandboxedTokenHandle as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsServicePackVersionInstalled" => {
            let addr = RtlIsServicePackVersionInstalled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsStateSeparationEnabled" => {
            let addr = RtlIsStateSeparationEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsUntrustedObject" => {
            let addr = RtlIsUntrustedObject as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsValidOemCharacter" => {
            let addr = RtlIsValidOemCharacter as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlIsZeroMemory" => {
            let addr = RtlIsZeroMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLengthRequiredSid" => {
            let addr = RtlLengthRequiredSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLengthSecurityDescriptor" => {
            let addr = RtlLengthSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLengthSid" => {
            let addr = RtlLengthSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupElementGenericTable" => {
            let addr = RtlLookupElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupElementGenericTableAvl" => {
            let addr = RtlLookupElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupElementGenericTableFull" => {
            let addr = RtlLookupElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupElementGenericTableFullAvl" => {
            let addr = RtlLookupElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupEntryHashTable" => {
            let addr = RtlLookupEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlLookupFirstMatchingElementGenericTableAvl" => {
            let addr = RtlLookupFirstMatchingElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlMapGenericMask" => {
            let addr = RtlMapGenericMask as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlMoveVolatileMemory" => {
            let addr = RtlMoveVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlMultiByteToUnicodeN" => {
            let addr = RtlMultiByteToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlMultiByteToUnicodeSize" => {
            let addr = RtlMultiByteToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNextUnicodePrefix" => {
            let addr = RtlNextUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNormalizeSecurityDescriptor" => {
            let addr = RtlNormalizeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNormalizeString" => {
            let addr = RtlNormalizeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNtStatusToDosError" => {
            let addr = RtlNtStatusToDosError as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNtStatusToDosErrorNoTeb" => {
            let addr = RtlNtStatusToDosErrorNoTeb as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberGenericTableElements" => {
            let addr = RtlNumberGenericTableElements as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberGenericTableElementsAvl" => {
            let addr = RtlNumberGenericTableElementsAvl as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberOfClearBits" => {
            let addr = RtlNumberOfClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberOfClearBitsInRange" => {
            let addr = RtlNumberOfClearBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberOfSetBits" => {
            let addr = RtlNumberOfSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberOfSetBitsInRange" => {
            let addr = RtlNumberOfSetBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlNumberOfSetBitsUlongPtr" => {
            let addr = RtlNumberOfSetBitsUlongPtr as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlOemStringToCountedUnicodeString" => {
            let addr = RtlOemStringToCountedUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlOemStringToUnicodeString" => {
            let addr = RtlOemStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlOemToUnicodeN" => {
            let addr = RtlOemToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlOsDeploymentState" => {
            let addr = RtlOsDeploymentState as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlPrefetchMemoryNonTemporal" => {
            let addr = RtlPrefetchMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlPrefixString" => {
            let addr = RtlPrefixString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlPrefixUnicodeString" => {
            let addr = RtlPrefixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryPackageIdentity" => {
            let addr = RtlQueryPackageIdentity as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryPackageIdentityEx" => {
            let addr = RtlQueryPackageIdentityEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryProcessPlaceholderCompatibilityMode" => {
            let addr = RtlQueryProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryRegistryValueWithFallback" => {
            let addr = RtlQueryRegistryValueWithFallback as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryRegistryValues" => {
            let addr = RtlQueryRegistryValues as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryThreadPlaceholderCompatibilityMode" => {
            let addr = RtlQueryThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlQueryValidationRunlevel" => {
            let addr = RtlQueryValidationRunlevel as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRaiseCustomSystemEventTrigger" => {
            let addr = RtlRaiseCustomSystemEventTrigger as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRandom" => {
            let addr = RtlRandom as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRandomEx" => {
            let addr = RtlRandomEx as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRealPredecessor" => {
            let addr = RtlRealPredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRealSuccessor" => {
            let addr = RtlRealSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRemoveEntryHashTable" => {
            let addr = RtlRemoveEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRemoveUnicodePrefix" => {
            let addr = RtlRemoveUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlReplaceSidInSd" => {
            let addr = RtlReplaceSidInSd as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlReserveChunk" => {
            let addr = RtlReserveChunk as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRunOnceBeginInitialize" => {
            let addr = RtlRunOnceBeginInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRunOnceComplete" => {
            let addr = RtlRunOnceComplete as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRunOnceExecuteOnce" => {
            let addr = RtlRunOnceExecuteOnce as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlRunOnceInitialize" => {
            let addr = RtlRunOnceInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSecondsSince1970ToTime" => {
            let addr = RtlSecondsSince1970ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSecondsSince1980ToTime" => {
            let addr = RtlSecondsSince1980ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSelfRelativeToAbsoluteSD" => {
            let addr = RtlSelfRelativeToAbsoluteSD as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetAllBits" => {
            let addr = RtlSetAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetBit" => {
            let addr = RtlSetBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetBits" => {
            let addr = RtlSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetDaclSecurityDescriptor" => {
            let addr = RtlSetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetGroupSecurityDescriptor" => {
            let addr = RtlSetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetOwnerSecurityDescriptor" => {
            let addr = RtlSetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetProcessPlaceholderCompatibilityMode" => {
            let addr = RtlSetProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetSystemGlobalData" => {
            let addr = RtlSetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetThreadPlaceholderCompatibilityMode" => {
            let addr = RtlSetThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSetVolatileMemory" => {
            let addr = RtlSetVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSplay" => {
            let addr = RtlSplay as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlStringFromGUID" => {
            let addr = RtlStringFromGUID as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlStronglyEnumerateEntryHashTable" => {
            let addr = RtlStronglyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSubAuthorityCountSid" => {
            let addr = RtlSubAuthorityCountSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSubAuthoritySid" => {
            let addr = RtlSubAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSubtreePredecessor" => {
            let addr = RtlSubtreePredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSubtreeSuccessor" => {
            let addr = RtlSubtreeSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlSuffixUnicodeString" => {
            let addr = RtlSuffixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlTestBit" => {
            let addr = RtlTestBit as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlTimeFieldsToTime" => {
            let addr = RtlTimeFieldsToTime as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlTimeToSecondsSince1970" => {
            let addr = RtlTimeToSecondsSince1970 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlTimeToSecondsSince1980" => {
            let addr = RtlTimeToSecondsSince1980 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlTimeToTimeFields" => {
            let addr = RtlTimeToTimeFields as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUTF8StringToUnicodeString" => {
            let addr = RtlUTF8StringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUTF8ToUnicodeN" => {
            let addr = RtlUTF8ToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToAnsiString" => {
            let addr = RtlUnicodeStringToAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToCountedOemString" => {
            let addr = RtlUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToInt64" => {
            let addr = RtlUnicodeStringToInt64 as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToInteger" => {
            let addr = RtlUnicodeStringToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToOemString" => {
            let addr = RtlUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeStringToUTF8String" => {
            let addr = RtlUnicodeStringToUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeToCustomCPN" => {
            let addr = RtlUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeToMultiByteN" => {
            let addr = RtlUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeToMultiByteSize" => {
            let addr = RtlUnicodeToMultiByteSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeToOemN" => {
            let addr = RtlUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUnicodeToUTF8N" => {
            let addr = RtlUnicodeToUTF8N as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeChar" => {
            let addr = RtlUpcaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeString" => {
            let addr = RtlUpcaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeStringToCountedOemString" => {
            let addr = RtlUpcaseUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeStringToOemString" => {
            let addr = RtlUpcaseUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeToCustomCPN" => {
            let addr = RtlUpcaseUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeToMultiByteN" => {
            let addr = RtlUpcaseUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpcaseUnicodeToOemN" => {
            let addr = RtlUpcaseUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpperChar" => {
            let addr = RtlUpperChar as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlUpperString" => {
            let addr = RtlUpperString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlValidRelativeSecurityDescriptor" => {
            let addr = RtlValidRelativeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlValidSecurityDescriptor" => {
            let addr = RtlValidSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlValidSid" => {
            let addr = RtlValidSid as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlValidateCorrelationVector" => {
            let addr = RtlValidateCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlValidateUnicodeString" => {
            let addr = RtlValidateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlVerifyVersionInfo" => {
            let addr = RtlVerifyVersionInfo as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlVolumeDeviceToDosName" => {
            let addr = RtlVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlWalkFrameChain" => {
            let addr = RtlWalkFrameChain as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlWeaklyEnumerateEntryHashTable" => {
            let addr = RtlWeaklyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlWriteNonVolatileMemory" => {
            let addr = RtlWriteNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlWriteRegistryValue" => {
            let addr = RtlWriteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlxAnsiStringToUnicodeSize" => {
            let addr = RtlxAnsiStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlxOemStringToUnicodeSize" => {
            let addr = RtlxOemStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlxUnicodeStringToAnsiSize" => {
            let addr = RtlxUnicodeStringToAnsiSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        "RtlxUnicodeStringToOemSize" => {
            let addr = RtlxUnicodeStringToOemSize as u64;
            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }
        }
        _ => Err(HookError::InvalidAddress),
    }
}

pub fn remove_inline_hook_by_name(name: &str) -> Result<(), HookError> {
    match name {
        "DbgBreakPointWithStatus" => {
            let addr = DbgBreakPointWithStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgPrint" => {
            let addr = DbgPrint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgPrintEx" => {
            let addr = DbgPrintEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgPrintReturnControlC" => {
            let addr = DbgPrintReturnControlC as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgPrompt" => {
            let addr = DbgPrompt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgQueryDebugFilterState" => {
            let addr = DbgQueryDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgSetDebugFilterState" => {
            let addr = DbgSetDebugFilterState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "DbgSetDebugPrintCallback" => {
            let addr = DbgSetDebugPrintCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireFastMutex" => {
            let addr = ExAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireFastMutexUnsafe" => {
            let addr = ExAcquireFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquirePushLockExclusiveEx" => {
            let addr = ExAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquirePushLockSharedEx" => {
            let addr = ExAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireResourceExclusiveLite" => {
            let addr = ExAcquireResourceExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireResourceSharedLite" => {
            let addr = ExAcquireResourceSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireRundownProtection" => {
            let addr = ExAcquireRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireRundownProtectionCacheAware" => {
            let addr = ExAcquireRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireRundownProtectionCacheAwareEx" => {
            let addr = ExAcquireRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireRundownProtectionEx" => {
            let addr = ExAcquireRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSharedStarveExclusive" => {
            let addr = ExAcquireSharedStarveExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSharedWaitForExclusive" => {
            let addr = ExAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSpinLockExclusive" => {
            let addr = ExAcquireSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSpinLockShared" => {
            let addr = ExAcquireSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAdjustLookasideDepth" => {
            let addr = ExAdjustLookasideDepth as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocateCacheAwareRundownProtection" => {
            let addr = ExAllocateCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocateFromLookasideListEx" => {
            let addr = ExAllocateFromLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocateFromNPagedLookasideList" => {
            let addr = ExAllocateFromNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocateFromPagedLookasideList" => {
            let addr = ExAllocateFromPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocatePool2" => {
            let addr = ExAllocatePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocatePool3" => {
            let addr = ExAllocatePool3 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocatePoolWithQuota" => {
            let addr = ExAllocatePoolWithQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExAllocateTimer" => {
            let addr = ExAllocateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExCancelTimer" => {
            let addr = ExCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExCleanupRundownProtectionCacheAware" => {
            let addr = ExCleanupRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExConvertExclusiveToSharedLite" => {
            let addr = ExConvertExclusiveToSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExCreateCallback" => {
            let addr = ExCreateCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExCreatePool" => {
            let addr = ExCreatePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDeleteLookasideListEx" => {
            let addr = ExDeleteLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDeleteNPagedLookasideList" => {
            let addr = ExDeleteNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDeletePagedLookasideList" => {
            let addr = ExDeletePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDeleteResourceLite" => {
            let addr = ExDeleteResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDeleteTimer" => {
            let addr = ExDeleteTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDestroyPool" => {
            let addr = ExDestroyPool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExDisableResourceBoostLite" => {
            let addr = ExDisableResourceBoostLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireResourceExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireResourceShared" => {
            let addr = ExEnterCriticalRegionAndAcquireResourceShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExEnterCriticalRegionAndAcquireSharedWaitForExclusive" => {
            let addr = ExEnterCriticalRegionAndAcquireSharedWaitForExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExEnumerateSystemFirmwareTables" => {
            let addr = ExEnumerateSystemFirmwareTables as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExExtendZone" => {
            let addr = ExExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFlushLookasideListEx" => {
            let addr = ExFlushLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreeCacheAwareRundownProtection" => {
            let addr = ExFreeCacheAwareRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreePool" => {
            let addr = ExFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreePool2" => {
            let addr = ExFreePool2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreePoolWithTag" => {
            let addr = ExFreePoolWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreeToLookasideListEx" => {
            let addr = ExFreeToLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreeToNPagedLookasideList" => {
            let addr = ExFreeToNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExFreeToPagedLookasideList" => {
            let addr = ExFreeToPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetExclusiveWaiterCount" => {
            let addr = ExGetExclusiveWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetFirmwareEnvironmentVariable" => {
            let addr = ExGetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetFirmwareType" => {
            let addr = ExGetFirmwareType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetPreviousMode" => {
            let addr = ExGetPreviousMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetSharedWaiterCount" => {
            let addr = ExGetSharedWaiterCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExGetSystemFirmwareTable" => {
            let addr = ExGetSystemFirmwareTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeDeviceAts" => {
            let addr = ExInitializeDeviceAts as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeLookasideListEx" => {
            let addr = ExInitializeLookasideListEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeNPagedLookasideList" => {
            let addr = ExInitializeNPagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializePagedLookasideList" => {
            let addr = ExInitializePagedLookasideList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializePushLock" => {
            let addr = ExInitializePushLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeResourceLite" => {
            let addr = ExInitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeRundownProtection" => {
            let addr = ExInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeRundownProtectionCacheAware" => {
            let addr = ExInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeRundownProtectionCacheAwareEx" => {
            let addr = ExInitializeRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInitializeZone" => {
            let addr = ExInitializeZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedAddLargeInteger" => {
            let addr = ExInterlockedAddLargeInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedAddUlong" => {
            let addr = ExInterlockedAddUlong as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedExtendZone" => {
            let addr = ExInterlockedExtendZone as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedInsertHeadList" => {
            let addr = ExInterlockedInsertHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedInsertTailList" => {
            let addr = ExInterlockedInsertTailList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedPopEntryList" => {
            let addr = ExInterlockedPopEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedPushEntryList" => {
            let addr = ExInterlockedPushEntryList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExInterlockedRemoveHeadList" => {
            let addr = ExInterlockedRemoveHeadList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExIsManufacturingModeEnabled" => {
            let addr = ExIsManufacturingModeEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExIsProcessorFeaturePresent" => {
            let addr = ExIsProcessorFeaturePresent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExIsResourceAcquiredExclusiveLite" => {
            let addr = ExIsResourceAcquiredExclusiveLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExIsResourceAcquiredSharedLite" => {
            let addr = ExIsResourceAcquiredSharedLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExIsSoftBoot" => {
            let addr = ExIsSoftBoot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExLocalTimeToSystemTime" => {
            let addr = ExLocalTimeToSystemTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExNotifyCallback" => {
            let addr = ExNotifyCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExQueryDepthSList" => {
            let addr = ExQueryDepthSList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExQueryPoolBlockSize" => {
            let addr = ExQueryPoolBlockSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExQueryTimerResolution" => {
            let addr = ExQueryTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExQueueWorkItem" => {
            let addr = ExQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRaiseAccessViolation" => {
            let addr = ExRaiseAccessViolation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRaiseDatatypeMisalignment" => {
            let addr = ExRaiseDatatypeMisalignment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRaiseStatus" => {
            let addr = ExRaiseStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRcuFreePool" => {
            let addr = ExRcuFreePool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReInitializeRundownProtection" => {
            let addr = ExReInitializeRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReInitializeRundownProtectionCacheAware" => {
            let addr = ExReInitializeRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRegisterCallback" => {
            let addr = ExRegisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReinitializeResourceLite" => {
            let addr = ExReinitializeResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseFastMutex" => {
            let addr = ExReleaseFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseFastMutexUnsafe" => {
            let addr = ExReleaseFastMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleasePushLockExclusiveEx" => {
            let addr = ExReleasePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleasePushLockSharedEx" => {
            let addr = ExReleasePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseResourceAndLeaveCriticalRegion" => {
            let addr = ExReleaseResourceAndLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseResourceForThreadLite" => {
            let addr = ExReleaseResourceForThreadLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseResourceLite" => {
            let addr = ExReleaseResourceLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseRundownProtection" => {
            let addr = ExReleaseRundownProtection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseRundownProtectionCacheAware" => {
            let addr = ExReleaseRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseRundownProtectionCacheAwareEx" => {
            let addr = ExReleaseRundownProtectionCacheAwareEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseRundownProtectionEx" => {
            let addr = ExReleaseRundownProtectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseSpinLockExclusive" => {
            let addr = ExReleaseSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseSpinLockExclusiveFromDpcLevel" => {
            let addr = ExReleaseSpinLockExclusiveFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseSpinLockShared" => {
            let addr = ExReleaseSpinLockShared as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExReleaseSpinLockSharedFromDpcLevel" => {
            let addr = ExReleaseSpinLockSharedFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRundownCompleted" => {
            let addr = ExRundownCompleted as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExRundownCompletedCacheAware" => {
            let addr = ExRundownCompletedCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSecurePoolUpdate" => {
            let addr = ExSecurePoolUpdate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSecurePoolValidate" => {
            let addr = ExSecurePoolValidate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSetFirmwareEnvironmentVariable" => {
            let addr = ExSetFirmwareEnvironmentVariable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSetResourceOwnerPointer" => {
            let addr = ExSetResourceOwnerPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSetResourceOwnerPointerEx" => {
            let addr = ExSetResourceOwnerPointerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSetTimer" => {
            let addr = ExSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSetTimerResolution" => {
            let addr = ExSetTimerResolution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSizeOfRundownProtectionCacheAware" => {
            let addr = ExSizeOfRundownProtectionCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExSystemTimeToLocalTime" => {
            let addr = ExSystemTimeToLocalTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryAcquirePushLockExclusiveEx" => {
            let addr = ExTryAcquirePushLockExclusiveEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryAcquirePushLockSharedEx" => {
            let addr = ExTryAcquirePushLockSharedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryAcquireSpinLockExclusiveAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockExclusiveAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryAcquireSpinLockSharedAtDpcLevel" => {
            let addr = ExTryAcquireSpinLockSharedAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryConvertSharedSpinLockExclusive" => {
            let addr = ExTryConvertSharedSpinLockExclusive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExTryToAcquireFastMutex" => {
            let addr = ExTryToAcquireFastMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExUnregisterCallback" => {
            let addr = ExUnregisterCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExUuidCreate" => {
            let addr = ExUuidCreate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExVerifySuite" => {
            let addr = ExVerifySuite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExWaitForRundownProtectionRelease" => {
            let addr = ExWaitForRundownProtectionRelease as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExWaitForRundownProtectionReleaseCacheAware" => {
            let addr = ExWaitForRundownProtectionReleaseCacheAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExpInterlockedFlushSList" => {
            let addr = ExpInterlockedFlushSList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExpInterlockedPopEntrySList" => {
            let addr = ExpInterlockedPopEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExpInterlockedPushEntrySList" => {
            let addr = ExpInterlockedPushEntrySList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ExportSecurityContext" => {
            let addr = ExportSecurityContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAcquireCancelSpinLock" => {
            let addr = IoAcquireCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAcquireKsrPersistentMemory" => {
            let addr = IoAcquireKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAcquireKsrPersistentMemoryEx" => {
            let addr = IoAcquireKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAcquireRemoveLockEx" => {
            let addr = IoAcquireRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAcquireVpbSpinLock" => {
            let addr = IoAcquireVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateController" => {
            let addr = IoAllocateController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateDriverObjectExtension" => {
            let addr = IoAllocateDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateErrorLogEntry" => {
            let addr = IoAllocateErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateIrp" => {
            let addr = IoAllocateIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateIrpEx" => {
            let addr = IoAllocateIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateMdl" => {
            let addr = IoAllocateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateMiniCompletionPacket" => {
            let addr = IoAllocateMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateSfioStreamIdentifier" => {
            let addr = IoAllocateSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAllocateWorkItem" => {
            let addr = IoAllocateWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoApplyPriorityInfoThread" => {
            let addr = IoApplyPriorityInfoThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAssignResources" => {
            let addr = IoAssignResources as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAttachDevice" => {
            let addr = IoAttachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAttachDeviceByPointer" => {
            let addr = IoAttachDeviceByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAttachDeviceToDeviceStack" => {
            let addr = IoAttachDeviceToDeviceStack as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoAttachDeviceToDeviceStackSafe" => {
            let addr = IoAttachDeviceToDeviceStackSafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoBuildAsynchronousFsdRequest" => {
            let addr = IoBuildAsynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoBuildDeviceIoControlRequest" => {
            let addr = IoBuildDeviceIoControlRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoBuildPartialMdl" => {
            let addr = IoBuildPartialMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoBuildSynchronousFsdRequest" => {
            let addr = IoBuildSynchronousFsdRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCancelFileOpen" => {
            let addr = IoCancelFileOpen as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCancelIrp" => {
            let addr = IoCancelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckDesiredAccess" => {
            let addr = IoCheckDesiredAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckEaBufferValidity" => {
            let addr = IoCheckEaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckFileObjectOpenedAsCopyDestination" => {
            let addr = IoCheckFileObjectOpenedAsCopyDestination as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckFileObjectOpenedAsCopySource" => {
            let addr = IoCheckFileObjectOpenedAsCopySource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckFunctionAccess" => {
            let addr = IoCheckFunctionAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckLinkShareAccess" => {
            let addr = IoCheckLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckQuerySetFileInformation" => {
            let addr = IoCheckQuerySetFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckQuerySetVolumeInformation" => {
            let addr = IoCheckQuerySetVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckQuotaBufferValidity" => {
            let addr = IoCheckQuotaBufferValidity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckShareAccess" => {
            let addr = IoCheckShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCheckShareAccessEx" => {
            let addr = IoCheckShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCleanupIrp" => {
            let addr = IoCleanupIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoClearActivityIdThread" => {
            let addr = IoClearActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoClearFsTrackOffsetState" => {
            let addr = IoClearFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoClearIrpExtraCreateParameter" => {
            let addr = IoClearIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoConnectInterrupt" => {
            let addr = IoConnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoConnectInterruptEx" => {
            let addr = IoConnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateController" => {
            let addr = IoCreateController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateDevice" => {
            let addr = IoCreateDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateDisk" => {
            let addr = IoCreateDisk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateDriverProxyExtension" => {
            let addr = IoCreateDriverProxyExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateFile" => {
            let addr = IoCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateFileEx" => {
            let addr = IoCreateFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateFileSpecifyDeviceObjectHint" => {
            let addr = IoCreateFileSpecifyDeviceObjectHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateNotificationEvent" => {
            let addr = IoCreateNotificationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateStreamFileObject" => {
            let addr = IoCreateStreamFileObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateStreamFileObjectEx" => {
            let addr = IoCreateStreamFileObjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateStreamFileObjectEx2" => {
            let addr = IoCreateStreamFileObjectEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateStreamFileObjectLite" => {
            let addr = IoCreateStreamFileObjectLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateSymbolicLink" => {
            let addr = IoCreateSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateSynchronizationEvent" => {
            let addr = IoCreateSynchronizationEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateSystemThread" => {
            let addr = IoCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCreateUnprotectedSymbolicLink" => {
            let addr = IoCreateUnprotectedSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqInitialize" => {
            let addr = IoCsqInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqInitializeEx" => {
            let addr = IoCsqInitializeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqInsertIrp" => {
            let addr = IoCsqInsertIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqInsertIrpEx" => {
            let addr = IoCsqInsertIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqRemoveIrp" => {
            let addr = IoCsqRemoveIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoCsqRemoveNextIrp" => {
            let addr = IoCsqRemoveNextIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDecrementKeepAliveCount" => {
            let addr = IoDecrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDeleteController" => {
            let addr = IoDeleteController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDeleteDevice" => {
            let addr = IoDeleteDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDeleteSymbolicLink" => {
            let addr = IoDeleteSymbolicLink as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDetachDevice" => {
            let addr = IoDetachDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDisconnectInterrupt" => {
            let addr = IoDisconnectInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDisconnectInterruptEx" => {
            let addr = IoDisconnectInterruptEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoDriverProxyCreateHotSwappableWorkerThread" => {
            let addr = IoDriverProxyCreateHotSwappableWorkerThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoEnumerateDeviceObjectList" => {
            let addr = IoEnumerateDeviceObjectList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoEnumerateKsrPersistentMemoryEx" => {
            let addr = IoEnumerateKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoEnumerateRegisteredFiltersList" => {
            let addr = IoEnumerateRegisteredFiltersList as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFastQueryNetworkAttributes" => {
            let addr = IoFastQueryNetworkAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoForwardIrpSynchronously" => {
            let addr = IoForwardIrpSynchronously as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeController" => {
            let addr = IoFreeController as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeErrorLogEntry" => {
            let addr = IoFreeErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeIrp" => {
            let addr = IoFreeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeKsrPersistentMemory" => {
            let addr = IoFreeKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeMdl" => {
            let addr = IoFreeMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeMiniCompletionPacket" => {
            let addr = IoFreeMiniCompletionPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeSfioStreamIdentifier" => {
            let addr = IoFreeSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoFreeWorkItem" => {
            let addr = IoFreeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetActivityIdIrp" => {
            let addr = IoGetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetActivityIdThread" => {
            let addr = IoGetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetAffinityInterrupt" => {
            let addr = IoGetAffinityInterrupt as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetAttachedDevice" => {
            let addr = IoGetAttachedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetAttachedDeviceReference" => {
            let addr = IoGetAttachedDeviceReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetBaseFileSystemDeviceObject" => {
            let addr = IoGetBaseFileSystemDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetBootDiskInformation" => {
            let addr = IoGetBootDiskInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetBootDiskInformationLite" => {
            let addr = IoGetBootDiskInformationLite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetConfigurationInformation" => {
            let addr = IoGetConfigurationInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetContainerInformation" => {
            let addr = IoGetContainerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetCopyInformationExtension" => {
            let addr = IoGetCopyInformationExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetCurrentProcess" => {
            let addr = IoGetCurrentProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceAttachmentBaseRef" => {
            let addr = IoGetDeviceAttachmentBaseRef as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceDirectory" => {
            let addr = IoGetDeviceDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceInterfaceAlias" => {
            let addr = IoGetDeviceInterfaceAlias as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceInterfacePropertyData" => {
            let addr = IoGetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceInterfaces" => {
            let addr = IoGetDeviceInterfaces as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceNumaNode" => {
            let addr = IoGetDeviceNumaNode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceObjectPointer" => {
            let addr = IoGetDeviceObjectPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceProperty" => {
            let addr = IoGetDeviceProperty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDevicePropertyData" => {
            let addr = IoGetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDeviceToVerify" => {
            let addr = IoGetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDiskDeviceObject" => {
            let addr = IoGetDiskDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDmaAdapter" => {
            let addr = IoGetDmaAdapter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverDirectory" => {
            let addr = IoGetDriverDirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverObjectExtension" => {
            let addr = IoGetDriverObjectExtension as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverProxyEndpointWrapper" => {
            let addr = IoGetDriverProxyEndpointWrapper as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverProxyExtensionFromDriverObject" => {
            let addr = IoGetDriverProxyExtensionFromDriverObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverProxyExtensionVersion" => {
            let addr = IoGetDriverProxyExtensionVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetDriverProxyFeatures" => {
            let addr = IoGetDriverProxyFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetFileObjectGenericMapping" => {
            let addr = IoGetFileObjectGenericMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetFsTrackOffsetState" => {
            let addr = IoGetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetFsZeroingOffset" => {
            let addr = IoGetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetInitialStack" => {
            let addr = IoGetInitialStack as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetInitiatorProcess" => {
            let addr = IoGetInitiatorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetIoAttributionHandle" => {
            let addr = IoGetIoAttributionHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetIoPriorityHint" => {
            let addr = IoGetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetIommuInterface" => {
            let addr = IoGetIommuInterface as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetIommuInterfaceEx" => {
            let addr = IoGetIommuInterfaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetIrpExtraCreateParameter" => {
            let addr = IoGetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetKsrPersistentMemoryBuffer" => {
            let addr = IoGetKsrPersistentMemoryBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetLowerDeviceObject" => {
            let addr = IoGetLowerDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetOplockKeyContext" => {
            let addr = IoGetOplockKeyContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetOplockKeyContextEx" => {
            let addr = IoGetOplockKeyContextEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetPagingIoPriority" => {
            let addr = IoGetPagingIoPriority as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetRelatedDeviceObject" => {
            let addr = IoGetRelatedDeviceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetRequestorProcess" => {
            let addr = IoGetRequestorProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetRequestorProcessId" => {
            let addr = IoGetRequestorProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetRequestorSessionId" => {
            let addr = IoGetRequestorSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetSfioStreamIdentifier" => {
            let addr = IoGetSfioStreamIdentifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetShadowFileInformation" => {
            let addr = IoGetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetSilo" => {
            let addr = IoGetSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetSiloParameters" => {
            let addr = IoGetSiloParameters as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetStackLimits" => {
            let addr = IoGetStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetTopLevelIrp" => {
            let addr = IoGetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoGetTransactionParameterBlock" => {
            let addr = IoGetTransactionParameterBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoHotSwapDriverProxyEndpoints" => {
            let addr = IoHotSwapDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIncrementKeepAliveCount" => {
            let addr = IoIncrementKeepAliveCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInitializeIrp" => {
            let addr = IoInitializeIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInitializeIrpEx" => {
            let addr = IoInitializeIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInitializeRemoveLockEx" => {
            let addr = IoInitializeRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInitializeTimer" => {
            let addr = IoInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInitializeWorkItem" => {
            let addr = IoInitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInvalidateDeviceRelations" => {
            let addr = IoInvalidateDeviceRelations as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoInvalidateDeviceState" => {
            let addr = IoInvalidateDeviceState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIrpHasFsTrackOffsetExtensionType" => {
            let addr = IoIrpHasFsTrackOffsetExtensionType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIs32bitProcess" => {
            let addr = IoIs32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsFileObjectIgnoringSharing" => {
            let addr = IoIsFileObjectIgnoringSharing as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsFileOriginRemote" => {
            let addr = IoIsFileOriginRemote as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsInitiator32bitProcess" => {
            let addr = IoIsInitiator32bitProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsOperationSynchronous" => {
            let addr = IoIsOperationSynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsSystemThread" => {
            let addr = IoIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsValidIrpStatus" => {
            let addr = IoIsValidIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsValidNameGraftingBuffer" => {
            let addr = IoIsValidNameGraftingBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoIsWdmVersionAvailable" => {
            let addr = IoIsWdmVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoMakeAssociatedIrp" => {
            let addr = IoMakeAssociatedIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoMakeAssociatedIrpEx" => {
            let addr = IoMakeAssociatedIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoMapKsrPersistentMemoryEx" => {
            let addr = IoMapKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoOpenDeviceInterfaceRegistryKey" => {
            let addr = IoOpenDeviceInterfaceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoOpenDeviceRegistryKey" => {
            let addr = IoOpenDeviceRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoOpenDriverRegistryKey" => {
            let addr = IoOpenDriverRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoPageRead" => {
            let addr = IoPageRead as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoPropagateActivityIdToThread" => {
            let addr = IoPropagateActivityIdToThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryDeviceDescription" => {
            let addr = IoQueryDeviceDescription as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryDmaFeatureSupport" => {
            let addr = IoQueryDmaFeatureSupport as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryFileDosDeviceName" => {
            let addr = IoQueryFileDosDeviceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryFileInformation" => {
            let addr = IoQueryFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryFullDriverPath" => {
            let addr = IoQueryFullDriverPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryInformationByName" => {
            let addr = IoQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryKsrPersistentMemorySize" => {
            let addr = IoQueryKsrPersistentMemorySize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryKsrPersistentMemorySizeEx" => {
            let addr = IoQueryKsrPersistentMemorySizeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueryVolumeInformation" => {
            let addr = IoQueryVolumeInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueueThreadIrp" => {
            let addr = IoQueueThreadIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueueWorkItem" => {
            let addr = IoQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoQueueWorkItemEx" => {
            let addr = IoQueueWorkItemEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRaiseHardError" => {
            let addr = IoRaiseHardError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRaiseInformationalHardError" => {
            let addr = IoRaiseInformationalHardError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReadDiskSignature" => {
            let addr = IoReadDiskSignature as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReadPartitionTable" => {
            let addr = IoReadPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReadPartitionTableEx" => {
            let addr = IoReadPartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRecordIoAttribution" => {
            let addr = IoRecordIoAttribution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterBootDriverCallback" => {
            let addr = IoRegisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterBootDriverReinitialization" => {
            let addr = IoRegisterBootDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterContainerNotification" => {
            let addr = IoRegisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterDeviceInterface" => {
            let addr = IoRegisterDeviceInterface as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterDriverProxyEndpoints" => {
            let addr = IoRegisterDriverProxyEndpoints as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterDriverReinitialization" => {
            let addr = IoRegisterDriverReinitialization as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterFileSystem" => {
            let addr = IoRegisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterFsRegistrationChange" => {
            let addr = IoRegisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterFsRegistrationChangeMountAware" => {
            let addr = IoRegisterFsRegistrationChangeMountAware as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterLastChanceShutdownNotification" => {
            let addr = IoRegisterLastChanceShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterPlugPlayNotification" => {
            let addr = IoRegisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRegisterShutdownNotification" => {
            let addr = IoRegisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReleaseCancelSpinLock" => {
            let addr = IoReleaseCancelSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReleaseRemoveLockAndWaitEx" => {
            let addr = IoReleaseRemoveLockAndWaitEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReleaseRemoveLockEx" => {
            let addr = IoReleaseRemoveLockEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReleaseVpbSpinLock" => {
            let addr = IoReleaseVpbSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRemoveIoCompletion" => {
            let addr = IoRemoveIoCompletion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRemoveLinkShareAccess" => {
            let addr = IoRemoveLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRemoveLinkShareAccessEx" => {
            let addr = IoRemoveLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRemoveShareAccess" => {
            let addr = IoRemoveShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReplaceFileObjectName" => {
            let addr = IoReplaceFileObjectName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReplacePartitionUnit" => {
            let addr = IoReplacePartitionUnit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportDetectedDevice" => {
            let addr = IoReportDetectedDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportInterruptActive" => {
            let addr = IoReportInterruptActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportInterruptInactive" => {
            let addr = IoReportInterruptInactive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportResourceForDetection" => {
            let addr = IoReportResourceForDetection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportResourceUsage" => {
            let addr = IoReportResourceUsage as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportRootDevice" => {
            let addr = IoReportRootDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportTargetDeviceChange" => {
            let addr = IoReportTargetDeviceChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReportTargetDeviceChangeAsynchronous" => {
            let addr = IoReportTargetDeviceChangeAsynchronous as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRequestDeviceEject" => {
            let addr = IoRequestDeviceEject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRequestDeviceEjectEx" => {
            let addr = IoRequestDeviceEjectEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRequestDeviceRemovalForReset" => {
            let addr = IoRequestDeviceRemovalForReset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReserveKsrPersistentMemory" => {
            let addr = IoReserveKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReserveKsrPersistentMemoryEx" => {
            let addr = IoReserveKsrPersistentMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoRetrievePriorityInfo" => {
            let addr = IoRetrievePriorityInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoReuseIrp" => {
            let addr = IoReuseIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetActivityIdIrp" => {
            let addr = IoSetActivityIdIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetActivityIdThread" => {
            let addr = IoSetActivityIdThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetCompletionRoutineEx" => {
            let addr = IoSetCompletionRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetDeviceInterfacePropertyData" => {
            let addr = IoSetDeviceInterfacePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetDeviceInterfaceState" => {
            let addr = IoSetDeviceInterfaceState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetDevicePropertyData" => {
            let addr = IoSetDevicePropertyData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetDeviceToVerify" => {
            let addr = IoSetDeviceToVerify as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetFileObjectIgnoreSharing" => {
            let addr = IoSetFileObjectIgnoreSharing as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetFileOrigin" => {
            let addr = IoSetFileOrigin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetFsTrackOffsetState" => {
            let addr = IoSetFsTrackOffsetState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetFsZeroingOffset" => {
            let addr = IoSetFsZeroingOffset as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetFsZeroingOffsetRequired" => {
            let addr = IoSetFsZeroingOffsetRequired as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetHardErrorOrVerifyDevice" => {
            let addr = IoSetHardErrorOrVerifyDevice as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetInformation" => {
            let addr = IoSetInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetIoAttributionIrp" => {
            let addr = IoSetIoAttributionIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetIoCompletionEx" => {
            let addr = IoSetIoCompletionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetIoPriorityHint" => {
            let addr = IoSetIoPriorityHint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetIrpExtraCreateParameter" => {
            let addr = IoSetIrpExtraCreateParameter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetLinkShareAccess" => {
            let addr = IoSetLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetMasterIrpStatus" => {
            let addr = IoSetMasterIrpStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetPartitionInformation" => {
            let addr = IoSetPartitionInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetPartitionInformationEx" => {
            let addr = IoSetPartitionInformationEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetShadowFileInformation" => {
            let addr = IoSetShadowFileInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetShareAccess" => {
            let addr = IoSetShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetShareAccessEx" => {
            let addr = IoSetShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetStartIoAttributes" => {
            let addr = IoSetStartIoAttributes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetSystemPartition" => {
            let addr = IoSetSystemPartition as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetThreadHardErrorMode" => {
            let addr = IoSetThreadHardErrorMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSetTopLevelIrp" => {
            let addr = IoSetTopLevelIrp as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSizeOfIrpEx" => {
            let addr = IoSizeOfIrpEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSizeofWorkItem" => {
            let addr = IoSizeofWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoStartNextPacket" => {
            let addr = IoStartNextPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoStartNextPacketByKey" => {
            let addr = IoStartNextPacketByKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoStartPacket" => {
            let addr = IoStartPacket as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoStartTimer" => {
            let addr = IoStartTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoStopTimer" => {
            let addr = IoStopTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSynchronousCallDriver" => {
            let addr = IoSynchronousCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoSynchronousPageWrite" => {
            let addr = IoSynchronousPageWrite as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoThreadToProcess" => {
            let addr = IoThreadToProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoTransferActivityId" => {
            let addr = IoTransferActivityId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoTranslateBusAddress" => {
            let addr = IoTranslateBusAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoTryQueueWorkItem" => {
            let addr = IoTryQueueWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUninitializeWorkItem" => {
            let addr = IoUninitializeWorkItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterBootDriverCallback" => {
            let addr = IoUnregisterBootDriverCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterContainerNotification" => {
            let addr = IoUnregisterContainerNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterFileSystem" => {
            let addr = IoUnregisterFileSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterFsRegistrationChange" => {
            let addr = IoUnregisterFsRegistrationChange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterPlugPlayNotification" => {
            let addr = IoUnregisterPlugPlayNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterPlugPlayNotificationEx" => {
            let addr = IoUnregisterPlugPlayNotificationEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUnregisterShutdownNotification" => {
            let addr = IoUnregisterShutdownNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUpdateLinkShareAccess" => {
            let addr = IoUpdateLinkShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUpdateLinkShareAccessEx" => {
            let addr = IoUpdateLinkShareAccessEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoUpdateShareAccess" => {
            let addr = IoUpdateShareAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoValidateDeviceIoControlAccess" => {
            let addr = IoValidateDeviceIoControlAccess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVerifyPartitionTable" => {
            let addr = IoVerifyPartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVerifyVolume" => {
            let addr = IoVerifyVolume as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVolumeDeviceNameToGuid" => {
            let addr = IoVolumeDeviceNameToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVolumeDeviceNameToGuidPath" => {
            let addr = IoVolumeDeviceNameToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVolumeDeviceToDosName" => {
            let addr = IoVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVolumeDeviceToGuid" => {
            let addr = IoVolumeDeviceToGuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoVolumeDeviceToGuidPath" => {
            let addr = IoVolumeDeviceToGuidPath as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIAllocateInstanceIds" => {
            let addr = IoWMIAllocateInstanceIds as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIDeviceObjectToInstanceName" => {
            let addr = IoWMIDeviceObjectToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIDeviceObjectToProviderId" => {
            let addr = IoWMIDeviceObjectToProviderId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIExecuteMethod" => {
            let addr = IoWMIExecuteMethod as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIHandleToInstanceName" => {
            let addr = IoWMIHandleToInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIOpenBlock" => {
            let addr = IoWMIOpenBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIQueryAllData" => {
            let addr = IoWMIQueryAllData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIQueryAllDataMultiple" => {
            let addr = IoWMIQueryAllDataMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIQuerySingleInstance" => {
            let addr = IoWMIQuerySingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIQuerySingleInstanceMultiple" => {
            let addr = IoWMIQuerySingleInstanceMultiple as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIRegistrationControl" => {
            let addr = IoWMIRegistrationControl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMISetNotificationCallback" => {
            let addr = IoWMISetNotificationCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMISetSingleInstance" => {
            let addr = IoWMISetSingleInstance as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMISetSingleItem" => {
            let addr = IoWMISetSingleItem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMISuggestInstanceName" => {
            let addr = IoWMISuggestInstanceName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWMIWriteEvent" => {
            let addr = IoWMIWriteEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWithinStackLimits" => {
            let addr = IoWithinStackLimits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWriteErrorLogEntry" => {
            let addr = IoWriteErrorLogEntry as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWriteKsrPersistentMemory" => {
            let addr = IoWriteKsrPersistentMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWritePartitionTable" => {
            let addr = IoWritePartitionTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IoWritePartitionTableEx" => {
            let addr = IoWritePartitionTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IofCallDriver" => {
            let addr = IofCallDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IofCompleteRequest" => {
            let addr = IofCompleteRequest as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "IofGetDriverProxyWrapperFromEndpoint" => {
            let addr = IofGetDriverProxyWrapperFromEndpoint as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireGuardedMutex" => {
            let addr = KeAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireGuardedMutexUnsafe" => {
            let addr = KeAcquireGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLock" => {
            let addr = KeAcquireInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLockAtDpcLevel" => {
            let addr = KeAcquireInStackQueuedSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireInStackQueuedSpinLockForDpc" => {
            let addr = KeAcquireInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireInterruptSpinLock" => {
            let addr = KeAcquireInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireQueuedSpinLock" => {
            let addr = KeAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireSpinLockAtDpcLevel" => {
            let addr = KeAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireSpinLockForDpc" => {
            let addr = KeAcquireSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireSpinLockRaiseToDpc" => {
            let addr = KeAcquireSpinLockRaiseToDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAcquireSpinLockRaiseToSynch" => {
            let addr = KeAcquireSpinLockRaiseToSynch as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAddTriageDumpDataBlock" => {
            let addr = KeAddTriageDumpDataBlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAreAllApcsDisabled" => {
            let addr = KeAreAllApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAreApcsDisabled" => {
            let addr = KeAreApcsDisabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeAttachProcess" => {
            let addr = KeAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeBugCheck" => {
            let addr = KeBugCheck as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeBugCheckEx" => {
            let addr = KeBugCheckEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeCancelTimer" => {
            let addr = KeCancelTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeClearEvent" => {
            let addr = KeClearEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeConvertAuxiliaryCounterToPerformanceCounter" => {
            let addr = KeConvertAuxiliaryCounterToPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeConvertPerformanceCounterToAuxiliaryCounter" => {
            let addr = KeConvertPerformanceCounterToAuxiliaryCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDelayExecutionThread" => {
            let addr = KeDelayExecutionThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDeregisterBoundCallback" => {
            let addr = KeDeregisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDeregisterBugCheckCallback" => {
            let addr = KeDeregisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDeregisterBugCheckReasonCallback" => {
            let addr = KeDeregisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDeregisterNmiCallback" => {
            let addr = KeDeregisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDeregisterProcessorChangeCallback" => {
            let addr = KeDeregisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeDetachProcess" => {
            let addr = KeDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeEnterCriticalRegion" => {
            let addr = KeEnterCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeEnterGuardedRegion" => {
            let addr = KeEnterGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeExpandKernelStackAndCallout" => {
            let addr = KeExpandKernelStackAndCallout as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeExpandKernelStackAndCalloutEx" => {
            let addr = KeExpandKernelStackAndCalloutEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeFlushIoBuffers" => {
            let addr = KeFlushIoBuffers as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeFlushQueuedDpcs" => {
            let addr = KeFlushQueuedDpcs as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeFlushWriteBuffer" => {
            let addr = KeFlushWriteBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetCurrentIrql" => {
            let addr = KeGetCurrentIrql as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetCurrentNodeNumber" => {
            let addr = KeGetCurrentNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetCurrentProcessorNumberEx" => {
            let addr = KeGetCurrentProcessorNumberEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetProcessorIndexFromNumber" => {
            let addr = KeGetProcessorIndexFromNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetProcessorNumberFromIndex" => {
            let addr = KeGetProcessorNumberFromIndex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeGetRecommendedSharedDataAlignment" => {
            let addr = KeGetRecommendedSharedDataAlignment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeCrashDumpHeader" => {
            let addr = KeInitializeCrashDumpHeader as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeDeviceQueue" => {
            let addr = KeInitializeDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeDpc" => {
            let addr = KeInitializeDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeEvent" => {
            let addr = KeInitializeEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeGuardedMutex" => {
            let addr = KeInitializeGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeMutant" => {
            let addr = KeInitializeMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeMutex" => {
            let addr = KeInitializeMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeQueue" => {
            let addr = KeInitializeQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeSemaphore" => {
            let addr = KeInitializeSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeSpinLock" => {
            let addr = KeInitializeSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeThreadedDpc" => {
            let addr = KeInitializeThreadedDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeTimer" => {
            let addr = KeInitializeTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeTimerEx" => {
            let addr = KeInitializeTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInitializeTriageDumpDataArray" => {
            let addr = KeInitializeTriageDumpDataArray as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInsertByKeyDeviceQueue" => {
            let addr = KeInsertByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInsertDeviceQueue" => {
            let addr = KeInsertDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInsertHeadQueue" => {
            let addr = KeInsertHeadQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInsertQueue" => {
            let addr = KeInsertQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInsertQueueDpc" => {
            let addr = KeInsertQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInvalidateAllCaches" => {
            let addr = KeInvalidateAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeInvalidateRangeAllCaches" => {
            let addr = KeInvalidateRangeAllCaches as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeIpiGenericCall" => {
            let addr = KeIpiGenericCall as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeIsExecutingDpc" => {
            let addr = KeIsExecutingDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeLeaveCriticalRegion" => {
            let addr = KeLeaveCriticalRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeLeaveGuardedRegion" => {
            let addr = KeLeaveGuardedRegion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeLowerIrql" => {
            let addr = KeLowerIrql as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KePulseEvent" => {
            let addr = KePulseEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryActiveGroupCount" => {
            let addr = KeQueryActiveGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryActiveProcessorCount" => {
            let addr = KeQueryActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryActiveProcessorCountEx" => {
            let addr = KeQueryActiveProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryActiveProcessors" => {
            let addr = KeQueryActiveProcessors as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryAuxiliaryCounterFrequency" => {
            let addr = KeQueryAuxiliaryCounterFrequency as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryDpcWatchdogInformation" => {
            let addr = KeQueryDpcWatchdogInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryGroupAffinity" => {
            let addr = KeQueryGroupAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryHardwareCounterConfiguration" => {
            let addr = KeQueryHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryHighestNodeNumber" => {
            let addr = KeQueryHighestNodeNumber as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryInterruptTimePrecise" => {
            let addr = KeQueryInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryLogicalProcessorRelationship" => {
            let addr = KeQueryLogicalProcessorRelationship as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryMaximumGroupCount" => {
            let addr = KeQueryMaximumGroupCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryMaximumProcessorCount" => {
            let addr = KeQueryMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryMaximumProcessorCountEx" => {
            let addr = KeQueryMaximumProcessorCountEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryNodeActiveAffinity" => {
            let addr = KeQueryNodeActiveAffinity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryNodeActiveAffinity2" => {
            let addr = KeQueryNodeActiveAffinity2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryNodeActiveProcessorCount" => {
            let addr = KeQueryNodeActiveProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryNodeMaximumProcessorCount" => {
            let addr = KeQueryNodeMaximumProcessorCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryOwnerMutant" => {
            let addr = KeQueryOwnerMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryPerformanceCounter" => {
            let addr = KeQueryPerformanceCounter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryPriorityThread" => {
            let addr = KeQueryPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryRuntimeThread" => {
            let addr = KeQueryRuntimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQuerySystemTimePrecise" => {
            let addr = KeQuerySystemTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryTimeIncrement" => {
            let addr = KeQueryTimeIncrement as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryTotalCycleTimeThread" => {
            let addr = KeQueryTotalCycleTimeThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryUnbiasedInterruptTime" => {
            let addr = KeQueryUnbiasedInterruptTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeQueryUnbiasedInterruptTimePrecise" => {
            let addr = KeQueryUnbiasedInterruptTimePrecise as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRcuReadLock" => {
            let addr = KeRcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRcuReadLockAtDpcLevel" => {
            let addr = KeRcuReadLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRcuReadUnlock" => {
            let addr = KeRcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRcuSynchronize" => {
            let addr = KeRcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateEvent" => {
            let addr = KeReadStateEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateMutant" => {
            let addr = KeReadStateMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateMutex" => {
            let addr = KeReadStateMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateQueue" => {
            let addr = KeReadStateQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateSemaphore" => {
            let addr = KeReadStateSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReadStateTimer" => {
            let addr = KeReadStateTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRegisterBoundCallback" => {
            let addr = KeRegisterBoundCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRegisterBugCheckCallback" => {
            let addr = KeRegisterBugCheckCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRegisterBugCheckReasonCallback" => {
            let addr = KeRegisterBugCheckReasonCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRegisterNmiCallback" => {
            let addr = KeRegisterNmiCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRegisterProcessorChangeCallback" => {
            let addr = KeRegisterProcessorChangeCallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseGuardedMutex" => {
            let addr = KeReleaseGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseGuardedMutexUnsafe" => {
            let addr = KeReleaseGuardedMutexUnsafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLock" => {
            let addr = KeReleaseInStackQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLockForDpc" => {
            let addr = KeReleaseInStackQueuedSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseInStackQueuedSpinLockFromDpcLevel" => {
            let addr = KeReleaseInStackQueuedSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseInterruptSpinLock" => {
            let addr = KeReleaseInterruptSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseMutant" => {
            let addr = KeReleaseMutant as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseMutex" => {
            let addr = KeReleaseMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseQueuedSpinLock" => {
            let addr = KeReleaseQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseSemaphore" => {
            let addr = KeReleaseSemaphore as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseSpinLock" => {
            let addr = KeReleaseSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseSpinLockForDpc" => {
            let addr = KeReleaseSpinLockForDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeReleaseSpinLockFromDpcLevel" => {
            let addr = KeReleaseSpinLockFromDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveByKeyDeviceQueue" => {
            let addr = KeRemoveByKeyDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveByKeyDeviceQueueIfBusy" => {
            let addr = KeRemoveByKeyDeviceQueueIfBusy as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveDeviceQueue" => {
            let addr = KeRemoveDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveEntryDeviceQueue" => {
            let addr = KeRemoveEntryDeviceQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveQueue" => {
            let addr = KeRemoveQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveQueueDpc" => {
            let addr = KeRemoveQueueDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveQueueDpcEx" => {
            let addr = KeRemoveQueueDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRemoveQueueEx" => {
            let addr = KeRemoveQueueEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeResetEvent" => {
            let addr = KeResetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRestoreExtendedProcessorState" => {
            let addr = KeRestoreExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRevertToUserAffinityThread" => {
            let addr = KeRevertToUserAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRevertToUserAffinityThreadEx" => {
            let addr = KeRevertToUserAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRevertToUserGroupAffinityThread" => {
            let addr = KeRevertToUserGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeRundownQueue" => {
            let addr = KeRundownQueue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSaveExtendedProcessorState" => {
            let addr = KeSaveExtendedProcessorState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetBasePriorityThread" => {
            let addr = KeSetBasePriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetCoalescableTimer" => {
            let addr = KeSetCoalescableTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetEvent" => {
            let addr = KeSetEvent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetHardwareCounterConfiguration" => {
            let addr = KeSetHardwareCounterConfiguration as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetIdealProcessorThread" => {
            let addr = KeSetIdealProcessorThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetImportanceDpc" => {
            let addr = KeSetImportanceDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetKernelStackSwapEnable" => {
            let addr = KeSetKernelStackSwapEnable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetPriorityThread" => {
            let addr = KeSetPriorityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetSystemAffinityThread" => {
            let addr = KeSetSystemAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetSystemAffinityThreadEx" => {
            let addr = KeSetSystemAffinityThreadEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetSystemGroupAffinityThread" => {
            let addr = KeSetSystemGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetTargetProcessorDpc" => {
            let addr = KeSetTargetProcessorDpc as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetTargetProcessorDpcEx" => {
            let addr = KeSetTargetProcessorDpcEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetTimer" => {
            let addr = KeSetTimer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSetTimerEx" => {
            let addr = KeSetTimerEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeShouldYieldProcessor" => {
            let addr = KeShouldYieldProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSrcuAllocate" => {
            let addr = KeSrcuAllocate as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSrcuFree" => {
            let addr = KeSrcuFree as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSrcuReadLock" => {
            let addr = KeSrcuReadLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSrcuReadUnlock" => {
            let addr = KeSrcuReadUnlock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSrcuSynchronize" => {
            let addr = KeSrcuSynchronize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeStackAttachProcess" => {
            let addr = KeStackAttachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeStallExecutionProcessor" => {
            let addr = KeStallExecutionProcessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeSynchronizeExecution" => {
            let addr = KeSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeTestSpinLock" => {
            let addr = KeTestSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeTryToAcquireGuardedMutex" => {
            let addr = KeTryToAcquireGuardedMutex as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeTryToAcquireQueuedSpinLock" => {
            let addr = KeTryToAcquireQueuedSpinLock as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeTryToAcquireSpinLockAtDpcLevel" => {
            let addr = KeTryToAcquireSpinLockAtDpcLevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeUnstackDetachProcess" => {
            let addr = KeUnstackDetachProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeWaitForMultipleObjects" => {
            let addr = KeWaitForMultipleObjects as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "KeWaitForSingleObject" => {
            let addr = KeWaitForSingleObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAddPhysicalMemory" => {
            let addr = MmAddPhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAddVerifierSpecialThunks" => {
            let addr = MmAddVerifierSpecialThunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAddVerifierThunks" => {
            let addr = MmAddVerifierThunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAdvanceMdl" => {
            let addr = MmAdvanceMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateContiguousMemory" => {
            let addr = MmAllocateContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateContiguousMemoryEx" => {
            let addr = MmAllocateContiguousMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateContiguousMemorySpecifyCache" => {
            let addr = MmAllocateContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateContiguousMemorySpecifyCacheNode" => {
            let addr = MmAllocateContiguousMemorySpecifyCacheNode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateContiguousNodeMemory" => {
            let addr = MmAllocateContiguousNodeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateMappingAddress" => {
            let addr = MmAllocateMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateMappingAddressEx" => {
            let addr = MmAllocateMappingAddressEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateMdlForIoSpace" => {
            let addr = MmAllocateMdlForIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateNodePagesForMdlEx" => {
            let addr = MmAllocateNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocateNonCachedMemory" => {
            let addr = MmAllocateNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocatePagesForMdl" => {
            let addr = MmAllocatePagesForMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocatePagesForMdlEx" => {
            let addr = MmAllocatePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAllocatePartitionNodePagesForMdlEx" => {
            let addr = MmAllocatePartitionNodePagesForMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmAreMdlPagesCached" => {
            let addr = MmAreMdlPagesCached as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmBuildMdlForNonPagedPool" => {
            let addr = MmBuildMdlForNonPagedPool as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmCanFileBeTruncated" => {
            let addr = MmCanFileBeTruncated as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmCopyMemory" => {
            let addr = MmCopyMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmCreateMdl" => {
            let addr = MmCreateMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmCreateMirror" => {
            let addr = MmCreateMirror as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmDoesFileHaveUserWritableReferences" => {
            let addr = MmDoesFileHaveUserWritableReferences as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFlushImageSection" => {
            let addr = MmFlushImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmForceSectionClosed" => {
            let addr = MmForceSectionClosed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmForceSectionClosedEx" => {
            let addr = MmForceSectionClosedEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreeContiguousMemory" => {
            let addr = MmFreeContiguousMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreeContiguousMemorySpecifyCache" => {
            let addr = MmFreeContiguousMemorySpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreeMappingAddress" => {
            let addr = MmFreeMappingAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreeNonCachedMemory" => {
            let addr = MmFreeNonCachedMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreePagesFromMdl" => {
            let addr = MmFreePagesFromMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmFreePagesFromMdlEx" => {
            let addr = MmFreePagesFromMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetCacheAttribute" => {
            let addr = MmGetCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetCacheAttributeEx" => {
            let addr = MmGetCacheAttributeEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetMaximumFileSectionSize" => {
            let addr = MmGetMaximumFileSectionSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetModuleRoutineAddress" => {
            let addr = MmGetModuleRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetPhysicalAddress" => {
            let addr = MmGetPhysicalAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetPhysicalMemoryRanges" => {
            let addr = MmGetPhysicalMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetPhysicalMemoryRangesEx" => {
            let addr = MmGetPhysicalMemoryRangesEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetPhysicalMemoryRangesEx2" => {
            let addr = MmGetPhysicalMemoryRangesEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetSystemRoutineAddress" => {
            let addr = MmGetSystemRoutineAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmGetVirtualForPhysical" => {
            let addr = MmGetVirtualForPhysical as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsAddressValid" => {
            let addr = MmIsAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsDriverSuspectForVerifier" => {
            let addr = MmIsDriverSuspectForVerifier as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsDriverVerifying" => {
            let addr = MmIsDriverVerifying as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsDriverVerifyingByAddress" => {
            let addr = MmIsDriverVerifyingByAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsFileSectionActive" => {
            let addr = MmIsFileSectionActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsIoSpaceActive" => {
            let addr = MmIsIoSpaceActive as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsKernelAddress" => {
            let addr = MmIsKernelAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsNonPagedSystemAddressValid" => {
            let addr = MmIsNonPagedSystemAddressValid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsRecursiveIoFault" => {
            let addr = MmIsRecursiveIoFault as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsThisAnNtAsSystem" => {
            let addr = MmIsThisAnNtAsSystem as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsUserAddress" => {
            let addr = MmIsUserAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmIsVerifierEnabled" => {
            let addr = MmIsVerifierEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmLockPagableDataSection" => {
            let addr = MmLockPagableDataSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmLockPagableSectionByHandle" => {
            let addr = MmLockPagableSectionByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapIoSpace" => {
            let addr = MmMapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapIoSpaceEx" => {
            let addr = MmMapIoSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapLockedPages" => {
            let addr = MmMapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapLockedPagesSpecifyCache" => {
            let addr = MmMapLockedPagesSpecifyCache as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapLockedPagesWithReservedMapping" => {
            let addr = MmMapLockedPagesWithReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapMdl" => {
            let addr = MmMapMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapMemoryDumpMdlEx" => {
            let addr = MmMapMemoryDumpMdlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapUserAddressesToPage" => {
            let addr = MmMapUserAddressesToPage as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapVideoDisplay" => {
            let addr = MmMapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapViewInSessionSpace" => {
            let addr = MmMapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapViewInSessionSpaceEx" => {
            let addr = MmMapViewInSessionSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapViewInSystemSpace" => {
            let addr = MmMapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMapViewInSystemSpaceEx" => {
            let addr = MmMapViewInSystemSpaceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMdlPageContentsState" => {
            let addr = MmMdlPageContentsState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmMdlPagesAreZero" => {
            let addr = MmMdlPagesAreZero as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmPageEntireDriver" => {
            let addr = MmPageEntireDriver as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmPrefetchPages" => {
            let addr = MmPrefetchPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProbeAndLockPages" => {
            let addr = MmProbeAndLockPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProbeAndLockPagesEx" => {
            let addr = MmProbeAndLockPagesEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProbeAndLockProcessPages" => {
            let addr = MmProbeAndLockProcessPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProbeAndLockSelectedPages" => {
            let addr = MmProbeAndLockSelectedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProtectDriverSection" => {
            let addr = MmProtectDriverSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmProtectMdlSystemAddress" => {
            let addr = MmProtectMdlSystemAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmQuerySystemSize" => {
            let addr = MmQuerySystemSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmRemovePhysicalMemory" => {
            let addr = MmRemovePhysicalMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmResetDriverPaging" => {
            let addr = MmResetDriverPaging as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmRotatePhysicalView" => {
            let addr = MmRotatePhysicalView as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmSecureVirtualMemory" => {
            let addr = MmSecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmSecureVirtualMemoryEx" => {
            let addr = MmSecureVirtualMemoryEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmSetAddressRangeModified" => {
            let addr = MmSetAddressRangeModified as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmSetPermanentCacheAttribute" => {
            let addr = MmSetPermanentCacheAttribute as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmSizeOfMdl" => {
            let addr = MmSizeOfMdl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnlockPagableImageSection" => {
            let addr = MmUnlockPagableImageSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnlockPages" => {
            let addr = MmUnlockPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapIoSpace" => {
            let addr = MmUnmapIoSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapLockedPages" => {
            let addr = MmUnmapLockedPages as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapReservedMapping" => {
            let addr = MmUnmapReservedMapping as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapVideoDisplay" => {
            let addr = MmUnmapVideoDisplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapViewInSessionSpace" => {
            let addr = MmUnmapViewInSessionSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnmapViewInSystemSpace" => {
            let addr = MmUnmapViewInSystemSpace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "MmUnsecureVirtualMemory" => {
            let addr = MmUnsecureVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAccessCheckAndAuditAlarm" => {
            let addr = NtAccessCheckAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAccessCheckByTypeAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarm" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAccessCheckByTypeResultListAndAuditAlarmByHandle" => {
            let addr = NtAccessCheckByTypeResultListAndAuditAlarmByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAdjustGroupsToken" => {
            let addr = NtAdjustGroupsToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAdjustPrivilegesToken" => {
            let addr = NtAdjustPrivilegesToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtAllocateVirtualMemory" => {
            let addr = NtAllocateVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtClose" => {
            let addr = NtClose as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCloseObjectAuditAlarm" => {
            let addr = NtCloseObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCommitComplete" => {
            let addr = NtCommitComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCommitEnlistment" => {
            let addr = NtCommitEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCommitTransaction" => {
            let addr = NtCommitTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCopyFileChunk" => {
            let addr = NtCopyFileChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateEnlistment" => {
            let addr = NtCreateEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateFile" => {
            let addr = NtCreateFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateResourceManager" => {
            let addr = NtCreateResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateSection" => {
            let addr = NtCreateSection as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateSectionEx" => {
            let addr = NtCreateSectionEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateTransaction" => {
            let addr = NtCreateTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtCreateTransactionManager" => {
            let addr = NtCreateTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtDeleteObjectAuditAlarm" => {
            let addr = NtDeleteObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtDeviceIoControlFile" => {
            let addr = NtDeviceIoControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtDuplicateToken" => {
            let addr = NtDuplicateToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtEnumerateTransactionObject" => {
            let addr = NtEnumerateTransactionObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtFilterToken" => {
            let addr = NtFilterToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtFlushBuffersFileEx" => {
            let addr = NtFlushBuffersFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtFreeVirtualMemory" => {
            let addr = NtFreeVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtFsControlFile" => {
            let addr = NtFsControlFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtGetNotificationResourceManager" => {
            let addr = NtGetNotificationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtImpersonateAnonymousToken" => {
            let addr = NtImpersonateAnonymousToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtLockFile" => {
            let addr = NtLockFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtManagePartition" => {
            let addr = NtManagePartition as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenEnlistment" => {
            let addr = NtOpenEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenFile" => {
            let addr = NtOpenFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenJobObjectToken" => {
            let addr = NtOpenJobObjectToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenObjectAuditAlarm" => {
            let addr = NtOpenObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenProcess" => {
            let addr = NtOpenProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenProcessToken" => {
            let addr = NtOpenProcessToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenProcessTokenEx" => {
            let addr = NtOpenProcessTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenRegistryTransaction" => {
            let addr = NtOpenRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenResourceManager" => {
            let addr = NtOpenResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenThreadToken" => {
            let addr = NtOpenThreadToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenThreadTokenEx" => {
            let addr = NtOpenThreadTokenEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenTransaction" => {
            let addr = NtOpenTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtOpenTransactionManager" => {
            let addr = NtOpenTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPowerInformation" => {
            let addr = NtPowerInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrePrepareComplete" => {
            let addr = NtPrePrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrePrepareEnlistment" => {
            let addr = NtPrePrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrepareComplete" => {
            let addr = NtPrepareComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrepareEnlistment" => {
            let addr = NtPrepareEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrivilegeCheck" => {
            let addr = NtPrivilegeCheck as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrivilegeObjectAuditAlarm" => {
            let addr = NtPrivilegeObjectAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPrivilegedServiceAuditAlarm" => {
            let addr = NtPrivilegedServiceAuditAlarm as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPropagationComplete" => {
            let addr = NtPropagationComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtPropagationFailed" => {
            let addr = NtPropagationFailed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryDirectoryFile" => {
            let addr = NtQueryDirectoryFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryDirectoryFileEx" => {
            let addr = NtQueryDirectoryFileEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationByName" => {
            let addr = NtQueryInformationByName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationEnlistment" => {
            let addr = NtQueryInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationFile" => {
            let addr = NtQueryInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationResourceManager" => {
            let addr = NtQueryInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationToken" => {
            let addr = NtQueryInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationTransaction" => {
            let addr = NtQueryInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryInformationTransactionManager" => {
            let addr = NtQueryInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryObject" => {
            let addr = NtQueryObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryQuotaInformationFile" => {
            let addr = NtQueryQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQuerySecurityObject" => {
            let addr = NtQuerySecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryVirtualMemory" => {
            let addr = NtQueryVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtQueryVolumeInformationFile" => {
            let addr = NtQueryVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtReadFile" => {
            let addr = NtReadFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtReadOnlyEnlistment" => {
            let addr = NtReadOnlyEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRecoverEnlistment" => {
            let addr = NtRecoverEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRecoverResourceManager" => {
            let addr = NtRecoverResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRecoverTransactionManager" => {
            let addr = NtRecoverTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRegisterProtocolAddressInformation" => {
            let addr = NtRegisterProtocolAddressInformation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRenameTransactionManager" => {
            let addr = NtRenameTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRollbackComplete" => {
            let addr = NtRollbackComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRollbackEnlistment" => {
            let addr = NtRollbackEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRollbackRegistryTransaction" => {
            let addr = NtRollbackRegistryTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRollbackTransaction" => {
            let addr = NtRollbackTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtRollforwardTransactionManager" => {
            let addr = NtRollforwardTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationEnlistment" => {
            let addr = NtSetInformationEnlistment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationFile" => {
            let addr = NtSetInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationResourceManager" => {
            let addr = NtSetInformationResourceManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationThread" => {
            let addr = NtSetInformationThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationToken" => {
            let addr = NtSetInformationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationTransaction" => {
            let addr = NtSetInformationTransaction as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationTransactionManager" => {
            let addr = NtSetInformationTransactionManager as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetInformationVirtualMemory" => {
            let addr = NtSetInformationVirtualMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetQuotaInformationFile" => {
            let addr = NtSetQuotaInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetSecurityObject" => {
            let addr = NtSetSecurityObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSetVolumeInformationFile" => {
            let addr = NtSetVolumeInformationFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtSinglePhaseReject" => {
            let addr = NtSinglePhaseReject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtUnlockFile" => {
            let addr = NtUnlockFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "NtWriteFile" => {
            let addr = NtWriteFile as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObCloseHandle" => {
            let addr = ObCloseHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObCloseHandleWithResult" => {
            let addr = ObCloseHandleWithResult as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObDereferenceObjectDeferDelete" => {
            let addr = ObDereferenceObjectDeferDelete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObDereferenceObjectDeferDeleteWithTag" => {
            let addr = ObDereferenceObjectDeferDeleteWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObGetFilterVersion" => {
            let addr = ObGetFilterVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObGetObjectSecurity" => {
            let addr = ObGetObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObInsertObject" => {
            let addr = ObInsertObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObIsKernelHandle" => {
            let addr = ObIsKernelHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObMakeTemporaryObject" => {
            let addr = ObMakeTemporaryObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObOpenObjectByPointer" => {
            let addr = ObOpenObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObOpenObjectByPointerWithTag" => {
            let addr = ObOpenObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObQueryNameString" => {
            let addr = ObQueryNameString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObQueryObjectAuditingByHandle" => {
            let addr = ObQueryObjectAuditingByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectByHandle" => {
            let addr = ObReferenceObjectByHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectByHandleWithTag" => {
            let addr = ObReferenceObjectByHandleWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectByPointer" => {
            let addr = ObReferenceObjectByPointer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectByPointerWithTag" => {
            let addr = ObReferenceObjectByPointerWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectSafe" => {
            let addr = ObReferenceObjectSafe as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReferenceObjectSafeWithTag" => {
            let addr = ObReferenceObjectSafeWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObRegisterCallbacks" => {
            let addr = ObRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObReleaseObjectSecurity" => {
            let addr = ObReleaseObjectSecurity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObUnRegisterCallbacks" => {
            let addr = ObUnRegisterCallbacks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObfDereferenceObject" => {
            let addr = ObfDereferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObfDereferenceObjectWithTag" => {
            let addr = ObfDereferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObfReferenceObject" => {
            let addr = ObfReferenceObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "ObfReferenceObjectWithTag" => {
            let addr = ObfReferenceObjectWithTag as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsAcquireSiloHardReference" => {
            let addr = PsAcquireSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsAllocSiloContextSlot" => {
            let addr = PsAllocSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsAllocateAffinityToken" => {
            let addr = PsAllocateAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsAssignImpersonationToken" => {
            let addr = PsAssignImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsAttachSiloToCurrentThread" => {
            let addr = PsAttachSiloToCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsChargePoolQuota" => {
            let addr = PsChargePoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsChargeProcessPoolQuota" => {
            let addr = PsChargeProcessPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsCreateSiloContext" => {
            let addr = PsCreateSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsCreateSystemThread" => {
            let addr = PsCreateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsDereferenceImpersonationToken" => {
            let addr = PsDereferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsDereferencePrimaryToken" => {
            let addr = PsDereferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsDereferenceSiloContext" => {
            let addr = PsDereferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsDetachSiloFromCurrentThread" => {
            let addr = PsDetachSiloFromCurrentThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsDisableImpersonation" => {
            let addr = PsDisableImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsFreeAffinityToken" => {
            let addr = PsFreeAffinityToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsFreeSiloContextSlot" => {
            let addr = PsFreeSiloContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentProcessId" => {
            let addr = PsGetCurrentProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentServerSilo" => {
            let addr = PsGetCurrentServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentServerSiloName" => {
            let addr = PsGetCurrentServerSiloName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentSilo" => {
            let addr = PsGetCurrentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentThreadId" => {
            let addr = PsGetCurrentThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetCurrentThreadTeb" => {
            let addr = PsGetCurrentThreadTeb as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetEffectiveServerSilo" => {
            let addr = PsGetEffectiveServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetHostSilo" => {
            let addr = PsGetHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetJobServerSilo" => {
            let addr = PsGetJobServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetJobSilo" => {
            let addr = PsGetJobSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetParentSilo" => {
            let addr = PsGetParentSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetPermanentSiloContext" => {
            let addr = PsGetPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetProcessCreateTimeQuadPart" => {
            let addr = PsGetProcessCreateTimeQuadPart as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetProcessExitStatus" => {
            let addr = PsGetProcessExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetProcessExitTime" => {
            let addr = PsGetProcessExitTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetProcessId" => {
            let addr = PsGetProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetProcessStartKey" => {
            let addr = PsGetProcessStartKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetServerSiloActiveConsoleId" => {
            let addr = PsGetServerSiloActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetServerSiloServiceSessionId" => {
            let addr = PsGetServerSiloServiceSessionId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetSiloContainerId" => {
            let addr = PsGetSiloContainerId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetSiloContext" => {
            let addr = PsGetSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetSiloMonitorContextSlot" => {
            let addr = PsGetSiloMonitorContextSlot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadCreateTime" => {
            let addr = PsGetThreadCreateTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadExitStatus" => {
            let addr = PsGetThreadExitStatus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadId" => {
            let addr = PsGetThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadProcess" => {
            let addr = PsGetThreadProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadProcessId" => {
            let addr = PsGetThreadProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadProperty" => {
            let addr = PsGetThreadProperty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetThreadServerSilo" => {
            let addr = PsGetThreadServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsGetVersion" => {
            let addr = PsGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsImpersonateClient" => {
            let addr = PsImpersonateClient as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsInsertPermanentSiloContext" => {
            let addr = PsInsertPermanentSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsInsertSiloContext" => {
            let addr = PsInsertSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsCurrentThreadInServerSilo" => {
            let addr = PsIsCurrentThreadInServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsCurrentThreadPrefetching" => {
            let addr = PsIsCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsDiskCountersEnabled" => {
            let addr = PsIsDiskCountersEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsHostSilo" => {
            let addr = PsIsHostSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsSystemThread" => {
            let addr = PsIsSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsThreadAttachedToSpecificSilo" => {
            let addr = PsIsThreadAttachedToSpecificSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsIsThreadTerminating" => {
            let addr = PsIsThreadTerminating as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsLookupProcessByProcessId" => {
            let addr = PsLookupProcessByProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsLookupThreadByThreadId" => {
            let addr = PsLookupThreadByThreadId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsMakeSiloContextPermanent" => {
            let addr = PsMakeSiloContextPermanent as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsQueryProcessAvailableCpus" => {
            let addr = PsQueryProcessAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsQueryProcessAvailableCpusCount" => {
            let addr = PsQueryProcessAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsQuerySystemAvailableCpus" => {
            let addr = PsQuerySystemAvailableCpus as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsQuerySystemAvailableCpusCount" => {
            let addr = PsQuerySystemAvailableCpusCount as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsQueryTotalCycleTimeProcess" => {
            let addr = PsQueryTotalCycleTimeProcess as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReferenceImpersonationToken" => {
            let addr = PsReferenceImpersonationToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReferencePrimaryToken" => {
            let addr = PsReferencePrimaryToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReferenceSiloContext" => {
            let addr = PsReferenceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRegisterProcessAvailableCpusChangeNotification" => {
            let addr = PsRegisterProcessAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRegisterSiloMonitor" => {
            let addr = PsRegisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRegisterSystemAvailableCpusChangeNotification" => {
            let addr = PsRegisterSystemAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReleaseSiloHardReference" => {
            let addr = PsReleaseSiloHardReference as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRemoveCreateThreadNotifyRoutine" => {
            let addr = PsRemoveCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRemoveLoadImageNotifyRoutine" => {
            let addr = PsRemoveLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRemoveSiloContext" => {
            let addr = PsRemoveSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReplaceSiloContext" => {
            let addr = PsReplaceSiloContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRestoreImpersonation" => {
            let addr = PsRestoreImpersonation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsReturnPoolQuota" => {
            let addr = PsReturnPoolQuota as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRevertToSelf" => {
            let addr = PsRevertToSelf as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsRevertToUserMultipleGroupAffinityThread" => {
            let addr = PsRevertToUserMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutine" => {
            let addr = PsSetCreateProcessNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutineEx" => {
            let addr = PsSetCreateProcessNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCreateProcessNotifyRoutineEx2" => {
            let addr = PsSetCreateProcessNotifyRoutineEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCreateThreadNotifyRoutine" => {
            let addr = PsSetCreateThreadNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCreateThreadNotifyRoutineEx" => {
            let addr = PsSetCreateThreadNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetCurrentThreadPrefetching" => {
            let addr = PsSetCurrentThreadPrefetching as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetLoadImageNotifyRoutine" => {
            let addr = PsSetLoadImageNotifyRoutine as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetLoadImageNotifyRoutineEx" => {
            let addr = PsSetLoadImageNotifyRoutineEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsSetSystemMultipleGroupAffinityThread" => {
            let addr = PsSetSystemMultipleGroupAffinityThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsStartSiloMonitor" => {
            let addr = PsStartSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsTerminateServerSilo" => {
            let addr = PsTerminateServerSilo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsTerminateSystemThread" => {
            let addr = PsTerminateSystemThread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsUnregisterAvailableCpusChangeNotification" => {
            let addr = PsUnregisterAvailableCpusChangeNotification as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsUnregisterSiloMonitor" => {
            let addr = PsUnregisterSiloMonitor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsUpdateDiskCounters" => {
            let addr = PsUpdateDiskCounters as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PsWrapApcWow64Thread" => {
            let addr = PsWrapApcWow64Thread as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedAllocateMemory" => {
            let addr = PshedAllocateMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedFreeMemory" => {
            let addr = PshedFreeMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedIsSystemWheaEnabled" => {
            let addr = PshedIsSystemWheaEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedRegisterPlugin" => {
            let addr = PshedRegisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedSynchronizeExecution" => {
            let addr = PshedSynchronizeExecution as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "PshedUnregisterPlugin" => {
            let addr = PshedUnregisterPlugin as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAbsoluteToSelfRelativeSD" => {
            let addr = RtlAbsoluteToSelfRelativeSD as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAddAccessAllowedAce" => {
            let addr = RtlAddAccessAllowedAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAddAccessAllowedAceEx" => {
            let addr = RtlAddAccessAllowedAceEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAddAce" => {
            let addr = RtlAddAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAllocateAndInitializeSid" => {
            let addr = RtlAllocateAndInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAllocateAndInitializeSidEx" => {
            let addr = RtlAllocateAndInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAllocateHeap" => {
            let addr = RtlAllocateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAnsiStringToUnicodeString" => {
            let addr = RtlAnsiStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAppendStringToString" => {
            let addr = RtlAppendStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAppendUnicodeStringToString" => {
            let addr = RtlAppendUnicodeStringToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAppendUnicodeToString" => {
            let addr = RtlAppendUnicodeToString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAreBitsClear" => {
            let addr = RtlAreBitsClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAreBitsSet" => {
            let addr = RtlAreBitsSet as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlAssert" => {
            let addr = RtlAssert as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCaptureContext" => {
            let addr = RtlCaptureContext as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCaptureContext2" => {
            let addr = RtlCaptureContext2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCaptureStackBackTrace" => {
            let addr = RtlCaptureStackBackTrace as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCharToInteger" => {
            let addr = RtlCharToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCheckRegistryKey" => {
            let addr = RtlCheckRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlClearAllBits" => {
            let addr = RtlClearAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlClearBit" => {
            let addr = RtlClearBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlClearBits" => {
            let addr = RtlClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCmDecodeMemIoResource" => {
            let addr = RtlCmDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCmEncodeMemIoResource" => {
            let addr = RtlCmEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareAltitudes" => {
            let addr = RtlCompareAltitudes as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareMemory" => {
            let addr = RtlCompareMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareMemoryUlong" => {
            let addr = RtlCompareMemoryUlong as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareString" => {
            let addr = RtlCompareString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareUnicodeString" => {
            let addr = RtlCompareUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompareUnicodeStrings" => {
            let addr = RtlCompareUnicodeStrings as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompressBuffer" => {
            let addr = RtlCompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCompressChunks" => {
            let addr = RtlCompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlContractHashTable" => {
            let addr = RtlContractHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlConvertSidToUnicodeString" => {
            let addr = RtlConvertSidToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyBitMap" => {
            let addr = RtlCopyBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyDeviceMemory" => {
            let addr = RtlCopyDeviceMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyLuid" => {
            let addr = RtlCopyLuid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyMemoryNonTemporal" => {
            let addr = RtlCopyMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopySid" => {
            let addr = RtlCopySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyString" => {
            let addr = RtlCopyString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyUnicodeString" => {
            let addr = RtlCopyUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCopyVolatileMemory" => {
            let addr = RtlCopyVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCrc32" => {
            let addr = RtlCrc32 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCrc64" => {
            let addr = RtlCrc64 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateAcl" => {
            let addr = RtlCreateAcl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateHashTable" => {
            let addr = RtlCreateHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateHashTableEx" => {
            let addr = RtlCreateHashTableEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateHeap" => {
            let addr = RtlCreateHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateRegistryKey" => {
            let addr = RtlCreateRegistryKey as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateSecurityDescriptor" => {
            let addr = RtlCreateSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateSecurityDescriptorRelative" => {
            let addr = RtlCreateSecurityDescriptorRelative as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateServiceSid" => {
            let addr = RtlCreateServiceSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateSystemVolumeInformationFolder" => {
            let addr = RtlCreateSystemVolumeInformationFolder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateUnicodeString" => {
            let addr = RtlCreateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCreateVirtualAccountSid" => {
            let addr = RtlCreateVirtualAccountSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlCustomCPToUnicodeN" => {
            let addr = RtlCustomCPToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressBuffer" => {
            let addr = RtlDecompressBuffer as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressBufferEx" => {
            let addr = RtlDecompressBufferEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressBufferEx2" => {
            let addr = RtlDecompressBufferEx2 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressChunks" => {
            let addr = RtlDecompressChunks as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressFragment" => {
            let addr = RtlDecompressFragment as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDecompressFragmentEx" => {
            let addr = RtlDecompressFragmentEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDelete" => {
            let addr = RtlDelete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteAce" => {
            let addr = RtlDeleteAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteElementGenericTable" => {
            let addr = RtlDeleteElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteElementGenericTableAvl" => {
            let addr = RtlDeleteElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteElementGenericTableAvlEx" => {
            let addr = RtlDeleteElementGenericTableAvlEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteHashTable" => {
            let addr = RtlDeleteHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteNoSplay" => {
            let addr = RtlDeleteNoSplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDeleteRegistryValue" => {
            let addr = RtlDeleteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDescribeChunk" => {
            let addr = RtlDescribeChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDestroyHeap" => {
            let addr = RtlDestroyHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDowncaseUnicodeChar" => {
            let addr = RtlDowncaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDowncaseUnicodeString" => {
            let addr = RtlDowncaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDrainNonVolatileFlush" => {
            let addr = RtlDrainNonVolatileFlush as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlDuplicateUnicodeString" => {
            let addr = RtlDuplicateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEndEnumerationHashTable" => {
            let addr = RtlEndEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEndStrongEnumerationHashTable" => {
            let addr = RtlEndStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEndWeakEnumerationHashTable" => {
            let addr = RtlEndWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateEntryHashTable" => {
            let addr = RtlEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateGenericTable" => {
            let addr = RtlEnumerateGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateGenericTableAvl" => {
            let addr = RtlEnumerateGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateGenericTableLikeADirectory" => {
            let addr = RtlEnumerateGenericTableLikeADirectory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateGenericTableWithoutSplaying" => {
            let addr = RtlEnumerateGenericTableWithoutSplaying as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEnumerateGenericTableWithoutSplayingAvl" => {
            let addr = RtlEnumerateGenericTableWithoutSplayingAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEqualPrefixSid" => {
            let addr = RtlEqualPrefixSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEqualSid" => {
            let addr = RtlEqualSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEqualString" => {
            let addr = RtlEqualString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlEqualUnicodeString" => {
            let addr = RtlEqualUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlExpandHashTable" => {
            let addr = RtlExpandHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlExtendCorrelationVector" => {
            let addr = RtlExtendCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlExtractBitMap" => {
            let addr = RtlExtractBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFillMemoryNonTemporal" => {
            let addr = RtlFillMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFillNonVolatileMemory" => {
            let addr = RtlFillNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindClearBits" => {
            let addr = RtlFindClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindClearBitsAndSet" => {
            let addr = RtlFindClearBitsAndSet as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindClearRuns" => {
            let addr = RtlFindClearRuns as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindClosestEncodableLength" => {
            let addr = RtlFindClosestEncodableLength as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindFirstRunClear" => {
            let addr = RtlFindFirstRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindLastBackwardRunClear" => {
            let addr = RtlFindLastBackwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindLeastSignificantBit" => {
            let addr = RtlFindLeastSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindLongestRunClear" => {
            let addr = RtlFindLongestRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindMostSignificantBit" => {
            let addr = RtlFindMostSignificantBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindNextForwardRunClear" => {
            let addr = RtlFindNextForwardRunClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindSetBits" => {
            let addr = RtlFindSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindSetBitsAndClear" => {
            let addr = RtlFindSetBitsAndClear as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFindUnicodePrefix" => {
            let addr = RtlFindUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFlushNonVolatileMemory" => {
            let addr = RtlFlushNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFlushNonVolatileMemoryRanges" => {
            let addr = RtlFlushNonVolatileMemoryRanges as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeAnsiString" => {
            let addr = RtlFreeAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeHeap" => {
            let addr = RtlFreeHeap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeNonVolatileToken" => {
            let addr = RtlFreeNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeOemString" => {
            let addr = RtlFreeOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeSid" => {
            let addr = RtlFreeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeUTF8String" => {
            let addr = RtlFreeUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlFreeUnicodeString" => {
            let addr = RtlFreeUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGUIDFromString" => {
            let addr = RtlGUIDFromString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGenerate8dot3Name" => {
            let addr = RtlGenerate8dot3Name as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGenerateClass5Guid" => {
            let addr = RtlGenerateClass5Guid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetAce" => {
            let addr = RtlGetAce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetAcesBufferSize" => {
            let addr = RtlGetAcesBufferSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetActiveConsoleId" => {
            let addr = RtlGetActiveConsoleId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetCallersAddress" => {
            let addr = RtlGetCallersAddress as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetCompressionWorkSpaceSize" => {
            let addr = RtlGetCompressionWorkSpaceSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetConsoleSessionForegroundProcessId" => {
            let addr = RtlGetConsoleSessionForegroundProcessId as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetDaclSecurityDescriptor" => {
            let addr = RtlGetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetElementGenericTable" => {
            let addr = RtlGetElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetElementGenericTableAvl" => {
            let addr = RtlGetElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetEnabledExtendedFeatures" => {
            let addr = RtlGetEnabledExtendedFeatures as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetGroupSecurityDescriptor" => {
            let addr = RtlGetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetNextEntryHashTable" => {
            let addr = RtlGetNextEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetNonVolatileToken" => {
            let addr = RtlGetNonVolatileToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetNtProductType" => {
            let addr = RtlGetNtProductType as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetNtSystemRoot" => {
            let addr = RtlGetNtSystemRoot as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetOwnerSecurityDescriptor" => {
            let addr = RtlGetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetPersistedStateLocation" => {
            let addr = RtlGetPersistedStateLocation as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetProductInfo" => {
            let addr = RtlGetProductInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetSaclSecurityDescriptor" => {
            let addr = RtlGetSaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetSuiteMask" => {
            let addr = RtlGetSuiteMask as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetSystemGlobalData" => {
            let addr = RtlGetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlGetVersion" => {
            let addr = RtlGetVersion as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlHashUnicodeString" => {
            let addr = RtlHashUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIdentifierAuthoritySid" => {
            let addr = RtlIdentifierAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIdnToAscii" => {
            let addr = RtlIdnToAscii as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIdnToNameprepUnicode" => {
            let addr = RtlIdnToNameprepUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIdnToUnicode" => {
            let addr = RtlIdnToUnicode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIncrementCorrelationVector" => {
            let addr = RtlIncrementCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitAnsiString" => {
            let addr = RtlInitAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitAnsiStringEx" => {
            let addr = RtlInitAnsiStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitCodePageTable" => {
            let addr = RtlInitCodePageTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitEnumerationHashTable" => {
            let addr = RtlInitEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitString" => {
            let addr = RtlInitString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitStringEx" => {
            let addr = RtlInitStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitStrongEnumerationHashTable" => {
            let addr = RtlInitStrongEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitUTF8String" => {
            let addr = RtlInitUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitUTF8StringEx" => {
            let addr = RtlInitUTF8StringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitUnicodeString" => {
            let addr = RtlInitUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitUnicodeStringEx" => {
            let addr = RtlInitUnicodeStringEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitWeakEnumerationHashTable" => {
            let addr = RtlInitWeakEnumerationHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeBitMap" => {
            let addr = RtlInitializeBitMap as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeCorrelationVector" => {
            let addr = RtlInitializeCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeGenericTable" => {
            let addr = RtlInitializeGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeGenericTableAvl" => {
            let addr = RtlInitializeGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeSid" => {
            let addr = RtlInitializeSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeSidEx" => {
            let addr = RtlInitializeSidEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInitializeUnicodePrefix" => {
            let addr = RtlInitializeUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertElementGenericTable" => {
            let addr = RtlInsertElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertElementGenericTableAvl" => {
            let addr = RtlInsertElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertElementGenericTableFull" => {
            let addr = RtlInsertElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertElementGenericTableFullAvl" => {
            let addr = RtlInsertElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertEntryHashTable" => {
            let addr = RtlInsertEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInsertUnicodePrefix" => {
            let addr = RtlInsertUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlInt64ToUnicodeString" => {
            let addr = RtlInt64ToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIntegerToUnicodeString" => {
            let addr = RtlIntegerToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIoDecodeMemIoResource" => {
            let addr = RtlIoDecodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIoEncodeMemIoResource" => {
            let addr = RtlIoEncodeMemIoResource as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsApiSetImplemented" => {
            let addr = RtlIsApiSetImplemented as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsCloudFilesPlaceholder" => {
            let addr = RtlIsCloudFilesPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsGenericTableEmpty" => {
            let addr = RtlIsGenericTableEmpty as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsGenericTableEmptyAvl" => {
            let addr = RtlIsGenericTableEmptyAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsMultiSessionSku" => {
            let addr = RtlIsMultiSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsMultiUsersInSessionSku" => {
            let addr = RtlIsMultiUsersInSessionSku as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsNameLegalDOS8Dot3" => {
            let addr = RtlIsNameLegalDOS8Dot3 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsNonEmptyDirectoryReparsePointAllowed" => {
            let addr = RtlIsNonEmptyDirectoryReparsePointAllowed as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsNormalizedString" => {
            let addr = RtlIsNormalizedString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsNtDdiVersionAvailable" => {
            let addr = RtlIsNtDdiVersionAvailable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsPartialPlaceholder" => {
            let addr = RtlIsPartialPlaceholder as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsPartialPlaceholderFileHandle" => {
            let addr = RtlIsPartialPlaceholderFileHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsPartialPlaceholderFileInfo" => {
            let addr = RtlIsPartialPlaceholderFileInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsSandboxedToken" => {
            let addr = RtlIsSandboxedToken as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsSandboxedTokenHandle" => {
            let addr = RtlIsSandboxedTokenHandle as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsServicePackVersionInstalled" => {
            let addr = RtlIsServicePackVersionInstalled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsStateSeparationEnabled" => {
            let addr = RtlIsStateSeparationEnabled as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsUntrustedObject" => {
            let addr = RtlIsUntrustedObject as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsValidOemCharacter" => {
            let addr = RtlIsValidOemCharacter as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlIsZeroMemory" => {
            let addr = RtlIsZeroMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLengthRequiredSid" => {
            let addr = RtlLengthRequiredSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLengthSecurityDescriptor" => {
            let addr = RtlLengthSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLengthSid" => {
            let addr = RtlLengthSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupElementGenericTable" => {
            let addr = RtlLookupElementGenericTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupElementGenericTableAvl" => {
            let addr = RtlLookupElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupElementGenericTableFull" => {
            let addr = RtlLookupElementGenericTableFull as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupElementGenericTableFullAvl" => {
            let addr = RtlLookupElementGenericTableFullAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupEntryHashTable" => {
            let addr = RtlLookupEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlLookupFirstMatchingElementGenericTableAvl" => {
            let addr = RtlLookupFirstMatchingElementGenericTableAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlMapGenericMask" => {
            let addr = RtlMapGenericMask as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlMoveVolatileMemory" => {
            let addr = RtlMoveVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlMultiByteToUnicodeN" => {
            let addr = RtlMultiByteToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlMultiByteToUnicodeSize" => {
            let addr = RtlMultiByteToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNextUnicodePrefix" => {
            let addr = RtlNextUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNormalizeSecurityDescriptor" => {
            let addr = RtlNormalizeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNormalizeString" => {
            let addr = RtlNormalizeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNtStatusToDosError" => {
            let addr = RtlNtStatusToDosError as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNtStatusToDosErrorNoTeb" => {
            let addr = RtlNtStatusToDosErrorNoTeb as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberGenericTableElements" => {
            let addr = RtlNumberGenericTableElements as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberGenericTableElementsAvl" => {
            let addr = RtlNumberGenericTableElementsAvl as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberOfClearBits" => {
            let addr = RtlNumberOfClearBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberOfClearBitsInRange" => {
            let addr = RtlNumberOfClearBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberOfSetBits" => {
            let addr = RtlNumberOfSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberOfSetBitsInRange" => {
            let addr = RtlNumberOfSetBitsInRange as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlNumberOfSetBitsUlongPtr" => {
            let addr = RtlNumberOfSetBitsUlongPtr as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlOemStringToCountedUnicodeString" => {
            let addr = RtlOemStringToCountedUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlOemStringToUnicodeString" => {
            let addr = RtlOemStringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlOemToUnicodeN" => {
            let addr = RtlOemToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlOsDeploymentState" => {
            let addr = RtlOsDeploymentState as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlPrefetchMemoryNonTemporal" => {
            let addr = RtlPrefetchMemoryNonTemporal as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlPrefixString" => {
            let addr = RtlPrefixString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlPrefixUnicodeString" => {
            let addr = RtlPrefixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryPackageIdentity" => {
            let addr = RtlQueryPackageIdentity as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryPackageIdentityEx" => {
            let addr = RtlQueryPackageIdentityEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryProcessPlaceholderCompatibilityMode" => {
            let addr = RtlQueryProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryRegistryValueWithFallback" => {
            let addr = RtlQueryRegistryValueWithFallback as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryRegistryValues" => {
            let addr = RtlQueryRegistryValues as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryThreadPlaceholderCompatibilityMode" => {
            let addr = RtlQueryThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlQueryValidationRunlevel" => {
            let addr = RtlQueryValidationRunlevel as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRaiseCustomSystemEventTrigger" => {
            let addr = RtlRaiseCustomSystemEventTrigger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRandom" => {
            let addr = RtlRandom as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRandomEx" => {
            let addr = RtlRandomEx as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRealPredecessor" => {
            let addr = RtlRealPredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRealSuccessor" => {
            let addr = RtlRealSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRemoveEntryHashTable" => {
            let addr = RtlRemoveEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRemoveUnicodePrefix" => {
            let addr = RtlRemoveUnicodePrefix as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlReplaceSidInSd" => {
            let addr = RtlReplaceSidInSd as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlReserveChunk" => {
            let addr = RtlReserveChunk as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRunOnceBeginInitialize" => {
            let addr = RtlRunOnceBeginInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRunOnceComplete" => {
            let addr = RtlRunOnceComplete as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRunOnceExecuteOnce" => {
            let addr = RtlRunOnceExecuteOnce as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlRunOnceInitialize" => {
            let addr = RtlRunOnceInitialize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSecondsSince1970ToTime" => {
            let addr = RtlSecondsSince1970ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSecondsSince1980ToTime" => {
            let addr = RtlSecondsSince1980ToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSelfRelativeToAbsoluteSD" => {
            let addr = RtlSelfRelativeToAbsoluteSD as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetAllBits" => {
            let addr = RtlSetAllBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetBit" => {
            let addr = RtlSetBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetBits" => {
            let addr = RtlSetBits as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetDaclSecurityDescriptor" => {
            let addr = RtlSetDaclSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetGroupSecurityDescriptor" => {
            let addr = RtlSetGroupSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetOwnerSecurityDescriptor" => {
            let addr = RtlSetOwnerSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetProcessPlaceholderCompatibilityMode" => {
            let addr = RtlSetProcessPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetSystemGlobalData" => {
            let addr = RtlSetSystemGlobalData as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetThreadPlaceholderCompatibilityMode" => {
            let addr = RtlSetThreadPlaceholderCompatibilityMode as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSetVolatileMemory" => {
            let addr = RtlSetVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSplay" => {
            let addr = RtlSplay as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlStringFromGUID" => {
            let addr = RtlStringFromGUID as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlStronglyEnumerateEntryHashTable" => {
            let addr = RtlStronglyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSubAuthorityCountSid" => {
            let addr = RtlSubAuthorityCountSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSubAuthoritySid" => {
            let addr = RtlSubAuthoritySid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSubtreePredecessor" => {
            let addr = RtlSubtreePredecessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSubtreeSuccessor" => {
            let addr = RtlSubtreeSuccessor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlSuffixUnicodeString" => {
            let addr = RtlSuffixUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlTestBit" => {
            let addr = RtlTestBit as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlTimeFieldsToTime" => {
            let addr = RtlTimeFieldsToTime as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlTimeToSecondsSince1970" => {
            let addr = RtlTimeToSecondsSince1970 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlTimeToSecondsSince1980" => {
            let addr = RtlTimeToSecondsSince1980 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlTimeToTimeFields" => {
            let addr = RtlTimeToTimeFields as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUTF8StringToUnicodeString" => {
            let addr = RtlUTF8StringToUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUTF8ToUnicodeN" => {
            let addr = RtlUTF8ToUnicodeN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToAnsiString" => {
            let addr = RtlUnicodeStringToAnsiString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToCountedOemString" => {
            let addr = RtlUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToInt64" => {
            let addr = RtlUnicodeStringToInt64 as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToInteger" => {
            let addr = RtlUnicodeStringToInteger as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToOemString" => {
            let addr = RtlUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeStringToUTF8String" => {
            let addr = RtlUnicodeStringToUTF8String as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeToCustomCPN" => {
            let addr = RtlUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeToMultiByteN" => {
            let addr = RtlUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeToMultiByteSize" => {
            let addr = RtlUnicodeToMultiByteSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeToOemN" => {
            let addr = RtlUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUnicodeToUTF8N" => {
            let addr = RtlUnicodeToUTF8N as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeChar" => {
            let addr = RtlUpcaseUnicodeChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeString" => {
            let addr = RtlUpcaseUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeStringToCountedOemString" => {
            let addr = RtlUpcaseUnicodeStringToCountedOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeStringToOemString" => {
            let addr = RtlUpcaseUnicodeStringToOemString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeToCustomCPN" => {
            let addr = RtlUpcaseUnicodeToCustomCPN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeToMultiByteN" => {
            let addr = RtlUpcaseUnicodeToMultiByteN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpcaseUnicodeToOemN" => {
            let addr = RtlUpcaseUnicodeToOemN as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpperChar" => {
            let addr = RtlUpperChar as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlUpperString" => {
            let addr = RtlUpperString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlValidRelativeSecurityDescriptor" => {
            let addr = RtlValidRelativeSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlValidSecurityDescriptor" => {
            let addr = RtlValidSecurityDescriptor as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlValidSid" => {
            let addr = RtlValidSid as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlValidateCorrelationVector" => {
            let addr = RtlValidateCorrelationVector as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlValidateUnicodeString" => {
            let addr = RtlValidateUnicodeString as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlVerifyVersionInfo" => {
            let addr = RtlVerifyVersionInfo as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlVolumeDeviceToDosName" => {
            let addr = RtlVolumeDeviceToDosName as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlWalkFrameChain" => {
            let addr = RtlWalkFrameChain as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlWeaklyEnumerateEntryHashTable" => {
            let addr = RtlWeaklyEnumerateEntryHashTable as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlWriteNonVolatileMemory" => {
            let addr = RtlWriteNonVolatileMemory as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlWriteRegistryValue" => {
            let addr = RtlWriteRegistryValue as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlxAnsiStringToUnicodeSize" => {
            let addr = RtlxAnsiStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlxOemStringToUnicodeSize" => {
            let addr = RtlxOemStringToUnicodeSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlxUnicodeStringToAnsiSize" => {
            let addr = RtlxUnicodeStringToAnsiSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        "RtlxUnicodeStringToOemSize" => {
            let addr = RtlxUnicodeStringToOemSize as u64;
            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }
        }
        _ => Err(HookError::NotHooked),
    }
}
