package pdbex

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

// D:\ewdk\dist\sdk\dia

type Variant struct {
	VT         uint16
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	Val        uint64
}

const (
	VT_EMPTY    = 0
	VT_NULL     = 1
	VT_I1       = 16
	VT_I2       = 2
	VT_I4       = 3
	VT_I8       = 20
	VT_UI1      = 17
	VT_UI2      = 18
	VT_UI4      = 19
	VT_UI8      = 21
	VT_INT      = 22
	VT_UINT     = 23
	VT_R4       = 4
	VT_R8       = 5
	VT_BOOL     = 11
	VT_BSTR     = 8
	VT_DISPATCH = 9
	VT_UNKNOWN  = 13
)

var (
	ole32            = syscall.NewLazyDLL("ole32.dll")
	oleaut32         = syscall.NewLazyDLL("oleaut32.dll")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	coInitialize     = ole32.NewProc("CoInitialize")
	coUninitialize   = ole32.NewProc("CoUninitialize")
	coCreateInstance = ole32.NewProc("CoCreateInstance")
	coGetClassObject = ole32.NewProc("CoGetClassObject")
	sysAllocString   = oleaut32.NewProc("SysAllocString")
	sysFreeString    = oleaut32.NewProc("SysFreeString")
	loadLibrary      = kernel32.NewProc("LoadLibraryW")
	getProcAddress   = kernel32.NewProc("GetProcAddress")
	freeLibrary      = kernel32.NewProc("FreeLibrary")
)

var comInitialized bool

func initCOM() error {
	if comInitialized {
		return nil
	}

	ret, _, _ := coInitialize.Call(0)
	hr := int32(ret)
	if hr < 0 && hr != -2147417850 {
		return fmt.Errorf("CoInitialize failed: 0x%08X", hr)
	}

	comInitialized = true
	return nil
}

func uninitCOM() {
	if comInitialized {
		coUninitialize.Call()
		comInitialized = false
	}
}

func bstrToString(bstr uintptr) string {
	if bstr == 0 {
		return ""
	}

	ptr := (*uint16)(unsafe.Pointer(bstr))
	length := 0
	for p := ptr; *p != 0; p = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 2)) {
		length++
	}

	if length == 0 {
		return ""
	}

	buf := make([]uint16, length)
	for i := 0; i < length; i++ {
		buf[i] = *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i*2)))
	}

	return syscall.UTF16ToString(buf)
}

func stringToBstr(s string) uintptr {
	if s == "" {
		return 0
	}

	utf16, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return 0
	}

	ret, _, _ := sysAllocString.Call(uintptr(unsafe.Pointer(utf16)))
	return ret
}

func freeBstr(bstr uintptr) {
	if bstr != 0 {
		sysFreeString.Call(bstr)
	}
}

func (p *PDB) parseWithDIA() error {
	if err := initCOM(); err != nil {
		return fmt.Errorf("failed to initialize COM: %w", err)
	}

	dia, err := newDIASession(p.path)
	if err != nil {
		return fmt.Errorf("failed to create DIA session for: %s: %w", p.path, err)
	}

	p.machineType = dia.getMachineType()
	p.language = dia.getLanguage()
	p.dia = dia

	dia.enumerateSymbols(func(sym *Symbol) {
		p.registerSymbol(sym)
	})

	dia.enumerateFunctions(func(fn *FunctionInfo) {
		p.registerFunction(fn)
	})

	return nil
}

type diaSession struct {
	dataSource uintptr
	session    uintptr
	global     uintptr
	dllHandle  uintptr
	visited    map[uintptr]bool
	depth      int
}

const maxDepth = 5

