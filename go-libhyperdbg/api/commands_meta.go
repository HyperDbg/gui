// Package api — commands_meta.go
//
// 对应 debugger/commands/meta 包的 23 个 meta 命令的 typed API。
//
// 命令分组：
//
//	A) 已实装（core/commands 层有真实行为）— 本文件提供 typed 包装
//	B) stubs（commands 层只返回 ErrNotImplemented）— 本文件提供 typed 签名
//	   但方法体返回 ErrCommandNotImplemented，待 core 层补齐后替换
//
// 注意：与字符串命令路径（Exec("status")）不同，typed API 在编译期检查
// 参数类型。已实装的方法直接委托 core/registry；stub 方法仅保证签名稳定。
//
// 命令对照表（与 meta.go RegisterAll 顺序一致）：
//
//	.connect / connect  → Connect          (api.debugger.go, 已有)
//	load                → LoadVMM          (api.debugger.go, 已有)
//	unload              → UnloadVMM        (api.debugger.go, 已有)
//	g / go              → Continue         (api.debugger.go, 已有)
//	pause               → Pause            (api.debugger.go, 已有)
//	.logopen / logopen  → LogOpen          (api.debugger.go, 已有)
//	.logclose / logclose→ LogClose         (api.debugger.go, 已有)
//	.start              → StartProcess     (api.debugger.go, 已有)
//	status              → Status           (本文件)
//	cls / clear         → ClearScreen      (本文件)
//	help                → Help             (本文件)
//	exit                → Exit             (本文件)
//	attach              → Attach           (本文件, stub)
//	debug               → Debug            (本文件, stub)
//	detach              → Detach           (本文件, stub)
//	disconnect          → Disconnect       (本文件, stub)
//	dump                → Dump             (本文件, stub)
//	formats             → Formats          (本文件, stub)
//	kill                → Kill             (本文件, stub)
//	listen              → Listen           (本文件, stub)
//	pagein              → PageIn           (本文件, stub)
//	pe                  → Pe               (本文件, stub)
//	process             → Process          (本文件, stub)
//	restart             → Restart          (本文件, stub)
//	script              → Script           (本文件, stub)
//	switch              → Switch           (本文件, stub)
//	sym                 → Sym              (本文件, stub)
//	sympath             → SymPath          (本文件, stub)
//	thread              → Thread           (本文件, stub)
package api

import (
	"fmt"

	metacmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/meta"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
	"github.com/hyperdbg/go-libhyperdbg/debugger/userlevel"
)

// ============================================================
// A) 已实装命令的 typed API
// ============================================================

// Status 对应 'status' 命令：返回当前调试器状态。
func (d *Debugger) Status() (core.DebuggerState, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.State(), nil
}

// ClearScreen 对应 'cls'/'clear' 命令：向 output 写入 ANSI 清屏序列。
//
// 注意：是否真的清屏取决于 output 实现（CLI 终端会清，GUI/MCP 可能忽略）。
func (d *Debugger) ClearScreen() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.output.Write([]byte("\x1b[2J\x1b[H"))
	return err
}

// Help 对应 'help' 命令：
//   - cmdName 为空：列出所有可见命令（按 meta/debugging/extension/hwdbg 分组）
//   - cmdName 非空：打印该命令的详细帮助
//
// 帮助文本写入 d.output，返回写过程中遇到的错误。
func (d *Debugger) Help(cmdName string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// 复用 Registry 已构造的 help handler，保证行为与字符串命令路径一致。
	line := "help"
	if cmdName != "" {
		line = "help " + cmdName
	}
	return d.commands.Exec(d.core, line)
}

// Exit 对应 'exit'/'.exit'/'quit' 命令：返回 metacmds.ErrExit 信号。
//
// CLI 循环检查返回值是否为该错误以决定是否退出；GUI/MCP 调用方可忽略。
func (d *Debugger) Exit() error {
	return metacmds.ErrExit
}

// ============================================================
// B) Stubs — typed 签名 + ErrCommandNotImplemented
// ============================================================
//
// 以下方法签名参考 HyperDbg 官方命令语法（hyperdbg.com/docs/commands）。
// core 层补齐实装后，只需替换方法体（删除 return ErrCommandNotImplemented
// 行，改为委托 core 方法即可）。

// Attach 对应 'attach <pid>'：附加到已运行的进程。
func (d *Debugger) Attach(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Attach(pid)
}

