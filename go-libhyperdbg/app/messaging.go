// Package app — messaging.go
//
// Implements the kernel→user message-channel dispatcher. The C++ counterpart
// is libhyperdbg/code/app/messaging.cpp; it owns:
//   - g_MessageHandler             — function pointer to the user-supplied
//     callback that receives every ShowMessages
//     output
//   - g_MessageHandlerSharedBuffer — optional shared buffer the callback reads
//     instead of receiving a pointer argument
//     (avoids an extra copy on the hot path)
//   - g_LogOpened                  — whether .logopen has opened a log file
//   - g_IsConnectedToRemoteDebugger / g_IsSerialConnectedToRemoteDebugger
//     — whether messages must also be forwarded
//     to a remote debugger over TCP or serial
//
// In the Go rewrite the global state from the C side is owned by the
// Messaging struct so that multiple debugger instances can coexist (GUI/MCP
// requirement, see API design spec).
//
// ShowMessages is the single entry point used by every command handler: it
// formats the message, writes it to the local Output sink, optionally appends
// it to the log file, optionally forwards it to a remote debugger, and finally
// invokes the user-supplied callback (if any).
package app

import (
	"fmt"
	"os"
	"sync"
)

// CommunicationBufferSize mirrors COMMUNICATION_BUFFER_SIZE (PacketChunkSize + 0x100).
// It is the size of the shared buffer used by SetCallbackUsingSharedBuffer.
const CommunicationBufferSize = 4096 + 0x100

// TcpEndOfBufferCharsCount mirrors TCP_END_OF_BUFFER_CHARS_COUNT (4 trailing
// NUL bytes used to frame the end of a TCP-forwarded message).
const TcpEndOfBufferCharsCount = 4

// MessageCallback is the user-supplied handler invoked by ShowMessages when a
// message is produced. It mirrors SendMessageWithParamCallback from the C++
// code: it receives the formatted message string.
type MessageCallback func(msg string)

// MessageCallbackSharedBuffer is the variant that receives no argument; the
// callback reads the message from a pre-allocated shared buffer that the
// dispatcher fills before invoking the callback. Mirrors
// SendMessageWithSharedBufferCallback.
type MessageCallbackSharedBuffer func()

// Messaging owns the message dispatch state. All fields are guarded by mu.
// The zero value is not usable; use NewMessaging.
type Messaging struct {
	mu sync.Mutex

	out        Output
	callback   MessageCallback
	sharedMode bool
	sharedBuf  []byte

	logFile *os.File
	logPath string

	// Forwarding flags mirroring g_IsConnectedToRemoteDebugger and
	// g_IsSerialConnectedToRemoteDebugger. When true, ShowMessages also ships
	// the formatted message to the remote debugger over the configured
	// transport (TCP or serial). The transports themselves live in the
	// debugger/comm sub-packages; Messaging only knows whether forwarding is
	// active.
	connectedToRemoteDebugger       bool
	serialConnectedToRemoteDebugger bool

	// forwardTcp / forwardSerial are optional sinks used when the
	// corresponding *ConnectedToRemoteDebugger flag is true. They are set
	// by the higher-level api layer when a remote session is established.
	forwardTcp    func(msg []byte) error
	forwardSerial func(msg []byte) error
}

// NewMessaging constructs a Messaging writing to out. The returned value is
// ready for ShowMessages; callbacks and log file are configured separately.
func NewMessaging(out Output) *Messaging {
	if out == nil {
		out = discardOutput{}
	}
	return &Messaging{out: out}
}

// SetCallback installs a SendMessageWithParamCallback-style handler. The
// handler receives the formatted message string on every ShowMessages call.
// Pass nil to clear a previously installed handler. Mirrors
// SetTextMessageCallback.
func (m *Messaging) SetCallback(cb MessageCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callback = cb
	m.sharedMode = false
	m.sharedBuf = nil
}

// SetCallbackUsingSharedBuffer installs a SendMessageWithSharedBufferCallback-
// style handler and returns the shared buffer the dispatcher fills before
// each invocation. Returns nil if allocation fails. Mirrors
// SetTextMessageCallbackUsingSharedBuffer.
func (m *Messaging) SetCallbackUsingSharedBuffer(cb MessageCallbackSharedBuffer) []byte {
	m.mu.Lock()
	defer m.mu.Unlock()
	buf := make([]byte, CommunicationBufferSize+TcpEndOfBufferCharsCount)
	m.callback = func(_ string) { cb() }
	m.sharedMode = true
	m.sharedBuf = buf
	return buf
}

