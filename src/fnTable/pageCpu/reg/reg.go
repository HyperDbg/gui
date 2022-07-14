package reg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/fpabl0/sparky-go/swid"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func makeReg(name string, window fyne.Window) *swid.TextFormField {
	show := container.NewVBox(
		swid.NewTextFormField("value", ""),
		swid.NewTextFormField("byte", ""),
		swid.NewTextFormField("signed", ""),
		swid.NewTextFormField("unsigned", ""),
		swid.NewTextFormField("ascii", ""),
	)
	show = show
	reg := swid.NewTextFormField(name, "")
	//reg.OnChanged = func(s string) { dialog.ShowCustom("edit", "ok", show, window) }
	return reg
}

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	RAX := makeReg("RAX", window)
	RBX := makeReg("RBX", window)
	RCX := makeReg("RCX", window)
	RDX := makeReg("RDX", window)
	RBP := makeReg("RBP", window)
	RSP := makeReg("RSP", window)
	RSI := makeReg("RSI", window)
	RDI := makeReg("RDI", window)
	R8 := makeReg("R8", window)
	R9 := makeReg("R9", window)
	R10 := makeReg("R10", window)
	R11 := makeReg("R11", window)
	R12 := makeReg("R12", window)
	R13 := makeReg("R13", window)
	R14 := makeReg("R14", window)
	R15 := makeReg("R15", window)
	RIP := makeReg("RIP", window)
	RFLAGS := makeReg("RFLAGS", window)
	ZF := makeReg("ZF", window)
	OF := makeReg("OF", window)
	CF := makeReg("CF", window)
	PF := makeReg("PF", window)
	SF := makeReg("SF", window)
	TF := makeReg("TF", window)
	AF := makeReg("AF", window)
	DF := makeReg("DF", window)
	IF := makeReg("IF", window)
	LastError := makeReg("LastError", window)
	LastStatus := makeReg("LastStatus", window)
	GS := makeReg("GS", window)
	ES := makeReg("ES", window)
	CS := makeReg("CS", window)
	FS := makeReg("FS", window)
	DS := makeReg("DS", window)
	SS := makeReg("SS", window)
	ST0 := makeReg("ST(0)", window)
	ST1 := makeReg("ST(1)", window)
	ST2 := makeReg("ST(2)", window)
	ST3 := makeReg("ST(3)", window)
	ST4 := makeReg("ST(4)", window)
	ST5 := makeReg("ST(5)", window)
	ST6 := makeReg("ST(6)", window)
	ST7 := makeReg("ST(7)", window)
	x87TagWord := makeReg("x87TagWord", window)
	x87ControlWord := makeReg("x87ControlWord", window)
	x87StatusWord := makeReg("x87StatusWord", window)
	x87TW_0 := makeReg("x87TW_0", window)
	x87TW_1 := makeReg("x87TW_1", window)
	x87TW_2 := makeReg("x87TW_2", window)
	x87TW_3 := makeReg("x87TW_3", window)
	x87TW_4 := makeReg("x87TW_4", window)
	x87TW_5 := makeReg("x87TW_5", window)
	x87TW_6 := makeReg("x87TW_6", window)
	x87TW_7 := makeReg("x87TW_7", window)
	x87SW_B := makeReg("x87SW_B", window)
	x87SW_C3 := makeReg("x87SW_C3", window)
	x87SW_TOP := makeReg("x87SW_TOP", window)
	x87SW_C2 := makeReg("x87SW_C2", window)
	x87SW_C1 := makeReg("x87SW_C1", window)
	x87SW_O := makeReg("x87SW_O", window)
	x87SW_ES := makeReg("x87SW_ES", window)
	x87SW_SF := makeReg("x87SW_SF", window)
	x87SW_P := makeReg("x87SW_P", window)
	x87SW_U := makeReg("x87SW_U", window)
	x87SW_Z := makeReg("x87SW_Z", window)
	x87SW_D := makeReg("x87SW_D", window)
	x87SW_I := makeReg("x87SW_I", window)
	x87SW_C0 := makeReg("x87SW_C0", window)
	x87CW_IC := makeReg("x87CW_IC", window)
	x87CW_RC := makeReg("x87CW_RC", window)
	x87CW_PC := makeReg("x87CW_PC", window)
	x87CW_PM := makeReg("x87CW_PM", window)
	x87CW_UM := makeReg("x87CW_UM", window)
	x87CW_OM := makeReg("x87CW_OM", window)
	x87CW_ZM := makeReg("x87CW_ZM", window)
	x87CW_DM := makeReg("x87CW_DM", window)
	x87CW_IM := makeReg("x87CW_IM", window)
	MxCsr := makeReg("MxCsr : ", window)
	MxCsr_FZ := makeReg("MxCsr_FZ", window)
	MxCsr_PM := makeReg("MxCsr_PM", window)
	MxCsr_UM := makeReg("MxCsr_UM", window)
	MxCsr_OM := makeReg("MxCsr_OM", window)
	MxCsr_ZM := makeReg("MxCsr_ZM", window)
	MxCsr_IM := makeReg("MxCsr_IM", window)
	MxCsr_DM := makeReg("MxCsr_DM", window)
	MxCsr_DAZ := makeReg("MxCsr_DA", window)
	MxCsr_PE := makeReg("MxCsr_PE", window)
	MxCsr_UE := makeReg("MxCsr_UE", window)
	MxCsr_OE := makeReg("MxCsr_OE", window)
	MxCsr_ZE := makeReg("MxCsr_ZE", window)
	MxCsr_DE := makeReg("MxCsr_DE", window)
	MxCsr_IE := makeReg("MxCsr_IE", window)
	MxCsr_RC := makeReg("MxCsr_RC", window)
	XMM0 := makeReg("XMM0", window)
	XMM1 := makeReg("XMM1", window)
	XMM2 := makeReg("XMM2", window)
	XMM3 := makeReg("XMM3", window)
	XMM4 := makeReg("XMM4", window)
	XMM5 := makeReg("XMM5", window)
	XMM6 := makeReg("XMM6", window)
	XMM7 := makeReg("XMM7", window)
	XMM8 := makeReg("XMM8", window)
	XMM9 := makeReg("XMM9", window)
	XMM10 := makeReg("XMM10", window)
	XMM11 := makeReg("XMM11", window)
	XMM12 := makeReg("XMM12", window)
	XMM13 := makeReg("XMM13", window)
	XMM14 := makeReg("XMM14", window)
	XMM15 := makeReg("XMM15", window)
	YMM0 := makeReg("YMM0", window)
	YMM1 := makeReg("YMM1", window)
	YMM2 := makeReg("YMM2", window)
	YMM3 := makeReg("YMM3", window)
	YMM4 := makeReg("YMM4", window)
	YMM5 := makeReg("YMM5", window)
	YMM6 := makeReg("YMM6", window)
	YMM7 := makeReg("YMM7", window)
	YMM8 := makeReg("YMM8", window)
	YMM9 := makeReg("YMM9", window)
	YMM10 := makeReg("YMM10", window)
	YMM11 := makeReg("YMM11", window)
	YMM12 := makeReg("YMM12", window)
	YMM13 := makeReg("YMM13", window)
	YMM14 := makeReg("YMM14", window)
	YMM15 := makeReg("YMM15", window)
	DR0 := makeReg("DR0", window)
	DR1 := makeReg("DR1", window)
	DR2 := makeReg("DR2", window)
	DR3 := makeReg("DR3", window)
	DR6 := makeReg("DR6", window)
	DR7 := makeReg("DR7", window)
	regBox := container.NewVBox(
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
	)
	return container.NewVScroll(regBox)
}

func New() Interface { return &object{} }
