package Imports

import (
	"github.com/ddkwork/hyperdbgui/SDK/Headers"
	"unsafe"
)

func (s *symbol) SymSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	value := Call(api.Proc(SymSetTextMessageCallback), uintptr(Handler))
	//DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (s *symbol) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64 {
	value := Call(api.Proc(SymConvertNameToAddress),
		uintptr(unsafe.Pointer(FunctionOrVariableName)),
		uintptr(unsafe.Pointer(WasFound)),
	)
	//DecodeErrorCode(value)
	return uint64(value)
}

func (s *symbol) SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32 {
	value := Call(api.Proc(SymLoadFileSymbol),
		uintptr(BaseAddress),
		uintptr(unsafe.Pointer(PdbFileName)),
	)
	//DecodeErrorCode(value)
	return uint32(value)
}

func (s *symbol) SymUnloadAllSymbols() uint32 {
	value := Call(api.Proc(SymUnloadAllSymbols))
	//DecodeErrorCode(value)
	return uint32(value)
}

func (s *symbol) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	value := Call(api.Proc(SymUnloadModuleSymbol), uintptr(unsafe.Pointer(ModuleName)))
	//DecodeErrorCode(value)
	return uint32(value)
}

func (s *symbol) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	value := Call(api.Proc(SymSearchSymbolForMask), uintptr(unsafe.Pointer(SearchMask)))
	//DecodeErrorCode(value)
	return uint32(value)
}

func (s *symbol) SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	value := Call(api.Proc(SymGetFieldOffset),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(FieldName)),
		uintptr(unsafe.Pointer(FieldOffset)),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	value := Call(api.Proc(SymGetDataTypeSize),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(TypeSize)),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	value := Call(api.Proc(SymCreateSymbolTableForDisassembler), uintptr(CallbackFunction))
	//DecodeErrorCode(value) //todo need decode r2 as error code ?
	return value == 0
}

func (s *symbol) SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	value := Call(api.Proc(SymConvertFileToPdbPath),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(ResultPath)),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymConvertFileToPdbFileAndGuidAndAgeDetails(
	LocalFilePath,
	PdbFilePath,
	GuidAndAgeDetails *int8) {
	value := Call(api.Proc(SymConvertFileToPdbFileAndGuidAndAgeDetails),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(PdbFilePath)),
		uintptr(unsafe.Pointer(GuidAndAgeDetails)),
	)
	DecodeErrorCode(value)
	return
}

func (s *symbol) SymbolInitLoad(
	BufferToStoreDetails Headers.PVOID,
	StoredLength uint32,
	DownloadIfAvailable bool,
	SymbolPath *int8,
	IsSilentLoad bool,
) bool {
	value := Call(api.Proc(SymbolInitLoad),
		uintptr(BufferToStoreDetails),
		uintptr(StoredLength),
		Bool2UintPtr(DownloadIfAvailable),
		uintptr(unsafe.Pointer(SymbolPath)),
		Bool2UintPtr(IsSilentLoad),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymShowDataBasedOnSymbolTypes(
	TypeName *int8,
	Address uint64,
	IsStruct bool,
	BufferAddress Headers.PVOID,
	AdditionalParameters *int8,
) bool {
	value := Call(api.Proc(SymShowDataBasedOnSymbolTypes),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(Address),
		Bool2UintPtr(IsStruct),
		uintptr(BufferAddress),
		uintptr(unsafe.Pointer(AdditionalParameters)),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymQuerySizeof(
	StructNameOrTypeName *int8,
	SizeOfField *uint32) bool {
	value := Call(api.Proc(SymQuerySizeof),
		uintptr(unsafe.Pointer(StructNameOrTypeName)),
		uintptr(unsafe.Pointer(SizeOfField)),
	)
	//DecodeErrorCode(value)
	return value == 0
}

func (s *symbol) SymCastingQueryForFiledsAndTypes(
	StructName *int8,
	FiledOfStructName *int8,
	IsStructNamePointerOrNot bool,
	IsFiledOfStructNamePointerOrNot bool,
	NewStructOrTypeName **int8,
	OffsetOfFieldFromTop *uint32,
	SizeOfField *uint32,
) bool {
	value := Call(api.Proc(SymCastingQueryForFiledsAndTypes),
		uintptr(unsafe.Pointer(StructName)),
		uintptr(unsafe.Pointer(FiledOfStructName)),
		Bool2UintPtr(IsStructNamePointerOrNot),
		Bool2UintPtr(IsFiledOfStructNamePointerOrNot),
		uintptr(unsafe.Pointer(NewStructOrTypeName)), //todo check this.pointHyperDbgSymImports.h
		uintptr(unsafe.Pointer(OffsetOfFieldFromTop)),
		uintptr(unsafe.Pointer(SizeOfField)),
	)
	//DecodeErrorCode(value)
	return value == 0
}
func (s *script) ScriptEngineParse(str *int8) PSYMBOL_BUFFER {
	value := Call(api.Proc(ScriptEngineParse), uintptr(unsafe.Pointer(str)))
	//DecodeErrorCode(value)
	return PSYMBOL_BUFFER(unsafe.Pointer(value))
}

func (s *script) PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID {
	value := Call(api.Proc(PrintSymbolBuffer), uintptr(unsafe.Pointer(SymbolBuffer)))
	//DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (s *script) PrintSymbol(Symbol PSYMBOL) Headers.VOID {
	value := Call(api.Proc(PrintSymbol), uintptr(unsafe.Pointer(Symbol)))
	//DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (s *script) RemoveSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID {
	value := Call(api.Proc(RemoveSymbolBuffer), uintptr(unsafe.Pointer(SymbolBuffer)))
	//DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}
