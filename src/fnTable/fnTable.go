package fnTable

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
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
	return container.NewAppTabs(
		container.NewTabItemWithIcon("cpu", ico.cpu(), pageCpu.New().CanvasObject(window)),
		container.NewTabItemWithIcon("log", ico.log(), widget.NewButton("log", nil)),
		container.NewTabItemWithIcon("notes", ico.notes(), widget.NewButton("notes", nil)),
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
	}
	pageIcoObj struct {
	}
)

func (p *pageIcoObj) cpu() fyne.Resource     { return fyne.NewStaticResource("i23", cpu) }
func (p *pageIcoObj) log() fyne.Resource     { return fyne.NewStaticResource("i23", log) }
func (p *pageIcoObj) notes() fyne.Resource   { return fyne.NewStaticResource("i23", notes) }
func (p *pageIcoObj) breaks() fyne.Resource  { return fyne.NewStaticResource("i23", breaks) }
func (p *pageIcoObj) memory() fyne.Resource  { return fyne.NewStaticResource("i23", memory) }
func (p *pageIcoObj) stack() fyne.Resource   { return fyne.NewStaticResource("i23", stack) }
func (p *pageIcoObj) sehList() fyne.Resource { return fyne.NewStaticResource("i23", sehList) }
func (p *pageIcoObj) script() fyne.Resource  { return fyne.NewStaticResource("i23", script) }
func (p *pageIcoObj) symbols() fyne.Resource { return fyne.NewStaticResource("i23", symbols) }
func (p *pageIcoObj) source() fyne.Resource  { return fyne.NewStaticResource("i23", source) }
func (p *pageIcoObj) xFrom() fyne.Resource   { return fyne.NewStaticResource("i23", xFrom) }
func (p *pageIcoObj) thead() fyne.Resource   { return fyne.NewStaticResource("i23", thead) }
func (p *pageIcoObj) handle() fyne.Resource  { return fyne.NewStaticResource("i23", handle) }
func (p *pageIcoObj) trace() fyne.Resource   { return fyne.NewStaticResource("i23", trace) }
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
