package meta-commands
type (
Process interface{
CommandProcessHelp()(ok bool)//col:46
CommandProcess()(ok bool)//col:240
}

)
func NewProcess() { return & process{} }
func (p *process)CommandProcessHelp()(ok bool){//col:46
return true
}

func (p *process)CommandProcess()(ok bool){//col:240
return true
}

