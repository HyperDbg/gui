package debugging-commands
type (
Load interface{
CommandLoadHelp()(ok bool)//col:36
CommandLoadVmmModule()(ok bool)//col:113
CommandLoad()(ok bool)//col:182
}

)
func NewLoad() { return & load{} }
func (l *load)CommandLoadHelp()(ok bool){//col:36
return true
}

func (l *load)CommandLoadVmmModule()(ok bool){//col:113
return true
}

func (l *load)CommandLoad()(ok bool){//col:182
return true
}

