package windows

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PROCESS_ALL_ACCESS            = 0x1F0FFF
	PROCESS_CREATE_THREAD         = 0x0002
	PROCESS_QUERY_INFORMATION     = 0x0400
	PROCESS_VM_OPERATION          = 0x0008
	PROCESS_VM_WRITE              = 0x0020
	PROCESS_VM_READ               = 0x0010
	THREAD_GET_CONTEXT            = 0x0008
	THREAD_SET_CONTEXT            = 0x0010
	THREAD_QUERY_INFORMATION      = 0x0040
	THREAD_SUSPEND_RESUME         = 0x0002
	INFINITE                      = 0xFFFFFFFF
	DEBUG_PROCESS                 = 0x00000001
	DEBUG_ONLY_THIS_PROCESS       = 0x00000002
	CREATE_NEW_CONSOLE            = 0x00000010
	STARTF_USESTDHANDLES          = 0x00000100
	STARTF_USESHOWWINDOW          = 0x00000001
	SW_HIDE                       = 0
	SW_SHOW                       = 5
	DBG_CONTINUE                  = 0x00010002
	DBG_EXCEPTION_NOT_HANDLED     = 0x80010001
	EXCEPTION_DEBUG_EVENT         = 1
	CREATE_THREAD_DEBUG_EVENT     = 2
	CREATE_PROCESS_DEBUG_EVENT    = 3
	EXIT_THREAD_DEBUG_EVENT       = 4
	EXIT_PROCESS_DEBUG_EVENT      = 5
	LOAD_DLL_DEBUG_EVENT          = 6
	UNLOAD_DLL_DEBUG_EVENT        = 7
	OUTPUT_DEBUG_STRING_EVENT     = 8
	EXCEPTION_BREAKPOINT          = 0x80000003
	EXCEPTION_SINGLE_STEP         = 0x80000004
	EXCEPTION_ACCESS_VIOLATION    = 0xC0000005
	EXCEPTION_ILLEGAL_INSTRUCTION = 0xC000001D
	CONTEXT_FULL                  = 0x10007
	CONTEXT_DEBUG_REGISTERS       = 0x00010000
	PAGE_EXECUTE_READ             = 0x20
)

type Handle uintptr

type DebugEventCode uint32

// EXCEPTION_RECORD64 - https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-exception_record
type ExceptionRecord struct {
	ExceptionCode        uint32
	ExceptionFlags       uint32
	ExceptionRecord      uint64
	ExceptionAddress     uint64
	NumberParameters     uint32
	_UnusedAlignment     uint32
	ExceptionInformation [15]uint64
}

// EXCEPTION_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-exception_debug_info
type ExceptionDebugInfo struct {
	ExceptionRecord ExceptionRecord
	FirstChance     uint32
}

// CREATE_THREAD_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-create_thread_debug_info
type CreateThreadDebugInfo struct {
	ThreadHandle    Handle
	ThreadLocalBase uint64
	StartAddress    uint64
}

// CREATE_PROCESS_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-create_process_debug_info
type CreateProcessDebugInfo struct {
	FileHandle          uint32
	_                   [4]byte // Padding to 8 bytes
	ProcessHandle       uint32
	_                   [4]byte // Padding to 8 bytes
	ThreadHandle        uint32
	_                   [4]byte // Padding to 8 bytes
	BaseOfImage         uint32
	_                   [4]byte // Padding to 8 bytes
	DebugInfoFileOffset uint32
	DebugInfoSize       uint32
	ThreadLocalBase     uint32
	_                   [4]byte // Padding to 8 bytes
	StartAddress        uint32
	_                   [4]byte // Padding to 8 bytes
	ImageName           uint32
	_                   [4]byte // Padding to 8 bytes
	Unicode             uint16
	_                   [6]byte  // Padding to 8 bytes
	_                   [30]byte // Padding to make struct 96 bytes (66 + 30 = 96)
}

// EXIT_THREAD_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-exit_thread_debug_info
type ExitThreadDebugInfo struct {
	ExitCode uint32
}

// EXIT_PROCESS_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-exit_process_debug_info
type ExitProcessDebugInfo struct {
	ExitCode uint32
}

// LOAD_DLL_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-load_dll_debug_info
type LoadDllDebugInfo struct {
	FileHandle          Handle
	BaseOfDll           uint64
	DebugInfoFileOffset uint32
	DebugInfoSize       uint32
	ImageName           uint16
	Unicode             uint16
	_                   [60]byte // Padding to make the union 96 bytes
}

// UNLOAD_DLL_DEBUG_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-unload_dll_debug_info
type UnloadDllDebugInfo struct {
	BaseOfDll uint64
}

// OUTPUT_DEBUG_STRING_INFO - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-output_debug_string_info
type OutputDebugStringInfo struct {
	DebugString       *uint16
	Unicode           uint16
	DebugStringLength uint16
}

