//Code generated from mapPather - DO NOT EDIT.

package tabs

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/ddkwork/golibrary/skiaLib/widget/tabbar"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		Actions
	}
	Actions interface {
		LogButton() *unison.Button
		NotesButton() *unison.Button
		ProcessorCpuButton() *unison.Button
		SehChainButton() *unison.Button
		VtButton() *unison.Button
		Vt1Button() *unison.Button
		DownloadSymbolsButton() *unison.Button
		SourceButton() *unison.Button
		TraceButton() *unison.Button
		TraceintoButton() *unison.Button
		BreakpointButton() *unison.Button
		HandlesButton() *unison.Button
		MemoryMapButton() *unison.Button
		ThreadSwitchButton() *unison.Button
		ScriptCodeButton() *unison.Button
		StackButton() *unison.Button
	}
	object struct {
		vtButton              *unison.Button
		vt1Button             *unison.Button
		logButton             *unison.Button
		notesButton           *unison.Button
		processorCpuButton    *unison.Button
		sehChainButton        *unison.Button
		downloadSymbolsButton *unison.Button
		sourceButton          *unison.Button
		traceButton           *unison.Button
		traceintoButton       *unison.Button
		breakpointButton      *unison.Button
		handlesButton         *unison.Button
		memoryMapButton       *unison.Button
		threadSwitchButton    *unison.Button
		scriptCodeButton      *unison.Button
		stackButton           *unison.Button
	}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	o.scriptCodeButton = widget.CreateImageButton(scriptCodeKind.Image(), scriptCodeKind.String(), panel)
	o.stackButton = widget.CreateImageButton(stackKind.Image(), stackKind.String(), panel)
	o.sehChainButton = widget.CreateImageButton(sehChainKind.Image(), sehChainKind.String(), panel)
	o.vtButton = widget.CreateImageButton(vtKind.Image(), vtKind.String(), panel)
	o.vt1Button = widget.CreateImageButton(vt1Kind.Image(), vt1Kind.String(), panel)
	o.logButton = widget.CreateImageButton(logKind.Image(), logKind.String(), panel)
	o.notesButton = widget.CreateImageButton(notesKind.Image(), notesKind.String(), panel)
	o.processorCpuButton = widget.CreateImageButton(processorCpuKind.Image(), processorCpuKind.String(), panel)
	o.traceintoButton = widget.CreateImageButton(traceintoKind.Image(), traceintoKind.String(), panel)
	o.downloadSymbolsButton = widget.CreateImageButton(downloadSymbolsKind.Image(), downloadSymbolsKind.String(), panel)
	o.sourceButton = widget.CreateImageButton(sourceKind.Image(), sourceKind.String(), panel)
	o.traceButton = widget.CreateImageButton(traceKind.Image(), traceKind.String(), panel)
	o.threadSwitchButton = widget.CreateImageButton(threadSwitchKind.Image(), threadSwitchKind.String(), panel)
	o.breakpointButton = widget.CreateImageButton(breakpointKind.Image(), breakpointKind.String(), panel)
	o.handlesButton = widget.CreateImageButton(handlesKind.Image(), handlesKind.String(), panel)
	o.memoryMapButton = widget.CreateImageButton(memoryMapKind.Image(), memoryMapKind.String(), panel)
	return CreateBodyView(window).AsPanel()
	return panel
}

func CreateBodyView(window *unison.Window) *unison.Dock {
	var dock = unison.NewDock()
	cpu := tabbar.New(processorCpuKind.Name(), "main debugger view", unison.Green) //todo set TitleIcon
	cpupanel := cpu.AsPanel()
	cpupanel.SetLayout(&unison.FlexLayout{
		Columns:      1,
		HSpacing:     0,
		VSpacing:     0,
		HAlign:       0,
		VAlign:       0,
		EqualColumns: false,
	})
	cpupanel.AddChild(pageCpu.New().CanvasObject(window))
	dock.DockTo(cpu, nil, unison.BottomSide)

	log := tabbar.New(logKind.Name(), "main debug log", unison.Green)
	logpanel := log.AsPanel()
	logpanel.SetLayout(&unison.FlexLayout{
		Columns:      1,
		HSpacing:     0,
		VSpacing:     0,
		HAlign:       0,
		VAlign:       0,
		EqualColumns: false,
	})
	logpanel.AddChild(pageLog.New().CanvasObject(window))
	//todo another

	unison.Ancestor[*unison.DockContainer](cpu).Stack(log, -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(notesKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(breakpointKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(memoryMapKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(stackKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(sehChainKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(scriptCodeKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(downloadSymbolsKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(sourceKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(vt1Kind.Name(), "", unison.Green), -1) //todo
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(threadSwitchKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(handlesKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(traceKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).Stack(tabbar.New(vtKind.Name(), "", unison.Green), -1)
	unison.Ancestor[*unison.DockContainer](cpu).SetCurrentDockable(cpu)

	dock.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  200,
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	return dock
}

func (o *object) BreakpointButton() *unison.Button      { return o.breakpointButton }
func (o *object) HandlesButton() *unison.Button         { return o.handlesButton }
func (o *object) MemoryMapButton() *unison.Button       { return o.memoryMapButton }
func (o *object) ThreadSwitchButton() *unison.Button    { return o.threadSwitchButton }
func (o *object) ScriptCodeButton() *unison.Button      { return o.scriptCodeButton }
func (o *object) StackButton() *unison.Button           { return o.stackButton }
func (o *object) Vt1Button() *unison.Button             { return o.vt1Button }
func (o *object) LogButton() *unison.Button             { return o.logButton }
func (o *object) NotesButton() *unison.Button           { return o.notesButton }
func (o *object) ProcessorCpuButton() *unison.Button    { return o.processorCpuButton }
func (o *object) SehChainButton() *unison.Button        { return o.sehChainButton }
func (o *object) VtButton() *unison.Button              { return o.vtButton }
func (o *object) DownloadSymbolsButton() *unison.Button { return o.downloadSymbolsButton }
func (o *object) SourceButton() *unison.Button          { return o.sourceButton }
func (o *object) TraceButton() *unison.Button           { return o.traceButton }
func (o *object) TraceintoButton() *unison.Button       { return o.traceintoButton }