func newDIASession(pdbPath string) (*diaSession, error) {
	fmt.Printf("newDIASession called with path: %s\n", pdbPath)
	dia := &diaSession{
		visited: make(map[uintptr]bool),
	}

	clsidDiaSource := syscall.GUID{
		Data1: 0xe6756135,
		Data2: 0x1e65,
		Data3: 0x4d17,
		Data4: [8]byte{0x85, 0x76, 0x61, 0x07, 0x61, 0x39, 0x8c, 0x3c},
	}

	iidIDataSource := syscall.GUID{
		Data1: 0x79f1bb5f,
		Data2: 0xb66e,
		Data3: 0x48e5,
		Data4: [8]byte{0xb6, 0xa9, 0x15, 0x45, 0xc3, 0x23, 0xca, 0x3d},
	}

	fmt.Printf("Calling CoCreateInstance...\n")
	var dataSource uintptr
	ret, _, _ := coCreateInstance.Call(
		uintptr(unsafe.Pointer(&clsidDiaSource)),
		0,
		1,
		uintptr(unsafe.Pointer(&iidIDataSource)),
		uintptr(unsafe.Pointer(&dataSource)),
	)
	fmt.Printf("CoCreateInstance returned: 0x%08X, dataSource: 0x%X\n", ret, dataSource)

	if ret != 0 {
		dllHandle := loadMsdia140()
		if dllHandle == 0 {
			return nil, fmt.Errorf("failed to load msdia140.dll")
		}
		dia.dllHandle = dllHandle

		iidIClassFactory := syscall.GUID{
			Data1: 0x00000001,
			Data2: 0x0000,
			Data3: 0x0000,
			Data4: [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46},
		}

		dllGetClassObjectName, err := syscall.BytePtrFromString("DllGetClassObject")
		if err != nil {
			return nil, fmt.Errorf("failed to get DllGetClassObject name: %w", err)
		}

		dllGetClassObjectAddr, _, lastErr := getProcAddress.Call(dllHandle, uintptr(unsafe.Pointer(dllGetClassObjectName)))
		if dllGetClassObjectAddr == 0 {
			return nil, fmt.Errorf("failed to get DllGetClassObject address: %v", lastErr)
		}
		fmt.Printf("DllGetClassObject address: 0x%X\n", dllGetClassObjectAddr)
		fmt.Printf("CLSID: %v\n", clsidDiaSource)
		fmt.Printf("IID IClassFactory: %v\n", iidIClassFactory)

		var classFactory uintptr
		ret, _, _ = syscall.SyscallN(dllGetClassObjectAddr,
			uintptr(unsafe.Pointer(&clsidDiaSource)),
			uintptr(unsafe.Pointer(&iidIClassFactory)),
			uintptr(unsafe.Pointer(&classFactory)),
		)
		fmt.Printf("DllGetClassObject returned: 0x%08X, classFactory: 0x%X\n", ret, classFactory)
		if ret != 0 || classFactory == 0 {
			return nil, fmt.Errorf("DllGetClassObject failed: 0x%08X", ret)
		}
		defer release(classFactory)

		factoryVtable := getVtable(classFactory)
		createInstance := factoryVtable[3]

		ret, _, _ = syscall.SyscallN(createInstance,
			classFactory,
			0,
			uintptr(unsafe.Pointer(&iidIDataSource)),
			uintptr(unsafe.Pointer(&dataSource)),
		)
		if ret != 0 || dataSource == 0 {
			return nil, fmt.Errorf("CreateInstance failed: 0x%08X", ret)
		}
	}

	dia.dataSource = dataSource

	if dataSource == 0 {
		return nil, fmt.Errorf("dataSource is null")
	}

	fmt.Printf("dataSource is valid: 0x%X\n", dataSource)

	pdbPathBstr := stringToBstr(pdbPath)
	defer freeBstr(pdbPathBstr)

	fmt.Printf("pdbPathBstr: 0x%X\n", pdbPathBstr)

	vtable := getVtable(dataSource)
	fmt.Printf("dataSource vtable[0-9]:\n")
	for i := range 10 {
		fmt.Printf("  [%d] = 0x%X\n", i, vtable[i])
	}
	loadDataFromPdb := vtable[VTDataSource_loadDataFromPdb]

	fmt.Printf("Calling loadDataFromPdb at 0x%X\n", loadDataFromPdb)
	ret, _, _ = syscall.SyscallN(loadDataFromPdb, dataSource, pdbPathBstr)
	fmt.Printf("loadDataFromPdb returned: 0x%08X\n", ret)
	if ret != 0 {
		dia.close()
		return nil, fmt.Errorf("loadDataFromPdb failed: 0x%08X", ret)
	}

	var session uintptr
	fmt.Printf("Calling openSession at 0x%X\n", vtable[VTDataSource_openSession])
	ret, _, _ = syscall.SyscallN(vtable[VTDataSource_openSession], dataSource, uintptr(unsafe.Pointer(&session)))
	fmt.Printf("openSession returned: 0x%08X, session: 0x%X\n", ret, session)
	if ret != 0 {
		dia.close()
		return nil, fmt.Errorf("openSession failed: 0x%08X", ret)
	}

	dia.session = session

	sessionVtable := getVtable(session)
	var global uintptr
	ret, _, _ = syscall.SyscallN(sessionVtable[VTSession_get_globalScope], session, uintptr(unsafe.Pointer(&global)))
	if ret != 0 {
		dia.close()
		return nil, fmt.Errorf("getGlobalScope failed: 0x%08X", ret)
	}

	dia.global = global

	return dia, nil
}

