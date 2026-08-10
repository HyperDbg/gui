package register

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/structview"
	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

type RegisterContext struct {
	RAX            uint64
	RBX            uint64
	RCX            uint64
	RDX            uint64
	RBP            uint64
	RSP            uint64
	RSI            uint64
	RDI            uint64
	R8             uint64
	R9             uint64
	R10            uint64
	R11            uint64
	R12            uint64
	R13            uint64
	R14            uint64
	R15            uint64
	RIP            uint64
	RFLAGS         uint32
	ZF             bool
	OF             bool
	CF             bool
	PF             bool
	SF             bool
	TF             bool
	AF             bool
	DF             bool
	IF             bool
	LastError      uint32
	LastStatus     uint32
	GS             uint16
	ES             uint16
	CS             uint16
	FS             uint16
	DS             uint16
	SS             uint16
	ST0            uint64
	ST1            uint64
	ST2            uint64
	ST3            uint64
	ST4            uint64
	ST5            uint64
	ST6            uint64
	ST7            uint64
	X87TagWord     uint16
	X87ControlWord uint16
	X87StatusWord  uint16
	MxCsr          uint32
	XMM0           [16]byte
	XMM1           [16]byte
	XMM2           [16]byte
	XMM3           [16]byte
	XMM4           [16]byte
	XMM5           [16]byte
	XMM6           [16]byte
	XMM7           [16]byte
	XMM8           [16]byte
	XMM9           [16]byte
	XMM10          [16]byte
	XMM11          [16]byte
	XMM12          [16]byte
	XMM13          [16]byte
	XMM14          [16]byte
	XMM15          [16]byte
	YMM0           [32]byte
	YMM1           [32]byte
	YMM2           [32]byte
	YMM3           [32]byte
	YMM4           [32]byte
	YMM5           [32]byte
	YMM6           [32]byte
	YMM7           [32]byte
	YMM8           [32]byte
	YMM9           [32]byte
	YMM10          [32]byte
	YMM11          [32]byte
	YMM12          [32]byte
	YMM13          [32]byte
	YMM14          [32]byte
	YMM15          [32]byte
	DR0            uint64
	DR1            uint64
	DR2            uint64
	DR3            uint64
	DR6            uint64
	DR7            uint64
}

type Manager struct {
	threadContexts *safemap.M[windows.Handle, *RegisterContext]
	threadHandles  *safemap.M[uint32, windows.Handle]
	table          *structview.StructView
}

func New() api.Interface {
	m := &Manager{
		threadContexts: safemap.New[windows.Handle, *RegisterContext](),
		threadHandles:  safemap.New[uint32, windows.Handle](),
	}
	m.initTable()
	return m
}

func (m *Manager) GetThreadContext(threadHandle windows.Handle) (*RegisterContext, error) {
	ctx, err := windows.GetThreadContext(threadHandle)
	if err != nil {
		return nil, err
	}

	regCtx := &RegisterContext{
		RAX:    ctx.Rax,
		RBX:    ctx.Rbx,
		RCX:    ctx.Rcx,
		RDX:    ctx.Rdx,
		RBP:    ctx.Rbp,
		RSP:    ctx.Rsp,
		RSI:    ctx.Rsi,
		RDI:    ctx.Rdi,
		R8:     ctx.R8,
		R9:     ctx.R9,
		R10:    ctx.R10,
		R11:    ctx.R11,
		R12:    ctx.R12,
		R13:    ctx.R13,
		R14:    ctx.R14,
		R15:    ctx.R15,
		RIP:    ctx.Rip,
		RFLAGS: uint32(ctx.EFlags),
		GS:     ctx.SegGs,
		ES:     ctx.SegEs,
		CS:     ctx.SegCs,
		FS:     ctx.SegFs,
		DS:     ctx.SegDs,
		SS:     ctx.SegSs,
		DR0:    ctx.Dr0,
		DR1:    ctx.Dr1,
		DR2:    ctx.Dr2,
		DR3:    ctx.Dr3,
		DR6:    ctx.Dr6,
		DR7:    ctx.Dr7,
		MxCsr:  uint32(ctx.MxCsr),
	}

	m.parseFlags(regCtx)
	m.parseXMMRegisters(ctx, regCtx)

	m.threadContexts.Update(threadHandle, regCtx)
	return regCtx, nil
}

