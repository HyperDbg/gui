// Package kernellvl — kd.go
//
// Implements the kernel-mode debugger state machine and packet/IOCTL helpers.
// The C++ counterpart is libhyperdbg/code/debugger/kernel-level/kd.cpp; it
// owns:
//   - the kernel-debugger connection state (handshaking, running/paused,
//     current remote core, ignore-pause flag)
//   - the synchronization-event handle table used to wait for, and wake up,
//     paused-debuggee responses (DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_*)
//   - the IOCTL path to the local VMM driver on the debuggee side
//     (IOCTL_PREPARE_DEBUGGEE, IOCTL_PAUSE_PACKET_RECEIVED,
//     IOCTL_DEBUGGER_REGISTER_EVENT, IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,
//     IOCTL_DEBUGGER_MODIFY_EVENTS, IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER,
//     IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED,
//     IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER)
//   - the serial/named-pipe packet framing helpers (KdComputeDataChecksum,
//     KdCheckForTheEndOfTheBuffer, KdCommandPacketToDebuggee, ...)
//
// In the Go rewrite the global state from the C side
// (g_IsSerialConnectedToRemoteDebuggee, g_IsDebuggeeRunning,
// g_KernelSyncronizationObjectsHandleTable, g_CurrentRemoteCore, ...) is owned
// by the KdState struct so that multiple debugger instances can coexist
// (GUI/MCP requirement). All IOCTLs go through comm.Device; the serial/named-
// pipe transport is abstracted behind the Transport interface so a GUI/MCP
// host can plug its own wire.
package kernellvl

import (
	"encoding/binary"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
)

// Output abstracts command output. CLI passes os.Stdout (wrapped), GUI passes
// a widget writer, MCP passes a JSON channel. Implementations must be
// goroutine-safe if KdState is used concurrently.
type Output interface {
	Printf(format string, args ...any) error
}

// Constants mirroring the C++ #defines in Constants.h and debugger.h. They are
// not in go-libhyperdbg/types/sdk.go because they are plain numeric macros
// (not enum/struct types).
const (
	// MaxKernelSyncObjects mirrors DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS.
	MaxKernelSyncObjects = 0x40

	// DebuggerEventTagStartSeed mirrors DebuggerEventTagStartSeed from Constants.h.
	// Event tags returned to the user are (raw_tag - DebuggerEventTagStartSeed).
	DebuggerEventTagStartSeed = 0x1000000

	// DebuggerDebuggeeIsRunningNoCore mirrors DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE.
	DebuggerDebuggeeIsRunningNoCore = 0xffffffff

	// MaximumInstrSize mirrors MAXIMUM_INSTR_SIZE.
	MaximumInstrSize = 16

	// IndicatorOfHyperdbgPacket mirrors INDICATOR_OF_HYPERDBG_PACKET ("HYPERDBG").
	IndicatorOfHyperdbgPacket uint64 = 0x4859504552444247

	// Serial end-of-buffer markers, mirror SERIAL_END_OF_BUFFER_CHAR_*.
	SerialEndOfBufferChar1 = 0x00
	SerialEndOfBufferChar2 = 0x80
	SerialEndOfBufferChar3 = 0xEE
	SerialEndOfBufferChar4 = 0xFF
	// SerialEndOfBufferCharsCount mirrors SERIAL_END_OF_BUFFER_CHARS_COUNT.
	SerialEndOfBufferCharsCount = 4

	// MaxSerialPacketSize mirrors MaxSerialPacketSize (20 * NORMAL_PAGE_SIZE).
	MaxSerialPacketSize = 20 * 4096

	// DebuggerOperationWasSuccessful mirrors DEBUGGER_OPERATION_WAS_SUCCESSFUL.
	DebuggerOperationWasSuccessful uint32 = 0xFFFFFFFF
)

// KernelSyncObject enumerates the synchronization slots used by the kernel
// debugger. Mirrors DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_*.
type KernelSyncObject uint32

const (
	SyncObjectIsDebuggerRunning KernelSyncObject = iota
	SyncObjectStartedPacketReceived
	SyncObjectPausedDebuggeeDetails
	SyncObjectCoreSwitchingResult
	SyncObjectProcessSwitchingResult
	SyncObjectThreadSwitchingResult
	SyncObjectScriptRunningResult
	SyncObjectScriptFormatsResult
	SyncObjectDebuggeeFinishedCommandExecution
	SyncObjectFlushResult
	SyncObjectRegisterEvent
	SyncObjectAddActionToEvent
	SyncObjectModifyAndQueryEvent
	SyncObjectReadRegisters
	SyncObjectBp
	SyncObjectListOrModifyBreakpoints
	SyncObjectReadMemory
	SyncObjectEditMemory
	SyncObjectSymbolReload
	SyncObjectTestQuery
	SyncObjectCallstackResult
	SyncObjectSearchQueryResult
	SyncObjectVa2paAndPa2vaResult
	SyncObjectPteResult
	SyncObjectShortCircuitingEventState
	SyncObjectPageInState
	SyncObjectWriteRegister
	SyncObjectPcitreeResult
	SyncObjectApicActions
	SyncObjectPcidevinfoResult
	SyncObjectIdtEntries
	SyncObjectSmiOperationResult
	SyncObjectHypertraceLbrDumpResult
	SyncObjectHypertracePtOperationResult
	SyncObjectUserCpuidResult
)

// KernelSyncEventState mirrors DEBUGGER_SYNCRONIZATION_EVENTS_STATE: each slot
// is either free (IsOnWaitingState=false) or held by a goroutine waiting on
// the embedded event.
type KernelSyncEventState struct {
	IsOnWaitingState bool
	// done is signalled by ReceivedKernelResponse to wake the waiting
	// goroutine. A nil channel is equivalent to "no waiter".
	done chan struct{}
}

// BpInfo tracks a single kernel breakpoint registered with the debuggee.
type BpInfo struct {
	Tag     uint64
	Address uint64
	Enabled bool
}

// Transport is the wire-level transport for debugger→debuggee packets.
// Implementations wrap a serial port or a named pipe. The C++ side uses
// ReadFile/WriteFile directly on g_SerialRemoteComPortHandle; in Go we let
// the caller plug github.com/Microsoft/go-winio or any other transport.
type Transport interface {
	// Write sends len(buf) bytes to the debuggee. It must return an error if
	// fewer bytes were written.
	Write(buf []byte) error
	// Read reads up to len(buf) bytes from the debuggee. It blocks until at
	// least one byte is available.
	Read(buf []byte) (int, error)
	// Close releases the transport resource.
	Close() error
}

