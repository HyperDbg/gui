package HPRDBGCTRL

import "fmt"

type ErrorCodes int

const (
	DEBUGGER_OPERATION_WAS_SUCCESSFUL                                              ErrorCodes = 0xFFFFFFFF
	DEBUGGER_ERROR_TAG_NOT_EXISTS                                                             = 0xc0000000
	DEBUGGER_ERROR_INVALID_ACTION_TYPE                                                        = 0xc0000001
	DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO                                                 = 0xc0000002
	DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID                                                      = 0xc0000003
	DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT                                                     = 0xc0000004
	DEBUGGER_ERROR_INVALID_ADDRESS                                                            = 0xc0000005
	DEBUGGER_ERROR_INVALID_CORE_ID                                                            = 0xc0000006
	DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES                                    = 0xc0000007
	DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID                                               = 0xc0000008
	DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER                                          = 0xc0000009
	DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE                                                     = 0xc000000a
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER                                       = 0xc000000b
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS                = 0xc000000c
	DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS                  = 0xc000000d
	DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG                                                  = 0xc000000e
	DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION                                       = 0xc000000f
	DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER                                                 = 0xc0000010
	DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED                              = 0xc0000011
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE                                        = 0xc0000012
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT                                     = 0xc0000013
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE                          = 0xc0000014
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS                         = 0xc0000015
	DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT                                           = 0xc0000016
	DEBUGGER_ERROR_INVALID_REGISTER_NUMBER                                                    = 0xc0000017
	DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE                                        = 0xc0000018
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS                                   = 0xc0000019
	DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND                                                    = 0xc000001a
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED                                                = 0xc000001b
	DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED                                                 = 0xc000001c
	DEBUGGER_ERROR_MEMORY_TYPE_INVALID                                                        = 0xc000001d
	DEBUGGER_ERROR_INVALID_PROCESS_ID                                                         = 0xc000001e
	DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED                                                       = 0xc000001f
	DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER                                = 0xc0000020
	DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER                                 = 0xc0000021
	DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT                                = 0xc0000022
	DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY                                              = 0xc0000023
	DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES                            = 0xc0000024
	DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS                             = 0xc0000025
	DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE                                        = 0xc0000026
	DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK                                               = 0xc0000027
	DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE                                             = 0xc0000028
	DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX                                                   = 0xc0000029
	DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS                               = 0xc000002a
	DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED                              = 0xc000002b
	DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS                                                     = 0xc000002c
	DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED               = 0xc000002d
	DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS                                  = 0xc000002e
	DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS                                                 = 0xc000002f
	DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN                                             = 0xc0000030
	DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS                                        = 0xc0000031
	DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS                            = 0xc0000032
	DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS                              = 0xc0000033
	DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS                               = 0xc0000034
	DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID                        = 0xc0000035
	DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS                         = 0xc0000036
	DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS                                       = 0xc0000037
	DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK                                                    = 0xc0000038
	DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS                              = 0xc0000039
	DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN            = 0xc000003a
)

