package ui

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/ui/pages"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/appbar"
	"github.com/ddkwork/ux/widget/panel"
	"github.com/ddkwork/ux/widget/tab"
	"github.com/ddkwork/ux/widget/toolbar"
	"github.com/ddkwork/ux/widget/tooltip"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

//go:embed asserts/ico/ico_aaamain.png
var mainIcons []byte

//go:embed asserts/bar/*.png
var bar embed.FS

//go:embed asserts/pageico/*.png
var pageIco embed.FS

type TabPageType int

const (
	CpuType TabPageType = iota
	PeViewType
	LogType
	NotesType
	BreaksType
	MemoryType
	SehType
	ScriptType
	SymbolType
	SourceType
	ReferencesType
	ThreadType
	HandleType
	TraceType
	ArkType
	ScyllaType
	LabelsType
	CommentsType
	FunctionsType
	XrefsType
	TypesType
	WatchesType
	GraphsType
	ExceptionsType
	BookmarksType
	LoopsType
	EventsType
)

func (t TabPageType) String() string {
	switch t {
	case CpuType:
		return "CPU"
	case PeViewType:
		return "PE View"
	case LogType:
		return "Log"
	case NotesType:
		return "Notes"
	case BreaksType:
		return "Breakpoints"
	case MemoryType:
		return "Memory"
	case SehType:
		return "SEH"
	case ScriptType:
		return "Script"
	case SymbolType:
		return "Symbols"
	case SourceType:
		return "Source"
	case ReferencesType:
		return "References"
	case ThreadType:
		return "Threads"
	case HandleType:
		return "Handles"
	case TraceType:
		return "Trace"
	case ArkType:
		return "ARK"
	case ScyllaType:
		return "Scylla"
	case LabelsType:
		return "Labels"
	case CommentsType:
		return "Comments"
	case FunctionsType:
		return "Functions"
	case XrefsType:
		return "Xrefs"
	case TypesType:
		return "Types"
	case WatchesType:
		return "Watches"
	case GraphsType:
		return "Graphs"
	case ExceptionsType:
		return "Exceptions"
	case BookmarksType:
		return "Bookmarks"
	case LoopsType:
		return "Loops"
	case EventsType:
		return "Events"
	default:
		return "Unknown"
	}
}

func (t TabPageType) EnumTypes() []TabPageType {
	return []TabPageType{
		CpuType, PeViewType, LogType, NotesType, BreaksType, MemoryType,
		SehType, ScriptType, SymbolType, SourceType, ReferencesType,
		ThreadType, HandleType, TraceType, ArkType, ScyllaType,
		LabelsType, CommentsType, FunctionsType, XrefsType, TypesType,
		WatchesType, GraphsType, ExceptionsType, BookmarksType, LoopsType,
		EventsType,
	}
}

var (
	myAppBar    = appbar.New()
	mainToolbar = toolbar.New()
)

var (
	globalDbg        *api.Debugger
	globalOut        *uiOutput
	logPage          *pages.LogPage
	cpuPage          *pages.CpuPage
	eventsPage       *pages.EventsPage
	breaksPage       *pages.BreaksPage
	modulesPage      *pages.ModulesPage
	vtab             *tab.TabView
	enableConsoleLog = true
)

// 运行时状态（供工具栏回调使用）
var (
	currentExePath string                // 当前加载的进程路径（用于 restart）
	messagePump    *core.MessagePump     // 内核日志消息泵（hook 回调输出）
	globalProc     *core.Process         // 当前调试进程（退出时需 Terminate + Close）
	driverPath     = "Debug/hyperhv.sys" // VMM 驱动路径（可由 SetDriverPath 覆盖）
	vmmLoaded      bool                  // VMM 驱动是否已加载（避免重复 LoadVMM）
	isLoading      bool                  // loadProcess 正在进行中，工具栏按钮应忽略
)

