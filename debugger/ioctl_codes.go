package debugger

func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return ((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method)
}

const (
	FILE_DEVICE_UNKNOWN = 0x00000022
	METHOD_BUFFERED     = 0
	FILE_ANY_ACCESS     = 0
)

var (
	// IOCTL_REGISTER_EVENT - ioctl, register a new event
	IoctlRegisterEvent = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL - irp pending mechanism for reading from message tracing buffers
	IoctlReturnIrpPendingPacketsAndDisallowIoctl = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_TERMINATE_VMX - ioctl, to terminate vmx and exit form debugger
	IoctlTerminateVmx = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_READ_MEMORY - ioctl, request to read memory
	IoctlDebuggerReadMemory = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x803, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_READ_OR_WRITE_MSR - ioctl, request to read or write on a special MSR
	IoctlDebuggerReadOrWriteMsr = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x804, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS - ioctl, request to read page table entries
	IoctlDebuggerReadPageTableEntriesDetails = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x805, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_REGISTER_EVENT - ioctl, register an event
	IoctlDebuggerRegisterEvent = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x806, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT - ioctl, add action to event
	IoctlDebuggerAddActionToEvent = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x807, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER - ioctl, request to enable or disable transparent-mode
	IoctlDebuggerHideAndUnhideToTransparentTheDebugger = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x808, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS - ioctl, for !va2pa and !pa2va commands
	IoctlDebuggerVa2paAndPa2vaCommands = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x809, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_EDIT_MEMORY - ioctl, request to edit virtual and physical memory
	IoctlDebuggerEditMemory = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80a, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_SEARCH_MEMORY - ioctl, request to search virtual and physical memory
	IoctlDebuggerSearchMemory = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80b, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_MODIFY_EVENTS - ioctl, request to modify an event (enable/disable/clear)
	IoctlDebuggerModifyEvents = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80c, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS - ioctl, flush the kernel buffers
	IoctlDebuggerFlushLoggingBuffers = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80d, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS - ioctl, attach or detach user-mode processes
	IoctlDebuggerAttachDetachUserModeProcess = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80e, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_PRINT - ioctl, print states (Deprecated)
	IoctlDebuggerPrint = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80f, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PREPARE_DEBUGGEE - ioctl, prepare debuggee
	IoctlPrepareDebuggee = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x810, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PAUSE_PACKET_RECEIVED - ioctl, pause and halt the system
	IoctlPausePacketReceived = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x811, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED - ioctl, send a signal that execution of command finished
	IoctlSendSignalExecutionInDebuggeeFinished = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x812, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER - ioctl, send user-mode messages to the debugger
	IoctlSendUsermodeMessagesToDebugger = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x813, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER - ioctl, send general buffer from debuggee to debugger
	IoctlSendGeneralBufferFromDebuggeeToDebugger = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x814, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PERFORM_KERNEL_SIDE_TESTS - ioctl, to perform kernel-side tests
	IoctlPerformKernelSideTests = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x815, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_RESERVE_PRE_ALLOCATED_POOLS - ioctl, to reserve pre-allocated pools
	IoctlReservePreAllocatedPools = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x816, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SEND_USER_DEBUGGER_COMMANDS - ioctl, to send user debugger commands
	IoctlSendUserDebuggerCommands = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x817, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES - ioctl, to get active threads/processes that are debugging
	IoctlGetDetailOfActiveThreadsAndProcesses = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x818, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_GET_USER_MODE_MODULE_DETAILS - ioctl, to get user mode modules details
	IoctlGetUserModeModuleDetails = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x819, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS - ioctl, query count of active threads or processes
	IoctlQueryCountOfActiveProcessesOrThreads = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81a, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES - ioctl, to get list threads/processes
	IoctlGetListOfThreadsAndProcesses = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81b, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_QUERY_CURRENT_PROCESS - ioctl, query the current process details
	IoctlQueryCurrentProcess = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81c, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_QUERY_CURRENT_THREAD - ioctl, query the current thread details
	IoctlQueryCurrentThread = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81d, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_REQUEST_REV_MACHINE_SERVICE - ioctl, request service from the reversing machine
	IoctlRequestRevMachineService = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81e, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_DEBUGGER_BRING_PAGES_IN - ioctl, request to bring pages in
	IoctlDebuggerBringPagesIn = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81f, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PREACTIVATE_FUNCTIONALITY - ioctl, to preactivate a functionality
	IoctlPreactivateFunctionality = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PCIE_ENDPOINT_ENUM - ioctl, to enumerate PCIe endpoints
	IoctlPcieEndpointEnum = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x821, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PERFORM_ACTIONS_ON_APIC - ioctl, to perform actions related to APIC
	IoctlPerformActionsOnApic = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x822, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PCIDEVINFO_ENUM - ioctl, to query for PCI endpoint info
	IoctlPcidevinfoEnum = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x823, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_QUERY_IDT_ENTRY - ioctl, to query the IDT entries
	IoctlQueryIdtEntry = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x824, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SET_BREAKPOINT_USER_DEBUGGER - ioctl, to set breakpoint for the user debugger
	IoctlSetBreakpointUserDebugger = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x825, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_PERFORM_SMI_OPERATION - ioctl, to perform SMI operations
	IoctlPerformSmiOperation = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x826, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SWITCH_PROCESS - ioctl, switch process context
	IoctlSwitchProcess = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x827, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_SWITCH_THREAD - ioctl, switch thread context
	IoctlSwitchThread = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x828, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_READ_CONTROL_REGISTER - ioctl, read control register
	IoctlReadControlRegister = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x829, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_KILL_PROCESS - ioctl, kill process
	IoctlKillProcess = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x82a, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_RESTART_PROCESS - ioctl, restart process
	IoctlRestartProcess = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x82b, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// IOCTL_EVALUATE_EXPRESSION - ioctl, evaluate expression
	IoctlEvaluateExpression = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x82c, METHOD_BUFFERED, FILE_ANY_ACCESS)

	// CommunicationBufferSize - size of communication buffer
	CommunicationBufferSize = uint32(0x1000)
)
