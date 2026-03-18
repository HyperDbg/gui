package handle

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type HandleType int

const (
	HandleTypeFile HandleType = iota
	HandleTypeEvent
	HandleTypeMutex
	HandleTypeSemaphore
	HandleTypeThread
	HandleTypeProcess
	HandleTypeSection
	HandleTypeKey
	HandleTypePort
	HandleTypeWindowStation
	HandleTypeDesktop
	HandleTypeDirectory
	HandleTypeToken
	HandleTypeJob
	HandleTypeOther
)

type Handle struct {
	Value         uint64
	Type          HandleType
	Name          string
	ObjectType    string
	Attributes    uint32
	GrantedAccess uint32
	ProcessId     uint32
	ThreadId      uint32
}

type Manager struct {
	handles *safemap.M[uint64, *Handle]
}

func New() api.Interface {
	return &Manager{
		handles: safemap.New[uint64, *Handle](),
	}
}

func (m *Manager) AddHandle(value uint64, handleType HandleType, name, objectType string, attributes, grantedAccess uint32, processId, threadId uint32) error {
	if value == 0 {
		return fmt.Errorf("handle value cannot be zero")
	}

	m.handles.Update(value, &Handle{
		Value:         value,
		Type:          handleType,
		Name:          name,
		ObjectType:    objectType,
		Attributes:    attributes,
		GrantedAccess: grantedAccess,
		ProcessId:     processId,
		ThreadId:      threadId,
	})

	return nil
}

func (m *Manager) GetHandle(value uint64) *Handle {
	handle, _ := m.handles.Get(value)
	return handle
}

func (m *Manager) DeleteHandle(value uint64) {
	m.handles.Delete(value)
}

func (m *Manager) GetAllHandles() []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		result = append(result, handle)
	}
	return result
}

func (m *Manager) GetHandlesByType(handleType HandleType) []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		if handle.Type == handleType {
			result = append(result, handle)
		}
	}
	return result
}

func (m *Manager) GetHandlesByProcess(processId uint32) []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		if handle.ProcessId == processId {
			result = append(result, handle)
		}
	}
	return result
}

func (m *Manager) GetHandlesByThread(threadId uint32) []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		if handle.ThreadId == threadId {
			result = append(result, handle)
		}
	}
	return result
}

func (m *Manager) GetHandlesByName(name string) []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		if handle.Name == name {
			result = append(result, handle)
		}
	}
	return result
}

func (m *Manager) GetHandlesByObjectType(objectType string) []*Handle {
	result := make([]*Handle, 0)
	for _, handle := range m.handles.Range() {
		if handle.ObjectType == objectType {
			result = append(result, handle)
		}
	}
	return result
}

func (m *Manager) UpdateHandle(value uint64, name string, attributes, grantedAccess uint32) error {
	if handle, exists := m.handles.Get(value); exists {
		handle.Name = name
		handle.Attributes = attributes
		handle.GrantedAccess = grantedAccess
		return nil
	}
	return fmt.Errorf("handle not found: 0x%X", value)
}

func (m *Manager) HasHandle(value uint64) bool {
	_, exists := m.handles.Get(value)
	return exists
}

func (m *Manager) Clear() {
	m.handles.Reset()
}

func (m *Manager) GetHandleCount() int {
	count := 0
	for range m.handles.Range() {
		count++
	}
	return count
}

func (m *Manager) GetHandleCountByType(handleType HandleType) int {
	count := 0
	for _, handle := range m.handles.Range() {
		if handle.Type == handleType {
			count++
		}
	}
	return count
}

func (m *Manager) GetHandleCountByProcess(processId uint32) int {
	count := 0
	for _, handle := range m.handles.Range() {
		if handle.ProcessId == processId {
			count++
		}
	}
	return count
}

func (m *Manager) ClearByProcess(processId uint32) {
	for value, handle := range m.handles.Range() {
		if handle.ProcessId == processId {
			m.handles.Delete(value)
		}
	}
}

func (m *Manager) ClearByThread(threadId uint32) {
	for value, handle := range m.handles.Range() {
		if handle.ThreadId == threadId {
			m.handles.Delete(value)
		}
	}
}

func (m *Manager) GetHandleStatistics() map[string]int {
	stats := make(map[string]int)
	stats["total"] = m.GetHandleCount()
	stats["file"] = m.GetHandleCountByType(HandleTypeFile)
	stats["event"] = m.GetHandleCountByType(HandleTypeEvent)
	stats["mutex"] = m.GetHandleCountByType(HandleTypeMutex)
	stats["semaphore"] = m.GetHandleCountByType(HandleTypeSemaphore)
	stats["thread"] = m.GetHandleCountByType(HandleTypeThread)
	stats["process"] = m.GetHandleCountByType(HandleTypeProcess)
	stats["section"] = m.GetHandleCountByType(HandleTypeSection)
	stats["key"] = m.GetHandleCountByType(HandleTypeKey)
	stats["port"] = m.GetHandleCountByType(HandleTypePort)
	stats["other"] = m.GetHandleCountByType(HandleTypeOther)
	return stats
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