// SetDriverPath 设置 VMM 驱动 .sys 文件路径（用于内嵌释放后的路径传递）。
func SetDriverPath(p string) {
	driverPath = p
}

func init() {
	myAppBar.Title("").HeightLevel(appbar.HeightSmall)
}

func SetEnableConsoleLog(enable bool) {
	enableConsoleLog = enable
}

// statusf 同时输出到控制台与日志页，用于工具栏/命令的 UI 反馈。
func statusf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(msg)
	if logPage != nil {
		logPage.Println(msg)
	}
}

// Run 启动 HyperDbg UI。onReady 在 debugger 构造完成后、UI 循环开始前调用。
func Run(onReady func(*api.Debugger)) {
	fmt.Println("=== Starting HyperDbg UI ===")

	// 先创建日志页与输出适配器，再构造 debugger（命令输出路由到日志页）
	logPage = pages.NewLog()
	globalOut = newUIOutput(logPage, os.Stdout)

	var err error
	globalDbg, err = api.New(api.WithOutput(globalOut))
	if err != nil {
		fmt.Printf("Failed to create debugger: %v\n", err)
		return
	}

	// 注册暂停事件回调：所有暂停事件（断点命中、单步完成、OEP、
	// 手动 Pause、Continue 后再次暂停）都会自动刷新 CPU 页。
	// 回调运行在 MessagePump goroutine 中；autoRefreshCPU →
	// cpuPage.Refresh() 本身是 go refreshInternal()，不会阻塞 pump。
	globalDbg.SetOnPaused(func() {
		autoRefreshCPU()
	})

	// ExitCallback 在 Gio 的 DestroyEvent 中同步调用（goroutine 内）。
	// 但 app.Main() 可能在 goroutine 的 os.Exit(0) 之前返回，
	// 导致主 goroutine 直接退出而跳过 cleanup。defer 作为兜底：
	// 若 app.Main() 先返回，defer 会执行 cleanup（cleanupDebugger
	// 对 nil 安全，双重调用不会出问题）。
	app.ExitCallback(func() {
		fmt.Println("=== UI Exited, cleanup start ===")
		cleanupDebugger()
		fmt.Println("=== cleanup done ===")
	})
	// defer cleanupDebugger()

	runUI(onReady)
}

// cleanupDebugger 按照与 themida Unpacker.Run 相同的顺序释放资源：
//  1. 停止消息泵（必须在 UnloadVMM 前：Stop 需要主设备句柄发送
//     IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL，而 UnloadVMM
//     会停止驱动服务使该句柄失效）
//  3. UnloadVMM（内部顺序：清事件 → continue+detach → TERMINATE_VMX →
//     卸载驱动服务）
//  2. TerminateProcess（VMX 关闭后再杀调试进程，避免进程在 hook 已拆除
//     但 VMX 仍开时 fault）
//  4. proc.Close + dbg.Close
//
// 全程使用 context.Background()（不设超时）：终止 VMX 是慢操作，带超时的
// ctx 会让 IOCTL 立即失败，导致驱动残留 + VT-x 卡死（StopPending），
// 只能重启恢复。错误只打印不返回——退出路径必须尽力走完所有步骤。
func cleanupDebugger() {
	// 1. 停消息泵
	if messagePump != nil {
		fmt.Println("[cleanup] stopping message pump...")
		messagePump.Stop()
		messagePump = nil
	}

	// 2. 终止调试进程（UnloadVMM 已 detach，进程可能仍在跑）
	if globalProc != nil {
		fmt.Println("[cleanup] terminating debuggee...")
		_ = globalProc.Terminate()
		_ = globalProc.Close()
		globalProc = nil
	}

	// 3. UnloadVMM（TERMINATE_VMX + 卸载驱动）
	if globalDbg != nil && vmmLoaded {
		fmt.Println("[cleanup] UnloadVMM...")
		if err := globalDbg.UnloadVMM(context.Background()); err != nil {
			fmt.Printf("[cleanup] UnloadVMM error: %v\n", err)
		}
		vmmLoaded = false
	}

	// 4. 关闭 debugger（设备句柄 + 日志文件）
	if globalDbg != nil {
		fmt.Println("[cleanup] closing debugger...")
		_ = globalDbg.Close()
		globalDbg = nil
	}
}

