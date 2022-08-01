package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\debugger.h.back

const (
	DefaultSpeedOfReadingKernelMessages                                                = 30   //col:1
	DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS                            = 0x40 //col:2
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING                 = 0x0  //col:3
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED             = 0x1  //col:4
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS             = 0x2  //col:5
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT               = 0x3  //col:6
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT            = 0x4  //col:7
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT             = 0x5  //col:8
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT               = 0x6  //col:9
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT               = 0x7  //col:10
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION = 0x8  //col:11
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_FLUSH_RESULT                        = 0x9  //col:12
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT                      = 0xa  //col:13
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT                 = 0xb  //col:14
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_MODIFY_AND_QUERY_EVENT              = 0xc  //col:15
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS                      = 0xd  //col:16
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_BP                                  = 0xe  //col:17
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_LIST_OR_MODIFY_BREAKPOINTS          = 0xf  //col:18
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY                         = 0x10 //col:19
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY                         = 0x11 //col:20
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD                       = 0x12 //col:21
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_TEST_QUERY                          = 0x13 //col:22
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT                    = 0x14 //col:23
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SEARCH_QUERY_RESULT                 = 0x15 //col:24
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_VA2PA_AND_PA2VA_RESULT              = 0x16 //col:25
	DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PTE_RESULT                          = 0x17 //col:26
	DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS                              = 0x40 //col:27
	DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING                   = 0x30 //col:28
)

const (
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SUCCESSFUL_NO_ERROR                          = 0 //col:3
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SCRIPT_SYNTAX_ERROR                          = 1 //col:4
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_NO_INPUT                                     = 2 //col:5
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_MAXIMUM_INPUT_REACHED                        = 3 //col:6
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND                        = 4 //col:7
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_SOURCE_ALREADY_CLOSED                 = 5 //col:8
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ALLOCATION_ERROR                             = 6 //col:9
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR                                 = 7 //col:10
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ATTEMPT_TO_BREAK_ON_VMI_MODE                 = 8 //col:11
	DEBUGGER_EVENT_PARSING_ERROR_CAUSE_IMMEDIATE_MESSAGING_IN_EVENT_FORWARDING_MODE = 9 //col:12
)

type DEBUGGER_SYNCRONIZATION_EVENTS_STATE struct {
	EventHandle      HANDLE //col:3
	IsOnWaitingState bool   //col:4
}

