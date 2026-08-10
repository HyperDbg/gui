package userlevel

import (
	"testing"
)

// TestUdStateInitialise verifies that a fresh UdState reports as not
// initialised, then reports initialised=true after Initialise is called.
func TestUdStateInitialise(t *testing.T) {
	t.Parallel()
	u := NewUdState()
	if u.IsInitialised() {
		t.Fatal("fresh UdState reports initialised, want false")
	}
	u.Initialise()
	if !u.IsInitialised() {
		t.Fatal("IsInitialised() = false after Initialise(), want true")
	}
}

// TestUdStateInitialiseIdempotent verifies that calling Initialise multiple
// times does not panic and leaves the state initialised.
func TestUdStateInitialiseIdempotent(t *testing.T) {
	t.Parallel()
	u := NewUdState()
	u.Initialise()
	// Subsequent calls must be no-ops, not panics.
	u.Initialise()
	u.Initialise()
	if !u.IsInitialised() {
		t.Fatal("IsInitialised() = false after repeated Initialise(), want true")
	}
}

// TestUdStateSetActiveProcess verifies that SetActiveProcess stores the
// process record so it can be read back by ActiveProcess, and that
// ClearActiveProcess zeroes it.
func TestUdStateSetActiveProcess(t *testing.T) {
	t.Parallel()
	u := NewUdState()
	want := ActiveDebuggingProcess{
		ProcessId:     1234,
		ThreadId:      5678,
		ProcessHandle: 0x1000,
		ThreadHandle:  0x2000,
	}
	u.SetActiveProcess(want)
	got := u.ActiveProcess()
	if got.ProcessId != want.ProcessId || got.ThreadId != want.ThreadId ||
		got.ProcessHandle != want.ProcessHandle || got.ThreadHandle != want.ThreadHandle {
		t.Errorf("ActiveProcess() = %+v, want %+v", got, want)
	}

	u.ClearActiveProcess()
	got = u.ActiveProcess()
	if got.ProcessId != 0 || got.ThreadId != 0 {
		t.Errorf("ActiveProcess() after clear = %+v, want zero value", got)
	}
}

// TestUdStateSetPaused verifies that SetPaused flips the IsPaused flag on the
// active process and that the state is observable via ActiveProcess.
func TestUdStateSetPaused(t *testing.T) {
	t.Parallel()
	u := NewUdState()
	u.SetActiveProcess(ActiveDebuggingProcess{ProcessId: 42})
	if u.ActiveProcess().IsPaused {
		t.Error("IsPaused = true initially, want false")
	}
	u.SetPaused(true)
	if !u.ActiveProcess().IsPaused {
		t.Error("IsPaused = false after SetPaused(true), want true")
	}
	u.SetPaused(false)
	if u.ActiveProcess().IsPaused {
		t.Error("IsPaused = true after SetPaused(false), want false")
	}
}

// TestUdStateSyncEvent verifies the sync-object handle table: AcquireSyncSlot
// marks a slot as waiting (IsOnWaitingState=true), and ReleaseSyncSlot frees
// it again. Mirrors the C handle-table round-trip.
func TestUdStateSyncEvent(t *testing.T) {
	t.Parallel()
	u := NewUdState()
	u.Initialise()

	// Acquire must return a valid slot index.
	idx := u.AcquireSyncSlot()
	if idx < 0 {
		t.Fatal("AcquireSyncSlot() = -1, want a valid slot index")
	}

	// The same slot cannot be re-acquired while held; the next acquire must
	// return a different index.
	idx2 := u.AcquireSyncSlot()
	if idx2 < 0 {
		t.Fatal("second AcquireSyncSlot() = -1, want a valid slot index")
	}
	if idx2 == idx {
		t.Errorf("second AcquireSyncSlot() = %d, want a different slot from %d", idx2, idx)
	}

	// Releasing the first slot must make it available again. Because the
	// table is scanned linearly from index 0, the next acquire should reuse
	// the freed lower index.
	u.ReleaseSyncSlot(idx)
	idx3 := u.AcquireSyncSlot()
	if idx3 != idx {
		t.Errorf("AcquireSyncSlot() after release = %d, want %d (reused freed slot)", idx3, idx)
	}

	// Releasing an out-of-range index must be a safe no-op (no panic).
	u.ReleaseSyncSlot(-1)
	u.ReleaseSyncSlot(MaxSynchronisationObjects)
	u.ReleaseSyncSlot(MaxSynchronisationObjects + 10)
}
