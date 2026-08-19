// Package kernellvl — kernel_listening.go
//
// Implements the kernel-mode debugger event listener. The C++ counterpart is
// libhyperdbg/code/debugger/kernel-level/kernel-listening.cpp: a single
// ListeningSerialPortInDebugger function that loops reading packets from the
// serial/named-pipe transport, validates indicator+checksum+type, and switches
// on RequestedActionOfThePacket to update global state, print user-visible
// messages, and signal sync objects so the waiting IOCTL caller resumes.
//
// In the Go rewrite the listener is a goroutine that consumes
// KernelPausedPacket items from a channel. The producer (a transport-reading
// goroutine) owns the wire I/O; this separation lets GUI/MCP hosts plug their
// own transport (named pipe, TCP, mock) without dragging serial-port specifics
// into the listener. The listener:
//   - validates indicator + checksum + packet type
//   - updates KdState (current core, instruction bytes, paused/running flags,
//     result buffers for register-event / add-action / formats / etc.)
//   - calls the user-supplied KernelHandler for every user-visible pause so
//     CLI/GUI/MCP can render the pause differently
//   - signals the sync object the waiting command is blocked on
//
// All IOCTL-issued commands (KdSendBpPacket, KdSendFlushPacket, ...) block on
// WaitForKernelResponse(obj); the listener wakes them up by calling
// ReceivedKernelResponse(obj) at the end of each case arm, mirroring the C++
// DbgReceivedKernelResponse calls.
package kernellvl

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
)

// KernelEvent categorises why the debuggee paused. It is derived from the
// DEBUGGEE_PAUSING_REASON field of DEBUGGEE_KD_PAUSED_PACKET so the handler
// can switch on a stable Go enum instead of the wire numeric. Mirrors the
// relevant subset of DEBUGGEE_PAUSING_REASON_DEBUGGEE_*.
type KernelEvent uint32

const (
	// KernelEventUnknown is the zero value; used when the pausing reason does
	// not map to a more specific event.
	KernelEventUnknown KernelEvent = iota
	// KernelEventPaused is a user-triggered pause (CTRL+C or 'pause' command).
	KernelEventPaused
	// KernelEventRequestFromDebugger is a pause triggered by a debugger request
	// that needs the debuggee to drop into the interactive prompt.
	KernelEventRequestFromDebugger
	// KernelEventSingleStep is the result of a 't'/'p' step command.
	KernelEventSingleStep
	// KernelEventTrackingStep is the result of a 'tt' tracking step.
	KernelEventTrackingStep
	// KernelEventBreakpointHit is a software breakpoint hit.
	KernelEventBreakpointHit
	// KernelEventHardwareDebugRegisterHit is a hardware debug-register hit.
	KernelEventHardwareDebugRegisterHit
	// KernelEventCoreSwitched is the result of a '~' core switch.
	KernelEventCoreSwitched
	// KernelEventProcessSwitched is the result of a process switch.
	KernelEventProcessSwitched
	// KernelEventThreadSwitched is the result of a thread switch.
	KernelEventThreadSwitched
	// KernelEventCommandFinished is the result of a remote command finishing.
	KernelEventCommandFinished
	// KernelEventTriggered is an event (hook) triggered; the handler should
	// consult pkt.EventCallingStage for pre/post emulation.
	KernelEventTriggered
	// KernelEventStartingModuleLoaded is the initial module-load pause before
	// the entry-point breakpoint is set.
	KernelEventStartingModuleLoaded
	// KernelEventGeneralDebugBreak is a generic debug break.
	KernelEventGeneralDebugBreak
	// KernelEventGeneralThreadIntercepted is a thread intercept pause.
	KernelEventGeneralThreadIntercepted
	// KernelEventHardwareBasedGeneralBreak is a hardware-based debug break.
	KernelEventHardwareBasedGeneralBreak
)

// fromPausingReason maps a wire DEBUGGEE_PAUSING_REASON to a KernelEvent.
func fromPausingReason(r hyperdbgsdk.DEBUGGEE_PAUSING_REASON) KernelEvent {
	switch r {
	case hyperdbgsdk.DebuggeePausingReasonPause:
		return KernelEventPaused
	case hyperdbgsdk.DebuggeePausingReasonRequestFromDebugger:
		return KernelEventRequestFromDebugger
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeStepped:
		return KernelEventSingleStep
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeTrackingStepped:
		return KernelEventTrackingStep
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeSoftwareBreakpointHit:
		return KernelEventBreakpointHit
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeHardwareDebugRegisterHit:
		return KernelEventHardwareDebugRegisterHit
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeCoreSwitched:
		return KernelEventCoreSwitched
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeProcessSwitched:
		return KernelEventProcessSwitched
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeThreadSwitched:
		return KernelEventThreadSwitched
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeCommandExecutionFinished:
		return KernelEventCommandFinished
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeEventTriggered:
		return KernelEventTriggered
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeStartingModuleLoaded:
		return KernelEventStartingModuleLoaded
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeGeneralDebugBreak:
		return KernelEventGeneralDebugBreak
	case hyperdbgsdk.DebuggeePausingReasonDebuggeeGeneralThreadIntercepted:
		return KernelEventGeneralThreadIntercepted
	case hyperdbgsdk.DebuggeePausingReasonHardwareBasedDebuggeeGeneralBreak:
		return KernelEventHardwareBasedGeneralBreak
	default:
		return KernelEventUnknown
	}
}

