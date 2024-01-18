package main

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/states"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor"
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/golibrary/widget"
)

func pageCpu(parent *gi.Frame) {
	vSplits := widget.NewVSplits(parent) //有垂直布局的情况需要先垂直，水平的父级是垂直就合理了
	topSplits := widget.NewHSplits(vSplits)
	downSplits := widget.NewHSplits(vSplits)
	vSplits.SetSplits(.7, .3)

	splits := widget.NewVSplits(topSplits) //top is dismTable,bottom is Immediately count the list view window
	dismFrame := gi.NewFrame(splits)
	Immediately := gi.NewFrame(splits)
	gi.NewButton(Immediately) //todo there need a list widget

	dismFrame.Style(func(s *styles.Style) {
		//s.Direction = styles.Row
		s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
	})
	dismTable(dismFrame)
	splits.SetSplits(.9, .1)

	registerFrame := gi.NewFrame(topSplits)
	registerFrame.Style(func(s *styles.Style) {
		//s.Direction = styles.Row
		s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
	})
	newVSplits := widget.NewVSplits(registerFrame) //top is register,bottom is fastCall layout
	giv.NewStructView(newVSplits).SetStruct(&Register{
		RAX:            0,
		RBX:            0x00007FF88500B7F0, //"LdrpInitializeProcess"
		RCX:            0x00007FF884F6F814, //ntdll.00007FF884F6F814
		RDX:            0,
		RBP:            0,
		RSP:            0x000000F4B095EBB0,
		RSI:            0x000000F4B0A89000,
		RDI:            0x00007FF885007C08,
		R8:             0,
		R9:             0,
		R10:            0,
		R11:            0,
		R12:            0,
		R13:            0,
		R14:            0,
		R15:            0,
		RIP:            0x00007FF884FAB785,
		RFLAGS:         0,
		ZF:             0,
		OF:             0,
		CF:             0,
		PF:             0,
		SF:             0,
		TF:             0,
		AF:             0,
		DF:             0,
		IF:             0,
		LastError:      0,
		LastStatus:     0,
		GS:             0,
		ES:             0,
		CS:             0,
		FS:             0,
		DS:             0,
		SS:             0,
		ST0:            0,
		ST1:            0,
		ST2:            0,
		ST3:            0,
		ST4:            0,
		ST5:            0,
		ST6:            0,
		ST7:            0,
		x87TagWord:     0,
		x87ControlWord: 0,
		x87StatusWord:  0,
		x87TW_0:        0,
		x87TW_1:        0,
		x87TW_2:        0,
		x87TW_3:        0,
		x87TW_4:        0,
		x87TW_5:        0,
		x87TW_6:        0,
		x87TW_7:        0,
		x87SW_B:        0,
		x87SW_C3:       0,
		x87SW_TOP:      0,
		x87SW_C2:       0,
		x87SW_C1:       0,
		x87SW_O:        0,
		x87SW_ES:       0,
		x87SW_SF:       0,
		x87SW_P:        0,
		x87SW_U:        0,
		x87SW_Z:        0,
		x87SW_D:        0,
		x87SW_I:        0,
		x87SW_C0:       0,
		x87CW_IC:       0,
		x87CW_RC:       0,
		x87CW_PC:       0,
		x87CW_PM:       0,
		x87CW_UM:       0,
		x87CW_OM:       0,
		x87CW_ZM:       0,
		x87CW_DM:       0,
		x87CW_IM:       0,
		MxCsr:          0,
		MxCsr_FZ:       0,
		MxCsr_PM:       0,
		MxCsr_UM:       0,
		MxCsr_OM:       0,
		MxCsr_ZM:       0,
		MxCsr_IM:       0,
		MxCsr_DM:       0,
		MxCsr_D:        0,
		MxCsr_PE:       0,
		MxCsr_UE:       0,
		MxCsr_OE:       0,
		MxCsr_ZE:       0,
		MxCsr_DE:       0,
		MxCsr_IE:       0,
		MxCsr_RC:       0,
		XMM0:           0,
		XMM1:           0,
		XMM2:           0,
		XMM3:           0,
		XMM4:           0,
		XMM5:           0,
		XMM6:           0,
		XMM7:           0,
		XMM8:           0,
		XMM9:           0,
		XMM10:          0,
		XMM11:          0,
		XMM12:          0,
		XMM13:          0,
		XMM14:          0,
		XMM15:          0,
		YMM0:           0,
		YMM1:           0,
		YMM2:           0,
		YMM3:           0,
		YMM4:           0,
		YMM5:           0,
		YMM6:           0,
		YMM7:           0,
		YMM8:           0,
		YMM9:           0,
		YMM10:          0,
		YMM11:          0,
		YMM12:          0,
		YMM13:          0,
		YMM14:          0,
		YMM15:          0,
		DR0:            0,
		DR1:            0,
		DR2:            0,
		DR3:            0,
		DR6:            0,
		DR7:            0,
	}).SetReadOnly(true)
	fastCallTable(newVSplits)
	newVSplits.SetSplits(.7, .3)

	topSplits.SetSplits(.7, .3)

	tabs := gi.NewTabs(downSplits)
	mem1Tab := tabs.NewTab("memory1")
	hexEditFrame := gi.NewFrame(mem1Tab)
	hexEditFrame.Style(func(s *styles.Style) {
		//s.Direction = styles.Row
		s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
	})
	hexEditEditor := texteditor.NewEditor(hexEditFrame)
	hexEditBuf := texteditor.NewBuf()
	hexEditBuf.SetText([]byte(hex.Dump(testHexDat)))
	hexEditEditor.SetBuf(hexEditBuf)

	tabs.NewTab("memory1")
	tabs.NewTab("memory2")
	tabs.NewTab("memory3")
	tabs.NewTab("memory4")
	tabs.NewTab("memory5")
	tabs.NewTab("watch1")
	tabs.NewTab("var")
	tabs.NewTab("struct")

	stackFrame := gi.NewFrame(downSplits)
	stackTable(stackFrame) //stack
	downSplits.SetSplits(.6, .4)
}

