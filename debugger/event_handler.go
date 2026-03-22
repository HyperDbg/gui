package debugger

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/HyperDbg/debugger/thread"
	"github.com/ddkwork/golibrary/std/mylog"
)

type EventHandler struct {
	driver driver.Api
}

func NewEventHandler(driverHandle driver.Api) *EventHandler {
	return &EventHandler{
		driver: driverHandle,
	}
}

func (h *EventHandler) HandleSetBreakpoint(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructBreakpointBuffer(event)
	h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)

	mylog.Info(event.Pid, event.Address)
	return nil
}

func (h *EventHandler) HandleRemoveBreakpoint(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	eventBuffer := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&eventBuffer[0])) = event.Address

	h.driver.SendReceive(bytes.NewBuffer(eventBuffer), IoctlDebuggerModifyEvents)

	mylog.Info(event.Pid, event.Address)
	return nil
}

func (h *EventHandler) HandleSingleStep(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	req := DebuggerGeneralEventDetails{
		Tag:       0,
		ProcessId: 0,
		IsEnabled: false,
		Type:      uint32(event.Type),
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	h.driver.SendReceive(buf, 0x222004)

	mylog.Info(event.Pid)
	return nil
}

func (h *EventHandler) HandleContinue(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	h.driver.SendReceive(bytes.NewBuffer(nil), 0x222004)

	mylog.Info(event.Pid)
	return nil
}

func (h *EventHandler) constructBreakpointBuffer(event *Event) []byte {
	req := DebuggerEventRegisterRequest{
		EventType:      uint32(event.Type),
		Tag:            0,
		ProcessId:      event.Pid,
		CoreId:         0,
		IsEnabled:      false,
		EventStage:     0,
		OptionalParam1: event.Address,
		OptionalParam2: 0,
		OptionalParam3: 0,
		OptionalParam4: 0,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	return buf.Bytes()
}

type DebugEventHandler struct {
	driver    driver.Api
	regMgr    *register.Provider
	threadMgr *thread.Provider
}

func NewDebugEventHandler(driverHandle driver.Api, regMgr *register.Provider, threadMgr *thread.Provider) *DebugEventHandler {
	return &DebugEventHandler{
		driver:    driverHandle,
		regMgr:    regMgr,
		threadMgr: threadMgr,
	}
}

func (h *DebugEventHandler) RegisterDebugEvent(event *DebugEvent) (uint64, error) {
	if h.driver == nil || !h.driver.IsConnected() {
		return 0, fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	if response.Len() < 8 {
		return 0, fmt.Errorf("invalid response length")
	}

	tag := binary.LittleEndian.Uint64(response.Bytes()[0:8])
	return tag, nil
}

func (h *DebugEventHandler) constructDebugEventBuffer(event *DebugEvent) []byte {
	req := DebuggerEventRegisterRequest{
		EventType:      uint32(event.Type),
		Tag:            0,
		ProcessId:      event.ProcessId,
		CoreId:         event.CoreId,
		IsEnabled:      false,
		EventStage:     0,
		OptionalParam1: event.EIP,
		OptionalParam2: event.Options.OptionalParam1,
		OptionalParam3: event.Options.OptionalParam2,
		OptionalParam4: event.Options.OptionalParam3,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	return buf.Bytes()
}

func (h *DebugEventHandler) HandleEptHook(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEptHookResponse(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleSyscallHook(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleCpuid(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleRdmsr(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleWrmsr(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleIn(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleOut(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleException(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleInterrupt(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleDebugRegisters(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleTsc(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandlePmc(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleVmcall(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleXsetbv(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleControlRegister(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleCr3(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleTrap(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) HandleStep(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	response := h.driver.SendReceive(bytes.NewBuffer(buffer), IoctlDebuggerRegisterEvent)
	h.decodeEventActionResult(response.Bytes())
	return nil
}

func (h *DebugEventHandler) decodeEptHookResponse(data []byte) {
	if len(data) < 8 {
		return
	}
}

func (h *DebugEventHandler) decodeEventActionResult(data []byte) {
	if len(data) < 4 {
		return
	}
}
