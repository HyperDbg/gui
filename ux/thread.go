package ux

import (
	"fmt"
	"time"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutThread() unison.Paneler {
	table, header := widget.NewTable(Thread{}, widget.TableContext[Thread]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Thread]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum()
			}
			return []widget.CellData{
				{Text: node.Data.IndexName},
				{Text: fmt.Sprintf("%d", node.Data.Id)},
				{Text: fmt.Sprintf("%016X", node.Data.Entry)},
				{Text: fmt.Sprintf("%016X", node.Data.Teb)},
				{Text: fmt.Sprintf("%016X", node.Data.Rip)},
				{Text: fmt.Sprintf("%d", node.Data.PendingCount)},
				{Text: node.Data.Priority},
				{Text: node.Data.WaitForTheReason},
				{Text: node.Data.LastError},
				{Text: node.Data.UserTime.Format("2006-01-02 15:04:05")},
				{Text: node.Data.KernelTime.Format("2006-01-02 15:04:05")},
				{Text: node.Data.CreatTime.Format("2006-01-02 15:04:05")},
				{Text: fmt.Sprintf("%016X", node.Data.CPUCycles)},
			}
		},
		UnmarshalRow: func(node *widget.Node[Thread], values []string) {
			mylog.Todo("UnmarshalRow")
		},
		SelectionChangedCallback: func(root *widget.Node[Thread]) {
			mylog.Todo("SelectionChangedCallback")
		},
		SetRootRowsCallBack: func(root *widget.Node[Thread]) {
			for range 100 {
				ts := Thread{
					IndexName:        "主线程",
					Id:               12364,
					Entry:            0x00007FF7967E55C8,
					Teb:              0x000000F4B0A8A000,
					Rip:              0x00007FF884FAB785,
					PendingCount:     1,
					Priority:         "标准",
					WaitForTheReason: "Executive",
					LastError:        "",
					UserTime:         time.Now(),
					KernelTime:       time.Now(),
					CreatTime:        time.Now(),
					CPUCycles:        0x11A39874,
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "thread",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
}

type Thread struct {
	IndexName        string
	Id               int
	Entry            int
	Teb              int
	Rip              int
	PendingCount     int
	Priority         string
	WaitForTheReason string
	LastError        string
	UserTime         time.Time
	KernelTime       time.Time
	CreatTime        time.Time
	CPUCycles        int
}
