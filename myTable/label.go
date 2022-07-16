package myTable

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
	style            *StyleObj
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

func NewResizableLabel(s string, look *StyleObj, options ...ResizableLabelOption) *resizableLabel {
	label := &resizableLabel{
		text:           []rune(s),
		mux:            new(sync.RWMutex),
		partialDisplay: new(uint32),
		style:          look,
		actChan:        make(chan labelAct, labelChanDepth),
		dragging:       new(uint32),
	}
	atomic.StoreUint32(label.partialDisplay, 0)
	atomic.StoreUint32(label.dragging, 0)
	for _, o := range options {
		o(label)
	}
	label.ExtendBaseWidget(label)
	return label
}
func (l *resizableLabel) CreateRenderer() fyne.WidgetRenderer {
	l.ExtendBaseWidget(l)
	return newResizableLabelRender(l)
}

func (l *resizableLabel) Get() []rune {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.text
}

func (l *resizableLabel) Set(s string) []rune {
	l.mux.Lock()
	l.text = []rune(s)
	l.mux.Unlock()
	select {
	case l.actChan <- labelActReload:
		l.Refresh()
	default:
		break
	}
	return l.text
}

func (l *resizableLabel) Dragged(evt *fyne.DragEvent) {
	if l.draggingHandler == nil {
		return
	}
	atomic.StoreUint32(l.dragging, 1)
	// log.Printf("field %d dragged X %d", l.metaData, evt.DraggedX)
	l.draggingHandler(l.metaData, evt)

}
func (l *resizableLabel) DragEnd() {
	// atomic.StoreUint32(row.dragging, 0)

}

func (l *resizableLabel) Tapped(evt *fyne.PointEvent) {
	// log.Printf("mouse clicked")
	if atomic.LoadUint32(l.dragging) != 0 {
		atomic.StoreUint32(l.dragging, 0)
		return
	}
	if l.tapHandler != nil {
		l.tapHandler(l.metaData, evt)
	}
}

func (l *resizableLabel) DoubleTapped(evt *fyne.PointEvent) {
	if atomic.LoadUint32(l.dragging) != 0 {
		atomic.StoreUint32(l.dragging, 0)
		return
	}
	if l.doubleTapHandler != nil {
		l.doubleTapHandler(l.metaData, evt)
	}

}

func getUnitSize(fontsize float32, style fyne.TextStyle) fyne.Size {
	return fyne.MeasureText(measureChar, fontsize, style)
}

// IsPartial return true if only part of text is showed
func (l *resizableLabel) IsPartial() bool {
	return atomic.LoadUint32(l.partialDisplay) == 1
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
func (r *resizableLabelRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *resizableLabelRender) calUnitSize() {
	r.unitSize = getUnitSize(r.label.style.TextSize, r.label.style.TextStyle)
	return
}

func (r *resizableLabelRender) MinSize() fyne.Size {
	return r.unitSize
}
func (r *resizableLabelRender) Destroy() {
	// log.Printf("label destroyed")

}
func (r *resizableLabelRender) Objects() []fyne.CanvasObject {
	// log.Printf("label objects")
	r.mux.RLock()
	defer r.mux.RUnlock()
	return []fyne.CanvasObject{r.overallContainer}
}

func (r *resizableLabelRender) Refresh() {
	select {
	case act := <-r.label.actChan:
		switch act {
		case labelActReload:
			r.Layout(r.label.Size())
		}
	default:
		return
	}
	canvas.Refresh(r.label)
}
func (r *resizableLabelRender) Layout(layoutsize fyne.Size) {
	// log.Printf("label layout for size %v", layoutsize)
	if layoutsize.Width <= 0 {
		return
	}
	r.mux.Lock()
	defer r.mux.Unlock()
	r.calUnitSize()
	totalSize := fyne.MeasureText(string(r.label.Get()), r.label.style.TextSize, r.label.style.TextStyle)
	i := lineChars(r.label.Get(), 0, layoutsize.Width, r.unitSize.Width, r.label.style.TextSize, r.label.style.TextStyle)
	if i < len(r.label.Get())-1 {
		atomic.StoreUint32(r.label.partialDisplay, 1)
	} else {
		atomic.StoreUint32(r.label.partialDisplay, 0)
	}
	r.canvasText = canvas.NewText(string(r.label.Get()[:i]), r.label.style.TextColor)
	r.canvasText.TextStyle = r.label.style.TextStyle
	r.canvasText.TextSize = r.label.style.TextSize
	r.canvasText.Refresh()
	switch r.label.style.Alignment {
	case fyne.TextAlignCenter:
		if layoutsize.Width > totalSize.Width {
			r.canvasText.Move(fyne.NewPos((layoutsize.Width-totalSize.Width)/2, 0))
		}
	case fyne.TextAlignTrailing:
		if layoutsize.Width > totalSize.Width {
			r.canvasText.Move(fyne.NewPos(layoutsize.Width-totalSize.Width, 0))
		}
	default:
	}

	r.overallContainer = container.NewWithoutLayout()
	r.overallContainer.Add(r.canvasText)
}
