package kernel-level


type (
Kd interface{
KdCheckForTheEndOfTheBuffer()(ok bool)//col:88
KdCompareBufferWithString()(ok bool)//col:109
KdComputeDataChecksum()(ok bool)//col:131
KdInterpretPausedDebuggee()(ok bool)//col:147
KdSendContinuePacketToDebuggee()(ok bool)//col:174
KdSendSwitchCorePacketToDebuggee()(ok bool)//col:219
KdSendEventQueryAndModifyPacketToDebuggee()(ok bool)//col:273
KdSendFlushPacketToDebuggee()(ok bool)//col:304
KdSendCallStackPacketToDebuggee()(ok bool)//col:384
KdSendTestQueryPacketToDebuggee()(ok bool)//col:418
KdSendSymbolReloadPacketToDebuggee()(ok bool)//col:451
KdSendReadRegisterPacketToDebuggee()(ok bool)//col:480
KdSendReadMemoryPacketToDebuggee()(ok bool)//col:527
KdSendEditMemoryPacketToDebuggee()(ok bool)//col:558
KdSendRegisterEventPacketToDebuggee()(ok bool)//col:626
KdSendAddActionToEventPacketToDebuggee()(ok bool)//col:694
KdSendSwitchProcessPacketToDebuggee()(ok bool)//col:748
KdSendSwitchThreadPacketToDebuggee()(ok bool)//col:802
KdSendPtePacketToDebuggee()(ok bool)//col:832
KdSendVa2paAndPa2vaPacketToDebuggee()(ok bool)//col:862
KdSendBpPacketToDebuggee()(ok bool)//col:892
KdSendListOrModifyPacketToDebuggee()(ok bool)//col:923
KdSendScriptPacketToDebuggee()(ok bool)//col:989
KdSendUserInputPacketToDebuggee()(ok bool)//col:1049
KdSendSearchRequestPacketToDebuggee()(ok bool)//col:1080
KdSendStepPacketToDebuggee()(ok bool)//col:1140
KdSendPausePacketToDebuggee()(ok bool)//col:1167
KdGetWindowVersion()(ok bool)//col:1235
KdReceivePacketFromDebuggee()(ok bool)//col:1337
KdSendPacketToDebuggee()(ok bool)//col:1463
KdCommandPacketToDebuggee()(ok bool)//col:1509
KdCommandPacketAndBufferToDebuggee()(ok bool)//col:1581
KdBreakControlCheckAndPauseDebugger()(ok bool)//col:1611
KdSetStatusAndWaitForPause()(ok bool)//col:1631
KdBreakControlCheckAndContinueDebugger()(ok bool)//col:1662
KdTheRemoteSystemIsRunning()(ok bool)//col:1683
KdPrepareSerialConnectionToRemoteSystem()(ok bool)//col:1821
KdPrepareAndConnectDebugPort()(ok bool)//col:2216
KdSendGeneralBuffersFromDebuggeeToDebugger()(ok bool)//col:2321
KdReloadSymbolsInDebuggee()(ok bool)//col:2365
KdCloseConnection()(ok bool)//col:2444
KdRegisterEventInDebuggee()(ok bool)//col:2499
KdAddActionToEventInDebuggee()(ok bool)//col:2550
KdSendModifyEventInDebuggee()(ok bool)//col:2605
KdHandleUserInputInDebuggee()(ok bool)//col:2661
KdSendUsermodePrints()(ok bool)//col:2720
KdSendSymbolDetailPacket()(ok bool)//col:2773
KdUninitializeConnection()(ok bool)//col:2878
}

)

func NewKd() { return & kd{} }

func (k *kd)KdCheckForTheEndOfTheBuffer()(ok bool){//col:88






















return true
}


func (k *kd)KdCompareBufferWithString()(ok bool){//col:109









return true
}


func (k *kd)KdComputeDataChecksum()(ok bool){//col:131












return true
}


func (k *kd)KdInterpretPausedDebuggee()(ok bool){//col:147




return true
}


func (k *kd)KdSendContinuePacketToDebuggee()(ok bool){//col:174











return true
}


func (k *kd)KdSendSwitchCorePacketToDebuggee()(ok bool){//col:219





















return true
}


func (k *kd)KdSendEventQueryAndModifyPacketToDebuggee()(ok bool){//col:273
























return true
}


