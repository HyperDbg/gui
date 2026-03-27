package debugger

type BreakpointType uint32

const (
	BreakpointSoftware BreakpointType = iota
	BreakpointHardware
	BreakpointHidden
	BreakpointEpt
)

type BreakpointEvent struct {
	Header       EventHeader   `json:"header"`
	Address      Address       `json:"address"`
	BreakpointID uint64        `json:"breakpoint_id"`
	Registers    RegisterState `json:"registers"`
}
