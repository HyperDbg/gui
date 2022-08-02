package header



const (
	EVENT_FORWARDING_NAMEDPIPE = 1 //col:3
	EVENT_FORWARDING_FILE      = 2 //col:4
	EVENT_FORWARDING_TCP       = 3 //col:5
)

const (
	EVENT_FORWARDING_STATE_NOT_OPENED = 1 //col:9
	EVENT_FORWARDING_STATE_OPENED     = 2 //col:10
	EVENT_FORWARDING_CLOSED           = 3 //col:11
)

const (
	DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED = 1 //col:15
	DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED = 2 //col:16
	DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED      = 3 //col:17
	DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED      = 4 //col:18
	DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR       = 5 //col:19
)

type DEBUGGER_EVENT_FORWARDING struct {
	Type              DEBUGGER_EVENT_FORWARDING_TYPE                     //col:3
	State             DEBUGGER_EVENT_FORWARDING_STATE                    //col:4
	Handle            HANDLE                                             //col:5
	Socket            SOCKET                                             //col:6
	OutputUniqueTag   uint64                                             //col:7
	ListEntry         *list.List                                         //col:8
	OutputSourcesList byte                                               //col:9
	Name              [MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME]int8 //col:10
}


