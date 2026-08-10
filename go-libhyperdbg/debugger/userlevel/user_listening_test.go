package userlevel

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

// packetInSendTimeout is the upper bound for a goroutine to deliver a test
// packet through a channel. It keeps tests from hanging forever if the
// listener never reads.
const packetInSendTimeout = 2 * time.Second

// TestListenerRunCancel verifies that cancelling the context causes Run to
// return immediately (with context.Canceled), even when no packets have been
// delivered.
func TestListenerRunCancel(t *testing.T) {
	t.Parallel()
	state := NewUdState()
	state.Initialise()
	lst := NewListener(state, nil)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan PausedPacket)

	done := make(chan error, 1)
	go func() { done <- lst.Run(ctx, ch) }()

	// Cancel and expect Run to return promptly.
	cancel()
	select {
	case err := <-done:
		if err != context.Canceled {
			t.Errorf("Run returned %v, want context.Canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after ctx cancel")
	}
}

// TestListenerRunChannelClose verifies that closing the packet channel causes
// Run to return nil (clean shutdown).
func TestListenerRunChannelClose(t *testing.T) {
	t.Parallel()
	state := NewUdState()
	state.Initialise()
	lst := NewListener(state, nil)

	ctx := context.Background()
	ch := make(chan PausedPacket)

	done := make(chan error, 1)
	go func() { done <- lst.Run(ctx, ch) }()

	close(ch)
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Run returned %v, want nil after channel close", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after channel close")
	}
}

// TestListenerDispatchPacket verifies that a packet sent on the channel is
// delivered to the Handler, and that the listener marks the UdState as
// paused before invoking the handler.
func TestListenerDispatchPacket(t *testing.T) {
	t.Parallel()
	state := NewUdState()
	state.Initialise()

	want := PausedPacket{
		ProcessId:      0xABC,
		ThreadId:       0xDEF,
		PausingReason:  PausingReasonGeneralThreadIntercepted,
		InstructionLen: 4,
	}

	got := make(chan PausedPacket, 1)
	sawPaused := make(chan bool, 1)
	handler := func(ctx context.Context, pkt PausedPacket) error {
		// The listener sets paused=true on the state before calling the
		// handler, so the handler must observe IsPaused=true.
		sawPaused <- state.ActiveProcess().IsPaused
		got <- pkt
		return nil
	}
	lst := NewListener(state, handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan PausedPacket, 1)

	go func() { _ = lst.Run(ctx, ch) }()
	t.Cleanup(cancel)

	select {
	case ch <- want:
	case <-time.After(packetInSendTimeout):
		t.Fatal("timed out sending packet to listener channel")
	}

	select {
	case pkt := <-got:
		if pkt != want {
			t.Errorf("handler received %+v, want %+v", pkt, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("handler was not invoked")
	}

	select {
	case paused := <-sawPaused:
		if !paused {
			t.Error("handler observed IsPaused=false, want true (listener must SetPaused before dispatching)")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("handler did not report IsPaused observation")
	}
}

// TestDefaultHandlerStartingModuleLoaded verifies that DefaultHandler emits a
// message containing "entrypoint" for the StartingModuleLoaded reason.
func TestDefaultHandlerStartingModuleLoaded(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	out := func(format string, args ...any) {
		// Mirror fmt.Printf semantics so the format string is exercised.
		buf.WriteString(fmt.Sprintf(format, args...))
	}
	h := DefaultHandler(out)

	pkt := PausedPacket{
		ProcessId:      1,
		ThreadId:       1,
		PausingReason:  PausingReasonStartingModuleLoaded,
		InstructionLen: 1, // valid length so the handler returns nil
	}
	if err := h(context.Background(), pkt); err != nil {
		t.Fatalf("DefaultHandler returned %v, want nil", err)
	}
	if !strings.Contains(buf.String(), "entrypoint") {
		t.Errorf("output = %q, want it to contain %q", buf.String(), "entrypoint")
	}
}

// TestDefaultHandlerGeneralThreadIntercepted verifies that DefaultHandler emits
// a message containing "intercepted" for the GeneralThreadIntercepted reason.
func TestDefaultHandlerGeneralThreadIntercepted(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	out := func(format string, args ...any) {
		buf.WriteString(fmt.Sprintf(format, args...))
	}
	h := DefaultHandler(out)

	pkt := PausedPacket{
		ProcessId:      0x100,
		ThreadId:       0x200,
		PausingReason:  PausingReasonGeneralThreadIntercepted,
		InstructionLen: 2, // valid length so the handler returns nil
	}
	if err := h(context.Background(), pkt); err != nil {
		t.Fatalf("DefaultHandler returned %v, want nil", err)
	}
	if !strings.Contains(buf.String(), "intercepted") {
		t.Errorf("output = %q, want it to contain %q", buf.String(), "intercepted")
	}
}
