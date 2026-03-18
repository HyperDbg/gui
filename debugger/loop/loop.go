package loop

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type LoopType int

const (
	LoopNone LoopType = iota
	LoopBegin
	LoopMiddle
	LoopEntry
	LoopEnd
	LoopSingle
)

type Loop struct {
	Depth            int
	Start            uint64
	End              uint64
	InstructionCount uint64
	Parent           *Loop
	Children         []*Loop
	Module           string
	Auto             bool
	Type             LoopType
}

type Manager struct {
	loops *safemap.M[int, []*Loop]
}

func New() api.Interface {
	return &Manager{
		loops: safemap.New[int, []*Loop](),
	}
}

func (m *Manager) AddLoop(depth int, start, end uint64, auto bool) error {
	if start >= end {
		return fmt.Errorf("loop start address must be less than end address")
	}

	if depth < 0 {
		return fmt.Errorf("loop depth must be non-negative")
	}

	loops, _ := m.loops.Get(depth)
	if loops == nil {
		loops = make([]*Loop, 0)
		m.loops.Update(depth, loops)
	}

	loop := &Loop{
		Depth:            depth,
		Start:            start,
		End:              end,
		InstructionCount: 0,
		Children:         make([]*Loop, 0),
		Auto:             auto,
		Type:             LoopBegin,
	}

	loops = append(loops, loop)
	m.loops.Update(depth, loops)
	return nil
}

func (m *Manager) GetLoop(depth int, start uint64) *Loop {
	if loops, exists := m.loops.Get(depth); exists {
		for _, loop := range loops {
			if loop.Start == start {
				return loop
			}
		}
	}
	return nil
}

func (m *Manager) GetLoopByAddress(address uint64) *Loop {
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if address >= loop.Start && address <= loop.End {
				return loop
			}
		}
	}
	return nil
}

func (m *Manager) GetLoopsByDepth(depth int) []*Loop {
	if loops, exists := m.loops.Get(depth); exists {
		result := make([]*Loop, len(loops))
		copy(result, loops)
		return result
	}
	return nil
}

func (m *Manager) GetAllLoops() []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			result = append(result, loop)
		}
	}
	return result
}

func (m *Manager) DeleteLoop(depth int, start uint64) {
	loops, exists := m.loops.Get(depth)
	if exists {
		for i, loop := range loops {
			if loop.Start == start {
				loops = append(loops[:i], loops[i+1:]...)
				m.loops.Update(depth, loops)
				break
			}
		}
	}
}

func (m *Manager) DeleteLoopRange(start, end uint64) {
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for i, loop := range loops {
			if loop.Start >= start && loop.End <= end {
				loops = append(loops[:i], loops[i+1:]...)
				m.loops.Update(depth, loops)
				break
			}
		}
	}
}

func (m *Manager) Clear() {
	m.loops.Reset()
}

func (m *Manager) HasLoop(depth int, start uint64) bool {
	return m.GetLoop(depth, start) != nil
}

func (m *Manager) HasLoopAtAddress(address uint64) bool {
	return m.GetLoopByAddress(address) != nil
}

func (m *Manager) Overlaps(depth int, start, end uint64) bool {
	if loops, exists := m.loops.Get(depth); exists {
		for _, loop := range loops {
			if start < loop.End && end > loop.Start {
				return true
			}
		}
	}
	return false
}

func (m *Manager) GetOverlappingLoops(start, end uint64) []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if start < loop.End && end > loop.Start {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) SetLoopType(depth int, start uint64, loopType LoopType) {
	if loop := m.GetLoop(depth, start); loop != nil {
		loop.Type = loopType
	}
}

func (m *Manager) SetInstructionCount(depth int, start uint64, count uint64) {
	if loop := m.GetLoop(depth, start); loop != nil {
		loop.InstructionCount = count
	}
}

func (m *Manager) SetLoopModule(depth int, start uint64, module string) {
	if loop := m.GetLoop(depth, start); loop != nil {
		loop.Module = module
	}
}

func (m *Manager) SetParentLoop(depth int, start uint64, parent *Loop) {
	if loop := m.GetLoop(depth, start); loop != nil {
		loop.Parent = parent
	}
}

func (m *Manager) AddChildLoop(depth int, start uint64, child *Loop) {
	if loop := m.GetLoop(depth, start); loop != nil {
		loop.Children = append(loop.Children, child)
	}
}

