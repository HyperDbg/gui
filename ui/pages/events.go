package pages

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/codeeditor"
	"github.com/ddkwork/ux/widget/input"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// 默认 hook 回调源码（触发断点）
const defaultCallback = "func hook(ctx *HookCtx) { ctx.Break() }"

// HookRecord 记录一个已注册的 hook（UI 层跟踪，因 core 不暴露 tag 表）。
type HookRecord struct {
	Tag      uint64
	Type     string
	Address  uint64
	PID      uint32
	Callback string
	Enabled  bool
}

// EventsPage 是事件页：注册 hook + 列出已注册 hook + 管理事件。
//
// 由于 go-libhyperdbg 不暴露 ListEvents，UI 层自行维护已注册 hook 表。
// 仅跟踪通过本页面 typed API 注册的 hook；命令栏输入的 ! 命令不会出现在表中。
type EventsPage struct {
	dbg *api.Debugger

	addrInput *input.Input
	pidInput  *input.Input
	tagInput  *input.Input

	list *codeeditor.CodeEditor

	mu    sync.Mutex
	hooks []HookRecord

	clearBtn   *button.Button
	disableBtn *button.Button
	enableBtn  *button.Button
}

func NewEvents(dbg *api.Debugger) *EventsPage {
	p := &EventsPage{
		dbg:        dbg,
		addrInput:  input.CompactInput(),
		pidInput:   input.CompactInput(),
		tagInput:   input.CompactInput(),
		list:       codeeditor.New("（无已注册 hook）", "go"),
		clearBtn:   button.Filled(),
		disableBtn: button.Filled(),
		enableBtn:  button.Filled(),
	}
	p.addrInput.Editor.SetText("0x00010000")
	p.pidInput.Editor.SetText("0")
	p.tagInput.Editor.SetText("0x0")
	p.list.SetReadOnly(true)
	return p
}

func (p *EventsPage) Layout() layout.Widget {
	return p.layout
}

func (p *EventsPage) layout(gtx layout.Context) layout.Dimensions {
	// 事件管理按钮
	if p.clearBtn.Clicked(gtx) {
		go p.clearEvent()
	}
	if p.disableBtn.Clicked(gtx) {
		go p.disableEvent()
	}
	if p.enableBtn.Clicked(gtx) {
		go p.enableEvent()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// 参数输入区
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: 4}.Layout(gtx,
						layout.Rigid(p.fieldLabel("地址:")),
						layout.Rigid(p.fieldInput(p.addrInput, 120)),
						layout.Rigid(p.fieldLabel("PID:")),
						layout.Rigid(p.fieldInput(p.pidInput, 80)),
						layout.Rigid(p.fieldLabel("Tag:")),
						layout.Rigid(p.fieldInput(p.tagInput, 100)),
					)
				},
			)
		}),
		// hook 注册按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: 4}.Layout(gtx,
						layout.Rigid(p.hookBtn("EptHook", p.registerEptHook)),
						layout.Rigid(p.hookBtn("EptHookForProcess", p.registerEptHookForProcess)),
						layout.Rigid(p.hookBtn("SyscallHook", p.registerSyscallHook)),
						layout.Rigid(p.hookBtn("SysretHook(ForProc)", p.registerSysretHook)),
						layout.Rigid(p.hookBtn("MonitorRead", p.registerMonitorRead)),
						layout.Rigid(p.hookBtn("MonitorWrite", p.registerMonitorWrite)),
						layout.Rigid(p.hookBtn("MonitorExec", p.registerMonitorExec)),
						layout.Rigid(p.hookBtn("CpuidHook", p.registerCpuidHook)),
						layout.Rigid(p.hookBtn("DrHook", p.registerDrHook)),
						layout.Rigid(p.hookBtn("VmcallHook", p.registerVmcallHook)),
						layout.Rigid(p.hookBtn("ModeHook", p.registerModeHook)),
						layout.Rigid(p.hookBtn("MsrRead", p.registerMsrRead)),
						layout.Rigid(p.hookBtn("MsrWrite", p.registerMsrWrite)),
						layout.Rigid(p.hookBtn("Exception", p.registerException)),
						layout.Rigid(p.hookBtn("Interrupt", p.registerInterrupt)),
						layout.Rigid(p.hookBtn("IoIn", p.registerIoIn)),
						layout.Rigid(p.hookBtn("IoOut", p.registerIoOut)),
						layout.Rigid(p.hookBtn("Crwrite", p.registerCrwrite)),
						layout.Rigid(p.hookBtn("Xsetbv", p.registerXsetbv)),
					)
				},
			)
		}),
		// 事件管理按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: 4}.Layout(gtx,
						layout.Rigid(p.hookBtn("ClearEvent", p.clearEvent)),
						layout.Rigid(p.hookBtn("DisableEvent", p.disableEvent)),
						layout.Rigid(p.hookBtn("EnableEvent", p.enableEvent)),
					)
				},
			)
		}),
		// hook 列表
		layout.Flexed(1, p.list.Layout),
	)
}

