package symbolChoose

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/fpabl0/sparky-go/swid"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

/*
.sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbol
.sym download
.sym load
.sym reload
.connect local
.debug remote namedpipe \\.\pipe\HyperDbgDebug
.debug prepare serial 115200 com1
*/

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	path := swid.NewTextFormField("symbol path", "")
	load := widget.NewButtonWithIcon("load", nil, func() {})
	reLoad := widget.NewButtonWithIcon("reLoad", nil, func() {})
	local := widget.NewButtonWithIcon("local vmi", nil, func() {})
	clientPipeStart := widget.NewButtonWithIcon("clientPipeStart", nil, func() {})
	serverPipePort := swid.NewTextFormField("server Pipe Port", "")

	ntoskrnl := widget.NewCheck("ntoskrnl", func(b bool) {})
	win32k := widget.NewCheck("win32k", func(b bool) {})
	win32kfull := widget.NewCheck("win32kfull", func(b bool) {})
	fltmgr := widget.NewCheck("fltmgr", func(b bool) {})

	ntdll := widget.NewCheck("ntdll", func(b bool) {})
	kernel32 := widget.NewCheck("kernel32", func(b bool) {})
	kernelbase := widget.NewCheck("kernelbase", func(b bool) {})
	user32 := widget.NewCheck("user32", func(b bool) {})
	//todo add iphelp mac dll

	return container.NewGridWithColumns(4,
		ntoskrnl,
		win32k,
		win32kfull,
		fltmgr,
		ntdll,
		kernel32,
		kernelbase,
		user32,
		load,
		reLoad,
		local,
		clientPipeStart,
		path,
		serverPipePort,
	)
}

func New() Interface { return &object{} }
