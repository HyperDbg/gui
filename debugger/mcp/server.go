package mcp

import (
	"log/slog"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/memory"
)

type Server struct {
	debugger *debugger.HyperDbg
	driver   *driver.DriverHandle
	handlers map[EventType][]*Handler
}

func NewServer(debugger *debugger.HyperDbg, driverHandle *driver.DriverHandle) *Server {
	return &Server{
		debugger: debugger,
		driver:   driverHandle,
		handlers: make(map[EventType][]*Handler),
	}
}

func (s *Server) RegisterHandler(handler *Handler) {
	s.handlers[handler.eventType] = append(s.handlers[handler.eventType], handler)
}

func (s *Server) HandleEvent(event *Event) error {
	slog.Info("MCP: Handling event", "type", event.Type, "pid", event.Pid, "address", event.Address)

	handlers, ok := s.handlers[event.Type]
	if !ok {
		return nil
	}

	for _, handler := range handlers {
		if err := handler.Handle(event); err != nil {
			slog.Error("MCP: Handler error", "error", err)
			return err
		}
	}

	return nil
}

func (s *Server) ReadMemory(pid uint32, address uint64, size uint32) ([]byte, error) {
	event := &Event{
		Type:    EventReadMemory,
		Pid:     pid,
		Address: address,
		Size:    size,
	}

	if err := s.HandleEvent(event); err != nil {
		return nil, err
	}

	memProvider := s.debugger.GetMemory()
	if memProvider == nil {
		return nil, nil
	}

	return memProvider.Self().(*memory.Provider).ReadMemory(pid, address, size)
}

func (s *Server) WriteMemory(pid uint32, address uint64, data []byte) error {
	event := &Event{
		Type:    EventWriteMemory,
		Pid:     pid,
		Address: address,
		Size:    uint32(len(data)),
		Data:    data,
	}

	if err := s.HandleEvent(event); err != nil {
		return err
	}

	return nil
}
