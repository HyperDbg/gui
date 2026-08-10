package hyperdbgsdk

import (
	"fmt"
	"time"
)

// Source: HyperDbgSdk.h
type (
	_builtin_va_list                                        = *uintptr
	_predefined_size_t                                      = uint64
	_predefined_wchar_t                                     = uint16
	_predefined_ptrdiff_t                                   = int64
	Wchar_t                                                 = uint16
	Size_t                                                  = uint64
	SIZE_T                                                  = uintptr
	Ptrdiff_t                                               = int64
	Uintptr_t                                               = uint64
	Intptr_t                                                = int64
	Uint8_t                                                 = uint8
	Uint16_t                                                = uint16
	Uint32_t                                                = uint32
	Uint64_t                                                = uint64
	Int8_t                                                  = int8
	Int16_t                                                 = int16
	Int32_t                                                 = int32
	Int64_t                                                 = int64
	Time_t                                                  = time.Duration
	Errno_t                                                 = int32
	_time32_t                                               = int32
	_time64_t                                               = time.Duration
	Va_list                                                 = *int8
	PLIST_ENTRY                                             = *LIST_ENTRY
	BOOL                                                    = int32
	BOOLEAN                                                 = bool
	DWORD                                                   = uint32
	WORD                                                    = uint16
	PVOID                                                   = uintptr
	LPVOID                                                  = uintptr
	HANDLE                                                  = uintptr
	QWORD                                                   = uint64
	UINT64                                                  = uint64
	PUINT64                                                 = *uint64
	UINT32                                                  = uint32
	PUINT32                                                 = *uint32
	UINT16                                                  = uint16
	PUINT16                                                 = *uint16
	UINT8                                                   = uint8
	PUINT8                                                  = *uint8
	INT64                                                   = int64
	PINT64                                                  = *int64
	INT32                                                   = int32
	PINT32                                                  = *int32
	INT16                                                   = int16
	PINT16                                                  = *int16
	INT8                                                    = int8
	PINT8                                                   = *int8
	ULONG64                                                 = uint64
	PULONG64                                                = *uint64
	DWORD64                                                 = uint64
	PDWORD64                                                = *uint64
	CHAR                                                    = int8
	WCHAR                                                   = uint16
	UCHAR                                                   = uint8
	SHORT                                                   = int16
	USHORT                                                  = uint16
	INT                                                     = int32
	UINT                                                    = uint32
	PUINT                                                   = *uint32
	LONG                                                    = int32
	ULONG                                                   = uint32
	PBOOLEAN                                                = *uint8
	BYTE                                                    = uint8
	PGUEST_REGS                                             = *GUEST_REGS
	PGUEST_XMM_REGS                                         = *GUEST_XMM_REGS
	PGUEST_EXTRA_REGISTERS                                  = *GUEST_EXTRA_REGISTERS
	PSCRIPT_ENGINE_GENERAL_REGISTERS                        = *SCRIPT_ENGINE_GENERAL_REGISTERS
	PCR3_TYPE                                               = *CR3_TYPE
	PDEBUGGER_REMOTE_PACKET                                 = *DEBUGGER_REMOTE_PACKET
	PDEBUGGEE_USER_INPUT_PACKET                             = *DEBUGGEE_USER_INPUT_PACKET
	PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET     = *DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
	PDEBUGGER_PAUSE_PACKET_RECEIVED                         = *DEBUGGER_PAUSE_PACKET_RECEIVED
	PDEBUGGER_TRIGGERED_EVENT_DETAILS                       = *DEBUGGER_TRIGGERED_EVENT_DETAILS
	PDEBUGGEE_KD_PAUSED_PACKET                              = *DEBUGGEE_KD_PAUSED_PACKET
	PDEBUGGEE_UD_PAUSED_PACKET                              = *DEBUGGEE_UD_PAUSED_PACKET
	PDEBUGGEE_MESSAGE_PACKET                                = *DEBUGGEE_MESSAGE_PACKET
	PREGISTER_NOTIFY_BUFFER                                 = *REGISTER_NOTIFY_BUFFER
	PDIRECT_VMCALL_PARAMETERS                               = *DIRECT_VMCALL_PARAMETERS
	PSYSCALL_CALLBACK_CONTEXT_PARAMS                        = *SYSCALL_CALLBACK_CONTEXT_PARAMS
	PEPT_HOOKS_CONTEXT                                      = *EPT_HOOKS_CONTEXT
	PEPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR           = *EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR
	PEPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2                 = *EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2
	PEPT_SINGLE_HOOK_UNHOOKING_DETAILS                      = *EPT_SINGLE_HOOK_UNHOOKING_DETAILS
	PVMX_SEGMENT_SELECTOR                                   = *VMX_SEGMENT_SELECTOR
	PDEBUGGER_MODIFY_EVENTS                                 = *DEBUGGER_MODIFY_EVENTS
	PDEBUGGER_SHORT_CIRCUITING_EVENT                        = *DEBUGGER_SHORT_CIRCUITING_EVENT
	PDEBUGGER_EVENT_OPTIONS                                 = *DEBUGGER_EVENT_OPTIONS
	PDEBUGGER_GENERAL_EVENT_DETAIL                          = *DEBUGGER_GENERAL_EVENT_DETAIL
	PDEBUGGER_GENERAL_ACTION                                = *DEBUGGER_GENERAL_ACTION
	PDEBUGGER_EVENT_AND_ACTION_RESULT                       = *DEBUGGER_EVENT_AND_ACTION_RESULT
	PPORTABLE_PCI_COMMON_HEADER                             = *PORTABLE_PCI_COMMON_HEADER
	PPORTABLE_PCI_DEVICE_HEADER                             = *PORTABLE_PCI_DEVICE_HEADER
	PPORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL               = *PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL
	PPCI_DEV_MINIMAL                                        = *PCI_DEV_MINIMAL
	PPCI_DEV_MMIOBAR_INFO                                   = *PCI_DEV_MMIOBAR_INFO
	PPORTABLE_PCI_CONFIG_SPACE_HEADER                       = *PORTABLE_PCI_CONFIG_SPACE_HEADER
	PPCI_DEV                                                = *PCI_DEV
	PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS               = *DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
	PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS                      = *DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
	PDEBUGGER_PAGE_IN_REQUEST                               = *DEBUGGER_PAGE_IN_REQUEST
	PREVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST           = *REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST
	PDEBUGGER_DT_COMMAND_OPTIONS                            = *DEBUGGER_DT_COMMAND_OPTIONS
	PDEBUGGER_PREALLOC_COMMAND                              = *DEBUGGER_PREALLOC_COMMAND
	PDEBUGGER_PREACTIVATE_COMMAND                           = *DEBUGGER_PREACTIVATE_COMMAND
	PDEBUGGER_READ_MEMORY                                   = *DEBUGGER_READ_MEMORY
	PDEBUGGER_FLUSH_LOGGING_BUFFERS                         = *DEBUGGER_FLUSH_LOGGING_BUFFERS
	PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER                    = *DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
	PDEBUGGER_PERFORM_KERNEL_TESTS                          = *DEBUGGER_PERFORM_KERNEL_TESTS
	PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL        = *DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
	PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER = *DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
	PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER            = *DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
	PDEBUGGER_READ_AND_WRITE_ON_MSR                         = *DEBUGGER_READ_AND_WRITE_ON_MSR
	PDEBUGGER_EDIT_MEMORY                                   = *DEBUGGER_EDIT_MEMORY
	PDEBUGGER_SEARCH_MEMORY                                 = *DEBUGGER_SEARCH_MEMORY
	PSYSTEM_CALL_NUMBERS_INFORMATION                        = *SYSTEM_CALL_NUMBERS_INFORMATION
	PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE            = *DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	PDEBUGGER_PREPARE_DEBUGGEE                              = *DEBUGGER_PREPARE_DEBUGGEE
	PDEBUGGEE_CHANGE_CORE_PACKET                            = *DEBUGGEE_CHANGE_CORE_PACKET
	PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS               = *DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
	PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS                   = *DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
	PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS                    = *DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
	PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY                    = *DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY
	PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY                     = *DEBUGGEE_THREAD_LIST_DETAILS_ENTRY
	PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS             = *DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
	PDEBUGGER_SINGLE_CALLSTACK_FRAME                        = *DEBUGGER_SINGLE_CALLSTACK_FRAME
	PDEBUGGER_CALLSTACK_REQUEST                             = *DEBUGGER_CALLSTACK_REQUEST
	PUSERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS     = *USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS
	PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION         = *DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
	PDEBUGGER_EVENT_REQUEST_BUFFER                          = *DEBUGGER_EVENT_REQUEST_BUFFER
	PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE                     = *DEBUGGER_EVENT_REQUEST_CUSTOM_CODE
	PDEBUGGER_UD_COMMAND_ACTION                             = *DEBUGGER_UD_COMMAND_ACTION
	PDEBUGGER_UD_COMMAND_PACKET                             = *DEBUGGER_UD_COMMAND_PACKET
	PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET             = *DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
	PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET              = *DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
	PDEBUGGEE_STEP_PACKET                                   = *DEBUGGEE_STEP_PACKET
	PDEBUGGER_APIC_REQUEST                                  = *DEBUGGER_APIC_REQUEST
	PLAPIC_PAGE                                             = *LAPIC_PAGE
	PIO_APIC_ENTRY_PACKETS                                  = *IO_APIC_ENTRY_PACKETS
	PSMI_OPERATION_PACKETS                                  = *SMI_OPERATION_PACKETS
	PHYPERTRACE_LBR_OPERATION_PACKETS                       = *HYPERTRACE_LBR_OPERATION_PACKETS
	PHYPERTRACE_PT_OPERATION_PACKETS                        = *HYPERTRACE_PT_OPERATION_PACKETS
	PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS             = *INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS
	PDEBUGGEE_FORMATS_PACKET                                = *DEBUGGEE_FORMATS_PACKET
	PDEBUGGEE_SYMBOL_REQUEST_PACKET                         = *DEBUGGEE_SYMBOL_REQUEST_PACKET
	PDEBUGGEE_BP_PACKET                                     = *DEBUGGEE_BP_PACKET
	PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET                      = *DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
	PDEBUGGEE_SCRIPT_PACKET                                 = *DEBUGGEE_SCRIPT_PACKET
	PDEBUGGEE_RESULT_OF_SEARCH_PACKET                       = *DEBUGGEE_RESULT_OF_SEARCH_PACKET
	PDEBUGGEE_REGISTER_READ_DESCRIPTION                     = *DEBUGGEE_REGISTER_READ_DESCRIPTION
	PDEBUGGEE_REGISTER_WRITE_DESCRIPTION                    = *DEBUGGEE_REGISTER_WRITE_DESCRIPTION
	PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET               = *DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET
	PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET            = *DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET
	PMODULE_SYMBOL_DETAIL                                   = *MODULE_SYMBOL_DETAIL
	PUSERMODE_LOADED_MODULE_SYMBOLS                         = *USERMODE_LOADED_MODULE_SYMBOLS
	PUSERMODE_LOADED_MODULE_DETAILS                         = *USERMODE_LOADED_MODULE_DETAILS
	PDEBUGGER_UPDATE_SYMBOL_TABLE                           = *DEBUGGER_UPDATE_SYMBOL_TABLE
	PDEBUGGEE_SYMBOL_UPDATE_RESULT                          = *DEBUGGEE_SYMBOL_UPDATE_RESULT
	PHWDBG_PORT_INFORMATION_ITEMS                           = *HWDBG_PORT_INFORMATION_ITEMS
	PHWDBG_INSTANCE_INFORMATION                             = *HWDBG_INSTANCE_INFORMATION
	PHWDBG_SCRIPT_BUFFER                                    = *HWDBG_SCRIPT_BUFFER
	PSYMBOL                                                 = *SYMBOL
	PHWDBG_SHORT_SYMBOL                                     = *HWDBG_SHORT_SYMBOL
	PSYMBOL_BUFFER                                          = *SYMBOL_BUFFER
	PSYMBOL_MAP                                             = *SYMBOL_MAP
	PACTION_BUFFER                                          = *ACTION_BUFFER
	SymbolTypeNames                                         = [22]*int8
	FunctionNames                                           = [104]*int8
	RegistersNames                                          = [120]*int8
)

type SendMessageWithParamCallback func(*int8) uintptr

type SendMessageWithSharedBufferCallback func() uintptr

type SymbolMapCallback func(uint64, *int8, *int8, uint32) uintptr

// Source: unknown.h:584 -> _SEGMENT_REGISTERS
type SEGMENT_REGISTERS uint32

const (
	ES SEGMENT_REGISTERS = iota
	CS
	SS
	DS
	FS
	GS
	LDTR
	TR
)

func (s SEGMENT_REGISTERS) String() string {
	switch s {
	case ES:
		return "ES"
	case CS:
		return "CS"
	case SS:
		return "SS"
	case DS:
		return "DS"
	case FS:
		return "FS"
	case GS:
		return "GS"
	case LDTR:
		return "LDTR"
	case TR:
		return "TR"
	default:
		return fmt.Sprintf("SEGMENT_REGISTERS(0x%X)", uint32(s))
	}
}

// Source: unknown.h:18 -> _DEBUGGEE_PAUSING_REASON
type DEBUGGEE_PAUSING_REASON uint32

const (
	DebuggeePausingReasonNotPaused DEBUGGEE_PAUSING_REASON = iota
	DebuggeePausingReasonPause
	DebuggeePausingReasonRequestFromDebugger
	DebuggeePausingReasonDebuggeeStepped
	DebuggeePausingReasonDebuggeeTrackingStepped
	DebuggeePausingReasonDebuggeeSoftwareBreakpointHit
	DebuggeePausingReasonDebuggeeHardwareDebugRegisterHit
	DebuggeePausingReasonDebuggeeCoreSwitched
	DebuggeePausingReasonDebuggeeProcessSwitched
	DebuggeePausingReasonDebuggeeThreadSwitched
	DebuggeePausingReasonDebuggeeCommandExecutionFinished
	DebuggeePausingReasonDebuggeeEventTriggered
	DebuggeePausingReasonDebuggeeStartingModuleLoaded
	DebuggeePausingReasonDebuggeeGeneralDebugBreak
	DebuggeePausingReasonDebuggeeGeneralThreadIntercepted
	DebuggeePausingReasonHardwareBasedDebuggeeGeneralBreak
)

func (d DEBUGGEE_PAUSING_REASON) String() string {
	switch d {
	case DebuggeePausingReasonNotPaused:
		return "Debuggee Pausing Reason Not Paused"
	case DebuggeePausingReasonPause:
		return "Debuggee Pausing Reason Pause"
	case DebuggeePausingReasonRequestFromDebugger:
		return "Debuggee Pausing Reason Request From Debugger"
	case DebuggeePausingReasonDebuggeeStepped:
		return "Debuggee Pausing Reason Debuggee Stepped"
	case DebuggeePausingReasonDebuggeeTrackingStepped:
		return "Debuggee Pausing Reason Debuggee Tracking Stepped"
	case DebuggeePausingReasonDebuggeeSoftwareBreakpointHit:
		return "Debuggee Pausing Reason Debuggee Software Breakpoint Hit"
	case DebuggeePausingReasonDebuggeeHardwareDebugRegisterHit:
		return "Debuggee Pausing Reason Debuggee Hardware Debug Register Hit"
	case DebuggeePausingReasonDebuggeeCoreSwitched:
		return "Debuggee Pausing Reason Debuggee Core Switched"
	case DebuggeePausingReasonDebuggeeProcessSwitched:
		return "Debuggee Pausing Reason Debuggee Process Switched"
	case DebuggeePausingReasonDebuggeeThreadSwitched:
		return "Debuggee Pausing Reason Debuggee Thread Switched"
	case DebuggeePausingReasonDebuggeeCommandExecutionFinished:
		return "Debuggee Pausing Reason Debuggee Command Execution Finished"
	case DebuggeePausingReasonDebuggeeEventTriggered:
		return "Debuggee Pausing Reason Debuggee Event Triggered"
	case DebuggeePausingReasonDebuggeeStartingModuleLoaded:
		return "Debuggee Pausing Reason Debuggee Starting Module Loaded"
	case DebuggeePausingReasonDebuggeeGeneralDebugBreak:
		return "Debuggee Pausing Reason Debuggee General Debug Break"
	case DebuggeePausingReasonDebuggeeGeneralThreadIntercepted:
		return "Debuggee Pausing Reason Debuggee General Thread Intercepted"
	case DebuggeePausingReasonHardwareBasedDebuggeeGeneralBreak:
		return "Debuggee Pausing Reason Hardware Based Debuggee General Break"
	default:
		return fmt.Sprintf("DEBUGGEE_PAUSING_REASON(0x%X)", uint32(d))
	}
}

// Source: unknown.h:55 -> _DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
type DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION uint32

