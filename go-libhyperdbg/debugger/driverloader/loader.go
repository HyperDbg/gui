// Package driverloader manages HyperDbg kernel-driver service lifecycle
// (install/start/stop/remove) via the Windows Service Control Manager.
//
// It mirrors the C implementation in
// libhyperdbg/code/debugger/driver-loader/install.cpp but uses pure syscalls
// through golang.org/x/sys/windows (no cgo) and returns Go errors instead of
// BOOLEAN. The Driver type is API-compatible with ok/ddk.Driver so callers can
// switch over with minimal changes.
//
// All methods are safe for concurrent use only from a single goroutine; the
// underlying SCM handles are not guarded because the SCM workflow is strictly
// sequential (install → start → stop → remove).
package driverloader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows"
)

// Driver represents a kernel driver managed by the Service Control Manager.
//
// Path is the absolute path to the .sys file. Name is the service name
// registered with the SCM (defaults to the file basename without extension).
type Driver struct {
	Path string
	Name string
}

// NewDriver creates a Driver from the given .sys path. The service name is
// derived from the file basename (without extension), matching the C
// ManageDriver convention.
func NewDriver(driverPath string) *Driver {
	name := strings.TrimSuffix(filepath.Base(driverPath), filepath.Ext(driverPath))
	return &Driver{Path: driverPath, Name: name}
}

// withSCManager opens the SCM with ALL_ACCESS and invokes fn with the handle.
// The handle is always closed. The returned error wraps the SCM failure cause.
func (d *Driver) withSCManager(fn func(windows.Handle) error) error {
	sc, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("OpenSCManager failed: %w", err)
	}
	defer windows.CloseServiceHandle(sc)
	return fn(sc)
}

// Install registers the driver as a kernel service (SERVICE_DEMAND_START).
//
// Mirrors C++ InstallDriver (install.cpp): if the service already exists
// (ERROR_SERVICE_EXISTS), the stale service is stopped and removed first,
// then Install recurses to create a fresh registration pointing at the
// current .sys path. This avoids loading a stale driver image that may
// differ from the user-mode build — the most common cause of BSOD after a
// rebuild.
//
// Crash recovery: if a previous test/process was force-killed mid-session,
// the service is left RUNNING and DeleteService marks it
// MARKED_FOR_DELETE. SCM only purges a marked-for-delete service once it is
// STOPPED and all handles close. Install detects this state, stops the
// service (triggering driver VMX unload), waits for STOPPED, then waits for
// SCM to actually purge the service before retrying CreateService.
func (d *Driver) Install() error {
	return d.installRecursive(false)
}

// installRecursive is the recursive core of Install. The `retried` flag
// prevents infinite recursion if the second CreateService still reports
// EXISTS/MARKED_FOR_DELETE (should not happen, but guards against a
// misbehaving SCM).
func (d *Driver) installRecursive(retried bool) error {
	return d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		exePtr, err := windows.UTF16PtrFromString(d.Path)
		if err != nil {
			return err
		}
		svc, err := windows.CreateService(
			sc, namePtr, namePtr, windows.SERVICE_ALL_ACCESS,
			windows.SERVICE_KERNEL_DRIVER, windows.SERVICE_DEMAND_START,
			windows.SERVICE_ERROR_NORMAL, exePtr, nil, nil, nil, nil, nil)
		if err == nil {
			windows.CloseServiceHandle(svc)
			return nil
		}
		switch err {
		case windows.ERROR_SERVICE_MARKED_FOR_DELETE:
			// Previous run crashed and left the service marked-for-delete.
			// SCM completes deletion only after the service is STOPPED and
			// all handles close. Stop it (triggers driver VMX unload), wait
			// for STOPPED, then wait for SCM to purge the deletion before
			// retrying CreateService.
			if retried {
				return fmt.Errorf("Install: service %q stuck marked-for-delete after stop+wait (driver unload hung? reboot required)", d.Name)
			}
			_ = d.Stop()             // trigger driver unload if still running
			d.waitForServiceStopped()  // wait for STOPPED (max 15s)
			d.waitForServiceGone()     // wait for SCM to purge deletion (max 20s)
			return d.installRecursive(true)
		case windows.ERROR_SERVICE_EXISTS:
			// Stale service exists — mirrors C++: stop + remove, then recurse.
			if retried {
				// Second attempt still reports EXISTS; reuse the service
				// (fixed path guarantees the correct .sys image).
				return nil
			}
			if err := d.Stop(); err != nil {
				// Non-fatal: service may not be running.
			}
			if err := d.Remove(); err != nil {
				return fmt.Errorf("Install: failed to remove stale service %q: %w", d.Name, err)
			}
			d.waitForServiceGone() // wait for SCM to purge before retry
			return d.installRecursive(true)
		default:
			return fmt.Errorf("CreateService failed for %q: %w", d.Name, err)
		}
	})
}

// Remove deletes the service registration. A non-existent service is treated
// as success so Remove is idempotent and safe to call after a failed install.
func (d *Driver) Remove() error {
	return d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_ALL_ACCESS)
		if err != nil {
			if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
				return nil
			}
			return fmt.Errorf("OpenService failed in remove: %w", err)
		}
		defer windows.CloseServiceHandle(svc)
		if err := windows.DeleteService(svc); err != nil &&
			err != windows.ERROR_SERVICE_MARKED_FOR_DELETE {
			return fmt.Errorf("DeleteService failed: %w", err)
		}
		return nil
	})
}

