package ui

import (
	"embed"
	"path/filepath"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/ui/ark"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ddkwork/ux"
)

// //go:embed Fleet.ico
// var icon []byte

//go:embed asserts/ico/ico_aaamain.png
var mainIcons []byte

//go:embed asserts/bar/*.png
var bar embed.FS

//go:embed asserts/pageico/*.png
var pageIco embed.FS

func Run() {
	panel := ux.NewPanel()
	hPanel := ux.NewHPanel()
	panel.AddChild(hPanel)

	NewToolbar(hPanel)

	m := new(safemap.M[TabPageType, ux.Widget])
	for _, Type := range CpuType.EnumTypes() {
		switch Type {
		case CpuType:
			m.Set(Type, NewCpu())
		case PeViewType:
			m.Set(Type, NewPeView())
		case LogType:
			m.Set(Type, ux.NewLogView())
		case NotesType:
			m.Set(Type, new(zeroWidget))
		case BreaksType:
			m.Set(Type, NewBreak())
		case MemoryType:
			m.Set(Type, NewMemory())
		case StackType:
			m.Set(Type, NewStack())
		case SehType:
			m.Set(Type, NewSeh())
		case ScriptType:
			m.Set(Type, NewScript())
		case SymbolType:
			m.Set(Type, NewSymbol())
		case SourceType:
			m.Set(Type, NewSource())
		case ReferencesType:
			m.Set(Type, new(zeroWidget))
		case ThreadType:
			m.Set(Type, NewThread())
		case HandleType:
			m.Set(Type, NewHandle())
		case TraceType:
			m.Set(Type, NewTrace())
		case ArkType:
			m.Set(Type, ark.New())
		}
	}

	vtab := ux.NewTabView(layout.Horizontal)
	for k, v := range m.Range() {
		tab := ux.NewTabItem(k.String(), v)
		// mylog.Todo("set tab ico")
		// // tabFileMap := stream.ReadEmbedFileMap(bar, "asserts/pageico")
		// for _, tab := range p.Elems() {
		//	if tab == p.cpu {
		//		continue
		//	}
		//	LeftContainer.Stack(tab, -1)
		// }
		vtab.AddTab(tab)
	}
	panel.AddChild(vtab)
	ux.Run("hyperdbg", panel)

	// callback := purego.NewCallback(func(text *byte) int {
	//	fmt.Println("Received data:", sdk.BytePointerToString(text)) // todo check api name
	//	return 0
	// })
	// sdk.SetTextMessageCallback(unsafe.Pointer(callback))
	// app.RunWithIco("HyperDbg "+winver.WindowVersion(), mainIcons, func(w *unison.Window) {
	//	//pages := NewTabPage()
	//	//w.Content().AddChild(pages.Layout())
	// })
}

type zeroWidget struct{}

func (z *zeroWidget) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{}
}

var (
	th     = ux.NewTheme()
	appBar *ux.AppBar
)

