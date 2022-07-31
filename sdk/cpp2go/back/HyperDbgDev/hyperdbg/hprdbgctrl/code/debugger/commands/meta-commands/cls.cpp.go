package meta-commands
type (
Cls interface{
CommandClearScreenHelp()(ok bool)//col:25
CommandClearScreen()(ok bool)//col:38
}

)
func NewCls() { return & cls{} }
func (c *cls)CommandClearScreenHelp()(ok bool){//col:25
return true
}

func (c *cls)CommandClearScreen()(ok bool){//col:38
return true
}

