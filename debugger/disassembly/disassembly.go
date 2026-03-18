package disassembly

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/treetable"
)

type Instruction struct {
	Address    uint64
	Bytes      []byte
	Mnemonic   string
	Operands   string
	Comment    string
	Size       int
	IsBranch   bool
	IsCall     bool
	IsRet      bool
	IsJump     bool
	BranchType string
	Target     uint64
	Module     string
	Function   string
}

type DisassemblyBlock struct {
	StartAddress uint64
	EndAddress   uint64
	Instructions []*Instruction
}

type Provider struct {
	mu       sync.RWMutex
	blocks   *safemap.M[uint64, *DisassemblyBlock]
	cache    map[uint64]*Instruction
	current  *DisassemblyBlock
	baseAddr uint64
	table    *treetable.TreeTable[Instruction]
}

func New() api.Interface {
	m := &Provider{
		blocks: safemap.New[uint64, *DisassemblyBlock](),
		cache:  make(map[uint64]*Instruction),
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = treetable.NewTreeTable[Instruction]()
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Provider) Update() error {
	m.table.Root().SetChildren(nil)
	for _, instr := range m.cache {
		m.table.Root().AddChild(m.table.NewNode(*instr))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Provider) AddInstruction(instr *Instruction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache[instr.Address] = instr
}

func (m *Provider) GetInstruction(address uint64) *Instruction {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.cache[address]
}

func (m *Provider) GetInstructions(start, end uint64) []*Instruction {
	m.mu.RLock()
	defer m.mu.RUnlock()

	instrs := make([]*Instruction, 0)
	for addr := start; addr <= end; {
		if instr, ok := m.cache[addr]; ok {
			instrs = append(instrs, instr)
			addr += uint64(instr.Size)
		} else {
			break
		}
	}
	return instrs
}

func (m *Provider) AddBlock(block *DisassemblyBlock) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.blocks.Update(block.StartAddress, block)

	for _, instr := range block.Instructions {
		m.cache[instr.Address] = instr
	}
}

func (m *Provider) GetBlock(address uint64) *DisassemblyBlock {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.blocks.GetMust(address)
}

func (m *Provider) SetCurrentBlock(block *DisassemblyBlock) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.current = block
}

func (m *Provider) GetCurrentBlock() *DisassemblyBlock {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.current
}

func (m *Provider) SetBaseAddress(addr uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.baseAddr = addr
}

func (m *Provider) GetBaseAddress() uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.baseAddr
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.blocks.Reset()
	m.cache = make(map[uint64]*Instruction)
	m.current = nil
	m.baseAddr = 0
}