func (e ErrorCodes) String() string {
	switch e {
	case DEBUGGER_OPERATION_WAS_SUCCESSFUL:
		return "DebuggerOperationWasSuccessful"
	case DEBUGGER_ERROR_TAG_NOT_EXISTS:
		return "DebuggerErrorTagNotExists"
	case DEBUGGER_ERROR_INVALID_ACTION_TYPE:
		return "DebuggerErrorInvalidActionType"
	case DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO:
		return "DebuggerErrorActionBufferSizeIsZero"
	case DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID:
		return "DebuggerErrorEventTypeIsInvalid"
	case DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT:
		return "DebuggerErrorUnableToCreateEvent"
	case DEBUGGER_ERROR_INVALID_ADDRESS:
		return "DebuggerErrorInvalidAddress"
	case DEBUGGER_ERROR_INVALID_CORE_ID:
		return "DebuggerErrorInvalidCoreId"
	case DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES:
		return "DebuggerErrorExceptionIndexExceedFirst32Entries"
	case DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID:
		return "DebuggerErrorInterruptIndexIsNotValid"
	case DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER:
		return "DebuggerErrorUnableToHideOrUnhideDebugger"
	case DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE:
		return "DebuggerErrorDebuggerAlreadyUhide"
	case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER:
		return "DebuggerErrorEditMemoryStatusInvalidParameter"
	case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcess"
	case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcess"
	case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG:
		return "DebuggerErrorModifyEventsInvalidTag"
	case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION:
		return "DebuggerErrorModifyEventsInvalidTypeOfAction"
	case DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER:
		return "DebuggerErrorSteppingInvalidParameter"
	case DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED:
		return "DebuggerErrorSteppingsEitherThreadNotFoundOrDisabled"
	case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE:
		return "DebuggerErrorPreparingDebuggeeInvalidBaudrate"
	case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT:
		return "DebuggerErrorPreparingDebuggeeInvalidSerialPort"
	case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE:
		return "DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebugge"
	case DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS:
		return "DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcess"
	case DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT:
		return "DebuggerErrorPreparingDebuggeeToRunScript"
	case DEBUGGER_ERROR_INVALID_REGISTER_NUMBER:
		return "DebuggerErrorInvalidRegisterNumber"
	case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE:
		return "DebuggerErrorMaximumBreakpointWithoutContinue"
	case DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS:
		return "DebuggerErrorBreakpointAlreadyExistsOnTheAddress"
	case DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND:
		return "DebuggerErrorBreakpointIdNotFound"
	case DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED:
		return "DebuggerErrorBreakpointAlreadyDisabled"
	case DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED:
		return "DebuggerErrorBreakpointAlreadyEnabled"
	case DEBUGGER_ERROR_MEMORY_TYPE_INVALID:
		return "DebuggerErrorMemoryTypeInvalid"
	case DEBUGGER_ERROR_INVALID_PROCESS_ID:
		return "DebuggerErrorInvalidProcessId"
	case DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED:
		return "DebuggerErrorEventIsNotApplied"
	case DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER:
		return "DebuggerErrorDetailsOrSwitchProcessInvalidParameter"
	case DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER:
		return "DebuggerErrorDetailsOrSwitchThreadInvalidParameter"
	case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT:
		return "DebuggerErrorMaximumBreakpointForASinglePageIsHit"
	case DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY:
		return "DebuggerErrorPreAllocatedBufferIsEmpty"
	case DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES:
		return "DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPages"
	case DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS:
		return "DebuggerErrorEptFailedToGetPml1EntryOfTargetAddress"
	case DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE:
		return "DebuggerErrorEptMultipleHooksInASinglePage"
	case DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK:
		return "DebuggerErrorCouldNotBuildTheEptHook"
	case DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE:
		return "DebuggerErrorCouldNotFindAllocationType"
	case DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX:
		return "DebuggerErrorInvalidTestQueryIndex"
	case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS:
		return "DebuggerErrorUnableToAttachToTargetUserModeProcess"
	case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED:
		return "DebuggerErrorUnableToRemoveHooksEntrypointNotReached"
	case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS:
		return "DebuggerErrorUnableToRemoveHooks"
	case DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED:
		return "DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitialized"
	case DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS:
		return "DebuggerErrorUnableToDetect32BitOr64BitProcess"
	case DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS:
		return "DebuggerErrorUnableToKillTheProcess"
	case DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN:
		return "DebuggerErrorInvalidThreadDebuggingToken"
	case DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS:
		return "DebuggerErrorUnableToPauseTheProcessThreads"
	case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS:
		return "DebuggerErrorUnableToAttachToAnAlreadyAttachedProcess"
	case DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS:
		return "DebuggerErrorTheUserDebuggerNotAttachedToTheProcess"
	case DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS:
		return "DebuggerErrorUnableToDetachAsThereArePausedThreads"
	case DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID:
		return "DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalid"
	case DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS:
		return "DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcess"
	case DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS:
		return "DebuggerErrorUnableToGetModulesOfTheProcess"
	case DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK:
		return "DebuggerErrorUnableToGetCallstack"
	case DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS:
		return "DebuggerErrorUnableToQueryCountOfProcessesOrThreads"
	case DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN:
		return "DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddedn"
	default:
		return ""
		return fmt.Sprint("unknown error code " + fmt.Sprintf("%d", e))
	}
}
