package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/common"
	"github.com/ddkwork/HyperDbg/debugger/driver_loader"
	"github.com/ddkwork/HyperDbg/debugger/user_level"
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

const (
	KernelDebuggerDriverName            = "hprdbghv"
	KernelDebuggerDriverNameAndExt      = "hprdbghv.sys"
	CommunicationBufferSize             = 0x100000
	TcpEndOfBufferCharsCount            = 4
	DefaultSpeedOfReadingKernelMessages = 5
	UsermodeBufferSize                  = 0x100000
	MaxPath                             = 260
)

type (
	SendMessageCallback                 func(message string)
	SendMessageWithSharedBufferCallback func()
)

type Libhyperdbg struct {
	deviceHandle                      windows.Handle
	isDriverLoadedSuccessfully        windows.Handle
	isVmxOffProcessStart              bool
	isDebuggerModulesLoaded           bool
	privilegesAlreadyAdjusted         bool
	messageHandler                    SendMessageCallback
	messageHandlerSharedBuffer        []byte
	driverLocation                    string
	driverName                        string
	useCustomDriverLocation           bool
	isConnectedToRemoteDebugger       bool
	isSerialConnectedToRemoteDebugger bool
	logOpened                         bool
	breakPrintingOutput               bool
	mu                                sync.Mutex
}

func NewLibhyperdbg() *Libhyperdbg {
	return &Libhyperdbg{
		deviceHandle:               windows.InvalidHandle,
		isDriverLoadedSuccessfully: windows.InvalidHandle,
		driverName:                 KernelDebuggerDriverName,
	}
}

func SetDebugPrivilege() bool {
	if privilegesAdjusted {
		return true
	}
	adjustTokenPrivileges()
	privilegesAdjusted = true
	return true
}

func adjustTokenPrivileges() {
	var token windows.Token
	currentProcess, _ := windows.GetCurrentProcess()
	windows.OpenProcessToken(currentProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token)
	defer token.Close()

	var luid windows.LUID
	windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr("SeDebugPrivilege"), &luid)

	privileges := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges: [1]windows.LUIDAndAttributes{
			{Luid: luid, Attributes: windows.SE_PRIVILEGE_ENABLED},
		},
	}
	windows.AdjustTokenPrivileges(token, false, &privileges, 0, nil, nil)
}

// SetTextMessageCallback sets the function callback for showing messages
func (l *Libhyperdbg) SetTextMessageCallback(handler SendMessageCallback) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messageHandler = handler
}

// SetTextMessageCallbackUsingSharedBuffer sets the callback using shared buffer method
func (l *Libhyperdbg) SetTextMessageCallbackUsingSharedBuffer(handler SendMessageCallback) []byte {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.messageHandler = handler
	l.messageHandlerSharedBuffer = make([]byte, CommunicationBufferSize+TcpEndOfBufferCharsCount)

	if l.messageHandlerSharedBuffer == nil {
		l.messageHandler = nil
		return nil
	}

	return l.messageHandlerSharedBuffer
}

// UnsetTextMessageCallback unsets the message callback and frees the shared buffer
func (l *Libhyperdbg) UnsetTextMessageCallback() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.messageHandler = nil
	l.messageHandlerSharedBuffer = nil
}

// ShowMessages is the core message display function
func (l *Libhyperdbg) ShowMessages(format string, args ...any) {
	message := fmt.Sprintf(format, args...)

	if l.messageHandler == nil && !l.isConnectedToRemoteDebugger && !l.isSerialConnectedToRemoteDebugger {
		fmt.Print(message)
		if !l.logOpened {
			return
		}
	}

	if l.isConnectedToRemoteDebugger {
		// Send to remote debugger host - placeholder for remote connection
	}

	if l.isSerialConnectedToRemoteDebugger {
		// Send through serial connection - placeholder
	}

	if l.logOpened {
		// Write to log file - placeholder
	}

	if l.messageHandler != nil {
		if l.messageHandlerSharedBuffer == nil {
			l.messageHandler(message)
		} else {
			copy(l.messageHandlerSharedBuffer, []byte(message))
			l.messageHandler(string(l.messageHandlerSharedBuffer))
		}
	}
}

