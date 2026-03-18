package breakpoint

import (
	"fmt"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
)

type BreakpointType int

const (
	BreakpointTypeNone BreakpointType = iota
	BreakpointTypeSoftware
	BreakpointTypeHardware
	BreakpointTypeMemory
	BreakpointTypeDll
	BreakpointTypeException
)

type HardwareBreakpointType int

const (
	HardwareBreakpointAccess HardwareBreakpointType = iota
	HardwareBreakpointWrite
	HardwareBreakpointExecute
)

type HardwareBreakpointSize int

const (
	HardwareBreakpointByte HardwareBreakpointSize = iota
	HardwareBreakpointWord
	HardwareBreakpointDword
	HardwareBreakpointQword
)

type MemoryBreakpointType int

const (
	MemoryBreakpointAccess MemoryBreakpointType = iota
	MemoryBreakpointRead
	MemoryBreakpointWrite
	MemoryBreakpointExecute
)

type Breakpoint struct {
	Type           BreakpointType
	Address        uint64
	Enabled        bool
	SingleShot     bool
	Active         bool
	Name           string
	Module         string
	Slot           uint16
	HwType         HardwareBreakpointType
	HwSize         HardwareBreakpointSize
	HitCount       uint32
	FastResume     bool
	Silent         bool
	BreakCondition string
	LogText        string
	LogCondition   string
	CommandText    string
	OriginalByte   byte
}

type Provider struct {
	mu          sync.RWMutex
	breakpoints map[uint64]*Breakpoint
	hardwareBps map[int]*Breakpoint
	memoryBps   map[uint64]*Breakpoint
	nextSlot    uint16
}

func New() api.Interface {
	return &Provider{
		breakpoints: make(map[uint64]*Breakpoint),
		hardwareBps: make(map[int]*Breakpoint),
		memoryBps:   make(map[uint64]*Breakpoint),
		nextSlot:    1,
	}
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{}
	}
}

func (m *Provider) Update() error {
	return nil
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.breakpoints = make(map[uint64]*Breakpoint)
	m.hardwareBps = make(map[int]*Breakpoint)
	m.memoryBps = make(map[uint64]*Breakpoint)
	m.nextSlot = 1
}

func (m *Provider) AddBreakpoint(bp *Breakpoint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if bp.Type == BreakpointTypeHardware {
		if len(m.hardwareBps) >= 4 {
			return fmt.Errorf("maximum 4 hardware breakpoints allowed")
		}
		slot := m.findAvailableHardwareSlot()
		if slot == -1 {
			return fmt.Errorf("no available hardware breakpoint slot")
		}
		bp.Slot = uint16(slot)
		m.hardwareBps[slot] = bp
	} else if bp.Type == BreakpointTypeMemory {
		m.memoryBps[bp.Address] = bp
	} else {
		m.breakpoints[bp.Address] = bp
	}

	return nil
}

func (m *Provider) RemoveBreakpoint(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.breakpoints, address)
	delete(m.memoryBps, address)
}

func (m *Provider) GetBreakpoint(address uint64) *Breakpoint {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.breakpoints[address]
}

func (m *Provider) GetAllBreakpoints() []*Breakpoint {
	m.mu.RLock()
	defer m.mu.RUnlock()

	bps := make([]*Breakpoint, 0, len(m.breakpoints))
	for _, bp := range m.breakpoints {
		bps = append(bps, bp)
	}
	return bps
}

func (m *Provider) EnableBreakpoint(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if bp, ok := m.breakpoints[address]; ok {
		bp.Enabled = true
	}
}

func (m *Provider) DisableBreakpoint(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if bp, ok := m.breakpoints[address]; ok {
		bp.Enabled = false
	}
}

func (m *Provider) IncrementHitCount(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if bp, ok := m.breakpoints[address]; ok {
		bp.HitCount++
	}
}

func (m *Provider) findAvailableHardwareSlot() int {
	for i := range 4 {
		if _, ok := m.hardwareBps[i]; !ok {
			return i
		}
	}
	return -1
}
