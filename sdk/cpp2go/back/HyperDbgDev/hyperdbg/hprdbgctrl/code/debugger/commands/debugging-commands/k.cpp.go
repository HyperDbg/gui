package debugging-commands
type (
K interface{
CommandKHelp()(ok bool)//col:45
CommandK()(ok bool)//col:184
}

)
func NewK() { return & k{} }
func (k *k)CommandKHelp()(ok bool){//col:45
return true
}

func (k *k)CommandK()(ok bool){//col:184
return true
}

