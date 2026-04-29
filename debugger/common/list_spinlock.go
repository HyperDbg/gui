package common

import (
	"runtime"
	"sync/atomic"
)

var privilegesAlreadyAdjusted atomic.Bool

func SpinlockTryLock(lock *atomic.Int32) bool {
	unlocked := lock.CompareAndSwap(0, 1)
	return unlocked
}

func SpinlockLock(lock *atomic.Int32) {
	wait := uint32(1)
	for !SpinlockTryLock(lock) {
		for i := uint32(0); i < wait; i++ {
			runtime.Gosched()
		}
		if wait < 65536 {
			wait *= 2
		}
	}
}

func SpinlockUnlock(lock *atomic.Int32) {
	lock.Store(0)
}

type ListEntry struct {
	Flink *ListEntry
	Blink *ListEntry
}

func InitializeListHead(listHead *ListEntry) {
	listHead.Flink = listHead
	listHead.Blink = listHead
}

func InsertHeadList(listHead, entry *ListEntry) {
	flink := listHead.Flink
	entry.Flink = flink
	entry.Blink = listHead
	flink.Blink = entry
	listHead.Flink = entry
}

func InsertTailList(listHead, entry *ListEntry) {
	blink := listHead.Blink
	entry.Flink = listHead
	entry.Blink = blink
	blink.Flink = entry
	listHead.Blink = entry
}

func RemoveEntryList(entry *ListEntry) bool {
	prevEntry := entry.Blink
	nextEntry := entry.Flink
	nextEntry.Blink = prevEntry
	prevEntry.Flink = nextEntry
	return nextEntry != entry && prevEntry != entry
}

func IsListEmpty(listHead *ListEntry) bool {
	return listHead.Flink == listHead
}

func PopEntryList(listHead *ListEntry) *ListEntry {
	entry := listHead.Blink
	if entry == listHead {
		return nil
	}
	RemoveEntryList(entry)
	return entry
}
