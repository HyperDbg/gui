package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h.back

type _DEBUGGEE_USER_INPUT_PACKET struct {
	CommandLen           uint32 //col:8
	IgnoreFinishedSignal bool   //col:9
	Result               uint32 //col:10
}

type _DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct {
	Length uint32 //col:12
}

type _DEBUGGER_PAUSE_PACKET_RECEIVED struct {
	Result uint32 //col:16
}

type _DEBUGGEE_KD_PAUSED_PACKET struct {
	Rip                   uint64                    //col:27
	Is32BitAddress        bool                      //col:28
	PausingReason         DEBUGGEE_PAUSING_REASON   //col:29
	CurrentCore           uint32                    //col:30
	EventTag              uint64                    //col:31
	Rflags                uint64                    //col:32
	InstructionBytesOnRip [MAXIMUM_INSTR_SIZE]uint8 //col:33
	ReadInstructionLen    uint16                    //col:34
}

type _DEBUGGEE_UD_PAUSED_PACKET struct {
	Rip                   uint64                    //col:41
	ProcessDebuggingToken uint64                    //col:42
	Is32Bit               bool                      //col:43
	PausingReason         DEBUGGEE_PAUSING_REASON   //col:44
	ProcessId             uint32                    //col:45
	ThreadId              uint32                    //col:46
	EventTag              uint64                    //col:47
	Rflags                uint64                    //col:48
	InstructionBytesOnRip [MAXIMUM_INSTR_SIZE]uint8 //col:49
	ReadInstructionLen    uint16                    //col:50
	GuestRegs             GUEST_REGS                //col:51
}

type _DEBUGGEE_MESSAGE_PACKET struct {
	OperationCode uint32                //col:46
	Message       [PacketChunkSize]int8 //col:47
}

