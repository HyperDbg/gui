package ui

import (
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/mcp/bindings/go/sdk"
	"iter"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"gioui.org/app"
	"gioui.org/layout"
	"github.com/ddkwork/ddk/xed"

	"github.com/ddkwork/ux"
	"github.com/ddkwork/ux/languages"
	"github.com/ddkwork/ux/widget/material"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/saferwall/pe"
)

type Cpu struct {
	asmAndImmData ux.Split
	spTop         ux.Split
	spBottom      ux.Split
	sp            ux.Split
}

func NewCpu() *Cpu {
	asmAndImmData := ux.Split{
		Ratio:  0.85, // 布局比例，0 表示居中，-1 表示完全靠左，1 表示完全靠右
		Bar:    10,
		Axis:   layout.Vertical,
		First:  LayoutDisassemblyTable().Layout, // left 表格
		Second: NewImm().Layout,                 // right 寄存器 todo
	}

	spTop := ux.Split{
		Ratio:  0.3, // 布局比例，0 表示居中，-1 表示完全靠左，1 表示完全靠右
		Bar:    10,
		Axis:   layout.Horizontal,
		First:  asmAndImmData.Layout,    // left 表格
		Second: LayoutRegister().Layout, // right 寄存器 todo
	}
	spBottom := ux.Split{
		Ratio: 0.2, // 布局比例，0 表示居中，-1 表示完全靠左，1 表示完全靠右
		Bar:   10,
		Axis:  layout.Horizontal,
		// todo
		//	tabs.NewTab("memory2")
		//	tabs.NewTab("memory3")
		//	tabs.NewTab("memory4")
		//	tabs.NewTab("memory5")
		//	tabs.NewTab("watch1")
		//	tabs.NewTab("var")
		//	tabs.NewTab("struct")
		First:  ux.NewCodeEditor(hex.Dump(testHexDat), languages.GoKind).Layout, // left memory editor todo https://github.com/itchyny/bed
		Second: LayoutStackTable().Layout,                                       // right stack
	}

	sp := ux.Split{
		Ratio:  0.4, // 布局比例，0 表示居中，-1 表示完全靠左，1 表示完全靠右
		Bar:    10,
		Axis:   layout.Vertical,
		First:  spTop.Layout,    //
		Second: spBottom.Layout, //
	}
	return &Cpu{
		asmAndImmData: asmAndImmData,
		spTop:         spTop,
		spBottom:      spBottom,
		sp:            sp,
	}
}

func (c *Cpu) Layout(gtx layout.Context) layout.Dimensions {
	return c.sp.Layout(gtx)
}

func LayoutRegister() *ux.StructView[Register] {
	return ux.NewStructView("Register", testRegData, func(key string, field any) (value string) {
		return fmt.Sprintf("%016x", field)
	})
}

// todo D:\desk\go\golang.org\x\arch@v0.8.0\x86\x86asm\inst.go :L150
var testRegData = Register{
	RAX:            0,
	RBX:            0x00007FF88500B7F0, // "LdrpInitializeProcess"
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

type (
	Imm struct {
		*layout.List
		rows []ImmData
	}
	ImmData struct {
		Reg     string // todo use kind of  Register enum
		Address uint64
		Mem     string
	}
	FastCall struct {
		// Index    int
		Register string
		Address  int
		ImmData  string
	}
)

func NewImm() *Imm {
	return &Imm{List: &layout.List{Axis: layout.Vertical}}
}

func (m *Imm) Layout(gtx layout.Context) layout.Dimensions {
	if m.rows == nil {
		rows := make([]ImmData, 0)
		for range 6 {
			rows = append(rows, ImmData{
				Reg:     "RAX",
				Address: 0x1234567890,
				Mem:     hex.EncodeToString([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x0}),
			})
		}
		m.rows = rows
	}
	return m.List.Layout(gtx, len(m.rows), func(gtx layout.Context, i int) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   5,
					Alignment: 0,
					WeightSum: 0,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H6(ux.NewTheme(), "Reg:"+m.rows[i].Reg).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H6(ux.NewTheme(), "  Address:"+strconv.FormatUint(m.rows[i].Address, 16)).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H6(ux.NewTheme(), "  Mem:"+string(m.rows[i].Mem)).Layout(gtx)
					}),
				)
			}))
	})
}

