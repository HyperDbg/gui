package user_level

import (
	"fmt"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/sdk"
	"golang.org/x/sys/windows"
)

const (
	MaximumSyncObjects = 64
	MaximumInstrSize   = 16

	SyncObjectUserDebuggerIsDebuggerRunning = 0
)

// SyncronizationEventState manages synchronization event objects for user debugger
// This is Go-specific (not from C++ headers) for Windows event handle management
type SyncronizationEventState struct {
	IsOnWaitingState bool
	EventHandle      windows.Handle
}

// ActiveDebuggingProcess represents the current debugging process state
// Source: HyperDbg/libhyperdbg/header/ud.h -> _ACTIVE_DEBUGGING_PROCESS
type ActiveDebuggingProcess struct {
	IsActive              bool
	ProcessDebuggingToken uint64
	ProcessId             uint32
	ThreadId              uint32
	IsPaused              bool
	Registers             sdk.GUEST_REGS
	Context               uint64
	Rip                   uint64
	Is32Bit               bool
	InstructionBytesOnRip [MaximumInstrSize]byte
}

type UserDebugger struct {
	isInitialized               bool
	activeProcess               ActiveDebuggingProcess
	syncObjects                 [MaximumSyncObjects]SyncronizationEventState
	processIdOfLatestStarting   uint32
	resultOfEvaluatedExpression uint64
	errorStateOfEvaluatedExpr   uint32
	mu                          sync.RWMutex
}

func NewUserDebugger() *UserDebugger {
	return &UserDebugger{
		isInitialized: false,
		syncObjects:   [MaximumSyncObjects]SyncronizationEventState{},
	}
}

func (ud *UserDebugger) Initialize() {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.isInitialized {
		for i := range MaximumSyncObjects {
			eventHandle, _ := windows.CreateEvent(nil, 0, 0, nil)
			ud.syncObjects[i] = SyncronizationEventState{
				IsOnWaitingState: false,
				EventHandle:      eventHandle,
			}
		}
		ud.isInitialized = true
	}
}

func (ud *UserDebugger) Uninitialize() {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if ud.isInitialized {
		ud.activeProcess.IsActive = false

		for i := range MaximumSyncObjects {
			if ud.syncObjects[i].EventHandle != windows.InvalidHandle {
				if ud.syncObjects[i].IsOnWaitingState {
					windows.SetEvent(ud.syncObjects[i].EventHandle)
				}
				windows.CloseHandle(ud.syncObjects[i].EventHandle)
				ud.syncObjects[i].EventHandle = windows.InvalidHandle
			}
		}
		ud.isInitialized = false
	}
}

func (ud *UserDebugger) SetActiveDebuggingProcess(
	debuggingToken uint64,
	rip uint64,
	processId uint32,
	threadId uint32,
	is32Bit bool,
	isPaused bool,
	instructionBytes [MaximumInstrSize]byte,
) {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	ud.activeProcess.IsActive = true
	ud.activeProcess.ProcessDebuggingToken = debuggingToken
	ud.activeProcess.ProcessId = processId
	ud.activeProcess.ThreadId = threadId
	ud.activeProcess.IsPaused = isPaused
	ud.activeProcess.Rip = rip
	ud.activeProcess.Is32Bit = is32Bit
	copy(ud.activeProcess.InstructionBytesOnRip[:], instructionBytes[:])
}

func (ud *UserDebugger) RemoveActiveDebuggingProcess() {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.activeProcess.IsActive = false
}

func (ud *UserDebugger) GetActiveDebuggingProcess() ActiveDebuggingProcess {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.activeProcess
}

func (ud *UserDebugger) ReceivedUserResponse(objectIndex int) {
	if objectIndex >= 0 && objectIndex < MaximumSyncObjects {
		ud.syncObjects[objectIndex].IsOnWaitingState = false
		windows.SetEvent(ud.syncObjects[objectIndex].EventHandle)
	}
}

func (ud *UserDebugger) IsInitialized() bool {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.isInitialized
}

func PrintError() {
	errNum := windows.GetLastError()
	buf := make([]uint16, 256)
	n, _ := windows.FormatMessage(windows.FORMAT_MESSAGE_FROM_SYSTEM|windows.FORMAT_MESSAGE_IGNORE_INSERTS,
		0, uint32(errNum.(syscall.Errno)), 0x0409, buf, nil)
	if n > 0 {
		msg := windows.UTF16ToString(buf[:n])
		mylog.Warning("Win32 API failed with error", "code", fmt.Sprintf("%x", errNum), "msg", msg)
	} else {
		mylog.Warning("Win32 API failed with error", "code", fmt.Sprintf("%x", errNum))
	}
}