func (k *kd)KdSendFlushPacketToDebuggee()(ok bool){//col:304














return true
}


func (k *kd)KdSendCallStackPacketToDebuggee()(ok bool){//col:384














































return true
}


func (k *kd)KdSendTestQueryPacketToDebuggee()(ok bool){//col:418















return true
}


func (k *kd)KdSendSymbolReloadPacketToDebuggee()(ok bool){//col:451















return true
}


func (k *kd)KdSendReadRegisterPacketToDebuggee()(ok bool){//col:480













return true
}


func (k *kd)KdSendReadMemoryPacketToDebuggee()(ok bool){//col:527

























return true
}


func (k *kd)KdSendEditMemoryPacketToDebuggee()(ok bool){//col:558













return true
}


func (k *kd)KdSendRegisterEventPacketToDebuggee()(ok bool){//col:626

































return true
}


func (k *kd)KdSendAddActionToEventPacketToDebuggee()(ok bool){//col:694

































return true
}


func (k *kd)KdSendSwitchProcessPacketToDebuggee()(ok bool){//col:748


























return true
}


func (k *kd)KdSendSwitchThreadPacketToDebuggee()(ok bool){//col:802


























return true
}


func (k *kd)KdSendPtePacketToDebuggee()(ok bool){//col:832













return true
}


func (k *kd)KdSendVa2paAndPa2vaPacketToDebuggee()(ok bool){//col:862













return true
}


func (k *kd)KdSendBpPacketToDebuggee()(ok bool){//col:892













return true
}


func (k *kd)KdSendListOrModifyPacketToDebuggee()(ok bool){//col:923














return true
}


func (k *kd)KdSendScriptPacketToDebuggee()(ok bool){//col:989






























return true
}


func (k *kd)KdSendUserInputPacketToDebuggee()(ok bool){//col:1049




























return true
}


func (k *kd)KdSendSearchRequestPacketToDebuggee()(ok bool){//col:1080













return true
}


func (k *kd)KdSendStepPacketToDebuggee()(ok bool){//col:1140



























return true
}


func (k *kd)KdSendPausePacketToDebuggee()(ok bool){//col:1167











return true
}


func (k *kd)KdGetWindowVersion()(ok bool){//col:1235








































return true
}


func (k *kd)KdReceivePacketFromDebuggee()(ok bool){//col:1337












































return true
}


func (k *kd)KdSendPacketToDebuggee()(ok bool){//col:1463






























































return true
}


func (k *kd)KdCommandPacketToDebuggee()(ok bool){//col:1509



















return true
}


func (k *kd)KdCommandPacketAndBufferToDebuggee()(ok bool){//col:1581


































return true
}


func (k *kd)KdBreakControlCheckAndPauseDebugger()(ok bool){//col:1611











return true
}


func (k *kd)KdSetStatusAndWaitForPause()(ok bool){//col:1631





return true
}


func (k *kd)KdBreakControlCheckAndContinueDebugger()(ok bool){//col:1662














return true
}


func (k *kd)KdTheRemoteSystemIsRunning()(ok bool){//col:1683





return true
}


func (k *kd)KdPrepareSerialConnectionToRemoteSystem()(ok bool){//col:1821








































return true
}


func (k *kd)KdPrepareAndConnectDebugPort()(ok bool){//col:2216









































































































































return true
}


func (k *kd)KdSendGeneralBuffersFromDebuggeeToDebugger()(ok bool){//col:2321















































return true
}


func (k *kd)KdReloadSymbolsInDebuggee()(ok bool){//col:2365















return true
}


func (k *kd)KdCloseConnection()(ok bool){//col:2444






































return true
}


func (k *kd)KdRegisterEventInDebuggee()(ok bool){//col:2499





















return true
}


func (k *kd)KdAddActionToEventInDebuggee()(ok bool){//col:2550




















return true
}


func (k *kd)KdSendModifyEventInDebuggee()(ok bool){//col:2605

















return true
}


func (k *kd)KdHandleUserInputInDebuggee()(ok bool){//col:2661



















return true
}


func (k *kd)KdSendUsermodePrints()(ok bool){//col:2720

























return true
}


func (k *kd)KdSendSymbolDetailPacket()(ok bool){//col:2773


























return true
}


func (k *kd)KdUninitializeConnection()(ok bool){//col:2878













































return true
}




