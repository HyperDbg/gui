package user-level


type (
Attaching interface{
AttachingInitialize()(ok bool)//col:87
AttachingCreateProcessDebuggingDetails()(ok bool)//col:156
AttachingFindProcessDebuggingDetailsByToken()(ok bool)//col:180
AttachingFindProcessDebuggingDetailsByProcessId()(ok bool)//col:204
AttachingFindProcessDebuggingDetailsInStartingPhase()(ok bool)//col:225
AttachingRemoveAndFreeAllProcessDebuggingDetails()(ok bool)//col:253
AttachingRemoveProcessDebuggingDetailsByToken()(ok bool)//col:296
AttachingSetStartingPhaseOfProcessDebuggingDetailsByToken()(ok bool)//col:344
AttachingReachedToProcessEntrypoint()(ok bool)//col:387
AttachingHandleEntrypointDebugBreak()(ok bool)//col:535
AttachingAdjustNopSledBuffer()(ok bool)//col:599
AttachingCheckPageFaultsWithUserDebugger()(ok bool)//col:666
AttachingConfigureInterceptingThreads()(ok bool)//col:759
AttachingPerformAttachToProcess()(ok bool)//col:973
AttachingHandleCr3VmexitsForThreadInterception()(ok bool)//col:1040
AttachingRemoveHooks()(ok bool)//col:1102
AttachingPauseProcess()(ok bool)//col:1128
AttachingKillProcess()(ok bool)//col:1203
AttachingPerformDetach()(ok bool)//col:1265
AttachingSwitchProcess()(ok bool)//col:1354
AttachingQueryCountOfActiveDebuggingThreadsAndProcesses()(ok bool)//col:1380
AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses()(ok bool)//col:1413
AttachingTargetProcess()(ok bool)//col:1481
}













































)

func NewAttaching() { return & attaching{} }

func (a *attaching)AttachingInitialize()(ok bool){//col:87

































return true
}














































func (a *attaching)AttachingCreateProcessDebuggingDetails()(ok bool){//col:156






























return true
}














































func (a *attaching)AttachingFindProcessDebuggingDetailsByToken()(ok bool){//col:180











return true
}














































func (a *attaching)AttachingFindProcessDebuggingDetailsByProcessId()(ok bool){//col:204











return true
}














































func (a *attaching)AttachingFindProcessDebuggingDetailsInStartingPhase()(ok bool){//col:225











return true
}














































func (a *attaching)AttachingRemoveAndFreeAllProcessDebuggingDetails()(ok bool){//col:253









return true
}














































func (a *attaching)AttachingRemoveProcessDebuggingDetailsByToken()(ok bool){//col:296













return true
}














































func (a *attaching)AttachingSetStartingPhaseOfProcessDebuggingDetailsByToken()(ok bool){//col:344
















return true
}














































func (a *attaching)AttachingReachedToProcessEntrypoint()(ok bool){//col:387


















return true
}














































func (a *attaching)AttachingHandleEntrypointDebugBreak()(ok bool){//col:535
















































return true
}














































func (a *attaching)AttachingAdjustNopSledBuffer()(ok bool){//col:599
























return true
}














































func (a *attaching)AttachingCheckPageFaultsWithUserDebugger()(ok bool){//col:666




























return true
}














































func (a *attaching)AttachingConfigureInterceptingThreads()(ok bool){//col:759













































return true
}














































func (a *attaching)AttachingPerformAttachToProcess()(ok bool){//col:973












































































































return true
}














































func (a *attaching)AttachingHandleCr3VmexitsForThreadInterception()(ok bool){//col:1040




























return true
}














































func (a *attaching)AttachingRemoveHooks()(ok bool){//col:1102






























return true
}














































func (a *attaching)AttachingPauseProcess()(ok bool){//col:1128













return true
}














































func (a *attaching)AttachingKillProcess()(ok bool){//col:1203
































return true
}














































func (a *attaching)AttachingPerformDetach()(ok bool){//col:1265























return true
}














































func (a *attaching)AttachingSwitchProcess()(ok bool){//col:1354












































return true
}














































func (a *attaching)AttachingQueryCountOfActiveDebuggingThreadsAndProcesses()(ok bool){//col:1380








return true
}














































func (a *attaching)AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses()(ok bool){//col:1413











return true
}














































func (a *attaching)AttachingTargetProcess()(ok bool){//col:1481































return true
}
















































