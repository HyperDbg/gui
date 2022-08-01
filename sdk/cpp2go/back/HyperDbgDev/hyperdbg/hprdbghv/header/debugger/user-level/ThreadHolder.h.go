package user-level
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\ThreadHolder.h.back

type USERMODE_DEBUGGING_THREAD_DETAILS struct{
ThreadId uint32
ThreadRip uint64
IsPaused bool
IsRflagsTrapFlagsSet bool
UdAction[MAX_USER_ACTIONS_FOR_THREADS] DEBUGGER_UD_COMMAND_ACTION
}


type USERMODE_DEBUGGING_THREAD_HOLDER struct{
ThreadHolderList *list.List
Threads[MAX_THREADS_IN_A_PROCESS_HOLDER] USERMODE_DEBUGGING_THREAD_DETAILS
}




