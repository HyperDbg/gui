package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	Symbol interface {
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
	symbol struct{}
)

func NewSymbol() Symbol { return &symbol{} }

func (s *symbol) SymSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	api.Proc(SymSetTextMessageCallback).Call(Handler)
}

func (s *symbol) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64 {
	api.Proc(SymConvertNameToAddress).Call(FunctionOrVariableName, WasFound)
}

func (s *symbol) SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32 {
	api.Proc(SymLoadFileSymbol).Call(BaseAddress, PdbFileName)
}

func (s *symbol) SymUnloadAllSymbols() uint32 {
	api.Proc(SymUnloadAllSymbols).Call()
}

func (s *symbol) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	api.Proc(SymUnloadModuleSymbol).Call(ModuleName)
}

func (s *symbol) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	api.Proc(SymSearchSymbolForMask).Call(SearchMask)
}

// todo this arg is wrong
func (s *symbol) SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	api.Proc(SymGetFieldOffset).Call(TypeName, FieldOffset)
}

func (s *symbol) SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	api.Proc(SymGetDataTypeSize).Call(TypeName, TypeSize)
}

func (s *symbol) SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	api.Proc(SymCreateSymbolTableForDisassembler).Call(CallbackFunction)
}

func (s *symbol) SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	api.Proc(SymConvertFileToPdbPath).Call(LocalFilePath, ResultPath)
}

func (s *symbol) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) {
	api.Proc(SymConvertFileToPdbFileAndGuidAndAgeDetails).Call(LocalFilePath, PdbFilePath, GuidAndAgeDetails)
}

func (s *symbol) SymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	api.Proc(SymbolInitLoad).Call(BufferToStoreDetails, StoredLength, DownloadIfAvailable, SymbolPath, IsSilentLoad)
}

func (s *symbol) SymShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	api.Proc(SymShowDataBasedOnSymbolTypes).Call(TypeName, Address, IsStruct, BufferAddress, AdditionalParameters)
}

func (s *symbol) SymQuerySizeof(StructNameOrTypeName *int8, SizeOfField *uint32) bool {
	api.Proc(SymQuerySizeof).Call(StructNameOrTypeName, SizeOfField)
}

func (s *symbol) SymCastingQueryForFiledsAndTypes(StructName *int8, FiledOfStructName *int8, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName **int8, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) bool {
	api.Proc(SymCastingQueryForFiledsAndTypes).Call(StructName, FiledOfStructName, IsFiledOfStructNamePointerOrNot, IsFiledOfStructNamePointerOrNot, NewStructOrTypeName, OffsetOfFieldFromTop, SizeOfField)
}
