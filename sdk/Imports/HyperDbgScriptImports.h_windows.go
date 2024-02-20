package Imports

import (
	"unsafe"

	"github.com/ddkwork/hyperdbgui/sdk/Headers"
)

func (s *script) ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	value := Call(api.Proc(ScriptEngineSetTextMessageCallback), uintptr(Handler))
	// DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (s *script) ScriptEngineSymbolAbortLoading() Headers.VOID {
	value := Call(api.Proc(ScriptEngineSymbolAbortLoading))
	// DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (s *script) ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64 {
	// WasFound need input a point value ?
	value := Call(api.Proc(ScriptEngineConvertNameToAddress), uintptr(unsafe.Pointer(FunctionOrVariableName)), Bool2UintPtr(WasFound))
	// DecodeErrorCode(value)
	return uint64(value)
}

func (s *script) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32 {
	value := Call(api.Proc(ScriptEngineLoadFileSymbol), uintptr(BaseAddress), uintptr(unsafe.Pointer(PdbFileName)))
	// DecodeErrorCode(value)
	return uint32(value)
}

func (s *script) ScriptEngineUnloadAllSymbols() uint32 {
	value := Call(api.Proc(ScriptEngineUnloadAllSymbols))
	// DecodeErrorCode(value)
	return uint32(value)
}

func (s *script) ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32 {
	value := Call(api.Proc(ScriptEngineUnloadModuleSymbol), uintptr(unsafe.Pointer(ModuleName)))
	// DecodeErrorCode(value)
	return uint32(value)
}

func (s *script) ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32 {
	value := Call(api.Proc(ScriptEngineSearchSymbolForMask), uintptr(unsafe.Pointer(SearchMask)))
	// DecodeErrorCode(value)
	return uint32(value)
}

func (s *script) ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	value := Call(api.Proc(ScriptEngineGetFieldOffset),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(FieldName)),
		uintptr(unsafe.Pointer(FieldOffset)),
	)
	// DecodeErrorCode(value)
	return value == 0
}

func (s *script) ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	value := Call(api.Proc(ScriptEngineGetDataTypeSize),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(TypeSize)),
	)
	return value == 0
}

func (s *script) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	value := Call(api.Proc(ScriptEngineCreateSymbolTableForDisassembler), uintptr(CallbackFunction))
	// DecodeErrorCode(value)
	return value == 0
}

func (s *script) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	value := Call(api.Proc(ScriptEngineConvertFileToPdbPath),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(ResultPath)),
	)
	// DecodeErrorCode(value)
	return value == 0
}

func (s *script) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	value := Call(api.Proc(ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(PdbFilePath)),
		uintptr(unsafe.Pointer(GuidAndAgeDetails)),
	)
	// DecodeErrorCode(value)
	return value == 0
}

func (s *script) ScriptEngineSymbolInitLoad(
	BufferToStoreDetails Headers.PVOID,
	StoredLength uint32,
	DownloadIfAvailable bool,
	SymbolPath *int8,
	IsSilentLoad bool,
) bool {
	value := Call(api.Proc(ScriptEngineSymbolInitLoad),
		uintptr(BufferToStoreDetails),
		uintptr(StoredLength),
		Bool2UintPtr(DownloadIfAvailable),
		uintptr(unsafe.Pointer(SymbolPath)),
		Bool2UintPtr(IsSilentLoad),
	)
	// DecodeErrorCode(value)
	return value == 0
}

func (s *script) ScriptEngineShowDataBasedOnSymbolTypes(
	TypeName *int8,
	Address uint64,
	IsStruct bool,
	BufferAddress Headers.PVOID,
	AdditionalParameters *int8,
) bool {
	value := Call(api.Proc(ScriptEngineShowDataBasedOnSymbolTypes),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(Address),
		Bool2UintPtr(IsStruct),
		uintptr(BufferAddress),
		uintptr(unsafe.Pointer(AdditionalParameters)),
	)
	// DecodeErrorCode(value)
	return value == 0
}
