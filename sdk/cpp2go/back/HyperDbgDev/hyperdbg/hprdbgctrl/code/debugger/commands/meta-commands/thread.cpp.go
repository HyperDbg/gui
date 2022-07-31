package meta-commands
type (
Thread interface{
CommandThreadHelp()(ok bool)//col:47
CommandThreadListThreads()(ok bool)//col:109
CommandThread()(ok bool)//col:304
}

)
func NewThread() { return & thread{} }
func (t *thread)CommandThreadHelp()(ok bool){//col:47
return true
}

func (t *thread)CommandThreadListThreads()(ok bool){//col:109
return true
}

func (t *thread)CommandThread()(ok bool){//col:304
return true
}

