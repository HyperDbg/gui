package kernel-level
type (
	Kd interface {
		KdInitializeKernelDebugger() (ok bool)                      //col:62
		KdUninitializeKernelDebugger() (ok bool)                    //col:101
		KdDummyDPC() (ok bool)                                      //col:120
		KdFireDpc() (ok bool)                                       //col:139
		KdComputeDataChecksum() (ok bool)                           //col:161
		KdResponsePacketToDebugger() (ok bool)                      //col:241
		KdLoggingResponsePacketToDebugger() (ok bool)               //col:296
		KdHandleDebugEventsWhenKernelDebuggerIsAttached() (ok bool) //col:413
		KdApplyTasksPreHaltCore() (ok bool)                         //col:466
		KdApplyTasksPostContinueCore() (ok bool)                    //col:520
		*are continued ()
(ok bool) //col:579
KdContinueDebuggeeJustCurrentCore()(ok bool) //col:602
KdReadRegisters()(ok bool)                   //col:651
KdReadMemory()(ok bool) //col:700
KdSwitchCore()(ok bool) //col:775
KdCloseConnectionAndUnloadDebuggee()(ok bool) //col:792
KdReloadSymbolDetailsInDebuggee()(ok bool) //col:811
KdNotifyDebuggeeForUserInput()(ok bool)    //col:832
KdSendFormatsFunctionResult()(ok bool) //col:856
KdSendCommandFinishedSignal()(ok bool) //col:874
KdHandleHaltsWhenNmiReceivedFromVmxRoot()(ok bool) //col:927
KdCustomDebuggerBreakSpinlockLock()(ok bool)     //col:999
KdHandleBreakpointAndDebugBreakpoints()(ok bool) //col:1102
KdHandleNmi()(ok bool) //col:1149
KdGuaranteedStepInstruction()(ok bool) //col:1200
KdCheckGuestOperatingModeChanges()(ok bool) //col:1265
KdRegularStepInInstruction()(ok bool) //col:1329
KdRegularStepOver()(ok bool)          //col:1383
KdPerformRegisterEvent()(ok bool) //col:1398
KdPerformAddActionToEvent()(ok bool) //col:1413
KdQuerySystemState()(ok bool) //col:1461
KdPerformEventQueryAndModification()(ok bool)       //col:1587
KdDispatchAndPerformCommandsFromDebugger()(ok bool) //col:2359
KdIsGuestOnUsermode32Bit()(ok bool) //col:2411
KdManageSystemHaltOnVmxRoot()(ok bool) //col:2591
KdBroadcastHaltOnAllCores()(ok bool) //col:2604
KdHaltSystem()(ok bool) //col:2634
}

)
func NewKd() { return &kd{} }
func (k *kd) KdInitializeKernelDebugger() (ok bool) { //col:62
	return true
}

func (k *kd) KdUninitializeKernelDebugger() (ok bool) { //col:101
	return true
}

func (k *kd) KdDummyDPC() (ok bool) { //col:120
	return true
}

func (k *kd) KdFireDpc() (ok bool) { //col:139
	return true
}

func (k *kd) KdComputeDataChecksum() (ok bool) { //col:161
	return true
}

func (k *kd) KdResponsePacketToDebugger() (ok bool) { //col:241
	return true
}

func (k *kd) KdLoggingResponsePacketToDebugger() (ok bool) { //col:296
	return true
}

func (k *kd) KdHandleDebugEventsWhenKernelDebuggerIsAttached() (ok bool) { //col:413
	return true
}

func (k *kd) KdApplyTasksPreHaltCore() (ok bool) { //col:466
	return true
}

func (k *kd) KdApplyTasksPostContinueCore() (ok bool) { //col:520
	return true
}

func (k *kd) * are continued ()(ok bool) { //col:579
	return true
}

func (k *kd) KdContinueDebuggeeJustCurrentCore() (ok bool) { //col:602
	return true
}

func (k *kd) KdReadRegisters() (ok bool) { //col:651
	return true
}

func (k *kd) KdReadMemory() (ok bool) { //col:700
	return true
}

func (k *kd) KdSwitchCore() (ok bool) { //col:775
	return true
}

func (k *kd) KdCloseConnectionAndUnloadDebuggee() (ok bool) { //col:792
	return true
}

func (k *kd) KdReloadSymbolDetailsInDebuggee() (ok bool) { //col:811
	return true
}

func (k *kd) KdNotifyDebuggeeForUserInput() (ok bool) { //col:832
	return true
}

func (k *kd) KdSendFormatsFunctionResult() (ok bool) { //col:856
	return true
}

func (k *kd) KdSendCommandFinishedSignal() (ok bool) { //col:874
	return true
}

func (k *kd) KdHandleHaltsWhenNmiReceivedFromVmxRoot() (ok bool) { //col:927
	return true
}

func (k *kd) KdCustomDebuggerBreakSpinlockLock() (ok bool) { //col:999
	return true
}

func (k *kd) KdHandleBreakpointAndDebugBreakpoints() (ok bool) { //col:1102
	return true
}

func (k *kd) KdHandleNmi() (ok bool) { //col:1149
	return true
}

func (k *kd) KdGuaranteedStepInstruction() (ok bool) { //col:1200
	return true
}

func (k *kd) KdCheckGuestOperatingModeChanges() (ok bool) { //col:1265
	return true
}

func (k *kd) KdRegularStepInInstruction() (ok bool) { //col:1329
	return true
}

func (k *kd) KdRegularStepOver() (ok bool) { //col:1383
	return true
}

func (k *kd) KdPerformRegisterEvent() (ok bool) { //col:1398
	return true
}

func (k *kd) KdPerformAddActionToEvent() (ok bool) { //col:1413
	return true
}

func (k *kd) KdQuerySystemState() (ok bool) { //col:1461
	return true
}

func (k *kd) KdPerformEventQueryAndModification() (ok bool) { //col:1587
	return true
}

func (k *kd) KdDispatchAndPerformCommandsFromDebugger() (ok bool) { //col:2359
	return true
}

func (k *kd) KdIsGuestOnUsermode32Bit() (ok bool) { //col:2411
	return true
}

func (k *kd) KdManageSystemHaltOnVmxRoot() (ok bool) { //col:2591
	return true
}

func (k *kd) KdBroadcastHaltOnAllCores() (ok bool) { //col:2604
	return true
}

func (k *kd) KdHaltSystem() (ok bool) { //col:2634
	return true
}

