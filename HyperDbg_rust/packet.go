package hyperdbgrust

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/ddkwork/golibrary/std/mylog"
)

const (
	ProtocolVersion = 1
	DefaultPort     = 9527
	MaxMessageSize  = 4096
	HeaderSize      = 20

	DriverHTTPHost = "127.0.0.1"
	DriverHTTPPort = 50080
	HTTPTimeout    = 10 * time.Second
)

type MessageType uint32

const (
	MsgTypeBreakpointEvent   MessageType = 1
	MsgTypeExceptionEvent    MessageType = 2
	MsgTypeMemoryAccessEvent MessageType = 3
	MsgTypeSyscallEvent      MessageType = 4
	MsgTypeProcessEvent      MessageType = 5
	MsgTypeThreadEvent       MessageType = 6
	MsgTypeModuleEvent       MessageType = 7
	MsgTypeDebugPrintEvent   MessageType = 8
	MsgTypeVmxExitEvent      MessageType = 9
	MsgTypeLogEvent          MessageType = 10

	MsgTypeInitialize        MessageType = 100
	MsgTypeTerminate         MessageType = 101
	MsgTypeConnect           MessageType = 102
	MsgTypeDisconnect        MessageType = 103
	MsgTypeAttachProcess     MessageType = 104
	MsgTypeDetachProcess     MessageType = 105
	MsgTypeStartProcess      MessageType = 106
	MsgTypeKillProcess       MessageType = 107
	MsgTypeSetBreakpoint     MessageType = 108
	MsgTypeRemoveBreakpoint  MessageType = 109
	MsgTypeEnableBreakpoint  MessageType = 110
	MsgTypeDisableBreakpoint MessageType = 111
	MsgTypeListBreakpoints   MessageType = 112
	MsgTypeContinue          MessageType = 113
	MsgTypePause             MessageType = 114
	MsgTypeStepInto          MessageType = 115
	MsgTypeStepOver          MessageType = 116
	MsgTypeStepOut           MessageType = 117
	MsgTypeReadRegisters     MessageType = 118
	MsgTypeWriteRegisters    MessageType = 119
	MsgTypeReadMemory        MessageType = 120
	MsgTypeWriteMemory       MessageType = 121
	MsgTypeGetProcessList    MessageType = 122
	MsgTypeGetThreadList     MessageType = 123
	MsgTypeGetModuleList     MessageType = 124
	MsgTypeGetCallStack      MessageType = 125

	MsgTypeResponse MessageType = 200
	MsgTypeError    MessageType = 201
)

type BreakpointType uint32

const (
	BreakpointSoftware BreakpointType = 0
	BreakpointHardware BreakpointType = 1
	BreakpointHidden   BreakpointType = 2
	BreakpointEpt      BreakpointType = 3
)

type DebugState uint32

const (
	StateRunning    DebugState = 0
	StatePaused     DebugState = 1
	StateStepping   DebugState = 2
	StateTerminated DebugState = 3
)

type MessageHeader struct {
	Type     MessageType
	Length   uint32
	Sequence uint64
	Flags    uint32
}

type Message struct {
	Header  MessageHeader
	Payload []byte
}

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

type RegisterState struct {
	RAX    uint64 `json:"RAX"`
	RBX    uint64 `json:"RBX"`
	RCX    uint64 `json:"RCX"`
	RDX    uint64 `json:"RDX"`
	RSI    uint64 `json:"RSI"`
	RDI    uint64 `json:"RDI"`
	RBP    uint64 `json:"RBP"`
	RSP    uint64 `json:"RSP"`
	R8     uint64 `json:"R8"`
	R9     uint64 `json:"R9"`
	R10    uint64 `json:"R10"`
	R11    uint64 `json:"R11"`
	R12    uint64 `json:"R12"`
	R13    uint64 `json:"R13"`
	R14    uint64 `json:"R14"`
	R15    uint64 `json:"R15"`
	RIP    uint64 `json:"RIP"`
	RFLAGS uint64 `json:"RFLAGS"`
	CR0    uint64 `json:"CR0"`
	CR2    uint64 `json:"CR2"`
	CR3    uint64 `json:"CR3"`
	CR4    uint64 `json:"CR4"`
	DR0    uint64 `json:"DR0"`
	DR1    uint64 `json:"DR1"`
	DR2    uint64 `json:"DR2"`
	DR3    uint64 `json:"DR3"`
	DR6    uint64 `json:"DR6"`
	DR7    uint64 `json:"DR7"`
	GDTR   uint64 `json:"GDTR"`
	GSBase uint64 `json:"GSBase"`
	FSBase uint64 `json:"FSBase"`
}

