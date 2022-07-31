package debugging-commands
type (
P interface{
CommandPHelp()(ok bool)//col:42
CommandP()(ok bool)//col:166
}

)
func NewP() { return & p{} }
func (p *p)CommandPHelp()(ok bool){//col:42
return true
}

func (p *p)CommandP()(ok bool){//col:166
return true
}

