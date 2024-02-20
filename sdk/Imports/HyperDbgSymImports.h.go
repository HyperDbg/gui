package Imports

import (
	"github.com/ddkwork/hyperdbgui/sdk/Headers"
)

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
