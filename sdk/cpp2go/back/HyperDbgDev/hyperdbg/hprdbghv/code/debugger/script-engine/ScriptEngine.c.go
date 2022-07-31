package script-engine


type (
ScriptEngine interface{
ScriptEngineWrapperGetInstructionPointer()(ok bool)//col:42
ScriptEngineWrapperGetAddressOfReservedBuffer()(ok bool)//col:55
}
















)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)ScriptEngineWrapperGetInstructionPointer()(ok bool){//col:42















return true
}

















func (s *scriptEngine)ScriptEngineWrapperGetAddressOfReservedBuffer()(ok bool){//col:55




return true
}



















