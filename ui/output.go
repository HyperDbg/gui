package ui

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/ddkwork/ux/HyperDbgUnified/ui/pages"
)

// uiOutput 是 api.Debugger 的 Output 适配器。它把 debugger 的命令输出
// 路由到 LogPage，同时支持"捕获模式"：CPU 页等需要获取命令文本结果的
// 组件可临时切换到捕获模式，把输出导向 bytes.Buffer 而非日志页。
//
// 线程安全：debugger 在 goroutine 中执行命令，本类型的 Write/Printf
// 可被任意 goroutine 调用。捕获模式的切换同样受 mu 保护。
type uiOutput struct {
	logPage *pages.LogPage
	console io.Writer // 可选的镜像输出（如 os.Stdout），可为 nil

	mu      sync.Mutex
	capture *bytes.Buffer // 非 nil 时 Write 写入此 buffer
}

// newUIOutput 创建输出适配器。console 可为 nil。
func newUIOutput(logPage *pages.LogPage, console io.Writer) *uiOutput {
	return &uiOutput{
		logPage: logPage,
		console: console,
	}
}

// Write 实现 io.Writer。输出同时镜像到 console（若设置）。
func (u *uiOutput) Write(b []byte) (int, error) {
	u.mu.Lock()
	cap := u.capture
	u.mu.Unlock()

	if cap != nil {
		// 捕获模式：写入 buffer，不写日志页（避免噪音）
		if u.console != nil {
			u.console.Write(b)
		}
		return cap.Write(b)
	}

	if u.console != nil {
		u.console.Write(b)
	}
	return u.logPage.Write(b)
}

// Printf 实现 api.Output 的格式化输出。
func (u *uiOutput) Printf(format string, args ...any) error {
	msg := []byte(fmt.Sprintf(format, args...))
	_, err := u.Write(msg)
	return err
}

// StartCapture 切换到捕获模式，后续 Write 写入新 buffer。
func (u *uiOutput) StartCapture() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.capture = &bytes.Buffer{}
}

// StopCapture 结束捕获模式并返回捕获的文本。
func (u *uiOutput) StopCapture() string {
	u.mu.Lock()
	defer u.mu.Unlock()
	buf := u.capture
	u.capture = nil
	if buf == nil {
		return ""
	}
	return buf.String()
}
