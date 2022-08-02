package user_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\ThreadHolder.h.back

type _USERMODE_DEBUGGING_THREAD_DETAILS struct {
	ThreadId             uint32                                                   //col:10
	ThreadRip            uint64                                                   //col:11
	IsPaused             bool                                                     //col:12
	IsRflagsTrapFlagsSet bool                                                     //col:13
	UdAction             [MAX_USER_ACTIONS_FOR_THREADS]DEBUGGER_UD_COMMAND_ACTION //col:14
}

type _USERMODE_DEBUGGING_THREAD_HOLDER struct {
	ThreadHolderList *list.List                                                         //col:15
	Threads          [MAX_THREADS_IN_A_PROCESS_HOLDER]USERMODE_DEBUGGING_THREAD_DETAILS //col:16
}