type FastCall struct { //gti:add
	//Index    int
	Register string
	Address  int `format:"%016X"`
	Data     string
}

func fastCallTable(parent ki.Ki) *giv.TableView {
	fastCalls := make([]*FastCall, 0)
	f := new(FastCall)
	for i := 0; i < 5; i++ {
		//1: rcx 00007FF884F6F814 ntdll.00007FF884F6F814
		//2: rdx 0000000000000000 0000000000000000
		//3: r8 0000008F09EFEFB8 0000008F09EFEFB8
		//4: r9 0000000000000000 0000000000000000
		//5: [rsp+28] 0000000000000000 0000000000000000
		switch i { //mock
		case 0:
			f = &FastCall{
				//Index:    i + 1,
				Register: "rcx",
				Address:  0x00007FF884F6F814,
				Data:     "ntdll.00007FF884F6F814",
			}
		case 1:
			f = &FastCall{
				//Index:    i + 1,
				Register: "rdx",
				Address:  0x0000000000000000,
				Data:     "0000000000000000",
			}
		case 2:
			f = &FastCall{
				//Index:    i + 1,
				Register: "r8",
				Address:  0x0000008F09EFEFB8,
				Data:     "0000008F09EFEFB8",
			}
		case 3:
			f = &FastCall{
				//Index:    i + 1,
				Register: "r9",
				Address:  0x0000000000000000,
				Data:     "0000000000000000",
			}
		case 4:
			f = &FastCall{
				//Index:    i + 1,
				Register: "[rsp+28]",
				Address:  0x0000000000000000,
				Data:     "0000000000000000",
			}
		}
		fastCalls = append(fastCalls, f)
	}

	tv := giv.NewTableView(parent, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&fastCalls)
	return tv
}

type Disassembly struct { //gti:add
	Icon        string
	Address     int `format:"%016X"`
	Opcode      []byte
	Instruction string
	Comment     string
}

