package sdk

import (
	"strings"

	
)

// Code generated by GeneratedFile types - DO NOT EDIT.

type DebuggerErrorType uint8

const (
	TagNotExistsType DebuggerErrorType = iota
	InvalidActionTypeType
	ActionBufferSizeIsZeroType
	EventTypeIsInvalidType
	UnableToCreateEventType
	InvalidAddressType
	InvalidCoreIdType
	ExceptionIndexExceedFirst32EntriesType
	InterruptIndexIsNotValidType
	UnableToHideOrUnhideDebuggerType
	DebuggerAlreadyUhideType
	EditMemoryStatusInvalidParameterType
	EditMemoryStatusInvalidAddressBasedOnCurrentProcessType
	EditMemoryStatusInvalidAddressBasedOnOtherProcessType
	ModifyEventsInvalidTagType
	ModifyEventsInvalidTypeOfActionType
	SteppingInvalidParameterType
	SteppingsEitherThreadNotFoundOrDisabledType
	PreparingDebuggeeInvalidBaudrateType
	PreparingDebuggeeInvalidSerialPortType
	PreparingDebuggeeInvalidCoreInRemoteDebuggeType
	PreparingDebuggeeUnableToSwitchToNewProcessType
	PreparingDebuggeeToRunScriptType
	InvalidRegisterNumberType
	MaximumBreakpointWithoutContinueType
	BreakpointAlreadyExistsOnTheAddressType
	BreakpointIdNotFoundType
	BreakpointAlreadyDisabledType
	BreakpointAlreadyEnabledType
	MemoryTypeInvalidType
	InvalidProcessIdType
	EventIsNotAppliedType
	DetailsOrSwitchProcessInvalidParameterType
	DetailsOrSwitchThreadInvalidParameterType
	MaximumBreakpointForASinglePageIsHitType
	PreAllocatedBufferIsEmptyType
	EptCouldNotSplitTheLargePageTo4KbPagesType
	EptFailedToGetPml1EntryOfTargetAddressType
	EptMultipleHooksInASinglePageType
	CouldNotBuildTheEptHookType
	CouldNotFindAllocationTypeType
	InvalidTestQueryIndexType
	UnableToAttachToTargetUserModeProcessType
	UnableToRemoveHooksEntrypointNotReachedType
	UnableToRemoveHooksType
	FunctionsForInitializingPebAddressesAreNotInitializedType
	UnableToDetect32BitOr64BitProcessType
	UnableToKillTheProcessType
	InvalidThreadDebuggingTokenType
	UnableToPauseTheProcessThreadsType
	UnableToAttachToAnAlreadyAttachedProcessType
	TheUserDebuggerNotAttachedToTheProcessType
	UnableToDetachAsThereArePausedThreadsType
	UnableToSwitchProcessIdOrThreadIdIsInvalidType
	UnableToSwitchThereIsNoThreadOnTheProcessType
	UnableToGetModulesOfTheProcessType
	UnableToGetCallstackType
	UnableToQueryCountOfProcessesOrThreadsType
	UsingShortCircuitingEventWithPostEventModeIsForbiddednType
	UnknownTestQueryReceivedType
	ReadingMemoryInvalidParameterType
	TheTrapFlagListIsFullType
	UnableToKillTheProcessDoesNotExistsType
	ModeExecutionIsInvalidType
	ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeType
	InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsType
	InstantEventRegularPreallocatedBufferNotFoundType
	InstantEventBigPreallocatedBufferNotFoundType
	UnableToCreateActionCannotAllocateBufferType
	InstantEventActionRegularPreallocatedBufferNotFoundType
	InstantEventActionBigPreallocatedBufferNotFoundType
	InstantEventPreallocatedBufferIsNotEnoughForActionBufferType
	InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackType
	InstantEventRegularRequestedSafeBufferNotFoundType
	InstantEventBigRequestedSafeBufferNotFoundType
	InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferType
	UnableToAllocateRequestedSafeBufferType
	CouldNotFindPreactivationTypeType
	TheModeExecTrapIsNotInitializedType
	TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullType
	NotAllCoresAreLockedForApplyingInstantEventType
	TargetSwitchingCoreIsNotLockedType
	InvalidPhysicalAddressType
)

func (t DebuggerErrorType) Valid() bool {
	return t >= TagNotExistsType && t <= InvalidPhysicalAddressType
}