func (m *Manager) SetThreadHandle(threadId uint32, threadHandle windows.Handle) {
	m.threadHandles.Update(threadId, threadHandle)
}

func (m *Manager) GetThreadHandle(threadId uint32) windows.Handle {
	return m.threadHandles.GetMust(threadId)
}

func (m *Manager) SetThreadContext(threadHandle windows.Handle, regCtx *RegisterContext) error {
	ctx := &windows.Context{
		ContextFlags: windows.CONTEXT_FULL | windows.CONTEXT_DEBUG_REGISTERS,
		Rax:          regCtx.RAX,
		Rbx:          regCtx.RBX,
		Rcx:          regCtx.RCX,
		Rdx:          regCtx.RDX,
		Rbp:          regCtx.RBP,
		Rsp:          regCtx.RSP,
		Rsi:          regCtx.RSI,
		Rdi:          regCtx.RDI,
		R8:           regCtx.R8,
		R9:           regCtx.R9,
		R10:          regCtx.R10,
		R11:          regCtx.R11,
		R12:          regCtx.R12,
		R13:          regCtx.R13,
		R14:          regCtx.R14,
		R15:          regCtx.R15,
		Rip:          regCtx.RIP,
		EFlags:       uint32(regCtx.RFLAGS),
		SegGs:        regCtx.GS,
		SegEs:        regCtx.ES,
		SegCs:        regCtx.CS,
		SegFs:        regCtx.FS,
		SegDs:        regCtx.DS,
		SegSs:        regCtx.SS,
		Dr0:          regCtx.DR0,
		Dr1:          regCtx.DR1,
		Dr2:          regCtx.DR2,
		Dr3:          regCtx.DR3,
		Dr6:          regCtx.DR6,
		Dr7:          regCtx.DR7,
		MxCsr:        regCtx.MxCsr,
	}

	m.setFlags(regCtx, ctx)
	m.setXMMRegisters(regCtx, ctx)

	err := windows.SetThreadContext(threadHandle, ctx)
	if err != nil {
		return err
	}

	m.threadContexts.Update(threadHandle, regCtx)
	return nil
}

func (m *Manager) Clear() {
	m.threadContexts.Reset()
	m.table.SetData(RegisterContext{})
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) initTable() {
	m.table = structview.New(RegisterContext{}, *button.NewButtonsLayout(button.LayoutButton{Button: button.Text(), Label: ""}))
}

func (m *Manager) Layout() layout.Widget {
	return m.table.Layout
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) UpdateFromThreadHandle(threadHandle windows.Handle) error {
	regCtx, err := m.GetThreadContext(threadHandle)
	if err != nil {
		return err
	}
	m.table.SetData(*regCtx)
	return nil
}

func (m *Manager) GetRegister(threadHandle windows.Handle, name string) (uint64, error) {
	regCtx, exists := m.threadContexts.Get(threadHandle)
	if !exists {
		return 0, fmt.Errorf("thread context not found")
	}

	switch name {
	case "RAX":
		return regCtx.RAX, nil
	case "RBX":
		return regCtx.RBX, nil
	case "RCX":
		return regCtx.RCX, nil
	case "RDX":
		return regCtx.RDX, nil
	case "RBP":
		return regCtx.RBP, nil
	case "RSP":
		return regCtx.RSP, nil
	case "RSI":
		return regCtx.RSI, nil
	case "RDI":
		return regCtx.RDI, nil
	case "R8":
		return regCtx.R8, nil
	case "R9":
		return regCtx.R9, nil
	case "R10":
		return regCtx.R10, nil
	case "R11":
		return regCtx.R11, nil
	case "R12":
		return regCtx.R12, nil
	case "R13":
		return regCtx.R13, nil
	case "R14":
		return regCtx.R14, nil
	case "R15":
		return regCtx.R15, nil
	case "RIP":
		return regCtx.RIP, nil
	case "RFLAGS":
		return uint64(regCtx.RFLAGS), nil
	case "DR0":
		return regCtx.DR0, nil
	case "DR1":
		return regCtx.DR1, nil
	case "DR2":
		return regCtx.DR2, nil
	case "DR3":
		return regCtx.DR3, nil
	case "DR6":
		return regCtx.DR6, nil
	case "DR7":
		return regCtx.DR7, nil
	default:
		return 0, fmt.Errorf("unknown register: %s", name)
	}
}

