package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\ErrorCodes.h.back

const (
	DEBUGGER_OPERATION_WAS_SUCCESSFULL                                          = 0xFFFFFFFF //col:1
	DEBUGGER_ERROR_TAG_NOT_EXISTS                                               = 0xc0000000 //col:2
	DEBUGGER_ERROR_INVALID_ACTION_TYPE                                          = 0xc0000001 //col:3
	DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO                                   = 0xc0000002 //col:4
	DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID                                        = 0xc0000003 //col:5
	DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT                                       = 0xc0000004 //col:6
	DEBUGGER_ERROR_INVALID_ADDRESS                                              = 0xc0000005 //col:7
	DEBUGGER_ERROR_INVALID_CORE_ID                                              = 0xc0000006 //col:8
	DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES                      = 0xc0000007 //col:9
	DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID                                 = 0xc0000008 //col:10
	DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER                            = 0xc0000009 //col:11
	DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE                                       = 0xc000000a //col:12
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER                         = 0xc000000b //col:13
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS  = 0xc000000c //col:14
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS    = 0xc000000d //col:16
	DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG                                    = 0xc000000e //col:18
	DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION                         = 0xc000000f //col:19
	DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER                                   = 0xc0000010 //col:20
	DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED                = 0xc0000011 //col:21
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE                          = 0xc0000012 //col:22
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT                       = 0xc0000013 //col:23
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE            = 0xc0000014 //col:24
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS           = 0xc0000015 //col:26
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT                             = 0xc0000016 //col:28
	DEBUGGER_ERROR_INVALID_REGISTER_NUMBER                                      = 0xc0000017 //col:29
	DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE                          = 0xc0000018 //col:30
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS                     = 0xc0000019 //col:31
	DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND                                      = 0xc000001a //col:32
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED                                  = 0xc000001b //col:33
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED                                   = 0xc000001c //col:34
	DEBUGGER_ERROR_MEMORY_TYPE_INVALID                                          = 0xc000001d //col:35
	DEBUGGER_ERROR_INVALID_PROCESS_ID                                           = 0xc000001e //col:36
	DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED                                         = 0xc000001f //col:37
	DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER                  = 0xc0000020 //col:38
	DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER                   = 0xc0000021 //col:39
	DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT                  = 0xc0000022 //col:40
	DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY                                = 0xc0000023 //col:41
	DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES              = 0xc0000024 //col:42
	DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS               = 0xc0000025 //col:43
	DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE                          = 0xc0000026 //col:44
	DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK                                 = 0xc0000027 //col:45
	DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE                               = 0xc0000028 //col:46
	DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX                                     = 0xc0000029 //col:47
	DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS                 = 0xc000002a //col:48
	DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED                = 0xc000002b //col:49
	DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS                                       = 0xc000002c //col:50
	DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED = 0xc000002d //col:51
	DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS                    = 0xc000002e //col:52
	DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS                                   = 0xc000002f //col:53
	DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN                               = 0xc0000030 //col:54
	DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS                          = 0xc0000031 //col:55
	DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS              = 0xc0000032 //col:56
	DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS                = 0xc0000033 //col:57
	DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS                 = 0xc0000034 //col:58
	DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID          = 0xc0000035 //col:59
	DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS           = 0xc0000036 //col:60
	DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS                         = 0xc0000037 //col:61
	DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK                                      = 0xc0000038 //col:62
	DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS                = 0xc0000039 //col:63
)
