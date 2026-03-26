package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg_rust/server"
)

type Debugger struct {
	server       *server.Server
	driverID     uint64
	state        protocol.DebugState
	mu           sync.RWMutex
	scriptEngine *ScriptEngine
}

func NewDebugger() *Debugger {
	dbg := &Debugger{
		server: server.NewServer(),
		state:  protocol.StateTerminated,
	}
	dbg.scriptEngine = NewScriptEngine(dbg)
	return dbg
}

func (d *Debugger) Start(port int) error {
	if err := d.server.Start(port); err != nil {
		return err
	}

	d.server.RegisterCallback(protocol.MsgTypeBreakpointEvent, d.handleBreakpointEvent)
	d.server.RegisterCallback(protocol.MsgTypeExceptionEvent, d.handleExceptionEvent)
	d.server.RegisterCallback(protocol.MsgTypeDebugPrintEvent, d.handleDebugPrintEvent)
	d.server.RegisterCallback(protocol.MsgTypeResponse, d.handleResponse)
	d.server.RegisterCallback(protocol.MsgTypeError, d.handleError)

	return nil
}

func (d *Debugger) Stop() {
	d.server.Stop()
}

func (d *Debugger) WaitForDriver(timeout time.Duration) error {
	start := time.Now()
	for time.Since(start) < timeout {
		drivers := d.server.GetConnectedDrivers()
		if len(drivers) > 0 {
			d.mu.Lock()
			d.driverID = drivers[0]
			d.mu.Unlock()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("timeout waiting for driver connection")
}

func (d *Debugger) Initialize() error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if err := d.server.InitializeDriver(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateRunning
	d.mu.Unlock()

	return nil
}

func (d *Debugger) Terminate() error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return nil
	}

	if err := d.server.TerminateDriver(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateTerminated
	d.mu.Unlock()

	return nil
}

func (d *Debugger) AttachProcess(processID uint32) error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.AttachProcess(driverID, processID)
}

func (d *Debugger) DetachProcess() error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.DetachProcess(driverID)
}

func (d *Debugger) SetBreakpoint(address uint64, bpType protocol.BreakpointType) error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.SetBreakpoint(driverID, address, bpType)
}

func (d *Debugger) RemoveBreakpoint(breakpointID uint64) error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.RemoveBreakpoint(driverID, breakpointID)
}

