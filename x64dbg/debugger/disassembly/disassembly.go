package disassembly

import (
	"iter"
	"os"
	"slices"

	"gioui.org/layout"
	"github.com/ddkwork/ddk/xed"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
	"github.com/saferwall/pe"
	"golang.org/x/arch/x86/x86asm"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type State struct {
	table               *treetable.TreeTable[xed.Disassembly]
	currentSelectedData *xed.Disassembly
	currentSelectedAddr uint64
}

func NewState() *State {
	return &State{}
}

func (d *State) SetSelectedData(data *xed.Disassembly) {
	d.currentSelectedData = data
	if data != nil {
		d.currentSelectedAddr = data.Address
	}
}

func (d *State) GetSelectedAddress() uint64 {
	return d.currentSelectedAddr
}

func (d *State) GetSelectedData() *xed.Disassembly {
	return d.currentSelectedData
}

func (d *State) SetTable(t *treetable.TreeTable[xed.Disassembly]) {
	d.table = t
}

func (d *State) GetTable() *treetable.TreeTable[xed.Disassembly] {
	return d.table
}

func (d *State) Clear() {
	if d.table != nil {
		d.table.AirTable.SelectedNode = nil
		d.table.Root().SetChildren(nil)
	}
	d.currentSelectedData = nil
	d.currentSelectedAddr = 0
}

type Disassembler struct {
	state          *State
	table          *treetable.TreeTable[xed.Disassembly]
	baseAddress    uint64
	entryPoint     uint64
	exePath        string
	buffer         []byte
	rip            uint64
	moduleBase     uint64
	moduleSize     uint64
	OnRowSelected  func(data *xed.Disassembly)
	OnFollowInDump func(address uint64)
	ReadMemory     func(addr uint64, size int) []byte
}

func (d *Disassembler) Clear() {
	if d.state != nil {
		d.state.Clear()
	}
	if d.table != nil {
		d.table.AirTable.SelectedNode = nil
		d.table.Root().SetChildren(nil)
	}
	d.baseAddress = 0
	d.entryPoint = 0
	d.exePath = ""
	d.buffer = nil
	d.rip = 0
	d.moduleBase = 0
	d.moduleSize = 0
}

func (d *Disassembler) Self() any {
	return d
}

func (d *Disassembler) GetTable() *treetable.TreeTable[xed.Disassembly] {
	return d.table
}

func (d *Disassembler) SetUpdateParams(baseAddress, entryPoint uint64, exePath string) {
	d.baseAddress = baseAddress
	d.entryPoint = entryPoint
	d.exePath = exePath
}

func (d *Disassembler) SetRipBuffer(rip uint64, buffer []byte) {
	d.rip = rip
	d.buffer = buffer
}

func (d *Disassembler) SetModuleInfo(base, size uint64) {
	d.moduleBase = base
	d.moduleSize = size
}

func (d *Disassembler) IsAddressInRange(addr uint64) bool {
	if d.moduleBase == 0 || d.moduleSize == 0 {
		return false
	}
	return addr >= d.moduleBase && addr < d.moduleBase+d.moduleSize
}

func (d *Disassembler) SelectAddress(addr uint64) bool {
	d.rip = addr
	d.selectRipRow(d.table, addr)
	return true
}

func New() api.Interface {
	d := &Disassembler{
		state: NewState(),
	}
	d.initTable()
	return d
}

func (d *Disassembler) initTable() {
	d.table = treetable.NewTreeTable[xed.Disassembly]()

	d.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {
				if d.table.AirTable.SelectedNode == nil {
					return
				}
				data := d.table.GetData(d.table.AirTable.SelectedNode)

				followInDump := &menu.MenuItem{
					Label: "Follow in Dump",
					Do: func() {
						if d.OnFollowInDump != nil {
							d.OnFollowInDump(data.Address)
						}
					},
				}
				if !yield(followInDump) {
					return
				}

				if len(data.Opcode) > 0 {
					mode := 64
					inst, err := x86asm.Decode(data.Opcode, mode)
					if err == nil {
						for _, arg := range inst.Args {
							switch a := arg.(type) {
							case x86asm.Imm:
								addr := uint64(a)
								followImm := &menu.MenuItem{
									Label: "Follow Immediate in Dump",
									Do: func() {
										if d.OnFollowInDump != nil {
											d.OnFollowInDump(addr)
										}
									},
								}
								if !yield(followImm) {
									return
								}
							case x86asm.Mem:
								if a.Disp != 0 {
									addr := uint64(a.Disp)
									followMem := &menu.MenuItem{
										Label: "Follow Memory in Dump",
										Do: func() {
											if d.OnFollowInDump != nil {
												d.OnFollowInDump(addr)
											}
										},
									}
									if !yield(followMem) {
										return
									}
								}
							}
						}
					}
				}
			}
		},
		RowSelectedCallback: func() {
			if d.table.AirTable.SelectedNode != nil {
				data := d.table.GetData(d.table.AirTable.SelectedNode)
				d.state.SetSelectedData(&data)
				if d.OnRowSelected != nil {
					d.OnRowSelected(&data)
				}
			}
		},
		RowDoubleClickCallback: func() {
			addr := d.state.GetSelectedAddress()
			if d.table.AirTable.SelectedNode != nil && addr != 0 {
			}
		},
		SetRootRowsCallBack: func() { d.Update() },
		JsonName:            "cpu_dism_table",
		HideToolbar:         true,
		DisableBoxSelect:    true,
		DisableAddColumn:    true,
		DisableAddRow:       true,
		DisableSort:         true,
	}
}

