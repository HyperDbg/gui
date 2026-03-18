package mcp

type EventType int

const (
	EventReadMemory EventType = iota
	EventWriteMemory
	EventSetBreakpoint
	EventRemoveBreakpoint
	EventSingleStep
	EventContinue
	EventAttachProcess
	EventDetachProcess
)

type Event struct {
	Type     EventType
	Pid      uint32
	Address  uint64
	Size     uint32
	Data     []byte
	Metadata map[string]any
}
