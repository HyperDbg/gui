package extension-commands
type (
Tsc interface{
CommandTscHelp()(ok bool)//col:33
CommandTsc()(ok bool)//col:122
}

)
func NewTsc() { return & tsc{} }
func (t *tsc)CommandTscHelp()(ok bool){//col:33
return true
}

func (t *tsc)CommandTsc()(ok bool){//col:122
return true
}

