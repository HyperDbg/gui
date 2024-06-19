// Code generated by gengo. DO NOT EDIT.
package HPRDBGCTRL

import (
	"unsafe"
	"github.com/can1357/gengo/gengort"
)

const GengoLibraryName = "HPRDBGCTRL"

var GengoLibrary = gengort.NewLibrary(GengoLibraryName)

// @brief enum for reasons why debuggee is paused
type _DebuggeePausingReason int32

const (
	DEBUGGEE_PAUSING_REASON_NOT_PAUSED                            _DebuggeePausingReason = 0
	DEBUGGEE_PAUSING_REASON_PAUSE                                 _DebuggeePausingReason = 1
	DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER                 _DebuggeePausingReason = 2
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED                      _DebuggeePausingReason = 3
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_TRACKING_STEPPED             _DebuggeePausingReason = 4
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT      _DebuggeePausingReason = 5
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT  _DebuggeePausingReason = 6
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED                _DebuggeePausingReason = 7
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED             _DebuggeePausingReason = 8
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED              _DebuggeePausingReason = 9
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED   _DebuggeePausingReason = 10
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED              _DebuggeePausingReason = 11
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_STARTING_MODULE_LOADED       _DebuggeePausingReason = 12
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK          _DebuggeePausingReason = 13
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED   _DebuggeePausingReason = 14
	DEBUGGEE_PAUSING_REASON_HARDWARE_BASED_DEBUGGEE_GENERAL_BREAK _DebuggeePausingReason = 15
)

// @brief enum for requested action for HyperDbg packet
type _DebuggerRemotePacketRequestedAction int32

const (
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE                            _DebuggerRemotePacketRequestedAction = 1
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET           _DebuggerRemotePacketRequestedAction = 2
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DEBUGGER_VERSION                 _DebuggerRemotePacketRequestedAction = 3
	DEBUGGER_REMOTE_PACKET_PING_AND_SEND_SUPPORTED_VERSION                                _DebuggerRemotePacketRequestedAction = 4
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP                         _DebuggerRemotePacketRequestedAction = 5
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE                     _DebuggerRemotePacketRequestedAction = 6
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE    _DebuggerRemotePacketRequestedAction = 7
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE                  _DebuggerRemotePacketRequestedAction = 8
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS                _DebuggerRemotePacketRequestedAction = 9
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK                    _DebuggerRemotePacketRequestedAction = 10
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY                   _DebuggerRemotePacketRequestedAction = 11
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS               _DebuggerRemotePacketRequestedAction = 12
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD                _DebuggerRemotePacketRequestedAction = 13
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT                        _DebuggerRemotePacketRequestedAction = 14
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER                 _DebuggerRemotePacketRequestedAction = 15
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY                      _DebuggerRemotePacketRequestedAction = 16
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT                    _DebuggerRemotePacketRequestedAction = 17
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT               _DebuggerRemotePacketRequestedAction = 18
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT            _DebuggerRemotePacketRequestedAction = 19
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS                    _DebuggerRemotePacketRequestedAction = 20
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY                       _DebuggerRemotePacketRequestedAction = 21
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY                       _DebuggerRemotePacketRequestedAction = 22
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP                                _DebuggerRemotePacketRequestedAction = 23
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS        _DebuggerRemotePacketRequestedAction = 24
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD                     _DebuggerRemotePacketRequestedAction = 25
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA             _DebuggerRemotePacketRequestedAction = 26
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE                  _DebuggerRemotePacketRequestedAction = 27
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SET_SHORT_CIRCUITING_STATE        _DebuggerRemotePacketRequestedAction = 28
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_INJECT_PAGE_FAULT                 _DebuggerRemotePacketRequestedAction = 29
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION                                     _DebuggerRemotePacketRequestedAction = 30
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED                              _DebuggerRemotePacketRequestedAction = 31
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM                    _DebuggerRemotePacketRequestedAction = 32
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION       _DebuggerRemotePacketRequestedAction = 33
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE              _DebuggerRemotePacketRequestedAction = 34
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS           _DebuggerRemotePacketRequestedAction = 35
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD            _DebuggerRemotePacketRequestedAction = 36
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT             _DebuggerRemotePacketRequestedAction = 37
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS                    _DebuggerRemotePacketRequestedAction = 38
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH                      _DebuggerRemotePacketRequestedAction = 39
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK                  _DebuggerRemotePacketRequestedAction = 40
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY                    _DebuggerRemotePacketRequestedAction = 41
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT          _DebuggerRemotePacketRequestedAction = 42
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT     _DebuggerRemotePacketRequestedAction = 43
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT     _DebuggerRemotePacketRequestedAction = 44
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_EVENT     _DebuggerRemotePacketRequestedAction = 45
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS          _DebuggerRemotePacketRequestedAction = 46
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY             _DebuggerRemotePacketRequestedAction = 47
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY             _DebuggerRemotePacketRequestedAction = 48
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP                         _DebuggerRemotePacketRequestedAction = 49
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE     _DebuggerRemotePacketRequestedAction = 50
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS _DebuggerRemotePacketRequestedAction = 51
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO                   _DebuggerRemotePacketRequestedAction = 52
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED               _DebuggerRemotePacketRequestedAction = 53
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY                  _DebuggerRemotePacketRequestedAction = 54
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE                        _DebuggerRemotePacketRequestedAction = 55
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA            _DebuggerRemotePacketRequestedAction = 56
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BRINGING_PAGES_IN          _DebuggerRemotePacketRequestedAction = 57
)