func (d *Debugger) Continue() error {
	d.mu.RLock()
	driverID := d.driverID
	state := d.state
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if state != protocol.StatePaused {
		return fmt.Errorf("debugger not paused")
	}

	if err := d.server.Continue(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateRunning
	d.mu.Unlock()

	return nil
}

func (d *Debugger) Pause() error {
	d.mu.RLock()
	driverID := d.driverID
	state := d.state
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if state != protocol.StateRunning {
		return fmt.Errorf("debugger not running")
	}

	if err := d.server.Pause(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StatePaused
	d.mu.Unlock()

	return nil
}

func (d *Debugger) StepInto() error {
	d.mu.RLock()
	driverID := d.driverID
	state := d.state
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if state != protocol.StatePaused {
		return fmt.Errorf("debugger not paused")
	}

	if err := d.server.StepInto(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateStepping
	d.mu.Unlock()

	return nil
}

func (d *Debugger) StepOver() error {
	d.mu.RLock()
	driverID := d.driverID
	state := d.state
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if state != protocol.StatePaused {
		return fmt.Errorf("debugger not paused")
	}

	if err := d.server.StepOver(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateStepping
	d.mu.Unlock()

	return nil
}

func (d *Debugger) StepOut() error {
	d.mu.RLock()
	driverID := d.driverID
	state := d.state
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	if state != protocol.StatePaused {
		return fmt.Errorf("debugger not paused")
	}

	if err := d.server.StepOut(driverID); err != nil {
		return err
	}

	d.mu.Lock()
	d.state = protocol.StateStepping
	d.mu.Unlock()

	return nil
}

func (d *Debugger) ReadMemory(address uint64, size uint32) ([]byte, error) {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return nil, fmt.Errorf("no driver connected")
	}

	if err := d.server.ReadMemory(driverID, address, size); err != nil {
		return nil, err
	}

	msg := d.server.WaitForEvent(5 * time.Second)
	if msg == nil || msg.Header.Type != protocol.MsgTypeResponse {
		return nil, fmt.Errorf("no response from driver")
	}

	return msg.Payload, nil
}

func (d *Debugger) WriteMemory(address uint64, data []byte) error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.WriteMemory(driverID, address, data)
}

func (d *Debugger) ReadRegisters() (*protocol.RegisterState, error) {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return nil, fmt.Errorf("no driver connected")
	}

	if err := d.server.ReadRegisters(driverID); err != nil {
		return nil, err
	}

	msg := d.server.WaitForEvent(5 * time.Second)
	if msg == nil || msg.Header.Type != protocol.MsgTypeResponse {
		return nil, fmt.Errorf("no response from driver")
	}

	return parseRegisterState(msg.Payload), nil
}

func (d *Debugger) WriteRegisters(regs *protocol.RegisterState) error {
	d.mu.RLock()
	driverID := d.driverID
	d.mu.RUnlock()

	if driverID == 0 {
		return fmt.Errorf("no driver connected")
	}

	return d.server.WriteRegisters(driverID, regs)
}

func (d *Debugger) GetState() protocol.DebugState {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.state
}

func (d *Debugger) WaitForEvent(timeout time.Duration) *protocol.Message {
	return d.server.WaitForEvent(timeout)
}

func (d *Debugger) handleBreakpointEvent(msg *protocol.Message) {
	d.mu.Lock()
	d.state = protocol.StatePaused
	d.mu.Unlock()

	event := parseBreakpointEvent(msg.Payload)
	fmt.Printf("[Breakpoint] PID=%d TID=%d Address=0x%x\n",
		event.ProcessID, event.ThreadID, event.Address)
}

func (d *Debugger) handleExceptionEvent(msg *protocol.Message) {
	d.mu.Lock()
	d.state = protocol.StatePaused
	d.mu.Unlock()

	event := parseExceptionEvent(msg.Payload)
	fmt.Printf("[Exception] PID=%d TID=%d Code=0x%x Address=0x%x\n",
		event.ProcessID, event.ThreadID, event.ExceptionCode, event.Address)
}

func (d *Debugger) handleDebugPrintEvent(msg *protocol.Message) {
	event := parseDebugPrintEvent(msg.Payload)
	fmt.Printf("[DebugPrint] %s\n", event.Message)
}

func (d *Debugger) handleResponse(msg *protocol.Message) {
}

func (d *Debugger) handleError(msg *protocol.Message) {
	errMsg := string(msg.Payload)
	fmt.Printf("[Error] %s\n", errMsg)
}

func parseRegisterState(data []byte) *protocol.RegisterState {
	regs := &protocol.RegisterState{}
	offset := 0
	regs.RAX = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RBX = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RCX = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RDX = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RSI = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RDI = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RBP = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RSP = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R8 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R9 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R10 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R11 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R12 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R13 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R14 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.R15 = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RIP = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	regs.RFLAGS = binary.LittleEndian.Uint64(data[offset:])
	return regs
}

func parseBreakpointEvent(data []byte) *protocol.BreakpointEvent {
	event := &protocol.BreakpointEvent{}
	offset := 0
	event.ProcessID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.ThreadID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.CoreID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.Address = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	event.BreakpointID = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	return event
}

func parseExceptionEvent(data []byte) *protocol.ExceptionEvent {
	event := &protocol.ExceptionEvent{}
	offset := 0
	event.ProcessID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.ThreadID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.CoreID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.ExceptionCode = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.Address = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	event.ErrorCode = binary.LittleEndian.Uint64(data[offset:])
	return event
}

func parseDebugPrintEvent(data []byte) *protocol.DebugPrintEvent {
	event := &protocol.DebugPrintEvent{}
	offset := 0
	event.ProcessID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.ThreadID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.CoreID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.Level = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	msgLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	event.Message = string(data[offset : offset+msgLen])
	return event
}

func (d *Debugger) ExecuteScript(script string) (string, error) {
	return d.scriptEngine.Execute(script)
}

func (d *Debugger) ExecuteScriptWithContext(ctx context.Context, script string) (string, error) {
	return d.scriptEngine.ExecuteWithContext(ctx, script)
}
