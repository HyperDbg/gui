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
	"github.com/ddkwork/ux/widget/logview"
	"github.com/ddkwork/ux/widget/panel"
	"github.com/ddkwork/ux/widget/tab"
	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/ark"
)

//go:embed asserts/ico/ico_aaamain.png
var mainIcons []byte

//go:embed asserts/bar/*.png
var bar embed.FS

//go:embed asserts/pageico/*.png
var pageIco embed.FS

func Run(onReady func(*debugger.Debugger)) {
	p := panel.New()
	hPanel := panel.NewHPanel()
	p.AddChild(hPanel)

	dbg := debugger.New()
	dbg.OnTitleUpdate = func(title string) {
		app.SetWindowTitle(title)
	}

	NewToolbar(hPanel, dbg)

	if onReady != nil {
		onReady(dbg)
	}

	app.FileDropCallback(func(files []string) {
		if len(files) > 0 {
			exePath := files[0]
			go func() {
				dbg.CreateProcess(exePath, "")
			}()
		}
	})

	m := safemap.NewOrdered[TabPageType, layout.Widget](func(yield func(TabPageType, layout.Widget) bool) {
		for _, Type := range CpuType.EnumTypes() {
			switch Type {
			case CpuType:
				if !yield(Type, NewCpu(dbg).Layout) {
					return
				}
			case PeViewType:
				if !yield(Type, dbg.GetPeView().Layout()) {
					return
				}
			case LogType:
				if !yield(Type, logview.New().Layout) {
					return
				}
			case NotesType:
				if !yield(Type, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case BreaksType:
				if !yield(Type, dbg.GetBreakpoints().Layout()) {
					return
				}
			case MemoryType:
				if !yield(Type, dbg.GetMemory().Layout()) {
					return
				}
			case SehType:
				if !yield(Type, dbg.GetSEH().Layout()) {
					return
				}
			case ScriptType:
				if !yield(Type, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case SymbolType:
				if !yield(Type, dbg.GetSymbols().Layout()) {
					return
				}
			case SourceType:
				if !yield(Type, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case ReferencesType:
				if !yield(Type, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case ThreadType:
				if !yield(Type, dbg.GetThreads().Layout()) {
					return
				}
			case HandleType:
				if !yield(Type, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }) {
					return
				}
			case TraceType:
				if !yield(Type, dbg.GetTrace().Layout()) {
					return
				}
			case ArkType:
				arkMgr := ark.New()
				if !yield(Type, arkMgr.Layout()) {
					return
				}
			case ScyllaType:
				if !yield(Type, dbg.GetScylla().Layout()) {
					return
				}
			case LabelsType:
				if !yield(Type, dbg.GetLabels().Layout()) {
					return
				}
			case CommentsType:
				if !yield(Type, dbg.GetComments().Layout()) {
					return
				}
			case FunctionsType:
				if !yield(Type, dbg.GetFunctions().Layout()) {
					return
				}
			case XrefsType:
				if !yield(Type, dbg.GetXrefs().Layout()) {
					return
				}
			case TypesType:
				if !yield(Type, dbg.GetTypes().Layout()) {
					return
				}
			case WatchesType:
				if !yield(Type, dbg.GetWatches().Layout()) {
					return
				}
			case GraphsType:
				if !yield(Type, dbg.GetGraphs().Layout()) {
					return
				}
			case ExceptionsType:
				if !yield(Type, dbg.GetExceptions().Layout()) {
					return
				}
			case BookmarksType:
				if !yield(Type, dbg.GetBookmarks().Layout()) {
					return
				}
			case LoopsType:
				if !yield(Type, dbg.GetLoops().Layout()) {
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
	app.Run("hyperdbg", func(gtx layout.Context) {
		select {
		case <-dbg.GetEventChan():
			dbg.UpdateAllPages()
		default:
		}
		p.Layout(gtx)
		myAppBar.LayoutOverlays(gtx)
	})
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

func init() {
	myAppBar.Title("").HeightLevel(appbar.HeightSmall)
}

func NewToolbar(hpanel *panel.Panel, dbg *debugger.Debugger) {
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
	myAppBar.AddAction(appbar.AboutAction("HyperDbg v1.0\n\nA debugger for Windows."))
	hpanel.AddChild(myAppBar)
}

func toolbarButtons(m *safemap.M[string, []byte], dbg *debugger.Debugger) iter.Seq[*TipIconButton] {
	return func(yield func(*TipIconButton) bool) {
		yield(NewTooltipButton(m.GetMust("open.png"), "open", func() {
		}))
		yield(NewTooltipButton(m.GetMust("restart.png"), "restart", func() {
			exePath := dbg.GetExePath()
			if exePath != "" {
				go func() {
					if dbg.GetProcessHandle() != 0 {
						dbg.TerminateProcess(0)
					}
					dbg.CreateProcess(exePath, "")
				}()
			}
		}))
		yield(NewTooltipButton(m.GetMust("close.png"), "close", func() {
			if dbg.GetProcessHandle() != 0 {
				go dbg.Detach()
			}
		}))
		yield(NewTooltipButton(m.GetMust("run.png"), "run", func() {
			dbg.Continue()
		}))
		yield(NewTooltipButton(m.GetMust("runthread.png"), "runthread", func() {
		}))
		yield(NewTooltipButton(m.GetMust("pause.png"), "pause", func() {
		}))
		yield(NewTooltipButton(m.GetMust("stepin.png"), "stepin", func() {
			dbg.StepInto()
		}))
		yield(NewTooltipButton(m.GetMust("stepover.png"), "stepover", func() {
			dbg.StepOver()
		}))
		yield(NewTooltipButton(m.GetMust("trin.png"), "trin", func() {
		}))
		yield(NewTooltipButton(m.GetMust("trover.png"), "trover", func() {
		}))
		yield(NewTooltipButton(m.GetMust("tillret.png"), "tillret", func() {
		}))
		yield(NewTooltipButton(m.GetMust("tilluser.png"), "tilluser", func() {
		}))
		yield(NewTooltipButton(m.GetMust("log.png"), "log", func() {
		}))
		yield(NewTooltipButton(m.GetMust("modules.png"), "modules", func() {
		}))
		yield(NewTooltipButton(m.GetMust("windows.png"), "windows", func() {
		}))
		yield(NewTooltipButton(m.GetMust("threads.png"), "threads", func() {
		}))
		yield(NewTooltipButton(m.GetMust("cpu.png"), "cpu", func() {
		}))
		yield(NewTooltipButton(m.GetMust("search.png"), "search", func() {
		}))
		yield(NewTooltipButton(m.GetMust("trace.png"), "trace", func() {
		}))
		yield(NewTooltipButton(m.GetMust("bpoints.png"), "bpoints", func() {
		}))
		yield(NewTooltipButton(m.GetMust("bpmem.png"), "bpmem", func() {
		}))
		yield(NewTooltipButton(m.GetMust("bphard.png"), "bphard", func() {
		}))
		yield(NewTooltipButton(m.GetMust("options.png"), "options", func() {
		}))
		yield(NewTooltipButton(m.GetMust("scylla.png"), "scylla", func() {
		}))
		yield(NewTooltipButton(m.GetMust("about.png"), "about", func() {
			fmt.Println("About button clicked")
		}))
		yield(NewTooltipButton(m.GetMust("settings.png"), "settings", func() {
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