// HyperDbgInstallVmmDriver installs the VMM driver
func (l *Libhyperdbg) HyperDbgInstallVmmDriver() int {
	if !l.useCustomDriverLocation {
		driverLoc, ok := SetupPathForFileName(KernelDebuggerDriverNameAndExt, true)
		if !ok {
			return 1
		}
		l.driverLocation = driverLoc
		l.driverName = KernelDebuggerDriverName
	}

	if !driver_loader.ManageDriver(l.driverName, l.driverLocation, driver_loader.DriverFuncInstall) {
		mylog.Warning("unable to install VMM driver")
		driver_loader.ManageDriver(l.driverName, l.driverLocation, driver_loader.DriverFuncRemove)
		return 1
	}

	return 0
}

// HyperDbgStopDriver stops the specified driver
func (l *Libhyperdbg) HyperDbgStopDriver(driverName string) int {
	if l.driverLocation != "" && driver_loader.ManageDriver(driverName, l.driverLocation, driver_loader.DriverFuncStop) {
		return 0
	}
	return 1
}

// HyperDbgStopVmmDriver stops the VMM driver
func (l *Libhyperdbg) HyperDbgStopVmmDriver() int {
	return l.HyperDbgStopDriver(l.driverName)
}

// HyperDbgUninstallDriver uninstalls the specified driver
func (l *Libhyperdbg) HyperDbgUninstallDriver(driverName string) int {
	if l.driverLocation != "" && driver_loader.ManageDriver(driverName, l.driverLocation, driver_loader.DriverFuncRemove) {
		return 0
	}
	return 1
}

// HyperDbgUninstallVmmDriver uninstalls the VMM driver
func (l *Libhyperdbg) HyperDbgUninstallVmmDriver() int {
	return l.HyperDbgUninstallDriver(l.driverName)
}

// SetupPathForFileName sets the path for a driver filename
func SetupPathForFileName(fileName string, checkFileExists bool) (string, bool) {
	exePath, err := os.Executable()
	if err != nil {
		mylog.Warning("unable to get executable path", "error", err)
		return "", false
	}

	dir := filepath.Dir(exePath)
	fullPath := filepath.Join(dir, fileName)

	if checkFileExists {
		fileNamePtr, err2 := windows.UTF16PtrFromString(fullPath)
		if err2 != nil {
			mylog.Warning("unable to convert path", "error", err2)
			return "", false
		}

		fileHandle, err2 := windows.CreateFile(
			fileNamePtr,
			windows.GENERIC_READ,
			0,
			nil,
			windows.OPEN_EXISTING,
			windows.FILE_ATTRIBUTE_NORMAL,
			0,
		)
		if fileHandle == windows.InvalidHandle {
			mylog.Warning("target file is not loaded", "error", err2)
			return "", false
		}
		windows.CloseHandle(fileHandle)
	}

	return fullPath, true
}

// SetCustomDriverPath sets custom driver path and name
func (l *Libhyperdbg) SetCustomDriverPath(driverFilePath, driverName string) bool {
	if len(driverFilePath) > MaxPath {
		mylog.Warning("the driver path is too long, the maximum length is", "max", MaxPath)
		return false
	}
	if len(driverName) > MaxPath {
		mylog.Warning("the driver name is too long, the maximum length is", "max", MaxPath)
		return false
	}

	l.driverLocation = driverFilePath
	l.driverName = driverName
	l.useCustomDriverLocation = true
	return true
}

// UseDefaultDriverPath resets to using default driver path
func (l *Libhyperdbg) UseDefaultDriverPath() {
	l.useCustomDriverLocation = false
}

// IsDriverLoaded returns whether the driver is loaded
func (l *Libhyperdbg) IsDriverLoaded() bool {
	return l.isDebuggerModulesLoaded
}

// GetDeviceHandle returns the device handle
func (l *Libhyperdbg) GetDeviceHandle() windows.Handle {
	return l.deviceHandle
}

