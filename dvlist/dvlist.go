// Package dvlist provides a fyne widget for displaying data sets in a table format
package dvlist

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/dvlist/scrollbar"
	"image/color"
	"sync"
	"sync/atomic"
)

// Data is an interface for the data sets to be displayed in DVList
type Data interface {
	// Len return total number of items
	Len() int
	// Fields return the field names
	Fields() []string
	// Item return item with index i, as a slice of strings, each string represents a field's value
	Item(id int) []string
	// sort items based on field, impact following Item() call
	Sort(field int, ascend bool)
	// Filter filters items with kw in field i, after calling Filter, calling item might return empty value
	// if i==-1, then filter all fields;
	Filter(kw string, i int)
}

type listActType int

const (
	actScrollBody listActType = iota
	actSortData
	actSetSelection
	actSetData
)

type listAct struct {
	act listActType
	arg interface{}
}

type listCord struct {
	Row, Column int
}

const actChanDepth = 16

// DVList is a fyne Widget that display data sets in a table format
type DVList struct {
	widget.BaseWidget
	data    Data
	mux     *sync.RWMutex
	actChan chan *listAct
	// curOffset            *uint32 //current offset value of scrollbar
	curStartPoint        *uint32 //the index of first row showed
	curShowedLines       *uint32
	multiSelection       bool
	curSelections        []bool
	headerLook, bodyLook *Look
	selectionHandler     OnSelectionHandler
	doubleClickHandler   DoubleClickHandler
}

// Option is a function could be specified for NewDVList() to provide customization
type Option func(list *DVList)

// WithMultiSelections allows multiple selections
func WithMultiSelections() Option {
	return func(list *DVList) {
		list.multiSelection = true
	}
}

// WithHeaderLook specifies the look for the table header
func WithHeaderLook(look *Look) Option {
	return func(list *DVList) {
		list.headerLook = look
	}
}

// WithBodyLook specifies the look for the table body
func WithBodyLook(look *Look) Option {
	return func(list *DVList) {
		list.bodyLook = look
	}
}

// WithSelectionHandler specifies h as a handler function gets called when a row is selected
func WithSelectionHandler(h OnSelectionHandler) Option {
	return func(list *DVList) {
		list.selectionHandler = h
	}
}

// WithDoubleClickHandler specifies h as a handler function gets called when a row is double clicked
func WithDoubleClickHandler(h DoubleClickHandler) Option {
	return func(list *DVList) {
		list.doubleClickHandler = h
	}
}

// NewDVList return a new DVList instance with the specified data
func NewDVList(data Data, options ...Option) (*DVList, error) {
	if len(data.Fields()) > MaxFields {
		return nil, fmt.Errorf("data has more fields (%d) than allowed (%d)", len(data.Fields()), MaxFields)
	}
	list := &DVList{
		data:           data,
		mux:            new(sync.RWMutex),
		actChan:        make(chan *listAct, actChanDepth),
		curStartPoint:  new(uint32),
		curShowedLines: new(uint32),
		curSelections:  []bool{},
		headerLook:     defaultHeaderLook(),
		bodyLook:       defaultBodyLook(),
	}
	atomic.StoreUint32(list.curStartPoint, 0)
	for i := 0; i < data.Len(); i++ {
		list.curSelections = append(list.curSelections, false)
	}
	for _, o := range options {
		o(list)
	}
	list.ExtendBaseWidget(list)
	return list, nil
}

// CreateRenderer implments fyne.Widget interface
func (list *DVList) CreateRenderer() fyne.WidgetRenderer {
	list.ExtendBaseWidget(list)
	return newDVListRender(list)
}

func (list *DVList) onTapHeader(col int, asc bool) {
	list.data.Sort(col, asc)
	select {
	case list.actChan <- &listAct{
		act: actSortData,
	}:
		list.Refresh()
	default:
		return
	}
}

func (list *DVList) updateSelection(row int, selected bool) {
	list.mux.Lock()
	if row < 0 || row >= len(list.curSelections) {
		list.mux.Unlock()
		return
	}
	list.curSelections[row] = selected
	if !list.multiSelection {
	L1:
		// single selection
		for i, b := range list.curSelections {
			if b {
				if i != row {
					list.curSelections[i] = false
					// list.SetSelection(i, false)
					// log.Printf("set row %d off", i)

					break L1
				}
			}
		}
	}
	list.mux.Unlock()
	select {
	case list.actChan <- &listAct{
		act: actSetSelection,
	}:
		list.Refresh()
	default:
		return
	}
}

