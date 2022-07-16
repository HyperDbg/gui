package myTable

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"sync"
	"sync/atomic"
)

type StyleObj struct {
	TextSize  float32
	TextColor color.Color
	TextStyle fyne.TextStyle
	Alignment fyne.TextAlign
}

func defaultBodyStyle() *StyleObj {
	return &StyleObj{
		TextSize:  theme.TextSize(),
		TextColor: theme.ForegroundColor(),
		TextStyle: fyne.TextStyle{},
		Alignment: fyne.TextAlignLeading,
	}
}

func defaultHeaderStyle() *StyleObj {
	bodyLook := defaultBodyStyle()
	return &StyleObj{
		TextSize:  bodyLook.TextSize + 2,
		TextColor: bodyLook.TextColor,
		TextStyle: fyne.TextStyle{Bold: true},
		Alignment: fyne.TextAlignCenter,
	}
}

// type headerSep struct {
// 	*widget.Separator
// 	fieldID int
// 	row     *ListRow
// }

// func newHeaderSep(row *ListRow, field int) *headerSep {
// 	return &headerSep{
// 		Separator: widget.NewSeparator(),
// 		fieldID:   field,
// 		row:       row,
// 	}
// }

// func (sep *headerSep) Dragged(evt *fyne.DragEvent) {
// 	log.Printf("field %d draged X %d", sep.fieldID, evt.DraggedX)
// }

// func (sep *headerSep) DragEnd() {

// }

type rowActType int

const (
	rowActSetSelection rowActType = iota
	rowActHeaderSetSort
	rowActReLayout
)

type rowAct struct {
	act rowActType
	val interface{}
}

const (
	headerSortNone uint32 = iota
	headerSortAscend
	headerSortDescend
)

type OnSelectionHandler func(id int, selected bool)
type DoubleClickHandler func(id int)
type headerOnTapHandler func(id int, ascend bool)
type notifyListDragHandler func([]float32)

type tableRow struct {
	widget.BaseWidget
	rowID                int //this should be the index in the data
	fieldList            []string
	selected             *uint32 // non-zero means selected
	header               bool
	actChan              chan *rowAct
	rightPadding         float32
	selectHandler        OnSelectionHandler
	headerTapHandler     headerOnTapHandler
	headerDragHandler    notifyListDragHandler
	headerSortDirections []*uint32 // non-zero means ascend
	columnWidths         []float32 // percentage width for each field, sum<=MaxFields0
	mux                  *sync.RWMutex
	look                 *StyleObj
	doubleClickHandler   DoubleClickHandler
}

func (row *tableRow) SetSelectHandler(selectHandler OnSelectionHandler) {
	row.selectHandler = selectHandler
}

const (
	rowActChanDepth = 16
	MaxFields       = 100
)

type Header struct {
	Title       []string
	Padding     float32
	Handler     headerOnTapHandler
	DragHandler notifyListDragHandler
	Arr         []float32
	Style       *StyleObj
}

func NewHeader(header Header) *tableRow {
	return NewTableHeaderRow(
		header.Title,
		header.Padding,
		header.Handler,
		header.DragHandler,
		header.Arr,
		header.Style,
	)
}
func NewTableHeaderRow(header []string, padding float32, handler headerOnTapHandler, dh notifyListDragHandler, arr []float32, look *StyleObj) *tableRow {
	if len(header) > MaxFields {
		return nil
	}
	row := &tableRow{
		header:            true,
		actChan:           make(chan *rowAct, rowActChanDepth),
		rightPadding:      padding,
		headerTapHandler:  handler,
		headerDragHandler: dh,
		columnWidths:      []float32{},
		mux:               new(sync.RWMutex),
		look:              look,
	}
	for i, s := range header {
		row.fieldList = append(row.fieldList, s)
		row.headerSortDirections = append(row.headerSortDirections, new(uint32))
		atomic.StoreUint32(row.headerSortDirections[i], headerSortNone)
	}
	if len(arr) == 0 {
		for i := 0; i < len(header); i++ {
			row.columnWidths = append(row.columnWidths, float32(MaxFields/len(header)))
		}
	} else {
		if len(arr) != len(header) {
			return nil
		}
		for _, p := range arr {
			row.columnWidths = append(row.columnWidths, p)
		}
	}
	row.ExtendBaseWidget(row)
	return row
}

func NewTableBodyRow(id int, body []string,
	handler OnSelectionHandler,
	clickHandler DoubleClickHandler,
	arr []float32, look *StyleObj) *tableRow {
	row := &tableRow{
		rowID:              id,
		selected:           new(uint32),
		header:             false,
		actChan:            make(chan *rowAct, rowActChanDepth),
		selectHandler:      handler,
		columnWidths:       []float32{},
		mux:                new(sync.RWMutex),
		look:               look,
		doubleClickHandler: clickHandler,
	}

	if len(arr) != len(body) {
		return nil
	}
	for i, s := range body {
		row.fieldList = append(row.fieldList, s)
		row.columnWidths = append(row.columnWidths, arr[i])
	}
	for i := 0; i < len(body); i++ {
		row.columnWidths = append(row.columnWidths, float32(MaxFields/len(body)))
	}
	atomic.StoreUint32(row.selected, 0)
	row.ExtendBaseWidget(row)
	return row
}