func (m *Manager) SetRegister(threadHandle windows.Handle, name string, value uint64) error {
	regCtx, exists := m.threadContexts.Get(threadHandle)
	if !exists {
		return fmt.Errorf("thread context not found")
	}

	switch name {
	case "RAX":
		regCtx.RAX = value
	case "RBX":
		regCtx.RBX = value
	case "RCX":
		regCtx.RCX = value
	case "RDX":
		regCtx.RDX = value
	case "RBP":
		regCtx.RBP = value
	case "RSP":
		regCtx.RSP = value
	case "RSI":
		regCtx.RSI = value
	case "RDI":
		regCtx.RDI = value
	case "R8":
		regCtx.R8 = value
	case "R9":
		regCtx.R9 = value
	case "R10":
		regCtx.R10 = value
	case "R11":
		regCtx.R11 = value
	case "R12":
		regCtx.R12 = value
	case "R13":
		regCtx.R13 = value
	case "R14":
		regCtx.R14 = value
	case "R15":
		regCtx.R15 = value
	case "RIP":
		regCtx.RIP = value
	case "RFLAGS":
		regCtx.RFLAGS = uint32(value)
	case "DR0":
		regCtx.DR0 = value
	case "DR1":
		regCtx.DR1 = value
	case "DR2":
		regCtx.DR2 = value
	case "DR3":
		regCtx.DR3 = value
	case "DR6":
		regCtx.DR6 = value
	case "DR7":
		regCtx.DR7 = value
	default:
		return fmt.Errorf("unknown register: %s", name)
	}

	ctx := &windows.Context{
		ContextFlags: windows.CONTEXT_FULL | windows.CONTEXT_DEBUG_REGISTERS,
		Rax:          regCtx.RAX,
		Rbx:          regCtx.RBX,
		Rcx:          regCtx.RCX,
		Rdx:          regCtx.RDX,
		Rbp:          regCtx.RBP,
		Rsp:          regCtx.RSP,
		Rsi:          regCtx.RSI,
		Rdi:          regCtx.RDI,
		R8:           regCtx.R8,
		R9:           regCtx.R9,
		R10:          regCtx.R10,
		R11:          regCtx.R11,
		R12:          regCtx.R12,
		R13:          regCtx.R13,
		R14:          regCtx.R14,
		R15:          regCtx.R15,
		Rip:          regCtx.RIP,
		EFlags:       uint32(regCtx.RFLAGS),
		SegGs:        regCtx.GS,
		SegEs:        regCtx.ES,
		SegCs:        regCtx.CS,
		SegFs:        regCtx.FS,
		SegDs:        regCtx.DS,
		SegSs:        regCtx.SS,
		Dr0:          regCtx.DR0,
		Dr1:          regCtx.DR1,
		Dr2:          regCtx.DR2,
		Dr3:          regCtx.DR3,
		Dr6:          regCtx.DR6,
		Dr7:          regCtx.DR7,
		MxCsr:        regCtx.MxCsr,
	}

	m.setFlags(regCtx, ctx)
	m.setXMMRegisters(regCtx, ctx)

	return windows.SetThreadContext(threadHandle, ctx)
}

func (m *Manager) GetContext(threadHandle windows.Handle) (*RegisterContext, error) {
	regCtx, exists := m.threadContexts.Get(threadHandle)
	if !exists {
		return nil, fmt.Errorf("thread context not found")
	}

	return regCtx, nil
}

func (m *Manager) ClearContext(threadHandle windows.Handle) {
	m.threadContexts.Delete(threadHandle)
}

func (m *Manager) parseFlags(regCtx *RegisterContext) {
	flags := regCtx.RFLAGS
	regCtx.ZF = (flags & (1 << 6)) != 0
	regCtx.OF = (flags & (1 << 11)) != 0
	regCtx.CF = (flags & 1) != 0
	regCtx.PF = (flags & (1 << 2)) != 0
	regCtx.SF = (flags & (1 << 7)) != 0
	regCtx.TF = (flags & (1 << 8)) != 0
	regCtx.AF = (flags & (1 << 4)) != 0
	regCtx.DF = (flags & (1 << 10)) != 0
	regCtx.IF = (flags & (1 << 9)) != 0
}

