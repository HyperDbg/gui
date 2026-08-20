// Package userlevel — user_listening.go
//
// Implements the user-mode debugger event listener. The C++ counterpart is
// libhyperdbg/code/debugger/user-level/user-listening.cpp: it receives
// DEBUGGEE_UD_PAUSED_PACKET messages from the driver, formats them for the
// user, and signals the sync objects so the waiting command can resume.
//
// In the Go version the listener is a goroutine that reads from the
// comm.Device event channel (set up by core.Debugger) and dispatches each
// event to a Handler callback. The CLI provides a Handler that prints the
// pause info; GUI/MCP provide their own.
package userlevel

import (
	"fmt"
	"sync"
)

// PausingReason mirrors DEBUGGEE_PAUSING_REASON (subset relevant to the
// user-mode listener).
type PausingReason uint32

const (
	PausingReasonUnknown                  PausingReason = 0
	PausingReasonStartingModuleLoaded     PausingReason = 1
	PausingReasonGeneralThreadIntercepted PausingReason = 2
)

// PausedPacket mirrors the subset of DEBUGGEE_UD_PAUSED_PACKET the listener
// needs to format a message. The full struct lives in go-libhyperdbg/types.
type PausedPacket struct {
	ProcessId      uint32
	ThreadId       uint32
	PausingReason  PausingReason
	InstructionLen uint32
	Instruction    [16]byte
}

// Handler is called by the listener for every paused packet. Implementations
// write to the user output (CLI: stdout, GUI: widget, MCP: JSON channel).
// Returning an error stops the listener.
type Handler func(pkt PausedPacket) error

// Listener runs until ch is closed. Each paused packet read from ch is
// dispatched to handler. The listener is goroutine-safe; a single listener
// per UdState is the expected usage.
type Listener struct {
	mu      sync.Mutex
	handler Handler
	state   *UdState
}

// NewListener creates a listener for the given UdState. The handler is
// invoked for every packet; nil handler drops packets (useful for tests).
func NewListener(state *UdState, handler Handler) *Listener {
	return &Listener{state: state, handler: handler}
}

// Run blocks until ch is closed. For every packet read from ch it calls
// the handler; if the handler returns an error the listener returns it.
func (l *Listener) Run(ch <-chan PausedPacket) error {
	for pkt := range ch {
		l.state.SetPaused(true)
		if l.handler != nil {
			if err := l.handler(pkt); err != nil {
				return err
			}
		}
	}
	return nil
}

// DefaultHandler prints a CLI-style message for the packet, mirroring the
// ShowMessages calls in UdHandleUserDebuggerPausing.
func DefaultHandler(out func(format string, args ...any)) Handler {
	return func(pkt PausedPacket) error {
		switch pkt.PausingReason {
		case PausingReasonStartingModuleLoaded:
			out("the target module is loaded and a breakpoint is set to the entrypoint\n")
			out("press 'g' to reach to the entrypoint of the main module...\n")
		case PausingReasonGeneralThreadIntercepted:
			out("thread: %x from process: %x intercepted\n", pkt.ThreadId, pkt.ProcessId)
		}
		if pkt.InstructionLen == 0 || pkt.InstructionLen > 16 {
			out("err: invalid instruction length: %d\n", pkt.InstructionLen)
			return fmt.Errorf("invalid instruction length: %d", pkt.InstructionLen)
		}
		return nil
	}
}
