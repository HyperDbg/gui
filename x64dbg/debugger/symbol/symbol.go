package symbol

import (
	"fmt"
	"iter"
	"sync"
	"syscall"
	"unsafe"

	"gioui.org/layout"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

const (
	MAX_PATH = 260
)

type SymbolType int

const (
	SymbolTypeImport SymbolType = iota
	SymbolTypeExport
	SymbolTypeSymbol
)

type SymbolInfo struct {
	Type       SymbolType
	Address    uint64
	Size       uint32
	Name       string
	ModuleName string
	FileName   string
	LineNumber uint32
	LineOffset uint32
}

type ModuleInfo struct {
	BaseAddress   uint64
	Size          uint32
	Name          string
	FileName      string
	ImageName     string
	LoadedSymbols bool
}

type SymbolCallback func(symbol *SymbolInfo)

type Manager struct {
	processHandle windows.Handle
	modules       *safemap.M[uint64, *ModuleInfo]
	symbols       *safemap.M[uint64, *SymbolInfo]
	mu            sync.RWMutex
	initialized   bool
	table         *treetable.TreeTable[ModuleInfo]
}

type SYMBOL_INFO struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Index        uint64
	Size         uint64
	ModBase      uint64
	Flags        uint32
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	Name         [1]uint16
}

type IMAGEHLP_LINE64 struct {
	SizeOfStruct uint32
	Key          uint32
	LineNumber   uint32
	Address      uint64
	FileName     uintptr
}

type SYM_ENUMERATESYMBOLS_CALLBACK func(symbolInfo *SYMBOL_INFO, symbolSize uint32, userContext unsafe.Pointer) bool

var (
	dbgHelp                      = syscall.NewLazyDLL("dbghelp.dll")
	procSymInitialize            = dbgHelp.NewProc("SymInitialize")
	procSymCleanup               = dbgHelp.NewProc("SymCleanup")
	procSymLoadModuleExW         = dbgHelp.NewProc("SymLoadModuleExW")
	procSymUnloadModule64        = dbgHelp.NewProc("SymUnloadModule64")
	procSymFromAddr              = dbgHelp.NewProc("SymFromAddr")
	procSymGetModuleInfo64       = dbgHelp.NewProc("SymGetModuleInfo64")
	procSymEnumSymbols           = dbgHelp.NewProc("SymEnumSymbols")
	procSymSetOptions            = dbgHelp.NewProc("SymSetOptions")
	procSymFromName              = dbgHelp.NewProc("SymFromName")
	procSymGetLineFromAddr64     = dbgHelp.NewProc("SymGetLineFromAddr64")
	procSymGetModuleBase64       = dbgHelp.NewProc("SymGetModuleBase64")
	procSymFunctionTableAccess64 = dbgHelp.NewProc("SymFunctionTableAccess64")
)

const (
	SYMOPT_CASE_INSENSITIVE       = 0x00000001
	SYMOPT_UNDNAME                = 0x00000002
	SYMOPT_DEFERRED_LOADS         = 0x00000004
	SYMOPT_NO_CPP                 = 0x00000008
	SYMOPT_LOAD_LINES             = 0x00000010
	SYMOPT_OMIT_MAPPOINTS         = 0x00000020
	SYMOPT_ALLOW_ABSOLUTE_SYMBOLS = 0x00000800
	SYMOPT_FAIL_CRITICAL_ERRORS   = 0x00000200
	SYMOPT_INCLUDE_32BIT_MODULES  = 0x00002000
	SYMOPT_PUBLICS_ONLY           = 0x00004000
	SYMOPT_NO_PUBLICS             = 0x00008000
	SYMOPT_AUTO_PUBLICS           = 0x00010000
	SYMOPT_NO_IMAGE_SEARCH        = 0x00020000
	SYMOPT_SECURE                 = 0x00040000
	SYMOPT_DEBUG                  = 0x80000000
)

func New() api.Interface {
	m := &Manager{
		modules: safemap.New[uint64, *ModuleInfo](),
		symbols: safemap.New[uint64, *SymbolInfo](),
	}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[ModuleInfo]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			modules := m.GetAllModules()
			for _, module := range modules {
				m.table.Root().AddChild(m.table.NewNode(*module))
			}
		},
		JsonName: "Symbol",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Update() error {
	m.table.Root().SetChildren(nil)
	modules := m.GetAllModules()
	for _, module := range modules {
		m.table.Root().AddChild(m.table.NewNode(*module))
	}
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) Initialize(processHandle windows.Handle) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.initialized {
		return nil
	}

	ret, _, err := procSymInitialize.Call(
		uintptr(processHandle),
		0,
		1,
	)
	if ret == 0 {
		return fmt.Errorf("SymInitialize failed: %v", err)
	}

	ret, _, err = procSymSetOptions.Call(
		SYMOPT_UNDNAME | SYMOPT_DEFERRED_LOADS | SYMOPT_LOAD_LINES | SYMOPT_FAIL_CRITICAL_ERRORS,
	)
	if ret == 0 {
		return fmt.Errorf("SymSetOptions failed: %v", err)
	}

	m.processHandle = processHandle
	m.initialized = true
	return nil
}