func loadMsdia140() uintptr {
	var dllData []byte
	var dllName string
	var symSrvName string

	if unsafe.Sizeof(uintptr(0)) == 8 {
		dllData = msdia140x64
		dllName = "msdia140.dll"
		symSrvName = "symsrv.dll"
	} else {
		dllData = msdia140x86
		dllName = "msdia140.dll"
		symSrvName = "symsrv.dll"
	}

	tempDir := os.TempDir()
	diaDir := filepath.Join(tempDir, "dia")

	if _, err := os.Stat(diaDir); os.IsNotExist(err) {
		if err := os.MkdirAll(diaDir, 0o755); err != nil {
			fmt.Printf("Failed to create dia directory: %v\n", err)
			return 0
		}
	}

	dllPath := filepath.Join(diaDir, dllName)
	symSrvPath := filepath.Join(diaDir, symSrvName)

	if _, err := os.Stat(dllPath); os.IsNotExist(err) {
		if err := os.WriteFile(dllPath, dllData, 0o644); err != nil {
			fmt.Printf("Failed to write msdia140.dll: %v\n", err)
			return 0
		}
	}

	if symSrvData := getSymSrvDLL(); len(symSrvData) > 0 {
		if _, err := os.Stat(symSrvPath); os.IsNotExist(err) {
			if err := os.WriteFile(symSrvPath, symSrvData, 0o644); err != nil {
				fmt.Printf("Failed to write symsrv.dll: %v\n", err)
			}
		}
	}

	fmt.Printf("Loading DLL from: %s\n", dllPath)

	utf16Path, err := syscall.UTF16PtrFromString(dllPath)
	if err != nil {
		fmt.Printf("Failed to convert path to UTF16: %v\n", err)
		return 0
	}

	ret, _, lastErr := loadLibrary.Call(uintptr(unsafe.Pointer(utf16Path)))
	if ret == 0 {
		fmt.Printf("LoadLibrary failed: %v\n", lastErr)
	} else {
		fmt.Printf("LoadLibrary succeeded, handle: 0x%X\n", ret)
	}
	return ret
}

func getSymSrvDLL() []byte {
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return symsrvX64
	}
	return symsrvX86
}

func (d *diaSession) close() {
	if d.global != 0 {
		release(d.global)
	}
	if d.session != 0 {
		release(d.session)
	}
	if d.dataSource != 0 {
		release(d.dataSource)
	}
}

func release(ptr uintptr) {
	if ptr == 0 {
		return
	}
	vtable := getVtable(ptr)
	releaseFunc := vtable[VTDataSource_Release]
	if releaseFunc == 0 {
		return
	}
	syscall.SyscallN(releaseFunc, ptr)
}

