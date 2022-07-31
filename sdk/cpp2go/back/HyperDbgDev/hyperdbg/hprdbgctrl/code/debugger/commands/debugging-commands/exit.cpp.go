package debugging-commands
type (
	Exit interface {
		CommandExitHelp() (ok bool) //col:31
		CommandExit() (ok bool)     //col:59
	}
)
func NewExit() { return &exit{} }
func (e *exit) CommandExitHelp() (ok bool) { //col:31
	return true
}

func (e *exit) CommandExit() (ok bool) { //col:59
	return true
}