// Debug 对应 'debug <exe>'：以调试模式启动一个新进程（与 .start 的区别
// 在于 debug 会等待符号加载完成才返回）。
func (d *Debugger) Debug(exePath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Debug is StartProcess + (future) symbol-load wait. The returned Process
	// handle is intentionally discarded: core.StartProcess stores the kernel
	// token internally; the caller does not need the Win32 handles here.
	_, err := d.core.StartProcess(exePath)
	return err
}

// Detach 对应 'detach'：从当前调试目标分离（不终止进程）。
func (d *Debugger) Detach() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Detach()
}

// Disconnect 对应 'disconnect'：断开与远程/local 调试目标的连接。
func (d *Debugger) Disconnect() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Disconnect()
}

// Dump 对应 'dump <path>'：将目标进程内存转储到文件（minidump 格式）。
func (d *Debugger) Dump(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return fmt.Errorf("Dump: minidump not yet implemented (requires MiniDumpWriteDump wrapper)")
}

// Formats 对应 'formats <expr>'：以多种格式（hex/dec/oct/bin）显示表达式值。
func (d *Debugger) Formats(expr string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, "formats "+expr)
}

// Kill 对应 'kill <pid>'：终止目标进程。
func (d *Debugger) Kill(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Kill(pid)
}

// Listen 对应 'listen <ip> <port>'：在指定端口监听远程调试连接。
func (d *Debugger) Listen(ip string, port int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return fmt.Errorf("Listen: remote debugging listener not yet wired (see kernellvl.KernelListener)")
}

// PageIn 对应 'pagein <addr>'：强制将指定地址的页面调入物理内存。
func (d *Debugger) PageIn(addr uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.PageIn(addr)
}

// Pe 对应 'pe <path>'：解析 PE 文件头并显示其属性（入口点/段/导入表等）。
func (d *Debugger) DumpPe(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	pf, err := userlevel.PeOpen(path)
	if err != nil {
		return err
	}
	defer pf.Close()
	d.output.Printf("path: %s\n", pf.Path())
	d.output.Printf("image base: 0x%X\n", pf.ImageBase())
	d.output.Printf("entry point: 0x%X\n", pf.EntryPoint())
	d.output.Printf("64-bit: %v\n", pf.Is64Bit())
	d.output.Printf("subsystem: %s\n", pf.SubsystemName())
	d.output.Printf("dll characteristics: %s\n", pf.DllCharacteristicsNames())
	secs := pf.Sections()
	d.output.Printf("sections: %d\n", len(secs))
	for _, s := range secs {
		d.output.Printf("  %-8s va=0x%X vsize=0x%X rsize=0x%X\n", s.Name, s.VirtAddr, s.VirtSize, s.RawSize)
	}
	imps := pf.Imports()
	d.output.Printf("imports: %d\n", len(imps))
	exps := pf.Exports()
	d.output.Printf("exports: %d\n", len(exps))
	return nil
}

// Process 对应 'process'（无参）：列出系统所有进程（类似 !process 0 0）。
func (d *Debugger) Process() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, "process")
}

// Restart 对应 'restart'：重启当前调试目标。
func (d *Debugger) Restart() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return fmt.Errorf("Restart: not supported (no exePath tracked); use Kill + StartProcess")
}

// Script 对应 'script <path>'：执行 .ds 脚本文件。
func (d *Debugger) Script(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, "script "+path)
}

// Switch 对应 'switch <pid>'：切换当前调试目标到指定进程。
func (d *Debugger) Switch(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, fmt.Sprintf("switch %d", pid))
}

// Sym 对应 'sym <name>'：解析符号名到地址（需要 SymbolResolver）。
func (d *Debugger) ResolveSymbol(name string) (uint64, error) {
	d.mu.Lock()
	resolver := d.symbols
	d.mu.Unlock()
	if resolver == nil {
		return 0, fmt.Errorf("ResolveSymbol(%q): no symbol resolver injected (use WithSymbolResolver)", name)
	}
	return resolver.FromName(name)
}

// SymPath 对应 'sympath <path>'：设置符号搜索路径。
func (d *Debugger) SymPath(path string) error {
	d.mu.Lock()
	resolver := d.symbols
	d.mu.Unlock()
	if resolver == nil {
		return fmt.Errorf("SymPath(%q): no symbol resolver injected (use WithSymbolResolver)", path)
	}
	return resolver.Init(path)
}

// Thread 对应 'thread'（无参）：列出当前调试目标的线程。
func (d *Debugger) Thread() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, "thread")
}