// KernelPausedPacket is the unit the listener consumes. The producer (typically
// a transport-reading goroutine) parses the wire bytes into Header + Payload
// and pushes the result onto the channel.
//
// Payload holds the bytes that follow the DEBUGGER_REMOTE_PACKET header on the
// wire; the listener reinterprets them as the appropriate sub-struct per the
// RequestedAction in Header. Payload may be empty for actions that carry no
// additional data.
type KernelPausedPacket struct {
	Header  hyperdbgsdk.DEBUGGER_REMOTE_PACKET
	Payload []byte
}

// KernelHandler is invoked by the listener for every user-visible pause (i.e.
// when the debuggee is halted waiting for the next command). It runs after the
// listener has updated KdState (current core, instruction, paused flag) and
// before the listener signals the sync object that resumes the waiting
// command. Implementations produce the user-visible output:
//   - CLI: print the disassembly + breakpoint/event context
//   - GUI: highlight the current line, refresh register/memory views
//   - MCP: emit a JSON event to the AI agent
//
// Returning an error stops the listener.
type KernelHandler func(ctx context.Context, state *KdState, ev KernelEvent, pkt hyperdbgsdk.DEBUGGEE_KD_PAUSED_PACKET) error

// KernelListener dispatches packets received from the debuggee. It owns no
// transport — packets are pushed onto the channel by a separate reader
// goroutine. This separation lets GUI/MCP hosts plug their own wire (named
// pipe, TCP, mock) without dragging in serial-port I/O.
//
// A KernelListener is goroutine-safe for SetHandler; Run must be called from a
// single goroutine.
type KernelListener struct {
	mu      sync.Mutex
	state   *KdState
	handler KernelHandler
}

// NewKernelListener creates a listener for the given state. handler may be nil;
// in that case paused packets are processed silently (the sync objects are
// still signalled so waiting commands resume).
func NewKernelListener(state *KdState, handler KernelHandler) *KernelListener {
	return &KernelListener{state: state, handler: handler}
}

// SetHandler replaces the handler at runtime. Safe to call concurrently with
// Run; the new handler takes effect on the next packet.
func (l *KernelListener) SetHandler(h KernelHandler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handler = h
}

// Run consumes packets from ch until ctx is cancelled or ch is closed. For
// each packet it validates the indicator/checksum/type and dispatches on the
// RequestedAction. Mirrors ListeningSerialPortInDebugger in
// kernel-listening.cpp.
//
// Transient errors (bad checksum, unknown action) are logged via the state's
// Output and dropped, mirroring the C++ 'goto StartAgain' behaviour. An error
// is returned only if ctx is cancelled or the handler returns one.
func (l *KernelListener) Run(ctx context.Context, ch <-chan KernelPausedPacket) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case pkt, ok := <-ch:
			if !ok {
				return nil
			}
			if err := l.dispatch(ctx, pkt); err != nil {
				return err
			}
		}
	}
}

