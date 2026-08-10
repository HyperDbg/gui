// Package userlevel — ud.go
//
// Implements the user-mode debugger state machine. The C++ counterpart is
// libhyperdbg/code/debugger/user-level/ud.cpp; it tracks:
//   - whether the user debugger is initialised
//   - the active debugging process (pid, handle, entry-point breakpoint)
//   - the synchronization-event handle table used to wait for, and wake up,
//     paused debuggee threads
//
// In the Go rewrite the global state from the C side (g_ActiveProcessDebuggingState,
// g_IsUserDebuggerInitialized, g_UserSyncronizationObjectsHandleTable) is
// owned by the UdState struct so that multiple debuggers can coexist
// (GUI/MCP requirement).
package userlevel

import (
	"sync"
)

// UdState owns the user-mode debugger state. All fields are guarded by mu.
type UdState struct {
	mu          sync.Mutex
	initialised bool
	active      ActiveDebuggingProcess
	syncObjects [MaxSynchronisationObjects]SyncEventState
}

// ActiveDebuggingProcess mirrors the C ACTIVE_DEBUGGING_PROCESS: the process
// the user is currently debugging.
type ActiveDebuggingProcess struct {
	ProcessId      uint32
	ProcessHandle  uintptr
	ThreadId       uint32
	ThreadHandle   uintptr
	IsPaused       bool
	IsOnWaitingMtc bool
}

// SyncEventState mirrors DEBUGGER_SYNCRONIZATION_EVENTS_STATE: each slot is
// either free (IsOnWaitingState=false) or held by a thread waiting on
// EventHandle.
type SyncEventState struct {
	IsOnWaitingState bool
	EventHandle      uintptr // windows event handle (CreateEvent)
}

// MaxSynchronisationObjects mirrors DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS.
const MaxSynchronisationObjects = 64

// NewUdState returns a zero-initialised user-mode debugger state.
func NewUdState() *UdState {
	return &UdState{}
}

// Initialise sets up the synchronization event handle table. Mirrors
// UdInitializeUserDebugger. Idempotent — calling it twice is a no-op.
func (u *UdState) Initialise() {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.initialised {
		return
	}
	for i := range u.syncObjects {
		u.syncObjects[i] = SyncEventState{
			IsOnWaitingState: false,
			EventHandle:      0, // created lazily by the platform layer
		}
	}
	u.initialised = true
}

// IsInitialised reports whether Initialise has been called.
func (u *UdState) IsInitialised() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.initialised
}

// ActiveProcess returns the currently debugged process (zero-value if none).
func (u *UdState) ActiveProcess() ActiveDebuggingProcess {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.active
}

// SetActiveProcess records the process being debugged.
func (u *UdState) SetActiveProcess(p ActiveDebuggingProcess) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.active = p
}

// ClearActiveProcess removes the active process record.
func (u *UdState) ClearActiveProcess() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.active = ActiveDebuggingProcess{}
}

// SetPaused marks the active process as paused/running.
func (u *UdState) SetPaused(paused bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.active.IsPaused = paused
}

// AcquireSyncSlot finds a free sync slot, marks it as waiting, and returns
// its index. Returns -1 if all slots are in use.
func (u *UdState) AcquireSyncSlot() int {
	u.mu.Lock()
	defer u.mu.Unlock()
	for i := range u.syncObjects {
		if !u.syncObjects[i].IsOnWaitingState {
			u.syncObjects[i].IsOnWaitingState = true
			return i
		}
	}
	return -1
}

// ReleaseSyncSlot frees a previously acquired sync slot.
func (u *UdState) ReleaseSyncSlot(idx int) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if idx < 0 || idx >= len(u.syncObjects) {
		return
	}
	u.syncObjects[idx].IsOnWaitingState = false
}
