package debugger

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/process"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/HyperDbg/debugger/thread"
)

type EventHandler struct {
	driver *driver.Handle
}

func NewEventHandler(driverHandle *driver.Handle) *EventHandler {
	return &EventHandler{
		driver: driverHandle,
	}
}

func (h *EventHandler) HandleSetBreakpoint(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructBreakpointBuffer(event)
	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		slog.Error("Handler: Set breakpoint failed", "error", err)
		return err
	}

	slog.Info("Handler: Set breakpoint success", "pid", event.Pid, "address", event.Address)
	return nil
}

func (h *EventHandler) HandleRemoveBreakpoint(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	eventBuffer := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&eventBuffer[0])) = event.Address

	err := h.driver.SendBuffer(eventBuffer, IoctlUnregisterEvent)
	if err != nil {
		slog.Error("Handler: Remove breakpoint failed", "error", err)
		return err
	}

	slog.Info("Handler: Remove breakpoint success", "pid", event.Pid, "address", event.Address)
	return nil
}

func (h *EventHandler) HandleSingleStep(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = uint64(event.Type)
	*(*uint64)(unsafe.Pointer(&buffer[8])) = 0

	err := h.driver.SendBuffer(buffer, 0x222004)
	if err != nil {
		slog.Error("Handler: Single step failed", "error", err)
		return err
	}

	slog.Info("Handler: Single step success", "pid", event.Pid)
	return nil
}

func (h *EventHandler) HandleContinue(event *Event) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	h.driver.SendBuffer(nil, 0x222004)

	slog.Info("Handler: Continue", "pid", event.Pid)
	return nil
}

func (h *EventHandler) constructBreakpointBuffer(event *Event) []byte {
	buffer := make([]byte, 64)

	binary.LittleEndian.PutUint64(buffer[0:8], uint64(event.Type))
	binary.LittleEndian.PutUint32(buffer[8:12], event.Pid)
	binary.LittleEndian.PutUint32(buffer[12:16], 0)
	binary.LittleEndian.PutUint32(buffer[16:20], 0)
	binary.LittleEndian.PutUint64(buffer[20:28], event.Address)
	binary.LittleEndian.PutUint64(buffer[28:36], 0)
	binary.LittleEndian.PutUint64(buffer[36:44], 0)
	binary.LittleEndian.PutUint64(buffer[44:52], 0)

	return buffer
}

type DebugEventHandler struct {
	driver    *driver.Handle
	regMgr    *register.Provider
	procMgr   *process.Provider
	threadMgr *thread.Provider
}

func NewDebugEventHandler(driverHandle *driver.Handle, regMgr *register.Provider, procMgr *process.Provider, threadMgr *thread.Provider) *DebugEventHandler {
	return &DebugEventHandler{
		driver:    driverHandle,
		regMgr:    regMgr,
		procMgr:   procMgr,
		threadMgr: threadMgr,
	}
}