// dispatch validates the packet and switches on RequestedAction. Returning an
// error stops the listener; the C++ code never does this for packet-level
// issues (it just logs and continues), so dispatch only returns an error when
// the user-supplied KernelHandler does.
func (l *KernelListener) dispatch(ctx context.Context, p KernelPausedPacket) error {
	if p.Header.Indicator != IndicatorOfHyperdbgPacket {
		l.printf("err, invalid packet received\n")
		return nil
	}
	if !l.validateChecksum(p) {
		l.printf("err, checksum is invalid\n")
		return nil
	}
	if p.Header.TypeOfThePacket != hyperdbgsdk.DebuggerRemotePacketTypeDebuggeeToDebugger {
		l.printf("err, unknown packet received from the debuggee\n")
		return nil
	}

	switch p.Header.RequestedActionOfThePacket {
	case hyperdbgsdk.DebuggerRemotePacketPingAndSendSupportedVersion:
		if err := l.state.KdSendResponseOfThePingPacket(ctx); err != nil {
			l.printf("err, failed to send ping response: %v\n", err)
		}

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeStarted:
		l.handleDebuggeeStarted(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeLoggingMechanism:
		l.handleLoggingMechanism(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction:
		return l.handlePausedAndCurrentInstruction(ctx, p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingCore:
		l.handleChangeCoreResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingProcess:
		l.handleChangeProcessResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingThread:
		l.handleChangeThreadResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeReloadSearchQuery:
		l.handleSearchQueryResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfFlush:
		l.handleFlushResult(p)

	// NOTE: DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_USER_CPUID
	// is not yet modelled in go-libhyperdbg/types/sdk.go (the auto-generated
	// enum is missing it). When the types package is regenerated to include
	// it, add a `case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfUserCpuid:
	// l.handleCpuidResult(p)` arm here. Until then, the CPUID result packet
	// falls through to default and is logged as "unknown packet action".

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfCallstack:
		l.handleCallstackResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultTestQuery:
		l.handleTestQueryResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfRunningScript:
		l.handleScriptResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfFormats:
		l.handleFormatsResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfRegisteringEvent:
		l.handleRegisterEventResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfAddingActionToEvent:
		l.handleAddActionResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryAndModifyEvent:
		l.handleQueryAndModifyEventResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeReloadSymbolFinished:
		l.handleSymbolReloadFinished(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingRegisters:
		l.handleReadRegistersResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfWriteRegister:
		l.handleWriteRegisterResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfApicRequests:
		l.handleCallerBufferResult(p, SyncObjectApicActions)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryIdtEntriesRequests:
		l.handleCallerBufferResult(p, SyncObjectIdtEntries)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingMemory:
		l.handleCallerBufferResult(p, SyncObjectReadMemory)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfEditingMemory:
		l.handleCallerBufferResult(p, SyncObjectEditMemory)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfBp:
		l.handleBpResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingState:
		l.handleShortCircuitingResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPte:
		l.handlePteResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultSmiOperationRequests:
		l.handleCallerBufferResult(p, SyncObjectSmiOperationResult)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultHypertraceLbrOperationRequests:
		l.handleCallerBufferResult(p, SyncObjectHypertraceLbrDumpResult)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultHypertracePtOperationRequests:
		l.handleCallerBufferResult(p, SyncObjectHypertracePtOperationResult)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfBringingPagesIn:
		l.handlePageInResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfVa2paAndPa2va:
		l.handleVa2paPa2vaResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfListOrModifyBreakpoints:
		l.handleListOrModifyBpResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeUpdateSymbolInfo:
		// Symbol-table updates are a no-op in the Go version: the symbol
		// parser is a separate package and does not maintain a remote symbol
		// table cache here. The C++ SymbolBuildAndUpdateSymbolTable call is
		// only meaningful when the in-process symbol table is active.

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPcitree:
		l.handlePcitreeResult(p)

	case hyperdbgsdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPcidevinfo:
		l.handlePcidevinfoResult(p)

	default:
		l.printf("err, unknown packet action received from the debugger\n")
	}
	return nil
}

// validateChecksum recomputes the per-byte sum over everything except the
// leading Checksum byte and compares it to the Checksum field. Mirrors the
// KdComputeDataChecksum call in ListeningSerialPortInDebugger.
func (l *KernelListener) validateChecksum(p KernelPausedPacket) bool {
	hdrBytes := structAsBytes(unsafe.Pointer(&p.Header), unsafe.Sizeof(p.Header))
	if len(hdrBytes) < 1 {
		return false
	}
	var sum uint8
	for _, b := range hdrBytes[1:] { // skip Checksum byte
		sum += b
	}
	for _, b := range p.Payload {
		sum += b
	}
	return sum == p.Header.Checksum
}

// ----------------------------------------------------------------------------
// Per-action handlers. Each mirrors one case arm in the C++ switch.
// ----------------------------------------------------------------------------

// handleDebuggeeStarted mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED:
// parse the DEBUGGER_PREPARE_DEBUGGEE payload, cache the kernel base address,
// print the "connected to debuggee" message, and signal
// SyncObjectStartedPacketReceived.
func (l *KernelListener) handleDebuggeeStarted(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_PREPARE_DEBUGGEE{})) {
		l.printf("err, debuggee-started packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectStartedPacketReceived)
		return
	}
	var init hyperdbgsdk.DEBUGGER_PREPARE_DEBUGGEE
	bytesIntoStruct(unsafe.Pointer(&init), p.Payload, unsafe.Sizeof(init))
	l.state.SetKernelBaseAddress(init.KernelBaseAddress)
	l.printf("connected to debuggee %s\n", int8String(init.OsName[:]))
	l.state.ReceivedKernelResponse(SyncObjectStartedPacketReceived)
}

// handleLoggingMechanism mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM:
// parse the DEBUGGEE_MESSAGE_PACKET payload and print the message (unless the
// pause-flag has muted logging).
func (l *KernelListener) handleLoggingMechanism(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_MESSAGE_PACKET{})) {
		l.printf("err, logging-mechanism packet too small\n")
		return
	}
	var msg hyperdbgsdk.DEBUGGEE_MESSAGE_PACKET
	bytesIntoStruct(unsafe.Pointer(&msg), p.Payload, unsafe.Sizeof(msg))
	// g_IgnoreNewLoggingMessages is checked on the C++ side; we route through
	// KdState so the listener stays in sync with pause/continue transitions.
	if l.state.IsIgnoreNewLoggingMessages() {
		return
	}
	l.printf("%s", int8String(msg.Message[:]))
}

// handlePausedAndCurrentInstruction mirrors the big
// DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION case:
// mark the debuggee paused, save the current core + instruction, print the
// pre-disassembly context, call the user handler, then signal the sync object
// that matches the pausing reason.
func (l *KernelListener) handlePausedAndCurrentInstruction(ctx context.Context, p KernelPausedPacket) error {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_KD_PAUSED_PACKET{})) {
		l.printf("err, paused packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectIsDebuggerRunning)
		return nil
	}
	var pkt hyperdbgsdk.DEBUGGEE_KD_PAUSED_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))

	// Pause the logging mechanism (mirror g_IgnoreNewLoggingMessages = TRUE).
	l.state.SetIgnoreNewLoggingMessages(true)
	// Debuggee is no longer running.
	l.state.SetDebuggeeRunning(false)
	// Save the current core + instruction.
	l.state.SetCurrentRemoteCore(pkt.CurrentCore)
	var instr [MaximumInstrSize]byte
	copy(instr[:], pkt.InstructionBytesOnRip[:])
	l.state.SetCurrentRunningInstruction(instr, pkt.IsProcessorOn32BitMode)

	ev := fromPausingReason(pkt.PausingReason)

	// Print the pre-disassembly context messages. Mirrors the first switch in
	// the C++ case arm.
	switch ev {
	case KernelEventBreakpointHit:
		if pkt.EventTag != 0 {
			l.printf("breakpoint 0x%x hit\n", pkt.EventTag)
		}
	case KernelEventTriggered:
		if pkt.EventTag != 0 {
			if pkt.EventCallingStage == hyperdbgsdk.VmmCallbackCallingStagePostEventEmulation {
				l.printf("event 0x%x triggered (post)\n", pkt.EventTag-DebuggerEventTagStartSeed)
			} else {
				l.printf("event 0x%x triggered (pre)\n", pkt.EventTag-DebuggerEventTagStartSeed)
			}
		}
	case KernelEventProcessSwitched:
		l.printf("switched to the specified process\n")
	case KernelEventThreadSwitched:
		l.printf("switched to the specified thread\n")
	case KernelEventStartingModuleLoaded:
		l.printf("the target module is loaded and a breakpoint is set to the entrypoint\n")
		l.printf("press 'g' to reach to the entrypoint of the main module...\n")
	}

	// Call the user-supplied handler (CLI prints disassembly, GUI highlights
	// the line, MCP emits a JSON event). The C++ listener calls
	// HyperDbgDisassembler64/32 here; that logic now lives in the handler so
	// the listener stays disassembler-agnostic.
	l.mu.Lock()
	handler := l.handler
	l.mu.Unlock()
	if handler != nil && !pkt.IgnoreDisassembling {
		if err := handler(ctx, l.state, ev, pkt); err != nil {
			// The handler bailed; still signal the waiter so the waiting
			// command does not deadlock.
			l.signalPauseSync(ev)
			return err
		}
	}

	// Signal the sync object that matches the pausing reason.
	l.signalPauseSync(ev)
	return nil
}

