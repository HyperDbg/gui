package script-engine-wrapper
type (
Symbol interface{
SymbolInitialReload()(ok bool)//col:41
SymbolLocalReload()(ok bool)//col:60
SymbolPrepareDebuggerWithSymbolInfo()(ok bool)//col:76
SymbolCreateDisassemblerMapCallback()(ok bool)//col:131
SymbolCreateDisassemblerSymbolMap()(ok bool)//col:152
SymbolShowFunctionNameBasedOnAddress()(ok bool)//col:254
SymbolBuildAndShowSymbolTable()(ok bool)//col:286
SymbolLoadOrDownloadSymbols()(ok bool)//col:345
SymbolConvertNameOrExprToAddress()(ok bool)//col:422
SymbolDeleteSymTable()(ok bool)//col:445
SymbolBuildSymbolTable()(ok bool)//col:908
SymbolBuildAndUpdateSymbolTable()(ok bool)//col:975
SymbolReloadSymbolTableInDebuggerMode()(ok bool)//col:1003
}

)
func NewSymbol() { return & symbol{} }
func (s *symbol)SymbolInitialReload()(ok bool){//col:41
return true
}

func (s *symbol)SymbolLocalReload()(ok bool){//col:60
return true
}

func (s *symbol)SymbolPrepareDebuggerWithSymbolInfo()(ok bool){//col:76
return true
}

func (s *symbol)SymbolCreateDisassemblerMapCallback()(ok bool){//col:131
return true
}

func (s *symbol)SymbolCreateDisassemblerSymbolMap()(ok bool){//col:152
return true
}

func (s *symbol)SymbolShowFunctionNameBasedOnAddress()(ok bool){//col:254
return true
}

func (s *symbol)SymbolBuildAndShowSymbolTable()(ok bool){//col:286
return true
}

func (s *symbol)SymbolLoadOrDownloadSymbols()(ok bool){//col:345
return true
}

func (s *symbol)SymbolConvertNameOrExprToAddress()(ok bool){//col:422
return true
}

func (s *symbol)SymbolDeleteSymTable()(ok bool){//col:445
return true
}

func (s *symbol)SymbolBuildSymbolTable()(ok bool){//col:908
return true
}

func (s *symbol)SymbolBuildAndUpdateSymbolTable()(ok bool){//col:975
return true
}

func (s *symbol)SymbolReloadSymbolTableInDebuggerMode()(ok bool){//col:1003
return true
}

