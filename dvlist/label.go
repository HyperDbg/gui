package dvlist

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"sync"
	"sync/atomic"
)

const measureChar = "H"

// input val is a single line of rune, return number of chars n
// so that val[txtpos:txtpos+n] fits in width wid;
// unitWid is the avg single char widwth;
func lineChars(val []rune, txtpos int, wid, unitWid float32, fontsize float32, style fyne.TextStyle) int {

	amount := int(wid / unitWid)
	lineLen := len(val)
	// log.Printf("linechar for linelen %d wid %d unitwid %d", lineLen, wid, unitWid)
	if txtpos >= lineLen || txtpos < 0 {
		return 0
	}
	lastdirection := 0
	turns := 0
	for {
		if txtpos+amount > lineLen {
			amount = lineLen - txtpos
		}
		if txtpos >= lineLen || txtpos < 0 {
			return 0
		}
		tsize := fyne.MeasureText(string(val[txtpos:txtpos+amount]), fontsize, style)
		// log.Printf("tsize width %d, wid %d, unitWid %d, amount %d, linelen %d",
		// tsize.Width, wid, unitWid, amount, lineLen)
		if tsize.Width > wid {
			amount--
			if lastdirection == 1 {
				turns++
			}
			lastdirection = 0

		} else {
			if wid-tsize.Width < unitWid {
				return amount
			}

			if lastdirection == 0 {
				turns++
				if turns > 3 {
					return amount
				}
			}
			amount++
			lastdirection = 1
			if txtpos+amount > lineLen {
				return lineLen - txtpos
			}
		}

	}
}

type labelAct int

const (
	labelActReload labelAct = iota
)

const labelChanDepth = 4

type resizableLabel struct {
	widget.BaseWidget
	text             []rune
	mux              *sync.RWMutex
	partialDisplay   *uint32 // 1 if only part of text is showed, 0 otherwise
	look             *Look
	actChan          chan labelAct
	metaData         interface{}
	dragging         *uint32
	draggingHandler  LabelDraggingHandler
	tapHandler       LabelTapHandler
	doubleTapHandler LabelTapHandler
}

type LabelDraggingHandler func(d interface{}, evt *fyne.DragEvent)
type LabelTapHandler func(d interface{}, evt *fyne.PointEvent)

type ResizableLabelOption func(label *resizableLabel)

func WithMetaData(d interface{}) ResizableLabelOption {
	return func(label *resizableLabel) {
		label.metaData = d
	}
}

func WithDraggingHandler(h LabelDraggingHandler) ResizableLabelOption {
	return func(label *resizableLabel) {
		label.draggingHandler = h
	}
}

func WithTapHandler(h LabelTapHandler) ResizableLabelOption {
	return func(label *resizableLabel) {
		label.tapHandler = h
	}
}
func WithDoubleTapHandler(h LabelTapHandler) ResizableLabelOption {
	return func(label *resizableLabel) {
		label.doubleTapHandler = h
	}
}

func NewResizableLabel(s string, look *Look, options ...ResizableLabelOption) *resizableLabel {
	rlabel := &resizableLabel{
		text:           []rune(s),
		mux:            new(sync.RWMutex),
		partialDisplay: new(uint32),
		look:           look,
		actChan:        make(chan labelAct, labelChanDepth),
		dragging:       new(uint32),
	}
	atomic.StoreUint32(rlabel.partialDisplay, 0)
	atomic.StoreUint32(rlabel.dragging, 0)
	for _, o := range options {
		o(rlabel)
	}
	rlabel.ExtendBaseWidget(rlabel)
	return rlabel
}
func (rl *resizableLabel) CreateRenderer() fyne.WidgetRenderer {
	rl.ExtendBaseWidget(rl)
	return newResizableLabelRender(rl)
}

func (rl *resizableLabel) Get() []rune {
	rl.mux.RLock()
	defer rl.mux.RUnlock()
	return rl.text
}

func (rl *resizableLabel) Set(s string) []rune {
	rl.mux.Lock()
	rl.text = []rune(s)
	rl.mux.Unlock()
	select {
	case rl.actChan <- labelActReload:
		rl.Refresh()
	default:
		break
	}
	return rl.text
}

func (rl *resizableLabel) Dragged(evt *fyne.DragEvent) {
	if rl.draggingHandler == nil {
		return
	}
	atomic.StoreUint32(rl.dragging, 1)
	// log.Printf("field %d dragged X %d", rl.metaData, evt.DraggedX)
	rl.draggingHandler(rl.metaData, evt)

}
func (rl *resizableLabel) DragEnd() {
	// atomic.StoreUint32(row.dragging, 0)

}