func getVtable(ptr uintptr) *[200]uintptr {
	vtablePtr := *(*uintptr)(unsafe.Pointer(ptr))
	return (*[200]uintptr)(unsafe.Pointer(vtablePtr))
}

func (d *diaSession) getMachineType() uint16 {
	if d.global == 0 {
		return 0
	}

	vtable := getVtable(d.global)
	getMachineType := VTSymbol_get_machineType.Get(vtable)

	var machineType uint32
	syscall.SyscallN(getMachineType, d.global, uintptr(unsafe.Pointer(&machineType)))

	return uint16(machineType)
}

func (d *diaSession) getLanguage() uint32 {
	if d.global == 0 {
		return 0
	}

	vtable := getVtable(d.global)
	getLanguage := VTSymbol_get_language.Get(vtable)

	var language uint32
	syscall.SyscallN(getLanguage, d.global, uintptr(unsafe.Pointer(&language)))

	return language
}

func (d *diaSession) enumerateSymbols(callback func(*Symbol)) {
	if d.global == 0 {
		return
	}

	d.enumerateSymbolsByTag(SymTagEnumType, callback)
	d.enumerateSymbolsByTag(SymTagUDT, callback)
	d.enumerateSymbolsByTag(SymTagTypedef, callback)
}

func (d *diaSession) enumerateSymbolsByTag(tag SymTag, callback func(*Symbol)) {
	if d.global == 0 {
		return
	}

	vtable := getVtable(d.global)
	findChildren := VTSymbol_findChildren.Get(vtable)

	var enumSymbols uintptr
	ret, _, _ := syscall.SyscallN(findChildren, d.global, uintptr(tag), 0, 0, uintptr(unsafe.Pointer(&enumSymbols)))
	if ret != 0 || enumSymbols == 0 {
		return
	}
	defer release(enumSymbols)

	enumVtable := getVtable(enumSymbols)
	next := enumVtable[VTEnumSymbols_Next]
	getCount := enumVtable[VTEnumSymbols_get_Count]

	var count uint32
	syscall.SyscallN(getCount, enumSymbols, uintptr(unsafe.Pointer(&count)))

	for i := uint32(0); i < count; i++ {
		var symbol uintptr
		var fetched uint32
		ret, _, _ := syscall.SyscallN(next, enumSymbols, 1, uintptr(unsafe.Pointer(&symbol)), uintptr(unsafe.Pointer(&fetched)))
		if ret != 0 || fetched == 0 {
			continue
		}

		sym := d.parseSymbol(symbol)
		if sym != nil {
			callback(sym)
		}

		release(symbol)
	}
}

func (d *diaSession) enumerateFunctions(callback func(*FunctionInfo)) {
	if d.global == 0 {
		return
	}

	vtable := getVtable(d.global)
	findChildren := VTSymbol_findChildren.Get(vtable)

	var enumSymbols uintptr
	ret, _, _ := syscall.SyscallN(findChildren, d.global, uintptr(SymTagFunction), 0, 0, uintptr(unsafe.Pointer(&enumSymbols)))
	if ret != 0 || enumSymbols == 0 {
		return
	}
	defer release(enumSymbols)

	enumVtable := getVtable(enumSymbols)
	next := enumVtable[VTEnumSymbols_Next]
	getCount := enumVtable[VTEnumSymbols_get_Count]

	var count uint32
	syscall.SyscallN(getCount, enumSymbols, uintptr(unsafe.Pointer(&count)))

	for i := uint32(0); i < count; i++ {
		var symbol uintptr
		var fetched uint32
		ret, _, _ := syscall.SyscallN(next, enumSymbols, 1, uintptr(unsafe.Pointer(&symbol)), uintptr(unsafe.Pointer(&fetched)))
		if ret != 0 || fetched == 0 {
			continue
		}

		fn := d.parseFunction(symbol)
		if fn != nil {
			callback(fn)
		}

		release(symbol)
	}
}