// KdState owns the kernel-mode debugger state. All mutable fields are guarded
// by mu. A KdState is bound to one comm.Device (the local VMM handle on the
// debuggee side) and one optional Transport (the serial/named-pipe link to
// the remote side). Either may be nil if the corresponding role is not
// active.
type KdState struct {
	mu sync.Mutex

	device    *comm.Device
	transport Transport
	out       Output

	initialised bool

	// Connection flags mirroring the C++ g_Is* globals.
	isSerialConnectedToRemoteDebuggee bool
	isSerialConnectedToRemoteDebugger bool
	isDebuggeeRunning                 bool
	isDebuggeeInHandshakingPhase      bool
	isDebuggerConnectedToNamedPipe    bool
	serialConnectionAlreadyClosed     bool
	ignorePauseRequests               bool
	ignoreNewLoggingMessages          bool
	sharedEventStatus                 bool
	shouldPreviousCommandBeContinued  bool

	// Current operating core on the debuggee (DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE
	// when no core is selected).
	currentRemoteCore uint32

	// Current instruction bytes at RIP and 32-bit flag, populated by the
	// listener on each pause packet.
	currentRunningInstruction [MaximumInstrSize]byte
	isRunningInstruction32Bit bool

	// Breakpoints registered with the debuggee, keyed by event tag.
	breakpoints map[uint64]*BpInfo

	// Sync objects handle table.
	syncObjects [MaxKernelSyncObjects]KernelSyncEventState

	// Result buffers mirroring g_DebuggeeResultOf* in the C++ code. They are
	// written by the listener goroutine and read by the waiting IOCTL caller.
	resultOfRegisteringEvent    hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultOfAddingActionsTo     hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultOfEvaluatedExpression uint64
	errorStateOfEvaluatedExpr   uint32
	kernelBaseAddress           uint64

	// requestBuffers holds caller-supplied response buffers for the
	// DbgWaitSetKernelRequestData / DbgWaitGetKernelRequestData pair. Allocated
	// lazily by SetKernelRequestData. Mirrors the per-slot caller buffers used
	// in the C++ listener for read-register / read-memory / etc. responses.
	requestBuffers requestBuffersMap
}

// NewKdState creates a kernel-debugger state bound to the given device and
// output sink. Either may be nil: a nil device is useful for pure serial
// debugging (debugger side only); a nil output silences diagnostics. The
// returned KdState is ready for Initialise to be called.
func NewKdState(device *comm.Device, out Output) *KdState {
	return &KdState{
		device:            device,
		out:               out,
		currentRemoteCore: DebuggerDebuggeeIsRunningNoCore,
		breakpoints:       make(map[uint64]*BpInfo),
	}
}

// SetTransport installs the wire-level transport used by the packet-sending
// helpers (KdCommandPacketToDebuggee, KdSendPacketToDebuggee, ...). Pass nil
// to detach from the remote.
func (k *KdState) SetTransport(t Transport) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.transport = t
}

// SetDevice replaces the underlying VMM device handle. Useful when the VMM
// driver is loaded after the KdState is created.
func (k *KdState) SetDevice(d *comm.Device) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.device = d
}

// Device returns the underlying VMM device handle (may be nil).
func (k *KdState) Device() *comm.Device {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.device
}

// Initialise sets up the synchronization event handle table. Mirrors the
// handle-table init in KdPrepareSerialConnectionToRemoteSystem. Idempotent.
func (k *KdState) Initialise() {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.initialised {
		return
	}
	for i := range k.syncObjects {
		k.syncObjects[i] = KernelSyncEventState{
			IsOnWaitingState: false,
			done:             nil,
		}
	}
	k.initialised = true
}

// IsInitialised reports whether Initialise has been called.
func (k *KdState) IsInitialised() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.initialised
}

// IsDebuggeeRunning reports whether the remote debuggee is currently running
// (i.e. not paused). Mirrors g_IsDebuggeeRunning.
func (k *KdState) IsDebuggeeRunning() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.isDebuggeeRunning
}

// IsConnectedToRemoteDebuggee reports whether the debugger is connected to a
// remote debuggee (i.e. this side is the debugger). Mirrors
// g_IsSerialConnectedToRemoteDebuggee.
func (k *KdState) IsConnectedToRemoteDebuggee() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.isSerialConnectedToRemoteDebuggee
}

// IsConnectedToRemoteDebugger reports whether the debuggee is connected to a
// remote debugger (i.e. this side is the debuggee). Mirrors
// g_IsSerialConnectedToRemoteDebugger.
func (k *KdState) IsConnectedToRemoteDebugger() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.isSerialConnectedToRemoteDebugger
}

// CurrentRemoteCore returns the currently selected core on the debuggee, or
// DebuggerDebuggeeIsRunningNoCore if none.
func (k *KdState) CurrentRemoteCore() uint32 {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.currentRemoteCore
}

// SetCurrentRemoteCore records the currently selected core. Called by the
// listener when a pause packet arrives.
func (k *KdState) SetCurrentRemoteCore(core uint32) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.currentRemoteCore = core
}

// SetIgnorePauseRequests toggles the g_IgnorePauseRequests flag. When true,
// KdBreakControlCheckAndPauseDebugger is a no-op.
func (k *KdState) SetIgnorePauseRequests(ignore bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.ignorePauseRequests = ignore
}

// IsIgnoreNewLoggingMessages reports the g_IgnoreNewLoggingMessages flag.
// When true, the listener drops DEBUGGEE_MESSAGE_PACKET payloads without
// printing them. Set to true by the listener on each pause and reset to false
// by KdCommandGoto / KdContinue so logging resumes while the debuggee runs.
func (k *KdState) IsIgnoreNewLoggingMessages() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.ignoreNewLoggingMessages
}

// SetIgnoreNewLoggingMessages updates the g_IgnoreNewLoggingMessages flag.
// Called by the listener (on pause / on continue) and by KdSendPacketToDebuggee
// (which clears the flag so outbound traffic does not get muted).
func (k *KdState) SetIgnoreNewLoggingMessages(ignore bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.ignoreNewLoggingMessages = ignore
}

// SetConnectedToRemoteDebuggee marks the state as connected to a remote
// debuggee. Called by the connection-establishment code on success.
func (k *KdState) SetConnectedToRemoteDebuggee(v bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.isSerialConnectedToRemoteDebuggee = v
}

// SetConnectedToRemoteDebugger marks the state as connected to a remote
// debugger (this side is the debuggee).
func (k *KdState) SetConnectedToRemoteDebugger(v bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.isSerialConnectedToRemoteDebugger = v
}

// SetDebuggeeRunning records the running/paused state of the remote debuggee.
func (k *KdState) SetDebuggeeRunning(v bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.isDebuggeeRunning = v
}

