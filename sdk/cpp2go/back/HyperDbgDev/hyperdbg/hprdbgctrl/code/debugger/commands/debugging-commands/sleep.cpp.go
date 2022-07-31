package debugging-commands
type (
Sleep interface{
CommandSleepHelp()(ok bool)//col:27
CommandSleep()(ok bool)//col:56
}

)
func NewSleep() { return & sleep{} }
func (s *sleep)CommandSleepHelp()(ok bool){//col:27
return true
}

func (s *sleep)CommandSleep()(ok bool){//col:56
return true
}

