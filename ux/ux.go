package ux

import (
	"embed"
	"path/filepath"
	"time"

	"github.com/ddkwork/HyperDbg/sdk"

	"github.com/richardwilkes/unison/enums/align"

	"github.com/ddkwork/app/ms/hook/winver"

	"github.com/richardwilkes/unison/enums/side"

	"github.com/ddkwork/app"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/richardwilkes/unison"
)

////go:embed Fleet.ico
//var icon []byte

//go:embed asserts/ico/ico_aaamain.png
var mainIcons []byte

//go:embed asserts/bar/*.png
var bar embed.FS

//go:embed asserts/pageico/*.png
var pageIco embed.FS

func Run() {
	app.RunWithIco("HyperDbg "+winver.WindowVersion(), mainIcons, func(w *unison.Window) {
		pages := NewTabPage()
		w.Content().AddChild(pages.Layout())
	})
}

type (
	TagPage struct {
		dock   *unison.Dock
		cpu    *widget.Tab
		pe     *widget.Tab
		log    *widget.Tab
		notes  *widget.Tab
		breaks *widget.Tab
		memory *widget.Tab
		stack  *widget.Tab
		seh    *widget.Tab
		script *widget.Tab
		symbol *widget.Tab
		source *widget.Tab
		ref    *widget.Tab
		thread *widget.Tab
		handle *widget.Tab
		trace  *widget.Tab
		ark    *widget.Tab
	}
)

func (p *TagPage) Elems() []*widget.Tab {
	return []*widget.Tab{
		p.cpu,
		p.pe,
		p.log,
		p.notes,
		p.breaks,
		p.memory,
		p.stack,
		p.seh,
		p.script,
		p.symbol,
		p.source,
		p.ref,
		p.thread,
		p.handle,
		p.trace,
		p.ark,
	}
}

func (p *TagPage) Layout() unison.Paneler {
	panel := widget.NewPanel()
	t := newToolbar()
	panel.AddChild(widget.NewToolBar(t.Elems()...))
	p.dock.DockTo(p.cpu, nil, side.Left)

	LeftContainer := widget.NewDockContainer(p.cpu)

	mylog.Todo("set tab ico")
	// tabFileMap := stream.ReadEmbedFileMap(bar, "asserts/pageico")
	for _, tab := range p.Elems() {
		if tab == p.cpu {
			continue
		}
		LeftContainer.Stack(tab, -1)
	}
	LeftContainer.SetCurrentDockable(p.cpu)
	panel.AddChild(p.dock)
	return panel
}

func NewTabPage() *TagPage {
	dock := unison.NewDock()
	dock.AsPanel().SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: align.Fill,
		VAlign: align.Fill,
	})
	log := LayoutLog()
	logBuffer := stream.NewBuffer(mylog.Body())
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			log.AsPanel().RemoveAllChildren()
			logBuffer.WriteStringLn(mylog.Body())
			log.AsPanel().AddChild(widget.NewLogView(logBuffer.String()))
		}
	}()

	p := &TagPage{
		dock:   dock,
		cpu:    widget.NewTab("cpu", "", LayoutCpu()),
		pe:     widget.NewTab("peView", "", LayoutPeView()),
		log:    widget.NewTab("log", "", log),
		notes:  widget.NewTab("notes", "", LayoutNotes()),
		breaks: widget.NewTab("break", "", LayoutBreak()),
		memory: widget.NewTab("memory", "", LayoutMemory()),
		stack:  widget.NewTab("stack", "", LayoutStack()),
		seh:    widget.NewTab("seh", "", LayoutSeh()),
		script: widget.NewTab("script", "", LayoutScript()),
		symbol: widget.NewTab("symbol", "", LayoutSymbol()),
		source: widget.NewTab("source", "", LayoutSource()),
		ref:    widget.NewTab("references", "", LayoutReferences()),
		thread: widget.NewTab("thread", "", LayoutThread()),
		handle: widget.NewTab("handle", "", LayoutHandle()),
		trace:  widget.NewTab("trace", "", LayoutTrace()),
		ark:    widget.NewTab("ark", "", LayoutArk()),
	}
	return p
}

func (t *toolbar) Elems() []*unison.Button {
	return []*unison.Button{
		t.open,
		t.restart,
		t.close,
		t.run,
		t.runthread,
		t.pause,
		t.stepin,
		t.stepover,
		t.trin,
		t.trover,
		t.tillret,
		t.tilluser,
		t.log,
		t.modules,
		t.windows,
		t.threads,
		t.cpu,
		t.search,
		t.trace,
		t.bpoints,
		t.bpmem,
		t.bphard,
		t.options,
		t.scylla,
		t.about,
		t.settings,
	}
}