func dismTable(frame *gi.Frame) *giv.TableView {
	disassemblies := make([]*Disassembly, 100)
	for i := range disassemblies {
		ts := &Disassembly{
			Icon:        "",
			Address:     0x00007FF838E51030 + i,
			Opcode:      []byte{1, 2, 3},
			Instruction: "mov qword ptr ss:[rsp+0x18], rsi",
			Comment:     "comment " + fmt.Sprint(i),
		}
		disassemblies[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&disassemblies)
	return tv
}

type Stack struct { //gti:add
	Address int `format:"%016X"`
	Data    int `format:"%016X"`
	Context string
}

func stackTable(frame *gi.Frame) *giv.TableView {
	stacks := make([]*Stack, 100)
	for i := range stacks {
		ts := &Stack{
			Address: 0x00007FF838E51030 + i,
			Data:    0x00007FF838E51030 + i,
			Context: "返回到 ntdll.RtlGetImageFileMachines+4D9 自 ntdll.RtlAllocateHeap",
		}
		stacks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&stacks)
	return tv
}

type Register struct { //gti:add
	RAX            int `format:"%016X"`
	RBX            int `format:"%016X"`
	RCX            int `format:"%016X"`
	RDX            int `format:"%016X"`
	RBP            int `format:"%016X"`
	RSP            int `format:"%016X"`
	RSI            int `format:"%016X"`
	RDI            int `format:"%016X"`
	R8             int `format:"%016X"`
	R9             int `format:"%016X"`
	R10            int `format:"%016X"`
	R11            int `format:"%016X"`
	R12            int `format:"%016X"`
	R13            int `format:"%016X"`
	R14            int `format:"%016X"`
	R15            int `format:"%016X"`
	RIP            int `format:"%016X"`
	RFLAGS         int `format:"%016X"`
	ZF             int `format:"%016X"`
	OF             int `format:"%016X"`
	CF             int `format:"%016X"`
	PF             int `format:"%016X"`
	SF             int `format:"%016X"`
	TF             int `format:"%016X"`
	AF             int `format:"%016X"`
	DF             int `format:"%016X"`
	IF             int `format:"%016X"`
	LastError      int `format:"%016X"`
	LastStatus     int `format:"%016X"`
	GS             int `format:"%016X"`
	ES             int `format:"%016X"`
	CS             int `format:"%016X"`
	FS             int `format:"%016X"`
	DS             int `format:"%016X"`
	SS             int `format:"%016X"`
	ST0            int `format:"%016X"`
	ST1            int `format:"%016X"`
	ST2            int `format:"%016X"`
	ST3            int `format:"%016X"`
	ST4            int `format:"%016X"`
	ST5            int `format:"%016X"`
	ST6            int `format:"%016X"`
	ST7            int `format:"%016X"`
	x87TagWord     int `format:"%016X"`
	x87ControlWord int `format:"%016X"`
	x87StatusWord  int `format:"%016X"`
	x87TW_0        int `format:"%016X"`
	x87TW_1        int `format:"%016X"`
	x87TW_2        int `format:"%016X"`
	x87TW_3        int `format:"%016X"`
	x87TW_4        int `format:"%016X"`
	x87TW_5        int `format:"%016X"`
	x87TW_6        int `format:"%016X"`
	x87TW_7        int `format:"%016X"`
	x87SW_B        int `format:"%016X"`
	x87SW_C3       int `format:"%016X"`
	x87SW_TOP      int `format:"%016X"`
	x87SW_C2       int `format:"%016X"`
	x87SW_C1       int `format:"%016X"`
	x87SW_O        int `format:"%016X"`
	x87SW_ES       int `format:"%016X"`
	x87SW_SF       int `format:"%016X"`
	x87SW_P        int `format:"%016X"`
	x87SW_U        int `format:"%016X"`
	x87SW_Z        int `format:"%016X"`
	x87SW_D        int `format:"%016X"`
	x87SW_I        int `format:"%016X"`
	x87SW_C0       int `format:"%016X"`
	x87CW_IC       int `format:"%016X"`
	x87CW_RC       int `format:"%016X"`
	x87CW_PC       int `format:"%016X"`
	x87CW_PM       int `format:"%016X"`
	x87CW_UM       int `format:"%016X"`
	x87CW_OM       int `format:"%016X"`
	x87CW_ZM       int `format:"%016X"`
	x87CW_DM       int `format:"%016X"`
	x87CW_IM       int `format:"%016X"`
	MxCsr          int `format:"%016X"`
	MxCsr_FZ       int `format:"%016X"`
	MxCsr_PM       int `format:"%016X"`
	MxCsr_UM       int `format:"%016X"`
	MxCsr_OM       int `format:"%016X"`
	MxCsr_ZM       int `format:"%016X"`
	MxCsr_IM       int `format:"%016X"`
	MxCsr_DM       int `format:"%016X"`
	MxCsr_D        int `format:"%016X"`
	MxCsr_PE       int `format:"%016X"`
	MxCsr_UE       int `format:"%016X"`
	MxCsr_OE       int `format:"%016X"`
	MxCsr_ZE       int `format:"%016X"`
	MxCsr_DE       int `format:"%016X"`
	MxCsr_IE       int `format:"%016X"`
	MxCsr_RC       int `format:"%016X"`
	XMM0           int `format:"%016X"`
	XMM1           int `format:"%016X"`
	XMM2           int `format:"%016X"`
	XMM3           int `format:"%016X"`
	XMM4           int `format:"%016X"`
	XMM5           int `format:"%016X"`
	XMM6           int `format:"%016X"`
	XMM7           int `format:"%016X"`
	XMM8           int `format:"%016X"`
	XMM9           int `format:"%016X"`
	XMM10          int `format:"%016X"`
	XMM11          int `format:"%016X"`
	XMM12          int `format:"%016X"`
	XMM13          int `format:"%016X"`
	XMM14          int `format:"%016X"`
	XMM15          int `format:"%016X"`
	YMM0           int `format:"%016X"`
	YMM1           int `format:"%016X"`
	YMM2           int `format:"%016X"`
	YMM3           int `format:"%016X"`
	YMM4           int `format:"%016X"`
	YMM5           int `format:"%016X"`
	YMM6           int `format:"%016X"`
	YMM7           int `format:"%016X"`
	YMM8           int `format:"%016X"`
	YMM9           int `format:"%016X"`
	YMM10          int `format:"%016X"`
	YMM11          int `format:"%016X"`
	YMM12          int `format:"%016X"`
	YMM13          int `format:"%016X"`
	YMM14          int `format:"%016X"`
	YMM15          int `format:"%016X"`
	DR0            int `format:"%016X"`
	DR1            int `format:"%016X"`
	DR2            int `format:"%016X"`
	DR3            int `format:"%016X"`
	DR6            int `format:"%016X"`
	DR7            int `format:"%016X"`
}

var testHexDat = []byte{
	0x01, 0x00, 0x00, 0x0F, 0xB7, 0x1A, 0xB8, 0x00, 0x02, 0x00, 0x00, 0x41, 0x8B, 0xF9, 0x49, 0x8B,
	0xF0, 0x4C, 0x8B, 0xF1, 0x66, 0x3B, 0xD8, 0x0F, 0x83, 0x93, 0xF8, 0x0A, 0x00, 0x48, 0x8B, 0x52,
	0x08, 0x4C, 0x8D, 0x44, 0x24, 0x50, 0x44, 0x0F, 0xB7, 0xCB, 0xE8, 0xF1, 0x01, 0x00, 0x00, 0x45,
	0x33, 0xFF, 0x85, 0xC0, 0x78, 0x7A, 0x66, 0x44, 0x89, 0xBD, 0x50, 0x01, 0x00, 0x00, 0x85, 0xFF,
	0x0F, 0x85, 0x71, 0xF8, 0x0A, 0x00, 0x48, 0x8D, 0x44, 0x24, 0x50, 0x66, 0x89, 0x5C, 0x24, 0x42,
	0x48, 0x89, 0x44, 0x24, 0x48, 0x48, 0x8D, 0x54, 0x24, 0x40, 0x48, 0x8D, 0x46, 0x28, 0x66, 0x89,
	0x5C, 0x24, 0x40, 0x45, 0x33, 0xC0, 0x48, 0x89, 0x44, 0x24, 0x38, 0x48, 0x8D, 0x4C, 0x24, 0x30,
	0xC7, 0x44, 0x24, 0x30, 0x00, 0x00, 0x00, 0x01, 0xE8, 0x73, 0x76, 0x01, 0x00, 0x85, 0xC0, 0x78,
	0x2F, 0x0F, 0xB7, 0x4C, 0x24, 0x30, 0x48, 0x8B, 0x54, 0x24, 0x38, 0x48, 0x03, 0xCA, 0xEB, 0x0B,
	0x48, 0x8B, 0xC1, 0x48, 0xFF, 0xC9, 0x80, 0x39, 0x5C, 0x74, 0x07, 0x48, 0x3B, 0xCA, 0x77, 0xF0,
	0xEB, 0x03, 0x48, 0x8B, 0xC8, 0x66, 0x2B, 0x4C, 0x24, 0x38, 0x66, 0x89, 0x4E, 0x26, 0x33, 0xC0,
	0x48, 0x8B, 0x8D, 0x60, 0x01, 0x00, 0x00, 0x48, 0x33, 0xCC, 0xE8, 0xB1, 0xD7, 0x08, 0x00, 0x48,
	0x81, 0xC4, 0x70, 0x02, 0x00, 0x00, 0x41, 0x5F, 0x41, 0x5E, 0x41, 0x5C, 0x5F, 0x5E, 0x5B, 0x5D,
	0xC3, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0x71, 0xC8, 0x5F, 0x16, 0x06, 0x2E, 0xAC, 0xBE,
	0x41, 0xB9, 0x08, 0x00, 0x00, 0x00, 0xE9, 0x35, 0x01, 0x00, 0x00, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
	0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0x71, 0x23, 0xD9, 0x70, 0x7F, 0x0E, 0xEC, 0xFA,
	0x48, 0x89, 0x5C, 0x24, 0x08, 0x48, 0x89, 0x6C, 0x24, 0x10, 0x56, 0x57, 0x41, 0x56, 0x48, 0x81,
	0xEC, 0xB0, 0x01, 0x00, 0x00, 0x4D, 0x8B, 0xF0, 0x49, 0x8B, 0xE9, 0x48, 0x8B, 0xDA, 0x4C, 0x8D,
	0x44, 0x24, 0x70, 0x49, 0x8B, 0xD6, 0x41, 0xB9, 0x38, 0x01, 0x00, 0x00, 0x48, 0x8B, 0xF9, 0xE8,
	0xEC, 0x00, 0x00, 0x00, 0x85, 0xC0, 0x0F, 0x88, 0xBB, 0x00, 0x00, 0x00, 0x48, 0x8B, 0x94, 0x24,
	0x08, 0x01, 0x00, 0x00, 0x4C, 0x8D, 0x44, 0x24, 0x20, 0x41, 0xB9, 0x50, 0x00, 0x00, 0x00, 0x48,
	0x8B, 0xCF, 0xE8, 0xC9, 0x00, 0x00, 0x00, 0x85, 0xC0, 0x0F, 0x88, 0x98, 0x00, 0x00, 0x00, 0x48,
	0x8B, 0x84, 0x24, 0xA0, 0x00, 0x00, 0x00, 0xB9, 0xFF, 0xFF, 0x00, 0x00, 0x48, 0x89, 0x43, 0x10,
	0x8B, 0x84, 0x24, 0xB0, 0x00, 0x00, 0x00, 0x89, 0x43, 0x18, 0x8B, 0x84, 0x24, 0xD8, 0x00, 0x00,
	0x00, 0x89, 0x43, 0x1C, 0x8B, 0x44, 0x24, 0x38, 0x3B, 0xC1, 0x0F, 0x87, 0x80, 0x00, 0x00, 0x00,
	0x66, 0x89, 0x43, 0x24, 0xF6, 0x84, 0x24, 0xF0, 0x01, 0x00, 0x00, 0x02, 0x74, 0x43, 0xBE, 0x00,
	0x28, 0x00, 0x00, 0x48, 0x8B, 0xD5, 0xEB, 0x1F, 0x48, 0x8B, 0x94, 0x24, 0xE0, 0x01, 0x00, 0x00,
	0x48, 0x3B, 0xD5, 0x74, 0x2C, 0x66, 0xFF, 0x43, 0x22, 0x48, 0x8D, 0x42, 0xE0, 0x4C, 0x3B, 0xF0,
	0x74, 0x1F, 0x83, 0xC6, 0xFF, 0x74, 0x1A, 0x41, 0xB9, 0x08, 0x00, 0x00, 0x00, 0x4C, 0x8D, 0x84,
	0x24, 0xE0, 0x01, 0x00, 0x00, 0x48, 0x8B, 0xCF, 0xE8, 0x43, 0x00, 0x00, 0x00, 0x85, 0xC0, 0x79,
	0xC7, 0x45, 0x33, 0xC9, 0x48, 0x8D, 0x94, 0x24, 0xB8, 0x00, 0x00, 0x00, 0x4C, 0x8B, 0xC3, 0x48,
	0x8B, 0xCF, 0xE8, 0xE1, 0xFD, 0xFF, 0xFF, 0x4C, 0x8D, 0x9C, 0x24, 0xB0, 0x01, 0x00, 0x00, 0x49,
	0x8B, 0x5B, 0x20, 0x49, 0x8B, 0x6B, 0x28, 0x49, 0x8B, 0xE3, 0x41, 0x5E, 0x5F, 0x5E, 0xC3, 0xCC,
	0x66, 0x89, 0x4B, 0x24, 0xE9, 0x7B, 0xFF, 0xFF, 0xFF, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
	0x40, 0x53, 0x48, 0x83, 0xEC, 0x30, 0x49, 0x8B, 0xD9, 0x48, 0x8B, 0xC1, 0x49, 0xBA, 0x70, 0x58,
	0xD7, 0x50, 0xB7, 0x1E, 0xD1, 0xC5, 0x48, 0x8D, 0x4C, 0x24, 0x40, 0x48, 0x89, 0x4C, 0x24, 0x20,
	0x48, 0x8B, 0x08,
}