func (row *tableRow) CreateRenderer() fyne.WidgetRenderer {
	row.ExtendBaseWidget(row)
	return newTableRowRender(row)
}

func (row *tableRow) SetSelection(selected bool, admin bool) {
	if selected == (atomic.LoadUint32(row.selected) == 1) {
		//same state, no need update
		// log.Printf("row %d set selection done,no change", row.rowID)
		return
	}
	// log.Printf("row %d set selection done,need change", row.rowID)
	if selected {
		atomic.StoreUint32(row.selected, 1)
	} else {
		atomic.StoreUint32(row.selected, 0)
	}
	select {
	case row.actChan <- &rowAct{
		act: rowActSetSelection,
	}:
		row.Refresh()
		if row.selectHandler != nil && !admin {
			row.selectHandler(row.rowID, selected)
		}
	default:
		return
	}
}

func (row *tableRow) SetColumnWidth(id int, width float32) error {
	if width == 0 {
		return fmt.Errorf("not dragging")
	}
	row.mux.Lock()
	defer row.mux.Unlock()
	if id < 0 || id >= len(row.columnWidths)-1 {
		if id == len(row.columnWidths)-1 {
			return fmt.Errorf("can't drag last column")
		} else {
			return fmt.Errorf("invalid column id %d", id)
		}
	}
	var previousRightTotal float32 = 0
	for i := id + 1; i < len(row.columnWidths); i++ {
		previousRightTotal += row.columnWidths[i]
	}
	if previousRightTotal == 0 {
		if width > 0 {
			return fmt.Errorf("reached the end")
		}
		previousRightTotal = 1
	}
	if width > 0 {
		row.columnWidths[id]++
	} else {
		row.columnWidths[id]--
		if row.columnWidths[id] < 1 {
			row.columnWidths[id] = 1
		}
	}
	var allocated float32 = 0
	for i := 0; i <= id; i++ {
		allocated += row.columnWidths[i]
	}
	available := MaxFields - allocated
	for i := id + 1; i < len(row.columnWidths); i++ {
		row.columnWidths[i] = (available * row.columnWidths[i]) / previousRightTotal
		if row.columnWidths[i] < 1 {
			row.columnWidths[i] = 1
		}
	}
	// log.Printf("now columnWidths is %v", row.columnWidths)
	return nil
}

func (row *tableRow) reLayout(refresh bool) {
	select {
	case row.actChan <- &rowAct{act: rowActReLayout, val: refresh}:
		row.Refresh()
	default:
	}
}

func (row *tableRow) dragHandler(metadata interface{}, evt *fyne.DragEvent) {
	id := metadata.(listCord).Column
	err := row.SetColumnWidth(id, evt.Dragged.DX)
	if err == nil {
		row.reLayout(true)
		if row.headerDragHandler != nil {
			row.headerDragHandler(row.columnWidths)
		}
	} else {
		// log.Printf("drag skipped,%v", err)
	}
}

func (row *tableRow) doubleTapHandler(metadata interface{}, evt *fyne.PointEvent) {
	if row.doubleClickHandler != nil {
		row.doubleClickHandler(metadata.(listCord).Row)
	}
}
func (row *tableRow) tapHandler(metadata interface{}, evt *fyne.PointEvent) {
	// log.Printf("row %d clicked", row.rowID)
	if !row.header {
		//body
		if atomic.LoadUint32(row.selected) == 0 {
			row.SetSelection(true, false)
		} else {
			row.SetSelection(false, false)
		}
	} else {
		//header
		pos := metadata.(listCord).Column
		currDidirection := atomic.LoadUint32(row.headerSortDirections[pos])
		newDirection := headerSortAscend
		if currDidirection == headerSortNone {

			// atomic.StoreUint32(row.headerSortDirections[pos], headerSortAscend)
			// if row.headerTapHandler != nil {
			// 	row.headerTapHandler(pos, true)
			// }
		} else {
			if currDidirection == headerSortAscend {
				newDirection = headerSortDescend
				// atomic.StoreUint32(row.headerSortDirections[pos], headerSortDescend)
				// if row.headerTapHandler != nil {
				// 	row.headerTapHandler(pos, false)
				// }
			} else {
				newDirection = headerSortAscend
				// atomic.StoreUint32(row.headerSortDirections[pos], headerSortAscend)
				// if row.headerTapHandler != nil {
				// 	row.headerTapHandler(pos, true)
				// }
			}
		}
		for i := range row.headerSortDirections {
			if i != pos {
				atomic.StoreUint32(row.headerSortDirections[i], headerSortNone)
			} else {
				atomic.StoreUint32(row.headerSortDirections[pos], newDirection)
			}
		}
		if row.headerTapHandler != nil {
			row.headerTapHandler(pos, newDirection == headerSortAscend)
		}
		select {
		case row.actChan <- &rowAct{
			act: rowActHeaderSetSort,
		}:
		default:
		}
		row.Refresh()
	}
}

