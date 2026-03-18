package driver

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

const (
	DriverFuncInstall = iota
	DriverFuncStop
	DriverFuncRemove
)

type DriverProvider struct {
	serviceName string
	driverPath  string
}

func NewDriverProvider(serviceName, driverPath string) *DriverProvider {
	return &DriverProvider{
		serviceName: serviceName,
		driverPath:  driverPath,
	}
}

func (dp *DriverProvider) Install() error {
	scm, err := dp.OpenSCManager()
	if err != nil {
		return fmt.Errorf("failed to open service control manager: %w", err)
	}
	defer dp.CloseSCManager(scm)

	err = dp.InstallDriver(scm)
	if err != nil {
		return fmt.Errorf("failed to install driver: %w", err)
	}

	err = dp.StartDriver(scm)
	if err != nil {
		return fmt.Errorf("failed to start driver: %w", err)
	}

	return nil
}

func (dp *DriverProvider) Uninstall() error {
	scm, err := dp.OpenSCManager()
	if err != nil {
		return fmt.Errorf("failed to open service control manager: %w", err)
	}
	defer dp.CloseSCManager(scm)

	err = dp.StopDriver(scm)
	if err != nil {
		return fmt.Errorf("failed to stop driver: %w", err)
	}

	err = dp.RemoveDriver(scm)
	if err != nil {
		return fmt.Errorf("failed to remove driver: %w", err)
	}

	return nil
}

func (dp *DriverProvider) Manage(function int) error {
	scm, err := dp.OpenSCManager()
	if err != nil {
		return fmt.Errorf("failed to open service control manager: %w", err)
	}
	defer dp.CloseSCManager(scm)

	switch function {
	case DriverFuncInstall:
		err = dp.InstallDriver(scm)
		if err != nil {
			return err
		}
		return dp.StartDriver(scm)
	case DriverFuncStop:
		return dp.StopDriver(scm)
	case DriverFuncRemove:
		return dp.RemoveDriver(scm)
	default:
		return fmt.Errorf("unknown function: %d", function)
	}
}

func (dp *DriverProvider) OpenSCManager() (syscall.Handle, error) {
	handle, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		return syscall.InvalidHandle, fmt.Errorf("OpenSCManager failed: %w", err)
	}

	return syscall.Handle(handle), nil
}

func (dp *DriverProvider) CloseSCManager(handle syscall.Handle) error {
	if handle != syscall.InvalidHandle {
		return windows.CloseServiceHandle(windows.Handle(handle))
	}
	return nil
}

func (dp *DriverProvider) InstallDriver(scm syscall.Handle) error {
	service, err := dp.CreateService(scm)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno == windows.ERROR_SERVICE_EXISTS {
				service, err = dp.OpenService(scm)
				if err != nil {
					return fmt.Errorf("failed to open existing service: %w", err)
				}

				err = dp.StopDriverService(service)
				if err != nil {
					windows.CloseServiceHandle(service)
					return fmt.Errorf("failed to stop existing service: %w", err)
				}

				err = dp.RemoveDriverService(service)
				if err != nil {
					windows.CloseServiceHandle(service)
					return fmt.Errorf("failed to remove existing service: %w", err)
				}

				windows.CloseServiceHandle(service)

				time.Sleep(2 * time.Second)

				service, err = dp.CreateService(scm)
				if err != nil {
					if errno, ok := err.(syscall.Errno); ok {
						if errno == windows.ERROR_SERVICE_MARKED_FOR_DELETE {
							return fmt.Errorf("service is marked for deletion but not fully deleted yet")
						}
					}
					return fmt.Errorf("failed to create service after cleanup: %w", err)
				}
			} else {
				return fmt.Errorf("CreateService failed: %w", err)
			}
		} else {
			return fmt.Errorf("CreateService failed: %w", err)
		}
	}

	if service != windows.Handle(0) {
		windows.CloseServiceHandle(service)
	}

	return nil
}

func (dp *DriverProvider) StartExistingDriver(scm syscall.Handle) error {
	service, err := dp.OpenService(scm)
	if err != nil {
		return err
	}

	err = dp.EnsureServiceStopped(service)
	if err != nil {
		windows.CloseServiceHandle(service)
		return fmt.Errorf("failed to stop existing driver: %w", err)
	}

	driverPathPtr, err := syscall.UTF16PtrFromString(dp.driverPath)
	if err != nil {
		windows.CloseServiceHandle(service)
		return fmt.Errorf("failed to convert driver path: %w", err)
	}

	err = windows.ChangeServiceConfig(
		service,
		windows.SERVICE_KERNEL_DRIVER,
		windows.SERVICE_DEMAND_START,
		windows.SERVICE_ERROR_NORMAL,
		driverPathPtr,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		windows.CloseServiceHandle(service)
		return fmt.Errorf("ChangeServiceConfig failed: %w", err)
	}

	defer windows.CloseServiceHandle(service)
	return dp.StartDriverService(service)
}

func (dp *DriverProvider) OpenService(scm syscall.Handle) (windows.Handle, error) {
	serviceName, err := syscall.UTF16PtrFromString(dp.serviceName)
	if err != nil {
		return 0, fmt.Errorf("failed to convert service name: %w", err)
	}

	service, err := windows.OpenService(windows.Handle(scm), serviceName, windows.SERVICE_ALL_ACCESS)
	if err != nil {
		return 0, fmt.Errorf("OpenService failed: %w", err)
	}

	return service, nil
}

