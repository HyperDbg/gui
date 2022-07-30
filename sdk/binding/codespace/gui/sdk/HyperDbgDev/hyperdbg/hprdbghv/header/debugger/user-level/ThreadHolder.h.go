package user-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\ThreadHolder.h.back

type USERMODE_DEBUGGING_THREAD_DETAILS struct {
	ThreadId             uint32                                                   //col:3
	ThreadRip            uint64                                                   //col:4
	IsPaused             bool                                                     //col:5
	IsRflagsTrapFlagsSet bool                                                     //col:6
	UdAction             [MAX_USER_ACTIONS_FOR_THREADS]DEBUGGER_UD_COMMAND_ACTION //col:7
}

type USERMODE_DEBUGGING_THREAD_HOLDER struct {
	ThreadHolderList *list.List                                                         //col:11
	Threads          [MAX_THREADS_IN_A_PROCESS_HOLDER]USERMODE_DEBUGGING_THREAD_DETAILS //col:12
}
