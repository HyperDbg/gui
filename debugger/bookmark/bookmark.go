package bookmark

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type Bookmark struct {
	Address uint64
	Name    string
	Module  string
	Auto    bool
	Notes   string
}

type Manager struct {
	bookmarks *safemap.M[uint64, *Bookmark]
}

func New() api.Interface {
	return &Manager{
		bookmarks: safemap.New[uint64, *Bookmark](),
	}
}

func (m *Manager) SetBookmark(address uint64, name string, auto bool, notes string) error {
	if name == "" {
		return nil
	}

	m.bookmarks.Update(address, &Bookmark{
		Address: address,
		Name:    name,
		Auto:    auto,
		Notes:   notes,
	})

	return nil
}

func (m *Manager) GetBookmark(address uint64) *Bookmark {
	bookmark, _ := m.bookmarks.Get(address)
	return bookmark
}

func (m *Manager) DeleteBookmark(address uint64) {
	m.bookmarks.Delete(address)
}

func (m *Manager) GetAllBookmarks() []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		result = append(result, bookmark)
	}
	return result
}

func (m *Manager) Clear() {
	m.bookmarks.Reset()
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