func (dp *DriverProvider) EnsureServiceStopped(service windows.Handle) error {
	var status windows.SERVICE_STATUS
	err := windows.QueryServiceStatus(service, &status)
	if err != nil {
		return fmt.Errorf("QueryServiceStatus failed: %w", err)
	}

	if status.CurrentState != windows.SERVICE_STOPPED {
		return dp.StopDriverService(service)
	}

	return nil
}

func (dp *DriverProvider) CreateService(scm syscall.Handle) (windows.Handle, error) {
	serviceName, err := syscall.UTF16PtrFromString(dp.serviceName)
	if err != nil {
		return 0, fmt.Errorf("failed to convert service name: %w", err)
	}

	driverPathPtr, err := syscall.UTF16PtrFromString(dp.driverPath)
	if err != nil {
		return 0, fmt.Errorf("failed to convert driver path: %w", err)
	}

	service, err := windows.CreateService(
		windows.Handle(scm),
		serviceName,
		serviceName,
		windows.SERVICE_ALL_ACCESS,
		windows.SERVICE_KERNEL_DRIVER,
		windows.SERVICE_DEMAND_START,
		windows.SERVICE_ERROR_NORMAL,
		driverPathPtr,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno == windows.ERROR_SERVICE_EXISTS {
				return 0, fmt.Errorf("CreateService failed: %w. Service already exists", err)
			}
		}
		return 0, fmt.Errorf("CreateService failed: %w. Main reasons: 1) Driver path incorrect (path: %s), 2) Insufficient privileges (need Administrator)", err, dp.driverPath)
	}
	return service, nil
}

func (dp *DriverProvider) RemoveDriver(scm syscall.Handle) error {
	service, err := dp.OpenService(scm)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(service)

	return dp.RemoveDriverService(service)
}

func (dp *DriverProvider) RemoveDriverService(service windows.Handle) error {
	err := windows.DeleteService(service)
	if err != nil {
		return fmt.Errorf("DeleteService failed: %w", err)
	}
	return nil
}

func (dp *DriverProvider) StartDriver(scm syscall.Handle) error {
	service, err := dp.OpenService(scm)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(service)

	return dp.StartDriverService(service)
}

func (dp *DriverProvider) StartDriverService(service windows.Handle) error {
	err := windows.StartService(service, 0, nil)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno == windows.ERROR_SERVICE_ALREADY_RUNNING {
				return nil
			}
			if errno == windows.ERROR_FILE_NOT_FOUND || errno == windows.ERROR_PATH_NOT_FOUND {
				return fmt.Errorf("StartService failed: driver path incorrect or file not found. Error code: %d (0x%x)", uint32(errno), uint32(errno))
			}
			if errno == 0x800700C1 {
				return fmt.Errorf("StartService failed: driver signature verification failed. Error code 0x800700C1 (ERROR_INVALID_IMAGE_HASH)")
			}
		}
		return fmt.Errorf("StartService failed: %w", err)
	}
	return nil
}

func (dp *DriverProvider) StopDriver(scm syscall.Handle) error {
	serviceName, err := syscall.UTF16PtrFromString(dp.serviceName)
	if err != nil {
		return fmt.Errorf("failed to convert service name: %w", err)
	}

	service, err := windows.OpenService(windows.Handle(scm), serviceName, windows.SERVICE_ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("OpenService failed: %w", err)
	}
	defer windows.CloseServiceHandle(service)

	return dp.StopDriverService(service)
}

func (dp *DriverProvider) StopDriverService(service windows.Handle) error {
	var serviceStatus windows.SERVICE_STATUS
	err := windows.ControlService(service, windows.SERVICE_CONTROL_STOP, &serviceStatus)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno == windows.ERROR_SERVICE_NOT_ACTIVE {
				return nil
			}
		}
		return fmt.Errorf("ControlService failed: %w", err)
	}

	for range 10 {
		err = windows.QueryServiceStatus(service, &serviceStatus)
		if err != nil {
			return fmt.Errorf("QueryServiceStatus failed: %w", err)
		}

		if serviceStatus.CurrentState == windows.SERVICE_STOPPED {
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return fmt.Errorf("service did not stop within timeout")
}

func (dp *DriverProvider) SetupDriverPath(driverFileName string, checkFileExists bool) (string, error) {
	driverPath := filepath.Join(dp.driverPath, driverFileName)

	if checkFileExists {
		driverPathPtr, err := syscall.UTF16PtrFromString(driverPath)
		if err != nil {
			return "", fmt.Errorf("failed to convert driver path: %w", err)
		}
		handle, err := syscall.CreateFile(
			driverPathPtr,
			syscall.GENERIC_READ,
			0,
			nil,
			syscall.OPEN_EXISTING,
			syscall.FILE_ATTRIBUTE_NORMAL,
			0,
		)
		if err != nil {
			return "", fmt.Errorf("driver file not found: %w", err)
		}
		syscall.CloseHandle(handle)
	}

	return driverPath, nil
}

func (dp *DriverProvider) GetDriverPath() string {
	return dp.driverPath
}

func (dp *DriverProvider) GetServiceName() string {
	return dp.serviceName
}
