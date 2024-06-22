package HPRDBGCTRL

import (
	"fmt"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ddkwork/app/ms/hardwareIndo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream/bitfield"

	"github.com/winlabs/gowin32/wrappers"
)

func CTL_CODE(deviceType, function, method, access uint32) IoctlsKind {
	return IoctlsKind(((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method))
}

const (
	FILE_DEVICE_UNKNOWN = wrappers.FILE_DEVICE_UNKNOWN
	METHOD_BUFFERED     = wrappers.METHOD_BUFFERED
	FILE_ANY_ACCESS     = wrappers.FILE_ANY_ACCESS
)

type IoctlsKind uint32

func (e IoctlsKind) String() string {
	switch e {
	case IOCTL_REGISTER_EVENT:
		return "IoctlRegisterEvent"
	case IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL:
		return "IoctlReturnIrpPendingPacketsAndDisallowIoctl"
	case IOCTL_TERMINATE_VMX:
		return "IoctlTerminateVmx"
	case IOCTL_DEBUGGER_READ_MEMORY:
		return "IoctlDebuggerReadMemory"
	case IOCTL_DEBUGGER_READ_OR_WRITE_MSR:
		return "IoctlDebuggerReadOrWriteMsr"
	case IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS:
		return "IoctlDebuggerReadPageTableEntriesDetails"
	case IOCTL_DEBUGGER_REGISTER_EVENT:
		return "IoctlDebuggerRegisterEvent"
	case IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT:
		return "IoctlDebuggerAddActionToEvent"
	case IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER:
		return "IoctlDebuggerHideAndUnhideToTransparentTheDebugger"
	case IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS:
		return "IoctlDebuggerVa2PaAndPa2VaCommands"
	case IOCTL_DEBUGGER_EDIT_MEMORY:
		return "IoctlDebuggerEditMemory"
	case IOCTL_DEBUGGER_SEARCH_MEMORY:
		return "IoctlDebuggerSearchMemory"
	case IOCTL_DEBUGGER_MODIFY_EVENTS:
		return "IoctlDebuggerModifyEvents"
	case IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS:
		return "IoctlDebuggerFlushLoggingBuffers"
	case IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS:
		return "IoctlDebuggerAttachDetachUserModeProcess"
	case IOCTL_DEBUGGER_PRINT:
		return "IoctlDebuggerPrint"
	case IOCTL_PREPARE_DEBUGGEE:
		return "IoctlPrepareDebuggee"
	case IOCTL_PAUSE_PACKET_RECEIVED:
		return "IoctlPausePacketReceived"
	case IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED:
		return "IoctlSendSignalExecutionInDebuggeeFinished"
	case IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER:
		return "IoctlSendUsermodeMessagesToDebugger"
	case IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER:
		return "IoctlSendGeneralBufferFromDebuggeeToDebugger"
	// case IOCTL_SEND_GET_KERNEL_SIDE_TEST_INFORMATION://removed ?
	//	return "IoctlSendGetKernelSideTestInformation"
	case IOCTL_PERFROM_KERNEL_SIDE_TESTS:
		return "IoctlPerfromKernelSideTests"
	case IOCTL_RESERVE_PRE_ALLOCATED_POOLS:
		return "IoctlReservePreAllocatedPools"
	case IOCTL_SEND_USER_DEBUGGER_COMMANDS:
		return "IoctlSendUserDebuggerCommands"
	case IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES:
		return "IoctlGetDetailOfActiveThreadsAndProcesses"
	case IOCTL_GET_USER_MODE_MODULE_DETAILS:
		return "IoctlGetUserModeModuleDetails"
	case IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS:
		return "IoctlQueryCountOfActiveProcessesOrThreads"
	case IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES:
		return "IoctlGetListOfThreadsAndProcesses"
	case IOCTL_QUERY_CURRENT_PROCESS:
		return "IoctlQueryCurrentProcess"
	case IOCTL_QUERY_CURRENT_THREAD:
		return "IoctlQueryCurrentThread"
	default:
		return fmt.Sprint("known error code " + fmt.Sprintf("%d", e))
	}
}

func LOWORD(l uint32) uint16 { return uint16(l) }
func LOBYTE(l uint32) uint8  { return byte(l) }
func HIWORD(l uint32) uint16 { return uint16(l >> 16) }
func HIBYTE(l uint32) uint8  { return byte(l >> 24) }

func TestSizeof(t *testing.T) {
	// assert.Equal(t, 11, binary.Size(DEBUGGER_REMOTE_PACKET{}))
}

func TestHIBYTE(t *testing.T) {
	v := uint32(0x11223344)
	assert.Equal(t, byte(0x11), HIBYTE(v))
	assert.Equal(t, uint16(0x1122), HIWORD(v))
	assert.Equal(t, byte(0x44), LOBYTE(v))
	assert.Equal(t, uint16(0x3344), LOWORD(v))
}

func VmxSupportDetection() (ok bool) {
	hard := hardwareIndo.New()
	if !hard.CpuInfo.Get() {
		return
	}
	if hard.CpuInfo.Vendor != "GenuineIntel" {
		mylog.Check("this program is not designed to run in a non-VT-x environemnt !")
	}
	mylog.Info("", "virtualization technology is vt-x")
	field := bitfield.NewFromUint32(hard.CpuInfo.Cpu1.Ecx)
	if !field.Test(5) {
		mylog.Check("vmx operation is not supported by your processor")
	}
	mylog.Info("", "vmx operation is supported by your processor")
	return true
}

func DeviceName() string { return "HyperdbgHypervisorDevice" }
func LinkName() (*uint16, error) {
	return syscall.UTF16PtrFromString(`\\\\.\\` + DeviceName())
}