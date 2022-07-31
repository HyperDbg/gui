package extension-commands
type (
Msrwrite interface{
CommandMsrwriteHelp()(ok bool)//col:33
CommandMsrwrite()(ok bool)//col:161
}

)
func NewMsrwrite() { return & msrwrite{} }
func (m *msrwrite)CommandMsrwriteHelp()(ok bool){//col:33
return true
}

func (m *msrwrite)CommandMsrwrite()(ok bool){//col:161
return true
}

