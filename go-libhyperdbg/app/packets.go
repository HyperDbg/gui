// Package app — packets.go
//
// Implements the IRP-based kernel→user buffer pump. The C++ counterpart is
// libhyperdbg/code/app/packets.cpp; it owns:
//   - ReadIrpBasedBuffer()         — the blocking loop that opens a dedicated
//     device handle, sends IOCTL_REGISTER_EVENT
//     with Type=IRP_BASED, and dispatches the
//     returned buffer based on OperationCode
//   - IrpBasedBufferThread(Data)   — the Win32 thread wrapper around the loop
//
// In the Go rewrite the loop runs in a goroutine owned by PacketReader; the
// state machine mirrors the C++ switch (OperationCode) block. The reader is
// stoppable via Stop() (which flips g_IsMessageLoggingWindowClosed) and via
// context cancellation.
//
// Lifecycle:
//
//	pr := NewPacketReader(core, messaging, out)
//	pr.Start(ctx)   // spawns the goroutine
//	...
//	pr.Stop()       // signals the loop to exit at the next iteration
//
// PacketReader is safe for concurrent use; the loop goroutine is the only
// writer of the internal state, the public methods only flip the stop flag.
package app

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// Operation codes mirroring the C++ #defines in
// hyperdbg/include/SDK/headers/Constants.h. They are declared here (rather
// than in types/sdk.go) because they are plain numeric macros, not enum
// types — matching the existing convention used by the kernellvl package.
const (
	operationMandatoryDebuggeeBit                        uint32 = 1 << 31
	operationLogInfoMessage                              uint32 = 1
	operationLogWarningMessage                           uint32 = 2
	operationLogErrorMessage                             uint32 = 3
	operationLogNonImmediateMessage                      uint32 = 4
	operationLogWithTag                                  uint32 = 5
	operationLogMessageMandatory                         uint32 = 6 | operationMandatoryDebuggeeBit
	operationCommandFromDebuggerCloseAndUnloadVmm        uint32 = 7 | operationMandatoryDebuggeeBit
	operationDebuggeeUserInput                           uint32 = 8 | operationMandatoryDebuggeeBit
	operationDebuggeeRegisterEvent                       uint32 = 9 | operationMandatoryDebuggeeBit
	operationDebuggeeAddActionToEvent                    uint32 = 10 | operationMandatoryDebuggeeBit
	operationDebuggeeClearEvents                         uint32 = 11 | operationMandatoryDebuggeeBit
	operationDebuggeeClearEventsWithoutNotifyingDebugger uint32 = 12 | operationMandatoryDebuggeeBit
	operationHypervisorDriverIsSuccessfullyLoaded        uint32 = 13 | operationMandatoryDebuggeeBit
	operationHypervisorDriverEndOfIrps                   uint32 = 14 | operationMandatoryDebuggeeBit
	operationCommandFromDebuggerReloadSymbol             uint32 = 15 | operationMandatoryDebuggeeBit
	operationNotificationFromUserDebuggerPause           uint32 = 16 | operationMandatoryDebuggeeBit
)

// UsermodeBufferSize mirrors UsermodeBufferSize (sizeof(UINT32) + PacketChunkSize + 1).
// It is the size of the buffer the reader passes to DeviceIoControl.
const UsermodeBufferSize = 4 + 4096 + 1

// DefaultSpeedOfReadingKernelMessages mirrors
// DefaultSpeedOfReadingKernelMessages (30 ms). It is the sleep between
// non-mandatory packets to avoid eating all CPU.
const DefaultSpeedOfReadingKernelMessages = 30 * time.Millisecond

// SizeofRegisterEvent mirrors SIZEOF_REGISTER_EVENT * 2 (the driver is x64
// and the C++ code passes 2× the struct size to account for 64-bit values).
const SizeofRegisterEvent = 16 // sizeof(REGISTER_NOTIFY_BUFFER) == 16 on x64

// PacketReader owns the kernel→user buffer pump. It is constructed by App
// once the device is open and runs a single goroutine that blocks on
// IOCTL_REGISTER_EVENT. The reader is bound to one *core.Debugger (for the
// device handle), one *Messaging (for ShowMessages dispatch) and one Output
// (for diagnostic messages from the reader itself).
type PacketReader struct {
	mu sync.Mutex

	core      *core.Debugger
	messaging *Messaging
	out       Output

	// stop is closed by Stop() to signal the loop goroutine to exit at the
	// next iteration. A nil stop means Start has not been called.
	stop chan struct{}
	// done is closed by the loop goroutine when it has fully exited, so
	// Stop can wait for clean shutdown.
	done chan struct{}

	// breakPrintingOutput mirrors g_BreakPrintingOutput. When true, log/info/
	// warning/error packets are silently dropped (CTRL+C path).
	breakPrintingOutput bool
}

