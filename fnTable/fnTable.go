package fnTable

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/golibrary/src/driverTool"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/fnTable/pageBreaks"
	"github.com/ddkwork/hyperdbgui/fnTable/pageCpu"
	"github.com/ddkwork/hyperdbgui/fnTable/pageHandle"
	"github.com/ddkwork/hyperdbgui/fnTable/pageMemory"
	"github.com/ddkwork/hyperdbgui/fnTable/pageNotes"
	"github.com/ddkwork/hyperdbgui/fnTable/pageScript"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSehList"
	"github.com/ddkwork/hyperdbgui/fnTable/pageStack"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSymbol"
	"github.com/ddkwork/hyperdbgui/fnTable/pageThead"
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
	driver.SetLoadVmmTapped(func() {

	})
	driver.SetUnloadVmmTapped(func() {

	})
	return container.NewAppTabs(
		container.NewTabItemWithIcon(ico.cpu().Name(), ico.cpu(), pageCpu.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.log().Name(), ico.log(), widget.NewMultiLineEntry()),
		container.NewTabItemWithIcon(ico.notes().Name(), ico.notes(), pageNotes.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.breaks().Name(), ico.breaks(), pageBreaks.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.memory().Name(), ico.memory(), pageMemory.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.stack().Name(), ico.stack(), pageStack.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.sehList().Name(), ico.sehList(), pageSehList.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.script().Name(), ico.script(), pageScript.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.symbols().Name(), ico.symbols(), pageSymbol.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.source().Name(), ico.source(), widget.NewButton("source", nil)),
		container.NewTabItemWithIcon(ico.xFrom().Name(), ico.xFrom(), widget.NewButton("xFrom", nil)),
		container.NewTabItemWithIcon(ico.thead().Name(), ico.thead(), pageThead.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.handle().Name(), ico.handle(), pageHandle.New().CanvasObject(window)),
		container.NewTabItemWithIcon(ico.trace().Name(), ico.trace(), widget.NewButton("trace", nil)),
		container.NewTabItemWithIcon(ico.vt().Name(), ico.vt(), driver.CanvasObject(window)),
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

var (

	//go:embed pageIco/processor-cpu.png
	cpu []byte

	//go:embed pageIco/log.png
	log []byte

	//go:embed pageIco/notes.png
	notes []byte

	//go:embed pageIco/breakpoint.png
	breaks []byte

	//go:embed pageIco/memory-map.png
	memory []byte

	//go:embed pageIco/stack.png
	stack []byte

	//go:embed pageIco/seh-chain.png
	sehList []byte

	//go:embed pageIco/script-code.png
	script []byte

	//go:embed pageIco/download_symbols.png
	symbols []byte

	//go:embed pageIco/source.png
	source []byte

	// todo
	//
	//go:embed pageIco/log.png
	xFrom []byte

	//go:embed pageIco/thread-switch.png
	thead []byte

	//go:embed pageIco/handles.png
	handle []byte

	//go:embed pageIco/trace.png
	trace []byte

	//go:embed pageIco/vt.ico
	vt []byte
)