func (d *diaSession) parseSymbol(symbol uintptr) *Symbol {
	if symbol == 0 {
		return nil
	}

	if d.depth > maxDepth {
		return &Symbol{Name: "<max depth exceeded>"}
	}
	d.depth++
	defer func() { d.depth-- }()

	if d.visited[symbol] {
		return &Symbol{Name: "<circular reference>"}
	}
	d.visited[symbol] = true
	defer delete(d.visited, symbol)

	vtable := getVtable(symbol)
	getSymTag := VTSymbol_get_symTag.Get(vtable)
	getName := VTSymbol_get_name.Get(vtable)
	getTypeId := VTSymbol_get_typeId.Get(vtable)
	getLength := VTSymbol_get_length.Get(vtable)
	getConstType := VTSymbol_get_constType.Get(vtable)
	getVolatileType := VTSymbol_get_volatileType.Get(vtable)
	getDataKind := VTSymbol_get_dataKind.Get(vtable)
	getBaseType := VTSymbol_get_baseType.Get(vtable)

	sym := &Symbol{}

	var tag uint32
	syscall.SyscallN(getSymTag, symbol, uintptr(unsafe.Pointer(&tag)))
	sym.Tag = SymTag(tag)

	var nameBstr uintptr
	syscall.SyscallN(getName, symbol, uintptr(unsafe.Pointer(&nameBstr)))
	sym.Name = bstrToString(nameBstr)
	freeBstr(nameBstr)

	var typeId uint32
	syscall.SyscallN(getTypeId, symbol, uintptr(unsafe.Pointer(&typeId)))
	sym.TypeId = typeId

	var length uint64
	syscall.SyscallN(getLength, symbol, uintptr(unsafe.Pointer(&length)))
	sym.Size = uint32(length)

	var isConst int32
	syscall.SyscallN(getConstType, symbol, uintptr(unsafe.Pointer(&isConst)))
	sym.IsConst = isConst != 0

	var isVolatile int32
	syscall.SyscallN(getVolatileType, symbol, uintptr(unsafe.Pointer(&isVolatile)))
	sym.IsVolatile = isVolatile != 0

	var dataKind uint32
	syscall.SyscallN(getDataKind, symbol, uintptr(unsafe.Pointer(&dataKind)))
	sym.DataKind = DataKind(dataKind)

	var baseType uint32
	syscall.SyscallN(getBaseType, symbol, uintptr(unsafe.Pointer(&baseType)))
	sym.BaseType = BasicType(baseType)

	switch sym.Tag {
	case SymTagUDT:
		sym.Udt = d.parseUdt(symbol)
	case SymTagEnumType:
		sym.Enum = d.parseEnum(symbol)
	case SymTagPointerType:
		sym.Pointer = d.parsePointer(symbol)
	case SymTagArrayType:
		sym.Array = d.parseArray(symbol)
	case SymTagTypedef:
		sym.Typedef = d.parseTypedef(symbol)
	}

	return sym
}

func (d *diaSession) parseUdt(symbol uintptr) *SymbolUdt {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getUdtKind := VTSymbol_get_udtKind.Get(vtable)
	findChildren := VTSymbol_findChildren.Get(vtable)

	var kind uint32
	syscall.SyscallN(getUdtKind, symbol, uintptr(unsafe.Pointer(&kind)))

	udt := &SymbolUdt{
		Kind: UdtKind(kind),
	}

	var enumSymbols uintptr
	ret, _, _ := syscall.SyscallN(findChildren, symbol, uintptr(SymTagData), 0, 0, uintptr(unsafe.Pointer(&enumSymbols)))
	if ret != 0 || enumSymbols == 0 {
		return udt
	}
	defer release(enumSymbols)

	enumVtable := getVtable(enumSymbols)
	next := enumVtable[VTEnumSymbols_Next]
	getCount := enumVtable[VTEnumSymbols_get_Count]

	var count uint32
	syscall.SyscallN(getCount, enumSymbols, uintptr(unsafe.Pointer(&count)))

	udt.Fields = make([]*SymbolUdtField, 0, count)

	for i := uint32(0); i < count; i++ {
		var childSymbol uintptr
		var fetched uint32
		ret, _, _ := syscall.SyscallN(next, enumSymbols, 1, uintptr(unsafe.Pointer(&childSymbol)), uintptr(unsafe.Pointer(&fetched)))
		if ret != 0 || fetched == 0 {
			continue
		}

		field := d.parseUdtField(childSymbol)
		if field != nil {
			udt.Fields = append(udt.Fields, field)
		}

		release(childSymbol)
	}

	udt.FieldCount = uint32(len(udt.Fields))
	return udt
}