func (p *EventsPage) fieldLabel(text string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Right: unit.Dp(4), Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return button.Text().Layout(gtx, text)
		})
	}
}

func (p *EventsPage) fieldInput(inp *input.Input, widthDp int) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Dp(unit.Dp(widthDp))
		return inp.Layout(gtx)
	}
}

func (p *EventsPage) hookBtn(label string, fn func()) layout.Widget {
	btn := button.Outlined()
	return func(gtx layout.Context) layout.Dimensions {
		if btn.Clicked(gtx) {
			go fn()
		}
		return btn.Layout(gtx, label)
	}
}

// addHook 记录已注册的 hook 并刷新列表显示。
func (p *EventsPage) addHook(rec HookRecord) {
	p.mu.Lock()
	p.hooks = append(p.hooks, rec)
	p.mu.Unlock()
	p.refreshList()
}

func (p *EventsPage) refreshList() {
	p.mu.Lock()
	hooks := make([]HookRecord, len(p.hooks))
	copy(hooks, p.hooks)
	p.mu.Unlock()

	if len(hooks) == 0 {
		p.list.SetCode("（无已注册 hook）")
		return
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-18s %-22s %-18s %-10s %-8s\n", "Tag", "Type", "Address", "PID", "Enabled"))
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("0x%016X %-22s 0x%016X %-10d %t\n", h.Tag, h.Type, h.Address, h.PID, h.Enabled))
	}
	p.list.SetCode(sb.String())
	app.RequestRedraw()
}

// parseAddr 解析地址输入。
func (p *EventsPage) parseAddr() uint64 {
	s := strings.TrimSpace(p.addrInput.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}
	return v
}

func (p *EventsPage) parsePID() uint32 {
	s := strings.TrimSpace(p.pidInput.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	v, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		v, _ = strconv.ParseUint(s, 10, 32)
	}
	return uint32(v)
}

func (p *EventsPage) parseTag() uint64 {
	s := strings.TrimSpace(p.tagInput.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		v, _ = strconv.ParseUint(s, 10, 64)
	}
	return v
}

// ===== hook 注册方法 =====

func (p *EventsPage) registerEptHook() {
	addr := p.parseAddr()
	tag, err := p.dbg.EptHook(context.Background(), addr, defaultCallback)
	p.recordHook("EptHook", addr, 0, tag, err)
}

func (p *EventsPage) registerEptHookForProcess() {
	addr := p.parseAddr()
	pid := p.parsePID()
	tag, err := p.dbg.EptHookForProcess(context.Background(), addr, pid, defaultCallback)
	p.recordHook("EptHookForProcess", addr, pid, tag, err)
}

func (p *EventsPage) registerSyscallHook() {
	tag, err := p.dbg.SyscallHook(context.Background(), defaultCallback)
	p.recordHook("SyscallHook", 0, 0, tag, err)
}

func (p *EventsPage) registerSysretHook() {
	pid := p.parsePID()
	tag, err := p.dbg.SysretHookForProcess(context.Background(), pid, defaultCallback)
	p.recordHook("SysretHookForProcess", 0, pid, tag, err)
}

func (p *EventsPage) registerMonitorRead() {
	addr := p.parseAddr()
	pid := p.parsePID()
	tag, err := p.dbg.MonitorReadForProcess(context.Background(), addr, addr+0x100, pid, defaultCallback)
	p.recordHook("MonitorRead", addr, pid, tag, err)
}

func (p *EventsPage) registerMonitorWrite() {
	addr := p.parseAddr()
	pid := p.parsePID()
	tag, err := p.dbg.MonitorWrite(context.Background(), addr, addr+0x100, pid, defaultCallback)
	p.recordHook("MonitorWrite", addr, pid, tag, err)
}

func (p *EventsPage) registerMonitorExec() {
	addr := p.parseAddr()
	pid := p.parsePID()
	tag, err := p.dbg.MonitorExec(context.Background(), addr, pid, defaultCallback)
	p.recordHook("MonitorExec", addr, pid, tag, err)
}

