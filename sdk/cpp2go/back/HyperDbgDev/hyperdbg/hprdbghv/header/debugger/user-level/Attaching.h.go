package user-level
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\Attaching.h.back

const(
MAX_USER_ACTIONS_FOR_THREADS = 3 //col:1
MAX_THREADS_IN_A_PROCESS_HOLDER = 100 //col:2
MAX_CR3_IN_A_PROCESS = 4 //col:3
)

type USERMODE_DEBUGGING_PROCESS_DETAILS struct{
Token uint64
Enabled bool
PebAddressToMonitor PVOID
ActiveThreadId uint32
Registers GUEST_REGS
Context uint64
AttachedProcessList *list.List
UsermodeReservedBuffer uint64
EntrypointOfMainModule uint64
BaseAddressOfMainModule uint64
Eprocess PEPROCESS
ProcessId uint32
Is32Bit bool
IsOnTheStartingPhase bool
IsOnThreadInterceptingPhase bool
InterceptedCr3[MAX_CR3_IN_A_PROCESS] CR3_TYPE
ThreadsListHead *list.List
}




