package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	DriverHTTPHost = "127.0.0.1"
	DriverHTTPPort = 50080
	HTTPTimeout    = 10 * time.Second
)

type HTTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Command struct {
	Action string `json:"action"`
	Target string `json:"target,omitempty"`
	Size   int    `json:"size,omitempty"`
	Data   string `json:"data,omitempty"`
}

type EventCallback func(event interface{})

type Server struct {
	client      *http.Client
	baseURL     string
	driverID    atomic.Uint64
	running     atomic.Bool
	callbacks   map[string][]EventCallback
	callbacksMu sync.RWMutex
	eventChan   chan interface{}
	connected   atomic.Bool
}

func NewServer() *Server {
	s := &Server{
		client: &http.Client{
			Timeout: HTTPTimeout,
		},
		baseURL:   fmt.Sprintf("http://%s:%d", DriverHTTPHost, DriverHTTPPort),
		callbacks: make(map[string][]EventCallback),
		eventChan: make(chan interface{}, 1000),
	}
	s.driverID.Store(1)
	return s
}

func (s *Server) Start(port int) error {
	s.running.Store(true)

	if err := s.Ping(); err != nil {
		return fmt.Errorf("failed to connect to driver: %w", err)
	}

	s.connected.Store(true)
	fmt.Printf("Connected to driver at %s\n", s.baseURL)

	go s.eventPoller()

	return nil
}

func (s *Server) Stop() {
	s.running.Store(false)
	s.connected.Store(false)
}

func (s *Server) sendRequest(method, path string, body []byte) (HTTPResponse, error) {
	var resp HTTPResponse
	url := fmt.Sprintf("%s%s", s.baseURL, path)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return resp, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", DriverHTTPHost)

	response, err := s.client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("send failed: %w", err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return resp, fmt.Errorf("parse failed: %w", err)
	}

	return resp, nil
}

func (s *Server) sendCommand(cmd Command) (HTTPResponse, error) {
	data, err := json.Marshal(cmd)
	if err != nil {
		return HTTPResponse{}, fmt.Errorf("marshal failed: %w", err)
	}
	return s.sendRequest("POST", "/api", data)
}

func (s *Server) eventPoller() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for s.running.Load() {
		<-ticker.C
		if !s.connected.Load() {
			continue
		}
	}
}

func (s *Server) Ping() error {
	resp, err := s.sendRequest("GET", "/ping", nil)
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("ping failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) Status() (string, error) {
	resp, err := s.sendRequest("GET", "/status", nil)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (s *Server) RegisterCallback(eventType string, cb EventCallback) {
	s.callbacksMu.Lock()
	defer s.callbacksMu.Unlock()
	s.callbacks[eventType] = append(s.callbacks[eventType], cb)
}

func (s *Server) GetEvent() interface{} {
	select {
	case event := <-s.eventChan:
		return event
	default:
		return nil
	}
}

func (s *Server) WaitForEvent(timeout time.Duration) interface{} {
	select {
	case event := <-s.eventChan:
		return event
	case <-time.After(timeout):
		return nil
	}
}

func (s *Server) GetConnectedDrivers() []uint64 {
	if s.connected.Load() {
		return []uint64{s.driverID.Load()}
	}
	return nil
}

func (s *Server) GetDriverCount() int {
	if s.connected.Load() {
		return 1
	}
	return 0
}

func (s *Server) InitializeDriver(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "initialize"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("initialize failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) TerminateDriver(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "terminate"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("terminate failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) AttachProcess(driverID uint64, processID uint32) error {
	resp, err := s.sendCommand(Command{
		Action: "attach_process",
		Target: fmt.Sprintf("0x%x", processID),
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("attach process failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) DetachProcess(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "detach_process"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("detach process failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) SetBreakpoint(driverID uint64, address uint64, bpType uint32) error {
	resp, err := s.sendCommand(Command{
		Action: "set_breakpoint",
		Target: fmt.Sprintf("0x%x", address),
		Size:   int(bpType),
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("set breakpoint failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) RemoveBreakpoint(driverID uint64, breakpointID uint64) error {
	resp, err := s.sendCommand(Command{
		Action: "remove_breakpoint",
		Target: fmt.Sprintf("0x%x", breakpointID),
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("remove breakpoint failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) Continue(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "continue"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("continue failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) Pause(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "pause"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("pause failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) StepInto(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "step_into"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("step into failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) StepOver(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "step_over"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("step over failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) StepOut(driverID uint64) error {
	resp, err := s.sendCommand(Command{Action: "step_out"})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("step out failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) ReadMemory(driverID uint64, address uint64, size uint32) (string, error) {
	resp, err := s.sendCommand(Command{
		Action: "read_memory",
		Target: fmt.Sprintf("0x%x", address),
		Size:   int(size),
	})
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", fmt.Errorf("read memory failed: %s", resp.Message)
	}
	return resp.Message, nil
}

func (s *Server) WriteMemory(driverID uint64, address uint64, data string) error {
	resp, err := s.sendCommand(Command{
		Action: "write_memory",
		Target: fmt.Sprintf("0x%x", address),
		Data:   data,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("write memory failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) ReadRegisters(driverID uint64) (string, error) {
	resp, err := s.sendCommand(Command{Action: "read_registers"})
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", fmt.Errorf("read registers failed: %s", resp.Message)
	}
	return resp.Message, nil
}

func (s *Server) WriteRegisters(driverID uint64, regs string) error {
	resp, err := s.sendCommand(Command{
		Action: "write_registers",
		Data:   regs,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("write registers failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) Broadcast(action string) error {
	resp, err := s.sendCommand(Command{Action: action})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("broadcast failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) SendToDriver(driverID uint64, action string) error {
	resp, err := s.sendCommand(Command{Action: action})
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("send to driver failed: %s", resp.Message)
	}
	return nil
}

func (s *Server) IsConnected() bool {
	return s.connected.Load()
}
