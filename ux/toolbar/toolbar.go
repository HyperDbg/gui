//Code generated from mapPather - DO NOT EDIT.

package toolbar

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		Actions
	}
	Actions interface {
		HandlesButton() *unison.Button
		MemoryButton() *unison.Button
		RunButton() *unison.Button
		SourceButton() *unison.Button
		LogButton() *unison.Button
		TraceButton() *unison.Button
		StackButton() *unison.Button
		TillretButton() *unison.Button
		TroverButton() *unison.Button
		PauseButton() *unison.Button
		RefersButton() *unison.Button
		PatchesButton() *unison.Button
		HelpButton() *unison.Button
		StepoverButton() *unison.Button
		TrinButton() *unison.Button
		BreaksButton() *unison.Button
		OpenButton() *unison.Button
		RestartButton() *unison.Button
		ThreadsButton() *unison.Button
		CloseButton() *unison.Button
		OptionsButton() *unison.Button
		StepinButton() *unison.Button
		CpuButton() *unison.Button
		GotoButton() *unison.Button
		ModulesButton() *unison.Button
		WindowsButton() *unison.Button
		AppearButton() *unison.Button
	}
	object struct {
		closeButton    *unison.Button
		openButton     *unison.Button
		restartButton  *unison.Button
		threadsButton  *unison.Button
		cpuButton      *unison.Button
		optionsButton  *unison.Button
		stepinButton   *unison.Button
		appearButton   *unison.Button
		gotoButton     *unison.Button
		modulesButton  *unison.Button
		windowsButton  *unison.Button
		sourceButton   *unison.Button
		handlesButton  *unison.Button
		memoryButton   *unison.Button
		runButton      *unison.Button
		traceButton    *unison.Button
		logButton      *unison.Button
		pauseButton    *unison.Button
		stackButton    *unison.Button
		tillretButton  *unison.Button
		troverButton   *unison.Button
		patchesButton  *unison.Button
		refersButton   *unison.Button
		stepoverButton *unison.Button
		trinButton     *unison.Button
		breaksButton   *unison.Button
		helpButton     *unison.Button
	}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	//todo sort
	o.patchesButton = widget.CreateImageButton(patchesKind.Image(), patchesKind.String(), panel)
	o.refersButton = widget.CreateImageButton(refersKind.Image(), refersKind.String(), panel)
	o.stepoverButton = widget.CreateImageButton(stepoverKind.Image(), stepoverKind.String(), panel)
	o.trinButton = widget.CreateImageButton(trinKind.Image(), trinKind.String(), panel)
	o.breaksButton = widget.CreateImageButton(breaksKind.Image(), breaksKind.String(), panel)
	o.helpButton = widget.CreateImageButton(helpKind.Image(), helpKind.String(), panel)
	o.threadsButton = widget.CreateImageButton(threadsKind.Image(), threadsKind.String(), panel)
	o.closeButton = widget.CreateImageButton(closeKind.Image(), closeKind.String(), panel)
	o.openButton = widget.CreateImageButton(openKind.Image(), openKind.String(), panel)
	o.restartButton = widget.CreateImageButton(restartKind.Image(), restartKind.String(), panel)
	o.cpuButton = widget.CreateImageButton(cpuKind.Image(), cpuKind.String(), panel)
	o.optionsButton = widget.CreateImageButton(optionsKind.Image(), optionsKind.String(), panel)
	o.stepinButton = widget.CreateImageButton(stepinKind.Image(), stepinKind.String(), panel)
	o.appearButton = widget.CreateImageButton(appearKind.Image(), appearKind.String(), panel)
	o.gotoButton = widget.CreateImageButton(gotoKind.Image(), gotoKind.String(), panel)
	o.modulesButton = widget.CreateImageButton(modulesKind.Image(), modulesKind.String(), panel)
	o.windowsButton = widget.CreateImageButton(windowsKind.Image(), windowsKind.String(), panel)
	o.runButton = widget.CreateImageButton(runKind.Image(), runKind.String(), panel)
	o.sourceButton = widget.CreateImageButton(sourceKind.Image(), sourceKind.String(), panel)
	o.handlesButton = widget.CreateImageButton(handlesKind.Image(), handlesKind.String(), panel)
	o.memoryButton = widget.CreateImageButton(memoryKind.Image(), memoryKind.String(), panel)
	o.traceButton = widget.CreateImageButton(traceKind.Image(), traceKind.String(), panel)
	o.logButton = widget.CreateImageButton(logKind.Image(), logKind.String(), panel)
	o.pauseButton = widget.CreateImageButton(pauseKind.Image(), pauseKind.String(), panel)
	o.stackButton = widget.CreateImageButton(stackKind.Image(), stackKind.String(), panel)
	o.tillretButton = widget.CreateImageButton(tillretKind.Image(), tillretKind.String(), panel)
	o.troverButton = widget.CreateImageButton(troverKind.Image(), troverKind.String(), panel)
	return panel
}

func (o *object) AppearButton() *unison.Button   { return o.appearButton }
func (o *object) GotoButton() *unison.Button     { return o.gotoButton }
func (o *object) ModulesButton() *unison.Button  { return o.modulesButton }
func (o *object) WindowsButton() *unison.Button  { return o.windowsButton }
func (o *object) SourceButton() *unison.Button   { return o.sourceButton }
func (o *object) HandlesButton() *unison.Button  { return o.handlesButton }
func (o *object) MemoryButton() *unison.Button   { return o.memoryButton }
func (o *object) RunButton() *unison.Button      { return o.runButton }
func (o *object) TraceButton() *unison.Button    { return o.traceButton }
func (o *object) LogButton() *unison.Button      { return o.logButton }
func (o *object) PauseButton() *unison.Button    { return o.pauseButton }
func (o *object) StackButton() *unison.Button    { return o.stackButton }
func (o *object) TillretButton() *unison.Button  { return o.tillretButton }
func (o *object) TroverButton() *unison.Button   { return o.troverButton }
func (o *object) PatchesButton() *unison.Button  { return o.patchesButton }
func (o *object) RefersButton() *unison.Button   { return o.refersButton }
func (o *object) StepoverButton() *unison.Button { return o.stepoverButton }
func (o *object) TrinButton() *unison.Button     { return o.trinButton }
func (o *object) BreaksButton() *unison.Button   { return o.breaksButton }
func (o *object) HelpButton() *unison.Button     { return o.helpButton }
func (o *object) CloseButton() *unison.Button    { return o.closeButton }
func (o *object) OpenButton() *unison.Button     { return o.openButton }
func (o *object) RestartButton() *unison.Button  { return o.restartButton }
func (o *object) ThreadsButton() *unison.Button  { return o.threadsButton }
func (o *object) CpuButton() *unison.Button      { return o.cpuButton }
func (o *object) OptionsButton() *unison.Button  { return o.optionsButton }
func (o *object) StepinButton() *unison.Button   { return o.stepinButton }
