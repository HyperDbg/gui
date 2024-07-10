package ux

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ddkwork/HyperDbg/sdk"

	"github.com/saferwall/pe"

	"github.com/ddkwork/app/ms/xed"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutCpu(fileName string) unison.Paneler {
	////fastCallLayout := unison.NewList[ImmData]()
	//widget.NewButton(m).SetText("goto 00007FF885007C08")
	//"rdi=00007FF885007C08 \"minkernel\\\\ntdll\\\\ldrinit.c\"",
	//todo make unit test for fast call layout

	asm := LayoutDisassemblyTable(fileName)

	structView, _ := widget.NewStructView(testRegData, func(data Register) (values []widget.CellData) {
		return []widget.CellData{
			{Text: fmt.Sprintf("%016x", data.RAX)},
			{Text: fmt.Sprintf("%016x", data.RBX)},
			{Text: fmt.Sprintf("%016x", data.RCX)},
			{Text: fmt.Sprintf("%016x", data.RDX)},
			{Text: fmt.Sprintf("%016x", data.RBP)},
			{Text: fmt.Sprintf("%016x", data.RSP)},
			{Text: fmt.Sprintf("%016x", data.RSI)},
			{Text: fmt.Sprintf("%016x", data.RDI)},
			{Text: fmt.Sprintf("%016x", data.R8)},
			{Text: fmt.Sprintf("%016x", data.R9)},
			{Text: fmt.Sprintf("%016x", data.R10)},
			{Text: fmt.Sprintf("%016x", data.R11)},
			{Text: fmt.Sprintf("%016x", data.R12)},
			{Text: fmt.Sprintf("%016x", data.R13)},
			{Text: fmt.Sprintf("%016x", data.R14)},
			{Text: fmt.Sprintf("%016x", data.R15)},
			{Text: fmt.Sprintf("%016x", data.RIP)},
			{Text: fmt.Sprintf("%016x", data.RFLAGS)},
			{Text: fmt.Sprintf("%016x", data.ZF)},
			{Text: fmt.Sprintf("%016x", data.OF)},
			{Text: fmt.Sprintf("%016x", data.CF)},
			{Text: fmt.Sprintf("%016x", data.PF)},
			{Text: fmt.Sprintf("%016x", data.SF)},
			{Text: fmt.Sprintf("%016x", data.TF)},
			{Text: fmt.Sprintf("%016x", data.AF)},
			{Text: fmt.Sprintf("%016x", data.DF)},
			{Text: fmt.Sprintf("%016x", data.IF)},
			{Text: fmt.Sprintf("%016x", data.LastError)},
			{Text: fmt.Sprintf("%016x", data.LastStatus)},
			{Text: fmt.Sprintf("%016x", data.GS)},
			{Text: fmt.Sprintf("%016x", data.ES)},
			{Text: fmt.Sprintf("%016x", data.CS)},
			{Text: fmt.Sprintf("%016x", data.FS)},
			{Text: fmt.Sprintf("%016x", data.DS)},
			{Text: fmt.Sprintf("%016x", data.SS)},
			{Text: fmt.Sprintf("%016x", data.ST0)},
			{Text: fmt.Sprintf("%016x", data.ST1)},
			{Text: fmt.Sprintf("%016x", data.ST2)},
			{Text: fmt.Sprintf("%016x", data.ST3)},
			{Text: fmt.Sprintf("%016x", data.ST4)},
			{Text: fmt.Sprintf("%016x", data.ST5)},
			{Text: fmt.Sprintf("%016x", data.ST6)},
			{Text: fmt.Sprintf("%016x", data.ST7)},
			{Text: fmt.Sprintf("%016x", data.X87TagWord)},
			{Text: fmt.Sprintf("%016x", data.X87ControlWord)},
			{Text: fmt.Sprintf("%016x", data.X87StatusWord)},
			{Text: fmt.Sprintf("%016x", data.X87TW_0)},
			{Text: fmt.Sprintf("%016x", data.X87TW_1)},
			{Text: fmt.Sprintf("%016x", data.X87TW_2)},
			{Text: fmt.Sprintf("%016x", data.X87TW_3)},
			{Text: fmt.Sprintf("%016x", data.X87TW_4)},
			{Text: fmt.Sprintf("%016x", data.X87TW_5)},
			{Text: fmt.Sprintf("%016x", data.X87TW_6)},
			{Text: fmt.Sprintf("%016x", data.X87TW_7)},
			{Text: fmt.Sprintf("%016x", data.X87SW_B)},
			{Text: fmt.Sprintf("%016x", data.X87SW_C3)},
			{Text: fmt.Sprintf("%016x", data.X87SW_TOP)},
			{Text: fmt.Sprintf("%016x", data.X87SW_C2)},
			{Text: fmt.Sprintf("%016x", data.X87SW_C1)},
			{Text: fmt.Sprintf("%016x", data.X87SW_O)},
			{Text: fmt.Sprintf("%016x", data.X87SW_ES)},
			{Text: fmt.Sprintf("%016x", data.X87SW_SF)},
			{Text: fmt.Sprintf("%016x", data.X87SW_P)},
			{Text: fmt.Sprintf("%016x", data.X87SW_U)},
			{Text: fmt.Sprintf("%016x", data.X87SW_Z)},
			{Text: fmt.Sprintf("%016x", data.X87SW_D)},
			{Text: fmt.Sprintf("%016x", data.X87SW_I)},
			{Text: fmt.Sprintf("%016x", data.X87SW_C0)},
			{Text: fmt.Sprintf("%016x", data.X87CW_IC)},
			{Text: fmt.Sprintf("%016x", data.X87CW_RC)},
			{Text: fmt.Sprintf("%016x", data.X87CW_PC)},
			{Text: fmt.Sprintf("%016x", data.X87CW_PM)},
			{Text: fmt.Sprintf("%016x", data.X87CW_UM)},
			{Text: fmt.Sprintf("%016x", data.X87CW_OM)},
			{Text: fmt.Sprintf("%016x", data.X87CW_ZM)},
			{Text: fmt.Sprintf("%016x", data.X87CW_DM)},
			{Text: fmt.Sprintf("%016x", data.X87CW_IM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_FZ)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_PM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_UM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_OM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_ZM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_IM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_DM)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_D)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_PE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_UE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_OE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_ZE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_DE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_IE)},
			{Text: fmt.Sprintf("%016x", data.MxCsr_RC)},
			{Text: fmt.Sprintf("%016x", data.XMM0)},
			{Text: fmt.Sprintf("%016x", data.XMM1)},
			{Text: fmt.Sprintf("%016x", data.XMM2)},
			{Text: fmt.Sprintf("%016x", data.XMM3)},
			{Text: fmt.Sprintf("%016x", data.XMM4)},
			{Text: fmt.Sprintf("%016x", data.XMM5)},
			{Text: fmt.Sprintf("%016x", data.XMM6)},
			{Text: fmt.Sprintf("%016x", data.XMM7)},
			{Text: fmt.Sprintf("%016x", data.XMM8)},
			{Text: fmt.Sprintf("%016x", data.XMM9)},
			{Text: fmt.Sprintf("%016x", data.XMM10)},
			{Text: fmt.Sprintf("%016x", data.XMM11)},
			{Text: fmt.Sprintf("%016x", data.XMM12)},
			{Text: fmt.Sprintf("%016x", data.XMM13)},
			{Text: fmt.Sprintf("%016x", data.XMM14)},
			{Text: fmt.Sprintf("%016x", data.XMM15)},
			{Text: fmt.Sprintf("%016x", data.YMM0)},
			{Text: fmt.Sprintf("%016x", data.YMM1)},
			{Text: fmt.Sprintf("%016x", data.YMM2)},
			{Text: fmt.Sprintf("%016x", data.YMM3)},
			{Text: fmt.Sprintf("%016x", data.YMM4)},
			{Text: fmt.Sprintf("%016x", data.YMM5)},
			{Text: fmt.Sprintf("%016x", data.YMM6)},
			{Text: fmt.Sprintf("%016x", data.YMM7)},
			{Text: fmt.Sprintf("%016x", data.YMM8)},
			{Text: fmt.Sprintf("%016x", data.YMM9)},
			{Text: fmt.Sprintf("%016x", data.YMM10)},
			{Text: fmt.Sprintf("%016x", data.YMM11)},
			{Text: fmt.Sprintf("%016x", data.YMM12)},
			{Text: fmt.Sprintf("%016x", data.YMM13)},
			{Text: fmt.Sprintf("%016x", data.YMM14)},
			{Text: fmt.Sprintf("%016x", data.YMM15)},
			{Text: fmt.Sprintf("%016x", data.DR0)},
			{Text: fmt.Sprintf("%016x", data.DR1)},
			{Text: fmt.Sprintf("%016x", data.DR2)},
			{Text: fmt.Sprintf("%016x", data.DR3)},
			{Text: fmt.Sprintf("%016x", data.DR6)},
			{Text: fmt.Sprintf("%016x", data.DR7)},
		}
	})
	TopHSplit := widget.NewHSplit(
		widget.NewTab("cpu with fast call", "todo fast call layout", asm),
		widget.NewTab("Register", "Register", widget.NewScrollPanelFill(structView)),
		0.3)

	hexEditor := widget.NewCodeEditor("")
	hexEditor.Editor.SetText(hex.Dump(testHexDat))
	stackTable := LayoutStackTable()
	HexEditTab := widget.NewTab(" hex editor", "todo hex editor", hexEditor)
	// dockContainer := widget.NewDockContainer(hexEditor)//todo
	BottomHSplit := widget.NewHSplit(
		HexEditTab,
		widget.NewTab("stack", "todo stack test", stackTable),
		0.1)
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

	top := widget.NewTab("cpu and reg", "", TopHSplit)
	bottom := widget.NewTab("hex editor and stack", "", BottomHSplit)
	vSplit := widget.NewVSplit(top, bottom, 0.1)
	return vSplit.Dock
}

