package main

import (
	"encoding/hex"
	"fmt"

	"github.com/saferwall/pe"

	"github.com/ddkwork/app/ms/xed"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/richardwilkes/unison"
)

func LayoutCpu(fileName string) unison.Paneler {
	// asm := LayoutDisassemblyTable(fileName)
	// return asm

	////fastCallLayout := unison.NewList[ImmData]()
	//widget.NewButton(m).SetText("goto 00007FF885007C08")
	//"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
	//todo make unit test for fast call layout

	//RegisterView, panel := widget.NewStructView(parent.AsPanel().Window(), Register{}, func(data Register) (values []widget.CellData) {
	//
	//})

	TopHSplit := widget.NewHSplit(
		widget.NewTab("cpu with fast call", "todo fast call layout", true, LayoutDisassemblyTable(fileName)),
		widget.NewTab("reg", "todo reg", true, unison.NewPanel()),
		0.3)
	// parent.AsPanel().AddChild(TopHSplit) //todo bug
	return TopHSplit

	hexEditor := widget.NewCodeEditor("")
	hexEditor.Editor.SetText(hex.Dump(testHexDat))
	BottomHSplit := widget.NewHSplit(
		widget.NewTab(" hex editor", "todo hex editor", true, hexEditor),
		widget.NewTab("stack", "todo stack test", true, LayoutStackTable()),
		0.3)
	//todo add tab into hex editor and stack layout
	/*
		tabs := gi.NewTabs(downSplits)
		mem1Tab := tabs.NewTab("memory1")
		hexEditFrame := gi.NewFrame(mem1Tab)
		hexEditFrame.Style(func(s *styles.Style) {
			// s.Direction = styles.Row
			s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
		})
		hexEditEditor := texteditor.NewEditor(hexEditFrame)
		hexEditBuf := texteditor.NewBuf()
		hexEditBuf.SetText([]byte(hex.Dump(testHexDat)))
		hexEditEditor.SetBuf(hexEditBuf)

		tabs.NewTab("memory2")
		tabs.NewTab("memory3")
		tabs.NewTab("memory4")
		tabs.NewTab("memory5")
		tabs.NewTab("watch1")
		tabs.NewTab("var")
		tabs.NewTab("struct")
	*/

	top := widget.NewTab("cpu and reg", "", true, TopHSplit)
	bottom := widget.NewTab("hex editor and stack", "", true, BottomHSplit)
	vSplit := widget.NewVSplit(top, bottom, 0.1)
	return vSplit
}

var testRegData = Register{
	RAX:            0,
	RBX:            0x00007FF88500B7F0, //"LdrpInitializeProcess"
	RCX:            0x00007FF884F6F814, // ntdll.00007FF884F6F814
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
}

type ( // todo merge into FastCall
	ImmData struct {
		reg     Register
		address uint64
		mem     []byte
	}
)

type FastCall struct {
	// Index    int
	Register string
	Address  int
	ImmData  string
}