func (m *Manager) Cleanup() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.initialized {
		return nil
	}

	ret, _, err := procSymCleanup.Call(uintptr(m.processHandle))
	if ret == 0 {
		return fmt.Errorf("SymCleanup failed: %v", err)
	}

	m.initialized = false
	m.modules.Reset()
	m.symbols.Reset()

	return nil
}

func (m *Manager) LoadModuleForProcess(pid uint32, fileHandle windows.Handle, baseAddress uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.initialized {
		return fmt.Errorf("symbol manager not initialized")
	}

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(processHandle)

	return m.loadModule(processHandle, fileHandle, baseAddress)
}

func (m *Manager) LoadModule(pid uint32, fileHandle windows.Handle, baseAddress uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.initialized {
		return fmt.Errorf("symbol manager not initialized")
	}

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(processHandle)

	return m.loadModule(processHandle, fileHandle, baseAddress)
}

func (m *Manager) loadModule(processHandle windows.Handle, fileHandle windows.Handle, baseAddress uint64) error {
	ret, _, err := procSymLoadModuleExW.Call(
		uintptr(processHandle),
		uintptr(fileHandle),
		0,
		0,
		uintptr(baseAddress),
		0,
		0,
		0,
	)
	if ret == 0 {
		return fmt.Errorf("SymLoadModuleExW failed: %v", err)
	}

	moduleInfo, err := m.getModuleInfo(baseAddress)
	if err != nil {
		return err
	}

	m.modules.Update(baseAddress, moduleInfo)
	return nil
}

func (m *Manager) UnloadModule(baseAddress uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.initialized {
		return fmt.Errorf("symbol manager not initialized")
	}

	ret, _, err := procSymUnloadModule64.Call(uintptr(m.processHandle), uintptr(baseAddress))
	if ret == 0 {
		return fmt.Errorf("SymUnloadModule64 failed: %v", err)
	}

	m.modules.Delete(baseAddress)
	return nil
}

func (m *Manager) GetSymbolFromAddress(address uint64) (*SymbolInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.initialized {
		return nil, fmt.Errorf("symbol manager not initialized")
	}

	symbolInfo := &SYMBOL_INFO{
		SizeOfStruct: uint32(unsafe.Sizeof(SYMBOL_INFO{})),
	}

	var displacement uint64
	ret, _, err := procSymFromAddr.Call(
		uintptr(m.processHandle),
		uintptr(address),
		uintptr(unsafe.Pointer(&displacement)),
		uintptr(unsafe.Pointer(symbolInfo)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("SymFromAddr failed: %v", err)
	}

	name := syscall.UTF16ToString(symbolInfo.Name[:symbolInfo.NameLen])
	moduleBase := m.getModuleBase(address)

	symbol := &SymbolInfo{
		Type:       SymbolTypeSymbol,
		Address:    symbolInfo.Address,
		Size:       uint32(symbolInfo.Size),
		Name:       name,
		ModuleName: m.getModuleName(moduleBase),
	}

	m.symbols.Update(address, symbol)
	return symbol, nil
}

func (m *Manager) GetSymbolFromName(name string) (*SymbolInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.initialized {
		return nil, fmt.Errorf("symbol manager not initialized")
	}

	namePtr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	symbolInfo := &SYMBOL_INFO{
		SizeOfStruct: uint32(unsafe.Sizeof(SYMBOL_INFO{})),
	}

	ret, _, err := procSymFromName.Call(
		uintptr(m.processHandle),
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(unsafe.Pointer(symbolInfo)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("SymFromName failed: %v", err)
	}

	symbolName := syscall.UTF16ToString(symbolInfo.Name[:symbolInfo.NameLen])
	moduleBase := m.getModuleBase(symbolInfo.Address)

	symbol := &SymbolInfo{
		Type:       SymbolTypeSymbol,
		Address:    symbolInfo.Address,
		Size:       uint32(symbolInfo.Size),
		Name:       symbolName,
		ModuleName: m.getModuleName(moduleBase),
	}

	m.symbols.Update(symbolInfo.Address, symbol)
	return symbol, nil
}

func (m *Manager) EnumerateSymbols(baseAddress uint64, callback SymbolCallback) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.initialized {
		return fmt.Errorf("symbol manager not initialized")
	}

	enumCallback := func(symbolInfo *SYMBOL_INFO, symbolSize uint32, userContext unsafe.Pointer) bool {
		name := syscall.UTF16ToString(symbolInfo.Name[:symbolInfo.NameLen])
		moduleBase := m.getModuleBase(symbolInfo.Address)

		symbol := &SymbolInfo{
			Type:       SymbolTypeSymbol,
			Address:    symbolInfo.Address,
			Size:       uint32(symbolInfo.Size),
			Name:       name,
			ModuleName: m.getModuleName(moduleBase),
		}

		callback(symbol)
		return true
	}

	ret, _, err := procSymEnumSymbols.Call(
		uintptr(m.processHandle),
		uintptr(baseAddress),
		0,
		syscall.NewCallbackCDecl(enumCallback),
		0,
	)
	if ret == 0 {
		return fmt.Errorf("SymEnumSymbols failed: %v", err)
	}

	return nil
}

func (m *Manager) GetLineFromAddress(address uint64) (*SymbolInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.initialized {
		return nil, fmt.Errorf("symbol manager not initialized")
	}

	line := &IMAGEHLP_LINE64{
		SizeOfStruct: uint32(unsafe.Sizeof(IMAGEHLP_LINE64{})),
	}

	var displacement uint32
	ret, _, err := procSymGetLineFromAddr64.Call(
		uintptr(m.processHandle),
		uintptr(address),
		uintptr(unsafe.Pointer(&displacement)),
		uintptr(unsafe.Pointer(line)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("SymGetLineFromAddr64 failed: %v", err)
	}

	fileName := syscall.UTF16ToString((*[260]uint16)(unsafe.Pointer(line.FileName))[:])

	symbol := &SymbolInfo{
		Type:       SymbolTypeSymbol,
		Address:    address,
		FileName:   fileName,
		LineNumber: line.LineNumber,
		LineOffset: displacement,
	}

	return symbol, nil
}

func (m *Manager) GetModuleInfo(baseAddress uint64) (*ModuleInfo, error) {
	return m.getModuleInfo(baseAddress)
}

func (m *Manager) GetAllModules() []*ModuleInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	modules := make([]*ModuleInfo, 0)
	for _, module := range m.modules.Range() {
		modules = append(modules, module)
	}
	return modules
}

func (m *Manager) GetModuleByAddress(address uint64) *ModuleInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	baseAddress := m.getModuleBase(address)
	module, _ := m.modules.Get(baseAddress)
	return module
}

