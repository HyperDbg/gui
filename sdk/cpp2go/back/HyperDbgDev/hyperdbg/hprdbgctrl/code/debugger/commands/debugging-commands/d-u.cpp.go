package debugging-commands
type (
Du interface{
CommandReadMemoryAndDisassemblerHelp()(ok bool)//col:59
CommandReadMemoryAndDisassembler()(ok bool)//col:351
}

)
func NewDU() { return & du{} }
func (d *du)CommandReadMemoryAndDisassemblerHelp()(ok bool){//col:59
return true
}

func (d *du)CommandReadMemoryAndDisassembler()(ok bool){//col:351
return true
}

