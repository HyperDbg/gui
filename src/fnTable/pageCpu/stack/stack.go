package stack

import (
	"fyne.io/fyne/v2"
	"github.com/ddkwork/hyperdbgui/dvlist"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/librarygo/src/mycheck"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		dvlist.Data
	}
	ds struct {
		hite    string
		address string
		data    string
		dism    string
		notes   string
	}
	object struct {
		ds
		dss []ds
	}
)

func New() Interface {
	return &object{
		ds:  ds{},
		dss: make([]ds, 0),
	}
}

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject { //5,还差一个立即数显示的表格，传递全局变量改变宽度
	dvlist.ColumnWidth = map[int]float32{
		0: 300,
	} //todo
	o.dss = append(o.dss, ds{
		hite:    "hite",
		address: "address",
		data:    "data",
		dism:    "diasm",
		notes:   "notes",
	})
	o.dss = append(o.dss, ds{
		hite:    "hite1",
		address: "address1",
		data:    "data1",
		dism:    "diasm1",
		notes:   "notes1",
	})
	o.dss = append(o.dss, ds{
		hite:    "hite1",
		address: "address1",
		data:    "data1",
		dism:    "diasm1",
		notes:   "notes1",
	})
	list, err := dvlist.NewDVList(o, dvlist.WithSelectionHandler(func(id int, selected bool) {
		println(id)
	}))
	//list.SetData()
	if !mycheck.Error(err) {
		return nil
	}
	return list
}

func (o *object) Len() int {
	return len(o.dss)
}

func (o *object) Fields() []string {
	return []string{
		o.dss[0].hite,
		o.dss[0].address,
		o.dss[0].data,
		o.dss[0].dism,
		o.dss[0].notes,
	}
}

func (o *object) Item(id int) []string {
	return []string{
		o.dss[id].hite,
		o.dss[id].address,
		o.dss[id].data,
		o.dss[id].dism,
		o.dss[id].notes,
	}
}

func (o *object) Sort(field int, ascend bool) {
}

func (o *object) Filter(kw string, i int) {
}
