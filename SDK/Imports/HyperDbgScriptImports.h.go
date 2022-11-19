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
	v := 1
	if WasFound {
		v = 0
	}
	valu := Call(api.Proc(ScriptEngineConvertNameToAddress), uintptr(unsafe.Pointer(FunctionOrVariableName)), uintptr(v))
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
	valu := Call(api.Proc(ScriptEngineGetDataTypeSize), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	valu := Call(api.Proc(ScriptEngineCreateSymbolTableForDisassembler), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	valu := Call(api.Proc(ScriptEngineConvertFileToPdbPath), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	valu := Call(api.Proc(ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	valu := Call(api.Proc(ScriptEngineSymbolInitLoad), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *script) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	valu := Call(api.Proc(ScriptEngineShowDataBasedOnSymbolTypes), uintptr(Handler))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}
