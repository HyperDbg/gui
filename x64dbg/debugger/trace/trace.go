package trace

import (
	"fmt"
	"iter"
	"strings"
	"sync"
	"time"

	"gioui.org/layout"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type TraceEntry struct {
	Timestamp   time.Time
	ThreadId    uint32
	Address     uint64
	Instruction string
	Registers   map[string]uint64
	Notes       string
}

type TraceType int

const (
	TraceTypeInstruction TraceType = iota
	TraceTypeFunctionCall
	TraceTypeFunctionReturn
	TraceTypeException
)

type Manager struct {
	entries []*TraceEntry
	mu      sync.RWMutex
	enabled bool
	maxSize int
	table   *treetable.TreeTable[TraceEntry]
}

func New() api.Interface {
	m := &Manager{
		entries: make([]*TraceEntry, 0),
		enabled: false,
		maxSize: 10000,
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[TraceEntry]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			entries := m.GetAllEntries()
			for _, entry := range entries {
				m.table.Root().AddChild(m.table.NewNode(*entry))
			}
		},
		JsonName: "trace",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Update() error {
	m.table.Root().SetChildren(nil)
	entries := m.GetAllEntries()
	for _, entry := range entries {
		m.table.Root().AddChild(m.table.NewNode(*entry))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) Enable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = true
}

func (m *Manager) Disable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = false
}

func (m *Manager) IsEnabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.enabled
}

func (m *Manager) AddEntry(threadId uint32, address uint64, instruction string, registers map[string]uint64, notes string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.enabled {
		return
	}

	entry := &TraceEntry{
		Timestamp:   time.Now(),
		ThreadId:    threadId,
		Address:     address,
		Instruction: instruction,
		Registers:   registers,
		Notes:       notes,
	}

	m.entries = append(m.entries, entry)

	if len(m.entries) > m.maxSize {
		m.entries = m.entries[1:]
	}
}

func (m *Manager) GetAllEntries() []*TraceEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*TraceEntry, len(m.entries))
	copy(result, m.entries)
	return result
}

func (m *Manager) GetEntriesForThread(threadId uint32) []*TraceEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*TraceEntry, 0)
	for _, entry := range m.entries {
		if entry.ThreadId == threadId {
			result = append(result, entry)
		}
	}
	return result
}

func (m *Manager) GetEntriesInRange(start time.Time, end time.Time) []*TraceEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*TraceEntry, 0)
	for _, entry := range m.entries {
		if entry.Timestamp.After(start) && entry.Timestamp.Before(end) {
			result = append(result, entry)
		}
	}
	return result
}

func (m *Manager) GetEntryCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.entries)
}

func (m *Manager) SetMaxSize(size int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.maxSize = size

	if len(m.entries) > m.maxSize {
		m.entries = m.entries[len(m.entries)-m.maxSize:]
	}
}

func (m *Manager) ExportToText() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result strings.Builder
	for _, entry := range m.entries {
		result.WriteString(fmt.Sprintf("[%s] Thread:%d Address:0x%X %s %s\n",
			entry.Timestamp.Format("15:04:05.000"),
			entry.ThreadId,
			entry.Address,
			entry.Instruction,
			entry.Notes))
	}
	return result.String()
}