const (
	DebuggerRemotePacketRequestedActionOnUserModePause DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 1 + iota
	DebuggerRemotePacketRequestedActionOnUserModeDoNotReadAnyPacket
	DebuggerRemotePacketRequestedActionOnUserModeDebuggerVersion
	DebuggerRemotePacketPingAndSendSupportedVersion
	DebuggerRemotePacketRequestedActionOnVmxRootModeStep
	DebuggerRemotePacketRequestedActionOnVmxRootModeContinue
	DebuggerRemotePacketRequestedActionOnVmxRootModeCloseAndUnloadDebuggee
	DebuggerRemotePacketRequestedActionOnVmxRootModeChangeCore
	DebuggerRemotePacketRequestedActionOnVmxRootModeFlushBuffers
	DebuggerRemotePacketRequestedActionOnVmxRootModeCallstack
	DebuggerRemotePacketRequestedActionOnVmxRootModeTestQuery
	DebuggerRemotePacketRequestedActionOnVmxRootModeChangeProcess
	DebuggerRemotePacketRequestedActionOnVmxRootModeChangeThread
	DebuggerRemotePacketRequestedActionOnVmxRootRunScript
	DebuggerRemotePacketRequestedActionOnVmxRootUserInputBuffer
	DebuggerRemotePacketRequestedActionOnVmxRootSearchQuery
	DebuggerRemotePacketRequestedActionOnVmxRootRegisterEvent
	DebuggerRemotePacketRequestedActionOnVmxRootAddActionToEvent
	DebuggerRemotePacketRequestedActionOnVmxRootQueryAndModifyEvent
	DebuggerRemotePacketRequestedActionOnVmxRootReadRegisters
	DebuggerRemotePacketRequestedActionOnVmxRootReadMemory
	DebuggerRemotePacketRequestedActionOnVmxRootEditMemory
	DebuggerRemotePacketRequestedActionOnVmxRootBp
	DebuggerRemotePacketRequestedActionOnVmxRootListOrModifyBreakpoints
	DebuggerRemotePacketRequestedActionOnVmxRootSymbolReload
	DebuggerRemotePacketRequestedActionOnVmxRootQueryPa2vaAndVa2pa
	DebuggerRemotePacketRequestedActionOnVmxRootSymbolQueryPte
	DebuggerRemotePacketRequestedActionOnVmxRootSetShortCircuitingState
	DebuggerRemotePacketRequestedActionOnVmxRootInjectPageFault
	DebuggerRemotePacketRequestedActionOnVmxRootWriteRegister
	DebuggerRemotePacketRequestedActionOnVmxRootQueryPcitree
	DebuggerRemotePacketRequestedActionOnVmxRootPerformActionsOnApic
	DebuggerRemotePacketRequestedActionOnVmxRootQueryPcidevinfo
	DebuggerRemotePacketRequestedActionOnVmxRootReadIdtEntries
	DebuggerRemotePacketRequestedActionOnVmxRootPerformSmiOperation
	DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertraceLbrOperation
	DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertracePtOperation
	DebuggerRemotePacketRequestedActionNoAction
	DebuggerRemotePacketRequestedActionDebuggeeStarted
	DebuggerRemotePacketRequestedActionDebuggeeLoggingMechanism
	DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction
	DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingCore
	DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingProcess
	DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingThread
	DebuggerRemotePacketRequestedActionDebuggeeResultOfRunningScript
	DebuggerRemotePacketRequestedActionDebuggeeResultOfFormats
	DebuggerRemotePacketRequestedActionDebuggeeResultOfFlush
	DebuggerRemotePacketRequestedActionDebuggeeResultOfCallstack
	DebuggerRemotePacketRequestedActionDebuggeeResultTestQuery
	DebuggerRemotePacketRequestedActionDebuggeeResultOfRegisteringEvent
	DebuggerRemotePacketRequestedActionDebuggeeResultOfAddingActionToEvent
	DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryAndModifyEvent
	DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingEvent
	DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingRegisters
	DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingMemory
	DebuggerRemotePacketRequestedActionDebuggeeResultOfEditingMemory
	DebuggerRemotePacketRequestedActionDebuggeeResultOfBp
	DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingState
	DebuggerRemotePacketRequestedActionDebuggeeResultOfListOrModifyBreakpoints
	DebuggerRemotePacketRequestedActionDebuggeeUpdateSymbolInfo
	DebuggerRemotePacketRequestedActionDebuggeeReloadSymbolFinished
	DebuggerRemotePacketRequestedActionDebuggeeReloadSearchQuery
	DebuggerRemotePacketRequestedActionDebuggeeResultOfPte
	DebuggerRemotePacketRequestedActionDebuggeeResultOfVa2paAndPa2va
	DebuggerRemotePacketRequestedActionDebuggeeResultOfBringingPagesIn
	DebuggerRemotePacketRequestedActionDebuggeeResultOfWriteRegister
	DebuggerRemotePacketRequestedActionDebuggeeResultOfPcitree
	DebuggerRemotePacketRequestedActionDebuggeeResultOfApicRequests
	DebuggerRemotePacketRequestedActionDebuggeeResultOfPcidevinfo
	DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryIdtEntriesRequests
	DebuggerRemotePacketRequestedActionDebuggeeResultSmiOperationRequests
	DebuggerRemotePacketRequestedActionDebuggeeResultHypertraceLbrOperationRequests
	DebuggerRemotePacketRequestedActionDebuggeeResultHypertracePtOperationRequests
)

func (d DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION) String() string {
	switch d {
	case DebuggerRemotePacketRequestedActionOnUserModePause:
		return "Debugger Remote Packet Requested Action On User Mode Pause"
	case DebuggerRemotePacketRequestedActionOnUserModeDoNotReadAnyPacket:
		return "Debugger Remote Packet Requested Action On User Mode Do Not Read Any Packet"
	case DebuggerRemotePacketRequestedActionOnUserModeDebuggerVersion:
		return "Debugger Remote Packet Requested Action On User Mode Debugger Version"
	case DebuggerRemotePacketPingAndSendSupportedVersion:
		return "Debugger Remote Packet Ping And Send Supported Version"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeStep:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Step"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeContinue:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Continue"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeCloseAndUnloadDebuggee:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Close And Unload Debuggee"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeChangeCore:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Change Core"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeFlushBuffers:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Flush Buffers"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeCallstack:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Callstack"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeTestQuery:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Test Query"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeChangeProcess:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Change Process"
	case DebuggerRemotePacketRequestedActionOnVmxRootModeChangeThread:
		return "Debugger Remote Packet Requested Action On Vmx Root Mode Change Thread"
	case DebuggerRemotePacketRequestedActionOnVmxRootRunScript:
		return "Debugger Remote Packet Requested Action On Vmx Root Run Script"
	case DebuggerRemotePacketRequestedActionOnVmxRootUserInputBuffer:
		return "Debugger Remote Packet Requested Action On Vmx Root User Input Buffer"
	case DebuggerRemotePacketRequestedActionOnVmxRootSearchQuery:
		return "Debugger Remote Packet Requested Action On Vmx Root Search Query"
	case DebuggerRemotePacketRequestedActionOnVmxRootRegisterEvent:
		return "Debugger Remote Packet Requested Action On Vmx Root Register Event"
	case DebuggerRemotePacketRequestedActionOnVmxRootAddActionToEvent:
		return "Debugger Remote Packet Requested Action On Vmx Root Add Action To Event"
	case DebuggerRemotePacketRequestedActionOnVmxRootQueryAndModifyEvent:
		return "Debugger Remote Packet Requested Action On Vmx Root Query And Modify Event"
	case DebuggerRemotePacketRequestedActionOnVmxRootReadRegisters:
		return "Debugger Remote Packet Requested Action On Vmx Root Read Registers"
	case DebuggerRemotePacketRequestedActionOnVmxRootReadMemory:
		return "Debugger Remote Packet Requested Action On Vmx Root Read Memory"
	case DebuggerRemotePacketRequestedActionOnVmxRootEditMemory:
		return "Debugger Remote Packet Requested Action On Vmx Root Edit Memory"
	case DebuggerRemotePacketRequestedActionOnVmxRootBp:
		return "Debugger Remote Packet Requested Action On Vmx Root Bp"
	case DebuggerRemotePacketRequestedActionOnVmxRootListOrModifyBreakpoints:
		return "Debugger Remote Packet Requested Action On Vmx Root List Or Modify Breakpoints"
	case DebuggerRemotePacketRequestedActionOnVmxRootSymbolReload:
		return "Debugger Remote Packet Requested Action On Vmx Root Symbol Reload"
	case DebuggerRemotePacketRequestedActionOnVmxRootQueryPa2vaAndVa2pa:
		return "Debugger Remote Packet Requested Action On Vmx Root Query Pa 2va And Va 2pa"
	case DebuggerRemotePacketRequestedActionOnVmxRootSymbolQueryPte:
		return "Debugger Remote Packet Requested Action On Vmx Root Symbol Query Pte"
	case DebuggerRemotePacketRequestedActionOnVmxRootSetShortCircuitingState:
		return "Debugger Remote Packet Requested Action On Vmx Root Set Short Circuiting State"
	case DebuggerRemotePacketRequestedActionOnVmxRootInjectPageFault:
		return "Debugger Remote Packet Requested Action On Vmx Root Inject Page Fault"
	case DebuggerRemotePacketRequestedActionOnVmxRootWriteRegister:
		return "Debugger Remote Packet Requested Action On Vmx Root Write Register"
	case DebuggerRemotePacketRequestedActionOnVmxRootQueryPcitree:
		return "Debugger Remote Packet Requested Action On Vmx Root Query Pcitree"
	case DebuggerRemotePacketRequestedActionOnVmxRootPerformActionsOnApic:
		return "Debugger Remote Packet Requested Action On Vmx Root Perform Actions On Apic"
	case DebuggerRemotePacketRequestedActionOnVmxRootQueryPcidevinfo:
		return "Debugger Remote Packet Requested Action On Vmx Root Query Pcidevinfo"
	case DebuggerRemotePacketRequestedActionOnVmxRootReadIdtEntries:
		return "Debugger Remote Packet Requested Action On Vmx Root Read Idt Entries"
	case DebuggerRemotePacketRequestedActionOnVmxRootPerformSmiOperation:
		return "Debugger Remote Packet Requested Action On Vmx Root Perform Smi Operation"
	case DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertraceLbrOperation:
		return "Debugger Remote Packet Requested Action On Vmx Root Perform Hypertrace Lbr Operation"
	case DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertracePtOperation:
		return "Debugger Remote Packet Requested Action On Vmx Root Perform Hypertrace Pt Operation"
	case DebuggerRemotePacketRequestedActionNoAction:
		return "Debugger Remote Packet Requested Action No Action"
	case DebuggerRemotePacketRequestedActionDebuggeeStarted:
		return "Debugger Remote Packet Requested Action Debuggee Started"
	case DebuggerRemotePacketRequestedActionDebuggeeLoggingMechanism:
		return "Debugger Remote Packet Requested Action Debuggee Logging Mechanism"
	case DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction:
		return "Debugger Remote Packet Requested Action Debuggee Paused And Current Instruction"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingCore:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Changing Core"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingProcess:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Changing Process"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingThread:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Changing Thread"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfRunningScript:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Running Script"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfFormats:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Formats"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfFlush:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Flush"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfCallstack:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Callstack"
	case DebuggerRemotePacketRequestedActionDebuggeeResultTestQuery:
		return "Debugger Remote Packet Requested Action Debuggee Result Test Query"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfRegisteringEvent:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Registering Event"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfAddingActionToEvent:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Adding Action To Event"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryAndModifyEvent:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Query And Modify Event"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingEvent:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Short Circuiting Event"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingRegisters:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Reading Registers"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingMemory:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Reading Memory"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfEditingMemory:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Editing Memory"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfBp:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Bp"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingState:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Short Circuiting State"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfListOrModifyBreakpoints:
		return "Debugger Remote Packet Requested Action Debuggee Result Of List Or Modify Breakpoints"
	case DebuggerRemotePacketRequestedActionDebuggeeUpdateSymbolInfo:
		return "Debugger Remote Packet Requested Action Debuggee Update Symbol Info"
	case DebuggerRemotePacketRequestedActionDebuggeeReloadSymbolFinished:
		return "Debugger Remote Packet Requested Action Debuggee Reload Symbol Finished"
	case DebuggerRemotePacketRequestedActionDebuggeeReloadSearchQuery:
		return "Debugger Remote Packet Requested Action Debuggee Reload Search Query"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfPte:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Pte"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfVa2paAndPa2va:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Va 2pa And Pa 2va"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfBringingPagesIn:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Bringing Pages In"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfWriteRegister:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Write Register"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfPcitree:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Pcitree"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfApicRequests:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Apic Requests"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfPcidevinfo:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Pcidevinfo"
	case DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryIdtEntriesRequests:
		return "Debugger Remote Packet Requested Action Debuggee Result Of Query Idt Entries Requests"
	case DebuggerRemotePacketRequestedActionDebuggeeResultSmiOperationRequests:
		return "Debugger Remote Packet Requested Action Debuggee Result Smi Operation Requests"
	case DebuggerRemotePacketRequestedActionDebuggeeResultHypertraceLbrOperationRequests:
		return "Debugger Remote Packet Requested Action Debuggee Result Hypertrace Lbr Operation Requests"
	case DebuggerRemotePacketRequestedActionDebuggeeResultHypertracePtOperationRequests:
		return "Debugger Remote Packet Requested Action Debuggee Result Hypertrace Pt Operation Requests"
	default:
		return fmt.Sprintf("DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION(0x%X)", uint32(d))
	}
}

// Source: unknown.h:162 -> _DEBUGGER_REMOTE_PACKET_TYPE
type DEBUGGER_REMOTE_PACKET_TYPE uint32

const (
	DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot DEBUGGER_REMOTE_PACKET_TYPE = 1 + iota
	DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnUserMode
	DebuggerRemotePacketTypeDebuggeeToDebugger
	DebuggerRemotePacketTypeDebuggerToDebuggeeHardwareLevel
	DebuggerRemotePacketTypeDebuggeeToDebuggerHardwareLevel
)