// RunDriverOnly 仅加载 VMM 驱动，不启动 GUI（控制台模式）。
func RunDriverOnly() {
	fmt.Println("=== Driver Only Mode ===")

	logFilePath := "hyperdbg_driver_only.log"
	fmt.Printf("日志文件: %s\n", logFilePath)

	dbg, err := api.New(api.WithOutput(os.Stdout))
	if err != nil {
		fmt.Printf("Failed to create debugger: %v\n", err)
		return
	}
	globalDbg = dbg

	ctx := context.Background()

	driverPath := filepath.Join("Debug", "hyperhv.sys")
	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		fmt.Printf("LoadVMM failed: %v\n", err)
	} else {
		vmmLoaded = true
		fmt.Println("=== Driver Loaded ===")
	}

	app.ExitCallback(func() {
		fmt.Println("\n=== Unloading Driver ===")
		cleanupDebugger()
		fmt.Println("=== Driver Unloaded ===")
	})

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

	cleanupDebugger()
	fmt.Println("=== Driver Unloaded ===")
}

// switchTab 切换到指定 tab（用于工具栏按钮）。
func switchTab(t TabPageType) {
	if vtab != nil {
		vtab.SetSelectedByID(t.String())
		app.RequestRedraw()
	}
}

func runUI(onReady func(*api.Debugger)) {
	p := panel.New()
	hPanel := panel.NewHPanel()
	p.AddChild(hPanel)

	logFilePath := "hyperdbg_ui.log"
	fmt.Printf("日志文件: %s\n", logFilePath)
	_ = globalDbg.LogOpen(logFilePath)

	NewToolbar(hPanel, globalDbg)

	if onReady != nil {
		onReady(globalDbg)
	}

	app.FileDropCallback(func(files []string) {
		if len(files) > 0 {
			exePath := files[0]
			go loadProcess(exePath)
		}
	})

	// 创建各页面
	cpuPage = pages.NewCpu(globalDbg, globalOut)
	eventsPage = pages.NewEvents(globalDbg)
	breaksPage = pages.NewBreaks(globalDbg)
	modulesPage = pages.NewModules(globalDbg)

	vtab = tab.New(layout.Horizontal)
	for _, tabType := range CpuType.EnumTypes() {
		var content layout.Widget
		switch tabType {
		case CpuType:
			content = cpuPage.Layout
		case LogType:
			content = logPage.Layout()
		case EventsType:
			content = eventsPage.Layout()
		case BreaksType:
			content = breaksPage.Layout()
		case ThreadType:
			// Threads tab 复用 Modules 页（执行 thread 命令）
			content = modulesPage.Layout()
		default:
			content = placeholderLayout
		}
		item := tab.NewTabItem(tabType.String(), content)
		item.SetIdentifier(tabType.String())
		vtab.AddTab(item)
	}
	p.AddChild(vtab)

	app.Run("HyperDbg", func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(myAppBar.Layout),
			layout.Flexed(1, p.Layout),
		)
		myAppBar.LayoutOverlays(gtx)
	})
}

// stopCurrentDebuggee 同步清理当前调试进程：detach → terminate → close。
// 用于 loadProcess/restart/close 按钮在启动新进程前清理旧状态。
// 不卸载 VMM 驱动（保留给下次调试），仅清理进程级状态。
func stopCurrentDebuggee() {
	if globalProc == nil {
		return
	}
	ctx := context.Background()
	// 1. detach（让内核释放调试会话，拆除 PEB 监控 hook）
	if globalDbg != nil {
		if err := globalDbg.Detach(ctx); err != nil {
			statusf("stopCurrentDebuggee: Detach 失败: %v", err)
		}
	}
	// 2. terminate 进程（detach 后进程可能仍在运行）
	_ = globalProc.Terminate()
	_ = globalProc.Close()
	globalProc = nil
	statusf("已停止调试进程")
}

