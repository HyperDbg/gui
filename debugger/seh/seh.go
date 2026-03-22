package seh

import (
	"encoding/binary"
	"fmt"
	"iter"
	"sync"

	"github.com/ddkwork/golibrary/std/mylog"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
)

type ExceptionHandler struct {
	Address                   uint64
	ExceptionHandlingRoutines string
	Label                     string
	Notes                     string
	Valid                     bool
}

type Manager struct {
	handlers []*ExceptionHandler
	mu       sync.RWMutex
	table    *treetable.TreeTable[ExceptionHandler]
}

func New() api.Interface {
	m := &Manager{
		handlers: make([]*ExceptionHandler, 0),
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[ExceptionHandler]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack:    func() {},
		JsonName:               "Seh",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) GetTable() *treetable.TreeTable[ExceptionHandler] {
	return m.table
}

func (m *Manager) ScanSEH(ctx *register.RegisterContext, readMemory func(uint64, uint32) ([]byte, error)) ([]*ExceptionHandler, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	tebAddress := uint64(ctx.GS)
	if tebAddress == 0 {
		tebAddress = uint64(ctx.FS)
	}

	if tebAddress == 0 {
		return nil, fmt.Errorf("failed to get thread information block base")
	}

	tebData := mylog.Check2(readMemory(tebAddress, 0x30))

	sehListHead := binary.LittleEndian.Uint64(tebData[0x00:0x08])

	handlers := make([]*ExceptionHandler, 0)
	currentSeh := sehListHead

	for currentSeh != 0xFFFFFFFFFFFFFFFF {
		if len(handlers) > 1000 {
			break
		}

		sehData := mylog.Check2(readMemory(currentSeh, 0x10))

		nextSeh := binary.LittleEndian.Uint64(sehData[0x00:0x08])
		handler := binary.LittleEndian.Uint64(sehData[0x08:0x10])

		if handler == 0 {
			break
		}

		handlerInfo := &ExceptionHandler{
			Address:                   currentSeh,
			ExceptionHandlingRoutines: fmt.Sprintf("0x%X", handler),
			Label:                     "",
			Notes:                     "",
			Valid:                     true,
		}

		handlers = append(handlers, handlerInfo)

		currentSeh = nextSeh
	}

	m.handlers = handlers
	return handlers, nil
}

func (m *Manager) GetAllHandlers() []*ExceptionHandler {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.handlers
}

func (m *Manager) GetHandler(address uint64) *ExceptionHandler {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, handler := range m.handlers {
		if handler.Address == address {
			return handler
		}
	}
	return nil
}
