package broadcast

//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\broadcast\Broadcast.c.back

type (
	Broadcast interface {
		BroadcastVmxVirtualizationAllCores() (ok bool)                         //col:27
		BroadcastEnableDbAndBpExitingAllCores() (ok bool)                      //col:41
		BroadcastDisableDbAndBpExitingAllCores() (ok bool)                     //col:55
		BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores() (ok bool)  //col:69
		BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores() (ok bool) //col:83
		BroadcastEnableNmiExitingAllCores() (ok bool)                          //col:97
		BroadcastDisableNmiExitingAllCores() (ok bool)                         //col:111
		BroadcastNotifyAllToInvalidateEptAllCores() (ok bool)                  //col:125
	}
)

func NewBroadcast() { return &broadcast{} }

func (b *broadcast) BroadcastVmxVirtualizationAllCores() (ok bool) { //col:27
	/*BroadcastVmxVirtualizationAllCores()
	  {
	      KeGenericCallDpc(DpcRoutinePerformVirtualization, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastEnableDbAndBpExitingAllCores() (ok bool) { //col:41
	/*BroadcastEnableDbAndBpExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableDbAndBpExitingOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastDisableDbAndBpExitingAllCores() (ok bool) { //col:55
	/*BroadcastDisableDbAndBpExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableDbAndBpExitingOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores() (ok bool) { //col:69
	/*BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores() (ok bool) { //col:83
	/*BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastEnableNmiExitingAllCores() (ok bool) { //col:97
	/*BroadcastEnableNmiExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableNmiVmexitOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastDisableNmiExitingAllCores() (ok bool) { //col:111
	/*BroadcastDisableNmiExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableNmiVmexitOnAllCores, NULL);
	  }*/
	return true
}

func (b *broadcast) BroadcastNotifyAllToInvalidateEptAllCores() (ok bool) { //col:125
	/*BroadcastNotifyAllToInvalidateEptAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineInvalidateEptOnAllCores, g_EptState->EptPointer.AsUInt);
	  }*/
	return true
}
