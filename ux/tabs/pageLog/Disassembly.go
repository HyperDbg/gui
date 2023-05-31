package pageLog

import (
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/fynelib/myTable"
	"github.com/ddkwork/golibrary/src/go-zydis"
)

type (
	Disassembly interface {
		myTable.Interface
		Decode(buf []byte)
		Lines() []line
	}
	line struct {
		address string
		data    string
		dism    string
		notes   string
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

func (d *disassemblyObject) Decode(buf []byte) {
	major, minor, patch, build := zydis.Version()
	fmt.Printf("Version: %d.%d.%d.%d\n", major, minor, patch, build)
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

	for len(buf) > 0 {
		instr, err := decoder.Decode(buf)
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
		//	hex.EncodeToString(buf[:instr.Length]),
		//	str,
		//)
		l := line{
			address: fmt.Sprintf("%016x", runtimeAddress),
			data:    hex.EncodeToString(buf[:instr.Length]),
			dism:    str,
			notes:   "",
		}
		d.lines = append(d.lines, l)

		// Advance instruction pointer.
		runtimeAddress += uint64(instr.Length)

		// Drop decoded instructions from the stream.
		buf = buf[instr.Length:]
	}
}

func (d *disassemblyObject) Lines() []line     { return d.lines }
func newDisassemblyObject() *disassemblyObject { return &disassemblyObject{lines: make([]line, 0)} }
func (d *disassemblyObject) Append(data any)   { d.lines = append(d.lines, data.(line)) }
func (d *disassemblyObject) Header() []string {
	return []string{
		"address",
		"data",
		"dism",
		"notes",
	}
}
func (d *disassemblyObject) Rows(id int) []string {
	return []string{
		d.lines[id].address,
		d.lines[id].data,
		d.lines[id].dism,
		d.lines[id].notes,
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
