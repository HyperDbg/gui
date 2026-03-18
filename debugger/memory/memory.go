package memory

import (
	"fmt"
	"strings"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
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

type MemoryDump struct {
	Address uint64
	Data    []byte
	Size    int
}

type Provider struct {
	mu      sync.RWMutex
	regions map[uint64]*MemoryRegion
	dumps   map[uint64]*MemoryDump
	watches map[uint64]bool
}

func New() api.Interface {
	return &Provider{
		regions: make(map[uint64]*MemoryRegion),
		dumps:   make(map[uint64]*MemoryDump),
		watches: make(map[uint64]bool),
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

func (m *Provider) GetMemoryRegions() []*MemoryRegion {
	m.mu.RLock()
	defer m.mu.RUnlock()

	regions := make([]*MemoryRegion, 0, len(m.regions))
	for _, region := range m.regions {
		regions = append(regions, region)
	}
	return regions
}

func (m *Provider) AddMemoryRegion(region *MemoryRegion) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.regions[region.BaseAddress] = region
}

func (m *Provider) GetMemoryDump(address uint64) *MemoryDump {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.dumps[address]
}

func (m *Provider) SetMemoryDump(dump *MemoryDump) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dumps[dump.Address] = dump
}

func (m *Provider) FormatMemoryDump(address uint64, data []byte) string {
	var result strings.Builder
	for i := 0; i < len(data); i += 16 {
		result.WriteString(fmt.Sprintf("%016x: ", address+uint64(i)))

		var hexPart strings.Builder
		var asciiPart strings.Builder
		for j := range 16 {
			if i+j < len(data) {
				hexPart.WriteString(fmt.Sprintf("%02x ", data[i+j]))
				if data[i+j] >= 32 && data[i+j] <= 126 {
					asciiPart.WriteString(string(data[i+j]))
				} else {
					asciiPart.WriteString(".")
				}
			} else {
				hexPart.WriteString("   ")
			}
		}
		result.WriteString(hexPart.String() + " |" + asciiPart.String() + "|\n")
	}
	return result.String()
}

func (m *Provider) AddWatch(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.watches[address] = true
}

func (m *Provider) RemoveWatch(address uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.watches, address)
}

func (m *Provider) GetWatches() []uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	watches := make([]uint64, 0, len(m.watches))
	for addr := range m.watches {
		watches = append(watches, addr)
	}
	return watches
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.regions = make(map[uint64]*MemoryRegion)
	m.dumps = make(map[uint64]*MemoryDump)
	m.watches = make(map[uint64]bool)
}

func (m *Provider) ReadMemory(pid uint32, address uint64, size uint32) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if dump, exists := m.dumps[address]; exists {
		if dump.Size >= int(size) {
			return dump.Data[:size], nil
		}
	}

	return nil, fmt.Errorf("memory dump not found at address 0x%x", address)
}
