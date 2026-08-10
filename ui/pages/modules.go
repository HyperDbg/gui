package pages

import (
	"context"
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/codeeditor"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// ModulesPage 是模块/线程页：执行 lm（列出模块）或 thread（列出线程）命令，
// 输出通过 debugger 的 output 显示在日志页。本页提供快捷按钮。
type ModulesPage struct {
	dbg *api.Debugger

	list       *codeeditor.CodeEditor
	lmBtn      *button.Button
	threadBtn  *button.Button
	processBtn *button.Button
	cpuBtn     *button.Button
	helpBtn    *button.Button
}

func NewModules(dbg *api.Debugger) *ModulesPage {
	p := &ModulesPage{
		dbg:        dbg,
		list:       codeeditor.New("（点击按钮执行命令，输出见日志页）", "go"),
		lmBtn:      button.Filled(),
		threadBtn:  button.Filled(),
		processBtn: button.Filled(),
		cpuBtn:     button.Filled(),
		helpBtn:    button.Filled(),
	}
	p.list.SetReadOnly(true)
	return p
}

func (p *ModulesPage) Layout() layout.Widget {
	return p.layout
}

func (p *ModulesPage) layout(gtx layout.Context) layout.Dimensions {
	if p.lmBtn.Clicked(gtx) {
		go p.exec("lm")
	}
	if p.threadBtn.Clicked(gtx) {
		go p.exec("thread")
	}
	if p.processBtn.Clicked(gtx) {
		go p.exec("process")
	}
	if p.cpuBtn.Clicked(gtx) {
		go p.exec("cpu")
	}
	if p.helpBtn.Clicked(gtx) {
		go p.exec("help")
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: 4}.Layout(gtx,
						layout.Rigid(p.btn(p.lmBtn, "lm 列出模块")),
						layout.Rigid(p.btn(p.threadBtn, "thread 列出线程")),
						layout.Rigid(p.btn(p.processBtn, "process 列出进程")),
						layout.Rigid(p.btn(p.cpuBtn, "cpu CPU信息")),
						layout.Rigid(p.btn(p.helpBtn, "help 帮助")),
					)
				},
			)
		}),
		layout.Flexed(1, p.list.Layout),
	)
}

func (p *ModulesPage) btn(b *button.Button, label string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return b.Layout(gtx, label)
	}
}

func (p *ModulesPage) exec(cmd string) {
	ctx := context.Background()
	if err := p.dbg.Exec(ctx, cmd); err != nil {
		p.list.SetCode(fmt.Sprintf("%s 失败: %v", cmd, err))
	}
	app.RequestRedraw()
}

var _ = api.Module{}
