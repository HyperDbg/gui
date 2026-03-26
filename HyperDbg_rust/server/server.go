package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"hyperdbg-communicator/protocol"
)

type DriverConnection struct {
	conn         net.Conn
	driverID     uint64
	sendQueue    chan *protocol.Message
	recvQueue    chan *protocol.Message
	closeChan    chan struct{}
	closed       atomic.Bool
	lastActivity time.Time
}

type EventCallback func(event *protocol.Message)

type Server struct {
	listener     net.Listener
	drivers      map[uint64]*DriverConnection
	driversMu    sync.RWMutex
	nextDriverID atomic.Uint64
	eventChan    chan *protocol.Message
	callbacks    map[protocol.MessageType][]EventCallback
	callbacksMu  sync.RWMutex
	running      atomic.Bool
}

func NewServer() *Server {
	return &Server{
		drivers:   make(map[uint64]*DriverConnection),
		eventChan: make(chan *protocol.Message, 1000),
		callbacks: make(map[protocol.MessageType][]EventCallback),
	}
}

func (s *Server) Start(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	s.listener = listener
	s.running.Store(true)

	go s.acceptLoop()
	go s.eventDispatcher()

	return nil
}

func (s *Server) Stop() {
	s.running.Store(false)
	if s.listener != nil {
		s.listener.Close()
	}

	s.driversMu.Lock()
	for _, driver := range s.drivers {
		driver.Close()
	}
	s.drivers = make(map[uint64]*DriverConnection)
	s.driversMu.Unlock()
}

func (s *Server) acceptLoop() {
	for s.running.Load() {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.running.Load() {
				fmt.Printf("Accept error: %v\n", err)
			}
			continue
		}

		driver := s.newDriverConnection(conn)
		go driver.recvLoop()
		go driver.sendLoop()
	}
}

func (s *Server) newDriverConnection(conn net.Conn) *DriverConnection {
	driverID := s.nextDriverID.Add(1)
	driver := &DriverConnection{
		conn:         conn,
		driverID:     driverID,
		sendQueue:    make(chan *protocol.Message, 100),
		recvQueue:    make(chan *protocol.Message, 100),
		closeChan:    make(chan struct{}),
		lastActivity: time.Now(),
	}

	s.driversMu.Lock()
	s.drivers[driverID] = driver
	s.driversMu.Unlock()

	fmt.Printf("Driver %d connected from %s\n", driverID, conn.RemoteAddr())
	return driver
}

func (d *DriverConnection) recvLoop() {
	defer d.Close()

	for {
		msg, err := protocol.ReadMessage(d.conn)
		if err != nil {
			if err != io.EOF && d.closed.Load() == false {
				fmt.Printf("Driver %d recv error: %v\n", d.driverID, err)
			}
			return
		}

		d.lastActivity = time.Now()
		select {
		case d.recvQueue <- msg:
		default:
			fmt.Printf("Driver %d recv queue full, dropping message\n", d.driverID)
		}
	}
}

func (d *DriverConnection) sendLoop() {
	defer d.Close()

	for {
		select {
		case msg := <-d.sendQueue:
			if err := protocol.WriteMessage(d.conn, msg); err != nil {
				fmt.Printf("Driver %d send error: %v\n", d.driverID, err)
				return
			}
			d.lastActivity = time.Now()
		case <-d.closeChan:
			return
		}
	}
}

func (d *DriverConnection) Close() {
	if d.closed.Swap(true) {
		return
	}
	close(d.closeChan)
	d.conn.Close()
}

func (d *DriverConnection) Send(msg *protocol.Message) error {
	if d.closed.Load() {
		return fmt.Errorf("connection closed")
	}
	select {
	case d.sendQueue <- msg:
		return nil
	default:
		return fmt.Errorf("send queue full")
	}
}

func (d *DriverConnection) Recv() (*protocol.Message, error) {
	select {
	case msg := <-d.recvQueue:
		return msg, nil
	case <-d.closeChan:
		return nil, fmt.Errorf("connection closed")
	}
}

func (s *Server) eventDispatcher() {
	for s.running.Load() {
		s.driversMu.RLock()
		for _, driver := range s.drivers {
			select {
			case msg := <-driver.recvQueue:
				s.handleMessage(driver, msg)
			default:
			}
		}
		s.driversMu.RUnlock()
		time.Sleep(time.Millisecond)
	}
}

