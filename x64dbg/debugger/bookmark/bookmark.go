package bookmark

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
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
		return fmt.Errorf("bookmark name cannot be empty")
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

func (m *Manager) GetBookmarkName(address uint64) string {
	if bookmark, exists := m.bookmarks.Get(address); exists {
		return bookmark.Name
	}
	return ""
}

func (m *Manager) DeleteBookmark(address uint64) {
	m.bookmarks.Delete(address)
}

func (m *Manager) DeleteBookmarkRange(start, end uint64) {
	for addr := range m.bookmarks.Range() {
		if addr >= start && addr <= end {
			m.bookmarks.Delete(addr)
		}
	}
}

func (m *Manager) GetAllBookmarks() []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		result = append(result, bookmark)
	}
	return result
}

func (m *Manager) GetAutoBookmarks() []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		if bookmark.Auto {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) GetUserBookmarks() []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		if !bookmark.Auto {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.bookmarks.Reset()
}

func (m *Manager) ClearAutoBookmarks() {
	for addr, bookmark := range m.bookmarks.Range() {
		if bookmark.Auto {
			m.bookmarks.Delete(addr)
		}
	}
}

func (m *Manager) HasBookmark(address uint64) bool {
	_, exists := m.bookmarks.Get(address)
	return exists
}

func (m *Manager) FindBookmarksByName(name string) []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		if bookmark.Name == name {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) SetBookmarkModule(address uint64, module string) {
	if bookmark, exists := m.bookmarks.Get(address); exists {
		bookmark.Module = module
	}
}

func (m *Manager) SetBookmarkNotes(address uint64, notes string) {
	if bookmark, exists := m.bookmarks.Get(address); exists {
		bookmark.Notes = notes
	}
}

func (m *Manager) GetBookmarkCount() int {
	count := 0
	for range m.bookmarks.Range() {
		count++
	}
	return count
}

func (m *Manager) GetAutoBookmarkCount() int {
	count := 0
	for _, bookmark := range m.bookmarks.Range() {
		if bookmark.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) GetUserBookmarkCount() int {
	count := 0
	for _, bookmark := range m.bookmarks.Range() {
		if !bookmark.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) GetBookmarksInModule(module string) []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		if bookmark.Module == module {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) FindBookmarksByNotes(notes string) []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		if bookmark.Notes == notes {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) GetNextBookmark(address uint64) *Bookmark {
	nextAddr := uint64(0xFFFFFFFFFFFFFFFF)
	var nextBookmark *Bookmark

	for addr, bookmark := range m.bookmarks.Range() {
		if addr > address && addr < nextAddr {
			nextAddr = addr
			nextBookmark = bookmark
		}
	}

	return nextBookmark
}

func (m *Manager) GetPrevBookmark(address uint64) *Bookmark {
	prevAddr := uint64(0)
	var prevBookmark *Bookmark

	for addr, bookmark := range m.bookmarks.Range() {
		if addr < address && addr > prevAddr {
			prevAddr = addr
			prevBookmark = bookmark
		}
	}

	return prevBookmark
}

func (m *Manager) GetFirstBookmark() *Bookmark {
	firstAddr := uint64(0xFFFFFFFFFFFFFFFF)
	var firstBookmark *Bookmark

	for addr, bookmark := range m.bookmarks.Range() {
		if addr < firstAddr {
			firstAddr = addr
			firstBookmark = bookmark
		}
	}

	return firstBookmark
}

func (m *Manager) GetLastBookmark() *Bookmark {
	lastAddr := uint64(0)
	var lastBookmark *Bookmark

	for addr, bookmark := range m.bookmarks.Range() {
		if addr > lastAddr {
			lastAddr = addr
			lastBookmark = bookmark
		}
	}

	return lastBookmark
}

func (m *Manager) GetBookmarksInRange(start, end uint64) []*Bookmark {
	result := make([]*Bookmark, 0)
	for addr, bookmark := range m.bookmarks.Range() {
		if addr >= start && addr <= end {
			result = append(result, bookmark)
		}
	}
	return result
}

func (m *Manager) SetBookmarkName(address uint64, name string) error {
	if name == "" {
		return fmt.Errorf("bookmark name cannot be empty")
	}

	if bookmark, exists := m.bookmarks.Get(address); exists {
		bookmark.Name = name
		return nil
	}

	return fmt.Errorf("bookmark not found at address: 0x%X", address)
}

func (m *Manager) ToggleBookmark(address uint64) {
	if _, exists := m.bookmarks.Get(address); exists {
		m.bookmarks.Delete(address)
	} else {
		m.bookmarks.Update(address, &Bookmark{
			Address: address,
			Name:    fmt.Sprintf("bookmark_%X", address),
			Auto:    false,
			Notes:   "",
		})
	}
}

func (m *Manager) GetSortedBookmarks() []*Bookmark {
	result := make([]*Bookmark, 0)
	for _, bookmark := range m.bookmarks.Range() {
		result = append(result, bookmark)
	}

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].Address > result[j].Address {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

func (m *Manager) GetBookmarkIndex(address uint64) int {
	sorted := m.GetSortedBookmarks()
	for i, bookmark := range sorted {
		if bookmark.Address == address {
			return i
		}
	}
	return -1
}

func (m *Manager) GetBookmarkByIndex(index int) *Bookmark {
	sorted := m.GetSortedBookmarks()
	if index >= 0 && index < len(sorted) {
		return sorted[index]
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
