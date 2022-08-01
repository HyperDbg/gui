package Headers
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h.back

const(
SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED = sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED) //col:1
)

type DEBUGGEE_USER_INPUT_PACKET struct{
CommandLen uint32
IgnoreFinishedSignal bool
Result uint32
}


type DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct{
Length uint32
}


type DEBUGGER_PAUSE_PACKET_RECEIVED struct{
Result uint32
}


type DEBUGGEE_KD_PAUSED_PACKET struct{
Rip uint64
Is32BitAddress bool
PausingReason DEBUGGEE_PAUSING_REASON
CurrentCore ULONG
EventTag uint64
Rflags uint64
InstructionBytesOnRip[MAXIMUM_INSTR_SIZE] BYTE
ReadInstructionLen UINT16
}


type DEBUGGEE_UD_PAUSED_PACKET struct{
Rip uint64
ProcessDebuggingToken uint64
Is32Bit bool
PausingReason DEBUGGEE_PAUSING_REASON
ProcessId uint32
ThreadId uint32
EventTag uint64
Rflags uint64
InstructionBytesOnRip[MAXIMUM_INSTR_SIZE] BYTE
ReadInstructionLen UINT16
GuestRegs GUEST_REGS
}


type DEBUGGEE_MESSAGE_PACKET struct{
OperationCode uint32
Message[PacketChunkSize] CHAR
}




