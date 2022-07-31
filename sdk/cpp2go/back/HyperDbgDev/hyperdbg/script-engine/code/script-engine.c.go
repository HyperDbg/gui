package code
const(
)
type (
ScriptEngine interface{
ScriptEngineConvertNameToAddress()(ok bool)//col:33
ScriptEngineLoadFileSymbol()(ok bool)//col:51
ScriptEngineSetTextMessageCallback()(ok bool)//col:63
ScriptEngineUnloadAllSymbols()(ok bool)//col:77
ScriptEngineUnloadModuleSymbol()(ok bool)//col:92
ScriptEngineSearchSymbolForMask()(ok bool)//col:107
ScriptEngineGetFieldOffset()(ok bool)//col:124
ScriptEngineGetDataTypeSize()(ok bool)//col:140
ScriptEngineCreateSymbolTableForDisassembler()(ok bool)//col:155
ScriptEngineConvertFileToPdbPath()(ok bool)//col:171
ScriptEngineSymbolInitLoad()(ok bool)//col:194
ScriptEngineShowDataBasedOnSymbolTypes()(ok bool)//col:217
ScriptEngineSymbolAbortLoading()(ok bool)//col:231
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool)//col:248
ScriptEngineParse()(ok bool)//col:476
CodeGen()(ok bool)//col:1504
BooleanExpressionExtractEnd()(ok bool)//col:1552
ScriptEngineBooleanExpresssionParse()(ok bool)//col:1734
NewSymbol()(ok bool)//col:1749
NewStringSymbol()(ok bool)//col:1766
GetStringSymbolSize()(ok bool)//col:1779
RemoveSymbol()(ok bool)//col:1792
PrintSymbol()(ok bool)//col:1810
ToSymbol()(ok bool)//col:1879
NewSymbolBuffer()(ok bool)//col:1896
RemoveSymbolBuffer()(ok bool)//col:1909
PushSymbol()(ok bool)//col:2015
PrintSymbolBuffer()(ok bool)//col:2042
RegisterToInt()(ok bool)//col:2061
PseudoRegToInt()(ok bool)//col:2080
SemanticRuleToInt()(ok bool)//col:2099
HandleError()(ok bool)//col:2192
GetGlobalIdentifierVal()(ok bool)//col:2213
GetLocalIdentifierVal()(ok bool)//col:2234
NewGlobalIdentifier()(ok bool)//col:2248
NewLocalIdentifier()(ok bool)//col:2262
LalrGetRhsSize()(ok bool)//col:2283
LalrIsOperandType()(ok bool)//col:2340
}

)
func NewScriptEngine() { return & scriptEngine{} }
func (s *scriptEngine)ScriptEngineConvertNameToAddress()(ok bool){//col:33
return true
}

func (s *scriptEngine)ScriptEngineLoadFileSymbol()(ok bool){//col:51
return true
}

func (s *scriptEngine)ScriptEngineSetTextMessageCallback()(ok bool){//col:63
return true
}

func (s *scriptEngine)ScriptEngineUnloadAllSymbols()(ok bool){//col:77
return true
}

func (s *scriptEngine)ScriptEngineUnloadModuleSymbol()(ok bool){//col:92
return true
}

func (s *scriptEngine)ScriptEngineSearchSymbolForMask()(ok bool){//col:107
return true
}

func (s *scriptEngine)ScriptEngineGetFieldOffset()(ok bool){//col:124
return true
}

func (s *scriptEngine)ScriptEngineGetDataTypeSize()(ok bool){//col:140
return true
}

func (s *scriptEngine)ScriptEngineCreateSymbolTableForDisassembler()(ok bool){//col:155
return true
}

func (s *scriptEngine)ScriptEngineConvertFileToPdbPath()(ok bool){//col:171
return true
}

func (s *scriptEngine)ScriptEngineSymbolInitLoad()(ok bool){//col:194
return true
}

func (s *scriptEngine)ScriptEngineShowDataBasedOnSymbolTypes()(ok bool){//col:217
return true
}

func (s *scriptEngine)ScriptEngineSymbolAbortLoading()(ok bool){//col:231
return true
}

func (s *scriptEngine)ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails()(ok bool){//col:248
return true
}

func (s *scriptEngine)ScriptEngineParse()(ok bool){//col:476
return true
}

func (s *scriptEngine)CodeGen()(ok bool){//col:1504
return true
}

func (s *scriptEngine)BooleanExpressionExtractEnd()(ok bool){//col:1552
return true
}

func (s *scriptEngine)ScriptEngineBooleanExpresssionParse()(ok bool){//col:1734
return true
}

func (s *scriptEngine)NewSymbol()(ok bool){//col:1749
return true
}

func (s *scriptEngine)NewStringSymbol()(ok bool){//col:1766
return true
}

func (s *scriptEngine)GetStringSymbolSize()(ok bool){//col:1779
return true
}

func (s *scriptEngine)RemoveSymbol()(ok bool){//col:1792
return true
}

func (s *scriptEngine)PrintSymbol()(ok bool){//col:1810
return true
}

func (s *scriptEngine)ToSymbol()(ok bool){//col:1879
return true
}

func (s *scriptEngine)NewSymbolBuffer()(ok bool){//col:1896
return true
}

func (s *scriptEngine)RemoveSymbolBuffer()(ok bool){//col:1909
return true
}

func (s *scriptEngine)PushSymbol()(ok bool){//col:2015
return true
}

func (s *scriptEngine)PrintSymbolBuffer()(ok bool){//col:2042
return true
}

func (s *scriptEngine)RegisterToInt()(ok bool){//col:2061
return true
}

func (s *scriptEngine)PseudoRegToInt()(ok bool){//col:2080
return true
}

func (s *scriptEngine)SemanticRuleToInt()(ok bool){//col:2099
return true
}

func (s *scriptEngine)HandleError()(ok bool){//col:2192
return true
}

func (s *scriptEngine)GetGlobalIdentifierVal()(ok bool){//col:2213
return true
}

func (s *scriptEngine)GetLocalIdentifierVal()(ok bool){//col:2234
return true
}

func (s *scriptEngine)NewGlobalIdentifier()(ok bool){//col:2248
return true
}

func (s *scriptEngine)NewLocalIdentifier()(ok bool){//col:2262
return true
}

func (s *scriptEngine)LalrGetRhsSize()(ok bool){//col:2283
return true
}

func (s *scriptEngine)LalrIsOperandType()(ok bool){//col:2340
return true
}

