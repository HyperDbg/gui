package meta-commands
type (
Logopen interface{
CommandLogopenHelp()(ok bool)//col:33
CommandLogopen()(ok bool)//col:110
LogopenSaveToFile()(ok bool)//col:122
}

)
func NewLogopen() { return & logopen{} }
func (l *logopen)CommandLogopenHelp()(ok bool){//col:33
return true
}

func (l *logopen)CommandLogopen()(ok bool){//col:110
return true
}

func (l *logopen)LogopenSaveToFile()(ok bool){//col:122
return true
}