type ProcessInfo struct {
	ProcessID   uint32 `json:"process_id"`
	ImageName   string `json:"image_name"`
	BaseAddress string `json:"base_address"`
	Size        uint64 `json:"size"`
	CR3         uint64 `json:"cr3"`
}

type ThreadInfo struct {
	ThreadID     uint32 `json:"thread_id"`
	ProcessID    uint32 `json:"process_id"`
	StartAddress string `json:"start_address"`
	TEB          string `json:"teb"`
	State        uint32 `json:"state"`
}

type ModuleInfo struct {
	BaseAddress string `json:"base_address"`
	Size        uint64 `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
}

type BreakpointInfo struct {
	ID        uint64 `json:"id"`
	Address   string `json:"address"`
	Type      int    `json:"type"`
	ProcessID uint32 `json:"process_id"`
	Enabled   bool   `json:"enabled"`
	HitCount  uint64 `json:"hit_count"`
}

type BreakpointEvent struct {
	ProcessID    uint32        `json:"process_id"`
	ThreadID     uint32        `json:"thread_id"`
	CoreID       uint32        `json:"core_id"`
	Address      string        `json:"address"`
	BreakpointID uint64        `json:"breakpoint_id"`
	Registers    RegisterState `json:"registers"`
}

type ExceptionEvent struct {
	ProcessID     uint32        `json:"process_id"`
	ThreadID      uint32        `json:"thread_id"`
	CoreID        uint32        `json:"core_id"`
	ExceptionCode uint32        `json:"exception_code"`
	Address       string        `json:"address"`
	ErrorCode     uint64        `json:"error_code"`
	Registers     RegisterState `json:"registers"`
}

type DebugPrintEvent struct {
	ProcessID uint32 `json:"process_id"`
	ThreadID  uint32 `json:"thread_id"`
	CoreID    uint32 `json:"core_id"`
	Message   string `json:"message"`
	Level     uint32 `json:"level"`
}

type SyscallEvent struct {
	ProcessID     uint32 `json:"process_id"`
	ThreadID      uint32 `json:"thread_id"`
	CoreID        uint32 `json:"core_id"`
	SyscallNumber uint64 `json:"syscall_number"`
	RAX           uint64 `json:"rax"`
	RCX           uint64 `json:"rcx"`
	RDX           uint64 `json:"rdx"`
	R8            uint64 `json:"r8"`
	R9            uint64 `json:"r9"`
	R10           uint64 `json:"r10"`
	RSP           uint64 `json:"rsp"`
}

type Response[T ResponseType] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type Empty = struct{}

type ResponseType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string | ~bool |
		[]byte | *RegisterState | []ProcessInfo | []ThreadInfo | []ModuleInfo | []BreakpointInfo |
		Empty
}

type EventCallback func(event any)

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

func NewPacket() Debugger {
	p := &Packet{
		client: &http.Client{
			Timeout: HTTPTimeout,
		},
		baseURL:   fmt.Sprintf("http://%s:%d", DriverHTTPHost, DriverHTTPPort),
		callbacks: make(map[MessageType][]EventCallback),
		eventChan: make(chan any, 1000),
	}
	p.driverID.Store(1)
	p.state.Store(int32(StateTerminated))
	return p
}

func (p *Packet) Start() error {
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

func (p *Packet) Stop() {
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
	url := fmt.Sprintf("%s/api", p.baseURL)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Host", DriverHTTPHost)

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

func (p *Packet) Ping() error {
	url := fmt.Sprintf("%s/ping", p.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	response, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var resp Response[Empty]
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("ping failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) Status() (string, error) {
	url := fmt.Sprintf("%s/status", p.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := p.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var resp Response[Empty]
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (p *Packet) RegisterCallback(msgType MessageType, cb EventCallback) {
	p.callbacks[msgType] = append(p.callbacks[msgType], cb)
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

func (p *Packet) GetConnectedDrivers() []uint64 {
	if p.connected.Load() {
		return []uint64{p.driverID.Load()}
	}
	return nil
}

func (p *Packet) Initialize() error {
	data := mylog.Check2(json.Marshal(map[string]string{"action": "initialize"}))
	resp := SendReceive[Empty](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("initialize failed: %s", resp.Message)
	}
	p.state.Store(int32(StateRunning))
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

func (p *Packet) AttachProcess(processID uint32) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":     "attach_process",
		"process_id": fmt.Sprintf("%d", processID),
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

func (p *Packet) SetBreakpoint(address uint64, bpType BreakpointType) error {
	data := mylog.Check2(json.Marshal(map[string]any{
		"action":  "set_breakpoint",
		"address": fmt.Sprintf("0x%x", address),
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
		"breakpoint_id": fmt.Sprintf("0x%x", breakpointID),
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
		"address": fmt.Sprintf("0x%x", address),
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
		"address": fmt.Sprintf("0x%x", address),
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
		"process_id": fmt.Sprintf("%d", processID),
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
		"process_id": fmt.Sprintf("%d", processID),
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
		fmt.Printf("[Breakpoint] PID=%d TID=%d Address=0x%x\n",
			e.ProcessID, e.ThreadID, e.Address)
	}
}

func (p *Packet) handleExceptionEvent(event any) {
	p.state.Store(int32(StatePaused))

	if msg, ok := event.(*Message); ok {
		e := parseExceptionEvent(msg.Payload)
		fmt.Printf("[Exception] PID=%d TID=%d Code=0x%x Address=0x%x\n",
			e.ProcessID, e.ThreadID, e.ExceptionCode, e.Address)
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
		ProcessID:    binary.LittleEndian.Uint32(data[0:4]),
		ThreadID:     binary.LittleEndian.Uint32(data[4:8]),
		CoreID:       binary.LittleEndian.Uint32(data[8:12]),
		Address:      fmt.Sprintf("0x%x", binary.LittleEndian.Uint64(data[12:20])),
		BreakpointID: binary.LittleEndian.Uint64(data[20:28]),
	}
}

func parseExceptionEvent(data []byte) *ExceptionEvent {
	if len(data) < 40 {
		return &ExceptionEvent{}
	}
	return &ExceptionEvent{
		ProcessID:     binary.LittleEndian.Uint32(data[0:4]),
		ThreadID:      binary.LittleEndian.Uint32(data[4:8]),
		CoreID:        binary.LittleEndian.Uint32(data[8:12]),
		ExceptionCode: binary.LittleEndian.Uint32(data[12:16]),
		Address:       fmt.Sprintf("0x%x", binary.LittleEndian.Uint64(data[16:24])),
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
		ProcessID: binary.LittleEndian.Uint32(data[0:4]),
		ThreadID:  binary.LittleEndian.Uint32(data[4:8]),
		CoreID:    binary.LittleEndian.Uint32(data[8:12]),
		Level:     binary.LittleEndian.Uint32(data[12:16]),
		Message:   string(data[20 : 20+msgLen]),
	}
}

func (p *Packet) WaitForDriver(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if err := p.Ping(); err == nil {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("timeout waiting for driver")
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
