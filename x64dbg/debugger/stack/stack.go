package stack

import (
	"encoding/binary"
	"fmt"
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"

	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/register"
	"github.com/ddkwork/x64dbg/debugger/symbol"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

type MemoryEntry struct {
	Address    uint64
	Data       uint64
	Context    string
	IsApiOrArg bool
}

type CallStackFrame struct {
	ThreadId   uint32
	Address    uint64
	ReturnTo   uint64
	ReturnFrom uint64
	Size       uint32
	Level      string
	Notes      string
	Valid      bool
}

type Manager struct {
	frames      []*CallStackFrame
	table       *treetable.TreeTable[CallStackFrame]
	memoryTable *treetable.TreeTable[MemoryEntry]
	rsp         uint64
	count       int
}

func New() api.Interface {
	m := &Manager{
		frames: make([]*CallStackFrame, 0),
		rsp:    0,
		count:  16,
	}
	m.initTable()
	m.initMemoryTable()
	return m
}

func (m *Manager) initMemoryTable() {
	m.memoryTable = treetable.NewTreeTable[MemoryEntry]()
	m.memoryTable.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack:    func() {},
		JsonName:               "stack_memory",
		HideToolbar:            true,
		DisableBoxSelect:       true,
		DisableAddColumn:       true,
		DisableAddRow:          true,
		DisableSort:            true,
	}
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[CallStackFrame]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack:    func() {},
		JsonName:               "stack",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.memoryTable.AirTable.Layout
}

func (m *Manager) LayoutMemory() *treetable.TreeTable[MemoryEntry] {
	return m.memoryTable
}

func (m *Manager) GetTable() *treetable.TreeTable[CallStackFrame] {
	return m.table
}

func (m *Manager) SetUpdateParams(rsp uint64, count int) {
	m.rsp = rsp
	m.count = count
}

func (m *Manager) WalkStack(threadHandle windows.Handle, ctx *register.RegisterContext, readMemory func(uint64, uint32) ([]byte, error), getSymbol func(uint64) (*symbol.SymbolInfo, error), getModule func(uint64) *symbol.ModuleInfo) ([]*CallStackFrame, error) {
	frames := make([]*CallStackFrame, 0)
	frameCount := 0

	currentRbp := ctx.RBP
	currentRip := ctx.RIP

	for currentRbp != 0 && frameCount < 100 {
		frameData, err := readMemory(currentRbp, 0x10)
		if err != nil {
			break
		}

		if len(frameData) < 0x10 {
			break
		}

		returnAddress := binary.LittleEndian.Uint64(frameData[0x08:0x10])
		nextRbp := binary.LittleEndian.Uint64(frameData[0x00:0x08])

		frame := &CallStackFrame{
			ThreadId:   0,
			Address:    currentRbp,
			ReturnTo:   returnAddress,
			ReturnFrom: currentRip,
			Size:       uint32(nextRbp - currentRbp),
			Level:      m.getFrameLevel(returnAddress, getModule),
			Notes:      m.getFrameNotes(returnAddress, getSymbol),
			Valid:      true,
		}

		frames = append(frames, frame)

		currentRip = returnAddress
		currentRbp = nextRbp
		frameCount++
	}

	m.frames = frames
	return frames, nil
}

func (m *Manager) GetAllFrames() []*CallStackFrame {
	return m.frames
}

func (m *Manager) GetFrame(address uint64) *CallStackFrame {
	for _, frame := range m.frames {
		if frame.Address == address {
			return frame
		}
	}
	return nil
}

func (m *Manager) getFrameLevel(address uint64, getModule func(uint64) *symbol.ModuleInfo) string {
	if address == 0 {
		return "无效"
	}

	module := getModule(address)
	if module == nil {
		return "未知"
	}

	if module.FileName != "" {
		if module.FileName == "ntdll.dll" || module.FileName == "kernel32.dll" {
			return "系统模块"
		}
		return "用户模块"
	}

	return "未知"
}

func (m *Manager) getFrameNotes(address uint64, getSymbol func(uint64) (*symbol.SymbolInfo, error)) string {
	if address == 0 {
		return ""
	}

	symbol, err := getSymbol(address)
	if err != nil {
		return ""
	}

	if symbol != nil {
		return fmt.Sprintf("%s.%s+0x%X", symbol.ModuleName, symbol.Name, address-symbol.Address)
	}

	return ""
}

func (m *Manager) ReadStackMemory(rsp uint64, count int, readMemory func(uint64, uint32) ([]byte, error)) ([]MemoryEntry, error) {
	entries := make([]MemoryEntry, 0, count)
	for i := uint64(0); i < uint64(count); i++ {
		addr := rsp + i*8
		data, err := readMemory(addr, 8)
		if err != nil {
			break
		}

		if len(data) < 8 {
			break
		}

		var value uint64
		for j := 0; j < 8; j++ {
			value |= uint64(data[j]) << (uint(j) * 8)
		}

		entries = append(entries, MemoryEntry{
			Address: addr,
			Data:    value,
			Context: fmt.Sprintf("%016X", value),
		})
	}

	return entries, nil
}

func (m *Manager) UpdateStack(rsp uint64, count int, readMemory func(uint64, uint32) ([]byte, error)) error {
	stackData, err := m.ReadStackMemory(rsp, count, readMemory)
	if err != nil {
		return err
	}

	if m.memoryTable != nil {
		m.memoryTable.Root().SetChildren(nil)
		for _, entry := range stackData {
			node := m.memoryTable.NewContainerNode(fmt.Sprintf("%016X", entry.Address), entry)
			m.memoryTable.Root().SetChildren(append(m.memoryTable.Root().Children, node))
		}
		m.memoryTable.AirTable.Refresh()
	}

	return nil
}

func (m *Manager) Clear() {
	m.frames = make([]*CallStackFrame, 0)
	if m.table != nil {
		m.table.AirTable.SelectedNode = nil
		m.table.Root().SetChildren(nil)
	}
	if m.memoryTable != nil {
		m.memoryTable.AirTable.SelectedNode = nil
		m.memoryTable.Root().SetChildren(nil)
	}
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) Update() error {
	return nil
}