// @brief enum for different packet types in HyperDbg packets
//
// @warning used in hwdbg
type _DebuggerRemotePacketType int32

const (
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT  _DebuggerRemotePacketType = 1
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE _DebuggerRemotePacketType = 2
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER                      _DebuggerRemotePacketType = 3
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL       _DebuggerRemotePacketType = 4
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL       _DebuggerRemotePacketType = 5
)

// @brief Different levels of paging
type _PagingLevel int32

const (
	PAGING_LEVEL_PAGE_TABLE                   _PagingLevel = 0
	PAGING_LEVEL_PAGE_DIRECTORY               _PagingLevel = 1
	PAGING_LEVEL_PAGE_DIRECTORY_POINTER_TABLE _PagingLevel = 2
	PAGING_LEVEL_PAGE_MAP_LEVEL4              _PagingLevel = 3
)

// @brief Inum of intentions for buffers (buffer tag)
type _PoolAllocationIntention int32

const (
	TRACKING_HOOKED_PAGES                  _PoolAllocationIntention = 0
	EXEC_TRAMPOLINE                        _PoolAllocationIntention = 1
	SPLIT_2MB_PAGING_TO_4KB_PAGE           _PoolAllocationIntention = 2
	DETOUR_HOOK_DETAILS                    _PoolAllocationIntention = 3
	BREAKPOINT_DEFINITION_STRUCTURE        _PoolAllocationIntention = 4
	PROCESS_THREAD_HOLDER                  _PoolAllocationIntention = 5
	INSTANT_REGULAR_EVENT_BUFFER           _PoolAllocationIntention = 6
	INSTANT_BIG_EVENT_BUFFER               _PoolAllocationIntention = 7
	INSTANT_REGULAR_EVENT_ACTION_BUFFER    _PoolAllocationIntention = 8
	INSTANT_BIG_EVENT_ACTION_BUFFER        _PoolAllocationIntention = 9
	INSTANT_REGULAR_SAFE_BUFFER_FOR_EVENTS _PoolAllocationIntention = 10
	INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS     _PoolAllocationIntention = 11
)

// ///////////////////////////////////////////////
type _DebugRegisterType int32

const (
	BREAK_ON_INSTRUCTION_FETCH              _DebugRegisterType = 0
	BREAK_ON_WRITE_ONLY                     _DebugRegisterType = 1
	BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED _DebugRegisterType = 2
	BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH   _DebugRegisterType = 3
)

// ///////////////////////////////////////////////
type _VmxExecutionMode int32

const (
	VMX_EXECUTION_MODE_NON_ROOT _VmxExecutionMode = 0
	VMX_EXECUTION_MODE_ROOT     _VmxExecutionMode = 1
)

// @brief Type of calling the event
type _VmmCallbackEventCallingStageType int32

