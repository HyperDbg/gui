package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	g_DeviceHandleReversingMachine    syscall.Handle
	g_DeviceHandle                    syscall.Handle
	g_IsDriverLoadedSuccessfully      syscall.Handle
	g_IsReversingMachineModulesLoaded bool
	g_IsVmxOffProcessStart            bool
)

func ShowMessages(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func ReadIrpBasedBuffer() {
	var status bool
	var returnedLength uint32
	var registerEvent struct {
		hEvent *syscall.Handle
		Type   uint32
	}
	registerEvent.Type = IRP_BASED

	handle, err := syscall.CreateFile(
		syscall.StringToUTF16Ptr("\\\\.\\HyperDbgReversingMachineDevice"),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
		0)

	if err != nil {
		errorNum := err.(syscall.Errno)
		if errorNum == syscall.ERROR_ACCESS_DENIED {
			ShowMessages("err, access denied\nare you sure you have administrator rights?\n")
		} else if errorNum == syscall.ERROR_GEN_FAILURE {
			ShowMessages("err, a device attached to the system is not functioning\nvmx feature might be disabled from BIOS or VBS/HVCI is active\n")
		} else {
			ShowMessages("err, CreateFile failed with (%x)\n", errorNum)
		}
		g_DeviceHandle = 0
		return
	}

	outputBuffer := make([]byte, UsermodeBufferSize)

	defer syscall.CloseHandle(handle)
	for {
		time.Sleep(DefaultSpeedOfReadingKernelMessages * time.Millisecond)

		status, _, err = syscall.DeviceIoControl(
			handle,
			IOCTL_REGISTER_EVENT,
			(*byte)(unsafe.Pointer(&registerEvent)),
			uint32(unsafe.Sizeof(registerEvent)*2),
			(*byte)(unsafe.Pointer(&outputBuffer[0])),
			uint32(len(outputBuffer)),
			&returnedLength,
			nil,
		)

		if !status {
			continue
		}

		operationCode := *(*uint32)(unsafe.Pointer(&outputBuffer[0]))

		switch operationCode {
		case OPERATION_LOG_NON_IMMEDIATE_MESSAGE,
			OPERATION_LOG_INFO_MESSAGE,
			OPERATION_LOG_ERROR_MESSAGE,
			OPERATION_LOG_WARNING_MESSAGE:
			ShowMessages("%s", outputBuffer[4:returnedLength])
		case OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED:
			syscall.SetEvent(g_IsDriverLoadedSuccessfully)
		default:
			ShowMessages("err, unknown message is received in rev module\n")
		}
	}
}

func IrpBasedBufferThread(data interface{}) uint32 {
	ReadIrpBasedBuffer()
	return 0
}

func ReversingMachineStop() int {
	return 0
}

func ReversingMachineStart() int {
	if g_DeviceHandle != 0 {
		ShowMessages("handle of the driver found, if you use 'load' before, please unload it using 'unload'\n")
		return 1
	}

	g_IsVmxOffProcessStart = false

	handle, err := syscall.CreateFile(
		syscall.StringToUTF16Ptr("\\\\.\\HyperDbgReversingMachineDevice"),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
		0)

	if err != nil {
		errorNum := err.(syscall.Errno)
		if errorNum == syscall.ERROR_ACCESS_DENIED {
			ShowMessages("err, access denied\nare you sure you have administrator rights?\n")
		} else if errorNum == syscall.ERROR_GEN_FAILURE {
			ShowMessages("err, a device attached to the system is not functioning\nvmx feature might be disabled from BIOS or VBS/HVCI is active\n")
		} else {
			ShowMessages("err, CreateFile failed (%x)\n", errorNum)
		}
		g_DeviceHandle = 0
		return 1
	}

	g_DeviceHandle = handle

	event, err := syscall.CreateEvent(nil, 0, 0, nil)
	if err != nil {
		ShowMessages("err, unable to create event\n")
		syscall.CloseHandle(handle)
		return 1
	}
	g_IsDriverLoadedSuccessfully = event

	thread, err := syscall.CreateThread(nil, 0, syscall.NewCallback(IrpBasedBufferThread), nil, 0, nil)
	if err != nil {
		ShowMessages("err, unable to create thread\n")
		syscall.CloseHandle(handle)
		syscall.CloseHandle(event)
		return 1
	}
	syscall.CloseHandle(thread)

	syscall.WaitForSingleObject(g_IsDriverLoadedSuccessfully, syscall.INFINITE)

	syscall.CloseHandle(g_IsDriverLoadedSuccessfully)

	return 0
}

func HyperDbgUnloadReversingMachine() int {
	if !g_IsReversingMachineModulesLoaded {
		ShowMessages("handle of the driver not found, probably the driver is not loaded. Did you use 'load' command?\n")
		return 1
	}

	ShowMessages("start terminating...\n")
	ReversingMachineStop()
	ShowMessages("you're not on reversing machine's hypervisor anymore!\n")
	return 0
}

func HyperDbgLoadReversingMachine() int {
	var cpuId [13]byte

	if g_DeviceHandle != 0 {
		ShowMessages("handle of the driver found, if you use 'load' before, please unload it using 'unload'\n")
		return 1
	}

	CpuReadVendorString(cpuId[:])

	ShowMessages("current processor vendor is : %s\n", cpuId)

	if string(cpuId[:]) == "GenuineIntel" {
		ShowMessages("virtualization technology is vt-x\n")
	} else {
		ShowMessages("this program is not designed to run in a non-VT-x environment !\n")
		return 1
	}

	if VmxSupportDetection() {
		ShowMessages("vmx operation is supported by your processor\n")
	} else {
		ShowMessages("vmx operation is not supported by your processor\n")
		return 1
	}

	return ReversingMachineStart()
}

func main() {
	if HyperDbgLoadReversingMachine() == 0 {
		HyperDbgUnloadReversingMachine()
	} else {
		ShowMessages("err, in loading HyperDbg\n")
	}
}
