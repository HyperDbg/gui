package HPRDBGCTRL

const (
	MaxSerialPacketSize                                                                      = 10 * NORMAL_PAGE_SIZE
	PAGE_SIZE                                                                                = 4096
	DEBUGGEE_BP_APPLY_TO_ALL_CORES                                                           = 0xffffffff
	DEBUGGEE_BP_APPLY_TO_ALL_THREADS                                                         = 0xffffffff
	DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE                                                     = 0xffffffff
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED                                               = 0xc000001b
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS                                  = 0xc0000019
	DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK                                              = 0xc0000027
	DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE                                         = 0xc000004d
	DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER                               = 0xc0000020
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER                                      = 0xc000000b
	DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS                            = 0xc0000025
	DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED                                                      = 0xc000001f
	DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES                                   = 0xc0000007
	DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND                    = 0xc0000046
	DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND                           = 0xc0000043
	DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER         = 0xc0000047
	DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER = 0xc000004b
	DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND                     = 0xc0000049
	DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID                                              = 0xc0000008
	DEBUGGER_ERROR_INVALID_ADDRESS                                                           = 0xc0000005
	DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS                                                  = 0xc0000052
	DEBUGGER_ERROR_INVALID_REGISTER_NUMBER                                                   = 0xc0000017
	DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN                                            = 0xc0000030
	DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE                                       = 0xc0000018
	DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID                                                 = 0xc000003f
	DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION                                      = 0xc000000f
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE                                       = 0xc0000012
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT                                          = 0xc0000016
	DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE    = 0xc0000040
	DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED                             = 0xc0000011
	DEBUGGER_ERROR_TAG_NOT_EXISTS                                                            = 0xc0000000
	DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED                                     = 0xc000004e
	DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL                                                = 0xc000003d
	DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER                                  = 0xc000004c
	DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS                              = 0xc000002a
	DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT                                                    = 0xc0000004
	DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS                                 = 0xc000002e
	DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS                                      = 0xc0000037
	DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS                                                = 0xc000002f
	DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS                                       = 0xc0000031
	DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS                                                    = 0xc000002c
	DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID                       = 0xc0000035
	DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED                                               = 0xc000003b
	DEBUGGER_EVENT_ALL_IO_PORTS                                                              = 0xffffffff
	DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES                                                    = 0xffffffff
	DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS                                                = 0xffffffff
	DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG                                                  = 0xffffffffffffffff
	DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES                                           = 0xffffffff
	DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET                                              = 0x200
	DEFAULT_PORT                                                                             = "50000"
	DbgPrintLimitation                                                                       = 512
	DebuggerOutputSourceMaximumRemoteSourceForSingleEvent                                    = 0x5
	DebuggerScriptEngineMemcpyMovingBufferSize                                               = 64
	FALSE                                                                                    = 0
	MAXIMUM_BIG_INSTANT_EVENTS                                                               = 0
	MAXIMUM_CALL_INSTR_SIZE                                                                  = 7
	MAXIMUM_GUID_AND_AGE_SIZE                                                                = 60
	MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS                                         = 5
	MAXIMUM_REGULAR_INSTANT_EVENTS                                                           = 20
	MAX_FUNCTION_NAME_LENGTH                                                                 = 32
	MAX_STACK_BUFFER_COUNT                                                                   = 128
	MAX_VAR_COUNT                                                                            = 512
	MaximumPacketsCapacity                                                                   = 1000
	MaximumSearchResults                                                                     = 0x1000
	NULL64_ZERO                                                                              = 0
	OPERATION_LOG_NON_IMMEDIATE_MESSAGE                                                      = 4
	POOLTAG                                                                                  = 0x48444247
	SERIAL_END_OF_BUFFER_CHARS_COUNT                                                         = 0x4
	SERIAL_END_OF_BUFFER_CHAR_2                                                              = 0x80
	SERIAL_END_OF_BUFFER_CHAR_4                                                              = 0xFF
	TCP_END_OF_BUFFER_CHARS_COUNT                                                            = 0x4
	TCP_END_OF_BUFFER_CHAR_2                                                                 = 0x20
	TCP_END_OF_BUFFER_CHAR_4                                                                 = 0x44
	TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER                                                 = 0x00000200
	VERSION_MAJOR                                                                            = 1
	VERSION_PATCH                                                                            = 0
	X86_FLAGS_IOPL_SHIFT_2ND_BIT                                                             = (13)
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
	X86_FLAGS_RESERVED_BITS                                                                  = 0xffc38028
)
