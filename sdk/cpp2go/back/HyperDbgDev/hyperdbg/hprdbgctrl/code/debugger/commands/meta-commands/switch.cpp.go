package meta-commands
type (
SwitchA interface{
CommandSwitchHelp()(ok bool)//col:39
CommandSwitch()(ok bool)//col:116
}

)
func NewSwitchA() { return & switchA{} }
func (s *switchA)CommandSwitchHelp()(ok bool){//col:39
return true
}

func (s *switchA)CommandSwitch()(ok bool){//col:116
return true
}