func (d DEBUGGER_REMOTE_PACKET_TYPE) String() string {
	switch d {
	case DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot:
		return "Debugger Remote Packet Type Debugger To Debuggee Execute On Vmx Root"
	case DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnUserMode:
		return "Debugger Remote Packet Type Debugger To Debuggee Execute On User Mode"
	case DebuggerRemotePacketTypeDebuggeeToDebugger:
		return "Debugger Remote Packet Type Debuggee To Debugger"
	case DebuggerRemotePacketTypeDebuggerToDebuggeeHardwareLevel:
		return "Debugger Remote Packet Type Debugger To Debuggee Hardware Level"
	case DebuggerRemotePacketTypeDebuggeeToDebuggerHardwareLevel:
		return "Debugger Remote Packet Type Debuggee To Debugger Hardware Level"
	default:
		return fmt.Sprintf("DEBUGGER_REMOTE_PACKET_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:23 -> _PAGING_LEVEL
type PAGING_LEVEL uint32

const (
	PagingLevelPageTable PAGING_LEVEL = iota
	PagingLevelPageDirectory
	PagingLevelPageDirectoryPointerTable
	PagingLevelPageMapLevel4
)

func (p PAGING_LEVEL) String() string {
	switch p {
	case PagingLevelPageTable:
		return "Paging Level Page Table"
	case PagingLevelPageDirectory:
		return "Paging Level Page Directory"
	case PagingLevelPageDirectoryPointerTable:
		return "Paging Level Page Directory Pointer Table"
	case PagingLevelPageMapLevel4:
		return "Paging Level Page Map Level 4"
	default:
		return fmt.Sprintf("PAGING_LEVEL(0x%X)", uint32(p))
	}
}

// Source: unknown.h:57 -> _POOL_ALLOCATION_INTENTION
type POOL_ALLOCATION_INTENTION uint32

const (
	TrackingHookedPages POOL_ALLOCATION_INTENTION = iota
	ExecTrampoline
	Split2mbPagingTo4kbPage
	DetourHookDetails
	BreakpointDefinitionStructure
	ProcessThreadHolder
	InstantRegularEventBuffer
	InstantBigEventBuffer
	InstantRegularEventActionBuffer
	InstantBigEventActionBuffer
	InstantRegularSafeBufferForEvents
	InstantBigSafeBufferForEvents
)

func (p POOL_ALLOCATION_INTENTION) String() string {
	switch p {
	case TrackingHookedPages:
		return "Tracking Hooked Pages"
	case ExecTrampoline:
		return "Exec Trampoline"
	case Split2mbPagingTo4kbPage:
		return "Split 2mb Paging To 4kb Page"
	case DetourHookDetails:
		return "Detour Hook Details"
	case BreakpointDefinitionStructure:
		return "Breakpoint Definition Structure"
	case ProcessThreadHolder:
		return "Process Thread Holder"
	case InstantRegularEventBuffer:
		return "Instant Regular Event Buffer"
	case InstantBigEventBuffer:
		return "Instant Big Event Buffer"
	case InstantRegularEventActionBuffer:
		return "Instant Regular Event Action Buffer"
	case InstantBigEventActionBuffer:
		return "Instant Big Event Action Buffer"
	case InstantRegularSafeBufferForEvents:
		return "Instant Regular Safe Buffer For Events"
	case InstantBigSafeBufferForEvents:
		return "Instant Big Safe Buffer For Events"
	default:
		return fmt.Sprintf("POOL_ALLOCATION_INTENTION(0x%X)", uint32(p))
	}
}

// Source: unknown.h:86 -> _DEBUG_REGISTER_TYPE
type DEBUG_REGISTER_TYPE uint32

const (
	BreakOnInstructionFetch DEBUG_REGISTER_TYPE = iota
	BreakOnWriteOnly
	BreakOnIoReadOrWriteNotSupported
	BreakOnReadAndWriteButNotFetch
)

func (d DEBUG_REGISTER_TYPE) String() string {
	switch d {
	case BreakOnInstructionFetch:
		return "Break On Instruction Fetch"
	case BreakOnWriteOnly:
		return "Break On Write Only"
	case BreakOnIoReadOrWriteNotSupported:
		return "Break On Io Read Or Write Not Supported"
	case BreakOnReadAndWriteButNotFetch:
		return "Break On Read And Write But Not Fetch"
	default:
		return fmt.Sprintf("DEBUG_REGISTER_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:98 -> _VMX_EXECUTION_MODE
type VMX_EXECUTION_MODE uint32

const (
	VmxExecutionModeNonRoot VMX_EXECUTION_MODE = iota
	VmxExecutionModeRoot
)

func (v VMX_EXECUTION_MODE) String() string {
	switch v {
	case VmxExecutionModeNonRoot:
		return "Vmx Execution Mode Non Root"
	case VmxExecutionModeRoot:
		return "Vmx Execution Mode Root"
	default:
		return fmt.Sprintf("VMX_EXECUTION_MODE(0x%X)", uint32(v))
	}
}

// Source: unknown.h:108 -> _VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
type VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE uint32

const (
	VmmCallbackCallingStageInvalidEventEmulation VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE = iota
	VmmCallbackCallingStagePreEventEmulation
	VmmCallbackCallingStagePostEventEmulation
	VmmCallbackCallingStageAllEventEmulation
)

func (v VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE) String() string {
	switch v {
	case VmmCallbackCallingStageInvalidEventEmulation:
		return "Vmm Callback Calling Stage Invalid Event Emulation"
	case VmmCallbackCallingStagePreEventEmulation:
		return "Vmm Callback Calling Stage Pre Event Emulation"
	case VmmCallbackCallingStagePostEventEmulation:
		return "Vmm Callback Calling Stage Post Event Emulation"
	case VmmCallbackCallingStageAllEventEmulation:
		return "Vmm Callback Calling Stage All Event Emulation"
	default:
		return fmt.Sprintf("VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE(0x%X)", uint32(v))
	}
}

// Source: unknown.h:121 -> _DEBUGGER_THREAD_PROCESS_TRACING
type DEBUGGER_THREAD_PROCESS_TRACING uint32

const (
	DebuggerThreadProcessTracingInterceptClockInterruptsForThreadChange DEBUGGER_THREAD_PROCESS_TRACING = iota
	DebuggerThreadProcessTracingInterceptClockInterruptsForProcessChange
	DebuggerThreadProcessTracingInterceptClockDebugRegisterInterception
	DebuggerThreadProcessTracingInterceptClockWaitingForMovCr3VmExits
)

func (d DEBUGGER_THREAD_PROCESS_TRACING) String() string {
	switch d {
	case DebuggerThreadProcessTracingInterceptClockInterruptsForThreadChange:
		return "Debugger Thread Process Tracing Intercept Clock Interrupts For Thread Change"
	case DebuggerThreadProcessTracingInterceptClockInterruptsForProcessChange:
		return "Debugger Thread Process Tracing Intercept Clock Interrupts For Process Change"
	case DebuggerThreadProcessTracingInterceptClockDebugRegisterInterception:
		return "Debugger Thread Process Tracing Intercept Clock Debug Register Interception"
	case DebuggerThreadProcessTracingInterceptClockWaitingForMovCr3VmExits:
		return "Debugger Thread Process Tracing Intercept Clock Waiting For Mov Cr 3 Vm Exits"
	default:
		return fmt.Sprintf("DEBUGGER_THREAD_PROCESS_TRACING(0x%X)", uint32(d))
	}
}

// Source: unknown.h:270 -> _NOTIFY_TYPE
type NOTIFY_TYPE uint32

const (
	IrpBased NOTIFY_TYPE = iota
	EventBased
)

func (n NOTIFY_TYPE) String() string {
	switch n {
	case IrpBased:
		return "Irp Based"
	case EventBased:
		return "Event Based"
	default:
		return fmt.Sprintf("NOTIFY_TYPE(0x%X)", uint32(n))
	}
}

// Source: unknown.h:343 -> _DEBUGGER_HOOK_MEMORY_TYPE
type DEBUGGER_HOOK_MEMORY_TYPE uint32

const (
	DebuggerMemoryHookVirtualAddress DEBUGGER_HOOK_MEMORY_TYPE = iota
	DebuggerMemoryHookPhysicalAddress
)

func (d DEBUGGER_HOOK_MEMORY_TYPE) String() string {
	switch d {
	case DebuggerMemoryHookVirtualAddress:
		return "Debugger Memory Hook Virtual Address"
	case DebuggerMemoryHookPhysicalAddress:
		return "Debugger Memory Hook Physical Address"
	default:
		return fmt.Sprintf("DEBUGGER_HOOK_MEMORY_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:22 -> _EXCEPTION_VECTORS
type EXCEPTION_VECTORS uint32

const (
	ExceptionVectorDivideError EXCEPTION_VECTORS = iota
	ExceptionVectorDebugBreakpoint
	ExceptionVectorNmi
	ExceptionVectorBreakpoint
	ExceptionVectorOverflow
	ExceptionVectorBoundRangeExceeded
	ExceptionVectorUndefinedOpcode
	ExceptionVectorNoMathCoprocessor
	ExceptionVectorDoubleFault
	ExceptionVectorReserved0
	ExceptionVectorInvalidTaskSegmentSelector
	ExceptionVectorSegmentNotPresent
	ExceptionVectorStackSegmentFault
	ExceptionVectorGeneralProtectionFault
	ExceptionVectorPageFault
	ExceptionVectorReserved1
	ExceptionVectorMathFault
	ExceptionVectorAlignmentCheck
	ExceptionVectorMachineCheck
	ExceptionVectorSimdFloatingPointNumericError
	ExceptionVectorVirtualException
	ExceptionVectorReserved2
	ExceptionVectorReserved3
	ExceptionVectorReserved4
	ExceptionVectorReserved5
	ExceptionVectorReserved6
	ExceptionVectorReserved7
	ExceptionVectorReserved8
	ExceptionVectorReserved9
	ExceptionVectorReserved10
	ExceptionVectorReserved11
	ExceptionVectorReserved12
	ApcInterrupt   EXCEPTION_VECTORS = 31
	DpcInterrupt   EXCEPTION_VECTORS = 47
	ClockInterrupt EXCEPTION_VECTORS = 209
	IpiInterrupt   EXCEPTION_VECTORS = 225
	PmiInterrupt   EXCEPTION_VECTORS = 254
)

func (e EXCEPTION_VECTORS) String() string {
	switch e {
	case ExceptionVectorDivideError:
		return "Exception Vector Divide Error"
	case ExceptionVectorDebugBreakpoint:
		return "Exception Vector Debug Breakpoint"
	case ExceptionVectorNmi:
		return "Exception Vector Nmi"
	case ExceptionVectorBreakpoint:
		return "Exception Vector Breakpoint"
	case ExceptionVectorOverflow:
		return "Exception Vector Overflow"
	case ExceptionVectorBoundRangeExceeded:
		return "Exception Vector Bound Range Exceeded"
	case ExceptionVectorUndefinedOpcode:
		return "Exception Vector Undefined Opcode"
	case ExceptionVectorNoMathCoprocessor:
		return "Exception Vector No Math Coprocessor"
	case ExceptionVectorDoubleFault:
		return "Exception Vector Double Fault"
	case ExceptionVectorReserved0:
		return "Exception Vector Reserved 0"
	case ExceptionVectorInvalidTaskSegmentSelector:
		return "Exception Vector Invalid Task Segment Selector"
	case ExceptionVectorSegmentNotPresent:
		return "Exception Vector Segment Not Present"
	case ExceptionVectorStackSegmentFault:
		return "Exception Vector Stack Segment Fault"
	case ExceptionVectorGeneralProtectionFault:
		return "Exception Vector General Protection Fault"
	case ExceptionVectorPageFault:
		return "Exception Vector Page Fault"
	case ExceptionVectorReserved1:
		return "Exception Vector Reserved 1"
	case ExceptionVectorMathFault:
		return "Exception Vector Math Fault"
	case ExceptionVectorAlignmentCheck:
		return "Exception Vector Alignment Check"
	case ExceptionVectorMachineCheck:
		return "Exception Vector Machine Check"
	case ExceptionVectorSimdFloatingPointNumericError:
		return "Exception Vector Simd Floating Point Numeric Error"
	case ExceptionVectorVirtualException:
		return "Exception Vector Virtual Exception"
	case ExceptionVectorReserved2:
		return "Exception Vector Reserved 2"
	case ExceptionVectorReserved3:
		return "Exception Vector Reserved 3"
	case ExceptionVectorReserved4:
		return "Exception Vector Reserved 4"
	case ExceptionVectorReserved5:
		return "Exception Vector Reserved 5"
	case ExceptionVectorReserved6:
		return "Exception Vector Reserved 6"
	case ExceptionVectorReserved7:
		return "Exception Vector Reserved 7"
	case ExceptionVectorReserved8:
		return "Exception Vector Reserved 8"
	case ExceptionVectorReserved9:
		return "Exception Vector Reserved 9"
	case ExceptionVectorReserved10:
		return "Exception Vector Reserved 10"
	case ExceptionVectorReserved11:
		return "Exception Vector Reserved 11"
	case ExceptionVectorReserved12:
		return "Exception Vector Reserved 12"
	case DpcInterrupt:
		return "Dpc Interrupt"
	case ClockInterrupt:
		return "Clock Interrupt"
	case IpiInterrupt:
		return "Ipi Interrupt"
	case PmiInterrupt:
		return "Pmi Interrupt"
	default:
		return fmt.Sprintf("EXCEPTION_VECTORS(0x%X)", uint32(e))
	}
}

// Source: unknown.h:76 -> _VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE
type VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE uint32

const (
	VmmCallbackTriggeringEventStatusSuccessfulNoInitialized VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE = 0
	VmmCallbackTriggeringEventStatusSuccessful              VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE = 0
	VmmCallbackTriggeringEventStatusSuccessfulIgnoreEvent   VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE = 1
	VmmCallbackTriggeringEventStatusDebuggerNotEnabled      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE = 2
	VmmCallbackTriggeringEventStatusInvalidEventType        VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE = 3
)

func (v VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE) String() string {
	switch v {
	case VmmCallbackTriggeringEventStatusSuccessfulNoInitialized:
		return "Vmm Callback Triggering Event Status Successful No Initialized"
	case VmmCallbackTriggeringEventStatusSuccessfulIgnoreEvent:
		return "Vmm Callback Triggering Event Status Successful Ignore Event"
	case VmmCallbackTriggeringEventStatusDebuggerNotEnabled:
		return "Vmm Callback Triggering Event Status Debugger Not Enabled"
	case VmmCallbackTriggeringEventStatusInvalidEventType:
		return "Vmm Callback Triggering Event Status Invalid Event Type"
	default:
		return fmt.Sprintf("VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE(0x%X)", uint32(v))
	}
}

// Source: unknown.h:94 -> _VMM_EVENT_TYPE_ENUM
type VMM_EVENT_TYPE_ENUM uint32

const (
	HiddenHookReadAndWriteAndExecute VMM_EVENT_TYPE_ENUM = iota
	HiddenHookReadAndWrite
	HiddenHookReadAndExecute
	HiddenHookWriteAndExecute
	HiddenHookRead
	HiddenHookWrite
	HiddenHookExecute
	HiddenHookExecDetours
	HiddenHookExecCc
	SyscallHookEferSyscall
	SyscallHookEferSysret
	CpuidInstructionExecution
	RdmsrInstructionExecution
	WrmsrInstructionExecution
	InInstructionExecution
	OutInstructionExecution
	ExceptionOccurred
	ExternalInterruptOccurred
	DebugRegistersAccessed
	TscInstructionExecution
	PmcInstructionExecution
	VmcallInstructionExecution
	ControlRegisterModified
	ControlRegisterRead
	ControlRegister3Modified
	TrapExecutionModeChanged
	TrapExecutionInstructionTrace
	XsetbvInstructionExecution
)

func (v VMM_EVENT_TYPE_ENUM) String() string {
	switch v {
	case HiddenHookReadAndWriteAndExecute:
		return "Hidden Hook Read And Write And Execute"
	case HiddenHookReadAndWrite:
		return "Hidden Hook Read And Write"
	case HiddenHookReadAndExecute:
		return "Hidden Hook Read And Execute"
	case HiddenHookWriteAndExecute:
		return "Hidden Hook Write And Execute"
	case HiddenHookRead:
		return "Hidden Hook Read"
	case HiddenHookWrite:
		return "Hidden Hook Write"
	case HiddenHookExecute:
		return "Hidden Hook Execute"
	case HiddenHookExecDetours:
		return "Hidden Hook Exec Detours"
	case HiddenHookExecCc:
		return "Hidden Hook Exec Cc"
	case SyscallHookEferSyscall:
		return "Syscall Hook Efer Syscall"
	case SyscallHookEferSysret:
		return "Syscall Hook Efer Sysret"
	case CpuidInstructionExecution:
		return "Cpuid Instruction Execution"
	case RdmsrInstructionExecution:
		return "Rdmsr Instruction Execution"
	case WrmsrInstructionExecution:
		return "Wrmsr Instruction Execution"
	case InInstructionExecution:
		return "In Instruction Execution"
	case OutInstructionExecution:
		return "Out Instruction Execution"
	case ExceptionOccurred:
		return "Exception Occurred"
	case ExternalInterruptOccurred:
		return "External Interrupt Occurred"
	case DebugRegistersAccessed:
		return "Debug Registers Accessed"
	case TscInstructionExecution:
		return "Tsc Instruction Execution"
	case PmcInstructionExecution:
		return "Pmc Instruction Execution"
	case VmcallInstructionExecution:
		return "Vmcall Instruction Execution"
	case ControlRegisterModified:
		return "Control Register Modified"
	case ControlRegisterRead:
		return "Control Register Read"
	case ControlRegister3Modified:
		return "Control Register 3 Modified"
	case TrapExecutionModeChanged:
		return "Trap Execution Mode Changed"
	case TrapExecutionInstructionTrace:
		return "Trap Execution Instruction Trace"
	case XsetbvInstructionExecution:
		return "Xsetbv Instruction Execution"
	default:
		return fmt.Sprintf("VMM_EVENT_TYPE_ENUM(0x%X)", uint32(v))
	}
}

// Source: unknown.h:183 -> _DEBUGGER_EVENT_ACTION_TYPE_ENUM
type DEBUGGER_EVENT_ACTION_TYPE_ENUM uint32

const (
	BreakToDebugger DEBUGGER_EVENT_ACTION_TYPE_ENUM = iota
	RunScript
	RunCustomCode
)

func (d DEBUGGER_EVENT_ACTION_TYPE_ENUM) String() string {
	switch d {
	case BreakToDebugger:
		return "Break To Debugger"
	case RunScript:
		return "Run Script"
	case RunCustomCode:
		return "Run Custom Code"
	default:
		return fmt.Sprintf("DEBUGGER_EVENT_ACTION_TYPE_ENUM(0x%X)", uint32(d))
	}
}

// Source: unknown.h:195 -> _DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE
type DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE uint32

const (
	DebuggerEventSyscallSysretSafeAccessMemory DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE = iota
	DebuggerEventSyscallSysretHandleAllUd
)

func (d DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE) String() string {
	switch d {
	case DebuggerEventSyscallSysretSafeAccessMemory:
		return "Debugger Event Syscall Sysret Safe Access Memory"
	case DebuggerEventSyscallSysretHandleAllUd:
		return "Debugger Event Syscall Sysret Handle All Ud"
	default:
		return fmt.Sprintf("DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:208 -> _DEBUGGER_EVENT_MODE_TYPE
type DEBUGGER_EVENT_MODE_TYPE uint32

const (
	DebuggerEventModeTypeUserModeAndKernelMode DEBUGGER_EVENT_MODE_TYPE = 1
	DebuggerEventModeTypeUserMode              DEBUGGER_EVENT_MODE_TYPE = 3
	DebuggerEventModeTypeKernelMode            DEBUGGER_EVENT_MODE_TYPE = 0
	DebuggerEventModeTypeInvalid               DEBUGGER_EVENT_MODE_TYPE = 0xffffffff
)

func (d DEBUGGER_EVENT_MODE_TYPE) String() string {
	switch d {
	case DebuggerEventModeTypeUserModeAndKernelMode:
		return "Debugger Event Mode Type User Mode And Kernel Mode"
	case DebuggerEventModeTypeUserMode:
		return "Debugger Event Mode Type User Mode"
	case DebuggerEventModeTypeKernelMode:
		return "Debugger Event Mode Type Kernel Mode"
	case DebuggerEventModeTypeInvalid:
		return "Debugger Event Mode Type Invalid"
	default:
		return fmt.Sprintf("DEBUGGER_EVENT_MODE_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:221 -> _DEBUGGER_EVENT_TRACE_TYPE
type DEBUGGER_EVENT_TRACE_TYPE uint32

const (
	DebuggerEventTraceTypeInvalid DEBUGGER_EVENT_TRACE_TYPE = iota
	DebuggerEventTraceTypeStepIn
	DebuggerEventTraceTypeStepOut
	DebuggerEventTraceTypeInstrumentationStepIn
)

func (d DEBUGGER_EVENT_TRACE_TYPE) String() string {
	switch d {
	case DebuggerEventTraceTypeInvalid:
		return "Debugger Event Trace Type Invalid"
	case DebuggerEventTraceTypeStepIn:
		return "Debugger Event Trace Type Step In"
	case DebuggerEventTraceTypeStepOut:
		return "Debugger Event Trace Type Step Out"
	case DebuggerEventTraceTypeInstrumentationStepIn:
		return "Debugger Event Trace Type Instrumentation Step In"
	default:
		return fmt.Sprintf("DEBUGGER_EVENT_TRACE_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:234 -> _DEBUGGER_MODIFY_EVENTS_TYPE
type DEBUGGER_MODIFY_EVENTS_TYPE uint32

const (
	DebuggerModifyEventsQueryState DEBUGGER_MODIFY_EVENTS_TYPE = iota
	DebuggerModifyEventsEnable
	DebuggerModifyEventsDisable
	DebuggerModifyEventsClear
)

func (d DEBUGGER_MODIFY_EVENTS_TYPE) String() string {
	switch d {
	case DebuggerModifyEventsQueryState:
		return "Debugger Modify Events Query State"
	case DebuggerModifyEventsEnable:
		return "Debugger Modify Events Enable"
	case DebuggerModifyEventsDisable:
		return "Debugger Modify Events Disable"
	case DebuggerModifyEventsClear:
		return "Debugger Modify Events Clear"
	default:
		return fmt.Sprintf("DEBUGGER_MODIFY_EVENTS_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:294 -> _PROTECTED_HV_RESOURCES_PASSING_OVERS
type PROTECTED_HV_RESOURCES_PASSING_OVERS uint32

const (
	PassingOverNone PROTECTED_HV_RESOURCES_PASSING_OVERS = iota
	PassingOverUdExceptionsForSyscallSysretHook
	PassingOverExceptionEvents
	PassingOverInterruptEvents
	PassingOverTscEvents
	PassingOverMovToHwDebugRegsEvents
	PassingOverMovToControlRegsEvents
)

func (p PROTECTED_HV_RESOURCES_PASSING_OVERS) String() string {
	switch p {
	case PassingOverNone:
		return "Passing Over None"
	case PassingOverUdExceptionsForSyscallSysretHook:
		return "Passing Over Ud Exceptions For Syscall Sysret Hook"
	case PassingOverExceptionEvents:
		return "Passing Over Exception Events"
	case PassingOverInterruptEvents:
		return "Passing Over Interrupt Events"
	case PassingOverTscEvents:
		return "Passing Over Tsc Events"
	case PassingOverMovToHwDebugRegsEvents:
		return "Passing Over Mov To Hw Debug Regs Events"
	case PassingOverMovToControlRegsEvents:
		return "Passing Over Mov To Control Regs Events"
	default:
		return fmt.Sprintf("PROTECTED_HV_RESOURCES_PASSING_OVERS(0x%X)", uint32(p))
	}
}

// Source: unknown.h:329 -> _PROTECTED_HV_RESOURCES_TYPE
type PROTECTED_HV_RESOURCES_TYPE uint32

const (
	ProtectedHvResourcesExceptionBitmap PROTECTED_HV_RESOURCES_TYPE = iota
	ProtectedHvResourcesExternalInterruptExiting
	ProtectedHvResourcesRdtscRdtscpExiting
	ProtectedHvResourcesMovToDebugRegisterExiting
	ProtectedHvResourcesMovControlRegisterExiting
	ProtectedHvResourcesMovToCr3Exiting
	ProtectedHvResourcesSaveAndLoadDebugControls
)

func (p PROTECTED_HV_RESOURCES_TYPE) String() string {
	switch p {
	case ProtectedHvResourcesExceptionBitmap:
		return "Protected Hv Resources Exception Bitmap"
	case ProtectedHvResourcesExternalInterruptExiting:
		return "Protected Hv Resources External Interrupt Exiting"
	case ProtectedHvResourcesRdtscRdtscpExiting:
		return "Protected Hv Resources Rdtsc Rdtscp Exiting"
	case ProtectedHvResourcesMovToDebugRegisterExiting:
		return "Protected Hv Resources Mov To Debug Register Exiting"
	case ProtectedHvResourcesMovControlRegisterExiting:
		return "Protected Hv Resources Mov Control Register Exiting"
	case ProtectedHvResourcesMovToCr3Exiting:
		return "Protected Hv Resources Mov To Cr 3 Exiting"
	case ProtectedHvResourcesSaveAndLoadDebugControls:
		return "Protected Hv Resources Save And Load Debug Controls"
	default:
		return fmt.Sprintf("PROTECTED_HV_RESOURCES_TYPE(0x%X)", uint32(p))
	}
}

// Source: unknown.h:88 -> _REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE
type REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE uint32

const (
	ReversingMachineReconstructMemoryModeUnknown REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE = iota
	ReversingMachineReconstructMemoryModeUserMode
	ReversingMachineReconstructMemoryModeKernelMode
)

func (r REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE) String() string {
	switch r {
	case ReversingMachineReconstructMemoryModeUnknown:
		return "Reversing Machine Reconstruct Memory Mode Unknown"
	case ReversingMachineReconstructMemoryModeUserMode:
		return "Reversing Machine Reconstruct Memory Mode User Mode"
	case ReversingMachineReconstructMemoryModeKernelMode:
		return "Reversing Machine Reconstruct Memory Mode Kernel Mode"
	default:
		return fmt.Sprintf("REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE(0x%X)", uint32(r))
	}
}

// Source: unknown.h:99 -> _REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE
type REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE uint32

const (
	ReversingMachineReconstructMemoryTypeUnknown REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE = iota
	ReversingMachineReconstructMemoryTypeReconstruct
	ReversingMachineReconstructMemoryTypePattern
)

func (r REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE) String() string {
	switch r {
	case ReversingMachineReconstructMemoryTypeUnknown:
		return "Reversing Machine Reconstruct Memory Type Unknown"
	case ReversingMachineReconstructMemoryTypeReconstruct:
		return "Reversing Machine Reconstruct Memory Type Reconstruct"
	case ReversingMachineReconstructMemoryTypePattern:
		return "Reversing Machine Reconstruct Memory Type Pattern"
	default:
		return fmt.Sprintf("REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE(0x%X)", uint32(r))
	}
}

// Source: unknown.h:150 -> _DEBUGGER_PREALLOC_COMMAND_TYPE
type DEBUGGER_PREALLOC_COMMAND_TYPE uint32

const (
	DebuggerPreallocCommandTypeThreadInterception DEBUGGER_PREALLOC_COMMAND_TYPE = iota
	DebuggerPreallocCommandTypeMonitor
	DebuggerPreallocCommandTypeEpthook
	DebuggerPreallocCommandTypeEpthook2
	DebuggerPreallocCommandTypeRegularEvent
	DebuggerPreallocCommandTypeBigEvent
	DebuggerPreallocCommandTypeRegularSafeBuffer
	DebuggerPreallocCommandTypeBigSafeBuffer
)

func (d DEBUGGER_PREALLOC_COMMAND_TYPE) String() string {
	switch d {
	case DebuggerPreallocCommandTypeThreadInterception:
		return "Debugger Prealloc Command Type Thread Interception"
	case DebuggerPreallocCommandTypeMonitor:
		return "Debugger Prealloc Command Type Monitor"
	case DebuggerPreallocCommandTypeEpthook:
		return "Debugger Prealloc Command Type Epthook"
	case DebuggerPreallocCommandTypeEpthook2:
		return "Debugger Prealloc Command Type Epthook 2"
	case DebuggerPreallocCommandTypeRegularEvent:
		return "Debugger Prealloc Command Type Regular Event"
	case DebuggerPreallocCommandTypeBigEvent:
		return "Debugger Prealloc Command Type Big Event"
	case DebuggerPreallocCommandTypeRegularSafeBuffer:
		return "Debugger Prealloc Command Type Regular Safe Buffer"
	case DebuggerPreallocCommandTypeBigSafeBuffer:
		return "Debugger Prealloc Command Type Big Safe Buffer"
	default:
		return fmt.Sprintf("DEBUGGER_PREALLOC_COMMAND_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:184 -> _DEBUGGER_PREACTIVATE_COMMAND_TYPE
type DEBUGGER_PREACTIVATE_COMMAND_TYPE uint32

const (
	DebuggerPreactivateCommandTypeMode DEBUGGER_PREACTIVATE_COMMAND_TYPE = iota
)

func (d DEBUGGER_PREACTIVATE_COMMAND_TYPE) String() string {
	switch d {
	case DebuggerPreactivateCommandTypeMode:
		return "Debugger Preactivate Command Type Mode"
	default:
		return fmt.Sprintf("DEBUGGER_PREACTIVATE_COMMAND_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:212 -> _DEBUGGER_READ_READING_TYPE
type DEBUGGER_READ_READING_TYPE uint32

const (
	ReadFromKernel DEBUGGER_READ_READING_TYPE = iota
	ReadFromVmxRoot
)

func (d DEBUGGER_READ_READING_TYPE) String() string {
	switch d {
	case ReadFromKernel:
		return "Read From Kernel"
	case ReadFromVmxRoot:
		return "Read From Vmx Root"
	default:
		return fmt.Sprintf("DEBUGGER_READ_READING_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:222 -> _DEBUGGER_READ_MEMORY_TYPE
type DEBUGGER_READ_MEMORY_TYPE uint32

const (
	DebuggerReadPhysicalAddress DEBUGGER_READ_MEMORY_TYPE = iota
	DebuggerReadVirtualAddress
)

func (d DEBUGGER_READ_MEMORY_TYPE) String() string {
	switch d {
	case DebuggerReadPhysicalAddress:
		return "Debugger Read Physical Address"
	case DebuggerReadVirtualAddress:
		return "Debugger Read Virtual Address"
	default:
		return fmt.Sprintf("DEBUGGER_READ_MEMORY_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:232 -> _DEBUGGER_READ_MEMORY_ADDRESS_MODE
type DEBUGGER_READ_MEMORY_ADDRESS_MODE uint32

const (
	DebuggerReadAddressMode32Bit DEBUGGER_READ_MEMORY_ADDRESS_MODE = iota
	DebuggerReadAddressMode64Bit
)

func (d DEBUGGER_READ_MEMORY_ADDRESS_MODE) String() string {
	switch d {
	case DebuggerReadAddressMode32Bit:
		return "Debugger Read Address Mode 32 Bit"
	case DebuggerReadAddressMode64Bit:
		return "Debugger Read Address Mode 64 Bit"
	default:
		return fmt.Sprintf("DEBUGGER_READ_MEMORY_ADDRESS_MODE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:244 -> _DEBUGGER_SHOW_MEMORY_STYLE
type DEBUGGER_SHOW_MEMORY_STYLE uint32

const (
	DebuggerShowCommandDt DEBUGGER_SHOW_MEMORY_STYLE = 1 + iota
	DebuggerShowCommandDisassemble64
	DebuggerShowCommandDisassemble32
	DebuggerShowCommandDb
	DebuggerShowCommandDc
	DebuggerShowCommandDq
	DebuggerShowCommandDd
	DebuggerShowCommandDump
)

func (d DEBUGGER_SHOW_MEMORY_STYLE) String() string {
	switch d {
	case DebuggerShowCommandDt:
		return "Debugger Show Command Dt"
	case DebuggerShowCommandDisassemble64:
		return "Debugger Show Command Disassemble 64"
	case DebuggerShowCommandDisassemble32:
		return "Debugger Show Command Disassemble 32"
	case DebuggerShowCommandDb:
		return "Debugger Show Command Db"
	case DebuggerShowCommandDc:
		return "Debugger Show Command Dc"
	case DebuggerShowCommandDq:
		return "Debugger Show Command Dq"
	case DebuggerShowCommandDd:
		return "Debugger Show Command Dd"
	case DebuggerShowCommandDump:
		return "Debugger Show Command Dump"
	default:
		return fmt.Sprintf("DEBUGGER_SHOW_MEMORY_STYLE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:304 -> _DEBUGGER_TEST_QUERY_STATE
type DEBUGGER_TEST_QUERY_STATE uint32

const (
	TestQueryHaltingCoreStatus DEBUGGER_TEST_QUERY_STATE = 1 + iota
	TestQueryPreallocatedPoolState
	TestQueryTrapState
	TestBreakpointTurnOffBps
	TestBreakpointTurnOnBps
	TestBreakpointTurnOffBpsAndEventsForCommandsInRemoteComputer
	TestBreakpointTurnOnBpsAndEventsForCommandsInRemoteComputer
	TestSettingTargetTasksOnHaltedCoresSynchronous
	TestSettingTargetTasksOnHaltedCoresAsynchronous
	TestSettingTargetTasksOnTargetHaltedCores
	TestBreakpointTurnOffDbs
	TestBreakpointTurnOnDbs
)

func (d DEBUGGER_TEST_QUERY_STATE) String() string {
	switch d {
	case TestQueryHaltingCoreStatus:
		return "Test Query Halting Core Status"
	case TestQueryPreallocatedPoolState:
		return "Test Query Preallocated Pool State"
	case TestQueryTrapState:
		return "Test Query Trap State"
	case TestBreakpointTurnOffBps:
		return "Test Breakpoint Turn Off Bps"
	case TestBreakpointTurnOnBps:
		return "Test Breakpoint Turn On Bps"
	case TestBreakpointTurnOffBpsAndEventsForCommandsInRemoteComputer:
		return "Test Breakpoint Turn Off Bps And Events For Commands In Remote Computer"
	case TestBreakpointTurnOnBpsAndEventsForCommandsInRemoteComputer:
		return "Test Breakpoint Turn On Bps And Events For Commands In Remote Computer"
	case TestSettingTargetTasksOnHaltedCoresSynchronous:
		return "Test Setting Target Tasks On Halted Cores Synchronous"
	case TestSettingTargetTasksOnHaltedCoresAsynchronous:
		return "Test Setting Target Tasks On Halted Cores Asynchronous"
	case TestSettingTargetTasksOnTargetHaltedCores:
		return "Test Setting Target Tasks On Target Halted Cores"
	case TestBreakpointTurnOffDbs:
		return "Test Breakpoint Turn Off Dbs"
	case TestBreakpointTurnOnDbs:
		return "Test Breakpoint Turn On Dbs"
	default:
		return fmt.Sprintf("DEBUGGER_TEST_QUERY_STATE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:417 -> _DEBUGGER_MSR_ACTION_TYPE
type DEBUGGER_MSR_ACTION_TYPE uint32

const (
	DebuggerMsrRead DEBUGGER_MSR_ACTION_TYPE = iota
	DebuggerMsrWrite
)

func (d DEBUGGER_MSR_ACTION_TYPE) String() string {
	switch d {
	case DebuggerMsrRead:
		return "Debugger Msr Read"
	case DebuggerMsrWrite:
		return "Debugger Msr Write"
	default:
		return fmt.Sprintf("DEBUGGER_MSR_ACTION_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:447 -> _DEBUGGER_EDIT_MEMORY_TYPE
type DEBUGGER_EDIT_MEMORY_TYPE uint32

const (
	EditVirtualMemory DEBUGGER_EDIT_MEMORY_TYPE = iota
	EditPhysicalMemory
)

func (d DEBUGGER_EDIT_MEMORY_TYPE) String() string {
	switch d {
	case EditVirtualMemory:
		return "Edit Virtual Memory"
	case EditPhysicalMemory:
		return "Edit Physical Memory"
	default:
		return fmt.Sprintf("DEBUGGER_EDIT_MEMORY_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:457 -> _DEBUGGER_EDIT_MEMORY_BYTE_SIZE
type DEBUGGER_EDIT_MEMORY_BYTE_SIZE uint32

const (
	EditByte DEBUGGER_EDIT_MEMORY_BYTE_SIZE = iota
	EditDword
	EditQword
)

func (d DEBUGGER_EDIT_MEMORY_BYTE_SIZE) String() string {
	switch d {
	case EditByte:
		return "Edit Byte"
	case EditDword:
		return "Edit Dword"
	case EditQword:
		return "Edit Qword"
	default:
		return fmt.Sprintf("DEBUGGER_EDIT_MEMORY_BYTE_SIZE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:488 -> _DEBUGGER_SEARCH_MEMORY_TYPE
type DEBUGGER_SEARCH_MEMORY_TYPE uint32

const (
	SearchPhysicalMemory DEBUGGER_SEARCH_MEMORY_TYPE = iota
	SearchVirtualMemory
	SearchPhysicalFromVirtualMemory
)

func (d DEBUGGER_SEARCH_MEMORY_TYPE) String() string {
	switch d {
	case SearchPhysicalMemory:
		return "Search Physical Memory"
	case SearchVirtualMemory:
		return "Search Virtual Memory"
	case SearchPhysicalFromVirtualMemory:
		return "Search Physical From Virtual Memory"
	default:
		return fmt.Sprintf("DEBUGGER_SEARCH_MEMORY_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:500 -> _DEBUGGER_SEARCH_MEMORY_BYTE_SIZE
type DEBUGGER_SEARCH_MEMORY_BYTE_SIZE uint32

const (
	SearchByte DEBUGGER_SEARCH_MEMORY_BYTE_SIZE = iota
	SearchDword
	SearchQword
)

func (d DEBUGGER_SEARCH_MEMORY_BYTE_SIZE) String() string {
	switch d {
	case SearchByte:
		return "Search Byte"
	case SearchDword:
		return "Search Dword"
	case SearchQword:
		return "Search Qword"
	default:
		return fmt.Sprintf("DEBUGGER_SEARCH_MEMORY_BYTE_SIZE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:627 -> _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE
type DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE uint32

const (
	DebuggerAttachDetachUserModeProcessActionAttach DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE = iota
	DebuggerAttachDetachUserModeProcessActionDetach
	DebuggerAttachDetachUserModeProcessActionRemoveHooks
	DebuggerAttachDetachUserModeProcessActionKillProcess
	DebuggerAttachDetachUserModeProcessActionContinueProcess
	DebuggerAttachDetachUserModeProcessActionPauseProcess
	DebuggerAttachDetachUserModeProcessActionSwitchByProcessOrThread
	DebuggerAttachDetachUserModeProcessActionQueryCountOfActiveDebuggingThreads
)

func (d DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE) String() string {
	switch d {
	case DebuggerAttachDetachUserModeProcessActionAttach:
		return "Debugger Attach Detach User Mode Process Action Attach"
	case DebuggerAttachDetachUserModeProcessActionDetach:
		return "Debugger Attach Detach User Mode Process Action Detach"
	case DebuggerAttachDetachUserModeProcessActionRemoveHooks:
		return "Debugger Attach Detach User Mode Process Action Remove Hooks"
	case DebuggerAttachDetachUserModeProcessActionKillProcess:
		return "Debugger Attach Detach User Mode Process Action Kill Process"
	case DebuggerAttachDetachUserModeProcessActionContinueProcess:
		return "Debugger Attach Detach User Mode Process Action Continue Process"
	case DebuggerAttachDetachUserModeProcessActionPauseProcess:
		return "Debugger Attach Detach User Mode Process Action Pause Process"
	case DebuggerAttachDetachUserModeProcessActionSwitchByProcessOrThread:
		return "Debugger Attach Detach User Mode Process Action Switch By Process Or Thread"
	case DebuggerAttachDetachUserModeProcessActionQueryCountOfActiveDebuggingThreads:
		return "Debugger Attach Detach User Mode Process Action Query Count Of Active Debugging Threads"
	default:
		return fmt.Sprintf("DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:672 -> _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES
type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES uint32

const (
	DebuggerQueryActiveProcessesOrThreadsQueryProcessCount DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES = 1 + iota
	DebuggerQueryActiveProcessesOrThreadsQueryThreadCount
	DebuggerQueryActiveProcessesOrThreadsQueryProcessList
	DebuggerQueryActiveProcessesOrThreadsQueryThreadList
	DebuggerQueryActiveProcessesOrThreadsQueryCurrentProcess
	DebuggerQueryActiveProcessesOrThreadsQueryCurrentThread
)

func (d DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES) String() string {
	switch d {
	case DebuggerQueryActiveProcessesOrThreadsQueryProcessCount:
		return "Debugger Query Active Processes Or Threads Query Process Count"
	case DebuggerQueryActiveProcessesOrThreadsQueryThreadCount:
		return "Debugger Query Active Processes Or Threads Query Thread Count"
	case DebuggerQueryActiveProcessesOrThreadsQueryProcessList:
		return "Debugger Query Active Processes Or Threads Query Process List"
	case DebuggerQueryActiveProcessesOrThreadsQueryThreadList:
		return "Debugger Query Active Processes Or Threads Query Thread List"
	case DebuggerQueryActiveProcessesOrThreadsQueryCurrentProcess:
		return "Debugger Query Active Processes Or Threads Query Current Process"
	case DebuggerQueryActiveProcessesOrThreadsQueryCurrentThread:
		return "Debugger Query Active Processes Or Threads Query Current Thread"
	default:
		return fmt.Sprintf("DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES(0x%X)", uint32(d))
	}
}

// Source: unknown.h:687 -> _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS
type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS uint32

const (
	DebuggerQueryActiveProcessesOrThreadsActionShowInstantly DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS = 1 + iota
	DebuggerQueryActiveProcessesOrThreadsActionQueryCount
	DebuggerQueryActiveProcessesOrThreadsActionQuerySaveDetails
)

func (d DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS) String() string {
	switch d {
	case DebuggerQueryActiveProcessesOrThreadsActionShowInstantly:
		return "Debugger Query Active Processes Or Threads Action Show Instantly"
	case DebuggerQueryActiveProcessesOrThreadsActionQueryCount:
		return "Debugger Query Active Processes Or Threads Action Query Count"
	case DebuggerQueryActiveProcessesOrThreadsActionQuerySaveDetails:
		return "Debugger Query Active Processes Or Threads Action Query Save Details"
	default:
		return fmt.Sprintf("DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS(0x%X)", uint32(d))
	}
}

// Source: unknown.h:793 -> _DEBUGGER_CALLSTACK_DISPLAY_METHOD
type DEBUGGER_CALLSTACK_DISPLAY_METHOD uint32

const (
	DebuggerCallstackDisplayMethodWithoutParams DEBUGGER_CALLSTACK_DISPLAY_METHOD = iota
	DebuggerCallstackDisplayMethodWithParams
)

func (d DEBUGGER_CALLSTACK_DISPLAY_METHOD) String() string {
	switch d {
	case DebuggerCallstackDisplayMethodWithoutParams:
		return "Debugger Callstack Display Method Without Params"
	case DebuggerCallstackDisplayMethodWithParams:
		return "Debugger Callstack Display Method With Params"
	default:
		return fmt.Sprintf("DEBUGGER_CALLSTACK_DISPLAY_METHOD(0x%X)", uint32(d))
	}
}

// Source: unknown.h:880 -> _DEBUGGER_UD_COMMAND_ACTION_TYPE
type DEBUGGER_UD_COMMAND_ACTION_TYPE uint32

const (
	DebuggerUdCommandActionTypeNone DEBUGGER_UD_COMMAND_ACTION_TYPE = iota
	DebuggerUdCommandActionTypePause
	DebuggerUdCommandActionTypeRegularStep
	DebuggerUdCommandActionTypeReadRegisters
	DebuggerUdCommandActionTypeExecuteScriptBuffer
)

func (d DEBUGGER_UD_COMMAND_ACTION_TYPE) String() string {
	switch d {
	case DebuggerUdCommandActionTypeNone:
		return "Debugger Ud Command Action Type None"
	case DebuggerUdCommandActionTypePause:
		return "Debugger Ud Command Action Type Pause"
	case DebuggerUdCommandActionTypeRegularStep:
		return "Debugger Ud Command Action Type Regular Step"
	case DebuggerUdCommandActionTypeReadRegisters:
		return "Debugger Ud Command Action Type Read Registers"
	case DebuggerUdCommandActionTypeExecuteScriptBuffer:
		return "Debugger Ud Command Action Type Execute Script Buffer"
	default:
		return fmt.Sprintf("DEBUGGER_UD_COMMAND_ACTION_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:925 -> _DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE
type DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE uint32

const (
	DebuggeeDetailsAndSwitchProcessGetProcessDetails DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE = iota
	DebuggeeDetailsAndSwitchProcessGetProcessList
	DebuggeeDetailsAndSwitchProcessPerformSwitch
)

func (d DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE) String() string {
	switch d {
	case DebuggeeDetailsAndSwitchProcessGetProcessDetails:
		return "Debuggee Details And Switch Process Get Process Details"
	case DebuggeeDetailsAndSwitchProcessGetProcessList:
		return "Debuggee Details And Switch Process Get Process List"
	case DebuggeeDetailsAndSwitchProcessPerformSwitch:
		return "Debuggee Details And Switch Process Perform Switch"
	default:
		return fmt.Sprintf("DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:964 -> _DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE
type DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE uint32

const (
	DebuggeeDetailsAndSwitchThreadPerformSwitch DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE = iota
	DebuggeeDetailsAndSwitchThreadGetThreadDetails
	DebuggeeDetailsAndSwitchThreadGetThreadList
)

func (d DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE) String() string {
	switch d {
	case DebuggeeDetailsAndSwitchThreadPerformSwitch:
		return "Debuggee Details And Switch Thread Perform Switch"
	case DebuggeeDetailsAndSwitchThreadGetThreadDetails:
		return "Debuggee Details And Switch Thread Get Thread Details"
	case DebuggeeDetailsAndSwitchThreadGetThreadList:
		return "Debuggee Details And Switch Thread Get Thread List"
	default:
		return fmt.Sprintf("DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:1004 -> _DEBUGGER_REMOTE_STEPPING_REQUEST
type DEBUGGER_REMOTE_STEPPING_REQUEST uint32

const (
	DebuggerRemoteSteppingRequestStepIn DEBUGGER_REMOTE_STEPPING_REQUEST = iota
	DebuggerRemoteSteppingRequestInstrumentationStepIn
	DebuggerRemoteSteppingRequestInstrumentationStepInForTracking
	DebuggerRemoteSteppingRequestStepOver
	DebuggerRemoteSteppingRequestStepOverForGu
	DebuggerRemoteSteppingRequestStepOverForGuLastInstruction
)

func (d DEBUGGER_REMOTE_STEPPING_REQUEST) String() string {
	switch d {
	case DebuggerRemoteSteppingRequestStepIn:
		return "Debugger Remote Stepping Request Step In"
	case DebuggerRemoteSteppingRequestInstrumentationStepIn:
		return "Debugger Remote Stepping Request Instrumentation Step In"
	case DebuggerRemoteSteppingRequestInstrumentationStepInForTracking:
		return "Debugger Remote Stepping Request Instrumentation Step In For Tracking"
	case DebuggerRemoteSteppingRequestStepOver:
		return "Debugger Remote Stepping Request Step Over"
	case DebuggerRemoteSteppingRequestStepOverForGu:
		return "Debugger Remote Stepping Request Step Over For Gu"
	case DebuggerRemoteSteppingRequestStepOverForGuLastInstruction:
		return "Debugger Remote Stepping Request Step Over For Gu Last Instruction"
	default:
		return fmt.Sprintf("DEBUGGER_REMOTE_STEPPING_REQUEST(0x%X)", uint32(d))
	}
}

// Source: unknown.h:1045 -> _DEBUGGER_APIC_REQUEST_TYPE
type DEBUGGER_APIC_REQUEST_TYPE uint32

const (
	DebuggerApicRequestTypeReadLocalApic DEBUGGER_APIC_REQUEST_TYPE = iota
	DebuggerApicRequestTypeReadIoApic
)

func (d DEBUGGER_APIC_REQUEST_TYPE) String() string {
	switch d {
	case DebuggerApicRequestTypeReadLocalApic:
		return "Debugger Apic Request Type Read Local Apic"
	case DebuggerApicRequestTypeReadIoApic:
		return "Debugger Apic Request Type Read Io Apic"
	default:
		return fmt.Sprintf("DEBUGGER_APIC_REQUEST_TYPE(0x%X)", uint32(d))
	}
}

// Source: unknown.h:1212 -> _SMI_OPERATION_REQUEST_TYPE
type SMI_OPERATION_REQUEST_TYPE uint32

const (
	SmiOperationRequestTypeReadCount SMI_OPERATION_REQUEST_TYPE = iota
	SmiOperationRequestTypeTriggerPowerSmi
)

func (s SMI_OPERATION_REQUEST_TYPE) String() string {
	switch s {
	case SmiOperationRequestTypeReadCount:
		return "Smi Operation Request Type Read Count"
	case SmiOperationRequestTypeTriggerPowerSmi:
		return "Smi Operation Request Type Trigger Power Smi"
	default:
		return fmt.Sprintf("SMI_OPERATION_REQUEST_TYPE(0x%X)", uint32(s))
	}
}

// Source: unknown.h:1244 -> _HYPERTRACE_LBR_OPERATION_REQUEST_TYPE
type HYPERTRACE_LBR_OPERATION_REQUEST_TYPE uint32

const (
	HypertraceLbrOperationRequestTypeEnable HYPERTRACE_LBR_OPERATION_REQUEST_TYPE = iota
	HypertraceLbrOperationRequestTypeDisable
	HypertraceLbrOperationRequestTypeSave
	HypertraceLbrOperationRequestTypeDump
	HypertraceLbrOperationRequestTypeFlush
	HypertraceLbrOperationRequestTypeFilter
)

func (h HYPERTRACE_LBR_OPERATION_REQUEST_TYPE) String() string {
	switch h {
	case HypertraceLbrOperationRequestTypeEnable:
		return "Hypertrace Lbr Operation Request Type Enable"
	case HypertraceLbrOperationRequestTypeDisable:
		return "Hypertrace Lbr Operation Request Type Disable"
	case HypertraceLbrOperationRequestTypeSave:
		return "Hypertrace Lbr Operation Request Type Save"
	case HypertraceLbrOperationRequestTypeDump:
		return "Hypertrace Lbr Operation Request Type Dump"
	case HypertraceLbrOperationRequestTypeFlush:
		return "Hypertrace Lbr Operation Request Type Flush"
	case HypertraceLbrOperationRequestTypeFilter:
		return "Hypertrace Lbr Operation Request Type Filter"
	default:
		return fmt.Sprintf("HYPERTRACE_LBR_OPERATION_REQUEST_TYPE(0x%X)", uint32(h))
	}
}

// Source: unknown.h:1280 -> _HYPERTRACE_PT_OPERATION_REQUEST_TYPE
type HYPERTRACE_PT_OPERATION_REQUEST_TYPE uint32

const (
	HypertracePtOperationRequestTypeEnable HYPERTRACE_PT_OPERATION_REQUEST_TYPE = iota
	HypertracePtOperationRequestTypeDisable
	HypertracePtOperationRequestTypeSave
	HypertracePtOperationRequestTypeDump
)

func (h HYPERTRACE_PT_OPERATION_REQUEST_TYPE) String() string {
	switch h {
	case HypertracePtOperationRequestTypeEnable:
		return "Hypertrace Pt Operation Request Type Enable"
	case HypertracePtOperationRequestTypeDisable:
		return "Hypertrace Pt Operation Request Type Disable"
	case HypertracePtOperationRequestTypeSave:
		return "Hypertrace Pt Operation Request Type Save"
	case HypertracePtOperationRequestTypeDump:
		return "Hypertrace Pt Operation Request Type Dump"
	default:
		return fmt.Sprintf("HYPERTRACE_PT_OPERATION_REQUEST_TYPE(0x%X)", uint32(h))
	}
}

// Source: unknown.h:1394 -> _DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST
type DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST uint32

const (
	DebuggeeBreakpointModificationRequestListBreakpoints DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST = iota
	DebuggeeBreakpointModificationRequestEnable
	DebuggeeBreakpointModificationRequestDisable
	DebuggeeBreakpointModificationRequestClear
)

func (d DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST) String() string {
	switch d {
	case DebuggeeBreakpointModificationRequestListBreakpoints:
		return "Debuggee Breakpoint Modification Request List Breakpoints"
	case DebuggeeBreakpointModificationRequestEnable:
		return "Debuggee Breakpoint Modification Request Enable"
	case DebuggeeBreakpointModificationRequestDisable:
		return "Debuggee Breakpoint Modification Request Disable"
	case DebuggeeBreakpointModificationRequestClear:
		return "Debuggee Breakpoint Modification Request Clear"
	default:
		return fmt.Sprintf("DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST(0x%X)", uint32(d))
	}
}

// Source: unknown.h:1422 -> _DEBUGGER_CONDITIONAL_JUMP_STATUS
type DEBUGGER_CONDITIONAL_JUMP_STATUS uint32

const (
	DebuggerConditionalJumpStatusError DEBUGGER_CONDITIONAL_JUMP_STATUS = iota
	DebuggerConditionalJumpStatusNotConditionalJump
	DebuggerConditionalJumpStatusJumpIsTaken
	DebuggerConditionalJumpStatusJumpIsNotTaken
)

func (d DEBUGGER_CONDITIONAL_JUMP_STATUS) String() string {
	switch d {
	case DebuggerConditionalJumpStatusError:
		return "Debugger Conditional Jump Status Error"
	case DebuggerConditionalJumpStatusNotConditionalJump:
		return "Debugger Conditional Jump Status Not Conditional Jump"
	case DebuggerConditionalJumpStatusJumpIsTaken:
		return "Debugger Conditional Jump Status Jump Is Taken"
	case DebuggerConditionalJumpStatusJumpIsNotTaken:
		return "Debugger Conditional Jump Status Jump Is Not Taken"
	default:
		return fmt.Sprintf("DEBUGGER_CONDITIONAL_JUMP_STATUS(0x%X)", uint32(d))
	}
}

// Source: unknown.h:65 -> _HWDBG_ACTION_ENUMS
type HWDBG_ACTION_ENUMS uint32

const (
	HwdbgActionSendInstanceInfo HWDBG_ACTION_ENUMS = 1 + iota
	HwdbgActionConfigureScriptBuffer
)

func (h HWDBG_ACTION_ENUMS) String() string {
	switch h {
	case HwdbgActionSendInstanceInfo:
		return "Hwdbg Action Send Instance Info"
	case HwdbgActionConfigureScriptBuffer:
		return "Hwdbg Action Configure Script Buffer"
	default:
		return fmt.Sprintf("HWDBG_ACTION_ENUMS(0x%X)", uint32(h))
	}
}

// Source: unknown.h:77 -> _HWDBG_RESPONSE_ENUMS
type HWDBG_RESPONSE_ENUMS uint32

const (
	HwdbgResponseSuccessOrErrorMessage HWDBG_RESPONSE_ENUMS = 1 + iota
	HwdbgResponseInstanceInfo
)

func (h HWDBG_RESPONSE_ENUMS) String() string {
	switch h {
	case HwdbgResponseSuccessOrErrorMessage:
		return "Hwdbg Response Success Or Error Message"
	case HwdbgResponseInstanceInfo:
		return "Hwdbg Response Instance Info"
	default:
		return fmt.Sprintf("HWDBG_RESPONSE_ENUMS(0x%X)", uint32(h))
	}
}

// Source: unknown.h:89 -> _HWDBG_SUCCESS_OR_ERROR_ENUMS
type HWDBG_SUCCESS_OR_ERROR_ENUMS uint32

const (
	HwdbgOperationWasSuccessful HWDBG_SUCCESS_OR_ERROR_ENUMS = 2147483647
	HwdbgErrorInvalidPacket     HWDBG_SUCCESS_OR_ERROR_ENUMS = 1
)

func (h HWDBG_SUCCESS_OR_ERROR_ENUMS) String() string {
	switch h {
	case HwdbgOperationWasSuccessful:
		return "Hwdbg Operation Was Successful"
	case HwdbgErrorInvalidPacket:
		return "Hwdbg Error Invalid Packet"
	default:
		return fmt.Sprintf("HWDBG_SUCCESS_OR_ERROR_ENUMS(0x%X)", uint32(h))
	}
}

// Source: unknown.h:309 -> REGS_ENUM
type REGS_ENUM uint32

const (
	RegisterRax REGS_ENUM = iota
	RegisterEax
	RegisterAx
	RegisterAh
	RegisterAl
	RegisterRcx
	RegisterEcx
	RegisterCx
	RegisterCh
	RegisterCl
	RegisterRdx
	RegisterEdx
	RegisterDx
	RegisterDh
	RegisterDl
	RegisterRbx
	RegisterEbx
	RegisterBx
	RegisterBh
	RegisterBl
	RegisterRsp
	RegisterEsp
	RegisterSp
	RegisterSpl
	RegisterRbp
	RegisterEbp
	RegisterBp
	RegisterBpl
	RegisterRsi
	RegisterEsi
	RegisterSi
	RegisterSil
	RegisterRdi
	RegisterEdi
	RegisterDi
	RegisterDil
	RegisterR8
	RegisterR8d
	RegisterR8w
	RegisterR8h
	RegisterR8l
	RegisterR9
	RegisterR9d
	RegisterR9w
	RegisterR9h
	RegisterR9l
	RegisterR10
	RegisterR10d
	RegisterR10w
	RegisterR10h
	RegisterR10l
	RegisterR11
	RegisterR11d
	RegisterR11w
	RegisterR11h
	RegisterR11l
	RegisterR12
	RegisterR12d
	RegisterR12w
	RegisterR12h
	RegisterR12l
	RegisterR13
	RegisterR13d
	RegisterR13w
	RegisterR13h
	RegisterR13l
	RegisterR14
	RegisterR14d
	RegisterR14w
	RegisterR14h
	RegisterR14l
	RegisterR15
	RegisterR15d
	RegisterR15w
	RegisterR15h
	RegisterR15l
	RegisterDs
	RegisterEs
	RegisterFs
	RegisterGs
	RegisterCs
	RegisterSs
	RegisterRflags
	RegisterEflags
	RegisterFlags
	RegisterCf
	RegisterPf
	RegisterAf
	RegisterZf
	RegisterSf
	RegisterTf
	RegisterIf
	RegisterDf
	RegisterOf
	RegisterIopl
	RegisterNt
	RegisterRf
	RegisterVm
	RegisterAc
	RegisterVif
	RegisterVip
	RegisterId
	RegisterRip
	RegisterEip
	RegisterIp
	RegisterIdtr
	RegisterLdtr
	RegisterGdtr
	RegisterTr
	RegisterCr0
	RegisterCr2
	RegisterCr3
	RegisterCr4
	RegisterCr8
	RegisterDr0
	RegisterDr1
	RegisterDr2
	RegisterDr3
	RegisterDr6
	RegisterDr7
)

func (r REGS_ENUM) String() string {
	switch r {
	case RegisterRax:
		return "Register Rax"
	case RegisterEax:
		return "Register Eax"
	case RegisterAx:
		return "Register Ax"
	case RegisterAh:
		return "Register Ah"
	case RegisterAl:
		return "Register Al"
	case RegisterRcx:
		return "Register Rcx"
	case RegisterEcx:
		return "Register Ecx"
	case RegisterCx:
		return "Register Cx"
	case RegisterCh:
		return "Register Ch"
	case RegisterCl:
		return "Register Cl"
	case RegisterRdx:
		return "Register Rdx"
	case RegisterEdx:
		return "Register Edx"
	case RegisterDx:
		return "Register Dx"
	case RegisterDh:
		return "Register Dh"
	case RegisterDl:
		return "Register Dl"
	case RegisterRbx:
		return "Register Rbx"
	case RegisterEbx:
		return "Register Ebx"
	case RegisterBx:
		return "Register Bx"
	case RegisterBh:
		return "Register Bh"
	case RegisterBl:
		return "Register Bl"
	case RegisterRsp:
		return "Register Rsp"
	case RegisterEsp:
		return "Register Esp"
	case RegisterSp:
		return "Register Sp"
	case RegisterSpl:
		return "Register Spl"
	case RegisterRbp:
		return "Register Rbp"
	case RegisterEbp:
		return "Register Ebp"
	case RegisterBp:
		return "Register Bp"
	case RegisterBpl:
		return "Register Bpl"
	case RegisterRsi:
		return "Register Rsi"
	case RegisterEsi:
		return "Register Esi"
	case RegisterSi:
		return "Register Si"
	case RegisterSil:
		return "Register Sil"
	case RegisterRdi:
		return "Register Rdi"
	case RegisterEdi:
		return "Register Edi"
	case RegisterDi:
		return "Register Di"
	case RegisterDil:
		return "Register Dil"
	case RegisterR8:
		return "Register R8"
	case RegisterR8d:
		return "Register R8d"
	case RegisterR8w:
		return "Register R8w"
	case RegisterR8h:
		return "Register R8h"
	case RegisterR8l:
		return "Register R8l"
	case RegisterR9:
		return "Register R9"
	case RegisterR9d:
		return "Register R9d"
	case RegisterR9w:
		return "Register R9w"
	case RegisterR9h:
		return "Register R9h"
	case RegisterR9l:
		return "Register R9l"
	case RegisterR10:
		return "Register R10"
	case RegisterR10d:
		return "Register R10d"
	case RegisterR10w:
		return "Register R10w"
	case RegisterR10h:
		return "Register R10h"
	case RegisterR10l:
		return "Register R10l"
	case RegisterR11:
		return "Register R11"
	case RegisterR11d:
		return "Register R11d"
	case RegisterR11w:
		return "Register R11w"
	case RegisterR11h:
		return "Register R11h"
	case RegisterR11l:
		return "Register R11l"
	case RegisterR12:
		return "Register R12"
	case RegisterR12d:
		return "Register R12d"
	case RegisterR12w:
		return "Register R12w"
	case RegisterR12h:
		return "Register R12h"
	case RegisterR12l:
		return "Register R12l"
	case RegisterR13:
		return "Register R13"
	case RegisterR13d:
		return "Register R13d"
	case RegisterR13w:
		return "Register R13w"
	case RegisterR13h:
		return "Register R13h"
	case RegisterR13l:
		return "Register R13l"
	case RegisterR14:
		return "Register R14"
	case RegisterR14d:
		return "Register R14d"
	case RegisterR14w:
		return "Register R14w"
	case RegisterR14h:
		return "Register R14h"
	case RegisterR14l:
		return "Register R14l"
	case RegisterR15:
		return "Register R15"
	case RegisterR15d:
		return "Register R15d"
	case RegisterR15w:
		return "Register R15w"
	case RegisterR15h:
		return "Register R15h"
	case RegisterR15l:
		return "Register R15l"
	case RegisterDs:
		return "Register Ds"
	case RegisterEs:
		return "Register Es"
	case RegisterFs:
		return "Register Fs"
	case RegisterGs:
		return "Register Gs"
	case RegisterCs:
		return "Register Cs"
	case RegisterSs:
		return "Register Ss"
	case RegisterRflags:
		return "Register Rflags"
	case RegisterEflags:
		return "Register Eflags"
	case RegisterFlags:
		return "Register Flags"
	case RegisterCf:
		return "Register Cf"
	case RegisterPf:
		return "Register Pf"
	case RegisterAf:
		return "Register Af"
	case RegisterZf:
		return "Register Zf"
	case RegisterSf:
		return "Register Sf"
	case RegisterTf:
		return "Register Tf"
	case RegisterIf:
		return "Register If"
	case RegisterDf:
		return "Register Df"
	case RegisterOf:
		return "Register Of"
	case RegisterIopl:
		return "Register Iopl"
	case RegisterNt:
		return "Register Nt"
	case RegisterRf:
		return "Register Rf"
	case RegisterVm:
		return "Register Vm"
	case RegisterAc:
		return "Register Ac"
	case RegisterVif:
		return "Register Vif"
	case RegisterVip:
		return "Register Vip"
	case RegisterId:
		return "Register Id"
	case RegisterRip:
		return "Register Rip"
	case RegisterEip:
		return "Register Eip"
	case RegisterIp:
		return "Register Ip"
	case RegisterIdtr:
		return "Register Idtr"
	case RegisterLdtr:
		return "Register Ldtr"
	case RegisterGdtr:
		return "Register Gdtr"
	case RegisterTr:
		return "Register Tr"
	case RegisterCr0:
		return "Register Cr 0"
	case RegisterCr2:
		return "Register Cr 2"
	case RegisterCr3:
		return "Register Cr 3"
	case RegisterCr4:
		return "Register Cr 4"
	case RegisterCr8:
		return "Register Cr 8"
	case RegisterDr0:
		return "Register Dr 0"
	case RegisterDr1:
		return "Register Dr 1"
	case RegisterDr2:
		return "Register Dr 2"
	case RegisterDr3:
		return "Register Dr 3"
	case RegisterDr6:
		return "Register Dr 6"
	case RegisterDr7:
		return "Register Dr 7"
	default:
		return fmt.Sprintf("REGS_ENUM(0x%X)", uint32(r))
	}
}

func (c *CR3_TYPE) GetPcid() uint64 {
	return uint64(c.Flags & uint64(0xFFF))
}

func (c *CR3_TYPE) SetPcid(val uint64) {
	c.Flags = (c.Flags & ^uint64(0xFFF)) | uint64(uint64(val)&0xFFF)
}

func (c *CR3_TYPE) GetPageFrameNumber() uint64 {
	return uint64((c.Flags >> 12) & uint64(0xFFFFFFFFF))
}

func (c *CR3_TYPE) SetPageFrameNumber(val uint64) {
	c.Flags = (c.Flags & ^uint64(0xFFFFFFFFF<<12)) | (uint64(uint64(val)&0xFFFFFFFFF) << 12)
}

func (c *CR3_TYPE) GetReserved1() uint64 {
	return uint64((c.Flags >> 48) & uint64(0xFFF))
}

func (c *CR3_TYPE) SetReserved1(val uint64) {
	c.Flags = (c.Flags & ^uint64(0xFFF<<48)) | (uint64(uint64(val)&0xFFF) << 48)
}

func (c *CR3_TYPE) GetReserved2() uint64 {
	return uint64((c.Flags >> 60) & uint64(0x7))
}

func (c *CR3_TYPE) SetReserved2(val uint64) {
	c.Flags = (c.Flags & ^uint64(0x7<<60)) | (uint64(uint64(val)&0x7) << 60)
}

func (c *CR3_TYPE) GetPcidInvalidate() uint64 {
	return uint64((c.Flags >> 63) & uint64(0x1))
}

func (c *CR3_TYPE) SetPcidInvalidate(val uint64) {
	c.Flags = (c.Flags & ^uint64(0x1<<63)) | (uint64(uint64(val)&0x1) << 63)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetType() uint32 {
	return uint32(v.AsUInt & uint32(0xF))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetType(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0xF)) | uint32(uint32(val)&0xF)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetDescriptorType() uint32 {
	return uint32((v.AsUInt >> 4) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetDescriptorType(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<4)) | (uint32(uint32(val)&0x1) << 4)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetDescriptorPrivilegeLevel() uint32 {
	return uint32((v.AsUInt >> 5) & uint32(0x3))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetDescriptorPrivilegeLevel(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x3<<5)) | (uint32(uint32(val)&0x3) << 5)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetPresent() uint32 {
	return uint32((v.AsUInt >> 7) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetPresent(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<7)) | (uint32(uint32(val)&0x1) << 7)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetReserved1() uint32 {
	return uint32((v.AsUInt >> 8) & uint32(0xF))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetReserved1(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0xF<<8)) | (uint32(uint32(val)&0xF) << 8)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetAvailableBit() uint32 {
	return uint32((v.AsUInt >> 12) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetAvailableBit(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<12)) | (uint32(uint32(val)&0x1) << 12)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetLongMode() uint32 {
	return uint32((v.AsUInt >> 13) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetLongMode(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<13)) | (uint32(uint32(val)&0x1) << 13)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetDefaultBig() uint32 {
	return uint32((v.AsUInt >> 14) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetDefaultBig(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<14)) | (uint32(uint32(val)&0x1) << 14)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetGranularity() uint32 {
	return uint32((v.AsUInt >> 15) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetGranularity(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<15)) | (uint32(uint32(val)&0x1) << 15)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetUnusable() uint32 {
	return uint32((v.AsUInt >> 16) & uint32(0x1))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetUnusable(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x1<<16)) | (uint32(uint32(val)&0x1) << 16)
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) GetReserved2() uint32 {
	return uint32((v.AsUInt >> 17) & uint32(0x7FFF))
}

func (v *VMX_SEGMENT_ACCESS_RIGHTS_TYPE) SetReserved2(val uint32) {
	v.AsUInt = (v.AsUInt & ^uint32(0x7FFF<<17)) | (uint32(uint32(val)&0x7FFF) << 17)
}

func (p *PCI_DEV_MINIMAL) GetBus() uint8 {
	return uint8(p.BusBits & uint8(0xFF))
}

func (p *PCI_DEV_MINIMAL) SetBus(val uint8) {
	p.BusBits = (p.BusBits & ^uint8(0xFF)) | uint8(uint8(val)&0xFF)
}

func (p *PCI_DEV_MINIMAL) GetDevice() uint8 {
	return uint8(p.DeviceBits & uint8(0x1F))
}

func (p *PCI_DEV_MINIMAL) SetDevice(val uint8) {
	p.DeviceBits = (p.DeviceBits & ^uint8(0x1F)) | uint8(uint8(val)&0x1F)
}

func (p *PCI_DEV_MINIMAL) GetFunction() uint8 {
	return uint8((p.DeviceBits >> 5) & uint8(0x7))
}

func (p *PCI_DEV_MINIMAL) SetFunction(val uint8) {
	p.DeviceBits = (p.DeviceBits & ^uint8(0x7<<5)) | (uint8(uint8(val)&0x7) << 5)
}

func (p *PCI_DEV) GetBus() uint8 {
	return uint8(p.BusBits & uint8(0xFF))
}

func (p *PCI_DEV) SetBus(val uint8) {
	p.BusBits = (p.BusBits & ^uint8(0xFF)) | uint8(uint8(val)&0xFF)
}

func (p *PCI_DEV) GetDevice() uint8 {
	return uint8(p.DeviceBits & uint8(0x1F))
}

func (p *PCI_DEV) SetDevice(val uint8) {
	p.DeviceBits = (p.DeviceBits & ^uint8(0x1F)) | uint8(uint8(val)&0x1F)
}

func (p *PCI_DEV) GetFunction() uint8 {
	return uint8((p.DeviceBits >> 5) & uint8(0x7))
}

func (p *PCI_DEV) SetFunction(val uint8) {
	p.DeviceBits = (p.DeviceBits & ^uint8(0x7<<5)) | (uint8(uint8(val)&0x7) << 5)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetAssignLocalGlobalVar() uint64 {
	return uint64(h.Value & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetAssignLocalGlobalVar(val uint64) {
	h.Value = (h.Value & ^uint64(0x1)) | uint64(uint64(val)&0x1)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetAssignRegisters() uint64 {
	return uint64((h.Value >> 1) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetAssignRegisters(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<1)) | (uint64(uint64(val)&0x1) << 1)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetAssignPseudoRegisters() uint64 {
	return uint64((h.Value >> 2) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetAssignPseudoRegisters(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<2)) | (uint64(uint64(val)&0x1) << 2)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetConditionalStatementsAndComparisonOperators() uint64 {
	return uint64((h.Value >> 3) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetConditionalStatementsAndComparisonOperators(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<3)) | (uint64(uint64(val)&0x1) << 3)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetStackAssignments() uint64 {
	return uint64((h.Value >> 4) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetStackAssignments(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<4)) | (uint64(uint64(val)&0x1) << 4)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncOr() uint64 {
	return uint64((h.Value >> 5) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncOr(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<5)) | (uint64(uint64(val)&0x1) << 5)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncXor() uint64 {
	return uint64((h.Value >> 6) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncXor(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<6)) | (uint64(uint64(val)&0x1) << 6)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncAnd() uint64 {
	return uint64((h.Value >> 7) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncAnd(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<7)) | (uint64(uint64(val)&0x1) << 7)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncAsr() uint64 {
	return uint64((h.Value >> 8) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncAsr(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<8)) | (uint64(uint64(val)&0x1) << 8)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncAsl() uint64 {
	return uint64((h.Value >> 9) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncAsl(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<9)) | (uint64(uint64(val)&0x1) << 9)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncAdd() uint64 {
	return uint64((h.Value >> 10) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncAdd(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<10)) | (uint64(uint64(val)&0x1) << 10)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncSub() uint64 {
	return uint64((h.Value >> 11) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncSub(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<11)) | (uint64(uint64(val)&0x1) << 11)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncMul() uint64 {
	return uint64((h.Value >> 12) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncMul(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<12)) | (uint64(uint64(val)&0x1) << 12)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncDiv() uint64 {
	return uint64((h.Value >> 13) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncDiv(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<13)) | (uint64(uint64(val)&0x1) << 13)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncMod() uint64 {
	return uint64((h.Value >> 14) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncMod(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<14)) | (uint64(uint64(val)&0x1) << 14)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncGt() uint64 {
	return uint64((h.Value >> 15) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncGt(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<15)) | (uint64(uint64(val)&0x1) << 15)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncLt() uint64 {
	return uint64((h.Value >> 16) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncLt(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<16)) | (uint64(uint64(val)&0x1) << 16)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncEgt() uint64 {
	return uint64((h.Value >> 17) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncEgt(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<17)) | (uint64(uint64(val)&0x1) << 17)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncElt() uint64 {
	return uint64((h.Value >> 18) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncElt(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<18)) | (uint64(uint64(val)&0x1) << 18)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncEqual() uint64 {
	return uint64((h.Value >> 19) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncEqual(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<19)) | (uint64(uint64(val)&0x1) << 19)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncNeq() uint64 {
	return uint64((h.Value >> 20) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncNeq(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<20)) | (uint64(uint64(val)&0x1) << 20)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncJmp() uint64 {
	return uint64((h.Value >> 21) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncJmp(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<21)) | (uint64(uint64(val)&0x1) << 21)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncJz() uint64 {
	return uint64((h.Value >> 22) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncJz(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<22)) | (uint64(uint64(val)&0x1) << 22)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncJnz() uint64 {
	return uint64((h.Value >> 23) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncJnz(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<23)) | (uint64(uint64(val)&0x1) << 23)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncMov() uint64 {
	return uint64((h.Value >> 24) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncMov(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<24)) | (uint64(uint64(val)&0x1) << 24)
}

func (h *HWDBG_SCRIPT_CAPABILITIES) GetFuncPrintf() uint64 {
	return uint64((h.Value >> 25) & uint64(0x1))
}

func (h *HWDBG_SCRIPT_CAPABILITIES) SetFuncPrintf(val uint64) {
	h.Value = (h.Value & ^uint64(0x1<<25)) | (uint64(uint64(val)&0x1) << 25)
}

type (
	LIST_ENTRY struct {
		Flink *LIST_ENTRY
		Blink *LIST_ENTRY
	} // unknown.h:21 -> _LIST_ENTRY
	GUEST_REGS struct {
		Rax uint64
		Rcx uint64
		Rdx uint64
		Rbx uint64
		Rsp uint64
		Rbp uint64
		Rsi uint64
		Rdi uint64
		R8  uint64
		R9  uint64
		R10 uint64
		R11 uint64
		R12 uint64
		R13 uint64
		R14 uint64
		R15 uint64
	} // unknown.h:117 -> GUEST_REGS
	XMM_REG struct {
		XmmLow  uint64
		XmmHigh uint64
	} // unknown.h:147 -> XMM_REG
	GUEST_XMM_REGS struct {
		Xmm0          XMM_REG
		Xmm1          XMM_REG
		Xmm2          XMM_REG
		Xmm3          XMM_REG
		Xmm4          XMM_REG
		Xmm5          XMM_REG
		Xmm6NotSaved  XMM_REG
		Xmm7NotSaved  XMM_REG
		Xmm8NotSaved  XMM_REG
		Xmm9NotSaved  XMM_REG
		Xmm10NotSaved XMM_REG
		Xmm11NotSaved XMM_REG
		Xmm12NotSaved XMM_REG
		Xmm13NotSaved XMM_REG
		Xmm14NotSaved XMM_REG
		Xmm15NotSaved XMM_REG
		Mxcsr         uint32
		_             [4]byte
	} // unknown.h:155 -> GUEST_XMM_REGS
	GUEST_EXTRA_REGISTERS struct {
		CS     uint16
		DS     uint16
		FS     uint16
		GS     uint16
		ES     uint16
		SS     uint16
		_      [4]byte
		RFLAGS uint64
		RIP    uint64
	} // unknown.h:189 -> GUEST_EXTRA_REGISTERS
	SCRIPT_ENGINE_GENERAL_REGISTERS struct {
		StackBuffer         *uint64
		GlobalVariablesList *uint64
		StackIndx           uint64
		StackBaseIndx       uint64
		ReturnValue         uint64
	} // unknown.h:204 -> _SCRIPT_ENGINE_GENERAL_REGISTERS
	CR3_TYPE struct {
		Flags uint64
	} // unknown.h:217 -> _CR3_TYPE
	DEBUGGER_REMOTE_PACKET struct {
		Checksum                   uint8
		_                          [7]byte
		Indicator                  uint64
		TypeOfThePacket            DEBUGGER_REMOTE_PACKET_TYPE
		RequestedActionOfThePacket DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
	} // unknown.h:196 -> _DEBUGGER_REMOTE_PACKET
	DEBUGGEE_USER_INPUT_PACKET struct {
		CommandLen           uint32
		IgnoreFinishedSignal bool
		_                    [3]byte
		Result               uint32
	} // unknown.h:157 -> _DEBUGGEE_USER_INPUT_PACKET
	DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct {
		Length uint32
	} // unknown.h:173 -> _DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
	DEBUGGER_PAUSE_PACKET_RECEIVED struct {
		Result uint32
	} // unknown.h:195 -> _DEBUGGER_PAUSE_PACKET_RECEIVED
	DEBUGGER_TRIGGERED_EVENT_DETAILS struct {
		Tag     uint64
		Context uintptr
		Stage   VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
		_       [4]byte
	} // unknown.h:209 -> _DEBUGGER_TRIGGERED_EVENT_DETAILS
	DEBUGGEE_KD_PAUSED_PACKET struct {
		Rip                    uint64
		IsProcessorOn32BitMode bool
		IgnoreDisassembling    bool
		_                      [2]byte
		PausingReason          DEBUGGEE_PAUSING_REASON
		CurrentCore            uint32
		_                      [4]byte
		EventTag               uint64
		EventCallingStage      VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
		_                      [4]byte
		Rflags                 uint64
		InstructionBytesOnRip  [16]uint8
		ReadInstructionLen     uint16
		_                      [6]byte
	} // unknown.h:224 -> _DEBUGGEE_KD_PAUSED_PACKET
	DEBUGGEE_UD_PAUSED_PACKET struct {
		Rip                   uint64
		ProcessDebuggingToken uint64
		Is32Bit               bool
		_                     [3]byte
		PausingReason         DEBUGGEE_PAUSING_REASON
		ProcessId             uint32
		ThreadId              uint32
		Rflags                uint64
		EventTag              uint64
		EventCallingStage     VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
		InstructionBytesOnRip [16]uint8
		ReadInstructionLen    uint16
		_                     [2]byte
	} // unknown.h:246 -> _DEBUGGEE_UD_PAUSED_PACKET
	DEBUGGEE_MESSAGE_PACKET struct {
		OperationCode uint32
		Message       [4096]int8
	} // unknown.h:284 -> _DEBUGGEE_MESSAGE_PACKET
	REGISTER_NOTIFY_BUFFER struct {
		Type   NOTIFY_TYPE
		_      [4]byte
		HEvent uintptr
	} // unknown.h:295 -> _REGISTER_NOTIFY_BUFFER
	DIRECT_VMCALL_PARAMETERS struct {
		OptionalParam1 uint64
		OptionalParam2 uint64
		OptionalParam3 uint64
	} // unknown.h:310 -> _DIRECT_VMCALL_PARAMETERS
	SYSCALL_CALLBACK_CONTEXT_PARAMS struct {
		OptionalParam1 uint64
		OptionalParam2 uint64
		OptionalParam3 uint64
		OptionalParam4 uint64
	} // unknown.h:326 -> _SYSCALL_CALLBACK_CONTEXT_PARAMS
	EPT_HOOKS_CONTEXT struct {
		HookingTag      uint64
		PhysicalAddress uint64
		VirtualAddress  uint64
	} // unknown.h:353 -> _EPT_HOOKS_CONTEXT
	EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR struct {
		StartAddress    uint64
		EndAddress      uint64
		SetHookForRead  bool
		SetHookForWrite bool
		SetHookForExec  bool
		_               [1]byte
		MemoryType      DEBUGGER_HOOK_MEMORY_TYPE
		Tag             uint64
	} // unknown.h:364 -> _EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR
	EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 struct {
		TargetAddress uintptr
		HookFunction  uintptr
	} // unknown.h:380 -> _EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2
	EPT_SINGLE_HOOK_UNHOOKING_DETAILS struct {
		CallerNeedsToRestoreEntryAndInvalidateEpt bool
		RemoveBreakpointInterception              bool
		_                                         [6]byte
		PhysicalAddress                           uintptr
		OriginalEntry                             uint64
	} // unknown.h:391 -> _EPT_SINGLE_HOOK_UNHOOKING_DETAILS
	VMX_SEGMENT_ACCESS_RIGHTS_TYPE struct {
		AsUInt uint32
	} // unknown.h:0 -> VMX_SEGMENT_ACCESS_RIGHTS_TYPE
	VMX_SEGMENT_SELECTOR struct {
		Selector   uint16
		_          [2]byte
		Attributes VMX_SEGMENT_ACCESS_RIGHTS_TYPE
		Limit      uint32
		_          [4]byte
		Base       uint64
	} // unknown.h:469 -> _VMX_SEGMENT_SELECTOR
	DEBUGGER_MODIFY_EVENTS struct {
		Tag          uint64
		KernelStatus uint64
		TypeOfAction DEBUGGER_MODIFY_EVENTS_TYPE
		IsEnabled    bool
		_            [3]byte
	} // unknown.h:246 -> _DEBUGGER_MODIFY_EVENTS
	DEBUGGER_SHORT_CIRCUITING_EVENT struct {
		KernelStatus      uint64
		IsShortCircuiting bool
		_                 [7]byte
	} // unknown.h:260 -> _DEBUGGER_SHORT_CIRCUITING_EVENT
	DEBUGGER_EVENT_OPTIONS struct {
		OptionalParam1 uint64
		OptionalParam2 uint64
		OptionalParam3 uint64
		OptionalParam4 uint64
		OptionalParam5 uint64
		OptionalParam6 uint64
	} // unknown.h:275 -> _DEBUGGER_EVENT_OPTIONS
	DEBUGGER_GENERAL_EVENT_DETAIL struct {
		CommandsEventList     LIST_ENTRY
		CoreId                uint32
		ProcessId             uint32
		IsEnabled             bool
		EnableShortCircuiting bool
		_                     [2]byte
		EventStage            VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
		HasCustomOutput       bool
		_                     [7]byte
		OutputSourceTags      [5]uint64
		CountOfActions        uint32
		_                     [4]byte
		Tag                   uint64
		EventType             VMM_EVENT_TYPE_ENUM
		_                     [4]byte
		Options               DEBUGGER_EVENT_OPTIONS
		CommandStringBuffer   uintptr
		ConditionBufferSize   uint32
		_                     [4]byte
	} // unknown.h:356 -> _DEBUGGER_GENERAL_EVENT_DETAIL
	DEBUGGER_GENERAL_ACTION struct {
		EventTag                uint64
		ActionType              DEBUGGER_EVENT_ACTION_TYPE_ENUM
		ImmediateMessagePassing bool
		_                       [3]byte
		PreAllocatedBuffer      uint32
		CustomCodeBufferSize    uint32
		ScriptBufferSize        uint32
		ScriptBufferPointer     uint32
	} // unknown.h:410 -> _DEBUGGER_GENERAL_ACTION
	DEBUGGER_EVENT_AND_ACTION_RESULT struct {
		IsSuccessful bool
		_            [3]byte
		Error        uint32
	} // unknown.h:427 -> _DEBUGGER_EVENT_AND_ACTION_RESULT
	PORTABLE_PCI_COMMON_HEADER struct {
		VendorId            uint16
		DeviceId            uint16
		Command             uint16
		Status              uint16
		RevisionId          uint8
		ClassCode           [3]uint8
		CacheLineSize       uint8
		PrimaryLatencyTimer uint8
		HeaderType          uint8
		Bist                uint8
	} // unknown.h:49 -> _PORTABLE_PCI_COMMON_HEADER
	PORTABLE_PCI_DEVICE_HEADER_ struct {
		ConfigSpaceEp        PORTABLE_PCI_EP_HEADER
		ConfigSpacePtpBridge PORTABLE_PCI_BRIDGE_HEADER
	} // unknown.h:0 -> _PORTABLE_PCI_DEVICE_HEADER
	PORTABLE_PCI_EP_HEADER struct {
		Bar             [6]uint32
		CardBusCISPtr   uint32
		SubVendorId     uint16
		SubSystemId     uint16
		ROMBar          uint32
		CapabilitiesPtr uint8
		Reserved        [3]uint8
		Reserved1       uint32
		InterruptLine   uint8
		InterruptPin    uint8
		MinGnt          uint8
		MaxLat          uint8
	} // unknown.h:0 -> _PORTABLE_PCI_EP_HEADER
	PORTABLE_PCI_BRIDGE_HEADER struct {
		Bar                       [2]uint32
		PrimaryBusNumber          uint8
		SecondaryBusNumber        uint8
		SubordinateBusNumber      uint8
		SecondaryLatencyTimer     uint8
		IoBase                    uint8
		IoLimit                   uint8
		SecondaryStatus           uint16
		MemoryBase                uint16
		MemoryLimit               uint16
		PrefetchableMemoryBase    uint16
		PrefetchableMemoryLimit   uint16
		PrefetchableBaseUpper32b  uint32
		PrefetchableLimitUpper32b uint32
		IoBaseUpper16b            uint16
		IoLimitUpper16b           uint16
		CapabilityPtr             uint8
		Reserved                  [3]uint8
		ROMBar                    uint32
		InterruptLine             uint8
		InterruptPin              uint8
		BridgeControl             uint16
	} // unknown.h:0 -> _PORTABLE_PCI_BRIDGE_HEADER
	PORTABLE_PCI_DEVICE_HEADER struct {
		Data PORTABLE_PCI_EP_HEADER
	} // unknown.h:0 -> _PORTABLE_PCI_DEVICE_HEADER
	PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL struct {
		VendorId  uint16
		DeviceId  uint16
		ClassCode [3]uint8
		_         [1]byte
	} // unknown.h:118 -> _PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL
	PCI_DEV_MINIMAL struct {
		BusBits     uint8
		DeviceBits  uint8
		ConfigSpace PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL
	} // unknown.h:132 -> _PCI_DEV_MINIMAL
	PCI_DEV_MMIOBAR_INFO struct {
		Is64Bit      int32
		IsEnabled    int32
		BarOffsetEnd uint64
		BarSize      uint64
	} // unknown.h:144 -> _PCI_DEV_MMIOBAR_INFO
	PORTABLE_PCI_CONFIG_SPACE_HEADER struct {
		CommonHeader PORTABLE_PCI_COMMON_HEADER
		DeviceHeader PORTABLE_PCI_DEVICE_HEADER
	} // unknown.h:156 -> _PORTABLE_PCI_CONFIG_SPACE_HEADER
	PCI_DEV struct {
		BusBits               uint8
		_                     [3]byte
		ConfigSpace           PORTABLE_PCI_CONFIG_SPACE_HEADER
		DeviceBits            uint8
		_                     [66]byte
		ConfigSpaceAdditional [191]uint8
		_                     [5]byte
		MmioBarInfo           [6]PCI_DEV_MMIOBAR_INFO
	} // unknown.h:166 -> _PCI_DEV
	DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS struct {
		VirtualAddress      uint64
		ProcessId           uint32
		_                   [4]byte
		Pml4eVirtualAddress uint64
		Pml4eValue          uint64
		PdpteVirtualAddress uint64
		PdpteValue          uint64
		PdeVirtualAddress   uint64
		PdeValue            uint64
		PteVirtualAddress   uint64
		PteValue            uint64
		KernelStatus        uint32
		_                   [4]byte
	} // unknown.h:22 -> _DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
	DEBUGGER_VA2PA_AND_PA2VA_COMMANDS struct {
		VirtualAddress     uint64
		PhysicalAddress    uint64
		ProcessId          uint32
		IsVirtual2Physical bool
		_                  [3]byte
		KernelStatus       uint32
		_                  [4]byte
	} // unknown.h:53 -> _DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
	DEBUGGER_PAGE_IN_REQUEST struct {
		VirtualAddressFrom uint64
		VirtualAddressTo   uint64
		ProcessId          uint32
		PageFaultErrorCode uint32
		KernelStatus       uint32
		_                  [4]byte
	} // unknown.h:72 -> _DEBUGGER_PAGE_IN_REQUEST
	REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST struct {
		ProcessId    uint32
		Size         uint32
		Mode         REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE
		Type         REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE
		KernelStatus uint32
	} // unknown.h:113 -> _REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST
	DEBUGGER_DT_COMMAND_OPTIONS struct {
		TypeName             *int8
		SizeOfTypeName       uint64
		Address              uint64
		IsStruct             bool
		_                    [7]byte
		BufferAddress        uintptr
		TargetPid            uint32
		_                    [4]byte
		AdditionalParameters *int8
	} // unknown.h:132 -> _DEBUGGER_DT_COMMAND_OPTIONS
	DEBUGGER_PREALLOC_COMMAND struct {
		Type         DEBUGGER_PREALLOC_COMMAND_TYPE
		Count        uint32
		KernelStatus uint32
	} // unknown.h:170 -> _DEBUGGER_PREALLOC_COMMAND
	DEBUGGER_PREACTIVATE_COMMAND struct {
		Type         DEBUGGER_PREACTIVATE_COMMAND_TYPE
		KernelStatus uint32
	} // unknown.h:197 -> _DEBUGGER_PREACTIVATE_COMMAND
	DEBUGGER_READ_MEMORY struct {
		Pid            uint32
		_              [4]byte
		Address        uint64
		Size           uint32
		GetAddressMode bool
		_              [3]byte
		AddressMode    DEBUGGER_READ_MEMORY_ADDRESS_MODE
		MemoryType     DEBUGGER_READ_MEMORY_TYPE
		ReadingType    DEBUGGER_READ_READING_TYPE
		ReturnLength   uint32
		KernelStatus   uint32
		_              [4]byte
	} // unknown.h:260 -> _DEBUGGER_READ_MEMORY
	DEBUGGER_FLUSH_LOGGING_BUFFERS struct {
		KernelStatus                               uint32
		CountOfMessagesThatSetAsReadFromVmxRoot    uint32
		CountOfMessagesThatSetAsReadFromVmxNonRoot uint32
	} // unknown.h:287 -> _DEBUGGER_FLUSH_LOGGING_BUFFERS
	DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER struct {
		RequestType  DEBUGGER_TEST_QUERY_STATE
		_            [4]byte
		Context      uint64
		KernelStatus uint32
		_            [4]byte
	} // unknown.h:325 -> _DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
	DEBUGGER_PERFORM_KERNEL_TESTS struct {
		KernelStatus uint32
	} // unknown.h:342 -> _DEBUGGER_PERFORM_KERNEL_TESTS
	DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL struct {
		KernelStatus uint32
	} // unknown.h:357 -> _DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
	DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER struct {
		RequestedAction       DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
		LengthOfBuffer        uint32
		PauseDebuggeeWhenSent bool
		_                     [3]byte
		KernelResult          uint32
	} // unknown.h:373 -> _DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
	DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER struct {
		KernelStatus uint32
		Length       uint32
	} // unknown.h:396 -> _DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
	DEBUGGER_READ_AND_WRITE_ON_MSR struct {
		Msr        uint64
		CoreNumber uint32
		ActionType DEBUGGER_MSR_ACTION_TYPE
		Value      uint64
	} // unknown.h:427 -> _DEBUGGER_READ_AND_WRITE_ON_MSR
	DEBUGGER_EDIT_MEMORY struct {
		Result             uint32
		_                  [4]byte
		Address            uint64
		ProcessId          uint32
		MemoryType         DEBUGGER_EDIT_MEMORY_TYPE
		ByteSize           DEBUGGER_EDIT_MEMORY_BYTE_SIZE
		CountOf64Chunks    uint32
		FinalStructureSize uint32
		_                  [4]byte
	} // unknown.h:468 -> _DEBUGGER_EDIT_MEMORY
	DEBUGGER_SEARCH_MEMORY struct {
		Address            uint64
		Length             uint64
		ProcessId          uint32
		MemoryType         DEBUGGER_SEARCH_MEMORY_TYPE
		ByteSize           DEBUGGER_SEARCH_MEMORY_BYTE_SIZE
		CountOf64Chunks    uint32
		FinalStructureSize uint32
		_                  [4]byte
	} // unknown.h:512 -> _DEBUGGER_SEARCH_MEMORY
	SYSTEM_CALL_NUMBERS_INFORMATION struct {
		SysNtQuerySystemInformation   uint32
		SysNtQuerySystemInformationEx uint32
		SysNtSystemDebugControl       uint32
		SysNtQueryAttributesFile      uint32
		SysNtOpenDirectoryObject      uint32
		SysNtQueryDirectoryObject     uint32
		SysNtQueryInformationProcess  uint32
		SysNtSetInformationProcess    uint32
		SysNtQueryInformationThread   uint32
		SysNtSetInformationThread     uint32
		SysNtOpenFile                 uint32
		SysNtOpenKey                  uint32
		SysNtOpenKeyEx                uint32
		SysNtQueryValueKey            uint32
		SysNtEnumerateKey             uint32
	} // unknown.h:534 -> _SYSTEM_CALL_NUMBERS_INFORMATION
	DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE struct {
		IsHide                               bool
		TrueIfProcessIdAndFalseIfProcessName bool
		_                                    [2]byte
		ProcId                               uint32
		LengthOfProcessName                  uint32
		SystemCallNumbersInformation         SYSTEM_CALL_NUMBERS_INFORMATION
		KernelStatus                         uint32
	} // unknown.h:562 -> _DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	DEBUGGER_PREPARE_DEBUGGEE struct {
		PortAddress       uint32
		Baudrate          uint32
		KernelBaseAddress uint64
		Result            uint32
		OsName            [256]int8
		_                 [4]byte
	} // unknown.h:595 -> _DEBUGGER_PREPARE_DEBUGGEE
	DEBUGGEE_CHANGE_CORE_PACKET struct {
		NewCore uint32
		Result  uint32
	} // unknown.h:611 -> _DEBUGGEE_CHANGE_CORE_PACKET
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS struct {
		IsStartingNewProcess                      bool
		_                                         [3]byte
		ProcessId                                 uint32
		ThreadId                                  uint32
		CheckCallbackAtFirstInstruction           bool
		Is32Bit                                   bool
		_                                         [2]byte
		Rip                                       uint64
		InstructionBytesOnRip                     [16]uint8
		SizeOfInstruction                         uint32
		IsPaused                                  bool
		_                                         [3]byte
		Action                                    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE
		CountOfActiveDebuggingThreadsAndProcesses uint32
		Token                                     uint64
		Result                                    uint64
	} // unknown.h:644 -> _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
	DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS struct {
		PsActiveProcessHead      uint64
		ImageFileNameOffset      uint32
		UniquePidOffset          uint32
		ActiveProcessLinksOffset uint32
		_                        [4]byte
	} // unknown.h:700 -> _DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
	DEBUGGEE_THREAD_LIST_NEEDED_DETAILS struct {
		ThreadListHeadOffset     uint32
		ThreadListEntryOffset    uint32
		CidOffset                uint32
		_                        [4]byte
		PsActiveProcessHead      uint64
		ActiveProcessLinksOffset uint32
		_                        [4]byte
		Process                  uint64
	} // unknown.h:714 -> _DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
	DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY struct {
		Eprocess      uint64
		ProcessId     uint32
		_             [4]byte
		Cr3           uint64
		ImageFileName [16]uint8
	} // unknown.h:730 -> _DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY
	DEBUGGEE_THREAD_LIST_DETAILS_ENTRY struct {
		Eprocess      uint64
		Ethread       uint64
		ProcessId     uint32
		ThreadId      uint32
		ImageFileName [16]uint8
	} // unknown.h:744 -> _DEBUGGEE_THREAD_LIST_DETAILS_ENTRY
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS struct {
		ProcessListNeededDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
		ThreadListNeededDetails  DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
		QueryType                DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES
		QueryAction              DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS
		Count                    uint32
		_                        [4]byte
		Result                   uint64
	} // unknown.h:758 -> _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
	DEBUGGER_SINGLE_CALLSTACK_FRAME struct {
		IsStackAddressValid   bool
		IsValidAddress        bool
		IsExecutable          bool
		_                     [5]byte
		Value                 uint64
		InstructionBytesOnRip [7]uint8
		_                     [1]byte
	} // unknown.h:776 -> _DEBUGGER_SINGLE_CALLSTACK_FRAME
	DEBUGGER_CALLSTACK_REQUEST struct {
		Is32Bit       bool
		_             [3]byte
		KernelStatus  uint32
		DisplayMethod DEBUGGER_CALLSTACK_DISPLAY_METHOD
		Size          uint32
		FrameCount    uint32
		_             [4]byte
		BaseAddress   uint64
		BufferSize    uint64
	} // unknown.h:804 -> _DEBUGGER_CALLSTACK_REQUEST
	USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS struct {
		ProcessId                      uint32
		ThreadId                       uint32
		NumberOfBlockedContextSwitches uint64
		IsProcess                      bool
		_                              [7]byte
	} // unknown.h:825 -> _USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS
	DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION struct {
		ScriptBuffer                uint64
		ScriptLength                uint32
		ScriptPointer               uint32
		OptionalRequestedBufferSize uint32
		_                           [4]byte
	} // unknown.h:840 -> _DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
	DEBUGGER_EVENT_REQUEST_BUFFER struct {
		EnabledRequestBuffer bool
		_                    [3]byte
		RequestBufferSize    uint32
		RequstBufferAddress  uint64
	} // unknown.h:854 -> _DEBUGGER_EVENT_REQUEST_BUFFER
	DEBUGGER_EVENT_REQUEST_CUSTOM_CODE struct {
		CustomCodeBufferSize        uint32
		_                           [4]byte
		CustomCodeBufferAddress     uintptr
		OptionalRequestedBufferSize uint32
		_                           [4]byte
	} // unknown.h:866 -> _DEBUGGER_EVENT_REQUEST_CUSTOM_CODE
	DEBUGGER_UD_COMMAND_ACTION struct {
		ActionType     DEBUGGER_UD_COMMAND_ACTION_TYPE
		_              [4]byte
		OptionalParam1 uint64
		OptionalParam2 uint64
		OptionalParam3 uint64
		OptionalParam4 uint64
	} // unknown.h:894 -> _DEBUGGER_UD_COMMAND_ACTION
	DEBUGGER_UD_COMMAND_PACKET struct {
		UdAction                    DEBUGGER_UD_COMMAND_ACTION
		ProcessDebuggingDetailToken uint64
		TargetThreadId              uint32
		ApplyToAllPausedThreads     bool
		WaitForEventCompletion      bool
		_                           [2]byte
		Result                      uint32
		_                           [4]byte
	} // unknown.h:908 -> _DEBUGGER_UD_COMMAND_PACKET
	DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET struct {
		ActionType            DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE
		ProcessId             uint32
		Process               uint64
		IsSwitchByClkIntr     bool
		ProcessName           [16]uint8
		_                     [7]byte
		ProcessListSymDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
		Result                uint32
		_                     [4]byte
	} // unknown.h:939 -> _DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
	DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET struct {
		ActionType            DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE
		ThreadId              uint32
		ProcessId             uint32
		_                     [4]byte
		Thread                uint64
		Process               uint64
		CheckByClockInterrupt bool
		ProcessName           [16]uint8
		_                     [7]byte
		ThreadListSymDetails  DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
		Result                uint32
		_                     [4]byte
	} // unknown.h:977 -> _DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
	DEBUGGEE_STEP_PACKET struct {
		StepType                  DEBUGGER_REMOTE_STEPPING_REQUEST
		IsCurrentInstructionACall bool
		_                         [3]byte
		CallLength                uint32
	} // unknown.h:1020 -> _DEBUGGEE_STEP_PACKET
	DEBUGGER_APIC_REQUEST struct {
		ApicType      DEBUGGER_APIC_REQUEST_TYPE
		IsUsingX2APIC bool
		_             [3]byte
		KernelStatus  uint32
	} // unknown.h:1056 -> _DEBUGGER_APIC_REQUEST
	LAPIC_PAGE struct {
		Reserved000             [16]uint8
		Reserved010             [16]uint8
		Id                      uint32
		Reserved024             [12]uint8
		Version                 uint32
		Reserved034             [12]uint8
		Reserved040             [64]uint8
		TPR                     uint32
		Reserved084             [12]uint8
		ArbitrationPriority     uint32
		Reserved094             [12]uint8
		ProcessorPriority       uint32
		Reserved0A4             [12]uint8
		EOI                     uint32
		Reserved0B4             [12]uint8
		RemoteRead              uint32
		Reserved0C4             [12]uint8
		LogicalDestination      uint32
		Reserved0D4             [12]uint8
		DestinationFormat       uint32
		Reserved0E4             [12]uint8
		SpuriousInterruptVector uint32
		Reserved0F4             [12]uint8
		ISR                     [32]uint32
		TMR                     [32]uint32
		IRR                     [32]uint32
		ErrorStatus             uint32
		Reserved284             [12]uint8
		Reserved290             [96]uint8
		LvtCmci                 uint32
		Reserved2F4             [12]uint8
		IcrLow                  uint32
		Reserved304             [12]uint8
		IcrHigh                 uint32
		Reserved314             [12]uint8
		LvtTimer                uint32
		Reserved324             [12]uint8
		LvtThermalSensor        uint32
		Reserved334             [12]uint8
		LvtPerfMonCounters      uint32
		Reserved344             [12]uint8
		LvtLINT0                uint32
		Reserved354             [12]uint8
		LvtLINT1                uint32
		Reserved364             [12]uint8
		LvtError                uint32
		Reserved374             [12]uint8
		InitialCount            uint32
		Reserved384             [12]uint8
		CurrentCount            uint32
		Reserved394             [12]uint8
		Reserved3A0             [64]uint8
		DivideConfiguration     uint32
		Reserved3E4             [12]uint8
		SelfIpi                 uint32
		Reserved3F4             [12]uint8
	} // unknown.h:1083 -> _LAPIC_PAGE
	IO_APIC_ENTRY_PACKETS struct {
		ApicBasePa uint64
		ApicBaseVa uint64
		IoIdReg    uint32
		IoLl       uint32
		IoArbIdReg uint32
		_          [4]byte
		LlLhData   [400]uint64
	} // unknown.h:1188 -> _IO_APIC_ENTRY_PACKETS
	SMI_OPERATION_PACKETS struct {
		SmiOperationType SMI_OPERATION_REQUEST_TYPE
		_                [4]byte
		SmiCount         uint64
		KernelStatus     uint32
		_                [4]byte
	} // unknown.h:1223 -> _SMI_OPERATION_PACKETS
	HYPERTRACE_LBR_OPERATION_PACKETS struct {
		LbrOperationType HYPERTRACE_LBR_OPERATION_REQUEST_TYPE
		LbrFilterOptions uint32
		KernelStatus     uint32
	} // unknown.h:1259 -> _HYPERTRACE_LBR_OPERATION_PACKETS
	HYPERTRACE_PT_OPERATION_PACKETS struct {
		PtOperationType HYPERTRACE_PT_OPERATION_REQUEST_TYPE
		KernelStatus    uint32
	} // unknown.h:1293 -> _HYPERTRACE_PT_OPERATION_PACKETS
	INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS struct {
		KernelStatus uint32
		_            [4]byte
		IdtEntry     [256]uint64
	} // unknown.h:1319 -> _INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS
	DEBUGGEE_FORMATS_PACKET struct {
		Value  uint64
		Result uint32
		_      [4]byte
	} // unknown.h:1346 -> _DEBUGGEE_FORMATS_PACKET
	DEBUGGEE_SYMBOL_REQUEST_PACKET struct {
		ProcessId uint32
	} // unknown.h:1359 -> _DEBUGGEE_SYMBOL_REQUEST_PACKET
	DEBUGGEE_BP_PACKET struct {
		Address           uint64
		Pid               uint32
		Tid               uint32
		Core              uint32
		RemoveAfterHit    bool
		CheckForCallbacks bool
		_                 [2]byte
		Result            uint32
		_                 [4]byte
	} // unknown.h:1371 -> _DEBUGGEE_BP_PACKET
	DEBUGGEE_BP_LIST_OR_MODIFY_PACKET struct {
		BreakpointId uint64
		Request      DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST
		Result       uint32
	} // unknown.h:1408 -> _DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
	DEBUGGEE_SCRIPT_PACKET struct {
		ScriptBufferSize    uint32
		ScriptBufferPointer uint32
		IsFormat            bool
		_                   [7]byte
		FormatValue         uint64
		Result              uint32
		_                   [4]byte
	} // unknown.h:1438 -> _DEBUGGEE_SCRIPT_PACKET
	DEBUGGEE_RESULT_OF_SEARCH_PACKET struct {
		CountOfResults uint32
		Result         uint32
	} // unknown.h:1458 -> _DEBUGGEE_RESULT_OF_SEARCH_PACKET
	DEBUGGEE_REGISTER_READ_DESCRIPTION struct {
		RegisterId   uint32
		_            [4]byte
		Value        uint64
		KernelStatus uint32
		_            [4]byte
	} // unknown.h:1471 -> _DEBUGGEE_REGISTER_READ_DESCRIPTION
	DEBUGGEE_REGISTER_WRITE_DESCRIPTION struct {
		RegisterId   uint32
		_            [4]byte
		Value        uint64
		KernelStatus uint32
		_            [4]byte
	} // unknown.h:1485 -> _DEBUGGEE_REGISTER_WRITE_DESCRIPTION
	DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET struct {
		KernelStatus      uint32
		DeviceInfoListNum uint8
		_                 [1]byte
		DeviceInfoList    [255]PCI_DEV_MINIMAL
	} // unknown.h:1502 -> _DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET
	DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET struct {
		KernelStatus uint32
		PrintRaw     int32
		DeviceInfo   PCI_DEV
	} // unknown.h:1526 -> _DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET
	MODULE_SYMBOL_DETAIL struct {
		IsSymbolDetailsFound   bool
		IsLocalSymbolPath      bool
		IsSymbolPDBAvaliable   bool
		IsUserMode             bool
		Is32Bit                bool
		_                      [3]byte
		BaseAddress            uint64
		FilePath               [260]int8
		ModuleSymbolPath       [260]int8
		ModuleSymbolGuidAndAge [60]int8
		_                      [4]byte
	} // unknown.h:23 -> _MODULE_SYMBOL_DETAIL
	USERMODE_LOADED_MODULE_SYMBOLS struct {
		BaseAddress uint64
		Entrypoint  uint64
		FilePath    [260]uint16
	} // unknown.h:38 -> _USERMODE_LOADED_MODULE_SYMBOLS
	USERMODE_LOADED_MODULE_DETAILS struct {
		ProcessId        uint32
		OnlyCountModules bool
		Is32Bit          bool
		_                [2]byte
		ModulesCount     uint32
		Result           uint32
	} // unknown.h:46 -> _USERMODE_LOADED_MODULE_DETAILS
	DEBUGGER_UPDATE_SYMBOL_TABLE struct {
		TotalSymbols       uint32
		CurrentSymbolIndex uint32
		SymbolDetailPacket MODULE_SYMBOL_DETAIL
	} // unknown.h:72 -> _DEBUGGER_UPDATE_SYMBOL_TABLE
	DEBUGGEE_SYMBOL_UPDATE_RESULT struct {
		KernelStatus uint64
	} // unknown.h:88 -> _DEBUGGEE_SYMBOL_UPDATE_RESULT
	HWDBG_PORT_INFORMATION_ITEMS struct {
		PortSize uint32
	} // unknown.h:104 -> _HWDBG_PORT_INFORMATION_ITEMS
	HWDBG_INSTANCE_INFORMATION struct {
		Version                                    uint32
		MaximumNumberOfStages                      uint32
		ScriptVariableLength                       uint32
		NumberOfSupportedLocalAndGlobalVariables   uint32
		NumberOfSupportedTemporaryVariables        uint32
		MaximumNumberOfSupportedGetScriptOperators uint32
		MaximumNumberOfSupportedSetScriptOperators uint32
		SharedMemorySize                           uint32
		DebuggerAreaOffset                         uint32
		DebuggeeAreaOffset                         uint32
		NumberOfPins                               uint32
		NumberOfPorts                              uint32
		ScriptCapabilities                         HWDBG_SCRIPT_CAPABILITIES
		BramAddrWidth                              uint32
		BramDataWidth                              uint32
	} // unknown.h:115 -> _HWDBG_INSTANCE_INFORMATION
	HWDBG_SCRIPT_CAPABILITIES struct {
		Value uint64
	} // unknown.h:0 -> _HWDBG_SCRIPT_CAPABILITIES
	HWDBG_SCRIPT_BUFFER struct {
		ScriptNumberOfSymbols uint32
	} // unknown.h:194 -> _HWDBG_SCRIPT_BUFFER
	SYMBOL struct {
		Type  uint64
		Len   uint64
		Value uint64
	} // unknown.h:5 -> SYMBOL
	HWDBG_SHORT_SYMBOL struct {
		Type  uint64
		Value uint64
	} // unknown.h:15 -> HWDBG_SHORT_SYMBOL
	SYMBOL_BUFFER struct {
		Head    PSYMBOL
		Pointer uint32
		Size    uint32
		Message *int8
	} // unknown.h:22 -> SYMBOL_BUFFER
	SYMBOL_MAP struct {
		Name *int8
		Type uint64
	} // unknown.h:29 -> SYMBOL_MAP
	ACTION_BUFFER struct {
		Tag                       uint64
		CurrentAction             uint64
		ImmediatelySendTheResults int8
		_                         [7]byte
		Context                   uint64
		CallingStage              int8
		_                         [7]byte
	} // unknown.h:35 -> ACTION_BUFFER
)
