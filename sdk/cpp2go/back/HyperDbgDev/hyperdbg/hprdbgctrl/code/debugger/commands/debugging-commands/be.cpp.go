package debugging-commands


type (
Be interface{
CommandBeHelp()(ok bool)//col:34
CommandBe()(ok bool)//col:96
}






)

func NewBe() { return & be{} }

func (b *be)CommandBeHelp()(ok bool){//col:34








return true
}







func (b *be)CommandBe()(ok bool){//col:96




























return true
}









