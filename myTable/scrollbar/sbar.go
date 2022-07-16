package scrollbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"sync"
	"sync/atomic"
)

// OffsetChangeHandler is the type handler function get called whenever scroller offset changes
type OffsetChangeHandler func(newoffset uint32)

// OffsetResolution indicate max value of the scroll offset
const OffsetResolution = 100

// ScrollBar is the fyne widget for scroll bar
type ScrollBar struct {
	widget.BaseWidget
	offset     *uint32
	actChan    chan posAction
	offsetH    OffsetChangeHandler
	horizontal bool
}

const actChanDepth = 4

// NewSBar returns a new scrollbar instance,
// with the specified handler function handler, could be nil;
// horizontal scrollbar if h is true, vertical otherwise;
func NewSBar(handler OffsetChangeHandler, h bool) *ScrollBar {
	r := &ScrollBar{
		offset:     new(uint32),
		actChan:    make(chan posAction, actChanDepth),
		offsetH:    handler,
		horizontal: h,
	}
	atomic.StoreUint32(r.offset, 0)
	r.ExtendBaseWidget(r)
	return r
}

// CreateRenderer implments fyne.Widget interface
func (sbar *ScrollBar) CreateRenderer() fyne.WidgetRenderer {
	sbar.ExtendBaseWidget(sbar)
	return newSbarRender(sbar)
}

// DragEnd implements fyne.Draggable interface
func (sbar *ScrollBar) DragEnd() {
}

// Dragged implements fyne.Draggable interface
func (sbar *ScrollBar) Dragged(evt *fyne.DragEvent) {
	cord := evt.Position.Y + evt.Dragged.DY
	if sbar.horizontal {
		cord = evt.Position.X + evt.Dragged.DX
	}
	select {
	case sbar.actChan <- posAction{act: actionDrag, val: cord}:
	default:
		return
	}
	sbar.Refresh()
}

// IsHorizontal returns true if it is horizontal scroll bar
func (sbar *ScrollBar) IsHorizontal() bool {
	return sbar.horizontal
}

// SetOffset set the scroll offset;
// note this function won't trigger the OffsetChangeHandler function
func (sbar *ScrollBar) SetOffset(o uint32) {
	sbar.setOffset(o, false)
}
func (sbar *ScrollBar) setOffset(o uint32, callh bool) {
	var i = o
	if o > OffsetResolution {
		i = OffsetResolution
	}
	atomic.StoreUint32(sbar.offset, i)
	sbar.actChan <- posAction{
		act: actionAdmin,
	}
	sbar.Refresh()
	if callh {
		if sbar.offsetH != nil {
			sbar.offsetH(i)
		}
	}
}

// GetOffset returns current offset
func (sbar *ScrollBar) GetOffset() uint32 { return atomic.LoadUint32(sbar.offset) }

const (
	BarWidth  = 20 //竖滚动条宽度
	BarHeight = 50 //竖滚动条高度
)

const (
	actionDrag = iota
	actionAdmin
)

type posAction struct {
	act int
	val float32
}

type sbarRender struct {
	sbar             *ScrollBar
	bar              *canvas.Rectangle
	curBarSize       fyne.Size
	mux              *sync.RWMutex
	overallContainer *fyne.Container
}

func newSbarRender(b *ScrollBar) *sbarRender {
	r := &sbarRender{
		sbar: b,
		bar:  canvas.NewRectangle(theme.ScrollBarColor()),
		mux:  new(sync.RWMutex),
	}
	r.curBarSize = fyne.NewSize(BarWidth, BarHeight)
	if b.horizontal {
		r.curBarSize = fyne.NewSize(BarHeight, BarWidth)
	}
	r.overallContainer = fyne.NewContainerWithoutLayout(r.bar)
	return r
}

