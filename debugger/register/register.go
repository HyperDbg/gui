package register

import (
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/structview"
)

type RegisterContext struct {
	RAX    uint64
	RBX    uint64
	RCX    uint64
	RDX    uint64
	RBP    uint64
	RSP    uint64
	RSI    uint64
	RDI    uint64
	R8     uint64
	R9     uint64
	R10    uint64
	R11    uint64
	R12    uint64
	R13    uint64
	R14    uint64
	R15    uint64
	RIP    uint64
	RFLAGS uint32
	ZF     bool
	OF     bool
	CF     bool
	PF     bool
	SF     bool
	TF     bool
	AF     bool
	DF     bool
	IF     bool
	GS     uint16
	ES     uint16
	CS     uint16
	FS     uint16
	DS     uint16
	SS     uint16
	DR0    uint64
	DR1    uint64
	DR2    uint64
	DR3    uint64
	DR6    uint64
	DR7    uint64
}

type Provider struct {
	mu             sync.RWMutex
	threadContexts *safemap.M[uint32, *RegisterContext]
	currentContext *RegisterContext
	table          *structview.StructView
}

func New() api.Interface {
	m := &Provider{
		threadContexts: safemap.New[uint32, *RegisterContext](),
		currentContext: &RegisterContext{},
	}
	m.initTable()
	return m
}

func (m *Provider) initTable() {
	m.table = structview.New(RegisterContext{}, *button.NewButtonsLayout(button.LayoutButton{Button: button.Text(), Label: ""}))
}

func (m *Provider) Self() any {
	return m
}

func (m *Provider) Layout() layout.Widget {
	return m.table.Layout
}

func (m *Provider) Update() error {
	m.table.SetData(*m.currentContext)
	return nil
}

func (m *Provider) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.threadContexts.Reset()
	m.currentContext = &RegisterContext{}
	m.table.SetData(RegisterContext{})
}

func (m *Provider) GetRegisterContext() *RegisterContext {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.currentContext
}

func (m *Provider) SetRegisterContext(ctx *RegisterContext) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentContext = ctx
	m.table.SetData(*ctx)
}

func (m *Provider) GetThreadContext(threadId uint32) *RegisterContext {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.threadContexts.GetMust(threadId)
}

func (m *Provider) SetThreadContext(threadId uint32, ctx *RegisterContext) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.threadContexts.Update(threadId, ctx)
}

func (m *Provider) GetAllRegisters() map[string]uint64 {
	result := make(map[string]uint64)
	ctx := m.currentContext
	if ctx == nil {
		return result
	}
	result["RAX"] = ctx.RAX
	result["RBX"] = ctx.RBX
	result["RCX"] = ctx.RCX
	result["RDX"] = ctx.RDX
	result["RBP"] = ctx.RBP
	result["RSP"] = ctx.RSP
	result["RSI"] = ctx.RSI
	result["RDI"] = ctx.RDI
	result["R8"] = ctx.R8
	result["R9"] = ctx.R9
	result["R10"] = ctx.R10
	result["R11"] = ctx.R11
	result["R12"] = ctx.R12
	result["R13"] = ctx.R13
	result["R14"] = ctx.R14
	result["R15"] = ctx.R15
	result["RIP"] = ctx.RIP
	result["RFLAGS"] = uint64(ctx.RFLAGS)
	result["DR0"] = ctx.DR0
	result["DR1"] = ctx.DR1
	result["DR2"] = ctx.DR2
	result["DR3"] = ctx.DR3
	result["DR6"] = ctx.DR6
	result["DR7"] = ctx.DR7
	return result
}
