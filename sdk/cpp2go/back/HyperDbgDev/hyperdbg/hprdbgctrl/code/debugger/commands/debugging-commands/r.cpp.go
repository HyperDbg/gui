package debugging-commands


type (
R interface{
CommandRHelp()(ok bool)//col:164
ShowAllRegisters()(ok bool)//col:179
CommandR()(ok bool)//col:367
}

)

func NewR() { return & r{} }

func (r *r)CommandRHelp()(ok bool){//col:164










return true
}


func (r *r)ShowAllRegisters()(ok bool){//col:179






return true
}


func (r *r)CommandR()(ok bool){//col:367











































































































return true
}




