package API


type (
Terminal interface{
#if   defined()(ok bool)//col:96
ZyanStatus ZyanTerminalIsTTY()(ok bool)//col:155
}






)

func NewTerminal() { return & terminal{} }

func (t *terminal)#if   defined()(ok bool){//col:96






















































return true
}







func (t *terminal)ZyanStatus ZyanTerminalIsTTY()(ok bool){//col:155




















































return true
}









