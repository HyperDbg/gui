package user-level
type (
	Ud interface {
		UdInitializeUserDebugger() (ok bool)                  //col:68
		UdUninitializeUserDebugger() (ok bool)                //col:92
		UdRestoreToOriginalDirection() (ok bool)              //col:108
		UdContinueThread() (ok bool)                          //col:134
		UdStepInstructions() (ok bool)                        //col:191
		UdPerformCommand() (ok bool)                          //col:246
		UdCheckForCommand() (ok bool)                         //col:318
		UdDispatchUsermodeCommands() (ok bool)                //col:359
		UdSpinThreadOnNop() (ok bool)                         //col:386
		UdHandleAfterSteppingReason() (ok bool)               //col:415
		UdPrePausingReasons() (ok bool)                       //col:453
		UdCheckAndHandleBreakpointsAndDebugBreaks() (ok bool) //col:634
	}
)
func NewUd() { return &ud{} }
func (u *ud) UdInitializeUserDebugger() (ok bool) { //col:68
	return true
}

func (u *ud) UdUninitializeUserDebugger() (ok bool) { //col:92
	return true
}

func (u *ud) UdRestoreToOriginalDirection() (ok bool) { //col:108
	return true
}

func (u *ud) UdContinueThread() (ok bool) { //col:134
	return true
}

func (u *ud) UdStepInstructions() (ok bool) { //col:191
	return true
}

func (u *ud) UdPerformCommand() (ok bool) { //col:246
	return true
}

func (u *ud) UdCheckForCommand() (ok bool) { //col:318
	return true
}

func (u *ud) UdDispatchUsermodeCommands() (ok bool) { //col:359
	return true
}

func (u *ud) UdSpinThreadOnNop() (ok bool) { //col:386
	return true
}

func (u *ud) UdHandleAfterSteppingReason() (ok bool) { //col:415
	return true
}

func (u *ud) UdPrePausingReasons() (ok bool) { //col:453
	return true
}

func (u *ud) UdCheckAndHandleBreakpointsAndDebugBreaks() (ok bool) { //col:634
	return true
}

