package debugging-commands


type (
Bc interface{
CommandBcHelp()(ok bool)//col:34
CommandBc()(ok bool)//col:96
}

)

func NewBc() { return & bc{} }

func (b *bc)CommandBcHelp()(ok bool){//col:34








return true
}


func (b *bc)CommandBc()(ok bool){//col:96




























return true
}