var TargetExePath = ""

func LayoutDisassemblyTable() ux.Widget {
	t := ux.NewTreeTable(xed.Disassembly{})
	t.TableContext = ux.TableContext[xed.Disassembly]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[xed.Disassembly]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
				// yield(ux.ContextMenuItem{
				//	Title:         "todo",
				//	Icon:          nil,
				//	Can:           nil,
				//	Do:            nil,
				//	AppendDivider: false,
				//	Clickable:     widget.Clickable{},
				// })

				/*
						return []ux.ContextMenuItem{
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
								//BreakpointClearByID
								//	BreakpointDisableByID
								//	BreakpointEnableByID
								//	BreakpointList
								//	SetBreakpoint0xcc
								//	CoreOperatingProcessorForShowsAndChanges
								//	CpuFeaturesForCollectsReport
								//	FlushKernelModeBuffers
								//	ContinueDebugger
								//	StepOut
								//	StepIn
								//	CallstackOrThreadView //跟随当前地址的堆栈，同时刷新右下角的堆栈panel
								//	ListModules
								//	LoadDriversAndModules
								//	OutputEventForwardingInstance
								//	StepExecuteSingleInstruction
								//	PauseKernelEvents
								//	PreactivateSpecialFunctionality
								//	PreallocateBuffer
								//	PrintEvaluateExpressions
								//	RegistersReadOrModify
								//	ReadMsr
								//	SearchMemoryPattern
								//	SettingsManagement
								//	SleepDebugger
								//	StepInExecute
								//	TestHyperDbgFeatures
								//	UnloadKernelModules
								//	SearchSymbols
								//	CpuidExecutionMonitor
								//	ControlRegisterModificationMonitor
								//	DebugRegistersMonitor
								//	EptHook
								//	EptHook2
								//	IdtEntriesMonitor
								//	HideHyperDbg
								//	InterruptExternalMonitor
								//	IoInDetect
								//	IoOutDetect
								//	MeasureArgumentsForHide
								//	ModeInstructionsTrap
								//	MonitorMemoryAccess
								//	MsrRead
								//	MsrWrite
								//	Pa2Va
								//	PmcExecutionMonitor
								//	Pte
								//	ReversingMachineModuleUse
								//	Syscall
								//	SysRet
								//	TraceExecution
								//	TrackModeTransitionInstructions
								//	TscInstructionsMonitor
								//	UnHide
								//	Va2Pa
								//	VmCall
								//	HardwareClockDebugging
								//	AttachDebugThread
								//	ClearScreen
								//	ConnectToMachine
								//	DebugMachine
								//	DetachDebugging
								//	DisconnectSession
								//	DumpMemoryContext
								//	FormatsDiff
								//	HelpForCommand
								//	KillProcess
								//	ListenForClientConnection
								//	LogClose
								//	LogOpen
								//	PageAvailableInRam
								//	ParseExecutableFiles
								//	ProcessesView
								//	RestartProcess
								//	Script
								//	StartProcess
								//	Status
								//	SwitchThread
								//	Symbol
								//	SymbolPath
								//	Thread

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
							mylog.Warning("StartProcess", sdk.StartProcess_(""))
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
				*/
			}
		},
		MarshalRowCells: func(n *ux.Node[xed.Disassembly]) (cells []ux.CellData) {
			if n.Container() {
				n.SumChildren()
			}
			cells = ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				switch key {
				case "Address":
					return fmt.Sprintf("%016X", field)
				case "Opcode":
					return fmt.Sprintf("%X", field)
				}
				return ""
			})
			for i, cell := range cells {
				if cell.Key == "Instruction" {
					cell.IsNasm = true
					cells[i] = cell
				}
			}
			return cells
		},
		UnmarshalRowCells: func(n *ux.Node[xed.Disassembly], rows []ux.CellData) xed.Disassembly {
			return ux.UnmarshalRow[xed.Disassembly](rows, func(key, value string) (field any) {
				switch key {
				case "Address":
					return mylog.Check2(strconv.ParseUint(value, 16, 64))
				case "Opcode":
					return mylog.Check2(strconv.ParseUint(value, 16, 8))
				}
				return nil
			})
		},
		RowSelectedCallback: func() {
			mylog.Struct(t.SelectedNode.Data)
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			if TargetExePath == "" {
				return
			}

			f := xed.ParserPe(TargetExePath)
			// b := stream.NewBuffer(fileName)
			optionalHeader := f.NtHeader.OptionalHeader
			switch o := optionalHeader.(type) {
			case pe.ImageOptionalHeader32:
				oepRVA := o.AddressOfEntryPoint
				imageBase := o.ImageBase
				oepVA := imageBase + oepRVA
				var oepFileOffset uint64

				// for _, section := range f.Sections {
				//	if section.String() == ".text" {
				//		oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
				//		break
				//	}
				// }

				// for _, section := range f.Sections {
				//	if section.Header.Characteristics&section.Header.Characteristics == pe.ImageSectionMemExecute {
				//		oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
				//		break
				//	}
				// }
				for _, section := range f.Sections {
					// mylog.Info("PrettySectionFlags", section.PrettySectionFlags())
					if slices.Contains(section.PrettySectionFlags(), "Executable") {
						oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
					}
				}

				if oepFileOffset == 0 {
					mylog.Check("未找到包含OEP节区或计算偏移不正确")
					return
				}
				buffer := make([]byte, 200)
				file := mylog.Check2(os.Open(TargetExePath))
				defer file.Close()
				mylog.Check2(file.ReadAt(buffer, int64(oepFileOffset)))
				// fmt.Printf("OEP File Off %x\n", oepFileOffset)
				// fmt.Printf("OEP VA %x\n", oepVA)
				// fmt.Printf("Entry Point RVA %x\n", oepRVA)
				// fmt.Printf("OEP Data %x\n", buffer[:100])

				x := xed.New(buffer[:100])
				x.SetBaseAddress(uint64(oepVA))
				x.Decode32()
				for _, object := range x.AsmObjects {
					t.Root.AddChildByData(object)
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
					mylog.Check("未找到包含OEP节区或计算偏移不正确")
					return
				}
				buffer := make([]byte, 200)
				file := mylog.Check2(os.Open(TargetExePath))
				defer file.Close()
				mylog.Check2(file.ReadAt(buffer, int64(oepFileOffset)))
				// fmt.Printf("OEP File Off %x\n", oepFileOffset)
				// fmt.Printf("OEP VA %x\n", oepVA)
				// fmt.Printf("Entry Point RVA %x\n", oepRVA)
				// fmt.Printf("OEP Data %x\n", buffer[:100])

				x := xed.New(buffer[:100])
				x.SetBaseAddress(oepVA)
				x.Decode64()
				for _, object := range x.AsmObjects {
					t.Root.AddChildByData(object)
				}
				// oep_rva := uint64(o.AddressOfEntryPoint)
				// image_base := uint64(o.ImageBase)
				// oep_va := image_base + oep_rva
				// var oep_file_offset uint64
				//
				// // 查找 .text 节区
				// for _, section := range f.Sections {
				//	if section.String() == ".text" {
				//		oep_file_offset = uint64(section.Header.PointerToRawData) + (oep_rva - uint64(section.Header.VirtualAddress))
				//		mylog.Struct("todo",section)
				//		break
				//	}
				// }
				//
				// if oep_file_offset == 0 {
				//	mylog.Check("未找到 .text 节区信息或计算偏移不正确")
				// }
				//
				// // 确定文件中的偏移位置
				// buffer := stream.NewBuffer(fileName)
				// if oep_file_offset+200 > uint64(len(buffer.Bytes())) {
				//	mylog.Check("计算出的文件偏移超出文件大小")
				// }
				//
				// // 确认调试信息：
				// mylog.Hex("oep_file_offset", oep_file_offset)
				// mylog.Hex("oep_va", oep_va)
				// mylog.Hex("entry_point_rva", oep_rva)
				//
				// // 读取指定文件偏移二进制数据：
				// oep_data := buffer.Bytes()[oep_file_offset : oep_file_offset+200]
				//
				// fmt.Printf("oep_data: %x\n", oep_data[:20])
				// // 使用反汇编工具解析数据
				// x := xed.NewOrderedMap(oep_data)
				// x.SetBaseAddress(oep_va)
				// x.Decode64()
				// for _, object := range x.AsmObjects {
				//	root.AddChildByData(object)
				// }

			}
		},
		JsonName:   "cpu_dism_table",
		IsDocument: false,
	}

	app.FileDropCallback(func(files []string) {
		for _, file := range files {
			mylog.Info("FileDropCallback", file)
			switch filepath.Ext(file) {
			case ".exe", ".dll", ".sys":
				mylog.Info("dropped file: ", file)
				TargetExePath = file
				if !d.VmxSupportDetection() {
					panic("vmx not support")
				}
				mylog.Success("vmx support detected")
				t.SetRootRowsCallBack()

				d.ConnectLocalDebugger()
				d.LoadVmmModule()
				d.RunCommand("start")
			default:
				mylog.Check(file + " is not a valid file type")
			}
		}
	})
	return t
}

