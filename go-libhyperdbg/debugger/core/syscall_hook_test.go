// Package core — syscall_hook_test.go
//
// Go AST equivalent of d:\ux\monitor-deviceio.ds:
//
//	!syscall 0x7 script {
//	  printf("[pid:%x][%s] NtDeviceIoControlFile IoControlCode: 0x%x\n",
//	    $pid, $pname, dd(@rsp + 30));
//	}
package core

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"golang.org/x/sys/windows"
)

const (
	hyperkdSysPath = `d:\ux\HyperDbgUnified\Debug\hyperkd.sys`
	monitorLogPath = `d:\ux\HyperDbgUnified\go-libhyperdbg\debugger\core\syscall_monitor.log`
)

// Go AST callback — kernel writes this via IRP pump; user-mode resolves
// pid→name post-hoc (HookCtx has no GetProcessName).
const monitorDeviceIoCallbackSrc = `package hook
func hook(ctx *HookCtx) {
	if ctx.Reg("rax") != 0x7 { return }
	ioctl := ctx.StackReadDword(0x30)
	ctx.Printf("[pid:%x] NtDeviceIoControlFile IoControlCode: 0x%x\n", ctx.GetPid(), ioctl)
}
`

func TestSyscallHook(t *testing.T) {
	// 清空日志文件
	_ = os.Truncate(monitorLogPath, 0)
	logFile, err := os.OpenFile(monitorLogPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		t.Fatalf("open log: %v", err)
	}
	defer logFile.Close()

	// Load driver + open device + init VMM.
	d := driverloader.NewDriver(hyperkdSysPath)
	if err := d.Load(); err != nil {
		t.Fatalf("load: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload() })

	dev, err := comm.Open(comm.DeviceName)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() {
		_ = dev.Close()
	})

	var vmm initVmmRequest
	sz := uint32(unsafe.Sizeof(vmm))
	if _, err := dev.IoctlStruct(comm.IOCTL_CODE_INIT_VMM,
		unsafe.Pointer(&vmm), unsafe.Pointer(&vmm), sz, sz); err != nil {
		t.Skipf("init VMM: %v", err)
	}
	if vmm.KernelStatus != debuggerOperationWasSuccessful {
		t.Skipf("VMM init status=0x%08x", vmm.KernelStatus)
	}

	dbg := &Debugger{
		device: dev,
		state: StateVmmLoaded,
		OnPacket: func(opCode uint32, payload []byte) {
			if opCode == operationNotificationFromUserDebuggerPause {
				return
			}
			// NUL 终止的 C 字符串，去掉末尾 NUL
			if i := bytes.IndexByte(payload, 0); i >= 0 {
				payload = payload[:i]
			}
			logFile.Write(payload)
			logFile.Write([]byte{'\n'})
		},
	}

	mp, err := dbg.StartMessagePump()
	if err != nil {
		t.Fatalf("pump: %v", err)
	}
	// TERMINATE_VMX 必须在 mp.Stop() 之前：mp.Stop() 发 DISALLOW_IOCTL
	// 会阻止后续 TERMINATE_VMX → VMX 未清理 → 驱动卸载卡死。
	t.Cleanup(func() {
		_, _ = dev.IoctlStruct(comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		mp.Stop()
	})

	// Register !syscall 0x7 hook. MUST be 0x7 (specific), not 0xFFFFFFFF
	// (all) — see BSOD note below.
	tag, err := dbg.SyscallHook(0x7, monitorDeviceIoCallbackSrc)
	if err != nil {
		t.Fatalf("hook: %v", err)
	}
	t.Logf("hook registered tag=%d, vmm status=0x%08x", tag, vmm.KernelStatus)
	t.Cleanup(func() { clearEvent(t, dev, tag) })

	t.Logf("monitoring !syscall 0x7 for 10s...")
	time.Sleep(10 * time.Second)
	mp.Stop()

	// Parse log: extract pid + ioctl, resolve pid→name, drop self-hits
	// (MessagePump's IOCTL loop triggers 0x7 on every iteration).
	data, _ := os.ReadFile(monitorLogPath)
	t.Logf("log file size=%d bytes", len(data))
	pidRe := regexp.MustCompile(`\((\d{2}:\d{2}:\d{2}\.\d+).*\[pid:([0-9a-f]+)\] NtDeviceIoControlFile IoControlCode: (0x[0-9a-f]+)`)
	selfPid := uint32(os.Getpid())
	type hit struct {
		ts, name string
		pid      uint32
		ioctl    string
	}
	var hits []hit
	pidToName := map[uint32]string{}
	selfCount := 0
	// Match against raw bytes (FindAllSubmatch) — line-splitting is
	// unreliable because the kernel emits literal "\n" inside each message
	// and the pump appends a real '\n', producing mixed separators.
	for _, m := range pidRe.FindAllSubmatch(data, -1) {
		pid64, _ := strconv.ParseUint(string(m[2]), 16, 32)
		pid := uint32(pid64)
		if pid == selfPid {
			selfCount++
			continue
		}
		if _, ok := pidToName[pid]; !ok {
			pidToName[pid] = processName(pid)
		}
		hits = append(hits, hit{string(m[1]), pidToName[pid], pid, string(m[3])})
	}

	// Rewrite log: real hits only, with " | name: <exe>" appended.
	var b strings.Builder
	for _, h := range hits {
		b.WriteString(h.ts)
		b.WriteString("  pid:")
		b.WriteString(strconv.FormatUint(uint64(h.pid), 16))
		b.WriteString("  ")
		b.WriteString(h.name)
		b.WriteString("  ioctl:")
		b.WriteString(h.ioctl)
		b.WriteString("\n")
	}
	_ = os.WriteFile(monitorLogPath, []byte(b.String()), 0o644)

	// Summary.
	type stat struct {
		pid  uint32
		name string
		n    int
	}
	stats := map[uint32]*stat{}
	for _, h := range hits {
		if s, ok := stats[h.pid]; ok {
			s.n++
		} else {
			stats[h.pid] = &stat{h.pid, h.name, 1}
		}
	}
	flat := make([]*stat, 0, len(stats))
	for _, s := range stats {
		flat = append(flat, s)
	}
	sort.Slice(flat, func(i, j int) bool { return flat[i].n > flat[j].n })

	t.Logf("=== %d real hits / %d procs / 10s (self %d filtered) ===",
		len(hits), len(flat), selfCount)
	for _, s := range flat {
		t.Logf("  pid:%-5x  %-30s  %d hits", s.pid, s.name, s.n)
	}
}

// processName resolves pid→basename via OpenProcess+QueryFullProcessImageName.
func processName(pid uint32) string {
	const PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	h, err := windows.OpenProcess(PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return "<exited>"
	}
	defer windows.CloseHandle(h)
	var buf [windows.MAX_PATH]uint16
	n := uint32(len(buf))
	if err := windows.QueryFullProcessImageName(h, 0, &buf[0], &n); err != nil {
		return "<pid:" + strconv.FormatUint(uint64(pid), 16) + ">"
	}
	return filepath.Base(windows.UTF16ToString(buf[:n]))
}
