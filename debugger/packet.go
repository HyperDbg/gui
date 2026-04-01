package debugger

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
)

func NewMessage(msgType MessageType, payload []byte) *Message {
	return &Message{
		Header: MessageHeader{
			Type:     msgType,
			Length:   uint32(len(payload)),
			Sequence: 0,
			Flags:    0,
		},
		Payload: payload,
	}
}

func (m *Message) Serialize() []byte {
	buf := make([]byte, HeaderSize+len(m.Payload))
	binary.LittleEndian.PutUint32(buf[0:4], uint32(m.Header.Type))
	binary.LittleEndian.PutUint32(buf[4:8], m.Header.Length)
	binary.LittleEndian.PutUint64(buf[8:16], m.Header.Sequence)
	binary.LittleEndian.PutUint32(buf[16:20], m.Header.Flags)
	copy(buf[HeaderSize:], m.Payload)
	return buf
}

func ReadMessage(r io.Reader) (*Message, error) {
	headerBuf := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, headerBuf); err != nil {
		return nil, err
	}

	msg := &Message{
		Header: MessageHeader{
			Type:     MessageType(binary.LittleEndian.Uint32(headerBuf[0:4])),
			Length:   binary.LittleEndian.Uint32(headerBuf[4:8]),
			Sequence: binary.LittleEndian.Uint64(headerBuf[8:16]),
			Flags:    binary.LittleEndian.Uint32(headerBuf[16:20]),
		},
	}

	if msg.Header.Length > 0 {
		msg.Payload = make([]byte, msg.Header.Length)
		if _, err := io.ReadFull(r, msg.Payload); err != nil {
			return nil, err
		}
	}

	return msg, nil
}

func WriteMessage(w io.Writer, msg *Message) error {
	_, err := w.Write(msg.Serialize())
	return err
}

type Packet struct {
	client    *http.Client
	baseURL   string
	driverID  atomic.Uint64
	running   atomic.Bool
	connected atomic.Bool
	state     atomic.Int32
	callbacks map[MessageType][]EventCallback
	eventChan chan any
}

func NewPacket(baseURL string) Debugger {
	p := &Packet{
		client: &http.Client{
			Timeout: HTTPTimeout,
		},
		baseURL:   baseURL,
		callbacks: make(map[MessageType][]EventCallback),
		eventChan: make(chan any, 1000),
	}
	p.driverID.Store(1)
	p.state.Store(int32(StateTerminated))
	return p
}

func (p *Packet) Connect() error {
	p.running.Store(true)

	if err := p.Ping(); err != nil {
		return fmt.Errorf("failed to connect to driver: %w", err)
	}

	p.connected.Store(true)
	fmt.Printf("Connected to driver at %s\n", p.baseURL)

	go p.eventPoller()

	p.RegisterCallback(MsgTypeBreakpointEvent, p.handleBreakpointEvent)
	p.RegisterCallback(MsgTypeExceptionEvent, p.handleExceptionEvent)
	p.RegisterCallback(MsgTypeDebugPrintEvent, p.handleDebugPrintEvent)

	return nil
}

func (p *Packet) Disconnect() {
	p.running.Store(false)
	p.connected.Store(false)
}

func (p *Packet) IsConnected() bool {
	return p.connected.Load()
}

func (p *Packet) GetState() DebugState {
	return DebugState(p.state.Load())
}

func SendReceive[T ResponseType](p *Packet, jsonData []byte) *Response[T] {
	var action string
	var reqMap map[string]any
	if json.Unmarshal(jsonData, &reqMap) == nil {
		if a, ok := reqMap["action"].(string); ok {
			action = a
		}
	}

	url := fmt.Sprintf("%s/api/%s", p.baseURL, action)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Host", DriverHTTPHost)

	fmt.Printf("[HTTP] POST %s\n  body: %s\n", url, string(jsonData))

	response, err := p.client.Do(httpReq)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	var result Response[T]
	json.NewDecoder(response.Body).Decode(&result)
	return &result
}

func (p *Packet) eventPoller() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for p.running.Load() {
		<-ticker.C
		if !p.connected.Load() {
			continue
		}
	}
}

func (p *Packet) Status() (string, error) {
	data, _ := json.Marshal(map[string]string{"action": "status"})
	resp := SendReceive[Empty](p, data)
	if resp == nil {
		return "", fmt.Errorf("status request failed")
	}
	return resp.Message, nil
}