// signalPauseSync wakes up the goroutine blocked on the sync object that
// corresponds to the given pausing reason. Mirrors the second switch in the
// DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION case arm.
func (l *KernelListener) signalPauseSync(ev KernelEvent) {
	switch ev {
	case KernelEventBreakpointHit,
		KernelEventHardwareDebugRegisterHit,
		KernelEventTriggered,
		KernelEventSingleStep,
		KernelEventProcessSwitched,
		KernelEventThreadSwitched,
		KernelEventTrackingStep,
		KernelEventStartingModuleLoaded,
		KernelEventGeneralDebugBreak,
		KernelEventGeneralThreadIntercepted,
		KernelEventHardwareBasedGeneralBreak:
		l.state.ReceivedKernelResponse(SyncObjectIsDebuggerRunning)
	case KernelEventCoreSwitched:
		l.state.ReceivedKernelResponse(SyncObjectCoreSwitchingResult)
	case KernelEventCommandFinished:
		l.printf("\n")
		l.state.ReceivedKernelResponse(SyncObjectDebuggeeFinishedCommandExecution)
	case KernelEventRequestFromDebugger:
		l.state.ReceivedKernelResponse(SyncObjectPausedDebuggeeDetails)
	case KernelEventPaused:
		// Nothing — the pause was user-triggered and no command is waiting.
		// The interactive loop polls IsDebuggeeRunning() instead.
	default:
		l.printf("err, unknown pausing reason is received\n")
	}
}

// handleChangeCoreResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE.
func (l *KernelListener) handleChangeCoreResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_CHANGE_CORE_PACKET{})) {
		l.printf("err, change-core packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectCoreSwitchingResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_CHANGE_CORE_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result == DebuggerOperationWasSuccessful {
		l.printf("current operating core changed to 0x%x\n", pkt.NewCore)
	} else {
		l.printf("err, change-core failed with error 0x%x\n", pkt.Result)
		l.state.ReceivedKernelResponse(SyncObjectCoreSwitchingResult)
	}
}

// handleChangeProcessResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS.
func (l *KernelListener) handleChangeProcessResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET{})) {
		l.printf("err, change-process packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectProcessSwitchingResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result == DebuggerOperationWasSuccessful {
		switch pkt.ActionType {
		case hyperdbgsdk.DebuggeeDetailsAndSwitchProcessGetProcessDetails:
			l.printf("process id: %x\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
				pkt.ProcessId,
				separateTo64BitValue(pkt.Process),
				uint8String(pkt.ProcessName[:]))
		case hyperdbgsdk.DebuggeeDetailsAndSwitchProcessPerformSwitch:
			l.printf("press 'g' to continue the debuggee, if the pid or the process object address is valid then the debuggee will be automatically paused when it attached to the target process\n")
		}
	} else {
		l.printf("err, change-process failed with error 0x%x\n", pkt.Result)
	}
	l.state.ReceivedKernelResponse(SyncObjectProcessSwitchingResult)
}

// handleChangeThreadResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD.
func (l *KernelListener) handleChangeThreadResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET{})) {
		l.printf("err, change-thread packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectThreadSwitchingResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result == DebuggerOperationWasSuccessful {
		switch pkt.ActionType {
		case hyperdbgsdk.DebuggeeDetailsAndSwitchThreadGetThreadDetails:
			l.printf("thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
				pkt.ThreadId, pkt.ProcessId,
				separateTo64BitValue(pkt.Thread),
				separateTo64BitValue(pkt.Process),
				uint8String(pkt.ProcessName[:]))
		case hyperdbgsdk.DebuggeeDetailsAndSwitchThreadPerformSwitch:
			l.printf("press 'g' to continue the debuggee, if the tid or the thread object address is valid then the debuggee will be automatically paused when it attached to the target thread\n")
		}
	} else {
		l.printf("err, change-thread failed with error 0x%x\n", pkt.Result)
	}
	l.state.ReceivedKernelResponse(SyncObjectThreadSwitchingResult)
}

// handleSearchQueryResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY.
func (l *KernelListener) handleSearchQueryResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET{})) {
		l.printf("err, search-query packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectSearchQueryResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result == DebuggerOperationWasSuccessful {
		if pkt.CountOfResults == 0 {
			l.printf("not found\n")
		}
	} else {
		l.printf("err, search query failed with error 0x%x\n", pkt.Result)
	}
	l.state.ReceivedKernelResponse(SyncObjectSearchQueryResult)
}

// handleFlushResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH.
func (l *KernelListener) handleFlushResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_FLUSH_LOGGING_BUFFERS{})) {
		l.printf("err, flush packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectFlushResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_FLUSH_LOGGING_BUFFERS
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus == DebuggerOperationWasSuccessful {
		l.printf("flushing buffers was successful, total %d messages were cleared.\n",
			pkt.CountOfMessagesThatSetAsReadFromVmxNonRoot+pkt.CountOfMessagesThatSetAsReadFromVmxRoot)
	} else {
		l.printf("err, flush failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectFlushResult)
}

// handleCpuidResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_USER_CPUID.
// The Go types package does not yet define DEBUGGER_CPUID_REQUEST_RESPONSE, so
// we just signal the sync object. The caller can re-parse the payload from
// KdState if needed once the struct is added.
func (l *KernelListener) handleCpuidResult(p KernelPausedPacket) {
	// Best-effort: extract KernelStatus from the start of the payload (the
	// first uint32 of every result packet is KernelStatus in the C++ structs).
	if len(p.Payload) >= 4 {
		ks := binary.LittleEndian.Uint32(p.Payload[:4])
		if ks != DebuggerOperationWasSuccessful {
			l.printf("err, cpuid failed with error 0x%x\n", ks)
		}
	}
	l.state.ReceivedKernelResponse(SyncObjectUserCpuidResult)
}

// handleCallstackResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK.
// The detailed callstack rendering lives in the callstack package; the listener
// only validates the result and signals the sync object.
func (l *KernelListener) handleCallstackResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_CALLSTACK_REQUEST{})) {
		l.printf("err, callstack packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectCallstackResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_CALLSTACK_REQUEST
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != DebuggerOperationWasSuccessful {
		l.printf("err, callstack failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectCallstackResult)
}

// handleTestQueryResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY.
func (l *KernelListener) handleTestQueryResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER{})) {
		l.printf("err, test-query packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectTestQuery)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus == DebuggerOperationWasSuccessful {
		switch pkt.RequestType {
		case hyperdbgsdk.TestBreakpointTurnOffBps:
			l.printf("breakpoint interception (#BP) is deactivated\nfrom now, the breakpoints will be re-injected into the guest debuggee\n")
		case hyperdbgsdk.TestBreakpointTurnOnBps:
			l.printf("breakpoint interception (#BP) is activated\n")
		case hyperdbgsdk.TestBreakpointTurnOffDbs:
			l.printf("debug break interception (#DB) is deactivated\nfrom now, the debug breaks will be re-injected into the guest debuggee\n")
		case hyperdbgsdk.TestBreakpointTurnOnDbs:
			l.printf("debug break interception (#DB) is activated\n")
		}
	} else {
		l.printf("err, test-query failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectTestQuery)
}

// handleScriptResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT.
func (l *KernelListener) handleScriptResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_SCRIPT_PACKET{})) {
		l.printf("err, script-result packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectScriptRunningResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_SCRIPT_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result != DebuggerOperationWasSuccessful {
		l.printf("err, script failed with error 0x%x\n", pkt.Result)
	}
	if pkt.IsFormat {
		l.state.ReceivedKernelResponse(SyncObjectScriptFormatsResult)
	}
	l.state.ReceivedKernelResponse(SyncObjectScriptRunningResult)
}

// handleFormatsResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS:
// stash the evaluated value + error state in KdState for the waiting caller.
func (l *KernelListener) handleFormatsResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_FORMATS_PACKET{})) {
		l.printf("err, formats packet too small\n")
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_FORMATS_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	l.state.SetFormatsResult(pkt.Value, pkt.Result)
}

// handleRegisterEventResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT.
func (l *KernelListener) handleRegisterEventResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})) {
		l.printf("err, register-event packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectRegisterEvent)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	l.state.SetResultOfRegisteringEvent(pkt)
	l.state.ReceivedKernelResponse(SyncObjectRegisterEvent)
}

// handleAddActionResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT.
func (l *KernelListener) handleAddActionResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})) {
		l.printf("err, add-action packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectAddActionToEvent)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	l.state.SetResultOfAddingActions(pkt)
	l.state.ReceivedKernelResponse(SyncObjectAddActionToEvent)
}

// handleQueryAndModifyEventResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT.
func (l *KernelListener) handleQueryAndModifyEventResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_MODIFY_EVENTS{})) {
		l.printf("err, query-modify-event packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectModifyAndQueryEvent)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_MODIFY_EVENTS
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != uint64(DebuggerOperationWasSuccessful) {
		l.printf("err, query-modify-event failed with error 0x%x\n", uint32(pkt.KernelStatus))
	} else if pkt.TypeOfAction == hyperdbgsdk.DebuggerModifyEventsQueryState {
		l.state.SetSharedEventStatus(pkt.IsEnabled)
	}
	l.state.ReceivedKernelResponse(SyncObjectModifyAndQueryEvent)
}

// handleSymbolReloadFinished mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED.
// The actual symbol reload (SymbolInitialReload) is performed by the
// symbolparser package; the listener only reports the result here.
func (l *KernelListener) handleSymbolReloadFinished(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_SYMBOL_UPDATE_RESULT{})) {
		l.printf("err, symbol-reload packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectSymbolReload)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_SYMBOL_UPDATE_RESULT
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != uint64(DebuggerOperationWasSuccessful) {
		l.printf("err, symbol reload failed with error 0x%x\n", uint32(pkt.KernelStatus))
	}
	l.state.ReceivedKernelResponse(SyncObjectSymbolReload)
}

// handleReadRegistersResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS.
// The caller-supplied buffer (set via SetKernelRequestData) is filled with the
// payload so the waiting IOCTL caller sees the response.
func (l *KernelListener) handleReadRegistersResult(p KernelPausedPacket) {
	l.copyToCallerBuffer(p, SyncObjectReadRegisters)
}

// handleWriteRegisterResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_WRITE_REGISTER.
func (l *KernelListener) handleWriteRegisterResult(p KernelPausedPacket) {
	l.copyToCallerBuffer(p, SyncObjectWriteRegister)
}

// handleCallerBufferResult is the generic form of the read/write register,
// read/edit memory, APIC, IDT, SMI, HyperTrace LBR/PT handlers: copy the
// payload into the caller-supplied buffer (if any) and signal the sync object.
// Mirrors the DbgWaitGetKernelRequestData + memcpy + DbgReceivedKernelResponse
// pattern repeated across the C++ switch.
func (l *KernelListener) handleCallerBufferResult(p KernelPausedPacket, obj KernelSyncObject) {
	l.copyToCallerBuffer(p, obj)
}

// copyToCallerBuffer copies the payload into the buffer previously registered
// via SetKernelRequestData (capped at the buffer size) and signals the sync
// object. Mirrors the DbgWaitGetKernelRequestData + memcpy pattern.
func (l *KernelListener) copyToCallerBuffer(p KernelPausedPacket, obj KernelSyncObject) {
	callerBuf, ok := l.state.GetKernelRequestData(obj)
	if ok {
		n := len(p.Payload)
		if n > len(callerBuf) {
			n = len(callerBuf)
		}
		copy(callerBuf, p.Payload[:n])
	}
	l.state.ReceivedKernelResponse(obj)
}

// handleBpResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP.
func (l *KernelListener) handleBpResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_BP_PACKET{})) {
		l.printf("err, bp-result packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectBp)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_BP_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result != DebuggerOperationWasSuccessful {
		l.printf("err, bp failed with error 0x%x\n", pkt.Result)
	}
	l.state.ReceivedKernelResponse(SyncObjectBp)
}

// handleShortCircuitingResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE.
func (l *KernelListener) handleShortCircuitingResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_SHORT_CIRCUITING_EVENT{})) {
		l.printf("err, short-circuiting packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectShortCircuitingEventState)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_SHORT_CIRCUITING_EVENT
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus == uint64(DebuggerOperationWasSuccessful) {
		state := "'off'"
		if pkt.IsShortCircuiting {
			state = "'on'"
		}
		l.printf("the event's short-circuiting state changed to %s\n", state)
	} else {
		l.printf("err, short-circuiting failed with error 0x%x\n", uint32(pkt.KernelStatus))
	}
	l.state.ReceivedKernelResponse(SyncObjectShortCircuitingEventState)
}

// handlePteResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE.
// The detailed PTE rendering (CommandPteShowResults) lives in the commands
// package; the listener only validates the result and signals the sync object.
func (l *KernelListener) handlePteResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{})) {
		l.printf("err, pte packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectPteResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != DebuggerOperationWasSuccessful {
		l.printf("err, pte query failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectPteResult)
}

// handlePageInResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BRINGING_PAGES_IN.
func (l *KernelListener) handlePageInResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_PAGE_IN_REQUEST{})) {
		l.printf("err, page-in packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectPageInState)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_PAGE_IN_REQUEST
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus == DebuggerOperationWasSuccessful {
		l.printf("the page-fault is delivered to the target thread\npress 'g' to continue debuggee (the current thread will execute ONLY one instruction and will be halted again)...\n")
	} else {
		l.printf("err, page-in failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectPageInState)
}

// handleVa2paPa2vaResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA.
func (l *KernelListener) handleVa2paPa2vaResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{})) {
		l.printf("err, va2pa/pa2va packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectVa2paAndPa2vaResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus == DebuggerOperationWasSuccessful {
		if pkt.IsVirtual2Physical {
			l.printf("%llx\n", pkt.PhysicalAddress)
		} else {
			l.printf("%llx\n", pkt.VirtualAddress)
		}
	} else {
		l.printf("err, va2pa/pa2va failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectVa2paAndPa2vaResult)
}

// handleListOrModifyBpResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS.
func (l *KernelListener) handleListOrModifyBpResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{})) {
		l.printf("err, list-modify-bp packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectListOrModifyBreakpoints)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.Result != DebuggerOperationWasSuccessful {
		l.printf("err, list-modify-bp failed with error 0x%x\n", pkt.Result)
	}
	l.state.ReceivedKernelResponse(SyncObjectListOrModifyBreakpoints)
}

// handlePcitreeResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCITREE.
// The detailed pcitree rendering (vendor/device name lookup via the pci-id
// database) lives in the commands/extension package; the listener only
// validates the result and signals the sync object.
func (l *KernelListener) handlePcitreeResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET{})) {
		l.printf("err, pcitree packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectPcitreeResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != DebuggerOperationWasSuccessful {
		l.printf("err, pcitree failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectPcitreeResult)
}

// handlePcidevinfoResult mirrors DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCIDEVINFO.
// The detailed pcidevinfo rendering lives in the commands/extension package;
// the listener only validates the result and signals the sync object.
func (l *KernelListener) handlePcidevinfoResult(p KernelPausedPacket) {
	if uint32(len(p.Payload)) < uint32(unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET{})) {
		l.printf("err, pcidevinfo packet too small\n")
		l.state.ReceivedKernelResponse(SyncObjectPcidevinfoResult)
		return
	}
	var pkt hyperdbgsdk.DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET
	bytesIntoStruct(unsafe.Pointer(&pkt), p.Payload, unsafe.Sizeof(pkt))
	if pkt.KernelStatus != DebuggerOperationWasSuccessful {
		l.printf("err, pcidevinfo failed with error 0x%x\n", pkt.KernelStatus)
	}
	l.state.ReceivedKernelResponse(SyncObjectPcidevinfoResult)
}

// ----------------------------------------------------------------------------
// Default handler.
// ----------------------------------------------------------------------------

// DefaultKernelHandler returns a Handler that prints a CLI-style summary of
// the pause, mirroring the ShowMessages calls in the
// DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION case arm of
// ListeningSerialPortInDebugger (excluding the disassembly itself, which is
// delegated to the disassembler package and would be wired in by the CLI).
//
// The handler is intentionally lightweight: it formats the RIP, the current
// core, and the instruction bytes, so the listener's output stays useful even
// before a disassembler is plugged in. Pass nil to silence all output.
func DefaultKernelHandler(out Output) KernelHandler {
	return func(ctx context.Context, state *KdState, ev KernelEvent, pkt hyperdbgsdk.DEBUGGEE_KD_PAUSED_PACKET) error {
		_ = ctx
		if out == nil {
			return nil
		}
		mode := "64-bit"
		if pkt.IsProcessorOn32BitMode {
			mode = "32-bit"
		}
		out.Printf("rip@%016llx [%s, core=%d] :", pkt.Rip, mode, pkt.CurrentCore)
		for _, b := range pkt.InstructionBytesOnRip {
			out.Printf(" %02x", b)
		}
		out.Printf("\n")
		return nil
	}
}

// ----------------------------------------------------------------------------
// Internal helpers.
// ----------------------------------------------------------------------------

// printf writes to the state's Output sink. Best-effort: errors are silently
// dropped (the C++ ShowMessages has no error path either).
func (l *KernelListener) printf(format string, args ...any) {
	l.state.printf(format, args...)
}

// int8String converts a NUL-terminated []int8 (the C `char[]` shape used by
// the SDK structs) to a Go string. Mirrors the implicit char*→string
// conversion in the C++ ShowMessages("%s", ...) calls.
func int8String(b []int8) string {
	n := 0
	for n < len(b) && b[n] != 0 {
		n++
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = byte(b[i])
	}
	return string(out)
}

// uint8String converts a NUL-terminated []uint8 to a Go string. Used for the
// [16]uint8 ProcessName field on DEBUGGEE_DETAILS_AND_SWITCH_*_PACKET.
func uint8String(b []uint8) string {
	n := 0
	for n < len(b) && b[n] != 0 {
		n++
	}
	return string(b[:n])
}

// separateTo64BitValue mirrors the C++ SeparateTo64BitValue helper: it formats
// a uint64 as "high:low" so 32-bit shells can read it. The C++ version
// returns a std::string; the Go version returns the same shape.
func separateTo64BitValue(v uint64) string {
	hi := uint32(v >> 32)
	lo := uint32(v)
	return fmt.Sprintf("%x:%x", hi, lo)
}
