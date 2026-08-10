package memory

import (
	"encoding/hex"
	"fmt"
	"iter"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"

	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

type MemoryRegion struct {
	BaseAddress       uint64
	AllocationBase    uint64
	AllocationProtect uint32
	RegionSize        uint64
	State             uint32
	Protect           uint32
	Type              uint32
	Info              string
}

type Manager struct {
	mu     sync.Mutex
	handle windows.Handle
	cache  *safemap.M[uint64, *MemoryRegion]
	table  *treetable.TreeTable[MemoryRegion]
}

func New() api.Interface {
	m := &Manager{
		cache: safemap.New[uint64, *MemoryRegion](),
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[MemoryRegion]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			regions, err := m.GetMemoryMap()
			if err == nil {
				for _, region := range regions {
					m.table.Root().AddChild(m.table.NewNode(*region))
				}
			}
		},
		JsonName:    "memory",
		DisableSort: true,
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Update() error {
	m.table.Root().SetChildren(nil)
	regions, err := m.GetMemoryMap()
	if err == nil {
		for _, region := range regions {
			m.table.Root().AddChild(m.table.NewNode(*region))
		}
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Clear() {
	m.cache.Reset()
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) SetHandle(handle windows.Handle) {
	m.mu.Lock()
	m.handle = handle
	m.mu.Unlock()
}

func (m *Manager) GetHandle() windows.Handle {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.handle
}

func (m *Manager) ReadMemory(address uint64, size uint32) ([]byte, error) {
	m.mu.Lock()
	handle := m.handle
	m.mu.Unlock()

	if handle == 0 {
		return nil, fmt.Errorf("process handle not set")
	}

	return m.readMemorySafe(handle, address, size)
}

func (m *Manager) readMemorySafe(handle windows.Handle, address uint64, size uint32) ([]byte, error) {
	if handle == 0 {
		return nil, fmt.Errorf("process handle not set")
	}

	data, err := windows.ReadProcessMemory(handle, address, size)
	if err == nil {
		return data, nil
	}

	mbi, err := windows.VirtualQueryEx(handle, address)
	if err != nil {
		return nil, fmt.Errorf("VirtualQueryEx failed: %w", err)
	}

	oldProtect := mbi.Protect
	err = windows.VirtualProtectEx(handle, address, uint64(size), windows.PAGE_EXECUTE_READ, &oldProtect)
	if err != nil {
		return nil, fmt.Errorf("VirtualProtectEx failed: %w", err)
	}

	data, err = windows.ReadProcessMemory(handle, address, size)
	if err != nil {
		windows.VirtualProtectEx(handle, address, uint64(size), oldProtect, &oldProtect)
		return nil, fmt.Errorf("ReadProcessMemory after VirtualProtectEx failed: %w", err)
	}

	windows.VirtualProtectEx(handle, address, uint64(size), oldProtect, &oldProtect)
	return data, nil
}

func (m *Manager) WriteMemory(address uint64, data []byte) (uint32, error) {
	m.mu.Lock()
	handle := m.handle
	m.mu.Unlock()

	if handle == 0 {
		return 0, fmt.Errorf("process handle not set")
	}

	return windows.WriteProcessMemory(handle, address, data)
}

func (m *Manager) ReadMemoryString(address uint64, maxLen uint32) (string, error) {
	data, err := m.ReadMemory(address, maxLen)
	if err != nil {
		return "", err
	}

	nullIndex := len(data)
	for i, b := range data {
		if b == 0 {
			nullIndex = i
			break
		}
	}

	return string(data[:nullIndex]), nil
}

func (m *Manager) ReadMemoryWString(address uint64, maxLen uint32) (string, error) {
	data, err := m.ReadMemory(address, maxLen*2)
	if err != nil {
		return "", err
	}

	nullIndex := len(data)
	for i := 0; i < len(data)-1; i += 2 {
		if data[i] == 0 && data[i+1] == 0 {
			nullIndex = i
			break
		}
	}

	return string(data[:nullIndex]), nil
}

func (m *Manager) ReadMemoryUint8(address uint64) (uint8, error) {
	data, err := m.ReadMemory(address, 1)
	if err != nil {
		return 0, err
	}
	return data[0], nil
}

func (m *Manager) ReadMemoryUint16(address uint64) (uint16, error) {
	data, err := m.ReadMemory(address, 2)
	if err != nil {
		return 0, err
	}
	return uint16(data[0]) | uint16(data[1])<<8, nil
}

func (m *Manager) ReadMemoryUint32(address uint64) (uint32, error) {
	data, err := m.ReadMemory(address, 4)
	if err != nil {
		return 0, err
	}
	return uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24, nil
}

func (m *Manager) ReadMemoryUint64(address uint64) (uint64, error) {
	data, err := m.ReadMemory(address, 8)
	if err != nil {
		return 0, err
	}
	return uint64(data[0]) | uint64(data[1])<<8 | uint64(data[2])<<16 | uint64(data[3])<<24 |
		uint64(data[4])<<32 | uint64(data[5])<<40 | uint64(data[6])<<48 | uint64(data[7])<<56, nil
}

func (m *Manager) ReadMemoryInt8(address uint64) (int8, error) {
	val, err := m.ReadMemoryUint8(address)
	return int8(val), err
}

func (m *Manager) ReadMemoryInt16(address uint64) (int16, error) {
	val, err := m.ReadMemoryUint16(address)
	return int16(val), err
}

func (m *Manager) ReadMemoryInt32(address uint64) (int32, error) {
	val, err := m.ReadMemoryUint32(address)
	return int32(val), err
}

func (m *Manager) ReadMemoryInt64(address uint64) (int64, error) {
	val, err := m.ReadMemoryUint64(address)
	return int64(val), err
}

func (m *Manager) WriteMemoryUint8(address uint64, value uint8) (uint32, error) {
	return m.WriteMemory(address, []byte{value})
}

func (m *Manager) WriteMemoryUint16(address uint64, value uint16) (uint32, error) {
	data := []byte{byte(value), byte(value >> 8)}
	return m.WriteMemory(address, data)
}

func (m *Manager) WriteMemoryUint32(address uint64, value uint32) (uint32, error) {
	data := []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)}
	return m.WriteMemory(address, data)
}

func (m *Manager) WriteMemoryUint64(address uint64, value uint64) (uint32, error) {
	data := []byte{
		byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24),
		byte(value >> 32), byte(value >> 40), byte(value >> 48), byte(value >> 56),
	}
	return m.WriteMemory(address, data)
}

func (m *Manager) QueryMemory(address uint64) (*MemoryRegion, error) {
	m.mu.Lock()
	handle := m.handle
	m.mu.Unlock()

	if handle == 0 {
		return nil, fmt.Errorf("process handle not set")
	}

	mbi, err := windows.VirtualQueryEx(handle, address)
	if err != nil {
		return nil, err
	}

	region := &MemoryRegion{
		BaseAddress:       mbi.BaseAddress,
		AllocationBase:    mbi.AllocationBase,
		AllocationProtect: mbi.AllocationProtect,
		RegionSize:        mbi.RegionSize,
		State:             mbi.State,
		Protect:           mbi.Protect,
		Type:              mbi.Type,
		Info:              m.formatProtectionString(mbi.Protect),
	}

	m.cache.Update(address, region)
	return region, nil
}

func (m *Manager) GetMemoryMap() ([]*MemoryRegion, error) {
	m.mu.Lock()
	handle := m.handle
	m.mu.Unlock()

	if handle == 0 {
		return nil, fmt.Errorf("process handle not set")
	}

	regions := make([]*MemoryRegion, 0)
	currentAddress := uint64(0)

	for {
		region, err := m.QueryMemory(currentAddress)
		if err != nil {
			break
		}

		regions = append(regions, region)
		currentAddress = region.BaseAddress + region.RegionSize

		if currentAddress < region.BaseAddress {
			break
		}
	}

	return regions, nil
}

func (m *Manager) SearchMemory(pattern []byte, startAddress uint64, endAddress uint64) ([]uint64, error) {
	m.mu.Lock()
	handle := m.handle
	m.mu.Unlock()

	if handle == 0 {
		return nil, fmt.Errorf("process handle not set")
	}

	results := make([]uint64, 0)
	currentAddress := startAddress

	for currentAddress < endAddress {
		region, err := m.QueryMemory(currentAddress)
		if err != nil {
			break
		}

		if region.State != 0x1000 {
			currentAddress = region.BaseAddress + region.RegionSize
			continue
		}

		regionEnd := min(region.BaseAddress+region.RegionSize, endAddress)

		searchAddress := currentAddress
		for searchAddress < regionEnd {
			data, err := m.ReadMemory(searchAddress, uint32(len(pattern)))
			if err != nil {
				break
			}

			if len(data) >= len(pattern) && matchPattern(data, pattern) {
				results = append(results, searchAddress)
				searchAddress += uint64(len(pattern))
			} else {
				searchAddress++
			}
		}

		currentAddress = region.BaseAddress + region.RegionSize
	}

	return results, nil
}

func (m *Manager) DumpMemory(address uint64, size uint32) (string, error) {
	data, err := m.ReadMemory(address, size)
	if err != nil {
		return "", err
	}

	return hex.Dump(data), nil
}

func (m *Manager) GetMemoryRegionsInRange(startAddress uint64, endAddress uint64) ([]*MemoryRegion, error) {
	regions := make([]*MemoryRegion, 0)

	for addr, region := range m.cache.Range() {
		if addr >= startAddress && addr < endAddress {
			regions = append(regions, region)
		}
	}

	return regions, nil
}

func (m *Manager) ClearCache() {
	m.cache.Reset()
}

func (m *Manager) formatProtectionString(protect uint32) string {
	result := ""

	if protect&0x01 != 0 {
		result += "R"
	}
	if protect&0x02 != 0 {
		result += "W"
	}
	if protect&0x04 != 0 {
		result += "X"
	}
	if protect&0x08 != 0 {
		result += "G"
	}
	if protect&0x10 != 0 {
		result += "-"
	}

	if result == "" {
		result = "N"
	}

	return result
}

func matchPattern(data []byte, pattern []byte) bool {
	if len(data) < len(pattern) {
		return false
	}

	for i := range pattern {
		if data[i] != pattern[i] {
			return false
		}
	}

	return true
}

func (m *Manager) FindPattern(startAddress uint64, size uint32, pattern string) ([]uint64, error) {
	patternBytes, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}

	return m.SearchMemory(patternBytes, startAddress, startAddress+uint64(size))
}

func parsePattern(pattern string) ([]byte, error) {
	result := make([]byte, 0)
	i := 0

	for i < len(pattern) {
		if pattern[i] == ' ' || pattern[i] == '\t' {
			i++
			continue
		}

		if pattern[i] == '?' {
			if i+1 < len(pattern) && pattern[i+1] == '?' {
				result = append(result, 0x00)
				i += 2
			} else {
				result = append(result, 0x00)
				i++
			}
		} else if i+1 < len(pattern) {
			hexStr := pattern[i : i+2]
			b, err := hex.DecodeString(hexStr)
			if err != nil {
				return nil, err
			}
			result = append(result, b[0])
			i += 2
		} else {
			return nil, fmt.Errorf("invalid pattern format")
		}
	}

	return result, nil
}
