package label

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
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
		return fmt.Errorf("label name cannot be empty")
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

func (m *Manager) GetLabelName(address uint64) string {
	if label, exists := m.labels.Get(address); exists {
		return label.Name
	}
	return ""
}

func (m *Manager) DeleteLabel(address uint64) {
	m.labels.Delete(address)
}

func (m *Manager) DeleteLabelRange(start, end uint64) {
	for addr := range m.labels.Range() {
		if addr >= start && addr <= end {
			m.labels.Delete(addr)
		}
	}
}

func (m *Manager) GetAllLabels() []*Label {
	result := make([]*Label, 0)
	for _, label := range m.labels.Range() {
		result = append(result, label)
	}
	return result
}

func (m *Manager) GetAutoLabels() []*Label {
	result := make([]*Label, 0)
	for _, label := range m.labels.Range() {
		if label.Auto {
			result = append(result, label)
		}
	}
	return result
}

func (m *Manager) GetUserLabels() []*Label {
	result := make([]*Label, 0)
	for _, label := range m.labels.Range() {
		if !label.Auto {
			result = append(result, label)
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.labels.Reset()
}

func (m *Manager) ClearAutoLabels() {
	for addr, label := range m.labels.Range() {
		if label.Auto {
			m.labels.Delete(addr)
		}
	}
}

func (m *Manager) HasLabel(address uint64) bool {
	_, exists := m.labels.Get(address)
	return exists
}

func (m *Manager) FindLabelsByName(name string) []*Label {
	result := make([]*Label, 0)
	for _, label := range m.labels.Range() {
		if label.Name == name {
			result = append(result, label)
		}
	}
	return result
}

func (m *Manager) SetLabelModule(address uint64, module string) {
	if label, exists := m.labels.Get(address); exists {
		label.Module = module
	}
}

func (m *Manager) GetLabelCount() int {
	count := 0
	for range m.labels.Range() {
		count++
	}
	return count
}

func (m *Manager) GetAutoLabelCount() int {
	count := 0
	for _, label := range m.labels.Range() {
		if label.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) GetUserLabelCount() int {
	count := 0
	for _, label := range m.labels.Range() {
		if !label.Auto {
			count++
		}
	}
	return count
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
