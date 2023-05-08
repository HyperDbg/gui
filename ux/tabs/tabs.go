//Code generated from mapPather - DO NOT EDIT.

package tabs

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/ddkwork/golibrary/skiaLib/widget/tabbar"
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
	return CreateBodyView().AsPanel()
	return panel
}

func CreateBodyView() *unison.Dock {
	var dock = unison.NewDock()
	yellow := tabbar.New(processorCpuKind.Name(), "main debugger view", unison.Yellow) //todo set TitleIcon
	dock.DockTo(yellow, nil, unison.BottomSide)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(logKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(notesKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(breakpointKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(memoryMapKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(stackKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(sehChainKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(scriptCodeKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(downloadSymbolsKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(sourceKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(vt1Kind.Name(), "", unison.Yellow), -1) //todo
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(threadSwitchKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(handlesKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(traceKind.Name(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(vtKind.Name(), "", unison.Yellow), -1)
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
