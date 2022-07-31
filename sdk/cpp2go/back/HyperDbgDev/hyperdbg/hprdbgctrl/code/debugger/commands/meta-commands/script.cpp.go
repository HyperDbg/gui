package meta-commands


type (
Script interface{
CommandScriptHelp()(ok bool)//col:40
CommandScriptRunCommand()(ok bool)//col:106
HyperDbgScriptReadFileAndExecuteCommand()(ok bool)//col:222
CommandScript()(ok bool)//col:280
}






)

func NewScript() { return & script{} }

func (s *script)CommandScriptHelp()(ok bool){//col:40








return true
}







func (s *script)CommandScriptRunCommand()(ok bool){//col:106


























return true
}







func (s *script)HyperDbgScriptReadFileAndExecuteCommand()(ok bool){//col:222
















































return true
}







func (s *script)CommandScript()(ok bool){//col:280















return true
}









