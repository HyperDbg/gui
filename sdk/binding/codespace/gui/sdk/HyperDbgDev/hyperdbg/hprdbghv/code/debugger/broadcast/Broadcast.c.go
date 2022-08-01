package broadcast
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\broadcast\Broadcast.c.back

type (
Broadcast interface{
BroadcastVmxVirtualizationAllCores()(ok bool)//col:4
BroadcastEnableDbAndBpExitingAllCores()(ok bool)//col:8
BroadcastDisableDbAndBpExitingAllCores()(ok bool)//col:12
BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()(ok bool)//col:16
BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()(ok bool)//col:20
BroadcastEnableNmiExitingAllCores()(ok bool)//col:24
BroadcastDisableNmiExitingAllCores()(ok bool)//col:28
BroadcastNotifyAllToInvalidateEptAllCores()(ok bool)//col:32
}
broadcast struct{}
)

func NewBroadcast()Broadcast{ return & broadcast{} }

func (b *broadcast)BroadcastVmxVirtualizationAllCores()(ok bool){//col:4
/*BroadcastVmxVirtualizationAllCores()
{
    KeGenericCallDpc(DpcRoutinePerformVirtualization, NULL);
}*/
return true
}

func (b *broadcast)BroadcastEnableDbAndBpExitingAllCores()(ok bool){//col:8
/*BroadcastEnableDbAndBpExitingAllCores()
{
    KeGenericCallDpc(DpcRoutineEnableDbAndBpExitingOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastDisableDbAndBpExitingAllCores()(ok bool){//col:12
/*BroadcastDisableDbAndBpExitingAllCores()
{
    KeGenericCallDpc(DpcRoutineDisableDbAndBpExitingOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()(ok bool){//col:16
/*BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()
{
    KeGenericCallDpc(DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()(ok bool){//col:20
/*BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()
{
    KeGenericCallDpc(DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastEnableNmiExitingAllCores()(ok bool){//col:24
/*BroadcastEnableNmiExitingAllCores()
{
    KeGenericCallDpc(DpcRoutineEnableNmiVmexitOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastDisableNmiExitingAllCores()(ok bool){//col:28
/*BroadcastDisableNmiExitingAllCores()
{
    KeGenericCallDpc(DpcRoutineDisableNmiVmexitOnAllCores, NULL);
}*/
return true
}

func (b *broadcast)BroadcastNotifyAllToInvalidateEptAllCores()(ok bool){//col:32
/*BroadcastNotifyAllToInvalidateEptAllCores()
{
    KeGenericCallDpc(DpcRoutineInvalidateEptOnAllCores, g_EptState->EptPointer.AsUInt);
}*/
return true
}



