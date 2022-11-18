package Headers

type DEBUGGER_EVENT_TYPE_ENUM byte

const (
	HIDDEN_HOOK_READ_AND_WRITE DEBUGGER_EVENT_TYPE_ENUM = iota //todo check this value is 0 in c?
	HIDDEN_HOOK_READ
	HIDDEN_HOOK_WRITE

	HIDDEN_HOOK_EXEC_DETOURS
	HIDDEN_HOOK_EXEC_CC

	SYSCALL_HOOK_EFER_SYSCALL
	SYSCALL_HOOK_EFER_SYSRET

	CPUID_INSTRUCTION_EXECUTION

	RDMSR_INSTRUCTION_EXECUTION
	WRMSR_INSTRUCTION_EXECUTION

	IN_INSTRUCTION_EXECUTION
	OUT_INSTRUCTION_EXECUTION

	EXCEPTION_OCCURRED
	EXTERNAL_INTERRUPT_OCCURRED

	DEBUG_REGISTERS_ACCESSED

	TSC_INSTRUCTION_EXECUTION
	PMC_INSTRUCTION_EXECUTION

	VMCALL_INSTRUCTION_EXECUTION

	CONTROL_REGISTER_MODIFIED
	CONTROL_REGISTER_READ
)

type DEBUGGER_EVENT_ACTION_TYPE_ENUM byte

const (
	BREAK_TO_DEBUGGER DEBUGGER_EVENT_ACTION_TYPE_ENUM = iota
	RUN_SCRIPT
	RUN_CUSTOM_CODE
)

type DEBUGGER_EVENT_CALLING_STAGE_TYPE byte

const (
	DEBUGGER_CALLING_STAGE_PRE_EVENT_EMULATION DEBUGGER_EVENT_CALLING_STAGE_TYPE = iota
	DEBUGGER_CALLING_STAGE_POST_EVENT_EMULATION
)

type DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE byte

const (
	DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE = iota
	DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD                                         = 1
)

//todo
//#define SIZEOF_DEBUGGER_MODIFY_EVENTS sizeof(DEBUGGER_MODIFY_EVENTS)

type DEBUGGER_MODIFY_EVENTS_TYPE byte

const (
	DEBUGGER_MODIFY_EVENTS_QUERY_STATE DEBUGGER_MODIFY_EVENTS_TYPE = iota
	DEBUGGER_MODIFY_EVENTS_ENABLE
	DEBUGGER_MODIFY_EVENTS_DISABLE
	DEBUGGER_MODIFY_EVENTS_CLEAR
)

type (
	DEBUGGER_MODIFY_EVENTS struct {
		Tag          uint64                      // Tag of the target event that we want to modify
		KernelStatus uint64                      // Kerenl put the status in this field
		TypeOfAction DEBUGGER_MODIFY_EVENTS_TYPE // Determines what's the action (enable | disable | clear)
		IsEnabled    bool                        // Determines what's the action (enable | disable | clear)

	}
	PDEBUGGER_MODIFY_EVENTS *DEBUGGER_MODIFY_EVENTS
)

type (
	DEBUGGER_SHORT_CIRCUITING_EVENT struct {
		KernelStatus      uint64 // Kerenl put the status in this field
		IsShortCircuiting bool   // Determines whether to perform short circuting (on | off)
	}

	PDEBUGGER_SHORT_CIRCUITING_EVENT *DEBUGGER_SHORT_CIRCUITING_EVENT
)
