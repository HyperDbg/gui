package user-level


type (
ThreadHolder interface{
ThreadHolderAllocateThreadHoldingBuffers()(ok bool)//col:33
ThreadHolderAssignThreadHolderToProcessDebuggingDetails()(ok bool)//col:69
ThreadHolderIsAnyPausedThreadInProcess()(ok bool)//col:101
ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId()(ok bool)//col:151
ThreadHolderGetProcessFirstThreadDetailsByProcessId()(ok bool)//col:200
ThreadHolderGetProcessDebuggingDetailsByThreadId()(ok bool)//col:252
ThreadHolderFindOrCreateThreadDebuggingDetail()(ok bool)//col:358
ThreadHolderApplyActionToPausedThreads()(ok bool)//col:463
ThreadHolderFreeHoldingStructures()(ok bool)//col:491
ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses()(ok bool)//col:547
ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses()(ok bool)//col:630
}

)

func NewThreadHolder() { return & threadHolder{} }

func (t *threadHolder)ThreadHolderAllocateThreadHoldingBuffers()(ok bool){//col:33





return true
}


func (t *threadHolder)ThreadHolderAssignThreadHolderToProcessDebuggingDetails()(ok bool){//col:69













return true
}


func (t *threadHolder)ThreadHolderIsAnyPausedThreadInProcess()(ok bool){//col:101



















return true
}


func (t *threadHolder)ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId()(ok bool){//col:151

























return true
}


func (t *threadHolder)ThreadHolderGetProcessFirstThreadDetailsByProcessId()(ok bool){//col:200

























return true
}


func (t *threadHolder)ThreadHolderGetProcessDebuggingDetailsByThreadId()(ok bool){//col:252



























return true
}


func (t *threadHolder)ThreadHolderFindOrCreateThreadDebuggingDetail()(ok bool){//col:358

















































return true
}


func (t *threadHolder)ThreadHolderApplyActionToPausedThreads()(ok bool){//col:463























































return true
}


func (t *threadHolder)ThreadHolderFreeHoldingStructures()(ok bool){//col:491












return true
}


func (t *threadHolder)ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses()(ok bool){//col:547





























return true
}


func (t *threadHolder)ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses()(ok bool){//col:630















































return true
}




