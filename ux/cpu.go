package ux

import (
	"encoding/hex"
	"fmt"

	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutCpu(parent unison.Paneler) unison.Paneler {
	TopHSplit := widget.NewHSplit(
		widget.NewTab("cpu with fast call", "todo fast call layout", true, LayoutDismTable(parent)),
		widget.NewTab("reg", "todo reg", true, unison.NewPanel()),
		0.7)

	hexEditor := unison.NewField()
	hexEditor.SetText(hex.Dump(testHexDat))
	BottomHSplit := widget.NewHSplit(
		widget.NewTab(" hex editor", "todo hex editor", true, hexEditor),
		widget.NewTab("stack", "todo stack test", true, LayoutStackTable(parent)),
		0.6)

	top := widget.NewTab("cpu and reg", "", true, TopHSplit)
	bottom := widget.NewTab("hex editor and stack", "", true, BottomHSplit)
	vSplit := widget.NewVSplit(top, bottom, 0.7)
	return vSplit
	/*
		splits := widget.NewVSplits(topSplits) // top is dismTable,bottom is Immediately count the list view window
		dismFrame := gi.NewFrame(splits)
		Immediately := gi.NewFrame(splits)
		v := giv.NewSliceView(Immediately)
		v.AddContextMenu(func(m *gi.Scene) {
			widget.NewButton(m).SetText("goto 00007FF885007C08")
			widget.NewButton(m).SetText("goto 00007FF885007C08")
			widget.NewButton(m).SetText("goto 00007FF885007C08")
			widget.NewButton(m).SetText("goto 00007FF885007C08")
			widget.NewButton(m).SetText("goto 00007FF885007C08")
		})
		v.SetSlice([]string{
			"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
			"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
			"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
			"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
		}).SetReadOnly(true)
		// widget.NewButton(Immediately) //todo there need a list widget

		dismFrame.Style(func(s *styles.Style) {
			// s.Direction = styles.Row
			s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
		})
		dismTable(dismFrame)
		splits.SetSplits(.85, .15)

		registerFrame := gi.NewFrame(topSplits)
		registerFrame.Style(func(s *styles.Style) {
			// s.Direction = styles.Row
			s.Background = colors.C(colors.Scheme.SurfaceContainerLow)
		})
		newVSplits := widget.NewVSplits(registerFrame) // top is register,bottom is fastCall layout
		giv.NewStructView(newVSplits).SetStruct(&Register{
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
		}).SetReadOnly(true)
		// fastCallTable(newVSplits)
		fastCall := giv.NewSliceView(newVSplits)
		fastCall.SetReadOnly(true)
		fastCall.AddContextMenu(func(m *gi.Scene) {
			widget.NewButton(m).SetText("go 00007FF884F6F814")
			widget.NewButton(m).SetText("go 00007FF884F6F814")
			widget.NewButton(m).SetText("go 00007FF884F6F814")
		})
		fastCall.SetSlice([]string{
			"rcx 00007FF884F6F814 ntdll.00007FF884F6F814", // 1:
			"rdx 0000000000000000 0000000000000000",       // 2:
			"r8  0000008F09EFEFB8 0000008F09EFEFB8",       // 3:
			"r9  0000000000000000 0000000000000000",       // 4:
			"[rsp+28] 0000000000000000 0000000000000000",  // 5:
		})
		newVSplits.SetSplits(.7, .3)

		topSplits.SetSplits(.7, .3)

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

		stackFrame := gi.NewFrame(downSplits)
		stackTable(stackFrame) // stack
		downSplits.SetSplits(.6, .4)
	*/
}

//type FastCall struct {
//	//Index    int
//	Register string
//	Address  int
//	MetaData     string
//}

//func fastCallTable(parent ki.Ki) *giv.TableView {
//	fastCalls := make([]*FastCall, 0)
//	f := new(FastCall)
//	for i := 0; i < 5; i++ {
//		//1: rcx 00007FF884F6F814 ntdll.00007FF884F6F814
//		//2: rdx 0000000000000000 0000000000000000
//		//3: r8 0000008F09EFEFB8 0000008F09EFEFB8
//		//4: r9 0000000000000000 0000000000000000
//		//5: [rsp+28] 0000000000000000 0000000000000000
//		switch i { //mock
//		case 0:
//			f = &FastCall{
//				//Index:    i + 1,
//				Register: "rcx",
//				Address:  0x00007FF884F6F814,
//				MetaData:     "ntdll.00007FF884F6F814",
//			}
//		case 1:
//			f = &FastCall{
//				//Index:    i + 1,
//				Register: "rdx",
//				Address:  0x0000000000000000,
//				MetaData:     "0000000000000000",
//			}
//		case 2:
//			f = &FastCall{
//				//Index:    i + 1,
//				Register: "r8",
//				Address:  0x0000008F09EFEFB8,
//				MetaData:     "0000008F09EFEFB8",
//			}
//		case 3:
//			f = &FastCall{
//				//Index:    i + 1,
//				Register: "r9",
//				Address:  0x0000000000000000,
//				MetaData:     "0000000000000000",
//			}
//		case 4:
//			f = &FastCall{
//				//Index:    i + 1,
//				Register: "[rsp+28]",
//				Address:  0x0000000000000000,
//				MetaData:     "0000000000000000",
//			}
//		}
//		fastCalls = append(fastCalls, f)
//	}
//
//	tv := giv.NewTableView(parent, "tv")
//	tv.SetReadOnly(true)
//	tv.SetSlice(&fastCalls)
//	return tv
//}

type Disassembly struct {
	Icon        string
	Address     int
	Opcode      []byte
	Instruction string
	Comment     string
}

func LayoutDismTable(parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(Disassembly{}, widget.TableContext[Disassembly]{
		ContextMenuItems: func(node *widget.Node[Disassembly]) []widget.ContextMenuItem {
			return []widget.ContextMenuItem{
				{
					Title: "goto",
					Can:   func(any) bool { return true },
					Do:    func(a any) { mylog.Todo("goto 0x00007FF838E51030") },
				},
			}
		},
		MarshalRow: func(node *widget.Node[Disassembly]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: node.Data.Icon},
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("% X", node.Data.Opcode)},
				{Text: node.Data.Instruction},
				{Text: node.Data.Comment},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Disassembly]) {
			for i := range 100 {
				ts := Disassembly{
					Icon:        "",
					Address:     0x00007FF838E51030 + i,
					Opcode:      []byte{1, 2, 3},
					Instruction: "mov qword ptr ss:[rsp+0x18], rsi",
					Comment:     "comment " + fmt.Sprint(i),
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "cpu_dism_table",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}

type Stack struct {
	Address int
	Data    int
	Context string
}

func LayoutStackTable(parent unison.Paneler) unison.Paneler {
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
	return widget.NewTableScrollPanel(parent, table, header)
}

type Register struct {
	RAX            int `width:"50" format:"%016X"`
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