func (m *Manager) setFlags(regCtx *RegisterContext, ctx *windows.Context) {
	flags := uint32(0)
	if regCtx.ZF {
		flags |= (1 << 6)
	}
	if regCtx.OF {
		flags |= (1 << 11)
	}
	if regCtx.CF {
		flags |= 1
	}
	if regCtx.PF {
		flags |= (1 << 2)
	}
	if regCtx.SF {
		flags |= (1 << 7)
	}
	if regCtx.TF {
		flags |= (1 << 8)
	}
	if regCtx.AF {
		flags |= (1 << 4)
	}
	if regCtx.DF {
		flags |= (1 << 10)
	}
	if regCtx.IF {
		flags |= (1 << 9)
	}
	ctx.EFlags = flags
}

func (m *Manager) parseXMMRegisters(ctx *windows.Context, regCtx *RegisterContext) {
	xmmArrays := [][16]byte{
		regCtx.XMM0, regCtx.XMM1, regCtx.XMM2, regCtx.XMM3,
		regCtx.XMM4, regCtx.XMM5, regCtx.XMM6, regCtx.XMM7,
		regCtx.XMM8, regCtx.XMM9, regCtx.XMM10, regCtx.XMM11,
		regCtx.XMM12, regCtx.XMM13, regCtx.XMM14, regCtx.XMM15,
	}

	for i := range 16 {
		if i < len(ctx.VectorRegister) && i < len(xmmArrays) {
			vec := ctx.VectorRegister[i]
			for j := 0; j < 16 && j < 8; j++ {
				xmmArrays[i][j] = byte(vec.Low >> (j * 8))
				xmmArrays[i][j+8] = byte(uint64(vec.High) >> (j * 8))
			}
		}
	}
}

func (m *Manager) setXMMRegisters(regCtx *RegisterContext, ctx *windows.Context) {
	xmmArrays := [][16]byte{
		regCtx.XMM0, regCtx.XMM1, regCtx.XMM2, regCtx.XMM3,
		regCtx.XMM4, regCtx.XMM5, regCtx.XMM6, regCtx.XMM7,
		regCtx.XMM8, regCtx.XMM9, regCtx.XMM10, regCtx.XMM11,
		regCtx.XMM12, regCtx.XMM13, regCtx.XMM14, regCtx.XMM15,
	}

	for i := 0; i < 16 && i < len(ctx.VectorRegister) && i < len(xmmArrays); i++ {
		vec := &ctx.VectorRegister[i]
		vec.Low = 0
		vec.High = 0
		for j := 0; j < 16 && j < 8; j++ {
			vec.Low |= uint64(xmmArrays[i][j]) << (j * 8)
			vec.High |= int64(xmmArrays[i][j+8]) << (j * 8)
		}
	}
}

func (m *Manager) GetAllRegisters() map[string]uint64 {
	result := make(map[string]uint64)

	for _, regCtx := range m.threadContexts.Range() {
		result["RAX"] = regCtx.RAX
		result["RBX"] = regCtx.RBX
		result["RCX"] = regCtx.RCX
		result["RDX"] = regCtx.RDX
		result["RBP"] = regCtx.RBP
		result["RSP"] = regCtx.RSP
		result["RSI"] = regCtx.RSI
		result["RDI"] = regCtx.RDI
		result["R8"] = regCtx.R8
		result["R9"] = regCtx.R9
		result["R10"] = regCtx.R10
		result["R11"] = regCtx.R11
		result["R12"] = regCtx.R12
		result["R13"] = regCtx.R13
		result["R14"] = regCtx.R14
		result["R15"] = regCtx.R15
		result["RIP"] = regCtx.RIP
		result["RFLAGS"] = uint64(regCtx.RFLAGS)
		break
	}

	return result
}

func (m *Manager) GetRegisterContext() *RegisterContext {
	for _, regCtx := range m.threadContexts.Range() {
		return regCtx
	}
	return &RegisterContext{}
}
