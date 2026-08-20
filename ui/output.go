package ui

import (
	"fmt"
	"io"

	"github.com/ddkwork/HyperDbg/ui/pages"
)

// uiOutput 是 api.Debugger 的 Output 适配器。它把 debugger 的命令输出
// 路由到 LogPage，同时可选地镜像到 console（如 os.Stdout）。
//
// 线程安全：debugger 在 goroutine 中执行命令，本类型的 Write/Printf
// 可被任意 goroutine 调用。LogPage.Write 内部自带锁。
type uiOutput struct {
	logPage *pages.LogPage
	console io.Writer // 可选的镜像输出（如 os.Stdout），可为 nil
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