func (p *Packet) Ping() error {
	data, _ := json.Marshal(map[string]string{"action": "ping"})
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("ping failed")
	}
	return nil
}

func (p *Packet) RegisterCallback(msgType MessageType, cb EventCallback) {
	p.callbacks[msgType] = append(p.callbacks[msgType], cb)
}

func (p *Packet) RegisterCpuidCallback(cb EventCallback) {
	p.callbacks[MsgTypeCpuidEvent] = append(p.callbacks[MsgTypeCpuidEvent], cb)
}

func (p *Packet) GetEvent() any {
	select {
	case event := <-p.eventChan:
		return event
	default:
		return nil
	}
}

func (p *Packet) WaitForEvent(timeout time.Duration) *Message {
	select {
	case event := <-p.eventChan:
		if msg, ok := event.(*Message); ok {
			return msg
		}
		return nil
	case <-time.After(timeout):
		return nil
	}
}

func (p *Packet) LoadVmm() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "load_vmm"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("load_vmm failed: %s", resp.Message)
	}
	p.state.Store(int32(StateRunning))
	return nil
}

func (p *Packet) UnloadVmm() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "unload_vmm"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("unload_vmm failed: %s", resp.Message)
	}
	p.state.Store(int32(StateTerminated))
	return nil
}

func (p *Packet) AttachProcess(processID uint32) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "attach_process",
		"process_id": processID,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("attach process failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) DetachProcess() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "detach_process"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("detach process failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) StartProcess(exePath string) (uint32, error) {
	type ProcessInformation struct {
		ProcessHandle uintptr
		ThreadHandle  uintptr
		ProcessID     uint32
		ThreadID      uint32
	}

	type StartupInfo struct {
		Cb    uint32
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
		Flags uint32
		_     uint16
		_     uint16
		_     uintptr
		_     uintptr
		_     uintptr
		_     uintptr
	}

	const CREATE_SUSPENDED = 0x00000004

	var si StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	var pi ProcessInformation

	cmdLinePtr, _ := syscall.UTF16PtrFromString(exePath)

	ret, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("CreateProcessW").Call(
		0,
		uintptr(unsafe.Pointer(cmdLinePtr)),
		0, 0, 0,
		CREATE_SUSPENDED,
		0, 0,
		uintptr(unsafe.Pointer(&si)),
		uintptr(unsafe.Pointer(&pi)),
	)

	if ret == 0 {
		return 0, fmt.Errorf("CreateProcess failed: %v", err)
	}

	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "start_process",
		"process_id": pi.ProcessID,
		"thread_id":  pi.ThreadID,
		"exe_path":   exePath,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		syscall.CloseHandle(syscall.Handle(pi.ProcessHandle))
		syscall.CloseHandle(syscall.Handle(pi.ThreadHandle))
		return 0, fmt.Errorf("start process failed: %s", resp.Message)
	}

	ret, _, _ = syscall.NewLazyDLL("kernel32.dll").NewProc("ResumeThread").Call(pi.ThreadHandle)
	_ = ret

	syscall.CloseHandle(syscall.Handle(pi.ProcessHandle))
	syscall.CloseHandle(syscall.Handle(pi.ThreadHandle))

	return pi.ProcessID, nil
}

