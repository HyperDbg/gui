package scylla

import (
	"encoding/binary"
	"fmt"
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
)

type ImportModule struct {
	Name     string
	BaseAddr uint64
	Size     uint64
	Valid    bool
	Imports  []*ImportFunction
}

type ImportFunction struct {
	Name    string
	Ordinal uint16
	Hint    uint16
	Address uint64
	Valid   bool
	Suspect bool
	Module  *ImportModule
}

type ImportStatus int

const (
	ImportStatusValid ImportStatus = iota
	ImportStatusInvalid
	ImportStatusSuspect
)

type Manager struct {
	modules *safemap.M[string, *ImportModule]
	table   *treetable.TreeTable[ImportFunction]
}

func New() api.Interface {
	m := &Manager{
		modules: safemap.New[string, *ImportModule](),
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[ImportFunction]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			modules := m.GetAllModules()
			for _, module := range modules {
				for _, imp := range module.Imports {
					m.table.Root().AddChild(m.table.NewNode(*imp))
				}
			}
		},
		JsonName: "scylla",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) GetTable() *treetable.TreeTable[ImportFunction] {
	return m.table
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) ScanImports(baseAddress uint64, peData []byte, moduleName string, regionSize uint64) (*ImportModule, error) {
	module, err := m.scanModule(baseAddress, peData, moduleName, regionSize)
	if err != nil {
		return nil, err
	}

	m.modules.Update(module.Name, module)
	return module, nil
}

func (m *Manager) GetAllModules() []*ImportModule {
	result := make([]*ImportModule, 0)
	for _, module := range m.modules.Range() {
		result = append(result, module)
	}
	return result
}

func (m *Manager) FixImport(moduleName string, functionName string, newAddress uint64) error {
	module, exists := m.modules.Get(moduleName)
	if !exists {
		return fmt.Errorf("module not found: %s", moduleName)
	}

	for _, imp := range module.Imports {
		if imp.Name == functionName {
			imp.Address = newAddress
			imp.Valid = true
			imp.Suspect = false
			return nil
		}
	}

	return fmt.Errorf("function not found: %s", functionName)
}

func (m *Manager) ValidateImports() (valid, invalid, suspect int) {
	for _, module := range m.modules.Range() {
		for _, imp := range module.Imports {
			if imp.Valid {
				valid++
			} else if imp.Suspect {
				suspect++
			} else {
				invalid++
			}
		}
	}

	return
}

func (m *Manager) scanModule(baseAddress uint64, peData []byte, moduleName string, regionSize uint64) (*ImportModule, error) {
	moduleInfo := &ImportModule{
		Name:     moduleName,
		BaseAddr: baseAddress,
		Size:     regionSize,
		Valid:    true,
		Imports:  make([]*ImportFunction, 0),
	}

	imports, err := m.parseImportTable(peData, baseAddress)
	if err != nil {
		moduleInfo.Valid = false
		return moduleInfo, nil
	}

	moduleInfo.Imports = imports
	return moduleInfo, nil
}

func (m *Manager) parseImportTable(peData []byte, baseAddress uint64) ([]*ImportFunction, error) {
	imports := make([]*ImportFunction, 0)

	ntHeaders, err := parseNTHeaders(peData)
	if err != nil {
		return nil, err
	}

	if ntHeaders.OptionalHeader.Magic != 0x20b {
		return imports, nil
	}

	importDir := ntHeaders.OptionalHeader.DataDirectory[1]
	if importDir.VirtualAddress == 0 || importDir.Size == 0 {
		return imports, nil
	}

	importOffset := ntHeaders.OptionalHeader.ImageBase + uint64(importDir.VirtualAddress)
	importData := peData[importOffset : importOffset+uint64(importDir.Size)]

	for i := 0; i < len(importData); {
		impDesc, consumed := m.parseImportDescriptor(importData[i:])
		if consumed == 0 {
			break
		}

		if impDesc.Name != "" {
			moduleImports := m.parseImportFunctions(impDesc, importData[i:])
			imports = append(imports, moduleImports...)
		}

		i += consumed
	}

	return imports, nil
}

func (m *Manager) parseImportDescriptor(data []byte) (ImportDescriptor, int) {
	if len(data) < 20 {
		return ImportDescriptor{}, 0
	}

	desc := ImportDescriptor{
		OriginalFirstThunk: binary.LittleEndian.Uint32(data[0:4]),
		TimeDateStamp:      binary.LittleEndian.Uint32(data[4:8]),
		ForwarderChain:     binary.LittleEndian.Uint32(data[8:12]),
		Name:               m.readString(data[12:]),
		FirstThunk:         binary.LittleEndian.Uint32(data[16:20]),
	}

	if desc.Name == "" {
		return ImportDescriptor{}, 0
	}

	return desc, 20
}

