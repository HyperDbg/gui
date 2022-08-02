package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\ud.h.back

type _ACTIVE_DEBUGGING_PROCESS struct {
	IsActive              bool       //col:13
	ProcessDebuggingToken uint64     //col:14
	ProcessId             uint32     //col:15
	ThreadId              uint32     //col:16
	IsPaused              bool       //col:17
	Registers             GUEST_REGS //col:18
	Context               uint64     //col:19
	Is32Bit               bool       //col:20
}

