package extension-commands


type (
Msrread interface{
CommandMsrreadHelp()(ok bool)//col:33
CommandMsrread()(ok bool)//col:162
}

)

func NewMsrread() { return & msrread{} }

func (m *msrread)CommandMsrreadHelp()(ok bool){//col:33












return true
}


func (m *msrread)CommandMsrread()(ok bool){//col:162












































































return true
}




