package sdk

import "github.com/winlabs/gowin32/wrappers"

var (
	PageSize                                                                 = 4096
	WcharMin                                                                 = 0
	WcharMax                                                                 = 65535
	MaxPath                                                                  = 260
	NullZero                                                                 = 0
	Null64Zero                                                               = 0
	False                                                                    = 0
	True                                                                     = 1
	Upper56Bits                                                              = uint64(0xffffffffffffff00)
	Upper48Bits                                                              = uint64(0xffffffffffff0000)
	Upper32Bits                                                              = uint64(0xffffffff00000000)
	Lower32Bits                                                              = uint64(0x00000000ffffffff)
	Lower16Bits                                                              = uint64(0x000000000000ffff)
	Lower8Bits                                                               = uint64(0x00000000000000ff)
	SecondLower8Bits                                                         = uint64(0x000000000000ff00)
	Upper48BitsAndLower8Bits                                                 = uint64(0xffffffffffff00ff)
	VersionMajor                                                             = 0
	VersionMinor                                                             = 10
	VersionPatch                                                             = 0
	MaximumPacketsCapacity                                                   = 1000
	MaximumPacketsCapacityPriority                                           = 50
	NormalPageSize                                                           = 4096 // PAGE_SIZE
	PacketChunkSize                                                          = NormalPageSize
	MaxSerialPacketSize                                                      = 10 * NormalPageSize
	DbgPrintLimitation                                                       = 512
	DebuggerEventTagStartSeed                                                = 0x1000000
	DebuggerThreadDebuggingTagStartSeed                                      = 0x1000000
	DebuggerOutputSourceTagStartSeed                                         = 0x1
	DebuggerOutputSourceMaximumRemoteSourceForSingleEvent                    = 0x5
	DebuggerScriptEngineMemcpyMovingBufferSize                               = 64
	MaximumNumberOfInitialPreallocatedEptHooks                               = 5
	MaximumRegularInstantEvents                                              = 20
	MaximumBigInstantEvents                                                  = 0
	RegularInstantEventRequestedSafeBuffer                                   = PageSize
	BigInstantEventRequestedSafeBuffer                                       = MaxSerialPacketSize
	DefaultPort                                                              = "50000"
	CommunicationBufferSize                                                  = PacketChunkSize + 0x100
	TopLevelDriversVmcallStartingNumber                                      = 0x00000200
	OperationMandatoryDebuggeeBit                                            = 1 << 31
	OperationLogInfoMessage                                                  = 1
	OperationLogWarningMessage                                               = 2
	OperationLogErrorMessage                                                 = 3
	OperationLogNonImmediateMessage                                          = 4
	OperationLogWithTag                                                      = 5
	MaximumBreakpointsWithoutContinue                                        = 100
	MaximumNumberOfThreadInformationForTraps                                 = 200
	Pooltag                                                                  = 0x48444247 // [H]yper[DBG] (HDBG)
	SerialEndOfBufferCharsCount                                              = 0x4
	SerialEndOfBufferChar1                                                   = 0x00
	SerialEndOfBufferChar2                                                   = 0x80
	SerialEndOfBufferChar3                                                   = 0xEE
	SerialEndOfBufferChar4                                                   = 0xFF
	TcpEndOfBufferCharsCount                                                 = 0x4
	TcpEndOfBufferChar1                                                      = 0x10
	TcpEndOfBufferChar2                                                      = 0x20
	TcpEndOfBufferChar3                                                      = 0x33
	TcpEndOfBufferChar4                                                      = 0x44
	MaximumCharacterForOsName                                                = 256
	MaximumInstrSize                                                         = 16
	MaximumCallInstrSize                                                     = 7
	MaximumSupportedSymbols                                                  = 1000
	MaximumGuidAndAgeSize                                                    = 60
	IndicatorOfHyperdbgPacket                                                = 0x4859504552444247 // HYPERDBG = 0x4859504552444247
	MaximumSearchResults                                                     = 0x1000
	X86FlagsCf                                                               = 1 << 0
	X86FlagsPf                                                               = 1 << 2
	X86FlagsAf                                                               = 1 << 4
	X86FlagsZf                                                               = 1 << 6
	X86FlagsSf                                                               = 1 << 7
	X86FlagsTf                                                               = 1 << 8
	X86FlagsIf                                                               = 1 << 9
	X86FlagsDf                                                               = 1 << 10
	X86FlagsOf                                                               = 1 << 11
	X86FlagsStatusMask                                                       = 0xfff
	X86FlagsIoplMask                                                         = 3 << 12
	X86FlagsIoplShift                                                        = 12
	X86FlagsIoplShift2NdBit                                                  = 13
	X86FlagsNt                                                               = 1 << 14
	X86FlagsRf                                                               = 1 << 16
	X86FlagsVm                                                               = 1 << 17
	X86FlagsAc                                                               = 1 << 18
	X86FlagsVif                                                              = 1 << 19
	X86FlagsVip                                                              = 1 << 20
	X86FlagsId                                                               = 1 << 21
	X86FlagsReservedOnes                                                     = 0x2
	X86FlagsReserved                                                         = 0xffc0802a
	X86FlagsReservedBits                                                     = 0xffc38028
	X86FlagsFixed                                                            = 0x00000002
	MaxTempCount                                                             = 128
	MaxStackBufferCount                                                      = 128
	MaxVarCount                                                              = 512
	MaxFunctionNameLength                                                    = 32
	DebuggerModifyEventsApplyToAllTag                                        = uint64(0xffffffffffffffff)
	DisassemblyMaximumDistanceFromObjectName                                 = 0xffff
	DebuggerReadAndWriteOnMsrApplyAllCores                                   = 0xffffffff
	DebuggerDebuggeeIsRunningNoCore                                          = 0xffffffff
	DebuggerEventApplyToAllCores                                             = 0xffffffff
	DebuggerEventApplyToAllProcesses                                         = 0xffffffff
	DebuggerEventMsrReadOrWriteAllMsrs                                       = 0xffffffff
	DebuggerEventExceptionsAllFirst32Entries                                 = 0xffffffff
	DebuggerEventSyscallAllSysretOrSyscalls                                  = 0xffffffff
	DebuggerEventAllIoPorts                                                  = 0xffffffff
	DebuggeeBpApplyToAllCores                                                = 0xffffffff
	DebuggeeBpApplyToAllProcesses                                            = 0xffffffff
	DebuggeeBpApplyToAllThreads                                              = 0xffffffff
	DebuggeeShowAllRegisters                                                 = 0xffffffff
	DebuggerOperationWasSuccessful                                           = 0xFFFFFFFF
	TagNotExists                                                             = 0xc0000000
	InvalidActionType                                                        = 0xc0000001
	ActionBufferSizeIsZero                                                   = 0xc0000002
	EventTypeIsInvalid                                                       = 0xc0000003
	UnableToCreateEvent                                                      = 0xc0000004
	InvalidAddress                                                           = 0xc0000005
	InvalidCoreId                                                            = 0xc0000006
	ExceptionIndexExceedFirst32Entries                                       = 0xc0000007
	InterruptIndexIsNotValid                                                 = 0xc0000008
	UnableToHideOrUnhideDebugger                                             = 0xc0000009
	DebuggerAlreadyUhide                                                     = 0xc000000a
	EditMemoryStatusInvalidParameter                                         = 0xc000000b
	EditMemoryStatusInvalidAddressBasedOnCurrentProcess                      = 0xc000000c
	EditMemoryStatusInvalidAddressBasedOnOtherProcess                        = 0xc000000d
	ModifyEventsInvalidTag                                                   = 0xc000000e
	ModifyEventsInvalidTypeOfAction                                          = 0xc000000f
	SteppingInvalidParameter                                                 = 0xc0000010
	SteppingsEitherThreadNotFoundOrDisabled                                  = 0xc0000011
	PreparingDebuggeeInvalidBaudrate                                         = 0xc0000012
	PreparingDebuggeeInvalidSerialPort                                       = 0xc0000013
	PreparingDebuggeeInvalidCoreInRemoteDebugge                              = 0xc0000014
	PreparingDebuggeeUnableToSwitchToNewProcess                              = 0xc0000015
	PreparingDebuggeeToRunScript                                             = 0xc0000016
	InvalidRegisterNumber                                                    = 0xc0000017
	MaximumBreakpointWithoutContinue                                         = 0xc0000018
	BreakpointAlreadyExistsOnTheAddress                                      = 0xc0000019
	BreakpointIdNotFound                                                     = 0xc000001a
	BreakpointAlreadyDisabled                                                = 0xc000001b
	BreakpointAlreadyEnabled                                                 = 0xc000001c
	MemoryTypeInvalid                                                        = 0xc000001d
	InvalidProcessId                                                         = 0xc000001e
	EventIsNotApplied                                                        = 0xc000001f
	DetailsOrSwitchProcessInvalidParameter                                   = 0xc0000020
	DetailsOrSwitchThreadInvalidParameter                                    = 0xc0000021
	MaximumBreakpointForASinglePageIsHit                                     = 0xc0000022
	PreAllocatedBufferIsEmpty                                                = 0xc0000023
	EptCouldNotSplitTheLargePageTo4KbPages                                   = 0xc0000024
	EptFailedToGetPml1EntryOfTargetAddress                                   = 0xc0000025
	EptMultipleHooksInASinglePage                                            = 0xc0000026
	CouldNotBuildTheEptHook                                                  = 0xc0000027
	CouldNotFindAllocationType                                               = 0xc0000028
	InvalidTestQueryIndex                                                    = 0xc0000029
	UnableToAttachToTargetUserModeProcess                                    = 0xc000002a
	UnableToRemoveHooksEntrypointNotReached                                  = 0xc000002b
	UnableToRemoveHooks                                                      = 0xc000002c
	FunctionsForInitializingPebAddressesAreNotInitialized                    = 0xc000002d
	UnableToDetect32BitOr64BitProcess                                        = 0xc000002e
	UnableToKillTheProcess                                                   = 0xc000002f
	InvalidThreadDebuggingToken                                              = 0xc0000030
	UnableToPauseTheProcessThreads                                           = 0xc0000031
	UnableToAttachToAnAlreadyAttachedProcess                                 = 0xc0000032
	TheUserDebuggerNotAttachedToTheProcess                                   = 0xc0000033
	UnableToDetachAsThereArePausedThreads                                    = 0xc0000034
	UnableToSwitchProcessIdOrThreadIdIsInvalid                               = 0xc0000035
	UnableToSwitchThereIsNoThreadOnTheProcess                                = 0xc0000036
	UnableToGetModulesOfTheProcess                                           = 0xc0000037
	UnableToGetCallstack                                                     = 0xc0000038
	UnableToQueryCountOfProcessesOrThreads                                   = 0xc0000039
	UsingShortCircuitingEventWithPostEventModeIsForbiddedn                   = 0xc000003a
	UnknownTestQueryReceived                                                 = 0xc000003b
	ReadingMemoryInvalidParameter                                            = 0xc000003c
	TheTrapFlagListIsFull                                                    = 0xc000003d
	UnableToKillTheProcessDoesNotExists                                      = 0xc000003e
	ModeExecutionIsInvalid                                                   = 0xc000003f
	ProcessIdCannotBeSpecifiedWhileApplyingEventFromVmxRootMode              = 0xc0000040
	InstantEventPreallocatedBufferIsNotEnoughForEventAndConditionals         = 0xc0000041
	InstantEventRegularPreallocatedBufferNotFound                            = 0xc0000042
	InstantEventBigPreallocatedBufferNotFound                                = 0xc0000043
	UnableToCreateActionCannotAllocateBuffer                                 = 0xc0000044
	InstantEventActionRegularPreallocatedBufferNotFound                      = 0xc0000045
	InstantEventActionBigPreallocatedBufferNotFound                          = 0xc0000046
	InstantEventPreallocatedBufferIsNotEnoughForActionBuffer                 = 0xc0000047
	InstantEventRequestedOptionalBufferIsBiggerThanDebuggersSendReceiveStack = 0xc0000048
	InstantEventRegularRequestedSafeBufferNotFound                           = 0xc0000049
	InstantEventBigRequestedSafeBufferNotFound                               = 0xc000004a
	InstantEventPreallocatedBufferIsNotEnoughForRequestedSafeBuffer          = 0xc000004b
	UnableToAllocateRequestedSafeBuffer                                      = 0xc000004c
	CouldNotFindPreactivationType                                            = 0xc000004d
	TheModeExecTrapIsNotInitialized                                          = 0xc000004e
	TheTargetEventIsDisabledButCannotBeClearedPrirityBufferIsFull            = 0xc000004f
	NotAllCoresAreLockedForApplyingInstantEvent                              = 0xc0000050
	TargetSwitchingCoreIsNotLocked                                           = 0xc0000051
	InvalidPhysicalAddress                                                   = 0xc0000052
	DefaultInitialDebuggeeToDebuggerOffset                                   = 0x200
	DefaultInitialDebuggerToDebuggeeOffset                                   = 0x0
	IoctlRegisterEvent                                                       = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlReturnIrpPendingPacketsAndDisallowIoctl                             = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlTerminateVmx                                                        = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerReadMemory                                                  = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x803, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerReadOrWriteMsr                                              = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x804, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerReadPageTableEntriesDetails                                 = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x805, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerRegisterEvent                                               = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x806, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerAddActionToEvent                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x807, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerHideAndUnhideToTransparentTheDebugger                       = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x808, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerVa2PaAndPa2VaCommands                                       = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x809, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerEditMemory                                                  = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80a, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerSearchMemory                                                = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80b, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerModifyEvents                                                = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80c, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerFlushLoggingBuffers                                         = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80d, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerAttachDetachUserModeProcess                                 = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80e, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerPrint                                                       = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80f, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlPrepareDebuggee                                                     = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x810, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlPausePacketReceived                                                 = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x811, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlSendSignalExecutionInDebuggeeFinished                               = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x812, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlSendUsermodeMessagesToDebugger                                      = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x813, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlSendGeneralBufferFromDebuggeeToDebugger                             = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x814, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlPerfromKernelSideTests                                              = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x815, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlReservePreAllocatedPools                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x816, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlSendUserDebuggerCommands                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x817, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlGetDetailOfActiveThreadsAndProcesses                                = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x818, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlGetUserModeModuleDetails                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x819, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlQueryCountOfActiveProcessesOrThreads                                = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81a, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlGetListOfThreadsAndProcesses                                        = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81b, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlQueryCurrentProcess                                                 = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81c, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlQueryCurrentThread                                                  = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81d, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlRequestRevMachineService                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81e, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlDebuggerBringPagesIn                                                = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81f, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IoctlPreactivateFunctionality                                            = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)
	DebuggerRemoteTrackingDefaultCountOfStepping                             = 0xffffffff
)

func CTL_CODE(deviceType, function, method, access uint32) IoctlKind {
	return IoctlKind(((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method))
}

const (
	FILE_DEVICE_UNKNOWN = wrappers.FILE_DEVICE_UNKNOWN
	METHOD_BUFFERED     = wrappers.METHOD_BUFFERED
	FILE_ANY_ACCESS     = wrappers.FILE_ANY_ACCESS
)
