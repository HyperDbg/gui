package extension-commands


type (
Unhide interface{
CommandUnhideHelp()(ok bool)//col:28
CommandUnhide()(ok bool)//col:97
}
















)

func NewUnhide() { return & unhide{} }

func (u *unhide)CommandUnhideHelp()(ok bool){//col:28







return true
}

















func (u *unhide)CommandUnhide()(ok bool){//col:97


































return true
}



















