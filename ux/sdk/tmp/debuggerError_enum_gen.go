package libhyperdbg

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Code generated by GeneratedFile enum - DO NOT EDIT.

type DebuggerErrorKind byte

const (
	DebuggerErrorTagNotExistsKind DebuggerErrorKind = iota
	DebuggerErrorInvalidActionTypeKind
	DebuggerErrorActionBufferSizeIsZeroKind
	DebuggerErrorEventTypeIsInvalidKind
	DebuggerErrorUnableToCreateEventKind
	DebuggerErrorInvalidAddressKind
	DebuggerErrorInvalidCoreIdKind
	DebuggerErrorExceptionIndexExceedFirst32EntriesKind
	DebuggerErrorInterruptIndexIsNotValidKind
	DebuggerErrorUnableToHideOrUnhideDebuggerKind
	DebuggerErrorDebuggerAlreadyUhideKind
	DebuggerErrorEditMemoryStatusInvalidParameterKind
	DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcessKind
	DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcessKind
	DebuggerErrorModifyEventsInvalidTagKind
	DebuggerErrorModifyEventsInvalidTypeOfActionKind
	DebuggerErrorSteppingInvalidParameterKind
	DebuggerErrorSteppingsEitherThreadNotFoundOrDisabledKind
	DebuggerErrorPreparingDebuggeeInvalidBaudrateKind
	DebuggerErrorPreparingDebuggeeInvalidSerialPortKind
	DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebuggeKind
	DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcessKind
	DebuggerErrorPreparingDebuggeeToRunScriptKind
	DebuggerErrorInvalidRegisterNumberKind
	DebuggerErrorMaximumBreakpointWithoutContinueKind
	DebuggerErrorBreakpointAlreadyExistsOnTheAddressKind
	DebuggerErrorBreakpointIdNotFoundKind
	DebuggerErrorBreakpointAlreadyDisabledKind
	DebuggerErrorBreakpointAlreadyEnabledKind
	DebuggerErrorMemoryTypeInvalidKind
	DebuggerErrorInvalidProcessIdKind
	DebuggerErrorEventIsNotAppliedKind
	DebuggerErrorDetailsOrSwitchProcessInvalidParameterKind
	DebuggerErrorDetailsOrSwitchThreadInvalidParameterKind
	DebuggerErrorMaximumBreakpointForASinglePageIsHitKind
	DebuggerErrorPreAllocatedBufferIsEmptyKind
	DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPagesKind
	DebuggerErrorEptFailedToGetPml1EntryOfTargetAddressKind
	DebuggerErrorEptMultipleHooksInASinglePageKind
	DebuggerErrorCouldNotBuildTheEptHookKind
	DebuggerErrorCouldNotFindAllocationTypeKind
	DebuggerErrorInvalidTestQueryIndexKind
	DebuggerErrorUnableToAttachToTargetUserModeProcessKind
	DebuggerErrorUnableToRemoveHooksEntrypointNotReachedKind
	DebuggerErrorUnableToRemoveHooksKind
	DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitializedKind
	DebuggerErrorUnableToDetect32BitOr64BitProcessKind
	DebuggerErrorUnableToKillTheProcessKind
	DebuggerErrorInvalidThreadDebuggingTokenKind
	DebuggerErrorUnableToPauseTheProcessThreadsKind
	DebuggerErrorUnableToAttachToAnAlreadyAttachedProcessKind
	DebuggerErrorTheUserDebuggerNotAttachedToTheProcessKind
	DebuggerErrorUnableToDetachAsThereArePausedThreadsKind
	DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalidKind
	DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcessKind
	DebuggerErrorUnableToGetModulesOfTheProcessKind
	DebuggerErrorUnableToGetCallstackKind
	DebuggerErrorUnableToQueryCountOfProcessesOrThreadsKind
	DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddednKind
	DebuggerErrorUnknownTestQueryReceivedKind
	DebuggerErrorReadingMemoryInvalidParameterKind
	DebuggerErrorTheTrapFlagListIsFullKind
	DebuggerErrorUnableToKillTheProcessDoesNotExistsKind
	DebuggerErrorModeExecutionIsInvalidKind
	DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeKind
	DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsKind
	DebuggerErrorInstantEventRegularPreallocatedBufferNotFoundKind
	DebuggerErrorInstantEventBigPreallocatedBufferNotFoundKind
	DebuggerErrorUnableToCreateActionCannotAllocateBufferKind
	DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFoundKind
	DebuggerErrorInstantEventActionBigPreallocatedBufferNotFoundKind
	DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBufferKind
	DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackKind
	DebuggerErrorInstantEventRegularRequestedSafeBufferNotFoundKind
	DebuggerErrorInstantEventBigRequestedSafeBufferNotFoundKind
	DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferKind
	DebuggerErrorUnableToAllocateRequestedSafeBufferKind
	DebuggerErrorCouldNotFindPreactivationTypeKind
	DebuggerErrorTheModeExecTrapIsNotInitializedKind
	DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullKind
	DebuggerErrorNotAllCoresAreLockedForApplyingInstantEventKind
	DebuggerErrorTargetSwitchingCoreIsNotLockedKind
	DebuggerErrorInvalidPhysicalAddressKind
	InvalidDebuggerErrorKind
)

