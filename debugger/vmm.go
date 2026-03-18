package debugger

import (
	"fmt"
	"log/slog"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/HyperDbg/debugger/hardware"
	"github.com/ddkwork/ux/widget/treetable"
)

type VmmProvider struct {
	mu        sync.RWMutex
	packet    *Packet
	isRunning bool
	table     *treetable.TreeTable[VmmInfo]
}

type VmmInfo struct {
	Status       string
	CpuVendor    string
	VmxSupported bool
	Initializing bool
}

func New() api.Interface {
	return &VmmProvider{
		isRunning: false,
	}
}

func VmmProviderNew(packet *Packet) *VmmProvider {
	return &VmmProvider{
		packet:    packet,
		isRunning: false,
	}
}

func (v *VmmProvider) Self() any {
	return v
}

func (v *VmmProvider) Layout() layout.Widget {
	return v.table.AirTable.Layout
}

func (v *VmmProvider) Update() error {
	return nil
}

func (v *VmmProvider) Clear() {}

func (v *VmmProvider) Initialize() error {
	slog.Info("VMM: 开始初始化...")

	v.mu.Lock()
	defer v.mu.Unlock()

	if v.isRunning {
		slog.Warn("VMM: 已经在运行中")
		return nil
	}

	if v.packet == nil || !v.packet.IsConnected() {
		slog.Error("VMM: 驱动不可用")
		return fmt.Errorf("driver not available")
	}

	slog.Info("VMM: 检查 CPU 厂商...")
	if err := hardware.CheckCpuVendor(); err != nil {
		slog.Error("VMM: CPU 厂商检查失败", "err", err)
		return err
	}

	slog.Info("VMM: 检查 VMX 支持...")
	if !hardware.VmxSupportDetection() {
		slog.Error("VMM: VMX 不支持")
		return fmt.Errorf("VMX not supported")
	}

	slog.Info("VMM: 发送初始化 VMX IOCTL...")
	err := v.packet.InitializeVmx()
	if err != nil {
		slog.Error("VMM: 初始化 VMX 失败", "error", err)
		return fmt.Errorf("failed to initialize VMM: %w", err)
	}

	v.isRunning = true

	slog.Info("VMM: 初始化成功")
	return nil
}

func (v *VmmProvider) Terminate() error {
	slog.Info("VMM: 开始终止...")

	v.mu.Lock()
	defer v.mu.Unlock()

	if !v.isRunning {
		slog.Warn("VMM: 未在运行中")
		return nil
	}

	if v.packet == nil || !v.packet.IsConnected() {
		slog.Error("VMM: 驱动不可用")
		return fmt.Errorf("driver not available")
	}

	slog.Info("VMM: 发送终止 VMX IOCTL...")
	if err := v.packet.TerminateVmx(); err != nil {
		slog.Warn("VMM: 终止 VMX IOCTL 失败", "err", err)
	}

	slog.Info("VMM: 发送 IRP Pending 完成 IOCTL...")
	if err := v.packet.ReturnIrpPendingAndDisallowIoctl(); err != nil {
		slog.Warn("VMM: IRP Pending 完成 IOCTL 失败", "err", err)
	}

	v.isRunning = false

	slog.Info("VMM: 终止成功")
	return nil
}

func (v *VmmProvider) IsRunning() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.isRunning
}

func (v *VmmProvider) ReadVirtualMemory(pid uint32, address uint64, size uint32) ([]byte, error) {
	if v.packet == nil || !v.packet.IsConnected() {
		return nil, fmt.Errorf("driver not available")
	}

	return v.packet.ReadMemory(pid, address, size)
}

func (v *VmmProvider) WriteVirtualMemory(pid uint32, address uint64, data []byte) error {
	if v.packet == nil || !v.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	return v.packet.WriteMemory(pid, address, data)
}
