package pages

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/codeeditor"
	"github.com/ddkwork/ux/widget/input"
	"github.com/ddkwork/ux/widget/split"
	"github.com/ddkwork/ux/widget/table"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// Capturer 允许页面临时捕获 debugger 命令输出（而非写入日志页）。
// ui.uiOutput 实现此接口。
type Capturer interface {
	StartCapture()
	StopCapture() string
}

// CpuPage 是 CPU 标签页，镜像 x64dbg 的 4 窗格布局：
//
//	左上 反汇编 | 右上 寄存器
//	左下 Hex   | 右下 调用栈
//
// 反汇编与 Hex 使用 typed API（Unassemble/DumpMem，直接返回数据）；
// 寄存器与调用栈走字符串命令路径（Exec("r")/Exec("k")），
// 通过 Capturer 捕获输出显示在窗格中。
type CpuPage struct {
	dbg      *api.Debugger
	capturer Capturer

	addrEditor *input.Input
	refreshBtn *button.Button

	// 反汇编改用 table 控件（地址 | 指令 | 注释），更接近 OllyDbg/x64dbg。
	disasmTbl  *table.Table
	disasmRows [][]string

	regs    *codeeditor.CodeEditor
	hexdump *codeeditor.CodeEditor
	stack   *codeeditor.CodeEditor

	spTop    split.Split
	spBottom split.Split
	sp       split.Split

	// 后台刷新的数据（线程安全，Layout 时应用到 editor）
	mu                sync.Mutex
	hasPending        bool
	pendingRegs       string
	pendingDisasmRows [][]string
	pendingHex        string
	pendingStack      string
	pendingAddr       uint64

	// refreshVersion 用于丢弃过期刷新结果：每次 Refresh 递增版本号，
	// refreshInternal 完成时检查版本号是否仍为最新，否则丢弃（避免
	// 旧刷新覆盖新刷新，例如 Continue 后 runAsync 立即刷新读到的
	// 过期数据覆盖 OnPaused 回调刷新的新鲜数据）。
	refreshVersion uint64
}

// NewCpu 创建 CPU 页。capturer 用于捕获字符串命令输出（可为 nil）。
func NewCpu(dbg *api.Debugger, capturer Capturer) *CpuPage {
	c := &CpuPage{
		dbg:        dbg,
		capturer:   capturer,
		addrEditor: input.CompactInput(),
		refreshBtn: button.Filled(),
		disasmTbl: table.New([]table.Column{
			{Name: "地址", Width: 160, MinWidth: 120},
			{Name: "指令", Width: 240, MinWidth: 160},
			{Name: "注释", Width: 160, MinWidth: 80},
		}),
		regs:    codeeditor.New("（寄存器）点击刷新", "go"),
		hexdump: codeeditor.New("（Hex dump）点击刷新", "go"),
		stack:   codeeditor.New("（调用栈）点击刷新", "go"),
	}
	c.addrEditor.Editor.SetText("0x00010000")
	c.regs.SetReadOnly(true)
	c.hexdump.SetReadOnly(true)
	c.stack.SetReadOnly(true)

	spTop := split.Split{
		Ratio: 0.7, Bar: 10, Axis: layout.Horizontal,
		First:  c.layoutDisasm,
		Second: c.regs.Layout,
	}
	spBottom := split.Split{
		Ratio: 0.5, Bar: 10, Axis: layout.Horizontal,
		First:  c.hexdump.Layout,
		Second: c.stack.Layout,
	}
	sp := split.Split{
		Ratio: 0.7, Bar: 10, Axis: layout.Vertical,
		First:  spTop.Layout,
		Second: spBottom.Layout,
	}
	c.spTop = spTop
	c.spBottom = spBottom
	c.sp = sp
	return c
}

// layoutDisasm 渲染反汇编表格。disasmRows 由 applyPending 在 UI 线程更新。
func (c *CpuPage) layoutDisasm(gtx layout.Context) layout.Dimensions {
	rows := c.disasmRows
	c.disasmTbl.SetColumns(gtx, c.disasmTbl.Columns, rows)
	return table.SimpleTable(gtx, c.disasmTbl, len(rows), func(gtx layout.Context, row, col int) layout.Dimensions {
		if row < 0 || row >= len(rows) {
			return layout.Dimensions{}
		}
		cell := ""
		if col < len(rows[row]) {
			cell = rows[row][col]
		}
		return wdk.BodyM(gtx, cell)
	})
}

// parseDisasmRows 将 Unassemble 的输出解析为 [地址, 指令, 注释] 三列。
// 输出格式: "000000007FF41234  mov rax, rsp\n..."
// 注释列暂留空（无符号解析）。
func parseDisasmRows(text string) [][]string {
	if text == "" {
		return nil
	}
	lines := strings.Split(text, "\n")
	rows := make([][]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n ")
		if line == "" {
			continue
		}
		// 期望格式: 16 hex digits + 2 spaces + instruction
		addr, instr := line, ""
		if len(line) >= 18 && line[16] == ' ' {
			addr = line[:16]
			instr = strings.TrimLeft(line[16:], " ")
		}
		rows = append(rows, []string{addr, instr, ""})
	}
	return rows
}