// NewPacketReader constructs a reader bound to the given core/messaging/out
// triple. The returned reader is not running; call Start to spawn the loop.
func NewPacketReader(c *core.Debugger, m *Messaging, out Output) *PacketReader {
	if out == nil {
		out = discardOutput{}
	}
	return &PacketReader{core: c, messaging: m, out: out}
}

// Start spawns the IRP-reading goroutine. It is idempotent: calling it twice
// without an intervening Stop is a no-op. The ctx is used for the initial
// device-open check only; the loop itself uses the reader's stop channel for
// cancellation because DeviceIoControl is synchronous.
func (p *PacketReader) Start(ctx context.Context) {
	p.mu.Lock()
	if p.stop != nil {
		p.mu.Unlock()
		return
	}
	p.stop = make(chan struct{})
	p.done = make(chan struct{})
	stop := p.stop
	done := p.done
	p.mu.Unlock()

	go p.readIrpBasedBuffer(ctx, stop, done)
}

// Stop signals the loop goroutine to exit at the next iteration and waits
// for it to actually return. Safe to call multiple times; subsequent calls
// are no-ops.
func (p *PacketReader) Stop() {
	p.mu.Lock()
	if p.stop == nil {
		p.mu.Unlock()
		return
	}
	stop := p.stop
	done := p.done
	p.stop = nil
	p.mu.Unlock()
	close(stop)
	<-done
}

// SetBreakPrintingOutput mirrors toggling g_BreakPrintingOutput. When true,
// log/info/warning/error packets are silently dropped (used by the CTRL+C
// signal handler to suppress noise while the user is typing).
func (p *PacketReader) SetBreakPrintingOutput(v bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.breakPrintingOutput = v
}

// readIrpBasedBuffer is the main loop. It mirrors the C++ ReadIrpBasedBuffer:
// open a dedicated device handle (separate from the main IOCTL handle so the
// pending IRP doesn't block other commands), then loop sending
// IOCTL_REGISTER_EVENT with Type=IRP_BASED and dispatch the returned buffer.
func (p *PacketReader) readIrpBasedBuffer(ctx context.Context, stop chan struct{}, done chan struct{}) {
	defer close(done)

	// Open a dedicated handle for the packet reader. The C++ code opens
	// \\.\HyperDbgDebuggerDevice again here so the pending IOCTL doesn't
	// block the main handle. We do the same via comm.Open.
	dev, err := comm.Open(ctx, comm.DeviceName)
	if err != nil {
		p.printf("err, packet reader failed to open device: %v\n", err)
		return
	}
	defer dev.Close()

	// REGISTER_NOTIFY_BUFFER: Type=IRP_BASED, HEvent=NULL.
	var reg hyperdbgsdk.REGISTER_NOTIFY_BUFFER
	reg.Type = hyperdbgsdk.IrpBased
	regBuf := structAsBytes(unsafe.Pointer(&reg), unsafe.Sizeof(reg))

	out := make([]byte, UsermodeBufferSize)

	for {
		select {
		case <-stop:
			return
		default:
		}

		// Clear the buffer (mirrors ZeroMemory in the C++ loop).
		for i := range out {
			out[i] = 0
		}

		// Synchronous IOCTL — blocks until the driver has a packet to
		// deliver. ctx cancellation cannot interrupt DeviceIoControl; the
		// caller uses Stop() to flip g_IsMessageLoggingWindowClosed and
		// the driver then completes the IRP with
		// OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS.
		n, err := dev.Ioctl(ctx, comm.IOCTL_CODE_REGISTER_EVENT, regBuf, out)
		if err != nil {
			// Mirror the C++ behaviour: log nothing and continue. The
			// packet probably failed because of a flush command.
			continue
		}

		// Compute the operation code from the first 4 bytes of the buffer.
		var opCode uint32
		if n >= 4 {
			opCode = binary.LittleEndian.Uint32(out[:4])
		}

		// Sleep between non-mandatory packets to avoid eating all CPU
		// (mirrors PlatformSleep(DefaultSpeedOfReadingKernelMessages)).
		if (opCode & operationMandatoryDebuggeeBit) == 0 {
			select {
			case <-stop:
				return
			case <-time.After(DefaultSpeedOfReadingKernelMessages):
			}
		}

		if !p.dispatchPacket(opCode, out[:n], n) {
			return
		}
	}
}