func (list *DVList) onSelectRow(row int, selected bool) {
	// log.Printf("dvlist row %d got selected %v", row, selected)
	defer func() {
		if list.selectionHandler != nil {
			list.selectionHandler(row, selected)
		}
	}()
	list.updateSelection(row, selected)

}

type rowSelectionArgs struct {
	row      int
	selected bool
}

// SetSelection select the specified row if selected is true;
// unselect it otherwise;
// the rowid is the index of Data.Item()
func (list *DVList) SetSelection(rowid int, selected bool) {
	// list.onSelectRow(rowid, selected)
	list.updateSelection(rowid, selected)

}

// CurrentSelections returns a slice of bool, one for each record; true if selected
func (list *DVList) CurrentSelections() (selections []bool) {
	list.mux.RLock()
	defer list.mux.RUnlock()
	for _, b := range list.curSelections {
		selections = append(selections, b)
	}
	return
}

// FirstSelected returns record index in data.Item() of first selected row;
// -1 if none is selected
func (list *DVList) FirstSelected() int {
	list.mux.RLock()
	defer list.mux.RUnlock()
	for i, b := range list.curSelections {
		if b {
			return i
		}
	}
	return -1
}

func (list *DVList) onScroll(offset uint32) {
	// log.Printf("dvlist scroll to offset %d", offset)
	list.mux.RLock()
	atomic.StoreUint32(list.curStartPoint,
		uint32((int(offset)*list.data.Len())/scrollbar.OffsetResolution))
	list.mux.RUnlock()
	select {
	case list.actChan <- &listAct{
		act: actScrollBody,
	}:
		list.Refresh()
	default:
		return
	}
}

// Scrolled implment fyne.Scrollable interface (e.g. mouse wheel scrolling)
func (list *DVList) Scrolled(evt *fyne.ScrollEvent) {
	if evt.Scrolled.DY < 0 && atomic.LoadUint32(list.curStartPoint) < uint32(list.data.Len())-1 {
		//go down
		list.ScrollTo(int(atomic.AddUint32(list.curStartPoint, 1)))
	} else if evt.Scrolled.DY > 0 && atomic.LoadUint32(list.curStartPoint) > 0 {
		//go up
		list.ScrollTo(int(atomic.AddUint32(list.curStartPoint, ^uint32(0))))
	}
}

// TypedKey could be linked to Parent window keyevent by using window.Canvas().SetOnTypedKey()
func (list *DVList) TypedKey(evt *fyne.KeyEvent) {
	curpoint := int(atomic.LoadUint32(list.curStartPoint))
	switch evt.Name {
	case fyne.KeyUp:
		curpoint--
		if curpoint < 0 {
			return
		}
	case fyne.KeyDown:
		curpoint++
		if curpoint >= list.data.Len() {
			return
		}
	case fyne.KeyPageDown:
		curpoint += int(atomic.LoadUint32(list.curShowedLines)) - 1
	case fyne.KeyPageUp:
		curpoint -= int(atomic.LoadUint32(list.curShowedLines)) - 1
	case fyne.KeyHome:
		curpoint = 0
	case fyne.KeyEnd:
		curpoint = list.data.Len() - 1
	}

	if curpoint < 0 {
		curpoint = 0
	}
	if curpoint >= list.data.Len() {
		curpoint = list.data.Len() - 1
	}
	list.ScrollTo(curpoint)
}

// ScrollTo make the row id as first row displayed
func (list *DVList) ScrollTo(id int) {
	atomic.StoreUint32(list.curStartPoint, uint32(id))
	select {
	case list.actChan <- &listAct{
		act: actScrollBody,
	}:
		list.Refresh()
	default:
		return
	}
}
func (list *DVList) SetData(d Data) {
	list.mux.Lock()
	list.data = d
	list.curSelections = []bool{}
	for i := 0; i < d.Len(); i++ {
		list.curSelections = append(list.curSelections, false)
	}
	list.mux.Unlock()
	select {
	case list.actChan <- &listAct{
		act: actSetData,
	}:
		list.Refresh()
	default:
		return
	}
}