// CurrentRunningInstruction returns the cached instruction bytes at the
// paused RIP and whether the debuggee was in 32-bit mode. Updated by the
// listener on each pause packet.
func (k *KdState) CurrentRunningInstruction() ([MaximumInstrSize]byte, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.currentRunningInstruction, k.isRunningInstruction32Bit
}

// SetCurrentRunningInstruction caches the instruction bytes at the paused RIP.
// Called by the listener.
func (k *KdState) SetCurrentRunningInstruction(instr [MaximumInstrSize]byte, is32Bit bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.currentRunningInstruction = instr
	k.isRunningInstruction32Bit = is32Bit
}

// AddBreakpoint records a kernel breakpoint in the local table.
func (k *KdState) AddBreakpoint(tag uint64, addr uint64) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.breakpoints[tag] = &BpInfo{Tag: tag, Address: addr, Enabled: true}
}

// RemoveBreakpoint removes a kernel breakpoint from the local table.
func (k *KdState) RemoveBreakpoint(tag uint64) {
	k.mu.Lock()
	defer k.mu.Unlock()
	delete(k.breakpoints, tag)
}

// Breakpoints returns a snapshot of the registered breakpoints.
func (k *KdState) Breakpoints() []BpInfo {
	k.mu.Lock()
	defer k.mu.Unlock()
	out := make([]BpInfo, 0, len(k.breakpoints))
	for _, bp := range k.breakpoints {
		out = append(out, *bp)
	}
	return out
}

// KernelBaseAddress returns the cached base address of ntoskrnl on the
// debuggee. Set by the listener from the DEBUGGEE_STARTED packet.
func (k *KdState) KernelBaseAddress() uint64 {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.kernelBaseAddress
}

// SetKernelBaseAddress caches the ntoskrnl base address.
func (k *KdState) SetKernelBaseAddress(base uint64) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.kernelBaseAddress = base
}

// ----------------------------------------------------------------------------
// Sync-object wait/signal primitives.
//
// Mirrors DbgWaitForKernelResponse / DbgReceivedKernelResponse from the C++
// side. Each slot in syncObjects has its own "done" channel; the waiter
// creates the channel and blocks on it, the signaller closes it.
// ----------------------------------------------------------------------------

// WaitForKernelResponse blocks until ReceivedKernelResponse is called for the
// given sync object. Mirrors DbgWaitForKernelResponse.
func (k *KdState) WaitForKernelResponse(obj KernelSyncObject) error {
	if int(obj) >= len(k.syncObjects) {
		return fmt.Errorf("WaitForKernelResponse: invalid sync object %d", obj)
	}
	ch := make(chan struct{})
	k.mu.Lock()
	if k.syncObjects[obj].IsOnWaitingState {
		// Slot already in use; return the existing channel so the second
		// caller also wakes up on the same signal (best-effort re-entrancy).
		ch = k.syncObjects[obj].done
		k.mu.Unlock()
	} else {
		k.syncObjects[obj].IsOnWaitingState = true
		k.syncObjects[obj].done = ch
		k.mu.Unlock()
	}
	<-ch
	return nil
}

// ReceivedKernelResponse signals the goroutine waiting on the given sync
// object. Mirrors DbgReceivedKernelResponse. No-op if nobody is waiting.
func (k *KdState) ReceivedKernelResponse(obj KernelSyncObject) {
	if int(obj) >= len(k.syncObjects) {
		return
	}
	k.mu.Lock()
	defer k.mu.Unlock()
	if !k.syncObjects[obj].IsOnWaitingState {
		return
	}
	ch := k.syncObjects[obj].done
	k.syncObjects[obj].IsOnWaitingState = false
	k.syncObjects[obj].done = nil
	if ch != nil {
		close(ch)
	}
}

// SetKernelRequestData stores the caller-supplied buffer for the request/
// response pair. Mirrors DbgWaitSetKernelRequestData. The Go version stores
// the buffer in a per-slot field so the listener can copy the response into
// it; we keep it inline in the sync object slot via a small map.
func (k *KdState) SetKernelRequestData(obj KernelSyncObject, buf []byte) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.requestBuffers == nil {
		k.requestBuffers = make(map[KernelSyncObject][]byte)
	}
	k.requestBuffers[obj] = buf
}

// GetKernelRequestData returns and clears the caller-supplied buffer for the
// given sync object. Mirrors DbgWaitGetKernelRequestData.
func (k *KdState) GetKernelRequestData(obj KernelSyncObject) ([]byte, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	buf, ok := k.requestBuffers[obj]
	if ok {
		delete(k.requestBuffers, obj)
	}
	return buf, ok
}

// requestBuffers is a small per-slot map for caller-supplied response buffers.
// It is allocated lazily.
//
// NOTE: declared as a field of KdState via the requestBuffers line below; Go
// allows fields after methods, we keep this near the sync code for clarity.
type requestBuffersMap = map[KernelSyncObject][]byte

// ----------------------------------------------------------------------------
// IOCTL helpers (debuggee side).
//
// These mirror the PlatformDeviceIoControl calls in kd.cpp. They run on the
// debuggee machine and talk to the local VMM driver through comm.Device.
// ----------------------------------------------------------------------------

// KdPrepareDebuggee sends IOCTL_PREPARE_DEBUGGEE to the local VMM driver,
// telling it the serial/named-pipe parameters and the ntoskrnl base address.
// Mirrors the IOCTL portion of KdPrepareAndConnectDebugPort (IsPreparing=true
// branch).
func (k *KdState) KdPrepareDebuggee(portAddress, baudrate, kernelBase uint32, osName string) error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdPrepareDebuggee: no device")
	}
	var req hyperdbgsdk.DEBUGGER_PREPARE_DEBUGGEE
	req.PortAddress = portAddress
	req.Baudrate = baudrate
	req.KernelBaseAddress = uint64(kernelBase)
	// OsName is [256]int8; copy the null-terminated string.
	copyInt8FromString(req.OsName[:], osName)

	inBuf := structAsBytes(unsafe.Pointer(&req), unsafe.Sizeof(req))
	outBuf := make([]byte, unsafe.Sizeof(req))
	copy(outBuf, inBuf)
	if _, err := dev.Ioctl(comm.IOCTL_CODE_PREPARE_DEBUGGEE, inBuf, outBuf); err != nil {
		return fmt.Errorf("KdPrepareDebuggee: IOCTL_PREPARE_DEBUGGEE failed: %w", err)
	}
	// Re-read Result from the output buffer.
	var resp hyperdbgsdk.DEBUGGER_PREPARE_DEBUGGEE
	bytesIntoStruct(unsafe.Pointer(&resp), outBuf, unsafe.Sizeof(resp))
	if resp.Result != DebuggerOperationWasSuccessful {
		return fmt.Errorf("KdPrepareDebuggee: kernel returned error 0x%x", resp.Result)
	}
	return nil
}