func (c *CpuPage) Layout(gtx layout.Context) layout.Dimensions {
	if c.refreshBtn.Clicked(gtx) {
		go c.Refresh()
	}
	// 应用后台刷新的数据到 editor（在 UI 线程中安全更新）
	c.applyPending()
	// 工具栏：地址输入 + 刷新按钮
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Dp(200)
							return c.addrEditor.Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return c.refreshBtn.Layout(gtx, "刷新")
						}),
					)
				},
			)
		}),
		layout.Flexed(1, c.sp.Layout),
	)
}

// applyPending 将后台刷新的数据应用到 editor（必须在 UI 线程调用）。
func (c *CpuPage) applyPending() {
	c.mu.Lock()
	if !c.hasPending {
		c.mu.Unlock()
		return
	}
	regs := c.pendingRegs
	disasmRows := c.pendingDisasmRows
	hexT := c.pendingHex
	stack := c.pendingStack
	addr := c.pendingAddr
	c.hasPending = false
	c.mu.Unlock()

	if regs != "" {
		c.regs.SetCode(regs)
	}
	if disasmRows != nil {
		c.disasmRows = disasmRows
	}
	if hexT != "" {
		c.hexdump.SetCode(hexT)
	}
	if stack != "" {
		c.stack.SetCode(stack)
	}
	if addr != 0 {
		c.addrEditor.Editor.SetText(fmt.Sprintf("0x%X", addr))
	}
}

// Refresh 在 goroutine 中拉取 4 窗格数据并更新编辑器内容。
// 优先使用当前 RIP 作为反汇编/Hex dump 地址（模拟 OllyDbg 行为），
// RIP 解析失败时回退到地址输入框的值。
// 使用 5 秒超时防止 IOCTL 阻塞导致 UI 冻结。
func (c *CpuPage) Refresh() {
	c.mu.Lock()
	c.refreshVersion++
	version := c.refreshVersion
	c.mu.Unlock()
	go c.refreshInternal(version)
}

// refreshInternal 实际执行数据拉取，在后台 goroutine 中运行。
func (c *CpuPage) refreshInternal(version uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. 用 Register API 读取所有寄存器（输出到 capturer）
	var regText string
	if c.capturer != nil {
		c.capturer.StartCapture()
	}
	_, regErr := c.dbg.Register(ctx, "")
	if c.capturer != nil {
		regText = c.capturer.StopCapture()
	}

	// 2. 解析 RIP，作为反汇编/Hex 地址
	addr := parseRIP(regText)
	if addr == 0 {
		addr = c.parseAddr() // 回退到地址输入框
	}

	// 3. 反汇编（typed 返回，解析为表格行）
	var disasmRows [][]string
	if regErr != nil {
		disasmRows = [][]string{{"", fmt.Sprintf("寄存器读取失败: %v", regErr), ""}}
	} else if disasm, err := c.dbg.Unassemble(ctx, addr, 20); err == nil && disasm != "" {
		disasmRows = parseDisasmRows(disasm)
	} else if err != nil {
		disasmRows = [][]string{{"", fmt.Sprintf("反汇编失败: %v", err), ""}}
	}

	// 4. Hex dump（typed 返回）
	var hexText string
	if data, err := c.dbg.DumpMem(ctx, addr, 256); err == nil && len(data) > 0 {
		hexText = hex.Dump(data)
	} else if err != nil {
		hexText = fmt.Sprintf("DumpMem 失败: %v", err)
	}

	// 5. 调用栈（用 K API，输出到 capturer）
	var stackText string
	if c.capturer != nil {
		c.capturer.StartCapture()
	}
	_, _ = c.dbg.K(ctx, 16)
	if c.capturer != nil {
		stackText = c.capturer.StopCapture()
	}

	// 6. 在 UI 线程中更新 widget 状态（通过 RequestRedraw 触发）
	//    先存到临时字段，Layout 时再更新 editor。
	//    检查版本号：如果期间有新的 Refresh 调用，丢弃当前过期结果。
	c.mu.Lock()
	if version != c.refreshVersion {
		c.mu.Unlock()
		return
	}
	c.pendingRegs = regText
	c.pendingDisasmRows = disasmRows
	c.pendingHex = hexText
	c.pendingStack = stackText
	c.pendingAddr = addr
	c.hasPending = true
	c.mu.Unlock()

	app.RequestRedraw()
}

// parseRIP 从 "r" 命令输出中解析 RIP 值。
// 输出格式示例: "RIP=000000007FF41234 RFL=..." 或 "rip=7ff41234" 等。
func parseRIP(regText string) uint64 {
	if regText == "" {
		return 0
	}
	for _, line := range strings.Split(regText, "\n") {
		line = strings.TrimSpace(line)
		lower := strings.ToLower(line)
		// 查找 "rip=" 或 "rip:"
		for _, sep := range []string{"rip=", "rip:", "rip ="} {
			idx := strings.Index(lower, sep)
			if idx < 0 {
				continue
			}
			rest := line[idx+len(sep):]
			rest = strings.TrimSpace(rest)
			rest = strings.TrimPrefix(rest, "0x")
			rest = strings.TrimPrefix(rest, "0X")
			// 取第一个 token（到空格/逗号/tab 为止）
			for i, ch := range rest {
				if ch == ' ' || ch == ',' || ch == '\t' || ch == '\n' {
					rest = rest[:i]
					break
				}
			}
			v, err := strconv.ParseUint(rest, 16, 64)
			if err == nil {
				return v
			}
		}
	}
	return 0
}

func (c *CpuPage) parseAddr() uint64 {
	s := strings.TrimSpace(c.addrEditor.Editor.GetText())
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0x10000
	}
	return v
}
