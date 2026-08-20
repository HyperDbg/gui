package pages

import (
	"fmt"
	"strconv"
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/codeeditor"
	"github.com/ddkwork/ux/widget/input"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// BreaksPage 是断点页：设置/清除/启用/禁用软件断点，列出当前断点。
type BreaksPage struct {
	dbg *api.Debugger

	addrInput *input.Input
	tagInput  *input.Input
	list      *codeeditor.CodeEditor

	setBtn     *button.Button
	clearBtn   *button.Button
	disableBtn *button.Button
	enableBtn  *button.Button
	listBtn    *button.Button
}

func NewBreaks(dbg *api.Debugger) *BreaksPage {
	p := &BreaksPage{
		dbg:        dbg,
		addrInput:  input.CompactInput(),
		tagInput:   input.CompactInput(),
		list:       codeeditor.New("（点击\"列出断点\"查看）", "go"),
		setBtn:     button.Filled(),
		clearBtn:   button.Filled(),
		disableBtn: button.Filled(),
		enableBtn:  button.Filled(),
		listBtn:    button.Filled(),
	}
	p.addrInput.Editor.SetText("0x00010000")
	p.tagInput.Editor.SetText("0x0")
	p.list.SetReadOnly(true)
	return p
}

func (p *BreaksPage) Layout() layout.Widget {
	return p.layout
}

func (p *BreaksPage) layout(gtx layout.Context) layout.Dimensions {
	if p.setBtn.Clicked(gtx) {
		go p.bpSet()
	}
	if p.clearBtn.Clicked(gtx) {
		go p.bpClear()
	}
	if p.disableBtn.Clicked(gtx) {
		go p.bpDisable()
	}
	if p.enableBtn.Clicked(gtx) {
		go p.bpEnable()
	}
	if p.listBtn.Clicked(gtx) {
		go p.bpList()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: 4}.Layout(gtx,
						layout.Rigid(p.label("地址:")),
						layout.Rigid(p.input(p.addrInput, 160)),
						layout.Rigid(p.btn(p.setBtn, "设置断点 bp")),
						layout.Rigid(p.label("Tag:")),
						layout.Rigid(p.input(p.tagInput, 140)),
						layout.Rigid(p.btn(p.clearBtn, "清除 bc")),
						layout.Rigid(p.btn(p.disableBtn, "禁用 bd")),
						layout.Rigid(p.btn(p.enableBtn, "启用 be")),
						layout.Rigid(p.btn(p.listBtn, "列出 bl")),
					)
				},
			)
		}),
		layout.Flexed(1, p.list.Layout),
	)
}

func (p *BreaksPage) label(text string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Right: unit.Dp(4), Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return button.Text().Layout(gtx, text)
		})
	}
}

func (p *BreaksPage) input(inp *input.Input, widthDp int) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Dp(unit.Dp(widthDp))
		return inp.Layout(gtx)
	}
}

func (p *BreaksPage) btn(b *button.Button, label string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return b.Layout(gtx, label)
	}
}

func (p *BreaksPage) parseAddr() uint64 {
	s := strings.TrimSpace(p.addrInput.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}
	return v
}

func (p *BreaksPage) parseTag() uint64 {
	s := strings.TrimSpace(p.tagInput.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		v, _ = strconv.ParseUint(s, 10, 64)
	}
	return v
}

func (p *BreaksPage) bpSet() {
	tag, err := p.dbg.BpSet(p.parseAddr())
	if err != nil {
		fmt.Printf("BpSet 失败: %v\n", err)
		return
	}
	fmt.Printf("断点已设置 tag=0x%X\n", tag)
	p.bpList()
}

func (p *BreaksPage) bpClear() {
	if err := p.dbg.BpClear(p.parseTag()); err != nil {
		fmt.Printf("BpClear 失败: %v\n", err)
	}
	p.bpList()
}

func (p *BreaksPage) bpDisable() {
	if err := p.dbg.BpDisable(p.parseTag()); err != nil {
		fmt.Printf("BpDisable 失败: %v\n", err)
	}
	p.bpList()
}

func (p *BreaksPage) bpEnable() {
	if err := p.dbg.BpEnable(p.parseTag()); err != nil {
		fmt.Printf("BpEnable 失败: %v\n", err)
	}
	p.bpList()
}

func (p *BreaksPage) bpList() {
	// BpList 走字符串命令路径，输出到 debugger output（日志页）。
	// 这里额外用 Exec 捕获显示在断点页。
	if err := p.dbg.Exec("bl"); err != nil {
		p.list.SetCode(fmt.Sprintf("bl 失败: %v", err))
	}
	app.RequestRedraw()
}

var _ = api.Breakpoint{}