// KdSendUserInterfacePausePacket sends IOCTL_PAUSE_PACKET_RECEIVED to the
// local VMM driver. Mirrors the IOCTL call in the C++ pause path on the
// debuggee side.
func (k *KdState) KdSendUserInterfacePausePacket() error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdSendUserInterfacePausePacket: no device")
	}
	var pkt hyperdbgsdk.DEBUGGER_PAUSE_PACKET_RECEIVED
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if _, err := dev.Ioctl(comm.IOCTL_CODE_PAUSE_PACKET_RECEIVED, buf, buf); err != nil {
		return fmt.Errorf("KdSendUserInterfacePausePacket: IOCTL_PAUSE_PACKET_RECEIVED failed: %w", err)
	}
	return nil
}

// KdRegisterEventInDebuggee sends IOCTL_DEBUGGER_REGISTER_EVENT to the local
// VMM driver and forwards the result back to the remote debugger via
// KdSendGeneralBuffersFromDebuggeeToDebugger. Mirrors KdRegisterEventInDebuggee.
func (k *KdState) KdRegisterEventInDebuggee(eventBuf []byte) (hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error) {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT{}, fmt.Errorf("KdRegisterEventInDebuggee: no device")
	}
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf := structAsBytes(unsafe.Pointer(&result), unsafe.Sizeof(result))
	if _, err := dev.Ioctl(comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT, eventBuf, resultBuf); err != nil {
		return result, fmt.Errorf("KdRegisterEventInDebuggee: IOCTL failed: %w", err)
	}
	bytesIntoStruct(unsafe.Pointer(&result), resultBuf, unsafe.Sizeof(result))
	// Forward to the remote debugger.
	if err := k.KdSendGeneralBuffersFromDebuggeeToDebugger(
		hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfRegisteringEvent,
		resultBuf,
		true,
	); err != nil {
		return result, fmt.Errorf("KdRegisterEventInDebuggee: forward failed: %w", err)
	}
	return result, nil
}

// KdAddActionToEventInDebuggee sends IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT to
// the local VMM driver and forwards the result back to the remote debugger.
// Mirrors KdAddActionToEventInDebuggee.
func (k *KdState) KdAddActionToEventInDebuggee(actionBuf []byte) (hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error) {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT{}, fmt.Errorf("KdAddActionToEventInDebuggee: no device")
	}
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf := structAsBytes(unsafe.Pointer(&result), unsafe.Sizeof(result))
	if _, err := dev.Ioctl(comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT, actionBuf, resultBuf); err != nil {
		return result, fmt.Errorf("KdAddActionToEventInDebuggee: IOCTL failed: %w", err)
	}
	bytesIntoStruct(unsafe.Pointer(&result), resultBuf, unsafe.Sizeof(result))
	if err := k.KdSendGeneralBuffersFromDebuggeeToDebugger(
		hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfAddingActionToEvent,
		resultBuf,
		true,
	); err != nil {
		return result, fmt.Errorf("KdAddActionToEventInDebuggee: forward failed: %w", err)
	}
	return result, nil
}

// KdSendModifyEventInDebuggee sends IOCTL_DEBUGGER_MODIFY_EVENTS to the local
// VMM driver and optionally forwards the result back to the remote debugger.
// Mirrors KdSendModifyEventInDebuggee.
func (k *KdState) KdSendModifyEventInDebuggee(modify *hyperdbgsdk.DEBUGGER_MODIFY_EVENTS, sendResultToDebugger bool) error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdSendModifyEventInDebuggee: no device")
	}
	inBuf := structAsBytes(unsafe.Pointer(modify), unsafe.Sizeof(*modify))
	outBuf := make([]byte, unsafe.Sizeof(*modify))
	copy(outBuf, inBuf)
	if _, err := dev.Ioctl(comm.IOCTL_CODE_DEBUGGER_MODIFY_EVENTS, inBuf, outBuf); err != nil {
		return fmt.Errorf("KdSendModifyEventInDebuggee: IOCTL failed: %w", err)
	}
	bytesIntoStruct(unsafe.Pointer(modify), outBuf, unsafe.Sizeof(*modify))
	if sendResultToDebugger {
		if err := k.KdSendGeneralBuffersFromDebuggeeToDebugger(
			hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryAndModifyEvent,
			outBuf,
			true,
		); err != nil {
			return fmt.Errorf("KdSendModifyEventInDebuggee: forward failed: %w", err)
		}
	}
	return nil
}

// KdSendSignalExecutionFinished sends
// IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED to the local VMM driver.
// Mirrors the IOCTL call in KdHandleUserInputInDebuggee.
func (k *KdState) KdSendSignalExecutionFinished() error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdSendSignalExecutionFinished: no device")
	}
	var sig hyperdbgsdk.DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
	buf := structAsBytes(unsafe.Pointer(&sig), unsafe.Sizeof(sig))
	if _, err := dev.Ioctl(comm.IOCTL_CODE_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED, buf, buf); err != nil {
		return fmt.Errorf("KdSendSignalExecutionFinished: IOCTL failed: %w", err)
	}
	return nil
}

// KdSendUsermodePrints sends IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER to
// forward a chunk of user-mode output to the remote debugger. Mirrors
// KdSendUsermodePrints.
func (k *KdState) KdSendUsermodePrints(msg []byte) error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdSendUsermodePrints: no device")
	}
	// Build [DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER][msg]
	hdrSize := unsafe.Sizeof(hyperdbgsdk.DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER{})
	total := uint32(hdrSize) + uint32(len(msg))
	buf := make([]byte, total)
	// Header: KernelStatus (uint32) + Length (uint32). The driver fills in
	// KernelStatus; we set Length.
	binary.LittleEndian.PutUint32(buf[4:8], uint32(len(msg)))
	copy(buf[hdrSize:], msg)
	if _, err := dev.Ioctl(comm.IOCTL_CODE_SEND_USERMODE_MESSAGES_TO_DEBUGGER, buf, buf); err != nil {
		return fmt.Errorf("KdSendUsermodePrints: IOCTL failed: %w", err)
	}
	return nil
}

