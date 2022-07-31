package script-engine-wrapper


type (
Symbol interface{
SymbolInitialReload()(ok bool)//col:41
SymbolLocalReload()(ok bool)//col:61
SymbolPrepareDebuggerWithSymbolInfo()(ok bool)//col:78
SymbolCreateDisassemblerMapCallback()(ok bool)//col:134
SymbolCreateDisassemblerSymbolMap()(ok bool)//col:156
SymbolShowFunctionNameBasedOnAddress()(ok bool)//col:259
SymbolBuildAndShowSymbolTable()(ok bool)//col:292
SymbolLoadOrDownloadSymbols()(ok bool)//col:352
SymbolConvertNameOrExprToAddress()(ok bool)//col:430
SymbolDeleteSymTable()(ok bool)//col:454
SymbolBuildSymbolTable()(ok bool)//col:918
SymbolBuildAndUpdateSymbolTable()(ok bool)//col:986
SymbolReloadSymbolTableInDebuggerMode()(ok bool)//col:1015
}

)

func NewSymbol() { return & symbol{} }

func (s *symbol)SymbolInitialReload()(ok bool){//col:41





return true
}


func (s *symbol)SymbolLocalReload()(ok bool){//col:61






return true
}


func (s *symbol)SymbolPrepareDebuggerWithSymbolInfo()(ok bool){//col:78




return true
}


func (s *symbol)SymbolCreateDisassemblerMapCallback()(ok bool){//col:134























return true
}


func (s *symbol)SymbolCreateDisassemblerSymbolMap()(ok bool){//col:156






return true
}


func (s *symbol)SymbolShowFunctionNameBasedOnAddress()(ok bool){//col:259

























































return true
}


func (s *symbol)SymbolBuildAndShowSymbolTable()(ok bool){//col:292



















return true
}


func (s *symbol)SymbolLoadOrDownloadSymbols()(ok bool){//col:352
























return true
}


func (s *symbol)SymbolConvertNameOrExprToAddress()(ok bool){//col:430




































return true
}


func (s *symbol)SymbolDeleteSymTable()(ok bool){//col:454















return true
}


func (s *symbol)SymbolBuildSymbolTable()(ok bool){//col:918






























































































































































































































return true
}


func (s *symbol)SymbolBuildAndUpdateSymbolTable()(ok bool){//col:986
























return true
}


func (s *symbol)SymbolReloadSymbolTableInDebuggerMode()(ok bool){//col:1015













return true
}