func (d *diaSession) parseUdtField(symbol uintptr) *SymbolUdtField {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getName := VTSymbol_get_name.Get(vtable)
	getOffset := VTSymbol_get_offset.Get(vtable)
	getLength := VTSymbol_get_length.Get(vtable)
	getBitPosition := VTSymbol_get_bitPosition.Get(vtable)
	getType := VTSymbol_get_type.Get(vtable)

	field := &SymbolUdtField{}

	var nameBstr uintptr
	syscall.SyscallN(getName, symbol, uintptr(unsafe.Pointer(&nameBstr)))
	field.Name = bstrToString(nameBstr)
	freeBstr(nameBstr)

	var offset int32
	syscall.SyscallN(getOffset, symbol, uintptr(unsafe.Pointer(&offset)))
	field.Offset = uint32(offset)

	var length uint64
	syscall.SyscallN(getLength, symbol, uintptr(unsafe.Pointer(&length)))
	field.Bits = uint32(length)

	var bitPosition uint32
	syscall.SyscallN(getBitPosition, symbol, uintptr(unsafe.Pointer(&bitPosition)))
	field.BitPosition = bitPosition

	var fieldType uintptr
	ret, _, _ := syscall.SyscallN(getType, symbol, uintptr(unsafe.Pointer(&fieldType)))
	if ret == 0 && fieldType != 0 {
		field.Type = d.parseSymbol(fieldType)
		release(fieldType)
	}

	return field
}

func (d *diaSession) parseEnum(symbol uintptr) *SymbolEnum {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	findChildren := VTSymbol_findChildren.Get(vtable)

	enum := &SymbolEnum{}

	var enumSymbols uintptr
	ret, _, _ := syscall.SyscallN(findChildren, symbol, uintptr(SymTagNull), 0, 0, uintptr(unsafe.Pointer(&enumSymbols)))
	if ret != 0 || enumSymbols == 0 {
		return enum
	}
	defer release(enumSymbols)

	enumVtable := getVtable(enumSymbols)
	next := enumVtable[VTEnumSymbols_Next]
	getCount := enumVtable[VTEnumSymbols_get_Count]

	var count uint32
	syscall.SyscallN(getCount, enumSymbols, uintptr(unsafe.Pointer(&count)))

	enum.Fields = make([]*SymbolEnumField, 0, count)

	for i := uint32(0); i < count; i++ {
		var childSymbol uintptr
		var fetched uint32
		ret, _, _ := syscall.SyscallN(next, enumSymbols, 1, uintptr(unsafe.Pointer(&childSymbol)), uintptr(unsafe.Pointer(&fetched)))
		if ret != 0 || fetched == 0 {
			continue
		}

		field := d.parseEnumField(childSymbol)
		if field != nil {
			enum.Fields = append(enum.Fields, field)
		}

		release(childSymbol)
	}

	enum.FieldCount = uint32(len(enum.Fields))
	return enum
}

func (d *diaSession) parseEnumField(symbol uintptr) *SymbolEnumField {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getName := VTSymbol_get_name.Get(vtable)
	getValue := VTSymbol_get_value.Get(vtable)

	field := &SymbolEnumField{}

	var nameBstr uintptr
	syscall.SyscallN(getName, symbol, uintptr(unsafe.Pointer(&nameBstr)))
	field.Name = bstrToString(nameBstr)
	freeBstr(nameBstr)

	var value Variant
	syscall.SyscallN(getValue, symbol, uintptr(unsafe.Pointer(&value)))
	field.Value = variantToInterface(value)

	return field
}