func (p *Packet) KillProcess(processID uint32) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "kill_process",
		"process_id": processID,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("kill process failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) SetBreakpoint(address uint64, bpType BreakpointType) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":  "set_breakpoint",
		"address": address,
		"type":    int(bpType),
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("set breakpoint failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) RemoveBreakpoint(breakpointID uint64) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":        "remove_breakpoint",
		"breakpoint_id": breakpointID,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("remove breakpoint failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) Continue() error {
	if p.GetState() != StatePaused {
		return fmt.Errorf("debugger not paused")
	}
	data := mylog.Check2(json.Marshal(map[string]string{"action": "continue"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("continue failed: %s", resp.Message)
	}
	p.state.Store(int32(StateRunning))
	return nil
}

func (p *Packet) Pause() error {
	if p.GetState() != StateRunning {
		return fmt.Errorf("debugger not running")
	}
	data := mylog.Check2(json.Marshal(map[string]string{"action": "pause"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("pause failed: %s", resp.Message)
	}
	p.state.Store(int32(StatePaused))
	return nil
}

func (p *Packet) StepInto() error {
	if p.GetState() != StatePaused {
		return fmt.Errorf("debugger not paused")
	}
	data := mylog.Check2(json.Marshal(map[string]string{"action": "step_into"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("step into failed: %s", resp.Message)
	}
	p.state.Store(int32(StateStepping))
	return nil
}

func (p *Packet) StepOver() error {
	if p.GetState() != StatePaused {
		return fmt.Errorf("debugger not paused")
	}
	data := mylog.Check2(json.Marshal(map[string]string{"action": "step_over"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("step over failed: %s", resp.Message)
	}
	p.state.Store(int32(StateStepping))
	return nil
}

func (p *Packet) StepOut() error {
	if p.GetState() != StatePaused {
		return fmt.Errorf("debugger not paused")
	}
	data := mylog.Check2(json.Marshal(map[string]string{"action": "step_out"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("step out failed: %s", resp.Message)
	}
	p.state.Store(int32(StateStepping))
	return nil
}

func (p *Packet) ReadMemory(address uint64, size uint32) ([]byte, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":  "read_memory",
		"address": address,
		"size":    size,
	}))
	resp := SendReceive[[]byte](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("read memory failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) WriteMemory(address uint64, data []byte) error {
	jsonData := mylog.Check2(json.Marshal(map[string]any{
		"action":  "write_memory",
		"address": address,
		"data":    data,
	}))
	resp := SendReceive[Empty](p, jsonData)
	if resp == nil || !resp.Success {
		return fmt.Errorf("write memory failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) ReadRegisters() (*RegisterState, error) {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "read_registers"}))
	resp := SendReceive[*RegisterState](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("read registers failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) WriteRegisters(regs *RegisterState) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action": "write_registers",
		"data":   regs,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("write registers failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) GetProcessList() ([]ProcessInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "get_process_list"}))
	resp := SendReceive[[]ProcessInfo](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("get process list failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) GetThreadList(processID uint32) ([]ThreadInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "get_thread_list",
		"process_id": processID,
	}))
	resp := SendReceive[[]ThreadInfo](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("get thread list failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) GetModuleList(processID uint32) ([]ModuleInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "get_module_list",
		"process_id": processID,
	}))
	resp := SendReceive[[]ModuleInfo](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("get module list failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) handleBreakpointEvent(event any) {
	p.state.Store(int32(StatePaused))

	if msg, ok := event.(*Message); ok {
		e := parseBreakpointEvent(msg.Payload)
		fmt.Printf("[Breakpoint] PID=%d TID=%d Address=0x%X\n",
			e.Header.ProcessID, e.Header.ThreadID, e.Address)
	}
}

func (p *Packet) handleExceptionEvent(event any) {
	p.state.Store(int32(StatePaused))

	if msg, ok := event.(*Message); ok {
		e := parseExceptionEvent(msg.Payload)
		fmt.Printf("[Exception] PID=%d TID=%d Code=0x%x Address=0x%X\n",
			e.Header.ProcessID, e.Header.ThreadID, e.ExceptionCode, e.Address)
	}
}

func (p *Packet) handleDebugPrintEvent(event any) {
	if msg, ok := event.(*Message); ok {
		e := parseDebugPrintEvent(msg.Payload)
		fmt.Printf("[DebugPrint] %s\n", e.Message)
	}
}

func parseBreakpointEvent(data []byte) *BreakpointEvent {
	if len(data) < 32 {
		return &BreakpointEvent{}
	}
	return &BreakpointEvent{
		Header: EventHeader{
			ProcessID: ProcessId(binary.LittleEndian.Uint32(data[0:4])),
			ThreadID:  ThreadId(binary.LittleEndian.Uint32(data[4:8])),
			CoreID:    binary.LittleEndian.Uint32(data[8:12]),
		},
		Address:      Address(binary.LittleEndian.Uint64(data[12:20])),
		BreakpointID: binary.LittleEndian.Uint64(data[20:28]),
	}
}

func parseExceptionEvent(data []byte) *ExceptionEvent {
	if len(data) < 40 {
		return &ExceptionEvent{}
	}
	return &ExceptionEvent{
		Header: EventHeader{
			ProcessID: ProcessId(binary.LittleEndian.Uint32(data[0:4])),
			ThreadID:  ThreadId(binary.LittleEndian.Uint32(data[4:8])),
			CoreID:    binary.LittleEndian.Uint32(data[8:12]),
		},
		ExceptionCode: ExceptionCode(binary.LittleEndian.Uint32(data[12:16])),
		Address:       Address(binary.LittleEndian.Uint64(data[16:24])),
		ErrorCode:     binary.LittleEndian.Uint64(data[24:32]),
	}
}

func parseDebugPrintEvent(data []byte) *DebugPrintEvent {
	if len(data) < 20 {
		return &DebugPrintEvent{}
	}
	msgLen := binary.LittleEndian.Uint32(data[16:20])
	if int(20+msgLen) > len(data) {
		msgLen = uint32(len(data) - 20)
	}
	return &DebugPrintEvent{
		Header: EventHeader{
			ProcessID: ProcessId(binary.LittleEndian.Uint32(data[0:4])),
			ThreadID:  ThreadId(binary.LittleEndian.Uint32(data[4:8])),
			CoreID:    binary.LittleEndian.Uint32(data[8:12]),
		},
		Level:   LogLevel(binary.LittleEndian.Uint32(data[12:16])),
		Message: string(data[20 : 20+msgLen]),
	}
}

func (p *Packet) ExecuteScript(script string) (string, error) {
	return p.ExecuteScriptWithContext(nil, script)
}

func (p *Packet) ExecuteScriptWithContext(ctx context.Context, script string) (string, error) {
	engine := NewScriptEngine(p)
	if ctx != nil {
		return engine.ExecuteWithContext(ctx, script)
	}
	return engine.Execute(script)
}

func (p *Packet) Disassemble(address uint64, bytes []byte, maxInstructions uint32) ([]Instruction, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":           "disassemble",
		"address":          address,
		"data":             bytes,
		"max_instructions": maxInstructions,
	}))
	resp := SendReceive[[]Instruction](p, data)
	if resp == nil || !resp.Success {
		return nil, fmt.Errorf("disassemble failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) LoadSymbols(pdbPath string) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":   "load_symbols",
		"pdb_path": pdbPath,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("load symbols failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) UnloadSymbols() {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "unload_symbols"}))
	SendReceive[Empty](p, data)
}

func (p *Packet) GetSymbolByName(name string) (SymbolInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action": "get_symbol_by_name",
		"name":   name,
	}))
	resp := SendReceive[SymbolInfo](p, data)
	if resp == nil || !resp.Success {
		return SymbolInfo{}, fmt.Errorf("get symbol failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) GetSymbolByAddress(address uint64) (SymbolInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":  "get_symbol_by_address",
		"address": address,
	}))
	resp := SendReceive[SymbolInfo](p, data)
	if resp == nil || !resp.Success {
		return SymbolInfo{}, fmt.Errorf("get symbol failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) GetFunctionByAddress(address uint64) (FunctionInfo, error) {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":  "get_function_by_address",
		"address": address,
	}))
	resp := SendReceive[FunctionInfo](p, data)
	if resp == nil || !resp.Success {
		return FunctionInfo{}, fmt.Errorf("get function failed: %s", resp.Message)
	}
	return resp.Data, nil
}

func (p *Packet) InstallHookScript(script *HookScript) error {
	code := ExtractClosureBody(script.OnMatch)
	if code == "" {
		return fmt.Errorf("failed to extract closure body")
	}

	fmt.Printf("[HookScript] Extracted code from closure:\n%s\n", code)

	data := mylog.Check2(json.Marshal(map[string]any{
		"action":    "install_hook_script",
		"api_name":  script.ApiName,
		"hook_type": script.HookType,
		"filter":    script.Filter,
		"code":      code,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("install hook script failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) Initialize() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "initialize"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("initialize failed: %s", resp.Message)
	}
	p.state.Store(int32(StateInitialized))
	return nil
}

func (p *Packet) Terminate() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "terminate"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("terminate failed: %s", resp.Message)
	}
	p.state.Store(int32(StateTerminated))
	return nil
}

func (p *Packet) StartNetworkServer(addr string) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action": "start_network_server",
		"addr":   addr,
	}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("start network server failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) StopNetworkServer() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "stop_network_server"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("stop network server failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) KernelDebuggerInitialize() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "kernel_debugger_initialize"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("kernel debugger initialize failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) KernelDebuggerUninitialize() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "kernel_debugger_uninitialize"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("kernel debugger uninitialize failed: %s", resp.Message)
	}
	return nil
}
