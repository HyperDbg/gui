package source

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type SourceLine struct {
	LineNumber   int
	Address      uint64
	Content      string
	IsBreakpoint bool
}

type SourceFile struct {
	Path      string
	Module    string
	Lines     []*SourceLine
	LineCount int
}

type Manager struct {
	files *safemap.M[string, *SourceFile]
}

func New() api.Interface {
	return &Manager{
		files: safemap.New[string, *SourceFile](),
	}
}

func (m *Manager) LoadFile(path, module string) error {
	file := &SourceFile{
		Path:      path,
		Module:    module,
		Lines:     make([]*SourceLine, 0),
		LineCount: 0,
	}

	m.files.Update(path, file)
	return nil
}

func (m *Manager) GetFile(path string) *SourceFile {
	file, _ := m.files.Get(path)
	return file
}

func (m *Manager) GetFileByModule(module string) *SourceFile {
	for _, file := range m.files.Range() {
		if file.Module == module {
			return file
		}
	}
	return nil
}

func (m *Manager) GetAllFiles() []*SourceFile {
	result := make([]*SourceFile, 0)
	for _, file := range m.files.Range() {
		result = append(result, file)
	}
	return result
}

func (m *Manager) DeleteFile(path string) {
	m.files.Delete(path)
}

func (m *Manager) AddLine(path string, lineNumber int, address uint64, content string) error {
	if file, exists := m.files.Get(path); exists {
		line := &SourceLine{
			LineNumber:   lineNumber,
			Address:      address,
			Content:      content,
			IsBreakpoint: false,
		}
		file.Lines = append(file.Lines, line)
		file.LineCount++
		return nil
	}
	return fmt.Errorf("file not found: %s", path)
}

func (m *Manager) GetLine(path string, lineNumber int) *SourceLine {
	if file, exists := m.files.Get(path); exists {
		for _, line := range file.Lines {
			if line.LineNumber == lineNumber {
				return line
			}
		}
	}
	return nil
}

func (m *Manager) GetLineByAddress(address uint64) *SourceLine {
	for _, file := range m.files.Range() {
		for _, line := range file.Lines {
			if line.Address == address {
				return line
			}
		}
	}
	return nil
}

func (m *Manager) SetBreakpoint(path string, lineNumber int, isBreakpoint bool) error {
	if line := m.GetLine(path, lineNumber); line != nil {
		line.IsBreakpoint = isBreakpoint
		return nil
	}
	return fmt.Errorf("line not found")
}

func (m *Manager) HasBreakpoint(path string, lineNumber int) bool {
	if line := m.GetLine(path, lineNumber); line != nil {
		return line.IsBreakpoint
	}
	return false
}

func (m *Manager) ClearBreakpoints(path string) {
	if file, exists := m.files.Get(path); exists {
		for _, line := range file.Lines {
			line.IsBreakpoint = false
		}
	}
}

func (m *Manager) ClearAllBreakpoints() {
	for _, file := range m.files.Range() {
		for _, line := range file.Lines {
			line.IsBreakpoint = false
		}
	}
}

func (m *Manager) GetBreakpoints() []*SourceLine {
	result := make([]*SourceLine, 0)
	for _, file := range m.files.Range() {
		for _, line := range file.Lines {
			if line.IsBreakpoint {
				result = append(result, line)
			}
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.files.Reset()
}

func (m *Manager) GetFileCount() int {
	count := 0
	for range m.files.Range() {
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
