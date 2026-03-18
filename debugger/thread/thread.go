package thread

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type ThreadState int

const (
	ThreadStateRunning ThreadState = iota
	ThreadStateSuspended
	ThreadStateTerminated
	ThreadStateWaiting
)

type Thread struct {
	ThreadId     uint32
	ProcessId    uint32
	StartAddress uint64
	State        ThreadState
	Priority     int32
	Context      *ThreadContext
	IsMainThread bool
}

type ThreadContext struct {
	EAX    uint64
	EBX    uint64
	ECX    uint64
	EDX    uint64
	ESI    uint64
	EDI    uint64
	EBP    uint64
	ESP    uint64
	EIP    uint64
	EFlags uint32
}

type Provider struct {
	mu      sync.RWMutex
	threads *safemap.M[uint32, *Thread]
	current *Thread
	table   *treetable.TreeTable[Thread]
}

func New() api.Interface {
	m := &Provider{
		threads: safemap.New[uint32, *Thread](),
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = treetable.NewTreeTable[Thread]()
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	threads := m.GetAllThreads()
	for _, thread := range threads {
		m.table.Root().AddChild(m.table.NewNode(*thread))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) AddThread(thread *Thread) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.threads.Update(thread.ThreadId, thread)
}

func (m *Provider) RemoveThread(threadId uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.threads.Delete(threadId)
}

func (m *Provider) GetThread(threadId uint32) *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.threads.GetMust(threadId)
}

func (m *Provider) GetAllThreads() []*Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.threads.Values()
}

func (m *Provider) SetCurrentThread(threadId uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.current = m.threads.GetMust(threadId)
}

func (m *Provider) GetCurrentThread() *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.current
}

func (m *Provider) GetMainThread() *Thread {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, thread := range m.threads.Values() {
		if thread.IsMainThread {
			return thread
		}
	}
	return nil
}

func (m *Provider) UpdateThreadState(threadId uint32, state ThreadState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if thread := m.threads.GetMust(threadId); thread != nil {
		thread.State = state
	}
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.threads.Reset()
	m.current = nil
}
