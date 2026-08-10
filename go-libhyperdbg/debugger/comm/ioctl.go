// Package comm implements the user-mode side of the HyperDbg kernel-driver
// communication channel: it wraps CreateFile + DeviceIoControl (no cgo, pure
// syscalls via golang.org/x/sys/windows) and defines every IOCTL code exposed
// by the VMM/hyperkd driver.
//
// The IOCTL codes and the CTL_CODE macro mirror
// include/SDK/headers/Ioctls.h verbatim, so a buffer packed by the Go types
// package (go-libhyperdbg/types) can be sent to the driver with byte-for-byte
// the same semantics as the C libhyperdbg.
//
// Named-pipe and TCP transports (forwarding/remote-connection) live in
// sub-packages; this file owns only the IOCTL device path.
package comm

// CTL_CODE replicates the Windows SDK macro at runtime:
//
//	CTL_CODE(DeviceType, Function, Method, Access) =
//	    (DeviceType << 16) | (Access << 14) | (Function << 2) | Method
//
// HyperDbg uses FILE_DEVICE_UNKNOWN (0x22), METHOD_BUFFERED (0) and
// FILE_ANY_ACCESS (0) for every IOCTL, so the effective encoding is
// 0x220000 | (Function << 2). The const block below uses that inlined form
// because Go const blocks cannot call functions.
func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return (deviceType << 16) | (access << 14) | (function << 2) | method
}

// IOCTL parameter constants (see Ioctls.h).
const (
	fileDeviceUnknown = 0x00000022
	methodBuffered    = 0
	fileAnyAccess     = 0

	ioctlStartCode  = 0x800
	ioctlBasic      = ioctlStartCode + 0x00
	ioctlKD         = ioctlStartCode + 0x100
	ioctlVMM        = ioctlStartCode + 0x200
	ioctlHyperTrace = ioctlStartCode + 0x300

	// ioctlBase is the constant prefix common to every HyperDbg IOCTL:
	// fileDeviceUnknown<<16 | fileAnyAccess<<14 | methodBuffered == 0x220000.
	// Each IOCTL code below is ioctlBase | (function << 2), which is a valid
	// Go constant expression (no function call).
	ioctlBase = fileDeviceUnknown<<16 | fileAnyAccess<<14 | methodBuffered
)

