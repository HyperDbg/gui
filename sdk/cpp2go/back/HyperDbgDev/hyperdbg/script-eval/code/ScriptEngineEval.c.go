package code
type (
ScriptEngineEval interface{
GetPseudoRegValue()(ok bool)//col:69
GetValue()(ok bool)//col:132
SetValue()(ok bool)//col:161
ScriptEngineGetOperatorName()(ok bool)//col:200
ScriptEngineExecute()(ok bool)//col:1509
}

)
func NewScriptEngineEval() { return & scriptEngineEval{} }
func (s *scriptEngineEval)GetPseudoRegValue()(ok bool){//col:69
return true
}

func (s *scriptEngineEval)GetValue()(ok bool){//col:132
return true
}

func (s *scriptEngineEval)SetValue()(ok bool){//col:161
return true
}

func (s *scriptEngineEval)ScriptEngineGetOperatorName()(ok bool){//col:200
return true
}

func (s *scriptEngineEval)ScriptEngineExecute()(ok bool){//col:1509
return true
}

