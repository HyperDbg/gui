package commands

import (
	"github.com/ddkwork/golibrary/std/mylog"
)

type Event struct {
	Type     string
	Pid      uint32
	Address  uint64
	Size     uint32
	Data     []byte
	Metadata map[string]any
}

type EventCallback func(event *Event) error

type EventHandler struct {
	callbacks map[string][]EventCallback
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		callbacks: make(map[string][]EventCallback),
	}
}

func (h *EventHandler) RegisterCallback(eventType string, callback EventCallback) {
	h.callbacks[eventType] = append(h.callbacks[eventType], callback)
}

func (h *EventHandler) TriggerEvent(event *Event) error {
	mylog.Info(event.Type, event.Pid)

	callbacks, ok := h.callbacks[event.Type]
	if !ok {
		return nil
	}

	for _, callback := range callbacks {
		mylog.Check(callback(event))
	}

	return nil
}

func (h *EventHandler) OnReadMemory(callback EventCallback) {
	h.RegisterCallback("read_memory", callback)
}

func (h *EventHandler) OnWriteMemory(callback EventCallback) {
	h.RegisterCallback("write_memory", callback)
}

func (h *EventHandler) OnSetBreakpoint(callback EventCallback) {
	h.RegisterCallback("set_breakpoint", callback)
}
