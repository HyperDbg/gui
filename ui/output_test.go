package ui

import (
	"strings"
	"sync"
	"testing"

	"github.com/ddkwork/HyperDbg/ui/pages"
)

// newUIOutputWithMirror 创建 uiOutput，用 mirror 记录所有输出。
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

func TestUIOutput_ConcurrentWrites(t *testing.T) {
	out, _ := newUIOutputWithMirror()

	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			out.Printf("line %d\n", n)
		}(i)
	}
	wg.Wait()
	// 不 panic、不 data race 即可
}

func TestSetDriverPath(t *testing.T) {
	orig := driverPath
	defer func() { driverPath = orig }()

	SetDriverPath("/custom/path/hyperkd.sys")
	if driverPath != "/custom/path/hyperkd.sys" {
		t.Errorf("driverPath = %q, want '/custom/path/hyperkd.sys'", driverPath)
	}
}
