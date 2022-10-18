package kernel_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\kernel-level\Kd.h.back

type _DEBUGGEE_REQUEST_TO_CHANGE_PROCESS struct {
	ProcessId uint32 //col:7
	Process   uint64 //col:8
}

type _DEBUGGEE_REQUEST_TO_CHANGE_THREAD struct {
	ThreadId uint32 //col:12
	Thread   uint64 //col:13
}

type _DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT struct {
	PauseBreaksUntilSpecialMessageSent bool                                    //col:17
	SpeialEventResponse                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION //col:18
}

type _HARDWARE_DEBUG_REGISTER_DETAILS struct {
	Address   uint64 //col:23
	ProcessId uint32 //col:24
	ThreadId  uint32 //col:25
}