func (d *diaSession) parsePointer(symbol uintptr) *SymbolPointer {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getType := VTSymbol_get_type.Get(vtable)
	getReference := VTSymbol_get_reference.Get(vtable)

	ptr := &SymbolPointer{}

	var isRef int32
	syscall.SyscallN(getReference, symbol, uintptr(unsafe.Pointer(&isRef)))
	ptr.IsReference = isRef != 0

	var targetType uintptr
	ret, _, _ := syscall.SyscallN(getType, symbol, uintptr(unsafe.Pointer(&targetType)))
	if ret == 0 && targetType != 0 {
		ptr.Type = d.parseSymbol(targetType)
		release(targetType)
	}

	return ptr
}

func (d *diaSession) parseArray(symbol uintptr) *SymbolArray {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getType := VTSymbol_get_type.Get(vtable)
	getCount := VTSymbol_get_count.Get(vtable)

	arr := &SymbolArray{}

	var count uint32
	syscall.SyscallN(getCount, symbol, uintptr(unsafe.Pointer(&count)))
	arr.ElementCount = count

	var elementType uintptr
	ret, _, _ := syscall.SyscallN(getType, symbol, uintptr(unsafe.Pointer(&elementType)))
	if ret == 0 && elementType != 0 {
		arr.ElementType = d.parseSymbol(elementType)
		release(elementType)
	}

	return arr
}

func (d *diaSession) parseTypedef(symbol uintptr) *SymbolTypedef {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getType := VTSymbol_get_type.Get(vtable)

	typedef := &SymbolTypedef{}

	var underlyingType uintptr
	ret, _, _ := syscall.SyscallN(getType, symbol, uintptr(unsafe.Pointer(&underlyingType)))
	if ret == 0 && underlyingType != 0 {
		typedef.Type = d.parseSymbol(underlyingType)
		release(underlyingType)
	}

	return typedef
}

func (d *diaSession) parseFunction(symbol uintptr) *FunctionInfo {
	if symbol == 0 {
		return nil
	}

	vtable := getVtable(symbol)
	getName := VTSymbol_get_name.Get(vtable)
	getRelativeVirtualAddress := VTSymbol_get_relativeVirtualAddress.Get(vtable)
	getLength := VTSymbol_get_length.Get(vtable)

	fn := &FunctionInfo{}

	var nameBstr uintptr
	syscall.SyscallN(getName, symbol, uintptr(unsafe.Pointer(&nameBstr)))
	fn.Name = bstrToString(nameBstr)
	freeBstr(nameBstr)

	var rva uint32
	syscall.SyscallN(getRelativeVirtualAddress, symbol, uintptr(unsafe.Pointer(&rva)))
	fn.Address = uint64(rva)

	var length uint64
	syscall.SyscallN(getLength, symbol, uintptr(unsafe.Pointer(&length)))
	fn.Size = uint32(length)

	return fn
}

func variantToInterface(v Variant) any {
	switch v.VT {
	case VT_I1:
		return int8(v.Val)
	case VT_I2:
		return int16(v.Val)
	case VT_I4:
		return int32(v.Val)
	case VT_I8:
		return int64(v.Val)
	case VT_UI1:
		return uint8(v.Val)
	case VT_UI2:
		return uint16(v.Val)
	case VT_UI4:
		return uint32(v.Val)
	case VT_UI8:
		return uint64(v.Val)
	case VT_INT:
		return int(v.Val)
	case VT_UINT:
		return uint(v.Val)
	default:
		return v.Val
	}
}