const (
	VMM_CALLBACK_CALLING_STAGE_INVALID_EVENT_EMULATION _VmmCallbackEventCallingStageType = 0
	VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION     _VmmCallbackEventCallingStageType = 1
	VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION    _VmmCallbackEventCallingStageType = 2
	VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION     _VmmCallbackEventCallingStageType = 3
)

// @brief enum to query different process and thread interception mechanisms
type _DebuggerThreadProcessTracing int32

const (
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_THREAD_CHANGE  _DebuggerThreadProcessTracing = 0
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_PROCESS_CHANGE _DebuggerThreadProcessTracing = 1
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION   _DebuggerThreadProcessTracing = 2
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_WAITING_FOR_MOV_CR3_VM_EXITS  _DebuggerThreadProcessTracing = 3
)

// @brief Type of transferring buffer between user-to-kernel
type _NotifyType int32

const (
	IRP_BASED   _NotifyType = 0
	EVENT_BASED _NotifyType = 1
)

// @brief different type of memory addresses
type _DebuggerHookMemoryType int32

const (
	DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS  _DebuggerHookMemoryType = 0
	DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS _DebuggerHookMemoryType = 1
)

// @brief Exceptions enum
type _ExceptionVectors int32

const (
	EXCEPTION_VECTOR_DIVIDE_ERROR                      _ExceptionVectors = 0
	EXCEPTION_VECTOR_DEBUG_BREAKPOINT                  _ExceptionVectors = 1
	EXCEPTION_VECTOR_NMI                               _ExceptionVectors = 2
	EXCEPTION_VECTOR_BREAKPOINT                        _ExceptionVectors = 3
	EXCEPTION_VECTOR_OVERFLOW                          _ExceptionVectors = 4
	EXCEPTION_VECTOR_BOUND_RANGE_EXCEEDED              _ExceptionVectors = 5
	EXCEPTION_VECTOR_UNDEFINED_OPCODE                  _ExceptionVectors = 6
	EXCEPTION_VECTOR_NO_MATH_COPROCESSOR               _ExceptionVectors = 7
	EXCEPTION_VECTOR_DOUBLE_FAULT                      _ExceptionVectors = 8
	EXCEPTION_VECTOR_RESERVED0                         _ExceptionVectors = 9
	EXCEPTION_VECTOR_INVALID_TASK_SEGMENT_SELECTOR     _ExceptionVectors = 10
	EXCEPTION_VECTOR_SEGMENT_NOT_PRESENT               _ExceptionVectors = 11
	EXCEPTION_VECTOR_STACK_SEGMENT_FAULT               _ExceptionVectors = 12
	EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT          _ExceptionVectors = 13
	EXCEPTION_VECTOR_PAGE_FAULT                        _ExceptionVectors = 14
	EXCEPTION_VECTOR_RESERVED1                         _ExceptionVectors = 15
	EXCEPTION_VECTOR_MATH_FAULT                        _ExceptionVectors = 16
	EXCEPTION_VECTOR_ALIGNMENT_CHECK                   _ExceptionVectors = 17
	EXCEPTION_VECTOR_MACHINE_CHECK                     _ExceptionVectors = 18
	EXCEPTION_VECTOR_SIMD_FLOATING_POINT_NUMERIC_ERROR _ExceptionVectors = 19
	EXCEPTION_VECTOR_VIRTUAL_EXCEPTION                 _ExceptionVectors = 20
	EXCEPTION_VECTOR_RESERVED2                         _ExceptionVectors = 21
	EXCEPTION_VECTOR_RESERVED3                         _ExceptionVectors = 22
	EXCEPTION_VECTOR_RESERVED4                         _ExceptionVectors = 23
	EXCEPTION_VECTOR_RESERVED5                         _ExceptionVectors = 24
	EXCEPTION_VECTOR_RESERVED6                         _ExceptionVectors = 25
	EXCEPTION_VECTOR_RESERVED7                         _ExceptionVectors = 26
	EXCEPTION_VECTOR_RESERVED8                         _ExceptionVectors = 27
	EXCEPTION_VECTOR_RESERVED9                         _ExceptionVectors = 28
	EXCEPTION_VECTOR_RESERVED10                        _ExceptionVectors = 29
	EXCEPTION_VECTOR_RESERVED11                        _ExceptionVectors = 30
	EXCEPTION_VECTOR_RESERVED12                        _ExceptionVectors = 31
	APC_INTERRUPT                                      _ExceptionVectors = 31
	DPC_INTERRUPT                                      _ExceptionVectors = 47
	CLOCK_INTERRUPT                                    _ExceptionVectors = 209
	IPI_INTERRUPT                                      _ExceptionVectors = 225
	PMI_INTERRUPT                                      _ExceptionVectors = 254
)

