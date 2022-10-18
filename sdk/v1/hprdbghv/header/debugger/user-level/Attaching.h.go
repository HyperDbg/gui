package user_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\Attaching.h.back

type _USERMODE_DEBUGGING_PROCESS_DETAILS struct {
	Token                       uint64                         //col:22
	Enabled                     bool                           //col:23
	PebAddressToMonitor         uintptr                        //col:24
	ActiveThreadId              uint32                         //col:25
	Registers                   GUEST_REGS                     //col:26
	Context                     uint64                         //col:27
	AttachedProcessList         *list.List                     //col:28
	UsermodeReservedBuffer      uint64                         //col:29
	EntrypointOfMainModule      uint64                         //col:30
	BaseAddressOfMainModule     uint64                         //col:31
	Eprocess                    PEPROCESS                      //col:32
	ProcessId                   uint32                         //col:33
	Is32Bit                     bool                           //col:34
	IsOnTheStartingPhase        bool                           //col:35
	IsOnThreadInterceptingPhase bool                           //col:36
	InterceptedCr3              [MAX_CR3_IN_A_PROCESS]CR3_TYPE //col:37
	ThreadsListHead             *list.List                     //col:38
}