func (d *diaSession) FindSourceLineByRVA(rva uint32) (string, uint32, bool) {
	if d.session == 0 {
		return "", 0, false
	}

	vtable := getVtable(d.session)
	findLinesByRVA := VTSession_findLinesByRVA.Get(vtable)

	var enumLineNumbers uintptr
	ret, _, _ := syscall.SyscallN(findLinesByRVA, d.session, uintptr(rva), 1, uintptr(unsafe.Pointer(&enumLineNumbers)))
	if ret != 0 || enumLineNumbers == 0 {
		return "", 0, false
	}
	defer release(enumLineNumbers)

	enumVtable := getVtable(enumLineNumbers)
	next := enumVtable[VTEnumLineNumbers_Next]

	var lineNumber uintptr
	var fetched uint32
	ret, _, _ = syscall.SyscallN(next, enumLineNumbers, 1, uintptr(unsafe.Pointer(&lineNumber)), uintptr(unsafe.Pointer(&fetched)))
	if ret != 0 || fetched == 0 || lineNumber == 0 {
		return "", 0, false
	}
	defer release(lineNumber)

	lineVtable := getVtable(lineNumber)
	getSourceFile := VTLineNumber_get_sourceFile.Get(lineVtable)
	getLineNumber := VTLineNumber_get_lineNumber.Get(lineVtable)

	var lineNum uint32
	syscall.SyscallN(getLineNumber, lineNumber, uintptr(unsafe.Pointer(&lineNum)))

	var sourceFile uintptr
	ret, _, _ = syscall.SyscallN(getSourceFile, lineNumber, uintptr(unsafe.Pointer(&sourceFile)))
	if ret != 0 || sourceFile == 0 {
		return "", lineNum, true
	}
	defer release(sourceFile)

	fileVtable := getVtable(sourceFile)
	getFileName := VTSourceFile_get_fileName.Get(fileVtable)

	var fileNameBstr uintptr
	syscall.SyscallN(getFileName, sourceFile, uintptr(unsafe.Pointer(&fileNameBstr)))
	fileName := bstrToString(fileNameBstr)
	freeBstr(fileNameBstr)

	return fileName, lineNum, true
}

func (d *diaSession) FindSourceLineByVA(va uint64) (string, uint32, bool) {
	if d.session == 0 {
		return "", 0, false
	}

	vtable := getVtable(d.session)
	findLinesByVA := VTSession_findLinesByVA.Get(vtable)

	var enumLineNumbers uintptr
	ret, _, _ := syscall.SyscallN(findLinesByVA, d.session, uintptr(va), 1, uintptr(unsafe.Pointer(&enumLineNumbers)))
	if ret != 0 || enumLineNumbers == 0 {
		return "", 0, false
	}
	defer release(enumLineNumbers)

	enumVtable := getVtable(enumLineNumbers)
	next := enumVtable[VTEnumLineNumbers_Next]

	var lineNumber uintptr
	var fetched uint32
	ret, _, _ = syscall.SyscallN(next, enumLineNumbers, 1, uintptr(unsafe.Pointer(&lineNumber)), uintptr(unsafe.Pointer(&fetched)))
	if ret != 0 || fetched == 0 || lineNumber == 0 {
		return "", 0, false
	}
	defer release(lineNumber)

	lineVtable := getVtable(lineNumber)
	getSourceFile := VTLineNumber_get_sourceFile.Get(lineVtable)
	getLineNumber := VTLineNumber_get_lineNumber.Get(lineVtable)

	var lineNum uint32
	syscall.SyscallN(getLineNumber, lineNumber, uintptr(unsafe.Pointer(&lineNum)))

	var sourceFile uintptr
	ret, _, _ = syscall.SyscallN(getSourceFile, lineNumber, uintptr(unsafe.Pointer(&sourceFile)))
	if ret != 0 || sourceFile == 0 {
		return "", lineNum, true
	}
	defer release(sourceFile)

	fileVtable := getVtable(sourceFile)
	getFileName := VTSourceFile_get_fileName.Get(fileVtable)

	var fileNameBstr uintptr
	syscall.SyscallN(getFileName, sourceFile, uintptr(unsafe.Pointer(&fileNameBstr)))
	fileName := bstrToString(fileNameBstr)
	freeBstr(fileNameBstr)

	return fileName, lineNum, true
}
