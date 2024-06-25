package symbol

type (
	InterfaceWin32u interface {
		NtBindCompositionSurface() (ok bool)
		NtCloseCompositionInputSink() (ok bool)
		NtCompositionInputThread() (ok bool)
		NtCompositionSetDropTarget() (ok bool)
		NtCompositorNotifyExitWindows() (ok bool)
		NtConfigureInputSpace() (ok bool)
		NtCreateCompositionInputSink() (ok bool)
		NtCreateCompositionSurfaceHandle() (ok bool)
		NtCreateImplicitCompositionInputSink() (ok bool)
		NtDCompositionAddCrossDeviceVisualChild() (ok bool)
		NtDCompositionBeginFrame() (ok bool)
		NtDCompositionBoostCompositorClock() (ok bool)
		NtDCompositionCommitChannel() (ok bool)
		NtDCompositionCommitSynchronizationObject() (ok bool)
		NtDCompositionConfirmFrame() (ok bool)
		NtDCompositionConnectPipe() (ok bool)
		NtDCompositionCreateAndBindSharedSection() (ok bool)
		NtDCompositionCreateChannel() (ok bool)
		NtDCompositionCreateConnection() (ok bool)
		NtDCompositionCreateDwmChannel() (ok bool)
		NtDCompositionCreateSharedResourceHandle() (ok bool)
		NtDCompositionCreateSynchronizationObject() (ok bool)
		NtDCompositionDestroyChannel() (ok bool)
		NtDCompositionDestroyConnection() (ok bool)
		NtDCompositionDuplicateHandleToProcess() (ok bool)
		NtDCompositionDuplicateSwapchainHandleToDwm() (ok bool)
		NtDCompositionEnableMMCSS() (ok bool)
		NtDCompositionGetBatchId() (ok bool)
		NtDCompositionGetChannels() (ok bool)
		NtDCompositionGetConnectionBatch() (ok bool)
		NtDCompositionGetDeletedResources() (ok bool)
		NtDCompositionGetFrameId() (ok bool)
		NtDCompositionGetFrameIdFromBatchId() (ok bool)
		NtDCompositionGetFrameLegacyTokens() (ok bool)
		NtDCompositionGetFrameStatistics() (ok bool)
		NtDCompositionGetFrameSurfaceUpdates() (ok bool)
		NtDCompositionGetMaterialProperty() (ok bool)
		NtDCompositionGetStatistics() (ok bool)
		NtDCompositionGetTargetStatistics() (ok bool)
		NtDCompositionNotifySuperWetInkWork() (ok bool)
		NtDCompositionProcessChannelBatchBuffer() (ok bool)
		NtDCompositionReferenceSharedResourceOnDwmChannel() (ok bool)
		NtDCompositionRegisterThumbnailVisual() (ok bool)
		NtDCompositionRegisterVirtualDesktopVisual() (ok bool)
		NtDCompositionReleaseAllResources() (ok bool)
		NtDCompositionRemoveCrossDeviceVisualChild() (ok bool)
		NtDCompositionSetBlurredWallpaperSurface() (ok bool)
		NtDCompositionSetChannelCommitCompletionEvent() (ok bool)
		NtDCompositionSetChannelConnectionId() (ok bool)
		NtDCompositionSetChildRootVisual() (ok bool)
		NtDCompositionSetDebugCounter() (ok bool)
		NtDCompositionSetMaterialProperty() (ok bool)
		NtDCompositionSubmitDWMBatch() (ok bool)
		NtDCompositionSuspendAnimations() (ok bool)
		NtDCompositionSynchronize() (ok bool)
		NtDCompositionTelemetrySetApplicationId() (ok bool)
		NtDCompositionUpdatePointerCapture() (ok bool)
		NtDCompositionWaitForChannel() (ok bool)
		NtDCompositionWaitForCompositorClock() (ok bool)
		NtDesktopCaptureBits() (ok bool)
		NtDuplicateCompositionInputSink() (ok bool)
		NtDxgkCancelPresents() (ok bool)
		NtDxgkCheckSinglePlaneForMultiPlaneOverlaySupport() (ok bool)
		NtDxgkConnectDoorbell() (ok bool)
		NtDxgkCreateDoorbell() (ok bool)
		NtDxgkCreateNativeFence() (ok bool)
		NtDxgkCreateTrackedWorkload() (ok bool)
		NtDxgkDestroyDoorbell() (ok bool)
		NtDxgkDestroyTrackedWorkload() (ok bool)
		NtDxgkDispMgrOperation() (ok bool)
		NtDxgkDisplayPortOperation() (ok bool)
		NtDxgkDuplicateHandle() (ok bool)
		NtDxgkEnumAdapters3() (ok bool)
		NtDxgkEnumProcesses() (ok bool)
		NtDxgkGetAvailableTrackedWorkloadIndex() (ok bool)
		NtDxgkGetProcessList() (ok bool)
		NtDxgkGetProperties() (ok bool)
		NtDxgkGetTrackedWorkloadStatistics() (ok bool)
		NtDxgkNotifyWorkSubmission() (ok bool)
		NtDxgkOpenNativeFenceFromNtHandle() (ok bool)
		NtDxgkOutputDuplPresentToHwQueue() (ok bool)
		NtDxgkPinResources() (ok bool)
		NtDxgkRegisterVailProcess() (ok bool)
		NtDxgkResetTrackedWorkloadStatistics() (ok bool)
		NtDxgkSetProperties() (ok bool)
		NtDxgkSubmitPresentBltToHwQueue() (ok bool)
		NtDxgkSubmitPresentToHwQueue() (ok bool)
		NtDxgkUnpinResources() (ok bool)
		NtDxgkUpdateTrackedWorkload() (ok bool)
		NtDxgkVailConnect() (ok bool)
		NtDxgkVailDisconnect() (ok bool)
		NtDxgkVailPromoteCompositionSurface() (ok bool)
		NtEnableOneCoreTransformMode() (ok bool)
		NtFlipObjectAddContent() (ok bool)
		NtFlipObjectAddPoolBuffer() (ok bool)
		NtFlipObjectConsumerAcquirePresent() (ok bool)
		NtFlipObjectConsumerAdjustUsageReference() (ok bool)
		NtFlipObjectConsumerBeginProcessPresent() (ok bool)
		NtFlipObjectConsumerEndProcessPresent() (ok bool)
		NtFlipObjectConsumerPostMessage() (ok bool)
		NtFlipObjectConsumerQueryBufferInfo() (ok bool)
		NtFlipObjectCreate() (ok bool)
		NtFlipObjectDisconnectEndpoint() (ok bool)
		NtFlipObjectEnablePresentStatisticsType() (ok bool)
		NtFlipObjectOpen() (ok bool)
		NtFlipObjectPresentCancel() (ok bool)
		NtFlipObjectQueryBufferAvailableEvent() (ok bool)
		NtFlipObjectQueryEndpointConnected() (ok bool)
		NtFlipObjectQueryLostEvent() (ok bool)
		NtFlipObjectQueryNextMessageToProducer() (ok bool)
		NtFlipObjectReadNextMessageToProducer() (ok bool)
		NtFlipObjectRemoveContent() (ok bool)
		NtFlipObjectRemovePoolBuffer() (ok bool)
		NtFlipObjectSetContent() (ok bool)
		NtGdiAbortDoc() (ok bool)
		NtGdiAbortPath() (ok bool)
		NtGdiAddEmbFontToDC() (ok bool)
		NtGdiAddFontMemResourceEx() (ok bool)
		NtGdiAddFontResourceW() (ok bool)
		NtGdiAddInitialFonts() (ok bool)
		NtGdiAddRemoteFontToDC() (ok bool)
		NtGdiAddRemoteMMInstanceToDC() (ok bool)
		NtGdiAlphaBlend() (ok bool)
		NtGdiAngleArc() (ok bool)
		NtGdiAnyLinkedFonts() (ok bool)
		NtGdiArcInternal() (ok bool)
		NtGdiBRUSHOBJ_DeleteRbrush() (ok bool)
		NtGdiBRUSHOBJ_hGetColorTransform() (ok bool)
		NtGdiBRUSHOBJ_pvAllocRbrush() (ok bool)
		NtGdiBRUSHOBJ_pvGetRbrush() (ok bool)
		NtGdiBRUSHOBJ_ulGetBrushColor() (ok bool)
		NtGdiBeginGdiRendering() (ok bool)
		NtGdiBeginPath() (ok bool)
		NtGdiBitBlt() (ok bool)
		NtGdiCLIPOBJ_bEnum() (ok bool)
		NtGdiCLIPOBJ_cEnumStart() (ok bool)
		NtGdiCLIPOBJ_ppoGetPath() (ok bool)
		NtGdiCancelDC() (ok bool)
		NtGdiChangeGhostFont() (ok bool)
		NtGdiCheckBitmapBits() (ok bool)
		NtGdiClearBitmapAttributes() (ok bool)
		NtGdiClearBrushAttributes() (ok bool)
		NtGdiCloseFigure() (ok bool)
		NtGdiColorCorrectPalette() (ok bool)
		NtGdiCombineRgn() (ok bool)
		NtGdiCombineTransform() (ok bool)
		NtGdiComputeXformCoefficients() (ok bool)
		NtGdiConfigureOPMProtectedOutput() (ok bool)
		NtGdiConvertMetafileRect() (ok bool)
		NtGdiCreateBitmap() (ok bool)
		NtGdiCreateBitmapFromDxSurface() (ok bool)
		NtGdiCreateBitmapFromDxSurface2() (ok bool)
		NtGdiCreateClientObj() (ok bool)
		NtGdiCreateColorSpace() (ok bool)
		NtGdiCreateColorTransform() (ok bool)
		NtGdiCreateCompatibleBitmap() (ok bool)
		NtGdiCreateCompatibleDC() (ok bool)
		NtGdiCreateDIBBrush() (ok bool)
		NtGdiCreateDIBSection() (ok bool)
		NtGdiCreateDIBitmapInternal() (ok bool)
		NtGdiCreateEllipticRgn() (ok bool)
		NtGdiCreateHalftonePalette() (ok bool)
		NtGdiCreateHatchBrushInternal() (ok bool)
		NtGdiCreateMetafileDC() (ok bool)
		NtGdiCreateOPMProtectedOutput() (ok bool)
		NtGdiCreateOPMProtectedOutputs() (ok bool)
		NtGdiCreatePaletteInternal() (ok bool)
		NtGdiCreatePatternBrushInternal() (ok bool)
		NtGdiCreatePen() (ok bool)
		NtGdiCreateRectRgn() (ok bool)
		NtGdiCreateRoundRectRgn() (ok bool)
		NtGdiCreateServerMetaFile() (ok bool)
		NtGdiCreateSessionMappedDIBSection() (ok bool)
		NtGdiCreateSolidBrush() (ok bool)
		NtGdiDDCCIGetCapabilitiesString() (ok bool)
		NtGdiDDCCIGetCapabilitiesStringLength() (ok bool)
		NtGdiDDCCIGetTimingReport() (ok bool)
		NtGdiDDCCIGetVCPFeature() (ok bool)
		NtGdiDDCCISaveCurrentSettings() (ok bool)
		NtGdiDDCCISetVCPFeature() (ok bool)
		NtGdiDdCreateFullscreenSprite() (ok bool)
		NtGdiDdDDIAbandonSwapChain() (ok bool)
		NtGdiDdDDIAcquireKeyedMutex() (ok bool)
		NtGdiDdDDIAcquireKeyedMutex2() (ok bool)
		NtGdiDdDDIAcquireSwapChain() (ok bool)
		NtGdiDdDDIAddSurfaceToSwapChain() (ok bool)
		NtGdiDdDDIAdjustFullscreenGamma() (ok bool)
		NtGdiDdDDICacheHybridQueryValue() (ok bool)
		NtGdiDdDDIChangeVideoMemoryReservation() (ok bool)
		NtGdiDdDDICheckExclusiveOwnership() (ok bool)
		NtGdiDdDDICheckMonitorPowerState() (ok bool)
		NtGdiDdDDICheckMultiPlaneOverlaySupport() (ok bool)
		NtGdiDdDDICheckMultiPlaneOverlaySupport2() (ok bool)
		NtGdiDdDDICheckMultiPlaneOverlaySupport3() (ok bool)
		NtGdiDdDDICheckOcclusion() (ok bool)
		NtGdiDdDDICheckSharedResourceAccess() (ok bool)
		NtGdiDdDDICheckVidPnExclusiveOwnership() (ok bool)
		NtGdiDdDDICloseAdapter() (ok bool)
		NtGdiDdDDIConfigureSharedResource() (ok bool)
		NtGdiDdDDICreateAllocation() (ok bool)
		NtGdiDdDDICreateBundleObject() (ok bool)
		NtGdiDdDDICreateContext() (ok bool)
		NtGdiDdDDICreateContextVirtual() (ok bool)
		NtGdiDdDDICreateDCFromMemory() (ok bool)
		NtGdiDdDDICreateDevice() (ok bool)
		NtGdiDdDDICreateHwContext() (ok bool)
		NtGdiDdDDICreateHwQueue() (ok bool)
		NtGdiDdDDICreateKeyedMutex() (ok bool)
		NtGdiDdDDICreateKeyedMutex2() (ok bool)
		NtGdiDdDDICreateOutputDupl() (ok bool)
		NtGdiDdDDICreateOverlay() (ok bool)
		NtGdiDdDDICreatePagingQueue() (ok bool)
		NtGdiDdDDICreateProtectedSession() (ok bool)
		NtGdiDdDDICreateSwapChain() (ok bool)
		NtGdiDdDDICreateSynchronizationObject() (ok bool)
		NtGdiDdDDIDDisplayEnum() (ok bool)
		NtGdiDdDDIDestroyAllocation() (ok bool)
		NtGdiDdDDIDestroyAllocation2() (ok bool)
		NtGdiDdDDIDestroyContext() (ok bool)
		NtGdiDdDDIDestroyDCFromMemory() (ok bool)
		NtGdiDdDDIDestroyDevice() (ok bool)
		NtGdiDdDDIDestroyHwContext() (ok bool)
		NtGdiDdDDIDestroyHwQueue() (ok bool)
		NtGdiDdDDIDestroyKeyedMutex() (ok bool)
		NtGdiDdDDIDestroyOutputDupl() (ok bool)
		NtGdiDdDDIDestroyOverlay() (ok bool)
		NtGdiDdDDIDestroyPagingQueue() (ok bool)
		NtGdiDdDDIDestroyProtectedSession() (ok bool)
		NtGdiDdDDIDestroySynchronizationObject() (ok bool)
		NtGdiDdDDIDispMgrCreate() (ok bool)
		NtGdiDdDDIDispMgrSourceOperation() (ok bool)
		NtGdiDdDDIDispMgrTargetOperation() (ok bool)
		NtGdiDdDDIEnumAdapters() (ok bool)
		NtGdiDdDDIEnumAdapters2() (ok bool)
		NtGdiDdDDIEscape() (ok bool)
		NtGdiDdDDIEvict() (ok bool)
		NtGdiDdDDIExtractBundleObject() (ok bool)
		NtGdiDdDDIFlipOverlay() (ok bool)
		NtGdiDdDDIFlushHeapTransitions() (ok bool)
		NtGdiDdDDIFreeGpuVirtualAddress() (ok bool)
		NtGdiDdDDIGetAllocationPriority() (ok bool)
		NtGdiDdDDIGetCachedHybridQueryValue() (ok bool)
		NtGdiDdDDIGetContextInProcessSchedulingPriority() (ok bool)
		NtGdiDdDDIGetContextSchedulingPriority() (ok bool)
		NtGdiDdDDIGetDWMVerticalBlankEvent() (ok bool)
		NtGdiDdDDIGetDeviceState() (ok bool)
		NtGdiDdDDIGetDisplayModeList() (ok bool)
		NtGdiDdDDIGetMemoryBudgetTarget() (ok bool)
		NtGdiDdDDIGetMultiPlaneOverlayCaps() (ok bool)
		NtGdiDdDDIGetMultisampleMethodList() (ok bool)
		NtGdiDdDDIGetOverlayState() (ok bool)
		NtGdiDdDDIGetPostCompositionCaps() (ok bool)
		NtGdiDdDDIGetPresentHistory() (ok bool)
		NtGdiDdDDIGetPresentQueueEvent() (ok bool)
		NtGdiDdDDIGetProcessDeviceRemovalSupport() (ok bool)
		NtGdiDdDDIGetProcessSchedulingPriorityBand() (ok bool)
		NtGdiDdDDIGetProcessSchedulingPriorityClass() (ok bool)
		NtGdiDdDDIGetResourcePresentPrivateDriverData() (ok bool)
		NtGdiDdDDIGetRuntimeData() (ok bool)
		NtGdiDdDDIGetScanLine() (ok bool)
		NtGdiDdDDIGetSetSwapChainMetadata() (ok bool)
		NtGdiDdDDIGetSharedPrimaryHandle() (ok bool)
		NtGdiDdDDIGetSharedResourceAdapterLuid() (ok bool)
		NtGdiDdDDIGetSharedResourceAdapterLuidFlipManager() (ok bool)
		NtGdiDdDDIGetSwapChainSurfacePhysicalAddress() (ok bool)
		NtGdiDdDDIGetYieldPercentage() (ok bool)
		NtGdiDdDDIInvalidateActiveVidPn() (ok bool)
		NtGdiDdDDIInvalidateCache() (ok bool)
		NtGdiDdDDILock() (ok bool)
		NtGdiDdDDILock2() (ok bool)
		NtGdiDdDDIMakeResident() (ok bool)
		NtGdiDdDDIMapGpuVirtualAddress() (ok bool)
		NtGdiDdDDIMarkDeviceAsError() (ok bool)
		NtGdiDdDDINetDispGetNextChunkInfo() (ok bool)
		NtGdiDdDDINetDispQueryMiracastDisplayDeviceStatus() (ok bool)
		NtGdiDdDDINetDispQueryMiracastDisplayDeviceSupport() (ok bool)
		NtGdiDdDDINetDispStartMiracastDisplayDevice() (ok bool)
		NtGdiDdDDINetDispStopMiracastDisplayDevice() (ok bool)
		NtGdiDdDDIOfferAllocations() (ok bool)
		NtGdiDdDDIOpenAdapterFromDeviceName() (ok bool)
		NtGdiDdDDIOpenAdapterFromHdc() (ok bool)
		NtGdiDdDDIOpenAdapterFromLuid() (ok bool)
		NtGdiDdDDIOpenBundleObjectNtHandleFromName() (ok bool)
		NtGdiDdDDIOpenKeyedMutex() (ok bool)
		NtGdiDdDDIOpenKeyedMutex2() (ok bool)
		NtGdiDdDDIOpenKeyedMutexFromNtHandle() (ok bool)
		NtGdiDdDDIOpenNtHandleFromName() (ok bool)
		NtGdiDdDDIOpenProtectedSessionFromNtHandle() (ok bool)
		NtGdiDdDDIOpenResource() (ok bool)
		NtGdiDdDDIOpenResourceFromNtHandle() (ok bool)
		NtGdiDdDDIOpenSwapChain() (ok bool)
		NtGdiDdDDIOpenSyncObjectFromNtHandle() (ok bool)
		NtGdiDdDDIOpenSyncObjectFromNtHandle2() (ok bool)
		NtGdiDdDDIOpenSyncObjectNtHandleFromName() (ok bool)
		NtGdiDdDDIOpenSynchronizationObject() (ok bool)
		NtGdiDdDDIOutputDuplGetFrameInfo() (ok bool)
		NtGdiDdDDIOutputDuplGetMetaData() (ok bool)
		NtGdiDdDDIOutputDuplGetPointerShapeData() (ok bool)
		NtGdiDdDDIOutputDuplPresent() (ok bool)
		NtGdiDdDDIOutputDuplReleaseFrame() (ok bool)
		NtGdiDdDDIPollDisplayChildren() (ok bool)
		NtGdiDdDDIPresent() (ok bool)
		NtGdiDdDDIPresentMultiPlaneOverlay() (ok bool)
		NtGdiDdDDIPresentMultiPlaneOverlay2() (ok bool)
		NtGdiDdDDIPresentMultiPlaneOverlay3() (ok bool)
		NtGdiDdDDIPresentRedirected() (ok bool)
		NtGdiDdDDIQueryAdapterInfo() (ok bool)
		NtGdiDdDDIQueryAllocationResidency() (ok bool)
		NtGdiDdDDIQueryClockCalibration() (ok bool)
		NtGdiDdDDIQueryFSEBlock() (ok bool)
		NtGdiDdDDIQueryProcessOfferInfo() (ok bool)
		NtGdiDdDDIQueryProtectedSessionInfoFromNtHandle() (ok bool)
		NtGdiDdDDIQueryProtectedSessionStatus() (ok bool)
		NtGdiDdDDIQueryRemoteVidPnSourceFromGdiDisplayName() (ok bool)
		NtGdiDdDDIQueryResourceInfo() (ok bool)
		NtGdiDdDDIQueryResourceInfoFromNtHandle() (ok bool)
		NtGdiDdDDIQueryStatistics() (ok bool)
		NtGdiDdDDIQueryVidPnExclusiveOwnership() (ok bool)
		NtGdiDdDDIQueryVideoMemoryInfo() (ok bool)
		NtGdiDdDDIReclaimAllocations() (ok bool)
		NtGdiDdDDIReclaimAllocations2() (ok bool)
		NtGdiDdDDIReleaseKeyedMutex() (ok bool)
		NtGdiDdDDIReleaseKeyedMutex2() (ok bool)
		NtGdiDdDDIReleaseProcessVidPnSourceOwners() (ok bool)
		NtGdiDdDDIReleaseSwapChain() (ok bool)
		NtGdiDdDDIRemoveSurfaceFromSwapChain() (ok bool)
		NtGdiDdDDIRender() (ok bool)
		NtGdiDdDDIReserveGpuVirtualAddress() (ok bool)
		NtGdiDdDDISetAllocationPriority() (ok bool)
		NtGdiDdDDISetContextInProcessSchedulingPriority() (ok bool)
		NtGdiDdDDISetContextSchedulingPriority() (ok bool)
		NtGdiDdDDISetDisplayMode() (ok bool)
		NtGdiDdDDISetDodIndirectSwapchain() (ok bool)
		NtGdiDdDDISetFSEBlock() (ok bool)
		NtGdiDdDDISetGammaRamp() (ok bool)
		NtGdiDdDDISetHwProtectionTeardownRecovery() (ok bool)
		NtGdiDdDDISetMemoryBudgetTarget() (ok bool)
		NtGdiDdDDISetMonitorColorSpaceTransform() (ok bool)
		NtGdiDdDDISetProcessDeviceRemovalSupport() (ok bool)
		NtGdiDdDDISetProcessSchedulingPriorityBand() (ok bool)
		NtGdiDdDDISetProcessSchedulingPriorityClass() (ok bool)
		NtGdiDdDDISetQueuedLimit() (ok bool)
		NtGdiDdDDISetStablePowerState() (ok bool)
		NtGdiDdDDISetStereoEnabled() (ok bool)
		NtGdiDdDDISetSyncRefreshCountWaitTarget() (ok bool)
		NtGdiDdDDISetVidPnSourceHwProtection() (ok bool)
		NtGdiDdDDISetVidPnSourceOwner() (ok bool)
		NtGdiDdDDISetYieldPercentage() (ok bool)
		NtGdiDdDDIShareObjects() (ok bool)
		NtGdiDdDDISharedPrimaryLockNotification() (ok bool)
		NtGdiDdDDISharedPrimaryUnLockNotification() (ok bool)
		NtGdiDdDDISignalSynchronizationObject() (ok bool)
		NtGdiDdDDISignalSynchronizationObjectFromCpu() (ok bool)
		NtGdiDdDDISignalSynchronizationObjectFromGpu() (ok bool)
		NtGdiDdDDISignalSynchronizationObjectFromGpu2() (ok bool)
		NtGdiDdDDISubmitCommand() (ok bool)
		NtGdiDdDDISubmitCommandToHwQueue() (ok bool)
		NtGdiDdDDISubmitSignalSyncObjectsToHwQueue() (ok bool)
		NtGdiDdDDISubmitWaitForSyncObjectsToHwQueue() (ok bool)
		NtGdiDdDDITrimProcessCommitment() (ok bool)
		NtGdiDdDDIUnOrderedPresentSwapChain() (ok bool)
		NtGdiDdDDIUnlock() (ok bool)
		NtGdiDdDDIUnlock2() (ok bool)
		NtGdiDdDDIUpdateAllocationProperty() (ok bool)
		NtGdiDdDDIUpdateGpuVirtualAddress() (ok bool)
		NtGdiDdDDIUpdateOverlay() (ok bool)
		NtGdiDdDDIWaitForIdle() (ok bool)
		NtGdiDdDDIWaitForSynchronizationObject() (ok bool)
		NtGdiDdDDIWaitForSynchronizationObjectFromCpu() (ok bool)
		NtGdiDdDDIWaitForSynchronizationObjectFromGpu() (ok bool)
		NtGdiDdDDIWaitForVerticalBlankEvent() (ok bool)
		NtGdiDdDDIWaitForVerticalBlankEvent2() (ok bool)
		NtGdiDdDestroyFullscreenSprite() (ok bool)
		NtGdiDdNotifyFullscreenSpriteUpdate() (ok bool)
		NtGdiDdQueryVisRgnUniqueness() (ok bool)
		NtGdiDeleteClientObj() (ok bool)
		NtGdiDeleteColorSpace() (ok bool)
		NtGdiDeleteColorTransform() (ok bool)
		NtGdiDeleteObjectApp() (ok bool)
		NtGdiDescribePixelFormat() (ok bool)
		NtGdiDestroyOPMProtectedOutput() (ok bool)
		NtGdiDestroyPhysicalMonitor() (ok bool)
		NtGdiDoBanding() (ok bool)
		NtGdiDoPalette() (ok bool)
		NtGdiDrawEscape() (ok bool)
		NtGdiDrawStream() (ok bool)
		NtGdiDwmCreatedBitmapRemotingOutput() (ok bool)
		NtGdiEllipse() (ok bool)
		NtGdiEnableEudc() (ok bool)
		NtGdiEndDoc() (ok bool)
		NtGdiEndGdiRendering() (ok bool)
		NtGdiEndPage() (ok bool)
		NtGdiEndPath() (ok bool)
		NtGdiEngAlphaBlend() (ok bool)
		NtGdiEngAssociateSurface() (ok bool)
		NtGdiEngBitBlt() (ok bool)
		NtGdiEngCheckAbort() (ok bool)
		NtGdiEngComputeGlyphSet() (ok bool)
		NtGdiEngCopyBits() (ok bool)
		NtGdiEngCreateBitmap() (ok bool)
		NtGdiEngCreateClip() (ok bool)
		NtGdiEngCreateDeviceBitmap() (ok bool)
		NtGdiEngCreateDeviceSurface() (ok bool)
		NtGdiEngCreatePalette() (ok bool)
		NtGdiEngDeleteClip() (ok bool)
		NtGdiEngDeletePalette() (ok bool)
		NtGdiEngDeletePath() (ok bool)
		NtGdiEngDeleteSurface() (ok bool)
		NtGdiEngEraseSurface() (ok bool)
		NtGdiEngFillPath() (ok bool)
		NtGdiEngGradientFill() (ok bool)
		NtGdiEngLineTo() (ok bool)
		NtGdiEngLockSurface() (ok bool)
		NtGdiEngMarkBandingSurface() (ok bool)
		NtGdiEngPaint() (ok bool)
		NtGdiEngPlgBlt() (ok bool)
		NtGdiEngStretchBlt() (ok bool)
		NtGdiEngStretchBltROP() (ok bool)
		NtGdiEngStrokeAndFillPath() (ok bool)
		NtGdiEngStrokePath() (ok bool)
		NtGdiEngTextOut() (ok bool)
		NtGdiEngTransparentBlt() (ok bool)
		NtGdiEngUnlockSurface() (ok bool)
		NtGdiEnsureDpiDepDefaultGuiFontForPlateau() (ok bool)
		NtGdiEnumFonts() (ok bool)
		NtGdiEnumObjects() (ok bool)
		NtGdiEqualRgn() (ok bool)
		NtGdiEudcLoadUnloadLink() (ok bool)
		NtGdiExcludeClipRect() (ok bool)
		NtGdiExtCreatePen() (ok bool)
		NtGdiExtCreateRegion() (ok bool)
		NtGdiExtEscape() (ok bool)
		NtGdiExtFloodFill() (ok bool)
		NtGdiExtGetObjectW() (ok bool)
		NtGdiExtSelectClipRgn() (ok bool)
		NtGdiExtTextOutW() (ok bool)
		NtGdiFONTOBJ_cGetAllGlyphHandles() (ok bool)
		NtGdiFONTOBJ_cGetGlyphs() (ok bool)
		NtGdiFONTOBJ_pQueryGlyphAttrs() (ok bool)
		NtGdiFONTOBJ_pfdg() (ok bool)
		NtGdiFONTOBJ_pifi() (ok bool)
		NtGdiFONTOBJ_pvTrueTypeFontFile() (ok bool)
		NtGdiFONTOBJ_pxoGetXform() (ok bool)
		NtGdiFONTOBJ_vGetInfo() (ok bool)
		NtGdiFillPath() (ok bool)
		NtGdiFillRgn() (ok bool)
		NtGdiFlattenPath() (ok bool)
		NtGdiFlush() (ok bool)
		NtGdiFontIsLinked() (ok bool)
		NtGdiForceUFIMapping() (ok bool)
		NtGdiFrameRgn() (ok bool)
		NtGdiFullscreenControl() (ok bool)
		NtGdiGetAndSetDCDword() (ok bool)
		NtGdiGetAppClipBox() (ok bool)
		NtGdiGetBitmapBits() (ok bool)
		NtGdiGetBitmapDimension() (ok bool)
		NtGdiGetBitmapDpiScaleValue() (ok bool)
		NtGdiGetBoundsRect() (ok bool)
		NtGdiGetCOPPCompatibleOPMInformation() (ok bool)
		NtGdiGetCertificate() (ok bool)
		NtGdiGetCertificateByHandle() (ok bool)
		NtGdiGetCertificateSize() (ok bool)
		NtGdiGetCertificateSizeByHandle() (ok bool)
		NtGdiGetCharABCWidthsW() (ok bool)
		NtGdiGetCharSet() (ok bool)
		NtGdiGetCharWidthInfo() (ok bool)
		NtGdiGetCharWidthW() (ok bool)
		NtGdiGetCharacterPlacementW() (ok bool)
		NtGdiGetColorAdjustment() (ok bool)
		NtGdiGetColorSpaceforBitmap() (ok bool)
		NtGdiGetCurrentDpiInfo() (ok bool)
		NtGdiGetDCDpiScaleValue() (ok bool)
		NtGdiGetDCDword() (ok bool)
		NtGdiGetDCObject() (ok bool)
		NtGdiGetDCPoint() (ok bool)
		NtGdiGetDCforBitmap() (ok bool)
		NtGdiGetDIBitsInternal() (ok bool)
		NtGdiGetDeviceCaps() (ok bool)
		NtGdiGetDeviceCapsAll() (ok bool)
		NtGdiGetDeviceWidth() (ok bool)
		NtGdiGetDhpdev() (ok bool)
		NtGdiGetETM() (ok bool)
		NtGdiGetEmbUFI() (ok bool)
		NtGdiGetEmbedFonts() (ok bool)
		NtGdiGetEntry() (ok bool)
		NtGdiGetEudcTimeStampEx() (ok bool)
		NtGdiGetFontData() (ok bool)
		NtGdiGetFontFileData() (ok bool)
		NtGdiGetFontFileInfo() (ok bool)
		NtGdiGetFontResourceInfoInternalW() (ok bool)
		NtGdiGetFontUnicodeRanges() (ok bool)
		NtGdiGetGlyphIndicesW() (ok bool)
		NtGdiGetGlyphIndicesWInternal() (ok bool)
		NtGdiGetGlyphOutline() (ok bool)
		NtGdiGetKerningPairs() (ok bool)
		NtGdiGetLinkedUFIs() (ok bool)
		NtGdiGetMiterLimit() (ok bool)
		NtGdiGetMonitorID() (ok bool)
		NtGdiGetNearestColor() (ok bool)
		NtGdiGetNearestPaletteIndex() (ok bool)
		NtGdiGetNumberOfPhysicalMonitors() (ok bool)
		NtGdiGetOPMInformation() (ok bool)
		NtGdiGetOPMRandomNumber() (ok bool)
		NtGdiGetObjectBitmapHandle() (ok bool)
		NtGdiGetOutlineTextMetricsInternalW() (ok bool)
		NtGdiGetPath() (ok bool)
		NtGdiGetPerBandInfo() (ok bool)
		NtGdiGetPhysicalMonitorDescription() (ok bool)
		NtGdiGetPhysicalMonitors() (ok bool)
		NtGdiGetPixel() (ok bool)
		NtGdiGetProcessSessionFonts() (ok bool)
		NtGdiGetPublicFontTableChangeCookie() (ok bool)
		NtGdiGetRandomRgn() (ok bool)
		NtGdiGetRasterizerCaps() (ok bool)
		NtGdiGetRealizationInfo() (ok bool)
		NtGdiGetRegionData() (ok bool)
		NtGdiGetRgnBox() (ok bool)
		NtGdiGetServerMetaFileBits() (ok bool)
		NtGdiGetSpoolMessage() (ok bool)
		NtGdiGetStats() (ok bool)
		NtGdiGetStringBitmapW() (ok bool)
		NtGdiGetSuggestedOPMProtectedOutputArraySize() (ok bool)
		NtGdiGetSystemPaletteUse() (ok bool)
		NtGdiGetTextCharsetInfo() (ok bool)
		NtGdiGetTextExtent() (ok bool)
		NtGdiGetTextExtentExW() (ok bool)
		NtGdiGetTextFaceW() (ok bool)
		NtGdiGetTextMetricsW() (ok bool)
		NtGdiGetTransform() (ok bool)
		NtGdiGetUFI() (ok bool)
		NtGdiGetUFIPathname() (ok bool)
		NtGdiGetWidthTable() (ok bool)
		NtGdiGradientFill() (ok bool)
		NtGdiHLSurfGetInformation() (ok bool)
		NtGdiHLSurfSetInformation() (ok bool)
		NtGdiHT_Get8BPPFormatPalette() (ok bool)
		NtGdiHT_Get8BPPMaskPalette() (ok bool)
		NtGdiHfontCreate() (ok bool)
		NtGdiIcmBrushInfo() (ok bool)
		NtGdiInit() (ok bool)
		NtGdiInitSpool() (ok bool)
		NtGdiIntersectClipRect() (ok bool)
		NtGdiInvertRgn() (ok bool)
		NtGdiLineTo() (ok bool)
		NtGdiMakeFontDir() (ok bool)
		NtGdiMakeInfoDC() (ok bool)
		NtGdiMakeObjectUnXferable() (ok bool)
		NtGdiMakeObjectXferable() (ok bool)
		NtGdiMaskBlt() (ok bool)
		NtGdiMirrorWindowOrg() (ok bool)
		NtGdiModifyWorldTransform() (ok bool)
		NtGdiMonoBitmap() (ok bool)
		NtGdiMoveTo() (ok bool)
		NtGdiOffsetClipRgn() (ok bool)
		NtGdiOffsetRgn() (ok bool)
		NtGdiOpenDCW() (ok bool)
		NtGdiPATHOBJ_bEnum() (ok bool)
		NtGdiPATHOBJ_bEnumClipLines() (ok bool)
		NtGdiPATHOBJ_vEnumStart() (ok bool)
		NtGdiPATHOBJ_vEnumStartClipLines() (ok bool)
		NtGdiPATHOBJ_vGetBounds() (ok bool)
		NtGdiPatBlt() (ok bool)
		NtGdiPathToRegion() (ok bool)
		NtGdiPlgBlt() (ok bool)
		NtGdiPolyDraw() (ok bool)
		NtGdiPolyPatBlt() (ok bool)
		NtGdiPolyPolyDraw() (ok bool)
		NtGdiPolyTextOutW() (ok bool)
		NtGdiPtInRegion() (ok bool)
		NtGdiPtVisible() (ok bool)
		NtGdiQueryFontAssocInfo() (ok bool)
		NtGdiQueryFonts() (ok bool)
		NtGdiRectInRegion() (ok bool)
		NtGdiRectVisible() (ok bool)
		NtGdiRectangle() (ok bool)
		NtGdiRemoveFontMemResourceEx() (ok bool)
		NtGdiRemoveFontResourceW() (ok bool)
		NtGdiRemoveMergeFont() (ok bool)
		NtGdiResetDC() (ok bool)
		NtGdiResizePalette() (ok bool)
		NtGdiRestoreDC() (ok bool)
		NtGdiRoundRect() (ok bool)
		NtGdiSTROBJ_bEnum() (ok bool)
		NtGdiSTROBJ_bEnumPositionsOnly() (ok bool)
		NtGdiSTROBJ_bGetAdvanceWidths() (ok bool)
		NtGdiSTROBJ_dwGetCodePage() (ok bool)
		NtGdiSTROBJ_vEnumStart() (ok bool)
		NtGdiSaveDC() (ok bool)
		NtGdiScaleRgn() (ok bool)
		NtGdiScaleValues() (ok bool)
		NtGdiScaleViewportExtEx() (ok bool)
		NtGdiScaleWindowExtEx() (ok bool)
		NtGdiSelectBitmap() (ok bool)
		NtGdiSelectBrush() (ok bool)
		NtGdiSelectClipPath() (ok bool)
		NtGdiSelectFont() (ok bool)
		NtGdiSelectPen() (ok bool)
		NtGdiSetBitmapAttributes() (ok bool)
		NtGdiSetBitmapBits() (ok bool)
		NtGdiSetBitmapDimension() (ok bool)
		NtGdiSetBoundsRect() (ok bool)
		NtGdiSetBrushAttributes() (ok bool)
		NtGdiSetBrushOrg() (ok bool)
		NtGdiSetColorAdjustment() (ok bool)
		NtGdiSetColorSpace() (ok bool)
		NtGdiSetDIBitsToDeviceInternal() (ok bool)
		NtGdiSetFontEnumeration() (ok bool)
		NtGdiSetFontXform() (ok bool)
		NtGdiSetIcmMode() (ok bool)
		NtGdiSetLayout() (ok bool)
		NtGdiSetLinkedUFIs() (ok bool)
		NtGdiSetMagicColors() (ok bool)
		NtGdiSetMetaRgn() (ok bool)
		NtGdiSetMiterLimit() (ok bool)
		NtGdiSetOPMSigningKeyAndSequenceNumbers() (ok bool)
		NtGdiSetPUMPDOBJ() (ok bool)
		NtGdiSetPixel() (ok bool)
		NtGdiSetPixelFormat() (ok bool)
		NtGdiSetRectRgn() (ok bool)
		NtGdiSetSizeDevice() (ok bool)
		NtGdiSetSystemPaletteUse() (ok bool)
		NtGdiSetTextJustification() (ok bool)
		NtGdiSetUMPDSandboxState() (ok bool)
		NtGdiSetVirtualResolution() (ok bool)
		NtGdiStartDoc() (ok bool)
		NtGdiStartPage() (ok bool)
		NtGdiStretchBlt() (ok bool)
		NtGdiStretchDIBitsInternal() (ok bool)
		NtGdiStrokeAndFillPath() (ok bool)
		NtGdiStrokePath() (ok bool)
		NtGdiSwapBuffers() (ok bool)
		NtGdiTransformPoints() (ok bool)
		NtGdiTransparentBlt() (ok bool)
		NtGdiUMPDEngFreeUserMem() (ok bool)
		NtGdiUnloadPrinterDriver() (ok bool)
		NtGdiUnmapMemFont() (ok bool)
		NtGdiUnrealizeObject() (ok bool)
		NtGdiUpdateColors() (ok bool)
		NtGdiUpdateTransform() (ok bool)
		NtGdiWaitForTextReady() (ok bool)
		NtGdiWidenPath() (ok bool)
		NtGdiXFORMOBJ_bApplyXform() (ok bool)
		NtGdiXFORMOBJ_iGetXform() (ok bool)
		NtGdiXLATEOBJ_cGetPalette() (ok bool)
		NtGdiXLATEOBJ_hGetColorTransform() (ok bool)
		NtGdiXLATEOBJ_iXlate() (ok bool)
		NtHWCursorUpdatePointer() (ok bool)
		NtInputSpaceRegionFromPoint() (ok bool)
		NtIsOneCoreTransformMode() (ok bool)
		NtKSTInitialize() (ok bool)
		NtKSTWait() (ok bool)
		NtMITAccessibilityTimerNotification() (ok bool)
		NtMITActivateInputProcessing() (ok bool)
		NtMITConfigureVirtualTouchpad() (ok bool)
		NtMITCoreMsgKOpenConnectionTo() (ok bool)
		NtMITDeactivateInputProcessing() (ok bool)
		NtMITDisableMouseIntercept() (ok bool)
		NtMITDispatchCompletion() (ok bool)
		NtMITEnableMouseIntercept() (ok bool)
		NtMITGetCursorUpdateHandle() (ok bool)
		NtMITInitMinuserThread() (ok bool)
		NtMITMinuserSetInputTransformOffset() (ok bool)
		NtMITMinuserWindowDestroyed() (ok bool)
		NtMITPostMouseInputMessage() (ok bool)
		NtMITPostThreadEventMessage() (ok bool)
		NtMITPostWindowEventMessage() (ok bool)
		NtMITPrepareReceiveInputMessage() (ok bool)
		NtMITPrepareSendInputMessage() (ok bool)
		NtMITProcessDelegateCapturedPointers() (ok bool)
		NtMITSetInputCallbacks() (ok bool)
		NtMITSetInputDelegationMode() (ok bool)
		NtMITSetInputObservationState() (ok bool)
		NtMITSetKeyboardInputRoutingPolicy() (ok bool)
		NtMITSetKeyboardOverriderState() (ok bool)
		NtMITSetLastInputRecipient() (ok bool)
		NtMITSynthesizeKeyboardInput() (ok bool)
		NtMITSynthesizeMouseInput() (ok bool)
		NtMITSynthesizeTouchInput() (ok bool)
		NtMITUninitMinuserThread() (ok bool)
		NtMITUpdateInputGlobals() (ok bool)
		NtMapVisualRelativePoints() (ok bool)
		NtMinGetInputTransform() (ok bool)
		NtMinInteropCoreMessagingWithInput() (ok bool)
		NtMinQPeekForInput() (ok bool)
		NtMinQSuspendInputProcessing() (ok bool)
		NtMinQUpdateWakeMask() (ok bool)
		NtModerncoreBeginLayoutUpdate() (ok bool)
		NtModerncoreCreateDCompositionHwndTarget() (ok bool)
		NtModerncoreCreateGDIHwndTarget() (ok bool)
		NtModerncoreDestroyDCompositionHwndTarget() (ok bool)
		NtModerncoreDestroyGDIHwndTarget() (ok bool)
		NtModerncoreEnableResizeLayoutSynchronization() (ok bool)
		NtModerncoreGetNavigationWindowVisual() (ok bool)
		NtModerncoreGetResizeDCompositionSynchronizationObject() (ok bool)
		NtModerncoreGetWindowContentVisual() (ok bool)
		NtModerncoreIdleTimerThread() (ok bool)
		NtModerncoreIsResizeLayoutSynchronizationEnabled() (ok bool)
		NtModerncoreProcessConnect() (ok bool)
		NtModerncoreRegisterEnhancedNavigationWindowHandle() (ok bool)
		NtModerncoreRegisterNavigationWindowHandle() (ok bool)
		NtModerncoreSetNavigationServiceSid() (ok bool)
		NtModerncoreUnregisterNavigationWindowHandle() (ok bool)
		NtNotifyPresentToCompositionSurface() (ok bool)
		NtOpenCompositionSurfaceDirtyRegion() (ok bool)
		NtOpenCompositionSurfaceRealizationInfo() (ok bool)
		NtOpenCompositionSurfaceSectionInfo() (ok bool)
		NtQueryCompositionInputIsImplicit() (ok bool)
		NtQueryCompositionInputQueueAndTransform() (ok bool)
		NtQueryCompositionInputSink() (ok bool)
		NtQueryCompositionInputSinkLuid() (ok bool)
		NtQueryCompositionInputSinkViewId() (ok bool)
		NtQueryCompositionSurfaceBinding() (ok bool)
		NtQueryCompositionSurfaceFrameRate() (ok bool)
		NtQueryCompositionSurfaceHDRMetaData() (ok bool)
		NtQueryCompositionSurfaceRenderingRealization() (ok bool)
		NtQueryCompositionSurfaceStatistics() (ok bool)
		NtRIMAddInputObserver() (ok bool)
		NtRIMAreSiblingDevices() (ok bool)
		NtRIMDeviceIoControl() (ok bool)
		NtRIMEnableMonitorMappingForDevice() (ok bool)
		NtRIMFreeInputBuffer() (ok bool)
		NtRIMGetDevicePreparsedData() (ok bool)
		NtRIMGetDevicePreparsedDataLockfree() (ok bool)
		NtRIMGetDeviceProperties() (ok bool)
		NtRIMGetDevicePropertiesLockfree() (ok bool)
		NtRIMGetPhysicalDeviceRect() (ok bool)
		NtRIMGetSourceProcessId() (ok bool)
		NtRIMObserveNextInput() (ok bool)
		NtRIMOnAsyncPnpWorkNotification() (ok bool)
		NtRIMOnPnpNotification() (ok bool)
		NtRIMOnTimerNotification() (ok bool)
		NtRIMQueryDevicePath() (ok bool)
		NtRIMReadInput() (ok bool)
		NtRIMRegisterForInputEx() (ok bool)
		NtRIMRemoveInputObserver() (ok bool)
		NtRIMSetDeadzoneRotation() (ok bool)
		NtRIMSetExtendedDeviceProperty() (ok bool)
		NtRIMSetTestModeStatus() (ok bool)
		NtRIMUnregisterForInput() (ok bool)
		NtRIMUpdateInputObserverRegistration() (ok bool)
		NtSetCompositionSurfaceAnalogExclusive() (ok bool)
		NtSetCompositionSurfaceBufferUsage() (ok bool)
		NtSetCompositionSurfaceDirectFlipState() (ok bool)
		NtSetCompositionSurfaceIndependentFlipInfo() (ok bool)
		NtSetCompositionSurfaceStatistics() (ok bool)
		NtSetCursorInputSpace() (ok bool)
		NtSetPointerDeviceInputSpace() (ok bool)
		NtSetShellCursorState() (ok bool)
		NtTokenManagerConfirmOutstandingAnalogToken() (ok bool)
		NtTokenManagerCreateCompositionTokenHandle() (ok bool)
		NtTokenManagerCreateFlipObjectReturnTokenHandle() (ok bool)
		NtTokenManagerCreateFlipObjectTokenHandle() (ok bool)
		NtTokenManagerGetAnalogExclusiveSurfaceUpdates() (ok bool)
		NtTokenManagerGetAnalogExclusiveTokenEvent() (ok bool)
		NtTokenManagerOpenSectionAndEvents() (ok bool)
		NtTokenManagerThread() (ok bool)
		NtUnBindCompositionSurface() (ok bool)
		NtUpdateInputSinkTransforms() (ok bool)
		NtUserAcquireIAMKey() (ok bool)
		NtUserAcquireInteractiveControlBackgroundAccess() (ok bool)
		NtUserActivateKeyboardLayout() (ok bool)
		NtUserAddClipboardFormatListener() (ok bool)
		NtUserAddVisualIdentifier() (ok bool)
		NtUserAllowSetForegroundWindow() (ok bool)
		NtUserAlterWindowStyle() (ok bool)
		NtUserArrangeIconicWindows() (ok bool)
		NtUserAssociateInputContext() (ok bool)
		NtUserAttachThreadInput() (ok bool)
		NtUserAutoPromoteMouseInPointer() (ok bool)
		NtUserAutoRotateScreen() (ok bool)
		NtUserBeginDeferWindowPos() (ok bool)
		NtUserBeginLayoutUpdate() (ok bool)
		NtUserBeginPaint() (ok bool)
		NtUserBitBltSysBmp() (ok bool)
		NtUserBlockInput() (ok bool)
		NtUserBroadcastImeShowStatusChange() (ok bool)
		NtUserBroadcastThemeChangeEvent() (ok bool)
		NtUserBuildHimcList() (ok bool)
		NtUserBuildHwndList() (ok bool)
		NtUserBuildNameList() (ok bool)
		NtUserBuildPropList() (ok bool)
		NtUserCalcMenuBar() (ok bool)
		NtUserCalculatePopupWindowPosition() (ok bool)
		NtUserCallMsgFilter() (ok bool)
		NtUserCallNextHookEx() (ok bool)
		NtUserCanBrokerForceForeground() (ok bool)
		NtUserCancelQueueEventCompletionPacket() (ok bool)
		NtUserChangeClipboardChain() (ok bool)
		NtUserChangeDisplaySettings() (ok bool)
		NtUserChangeWindowMessageFilter() (ok bool)
		NtUserChangeWindowMessageFilterEx() (ok bool)
		NtUserCheckAccessForIntegrityLevel() (ok bool)
		NtUserCheckImeShowStatusInThread() (ok bool)
		NtUserCheckMenuItem() (ok bool)
		NtUserCheckProcessForClipboardAccess() (ok bool)
		NtUserCheckProcessSession() (ok bool)
		NtUserCheckWindowThreadDesktop() (ok bool)
		NtUserChildWindowFromPointEx() (ok bool)
		NtUserCitSetInfo() (ok bool)
		NtUserClearForeground() (ok bool)
		NtUserClearRunWakeBit() (ok bool)
		NtUserClearWakeMask() (ok bool)
		NtUserClearWindowState() (ok bool)
		NtUserClipCursor() (ok bool)
		NtUserCloseClipboard() (ok bool)
		NtUserCloseDesktop() (ok bool)
		NtUserCloseWindowStation() (ok bool)
		NtUserCompositionInputSinkLuidFromPoint() (ok bool)
		NtUserCompositionInputSinkViewInstanceIdFromPoint() (ok bool)
		NtUserConfigureActivationObject() (ok bool)
		NtUserConfirmResizeCommit() (ok bool)
		NtUserConsoleControl() (ok bool)
		NtUserConvertMemHandle() (ok bool)
		NtUserCopyAcceleratorTable() (ok bool)
		NtUserCountClipboardFormats() (ok bool)
		NtUserCreateAcceleratorTable() (ok bool)
		NtUserCreateActivationObject() (ok bool)
		NtUserCreateBaseWindow() (ok bool)
		NtUserCreateCaret() (ok bool)
		NtUserCreateDCompositionHwndTarget() (ok bool)
		NtUserCreateDesktopEx() (ok bool)
		NtUserCreateEmptyCursorObject() (ok bool)
		NtUserCreateInputContext() (ok bool)
		NtUserCreateLocalMemHandle() (ok bool)
		NtUserCreateMenu() (ok bool)
		NtUserCreatePalmRejectionDelayZone() (ok bool)
		NtUserCreatePopupMenu() (ok bool)
		NtUserCreateSystemThreads() (ok bool)
		NtUserCreateWindowEx() (ok bool)
		NtUserCreateWindowStation() (ok bool)
		NtUserCsDdeUninitialize() (ok bool)
		NtUserCtxDisplayIOCtl() (ok bool)
		NtUserDWP_GetEnabledPopupOffset() (ok bool)
		NtUserDdeInitialize() (ok bool)
		NtUserDefSetText() (ok bool)
		NtUserDeferWindowDpiChanges() (ok bool)
		NtUserDeferWindowPosAndBand() (ok bool)
		NtUserDeferredDesktopRotation() (ok bool)
		NtUserDelegateCapturePointers() (ok bool)
		NtUserDelegateInput() (ok bool)
		NtUserDeleteMenu() (ok bool)
		NtUserDeregisterShellHookWindow() (ok bool)
		NtUserDestroyAcceleratorTable() (ok bool)
		NtUserDestroyActivationObject() (ok bool)
		NtUserDestroyCaret() (ok bool)
		NtUserDestroyCursor() (ok bool)
		NtUserDestroyDCompositionHwndTarget() (ok bool)
		NtUserDestroyInputContext() (ok bool)
		NtUserDestroyMenu() (ok bool)
		NtUserDestroyPalmRejectionDelayZone() (ok bool)
		NtUserDestroyWindow() (ok bool)
		NtUserDisableImmersiveOwner() (ok bool)
		NtUserDisableProcessWindowFiltering() (ok bool)
		NtUserDisableProcessWindowsGhosting() (ok bool)
		NtUserDisableThreadIme() (ok bool)
		NtUserDiscardPointerFrameMessages() (ok bool)
		NtUserDispatchMessage() (ok bool)
		NtUserDisplayConfigGetDeviceInfo() (ok bool)
		NtUserDisplayConfigSetDeviceInfo() (ok bool)
		NtUserDoInitMessagePumpHook() (ok bool)
		NtUserDoSoundConnect() (ok bool)
		NtUserDoSoundDisconnect() (ok bool)
		NtUserDoUninitMessagePumpHook() (ok bool)
		NtUserDownlevelTouchpad() (ok bool)
		NtUserDragDetect() (ok bool)
		NtUserDragObject() (ok bool)
		NtUserDrainThreadCoreMessagingCompletions() (ok bool)
		NtUserDrawAnimatedRects() (ok bool)
		NtUserDrawCaption() (ok bool)
		NtUserDrawCaptionTemp() (ok bool)
		NtUserDrawIconEx() (ok bool)
		NtUserDrawMenuBar() (ok bool)
		NtUserDrawMenuBarTemp() (ok bool)
		NtUserDwmGetRemoteSessionOcclusionEvent() (ok bool)
		NtUserDwmGetRemoteSessionOcclusionState() (ok bool)
		NtUserDwmKernelShutdown() (ok bool)
		NtUserDwmKernelStartup() (ok bool)
		NtUserDwmLockScreenUpdates() (ok bool)
		NtUserDwmValidateWindow() (ok bool)
		NtUserEmptyClipboard() (ok bool)
		NtUserEnableChildWindowDpiMessage() (ok bool)
		NtUserEnableIAMAccess() (ok bool)
		NtUserEnableMenuItem() (ok bool)
		NtUserEnableModernAppWindowKeyboardIntercept() (ok bool)
		NtUserEnableMouseInPointer() (ok bool)
		NtUserEnableMouseInPointerForThread() (ok bool)
		NtUserEnableMouseInPointerForWindow() (ok bool)
		NtUserEnableMouseInputForCursorSuppression() (ok bool)
		NtUserEnableNonClientDpiScaling() (ok bool)
		NtUserEnableResizeLayoutSynchronization() (ok bool)
		NtUserEnableScrollBar() (ok bool)
		NtUserEnableSessionForMMCSS() (ok bool)
		NtUserEnableShellWindowManagementBehavior() (ok bool)
		NtUserEnableSoftwareCursorForScreenCapture() (ok bool)
		NtUserEnableTouchPad() (ok bool)
		NtUserEnableWindow() (ok bool)
		NtUserEnableWindowGDIScaledDpiMessage() (ok bool)
		NtUserEnableWindowResizeOptimization() (ok bool)
		NtUserEndDeferWindowPosEx() (ok bool)
		NtUserEndMenu() (ok bool)
		NtUserEndPaint() (ok bool)
		NtUserEnsureDpiDepSysMetCacheForPlateau() (ok bool)
		NtUserEnumClipboardFormats() (ok bool)
		NtUserEnumDisplayDevices() (ok bool)
		NtUserEnumDisplayMonitors() (ok bool)
		NtUserEnumDisplaySettings() (ok bool)
		NtUserEvent() (ok bool)
		NtUserExcludeUpdateRgn() (ok bool)
		NtUserFillWindow() (ok bool)
		NtUserFindExistingCursorIcon() (ok bool)
		NtUserFindWindowEx() (ok bool)
		NtUserFlashWindowEx() (ok bool)
		NtUserForceEnableNumpadTranslation() (ok bool)
		NtUserForceWindowToDpiForTest() (ok bool)
		NtUserFrostCrashedWindow() (ok bool)
		NtUserFunctionalizeDisplayConfig() (ok bool)
		NtUserGetActiveProcessesDpis() (ok bool)
		NtUserGetAltTabInfo() (ok bool)
		NtUserGetAncestor() (ok bool)
		NtUserGetAppImeLevel() (ok bool)
		NtUserGetAsyncKeyState() (ok bool)
		NtUserGetAtomName() (ok bool)
		NtUserGetAutoRotationState() (ok bool)
		NtUserGetCIMSSM() (ok bool)
		NtUserGetCPD() (ok bool)
		NtUserGetCaretBlinkTime() (ok bool)
		NtUserGetCaretPos() (ok bool)
		NtUserGetClassIcoCur() (ok bool)
		NtUserGetClassInfoEx() (ok bool)
		NtUserGetClassName() (ok bool)
		NtUserGetClipCursor() (ok bool)
		NtUserGetClipboardAccessToken() (ok bool)
		NtUserGetClipboardData() (ok bool)
		NtUserGetClipboardFormatName() (ok bool)
		NtUserGetClipboardMetadata() (ok bool)
		NtUserGetClipboardOwner() (ok bool)
		NtUserGetClipboardSequenceNumber() (ok bool)
		NtUserGetClipboardViewer() (ok bool)
		NtUserGetComboBoxInfo() (ok bool)
		NtUserGetControlBrush() (ok bool)
		NtUserGetControlColor() (ok bool)
		NtUserGetCurrentDpiInfoForWindow() (ok bool)
		NtUserGetCurrentInputMessageSource() (ok bool)
		NtUserGetCursor() (ok bool)
		NtUserGetCursorFrameInfo() (ok bool)
		NtUserGetCursorInfo() (ok bool)
		NtUserGetCursorPos() (ok bool)
		NtUserGetDC() (ok bool)
		NtUserGetDCEx() (ok bool)
		NtUserGetDCompositionHwndBitmap() (ok bool)
		NtUserGetDManipHookInitFunction() (ok bool)
		NtUserGetDesktopID() (ok bool)
		NtUserGetDesktopVisualTransform() (ok bool)
		NtUserGetDeviceChangeInfo() (ok bool)
		NtUserGetDisplayAutoRotationPreferences() (ok bool)
		NtUserGetDisplayAutoRotationPreferencesByProcessId() (ok bool)
		NtUserGetDisplayConfigBufferSizes() (ok bool)
		NtUserGetDoubleClickTime() (ok bool)
		NtUserGetDpiForCurrentProcess() (ok bool)
		NtUserGetDpiForMonitor() (ok bool)
		NtUserGetExtendedPointerDeviceProperty() (ok bool)
		NtUserGetForegroundWindow() (ok bool)
		NtUserGetGUIThreadInfo() (ok bool)
		NtUserGetGestureConfig() (ok bool)
		NtUserGetGestureExtArgs() (ok bool)
		NtUserGetGestureInfo() (ok bool)
		NtUserGetGuiResources() (ok bool)
		NtUserGetHDevName() (ok bool)
		NtUserGetHimetricScaleFactorFromPixelLocation() (ok bool)
		NtUserGetIMEShowStatus() (ok bool)
		NtUserGetIconInfo() (ok bool)
		NtUserGetIconSize() (ok bool)
		NtUserGetImeHotKey() (ok bool)
		NtUserGetImeInfoEx() (ok bool)
		NtUserGetInputContainerId() (ok bool)
		NtUserGetInputDesktop() (ok bool)
		NtUserGetInputEvent() (ok bool)
		NtUserGetInputLocaleInfo() (ok bool)
		NtUserGetInteractiveControlDeviceInfo() (ok bool)
		NtUserGetInteractiveControlInfo() (ok bool)
		NtUserGetInteractiveCtrlSupportedWaveforms() (ok bool)
		NtUserGetInternalWindowPos() (ok bool)
		NtUserGetKeyNameText() (ok bool)
		NtUserGetKeyState() (ok bool)
		NtUserGetKeyboardLayout() (ok bool)
		NtUserGetKeyboardLayoutList() (ok bool)
		NtUserGetKeyboardLayoutName() (ok bool)
		NtUserGetKeyboardState() (ok bool)
		NtUserGetKeyboardType() (ok bool)
		NtUserGetLayeredWindowAttributes() (ok bool)
		NtUserGetListBoxInfo() (ok bool)
		NtUserGetMenuBarInfo() (ok bool)
		NtUserGetMenuIndex() (ok bool)
		NtUserGetMenuItemRect() (ok bool)
		NtUserGetMessage() (ok bool)
		NtUserGetMessagePos() (ok bool)
		NtUserGetMinuserIdForBaseWindow() (ok bool)
		NtUserGetModernAppWindow() (ok bool)
		NtUserGetMouseMovePointsEx() (ok bool)
		NtUserGetObjectInformation() (ok bool)
		NtUserGetOemBitmapSize() (ok bool)
		NtUserGetOpenClipboardWindow() (ok bool)
		NtUserGetOwnerTransformedMonitorRect() (ok bool)
		NtUserGetPhysicalDeviceRect() (ok bool)
		NtUserGetPointerCursorId() (ok bool)
		NtUserGetPointerDevice() (ok bool)
		NtUserGetPointerDeviceCursors() (ok bool)
		NtUserGetPointerDeviceInputSpace() (ok bool)
		NtUserGetPointerDeviceOrientation() (ok bool)
		NtUserGetPointerDeviceProperties() (ok bool)
		NtUserGetPointerDeviceRects() (ok bool)
		NtUserGetPointerDevices() (ok bool)
		NtUserGetPointerFrameTimes() (ok bool)
		NtUserGetPointerInfoList() (ok bool)
		NtUserGetPointerInputTransform() (ok bool)
		NtUserGetPointerProprietaryId() (ok bool)
		NtUserGetPointerType() (ok bool)
		NtUserGetPrecisionTouchPadConfiguration() (ok bool)
		NtUserGetPriorityClipboardFormat() (ok bool)
		NtUserGetProcessDefaultLayout() (ok bool)
		NtUserGetProcessDpiAwarenessContext() (ok bool)
		NtUserGetProcessUIContextInformation() (ok bool)
		NtUserGetProcessWindowStation() (ok bool)
		NtUserGetProp() (ok bool)
		NtUserGetQueueIocp() (ok bool)
		NtUserGetQueueStatus() (ok bool)
		NtUserGetQueueStatusReadonly() (ok bool)
		NtUserGetRawInputBuffer() (ok bool)
		NtUserGetRawInputData() (ok bool)
		NtUserGetRawInputDeviceInfo() (ok bool)
		NtUserGetRawInputDeviceList() (ok bool)
		NtUserGetRawPointerDeviceData() (ok bool)
		NtUserGetRegisteredRawInputDevices() (ok bool)
		NtUserGetRequiredCursorSizes() (ok bool)
		NtUserGetResizeDCompositionSynchronizationObject() (ok bool)
		NtUserGetScrollBarInfo() (ok bool)
		NtUserGetSendMessageReceiver() (ok bool)
		NtUserGetSharedWindowData() (ok bool)
		NtUserGetSysMenuOffset() (ok bool)
		NtUserGetSystemContentRects() (ok bool)
		NtUserGetSystemDpiForProcess() (ok bool)
		NtUserGetSystemMenu() (ok bool)
		NtUserGetThreadDesktop() (ok bool)
		NtUserGetThreadState() (ok bool)
		NtUserGetTitleBarInfo() (ok bool)
		NtUserGetTopLevelWindow() (ok bool)
		NtUserGetTouchInputInfo() (ok bool)
		NtUserGetTouchValidationStatus() (ok bool)
		NtUserGetUniformSpaceMapping() (ok bool)
		NtUserGetUnpredictedMessagePos() (ok bool)
		NtUserGetUpdateRect() (ok bool)
		NtUserGetUpdateRgn() (ok bool)
		NtUserGetUpdatedClipboardFormats() (ok bool)
		NtUserGetWOWClass() (ok bool)
		NtUserGetWinStationInfo() (ok bool)
		NtUserGetWindowBand() (ok bool)
		NtUserGetWindowCompositionAttribute() (ok bool)
		NtUserGetWindowCompositionInfo() (ok bool)
		NtUserGetWindowContextHelpId() (ok bool)
		NtUserGetWindowDC() (ok bool)
		NtUserGetWindowDisplayAffinity() (ok bool)
		NtUserGetWindowFeedbackSetting() (ok bool)
		NtUserGetWindowMinimizeRect() (ok bool)
		NtUserGetWindowPlacement() (ok bool)
		NtUserGetWindowProcessHandle() (ok bool)
		NtUserGetWindowRgnEx() (ok bool)
		NtUserGetWindowThreadProcessId() (ok bool)
		NtUserGetWindowTrackInfoAsync() (ok bool)
		NtUserGhostWindowFromHungWindow() (ok bool)
		NtUserHandleDelegatedInput() (ok bool)
		NtUserHandleSystemThreadCreationFailure() (ok bool)
		NtUserHardErrorControl() (ok bool)
		NtUserHideCaret() (ok bool)
		NtUserHideCursorNoCapture() (ok bool)
		NtUserHidePointerContactVisualization() (ok bool)
		NtUserHiliteMenuItem() (ok bool)
		NtUserHungWindowFromGhostWindow() (ok bool)
		NtUserHwndQueryRedirectionInfo() (ok bool)
		NtUserHwndSetRedirectionInfo() (ok bool)
		NtUserImpersonateDdeClientWindow() (ok bool)
		NtUserInheritWindowMonitor() (ok bool)
		NtUserInitAnsiOem() (ok bool)
		NtUserInitRunThread() (ok bool)
		NtUserInitThreadCoreMessagingIocp() (ok bool)
		NtUserInitialize() (ok bool)
		NtUserInitializeClientPfnArrays() (ok bool)
		NtUserInitializeGenericHidInjection() (ok bool)
		NtUserInitializeInputDeviceInjection() (ok bool)
		NtUserInitializePointerDeviceInjection() (ok bool)
		NtUserInitializePointerDeviceInjectionEx() (ok bool)
		NtUserInitializeTouchInjection() (ok bool)
		NtUserInjectDeviceInput() (ok bool)
		NtUserInjectGenericHidInput() (ok bool)
		NtUserInjectGesture() (ok bool)
		NtUserInjectKeyboardInput() (ok bool)
		NtUserInjectMouseInput() (ok bool)
		NtUserInjectPointerInput() (ok bool)
		NtUserInjectTouchInput() (ok bool)
		NtUserInteractiveControlQueryUsage() (ok bool)
		NtUserInternalGetWindowIcon() (ok bool)
		NtUserInternalGetWindowText() (ok bool)
		NtUserInternalToUnicode() (ok bool)
		NtUserInvalidateRect() (ok bool)
		NtUserInvalidateRgn() (ok bool)
		NtUserIsChildWindowDpiMessageEnabled() (ok bool)
		NtUserIsClipboardFormatAvailable() (ok bool)
		NtUserIsMouseInPointerEnabled() (ok bool)
		NtUserIsMouseInputEnabled() (ok bool)
		NtUserIsNonClientDpiScalingEnabled() (ok bool)
		NtUserIsResizeLayoutSynchronizationEnabled() (ok bool)
		NtUserIsTopLevelWindow() (ok bool)
		NtUserIsTouchWindow() (ok bool)
		NtUserIsWindowBroadcastingDpiToChildren() (ok bool)
		NtUserIsWindowGDIScaledDpiMessageEnabled() (ok bool)
		NtUserKillSystemTimer() (ok bool)
		NtUserKillTimer() (ok bool)
		NtUserLW_LoadFonts() (ok bool)
		NtUserLayoutCompleted() (ok bool)
		NtUserLinkDpiCursor() (ok bool)
		NtUserLoadCursorsAndIcons() (ok bool)
		NtUserLoadKeyboardLayoutEx() (ok bool)
		NtUserLoadUserApiHook() (ok bool)
		NtUserLockCursor() (ok bool)
		NtUserLockSetForegroundWindow() (ok bool)
		NtUserLockWindowStation() (ok bool)
		NtUserLockWindowUpdate() (ok bool)
		NtUserLockWorkStation() (ok bool)
		NtUserLogicalToPerMonitorDPIPhysicalPoint() (ok bool)
		NtUserLogicalToPhysicalDpiPointForWindow() (ok bool)
		NtUserLogicalToPhysicalPoint() (ok bool)
		NtUserMNDragLeave() (ok bool)
		NtUserMNDragOver() (ok bool)
		NtUserMagControl() (ok bool)
		NtUserMagGetContextInformation() (ok bool)
		NtUserMagSetContextInformation() (ok bool)
		NtUserMapDesktopObject() (ok bool)
		NtUserMapPointsByVisualIdentifier() (ok bool)
		NtUserMapVirtualKeyEx() (ok bool)
		NtUserMarkWindowForRawMouse() (ok bool)
		NtUserMenuItemFromPoint() (ok bool)
		NtUserMessageBeep() (ok bool)
		NtUserMessageCall() (ok bool)
		NtUserMinInitialize() (ok bool)
		NtUserMinMaximize() (ok bool)
		NtUserModifyUserStartupInfoFlags() (ok bool)
		NtUserModifyWindowTouchCapability() (ok bool)
		NtUserMoveWindow() (ok bool)
		NtUserMsgWaitForMultipleObjectsEx() (ok bool)
		NtUserNavigateFocus() (ok bool)
		NtUserNlsKbdSendIMENotification() (ok bool)
		NtUserNotifyIMEStatus() (ok bool)
		NtUserNotifyOverlayWindow() (ok bool)
		NtUserNotifyProcessCreate() (ok bool)
		NtUserNotifyWinEvent() (ok bool)
		NtUserOpenClipboard() (ok bool)
		NtUserOpenDesktop() (ok bool)
		NtUserOpenInputDesktop() (ok bool)
		NtUserOpenThreadDesktop() (ok bool)
		NtUserOpenWindowStation() (ok bool)
		NtUserPaintDesktop() (ok bool)
		NtUserPaintMenuBar() (ok bool)
		NtUserPaintMonitor() (ok bool)
		NtUserPeekMessage() (ok bool)
		NtUserPerMonitorDPIPhysicalToLogicalPoint() (ok bool)
		NtUserPhysicalToLogicalDpiPointForWindow() (ok bool)
		NtUserPhysicalToLogicalPoint() (ok bool)
		NtUserPlayEventSound() (ok bool)
		NtUserPostKeyboardInputMessage() (ok bool)
		NtUserPostMessage() (ok bool)
		NtUserPostQuitMessage() (ok bool)
		NtUserPostThreadMessage() (ok bool)
		NtUserPrepareForLogoff() (ok bool)
		NtUserPrintWindow() (ok bool)
		NtUserProcessConnect() (ok bool)
		NtUserProcessInkFeedbackCommand() (ok bool)
		NtUserPromoteMouseInPointer() (ok bool)
		NtUserPromotePointer() (ok bool)
		NtUserQueryBSDRWindow() (ok bool)
		NtUserQueryDisplayConfig() (ok bool)
		NtUserQueryInformationThread() (ok bool)
		NtUserQueryInputContext() (ok bool)
		NtUserQuerySendMessage() (ok bool)
		NtUserQueryWindow() (ok bool)
		NtUserRealChildWindowFromPoint() (ok bool)
		NtUserRealInternalGetMessage() (ok bool)
		NtUserRealWaitMessageEx() (ok bool)
		NtUserRealizePalette() (ok bool)
		NtUserReassociateQueueEventCompletionPacket() (ok bool)
		NtUserRedrawFrame() (ok bool)
		NtUserRedrawFrameAndHook() (ok bool)
		NtUserRedrawTitle() (ok bool)
		NtUserRedrawWindow() (ok bool)
		NtUserRegisterBSDRWindow() (ok bool)
		NtUserRegisterClassExWOW() (ok bool)
		NtUserRegisterDManipHook() (ok bool)
		NtUserRegisterEdgy() (ok bool)
		NtUserRegisterErrorReportingDialog() (ok bool)
		NtUserRegisterForCustomDockTargets() (ok bool)
		NtUserRegisterForTooltipDismissNotification() (ok bool)
		NtUserRegisterGhostWindow() (ok bool)
		NtUserRegisterHotKey() (ok bool)
		NtUserRegisterLPK() (ok bool)
		NtUserRegisterLogonProcess() (ok bool)
		NtUserRegisterManipulationThread() (ok bool)
		NtUserRegisterPointerDeviceNotifications() (ok bool)
		NtUserRegisterPointerInputTarget() (ok bool)
		NtUserRegisterRawInputDevices() (ok bool)
		NtUserRegisterServicesProcess() (ok bool)
		NtUserRegisterSessionPort() (ok bool)
		NtUserRegisterShellHookWindow() (ok bool)
		NtUserRegisterShellPTPListener() (ok bool)
		NtUserRegisterSiblingFrostWindow() (ok bool)
		NtUserRegisterSystemThread() (ok bool)
		NtUserRegisterTasklist() (ok bool)
		NtUserRegisterTouchHitTestingWindow() (ok bool)
		NtUserRegisterTouchPadCapable() (ok bool)
		NtUserRegisterUserApiHook() (ok bool)
		NtUserRegisterWindowArrangementCallout() (ok bool)
		NtUserRegisterWindowMessage() (ok bool)
		NtUserReleaseCapture() (ok bool)
		NtUserReleaseDC() (ok bool)
		NtUserReleaseDwmHitTestWaiters() (ok bool)
		NtUserRemoteConnect() (ok bool)
		NtUserRemoteConnectState() (ok bool)
		NtUserRemoteConsoleShadowStop() (ok bool)
		NtUserRemoteDisconnect() (ok bool)
		NtUserRemoteNotify() (ok bool)
		NtUserRemotePassthruDisable() (ok bool)
		NtUserRemotePassthruEnable() (ok bool)
		NtUserRemoteReconnect() (ok bool)
		NtUserRemoteRedrawRectangle() (ok bool)
		NtUserRemoteRedrawScreen() (ok bool)
		NtUserRemoteShadowCleanup() (ok bool)
		NtUserRemoteShadowSetup() (ok bool)
		NtUserRemoteShadowStart() (ok bool)
		NtUserRemoteShadowStop() (ok bool)
		NtUserRemoteStopScreenUpdates() (ok bool)
		NtUserRemoteThinwireStats() (ok bool)
		NtUserRemoveClipboardFormatListener() (ok bool)
		NtUserRemoveInjectionDevice() (ok bool)
		NtUserRemoveMenu() (ok bool)
		NtUserRemoveProp() (ok bool)
		NtUserRemoveQueueCompletion() (ok bool)
		NtUserRemoveVisualIdentifier() (ok bool)
		NtUserReplyMessage() (ok bool)
		NtUserReportInertia() (ok bool)
		NtUserResetDblClk() (ok bool)
		NtUserResolveDesktopForWOW() (ok bool)
		NtUserRestoreWindowDpiChanges() (ok bool)
		NtUserSBGetParms() (ok bool)
		NtUserScaleSystemMetricForDPIWithoutCache() (ok bool)
		NtUserScheduleDispatchNotification() (ok bool)
		NtUserScrollDC() (ok bool)
		NtUserScrollWindowEx() (ok bool)
		NtUserSelectPalette() (ok bool)
		NtUserSendEventMessage() (ok bool)
		NtUserSendInput() (ok bool)
		NtUserSendInteractiveControlHapticsReport() (ok bool)
		NtUserSetActivationFilter() (ok bool)
		NtUserSetActiveProcessForMonitor() (ok bool)
		NtUserSetActiveWindow() (ok bool)
		NtUserSetAdditionalForegroundBoostProcesses() (ok bool)
		NtUserSetAppImeLevel() (ok bool)
		NtUserSetAutoRotation() (ok bool)
		NtUserSetBridgeWindowChild() (ok bool)
		NtUserSetBrokeredForeground() (ok bool)
		NtUserSetCalibrationData() (ok bool)
		NtUserSetCancelRotationDelayHintWindow() (ok bool)
		NtUserSetCapture() (ok bool)
		NtUserSetCaretBlinkTime() (ok bool)
		NtUserSetCaretPos() (ok bool)
		NtUserSetChildWindowNoActivate() (ok bool)
		NtUserSetClassLong() (ok bool)
		NtUserSetClassLongPtr() (ok bool)
		NtUserSetClassWord() (ok bool)
		NtUserSetClipboardData() (ok bool)
		NtUserSetClipboardViewer() (ok bool)
		NtUserSetCoreWindow() (ok bool)
		NtUserSetCoreWindowPartner() (ok bool)
		NtUserSetCursor() (ok bool)
		NtUserSetCursorIconData() (ok bool)
		NtUserSetCursorIconDataEx() (ok bool)
		NtUserSetCursorPos() (ok bool)
		NtUserSetDesktopColorTransform() (ok bool)
		NtUserSetDesktopVisualInputSink() (ok bool)
		NtUserSetDialogControlDpiChangeBehavior() (ok bool)
		NtUserSetDialogPointer() (ok bool)
		NtUserSetDialogSystemMenu() (ok bool)
		NtUserSetDisplayAutoRotationPreferences() (ok bool)
		NtUserSetDisplayConfig() (ok bool)
		NtUserSetDisplayMapping() (ok bool)
		NtUserSetDoubleClickTime() (ok bool)
		NtUserSetDpiForWindow() (ok bool)
		NtUserSetFallbackForeground() (ok bool)
		NtUserSetFeatureReportResponse() (ok bool)
		NtUserSetFocus() (ok bool)
		NtUserSetForegroundRedirectionForActivationObject() (ok bool)
		NtUserSetForegroundWindow() (ok bool)
		NtUserSetForegroundWindowForApplication() (ok bool)
		NtUserSetFullscreenMagnifierOffsetsDWMUpdated() (ok bool)
		NtUserSetGestureConfig() (ok bool)
		NtUserSetImeHotKey() (ok bool)
		NtUserSetImeInfoEx() (ok bool)
		NtUserSetImeOwnerWindow() (ok bool)
		NtUserSetInformationThread() (ok bool)
		NtUserSetInputServiceState() (ok bool)
		NtUserSetInteractiveControlFocus() (ok bool)
		NtUserSetInteractiveCtrlRotationAngle() (ok bool)
		NtUserSetInternalWindowPos() (ok bool)
		NtUserSetKeyboardState() (ok bool)
		NtUserSetLayeredWindowAttributes() (ok bool)
		NtUserSetMagnificationDesktopMagnifierOffsetsDWMUpdated() (ok bool)
		NtUserSetManipulationInputTarget() (ok bool)
		NtUserSetMenu() (ok bool)
		NtUserSetMenuContextHelpId() (ok bool)
		NtUserSetMenuDefaultItem() (ok bool)
		NtUserSetMenuFlagRtoL() (ok bool)
		NtUserSetMessageExtraInfo() (ok bool)
		NtUserSetMirrorRendering() (ok bool)
		NtUserSetModernAppWindow() (ok bool)
		NtUserSetMonitorWorkArea() (ok bool)
		NtUserSetMouseInputRateLimitingTimer() (ok bool)
		NtUserSetMsgBox() (ok bool)
		NtUserSetObjectInformation() (ok bool)
		NtUserSetParent() (ok bool)
		NtUserSetPrecisionTouchPadConfiguration() (ok bool)
		NtUserSetProcessDefaultLayout() (ok bool)
		NtUserSetProcessDpiAwarenessContext() (ok bool)
		NtUserSetProcessInteractionFlags() (ok bool)
		NtUserSetProcessLaunchForegroundPolicy() (ok bool)
		NtUserSetProcessMousewheelRoutingMode() (ok bool)
		NtUserSetProcessRestrictionExemption() (ok bool)
		NtUserSetProcessUIAccessZorder() (ok bool)
		NtUserSetProcessWindowStation() (ok bool)
		NtUserSetProgmanWindow() (ok bool)
		NtUserSetProp() (ok bool)
		NtUserSetScrollInfo() (ok bool)
		NtUserSetSensorPresence() (ok bool)
		NtUserSetSharedWindowData() (ok bool)
		NtUserSetShellChangeNotifyHWND() (ok bool)
		NtUserSetShellWindowEx() (ok bool)
		NtUserSetSysColors() (ok bool)
		NtUserSetSysMenu() (ok bool)
		NtUserSetSystemContentRects() (ok bool)
		NtUserSetSystemCursor() (ok bool)
		NtUserSetSystemMenu() (ok bool)
		NtUserSetSystemTimer() (ok bool)
		NtUserSetTSFEventState() (ok bool)
		NtUserSetTargetForResourceBrokering() (ok bool)
		NtUserSetTaskmanWindow() (ok bool)
		NtUserSetThreadDesktop() (ok bool)
		NtUserSetThreadInputBlocked() (ok bool)
		NtUserSetThreadLayoutHandles() (ok bool)
		NtUserSetThreadQueueMergeSetting() (ok bool)
		NtUserSetThreadState() (ok bool)
		NtUserSetTimer() (ok bool)
		NtUserSetVisible() (ok bool)
		NtUserSetWaitForQueueAttach() (ok bool)
		NtUserSetWatermarkStrings() (ok bool)
		NtUserSetWinEventHook() (ok bool)
		NtUserSetWindowBand() (ok bool)
		NtUserSetWindowCompositionAttribute() (ok bool)
		NtUserSetWindowCompositionTransition() (ok bool)
		NtUserSetWindowContextHelpId() (ok bool)
		NtUserSetWindowDisplayAffinity() (ok bool)
		NtUserSetWindowFNID() (ok bool)
		NtUserSetWindowFeedbackSetting() (ok bool)
		NtUserSetWindowLong() (ok bool)
		NtUserSetWindowLongPtr() (ok bool)
		NtUserSetWindowPlacement() (ok bool)
		NtUserSetWindowPos() (ok bool)
		NtUserSetWindowRgn() (ok bool)
		NtUserSetWindowRgnEx() (ok bool)
		NtUserSetWindowShowState() (ok bool)
		NtUserSetWindowState() (ok bool)
		NtUserSetWindowStationUser() (ok bool)
		NtUserSetWindowWord() (ok bool)
		NtUserSetWindowsHookAW() (ok bool)
		NtUserSetWindowsHookEx() (ok bool)
		NtUserShellMigrateWindow() (ok bool)
		NtUserShellSetWindowPos() (ok bool)
		NtUserShowCaret() (ok bool)
		NtUserShowCursor() (ok bool)
		NtUserShowOwnedPopups() (ok bool)
		NtUserShowScrollBar() (ok bool)
		NtUserShowStartGlass() (ok bool)
		NtUserShowSystemCursor() (ok bool)
		NtUserShowWindow() (ok bool)
		NtUserShowWindowAsync() (ok bool)
		NtUserShutdownBlockReasonCreate() (ok bool)
		NtUserShutdownBlockReasonQuery() (ok bool)
		NtUserShutdownReasonDestroy() (ok bool)
		NtUserSignalRedirectionStartComplete() (ok bool)
		NtUserSlicerControl() (ok bool)
		NtUserSoundSentry() (ok bool)
		NtUserStopAndEndInertia() (ok bool)
		NtUserSwapMouseButton() (ok bool)
		NtUserSwitchDesktop() (ok bool)
		NtUserSwitchToThisWindow() (ok bool)
		NtUserSystemParametersInfo() (ok bool)
		NtUserSystemParametersInfoForDpi() (ok bool)
		NtUserTestForInteractiveUser() (ok bool)
		NtUserThreadMessageQueueAttached() (ok bool)
		NtUserThunkedMenuInfo() (ok bool)
		NtUserThunkedMenuItemInfo() (ok bool)
		NtUserToUnicodeEx() (ok bool)
		NtUserTraceLoggingSendMixedModeTelemetry() (ok bool)
		NtUserTrackMouseEvent() (ok bool)
		NtUserTrackPopupMenuEx() (ok bool)
		NtUserTransformPoint() (ok bool)
		NtUserTransformRect() (ok bool)
		NtUserTranslateAccelerator() (ok bool)
		NtUserTranslateMessage() (ok bool)
		NtUserUndelegateInput() (ok bool)
		NtUserUnhookWinEvent() (ok bool)
		NtUserUnhookWindowsHook() (ok bool)
		NtUserUnhookWindowsHookEx() (ok bool)
		NtUserUnloadKeyboardLayout() (ok bool)
		NtUserUnlockWindowStation() (ok bool)
		NtUserUnregisterClass() (ok bool)
		NtUserUnregisterHotKey() (ok bool)
		NtUserUnregisterSessionPort() (ok bool)
		NtUserUnregisterUserApiHook() (ok bool)
		NtUserUpdateClientRect() (ok bool)
		NtUserUpdateDefaultDesktopThumbnail() (ok bool)
		NtUserUpdateInputContext() (ok bool)
		NtUserUpdateInstance() (ok bool)
		NtUserUpdateLayeredWindow() (ok bool)
		NtUserUpdatePerUserImmEnabling() (ok bool)
		NtUserUpdatePerUserSystemParameters() (ok bool)
		NtUserUpdateWindow() (ok bool)
		NtUserUpdateWindowInputSinkHints() (ok bool)
		NtUserUpdateWindowTrackingInfo() (ok bool)
		NtUserUpdateWindows() (ok bool)
		NtUserUserHandleGrantAccess() (ok bool)
		NtUserUserPowerCalloutWorker() (ok bool)
		NtUserValidateHandleSecure() (ok bool)
		NtUserValidateRect() (ok bool)
		NtUserValidateRgn() (ok bool)
		NtUserValidateTimerCallback() (ok bool)
		NtUserVkKeyScanEx() (ok bool)
		NtUserWaitAvailableMessageEx() (ok bool)
		NtUserWaitForInputIdle() (ok bool)
		NtUserWaitForRedirectionStartComplete() (ok bool)
		NtUserWaitMessage() (ok bool)
		NtUserWakeRITForShutdown() (ok bool)
		NtUserWindowFromDC() (ok bool)
		NtUserWindowFromPhysicalPoint() (ok bool)
		NtUserWindowFromPoint() (ok bool)
		NtUserZapActiveAndFocus() (ok bool)
		NtValidateCompositionSurfaceHandle() (ok bool)
		NtVisualCaptureBits() (ok bool)
	}
	objectWin32u struct{}
)