func ConvertInteger2DebuggerErrorKind[T constraints.Integer](v T) DebuggerErrorKind {
	return DebuggerErrorKind(v)
}

func (k DebuggerErrorKind) AssertKind(kinds string) DebuggerErrorKind {
	for _, kind := range k.Kinds() {
		if strings.ToLower(kinds) == strings.ToLower(kind.String()) {
			return kind
		}
	}
	return InvalidDebuggerErrorKind
}

func (k DebuggerErrorKind) String() string {
	switch k {
	case DebuggerErrorTagNotExistsKind:
		return "DebuggerErrorTagNotExists"
	case DebuggerErrorInvalidActionTypeKind:
		return "DebuggerErrorInvalidActionType"
	case DebuggerErrorActionBufferSizeIsZeroKind:
		return "DebuggerErrorActionBufferSizeIsZero"
	case DebuggerErrorEventTypeIsInvalidKind:
		return "DebuggerErrorEventTypeIsInvalid"
	case DebuggerErrorUnableToCreateEventKind:
		return "DebuggerErrorUnableToCreateEvent"
	case DebuggerErrorInvalidAddressKind:
		return "DebuggerErrorInvalidAddress"
	case DebuggerErrorInvalidCoreIdKind:
		return "DebuggerErrorInvalidCoreId"
	case DebuggerErrorExceptionIndexExceedFirst32EntriesKind:
		return "DebuggerErrorExceptionIndexExceedFirst32Entries"
	case DebuggerErrorInterruptIndexIsNotValidKind:
		return "DebuggerErrorInterruptIndexIsNotValid"
	case DebuggerErrorUnableToHideOrUnhideDebuggerKind:
		return "DebuggerErrorUnableToHideOrUnhideDebugger"
	case DebuggerErrorDebuggerAlreadyUhideKind:
		return "DebuggerErrorDebuggerAlreadyUhide"
	case DebuggerErrorEditMemoryStatusInvalidParameterKind:
		return "DebuggerErrorEditMemoryStatusInvalidParameter"
	case DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcessKind:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcess"
	case DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcessKind:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcess"
	case DebuggerErrorModifyEventsInvalidTagKind:
		return "DebuggerErrorModifyEventsInvalidTag"
	case DebuggerErrorModifyEventsInvalidTypeOfActionKind:
		return "DebuggerErrorModifyEventsInvalidTypeOfAction"
	case DebuggerErrorSteppingInvalidParameterKind:
		return "DebuggerErrorSteppingInvalidParameter"
	case DebuggerErrorSteppingsEitherThreadNotFoundOrDisabledKind:
		return "DebuggerErrorSteppingsEitherThreadNotFoundOrDisabled"
	case DebuggerErrorPreparingDebuggeeInvalidBaudrateKind:
		return "DebuggerErrorPreparingDebuggeeInvalidBaudrate"
	case DebuggerErrorPreparingDebuggeeInvalidSerialPortKind:
		return "DebuggerErrorPreparingDebuggeeInvalidSerialPort"
	case DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebuggeKind:
		return "DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebugge"
	case DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcessKind:
		return "DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcess"
	case DebuggerErrorPreparingDebuggeeToRunScriptKind:
		return "DebuggerErrorPreparingDebuggeeToRunScript"
	case DebuggerErrorInvalidRegisterNumberKind:
		return "DebuggerErrorInvalidRegisterNumber"
	case DebuggerErrorMaximumBreakpointWithoutContinueKind:
		return "DebuggerErrorMaximumBreakpointWithoutContinue"
	case DebuggerErrorBreakpointAlreadyExistsOnTheAddressKind:
		return "DebuggerErrorBreakpointAlreadyExistsOnTheAddress"
	case DebuggerErrorBreakpointIdNotFoundKind:
		return "DebuggerErrorBreakpointIdNotFound"
	case DebuggerErrorBreakpointAlreadyDisabledKind:
		return "DebuggerErrorBreakpointAlreadyDisabled"
	case DebuggerErrorBreakpointAlreadyEnabledKind:
		return "DebuggerErrorBreakpointAlreadyEnabled"
	case DebuggerErrorMemoryTypeInvalidKind:
		return "DebuggerErrorMemoryTypeInvalid"
	case DebuggerErrorInvalidProcessIdKind:
		return "DebuggerErrorInvalidProcessId"
	case DebuggerErrorEventIsNotAppliedKind:
		return "DebuggerErrorEventIsNotApplied"
	case DebuggerErrorDetailsOrSwitchProcessInvalidParameterKind:
		return "DebuggerErrorDetailsOrSwitchProcessInvalidParameter"
	case DebuggerErrorDetailsOrSwitchThreadInvalidParameterKind:
		return "DebuggerErrorDetailsOrSwitchThreadInvalidParameter"
	case DebuggerErrorMaximumBreakpointForASinglePageIsHitKind:
		return "DebuggerErrorMaximumBreakpointForASinglePageIsHit"
	case DebuggerErrorPreAllocatedBufferIsEmptyKind:
		return "DebuggerErrorPreAllocatedBufferIsEmpty"
	case DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPagesKind:
		return "DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPages"
	case DebuggerErrorEptFailedToGetPml1EntryOfTargetAddressKind:
		return "DebuggerErrorEptFailedToGetPml1EntryOfTargetAddress"
	case DebuggerErrorEptMultipleHooksInASinglePageKind:
		return "DebuggerErrorEptMultipleHooksInASinglePage"
	case DebuggerErrorCouldNotBuildTheEptHookKind:
		return "DebuggerErrorCouldNotBuildTheEptHook"
	case DebuggerErrorCouldNotFindAllocationTypeKind:
		return "DebuggerErrorCouldNotFindAllocationType"
	case DebuggerErrorInvalidTestQueryIndexKind:
		return "DebuggerErrorInvalidTestQueryIndex"
	case DebuggerErrorUnableToAttachToTargetUserModeProcessKind:
		return "DebuggerErrorUnableToAttachToTargetUserModeProcess"
	case DebuggerErrorUnableToRemoveHooksEntrypointNotReachedKind:
		return "DebuggerErrorUnableToRemoveHooksEntrypointNotReached"
	case DebuggerErrorUnableToRemoveHooksKind:
		return "DebuggerErrorUnableToRemoveHooks"
	case DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitializedKind:
		return "DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitialized"
	case DebuggerErrorUnableToDetect32BitOr64BitProcessKind:
		return "DebuggerErrorUnableToDetect32BitOr64BitProcess"
	case DebuggerErrorUnableToKillTheProcessKind:
		return "DebuggerErrorUnableToKillTheProcess"
	case DebuggerErrorInvalidThreadDebuggingTokenKind:
		return "DebuggerErrorInvalidThreadDebuggingToken"
	case DebuggerErrorUnableToPauseTheProcessThreadsKind:
		return "DebuggerErrorUnableToPauseTheProcessThreads"
	case DebuggerErrorUnableToAttachToAnAlreadyAttachedProcessKind:
		return "DebuggerErrorUnableToAttachToAnAlreadyAttachedProcess"
	case DebuggerErrorTheUserDebuggerNotAttachedToTheProcessKind:
		return "DebuggerErrorTheUserDebuggerNotAttachedToTheProcess"
	case DebuggerErrorUnableToDetachAsThereArePausedThreadsKind:
		return "DebuggerErrorUnableToDetachAsThereArePausedThreads"
	case DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalidKind:
		return "DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalid"
	case DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcessKind:
		return "DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcess"
	case DebuggerErrorUnableToGetModulesOfTheProcessKind:
		return "DebuggerErrorUnableToGetModulesOfTheProcess"
	case DebuggerErrorUnableToGetCallstackKind:
		return "DebuggerErrorUnableToGetCallstack"
	case DebuggerErrorUnableToQueryCountOfProcessesOrThreadsKind:
		return "DebuggerErrorUnableToQueryCountOfProcessesOrThreads"
	case DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddednKind:
		return "DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddedn"
	case DebuggerErrorUnknownTestQueryReceivedKind:
		return "DebuggerErrorUnknownTestQueryReceived"
	case DebuggerErrorReadingMemoryInvalidParameterKind:
		return "DebuggerErrorReadingMemoryInvalidParameter"
	case DebuggerErrorTheTrapFlagListIsFullKind:
		return "DebuggerErrorTheTrapFlagListIsFull"
	case DebuggerErrorUnableToKillTheProcessDoesNotExistsKind:
		return "DebuggerErrorUnableToKillTheProcessDoesNotExists"
	case DebuggerErrorModeExecutionIsInvalidKind:
		return "DebuggerErrorModeExecutionIsInvalid"
	case DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeKind:
		return "DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals"
	case DebuggerErrorInstantEventRegularPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventRegularPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventBigPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventBigPreallocatedBufferNotFound"
	case DebuggerErrorUnableToCreateActionCannotAllocateBufferKind:
		return "DebuggerErrorUnableToCreateActionCannotAllocateBuffer"
	case DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventActionBigPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventActionBigPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBufferKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBuffer"
	case DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackKind:
		return "DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack"
	case DebuggerErrorInstantEventRegularRequestedSafeBufferNotFoundKind:
		return "DebuggerErrorInstantEventRegularRequestedSafeBufferNotFound"
	case DebuggerErrorInstantEventBigRequestedSafeBufferNotFoundKind:
		return "DebuggerErrorInstantEventBigRequestedSafeBufferNotFound"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer"
	case DebuggerErrorUnableToAllocateRequestedSafeBufferKind:
		return "DebuggerErrorUnableToAllocateRequestedSafeBuffer"
	case DebuggerErrorCouldNotFindPreactivationTypeKind:
		return "DebuggerErrorCouldNotFindPreactivationType"
	case DebuggerErrorTheModeExecTrapIsNotInitializedKind:
		return "DebuggerErrorTheModeExecTrapIsNotInitialized"
	case DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullKind:
		return "DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull"
	case DebuggerErrorNotAllCoresAreLockedForApplyingInstantEventKind:
		return "DebuggerErrorNotAllCoresAreLockedForApplyingInstantEvent"
	case DebuggerErrorTargetSwitchingCoreIsNotLockedKind:
		return "DebuggerErrorTargetSwitchingCoreIsNotLocked"
	case DebuggerErrorInvalidPhysicalAddressKind:
		return "DebuggerErrorInvalidPhysicalAddress"
	default:
		return "InvalidDebuggerErrorKind"
	}
}

