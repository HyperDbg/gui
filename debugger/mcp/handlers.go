package mcp

type EventHandler func(event *Event) error

type Handler struct {
	eventType EventType
	handler   EventHandler
}

func NewHandler(eventType EventType, handler EventHandler) *Handler {
	return &Handler{
		eventType: eventType,
		handler:   handler,
	}
}

func (h *Handler) Handle(event *Event) error {
	if event.Type != h.eventType {
		return nil
	}
	return h.handler(event)
}
