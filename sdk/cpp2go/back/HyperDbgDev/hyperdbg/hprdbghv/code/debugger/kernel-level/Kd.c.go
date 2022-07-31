package kernel-level


type (
Kd interface{
KdInitializeKernelDebugger()(ok bool)//col:62
KdUninitializeKernelDebugger()(ok bool)//col:102
KdDummyDPC()(ok bool)//col:122
KdFireDpc()(ok bool)//col:142
KdComputeDataChecksum()(ok bool)//col:165
KdResponsePacketToDebugger()(ok bool)//col:246
KdLoggingResponsePacketToDebugger()(ok bool)//col:302
KdHandleDebugEventsWhenKernelDebuggerIsAttached()(ok bool)//col:420
KdApplyTasksPreHaltCore()(ok bool)//col:474
KdApplyTasksPostContinueCore()(ok bool)//col:529
KdContinueDebuggee()(ok bool)//col:589
KdContinueDebuggeeJustCurrentCore()(ok bool)//col:613
KdReadRegisters()(ok bool)//col:663
KdReadMemory()(ok bool)//col:713
KdSwitchCore()(ok bool)//col:789
KdCloseConnectionAndUnloadDebuggee()(ok bool)//col:807
KdReloadSymbolDetailsInDebuggee()(ok bool)//col:827
KdNotifyDebuggeeForUserInput()(ok bool)//col:849
KdSendFormatsFunctionResult()(ok bool)//col:874
KdSendCommandFinishedSignal()(ok bool)//col:893
KdHandleHaltsWhenNmiReceivedFromVmxRoot()(ok bool)//col:947
KdCustomDebuggerBreakSpinlockLock()(ok bool)//col:1020
KdHandleBreakpointAndDebugBreakpoints()(ok bool)//col:1124
KdHandleNmi()(ok bool)//col:1172
KdGuaranteedStepInstruction()(ok bool)//col:1224
KdCheckGuestOperatingModeChanges()(ok bool)//col:1290
KdRegularStepInInstruction()(ok bool)//col:1355
KdRegularStepOver()(ok bool)//col:1410
KdPerformRegisterEvent()(ok bool)//col:1426
KdPerformAddActionToEvent()(ok bool)//col:1442
KdQuerySystemState()(ok bool)//col:1491
KdPerformEventQueryAndModification()(ok bool)//col:1618
KdDispatchAndPerformCommandsFromDebugger()(ok bool)//col:2391
KdIsGuestOnUsermode32Bit()(ok bool)//col:2444
KdManageSystemHaltOnVmxRoot()(ok bool)//col:2625
KdBroadcastHaltOnAllCores()(ok bool)//col:2639
KdHaltSystem()(ok bool)//col:2670
}

)

func NewKd() { return & kd{} }

func (k *kd)KdInitializeKernelDebugger()(ok bool){//col:62









return true
}


func (k *kd)KdUninitializeKernelDebugger()(ok bool){//col:102












return true
}


func (k *kd)KdDummyDPC()(ok bool){//col:122







return true
}


func (k *kd)KdFireDpc()(ok bool){//col:142







return true
}


func (k *kd)KdComputeDataChecksum()(ok bool){//col:165












return true
}


func (k *kd)KdResponsePacketToDebugger()(ok bool){//col:246







































return true
}


func (k *kd)KdLoggingResponsePacketToDebugger()(ok bool){//col:302
























return true
}


func (k *kd)KdHandleDebugEventsWhenKernelDebuggerIsAttached()(ok bool){//col:420






























































return true
}


func (k *kd)KdApplyTasksPreHaltCore()(ok bool){//col:474





















return true
}


func (k *kd)KdApplyTasksPostContinueCore()(ok bool){//col:529

























return true
}


func (k *kd)KdContinueDebuggee()(ok bool){//col:589


























return true
}


func (k *kd)KdContinueDebuggeeJustCurrentCore()(ok bool){//col:613







return true
}


func (k *kd)KdReadRegisters()(ok bool){//col:663


























return true
}


func (k *kd)KdReadMemory()(ok bool){//col:713


























return true
}


func (k *kd)KdSwitchCore()(ok bool){//col:789





















return true
}


func (k *kd)KdCloseConnectionAndUnloadDebuggee()(ok bool){//col:807







return true
}


func (k *kd)KdReloadSymbolDetailsInDebuggee()(ok bool){//col:827







return true
}


func (k *kd)KdNotifyDebuggeeForUserInput()(ok bool){//col:849







return true
}


func (k *kd)KdSendFormatsFunctionResult()(ok bool){//col:874











return true
}


func (k *kd)KdSendCommandFinishedSignal()(ok bool){//col:893








return true
}


func (k *kd)KdHandleHaltsWhenNmiReceivedFromVmxRoot()(ok bool){//col:947






return true
}


func (k *kd)KdCustomDebuggerBreakSpinlockLock()(ok bool){//col:1020

































return true
}


func (k *kd)KdHandleBreakpointAndDebugBreakpoints()(ok bool){//col:1124








































return true
}


func (k *kd)KdHandleNmi()(ok bool){//col:1172














return true
}


func (k *kd)KdGuaranteedStepInstruction()(ok bool){//col:1224














return true
}


func (k *kd)KdCheckGuestOperatingModeChanges()(ok bool){//col:1290






























return true
}


func (k *kd)KdRegularStepInInstruction()(ok bool){//col:1355


























return true
}


func (k *kd)KdRegularStepOver()(ok bool){//col:1410
























return true
}


func (k *kd)KdPerformRegisterEvent()(ok bool){//col:1426







return true
}


func (k *kd)KdPerformAddActionToEvent()(ok bool){//col:1442







return true
}


func (k *kd)KdQuerySystemState()(ok bool){//col:1491





























return true
}


func (k *kd)KdPerformEventQueryAndModification()(ok bool){//col:1618



































































return true
}


func (k *kd)KdDispatchAndPerformCommandsFromDebugger()(ok bool){//col:2391

































































































































































































































































































































































































return true
}


func (k *kd)KdIsGuestOnUsermode32Bit()(ok bool){//col:2444






















return true
}


func (k *kd)KdManageSystemHaltOnVmxRoot()(ok bool){//col:2625






































































return true
}


func (k *kd)KdBroadcastHaltOnAllCores()(ok bool){//col:2639




return true
}


func (k *kd)KdHaltSystem()(ok bool){//col:2670





return true
}




