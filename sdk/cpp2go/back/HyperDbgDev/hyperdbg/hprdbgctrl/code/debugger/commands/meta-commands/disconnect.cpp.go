package meta-commands


type (
Disconnect interface{
CommandDisconnectHelp()(ok bool)//col:34
CommandDisconnect()(ok bool)//col:98
}
















)

func NewDisconnect() { return & disconnect{} }

func (d *disconnect)CommandDisconnectHelp()(ok bool){//col:34






return true
}

















func (d *disconnect)CommandDisconnect()(ok bool){//col:98

































return true
}



















