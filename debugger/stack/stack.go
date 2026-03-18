package stack

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type StackFrame struct {
	Address    uint64
	ReturnAddr uint64
	Function   string
	Module     string
	FileName   string
	LineNumber int
	FrameSize  uint32
	Params     []StackParam
	Locals     []StackLocal
}

type StackParam struct {
	Name  string
	Type  string
	Value uint64
}

type StackLocal struct {
	Name  string
	Type  string
	Value uint64
}

type StackEntry struct {
	Address uint64
	Value   uint64
	Comment string
}

type Provider struct {
	mu      sync.RWMutex
	frames  *safemap.M[uint64, *StackFrame]
	entries map[uint64]*StackEntry
	table   *treetable.TreeTable[StackFrame]
}

func New() api.Interface {
	m := &Provider{
		frames:  safemap.New[uint64, *StackFrame](),
		entries: make(map[uint64]*StackEntry),
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = treetable.NewTreeTable[StackFrame]()
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	frames := m.GetFrames()
	for _, frame := range frames {
		m.table.Root().AddChild(m.table.NewNode(*frame))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) AddFrame(frame *StackFrame) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.frames.Update(frame.Address, frame)
}

func (m *Provider) GetFrames() []*StackFrame {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.frames.Values()
}

func (m *Provider) ClearFrames() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.frames.Reset()
}

func (m *Provider) AddEntry(entry *StackEntry) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.entries[entry.Address] = entry
}

func (m *Provider) GetEntry(address uint64) *StackEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.entries[address]
}

func (m *Provider) GetEntries() []*StackEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entries := make([]*StackEntry, 0, len(m.entries))
	for _, entry := range m.entries {
		entries = append(entries, entry)
	}
	return entries
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.frames.Reset()
	m.entries = make(map[uint64]*StackEntry)
}
