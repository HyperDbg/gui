package communication

import (
	"fmt"
	"maps"
	"os"
	"sync"
)

type ForwardingType int

const (
	ForwardingFile ForwardingType = iota
	ForwardingNamedPipe
	ForwardingTCP
	ForwardingModule
)

type ForwardingState int

const (
	ForwardingStateNotOpened ForwardingState = iota
	ForwardingStateOpened
	ForwardingClosed
)

type OutputSourceStatus int

const (
	OutputSourceStatusSuccessfullyOpened OutputSourceStatus = iota
	OutputSourceStatusAlreadyOpened
	OutputSourceStatusAlreadyClosed
	OutputSourceStatusSuccessfullyClosed
	OutputSourceStatusUnknownError
)

type EventForwarding struct {
	Type        ForwardingType
	State       ForwardingState
	Description string
	Handle      *os.File
	Socket      *TcpClient
	ModuleName  string
	mu          sync.RWMutex
}

type ForwardingManager struct {
	sources    map[uint64]*EventForwarding
	currentTag uint64
	mu         sync.RWMutex
}

func NewForwardingManager() *ForwardingManager {
	return &ForwardingManager{
		sources:    make(map[uint64]*EventForwarding),
		currentTag: 0,
	}
}

func (fm *ForwardingManager) GetNewOutputSourceTag() uint64 {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	tag := fm.currentTag
	fm.currentTag++
	return tag
}

func (fm *ForwardingManager) OpenOutputSource(sourceDescriptor *EventForwarding) OutputSourceStatus {
	sourceDescriptor.mu.Lock()
	defer sourceDescriptor.mu.Unlock()

	if sourceDescriptor.State == ForwardingClosed {
		return OutputSourceStatusAlreadyClosed
	}

	if sourceDescriptor.State == ForwardingStateOpened {
		return OutputSourceStatusAlreadyOpened
	}

	sourceDescriptor.State = ForwardingStateOpened

	switch sourceDescriptor.Type {
	case ForwardingFile:
		return OutputSourceStatusSuccessfullyOpened
	case ForwardingNamedPipe:
		return OutputSourceStatusSuccessfullyOpened
	case ForwardingTCP:
		if sourceDescriptor.Socket == nil || !sourceDescriptor.Socket.IsConnected() {
			return OutputSourceStatusUnknownError
		}
		return OutputSourceStatusSuccessfullyOpened
	case ForwardingModule:
		return OutputSourceStatusSuccessfullyOpened
	default:
		return OutputSourceStatusUnknownError
	}
}

func (fm *ForwardingManager) CloseOutputSource(sourceDescriptor *EventForwarding) OutputSourceStatus {
	sourceDescriptor.mu.Lock()
	defer sourceDescriptor.mu.Unlock()

	if sourceDescriptor.State == ForwardingClosed {
		return OutputSourceStatusAlreadyClosed
	}

	if sourceDescriptor.State != ForwardingStateOpened {
		return OutputSourceStatusUnknownError
	}

	sourceDescriptor.State = ForwardingClosed

	switch sourceDescriptor.Type {
	case ForwardingFile:
		if sourceDescriptor.Handle != nil {
			err := sourceDescriptor.Handle.Close()
			if err != nil {
				fmt.Printf("err, failed to close file: %v\n", err)
				return OutputSourceStatusUnknownError
			}
			sourceDescriptor.Handle = nil
		}
		return OutputSourceStatusSuccessfullyClosed

	case ForwardingTCP:
		if sourceDescriptor.Socket != nil {
			err := sourceDescriptor.Socket.ShutdownConnection()
			if err != nil {
				fmt.Printf("err, failed to shutdown TCP connection: %v\n", err)
				return OutputSourceStatusUnknownError
			}
			sourceDescriptor.Socket = nil
		}
		return OutputSourceStatusSuccessfullyClosed

	case ForwardingNamedPipe:
		if sourceDescriptor.Handle != nil {
			err := sourceDescriptor.Handle.Close()
			if err != nil {
				fmt.Printf("err, failed to close named pipe: %v\n", err)
				return OutputSourceStatusUnknownError
			}
			sourceDescriptor.Handle = nil
		}
		return OutputSourceStatusSuccessfullyClosed

	case ForwardingModule:
		sourceDescriptor.ModuleName = ""
		return OutputSourceStatusSuccessfullyClosed

	default:
		return OutputSourceStatusUnknownError
	}
}

func (fm *ForwardingManager) CreateOutputSource(sourceType ForwardingType, description string, socket *TcpClient) (*EventForwarding, error) {
	source := &EventForwarding{
		Type:        sourceType,
		State:       ForwardingStateNotOpened,
		Description: description,
	}

	switch sourceType {
	case ForwardingFile:
		file, err := os.OpenFile(description, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return nil, fmt.Errorf("unable to create file: %w", err)
		}
		source.Handle = file

	case ForwardingNamedPipe:
		return nil, fmt.Errorf("named pipe creation not implemented")

	case ForwardingTCP:
		if socket == nil || !socket.IsConnected() {
			return nil, fmt.Errorf("invalid or unconnected TCP socket")
		}
		source.Socket = socket

	case ForwardingModule:
		source.ModuleName = description

	default:
		return nil, fmt.Errorf("unknown forwarding type")
	}

	tag := fm.GetNewOutputSourceTag()
	fm.mu.Lock()
	fm.sources[tag] = source
	fm.mu.Unlock()

	return source, nil
}

func (fm *ForwardingManager) GetSource(tag uint64) (*EventForwarding, bool) {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	source, exists := fm.sources[tag]
	return source, exists
}

func (fm *ForwardingManager) RemoveSource(tag uint64) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	source, exists := fm.sources[tag]
	if !exists {
		return fmt.Errorf("source with tag %d not found", tag)
	}

	if source.State != ForwardingClosed {
		fm.CloseOutputSource(source)
	}

	delete(fm.sources, tag)
	return nil
}

func (fm *ForwardingManager) GetAllSources() map[uint64]*EventForwarding {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	result := make(map[uint64]*EventForwarding)
	maps.Copy(result, fm.sources)
	return result
}

func (fm *ForwardingManager) CloseAllSources() {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	for _, source := range fm.sources {
		if source.State != ForwardingClosed {
			fm.CloseOutputSource(source)
		}
	}

	fm.sources = make(map[uint64]*EventForwarding)
}
