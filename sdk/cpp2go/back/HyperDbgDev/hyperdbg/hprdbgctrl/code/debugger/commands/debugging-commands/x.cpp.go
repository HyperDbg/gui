package debugging-commands


type (
X interface{
CommandXHelp()(ok bool)//col:30
CommandX()(ok bool)//col:69
}

)

func NewX() { return & x{} }

func (x *x)CommandXHelp()(ok bool){//col:30









return true
}


func (x *x)CommandX()(ok bool){//col:69













return true
}




