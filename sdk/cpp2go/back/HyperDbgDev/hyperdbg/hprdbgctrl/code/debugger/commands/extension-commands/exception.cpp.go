package extension-commands
type (
Exception interface{
CommandExceptionHelp()(ok bool)//col:37
CommandException()(ok bool)//col:180
}

)
func NewException() { return & exception{} }
func (e *exception)CommandExceptionHelp()(ok bool){//col:37
return true
}

func (e *exception)CommandException()(ok bool){//col:180
return true
}