// dispatchPacket handles a single decoded kernel packet. Returns false to
// signal the loop should exit (mirrors the C++ loop's
// g_IsMessageLoggingWindowClosed check).
func (p *PacketReader) dispatchPacket(opCode uint32, buf []byte, n uint32) bool {
	// payload is the message body that follows the 4-byte OperationCode
	// prefix (mirrors OutputBuffer + sizeof(UINT32) in the C++ code).
	var payload []byte
	if uint32(len(buf)) > 4 {
		payload = buf[4:]
	}

	switch opCode {
	case operationLogNonImmediateMessage,
		operationLogInfoMessage,
		operationLogErrorMessage,
		operationLogWarningMessage:
		// Drop these when the user is mid-CTRL+C.
		p.mu.Lock()
		breakPrinting := p.breakPrintingOutput
		p.mu.Unlock()
		if breakPrinting {
			return true
		}
		p.showMessages("%s", trimNul(payload))

	case operationLogMessageMandatory:
		// Mandatory messages bypass the break-printing flag.
		p.showMessages("%s", trimNul(payload))

	case operationCommandFromDebuggerCloseAndUnloadVmm:
		// KdCloseConnection — handled by the higher-level api layer in the
		// full implementation. For the framework stub we just log.
		p.printf("packet: COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM received\n")

	case operationDebuggeeUserInput:
		// KdHandleUserInputInDebuggee — Phase C.3.
		p.printf("packet: DEBUGGEE_USER_INPUT received (len=%d)\n", n)

	case operationDebuggeeRegisterEvent:
		// KdRegisterEventInDebuggee — Phase C.3.
		p.printf("packet: DEBUGGEE_REGISTER_EVENT received (len=%d)\n", n)

	case operationDebuggeeAddActionToEvent:
		// KdAddActionToEventInDebuggee — Phase C.3.
		p.printf("packet: DEBUGGEE_ADD_ACTION_TO_EVENT received (len=%d)\n", n)

	case operationDebuggeeClearEvents,
		operationDebuggeeClearEventsWithoutNotifyingDebugger:
		// KdSendModifyEventInDebuggee — Phase C.3.
		p.printf("packet: DEBUGGEE_CLEAR_EVENTS received (len=%d)\n", n)

	case operationHypervisorDriverIsSuccessfullyLoaded:
		// The C++ code signals g_IsDriverLoadedSuccessfully here. The Go
		// api layer wires this to a sync.Cond / channel in the future;
		// for the framework stub we just log.
		p.printf("packet: HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED\n")

	case operationHypervisorDriverEndOfIrps:
		// End of receiving messages — the driver completed the pending IRP
		// because g_IsMessageLoggingWindowClosed was flipped. Exit the loop.
		return false

	case operationCommandFromDebuggerReloadSymbol:
		// KdReloadSymbolsInDebuggee(TRUE, ProcessId) — Phase C.3.
		p.printf("packet: COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL received\n")

	case operationNotificationFromUserDebuggerPause:
		// UdHandleUserDebuggerPausing — Phase C.3. The user debugger paused
		// the debuggee; the higher-level layer reads the
		// DEBUGGEE_UD_PAUSED_PACKET from the payload.
		p.printf("packet: NOTIFICATION_FROM_USER_DEBUGGER_PAUSE received (len=%d)\n", n)

	default:
		// ForwardingCheckAndPerformEventForwarding path — for the framework
		// stub we just print the payload (the C++ default branch falls back
		// to ShowMessages if no output source consumed the packet).
		breakPrinting := p.breakPrintingOutput
		if breakPrinting {
			return true
		}
		p.showMessages("%s", trimNul(payload))
	}
	return true
}

// printf writes a diagnostic line from the reader itself (not a kernel
// message) to the local Output. Best-effort: errors are dropped.
func (p *PacketReader) printf(format string, args ...any) {
	_, _ = p.out.Write([]byte(fmt.Sprintf(format, args...)))
}

// showMessages routes a kernel message through the Messaging dispatcher.
func (p *PacketReader) showMessages(format string, args ...any) {
	if p.messaging == nil {
		_, _ = p.out.Write([]byte(fmt.Sprintf(format, args...)))
		return
	}
	p.messaging.ShowMessages(format, args...)
}

// trimNul returns buf truncated at the first NUL byte and with any trailing
// NULs stripped. Kernel messages are NUL-terminated C strings.
func trimNul(buf []byte) []byte {
	for i, b := range buf {
		if b == 0 {
			return buf[:i]
		}
	}
	return buf
}

// structAsBytes returns a byte slice aliasing the memory at ptr for size
// bytes. Mirrors the helper in debugger/kernellvl/kd.go.
func structAsBytes(ptr unsafe.Pointer, size uintptr) []byte {
	if ptr == nil || size == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(ptr), size)
}
