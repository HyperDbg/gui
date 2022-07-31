package debugging-commands
type (
Lm interface{
CommandLmHelp()(ok bool)//col:41
CommandLmShowUserModeModule()(ok bool)//col:220
CommandLmShowKernelModeModule()(ok bool)//col:330
CommandLm()(ok bool)//col:484
}

)
func NewLm() { return & lm{} }
func (l *lm)CommandLmHelp()(ok bool){//col:41
return true
}

func (l *lm)CommandLmShowUserModeModule()(ok bool){//col:220
return true
}

func (l *lm)CommandLmShowKernelModeModule()(ok bool){//col:330
return true
}

func (l *lm)CommandLm()(ok bool){//col:484
return true
}

