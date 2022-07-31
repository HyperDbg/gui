package debugging-commands
type (
T interface{
CommandTHelp()(ok bool)//col:42
CommandT()(ok bool)//col:166
}

)
func NewT() { return & t{} }
func (t *t)CommandTHelp()(ok bool){//col:42
return true
}

func (t *t)CommandT()(ok bool){//col:166
return true
}

