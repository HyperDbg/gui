package debugger

import (
	"fmt"
	"sync"

	"golang.org/x/sys/windows"
)

const (
	MaximumSyncObjects = 64
	MaximumInstrSize   = 16
)

type PausingReason int

const (
	PausingReasonGeneralThreadIntercepted PausingReason = iota
	PausingReasonDebuggeeStartingModuleLoaded
	PausingReasonBreakpointHit
	PausingReasonException
	PausingReasonStep
)

type UserDebuggingPacket struct {
	ProcessId             uint32
	ThreadId              uint32
	Rip                   uint64
	Rflags                uint64
	Is32Bit               bool
	ReadInstructionLen    uint8
	InstructionBytesOnRip [MaximumInstrSize]byte
	PausingReason         PausingReason
	ProcessDebuggingToken any
}

type SyncronizationEventState struct {
	IsOnWaitingState bool
	EventHandle      windows.Handle
}

type UserProvider struct {
	mu                        sync.RWMutex
	isInitialized             bool
	activeProcess             UserActiveDebuggingProcess
	syncObjects               [MaximumSyncObjects]SyncronizationEventState
	isUserDebuggerInitialized bool
}

type UserActiveDebuggingProcess struct {
	IsActive         bool
	IsPaused         bool
	ProcessId        uint32
	ThreadId         uint32
	Rip              uint64
	Is32Bit          bool
	InstructionBytes [MaximumInstrSize]byte
	DebuggingToken   any
}

func NewUserProvider() *UserProvider {
	ud := &UserProvider{
		isInitialized:             false,
		isUserDebuggerInitialized: false,
	}

	for i := range MaximumSyncObjects {
		ud.syncObjects[i] = SyncronizationEventState{
			IsOnWaitingState: false,
			EventHandle:      windows.InvalidHandle,
		}
	}

	return ud
}

func (ud *UserProvider) Initialize() error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if ud.isUserDebuggerInitialized {
		return nil
	}

	for i := range MaximumSyncObjects {
		event, err := windows.CreateEvent(nil, 0, 0, nil)
		if err != nil {
			return fmt.Errorf("failed to create event: %w", err)
		}
		ud.syncObjects[i].EventHandle = event
		ud.syncObjects[i].IsOnWaitingState = false
	}

	ud.isUserDebuggerInitialized = true
	return nil
}

func (ud *UserProvider) Uninitialize() error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.isUserDebuggerInitialized {
		return nil
	}

	ud.RemoveActiveDebuggingProcess()

	for i := range MaximumSyncObjects {
		if ud.syncObjects[i].EventHandle != windows.InvalidHandle {
			if ud.syncObjects[i].IsOnWaitingState {
				ud.ReceivedUserResponse(i)
			}
			windows.CloseHandle(ud.syncObjects[i].EventHandle)
			ud.syncObjects[i].EventHandle = windows.InvalidHandle
		}
	}

	ud.isUserDebuggerInitialized = false
	return nil
}

func (ud *UserProvider) SetActiveDebuggingProcess(process UserActiveDebuggingProcess) {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.activeProcess = process
}

func (ud *UserProvider) GetActiveDebuggingProcess() UserActiveDebuggingProcess {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.activeProcess
}

func (ud *UserProvider) RemoveActiveDebuggingProcess() {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.activeProcess = UserActiveDebuggingProcess{}
}

func (ud *UserProvider) ContinueProcess(processId uint32) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.activeProcess.IsActive {
		return fmt.Errorf("no active process")
	}

	if ud.activeProcess.ProcessId != processId {
		return fmt.Errorf("process id mismatch")
	}

	ud.activeProcess.IsPaused = false
	return nil
}

func (ud *UserProvider) PauseProcess(processId uint32) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.activeProcess.IsActive {
		return fmt.Errorf("no active process")
	}

	if ud.activeProcess.ProcessId != processId {
		return fmt.Errorf("process id mismatch")
	}

	ud.activeProcess.IsPaused = true
	return nil
}

func (ud *UserProvider) StepIn(processId uint32) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.activeProcess.IsActive {
		return fmt.Errorf("no active process")
	}

	if !ud.activeProcess.IsPaused {
		return fmt.Errorf("process is not paused")
	}

	ud.activeProcess.IsPaused = false
	return nil
}

func (ud *UserProvider) StepOver(processId uint32) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.activeProcess.IsActive {
		return fmt.Errorf("no active process")
	}

	if !ud.activeProcess.IsPaused {
		return fmt.Errorf("process is not paused")
	}

	ud.activeProcess.IsPaused = false
	return nil
}

func (ud *UserProvider) HandleUserDebuggerPausing(packet *UserDebuggingPacket) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	switch packet.PausingReason {
	case PausingReasonDebuggeeStartingModuleLoaded:
	case PausingReasonGeneralThreadIntercepted:
	case PausingReasonBreakpointHit:
	case PausingReasonException:
	case PausingReasonStep:
	}

	if !ud.activeProcess.IsPaused || packet.ThreadId == ud.activeProcess.ThreadId {
		ud.SetActiveDebuggingProcess(UserActiveDebuggingProcess{
			IsActive:       true,
			IsPaused:       true,
			ProcessId:      packet.ProcessId,
			ThreadId:       packet.ThreadId,
			Rip:            packet.Rip,
			Is32Bit:        packet.Is32Bit,
			DebuggingToken: packet.ProcessDebuggingToken,
		})
	}

	return nil
}

func (ud *UserProvider) WaitForUserResponse(index int) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if index < 0 || index >= MaximumSyncObjects {
		return fmt.Errorf("invalid sync object index")
	}

	ud.syncObjects[index].IsOnWaitingState = true
	return nil
}

func (ud *UserProvider) ReceivedUserResponse(index int) error {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if index < 0 || index >= MaximumSyncObjects {
		return fmt.Errorf("invalid sync object index")
	}

	if ud.syncObjects[index].EventHandle != windows.InvalidHandle {
		windows.SetEvent(ud.syncObjects[index].EventHandle)
	}
	ud.syncObjects[index].IsOnWaitingState = false
	return nil
}
