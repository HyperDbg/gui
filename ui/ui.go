package ui

import (
	"embed"
	"fmt"
	"iter"
	"path/filepath"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/appbar"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/panel"
	"github.com/ddkwork/ux/widget/tab"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/ui/pages"
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
		CpuType,
		PeViewType,
		LogType,
		NotesType,
		BreaksType,
		MemoryType,
		SehType,
		ScriptType,
		SymbolType,
		SourceType,
		ReferencesType,
		ThreadType,
		HandleType,
		TraceType,
		ArkType,
		ScyllaType,
		LabelsType,
		CommentsType,
		FunctionsType,
		XrefsType,
		TypesType,
		WatchesType,
		GraphsType,
		ExceptionsType,
		BookmarksType,
		LoopsType,
		EventsType,
	}
}

type TipIconButton struct {
	btn      *button.Button
	iconData []byte
	tip      string
	callback func()
}

func NewTooltipButton(iconData []byte, tip string, callback func()) *TipIconButton {
	return &TipIconButton{
		btn:      button.Text(),
		iconData: iconData,
		tip:      tip,
		callback: callback,
	}
}

func (t *TipIconButton) Layout(gtx layout.Context) layout.Dimensions {
	if t.btn.Clicked(gtx) && t.callback != nil {
		t.callback()
	}
	return t.btn.LayoutIconOnly(gtx, t.tip, wdk.RequireIconWidget(t.iconData))
}

var myAppBar = appbar.New()

var (
	globalDbg        *debugger.HyperDbg
	logPage          *pages.LogPage
	enableConsoleLog = true
)

func init() {
	myAppBar.Title("").HeightLevel(appbar.HeightSmall)
}

func SetEnableConsoleLog(enable bool) {
	enableConsoleLog = enable
}

func Run(onReady func(*debugger.HyperDbg)) {
	fmt.Println("=== Starting HyperDbg UI ===")

	app.ExitCallback(func() {
		fmt.Println("=== UI Exited, Unloading Driver ===")
		if globalDbg != nil {
			globalDbg.UnloadDriver()
		}
		fmt.Println("=== Driver Unloaded ===")
	})

	runUI(onReady)
}

func RunDriverOnly() {
	fmt.Println("=== Driver Only Mode ===")

	logFilePath := "hyperdbg_driver_only.log"

	fmt.Printf("日志文件: %s\n", logFilePath)

	globalDbg = debugger.NewHyperDbgWithOptions(true)

	app.ExitCallback(func() {
		fmt.Println("\n=== Unloading Driver ===")
		if globalDbg != nil {
			globalDbg.UnloadDriver()
		}
		fmt.Println("=== Driver Unloaded ===")
	})

	fmt.Println("=== Driver Loaded ===")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

	if globalDbg != nil {
		globalDbg.UnloadDriver()
	}

	fmt.Println("=== Driver Unloaded ===")
}

