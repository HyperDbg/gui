package meta-commands
type (
Attach interface{
CommandAttachHelp()(ok bool)//col:34
CommandAttach()(ok bool)//col:126
}

)
func NewAttach() { return & attach{} }
func (a *attach)CommandAttachHelp()(ok bool){//col:34
return true
}

func (a *attach)CommandAttach()(ok bool){//col:126
return true
}

