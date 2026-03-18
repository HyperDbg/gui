package debugger

import (
	"log/slog"
	"sync"
)

type EventLoop struct {
	mu            sync.Mutex
	running       bool
	eventManager  *EventManager
	eventHandlers map[EventType][]func(*DebugEvent)
	stopChan      chan struct{}
}

func NewEventLoop(eventManager *EventManager) *EventLoop {
	return &EventLoop{
		eventManager:  eventManager,
		eventHandlers: make(map[EventType][]func(*DebugEvent)),
		stopChan:      make(chan struct{}),
	}
}

func (el *EventLoop) Start() {
	el.mu.Lock()
	defer el.mu.Unlock()

	if el.running {
		return
	}

	el.running = true
	slog.Info("EventLoop: Starting")

	go el.run()
}

func (el *EventLoop) Stop() {
	el.mu.Lock()
	defer el.mu.Unlock()

	if !el.running {
		return
	}

	el.running = false
	slog.Info("EventLoop: Stopping")
	close(el.stopChan)
}

func (el *EventLoop) run() {
	eventChan := el.eventManager.DebugEventChannel()

	for {
		select {
		case event := <-eventChan:
			el.handleEvent(event)
		case <-el.stopChan:
			slog.Info("EventLoop: Stopped")
			return
		}
	}
}

func (el *EventLoop) handleEvent(event *DebugEvent) {
	slog.Info("EventLoop: Handling event", "type", event.Type, "tag", event.Tag, "pid", event.ProcessId)

	el.mu.Lock()
	handlers, ok := el.eventHandlers[event.Type]
	el.mu.Unlock()

	if !ok {
		return
	}

	for _, handler := range handlers {
		handler(event)
	}
}

func (el *EventLoop) RegisterHandler(eventType EventType, handler func(*DebugEvent)) {
	el.mu.Lock()
	defer el.mu.Unlock()

	el.eventHandlers[eventType] = append(el.eventHandlers[eventType], handler)
	slog.Info("EventLoop: Handler registered", "event", eventType.String())
}

func (el *EventLoop) UnregisterHandler(eventType EventType, handler func(*DebugEvent)) {
	el.mu.Lock()
	defer el.mu.Unlock()

	_, ok := el.eventHandlers[eventType]
	if !ok {
		return
	}

	delete(el.eventHandlers, eventType)
	slog.Info("EventLoop: All handlers unregistered for type", "type", eventType)
}

func (el *EventLoop) IsRunning() bool {
	el.mu.Lock()
	defer el.mu.Unlock()
	return el.running
}