func (h *DebugEventHandler) RegisterDebugEvent(event *DebugEvent) (uint64, error) {
	if h.driver == nil || !h.driver.IsConnected() {
		return 0, fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return 0, fmt.Errorf("failed to send event registration: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return 0, fmt.Errorf("failed to receive event registration response: %v", err)
	}

	tag := binary.LittleEndian.Uint64(response[0:8])
	return tag, nil
}

func (h *DebugEventHandler) constructDebugEventBuffer(event *DebugEvent) []byte {
	buffer := make([]byte, 64)

	binary.LittleEndian.PutUint64(buffer[0:8], uint64(IoctlDebuggerRegisterEvent))
	binary.LittleEndian.PutUint64(buffer[8:16], uint64(event.Type))
	binary.LittleEndian.PutUint32(buffer[16:20], event.ProcessId)
	binary.LittleEndian.PutUint32(buffer[20:24], event.ThreadId)
	binary.LittleEndian.PutUint32(buffer[24:28], event.CoreId)
	binary.LittleEndian.PutUint64(buffer[28:36], event.EIP)
	binary.LittleEndian.PutUint64(buffer[36:44], event.Options.OptionalParam1)
	binary.LittleEndian.PutUint64(buffer[44:52], event.Options.OptionalParam2)
	binary.LittleEndian.PutUint64(buffer[52:60], event.Options.OptionalParam3)

	return buffer
}

func (h *DebugEventHandler) HandleEptHook(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleSyscallHook(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send syscall hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive syscall hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleCpuid(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send CPUID hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive CPUID hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleRdmsr(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send RDMSR hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive RDMSR hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleWrmsr(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send WRMSR hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive WRMSR hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleIn(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send IN hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive IN hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleOut(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send OUT hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive OUT hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleException(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send exception hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive exception hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleInterrupt(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send interrupt hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive interrupt hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleDebugRegisters(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send debug registers hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive debug registers hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleTsc(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send TSC hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive TSC hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandlePmc(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send PMC hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive PMC hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleVmcall(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send VMCALL hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive VMCALL hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleXsetbv(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send XSETBV hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive XSETBV hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleControlRegister(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send control register hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive control register hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleCr3(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send CR3 hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive CR3 hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleTrap(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send trap hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive trap hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleStep(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send step hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive step hook response: %v", err)
	}

	h.decodeStepResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeStepResponse(buffer []byte) {
	if len(buffer) < 48 {
		return
	}

	pausedPacket := &DebuggeeKdPausedPacket{
		Rip:                    binary.LittleEndian.Uint64(buffer[0:8]),
		IsProcessorOn32BitMode: buffer[8] != 0,
		IgnoreDisassembling:    buffer[9] != 0,
		PausingReason:          binary.LittleEndian.Uint32(buffer[12:16]),
		CurrentCore:            binary.LittleEndian.Uint32(buffer[16:20]),
		EventTag:               binary.LittleEndian.Uint64(buffer[20:28]),
		EventCallingStage:      binary.LittleEndian.Uint32(buffer[28:32]),
		Rflags:                 binary.LittleEndian.Uint64(buffer[32:40]),
	}

	if len(buffer) >= 48 {
		copy(pausedPacket.InstructionBytesOnRip[:], buffer[40:55])
		pausedPacket.ReadInstructionLen = binary.LittleEndian.Uint16(buffer[55:57])
	}

	if h.regMgr != nil {
		regCtx := h.regMgr.GetRegisterContext()
		regCtx.RIP = pausedPacket.Rip
		regCtx.RFLAGS = uint32(pausedPacket.Rflags)
		h.regMgr.SetRegisterContext(regCtx)
	}
}

func (h *DebugEventHandler) HandleBreakpoint(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send breakpoint hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive breakpoint hook response: %v", err)
	}

	h.decodeBreakpointResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeBreakpointResponse(buffer []byte) {
	if len(buffer) < 48 {
		return
	}

	pausedPacket := &DebuggeeKdPausedPacket{
		Rip:                    binary.LittleEndian.Uint64(buffer[0:8]),
		IsProcessorOn32BitMode: buffer[8] != 0,
		IgnoreDisassembling:    buffer[9] != 0,
		PausingReason:          binary.LittleEndian.Uint32(buffer[12:16]),
		CurrentCore:            binary.LittleEndian.Uint32(buffer[16:20]),
		EventTag:               binary.LittleEndian.Uint64(buffer[20:28]),
		EventCallingStage:      binary.LittleEndian.Uint32(buffer[28:32]),
		Rflags:                 binary.LittleEndian.Uint64(buffer[32:40]),
	}

	if len(buffer) >= 48 {
		copy(pausedPacket.InstructionBytesOnRip[:], buffer[40:55])
		pausedPacket.ReadInstructionLen = binary.LittleEndian.Uint16(buffer[55:57])
	}

	if h.regMgr != nil {
		regCtx := h.regMgr.GetRegisterContext()
		regCtx.RIP = pausedPacket.Rip
		regCtx.RFLAGS = uint32(pausedPacket.Rflags)
		h.regMgr.SetRegisterContext(regCtx)
	}
}

func (h *DebugEventHandler) HandleProcessCreated(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send process created hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive process created hook response: %v", err)
	}

	h.decodeProcessCreatedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeProcessCreatedResponse(buffer []byte) {
	if len(buffer) < 32 {
		return
	}

	processPacket := &DebuggeeDetailsAndSwitchProcessPacket{
		ActionType:            binary.LittleEndian.Uint32(buffer[0:4]),
		ProcessId:             binary.LittleEndian.Uint32(buffer[4:8]),
		Process:               binary.LittleEndian.Uint64(buffer[8:16]),
		IsSwitchByClkIntr:     buffer[16] != 0,
		ProcessListSymDetails: binary.LittleEndian.Uint64(buffer[20:28]),
		Result:                binary.LittleEndian.Uint32(buffer[28:32]),
	}

	if len(buffer) >= 36 {
		copy(processPacket.ProcessName[:], buffer[17:33])
	}

	if h.procMgr != nil {
		procName := string(processPacket.ProcessName[:])
		end := 0
		for i, b := range processPacket.ProcessName[:] {
			if b == 0 {
				end = i
				break
			}
		}
		if end > 0 {
			procName = string(processPacket.ProcessName[:end])
		}

		procInfo := &process.ProcessInfo{
			PID:  processPacket.ProcessId,
			Name: procName,
			Path: fmt.Sprintf("0x%x", processPacket.Process),
		}
		h.procMgr.AddProcess(procInfo)
	}
}

func (h *DebugEventHandler) HandleProcessExited(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send process exited hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive process exited hook response: %v", err)
	}

	h.decodeProcessExitedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeProcessExitedResponse(buffer []byte) {
	if len(buffer) < 32 {
		return
	}

	processPacket := &DebuggeeDetailsAndSwitchProcessPacket{
		ActionType:            binary.LittleEndian.Uint32(buffer[0:4]),
		ProcessId:             binary.LittleEndian.Uint32(buffer[4:8]),
		Process:               binary.LittleEndian.Uint64(buffer[8:16]),
		IsSwitchByClkIntr:     buffer[16] != 0,
		ProcessListSymDetails: binary.LittleEndian.Uint64(buffer[20:28]),
		Result:                binary.LittleEndian.Uint32(buffer[28:32]),
	}

	if len(buffer) >= 36 {
		copy(processPacket.ProcessName[:], buffer[17:33])
	}

	if h.procMgr != nil {
		h.procMgr.RemoveProcess(processPacket.ProcessId)
	}
}

func (h *DebugEventHandler) HandleThreadCreated(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send thread created hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive thread created hook response: %v", err)
	}

	h.decodeThreadCreatedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeThreadCreatedResponse(buffer []byte) {
	if len(buffer) < 40 {
		return
	}

	threadPacket := &DebuggeeDetailsAndSwitchThreadPacket{
		ActionType: binary.LittleEndian.Uint32(buffer[0:4]),
		ThreadId:   binary.LittleEndian.Uint32(buffer[4:8]),
		Thread:     binary.LittleEndian.Uint64(buffer[8:16]),
		ProcessId:  binary.LittleEndian.Uint32(buffer[16:20]),
		Process:    binary.LittleEndian.Uint64(buffer[20:28]),
		Result:     binary.LittleEndian.Uint32(buffer[36:40]),
	}

	if len(buffer) >= 52 {
		copy(threadPacket.ProcessName[:], buffer[28:44])
	}

	if h.threadMgr != nil {
		threadInfo := &thread.Thread{
			ThreadId:     threadPacket.ThreadId,
			ProcessId:    threadPacket.ProcessId,
			StartAddress: threadPacket.Thread,
			State:        thread.ThreadStateRunning,
		}
		h.threadMgr.AddThread(threadInfo)
	}
}

func (h *DebugEventHandler) HandleThreadExited(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send thread exited hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return err
	}

	h.decodeThreadExitedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeThreadExitedResponse(buffer []byte) {
	if len(buffer) < 40 {
		return
	}

	threadPacket := &DebuggeeDetailsAndSwitchThreadPacket{
		ActionType: binary.LittleEndian.Uint32(buffer[0:4]),
		ThreadId:   binary.LittleEndian.Uint32(buffer[4:8]),
		Thread:     binary.LittleEndian.Uint64(buffer[8:16]),
		ProcessId:  binary.LittleEndian.Uint32(buffer[16:20]),
		Process:    binary.LittleEndian.Uint64(buffer[20:28]),
		Result:     binary.LittleEndian.Uint32(buffer[36:40]),
	}

	if len(buffer) >= 52 {
		copy(threadPacket.ProcessName[:], buffer[28:44])
	}

	if h.threadMgr != nil {
		h.threadMgr.RemoveThread(threadPacket.ThreadId)
	}
}

func (h *DebugEventHandler) HandleModuleLoaded(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send module loaded hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive module loaded hook response: %v", err)
	}

	h.decodeModuleLoadedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeModuleLoadedResponse(buffer []byte) {
	if len(buffer) < 5 {
		return
	}

	result := &EventAndActionResult{
		IsSuccessful: buffer[0] != 0,
		Error:        binary.LittleEndian.Uint32(buffer[1:5]),
	}

	if result.IsSuccessful {
		slog.Info("Module loaded event registered successfully")
	} else {
		slog.Error("Module loaded event registration failed", "error", result.Error)
	}
}

func (h *DebugEventHandler) HandleModuleUnloaded(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send module unloaded hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive module unloaded hook response: %v", err)
	}

	h.decodeModuleUnloadedResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeModuleUnloadedResponse(buffer []byte) {
	if len(buffer) < 5 {
		return
	}

	result := &EventAndActionResult{
		IsSuccessful: buffer[0] != 0,
		Error:        binary.LittleEndian.Uint32(buffer[1:5]),
	}

	if result.IsSuccessful {
		slog.Info("Module unloaded event registered successfully")
	} else {
		slog.Error("Module unloaded event registration failed", "error", result.Error)
	}
}

func (h *DebugEventHandler) HandleEptHookReadWrite(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT read-write hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT read-write hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) decodeEptHookResponse(buffer []byte) {
	if len(buffer) < 5 {
		return
	}

	result := &EventAndActionResult{
		IsSuccessful: buffer[0] != 0,
		Error:        binary.LittleEndian.Uint32(buffer[1:5]),
	}

	if result.IsSuccessful {
		slog.Info("EPT hook event registered successfully")
	} else {
		slog.Error("EPT hook event registration failed", "error", result.Error)
	}
}

func (h *DebugEventHandler) HandleEptHookReadAndExecute(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT read-execute hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT read-execute hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookWriteAndExecute(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT write-execute hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT write-execute hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookRead(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT read hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT read hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookWrite(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT write hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT write hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookExecute(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT execute hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT execute hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookExecDetours(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT exec detours hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT exec detours hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleEptHookExecCC(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send EPT exec CC hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive EPT exec CC hook response: %v", err)
	}

	h.decodeEptHookResponse(response)
	return nil
}

func (h *DebugEventHandler) HandleSyscallHookEferSysret(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send SYSRET hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive SYSRET hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) HandleControlRegisterRead(event *DebugEvent) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send control register read hook request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive control register read hook response: %v", err)
	}

	h.decodeEventActionResult(response)
	return nil
}

func (h *DebugEventHandler) decodeEventActionResult(buffer []byte) {
	if len(buffer) < 5 {
		return
	}

	result := &EventAndActionResult{
		IsSuccessful: buffer[0] != 0,
		Error:        binary.LittleEndian.Uint32(buffer[1:5]),
	}

	if result.IsSuccessful {
		slog.Info("Event registered successfully")
	} else {
		slog.Error("Event registration failed", "error", result.Error)
	}
}

func (h *DebugEventHandler) Handle(event *DebugEvent) bool {
	if !event.IsEnabled {
		return false
	}

	switch event.Type {
	case EventStep:
		h.HandleStep(event)
	case EventBreakpoint:
		h.HandleBreakpoint(event)
	case EventException:
		h.HandleException(event)
	case EventProcessCreated:
		h.HandleProcessCreated(event)
	case EventProcessExited:
		h.HandleProcessExited(event)
	case EventThreadCreated:
		h.HandleThreadCreated(event)
	case EventThreadExited:
		h.HandleThreadExited(event)
	case EventModuleLoaded:
		h.HandleModuleLoaded(event)
	case EventModuleUnloaded:
		h.HandleModuleUnloaded(event)
	case EptHookReadWriteAndExecute:
		h.HandleEptHook(event)
	case EptHookReadWrite:
		h.HandleEptHookReadWrite(event)
	case EptHookReadAndExecute:
		h.HandleEptHookReadAndExecute(event)
	case EptHookWriteAndExecute:
		h.HandleEptHookWriteAndExecute(event)
	case EptHookRead:
		h.HandleEptHookRead(event)
	case EptHookWrite:
		h.HandleEptHookWrite(event)
	case EptHookExecute:
		h.HandleEptHookExecute(event)
	case EptHookExecDetours:
		h.HandleEptHookExecDetours(event)
	case EptHookExecCC:
		h.HandleEptHookExecCC(event)
	case SyscallHookEferSyscall:
		h.HandleSyscallHook(event)
	case SyscallHookEferSysret:
		h.HandleSyscallHookEferSysret(event)
	case CpuidInstructionExecution:
		h.HandleCpuid(event)
	case RdmsrInstructionExecution:
		h.HandleRdmsr(event)
	case WrmsrInstructionExecution:
		h.HandleWrmsr(event)
	case InInstructionExecution:
		h.HandleIn(event)
	case OutInstructionExecution:
		h.HandleOut(event)
	case ExceptionOccurred:
		h.HandleException(event)
	case ExternalInterruptOccurred:
		h.HandleInterrupt(event)
	case DebugRegistersAccessed:
		h.HandleDebugRegisters(event)
	case TscInstructionExecution:
		h.HandleTsc(event)
	case PmcInstructionExecution:
		h.HandlePmc(event)
	case VmcallInstructionExecution:
		h.HandleVmcall(event)
	case XsetbvInstructionExecution:
		h.HandleXsetbv(event)
	case ControlRegisterModified:
		h.HandleControlRegister(event)
	case ControlRegisterRead:
		h.HandleControlRegisterRead(event)
	case ControlRegister3Modified:
		h.HandleCr3(event)
	case TrapExecutionModeChanged:
		h.HandleTrap(event)
	case TrapExecutionInstructionTrace:
		h.HandleStep(event)
	default:
		return false
	}

	return true
}

func (h *DebugEventHandler) HandleRegisterEvent(event *DebugEvent, registerMgr *register.Provider) error {
	if h.driver == nil || !h.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := h.constructDebugEventBuffer(event)

	err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send register event request: %v", err)
	}

	response, err := h.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive register event response: %v", err)
	}

	regCtx := h.decodeRegisterContext(response)
	registerMgr.SetRegisterContext(regCtx)

	return nil
}

func (h *DebugEventHandler) decodeRegisterContext(buffer []byte) *register.RegisterContext {
	if len(buffer) < 200 {
		return &register.RegisterContext{}
	}

	regCtx := &register.RegisterContext{
		RAX: binary.LittleEndian.Uint64(buffer[0:8]),
		RCX: binary.LittleEndian.Uint64(buffer[8:16]),
		RDX: binary.LittleEndian.Uint64(buffer[16:24]),
		RBX: binary.LittleEndian.Uint64(buffer[24:32]),
		RSP: binary.LittleEndian.Uint64(buffer[32:40]),
		RBP: binary.LittleEndian.Uint64(buffer[40:48]),
		RSI: binary.LittleEndian.Uint64(buffer[48:56]),
		RDI: binary.LittleEndian.Uint64(buffer[56:64]),
		R8:  binary.LittleEndian.Uint64(buffer[64:72]),
		R9:  binary.LittleEndian.Uint64(buffer[72:80]),
		R10: binary.LittleEndian.Uint64(buffer[80:88]),
		R11: binary.LittleEndian.Uint64(buffer[88:96]),
		R12: binary.LittleEndian.Uint64(buffer[96:104]),
		R13: binary.LittleEndian.Uint64(buffer[104:112]),
		R14: binary.LittleEndian.Uint64(buffer[112:120]),
		R15: binary.LittleEndian.Uint64(buffer[120:128]),
	}

	if len(buffer) >= 144 {
		regCtx.RIP = binary.LittleEndian.Uint64(buffer[136:144])
	}

	if len(buffer) >= 152 {
		regCtx.RFLAGS = binary.LittleEndian.Uint32(buffer[144:148])
	}

	if len(buffer) >= 160 {
		regCtx.CS = binary.LittleEndian.Uint16(buffer[148:150])
		regCtx.DS = binary.LittleEndian.Uint16(buffer[150:152])
		regCtx.FS = binary.LittleEndian.Uint16(buffer[152:154])
		regCtx.GS = binary.LittleEndian.Uint16(buffer[154:156])
		regCtx.ES = binary.LittleEndian.Uint16(buffer[156:158])
		regCtx.SS = binary.LittleEndian.Uint16(buffer[158:160])
	}

	h.parseFlags(regCtx)

	return regCtx
}

func (h *DebugEventHandler) parseFlags(regCtx *register.RegisterContext) {
	flags := regCtx.RFLAGS
	regCtx.ZF = (flags & (1 << 6)) != 0
	regCtx.OF = (flags & (1 << 11)) != 0
	regCtx.CF = (flags & 1) != 0
	regCtx.PF = (flags & (1 << 2)) != 0
	regCtx.SF = (flags & (1 << 7)) != 0
	regCtx.TF = (flags & (1 << 8)) != 0
	regCtx.AF = (flags & (1 << 4)) != 0
	regCtx.DF = (flags & (1 << 10)) != 0
	regCtx.IF = (flags & (1 << 9)) != 0
}
