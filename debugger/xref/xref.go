package xref

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
)

type Xref struct {
	FromAddress uint64
	ToAddress   uint64
	Type        string
}

type Manager struct {
	xrefs []Xref
}

func New() api.Interface {
	return &Manager{
		xrefs: make([]Xref, 0),
	}
}

func (m *Manager) AddXref(from, to uint64, xrefType string) {
	m.xrefs = append(m.xrefs, Xref{
		FromAddress: from,
		ToAddress:   to,
		Type:        xrefType,
	})
}

func (m *Manager) GetXrefsTo(address uint64) []Xref {
	result := make([]Xref, 0)
	for _, xref := range m.xrefs {
		if xref.ToAddress == address {
			result = append(result, xref)
		}
	}
	return result
}

func (m *Manager) GetXrefsFrom(address uint64) []Xref {
	result := make([]Xref, 0)
	for _, xref := range m.xrefs {
		if xref.FromAddress == address {
			result = append(result, xref)
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.xrefs = make([]Xref, 0)
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
