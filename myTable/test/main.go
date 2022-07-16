package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/myTable"
	"github.com/ddkwork/librarygo/src/fynelib/fyneTheme"
	"net/http"
	"sort"
	"time"
)

type (
	a interface {
		myTable.Interface
		AddRow(info PacketInfo)
	}
	DecodedInfo struct {
		Head      string
		Pb2       string
		Pb3       string
		Tdf       string
		Taf       string
		Acc       string
		Text      string
		Json      string
		Websocket string
		Msgpack   string
		HexDump   string
	}
	PacketInfo struct {
		PacketIndex int
		Method      string
		Host        string
		Path        string
		ConnectType string
		Size        int
		PadTime     time.Duration
		StartTime   time.Time
		Status      string
		StatusCode  int
		Note        string
		Req         DecodedInfo
		Resp        DecodedInfo
	}
	packets []*PacketInfo
)

func (p *packets) Append(data any) {
	//TODO implement me
	panic("implement me")
}

func main() {
	a := app.NewWithID("com.rows.app")
	a.SetIcon(nil)
	fyneTheme.New().SetDarkTheme(a)
	w := a.NewWindow("app")
	w.Resize(fyne.NewSize(1040, 780))
	w.SetMaster()
	w.CenterOnScreen()
	p := new(packets)
	selectionHandler := myTable.WithSelectionHandler(func(id int, selected bool) {
		println(id)
		println(selected)
	})
	doubleClickHandler := myTable.WithDoubleClickHandler(func(id int) {
		popUpMenu := widget.NewPopUpMenu(
			fyne.NewMenu("pop",
				fyne.NewMenuItem("copy", func() {
					println("copy")
				}),
				fyne.NewMenuItem("cut", func() {
					println("cut")
				}),
				fyne.NewMenuItem("no", func() {
					println("no")
				}),
			),
			nil,
		)
		popUpMenu.Show()
	})
	list, err := myTable.NewDVList(p, selectionHandler, doubleClickHandler)
	if err != nil {
		panic(err.Error())
	}
	go func() {
		myTable.ColumnWidth = map[int]float32{
			2: 300,
		} //todo
		ticker := time.NewTicker(1 * time.Second)
		defer func() { ticker.Stop() }()
		for i := 0; i < 10; i++ {
			now := time.Now()
			p.Add(&PacketInfo{
				PacketIndex: i + 1,
				Method:      http.MethodGet,
				Host:        "www.baidu.com",
				Path:        "/login",
				ConnectType: "json",
				Size:        246,
				PadTime:     now.Sub(time.Now()),
				StartTime:   now,
				Status:      http.StatusText(http.StatusOK),
				StatusCode:  http.StatusOK,
				Note:        "good",
				Req:         DecodedInfo{},
				Resp:        DecodedInfo{},
			})
			p.Add(&PacketInfo{
				PacketIndex: i + 2,
				Method:      http.MethodPost,
				Host:        "www.baidu.com",
				Path:        "/login",
				ConnectType: "json",
				Size:        246,
				PadTime:     now.Sub(time.Now()),
				StartTime:   now,
				Status:      http.StatusText(http.StatusOK),
				StatusCode:  http.StatusOK,
				Note:        "good",
				Req:         DecodedInfo{},
				Resp:        DecodedInfo{},
			})
			list.SetData(p)
			p.Filter(http.MethodPost, 1)
			<-ticker.C
			list.Refresh()
		}
	}()
	w.SetContent(list)
	w.ShowAndRun()
}

func (p *packets) Add(info *PacketInfo) { *p = append(*p, info) }
func (p *packets) Len() int             { return len(*p) }

type (
	FieldName string
)

func (FieldName) PacketIndex() string { return "PacketIndex" }
func (FieldName) Method() string      { return "Method" }
func (FieldName) Host() string        { return "Host" }
func (FieldName) Path() string        { return "Path" }
func (FieldName) ConnectType() string { return "ConnectType" }
func (FieldName) Size() string        { return "Size" }
func (FieldName) PadTime() string     { return "PadTime" }
func (FieldName) StartTime() string   { return "StartTime" }
func (FieldName) Status() string      { return "Status" }
func (FieldName) StatusCode() string  { return "StatusCode" }
func (FieldName) Note() string        { return "Note" }

var fieldName FieldName

func (p packets) Header() []string {
	return []string{
		fieldName.PacketIndex(),
		fieldName.Method(),
		fieldName.Host(),
		fieldName.Path(),
		fieldName.ConnectType(),
		fieldName.Size(),
		fieldName.PadTime(),
		fieldName.StartTime(),
		fieldName.Status(),
		fieldName.StatusCode(),
		fieldName.Note(),
	}
}

func (p packets) Rows(id int) []string {
	if id < 0 || id >= p.Len() {
		return nil
	}
	return []string{
		fmt.Sprintf("%03d", p[id].PacketIndex),
		p[id].Method,
		p[id].Host,
		p[id].Path,
		p[id].ConnectType,
		fmt.Sprint(p[id].Size),
		p[id].PadTime.String(),
		p[id].StartTime.String(),
		p[id].Status,
		fmt.Sprint(p[id].StatusCode),
		p[id].Note,
	}
}

func (p packets) Sort(field int, ascend bool) {
	switch p.Header()[field] {
	case fieldName.PacketIndex():
		sort.Slice(p, func(i, j int) bool {
			return p[i].PacketIndex > p[j].PacketIndex
		})
	case fieldName.Method():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Method > p[j].Method
		})
	case fieldName.Host():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Host > p[j].Host
		})
	case fieldName.Path():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Path > p[j].Path
		})
	case fieldName.ConnectType():
		sort.Slice(p, func(i, j int) bool {
			return p[i].ConnectType > p[j].ConnectType
		})
	case fieldName.Size():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Size > p[j].Size
		})
	case fieldName.PadTime():
		sort.Slice(p, func(i, j int) bool {
			return p[i].PadTime > p[j].PadTime
		})
	case fieldName.StartTime():
		sort.Slice(p, func(i, j int) bool {
			return p[i].StartTime.Unix() > p[j].StartTime.Unix()
		})
	case fieldName.Status():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Status > p[j].Status
		})
	case fieldName.StatusCode():
		sort.Slice(p, func(i, j int) bool {
			return p[i].StatusCode > p[j].StatusCode
		})
	case fieldName.Note():
		sort.Slice(p, func(i, j int) bool {
			return p[i].Note > p[j].Note
		})
	}
}
func (p packets) Filter(kw string, i int) {
	//for lineIndex, s := range p.Rows(i) {
	//	println(s)
	//	if strings.Contains(kw, s) {
	//	}
	//}
	//tocdiag.list.ScrollTo(i)
	//tocdiag.list.SetSelection(i, true
	//save
	//delete
	//update
	//right meanu
	//copy clomn
}