func (m *Manager) GetLoopCount() int {
	count := 0
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		count += len(loops)
	}
	return count
}

func (m *Manager) GetLoopCountByDepth(depth int) int {
	if loops, exists := m.loops.Get(depth); exists {
		return len(loops)
	}
	return 0
}

func (m *Manager) GetMaxDepth() int {
	maxDepth := 0
	for depth := range m.loops.Range() {
		if depth > maxDepth {
			maxDepth = depth
		}
	}
	return maxDepth
}

func (m *Manager) GetAutoLoops() []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if loop.Auto {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) GetUserLoops() []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if !loop.Auto {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) ClearAutoLoops() {
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for i, loop := range loops {
			if loop.Auto {
				loops = append(loops[:i], loops[i+1:]...)
				m.loops.Update(depth, loops)
				break
			}
		}
	}
}

func (m *Manager) GetLoopsInModule(module string) []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if loop.Module == module {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) GetNestedLoops(depth int, start uint64) []*Loop {
	result := make([]*Loop, 0)
	loop := m.GetLoop(depth, start)
	if loop != nil {
		for _, child := range loop.Children {
			result = append(result, child)
			result = append(result, m.getNestedLoopsRecursive(child)...)
		}
	}
	return result
}

func (m *Manager) getNestedLoopsRecursive(loop *Loop) []*Loop {
	result := make([]*Loop, 0)
	for _, child := range loop.Children {
		result = append(result, child)
		result = append(result, m.getNestedLoopsRecursive(child)...)
	}
	return result
}

func (m *Manager) GetParentLoop(depth int, start uint64) *Loop {
	if loop := m.GetLoop(depth, start); loop != nil {
		return loop.Parent
	}
	return nil
}

func (m *Manager) GetRootLoops() []*Loop {
	result := make([]*Loop, 0)
	loops, _ := m.loops.Get(0)
	for _, loop := range loops {
		if loop.Parent == nil {
			result = append(result, loop)
		}
	}
	return result
}

func (m *Manager) GetLoopDepth(address uint64) int {
	maxDepth := -1
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if address >= loop.Start && address <= loop.End {
				if depth > maxDepth {
					maxDepth = depth
				}
			}
		}
	}
	return maxDepth
}

func (m *Manager) GetLoopsInRange(start, end uint64) []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if loop.Start >= start && loop.End <= end {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) FindLoopsByType(loopType LoopType) []*Loop {
	result := make([]*Loop, 0)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		for _, loop := range loops {
			if loop.Type == loopType {
				result = append(result, loop)
			}
		}
	}
	return result
}

func (m *Manager) AnalyzeLoopNesting() map[int]int {
	result := make(map[int]int)
	for depth := range m.loops.Range() {
		loops, _ := m.loops.Get(depth)
		result[depth] = len(loops)
	}
	return result
}

func (m *Manager) GetLoopStatistics() map[string]int {
	stats := make(map[string]int)
	stats["total"] = m.GetLoopCount()
	stats["maxDepth"] = m.GetMaxDepth()
	stats["auto"] = len(m.GetAutoLoops())
	stats["user"] = len(m.GetUserLoops())

	for depth := 0; depth <= m.GetMaxDepth(); depth++ {
		stats[fmt.Sprintf("depth_%d", depth)] = m.GetLoopCountByDepth(depth)
	}

	return stats
}

func (m *Manager) UpdateLoopBounds(depth int, start, newStart, newEnd uint64) error {
	if newStart >= newEnd {
		return fmt.Errorf("loop start address must be less than end address")
	}

	if loop := m.GetLoop(depth, start); loop != nil {
		loop.Start = newStart
		loop.End = newEnd
		return nil
	}

	return fmt.Errorf("loop not found at depth %d, start 0x%X", depth, start)
}

func (m *Manager) MergeLoops(depth int, start1, start2 uint64) error {
	loop1 := m.GetLoop(depth, start1)
	loop2 := m.GetLoop(depth, start2)

	if loop1 == nil || loop2 == nil {
		return fmt.Errorf("one or both loops not found")
	}

	if loop1.End != loop2.Start {
		return fmt.Errorf("loops are not adjacent")
	}

	loop1.End = loop2.End
	loop1.InstructionCount += loop2.InstructionCount

	loops, _ := m.loops.Get(depth)
	for i, loop := range loops {
		if loop.Start == start2 {
			loops = append(loops[:i], loops[i+1:]...)
			m.loops.Update(depth, loops)
			break
		}
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
