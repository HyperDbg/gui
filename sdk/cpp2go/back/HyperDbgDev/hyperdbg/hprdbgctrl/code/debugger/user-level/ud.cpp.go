package user-level
type (
Ud interface{
UdInitializeUserDebugger()(ok bool)//col:48
UdUninitializeUserDebugger()(ok bool)//col:88
UdSetActiveDebuggingProcess()(ok bool)//col:120
UdRemoveActiveDebuggingProcess()(ok bool)//col:135
UdPrintError()(ok bool)//col:174
UdListProcessThreads()(ok bool)//col:232
UdCheckThreadByProcessId()(ok bool)//col:295
UdCreateSuspendedProcess()(ok bool)//col:335
UdAttachToProcess()(ok bool)//col:536
UdKillProcess()(ok bool)//col:611
UdDetachProcess()(ok bool)//col:693
UdPauseProcess()(ok bool)//col:761
UdSendCommand()(ok bool)//col:828
UdContinueDebuggee()(ok bool)//col:850
UdSendStepPacketToDebuggee()(ok bool)//col:887
UdSetActiveDebuggingThreadByPidOrTid()(ok bool)//col:972
UdShowListActiveDebuggingProcessesAndThreads()(ok bool)//col:1122
}

)
func NewUd() { return & ud{} }
func (u *ud)UdInitializeUserDebugger()(ok bool){//col:48
return true
}

func (u *ud)UdUninitializeUserDebugger()(ok bool){//col:88
return true
}

func (u *ud)UdSetActiveDebuggingProcess()(ok bool){//col:120
return true
}

func (u *ud)UdRemoveActiveDebuggingProcess()(ok bool){//col:135
return true
}

func (u *ud)UdPrintError()(ok bool){//col:174
return true
}

func (u *ud)UdListProcessThreads()(ok bool){//col:232
return true
}

func (u *ud)UdCheckThreadByProcessId()(ok bool){//col:295
return true
}

func (u *ud)UdCreateSuspendedProcess()(ok bool){//col:335
return true
}

func (u *ud)UdAttachToProcess()(ok bool){//col:536
return true
}

func (u *ud)UdKillProcess()(ok bool){//col:611
return true
}

func (u *ud)UdDetachProcess()(ok bool){//col:693
return true
}

func (u *ud)UdPauseProcess()(ok bool){//col:761
return true
}

func (u *ud)UdSendCommand()(ok bool){//col:828
return true
}

func (u *ud)UdContinueDebuggee()(ok bool){//col:850
return true
}

func (u *ud)UdSendStepPacketToDebuggee()(ok bool){//col:887
return true
}

func (u *ud)UdSetActiveDebuggingThreadByPidOrTid()(ok bool){//col:972
return true
}

func (u *ud)UdShowListActiveDebuggingProcessesAndThreads()(ok bool){//col:1122
return true
}

