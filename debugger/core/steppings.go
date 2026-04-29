package core

type SteppingRequest uint32

const (
	StepInstrumentationStepIn            SteppingRequest = 0
	StepInstrumentationStepInForTracking SteppingRequest = 1
	StepRegularStepIn                    SteppingRequest = 2
	StepStepOver                         SteppingRequest = 3
	StepInstrumentationStepOut           SteppingRequest = 4
	StepRegularStepOut                   SteppingRequest = 5
	StepStepOverForGu                    SteppingRequest = 6
	StepStepOverForGuLastInstruction     SteppingRequest = 7
)

type DebuggerState struct {
	IsActive       bool
	IsPaused       bool
	ProcessId      uint32
	ThreadId       uint32
	Rip            uint64
	Is32Bit        bool
	DebuggingToken uint64
}

type Core struct {
	activeProcess     DebuggerState
	isSerialConnected bool
	sendStepPacket    func(SteppingRequest) bool
}

func NewCore(sendStepPacket func(SteppingRequest) bool) *Core {
	return &Core{
		sendStepPacket: sendStepPacket,
	}
}

func (c *Core) SetActiveProcess(state DebuggerState) {
	c.activeProcess = state
}

func (c *Core) GetActiveProcess() DebuggerState {
	return c.activeProcess
}

func (c *Core) InstrumentationStepIn() bool {
	if c.activeProcess.IsActive {
		return false
	}
	return c.sendStepPacket(StepInstrumentationStepIn)
}

func (c *Core) InstrumentationStepInForTracking() bool {
	if c.activeProcess.IsActive {
		return false
	}
	return c.sendStepPacket(StepInstrumentationStepInForTracking)
}

func (c *Core) RegularStepIn() bool {
	if c.isSerialConnected {
		return c.sendStepPacket(StepRegularStepIn)
	}
	if c.activeProcess.IsActive && c.activeProcess.IsPaused {
		return c.sendStepPacket(StepRegularStepIn)
	}
	return false
}

func (c *Core) StepOver() bool {
	if c.isSerialConnected {
		return c.sendStepPacket(StepStepOver)
	}
	if c.activeProcess.IsActive && c.activeProcess.IsPaused {
		return c.sendStepPacket(StepStepOver)
	}
	return false
}

func (c *Core) InstrumentationStepOut() bool {
	if c.activeProcess.IsActive {
		return false
	}
	return c.sendStepPacket(StepInstrumentationStepOut)
}

func (c *Core) RegularStepOut() bool {
	if c.isSerialConnected {
		return c.sendStepPacket(StepRegularStepOut)
	}
	if c.activeProcess.IsActive && c.activeProcess.IsPaused {
		return c.sendStepPacket(StepRegularStepOut)
	}
	return false
}

func (c *Core) StepOverForGu(lastInstruction bool) bool {
	var requestFormat SteppingRequest

	if !lastInstruction {
		requestFormat = StepStepOverForGu
	} else {
		requestFormat = StepStepOverForGuLastInstruction
	}

	if c.isSerialConnected {
		return c.sendStepPacket(requestFormat)
	} else if c.activeProcess.IsActive && c.activeProcess.IsPaused {
		return c.sendStepPacket(requestFormat)
	} else {
		return false
	}
}
