package header


const(
    SCRIPT_ENGINE_ERROR_FREE = 1  //col:55
    SCRIPT_ENGINE_ERROR_SYNTAX = 2  //col:56
    SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN = 3  //col:57
    SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE = 4  //col:58
    SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE = 5  //col:59
    SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL = 6  //col:60
)



type (
ScriptEngine interface{
__declspec()(ok bool)//col:61
}

)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)__declspec()(ok bool){//col:61









































return true
}




