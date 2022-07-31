package debugging-commands
type (
	S interface {
		CommandSearchMemoryHelp() (ok bool)  //col:52
		CommandSearchSendRequest() (ok bool) //col:128
		CommandSearchMemory() (ok bool)      //col:502
	}
)
func NewS() { return &s{} }
func (s *s) CommandSearchMemoryHelp() (ok bool) { //col:52
	return true
}

func (s *s) CommandSearchSendRequest() (ok bool) { //col:128
	return true
}

func (s *s) CommandSearchMemory() (ok bool) { //col:502
	return true
}

