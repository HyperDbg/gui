// readmem_helper 是独立的驱动 API 调用工具。
// 编译: go build -o Debug\readmem_helper.exe ./debugger/readmem/cmd/readmem_helper/
// 用法: readmem_helper.exe
//
// 流程: 加载 hyperkd.sys → 打开设备 → IOCTL_INIT_VMM → ReadMemory 读 KUSER_SHARED_DATA → 输出 JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
	"golang.org/x/sys/windows"
)

type result struct {
	InitStatus   string `json:"init_status"`
	KernelStatus string `json:"kernel_status"`
	NtMajor      uint32 `json:"nt_major"`
	NtMinor      uint32 `json:"nt_minor"`
	Error        string `json:"error"`
}

func main() {
	driverPath := filepath.Join(filepath.Dir(os.Args[0]), "hyperkd.sys")
	if _, err := os.Stat(driverPath); err != nil {
		// 尝试上级目录
		driverPath = filepath.Join(filepath.Dir(filepath.Dir(os.Args[0])), "Debug", "hyperkd.sys")
	}

	d := driverloader.NewDriver(driverPath)
	_ = d.Unload()
	time.Sleep(2 * time.Second)

	if err := d.Load(); err != nil {
		out, _ := json.Marshal(result{Error: "Load failed: " + err.Error()})
		fmt.Println(string(out))
		return
	}
	defer d.Unload()
	time.Sleep(1 * time.Second)

	dev, err := comm.Open(comm.DeviceName)
	if err != nil {
		out, _ := json.Marshal(result{Error: "Open failed: " + err.Error()})
		fmt.Println(string(out))
		return
	}
	defer dev.Close()

	// IOCTL_INIT_VMM
	vmmBuf := make([]byte, 4)
	_, _ = dev.Ioctl(comm.IOCTL_CODE_INIT_VMM, vmmBuf, vmmBuf)
	vmmStatus := *(*uint32)(unsafe.Pointer(&vmmBuf[0]))
	if vmmStatus != 0xFFFFFFFF {
		out, _ := json.Marshal(result{Error: "VMM init failed", InitStatus: fmt.Sprintf("0x%08X", vmmStatus)})
		fmt.Println(string(out))
		return
	}

	// ReadMemory KUSER_SHARED_DATA
	addr := uint64(0x7FFE026C)
	pid := windows.GetCurrentProcessId()
	data, _, err := readmem.ReadMemory(dev, addr, pid, 8,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
	if err != nil {
		out, _ := json.Marshal(result{Error: "ReadMemory failed: " + err.Error(), InitStatus: "0xFFFFFFFF"})
		fmt.Println(string(out))
		return
	}

	ntMajor := *(*uint32)(unsafe.Pointer(&data[0]))
	ntMinor := *(*uint32)(unsafe.Pointer(&data[4]))

	// TERMINATE_VMX
	_, _ = dev.Ioctl(comm.IOCTL_CODE_TERMINATE_VMX, nil, nil)

	out, _ := json.Marshal(result{
		InitStatus:   "0xFFFFFFFF",
		KernelStatus: "0xFFFFFFFF",
		NtMajor:      ntMajor,
		NtMinor:      ntMinor,
	})
	fmt.Println(string(out))
}
