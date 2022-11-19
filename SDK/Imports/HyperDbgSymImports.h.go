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
	//TODO implement me
	panic("implement me")
}

// todo goto windows

func (h *hyperDbgSymImports) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymUnloadAllSymbols() uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymQuerySizeof(StructNameOrTypeName *int8, SizeOfField *uint32) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgSymImports) SymCastingQueryForFiledsAndTypes(StructName *int8, FiledOfStructName *int8, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName **int8, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) bool {
	//TODO implement me
	panic("implement me")
}

func newHyperDbgSymImports() HyperDbgSymImports {
	return &hyperDbgSymImports{}
}