func DebuggerErrorTypeBy[T stream.Integer](v T) DebuggerErrorType {
	return DebuggerErrorType(v)
}

func (t DebuggerErrorType) AssertBy(name string) DebuggerErrorType {
	name = strings.TrimSuffix(name, "Type")
	for _, n := range t.EnumTypes() {
		if strings.ToLower(name) == strings.ToLower(n.String()) {
			return n
		}
	}
	panic("InvalidType")
}

func (t DebuggerErrorType) String() string {
	switch t {
	case TagNotExistsType:
		return "TagNotExists"
	case InvalidActionTypeType:
		return "InvalidActionType"
	case ActionBufferSizeIsZeroType:
		return "ActionBufferSizeIsZero"
	case EventTypeIsInvalidType:
		return "EventTypeIsInvalid"
	case UnableToCreateEventType:
		return "UnableToCreateEvent"
	case InvalidAddressType:
		return "InvalidAddress"
	case InvalidCoreIdType:
		return "InvalidCoreId"
	case ExceptionIndexExceedFirst32EntriesType:
		return "ExceptionIndexExceedFirst32Entries"
	case InterruptIndexIsNotValidType:
		return "InterruptIndexIsNotValid"
	case UnableToHideOrUnhideDebuggerType:
		return "UnableToHideOrUnhideDebugger"
	case DebuggerAlreadyUhideType:
		return "DebuggerAlreadyUhide"
	case EditMemoryStatusInvalidParameterType:
		return "EditMemoryStatusInvalidParameter"
	case EditMemoryStatusInvalidAddressBasedOnCurrentProcessType:
		return "EditMemoryStatusInvalidAddressBasedOnCurrentProcess"
	case EditMemoryStatusInvalidAddressBasedOnOtherProcessType:
		return "EditMemoryStatusInvalidAddressBasedOnOtherProcess"
	case ModifyEventsInvalidTagType:
		return "ModifyEventsInvalidTag"
	case ModifyEventsInvalidTypeOfActionType:
		return "ModifyEventsInvalidTypeOfAction"
	case SteppingInvalidParameterType:
		return "SteppingInvalidParameter"
	case SteppingsEitherThreadNotFoundOrDisabledType:
		return "SteppingsEitherThreadNotFoundOrDisabled"
	case PreparingDebuggeeInvalidBaudrateType:
		return "PreparingDebuggeeInvalidBaudrate"
	case PreparingDebuggeeInvalidSerialPortType:
		return "PreparingDebuggeeInvalidSerialPort"
	case PreparingDebuggeeInvalidCoreInRemoteDebuggeType:
		return "PreparingDebuggeeInvalidCoreInRemoteDebugge"
	case PreparingDebuggeeUnableToSwitchToNewProcessType:
		return "PreparingDebuggeeUnableToSwitchToNewProcess"
	case PreparingDebuggeeToRunScriptType:
		return "PreparingDebuggeeToRunScript"
	case InvalidRegisterNumberType:
		return "InvalidRegisterNumber"
	case MaximumBreakpointWithoutContinueType:
		return "MaximumBreakpointWithoutContinue"
	case BreakpointAlreadyExistsOnTheAddressType:
		return "BreakpointAlreadyExistsOnTheAddress"
	case BreakpointIdNotFoundType:
		return "BreakpointIdNotFound"
	case BreakpointAlreadyDisabledType:
		return "BreakpointAlreadyDisabled"
	case BreakpointAlreadyEnabledType:
		return "BreakpointAlreadyEnabled"
	case MemoryTypeInvalidType:
		return "MemoryTypeInvalid"
	case InvalidProcessIdType:
		return "InvalidProcessId"
	case EventIsNotAppliedType:
		return "EventIsNotApplied"
	case DetailsOrSwitchProcessInvalidParameterType:
		return "DetailsOrSwitchProcessInvalidParameter"
	case DetailsOrSwitchThreadInvalidParameterType:
		return "DetailsOrSwitchThreadInvalidParameter"
	case MaximumBreakpointForASinglePageIsHitType:
		return "MaximumBreakpointForASinglePageIsHit"
	case PreAllocatedBufferIsEmptyType:
		return "PreAllocatedBufferIsEmpty"
	case EptCouldNotSplitTheLargePageTo4KbPagesType:
		return "EptCouldNotSplitTheLargePageTo4KbPages"
	case EptFailedToGetPml1EntryOfTargetAddressType:
		return "EptFailedToGetPml1EntryOfTargetAddress"
	case EptMultipleHooksInASinglePageType:
		return "EptMultipleHooksInASinglePage"
	case CouldNotBuildTheEptHookType:
		return "CouldNotBuildTheEptHook"
	case CouldNotFindAllocationTypeType:
		return "CouldNotFindAllocationType"
	case InvalidTestQueryIndexType:
		return "InvalidTestQueryIndex"
	case UnableToAttachToTargetUserModeProcessType:
		return "UnableToAttachToTargetUserModeProcess"
	case UnableToRemoveHooksEntrypointNotReachedType:
		return "UnableToRemoveHooksEntrypointNotReached"
	case UnableToRemoveHooksType:
		return "UnableToRemoveHooks"
	case FunctionsForInitializingPebAddressesAreNotInitializedType:
		return "FunctionsForInitializingPebAddressesAreNotInitialized"
	case UnableToDetect32BitOr64BitProcessType:
		return "UnableToDetect32BitOr64BitProcess"
	case UnableToKillTheProcessType:
		return "UnableToKillTheProcess"
	case InvalidThreadDebuggingTokenType:
		return "InvalidThreadDebuggingToken"
	case UnableToPauseTheProcessThreadsType:
		return "UnableToPauseTheProcessThreads"
	case UnableToAttachToAnAlreadyAttachedProcessType:
		return "UnableToAttachToAnAlreadyAttachedProcess"
	case TheUserDebuggerNotAttachedToTheProcessType:
		return "TheUserDebuggerNotAttachedToTheProcess"
	case UnableToDetachAsThereArePausedThreadsType:
		return "UnableToDetachAsThereArePausedThreads"
	case UnableToSwitchProcessIdOrThreadIdIsInvalidType:
		return "UnableToSwitchProcessIdOrThreadIdIsInvalid"
	case UnableToSwitchThereIsNoThreadOnTheProcessType:
		return "UnableToSwitchThereIsNoThreadOnTheProcess"
	case UnableToGetModulesOfTheProcessType:
		return "UnableToGetModulesOfTheProcess"
	case UnableToGetCallstackType:
		return "UnableToGetCallstack"
	case UnableToQueryCountOfProcessesOrThreadsType:
		return "UnableToQueryCountOfProcessesOrThreads"
	case UsingShortCircuitingEventWithPostEventModeIsForbiddednType:
		return "UsingShortCircuitingEventWithPostEventModeIsForbiddedn"
	case UnknownTestQueryReceivedType:
		return "UnknownTestQueryReceived"
	case ReadingMemoryInvalidParameterType:
		return "ReadingMemoryInvalidParameter"
	case TheTrapFlagListIsFullType:
		return "TheTrapFlagListIsFull"
	case UnableToKillTheProcessDoesNotExistsType:
		return "UnableToKillTheProcessDoesNotExists"
	case ModeExecutionIsInvalidType:
		return "ModeExecutionIsInvalid"
	case ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeType:
		return "ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode"
	case InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsType:
		return "InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals"
	case InstantEventRegularPreallocatedBufferNotFoundType:
		return "InstantEventRegularPreallocatedBufferNotFound"
	case InstantEventBigPreallocatedBufferNotFoundType:
		return "InstantEventBigPreallocatedBufferNotFound"
	case UnableToCreateActionCannotAllocateBufferType:
		return "UnableToCreateActionCannotAllocateBuffer"
	case InstantEventActionRegularPreallocatedBufferNotFoundType:
		return "InstantEventActionRegularPreallocatedBufferNotFound"
	case InstantEventActionBigPreallocatedBufferNotFoundType:
		return "InstantEventActionBigPreallocatedBufferNotFound"
	case InstantEventPreallocatedBufferIsNotEnoughForActionBufferType:
		return "InstantEventPreallocatedBufferIsNotEnoughForActionBuffer"
	case InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackType:
		return "InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack"
	case InstantEventRegularRequestedSafeBufferNotFoundType:
		return "InstantEventRegularRequestedSafeBufferNotFound"
	case InstantEventBigRequestedSafeBufferNotFoundType:
		return "InstantEventBigRequestedSafeBufferNotFound"
	case InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferType:
		return "InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer"
	case UnableToAllocateRequestedSafeBufferType:
		return "UnableToAllocateRequestedSafeBuffer"
	case CouldNotFindPreactivationTypeType:
		return "CouldNotFindPreactivationType"
	case TheModeExecTrapIsNotInitializedType:
		return "TheModeExecTrapIsNotInitialized"
	case TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullType:
		return "TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull"
	case NotAllCoresAreLockedForApplyingInstantEventType:
		return "NotAllCoresAreLockedForApplyingInstantEvent"
	case TargetSwitchingCoreIsNotLockedType:
		return "TargetSwitchingCoreIsNotLocked"
	case InvalidPhysicalAddressType:
		return "InvalidPhysicalAddress"
	default:
		panic("InvalidType")
	}
}

