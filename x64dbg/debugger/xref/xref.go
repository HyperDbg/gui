package xref

import (
	"fmt"

	"gioui.org/layout"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type XrefType int

const (
	XrefNone XrefType = iota
	XrefData
	XrefJmp
	XrefCall
)

type Xref struct {
	Address uint64
	From    uint64
	Type    XrefType
	Module  string
}

type Manager struct {
	xrefs *safemap.M[uint64, []*Xref]
}

func New() api.Interface {
	return &Manager{
		xrefs: safemap.New[uint64, []*Xref](),
	}
}

func (m *Manager) AddXref(address, from uint64, xrefType XrefType) {
	xrefs, _ := m.xrefs.Get(address)
	if xrefs == nil {
		xrefs = make([]*Xref, 0)
		m.xrefs.Update(address, xrefs)
	}

	xrefs = append(xrefs, &Xref{
		Address: address,
		From:    from,
		Type:    xrefType,
	})
	m.xrefs.Update(address, xrefs)
}

func (m *Manager) GetXrefs(address uint64) []*Xref {
	if xrefs, exists := m.xrefs.Get(address); exists {
		result := make([]*Xref, len(xrefs))
		copy(result, xrefs)
		return result
	}
	return nil
}

func (m *Manager) GetXrefCount(address uint64) int {
	if xrefs, exists := m.xrefs.Get(address); exists {
		return len(xrefs)
	}
	return 0
}

func (m *Manager) GetXrefType(address uint64) XrefType {
	if xrefs, exists := m.xrefs.Get(address); exists && len(xrefs) > 0 {
		return xrefs[0].Type
	}
	return XrefNone
}

func (m *Manager) DeleteXref(address uint64) {
	m.xrefs.Delete(address)
}

func (m *Manager) DeleteAllXrefs() {
	m.xrefs.Reset()
}

func (m *Manager) GetAllXrefs() map[uint64][]*Xref {
	result := make(map[uint64][]*Xref)
	for addr, xrefs := range m.xrefs.Range() {
		result[addr] = make([]*Xref, len(xrefs))
		copy(result[addr], xrefs)
	}
	return result
}

func (m *Manager) GetDataXrefs(address uint64) []*Xref {
	if xrefs, exists := m.xrefs.Get(address); exists {
		result := make([]*Xref, 0)
		for _, xref := range xrefs {
			if xref.Type == XrefData {
				result = append(result, xref)
			}
		}
		return result
	}
	return nil
}

func (m *Manager) GetJmpXrefs(address uint64) []*Xref {
	if xrefs, exists := m.xrefs.Get(address); exists {
		result := make([]*Xref, 0)
		for _, xref := range xrefs {
			if xref.Type == XrefJmp {
				result = append(result, xref)
			}
		}
		return result
	}
	return nil
}

func (m *Manager) GetCallXrefs(address uint64) []*Xref {
	if xrefs, exists := m.xrefs.Get(address); exists {
		result := make([]*Xref, 0)
		for _, xref := range xrefs {
			if xref.Type == XrefCall {
				result = append(result, xref)
			}
		}
		return result
	}
	return nil
}

func (m *Manager) HasXref(address uint64) bool {
	_, exists := m.xrefs.Get(address)
	return exists
}

func (m *Manager) HasDataXref(address uint64) bool {
	if xrefs, exists := m.xrefs.Get(address); exists {
		for _, xref := range xrefs {
			if xref.Type == XrefData {
				return true
			}
		}
	}
	return false
}

func (m *Manager) HasJmpXref(address uint64) bool {
	if xrefs, exists := m.xrefs.Get(address); exists {
		for _, xref := range xrefs {
			if xref.Type == XrefJmp {
				return true
			}
		}
	}
	return false
}

func (m *Manager) HasCallXref(address uint64) bool {
	if xrefs, exists := m.xrefs.Get(address); exists {
		for _, xref := range xrefs {
			if xref.Type == XrefCall {
				return true
			}
		}
	}
	return false
}

func (m *Manager) SetXrefModule(address uint64, module string) {
	if xrefs, exists := m.xrefs.Get(address); exists {
		for _, xref := range xrefs {
			xref.Module = module
		}
	}
}

func (m *Manager) GetXrefsFrom(from uint64) []*Xref {
	result := make([]*Xref, 0)
	for _, xrefs := range m.xrefs.Range() {
		for _, xref := range xrefs {
			if xref.From == from {
				result = append(result, xref)
			}
		}
	}
	return result
}

func (m *Manager) GetTotalXrefCount() int {
	count := 0
	for _, xrefs := range m.xrefs.Range() {
		count += len(xrefs)
	}
	return count
}

func (m *Manager) GetReferencedAddressCount() int {
	count := 0
	for range m.xrefs.Range() {
		count++
	}
	return count
}

func (m *Manager) Clear() {
	m.xrefs.Reset()
}

func (m *Manager) RemoveXref(address, from uint64) {
	xrefs, exists := m.xrefs.Get(address)
	if exists {
		for i, xref := range xrefs {
			if xref.From == from {
				xrefs = append(xrefs[:i], xrefs[i+1:]...)
				m.xrefs.Update(address, xrefs)
				break
			}
		}
	}
}

func (m *Manager) GetXrefsInModule(module string) []*Xref {
	result := make([]*Xref, 0)
	for _, xrefs := range m.xrefs.Range() {
		for _, xref := range xrefs {
			if xref.Module == module {
				result = append(result, xref)
			}
		}
	}
	return result
}

func (m *Manager) AddXrefs(address uint64, from []uint64, xrefTypes []XrefType) error {
	if len(from) != len(xrefTypes) {
		return fmt.Errorf("from and xrefTypes must have the same length")
	}

	xrefs, _ := m.xrefs.Get(address)
	if xrefs == nil {
		xrefs = make([]*Xref, 0)
		m.xrefs.Update(address, xrefs)
	}

	for i := range from {
		xrefs = append(xrefs, &Xref{
			Address: address,
			From:    from[i],
			Type:    xrefTypes[i],
		})
		m.xrefs.Update(address, xrefs)
	}

	return nil
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
