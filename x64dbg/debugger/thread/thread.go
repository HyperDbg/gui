package thread

import (
	"fmt"
	"iter"
	"sync"
	"time"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

type ThreadPriority int

const (
	PriorityIdle         ThreadPriority = -15
	PriorityBelowNormal  ThreadPriority = -1
	PriorityNormal       ThreadPriority = 0
	PriorityAboveNormal  ThreadPriority = 1
	PriorityHighest      ThreadPriority = 2
	PriorityTimeCritical ThreadPriority = 15
	PriorityUnknown      ThreadPriority = 0x7FFFFFFF
)

type ThreadWaitReason int

const (
	WaitReasonExecutive      ThreadWaitReason = 0
	WaitReasonFreePage       ThreadWaitReason = 1
	WaitReasonPageIn         ThreadWaitReason = 2
	WaitReasonPoolAllocation ThreadWaitReason = 3
	WaitReasonDelayExecution ThreadWaitReason = 4
	WaitReasonSuspended      ThreadWaitReason = 5
	WaitReasonUserRequest    ThreadWaitReason = 6
	WaitReasonExecutiveWrap  ThreadWaitReason = 7
	WaitReasonFreePageWrap   ThreadWaitReason = 8
	WaitReasonPageInWrap     ThreadWaitReason = 9
	WaitReasonPoolAllocWrap  ThreadWaitReason = 10
	WaitReasonDelayExecWrap  ThreadWaitReason = 11
	WaitReasonSuspendedWrap  ThreadWaitReason = 12
	WaitReasonUserReqWrap    ThreadWaitReason = 13
	WaitReasonEventPair      ThreadWaitReason = 14
	WaitReasonQueue          ThreadWaitReason = 15
	WaitReasonLpcReceive     ThreadWaitReason = 16
	WaitReasonLpcReply       ThreadWaitReason = 17
	WaitReasonVirtualMemory  ThreadWaitReason = 18
	WaitReasonPageOut        ThreadWaitReason = 19
	WaitReasonRendezvous     ThreadWaitReason = 20
)

type Thread struct {
	IndexName        string
	Id               uint32
	Handle           windows.Handle
	Entry            uint64
	Teb              uint64
	Rip              uint64
	PendingCount     int32
	Priority         ThreadPriority
	WaitForTheReason ThreadWaitReason
	LastError        uint32
	UserTime         time.Time
	KernelTime       time.Time
	CreatTime        time.Time
	CPUCycles        uint64
	Suspended        bool
}

type Manager struct {
	threads       *safemap.M[uint32, *Thread]
	threadHandles *safemap.M[windows.Handle, *Thread]
	mu            sync.RWMutex
	table         *treetable.TreeTable[Thread]
}

func New() api.Interface {
	m := &Manager{
		threads:       safemap.New[uint32, *Thread](),
		threadHandles: safemap.New[windows.Handle, *Thread](),
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[Thread]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			threads := m.GetAllThreads()
			for _, th := range threads {
				m.table.Root().AddChild(m.table.NewNode(*th))
			}
		},
		JsonName: "Thread",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.threads.Reset()
	m.threadHandles.Reset()
}

func (m *Manager) Update() error {
	m.table.Root().SetChildren(nil)
	threads := m.GetAllThreads()
	for _, th := range threads {
		m.table.Root().AddChild(m.table.NewNode(*th))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) AddThread(threadId uint32, handle windows.Handle, teb uint64) *Thread {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread := &Thread{
		Id:        threadId,
		Handle:    handle,
		Teb:       teb,
		IndexName: fmt.Sprintf("Thread_%d", threadId),
		Priority:  PriorityNormal,
		CreatTime: time.Now(),
		Suspended: false,
	}

	m.threads.Update(threadId, thread)
	m.threadHandles.Update(handle, thread)

	return thread
}

func (m *Manager) SetThreadHandle(threadId uint32, handle windows.Handle) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if thread, exists := m.threads.Get(threadId); exists {
		thread.Handle = handle
		m.threadHandles.Update(handle, thread)
	}
}

func (m *Manager) RemoveThread(threadId uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if thread, exists := m.threads.Get(threadId); exists {
		m.threadHandles.Delete(thread.Handle)
		m.threads.Delete(threadId)
	}
}

func (m *Manager) GetThread(threadId uint32) *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()

	thread, _ := m.threads.Get(threadId)
	return thread
}

