package kernel-level
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\kernel-level\Kd.h.back

type DEBUGGEE_REQUEST_TO_CHANGE_PROCESS struct{
ProcessId UINT32
Process uint64
}


type DEBUGGEE_REQUEST_TO_CHANGE_THREAD struct{
ThreadId UINT32
Thread uint64
}


type DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT struct{
BOOLEAN volatile
SpeialEventResponse DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
}


type HARDWARE_DEBUG_REGISTER_DETAILS struct{
Address uint64
ProcessId UINT32
ThreadId UINT32
}