// loadProcess 加载 VMM 驱动（仅首次）并启动目标进程，同时启动消息泵接收 hook 回调。
func loadProcess(exePath string) {
	ctx := context.Background()
	currentExePath = exePath
	isLoading = true
	defer func() { isLoading = false }()
	statusf("加载进程: %s", exePath)

	// 清理旧调试进程（避免句柄泄漏 + 内核调试会话残留）
	stopCurrentDebuggee()

	// 仅首次加载 VMM 驱动；重复 LoadVMM 会因驱动服务已存在而失败
	if !vmmLoaded {
		if err := globalDbg.LoadVMM(ctx, driverPath); err != nil {
			statusf("LoadVMM 失败: %v", err)
			return
		}
		vmmLoaded = true
		statusf("VMM 驱动已加载")
	}

	// 启动消息泵（hook 回调输出到日志文件）
	if messagePump == nil {
		if pump, err := globalDbg.StartMessagePump(ctx); err != nil {
			statusf("StartMessagePump 失败: %v", err)
		} else {
			messagePump = pump
			statusf("消息泵已启动")
		}
	}

	if proc, err := globalDbg.StartProcess(ctx, exePath); err != nil {
		statusf("启动进程失败: %v", err)
		return
	} else {
		// 保存进程句柄，退出时需要 Terminate + Close（否则调试进程残留，
		// UnloadVMM 的 detach/TERMINATE_VMX 会卡死，导致驱动 StopPending）。
		globalProc = &proc
	}
	// StartProcess 的 ResumeThread 后，内核在第一条指令拦截并发 PAUSED 包，
	// 但 StartProcess 不等 PAUSED 就返回——此时 pausedRIP=0。
	// 等待 PAUSED 到达（MessagePump 设置 pausedRIP），否则 StepOver 的
	// detectCallAtPausedRip 会因 pausedRIP=0 失败，退化为 STEP_IN 后卡死。
	time.Sleep(2 * time.Second)
	pauseCtx, pauseCancel := context.WithTimeout(ctx, 5*time.Second)
	if err := globalDbg.Pause(pauseCtx); err != nil {
		statusf("初始 Pause: %v", err)
	}
	pauseCancel()
	statusf("进程已启动: %s（已暂停）", exePath)
	// 进程停在 OEP，自动刷新 CPU 页显示入口点的反汇编/寄存器/堆栈/内存
	autoRefreshCPU()
	switchTab(CpuType)
}

// placeholderLayout 渲染空白占位（未实装的 tab 页）。
func placeholderLayout(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{Size: gtx.Constraints.Min}
}

// aboutText 是 About 对话框的内容，包含项目信息、核心特性、技术栈与链接。
const aboutText = `HyperDbg v1.0

A hypervisor-based debugger for Windows.

基于 Intel VT-x / EPT 的内核级调试器，支持硬件断点、内存监控、
syscall hook、EPT hook、MSR hook 等，无需修改目标二进制即可
在 VMX-root 模式下拦截并控制任意指令的执行。

Core Features:
  • VMM 驱动 (hyperhv.sys) — Intel VT-x 虚拟化层
  • EPT Hook — 隐藏断点，对目标进程透明
  • !syscall / !sysret — 系统调用拦截与过滤
  • !monitor r/w/e — 内存读写执行监控
  • !epthook2 — 无痕迹函数钩子
  • Go AST 回调 — 内核态 Go 子集解释器执行 hook 脚本
  • 符号解析 — PDB 自动下载与地址反查

Tech Stack:
  • Driver: C/C++ (HyperDbg/hyperhv)
  • User-mode Library: Go (go-libhyperdbg)
  • GUI: Gio + Material Design 3
  • Script Engine: Go subset interpreter (in-kernel)

GitHub: https://github.com/HyperDbg
License: GPL-3.0

Made with Gio + go-libhyperdbg`

