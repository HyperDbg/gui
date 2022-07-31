package debugging-commands
type (
Flush interface{
CommandFlushHelp()(ok bool)//col:31
CommandFlushRequestFlush()(ok bool)//col:99
CommandFlush()(ok bool)//col:122
}

)
func NewFlush() { return & flush{} }
func (f *flush)CommandFlushHelp()(ok bool){//col:31
return true
}

func (f *flush)CommandFlushRequestFlush()(ok bool){//col:99
return true
}

func (f *flush)CommandFlush()(ok bool){//col:122
return true
}