// Pre-computed IOCTL codes. These are the values actually passed to
// DeviceIoControl. They are declared as consts (using the inlined
// ioctlBase | (function<<2) form) so they can be used in array sizes and
// switch cases.
const (
	// Basic group
	IOCTL_CODE_INIT_VMM                                      = ioctlBase | (ioctlBasic+0x01)<<2
	IOCTL_CODE_INIT_HYPERTRACE                               = ioctlBase | (ioctlBasic+0x02)<<2
	IOCTL_CODE_REGISTER_EVENT                                = ioctlBase | (ioctlBasic+0x03)<<2
	IOCTL_CODE_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL = ioctlBase | (ioctlBasic+0x04)<<2

	// VMM group
	IOCTL_CODE_TERMINATE_VMX                                        = ioctlBase | (ioctlVMM+0x01)<<2
	IOCTL_CODE_DEBUGGER_READ_MEMORY                                 = ioctlBase | (ioctlVMM+0x02)<<2
	IOCTL_CODE_DEBUGGER_READ_OR_WRITE_MSR                           = ioctlBase | (ioctlVMM+0x03)<<2
	IOCTL_CODE_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS             = ioctlBase | (ioctlVMM+0x04)<<2
	IOCTL_CODE_DEBUGGER_REGISTER_EVENT                              = ioctlBase | (ioctlVMM+0x05)<<2
	IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT                         = ioctlBase | (ioctlVMM+0x06)<<2
	IOCTL_CODE_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER = ioctlBase | (ioctlVMM+0x07)<<2
	IOCTL_CODE_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS                    = ioctlBase | (ioctlVMM+0x08)<<2
	IOCTL_CODE_DEBUGGER_EDIT_MEMORY                                 = ioctlBase | (ioctlVMM+0x09)<<2
	IOCTL_CODE_DEBUGGER_SEARCH_MEMORY                               = ioctlBase | (ioctlVMM+0x0a)<<2
	IOCTL_CODE_DEBUGGER_MODIFY_EVENTS                               = ioctlBase | (ioctlVMM+0x0b)<<2
	IOCTL_CODE_DEBUGGER_FLUSH_LOGGING_BUFFERS                       = ioctlBase | (ioctlVMM+0x0c)<<2
	IOCTL_CODE_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS             = ioctlBase | (ioctlVMM+0x0d)<<2
	IOCTL_CODE_DEBUGGER_PRINT                                       = ioctlBase | (ioctlVMM+0x0e)<<2
	IOCTL_CODE_PREPARE_DEBUGGEE                                     = ioctlBase | (ioctlVMM+0x0f)<<2
	IOCTL_CODE_PAUSE_PACKET_RECEIVED                                = ioctlBase | (ioctlVMM+0x10)<<2
	IOCTL_CODE_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED           = ioctlBase | (ioctlVMM+0x11)<<2
	IOCTL_CODE_SEND_USERMODE_MESSAGES_TO_DEBUGGER                   = ioctlBase | (ioctlVMM+0x12)<<2
	IOCTL_CODE_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER        = ioctlBase | (ioctlVMM+0x13)<<2
	IOCTL_CODE_PERFORM_KERNEL_SIDE_TESTS                            = ioctlBase | (ioctlVMM+0x14)<<2
	IOCTL_CODE_RESERVE_PRE_ALLOCATED_POOLS                          = ioctlBase | (ioctlVMM+0x15)<<2
	IOCTL_CODE_SEND_USER_DEBUGGER_COMMANDS                          = ioctlBase | (ioctlVMM+0x16)<<2
	IOCTL_CODE_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES           = ioctlBase | (ioctlVMM+0x17)<<2
	IOCTL_CODE_GET_USER_MODE_MODULE_DETAILS                         = ioctlBase | (ioctlVMM+0x18)<<2
	IOCTL_CODE_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS           = ioctlBase | (ioctlVMM+0x19)<<2
	IOCTL_CODE_GET_LIST_OF_THREADS_AND_PROCESSES                    = ioctlBase | (ioctlVMM+0x1a)<<2
	IOCTL_CODE_QUERY_CURRENT_PROCESS                                = ioctlBase | (ioctlVMM+0x1b)<<2
	IOCTL_CODE_QUERY_CURRENT_THREAD                                 = ioctlBase | (ioctlVMM+0x1c)<<2
	IOCTL_CODE_REQUEST_REV_MACHINE_SERVICE                          = ioctlBase | (ioctlVMM+0x1d)<<2
	IOCTL_CODE_DEBUGGER_BRING_PAGES_IN                              = ioctlBase | (ioctlVMM+0x1e)<<2
	IOCTL_CODE_PREACTIVATE_FUNCTIONALITY                            = ioctlBase | (ioctlVMM+0x1f)<<2
	IOCTL_CODE_PCIE_ENDPOINT_ENUM                                   = ioctlBase | (ioctlVMM+0x20)<<2
	IOCTL_CODE_PERFORM_ACTIONS_ON_APIC                              = ioctlBase | (ioctlVMM+0x21)<<2
	IOCTL_CODE_PCIDEVINFO_ENUM                                      = ioctlBase | (ioctlVMM+0x22)<<2
	IOCTL_CODE_QUERY_IDT_ENTRY                                      = ioctlBase | (ioctlVMM+0x24)<<2
	IOCTL_CODE_SET_BREAKPOINT_USER_DEBUGGER                         = ioctlBase | (ioctlVMM+0x25)<<2
	IOCTL_CODE_PERFORM_SMI_OPERATION                                = ioctlBase | (ioctlVMM+0x26)<<2
	IOCTL_CODE_DEBUGGER_CPUID                                       = ioctlBase | (ioctlVMM+0x27)<<2

	// HyperTrace group
	IOCTL_CODE_PERFORM_HYPERTRACE_UNLOAD        = ioctlBase | (ioctlHyperTrace+0x01)<<2
	IOCTL_CODE_PERFORM_HYPERTRACE_LBR_OPERATION = ioctlBase | (ioctlHyperTrace+0x02)<<2
	IOCTL_CODE_PERFORM_HYPERTRACE_LBR_DUMP      = ioctlBase | (ioctlHyperTrace+0x03)<<2
	IOCTL_CODE_PERFORM_HYPERTRACE_PT_OPERATION  = ioctlBase | (ioctlHyperTrace+0x04)<<2
	IOCTL_CODE_PERFORM_HYPERTRACE_PT_MMAP       = ioctlBase | (ioctlHyperTrace+0x05)<<2
)

// Device path constants. HyperDbg's VMM driver exposes its device at
// \\.\HyperDbgDebuggerDevice (see libhyperdbg/code/app/libhyperdbg.cpp and
// packets.cpp). The named-pipe endpoints used for remote/event forwarding
// live in the namedpipe sub-package.
const (
	// DeviceName is the Win32 device path for the HyperDbg VMM driver.
	DeviceName = `\\.\HyperDbgDebuggerDevice`

	// HyperDbgPipe is the default named pipe for remote debugging sessions.
	HyperDbgPipe = `\\.\pipe\HyperDbgPipe`

	// HyperDbgOutputPipe is the named pipe used for output forwarding.
	HyperDbgOutputPipe = `\\.\Pipe\HyperDbgOutput`

	// HyperDbgTestsPipe is the named pipe used by the test harness.
	HyperDbgTestsPipe = `\\.\Pipe\HyperDbgTests`
)
