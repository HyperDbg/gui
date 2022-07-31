package broadcast
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
	return true
}

func (b *broadcast) BroadcastEnableDbAndBpExitingAllCores() (ok bool) { //col:41
	return true
}

func (b *broadcast) BroadcastDisableDbAndBpExitingAllCores() (ok bool) { //col:55
	return true
}

func (b *broadcast) BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores() (ok bool) { //col:69
	return true
}

func (b *broadcast) BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores() (ok bool) { //col:83
	return true
}

func (b *broadcast) BroadcastEnableNmiExitingAllCores() (ok bool) { //col:97
	return true
}

func (b *broadcast) BroadcastDisableNmiExitingAllCores() (ok bool) { //col:111
	return true
}

func (b *broadcast) BroadcastNotifyAllToInvalidateEptAllCores() (ok bool) { //col:125
	return true
}

