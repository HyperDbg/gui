package extension-commands


type (
Cpuid interface{
CommandCpuidHelp()(ok bool)//col:34
CommandCpuid()(ok bool)//col:123
}
















)

func NewCpuid() { return & cpuid{} }

func (c *cpuid)CommandCpuidHelp()(ok bool){//col:34













return true
}

















func (c *cpuid)CommandCpuid()(ok bool){//col:123




















































return true
}



















