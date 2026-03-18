package symbol

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type SymbolType int

const (
	SymbolTypeNone SymbolType = iota
	SymbolTypeFunction
	SymbolTypeVariable
	SymbolTypeParameter
	SymbolTypeLabel
	SymbolTypeExport
	SymbolTypeImport
)

type Symbol struct {
	Name       string
	Address    uint64
	Size       uint64
	Type       SymbolType
	Module     string
	FileName   string
	LineNumber int
}

type Module struct {
	Name        string
	BaseAddress uint64
	Size        uint64
	Path        string
	Symbols     map[string]*Symbol
}

type Provider struct {
	mu      sync.RWMutex
	symbols *safemap.M[uint64, *Symbol]
	modules *safemap.M[string, *Module]
	table   *treetable.TreeTable[Symbol]
}

func New() api.Interface {
	m := &Provider{
		symbols: safemap.New[uint64, *Symbol](),
		modules: safemap.New[string, *Module](),
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = treetable.NewTreeTable[Symbol]()
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	symbols := m.GetAllSymbols()
	for _, symbol := range symbols {
		m.table.Root().AddChild(m.table.NewNode(*symbol))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) AddSymbol(symbol *Symbol) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.symbols.Update(symbol.Address, symbol)
}

func (m *Provider) GetSymbol(address uint64) *Symbol {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.symbols.GetMust(address)
}

func (m *Provider) GetSymbolByName(name string) *Symbol {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, symbol := range m.symbols.Values() {
		if symbol.Name == name {
			return symbol
		}
	}
	return nil
}

func (m *Provider) FindNearestSymbol(address uint64) *Symbol {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var nearest *Symbol
	for _, symbol := range m.symbols.Values() {
		if address >= symbol.Address && address < symbol.Address+symbol.Size {
			nearest = symbol
			break
		}
	}
	return nearest
}

func (m *Provider) GetAllSymbols() []*Symbol {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.symbols.Values()
}

func (m *Provider) AddModule(module *Module) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.modules.Update(module.Name, module)
}

func (m *Provider) GetModule(name string) *Module {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.modules.GetMust(name)
}

func (m *Provider) GetModuleByAddress(address uint64) *Module {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, module := range m.modules.Values() {
		if address >= module.BaseAddress && address < module.BaseAddress+module.Size {
			return module
		}
	}
	return nil
}

func (m *Provider) GetAllModules() []*Module {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.modules.Values()
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.symbols.Reset()
	m.modules.Reset()
}
