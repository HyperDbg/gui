package extension-commands


type (
Pmc interface{
CommandPmcHelp()(ok bool)//col:33
CommandPmc()(ok bool)//col:123
}
















)

func NewPmc() { return & pmc{} }

func (p *pmc)CommandPmcHelp()(ok bool){//col:33












return true
}

















func (p *pmc)CommandPmc()(ok bool){//col:123




















































return true
}



