// KdSendGeneralBuffersFromDebuggeeToDebugger sends
// IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER to forward an
// arbitrary buffer from the debuggee to the debugger. Mirrors
// KdSendGeneralBuffersFromDebuggeeToDebugger.
func (k *KdState) KdSendGeneralBuffersFromDebuggeeToDebugger(
	action hyperdbgsdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION,
	buf []byte,
	pauseDebuggeeWhenSent bool,
) error {
	k.mu.Lock()
	dev := k.device
	k.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("KdSendGeneralBuffersFromDebuggeeToDebugger: no device")
	}
	hdrSize := unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER{})
	total := uint32(hdrSize) + uint32(len(buf))
	packet := make([]byte, total)
	// Fill header: RequestedAction (uint32) + LengthOfBuffer (uint32) +
	// PauseDebuggeeWhenSent (bool, 1 byte) + 3 bytes padding + KernelResult (uint32).
	binary.LittleEndian.PutUint32(packet[0:4], uint32(action))
	binary.LittleEndian.PutUint32(packet[4:8], uint32(len(buf)))
	if pauseDebuggeeWhenSent {
		packet[8] = 1
	}
	// KernelResult at offset 12 stays 0; the driver fills it in.
	copy(packet[hdrSize:], buf)
	outBuf := make([]byte, hdrSize) // driver returns just the header
	copy(outBuf, packet[:hdrSize])
	if _, err := dev.Ioctl(comm.IOCTL_CODE_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER, packet, outBuf); err != nil {
		return fmt.Errorf("KdSendGeneralBuffersFromDebuggeeToDebugger: IOCTL failed: %w", err)
	}
	// Check KernelResult at offset 12.
	kernelResult := binary.LittleEndian.Uint32(outBuf[12:16])
	if kernelResult != DebuggerOperationWasSuccessful {
		return fmt.Errorf("KdSendGeneralBuffersFromDebuggeeToDebugger: kernel returned 0x%x", kernelResult)
	}
	return nil
}

// ----------------------------------------------------------------------------
// High-level control flow: Continue / Pause / Step.
//
// These mirror the C++ control flow:
//   KdSendContinuePacketToDebuggee   -> KdCommandGoto
//   KdSendStepPacketToDebuggee       -> KdCommandStep
//   KdSendPausePacketToDebuggee      -> KdCommandPause
//   KdBreakControlCheckAndPauseDebugger / KdBreakControlCheckAndContinueDebugger
//   KdSetStatusAndWaitForPause
//   KdCloseConnection
// ----------------------------------------------------------------------------

// KdCommandGoto sends a 'g' (continue) packet to the debuggee. Mirrors
// KdSendContinuePacketToDebuggee.
func (k *KdState) KdCommandGoto() error {
	k.mu.Lock()
	k.currentRemoteCore = DebuggerDebuggeeIsRunningNoCore
	t := k.transport
	k.mu.Unlock()
	if t == nil {
		return fmt.Errorf("KdCommandGoto: no transport")
	}
	if err := k.KdCommandPacketToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootModeContinue,
	); err != nil {
		return err
	}
	return nil
}

// KdCommandStep sends a 't'/'p' (step-in / step-over) packet to the debuggee.
// Mirrors KdSendStepPacketToDebuggee. callInstrSize is the size of the current
// CALL instruction (0 if not a CALL); it is only meaningful for step-over.
func (k *KdState) KdCommandStep(stepType hyperdbgsdk.DEBUGGER_REMOTE_STEPPING_REQUEST, isCall bool, callInstrSize uint32) error {
	k.mu.Lock()
	t := k.transport
	k.mu.Unlock()
	if t == nil {
		return fmt.Errorf("KdCommandStep: no transport")
	}
	var pkt hyperdbgsdk.DEBUGGEE_STEP_PACKET
	pkt.StepType = stepType
	pkt.IsCurrentInstructionACall = isCall
	if isCall {
		pkt.CallLength = callInstrSize
	}
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	// Mark the debuggee as running; the step packet will resume it.
	k.mu.Lock()
	switch stepType {
	case hyperdbgsdk.DebuggerRemoteSteppingRequestStepIn,
		hyperdbgsdk.DebuggerRemoteSteppingRequestStepOver,
		hyperdbgsdk.DebuggerRemoteSteppingRequestStepOverForGu:
		k.isDebuggeeRunning = true
	}
	k.mu.Unlock()
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootModeStep,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectIsDebuggerRunning)
}

// KdCommandPause sends a PAUSE packet to the debuggee and then handles the
// paused state. Mirrors KdSendPausePacketToDebuggee.
func (k *KdState) KdCommandPause() error {
	if err := k.KdCommandPacketToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnUserMode,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnUserModePause,
	); err != nil {
		return err
	}
	// Handle the paused state (show rip, instructions, etc.) — the listener
	// goroutine will receive the paused packet and signal
	// SyncObjectPausedDebuggeeDetails. We mirror KdInterpretPausedDebuggee by
	// waiting on that sync object.
	return k.WaitForKernelResponse(SyncObjectPausedDebuggeeDetails)
}

// KdContinue is the high-level "go" entry point: it sends the continue packet
// and then waits for the debuggee to pause again (CTRL+C or breakpoint hit).
// Mirrors KdBreakControlCheckAndContinueDebugger + KdSetStatusAndWaitForPause.
func (k *KdState) KdContinue() error {
	k.mu.Lock()
	if k.isDebuggeeRunning {
		k.mu.Unlock()
		return nil // already running, nothing to do
	}
	k.mu.Unlock()
	if err := k.KdCommandGoto(); err != nil {
		k.printf("err, unable to continue the debuggee: %v\n", err)
		return err
	}
	return k.KdSetStatusAndWaitForPause()
}

// KdSetStatusAndWaitForPause marks the debuggee as running and blocks until
// the listener signals a pause. Mirrors KdSetStatusAndWaitForPause +
// KdTheRemoteSystemIsRunning.
func (k *KdState) KdSetStatusAndWaitForPause() error {
	k.mu.Lock()
	k.isDebuggeeRunning = true
	k.mu.Unlock()
	k.printf("debuggee is running...\n")
	return k.WaitForKernelResponse(SyncObjectIsDebuggerRunning)
}

// KdBreakControlCheckAndPauseDebugger sends a pause request to the debuggee
// if it is currently running. Mirrors KdBreakControlCheckAndPauseDebugger.
// signalRunningFlag controls whether the IsDebuggerRunning sync object is
// signalled (used by the post-connection pause path).
func (k *KdState) KdBreakControlCheckAndPauseDebugger(signalRunningFlag bool) error {
	k.mu.Lock()
	running := k.isDebuggeeRunning
	ignore := k.ignorePauseRequests
	k.mu.Unlock()
	if ignore {
		return nil
	}
	if !running {
		return nil
	}
	if err := k.KdCommandPause(); err != nil {
		k.printf("err, unable to pause the debuggee: %v\n", err)
		return err
	}
	if signalRunningFlag {
		k.ReceivedKernelResponse(SyncObjectIsDebuggerRunning)
	}
	return nil
}

