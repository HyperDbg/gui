package Imports

func (s *symbol) SymSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	valu := Call(api.Proc(SymSetTextMessageCallback), uintptr(Handler))
	//DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (s *symbol) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound *bool) uint64 {
	valu := Call(api.Proc(SymConvertNameToAddress),
		uintptr(unsafe.Pointer(FunctionOrVariableName)),
		uintptr(unsafe.Pointer(WasFound)),
	)
	//DecodeErrorCode(valu)
	return uint64(valu)
}

func (s *symbol) SymLoadFileSymbol(BaseAddress uint64, PdbFileName *byte) uint32 {
	valu := Call(api.Proc(SymLoadFileSymbol),
		uintptr(BaseAddress),
		uintptr(unsafe.Pointer(PdbFileName)),
	)
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *symbol) SymUnloadAllSymbols() uint32 {
	valu := Call(api.Proc(SymUnloadAllSymbols))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *symbol) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	valu := Call(api.Proc(SymUnloadModuleSymbol), uintptr(unsafe.Pointer(ModuleName)))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

func (s *symbol) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	valu := Call(api.Proc(SymSearchSymbolForMask), uintptr(unsafe.Pointer(SearchMask)))
	//DecodeErrorCode(valu)
	return uint32(valu)
}

// todo this arg is wrong
func (s *symbol) SymGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	valu := Call(api.Proc(SymGetFieldOffset),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(FieldName)),
		uintptr(unsafe.Pointer(FieldOffset)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *symbol) SymGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	valu := Call(api.Proc(SymGetDataTypeSize),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(unsafe.Pointer(TypeSize)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *symbol) SymCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	valu := Call(api.Proc(SymCreateSymbolTableForDisassembler), uintptr(CallbackFunction))
	//DecodeErrorCode(valu) //todo need decode r2 as error code ?
	return valu == 0
}

func (s *symbol) SymConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	valu := Call(api.Proc(SymConvertFileToPdbPath),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(ResultPath)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *symbol) SymConvertFileToPdbFileAndGuidAndAgeDetails(
	LocalFilePath,
	PdbFilePath,
	GuidAndAgeDetails *int8) {
	valu := Call(api.Proc(SymConvertFileToPdbFileAndGuidAndAgeDetails),
		uintptr(unsafe.Pointer(LocalFilePath)),
		uintptr(unsafe.Pointer(PdbFilePath)),
		uintptr(unsafe.Pointer(GuidAndAgeDetails)),
	)
	DecodeErrorCode(valu)
	return //todonot ret code ?
}

func (s *symbol) SymbolInitLoad(
	BufferToStoreDetails Headers.PVOID,
	StoredLength uint32,
	DownloadIfAvailable bool,
	SymbolPath *int8,
	IsSilentLoad bool,
) bool {
	valu := Call(api.Proc(SymbolInitLoad),
		uintptr(BufferToStoreDetails),
		uintptr(StoredLength),
		Bool2UintPtr(DownloadIfAvailable),
		uintptr(unsafe.Pointer(SymbolPath)),
		Bool2UintPtr(IsSilentLoad),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *symbol) SymShowDataBasedOnSymbolTypes(
	TypeName *int8,
	Address uint64,
	IsStruct bool,
	BufferAddress Headers.PVOID,
	AdditionalParameters *int8,
) bool {
	valu := Call(api.Proc(SymShowDataBasedOnSymbolTypes),
		uintptr(unsafe.Pointer(TypeName)),
		uintptr(Address),
		Bool2UintPtr(IsStruct),
		uintptr(BufferAddress),
		uintptr(unsafe.Pointer(AdditionalParameters)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}

func (s *symbol) SymQuerySizeof(
	StructNameOrTypeName *int8,
	SizeOfField *uint32) bool {
	valu := Call(api.Proc(SymQuerySizeof),
		uintptr(unsafe.Pointer(StructNameOrTypeName)),
		uintptr(unsafe.Pointer(SizeOfField)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
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
	valu := Call(api.Proc(SymCastingQueryForFiledsAndTypes),
		uintptr(unsafe.Pointer(StructName)),
		uintptr(unsafe.Pointer(FiledOfStructName)),
		Bool2UintPtr(IsStructNamePointerOrNot),
		Bool2UintPtr(IsFiledOfStructNamePointerOrNot),
		uintptr(unsafe.Pointer(NewStructOrTypeName)), //todo check this.pointHyperDbgSymImports.h
		uintptr(unsafe.Pointer(OffsetOfFieldFromTop)),
		uintptr(unsafe.Pointer(SizeOfField)),
	)
	//DecodeErrorCode(valu)
	return valu == 0
}
