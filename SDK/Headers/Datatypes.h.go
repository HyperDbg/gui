package Headers

type (
	CallBack func(Text *int8) int
)

type (
	DEBUGGEE_USER_INPUT_PACKET struct {
		CommandLen           uint32
		IgnoreFinishedSignal bool
		Result               uint32
		// The user's input is here
	}
	PDEBUGGEE_USER_INPUT_PACKET *DEBUGGEE_USER_INPUT_PACKET

	DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct {
		Length uint32
		// The buffer for event and action is here
	}
	PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET *DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
)

//todo add a const with a unit test
//SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED \
//    sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED)

type (
	DEBUGGER_PAUSE_PACKET_RECEIVED struct {
		Result uint32 // Result from kernel
	}
	PDEBUGGER_PAUSE_PACKET_RECEIVED *DEBUGGER_PAUSE_PACKET_RECEIVED
)

type (
	DEBUGGEE_KD_PAUSED_PACKET struct {
		Rip            uint64
		Is32BitAddress bool // if true shows that the address should be interpreted in 32-bit mode
		PausingReason  DEBUGGEE_PAUSING_REASON
		CurrentCore    uint32
		EventTag       uint64
		Rflags         uint64
		//InstructionBytesOnRip[MAXIMUM_INSTR_SIZE]byte       //todo
		ReadInstructionLen uint16
	}
	PDEBUGGEE_KD_PAUSED_PACKET *DEBUGGEE_KD_PAUSED_PACKET
)

type (
	DEBUGGEE_UD_PAUSED_PACKET struct {
		Rip                   uint64
		ProcessDebuggingToken uint64
		Is32Bit               bool // if true shows that the address should be interpreted in 32-bit mode
		PausingReason         DEBUGGEE_PAUSING_REASON
		ProcessId             uint32
		ThreadId              uint32
		EventTag              uint64
		Rflags                uint64
		//InstructionBytesOnRip[MAXIMUM_INSTR_SIZE]byte   //todo make fn get it ?
		ReadInstructionLen uint16
		GuestRegs          GUEST_REGS
	}
	PDEBUGGEE_UD_PAUSED_PACKET *DEBUGGEE_UD_PAUSED_PACKET
)

//todo
//static_assert(sizeof(DEBUGGEE_UD_PAUSED_PACKET) < PacketChunkSize,
//   "err (static_assert), size of PacketChunkSize should be bigger than DEBUGGEE_UD_PAUSED_PACKET");

type (
	DEBUGGEE_MESSAGE_PACKET struct {
		OperationCode uint32
		//Message[PacketChunkSize]   byte      //todo
	}
	PDEBUGGEE_MESSAGE_PACKET *PDEBUGGEE_MESSAGE_PACKET
)