var d = sdk.Debugger{}

type Stack struct {
	Address    int
	Data       int
	Context    string // api Name from symbol or address
	IsApiOrArg bool   // decode args from plugin pkg
}

func LayoutStackTable() ux.Widget {
	t := ux.NewTreeTable(Stack{})
	t.TableContext = ux.TableContext[Stack]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Stack]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Stack]) (cells []ux.CellData) {
			addressFmt := fmt.Sprintf("%016X", n.Data.Address)
			if n.Container() {
				addressFmt = n.SumChildren()
			}
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				switch key {
				case "Address":
					return addressFmt
				case "Data":
					return fmt.Sprintf("%X", field)
				}
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Stack], rows []ux.CellData) Stack {
			return ux.UnmarshalRow[Stack](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for i := 0; i < len(stackMock); i++ {
				stack := stackMock[i]
				if stack.IsApiOrArg { // todo decoe args https://docs.hyperdbg.org/commands/debugging-commands/k
					container := ux.NewContainerNode(fmt.Sprintf("%016X", stack.Address), stack)
					t.Root.AddChild(container)
					stacks := stackMock[i+1:]
					for j, s := range stacks {
						if !s.IsApiOrArg {
							i += j
							break
						}
						container.AddChildByData(s)
					}
					continue
				}
				t.Root.AddChildByData(stack)
			}
		},
		JsonName:   "stack",
		IsDocument: false,
	}
	return t
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

