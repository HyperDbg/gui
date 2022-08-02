package Headers



const (
	HIDDEN_HOOK_READ_AND_WRITE   = 1  //col:3
	HIDDEN_HOOK_READ             = 2  //col:4
	HIDDEN_HOOK_WRITE            = 3  //col:5
	HIDDEN_HOOK_EXEC_DETOURS     = 4  //col:6
	HIDDEN_HOOK_EXEC_CC          = 5  //col:7
	SYSCALL_HOOK_EFER_SYSCALL    = 6  //col:8
	SYSCALL_HOOK_EFER_SYSRET     = 7  //col:9
	CPUID_INSTRUCTION_EXECUTION  = 8  //col:10
	RDMSR_INSTRUCTION_EXECUTION  = 9  //col:11
	WRMSR_INSTRUCTION_EXECUTION  = 10 //col:12
	IN_INSTRUCTION_EXECUTION     = 11 //col:13
	OUT_INSTRUCTION_EXECUTION    = 12 //col:14
	EXCEPTION_OCCURRED           = 13 //col:15
	EXTERNAL_INTERRUPT_OCCURRED  = 14 //col:16
	DEBUG_REGISTERS_ACCESSED     = 15 //col:17
	TSC_INSTRUCTION_EXECUTION    = 16 //col:18
	PMC_INSTRUCTION_EXECUTION    = 17 //col:19
	VMCALL_INSTRUCTION_EXECUTION = 18 //col:20
	CONTROL_REGISTER_MODIFIED    = 19 //col:21
)

const (
	BREAK_TO_DEBUGGER = 1 //col:25
	RUN_SCRIPT        = 2 //col:26
	RUN_CUSTOM_CODE   = 3 //col:27
)

const (
	DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY = 0 //col:31
	DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD      = 1 //col:32
)

const (
	DEBUGGER_MODIFY_EVENTS_QUERY_STATE = 1 //col:36
	DEBUGGER_MODIFY_EVENTS_ENABLE      = 2 //col:37
	DEBUGGER_MODIFY_EVENTS_DISABLE     = 3 //col:38
	DEBUGGER_MODIFY_EVENTS_CLEAR       = 4 //col:39
)

type DEBUGGER_MODIFY_EVENTS struct {
	Tag                      uint64                      //col:3
	KernelStatus             uint64                      //col:4
	DebuggerModifyEventsType DEBUGGER_MODIFY_EVENTS_TYPE //col:5
	TypeOfAction             byte                        //col:6
	IsEnabled                bool                        //col:7
}


