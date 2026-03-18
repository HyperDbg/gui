package notes

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type Note struct {
	Address uint64
	Module  string
	Content string
}

type Manager struct {
	notes *safemap.M[uint64, *Note]
}

func New() api.Interface {
	return &Manager{
		notes: safemap.New[uint64, *Note](),
	}
}

func (m *Manager) AddNote(address uint64, module, content string) error {
	if content == "" {
		return fmt.Errorf("note content cannot be empty")
	}

	m.notes.Update(address, &Note{
		Address: address,
		Module:  module,
		Content: content,
	})

	return nil
}

func (m *Manager) GetNote(address uint64) *Note {
	note, _ := m.notes.Get(address)
	return note
}

func (m *Manager) DeleteNote(address uint64) {
	m.notes.Delete(address)
}

func (m *Manager) GetAllNotes() []*Note {
	result := make([]*Note, 0)
	for _, note := range m.notes.Range() {
		result = append(result, note)
	}
	return result
}

func (m *Manager) GetNotesByModule(module string) []*Note {
	result := make([]*Note, 0)
	for _, note := range m.notes.Range() {
		if note.Module == module {
			result = append(result, note)
		}
	}
	return result
}

func (m *Manager) UpdateNote(address uint64, content string) error {
	if note, exists := m.notes.Get(address); exists {
		note.Content = content
		return nil
	}
	return fmt.Errorf("note not found at address 0x%X", address)
}

func (m *Manager) HasNote(address uint64) bool {
	_, exists := m.notes.Get(address)
	return exists
}

func (m *Manager) Clear() {
	m.notes.Reset()
}

func (m *Manager) GetNoteCount() int {
	count := 0
	for range m.notes.Range() {
		count++
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
