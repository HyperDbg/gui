package debugging-commands
type (
	Rdmsr interface {
		CommandRdmsrHelp() (ok bool) //col:29
		CommandRdmsr() (ok bool)     //col:171
	}
)
func NewRdmsr() { return &rdmsr{} }
func (r *rdmsr) CommandRdmsrHelp() (ok bool) { //col:29
	return true
}

func (r *rdmsr) CommandRdmsr() (ok bool) { //col:171
	return true
}

