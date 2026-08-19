package function

import (
	"fmt"

	"gioui.org/layout"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type FunctionType int

const (
	FunctionNone FunctionType = iota
	FunctionBegin
	FunctionMiddle
	FunctionEnd
	FunctionSingle
)

type Function struct {
	Start            uint64
	End              uint64
	InstructionCount uint64
	Name             string
	Module           string
	Type             FunctionType
	Auto             bool
}

type Manager struct {
	functions *safemap.M[uint64, *Function]
}

func New() api.Interface {
	return &Manager{
		functions: safemap.New[uint64, *Function](),
	}
}

func (m *Manager) AddFunction(start, end uint64, name string, auto bool) error {
	if start >= end {
		return fmt.Errorf("function start address must be less than end address")
	}

	if name == "" {
		name = fmt.Sprintf("sub_%X", start)
	}

	m.functions.Update(start, &Function{
		Start:            start,
		End:              end,
		InstructionCount: 0,
		Name:             name,
		Type:             FunctionBegin,
		Auto:             auto,
	})

	return nil
}

func (m *Manager) GetFunction(address uint64) *Function {
	for _, fn := range m.functions.Range() {
		if address >= fn.Start && address <= fn.End {
			return fn
		}
	}
	return nil
}

func (m *Manager) GetFunctionByStart(start uint64) *Function {
	fn, _ := m.functions.Get(start)
	return fn
}

func (m *Manager) GetFunctionByEnd(end uint64) *Function {
	for _, fn := range m.functions.Range() {
		if fn.End == end {
			return fn
		}
	}
	return nil
}

func (m *Manager) DeleteFunction(start uint64) {
	m.functions.Delete(start)
}

func (m *Manager) DeleteFunctionRange(start, end uint64) {
	for addr, fn := range m.functions.Range() {
		if fn.Start >= start && fn.End <= end {
			m.functions.Delete(addr)
		}
	}
}

func (m *Manager) GetAllFunctions() []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		result = append(result, fn)
	}
	return result
}

func (m *Manager) GetAutoFunctions() []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		if fn.Auto {
			result = append(result, fn)
		}
	}
	return result
}

func (m *Manager) GetUserFunctions() []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		if !fn.Auto {
			result = append(result, fn)
		}
	}
	return result
}

func (m *Manager) Clear() {
	m.functions.Reset()
}

func (m *Manager) ClearAutoFunctions() {
	for addr, fn := range m.functions.Range() {
		if fn.Auto {
			m.functions.Delete(addr)
		}
	}
}

func (m *Manager) HasFunction(address uint64) bool {
	return m.GetFunction(address) != nil
}

func (m *Manager) Overlaps(start, end uint64) bool {
	for _, fn := range m.functions.Range() {
		if start < fn.End && end > fn.Start {
			return true
		}
	}
	return false
}

func (m *Manager) GetOverlappingFunctions(start, end uint64) []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		if start < fn.End && end > fn.Start {
			result = append(result, fn)
		}
	}
	return result
}

func (m *Manager) SetFunctionType(start uint64, fnType FunctionType) {
	if fn, exists := m.functions.Get(start); exists {
		fn.Type = fnType
	}
}

func (m *Manager) SetFunctionName(start uint64, name string) {
	if fn, exists := m.functions.Get(start); exists {
		fn.Name = name
	}
}

func (m *Manager) SetInstructionCount(start uint64, count uint64) {
	if fn, exists := m.functions.Get(start); exists {
		fn.InstructionCount = count
	}
}

func (m *Manager) GetFunctionCount() int {
	count := 0
	for range m.functions.Range() {
		count++
	}
	return count
}

func (m *Manager) GetAutoFunctionCount() int {
	count := 0
	for _, fn := range m.functions.Range() {
		if fn.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) GetUserFunctionCount() int {
	count := 0
	for _, fn := range m.functions.Range() {
		if !fn.Auto {
			count++
		}
	}
	return count
}

func (m *Manager) FindFunctionsByName(name string) []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		if fn.Name == name {
			result = append(result, fn)
		}
	}
	return result
}

func (m *Manager) GetFunctionsInModule(module string) []*Function {
	result := make([]*Function, 0)
	for _, fn := range m.functions.Range() {
		if fn.Module == module {
			result = append(result, fn)
		}
	}
	return result
}

func (m *Manager) SetFunctionModule(start uint64, module string) {
	if fn, exists := m.functions.Get(start); exists {
		fn.Module = module
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
