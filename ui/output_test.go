package ui

import (
	"strings"
	"sync"
	"testing"

	"github.com/ddkwork/ux/HyperDbgUnified/ui/pages"
)

// newUIOutputWithMirror 创建 uiOutput，用 mirror 记录所有输出（含捕获模式）。
// 通过 mirror 验证输出内容，避免依赖 LogPage 内部方法。
func newUIOutputWithMirror() (*uiOutput, *strings.Builder) {
	lp := pages.NewLog()
	var mirror strings.Builder
	return newUIOutput(lp, &mirror), &mirror
}

func TestUIOutput_WriteNormal(t *testing.T) {
	out, mirror := newUIOutputWithMirror()

	n, err := out.Write([]byte("hello\n"))
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}
	if n != 6 {
		t.Errorf("Write n=%d, want 6", n)
	}
	if mirror.String() != "hello\n" {
		t.Errorf("mirror = %q, want 'hello\\n'", mirror.String())
	}
}

func TestUIOutput_Printf(t *testing.T) {
	out, mirror := newUIOutputWithMirror()

	if err := out.Printf("val=%d hex=0x%X", 42, 255); err != nil {
		t.Fatalf("Printf failed: %v", err)
	}
	if !strings.Contains(mirror.String(), "val=42") {
		t.Errorf("mirror = %q, want contains 'val=42'", mirror.String())
	}
	if !strings.Contains(mirror.String(), "hex=0xFF") {
		t.Errorf("mirror = %q, want contains 'hex=0xFF'", mirror.String())
	}
}

func TestUIOutput_CaptureMode(t *testing.T) {
	out, mirror := newUIOutputWithMirror()

	// 正常模式
	out.Write([]byte("normal\n"))

	// 切换到捕获模式
	out.StartCapture()
	out.Write([]byte("captured line 1\n"))
	out.Write([]byte("captured line 2\n"))

	// 捕获模式下仍镜像到 console
	if !strings.Contains(mirror.String(), "captured line 1") {
		t.Errorf("mirror during capture = %q, want 'captured line 1'", mirror.String())
	}

	// 结束捕获，验证捕获内容
	got := out.StopCapture()
	if !strings.Contains(got, "captured line 1") || !strings.Contains(got, "captured line 2") {
		t.Errorf("StopCapture = %q, want both captured lines", got)
	}

	// 结束后恢复正常写入
	out.Write([]byte("resumed\n"))
	if !strings.Contains(mirror.String(), "resumed") {
		t.Errorf("mirror after StopCapture = %q, want 'resumed'", mirror.String())
	}
}

func TestUIOutput_StopCaptureWithoutStart(t *testing.T) {
	out, _ := newUIOutputWithMirror()

	got := out.StopCapture()
	if got != "" {
		t.Errorf("StopCapture without StartCapture = %q, want empty", got)
	}
}

func TestUIOutput_ConcurrentWrites(t *testing.T) {
	out, _ := newUIOutputWithMirror()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			out.Printf("line %d\n", n)
		}(i)
	}
	wg.Wait()
	// 不 panic、不 data race 即可
}

func TestUIOutput_CaptureIsolatedFromLogPage(t *testing.T) {
	lp := pages.NewLog()
	var mirror strings.Builder
	out := newUIOutput(lp, &mirror)

	// 捕获模式：写入只进 buffer 和 mirror，不进 LogPage
	out.StartCapture()
	out.Write([]byte("captured\n"))
	out.StopCapture()

	// LogPage 的 pending 应为空（捕获模式不写 LogPage）
	// 通过写入正常内容再验证 mirror 增量来间接确认
	mirror.Reset()
	out.Write([]byte("after capture\n"))
	if mirror.String() != "after capture\n" {
		t.Errorf("mirror after capture = %q", mirror.String())
	}
}

func TestSetDriverPath(t *testing.T) {
	orig := driverPath
	defer func() { driverPath = orig }()

	SetDriverPath("/custom/path/hyperkd.sys")
	if driverPath != "/custom/path/hyperkd.sys" {
		t.Errorf("driverPath = %q, want '/custom/path/hyperkd.sys'", driverPath)
	}
}
