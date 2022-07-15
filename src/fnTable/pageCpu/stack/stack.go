package stack

import (
	"encoding/hex"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"zydis"
)

type (
	Interface interface {
		canvasobjectapi.Interface
	}
	ds struct {
		hite    string
		address string
		data    string
		dism    string
		notes   string
	}
	object struct {
		ds
		lines []ds
	}
)

func (o *object) Lines() []ds {
	return o.lines
}

func (o *object) SetLines(lines []ds) {
	o.lines = append(o.lines, dd()...)
}

func New() Interface {
	return &object{
		ds:    ds{},
		lines: make([]ds, 0),
	}
}

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	o.SetLines(nil)
	t := widget.NewTable(
		func() (int, int) { return len(o.lines), 5 },
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(5, widget.NewLabel("Cell 000, 000"))
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			c := cell.(*fyne.Container)
			switch id.Col {
			case 0:
				c.Objects[0].(*widget.Label).SetText(o.lines[id.Row].hite)
			case 1:
				c.Objects[0].(*widget.Label).SetText(o.lines[id.Row].address)
			case 2:
				c.Objects[0].(*widget.Label).SetText(o.lines[id.Row].data)
			case 3:
				c.Objects[0].(*widget.Label).SetText(o.lines[id.Row].dism)
			case 4:
				c.Objects[0].(*widget.Label).SetText(o.lines[id.Row].notes)
			}
		})
	t.SetColumnWidth(0, 40)
	t.SetColumnWidth(1, 130)
	t.SetColumnWidth(2, 160)
	t.SetColumnWidth(3, 240)
	t.SetColumnWidth(4, 160)
	t.OnSelected = func(id widget.TableCellID) {
		println(id.Col)
		println(id.Row)
	}
	return t
}

func dd() (lines []ds) {
	lines = make([]ds, 0)
	// Show Zydis library version.
	major, minor, patch, build := zydis.Version()
	fmt.Printf("Version: %d.%d.%d.%d\n", major, minor, patch, build)

	// Instruction stream to decode.
	data := []byte{
		0x55,             // pushq %rbp
		0x48, 0x89, 0xe5, // movq %rsp, %rbp
		0x48, 0x83, 0xe4, 0xe0, // andq $-32, %rsp
		0x48, 0x81, 0xec, 0xa0, 0x00, 0x00, 0x00, // subq $160, %rsp
		0xc5, 0xf9, 0xef, 0xc0, // vpxor   %xmm0, %xmm0, %xmm0
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x60, // vmovdqa %ymm0, 96(%rsp)
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x40, // vmovdqa %ymm0, 64(%rsp)
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x20, // vmovdqa %ymm0, 32(%rsp)
		0xc5, 0xfd, 0x7f, 0x04, 0x24, // vmovdqa %ymm0, (%rsp)
		0x4c, 0x8d, 0x05, 0x00, 0x00, 0x00, 0x00, // leaq (%rip), %r8
	}

	// Initialize decoder context.
	decoder := zydis.NewDecoder(zydis.MachineMode64, zydis.AddressWidth64)

	// Initialize formatter.
	// Only required when instruction formatting ("disassembling"), like below.
	formatter, err := zydis.NewFormatter(zydis.FormatterStyleIntel)
	if err != nil {
		panic(err)
	}

	// Loop over the instructions in our buffer.
	// The runtime-address (instruction pointer) is chosen arbitrary here in
	// order to better visualize relative addressing.
	runtimeAddress := uint64(0x007FFFFFFF400000)

	for len(data) > 0 {
		instr, err := decoder.Decode(data)
		if err != nil {
			break
		}

		// Format and print the binary instruction structure.
		str, err := formatter.FormatInstruction(instr, runtimeAddress)
		if err != nil {
			panic(err)
		}
		//s := fmt.Sprintf(
		//	"%016x %s %s\n",
		//	runtimeAddress,
		//	hex.EncodeToString(data[:instr.Length]),
		//	str,
		//)
		d := ds{
			hite:    "",
			address: fmt.Sprintf("%016x", runtimeAddress),
			data:    hex.EncodeToString(data[:instr.Length]),
			dism:    str,
			notes:   "",
		}
		lines = append(lines, d)

		// Advance instruction pointer.
		runtimeAddress += uint64(instr.Length)

		// Drop decoded instructions from the stream.
		data = data[instr.Length:]
	}
	return
}
