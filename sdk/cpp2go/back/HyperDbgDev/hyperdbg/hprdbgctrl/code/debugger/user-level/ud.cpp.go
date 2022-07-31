package user-level


type (
Ud interface{
UdInitializeUserDebugger()(ok bool)//col:48
UdUninitializeUserDebugger()(ok bool)//col:89
UdSetActiveDebuggingProcess()(ok bool)//col:122
UdRemoveActiveDebuggingProcess()(ok bool)//col:138
UdPrintError()(ok bool)//col:178
UdListProcessThreads()(ok bool)//col:237
UdCheckThreadByProcessId()(ok bool)//col:301
UdCreateSuspendedProcess()(ok bool)//col:342
UdAttachToProcess()(ok bool)//col:544
UdKillProcess()(ok bool)//col:620
UdDetachProcess()(ok bool)//col:703
UdPauseProcess()(ok bool)//col:772
UdSendCommand()(ok bool)//col:840
UdContinueDebuggee()(ok bool)//col:863
UdSendStepPacketToDebuggee()(ok bool)//col:901
UdSetActiveDebuggingThreadByPidOrTid()(ok bool)//col:987
UdShowListActiveDebuggingProcessesAndThreads()(ok bool)//col:1138
}






)

func NewUd() { return & ud{} }

func (u *ud)UdInitializeUserDebugger()(ok bool){//col:48













return true
}







func (u *ud)UdUninitializeUserDebugger()(ok bool){//col:89




















return true
}







func (u *ud)UdSetActiveDebuggingProcess()(ok bool){//col:122













return true
}







func (u *ud)UdRemoveActiveDebuggingProcess()(ok bool){//col:138




return true
}







func (u *ud)UdPrintError()(ok bool){//col:178





















return true
}







func (u *ud)UdListProcessThreads()(ok bool){//col:237
























return true
}







func (u *ud)UdCheckThreadByProcessId()(ok bool){//col:301



























return true
}







func (u *ud)UdCreateSuspendedProcess()(ok bool){//col:342























return true
}







func (u *ud)UdAttachToProcess()(ok bool){//col:544






















































































return true
}







func (u *ud)UdKillProcess()(ok bool){//col:620


























return true
}







func (u *ud)UdDetachProcess()(ok bool){//col:703



























return true
}







func (u *ud)UdPauseProcess()(ok bool){//col:772

























return true
}







func (u *ud)UdSendCommand()(ok bool){//col:840






























return true
}







func (u *ud)UdContinueDebuggee()(ok bool){//col:863











return true
}







func (u *ud)UdSendStepPacketToDebuggee()(ok bool){//col:901



















return true
}







func (u *ud)UdSetActiveDebuggingThreadByPidOrTid()(ok bool){//col:987





































return true
}







func (u *ud)UdShowListActiveDebuggingProcessesAndThreads()(ok bool){//col:1138














































































return true
}









