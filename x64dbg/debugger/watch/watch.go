package watch

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type WatchVarType int

const (
	WatchVarTypeUint WatchVarType = iota
	WatchVarTypeInt
	WatchVarTypeFloat
	WatchVarTypeAscii
	WatchVarTypeUnicode
	WatchVarTypeInvalid
)

type WatchdogMode int

const (
	WatchdogModeDisabled WatchdogMode = iota
	WatchdogModeIsTrue
	WatchdogModeIsFalse
	WatchdogModeChanged
	WatchdogModeUnchanged
)

type WatchInfo struct {
	Name              string
	Expression        string
	Window            uint32
	ID                uint32
	VarType           WatchVarType
	WatchdogMode      WatchdogMode
	Value             uint64
	WatchdogTriggered bool
}

type Manager struct {
	watches *safemap.M[uint32, *WatchInfo]
	nextID  uint32
}

func New() api.Interface {
	return &Manager{
		watches: safemap.New[uint32, *WatchInfo](),
		nextID:  1,
	}
}

func (m *Manager) AddWatch(name, expression string, window uint32, varType WatchVarType, watchdogMode WatchdogMode) (uint32, error) {
	if name == "" {
		return 0, fmt.Errorf("watch name cannot be empty")
	}

	if expression == "" {
		return 0, fmt.Errorf("watch expression cannot be empty")
	}

	id := m.nextID
	m.nextID++

	m.watches.Update(id, &WatchInfo{
		Name:              name,
		Expression:        expression,
		Window:            window,
		ID:                id,
		VarType:           varType,
		WatchdogMode:      watchdogMode,
		Value:             0,
		WatchdogTriggered: false,
	})

	return id, nil
}

func (m *Manager) GetWatch(id uint32) *WatchInfo {
	watch, _ := m.watches.Get(id)
	return watch
}

func (m *Manager) GetWatchByName(name string) *WatchInfo {
	for _, watch := range m.watches.Range() {
		if watch.Name == name {
			return watch
		}
	}
	return nil
}

func (m *Manager) DeleteWatch(id uint32) {
	m.watches.Delete(id)
}

func (m *Manager) DeleteWatchByName(name string) {
	for id, watch := range m.watches.Range() {
		if watch.Name == name {
			m.watches.Delete(id)
			break
		}
	}
}

func (m *Manager) GetAllWatches() []*WatchInfo {
	result := make([]*WatchInfo, 0)
	for _, watch := range m.watches.Range() {
		result = append(result, watch)
	}
	return result
}

func (m *Manager) GetWatchesByWindow(window uint32) []*WatchInfo {
	result := make([]*WatchInfo, 0)
	for _, watch := range m.watches.Range() {
		if watch.Window == window {
			result = append(result, watch)
		}
	}
	return result
}

func (m *Manager) HasWatch(id uint32) bool {
	_, exists := m.watches.Get(id)
	return exists
}

func (m *Manager) HasWatchByName(name string) bool {
	return m.GetWatchByName(name) != nil
}

func (m *Manager) UpdateWatchValue(id uint32, value uint64) {
	if watch, exists := m.watches.Get(id); exists {
		oldValue := watch.Value
		watch.Value = value

		switch watch.WatchdogMode {
		case WatchdogModeIsTrue:
			watch.WatchdogTriggered = (value != 0)
		case WatchdogModeIsFalse:
			watch.WatchdogTriggered = (value == 0)
		case WatchdogModeChanged:
			watch.WatchdogTriggered = (value != oldValue)
		case WatchdogModeUnchanged:
			watch.WatchdogTriggered = (value == oldValue)
		}
	}
}

func (m *Manager) SetWatchdogMode(id uint32, mode WatchdogMode) {
	if watch, exists := m.watches.Get(id); exists {
		watch.WatchdogMode = mode
		watch.WatchdogTriggered = false
	}
}

func (m *Manager) SetWatchVarType(id uint32, varType WatchVarType) {
	if watch, exists := m.watches.Get(id); exists {
		watch.VarType = varType
	}
}

func (m *Manager) SetWatchExpression(id uint32, expression string) {
	if watch, exists := m.watches.Get(id); exists {
		watch.Expression = expression
	}
}

func (m *Manager) SetWatchWindow(id uint32, window uint32) {
	if watch, exists := m.watches.Get(id); exists {
		watch.Window = window
	}
}

func (m *Manager) SetWatchName(id uint32, name string) {
	if watch, exists := m.watches.Get(id); exists {
		watch.Name = name
	}
}

func (m *Manager) GetWatchCount() int {
	count := 0
	for range m.watches.Range() {
		count++
	}
	return count
}

func (m *Manager) GetWatchCountByWindow(window uint32) int {
	count := 0
	for _, watch := range m.watches.Range() {
		if watch.Window == window {
			count++
		}
	}
	return count
}

func (m *Manager) GetTriggeredWatches() []*WatchInfo {
	result := make([]*WatchInfo, 0)
	for _, watch := range m.watches.Range() {
		if watch.WatchdogTriggered && watch.WatchdogMode != WatchdogModeDisabled {
			result = append(result, watch)
		}
	}
	return result
}

func (m *Manager) ClearWatchdogTriggered(id uint32) {
	if watch, exists := m.watches.Get(id); exists {
		watch.WatchdogTriggered = false
	}
}

func (m *Manager) ClearAllWatchdogTriggered() {
	for _, watch := range m.watches.Range() {
		watch.WatchdogTriggered = false
	}
}

func (m *Manager) GetWatchesByType(varType WatchVarType) []*WatchInfo {
	result := make([]*WatchInfo, 0)
	for _, watch := range m.watches.Range() {
		if watch.VarType == varType {
			result = append(result, watch)
		}
	}
	return result
}

func (m *Manager) GetWatchesByWatchdogMode(mode WatchdogMode) []*WatchInfo {
	result := make([]*WatchInfo, 0)
	for _, watch := range m.watches.Range() {
		if watch.WatchdogMode == mode {
			result = append(result, watch)
		}
	}
	return result
}

func (m *Manager) EvaluateAllWatches(evaluateFunc func(expression string) (uint64, error)) {
	for _, watch := range m.watches.Range() {
		if value, err := evaluateFunc(watch.Expression); err == nil {
			oldValue := watch.Value
			watch.Value = value

			switch watch.WatchdogMode {
			case WatchdogModeIsTrue:
				watch.WatchdogTriggered = (value != 0)
			case WatchdogModeIsFalse:
				watch.WatchdogTriggered = (value == 0)
			case WatchdogModeChanged:
				watch.WatchdogTriggered = (value != oldValue)
			case WatchdogModeUnchanged:
				watch.WatchdogTriggered = (value == oldValue)
			}
		}
	}
}

func (m *Manager) GetNextID() uint32 {
	return m.nextID
}

func (m *Manager) SetWatchValue(id uint32, value uint64) {
	if watch, exists := m.watches.Get(id); exists {
		watch.Value = value
	}
}

func (m *Manager) GetWatchValue(id uint32) uint64 {
	if watch, exists := m.watches.Get(id); exists {
		return watch.Value
	}
	return 0
}

func (m *Manager) IsWatchdogTriggered(id uint32) bool {
	if watch, exists := m.watches.Get(id); exists {
		return watch.WatchdogTriggered
	}
	return false
}

func (m *Manager) Layout() layout.Widget {
	return nil
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Clear() {
	m.watches.Reset()
	m.nextID = 1
}

func (m *Manager) Self() any {
	return m
}