func NewToolbar(hpanel *panel.Panel, dbg *api.Debugger) {
	m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	mainToolbar.Items(toolbarItems(m, dbg))

	// 把 toolbar 作为一个靠左的 action 塞进 appbar，
	// tooltip 包裹在外层，hover 时显示气泡。
	// PositionBelow：appbar 在窗口顶部，气泡往下弹避免被裁剪。
	// AnchorX：锚定到当前 hover item 的中心（由 toolbar 记录），
	// 气泡跟随 item 而不是停在 toolbar 中点。
	// GetHoveredTooltip/GetHoveredItemOffset 拿上一帧的状态（toolbar.Layout
	// 内部才更新 hover），一帧延迟用户无感。
	myAppBar.AddAction(appbar.Action{
		Widget: func(gtx layout.Context) layout.Dimensions {
			tipText := mainToolbar.GetHoveredTooltip()
			itemX, itemW := mainToolbar.GetHoveredItemOffset()
			tip := tooltip.New(tipText)
			tip.Position = tooltip.PositionBelow
			if itemW > 0 {
				tip.AnchorX = itemX + itemW/2
			}
			return tip.Layout(gtx, mainToolbar.Layout, tipText != "")
		},
		AlignLeft: true,
	})

	myAppBar.AddAction(appbar.SettingsAction())
	myAppBar.AddAction(appbar.ThemeToggleAction())
	myAppBar.AddAction(appbar.AboutAction(aboutText))
}

// runAsync 在 goroutine 中执行 debugger 命令并反馈错误到状态栏。
// 命令完成后自动刷新 CPU 页。与 OnPaused 回调的刷新通过版本号去重：
// 过期刷新（如 Continue 后进程仍在运行时读到的错误数据）会被
// CpuPage.refreshVersion 机制丢弃，不会覆盖 OnPaused 刷新的新鲜数据。
// 如果 isLoading=true（loadProcess 正在进行中），忽略操作避免竞态。
func runAsync(name string, fn func(ctx context.Context) error) {
	if isLoading {
		statusf("%s: 进程加载中，请稍候", name)
		return
	}
	if globalProc == nil {
		statusf("%s: 无调试进程（请拖入 exe）", name)
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		if err := fn(ctx); err != nil {
			statusf("%s 失败: %v", name, err)
		} else {
			statusf("%s 完成", name)
		}
		// 命令完成后刷新 CPU 页：
		// - Pause 后进程已暂停，立即刷新可得正确数据
		// - Continue/Step 后进程可能尚未暂停，刷新读到错误数据，
		//   但 OnPaused 回调会触发新版本刷新覆盖（版本号去重）
		// - lm/process/thread 等查询命令不改变进程状态，刷新无害
		autoRefreshCPU()
	}()
}

// autoRefreshCPU 刷新 CPU 页面。在进程暂停（断点命中/单步完成/Continue 返回）时调用。
func autoRefreshCPU() {
	if cpuPage != nil {
		cpuPage.Refresh()
	}
}