func (m *Manager) parseImportFunctions(desc ImportDescriptor, data []byte) []*ImportFunction {
	functions := make([]*ImportFunction, 0)

	if desc.FirstThunk == 0 {
		return functions
	}

	thunkOffset := desc.OriginalFirstThunk
	if thunkOffset == 0 {
		thunkOffset = desc.FirstThunk
	}

	for i := uint32(0); ; i++ {
		thunkData := data[thunkOffset+i*8 : thunkOffset+i*8+8]
		if len(thunkData) < 8 {
			break
		}

		ordinal := binary.LittleEndian.Uint16(thunkData[0:2])
		address := binary.LittleEndian.Uint64(thunkData[0:8])

		if address == 0 {
			break
		}

		isOrdinal := (ordinal & 0x8000) != 0
		funcName := ""

		if !isOrdinal {
			nameOffset := address & 0x7FFFFFFFFFFFFFFF
			if nameOffset < uint64(len(data)) {
				funcName = m.readString(data[nameOffset:])
			}
		} else {
			funcName = fmt.Sprintf("Ordinal_%d", ordinal&0x7FFF)
		}

		valid := true
		suspect := false

		if isOrdinal {
			suspect = true
		}

		functions = append(functions, &ImportFunction{
			Name:    funcName,
			Ordinal: ordinal,
			Hint:    0,
			Address: address,
			Valid:   valid,
			Suspect: suspect,
		})
	}

	return functions
}

func (m *Manager) readString(data []byte) string {
	end := 0
	for i, b := range data {
		if b == 0 {
			end = i
			break
		}
	}
	if end == 0 {
		return ""
	}
	return string(data[:end])
}

type ImportDescriptor struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               string
	FirstThunk         uint32
}

type NTHeaders struct {
	Signature      uint32
	FileHeader     FileHeader
	OptionalHeader OptionalHeader64
}

type FileHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type OptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          uint16
	MinorLinkerVersion          uint16
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]DataDirectory
}

type DataDirectory struct {
	VirtualAddress uint32
	Size           uint32
}

func parseNTHeaders(data []byte) (*NTHeaders, error) {
	if len(data) < 64 {
		return nil, fmt.Errorf("data too short")
	}

	peOffset := binary.LittleEndian.Uint32(data[60:64])
	if peOffset >= uint32(len(data)) {
		return nil, fmt.Errorf("invalid PE offset")
	}

	signature := binary.LittleEndian.Uint32(data[peOffset : peOffset+4])
	if signature != 0x4550 {
		return nil, fmt.Errorf("invalid PE signature")
	}

	fileHeader := FileHeader{
		Machine:              binary.LittleEndian.Uint16(data[peOffset+4 : peOffset+6]),
		NumberOfSections:     binary.LittleEndian.Uint16(data[peOffset+6 : peOffset+8]),
		TimeDateStamp:        binary.LittleEndian.Uint32(data[peOffset+8 : peOffset+12]),
		PointerToSymbolTable: binary.LittleEndian.Uint32(data[peOffset+12 : peOffset+16]),
		NumberOfSymbols:      binary.LittleEndian.Uint32(data[peOffset+16 : peOffset+20]),
		SizeOfOptionalHeader: binary.LittleEndian.Uint16(data[peOffset+20 : peOffset+22]),
		Characteristics:      binary.LittleEndian.Uint16(data[peOffset+22 : peOffset+24]),
	}

	optHeaderOffset := peOffset + 24
	if len(data) < int(optHeaderOffset)+int(fileHeader.SizeOfOptionalHeader) {
		return nil, fmt.Errorf("optional header truncated")
	}

	optData := data[optHeaderOffset : optHeaderOffset+uint32(fileHeader.SizeOfOptionalHeader)]
	if len(optData) < 112 {
		return nil, fmt.Errorf("optional header too short")
	}

	optionalHeader := OptionalHeader64{
		Magic:               binary.LittleEndian.Uint16(optData[0:2]),
		ImageBase:           binary.LittleEndian.Uint64(optData[24:32]),
		NumberOfRvaAndSizes: binary.LittleEndian.Uint32(optData[108:112]),
	}

	if optionalHeader.Magic == 0x10b {
		for i := range 16 {
			optionalHeader.DataDirectory[i] = DataDirectory{
				VirtualAddress: binary.LittleEndian.Uint32(optData[112+i*8 : 112+i*8+4]),
				Size:           binary.LittleEndian.Uint32(optData[112+i*8+4 : 112+i*8+8]),
			}
		}
	}

	return &NTHeaders{
		Signature:      signature,
		FileHeader:     fileHeader,
		OptionalHeader: optionalHeader,
	}, nil
}
