package meta-commands


type (
Pe interface{
CommandPeHelp()(ok bool)//col:33
CommandPe()(ok bool)//col:147
}






)

func NewPe() { return & pe{} }

func (p *pe)CommandPeHelp()(ok bool){//col:33










return true
}







func (p *pe)CommandPe()(ok bool){//col:147


























































return true
}









