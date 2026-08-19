package breakpoint

import (
	"fmt"
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"

	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/windows"
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
	Type             BreakpointType
	Address          uint64
	Enabled          bool
	SingleShot       bool
	Active           bool
	Name             string
	Module           string
	Slot             uint16
	TypeEx           int
	HwSize           HardwareBreakpointSize
	HitCount         uint32
	FastResume       bool
	Silent           bool
	BreakCondition   string
	LogText          string
	LogCondition     string
	CommandText      string
	CommandCondition string
	OriginalByte     byte
}

type Manager struct {
	breakpoints *safemap.M[uint64, *Breakpoint]
	hardwareBps *safemap.M[int, *Breakpoint]
	memoryBps   *safemap.M[uint64, *Breakpoint]
	nextSlot    uint16
	table       *treetable.TreeTable[Breakpoint]
}

func New() api.Interface {
	m := &Manager{
		breakpoints: safemap.New[uint64, *Breakpoint](),
		hardwareBps: safemap.New[int, *Breakpoint](),
		memoryBps:   safemap.New[uint64, *Breakpoint](),
		nextSlot:    1,
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[Breakpoint]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			bps := m.GetAllBreakpoints()
			for _, bp := range bps {
				m.table.Root().AddChild(m.table.NewNode(*bp))
			}
		},
		JsonName:    "break",
		DisableSort: true,
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Update() error {
	m.table.Root().SetChildren(nil)
	bps := m.GetAllBreakpoints()
	for _, bp := range bps {
		m.table.Root().AddChild(m.table.NewNode(*bp))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Clear() {
	m.breakpoints.Reset()
	m.hardwareBps.Reset()
	m.memoryBps.Reset()
	m.nextSlot = 1
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) SetSoftwareBreakpoint(address uint64, enabled bool) (*Breakpoint, error) {
	if bp, exists := m.breakpoints.Get(address); exists {
		if bp.Type == BreakpointTypeSoftware {
			bp.Enabled = enabled
			return bp, nil
		}
		return nil, fmt.Errorf("breakpoint already exists at address 0x%X with different type", address)
	}

	bp := &Breakpoint{
		Type:         BreakpointTypeSoftware,
		Address:      address,
		Enabled:      enabled,
		SingleShot:   false,
		Active:       false,
		Name:         fmt.Sprintf("BP_0x%X", address),
		Slot:         m.nextSlot,
		HitCount:     0,
		OriginalByte: 0xCC,
	}

	m.breakpoints.Update(address, bp)
	m.nextSlot++

	return bp, nil
}

func (m *Manager) SetBreakpointAtAddress(address uint64) (*Breakpoint, error) {
	return m.SetSoftwareBreakpoint(address, true)
}

func (m *Manager) SetHardwareBreakpoint(address uint64, bpType HardwareBreakpointType, size HardwareBreakpointSize, enabled bool) (*Breakpoint, error) {
	slot := m.findAvailableHardwareSlot()
	if slot == -1 {
		return nil, fmt.Errorf("no available hardware breakpoint slots")
	}

	bp := &Breakpoint{
		Type:     BreakpointTypeHardware,
		Address:  address,
		Enabled:  enabled,
		Slot:     uint16(slot),
		TypeEx:   int(bpType),
		HwSize:   size,
		Name:     fmt.Sprintf("HWBP_0x%X", address),
		HitCount: 0,
	}

	m.breakpoints.Update(address, bp)
	m.hardwareBps.Update(slot, bp)
	m.nextSlot++

	return bp, nil
}

func (m *Manager) SetMemoryBreakpoint(address uint64, size uint64, bpType MemoryBreakpointType, enabled bool) (*Breakpoint, error) {
	for addr := address; addr < address+size; addr++ {
		if _, exists := m.memoryBps.Get(addr); exists {
			return nil, fmt.Errorf("memory breakpoint already exists in range 0x%X-0x%X", address, address+size-1)
		}
	}

	bp := &Breakpoint{
		Type:     BreakpointTypeMemory,
		Address:  address,
		Enabled:  enabled,
		TypeEx:   int(bpType),
		Name:     fmt.Sprintf("MEMBP_0x%X", address),
		HitCount: 0,
	}

	m.breakpoints.Update(address, bp)
	m.memoryBps.Update(address, bp)
	m.nextSlot++

	return bp, nil
}

func (m *Manager) RemoveBreakpoint(address uint64) error {
	bp, exists := m.breakpoints.Get(address)
	if !exists {
		return fmt.Errorf("breakpoint not found at address 0x%X", address)
	}

	switch bp.Type {
	case BreakpointTypeSoftware:
		m.breakpoints.Delete(address)
	case BreakpointTypeHardware:
		m.breakpoints.Delete(address)
		m.hardwareBps.Delete(int(bp.Slot))
	case BreakpointTypeMemory:
		m.breakpoints.Delete(address)
		m.memoryBps.Delete(address)
	default:
		return fmt.Errorf("unsupported breakpoint type")
	}

	return nil
}

func (m *Manager) EnableBreakpoint(address uint64) error {
	bp, exists := m.breakpoints.Get(address)
	if !exists {
		return fmt.Errorf("breakpoint not found at address 0x%X", address)
	}

	bp.Enabled = true
	return nil
}

func (m *Manager) DisableBreakpoint(address uint64) error {
	bp, exists := m.breakpoints.Get(address)
	if !exists {
		return fmt.Errorf("breakpoint not found at address 0x%X", address)
	}

	bp.Enabled = false
	return nil
}

func (m *Manager) FindByAddress(address uint64) *Breakpoint {
	bp, _ := m.breakpoints.Get(address)
	return bp
}

func (m *Manager) FindBySlot(slot uint16) *Breakpoint {
	for _, bp := range m.breakpoints.Range() {
		if bp.Slot == slot {
			return bp
		}
	}
	return nil
}

func (m *Manager) GetAllBreakpoints() []*Breakpoint {
	bps := make([]*Breakpoint, 0)
	for _, bp := range m.breakpoints.Range() {
		bps = append(bps, bp)
	}
	return bps
}

func (m *Manager) GetEnabledBreakpoints() []*Breakpoint {
	bps := make([]*Breakpoint, 0)
	for _, bp := range m.breakpoints.Range() {
		if bp.Enabled {
			bps = append(bps, bp)
		}
	}
	return bps
}

func (m *Manager) ClearAll() {
	m.breakpoints.Reset()
	m.hardwareBps.Reset()
	m.memoryBps.Reset()
	m.nextSlot = 1
}

func (m *Manager) GetHardwareBreakpointForSlot(slot int) *Breakpoint {
	bp, _ := m.hardwareBps.Get(slot)
	return bp
}

func (m *Manager) GetMemoryBreakpoints() []*Breakpoint {
	bps := make([]*Breakpoint, 0)
	for _, bp := range m.memoryBps.Range() {
		bps = append(bps, bp)
	}
	return bps
}

func (m *Manager) findAvailableHardwareSlot() int {
	for i := range 4 {
		if _, exists := m.hardwareBps.Get(i); !exists {
			return i
		}
	}
	return -1
}

func (m *Manager) ApplySoftwareBreakpoint(handle windows.Handle, bp *Breakpoint) error {
	if !bp.Enabled || bp.Type != BreakpointTypeSoftware {
		return nil
	}

	originalBytes, err := windows.ReadProcessMemory(handle, bp.Address, 1)
	if err != nil {
		return fmt.Errorf("failed to read original bytes: %v", err)
	}

	bp.OriginalByte = originalBytes[0]

	int3 := []byte{0xCC}
	_, err = windows.WriteProcessMemory(handle, bp.Address, int3)
	if err != nil {
		return fmt.Errorf("failed to write INT3: %v", err)
	}

	bp.Active = true
	return nil
}

func (m *Manager) RemoveSoftwareBreakpoint(handle windows.Handle, bp *Breakpoint) error {
	if !bp.Active || bp.Type != BreakpointTypeSoftware {
		return nil
	}

	original := []byte{bp.OriginalByte}
	_, err := windows.WriteProcessMemory(handle, bp.Address, original)
	if err != nil {
		return fmt.Errorf("failed to restore original byte: %v", err)
	}

	bp.Active = false
	return nil
}

func (m *Manager) ApplyHardwareBreakpoint(ctx *windows.Context, bp *Breakpoint) error {
	if !bp.Enabled || bp.Type != BreakpointTypeHardware {
		return nil
	}

	slot := int(bp.Slot)

	switch slot {
	case 0:
		ctx.Dr0 = bp.Address
	case 1:
		ctx.Dr1 = bp.Address
	case 2:
		ctx.Dr2 = bp.Address
	case 3:
		ctx.Dr3 = bp.Address
	}

	var sizeBits uint64
	switch bp.HwSize {
	case HardwareBreakpointByte:
		sizeBits = 0
	case HardwareBreakpointWord:
		sizeBits = 1
	case HardwareBreakpointDword:
		sizeBits = 2
	case HardwareBreakpointQword:
		sizeBits = 3
	}

	var typeBits uint64
	switch HardwareBreakpointType(bp.TypeEx) {
	case HardwareBreakpointExecute:
		typeBits = 0
	case HardwareBreakpointWrite:
		typeBits = 1
	case HardwareBreakpointAccess:
		typeBits = 3
	}

	shift := uint64(slot * 4)
	mask := uint64(0xF << shift)
	value := (typeBits | (sizeBits << 2)) << shift

	ctx.Dr7 = (ctx.Dr7 & ^mask) | value
	ctx.Dr7 |= (1 << (slot + 1))

	bp.Active = true
	return nil
}

func (m *Manager) RemoveHardwareBreakpoint(ctx *windows.Context, bp *Breakpoint) error {
	if !bp.Active || bp.Type != BreakpointTypeHardware {
		return nil
	}

	slot := int(bp.Slot)

	switch slot {
	case 0:
		ctx.Dr0 = 0
	case 1:
		ctx.Dr1 = 0
	case 2:
		ctx.Dr2 = 0
	case 3:
		ctx.Dr3 = 0
	}

	ctx.Dr7 &= ^(0xF << (slot * 4))
	ctx.Dr7 &= ^(1 << (slot + 1))

	bp.Active = false
	return nil
}

func (m *Manager) GetBreakpointCount() int {
	count := 0
	for range m.breakpoints.Range() {
		count++
	}
	return count
}

func (m *Manager) GetEnabledBreakpointCount() int {
	count := 0
	for _, bp := range m.breakpoints.Range() {
		if bp.Enabled {
			count++
		}
	}
	return count
}
