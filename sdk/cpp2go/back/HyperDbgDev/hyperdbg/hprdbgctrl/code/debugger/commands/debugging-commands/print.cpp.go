package debugging-commands
type (
Print interface{
CommandPrintHelp()(ok bool)//col:36
CommandPrint()(ok bool)//col:131
}

)
func NewPrint() { return & print{} }
func (p *print)CommandPrintHelp()(ok bool){//col:36
return true
}

func (p *print)CommandPrint()(ok bool){//col:131
return true
}