func (k DebuggerErrorKind) Keys() []string {
	return []string{
		"DebuggerErrorTagNotExists",
		"DebuggerErrorInvalidActionType",
		"DebuggerErrorActionBufferSizeIsZero",
		"DebuggerErrorEventTypeIsInvalid",
		"DebuggerErrorUnableToCreateEvent",
		"DebuggerErrorInvalidAddress",
		"DebuggerErrorInvalidCoreId",
		"DebuggerErrorExceptionIndexExceedFirst32Entries",
		"DebuggerErrorInterruptIndexIsNotValid",
		"DebuggerErrorUnableToHideOrUnhideDebugger",
		"DebuggerErrorDebuggerAlreadyUhide",
		"DebuggerErrorEditMemoryStatusInvalidParameter",
		"DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcess",
		"DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcess",
		"DebuggerErrorModifyEventsInvalidTag",
		"DebuggerErrorModifyEventsInvalidTypeOfAction",
		"DebuggerErrorSteppingInvalidParameter",
		"DebuggerErrorSteppingsEitherThreadNotFoundOrDisabled",
		"DebuggerErrorPreparingDebuggeeInvalidBaudrate",
		"DebuggerErrorPreparingDebuggeeInvalidSerialPort",
		"DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebugge",
		"DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcess",
		"DebuggerErrorPreparingDebuggeeToRunScript",
		"DebuggerErrorInvalidRegisterNumber",
		"DebuggerErrorMaximumBreakpointWithoutContinue",
		"DebuggerErrorBreakpointAlreadyExistsOnTheAddress",
		"DebuggerErrorBreakpointIdNotFound",
		"DebuggerErrorBreakpointAlreadyDisabled",
		"DebuggerErrorBreakpointAlreadyEnabled",
		"DebuggerErrorMemoryTypeInvalid",
		"DebuggerErrorInvalidProcessId",
		"DebuggerErrorEventIsNotApplied",
		"DebuggerErrorDetailsOrSwitchProcessInvalidParameter",
		"DebuggerErrorDetailsOrSwitchThreadInvalidParameter",
		"DebuggerErrorMaximumBreakpointForASinglePageIsHit",
		"DebuggerErrorPreAllocatedBufferIsEmpty",
		"DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPages",
		"DebuggerErrorEptFailedToGetPml1EntryOfTargetAddress",
		"DebuggerErrorEptMultipleHooksInASinglePage",
		"DebuggerErrorCouldNotBuildTheEptHook",
		"DebuggerErrorCouldNotFindAllocationType",
		"DebuggerErrorInvalidTestQueryIndex",
		"DebuggerErrorUnableToAttachToTargetUserModeProcess",
		"DebuggerErrorUnableToRemoveHooksEntrypointNotReached",
		"DebuggerErrorUnableToRemoveHooks",
		"DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitialized",
		"DebuggerErrorUnableToDetect32BitOr64BitProcess",
		"DebuggerErrorUnableToKillTheProcess",
		"DebuggerErrorInvalidThreadDebuggingToken",
		"DebuggerErrorUnableToPauseTheProcessThreads",
		"DebuggerErrorUnableToAttachToAnAlreadyAttachedProcess",
		"DebuggerErrorTheUserDebuggerNotAttachedToTheProcess",
		"DebuggerErrorUnableToDetachAsThereArePausedThreads",
		"DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalid",
		"DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcess",
		"DebuggerErrorUnableToGetModulesOfTheProcess",
		"DebuggerErrorUnableToGetCallstack",
		"DebuggerErrorUnableToQueryCountOfProcessesOrThreads",
		"DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddedn",
		"DebuggerErrorUnknownTestQueryReceived",
		"DebuggerErrorReadingMemoryInvalidParameter",
		"DebuggerErrorTheTrapFlagListIsFull",
		"DebuggerErrorUnableToKillTheProcessDoesNotExists",
		"DebuggerErrorModeExecutionIsInvalid",
		"DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode",
		"DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals",
		"DebuggerErrorInstantEventRegularPreallocatedBufferNotFound",
		"DebuggerErrorInstantEventBigPreallocatedBufferNotFound",
		"DebuggerErrorUnableToCreateActionCannotAllocateBuffer",
		"DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFound",
		"DebuggerErrorInstantEventActionBigPreallocatedBufferNotFound",
		"DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBuffer",
		"DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack",
		"DebuggerErrorInstantEventRegularRequestedSafeBufferNotFound",
		"DebuggerErrorInstantEventBigRequestedSafeBufferNotFound",
		"DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer",
		"DebuggerErrorUnableToAllocateRequestedSafeBuffer",
		"DebuggerErrorCouldNotFindPreactivationType",
		"DebuggerErrorTheModeExecTrapIsNotInitialized",
		"DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull",
		"DebuggerErrorNotAllCoresAreLockedForApplyingInstantEvent",
		"DebuggerErrorTargetSwitchingCoreIsNotLocked",
		"DebuggerErrorInvalidPhysicalAddress",
		"InvalidDebuggerErrorKind",
	}
}

