package pages

import (
	"fmt"
	"strings"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/widget/logview"
)

// LogPage 是日志页。它既作为 UI 组件展示日志，也作为 api.Debugger 的
// Output 接收端（通过 uiOutput 适配器）。
//
// 由于 debugger 的命令在 goroutine 中执行（工具栏按钮回调），而 gio 的
// widget 只能在 UI 线程访问，LogPage 维护一个线程安全的 pending 缓冲：
//   - Write/Printf（任意 goroutine）→ 追加到 pending
//   - Layout（UI 线程）→ Flush 把 pending 排空到 logView
type LogPage struct {
	logView *logview.LogView
	mu      sync.Mutex
	pending []string // 待 flush 的日志行
}

func NewLog() *LogPage {
	return &LogPage{
		logView: logview.New(),
	}
}

// LogView 返回内部 logView（供需要直接调用的场景使用）。
func (p *LogPage) LogView() *logview.LogView {
	return p.logView
}

// Write 实现 io.Writer。字节流按行拆分后追加到 pending 缓冲。
// 该方法可被任意 goroutine 调用（debugger 命令输出在 goroutine 中产生）。
func (p *LogPage) Write(b []byte) (int, error) {
	text := string(b)
	p.mu.Lock()
	if strings.HasSuffix(text, "\n") {
		text = text[:len(text)-1]
	}
	if text != "" {
		for _, line := range strings.Split(text, "\n") {
			p.pending = append(p.pending, line)
		}
	}
	p.mu.Unlock()
	app.RequestRedraw()
	return len(b), nil
}

// Printf 实现 api.Output 的格式化输出（被 uiOutput.Printf 委托调用）。
func (p *LogPage) Printf(format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	_, err := p.Write([]byte(msg))
	return err
}

// flush 把 pending 缓冲排空到 logView。只能在 UI 线程调用。
func (p *LogPage) flush() {
	p.mu.Lock()
	lines := p.pending
	p.pending = nil
	p.mu.Unlock()
	for _, line := range lines {
		p.logView.Info(line)
	}
}

func (p *LogPage) Println(a ...any) {
	msg := fmt.Sprintln(a...)
	p.Write([]byte(msg))
}

func (p *LogPage) Print(a ...any) {
	msg := fmt.Sprint(a...)
	p.Write([]byte(msg))
}

// Clear 清空日志视图与 pending 缓冲。
func (p *LogPage) Clear() {
	p.mu.Lock()
	p.pending = nil
	p.mu.Unlock()
	p.logView.Clear()
}

func (p *LogPage) Layout() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		p.flush()
		return p.logView.Layout(gtx)
	}
}