type RowRender struct {
	row              *tableRow
	labelList        []*resizableLabel
	background       *canvas.Rectangle
	innerContainer   *fyne.Container
	overallContainer *fyne.Container
	headerSepList    []*widget.Separator
}

func newTableRowRender(row *tableRow) *RowRender {
	r := &RowRender{
		row: row,
	}
	r.innerContainer = container.NewWithoutLayout()
	for i, s := range row.fieldList {
		options := []ResizableLabelOption{
			WithMetaData(listCord{Row: row.rowID, Column: i}),
			WithTapHandler(row.tapHandler),
		}
		if r.row.header {
			options = append(options, WithDraggingHandler(row.dragHandler))
		} else {
			options = append(options, WithDoubleTapHandler(row.doubleTapHandler))
		}
		label := NewResizableLabel(s, row.look, options...)
		r.innerContainer.Add(label)
		r.labelList = append(r.labelList, label)
		// if r.row.header {
		if i < len(row.fieldList)-1 {
			// if !r.row.header {
			sep := widget.NewSeparator()
			r.headerSepList = append(r.headerSepList, sep)
			r.innerContainer.Add(sep)
			// } else {
			// sep := newHeaderSep(row, i)
			// r.headerSepList = append(r.headerSepList, sep)
			// r.innerContainer.AddRow(sep)
			// }
			// }
		}
	}
	r.background = canvas.NewRectangle(theme.BackgroundColor())
	r.overallContainer = container.NewMax(r.background, r.innerContainer)
	return r
}

func (r *RowRender) BackgroundColor() color.Color {
	return color.Transparent
}
func (r *RowRender) Destroy() {

}

const (
	headerSepWidth         = 1
	sepGapWidth    float32 = 1
)

var ColumnWidth = make(map[int]float32) //SetColumnWidth(id int, width float32)

func (r *RowRender) Layout(layoutSize fyne.Size) {
	newSize := fyne.NewSize(layoutSize.Width-r.row.rightPadding, layoutSize.Height)
	if newSize.Width <= 0 || newSize.Height <= 0 {
		return
	}
	r.background.Resize(newSize)
	r.row.mux.RLock()
	defer r.row.mux.RUnlock()
	var startx float32 = 0
	for i, label := range r.labelList {
		cellWidth := (r.row.columnWidths[i] * newSize.Width) / MaxFields
		if w, ok := ColumnWidth[i]; ok {
			if w != 0 {
				cellWidth = w
			}
		}
		cellWidth -= 20
		label.Resize(fyne.NewSize(cellWidth-2*sepGapWidth-headerSepWidth, newSize.Height))
		label.Move(fyne.NewPos(startx+sepGapWidth, 0))
		label.Refresh()
		startx += cellWidth
		if i < len(r.labelList)-1 {
			r.headerSepList[i].Resize(fyne.NewSize(headerSepWidth, newSize.Height))
			r.headerSepList[i].Move(fyne.NewPos(startx-headerSepWidth-sepGapWidth, 0))
			r.headerSepList[i].Refresh()
		}
	}
}

func (r *RowRender) MinSize() fyne.Size {
	h := r.labelList[0].MinSize().Height
	var w float32 = 0
	for _, label := range r.labelList {
		w += label.MinSize().Width
	}
	return fyne.NewSize(w, h)
	// return r.innerContainer.MinSize()
}
func (r *RowRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.overallContainer}
}

const (
	ascendChar  = rune(0x02c4)
	descendChar = rune(0x02c5)
)

func (r *RowRender) Refresh() {
	select {
	case act := <-r.row.actChan:
		switch act.act {
		case rowActSetSelection:
			if atomic.LoadUint32(r.row.selected) != 0 {
				r.background.FillColor = theme.FocusColor()
			} else {
				r.background.FillColor = theme.BackgroundColor()
			}
			r.background.Refresh()
			canvas.Refresh(r.background)
			canvas.Refresh(r.row)
			canvas.Refresh(r.background)
		case rowActHeaderSetSort:
			r.row.mux.Lock()
			for i := range r.row.headerSortDirections {
				newDirection := atomic.LoadUint32(r.row.headerSortDirections[i])
				newText := r.labelList[i].Get()
				if newText[len(newText)-1] == ascendChar || newText[len(newText)-1] == descendChar {
					newText = newText[:len(newText)-1]
				}
				switch newDirection {
				case headerSortAscend:
					newText = append(newText, ascendChar)
				case headerSortDescend:
					newText = append(newText, descendChar)
				}
				r.labelList[i].Set(string(newText))
			}
			r.row.mux.Unlock()
		case rowActReLayout:
			r.Layout(r.row.Size())

		}
	default:
		return
	}
	canvas.Refresh(r.row)
}
