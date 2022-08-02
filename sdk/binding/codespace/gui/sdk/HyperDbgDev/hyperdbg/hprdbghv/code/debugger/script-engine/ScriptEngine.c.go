package script_engine



type (
	ScriptEngine interface {
		ScriptEngineWrapperGetInstructionPointer() (ok bool) //col:12
	}
	scriptEngine struct{}
)

func NewScriptEngine() ScriptEngine { return &scriptEngine{} }

func (s *scriptEngine) ScriptEngineWrapperGetInstructionPointer() (ok bool) { //col:12

















	return true
}


