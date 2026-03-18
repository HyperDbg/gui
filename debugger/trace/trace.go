package trace

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type TraceEntry struct {
	Index       uint64
	Address     uint64
	ThreadId    uint32
	Timestamp   int64
	Registers   map[string]uint64
	Instruction string
	Comment     string
}

type TraceRecord struct {
	Entries    []*TraceEntry
	StartIndex uint64
	EndIndex   uint64
	IsActive   bool
	ProcessId  uint32
	ThreadId   uint32
}

type TraceCondition struct {
	Expression string
	Enabled    bool
}

type Provider struct {
	mu          sync.RWMutex
	records     *safemap.M[uint64, *TraceRecord]
	conditions  []*TraceCondition
	activeTrace *TraceRecord
	nextIndex   uint64
	table       *treetable.TreeTable[TraceEntry]
}

func New() api.Interface {
	m := &Provider{
		records:    safemap.New[uint64, *TraceRecord](),
		conditions: make([]*TraceCondition, 0),
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = treetable.NewTreeTable[TraceEntry]()
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	if m.activeTrace != nil {
		for _, entry := range m.activeTrace.Entries {
			m.table.Root().AddChild(m.table.NewNode(*entry))
		}
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) StartTrace(processId, threadId uint32) *TraceRecord {
	m.mu.Lock()
	defer m.mu.Unlock()

	record := &TraceRecord{
		Entries:    make([]*TraceEntry, 0),
		StartIndex: m.nextIndex,
		IsActive:   true,
		ProcessId:  processId,
		ThreadId:   threadId,
	}

	m.records.Update(m.nextIndex, record)
	m.activeTrace = record
	return record
}

func (m *Provider) StopTrace() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.activeTrace != nil {
		m.activeTrace.IsActive = false
		m.activeTrace.EndIndex = m.nextIndex
		m.activeTrace = nil
	}
}

func (m *Provider) AddTraceEntry(entry *TraceEntry) {
	m.mu.Lock()
	defer m.mu.Unlock()

	entry.Index = m.nextIndex
	m.nextIndex++

	if m.activeTrace != nil {
		m.activeTrace.Entries = append(m.activeTrace.Entries, entry)
	}
}

func (m *Provider) GetTrace(index uint64) *TraceRecord {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.records.GetMust(index)
}

func (m *Provider) GetActiveTrace() *TraceRecord {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.activeTrace
}

func (m *Provider) GetAllTraces() []*TraceRecord {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.records.Values()
}

func (m *Provider) AddCondition(condition *TraceCondition) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.conditions = append(m.conditions, condition)
}

func (m *Provider) RemoveCondition(index int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if index >= 0 && index < len(m.conditions) {
		m.conditions = append(m.conditions[:index], m.conditions[index+1:]...)
	}
}

func (m *Provider) GetConditions() []*TraceCondition {
	m.mu.RLock()
	defer m.mu.RUnlock()

	conditions := make([]*TraceCondition, len(m.conditions))
	copy(conditions, m.conditions)
	return conditions
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.records.Reset()
	m.conditions = make([]*TraceCondition, 0)
	m.activeTrace = nil
	m.nextIndex = 0
}
