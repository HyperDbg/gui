package meta-commands


type (
Thread interface{
CommandThreadHelp()(ok bool)//col:47
CommandThreadListThreads()(ok bool)//col:110
CommandThread()(ok bool)//col:306
}
















)

func NewThread() { return & thread{} }

func (t *thread)CommandThreadHelp()(ok bool){//col:47





















return true
}

















func (t *thread)CommandThreadListThreads()(ok bool){//col:110





































return true
}

















func (t *thread)CommandThread()(ok bool){//col:306



























































































































































return true
}



