func (l *Libhyperdbg) LoadVmmModule(deviceName string) int {
	if !SetDebugPrivilege() {
		mylog.Warning("couldn't set debug privilege")
		return 1
	}

	vendorString := common.CpuReadVendorString()
	mylog.Info("current processor vendor is : " + string(vendorString[:12]))

	vendorStr := string(vendorString[:12])
	switch vendorStr {
	case "GenuineIntel":
		mylog.Info("virtualization technology is vt-x")
	default:
		mylog.Warning("this program is not designed to run in a non-VT-x environment !")
		return 1
	}

	if common.VmxSupportDetection() {
		mylog.Info("vmx operation is supported by your processor")
	} else {
		mylog.Warning("vmx operation is not supported by your processor")
		return 1
	}

	eventHandle, err := windows.CreateEvent(nil, 1, 0, nil)
	if err != nil {
		mylog.Warning("CreateEvent failed", "error", err)
		return 1
	}
	l.isDriverLoadedSuccessfully = eventHandle

	if err := l.CreateHandleFromVmmModule(deviceName); err != nil {
		windows.CloseHandle(eventHandle)
		return 1
	}

	ev, _ := windows.WaitForSingleObject(l.isDriverLoadedSuccessfully, windows.INFINITE)
	windows.CloseHandle(l.isDriverLoadedSuccessfully)
	if ev != windows.WAIT_OBJECT_0 {
		return 1
	}

	l.isDebuggerModulesLoaded = true
	mylog.Success("vmm module is running...")
	return 0
}

func (l *Libhyperdbg) CreateHandleFromVmmModule(deviceName string) error {
	if l.deviceHandle != windows.InvalidHandle {
		mylog.Warning("handle of the driver found, if you use 'load' before, please unload it using 'unload'")
		return windows.ERROR_ALREADY_EXISTS
	}

	l.isVmxOffProcessStart = false

	deviceNamePtr, err := windows.UTF16PtrFromString(deviceName)
	if err != nil {
		return err
	}

	handle, err := windows.CreateFile(
		deviceNamePtr,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL|windows.FILE_FLAG_OVERLAPPED,
		0,
	)
	if err != nil {
		switch err {
		case windows.ERROR_ACCESS_DENIED:
			mylog.Warning("err, access denied\nare you sure you have administrator rights?")
		case windows.ERROR_GEN_FAILURE:
			mylog.Warning("err, a device attached to the system is not functioning\nvmx feature might be disabled from BIOS or VBS/HVCI is active")
		default:
			mylog.Warning("err, CreateFile failed", "error", err)
		}
		l.deviceHandle = windows.InvalidHandle
		return err
	}

	l.deviceHandle = handle
	return nil
}

func (l *Libhyperdbg) UnloadVmm(driverSendReceive func(buffer []byte, ioctlCode uint32) ([]byte, error)) int {
	if l.deviceHandle == windows.InvalidHandle {
		mylog.Warning("driver handle is invalid")
		return 1
	}

	mylog.Info("start terminating...")

	UdUninitializeUserDebugger()

	driverSendReceive(nil, IoctlTerminateVmx)
	driverSendReceive(nil, IoctlReturnIrpPendingPacketsAndDisallowIoctl)

	l.isVmxOffProcessStart = true
	time.Sleep(1 * time.Second)

	if err := windows.CloseHandle(l.deviceHandle); err != nil {
		mylog.Warning("err, closing handle", "error", err)
		return 1
	}

	l.deviceHandle = windows.InvalidHandle
	l.isDebuggerModulesLoaded = false

	mylog.Info("you're not on HyperDbg's hypervisor anymore!")
	return 0
}

var (
	IoctlTerminateVmx                            uint32 = 0x00222008
	IoctlReturnIrpPendingPacketsAndDisallowIoctl uint32 = 0x00222004
	privilegesAdjusted                           bool
)

func UdUninitializeUserDebugger() {
	user_level.UdUninitializeUserDebugger()
}