func NewToolbar(hpanel *ux.Panel) {
	m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	appBar = ux.InitAppBar(hpanel, func(yield func(*ux.TipIconButton) bool) {
		yield(ux.NewTooltipButton(m.GetMust("open.png"), "open", nil))
		yield(ux.NewTooltipButton(m.GetMust("restart.png"), "restart", nil))
		yield(ux.NewTooltipButton(m.GetMust("close.png"), "close", nil))
		yield(ux.NewTooltipButton(m.GetMust("run.png"), "run", nil))
		yield(ux.NewTooltipButton(m.GetMust("runthread.png"), "runthread", nil))
		yield(ux.NewTooltipButton(m.GetMust("pause.png"), "pause", nil))
		yield(ux.NewTooltipButton(m.GetMust("stepin.png"), "stepin", nil))
		yield(ux.NewTooltipButton(m.GetMust("stepover.png"), "stepover", nil))
		yield(ux.NewTooltipButton(m.GetMust("trin.png"), "trin", nil))
		yield(ux.NewTooltipButton(m.GetMust("trover.png"), "trover", nil))
		yield(ux.NewTooltipButton(m.GetMust("tillret.png"), "tillret", nil))
		yield(ux.NewTooltipButton(m.GetMust("tilluser.png"), "tilluser", nil))
		yield(ux.NewTooltipButton(m.GetMust("log.png"), "log", nil))
		yield(ux.NewTooltipButton(m.GetMust("modules.png"), "modules", nil))
		yield(ux.NewTooltipButton(m.GetMust("windows.png"), "windows", nil))
		yield(ux.NewTooltipButton(m.GetMust("threads.png"), "threads", nil))
		yield(ux.NewTooltipButton(m.GetMust("cpu.png"), "cpu", nil))
		yield(ux.NewTooltipButton(m.GetMust("search.png"), "search", nil))
		yield(ux.NewTooltipButton(m.GetMust("trace.png"), "trace", nil))
		yield(ux.NewTooltipButton(m.GetMust("bpoints.png"), "bpoints", nil))
		yield(ux.NewTooltipButton(m.GetMust("bpmem.png"), "bpmem", nil))
		yield(ux.NewTooltipButton(m.GetMust("bphard.png"), "bphard", nil))
		yield(ux.NewTooltipButton(m.GetMust("options.png"), "options", nil))
		yield(ux.NewTooltipButton(m.GetMust("scylla.png"), "scylla", nil))
		yield(ux.NewTooltipButton(m.GetMust("about.png"), "about", nil))
		yield(ux.NewTooltipButton(m.GetMust("settings.png"), "settings", nil))
	}, "hyperdbg is a debugger for windows")

	// m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	// return &toolbar{
	//	open: ux.NewImageButton("open", m.GetMust("open.png"), func() {}),
	//	restart: ux.NewImageButton("restart", m.GetMust("restart.png"), func() {
	//		mylog.Warning("RestartProcess", sdk.RestartProcess())
	//	}),
	//	close: ux.NewImageButton("close", m.GetMust("close.png"), func() { // exit command ?
	//		mylog.Warning("KillProcess", sdk.KillProcess())
	//	}),
	//	run: ux.NewImageButton("run", m.GetMust("run.png"), func() {
	//		targetExePathPtr := windows.StringToUTF16Ptr(TargetExePath)
	//		targetExePathInt32Ptr := (*int32)(unsafe.Pointer(targetExePathPtr))
	//		mylog.Warning("StartProcess", sdk.StartProcess(targetExePathInt32Ptr))
	//	}),
	//	runthread: ux.NewImageButton("runthread", m.GetMust("runthread.png"), func() {}),
	//	pause: ux.NewImageButton("pause", m.GetMust("pause.png"), func() {
	//		mylog.Warning("PauseKernelEvents", sdk.PauseKernelEvents())
	//	}),
	//	stepin: ux.NewImageButton("stepin", m.GetMust("stepin.png"), func() {
	//		// todo set F7 shortcut
	//		mylog.Warning("StepInto", sdk.StepIn())
	//	}),
	//	stepover: ux.NewImageButton("stepover", m.GetMust("stepover.png"), func() {
	//		mylog.Warning("StepOut", sdk.StepOut()) // todo set args
	//	}),
	//	trin:     ux.NewImageButton("trin", m.GetMust("trin.png"), func() {}),
	//	trover:   ux.NewImageButton("trover", m.GetMust("trover.png"), func() {}),
	//	tillret:  ux.NewImageButton("tillret", m.GetMust("tillret.png"), func() {}),
	//	tilluser: ux.NewImageButton("tilluser", m.GetMust("tilluser.png"), func() {}),
	//	log:      ux.NewImageButton("log", m.GetMust("log.png"), func() {}),
	//	modules:  ux.NewImageButton("modules", m.GetMust("modules.png"), func() {}),
	//	windows:  ux.NewImageButton("windows", m.GetMust("windows.png"), func() {}),
	//	threads:  ux.NewImageButton("threads", m.GetMust("threads.png"), func() {}),
	//	cpu: ux.NewImageButton("cpu", m.GetMust("cpu.png"), func() {
	//		mylog.Todo("goto cpu tab page") // dock.SetCurrentDockable(p.cpu)
	//	}),
	//	search: ux.NewImageButton("search", m.GetMust("search.png"), func() {}),
	//	trace:  ux.NewImageButton("trace", m.GetMust("trace.png"), func() {}),
	//	bpoints: ux.NewImageButton("bpoints", m.GetMust("bpoints.png"), func() {
	//		mylog.Warning("breakpoint list", sdk.BreakpointList())
	//	}),
	//	bpmem:   ux.NewImageButton("bpmem", m.GetMust("bpmem.png"), func() {}),
	//	bphard:  ux.NewImageButton("bphard", m.GetMust("bphard.png"), func() {}),
	//	options: ux.NewImageButton("options", m.GetMust("options.png"), func() {}),
	//	scylla:  ux.NewImageButton("scylla", m.GetMust("scylla.png"), func() {}),
	//	about:   ux.NewImageButton("about", m.GetMust("about.png"), func() {}),
	//	settings: ux.NewImageButton("settings", m.GetMust("settings.png"), func() {
	//		app.RunWithIco("settings", m.GetMust("settings.png"), func(w *unison.Window) {
	//			content := w.Content()
	//			content.SetLayout(&unison.FlexLayout{Columns: 1})
	//			content.AddChild(ux.NewVSpacer())
	//
	//			newPanel := ux.NewKeyValuePanel()
	//			content.AddChild(newPanel)
	//
	//			// .debug remote namedpipe \\.\pipe\HyperDbgDebug
	//			// .debug prepare serial 115200 com1
	//			type remoteData struct {
	//				Host string
	//				Port string
	//			}
	//			remoteEditor, _ := ux.NewStructView(
	//				remoteData{Host: "127.0.0.1", Port: "8080"},
	//				func(data remoteData) (values []ux.CellData) {
	//					return []ux.CellData{{Text: data.Host}, {Text: data.Port}}
	//				},
	//			)
	//			content.AddChild(remoteEditor)
	//
	//			newPanel.AddChild(ux.NewButton("Install Driver", func() {
	//				mylog.Check(sdk.VmxSupportDetection())
	//				mylog.Check(sdk.SetCustomDriverPathEx(sdk.TargetFilePath))
	//
	//				mylog.Trace("Install Driver", sdk.InstallVmmDriver())
	//			}))
	//			newPanel.AddChild(ux.NewButton("Uninstall Driver", func() {
	//				mylog.Trace("UninstallVmmDriver", sdk.UninstallVmmDriver())
	//			}))
	//
	//			newPanel.AddChild(ux.NewButton("Load Vmm", func() {
	//				mylog.Trace("LoadVmm", sdk.LoadVmm())
	//			}))
	//			newPanel.AddChild(ux.NewButton("UnLoad Vmm", func() {
	//				mylog.Trace("UnloadVmm", sdk.UnloadVmm())
	//			}))
	//			newPanel.AddChild(ux.NewButton("Stop Vmm", func() { //?
	//				mylog.Trace("StopVmmDriver", sdk.StopVmmDriver())
	//			}))
	//			newPanel.AddChild(ux.NewButton("xxxxxxxoo", func() {}))
	//
	//			newPanel.AddChild(ux.NewButton("ConnectLocalDebugger", func() {
	//				mylog.Trace("ConnectLocalDebugger")
	//				sdk.ConnectLocalDebugger()
	//			}))
	//			port := ux.NewButton("ConnectRemoteDebugger", func() {
	//				mylog.Trace("ConnectRemoteDebugger", sdk.ConnectRemoteDebuggerEx(remoteEditor.MetaData.Host, remoteEditor.MetaData.Port))
	//			})
	//			newPanel.AddChild(port)
	//
	//			newPanel.AddChild(ux.NewButton("Register context menu", func() { registerContextMenu(true) }))
	//			newPanel.AddChild(ux.NewButton("UnRegister context menu", func() { registerContextMenu(false) }))
	//		})
	//	}),
	// }
}

func registerContextMenu(enable bool) {
	remove := ""
	if !enable {
		remove = "-"
	}
	path := stream.RunDir() // need process location?
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

	// todo bug
	// powershell -Command "Start-Process cmd -ArgumentList '/c C:\Scripts\MyScript.bat' -Verb RunAs"
	// os.TempDir()//todo give it a dir
	stream.WriteTruncate("open.reg", g.Buffer)
	// reg import open.reg
	// reg export HKEY_CLASSES_ROOT HKCR_backup.reg
	stream.RunCommand("reg import open.reg")
}
