// Package api — commands_hwdbg.go
//
// 对应 debugger/commands/hwdbg 包的 2 个 hwdbg 命令的 typed API。
//
// 命令对照表（与 hwdbg.go RegisterAll 顺序一致）：
//
//	hw      → Hw      (本文件, stub)
//	hw_clk  → HwClk   (本文件, stub)
//
// hwdbg 是 HyperDbg 的硬件调试子模块（基于 FPGA/专用硬件的调试设备），
// 与 VMM 软件调试平行。当前 Go 实装仅注册了 stubs。
package api

import (
	"fmt"
)

// Hw 对应 'hw'：启动/查询硬件调试设备。
func (d *Debugger) Hw() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return fmt.Errorf("Hw: hwdbg hardware device not present (FPGA-based; see hwdbg package)")
}

// HwClk 对应 'hw_clk'：硬件调试设备的时钟配置。
func (d *Debugger) HwClk() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return fmt.Errorf("HwClk: hwdbg hardware device not present")
}