func (o *objectWin32u) NtBindCompositionSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCloseCompositionInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCompositionInputThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCompositionSetDropTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCompositorNotifyExitWindows() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtConfigureInputSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCreateCompositionInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCreateCompositionSurfaceHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtCreateImplicitCompositionInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionAddCrossDeviceVisualChild() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionBeginFrame() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionBoostCompositorClock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCommitChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCommitSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionConfirmFrame() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionConnectPipe() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateAndBindSharedSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateConnection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateDwmChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateSharedResourceHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionCreateSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionDestroyChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionDestroyConnection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionDuplicateHandleToProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionDuplicateSwapchainHandleToDwm() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionEnableMMCSS() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetBatchId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetChannels() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetConnectionBatch() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetDeletedResources() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetFrameId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetFrameIdFromBatchId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetFrameLegacyTokens() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetFrameStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetFrameSurfaceUpdates() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetMaterialProperty() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionGetTargetStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionNotifySuperWetInkWork() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionProcessChannelBatchBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionReferenceSharedResourceOnDwmChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionRegisterThumbnailVisual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionRegisterVirtualDesktopVisual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionReleaseAllResources() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionRemoveCrossDeviceVisualChild() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetBlurredWallpaperSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetChannelCommitCompletionEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetChannelConnectionId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetChildRootVisual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetDebugCounter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSetMaterialProperty() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSubmitDWMBatch() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSuspendAnimations() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionSynchronize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionTelemetrySetApplicationId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionUpdatePointerCapture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionWaitForChannel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDCompositionWaitForCompositorClock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDesktopCaptureBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDuplicateCompositionInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkCancelPresents() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkCheckSinglePlaneForMultiPlaneOverlaySupport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkConnectDoorbell() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkCreateDoorbell() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkCreateNativeFence() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkCreateTrackedWorkload() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkDestroyDoorbell() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkDestroyTrackedWorkload() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkDispMgrOperation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkDisplayPortOperation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkDuplicateHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkEnumAdapters3() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkEnumProcesses() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkGetAvailableTrackedWorkloadIndex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkGetProcessList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkGetProperties() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkGetTrackedWorkloadStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkNotifyWorkSubmission() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkOpenNativeFenceFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkOutputDuplPresentToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkPinResources() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkRegisterVailProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkResetTrackedWorkloadStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkSetProperties() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkSubmitPresentBltToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkSubmitPresentToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkUnpinResources() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkUpdateTrackedWorkload() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkVailConnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkVailDisconnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtDxgkVailPromoteCompositionSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtEnableOneCoreTransformMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectAddContent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectAddPoolBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerAcquirePresent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerAdjustUsageReference() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerBeginProcessPresent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerEndProcessPresent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerPostMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectConsumerQueryBufferInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectCreate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectDisconnectEndpoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectEnablePresentStatisticsType() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectOpen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectPresentCancel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectQueryBufferAvailableEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectQueryEndpointConnected() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectQueryLostEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectQueryNextMessageToProducer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectReadNextMessageToProducer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectRemoveContent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectRemovePoolBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtFlipObjectSetContent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAbortDoc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAbortPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddEmbFontToDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddFontMemResourceEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddFontResourceW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddInitialFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddRemoteFontToDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAddRemoteMMInstanceToDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAlphaBlend() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAngleArc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiAnyLinkedFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiArcInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBRUSHOBJ_DeleteRbrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBRUSHOBJ_hGetColorTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBRUSHOBJ_pvAllocRbrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBRUSHOBJ_pvGetRbrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBRUSHOBJ_ulGetBrushColor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBeginGdiRendering() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBeginPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiBitBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCLIPOBJ_bEnum() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCLIPOBJ_cEnumStart() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCLIPOBJ_ppoGetPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCancelDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiChangeGhostFont() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCheckBitmapBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiClearBitmapAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiClearBrushAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCloseFigure() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiColorCorrectPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCombineRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCombineTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiComputeXformCoefficients() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiConfigureOPMProtectedOutput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiConvertMetafileRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateBitmapFromDxSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateBitmapFromDxSurface2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateClientObj() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateColorSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateColorTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateCompatibleBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateCompatibleDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateDIBBrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateDIBSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateDIBitmapInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateEllipticRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateHalftonePalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateHatchBrushInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateMetafileDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateOPMProtectedOutput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateOPMProtectedOutputs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreatePaletteInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreatePatternBrushInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreatePen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateRectRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateRoundRectRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateServerMetaFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateSessionMappedDIBSection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiCreateSolidBrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCIGetCapabilitiesString() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCIGetCapabilitiesStringLength() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCIGetTimingReport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCIGetVCPFeature() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCISaveCurrentSettings() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDDCCISetVCPFeature() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdCreateFullscreenSprite() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAbandonSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAcquireKeyedMutex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAcquireKeyedMutex2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAcquireSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAddSurfaceToSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIAdjustFullscreenGamma() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICacheHybridQueryValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIChangeVideoMemoryReservation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckExclusiveOwnership() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckMonitorPowerState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckMultiPlaneOverlaySupport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckMultiPlaneOverlaySupport2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckMultiPlaneOverlaySupport3() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckOcclusion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckSharedResourceAccess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICheckVidPnExclusiveOwnership() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICloseAdapter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIConfigureSharedResource() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateAllocation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateBundleObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateContextVirtual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateDCFromMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateHwContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateKeyedMutex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateKeyedMutex2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateOutputDupl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateOverlay() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreatePagingQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateProtectedSession() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDICreateSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDDisplayEnum() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyAllocation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyAllocation2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyDCFromMemory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyHwContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyKeyedMutex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyOutputDupl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyOverlay() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyPagingQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroyProtectedSession() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDestroySynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDispMgrCreate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDispMgrSourceOperation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIDispMgrTargetOperation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIEnumAdapters() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIEnumAdapters2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIEscape() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIEvict() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIExtractBundleObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIFlipOverlay() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIFlushHeapTransitions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIFreeGpuVirtualAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetAllocationPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetCachedHybridQueryValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetContextInProcessSchedulingPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetContextSchedulingPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetDWMVerticalBlankEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetDeviceState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetDisplayModeList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetMemoryBudgetTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetMultiPlaneOverlayCaps() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetMultisampleMethodList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetOverlayState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetPostCompositionCaps() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetPresentHistory() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetPresentQueueEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetProcessDeviceRemovalSupport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetProcessSchedulingPriorityBand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetProcessSchedulingPriorityClass() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetResourcePresentPrivateDriverData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetRuntimeData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetScanLine() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetSetSwapChainMetadata() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetSharedPrimaryHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetSharedResourceAdapterLuid() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetSharedResourceAdapterLuidFlipManager() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetSwapChainSurfacePhysicalAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIGetYieldPercentage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIInvalidateActiveVidPn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIInvalidateCache() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDILock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDILock2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIMakeResident() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIMapGpuVirtualAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIMarkDeviceAsError() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDINetDispGetNextChunkInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDINetDispQueryMiracastDisplayDeviceStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDINetDispQueryMiracastDisplayDeviceSupport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDINetDispStartMiracastDisplayDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDINetDispStopMiracastDisplayDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOfferAllocations() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenAdapterFromDeviceName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenAdapterFromHdc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenAdapterFromLuid() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenBundleObjectNtHandleFromName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenKeyedMutex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenKeyedMutex2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenKeyedMutexFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenNtHandleFromName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenProtectedSessionFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenResource() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenResourceFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenSyncObjectFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenSyncObjectFromNtHandle2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenSyncObjectNtHandleFromName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOpenSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOutputDuplGetFrameInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOutputDuplGetMetaData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOutputDuplGetPointerShapeData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOutputDuplPresent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIOutputDuplReleaseFrame() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPollDisplayChildren() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPresent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPresentMultiPlaneOverlay() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPresentMultiPlaneOverlay2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPresentMultiPlaneOverlay3() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIPresentRedirected() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryAdapterInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryAllocationResidency() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryClockCalibration() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryFSEBlock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryProcessOfferInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryProtectedSessionInfoFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryProtectedSessionStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryRemoteVidPnSourceFromGdiDisplayName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryResourceInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryResourceInfoFromNtHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryVidPnExclusiveOwnership() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIQueryVideoMemoryInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReclaimAllocations() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReclaimAllocations2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReleaseKeyedMutex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReleaseKeyedMutex2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReleaseProcessVidPnSourceOwners() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReleaseSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIRemoveSurfaceFromSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIRender() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIReserveGpuVirtualAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetAllocationPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetContextInProcessSchedulingPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetContextSchedulingPriority() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetDisplayMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetDodIndirectSwapchain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetFSEBlock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetGammaRamp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetHwProtectionTeardownRecovery() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetMemoryBudgetTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetMonitorColorSpaceTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetProcessDeviceRemovalSupport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetProcessSchedulingPriorityBand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetProcessSchedulingPriorityClass() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetQueuedLimit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetStablePowerState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetStereoEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetSyncRefreshCountWaitTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetVidPnSourceHwProtection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetVidPnSourceOwner() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISetYieldPercentage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIShareObjects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISharedPrimaryLockNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISharedPrimaryUnLockNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISignalSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISignalSynchronizationObjectFromCpu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISignalSynchronizationObjectFromGpu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISignalSynchronizationObjectFromGpu2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISubmitCommand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISubmitCommandToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISubmitSignalSyncObjectsToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDISubmitWaitForSyncObjectsToHwQueue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDITrimProcessCommitment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUnOrderedPresentSwapChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUnlock() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUnlock2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUpdateAllocationProperty() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUpdateGpuVirtualAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIUpdateOverlay() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForIdle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForSynchronizationObjectFromCpu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForSynchronizationObjectFromGpu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForVerticalBlankEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDDIWaitForVerticalBlankEvent2() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdDestroyFullscreenSprite() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdNotifyFullscreenSpriteUpdate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDdQueryVisRgnUniqueness() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDeleteClientObj() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDeleteColorSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDeleteColorTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDeleteObjectApp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDescribePixelFormat() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDestroyOPMProtectedOutput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDestroyPhysicalMonitor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDoBanding() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDoPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDrawEscape() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDrawStream() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiDwmCreatedBitmapRemotingOutput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEllipse() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEnableEudc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEndDoc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEndGdiRendering() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEndPage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEndPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngAlphaBlend() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngAssociateSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngBitBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCheckAbort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngComputeGlyphSet() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCopyBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCreateBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCreateClip() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCreateDeviceBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCreateDeviceSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngCreatePalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngDeleteClip() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngDeletePalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngDeletePath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngDeleteSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngEraseSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngFillPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngGradientFill() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngLineTo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngLockSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngMarkBandingSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngPaint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngPlgBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngStretchBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngStretchBltROP() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngStrokeAndFillPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngStrokePath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngTextOut() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngTransparentBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEngUnlockSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEnsureDpiDepDefaultGuiFontForPlateau() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEnumFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEnumObjects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEqualRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiEudcLoadUnloadLink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExcludeClipRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtCreatePen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtCreateRegion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtEscape() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtFloodFill() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtGetObjectW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtSelectClipRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiExtTextOutW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_cGetAllGlyphHandles() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_cGetGlyphs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_pQueryGlyphAttrs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_pfdg() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_pifi() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_pvTrueTypeFontFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_pxoGetXform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFONTOBJ_vGetInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFillPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFillRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFlattenPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFlush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFontIsLinked() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiForceUFIMapping() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFrameRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiFullscreenControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetAndSetDCDword() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetAppClipBox() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetBitmapBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetBitmapDimension() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetBitmapDpiScaleValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetBoundsRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCOPPCompatibleOPMInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCertificate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCertificateByHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCertificateSize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCertificateSizeByHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCharABCWidthsW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCharSet() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCharWidthInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCharWidthW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCharacterPlacementW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetColorAdjustment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetColorSpaceforBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetCurrentDpiInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDCDpiScaleValue() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDCDword() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDCObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDCPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDCforBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDIBitsInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDeviceCaps() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDeviceCapsAll() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDeviceWidth() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetDhpdev() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetETM() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetEmbUFI() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetEmbedFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetEntry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetEudcTimeStampEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetFontData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetFontFileData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetFontFileInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetFontResourceInfoInternalW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetFontUnicodeRanges() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetGlyphIndicesW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetGlyphIndicesWInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetGlyphOutline() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetKerningPairs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetLinkedUFIs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetMiterLimit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetMonitorID() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetNearestColor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetNearestPaletteIndex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetNumberOfPhysicalMonitors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetOPMInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetOPMRandomNumber() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetObjectBitmapHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetOutlineTextMetricsInternalW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPerBandInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPhysicalMonitorDescription() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPhysicalMonitors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPixel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetProcessSessionFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetPublicFontTableChangeCookie() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetRandomRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetRasterizerCaps() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetRealizationInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetRegionData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetRgnBox() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetServerMetaFileBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetSpoolMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetStats() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetStringBitmapW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetSuggestedOPMProtectedOutputArraySize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetSystemPaletteUse() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTextCharsetInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTextExtent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTextExtentExW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTextFaceW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTextMetricsW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetUFI() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetUFIPathname() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGetWidthTable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiGradientFill() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiHLSurfGetInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiHLSurfSetInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiHT_Get8BPPFormatPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiHT_Get8BPPMaskPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiHfontCreate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiIcmBrushInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiInit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiInitSpool() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiIntersectClipRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiInvertRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiLineTo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMakeFontDir() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMakeInfoDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMakeObjectUnXferable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMakeObjectXferable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMaskBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMirrorWindowOrg() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiModifyWorldTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMonoBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiMoveTo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiOffsetClipRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiOffsetRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiOpenDCW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPATHOBJ_bEnum() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPATHOBJ_bEnumClipLines() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPATHOBJ_vEnumStart() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPATHOBJ_vEnumStartClipLines() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPATHOBJ_vGetBounds() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPatBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPathToRegion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPlgBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPolyDraw() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPolyPatBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPolyPolyDraw() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPolyTextOutW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPtInRegion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiPtVisible() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiQueryFontAssocInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiQueryFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRectInRegion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRectVisible() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRectangle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRemoveFontMemResourceEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRemoveFontResourceW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRemoveMergeFont() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiResetDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiResizePalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRestoreDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiRoundRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSTROBJ_bEnum() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSTROBJ_bEnumPositionsOnly() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSTROBJ_bGetAdvanceWidths() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSTROBJ_dwGetCodePage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSTROBJ_vEnumStart() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSaveDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiScaleRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiScaleValues() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiScaleViewportExtEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiScaleWindowExtEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSelectBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSelectBrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSelectClipPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSelectFont() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSelectPen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBitmapAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBitmapBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBitmapDimension() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBoundsRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBrushAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetBrushOrg() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetColorAdjustment() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetColorSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetDIBitsToDeviceInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetFontEnumeration() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetFontXform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetIcmMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetLinkedUFIs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetMagicColors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetMetaRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetMiterLimit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetOPMSigningKeyAndSequenceNumbers() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetPUMPDOBJ() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetPixel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetPixelFormat() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetRectRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetSizeDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetSystemPaletteUse() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetTextJustification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetUMPDSandboxState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSetVirtualResolution() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStartDoc() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStartPage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStretchBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStretchDIBitsInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStrokeAndFillPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiStrokePath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiSwapBuffers() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiTransformPoints() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiTransparentBlt() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUMPDEngFreeUserMem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUnloadPrinterDriver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUnmapMemFont() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUnrealizeObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUpdateColors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiUpdateTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiWaitForTextReady() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiWidenPath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiXFORMOBJ_bApplyXform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiXFORMOBJ_iGetXform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiXLATEOBJ_cGetPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiXLATEOBJ_hGetColorTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtGdiXLATEOBJ_iXlate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtHWCursorUpdatePointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtInputSpaceRegionFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtIsOneCoreTransformMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtKSTInitialize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtKSTWait() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITAccessibilityTimerNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITActivateInputProcessing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITConfigureVirtualTouchpad() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITCoreMsgKOpenConnectionTo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITDeactivateInputProcessing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITDisableMouseIntercept() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITDispatchCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITEnableMouseIntercept() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITGetCursorUpdateHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITInitMinuserThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITMinuserSetInputTransformOffset() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITMinuserWindowDestroyed() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITPostMouseInputMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITPostThreadEventMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITPostWindowEventMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITPrepareReceiveInputMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITPrepareSendInputMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITProcessDelegateCapturedPointers() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetInputCallbacks() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetInputDelegationMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetInputObservationState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetKeyboardInputRoutingPolicy() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetKeyboardOverriderState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSetLastInputRecipient() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSynthesizeKeyboardInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSynthesizeMouseInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITSynthesizeTouchInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITUninitMinuserThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMITUpdateInputGlobals() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMapVisualRelativePoints() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMinGetInputTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMinInteropCoreMessagingWithInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMinQPeekForInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMinQSuspendInputProcessing() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtMinQUpdateWakeMask() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreBeginLayoutUpdate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreCreateDCompositionHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreCreateGDIHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreDestroyDCompositionHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreDestroyGDIHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreEnableResizeLayoutSynchronization() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreGetNavigationWindowVisual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreGetResizeDCompositionSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreGetWindowContentVisual() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreIdleTimerThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreIsResizeLayoutSynchronizationEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreProcessConnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreRegisterEnhancedNavigationWindowHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreRegisterNavigationWindowHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreSetNavigationServiceSid() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtModerncoreUnregisterNavigationWindowHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtNotifyPresentToCompositionSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtOpenCompositionSurfaceDirtyRegion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtOpenCompositionSurfaceRealizationInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtOpenCompositionSurfaceSectionInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionInputIsImplicit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionInputQueueAndTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionInputSinkLuid() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionInputSinkViewId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionSurfaceBinding() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionSurfaceFrameRate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionSurfaceHDRMetaData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionSurfaceRenderingRealization() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtQueryCompositionSurfaceStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMAddInputObserver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMAreSiblingDevices() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMDeviceIoControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMEnableMonitorMappingForDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMFreeInputBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetDevicePreparsedData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetDevicePreparsedDataLockfree() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetDeviceProperties() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetDevicePropertiesLockfree() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetPhysicalDeviceRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMGetSourceProcessId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMObserveNextInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMOnAsyncPnpWorkNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMOnPnpNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMOnTimerNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMQueryDevicePath() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMReadInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMRegisterForInputEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMRemoveInputObserver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMSetDeadzoneRotation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMSetExtendedDeviceProperty() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMSetTestModeStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMUnregisterForInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtRIMUpdateInputObserverRegistration() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCompositionSurfaceAnalogExclusive() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCompositionSurfaceBufferUsage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCompositionSurfaceDirectFlipState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCompositionSurfaceIndependentFlipInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCompositionSurfaceStatistics() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetCursorInputSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetPointerDeviceInputSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtSetShellCursorState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerConfirmOutstandingAnalogToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerCreateCompositionTokenHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerCreateFlipObjectReturnTokenHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerCreateFlipObjectTokenHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerGetAnalogExclusiveSurfaceUpdates() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerGetAnalogExclusiveTokenEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerOpenSectionAndEvents() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtTokenManagerThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUnBindCompositionSurface() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUpdateInputSinkTransforms() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAcquireIAMKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAcquireInteractiveControlBackgroundAccess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserActivateKeyboardLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAddClipboardFormatListener() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAddVisualIdentifier() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAllowSetForegroundWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAlterWindowStyle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserArrangeIconicWindows() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAssociateInputContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAttachThreadInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAutoPromoteMouseInPointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserAutoRotateScreen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBeginDeferWindowPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBeginLayoutUpdate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBeginPaint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBitBltSysBmp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBlockInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBroadcastImeShowStatusChange() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBroadcastThemeChangeEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBuildHimcList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBuildHwndList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBuildNameList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserBuildPropList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCalcMenuBar() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCalculatePopupWindowPosition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCallMsgFilter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCallNextHookEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCanBrokerForceForeground() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCancelQueueEventCompletionPacket() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserChangeClipboardChain() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserChangeDisplaySettings() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserChangeWindowMessageFilter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserChangeWindowMessageFilterEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckAccessForIntegrityLevel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckImeShowStatusInThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckMenuItem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckProcessForClipboardAccess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckProcessSession() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCheckWindowThreadDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserChildWindowFromPointEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCitSetInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserClearForeground() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserClearRunWakeBit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserClearWakeMask() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserClearWindowState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserClipCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCloseClipboard() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCloseDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCloseWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCompositionInputSinkLuidFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCompositionInputSinkViewInstanceIdFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserConfigureActivationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserConfirmResizeCommit() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserConsoleControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserConvertMemHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCopyAcceleratorTable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCountClipboardFormats() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateAcceleratorTable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateActivationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateBaseWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateCaret() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateDCompositionHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateDesktopEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateEmptyCursorObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateInputContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateLocalMemHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreatePalmRejectionDelayZone() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreatePopupMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateSystemThreads() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateWindowEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCreateWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCsDdeUninitialize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserCtxDisplayIOCtl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDWP_GetEnabledPopupOffset() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDdeInitialize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDefSetText() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDeferWindowDpiChanges() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDeferWindowPosAndBand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDeferredDesktopRotation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDelegateCapturePointers() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDelegateInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDeleteMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDeregisterShellHookWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyAcceleratorTable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyActivationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyCaret() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyDCompositionHwndTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyInputContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyPalmRejectionDelayZone() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDestroyWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisableImmersiveOwner() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisableProcessWindowFiltering() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisableProcessWindowsGhosting() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisableThreadIme() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDiscardPointerFrameMessages() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDispatchMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisplayConfigGetDeviceInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDisplayConfigSetDeviceInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDoInitMessagePumpHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDoSoundConnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDoSoundDisconnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDoUninitMessagePumpHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDownlevelTouchpad() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDragDetect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDragObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrainThreadCoreMessagingCompletions() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawAnimatedRects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawCaption() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawCaptionTemp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawIconEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawMenuBar() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDrawMenuBarTemp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmGetRemoteSessionOcclusionEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmGetRemoteSessionOcclusionState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmKernelShutdown() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmKernelStartup() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmLockScreenUpdates() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserDwmValidateWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEmptyClipboard() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableChildWindowDpiMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableIAMAccess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableMenuItem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableModernAppWindowKeyboardIntercept() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableMouseInPointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableMouseInPointerForThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableMouseInPointerForWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableMouseInputForCursorSuppression() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableNonClientDpiScaling() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableResizeLayoutSynchronization() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableScrollBar() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableSessionForMMCSS() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableShellWindowManagementBehavior() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableSoftwareCursorForScreenCapture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableTouchPad() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableWindowGDIScaledDpiMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnableWindowResizeOptimization() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEndDeferWindowPosEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEndMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEndPaint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnsureDpiDepSysMetCacheForPlateau() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnumClipboardFormats() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnumDisplayDevices() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnumDisplayMonitors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEnumDisplaySettings() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserExcludeUpdateRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFillWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFindExistingCursorIcon() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFindWindowEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFlashWindowEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserForceEnableNumpadTranslation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserForceWindowToDpiForTest() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFrostCrashedWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserFunctionalizeDisplayConfig() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetActiveProcessesDpis() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAltTabInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAncestor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAppImeLevel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAsyncKeyState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAtomName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetAutoRotationState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCIMSSM() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCPD() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCaretBlinkTime() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCaretPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClassIcoCur() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClassInfoEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClassName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardAccessToken() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardFormatName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardMetadata() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardOwner() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardSequenceNumber() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetClipboardViewer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetComboBoxInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetControlBrush() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetControlColor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCurrentDpiInfoForWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCurrentInputMessageSource() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCursorFrameInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCursorInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetCursorPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDCEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDCompositionHwndBitmap() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDManipHookInitFunction() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDesktopID() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDesktopVisualTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDeviceChangeInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDisplayAutoRotationPreferences() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDisplayAutoRotationPreferencesByProcessId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDisplayConfigBufferSizes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDoubleClickTime() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDpiForCurrentProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetDpiForMonitor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetExtendedPointerDeviceProperty() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetForegroundWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetGUIThreadInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetGestureConfig() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetGestureExtArgs() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetGestureInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetGuiResources() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetHDevName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetHimetricScaleFactorFromPixelLocation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetIMEShowStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetIconInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetIconSize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetImeHotKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetImeInfoEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInputContainerId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInputDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInputEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInputLocaleInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInteractiveControlDeviceInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInteractiveControlInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInteractiveCtrlSupportedWaveforms() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetInternalWindowPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyNameText() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyboardLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyboardLayoutList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyboardLayoutName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyboardState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetKeyboardType() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetLayeredWindowAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetListBoxInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMenuBarInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMenuIndex() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMenuItemRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMessagePos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMinuserIdForBaseWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetModernAppWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetMouseMovePointsEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetObjectInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetOemBitmapSize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetOpenClipboardWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetOwnerTransformedMonitorRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPhysicalDeviceRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerCursorId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDeviceCursors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDeviceInputSpace() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDeviceOrientation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDeviceProperties() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDeviceRects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerDevices() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerFrameTimes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerInfoList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerInputTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerProprietaryId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPointerType() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPrecisionTouchPadConfiguration() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetPriorityClipboardFormat() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetProcessDefaultLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetProcessDpiAwarenessContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetProcessUIContextInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetProcessWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetProp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetQueueIocp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetQueueStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetQueueStatusReadonly() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRawInputBuffer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRawInputData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRawInputDeviceInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRawInputDeviceList() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRawPointerDeviceData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRegisteredRawInputDevices() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetRequiredCursorSizes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetResizeDCompositionSynchronizationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetScrollBarInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSendMessageReceiver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSharedWindowData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSysMenuOffset() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSystemContentRects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSystemDpiForProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetSystemMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetThreadDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetThreadState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetTitleBarInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetTopLevelWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetTouchInputInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetTouchValidationStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetUniformSpaceMapping() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetUnpredictedMessagePos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetUpdateRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetUpdateRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetUpdatedClipboardFormats() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWOWClass() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWinStationInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowBand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowCompositionAttribute() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowCompositionInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowContextHelpId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowDisplayAffinity() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowFeedbackSetting() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowMinimizeRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowPlacement() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowProcessHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowRgnEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowThreadProcessId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGetWindowTrackInfoAsync() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserGhostWindowFromHungWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHandleDelegatedInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHandleSystemThreadCreationFailure() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHardErrorControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHideCaret() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHideCursorNoCapture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHidePointerContactVisualization() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHiliteMenuItem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHungWindowFromGhostWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHwndQueryRedirectionInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserHwndSetRedirectionInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserImpersonateDdeClientWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInheritWindowMonitor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitAnsiOem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitRunThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitThreadCoreMessagingIocp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitialize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializeClientPfnArrays() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializeGenericHidInjection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializeInputDeviceInjection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializePointerDeviceInjection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializePointerDeviceInjectionEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInitializeTouchInjection() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectDeviceInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectGenericHidInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectGesture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectKeyboardInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectMouseInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectPointerInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInjectTouchInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInteractiveControlQueryUsage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInternalGetWindowIcon() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInternalGetWindowText() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInternalToUnicode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInvalidateRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserInvalidateRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsChildWindowDpiMessageEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsClipboardFormatAvailable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsMouseInPointerEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsMouseInputEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsNonClientDpiScalingEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsResizeLayoutSynchronizationEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsTopLevelWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsTouchWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsWindowBroadcastingDpiToChildren() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserIsWindowGDIScaledDpiMessageEnabled() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserKillSystemTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserKillTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLW_LoadFonts() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLayoutCompleted() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLinkDpiCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLoadCursorsAndIcons() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLoadKeyboardLayoutEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLoadUserApiHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLockCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLockSetForegroundWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLockWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLockWindowUpdate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLockWorkStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLogicalToPerMonitorDPIPhysicalPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLogicalToPhysicalDpiPointForWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserLogicalToPhysicalPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMNDragLeave() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMNDragOver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMagControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMagGetContextInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMagSetContextInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMapDesktopObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMapPointsByVisualIdentifier() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMapVirtualKeyEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMarkWindowForRawMouse() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMenuItemFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMessageBeep() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMessageCall() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMinInitialize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMinMaximize() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserModifyUserStartupInfoFlags() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserModifyWindowTouchCapability() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMoveWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserMsgWaitForMultipleObjectsEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNavigateFocus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNlsKbdSendIMENotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNotifyIMEStatus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNotifyOverlayWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNotifyProcessCreate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserNotifyWinEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserOpenClipboard() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserOpenDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserOpenInputDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserOpenThreadDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserOpenWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPaintDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPaintMenuBar() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPaintMonitor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPeekMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPerMonitorDPIPhysicalToLogicalPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPhysicalToLogicalDpiPointForWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPhysicalToLogicalPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPlayEventSound() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPostKeyboardInputMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPostMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPostQuitMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPostThreadMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPrepareForLogoff() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPrintWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserProcessConnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserProcessInkFeedbackCommand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPromoteMouseInPointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserPromotePointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQueryBSDRWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQueryDisplayConfig() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQueryInformationThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQueryInputContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQuerySendMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserQueryWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRealChildWindowFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRealInternalGetMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRealWaitMessageEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRealizePalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReassociateQueueEventCompletionPacket() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRedrawFrame() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRedrawFrameAndHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRedrawTitle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRedrawWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterBSDRWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterClassExWOW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterDManipHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterEdgy() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterErrorReportingDialog() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterForCustomDockTargets() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterForTooltipDismissNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterGhostWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterHotKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterLPK() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterLogonProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterManipulationThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterPointerDeviceNotifications() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterPointerInputTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterRawInputDevices() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterServicesProcess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterSessionPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterShellHookWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterShellPTPListener() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterSiblingFrostWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterSystemThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterTasklist() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterTouchHitTestingWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterTouchPadCapable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterUserApiHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterWindowArrangementCallout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRegisterWindowMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReleaseCapture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReleaseDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReleaseDwmHitTestWaiters() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteConnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteConnectState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteConsoleShadowStop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteDisconnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteNotify() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemotePassthruDisable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemotePassthruEnable() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteReconnect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteRedrawRectangle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteRedrawScreen() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteShadowCleanup() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteShadowSetup() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteShadowStart() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteShadowStop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteStopScreenUpdates() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoteThinwireStats() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveClipboardFormatListener() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveInjectionDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveProp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveQueueCompletion() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRemoveVisualIdentifier() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReplyMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserReportInertia() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserResetDblClk() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserResolveDesktopForWOW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserRestoreWindowDpiChanges() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSBGetParms() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserScaleSystemMetricForDPIWithoutCache() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserScheduleDispatchNotification() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserScrollDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserScrollWindowEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSelectPalette() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSendEventMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSendInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSendInteractiveControlHapticsReport() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetActivationFilter() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetActiveProcessForMonitor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetActiveWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetAdditionalForegroundBoostProcesses() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetAppImeLevel() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetAutoRotation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetBridgeWindowChild() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetBrokeredForeground() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCalibrationData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCancelRotationDelayHintWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCapture() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCaretBlinkTime() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCaretPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetChildWindowNoActivate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetClassLong() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetClassLongPtr() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetClassWord() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetClipboardData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetClipboardViewer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCoreWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCoreWindowPartner() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCursorIconData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCursorIconDataEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetCursorPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDesktopColorTransform() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDesktopVisualInputSink() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDialogControlDpiChangeBehavior() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDialogPointer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDialogSystemMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDisplayAutoRotationPreferences() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDisplayConfig() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDisplayMapping() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDoubleClickTime() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetDpiForWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetFallbackForeground() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetFeatureReportResponse() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetFocus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetForegroundRedirectionForActivationObject() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetForegroundWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetForegroundWindowForApplication() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetFullscreenMagnifierOffsetsDWMUpdated() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetGestureConfig() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetImeHotKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetImeInfoEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetImeOwnerWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetInformationThread() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetInputServiceState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetInteractiveControlFocus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetInteractiveCtrlRotationAngle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetInternalWindowPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetKeyboardState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetLayeredWindowAttributes() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMagnificationDesktopMagnifierOffsetsDWMUpdated() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetManipulationInputTarget() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMenuContextHelpId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMenuDefaultItem() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMenuFlagRtoL() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMessageExtraInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMirrorRendering() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetModernAppWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMonitorWorkArea() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMouseInputRateLimitingTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetMsgBox() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetObjectInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetParent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetPrecisionTouchPadConfiguration() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessDefaultLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessDpiAwarenessContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessInteractionFlags() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessLaunchForegroundPolicy() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessMousewheelRoutingMode() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessRestrictionExemption() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessUIAccessZorder() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProcessWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProgmanWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetProp() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetScrollInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSensorPresence() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSharedWindowData() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetShellChangeNotifyHWND() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetShellWindowEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSysColors() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSysMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSystemContentRects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSystemCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSystemMenu() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetSystemTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetTSFEventState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetTargetForResourceBrokering() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetTaskmanWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetThreadDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetThreadInputBlocked() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetThreadLayoutHandles() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetThreadQueueMergeSetting() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetThreadState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetTimer() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetVisible() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWaitForQueueAttach() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWatermarkStrings() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWinEventHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowBand() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowCompositionAttribute() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowCompositionTransition() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowContextHelpId() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowDisplayAffinity() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowFNID() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowFeedbackSetting() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowLong() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowLongPtr() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowPlacement() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowRgnEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowShowState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowState() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowStationUser() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowWord() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowsHookAW() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSetWindowsHookEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShellMigrateWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShellSetWindowPos() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowCaret() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowOwnedPopups() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowScrollBar() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowStartGlass() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowSystemCursor() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShowWindowAsync() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShutdownBlockReasonCreate() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShutdownBlockReasonQuery() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserShutdownReasonDestroy() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSignalRedirectionStartComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSlicerControl() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSoundSentry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserStopAndEndInertia() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSwapMouseButton() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSwitchDesktop() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSwitchToThisWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSystemParametersInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserSystemParametersInfoForDpi() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTestForInteractiveUser() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserThreadMessageQueueAttached() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserThunkedMenuInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserThunkedMenuItemInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserToUnicodeEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTraceLoggingSendMixedModeTelemetry() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTrackMouseEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTrackPopupMenuEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTransformPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTransformRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTranslateAccelerator() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserTranslateMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUndelegateInput() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnhookWinEvent() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnhookWindowsHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnhookWindowsHookEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnloadKeyboardLayout() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnlockWindowStation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnregisterClass() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnregisterHotKey() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnregisterSessionPort() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUnregisterUserApiHook() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateClientRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateDefaultDesktopThumbnail() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateInputContext() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateInstance() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateLayeredWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdatePerUserImmEnabling() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdatePerUserSystemParameters() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateWindow() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateWindowInputSinkHints() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateWindowTrackingInfo() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUpdateWindows() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUserHandleGrantAccess() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserUserPowerCalloutWorker() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserValidateHandleSecure() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserValidateRect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserValidateRgn() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserValidateTimerCallback() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserVkKeyScanEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWaitAvailableMessageEx() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWaitForInputIdle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWaitForRedirectionStartComplete() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWaitMessage() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWakeRITForShutdown() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWindowFromDC() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWindowFromPhysicalPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserWindowFromPoint() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtUserZapActiveAndFocus() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtValidateCompositionSurfaceHandle() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *objectWin32u) NtVisualCaptureBits() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func NewWin32u() InterfaceWin32u { return &objectWin32u{} }
