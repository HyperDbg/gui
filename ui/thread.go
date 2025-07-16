package ui

import (
	"iter"
	"time"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func NewThread() ux.Widget {
	t := ux.NewTreeTable(Thread{})
	t.TableContext = ux.TableContext[Thread]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Thread]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Thread]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Thread], rows []ux.CellData) Thread {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: n.Data.IndexName},
			//	{Text: fmt.Sprintf("%d", n.Data.Id)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Entry)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Teb)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Rip)},
			//	{Text: fmt.Sprintf("%d", n.Data.PendingCount)},
			//	{Text: n.Data.Priority},
			//	{Text: n.Data.WaitForTheReason},
			//	{Text: n.Data.LastError},
			//	{Text: n.Data.UserTime.Format("2006-01-02 15:04:05")},
			//	{Text: n.Data.KernelTime.Format("2006-01-02 15:04:05")},
			//	{Text: n.Data.CreatTime.Format("2006-01-02 15:04:05")},
			//	{Text: fmt.Sprintf("%016X", n.Data.CPUCycles)},
			// }
			return ux.UnmarshalRow[Thread](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
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
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "Thread",
		IsDocument: false,
	}
	return t
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