func toolbarItems(m *safemap.M[string, []byte], dbg *api.Debugger) []*toolbar.Item {
	item := func(icon, tip string, onClick func()) *toolbar.Item {
		return &toolbar.Item{
			Icon:    wdk.RequireIconWidget(m.GetMust(icon)),
			Tooltip: tip,
			OnClick: onClick,
		}
	}
	return []*toolbar.Item{
		item("open.png", "open", func() {
			statusf("请将 exe 文件拖入窗口以加载进程")
		}),
		item("restart.png", "restart", func() {
			if currentExePath == "" {
				statusf("restart: 无已加载进程")
				return
			}
			// 同步顺序执行：先停止旧进程，再启动新进程。
			// 不能并发——detach 和 StartProcess 共用同一设备句柄，
			// 并发会导致 IOCTL 竞争 + 内核调试会话状态混乱。
			go func() {
				stopCurrentDebuggee()
				loadProcess(currentExePath)
			}()
		}),
		item("close.png", "close", func() {
			// 停止调试进程（detach + terminate），保留 VMM 驱动供下次调试。
			go func() {
				stopCurrentDebuggee()
				autoRefreshCPU()
			}()
		}),
		item("run.png", "run", func() { runAsync("Continue", dbg.Continue) }),
		item("runthread.png", "runthread", func() { runAsync("Continue", dbg.Continue) }),
		item("pause.png", "pause", func() { runAsync("Pause", dbg.Pause) }),
		item("stepin.png", "stepin", func() { runAsync("TraceInto", dbg.TraceInto) }),
		item("stepover.png", "stepover", func() { runAsync("StepOver", dbg.StepOver) }),
		item("trin.png", "trin", func() { runAsync("TraceInto", dbg.TraceInto) }),
		item("trover.png", "trover", func() { runAsync("StepOver", dbg.StepOver) }),
		item("tillret.png", "tillret", func() { runAsync("Gu", dbg.Gu) }),
		item("tilluser.png", "tilluser", func() { statusf("tilluser: 暂未实装") }),
		item("log.png", "log", func() { switchTab(LogType) }),
		item("modules.png", "modules", func() {
			runAsync("lm", func(ctx context.Context) error { return dbg.Exec(ctx, "lm") })
			switchTab(LogType)
		}),
		item("windows.png", "windows", func() {
			runAsync("process", func(ctx context.Context) error { return dbg.Exec(ctx, "process") })
			switchTab(LogType)
		}),
		item("threads.png", "threads", func() {
			runAsync("thread", func(ctx context.Context) error { return dbg.Exec(ctx, "thread") })
			switchTab(LogType)
		}),
		item("cpu.png", "cpu", func() { switchTab(CpuType) }),
		item("search.png", "search", func() {
			statusf("search: 通过 Go API dbg.Search() 调用")
			switchTab(LogType)
		}),
		item("trace.png", "trace", func() { runAsync("Trace", dbg.Trace) }),
		item("bpoints.png", "bpoints", func() { switchTab(BreaksType) }),
		item("bpmem.png", "bpmem", func() {
			statusf("bpmem: 通过 Go API dbg.Monitor() 调用")
			switchTab(EventsType)
		}),
		item("bphard.png", "bphard", func() {
			statusf("bphard: 通过 Go API dbg.DrHook() 调用")
			switchTab(EventsType)
		}),
		item("options.png", "options", func() { statusf("options: 暂未实装") }),
		item("scylla.png", "scylla", func() { statusf("scylla: 暂未实装") }),
		item("about.png", "about", func() { statusf("HyperDbg v1.0 — 基于 hypervisor 的调试器") }),
		item("settings.png", "settings", func() {
			runAsync("settings", func(ctx context.Context) error {
				s, err := dbg.Settings(ctx)
				if err != nil {
					return err
				}
				statusf("状态: %v", s.State)
				return nil
			})
		}),
	}
}

func registerContextMenu(enable bool) {
	remove := ""
	if !enable {
		remove = "-"
	}
	path := stream.RunDir()
	path += string(filepath.Separator)
	g := stream.NewGeneratedFile()
	g.P("Windows Registry Editor Version 5.00")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\HyperDbg]")
	g.P("@=\"Run HyperDbg Here\"")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe --cd=\\\"%V\\\"\"")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\Directory\\shell\\HyperDbg]")
	g.P("@=\"Run HyperDbg Here\"")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\Directory\\shell\\HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe --cd=%V\"")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\*\\shell\\Open with HyperDbg]")
	g.P("")

	g.P("[", remove, "HKEY_CLASSES_ROOT\\*\\shell\\Open with HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe \\\"%1\\\"\"")
	g.P("")

	stream.WriteTruncate("open.reg", g.Buffer)
	stream.RunCommand("reg import open.reg")
}
