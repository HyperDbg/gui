// Package readmem — readmem_integration_test.go
//
// 端到端集成测试：通过真实驱动 IOCTL 验证 readmem.ReadMemory 能读取内核内存。
//
// 流程：
//  1. 加载驱动  2. 加载VMM  3. readmem.ReadMemory 读取 KUSER_SHARED_DATA
//  4. 卸载VMM  5. 卸载驱动
package readmem

import (
	"testing"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"golang.org/x/sys/windows"
)

const (
	kuserSharedDataAddr = 0x7FFE0000
	ntMajorVersionOff   = 0x26C
	kuserReadSize       = 8
)

const driverPath = `C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys`

// TestReadMemory_KuserSharedData 通过驱动 IOCTL 读取 KUSER_SHARED_DATA
// 验证 NtMajorVersion==10, NtMinorVersion==0。
func TestReadMemory_KuserSharedData(t *testing.T) {
	// 1. 加载驱动
	d := driverloader.NewDriver(driverPath)
	if err := d.Load(); err != nil {
		t.Fatalf("加载驱动失败: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload() })

	// 2. 加载VMM
	dev, err := comm.Open(comm.DeviceName)
	if err != nil {
		t.Fatalf("打开设备失败: %v", err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	vmmBuf := make([]byte, 4)
	if _, err := dev.Ioctl(hyperdbgsdk.IoctlInitVmm, vmmBuf, vmmBuf); err != nil {
		t.Skipf("IOCTL_INIT_VMM failed: %v", err)
	}
	vmmStatus := *(*uint32)(unsafe.Pointer(&vmmBuf[0]))
	if vmmStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		t.Skipf("VMM init failed (0x%08x); system lacks VT-x", vmmStatus)
	}
	t.Cleanup(func() {
		_, _ = dev.Ioctl(hyperdbgsdk.IoctlTerminateVmx, nil, nil)
	})

	// 3. ReadMemory 读取 KUSER_SHARED_DATA
	addr := uint64(kuserSharedDataAddr) + ntMajorVersionOff
	pid := windows.GetCurrentProcessId()
	data, _, err := ReadMemory(dev, addr, pid, kuserReadSize,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
	if err != nil {
		t.Fatalf("ReadMemory failed: %v", err)
	}
	if len(data) != kuserReadSize {
		t.Fatalf("ReadMemory returned %d bytes, want %d", len(data), kuserReadSize)
	}

	ntMajor := *(*uint32)(unsafe.Pointer(&data[0]))
	ntMinor := *(*uint32)(unsafe.Pointer(&data[4]))
	t.Logf("KUSER_SHARED_DATA: NtMajorVersion=%d, NtMinorVersion=%d", ntMajor, ntMinor)

	if ntMajor != 10 {
		t.Errorf("NtMajorVersion = %d, want 10", ntMajor)
	}
	if ntMinor != 0 {
		t.Errorf("NtMinorVersion = %d, want 0", ntMinor)
	}
}
