package comment

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type Comment struct {
	Address uint64
	Text    string
	Auto    bool
	Module  string
}

type Manager struct {
	comments *safemap.M[uint64, *Comment]
}

func New() api.Interface {
	return &Manager{
		comments: safemap.New[uint64, *Comment](),
	}
}

func (m *Manager) SetComment(address uint64, text string, auto bool) error {
	if text == "" {
		return fmt.Errorf("comment text cannot be empty")
	}

	m.comments.Update(address, &Comment{
		Address: address,
		Text:    text,
		Auto:    auto,
	})

	return nil
}

func (m *Manager) GetComment(address uint64) *Comment {
	comment, _ := m.comments.Get(address)
	return comment
}

func (m *Manager) GetCommentText(address uint64) string {
	if comment, exists := m.comments.Get(address); exists {
		return comment.Text
	}
	return ""
}

func (m *Manager) DeleteComment(address uint64) {
	m.comments.Delete(address)
}

func (m *Manager) DeleteCommentRange(start, end uint64) {
	for addr := range m.comments.Range() {
		if addr >= start && addr <= end {
			m.comments.Delete(addr)
		}
	}
}

func (m *Manager) GetAllComments() []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range m.comments.Range() {
		result = append(result, comment)
	}
	return result
}

func (m *Manager) GetAutoComments() []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range m.comments.Range() {
		if comment.Auto {
			result = append(result, comment)
		}
	}
	return result
}

func (m *Manager) GetUserComments() []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range m.comments.Range() {
		if !comment.Auto {
			result = append(result, comment)
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.comments.Reset()
}

func (m *Manager) ClearAutoComments() {
	for addr, comment := range m.comments.Range() {
		if comment.Auto {
			m.comments.Delete(addr)
		}
	}
}

func (m *Manager) HasComment(address uint64) bool {
	_, exists := m.comments.Get(address)
	return exists
}

func (m *Manager) FindCommentsByText(text string) []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range m.comments.Range() {
		if comment.Text == text {
			result = append(result, comment)
		}
	}
	return result
}

func (m *Manager) SetCommentModule(address uint64, module string) {
	if comment, exists := m.comments.Get(address); exists {
		comment.Module = module
	}
}

func (m *Manager) GetCommentCount() int {
	count := 0
	for range m.comments.Range() {
		count++
	}
	return count
}

func (m *Manager) GetAutoCommentCount() int {
	count := 0
	for _, comment := range m.comments.Range() {
		if comment.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) GetUserCommentCount() int {
	count := 0
	for _, comment := range m.comments.Range() {
		if !comment.Auto {
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
