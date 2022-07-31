package debugging-commands
type (
Bl interface{
CommandBlHelp()(ok bool)//col:30
CommandBl()(ok bool)//col:75
}

)
func NewBl() { return & bl{} }
func (b *bl)CommandBlHelp()(ok bool){//col:30
return true
}

func (b *bl)CommandBl()(ok bool){//col:75
return true
}