func (k DebuggerErrorKind) Kinds() []DebuggerErrorKind {
	return []DebuggerErrorKind{
		DebuggerErrorTagNotExistsKind,
		DebuggerErrorInvalidActionTypeKind,
		DebuggerErrorActionBufferSizeIsZeroKind,
		DebuggerErrorEventTypeIsInvalidKind,
		DebuggerErrorUnableToCreateEventKind,
		DebuggerErrorInvalidAddressKind,
		DebuggerErrorInvalidCoreIdKind,
		DebuggerErrorExceptionIndexExceedFirst32EntriesKind,
		DebuggerErrorInterruptIndexIsNotValidKind,
		DebuggerErrorUnableToHideOrUnhideDebuggerKind,
		DebuggerErrorDebuggerAlreadyUhideKind,
		DebuggerErrorEditMemoryStatusInvalidParameterKind,
		DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcessKind,
		DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcessKind,
		DebuggerErrorModifyEventsInvalidTagKind,
		DebuggerErrorModifyEventsInvalidTypeOfActionKind,
		DebuggerErrorSteppingInvalidParameterKind,
		DebuggerErrorSteppingsEitherThreadNotFoundOrDisabledKind,
		DebuggerErrorPreparingDebuggeeInvalidBaudrateKind,
		DebuggerErrorPreparingDebuggeeInvalidSerialPortKind,
		DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebuggeKind,
		DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcessKind,
		DebuggerErrorPreparingDebuggeeToRunScriptKind,
		DebuggerErrorInvalidRegisterNumberKind,
		DebuggerErrorMaximumBreakpointWithoutContinueKind,
		DebuggerErrorBreakpointAlreadyExistsOnTheAddressKind,
		DebuggerErrorBreakpointIdNotFoundKind,
		DebuggerErrorBreakpointAlreadyDisabledKind,
		DebuggerErrorBreakpointAlreadyEnabledKind,
		DebuggerErrorMemoryTypeInvalidKind,
		DebuggerErrorInvalidProcessIdKind,
		DebuggerErrorEventIsNotAppliedKind,
		DebuggerErrorDetailsOrSwitchProcessInvalidParameterKind,
		DebuggerErrorDetailsOrSwitchThreadInvalidParameterKind,
		DebuggerErrorMaximumBreakpointForASinglePageIsHitKind,
		DebuggerErrorPreAllocatedBufferIsEmptyKind,
		DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPagesKind,
		DebuggerErrorEptFailedToGetPml1EntryOfTargetAddressKind,
		DebuggerErrorEptMultipleHooksInASinglePageKind,
		DebuggerErrorCouldNotBuildTheEptHookKind,
		DebuggerErrorCouldNotFindAllocationTypeKind,
		DebuggerErrorInvalidTestQueryIndexKind,
		DebuggerErrorUnableToAttachToTargetUserModeProcessKind,
		DebuggerErrorUnableToRemoveHooksEntrypointNotReachedKind,
		DebuggerErrorUnableToRemoveHooksKind,
		DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitializedKind,
		DebuggerErrorUnableToDetect32BitOr64BitProcessKind,
		DebuggerErrorUnableToKillTheProcessKind,
		DebuggerErrorInvalidThreadDebuggingTokenKind,
		DebuggerErrorUnableToPauseTheProcessThreadsKind,
		DebuggerErrorUnableToAttachToAnAlreadyAttachedProcessKind,
		DebuggerErrorTheUserDebuggerNotAttachedToTheProcessKind,
		DebuggerErrorUnableToDetachAsThereArePausedThreadsKind,
		DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalidKind,
		DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcessKind,
		DebuggerErrorUnableToGetModulesOfTheProcessKind,
		DebuggerErrorUnableToGetCallstackKind,
		DebuggerErrorUnableToQueryCountOfProcessesOrThreadsKind,
		DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddednKind,
		DebuggerErrorUnknownTestQueryReceivedKind,
		DebuggerErrorReadingMemoryInvalidParameterKind,
		DebuggerErrorTheTrapFlagListIsFullKind,
		DebuggerErrorUnableToKillTheProcessDoesNotExistsKind,
		DebuggerErrorModeExecutionIsInvalidKind,
		DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeKind,
		DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsKind,
		DebuggerErrorInstantEventRegularPreallocatedBufferNotFoundKind,
		DebuggerErrorInstantEventBigPreallocatedBufferNotFoundKind,
		DebuggerErrorUnableToCreateActionCannotAllocateBufferKind,
		DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFoundKind,
		DebuggerErrorInstantEventActionBigPreallocatedBufferNotFoundKind,
		DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBufferKind,
		DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackKind,
		DebuggerErrorInstantEventRegularRequestedSafeBufferNotFoundKind,
		DebuggerErrorInstantEventBigRequestedSafeBufferNotFoundKind,
		DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferKind,
		DebuggerErrorUnableToAllocateRequestedSafeBufferKind,
		DebuggerErrorCouldNotFindPreactivationTypeKind,
		DebuggerErrorTheModeExecTrapIsNotInitializedKind,
		DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullKind,
		DebuggerErrorNotAllCoresAreLockedForApplyingInstantEventKind,
		DebuggerErrorTargetSwitchingCoreIsNotLockedKind,
		DebuggerErrorInvalidPhysicalAddressKind,
		InvalidDebuggerErrorKind,
	}
}

