package code


type (
SymbolParser interface{
SymSetTextMessageCallback()(ok bool)//col:41
ShowMessages()(ok bool)//col:76
SymGetModuleBaseFromSearchMask()(ok bool)//col:191
SymGetFieldOffsetFromModule()(ok bool)//col:323
SymGetDataTypeSizeFromModule()(ok bool)//col:371
SymLoadFileSymbol()(ok bool)//col:562
SymUnloadModuleSymbol()(ok bool)//col:623
SymUnloadAllSymbols()(ok bool)//col:669
SymConvertNameToAddress()(ok bool)//col:739
SymGetFieldOffset()(ok bool)//col:815
SymGetDataTypeSize()(ok bool)//col:880
SymSearchSymbolForMask()(ok bool)//col:927
SymCreateSymbolTableForDisassembler()(ok bool)//col:981
SymSeparateTo64BitValue()(ok bool)//col:1001
SymGetFileParams()(ok bool)//col:1057
SymGetFileSize()(ok bool)//col:1117
SymShowSymbolInfo()(ok bool)//col:1266
SymDisplayMaskSymbolsCallback()(ok bool)//col:1290
SymDeliverDisassemblerSymbolMapCallback()(ok bool)//col:1317
SymShowSymbolDetails()(ok bool)//col:1363
SymTagStr()(ok bool)//col:1476
SymConvertFileToPdbPath()(ok bool)//col:1530
SymConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool)//col:1587
SymbolInitLoad()(ok bool)//col:1740
SymbolPDBDownload()(ok bool)//col:1802
SymbolAbortLoading()(ok bool)//col:1819
SymShowDataBasedOnSymbolTypes()(ok bool)//col:1944
}






)

func NewSymbolParser() { return & symbolParser{} }

func (s *symbolParser)SymSetTextMessageCallback()(ok bool){//col:41





return true
}







func (s *symbolParser)ShowMessages()(ok bool){//col:76






















return true
}







func (s *symbolParser)SymGetModuleBaseFromSearchMask()(ok bool){//col:191





















































return true
}







func (s *symbolParser)SymGetFieldOffsetFromModule()(ok bool){//col:323




















































return true
}







func (s *symbolParser)SymGetDataTypeSizeFromModule()(ok bool){//col:371


















return true
}







func (s *symbolParser)SymLoadFileSymbol()(ok bool){//col:562








































































return true
}







func (s *symbolParser)SymUnloadModuleSymbol()(ok bool){//col:623































return true
}







func (s *symbolParser)SymUnloadAllSymbols()(ok bool){//col:669






















return true
}







func (s *symbolParser)SymConvertNameToAddress()(ok bool){//col:739































return true
}







func (s *symbolParser)SymGetFieldOffset()(ok bool){//col:815

































return true
}







func (s *symbolParser)SymGetDataTypeSize()(ok bool){//col:880




























return true
}







func (s *symbolParser)SymSearchSymbolForMask()(ok bool){//col:927


















return true
}







func (s *symbolParser)SymCreateSymbolTableForDisassembler()(ok bool){//col:981

















return true
}







func (s *symbolParser)SymSeparateTo64BitValue()(ok bool){//col:1001









return true
}







func (s *symbolParser)SymGetFileParams()(ok bool){//col:1057






















return true
}







func (s *symbolParser)SymGetFileSize()(ok bool){//col:1117























return true
}







func (s *symbolParser)SymShowSymbolInfo()(ok bool){//col:1266








































































return true
}







func (s *symbolParser)SymDisplayMaskSymbolsCallback()(ok bool){//col:1290








return true
}







func (s *symbolParser)SymDeliverDisassemblerSymbolMapCallback()(ok bool){//col:1317








return true
}







func (s *symbolParser)SymShowSymbolDetails()(ok bool){//col:1363















return true
}







func (s *symbolParser)SymTagStr()(ok bool){//col:1476







































































return true
}







func (s *symbolParser)SymConvertFileToPdbPath()(ok bool){//col:1530

































return true
}







func (s *symbolParser)SymConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool){//col:1587

































return true
}







func (s *symbolParser)SymbolInitLoad()(ok bool){//col:1740

































































































return true
}







func (s *symbolParser)SymbolPDBDownload()(ok bool){//col:1802











































return true
}







func (s *symbolParser)SymbolAbortLoading()(ok bool){//col:1819








return true
}







func (s *symbolParser)SymShowDataBasedOnSymbolTypes()(ok bool){//col:1944





















































return true
}









