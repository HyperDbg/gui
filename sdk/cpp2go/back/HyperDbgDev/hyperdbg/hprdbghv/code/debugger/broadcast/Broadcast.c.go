package broadcast


type (
Broadcast interface{
BroadcastVmxVirtualizationAllCores()(ok bool)//col:27
BroadcastEnableDbAndBpExitingAllCores()(ok bool)//col:42
BroadcastDisableDbAndBpExitingAllCores()(ok bool)//col:57
BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()(ok bool)//col:72
BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()(ok bool)//col:87
BroadcastEnableNmiExitingAllCores()(ok bool)//col:102
BroadcastDisableNmiExitingAllCores()(ok bool)//col:117
BroadcastNotifyAllToInvalidateEptAllCores()(ok bool)//col:132
}
















)

func NewBroadcast() { return & broadcast{} }

func (b *broadcast)BroadcastVmxVirtualizationAllCores()(ok bool){//col:27




return true
}

















func (b *broadcast)BroadcastEnableDbAndBpExitingAllCores()(ok bool){//col:42




return true
}

















func (b *broadcast)BroadcastDisableDbAndBpExitingAllCores()(ok bool){//col:57




return true
}

















func (b *broadcast)BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()(ok bool){//col:72




return true
}

















func (b *broadcast)BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()(ok bool){//col:87




return true
}

















func (b *broadcast)BroadcastEnableNmiExitingAllCores()(ok bool){//col:102




return true
}

















func (b *broadcast)BroadcastDisableNmiExitingAllCores()(ok bool){//col:117




return true
}

















func (b *broadcast)BroadcastNotifyAllToInvalidateEptAllCores()(ok bool){//col:132




return true
}



















