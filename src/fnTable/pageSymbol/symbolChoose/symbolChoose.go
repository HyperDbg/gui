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
	path := swid.NewTextFormField("path", "")
	load := widget.NewButtonWithIcon("load", nil, func() {})
	reLoad := widget.NewButtonWithIcon("reLoad", nil, func() {})
	local := widget.NewButtonWithIcon("local", nil, func() {})
	clientPipeStart := widget.NewButtonWithIcon("clientPipeStart", nil, func() {})
	serverPipePort := swid.NewTextFormField("server Pipe Port", "")
	return container.NewGridWithRows(3,
		path,
		load,
		reLoad,
		local,
		clientPipeStart,
		serverPipePort,
	)
}

func New() Interface { return &object{} }
