package dvlist

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

type Look struct {
	TextSize  float32
	TextColor color.Color
	TextStyle fyne.TextStyle
	Alignment fyne.TextAlign
}

func defaultBodyLook() *Look {

	return &Look{
		TextSize:  theme.TextSize(),
		TextColor: theme.ForegroundColor(),
		TextStyle: fyne.TextStyle{},
		Alignment: fyne.TextAlignLeading,
	}
}

func defaultHeaderLook() *Look {
	bodyLook := defaultBodyLook()
	return &Look{
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
type headerOnTapHandler func(colid int, ascend bool)
type notifyListDraghandler func([]float32)

type listRow struct {
	widget.BaseWidget
	rowID                int //this should be the index in the data
	fieldList            []string
	selected             *uint32 // non-zero means selected
	header               bool
	actChan              chan *rowAct
	rightPadding         float32
	selectHandler        OnSelectionHandler
	headerTapHandler     headerOnTapHandler
	headerDragHandler    notifyListDraghandler
	headerSortDirections []*uint32 // non-zero means ascend
	columnWidths         []float32 // percentage width for each field, sum<=MaxFields0
	mux                  *sync.RWMutex
	look                 *Look
	doubleClickHandler   DoubleClickHandler
}

func (row *listRow) SetSelectHandler(selectHandler OnSelectionHandler) {
	row.selectHandler = selectHandler
}

const (
	rowActChanDepth = 16
	MaxFields       = 100
)

type Header struct {
	List    []string
	Padding float32
	Handler headerOnTapHandler
	Dh      notifyListDraghandler
	Arr     []float32
	Look    *Look
}

func NewHeader(header Header) *listRow {
	return NewListHeaderRow(
		header.List,
		header.Padding,
		header.Handler,
		header.Dh,
		header.Arr,
		header.Look,
	)
}
func NewListHeaderRow(list []string, padding float32, handler headerOnTapHandler, dh notifyListDraghandler, arr []float32, look *Look) *listRow {
	if len(list) > MaxFields {
		return nil
	}
	row := &listRow{
		header:            true,
		actChan:           make(chan *rowAct, rowActChanDepth),
		rightPadding:      padding,
		headerTapHandler:  handler,
		headerDragHandler: dh,
		columnWidths:      []float32{},
		mux:               new(sync.RWMutex),
		look:              look,
	}
	for i, s := range list {
		row.fieldList = append(row.fieldList, s)
		row.headerSortDirections = append(row.headerSortDirections, new(uint32))
		atomic.StoreUint32(row.headerSortDirections[i], headerSortNone)
	}
	if len(arr) == 0 {
		for i := 0; i < len(list); i++ {
			row.columnWidths = append(row.columnWidths, float32(MaxFields/len(list)))
		}
	} else {
		if len(arr) != len(list) {
			return nil
		}
		for _, p := range arr {
			row.columnWidths = append(row.columnWidths, p)
		}
	}
	row.ExtendBaseWidget(row)
	return row
}

func NewListBodyRow(id int, list []string,
	handler OnSelectionHandler,
	dclickHandler DoubleClickHandler,
	arr []float32, look *Look) *listRow {
	row := &listRow{
		rowID:              id,
		selected:           new(uint32),
		header:             false,
		actChan:            make(chan *rowAct, rowActChanDepth),
		selectHandler:      handler,
		columnWidths:       []float32{},
		mux:                new(sync.RWMutex),
		look:               look,
		doubleClickHandler: dclickHandler,
	}

	if len(arr) != len(list) {
		return nil
	}
	for i, s := range list {
		row.fieldList = append(row.fieldList, s)
		row.columnWidths = append(row.columnWidths, arr[i])
	}
	for i := 0; i < len(list); i++ {
		row.columnWidths = append(row.columnWidths, float32(MaxFields/len(list)))
	}
	atomic.StoreUint32(row.selected, 0)
	row.ExtendBaseWidget(row)
	return row
}

func (row *listRow) CreateRenderer() fyne.WidgetRenderer {
	row.ExtendBaseWidget(row)
	return newListRowRender(row)
}

func (row *listRow) SetSelection(selected bool, admin bool) {
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

func (row *listRow) SetColumnWidth(col int, delta float32) error {
	if delta == 0 {
		return fmt.Errorf("not dragging")
	}
	row.mux.Lock()
	defer row.mux.Unlock()
	if col < 0 || col >= len(row.columnWidths)-1 {
		if col == len(row.columnWidths)-1 {
			return fmt.Errorf("can't drag last column")
		} else {
			return fmt.Errorf("invalid column id %d", col)
		}
	}
	var previousRightTotal float32 = 0
	for i := col + 1; i < len(row.columnWidths); i++ {
		previousRightTotal += row.columnWidths[i]
	}
	if previousRightTotal == 0 {
		if delta > 0 {
			return fmt.Errorf("reached the end")
		}
		previousRightTotal = 1
	}

	if delta > 0 {
		row.columnWidths[col]++
	} else {
		row.columnWidths[col]--
		if row.columnWidths[col] < 1 {
			row.columnWidths[col] = 1
		}
	}
	var allocated float32 = 0
	for i := 0; i <= col; i++ {
		allocated += row.columnWidths[i]
	}
	available := MaxFields - allocated
	for i := col + 1; i < len(row.columnWidths); i++ {
		row.columnWidths[i] = (available * row.columnWidths[i]) / previousRightTotal
		if row.columnWidths[i] < 1 {
			row.columnWidths[i] = 1
		}
	}
	// log.Printf("now columnWidths is %v", row.columnWidths)
	return nil
}

func (row *listRow) reLayout(refresh bool) {
	select {
	case row.actChan <- &rowAct{act: rowActReLayout, val: refresh}:
		row.Refresh()
	default:
	}
}

func (row *listRow) dragHandler(metadata interface{}, evt *fyne.DragEvent) {
	col := metadata.(listCord).Column
	err := row.SetColumnWidth(col, evt.Dragged.DX)
	if err == nil {
		row.reLayout(true)
		if row.headerDragHandler != nil {
			row.headerDragHandler(row.columnWidths)
		}
	} else {
		// log.Printf("drag skipped,%v", err)
	}
}

func (row *listRow) doubleTapHandler(metadata interface{}, evt *fyne.PointEvent) {
	if row.doubleClickHandler != nil {
		row.doubleClickHandler(metadata.(listCord).Row)
	}
}
func (row *listRow) tapHandler(metadata interface{}, evt *fyne.PointEvent) {
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

type listRowRender struct {
	row              *listRow
	labelList        []*resizableLabel
	background       *canvas.Rectangle
	innerContainer   *fyne.Container
	overallContainer *fyne.Container
	headerSepList    []*widget.Separator
}

func newListRowRender(row *listRow) *listRowRender {
	lrr := &listRowRender{
		row: row,
	}
	lrr.innerContainer = container.NewWithoutLayout()
	for i, s := range row.fieldList {
		options := []ResizableLabelOption{
			WithMetaData(listCord{Row: row.rowID, Column: i}),
			WithTapHandler(row.tapHandler),
		}
		if lrr.row.header {
			options = append(options, WithDraggingHandler(row.dragHandler))
		} else {
			options = append(options, WithDoubleTapHandler(row.doubleTapHandler))
		}
		label := NewResizableLabel(s, row.look, options...)
		lrr.innerContainer.Add(label)
		lrr.labelList = append(lrr.labelList, label)
		// if lrr.row.header {
		if i < len(row.fieldList)-1 {
			// if !lrr.row.header {
			sep := widget.NewSeparator()
			lrr.headerSepList = append(lrr.headerSepList, sep)
			lrr.innerContainer.Add(sep)
			// } else {
			// sep := newHeaderSep(row, i)
			// lrr.headerSepList = append(lrr.headerSepList, sep)
			// lrr.innerContainer.Add(sep)
			// }
			// }
		}
	}
	lrr.background = canvas.NewRectangle(theme.BackgroundColor())
	lrr.overallContainer = container.NewMax(lrr.background, lrr.innerContainer)
	return lrr
}

func (lrr *listRowRender) BackgroundColor() color.Color {
	return color.Transparent
}
func (lrr *listRowRender) Destroy() {

}

const (
	headerSepWidth         = 1
	sepGapWidth    float32 = 1
)

var ColumnWidth = make(map[int]float32) //SetColumnWidth(id int, width float32)

func (lrr *listRowRender) Layout(layoutsize fyne.Size) {
	newsize := fyne.NewSize(layoutsize.Width-lrr.row.rightPadding, layoutsize.Height)
	if newsize.Width <= 0 || newsize.Height <= 0 {
		return
	}
	lrr.background.Resize(newsize)
	lrr.row.mux.RLock()
	defer lrr.row.mux.RUnlock()
	var startx float32 = 0
	for i, label := range lrr.labelList {
		cellWidth := (lrr.row.columnWidths[i] * newsize.Width) / MaxFields
		if w, ok := ColumnWidth[i]; ok {
			if w != 0 {
				cellWidth = w
				//ColumnWidth[i] = 0
			}
			//ColumnWidth[i] = 0
		}
		//if i == 2 {
		//	cellWidth = 400
		//}
		cellWidth -= 20
		label.Resize(fyne.NewSize(cellWidth-2*sepGapWidth-headerSepWidth, newsize.Height))
		label.Move(fyne.NewPos(startx+sepGapWidth, 0))
		label.Refresh()
		startx += cellWidth
		if i < len(lrr.labelList)-1 {
			lrr.headerSepList[i].Resize(fyne.NewSize(headerSepWidth, newsize.Height))
			lrr.headerSepList[i].Move(fyne.NewPos(startx-headerSepWidth-sepGapWidth, 0))
			lrr.headerSepList[i].Refresh()
		}
	}
}

func (lrr *listRowRender) MinSize() fyne.Size {
	h := lrr.labelList[0].MinSize().Height
	var w float32 = 0
	for _, label := range lrr.labelList {
		w += label.MinSize().Width
	}
	return fyne.NewSize(w, h)
	// return lrr.innerContainer.MinSize()
}
func (lrr *listRowRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{lrr.overallContainer}
}

const (
	ascendChar  = rune(0x02c4)
	descendChar = rune(0x02c5)
)

func (lrr *listRowRender) Refresh() {
	select {
	case act := <-lrr.row.actChan:
		switch act.act {
		case rowActSetSelection:
			if atomic.LoadUint32(lrr.row.selected) != 0 {
				lrr.background.FillColor = theme.FocusColor()
			} else {
				lrr.background.FillColor = theme.BackgroundColor()
			}
			lrr.background.Refresh()
			canvas.Refresh(lrr.background)
			canvas.Refresh(lrr.row)
			canvas.Refresh(lrr.background)
		case rowActHeaderSetSort:
			lrr.row.mux.Lock()
			for i := range lrr.row.headerSortDirections {
				newdirection := atomic.LoadUint32(lrr.row.headerSortDirections[i])
				newtext := lrr.labelList[i].Get()
				if newtext[len(newtext)-1] == ascendChar || newtext[len(newtext)-1] == descendChar {
					newtext = newtext[:len(newtext)-1]
				}
				switch newdirection {
				case headerSortAscend:
					newtext = append(newtext, ascendChar)
				case headerSortDescend:
					newtext = append(newtext, descendChar)
				}
				lrr.labelList[i].Set(string(newtext))
			}
			lrr.row.mux.Unlock()
		case rowActReLayout:
			lrr.Layout(lrr.row.Size())

		}
	default:
		return
	}
	canvas.Refresh(lrr.row)
}
