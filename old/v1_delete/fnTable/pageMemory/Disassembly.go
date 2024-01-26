package pageMemory

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/fynelib/myTable"
	"github.com/ddkwork/golibrary/src/go-zydis"
)

type (
	Disassembly interface {
		myTable.Interface
		SetLines(lines []line)
		Lines() []line
	}
	line struct {
		address        string
		Size           string
		pageInfo       string
		content        string
		Type           string
		pageProtect    string
		defaultProtect string
	}
	disassemblyObject struct {
		lines []line
	}
)

func (d *disassemblyObject) ColumnLen() int {
	//TODO implement me
	panic("implement me")
}

func (d *disassemblyObject) ColumnWidths() []float32 {
	//TODO implement me
	panic("implement me")
}

func (d *disassemblyObject) SetLines(lines []line) {
	major, minor, patch, build := zydis.Version()
	fmt.Printf("Version: %d.%d.%d.%d\n", major, minor, patch, build)
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
	if !mylog.Error(err) {
		panic(err)
	}

	// Loop over the instructions in our buffer.
	// The runtime-address (instruction pointer) is chosen arbitrary here in
	// order to better visualize relative addressing.
	runtimeAddress := uint64(0x007FFFFFFF400000)

	for len(data) > 0 {
		instr, err := decoder.Decode(data)
		if !mylog.Error(err) {
			break
		}

		// Format and print the binary instruction structure.
		str, err := formatter.FormatInstruction(instr, runtimeAddress)
		if !mylog.Error(err) {
			panic(err)
		}
		//s := fmt.Sprintf(
		//	"%016x %s %s\n",
		//	runtimeAddress,
		//	hex.EncodeToString(data[:instr.Length]),
		//	str,
		//)
		l := line{
			address:        fmt.Sprintf("%016x", runtimeAddress),
			Size:           "",
			pageInfo:       "",
			content:        "",
			Type:           "",
			pageProtect:    "",
			defaultProtect: str,
		}
		d.lines = append(d.lines, l)

		// Advance instruction pointer.
		runtimeAddress += uint64(instr.Length)

		// Drop decoded instructions from the stream.
		data = data[instr.Length:]
	}
}

func (d *disassemblyObject) Lines() []line     { return d.lines }
func newDisassemblyObject() *disassemblyObject { return &disassemblyObject{lines: make([]line, 0)} }
func (d *disassemblyObject) Append(data any)   { d.lines = append(d.lines, data.(line)) }
func (d *disassemblyObject) Header() []string {
	return []string{
		"address",
		"size",
		"page info",
		"content",
		"type",
		"page protect",
		"default protect",
	}
}
func (d *disassemblyObject) Rows(id int) []string {
	return []string{
		d.lines[id].address,
		d.lines[id].Size,
		d.lines[id].pageInfo,
		d.lines[id].content,
		d.lines[id].Type,
		d.lines[id].pageProtect,
		d.lines[id].defaultProtect,
	}
}
func (d *disassemblyObject) Len() int { return len(d.lines) }
func (d *disassemblyObject) Sort(id int, ascend bool) {
	//TODO implement me
}
func (d *disassemblyObject) Filter(row string, id int) {
	//TODO implement me
}
func NewDisassembly() Disassembly { return &disassemblyObject{} }