type dvListRender struct {
	list             *DVList
	header           *listRow
	body             []*listRow
	scrollbar        *scrollbar.SBar
	overallContainer *fyne.Container
	// bottomContainer  *fyne.Container
	lastStartPoint uint32
	lastSize       fyne.Size
}

func newDVListRender(list *DVList) *dvListRender {
	dlr := &dvListRender{
		list:      list,
		scrollbar: scrollbar.NewSBar(list.onScroll, false),
	}
	dlr.overallContainer = container.NewWithoutLayout()
	dlr.header = NewListHeaderRow(list.data.Fields(),
		scrollbar.BarWidth, list.onTapHeader, dlr.onDragHeader, nil,
		list.headerLook)
	return dlr
}

func (dlr *dvListRender) onDragHeader(newarr []float32) {
	for _, row := range dlr.body {
		row.mux.Lock()
		for i, r := range newarr {
			row.columnWidths[i] = r
		}
		row.mux.Unlock()
		// row.reLayout(false)
	}
	dlr.layout(dlr.list.Size(), listLayoutRedoBody)
	// canvas.Refresh(dlr.list)
}

func (dlr *dvListRender) loadData(layoutsize fyne.Size) {
	// bodyfsize, bodystyle, _ := getBodyLook()
	// bodyunitSize := getUnitSize(dlr.list.bodyLook.TextSize, dlr.list.bodyLook.TextStyle)
	//log.Printf("list loading data for %v", layoutsize)
	dlr.list.mux.Lock()
	defer dlr.list.mux.Unlock()
	dlr.body = []*listRow{}
	startPoint := int(atomic.LoadUint32(dlr.list.curStartPoint))
	// if dlr.list.data.Len()-startPoint <= layoutsize.Height/bodyunitSize.Height {
	// 	startPoint = dlr.list.data.Len() - (layoutsize.Height/bodyunitSize.Height + 1)
	// }
	// if startPoint < 0 {
	// 	startPoint = 0
	// }
	// atomic.StoreUint32(dlr.list.curStartPoint, uint32(startPoint))
	// log.Printf("adjust start pt is %d", startPoint)
	var h float32 = 0
	defer func() {
		if dlr.list.data.Len() == 0 {
			return
		}
		newoffset := uint32((startPoint * scrollbar.OffsetResolution) / dlr.list.data.Len())
		// log.Printf("set new offset %d", newoffset)
		dlr.scrollbar.SetOffset(newoffset)
	}()
	showedLine := 0
	for i := startPoint; i < dlr.list.data.Len(); i++ {
		row := NewListBodyRow(i, dlr.list.data.Item(i), dlr.list.onSelectRow,
			dlr.list.doubleClickHandler, dlr.header.columnWidths,
			dlr.list.bodyLook)
		dlr.body = append(dlr.body, row)
		showedLine++
		h += row.MinSize().Height
		if h > layoutsize.Height {
			atomic.StoreUint32(dlr.list.curShowedLines, uint32(showedLine))
			return
		}
	}
}
func (dlr *dvListRender) BackgroundColor() color.Color {
	return color.Transparent
}
func (dlr *dvListRender) Destroy() {
}

const (
	sepHeight        = 1
	sectionGapHeight = 1
)

func (dlr *dvListRender) getBottomAllowanceSize(layoutsize fyne.Size) fyne.Size {
	// minus header, speroator section and scrollbar
	newwidth := layoutsize.Width - scrollbar.BarWidth
	newheight := layoutsize.Height - dlr.header.MinSize().Height - sepHeight - sectionGapHeight
	return fyne.NewSize(newwidth, newheight)
}

const (
	listLayoutRedoALL = iota
	listLayoutRedoBody
	listLayoutLazy
	listLayoutUpdateSelection
)

