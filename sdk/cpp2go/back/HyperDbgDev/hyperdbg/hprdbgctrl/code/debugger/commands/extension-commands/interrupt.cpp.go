package extension-commands


type (
Interrupt interface{
CommandInterruptHelp()(ok bool)//col:35
CommandInterrupt()(ok bool)//col:196
}

)

func NewInterrupt() { return & interrupt{} }

func (i *interrupt)CommandInterruptHelp()(ok bool){//col:35













return true
}


func (i *interrupt)CommandInterrupt()(ok bool){//col:196





























































































return true
}




