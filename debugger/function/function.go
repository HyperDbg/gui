package function

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type Function struct {
	Address uint64
	Name    string
	Module  string
	Size    uint64
}

type Manager struct {
	functions *safemap.M[uint64, *Function]
}

func New() api.Interface {
	return &Manager{
		functions: safemap.New[uint64, *Function](),
	}
}

func (m *Manager) AddFunction(address uint64, name string, module string, size uint64) error {
	if name == "" {
		return nil
	}

	m.functions.Update(address, &Function{
		Address: address,
		Name:    name,
		Module:  module,
		Size:    size,
	})

	return nil
}

func (m *Manager) GetFunction(address uint64) *Function {
	function, _ := m.functions.Get(address)
	return function
}

func (m *Manager) DeleteFunction(address uint64) {
	m.functions.Delete(address)
}

func (m *Manager) GetAllFunctions() []*Function {
	result := make([]*Function, 0)
	for _, function := range m.functions.Range() {
		result = append(result, function)
	}
	return result
}

func (m *Manager) Clear() {
	m.functions.Reset()
}

func (m *Manager) Layout() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{}
	}
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Self() any {
	return m
}