func (k DebuggerErrorKind) SvgFileName() string {
	switch k {
	case DebuggerErrorTagNotExistsKind:
		return "DebuggerErrorTagNotExists"
	case DebuggerErrorInvalidActionTypeKind:
		return "DebuggerErrorInvalidActionType"
	case DebuggerErrorActionBufferSizeIsZeroKind:
		return "DebuggerErrorActionBufferSizeIsZero"
	case DebuggerErrorEventTypeIsInvalidKind:
		return "DebuggerErrorEventTypeIsInvalid"
	case DebuggerErrorUnableToCreateEventKind:
		return "DebuggerErrorUnableToCreateEvent"
	case DebuggerErrorInvalidAddressKind:
		return "DebuggerErrorInvalidAddress"
	case DebuggerErrorInvalidCoreIdKind:
		return "DebuggerErrorInvalidCoreId"
	case DebuggerErrorExceptionIndexExceedFirst32EntriesKind:
		return "DebuggerErrorExceptionIndexExceedFirst32Entries"
	case DebuggerErrorInterruptIndexIsNotValidKind:
		return "DebuggerErrorInterruptIndexIsNotValid"
	case DebuggerErrorUnableToHideOrUnhideDebuggerKind:
		return "DebuggerErrorUnableToHideOrUnhideDebugger"
	case DebuggerErrorDebuggerAlreadyUhideKind:
		return "DebuggerErrorDebuggerAlreadyUhide"
	case DebuggerErrorEditMemoryStatusInvalidParameterKind:
		return "DebuggerErrorEditMemoryStatusInvalidParameter"
	case DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcessKind:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnCurrentProcess"
	case DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcessKind:
		return "DebuggerErrorEditMemoryStatusInvalidAddressBasedOnOtherProcess"
	case DebuggerErrorModifyEventsInvalidTagKind:
		return "DebuggerErrorModifyEventsInvalidTag"
	case DebuggerErrorModifyEventsInvalidTypeOfActionKind:
		return "DebuggerErrorModifyEventsInvalidTypeOfAction"
	case DebuggerErrorSteppingInvalidParameterKind:
		return "DebuggerErrorSteppingInvalidParameter"
	case DebuggerErrorSteppingsEitherThreadNotFoundOrDisabledKind:
		return "DebuggerErrorSteppingsEitherThreadNotFoundOrDisabled"
	case DebuggerErrorPreparingDebuggeeInvalidBaudrateKind:
		return "DebuggerErrorPreparingDebuggeeInvalidBaudrate"
	case DebuggerErrorPreparingDebuggeeInvalidSerialPortKind:
		return "DebuggerErrorPreparingDebuggeeInvalidSerialPort"
	case DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebuggeKind:
		return "DebuggerErrorPreparingDebuggeeInvalidCoreInRemoteDebugge"
	case DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcessKind:
		return "DebuggerErrorPreparingDebuggeeUnableToSwitchToNewProcess"
	case DebuggerErrorPreparingDebuggeeToRunScriptKind:
		return "DebuggerErrorPreparingDebuggeeToRunScript"
	case DebuggerErrorInvalidRegisterNumberKind:
		return "DebuggerErrorInvalidRegisterNumber"
	case DebuggerErrorMaximumBreakpointWithoutContinueKind:
		return "DebuggerErrorMaximumBreakpointWithoutContinue"
	case DebuggerErrorBreakpointAlreadyExistsOnTheAddressKind:
		return "DebuggerErrorBreakpointAlreadyExistsOnTheAddress"
	case DebuggerErrorBreakpointIdNotFoundKind:
		return "DebuggerErrorBreakpointIdNotFound"
	case DebuggerErrorBreakpointAlreadyDisabledKind:
		return "DebuggerErrorBreakpointAlreadyDisabled"
	case DebuggerErrorBreakpointAlreadyEnabledKind:
		return "DebuggerErrorBreakpointAlreadyEnabled"
	case DebuggerErrorMemoryTypeInvalidKind:
		return "DebuggerErrorMemoryTypeInvalid"
	case DebuggerErrorInvalidProcessIdKind:
		return "DebuggerErrorInvalidProcessId"
	case DebuggerErrorEventIsNotAppliedKind:
		return "DebuggerErrorEventIsNotApplied"
	case DebuggerErrorDetailsOrSwitchProcessInvalidParameterKind:
		return "DebuggerErrorDetailsOrSwitchProcessInvalidParameter"
	case DebuggerErrorDetailsOrSwitchThreadInvalidParameterKind:
		return "DebuggerErrorDetailsOrSwitchThreadInvalidParameter"
	case DebuggerErrorMaximumBreakpointForASinglePageIsHitKind:
		return "DebuggerErrorMaximumBreakpointForASinglePageIsHit"
	case DebuggerErrorPreAllocatedBufferIsEmptyKind:
		return "DebuggerErrorPreAllocatedBufferIsEmpty"
	case DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPagesKind:
		return "DebuggerErrorEptCouldNotSplitTheLargePageTo4KbPages"
	case DebuggerErrorEptFailedToGetPml1EntryOfTargetAddressKind:
		return "DebuggerErrorEptFailedToGetPml1EntryOfTargetAddress"
	case DebuggerErrorEptMultipleHooksInASinglePageKind:
		return "DebuggerErrorEptMultipleHooksInASinglePage"
	case DebuggerErrorCouldNotBuildTheEptHookKind:
		return "DebuggerErrorCouldNotBuildTheEptHook"
	case DebuggerErrorCouldNotFindAllocationTypeKind:
		return "DebuggerErrorCouldNotFindAllocationType"
	case DebuggerErrorInvalidTestQueryIndexKind:
		return "DebuggerErrorInvalidTestQueryIndex"
	case DebuggerErrorUnableToAttachToTargetUserModeProcessKind:
		return "DebuggerErrorUnableToAttachToTargetUserModeProcess"
	case DebuggerErrorUnableToRemoveHooksEntrypointNotReachedKind:
		return "DebuggerErrorUnableToRemoveHooksEntrypointNotReached"
	case DebuggerErrorUnableToRemoveHooksKind:
		return "DebuggerErrorUnableToRemoveHooks"
	case DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitializedKind:
		return "DebuggerErrorFunctionsForInitializingPebAddressesAreNotInitialized"
	case DebuggerErrorUnableToDetect32BitOr64BitProcessKind:
		return "DebuggerErrorUnableToDetect32BitOr64BitProcess"
	case DebuggerErrorUnableToKillTheProcessKind:
		return "DebuggerErrorUnableToKillTheProcess"
	case DebuggerErrorInvalidThreadDebuggingTokenKind:
		return "DebuggerErrorInvalidThreadDebuggingToken"
	case DebuggerErrorUnableToPauseTheProcessThreadsKind:
		return "DebuggerErrorUnableToPauseTheProcessThreads"
	case DebuggerErrorUnableToAttachToAnAlreadyAttachedProcessKind:
		return "DebuggerErrorUnableToAttachToAnAlreadyAttachedProcess"
	case DebuggerErrorTheUserDebuggerNotAttachedToTheProcessKind:
		return "DebuggerErrorTheUserDebuggerNotAttachedToTheProcess"
	case DebuggerErrorUnableToDetachAsThereArePausedThreadsKind:
		return "DebuggerErrorUnableToDetachAsThereArePausedThreads"
	case DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalidKind:
		return "DebuggerErrorUnableToSwitchProcessIdOrThreadIdIsInvalid"
	case DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcessKind:
		return "DebuggerErrorUnableToSwitchThereIsNoThreadOnTheProcess"
	case DebuggerErrorUnableToGetModulesOfTheProcessKind:
		return "DebuggerErrorUnableToGetModulesOfTheProcess"
	case DebuggerErrorUnableToGetCallstackKind:
		return "DebuggerErrorUnableToGetCallstack"
	case DebuggerErrorUnableToQueryCountOfProcessesOrThreadsKind:
		return "DebuggerErrorUnableToQueryCountOfProcessesOrThreads"
	case DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddednKind:
		return "DebuggerErrorUsingShortCircuitingEventWithPostEventModeIsForbiddedn"
	case DebuggerErrorUnknownTestQueryReceivedKind:
		return "DebuggerErrorUnknownTestQueryReceived"
	case DebuggerErrorReadingMemoryInvalidParameterKind:
		return "DebuggerErrorReadingMemoryInvalidParameter"
	case DebuggerErrorTheTrapFlagListIsFullKind:
		return "DebuggerErrorTheTrapFlagListIsFull"
	case DebuggerErrorUnableToKillTheProcessDoesNotExistsKind:
		return "DebuggerErrorUnableToKillTheProcessDoesNotExists"
	case DebuggerErrorModeExecutionIsInvalidKind:
		return "DebuggerErrorModeExecutionIsInvalid"
	case DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootModeKind:
		return "DebuggerErrorProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionalsKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals"
	case DebuggerErrorInstantEventRegularPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventRegularPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventBigPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventBigPreallocatedBufferNotFound"
	case DebuggerErrorUnableToCreateActionCannotAllocateBufferKind:
		return "DebuggerErrorUnableToCreateActionCannotAllocateBuffer"
	case DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventActionRegularPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventActionBigPreallocatedBufferNotFoundKind:
		return "DebuggerErrorInstantEventActionBigPreallocatedBufferNotFound"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBufferKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForActionBuffer"
	case DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStackKind:
		return "DebuggerErrorInstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack"
	case DebuggerErrorInstantEventRegularRequestedSafeBufferNotFoundKind:
		return "DebuggerErrorInstantEventRegularRequestedSafeBufferNotFound"
	case DebuggerErrorInstantEventBigRequestedSafeBufferNotFoundKind:
		return "DebuggerErrorInstantEventBigRequestedSafeBufferNotFound"
	case DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBufferKind:
		return "DebuggerErrorInstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer"
	case DebuggerErrorUnableToAllocateRequestedSafeBufferKind:
		return "DebuggerErrorUnableToAllocateRequestedSafeBuffer"
	case DebuggerErrorCouldNotFindPreactivationTypeKind:
		return "DebuggerErrorCouldNotFindPreactivationType"
	case DebuggerErrorTheModeExecTrapIsNotInitializedKind:
		return "DebuggerErrorTheModeExecTrapIsNotInitialized"
	case DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFullKind:
		return "DebuggerErrorTheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull"
	case DebuggerErrorNotAllCoresAreLockedForApplyingInstantEventKind:
		return "DebuggerErrorNotAllCoresAreLockedForApplyingInstantEvent"
	case DebuggerErrorTargetSwitchingCoreIsNotLockedKind:
		return "DebuggerErrorTargetSwitchingCoreIsNotLocked"
	case DebuggerErrorInvalidPhysicalAddressKind:
		return "DebuggerErrorInvalidPhysicalAddress"
	default:
		return "InvalidDebuggerErrorKind"
	}
}