// DEBUG_EVENT - https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-debug_event
// 布局: offset 0: DebugEventCode, offset 4: ProcessId, offset 8: ThreadId, offset 12: Union
// 注意：Windows API 中 Union 从 offset 8 开始（但 64 位系统可能需要对齐）
type DebugEvent struct {
	data [172]byte // 固定大小：4 + 4 + 4 + 160 = 172
}

// 字段访问器 - 直接访问 data 数组中的特定偏移位置
func (e *DebugEvent) ProcessId() uint32 {
	return binary.LittleEndian.Uint32(e.data[4:8])
}

func (e *DebugEvent) ThreadId() uint32 {
	return binary.LittleEndian.Uint32(e.data[8:12])
}

func (e *DebugEvent) DebugEventCode() uint32 {
	return binary.LittleEndian.Uint32(e.data[0:4])
}

func (e *DebugEvent) Exception() *ExceptionDebugInfo {
	// Union 从 offset 16 开始（根据实际测试调整）
	info, err := DecodeExceptionDebugInfo(e.data[16:172])
	if err != nil {
		return &ExceptionDebugInfo{}
	}
	return info
}

func (e *DebugEvent) CreateThread() *CreateThreadDebugInfo {
	return (*CreateThreadDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) CreateProcess() *CreateProcessDebugInfo {
	return (*CreateProcessDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) ExitThread() *ExitThreadDebugInfo {
	return (*ExitThreadDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) ExitProcess() *ExitProcessDebugInfo {
	return (*ExitProcessDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) LoadDll() *LoadDllDebugInfo {
	return (*LoadDllDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) UnloadDll() *UnloadDllDebugInfo {
	return (*UnloadDllDebugInfo)(unsafe.Pointer(&e.data[16]))
}

func (e *DebugEvent) OutputString() *OutputDebugStringInfo {
	return (*OutputDebugStringInfo)(unsafe.Pointer(&e.data[16]))
}

// M128A - https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-m128a
type M128A struct {
	Low  uint64
	High int64
}

// FLOATING_SAVE_AREA - https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-floating_save_area
type FloatingSaveArea struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  [80]byte
	Cr0NpxState   uint32
}

// CONTEXT - https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-arm64_ntcontext
type Context struct {
	P1Home               uint64
	P2Home               uint64
	P3Home               uint64
	P4Home               uint64
	P5Home               uint64
	P6Home               uint64
	ContextFlags         uint32
	MxCsr                uint32
	SegCs                uint16
	SegDs                uint16
	SegEs                uint16
	SegFs                uint16
	SegGs                uint16
	SegSs                uint16
	EFlags               uint32
	Dr0                  uint64
	Dr1                  uint64
	Dr2                  uint64
	Dr3                  uint64
	Dr6                  uint64
	Dr7                  uint64
	Rax                  uint64
	Rcx                  uint64
	Rdx                  uint64
	Rbx                  uint64
	Rsp                  uint64
	Rbp                  uint64
	Rsi                  uint64
	Rdi                  uint64
	R8                   uint64
	R9                   uint64
	R10                  uint64
	R11                  uint64
	R12                  uint64
	R13                  uint64
	R14                  uint64
	R15                  uint64
	Rip                  uint64
	FltSave              FloatingSaveArea
	VectorRegister       [26]M128A
	VectorControl        uint64
	DebugControl         uint64
	LastBranchToRip      uint64
	LastBranchFromRip    uint64
	LastExceptionToRip   uint64
	LastExceptionFromRip uint64
}

// STARTUPINFO - https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfoa
type StartupInfo struct {
	Cb            uint32
	Reserved      *uint16
	Desktop       *uint16
	Title         *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	CbReserved2   uint16
	Reserved2     [2]uintptr
	StdInput      Handle
	StdOutput     Handle
	StdError      Handle
}

// PROCESS_INFORMATION - https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-process_information
type ProcessInformation struct {
	Process   Handle
	Thread    Handle
	ProcessId uint32
	ThreadId  uint32
}

// MEMORY_BASIC_INFORMATION - https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/ns-memoryapi-memory_basic_information
type MemoryBasicInformation struct {
	BaseAddress       uint64
	AllocationBase    uint64
	AllocationProtect uint32
	RegionSize        uint64
	State             uint32
	Protect           uint32
	Type              uint32
}

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procDebugActiveProcess       = kernel32.NewProc("DebugActiveProcess")
	procDebugActiveProcessStop   = kernel32.NewProc("DebugActiveProcessStop")
	procWaitForDebugEvent        = kernel32.NewProc("WaitForDebugEvent")
	procContinueDebugEvent       = kernel32.NewProc("ContinueDebugEvent")
	procCreateProcessW           = kernel32.NewProc("CreateProcessW")
	procTerminateProcess         = kernel32.NewProc("TerminateProcess")
	procReadProcessMemory        = kernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory       = kernel32.NewProc("WriteProcessMemory")
	procVirtualQueryEx           = kernel32.NewProc("VirtualQueryEx")
	procVirtualProtectEx         = kernel32.NewProc("VirtualProtectEx")
	procGetThreadContext         = kernel32.NewProc("GetThreadContext")
	procSetThreadContext         = kernel32.NewProc("SetThreadContext")
	procOpenThread               = kernel32.NewProc("OpenThread")
	procSuspendThread            = kernel32.NewProc("SuspendThread")
	procResumeThread             = kernel32.NewProc("ResumeThread")
	procCloseHandle              = kernel32.NewProc("CloseHandle")
	procGetModuleFileNameExW     = kernel32.NewProc("GetModuleFileNameExW")
	procGetModuleBaseNameW       = kernel32.NewProc("GetModuleBaseNameW")
	procEnumProcessModules       = kernel32.NewProc("EnumProcessModules")
	procSymInitialize            = kernel32.NewProc("SymInitialize")
	procSymCleanup               = kernel32.NewProc("SymCleanup")
	procSymLoadModuleExW         = kernel32.NewProc("SymLoadModuleExW")
	procSymUnloadModule64        = kernel32.NewProc("SymUnloadModule64")
	procSymFromAddr              = kernel32.NewProc("SymFromAddr")
	procSymGetModuleInfo64       = kernel32.NewProc("SymGetModuleInfo64")
	procSymEnumSymbols           = kernel32.NewProc("SymEnumSymbols")
	procSymSetOptions            = kernel32.NewProc("SymSetOptions")
	procStackWalk64              = kernel32.NewProc("StackWalk64")
	procSymFunctionTableAccess64 = kernel32.NewProc("SymFunctionTableAccess64")
	procSymGetModuleBase64       = kernel32.NewProc("SymGetModuleBase64")
)

func DebugActiveProcess(pid uint32) (Handle, error) {
	ret, _, err := procDebugActiveProcess.Call(uintptr(pid))
	if ret == 0 {
		return 0, fmt.Errorf("DebugActiveProcess failed: %v", err)
	}

	handle, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}

	return handle, nil
}

func DebugActiveProcessStop(pid uint32) error {
	ret, _, err := procDebugActiveProcessStop.Call(uintptr(pid))
	if ret == 0 {
		return fmt.Errorf("DebugActiveProcessStop failed: %v", err)
	}
	return nil
}

func WaitForDebugEvent() (*DebugEvent, error) {
	var event DebugEvent
	ret, _, err := procWaitForDebugEvent.Call(
		uintptr(unsafe.Pointer(&event)),
		uintptr(100), // 100ms timeout like x64dbg
	)
	if ret == 0 {
		if errno, ok := err.(syscall.Errno); ok && errno == 0x79 {
			return nil, nil // timeout, no event
		}
		return nil, fmt.Errorf("WaitForDebugEvent failed: %v", err)
	}
	return &event, nil
}

func ContinueDebugEvent(pid uint32, tid uint32, status uint32) error {
	ret, _, err := procContinueDebugEvent.Call(
		uintptr(pid),
		uintptr(tid),
		uintptr(status),
	)
	if ret == 0 {
		return fmt.Errorf("ContinueDebugEvent failed: %v", err)
	}
	return nil
}

func CreateProcessDebug(exePath string, cmdLine string) (Handle, uint32, Handle, uint32, error) {
	var si syscall.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Flags = syscall.STARTF_USESHOWWINDOW
	si.ShowWindow = syscall.SW_SHOW

	var pi syscall.ProcessInformation

	var exePathPtr *uint16
	var cmdLinePtr *uint16

	if exePath != "" {
		exePathPtr, _ = syscall.UTF16PtrFromString(exePath)
	}
	if cmdLine != "" {
		cmdLinePtr, _ = syscall.UTF16PtrFromString(cmdLine)
	}

	err := syscall.CreateProcess(
		exePathPtr,
		cmdLinePtr,
		nil,
		nil,
		false,
		0,
		nil,
		nil,
		&si,
		&pi,
	)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("CreateProcess failed: %v", err)
	}

	handle, err := DebugActiveProcess(pi.ProcessId)
	if err != nil {
		syscall.CloseHandle(pi.Process)
		syscall.CloseHandle(pi.Thread)
		return 0, 0, 0, 0, err
	}

	syscall.CloseHandle(pi.Process)

	return handle, pi.ProcessId, Handle(pi.Thread), pi.ThreadId, nil
}

func CreateProcessNoDebug(exePath string, cmdLine string) (Handle, uint32, error) {
	var si syscall.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Flags = syscall.STARTF_USESHOWWINDOW
	si.ShowWindow = syscall.SW_SHOW

	var pi syscall.ProcessInformation

	var exePathPtr *uint16
	var cmdLinePtr *uint16

	if exePath != "" {
		exePathPtr, _ = syscall.UTF16PtrFromString(exePath)
	}
	if cmdLine != "" {
		cmdLinePtr, _ = syscall.UTF16PtrFromString(cmdLine)
	}

	err := syscall.CreateProcess(
		exePathPtr,
		cmdLinePtr,
		nil,
		nil,
		false,
		0,
		nil,
		nil,
		&si,
		&pi,
	)
	if err != nil {
		return 0, 0, fmt.Errorf("CreateProcess failed: %v", err)
	}

	return Handle(pi.Process), pi.ProcessId, nil
}

func TerminateProcess(handle Handle, exitCode uint32) error {
	ret, _, err := procTerminateProcess.Call(
		uintptr(handle),
		uintptr(exitCode),
	)
	if ret == 0 {
		return fmt.Errorf("TerminateProcess failed: %v", err)
	}
	return nil
}

func ReadProcessMemory(handle Handle, address uint64, size uint32) ([]byte, error) {
	buffer := make([]byte, size)
	var bytesRead uint32

	ret, _, err := procReadProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("ReadProcessMemory failed: %v", err)
	}

	return buffer[:bytesRead], nil
}

