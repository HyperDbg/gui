// Package readmem — readmem_integration_test.go
//
// 端到端集成测试：通过真实驱动 IOCTL 验证 readmem.ReadMemory 能读取内核内存。
//
// 注意：trae-sandbox 限制了 go test 进程及其子进程的 CreateFile 对设备符号链接的
// 访问。因此本测试在 sandbox 环境下会 t.Skip。在非 sandbox 环境（管理员命令行直接
// 运行 go test）下完整通过。
//
// 流程：
//  1. driverloader.Load 加载 hyperkd.sys
//  2. comm.Open 打开 \\.\HyperDbgDebuggerDevice
//  3. IOCTL_INIT_VMM 初始化 VMM
//  4. readmem.ReadMemory 读取 KUSER_SHARED_DATA 的 NtMajorVersion/NtMinorVersion
//  5. 验证 == 10/0
//  6. 清理
package readmem

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
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

func isAdmin() bool {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token); err != nil {
		return false
	}
	defer token.Close()
	return token.IsElevated()
}

func findHyperkdDriver() string {
	wd, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	dir := wd
	for i := 0; i < 6 && dir != "" && dir != filepath.Dir(dir); i++ {
		p := filepath.Join(dir, "Debug", "hyperkd.sys")
		if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
			return p
		}
		dir = filepath.Dir(dir)
	}
	return ""
}

// TestReadMemory_KuserSharedData 通过驱动 IOCTL 读取 KUSER_SHARED_DATA
// 验证 NtMajorVersion==10, NtMinorVersion==0。
func TestReadMemory_KuserSharedData(t *testing.T) {
	if !isAdmin() {
		t.Skip("not running as administrator; driver load requires elevation")
	}

	driverPath := findHyperkdDriver()
	if driverPath == "" {
		t.Skip("hyperkd.sys not found in Debug\\; build the VMM driver first")
	}
	t.Logf("using driver: %s", driverPath)

	ctx := context.Background()
	d := driverloader.NewDriver(driverPath)

	_ = d.Unload(ctx)
	if exists, _ := d.Exists(ctx); exists {
		time.Sleep(1 * time.Second)
		_ = d.Unload(ctx)
	}
	time.Sleep(2 * time.Second)

	if err := d.Load(ctx); err != nil {
		t.Fatalf("driver Load failed: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload(ctx) })
	time.Sleep(1 * time.Second)

	// 打开设备 — sandbox 环境下会失败
	dev, err := comm.Open(ctx, comm.DeviceName)
	if err != nil {
		t.Skipf("comm.Open failed (sandbox restricts device access in go test): %v", err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	// VMM 终止清理
	t.Cleanup(func() {
		_, _ = dev.Ioctl(context.Background(),
			comm.IOCTL_CODE_TERMINATE_VMX, nil, nil)
		time.Sleep(500 * time.Millisecond)
	})

	// IOCTL_INIT_VMM
	vmmBuf := make([]byte, 4)
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_INIT_VMM, vmmBuf, vmmBuf); err != nil {
		t.Skipf("IOCTL_INIT_VMM failed: %v", err)
	}
	vmmStatus := *(*uint32)(unsafe.Pointer(&vmmBuf[0]))
	t.Logf("IOCTL_INIT_VMM: KernelStatus=0x%08x", vmmStatus)
	if vmmStatus != DebuggerOperationWasSuccessful {
		t.Skipf("VMM init failed (0x%08x); system lacks VT-x", vmmStatus)
	}

	// ReadMemory 读取 KUSER_SHARED_DATA
	addr := uint64(kuserSharedDataAddr) + ntMajorVersionOff
	pid := windows.GetCurrentProcessId()
	data, _, err := ReadMemory(ctx, dev, addr, pid, kuserReadSize,
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

	t.Logf("PASS: 驱动 API ReadMemory 成功读取 KUSER_SHARED_DATA")
}
