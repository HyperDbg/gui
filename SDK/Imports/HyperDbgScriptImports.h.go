package Imports

import (
	"github.com/ddkwork/hyperdbgui/SDK/Headers"
	"unsafe"
)

type (
	pdbParser interface {
		ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID
		ScriptEngineSymbolAbortLoading() Headers.VOID
		ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64
		ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32
		ScriptEngineUnloadAllSymbols() uint32
		ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32
		ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32
		ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool
		ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool
		ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool
		ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool
		ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool
		ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool
		ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool
	}
	Script interface {
		//ScriptEngineParse(str *int8) PSYMBOL_BUFFER
		//PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID
		//PrintSymbol(PSYMBOL Symbol) Headers.VOID
		//RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer) Headers.VOID
		pdbParser
	}
	script struct{}
)

func NewScript() Script { return &script{} }

func (s *script) ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	valu := Call(api.Proc(ScriptEngineSetTextMessageCallback), uintptr(Handler))
	//DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineSymbolAbortLoading() Headers.VOID {
	valu := Call(api.Proc(ScriptEngineSymbolAbortLoading))
	//DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64 {
	//WasFound need input a point value ?
	valu := Call(api.Proc(ScriptEngineConvertNameToAddress), uintptr(unsafe.Pointer(FunctionOrVariableName)), Bool2UintPtr(WasFound))
	//DecodeErrorCode(valu)
	return uint64(valu)
}

func (s *script) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32 {
	valu := Call(api.Proc(ScriptEngineLoadFileSymbol), uintptr(BaseAddress), uintptr(unsafe.Pointer(PdbFileName)))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *script) ScriptEngineUnloadAllSymbols() uint32 {
	valu := Call(api.Proc(ScriptEngineUnloadAllSymbols))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *script) ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32 {
	valu := Call(api.Proc(ScriptEngineUnloadModuleSymbol), uintptr(unsafe.Pointer(ModuleName)))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *script) ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32 {
	valu := Call(api.Proc(ScriptEngineSearchSymbolForMask), uintptr(unsafe.Pointer(SearchMask)))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *script) ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	valu := Call(api.Proc(ScriptEngineGetFieldOffset),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(FieldName)),
		uintptr(unsafe.Pointer(FieldOffset)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	valu := Call(api.Proc(ScriptEngineGetDataTypeSize),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(TypeSize)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	valu := Call(api.Proc(ScriptEngineCreateSymbolTableForDisassembler), uintptr(CallbackFunction))
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	valu := Call(api.Proc(ScriptEngineConvertFileToPdbPath),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(ResultPath)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	valu := Call(api.Proc(ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(PdbFilePath)),
		uintptr(unsafe.Pointer(GuidAndAgeDetails)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineSymbolInitLoad(
	BufferToStoreDetails Headers.PVOID,
	StoredLength uint32,
	DownloadIfAvailable bool,
	SymbolPath *int8,
	IsSilentLoad bool,
) bool {
	valu := Call(api.Proc(ScriptEngineSymbolInitLoad),
		uintptr(BufferToStoreDetails),
		uintptr(StoredLength),
		Bool2UintPtr(DownloadIfAvailable),
		uintptr(unsafe.Pointer(SymbolPath)),
		Bool2UintPtr(IsSilentLoad),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *script) ScriptEngineShowDataBasedOnSymbolTypes(
	TypeName *int8,
	Address uint64,
	IsStruct bool,
	BufferAddress Headers.PVOID,
	AdditionalParameters *int8,
) bool {
	valu := Call(api.Proc(ScriptEngineShowDataBasedOnSymbolTypes),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(Address),
		Bool2UintPtr(IsStruct),
		uintptr(BufferAddress),
		uintptr(unsafe.Pointer(AdditionalParameters)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}
