package debugging-commands
type (
Output interface{
CommandOutputHelp()(ok bool)//col:42
CommandOutput()(ok bool)//col:462
}

)
func NewOutput() { return & output{} }
func (o *output)CommandOutputHelp()(ok bool){//col:42
return true
}

func (o *output)CommandOutput()(ok bool){//col:462
return true
}

