package code
type (
SymbolParser interface{
SymSetTextMessageCallback()(ok bool)//col:41
ShowMessages()(ok bool)//col:75
SymGetModuleBaseFromSearchMask()(ok bool)//col:189
SymGetFieldOffsetFromModule()(ok bool)//col:320
SymGetDataTypeSizeFromModule()(ok bool)//col:367
SymLoadFileSymbol()(ok bool)//col:557
SymUnloadModuleSymbol()(ok bool)//col:617
SymUnloadAllSymbols()(ok bool)//col:662
SymConvertNameToAddress()(ok bool)//col:731
SymGetFieldOffset()(ok bool)//col:806
SymGetDataTypeSize()(ok bool)//col:870
SymSearchSymbolForMask()(ok bool)//col:916
SymCreateSymbolTableForDisassembler()(ok bool)//col:969
SymSeparateTo64BitValue()(ok bool)//col:988
SymGetFileParams()(ok bool)//col:1043
SymGetFileSize()(ok bool)//col:1102
SymShowSymbolInfo()(ok bool)//col:1250
SymDisplayMaskSymbolsCallback()(ok bool)//col:1273
SymDeliverDisassemblerSymbolMapCallback()(ok bool)//col:1299
SymShowSymbolDetails()(ok bool)//col:1344
SymTagStr()(ok bool)//col:1456
SymConvertFileToPdbPath()(ok bool)//col:1509
SymConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool)//col:1565
SymbolInitLoad()(ok bool)//col:1717
SymbolPDBDownload()(ok bool)//col:1778
SymbolAbortLoading()(ok bool)//col:1794
SymShowDataBasedOnSymbolTypes()(ok bool)//col:1918
}

)
func NewSymbolParser() { return & symbolParser{} }
func (s *symbolParser)SymSetTextMessageCallback()(ok bool){//col:41
return true
}

func (s *symbolParser)ShowMessages()(ok bool){//col:75
return true
}

func (s *symbolParser)SymGetModuleBaseFromSearchMask()(ok bool){//col:189
return true
}

func (s *symbolParser)SymGetFieldOffsetFromModule()(ok bool){//col:320
return true
}

func (s *symbolParser)SymGetDataTypeSizeFromModule()(ok bool){//col:367
return true
}

func (s *symbolParser)SymLoadFileSymbol()(ok bool){//col:557
return true
}

func (s *symbolParser)SymUnloadModuleSymbol()(ok bool){//col:617
return true
}

func (s *symbolParser)SymUnloadAllSymbols()(ok bool){//col:662
return true
}

func (s *symbolParser)SymConvertNameToAddress()(ok bool){//col:731
return true
}

func (s *symbolParser)SymGetFieldOffset()(ok bool){//col:806
return true
}

func (s *symbolParser)SymGetDataTypeSize()(ok bool){//col:870
return true
}

func (s *symbolParser)SymSearchSymbolForMask()(ok bool){//col:916
return true
}

func (s *symbolParser)SymCreateSymbolTableForDisassembler()(ok bool){//col:969
return true
}

func (s *symbolParser)SymSeparateTo64BitValue()(ok bool){//col:988
return true
}

func (s *symbolParser)SymGetFileParams()(ok bool){//col:1043
return true
}

func (s *symbolParser)SymGetFileSize()(ok bool){//col:1102
return true
}

func (s *symbolParser)SymShowSymbolInfo()(ok bool){//col:1250
return true
}

func (s *symbolParser)SymDisplayMaskSymbolsCallback()(ok bool){//col:1273
return true
}

func (s *symbolParser)SymDeliverDisassemblerSymbolMapCallback()(ok bool){//col:1299
return true
}

func (s *symbolParser)SymShowSymbolDetails()(ok bool){//col:1344
return true
}

func (s *symbolParser)SymTagStr()(ok bool){//col:1456
return true
}

func (s *symbolParser)SymConvertFileToPdbPath()(ok bool){//col:1509
return true
}

func (s *symbolParser)SymConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool){//col:1565
return true
}

func (s *symbolParser)SymbolInitLoad()(ok bool){//col:1717
return true
}

func (s *symbolParser)SymbolPDBDownload()(ok bool){//col:1778
return true
}

func (s *symbolParser)SymbolAbortLoading()(ok bool){//col:1794
return true
}

func (s *symbolParser)SymShowDataBasedOnSymbolTypes()(ok bool){//col:1918
return true
}

