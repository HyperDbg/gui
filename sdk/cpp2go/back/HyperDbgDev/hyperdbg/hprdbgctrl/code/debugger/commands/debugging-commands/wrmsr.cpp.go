package debugging-commands
type (
Wrmsr interface{
CommandWrmsrHelp()(ok bool)//col:31
CommandWrmsr()(ok bool)//col:167
}

)
func NewWrmsr() { return & wrmsr{} }
func (w *wrmsr)CommandWrmsrHelp()(ok bool){//col:31
return true
}

func (w *wrmsr)CommandWrmsr()(ok bool){//col:167
return true
}