func (d *Disassembler) Layout() layout.Widget {
	return d.table.Layout
}

func (d *Disassembler) LayoutDisassembly() layout.Widget {
	return d.table.AirTable.Layout
}

func (d *Disassembler) Update() error {
	if d.buffer != nil && d.rip != 0 {
		if d.IsAddressInRange(d.rip) {
			for _, node := range d.table.Root().Children {
				data := d.table.GetData(node)
				if data.Address == d.rip {
					d.selectRipRow(d.table, d.rip)
					return nil
				}
			}
		}

		d.table.Root().ResetChildren()
		size := len(d.buffer)
		if size > 2048 {
			size = 2048
		}
		x := xed.New(d.buffer[:size])
		x.SetBaseAddress(d.rip)
		x.SetIsFilterModel(false)
		x.Decode64()
		for _, object := range x.AsmObjects {
			d.table.AddChild(d.table.Root(), object)
		}
		d.table.AirTable.Refresh()
		d.selectRipRow(d.table, d.rip)
		return nil
	}

	if d.baseAddress != 0 && d.entryPoint != 0 && d.exePath == "" {
		if d.buffer != nil {
			if d.IsAddressInRange(d.entryPoint) {
				for _, node := range d.table.Root().Children {
					data := d.table.GetData(node)
					if data.Address == d.entryPoint {
						d.selectRipRow(d.table, d.entryPoint)
						return nil
					}
				}
			}

			d.table.Root().ResetChildren()
			size := len(d.buffer)
			if size > 2048 {
				size = 2048
			}
			x := xed.New(d.buffer[:size])
			x.SetBaseAddress(d.baseAddress)
			x.SetIsFilterModel(false)
			x.Decode64()
			for _, object := range x.AsmObjects {
				d.table.AddChild(d.table.Root(), object)
			}
			d.table.AirTable.Refresh()
			d.selectRipRow(d.table, d.entryPoint)
			return nil
		}
	}

	if d.exePath != "" {
		d.table.Root().ResetChildren()
		d.populateFromFile(d.table, d.exePath)
		d.table.AirTable.Refresh()
	}
	return nil
}

func (d *Disassembler) populateFromFile(t *treetable.TreeTable[xed.Disassembly], exePath string) {
	f := xed.ParserPe(exePath)
	optionalHeader := f.NtHeader.OptionalHeader
	switch o := optionalHeader.(type) {
	case pe.ImageOptionalHeader32:
		d.populateFromPE32(t, f, o, exePath)
	case pe.ImageOptionalHeader64:
		d.populateFromPE64(t, f, o, exePath)
	}
}

func (d *Disassembler) populateFromPE32(t *treetable.TreeTable[xed.Disassembly], f *pe.File, o pe.ImageOptionalHeader32, exePath string) {
	oepRVA := o.AddressOfEntryPoint
	imageBase := o.ImageBase
	oepVA := imageBase + oepRVA
	var oepFileOffset uint64

	for _, section := range f.Sections {
		if slices.Contains(section.PrettySectionFlags(), "Executable") {
			oepFileOffset = uint64(section.Header.PointerToRawData) + (uint64(oepRVA) - uint64(section.Header.VirtualAddress))
		}
	}

	if oepFileOffset == 0 {
		mylog.Check("未找到包含OEP节区或计算偏移不正确")
		return
	}
	buffer := make([]byte, 200)
	file := mylog.Check2(os.Open(exePath))
	defer file.Close()
	_, _ = file.ReadAt(buffer, int64(oepFileOffset))
	x := xed.New(buffer[:100])
	x.SetBaseAddress(uint64(oepVA))
	x.Decode32()
	for _, object := range x.AsmObjects {
		t.AddChild(t.Root(), object)
	}
	d.selectRipRow(t, uint64(oepVA))
}

func (d *Disassembler) populateFromPE64(t *treetable.TreeTable[xed.Disassembly], f *pe.File, o pe.ImageOptionalHeader64, exePath string) {
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
	file := mylog.Check2(os.Open(exePath))
	defer file.Close()
	_, _ = file.ReadAt(buffer, int64(oepFileOffset))

	x := xed.New(buffer[:100])
	x.SetBaseAddress(oepVA)
	x.Decode64()
	for _, object := range x.AsmObjects {
		t.AddChild(t.Root(), object)
	}
	d.selectRipRow(t, oepVA)
}

func (d *Disassembler) selectRipRow(t *treetable.TreeTable[xed.Disassembly], rip uint64) {
	for index, node := range t.Root().Children {
		data := t.GetData(node)
		if data.Address == rip {
			t.AirTable.SelectedNode = node
			d.state.SetSelectedData(&data)

			first, last := t.AirTable.RowList.VisibleRange()
			if index < first || index > last {
				t.AirTable.RowList.ScrollToIndex(index)
			}
			return
		}
	}
}