// @brief The status of triggering events
type _VmmCallbackTriggeringEventStatusType int32

const (
	VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_NO_INITIALIZED _VmmCallbackTriggeringEventStatusType = 0
	VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL                _VmmCallbackTriggeringEventStatusType = 0
	VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT   _VmmCallbackTriggeringEventStatusType = 1
	VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED      _VmmCallbackTriggeringEventStatusType = 2
	VMM_CALLBACK_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE        _VmmCallbackTriggeringEventStatusType = 3
)

// @brief enum to show type of all HyperDbg events
type _VmmEventTypeEnum int32

const (
	HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE _VmmEventTypeEnum = 0
	HIDDEN_HOOK_READ_AND_WRITE             _VmmEventTypeEnum = 1
	HIDDEN_HOOK_READ_AND_EXECUTE           _VmmEventTypeEnum = 2
	HIDDEN_HOOK_WRITE_AND_EXECUTE          _VmmEventTypeEnum = 3
	HIDDEN_HOOK_READ                       _VmmEventTypeEnum = 4
	HIDDEN_HOOK_WRITE                      _VmmEventTypeEnum = 5
	HIDDEN_HOOK_EXECUTE                    _VmmEventTypeEnum = 6
	HIDDEN_HOOK_EXEC_DETOURS               _VmmEventTypeEnum = 7
	HIDDEN_HOOK_EXEC_CC                    _VmmEventTypeEnum = 8
	SYSCALL_HOOK_EFER_SYSCALL              _VmmEventTypeEnum = 9
	SYSCALL_HOOK_EFER_SYSRET               _VmmEventTypeEnum = 10
	CPUID_INSTRUCTION_EXECUTION            _VmmEventTypeEnum = 11
	RDMSR_INSTRUCTION_EXECUTION            _VmmEventTypeEnum = 12
	WRMSR_INSTRUCTION_EXECUTION            _VmmEventTypeEnum = 13
	IN_INSTRUCTION_EXECUTION               _VmmEventTypeEnum = 14
	OUT_INSTRUCTION_EXECUTION              _VmmEventTypeEnum = 15
	EXCEPTION_OCCURRED                     _VmmEventTypeEnum = 16
	EXTERNAL_INTERRUPT_OCCURRED            _VmmEventTypeEnum = 17
	DEBUG_REGISTERS_ACCESSED               _VmmEventTypeEnum = 18
	TSC_INSTRUCTION_EXECUTION              _VmmEventTypeEnum = 19
	PMC_INSTRUCTION_EXECUTION              _VmmEventTypeEnum = 20
	VMCALL_INSTRUCTION_EXECUTION           _VmmEventTypeEnum = 21
	CONTROL_REGISTER_MODIFIED              _VmmEventTypeEnum = 22
	CONTROL_REGISTER_READ                  _VmmEventTypeEnum = 23
	CONTROL_REGISTER_3_MODIFIED            _VmmEventTypeEnum = 24
	TRAP_EXECUTION_MODE_CHANGED            _VmmEventTypeEnum = 25
	TRAP_EXECUTION_INSTRUCTION_TRACE       _VmmEventTypeEnum = 26
)

// @brief Type of Actions
type _DebuggerEventActionTypeEnum int32

const (
	BREAK_TO_DEBUGGER _DebuggerEventActionTypeEnum = 0
	RUN_SCRIPT        _DebuggerEventActionTypeEnum = 1
	RUN_CUSTOM_CODE   _DebuggerEventActionTypeEnum = 2
)

// @brief Type of handling !syscall or !sysret
type _DebuggerEventSyscallSysretType int32

const (
	DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY _DebuggerEventSyscallSysretType = 0
	DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD      _DebuggerEventSyscallSysretType = 1
)

// @brief Type of mode change traps
type _DebuggerEventModeType int32
