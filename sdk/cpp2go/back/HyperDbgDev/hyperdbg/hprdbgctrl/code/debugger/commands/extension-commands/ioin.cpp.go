package extension-commands
type (
Ioin interface{
CommandIoinHelp()(ok bool)//col:34
CommandIoin()(ok bool)//col:160
}

)
func NewIoin() { return & ioin{} }
func (i *ioin)CommandIoinHelp()(ok bool){//col:34
return true
}

func (i *ioin)CommandIoin()(ok bool){//col:160
return true
}

