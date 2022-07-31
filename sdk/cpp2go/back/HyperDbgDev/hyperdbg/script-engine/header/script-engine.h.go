package header
type     SCRIPT_ENGINE_ERROR_FREE uint32
const(
    SCRIPT_ENGINE_ERROR_FREE SCRIPT_ENGINE_ERROR_TYPE = 1  //col:55
    SCRIPT_ENGINE_ERROR_SYNTAX SCRIPT_ENGINE_ERROR_TYPE = 2  //col:56
    SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN SCRIPT_ENGINE_ERROR_TYPE = 3  //col:57
    SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE SCRIPT_ENGINE_ERROR_TYPE = 4  //col:58
    SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE SCRIPT_ENGINE_ERROR_TYPE = 5  //col:59
    SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL SCRIPT_ENGINE_ERROR_TYPE = 6  //col:60
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

