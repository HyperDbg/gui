package debugging-commands


type (
Flush interface{
CommandFlushHelp()(ok bool)//col:31
CommandFlushRequestFlush()(ok bool)//col:100
CommandFlush()(ok bool)//col:124
}

)

func NewFlush() { return & flush{} }

func (f *flush)CommandFlushHelp()(ok bool){//col:31






return true
}


func (f *flush)CommandFlushRequestFlush()(ok bool){//col:100
































return true
}


func (f *flush)CommandFlush()(ok bool){//col:124










return true
}




