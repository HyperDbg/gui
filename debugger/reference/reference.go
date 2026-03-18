package reference

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type ReferenceType int

const (
	ReferenceCode ReferenceType = iota
	ReferenceData
	ReferenceString
	ReferenceImport
	ReferenceExport
)

type Reference struct {
	Address       uint64
	Type          ReferenceType
	Module        string
	TargetAddress uint64
	TargetModule  string
	TargetName    string
}

type Manager struct {
	refs *safemap.M[uint64, []*Reference]
}

func New() api.Interface {
	return &Manager{
		refs: safemap.New[uint64, []*Reference](),
	}
}

func (m *Manager) AddReference(address uint64, refType ReferenceType, module string, targetAddress uint64, targetModule, targetName string) {
	refs, _ := m.refs.Get(address)
	if refs == nil {
		refs = make([]*Reference, 0)
	}

	ref := &Reference{
		Address:       address,
		Type:          refType,
		Module:        module,
		TargetAddress: targetAddress,
		TargetModule:  targetModule,
		TargetName:    targetName,
	}

	refs = append(refs, ref)
	m.refs.Update(address, refs)
}

func (m *Manager) GetReferences(address uint64) []*Reference {
	if refs, exists := m.refs.Get(address); exists {
		result := make([]*Reference, len(refs))
		copy(result, refs)
		return result
	}
	return nil
}

func (m *Manager) GetReferencesByType(refType ReferenceType) []*Reference {
	result := make([]*Reference, 0)
	for _, refs := range m.refs.Range() {
		for _, ref := range refs {
			if ref.Type == refType {
				result = append(result, ref)
			}
		}
	}
	return result
}

func (m *Manager) GetReferencesByModule(module string) []*Reference {
	result := make([]*Reference, 0)
	for _, refs := range m.refs.Range() {
		for _, ref := range refs {
			if ref.Module == module {
				result = append(result, ref)
			}
		}
	}
	return result
}

func (m *Manager) GetReferencesByTargetModule(targetModule string) []*Reference {
	result := make([]*Reference, 0)
	for _, refs := range m.refs.Range() {
		for _, ref := range refs {
			if ref.TargetModule == targetModule {
				result = append(result, ref)
			}
		}
	}
	return result
}

func (m *Manager) GetReferencesByTargetName(targetName string) []*Reference {
	result := make([]*Reference, 0)
	for _, refs := range m.refs.Range() {
		for _, ref := range refs {
			if ref.TargetName == targetName {
				result = append(result, ref)
			}
		}
	}
	return result
}

func (m *Manager) DeleteReference(address uint64, targetAddress uint64) {
	if refs, exists := m.refs.Get(address); exists {
		for i, ref := range refs {
			if ref.TargetAddress == targetAddress {
				refs = append(refs[:i], refs[i+1:]...)
				m.refs.Update(address, refs)
				break
			}
		}
	}
}

func (m *Manager) DeleteReferences(address uint64) {
	m.refs.Delete(address)
}

func (m *Manager) GetAllReferences() []*Reference {
	result := make([]*Reference, 0)
	for _, refs := range m.refs.Range() {
		result = append(result, refs...)
	}
	return result
}

func (m *Manager) HasReference(address uint64) bool {
	_, exists := m.refs.Get(address)
	return exists
}

func (m *Manager) HasReferenceToTarget(address uint64, targetAddress uint64) bool {
	if refs, exists := m.refs.Get(address); exists {
		for _, ref := range refs {
			if ref.TargetAddress == targetAddress {
				return true
			}
		}
	}
	return false
}

func (m *Manager) GetReferenceCount(address uint64) int {
	if refs, exists := m.refs.Get(address); exists {
		return len(refs)
	}
	return 0
}

func (m *Manager) GetTotalReferenceCount() int {
	count := 0
	for _, refs := range m.refs.Range() {
		count += len(refs)
	}
	return count
}

func (m *Manager) GetReferenceCountByType(refType ReferenceType) int {
	count := 0
	for _, refs := range m.refs.Range() {
		for _, ref := range refs {
			if ref.Type == refType {
				count++
			}
		}
	}
	return count
}

func (m *Manager) Clear() {
	m.refs.Reset()
}

func (m *Manager) ClearByModule(module string) {
	for address, refs := range m.refs.Range() {
		filtered := make([]*Reference, 0)
		for _, ref := range refs {
			if ref.Module != module {
				filtered = append(filtered, ref)
			}
		}
		if len(filtered) == 0 {
			m.refs.Delete(address)
		} else {
			m.refs.Update(address, filtered)
		}
	}
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
