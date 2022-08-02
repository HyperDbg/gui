package user_level



type USERMODE_DEBUGGING_PROCESS_DETAILS struct {
	Token                       uint64                         //col:3
	Enabled                     bool                           //col:4
	PebAddressToMonitor         PVOID                          //col:5
	ActiveThreadId              uint32                         //col:6
	Registers                   GUEST_REGS                     //col:7
	Context                     uint64                         //col:8
	AttachedProcessList         *list.List                     //col:9
	UsermodeReservedBuffer      uint64                         //col:10
	EntrypointOfMainModule      uint64                         //col:11
	BaseAddressOfMainModule     uint64                         //col:12
	Eprocess                    PEPROCESS                      //col:13
	ProcessId                   uint32                         //col:14
	Is32Bit                     bool                           //col:15
	IsOnTheStartingPhase        bool                           //col:16
	IsOnThreadInterceptingPhase bool                           //col:17
	InterceptedCr3              [MAX_CR3_IN_A_PROCESS]CR3_TYPE //col:18
	ThreadsListHead             *list.List                     //col:19
}


