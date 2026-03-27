package hyperdbgrust

type BreakpointType uint32

const (
	BreakpointSoftware BreakpointType = 0
	BreakpointHardware BreakpointType = 1
	BreakpointHidden   BreakpointType = 2
	BreakpointEpt      BreakpointType = 3
)

type BreakpointEvent struct {
	Header       EventHeader   `json:"header"`
	Address      Address       `json:"address"`
	BreakpointID uint64        `json:"breakpoint_id"`
	Registers    RegisterState `json:"registers"`
}