// KdBreakControlCheckAndContinueDebugger sends a continue request if the
// debuggee is currently paused. Mirrors KdBreakControlCheckAndContinueDebugger.
func (k *KdState) KdBreakControlCheckAndContinueDebugger() error {
	k.mu.Lock()
	running := k.isDebuggeeRunning
	k.mu.Unlock()
	if running {
		return nil
	}
	if err := k.KdCommandGoto(); err != nil {
		k.printf("err, unable to continue the debuggee: %v\n", err)
		return err
	}
	return k.KdSetStatusAndWaitForPause()
}

// KdCloseConnection closes the connection in both debuggee and debugger
// roles. Mirrors KdCloseConnection. Safe to call multiple times.
func (k *KdState) KdCloseConnection() error {
	k.mu.Lock()
	if k.serialConnectionAlreadyClosed {
		k.mu.Unlock()
		return nil
	}
	k.serialConnectionAlreadyClosed = true
	isDebuggee := k.isSerialConnectedToRemoteDebugger
	isDebugger := k.isSerialConnectedToRemoteDebuggee
	t := k.transport
	k.mu.Unlock()

	if isDebuggee {
		// This side is the debuggee; nothing more to do here, the local VMM
		// driver is unloaded by the caller.
	} else if isDebugger {
		// This side is the debugger; send the close-and-unload packet to the
		// debuggee.
		if t != nil {
			_ = k.KdCommandPacketToDebuggee(
				hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
				hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootModeCloseAndUnloadDebuggee,
			)
			k.printf("unloading debugger vmm module on debuggee...")
			_ = k.KdCommandPacketToDebuggee(
				hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnUserMode,
				hyperdbgsdk.DebuggerRemotePacketRequestedActionOnUserModeDoNotReadAnyPacket,
			)
		}
	} else {
		k.printf("err, start packet not received but the target machine closed the connection\n")
		k.ReceivedKernelResponse(SyncObjectStartedPacketReceived)
	}

	k.KdUninitializeConnection()
	return nil
}

// KdUninitializeConnection resets all connection state and releases the
// transport. Mirrors KdUninitializeConnection.
func (k *KdState) KdUninitializeConnection() {
	k.mu.Lock()
	t := k.transport
	k.transport = nil
	k.isSerialConnectedToRemoteDebugger = false
	k.isSerialConnectedToRemoteDebuggee = false
	k.isDebuggeeRunning = false
	k.isDebuggeeInHandshakingPhase = false
	k.isDebuggerConnectedToNamedPipe = false
	k.shouldPreviousCommandBeContinued = false
	k.ignoreNewLoggingMessages = false
	k.currentRemoteCore = DebuggerDebuggeeIsRunningNoCore
	// Wake up any goroutines still waiting on sync objects.
	for i := range k.syncObjects {
		if k.syncObjects[i].IsOnWaitingState {
			ch := k.syncObjects[i].done
			k.syncObjects[i].IsOnWaitingState = false
			k.syncObjects[i].done = nil
			if ch != nil {
				close(ch)
			}
		}
	}
	k.breakpoints = make(map[uint64]*BpInfo)
	k.mu.Unlock()

	if t != nil {
		_ = t.Close()
	}
	// Best-effort wake-up of the IsDebuggerRunning waiter so callers blocked
	// on KdSetStatusAndWaitForPause return after a close.
	k.ReceivedKernelResponse(SyncObjectIsDebuggerRunning)
}

// ----------------------------------------------------------------------------
// Packet framing helpers (serial/named-pipe transport).
//
// Mirror KdComputeDataChecksum, KdCheckForTheEndOfTheBuffer,
// KdSendPacketToDebuggee, KdCommandPacketToDebuggee,
// KdCommandPacketAndBufferToDebuggee.
// ----------------------------------------------------------------------------

// KdComputeDataChecksum computes the per-byte sum of buf. Mirrors
// KdComputeDataChecksum.
func KdComputeDataChecksum(buf []byte) uint8 {
	var sum uint8
	for _, b := range buf {
		sum += b
	}
	return sum
}

// KdCheckForTheEndOfTheBuffer reports whether the last 4 bytes of buf form the
// SERIAL_END_OF_BUFFER_CHAR_{1,2,3,4} signature; if so, the trailing 4 bytes
// are stripped from buf and true is returned. Mirrors
// KdCheckForTheEndOfTheBuffer.
func KdCheckForTheEndOfTheBuffer(buf *[]byte) bool {
	if len(*buf) < 4 {
		return false
	}
	n := len(*buf)
	if (*buf)[n-1] == SerialEndOfBufferChar4 &&
		(*buf)[n-2] == SerialEndOfBufferChar3 &&
		(*buf)[n-3] == SerialEndOfBufferChar2 &&
		(*buf)[n-4] == SerialEndOfBufferChar1 {
		*buf = (*buf)[:n-4]
		return true
	}
	return false
}

// KdSendPacketToDebuggee writes buf to the transport and optionally appends
// the -byte end-of-buffer marker. Mirrors KdSendPacketToDebuggee.
func (k *KdState) KdSendPacketToDebuggee(buf []byte, sendEndOfBuffer bool) error {
	k.mu.Lock()
	t := k.transport
	k.ignoreNewLoggingMessages = false
	k.mu.Unlock()
	if t == nil {
		return fmt.Errorf("KdSendPacketToDebuggee: no transport")
	}
	if uint32(len(buf))+SerialEndOfBufferCharsCount > MaxSerialPacketSize {
		return fmt.Errorf("KdSendPacketToDebuggee: buffer too large (%d > %d)", len(buf), MaxSerialPacketSize)
	}
	if err := t.Write(buf); err != nil {
		return fmt.Errorf("KdSendPacketToDebuggee: write failed: %w", err)
	}
	if sendEndOfBuffer {
		marker := []byte{SerialEndOfBufferChar1, SerialEndOfBufferChar2, SerialEndOfBufferChar3, SerialEndOfBufferChar4}
		if err := t.Write(marker); err != nil {
			return fmt.Errorf("KdSendPacketToDebuggee: end-of-buffer write failed: %w", err)
		}
	}
	return nil
}

