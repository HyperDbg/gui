package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\debugger.h.back

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

type _DEBUGGER_SYNCRONIZATION_EVENTS_STATE struct {
	EventHandle      uintptr //col:7
	IsOnWaitingState bool    //col:8
}

