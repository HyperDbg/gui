package debugging-commands
type (
	E interface {
		CommandEditMemoryHelp() (ok bool) //col:49
		CommandEditMemory() (ok bool)     //col:410
	}
)
func NewE() { return &e{} }
func (e *e) CommandEditMemoryHelp() (ok bool) { //col:49
	return true
}

func (e *e) CommandEditMemory() (ok bool) { //col:410
	return true
}