// KdCommandPacketToDebuggee builds and sends a DEBUGGER_REMOTE_PACKET with no
// additional payload. Mirrors KdCommandPacketToDebuggee.
func (k *KdState) KdCommandPacketToDebuggee(
	packetType hyperdbgsdk.DEBUGGER_REMOTE_PACKET_TYPE,
	action hyperdbgsdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION,
) error {
	var pkt hyperdbgsdk.DEBUGGER_REMOTE_PACKET
	pkt.Indicator = IndicatorOfHyperdbgPacket
	pkt.TypeOfThePacket = packetType
	pkt.RequestedActionOfThePacket = action
	// Checksum is computed over the packet bytes excluding the leading
	// Checksum byte.
	pktBytes := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	pkt.Checksum = KdComputeDataChecksum(pktBytes[1:])
	return k.KdSendPacketToDebuggee(pktBytes, true)
}

// KdCommandPacketAndBufferToDebuggee builds and sends a DEBUGGER_REMOTE_PACKET
// followed by an additional payload buffer. Mirrors
// KdCommandPacketAndBufferToDebuggee.
func (k *KdState) KdCommandPacketAndBufferToDebuggee(
	packetType hyperdbgsdk.DEBUGGER_REMOTE_PACKET_TYPE,
	action hyperdbgsdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION,
	payload []byte,
) error {
	pktSize := unsafe.Sizeof(hyperdbgsdk.DEBUGGER_REMOTE_PACKET{})
	if uint32(pktSize)+uint32(len(payload))+SerialEndOfBufferCharsCount > MaxSerialPacketSize {
		return fmt.Errorf("KdCommandPacketAndBufferToDebuggee: buffer too large")
	}
	var pkt hyperdbgsdk.DEBUGGER_REMOTE_PACKET
	pkt.Indicator = IndicatorOfHyperdbgPacket
	pkt.TypeOfThePacket = packetType
	pkt.RequestedActionOfThePacket = action
	pktBytes := structAsBytes(unsafe.Pointer(&pkt), pktSize)
	pkt.Checksum = KdComputeDataChecksum(pktBytes[1:])
	pkt.Checksum += KdComputeDataChecksum(payload)
	// Send the packet header without end-of-buffer...
	if err := k.KdSendPacketToDebuggee(pktBytes, false); err != nil {
		return err
	}
	// ...then the payload with end-of-buffer.
	return k.KdSendPacketToDebuggee(payload, true)
}

// ----------------------------------------------------------------------------
// Breakpoint / event management (high-level wrappers).
// ----------------------------------------------------------------------------

// KdSendBpPacket sends a 'bp' breakpoint packet to the debuggee and waits for
// the result. Mirrors KdSendBpPacketToDebuggee.
func (k *KdState) KdSendBpPacket(bp *hyperdbgsdk.DEBUGGEE_BP_PACKET) error {
	buf := structAsBytes(unsafe.Pointer(bp), unsafe.Sizeof(*bp))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootBp,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectBp)
}

// KdSendListOrModifyBpPacket sends a 'bc'/'bd'/'be'/'bl' packet to the
// debuggee. Mirrors KdSendListOrModifyPacketToDebuggee.
func (k *KdState) KdSendListOrModifyBpPacket(pkt *hyperdbgsdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET) error {
	buf := structAsBytes(unsafe.Pointer(pkt), unsafe.Sizeof(*pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootListOrModifyBreakpoints,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectListOrModifyBreakpoints)
}

// KdSendFlushPacket sends a 'flush' packet to the debuggee. Mirrors
// KdSendFlushPacketToDebuggee.
func (k *KdState) KdSendFlushPacket() error {
	var pkt hyperdbgsdk.DEBUGGER_FLUSH_LOGGING_BUFFERS
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootModeFlushBuffers,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectFlushResult)
}

// KdSendSwitchCorePacket sends a '~' (switch core) packet to the debuggee.
// Mirrors KdSendSwitchCorePacketToDebuggee.
func (k *KdState) KdSendSwitchCorePacket(newCore uint32) error {
	k.mu.Lock()
	current := k.currentRemoteCore
	k.mu.Unlock()
	if newCore == current {
		k.printf("the current operating core is %x (not changed)\n", newCore)
		return fmt.Errorf("KdSendSwitchCorePacket: core unchanged")
	}
	var pkt hyperdbgsdk.DEBUGGEE_CHANGE_CORE_PACKET
	pkt.NewCore = newCore
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootModeChangeCore,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectCoreSwitchingResult)
}

// KdSendShortCircuitingEvent sends a short-circuiting enable/disable packet
// to the debuggee. Mirrors KdSendShortCircuitingEventToDebuggee.
func (k *KdState) KdSendShortCircuitingEvent(isEnabled bool) error {
	var pkt hyperdbgsdk.DEBUGGER_SHORT_CIRCUITING_EVENT
	pkt.IsShortCircuiting = isEnabled
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootSetShortCircuitingState,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectShortCircuitingEventState)
}

// KdSendEventQueryAndModifyPacket sends a query/modify event packet to the
// debuggee and returns the queried state (only meaningful for query actions).
// Mirrors KdSendEventQueryAndModifyPacketToDebuggee.
func (k *KdState) KdSendEventQueryAndModifyPacket(
	tag uint64,
	actionType hyperdbgsdk.DEBUGGER_MODIFY_EVENTS_TYPE,
) (bool, error) {
	k.mu.Lock()
	k.sharedEventStatus = false
	k.mu.Unlock()
	var pkt hyperdbgsdk.DEBUGGER_MODIFY_EVENTS
	pkt.Tag = tag
	pkt.TypeOfAction = actionType
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootQueryAndModifyEvent,
		buf,
	); err != nil {
		return false, err
	}
	if err := k.WaitForKernelResponse(SyncObjectModifyAndQueryEvent); err != nil {
		return false, err
	}
	k.mu.Lock()
	defer k.mu.Unlock()
	if actionType == hyperdbgsdk.DebuggerModifyEventsQueryState {
		return k.sharedEventStatus, nil
	}
	return false, nil
}

// KdSendUserInputPacket sends a user-input command string to the debuggee.
// Mirrors KdSendUserInputPacketToDebuggee. ignoreBreakingAgain corresponds to
// the IgnoreFinishedSignal field of DEBUGGEE_USER_INPUT_PACKET.
func (k *KdState) KdSendUserInputPacket(input string, ignoreBreakingAgain bool) error {
	var hdr hyperdbgsdk.DEBUGGEE_USER_INPUT_PACKET
	hdr.CommandLen = uint32(len(input))
	hdr.IgnoreFinishedSignal = ignoreBreakingAgain
	hdrSize := unsafe.Sizeof(hdr)
	total := uint32(hdrSize) + uint32(len(input))
	buf := make([]byte, total)
	copy(buf[:hdrSize], structAsBytes(unsafe.Pointer(&hdr), hdrSize))
	copy(buf[hdrSize:], []byte(input))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootUserInputBuffer,
		buf,
	); err != nil {
		return err
	}
	if !ignoreBreakingAgain {
		return k.WaitForKernelResponse(SyncObjectDebuggeeFinishedCommandExecution)
	}
	return nil
}

