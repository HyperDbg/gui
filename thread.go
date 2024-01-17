package main

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/states"
	"time"
)

func pageThread(parent *gi.Frame) {
	ThreadTable(parent)
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

func ThreadTable(frame *gi.Frame) *giv.TableView {
	threads := make([]*Thread, 100)
	for i := range threads {
		ts := &Thread{
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
		threads[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&threads)
	return tv
}