//var TargetExePath = sdk.TargetFilePath

/*
0019ECEC   77715BFE  /CALL 到 DeviceIoControl 来自 77715BF8
0019ECF0   000001A8  |hDevice = 000001A8 (window)
0019ECF4   00390008  |IoControlCode = 390008
0019ECF8   00000000  |InBuffer = NULL
0019ECFC   00000000  |InBufferSize = 0
0019ED00   0019ED5C  |OutBuffer = 0019ED5C
0019ED04   00000030  |OutBufferSize = 30 (48.)
0019ED08   0019ED28  |pBytesReturned = 0019ED28
0019ED0C   00000000  \pOverlapped = NULL
0019ED10   00000000
0019ED14   00000030
0019ED18   0019ED2C
0019ED1C   0000000F
0019ED20   00000000
0019ED24   00000000
0019ED28   00000030
0019ED2C  /0019ED90
0019ED30  |77715B60  返回到 77715B60
0019ED34  |0019ED5C
0019ED38  |00000030
0019ED3C  |0019ED50
0019ED40  |00000000
0019ED44  |0019EE58
0019ED48  |00000000
0019ED4C  |0019ED68
0019ED50  |00000000
0019ED54  |00000000
0019ED58  |00000000
0019ED5C  |00000000
0019ED60  |00000000
0019ED64  |00000000
0019ED68  |00000000
0019ED6C  |00000000
0019ED70  |00000000
0019ED74  |00000000
0019ED78  |00000000
0019ED7C  |00000000
0019ED80  |00000000
0019ED84  |00000000
0019ED88  |00000000
0019ED8C  |7F1CF281
0019ED90  ]0019EE30
0019ED94  |77715D3A  返回到 77715D3A 来自 77715B12
0019ED98  |77727520
0019ED9C  |0019EE58
0019EDA0  |00000000
0019EDA4  |0074C301
0019EDA8  |00000000
0019EDAC  |0019EDF8
0019EDB0  |00000000
0019EDB4  |0019EDE4
0019EDB8  |00000000
0019EDBC  |00000000
0019EDC0  |00000000
0019EDC4  |00000000
0019EDC8  |00000001
0019EDCC  |00000000
0019EDD0  |777D7B89  返回到 ntdll.777D7B89 来自 ntdll.777E8420
0019EDD4  |7772E070
*/
var stackMock = []Stack{
	{Address: 0x0019ECEC, Data: 0x77715BFE, Context: "CALL 到 DeviceIoControl 来自 77715BF8", IsApiOrArg: true},
	{Address: 0x0019ECF0, Data: 0x000001A8, Context: "|hDevice = 000001A8 (window)", IsApiOrArg: true},
	{Address: 0x0019ECF4, Data: 0x00390008, Context: "|IoControlCode = 390008", IsApiOrArg: true},
	{Address: 0x0019ECF8, Data: 0x00390009, Context: "|InBuffer = NULL", IsApiOrArg: true},
	{Address: 0x0019ECFC, Data: 0x00000000, Context: "|InBufferSize = 0", IsApiOrArg: true},
	{Address: 0x0019ED00, Data: 0x0019ED5C, Context: "|OutBuffer = 0019ED5C", IsApiOrArg: true},
	{Address: 0x0019ED04, Data: 0x00000030, Context: "|OutBufferSize = 30 (48.)", IsApiOrArg: true},
	{Address: 0x0019ED08, Data: 0x0019ED28, Context: "|pBytesReturned = 0019ED28", IsApiOrArg: true},
	{Address: 0x0019ED0C, Data: 0x00000000, Context: "pOverlapped = NULL", IsApiOrArg: true},
	{Address: 0x0019ED10, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED14, Data: 0x00000030, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED18, Data: 0x0019ED2C, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED1C, Data: 0x0000000F, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED20, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED24, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED28, Data: 0x00000030, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED2C, Data: 0x0019ED90, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED30, Data: 0x77715B60, Context: "返回到 77715B60", IsApiOrArg: true},
	{Address: 0x0019ED34, Data: 0x0019ED5C, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED38, Data: 0x00000030, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED3C, Data: 0x0019ED50, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED40, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED44, Data: 0x0019EE58, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED48, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED4C, Data: 0x0019ED68, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED50, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED54, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED58, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED5C, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED60, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED64, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED68, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED6C, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED70, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED74, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED78, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED7C, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED80, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED84, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED88, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED8C, Data: 0x7F1CF281, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED90, Data: 0x0019EE30, Context: "]", IsApiOrArg: false},
	{Address: 0x0019ED94, Data: 0x77715D3A, Context: "返回到 77715D3A 来自 77715B12", IsApiOrArg: true},
	{Address: 0x0019ED98, Data: 0x77727520, Context: "", IsApiOrArg: false},
	{Address: 0x0019ED9C, Data: 0x0019EE58, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDA0, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDA4, Data: 0x0074C301, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDA8, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDAC, Data: 0x0019EDF8, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDB0, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDB4, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDB8, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDBC, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDC0, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDC4, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDC8, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDCC, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDD0, Data: 0x00000000, Context: "", IsApiOrArg: false},
	{Address: 0x0019EDD4, Data: 0x00000000, Context: "", IsApiOrArg: false},
}