// KdSendScriptPacket sends a script buffer to the debuggee. Mirrors
// KdSendScriptPacketToDebuggee. isFormat distinguishes the '.formats' path
// which waits on an extra sync object.
func (k *KdState) KdSendScriptPacket(scriptBuf []byte, pointer uint32, isFormat bool) error {
	var hdr hyperdbgsdk.DEBUGGEE_SCRIPT_PACKET
	hdr.ScriptBufferSize = uint32(len(scriptBuf))
	hdr.ScriptBufferPointer = pointer
	hdr.IsFormat = isFormat
	hdrSize := unsafe.Sizeof(hdr)
	total := uint32(hdrSize) + uint32(len(scriptBuf))
	buf := make([]byte, total)
	copy(buf[:hdrSize], structAsBytes(unsafe.Pointer(&hdr), hdrSize))
	copy(buf[hdrSize:], scriptBuf)
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootRunScript,
		buf,
	); err != nil {
		return err
	}
	if isFormat {
		if err := k.WaitForKernelResponse(SyncObjectScriptFormatsResult); err != nil {
			return err
		}
	}
	return k.WaitForKernelResponse(SyncObjectScriptRunningResult)
}

// KdSendSymbolReloadPacket sends a '.sym reload' packet to the debuggee.
// Mirrors KdSendSymbolReloadPacketToDebuggee.
func (k *KdState) KdSendSymbolReloadPacket(processId uint32) error {
	var pkt hyperdbgsdk.DEBUGGEE_SYMBOL_REQUEST_PACKET
	pkt.ProcessId = processId
	buf := structAsBytes(unsafe.Pointer(&pkt), unsafe.Sizeof(pkt))
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnVmxRootSymbolReload,
		buf,
	); err != nil {
		return err
	}
	return k.WaitForKernelResponse(SyncObjectSymbolReload)
}

// SetSharedEventStatus stores the result of a query-state event action.
// Called by the listener when DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT is
// received.
func (k *KdState) SetSharedEventStatus(v bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.sharedEventStatus = v
}

// SetResultOfRegisteringEvent stores the result buffer of a register-event
// action. Called by the listener.
func (k *KdState) SetResultOfRegisteringEvent(r hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.resultOfRegisteringEvent = r
}

// ResultOfRegisteringEvent returns the cached result of the last
// register-event action.
func (k *KdState) ResultOfRegisteringEvent() hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.resultOfRegisteringEvent
}

// SetResultOfAddingActions stores the result buffer of an add-action-to-event
// action. Called by the listener.
func (k *KdState) SetResultOfAddingActions(r hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.resultOfAddingActionsTo = r
}

// ResultOfAddingActions returns the cached result of the last
// add-action-to-event action.
func (k *KdState) ResultOfAddingActions() hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.resultOfAddingActionsTo
}

// SetFormatsResult stores the result of a '.formats' evaluation.
// Called by the listener.
func (k *KdState) SetFormatsResult(value uint64, errState uint32) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.resultOfEvaluatedExpression = value
	k.errorStateOfEvaluatedExpr = errState
}

// FormatsResult returns the cached result of the last '.formats' evaluation.
func (k *KdState) FormatsResult() (value uint64, errState uint32) {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.resultOfEvaluatedExpression, k.errorStateOfEvaluatedExpr
}

// ----------------------------------------------------------------------------
// Internal helpers.
// ----------------------------------------------------------------------------

// printf writes to the output sink if one is configured. Best-effort: errors
// are silently dropped (the C++ ShowMessages has no error path either).
func (k *KdState) printf(format string, args ...any) {
	k.mu.Lock()
	out := k.out
	k.mu.Unlock()
	if out == nil {
		return
	}
	_ = out.Printf(format, args...)
}

// structAsBytes returns a byte slice aliasing the memory at ptr for size
// bytes. The result is only valid while the source memory stays live.
func structAsBytes(ptr unsafe.Pointer, size uintptr) []byte {
	return (*[1 << 30]byte)(ptr)[:size:size]
}

// bytesIntoStruct copies len(src) bytes (capped at size) into the memory at
// ptr. Mirrors a memcpy into a typed struct.
func bytesIntoStruct(ptr unsafe.Pointer, src []byte, size uintptr) {
	n := uintptr(len(src))
	if n > size {
		n = size
	}
	dst := (*[1 << 30]byte)(ptr)[:size:size]
	copy(dst, src[:n])
}

// copyInt8FromString copies s (truncated, NUL-padded) into a []int8 buffer.
// Used to fill the OsName field of DEBUGGER_PREPARE_DEBUGGEE.
func copyInt8FromString(dst []int8, s string) {
	for i := range dst {
		dst[i] = 0
	}
	n := len(s)
	if n > len(dst)-1 {
		n = len(dst) - 1
	}
	for i := 0; i < n; i++ {
		dst[i] = int8(s[i])
	}
}

// requestBuffersMap is the per-slot map type used by SetKernelRequestData /
// GetKernelRequestData to ferry caller-supplied response buffers from the
// IOCTL-issuing goroutine to the listener goroutine. Allocated lazily inside
// KdState.requestBuffers.
type _ = requestBuffersMap

// HyperdbgBuildSignature mirrors the C++ BuildSignature byte array (a
// "M.m.p-YYYYMMDD.HHMM" string built from the VERSION_* / __DATE__ / __TIME__
// macros). The Go version reports the same shape so the debuggee cannot
// distinguish a Go-built debugger from a C++-built one at the handshake.
var HyperdbgBuildSignature = []byte("0.22.0-20260101.0000\x00")

// KdSendResponseOfThePingPacket answers a ping from the debuggee by sending
// back the build signature. Mirrors KdSendResponseOfThePingPacket in kd.cpp.
// Called by the listener when DEBUGGER_REMOTE_PACKET_PING_AND_SEND_SUPPORTED_VERSION
// is received.
func (k *KdState) KdSendResponseOfThePingPacket() error {
	sig := HyperdbgBuildSignature
	if err := k.KdCommandPacketAndBufferToDebuggee(
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnUserMode,
		hyperdbgsdk.DebuggerRemotePacketRequestedActionOnUserModeDebuggerVersion,
		sig,
	); err != nil {
		k.printf("err, unable to send response to the ping packet: %v\n", err)
		return err
	}
	return nil
}
