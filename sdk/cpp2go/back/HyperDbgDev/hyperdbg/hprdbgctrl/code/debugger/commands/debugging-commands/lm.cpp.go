package debugging-commands


type (
Lm interface{
CommandLmHelp()(ok bool)//col:41
CommandLmShowUserModeModule()(ok bool)//col:221
CommandLmShowKernelModeModule()(ok bool)//col:332
CommandLm()(ok bool)//col:487
}






)

func NewLm() { return & lm{} }

func (l *lm)CommandLmHelp()(ok bool){//col:41













return true
}







func (l *lm)CommandLmShowUserModeModule()(ok bool){//col:221































































































return true
}







func (l *lm)CommandLmShowKernelModeModule()(ok bool){//col:332






































































return true
}







func (l *lm)CommandLm()(ok bool){//col:487



















































































































return true
}









