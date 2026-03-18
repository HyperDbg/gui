package commands

import (
	"log/slog"
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
	slog.Info("Commands: Triggering event", "type", event.Type, "pid", event.Pid)

	callbacks, ok := h.callbacks[event.Type]
	if !ok {
		return nil
	}

	for _, callback := range callbacks {
		if err := callback(event); err != nil {
			slog.Error("Commands: Callback error", "error", err)
			return err
		}
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

func (h *EventHandler) OnRemoveBreakpoint(callback EventCallback) {
	h.RegisterCallback("remove_breakpoint", callback)
}

func (h *EventHandler) OnSingleStep(callback EventCallback) {
	h.RegisterCallback("single_step", callback)
}

func (h *EventHandler) OnContinue(callback EventCallback) {
	h.RegisterCallback("continue", callback)
}

func (h *EventHandler) OnAttachProcess(callback EventCallback) {
	h.RegisterCallback("attach_process", callback)
}

func (h *EventHandler) OnDetachProcess(callback EventCallback) {
	h.RegisterCallback("detach_process", callback)
}