func (t DebuggerErrorType) Tooltip() string {
	switch t {
	case TagNotExistsType:
		return "_TAG_NOT_EXISTS"
	case InvalidActionTypeType:
		return "_INVALID_ACTION_TYPE"
	case ActionBufferSizeIsZeroType:
		return "_ACTION_BUFFER_SIZE_IS_ZERO"
	case EventTypeIsInvalidType:
		return "_EVENT_TYPE_IS_INVALID"
	case UnableToCreateEventType:
		return "_UNABLE_TO_CREATE_EVENT"
	case InvalidAddressType:
		return "_INVALID_ADDRESS"
	case InvalidCoreIdType:
		return "_INVALID_CORE_ID"
	case ExceptionIndexExceedFirst32EntriesType:
		return "_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES"
	case InterruptIndexIsNotValidType:
		return "_INTERRUPT_INDEX_IS_NOT_VALID"
	case UnableToHideOrUnhideDebuggerType:
		return "_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER"
	case DebuggerAlreadyUhideType:
		return "_DEBUGGER_ALREADY_UHIDE"
	case EditMemoryStatusInvalidParameterType:
		return "_EDIT_MEMORY_STATUS_INVALID_PARAMETER"
	case EditMemoryStatusInvalidAddressBasedOnCurrentProcessType:
		return "_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS"
	case EditMemoryStatusInvalidAddressBasedOnOtherProcessType:
		return "_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS"
	case ModifyEventsInvalidTagType:
		return "_MODIFY_EVENTS_INVALID_TAG"
	case ModifyEventsInvalidTypeOfActionType:
		return "_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION"
	case SteppingInvalidParameterType:
		return "_STEPPING_INVALID_PARAMETER"
	case SteppingsEitherThreadNotFoundOrDisabledType:
		return "_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED"
	case PreparingDebuggeeInvalidBaudrateType:
		return "_PREPARING_DEBUGGEE_INVALID_BAUDRATE"
	case PreparingDebuggeeInvalidSerialPortType:
		return "_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT"
	case PreparingDebuggeeInvalidCoreInRemoteDebuggeType:
		return "_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE"
	case PreparingDebuggeeUnableToSwitchToNewProcessType:
		return "_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS"
	case PreparingDebuggeeToRunScriptType:
		return "_PREPARING_DEBUGGEE_TO_RUN_SCRIPT"
	case InvalidRegisterNumberType:
		return "_INVALID_REGISTER_NUMBER"
	case MaximumBreakpointWithoutContinueType:
		return "_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE"
	case BreakpointAlreadyExistsOnTheAddressType:
		return "_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS"
	case BreakpointIdNotFoundType:
		return "_BREAKPOINT_ID_NOT_FOUND"
	case BreakpointAlreadyDisabledType:
		return "_BREAKPOINT_ALREADY_DISABLED"
	case BreakpointAlreadyEnabledType:
		return "_BREAKPOINT_ALREADY_ENABLED"
	case MemoryTypeInvalidType:
		return "_MEMORY_TYPE_INVALID"
	case InvalidProcessIdType:
		return "_INVALID_PROCESS_ID"
	case EventIsNotAppliedType:
		return "_EVENT_IS_NOT_APPLIED"
	case DetailsOrSwitchProcessInvalidParameterType:
		return "_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER"
	case DetailsOrSwitchThreadInvalidParameterType:
		return "_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER"
	case MaximumBreakpointForASinglePageIsHitType:
		return "_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT"
	case PreAllocatedBufferIsEmptyType:
		return "_PRE_ALLOCATED_BUFFER_IS_EMPTY"
	case EptCouldNotSplitTheLargePageTo4KbPagesType:
		return "_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES"
	case EptFailedToGetPml1EntryOfTargetAddressType:
		return "_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS"
	case EptMultipleHooksInASinglePageType:
		return "_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE"
	case CouldNotBuildTheEptHookType:
		return "_COULD_NOT_BUILD_THE_EPT_HOOK"
	case CouldNotFindAllocationTypeType:
		return "_COULD_NOT_FIND_ALLOCATION_TYPE"
	case InvalidTestQueryIndexType:
		return "_INVALID_TEST_QUERY_INDEX"
	case UnableToAttachToTargetUserModeProcessType:
		return "_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS"
	case UnableToRemoveHooksEntrypointNotReachedType:
		return "_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED"
	case UnableToRemoveHooksType:
		return "_UNABLE_TO_REMOVE_HOOKS"
	case FunctionsForInitializingPebAddressesAreNotInitializedType:
		return "_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED"
	case UnableToDetect32BitOr64BitProcessType:
		return "_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS"
	case UnableToKillTheProcessType:
		return "_UNABLE_TO_KILL_THE_PROCESS"
	case InvalidThreadDebuggingTokenType:
		return "_INVALID_THREAD_DEBUGGING_TOKEN"
	case UnableToPauseTheProcessThreadsType:
		return "_UNABLE_TO_PAUSE_THE_PROCESS_THREADS"
	case UnableToAttachToAnAlreadyAttachedProcessType:
		return "_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS"
	case TheUserDebuggerNotAttachedToTheProcessType:
		return "_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS"
	case UnableToDetachAsThereArePausedThreadsType:
		return "_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS"
	case UnableToSwitchProcessIdOrThreadIdIsInvalidType:
		return "_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID"
	case UnableToSwitchThereIsNoThreadOnTheProcessType:
		return "_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS"
	case UnableToGetModulesOfTheProcessType:
		return "_UNABLE_TO_GET_MODULES_OF_THE_PROCESS"
	case UnableToGetCallstackType:
		return "_UNABLE_TO_GET_CALLSTACK"
	case UnableToQueryCountOfProcessesOrThreadsType:
		return "_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS"
	case UsingShortCircuitingEventWithPostEventModeIsForbiddednType:
		return "_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN"
	case UnknownTestQueryReceivedType:
		return "_UNKNOWN_TEST_QUERY_RECEIVED"
	case ReadingMemoryInvalidParameterType:
		return "_READING_MEMORY_INVALID_PARAMETER"
	case TheTrapFlagListIsFullType:
		return "_THE_TRAP_FLAG_LIST_IS_FULL"
	case UnableToKillTheProcessDoesNotExistsType:
		return "_UNABLE_TO_KILL_THE_PROCESS_DOES_NOT_EXISTS"
	case ModeExecutionIsInvalidType:
		return "_MODE_EXECUTION_IS_INVALID"
	case ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeType:
		return "_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE"
	case InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsType:
		return "_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS"
	case InstantEventRegularPreallocatedBufferNotFoundType:
		return "_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND"
	case InstantEventBigPreallocatedBufferNotFoundType:
		return "_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND"
	case UnableToCreateActionCannotAllocateBufferType:
		return "_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER"
	case InstantEventActionRegularPreallocatedBufferNotFoundType:
		return "_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND"
	case InstantEventActionBigPreallocatedBufferNotFoundType:
		return "_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND"
	case InstantEventPreallocatedBufferIsNotEnoughForActionBufferType:
		return "_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER"
	case InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackType:
		return "_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK"
	case InstantEventRegularRequestedSafeBufferNotFoundType:
		return "_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND"
	case InstantEventBigRequestedSafeBufferNotFoundType:
		return "_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND"
	case InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferType:
		return "_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER"
	case UnableToAllocateRequestedSafeBufferType:
		return "_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER"
	case CouldNotFindPreactivationTypeType:
		return "_COULD_NOT_FIND_PREACTIVATION_TYPE"
	case TheModeExecTrapIsNotInitializedType:
		return "_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED"
	case TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullType:
		return "_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL"
	case NotAllCoresAreLockedForApplyingInstantEventType:
		return "_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT"
	case TargetSwitchingCoreIsNotLockedType:
		return "_TARGET_SWITCHING_CORE_IS_NOT_LOCKED"
	case InvalidPhysicalAddressType:
		return "_INVALID_PHYSICAL_ADDRESS"
	default:
		panic("InvalidType")
	}
}

