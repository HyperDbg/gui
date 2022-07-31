package extension-commands
type (
Ioout interface{
CommandIooutHelp()(ok bool)//col:34
CommandIoout()(ok bool)//col:162
}

)
func NewIoout() { return & ioout{} }
func (i *ioout)CommandIooutHelp()(ok bool){//col:34
return true
}

func (i *ioout)CommandIoout()(ok bool){//col:162
return true
}

