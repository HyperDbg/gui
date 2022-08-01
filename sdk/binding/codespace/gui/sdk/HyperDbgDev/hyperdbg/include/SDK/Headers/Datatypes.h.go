package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h.back

const (
	SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED = sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED) //col:1
)

type DEBUGGEE_USER_INPUT_PACKET struct {
	CommandLen           uint32 //col:3
	IgnoreFinishedSignal bool   //col:4
	Result               uint32 //col:5
}

type DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct {
	Length uint32 //col:9
}

type DEBUGGER_PAUSE_PACKET_RECEIVED struct {
	Result uint32 //col:13
}

type DEBUGGEE_KD_PAUSED_PACKET struct {
	Rip                   uint64                    //col:17
	Is32BitAddress        bool                      //col:18
	PausingReason         DEBUGGEE_PAUSING_REASON   //col:19
	CurrentCore           uint32                    //col:20
	EventTag              uint64                    //col:21
	Rflags                uint64                    //col:22
	InstructionBytesOnRip [MAXIMUM_INSTR_SIZE]uint8 //col:23
	ReadInstructionLen    uint16                    //col:24
}

type DEBUGGEE_UD_PAUSED_PACKET struct {
	Rip                   uint64                    //col:28
	ProcessDebuggingToken uint64                    //col:29
	Is32Bit               bool                      //col:30
	PausingReason         DEBUGGEE_PAUSING_REASON   //col:31
	ProcessId             uint32                    //col:32
	ThreadId              uint32                    //col:33
	EventTag              uint64                    //col:34
	Rflags                uint64                    //col:35
	InstructionBytesOnRip [MAXIMUM_INSTR_SIZE]uint8 //col:36
	ReadInstructionLen    uint16                    //col:37
	GuestRegs             GUEST_REGS                //col:38
}

type DEBUGGEE_MESSAGE_PACKET struct {
	OperationCode uint32                //col:42
	Message       [PacketChunkSize]int8 //col:43
}

