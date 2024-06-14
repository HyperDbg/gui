package ux

import (
	"embed"
	"path/filepath"

	"github.com/richardwilkes/unison/enums/align"
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
	app.RunWithIco("HyperDbg", mainIcons, func(w *unison.Window) {
		pages := NewPage()
		w.Content().FileDropCallback = func(files []string) {
			switch filepath.Ext(files[0]) {
			case ".exe", ".dll", ".sys":
				mylog.Trace("dropped file", files[0])
				pages.pe.SetContent(layoutPeView(pages.dock)) //todo test is dock or cpu tab page
			default:
				mylog.Check("not support file type")
			}
		}
		pages.Layout(w.Content())
	})
}

type (
	Page struct {
		dock   *unison.Dock
		cpu    *widget.Tab
		pe     *widget.Tab
		log    *widget.Tab
		notes  *widget.Tab
		breakp *widget.Tab
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
	}
)

func (p *Page) Elems() []*widget.Tab {
	return []*widget.Tab{
		p.cpu,
		p.pe,
		p.log,
		p.notes,
		p.breakp,
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
	}
}

func (p *Page) Layout(parent unison.Paneler) unison.Paneler {
	t := newToolbar()
	widget.NewToolBar(parent, t.Elems()...) // make toolbar

	///make tabs
	cpuTab := widget.NewTab("cpu", "", false, LayoutCpu(p.dock))
	p.dock.DockTo(cpuTab, nil, side.Left)
	parent.AsPanel().AddChild(p.dock)
	LeftContainer := widget.NewDockContainer(cpuTab)

	mylog.Todo("set tab ico")
	// tabFileMap := stream.ReadEmbedFileMap(bar, "asserts/pageico")
	for _, tab := range p.Elems() {
		LeftContainer.Stack(tab, -1)
	}
	LeftContainer.SetCurrentDockable(cpuTab)
	return nil
}

func NewPage() *Page {
	dock := unison.NewDock()
	dock.AsPanel().SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	p := &Page{
		dock:   dock,
		cpu:    widget.NewTab("cpu", "", false, LayoutCpu(dock)),
		pe:     widget.NewTab("peView", "", false, layoutPeView(dock)),
		log:    widget.NewTab("log", "", false, LayoutLog(dock)),
		notes:  widget.NewTab("notes", "", false, LayoutNotes(dock)),
		breakp: widget.NewTab("break", "", false, LayoutBreak(dock)),
		memory: widget.NewTab("memory", "", false, LayoutMemory(dock)),
		stack:  widget.NewTab("stack", "", false, LayoutStack(dock)),
		seh:    widget.NewTab("seh", "", false, LayoutSeh(dock)),
		script: widget.NewTab("script", "", false, LayoutScript(dock)),
		symbol: widget.NewTab("symbol", "", false, LayoutSymbol(dock)),
		source: widget.NewTab("source", "", false, LayoutSource(dock)),
		ref:    widget.NewTab("references", "", false, LayoutReferences(dock)),
		thread: widget.NewTab("thread", "", false, LayoutThread(dock)),
		handle: widget.NewTab("handle", "", false, LayoutHandle(dock)),
		trace:  widget.NewTab("trace", "", false, LayoutTrace(dock)),
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
	}
}

func newToolbar() *toolbar {
	m := stream.ReadEmbedFileMap(bar, "asserts/bar")
	return &toolbar{
		open:      widget.NewImageButton(m.Get("open.png"), func() {}),
		restart:   widget.NewImageButton(m.Get("restart.png"), func() {}),
		close:     widget.NewImageButton(m.Get("close.png"), func() {}),
		run:       widget.NewImageButton(m.Get("run.png"), func() {}),
		runthread: widget.NewImageButton(m.Get("runthread.png"), func() {}),
		pause:     widget.NewImageButton(m.Get("pause.png"), func() {}),
		stepin:    widget.NewImageButton(m.Get("stepin.png"), func() {}),
		stepover:  widget.NewImageButton(m.Get("stepover.png"), func() {}),
		trin:      widget.NewImageButton(m.Get("trin.png"), func() {}),
		trover:    widget.NewImageButton(m.Get("trover.png"), func() {}),
		tillret:   widget.NewImageButton(m.Get("tillret.png"), func() {}),
		tilluser:  widget.NewImageButton(m.Get("tilluser.png"), func() {}),
		log:       widget.NewImageButton(m.Get("log.png"), func() {}),
		modules:   widget.NewImageButton(m.Get("modules.png"), func() {}),
		windows:   widget.NewImageButton(m.Get("windows.png"), func() {}),
		threads:   widget.NewImageButton(m.Get("threads.png"), func() {}),
		cpu:       widget.NewImageButton(m.Get("cpu.png"), func() {}),
		search:    widget.NewImageButton(m.Get("search.png"), func() {}),
		trace:     widget.NewImageButton(m.Get("trace.png"), func() {}),
		bpoints:   widget.NewImageButton(m.Get("bpoints.png"), func() {}),
		bpmem:     widget.NewImageButton(m.Get("bpmem.png"), func() {}),
		bphard:    widget.NewImageButton(m.Get("bphard.png"), func() {}),
		options:   widget.NewImageButton(m.Get("options.png"), func() {}),
		scylla:    widget.NewImageButton(m.Get("scylla.png"), func() {}),
		about:     widget.NewImageButton(m.Get("about.png"), func() {}),
	}
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
	}
)
