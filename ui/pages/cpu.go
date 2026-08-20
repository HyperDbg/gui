package pages

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/codeeditor"
	"github.com/ddkwork/ux/widget/split"
	"github.com/ddkwork/ux/widget/table"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// CpuPage 是 CPU 标签页，镜像 x64dbg 的 4 窗格布局：
//
//	左上 反汇编 | 右上 寄存器
//	左下 Hex   | 右下 调用栈
//
// 全部使用 typed API：Unassemble/DumpMem 直接返回数据，
// Register("RIP") 返回 uint64，CallStack(16) 返回 []CallFrame。
type CpuPage struct {
	dbg *api.Debugger

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
	pendingRIP        uint64 // 本次刷新读取到的 RIP（用于计算选中行）

	// refreshVersion 用于丢弃过期刷新结果：每次 Refresh 递增版本号，
	// refreshInternal 完成时检查版本号是否仍为最新，否则丢弃（避免
	// 旧刷新覆盖新刷新，例如 Continue 后 runAsync 立即刷新读到的
	// 过期数据覆盖 OnPaused 回调刷新的新鲜数据）。
	refreshVersion uint64

	// currentRIP 是当前已应用到 UI 的 RIP 值，用于 layoutDisasm
	// 计算反汇编表格的选中行（OllyDbg 风格：高亮当前指令所在行）。
	currentRIP uint64
}

// NewCpu 创建 CPU 页。
func NewCpu(dbg *api.Debugger) *CpuPage {
	c := &CpuPage{
		dbg: dbg,
		disasmTbl: table.New([]table.Column{
			{Name: "地址", Width: 160, MinWidth: 120},
			{Name: "指令", Width: 240, MinWidth: 160},
			{Name: "注释", Width: 160, MinWidth: 80},
		}),
		regs:    codeeditor.New("（寄存器）点击刷新", "go"),
		hexdump: codeeditor.New("（Hex dump）点击刷新", "go"),
		stack:   codeeditor.New("（调用栈）点击刷新", "go"),
	}
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
// 当前 RIP 所在行被高亮选中（OllyDbg 风格）。
// 由于 Unassemble 从 currentRIP 开始反汇编，RIP 行通常为第 0 行。
func (c *CpuPage) layoutDisasm(gtx layout.Context) layout.Dimensions {
	rows := c.disasmRows

	// 根据 currentRIP 查找选中行索引
	ripStr := fmt.Sprintf("%016X", c.currentRIP)
	selectedRow := -1
	for i, row := range rows {
		if len(row) > 0 && strings.EqualFold(row[0], ripStr) {
			selectedRow = i
			break
		}
	}
	c.disasmTbl.SelectedRow = selectedRow

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
	// 应用后台刷新的数据到 editor（在 UI 线程中安全更新）
	c.applyPending()
	return c.sp.Layout(gtx)
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
	rip := c.pendingRIP
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
	if rip != 0 {
		c.currentRIP = rip
	}
}

// Refresh 在 goroutine 中拉取 4 窗格数据并更新编辑器内容。
// 优先使用当前 RIP 作为反汇编/Hex dump 地址（模拟 OllyDbg 行为），
// RIP 读取失败时回退到固定地址 0x10000。
func (c *CpuPage) Refresh() {
	c.mu.Lock()
	c.refreshVersion++
	version := c.refreshVersion
	c.mu.Unlock()
	go c.refreshInternal(version)
}

// refreshInternal 实际执行数据拉取，在后台 goroutine 中运行。
func (c *CpuPage) refreshInternal(version uint64) {
	// 1. 读取所有寄存器（文本），并解析 RIP 地址
	regText, regErr := c.dbg.AllRegisters()
	addr := parseRIP(regText)
	if regErr != nil || addr == 0 {
		addr = 0x10000 // RIP 读取失败时回退到固定地址
	}

	// 2. 反汇编（typed 返回，解析为表格行）
	var disasmRows [][]string
	if regErr != nil {
		disasmRows = [][]string{{"", fmt.Sprintf("寄存器读取失败: %v", regErr), ""}}
	} else if disasm, err := c.dbg.Unassemble(addr, 20); err == nil && disasm != "" {
		disasmRows = parseDisasmRows(disasm)
	} else if err != nil {
		disasmRows = [][]string{{"", fmt.Sprintf("反汇编失败: %v", err), ""}}
	}

	// 3. Hex dump（typed 返回）
	var hexText string
	if data, err := c.dbg.DumpMem(addr, 256); err == nil && len(data) > 0 {
		hexText = hex.Dump(data)
	} else if err != nil {
		hexText = fmt.Sprintf("DumpMem 失败: %v", err)
	}

	// 4. 调用栈（K 返回 []CallFrame，格式化为文本显示）
	var stackText string
	if frames, err := c.dbg.CallStack(16); err == nil && len(frames) > 0 {
		var sb strings.Builder
		for _, f := range frames {
			fmt.Fprintf(&sb, "%016X  %s\n", f.Address, f.Symbol)
		}
		stackText = sb.String()
	} else if err != nil {
		stackText = fmt.Sprintf("K 失败: %v", err)
	}

	// 5. 在 UI 线程中更新 widget 状态（通过 RequestRedraw 触发）
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
	c.pendingRIP = addr
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
	for line := range strings.SplitSeq(regText, "\n") {
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
