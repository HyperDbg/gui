package extension-commands
type (
Dr interface{
CommandDrHelp()(ok bool)//col:33
CommandDr()(ok bool)//col:122
}

)
func NewDr() { return & dr{} }
func (d *dr)CommandDrHelp()(ok bool){//col:33
return true
}

func (d *dr)CommandDr()(ok bool){//col:122
return true
}

