package debugging-commands


type (
Unload interface{
CommandUnloadHelp()(ok bool)//col:38
CommandUnload()(ok bool)//col:125
}













































)

func NewUnload() { return & unload{} }

func (u *unload)CommandUnloadHelp()(ok bool){//col:38









return true
}














































func (u *unload)CommandUnload()(ok bool){//col:125





















































return true
}
















