func (p *EventsPage) registerCpuidHook() {
	tag, err := p.dbg.CpuidHook(context.Background(), defaultCallback)
	p.recordHook("CpuidHook", 0, 0, tag, err)
}

func (p *EventsPage) registerDrHook() {
	tag, err := p.dbg.DrHook(context.Background(), defaultCallback)
	p.recordHook("DrHook", 0, 0, tag, err)
}

func (p *EventsPage) registerVmcallHook() {
	tag, err := p.dbg.VmcallHook(context.Background(), defaultCallback)
	p.recordHook("VmcallHook", 0, 0, tag, err)
}

func (p *EventsPage) registerModeHook() {
	tag, err := p.dbg.ModeHook(context.Background(), defaultCallback)
	p.recordHook("ModeHook", 0, 0, tag, err)
}

func (p *EventsPage) registerMsrRead() {
	addr := p.parseAddr()
	tag, err := p.dbg.MsrReadHook(context.Background(), uint32(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("MsrReadHook(0x%X)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerMsrWrite() {
	addr := p.parseAddr()
	tag, err := p.dbg.MsrWriteHook(context.Background(), uint32(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("MsrWriteHook(0x%X)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerException() {
	addr := p.parseAddr()
	tag, err := p.dbg.ExceptionHook(context.Background(), uint32(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("ExceptionHook(%d)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerInterrupt() {
	addr := p.parseAddr()
	tag, err := p.dbg.InterruptHook(context.Background(), uint32(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("InterruptHook(%d)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerIoIn() {
	addr := p.parseAddr()
	tag, err := p.dbg.IoInHook(context.Background(), uint16(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("IoInHook(0x%X)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerIoOut() {
	addr := p.parseAddr()
	tag, err := p.dbg.IoOutHook(context.Background(), uint16(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("IoOutHook(0x%X)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerCrwrite() {
	addr := p.parseAddr()
	tag, err := p.dbg.CrwriteHook(context.Background(), uint32(addr), defaultCallback)
	p.recordHook(fmt.Sprintf("CrwriteHook(%d)", addr), 0, 0, tag, err)
}

func (p *EventsPage) registerXsetbv() {
	tag, err := p.dbg.XsetbvHook(context.Background(), defaultCallback)
	p.recordHook("XsetbvHook", 0, 0, tag, err)
}

func (p *EventsPage) recordHook(hookType string, addr uint64, pid uint32, tag uint64, err error) {
	if err != nil {
		fmt.Printf("[%s] 注册失败: %v\n", hookType, err)
		return
	}
	p.addHook(HookRecord{
		Tag:      tag,
		Type:     hookType,
		Address:  addr,
		PID:      pid,
		Callback: defaultCallback,
		Enabled:  true,
	})
	fmt.Printf("[%s] 已注册 tag=0x%X\n", hookType, tag)
}

// ===== 事件管理 =====

func (p *EventsPage) clearEvent() {
	tag := p.parseTag()
	if err := p.dbg.ClearEvent(context.Background(), tag); err != nil {
		fmt.Printf("ClearEvent 失败: %v\n", err)
		return
	}
	p.mu.Lock()
	for i, h := range p.hooks {
		if h.Tag == tag {
			p.hooks = append(p.hooks[:i], p.hooks[i+1:]...)
			break
		}
	}
	p.mu.Unlock()
	p.refreshList()
}

func (p *EventsPage) disableEvent() {
	tag := p.parseTag()
	if err := p.dbg.DisableEvent(context.Background(), tag); err != nil {
		fmt.Printf("DisableEvent 失败: %v\n", err)
		return
	}
	p.mu.Lock()
	for i := range p.hooks {
		if p.hooks[i].Tag == tag {
			p.hooks[i].Enabled = false
			break
		}
	}
	p.mu.Unlock()
	p.refreshList()
}

func (p *EventsPage) enableEvent() {
	tag := p.parseTag()
	if err := p.dbg.EnableEvent(context.Background(), tag); err != nil {
		fmt.Printf("EnableEvent 失败: %v\n", err)
		return
	}
	p.mu.Lock()
	for i := range p.hooks {
		if p.hooks[i].Tag == tag {
			p.hooks[i].Enabled = true
			break
		}
	}
	p.mu.Unlock()
	p.refreshList()
}

var _ = api.Breakpoint{}
