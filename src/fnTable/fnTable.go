package fnTable

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageNotes"
	"github.com/ddkwork/librarygo/src/driverTool"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
)

type (
	Interface interface {
		canvasobjectapi.Interface
	}
	object struct{}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	ico := newPageIcoObj()
	driver := driverTool.New()
	driver.Driver().DeviceName = "HyperdbgHypervisorDevice"
	return container.NewAppTabs(
		container.NewTabItemWithIcon("cpu", ico.cpu(), pageCpu.New().CanvasObject(window)),
		container.NewTabItemWithIcon("log", ico.log(), widget.NewMultiLineEntry()), //todo export for set log
		container.NewTabItemWithIcon("notes", ico.notes(), pageNotes.New().CanvasObject(window)),
		container.NewTabItemWithIcon("breaks", ico.breaks(), widget.NewButton("breaks", nil)),
		container.NewTabItemWithIcon("memory", ico.memory(), widget.NewButton("memory", nil)),
		container.NewTabItemWithIcon("stack", ico.stack(), widget.NewButton("stack", nil)),
		container.NewTabItemWithIcon("sehList", ico.sehList(), widget.NewButton("sehList", nil)),
		container.NewTabItemWithIcon("script", ico.script(), widget.NewButton("script", nil)),
		container.NewTabItemWithIcon("symbols", ico.symbols(), widget.NewButton("symbols", nil)),
		container.NewTabItemWithIcon("source", ico.source(), widget.NewButton("source", nil)),
		container.NewTabItemWithIcon("xFrom", ico.xFrom(), widget.NewButton("xFrom", nil)),
		container.NewTabItemWithIcon("thead", ico.thead(), widget.NewButton("thead", nil)),
		container.NewTabItemWithIcon("handle", ico.handle(), widget.NewButton("handle", nil)),
		container.NewTabItemWithIcon("trace", ico.trace(), widget.NewButton("trace", nil)),
		container.NewTabItemWithIcon("driver control", ico.vt(), driver.CanvasObject(window)),
	)
}

type (
	pageIco interface {
		cpu() fyne.Resource
		log() fyne.Resource
		notes() fyne.Resource
		breaks() fyne.Resource
		memory() fyne.Resource
		stack() fyne.Resource
		sehList() fyne.Resource
		script() fyne.Resource
		symbols() fyne.Resource
		source() fyne.Resource
		xFrom() fyne.Resource
		thead() fyne.Resource
		handle() fyne.Resource
		trace() fyne.Resource
		vt() fyne.Resource
	}
	pageIcoObj struct {
	}
)

func (p *pageIcoObj) cpu() fyne.Resource     { return fyne.NewStaticResource("cpu", cpu) }
func (p *pageIcoObj) log() fyne.Resource     { return fyne.NewStaticResource("log", log) }
func (p *pageIcoObj) notes() fyne.Resource   { return fyne.NewStaticResource("notes", notes) }
func (p *pageIcoObj) breaks() fyne.Resource  { return fyne.NewStaticResource("breaks", breaks) }
func (p *pageIcoObj) memory() fyne.Resource  { return fyne.NewStaticResource("memory", memory) }
func (p *pageIcoObj) stack() fyne.Resource   { return fyne.NewStaticResource("stack", stack) }
func (p *pageIcoObj) sehList() fyne.Resource { return fyne.NewStaticResource("sehList", sehList) }
func (p *pageIcoObj) script() fyne.Resource  { return fyne.NewStaticResource("script", script) }
func (p *pageIcoObj) symbols() fyne.Resource { return fyne.NewStaticResource("symbols", symbols) }
func (p *pageIcoObj) source() fyne.Resource  { return fyne.NewStaticResource("source", source) }
func (p *pageIcoObj) xFrom() fyne.Resource   { return fyne.NewStaticResource("xFrom", xFrom) }
func (p *pageIcoObj) thead() fyne.Resource   { return fyne.NewStaticResource("thead", thead) }
func (p *pageIcoObj) handle() fyne.Resource  { return fyne.NewStaticResource("handle", handle) }
func (p *pageIcoObj) trace() fyne.Resource   { return fyne.NewStaticResource("trace", trace) }
func (p *pageIcoObj) vt() fyne.Resource      { return fyne.NewStaticResource("vt", vt) }
func newPageIcoObj() pageIco {
	return &pageIcoObj{}
}

//go:embed pageIco/processor-cpu.png
var cpu []byte

//go:embed pageIco/log.png
var log []byte

//go:embed pageIco/notes.png
var notes []byte

//go:embed pageIco/breakpoint.png
var breaks []byte

//go:embed pageIco/memory-map.png
var memory []byte

//go:embed pageIco/stack.png
var stack []byte

//go:embed pageIco/seh-chain.png
var sehList []byte

//go:embed pageIco/script-code.png
var script []byte

//go:embed pageIco/download_symbols.png
var symbols []byte

//go:embed pageIco/source.png
var source []byte

//todo
//go:embed pageIco/log.png
var xFrom []byte

//go:embed pageIco/thread-switch.png
var thead []byte

//go:embed pageIco/handles.png
var handle []byte

//go:embed pageIco/trace.png
var trace []byte

//go:embed pageIco/vt.ico
var vt []byte
