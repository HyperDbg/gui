package label

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type Label struct {
	Address uint64
	Name    string
	Module  string
	Auto    bool
}

type Manager struct {
	labels *safemap.M[uint64, *Label]
}

func New() api.Interface {
	return &Manager{
		labels: safemap.New[uint64, *Label](),
	}
}

func (m *Manager) SetLabel(address uint64, name string, auto bool) error {
	if name == "" {
		return nil
	}

	m.labels.Update(address, &Label{
		Address: address,
		Name:    name,
		Auto:    auto,
	})

	return nil
}

func (m *Manager) GetLabel(address uint64) *Label {
	label, _ := m.labels.Get(address)
	return label
}

func (m *Manager) DeleteLabel(address uint64) {
	m.labels.Delete(address)
}

func (m *Manager) GetAllLabels() []*Label {
	result := make([]*Label, 0)
	for _, label := range m.labels.Range() {
		result = append(result, label)
	}
	return result
}

func (m *Manager) Clear() {
	m.labels.Reset()
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
