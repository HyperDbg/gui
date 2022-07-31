package user-level
type (
	ThreadHolder interface {
		ThreadHolderAllocateThreadHoldingBuffers() (ok bool)                     //col:33
		ThreadHolderAssignThreadHolderToProcessDebuggingDetails() (ok bool)      //col:68
		ThreadHolderIsAnyPausedThreadInProcess() (ok bool)                       //col:99
		ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId() (ok bool)    //col:148
		ThreadHolderGetProcessFirstThreadDetailsByProcessId() (ok bool)          //col:196
		ThreadHolderGetProcessDebuggingDetailsByThreadId() (ok bool)             //col:247
		ThreadHolderFindOrCreateThreadDebuggingDetail() (ok bool)                //col:352
		ThreadHolderApplyActionToPausedThreads() (ok bool)                       //col:456
		ThreadHolderFreeHoldingStructures() (ok bool)                            //col:483
		ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses() (ok bool)   //col:538
		ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses() (ok bool) //col:620
	}
)
func NewThreadHolder() { return &threadHolder{} }
func (t *threadHolder) ThreadHolderAllocateThreadHoldingBuffers() (ok bool) { //col:33
	return true
}

func (t *threadHolder) ThreadHolderAssignThreadHolderToProcessDebuggingDetails() (ok bool) { //col:68
	return true
}

func (t *threadHolder) ThreadHolderIsAnyPausedThreadInProcess() (ok bool) { //col:99
	return true
}

func (t *threadHolder) ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId() (ok bool) { //col:148
	return true
}

func (t *threadHolder) ThreadHolderGetProcessFirstThreadDetailsByProcessId() (ok bool) { //col:196
	return true
}

func (t *threadHolder) ThreadHolderGetProcessDebuggingDetailsByThreadId() (ok bool) { //col:247
	return true
}

func (t *threadHolder) ThreadHolderFindOrCreateThreadDebuggingDetail() (ok bool) { //col:352
	return true
}

func (t *threadHolder) ThreadHolderApplyActionToPausedThreads() (ok bool) { //col:456
	return true
}

func (t *threadHolder) ThreadHolderFreeHoldingStructures() (ok bool) { //col:483
	return true
}

func (t *threadHolder) ThreadHolderQueryCountOfActiveDebuggingThreadsAndProcesses() (ok bool) { //col:538
	return true
}

func (t *threadHolder) ThreadHolderQueryDetailsOfActiveDebuggingThreadsAndProcesses() (ok bool) { //col:620
	return true
}

