package user-level


type (
Ud interface{
UdInitializeUserDebugger()(ok bool)//col:68
UdUninitializeUserDebugger()(ok bool)//col:93
UdRestoreToOriginalDirection()(ok bool)//col:110
UdContinueThread()(ok bool)//col:137
UdStepInstructions()(ok bool)//col:195
UdPerformCommand()(ok bool)//col:251
UdCheckForCommand()(ok bool)//col:324
UdDispatchUsermodeCommands()(ok bool)//col:366
UdSpinThreadOnNop()(ok bool)//col:394
UdHandleAfterSteppingReason()(ok bool)//col:424
UdPrePausingReasons()(ok bool)//col:463
UdCheckAndHandleBreakpointsAndDebugBreaks()(ok bool)//col:645
}

)

func NewUd() { return & ud{} }

func (u *ud)UdInitializeUserDebugger()(ok bool){//col:68

















return true
}


func (u *ud)UdUninitializeUserDebugger()(ok bool){//col:93








return true
}


func (u *ud)UdRestoreToOriginalDirection()(ok bool){//col:110




return true
}


func (u *ud)UdContinueThread()(ok bool){//col:137






return true
}


func (u *ud)UdStepInstructions()(ok bool){//col:195





















return true
}


func (u *ud)UdPerformCommand()(ok bool){//col:251





















return true
}


func (u *ud)UdCheckForCommand()(ok bool){//col:324


































return true
}


func (u *ud)UdDispatchUsermodeCommands()(ok bool){//col:366














return true
}


func (u *ud)UdSpinThreadOnNop()(ok bool){//col:394







return true
}


func (u *ud)UdHandleAfterSteppingReason()(ok bool){//col:424









return true
}


func (u *ud)UdPrePausingReasons()(ok bool){//col:463


















return true
}


func (u *ud)UdCheckAndHandleBreakpointsAndDebugBreaks()(ok bool){//col:645







































































return true
}