func (rl *resizableLabel) Tapped(evt *fyne.PointEvent) {
	// log.Printf("mouse clicked")
	if atomic.LoadUint32(rl.dragging) != 0 {
		atomic.StoreUint32(rl.dragging, 0)
		return
	}
	if rl.tapHandler != nil {
		rl.tapHandler(rl.metaData, evt)
	}
}

func (rl *resizableLabel) DoubleTapped(evt *fyne.PointEvent) {
	if atomic.LoadUint32(rl.dragging) != 0 {
		atomic.StoreUint32(rl.dragging, 0)
		return
	}
	if rl.doubleTapHandler != nil {
		rl.doubleTapHandler(rl.metaData, evt)
	}

}

func getUnitSize(fontsize float32, style fyne.TextStyle) fyne.Size {
	return fyne.MeasureText(measureChar, fontsize, style)
}

// IsPartial return true if only part of text is showed
func (rl *resizableLabel) IsPartial() bool {
	return atomic.LoadUint32(rl.partialDisplay) == 1
}

type resizableLabelRender struct {
	label            *resizableLabel
	canvasText       *canvas.Text
	overallContainer *fyne.Container
	unitSize         fyne.Size
	mux              *sync.RWMutex
}

func newResizableLabelRender(l *resizableLabel) *resizableLabelRender {
	r := &resizableLabelRender{
		label:            l,
		overallContainer: container.NewWithoutLayout(),
		canvasText:       canvas.NewText(string(l.Get()), theme.ForegroundColor()),
		mux:              new(sync.RWMutex),
	}
	r.overallContainer.Add(r.canvasText)
	r.calUnitSize()
	return r
}
func (rlr *resizableLabelRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (rlr *resizableLabelRender) calUnitSize() {
	rlr.unitSize = getUnitSize(rlr.label.look.TextSize, rlr.label.look.TextStyle)
	return
}

func (rlr *resizableLabelRender) MinSize() fyne.Size {
	return rlr.unitSize
}
func (rlr *resizableLabelRender) Destroy() {
	// log.Printf("label destroyed")

}
func (rlr *resizableLabelRender) Objects() []fyne.CanvasObject {
	// log.Printf("label objects")
	rlr.mux.RLock()
	defer rlr.mux.RUnlock()
	return []fyne.CanvasObject{rlr.overallContainer}
}

func (rlr *resizableLabelRender) Refresh() {
	select {
	case act := <-rlr.label.actChan:
		switch act {
		case labelActReload:
			rlr.Layout(rlr.label.Size())
		}
	default:
		return
	}
	canvas.Refresh(rlr.label)
}
func (rlr *resizableLabelRender) Layout(layoutsize fyne.Size) {
	// log.Printf("label layout for size %v", layoutsize)
	if layoutsize.Width <= 0 {
		return
	}
	rlr.mux.Lock()
	defer rlr.mux.Unlock()
	rlr.calUnitSize()
	totalSize := fyne.MeasureText(string(rlr.label.Get()), rlr.label.look.TextSize, rlr.label.look.TextStyle)
	i := lineChars(rlr.label.Get(), 0, layoutsize.Width, rlr.unitSize.Width, rlr.label.look.TextSize, rlr.label.look.TextStyle)
	if i < len(rlr.label.Get())-1 {
		atomic.StoreUint32(rlr.label.partialDisplay, 1)
	} else {
		atomic.StoreUint32(rlr.label.partialDisplay, 0)
	}
	rlr.canvasText = canvas.NewText(string(rlr.label.Get()[:i]), rlr.label.look.TextColor)
	rlr.canvasText.TextStyle = rlr.label.look.TextStyle
	rlr.canvasText.TextSize = rlr.label.look.TextSize
	rlr.canvasText.Refresh()
	switch rlr.label.look.Alignment {
	case fyne.TextAlignCenter:
		if layoutsize.Width > totalSize.Width {
			rlr.canvasText.Move(fyne.NewPos((layoutsize.Width-totalSize.Width)/2, 0))
		}
	case fyne.TextAlignTrailing:
		if layoutsize.Width > totalSize.Width {
			rlr.canvasText.Move(fyne.NewPos(layoutsize.Width-totalSize.Width, 0))
		}
	default:
	}

	rlr.overallContainer = container.NewWithoutLayout()
	rlr.overallContainer.Add(rlr.canvasText)
}