func (t DebuggerErrorType) Names() []string {
	return []string{
		"TagNotExists",
		"InvalidActionType",
		"ActionBufferSizeIsZero",
		"EventTypeIsInvalid",
		"UnableToCreateEvent",
		"InvalidAddress",
		"InvalidCoreId",
		"ExceptionIndexExceedFirst32Entries",
		"InterruptIndexIsNotValid",
		"UnableToHideOrUnhideDebugger",
		"DebuggerAlreadyUhide",
		"EditMemoryStatusInvalidParameter",
		"EditMemoryStatusInvalidAddressBasedOnCurrentProcess",
		"EditMemoryStatusInvalidAddressBasedOnOtherProcess",
		"ModifyEventsInvalidTag",
		"ModifyEventsInvalidTypeOfAction",
		"SteppingInvalidParameter",
		"SteppingsEitherThreadNotFoundOrDisabled",
		"PreparingDebuggeeInvalidBaudrate",
		"PreparingDebuggeeInvalidSerialPort",
		"PreparingDebuggeeInvalidCoreInRemoteDebugge",
		"PreparingDebuggeeUnableToSwitchToNewProcess",
		"PreparingDebuggeeToRunScript",
		"InvalidRegisterNumber",
		"MaximumBreakpointWithoutContinue",
		"BreakpointAlreadyExistsOnTheAddress",
		"BreakpointIdNotFound",
		"BreakpointAlreadyDisabled",
		"BreakpointAlreadyEnabled",
		"MemoryTypeInvalid",
		"InvalidProcessId",
		"EventIsNotApplied",
		"DetailsOrSwitchProcessInvalidParameter",
		"DetailsOrSwitchThreadInvalidParameter",
		"MaximumBreakpointForASinglePageIsHit",
		"PreAllocatedBufferIsEmpty",
		"EptCouldNotSplitTheLargePageTo4KbPages",
		"EptFailedToGetPml1EntryOfTargetAddress",
		"EptMultipleHooksInASinglePage",
		"CouldNotBuildTheEptHook",
		"CouldNotFindAllocationType",
		"InvalidTestQueryIndex",
		"UnableToAttachToTargetUserModeProcess",
		"UnableToRemoveHooksEntrypointNotReached",
		"UnableToRemoveHooks",
		"FunctionsForInitializingPebAddressesAreNotInitialized",
		"UnableToDetect32BitOr64BitProcess",
		"UnableToKillTheProcess",
		"InvalidThreadDebuggingToken",
		"UnableToPauseTheProcessThreads",
		"UnableToAttachToAnAlreadyAttachedProcess",
		"TheUserDebuggerNotAttachedToTheProcess",
		"UnableToDetachAsThereArePausedThreads",
		"UnableToSwitchProcessIdOrThreadIdIsInvalid",
		"UnableToSwitchThereIsNoThreadOnTheProcess",
		"UnableToGetModulesOfTheProcess",
		"UnableToGetCallstack",
		"UnableToQueryCountOfProcessesOrThreads",
		"UsingShortCircuitingEventWithPostEventModeIsForbiddedn",
		"UnknownTestQueryReceived",
		"ReadingMemoryInvalidParameter",
		"TheTrapFlagListIsFull",
		"UnableToKillTheProcessDoesNotExists",
		"ModeExecutionIsInvalid",
		"ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode",
		"InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals",
		"InstantEventRegularPreallocatedBufferNotFound",
		"InstantEventBigPreallocatedBufferNotFound",
		"UnableToCreateActionCannotAllocateBuffer",
		"InstantEventActionRegularPreallocatedBufferNotFound",
		"InstantEventActionBigPreallocatedBufferNotFound",
		"InstantEventPreallocatedBufferIsNotEnoughForActionBuffer",
		"InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack",
		"InstantEventRegularRequestedSafeBufferNotFound",
		"InstantEventBigRequestedSafeBufferNotFound",
		"InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer",
		"UnableToAllocateRequestedSafeBuffer",
		"CouldNotFindPreactivationType",
		"TheModeExecTrapIsNotInitialized",
		"TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull",
		"NotAllCoresAreLockedForApplyingInstantEvent",
		"TargetSwitchingCoreIsNotLocked",
		"InvalidPhysicalAddress",
	}
}

