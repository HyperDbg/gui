package debugging-commands


type (
Bp interface{
CommandBpHelp()(ok bool)//col:44
CommandBp()(ok bool)//col:216
}

)

func NewBp() { return & bp{} }

func (b *bp)CommandBpHelp()(ok bool){//col:44
















return true
}


func (b *bp)CommandBp()(ok bool){//col:216


































































































































return true
}




