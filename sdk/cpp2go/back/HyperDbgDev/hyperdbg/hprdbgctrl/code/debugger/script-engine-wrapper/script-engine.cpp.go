package script-engine-wrapper


type (
ScriptEngine interface{
ScriptEngineEvalSingleExpression()(ok bool)//col:125
}






)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)ScriptEngineEvalSingleExpression()(ok bool){//col:125











































return true
}









