package breakpoint

import (
	"fmt"
	"log/slog"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

type Control struct {
	driver *driver.DriverHandle
}

func NewControl(driverHandle *driver.DriverHandle) *Control {
	return &Control{
		driver: driverHandle,
	}
}

func (c *Control) EnableBreakpoint(tag uint32) error {
	slog.Info("Breakpoint: Enabling breakpoint", "tag", tag)

	if c.driver == nil || !c.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	return fmt.Errorf("ModifyEvent not implemented")
}

func (c *Control) DisableBreakpoint(tag uint32) error {
	slog.Info("Breakpoint: Disabling breakpoint", "tag", tag)

	if c.driver == nil || !c.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	return fmt.Errorf("ModifyEvent not implemented")
}

func (c *Control) RemoveBreakpoint(tag uint32) error {
	slog.Info("Breakpoint: Removing breakpoint", "tag", tag)

	if c.driver == nil || !c.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	return fmt.Errorf("UnregisterEvent not implemented")
}

func (c *Control) ClearAllBreakpoints() error {
	slog.Info("Breakpoint: Clearing all breakpoints")

	if c.driver == nil || !c.driver.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	return fmt.Errorf("ClearEvents not implemented")
}