func (m *Manager) getModuleInfo(baseAddress uint64) (*ModuleInfo, error) {
	type IMAGEHLP_MODULE64 struct {
		SizeOfStruct    uint32
		BaseOfImage     uint64
		ImageSize       uint32
		TimeDateStamp   uint32
		CheckSum        uint32
		NumSyms         uint32
		SymType         uint32
		ModuleName      [32]uint16
		ImageName       [256]uint16
		LoadedImageName [256]uint16
		LoadedPdbName   [256]uint16
		CVSig           uint32
		CVData          [MAX_PATH * 3]uint16
		PdbSig          uint32
		PdbSig70        [16]byte
		PdbAge          uint32
		PdbUnmatched    bool
		DbgUnmatched    bool
		LineNumbers     bool
		GlobalSymbols   bool
		TypeInfo        bool
		SourceIndexed   bool
		PublicSymbols   bool
		Loaded          bool
	}

	const MAX_PATH = 260

	moduleInfo := &IMAGEHLP_MODULE64{
		SizeOfStruct: uint32(unsafe.Sizeof(IMAGEHLP_MODULE64{})),
	}

	ret, _, err := procSymGetModuleInfo64.Call(
		uintptr(m.processHandle),
		uintptr(baseAddress),
		uintptr(unsafe.Pointer(moduleInfo)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("SymGetModuleInfo64 failed: %v", err)
	}

	moduleName := syscall.UTF16ToString(moduleInfo.ModuleName[:])
	imageName := syscall.UTF16ToString(moduleInfo.ImageName[:])

	return &ModuleInfo{
		BaseAddress:   moduleInfo.BaseOfImage,
		Size:          moduleInfo.ImageSize,
		Name:          moduleName,
		ImageName:     imageName,
		LoadedSymbols: moduleInfo.Loaded,
	}, nil
}

func (m *Manager) getModuleBase(address uint64) uint64 {
	ret, _, _ := procSymGetModuleBase64.Call(
		uintptr(m.processHandle),
		uintptr(address),
	)
	return uint64(ret)
}

func (m *Manager) getModuleName(baseAddress uint64) string {
	if module, exists := m.modules.Get(baseAddress); exists {
		return module.Name
	}
	return ""
}

func (m *Manager) IsInitialized() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.initialized
}

func (m *Manager) GetSymbolCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	count := 0
	for range m.symbols.Range() {
		count++
	}
	return count
}

func (m *Manager) GetModuleCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	count := 0
	for range m.modules.Range() {
		count++
	}
	return count
}

func (m *Manager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.symbols.Reset()
}
