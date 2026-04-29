package driver_loader

import (
	"time"

	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

const (
	DriverFuncStop DriverFunc = iota
	DriverFuncRemove
	DriverFuncStart
	DriverFuncInstall
)

type DriverFunc int

func ManageDriver(driverName, driverPath string, funcType DriverFunc) bool {
	scManager, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		mylog.Warning("OpenSCManager failed", "error", err)
		return false
	}
	defer windows.CloseServiceHandle(scManager)

	switch funcType {
	case DriverFuncStop:
		return StopDriver(scManager, driverName)
	case DriverFuncRemove:
		return RemoveDriver(scManager, driverName)
	case DriverFuncStart:
		return StartDriver(scManager, driverName)
	case DriverFuncInstall:
		return InstallDriver(scManager, driverName, driverPath)
	default:
		return false
	}
}

func InstallDriver(scManager windows.Handle, driverName, serviceExe string) bool {
	driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(driverName))
	serviceExePtr := mylog.Check2(windows.UTF16PtrFromString(serviceExe))

	schService, err := windows.CreateService(
		scManager,
		driverNamePtr,
		driverNamePtr,
		windows.SERVICE_ALL_ACCESS,
		windows.SERVICE_KERNEL_DRIVER,
		windows.SERVICE_DEMAND_START,
		windows.SERVICE_ERROR_NORMAL,
		serviceExePtr,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		if err == windows.ERROR_SERVICE_EXISTS {
			mylog.Info("the service (driver) already exists, trying to remove old instance")

			StopDriver(scManager, driverName)
			time.Sleep(1 * time.Second)
			RemoveDriver(scManager, driverName)
			time.Sleep(1 * time.Second)

			return InstallDriver(scManager, driverName, serviceExe)
		} else if err == windows.ERROR_SERVICE_MARKED_FOR_DELETE {
			mylog.Warning("previous instance of the service is not fully deleted. Try again...")
			return false
		} else {
			mylog.Warning("CreateService failed", "error", err)
			return false
		}
	}

	if schService != 0 {
		windows.CloseServiceHandle(schService)
	}

	mylog.Success("driver installed successfully")
	return true
}

func RemoveDriver(scManager windows.Handle, driverName string) bool {
	driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(driverName))

	schService, err := windows.OpenService(scManager, driverNamePtr, windows.SERVICE_ALL_ACCESS)
	if err != nil {
		if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
			return true
		}
		mylog.Warning("OpenService failed in remove", "error", err)
		return false
	}
	defer windows.CloseServiceHandle(schService)

	err = windows.DeleteService(schService)
	if err != nil && err != windows.ERROR_SERVICE_MARKED_FOR_DELETE {
		mylog.Warning("DeleteService failed", "error", err)
		return false
	}

	mylog.Success("driver removed successfully")
	return true
}

func StartDriver(scManager windows.Handle, driverName string) bool {
	driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(driverName))

	schService, err := windows.OpenService(scManager, driverNamePtr, windows.SERVICE_ALL_ACCESS)
	if err != nil {
		if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
			mylog.Info("service does not exist, trying to install")
			return false
		}
		mylog.Warning("OpenService failed in start", "error", err)
		return false
	}
	defer windows.CloseServiceHandle(schService)

	err = windows.StartService(schService, 0, nil)
	if err != nil {
		if err == windows.ERROR_SERVICE_ALREADY_RUNNING {
			mylog.Info("service is already running")
		} else {
			mylog.Warning("StartService failed", "error", err)
			return false
		}
	} else {
		mylog.Success("driver started successfully")
	}
	return true
}

func StopDriver(scManager windows.Handle, driverName string) bool {
	driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(driverName))

	schService, err := windows.OpenService(scManager, driverNamePtr, windows.SERVICE_ALL_ACCESS)
	if err != nil {
		if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
			return true
		}
		mylog.Warning("OpenService failed in stop", "error", err)
		return false
	}
	defer windows.CloseServiceHandle(schService)

	var serviceStatus windows.SERVICE_STATUS
	err = windows.ControlService(schService, windows.SERVICE_CONTROL_STOP, &serviceStatus)
	if err != nil {
		if err.Error() == "The service has not been started." {
			mylog.Info("service is not started, no need to stop")
		} else {
			mylog.Warning("ControlService failed", "error", err)
			return false
		}
	} else {
		mylog.Success("driver stopped successfully")
	}
	return true
}

// SetupPathForFileName builds the full path for a driver file by locating it
// relative to the current executable. If checkFileExists is true, it verifies
// the file is accessible before returning.
func SetupPathForFileName(fileName string, checkFileExists bool) (string, bool) {
	var buf [windows.MAX_PATH]uint16
	_, err := windows.GetModuleFileName(0, &buf[0], windows.MAX_PATH)
	if err != nil {
		mylog.Warning("unable to get module file name", "error", err)
		return "", false
	}

	exePath := windows.UTF16ToString(buf[:])

	// Find last backslash and truncate to directory
	lastSlash := -1
	for i := 0; i < len(exePath); i++ {
		if exePath[i] == '\\' {
			lastSlash = i
		}
	}
	if lastSlash >= 0 {
		exePath = exePath[:lastSlash]
	}

	// Build full path
	fullPath := exePath + "\\" + fileName

	if checkFileExists {
		fileNamePtr, err := windows.UTF16PtrFromString(fullPath)
		if err != nil {
			mylog.Warning("unable to convert path", "error", err)
			return "", false
		}

		fileHandle, err := windows.CreateFile(
			fileNamePtr,
			windows.GENERIC_READ,
			0,
			nil,
			windows.OPEN_EXISTING,
			windows.FILE_ATTRIBUTE_NORMAL,
			0,
		)
		if fileHandle == windows.InvalidHandle {
			mylog.Warning("target file is not loaded")
			return "", false
		}
		windows.CloseHandle(fileHandle)
	}

	return fullPath, true
}
