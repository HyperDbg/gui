package debugging-commands
type (
R interface{
CommandRHelp()(ok bool)//col:164
ShowAllRegisters()(ok bool)//col:178
CommandR()(ok bool)//col:365
}

)
func NewR() { return & r{} }
func (r *r)CommandRHelp()(ok bool){//col:164
return true
}

func (r *r)ShowAllRegisters()(ok bool){//col:178
return true
}

func (r *r)CommandR()(ok bool){//col:365
return true
}