// UnsetCallback clears any previously installed handler and frees the shared
// buffer (if any). Mirrors UnsetTextMessageCallback.
func (m *Messaging) UnsetCallback() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callback = nil
	m.sharedMode = false
	m.sharedBuf = nil
}

// SetLogOpen opens path for appending and routes every ShowMessages call to
// it in addition to the local Output. Mirrors the .logopen command. Returns
// an error if the file cannot be opened.
func (m *Messaging) SetLogOpen(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("messaging: open log %q: %w", path, err)
	}
	m.mu.Lock()
	// Close any previously open log file.
	if m.logFile != nil {
		_ = m.logFile.Close()
	}
	m.logFile = f
	m.logPath = path
	m.mu.Unlock()
	return nil
}

// SetLogClose closes the open log file (if any). Mirrors the .logclose
// command. Safe to call when no log is open.
func (m *Messaging) SetLogClose() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.logFile == nil {
		return nil
	}
	err := m.logFile.Close()
	m.logFile = nil
	m.logPath = ""
	return err
}

// IsLogOpen reports whether a log file is currently open.
func (m *Messaging) IsLogOpen() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.logFile != nil
}

// SetConnectedToRemoteDebugger toggles the g_IsConnectedToRemoteDebugger
// flag. When true, ShowMessages also invokes forwardTcp (if set) so the
// remote debugger sees the message.
func (m *Messaging) SetConnectedToRemoteDebugger(v bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.connectedToRemoteDebugger = v
}

// SetSerialConnectedToRemoteDebugger toggles the
// g_IsSerialConnectedToRemoteDebugger flag. When true, ShowMessages also
// invokes forwardSerial (if set) so the remote debugger sees the message
// over the serial link.
func (m *Messaging) SetSerialConnectedToRemoteDebugger(v bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.serialConnectedToRemoteDebugger = v
}

// SetForwarders installs the TCP and serial forwarding sinks. Pass nil to
// disable a transport. The higher-level api layer wires these to the
// debugger/comm named-pipe / TCP transports when a remote session is
// established.
func (m *Messaging) SetForwarders(tcp, serial func(msg []byte) error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.forwardTcp = tcp
	m.forwardSerial = serial
}

// ShowMessages formats and dispatches a message. It mirrors the C++
// ShowMessages entry point: when no callback and no remote connection is
// active, the message is printed to the local Output (vprintf path); when a
// remote connection is active, the message is forwarded; when a log file is
// open, the message is appended; when a callback is installed, the callback
// is invoked with the formatted message (or with the shared buffer filled).
//
// Errors from the optional forwarders / log file are silently dropped to
// match the C++ behaviour (ShowMessages has no error path).
func (m *Messaging) ShowMessages(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)

	m.mu.Lock()
	cb := m.callback
	sharedMode := m.sharedMode
	sharedBuf := m.sharedBuf
	logFile := m.logFile
	tcpActive := m.connectedToRemoteDebugger
	serialActive := m.serialConnectedToRemoteDebugger
	tcpFn := m.forwardTcp
	serialFn := m.forwardSerial
	out := m.out
	m.mu.Unlock()

	// Mirror the C++ "no handler and no remote connection → vprintf" path:
	// when neither a callback nor a remote connection is active, route the
	// message to the local Output sink.
	if cb == nil && !tcpActive && !serialActive {
		_, _ = out.Write([]byte(msg))
		if logFile == nil {
			return
		}
	}

	// Forward to remote debugger over TCP / serial.
	if tcpActive && tcpFn != nil {
		_ = tcpFn([]byte(msg))
	}
	if serialActive && serialFn != nil {
		_ = serialFn([]byte(msg))
	}

	// Append to log file.
	if logFile != nil {
		_, _ = logFile.WriteString(msg)
	}

	// Invoke user callback.
	if cb != nil {
		if sharedMode && sharedBuf != nil {
			// Copy the message into the shared buffer (NUL-terminated) and
			// invoke the no-arg callback.
			n := len(msg)
			if n >= len(sharedBuf) {
				n = len(sharedBuf) - 1
			}
			copy(sharedBuf, msg[:n])
			sharedBuf[n] = 0
			cb("")
		} else {
			cb(msg)
		}
	}
}
