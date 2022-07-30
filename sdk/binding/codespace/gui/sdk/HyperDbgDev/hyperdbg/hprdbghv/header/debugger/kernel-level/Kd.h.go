package kernel-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\kernel-level\Kd.h.back

type DEBUGGEE_REQUEST_TO_CHANGE_PROCESS struct {
	ProcessId uint32 //col:3
	Process   uint64 //col:4
}

type DEBUGGEE_REQUEST_TO_CHANGE_THREAD struct {
	ThreadId uint32 //col:8
	Thread   uint64 //col:9
}

type DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT struct {
	BOOLEAN             volatile                                //col:13
	SpeialEventResponse DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION //col:14
}

type HARDWARE_DEBUG_REGISTER_DETAILS struct {
	Address   uint64 //col:18
	ProcessId uint32 //col:19
	ThreadId  uint32 //col:20
}
