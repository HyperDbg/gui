package extension-commands


type (
Exception interface{
CommandExceptionHelp()(ok bool)//col:37
CommandException()(ok bool)//col:181
}
















)

func NewException() { return & exception{} }

func (e *exception)CommandExceptionHelp()(ok bool){//col:37















return true
}

















func (e *exception)CommandException()(ok bool){//col:181




















































































return true
}



