func newToolbar() *toolbar {
	m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	return &toolbar{
		open: widget.NewImageButton("open", m.Get("open.png"), func() {}),
		restart: widget.NewImageButton("restart", m.Get("restart.png"), func() {
			mylog.Warning("RestartProcess", sdk.RestartProcess())
		}),
		close: widget.NewImageButton("close", m.Get("close.png"), func() { // exit command ?
			mylog.Warning("KillProcess", sdk.KillProcess())
		}),
		run: widget.NewImageButton("run", m.Get("run.png"), func() {
			mylog.Warning("StartProcess", sdk.StartProcess(TargetExePath))
		}),
		runthread: widget.NewImageButton("runthread", m.Get("runthread.png"), func() {}),
		pause: widget.NewImageButton("pause", m.Get("pause.png"), func() {
			mylog.Warning("PauseKernelEvents", sdk.PauseKernelEvents())
		}),
		stepin: widget.NewImageButton("stepin", m.Get("stepin.png"), func() {
			// todo set F7 shortcut
			mylog.Warning("StepInto", sdk.StepIn())
		}),
		stepover: widget.NewImageButton("stepover", m.Get("stepover.png"), func() {
			mylog.Warning("StepOut", sdk.StepOut()) // todo set args
		}),
		trin:     widget.NewImageButton("trin", m.Get("trin.png"), func() {}),
		trover:   widget.NewImageButton("trover", m.Get("trover.png"), func() {}),
		tillret:  widget.NewImageButton("tillret", m.Get("tillret.png"), func() {}),
		tilluser: widget.NewImageButton("tilluser", m.Get("tilluser.png"), func() {}),
		log:      widget.NewImageButton("log", m.Get("log.png"), func() {}),
		modules:  widget.NewImageButton("modules", m.Get("modules.png"), func() {}),
		windows:  widget.NewImageButton("windows", m.Get("windows.png"), func() {}),
		threads:  widget.NewImageButton("threads", m.Get("threads.png"), func() {}),
		cpu: widget.NewImageButton("cpu", m.Get("cpu.png"), func() {
			mylog.Todo("goto cpu tab page") // dock.SetCurrentDockable(p.cpu)
		}),
		search: widget.NewImageButton("search", m.Get("search.png"), func() {}),
		trace:  widget.NewImageButton("trace", m.Get("trace.png"), func() {}),
		bpoints: widget.NewImageButton("bpoints", m.Get("bpoints.png"), func() {
			mylog.Warning("breakpoint list", sdk.BreakpointList())
		}),
		bpmem:   widget.NewImageButton("bpmem", m.Get("bpmem.png"), func() {}),
		bphard:  widget.NewImageButton("bphard", m.Get("bphard.png"), func() {}),
		options: widget.NewImageButton("options", m.Get("options.png"), func() {}),
		scylla:  widget.NewImageButton("scylla", m.Get("scylla.png"), func() {}),
		about:   widget.NewImageButton("about", m.Get("about.png"), func() {}),
		settings: widget.NewImageButton("settings", m.Get("settings.png"), func() {
			app.RunWithIco("settings", m.Get("settings.png"), func(w *unison.Window) {
				content := w.Content()
				content.SetLayout(&unison.FlexLayout{Columns: 1})
				content.AddChild(widget.NewVSpacer())

				newPanel := widget.NewKeyValuePanel()
				content.AddChild(newPanel)

				// .debug remote namedpipe \\.\pipe\HyperDbgDebug
				// .debug prepare serial 115200 com1
				type remoteData struct {
					Host string
					Port string
				}
				remoteEditor, _ := widget.NewStructView(
					remoteData{Host: "127.0.0.1", Port: "8080"},
					func(data remoteData) (values []widget.CellData) {
						return []widget.CellData{{Text: data.Host}, {Text: data.Port}}
					},
				)
				content.AddChild(remoteEditor)

				newPanel.AddChild(widget.NewButton("Install Driver", func() {
					mylog.Check(sdk.VmxSupportDetection())
					mylog.Check(sdk.SetCustomDriverPathEx(sdk.SysPath))
					mylog.Trace("InstallVmmDriver", sdk.InstallVmmDriver())
				}))
				newPanel.AddChild(widget.NewButton("Uninstall Driver", func() {
					mylog.Trace("UninstallVmmDriver", sdk.UninstallVmmDriver())
				}))

				newPanel.AddChild(widget.NewButton("Load Vmm", func() {
					mylog.Trace("LoadVmm", sdk.LoadVmm())
				}))
				newPanel.AddChild(widget.NewButton("UnLoad Vmm", func() {
					mylog.Trace("UnloadVmm", sdk.UnloadVmm())
				}))
				newPanel.AddChild(widget.NewButton("Stop Vmm", func() { //?
					mylog.Trace("StopVmmDriver", sdk.StopVmmDriver())
				}))
				newPanel.AddChild(widget.NewButton("xxxxxxxoo", func() {}))

				newPanel.AddChild(widget.NewButton("ConnectLocalDebugger", func() {
					mylog.Trace("ConnectLocalDebugger")
					sdk.ConnectLocalDebugger()
				}))
				port := widget.NewButton("ConnectRemoteDebugger", func() {
					mylog.Trace("ConnectRemoteDebugger", sdk.ConnectRemoteDebuggerEx(remoteEditor.MetaData.Host, remoteEditor.MetaData.Port))
				})
				newPanel.AddChild(port)

				newPanel.AddChild(widget.NewButton("Register context menu", func() { registerContextMenu(true) }))
				newPanel.AddChild(widget.NewButton("UnRegister context menu", func() { registerContextMenu(false) }))
			})
		}),
	}
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

type (
	toolbar struct {
		open      *unison.Button
		restart   *unison.Button
		close     *unison.Button
		run       *unison.Button
		runthread *unison.Button
		pause     *unison.Button
		stepin    *unison.Button
		stepover  *unison.Button
		trin      *unison.Button
		trover    *unison.Button
		tillret   *unison.Button
		tilluser  *unison.Button
		log       *unison.Button
		modules   *unison.Button
		windows   *unison.Button
		threads   *unison.Button
		cpu       *unison.Button
		search    *unison.Button
		trace     *unison.Button
		bpoints   *unison.Button
		bpmem     *unison.Button
		bphard    *unison.Button
		options   *unison.Button
		scylla    *unison.Button
		about     *unison.Button
		settings  *unison.Button
	}
)
