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
	"context"
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
func (d *Driver) withSCManager(ctx context.Context, fn func(windows.Handle) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	sc, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("OpenSCManager failed: %w", err)
	}
	defer windows.CloseServiceHandle(sc)
	return fn(sc)
}

// Install registers the driver as a kernel service (SERVICE_DEMAND_START).
//
// If the service already exists it is treated as success so that Install is
// idempotent — matching HyperDbg's behaviour where a stale service entry is
// reused rather than causing a hard failure. Use Remove before Install to
// guarantee a fresh registration.
func (d *Driver) Install(ctx context.Context) error {
	return d.withSCManager(ctx, func(sc windows.Handle) error {
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
		if err != nil {
			if err == windows.ERROR_SERVICE_EXISTS {
				return nil
			}
			if err == windows.ERROR_SERVICE_MARKED_FOR_DELETE {
				// The previous instance is still being torn down asynchronously.
				// Wait and retry once.
				time.Sleep(3 * time.Second)
				svc, err = windows.CreateService(
					sc, namePtr, namePtr, windows.SERVICE_ALL_ACCESS,
					windows.SERVICE_KERNEL_DRIVER, windows.SERVICE_DEMAND_START,
					windows.SERVICE_ERROR_NORMAL, exePtr, nil, nil, nil, nil, nil)
				if err != nil {
					if err == windows.ERROR_SERVICE_EXISTS || err == windows.ERROR_SERVICE_MARKED_FOR_DELETE {
						return nil
					}
					return fmt.Errorf("CreateService (retry) failed for %q: %w", d.Name, err)
				}
				windows.CloseServiceHandle(svc)
				return nil
			}
			return fmt.Errorf("CreateService failed for %q: %w", d.Name, err)
		}
		windows.CloseServiceHandle(svc)
		return nil
	})
}

// Remove deletes the service registration. A non-existent service is treated
// as success so Remove is idempotent and safe to call after a failed install.
func (d *Driver) Remove(ctx context.Context) error {
	return d.withSCManager(ctx, func(sc windows.Handle) error {
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

// Start starts the driver service. If the service is already running this is
// treated as success (ERROR_SERVICE_ALREADY_RUNNING).
//
// Common failure causes returned:
//   - ERROR_PATH_NOT_FOUND        — driver file missing or AV blocking access
//   - ERROR_INVALID_IMAGE_HASH    — driver signature enforcement / HVCI active
func (d *Driver) Start(ctx context.Context) error {
	return d.withSCManager(ctx, func(sc windows.Handle) error {
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

// Stop sends SERVICE_CONTROL_STOP to the driver service. A service that is not
// running or not removable returns an error; callers may treat such errors as
// non-fatal (the C version also just warns and proceeds to remove).
func (d *Driver) Stop(ctx context.Context) error {
	return d.withSCManager(ctx, func(sc windows.Handle) error {
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
		return nil
	})
}

// Load is a convenience helper that installs and starts the driver. It is the
// Go equivalent of the C ManageDriver(DRIVER_FUNC_INSTALL) flow.
//
// NOTE: Do NOT call Stop here to "clean up" a stale running instance.
// Stopping a driver that still has VMX active causes a BSOD
// (DRIVER_IRQL_NOT_LESS_OR_EQUAL in AsmVmexitHandler) because the OS
// unmaps the driver's code section while VMX is still enabled — the next
// VMX exit jumps to unmapped code. The stale-driver state
// (g_HyperLogInitialized==TRUE after DrvClose) must instead be fixed in
// the driver source (Loader.c: LoaderInitHyperLog must return TRUE when
// already initialized).
func (d *Driver) Load(ctx context.Context) error {
	if _, err := os.Stat(d.Path); err != nil {
		return fmt.Errorf("driver file not accessible %q: %w", d.Path, err)
	}
	// Retry Install+Start a few times: immediately after a previous
	// TERMINATE_VMX the VMX state may not have fully settled, causing
	// DriverEntry (and thus StartService) to fail. A short delay + retry
	// avoids requiring a manual `sc delete` between runs.
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			// Clean up any stale service from the previous failed attempt.
			_ = d.Remove(ctx)
			time.Sleep(2 * time.Second)
		}
		if err := d.Install(ctx); err != nil {
			lastErr = err
			continue
		}
		if err := d.Start(ctx); err != nil {
			_ = d.Remove(ctx)
			lastErr = err
			continue
		}
		return nil
	}
	return lastErr
}

// Unload is a convenience helper that stops and removes the driver. It is the
// Go equivalent of the C stop+remove flow. Stop failures are tolerated so
// that Remove still runs (matching HyperDbg's tolerant unload behaviour).
func (d *Driver) Unload(ctx context.Context) error {
	if err := d.Stop(ctx); err != nil {
		// Non-fatal: driver may not be running. Continue to remove.
	}
	return d.Remove(ctx)
}

// Exists reports whether the service is currently registered with the SCM.
func (d *Driver) Exists(ctx context.Context) (bool, error) {
	var exists bool
	err := d.withSCManager(ctx, func(sc windows.Handle) error {
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
