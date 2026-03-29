package debugger

import (
	"bytes"
	"encoding/binary"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/golibrary/std/mylog"
)

type EventLoop struct {
	mu            sync.Mutex
	running       bool
	eventManager  *EventManager
	eventHandlers map[EventType][]func(*DebugEvent)
	stopChan      chan struct{}
	driver        driver.Api
}

func NewEventLoop(eventManager *EventManager, driver driver.Api) *EventLoop {
	return &EventLoop{
		eventManager:  eventManager,
		eventHandlers: make(map[EventType][]func(*DebugEvent)),
		stopChan:      make(chan struct{}),
		driver:        driver,
	}
}

func (el *EventLoop) Start() {
	el.mu.Lock()
	defer el.mu.Unlock()

	if el.running {
		return
	}

	el.running = true
	mylog.Info("EventLoop: Starting")

	go el.run()
}

func (el *EventLoop) Stop() {
	el.mu.Lock()
	defer el.mu.Unlock()

	if !el.running {
		return
	}

	el.running = false
	mylog.Info("EventLoop: Stopping")
	close(el.stopChan)
}

func (el *EventLoop) run() {
	eventChan := el.eventManager.DebugEventChannel()

	go el.receiveEvents()

	for {
		select {
		case event := <-eventChan:
			el.handleEvent(event)
		case <-el.stopChan:
			mylog.Info("EventLoop: Stopped")
			return
		}
	}
}

func (el *EventLoop) receiveEvents() {
	for el.IsRunning() {
		if el.driver == nil || !el.driver.IsConnected() {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		response := el.driver.Receive(IoctlRegisterEvent)
		if response == nil || response.Len() == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		el.handleKernelEvent(response)
	}
}

func (el *EventLoop) handleKernelEvent(response *bytes.Buffer) {
	if response.Len() < 4 {
		return
	}

	operationCode := binary.LittleEndian.Uint32(response.Next(4))

	if (operationCode & 0x80000000) != 0 {
		el.handleMandatoryEvent(operationCode, response)
	} else {
		time.Sleep(10 * time.Millisecond)
	}
}

func (el *EventLoop) handleMandatoryEvent(operationCode uint32, response *bytes.Buffer) {
	switch operationCode {
	case 0x80000010: // OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE
		el.handleUserDebuggerPause(response)
	default:
		mylog.Info("Unknown mandatory event", operationCode)
	}
}

func (el *EventLoop) handleUserDebuggerPause(response *bytes.Buffer) {
	if response.Len() < 90 {
		return
	}

	event := &DebugEvent{
		Type:      EventBreakpoint,
		ProcessId: binary.LittleEndian.Uint32(response.Next(4)),
		Tag:       binary.LittleEndian.Uint64(response.Next(8)),
	}

	response.Next(4) // Skip ThreadId
	response.Next(8) // Skip Token
	event.EIP = binary.LittleEndian.Uint64(response.Next(8))

	response.Next(1)  // Skip Is32Bit
	response.Next(4)  // Skip PausingReason
	response.Next(8)  // Skip Rflags
	response.Next(8)  // Skip EventTag
	response.Next(8)  // Skip EventCallingStage
	response.Next(16) // Skip InstructionBytesOnRip
	response.Next(2)  // Skip ReadInstructionLen

	el.eventManager.TriggerDebugEvent(event)
}

func (el *EventLoop) handleEvent(event *DebugEvent) {
	mylog.Info(event.Type, event.Tag, event.ProcessId)

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
	mylog.Info(eventType.String())
}

func (el *EventLoop) UnregisterHandler(eventType EventType, handler func(*DebugEvent)) {
	el.mu.Lock()
	defer el.mu.Unlock()

	_, ok := el.eventHandlers[eventType]
	if !ok {
		return
	}

	delete(el.eventHandlers, eventType)
	mylog.Info(eventType)
}

func (el *EventLoop) IsRunning() bool {
	el.mu.Lock()
	defer el.mu.Unlock()
	return el.running
}
