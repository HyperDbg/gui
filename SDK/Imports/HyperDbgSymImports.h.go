package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	HyperDbgSymImports interface {
		SymSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID
		SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64
		SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32
		SymUnloadAllSymbols() uint32
		SymUnloadModuleSymbol(ModuleName *byte) uint32
		SymSearchSymbolForMask(SearchMask *byte) uint32
		SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool
		SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool
		SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool
		SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool
		SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8)
		SymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool
		SymShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool
		SymQuerySizeof(StructNameOrTypeName *int8, SizeOfField *uint32) bool
		SymCastingQueryForFiledsAndTypes(StructName *int8, FiledOfStructName *int8, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName **int8, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) bool
	}
	hyperDbgSymImports struct {
	}
)

func (h *hyperDbgSymImports) SymSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	api.Proc(SymSetTextMessageCallback).Call()
}

func (h *hyperDbgSymImports) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64 {
	api.Proc(SymConvertNameToAddress).Call()
}

func (h *hyperDbgSymImports) SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32 {
	api.Proc(SymLoadFileSymbol).Call()
}

func (h *hyperDbgSymImports) SymUnloadAllSymbols() uint32 {
	api.Proc(SymUnloadAllSymbols).Call()
}

func (h *hyperDbgSymImports) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	api.Proc(SymUnloadModuleSymbol).Call()
}

func (h *hyperDbgSymImports) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	api.Proc(SymSearchSymbolForMask).Call()
}

func (h *hyperDbgSymImports) SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	api.Proc(SymGetFieldOffset).Call()
}

func (h *hyperDbgSymImports) SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	api.Proc(SymGetDataTypeSize).Call()
}

func (h *hyperDbgSymImports) SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	api.Proc(SymCreateSymbolTableForDisassembler).Call()
}

func (h *hyperDbgSymImports) SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	api.Proc(SymConvertFileToPdbPath).Call()
}

func (h *hyperDbgSymImports) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) {
	api.Proc(SymConvertFileToPdbFileAndGuidAndAgeDetails).Call()
}

func (h *hyperDbgSymImports) SymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	api.Proc(SymbolInitLoad).Call()
}

func (h *hyperDbgSymImports) SymShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	api.Proc(SymShowDataBasedOnSymbolTypes).Call()
}

func (h *hyperDbgSymImports) SymQuerySizeof(StructNameOrTypeName *int8, SizeOfField *uint32) bool {
	api.Proc(SymQuerySizeof).Call()
}

func (h *hyperDbgSymImports) SymCastingQueryForFiledsAndTypes(StructName *int8, FiledOfStructName *int8, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName **int8, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) bool {
	api.Proc(SymCastingQueryForFiledsAndTypes).Call()
}

func newHyperDbgSymImports() HyperDbgSymImports {
	return &hyperDbgSymImports{}
}