func LayoutDisassemblyTable(fileName string) unison.Paneler {
	table, header := widget.NewTable(xed.Disassembly{}, widget.TableContext[xed.Disassembly]{
		ContextMenuItems: func(node *widget.Node[xed.Disassembly]) []widget.ContextMenuItem {
			return []widget.ContextMenuItem{
				{
					Title: "goto",
					Can:   func(any) bool { return true },
					Do:    func(a any) { mylog.Todo("goto 0x00007FF838E51030") },
				},
			}
		},
		MarshalRow: func(node *widget.Node[xed.Disassembly]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("%X", node.Data.Opcode)},
				{Text: node.Data.Instruction},
				{Text: node.Data.Comment},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[xed.Disassembly]) {
			f := xed.ParserPe(fileName)
			b := stream.NewBuffer(fileName)
			optionalHeader := f.NtHeader.OptionalHeader
			switch o := optionalHeader.(type) {
			case pe.ImageOptionalHeader32:
				oep := o.ImageBase + o.AddressOfEntryPoint
				x := xed.New(b.Bytes()[oep:])
				x.Decode64()
				for _, object := range x.AsmObjects {
					root.AddChildByData(object)
				}
			case pe.ImageOptionalHeader64:
				oep := o.ImageBase + uint64(o.AddressOfEntryPoint)
				mylog.Struct(o)
				mylog.Hex("oep", oep)
				x := xed.New(b.Bytes()[o.AddressOfEntryPoint:]) // todo bug
				// x := xed.New(b.Bytes()[o.AddressOfEntryPoint:])
				x.Decode64()
				for _, object := range x.AsmObjects {
					root.AddChildByData(object)
				}
			}
			// hyperdbg-cli.exe size is 2mb
			// now 1-2s need
			// todo show iat table,dump overlay

			//x := xed.New(buf)
			//if f.Is64 {
			//	x.Decode64()
			//} else {
			//	x.Decode32()
			//}
		},
		JsonName:   "cpu_dism_table",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
}

type Stack struct {
	Address int
	Data    int
	Context string
}

func LayoutStackTable() unison.Paneler {
	table, header := widget.NewTable(Stack{}, widget.TableContext[Stack]{
		ContextMenuItems: nil, // todo goto 0x00007FF838E51030
		MarshalRow: func(node *widget.Node[Stack]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("%016X", node.Data.Data)},
				{Text: node.Data.Context},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Stack]) {
			for i := range 100 {
				ts := Stack{
					Address: 0x00007FF838E51030 + i,
					Data:    0x00007FF838E51030 + i,
					Context: "返回到 ntdll.RtlGetImageFileMachines+4D9 自 ntdll.RtlAllocateHeap",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "cpu_stack_table",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
}

type Register struct {
	RAX            int
	RBX            int
	RCX            int
	RDX            int
	RBP            int
	RSP            int
	RSI            int
	RDI            int
	R8             int
	R9             int
	R10            int
	R11            int
	R12            int
	R13            int
	R14            int
	R15            int
	RIP            int
	RFLAGS         int
	ZF             int
	OF             int
	CF             int
	PF             int
	SF             int
	TF             int
	AF             int
	DF             int
	IF             int
	LastError      int
	LastStatus     int
	GS             int
	ES             int
	CS             int
	FS             int
	DS             int
	SS             int
	ST0            int
	ST1            int
	ST2            int
	ST3            int
	ST4            int
	ST5            int
	ST6            int
	ST7            int
	x87TagWord     int
	x87ControlWord int
	x87StatusWord  int
	x87TW_0        int
	x87TW_1        int
	x87TW_2        int
	x87TW_3        int
	x87TW_4        int
	x87TW_5        int
	x87TW_6        int
	x87TW_7        int
	x87SW_B        int
	x87SW_C3       int
	x87SW_TOP      int
	x87SW_C2       int
	x87SW_C1       int
	x87SW_O        int
	x87SW_ES       int
	x87SW_SF       int
	x87SW_P        int
	x87SW_U        int
	x87SW_Z        int
	x87SW_D        int
	x87SW_I        int
	x87SW_C0       int
	x87CW_IC       int
	x87CW_RC       int
	x87CW_PC       int
	x87CW_PM       int
	x87CW_UM       int
	x87CW_OM       int
	x87CW_ZM       int
	x87CW_DM       int
	x87CW_IM       int
	MxCsr          int
	MxCsr_FZ       int
	MxCsr_PM       int
	MxCsr_UM       int
	MxCsr_OM       int
	MxCsr_ZM       int
	MxCsr_IM       int
	MxCsr_DM       int
	MxCsr_D        int
	MxCsr_PE       int
	MxCsr_UE       int
	MxCsr_OE       int
	MxCsr_ZE       int
	MxCsr_DE       int
	MxCsr_IE       int
	MxCsr_RC       int
	XMM0           int
	XMM1           int
	XMM2           int
	XMM3           int
	XMM4           int
	XMM5           int
	XMM6           int
	XMM7           int
	XMM8           int
	XMM9           int
	XMM10          int
	XMM11          int
	XMM12          int
	XMM13          int
	XMM14          int
	XMM15          int
	YMM0           int
	YMM1           int
	YMM2           int
	YMM3           int
	YMM4           int
	YMM5           int
	YMM6           int
	YMM7           int
	YMM8           int
	YMM9           int
	YMM10          int
	YMM11          int
	YMM12          int
	YMM13          int
	YMM14          int
	YMM15          int
	DR0            int
	DR1            int
	DR2            int
	DR3            int
	DR6            int
	DR7            int
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
}