func (m *Manager) GetThreadByHandle(handle windows.Handle) *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()

	thread, _ := m.threadHandles.Get(handle)
	return thread
}

func (m *Manager) GetAllThreads() []*Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()

	threads := make([]*Thread, 0)
	for _, thread := range m.threads.Range() {
		threads = append(threads, thread)
	}
	return threads
}

func (m *Manager) GetMainThread() *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, thread := range m.threads.Range() {
		if thread.Id == 0 {
			return thread
		}
	}

	count := 0
	for range m.threads.Range() {
		count++
	}
	if count > 0 {
		for _, thread := range m.threads.Range() {
			return thread
		}
	}

	return nil
}

func (m *Manager) SuspendThread(threadId uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	if thread.Suspended {
		return nil
	}

	count, err := windows.SuspendThread(thread.Handle)
	if err != nil {
		return err
	}

	thread.PendingCount = int32(count)
	thread.Suspended = true

	return nil
}

func (m *Manager) ResumeThread(threadId uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	if !thread.Suspended {
		return nil
	}

	count, err := windows.ResumeThread(thread.Handle)
	if err != nil {
		return err
	}

	thread.PendingCount = int32(count)
	thread.Suspended = false

	return nil
}

func (m *Manager) SuspendAllThreads() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, thread := range m.threads.Range() {
		if !thread.Suspended {
			count, err := windows.SuspendThread(thread.Handle)
			if err != nil {
				return err
			}
			thread.PendingCount = int32(count)
			thread.Suspended = true
		}
	}

	return nil
}

func (m *Manager) ResumeAllThreads() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, thread := range m.threads.Range() {
		if thread.Suspended {
			count, err := windows.ResumeThread(thread.Handle)
			if err != nil {
				return err
			}
			thread.PendingCount = int32(count)
			thread.Suspended = false
		}
	}

	return nil
}

func (m *Manager) UpdateThreadContext(threadId uint32, rip uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	thread.Rip = rip
	return nil
}

func (m *Manager) SetThreadPriority(threadId uint32, priority ThreadPriority) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	thread.Priority = priority
	return nil
}

func (m *Manager) GetThreadCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for range m.threads.Range() {
		count++
	}
	return count
}

func (m *Manager) GetSuspendedThreadCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, thread := range m.threads.Range() {
		if thread.Suspended {
			count++
		}
	}
	return count
}

func (m *Manager) UpdateThreadInfo(threadId uint32, entry uint64, waitReason ThreadWaitReason, lastError uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	thread.Entry = entry
	thread.WaitForTheReason = waitReason
	thread.LastError = lastError

	return nil
}

func (m *Manager) SetThreadName(threadId uint32, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	thread.IndexName = name
	return nil
}

func (m *Manager) GetThreadName(threadId uint32) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return ""
	}

	return thread.IndexName
}

func (m *Manager) CloseThreadHandle(threadId uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	thread, exists := m.threads.Get(threadId)
	if !exists {
		return fmt.Errorf("thread not found: %d", threadId)
	}

	err := windows.CloseHandle(thread.Handle)
	if err != nil {
		return err
	}

	m.threadHandles.Delete(thread.Handle)
	thread.Handle = 0

	return nil
}

func (m *Manager) CloseAllThreadHandles() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, thread := range m.threads.Range() {
		if thread.Handle != 0 {
			err := windows.CloseHandle(thread.Handle)
			if err != nil {
				return err
			}
			m.threadHandles.Delete(thread.Handle)
			thread.Handle = 0
		}
	}

	return nil
}