func ListProcessThreads(ownerPID uint32) ([]uint32, error) {
	threadSnap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPTHREAD, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(threadSnap)

	var te32 windows.ThreadEntry32
	te32.Size = uint32(unsafe.Sizeof(te32))

	err = windows.Thread32First(threadSnap, &te32)
	if err != nil {
		PrintError()
		return nil, err
	}

	var threadIDs []uint32
	for {
		if te32.OwnerProcessID == ownerPID {
			threadIDs = append(threadIDs, te32.ThreadID)
		}
		err = windows.Thread32Next(threadSnap, &te32)
		if err != nil {
			break
		}
	}
	return threadIDs, nil
}

func CheckThreadByProcessId(pid, tid uint32) (bool, error) {
	threadSnap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPTHREAD, 0)
	if err != nil {
		return false, err
	}
	defer windows.CloseHandle(threadSnap)

	var te32 windows.ThreadEntry32
	te32.Size = uint32(unsafe.Sizeof(te32))

	err = windows.Thread32First(threadSnap, &te32)
	if err != nil {
		PrintError()
		return false, err
	}

	for {
		if te32.OwnerProcessID == pid && te32.ThreadID == tid {
			return true, nil
		}
		err = windows.Thread32Next(threadSnap, &te32)
		if err != nil {
			break
		}
	}
	return false, nil
}

func CreateSuspendedProcess(fileName string, commandLine string) (*windows.ProcessInformation, error) {
	fileNamePtr, _ := windows.UTF16PtrFromString(fileName)
	var commandLinePtr *uint16
	if commandLine != "" {
		commandLinePtr, _ = windows.UTF16PtrFromString(commandLine)
	}

	startupInfo := windows.StartupInfo{Cb: uint32(unsafe.Sizeof(windows.StartupInfo{}))}
	var procInfo windows.ProcessInformation

	err := windows.CreateProcess(
		fileNamePtr,
		commandLinePtr,
		nil,
		nil,
		false,
		windows.CREATE_SUSPENDED|windows.CREATE_NEW_CONSOLE,
		nil,
		nil,
		&startupInfo,
		&procInfo,
	)
	if err != nil {
		return nil, fmt.Errorf("start process failed (%w)", err)
	}
	return &procInfo, nil
}

func (ud *UserDebugger) AttachToProcess(
	targetPid uint32,
	targetFileAddress string,
	commandLine string,
	runCallbackAtFirstInstruction bool,
	sendIoctl func(*sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS) error,
) error {
	ud.Initialize()

	attachReq := &sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{}
	var procInfo *windows.ProcessInformation

	attachReq.IsStartingNewProcess = (targetFileAddress == "")

	attachReq.CheckCallbackAtFirstInstruction = runCallbackAtFirstInstruction
	attachReq.Action = sdk.DebuggerAttachDetachUserModeProcessActionAttach

	if attachReq.IsStartingNewProcess {
		var err error
		procInfo, err = CreateSuspendedProcess(targetFileAddress, commandLine)
		if err != nil {
			return err
		}
		attachReq.ProcessId = procInfo.ProcessId
		attachReq.ThreadId = procInfo.ThreadId
	} else {
		attachReq.ProcessId = targetPid
	}

	err := sendIoctl(attachReq)
	if err != nil {
		return fmt.Errorf("ioctl failed: %w", err)
	}

	if attachReq.Result == uint64(sdk.DebuggerOperationWasSuccessful) {
		if !attachReq.IsStartingNewProcess {
			mylog.Info("successfully attached to the target process!")
			return nil
		}

		if procInfo.Thread != windows.InvalidHandle {
			_, _ = windows.ResumeThread(procInfo.Thread)
		} else {
			return fmt.Errorf("thread handle is empty")
		}

		for {
			attachReq.Action = sdk.DebuggerAttachDetachUserModeProcessActionRemoveHooks

			err := sendIoctl(attachReq)
			if err != nil {
				return fmt.Errorf("ioctl failed in remove hooks loop: %w", err)
			}

			if attachReq.Result == uint64(sdk.DebuggerOperationWasSuccessful) {
				break
			} else if attachReq.Result == uint64(sdk.DebuggerErrorUnableToRemoveHooksEntrypointNotReached) {
				time.Sleep(1 * time.Second)
				continue
			} else {
				fmt.Println(sdk.DebuggerErrorCode(attachReq.Result).String())
				return fmt.Errorf("remove hooks failed with result %d", attachReq.Result)
			}
		}

		ud.mu.Lock()
		ud.processIdOfLatestStarting = procInfo.ProcessId
		ud.mu.Unlock()

		return nil
	} else {
		fmt.Println(sdk.DebuggerErrorCode(attachReq.Result).String())
		return fmt.Errorf("attach failed with result %d", attachReq.Result)
	}
}

