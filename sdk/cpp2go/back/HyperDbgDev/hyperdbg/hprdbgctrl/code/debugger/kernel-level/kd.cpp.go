package kernel-level
type (
Kd interface{
KdCheckForTheEndOfTheBuffer()(ok bool)//col:88
KdCompareBufferWithString()(ok bool)//col:108
KdComputeDataChecksum()(ok bool)//col:129
KdInterpretPausedDebuggee()(ok bool)//col:144
KdSendContinuePacketToDebuggee()(ok bool)//col:170
KdSendSwitchCorePacketToDebuggee()(ok bool)//col:214
KdSendEventQueryAndModifyPacketToDebuggee()(ok bool)//col:267
KdSendFlushPacketToDebuggee()(ok bool)//col:297
KdSendCallStackPacketToDebuggee()(ok bool)//col:376
KdSendTestQueryPacketToDebuggee()(ok bool)//col:409
KdSendSymbolReloadPacketToDebuggee()(ok bool)//col:441
KdSendReadRegisterPacketToDebuggee()(ok bool)//col:469
KdSendReadMemoryPacketToDebuggee()(ok bool)//col:515
KdSendEditMemoryPacketToDebuggee()(ok bool)//col:545
KdSendRegisterEventPacketToDebuggee()(ok bool)//col:612
KdSendAddActionToEventPacketToDebuggee()(ok bool)//col:679
KdSendSwitchProcessPacketToDebuggee()(ok bool)//col:732
KdSendSwitchThreadPacketToDebuggee()(ok bool)//col:785
KdSendPtePacketToDebuggee()(ok bool)//col:814
KdSendVa2paAndPa2vaPacketToDebuggee()(ok bool)//col:843
KdSendBpPacketToDebuggee()(ok bool)//col:872
KdSendListOrModifyPacketToDebuggee()(ok bool)//col:902
KdSendScriptPacketToDebuggee()(ok bool)//col:967
KdSendUserInputPacketToDebuggee()(ok bool)//col:1026
KdSendSearchRequestPacketToDebuggee()(ok bool)//col:1056
KdSendStepPacketToDebuggee()(ok bool)//col:1115
KdSendPausePacketToDebuggee()(ok bool)//col:1141
KdGetWindowVersion()(ok bool)//col:1208
KdReceivePacketFromDebuggee()(ok bool)//col:1309
KdSendPacketToDebuggee()(ok bool)//col:1434
KdCommandPacketToDebuggee()(ok bool)//col:1479
KdCommandPacketAndBufferToDebuggee()(ok bool)//col:1550
KdBreakControlCheckAndPauseDebugger()(ok bool)//col:1579
KdSetStatusAndWaitForPause()(ok bool)//col:1598
KdBreakControlCheckAndContinueDebugger()(ok bool)//col:1628
KdTheRemoteSystemIsRunning()(ok bool)//col:1648
KdPrepareSerialConnectionToRemoteSystem()(ok bool)//col:1785
KdPrepareAndConnectDebugPort()(ok bool)//col:2179
KdSendGeneralBuffersFromDebuggeeToDebugger()(ok bool)//col:2283
KdReloadSymbolsInDebuggee()(ok bool)//col:2326
KdCloseConnection()(ok bool)//col:2404
KdRegisterEventInDebuggee()(ok bool)//col:2458
KdAddActionToEventInDebuggee()(ok bool)//col:2508
KdSendModifyEventInDebuggee()(ok bool)//col:2562
KdHandleUserInputInDebuggee()(ok bool)//col:2617
KdSendUsermodePrints()(ok bool)//col:2675
KdSendSymbolDetailPacket()(ok bool)//col:2727
KdUninitializeConnection()(ok bool)//col:2831
}

)
func NewKd() { return & kd{} }
func (k *kd)KdCheckForTheEndOfTheBuffer()(ok bool){//col:88
return true
}

func (k *kd)KdCompareBufferWithString()(ok bool){//col:108
return true
}

func (k *kd)KdComputeDataChecksum()(ok bool){//col:129
return true
}

func (k *kd)KdInterpretPausedDebuggee()(ok bool){//col:144
return true
}

func (k *kd)KdSendContinuePacketToDebuggee()(ok bool){//col:170
return true
}

func (k *kd)KdSendSwitchCorePacketToDebuggee()(ok bool){//col:214
return true
}

