package extension-commands
type (
SyscallSysret interface{
CommandSyscallHelp()(ok bool)//col:42
CommandSysretHelp()(ok bool)//col:69
CommandSyscallAndSysret()(ok bool)//col:268
}

)
func NewSyscallSysret() { return & syscallSysret{} }
func (s *syscallSysret)CommandSyscallHelp()(ok bool){//col:42
return true
}

func (s *syscallSysret)CommandSysretHelp()(ok bool){//col:69
return true
}

func (s *syscallSysret)CommandSyscallAndSysret()(ok bool){//col:268
return true
}

