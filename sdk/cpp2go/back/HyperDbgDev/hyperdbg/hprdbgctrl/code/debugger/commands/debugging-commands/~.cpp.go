package debugging-commands


type (
Unknown interface{
CommandCoreHelp()(ok bool)//col:36
CommandCore()(ok bool)//col:86
}

)

func NewUnknown() { return & unknown{} }

func (u *unknown)CommandCoreHelp()(ok bool){//col:36









return true
}


func (u *unknown)CommandCore()(ok bool){//col:86






























return true
}




