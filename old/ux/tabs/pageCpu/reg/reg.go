package reg

import (
	"github.com/ddkwork/GolandProjects/ui/widget"
)

type (
	Row struct {
		name    *unison.Label
		address *unison.Button
		data    *unison.Label
	}
	object struct{}
)

func makeReg(name string) *unison.Panel {
	row := Row{
		name:    unison.NewLabel(),
		address: unison.NewButton(),
		data:    unison.NewLabel(),
	}
	row.address.ClickCallback = func() {
		//edit := container.NewVBox(
		//	swid.NewTextFormField("value", ""),
		//	swid.NewTextFormField("byte", ""),
		//	swid.NewTextFormField("signed", ""),
		//	swid.NewTextFormField("unsigned", ""),
		//	swid.NewTextFormField("ascii", ""),
		//)
	}
	row.name.Text = name
	//row.address.Text = name //todo input data
	//row.data.Text = name
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 3})
	panel.AddChild(row.name)
	panel.AddChild(row.address)
	panel.AddChild(row.data)
	return panel
}

func (o *object) CanvasObject() *unison.Panel {
	RAX := makeReg("RAX")
	RBX := makeReg("RBX")
	RCX := makeReg("RCX")
	RDX := makeReg("RDX")
	RBP := makeReg("RBP")
	RSP := makeReg("RSP")
	RSI := makeReg("RSI")
	RDI := makeReg("RDI")
	R8 := makeReg("R8")
	R9 := makeReg("R9")
	R10 := makeReg("R10")
	R11 := makeReg("R11")
	R12 := makeReg("R12")
	R13 := makeReg("R13")
	R14 := makeReg("R14")
	R15 := makeReg("R15")
	RIP := makeReg("RIP")
	RFLAGS := makeReg("RFLAGS")
	ZF := makeReg("ZF")
	OF := makeReg("OF")
	CF := makeReg("CF")
	PF := makeReg("PF")
	SF := makeReg("SF")
	TF := makeReg("TF")
	AF := makeReg("AF")
	DF := makeReg("DF")
	IF := makeReg("IF")
	LastError := makeReg("LastError")
	LastStatus := makeReg("LastStatus")
	GS := makeReg("GS")
	ES := makeReg("ES")
	CS := makeReg("CS")
	FS := makeReg("FS")
	DS := makeReg("DS")
	SS := makeReg("SS")
	ST0 := makeReg("ST(0)")
	ST1 := makeReg("ST(1)")
	ST2 := makeReg("ST(2)")
	ST3 := makeReg("ST(3)")
	ST4 := makeReg("ST(4)")
	ST5 := makeReg("ST(5)")
	ST6 := makeReg("ST(6)")
	ST7 := makeReg("ST(7)")
	x87TagWord := makeReg("x87TagWord")
	x87ControlWord := makeReg("x87ControlWord")
	x87StatusWord := makeReg("x87StatusWord")
	x87TW_0 := makeReg("x87TW_0")
	x87TW_1 := makeReg("x87TW_1")
	x87TW_2 := makeReg("x87TW_2")
	x87TW_3 := makeReg("x87TW_3")
	x87TW_4 := makeReg("x87TW_4")
	x87TW_5 := makeReg("x87TW_5")
	x87TW_6 := makeReg("x87TW_6")
	x87TW_7 := makeReg("x87TW_7")
	x87SW_B := makeReg("x87SW_B")
	x87SW_C3 := makeReg("x87SW_C3")
	x87SW_TOP := makeReg("x87SW_TOP")
	x87SW_C2 := makeReg("x87SW_C2")
	x87SW_C1 := makeReg("x87SW_C1")
	x87SW_O := makeReg("x87SW_O")
	x87SW_ES := makeReg("x87SW_ES")
	x87SW_SF := makeReg("x87SW_SF")
	x87SW_P := makeReg("x87SW_P")
	x87SW_U := makeReg("x87SW_U")
	x87SW_Z := makeReg("x87SW_Z")
	x87SW_D := makeReg("x87SW_D")
	x87SW_I := makeReg("x87SW_I")
	x87SW_C0 := makeReg("x87SW_C0")
	x87CW_IC := makeReg("x87CW_IC")
	x87CW_RC := makeReg("x87CW_RC")
	x87CW_PC := makeReg("x87CW_PC")
	x87CW_PM := makeReg("x87CW_PM")
	x87CW_UM := makeReg("x87CW_UM")
	x87CW_OM := makeReg("x87CW_OM")
	x87CW_ZM := makeReg("x87CW_ZM")
	x87CW_DM := makeReg("x87CW_DM")
	x87CW_IM := makeReg("x87CW_IM")
	MxCsr := makeReg("MxCsr : ")
	MxCsr_FZ := makeReg("MxCsr_FZ")
	MxCsr_PM := makeReg("MxCsr_PM")
	MxCsr_UM := makeReg("MxCsr_UM")
	MxCsr_OM := makeReg("MxCsr_OM")
	MxCsr_ZM := makeReg("MxCsr_ZM")
	MxCsr_IM := makeReg("MxCsr_IM")
	MxCsr_DM := makeReg("MxCsr_DM")
	MxCsr_DAZ := makeReg("MxCsr_DA")
	MxCsr_PE := makeReg("MxCsr_PE")
	MxCsr_UE := makeReg("MxCsr_UE")
	MxCsr_OE := makeReg("MxCsr_OE")
	MxCsr_ZE := makeReg("MxCsr_ZE")
	MxCsr_DE := makeReg("MxCsr_DE")
	MxCsr_IE := makeReg("MxCsr_IE")
	MxCsr_RC := makeReg("MxCsr_RC")
	XMM0 := makeReg("XMM0")
	XMM1 := makeReg("XMM1")
	XMM2 := makeReg("XMM2")
	XMM3 := makeReg("XMM3")
	XMM4 := makeReg("XMM4")
	XMM5 := makeReg("XMM5")
	XMM6 := makeReg("XMM6")
	XMM7 := makeReg("XMM7")
	XMM8 := makeReg("XMM8")
	XMM9 := makeReg("XMM9")
	XMM10 := makeReg("XMM10")
	XMM11 := makeReg("XMM11")
	XMM12 := makeReg("XMM12")
	XMM13 := makeReg("XMM13")
	XMM14 := makeReg("XMM14")
	XMM15 := makeReg("XMM15")
	YMM0 := makeReg("YMM0")
	YMM1 := makeReg("YMM1")
	YMM2 := makeReg("YMM2")
	YMM3 := makeReg("YMM3")
	YMM4 := makeReg("YMM4")
	YMM5 := makeReg("YMM5")
	YMM6 := makeReg("YMM6")
	YMM7 := makeReg("YMM7")
	YMM8 := makeReg("YMM8")
	YMM9 := makeReg("YMM9")
	YMM10 := makeReg("YMM10")
	YMM11 := makeReg("YMM11")
	YMM12 := makeReg("YMM12")
	YMM13 := makeReg("YMM13")
	YMM14 := makeReg("YMM14")
	YMM15 := makeReg("YMM15")
	DR0 := makeReg("DR0")
	DR1 := makeReg("DR1")
	DR2 := makeReg("DR2")
	DR3 := makeReg("DR3")
	DR6 := makeReg("DR6")
	DR7 := makeReg("DR7")

	regs := []*unison.Panel{
		RAX,
		RBX,
		RCX,
		RDX,
		RBP,
		RSP,
		RSI,
		RDI,
		R8,
		R9,
		R10,
		R11,
		R12,
		R13,
		R14,
		R15,
		RIP,
		RFLAGS,
		ZF,
		OF,
		CF,
		PF,
		SF,
		TF,
		AF,
		DF,
		IF,
		LastError,
		LastStatus,
		GS,
		ES,
		CS,
		FS,
		DS,
		SS,
		ST0,
		ST1,
		ST2,
		ST3,
		ST4,
		ST5,
		ST6,
		ST7,
		x87TagWord,
		x87ControlWord,
		x87StatusWord,
		x87TW_0,
		x87TW_1,
		x87TW_2,
		x87TW_3,
		x87TW_4,
		x87TW_5,
		x87TW_6,
		x87TW_7,
		x87SW_B,
		x87SW_C3,
		x87SW_TOP,
		x87SW_C2,
		x87SW_C1,
		x87SW_O,
		x87SW_ES,
		x87SW_SF,
		x87SW_P,
		x87SW_U,
		x87SW_Z,
		x87SW_D,
		x87SW_I,
		x87SW_C0,
		x87CW_IC,
		x87CW_RC,
		x87CW_PC,
		x87CW_PM,
		x87CW_UM,
		x87CW_OM,
		x87CW_ZM,
		x87CW_DM,
		x87CW_IM,
		MxCsr,
		MxCsr_FZ,
		MxCsr_PM,
		MxCsr_UM,
		MxCsr_OM,
		MxCsr_ZM,
		MxCsr_IM,
		MxCsr_DM,
		MxCsr_DAZ,
		MxCsr_PE,
		MxCsr_UE,
		MxCsr_OE,
		MxCsr_ZE,
		MxCsr_DE,
		MxCsr_IE,
		MxCsr_RC,
		XMM0,
		XMM1,
		XMM2,
		XMM3,
		XMM4,
		XMM5,
		XMM6,
		XMM7,
		XMM8,
		XMM9,
		XMM10,
		XMM11,
		XMM12,
		XMM13,
		XMM14,
		XMM15,
		YMM0,
		YMM1,
		YMM2,
		YMM3,
		YMM4,
		YMM5,
		YMM6,
		YMM7,
		YMM8,
		YMM9,
		YMM10,
		YMM11,
		YMM12,
		YMM13,
		YMM14,
		YMM15,
		DR0,
		DR1,
		DR2,
		DR3,
		DR6,
		DR7,
	}
	panel := unison.NewPanel()
	panel.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		pref.Width = 190
		pref.Height = 190
		return min, pref, pref
	})
	panel.SetLayout(&unison.FlexLayout{Columns: 9})
	for _, reg := range regs {
		panel.AddChild(reg)
	}

	scrollPanel := widget.NewScrollPanel(panel)

	//panel.AddChild(scrollPanel)
	return scrollPanel.AsPanel()
}

func New() *object { return &object{} }
