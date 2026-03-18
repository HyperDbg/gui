package watch

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
)

type Watch struct {
	Address uint64
	Name    string
	Size    uint64
	Type    string
}

type Manager struct {
	watches []Watch
}

func New() api.Interface {
	return &Manager{
		watches: make([]Watch, 0),
	}
}

func (m *Manager) AddWatch(address uint64, name string, size uint64, watchType string) {
	m.watches = append(m.watches, Watch{
		Address: address,
		Name:    name,
		Size:    size,
		Type:    watchType,
	})
}

func (m *Manager) RemoveWatch(address uint64) {
	for i, watch := range m.watches {
		if watch.Address == address {
			m.watches = append(m.watches[:i], m.watches[i+1:]...)
			break
		}
	}
}

func (m *Manager) GetAllWatches() []Watch {
	return m.watches
}

func (m *Manager) Clear() {
	m.watches = make([]Watch, 0)
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
