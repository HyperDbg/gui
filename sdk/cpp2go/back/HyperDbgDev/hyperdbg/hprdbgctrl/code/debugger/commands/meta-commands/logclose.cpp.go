package meta-commands


type (
Logclose interface{
CommandLogcloseHelp()(ok bool)//col:31
CommandLogclose()(ok bool)//col:78
}






)

func NewLogclose() { return & logclose{} }

func (l *logclose)CommandLogcloseHelp()(ok bool){//col:31





return true
}







func (l *logclose)CommandLogclose()(ok bool){//col:78


























return true
}









