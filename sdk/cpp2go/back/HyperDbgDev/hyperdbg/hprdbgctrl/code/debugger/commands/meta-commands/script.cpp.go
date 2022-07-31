package meta-commands
type (
Script interface{
CommandScriptHelp()(ok bool)//col:40
CommandScriptRunCommand()(ok bool)//col:105
HyperDbgScriptReadFileAndExecuteCommand()(ok bool)//col:220
CommandScript()(ok bool)//col:277
}

)
func NewScript() { return & script{} }
func (s *script)CommandScriptHelp()(ok bool){//col:40
return true
}

func (s *script)CommandScriptRunCommand()(ok bool){//col:105
return true
}

func (s *script)HyperDbgScriptReadFileAndExecuteCommand()(ok bool){//col:220
return true
}

func (s *script)CommandScript()(ok bool){//col:277
return true
}

