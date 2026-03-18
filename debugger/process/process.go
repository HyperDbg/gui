package process

import (
	"fmt"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type ProcessInfo struct {
	PID        uint32
	Name       string
	Path       string
	Modules    []*ModuleInfo
	Threads    []*ThreadInfo
	Handles    map[string]uint64
	MemoryUsed uint64
}

type ModuleInfo struct {
	Name        string
	BaseAddress uint64
	Size        uint64
	Path        string
}

type ThreadInfo struct {
	TID          uint32
	StartAddress uint64
	State        string
}

type Provider struct {
	driver *driver.DriverHandle
	mu     sync.RWMutex
	procs  *safemap.M[uint32, *ProcessInfo]
	table  *treetable.TreeTable[ProcessInfo]
}

func New() api.Interface {
	return &Provider{
		procs: safemap.New[uint32, *ProcessInfo](),
	}
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	procs := m.GetAllProcesses()
	for _, proc := range procs {
		m.table.Root().AddChild(m.table.NewNode(*proc))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.procs.Reset()
}

func (m *Provider) AttachProcess(pid uint32) error {
	return fmt.Errorf("AttachProcess not implemented")
}

func (m *Provider) DetachProcess(pid uint32) error {
	return fmt.Errorf("DetachProcess not implemented")
}

func (m *Provider) EditMemory(pid uint32, address uint64, data []byte) error {
	return fmt.Errorf("EditMemory not implemented")
}

func (m *Provider) SearchMemory(pid uint32, address uint64, size uint32, searchBytes []byte) ([]uint64, error) {
	return nil, fmt.Errorf("SearchMemory not implemented")
}

func (m *Provider) HideToTransparentDebugger(pid uint32) error {
	return fmt.Errorf("HideToTransparentDebugger not implemented")
}

func (m *Provider) UnhideFromTransparentDebugger(pid uint32) error {
	return fmt.Errorf("UnhideFromTransparentDebugger not implemented")
}

func (m *Provider) FlushLoggingBuffers() error {
	return fmt.Errorf("FlushLoggingBuffers not implemented")
}

func (m *Provider) AddProcess(info *ProcessInfo) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.procs.Update(info.PID, info)
}

func (m *Provider) GetProcess(pid uint32) *ProcessInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.procs.GetMust(pid)
}

func (m *Provider) GetAllProcesses() []*ProcessInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.procs.Values()
}

func (m *Provider) RemoveProcess(pid uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.procs.Delete(pid)
}
