package debugger

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

type EventCallback func(event *Event) error

type EventManager struct {
	driver         *driver.Handle
	handlers       map[EventType][]EventCallback
	mu             sync.RWMutex
	events         map[uint64]*DebugEvent
	eventTag       uint64
	eventCallbacks []DebugEventCallback
	eventChan      chan *DebugEvent
}

const (
	DebuggerEventTagStartSeed                             = 0x1000
	DebuggerOutputSourceMaximumRemoteSourceForSingleEvent = 4
)

func NewEventManager(driverHandle *driver.Handle) *EventManager {
	return &EventManager{
		driver:         driverHandle,
		handlers:       make(map[EventType][]EventCallback),
		events:         make(map[uint64]*DebugEvent),
		eventTag:       DebuggerEventTagStartSeed,
		eventCallbacks: make([]DebugEventCallback, 0),
		eventChan:      make(chan *DebugEvent, 100),
	}
}

func (m *EventManager) RegisterHandler(eventType EventType, handler EventCallback) {
	m.handlers[eventType] = append(m.handlers[eventType], handler)
}

func (m *EventManager) TriggerEvent(event *Event) error {
	slog.Info("Event: Triggering event", "type", event.Type, "pid", event.Pid)

	handlers, ok := m.handlers[event.Type]
	if !ok {
		return nil
	}

	for _, handler := range handlers {
		if err := handler(event); err != nil {
			slog.Error("Event: Handler error", "error", err)
			return err
		}
	}

	return nil
}

func (m *EventManager) SendEvent(event *Event) error {
	return m.TriggerEvent(event)
}

func (m *EventManager) RegisterDebugEvent(event *DebugEvent) (uint64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	event.Tag = m.eventTag
	event.CreationTime = time.Now().Unix()
	event.IsEnabled = true
	m.events[event.Tag] = event
	m.eventTag++

	return event.Tag, nil
}

func (m *EventManager) ModifyDebugEvent(tag uint64, modifyType EventModifyType) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	event, exists := m.events[tag]
	if !exists {
		return false, fmt.Errorf("event with tag %x not found", tag)
	}

	switch modifyType {
	case ModifyEnable:
		event.IsEnabled = true
		return true, nil
	case ModifyDisable:
		event.IsEnabled = false
		return false, nil
	case ModifyClear:
		delete(m.events, tag)
		return true, nil
	case ModifyQueryState:
		return event.IsEnabled, nil
	}

	return false, nil
}

func (m *EventManager) ModifyAllDebugEvents(modifyType EventModifyType) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch modifyType {
	case ModifyEnable:
		for _, event := range m.events {
			event.IsEnabled = true
		}
	case ModifyDisable:
		for _, event := range m.events {
			event.IsEnabled = false
		}
	case ModifyClear:
		m.events = make(map[uint64]*DebugEvent)
		m.eventTag = DebuggerEventTagStartSeed
	}

	return nil
}

func (m *EventManager) QueryDebugEventState(tag uint64) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	event, exists := m.events[tag]
	if !exists {
		return false, fmt.Errorf("event with tag %x not found", tag)
	}

	return event.IsEnabled, nil
}

func (m *EventManager) GetDebugEvents() []*DebugEvent {
	m.mu.RLock()
	defer m.mu.RUnlock()

	events := make([]*DebugEvent, 0, len(m.events))
	for _, event := range m.events {
		events = append(events, event)
	}
	return events
}

func (m *EventManager) GetDebugEvent(tag uint64) (*DebugEvent, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	event, exists := m.events[tag]
	if !exists {
		return nil, fmt.Errorf("event with tag %x not found", tag)
	}

	return event, nil
}

func (m *EventManager) AddDebugEventCallback(callback DebugEventCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.eventCallbacks = append(m.eventCallbacks, callback)
}

func (m *EventManager) RemoveDebugEventCallback(callback DebugEventCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.eventCallbacks = m.eventCallbacks[:0]
}

func (m *EventManager) TriggerDebugEvent(event *DebugEvent) {
	m.mu.RLock()
	callbacks := make([]DebugEventCallback, len(m.eventCallbacks))
	copy(callbacks, m.eventCallbacks)
	m.mu.RUnlock()

	for _, callback := range callbacks {
		callback(event)
	}

	select {
	case m.eventChan <- event:
	default:
		fmt.Printf("Warning: event channel full, dropping event\n")
	}
}

func (m *EventManager) DebugEventChannel() <-chan *DebugEvent {
	return m.eventChan
}

func (m *EventManager) ClearAllDebugEvents() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.events = make(map[uint64]*DebugEvent)
	m.eventTag = DebuggerEventTagStartSeed
}

func (m *EventManager) IsDebugEventTagExist(tag uint64) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, exists := m.events[tag]
	return exists
}
