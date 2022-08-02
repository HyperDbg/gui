package header



type ACTIVE_DEBUGGING_PROCESS struct {
	IsActive              bool       //col:3
	ProcessDebuggingToken uint64     //col:4
	ProcessId             uint32     //col:5
	ThreadId              uint32     //col:6
	IsPaused              bool       //col:7
	Registers             GUEST_REGS //col:8
	Context               uint64     //col:9
	Is32Bit               bool       //col:10
}


