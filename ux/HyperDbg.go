package ux

import (
	"embed"

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
		Layout(w.Content())
	})
}

func Layout(parent unison.Paneler) unison.Paneler {
	t := newToolbar()
	widget.NewToolBar(parent, t.Elems()...) // make toolbar

	///make tabs
	widget.NewTabs()

	left := widget.NewTab("cpu", "", false, nil)
	right := widget.NewTab("log", "", false, nil)
	hSplit := widget.NewHSplit(left, right, 0.3)
	parent.AsPanel().AddChild(hSplit.Dock)

	mylog.Todo("set tab ico")
	// tabFileMap := stream.ReadEmbedFileMap(bar, "asserts/pageico")
	// hSplit.AddLeftItem(widget.NewTab("cpu", "", false, LayoutCpu(hSplit)))
	// hSplit.AddRightItem(widget.NewTab("log", "", false, LayoutLog(hSplit)))
	hSplit.AddRightItem(widget.NewTab("notes", "", false, LayoutNotes(hSplit)))
	hSplit.AddRightItem(widget.NewTab("break", "", false, LayoutBreak(hSplit)))
	hSplit.AddRightItem(widget.NewTab("memory", "", false, LayoutMemory(hSplit)))
	hSplit.AddRightItem(widget.NewTab("stack", "", false, LayoutStack(hSplit)))
	hSplit.AddRightItem(widget.NewTab("seh", "", false, LayoutSeh(hSplit)))
	hSplit.AddRightItem(widget.NewTab("script", "", false, LayoutScript(hSplit)))
	hSplit.AddRightItem(widget.NewTab("symbol", "", false, LayoutSymbol(hSplit)))
	hSplit.AddRightItem(widget.NewTab("source", "", false, LayoutSource(hSplit)))
	hSplit.AddRightItem(widget.NewTab("references", "", false, LayoutReferences(hSplit)))
	hSplit.AddRightItem(widget.NewTab("thread", "", false, LayoutThread(hSplit)))
	hSplit.AddRightItem(widget.NewTab("handle", "", false, LayoutHandle(hSplit)))
	hSplit.AddRightItem(widget.NewTab("trace", "", false, LayoutTrace(hSplit)))

	parent.AsPanel().AddChild(hSplit)
	return nil
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
