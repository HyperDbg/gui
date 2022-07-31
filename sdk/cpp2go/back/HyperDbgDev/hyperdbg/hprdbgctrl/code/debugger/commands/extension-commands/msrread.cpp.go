package extension-commands
type (
Msrread interface{
CommandMsrreadHelp()(ok bool)//col:33
CommandMsrread()(ok bool)//col:161
}

)
func NewMsrread() { return & msrread{} }
func (m *msrread)CommandMsrreadHelp()(ok bool){//col:33
return true
}

func (m *msrread)CommandMsrread()(ok bool){//col:161
return true
}

