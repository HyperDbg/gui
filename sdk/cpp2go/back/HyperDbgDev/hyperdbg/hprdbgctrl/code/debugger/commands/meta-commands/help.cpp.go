package meta-commands
type (
Help interface{
CommandHelpHelp()(ok bool)//col:30
}

)
func NewHelp() { return & help{} }
func (h *help)CommandHelpHelp()(ok bool){//col:30
return true
}

