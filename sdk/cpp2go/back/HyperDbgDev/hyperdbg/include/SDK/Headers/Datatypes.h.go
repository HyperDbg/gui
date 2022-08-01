package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h.back

const(
SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED = sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED) //col:1
)

type DEBUGGEE_USER_INPUT_PACKET struct{
CommandLen UINT32
IgnoreFinishedSignal bool
Result UINT32
}


type DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct{
Length UINT32
}


type DEBUGGER_PAUSE_PACKET_RECEIVED struct{
Result UINT32
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
ProcessId UINT32
ThreadId UINT32
EventTag uint64
Rflags uint64
InstructionBytesOnRip[MAXIMUM_INSTR_SIZE] BYTE
ReadInstructionLen UINT16
GuestRegs GUEST_REGS
}


type DEBUGGEE_MESSAGE_PACKET struct{
OperationCode UINT32
Message[PacketChunkSize] CHAR
}