// waitForServiceStopped polls QueryServiceStatus until the service reaches
// SERVICE_STOPPED (max 15s). Used by Install when a stale service is
// marked-for-delete and by Stop to ensure the driver unload routine
// completes. Unlike the previous version (which only waited for
// STOP_PENDING to clear), this waits for the actual STOPPED state so a
// still-RUNNING marked-for-delete service is properly waited on.
func (d *Driver) waitForServiceStopped() {
	_ = d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_QUERY_STATUS)
		if err != nil {
			return err
		}
		defer windows.CloseServiceHandle(svc)
		var status windows.SERVICE_STATUS
		for i := 0; i < 150; i++ { // 150 * 100ms = 15s
			if err := windows.QueryServiceStatus(svc, &status); err != nil {
				break
			}
			if status.CurrentState == windows.SERVICE_STOPPED {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	})
}

// waitForServiceGone polls OpenService until the service no longer exists
// (ERROR_SERVICE_DOES_NOT_EXIST or unopenable), up to 20s. Used after
// DeleteService to ensure SCM has fully purged the marked-for-delete
// service before CreateService is retried — otherwise the retry returns
// ERROR_SERVICE_MARKED_FOR_DELETE again (the exact bug that left the test
// unable to reload the driver after a force-stop).
func (d *Driver) waitForServiceGone() {
	_ = d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		for i := 0; i < 200; i++ { // 200 * 100ms = 20s
			svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_QUERY_STATUS)
			if err != nil {
				return nil // service gone (DOES_NOT_EXIST) or unopenable
			}
			windows.CloseServiceHandle(svc)
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	})
}

// Start starts the driver service. If the service is already running this is
// treated as success (ERROR_SERVICE_ALREADY_RUNNING).
//
// Common failure causes returned:
//   - ERROR_PATH_NOT_FOUND        — driver file missing or AV blocking access
//   - ERROR_INVALID_IMAGE_HASH    — driver signature enforcement / HVCI active
func (d *Driver) Start() error {
	return d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_ALL_ACCESS)
		if err != nil {
			return fmt.Errorf("OpenService failed in start: %w", err)
		}
		defer windows.CloseServiceHandle(svc)
		if err := windows.StartService(svc, 0, nil); err != nil {
			if err == windows.ERROR_SERVICE_ALREADY_RUNNING {
				return nil
			}
			return fmt.Errorf("StartService failed for %q: %w", d.Name, err)
		}
		return nil
	})
}

// Stop sends SERVICE_CONTROL_STOP and waits for the service to reach
// STOPPED state. VMX teardown inside the driver's unload routine can take
// several seconds; without waiting, the service lingers in STOP_PENDING,
// causing the next Install to fail with MARKED_FOR_DELETE.
func (d *Driver) Stop() error {
	return d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_ALL_ACCESS)
		if err != nil {
			return fmt.Errorf("OpenService failed in stop: %w", err)
		}
		defer windows.CloseServiceHandle(svc)
		var status windows.SERVICE_STATUS
		if err := windows.ControlService(svc, windows.SERVICE_CONTROL_STOP, &status); err != nil {
			return fmt.Errorf("ControlService(STOP) failed for %q: %w", d.Name, err)
		}
		// Poll for STOPPED (max 15s). VMX cleanup is synchronous in the
		// driver's Unload routine but may take several seconds. Wait for
		// the actual STOPPED state, not just for STOP_PENDING to clear.
		for i := 0; i < 150; i++ {
			if err := windows.QueryServiceStatus(svc, &status); err != nil {
				break
			}
			if status.CurrentState == windows.SERVICE_STOPPED {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	})
}

// Load installs and starts the driver — the Go equivalent of C++
// ManageDriver(DRIVER_FUNC_INSTALL): single Install + Start, no retry.
// Install itself handles stale-service cleanup (Stop+Remove+recreate)
// so that the freshly built .sys is always loaded.
func (d *Driver) Load() error {
	if _, err := os.Stat(d.Path); err != nil {
		return fmt.Errorf("driver file not accessible %q: %w", d.Path, err)
	}
	if err := d.Install(); err != nil {
		return err
	}
	return d.Start()
}

// Unload is a convenience helper that stops and removes the driver. It is the
// Go equivalent of the C stop+remove flow. Stop failures are tolerated so
// that Remove still runs (matching HyperDbg's tolerant unload behaviour).
func (d *Driver) Unload() error {
	if err := d.Stop(); err != nil {
		// Non-fatal: driver may not be running. Continue to remove.
	}
	return d.Remove()
}

// Exists reports whether the service is currently registered with the SCM.
func (d *Driver) Exists() (bool, error) {
	var exists bool
	err := d.withSCManager(func(sc windows.Handle) error {
		namePtr, err := windows.UTF16PtrFromString(d.Name)
		if err != nil {
			return err
		}
		svc, err := windows.OpenService(sc, namePtr, windows.SERVICE_ALL_ACCESS)
		if err != nil {
			if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
				exists = false
				return nil
			}
			return err
		}
		windows.CloseServiceHandle(svc)
		exists = true
		return nil
	})
	return exists, err
}
