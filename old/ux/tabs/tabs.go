package tabs

import (
	_ "embed"
	"github.com/ddkwork/GolandProjects/ui/widget"
	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/side"

	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
)

type (
	Object struct {
		VtButton              *unison.Button
		Vt1Button             *unison.Button
		LogButton             *unison.Button
		NotesButton           *unison.Button
		ProcessorCpuButton    *unison.Button
		SehChainButton        *unison.Button
		DownloadSymbolsButton *unison.Button
		SourceButton          *unison.Button
		TraceButton           *unison.Button
		TraceintoButton       *unison.Button
		BreakpointButton      *unison.Button
		HandlesButton         *unison.Button
		MemoryMapButton       *unison.Button
		ThreadSwitchButton    *unison.Button
		ScriptCodeButton      *unison.Button
		StackButton           *unison.Button
	}
)

func New() *Object {
	return &Object{
		VtButton:              VtKind.Button(),
		Vt1Button:             Vt1Kind.Button(),
		LogButton:             LogKind.Button(),
		NotesButton:           NotesKind.Button(),
		ProcessorCpuButton:    CpuKind.Button(),
		SehChainButton:        SehKind.Button(),
		DownloadSymbolsButton: SymbolsKind.Button(),
		SourceButton:          SourceKind.Button(),
		TraceButton:           TraceKind.Button(),
		TraceintoButton:       TraceintoKind.Button(),
		BreakpointButton:      BreakpointKind.Button(),
		HandlesButton:         HandlesKind.Button(),
		MemoryMapButton:       MemoryKind.Button(),
		ThreadSwitchButton:    ThreadKind.Button(),
		ScriptCodeButton:      ScriptKind.Button(),
		StackButton:           StackKind.Button(),
	}
}

func (o *Object) CanvasObject() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	panel.AddChild(o.ScriptCodeButton)
	panel.AddChild(o.StackButton)
	panel.AddChild(o.SehChainButton)
	panel.AddChild(o.VtButton)
	panel.AddChild(o.Vt1Button)
	panel.AddChild(o.LogButton)
	panel.AddChild(o.NotesButton)
	panel.AddChild(o.ProcessorCpuButton)
	panel.AddChild(o.TraceintoButton)
	panel.AddChild(o.DownloadSymbolsButton)
	panel.AddChild(o.SourceButton)
	panel.AddChild(o.TraceButton)
	panel.AddChild(o.ThreadSwitchButton)
	panel.AddChild(o.BreakpointButton)
	panel.AddChild(o.HandlesButton)
	panel.AddChild(o.MemoryMapButton)
	return CreateBodyView().AsPanel()
	return panel
}

func CreateBodyView() *unison.Dock {
	cpu := widget.NewAppTabs(CpuKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(CpuKind.Buffer()) })
	log := widget.NewAppTabs(LogKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(LogKind.Buffer()) })
	Notes := widget.NewAppTabs(NotesKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(NotesKind.Buffer()) })
	Breakpoint := widget.NewAppTabs(BreakpointKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(BreakpointKind.Buffer()) })
	MemoryMap := widget.NewAppTabs(MemoryKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(MemoryKind.Buffer()) })
	Stack := widget.NewAppTabs(StackKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(StackKind.Buffer()) })
	SehChain := widget.NewAppTabs(SehKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(SehKind.Buffer()) })
	ScriptCode := widget.NewAppTabs(ScriptKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(ScriptKind.Buffer()) })
	DownloadSymbols := widget.NewAppTabs(SymbolsKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(SymbolsKind.Buffer()) })
	Source := widget.NewAppTabs(SourceKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(SourceKind.Buffer()) })
	Vt1 := widget.NewAppTabs(Vt1Kind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(Vt1Kind.Buffer()) })
	ThreadSwitch := widget.NewAppTabs(ThreadKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(ThreadKind.Buffer()) })
	Handles := widget.NewAppTabs(HandlesKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(HandlesKind.Buffer()) })
	Traceinto := widget.NewAppTabs(TraceintoKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(TraceintoKind.Buffer()) })
	Vt := widget.NewAppTabs(VtKind.String(), "", nil, func() unison.Drawable { return widget.SizedDrawable(VtKind.Buffer()) })

	dock := unison.NewDock()
	dock.DockTo(cpu, nil, side.Left)

	unison.Ancestor[*unison.DockContainer](cpu).Stack(log, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Notes, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Breakpoint, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(MemoryMap, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Stack, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(SehChain, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(ScriptCode, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(DownloadSymbols, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Source, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Vt1, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(ThreadSwitch, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Handles, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Traceinto, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(Vt, -1)
	unison.Ancestor[*unison.DockContainer](cpu).SetCurrentDockable(cpu)

	cpu.AddChild(pageCpu.New().CanvasObject())
	log.AddChild(pageLog.New().CanvasObject())

	dock.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  200,
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	return dock
}
