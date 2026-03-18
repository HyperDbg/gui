package comment

import (
	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
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
		return nil
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

func (m *Manager) GetAllComments() []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range m.comments.Range() {
		result = append(result, comment)
	}
	return result
}

func (m *Manager) Clear() {
	m.comments.Reset()
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