func (k *kd)KdSendEventQueryAndModifyPacketToDebuggee()(ok bool){//col:267
return true
}

func (k *kd)KdSendFlushPacketToDebuggee()(ok bool){//col:297
return true
}

func (k *kd)KdSendCallStackPacketToDebuggee()(ok bool){//col:376
return true
}

func (k *kd)KdSendTestQueryPacketToDebuggee()(ok bool){//col:409
return true
}

func (k *kd)KdSendSymbolReloadPacketToDebuggee()(ok bool){//col:441
return true
}

func (k *kd)KdSendReadRegisterPacketToDebuggee()(ok bool){//col:469
return true
}

func (k *kd)KdSendReadMemoryPacketToDebuggee()(ok bool){//col:515
return true
}

func (k *kd)KdSendEditMemoryPacketToDebuggee()(ok bool){//col:545
return true
}

func (k *kd)KdSendRegisterEventPacketToDebuggee()(ok bool){//col:612
return true
}

func (k *kd)KdSendAddActionToEventPacketToDebuggee()(ok bool){//col:679
return true
}

func (k *kd)KdSendSwitchProcessPacketToDebuggee()(ok bool){//col:732
return true
}

func (k *kd)KdSendSwitchThreadPacketToDebuggee()(ok bool){//col:785
return true
}

func (k *kd)KdSendPtePacketToDebuggee()(ok bool){//col:814
return true
}

func (k *kd)KdSendVa2paAndPa2vaPacketToDebuggee()(ok bool){//col:843
return true
}

func (k *kd)KdSendBpPacketToDebuggee()(ok bool){//col:872
return true
}

func (k *kd)KdSendListOrModifyPacketToDebuggee()(ok bool){//col:902
return true
}

func (k *kd)KdSendScriptPacketToDebuggee()(ok bool){//col:967
return true
}

func (k *kd)KdSendUserInputPacketToDebuggee()(ok bool){//col:1026
return true
}

func (k *kd)KdSendSearchRequestPacketToDebuggee()(ok bool){//col:1056
return true
}

func (k *kd)KdSendStepPacketToDebuggee()(ok bool){//col:1115
return true
}

func (k *kd)KdSendPausePacketToDebuggee()(ok bool){//col:1141
return true
}

func (k *kd)KdGetWindowVersion()(ok bool){//col:1208
return true
}

func (k *kd)KdReceivePacketFromDebuggee()(ok bool){//col:1309
return true
}

func (k *kd)KdSendPacketToDebuggee()(ok bool){//col:1434
return true
}

func (k *kd)KdCommandPacketToDebuggee()(ok bool){//col:1479
return true
}

func (k *kd)KdCommandPacketAndBufferToDebuggee()(ok bool){//col:1550
return true
}

func (k *kd)KdBreakControlCheckAndPauseDebugger()(ok bool){//col:1579
return true
}

func (k *kd)KdSetStatusAndWaitForPause()(ok bool){//col:1598
return true
}

func (k *kd)KdBreakControlCheckAndContinueDebugger()(ok bool){//col:1628
return true
}

func (k *kd)KdTheRemoteSystemIsRunning()(ok bool){//col:1648
return true
}

func (k *kd)KdPrepareSerialConnectionToRemoteSystem()(ok bool){//col:1785
return true
}

func (k *kd)KdPrepareAndConnectDebugPort()(ok bool){//col:2179
return true
}

func (k *kd)KdSendGeneralBuffersFromDebuggeeToDebugger()(ok bool){//col:2283
return true
}

func (k *kd)KdReloadSymbolsInDebuggee()(ok bool){//col:2326
return true
}

func (k *kd)KdCloseConnection()(ok bool){//col:2404
return true
}

func (k *kd)KdRegisterEventInDebuggee()(ok bool){//col:2458
return true
}

func (k *kd)KdAddActionToEventInDebuggee()(ok bool){//col:2508
return true
}

func (k *kd)KdSendModifyEventInDebuggee()(ok bool){//col:2562
return true
}

func (k *kd)KdHandleUserInputInDebuggee()(ok bool){//col:2617
return true
}

func (k *kd)KdSendUsermodePrints()(ok bool){//col:2675
return true
}

func (k *kd)KdSendSymbolDetailPacket()(ok bool){//col:2727
return true
}

func (k *kd)KdUninitializeConnection()(ok bool){//col:2831
return true
}

