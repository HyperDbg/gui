package Headers

type DEBUGGEE_PAUSING_REASON byte

const (
	// For both kernel & user debugger
	DEBUGGEE_PAUSING_REASON_NOT_PAUSED DEBUGGEE_PAUSING_REASON = iota
	DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM
	DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED

	// Only for user-debugger
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED

	// Only used for hardware debugging
	DEBUGGEE_PAUSING_REASON_HARDWARE_BASED_DEBUGGEE_GENERAL_BREAK
)

// add for decode status by error codes
//func (pr DEBUGGEE_PAUSING_REASON) String() string { return fmt.Sprint(pr) }

type DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION byte

const (
	// Debugger to debuggee (user-mode execution)
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_bad DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = iota
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET

	// Debugger to debuggee (vmx-root mode execution)
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SET_SHORT_CIRCUITING_STATE
)

const ( //if merge the value is wrong
	// Debuggee to debugger
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = iota
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_EVENT
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA

	// hardware debuggee to debugger

	// hardware debugger to debuggee
)

const (
	name = iota
)

type (
	DEBUGGER_REMOTE_PA          byte
	DEBUGGER_REMOTE_PACKET_TYPE DEBUGGER_REMOTE_PA
)

const (
	DEBUGGER_REMOTE_PA_bad DEBUGGER_REMOTE_PA = iota
	// Debugger to debuggee (vmx-root)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT

	// Debugger to debuggee (user-mode)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE

	// Debuggee to debugger (user-mode and kernel-mode, vmx-root mode)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER

	// Debugger to debuggee (hardware)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL = 1

	// Debuggee to debugger (hardware)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL = 2 //go syntax not support nested unum
)

//
//func (pr DEBUGGER_REMOTE_PA) String() string { return fmt.Sprint(pr) }

type (
	DEBUGGER_REMOTE_PACKET struct {
		Checksum                   byte
		Indicator                  uint64
		TypeOfThePacket            DEBUGGER_REMOTE_PACKET_TYPE
		RequestedActionOfThePacket DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
	}
	PDEBUGGER_REMOTE_PACKET *DEBUGGER_REMOTE_PACKET
)