func runUI(onReady func(*debugger.HyperDbg)) {
	p := panel.New()
	hPanel := panel.NewHPanel()
	p.AddChild(hPanel)

	logPage = pages.NewLog()
	logFilePath := "hyperdbg_ui.log"

	fmt.Printf("日志文件: %s\n", logFilePath)

	globalDbg = debugger.NewHyperDbgWithOptions(true)
	globalDbg.OnTitleUpdate = func(title string) {
		app.SetWindowTitle(title)
	}

	NewToolbar(hPanel, globalDbg)

	if onReady != nil {
		onReady(globalDbg)
	}

	app.FileDropCallback(func(files []string) {
		if len(files) > 0 {
			exePath := files[0]
			go func() {
				fmt.Printf("Loading process: %s\n", exePath)
				if err := globalDbg.LoadProcess(exePath); err != nil {
					fmt.Printf("Failed to load process: %v\n", err)
				} else {
					globalDbg.UpdateAllPages()
				}
			}()
		}
	})

	m := safemap.NewOrdered[TabPageType, layout.Widget](func(yield func(TabPageType, layout.Widget) bool) {
		for _, tabType := range CpuType.EnumTypes() {
			switch tabType {
			case CpuType:
				if !yield(tabType, pages.NewCpu(globalDbg).Layout) {
					return
				}
			case PeViewType:
				if !yield(tabType, globalDbg.PeView.Layout()) {
					return
				}
			case LogType:
				if !yield(tabType, logPage.Layout()) {
					return
				}
			case NotesType:
				if !yield(tabType, globalDbg.Notes.Layout()) {
					return
				}
			case BreaksType:
				if !yield(tabType, globalDbg.Breakpoints.Layout()) {
					return
				}
			case MemoryType:
				if !yield(tabType, globalDbg.Memory.Layout()) {
					return
				}
			case SehType:
				if !yield(tabType, globalDbg.Seh.Layout()) {
					return
				}
			case ScriptType:
				if !yield(tabType, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case SymbolType:
				if !yield(tabType, globalDbg.Symbols.Layout()) {
					return
				}
			case SourceType:
				if !yield(tabType, globalDbg.Source.Layout()) {
					return
				}
			case ReferencesType:
				if !yield(tabType, globalDbg.References.Layout()) {
					return
				}
			case ThreadType:
				if !yield(tabType, globalDbg.Threads.Layout()) {
					return
				}
			case HandleType:
				if !yield(tabType, globalDbg.Handles.Layout()) {
					return
				}
			case TraceType:
				if !yield(tabType, globalDbg.Trace.Layout()) {
					return
				}
			case ArkType:
				if !yield(tabType, globalDbg.Ark.Layout()) {
					return
				}
			case ScyllaType:
				if !yield(tabType, globalDbg.Scylla.Layout()) {
					return
				}
			case LabelsType:
				if !yield(tabType, globalDbg.Labels.Layout()) {
					return
				}
			case CommentsType:
				if !yield(tabType, globalDbg.Comments.Layout()) {
					return
				}
			case FunctionsType:
				if !yield(tabType, globalDbg.Functions.Layout()) {
					return
				}
			case XrefsType:
				if !yield(tabType, globalDbg.Xrefs.Layout()) {
					return
				}
			case TypesType:
				if !yield(tabType, globalDbg.Types.Layout()) {
					return
				}
			case WatchesType:
				if !yield(tabType, globalDbg.Watches.Layout()) {
					return
				}
			case GraphsType:
				if !yield(tabType, globalDbg.Graphs.Layout()) {
					return
				}
			case ExceptionsType:
				if !yield(tabType, globalDbg.Exceptions.Layout()) {
					return
				}
			case BookmarksType:
				if !yield(tabType, globalDbg.Bookmarks.Layout()) {
					return
				}
			case LoopsType:
				if !yield(tabType, globalDbg.Loops.Layout()) {
					return
				}
			case EventsType:
				if !yield(tabType, pages.NewEvents(globalDbg).Layout()) {
					return
				}
			}
		}
	})

	vtab := tab.New(layout.Horizontal)
	for k, v := range m.Range() {
		tabItem := tab.NewTabItem(k.String(), v)
		vtab.AddTab(tabItem)
	}
	p.AddChild(vtab)

	app.Run("HyperDbg", func(gtx layout.Context) {
		select {
		case <-globalDbg.GetEventChan():
			globalDbg.UpdateAllPages()
		default:
		}
		p.Layout(gtx)
		myAppBar.LayoutOverlays(gtx)
	})
}

func NewToolbar(hpanel *panel.Panel, dbg *debugger.HyperDbg) {
	m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	for tipBtn := range toolbarButtons(m, dbg) {
		action := appbar.Action{
			Widget:    tipBtn.Layout,
			Update:    func(gtx layout.Context) {},
			AlignLeft: true,
		}
		myAppBar.AddAction(action)
	}
	myAppBar.AddAction(appbar.SettingsAction())
	myAppBar.AddAction(appbar.ThemeToggleAction())
	myAppBar.AddAction(appbar.AboutAction("HyperDbg v1.0\n\nA hypervisor-based debugger for Windows."))
	hpanel.AddChild(myAppBar)
}

func toolbarButtons(m *safemap.M[string, []byte], dbg *debugger.HyperDbg) iter.Seq[*TipIconButton] {
	return func(yield func(*TipIconButton) bool) {
		yield(NewTooltipButton(m.GetMust("open.png"), "open", func() {
			fmt.Println("Open button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("restart.png"), "restart", func() {
			fmt.Println("Restart button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("close.png"), "close", func() {
			fmt.Println("Close button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("run.png"), "run", func() {
			fmt.Println("Run button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("runthread.png"), "runthread", func() {
			fmt.Println("Run thread button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("pause.png"), "pause", func() {
			fmt.Println("Pause button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("stepin.png"), "stepin", func() {
			fmt.Println("Step in button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("stepover.png"), "stepover", func() {
			fmt.Println("Step over button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("trin.png"), "trin", func() {
			fmt.Println("Trace in button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("trover.png"), "trover", func() {
			fmt.Println("Trace over button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("tillret.png"), "tillret", func() {
			fmt.Println("Till return button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("tilluser.png"), "tilluser", func() {
			fmt.Println("Till user button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("log.png"), "log", func() {
			fmt.Println("Log button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("modules.png"), "modules", func() {
			fmt.Println("Modules button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("windows.png"), "windows", func() {
			fmt.Println("Windows button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("threads.png"), "threads", func() {
			fmt.Println("Threads button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("cpu.png"), "cpu", func() {
			fmt.Println("CPU button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("search.png"), "search", func() {
			fmt.Println("Search button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("trace.png"), "trace", func() {
			fmt.Println("Trace button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("bpoints.png"), "bpoints", func() {
			fmt.Println("Breakpoints button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("bpmem.png"), "bpmem", func() {
			fmt.Println("Memory breakpoints button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("bphard.png"), "bphard", func() {
			fmt.Println("Hardware breakpoints button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("options.png"), "options", func() {
			fmt.Println("Options button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("scylla.png"), "scylla", func() {
			fmt.Println("Scylla button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("about.png"), "about", func() {
			fmt.Println("About button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("settings.png"), "settings", func() {
			fmt.Println("Settings button clicked")
		}))
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