// todo D:\desk\go\golang.org\x\arch@v0.8.0\x86\x86asm\inst.go :L150
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
	X87TagWord:     0,
	X87ControlWord: 0,
	X87StatusWord:  0,
	X87TW_0:        0,
	X87TW_1:        0,
	X87TW_2:        0,
	X87TW_3:        0,
	X87TW_4:        0,
	X87TW_5:        0,
	X87TW_6:        0,
	X87TW_7:        0,
	X87SW_B:        0,
	X87SW_C3:       0,
	X87SW_TOP:      0,
	X87SW_C2:       0,
	X87SW_C1:       0,
	X87SW_O:        0,
	X87SW_ES:       0,
	X87SW_SF:       0,
	X87SW_P:        0,
	X87SW_U:        0,
	X87SW_Z:        0,
	X87SW_D:        0,
	X87SW_I:        0,
	X87SW_C0:       0,
	X87CW_IC:       0,
	X87CW_RC:       0,
	X87CW_PC:       0,
	X87CW_PM:       0,
	X87CW_UM:       0,
	X87CW_OM:       0,
	X87CW_ZM:       0,
	X87CW_DM:       0,
	X87CW_IM:       0,
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
	SetRootRowsCallBack := func(root *widget.Node[xed.Disassembly]) {
		f := xed.ParserPe(fileName)
		// b := stream.NewBuffer(fileName)
		optionalHeader := f.NtHeader.OptionalHeader
		switch o := optionalHeader.(type) {
		case pe.ImageOptionalHeader32:
			oepRVA := o.AddressOfEntryPoint
			imageBase := o.ImageBase
			oepVA := imageBase + oepRVA
			var oepFileOffset uint64

			//for _, section := range f.Sections {
			//	if section.String() == ".text" {
			//		oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
			//		break
			//	}
			//}

			//for _, section := range f.Sections {
			//	if section.Header.Characteristics&section.Header.Characteristics == pe.ImageSectionMemExecute {
			//		oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
			//		break
			//	}
			//}
			for _, section := range f.Sections {
				mylog.Info("PrettySectionFlags", section.PrettySectionFlags())
				for _, s := range section.PrettySectionFlags() {
					if s == "Executable" {
						oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
						break
					}
				}
			}

			if oepFileOffset == 0 {
				fmt.Println("未找到包含OEP节区或计算偏移不正确")
				return
			}
			buffer := make([]byte, 200)
			file := mylog.Check2(os.Open(fileName))
			defer file.Close()
			mylog.Check2(file.ReadAt(buffer, int64(oepFileOffset)))
			fmt.Printf("OEP File Off %x\n", oepFileOffset)
			fmt.Printf("OEP VA %x\n", oepVA)
			fmt.Printf("Entry Point RVA %x\n", oepRVA)
			fmt.Printf("OEP Data %x\n", buffer[:100])

			x := xed.New(buffer[:100])
			x.SetBaseAddress(uint64(oepVA))
			x.Decode32()
			for _, object := range x.AsmObjects {
				root.AddChildByData(object)
			}
		case pe.ImageOptionalHeader64:
			oepRVA := o.AddressOfEntryPoint
			imageBase := o.ImageBase
			oepVA := imageBase + uint64(oepRVA)
			var oepFileOffset uint64
			for _, section := range f.Sections {
				if section.String() == ".text" {
					oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
					break
				}
			}
			if oepFileOffset == 0 {
				fmt.Println("未找到包含OEP节区或计算偏移不正确")
				return
			}
			buffer := make([]byte, 200)
			file := mylog.Check2(os.Open(fileName))
			defer file.Close()
			mylog.Check2(file.ReadAt(buffer, int64(oepFileOffset)))
			fmt.Printf("OEP File Off %x\n", oepFileOffset)
			fmt.Printf("OEP VA %x\n", oepVA)
			fmt.Printf("Entry Point RVA %x\n", oepRVA)
			fmt.Printf("OEP Data %x\n", buffer[:100])

			x := xed.New(buffer[:100])
			x.SetBaseAddress(oepVA)
			x.Decode64()
			for _, object := range x.AsmObjects {
				root.AddChildByData(object)
			}
			//oep_rva := uint64(o.AddressOfEntryPoint)
			//image_base := uint64(o.ImageBase)
			//oep_va := image_base + oep_rva
			//var oep_file_offset uint64
			//
			//// 查找 .text 节区
			//for _, section := range f.Sections {
			//	if section.String() == ".text" {
			//		oep_file_offset = uint64(section.Header.PointerToRawData) + (oep_rva - uint64(section.Header.VirtualAddress))
			//		mylog.Struct(section)
			//		break
			//	}
			//}
			//
			//if oep_file_offset == 0 {
			//	mylog.Check("未找到 .text 节区信息或计算偏移不正确")
			//}
			//
			//// 确定文件中的偏移位置
			//buffer := stream.NewBuffer(fileName)
			//if oep_file_offset+200 > uint64(len(buffer.Bytes())) {
			//	mylog.Check("计算出的文件偏移超出文件大小")
			//}
			//
			//// 确认调试信息：
			//mylog.Hex("oep_file_offset", oep_file_offset)
			//mylog.Hex("oep_va", oep_va)
			//mylog.Hex("entry_point_rva", oep_rva)
			//
			//// 读取指定文件偏移二进制数据：
			//oep_data := buffer.Bytes()[oep_file_offset : oep_file_offset+200]
			//
			//fmt.Printf("oep_data: %x\n", oep_data[:20])
			//// 使用反汇编工具解析数据
			//x := xed.New(oep_data)
			//x.SetBaseAddress(oep_va)
			//x.Decode64()
			//for _, object := range x.AsmObjects {
			//	root.AddChildByData(object)
			//}

		}
	}

	table, header := widget.NewTable(xed.Disassembly{}, widget.TableContext[xed.Disassembly]{
		ContextMenuItems: func(node *widget.Node[xed.Disassembly]) []widget.ContextMenuItem {
			return []widget.ContextMenuItem{
				{
					Title: "copy address",
					Can:   func(any) bool { return true },
					Do:    func(a any) { mylog.Todo("copy address 0x00007FF838E51030") },
				},
				{
					Title: "copy row",
					Can:   func(any) bool { return true },
					Do:    func(a any) { mylog.Todo("copy row 0x00007FF838E51030") },
				},
				{
					Title: "goto",
					Can:   func(any) bool { return true },
					Do:    func(a any) { mylog.Todo("goto 0x00007FF838E51030") },
				},
				{
					Title: "set breakpoint",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SetBreakpoint0xcc", sdk.SetBreakpoint0xcc()) // todo use SetBreakpoint
					},
				},
				{
					Title: "EptHook",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						// todo make a args panel ?
						mylog.Warning("EptHook", sdk.EptHook()) // todo make args with type
					},
				},
				{
					Title: "EptHook2",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("EptHook2", sdk.EptHook2()) // todo make args with type
					},
				},
				/*
					BreakpointClearByID
						BreakpointDisableByID
						BreakpointEnableByID
						BreakpointList
						SetBreakpoint0xcc
						CoreOperatingProcessorForShowsAndChanges
						CpuFeaturesForCollectsReport
						FlushKernelModeBuffers
						ContinueDebugger
						StepOut
						StepIn
						CallstackOrThreadView //跟随当前地址的堆栈，同时刷新右下角的堆栈panel
						ListModules
						LoadDriversAndModules
						OutputEventForwardingInstance
						StepExecuteSingleInstruction
						PauseKernelEvents
						PreactivateSpecialFunctionality
						PreallocateBuffer
						PrintEvaluateExpressions
						RegistersReadOrModify
						ReadMsr
						SearchMemoryPattern
						SettingsManagement
						SleepDebugger
						StepInExecute
						TestHyperDbgFeatures
						UnloadKernelModules
						SearchSymbols
						CpuidExecutionMonitor
						ControlRegisterModificationMonitor
						DebugRegistersMonitor
						EptHook
						EptHook2
						IdtEntriesMonitor
						HideHyperDbg
						InterruptExternalMonitor
						IoInDetect
						IoOutDetect
						MeasureArgumentsForHide
						ModeInstructionsTrap
						MonitorMemoryAccess
						MsrRead
						MsrWrite
						Pa2Va
						PmcExecutionMonitor
						Pte
						ReversingMachineModuleUse
						Syscall
						SysRet
						TraceExecution
						TrackModeTransitionInstructions
						TscInstructionsMonitor
						UnHide
						Va2Pa
						VmCall
						HardwareClockDebugging
						AttachDebugThread
						ClearScreen
						ConnectToMachine
						DebugMachine
						DetachDebugging
						DisconnectSession
						DumpMemoryContext
						FormatsDiff
						HelpForCommand
						KillProcess
						ListenForClientConnection
						LogClose
						LogOpen
						PageAvailableInRam
						ParseExecutableFiles
						ProcessesView
						RestartProcess
						Script
						StartProcess
						Status
						SwitchThread
						Symbol
						SymbolPath
						Thread
				*/
				{
					Title: "BreakpointClearByID",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("BreakpointClearByID", sdk.BreakpointClearByID())
					},
				},
				{
					Title: "BreakpointDisableByID",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("BreakpointDisableByID", sdk.BreakpointDisableByID())
					},
				},
				{
					Title: "BreakpointEnableByID",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("BreakpointEnableByID", sdk.BreakpointEnableByID())
					},
				},
				{
					Title: "BreakpointList",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("BreakpointList", sdk.BreakpointList())
					},
				},
				{
					Title: "CoreOperatingProcessorForShowsAndChanges",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("CoreOperatingProcessorForShowsAndChanges", sdk.CoreOperatingProcessorForShowsAndChanges())
					},
				},
				{
					Title: "CpuFeaturesForCollectsReport",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("CpuFeaturesForCollectsReport", sdk.CpuFeaturesForCollectsReport())
					},
				},
				{
					Title: "FlushKernelModeBuffers",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("FlushKernelModeBuffers", sdk.FlushKernelModeBuffers())
					},
				},
				{
					Title: "ContinueDebugger",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ContinueDebugger", sdk.ContinueDebugger())
					},
				},
				{
					Title: "StepOut",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("StepOut", sdk.StepOut())
					},
				},
				{
					Title: "StepIn",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("StepIn", sdk.StepIn())
					},
				},
				{
					Title: "CallstackOrThreadView",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("CallstackOrThreadView", sdk.CallstackOrThreadView())
					},
				},
				{
					Title: "ListModules",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ListModules", sdk.ListModules())
					},
				},
				{
					Title: "LoadDriversAndModules",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("LoadDriversAndModules", sdk.LoadDriversAndModules())
					},
				},
				{
					Title: "OutputEventForwardingInstance",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("OutputEventForwardingInstance", sdk.OutputEventForwardingInstance())
					},
				},
				{
					Title: "StepExecuteSingleInstruction",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("StepExecuteSingleInstruction", sdk.StepExecuteSingleInstruction())
					},
				},
				{
					Title: "PauseKernelEvents",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PauseKernelEvents", sdk.PauseKernelEvents())
					},
				},
				{
					Title: "PreactivateSpecialFunctionality",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PreactivateSpecialFunctionality", sdk.PreactivateSpecialFunctionality())
					},
				},
				{
					Title: "PreallocateBuffer",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PreallocateBuffer", sdk.PreallocateBuffer())
					},
				},
				{
					Title: "PrintEvaluateExpressions",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PrintEvaluateExpressions", sdk.PrintEvaluateExpressions())
					},
				},
				{
					Title: "RegistersReadOrModify",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("RegistersReadOrModify", sdk.RegistersReadOrModify())
					},
				},
				{
					Title: "ReadMsr",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ReadMsr", sdk.ReadMsr())
					},
				},
				{
					Title: "SearchMemoryPattern",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SearchMemoryPattern", sdk.SearchMemoryPattern())
					},
				},
				{
					Title: "SettingsManagement",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SettingsManagement", sdk.SettingsManagement())
					},
				},
				{
					Title: "SleepDebugger",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SleepDebugger", sdk.SleepDebugger())
					},
				},
				{
					Title: "StepInExecute",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("StepInExecute", sdk.StepInExecute())
					},
				},
				{
					Title: "TestHyperDbgFeatures",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("TestHyperDbgFeatures", sdk.TestHyperDbgFeatures())
					},
				},
				{
					Title: "UnloadKernelModules",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("UnloadKernelModules", sdk.UnloadKernelModules())
					},
				},
				{
					Title: "SearchSymbols",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SearchSymbols", sdk.SearchSymbols())
					},
				},
				{
					Title: "CpuidExecutionMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("CpuidExecutionMonitor", sdk.CpuidExecutionMonitor())
					},
				},
				{
					Title: "ControlRegisterModificationMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ControlRegisterModificationMonitor", sdk.ControlRegisterModificationMonitor())
					},
				},
				{
					Title: "DebugRegistersMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("DebugRegistersMonitor", sdk.DebugRegistersMonitor())
					},
				},
				{
					Title: "EptHook",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("EptHook", sdk.EptHook())
					},
				},
				{
					Title: "EptHook2",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("EptHook2", sdk.EptHook2())
					},
				},
				{
					Title: "IdtEntriesMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("IdtEntriesMonitor", sdk.IdtEntriesMonitor())
					},
				},
				{
					Title: "HideHyperDbg",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("HideHyperDbg", sdk.HideHyperDbg())
					},
				},
				{
					Title: "InterruptExternalMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("InterruptExternalMonitor", sdk.InterruptExternalMonitor())
					},
				},
				{
					Title: "IoInDetect",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("IoInDetect", sdk.IoInDetect())
					},
				},
				{
					Title: "IoOutDetect",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("IoOutDetect", sdk.IoOutDetect())
					},
				},
				{
					Title: "MeasureArgumentsForHide",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("MeasureArgumentsForHide", sdk.MeasureArgumentsForHide())
					},
				},
				{
					Title: "ModeInstructionsTrap",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ModeInstructionsTrap", sdk.ModeInstructionsTrap())
					},
				},
				{
					Title: "MonitorMemoryAccess",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("MonitorMemoryAccess", sdk.MonitorMemoryAccess())
					},
				},
				{
					Title: "MsrRead",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("MsrRead", sdk.MsrRead())
					},
				},
				{
					Title: "MsrWrite",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("MsrWrite", sdk.MsrWrite())
					},
				},
				{
					Title: "Pa2Va",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Pa2Va", sdk.Pa2Va())
					},
				},
				{
					Title: "PmcExecutionMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PmcExecutionMonitor", sdk.PmcExecutionMonitor())
					},
				},
				{
					Title: "Pte",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Pte", sdk.Pte())
					},
				},
				{
					Title: "ReversingMachineModuleUse",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ReversingMachineModuleUse", sdk.ReversingMachineModuleUse())
					},
				},
				{
					Title: "Syscall",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Syscall", sdk.Syscall())
					},
				},
				{
					Title: "SysRet",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SysRet", sdk.SysRet())
					},
				},
				{
					Title: "TraceExecution",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("TraceExecution", sdk.TraceExecution())
					},
				},
				{
					Title: "TrackModeTransitionInstructions",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("TrackModeTransitionInstructions", sdk.TrackModeTransitionInstructions())
					},
				},
				{
					Title: "TscInstructionsMonitor",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("TscInstructionsMonitor", sdk.TscInstructionsMonitor())
					},
				},
				{
					Title: "UnHide",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("UnHide", sdk.UnHide())
					},
				},
				{
					Title: "Va2Pa",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Va2Pa", sdk.Va2Pa())
					},
				},
				{
					Title: "VmCall",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("VmCall", sdk.VmCall())
					},
				},
				{
					Title: "HardwareClockDebugging",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("HardwareClockDebugging", sdk.HardwareClockDebugging())
					},
				},
				{
					Title: "AttachDebugThread",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("AttachDebugThread", sdk.AttachDebugThread())
					},
				},
				{
					Title: "ClearScreen",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ClearScreen", sdk.ClearScreen())
					},
				},
				{
					Title: "ConnectToMachine",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ConnectToMachine", sdk.ConnectToMachine())
					},
				},
				{
					Title: "DebugMachine",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("DebugMachine", sdk.DebugMachine())
					},
				},
				{
					Title: "DetachDebugging",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("DetachDebugging", sdk.DetachDebugging())
					},
				},
				{
					Title: "DisconnectSession",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("DisconnectSession", sdk.DisconnectSession())
					},
				},
				{
					Title: "DumpMemoryContext",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("DumpMemoryContext", sdk.DumpMemoryContext())
					},
				},
				{
					Title: "FormatsDiff",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("FormatsDiff", sdk.FormatsDiff())
					},
				},
				{
					Title: "HelpForCommand",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("HelpForCommand", sdk.HelpForCommand())
					},
				},
				{
					Title: "KillProcess",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("KillProcess", sdk.KillProcess())
					},
				},
				{
					Title: "ListenForClientConnection",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ListenForClientConnection", sdk.ListenForClientConnection())
					},
				},
				{
					Title: "LogClose",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("LogClose", sdk.LogClose())
					},
				},
				{
					Title: "LogOpen",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("LogOpen", sdk.LogOpen())
					},
				},
				{
					Title: "PageAvailableInRam",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("PageAvailableInRam", sdk.PageAvailableInRam())
					},
				},
				{
					Title: "ParseExecutableFiles",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ParseExecutableFiles", sdk.ParseExecutableFiles())
					},
				},
				{
					Title: "ProcessesView",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("ProcessesView", sdk.ProcessesView())
					},
				},
				{
					Title: "RestartProcess",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("RestartProcess", sdk.RestartProcess())
					},
				},
				{
					Title: "Script",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Script", sdk.Script())
					},
				},
				{
					Title: "StartProcess",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("StartProcess", sdk.StartProcess())
					},
				},
				{
					Title: "Status",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Status", sdk.Status())
					},
				},
				{
					Title: "SwitchThread",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SwitchThread", sdk.SwitchThread())
					},
				},
				{
					Title: "Symbol",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Symbol", sdk.SymbolActions())
					},
				},
				{
					Title: "SymbolPath",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("SymbolPath", sdk.SymbolPath())
					},
				},
				{
					Title: "Thread",
					Can:   func(any) bool { return true },
					Do: func(a any) {
						mylog.Warning("Thread", sdk.Thread())
					},
				},
			}
		},
		MarshalRow: func(node *widget.Node[xed.Disassembly]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum("todo")
			}
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("%X", node.Data.Opcode)},
				{Text: node.Data.Instruction},
				{Text: node.Data.Comment},
			}
		},
		UnmarshalRow: nil,
		SelectionChangedCallback: func(root *widget.Node[xed.Disassembly]) {
			for _, n := range root.SelectedRows(false) {
				mylog.Struct(n.Data)
			}
		},
		SetRootRowsCallBack: SetRootRowsCallBack,
		JsonName:            "cpu_dism_table",
		IsDocument:          false,
	})
	table.FileDropCallback = func(files []string) {
		f := files[0]
		ext := filepath.Ext(f)
		switch ext {
		case ".exe", ".dll", ".sys":
			fileName = f
			SetRootRowsCallBack(table)
		}
	}
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
			if node.Container() {
				node.Sum("todo")
			}
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
	X87TagWord     int
	X87ControlWord int
	X87StatusWord  int
	X87TW_0        int
	X87TW_1        int
	X87TW_2        int
	X87TW_3        int
	X87TW_4        int
	X87TW_5        int
	X87TW_6        int
	X87TW_7        int
	X87SW_B        int
	X87SW_C3       int
	X87SW_TOP      int
	X87SW_C2       int
	X87SW_C1       int
	X87SW_O        int
	X87SW_ES       int
	X87SW_SF       int
	X87SW_P        int
	X87SW_U        int
	X87SW_Z        int
	X87SW_D        int
	X87SW_I        int
	X87SW_C0       int
	X87CW_IC       int
	X87CW_RC       int
	X87CW_PC       int
	X87CW_PM       int
	X87CW_UM       int
	X87CW_OM       int
	X87CW_ZM       int
	X87CW_DM       int
	X87CW_IM       int
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