// recreate everything if force is true
func (dlr *dvListRender) layout(layoutsize fyne.Size, mode int) {
	// debug.PrintStack()
	// log.Printf("body layout for %v, with startpoint %d with mode %d", layoutsize, atomic.LoadUint32(dlr.list.curStartPoint), mode)
	switch mode {
	case listLayoutUpdateSelection:
		for _, row := range dlr.body {
			if dlr.list.curSelections[row.rowID] {
				row.SetSelection(true, true)
			} else {
				row.SetSelection(false, true)
			}
			// dlr.overallContainer.Add(row)

		}
		return
	}
	if mode == listLayoutLazy {
		dlr.list.mux.RLock()
		if dlr.overallContainer != nil &&
			dlr.lastStartPoint == atomic.LoadUint32(dlr.list.curStartPoint) &&
			dlr.lastSize.Subtract(layoutsize).IsZero() {
			dlr.list.mux.RUnlock()
			return
		}
		if dlr.overallContainer != nil &&
			dlr.lastSize.Subtract(layoutsize).IsZero() &&
			len(dlr.body) == dlr.list.data.Len() {
			dlr.list.mux.RUnlock()
			return
		}
		dlr.list.mux.RUnlock()
	}

	if mode == listLayoutRedoALL {
		dlr.header = NewListHeaderRow(dlr.list.data.Fields(),
			scrollbar.BarWidth, dlr.list.onTapHeader, dlr.onDragHeader, nil,
			dlr.list.headerLook)
	}
	dlr.lastSize = layoutsize
	mainbodysize := dlr.getBottomAllowanceSize(layoutsize)
	dlr.loadData(mainbodysize)
	// headerfsize, headerstyle, _ := getHeaderLook()
	unitSize := getUnitSize(dlr.list.headerLook.TextSize, dlr.list.headerLook.TextStyle)
	dlr.list.mux.Lock()
	defer dlr.list.mux.Unlock()
	dlr.overallContainer = container.NewWithoutLayout()
	dlr.header.Resize(fyne.NewSize(layoutsize.Width, unitSize.Height))
	dlr.header.Move(fyne.NewPos(0, 0))
	dlr.overallContainer.Add(dlr.header)
	sep := widget.NewSeparator()
	sep.Resize(fyne.NewSize(layoutsize.Width, sepHeight))
	sep.Move(fyne.NewPos(0, dlr.header.MinSize().Height+sectionGapHeight))
	dlr.overallContainer.Add(sep)
	for i, row := range dlr.body {
		row.Resize(fyne.NewSize(mainbodysize.Width, row.MinSize().Height))
		row.Move(fyne.NewPos(0, (layoutsize.Height-mainbodysize.Height)+float32(i)*row.MinSize().Height))
		if dlr.list.curSelections[row.rowID] {
			row.SetSelection(true, true)
		} else {
			row.SetSelection(false, true)
		}
		dlr.overallContainer.Add(row)
	}

	dlr.scrollbar.Resize(fyne.NewSize(scrollbar.BarWidth, mainbodysize.Height))
	dlr.scrollbar.Move(fyne.NewPos(mainbodysize.Width, layoutsize.Height-mainbodysize.Height))
	dlr.overallContainer.Add(dlr.scrollbar)

}

func (dlr *dvListRender) Layout(layoutsize fyne.Size) {
	dlr.layout(layoutsize, listLayoutLazy)
}
func (dlr *dvListRender) MinSize() fyne.Size {
	dlr.list.mux.RLock()
	defer dlr.list.mux.RUnlock()
	if len(dlr.body) > 0 {
		return dlr.header.MinSize().Add(dlr.body[0].MinSize())
	}

	if dlr.header != nil {
		dlr.header.MinSize()
	}
	return fyne.NewSize(10, 10)
}

func (dlr *dvListRender) Objects() []fyne.CanvasObject {
	dlr.list.mux.RLock()
	defer dlr.list.mux.RUnlock()
	return []fyne.CanvasObject{dlr.overallContainer}
}
func (dlr *dvListRender) Refresh() {
	select {
	case act := <-dlr.list.actChan:
		switch act.act {
		case actSetData:
			dlr.layout(dlr.list.Size(), listLayoutRedoALL)
		case actScrollBody:
			dlr.layout(dlr.list.Size(), listLayoutRedoBody)
		case actSortData:
			dlr.layout(dlr.list.Size(), listLayoutRedoBody)
		case actSetSelection:
			dlr.layout(dlr.list.Size(), listLayoutUpdateSelection)
			// dlr.layout(dlr.list.Size(), listLayoutRedoBody)
			// rowindex := act.arg.(rowSelectionArgs).row - int(atomic.LoadUint32(dlr.list.curStartPoint))
			// if rowindex < 0 || rowindex >= len(dlr.body) {
			// 	return
			// }
			// log.Printf("set row %d selection true", rowindex)
			// dlr.body[rowindex].SetSelection(act.arg.(rowSelectionArgs).selected, true)
		}
	default:
	}
	canvas.Refresh(dlr.list)
}
