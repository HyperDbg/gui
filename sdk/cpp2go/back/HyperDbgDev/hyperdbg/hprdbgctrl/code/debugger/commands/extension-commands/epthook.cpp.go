package extension-commands
type (
Epthook interface{
CommandEptHookHelp()(ok bool)//col:36
CommandEptHook()(ok bool)//col:182
}

)
func NewEpthook() { return & epthook{} }
func (e *epthook)CommandEptHookHelp()(ok bool){//col:36
return true
}

func (e *epthook)CommandEptHook()(ok bool){//col:182
return true
}