func (t DebuggerErrorType) EnumTypes() []DebuggerErrorType {
	return []DebuggerErrorType{
		TagNotExistsType,
		InvalidActionTypeType,
		ActionBufferSizeIsZeroType,
		EventTypeIsInvalidType,
		UnableToCreateEventType,
		InvalidAddressType,
		InvalidCoreIdType,
		ExceptionIndexExceedFirst32EntriesType,
		InterruptIndexIsNotValidType,
		UnableToHideOrUnhideDebuggerType,
		DebuggerAlreadyUhideType,
		EditMemoryStatusInvalidParameterType,
		EditMemoryStatusInvalidAddressBasedOnCurrentProcessType,
		EditMemoryStatusInvalidAddressBasedOnOtherProcessType,
		ModifyEventsInvalidTagType,
		ModifyEventsInvalidTypeOfActionType,
		SteppingInvalidParameterType,
		SteppingsEitherThreadNotFoundOrDisabledType,
		PreparingDebuggeeInvalidBaudrateType,
		PreparingDebuggeeInvalidSerialPortType,
		PreparingDebuggeeInvalidCoreInRemoteDebuggeType,
		PreparingDebuggeeUnableToSwitchToNewProcessType,
		PreparingDebuggeeToRunScriptType,
		InvalidRegisterNumberType,
		MaximumBreakpointWithoutContinueType,
		BreakpointAlreadyExistsOnTheAddressType,
		BreakpointIdNotFoundType,
		BreakpointAlreadyDisabledType,
		BreakpointAlreadyEnabledType,
		MemoryTypeInvalidType,
		InvalidProcessIdType,
		EventIsNotAppliedType,
		DetailsOrSwitchProcessInvalidParameterType,
		DetailsOrSwitchThreadInvalidParameterType,
		MaximumBreakpointForASinglePageIsHitType,
		PreAllocatedBufferIsEmptyType,
		EptCouldNotSplitTheLargePageTo4KbPagesType,
		EptFailedToGetPml1EntryOfTargetAddressType,
		EptMultipleHooksInASinglePageType,
		CouldNotBuildTheEptHookType,
		CouldNotFindAllocationTypeType,
		InvalidTestQueryIndexType,
		UnableToAttachToTargetUserModeProcessType,
		UnableToRemoveHooksEntrypointNotReachedType,
		UnableToRemoveHooksType,
		FunctionsForInitializingPebAddressesAreNotInitializedType,
		UnableToDetect32BitOr64BitProcessType,
		UnableToKillTheProcessType,
		InvalidThreadDebuggingTokenType,
		UnableToPauseTheProcessThreadsType,
		UnableToAttachToAnAlreadyAttachedProcessType,
		TheUserDebuggerNotAttachedToTheProcessType,
		UnableToDetachAsThereArePausedThreadsType,
		UnableToSwitchProcessIdOrThreadIdIsInvalidType,
		UnableToSwitchThereIsNoThreadOnTheProcessType,
		UnableToGetModulesOfTheProcessType,
		UnableToGetCallstackType,
		UnableToQueryCountOfProcessesOrThreadsType,
		UsingShortCircuitingEventWithPostEventModeIsForbiddednType,
		UnknownTestQueryReceivedType,
		ReadingMemoryInvalidParameterType,
		TheTrapFlagListIsFullType,
		UnableToKillTheProcessDoesNotExistsType,
		ModeExecutionIsInvalidType,
		ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeType,
		InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsType,
		InstantEventRegularPreallocatedBufferNotFoundType,
		InstantEventBigPreallocatedBufferNotFoundType,
		UnableToCreateActionCannotAllocateBufferType,
		InstantEventActionRegularPreallocatedBufferNotFoundType,
		InstantEventActionBigPreallocatedBufferNotFoundType,
		InstantEventPreallocatedBufferIsNotEnoughForActionBufferType,
		InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackType,
		InstantEventRegularRequestedSafeBufferNotFoundType,
		InstantEventBigRequestedSafeBufferNotFoundType,
		InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferType,
		UnableToAllocateRequestedSafeBufferType,
		CouldNotFindPreactivationTypeType,
		TheModeExecTrapIsNotInitializedType,
		TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullType,
		NotAllCoresAreLockedForApplyingInstantEventType,
		TargetSwitchingCoreIsNotLockedType,
		InvalidPhysicalAddressType,
	}
}

func (t DebuggerErrorType) SvgFileName() string {
	return t.String() + ".svg"
}
