package debugging-commands


type (
Bd interface{
CommandBdHelp()(ok bool)//col:34
CommandBd()(ok bool)//col:96
}






)

func NewBd() { return & bd{} }

func (b *bd)CommandBdHelp()(ok bool){//col:34








return true
}







func (b *bd)CommandBd()(ok bool){//col:96




























return true
}









