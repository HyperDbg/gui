package myTable

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/myTable/scrollbar"
	"image/color"
	"sync"
	"sync/atomic"
)

type Interface interface {
	Header() []string
	Rows(id int) []string
	Len() int
	Sort(id int, ascend bool)
	Filter(row string, id int)
	Append(data any)
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

type MyTable struct {
	widget.BaseWidget
	data    Interface
	mux     *sync.RWMutex
	actChan chan *listAct
	// curOffset            *uint32 //current offset value of scrollbar
	curStartPoint          *uint32 //the index of first row showed
	curShowedLines         *uint32
	multiSelection         bool
	curSelections          []bool
	headerStyle, bodyStyle *StyleObj
	selectionHandler       OnSelectionHandler
	doubleClickHandler     DoubleClickHandler
}

type Option func(table *MyTable)

func WithMultiSelections() Option {
	return func(list *MyTable) {
		list.multiSelection = true
	}
}

func WithHeaderStyle(style *StyleObj) Option {
	return func(list *MyTable) {
		list.headerStyle = style
	}
}

func WithBodyLook(style *StyleObj) Option {
	return func(list *MyTable) {
		list.bodyStyle = style
	}
}

// WithSelectionHandler specifies h as a handler function gets called when a row is selected
func WithSelectionHandler(h OnSelectionHandler) Option {
	return func(list *MyTable) {
		list.selectionHandler = h
	}
}

// WithDoubleClickHandler specifies h as a handler function gets called when a row is double clicked
func WithDoubleClickHandler(h DoubleClickHandler) Option {
	return func(table *MyTable) {
		table.doubleClickHandler = h
	}
}

// NewDVList return a new MyTable instance with the specified data
func NewDVList(data Interface, options ...Option) (*MyTable, error) {
	if len(data.Header()) > MaxFields {
		return nil, fmt.Errorf("data has more fields (%d) than allowed (%d)", len(data.Header()), MaxFields)
	}
	list := &MyTable{
		data:           data,
		mux:            new(sync.RWMutex),
		actChan:        make(chan *listAct, actChanDepth),
		curStartPoint:  new(uint32),
		curShowedLines: new(uint32),
		curSelections:  []bool{},
		headerStyle:    defaultHeaderStyle(),
		bodyStyle:      defaultBodyStyle(),
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

func (t *MyTable) CreateRenderer() fyne.WidgetRenderer {
	t.ExtendBaseWidget(t)
	return newTableRender(t)
}

func (t *MyTable) onTapHeader(col int, asc bool) {
	t.data.Sort(col, asc)
	select {
	case t.actChan <- &listAct{
		act: actSortData,
	}:
		t.Refresh()
	default:
		return
	}
}

func (t *MyTable) updateSelection(row int, selected bool) {
	t.mux.Lock()
	if row < 0 || row >= len(t.curSelections) {
		t.mux.Unlock()
		return
	}
	t.curSelections[row] = selected
	if !t.multiSelection {
	L1:
		// single selection
		for i, b := range t.curSelections {
			if b {
				if i != row {
					t.curSelections[i] = false
					// t.SetSelection(i, false)
					// log.Printf("set row %d off", i)

					break L1
				}
			}
		}
	}
	t.mux.Unlock()
	select {
	case t.actChan <- &listAct{
		act: actSetSelection,
	}:
		t.Refresh()
	default:
		return
	}
}

func (t *MyTable) onSelectRow(row int, selected bool) {
	// log.Printf("dvlist row %d got selected %v", row, selected)
	defer func() {
		if t.selectionHandler != nil {
			t.selectionHandler(row, selected)
		}
	}()
	t.updateSelection(row, selected)

}

type rowSelectionArgs struct {
	row      int
	selected bool
}

// SetSelection select the specified row if selected is true;
// unselect it otherwise;
// the rowid is the index of Interface.Item()
func (t *MyTable) SetSelection(rowid int, selected bool) {
	// t.onSelectRow(rowid, selected)
	t.updateSelection(rowid, selected)

}

// CurrentSelections returns a slice of bool, one for each record; true if selected
func (t *MyTable) CurrentSelections() (selections []bool) {
	t.mux.RLock()
	defer t.mux.RUnlock()
	for _, b := range t.curSelections {
		selections = append(selections, b)
	}
	return
}

// FirstSelected returns record index in data.Rows() of first selected row;
// -1 if none is selected
func (t *MyTable) FirstSelected() int {
	t.mux.RLock()
	defer t.mux.RUnlock()
	for i, b := range t.curSelections {
		if b {
			return i
		}
	}
	return -1
}

func (t *MyTable) onScroll(offset uint32) {
	// log.Printf("dvlist scroll to offset %d", offset)
	t.mux.RLock()
	atomic.StoreUint32(t.curStartPoint,
		uint32((int(offset)*t.data.Len())/scrollbar.OffsetResolution))
	t.mux.RUnlock()
	select {
	case t.actChan <- &listAct{
		act: actScrollBody,
	}:
		t.Refresh()
	default:
		return
	}
}

// Scrolled implment fyne.Scrollable interface (e.g. mouse wheel scrolling)
func (t *MyTable) Scrolled(evt *fyne.ScrollEvent) {
	if evt.Scrolled.DY < 0 && atomic.LoadUint32(t.curStartPoint) < uint32(t.data.Len())-1 {
		//go down
		t.ScrollTo(int(atomic.AddUint32(t.curStartPoint, 1)))
	} else if evt.Scrolled.DY > 0 && atomic.LoadUint32(t.curStartPoint) > 0 {
		//go up
		t.ScrollTo(int(atomic.AddUint32(t.curStartPoint, ^uint32(0))))
	}
}

// TypedKey could be linked to Parent window keyevent by using window.Canvas().SetOnTypedKey()
func (t *MyTable) TypedKey(evt *fyne.KeyEvent) {
	curpoint := int(atomic.LoadUint32(t.curStartPoint))
	switch evt.Name {
	case fyne.KeyUp:
		curpoint--
		if curpoint < 0 {
			return
		}
	case fyne.KeyDown:
		curpoint++
		if curpoint >= t.data.Len() {
			return
		}
	case fyne.KeyPageDown:
		curpoint += int(atomic.LoadUint32(t.curShowedLines)) - 1
	case fyne.KeyPageUp:
		curpoint -= int(atomic.LoadUint32(t.curShowedLines)) - 1
	case fyne.KeyHome:
		curpoint = 0
	case fyne.KeyEnd:
		curpoint = t.data.Len() - 1
	}

	if curpoint < 0 {
		curpoint = 0
	}
	if curpoint >= t.data.Len() {
		curpoint = t.data.Len() - 1
	}
	t.ScrollTo(curpoint)
}

// ScrollTo make the row id as first row displayed
func (t *MyTable) ScrollTo(id int) {
	atomic.StoreUint32(t.curStartPoint, uint32(id))
	select {
	case t.actChan <- &listAct{
		act: actScrollBody,
	}:
		t.Refresh()
	default:
		return
	}
}
func (t *MyTable) SetData(d Interface) {
	t.mux.Lock()
	t.data = d
	t.curSelections = []bool{}
	for i := 0; i < d.Len(); i++ {
		t.curSelections = append(t.curSelections, false)
	}
	t.mux.Unlock()
	select {
	case t.actChan <- &listAct{
		act: actSetData,
	}:
		t.Refresh()
	default:
		return
	}
}

type TableRender struct {
	table            *MyTable
	header           *tableRow
	body             []*tableRow
	scrollbar        *scrollbar.ScrollBar
	overallContainer *fyne.Container
	// bottomContainer  *fyne.Container
	lastStartPoint uint32
	lastSize       fyne.Size
}

func newTableRender(list *MyTable) *TableRender {
	dlr := &TableRender{
		table:     list,
		scrollbar: scrollbar.NewSBar(list.onScroll, false),
	}
	dlr.overallContainer = container.NewWithoutLayout()
	dlr.header = NewTableHeaderRow(list.data.Header(),
		scrollbar.BarWidth, list.onTapHeader, dlr.onDragHeader, nil,
		list.headerStyle)
	return dlr
}

func (r *TableRender) onDragHeader(widths []float32) {
	for _, row := range r.body {
		row.mux.Lock()
		for i, r := range widths {
			row.columnWidths[i] = r
		}
		row.mux.Unlock()
		// row.reLayout(false)
	}
	r.layout(r.table.Size(), listLayoutRedoBody)
	// canvas.Refresh(r.table)
}

func (r *TableRender) loadData(layoutSize fyne.Size) {
	// bodyfsize, bodystyle, _ := getBodyLook()
	// bodyunitSize := getUnitSize(r.table.bodyStyle.TextSize, r.table.bodyStyle.TextStyle)
	//log.Printf("table loading data for %v", layoutSize)
	r.table.mux.Lock()
	defer r.table.mux.Unlock()
	r.body = []*tableRow{}
	startPoint := int(atomic.LoadUint32(r.table.curStartPoint))
	// if r.table.data.Len()-startPoint <= layoutSize.Height/bodyunitSize.Height {
	// 	startPoint = r.table.data.Len() - (layoutSize.Height/bodyunitSize.Height + 1)
	// }
	// if startPoint < 0 {
	// 	startPoint = 0
	// }
	// atomic.StoreUint32(r.table.curStartPoint, uint32(startPoint))
	// log.Printf("adjust start pt is %d", startPoint)
	var h float32 = 0
	defer func() {
		if r.table.data.Len() == 0 {
			return
		}
		newOffset := uint32((startPoint * scrollbar.OffsetResolution) / r.table.data.Len())
		// log.Printf("set new offset %d", newOffset)
		r.scrollbar.SetOffset(newOffset)
	}()
	showedLine := 0
	for i := startPoint; i < r.table.data.Len(); i++ {
		row := NewTableBodyRow(i, r.table.data.Rows(i), r.table.onSelectRow,
			r.table.doubleClickHandler, r.header.columnWidths,
			r.table.bodyStyle)
		r.body = append(r.body, row)
		showedLine++
		h += row.MinSize().Height
		if h > layoutSize.Height {
			atomic.StoreUint32(r.table.curShowedLines, uint32(showedLine))
			return
		}
	}
}
func (r *TableRender) BackgroundColor() color.Color {
	return color.Transparent
}
func (r *TableRender) Destroy() {
}

const (
	sepHeight        = 1
	sectionGapHeight = 1
)

func (r *TableRender) getBottomAllowanceSize(layoutsize fyne.Size) fyne.Size {
	// minus header, speroator section and scrollbar
	newwidth := layoutsize.Width - scrollbar.BarWidth
	newheight := layoutsize.Height - r.header.MinSize().Height - sepHeight - sectionGapHeight
	return fyne.NewSize(newwidth, newheight)
}

const (
	listLayoutRedoALL = iota
	listLayoutRedoBody
	listLayoutLazy
	listLayoutUpdateSelection
)

// recreate everything if force is true
func (r *TableRender) layout(layoutsize fyne.Size, mode int) {
	// debug.PrintStack()
	// log.Printf("body layout for %v, with startpoint %d with mode %d", layoutsize, atomic.LoadUint32(r.table.curStartPoint), mode)
	switch mode {
	case listLayoutUpdateSelection:
		for _, row := range r.body {
			if r.table.curSelections[row.rowID] {
				row.SetSelection(true, true)
			} else {
				row.SetSelection(false, true)
			}
			// r.overallContainer.AddRow(row)

		}
		return
	}
	if mode == listLayoutLazy {
		r.table.mux.RLock()
		if r.overallContainer != nil &&
			r.lastStartPoint == atomic.LoadUint32(r.table.curStartPoint) &&
			r.lastSize.Subtract(layoutsize).IsZero() {
			r.table.mux.RUnlock()
			return
		}
		if r.overallContainer != nil &&
			r.lastSize.Subtract(layoutsize).IsZero() &&
			len(r.body) == r.table.data.Len() {
			r.table.mux.RUnlock()
			return
		}
		r.table.mux.RUnlock()
	}

	if mode == listLayoutRedoALL {
		r.header = NewTableHeaderRow(r.table.data.Header(),
			scrollbar.BarWidth, r.table.onTapHeader, r.onDragHeader, nil,
			r.table.headerStyle)
	}
	r.lastSize = layoutsize
	mainbodysize := r.getBottomAllowanceSize(layoutsize)
	r.loadData(mainbodysize)
	// headerfsize, headerstyle, _ := getHeaderLook()
	unitSize := getUnitSize(r.table.headerStyle.TextSize, r.table.headerStyle.TextStyle)
	r.table.mux.Lock()
	defer r.table.mux.Unlock()
	r.overallContainer = container.NewWithoutLayout()
	r.header.Resize(fyne.NewSize(layoutsize.Width, unitSize.Height))
	r.header.Move(fyne.NewPos(0, 0))
	r.overallContainer.Add(r.header)
	sep := widget.NewSeparator()
	sep.Resize(fyne.NewSize(layoutsize.Width, sepHeight))
	sep.Move(fyne.NewPos(0, r.header.MinSize().Height+sectionGapHeight))
	r.overallContainer.Add(sep)
	for i, row := range r.body {
		row.Resize(fyne.NewSize(mainbodysize.Width, row.MinSize().Height))
		row.Move(fyne.NewPos(0, (layoutsize.Height-mainbodysize.Height)+float32(i)*row.MinSize().Height))
		if r.table.curSelections[row.rowID] {
			row.SetSelection(true, true)
		} else {
			row.SetSelection(false, true)
		}
		r.overallContainer.Add(row)
	}

	r.scrollbar.Resize(fyne.NewSize(scrollbar.BarWidth, mainbodysize.Height))
	r.scrollbar.Move(fyne.NewPos(mainbodysize.Width, layoutsize.Height-mainbodysize.Height))
	r.overallContainer.Add(r.scrollbar)

}

func (r *TableRender) Layout(layoutsize fyne.Size) {
	r.layout(layoutsize, listLayoutLazy)
}
func (r *TableRender) MinSize() fyne.Size {
	r.table.mux.RLock()
	defer r.table.mux.RUnlock()
	if len(r.body) > 0 {
		return r.header.MinSize().Add(r.body[0].MinSize())
	}

	if r.header != nil {
		r.header.MinSize()
	}
	return fyne.NewSize(10, 10)
}

func (r *TableRender) Objects() []fyne.CanvasObject {
	r.table.mux.RLock()
	defer r.table.mux.RUnlock()
	return []fyne.CanvasObject{r.overallContainer}
}
func (r *TableRender) Refresh() {
	select {
	case act := <-r.table.actChan:
		switch act.act {
		case actSetData:
			r.layout(r.table.Size(), listLayoutRedoALL)
		case actScrollBody:
			r.layout(r.table.Size(), listLayoutRedoBody)
		case actSortData:
			r.layout(r.table.Size(), listLayoutRedoBody)
		case actSetSelection:
			r.layout(r.table.Size(), listLayoutUpdateSelection)
			// r.layout(r.table.Size(), listLayoutRedoBody)
			// rowindex := act.arg.(rowSelectionArgs).row - int(atomic.LoadUint32(r.table.curStartPoint))
			// if rowindex < 0 || rowindex >= len(r.body) {
			// 	return
			// }
			// log.Printf("set row %d selection true", rowindex)
			// r.body[rowindex].SetSelection(act.arg.(rowSelectionArgs).selected, true)
		}
	default:
	}
	canvas.Refresh(r.table)
}