func (s *Server) handleMessage(driver *DriverConnection, msg *protocol.Message) {
	s.callbacksMu.RLock()
	callbacks, ok := s.callbacks[msg.Header.Type]
	s.callbacksMu.RUnlock()

	if ok {
		for _, cb := range callbacks {
			go cb(msg)
		}
	}

	select {
	case s.eventChan <- msg:
	default:
		fmt.Printf("Event channel full, dropping event type %d\n", msg.Header.Type)
	}
}

func (s *Server) RegisterCallback(msgType protocol.MessageType, cb EventCallback) {
	s.callbacksMu.Lock()
	defer s.callbacksMu.Unlock()
	s.callbacks[msgType] = append(s.callbacks[msgType], cb)
}

func (s *Server) GetEvent() *protocol.Message {
	select {
	case msg := <-s.eventChan:
		return msg
	default:
		return nil
	}
}

func (s *Server) WaitForEvent(timeout time.Duration) *protocol.Message {
	select {
	case msg := <-s.eventChan:
		return msg
	case <-time.After(timeout):
		return nil
	}
}

func (s *Server) Broadcast(msg *protocol.Message) error {
	s.driversMu.RLock()
	defer s.driversMu.RUnlock()

	var lastErr error
	for _, driver := range s.drivers {
		if err := driver.Send(msg); err != nil {
			lastErr = err
		}
	}
	return lastErr
}

func (s *Server) SendToDriver(driverID uint64, msg *protocol.Message) error {
	s.driversMu.RLock()
	driver, ok := s.drivers[driverID]
	s.driversMu.RUnlock()

	if !ok {
		return fmt.Errorf("driver %d not found", driverID)
	}
	return driver.Send(msg)
}

func (s *Server) GetConnectedDrivers() []uint64 {
	s.driversMu.RLock()
	defer s.driversMu.RUnlock()

	ids := make([]uint64, 0, len(s.drivers))
	for id := range s.drivers {
		ids = append(ids, id)
	}
	return ids
}

func (s *Server) GetDriverCount() int {
	s.driversMu.RLock()
	defer s.driversMu.RUnlock()
	return len(s.drivers)
}

func (s *Server) InitializeDriver(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeInitialize, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) TerminateDriver(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeTerminate, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) AttachProcess(driverID uint64, processID uint32) error {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint32(payload, processID)
	msg := protocol.NewMessage(protocol.MsgTypeAttachProcess, payload)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) DetachProcess(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeDetachProcess, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) SetBreakpoint(driverID uint64, address uint64, bpType protocol.BreakpointType) error {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint64(payload[0:8], address)
	binary.LittleEndian.PutUint32(payload[8:12], uint32(bpType))
	msg := protocol.NewMessage(protocol.MsgTypeSetBreakpoint, payload)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) RemoveBreakpoint(driverID uint64, breakpointID uint64) error {
	payload := make([]byte, 8)
	binary.LittleEndian.PutUint64(payload, breakpointID)
	msg := protocol.NewMessage(protocol.MsgTypeRemoveBreakpoint, payload)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) Continue(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeContinue, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) Pause(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypePause, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) StepInto(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeStepInto, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) StepOver(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeStepOver, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) StepOut(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeStepOut, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) ReadMemory(driverID uint64, address uint64, size uint32) error {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint64(payload[0:8], address)
	binary.LittleEndian.PutUint32(payload[8:12], size)
	msg := protocol.NewMessage(protocol.MsgTypeReadMemory, payload)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) WriteMemory(driverID uint64, address uint64, data []byte) error {
	payload := make([]byte, 8+len(data))
	binary.LittleEndian.PutUint64(payload[0:8], address)
	copy(payload[8:], data)
	msg := protocol.NewMessage(protocol.MsgTypeWriteMemory, payload)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) ReadRegisters(driverID uint64) error {
	msg := protocol.NewMessage(protocol.MsgTypeReadRegisters, nil)
	return s.SendToDriver(driverID, msg)
}

func (s *Server) WriteRegisters(driverID uint64, regs *protocol.RegisterState) error {
	payload := make([]byte, 272)
	offset := 0
	binary.LittleEndian.PutUint64(payload[offset:], regs.RAX); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RBX); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RCX); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RDX); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RSI); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RDI); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RBP); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RSP); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R8); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R9); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R10); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R11); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R12); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R13); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R14); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.R15); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RIP); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.RFLAGS); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.CR0); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.CR2); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.CR3); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.CR4); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR0); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR1); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR2); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR3); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR6); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.DR7); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.GDTR); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.GSBase); offset += 8
	binary.LittleEndian.PutUint64(payload[offset:], regs.FSBase)
	msg := protocol.NewMessage(protocol.MsgTypeWriteRegisters, payload)
	return s.SendToDriver(driverID, msg)
}