func TerminateProcessByPid(targetPid uint32) bool {
	handle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, targetPid)
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)

	err = windows.TerminateProcess(handle, 0)
	return err == nil
}

func DoesProcessExistByPid(targetPid uint32) bool {
	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, targetPid)
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)

	var exitCode uint32
	err = windows.GetExitCodeProcess(handle, &exitCode)
	if err != nil {
		return false
	}
	return exitCode == 259
}

func DoesProcessExistByHandle(targetHandle windows.Handle) bool {
	var exitCode uint32
	err := windows.GetExitCodeProcess(targetHandle, &exitCode)
	if err != nil {
		return false
	}
	return exitCode == 259
}

func (ud *UserDebugger) KillProcess(
	targetPid uint32,
	sendIoctl func(*sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS) error,
) error {
	killReq := &sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionKillProcess,
		ProcessId: targetPid,
	}

	if !DoesProcessExistByPid(targetPid) {
		return fmt.Errorf("process does not exist")
	}

	TerminateProcessByPid(targetPid)

	if !DoesProcessExistByPid(targetPid) {
		return nil
	}

	err := sendIoctl(killReq)
	if err != nil {
		return fmt.Errorf("ioctl kill failed: %w", err)
	}

	if killReq.Result == uint64(sdk.DebuggerOperationWasSuccessful) {
		return nil
	}
	fmt.Println(sdk.DebuggerErrorCode(killReq.Result).String())
	return fmt.Errorf("kill from kernel failed with result %d", killReq.Result)
}

func (ud *UserDebugger) DetachFromProcess(
	targetPid uint32,
	sendIoctl func(*sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS) error,
) error {
	detachReq := &sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionDetach,
		ProcessId: targetPid,
	}

	err := sendIoctl(detachReq)
	if err != nil {
		return fmt.Errorf("ioctl detach failed: %w", err)
	}

	if detachReq.Result == uint64(sdk.DebuggerOperationWasSuccessful) {
		ud.RemoveActiveDebuggingProcess()
		return nil
	}
	fmt.Println(sdk.DebuggerErrorCode(detachReq.Result).String())
	return fmt.Errorf("detach failed with result %d", detachReq.Result)
}

func (ud *UserDebugger) GetProcessIdOfLatestStarting() uint32 {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.processIdOfLatestStarting
}

func (ud *UserDebugger) SetResultOfEvaluatedExpression(value uint64, errorState uint32) {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.resultOfEvaluatedExpression = value
	ud.errorStateOfEvaluatedExpr = errorState
}

func (ud *UserDebugger) GetResultOfEvaluatedExpression() (uint64, uint32) {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.resultOfEvaluatedExpression, ud.errorStateOfEvaluatedExpr
}

func IsFileExist(filePath string) bool {
	attr, err := windows.GetFileAttributes(mylog.Check2(windows.UTF16PtrFromString(filePath)))
	if err != nil {
		return false
	}
	return attr&0xFFFFFFFF != 0xFFFFFFFF
}

// UdAttachToProcess is the exported function for attaching to a process
// This matches the C++ export: hyperdbg_u_start_process / hyperdbg_u_start_process_with_args
func UdAttachToProcess(targetPid uint32, path, commandLine string, runCallbackAtFirstInstruction bool) bool {
	ud := NewUserDebugger()
	ud.Initialize()

	err := ud.AttachToProcess(
		targetPid,
		path,
		commandLine,
		runCallbackAtFirstInstruction,
		nil, // sendIoctl - will need actual IOCTL implementation
	)
	return err == nil
}

// defaultUserDebugger is the global user debugger instance
var defaultUserDebugger *UserDebugger

// UdUninitializeUserDebugger uninitializes the user debugger
func UdUninitializeUserDebugger() {
	if defaultUserDebugger != nil {
		defaultUserDebugger.Uninitialize()
		defaultUserDebugger = nil
	}
}