func (sbrr *sbarRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (sbrr *sbarRender) Destroy() {
}
func (sbrr *sbarRender) Layout(layoutsize fyne.Size) {
	sbrr.mux.RLock()
	sbrr.bar.Resize(sbrr.curBarSize)
	x := layoutsize.Width - sbrr.curBarSize.Width
	y := sbrr.offsetToCord(int(layoutsize.Height))
	if sbrr.sbar.horizontal {
		y = layoutsize.Height - sbrr.curBarSize.Height
		x = sbrr.offsetToCord(int(layoutsize.Width))
	}
	sbrr.mux.RUnlock()
	sbrr.bar.Move(fyne.NewPos(x, y))
	sbrr.bar.Refresh()
}

func (sbrr *sbarRender) MinSize() fyne.Size {
	sbrr.mux.RLock()
	defer sbrr.mux.RUnlock()
	if sbrr.sbar.horizontal {
		return fyne.NewSize(BarHeight, BarWidth)
	}
	return fyne.NewSize(BarWidth, BarHeight)
}

func (sbrr *sbarRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{sbrr.overallContainer}
}
func (sbrr *sbarRender) calCordToOffset(newCord float32) uint32 {
	// debug.PrintStack()
	var newoffset uint32
	if sbrr.sbar.horizontal {
		newoffset = uint32((newCord * OffsetResolution) / (sbrr.sbar.Size().Width - sbrr.curBarSize.Width))
	} else {
		newoffset = uint32((newCord * OffsetResolution) / (sbrr.sbar.Size().Height - sbrr.curBarSize.Height))
	}
	if newoffset < 0 {
		newoffset = 0
	}
	if newoffset > OffsetResolution {
		newoffset = OffsetResolution
	}
	return newoffset
}

func (sbrr *sbarRender) cordToOffset(newCord float32) {
	newoffset := sbrr.calCordToOffset(newCord)
	atomic.StoreUint32(sbrr.sbar.offset, newoffset)
	if sbrr.sbar.offsetH != nil {
		sbrr.sbar.offsetH(newoffset)
	}
}

func (sbrr *sbarRender) offsetToCord(overall int) float32 {
	r := ((overall - int(sbrr.curBarSize.Height)) * int(atomic.LoadUint32(sbrr.sbar.offset))) / OffsetResolution
	if sbrr.sbar.horizontal {
		r = ((overall - int(sbrr.curBarSize.Width)) * int(atomic.LoadUint32(sbrr.sbar.offset))) / OffsetResolution
	}
	return float32(r)
}

func (sbrr *sbarRender) Refresh() {
	select {
	case a := <-sbrr.sbar.actChan:
		switch a.act {
		case actionDrag:
			var newx, newy float32
			pos := sbrr.bar.Position()
			if !sbrr.sbar.horizontal {
				newx = pos.X
				newy = a.val
				if newy < 0 {
					newy = 0
				}
				if newy > sbrr.sbar.Size().Height-sbrr.curBarSize.Height {
					newy = sbrr.sbar.Size().Height - sbrr.curBarSize.Height
				}
			} else {
				newy = pos.Y
				newx = a.val
				if newx < 0 {
					newx = 0
				}
				if newx > sbrr.sbar.Size().Width-sbrr.curBarSize.Width {
					newx = sbrr.sbar.Size().Width - sbrr.curBarSize.Width
				}
			}
			sbrr.bar.Move(fyne.NewPos(newx, newy))
			sbrr.bar.Refresh()
			if !sbrr.sbar.horizontal {
				sbrr.cordToOffset(newy)
			} else {
				sbrr.cordToOffset(newx)
			}
		case actionAdmin:
			var newx, newy float32
			if !sbrr.sbar.horizontal {
				newx = sbrr.bar.Position().X
				newy = sbrr.offsetToCord(int(sbrr.sbar.Size().Height))
			} else {
				newy = sbrr.bar.Position().Y
				newx = sbrr.offsetToCord(int(sbrr.sbar.Size().Width))
			}
			sbrr.bar.Move(fyne.NewPos(newx, newy))
			sbrr.bar.Refresh()
		}
	default:
		break
	}
	canvas.Refresh(sbrr.sbar)
}