func WriteProcessMemory(handle Handle, address uint64, data []byte) (uint32, error) {
	var bytesWritten uint32

	ret, _, err := procWriteProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	if ret == 0 {
		return 0, fmt.Errorf("WriteProcessMemory failed: %v", err)
	}

	return bytesWritten, nil
}

func VirtualQueryEx(handle Handle, address uint64) (*MemoryBasicInformation, error) {
	var mbi MemoryBasicInformation

	ret, _, err := procVirtualQueryEx.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&mbi)),
		uintptr(unsafe.Sizeof(mbi)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("VirtualQueryEx failed: %v", err)
	}

	return &mbi, nil
}

func VirtualProtectEx(handle Handle, address uint64, size uint64, newProtect uint32, oldProtect *uint32) error {
	ret, _, err := procVirtualProtectEx.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(size),
		uintptr(newProtect),
		uintptr(unsafe.Pointer(oldProtect)),
	)
	if ret == 0 {
		return fmt.Errorf("VirtualProtectEx failed: %v", err)
	}
	return nil
}

func GetThreadContext(threadHandle Handle) (*Context, error) {
	var ctx Context
	ctx.ContextFlags = CONTEXT_FULL | CONTEXT_DEBUG_REGISTERS

	ret, _, err := procGetThreadContext.Call(
		uintptr(threadHandle),
		uintptr(unsafe.Pointer(&ctx)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("GetThreadContext failed: %v", err)
	}

	return &ctx, nil
}

func SetThreadContext(threadHandle Handle, ctx *Context) error {
	ret, _, err := procSetThreadContext.Call(
		uintptr(threadHandle),
		uintptr(unsafe.Pointer(ctx)),
	)
	if ret == 0 {
		return fmt.Errorf("SetThreadContext failed: %v", err)
	}
	return nil
}

func OpenThread(desiredAccess uint32, inheritHandle bool, threadId uint32) (Handle, error) {
	var inherit uint32
	if inheritHandle {
		inherit = 1
	}

	ret, _, err := procOpenThread.Call(
		uintptr(desiredAccess),
		uintptr(inherit),
		uintptr(threadId),
	)
	if ret == 0 {
		return 0, fmt.Errorf("OpenThread failed: %v", err)
	}

	return Handle(ret), nil
}

func SuspendThread(threadHandle Handle) (uint32, error) {
	ret, _, err := procSuspendThread.Call(uintptr(threadHandle))
	if ret == 0xFFFFFFFF {
		return 0, fmt.Errorf("SuspendThread failed: %v", err)
	}
	return uint32(ret), nil
}

func ResumeThread(threadHandle Handle) (uint32, error) {
	ret, _, err := procResumeThread.Call(uintptr(threadHandle))
	if ret == 0xFFFFFFFF {
		return 0, fmt.Errorf("ResumeThread failed: %v", err)
	}
	return uint32(ret), nil
}

func CloseHandle(handle Handle) error {
	ret, _, err := procCloseHandle.Call(uintptr(handle))
	if ret == 0 {
		return fmt.Errorf("CloseHandle failed: %v", err)
	}
	return nil
}

func OpenProcess(desiredAccess uint32, inheritHandle bool, pid uint32) (Handle, error) {
	ret, err := syscall.OpenProcess(desiredAccess, inheritHandle, pid)
	if ret == 0 {
		return 0, fmt.Errorf("OpenProcess failed: %v", err)
	}

	return Handle(ret), nil
}
